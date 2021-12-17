"use strict";

Object.defineProperty(exports, "__esModule", {
    value: true
});
exports.Abi = Abi;
exports.FunctionSignature = FunctionSignature;
exports.Parameter = Parameter;
exports.FunctionType = FunctionType;
exports.RequestType = RequestType;
exports.ResponseType = ResponseType;
exports.Request = Request;
exports.Response = Response;
exports.Block = Block;
exports.BlockHeader = BlockHeader;
exports.BlockHeight = BlockHeight;
exports.Signature = Signature;
exports.ID = ID;
exports.Hash = Hash;
exports.Argument = Argument;
exports.StatusInfo = StatusInfo;
exports.Status = Status;
exports.Receipt = Receipt;
exports.Call = Call;
exports.Config = Config;
exports.Contract = Contract;
exports.Authorization = Authorization;
exports.Data = Data;
exports.Transaction = Transaction;
exports.CategoryType = CategoryType;
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
        3: "RECEIPT",
        4: "BLOCK",
        5: "BLOCKLIST",
        6: "BLOCKHEADER",
        7: "BLOCKHEADERLIST",
        8: "CONFIG",
        9: "HEIGHT",
        10: "ABI"
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
        "RECEIPT": () => {
            return Receipt();
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
        "CONFIG": () => {
            return Config();
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
    return new _xdrJsSerialize2.default.Struct(["blockHeight", "transactionHeight", "consensusSequenceNumber", "transactionsMerkleRoot", "transactionsReceiptRoot", "stateRoot", "previousHeader", "status"], [new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper(), Hash(), Hash(), Hash(), Hash(), Status()]);
}
function BlockHeight() {
    return new _xdrJsSerialize2.default.Struct(["height"], [new _xdrJsSerialize2.default.UHyper()]);
}
// End struct section

// Start enum section

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
function Status() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "SUCCESS",
        2: "FAILURE",
        3: "PENDING",
        4: "FINALIZED"
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
function Receipt() {
    return new _xdrJsSerialize2.default.Struct(["transactionID", "status", "stateRoot", "result", "statusInfo"], [ID(), Status(), Hash(), new _xdrJsSerialize2.default.Str('', 2147483647), StatusInfo()]);
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
function Call() {
    return new _xdrJsSerialize2.default.Struct(["function", "arguments"], [new _xdrJsSerialize2.default.Str('', 256), new _xdrJsSerialize2.default.VarArray(2147483647, Argument)]);
}
function Config() {
    return new _xdrJsSerialize2.default.Struct(["owner", "admins"], [ID(), new _xdrJsSerialize2.default.VarArray(32, ID)]);
}
function Contract() {
    return new _xdrJsSerialize2.default.Struct(["contractBytes", "abi", "contractHash", "version"], [new _xdrJsSerialize2.default.VarOpaque(2147483647), Abi(), Hash(), new _xdrJsSerialize2.default.Str('', 100)]);
}
function Authorization() {
    return new _xdrJsSerialize2.default.Struct(["account", "authorize"], [ID(), new _xdrJsSerialize2.default.Bool()]);
}
function Data() {
    return new _xdrJsSerialize2.default.Struct(["channelID", "nonce", "blockExpirationNumber", "category"], [ID(), new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper(), Category()]);
}
function Transaction() {
    return new _xdrJsSerialize2.default.Struct(["sender", "signature", "data"], [ID(), Signature(), Data()]);
}
// End struct section

// Start enum section
function CategoryType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "CALL",
        2: "CONTRACT",
        3: "CONFIG"
    });
}

// End enum section

// Start union section

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
            return Config();
        }
    });
}
// End union section

// End namespace mazzaroth