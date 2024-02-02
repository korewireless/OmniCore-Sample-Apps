# \RegistryApi

All URIs are relative to *https://api.korewireless.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegistry**](RegistryApi.md#CreateRegistry) | **Post** /subscriptions/{subscriptionId}/registries | 
[**DeleteRegistry**](RegistryApi.md#DeleteRegistry) | **Delete** /subscriptions/{subscriptionId}/registries/{registryId} | 
[**GetRegistries**](RegistryApi.md#GetRegistries) | **Get** /subscriptions/{subscriptionId}/registries | 
[**GetRegistry**](RegistryApi.md#GetRegistry) | **Get** /subscriptions/{subscriptionId}/registries/{registryId} | 
[**SendBroadcastToDevices**](RegistryApi.md#SendBroadcastToDevices) | **Post** /subscriptions/{subscriptionid}/registries/{registryId}/sendBroadcastToDevice | 
[**UpdateRegistry**](RegistryApi.md#UpdateRegistry) | **Patch** /subscriptions/{subscriptionId}/registries/{registryId} | 



## CreateRegistry

> DeviceRegistry CreateRegistry(ctx, subscriptionId).Registry(registry).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    registry := *openapiclient.NewDeviceRegistry("Id_example") // DeviceRegistry | application/json (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.CreateRegistry(context.Background(), subscriptionId).Registry(registry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.CreateRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegistry`: DeviceRegistry
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.CreateRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registry** | [**DeviceRegistry**](DeviceRegistry.md) | application/json | 

### Return type

[**DeviceRegistry**](DeviceRegistry.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegistry

> Info DeleteRegistry(ctx, subscriptionId, registryId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    registryId := "registryId_example" // string | Registry ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.DeleteRegistry(context.Background(), subscriptionId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.DeleteRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRegistry`: Info
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.DeleteRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistries

> ListDeviceRegistries GetRegistries(ctx, subscriptionId).PageNumber(pageNumber).PageSize(pageSize).RegistryIds(registryIds).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    pageNumber := int32(56) // int32 | Page Number (optional)
    pageSize := int32(56) // int32 | Page Size (optional)
    registryIds := []string{"Inner_example"} // []string | A list of registry string IDs. For example, ['registry0', 'registry12']. If empty, this field is ignored. Maximum IDs: 10,000 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.GetRegistries(context.Background(), subscriptionId).PageNumber(pageNumber).PageSize(pageSize).RegistryIds(registryIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetRegistries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistries`: ListDeviceRegistries
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetRegistries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **int32** | Page Number | 
 **pageSize** | **int32** | Page Size | 
 **registryIds** | **[]string** | A list of registry string IDs. For example, [&#39;registry0&#39;, &#39;registry12&#39;]. If empty, this field is ignored. Maximum IDs: 10,000 | 

### Return type

[**ListDeviceRegistries**](ListDeviceRegistries.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistry

> DeviceRegistry GetRegistry(ctx, subscriptionId, registryId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    registryId := "registryId_example" // string | Registry ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.GetRegistry(context.Background(), subscriptionId, registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistry`: DeviceRegistry
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceRegistry**](DeviceRegistry.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendBroadcastToDevices

> map[string]interface{} SendBroadcastToDevices(ctx, subscriptionid, registryId).Registry(registry).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionid := "subscriptionid_example" // string | Subscription ID
    registryId := "registryId_example" // string | Registry ID
    registry := *openapiclient.NewDeviceCommand("BinaryData_example") // DeviceCommand | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.SendBroadcastToDevices(context.Background(), subscriptionid, registryId).Registry(registry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.SendBroadcastToDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendBroadcastToDevices`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.SendBroadcastToDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendBroadcastToDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **registry** | [**DeviceCommand**](DeviceCommand.md) | application/json | 

### Return type

**map[string]interface{}**

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegistry

> DeviceRegistry UpdateRegistry(ctx, subscriptionId, registryId).UpdateMask(updateMask).Registry(registry).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/korewireless/OmniCore-Go-SDK"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    registryId := "registryId_example" // string | Registry ID
    updateMask := "updateMask_example" // string | values to be updated: eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,customOnboardNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials,customOnboardEnabled
    registry := *openapiclient.NewDeviceRegistry("Id_example") // DeviceRegistry | application/json (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.UpdateRegistry(context.Background(), subscriptionId, registryId).UpdateMask(updateMask).Registry(registry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.UpdateRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRegistry`: DeviceRegistry
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.UpdateRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMask** | **string** | values to be updated: eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,customOnboardNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials,customOnboardEnabled | 
 **registry** | [**DeviceRegistry**](DeviceRegistry.md) | application/json | 

### Return type

[**DeviceRegistry**](DeviceRegistry.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

