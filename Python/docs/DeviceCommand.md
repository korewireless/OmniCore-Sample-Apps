# DeviceCommand


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**binary_data** | **str** | Base64 Encoded Command String | 
**subfolder** | **str** |  | [optional] 

## Example

```python
from OmniCore.models.device_command import DeviceCommand

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceCommand from a JSON string
device_command_instance = DeviceCommand.from_json(json)
# print the JSON string representation of the object
print DeviceCommand.to_json()

# convert the object into a dict
device_command_dict = device_command_instance.to_dict()
# create an instance of DeviceCommand from a dict
device_command_form_dict = device_command.from_dict(device_command_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


