%#include "common.x"

namespace mazzaroth
{

  // ContractMetadata stores contract metadata in state for a contract
  struct ContractMetadata
  {
    // Hash is the hash of the contract bytes
    Hash hash;
    // Version number assigned to this contract
    unsigned hyper version;
  };
}
