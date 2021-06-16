
%#include "common.x"

namespace mazzaroth
{
  // A transaction that calls a function on a user defined contract.
  struct Call
  {
    // Contract function to execute.
    string function<256>;

    // Arguments to the contract function. The serialization format is defined
    // by the contract itself.
    Argument arguments<>;
  };

  enum UpdateType
  {
    NONE = 0,
    CONTRACT = 1,
    CONFIG = 2,
    PERMISSION = 3
  };

  union Update switch (UpdateType Type)
  {
    case NONE:
      void;
    case CONTRACT:
      Contract contract;
    case CONFIG:
      ChannelConfig channelConfig;
    case PERMISSION:
      Permission permission;
  };

  // An update transaction that provides a contract as a wasm binary.
  struct Contract
  {
    // Contract binary bytes.
    opaque contractBytes<>;

    // Contract ABI as JSON string
    string abi<>;

    // Sha3 256 Hash of the contract bytes, verified on execution
    Hash contractHash;

    // Version number of the contract, specified by owner
    string version<100>;
  }

  enum PermissionAction
  {
    REVOKE = 0,
    GRANT = 1,
  };

  struct Permission
  {
    ID key;

    PermissionAction action;
  };

  enum ActionCategoryType
  {
    NONE = 0,
    CALL = 1,
    UPDATE = 2
  };

  union ActionCategory switch (ActionCategoryType Type)
  {
    case NONE:
      void;
    case CALL:
      Call call;
    case UPDATE:
      Update update;
  };

  // The Action data of a transaction
  // Set by the client to form a transaction
  struct Action 
  {
    // Byte array representing the id of the sender, this also happens
    // to be the sender's account public key.
    ID address;

    ID channelID;

    unsigned hyper nonce;

    // Highest block number in which to accept this transaction
    unsigned hyper blockExpirationNumber;

    ActionCategory category;

  };

  enum AuthorityType
  {
    NONE = 0,

    PERMISSIONED = 1,
  };

  union Authority switch (AuthorityType Type)
  {
    case NONE:
      void;
    case PERMISSIONED:
      ID origin;
  };

  // A transaction that calls a function on a user defined contract.
  struct Transaction
  {
    // Byte array signature of the Transaction bytes signed by the Transaction 
    // sender's private key.
    Signature signature;

    // Authority of the signer of the transaction. This will indicate if this
    // transaction is being signed by a key that the original account owner gave
    // permission to.
    Authority signer;

    // The action data for this transaction
    Action action;
  };
}
