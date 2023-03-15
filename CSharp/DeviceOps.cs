using System.Collections.Generic;
using System.Diagnostics;
using OmniCore.Api;
using OmniCore.Client;
using OmniCore.Model;

namespace Example
{
    public class DeviceOps
    {
        internal static void GetDevices(Configuration config , String subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "provide-registry-name";  

            try
            {
                // Get Device
                ListDevicesResponse devices = apiInstance.GetDevices(registryId, subscriptionId);                
                foreach (Device oDevice in devices.Devices)
                {
                    Console.WriteLine("oDevice.Id: " + oDevice.Id);
                    Console.WriteLine("oDevice.CreatedOn: " + oDevice.CreatedOn);
                    Console.WriteLine("oDevice.DeviceErrors: " + oDevice.DeviceErrors);
                    Console.WriteLine("oDevice.IsGateway: " + oDevice.IsGateway);
                    Console.WriteLine("oDevice.LastConfigAckTime: " + oDevice.LastConfigAckTime);
                    Console.WriteLine("oDevice.LastConfigSendTime: " + oDevice.LastConfigSendTime);
                    Console.WriteLine("");
                     
                    if (oDevice.Credentials !=null) 
                    {
                        foreach (DeviceCredential oDeviceCredential in oDevice.Credentials)
                        {    
                            Console.WriteLine(oDeviceCredential.PublicKey);
                            Console.WriteLine(oDeviceCredential.PublicKey.Format);
                            Console.WriteLine(oDeviceCredential.PublicKey.Key);
                            Console.WriteLine(oDeviceCredential.ExpirationTime);                    
                        }
                    }                  
                }
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling DeviceApi.GetDevice: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }
        internal static void GetDevice(Configuration config , String subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "provide-registry-name";  
            var deviceId = "provide-device-name";  

            try
            {
                // Get Device
                Device device = apiInstance.GetDevice(registryId, subscriptionId, deviceId);
                
                foreach (DeviceCredential oDeviceCredential in device.Credentials)
                {
                    Console.WriteLine(oDeviceCredential.PublicKey);
                    Console.WriteLine(oDeviceCredential.PublicKey.Format);
                    Console.WriteLine(oDeviceCredential.PublicKey.Key);
                    Console.WriteLine(oDeviceCredential.ExpirationTime);                    
                }
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling DeviceApi.GetDevice: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }
        internal static void GetStates(Configuration config , String subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "provide-registry-name";  
            var deviceId = "provide-device-name";   
            
            try
            {                
                ListDeviceStatesResponse statesResponse = apiInstance.GetStates(subscriptionId, registryId, deviceId, 0); //issue name should not have List
                Console.WriteLine("DeviceStates: " + statesResponse.DeviceStates);
                if (statesResponse.DeviceStates !=null)
                {
                    foreach (DeviceState state in statesResponse.DeviceStates)
                    {
                        Console.WriteLine("state binary: " + state.BinaryData);
                        Console.WriteLine("state updatetime: " + state.UpdateTime);
                    }
                }
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling DeviceApi.GetStates: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }
        internal static void GetConfig(Configuration config , string subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "provide-registry-name";  
            var deviceId = "provide-device-name";   
            
            try
            {
                ListDeviceConfigVersionsResponse result = apiInstance.GetConfig(subscriptionId, registryId, deviceId, 0);
                 
                foreach (DeviceConfig dc in result.DeviceConfigs)
                {
                    Console.WriteLine("_Version :" +dc._Version);
                    Console.WriteLine("Acknowledged :" + dc.Acknowledged);
                    Console.WriteLine("BinaryData :" + dc.BinaryData);
                    Console.WriteLine("DeviceAckTime :" + dc.DeviceAckTime);
                } 
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling DeviceApi.GetConfig: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }
        /************* Add ********************/
        internal static void CreateDevice(Configuration config , string subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "provide-registry-name";  
            var deviceId = "provide-device-name";    

            try
            {
                // Add New Device
                Device newDevice = new Device(
                                                deviceId,
                                                false,                    
                                                getCredentials(),
                                                null,
                                                null,
                                                false,
                                                null,
                                                LogLevel.ERROR,
                                                getMetaData()
                                           );
                Device result = apiInstance.CreateDevice(subscriptionId, registryId, newDevice);
                Console.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling DeviceApi.CreateDevice: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }
         /************* Update ********************/
        internal static void UpdateDevice(Configuration config , string subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "provide-registry-name";  
            var deviceId = "provide-device-name";  

            var updateMask = "logLevel,blocked,credentials,metadata";  // string | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,log_level, blocked, and metadata
            
            try
            {
            // Modify Device
                var device =    new Device( "",
                                            false,
                                            getCredentials(),
                                            null,
                                            null,
                                            false,
                                            null,
                                            LogLevel.INFO,
                                            getMetaData()
                                            );                  
                Device result = apiInstance.UpdateDevice(subscriptionId, registryId, deviceId, updateMask, device);
                Console.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling DeviceApi.UpdateDevice: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }
        /************* Delete ********************/
        internal static void DeleteDevice(Configuration config , string subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "provide-registry-name";  
            var deviceId = "provide-device-name"; 

            try
            {   
                Info result = apiInstance.DeleteDevice(subscriptionId, registryId, deviceId);
                Console.WriteLine("Info: " + result._Info);
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling DeviceApi.DeleteDevice: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }
         /************** Configuration ******************/
        internal static void UpdateConfiguration(Configuration config , string subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "provide-registry-name";  
            var deviceId = "provide-device-name"; 
            var deviceConfiguration = new DeviceConfiguration("c29tZXRoaW5nIGlzIGdvaW5nIG9u", "sd", "0");    

            try
            {                
                Object result = apiInstance.UpdateConfigurationToDevice(subscriptionId, registryId, deviceId, deviceConfiguration, 0);
                Console.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling DeviceApi.UpdateConfiguration: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }
        internal static void SendCommand(Configuration config , string subscriptionId)
        {
            var apiInstance = new DeviceApi(config); 
            var registryId = "meta-registry";  
            var deviceId = "provide-device-name";  
            DeviceCommand dCommand = new DeviceCommand("c29tZXRoaW5nIGlzIGdvaW5nIG9u", "sd");    

            try
            {                 
                Object result = apiInstance.SendCommandToDevice(subscriptionId, registryId, deviceId,  dCommand, 0);
                Console.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling DeviceApi.ModifyCloudToDeviceConfig: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
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
            for (int i=0; i<10; i++)
                dc.Add("key" + i, "value" + i);
            return dc;
        }
    }
}