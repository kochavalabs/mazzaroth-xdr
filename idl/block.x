%#include "common.x"

namespace mazzaroth
{

  // BlockHeader contains fields that describe the block
  // TODO: Transaction Bloom and Receipt Bloom
  struct BlockHeader
  {

    string timestamp<256>; 

    uint64 blockHeight;

    Hash txMerkleRoot;

    Hash txReceiptRoot;

    Hash stateRoot;

    Hash previousHeader;

    ID blockProducerAddress;
    
  };
}
