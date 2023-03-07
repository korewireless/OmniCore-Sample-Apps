# UpdateRegistryRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**credentials** | [**List[RegistryCredential]**](RegistryCredential.md) |  | [optional] 
**http_config** | [**HttpConfig**](HttpConfig.md) |  | [optional] 
**mqtt_config** | [**MqttConfig**](MqttConfig.md) |  | [optional] 
**log_level** | [**LogLevel**](LogLevel.md) |  | [optional] 
**event_notification_configs** | [**List[EventNotificationConfig]**](EventNotificationConfig.md) |  | [optional] 
**log_notification_config** | [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**state_notification_config** | [**NotificationConfig**](NotificationConfig.md) |  | [optional] 

## Example

```python
from OmniCore.models.update_registry_request import UpdateRegistryRequest

# TODO update the JSON string below
json = "{}"
# create an instance of UpdateRegistryRequest from a JSON string
update_registry_request_instance = UpdateRegistryRequest.from_json(json)
# print the JSON string representation of the object
print UpdateRegistryRequest.to_json()

# convert the object into a dict
update_registry_request_dict = update_registry_request_instance.to_dict()
# create an instance of UpdateRegistryRequest from a dict
update_registry_request_form_dict = update_registry_request.from_dict(update_registry_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


