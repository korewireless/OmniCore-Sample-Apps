# OmniCore.DeviceApi

All URIs are relative to *https://demo-api.omnicore.cloud.korewireless.com/model-state-management*

Method | HTTP request | Description
------------- | ------------- | -------------
[**bind_device**](DeviceApi.md#bind_device) | **POST** /subscriptions/{subscriptionId}/registries/{registryId}/bindDeviceToGateway | 
[**bind_devices**](DeviceApi.md#bind_devices) | **POST** /subscriptions/{subscriptionId}/registries/{registryId}/bindDevicesToGateway | 
[**block_device_communcation**](DeviceApi.md#block_device_communcation) | **PUT** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/communication | 
[**create_device**](DeviceApi.md#create_device) | **POST** /subscriptions/{subscriptionId}/registries/{registryId}/devices | 
[**delete_device**](DeviceApi.md#delete_device) | **DELETE** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
[**get_config**](DeviceApi.md#get_config) | **GET** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/configVersions | 
[**get_device**](DeviceApi.md#get_device) | **GET** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 
[**get_devices**](DeviceApi.md#get_devices) | **GET** /subscriptions/{subscriptionId}/registries/{registryId}/devices | 
[**get_states**](DeviceApi.md#get_states) | **GET** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/states | 
[**send_command_to_device**](DeviceApi.md#send_command_to_device) | **POST** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/sendCommandToDevice | 
[**un_bind_device**](DeviceApi.md#un_bind_device) | **POST** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDeviceFromGateway | 
[**un_bind_devices**](DeviceApi.md#un_bind_devices) | **POST** /subscriptions/{subscriptionId}/registries/{registryId}/unbindDevicesFromGateway | 
[**update_configuration_to_device**](DeviceApi.md#update_configuration_to_device) | **POST** /subscriptions/{subscriptionid}/registries/{registryId}/devices/{deviceId}/updateConfigurationToDevice | 
[**update_device**](DeviceApi.md#update_device) | **PATCH** /subscriptions/{subscriptionId}/registries/{registryId}/devices/{deviceId} | 


# **bind_device**
> Info bind_device(subscription_id, registry_id, device)



Bind  a device to a gateway under a registry

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device = OmniCore.BindRequest() # BindRequest | application/json

    try:
        api_response = api_instance.bind_device(subscription_id, registry_id, device)
        print("The response of DeviceApi->bind_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->bind_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device** | [**BindRequest**](BindRequest.md)| application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **bind_devices**
> Info bind_devices(subscription_id, registry_id, device)



Bind devices to a gateway under a registry

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device = OmniCore.BindRequestIdsGateway() # BindRequestIdsGateway | application/json

    try:
        api_response = api_instance.bind_devices(subscription_id, registry_id, device)
        print("The response of DeviceApi->bind_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->bind_devices: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device** | [**BindRequestIdsGateway**](BindRequestIdsGateway.md)| application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **block_device_communcation**
> object block_device_communcation(subscriptionid, registry_id, device_id, device)



Blocks All Communication From A Device

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscriptionid = 'subscriptionid_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device_id = 'device_id_example' # str | Device ID
    device = OmniCore.BlockCommunicationBody() # BlockCommunicationBody | application/json

    try:
        api_response = api_instance.block_device_communcation(subscriptionid, registry_id, device_id, device)
        print("The response of DeviceApi->block_device_communcation:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->block_device_communcation: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device_id** | **str**| Device ID | 
 **device** | [**BlockCommunicationBody**](BlockCommunicationBody.md)| application/json | 

### Return type

**object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_device**
> Device create_device(subscription_id, registry_id, device)



Create a device under a registry

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device = OmniCore.CreateNewDevice() # CreateNewDevice | application/json

    try:
        api_response = api_instance.create_device(subscription_id, registry_id, device)
        print("The response of DeviceApi->create_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->create_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device** | [**CreateNewDevice**](CreateNewDevice.md)| application/json | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_device**
> Info delete_device(subscription_id, registry_id, device_id)



Delete a device under a registry

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device_id = 'device_id_example' # str | Device ID

    try:
        api_response = api_instance.delete_device(subscription_id, registry_id, device_id)
        print("The response of DeviceApi->delete_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->delete_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device_id** | **str**| Device ID | 

### Return type

[**Info**](Info.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_config**
> ListDeviceConfigVersionsResponse get_config(subscriptionid, registry_id, device_id, num_versions)



Get Configs Of Devices

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscriptionid = 'subscriptionid_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device_id = 'device_id_example' # str | Device ID
    num_versions = 56 # int | Device ID

    try:
        api_response = api_instance.get_config(subscriptionid, registry_id, device_id, num_versions)
        print("The response of DeviceApi->get_config:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->get_config: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device_id** | **str**| Device ID | 
 **num_versions** | **int**| Device ID | 

### Return type

[**ListDeviceConfigVersionsResponse**](ListDeviceConfigVersionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device**
> Device get_device(registry_id, subscription_id, device_id)



Get a device under a registry

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    registry_id = 'registry_id_example' # str | Registry ID
    subscription_id = 'subscription_id_example' # str | Subscription ID
    device_id = 'device_id_example' # str | Device ID

    try:
        api_response = api_instance.get_device(registry_id, subscription_id, device_id)
        print("The response of DeviceApi->get_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->get_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registry_id** | **str**| Registry ID | 
 **subscription_id** | **str**| Subscription ID | 
 **device_id** | **str**| Device ID | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_devices**
> ListDevicesResponse get_devices(registry_id, subscription_id, page_number=page_number, page_size=page_size, field_mask=field_mask, device_ids=device_ids, device_num_ids=device_num_ids, gateway_list_options_associations_device_id=gateway_list_options_associations_device_id, gateway_list_options_associations_gateway_id=gateway_list_options_associations_gateway_id, gateway_list_options_gateway_type=gateway_list_options_gateway_type)



Get all devices under a registry

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    registry_id = 'registry_id_example' # str | Registry ID
    subscription_id = 'subscription_id_example' # str | Subscription ID
    page_number = 56 # int | Page Number (optional)
    page_size = 56 # int | The maximum number of devices to return in the response. If this value is zero, the service will select a default size.  (optional)
    field_mask = 'field_mask_example' # str | The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example:  (optional)
    device_ids = ['device_ids_example'] # List[str] | A list of device string IDs. For example, ['device0', 'device12']. If empty, this field is ignored. Maximum IDs: 10,000 (optional)
    device_num_ids = ['device_num_ids_example'] # List[str] | A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000. (optional)
    gateway_list_options_associations_device_id = 'gateway_list_options_associations_device_id_example' # str | If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound. (optional)
    gateway_list_options_associations_gateway_id = 'gateway_list_options_associations_gateway_id_example' # str | If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned (optional)
    gateway_list_options_gateway_type = 'gateway_list_options_gateway_type_example' # str | If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned. (optional)

    try:
        api_response = api_instance.get_devices(registry_id, subscription_id, page_number=page_number, page_size=page_size, field_mask=field_mask, device_ids=device_ids, device_num_ids=device_num_ids, gateway_list_options_associations_device_id=gateway_list_options_associations_device_id, gateway_list_options_associations_gateway_id=gateway_list_options_associations_gateway_id, gateway_list_options_gateway_type=gateway_list_options_gateway_type)
        print("The response of DeviceApi->get_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->get_devices: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registry_id** | **str**| Registry ID | 
 **subscription_id** | **str**| Subscription ID | 
 **page_number** | **int**| Page Number | [optional] 
 **page_size** | **int**| The maximum number of devices to return in the response. If this value is zero, the service will select a default size.  | [optional] 
 **field_mask** | **str**| The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example:  | [optional] 
 **device_ids** | [**List[str]**](str.md)| A list of device string IDs. For example, [&#39;device0&#39;, &#39;device12&#39;]. If empty, this field is ignored. Maximum IDs: 10,000 | [optional] 
 **device_num_ids** | [**List[str]**](str.md)| A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000. | [optional] 
 **gateway_list_options_associations_device_id** | **str**| If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound. | [optional] 
 **gateway_list_options_associations_gateway_id** | **str**| If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned | [optional] 
 **gateway_list_options_gateway_type** | **str**| If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned. | [optional] 

### Return type

[**ListDevicesResponse**](ListDevicesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_states**
> ListDeviceStatesResponse get_states(subscriptionid, registry_id, device_id, num_states=num_states)



Get States Of Devices

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscriptionid = 'subscriptionid_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device_id = 'device_id_example' # str | Device ID
    num_states = 56 # int | The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available. (optional)

    try:
        api_response = api_instance.get_states(subscriptionid, registry_id, device_id, num_states=num_states)
        print("The response of DeviceApi->get_states:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->get_states: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device_id** | **str**| Device ID | 
 **num_states** | **int**| The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available. | [optional] 

### Return type

[**ListDeviceStatesResponse**](ListDeviceStatesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **send_command_to_device**
> object send_command_to_device(subscriptionid, registry_id, device_id, device)



Send A Command To A Device

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscriptionid = 'subscriptionid_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device_id = 'device_id_example' # str | Device ID
    device = OmniCore.DeviceCommand() # DeviceCommand | application/json

    try:
        api_response = api_instance.send_command_to_device(subscriptionid, registry_id, device_id, device)
        print("The response of DeviceApi->send_command_to_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->send_command_to_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device_id** | **str**| Device ID | 
 **device** | [**DeviceCommand**](DeviceCommand.md)| application/json | 

### Return type

**object**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **un_bind_device**
> Info un_bind_device(subscription_id, registry_id, device)



UnBind  a device from a gateway under a registry

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device = OmniCore.BindRequest() # BindRequest | application/json

    try:
        api_response = api_instance.un_bind_device(subscription_id, registry_id, device)
        print("The response of DeviceApi->un_bind_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->un_bind_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device** | [**BindRequest**](BindRequest.md)| application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **un_bind_devices**
> Info un_bind_devices(subscription_id, registry_id, device)



UnBind devices from a gateway under a registry

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device = OmniCore.BindRequestIdsGateway() # BindRequestIdsGateway | application/json

    try:
        api_response = api_instance.un_bind_devices(subscription_id, registry_id, device)
        print("The response of DeviceApi->un_bind_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->un_bind_devices: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device** | [**BindRequestIdsGateway**](BindRequestIdsGateway.md)| application/json | 

### Return type

[**Info**](Info.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_configuration_to_device**
> DeviceConfig update_configuration_to_device(subscriptionid, registry_id, device_id, device)



Update A Configuration Of A Device

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscriptionid = 'subscriptionid_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device_id = 'device_id_example' # str | Device ID
    device = OmniCore.DeviceConfiguration() # DeviceConfiguration | application/json

    try:
        api_response = api_instance.update_configuration_to_device(subscriptionid, registry_id, device_id, device)
        print("The response of DeviceApi->update_configuration_to_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->update_configuration_to_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionid** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device_id** | **str**| Device ID | 
 **device** | [**DeviceConfiguration**](DeviceConfiguration.md)| application/json | 

### Return type

[**DeviceConfig**](DeviceConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_device**
> Device update_device(subscription_id, registry_id, device_id, update_mask, device)



Modify device under a registry

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://demo-api.omnicore.cloud.korewireless.com/model-state-management"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    device_id = 'device_id_example' # str | Device ID
    update_mask = 'update_mask_example' # str | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,log_level, blocked, and metadata
    device = OmniCore.UpdateDevice() # UpdateDevice | application/json

    try:
        api_response = api_instance.update_device(subscription_id, registry_id, device_id, update_mask, device)
        print("The response of DeviceApi->update_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->update_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **device_id** | **str**| Device ID | 
 **update_mask** | **str**| Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,log_level, blocked, and metadata | 
 **device** | [**UpdateDevice**](UpdateDevice.md)| application/json | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**400** | Bad Request |  -  |
**404** | Not Found |  -  |
**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

