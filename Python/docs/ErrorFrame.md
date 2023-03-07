# ErrorFrame


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **int** |  | 
**details** | **str** |  | 
**message** | **str** |  | 
**status** | **str** |  | 

## Example

```python
from OmniCore.models.error_frame import ErrorFrame

# TODO update the JSON string below
json = "{}"
# create an instance of ErrorFrame from a JSON string
error_frame_instance = ErrorFrame.from_json(json)
# print the JSON string representation of the object
print ErrorFrame.to_json()

# convert the object into a dict
error_frame_dict = error_frame_instance.to_dict()
# create an instance of ErrorFrame from a dict
error_frame_form_dict = error_frame.from_dict(error_frame_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


