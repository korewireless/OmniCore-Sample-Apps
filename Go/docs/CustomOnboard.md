# CustomOnboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**NumId** | Pointer to **string** |  | [optional] [readonly] 
**Parent** | Pointer to **string** |  | [optional] [readonly] 
**Registry** | Pointer to **string** |  | [optional] [readonly] 
**Blocked** | Pointer to **bool** |  | [optional] 
**Capresent** | Pointer to **bool** |  | [optional] [readonly] 
**Subscription** | Pointer to **string** |  | [optional] [readonly] 
**CreatedOn** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedOn** | Pointer to **string** |  | [optional] [readonly] 
**Credentials** | Pointer to [**[]DeviceCredential**](DeviceCredential.md) |  | [optional] 
**Gateway** | Pointer to **[]string** |  | [optional] 
**GatewayConfig** | Pointer to [**GatewayConfig**](GatewayConfig.md) |  | [optional] 
**IsGateway** | Pointer to **bool** |  | [optional] 
**DeviceErrors** | Pointer to **string** |  | [optional] [readonly] 
**ClientOnline** | Pointer to **bool** |  | [optional] [readonly] 
**LastConfigAckTime** | Pointer to **string** |  | [optional] [readonly] 
**LastConfigSendTime** | Pointer to **string** |  | [optional] [readonly] 
**LastErrorStatus** | Pointer to [**ErrorStatus**](ErrorStatus.md) |  | [optional] 
**LastErrorTime** | Pointer to **string** |  | [optional] [readonly] 
**LastEventTime** | Pointer to **string** |  | [optional] [readonly] 
**LastHeartbeatTime** | Pointer to **string** |  | [optional] [readonly] 
**LastStateTime** | Pointer to **string** |  | [optional] [readonly] 
**LogLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Config** | Pointer to [**DeviceConfig**](DeviceConfig.md) |  | [optional] 
**State** | Pointer to [**DeviceState**](DeviceState.md) |  | [optional] 
**Policy** | Pointer to [**Policy**](Policy.md) |  | [optional] 
**CustomOnboardData** | Pointer to **string** |  | [optional] 
**IsApprove** | Pointer to **bool** |  | [optional] 

## Methods

### NewCustomOnboard

`func NewCustomOnboard(id string, ) *CustomOnboard`

NewCustomOnboard instantiates a new CustomOnboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomOnboardWithDefaults

`func NewCustomOnboardWithDefaults() *CustomOnboard`

NewCustomOnboardWithDefaults instantiates a new CustomOnboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomOnboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomOnboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomOnboard) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomOnboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomOnboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomOnboard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomOnboard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumId

`func (o *CustomOnboard) GetNumId() string`

GetNumId returns the NumId field if non-nil, zero value otherwise.

### GetNumIdOk

`func (o *CustomOnboard) GetNumIdOk() (*string, bool)`

GetNumIdOk returns a tuple with the NumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumId

`func (o *CustomOnboard) SetNumId(v string)`

SetNumId sets NumId field to given value.

### HasNumId

`func (o *CustomOnboard) HasNumId() bool`

HasNumId returns a boolean if a field has been set.

### GetParent

`func (o *CustomOnboard) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CustomOnboard) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CustomOnboard) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *CustomOnboard) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRegistry

`func (o *CustomOnboard) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *CustomOnboard) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *CustomOnboard) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *CustomOnboard) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetBlocked

`func (o *CustomOnboard) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *CustomOnboard) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *CustomOnboard) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *CustomOnboard) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetCapresent

`func (o *CustomOnboard) GetCapresent() bool`

GetCapresent returns the Capresent field if non-nil, zero value otherwise.

### GetCapresentOk

`func (o *CustomOnboard) GetCapresentOk() (*bool, bool)`

GetCapresentOk returns a tuple with the Capresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapresent

`func (o *CustomOnboard) SetCapresent(v bool)`

SetCapresent sets Capresent field to given value.

### HasCapresent

`func (o *CustomOnboard) HasCapresent() bool`

HasCapresent returns a boolean if a field has been set.

### GetSubscription

`func (o *CustomOnboard) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CustomOnboard) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CustomOnboard) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *CustomOnboard) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetCreatedOn

`func (o *CustomOnboard) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *CustomOnboard) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *CustomOnboard) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *CustomOnboard) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *CustomOnboard) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *CustomOnboard) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *CustomOnboard) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *CustomOnboard) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetCredentials

`func (o *CustomOnboard) GetCredentials() []DeviceCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CustomOnboard) GetCredentialsOk() (*[]DeviceCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CustomOnboard) SetCredentials(v []DeviceCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CustomOnboard) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetGateway

