# DeviceConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**acknowledged** | **bool** |  | [optional] 
**binary_data** | **str** | Base64 Encoded Comnfig String | [optional] 
**cloud_update_time** | **str** |  | [optional] 
**device_ack_time** | **str** |  | [optional] 
**version** | **int** |  | [optional] 

## Example

```python
from OmniCore.models.device_config import DeviceConfig

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceConfig from a JSON string
device_config_instance = DeviceConfig.from_json(json)
# print the JSON string representation of the object
print DeviceConfig.to_json()

# convert the object into a dict
device_config_dict = device_config_instance.to_dict()
# create an instance of DeviceConfig from a dict
device_config_form_dict = device_config.from_dict(device_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


