"use strict";

Object.defineProperty(exports, "__esModule", {
    value: true
});
exports.Abi = Abi;
exports.FunctionSignature = FunctionSignature;
exports.Parameter = Parameter;
exports.FunctionType = FunctionType;
exports.Account = Account;
exports.AuthorizedAccount = AuthorizedAccount;
exports.RequestType = RequestType;
exports.ResponseType = ResponseType;
exports.Request = Request;
exports.Response = Response;
exports.Block = Block;
exports.BlockHeader = BlockHeader;
exports.BlockHeight = BlockHeight;
exports.BlockStatus = BlockStatus;
exports.Signature = Signature;
exports.ID = ID;
exports.Hash = Hash;
exports.Argument = Argument;
exports.StatusInfo = StatusInfo;
exports.ChannelConfig = ChannelConfig;
exports.Receipt = Receipt;
exports.ReceiptStatus = ReceiptStatus;
exports.Call = Call;
exports.Contract = Contract;
exports.Authorization = Authorization;
exports.Action = Action;
exports.Transaction = Transaction;
exports.AccountUpdateType = AccountUpdateType;
exports.CategoryType = CategoryType;
exports.AccountUpdate = AccountUpdate;
exports.Category = Category;

var _xdrJsSerialize = require("xdr-js-serialize");

var _xdrJsSerialize2 = _interopRequireDefault(_xdrJsSerialize);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Abi() {
    return new _xdrJsSerialize2.default.Struct(["version", "functions"], [new _xdrJsSerialize2.default.Str('', 0), new _xdrJsSerialize2.default.VarArray(2147483647, FunctionSignature)]);
}
function FunctionSignature() {
    return new _xdrJsSerialize2.default.Struct(["functionType", "functionName", "parameters", "returns"], [FunctionType(), new _xdrJsSerialize2.default.Str('', 2147483647), new _xdrJsSerialize2.default.VarArray(2147483647, Parameter), new _xdrJsSerialize2.default.VarArray(2147483647, Parameter)]);
}
function Parameter() {
    return new _xdrJsSerialize2.default.Struct(["parameterName", "parameterType"], [new _xdrJsSerialize2.default.Str('', 2147483647), new _xdrJsSerialize2.default.Str('', 2147483647)]);
}

// End struct section

// Start enum section

function FunctionType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "READ",
        2: "WRITE"

    });
}

// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Account() {
    return new _xdrJsSerialize2.default.Struct(["alias", "transactionCount", "authorizedAccounts"], [new _xdrJsSerialize2.default.Str('', 32), new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.VarArray(32, AuthorizedAccount)]);
}
function AuthorizedAccount() {
    return new _xdrJsSerialize2.default.Struct(["key", "alias"], [ID(), new _xdrJsSerialize2.default.Str('', 32)]);
}

// End struct section

// Start enum section


// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section

// End struct section

// Start enum section

function RequestType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "TRANSACTION"

    });
}

function ResponseType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "TRANSACTIONID",
        2: "TRANSACTION",
        3: "TRANSACTIONLIST",
        4: "RECEIPT",
        5: "RECEIPTLIST",
        6: "BLOCK",
        7: "BLOCKLIST",
        8: "BLOCKHEADER",
        9: "BLOCKHEADERLIST",
        10: "CHANNEL",
        11: "CHANNELLIST",
        12: "ACCOUNT",
        13: "HEIGHT",
        14: "ABI"

    });
}

// End enum section

// Start union section


function Request() {
    return new _xdrJsSerialize2.default.Union(RequestType(), {

        "UNKNOWN": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "TRANSACTION": () => {
            return Transaction();
        }

    });
}

function Response() {
    return new _xdrJsSerialize2.default.Union(ResponseType(), {

        "UNKNOWN": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "TRANSACTIONID": () => {
            return ID();
        },

        "TRANSACTION": () => {
            return Transaction();
        },

        "TRANSACTIONLIST": () => {
            return new _xdrJsSerialize2.default.VarArray(2147483647, Transaction);
        },

        "RECEIPT": () => {
            return Receipt();
        },

        "RECEIPTLIST": () => {
            return new _xdrJsSerialize2.default.VarArray(2147483647, Receipt);
        },

        "BLOCK": () => {
            return Block();
        },

        "BLOCKLIST": () => {
            return new _xdrJsSerialize2.default.VarArray(2147483647, Block);
        },

        "BLOCKHEADER": () => {
            return BlockHeader();
        },

        "BLOCKHEADERLIST": () => {
            return new _xdrJsSerialize2.default.VarArray(2147483647, BlockHeader);
        },

        "CHANNEL": () => {
            return ChannelConfig();
        },

        "CHANNELLIST": () => {
            return new _xdrJsSerialize2.default.VarArray(2147483647, ChannelConfig);
        },

        "ACCOUNT": () => {
            return Account();
        },

        "HEIGHT": () => {
            return BlockHeight();
        },

        "ABI": () => {
            return Abi();
        }

    });
}

