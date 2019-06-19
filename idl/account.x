namespace mazzaroth
{

  // Account stores account object data stored in state by ID (Public Key)
  struct Account
  {
    string name;
    unsigned hyper nonce;
    // Storage for an account can hold various structs
    // Must be ordered by key to ensure deterministic serialization
    StorageItem storage<>;
  };

  // StorageItem stores an object in account Storage
  struct StorageItem
  {
    string key;
    opaque value;
  }
}
