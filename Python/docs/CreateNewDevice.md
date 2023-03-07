# CreateNewDevice


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**blocked** | **bool** |  | [optional] 
**credentials** | [**List[DeviceCredential]**](DeviceCredential.md) |  | [optional] 
**gateway_config** | [**GatewayConfig**](GatewayConfig.md) |  | [optional] 
**config** | [**DeviceConfig**](DeviceConfig.md) |  | [optional] 
**log_level** | [**LogLevel**](LogLevel.md) |  | [optional] 
**metadata** | **Dict[str, str]** |  | [optional] 

## Example

```python
from OmniCore.models.create_new_device import CreateNewDevice

# TODO update the JSON string below
json = "{}"
# create an instance of CreateNewDevice from a JSON string
create_new_device_instance = CreateNewDevice.from_json(json)
# print the JSON string representation of the object
print CreateNewDevice.to_json()

# convert the object into a dict
create_new_device_dict = create_new_device_instance.to_dict()
# create an instance of CreateNewDevice from a dict
create_new_device_form_dict = create_new_device.from_dict(create_new_device_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


