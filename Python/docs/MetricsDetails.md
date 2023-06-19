# MetricsDetails


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**no_of_messages_for30_minutes** | **List[object]** |  | [optional] 
**no_of_messages_for48_hours** | **List[object]** |  | [optional] 
**billable_bytes_received** | **int** |  | [optional] 
**billable_bytes_sent** | **int** |  | [optional] 
**billable_message_size** | **int** |  | [optional] 
**bytes_received** | **int** |  | [optional] 
**bytes_sent** | **int** |  | [optional] 
**message_size** | **int** |  | [optional] 
**no_of_ack_messages** | **int** |  | [optional] 
**no_of_command_messages** | **int** |  | [optional] 
**no_of_config_messages** | **int** |  | [optional] 
**no_of_device_connections_failed** | **int** |  | [optional] 
**no_of_devices** | **int** |  | [optional] 
**no_of_dis_connections** | **int** |  | [optional] 
**no_of_event_messages** | **int** |  | [optional] 
**no_of_gateway_connections_failed** | **int** |  | [optional] 
**no_of_gateways** | **int** |  | [optional] 
**no_of_loop_back_messages** | **int** |  | [optional] 
**no_of_messages** | **int** |  | [optional] 
**no_of_publish_errors** | **int** |  | [optional] 
**no_of_registries** | **int** |  | [optional] 
**no_of_state_messages** | **int** |  | [optional] 
**no_of_subscribe** | **int** |  | [optional] 
**no_of_successful_connections** | **int** |  | [optional] 
**no_of_un_subscribe** | **int** |  | [optional] 
**subscription_id** | **str** |  | [optional] 

## Example

```python
from OmniCore.models.metrics_details import MetricsDetails

# TODO update the JSON string below
json = "{}"
# create an instance of MetricsDetails from a JSON string
metrics_details_instance = MetricsDetails.from_json(json)
# print the JSON string representation of the object
print MetricsDetails.to_json()

# convert the object into a dict
metrics_details_dict = metrics_details_instance.to_dict()
# create an instance of MetricsDetails from a dict
metrics_details_form_dict = metrics_details.from_dict(metrics_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


