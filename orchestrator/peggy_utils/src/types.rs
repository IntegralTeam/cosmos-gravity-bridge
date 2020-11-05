use std::{cmp::Ordering, collections::HashMap, fmt};

use clarity::Signature as EthSignature;
use clarity::{abi::Token, Address as EthAddress};
use contact::{jsonrpc::error::JsonRpcError, types::parse_val};
use deep_space::address::Address as CosmosAddress;
use num256::Uint256;
use web30::types::Log;

use crate::error::OrchestratorError;

/// the response we get when querying for a valset confirmation
#[derive(Serialize, Deserialize, Debug, Default, Clone)]
pub struct ValsetConfirmResponse {
    #[serde(deserialize_with = "parse_val")]
    pub validator: CosmosAddress,
    #[serde(deserialize_with = "parse_val")]
    pub eth_address: EthAddress,
    #[serde(deserialize_with = "parse_val")]
    pub nonce: Uint256,
    #[serde(deserialize_with = "parse_val", rename = "signature")]
    pub eth_signature: EthSignature,
}

/// a list of validators, powers, and eth addresses at a given block height
#[derive(Serialize, Deserialize, Debug, Default, Clone)]
pub struct Valset {
    pub nonce: u64,
    pub members: Vec<ValsetMember>,
}

impl Valset {
    /// Takes an array of Option<EthAddress> and converts to EthAddress and replaces with zeros
    /// when none is found, Zeros are interpreted by the contract as 'no signature provided' and
    /// signature checks can pass with up to 33% of all voting power presented as zeroed addresses
    pub fn filter_empty_addresses(&self) -> (Vec<EthAddress>, Vec<u64>) {
        let mut addresses = Vec::new();
        let mut powers = Vec::new();
        for val in self.members.iter() {
            match val.eth_address {
                Some(a) => {
                    addresses.push(a);
                    powers.push(val.power);
                }
                None => {
                    addresses.push(EthAddress::default());
                    powers.push(val.power);
                }
            }
        }
        (addresses, powers)
    }

    pub fn get_power(&self, address: EthAddress) -> Result<u64, JsonRpcError> {
        for val in self.members.iter() {
            if val.eth_address == Some(address) {
                return Ok(val.power);
            }
        }
        Err(JsonRpcError::BadInput(
            "All Eth Addresses must be set".to_string(),
        ))
    }

    /// combines the provided signatures with the valset ensuring that ordering and signature data is correct
    pub fn order_valset_sigs(
        &self,
        signatures: &[ValsetConfirmResponse],
    ) -> Result<Vec<PeggySignature>, JsonRpcError> {
        let mut out = Vec::new();
        let mut members = HashMap::new();
        for member in self.members.iter() {
            if let Some(address) = member.eth_address {
                members.insert(address, member);
            } else {
                return Err(JsonRpcError::BadInput(
                    "All Eth Addresses must be set".to_string(),
                ));
            }
        }
        for sig in signatures {
            if let Some(val) = members.get(&sig.eth_address) {
                out.push(PeggySignature {
                    power: val.power,
                    eth_address: sig.eth_address,
                    v: sig.eth_signature.v.clone(),
                    r: sig.eth_signature.r.clone(),
                    s: sig.eth_signature.s.clone(),
                })
            } else {
                return Err(JsonRpcError::BadInput(format!(
                    "No Match for sig! {} and {}",
                    sig.eth_address,
                    ValsetMember::display_vec(&self.members)
                )));
            }
        }
        // sort by power so that it is accepted by the contract
        out.sort();
        // go code sorts descending, rust sorts ascending, annoying
        out.reverse();

        Ok(out)
    }

