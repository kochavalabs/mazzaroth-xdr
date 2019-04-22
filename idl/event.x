
namespace mazzaroth
{

  struct Event {
    // Name of Event (Function Name)
    string key<256>;

    opaque values<>;
  };
}
