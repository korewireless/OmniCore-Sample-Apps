# ListDeviceRegistriesResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_registries** | [**List[DeviceRegistry]**](DeviceRegistry.md) |  | 
**page_number** | **int** |  | [optional] 
**page_size** | **int** |  | [optional] 
**total_count** | **int** |  | [optional] 

## Example

```python
from OmniCore.models.list_device_registries_response import ListDeviceRegistriesResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ListDeviceRegistriesResponse from a JSON string
list_device_registries_response_instance = ListDeviceRegistriesResponse.from_json(json)
# print the JSON string representation of the object
print ListDeviceRegistriesResponse.to_json()

# convert the object into a dict
list_device_registries_response_dict = list_device_registries_response_instance.to_dict()
# create an instance of ListDeviceRegistriesResponse from a dict
list_device_registries_response_form_dict = list_device_registries_response.from_dict(list_device_registries_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


