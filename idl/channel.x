%#include "common.x"

namespace mazzaroth
{
  // ChannelConfig stores channel configuration in state
  struct ChannelConfig
  {
    ID channelID;
    ID owner;
    ConsensusConfig consensusConfig;
  };

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
    // List of currently selected validators participating in pbft consensus
    ID validators<>;
    // Number of sequence numbers between checkpoints
    unsigned hyper checkpointPeriod;
    // Max number of sequence number to accept while waiting on checkpoint
    unsigned hyper watermarkRange;
  };
}