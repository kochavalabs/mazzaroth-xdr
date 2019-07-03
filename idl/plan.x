
namespace mazzaroth
{

  struct ExecutionPlan 
  {
    string host<256>;

    ID channelID;

    Call calls<100>;      
  };

}
