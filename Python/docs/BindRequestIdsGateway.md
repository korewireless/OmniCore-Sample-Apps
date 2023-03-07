# BindRequestIdsGateway


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_ids** | **List[str]** |  | 
**gateway_id** | **str** |  | 

## Example

```python
from OmniCore.models.bind_request_ids_gateway import BindRequestIdsGateway

# TODO update the JSON string below
json = "{}"
# create an instance of BindRequestIdsGateway from a JSON string
bind_request_ids_gateway_instance = BindRequestIdsGateway.from_json(json)
# print the JSON string representation of the object
print BindRequestIdsGateway.to_json()

# convert the object into a dict
bind_request_ids_gateway_dict = bind_request_ids_gateway_instance.to_dict()
# create an instance of BindRequestIdsGateway from a dict
bind_request_ids_gateway_form_dict = bind_request_ids_gateway.from_dict(bind_request_ids_gateway_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


