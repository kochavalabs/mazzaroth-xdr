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
  //   typedef unsigned hyper uint64;
  //
  // ===========================================================================
  xdr.typedef("Uint64", xdr.uhyper());

  // === xdr source ============================================================
  //
  //   struct Call
  //     {
  //       string function<256>;
  //   
  //       opaque parameters<>;
  //     };
  //
  // ===========================================================================
  xdr.struct("Call", [["function", xdr.string(256)], ["parameters", xdr.varOpaque()]]);

  // === xdr source ============================================================
  //
  //   struct Update
  //     {
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
  //       ID channelID;    
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
  //       Signature signature;
  //   
  //       ID address;
  //   
  //       Action action;
  //     };
  //
  // ===========================================================================
  xdr.struct("Transaction", [["signature", xdr.lookup("Signature")], ["address", xdr.lookup("Id")], ["action", xdr.lookup("Action")]]);
}); // Automatically generated on 2019-04-22T09:58:16-07:00
// DO NOT EDIT or your changes may be overwritten

/* jshint maxstatements:2147483647  */
/* jshint esnext:true  */

exports.default = types;