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

// Namespace start mazzaroth

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

// Namespace end mazzaroth
// Namespace start mazzaroth

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
    RECEIPT = 3,
    BLOCK = 4,
    BLOCKLIST = 5,
    BLOCKHEADER = 6,
    BLOCKHEADERLIST = 7,
    CONTRACT = 8,
    HEIGHT = 9,
    ABI = 10,
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

    RECEIPT(Receipt),

    BLOCK(Block),

    #[array(var = 2147483647)]
    BLOCKLIST(Vec<Block>),

    BLOCKHEADER(BlockHeader),

    #[array(var = 2147483647)]
    BLOCKHEADERLIST(Vec<BlockHeader>),

    CONTRACT(Contract),

    HEIGHT(BlockHeight),

    ABI(Abi),
}

impl Default for Response {
    fn default() -> Self {
        Response::UNKNOWN(())
    }
}
// End union section

// Namespace end mazzaroth
// Namespace start mazzaroth

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

// Namespace end mazzaroth
// Namespace start mazzaroth

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
// End struct section

// Start union section

// End union section

// Namespace end mazzaroth
// Namespace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Channel {
    pub channelID: ID,
    #[array(var = 32)]
    pub name: String,
    #[array(var = 2147483647)]
    pub configuration: Vec<Config>,
}
// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum ConfigType {
    UNKNOWN = 0,
    ADMIN = 1,
}

impl Default for ConfigType {
    fn default() -> Self {
        ConfigType::UNKNOWN
    }
}
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Config {
    UNKNOWN(()),

    #[array(var = 8)]
    ADMIN(Vec<ID>),
}

impl Default for Config {
    fn default() -> Self {
        Config::UNKNOWN(())
    }
}
// End union section

// Namespace end mazzaroth
// Namespace start mazzaroth

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

// Namespace end mazzaroth
// Namespace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Contract {
    #[array(var = 100)]
    pub version: String,
    pub abi: Abi,
    pub contractHash: Hash,
    #[array(var = 2147483647)]
    pub contractBytes: Vec<u8>,
}
// End struct section

// Start union section

// End union section

// Namespace end mazzaroth
// Namespace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

#[derive(PartialEq, Clone, Default, Debug, XDROut, XDRIn)]
pub struct Receipt {
    pub transactionID: ID,
    pub status: Status,
    pub stateRoot: Hash,
    #[array(var = 2147483647)]
    pub result: String,
    pub statusInfo: StatusInfo,
}
// End struct section

// Start union section

// End union section

// Namespace end mazzaroth
// Namespace start mazzaroth

// Start typedef section

// End typedef section

// Start struct section

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
    pub signature: Signature,
    pub data: Data,
}
// End struct section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum CategoryType {
    UNKNOWN = 0,
    CALL = 1,
    DEPLOY = 2,
    PAUSE = 3,
    DELETE = 4,
    CHANNEL = 5,
}

impl Default for CategoryType {
    fn default() -> Self {
        CategoryType::UNKNOWN
    }
}
// Start union section

#[derive(PartialEq, Clone, Debug, XDROut, XDRIn)]
pub enum Category {
    UNKNOWN(()),
    CALL(Call),

    DEPLOY(Contract),

    PAUSE(bool),

    DELETE(()),
    CHANNEL(Channel),
}

impl Default for Category {
    fn default() -> Self {
        Category::UNKNOWN(())
    }
}
// End union section

// Namespace end mazzaroth
