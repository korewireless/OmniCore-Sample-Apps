# BindRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** |  | 
**GatewayId** | **string** |  | 

## Methods

### NewBindRequest

`func NewBindRequest(deviceId string, gatewayId string, ) *BindRequest`

NewBindRequest instantiates a new BindRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindRequestWithDefaults

`func NewBindRequestWithDefaults() *BindRequest`

NewBindRequestWithDefaults instantiates a new BindRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *BindRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BindRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BindRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetGatewayId

`func (o *BindRequest) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *BindRequest) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *BindRequest) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