    /// combines the provided signatures with the valset ensuring that ordering and signature data is correct
    pub fn order_batch_sigs(
        &self,
        signed_batch: SignedTransactionBatch,
    ) -> Result<Vec<PeggySignature>, JsonRpcError> {
        let mut out = Vec::new();
        let mut members = HashMap::new();
        for member in self.members.iter() {
            if let Some(address) = member.eth_address {
                members.insert(address, member);
            } else {
                return Err(JsonRpcError::BadInput(
                    "All Eth Addresses must be set".to_string(),
                ));
            }
        }
        for sig in signed_batch.signatures {
            if let Some(val) = members.get(&sig.eth_address) {
                out.push(PeggySignature {
                    power: val.power,
                    eth_address: sig.eth_address,
                    v: sig.eth_signature.v.clone(),
                    r: sig.eth_signature.r.clone(),
                    s: sig.eth_signature.s.clone(),
                })
            } else {
                return Err(JsonRpcError::BadInput(format!(
                    "No Match for sig! {} and {}",
                    sig.eth_address,
                    ValsetMember::display_vec(&self.members)
                )));
            }
        }
        // sort by power so that it is accepted by the contract
        out.sort();
        // go code sorts descending, rust sorts ascending, annoying
        out.reverse();

        Ok(out)
    }
}

/// A sortable struct of a validator and it's signatures
/// this can be used for either transaction batch or validator
/// set signatures
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct PeggySignature {
    // ord sorts on the first member first, so this produces the correct sorting
    power: u64,
    eth_address: EthAddress,
    v: Uint256,
    r: Uint256,
    s: Uint256,
}

impl Ord for PeggySignature {
    // Alex wrote the Go sorting implementation for validator
    // sets as Greatest to Least, now this isn't the convention
    // for any standard sorting implementation and Rust doesn't
    // really like it when you implement sort yourself. It prefers
    // Ord. So here we implement Ord with the Eth address sorting
    // reversed, since they are also sorted greatest to least in
    // the Cosmos module. Then we can call .sort and .reverse and get
    // the same sorting as the Cosmos module.
    fn cmp(&self, other: &Self) -> Ordering {
        if self.power != other.power {
            self.power.cmp(&other.power)
        } else {
            self.eth_address.cmp(&other.eth_address).reverse()
        }
    }
}

