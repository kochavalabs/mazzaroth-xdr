namespace mazzaroth
{

  // Account stores account object data stored in state by ID (Public Key)
  struct Account
  {
    // Allow account holder to set their own alias
    string alias<32>;
  };

  // Authorized stores the list of authorized IDs of a given account
  struct Authorized {
    ID accounts<>;
  }
}
