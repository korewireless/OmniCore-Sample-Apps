# OmniCore.RegistryApi

All URIs are relative to *https://demo-api.omnicore.cloud.korewireless.com/model-state-management*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_registry**](RegistryApi.md#create_registry) | **POST** /subscriptions/{subscriptionId}/registries | 
[**delete_registry**](RegistryApi.md#delete_registry) | **DELETE** /subscriptions/{subscriptionId}/registries/{registryId} | 
[**get_registries**](RegistryApi.md#get_registries) | **GET** /subscriptions/{subscriptionId}/registries | 
[**get_registry**](RegistryApi.md#get_registry) | **GET** /subscriptions/{subscriptionId}/registries/{registryId} | 
[**update_registry**](RegistryApi.md#update_registry) | **PATCH** /subscriptions/{subscriptionId}/registries/{registryId} | 


# **create_registry**
> DeviceRegistry create_registry(subscription_id, registry=registry)



Create a registry

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
    api_instance = OmniCore.RegistryApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry = OmniCore.DeviceRegistry() # DeviceRegistry | application/json (optional)

    try:
        api_response = api_instance.create_registry(subscription_id, registry=registry)
        print("The response of RegistryApi->create_registry:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RegistryApi->create_registry: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry** | [**DeviceRegistry**](DeviceRegistry.md)| application/json | [optional] 

### Return type

[**DeviceRegistry**](DeviceRegistry.md)

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

# **delete_registry**
> Info delete_registry(subscription_id, registry_id)



Delete a registry

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
    api_instance = OmniCore.RegistryApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID

    try:
        api_response = api_instance.delete_registry(subscription_id, registry_id)
        print("The response of RegistryApi->delete_registry:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RegistryApi->delete_registry: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 

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

# **get_registries**
> ListDeviceRegistries get_registries(subscription_id, page_number=page_number, page_size=page_size)



Get all registries under a subscription

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
    api_instance = OmniCore.RegistryApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    page_number = 56 # int | Page Number (optional)
    page_size = 56 # int | Page Size (optional)

    try:
        api_response = api_instance.get_registries(subscription_id, page_number=page_number, page_size=page_size)
        print("The response of RegistryApi->get_registries:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RegistryApi->get_registries: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **page_number** | **int**| Page Number | [optional] 
 **page_size** | **int**| Page Size | [optional] 

### Return type

[**ListDeviceRegistries**](ListDeviceRegistries.md)

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

# **get_registry**
> DeviceRegistry get_registry(subscription_id, registry_id)



Get a registry

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
    api_instance = OmniCore.RegistryApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID

    try:
        api_response = api_instance.get_registry(subscription_id, registry_id)
        print("The response of RegistryApi->get_registry:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RegistryApi->get_registry: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 

### Return type

[**DeviceRegistry**](DeviceRegistry.md)

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

# **update_registry**
> DeviceRegistry update_registry(subscription_id, registry_id, update_mask, registry=registry)



Update a registry

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
    api_instance = OmniCore.RegistryApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    registry_id = 'registry_id_example' # str | Registry ID
    update_mask = 'update_mask_example' # str | values to be updated: eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials
    registry = OmniCore.DeviceRegistry() # DeviceRegistry | application/json (optional)

    try:
        api_response = api_instance.update_registry(subscription_id, registry_id, update_mask, registry=registry)
        print("The response of RegistryApi->update_registry:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RegistryApi->update_registry: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **registry_id** | **str**| Registry ID | 
 **update_mask** | **str**| values to be updated: eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials | 
 **registry** | [**DeviceRegistry**](DeviceRegistry.md)| application/json | [optional] 

### Return type

[**DeviceRegistry**](DeviceRegistry.md)

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

