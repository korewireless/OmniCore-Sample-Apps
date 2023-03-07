# ListDeviceRegistries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceRegistries** | [**[]DeviceRegistry**](DeviceRegistry.md) |  | 
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewListDeviceRegistries

`func NewListDeviceRegistries(deviceRegistries []DeviceRegistry, ) *ListDeviceRegistries`

NewListDeviceRegistries instantiates a new ListDeviceRegistries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeviceRegistriesWithDefaults

`func NewListDeviceRegistriesWithDefaults() *ListDeviceRegistries`

NewListDeviceRegistriesWithDefaults instantiates a new ListDeviceRegistries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceRegistries

`func (o *ListDeviceRegistries) GetDeviceRegistries() []DeviceRegistry`

GetDeviceRegistries returns the DeviceRegistries field if non-nil, zero value otherwise.

### GetDeviceRegistriesOk

`func (o *ListDeviceRegistries) GetDeviceRegistriesOk() (*[]DeviceRegistry, bool)`

GetDeviceRegistriesOk returns a tuple with the DeviceRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistries

`func (o *ListDeviceRegistries) SetDeviceRegistries(v []DeviceRegistry)`

SetDeviceRegistries sets DeviceRegistries field to given value.


### GetPageNumber

`func (o *ListDeviceRegistries) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ListDeviceRegistries) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ListDeviceRegistries) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ListDeviceRegistries) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ListDeviceRegistries) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListDeviceRegistries) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListDeviceRegistries) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListDeviceRegistries) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListDeviceRegistries) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListDeviceRegistries) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListDeviceRegistries) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListDeviceRegistries) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


