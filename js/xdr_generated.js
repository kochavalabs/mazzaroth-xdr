"use strict";

Object.defineProperty(exports, "__esModule", {
    value: true
});
exports.Abi = Abi;
exports.FunctionSignature = FunctionSignature;
exports.Parameter = Parameter;
exports.Account = Account;
exports.Block = Block;
exports.BlockHeader = BlockHeader;
exports.Signature = Signature;
exports.ID = ID;
exports.Hash = Hash;
exports.Argument = Argument;
exports.Hash32 = Hash32;
exports.Hash64 = Hash64;
exports.StatusInfo = StatusInfo;
exports.ChannelConfig = ChannelConfig;
exports.DownloadRequest = DownloadRequest;
exports.BatchesRequest = BatchesRequest;
exports.DownloadResponse = DownloadResponse;
exports.DownloadHeight = DownloadHeight;
exports.DownloadRequestType = DownloadRequestType;
exports.DownloadStatus = DownloadStatus;
exports.DownloadRequestPayload = DownloadRequestPayload;
exports.DownloadResponsePayload = DownloadResponsePayload;
exports.Receipt = Receipt;
exports.ReceiptStatus = ReceiptStatus;
exports.StateStatus = StateStatus;
exports.BlockLookupRequest = BlockLookupRequest;
exports.BlockHeaderLookupRequest = BlockHeaderLookupRequest;
exports.BlockLookupResponse = BlockLookupResponse;
exports.BlockHeaderLookupResponse = BlockHeaderLookupResponse;
exports.TransactionLookupRequest = TransactionLookupRequest;
exports.TransactionLookupResponse = TransactionLookupResponse;
exports.TransactionSubmitRequest = TransactionSubmitRequest;
exports.TransactionSubmitResponse = TransactionSubmitResponse;
exports.ReceiptLookupRequest = ReceiptLookupRequest;
exports.ReceiptLookupResponse = ReceiptLookupResponse;
exports.AccountInfoLookupRequest = AccountInfoLookupRequest;
exports.AccountInfoLookupResponse = AccountInfoLookupResponse;
exports.ChannelInfoLookupRequest = ChannelInfoLookupRequest;
exports.ChannelInfoLookupResponse = ChannelInfoLookupResponse;
exports.IdentifierType = IdentifierType;
exports.BlockStatus = BlockStatus;
exports.TransactionType = TransactionType;
exports.TransactionStatus = TransactionStatus;
exports.ReceiptLookupStatus = ReceiptLookupStatus;
exports.ChannelInfoType = ChannelInfoType;
exports.InfoLookupStatus = InfoLookupStatus;
exports.Identifier = Identifier;
exports.TransactionInfo = TransactionInfo;
exports.ChannelInfo = ChannelInfo;
exports.ReceiptSubscription = ReceiptSubscription;
exports.ReceiptSubscriptionResult = ReceiptSubscriptionResult;
exports.ReceiptValueFilter = ReceiptValueFilter;
exports.ActionFilter = ActionFilter;
exports.ContractFilter = ContractFilter;
exports.ConfigFilter = ConfigFilter;
exports.PermissionFilter = PermissionFilter;
exports.CallFilter = CallFilter;
exports.ValueFilterType = ValueFilterType;
exports.TransactionFilterType = TransactionFilterType;
exports.ReceiptFilterType = ReceiptFilterType;
exports.ValueFilter = ValueFilter;
exports.TransactionFilter = TransactionFilter;
exports.ReceiptFilter = ReceiptFilter;
exports.Call = Call;
exports.Contract = Contract;
exports.Permission = Permission;
exports.Action = Action;
exports.Transaction = Transaction;
exports.UpdateType = UpdateType;
exports.PermissionAction = PermissionAction;
exports.ActionCategoryType = ActionCategoryType;
exports.AuthorityType = AuthorityType;
exports.Update = Update;
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
    return new _xdrJsSerialize2.default.Struct(["functions"], [new _xdrJsSerialize2.default.VarArray(2147483647, FunctionSignature)]);
}
function FunctionSignature() {
    return new _xdrJsSerialize2.default.Struct(["functionType", "name", "inputs", "outputs"], [new _xdrJsSerialize2.default.Str('', 2147483647), new _xdrJsSerialize2.default.Str('', 2147483647), new _xdrJsSerialize2.default.VarArray(2147483647, Parameter), new _xdrJsSerialize2.default.VarArray(2147483647, Parameter)]);
}
function Parameter() {
    return new _xdrJsSerialize2.default.Struct(["name", "parameterType", "codec"], [new _xdrJsSerialize2.default.Str('', 2147483647), new _xdrJsSerialize2.default.Str('', 2147483647), new _xdrJsSerialize2.default.Str('', 2147483647)]);
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
function Account() {
    return new _xdrJsSerialize2.default.Struct(["permissionedKeys"], [new _xdrJsSerialize2.default.VarArray(2147483647, ID)]);
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
function Block() {
    return new _xdrJsSerialize2.default.Struct(["header", "transactions"], [BlockHeader(), new _xdrJsSerialize2.default.VarArray(2147483647, Transaction)]);
}
function BlockHeader() {
    return new _xdrJsSerialize2.default.Struct(["timestamp", "blockHeight", "transactionHeight", "consensusSequenceNumber", "txMerkleRoot", "txReceiptRoot", "stateRoot", "previousHeader", "blockProducerAddress"], [new _xdrJsSerialize2.default.Str('', 256), new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper(), Hash(), Hash(), Hash(), Hash(), ID()]);
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

function Parameter() {
    return new _xdrJsSerialize2.default.Str('', 2147483647);
}

function Hash32() {
    return new _xdrJsSerialize2.default.FixedOpaque(32);
}

function Hash64() {
    return new _xdrJsSerialize2.default.FixedOpaque(64);
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
    return new _xdrJsSerialize2.default.Struct(["owner", "channelName", "admins"], [ID(), new _xdrJsSerialize2.default.Str('', 200), new _xdrJsSerialize2.default.VarArray(200, ID)]);
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
function DownloadRequest() {
    return new _xdrJsSerialize2.default.Struct(["downloadRequestPayload"], [DownloadRequestPayload()]);
}
function BatchesRequest() {
    return new _xdrJsSerialize2.default.Struct(["seqNum", "id", "ip", "port"], [new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.Str('', 0), new _xdrJsSerialize2.default.Str('', 0), new _xdrJsSerialize2.default.UHyper()]);
}
function DownloadResponse() {
    return new _xdrJsSerialize2.default.Struct(["downloadStatus", "downloadResponsePayload"], [DownloadStatus(), DownloadResponsePayload()]);
}
function DownloadHeight() {
    return new _xdrJsSerialize2.default.Struct(["height", "seqNum"], [new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper()]);
}

// End struct section

// Start enum section

function DownloadRequestType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "HEIGHT",
        2: "BLOCK",
        3: "BATCHES"

    });
}

function DownloadStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "SUCCESS",
        2: "FAILURE"

    });
}

