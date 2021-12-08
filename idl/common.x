
namespace mazzaroth
{
  typedef opaque Signature[64];
  typedef opaque ID[32];
  typedef opaque Hash[32];
  typedef string Argument<>;
  typedef string StatusInfo<256>;
  typedef Authorized AuthAccount;

  enum Status
  {
    UNKNOWN = 0,
    SUCCESS = 1,
    FAILURE = 2,
    PENDING = 3,
    FINALIZED = 4
  }
}
