# X509CertificateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryTime** | Pointer to **string** | ExpiryTime: The time the certificate becomes invalid. | [optional] 
**Issuer** | Pointer to **string** | Issuer: The entity that signed the certificate. | [optional] 
**PublicKeyType** | Pointer to **string** | PublicKeyType: The type of public key in the certificate. | [optional] 
**SignatureAlgorithm** | Pointer to **string** | SignatureAlgorithm: The algorithm used to sign the certificate. | [optional] 
**StartTime** | Pointer to **string** | StartTime: The time the certificate becomes valid. | [optional] 
**Subject** | Pointer to **string** | Subject: The entity the certificate and public key belong to. | [optional] 

## Methods

### NewX509CertificateDetails

`func NewX509CertificateDetails() *X509CertificateDetails`

NewX509CertificateDetails instantiates a new X509CertificateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509CertificateDetailsWithDefaults

`func NewX509CertificateDetailsWithDefaults() *X509CertificateDetails`

NewX509CertificateDetailsWithDefaults instantiates a new X509CertificateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryTime

`func (o *X509CertificateDetails) GetExpiryTime() string`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *X509CertificateDetails) GetExpiryTimeOk() (*string, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *X509CertificateDetails) SetExpiryTime(v string)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *X509CertificateDetails) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetIssuer

`func (o *X509CertificateDetails) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *X509CertificateDetails) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *X509CertificateDetails) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *X509CertificateDetails) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetPublicKeyType

`func (o *X509CertificateDetails) GetPublicKeyType() string`

GetPublicKeyType returns the PublicKeyType field if non-nil, zero value otherwise.

### GetPublicKeyTypeOk

`func (o *X509CertificateDetails) GetPublicKeyTypeOk() (*string, bool)`

GetPublicKeyTypeOk returns a tuple with the PublicKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyType

`func (o *X509CertificateDetails) SetPublicKeyType(v string)`

SetPublicKeyType sets PublicKeyType field to given value.

### HasPublicKeyType

`func (o *X509CertificateDetails) HasPublicKeyType() bool`

HasPublicKeyType returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *X509CertificateDetails) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *X509CertificateDetails) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *X509CertificateDetails) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *X509CertificateDetails) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetStartTime

`func (o *X509CertificateDetails) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *X509CertificateDetails) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *X509CertificateDetails) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *X509CertificateDetails) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSubject

`func (o *X509CertificateDetails) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *X509CertificateDetails) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *X509CertificateDetails) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *X509CertificateDetails) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


