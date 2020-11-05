use crate::messages::*;
use clarity::{abi::encode_tokens, abi::Token, PrivateKey as EthPrivateKey};
use clarity::{Address as EthAddress, Uint256};
use contact::jsonrpc::error::JsonRpcError;
use contact::types::TXSendResponse;
use contact::{client::Contact, utils::maybe_get_optional_tx_info};
use deep_space::private_key::PrivateKey;
use deep_space::stdfee::StdFee;
use deep_space::stdsignmsg::StdSignMsg;
use deep_space::transaction::TransactionSendType;
use deep_space::{coin::Coin, utils::bytes_to_hex_str};
use peggy_utils::types::*;
use sha3::{Digest, Keccak256};

/// Send a transaction updating the eth address for the sending
/// Cosmos address. The sending Cosmos address should be a validator
pub async fn update_peggy_eth_address(
    contact: &Contact,
    eth_private_key: EthPrivateKey,
    private_key: PrivateKey,
    fee: Coin,
    chain_id: Option<String>,
    account_number: Option<u128>,
    sequence: Option<u128>,
) -> Result<TXSendResponse, JsonRpcError> {
    trace!("Updating Peggy ETH address");
    let our_address = private_key
        .to_public_key()
        .expect("Invalid private key!")
        .to_address();

    let tx_info =
        maybe_get_optional_tx_info(our_address, chain_id, account_number, sequence, &contact)
            .await?;
    trace!("got optional tx info");

    let eth_address = eth_private_key.to_public_key().unwrap();
    let eth_signature = eth_private_key.sign_ethereum_msg(our_address.as_bytes());
    trace!(
        "sig: {} address: {}",
        clarity::utils::bytes_to_hex_str(&eth_signature.to_bytes()),
        clarity::utils::bytes_to_hex_str(eth_address.as_bytes())
    );

    let std_sign_msg = StdSignMsg {
        chain_id: tx_info.chain_id,
        account_number: tx_info.account_number,
        sequence: tx_info.sequence,
        fee: StdFee {
            amount: vec![fee],
            gas: 500_000u64.into(),
        },
        msgs: vec![PeggyMsg::SetEthAddressMsg(SetEthAddressMsg {
            eth_address,
            validator: our_address,
            eth_signature: bytes_to_hex_str(&eth_signature.to_bytes()),
        })],
        memo: String::new(),
    };

    let tx = private_key
        .sign_std_msg(std_sign_msg, TransactionSendType::Block)
        .unwrap();

    contact.retry_on_block(tx).await
}

/// Send a transaction requesting that a valset be formed for a given block
/// height
pub async fn send_valset_request(
    contact: &Contact,
    private_key: PrivateKey,
    fee: Coin,
) -> Result<TXSendResponse, JsonRpcError> {
    let our_address = private_key
        .to_public_key()
        .expect("Invalid private key!")
        .to_address();

    let tx_info = maybe_get_optional_tx_info(our_address, None, None, None, &contact).await?;

    let std_sign_msg = StdSignMsg {
        chain_id: tx_info.chain_id,
        account_number: tx_info.account_number,
        sequence: tx_info.sequence,
        fee: StdFee {
            amount: vec![fee],
            gas: 500_000u64.into(),
        },
        msgs: vec![PeggyMsg::ValsetRequestMsg(ValsetRequestMsg {
            requester: our_address,
        })],
        memo: String::new(),
    };

    let tx = private_key
        .sign_std_msg(std_sign_msg.clone(), TransactionSendType::Block)
        .unwrap();

    contact.retry_on_block(tx).await
}

/// Send in a confirmation for a specific validator set for a specific block height
#[allow(clippy::too_many_arguments)]
pub async fn send_valset_confirm(
    contact: &Contact,
    eth_private_key: EthPrivateKey,
    fee: Coin,
    valset: Valset,
    private_key: PrivateKey,
    peggy_id: String,
) -> Result<TXSendResponse, JsonRpcError> {
    let our_address = private_key
        .to_public_key()
        .expect("Invalid private key!")
        .to_address();
    let our_eth_address = eth_private_key.to_public_key().unwrap();

    let tx_info = maybe_get_optional_tx_info(our_address, None, None, None, contact).await?;

    let (eth_addresses, powers) = valset.filter_empty_addresses();
    let message = encode_tokens(&[
        Token::FixedString(peggy_id),
        Token::FixedString("checkpoint".to_string()),
        valset.nonce.into(),
        eth_addresses.into(),
        powers.into(),
    ]);
    let eth_signature = eth_private_key.sign_ethereum_msg(&message);

    let std_sign_msg = StdSignMsg {
        chain_id: tx_info.chain_id,
        account_number: tx_info.account_number,
        sequence: tx_info.sequence,
        fee: StdFee {
            amount: vec![fee],
            gas: 500_000u64.into(),
        },
        msgs: vec![PeggyMsg::ValsetConfirmMsg(ValsetConfirmMsg {
            validator: our_address,
            eth_address: our_eth_address,
            nonce: valset.nonce.into(),
            eth_signature: bytes_to_hex_str(&eth_signature.to_bytes()),
        })],
        memo: String::new(),
    };

    let tx = private_key
        .sign_std_msg(std_sign_msg, TransactionSendType::Block)
        .unwrap();

    contact.retry_on_block(tx).await
}

