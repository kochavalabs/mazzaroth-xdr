%#include "common.x"
%#include "block.x"
%#include "receipt.x"
%#include "transaction.x"

namespace mazzaroth
{
  typedef string StatusInfo<256>;

  struct StateStatus
  {
    Hash previousBlock;

    unsigned hyper transactionCount;

  };

  enum IdentifierType
  {
    NONE = 0, 
    NUMBER = 1,
    HASH = 2
  };

  union Identifier switch (IdentifierType Type)
  {
    case NONE:
      void;
    case NUMBER:
      unsigned hyper number;
    case HASH:
      Hash hash;
  };

  struct BlockLookupRequest
  {
    Identifier ID;
  };

  struct BlockHeaderLookupRequest
  {
    Identifier ID;
  };

  struct BlockLookupResponse
  {

    // Block that was requested if status is found.
    Block block;

    // Current state status
    StateStatus stateStatus;

    // Status for the requested block.
    BlockStatus status;

    // Human readable information to help understand the block status.
    StatusInfo statusInfo;
  };

  struct BlockHeaderLookupResponse
  {

    // Block header that was requested if status is found.
    BlockHeader header;

    // Current state status
    StateStatus stateStatus;

    // Status for the requested block.
    BlockStatus status;

    // Human readable information to help understand the block status.
    StatusInfo statusInfo;

  };

  // Status of a block.
  enum BlockStatus
  {

    // Status of the block is unkown.
    UNKNOWN = 0,

    // This block has been created.
    CREATED = 1,

    // This block has not been created yet.
    FUTURE = 2,

    // The block that was looked up was not found.
    NOT_FOUND = 3
  };

  // Request for a node to look up the status and value of a transaction.
  struct TransactionLookupRequest 
  {
    // Unique transaction identifier.
    ID transactionID;
  };

  // Response to lookup request.
  struct TransactionLookupResponse
  {
    // Final transaction written to the blockchain.
    Transaction transaction;

    // Current state status
    StateStatus stateStatus;

    // Current status of the transaction.
    TransactionStatus status;

    // Human readable information to help understand the transaction status.
    StatusInfo statusInfo;
  };

  // Message sent to a node to submit a transaction.
  struct TransactionSubmitRequest
  {
    Transaction transaction;
  };

  // Response from a node from a transaction Request.
  struct TransactionSubmitResponse
  {
    // Final transaction written to the blockchain. (if successful)
    ID transactionID;

    // Current status of the transaction.
    TransactionStatus status;

    // Human readable information to help understand the transaction status.
    StatusInfo statusInfo;
  };

  // Message sent to a node to submit a readonly request.
  struct ReadonlyRequest
  {
    // Reaonly Request just takes call
    Call call;
  };

  // Response from a node for a readonly request.
  struct ReadonlyResponse
  {
    // Return results of execution
    opaque result<>;

    // Current state status
    StateStatus stateStatus;

    // Status of the request.
    ReadonlyStatus status;

    // Human readable information to help understand the transaction status.
    StatusInfo statusInfo;
  };

  // Status of a transaction.
  enum TransactionStatus
  {
    // The transaction status is either not known or not set.
    UNKNOWN = 0,

    // The transaction has been accepted by a node and is in the process of being
    // submitted to the blockchain.
    ACCEPTED = 1,

    // This transaction was not accepted by the blockchain.
    REJECTED = 2,

    // The transaction has succesfully been added to the blockchain.
    CONFIRMED = 3,

    // This transaction was not found.
    NOT_FOUND = 4
  };

  // Status of a readonly request.
  enum ReadonlyStatus
  {
    // The status is either not known or not set.
    UNKNOWN = 0,

    // Readonly request was successfully executed.
    SUCCESS = 1,

    // Readonly request did not execute successfully.
    FAILURE = 2,
  };

  // Request for a node to look up a transaction receipt.
  struct ReceiptLookupRequest 
  {
    // Unique transaction identifier.
    ID transactionID;
  };

  // Response to receipt lookup request.
  struct ReceiptLookupResponse
  {
    // Final receipt written to the blockchain.
    Receipt receipt; 

    // Current state status
    StateStatus stateStatus;

    // Current status of the receipt
    ReceiptLookupStatus status;

    // Human readable information to help understand the receipt status.
    StatusInfo statusInfo;
  };

  // Status of a receipt.
  enum ReceiptLookupStatus
  {
    // The receipt status is either not known or not set.
    UNKNOWN = 0,

    // The transaction receipt was found
    FOUND = 1,

    // This transaction receipt was not found.
    NOT_FOUND = 2
  };

  // Request for a node to look up an account nonce.
  struct AccountNonceLookupRequest 
  {
    // ID of the account
    ID account;
  };

  // Response to account nonce lookup request.
  struct AccountNonceLookupResponse
  {
    // Final receipt written to the blockchain.
    unsigned hyper nonce; 

    // Current state status
    StateStatus stateStatus;

    // Status of the lookup
    NonceLookupStatus status;

    // Human readable information to help understand the status.
    StatusInfo statusInfo;
  };

  // Status of a nonce lookup.
  enum NonceLookupStatus
  {
    // The status is either not known or not set.
    UNKNOWN = 0,

    // The account nonce was found.
    FOUND = 1,

    // The account nonce was not found.
    NOT_FOUND = 2
  };

  // Request for a node to look up account info.
  struct AccountInfoLookupRequest
  {
    // ID of the account
    ID account;
  };

  // Response to account info lookup request.
  struct AccountInfoLookupResponse
  {
    // Final receipt written to the blockchain.
    Account accountInfo;

    // Current state status
    StateStatus stateStatus;

    // Status of the lookup
    InfoLookupStatus status;

    // Human readable information to help understand the status.
    StatusInfo statusInfo;
  };

  // Status of a info lookup.
  enum InfoLookupStatus
  {
    // The status is either not known or not set.
    UNKNOWN = 0,

    // The account info was found.
    FOUND = 1,

    // The account info was not found.
    NOT_FOUND = 2
  };

}