// End union section

// End namespace mazzaroth
// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Block() {
    return new _xdrJsSerialize2.default.Struct(["header", "transactions"], [BlockHeader(), new _xdrJsSerialize2.default.VarArray(2147483647, Transaction)]);
}
function BlockHeader() {
    return new _xdrJsSerialize2.default.Struct(["blockHeight", "transactionHeight", "consensusSequenceNumber", "transactionsMerkleRoot", "transactionsReceiptRoot", "stateRoot", "previousHeader", "status"], [new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper(), Hash(), Hash(), Hash(), Hash(), BlockStatus()]);
}
function BlockHeight() {
    return new _xdrJsSerialize2.default.Struct(["height"], [new _xdrJsSerialize2.default.UHyper()]);
}

// End struct section

// Start enum section

function BlockStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "PENDING",
        2: "FINALIZED"

    });
}

// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namespace start mazzaroth

// Start typedef section

function Signature() {
    return new _xdrJsSerialize2.default.FixedOpaque(64);
}

function ID() {
    return new _xdrJsSerialize2.default.FixedOpaque(32);
}

function Hash() {
    return new _xdrJsSerialize2.default.FixedOpaque(32);
}

function Argument() {
    return new _xdrJsSerialize2.default.Str('', 2147483647);
}

function StatusInfo() {
    return new _xdrJsSerialize2.default.Str('', 256);
}
// End typedef section

// Start struct section

// End struct section

// Start enum section


// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function ChannelConfig() {
    return new _xdrJsSerialize2.default.Struct(["owner", "admins"], [ID(), new _xdrJsSerialize2.default.VarArray(32, ID)]);
}

// End struct section

// Start enum section


// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Receipt() {
    return new _xdrJsSerialize2.default.Struct(["status", "stateRoot", "result", "statusInfo"], [ReceiptStatus(), Hash(), new _xdrJsSerialize2.default.Str('', 2147483647), StatusInfo()]);
}

// End struct section

// Start enum section

function ReceiptStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "FAILURE",
        2: "SUCCESS"

    });
}

// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Call() {
    return new _xdrJsSerialize2.default.Struct(["function", "arguments"], [new _xdrJsSerialize2.default.Str('', 256), new _xdrJsSerialize2.default.VarArray(2147483647, Argument)]);
}
function Contract() {
    return new _xdrJsSerialize2.default.Struct(["contractBytes", "abi", "contractHash", "version"], [new _xdrJsSerialize2.default.VarOpaque(2147483647), Abi(), Hash(), new _xdrJsSerialize2.default.Str('', 100)]);
}
function Authorization() {
    return new _xdrJsSerialize2.default.Struct(["account", "authorize"], [AuthorizedAccount(), new _xdrJsSerialize2.default.Bool()]);
}
function Action() {
    return new _xdrJsSerialize2.default.Struct(["channelID", "nonce", "blockExpirationNumber", "category"], [ID(), new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper(), Category()]);
}
function Transaction() {
    return new _xdrJsSerialize2.default.Struct(["signature", "sender", "signer", "action"], [Signature(), ID(), ID(), Action()]);
}

// End struct section

// Start enum section

function AccountUpdateType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "ALIAS",
        2: "AUTHORIZATION"

    });
}

function CategoryType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "CALL",
        2: "CONTRACT",
        3: "CONFIG",
        4: "ACCOUNT"

    });
}

// End enum section

// Start union section


function AccountUpdate() {
    return new _xdrJsSerialize2.default.Union(AccountUpdateType(), {

        "UNKNOWN": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "ALIAS": () => {
            return new _xdrJsSerialize2.default.Str('', 0);
        },

        "AUTHORIZATION": () => {
            return Authorization();
        }

    });
}

function Category() {
    return new _xdrJsSerialize2.default.Union(CategoryType(), {

        "UNKNOWN": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "CALL": () => {
            return Call();
        },

        "CONTRACT": () => {
            return Contract();
        },

        "CONFIG": () => {
            return ChannelConfig();
        },

        "ACCOUNT": () => {
            return AccountUpdate();
        }

    });
}

// End union section

// End namespace mazzaroth