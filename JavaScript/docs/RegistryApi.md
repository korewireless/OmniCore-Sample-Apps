# OmniCoreModelAndStateManagementApi.RegistryApi

All URIs are relative to *https://api.korewireless.com/omnicore*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createRegistry**](RegistryApi.md#createRegistry) | **POST** /subscriptions/{subscriptionId}/registries | 
[**deleteRegistry**](RegistryApi.md#deleteRegistry) | **DELETE** /subscriptions/{subscriptionId}/registries/{registryId} | 
[**getRegistries**](RegistryApi.md#getRegistries) | **GET** /subscriptions/{subscriptionId}/registries | 
[**getRegistry**](RegistryApi.md#getRegistry) | **GET** /subscriptions/{subscriptionId}/registries/{registryId} | 
[**sendBroadcastToDevices**](RegistryApi.md#sendBroadcastToDevices) | **POST** /subscriptions/{subscriptionid}/registries/{registryId}/sendBroadcastToDevice | 
[**updateRegistry**](RegistryApi.md#updateRegistry) | **PATCH** /subscriptions/{subscriptionId}/registries/{registryId} | 



## createRegistry

> DeviceRegistry createRegistry(subscriptionId, opts)



Create a registry

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

let apiInstance = new OmniCoreModelAndStateManagementApi.RegistryApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let opts = {
  'registry': new OmniCoreModelAndStateManagementApi.DeviceRegistry() // DeviceRegistry | application/json
};
apiInstance.createRegistry(subscriptionId, opts, (error, data, response) => {
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
 **registry** | [**DeviceRegistry**](DeviceRegistry.md)| application/json | [optional] 

### Return type

[**DeviceRegistry**](DeviceRegistry.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## deleteRegistry

> Info deleteRegistry(subscriptionId, registryId)



Delete a registry

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

let apiInstance = new OmniCoreModelAndStateManagementApi.RegistryApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
apiInstance.deleteRegistry(subscriptionId, registryId, (error, data, response) => {
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

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getRegistries

> ListDeviceRegistries getRegistries(subscriptionId, opts)



Get all registries under a subscription

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

let apiInstance = new OmniCoreModelAndStateManagementApi.RegistryApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let opts = {
  'pageNumber': 56, // Number | Page Number
  'pageSize': 56, // Number | Page Size
  'registryIds': ["null"] // [String] | A list of registry string IDs. For example, ['registry0', 'registry12']. If empty, this field is ignored. Maximum IDs: 10,000
};
apiInstance.getRegistries(subscriptionId, opts, (error, data, response) => {
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
 **pageNumber** | **Number**| Page Number | [optional] 
 **pageSize** | **Number**| Page Size | [optional] 
 **registryIds** | [**[String]**](String.md)| A list of registry string IDs. For example, [&#39;registry0&#39;, &#39;registry12&#39;]. If empty, this field is ignored. Maximum IDs: 10,000 | [optional] 

### Return type

[**ListDeviceRegistries**](ListDeviceRegistries.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getRegistry

> DeviceRegistry getRegistry(subscriptionId, registryId)



Get a registry

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

let apiInstance = new OmniCoreModelAndStateManagementApi.RegistryApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
apiInstance.getRegistry(subscriptionId, registryId, (error, data, response) => {
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

### Return type

[**DeviceRegistry**](DeviceRegistry.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## sendBroadcastToDevices

> Object sendBroadcastToDevices(subscriptionid, registryId, registry)



Send  Broadcast To Devices

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

let apiInstance = new OmniCoreModelAndStateManagementApi.RegistryApi();
let subscriptionid = "subscriptionid_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let registry = new OmniCoreModelAndStateManagementApi.DeviceCommand(); // DeviceCommand | application/json
apiInstance.sendBroadcastToDevices(subscriptionid, registryId, registry, (error, data, response) => {
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
 **registry** | [**DeviceCommand**](DeviceCommand.md)| application/json | 

### Return type

**Object**

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## updateRegistry

> DeviceRegistry updateRegistry(subscriptionId, registryId, updateMask, opts)



Update a registry

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

let apiInstance = new OmniCoreModelAndStateManagementApi.RegistryApi();
let subscriptionId = "subscriptionId_example"; // String | Subscription ID
let registryId = "registryId_example"; // String | Registry ID
let updateMask = "updateMask_example"; // String | values to be updated: eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,customOnboardNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials,customOnboardEnabled
let opts = {
  'registry': new OmniCoreModelAndStateManagementApi.DeviceRegistry() // DeviceRegistry | application/json
};
apiInstance.updateRegistry(subscriptionId, registryId, updateMask, opts, (error, data, response) => {
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
 **updateMask** | **String**| values to be updated: eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,customOnboardNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials,customOnboardEnabled | 
 **registry** | [**DeviceRegistry**](DeviceRegistry.md)| application/json | [optional] 

### Return type

[**DeviceRegistry**](DeviceRegistry.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

