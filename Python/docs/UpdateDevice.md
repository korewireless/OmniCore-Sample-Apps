# UpdateDevice


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**blocked** | **bool** |  | [optional] 
**credentials** | [**List[DeviceCredential]**](DeviceCredential.md) |  | [optional] 
**log_level** | [**LogLevel**](LogLevel.md) |  | [optional] 
**metadata** | **Dict[str, str]** |  | [optional] 

## Example

```python
from OmniCore.models.update_device import UpdateDevice

# TODO update the JSON string below
json = "{}"
# create an instance of UpdateDevice from a JSON string
update_device_instance = UpdateDevice.from_json(json)
# print the JSON string representation of the object
print UpdateDevice.to_json()

# convert the object into a dict
update_device_dict = update_device_instance.to_dict()
# create an instance of UpdateDevice from a dict
update_device_form_dict = update_device.from_dict(update_device_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


