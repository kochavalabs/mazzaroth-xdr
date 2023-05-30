
namespace mazzaroth
{
  enum CategoryType
  {
    UNKNOWN = 0,
    CALL = 1,
    DEPLOY = 2,
    PAUSE = 3,
    DELETE = 4,
    CHANNEL = 5
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
    case CHANNEL:
      // Channel
      Channel channel;
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
