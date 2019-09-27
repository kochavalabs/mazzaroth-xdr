namespace mazzaroth
{

  // Account stores account object data stored in state by ID (Public Key)
  struct Account
  {
    // Registered name of the account, separate from public key
    string name;
    // Nonce of the account tracked by Mazzaroth for ordering transactions, starts at 0
    unsigned hyper nonce;
    // List of IDs with permission to act on behalf of this account
    ID permissionedKeys<>;
  };
}