impl PartialOrd for PeggySignature {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

/// signatures in array formats ready to be
/// submitted to the Peggy Ethereum Contract
pub struct PeggySignatureArrays {
    pub addresses: Vec<EthAddress>,
    pub powers: Vec<u64>,
    pub v: Token,
    pub r: Token,
    pub s: Token,
}

/// This function handles converting the PeggySignature type into an Ethereum
/// submittable arrays, including the finicky token encoding tricks you need to
/// perform in order to distinguish between a uint8[] and bytes32[]
pub fn to_arrays(input: Vec<PeggySignature>) -> PeggySignatureArrays {
    let mut addresses = Vec::new();
    let mut powers = Vec::new();
    let mut v = Vec::new();
    let mut r = Vec::new();
    let mut s = Vec::new();
    for val in input {
        addresses.push(val.eth_address);
        powers.push(val.power);
        v.push(val.v);
        r.push(Token::Bytes(val.r.to_bytes_be()));
        s.push(Token::Bytes(val.s.to_bytes_be()));
    }
    PeggySignatureArrays {
        addresses,
        powers,
        v: v.into(),
        r: Token::Dynamic(r),
        s: Token::Dynamic(s),
    }
}

/// a list of validators, powers, and eth addresses at a given block height
#[derive(Serialize, Deserialize, Debug, Default, Clone, Eq, PartialEq)]
pub struct ValsetMember {
    // ord sorts on the first member first, so this produces the correct sorting
    power: u64,
    eth_address: Option<EthAddress>,
}

impl Ord for ValsetMember {
    // Alex wrote the Go sorting implementation for validator
    // sets as Greatest to Least, now this isn't the convention
    // for any standard sorting implementation and Rust doesn't
    // really like it when you implement sort yourself. It prefers
    // Ord. So here we implement Ord with the Eth address sorting
    // reversed, since they are also sorted greatest to least in
    // the Cosmos module. Then we can call .sort and .reverse and get
    // the same sorting as the Cosmos module.
    fn cmp(&self, other: &Self) -> Ordering {
        if self.power != other.power {
            self.power.cmp(&other.power)
        } else {
            self.eth_address.cmp(&other.eth_address).reverse()
        }
    }
}

impl PartialOrd for ValsetMember {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl ValsetMember {
    fn display_vec(input: &[ValsetMember]) -> String {
        let mut out = String::new();
        for val in input.iter() {
            out += &val.to_string()
        }
        out
    }
}

impl fmt::Display for ValsetMember {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self.eth_address {
            Some(a) => write!(f, "Address: {} Power: {}", a, self.power),
            None => write!(f, "Address: None Power: {}", self.power),
        }
    }
}

/// a list of validators, powers, and eth addresses at a given block height
#[derive(Serialize, Deserialize, Debug, Default, Clone)]
pub struct ValsetMemberUnparsed {
    ethereum_address: String,
    #[serde(deserialize_with = "parse_val")]
    power: u64,
}

/// a list of validators, powers, and eth addresses at a given block height
/// this version is used by the endpoint to get the data and is then processed
/// by "convert" into ValsetResponse. Making this struct purely internal
#[derive(Serialize, Deserialize, Debug, Default, Clone)]
pub struct ValsetUnparsed {
    #[serde(deserialize_with = "parse_val")]
    nonce: u64,
    members: Vec<ValsetMemberUnparsed>,
}

impl ValsetUnparsed {
    pub fn convert(self) -> Valset {
        let mut out = Vec::new();
        for member in self.members {
            if member.ethereum_address.is_empty() {
                out.push(ValsetMember {
                    power: member.power,
                    eth_address: None,
                });
            } else {
                match member.ethereum_address.parse() {
                    Ok(val) => out.push(ValsetMember {
                        power: member.power,
                        eth_address: Some(val),
                    }),
                    Err(_e) => out.push(ValsetMember {
                        power: member.power,
                        eth_address: None,
                    }),
                }
            }
        }
        Valset {
            nonce: self.nonce,
            members: out,
        }
    }
}

/// the query struct required to get the valset request sent by a specific
/// validator. This is required because the url encoded get methods don't
/// parse addresses well. So there's no way to get an individual validators
/// address without sending over a json body
#[derive(Serialize, Deserialize, Debug, Default, Clone)]
pub struct QueryValsetConfirm {
    pub nonce: String,
    pub address: String,
}

/// A parsed struct representing the Ethereum event fired by the Peggy contract
/// when the validator set is updated.
pub struct ValsetUpdatedEvent {
    pub nonce: Uint256,
    // we currently don't parse members, but the data is there
    //members: Vec<ValsetMember>,
}

impl ValsetUpdatedEvent {
    pub fn from_log(input: &Log) -> Result<ValsetUpdatedEvent, OrchestratorError> {
        // we have one indexed event so we should fine two indexes, one the event itself
        // and one the indexed nonce
        if let Some(nonce_data) = input.topics.get(1) {
            let nonce = Uint256::from_bytes_be(nonce_data);
            Ok(ValsetUpdatedEvent { nonce })
        } else {
            Err(OrchestratorError::InvalidEventLogError(
                "Too few topics".to_string(),
            ))
        }
    }
    pub fn from_logs(input: &[Log]) -> Result<Vec<ValsetUpdatedEvent>, OrchestratorError> {
        let mut res = Vec::new();
        for item in input {
            res.push(ValsetUpdatedEvent::from_log(item)?);
        }
        Ok(res)
    }
}

/// A parsed struct representing the Ethereum event fired by the Peggy contract when
/// a transaction batch is executed.
pub struct TransactionBatchExecutedEvent {
    pub nonce: Uint256,
    /// The ERC20 token contract address for the batch executed, since batches are uniform
    /// in token type there is only one
    pub erc20: EthAddress,
}

impl TransactionBatchExecutedEvent {
    pub fn from_log(input: &Log) -> Result<TransactionBatchExecutedEvent, OrchestratorError> {
        if let (Some(nonce_data), Some(erc20_data)) = (input.topics.get(1), input.topics.get(2)) {
            let nonce = Uint256::from_bytes_be(nonce_data);
            let erc20 = EthAddress::from_slice(&erc20_data)?;
            Ok(TransactionBatchExecutedEvent { nonce, erc20 })
        } else {
            Err(OrchestratorError::InvalidEventLogError(
                "Too few topics".to_string(),
            ))
        }
    }
    pub fn from_logs(
        input: &[Log],
    ) -> Result<Vec<TransactionBatchExecutedEvent>, OrchestratorError> {
        let mut res = Vec::new();
        for item in input {
            res.push(TransactionBatchExecutedEvent::from_log(item)?);
        }
        Ok(res)
    }
}

/// A parsed struct representing the Ethereum event fired when someone makes a deposit
/// on the Peggy contract
pub struct SendToCosmosEvent {
    /// The token contract address for the deposit
    pub erc20: EthAddress,
    /// The Ethereum Sender
    pub sender: EthAddress,
    /// The Cosmos destination
    pub destination: CosmosAddress,
    /// The amount of the erc20 token that is being sent
    pub amount: Uint256,
}

impl SendToCosmosEvent {
    pub fn from_log(input: &Log) -> Result<SendToCosmosEvent, OrchestratorError> {
        let topics = (
            input.topics.get(1),
            input.topics.get(2),
            input.topics.get(3),
        );
        if let (Some(erc20_data), Some(sender_data), Some(destination_data)) = topics {
            let erc20 = EthAddress::from_slice(&erc20_data[12..32])?;
            let sender = EthAddress::from_slice(&sender_data[12..32])?;
            let mut c_address_bytes: [u8; 20] = [0; 20];
            // this is little endian encoded
            c_address_bytes.copy_from_slice(&destination_data[0..20]);
            let destination = CosmosAddress::from_bytes(c_address_bytes);
            let amount = Uint256::from_bytes_be(&input.data);
            Ok(SendToCosmosEvent {
                erc20,
                sender,
                destination,
                amount,
            })
        } else {
            Err(OrchestratorError::InvalidEventLogError(
                "Too few topics".to_string(),
            ))
        }
    }
    pub fn from_logs(input: &[Log]) -> Result<Vec<SendToCosmosEvent>, OrchestratorError> {
        let mut res = Vec::new();
        for item in input {
            res.push(SendToCosmosEvent::from_log(item)?);
        }
        Ok(res)
    }
}

/// This represents an individual transaction being bridged over to Ethereum
/// parallel is the OutgoingTransferTx in x/peggy/types/batch.go
#[derive(Serialize, Deserialize, Debug, Default, Clone)]
pub struct BatchTransaction {
    pub txid: Uint256,
    pub sender: CosmosAddress,
    pub destination: EthAddress,
    pub amount: ERC20Token,
    pub bridge_fee: ERC20Token,
}

/// the response we get when querying for a valset confirmation
#[derive(Serialize, Deserialize, Debug, Default, Clone)]
pub struct TransactionBatch {
    pub nonce: Uint256,
    pub elements: Vec<BatchTransaction>,
    pub created_at: String,
    pub total_fee: ERC20Token,
    pub bridged_denominator: String,
    pub batch_status: String,
    pub valset: Valset,
    pub token_contract: EthAddress,
}

impl TransactionBatch {
    /// extracts the amounts, destinations and fees as submitted to the Ethereum contract
    /// and used for signatures
    pub fn get_checkpoint_values(&self) -> (Token, Token, Token) {
        let mut amounts = Vec::new();
        let mut destinations = Vec::new();
        let mut fees = Vec::new();
        for item in self.elements.iter() {
            amounts.push(Token::Bytes(item.amount.amount.clone().to_bytes_be()));
            fees.push(Token::Bytes(item.bridge_fee.amount.clone().to_bytes_be()));
            destinations.push(item.destination)
        }
        (
            Token::Dynamic(amounts),
            destinations.into(),
            Token::Dynamic(fees),
        )
    }
}

/// the response we get when querying for a valset confirmation
#[derive(Serialize, Deserialize, Debug, Default, Clone)]
pub struct SignedTransactionBatch {
    pub batch: TransactionBatch,
    pub signatures: Vec<SigWithAddress>,
}

#[derive(Serialize, Deserialize, Debug, Default, Clone, Eq, PartialEq, Hash)]
pub struct SigWithAddress {
    eth_address: EthAddress,
    eth_signature: EthSignature,
}

#[derive(Serialize, Deserialize, Debug, Default, Clone, Eq, PartialEq, Hash)]
pub struct ERC20Token {
    pub amount: Uint256,
    pub symbol: String,
    #[serde(rename = "token_contract_address")]
    pub token_contract_address: EthAddress,
}

#[cfg(test)]
mod tests {
    use super::*;
    use rand::seq::SliceRandom;
    use rand::thread_rng;

