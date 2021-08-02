namespace mazzaroth
{

  // Account stores account object data stored in state by ID (Public Key)
  struct Account
  {
    // List of IDs with permission to act on behalf of this account
    ID permissionedKeys<32>;
  };
}
