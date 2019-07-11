#![allow(non_snake_case)]
#![allow(non_camel_case_types)]
#[macro_use]
extern crate ex_dee_derive;
use ex_dee::de::{read_fixed_array, read_var_array, read_var_string, XDRIn};
use ex_dee::error::Error;
use ex_dee::ser::{write_fixed_array, write_var_array, write_var_string, XDROut};

// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Account {
    pub name: String,

    pub nonce: u64,

    #[array(var = 2147483647)]
    pub permissioned_keys: Vec<ID>,
}

// End struct section

// Start union section

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Block {
    pub header: BlockHeader,

    #[array(var = 2147483647)]
    pub transactions: Vec<Transaction>,
}

#[derive(Default, Debug, XDROut, XDRIn)]
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

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct ChannelConfig {
    pub owner: ID,

    #[array(var = 2147483647)]
    pub validators: Vec<ID>,

    pub consensusConfig: ConsensusConfig,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct PBFTConfig {
    pub checkpointPeriod: u64,
}

// End struct section

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Signature {
    #[array(fixed = 64)]
    pub t: Vec<u8>,
}
#[derive(Default, Debug, XDROut, XDRIn)]
pub struct ID {
    #[array(fixed = 32)]
    pub t: Vec<u8>,
}
#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Hash {
    #[array(fixed = 32)]
    pub t: Vec<u8>,
}
#[derive(Default, Debug, XDROut, XDRIn)]
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

#[derive(Default, Debug, XDROut, XDRIn)]
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

#[derive(Default, Debug, XDROut, XDRIn)]
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

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Receipt {
    pub status: ReceiptStatus,

    pub stateRoot: Hash,

    #[array(var = 2147483647)]
    pub events: Vec<Event>,

    #[array(var = 2147483647)]
    pub result: Vec<u8>,
}

// End struct section

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct StatusInfo {
    #[array(var = 256)]
    pub t: String,
}

// End typedef section

// Start struct section

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct BlockLookupRequest {
    pub ID: Identifier,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct BlockHeaderLookupRequest {
    pub ID: Identifier,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct BlockLookupResponse {
    pub block: Block,

    pub status: BlockStatus,

    pub statusInfo: StatusInfo,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct BlockHeaderLookupResponse {
    pub header: BlockHeader,

    pub status: BlockStatus,

    pub statusInfo: StatusInfo,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct TransactionLookupRequest {
    pub transactionID: ID,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct TransactionLookupResponse {
    pub transaction: Transaction,

    pub status: TransactionStatus,

    pub statusInfo: StatusInfo,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct TransactionSubmitRequest {
    pub transaction: Transaction,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct TransactionSubmitResponse {
    pub transactionID: ID,

    pub status: TransactionStatus,

    pub statusInfo: StatusInfo,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct ReadonlyRequest {
    pub call: Call,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct ReadonlyResponse {
    #[array(var = 2147483647)]
    pub result: Vec<u8>,

    pub stateRoot: Hash,

    pub status: ReadonlyStatus,

    pub statusInfo: StatusInfo,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct ReceiptLookupRequest {
    pub transactionID: ID,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct ReceiptLookupResponse {
    pub receipt: Receipt,

    pub status: ReceiptLookupStatus,

    pub statusInfo: StatusInfo,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct AccountNonceLookupRequest {
    pub account: ID,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct AccountNonceLookupResponse {
    pub nonce: u64,

    pub status: NonceLookupStatus,

    pub statusInfo: StatusInfo,
}

// End struct section

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Debug, XDROut, XDRIn)]
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
// Start union section

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Call {
    #[array(var = 256)]
    pub function: String,

    #[array(var = 2147483647)]
    pub parameters: Vec<Parameter>,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Update {
    #[array(var = 2147483647)]
    pub contract: Vec<u8>,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Permission {
    pub key: ID,

    pub action: PermissionAction,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Action {
    pub channelID: ID,

    pub nonce: u64,

    pub category: ActionCategory,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Transaction {
    pub signature: Signature,

    pub onBehalfOf: Authority,

    pub address: ID,

    pub action: Action,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct CommittedTransaction {
    pub transaction: Transaction,

    pub sequenceNumber: u64,

    #[array(var = 25)]
    pub receiptID: Vec<ID>,

    pub currentTransactionRoot: Hash,

    #[array(var = 2147483647)]
    pub signatures: Vec<Signature>,
}

#[derive(Default, Debug, XDROut, XDRIn)]
pub struct Input {
    pub inputType: InputType,

    #[array(var = 256)]
    pub function: String,

    #[array(var = 2147483647)]
    pub parameters: Vec<Parameter>,
}

// End struct section

#[derive(Debug, XDROut, XDRIn)]
pub enum PermissionAction {
    REVOKE = 0,
    GRANT = 1,
}

impl Default for PermissionAction {
    fn default() -> Self {
        PermissionAction::REVOKE
    }
}

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Debug, XDROut, XDRIn)]
pub enum AuthortiyType {
    NONE = 0,
    PERMISSIONED = 1,
}

impl Default for AuthortiyType {
    fn default() -> Self {
        AuthortiyType::NONE
    }
}

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Debug, XDROut, XDRIn)]
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

#[derive(Debug, XDROut, XDRIn)]
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
