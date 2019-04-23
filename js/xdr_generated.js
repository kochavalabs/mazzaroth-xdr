"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _jsXdr = require("js-xdr");

var XDR = _interopRequireWildcard(_jsXdr);

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) newObj[key] = obj[key]; } } newObj.default = obj; return newObj; } }

var types = XDR.config(xdr => {

  // === xdr source ============================================================
  //
  //   struct Block
  //     {
  //       // Order is preserved in repeated fields
  //       BlockHeader header;
  //       Transaction transactions<>;
  //     };
  //
  // ===========================================================================
  xdr.struct("Block", [["header", xdr.lookup("BlockHeader")], ["transactions", xdr.varArray(xdr.lookup("Transaction"), 2147483647)]]);

  // === xdr source ============================================================
  //
  //   struct BlockHeader
  //     {
  //   
  //       string timestamp<256>; 
  //   
  //       unsigned hyper blockHeight;
  //   
  //       Hash txMerkleRoot;
  //   
  //       Hash txReceiptRoot;
  //   
  //       Hash stateRoot;
  //   
  //       Hash previousHeader;
  //   
  //       ID blockProducerAddress;
  //       
  //     };
  //
  // ===========================================================================
  xdr.struct("BlockHeader", [["timestamp", xdr.string(256)], ["blockHeight", xdr.uhyper()], ["txMerkleRoot", xdr.lookup("Hash")], ["txReceiptRoot", xdr.lookup("Hash")], ["stateRoot", xdr.lookup("Hash")], ["previousHeader", xdr.lookup("Hash")], ["blockProducerAddress", xdr.lookup("Id")]]);

  // === xdr source ============================================================
  //
  //   typedef opaque Signature[64];
  //
  // ===========================================================================
  xdr.typedef("Signature", xdr.opaque(64));

  // === xdr source ============================================================
  //
  //   typedef opaque ID[32];
  //
  // ===========================================================================
  xdr.typedef("Id", xdr.opaque(32));

  // === xdr source ============================================================
  //
  //   typedef opaque Hash[32];
  //
  // ===========================================================================
  xdr.typedef("Hash", xdr.opaque(32));

  // === xdr source ============================================================
  //
  //   struct Event {
  //       // Name of Event (Function Name)
  //       string key<256>;
  //   
  //       opaque values<>;
  //     };
  //
  // ===========================================================================
  xdr.struct("Event", [["key", xdr.string(256)], ["values", xdr.varOpaque()]]);

  // === xdr source ============================================================
  //
  //   struct Receipt
  //     {
  //       // Status failure or success
  //       ReceiptStatus status;
  //    
  //       // The state root after execution of the transaction
  //       Hash stateRoot;
  //    
  //       // The list of events fired during execution of this transaction
  //       Event events<>;
  //    
  //       // Return results of execution if there is one for function called
  //       opaque result<>;
  //     };
  //
  // ===========================================================================
  xdr.struct("Receipt", [["status", xdr.lookup("ReceiptStatus")], ["stateRoot", xdr.lookup("Hash")], ["events", xdr.varArray(xdr.lookup("Event"), 2147483647)], ["result", xdr.varOpaque()]]);

  // === xdr source ============================================================
  //
  //   enum ReceiptStatus
  //     {
  //       FAILURE = 0,
  //       SUCCESS = 1
  //     };
  //
  // ===========================================================================
  xdr.enum("ReceiptStatus", {
    failure: 0,
    success: 1
  });

  // === xdr source ============================================================
  //
  //   typedef string StatusInfo<256>;
  //
  // ===========================================================================
  xdr.typedef("StatusInfo", xdr.string(256));

  // === xdr source ============================================================
  //
  //   enum IdentifierType
  //     {
  //       NUMBER = 0,
  //       HASH = 1
  //     };
  //
  // ===========================================================================
  xdr.enum("IdentifierType", {
    number: 0,
    hash: 1
  });

  // === xdr source ============================================================
  //
  //   union Identifier switch (IdentifierType type)
  //     {
  //       case NUMBER:
  //         unsigned hyper number;
  //       case HASH:
  //         Hash hash;
  //     };
  //
  // ===========================================================================
  xdr.union("Identifier", {
    switchOn: xdr.lookup("IdentifierType"),
    switchName: "type",
    switches: [["number", "number"], ["hash", "hash"]],
    arms: {
      number: xdr.uhyper(),
      hash: xdr.lookup("Hash")
    }
  });

  // === xdr source ============================================================
  //
  //   struct BlockLookupRequest
  //     {
  //       Identifier id; 
  //     };
  //
  // ===========================================================================
  xdr.struct("BlockLookupRequest", [["id", xdr.lookup("Identifier")]]);

  // === xdr source ============================================================
  //
  //   struct BlockHeaderLookupRequest
  //     {
  //       Identifier id; 
  //     };
  //
  // ===========================================================================
  xdr.struct("BlockHeaderLookupRequest", [["id", xdr.lookup("Identifier")]]);

  // === xdr source ============================================================
  //
  //   struct BlockLookupResponse
  //     {
  //   
  //       // Block that was requested if status is found.
  //       Block block;
  //   
  //       // Status for the requested block.
  //       BlockStatus status;
  //   
  //       // Human readable information to help understand the block status.
  //       StatusInfo statusInfo;
  //     };
  //
  // ===========================================================================
  xdr.struct("BlockLookupResponse", [["block", xdr.lookup("Block")], ["status", xdr.lookup("BlockStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);

  // === xdr source ============================================================
  //
  //   struct BlockHeaderLookupResponse
  //     {
  //   
  //       // Block header that was requested if status is found.
  //       BlockHeader header;
  //   
  //       // Status for the requested block.
  //       BlockStatus status;
  //   
  //       // Human readable information to help understand the block status.
  //       StatusInfo statusInfo;
  //   
  //     };
  //
  // ===========================================================================
  xdr.struct("BlockHeaderLookupResponse", [["header", xdr.lookup("BlockHeader")], ["status", xdr.lookup("BlockStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);

  // === xdr source ============================================================
  //
  //   enum BlockStatus
  //     {
  //   
  //       // Status of the block is unkown.
  //       UNKNOWN = 0,
  //   
  //       // This block has been created.
  //       CREATED = 1,
  //   
  //       // This block has not been created yet.
  //       FUTURE = 2,
  //   
  //       // The block that was looked up was not found.
  //       NOT_FOUND = 3
  //     };
  //
  // ===========================================================================
  xdr.enum("BlockStatus", {
    unknown: 0,
    created: 1,
    future: 2,
    notFound: 3
  });

  // === xdr source ============================================================
  //
  //   struct TransactionLookupRequest 
  //     {
  //       // Unique transaction identifier.
  //       ID transactionId;
  //     };
  //
  // ===========================================================================
  xdr.struct("TransactionLookupRequest", [["transactionId", xdr.lookup("Id")]]);

  // === xdr source ============================================================
  //
  //   struct TransactionLookupResponse
  //     {
  //       // Final transaction written to the blockchain.
  //       Transaction transaction;
  //   
  //       // Current status of the transaction.
  //       TransactionStatus status;
  //   
  //       // Human readable information to help understand the transaction status.
  //       StatusInfo statusInfo;
  //     };
  //
  // ===========================================================================
  xdr.struct("TransactionLookupResponse", [["transaction", xdr.lookup("Transaction")], ["status", xdr.lookup("TransactionStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);

  // === xdr source ============================================================
  //
  //   struct TransactionSubmitRequest
  //     {
  //       Transaction transaction;
  //     };
  //
  // ===========================================================================
  xdr.struct("TransactionSubmitRequest", [["transaction", xdr.lookup("Transaction")]]);

  // === xdr source ============================================================
  //
  //   struct TransactionSubmitResponse
  //     {
  //       // Final transaction written to the blockchain. (if successful)
  //       ID transactionId;
  //   
  //       // Current status of the transaction.
  //       TransactionStatus status;
  //   
  //       // Human readable information to help understand the transaction status.
  //       StatusInfo statusInfo;
  //     };
  //
  // ===========================================================================
  xdr.struct("TransactionSubmitResponse", [["transactionId", xdr.lookup("Id")], ["status", xdr.lookup("TransactionStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);

  // === xdr source ============================================================
  //
  //   enum TransactionStatus
  //     {
  //   
  //       // The transaction status is either not known or not set.
  //       UNKNOWN = 0,
  //   
  //       // The transaction has been accepted by a node and is in the process of being
  //       // submitted to the blockchain.
  //       ACCEPTED = 1,
  //   
  //       // This transaction was not accepted by the blockchain.
  //       REJECTED = 2,
  //   
  //       // The transaction has succesfully been added to the blockchain.
  //       CONFIRMED = 3,
  //   
  //       // This transaction was not found.
  //       NOT_FOUND = 4
  //     };
  //
  // ===========================================================================
  xdr.enum("TransactionStatus", {
    unknown: 0,
    accepted: 1,
    rejected: 2,
    confirmed: 3,
    notFound: 4
  });

  // === xdr source ============================================================
  //
  //   struct ReceiptLookupRequest 
  //     {
  //       // Unique receipt identifier.
  //       ID recepitId;
  //     };
  //
  // ===========================================================================
  xdr.struct("ReceiptLookupRequest", [["recepitId", xdr.lookup("Id")]]);

  // === xdr source ============================================================
  //
  //   struct ReceiptLookupResponse
  //     {
  //       // Final receipt written to the blockchain.
  //       Receipt receipt; 
  //   
  //       // Current status of the receipt
  //       ReceiptStatus status;
  //   
  //       // Human readable information to help understand the receipt status.
  //       StatusInfo statusInfo;
  //     };
  //
  // ===========================================================================
  xdr.struct("ReceiptLookupResponse", [["receipt", xdr.lookup("Receipt")], ["status", xdr.lookup("ReceiptStatus")], ["statusInfo", xdr.lookup("StatusInfo")]]);

  // === xdr source ============================================================
  //
  //   enum ReceiptLookupStatus
  //     {
  //       // The receipt status is either not known or not set.
  //       UNKNOWN = 0,
  //   
  //       // The transaction receipt was found
  //       FOUND = 1,
  //   
  //       // This transaction receipt was not found.
  //       NOT_FOUND = 2
  //     };
  //
  // ===========================================================================
  xdr.enum("ReceiptLookupStatus", {
    unknown: 0,
    found: 1,
    notFound: 2
  });

  // === xdr source ============================================================
  //
  //   struct Call
  //     {
  //       // Contract function to execute.
  //       string function<256>;
  //   
  //       // Parameters to the contract function. The serialization format is defined
  //       // by the contract itself.
  //       opaque parameters<>;
  //     };
  //
  // ===========================================================================
  xdr.struct("Call", [["function", xdr.string(256)], ["parameters", xdr.varOpaque()]]);

  // === xdr source ============================================================
  //
  //   struct Update
  //     {
  //       // Contract binary bytes.
  //       opaque contract<>;
  //     };
  //
  // ===========================================================================
  xdr.struct("Update", [["contract", xdr.varOpaque()]]);

  // === xdr source ============================================================
  //
  //   enum ActionCategoryType
  //     {
  //       CALL = 0,
  //       UPDATE = 1
  //     };
  //
  // ===========================================================================
  xdr.enum("ActionCategoryType", {
    call: 0,
    update: 1
  });

  // === xdr source ============================================================
  //
  //   union ActionCategory switch (ActionCategoryType type)
  //     {
  //       case CALL:
  //           Call call;
  //       case UPDATE:
  //           Update update;
  //     };
  //
  // ===========================================================================
  xdr.union("ActionCategory", {
    switchOn: xdr.lookup("ActionCategoryType"),
    switchName: "type",
    switches: [["call", "call"], ["update", "update"]],
    arms: {
      call: xdr.lookup("Call"),
      update: xdr.lookup("Update")
    }
  });

  // === xdr source ============================================================
  //
  //   struct Action 
  //     {
  //       ID channelId;    
  //   
  //       unsigned hyper nonce;
  //   
  //       ActionCategory category;
  //   
  //     };
  //
  // ===========================================================================
  xdr.struct("Action", [["channelId", xdr.lookup("Id")], ["nonce", xdr.uhyper()], ["category", xdr.lookup("ActionCategory")]]);

  // === xdr source ============================================================
  //
  //   struct Transaction
  //     {
  //       // Byte array signature of the Transaction bytes signed by the Transaction 
  //       // sender's private key.
  //       Signature signature;
  //   
  //       // Byte array representing the id of the sender, this also happens
  //       // to be the sender's account public key.
  //       ID address;
  //   
  //       // The action data for this transaction
  //       Action action;
  //     };
  //
  // ===========================================================================
  xdr.struct("Transaction", [["signature", xdr.lookup("Signature")], ["address", xdr.lookup("Id")], ["action", xdr.lookup("Action")]]);

  // === xdr source ============================================================
  //
  //   struct CommittedTransaction
  //     {
  //       // The transaction itself
  //       Transaction transaction;
  //   
  //       // The execution ordering of the transaction, provided by consensus
  //       unsigned hyper sequenceNumber;
  //   
  //       // The receipt hash generated by consensus, to be compared to local execution results
  //       ID receiptId;
  //   
  //       // The transaction merkle root after adding this transaction to the merkle tree, for validation
  //       Hash currentTransactionRoot;
  //   
  //        // Consensus signatures
  //       Signature signatures<>;
  //     };
  //
  // ===========================================================================
  xdr.struct("CommittedTransaction", [["transaction", xdr.lookup("Transaction")], ["sequenceNumber", xdr.uhyper()], ["receiptId", xdr.lookup("Id")], ["currentTransactionRoot", xdr.lookup("Hash")], ["signatures", xdr.varArray(xdr.lookup("Signature"), 2147483647)]]);
}); // Automatically generated on 2019-04-23T10:33:23-07:00
// DO NOT EDIT or your changes may be overwritten

/* jshint maxstatements:2147483647  */
/* jshint esnext:true  */

exports.default = types;