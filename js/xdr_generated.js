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
exports.Receipt = Receipt;
exports.ReceiptStatus = ReceiptStatus;
exports.StatusInfo = StatusInfo;
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
exports.IdentifierType = IdentifierType;
exports.BlockStatus = BlockStatus;
exports.TransactionStatus = TransactionStatus;
exports.ReadonlyStatus = ReadonlyStatus;
exports.ReceiptLookupStatus = ReceiptLookupStatus;
exports.NonceLookupStatus = NonceLookupStatus;
exports.Identifier = Identifier;
exports.Call = Call;
exports.Update = Update;
exports.Action = Action;
exports.Transaction = Transaction;
exports.CommittedTransaction = CommittedTransaction;
exports.ActionCategoryType = ActionCategoryType;
exports.ActionCategory = ActionCategory;

var _jsXdr = require("js-xdr");

var _jsXdr2 = _interopRequireDefault(_jsXdr);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Account() {
    return new _jsXdr2.default.Struct(["name", "nonce", "value"], [new _jsXdr2.default.Str(0), new _jsXdr2.default.UHyper(), new _jsXdr2.default.UHyper()]);
}

// End struct section

// Start enum section


// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Block() {
    return new _jsXdr2.default.Struct(["header", "transactions"], [BlockHeader(), new _jsXdr2.default.VarArray(2147483647, Transaction)]);
}
function BlockHeader() {
    return new _jsXdr2.default.Struct(["timestamp", "blockHeight", "txMerkleRoot", "txReceiptRoot", "stateRoot", "previousHeader", "blockProducerAddress"], [new _jsXdr2.default.Str(256), new _jsXdr2.default.UHyper(), Hash(), Hash(), Hash(), Hash(), ID()]);
}

// End struct section

// Start enum section


// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function ChannelConfig() {
    return new _jsXdr2.default.Struct(["owner", "validators", "consensusConfig"], [ID(), new _jsXdr2.default.VarArray(2147483647, ID), ConsensusConfig()]);
}
function PBFTConfig() {
    return new _jsXdr2.default.Struct(["checkpointPeriod"], [new _jsXdr2.default.UHyper()]);
}

// End struct section

// Start enum section

function ConsensusConfigType() {
    return new _jsXdr2.default.Enum({
        0: "NONE",
        1: "PBFT"

    });
}

// End enum section

// Start union section


function ConsensusConfig() {
    return new _jsXdr2.default.Union(ConsensusConfigType(), {

        "NONE": new _jsXdr2.default.Void(),

        "PBFT": PBFTConfig()

    });
}

// End union section

// End namespace mazzaroth
// Namspace start mazzaroth

// Start typedef section

function Signature() {
    return new _jsXdr2.default.FixedOpaque(64);
}

function ID() {
    return new _jsXdr2.default.FixedOpaque(32);
}

function Hash() {
    return new _jsXdr2.default.FixedOpaque(32);
}

function Parameter() {
    return new _jsXdr2.default.VarOpaque(2147483647);
}
// End typedef section

// Start struct section

// End struct section

// Start enum section


// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function ContractMetadata() {
    return new _jsXdr2.default.Struct(["hash", "version"], [Hash(), new _jsXdr2.default.UHyper()]);
}

// End struct section

// Start enum section


// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Event() {
    return new _jsXdr2.default.Struct(["key", "parameters"], [new _jsXdr2.default.Str(256), new _jsXdr2.default.VarArray(2147483647, Parameter)]);
}

// End struct section

// Start enum section


// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Receipt() {
    return new _jsXdr2.default.Struct(["status", "stateRoot", "events", "result"], [ReceiptStatus(), Hash(), new _jsXdr2.default.VarArray(2147483647, Event), new _jsXdr2.default.VarOpaque(2147483647)]);
}

// End struct section

// Start enum section

function ReceiptStatus() {
    return new _jsXdr2.default.Enum({
        0: "FAILURE",
        1: "SUCCESS"

    });
}

// End enum section

// Start union section


// End union section

// End namespace mazzaroth
// Namspace start mazzaroth

// Start typedef section

function StatusInfo() {
    return new _jsXdr2.default.Str(256);
}
// End typedef section

