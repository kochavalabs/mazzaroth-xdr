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
  //   struct BlockHeader
  //     {
  //   
  //       string timestamp<256>; 
  //   
  //       uint64 blockHeight;
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
  xdr.struct("BlockHeader", [["timestamp", xdr.string(256)], ["blockHeight", xdr.lookup("Uint64")], ["txMerkleRoot", xdr.lookup("Hash")], ["txReceiptRoot", xdr.lookup("Hash")], ["stateRoot", xdr.lookup("Hash")], ["previousHeader", xdr.lookup("Hash")], ["blockProducerAddress", xdr.lookup("Id")]]);

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
  //   typedef unsigned hyper uint64;
  //
  // ===========================================================================
  xdr.typedef("Uint64", xdr.uhyper());

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
  //       uint64 nonce;
  //   
  //       ActionCategory category;
  //   
  //     };
  //
  // ===========================================================================
  xdr.struct("Action", [["channelId", xdr.lookup("Id")], ["nonce", xdr.lookup("Uint64")], ["category", xdr.lookup("ActionCategory")]]);

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
  //       uint64 sequenceNumber;
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
  xdr.struct("CommittedTransaction", [["transaction", xdr.lookup("Transaction")], ["sequenceNumber", xdr.lookup("Uint64")], ["receiptId", xdr.lookup("Id")], ["currentTransactionRoot", xdr.lookup("Hash")], ["signatures", xdr.varArray(xdr.lookup("Signature"), 2147483647)]]);
}); // Automatically generated on 2019-04-22T12:41:35-07:00
// DO NOT EDIT or your changes may be overwritten

/* jshint maxstatements:2147483647  */
/* jshint esnext:true  */

exports.default = types;