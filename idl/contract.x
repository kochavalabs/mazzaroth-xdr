namespace mazzaroth
{
  // An update transaction that provides a contract as a wasm binary.
  struct Contract
  {
    // Version number of the contract, specified by owner
    string version<100>;

    // Contract ABI
    Abi abi;

    // Sha3 256 Hash of the contract bytes, verified on execution
    Hash contractHash;

    // Contract binary bytes.
    opaque contractBytes<>;
  };
}