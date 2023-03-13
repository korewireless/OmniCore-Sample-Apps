var OmnicoreModelAndStateManagementApi = require('omni_core_model_and_state_management_api');;

module.exports.createDevice =(token,subId,regId,devId,hostUrl)=>{
let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
defaultClient.basePath = hostUrl // API url end point
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = token
let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
let subscriptionId = subId; // String | Subscription ID
let registryId = regId; // String | Registry ID
let deviceObj = {
    id:devId,
    credentials:[
      {
        expirationTime: "",
        publicKey: {
            format: "RSA_X509_PEM",
            key: "-----BEGIN CERTIFICATE-----\nMIIDujCCAqKgAwIBAgITaOszYpBme+SRHZUkFWLgDs7EMDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzQy\nMloXDTMyMDgwMTEwMzIwOFowADCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAMFvTBHPdgH7+5wVlUnEdIS/0a4p9fkVzMdEMdDVr5s62VoGO7nZWxMCaxxU\nXqQiGuX3N7SINyD7h8LI8CxQsn5zyDda3QVNGU7I96iWjwzOYJmNHAN1nRI2hRDY\n8fJoQgTZI+IiRDBmgkmL9yjTY04qY7UP8zpofuMnKRuTwP6Ey1eFEMBqFfvgwrVl\niLNcq9At0bd/vlQ0VUnKV6oKqSTq9ZDPB6Cxu5amhejVwTeE6p5GGmiKw5vskmtB\ndGNgsom1K/pJdOMes8lODVp00tIVnsplL3jLgrWfbfCPALRnGz/C5XlKW8fNKEuW\nqFw2Lhnk51dtobw/oBo7vJcx2w0CAwEAAaOCAQ0wggEJMA4GA1UdDwEB/wQEAwIF\noDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQUiRS8X3OWpJNpjQoJ+22xGb3xXTAf\nBgNVHSMEGDAWgBTKO7S10CConGVgZZli7NVAcim/AzCBjQYIKwYBBQUHAQEEgYAw\nfjB8BggrBgEFBQcwAoZwaHR0cDovL3ByaXZhdGVjYS1jb250ZW50LTYyZTM5YmRh\nLTAwMDAtMjI0My1hNjFhLTNjMjg2ZDRlZWUwYS5zdG9yYWdlLmdvb2dsZWFwaXMu\nY29tL2MwZGQxZjg3ZDcwZGZhMDEwNGEwL2NhLmNydDAZBgNVHREBAf8EDzANggtn\nYWRnZW9uLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAK82b/xGn8B6Nfogw0myKjy3O\nWg53YPXuct3E04qRmD3JJtFzpSkjI2WyRmIkRX1b5SKF+ImOmGzvENZDkjT/Y2I/\nnsBL639OlXnz/+GYSq4rL6fVxXistP4LGA+khoBYSfHFZb7EYoVOYJFzZjnvJbtz\n7XG0jTMeHo8KhCBPxrNWkOERrcc7OWqREldQ36yg7zdbRLjDOjeD6FByTrpRhbDC\n0AeozF9ug9W/gPYtnkI++ksUqjJcV06uGd+9XLJPGcjH0Bai1alxROh+dkWx6TcB\nHC94el4KR6EJijMvylmnOyKHedmYaDvb52+B6zXTW9rQkh1UycmONlmAlA3OeQ==\n-----END CERTIFICATE-----"
        }
      }
    ], // Json | Credentials
    blocked:false, // Boolean | Blocked
    metadata:{},   // Json | Metadata
    loglevel:"INFO", // String | Loglevel
	gatewayConfig:{
        gatewayType:"NON_GATEWAY",   // String | Gateway Type
       gatewayAuthMethod: "GATEWAY_AUTH_METHOD_UNSPECIFIED" // String | Gateway Auth Method
    }   
}
apiInstance.createDevice(subscriptionId, registryId, deviceObj, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: \n');
    console.log('Device Id: ',data.id);
    console.log('Parent: ',data.parent);
    console.log('Registry Id:',data.registry);
    console.log('Subscription Id:',data.subscription);
    console.log('blocked: ',data.blocked);
    console.log('capresent: ',data.capresent);
    console.log('clientOnline: ',data.clientOnline);
    console.log('config:',data.config);
    console.log('createdOn:',data.createdOn);
    console.log('credentials: ',data.credentials);
    console.log('deviceErrors: ',data.deviceErrors);
    console.log('gateway: ',data.gateway);
    console.log('gatewayConfig: ',data.gatewayConfig);
    console.log('isGateway:',data.isGateway);
    console.log('lastConfigAckTime:',data.lastConfigAckTime);
    console.log('lastConfigSendTime: ',data.lastConfigSendTime);
    console.log('lastErrorStatus: ',data.lastErrorStatus)
    console.log('lastErrorTime: ',data.lastErrorTime);
    console.log('lastEventTime: ',data.lastEventTime);
    console.log('lastHeartbeatTime: ',data.lastHeartbeatTime);
    console.log('lastStateTime: ',data.lastStateTime);
    console.log('logLevel:',data.logLevel);
    console.log('metadata:',data.metadata);
    console.log('name: ',data.name);
    console.log('state: ',data.state);
    console.log('subscriptions: ',data.subscriptions);
    console.log('updatedOn:',data.updatedOn);
  }
});
}

