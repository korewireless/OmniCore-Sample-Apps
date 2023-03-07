# Go API client for OmniCore

This is an OmniCore Model and State Management server.

## Overview

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [http://www.korewireless.com](http://www.korewireless.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import OmniCore "github.com/korewireless/OmniCore-Go-SDK"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), OmniCore.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), OmniCore.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), OmniCore.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), OmniCore.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://demo-api.omnicore.cloud.korewireless.com/model-state-management*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DeviceApi* | [**BindDevice**](docs/DeviceApi.md#binddevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/bindDeviceToGateway | 
*DeviceApi* | [**BindDevices**](docs/DeviceApi.md#binddevices) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/bindDevicesToGateway | 
*DeviceApi* | [**BlockDeviceCommuncation**](docs/DeviceApi.md#blockdevicecommuncation) | **Put** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/communication | 
*DeviceApi* | [**CreateDevice**](docs/DeviceApi.md#createdevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/devices | 
*DeviceApi* | [**DeleteDevice**](docs/DeviceApi.md#deletedevice) | **Delete** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
*DeviceApi* | [**GetConfig**](docs/DeviceApi.md#getconfig) | **Get** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/configVersions | 
*DeviceApi* | [**GetDevice**](docs/DeviceApi.md#getdevice) | **Get** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
*DeviceApi* | [**GetDevices**](docs/DeviceApi.md#getdevices) | **Get** /subscriptions/{subscriptionId}/registries/{registryId}/devices | 
*DeviceApi* | [**GetStates**](docs/DeviceApi.md#getstates) | **Get** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/states | 
*DeviceApi* | [**SendCommandToDevice**](docs/DeviceApi.md#sendcommandtodevice) | **Post** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendCommandToDevice | 
*DeviceApi* | [**UnBindDevice**](docs/DeviceApi.md#unbinddevice) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDeviceFromGateway | 
*DeviceApi* | [**UnBindDevices**](docs/DeviceApi.md#unbinddevices) | **Post** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDevicesFromGateway | 
*DeviceApi* | [**UpdateConfigurationToDevice**](docs/DeviceApi.md#updateconfigurationtodevice) | **Post** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/updateConfigurationToDevice | 
*DeviceApi* | [**UpdateDevice**](docs/DeviceApi.md#updatedevice) | **Patch** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
*RegistryApi* | [**CreateRegistry**](docs/RegistryApi.md#createregistry) | **Post** /subscriptions/{subscriptionId}/registries | 
*RegistryApi* | [**DeleteRegistry**](docs/RegistryApi.md#deleteregistry) | **Delete** /subscriptions/{subscriptionId}/registries/{registryId} | 
*RegistryApi* | [**GetRegistries**](docs/RegistryApi.md#getregistries) | **Get** /subscriptions/{subscriptionId}/registries | 
*RegistryApi* | [**GetRegistry**](docs/RegistryApi.md#getregistry) | **Get** /subscriptions/{subscriptionId}/registries/{registryId} | 
*RegistryApi* | [**UpdateRegistry**](docs/RegistryApi.md#updateregistry) | **Patch** /subscriptions/{subscriptionId}/registries/{registryId} | 


## Documentation For Models

 - [BindRequest](docs/BindRequest.md)
 - [BindRequestIdsGateway](docs/BindRequestIdsGateway.md)
 - [BlockCommunicationBody](docs/BlockCommunicationBody.md)
 - [CreateNewDevice](docs/CreateNewDevice.md)
 - [CreateRegistry200Response](docs/CreateRegistry200Response.md)
 - [CreateRegistryRequest](docs/CreateRegistryRequest.md)
 - [Device](docs/Device.md)
 - [DeviceCertificate](docs/DeviceCertificate.md)
 - [DeviceCommand](docs/DeviceCommand.md)
 - [DeviceConfig](docs/DeviceConfig.md)
 - [DeviceConfiguration](docs/DeviceConfiguration.md)
 - [DeviceCredential](docs/DeviceCredential.md)
 - [DeviceRegistry](docs/DeviceRegistry.md)
 - [DeviceState](docs/DeviceState.md)
 - [ErrorFrame](docs/ErrorFrame.md)
 - [ErrorStatus](docs/ErrorStatus.md)
 - [EventNotificationConfig](docs/EventNotificationConfig.md)
 - [GatewayConfig](docs/GatewayConfig.md)
 - [GenericErrorResponse](docs/GenericErrorResponse.md)
 - [HttpConfig](docs/HttpConfig.md)
 - [Info](docs/Info.md)
 - [ListDeviceConfigVersionsResponse](docs/ListDeviceConfigVersionsResponse.md)
 - [ListDeviceRegistries](docs/ListDeviceRegistries.md)
 - [ListDeviceStatesResponse](docs/ListDeviceStatesResponse.md)
 - [ListDevicesResponse](docs/ListDevicesResponse.md)
 - [LogLevel](docs/LogLevel.md)
 - [MqttConfig](docs/MqttConfig.md)
 - [NewRegistry](docs/NewRegistry.md)
 - [NotificationConfig](docs/NotificationConfig.md)
 - [PublicKeyCertificate](docs/PublicKeyCertificate.md)
 - [PublicKeyCredential](docs/PublicKeyCredential.md)
 - [RegistryCertificate](docs/RegistryCertificate.md)
 - [RegistryCredential](docs/RegistryCredential.md)
 - [UpdateDevice](docs/UpdateDevice.md)
 - [UpdateRegistry](docs/UpdateRegistry.md)
 - [UpdateRegistryRequest](docs/UpdateRegistryRequest.md)
 - [X509CertificateDetails](docs/X509CertificateDetails.md)


## Documentation For Authorization



### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@korewireless.com

