# PublicKeyCredential


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**format** | **str** | Format: The format of the key.  Possible values:   \&quot;UNSPECIFIED_PUBLIC_KEY_FORMAT\&quot; - The format has not been specified. This is an invalid default value and must not be used.   \&quot;RSA_PEM\&quot; - An RSA public key encoded in base64, and wrapped by &#x60;-----BEGIN PUBLIC KEY-----&#x60; and &#x60;-----END PUBLIC KEY-----&#x60;. This can be used to verify &#x60;RS256&#x60; signatures in JWT tokens ([RFC7518]( https://www.ietf.org/rfc/rfc7518.txt)).   \&quot;RSA_X509_PEM\&quot; - As RSA_PEM, but wrapped in an X.509v3 certificate ([RFC5280]( https://www.ietf.org/rfc/rfc5280.txt)), encoded in base64, and wrapped by &#x60;-----BEGIN CERTIFICATE-----&#x60; and &#x60;-----END CERTIFICATE-----&#x60;.   \&quot;ES256_PEM\&quot; - Public key for the ECDSA algorithm using P-256 and SHA-256, encoded in base64, and wrapped by &#x60;-----BEGIN PUBLIC KEY-----&#x60; and &#x60;-----END PUBLIC KEY-----&#x60;. This can be used to verify JWT tokens with the &#x60;ES256&#x60; algorithm ([RFC7518](https://www.ietf.org/rfc/rfc7518.txt)). This curve is defined in [OpenSSL](https://www.openssl.org/) as the &#x60;prime256v1&#x60; curve.   \&quot;ES256_X509_PEM\&quot; - As ES256_PEM, but wrapped in an X.509v3 certificate ([RFC5280]( https://www.ietf.org/rfc/rfc5280.txt)), encoded in base64, and wrapped by &#x60;-----BEGIN CERTIFICATE-----&#x60; and &#x60;-----END CERTIFICATE-----&#x60;. | 
**key** | **str** | Key: The key data. | [optional] 

## Example

```python
from OmniCore.models.public_key_credential import PublicKeyCredential

# TODO update the JSON string below
json = "{}"
# create an instance of PublicKeyCredential from a JSON string
public_key_credential_instance = PublicKeyCredential.from_json(json)
# print the JSON string representation of the object
print PublicKeyCredential.to_json()

# convert the object into a dict
public_key_credential_dict = public_key_credential_instance.to_dict()
# create an instance of PublicKeyCredential from a dict
public_key_credential_form_dict = public_key_credential.from_dict(public_key_credential_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