// End enum section

// Start union section


function DownloadRequestPayload() {
    return new _xdrJsSerialize2.default.Union(DownloadRequestType(), {

        "UNKNOWN": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "HEIGHT": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "BLOCK": () => {
            return new _xdrJsSerialize2.default.UHyper();
        },

        "BATCHES": () => {
            return BatchesRequest();
        }

    });
}

function DownloadResponsePayload() {
    return new _xdrJsSerialize2.default.Union(DownloadRequestType(), {

        "UNKNOWN": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "HEIGHT": () => {
            return DownloadHeight();
        },

        "BLOCK": () => {
            return Block();
        },

        "BATCHES": () => {
            return new _xdrJsSerialize2.default.Void();
        }

    });
}

// End union section

// End namespace mazzaroth
// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Receipt() {
    return new _xdrJsSerialize2.default.Struct(["status", "stateRoot", "result"], [ReceiptStatus(), Hash(), new _xdrJsSerialize2.default.Str('', 2147483647)]);
}

// End struct section

// Start enum section

function ReceiptStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "FAILURE",
        1: "SUCCESS"

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
function StateStatus() {
    return new _xdrJsSerialize2.default.Struct(["previousBlock", "transactionCount"], [new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.UHyper()]);
}
function BlockLookupRequest() {
    return new _xdrJsSerialize2.default.Struct(["ID"], [Identifier()]);
}
function BlockHeaderLookupRequest() {
    return new _xdrJsSerialize2.default.Struct(["ID"], [Identifier()]);
}
function BlockLookupResponse() {
    return new _xdrJsSerialize2.default.Struct(["block", "stateStatus", "status", "statusInfo"], [Block(), StateStatus(), BlockStatus(), StatusInfo()]);
}
function BlockHeaderLookupResponse() {
    return new _xdrJsSerialize2.default.Struct(["header", "stateStatus", "status", "statusInfo"], [BlockHeader(), StateStatus(), BlockStatus(), StatusInfo()]);
}
function TransactionLookupRequest() {
    return new _xdrJsSerialize2.default.Struct(["transactionID"], [ID()]);
}
function TransactionLookupResponse() {
    return new _xdrJsSerialize2.default.Struct(["transaction", "stateStatus", "status", "statusInfo"], [Transaction(), StateStatus(), TransactionStatus(), StatusInfo()]);
}
function TransactionSubmitRequest() {
    return new _xdrJsSerialize2.default.Struct(["transaction"], [Transaction()]);
}
function TransactionSubmitResponse() {
    return new _xdrJsSerialize2.default.Struct(["transactionInfo", "status", "statusInfo"], [TransactionInfo(), TransactionStatus(), StatusInfo()]);
}
function ReceiptLookupRequest() {
    return new _xdrJsSerialize2.default.Struct(["transactionID"], [ID()]);
}
function ReceiptLookupResponse() {
    return new _xdrJsSerialize2.default.Struct(["receipt", "stateStatus", "status", "statusInfo"], [Receipt(), StateStatus(), ReceiptLookupStatus(), StatusInfo()]);
}
function AccountInfoLookupRequest() {
    return new _xdrJsSerialize2.default.Struct(["account"], [ID()]);
}
function AccountInfoLookupResponse() {
    return new _xdrJsSerialize2.default.Struct(["accountInfo", "stateStatus", "status", "statusInfo"], [Account(), StateStatus(), InfoLookupStatus(), StatusInfo()]);
}
function ChannelInfoLookupRequest() {
    return new _xdrJsSerialize2.default.Struct(["infoType"], [ChannelInfoType()]);
}
function ChannelInfoLookupResponse() {
    return new _xdrJsSerialize2.default.Struct(["channelInfo", "stateStatus", "status", "statusInfo"], [ChannelInfo(), StateStatus(), InfoLookupStatus(), StatusInfo()]);
}

