# \DeviceApi

All URIs are relative to *https://api.korewireless.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindDevice**](DeviceApi.md#BindDevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/bindDeviceToGateway | 
[**BindDevices**](DeviceApi.md#BindDevices) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/bindDevicesToGateway | 
[**BlockDeviceCommuncation**](DeviceApi.md#BlockDeviceCommuncation) | **Put** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/communication | 
[**CreateDevice**](DeviceApi.md#CreateDevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/devices | 
[**DeleteDevice**](DeviceApi.md#DeleteDevice) | **Delete** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
[**GetConfig**](DeviceApi.md#GetConfig) | **Get** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/configVersions | 
[**GetDevice**](DeviceApi.md#GetDevice) | **Get** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
[**GetDevices**](DeviceApi.md#GetDevices) | **Get** /subscriptions/{subscriptionId}/registries/{registryId}/devices | 
[**GetStates**](DeviceApi.md#GetStates) | **Get** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/states | 
[**SendCommandToDevice**](DeviceApi.md#SendCommandToDevice) | **Post** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendCommandToDevice | 
[**UnBindDevice**](DeviceApi.md#UnBindDevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDeviceFromGateway | 
[**UnBindDevices**](DeviceApi.md#UnBindDevices) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDevicesFromGateway | 
[**UpdateConfigurationToDevice**](DeviceApi.md#UpdateConfigurationToDevice) | **Post** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/updateConfigurationToDevice | 
[**UpdateCustomOnboardRequest**](DeviceApi.md#UpdateCustomOnboardRequest) | **Post** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/updateCustomOnboardRequest | 
[**UpdateDevice**](DeviceApi.md#UpdateDevice) | **Patch** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 



## BindDevice

> Info BindDevice(ctx, subscriptionId, registryId).Device(device).Execute()





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
    device := *openapiclient.NewBindRequest("DeviceId_example", "GatewayId_example") // BindRequest | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.BindDevice(context.Background(), subscriptionId, registryId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.BindDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BindDevice`: Info
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.BindDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**BindRequest**](BindRequest.md) | application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BindDevices

> Info BindDevices(ctx, subscriptionId, registryId).Device(device).Execute()





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
    device := *openapiclient.NewBindRequestIdsGateway([]string{"DeviceIds_example"}, "GatewayId_example") // BindRequestIdsGateway | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.BindDevices(context.Background(), subscriptionId, registryId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.BindDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BindDevices`: Info
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.BindDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**BindRequestIdsGateway**](BindRequestIdsGateway.md) | application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockDeviceCommuncation

> map[string]interface{} BlockDeviceCommuncation(ctx, subscriptionid, registryId, deviceId).Device(device).Execute()





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
    deviceId := "deviceId_example" // string | Device ID
    device := *openapiclient.NewBlockCommunicationBody() // BlockCommunicationBody | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.BlockDeviceCommuncation(context.Background(), subscriptionid, registryId, deviceId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.BlockDeviceCommuncation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BlockDeviceCommuncation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.BlockDeviceCommuncation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlockDeviceCommuncationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **device** | [**BlockCommunicationBody**](BlockCommunicationBody.md) | application/json | 

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


## CreateDevice

> Device CreateDevice(ctx, subscriptionId, registryId).Device(device).Execute()





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
    device := *openapiclient.NewDevice("Id_example") // Device | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.CreateDevice(context.Background(), subscriptionId, registryId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.CreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.CreateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**Device**](Device.md) | application/json | 

### Return type

[**Device**](Device.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> Info DeleteDevice(ctx, subscriptionId, registryId, deviceId).Execute()





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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.DeleteDevice(context.Background(), subscriptionId, registryId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDevice`: Info
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeleteDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


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


## GetConfig

> ListDeviceConfigVersionsResponse GetConfig(ctx, subscriptionid, registryId, deviceId).NumVersions(numVersions).Execute()





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
    deviceId := "deviceId_example" // string | Device ID
    numVersions := int32(56) // int32 | The number of versions to list. Versions are listed in decreasing order of the version number. The maximum number of versions retained is 10. If this value is zero, it will return all the versions available. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetConfig(context.Background(), subscriptionid, registryId, deviceId).NumVersions(numVersions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfig`: ListDeviceConfigVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **numVersions** | **int32** | The number of versions to list. Versions are listed in decreasing order of the version number. The maximum number of versions retained is 10. If this value is zero, it will return all the versions available. | 

### Return type

[**ListDeviceConfigVersionsResponse**](ListDeviceConfigVersionsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevice

> Device GetDevice(ctx, registryId, subscriptionId, deviceId).Execute()





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
    registryId := "registryId_example" // string | Registry ID
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDevice(context.Background(), registryId, subscriptionId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | Registry ID | 
**subscriptionId** | **string** | Subscription ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Device**](Device.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> ListDevicesResponse GetDevices(ctx, registryId, subscriptionId).PageNumber(pageNumber).PageSize(pageSize).FieldMask(fieldMask).DeviceIds(deviceIds).DeviceNumIds(deviceNumIds).GatewayListOptionsAssociationsDeviceId(gatewayListOptionsAssociationsDeviceId).GatewayListOptionsAssociationsGatewayId(gatewayListOptionsAssociationsGatewayId).GatewayListOptionsGatewayType(gatewayListOptionsGatewayType).Execute()





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
    registryId := "registryId_example" // string | Registry ID
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    pageNumber := int32(56) // int32 | Page Number (optional)
    pageSize := int32(56) // int32 | The maximum number of devices to return in the response. If this value is zero, the service will select a default size.  (optional)
    fieldMask := "fieldMask_example" // string | The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example:  (optional)
    deviceIds := []string{"Inner_example"} // []string | A list of device string IDs. For example, ['device0', 'device12']. If empty, this field is ignored. Maximum IDs: 10,000 (optional)
    deviceNumIds := []string{"Inner_example"} // []string | A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000. (optional)
    gatewayListOptionsAssociationsDeviceId := "gatewayListOptionsAssociationsDeviceId_example" // string | If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound. (optional)
    gatewayListOptionsAssociationsGatewayId := "gatewayListOptionsAssociationsGatewayId_example" // string | If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned (optional)
    gatewayListOptionsGatewayType := "gatewayListOptionsGatewayType_example" // string | If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDevices(context.Background(), registryId, subscriptionId).PageNumber(pageNumber).PageSize(pageSize).FieldMask(fieldMask).DeviceIds(deviceIds).DeviceNumIds(deviceNumIds).GatewayListOptionsAssociationsDeviceId(gatewayListOptionsAssociationsDeviceId).GatewayListOptionsAssociationsGatewayId(gatewayListOptionsAssociationsGatewayId).GatewayListOptionsGatewayType(gatewayListOptionsGatewayType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevices`: ListDevicesResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** | Registry ID | 
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **int32** | Page Number | 
 **pageSize** | **int32** | The maximum number of devices to return in the response. If this value is zero, the service will select a default size.  | 
 **fieldMask** | **string** | The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example:  | 
 **deviceIds** | **[]string** | A list of device string IDs. For example, [&#39;device0&#39;, &#39;device12&#39;]. If empty, this field is ignored. Maximum IDs: 10,000 | 
 **deviceNumIds** | **[]string** | A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000. | 
 **gatewayListOptionsAssociationsDeviceId** | **string** | If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound. | 
 **gatewayListOptionsAssociationsGatewayId** | **string** | If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned | 
 **gatewayListOptionsGatewayType** | **string** | If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned. | 

### Return type

[**ListDevicesResponse**](ListDevicesResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStates

> ListDeviceStatesResponse GetStates(ctx, subscriptionid, registryId, deviceId).NumStates(numStates).Execute()





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
    deviceId := "deviceId_example" // string | Device ID
    numStates := int32(56) // int32 | The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetStates(context.Background(), subscriptionid, registryId, deviceId).NumStates(numStates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStates`: ListDeviceStatesResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **numStates** | **int32** | The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available. | 

### Return type

[**ListDeviceStatesResponse**](ListDeviceStatesResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCommandToDevice

> map[string]interface{} SendCommandToDevice(ctx, subscriptionid, registryId, deviceId).Device(device).Execute()





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
    deviceId := "deviceId_example" // string | Device ID
    device := *openapiclient.NewDeviceCommand("BinaryData_example") // DeviceCommand | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.SendCommandToDevice(context.Background(), subscriptionid, registryId, deviceId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.SendCommandToDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendCommandToDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.SendCommandToDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendCommandToDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **device** | [**DeviceCommand**](DeviceCommand.md) | application/json | 

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


## UnBindDevice

> Info UnBindDevice(ctx, subscriptionId, registryId).Device(device).Execute()





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
    device := *openapiclient.NewBindRequest("DeviceId_example", "GatewayId_example") // BindRequest | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UnBindDevice(context.Background(), subscriptionId, registryId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UnBindDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnBindDevice`: Info
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UnBindDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnBindDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**BindRequest**](BindRequest.md) | application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnBindDevices

> Info UnBindDevices(ctx, subscriptionId, registryId).Device(device).Execute()





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
    device := *openapiclient.NewBindRequestIdsGateway([]string{"DeviceIds_example"}, "GatewayId_example") // BindRequestIdsGateway | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UnBindDevices(context.Background(), subscriptionId, registryId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UnBindDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnBindDevices`: Info
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UnBindDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnBindDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **device** | [**BindRequestIdsGateway**](BindRequestIdsGateway.md) | application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationToDevice

> DeviceConfig UpdateConfigurationToDevice(ctx, subscriptionid, registryId, deviceId).Device(device).Execute()





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
    deviceId := "deviceId_example" // string | Device ID
    device := *openapiclient.NewDeviceConfiguration("BinaryData_example") // DeviceConfiguration | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UpdateConfigurationToDevice(context.Background(), subscriptionid, registryId, deviceId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UpdateConfigurationToDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigurationToDevice`: DeviceConfig
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UpdateConfigurationToDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationToDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **device** | [**DeviceConfiguration**](DeviceConfiguration.md) | application/json | 

### Return type

[**DeviceConfig**](DeviceConfig.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomOnboardRequest

> Info UpdateCustomOnboardRequest(ctx, subscriptionid, registryId, deviceId).Device(device).Execute()





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
    deviceId := "deviceId_example" // string | Device ID
    device := *openapiclient.NewCustomOnboard("Id_example") // CustomOnboard | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UpdateCustomOnboardRequest(context.Background(), subscriptionid, registryId, deviceId).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UpdateCustomOnboardRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomOnboardRequest`: Info
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UpdateCustomOnboardRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomOnboardRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **device** | [**CustomOnboard**](CustomOnboard.md) | application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> Device UpdateDevice(ctx, subscriptionId, registryId, deviceId).UpdateMask(updateMask).Device(device).Execute()





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
    deviceId := "deviceId_example" // string | Device ID
    updateMask := "updateMask_example" // string | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,logLevel, blocked,policy and metadata
    device := *openapiclient.NewDevice("Id_example") // Device | application/json

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UpdateDevice(context.Background(), subscriptionId, registryId, deviceId).UpdateMask(updateMask).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UpdateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 
**registryId** | **string** | Registry ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateMask** | **string** | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,logLevel, blocked,policy and metadata | 
 **device** | [**Device**](Device.md) | application/json | 

### Return type

[**Device**](Device.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

