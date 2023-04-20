# DeviceRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Parent** | Pointer to **string** |  | [optional] [readonly] 
**CreatedOn** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedOn** | Pointer to **string** |  | [optional] [readonly] 
**Credentials** | Pointer to [**[]RegistryCredential**](RegistryCredential.md) |  | [optional] 
**HttpConfig** | Pointer to [**HttpConfig**](HttpConfig.md) |  | [optional] 
**MqttConfig** | Pointer to [**MqttConfig**](MqttConfig.md) |  | [optional] 
**LogLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**EventNotificationConfigs** | Pointer to [**[]EventNotificationConfig**](EventNotificationConfig.md) |  | [optional] 
**LogNotificationConfig** | Pointer to [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**StateNotificationConfig** | Pointer to [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**CustomOnboardNotificationConfig** | Pointer to [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**CustomOnboardEnabled** | Pointer to **bool** |  | [optional] 
**NumberOfDevices** | Pointer to **int32** |  | [optional] [readonly] 
**NumberOfGateways** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewDeviceRegistry

`func NewDeviceRegistry(id string, ) *DeviceRegistry`

NewDeviceRegistry instantiates a new DeviceRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRegistryWithDefaults

`func NewDeviceRegistryWithDefaults() *DeviceRegistry`

NewDeviceRegistryWithDefaults instantiates a new DeviceRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceRegistry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceRegistry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceRegistry) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeviceRegistry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceRegistry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceRegistry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceRegistry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *DeviceRegistry) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *DeviceRegistry) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *DeviceRegistry) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *DeviceRegistry) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetCreatedOn

`func (o *DeviceRegistry) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *DeviceRegistry) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *DeviceRegistry) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *DeviceRegistry) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *DeviceRegistry) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *DeviceRegistry) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *DeviceRegistry) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *DeviceRegistry) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetCredentials

`func (o *DeviceRegistry) GetCredentials() []RegistryCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DeviceRegistry) GetCredentialsOk() (*[]RegistryCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DeviceRegistry) SetCredentials(v []RegistryCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *DeviceRegistry) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetHttpConfig

`func (o *DeviceRegistry) GetHttpConfig() HttpConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *DeviceRegistry) GetHttpConfigOk() (*HttpConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *DeviceRegistry) SetHttpConfig(v HttpConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *DeviceRegistry) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMqttConfig

`func (o *DeviceRegistry) GetMqttConfig() MqttConfig`

GetMqttConfig returns the MqttConfig field if non-nil, zero value otherwise.

### GetMqttConfigOk

`func (o *DeviceRegistry) GetMqttConfigOk() (*MqttConfig, bool)`

GetMqttConfigOk returns a tuple with the MqttConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConfig

`func (o *DeviceRegistry) SetMqttConfig(v MqttConfig)`

SetMqttConfig sets MqttConfig field to given value.

### HasMqttConfig

`func (o *DeviceRegistry) HasMqttConfig() bool`

HasMqttConfig returns a boolean if a field has been set.

### GetLogLevel

`func (o *DeviceRegistry) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *DeviceRegistry) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *DeviceRegistry) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *DeviceRegistry) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetEventNotificationConfigs

`func (o *DeviceRegistry) GetEventNotificationConfigs() []EventNotificationConfig`

GetEventNotificationConfigs returns the EventNotificationConfigs field if non-nil, zero value otherwise.

### GetEventNotificationConfigsOk

`func (o *DeviceRegistry) GetEventNotificationConfigsOk() (*[]EventNotificationConfig, bool)`

GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationConfigs

`func (o *DeviceRegistry) SetEventNotificationConfigs(v []EventNotificationConfig)`

SetEventNotificationConfigs sets EventNotificationConfigs field to given value.

### HasEventNotificationConfigs

`func (o *DeviceRegistry) HasEventNotificationConfigs() bool`

HasEventNotificationConfigs returns a boolean if a field has been set.

### GetLogNotificationConfig

`func (o *DeviceRegistry) GetLogNotificationConfig() NotificationConfig`

GetLogNotificationConfig returns the LogNotificationConfig field if non-nil, zero value otherwise.

### GetLogNotificationConfigOk

`func (o *DeviceRegistry) GetLogNotificationConfigOk() (*NotificationConfig, bool)`

GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotificationConfig

`func (o *DeviceRegistry) SetLogNotificationConfig(v NotificationConfig)`

SetLogNotificationConfig sets LogNotificationConfig field to given value.

### HasLogNotificationConfig

`func (o *DeviceRegistry) HasLogNotificationConfig() bool`

HasLogNotificationConfig returns a boolean if a field has been set.

### GetStateNotificationConfig

`func (o *DeviceRegistry) GetStateNotificationConfig() NotificationConfig`

GetStateNotificationConfig returns the StateNotificationConfig field if non-nil, zero value otherwise.

### GetStateNotificationConfigOk

`func (o *DeviceRegistry) GetStateNotificationConfigOk() (*NotificationConfig, bool)`

GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateNotificationConfig

`func (o *DeviceRegistry) SetStateNotificationConfig(v NotificationConfig)`

SetStateNotificationConfig sets StateNotificationConfig field to given value.

### HasStateNotificationConfig

`func (o *DeviceRegistry) HasStateNotificationConfig() bool`

HasStateNotificationConfig returns a boolean if a field has been set.

### GetCustomOnboardNotificationConfig

`func (o *DeviceRegistry) GetCustomOnboardNotificationConfig() NotificationConfig`

GetCustomOnboardNotificationConfig returns the CustomOnboardNotificationConfig field if non-nil, zero value otherwise.

### GetCustomOnboardNotificationConfigOk

`func (o *DeviceRegistry) GetCustomOnboardNotificationConfigOk() (*NotificationConfig, bool)`

GetCustomOnboardNotificationConfigOk returns a tuple with the CustomOnboardNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOnboardNotificationConfig

`func (o *DeviceRegistry) SetCustomOnboardNotificationConfig(v NotificationConfig)`

SetCustomOnboardNotificationConfig sets CustomOnboardNotificationConfig field to given value.

### HasCustomOnboardNotificationConfig

`func (o *DeviceRegistry) HasCustomOnboardNotificationConfig() bool`

HasCustomOnboardNotificationConfig returns a boolean if a field has been set.

### GetCustomOnboardEnabled

`func (o *DeviceRegistry) GetCustomOnboardEnabled() bool`

GetCustomOnboardEnabled returns the CustomOnboardEnabled field if non-nil, zero value otherwise.

### GetCustomOnboardEnabledOk

`func (o *DeviceRegistry) GetCustomOnboardEnabledOk() (*bool, bool)`

GetCustomOnboardEnabledOk returns a tuple with the CustomOnboardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOnboardEnabled

`func (o *DeviceRegistry) SetCustomOnboardEnabled(v bool)`

SetCustomOnboardEnabled sets CustomOnboardEnabled field to given value.

### HasCustomOnboardEnabled

`func (o *DeviceRegistry) HasCustomOnboardEnabled() bool`

HasCustomOnboardEnabled returns a boolean if a field has been set.

### GetNumberOfDevices

`func (o *DeviceRegistry) GetNumberOfDevices() int32`

GetNumberOfDevices returns the NumberOfDevices field if non-nil, zero value otherwise.

### GetNumberOfDevicesOk

`func (o *DeviceRegistry) GetNumberOfDevicesOk() (*int32, bool)`

GetNumberOfDevicesOk returns a tuple with the NumberOfDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDevices

`func (o *DeviceRegistry) SetNumberOfDevices(v int32)`

SetNumberOfDevices sets NumberOfDevices field to given value.

### HasNumberOfDevices

`func (o *DeviceRegistry) HasNumberOfDevices() bool`

HasNumberOfDevices returns a boolean if a field has been set.

### GetNumberOfGateways

`func (o *DeviceRegistry) GetNumberOfGateways() int32`

GetNumberOfGateways returns the NumberOfGateways field if non-nil, zero value otherwise.

### GetNumberOfGatewaysOk

`func (o *DeviceRegistry) GetNumberOfGatewaysOk() (*int32, bool)`

GetNumberOfGatewaysOk returns a tuple with the NumberOfGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfGateways

`func (o *DeviceRegistry) SetNumberOfGateways(v int32)`

SetNumberOfGateways sets NumberOfGateways field to given value.

### HasNumberOfGateways

`func (o *DeviceRegistry) HasNumberOfGateways() bool`

HasNumberOfGateways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


