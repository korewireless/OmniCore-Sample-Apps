# UpdateRegistryRequest

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

### NewUpdateRegistryRequest

`func NewUpdateRegistryRequest() *UpdateRegistryRequest`

NewUpdateRegistryRequest instantiates a new UpdateRegistryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRegistryRequestWithDefaults

`func NewUpdateRegistryRequestWithDefaults() *UpdateRegistryRequest`

NewUpdateRegistryRequestWithDefaults instantiates a new UpdateRegistryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *UpdateRegistryRequest) GetCredentials() []RegistryCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UpdateRegistryRequest) GetCredentialsOk() (*[]RegistryCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UpdateRegistryRequest) SetCredentials(v []RegistryCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UpdateRegistryRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetHttpConfig

`func (o *UpdateRegistryRequest) GetHttpConfig() HttpConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *UpdateRegistryRequest) GetHttpConfigOk() (*HttpConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *UpdateRegistryRequest) SetHttpConfig(v HttpConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *UpdateRegistryRequest) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMqttConfig

`func (o *UpdateRegistryRequest) GetMqttConfig() MqttConfig`

GetMqttConfig returns the MqttConfig field if non-nil, zero value otherwise.

### GetMqttConfigOk

`func (o *UpdateRegistryRequest) GetMqttConfigOk() (*MqttConfig, bool)`

GetMqttConfigOk returns a tuple with the MqttConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConfig

`func (o *UpdateRegistryRequest) SetMqttConfig(v MqttConfig)`

SetMqttConfig sets MqttConfig field to given value.

### HasMqttConfig

`func (o *UpdateRegistryRequest) HasMqttConfig() bool`

HasMqttConfig returns a boolean if a field has been set.

### GetLogLevel

`func (o *UpdateRegistryRequest) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *UpdateRegistryRequest) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *UpdateRegistryRequest) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *UpdateRegistryRequest) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetEventNotificationConfigs

`func (o *UpdateRegistryRequest) GetEventNotificationConfigs() []EventNotificationConfig`

GetEventNotificationConfigs returns the EventNotificationConfigs field if non-nil, zero value otherwise.

### GetEventNotificationConfigsOk

`func (o *UpdateRegistryRequest) GetEventNotificationConfigsOk() (*[]EventNotificationConfig, bool)`

GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationConfigs

`func (o *UpdateRegistryRequest) SetEventNotificationConfigs(v []EventNotificationConfig)`

SetEventNotificationConfigs sets EventNotificationConfigs field to given value.

### HasEventNotificationConfigs

`func (o *UpdateRegistryRequest) HasEventNotificationConfigs() bool`

HasEventNotificationConfigs returns a boolean if a field has been set.

### GetLogNotificationConfig

`func (o *UpdateRegistryRequest) GetLogNotificationConfig() NotificationConfig`

GetLogNotificationConfig returns the LogNotificationConfig field if non-nil, zero value otherwise.

### GetLogNotificationConfigOk

`func (o *UpdateRegistryRequest) GetLogNotificationConfigOk() (*NotificationConfig, bool)`

GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotificationConfig

`func (o *UpdateRegistryRequest) SetLogNotificationConfig(v NotificationConfig)`

SetLogNotificationConfig sets LogNotificationConfig field to given value.

### HasLogNotificationConfig

`func (o *UpdateRegistryRequest) HasLogNotificationConfig() bool`

HasLogNotificationConfig returns a boolean if a field has been set.

### GetStateNotificationConfig

`func (o *UpdateRegistryRequest) GetStateNotificationConfig() NotificationConfig`

GetStateNotificationConfig returns the StateNotificationConfig field if non-nil, zero value otherwise.

### GetStateNotificationConfigOk

`func (o *UpdateRegistryRequest) GetStateNotificationConfigOk() (*NotificationConfig, bool)`

GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateNotificationConfig

`func (o *UpdateRegistryRequest) SetStateNotificationConfig(v NotificationConfig)`

SetStateNotificationConfig sets StateNotificationConfig field to given value.

### HasStateNotificationConfig

`func (o *UpdateRegistryRequest) HasStateNotificationConfig() bool`

HasStateNotificationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


