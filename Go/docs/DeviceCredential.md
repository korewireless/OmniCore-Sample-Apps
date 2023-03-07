# DeviceCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTime** | Pointer to **string** | ExpirationTime: [Optional] The time at which this credential becomes invalid. This credential will be ignored for new client authentication requests after this timestamp; however, it will not be automatically deleted. | [optional] 
**PublicKey** | Pointer to [**PublicKeyCredential**](PublicKeyCredential.md) |  | [optional] 

## Methods

### NewDeviceCredential

`func NewDeviceCredential() *DeviceCredential`

NewDeviceCredential instantiates a new DeviceCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCredentialWithDefaults

`func NewDeviceCredentialWithDefaults() *DeviceCredential`

NewDeviceCredentialWithDefaults instantiates a new DeviceCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTime

`func (o *DeviceCredential) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *DeviceCredential) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *DeviceCredential) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *DeviceCredential) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetPublicKey

`func (o *DeviceCredential) GetPublicKey() PublicKeyCredential`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DeviceCredential) GetPublicKeyOk() (*PublicKeyCredential, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DeviceCredential) SetPublicKey(v PublicKeyCredential)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *DeviceCredential) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


