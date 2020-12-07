use clarity::Address as EthAddress;
use deep_space::address::Address;
use peggy_proto::peggy::query_client::QueryClient as PeggyQueryClient;
use peggy_proto::peggy::QueryBatchConfirmsRequest;
use peggy_proto::peggy::QueryLastPendingBatchRequestByAddrRequest;
use peggy_proto::peggy::QueryLastPendingValsetRequestByAddrRequest;
use peggy_proto::peggy::QueryLastValsetRequestsRequest;
use peggy_proto::peggy::QueryOutgoingTxBatchesRequest;
use peggy_proto::peggy::QueryValsetConfirmsByNonceRequest;
use peggy_proto::peggy::QueryValsetRequestRequest;
use peggy_utils::error::PeggyError;
use peggy_utils::types::*;
use tonic::transport::Channel;

/// get the valset for a given nonce (block) height
pub async fn get_valset(
    client: &mut PeggyQueryClient<Channel>,
    nonce: u64,
) -> Result<Option<Valset>, PeggyError> {
    let request = client
        .valset_request(QueryValsetRequestRequest { nonce })
        .await?;
    let valset = request.into_inner().valset;
    let valset = match valset {
        Some(v) => Some(v.into()),
        None => None,
    };
    Ok(valset)
}

/// This hits the /pending_valset_requests endpoint and will provide the oldest
/// validator set we have not yet signed.
pub async fn get_oldest_unsigned_valset(
    client: &mut PeggyQueryClient<Channel>,
    address: Address,
) -> Result<Option<Valset>, PeggyError> {
    let request = client
        .last_pending_valset_request_by_addr(QueryLastPendingValsetRequestByAddrRequest {
            address: address.to_string(),
        })
        .await?;
    let valset = request.into_inner().valset;
    let valset = match valset {
        Some(v) => Some(v.into()),
        None => None,
    };
    Ok(valset)
}

/// this input views the last five valest requests that have been made, useful if you're
/// a relayer looking to ferry confirmations
pub async fn get_latest_valsets(
    client: &mut PeggyQueryClient<Channel>,
) -> Result<Vec<Valset>, PeggyError> {
    let request = client
        .last_valset_requests(QueryLastValsetRequestsRequest {})
        .await?;
    let valsets = request.into_inner().valsets;
    Ok(valsets.iter().map(|v| v.into()).collect())
}

/// get all valset confirmations for a given nonce
pub async fn get_all_valset_confirms(
    client: &mut PeggyQueryClient<Channel>,
    nonce: u64,
) -> Result<Vec<ValsetConfirmResponse>, PeggyError> {
    let request = client
        .valset_confirms_by_nonce(QueryValsetConfirmsByNonceRequest { nonce })
        .await?;
    let confirms = request.into_inner().confirms;
    let mut parsed_confirms = Vec::new();
    for item in confirms {
        parsed_confirms.push(ValsetConfirmResponse::from_proto(item)?)
    }
    Ok(parsed_confirms)
}

pub async fn get_oldest_unsigned_transaction_batch(
    client: &mut PeggyQueryClient<Channel>,
    address: Address,
) -> Result<Option<TransactionBatch>, PeggyError> {
    let request = client
        .last_pending_batch_request_by_addr(QueryLastPendingBatchRequestByAddrRequest {
            address: address.to_string(),
        })
        .await?;
    let batch = request.into_inner().batch;
    match batch {
        Some(batch) => Ok(Some(TransactionBatch::from_proto(batch)?)),
        None => Ok(None),
    }
}

pub async fn get_latest_transaction_batches(
    client: &mut PeggyQueryClient<Channel>,
) -> Result<Vec<TransactionBatch>, PeggyError> {
    let request = client
        .outgoing_tx_batches(QueryOutgoingTxBatchesRequest {})
        .await?;
    let batches = request.into_inner().batches;
    let mut out = Vec::new();
    for batch in batches {
        out.push(TransactionBatch::from_proto(batch)?)
    }
    Ok(out)
}

/// get all batch confirmations for a given nonce and denom
pub async fn get_transaction_batch_signatures(
    client: &mut PeggyQueryClient<Channel>,
    nonce: u64,
    contract_address: EthAddress,
) -> Result<Vec<BatchConfirmResponse>, PeggyError> {
    let request = client
        .batch_confirms(QueryBatchConfirmsRequest {
            nonce,
            contract_address: contract_address.to_string(),
        })
        .await?;
    let batch_confirms = request.into_inner().confirms;
    let mut out = Vec::new();
    for confirm in batch_confirms {
        out.push(BatchConfirmResponse::from_proto(confirm)?)
    }
    Ok(out)
}
