
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
    UNKNOWN = 0,
    CONTRACT = 1,
    CONFIG = 2,
    ACCOUNT = 3
  };

  union Update switch (UpdateType Type)
  {
    case UNKNOWN:
      void;
    case CONTRACT:
      Contract contract;
    case CONFIG:
      ChannelConfig channelConfig;
    case ACCOUNT:
      AccountUpdate account;
  };

  // An update transaction that provides a contract as a wasm binary.
  struct Contract
  {
    // Contract binary bytes.
    opaque contractBytes<>;

    // Contract ABI
    Abi abi;

    // Sha3 256 Hash of the contract bytes, verified on execution
    Hash contractHash;

    // Version number of the contract, specified by owner
    string version<100>;
  }

  enum AccountUpdateType
  {
    UNKNOWN = 0,
    ALIAS = 1,
    AUTHORIZATION = 2
  };

  union AccountUpdate switch (AccountUpdateType Type)
  {
    case UNKNOWN:
      void;
    case ALIAS:
      string alias;
    case AUTHORIZATION:
      Authorization authorization;
  };

  struct Authorization
  {
    AuthorizedAccount account;

    boolean authorize;
  };

  enum ActionCategoryType
  {
    UNKNOWN = 0,
    CALL = 1,
    UPDATE = 2
  };

  union ActionCategory switch (ActionCategoryType Type)
  {
    case UNKNOWN:
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
    UNKNOWN = 0,

    SELF = 1,

    AUTHORIZED = 2
  };

  union Authority switch (AuthorityType Type)
  {
    case UNKNOWN:
      void;
    case SELF:
      void;
    case AUTHORIZED:
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
