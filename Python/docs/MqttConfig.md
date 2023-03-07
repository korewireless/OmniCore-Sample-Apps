# MqttConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**mqtt_enabled_state** | **str** | MqttEnabledState: If enabled, allows connections using the MQTT protocol. Otherwise, MQTT connections to this registry will fail.  Possible values:   \&quot;MQTT_STATE_UNSPECIFIED\&quot; - No MQTT state specified. If not specified, MQTT will be enabled by default.   \&quot;MQTT_ENABLED\&quot; - Enables a MQTT connection.   \&quot;MQTT_DISABLED\&quot; - Disables a MQTT connection. | [optional] 

## Example

```python
from OmniCore.models.mqtt_config import MqttConfig

# TODO update the JSON string below
json = "{}"
# create an instance of MqttConfig from a JSON string
mqtt_config_instance = MqttConfig.from_json(json)
# print the JSON string representation of the object
print MqttConfig.to_json()

# convert the object into a dict
mqtt_config_dict = mqtt_config_instance.to_dict()
# create an instance of MqttConfig from a dict
mqtt_config_form_dict = mqtt_config.from_dict(mqtt_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


