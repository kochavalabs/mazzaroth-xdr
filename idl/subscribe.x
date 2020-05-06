namespace mazzaroth
{
  struct ReceiptSubscription
  {
    TransactionFilter transactionFilter;
 
    ReceiptFilter receiptFilter;
  };

  struct ReceiptSubscriptionResult
  {
    Receipt receipt;

    ID transactionID;
  };

  enum ValueFilterType
  {
    NONE = 0,
    STRING = 1,
    HASH32 = 2,
    HASH64 = 3,
    UHYPER = 4,
    INT = 5,
  };

  union ValueFilter switch (ValueFilterType Type) 
  {
    case NONE:
      void;
    case STRING:
      string regex<256>;
    case HASH32:
      Hash32 hash32Value;
    case HASH64:
      Hash64 hash64Value;
    case UHYPER:
      unsigned hyper uhyperValue;
    case INT:
      int intValue;
  };

  struct ReceiptValueFilter 
  {
    ValueFilter status;

    ValueFilter stateRoot;
  };

  struct ActionFilter
  {
    ValueFilter signature; 

    ValueFilter signer;

    ValueFilter address;

    ValueFilter channelID;

    ValueFilter nonce;
  };

  struct ContractFilter
  {
    ActionFilter actionFilter;

    ValueFilter version;
  };

  struct ConfigFilter 
  {
    ActionFilter actionFilter;
  };

  struct PermissionFilter
  {
    ActionFilter actionFilter;

    ValueFilter key;

    ValueFilter action;
  };

  struct CallFilter
  {
    ActionFilter actionFilter;

    ValueFilter function;
  };

  enum TransactionFilterType
  {
    NONE = 0,
    GENERIC = 1,
    CONTRACT = 2,
    CONFIG = 3,
    PERMISSION = 4,
    CALL = 5,
  };

  union TransactionFilter switch (TransactionFilterType Type)
  {
    case NONE:
      void;
    case GENERIC:
      ActionFilter genericFilter;
    case CONTRACT:
      ContractFilter contractFilter; 
    case CONFIG:
      ConfigFilter configFilter; 
    case PERMISSION:
      PermissionFilter permissionFilter; 
    case CALL:
      CallFilter callFilter; 
  };

  enum ReceiptFilterType
  {
    NONE = 0,
    RECEIPT = 1,
  };

  union ReceiptFilter switch (ReceiptFilterType Type)
  {
    case NONE:
      void;
    case RECEIPT:
      ReceiptValueFilter valueFilter; 
  };
}

