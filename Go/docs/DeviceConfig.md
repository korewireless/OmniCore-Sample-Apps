# DeviceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledged** | Pointer to **bool** |  | [optional] 
**BinaryData** | Pointer to **string** | Base64 Encoded Comnfig String | [optional] 
**CloudUpdateTime** | Pointer to **string** |  | [optional] 
**DeviceAckTime** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewDeviceConfig

`func NewDeviceConfig() *DeviceConfig`

NewDeviceConfig instantiates a new DeviceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceConfigWithDefaults

`func NewDeviceConfigWithDefaults() *DeviceConfig`

NewDeviceConfigWithDefaults instantiates a new DeviceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledged

`func (o *DeviceConfig) GetAcknowledged() bool`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *DeviceConfig) GetAcknowledgedOk() (*bool, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *DeviceConfig) SetAcknowledged(v bool)`

SetAcknowledged sets Acknowledged field to given value.

### HasAcknowledged

`func (o *DeviceConfig) HasAcknowledged() bool`

HasAcknowledged returns a boolean if a field has been set.

### GetBinaryData

`func (o *DeviceConfig) GetBinaryData() string`

GetBinaryData returns the BinaryData field if non-nil, zero value otherwise.

### GetBinaryDataOk

`func (o *DeviceConfig) GetBinaryDataOk() (*string, bool)`

GetBinaryDataOk returns a tuple with the BinaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryData

`func (o *DeviceConfig) SetBinaryData(v string)`

SetBinaryData sets BinaryData field to given value.

### HasBinaryData

`func (o *DeviceConfig) HasBinaryData() bool`

HasBinaryData returns a boolean if a field has been set.

### GetCloudUpdateTime

`func (o *DeviceConfig) GetCloudUpdateTime() string`

GetCloudUpdateTime returns the CloudUpdateTime field if non-nil, zero value otherwise.

### GetCloudUpdateTimeOk

`func (o *DeviceConfig) GetCloudUpdateTimeOk() (*string, bool)`

GetCloudUpdateTimeOk returns a tuple with the CloudUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudUpdateTime

`func (o *DeviceConfig) SetCloudUpdateTime(v string)`

SetCloudUpdateTime sets CloudUpdateTime field to given value.

### HasCloudUpdateTime

`func (o *DeviceConfig) HasCloudUpdateTime() bool`

HasCloudUpdateTime returns a boolean if a field has been set.

### GetDeviceAckTime

`func (o *DeviceConfig) GetDeviceAckTime() string`

GetDeviceAckTime returns the DeviceAckTime field if non-nil, zero value otherwise.

### GetDeviceAckTimeOk

`func (o *DeviceConfig) GetDeviceAckTimeOk() (*string, bool)`

GetDeviceAckTimeOk returns a tuple with the DeviceAckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAckTime

`func (o *DeviceConfig) SetDeviceAckTime(v string)`

SetDeviceAckTime sets DeviceAckTime field to given value.

### HasDeviceAckTime

`func (o *DeviceConfig) HasDeviceAckTime() bool`

HasDeviceAckTime returns a boolean if a field has been set.

### GetVersion

`func (o *DeviceConfig) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeviceConfig) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeviceConfig) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeviceConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


