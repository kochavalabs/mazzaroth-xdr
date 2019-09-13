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
    // Timestamp stored for this block
    string timestamp<256>; 

    // Block height is the number of blocks preceding this block
    unsigned hyper blockHeight;

    // The merkle root of the transaction merkle tree in this block
    Hash txMerkleRoot;

    // The merkle root of the receipt merkle tree in this block
    Hash txReceiptRoot;

    // The state root of the Mazzaroth statedb after all transactions have been executed
    Hash stateRoot;

    // The hash of the previous block header
    Hash previousHeader;

    // Address associated with the consensus primary at time of block confirmation
    ID blockProducerAddress;
  };
}
