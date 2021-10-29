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

    TRANSACTIONID = 1,

    TRANSACTION = 2,

    TRANSACTIONLIST = 3,

    RECEIPT = 4,

    RECEIPTLIST = 5,

    BLOCK = 6,

    BLOCKLIST = 7,

    BLOCKHEADER = 8,

    BLOCKHEADERLIST = 9,

    CHANNEL = 10,

    CHANNELLIST = 11,

    ACCOUNT = 12,

    HEIGHT = 13,

    ABI = 14

  };

  union Response switch (ResponseType Type)
  {
    case UNKNOWN:
      void;
    case TRANSACTIONID:
      ID transactionID;
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
