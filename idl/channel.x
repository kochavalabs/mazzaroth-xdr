%#include "common.x"

namespace mazzaroth
{
  // ChannelConfig stores channel configuration in state
  struct ChannelConfig
  {
    ID owner;
    ID validators<>;
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
    // number of sequence numbers between checkpoints
    unsigned hyper checkpointPeriod;
  };
}