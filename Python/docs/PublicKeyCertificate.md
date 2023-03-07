# PublicKeyCertificate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**certificate** | **str** | Certificate: The certificate data. | [optional] 
**format** | **str** | Format: The certificate format.  Possible values:   \&quot;UNSPECIFIED_PUBLIC_KEY_CERTIFICATE_FORMAT\&quot; - The format has not been specified. This is an invalid default value and must not be used.   \&quot;X509_CERTIFICATE_PEM\&quot; - An X.509v3 certificate ([RFC5280](https://www.ietf.org/rfc/rfc5280.txt)), encoded in base64, and wrapped by &#x60;-----BEGIN CERTIFICATE-----&#x60; and &#x60;-----END CERTIFICATE-----&#x60;. | [optional] 
**x509_details** | [**X509CertificateDetails**](X509CertificateDetails.md) |  | [optional] 

## Example

```python
from OmniCore.models.public_key_certificate import PublicKeyCertificate

# TODO update the JSON string below
json = "{}"
# create an instance of PublicKeyCertificate from a JSON string
public_key_certificate_instance = PublicKeyCertificate.from_json(json)
# print the JSON string representation of the object
print PublicKeyCertificate.to_json()

# convert the object into a dict
public_key_certificate_dict = public_key_certificate_instance.to_dict()
# create an instance of PublicKeyCertificate from a dict
public_key_certificate_form_dict = public_key_certificate.from_dict(public_key_certificate_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


