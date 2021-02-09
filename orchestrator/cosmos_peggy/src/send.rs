use crate::messages::*;
use clarity::Address as EthAddress;
use clarity::PrivateKey as EthPrivateKey;
use contact::jsonrpc::error::JsonRpcError;
use contact::types::TXSendResponse;
use contact::{client::Contact, utils::maybe_get_optional_tx_info};
use deep_space::address::Address;
use deep_space::private_key::PrivateKey;
use deep_space::stdfee::StdFee;
use deep_space::stdsignmsg::StdSignMsg;
use deep_space::transaction::TransactionSendType;
use deep_space::{coin::Coin, utils::bytes_to_hex_str};
use ethereum_peggy::message_signatures::{encode_tx_batch_confirm, encode_valset_confirm};
use peggy_utils::types::*;

/// Send a transaction updating the eth address for the sending
/// Cosmos address. The sending Cosmos address should be a validator
pub async fn update_peggy_delegate_addresses(
    contact: &Contact,
    delegate_eth_address: EthAddress,
    delegate_cosmos_address: Address,
    private_key: PrivateKey,
    fee: Coin,
) -> Result<TXSendResponse, JsonRpcError> {
    trace!("Updating Peggy Delegate addresses");
    let our_valoper_address = private_key
        .to_public_key()
        .expect("Invalid private key!")
        .to_address()
        .to_bech32("cosmosvaloper")
        .unwrap();
    let our_address = private_key
        .to_public_key()
        .expect("Invalid private key!")
        .to_address();

    let tx_info = maybe_get_optional_tx_info(our_address, None, None, None, &contact).await?;
    trace!("got optional tx info");

    let std_sign_msg = StdSignMsg {
        chain_id: tx_info.chain_id,
        account_number: tx_info.account_number,
        sequence: tx_info.sequence,
        fee: StdFee {
            amount: vec![fee],
            gas: 500_000u64.into(),
        },
        msgs: vec![PeggyMsg::SetOrchestratorAddressMsg(
            SetOrchestratorAddressMsg {
                eth_address: delegate_eth_address,
                validator: our_valoper_address,
                orchestrator: delegate_cosmos_address,
            },
        )],
        memo: String::new(),
    };

    let tx = private_key
        .sign_std_msg(std_sign_msg, TransactionSendType::Block)
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

    let message = encode_valset_confirm(peggy_id, valset.clone());
    let eth_signature = eth_private_key.sign_ethereum_msg(&message);

    trace!(
        "Sent valset update with address {} and sig {}",
        our_eth_address,
        bytes_to_hex_str(&eth_signature.to_bytes())
    );
    let std_sign_msg = StdSignMsg {
        chain_id: tx_info.chain_id,
        account_number: tx_info.account_number,
        sequence: tx_info.sequence,
        fee: StdFee {
            amount: vec![fee],
            gas: 500_000u64.into(),
        },
        msgs: vec![PeggyMsg::ValsetConfirmMsg(ValsetConfirmMsg {
            orchestrator: our_address,
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

    let batch_checkpoint = encode_tx_batch_confirm(peggy_id.clone(), transaction_batch.clone());
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
            orchestrator: our_address,
            token_contract: transaction_batch.token_contract,
            eth_signer: our_eth_address,
            nonce: transaction_batch.nonce.into(),
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
    private_key: PrivateKey,
    deposits: Vec<SendToCosmosEvent>,
    withdraws: Vec<TransactionBatchExecutedEvent>,
    erc20_deploys: Vec<ERC20DeployedEvent>,
    fee: Coin,
) -> Result<TXSendResponse, JsonRpcError> {
    let our_address = private_key
        .to_public_key()
        .expect("Invalid private key!")
        .to_address();

    let tx_info = maybe_get_optional_tx_info(our_address, None, None, None, contact).await?;

    let mut msgs = Vec::new();
    for deposit in deposits {
        msgs.push(PeggyMsg::DepositClaimMsg(DepositClaimMsg::from_event(
            deposit,
            our_address,
        )))
    }
    for withdraw in withdraws {
        msgs.push(PeggyMsg::WithdrawClaimMsg(WithdrawClaimMsg::from_event(
            withdraw,
            our_address,
        )))
    }
    for deploy in erc20_deploys {
        msgs.push(PeggyMsg::ERC20DeployedClaimMsg(
            ERC20DeployedClaimMsg::from_event(deploy, our_address),
        ))
    }

    let std_sign_msg = StdSignMsg {
        chain_id: tx_info.chain_id,
        account_number: tx_info.account_number,
        sequence: tx_info.sequence,
        fee: StdFee {
            amount: vec![fee],
            gas: 500_000_000u64.into(),
        },
        msgs,
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
    if amount.denom != fee.denom || !amount.denom.contains("peggy") {
        return Err(JsonRpcError::BadInput(format!(
            "{} {} is an invalid denom set for SendToEth",
            amount.denom, fee.denom,
        )));
    }
    let balances = contact.get_balances(our_address).await.unwrap().result;
    let mut found = false;
    for balance in balances {
        if balance.denom == amount.denom {
            let total_amount = amount.amount.clone() + (fee.amount.clone() * 2u8.into());
            if balance.amount < total_amount {
                return Err(JsonRpcError::BadInput(format!(
                    "Insufficient balance of {} to send {}",
                    amount.denom, total_amount,
                )));
            }
            found = true;
        }
    }
    if !found {
        return Err(JsonRpcError::BadInput(format!(
            "No balance of {} to send",
            amount.denom,
        )));
    }

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
            eth_dest: destination,
            amount,
            bridge_fee: fee,
        })],
        memo: String::new(),
    };

    let tx = private_key
        .sign_std_msg(std_sign_msg, TransactionSendType::Block)
        .unwrap();

    contact.retry_on_block(tx).await
}

pub async fn send_request_batch(
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
            gas: 500_000_000u64.into(),
        },
        msgs: vec![PeggyMsg::RequestBatchMsg(RequestBatchMsg {
            denom,
            orchestrator: our_address,
        })],
        memo: String::new(),
    };

    let tx = private_key
        .sign_std_msg(std_sign_msg, TransactionSendType::Block)
        .unwrap();

    contact.retry_on_block(tx).await
}