/// Send in a confirmation for a specific transaction batch set for a specific block height
/// since transaction batches also include validator sets this has all the arguments
#[allow(clippy::too_many_arguments)]
pub async fn send_batch_confirm(
    contact: &Contact,
    eth_private_key: EthPrivateKey,
    fee: Coin,
    transaction_batch: TransactionBatch,
    private_key: PrivateKey,
    peggy_id: String,
) -> Result<TXSendResponse, JsonRpcError> {
    let our_address = private_key
        .to_public_key()
        .expect("Invalid private key!")
        .to_address();
    let our_eth_address = eth_private_key.to_public_key().unwrap();

    let tx_info = maybe_get_optional_tx_info(our_address, None, None, None, contact).await?;

    // transaction batches include a validator set update, the way this is verified is that the valset checkpoint
    // (encoded ethereum data) is included within the batch signature, which is itself a checkpoint over the batch data
    let valset = transaction_batch.valset.clone();
    let (eth_addresses, powers) = valset.filter_empty_addresses();
    let valset_checkpoint = encode_tokens(&[
        Token::FixedString(peggy_id.clone()),
        Token::FixedString("checkpoint".to_string()),
        valset.nonce.into(),
        eth_addresses.into(),
        powers.into(),
    ]);
    let valset_digest = Keccak256::digest(&valset_checkpoint).as_slice().to_vec();
    let (amounts, destinations, fees) = transaction_batch.get_checkpoint_values();
    let batch_checkpoint = encode_tokens(&[
        Token::FixedString(peggy_id),
        Token::FixedString("valsetAndTransactionBatch".to_string()),
        Token::Bytes(valset_digest),
        amounts,
        destinations,
        fees,
        transaction_batch.nonce.clone().into(),
        transaction_batch.token_contract.into(),
    ]);
    let eth_signature = eth_private_key.sign_ethereum_msg(&batch_checkpoint);

    let std_sign_msg = StdSignMsg {
        chain_id: tx_info.chain_id,
        account_number: tx_info.account_number,
        sequence: tx_info.sequence,
        fee: StdFee {
            amount: vec![fee],
            gas: 500_000u64.into(),
        },
        msgs: vec![PeggyMsg::ConfirmBatchMsg(ConfirmBatchMsg {
            validator: our_address,
            address: our_eth_address,
            nonce: transaction_batch.nonce,
            eth_signature: bytes_to_hex_str(&eth_signature.to_bytes()),
        })],
        memo: String::new(),
    };

    let tx = private_key
        .sign_std_msg(std_sign_msg, TransactionSendType::Block)
        .unwrap();

    contact.retry_on_block(tx).await
}

pub async fn send_ethereum_claims(
    contact: &Contact,
    eth_chain_id: Uint256,
    peggy_contract: EthAddress,
    private_key: PrivateKey,
    claims: Vec<EthereumBridgeClaim>,
    fee: Coin,
) -> Result<TXSendResponse, JsonRpcError> {
    let our_address = private_key
        .to_public_key()
        .expect("Invalid private key!")
        .to_address();

    let tx_info = maybe_get_optional_tx_info(our_address, None, None, None, contact).await?;

    let std_sign_msg = StdSignMsg {
        chain_id: tx_info.chain_id,
        account_number: tx_info.account_number,
        sequence: tx_info.sequence,
        fee: StdFee {
            amount: vec![fee],
            gas: 500_000u64.into(),
        },
        msgs: vec![PeggyMsg::CreateEthereumClaimsMsg(CreateEthereumClaimsMsg {
            ethereum_chain_id: eth_chain_id,
            bridge_contract_address: peggy_contract,
            orchestrator: our_address,
            claims,
        })],
        memo: String::new(),
    };

    let tx = private_key
        .sign_std_msg(std_sign_msg, TransactionSendType::Block)
        .unwrap();

    contact.retry_on_block(tx).await
}

/// Sends tokens from Cosmos to Ethereum. These tokens will not be sent immediately instead
/// they will require some time to be included in a batch
pub async fn send_to_eth(
    private_key: PrivateKey,
    destination: EthAddress,
    amount: Coin,
    fee: Coin,
    contact: &Contact,
) -> Result<TXSendResponse, JsonRpcError> {
    let our_address = private_key
        .to_public_key()
        .expect("Invalid private key!")
        .to_address();
    let tx_info = maybe_get_optional_tx_info(our_address, None, None, None, contact).await?;

    let std_sign_msg = StdSignMsg {
        chain_id: tx_info.chain_id,
        account_number: tx_info.account_number,
        sequence: tx_info.sequence,
        fee: StdFee {
            amount: vec![fee.clone()],
            gas: 500_000u64.into(),
        },
        msgs: vec![PeggyMsg::SendToEthMsg(SendToEthMsg {
            sender: our_address,
            dest_address: destination,
            send: amount,
            bridge_fee: fee,
        })],
        memo: String::new(),
    };

    let tx = private_key
        .sign_std_msg(std_sign_msg, TransactionSendType::Block)
        .unwrap();

    contact.retry_on_block(tx).await
}

pub async fn request_batch(
    private_key: PrivateKey,
    denom: String,
    fee: Coin,
    contact: &Contact,
) -> Result<TXSendResponse, JsonRpcError> {
    let our_address = private_key
        .to_public_key()
        .expect("Invalid private key!")
        .to_address();
    let tx_info = maybe_get_optional_tx_info(our_address, None, None, None, contact).await?;

    let std_sign_msg = StdSignMsg {
        chain_id: tx_info.chain_id,
        account_number: tx_info.account_number,
        sequence: tx_info.sequence,
        fee: StdFee {
            amount: vec![fee.clone()],
            gas: 500_000u64.into(),
        },
        msgs: vec![PeggyMsg::RequestBatchMsg(RequestBatchMsg {
            denom,
            requester: our_address,
        })],
        memo: String::new(),
    };

    let tx = private_key
        .sign_std_msg(std_sign_msg, TransactionSendType::Block)
        .unwrap();

    contact.retry_on_block(tx).await
}
