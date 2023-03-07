# DeviceCertificate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**credentials** | [**DeviceCredential**](DeviceCredential.md) |  | 

## Example

```python
from OmniCore.models.device_certificate import DeviceCertificate

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceCertificate from a JSON string
device_certificate_instance = DeviceCertificate.from_json(json)
# print the JSON string representation of the object
print DeviceCertificate.to_json()

# convert the object into a dict
device_certificate_dict = device_certificate_instance.to_dict()
# create an instance of DeviceCertificate from a dict
device_certificate_form_dict = device_certificate.from_dict(device_certificate_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


