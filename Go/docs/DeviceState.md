# DeviceState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryData** | Pointer to **string** | Base64 Encoded State String | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceState

`func NewDeviceState() *DeviceState`

NewDeviceState instantiates a new DeviceState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStateWithDefaults

`func NewDeviceStateWithDefaults() *DeviceState`

NewDeviceStateWithDefaults instantiates a new DeviceState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryData

`func (o *DeviceState) GetBinaryData() string`

GetBinaryData returns the BinaryData field if non-nil, zero value otherwise.

### GetBinaryDataOk

`func (o *DeviceState) GetBinaryDataOk() (*string, bool)`

GetBinaryDataOk returns a tuple with the BinaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryData

`func (o *DeviceState) SetBinaryData(v string)`

SetBinaryData sets BinaryData field to given value.

### HasBinaryData

`func (o *DeviceState) HasBinaryData() bool`

HasBinaryData returns a boolean if a field has been set.

### GetUpdateTime

`func (o *DeviceState) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *DeviceState) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *DeviceState) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *DeviceState) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


