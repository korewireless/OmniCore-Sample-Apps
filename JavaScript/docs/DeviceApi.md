# OmniCoreModelAndStateManagementApi.DeviceApi

All URIs are relative to *https://api.korewireless.com/omnicore*

Method | HTTP request | Description
------------- | ------------- | -------------
[**bindDevice**](DeviceApi.md#bindDevice) | **POST** /subscriptions/{subscriptionId}/registries/{registryId}/bindDeviceToGateway | 
[**bindDevices**](DeviceApi.md#bindDevices) | **POST** /subscriptions/{subscriptionId}/registries/{registryId}/bindDevicesToGateway | 
[**blockDeviceCommuncation**](DeviceApi.md#blockDeviceCommuncation) | **PUT** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/communication | 
[**createDevice**](DeviceApi.md#createDevice) | **POST** /subscriptions/{subscriptionId}/registries/{registryId}/devices | 
[**deleteDevice**](DeviceApi.md#deleteDevice) | **DELETE** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
[**getConfig**](DeviceApi.md#getConfig) | **GET** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/configVersions | 
[**getDevice**](DeviceApi.md#getDevice) | **GET** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
[**getDevices**](DeviceApi.md#getDevices) | **GET** /subscriptions/{subscriptionId}/registries/{registryId}/devices | 
[**getStates**](DeviceApi.md#getStates) | **GET** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/states | 
[**sendCommandToDevice**](DeviceApi.md#sendCommandToDevice) | **POST** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendCommandToDevice | 
[**unBindDevice**](DeviceApi.md#unBindDevice) | **POST** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDeviceFromGateway | 
[**unBindDevices**](DeviceApi.md#unBindDevices) | **POST** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDevicesFromGateway | 
[**updateConfigurationToDevice**](DeviceApi.md#updateConfigurationToDevice) | **POST** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/updateConfigurationToDevice | 
[**updateCustomOnboardRequest**](DeviceApi.md#updateCustomOnboardRequest) | **POST** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/updateCustomOnboardRequest | 
[**updateDevice**](DeviceApi.md#updateDevice) | **PATCH** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 



## bindDevice

> Info bindDevice(subscriptionId, registryId, device)



Bind  a device to a gateway under a registry

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let device = new OmniCoreModelAndStateManagementApi.BindRequest(); // BindRequest | application/json
apiInstance.bindDevice(subscriptionId, registryId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **device** | [**BindRequest**](BindRequest.md)| application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## bindDevices

> Info bindDevices(subscriptionId, registryId, device)



Bind devices to a gateway under a registry

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let device = new OmniCoreModelAndStateManagementApi.BindRequestIdsGateway(); // BindRequestIdsGateway | application/json
apiInstance.bindDevices(subscriptionId, registryId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **device** | [**BindRequestIdsGateway**](BindRequestIdsGateway.md)| application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## blockDeviceCommuncation

> Object blockDeviceCommuncation(subscriptionid, registryId, deviceId, device)



Blocks All Communication From A Device

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = "subscriptionid_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let deviceId = "deviceId_example"; // String | Device ID
let device = new OmniCoreModelAndStateManagementApi.BlockCommunicationBody(); // BlockCommunicationBody | application/json
apiInstance.blockDeviceCommuncation(subscriptionid, registryId, deviceId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **deviceId** | **String**| Device ID | 
 **device** | [**BlockCommunicationBody**](BlockCommunicationBody.md)| application/json | 

### Return type

**Object**

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## createDevice

> Device createDevice(subscriptionId, registryId, device)



Create a device under a registry

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let device = new OmniCoreModelAndStateManagementApi.Device(); // Device | application/json
apiInstance.createDevice(subscriptionId, registryId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **device** | [**Device**](Device.md)| application/json | 

### Return type

[**Device**](Device.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteDevice

> Info deleteDevice(subscriptionId, registryId, deviceId)



Delete a device under a registry

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let deviceId = "deviceId_example"; // String | Device ID
apiInstance.deleteDevice(subscriptionId, registryId, deviceId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **deviceId** | **String**| Device ID | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getConfig

> ListDeviceConfigVersionsResponse getConfig(subscriptionid, registryId, deviceId, opts)



Get Configs Of Devices

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = "subscriptionid_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let deviceId = "deviceId_example"; // String | Device ID
let opts = {
  'numVersions': 56 // Number | The number of versions to list. Versions are listed in decreasing order of the version number. The maximum number of versions retained is 10. If this value is zero, it will return all the versions available.
};
apiInstance.getConfig(subscriptionid, registryId, deviceId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **deviceId** | **String**| Device ID | 
 **numVersions** | **Number**| The number of versions to list. Versions are listed in decreasing order of the version number. The maximum number of versions retained is 10. If this value is zero, it will return all the versions available. | [optional] 

### Return type

[**ListDeviceConfigVersionsResponse**](ListDeviceConfigVersionsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getDevice

> Device getDevice(registryId, subscriptionId, deviceId)



Get a device under a registry

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let registryId = "registryId_example"; // String | Registry ID
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let deviceId = "deviceId_example"; // String | Device ID
apiInstance.getDevice(registryId, subscriptionId, deviceId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryId** | **String**| Registry ID | 
 **subscriptionId** | **String**| Subscription ID | 
 **deviceId** | **String**| Device ID | 

### Return type

[**Device**](Device.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getDevices

> ListDevicesResponse getDevices(registryId, subscriptionId, opts)



Get all devices under a registry

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let registryId = "registryId_example"; // String | Registry ID
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let opts = {
  'pageNumber': 56, // Number | Page Number
  'pageSize': 56, // Number | The maximum number of devices to return in the response. If this value is zero, the service will select a default size. 
  'fieldMask': "fieldMask_example", // String | The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example: 
  'deviceIds': ["null"], // [String] | A list of device string IDs. For example, ['device0', 'device12']. If empty, this field is ignored. Maximum IDs: 10,000
  'deviceNumIds': ["null"], // [String] | A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000.
  'gatewayListOptionsAssociationsDeviceId': "gatewayListOptionsAssociationsDeviceId_example", // String | If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound.
  'gatewayListOptionsAssociationsGatewayId': "gatewayListOptionsAssociationsGatewayId_example", // String | If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned
  'gatewayListOptionsGatewayType': "gatewayListOptionsGatewayType_example" // String | If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned.
};
apiInstance.getDevices(registryId, subscriptionId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryId** | **String**| Registry ID | 
 **subscriptionId** | **String**| Subscription ID | 
 **pageNumber** | **Number**| Page Number | [optional] 
 **pageSize** | **Number**| The maximum number of devices to return in the response. If this value is zero, the service will select a default size.  | [optional] 
 **fieldMask** | **String**| The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example:  | [optional] 
 **deviceIds** | [**[String]**](String.md)| A list of device string IDs. For example, [&#39;device0&#39;, &#39;device12&#39;]. If empty, this field is ignored. Maximum IDs: 10,000 | [optional] 
 **deviceNumIds** | [**[String]**](String.md)| A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000. | [optional] 
 **gatewayListOptionsAssociationsDeviceId** | **String**| If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound. | [optional] 
 **gatewayListOptionsAssociationsGatewayId** | **String**| If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned | [optional] 
 **gatewayListOptionsGatewayType** | **String**| If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned. | [optional] 

### Return type

[**ListDevicesResponse**](ListDevicesResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getStates

> ListDeviceStatesResponse getStates(subscriptionid, registryId, deviceId, opts)



Get States Of Devices

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = "subscriptionid_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let deviceId = "deviceId_example"; // String | Device ID
let opts = {
  'numStates': 56 // Number | The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available.
};
apiInstance.getStates(subscriptionid, registryId, deviceId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **deviceId** | **String**| Device ID | 
 **numStates** | **Number**| The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available. | [optional] 

### Return type

[**ListDeviceStatesResponse**](ListDeviceStatesResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## sendCommandToDevice

> Object sendCommandToDevice(subscriptionid, registryId, deviceId, device)



Send A Command To A Device

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = "subscriptionid_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let deviceId = "deviceId_example"; // String | Device ID
let device = new OmniCoreModelAndStateManagementApi.DeviceCommand(); // DeviceCommand | application/json
apiInstance.sendCommandToDevice(subscriptionid, registryId, deviceId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **deviceId** | **String**| Device ID | 
 **device** | [**DeviceCommand**](DeviceCommand.md)| application/json | 

### Return type

**Object**

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## unBindDevice

> Info unBindDevice(subscriptionId, registryId, device)



UnBind  a device from a gateway under a registry

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let device = new OmniCoreModelAndStateManagementApi.BindRequest(); // BindRequest | application/json
apiInstance.unBindDevice(subscriptionId, registryId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **device** | [**BindRequest**](BindRequest.md)| application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## unBindDevices

> Info unBindDevices(subscriptionId, registryId, device)



UnBind devices from a gateway under a registry

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let device = new OmniCoreModelAndStateManagementApi.BindRequestIdsGateway(); // BindRequestIdsGateway | application/json
apiInstance.unBindDevices(subscriptionId, registryId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **device** | [**BindRequestIdsGateway**](BindRequestIdsGateway.md)| application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateConfigurationToDevice

> DeviceConfig updateConfigurationToDevice(subscriptionid, registryId, deviceId, device)



Update A Configuration Of A Device

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = "subscriptionid_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let deviceId = "deviceId_example"; // String | Device ID
let device = new OmniCoreModelAndStateManagementApi.DeviceConfiguration(); // DeviceConfiguration | application/json
apiInstance.updateConfigurationToDevice(subscriptionid, registryId, deviceId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **deviceId** | **String**| Device ID | 
 **device** | [**DeviceConfiguration**](DeviceConfiguration.md)| application/json | 

### Return type

[**DeviceConfig**](DeviceConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateCustomOnboardRequest

> Info updateCustomOnboardRequest(subscriptionid, registryId, deviceId, device)



Approve/Reject a Custom Onboard Request

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionid = "subscriptionid_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let deviceId = "deviceId_example"; // String | Device ID
let device = new OmniCoreModelAndStateManagementApi.CustomOnboard(); // CustomOnboard | application/json
apiInstance.updateCustomOnboardRequest(subscriptionid, registryId, deviceId, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **deviceId** | **String**| Device ID | 
 **device** | [**CustomOnboard**](CustomOnboard.md)| application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateDevice

> Device updateDevice(subscriptionId, registryId, deviceId, updateMask, device)



Modify device under a registry

### Example

```javascript
import OmniCoreModelAndStateManagementApi from 'omni_core_model_and_state_management_api';
let defaultClient = OmniCoreModelAndStateManagementApi.ApiClient.instance;
// Configure API key authorization: apiKey
let apiKey = defaultClient.authentications['apiKey'];
apiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.apiKeyPrefix = 'Token';
// Configure Bearer (JWT) access token for authorization: bearerAuth
let bearerAuth = defaultClient.authentications['bearerAuth'];
bearerAuth.accessToken = "YOUR ACCESS TOKEN"

let apiInstance = new OmniCoreModelAndStateManagementApi.DeviceApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let deviceId = "deviceId_example"; // String | Device ID
let updateMask = "updateMask_example"; // String | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,logLevel, blocked,policy and metadata
let device = new OmniCoreModelAndStateManagementApi.Device(); // Device | application/json
apiInstance.updateDevice(subscriptionId, registryId, deviceId, updateMask, device, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **String**| Subscription ID | 
 **registryId** | **String**| Registry ID | 
 **deviceId** | **String**| Device ID | 
 **updateMask** | **String**| Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,logLevel, blocked,policy and metadata | 
 **device** | [**Device**](Device.md)| application/json | 

### Return type

[**Device**](Device.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

