var OmnicoreModelAndStateManagementApi = require('omni_core_model_and_state_management_api');

module.exports.createRegistry = (token,subId,regId,hostUrl,key) => {
  let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
  defaultClient.basePath = hostUrl
  // Configure API key authorization: apiKey
  var apiKey = defaultClient.authentications['apiKey'];
  apiKey.apiKey = key
  // Configure Bearer (JWT) access token for authorization: bearerAuth
  let bearerAuth = defaultClient.authentications['bearerAuth'];
  bearerAuth.accessToken = token
  
  let apiInstance = new OmnicoreModelAndStateManagementApi.RegistryApi();
  let subscriptionId = subId; // String | Subscription ID
  let registry = {
  id: regId, //String | Registry Id
  mqttConfig: {
    mqttEnabledState: "MQTT_ENABLED" //String | Mqtt Config
  },
  httpConfig: {
    httpEnabledState: "HTTP_ENABLED" //String | http Config
  },
  logLevel: "INFO", //String | Mqtt Config
  stateNotificationConfig: {
    pubsubTopicName: ""     // String | State pubsub Topic
  },
  logNotificationConfig: {
    pubsubTopicName: ""     // String | Log Notification pubsub Topic
  },
  customOnboardEnabled: false,
  customOnboardNotificationConfig: {
    pubsubTopicName: ""
  },
  isNatsRout: false, // If the field is false, the message is routs to Cloud Pub/Sub, if true, message routes to The Marina (OmniCore Stream Analytics)
  eventNotificationConfigs: [
    {
      pubsubTopicName: "" // String | Event Notification pubsub Topic
    }
  ],
  credentials: [
    {
      publicKeyCertificate: {
        format: "X509_CERTIFICATE_PEM",  // String Certificate Format
        certificate: "-----BEGIN CERTIFICATE-----\nMIIDLDCCAhSgAwIBAgITMNZHEBss/CrpIrK9eX8aMQtGNDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzIw\nOVoXDTMyMDgwMTEwMzIwOFowHjENMAsGA1UEChMEa29yZTENMAsGA1UEAxMEa29y\nZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL69ASS6Fz8EXBI1yPqm\nKASYIpyOVZN5VcMnV1spqn7uLtVHVnyeLuQon+aHTqb59TjafP+d7az8hyByL2Zm\n98juwo9vCnT4rKceskZM55olnfF07dbd+DpIT/3hbDh2RlU+rSRYwYoXZ3DOF7jd\nLsEJN9tFpExInAk5GK0ZaPXiZzCY+a9hjg+6RSPPqBhoTWuI6zPrPYeWSTKHvqNG\ngWiytwriss3hBp34Gdg4sO8COD+0uf9/Ia0hB/tpcr0Myyshgzw/SqCDK4mWVqNs\nLcsBeqZkPwNvbC3ZFvUI8NJMpeAyghW25qvswyQZWxIsvJgMI9SmSXgYCBf0DyFV\nGH8CAwEAAaNjMGEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYD\nVR0OBBYEFMo7tLXQIKicZWBlmWLs1UByKb8DMB8GA1UdIwQYMBaAFMo7tLXQIKic\nZWBlmWLs1UByKb8DMA0GCSqGSIb3DQEBCwUAA4IBAQCS7aFBnvAZ66HEC5ssutgC\nJ0Ak0Li7x8YThz0+AqWyU42/8woitwuxmyQOYP4g9CJty6bqM1LoKc3bOvb1GOMm\nZ0xhAe/+H3ZIKs6g5zXon7mZOEpGJSZQLMuyPI5searFIp+3mgIo4UpgcjYFrBTE\nYJezh93GxirAFVUQ2ZcOltvt13LiavjATlSNNwhSNKZ7lUzD2Y9d5VBkPIYBZw4U\neeJ4GnDPg3IX62rJWWpJ/unJQcmxTwjY8CT85P8C0oROjqCJc93dm8aNXpw7afVq\n+0R+VZwrzkHCmfJEV6ogrR0fUQODUWSvLmR8Z956s/OijpSqRmr88mzq1UZtbnwA\n-----END CERTIFICATE-----" // String | Certificate
      }
    }
  ]
  } // NewRegistry | application/json
  let registryData = {
    registry,
  }
  apiInstance.createRegistry(subscriptionId, registryData, (error, data, response) => {
    if (error) {
      console.error(error);
    } else {
      console.log('\nAPI called successfully. Returned data: \n');
      console.log(" Registry Id: ",data.id)
      console.log(" Parent: ",data.parent)
      console.log(" Credentials: ",data.credentials)
      console.log(" eventNotificationConfigs:",data.eventNotificationConfigs)
      console.log(" httpConfig: ",data.httpConfig)
      console.log(" logLevel: ",data.logLevel)
      console.log(" logNotificationConfig: ",data.logNotificationConfig)
      console.log(" customOnboardEnabled: ",data.customOnboardEnabled)
      console.log(" customOnboardNotificationConfig: ",data.customOnboardNotificationConfig)
      console.log(" mqttConfig: ",data.mqttConfig)
      console.log(" name:",data.name)
      console.log(" numberOfDevices: ",data.numberOfDevices)
      console.log(" numberOfGateways: ",data.numberOfGateways)
      console.log(" stateNotificationConfig: ",data.stateNotificationConfig)
      console.log(" updatedOn: ",data.updatedOn)
    }
  }); 
}

