# UpdateDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to **bool** |  | [optional] 
**Credentials** | Pointer to [**[]DeviceCredential**](DeviceCredential.md) |  | [optional] 
**LogLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUpdateDevice

`func NewUpdateDevice() *UpdateDevice`

NewUpdateDevice instantiates a new UpdateDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceWithDefaults

`func NewUpdateDeviceWithDefaults() *UpdateDevice`

NewUpdateDeviceWithDefaults instantiates a new UpdateDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *UpdateDevice) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *UpdateDevice) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *UpdateDevice) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *UpdateDevice) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetCredentials

`func (o *UpdateDevice) GetCredentials() []DeviceCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UpdateDevice) GetCredentialsOk() (*[]DeviceCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UpdateDevice) SetCredentials(v []DeviceCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UpdateDevice) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetLogLevel

`func (o *UpdateDevice) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *UpdateDevice) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *UpdateDevice) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *UpdateDevice) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateDevice) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateDevice) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateDevice) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateDevice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


