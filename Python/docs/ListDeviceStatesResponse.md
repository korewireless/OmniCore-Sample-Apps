# ListDeviceStatesResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_states** | [**List[DeviceState]**](DeviceState.md) |  | [optional] 

## Example

```python
from OmniCore.models.list_device_states_response import ListDeviceStatesResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ListDeviceStatesResponse from a JSON string
list_device_states_response_instance = ListDeviceStatesResponse.from_json(json)
# print the JSON string representation of the object
print ListDeviceStatesResponse.to_json()

# convert the object into a dict
list_device_states_response_dict = list_device_states_response_instance.to_dict()
# create an instance of ListDeviceStatesResponse from a dict
list_device_states_response_form_dict = list_device_states_response.from_dict(list_device_states_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