`func (o *CustomOnboard) GetGateway() []string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CustomOnboard) GetGatewayOk() (*[]string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CustomOnboard) SetGateway(v []string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CustomOnboard) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayConfig

`func (o *CustomOnboard) GetGatewayConfig() GatewayConfig`

GetGatewayConfig returns the GatewayConfig field if non-nil, zero value otherwise.

### GetGatewayConfigOk

`func (o *CustomOnboard) GetGatewayConfigOk() (*GatewayConfig, bool)`

GetGatewayConfigOk returns a tuple with the GatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConfig

`func (o *CustomOnboard) SetGatewayConfig(v GatewayConfig)`

SetGatewayConfig sets GatewayConfig field to given value.

### HasGatewayConfig

`func (o *CustomOnboard) HasGatewayConfig() bool`

HasGatewayConfig returns a boolean if a field has been set.

### GetIsGateway

`func (o *CustomOnboard) GetIsGateway() bool`

GetIsGateway returns the IsGateway field if non-nil, zero value otherwise.

### GetIsGatewayOk

`func (o *CustomOnboard) GetIsGatewayOk() (*bool, bool)`

GetIsGatewayOk returns a tuple with the IsGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGateway

`func (o *CustomOnboard) SetIsGateway(v bool)`

SetIsGateway sets IsGateway field to given value.

### HasIsGateway

`func (o *CustomOnboard) HasIsGateway() bool`

HasIsGateway returns a boolean if a field has been set.

### GetDeviceErrors

`func (o *CustomOnboard) GetDeviceErrors() string`

GetDeviceErrors returns the DeviceErrors field if non-nil, zero value otherwise.

### GetDeviceErrorsOk

`func (o *CustomOnboard) GetDeviceErrorsOk() (*string, bool)`

GetDeviceErrorsOk returns a tuple with the DeviceErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceErrors

`func (o *CustomOnboard) SetDeviceErrors(v string)`

SetDeviceErrors sets DeviceErrors field to given value.

### HasDeviceErrors

`func (o *CustomOnboard) HasDeviceErrors() bool`

HasDeviceErrors returns a boolean if a field has been set.

### GetClientOnline

`func (o *CustomOnboard) GetClientOnline() bool`

GetClientOnline returns the ClientOnline field if non-nil, zero value otherwise.

### GetClientOnlineOk

`func (o *CustomOnboard) GetClientOnlineOk() (*bool, bool)`

GetClientOnlineOk returns a tuple with the ClientOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOnline

`func (o *CustomOnboard) SetClientOnline(v bool)`

SetClientOnline sets ClientOnline field to given value.

### HasClientOnline

`func (o *CustomOnboard) HasClientOnline() bool`

HasClientOnline returns a boolean if a field has been set.

### GetLastConfigAckTime

`func (o *CustomOnboard) GetLastConfigAckTime() string`

GetLastConfigAckTime returns the LastConfigAckTime field if non-nil, zero value otherwise.

### GetLastConfigAckTimeOk

`func (o *CustomOnboard) GetLastConfigAckTimeOk() (*string, bool)`

GetLastConfigAckTimeOk returns a tuple with the LastConfigAckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfigAckTime

`func (o *CustomOnboard) SetLastConfigAckTime(v string)`

SetLastConfigAckTime sets LastConfigAckTime field to given value.

### HasLastConfigAckTime

`func (o *CustomOnboard) HasLastConfigAckTime() bool`

HasLastConfigAckTime returns a boolean if a field has been set.

### GetLastConfigSendTime

`func (o *CustomOnboard) GetLastConfigSendTime() string`

GetLastConfigSendTime returns the LastConfigSendTime field if non-nil, zero value otherwise.

### GetLastConfigSendTimeOk

`func (o *CustomOnboard) GetLastConfigSendTimeOk() (*string, bool)`

GetLastConfigSendTimeOk returns a tuple with the LastConfigSendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfigSendTime

`func (o *CustomOnboard) SetLastConfigSendTime(v string)`

SetLastConfigSendTime sets LastConfigSendTime field to given value.

### HasLastConfigSendTime

`func (o *CustomOnboard) HasLastConfigSendTime() bool`

HasLastConfigSendTime returns a boolean if a field has been set.

### GetLastErrorStatus

`func (o *CustomOnboard) GetLastErrorStatus() ErrorStatus`

GetLastErrorStatus returns the LastErrorStatus field if non-nil, zero value otherwise.

### GetLastErrorStatusOk

`func (o *CustomOnboard) GetLastErrorStatusOk() (*ErrorStatus, bool)`

GetLastErrorStatusOk returns a tuple with the LastErrorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorStatus

`func (o *CustomOnboard) SetLastErrorStatus(v ErrorStatus)`

SetLastErrorStatus sets LastErrorStatus field to given value.

### HasLastErrorStatus

`func (o *CustomOnboard) HasLastErrorStatus() bool`

HasLastErrorStatus returns a boolean if a field has been set.

### GetLastErrorTime

`func (o *CustomOnboard) GetLastErrorTime() string`

GetLastErrorTime returns the LastErrorTime field if non-nil, zero value otherwise.

### GetLastErrorTimeOk

`func (o *CustomOnboard) GetLastErrorTimeOk() (*string, bool)`

GetLastErrorTimeOk returns a tuple with the LastErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorTime

`func (o *CustomOnboard) SetLastErrorTime(v string)`

SetLastErrorTime sets LastErrorTime field to given value.

### HasLastErrorTime

`func (o *CustomOnboard) HasLastErrorTime() bool`

HasLastErrorTime returns a boolean if a field has been set.

### GetLastEventTime

`func (o *CustomOnboard) GetLastEventTime() string`

GetLastEventTime returns the LastEventTime field if non-nil, zero value otherwise.

### GetLastEventTimeOk

`func (o *CustomOnboard) GetLastEventTimeOk() (*string, bool)`

GetLastEventTimeOk returns a tuple with the LastEventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventTime

`func (o *CustomOnboard) SetLastEventTime(v string)`

SetLastEventTime sets LastEventTime field to given value.

### HasLastEventTime

`func (o *CustomOnboard) HasLastEventTime() bool`

HasLastEventTime returns a boolean if a field has been set.

### GetLastHeartbeatTime

`func (o *CustomOnboard) GetLastHeartbeatTime() string`

GetLastHeartbeatTime returns the LastHeartbeatTime field if non-nil, zero value otherwise.

### GetLastHeartbeatTimeOk

`func (o *CustomOnboard) GetLastHeartbeatTimeOk() (*string, bool)`

GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatTime

`func (o *CustomOnboard) SetLastHeartbeatTime(v string)`

SetLastHeartbeatTime sets LastHeartbeatTime field to given value.

### HasLastHeartbeatTime

`func (o *CustomOnboard) HasLastHeartbeatTime() bool`

HasLastHeartbeatTime returns a boolean if a field has been set.

### GetLastStateTime

`func (o *CustomOnboard) GetLastStateTime() string`

GetLastStateTime returns the LastStateTime field if non-nil, zero value otherwise.

### GetLastStateTimeOk

`func (o *CustomOnboard) GetLastStateTimeOk() (*string, bool)`

GetLastStateTimeOk returns a tuple with the LastStateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStateTime

`func (o *CustomOnboard) SetLastStateTime(v string)`

SetLastStateTime sets LastStateTime field to given value.

### HasLastStateTime

`func (o *CustomOnboard) HasLastStateTime() bool`

HasLastStateTime returns a boolean if a field has been set.

### GetLogLevel

`func (o *CustomOnboard) GetLogLevel() LogLevel`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *CustomOnboard) GetLogLevelOk() (*LogLevel, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *CustomOnboard) SetLogLevel(v LogLevel)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *CustomOnboard) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomOnboard) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomOnboard) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomOnboard) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomOnboard) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConfig

`func (o *CustomOnboard) GetConfig() DeviceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CustomOnboard) GetConfigOk() (*DeviceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CustomOnboard) SetConfig(v DeviceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CustomOnboard) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetState

