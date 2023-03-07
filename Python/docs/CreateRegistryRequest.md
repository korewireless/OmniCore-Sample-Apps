# CreateRegistryRequest


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
from OmniCore.models.create_registry_request import CreateRegistryRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CreateRegistryRequest from a JSON string
create_registry_request_instance = CreateRegistryRequest.from_json(json)
# print the JSON string representation of the object
print CreateRegistryRequest.to_json()

# convert the object into a dict
create_registry_request_dict = create_registry_request_instance.to_dict()
# create an instance of CreateRegistryRequest from a dict
create_registry_request_form_dict = create_registry_request.from_dict(create_registry_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


