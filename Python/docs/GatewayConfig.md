# GatewayConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**gateway_auth_method** | **str** |  | [optional] 
**gateway_type** | **str** |  | [optional] 

## Example

```python
from OmniCore.models.gateway_config import GatewayConfig

# TODO update the JSON string below
json = "{}"
# create an instance of GatewayConfig from a JSON string
gateway_config_instance = GatewayConfig.from_json(json)
# print the JSON string representation of the object
print GatewayConfig.to_json()

# convert the object into a dict
gateway_config_dict = gateway_config_instance.to_dict()
# create an instance of GatewayConfig from a dict
gateway_config_form_dict = gateway_config.from_dict(gateway_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


