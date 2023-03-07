# ListDevicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]Device**](Device.md) |  | 
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewListDevicesResponse

`func NewListDevicesResponse(devices []Device, ) *ListDevicesResponse`

NewListDevicesResponse instantiates a new ListDevicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDevicesResponseWithDefaults

`func NewListDevicesResponseWithDefaults() *ListDevicesResponse`

NewListDevicesResponseWithDefaults instantiates a new ListDevicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *ListDevicesResponse) GetDevices() []Device`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ListDevicesResponse) GetDevicesOk() (*[]Device, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ListDevicesResponse) SetDevices(v []Device)`

SetDevices sets Devices field to given value.


### GetPageNumber

`func (o *ListDevicesResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ListDevicesResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ListDevicesResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ListDevicesResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ListDevicesResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListDevicesResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListDevicesResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListDevicesResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListDevicesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListDevicesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListDevicesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListDevicesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


