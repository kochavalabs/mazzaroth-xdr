%#include "common.x"

namespace mazzaroth
{

  // ContractMetadata stores contract metadata in state for a contract
  struct ContractMetadata
  {
    string name;
    ID owner;
    unsigned hyper version;
  };
}