`func (o *CustomOnboard) GetState() DeviceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomOnboard) GetStateOk() (*DeviceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomOnboard) SetState(v DeviceState)`

SetState sets State field to given value.

### HasState

`func (o *CustomOnboard) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPolicy

`func (o *CustomOnboard) GetPolicy() Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CustomOnboard) GetPolicyOk() (*Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CustomOnboard) SetPolicy(v Policy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CustomOnboard) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetCustomOnboardData

`func (o *CustomOnboard) GetCustomOnboardData() string`

GetCustomOnboardData returns the CustomOnboardData field if non-nil, zero value otherwise.

### GetCustomOnboardDataOk

`func (o *CustomOnboard) GetCustomOnboardDataOk() (*string, bool)`

GetCustomOnboardDataOk returns a tuple with the CustomOnboardData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOnboardData

`func (o *CustomOnboard) SetCustomOnboardData(v string)`

SetCustomOnboardData sets CustomOnboardData field to given value.

### HasCustomOnboardData

`func (o *CustomOnboard) HasCustomOnboardData() bool`

HasCustomOnboardData returns a boolean if a field has been set.

### GetIsApprove

`func (o *CustomOnboard) GetIsApprove() bool`

GetIsApprove returns the IsApprove field if non-nil, zero value otherwise.

### GetIsApproveOk

`func (o *CustomOnboard) GetIsApproveOk() (*bool, bool)`

GetIsApproveOk returns a tuple with the IsApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApprove

`func (o *CustomOnboard) SetIsApprove(v bool)`

SetIsApprove sets IsApprove field to given value.

### HasIsApprove

`func (o *CustomOnboard) HasIsApprove() bool`

HasIsApprove returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


