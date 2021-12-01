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

    RECEIPT = 3,

    BLOCK = 4,

    BLOCKLIST = 5,

    BLOCKHEADER = 6,

    BLOCKHEADERLIST = 7,

    CONFIG = 8,

    ACCOUNT = 9,

    HEIGHT = 10,

    ABI = 11

  };

  union Response switch (ResponseType Type)
  {
    case UNKNOWN:
      void;
    case TRANSACTIONID:
      ID transactionID;
    case TRANSACTION:
      Transaction transaction;
    case RECEIPT:
      Receipt receipt;
    case BLOCK:
      Block block;
    case BLOCKLIST:
      Block blocks<>;
    case BLOCKHEADER:
      BlockHeader blockHeader;
    case BLOCKHEADERLIST:
      BlockHeader blockHeaders<>;
    case CONFIG:
      Config config;
    case ACCOUNT:
      Account account;
    case HEIGHT:
      BlockHeight height;
    case ABI:
      Abi abi;
  };
}
