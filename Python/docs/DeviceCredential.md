# DeviceCredential


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**expiration_time** | **str** | ExpirationTime: [Optional] The time at which this credential becomes invalid. This credential will be ignored for new client authentication requests after this timestamp; however, it will not be automatically deleted. | [optional] 
**id** | **str** |  | [optional] [readonly] 
**public_key** | [**PublicKeyCredential**](PublicKeyCredential.md) |  | [optional] 

## Example

```python
from OmniCore.models.device_credential import DeviceCredential

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceCredential from a JSON string
device_credential_instance = DeviceCredential.from_json(json)
# print the JSON string representation of the object
print DeviceCredential.to_json()

# convert the object into a dict
device_credential_dict = device_credential_instance.to_dict()
# create an instance of DeviceCredential from a dict
device_credential_form_dict = device_credential.from_dict(device_credential_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