// End struct section

// Start enum section

function IdentifierType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "NUMBER",
        2: "HASH"

    });
}

function BlockStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "CREATED",
        2: "FUTURE",
        3: "NOT_FOUND"

    });
}

function TransactionType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "READ",
        2: "WRITE"

    });
}

function TransactionStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "ACCEPTED",
        2: "REJECTED",
        3: "CONFIRMED",
        4: "NOT_FOUND"

    });
}

function ReceiptLookupStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "FOUND",
        2: "NOT_FOUND"

    });
}

function ChannelInfoType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "CONTRACT",
        2: "CONFIG"

    });
}

function InfoLookupStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "FOUND",
        2: "NOT_FOUND"

    });
}

// End enum section

// Start union section


function Identifier() {
    return new _xdrJsSerialize2.default.Union(IdentifierType(), {

        "NONE": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "NUMBER": () => {
            return new _xdrJsSerialize2.default.UHyper();
        },

        "HASH": () => {
            return Hash();
        }

    });
}

function TransactionInfo() {
    return new _xdrJsSerialize2.default.Union(TransactionType(), {

        "NONE": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "READ": () => {
            return Receipt();
        },

        "WRITE": () => {
            return ID();
        }

    });
}

function ChannelInfo() {
    return new _xdrJsSerialize2.default.Union(ChannelInfoType(), {

        "NONE": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "CONTRACT": () => {
            return Contract();
        },

        "CONFIG": () => {
            return ChannelConfig();
        }

    });
}

// End union section

