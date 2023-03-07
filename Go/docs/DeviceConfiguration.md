# DeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryData** | **string** | Base64 Encoded Config String | 
**Subfolder** | Pointer to **string** |  | [optional] 
**VersionToUpdate** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceConfiguration

`func NewDeviceConfiguration(binaryData string, ) *DeviceConfiguration`

NewDeviceConfiguration instantiates a new DeviceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceConfigurationWithDefaults

`func NewDeviceConfigurationWithDefaults() *DeviceConfiguration`

NewDeviceConfigurationWithDefaults instantiates a new DeviceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryData

`func (o *DeviceConfiguration) GetBinaryData() string`

GetBinaryData returns the BinaryData field if non-nil, zero value otherwise.

### GetBinaryDataOk

`func (o *DeviceConfiguration) GetBinaryDataOk() (*string, bool)`

GetBinaryDataOk returns a tuple with the BinaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryData

`func (o *DeviceConfiguration) SetBinaryData(v string)`

SetBinaryData sets BinaryData field to given value.


### GetSubfolder

`func (o *DeviceConfiguration) GetSubfolder() string`

GetSubfolder returns the Subfolder field if non-nil, zero value otherwise.

### GetSubfolderOk

`func (o *DeviceConfiguration) GetSubfolderOk() (*string, bool)`

GetSubfolderOk returns a tuple with the Subfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubfolder

`func (o *DeviceConfiguration) SetSubfolder(v string)`

SetSubfolder sets Subfolder field to given value.

### HasSubfolder

`func (o *DeviceConfiguration) HasSubfolder() bool`

HasSubfolder returns a boolean if a field has been set.

### GetVersionToUpdate

`func (o *DeviceConfiguration) GetVersionToUpdate() string`

GetVersionToUpdate returns the VersionToUpdate field if non-nil, zero value otherwise.

### GetVersionToUpdateOk

`func (o *DeviceConfiguration) GetVersionToUpdateOk() (*string, bool)`

GetVersionToUpdateOk returns a tuple with the VersionToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionToUpdate

`func (o *DeviceConfiguration) SetVersionToUpdate(v string)`

SetVersionToUpdate sets VersionToUpdate field to given value.

### HasVersionToUpdate

`func (o *DeviceConfiguration) HasVersionToUpdate() bool`

HasVersionToUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


