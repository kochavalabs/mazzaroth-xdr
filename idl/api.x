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

    CONTRACT = 8,

    HEIGHT = 9,

    ABI = 10

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
    case CONTRACT:
      Contract Contract;
    case HEIGHT:
      BlockHeight height;
    case ABI:
      Abi abi;
  };
}
