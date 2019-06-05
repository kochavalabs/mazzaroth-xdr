%#include "common.x"

namespace mazzaroth
{

  // ContractMetadata stores contract metadata in state for a contract
  struct ContractMetadata
  {
    Hash hash;
    unsigned hyper version;
  };
}
