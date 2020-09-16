#![allow(non_snake_case)]
#![allow(non_camel_case_types)]
#[macro_use]
extern crate xdr_rs_serialize_derive;
#[allow(unused_imports)]
use std::io::Write;
#[allow(unused_imports)]
use xdr_rs_serialize::de::{
    read_fixed_array, read_fixed_array_json, read_fixed_opaque, read_fixed_opaque_json,
    read_var_array, read_var_array_json, read_var_opaque, read_var_opaque_json, read_var_string,
    read_var_string_json, XDRIn,
};
use xdr_rs_serialize::error::Error;
#[allow(unused_imports)]
use xdr_rs_serialize::ser::{
    write_fixed_array, write_fixed_array_json, write_fixed_opaque, write_fixed_opaque_json,
    write_var_array, write_var_array_json, write_var_opaque, write_var_opaque_json,
    write_var_string, write_var_string_json, XDROut,
};

extern crate json;

// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Account {
    pub name: String,

    pub nonce: u64,

    #[array(var = 2147483647)]
    pub permissionedKeys: Vec<ID>,
}

// End struct section

// Start union section

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Block {
    pub header: BlockHeader,

    #[array(var = 2147483647)]
    pub transactions: Vec<Transaction>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct BlockHeader {
    #[array(var = 256)]
    pub timestamp: String,

    pub blockHeight: u64,

    pub transactionHeight: u64,

    pub consensusSequenceNumber: u64,

    pub txMerkleRoot: Hash,

    pub txReceiptRoot: Hash,

    pub stateRoot: Hash,

    pub previousHeader: Hash,

    pub blockProducerAddress: ID,
}

// End struct section

// Start union section

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Signature {
    #[array(fixed = 64)]
    pub t: Vec<u8>,
}
#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ID {
    #[array(fixed = 32)]
    pub t: Vec<u8>,
}
#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Hash {
    #[array(fixed = 32)]
    pub t: Vec<u8>,
}
#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Parameter {
    #[array(var = 2147483647)]
    pub t: String,
}
#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Hash32 {
    #[array(fixed = 32)]
    pub t: Vec<u8>,
}
#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Hash64 {
    #[array(fixed = 64)]
    pub t: Vec<u8>,
}
#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct StatusInfo {
    #[array(var = 256)]
    pub t: String,
}

// End typedef section

// Start struct section

// End struct section

