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

    Hash transactionID;
  };

  enum ValueFilterType
  {
    NONE = 0,
    STRING = 1,
    HASH32 = 2,
    HASH64 = 3,
    UHYPER = 4,
    BOOL = 5,
  };

  union ValueFilter switch (ValueFilterType Type) 
  {
    case NONE:
      void;
    case STRING:
      string regex<256>;
    case HASH32:
      Hash hash32Value;
    case HASH64:
      Hash hash64Value;
    case UHYPER:
      unsigned hyper uhyperVaue;
    case BOOL:
      bool boolValue;
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
    CONTRACT = 1,
    CONFIG = 2,
    PERMISSION = 3,
    CALL = 4,
  };

  union TransactionFilter switch (TransactionFilterType Type)
  {
    case NONE:
      void;
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

