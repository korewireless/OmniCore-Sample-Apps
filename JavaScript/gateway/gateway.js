var OmnicoreModelAndStateManagementApi = require('omni_core_model_and_state_management_api');

module.exports.bindDevice = (token,subId,regId,devId,gwId,hostUrl,key) => {
  let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
  defaultClient.basePath = hostUrl
  // Configure API key authorization: apiKey
  var apiKey = defaultClient.authentications['apiKey'];
  apiKey.apiKey = key
  // Configure Bearer (JWT) access token for authorization: bearerAuth
  let bearerAuth = defaultClient.authentications['bearerAuth'];
  bearerAuth.accessToken = token

  let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();

  let subscriptionId = subId; // String | Subscription ID
  let registryId = regId; // String | Registry ID
  let device = {
    deviceId: devId,
    gatewayId:gwId
  }; // BindRequest | application/json

  apiInstance.bindDevice(subscriptionId, registryId, device, (error, data, response) => {
    if (error) {
      console.error(error);
    } else {
      console.log('API called successfully. Returned data: ' + data.info);
    }
  });
}

module.exports.unbindDevice = (token,subId,regId,devId,gwId,hostUrl,key) => {
  let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
  defaultClient.basePath = hostUrl
  // Configure API key authorization: apiKey
  var apiKey = defaultClient.authentications['apiKey'];
  apiKey.apiKey = key
  // Configure Bearer (JWT) access token for authorization: bearerAuth
  let bearerAuth = defaultClient.authentications['bearerAuth'];
  bearerAuth.accessToken = token
  
  let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
  let subscriptionId = subId; // String | Subscription ID
  let registryId = regId; // String | Registry ID
  let device = {
          deviceId: devId,
          gatewayId:gwId
  }; // BindRequest | application/json
  apiInstance.unBindDevice(subscriptionId, registryId, device, (error, data, response) => {
    if (error) {
      console.error(error);
    } else {
      console.log('API called successfully. Returned data: ' + data.info);
    }
  });
}

module.exports.bindDevices = (token,subId,regId,devId,gwId,hostUrl,key) => {
  let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
  defaultClient.basePath = hostUrl
  // Configure API key authorization: apiKey
  var apiKey = defaultClient.authentications['apiKey'];
  apiKey.apiKey = key
  // Configure Bearer (JWT) access token for authorization: bearerAuth
  let bearerAuth = defaultClient.authentications['bearerAuth'];
  bearerAuth.accessToken = token
  
  let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
  let subscriptionId = subId; // String | Subscription ID
  let registryId = regId; // String | Registry ID
  let device = {
          deviceIds: [devId], // Array  | Device Ids
          gatewayId:gwId      // String | Gateway Ids
  }; // BindRequest | application/json
  apiInstance.bindDevices(subscriptionId, registryId, device, (error, data, response) => {
    if (error) {
      console.error(error);
    } else {
      console.log('API called successfully. Returned data: ' +data.info);
    }
  });
}

module.exports.unbindDevices = (token,subId,regId,devId,gwId,hostUrl,key) => {
  let defaultClient = OmnicoreModelAndStateManagementApi.ApiClient.instance;
  defaultClient.basePath = hostUrl
  // Configure API key authorization: apiKey
  var apiKey = defaultClient.authentications['apiKey'];
  apiKey.apiKey = key
  // Configure Bearer (JWT) access token for authorization: bearerAuth
  let bearerAuth = defaultClient.authentications['bearerAuth'];
  bearerAuth.accessToken = token
  
  let apiInstance = new OmnicoreModelAndStateManagementApi.DeviceApi();
  let subscriptionId = subId; // String | Subscription ID
  let registryId = regId; // String | Registry ID
  let device = {
    deviceIds: [devId], // Array  | Device Ids
    gatewayId:gwId      // String | Gateway Ids
  }; // BindRequest | application/json
  apiInstance.unBindDevices(subscriptionId, registryId, device, (error, data, response) => {
    if (error) {
      console.error(error);
    } else {
      console.log('API called successfully. Returned data: ' +data.info);
    }
  });
}