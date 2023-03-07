# CreateRegistry200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Parent** | **string** |  | 
**CreatedOn** | Pointer to **string** |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**[]RegistryCredential**](RegistryCredential.md) |  | [optional] 
**HttpConfig** | Pointer to [**HttpConfig**](HttpConfig.md) |  | [optional] 
**MqttConfig** | Pointer to [**MqttConfig**](MqttConfig.md) |  | [optional] 
**LogLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**EventNotificationConfigs** | Pointer to [**[]EventNotificationConfig**](EventNotificationConfig.md) |  | [optional] 
**LogNotificationConfig** | Pointer to [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**StateNotificationConfig** | Pointer to [**NotificationConfig**](NotificationConfig.md) |  | [optional] 
**NumberOfDevices** | Pointer to **int32** |  | [optional] 
**NumberOfGateways** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateRegistry200Response

`func NewCreateRegistry200Response(id string, parent string, ) *CreateRegistry200Response`

NewCreateRegistry200Response instantiates a new CreateRegistry200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRegistry200ResponseWithDefaults

`func NewCreateRegistry200ResponseWithDefaults() *CreateRegistry200Response`

NewCreateRegistry200ResponseWithDefaults instantiates a new CreateRegistry200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateRegistry200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRegistry200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRegistry200Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateRegistry200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRegistry200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRegistry200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateRegistry200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *CreateRegistry200Response) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CreateRegistry200Response) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CreateRegistry200Response) SetParent(v string)`

SetParent sets Parent field to given value.


### GetCreatedOn

`func (o *CreateRegistry200Response) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *CreateRegistry200Response) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *CreateRegistry200Response) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *CreateRegistry200Response) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *CreateRegistry200Response) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *CreateRegistry200Response) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *CreateRegistry200Response) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *CreateRegistry200Response) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetCredentials

`func (o *CreateRegistry200Response) GetCredentials() []RegistryCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateRegistry200Response) GetCredentialsOk() (*[]RegistryCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateRegistry200Response) SetCredentials(v []RegistryCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateRegistry200Response) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetHttpConfig

`func (o *CreateRegistry200Response) GetHttpConfig() HttpConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *CreateRegistry200Response) GetHttpConfigOk() (*HttpConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *CreateRegistry200Response) SetHttpConfig(v HttpConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *CreateRegistry200Response) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMqttConfig

`func (o *CreateRegistry200Response) GetMqttConfig() MqttConfig`

GetMqttConfig returns the MqttConfig field if non-nil, zero value otherwise.

### GetMqttConfigOk

`func (o *CreateRegistry200Response) GetMqttConfigOk() (*MqttConfig, bool)`

GetMqttConfigOk returns a tuple with the MqttConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConfig

`func (o *CreateRegistry200Response) SetMqttConfig(v MqttConfig)`

SetMqttConfig sets MqttConfig field to given value.

### HasMqttConfig

`func (o *CreateRegistry200Response) HasMqttConfig() bool`

HasMqttConfig returns a boolean if a field has been set.

### GetLogLevel

`func (o *CreateRegistry200Response) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *CreateRegistry200Response) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *CreateRegistry200Response) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *CreateRegistry200Response) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetEventNotificationConfigs

`func (o *CreateRegistry200Response) GetEventNotificationConfigs() []EventNotificationConfig`

GetEventNotificationConfigs returns the EventNotificationConfigs field if non-nil, zero value otherwise.

### GetEventNotificationConfigsOk

`func (o *CreateRegistry200Response) GetEventNotificationConfigsOk() (*[]EventNotificationConfig, bool)`

GetEventNotificationConfigsOk returns a tuple with the EventNotificationConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationConfigs

`func (o *CreateRegistry200Response) SetEventNotificationConfigs(v []EventNotificationConfig)`

SetEventNotificationConfigs sets EventNotificationConfigs field to given value.

### HasEventNotificationConfigs

`func (o *CreateRegistry200Response) HasEventNotificationConfigs() bool`

HasEventNotificationConfigs returns a boolean if a field has been set.

### GetLogNotificationConfig

`func (o *CreateRegistry200Response) GetLogNotificationConfig() NotificationConfig`

GetLogNotificationConfig returns the LogNotificationConfig field if non-nil, zero value otherwise.

### GetLogNotificationConfigOk

`func (o *CreateRegistry200Response) GetLogNotificationConfigOk() (*NotificationConfig, bool)`

GetLogNotificationConfigOk returns a tuple with the LogNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogNotificationConfig

`func (o *CreateRegistry200Response) SetLogNotificationConfig(v NotificationConfig)`

SetLogNotificationConfig sets LogNotificationConfig field to given value.

### HasLogNotificationConfig

`func (o *CreateRegistry200Response) HasLogNotificationConfig() bool`

HasLogNotificationConfig returns a boolean if a field has been set.

### GetStateNotificationConfig

`func (o *CreateRegistry200Response) GetStateNotificationConfig() NotificationConfig`

GetStateNotificationConfig returns the StateNotificationConfig field if non-nil, zero value otherwise.

### GetStateNotificationConfigOk

`func (o *CreateRegistry200Response) GetStateNotificationConfigOk() (*NotificationConfig, bool)`

GetStateNotificationConfigOk returns a tuple with the StateNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateNotificationConfig

`func (o *CreateRegistry200Response) SetStateNotificationConfig(v NotificationConfig)`

SetStateNotificationConfig sets StateNotificationConfig field to given value.

### HasStateNotificationConfig

`func (o *CreateRegistry200Response) HasStateNotificationConfig() bool`

HasStateNotificationConfig returns a boolean if a field has been set.

### GetNumberOfDevices

`func (o *CreateRegistry200Response) GetNumberOfDevices() int32`

GetNumberOfDevices returns the NumberOfDevices field if non-nil, zero value otherwise.

### GetNumberOfDevicesOk

`func (o *CreateRegistry200Response) GetNumberOfDevicesOk() (*int32, bool)`

GetNumberOfDevicesOk returns a tuple with the NumberOfDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDevices

`func (o *CreateRegistry200Response) SetNumberOfDevices(v int32)`

SetNumberOfDevices sets NumberOfDevices field to given value.

### HasNumberOfDevices

`func (o *CreateRegistry200Response) HasNumberOfDevices() bool`

HasNumberOfDevices returns a boolean if a field has been set.

### GetNumberOfGateways

`func (o *CreateRegistry200Response) GetNumberOfGateways() int32`

GetNumberOfGateways returns the NumberOfGateways field if non-nil, zero value otherwise.

### GetNumberOfGatewaysOk

`func (o *CreateRegistry200Response) GetNumberOfGatewaysOk() (*int32, bool)`

GetNumberOfGatewaysOk returns a tuple with the NumberOfGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfGateways

`func (o *CreateRegistry200Response) SetNumberOfGateways(v int32)`

SetNumberOfGateways sets NumberOfGateways field to given value.

### HasNumberOfGateways

`func (o *CreateRegistry200Response) HasNumberOfGateways() bool`

HasNumberOfGateways returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