module.exports.getDevice =(token,subId,regId,devId,hostUrl)=>{
    let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
    defaultClient.basePath = hostUrl
    // Configure Bearer (JWT) access token for authorization: bearerAuth
    let bearerAuth = defaultClient.authentications['bearerAuth'];
    bearerAuth.accessToken = token
    
    let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
    let registryId = regId; // String | Registry ID
    let subscriptionId = subId; // String | Subscription ID
    let deviceId = devId; // String | Device ID
    apiInstance.getDevice(registryId, subscriptionId, deviceId, (error, data, response) => {
      if (error) {
        console.error(error);
      } else {
        console.log('API called successfully. Returned data: \n');
        console.log('Device Id: ',data.id);
        console.log('Parent: ',data.parent);
        console.log('Registry Id:',data.registry);
        console.log('Subscription Id:',data.subscription);
        console.log('blocked: ',data.blocked);
        console.log('capresent: ',data.capresent);
        console.log('clientOnline: ',data.clientOnline);
        console.log('config:',data.config);
        console.log('createdOn:',data.createdOn);
        console.log('credentials: ',data.credentials);
        console.log('deviceErrors: ',data.deviceErrors);
        console.log('gateway: ',data.gateway);
        console.log('gatewayConfig: ',data.gatewayConfig);
        console.log('isGateway:',data.isGateway);
        console.log('lastConfigAckTime:',data.lastConfigAckTime);
        console.log('lastConfigSendTime: ',data.lastConfigSendTime);
        console.log('lastErrorStatus: ',data.lastErrorStatus)
        console.log('lastErrorTime: ',data.lastErrorTime);
        console.log('lastEventTime: ',data.lastEventTime);
        console.log('lastHeartbeatTime: ',data.lastHeartbeatTime);
        console.log('lastStateTime: ',data.lastStateTime);
        console.log('logLevel:',data.logLevel);
        console.log('metadata:',data.metadata);
        console.log('name: ',data.name);
        console.log('numId: ',data.numId);
        console.log('state: ',data.state);
        console.log('subscriptions: ',data.subscriptions);
        console.log('updatedOn:',data.updatedOn);
      }
    });
}

