# PublicKeyCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | Certificate: The certificate data. | [optional] 
**Format** | Pointer to **string** | Format: The certificate format.  Possible values:   \&quot;UNSPECIFIED_PUBLIC_KEY_CERTIFICATE_FORMAT\&quot; - The format has not been specified. This is an invalid default value and must not be used.   \&quot;X509_CERTIFICATE_PEM\&quot; - An X.509v3 certificate ([RFC5280](https://www.ietf.org/rfc/rfc5280.txt)), encoded in base64, and wrapped by &#x60;-----BEGIN CERTIFICATE-----&#x60; and &#x60;-----END CERTIFICATE-----&#x60;. | [optional] 
**X509Details** | Pointer to [**X509CertificateDetails**](X509CertificateDetails.md) |  | [optional] 

## Methods

### NewPublicKeyCertificate

`func NewPublicKeyCertificate() *PublicKeyCertificate`

NewPublicKeyCertificate instantiates a new PublicKeyCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicKeyCertificateWithDefaults

`func NewPublicKeyCertificateWithDefaults() *PublicKeyCertificate`

NewPublicKeyCertificateWithDefaults instantiates a new PublicKeyCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *PublicKeyCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PublicKeyCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PublicKeyCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *PublicKeyCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetFormat

`func (o *PublicKeyCertificate) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PublicKeyCertificate) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PublicKeyCertificate) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PublicKeyCertificate) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetX509Details

`func (o *PublicKeyCertificate) GetX509Details() X509CertificateDetails`

GetX509Details returns the X509Details field if non-nil, zero value otherwise.

### GetX509DetailsOk

`func (o *PublicKeyCertificate) GetX509DetailsOk() (*X509CertificateDetails, bool)`

GetX509DetailsOk returns a tuple with the X509Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Details

`func (o *PublicKeyCertificate) SetX509Details(v X509CertificateDetails)`

SetX509Details sets X509Details field to given value.

### HasX509Details

`func (o *PublicKeyCertificate) HasX509Details() bool`

HasX509Details returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


