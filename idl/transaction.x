
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

  enum ActionCategoryType
  {
    CALL = 0,
    UPDATE = 1
  };

  union ActionCategory switch (ActionCategoryType type)
  {
    case CALL:
        Call call;
    case UPDATE:
        Update update;
  };

  // The Action data of a transaction
  // Set by the client to form a transaction
  struct Action 
  {
    ID channelId;    

    unsigned hyper nonce;

    ActionCategory category;

  };

  // A transaction that calls a function on a user defined contract.
  struct Transaction
  {
    // Byte array signature of the Transaction bytes signed by the Transaction 
    // sender's private key.
    Signature signature;

    // Byte array representing the id of the sender, this also happens
    // to be the sender's account public key.
    ID address;

    // The action data for this transaction
    Action action;
  };


  // A transaction that has been executed, contains a receipt, and is
  // ready to be stored in the ledger.
  struct CommittedTransaction
  {
    // The transaction itself
    Transaction transaction;

    // The execution ordering of the transaction, provided by consensus
    unsigned hyper sequenceNumber;

    // The receipt hash generated by consensus, to be compared to local execution results
    ID receiptId;

    // The transaction merkle root after adding this transaction to the merkle tree, for validation
    Hash currentTransactionRoot;

     // Consensus signatures
    Signature signatures<>;
  };
}
