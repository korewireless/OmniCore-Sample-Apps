# OmniCoreModelAndStateManagementApi.MqttConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**mqttEnabledState** | **String** | MqttEnabledState: If enabled, allows connections using the MQTT protocol. Otherwise, MQTT connections to this registry will fail.  Possible values:   \&quot;MQTT_STATE_UNSPECIFIED\&quot; - No MQTT state specified. If not specified, MQTT will be enabled by default.   \&quot;MQTT_ENABLED\&quot; - Enables a MQTT connection.   \&quot;MQTT_DISABLED\&quot; - Disables a MQTT connection. | [optional] 



## Enum: MqttEnabledStateEnum


* `ENABLED` (value: `"MQTT_ENABLED"`)

* `DISABLED` (value: `"MQTT_DISABLED"`)

* `STATE_UNSPECIFIED` (value: `"MQTT_STATE_UNSPECIFIED"`)




