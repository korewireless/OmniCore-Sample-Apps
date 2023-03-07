# RegistryCertificate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**credentials** | [**RegistryCredential**](RegistryCredential.md) |  | 

## Example

```python
from OmniCore.models.registry_certificate import RegistryCertificate

# TODO update the JSON string below
json = "{}"
# create an instance of RegistryCertificate from a JSON string
registry_certificate_instance = RegistryCertificate.from_json(json)
# print the JSON string representation of the object
print RegistryCertificate.to_json()

# convert the object into a dict
registry_certificate_dict = registry_certificate_instance.to_dict()
# create an instance of RegistryCertificate from a dict
registry_certificate_form_dict = registry_certificate.from_dict(registry_certificate_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


