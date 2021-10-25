namespace mazzaroth
{
  struct Abi
  {
    // SemVer version of this ABI Schema
    // Will match the mazzaroth-xdr version used to generate ABI
    string version;

    // Registered name of the account, separate from public key
    FunctionSignature functions<>;
  };

  struct FunctionSignature
  {
    FunctionType functionType;

    string functionName<>;

    Parameter parameters<>;

    Parameter returns<>;
  }

  enum FunctionType
  {
    UNKNOWN = 0,
    READ = 1,
    WRITE = 2
  }

  struct Parameter 
  {
    string parameterName<>;

    string parameterType<>;
  }
}
