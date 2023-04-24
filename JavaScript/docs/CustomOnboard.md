# OmniCoreModelAndStateManagementApi.CustomOnboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **String** |  | 
**name** | **String** |  | [optional] [readonly] 
**numId** | **String** |  | [optional] [readonly] 
**parent** | **String** |  | [optional] [readonly] 
**registry** | **String** |  | [optional] [readonly] 
**blocked** | **Boolean** |  | [optional] 
**capresent** | **Boolean** |  | [optional] [readonly] 
**subscription** | **String** |  | [optional] [readonly] 
**createdOn** | **String** |  | [optional] [readonly] 
**updatedOn** | **String** |  | [optional] [readonly] 
**credentials** | [**[DeviceCredential]**](DeviceCredential.md) |  | [optional] 
**gateway** | **[String]** |  | [optional] 
**gatewayConfig** | [**GatewayConfig**](GatewayConfig.md) |  | [optional] 
**isGateway** | **Boolean** |  | [optional] 
**deviceErrors** | **String** |  | [optional] [readonly] 
**clientOnline** | **Boolean** |  | [optional] [readonly] 
**lastConfigAckTime** | **String** |  | [optional] [readonly] 
**lastConfigSendTime** | **String** |  | [optional] [readonly] 
**lastErrorStatus** | [**ErrorStatus**](ErrorStatus.md) |  | [optional] 
**lastErrorTime** | **String** |  | [optional] [readonly] 
**lastEventTime** | **String** |  | [optional] [readonly] 
**lastHeartbeatTime** | **String** |  | [optional] [readonly] 
**lastStateTime** | **String** |  | [optional] [readonly] 
**logLevel** | [**LogLevel**](LogLevel.md) |  | [optional] 
**metadata** | **{String: String}** |  | [optional] 
**config** | [**DeviceConfig**](DeviceConfig.md) |  | [optional] 
**state** | [**DeviceState**](DeviceState.md) |  | [optional] 
**policy** | [**Policy**](Policy.md) |  | [optional] 
**customOnboardData** | **String** |  | [optional] 
**isApprove** | **Boolean** |  | [optional] 


