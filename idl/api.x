namespace mazzaroth
{
  enum RequestType
  {
    UNKNOWN = 0,

    TRANSACTION = 1
  };

  union Request switch (RequestType Type)
  {
      case UNKNOWN:
        void;
      case TRANSACTION:
        Transaction transaction;
  };

  enum ResponseType
  {
    UNKNOWN = 0,

    TRANSACTION = 1,

    TRANSACTIONLIST = 2,

    RECEIPT = 3,

    RECEIPTLIST = 4,

    BLOCK = 5,

    BLOCKLIST = 6,

    BLOCKHEADER = 7,

    BLOCKHEADERLIST = 8,

    CHANNEL = 9,

    CHANNELLIST = 10,

    ACCOUNT = 11,

    HEIGHT = 12,

    ABI = 13

  };

  union Response switch (ResponseType Type)
  {
    case UNKNOWN:
      void;
    case TRANSACTION:
      Transaction transaction;
    case TRANSACTIONLIST:
      Transaction transactions<>;
    case RECEIPT:
      Receipt receipt;
    case RECEIPTLIST:
      Receipt receipts<>;
    case BLOCK:
      Block block;
    case BLOCKLIST:
      Block blocks<>;
    case BLOCKHEADER:
      BlockHeader blockHeader;
    case BLOCKHEADERLIST:
      BlockHeader blockHeaders<>;
    case CHANNEL:
      ChannelConfig channel;
    case CHANNELLIST:
      ChannelConfig channels<>;
    case ACCOUNT:
      Account account;
    case HEIGHT:
      BlockHeight height;
    case ABI:
      Abi abi;
  };
}
