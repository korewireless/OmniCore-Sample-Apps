# Device


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**name** | **str** |  | [optional] [readonly] 
**num_id** | **str** |  | [optional] [readonly] 
**parent** | **str** |  | [optional] [readonly] 
**registry** | **str** |  | [optional] [readonly] 
**blocked** | **bool** |  | [optional] 
**capresent** | **bool** |  | [optional] [readonly] 
**subscription** | **str** |  | [optional] [readonly] 
**created_on** | **str** |  | [optional] [readonly] 
**updated_on** | **str** |  | [optional] [readonly] 
**credentials** | [**List[DeviceCredential]**](DeviceCredential.md) |  | [optional] 
**gateway** | **List[str]** |  | [optional] 
**gateway_config** | [**GatewayConfig**](GatewayConfig.md) |  | [optional] 
**is_gateway** | **bool** |  | [optional] 
**device_errors** | **str** |  | [optional] [readonly] 
**client_online** | **bool** |  | [optional] [readonly] 
**last_config_ack_time** | **str** |  | [optional] [readonly] 
**last_config_send_time** | **str** |  | [optional] [readonly] 
**last_error_status** | [**ErrorStatus**](ErrorStatus.md) |  | [optional] 
**last_error_time** | **str** |  | [optional] [readonly] 
**last_event_time** | **str** |  | [optional] [readonly] 
**last_heartbeat_time** | **str** |  | [optional] [readonly] 
**last_state_time** | **str** |  | [optional] [readonly] 
**log_level** | [**LogLevel**](LogLevel.md) |  | [optional] 
**metadata** | **Dict[str, str]** |  | [optional] 
**config** | [**DeviceConfig**](DeviceConfig.md) |  | [optional] 
**state** | [**DeviceState**](DeviceState.md) |  | [optional] 

## Example

```python
from OmniCore.models.device import Device

# TODO update the JSON string below
json = "{}"
# create an instance of Device from a JSON string
device_instance = Device.from_json(json)
# print the JSON string representation of the object
print Device.to_json()

# convert the object into a dict
device_dict = device_instance.to_dict()
# create an instance of Device from a dict
device_form_dict = device.from_dict(device_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


