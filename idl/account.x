namespace mazzaroth
{

  // Account stores account object data stored in state by ID (Public Key)
  struct Account
  {
    // Allow account holder to set their own alias
    string alias<32>;

    // Keep track of the number write transactions sent by an account
    unsigned hyper transactionCount;

    // List of IDs with permission to act on behalf of this account
    AuthorizedAccount authorizedAccounts<32>;
  };

  struct AuthorizedAccount {
    // Public Key of the authorized account
    ID key;

    // Alias that we give to an authorized account
    string alias<32>;
  }
}
