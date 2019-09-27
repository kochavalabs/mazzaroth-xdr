
namespace mazzaroth
{

  // Event is a message object for sending out events defined during contract execution
  struct Event {
    // Name of Event (Function Name)
    string key<256>;
    // Parameters supplied to event call
    Parameter parameters<>;
  };
}
