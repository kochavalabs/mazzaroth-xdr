namespace mazzaroth
{

  // Account stores account object data stored in state by ID (Public Key)
  struct Account
  {
    // Allow account holder to set their own alias
    string alias<32>;

    // List of IDs with permission to act on behalf of this account
    AuthorizedAccount authorizedAccounts<32>;
  };

  struct AuthorizedAccount {
    // Alias that we give to an authorized account
    string alias<32>;
    
    // Public Key of the authorized account
    ID key;
  }
}