module.exports.getRegistry = (token,subId,regId,hostUrl,key) => {
  let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
  defaultClient.basePath = hostUrl
  // Configure API key authorization: apiKey
  var apiKey = defaultClient.authentications['apiKey'];
  apiKey.apiKey = key
  // Configure Bearer (JWT) access token for authorization: bearerAuth
  let bearerAuth = defaultClient.authentications['bearerAuth'];
  bearerAuth.accessToken = token

  let apiInstance = new OmnicoreModelAndStateManagementApi.RegistryApi();
  let subscriptionId = subId; // String | Subscription ID
  let registryId = regId; // String | Registry ID

  apiInstance.getRegistry(subscriptionId, registryId, (error, data, response) => {
    if (error) {
      console.error(error);
    } else {
      console.log('\nAPI called successfully. Returned data: \n');
      console.log(" Registry Id: ",data.id)
      console.log(" Parent: ",data.parent)
      console.log(" Credentials: ",data.credentials)
      console.log(" eventNotificationConfigs:",data.eventNotificationConfigs)
      console.log(" httpConfig: ",data.httpConfig)
      console.log(" logLevel: ",data.logLevel)
      console.log(" logNotificationConfig: ",data.logNotificationConfig)
      console.log(" customOnboardEnabled: ",data.customOnboardEnabled)
      console.log(" customOnboardNotificationConfig: ",data.customOnboardNotificationConfig)
      console.log(" mqttConfig: ",data.mqttConfig)
      console.log(" name:",data.name)
      console.log(" numberOfDevices: ",data.numberOfDevices)
      console.log(" numberOfGateways: ",data.numberOfGateways)
      console.log(" stateNotificationConfig: ",data.stateNotificationConfig)
      console.log(" updatedOn: ",data.updatedOn)
    }
  });
}

module.exports.getRegistries = (token,subId,regId,hostUrl,key) => {
  let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
  defaultClient.basePath = hostUrl
  // Configure API key authorization: apiKey
  var apiKey = defaultClient.authentications['apiKey'];
  apiKey.apiKey = key
  // Configure Bearer (JWT) access token for authorization: bearerAuth
  let bearerAuth = defaultClient.authentications['bearerAuth'];
  bearerAuth.accessToken = token

  let apiInstance = new OmnicoreModelAndStateManagementApi.RegistryApi();
  let subscriptionId = subId; // String | Subscription ID
  let opts = {
    'pageNumber': 1, // Number | Page Number
    'pageSize': 10 // Number | Page Size
  };
  apiInstance.getRegistries(subscriptionId, opts, (error, data, response) => {
    if (error) {
      console.error(error);
    } else {
      console.log('\nAPI called successfully. Returned data: \n');
      console.log(" Registry Id: ",data.deviceRegistries[0].id)
      console.log(" Parent: ",data.deviceRegistries[0].parent)
      console.log(" Credentials: ",data.deviceRegistries[0].credentials)
      console.log(" eventNotificationConfigs:",data.deviceRegistries[0].eventNotificationConfigs)
      console.log(" httpConfig: ",data.deviceRegistries[0].httpConfig)
      console.log(" logLevel: ",data.deviceRegistries[0].logLevel)
      console.log(" logNotificationConfig: ",data.deviceRegistries[0].logNotificationConfig)
      console.log(" customOnboardEnabled: ",data.deviceRegistries[0].customOnboardEnabled)
      console.log(" customOnboardNotificationConfig: ",data.deviceRegistries[0].customOnboardNotificationConfig)
      console.log(" mqttConfig: ",data.deviceRegistries[0].mqttConfig)
      console.log(" name:",data.deviceRegistries[0].name)
      console.log(" numberOfDevices: ",data.deviceRegistries[0].numberOfDevices)
      console.log(" numberOfGateways: ",data.deviceRegistries[0].numberOfGateways)
      console.log(" stateNotificationConfig: ",data.deviceRegistries[0].stateNotificationConfig)
      console.log(" updatedOn: ",data.deviceRegistries[0].updatedOn)
    }
  });
}

