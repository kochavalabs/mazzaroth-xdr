namespace mazzaroth
{

  // Account stores account object data stored in state by ID (Public Key)
  struct Account
  {
    string name;
    unsigned hyper nonce;
    unsigned hyper value;
  };
}