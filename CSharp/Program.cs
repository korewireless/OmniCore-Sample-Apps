using OmniCore.Client; 

namespace Example
{
    public class Program
    {
        public static async Task Main()
        {  
            string tokenUrlPath = "https://api.korewireless.com/Api/token"; 
            Configuration config = new Configuration();
            config.BasePath = "https://api.korewireless.com/omnicore";
            config.ApiKey.Add("x-api-key", "provide-your-api-key");
            
            string token = await TokenHelper.FetchToken("", "provide-your-clientid", "provide-your-clientsecret",  tokenUrlPath);
            config.AccessToken = token;
            var subscriptionId =   "provide-your-subscription-id";  
            Console.WriteLine("Token: " + token);
            /* Registry operations */
            RegistryOps.GetRegistries(config, subscriptionId);
            //RegistryOps.GetRegistry(config, subscriptionId);
            //RegistryOps.CreateRegistry(config, subscriptionId);
            //RegistryOps.UpdateRegistry(config, subscriptionId);
            //RegistryOps.BroadcastToDevices(config, subscriptionId);

            /* Device operations */
            //DeviceOps.GetDevices(config, subscriptionId);
            //DeviceOps.GetDevice(config, subscriptionId);
            //DeviceOps.GetStates(config, subscriptionId);
            //DeviceOps.GetConfig(config, subscriptionId); 

            /* Provision Device operations */
            //DeviceOps.CreateDevice(config, subscriptionId);
            //DeviceOps.UpdateDevice(config, subscriptionId);
            //DeviceOps.DeleteDevice(config, subscriptionId);
            //DeviceOps.UpdateCustomOnboardRequest(config, subscriptionId); //custom provision

            /* Manage Device operations */
            //DeviceOps.UpdateConfiguration(config, subscriptionId);
            //DeviceOps.SendCommand(config, subscriptionId);
 
            /*
             Gateway operations */
            //GatewayOps.CreateGateway(config, subscriptionId);
            //GatewayOps.BindDevice(config, subscriptionId);
        }
    }
}