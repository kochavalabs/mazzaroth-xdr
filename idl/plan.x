
namespace mazzaroth
{

  // ExecutionPlan is used to define a set of transaction calls to execute as a batch
  // Used by the mazzaroth-js client
  struct ExecutionPlan 
  {
    // The host ip of the Mazzaroth node to target
    string host<256>;

    // The id of the channel for the transactions
    ID channelID;

    // The list of calls to submit to the Mazzaroth node
    Call calls<100>;      
  };

}
