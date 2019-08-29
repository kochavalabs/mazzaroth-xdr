#![allow(non_snake_case)]
#![allow(non_camel_case_types)]
#[macro_use]
extern crate ex_dee_derive;
#[allow(unused_imports)]
use ex_dee::de::{
    read_fixed_array, read_fixed_opaque, read_var_array, read_var_opaque, read_var_string, XDRIn,
};
use ex_dee::error::Error;
#[allow(unused_imports)]
use ex_dee::ser::{
    write_fixed_array, write_fixed_opaque, write_var_array, write_var_opaque, write_var_string,
    XDROut,
};

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

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ChannelConfig {
    pub owner: ID,

    #[array(var = 2147483647)]
    pub validators: Vec<ID>,

    pub consensusConfig: ConsensusConfig,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct PBFTConfig {
    pub checkpointPeriod: u64,
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
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ConsensusConfig {
    NONE(()),

    PBFT(PBFTConfig),
}

impl Default for ConsensusConfig {
    fn default() -> Self {
        ConsensusConfig::NONE(())
    }
}
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
    pub t: Vec<u8>,
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
pub struct ContractMetadata {
    pub hash: Hash,

    pub version: u64,
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
pub struct Event {
    #[array(var = 256)]
    pub key: String,

    #[array(var = 2147483647)]
    pub parameters: Vec<Parameter>,
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
pub struct ExecutionPlan {
    #[array(var = 256)]
    pub host: String,

    pub channelID: ID,

    #[array(var = 100)]
    pub calls: Vec<Call>,
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
    pub events: Vec<Event>,

    #[array(var = 2147483647)]
    pub result: Vec<u8>,
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

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct StatusInfo {
    #[array(var = 256)]
    pub t: String,
}

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
    pub result: Vec<u8>,

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
// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct BasicColumn {
    #[array(var = 40)]
    pub name: String,

    pub typ: BasicType,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct TypedefColumn {
    #[array(fixed = 1)]
    pub parent: Vec<Column>,

    #[array(fixed = 1)]
    pub child: Vec<Column>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct StructColumn {
    #[array(var = 40)]
    pub name: String,

    #[array(var = 40)]
    pub columns: Vec<Column>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct ArrayColumn {
    #[array(var = 40)]
    pub name: String,

    pub fixed: bool,

    pub length: u32,

    #[array(fixed = 1)]
    pub column: Vec<Column>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Table {
    #[array(var = 40)]
    pub name: String,

    #[array(var = 40)]
    pub columns: Vec<Column>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Schema {
    #[array(var = 40)]
    pub tables: Vec<Table>,
}

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum BasicType {
    BOOLEAN = 0,
    STRING = 1,
    OPAQUE = 2,
    INT = 3,
    UNSIGNED_INT = 4,
    HYPER = 5,
    UNSIGNED_HYPER = 6,
    FLOAT = 7,
    DOUBLE = 8,
}

impl Default for BasicType {
    fn default() -> Self {
        BasicType::BOOLEAN
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ColumnType {
    BASIC = 0,
    STRUCT = 1,
    ARRAY = 2,
}

impl Default for ColumnType {
    fn default() -> Self {
        ColumnType::BASIC
    }
}
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Column {
    BASIC(BasicColumn),

    ARRAY(ArrayColumn),

    STRUCT(StructColumn),

    TYPEDEF(TypedefColumn),
}

impl Default for Column {
    fn default() -> Self {
        Column::BASIC(BasicColumn::default())
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
pub struct Update {
    #[array(var = 2147483647)]
    pub contract: Vec<u8>,
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
pub struct CommittedTransaction {
    pub transaction: Transaction,

    pub sequenceNumber: u64,

    #[array(var = 25)]
    pub receiptID: Vec<ID>,

    pub currentTransactionRoot: Hash,

    #[array(var = 2147483647)]
    pub signatures: Vec<Signature>,
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
    PERMISSION = 3,
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
pub enum ActionCategory {
    NONE(()),

    CALL(Call),

    UPDATE(Update),

    PERMISSION(Permission),
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
