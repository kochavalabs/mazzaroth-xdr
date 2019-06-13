"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _jsXdr = require("js-xdr");

var XDR = _interopRequireWildcard(_jsXdr);

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) newObj[key] = obj[key]; } } newObj.default = obj; return newObj; } }

var types = XDR.config(xdr => {
  // Namspace start mazzaroth

  // Start typedef section
  // End typedef section

  // Start struct section
  xdr.struct("Account", [["name", xdr.string(0)], ["nonce", xdr.uhyper()], ["value", xdr.uhyper()]]);

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
  xdr.struct("Block", [["header", xdr.lookup("BlockHeader")], ["transactions", xdr.varArray(xdr.lookup("Transaction"), 2147483647)]]);
  xdr.struct("BlockHeader", [["timestamp", xdr.string(256)], ["blockHeight", xdr.uhyper()], ["txMerkleRoot", xdr.lookup("Hash")], ["txReceiptRoot", xdr.lookup("Hash")], ["stateRoot", xdr.lookup("Hash")], ["previousHeader", xdr.lookup("Hash")], ["blockProducerAddress", xdr.lookup("ID")]]);

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
  xdr.struct("ChannelConfig", [["owner", xdr.lookup("ID")], ["validators", xdr.varArray(xdr.lookup("ID"), 2147483647)], ["consensusConfig", xdr.lookup("ConsensusConfig")]]);
  xdr.struct("PBFTConfig", [["checkpointPeriod", xdr.uhyper()]]);

  // End struct section

  // Start enum section

  xdr.enum("ConsensusConfigType", {

    NONE: 0,

    PBFT: 1
  });

  // End enum section

  // Start union section


  xdr.union("ConsensusConfig", {
    switchOn: xdr.lookup("ConsensusConfigType"),
    switchName: "Type",
    switches: [["NONE", xdr.void()], ["PBFT", "PBFT"]],
    arms: {

      PBFT: xdr.lookup("PBFTConfig")

    }
  });

  // End union section

  // End namespace mazzaroth
  // Namspace start mazzaroth

  // Start typedef section
  xdr.typedef("Signature", xdr.opaque(64));

  xdr.typedef("ID", xdr.opaque(32));

  xdr.typedef("Hash", xdr.opaque(32));

  xdr.typedef("Parameter", xdr.varOpaque());

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
  xdr.struct("ContractMetadata", [["hash", xdr.lookup("Hash")], ["version", xdr.uhyper()]]);

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
  xdr.struct("Event", [["key", xdr.string(256)], ["parameters", xdr.varArray(xdr.lookup("Parameter"), 2147483647)]]);

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
  xdr.struct("Receipt", [["status", xdr.lookup("ReceiptStatus")], ["stateRoot", xdr.lookup("Hash")], ["events", xdr.varArray(xdr.lookup("Event"), 2147483647)], ["result", xdr.varOpaque()]]);

  // End struct section

  // Start enum section

  xdr.enum("ReceiptStatus", {

    FAILURE: 0,

    SUCCESS: 1
  });

  // End enum section

  // Start union section


  // End union section

  // End namespace mazzaroth
  // Namspace start mazzaroth

  // Start typedef section
  xdr.typedef("StatusInfo", xdr.string(256));

  // End typedef section

  // Start struct section
  xdr.struct("BlockLookupRequest", [["ID", xdr.lookup("Identifier")]]);
  xdr.struct("BlockHeaderLookupRequest", [["ID", xdr.lookup("Identifier")]]);
  xdr.struct("BlockLookupResponse", [["block", xdr.lookup("Block")], ["status", xdr.lookup("BlockStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);
  xdr.struct("BlockHeaderLookupResponse", [["header", xdr.lookup("BlockHeader")], ["status", xdr.lookup("BlockStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);
  xdr.struct("TransactionLookupRequest", [["transactionID", xdr.lookup("ID")]]);
  xdr.struct("TransactionLookupResponse", [["transaction", xdr.lookup("Transaction")], ["status", xdr.lookup("TransactionStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);
  xdr.struct("TransactionSubmitRequest", [["transaction", xdr.lookup("Transaction")]]);
  xdr.struct("TransactionSubmitResponse", [["transactionID", xdr.lookup("ID")], ["status", xdr.lookup("TransactionStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);
  xdr.struct("TransactionReadonlyRequest", [["sender", xdr.lookup("ID")], ["call", xdr.lookup("Call")]]);
  xdr.struct("TransactionReadonlyResponse", [["result", xdr.varOpaque()], ["stateRoot", xdr.lookup("Hash")], ["status", xdr.lookup("TransactionStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);
  xdr.struct("ReceiptLookupRequest", [["transactionID", xdr.lookup("ID")]]);
  xdr.struct("ReceiptLookupResponse", [["receipt", xdr.lookup("Receipt")], ["status", xdr.lookup("ReceiptLookupStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);
  xdr.struct("AccountNonceLookupRequest", [["account", xdr.lookup("ID")]]);
  xdr.struct("AccountNonceLookupResponse", [["nonce", xdr.uhyper()], ["status", xdr.lookup("NonceLookupStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);

  // End struct section

  // Start enum section

  xdr.enum("IdentifierType", {

    NONE: 0,

    NUMBER: 1,

    HASH: 2
  });

  xdr.enum("BlockStatus", {

    UNKNOWN: 0,

    CREATED: 1,

    FUTURE: 2,

    NOT_FOUND: 3
  });

  xdr.enum("TransactionStatus", {

    UNKNOWN: 0,

    ACCEPTED: 1,

    REJECTED: 2,

    CONFIRMED: 3,

    NOT_FOUND: 4
  });

  xdr.enum("ReceiptLookupStatus", {

    UNKNOWN: 0,

    FOUND: 1,

    NOT_FOUND: 2
  });

  xdr.enum("NonceLookupStatus", {

    UNKNOWN: 0,

    FOUND: 1,

    NOT_FOUND: 2
  });

  // End enum section

  // Start union section


  xdr.union("Identifier", {
    switchOn: xdr.lookup("IdentifierType"),
    switchName: "Type",
    switches: [["NONE", xdr.void()], ["NUMBER", "NUMBER"], ["HASH", "HASH"]],
    arms: {

      NUMBER: xdr.uhyper(),

      HASH: xdr.lookup("Hash")

    }
  });

  // End union section

  // End namespace mazzaroth
  // Namspace start mazzaroth

  // Start typedef section
  // End typedef section

  // Start struct section
  xdr.struct("Call", [["function", xdr.string(256)], ["parameters", xdr.varArray(xdr.lookup("Parameter"), 2147483647)]]);
  xdr.struct("Update", [["contract", xdr.varOpaque()]]);
  xdr.struct("Action", [["channelID", xdr.lookup("ID")], ["nonce", xdr.uhyper()], ["category", xdr.lookup("ActionCategory")]]);
  xdr.struct("Transaction", [["signature", xdr.lookup("Signature")], ["address", xdr.lookup("ID")], ["action", xdr.lookup("Action")]]);
  xdr.struct("CommittedTransaction", [["transaction", xdr.lookup("Transaction")], ["sequenceNumber", xdr.uhyper()], ["receiptID", xdr.varArray(xdr.lookup("ID"), 25)], ["currentTransactionRoot", xdr.lookup("Hash")], ["signatures", xdr.varArray(xdr.lookup("Signature"), 2147483647)]]);

  // End struct section

  // Start enum section

  xdr.enum("ActionCategoryType", {

    NONE: 0,

    CALL: 1,

    UPDATE: 2
  });

  // End enum section

  // Start union section


  xdr.union("ActionCategory", {
    switchOn: xdr.lookup("ActionCategoryType"),
    switchName: "Type",
    switches: [["NONE", xdr.void()], ["CALL", "CALL"], ["UPDATE", "UPDATE"]],
    arms: {

      CALL: xdr.lookup("Call"),

      UPDATE: xdr.lookup("Update")

    }
  });

  // End union section

  // End namespace mazzaroth
});
exports.default = types;