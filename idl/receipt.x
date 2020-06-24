namespace mazzaroth
{
  struct Receipt
  {
    // Status failure or success
    ReceiptStatus status;
 
    // The state root after execution of the transaction
    Hash stateRoot;
 
    // Return results of execution if there is one for function called
    string result<>;
  };

  enum ReceiptStatus
  {
    FAILURE = 0,
    SUCCESS = 1
  };
}
