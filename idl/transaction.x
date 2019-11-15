
%#include "common.x"

namespace mazzaroth
{
  // A transaction that calls a function on a user defined contract.
  struct Call
  {
    // Contract function to execute.
    string function<256>;

    // Parameters to the contract function. The serialization format is defined
    // by the contract itself.
    Parameter parameters<>;
  };

  // A transaction that provides a contract as a wasm binary.
  struct Update
  {
    // Contract binary bytes.
    opaque contract<>;
  };

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

  enum ConfigType
  {
    NONE = 0,
    CHANNELID = 1,
    OWNER = 2,
    CONSENSUS = 3
  };

  union ConfigAction switch (ConfigType Type)
  {
    case NONE:
      void;
    case CHANNELID:
      ID channelID;
    case OWNER:
      ID owner;
    case CONSENSUS:
      ConsensusConfig consensusConfig;
  };

  // A transaction that defines a channel configuration change
  struct Config
  {
    ConfigAction action;
  };

  enum ActionCategoryType
  {
    NONE = 0,
    CALL = 1,
    UPDATE = 2,
    PERMISSION = 3,
    CONFIG = 4
  };

  union ActionCategory switch (ActionCategoryType Type)
  {
    case NONE:
      void;
    case CALL:
      Call call;
    case UPDATE:
      Update update;
    case PERMISSION:
      Permission permission;
    case CONFIG:
      Config config;
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

  // Input for execution in a user defined contract.
  struct Input
  {
    // Type of input: readonly or write transaction
    InputType inputType;

    // Contract function to execute.
    string function<256>;

    // Parameters to the contract function. The serialization format is defined
    // by the contract itself.
    Parameter parameters<>;
  };

  enum InputType
  {
    NONE = 0,
    READONLY = 1,
    EXECUTE = 2,
    CONSTRUCTOR = 3
  };
}
