# RegistryCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKeyCertificate** | Pointer to [**PublicKeyCertificate**](PublicKeyCertificate.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRegistryCredential

`func NewRegistryCredential() *RegistryCredential`

NewRegistryCredential instantiates a new RegistryCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryCredentialWithDefaults

`func NewRegistryCredentialWithDefaults() *RegistryCredential`

NewRegistryCredentialWithDefaults instantiates a new RegistryCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKeyCertificate

`func (o *RegistryCredential) GetPublicKeyCertificate() PublicKeyCertificate`

GetPublicKeyCertificate returns the PublicKeyCertificate field if non-nil, zero value otherwise.

### GetPublicKeyCertificateOk

`func (o *RegistryCredential) GetPublicKeyCertificateOk() (*PublicKeyCertificate, bool)`

GetPublicKeyCertificateOk returns a tuple with the PublicKeyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyCertificate

`func (o *RegistryCredential) SetPublicKeyCertificate(v PublicKeyCertificate)`

SetPublicKeyCertificate sets PublicKeyCertificate field to given value.

### HasPublicKeyCertificate

`func (o *RegistryCredential) HasPublicKeyCertificate() bool`

HasPublicKeyCertificate returns a boolean if a field has been set.

### GetId

`func (o *RegistryCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistryCredential) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


