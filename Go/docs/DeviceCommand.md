# DeviceCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryData** | **string** | Base64 Encoded Command String | 
**Subfolder** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceCommand

`func NewDeviceCommand(binaryData string, ) *DeviceCommand`

NewDeviceCommand instantiates a new DeviceCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCommandWithDefaults

`func NewDeviceCommandWithDefaults() *DeviceCommand`

NewDeviceCommandWithDefaults instantiates a new DeviceCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryData

`func (o *DeviceCommand) GetBinaryData() string`

GetBinaryData returns the BinaryData field if non-nil, zero value otherwise.

### GetBinaryDataOk

`func (o *DeviceCommand) GetBinaryDataOk() (*string, bool)`

GetBinaryDataOk returns a tuple with the BinaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryData

`func (o *DeviceCommand) SetBinaryData(v string)`

SetBinaryData sets BinaryData field to given value.


### GetSubfolder

`func (o *DeviceCommand) GetSubfolder() string`

GetSubfolder returns the Subfolder field if non-nil, zero value otherwise.

### GetSubfolderOk

`func (o *DeviceCommand) GetSubfolderOk() (*string, bool)`

GetSubfolderOk returns a tuple with the Subfolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubfolder

`func (o *DeviceCommand) SetSubfolder(v string)`

SetSubfolder sets Subfolder field to given value.

### HasSubfolder

`func (o *DeviceCommand) HasSubfolder() bool`

HasSubfolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


