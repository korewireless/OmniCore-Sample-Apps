# NewRegistry


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**credentials** | [**List[RegistryCredential]**](RegistryCredential.md) |  | [optional] 
**http_config** | [**HttpConfig**](HttpConfig.md) |  | [optional] 
**mqtt_config** | [**MqttConfig**](MqttConfig.md) |  | [optional] 
**log_level** | [**LogLevel**](LogLevel.md) |  | [optional] 
**event_notification_configs** | [**List[EventNotificationConfig]**](EventNotificationConfig.md) |  | [optional] 
**log_notification_config** | [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**state_notification_config** | [**NotificationConfig**](NotificationConfig.md) |  | [optional] 

## Example

```python
from OmniCore.models.new_registry import NewRegistry

# TODO update the JSON string below
json = "{}"
# create an instance of NewRegistry from a JSON string
new_registry_instance = NewRegistry.from_json(json)
# print the JSON string representation of the object
print NewRegistry.to_json()

# convert the object into a dict
new_registry_dict = new_registry_instance.to_dict()
# create an instance of NewRegistry from a dict
new_registry_form_dict = new_registry.from_dict(new_registry_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


