# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**NumId** | Pointer to **string** |  | [optional] 
**Parent** | **string** |  | 
**Registry** | **string** |  | 
**Blocked** | Pointer to **bool** |  | [optional] 
**Capresent** | Pointer to **bool** |  | [optional] 
**Subscription** | **string** |  | 
**CreatedOn** | Pointer to **string** |  | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**[]DeviceCredential**](DeviceCredential.md) |  | [optional] 
**Gateway** | Pointer to **[]string** |  | [optional] 
**GatewayConfig** | Pointer to [**GatewayConfig**](GatewayConfig.md) |  | [optional] 
**IsGateway** | Pointer to **bool** |  | [optional] 
**DeviceErrors** | Pointer to **string** |  | [optional] 
**ClientOnline** | Pointer to **bool** |  | [optional] 
**LastConfigAckTime** | Pointer to **string** |  | [optional] 
**LastConfigSendTime** | Pointer to **string** |  | [optional] 
**LastErrorStatus** | Pointer to [**ErrorStatus**](ErrorStatus.md) |  | [optional] 
**LastErrorTime** | Pointer to **string** |  | [optional] 
**LastEventTime** | Pointer to **string** |  | [optional] 
**LastHeartbeatTime** | Pointer to **string** |  | [optional] 
**LastStateTime** | Pointer to **string** |  | [optional] 
**LogLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Config** | Pointer to [**DeviceConfig**](DeviceConfig.md) |  | [optional] 
**State** | Pointer to [**DeviceState**](DeviceState.md) |  | [optional] 
**Subscriptions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDevice

`func NewDevice(id string, parent string, registry string, subscription string, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Device) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumId

`func (o *Device) GetNumId() string`

GetNumId returns the NumId field if non-nil, zero value otherwise.

### GetNumIdOk

`func (o *Device) GetNumIdOk() (*string, bool)`

GetNumIdOk returns a tuple with the NumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumId

`func (o *Device) SetNumId(v string)`

SetNumId sets NumId field to given value.

### HasNumId

`func (o *Device) HasNumId() bool`

HasNumId returns a boolean if a field has been set.

### GetParent

`func (o *Device) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Device) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Device) SetParent(v string)`

SetParent sets Parent field to given value.


### GetRegistry

`func (o *Device) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *Device) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *Device) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetBlocked

`func (o *Device) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *Device) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *Device) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *Device) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetCapresent

`func (o *Device) GetCapresent() bool`

GetCapresent returns the Capresent field if non-nil, zero value otherwise.

### GetCapresentOk

`func (o *Device) GetCapresentOk() (*bool, bool)`

GetCapresentOk returns a tuple with the Capresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapresent

`func (o *Device) SetCapresent(v bool)`

SetCapresent sets Capresent field to given value.

### HasCapresent

`func (o *Device) HasCapresent() bool`

HasCapresent returns a boolean if a field has been set.

### GetSubscription

`func (o *Device) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Device) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Device) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetCreatedOn

`func (o *Device) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Device) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Device) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Device) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Device) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Device) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Device) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Device) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetCredentials

`func (o *Device) GetCredentials() []DeviceCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Device) GetCredentialsOk() (*[]DeviceCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Device) SetCredentials(v []DeviceCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Device) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetGateway

`func (o *Device) GetGateway() []string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Device) GetGatewayOk() (*[]string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Device) SetGateway(v []string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Device) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayConfig

`func (o *Device) GetGatewayConfig() GatewayConfig`

GetGatewayConfig returns the GatewayConfig field if non-nil, zero value otherwise.

### GetGatewayConfigOk

`func (o *Device) GetGatewayConfigOk() (*GatewayConfig, bool)`

GetGatewayConfigOk returns a tuple with the GatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConfig

`func (o *Device) SetGatewayConfig(v GatewayConfig)`

SetGatewayConfig sets GatewayConfig field to given value.

### HasGatewayConfig

`func (o *Device) HasGatewayConfig() bool`

HasGatewayConfig returns a boolean if a field has been set.

### GetIsGateway

`func (o *Device) GetIsGateway() bool`

GetIsGateway returns the IsGateway field if non-nil, zero value otherwise.

### GetIsGatewayOk

`func (o *Device) GetIsGatewayOk() (*bool, bool)`

GetIsGatewayOk returns a tuple with the IsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGateway

`func (o *Device) SetIsGateway(v bool)`

SetIsGateway sets IsGateway field to given value.

### HasIsGateway

`func (o *Device) HasIsGateway() bool`

HasIsGateway returns a boolean if a field has been set.

### GetDeviceErrors

`func (o *Device) GetDeviceErrors() string`

GetDeviceErrors returns the DeviceErrors field if non-nil, zero value otherwise.

### GetDeviceErrorsOk

`func (o *Device) GetDeviceErrorsOk() (*string, bool)`

GetDeviceErrorsOk returns a tuple with the DeviceErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceErrors

`func (o *Device) SetDeviceErrors(v string)`

SetDeviceErrors sets DeviceErrors field to given value.

### HasDeviceErrors

`func (o *Device) HasDeviceErrors() bool`

HasDeviceErrors returns a boolean if a field has been set.

### GetClientOnline

`func (o *Device) GetClientOnline() bool`

GetClientOnline returns the ClientOnline field if non-nil, zero value otherwise.

### GetClientOnlineOk

`func (o *Device) GetClientOnlineOk() (*bool, bool)`

GetClientOnlineOk returns a tuple with the ClientOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOnline

`func (o *Device) SetClientOnline(v bool)`

SetClientOnline sets ClientOnline field to given value.

### HasClientOnline

`func (o *Device) HasClientOnline() bool`

HasClientOnline returns a boolean if a field has been set.

### GetLastConfigAckTime

`func (o *Device) GetLastConfigAckTime() string`

GetLastConfigAckTime returns the LastConfigAckTime field if non-nil, zero value otherwise.

### GetLastConfigAckTimeOk

`func (o *Device) GetLastConfigAckTimeOk() (*string, bool)`

GetLastConfigAckTimeOk returns a tuple with the LastConfigAckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfigAckTime

`func (o *Device) SetLastConfigAckTime(v string)`

SetLastConfigAckTime sets LastConfigAckTime field to given value.

### HasLastConfigAckTime

`func (o *Device) HasLastConfigAckTime() bool`

HasLastConfigAckTime returns a boolean if a field has been set.

### GetLastConfigSendTime

`func (o *Device) GetLastConfigSendTime() string`

GetLastConfigSendTime returns the LastConfigSendTime field if non-nil, zero value otherwise.

### GetLastConfigSendTimeOk

`func (o *Device) GetLastConfigSendTimeOk() (*string, bool)`

GetLastConfigSendTimeOk returns a tuple with the LastConfigSendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfigSendTime

`func (o *Device) SetLastConfigSendTime(v string)`

SetLastConfigSendTime sets LastConfigSendTime field to given value.

### HasLastConfigSendTime

`func (o *Device) HasLastConfigSendTime() bool`

HasLastConfigSendTime returns a boolean if a field has been set.

### GetLastErrorStatus

`func (o *Device) GetLastErrorStatus() ErrorStatus`

GetLastErrorStatus returns the LastErrorStatus field if non-nil, zero value otherwise.

### GetLastErrorStatusOk

`func (o *Device) GetLastErrorStatusOk() (*ErrorStatus, bool)`

GetLastErrorStatusOk returns a tuple with the LastErrorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorStatus

`func (o *Device) SetLastErrorStatus(v ErrorStatus)`

SetLastErrorStatus sets LastErrorStatus field to given value.

### HasLastErrorStatus

`func (o *Device) HasLastErrorStatus() bool`

HasLastErrorStatus returns a boolean if a field has been set.

### GetLastErrorTime

`func (o *Device) GetLastErrorTime() string`

GetLastErrorTime returns the LastErrorTime field if non-nil, zero value otherwise.

### GetLastErrorTimeOk

`func (o *Device) GetLastErrorTimeOk() (*string, bool)`

GetLastErrorTimeOk returns a tuple with the LastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorTime

`func (o *Device) SetLastErrorTime(v string)`

SetLastErrorTime sets LastErrorTime field to given value.

### HasLastErrorTime

`func (o *Device) HasLastErrorTime() bool`

HasLastErrorTime returns a boolean if a field has been set.

### GetLastEventTime

`func (o *Device) GetLastEventTime() string`

GetLastEventTime returns the LastEventTime field if non-nil, zero value otherwise.

### GetLastEventTimeOk

`func (o *Device) GetLastEventTimeOk() (*string, bool)`

GetLastEventTimeOk returns a tuple with the LastEventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventTime

`func (o *Device) SetLastEventTime(v string)`

SetLastEventTime sets LastEventTime field to given value.

### HasLastEventTime

`func (o *Device) HasLastEventTime() bool`

HasLastEventTime returns a boolean if a field has been set.

### GetLastHeartbeatTime

`func (o *Device) GetLastHeartbeatTime() string`

GetLastHeartbeatTime returns the LastHeartbeatTime field if non-nil, zero value otherwise.

### GetLastHeartbeatTimeOk

`func (o *Device) GetLastHeartbeatTimeOk() (*string, bool)`

GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatTime

`func (o *Device) SetLastHeartbeatTime(v string)`

SetLastHeartbeatTime sets LastHeartbeatTime field to given value.

### HasLastHeartbeatTime

`func (o *Device) HasLastHeartbeatTime() bool`

HasLastHeartbeatTime returns a boolean if a field has been set.

### GetLastStateTime

`func (o *Device) GetLastStateTime() string`

GetLastStateTime returns the LastStateTime field if non-nil, zero value otherwise.

### GetLastStateTimeOk

`func (o *Device) GetLastStateTimeOk() (*string, bool)`

GetLastStateTimeOk returns a tuple with the LastStateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateTime

`func (o *Device) SetLastStateTime(v string)`

SetLastStateTime sets LastStateTime field to given value.

### HasLastStateTime

`func (o *Device) HasLastStateTime() bool`

HasLastStateTime returns a boolean if a field has been set.

### GetLogLevel

`func (o *Device) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *Device) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *Device) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *Device) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetMetadata

`func (o *Device) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Device) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Device) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Device) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConfig

`func (o *Device) GetConfig() DeviceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Device) GetConfigOk() (*DeviceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Device) SetConfig(v DeviceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Device) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetState

`func (o *Device) GetState() DeviceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Device) GetStateOk() (*DeviceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Device) SetState(v DeviceState)`

SetState sets State field to given value.

### HasState

`func (o *Device) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubscriptions

`func (o *Device) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *Device) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *Device) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *Device) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


