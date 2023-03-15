using System.Collections.Generic;
using System.Diagnostics;
using OmniCore.Api;
using OmniCore.Client;
using OmniCore.Model;

namespace Example
{
    public class GatewayOps 
    {
        internal static void CreateGateway(Configuration config , string subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "provide-registry-name";  // string | Registry ID
            var deviceId    = "provide-gateway-name";  // string | Device ID

            try
            {
                GatewayConfig gC = new GatewayConfig(GatewayConfig.GatewayAuthMethodEnum.ASSOCIATIONANDDEVICEAUTHTOKEN, GatewayConfig.GatewayTypeEnum.GATEWAY);
                Device device = new Device(
                                            deviceId,
                                            false,                    
                                            getCredentials(),
                                            null, 
                                            gC,
                                            true,
                                            null,
                                            LogLevel.ERROR,
                                            getMetaData()
                                           );
                Device result = apiInstance.CreateDevice(subscriptionId, registryId,  device);
                Console.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling DeviceApi.CreateGateway: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }  

        internal static void BindDevice(Configuration config , string subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "provide-registry-name";  // string | Registry ID
            var deviceId = "provide-device-name";  // string | Device ID
            var gatewayId = "provide-gateway-name";
            var device = new BindRequest(deviceId, gatewayId);  

            try
            {
                Info result = apiInstance.BindDevice(subscriptionId, registryId, device, 0);
                Debug.WriteLine(result);
            }
            catch (ApiException e)
            {
                Console.WriteLine("Exception when calling BindDevice: " + e.Message );
                Console.WriteLine("Status Code: "+ e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }

       
         /************* Helpers ********************/
         // info on creating key pay
         // https://docs.omnicore.cloud.korewireless.com/docs/Guides/Connect/Managing%20Credentials/create-key-pairs/index.html
        private static List<DeviceCredential> getCredentials() 
        {
            List<DeviceCredential> listDC = new List<DeviceCredential> ();
            PublicKeyCredential pKC1 = new PublicKeyCredential(PublicKeyCredential.FormatEnum.RSAX509PEM, File.ReadAllText("C:\\omnicore\\cert\\roots.pem")); 
            DeviceCredential dc1 = new DeviceCredential(DateTime.Now.AddYears(1).ToString(), pKC1);
            listDC.Add(dc1);            
            return listDC;
        }

        private static Dictionary<string, string> getMetaData()
        {
            var dc = new  Dictionary<string, string>();
            for (int i=0; i<5; i++)
            dc.Add("key" + i, "value" + i);
            return dc;
        }
    }
}