module.exports.updateRegistry = (token,subId,regId,hostUrl,key) => {
  let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
  defaultClient.basePath = hostUrl
  // Configure API key authorization: apiKey
  var apiKey = defaultClient.authentications['apiKey'];
  apiKey.apiKey = key
  // Configure Bearer (JWT) access token for authorization: bearerAuth
  let bearerAuth = defaultClient.authentications['bearerAuth'];
  bearerAuth.accessToken = token

  let apiInstance = new OmnicoreModelAndStateManagementApi.RegistryApi();
  let subscriptionId = subId; // String | Subscription ID
  let registryId = regId; // String | Registry ID
  let updateMask = 'eventNotificationConfigs, stateNotificationConfig.pubsub_topic_name, logNotificationConfig.pubsub_topic_name, mqttConfig.mqtt_enabled_state, httpConfig.http_enabled_state, logLevel, credentials; // String | values to be updated: eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,customOnboardEnabled,customOnboardNotificationConfig.pubsub_topic_name'
  let registry = {
    mqttConfig: {
      mqttEnabledState: "MQTT_ENABLED" //String | Mqtt Config
    },
    httpConfig: {
        httpEnabledState: "HTTP_DISABLED" //String | Http Config 
    },
    logLevel: "ERROR", // String | Loglevel
    stateNotificationConfig: {
      pubsubTopicName: ""  // String | State Pubsub topic
    },
    logNotificationConfig: {
      pubsubTopicName: "" // String | Log notification Pubsub topic 
    },
    customOnboardEnabled: false,
    customOnboardNotificationConfig: {
      pubsubTopicName: ""
		},
    eventNotificationConfigs: [
      {
        pubsubTopicName: "" // String | Event Notification Pusub topic
      }
    ]
  }
  let opts = {
    registry,
  }

  apiInstance.updateRegistry(subscriptionId, registryId, updateMask, opts, (error, data, response) => {
    if (error) {
      console.error(error);
    } else {
      console.log('\nAPI called successfully. Returned data: \n');
      console.log(" Registry Id: ",data.id)
      console.log(" Parent: ",data.parent)
      console.log(" Credentials: ",data.credentials)
      console.log(" eventNotificationConfigs:",data.eventNotificationConfigs)
      console.log(" httpConfig: ",data.httpConfig)
      console.log(" logLevel: ",data.logLevel)
      console.log(" logNotificationConfig: ",data.logNotificationConfig)
      console.log(" customOnboardEnabled: ",data.customOnboardEnabled)
      console.log(" customOnboardNotificationConfig: ",data.customOnboardNotificationConfig)
      console.log(" mqttConfig: ",data.mqttConfig)
      console.log(" name:",data.name)
      console.log(" numberOfDevices: ",data.numberOfDevices)
      console.log(" numberOfGateways: ",data.numberOfGateways)
      console.log(" stateNotificationConfig: ",data.stateNotificationConfig)
      console.log(" updatedOn: ",data.updatedOn)
    }
  });
}

module.exports.deleteRegistry = (token,subId,regId,hostUrl,key)=>{
  let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
  defaultClient.basePath = hostUrl
  // Configure API key authorization: apiKey
  var apiKey = defaultClient.authentications['apiKey'];
  apiKey.apiKey = key
  // Configure Bearer (JWT) access token for authorization: bearerAuth
  let bearerAuth = defaultClient.authentications['bearerAuth'];
  bearerAuth.accessToken = token
  
  let apiInstance = new OmnicoreModelAndStateManagementApi.RegistryApi();
  let subscriptionId = subId; // String | Subscription ID
  let registryId = regId; // String | Registry ID
  apiInstance.deleteRegistry(subscriptionId, registryId, (error, data, response) => {
    if (error) {
      console.error(error);
    } else {
      console.log('API called successfully. Returned data: ' + data.info);
    }
  });
}

module.exports.sendBroadcast = async (token, subscriptionId, registryId, hostUrl, key) => {
  let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
  defaultClient.basePath = hostUrl
  // Configure API key authorization: apiKey
  let apiKey = defaultClient.authentications['apiKey'];
  apiKey.apiKey = key
  // Configure Bearer (JWT) access token for authorization: bearerAuth
  let bearerAuth = defaultClient.authentications['bearerAuth'];
  bearerAuth.accessToken = token

  let deviceCommand = {
    binaryData:"aGVsbG8gd29ybGQ=", //Base64 encoded String | binary data
    subfolder:""
  }; // DeviceCommand | application/json

  apiInstance.sendBroadcastToDevices(subscriptionId, registryId, deviceCommand, (error, data) => {
    if (error) {
      console.error(error);
    } else {
      console.log('API called successfully. Returned data: ' + data);
    }
  });
}
