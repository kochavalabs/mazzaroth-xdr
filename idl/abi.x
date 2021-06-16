namespace mazzaroth
{
  struct Abi
  {
    // Registered name of the account, separate from public key
    FunctionSignature functions<>;
  };

  struct FunctionSignature
  {
    string functionType<>;

    string name<>;

    Parameter inputs<>;

    Parameter outputs<>;
  }

  struct Parameter 
  {
    string name<>;

    string parameterType<>;

    string codec<>;
  }
}