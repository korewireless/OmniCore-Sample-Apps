using OmniCore.Api;
using OmniCore.Client;
using OmniCore.Model;

namespace Example
{
    public class RegistryOps
    {
        internal static void GetRegistries(Configuration config , String subscriptionId)
        {
            var apiInstance = new RegistryApi(config);
            var registryId = "provide-registry-name"; // string | Registry ID
            var pageNumber = 1;  // int? | Page Number (optional) 
            var pageSize = 50;  // int? | Page Size (optional) 

            try
            {
                // Get All Registries
                ListDeviceRegistries result = apiInstance.GetRegistries(subscriptionId, pageNumber, pageSize);
                foreach (DeviceRegistry r in result.DeviceRegistries)
                {
                   Console.WriteLine(r.Name);
                }
            }
            catch(ApiException e) 
            {
                Console.WriteLine("Exception when calling RegistryApi.GetRegistries: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);                 
                Console.WriteLine(e.StackTrace);
            }
        }

        internal static void GetRegistry(Configuration config, String subscriptionId)
        {
            var apiInstance = new RegistryApi(config); 
            var registryId = "provide-registry-name"; // string | Registry ID

            try
            {
                // Get All Registries
                DeviceRegistry DeviceRegistry = apiInstance.GetRegistry(subscriptionId, registryId, 0);                 
                Console.WriteLine("DeviceRegistry.Name: " + DeviceRegistry.Name);
                Console.WriteLine("DeviceRegistry.CreatedOn: " + DeviceRegistry.CreatedOn);                  

                if (DeviceRegistry.Credentials !=null) 
                {
                    foreach (RegistryCredential oRegistryCredential in DeviceRegistry.Credentials)
                    {
                        Console.WriteLine("");
                        Console.WriteLine("oRegistryCredential.PublicKeyCertificate.X509Details"); 
                        Console.WriteLine(oRegistryCredential.PublicKeyCertificate.X509Details);
                        Console.WriteLine("oRegistryCredential.PublicKeyCertificate.Certificate");
                        Console.WriteLine(oRegistryCredential.PublicKeyCertificate.Certificate);
                        Console.WriteLine("oRegistryCredential.PublicKeyCertificate.Format");
                        Console.WriteLine(oRegistryCredential.PublicKeyCertificate.Format);                    
                    }
                }
                Console.WriteLine(DeviceRegistry);
            }
            catch(ApiException e) 
            {
                Console.WriteLine("Exception when calling RegistryApi.GetRegistries: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);                 
                Console.WriteLine(e.StackTrace);
            }
        }

        internal static void CreateRegistry(Configuration config, String subscriptionId)
        {
            var apiInstance = new RegistryApi(config); 
            var registryId = "provide-registry-name";  // string | Registry ID

            try
            {
                DeviceRegistry newRegistry = new DeviceRegistry(registryId,
                                                            getRegistryCredentials(),                                                            
                                                            new HttpConfig(HttpConfig.HttpEnabledStateEnum.DISABLED),
                                                            new MqttConfig(MqttConfig.MqttEnabledStateEnum.ENABLED),
                                                            LogLevel.ERROR,
                                                            getEventNotificationConfigs(),
                                                            getLogNotificationConfig(),
                                                            getStateNotificationConfig()
                                                            );  
                DeviceRegistry result = apiInstance.CreateRegistry(subscriptionId, newRegistry, 0); 
                Console.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling RegistryApi.CreateRegistry: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        } 

        internal static void UpdateRegistry(Configuration config, string subscriptionId)
        {
            var apiInstance = new RegistryApi(config); 
            var registryId = "provide-registry-name";  // string | Registry ID
            var updateMask = "credentials,logLevel";  // string | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,log_level, blocked, and metadata
    
            DeviceRegistry uR = new DeviceRegistry(registryId, getRegistryCredentials(), new HttpConfig(HttpConfig.HttpEnabledStateEnum.ENABLED), new MqttConfig(MqttConfig.MqttEnabledStateEnum.DISABLED),
                                                    LogLevel.INFO, getEventNotificationConfigs(), getLogNotificationConfig(), getStateNotificationConfig());

            try
            {
                DeviceRegistry result = apiInstance.UpdateRegistry(subscriptionId, registryId, updateMask, uR, 0);
                Console.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling RegistryApi.UpdateRegistry: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }

        internal static void BroadcastToDevices(Configuration config, string subscriptionId)
        {
            var apiInstance = new RegistryApi(config); 
            var registryId = "provide-registry-name";  // string | Registry ID

            DeviceRegistry uR = new DeviceRegistry(registryId, getRegistryCredentials(), new HttpConfig(HttpConfig.HttpEnabledStateEnum.ENABLED), new MqttConfig(MqttConfig.MqttEnabledStateEnum.DISABLED),
                                                    LogLevel.INFO, getEventNotificationConfigs(), getLogNotificationConfig(), getStateNotificationConfig());

            try
            {
                DeviceCommand dCommand = new DeviceCommand("c29tZXRoaW5nIGlzIGdvaW5nIG9u", "");  
                Object result = apiInstance.SendBroadcastToDevices(subscriptionId,  registryId, dCommand, 0);
                Console.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Console.WriteLine("Exception when calling RegistryApi.UpdateRegistry: " + e.Message);
                Console.WriteLine("Status Code: " + e.ErrorCode);
                Console.WriteLine(e.StackTrace);
            }
        }

        private static List<EventNotificationConfig> getEventNotificationConfigs()
        {
            EventNotificationConfig ev =new EventNotificationConfig();
            ev.PubsubTopicName = "projects/project-name/topics/test-topic";
            ev.SubfolderMatches = null;
            List<EventNotificationConfig> l= new List<EventNotificationConfig>();
            l.Add(ev);
            return l; 
        } 

        private static NotificationConfig getStateNotificationConfig()
        {
            NotificationConfig ev =new NotificationConfig();
            ev.PubsubTopicName = "projects/project-name/topics/test";             
            return ev;
        }

        private static NotificationConfig getLogNotificationConfig()
        {
            NotificationConfig ev =new NotificationConfig();
            ev.PubsubTopicName = "projects/project-name/topics/test-topic";            
            return ev;
        }
 
        private static List<RegistryCredential> getRegistryCredentials()
        {
            List<RegistryCredential> listDC = new List<RegistryCredential> ();
            PublicKeyCertificate pKC1 = new PublicKeyCertificate(
                                                                    File.ReadAllText("C:\\omnicore\\cert\\roots.pem"),
                                                                    PublicKeyCertificate.FormatEnum.X509CERTIFICATEPEM,
                                                                    null
                                                                ); 
            PublicKeyCertificate pKC2 = new PublicKeyCertificate(
                                                                    File.ReadAllText("C:\\omnicore\\cert\\roots.pem"),
                                                                    PublicKeyCertificate.FormatEnum.X509CERTIFICATEPEM,
                                                                    null
                                                                ); 
            RegistryCredential dc1 = new RegistryCredential(pKC1);
            RegistryCredential dc2 = new RegistryCredential(pKC2); 
            listDC.Add(dc1);
            listDC.Add(dc2);
            
            return listDC;
        }
    }
}