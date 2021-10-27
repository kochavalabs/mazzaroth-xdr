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
exports.UpdateType = UpdateType;
exports.AccountUpdateType = AccountUpdateType;
exports.ActionCategoryType = ActionCategoryType;
exports.AuthorityType = AuthorityType;
exports.Update = Update;
exports.AccountUpdate = AccountUpdate;
exports.ActionCategory = ActionCategory;
exports.Authority = Authority;

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
        1: "TRANSACTION",
        2: "TRANSACTIONLIST",
        3: "RECEIPT",
        4: "RECEIPTLIST",
        5: "BLOCK",
        6: "BLOCKLIST",
        7: "BLOCKHEADER",
        8: "BLOCKHEADERLIST",
        9: "CHANNEL",
        10: "CHANNELLIST",
        11: "ACCOUNT",
        12: "HEIGHT",
        13: "ABI"

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
    return new _xdrJsSerialize2.default.Struct(["account", "authorize"], [AuthorizedAccount(), bool()]);
}
function Action() {
    return new _xdrJsSerialize2.default.Struct(["address", "channelID", "nonce", "blockExpirationNumber", "category"], [ID(), ID(), new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper(), ActionCategory()]);
}
function Transaction() {
    return new _xdrJsSerialize2.default.Struct(["signature", "signer", "action"], [Signature(), Authority(), Action()]);
}

// End struct section

// Start enum section

function UpdateType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "CONTRACT",
        2: "CONFIG",
        3: "ACCOUNT"

    });
}

function AccountUpdateType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "ALIAS",
        2: "AUTHORIZATION"

    });
}

function ActionCategoryType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "CALL",
        2: "UPDATE"

    });
}

function AuthorityType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "SELF",
        2: "AUTHORIZED"

    });
}

// End enum section

// Start union section


function Update() {
    return new _xdrJsSerialize2.default.Union(UpdateType(), {

        "UNKNOWN": () => {
            return new _xdrJsSerialize2.default.Void();
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

function ActionCategory() {
    return new _xdrJsSerialize2.default.Union(ActionCategoryType(), {

        "UNKNOWN": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "CALL": () => {
            return Call();
        },

        "UPDATE": () => {
            return Update();
        }

    });
}

function Authority() {
    return new _xdrJsSerialize2.default.Union(AuthorityType(), {

        "UNKNOWN": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "SELF": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "AUTHORIZED": () => {
            return ID();
        }

    });
}

// End union section

// End namespace mazzaroth