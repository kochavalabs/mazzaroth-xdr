namespace mazzaroth
{
  struct Receipt
  {
    // ID of the transaction the receipt belongs to
    ID transactionID; 
    
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
