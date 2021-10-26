%#include "common.x"

namespace mazzaroth
{

  // Block contains a block header and transactions committed to this block
  struct Block
  {
    // BlockHeader for this block
    BlockHeader header;
    // List of transactions in the order they were executed
    Transaction transactions<>;
  };

  // BlockHeader contains fields that describe the block
  struct BlockHeader
  {
    // Block height is the number of blocks preceding this block
    unsigned hyper blockHeight;

    // Ending transaction height of this block
    unsigned hyper transactionHeight;

    // Ending consensus sequence number for commits in this block
    unsigned hyper consensusSequenceNumber;

    // The merkle root of the transaction merkle tree in this block
    Hash transactionsMerkleRoot;

    // The merkle root of the receipt merkle tree in this block
    Hash transactionsReceiptRoot;

    // The state root of the Mazzaroth statedb after all transactions have been executed
    Hash stateRoot;

    // The hash of the previous block header
    Hash previousHeader;

    BlockStatus status;
  };

  enum BlockStatus
  {
    UNKNOWN = 0,
    PENDING = 1,
    FINALIZED = 2
  }

  struct BlockHeight
  {
    unsigned hyper height;
  }
}
