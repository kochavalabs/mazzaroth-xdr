
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

  // Config stores contract channel configuration in state and is 
  // accessible through host contract host functions
  struct Config
  {
    // Public Key ID of the channel owner. Only owner can change this to transfer ownership of channel
    ID owner;
    // Public Keys of IDs approved by owner able to modify channel
    ID admins<32>;
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

  struct Authorization
  {
    ID account;

    boolean authorize;
  };

  enum CategoryType
  {
    UNKNOWN = 0,
    CALL = 1,
    CONTRACT = 2,
    CONFIG = 3,
    ACCOUNT = 4,
    AUTHORIZATION = 5
  };

  union Category switch (CategoryType Type)
  {
    case UNKNOWN:
      void;
    case CALL:
      Call call;
    case CONTRACT:
      Contract contract;
    case CONFIG:
      Config config;
    case ACCOUNT:
      Account account;
    case AUTHORIZATION:
      Authorization authorization; 
  };

  // The data of a transaction
  // Set by the client to form a transaction
  // This is marshaled to XDR bytes to sign
  struct Data
  {
    ID channelID;

    unsigned hyper nonce;

    // Highest block number in which to accept this transaction
    unsigned hyper blockExpirationNumber;

    Category category;
  }

  // A transaction that calls a function on a user defined contract.
  struct Transaction
  {
    // ID (public key) of the sender of the transaction
    ID sender;

    // Signer ID (public key) of the transaction signer
    // This should either match the sender or be a key that has been
    // authorized to sign on behalf of the sender through an authorization transaction
    ID signer;

    // Byte array signature of the Transaction bytes signed by the Transaction 
    // sender's private key.
    Signature signature;

    // The transaction data
    Data data;
  };
}
