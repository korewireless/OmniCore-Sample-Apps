# NewRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Credentials** | Pointer to [**[]RegistryCredential**](RegistryCredential.md) |  | [optional] 
**HttpConfig** | Pointer to [**HttpConfig**](HttpConfig.md) |  | [optional] 
**MqttConfig** | Pointer to [**MqttConfig**](MqttConfig.md) |  | [optional] 
**LogLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**EventNotificationConfigs** | Pointer to [**[]EventNotificationConfig**](EventNotificationConfig.md) |  | [optional] 
**LogNotificationConfig** | Pointer to [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**StateNotificationConfig** | Pointer to [**NotificationConfig**](NotificationConfig.md) |  | [optional] 

## Methods

### NewNewRegistry

`func NewNewRegistry(id string, ) *NewRegistry`

NewNewRegistry instantiates a new NewRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewRegistryWithDefaults

`func NewNewRegistryWithDefaults() *NewRegistry`

NewNewRegistryWithDefaults instantiates a new NewRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewRegistry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewRegistry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewRegistry) SetId(v string)`

SetId sets Id field to given value.


### GetCredentials

`func (o *NewRegistry) GetCredentials() []RegistryCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *NewRegistry) GetCredentialsOk() (*[]RegistryCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *NewRegistry) SetCredentials(v []RegistryCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *NewRegistry) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetHttpConfig

`func (o *NewRegistry) GetHttpConfig() HttpConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *NewRegistry) GetHttpConfigOk() (*HttpConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *NewRegistry) SetHttpConfig(v HttpConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *NewRegistry) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMqttConfig

`func (o *NewRegistry) GetMqttConfig() MqttConfig`

GetMqttConfig returns the MqttConfig field if non-nil, zero value otherwise.

### GetMqttConfigOk

`func (o *NewRegistry) GetMqttConfigOk() (*MqttConfig, bool)`

GetMqttConfigOk returns a tuple with the MqttConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConfig

`func (o *NewRegistry) SetMqttConfig(v MqttConfig)`

SetMqttConfig sets MqttConfig field to given value.

### HasMqttConfig

`func (o *NewRegistry) HasMqttConfig() bool`

HasMqttConfig returns a boolean if a field has been set.

### GetLogLevel

`func (o *NewRegistry) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *NewRegistry) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *NewRegistry) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *NewRegistry) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetEventNotificationConfigs

`func (o *NewRegistry) GetEventNotificationConfigs() []EventNotificationConfig`

GetEventNotificationConfigs returns the EventNotificationConfigs field if non-nil, zero value otherwise.

### GetEventNotificationConfigsOk

`func (o *NewRegistry) GetEventNotificationConfigsOk() (*[]EventNotificationConfig, bool)`

GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationConfigs

`func (o *NewRegistry) SetEventNotificationConfigs(v []EventNotificationConfig)`

SetEventNotificationConfigs sets EventNotificationConfigs field to given value.

### HasEventNotificationConfigs

`func (o *NewRegistry) HasEventNotificationConfigs() bool`

HasEventNotificationConfigs returns a boolean if a field has been set.

### GetLogNotificationConfig

`func (o *NewRegistry) GetLogNotificationConfig() NotificationConfig`

GetLogNotificationConfig returns the LogNotificationConfig field if non-nil, zero value otherwise.

### GetLogNotificationConfigOk

`func (o *NewRegistry) GetLogNotificationConfigOk() (*NotificationConfig, bool)`

GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotificationConfig

`func (o *NewRegistry) SetLogNotificationConfig(v NotificationConfig)`

SetLogNotificationConfig sets LogNotificationConfig field to given value.

### HasLogNotificationConfig

`func (o *NewRegistry) HasLogNotificationConfig() bool`

HasLogNotificationConfig returns a boolean if a field has been set.

### GetStateNotificationConfig

`func (o *NewRegistry) GetStateNotificationConfig() NotificationConfig`

GetStateNotificationConfig returns the StateNotificationConfig field if non-nil, zero value otherwise.

### GetStateNotificationConfigOk

`func (o *NewRegistry) GetStateNotificationConfigOk() (*NotificationConfig, bool)`

GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateNotificationConfig

`func (o *NewRegistry) SetStateNotificationConfig(v NotificationConfig)`

SetStateNotificationConfig sets StateNotificationConfig field to given value.

### HasStateNotificationConfig

`func (o *NewRegistry) HasStateNotificationConfig() bool`

HasStateNotificationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