// Start struct section
function BlockLookupRequest() {
    return new _jsXdr2.default.Struct(["ID"], [Identifier()]);
}
function BlockHeaderLookupRequest() {
    return new _jsXdr2.default.Struct(["ID"], [Identifier()]);
}
function BlockLookupResponse() {
    return new _jsXdr2.default.Struct(["block", "status", "statusInfo"], [Block(), BlockStatus(), StatusInfo()]);
}
function BlockHeaderLookupResponse() {
    return new _jsXdr2.default.Struct(["header", "status", "statusInfo"], [BlockHeader(), BlockStatus(), StatusInfo()]);
}
function TransactionLookupRequest() {
    return new _jsXdr2.default.Struct(["transactionID"], [ID()]);
}
function TransactionLookupResponse() {
    return new _jsXdr2.default.Struct(["transaction", "status", "statusInfo"], [Transaction(), TransactionStatus(), StatusInfo()]);
}
function TransactionSubmitRequest() {
    return new _jsXdr2.default.Struct(["transaction"], [Transaction()]);
}
function TransactionSubmitResponse() {
    return new _jsXdr2.default.Struct(["transactionID", "status", "statusInfo"], [ID(), TransactionStatus(), StatusInfo()]);
}
function ReadonlyRequest() {
    return new _jsXdr2.default.Struct(["call"], [Call()]);
}
function ReadonlyResponse() {
    return new _jsXdr2.default.Struct(["result", "stateRoot", "status", "statusInfo"], [new _jsXdr2.default.VarOpaque(2147483647), Hash(), ReadonlyStatus(), StatusInfo()]);
}
function ReceiptLookupRequest() {
    return new _jsXdr2.default.Struct(["transactionID"], [ID()]);
}
function ReceiptLookupResponse() {
    return new _jsXdr2.default.Struct(["receipt", "status", "statusInfo"], [Receipt(), ReceiptLookupStatus(), StatusInfo()]);
}
function AccountNonceLookupRequest() {
    return new _jsXdr2.default.Struct(["account"], [ID()]);
}
function AccountNonceLookupResponse() {
    return new _jsXdr2.default.Struct(["nonce", "status", "statusInfo"], [new _jsXdr2.default.UHyper(), NonceLookupStatus(), StatusInfo()]);
}

// End struct section

// Start enum section

function IdentifierType() {
    return new _jsXdr2.default.Enum({
        0: "NONE",
        1: "NUMBER",
        2: "HASH"

    });
}

function BlockStatus() {
    return new _jsXdr2.default.Enum({
        0: "UNKNOWN",
        1: "CREATED",
        2: "FUTURE",
        3: "NOT_FOUND"

    });
}

function TransactionStatus() {
    return new _jsXdr2.default.Enum({
        0: "UNKNOWN",
        1: "ACCEPTED",
        2: "REJECTED",
        3: "CONFIRMED",
        4: "NOT_FOUND"

    });
}

function ReadonlyStatus() {
    return new _jsXdr2.default.Enum({
        0: "UNKNOWN",
        1: "SUCCESS",
        2: "FAILURE"

    });
}

function ReceiptLookupStatus() {
    return new _jsXdr2.default.Enum({
        0: "UNKNOWN",
        1: "FOUND",
        2: "NOT_FOUND"

    });
}

function NonceLookupStatus() {
    return new _jsXdr2.default.Enum({
        0: "UNKNOWN",
        1: "FOUND",
        2: "NOT_FOUND"

    });
}

// End enum section

// Start union section


function Identifier() {
    return new _jsXdr2.default.Union(IdentifierType(), {

        "NONE": new _jsXdr2.default.Void(),

        "NUMBER": new _jsXdr2.default.UHyper(),

        "HASH": Hash()

    });
}

// End union section

// End namespace mazzaroth
// Namspace start mazzaroth

// Start typedef section
// End typedef section

// Start struct section
function Call() {
    return new _jsXdr2.default.Struct(["function", "parameters"], [new _jsXdr2.default.Str(256), new _jsXdr2.default.VarArray(2147483647, Parameter)]);
}
function Update() {
    return new _jsXdr2.default.Struct(["contract"], [new _jsXdr2.default.VarOpaque(2147483647)]);
}
function Action() {
    return new _jsXdr2.default.Struct(["channelID", "nonce", "category"], [ID(), new _jsXdr2.default.UHyper(), ActionCategory()]);
}
function Transaction() {
    return new _jsXdr2.default.Struct(["signature", "address", "action"], [Signature(), ID(), Action()]);
}
function CommittedTransaction() {
    return new _jsXdr2.default.Struct(["transaction", "sequenceNumber", "receiptID", "currentTransactionRoot", "signatures"], [Transaction(), new _jsXdr2.default.UHyper(), new _jsXdr2.default.VarArray(25, ID), Hash(), new _jsXdr2.default.VarArray(2147483647, Signature)]);
}

// End struct section

// Start enum section

function ActionCategoryType() {
    return new _jsXdr2.default.Enum({
        0: "NONE",
        1: "CALL",
        2: "UPDATE"

    });
}

// End enum section

// Start union section


function ActionCategory() {
    return new _jsXdr2.default.Union(ActionCategoryType(), {

        "NONE": new _jsXdr2.default.Void(),

        "CALL": Call(),

        "UPDATE": Update()

    });
}

// End union section

// End namespace mazzaroth