namespace mazzaroth
{
  struct Receipt
  {
    // Status failure or success
    Status status;
 
    // The state root after execution of the transaction
    Hash stateRoot;
 
    // Return results of execution if there is one for function called
    string result<>;

    // Human readable information to help understand the receipt status.
    StatusInfo statusInfo;
  };
}