module.exports.updateDevice = (token,subId,regId,devId,hostUrl) => {
    let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
    defaultClient.basePath = hostUrl
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = token

let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
let subscriptionId = subId; // String | Subscription ID
let registryId = regId; // String | Registry ID
let deviceId = devId; // String | Device ID
let updateMask = "blocked,metadata,logLevel"; // String | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,log_level, blocked, and metadata
let device = {
    id: devId, // String | Device ID
    credentials: [], // Json | Credentials
    blocked: false, // Boolean | Blocked
    metadata: {}, // Json | Metadata
    logLevel: "ERROR", // String | Loglevel
    gatewayConfig: {
        gatewayAuthMethod: "GATEWAY_AUTH_METHOD_UNSPECIFIED", // String | Gateway Auth Method
        gatewayType: "NON_GATEWAY" // String | Gateway Type
    }
} // device | application/json
apiInstance.updateDevice(subscriptionId, registryId, deviceId, updateMask, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: \n');
    console.log('Device Id: ',data.id);
    console.log('Parent: ',data.parent);
    console.log('Registry Id:',data.registry);
    console.log('Subscription Id:',data.subscription);
    console.log('blocked: ',data.blocked);
    console.log('capresent: ',data.capresent);
    console.log('clientOnline: ',data.clientOnline);
    console.log('config:',data.config);
    console.log('createdOn:',data.createdOn);
    console.log('credentials: ',data.credentials);
    console.log('deviceErrors: ',data.deviceErrors);
    console.log('gateway: ',data.gateway);
    console.log('gatewayConfig: ',data.gatewayConfig);
    console.log('isGateway:',data.isGateway);
    console.log('lastConfigAckTime:',data.lastConfigAckTime);
    console.log('lastConfigSendTime: ',data.lastConfigSendTime);
    console.log('lastErrorStatus: ',data.lastErrorStatus)
    console.log('lastErrorTime: ',data.lastErrorTime);
    console.log('lastEventTime: ',data.lastEventTime);
    console.log('lastHeartbeatTime: ',data.lastHeartbeatTime);
    console.log('lastStateTime: ',data.lastStateTime);
    console.log('logLevel:',data.logLevel);
    console.log('metadata:',data.metadata);
    console.log('name: ',data.name);
    console.log('numId: ',data.numId);
    console.log('state: ',data.state);
    console.log('subscriptions: ',data.subscriptions);
    console.log('updatedOn:',data.updatedOn);
  }
});
}

module.exports.getDevices = (token,subId,regId,devId,hostUrl) => {
let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
defaultClient.basePath = hostUrl   
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = token

let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
let registryId = regId; // String | Registry ID
let subscriptionId = subId; // String | Subscription ID
let opts = {
  'pageNumber': 1, // Number | Page Number
  'pageSize': 10, // Number | The maximum number of devices to return in the response. If this value is zero, the service will select a default size. 
  'fieldMask': "", // String | The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example: 
  'deviceIds': [], // [String] | A list of device string IDs. For example, ['device0', 'device12']. If empty, this field is ignored. Maximum IDs: 10,000
  'deviceNumIds':[], // [String] | A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000.
  'gatewayListOptionsAssociationsDeviceId': "", // String | If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound.
  'gatewayListOptionsAssociationsGatewayId': "", // String | If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned
  'gatewayListOptionsGatewayType': "" // String | If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned.
};
apiInstance.getDevices(registryId, subscriptionId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: \n');
    console.log(JSON.stringify(data))
    console.log('Device Id: ',data.devices[0].id);
    console.log('Parent: ',data.devices[0].parent);
    console.log('Registry Id:',data.devices[0].registry);
    console.log('Subscription Id:',data.devices[0].subscription);
    console.log('blocked: ',data.devices[0].blocked);
    console.log('capresent: ',data.devices[0].capresent);
    console.log('clientOnline: ',data.devices[0].clientOnline);
    console.log('config:',data.devices[0].config);
    console.log('createdOn:',data.devices[0].createdOn);
    console.log('credentials: ',data.devices[0].credentials);
    console.log('deviceErrors: ',data.devices[0].deviceErrors);
    console.log('gateway: ',data.devices[0].gateway);
    console.log('gatewayConfig: ',data.devices[0].gatewayConfig);
    console.log('isGateway:',data.devices[0].isGateway);
    console.log('lastConfigAckTime:',data.devices[0].lastConfigAckTime);
    console.log('lastConfigSendTime: ',data.devices[0].lastConfigSendTime);
    console.log('lastErrorStatus: ',data.devices[0].lastErrorStatus)
    console.log('lastErrorTime: ',data.devices[0].lastErrorTime);
    console.log('lastEventTime: ',data.devices[0].lastEventTime);
    console.log('lastHeartbeatTime: ',data.devices[0].lastHeartbeatTime);
    console.log('lastStateTime: ',data.devices[0].lastStateTime);
    console.log('logLevel:',data.devices[0].logLevel);
    console.log('metadata:',data.devices[0].metadata);
    console.log('name: ',data.devices[0].name);
    console.log('numId: ',data.devices[0].numId);
    console.log('state: ',data.devices[0].state);
    console.log('subscriptions: ',data.devices[0].subscriptions);
    console.log('updatedOn:',data.devices[0].updatedOn);
  }
});
}

