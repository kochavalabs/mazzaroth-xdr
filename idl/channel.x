%#include "common.x"

namespace mazzaroth
{
  // ChannelConfig stores channel configuration in state
  struct ChannelConfig
  {
    // ID (public Key) of the owner of the channel
    ID owner;
    // ID (public key) of the validators participating in consensus
    ID validators<>;
    // Consensus Config type for this channel
    ConsensusConfig consensusConfig;
  };

  // ConsensusConfigType defines the options for consensus in a channel
  // currently only PBFT is supported
  enum ConsensusConfigType
  {
    NONE = 0,
    PBFT = 1
  };

  union ConsensusConfig switch (ConsensusConfigType Type)
  {
  case NONE:
      void;
  case PBFT:
      PBFTConfig pbftConfig;
  };

  // Configuration for PBFT to store in state
  struct PBFTConfig
  {
    // number of sequence numbers between checkpoints
    unsigned hyper checkpointPeriod;
  };
}