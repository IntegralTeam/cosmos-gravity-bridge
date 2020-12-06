/// Attestation is an aggregate of `claims` that eventually becomes `observed` by all orchestrators
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Attestation {
    /// ClaimType is the type of attestation either WithdrawBatch or BridgeDepoist
    /// TODO: we could remove ClaimType and infer the type of Claim from
    /// the details field and type assertions
    #[prost(enumeration="ClaimType", tag="1")]
    pub claim_type: i32,
    /// EventNonce is the nonce from/for...?
    #[prost(uint64, tag="2")]
    pub event_nonce: u64,
    /// Observed says...?
    #[prost(bool, tag="3")]
    pub observed: bool,
    /// Validator votes for the attestation
    #[prost(string, repeated, tag="4")]
    pub votes: ::std::vec::Vec<std::string::String>,
    /// The attestation details interface type
    #[prost(message, optional, tag="5")]
    pub details: ::std::option::Option<::prost_types::Any>,
}
/// WithdrawalBatch is an attestation detail that marks a batch of outgoing transactions executed and
/// frees earlier unexecuted batches
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct WithdrawalBatch {
    /// BatchNonce is the nonce of the batch on Peggy
    #[prost(uint64, tag="1")]
    pub batch_nonce: u64,
    /// The ERC20 token being sent back to ETH
    #[prost(message, optional, tag="2")]
    pub erc20_token: ::std::option::Option<Erc20Token>,
}
/// // BridgeDeposit is an attestation detail that adds vouchers to an account when executed
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BridgeDeposit {
    /// The ERC20 token being sent to Peggy
    #[prost(message, optional, tag="1")]
    pub erc20_token: ::std::option::Option<Erc20Token>,
    /// The address on ETH that sent the transaction
    #[prost(string, tag="2")]
    pub ethereum_sender: std::string::String,
    /// The address on Peggy recieving the funds
    #[prost(string, tag="3")]
    pub cosmos_receiver: std::string::String,
}
/// ERC20Token unique identifier for an Ethereum ERC20 token.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Erc20Token {
    /// The amount of the ERC20 token 
    #[prost(string, tag="1")]
    pub amount: std::string::String,
    /// The contract address on ETH of the token (note: developers should look up the token symbol using the address on ETH to display for UI)
    #[prost(string, tag="2")]
    pub contract: std::string::String,
}
/// ClaimType is the cosmos type of an event from the counterpart chain that can be handled
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum ClaimType {
    Unknown = 0,
    EthereumBridgeDeposit = 1,
    EthereumBridgeWithdrawalBatch = 2,
}
/// OutgoingTxBatch represents a batch of transactions going from Peggy to ETH
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OutgoingTxBatch {
    /// The nonce of of the individual batch
    #[prost(uint64, tag="1")]
    pub batch_nonce: u64,
    /// The individual operation of the transaction batch
    #[prost(message, repeated, tag="2")]
    pub transactions: ::std::vec::Vec<OutgoingTransferTx>,
    /// The token contract on ETH for amounts in transactions
    #[prost(string, tag="4")]
    pub token_contract: std::string::String,
}
/// OutgoingTransferTx represents an individual send from Peggy to ETH
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OutgoingTransferTx {
    /// The nonce on the peggy side for this individual transfer
    #[prost(uint64, tag="1")]
    pub id: u64,
    /// The peggy address of the sender of this transaction
    #[prost(string, tag="2")]
    pub sender: std::string::String,
    /// The address on ETH where the transfer is bound
    #[prost(string, tag="3")]
    pub dest_address: std::string::String,
    /// The ERC20 token amount to be sent back to ETH
    #[prost(message, optional, tag="4")]
    pub erc20_token: ::std::option::Option<Erc20Token>,
    /// The ERC20 fee paid to the bridge
    #[prost(message, optional, tag="5")]
    pub erc20_fee: ::std::option::Option<Erc20Token>,
}
/// SignType defines messages that have been signed by an orchestrator
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum SignType {
    Unknown = 0,
    OrchestratorSignedMultiSigUpdate = 1,
    OrchestratorSignedWithdrawBatch = 2,
}
/// Params represent the peggy genesis and store parameters
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Params {
    /// PeggyID is a random 32 byte value to prevent signature reuse
    #[prost(string, tag="1")]
    pub peggy_id: std::string::String,
    /// ContractHash is the code hash of a known good version of the Peggy contract solidity code.
    /// It will be used to verify exactly which version of the bridge will be deployed.
    #[prost(string, tag="2")]
    pub contract_source_hash: std::string::String,
    /// StartThreshold is the percentage of total voting power that must be online and participating in
    /// Peggy operations before a bridge can start operating
    #[prost(uint64, tag="3")]
    pub start_threshold: u64,
    /// BridgeContractAddress is address of the bridge contract on the Ethereum side
    #[prost(string, tag="4")]
    pub ethereum_address: std::string::String,
    /// BridgeChainID is the unique identifier of the Ethereum chain
    #[prost(uint64, tag="5")]
    pub bridge_chain_id: u64,
}
/// GenesisState struct
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GenesisState {
    /// The above params are held in GenesisState
    #[prost(message, optional, tag="1")]
    pub params: ::std::option::Option<Params>,
}
/// MsgValsetConfirm
/// this is the message sent by the validators when they wish to submit their
/// signatures over the validator set at a given block height. A validator must
/// first call MsgSetEthAddress to set their Ethereum address to be used for
/// signing. Then someone (anyone) must make a ValsetRequest the request is
/// essentially a messaging mechanism to determine which block all validators
/// should submit signatures over. Finally validators sign the validator set,
/// powers, and Ethereum addresses of the entire validator set at the height of a
/// ValsetRequest and submit that signature with this message.
///
/// If a sufficient number of validators (66% of voting power) (A) have set
/// Ethereum addresses and (B) submit ValsetConfirm messages with their
/// signatures it is then possible for anyone to view these signatures in the
/// chain store and submit them to Ethereum to update the validator set
/// -------------
/// deprecated should use MsgBridgeSignatureSubmission instead
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgValsetConfirm {
    #[prost(uint64, tag="1")]
    pub nonce: u64,
    #[prost(string, tag="2")]
    pub validator: std::string::String,
    #[prost(string, tag="3")]
    pub eth_address: std::string::String,
    #[prost(string, tag="4")]
    pub signature: std::string::String,
}
/// ValsetRequest
/// This message starts off the validator set update process by coordinating a
/// block height around which signatures over the validators, powers, and
/// ethereum addresses will be made and submitted using a ValsetConfirm. Anyone
/// can send this message as it is not authenticated except as a valid tx. In
/// theory people could spam it and the validators will have to determine which
/// block to actually coordinate around by looking over the valset requests and
/// seeing which one some other validator has already submitted a ValsetResponse
/// for.
/// -------------
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgValsetRequest {
    #[prost(string, tag="1")]
    pub requester: std::string::String,
}
/// SetEthAddress
/// This is used by the validators to set the Ethereum address that represents
/// them on the Ethereum side of the bridge. They must sign their Cosmos address
/// using the Ethereum address they have submitted. Like ValsetResponse this
/// message can in theory be submitted by anyone, but only the current validator
/// sets submissions carry any weight.
/// -------------
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetEthAddress {
    #[prost(string, tag="1")]
    pub address: std::string::String,
    #[prost(string, tag="2")]
    pub validator: std::string::String,
    #[prost(string, tag="3")]
    pub signature: std::string::String,
}
/// MsgSendToEth
/// This is the message that a user calls when they want to bridge an asset
/// TODO right now this needs to be locked to a single ERC20
/// TODO fixed fee amounts for now, variable fee amounts in the fee field later
/// TODO actually remove amounts form the users bank balances
/// TODO this message modifies the on chain store by adding itself to a txpool
/// it will later be removed when it is included in a batch and successfully
/// submitted tokens are removed from the users balance immediately
/// -------------
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSendToEth {
    /// the source address on Cosmos
    #[prost(string, tag="1")]
    pub sender: std::string::String,
    /// the destination address on Ethereum
    #[prost(string, tag="2")]
    pub eth_dest: std::string::String,
    /// the coin to send across the bridge, note the restriction that this is a
    /// single coin not a set of coins that is normal in other Cosmos messages
    #[prost(message, optional, tag="3")]
    pub amount: ::std::option::Option<cosmos_sdk_proto::cosmos::base::v1beta1::Coin>,
    /// the fee paid for the bridge, distinct from the fee paid to the chain to
    /// actually send this message in the first place. So a successful send has
    /// two layers of fees for the user
    #[prost(message, optional, tag="4")]
    pub bridge_fee: ::std::option::Option<cosmos_sdk_proto::cosmos::base::v1beta1::Coin>,
}
/// MsgRequestBatch
/// this is a message anyone can send that requests a batch of transactions to
/// send across the bridge be created for whatever block height this message is
/// included in. This acts as a coordination point, the handler for this message
/// looks at the AddToOutgoingPool tx's in the store and generates a batch, also
/// available in the store tied to this message. The validators then grab this
/// batch, sign it, submit the signatures with a MsgConfirmBatch before a relayer
/// can finally submit the batch
/// -------------
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgRequestBatch {
    #[prost(string, tag="1")]
    pub requester: std::string::String,
    #[prost(string, tag="2")]
    pub denom: std::string::String,
}
/// MsgConfirmBatch
/// When validators observe a MsgRequestBatch they form a batch by ordering
/// transactions currently in the txqueue in order of highest to lowest fee,
/// cutting off when the batch either reaches a hardcoded maximum size (to be
/// decided, probably around 100) or when transactions stop being profitable
/// (TODO determine this without nondeterminism) This message includes the batch
/// as well as an Ethereum signature over this batch by the validator
/// -------------
/// deprecated should use MsgBridgeSignatureSubmission instead
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgConfirmBatch {
    #[prost(uint64, tag="1")]
    pub nonce: u64,
    #[prost(string, tag="2")]
    pub token_contract: std::string::String,
    #[prost(string, tag="3")]
    pub eth_signer: std::string::String,
    #[prost(string, tag="4")]
    pub validator: std::string::String,
    #[prost(string, tag="5")]
    pub signature: std::string::String,
}
/// EthereumBridgeDepositClaim
/// When more than 66% of the active validator set has
/// claimed to have seen the deposit enter the ethereum blockchain coins are
/// issued to the Cosmos address in question
/// -------------
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EthereumBridgeDepositClaim {
    #[prost(uint64, tag="1")]
    pub nonce: u64,
    #[prost(message, optional, tag="2")]
    pub erc20_token: ::std::option::Option<Erc20Token>,
    #[prost(string, tag="3")]
    pub ethereum_sender: std::string::String,
    #[prost(string, tag="4")]
    pub cosmos_receiver: std::string::String,
}
/// EthereumBridgeWithdrawalBatchClaim claims that a batch of withdrawal
/// operations on the bridge contract was executed.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct EthereumBridgeWithdrawalBatchClaim {
    #[prost(uint64, tag="1")]
    pub event_nonce: u64,
    #[prost(uint64, tag="2")]
    pub batch_nonce: u64,
}
/// MsgCreateEthereumClaims
/// this message essentially acts as the oracle between Ethereum and Cosmos, when
/// an orchestrator sees that a batch/ deposit/ multisig set update has been
/// submitted on to the Ethereum blockchain they will submit this message which
/// acts as their oracle attestation. When more than 66% of the active validator
/// set has claimed to have seen the transaction enter the ethereum blockchain it
/// is "observed" and state transitions and operations are triggered on the
/// cosmos side.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgCreateEthereumClaims {
    #[prost(uint64, tag="1")]
    pub ethereum_chain_id: u64,
    /// do we need to specify contract address? can we get this from the store?
    #[prost(string, tag="2")]
    pub bridge_contract_address: std::string::String,
    #[prost(string, tag="3")]
    pub orchestrator: std::string::String,
    #[prost(message, repeated, tag="4")]
    pub claims: ::std::vec::Vec<::prost_types::Any>,
}
/// MsgBridgeSignatureSubmission submits the Ethereum signature for a given nonce
/// an claim type.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgBridgeSignatureSubmission {
    #[prost(uint64, tag="1")]
    pub nonce: u64,
    #[prost(enumeration="SignType", tag="2")]
    pub sign_type: i32,
    #[prost(string, tag="3")]
    pub orchestrator: std::string::String,
    #[prost(string, tag="4")]
    pub ethereum_signature: std::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgValsetConfirmResponse {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgValsetRequestResponse {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetEthAddressResponse {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSendToEthResponse {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgRequestBatchResponse {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgConfirmBatchResponse {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgCreateEthereumClaimsResponse {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgBridgeSignatureSubmissionResponse {
}
# [doc = r" Generated client implementations."] pub mod msg_client { # ! [allow (unused_variables , dead_code , missing_docs)] use tonic :: codegen :: * ; pub struct MsgClient < T > { inner : tonic :: client :: Grpc < T > , } impl MsgClient < tonic :: transport :: Channel > { # [doc = r" Attempt to create a new client by connecting to a given endpoint."] pub async fn connect < D > (dst : D) -> Result < Self , tonic :: transport :: Error > where D : std :: convert :: TryInto < tonic :: transport :: Endpoint > , D :: Error : Into < StdError > , { let conn = tonic :: transport :: Endpoint :: new (dst) ? . connect () . await ? ; Ok (Self :: new (conn)) } } impl < T > MsgClient < T > where T : tonic :: client :: GrpcService < tonic :: body :: BoxBody > , T :: ResponseBody : Body + HttpBody + Send + 'static , T :: Error : Into < StdError > , < T :: ResponseBody as HttpBody > :: Error : Into < StdError > + Send , { pub fn new (inner : T) -> Self { let inner = tonic :: client :: Grpc :: new (inner) ; Self { inner } } pub fn with_interceptor (inner : T , interceptor : impl Into < tonic :: Interceptor >) -> Self { let inner = tonic :: client :: Grpc :: with_interceptor (inner , interceptor) ; Self { inner } } pub async fn valset_confirm (& mut self , request : impl tonic :: IntoRequest < super :: MsgValsetConfirm > ,) -> Result < tonic :: Response < super :: MsgValsetConfirmResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Msg/ValsetConfirm") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn valset_request (& mut self , request : impl tonic :: IntoRequest < super :: MsgValsetRequest > ,) -> Result < tonic :: Response < super :: MsgValsetRequestResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Msg/ValsetRequest") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn set_eth_address (& mut self , request : impl tonic :: IntoRequest < super :: MsgSetEthAddress > ,) -> Result < tonic :: Response < super :: MsgSetEthAddressResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Msg/SetEthAddress") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn send_to_eth (& mut self , request : impl tonic :: IntoRequest < super :: MsgSendToEth > ,) -> Result < tonic :: Response < super :: MsgSendToEthResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Msg/SendToEth") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn request_batch (& mut self , request : impl tonic :: IntoRequest < super :: MsgRequestBatch > ,) -> Result < tonic :: Response < super :: MsgRequestBatchResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Msg/RequestBatch") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn confirm_batch (& mut self , request : impl tonic :: IntoRequest < super :: MsgConfirmBatch > ,) -> Result < tonic :: Response < super :: MsgConfirmBatchResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Msg/ConfirmBatch") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn create_ethereum_claims (& mut self , request : impl tonic :: IntoRequest < super :: MsgCreateEthereumClaims > ,) -> Result < tonic :: Response < super :: MsgCreateEthereumClaimsResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Msg/CreateEthereumClaims") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn bridge_signature_submission (& mut self , request : impl tonic :: IntoRequest < super :: MsgBridgeSignatureSubmission > ,) -> Result < tonic :: Response < super :: MsgBridgeSignatureSubmissionResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Msg/BridgeSignatureSubmission") ; self . inner . unary (request . into_request () , path , codec) . await } } impl < T : Clone > Clone for MsgClient < T > { fn clone (& self) -> Self { Self { inner : self . inner . clone () , } } } impl < T > std :: fmt :: Debug for MsgClient < T > { fn fmt (& self , f : & mut std :: fmt :: Formatter < '_ >) -> std :: fmt :: Result { write ! (f , "MsgClient {{ ... }}") } } }/// OutgoingTx is a withdrawal on the bridged contract
/// TODO: can this type be replaced by outgoing transfer tx
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OutgoingTx {
    #[prost(string, tag="1")]
    pub sender: std::string::String,
    #[prost(string, tag="2")]
    pub dest_addr: std::string::String,
    #[prost(message, optional, tag="3")]
    pub amount: ::std::option::Option<cosmos_sdk_proto::cosmos::base::v1beta1::Coin>,
    #[prost(message, optional, tag="4")]
    pub bridge_fee: ::std::option::Option<cosmos_sdk_proto::cosmos::base::v1beta1::Coin>,
}
/// IDSet represents a set of IDs
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct IdSet {
    #[prost(uint64, repeated, tag="1")]
    pub ids: ::std::vec::Vec<u64>,
}
/// BridgeValidator represents a validator's ETH address and its power
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct BridgeValidator {
    #[prost(uint64, tag="1")]
    pub power: u64,
    #[prost(string, tag="2")]
    pub ethereum_address: std::string::String,
}
/// Valset is the Ethereum Bridge Multsig Set, each peggy validator also
/// maintains an ETH key to sign messages, these are used to check signatures on
/// ETH because of the significant gas savings
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Valset {
    /// Valsets are stored in the store at a nonce
    #[prost(uint64, tag="1")]
    pub nonce: u64,
    /// It is made up of BridgeValidators
    #[prost(message, repeated, tag="2")]
    pub members: ::std::vec::Vec<BridgeValidator>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryParamsRequest {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryParamsResponse {
    #[prost(message, optional, tag="1")]
    pub params: ::std::option::Option<Params>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryCurrentValsetRequest {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryCurrentValsetResponse {
    #[prost(message, optional, tag="1")]
    pub valset: ::std::option::Option<Valset>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryValsetRequestRequest {
    #[prost(uint64, tag="1")]
    pub nonce: u64,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryValsetRequestResponse {
    #[prost(message, optional, tag="1")]
    pub valset: ::std::option::Option<Valset>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryValsetConfirmRequest {
    #[prost(uint64, tag="1")]
    pub nonce: u64,
    #[prost(string, tag="2")]
    pub address: std::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryValsetConfirmResponse {
    #[prost(message, optional, tag="1")]
    pub confirm: ::std::option::Option<MsgValsetConfirm>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryValsetConfirmsByNonceRequest {
    #[prost(uint64, tag="1")]
    pub nonce: u64,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryValsetConfirmsByNonceResponse {
    #[prost(message, repeated, tag="1")]
    pub confirms: ::std::vec::Vec<MsgValsetConfirm>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryLastValsetRequestsRequest {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryLastValsetRequestsResponse {
    #[prost(message, repeated, tag="1")]
    pub valsets: ::std::vec::Vec<Valset>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryLastPendingValsetRequestByAddrRequest {
    #[prost(string, tag="1")]
    pub address: std::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryLastPendingValsetRequestByAddrResponse {
    #[prost(message, optional, tag="1")]
    pub valset: ::std::option::Option<Valset>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryLastPendingBatchRequestByAddrRequest {
    #[prost(string, tag="1")]
    pub address: std::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryLastPendingBatchRequestByAddrResponse {
    #[prost(message, optional, tag="1")]
    pub batch: ::std::option::Option<OutgoingTxBatch>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryOutgoingTxBatchesRequest {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryOutgoingTxBatchesResponse {
    #[prost(message, repeated, tag="1")]
    pub batches: ::std::vec::Vec<OutgoingTxBatch>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryBatchRequestByNonceRequest {
    #[prost(uint64, tag="1")]
    pub nonce: u64,
    #[prost(string, tag="2")]
    pub contract_address: std::string::String,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct QueryBatchRequestByNonceResponse {
    #[prost(message, optional, tag="1")]
    pub batch: ::std::option::Option<OutgoingTxBatch>,
}
# [doc = r" Generated client implementations."] pub mod query_client { # ! [allow (unused_variables , dead_code , missing_docs)] use tonic :: codegen :: * ; # [doc = " Query defines the gRPC querier service"] pub struct QueryClient < T > { inner : tonic :: client :: Grpc < T > , } impl QueryClient < tonic :: transport :: Channel > { # [doc = r" Attempt to create a new client by connecting to a given endpoint."] pub async fn connect < D > (dst : D) -> Result < Self , tonic :: transport :: Error > where D : std :: convert :: TryInto < tonic :: transport :: Endpoint > , D :: Error : Into < StdError > , { let conn = tonic :: transport :: Endpoint :: new (dst) ? . connect () . await ? ; Ok (Self :: new (conn)) } } impl < T > QueryClient < T > where T : tonic :: client :: GrpcService < tonic :: body :: BoxBody > , T :: ResponseBody : Body + HttpBody + Send + 'static , T :: Error : Into < StdError > , < T :: ResponseBody as HttpBody > :: Error : Into < StdError > + Send , { pub fn new (inner : T) -> Self { let inner = tonic :: client :: Grpc :: new (inner) ; Self { inner } } pub fn with_interceptor (inner : T , interceptor : impl Into < tonic :: Interceptor >) -> Self { let inner = tonic :: client :: Grpc :: with_interceptor (inner , interceptor) ; Self { inner } } # [doc = " Deployments queries deployments"] pub async fn params (& mut self , request : impl tonic :: IntoRequest < super :: QueryParamsRequest > ,) -> Result < tonic :: Response < super :: QueryParamsResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Query/Params") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn current_valset (& mut self , request : impl tonic :: IntoRequest < super :: QueryCurrentValsetRequest > ,) -> Result < tonic :: Response < super :: QueryCurrentValsetResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Query/CurrentValset") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn valset_request (& mut self , request : impl tonic :: IntoRequest < super :: QueryValsetRequestRequest > ,) -> Result < tonic :: Response < super :: QueryValsetRequestResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Query/ValsetRequest") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn valset_confirm (& mut self , request : impl tonic :: IntoRequest < super :: QueryValsetConfirmRequest > ,) -> Result < tonic :: Response < super :: QueryValsetConfirmResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Query/ValsetConfirm") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn valset_confirms_by_nonce (& mut self , request : impl tonic :: IntoRequest < super :: QueryValsetConfirmsByNonceRequest > ,) -> Result < tonic :: Response < super :: QueryValsetConfirmsByNonceResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Query/ValsetConfirmsByNonce") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn last_valset_requests (& mut self , request : impl tonic :: IntoRequest < super :: QueryLastValsetRequestsRequest > ,) -> Result < tonic :: Response < super :: QueryLastValsetRequestsResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Query/LastValsetRequests") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn last_pending_valset_request_by_addr (& mut self , request : impl tonic :: IntoRequest < super :: QueryLastPendingValsetRequestByAddrRequest > ,) -> Result < tonic :: Response < super :: QueryLastPendingValsetRequestByAddrResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Query/LastPendingValsetRequestByAddr") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn last_pending_batch_request_by_addr (& mut self , request : impl tonic :: IntoRequest < super :: QueryLastPendingBatchRequestByAddrRequest > ,) -> Result < tonic :: Response < super :: QueryLastPendingBatchRequestByAddrResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Query/LastPendingBatchRequestByAddr") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn outgoing_tx_batches (& mut self , request : impl tonic :: IntoRequest < super :: QueryOutgoingTxBatchesRequest > ,) -> Result < tonic :: Response < super :: QueryOutgoingTxBatchesResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Query/OutgoingTxBatches") ; self . inner . unary (request . into_request () , path , codec) . await } pub async fn batch_request_by_nonce (& mut self , request : impl tonic :: IntoRequest < super :: QueryBatchRequestByNonceRequest > ,) -> Result < tonic :: Response < super :: QueryBatchRequestByNonceResponse > , tonic :: Status > { self . inner . ready () . await . map_err (| e | { tonic :: Status :: new (tonic :: Code :: Unknown , format ! ("Service was not ready: {}" , e . into ())) }) ? ; let codec = tonic :: codec :: ProstCodec :: default () ; let path = http :: uri :: PathAndQuery :: from_static ("/peggy.v1.Query/BatchRequestByNonce") ; self . inner . unary (request . into_request () , path , codec) . await } } impl < T : Clone > Clone for QueryClient < T > { fn clone (& self) -> Self { Self { inner : self . inner . clone () , } } } impl < T > std :: fmt :: Debug for QueryClient < T > { fn fmt (& self , f : & mut std :: fmt :: Formatter < '_ >) -> std :: fmt :: Result { write ! (f , "QueryClient {{ ... }}") } } }