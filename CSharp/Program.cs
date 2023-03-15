using OmniCore.Client; 

namespace Example
{
    public class Program
    {
        public static void Main()
        {
           
            Configuration config = new Configuration();
            config.BasePath = "https://api.omnicore.cloud.korewireless.com";
            config.ApiKey.Add("x-api-key", "provide-api-key");
            config.AccessToken = "provide-access-token";
            var subscriptionId =   "korewireless-staging";  
            
            /* Registry operations */
            //RegistryOps.GetRegistries(config, subscriptionId);
            //RegistryOps.GetRegistry(config, subscriptionId);
            //RegistryOps.CreateRegistry(config, subscriptionId);
            //RegistryOps.UpdateRegistry(config, subscriptionId);
            
            /* Device operations */
            //DeviceOps.GetDevices(config, subscriptionId);
            //DeviceOps.GetDevice(config, subscriptionId);
            //DeviceOps.GetStates(config, subscriptionId);
            //DeviceOps.GetConfig(config, subscriptionId); 
            
            /* Provision Device operations */
            //DeviceOps.CreateDevice(config, subscriptionId);
            //DeviceOps.UpdateDevice(config, subscriptionId);
            //DeviceOps.DeleteDevice(config, subscriptionId);

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