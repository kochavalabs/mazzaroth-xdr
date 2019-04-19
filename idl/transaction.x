namespace mazzaroth
{
  typedef opaque Signature[64];
  typedef opaque ID[32];

  struct Call
  {
    string function<256>;

    opaque parameters<>;
  };

  struct Update
  {
    opaque contract<>;
  };

  enum ActionType
  {
    ACTION_TYPE_CALL = 0,
    ACTION_TYPE_UPDATE = 1
  };

  union Action switch (ActionType type)
  {
    case ACTION_TYPE_CALL:
        Call call;
    case ACTION_TYPE_UPDATE:
        Update update;
  };

  struct Transaction
  {
    Signature signature;

    ID address;

    Action action;
  };
}
