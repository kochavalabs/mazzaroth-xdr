
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

  // An update transaction that provides a contract as a wasm binary.
  struct Contract
  {
    // Version number of the contract, specified by owner
    string version<100>;

     // Public Key ID of the channel owner. Only owner can change this to transfer ownership of channel
    ID owner;
   
    // Contract ABI
    Abi abi;

    // Sha3 256 Hash of the contract bytes, verified on execution
    Hash contractHash;

    // Contract binary bytes.
    opaque contractBytes<>;
  }

  enum CategoryType
  {
    UNKNOWN = 0,
    CALL = 1,
    DEPLOY = 2,
    PAUSE = 3,
    DELETE = 4
  };

  union Category switch (CategoryType Type)
  {
    case UNKNOWN:
      void;
    case CALL:
      Call call;
    case DEPLOY:
      Contract contract;
    case PAUSE:
      boolean pause;
    case DELETE:
      void;
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

    // Byte array signature of the Transaction bytes signed by the Transaction 
    // sender's private key.
    Signature signature;

    // The transaction data
    Data data;
  };
}
