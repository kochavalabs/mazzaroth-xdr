
namespace mazzaroth
{
  // A transaction that calls a function on a user defined contract.
  struct Call
  {
    // Contract function to execute.
    string function<256>;

    // Arguments to the contract function. The serialization format is defined
    // by the contract itself.
    Argument arguments<>;
  };
}
