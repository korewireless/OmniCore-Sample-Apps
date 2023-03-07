# BindRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_id** | **str** |  | 
**gateway_id** | **str** |  | 

## Example

```python
from OmniCore.models.bind_request import BindRequest

# TODO update the JSON string below
json = "{}"
# create an instance of BindRequest from a JSON string
bind_request_instance = BindRequest.from_json(json)
# print the JSON string representation of the object
print BindRequest.to_json()

# convert the object into a dict
bind_request_dict = bind_request_instance.to_dict()
# create an instance of BindRequest from a dict
bind_request_form_dict = bind_request.from_dict(bind_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


