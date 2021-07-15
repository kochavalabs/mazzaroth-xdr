%#include "common.x"

namespace mazzaroth
{
  // ChannelConfig stores contract channel configuration in state and is 
  // accessible through host contract host functions
  struct ChannelConfig
  {
    // Public Key ID of the channel owner. Only owner can change this to transfer ownership of channel
    ID owner;
    // Public Keys of IDs approved by owner able to modify channel
    ID admins<32>;
    // The Channel Governance type, controls permissions on the channel
    ChannelType channelType;
  };

  enum Governance
  {
    NONE = 0,
    PUBLIC = 1,
    PRIVATE = 2,
    PERMISSIONED = 3
  };

  union ChannelType switch (Governance Type)
  {
    case NONE:
      void;
    case PUBLIC:
      void;
    case PRIVATE:
      void;
    case PERMISSIONED:
      void;
  };
}
