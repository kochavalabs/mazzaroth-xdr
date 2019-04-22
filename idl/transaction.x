namespace mazzaroth
{
  typedef opaque Signature[64];
  typedef opaque ID[32];
  typedef unsigned hyper uint64;

  struct Call
  {
    string function<256>;

    opaque parameters<>;
  };

  struct Update
  {
    opaque contract<>;
  };

  enum ActionCategoryType
  {
    CALL = 0,
    UPDATE = 1
  };

  union ActionCategory switch (ActionCategoryType type)
  {
    case CALL:
        Call call;
    case UPDATE:
        Update update;
  };

  struct Action 
  {
    ID channelID;    

    uint64 nonce;

    ActionCategory category;

  };

  struct Transaction
  {
    Signature signature;

    ID address;

    Action action;
  };
}