    #[test]
    fn test_valset_sort() {
        let correct: [PeggySignature; 8] = [
            PeggySignature {
                power: 685294939,
                eth_address: "0x479FFc856Cdfa0f5D1AE6Fa61915b01351A7773D"
                    .parse()
                    .unwrap(),
                v: 0u64.into(),
                r: 0u64.into(),
                s: 0u64.into(),
            },
            PeggySignature {
                power: 678509841,
                eth_address: "0x6db48cBBCeD754bDc760720e38E456144e83269b"
                    .parse()
                    .unwrap(),
                v: 0u64.into(),
                r: 0u64.into(),
                s: 0u64.into(),
            },
            PeggySignature {
                power: 671724742,
                eth_address: "0x0A7254b318dd742A3086882321C27779B4B642a6"
                    .parse()
                    .unwrap(),
                v: 0u64.into(),
                r: 0u64.into(),
                s: 0u64.into(),
            },
            PeggySignature {
                power: 671724742,
                eth_address: "0x454330deAaB759468065d08F2b3B0562caBe1dD1"
                    .parse()
                    .unwrap(),
                v: 0u64.into(),
                r: 0u64.into(),
                s: 0u64.into(),
            },
            PeggySignature {
                power: 671724742,
                eth_address: "0x8E91960d704Df3fF24ECAb78AB9df1B5D9144140"
                    .parse()
                    .unwrap(),
                v: 0u64.into(),
                r: 0u64.into(),
                s: 0u64.into(),
            },
            PeggySignature {
                power: 617443955,
                eth_address: "0x3511A211A6759d48d107898302042d1301187BA9"
                    .parse()
                    .unwrap(),
                v: 0u64.into(),
                r: 0u64.into(),
                s: 0u64.into(),
            },
            PeggySignature {
                power: 291759231,
                eth_address: "0xF14879a175A2F1cEFC7c616f35b6d9c2b0Fd8326"
                    .parse()
                    .unwrap(),
                v: 0u64.into(),
                r: 0u64.into(),
                s: 0u64.into(),
            },
            PeggySignature {
                power: 6785098,
                eth_address: "0x37A0603dA2ff6377E5C7f75698dabA8EE4Ba97B8"
                    .parse()
                    .unwrap(),
                v: 0u64.into(),
                r: 0u64.into(),
                s: 0u64.into(),
            },
        ];
        let mut rng = thread_rng();
        let mut incorrect = correct.clone();

        incorrect.shuffle(&mut rng);
        assert_ne!(incorrect, correct);

        incorrect.sort();
        incorrect.reverse();
        assert_eq!(incorrect, correct);
    }
}
