# OmniCoreModelAndStateManagementApi.DeviceRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **String** |  | 
**name** | **String** |  | [optional] [readonly] 
**parent** | **String** |  | [optional] [readonly] 
**createdOn** | **String** |  | [optional] [readonly] 
**updatedOn** | **String** |  | [optional] [readonly] 
**credentials** | [**[RegistryCredential]**](RegistryCredential.md) |  | [optional] 
**httpConfig** | [**HttpConfig**](HttpConfig.md) |  | [optional] 
**mqttConfig** | [**MqttConfig**](MqttConfig.md) |  | [optional] 
**logLevel** | [**LogLevel**](LogLevel.md) |  | [optional] 
**eventNotificationConfigs** | [**[EventNotificationConfig]**](EventNotificationConfig.md) |  | [optional] 
**logNotificationConfig** | [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**stateNotificationConfig** | [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**customOnboardNotificationConfig** | [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**customOnboardEnabled** | **Boolean** |  | [optional] 
**numberOfDevices** | **Number** |  | [optional] [readonly] 
**numberOfGateways** | **Number** |  | [optional] [readonly] 


