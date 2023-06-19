# ListSinksSinksInner


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | [optional] [readonly] 
**subscription** | **str** |  | [optional] [readonly] 
**sink** | **str** |  | [optional] 
**config** | [**ListSinksSinksInnerConfig**](ListSinksSinksInnerConfig.md) |  | [optional] 
**status** | **bool** |  | [optional] [readonly] 
**createdon** | **str** |  | [optional] [readonly] 
**updatedon** | **str** |  | [optional] [readonly] 

## Example

```python
from OmniCore.models.list_sinks_sinks_inner import ListSinksSinksInner

# TODO update the JSON string below
json = "{}"
# create an instance of ListSinksSinksInner from a JSON string
list_sinks_sinks_inner_instance = ListSinksSinksInner.from_json(json)
# print the JSON string representation of the object
print ListSinksSinksInner.to_json()

# convert the object into a dict
list_sinks_sinks_inner_dict = list_sinks_sinks_inner_instance.to_dict()
# create an instance of ListSinksSinksInner from a dict
list_sinks_sinks_inner_form_dict = list_sinks_sinks_inner.from_dict(list_sinks_sinks_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


