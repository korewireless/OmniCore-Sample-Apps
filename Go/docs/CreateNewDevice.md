# CreateNewDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Blocked** | Pointer to **bool** |  | [optional] 
**Credentials** | Pointer to [**[]DeviceCredential**](DeviceCredential.md) |  | [optional] 
**GatewayConfig** | Pointer to [**GatewayConfig**](GatewayConfig.md) |  | [optional] 
**Config** | Pointer to [**DeviceConfig**](DeviceConfig.md) |  | [optional] 
**LogLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateNewDevice

`func NewCreateNewDevice(id string, ) *CreateNewDevice`

NewCreateNewDevice instantiates a new CreateNewDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewDeviceWithDefaults

`func NewCreateNewDeviceWithDefaults() *CreateNewDevice`

NewCreateNewDeviceWithDefaults instantiates a new CreateNewDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNewDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNewDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNewDevice) SetId(v string)`

SetId sets Id field to given value.


### GetBlocked

`func (o *CreateNewDevice) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *CreateNewDevice) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *CreateNewDevice) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *CreateNewDevice) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetCredentials

`func (o *CreateNewDevice) GetCredentials() []DeviceCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateNewDevice) GetCredentialsOk() (*[]DeviceCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateNewDevice) SetCredentials(v []DeviceCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateNewDevice) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetGatewayConfig

`func (o *CreateNewDevice) GetGatewayConfig() GatewayConfig`

GetGatewayConfig returns the GatewayConfig field if non-nil, zero value otherwise.

### GetGatewayConfigOk

`func (o *CreateNewDevice) GetGatewayConfigOk() (*GatewayConfig, bool)`

GetGatewayConfigOk returns a tuple with the GatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConfig

`func (o *CreateNewDevice) SetGatewayConfig(v GatewayConfig)`

SetGatewayConfig sets GatewayConfig field to given value.

### HasGatewayConfig

`func (o *CreateNewDevice) HasGatewayConfig() bool`

HasGatewayConfig returns a boolean if a field has been set.

### GetConfig

`func (o *CreateNewDevice) GetConfig() DeviceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateNewDevice) GetConfigOk() (*DeviceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateNewDevice) SetConfig(v DeviceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateNewDevice) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLogLevel

`func (o *CreateNewDevice) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *CreateNewDevice) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *CreateNewDevice) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *CreateNewDevice) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateNewDevice) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateNewDevice) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateNewDevice) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateNewDevice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