// Start union section

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ChannelConfig {
    pub channelID: ID,

    pub contractHash: Hash,

    #[array(var = 200)]
    pub version: String,

    pub owner: ID,

    #[array(var = 200)]
    pub channelName: String,

    #[array(var = 200)]
    pub admins: Vec<ID>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct GovernanceConfig {
    pub maxBlockSize: u64,

    pub consensus: ConsensusConfigType,

    pub permissioning: Permissioning,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct PermissionedIDs {
    #[array(var = 2147483647)]
    pub allowedIDs: Vec<ID>,

    #[array(var = 2147483647)]
    pub validators: Vec<ID>,
}

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ConsensusConfigType {
    NONE = 0,
    PBFT = 1,
}

impl Default for ConsensusConfigType {
    fn default() -> Self {
        ConsensusConfigType::NONE
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum PermissioningType {
    PUBLIC = 0,
    PRIVATE = 1,
    PERMISSIONED = 2,
}

impl Default for PermissioningType {
    fn default() -> Self {
        PermissioningType::PUBLIC
    }
}
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Permissioning {
    PUBLIC(()),

    PRIVATE(()),

    PERMISSIONED(PermissionedIDs),
}

impl Default for Permissioning {
    fn default() -> Self {
        Permissioning::PUBLIC(())
    }
}
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct DownloadRequest {
    pub downloadRequestPayload: DownloadRequestPayload,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct BatchesRequest {
    pub seqNum: u64,

    pub id: String,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct DownloadResponse {
    pub downloadStatus: DownloadStatus,

    pub downloadResponsePayload: DownloadResponsePayload,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct DownloadHeight {
    pub height: u64,

    pub seqNum: u64,
}

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum DownloadRequestType {
    UNKNOWN = 0,
    HEIGHT = 1,
    BLOCK = 2,
    BATCHES = 3,
}

impl Default for DownloadRequestType {
    fn default() -> Self {
        DownloadRequestType::UNKNOWN
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum DownloadStatus {
    UNKNOWN = 0,
    SUCCESS = 1,
    FAILURE = 2,
}

impl Default for DownloadStatus {
    fn default() -> Self {
        DownloadStatus::UNKNOWN
    }
}
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum DownloadRequestPayload {
    UNKNOWN(()),

    HEIGHT(()),

    BLOCK(u64),

    BATCHES(BatchesRequest),
}

impl Default for DownloadRequestPayload {
    fn default() -> Self {
        DownloadRequestPayload::UNKNOWN(())
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum DownloadResponsePayload {
    UNKNOWN(()),

    HEIGHT(DownloadHeight),

    BLOCK(Block),

    BATCHES(()),
}

impl Default for DownloadResponsePayload {
    fn default() -> Self {
        DownloadResponsePayload::UNKNOWN(())
    }
}
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ExecutionPlan {
    #[array(var = 256)]
    pub host: String,

    #[array(var = 100)]
    pub actions: Vec<Action>,
}

// End struct section

// Start union section

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Receipt {
    pub status: ReceiptStatus,

    pub stateRoot: Hash,

    #[array(var = 2147483647)]
    pub result: String,

    pub statusInfo: StatusInfo,
}

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ReceiptStatus {
    FAILURE = 0,
    SUCCESS = 1,
}

impl Default for ReceiptStatus {
    fn default() -> Self {
        ReceiptStatus::FAILURE
    }
}
// Start union section

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct StateStatus {
    pub previousBlock: u64,

    pub transactionCount: u64,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct BlockLookupRequest {
    pub ID: Identifier,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct BlockHeaderLookupRequest {
    pub ID: Identifier,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct BlockLookupResponse {
    pub block: Block,

    pub stateStatus: StateStatus,

    pub status: BlockStatus,

    pub statusInfo: StatusInfo,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct BlockHeaderLookupResponse {
    pub header: BlockHeader,

    pub stateStatus: StateStatus,

    pub status: BlockStatus,

    pub statusInfo: StatusInfo,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct TransactionLookupRequest {
    pub transactionID: ID,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct TransactionLookupResponse {
    pub transaction: Transaction,

    pub stateStatus: StateStatus,

    pub status: TransactionStatus,

    pub statusInfo: StatusInfo,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct TransactionSubmitRequest {
    pub transaction: Transaction,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct TransactionSubmitResponse {
    pub transactionID: ID,

    pub status: TransactionStatus,

    pub statusInfo: StatusInfo,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ReadonlyRequest {
    pub call: Call,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ReadonlyResponse {
    #[array(var = 2147483647)]
    pub result: String,

    pub stateStatus: StateStatus,

    pub status: ReadonlyStatus,

    pub statusInfo: StatusInfo,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ReceiptLookupRequest {
    pub transactionID: ID,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ReceiptLookupResponse {
    pub receipt: Receipt,

    pub stateStatus: StateStatus,

    pub status: ReceiptLookupStatus,

    pub statusInfo: StatusInfo,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct AccountNonceLookupRequest {
    pub account: ID,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct AccountNonceLookupResponse {
    pub nonce: u64,

    pub stateStatus: StateStatus,

    pub status: NonceLookupStatus,

    pub statusInfo: StatusInfo,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct AccountInfoLookupRequest {
    pub account: ID,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct AccountInfoLookupResponse {
    pub accountInfo: Account,

    pub stateStatus: StateStatus,

    pub status: InfoLookupStatus,

    pub statusInfo: StatusInfo,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ChannelInfoLookupRequest {
    pub infoType: ChannelInfoType,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ChannelInfoLookupResponse {
    pub channelInfo: ChannelInfo,

    pub stateStatus: StateStatus,

    pub status: InfoLookupStatus,

    pub statusInfo: StatusInfo,
}

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum IdentifierType {
    NONE = 0,
    NUMBER = 1,
    HASH = 2,
}

impl Default for IdentifierType {
    fn default() -> Self {
        IdentifierType::NONE
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum BlockStatus {
    UNKNOWN = 0,
    CREATED = 1,
    FUTURE = 2,
    NOT_FOUND = 3,
}

impl Default for BlockStatus {
    fn default() -> Self {
        BlockStatus::UNKNOWN
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum TransactionStatus {
    UNKNOWN = 0,
    ACCEPTED = 1,
    REJECTED = 2,
    CONFIRMED = 3,
    NOT_FOUND = 4,
}

impl Default for TransactionStatus {
    fn default() -> Self {
        TransactionStatus::UNKNOWN
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ReadonlyStatus {
    UNKNOWN = 0,
    SUCCESS = 1,
    FAILURE = 2,
}

impl Default for ReadonlyStatus {
    fn default() -> Self {
        ReadonlyStatus::UNKNOWN
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ReceiptLookupStatus {
    UNKNOWN = 0,
    FOUND = 1,
    NOT_FOUND = 2,
}

impl Default for ReceiptLookupStatus {
    fn default() -> Self {
        ReceiptLookupStatus::UNKNOWN
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum NonceLookupStatus {
    UNKNOWN = 0,
    FOUND = 1,
    NOT_FOUND = 2,
}

impl Default for NonceLookupStatus {
    fn default() -> Self {
        NonceLookupStatus::UNKNOWN
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ChannelInfoType {
    NONE = 0,
    CONTRACT = 1,
    CONFIG = 2,
}

impl Default for ChannelInfoType {
    fn default() -> Self {
        ChannelInfoType::NONE
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum InfoLookupStatus {
    UNKNOWN = 0,
    FOUND = 1,
    NOT_FOUND = 2,
}

impl Default for InfoLookupStatus {
    fn default() -> Self {
        InfoLookupStatus::UNKNOWN
    }
}
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Identifier {
    NONE(()),

    NUMBER(u64),

    HASH(Hash),
}

impl Default for Identifier {
    fn default() -> Self {
        Identifier::NONE(())
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ChannelInfo {
    NONE(()),

    CONTRACT(Contract),

    CONFIG(ChannelConfig),
}

impl Default for ChannelInfo {
    fn default() -> Self {
        ChannelInfo::NONE(())
    }
}
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ReceiptSubscription {
    pub transactionFilter: TransactionFilter,

    pub receiptFilter: ReceiptFilter,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ReceiptSubscriptionResult {
    pub receipt: Receipt,

    pub transactionID: ID,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ReceiptValueFilter {
    pub status: ValueFilter,

    pub stateRoot: ValueFilter,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ActionFilter {
    pub signature: ValueFilter,

    pub signer: ValueFilter,

    pub address: ValueFilter,

    pub channelID: ValueFilter,

    pub nonce: ValueFilter,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ContractFilter {
    pub actionFilter: ActionFilter,

    pub version: ValueFilter,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ConfigFilter {
    pub actionFilter: ActionFilter,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct PermissionFilter {
    pub actionFilter: ActionFilter,

    pub key: ValueFilter,

    pub action: ValueFilter,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct CallFilter {
    pub actionFilter: ActionFilter,

    pub function: ValueFilter,
}

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ValueFilterType {
    NONE = 0,
    STRING = 1,
    HASH32 = 2,
    HASH64 = 3,
    UHYPER = 4,
    INT = 5,
}

impl Default for ValueFilterType {
    fn default() -> Self {
        ValueFilterType::NONE
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum TransactionFilterType {
    NONE = 0,
    GENERIC = 1,
    CONTRACT = 2,
    CONFIG = 3,
    PERMISSION = 4,
    CALL = 5,
}

impl Default for TransactionFilterType {
    fn default() -> Self {
        TransactionFilterType::NONE
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ReceiptFilterType {
    NONE = 0,
    RECEIPT = 1,
}

impl Default for ReceiptFilterType {
    fn default() -> Self {
        ReceiptFilterType::NONE
    }
}
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ValueFilter {
    NONE(()),

    STRING(String),

    HASH32(Hash32),

    HASH64(Hash64),

    UHYPER(u64),

    INT(i32),
}

impl Default for ValueFilter {
    fn default() -> Self {
        ValueFilter::NONE(())
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum TransactionFilter {
    NONE(()),

    GENERIC(ActionFilter),

    CONTRACT(ContractFilter),

    CONFIG(ConfigFilter),

    PERMISSION(PermissionFilter),

    CALL(CallFilter),
}

impl Default for TransactionFilter {
    fn default() -> Self {
        TransactionFilter::NONE(())
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ReceiptFilter {
    NONE(()),

    RECEIPT(ReceiptValueFilter),
}

impl Default for ReceiptFilter {
    fn default() -> Self {
        ReceiptFilter::NONE(())
    }
}
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Call {
    #[array(var = 256)]
    pub function: String,

    #[array(var = 2147483647)]
    pub parameters: Vec<Parameter>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Contract {
    #[array(var = 2147483647)]
    pub contract: Vec<u8>,

    pub version: String,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Permission {
    pub key: ID,

    pub action: PermissionAction,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Action {
    pub address: ID,

    pub channelID: ID,

    pub nonce: u64,

    pub category: ActionCategory,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Transaction {
    pub signature: Signature,

    pub signer: Authority,

    pub action: Action,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Input {
    pub inputType: InputType,

    #[array(var = 256)]
    pub function: String,

    #[array(var = 2147483647)]
    pub parameters: Vec<Parameter>,
}

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum UpdateType {
    NONE = 0,
    CONTRACT = 1,
    CONFIG = 2,
    PERMISSION = 3,
}

impl Default for UpdateType {
    fn default() -> Self {
        UpdateType::NONE
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum PermissionAction {
    REVOKE = 0,
    GRANT = 1,
}

impl Default for PermissionAction {
    fn default() -> Self {
        PermissionAction::REVOKE
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ActionCategoryType {
    NONE = 0,
    CALL = 1,
    UPDATE = 2,
}

impl Default for ActionCategoryType {
    fn default() -> Self {
        ActionCategoryType::NONE
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum AuthorityType {
    NONE = 0,
    PERMISSIONED = 1,
}

impl Default for AuthorityType {
    fn default() -> Self {
        AuthorityType::NONE
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum InputType {
    NONE = 0,
    READONLY = 1,
    EXECUTE = 2,
    CONSTRUCTOR = 3,
}

impl Default for InputType {
    fn default() -> Self {
        InputType::NONE
    }
}
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Update {
    NONE(()),

    CONTRACT(Contract),

    CONFIG(ChannelConfig),

    PERMISSION(Permission),
}

impl Default for Update {
    fn default() -> Self {
        Update::NONE(())
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ActionCategory {
    NONE(()),

    CALL(Call),

    UPDATE(Update),
}

impl Default for ActionCategory {
    fn default() -> Self {
        ActionCategory::NONE(())
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Authority {
    NONE(()),

    PERMISSIONED(ID),
}

impl Default for Authority {
    fn default() -> Self {
        Authority::NONE(())
    }
}
// End union section

// Namspace end mazzaroth
