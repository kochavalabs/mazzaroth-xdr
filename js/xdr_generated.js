"use strict";

Object.defineProperty(exports, "__esModule", {
    value: true
});
exports.Account = Account;
exports.Block = Block;
exports.BlockHeader = BlockHeader;
exports.ChannelConfig = ChannelConfig;
exports.PBFTConfig = PBFTConfig;
exports.ConsensusConfigType = ConsensusConfigType;
exports.ConsensusConfig = ConsensusConfig;
exports.Signature = Signature;
exports.ID = ID;
exports.Hash = Hash;
exports.Parameter = Parameter;
exports.ContractMetadata = ContractMetadata;
exports.Event = Event;
exports.ExecutionPlan = ExecutionPlan;
exports.Receipt = Receipt;
exports.ReceiptStatus = ReceiptStatus;
exports.StatusInfo = StatusInfo;
exports.StateStatus = StateStatus;
exports.BlockLookupRequest = BlockLookupRequest;
exports.BlockHeaderLookupRequest = BlockHeaderLookupRequest;
exports.BlockLookupResponse = BlockLookupResponse;
exports.BlockHeaderLookupResponse = BlockHeaderLookupResponse;
exports.TransactionLookupRequest = TransactionLookupRequest;
exports.TransactionLookupResponse = TransactionLookupResponse;
exports.TransactionSubmitRequest = TransactionSubmitRequest;
exports.TransactionSubmitResponse = TransactionSubmitResponse;
exports.ReadonlyRequest = ReadonlyRequest;
exports.ReadonlyResponse = ReadonlyResponse;
exports.ReceiptLookupRequest = ReceiptLookupRequest;
exports.ReceiptLookupResponse = ReceiptLookupResponse;
exports.AccountNonceLookupRequest = AccountNonceLookupRequest;
exports.AccountNonceLookupResponse = AccountNonceLookupResponse;
exports.AccountInfoLookupRequest = AccountInfoLookupRequest;
exports.AccountInfoLookupResponse = AccountInfoLookupResponse;
exports.IdentifierType = IdentifierType;
exports.BlockStatus = BlockStatus;
exports.TransactionStatus = TransactionStatus;
exports.ReadonlyStatus = ReadonlyStatus;
exports.ReceiptLookupStatus = ReceiptLookupStatus;
exports.NonceLookupStatus = NonceLookupStatus;
exports.InfoLookupStatus = InfoLookupStatus;
exports.Identifier = Identifier;
exports.Call = Call;
exports.Update = Update;
exports.Permission = Permission;
exports.Action = Action;
exports.Transaction = Transaction;
exports.CommittedTransaction = CommittedTransaction;
exports.Input = Input;
exports.PermissionAction = PermissionAction;
exports.ActionCategoryType = ActionCategoryType;
exports.AuthorityType = AuthorityType;
exports.InputType = InputType;
exports.ActionCategory = ActionCategory;
exports.Authority = Authority;

var _xdrJsSerialize = require("xdr-js-serialize");

var _xdrJsSerialize2 = _interopRequireDefault(_xdrJsSerialize);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Account() {
    return new _xdrJsSerialize2.default.Struct(["name", "nonce", "permissionedKeys"], [new _xdrJsSerialize2.default.Str('', 0), new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.VarArray(2147483647, ID)]);
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
    return new _xdrJsSerialize2.default.Struct(["timestamp", "blockHeight", "txMerkleRoot", "txReceiptRoot", "stateRoot", "previousHeader", "blockProducerAddress"], [new _xdrJsSerialize2.default.Str('', 256), new _xdrJsSerialize2.default.UHyper(), Hash(), Hash(), Hash(), Hash(), ID()]);
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
function ChannelConfig() {
    return new _xdrJsSerialize2.default.Struct(["owner", "validators", "consensusConfig"], [ID(), new _xdrJsSerialize2.default.VarArray(2147483647, ID), ConsensusConfig()]);
}
function PBFTConfig() {
    return new _xdrJsSerialize2.default.Struct(["checkpointPeriod"], [new _xdrJsSerialize2.default.UHyper()]);
}

// End struct section

// Start enum section

function ConsensusConfigType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "PBFT"

    });
}

// End enum section

// Start union section


