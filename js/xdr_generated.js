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
  //   enum ActionType
  //     {
  //       ACTION_TYPE_CALL = 0,
  //       ACTION_TYPE_UPDATE = 1
  //     };
  //
  // ===========================================================================
  xdr.enum("ActionType", {
    actionTypeCall: 0,
    actionTypeUpdate: 1
  });

  // === xdr source ============================================================
  //
  //   union Action switch (ActionType type)
  //     {
  //       case ACTION_TYPE_CALL:
  //           Call call; 
  //       case ACTION_TYPE_UPDATE:
  //           Update update; 
  //     };
  //
  // ===========================================================================
  xdr.union("Action", {
    switchOn: xdr.lookup("ActionType"),
    switchName: "type",
    switches: [["actionTypeCall", "call"], ["actionTypeUpdate", "update"]],
    arms: {
      call: xdr.lookup("Call"),
      update: xdr.lookup("Update")
    }
  });

  // === xdr source ============================================================
  //
  //   struct Transaction
  //     {
  //       Signature signature;
  //   
  //       ID address;
  //     };
  //
  // ===========================================================================
  xdr.struct("Transaction", [["signature", xdr.lookup("Signature")], ["address", xdr.lookup("Id")]]);
}); // Automatically generated on 2019-04-17T16:50:58-07:00
// DO NOT EDIT or your changes may be overwritten

/* jshint maxstatements:2147483647  */
/* jshint esnext:true  */

exports.default = types;
