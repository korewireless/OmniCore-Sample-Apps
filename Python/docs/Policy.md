# Policy


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**connect** | **bool** |  | 
**publish_state** | **bool** |  | 
**publish_events** | **bool** |  | 
**publish_events_regex** | **str** |  | 
**publish_loopback** | **bool** |  | 
**subscribe_command** | **bool** |  | 
**subscribe_command_regex** | **str** |  | 
**subscribe_broadcast** | **bool** |  | 
**subscribe_broadcast_regex** | **str** |  | 
**subscribe_config** | **bool** |  | 

## Example

```python
from OmniCore.models.policy import Policy

# TODO update the JSON string below
json = "{}"
# create an instance of Policy from a JSON string
policy_instance = Policy.from_json(json)
# print the JSON string representation of the object
print Policy.to_json()

# convert the object into a dict
policy_dict = policy_instance.to_dict()
# create an instance of Policy from a dict
policy_form_dict = policy.from_dict(policy_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


