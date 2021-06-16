%#include "common.x"

namespace mazzaroth
{
  // ChannelConfig stores contract channel configuration in state and is 
  // accessible through host contract host functions
  struct ChannelConfig
  {
    // Public Key ID of the channel owner. Only owner can change this to transfer ownership of channel
    ID owner;
    // Human readable channel name
    string channelName<200>;
    // Public Keys of IDs approved by owner able to modify channel
    ID admins<200>;
  };
}