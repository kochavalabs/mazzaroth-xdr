%#include "common.x"

namespace mazzaroth
{
  // ContractChannelConfig stores contract channel configuration in state and is 
  // accessible through host contract host functions
  struct ContractChannelConfig
  {
    // Unique channel id which cannot be changed
    ID channelID;
    // Contract binary bytes
    opaque contract<>;
    // Version number of the contract, specified by owner
    string version;
    // Public Key ID of the channel owner. Only owner can change this to transfer ownership of channel
    ID owner;
    // Human readable channel name
    string channelName;
    // Public Keys of IDs approved by owner able to modify channel
    ID admins<>;
  };

  // GovernanceConfig stores governance information about a channel 
  struct GovernanceConfig
  {
    // Max number of transactions to store in a block
    unsigned hyper maxBlockSize;
    // Consensus type being used
    ConsensusConfigType consensus;
    // Permissioning configurations
    Permissioning permissioning;
  }

  // Consensus type
  enum ConsensusConfigType
  {
    NONE = 0,
    PBFT = 1
  };

  // Permissioning types
  enum PermissioningType
  {
    PUBLIC = 0,
    PRIVATE = 1
  };

  struct PrivatePermissioning 
  {
    // List of public key ids authorized to participate in channel
    ID allowedIDs<>;
    // List of public key ids authorized to validate in consensus
    ID validators<>;
  };

  union Permissioning switch (PermissioningType Type)
  {
  case PUBLIC:
      void;
  case PRIVATE:
      PrivatePermissioning privatePermissioning;
  };
}