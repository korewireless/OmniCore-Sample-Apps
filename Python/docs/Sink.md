# Sink


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | [optional] [readonly] 
**subscription** | **str** |  | [optional] [readonly] 
**sink** | **str** |  | [optional] 
**config** | [**Config**](Config.md) |  | [optional] 
**status** | **bool** |  | [optional] [readonly] 
**createdon** | **str** |  | [optional] [readonly] 
**updatedon** | **str** |  | [optional] [readonly] 

## Example

```python
from OmniCore.models.sink import Sink

# TODO update the JSON string below
json = "{}"
# create an instance of Sink from a JSON string
sink_instance = Sink.from_json(json)
# print the JSON string representation of the object
print Sink.to_json()

# convert the object into a dict
sink_dict = sink_instance.to_dict()
# create an instance of Sink from a dict
sink_form_dict = sink.from_dict(sink_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


