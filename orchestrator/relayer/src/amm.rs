use clarity::Address;
use lazy_static::lazy_static;

// Addresses for Uniswap V3
lazy_static! {
    /// Uniswap V3's Quoter interface for checking current swap prices, from prod Ethereum
    pub static ref UNISWAP_QUOTER_ADDRESS: Address =
        Address::parse_and_validate("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6").unwrap();
    /// Uniswap V3's Router interface for swapping tokens, from prod Ethereum
    pub static ref UNISWAP_ROUTER_ADDRESS: Address =
        Address::parse_and_validate("0xE592427A0AEce92De3Edee1F18E0157C05861564").unwrap();
    /// The DAI V2 Token's address, on prod Ethereum
    pub static ref DAI_CONTRACT_ADDRESS: Address =
        Address::parse_and_validate("0x6B175474E89094C44Da98b954EedeAC495271d0F").unwrap();
    /// The Wrapped Ether's address, on prod Ethereum
    pub static ref WETH_CONTRACT_ADDRESS: Address =
        Address::parse_and_validate("0xc778417E063141139Fce010982780140Aa0cD5Ab").unwrap();
}
