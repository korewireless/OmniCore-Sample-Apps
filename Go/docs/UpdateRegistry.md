# UpdateRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]RegistryCredential**](RegistryCredential.md) |  | [optional] 
**HttpConfig** | Pointer to [**HttpConfig**](HttpConfig.md) |  | [optional] 
**MqttConfig** | Pointer to [**MqttConfig**](MqttConfig.md) |  | [optional] 
**LogLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**EventNotificationConfigs** | Pointer to [**[]EventNotificationConfig**](EventNotificationConfig.md) |  | [optional] 
**LogNotificationConfig** | Pointer to [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**StateNotificationConfig** | Pointer to [**NotificationConfig**](NotificationConfig.md) |  | [optional] 

## Methods

### NewUpdateRegistry

`func NewUpdateRegistry() *UpdateRegistry`

NewUpdateRegistry instantiates a new UpdateRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRegistryWithDefaults

`func NewUpdateRegistryWithDefaults() *UpdateRegistry`

NewUpdateRegistryWithDefaults instantiates a new UpdateRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *UpdateRegistry) GetCredentials() []RegistryCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UpdateRegistry) GetCredentialsOk() (*[]RegistryCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UpdateRegistry) SetCredentials(v []RegistryCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UpdateRegistry) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetHttpConfig

`func (o *UpdateRegistry) GetHttpConfig() HttpConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *UpdateRegistry) GetHttpConfigOk() (*HttpConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *UpdateRegistry) SetHttpConfig(v HttpConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *UpdateRegistry) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMqttConfig

`func (o *UpdateRegistry) GetMqttConfig() MqttConfig`

GetMqttConfig returns the MqttConfig field if non-nil, zero value otherwise.

### GetMqttConfigOk

`func (o *UpdateRegistry) GetMqttConfigOk() (*MqttConfig, bool)`

GetMqttConfigOk returns a tuple with the MqttConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConfig

`func (o *UpdateRegistry) SetMqttConfig(v MqttConfig)`

SetMqttConfig sets MqttConfig field to given value.

### HasMqttConfig

`func (o *UpdateRegistry) HasMqttConfig() bool`

HasMqttConfig returns a boolean if a field has been set.

### GetLogLevel

`func (o *UpdateRegistry) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *UpdateRegistry) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *UpdateRegistry) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *UpdateRegistry) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetEventNotificationConfigs

`func (o *UpdateRegistry) GetEventNotificationConfigs() []EventNotificationConfig`

GetEventNotificationConfigs returns the EventNotificationConfigs field if non-nil, zero value otherwise.

### GetEventNotificationConfigsOk

`func (o *UpdateRegistry) GetEventNotificationConfigsOk() (*[]EventNotificationConfig, bool)`

GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationConfigs

`func (o *UpdateRegistry) SetEventNotificationConfigs(v []EventNotificationConfig)`

SetEventNotificationConfigs sets EventNotificationConfigs field to given value.

### HasEventNotificationConfigs

`func (o *UpdateRegistry) HasEventNotificationConfigs() bool`

HasEventNotificationConfigs returns a boolean if a field has been set.

### GetLogNotificationConfig

`func (o *UpdateRegistry) GetLogNotificationConfig() NotificationConfig`

GetLogNotificationConfig returns the LogNotificationConfig field if non-nil, zero value otherwise.

### GetLogNotificationConfigOk

`func (o *UpdateRegistry) GetLogNotificationConfigOk() (*NotificationConfig, bool)`

GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotificationConfig

`func (o *UpdateRegistry) SetLogNotificationConfig(v NotificationConfig)`

SetLogNotificationConfig sets LogNotificationConfig field to given value.

### HasLogNotificationConfig

`func (o *UpdateRegistry) HasLogNotificationConfig() bool`

HasLogNotificationConfig returns a boolean if a field has been set.

### GetStateNotificationConfig

`func (o *UpdateRegistry) GetStateNotificationConfig() NotificationConfig`

GetStateNotificationConfig returns the StateNotificationConfig field if non-nil, zero value otherwise.

### GetStateNotificationConfigOk

`func (o *UpdateRegistry) GetStateNotificationConfigOk() (*NotificationConfig, bool)`

GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateNotificationConfig

`func (o *UpdateRegistry) SetStateNotificationConfig(v NotificationConfig)`

SetStateNotificationConfig sets StateNotificationConfig field to given value.

### HasStateNotificationConfig

`func (o *UpdateRegistry) HasStateNotificationConfig() bool`

HasStateNotificationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