module.exports.sendCommandToDevice =(token,subId,regId,devId,hostUrl)=>{
let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
defaultClient.basePath = hostUrl
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = token;

let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = subId; // String | Subscription ID
let registryId = regId; // String | Registry ID
let deviceId = devId; // String | Device ID

let device = {
binaryData:"aGVsbG8gd29ybGQ=", //Base64 encoded String | binary data
subfolder:""
}; // DeviceCommand | application/json

apiInstance.sendCommandToDevice(subscriptionid, registryId, deviceId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
}

module.exports.updateConfigurationToDevice = (token,subId,regId,devId,hostUrl)=>{
let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
defaultClient.basePath = hostUrl
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = token;
  
let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = subId; // String | Subscription ID
let registryId = regId; // String | Registry ID
let deviceId = devId; // String | Device ID

let device = {
    binaryData: "aGVsbG8gd29ybGQ=",
    subfolder: "", // Optional
    versionToUpdate: "" //Optional
}; // DeviceConfiguration | application/json

apiInstance.updateConfigurationToDevice(subscriptionid, registryId, deviceId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ');
    console.log('Acknowledged: ',data.acknowledged);
    console.log('Binary Data: ',data.binaryData);
    console.log('Cloud Update Time: ',data.cloudUpdateTime);
    console.log('Version: ',data.version)
  }
});
}

module.exports.getState = (token,subId,regId,devId,hostUrl)=>{
let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
defaultClient.basePath = hostUrl
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = token;

let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = subId; // String | Subscription ID
let registryId = regId; // String | Registry ID
let deviceId = devId; // String | Device ID
let opts = {
  'numStates': 0// Number | The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available.
};
apiInstance.getStates(subscriptionid, registryId, deviceId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ');
      console.log("Binary Data: ",data.deviceStates ? data.deviceStates[0].binaryData:"")
      console.log("Update Time: ",data.deviceStates ? data.deviceStates[0].updateTime:"")
  }
});
}

module.exports.getConfig=(token,subId,regId,devId,hostUrl)=>{
let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
defaultClient.basePath = hostUrl
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = token

let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = subId; // String | Subscription ID
let registryId = regId; // String | Registry ID
let deviceId = devId; // String | Device ID
let numVersions = 0; // Number | Device ID
apiInstance.getConfig(subscriptionid, registryId, deviceId, numVersions, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ');
    console.log("Acknowledged: ",data.deviceConfigs[0].acknowledged)
    console.log("Cloud Update Time: ",data.deviceConfigs[0].cloudUpdateTime)
    console.log("Binary Data: ",data.deviceConfigs[0].binaryData)
    console.log("Version : ",data.deviceConfigs[0].version)
  }
});
}

module.exports.blockDeviceCommunication=(token,subId,regId,devId,hostUrl)=>{
let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
defaultClient.basePath = hostUrl
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = token

let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = subId; // String | Subscription ID
let registryId = regId; // String | Registry ID
let deviceId = devId; // String | Device ID
let device = {
  "isblocked": false // Boolean | isblocked
}; // BlockCommunicationBody | application/json
apiInstance.blockDeviceCommuncation(subscriptionid, registryId, deviceId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data.info);
  }
});
}

module.exports.deleteDevice = (token,subId,regId,devId,hostUrl)=>{
let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
defaultClient.basePath = hostUrl
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = token

let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
let subscriptionId = subId; // String | Subscription ID
let registryId = regId; // String | Registry ID
let deviceId = devId; // String | Device ID
apiInstance.deleteDevice(subscriptionId, registryId, deviceId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data.info);
  }
});
}

