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
pub struct Abi {
    pub version: String,

    #[array(var = 2147483647)]
    pub functions: Vec<FunctionSignature>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct FunctionSignature {
    pub functionType: FunctionType,

    #[array(var = 2147483647)]
    pub functionName: String,

    #[array(var = 2147483647)]
    pub parameters: Vec<Parameter>,

    #[array(var = 2147483647)]
    pub returns: Vec<Parameter>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Parameter {
    #[array(var = 2147483647)]
    pub parameterName: String,

    #[array(var = 2147483647)]
    pub parameterType: String,
}

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum FunctionType {
    UNKNOWN = 0,
    READ = 1,
    WRITE = 2,
}

impl Default for FunctionType {
    fn default() -> Self {
        FunctionType::UNKNOWN
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
pub struct Account {
    #[array(var = 32)]
    pub alias: String,

    pub transactionCount: u64,

    #[array(var = 32)]
    pub authorizedAccounts: Vec<AuthorizedAccount>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct AuthorizedAccount {
    pub key: ID,

    #[array(var = 32)]
    pub alias: String,
}

// End struct section

// Start union section

// End union section

// Namspace end mazzaroth
// Namspace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum RequestType {
    UNKNOWN = 0,
    TRANSACTION = 1,
}

impl Default for RequestType {
    fn default() -> Self {
        RequestType::UNKNOWN
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ResponseType {
    UNKNOWN = 0,
    TRANSACTIONID = 1,
    TRANSACTION = 2,
    TRANSACTIONLIST = 3,
    RECEIPT = 4,
    RECEIPTLIST = 5,
    BLOCK = 6,
    BLOCKLIST = 7,
    BLOCKHEADER = 8,
    BLOCKHEADERLIST = 9,
    CONFIG = 10,
    ACCOUNT = 11,
    HEIGHT = 12,
    ABI = 13,
}

impl Default for ResponseType {
    fn default() -> Self {
        ResponseType::UNKNOWN
    }
}
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Request {
    UNKNOWN(()),

    TRANSACTION(Transaction),
}

impl Default for Request {
    fn default() -> Self {
        Request::UNKNOWN(())
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Response {
    UNKNOWN(()),

    TRANSACTIONID(ID),

    TRANSACTION(Transaction),

    #[array(var = 2147483647)]
    TRANSACTIONLIST(Vec<Transaction>),

    RECEIPT(Receipt),

    #[array(var = 2147483647)]
    RECEIPTLIST(Vec<Receipt>),

    BLOCK(Block),

    #[array(var = 2147483647)]
    BLOCKLIST(Vec<Block>),

    BLOCKHEADER(BlockHeader),

    #[array(var = 2147483647)]
    BLOCKHEADERLIST(Vec<BlockHeader>),

    CONFIG(Config),

    ACCOUNT(Account),

    HEIGHT(BlockHeight),

    ABI(Abi),
}

impl Default for Response {
    fn default() -> Self {
        Response::UNKNOWN(())
    }
}
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
    pub blockHeight: u64,

    pub transactionHeight: u64,

    pub consensusSequenceNumber: u64,

    pub transactionsMerkleRoot: Hash,

    pub transactionsReceiptRoot: Hash,

    pub stateRoot: Hash,

    pub previousHeader: Hash,

    pub status: Status,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct BlockHeight {
    pub height: u64,
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
pub struct Argument {
    #[array(var = 2147483647)]
    pub t: String,
}
#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct StatusInfo {
    #[array(var = 256)]
    pub t: String,
}

// End typedef section

// Start struct section

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Status {
    UNKNOWN = 0,
    SUCCESS = 1,
    FAILURE = 2,
    PENDING = 3,
    FINALIZED = 4,
}

impl Default for Status {
    fn default() -> Self {
        Status::UNKNOWN
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
pub struct Receipt {
    pub status: Status,

    pub stateRoot: Hash,

    #[array(var = 2147483647)]
    pub result: String,

    pub statusInfo: StatusInfo,
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
pub struct Call {
    #[array(var = 256)]
    pub function: String,

    #[array(var = 2147483647)]
    pub arguments: Vec<Argument>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Config {
    pub owner: ID,

    #[array(var = 32)]
    pub admins: Vec<ID>,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Contract {
    #[array(var = 2147483647)]
    pub contractBytes: Vec<u8>,

    pub abi: Abi,

    pub contractHash: Hash,

    #[array(var = 100)]
    pub version: String,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Authorization {
    pub account: AuthorizedAccount,

    pub authorize: bool,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Data {
    pub channelID: ID,

    pub nonce: u64,

    pub blockExpirationNumber: u64,

    pub category: Category,
}

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Transaction {
    pub sender: ID,

    pub signer: ID,

    pub signature: Signature,

    pub data: Data,
}

// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum AccountUpdateType {
    UNKNOWN = 0,
    ALIAS = 1,
    AUTHORIZATION = 2,
}

impl Default for AccountUpdateType {
    fn default() -> Self {
        AccountUpdateType::UNKNOWN
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum CategoryType {
    UNKNOWN = 0,
    CALL = 1,
    CONTRACT = 2,
    CONFIG = 3,
    ACCOUNT = 4,
}

impl Default for CategoryType {
    fn default() -> Self {
        CategoryType::UNKNOWN
    }
}
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum AccountUpdate {
    UNKNOWN(()),

    ALIAS(String),

    AUTHORIZATION(Authorization),
}

impl Default for AccountUpdate {
    fn default() -> Self {
        AccountUpdate::UNKNOWN(())
    }
}

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Category {
    UNKNOWN(()),

    CALL(Call),

    CONTRACT(Contract),

    CONFIG(Config),

    ACCOUNT(AccountUpdate),
}

impl Default for Category {
    fn default() -> Self {
        Category::UNKNOWN(())
    }
}
// End union section

// Namspace end mazzaroth
