
namespace mazzaroth
{
  struct Channel
  {
    ID channelID;

    string name<32>;

    Config configuration<>;
  };

  enum ConfigType
  {
    UNKNOWN = 0,
    ADMIN = 1
  };

  union Config switch (ConfigType Type)
  {
    case UNKNOWN:
      void;
    case ADMIN:
      ID admin<8>;
  };
}
