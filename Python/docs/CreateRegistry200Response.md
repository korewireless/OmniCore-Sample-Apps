# CreateRegistry200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**name** | **str** |  | [optional] 
**parent** | **str** |  | 
**created_on** | **str** |  | [optional] 
**updated_on** | **str** |  | [optional] 
**credentials** | [**List[RegistryCredential]**](RegistryCredential.md) |  | [optional] 
**http_config** | [**HttpConfig**](HttpConfig.md) |  | [optional] 
**mqtt_config** | [**MqttConfig**](MqttConfig.md) |  | [optional] 
**log_level** | [**LogLevel**](LogLevel.md) |  | [optional] 
**event_notification_configs** | [**List[EventNotificationConfig]**](EventNotificationConfig.md) |  | [optional] 
**log_notification_config** | [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**state_notification_config** | [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**number_of_devices** | **int** |  | [optional] 
**number_of_gateways** | **int** |  | [optional] 

## Example

```python
from OmniCore.models.create_registry200_response import CreateRegistry200Response

# TODO update the JSON string below
json = "{}"
# create an instance of CreateRegistry200Response from a JSON string
create_registry200_response_instance = CreateRegistry200Response.from_json(json)
# print the JSON string representation of the object
print CreateRegistry200Response.to_json()

# convert the object into a dict
create_registry200_response_dict = create_registry200_response_instance.to_dict()
# create an instance of CreateRegistry200Response from a dict
create_registry200_response_form_dict = create_registry200_response.from_dict(create_registry200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


