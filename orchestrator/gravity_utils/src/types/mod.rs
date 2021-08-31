use clarity::Address as EthAddress;
use num256::Uint256;
use std::convert::TryFrom;
mod batches;
mod config;
mod ethereum_events;
pub mod event_signatures;
mod logic_call;
mod signatures;
mod valsets;
use crate::error::GravityError;

pub use batches::*;
pub use config::*;
pub use ethereum_events::*;
pub use logic_call::*;
pub use signatures::*;
pub use valsets::*;

#[derive(Serialize, Deserialize, Debug, Default, Clone, Eq, PartialEq, Hash)]
pub struct Erc20Token {
    pub amount: Uint256,
    #[serde(rename = "contract")]
    pub token_contract_address: EthAddress,
}

impl TryFrom<gravity_proto::gravity::Erc20Token> for Erc20Token {
    type Error = GravityError;
    fn try_from(input: gravity_proto::gravity::Erc20Token) -> Result<Erc20Token, GravityError> {
        Ok(Erc20Token {
            amount: input.amount.parse()?,
            token_contract_address: input.contract.parse()?,
        })
    }
}
