# ListDeviceConfigVersionsResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_configs** | [**List[DeviceConfig]**](DeviceConfig.md) |  | [optional] 

## Example

```python
from OmniCore.models.list_device_config_versions_response import ListDeviceConfigVersionsResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ListDeviceConfigVersionsResponse from a JSON string
list_device_config_versions_response_instance = ListDeviceConfigVersionsResponse.from_json(json)
# print the JSON string representation of the object
print ListDeviceConfigVersionsResponse.to_json()

# convert the object into a dict
list_device_config_versions_response_dict = list_device_config_versions_response_instance.to_dict()
# create an instance of ListDeviceConfigVersionsResponse from a dict
list_device_config_versions_response_form_dict = list_device_config_versions_response.from_dict(list_device_config_versions_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


