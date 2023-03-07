# HttpConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**http_enabled_state** | **str** | HttpEnabledState: If enabled, allows devices to use DeviceService via the HTTP protocol. Otherwise, any requests to DeviceService will fail for this registry.  Possible values:   \&quot;HTTP_STATE_UNSPECIFIED\&quot; - No HTTP state specified. If not specified, DeviceService will be enabled by default.   \&quot;HTTP_ENABLED\&quot; - Enables DeviceService (HTTP) service for the registry.   \&quot;HTTP_DISABLED\&quot; - Disables DeviceService (HTTP) service for the registry. | [optional] 

## Example

```python
from OmniCore.models.http_config import HttpConfig

# TODO update the JSON string below
json = "{}"
# create an instance of HttpConfig from a JSON string
http_config_instance = HttpConfig.from_json(json)
# print the JSON string representation of the object
print HttpConfig.to_json()

# convert the object into a dict
http_config_dict = http_config_instance.to_dict()
# create an instance of HttpConfig from a dict
http_config_form_dict = http_config.from_dict(http_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