function ConsensusConfig() {
    return new _xdrJsSerialize2.default.Union(ConsensusConfigType(), {

        "NONE": () => {
            return new _xdrJsSerialize2.default.Void();
        },

        "PBFT": () => {
            return PBFTConfig();
        }

    });
}

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
    return new _xdrJsSerialize2.default.VarOpaque(2147483647);
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
function ContractMetadata() {
    return new _xdrJsSerialize2.default.Struct(["hash", "version"], [Hash(), new _xdrJsSerialize2.default.UHyper()]);
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
function Event() {
    return new _xdrJsSerialize2.default.Struct(["key", "parameters"], [new _xdrJsSerialize2.default.Str('', 256), new _xdrJsSerialize2.default.VarArray(2147483647, Parameter)]);
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
function ExecutionPlan() {
    return new _xdrJsSerialize2.default.Struct(["host", "channelID", "calls"], [new _xdrJsSerialize2.default.Str('', 256), ID(), new _xdrJsSerialize2.default.VarArray(100, Call)]);
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
    return new _xdrJsSerialize2.default.Struct(["status", "stateRoot", "events", "result"], [ReceiptStatus(), Hash(), new _xdrJsSerialize2.default.VarArray(2147483647, Event), new _xdrJsSerialize2.default.VarOpaque(2147483647)]);
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

function StatusInfo() {
    return new _xdrJsSerialize2.default.Str('', 256);
}
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
    return new _xdrJsSerialize2.default.Struct(["transactionID", "status", "statusInfo"], [ID(), TransactionStatus(), StatusInfo()]);
}
function ReadonlyRequest() {
    return new _xdrJsSerialize2.default.Struct(["call"], [Call()]);
}
function ReadonlyResponse() {
    return new _xdrJsSerialize2.default.Struct(["result", "stateStatus", "status", "statusInfo"], [new _xdrJsSerialize2.default.VarOpaque(2147483647), StateStatus(), ReadonlyStatus(), StatusInfo()]);
}
function ReceiptLookupRequest() {
    return new _xdrJsSerialize2.default.Struct(["transactionID"], [ID()]);
}
function ReceiptLookupResponse() {
    return new _xdrJsSerialize2.default.Struct(["receipt", "stateStatus", "status", "statusInfo"], [Receipt(), StateStatus(), ReceiptLookupStatus(), StatusInfo()]);
}
function AccountNonceLookupRequest() {
    return new _xdrJsSerialize2.default.Struct(["account"], [ID()]);
}
function AccountNonceLookupResponse() {
    return new _xdrJsSerialize2.default.Struct(["nonce", "stateStatus", "status", "statusInfo"], [new _xdrJsSerialize2.default.UHyper(), StateStatus(), NonceLookupStatus(), StatusInfo()]);
}
function AccountInfoLookupRequest() {
    return new _xdrJsSerialize2.default.Struct(["account"], [ID()]);
}
function AccountInfoLookupResponse() {
    return new _xdrJsSerialize2.default.Struct(["accountInfo", "stateStatus", "status", "statusInfo"], [Account(), StateStatus(), InfoLookupStatus(), StatusInfo()]);
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

function TransactionStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "ACCEPTED",
        2: "REJECTED",
        3: "CONFIRMED",
        4: "NOT_FOUND"

    });
}

function ReadonlyStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "SUCCESS",
        2: "FAILURE"

    });
}

function ReceiptLookupStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "FOUND",
        2: "NOT_FOUND"

    });
}

function NonceLookupStatus() {
    return new _xdrJsSerialize2.default.Enum({
        0: "UNKNOWN",
        1: "FOUND",
        2: "NOT_FOUND"

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

// End union section

// End namespace mazzaroth
// Namespace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Call() {
    return new _xdrJsSerialize2.default.Struct(["function", "parameters"], [new _xdrJsSerialize2.default.Str('', 256), new _xdrJsSerialize2.default.VarArray(2147483647, Parameter)]);
}
function Update() {
    return new _xdrJsSerialize2.default.Struct(["contract"], [new _xdrJsSerialize2.default.VarOpaque(2147483647)]);
}
function Permission() {
    return new _xdrJsSerialize2.default.Struct(["key", "action"], [ID(), PermissionAction()]);
}
function Action() {
    return new _xdrJsSerialize2.default.Struct(["address", "channelID", "nonce", "category"], [ID(), ID(), new _xdrJsSerialize2.default.UHyper(), ActionCategory()]);
}
function Transaction() {
    return new _xdrJsSerialize2.default.Struct(["signature", "signer", "action"], [Signature(), Authority(), Action()]);
}
function CommittedTransaction() {
    return new _xdrJsSerialize2.default.Struct(["transaction", "sequenceNumber", "receiptID", "currentTransactionRoot", "signatures"], [Transaction(), new _xdrJsSerialize2.default.UHyper(), new _xdrJsSerialize2.default.VarArray(25, ID), Hash(), new _xdrJsSerialize2.default.VarArray(2147483647, Signature)]);
}
function Input() {
    return new _xdrJsSerialize2.default.Struct(["inputType", "function", "parameters"], [InputType(), new _xdrJsSerialize2.default.Str('', 256), new _xdrJsSerialize2.default.VarArray(2147483647, Parameter)]);
}

// End struct section

// Start enum section

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
        2: "UPDATE",
        3: "PERMISSION"

    });
}

function AuthorityType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "PERMISSIONED"

    });
}

function InputType() {
    return new _xdrJsSerialize2.default.Enum({
        0: "NONE",
        1: "READONLY",
        2: "EXECUTE",
        3: "CONSTRUCTOR"

    });
}

// End enum section

// Start union section


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
        },

        "PERMISSION": () => {
            return Permission();
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