// End namespace mazzaroth
// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function ReceiptSubscription() {
    return new _xdrJsSerialize2.default.Struct(["transactionFilter", "receiptFilter"], [TransactionFilter(), ReceiptFilter()]);
}
function ReceiptSubscriptionResult() {
    return new _xdrJsSerialize2.default.Struct(["receipt", "transactionID"], [Receipt(), ID()]);
}
function ReceiptValueFilter() {
    return new _xdrJsSerialize2.default.Struct(["status", "stateRoot"], [ValueFilter(), ValueFilter()]);
}
function ActionFilter() {
    return new _xdrJsSerialize2.default.Struct(["signature", "signer", "address", "channelID", "nonce"], [ValueFilter(), ValueFilter(), ValueFilter(), ValueFilter(), ValueFilter()]);
}
function ContractFilter() {
    return new _xdrJsSerialize2.default.Struct(["actionFilter", "version"], [ActionFilter(), ValueFilter()]);
}
function ConfigFilter() {
    return new _xdrJsSerialize2.default.Struct(["actionFilter"], [ActionFilter()]);
}
function PermissionFilter() {
    return new _xdrJsSerialize2.default.Struct(["actionFilter", "key", "action"], [ActionFilter(), ValueFilter(), ValueFilter()]);
}
function CallFilter() {
    return new _xdrJsSerialize2.default.Struct(["actionFilter", "function"], [ActionFilter(), ValueFilter()]);
}

// End struct section

// Start enum section

function ValueFilterType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "STRING",
        2: "HASH32",
        3: "HASH64",
        4: "UHYPER",
        5: "INT"

    });
}

function TransactionFilterType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "GENERIC",
        2: "CONTRACT",
        3: "CONFIG",
        4: "PERMISSION",
        5: "CALL"

    });
}

function ReceiptFilterType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "RECEIPT"

    });
}

// End enum section

// Start union section


function ValueFilter() {
    return new _xdrJsSerialize2.default.Union(ValueFilterType(), {

        "NONE": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "STRING": () => {
            return new _xdrJsSerialize2.default.Str('', 256);
        },

        "HASH32": () => {
            return Hash32();
        },

        "HASH64": () => {
            return Hash64();
        },

        "UHYPER": () => {
            return new _xdrJsSerialize2.default.UHyper();
        },

        "INT": () => {
            return new _xdrJsSerialize2.default.Int();
        }

    });
}

function TransactionFilter() {
    return new _xdrJsSerialize2.default.Union(TransactionFilterType(), {

        "NONE": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "GENERIC": () => {
            return ActionFilter();
        },

        "CONTRACT": () => {
            return ContractFilter();
        },

        "CONFIG": () => {
            return ConfigFilter();
        },

        "PERMISSION": () => {
            return PermissionFilter();
        },

        "CALL": () => {
            return CallFilter();
        }

    });
}

function ReceiptFilter() {
    return new _xdrJsSerialize2.default.Union(ReceiptFilterType(), {

        "NONE": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "RECEIPT": () => {
            return ReceiptValueFilter();
        }

    });
}

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
function Permission() {
    return new _xdrJsSerialize2.default.Struct(["key", "action"], [ID(), PermissionAction()]);
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
        0: "NONE",
        1: "CONTRACT",
        2: "CONFIG",
        3: "PERMISSION"

    });
}

function PermissionAction() {
    return new _xdrJsSerialize2.default.Enum({
        0: "REVOKE",
        1: "GRANT"

    });
}

function ActionCategoryType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "CALL",
        2: "UPDATE"

    });
}

function AuthorityType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "PERMISSIONED"

    });
}

// End enum section

// Start union section


function Update() {
    return new _xdrJsSerialize2.default.Union(UpdateType(), {

        "NONE": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "CONTRACT": () => {
            return Contract();
        },

        "CONFIG": () => {
            return ChannelConfig();
        },

        "PERMISSION": () => {
            return Permission();
        }

    });
}

function ActionCategory() {
    return new _xdrJsSerialize2.default.Union(ActionCategoryType(), {

        "NONE": () => {
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

        "NONE": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "PERMISSIONED": () => {
            return ID();
        }

    });
}

// End union section

// End namespace mazzaroth