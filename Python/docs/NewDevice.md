# NewDevice


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
from OmniCore.models.new_device import NewDevice

# TODO update the JSON string below
json = "{}"
# create an instance of NewDevice from a JSON string
new_device_instance = NewDevice.from_json(json)
# print the JSON string representation of the object
print NewDevice.to_json()

# convert the object into a dict
new_device_dict = new_device_instance.to_dict()
# create an instance of NewDevice from a dict
new_device_form_dict = new_device.from_dict(new_device_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


