# OmniCore.SinkApi

All URIs are relative to *https://api.korewireless.com/omnicore*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_sink**](SinkApi.md#create_sink) | **POST** /subscriptions/{subscriptionId}/sinks | 
[**delete_sink**](SinkApi.md#delete_sink) | **DELETE** /subscriptions/{subscriptionId}/sinks/{sinkId} | 
[**get_a_sink**](SinkApi.md#get_a_sink) | **GET** /subscriptions/{subscriptionId}/sinks/{sinkId} | 
[**get_sinks**](SinkApi.md#get_sinks) | **GET** /subscriptions/{subscriptionId}/sinks | Get All Sinks


# **create_sink**
> Sink create_sink(subscription_id, sink=sink)



Create a Sink

### Example

* Api Key Authentication (apiKey):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.korewireless.com/omnicore
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://api.korewireless.com/omnicore"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apiKey
configuration.api_key['apiKey'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apiKey'] = 'Bearer'

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.SinkApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    sink = OmniCore.Sink() # Sink |  (optional)

    try:
        api_response = api_instance.create_sink(subscription_id, sink=sink)
        print("The response of SinkApi->create_sink:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SinkApi->create_sink: %s\n" % e)
```

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.korewireless.com/omnicore
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://api.korewireless.com/omnicore"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apiKey
configuration.api_key['apiKey'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apiKey'] = 'Bearer'

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.SinkApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    sink = OmniCore.Sink() # Sink |  (optional)

    try:
        api_response = api_instance.create_sink(subscription_id, sink=sink)
        print("The response of SinkApi->create_sink:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SinkApi->create_sink: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **sink** | [**Sink**](Sink.md)|  | [optional] 

### Return type

[**Sink**](Sink.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**400** | Bad Request |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**404** | Not Found |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**500** | Internal Server Error |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_sink**
> object delete_sink(subscription_id, sink_id)



Delete a Sink

### Example

* Api Key Authentication (apiKey):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.korewireless.com/omnicore
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://api.korewireless.com/omnicore"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apiKey
configuration.api_key['apiKey'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apiKey'] = 'Bearer'

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.SinkApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    sink_id = 'sink_id_example' # str | Sink ID

    try:
        api_response = api_instance.delete_sink(subscription_id, sink_id)
        print("The response of SinkApi->delete_sink:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SinkApi->delete_sink: %s\n" % e)
```

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.korewireless.com/omnicore
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://api.korewireless.com/omnicore"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apiKey
configuration.api_key['apiKey'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apiKey'] = 'Bearer'

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.SinkApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    sink_id = 'sink_id_example' # str | Sink ID

    try:
        api_response = api_instance.delete_sink(subscription_id, sink_id)
        print("The response of SinkApi->delete_sink:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SinkApi->delete_sink: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **sink_id** | **str**| Sink ID | 

### Return type

**object**

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**400** | Bad Request |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**404** | Not Found |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**500** | Internal Server Error |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_a_sink**
> Sink get_a_sink(subscription_id, sink_id)



Get A Sink

### Example

* Api Key Authentication (apiKey):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.korewireless.com/omnicore
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://api.korewireless.com/omnicore"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apiKey
configuration.api_key['apiKey'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apiKey'] = 'Bearer'

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.SinkApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    sink_id = 'sink_id_example' # str | Sink ID

    try:
        api_response = api_instance.get_a_sink(subscription_id, sink_id)
        print("The response of SinkApi->get_a_sink:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SinkApi->get_a_sink: %s\n" % e)
```

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.korewireless.com/omnicore
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://api.korewireless.com/omnicore"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apiKey
configuration.api_key['apiKey'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apiKey'] = 'Bearer'

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.SinkApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID
    sink_id = 'sink_id_example' # str | Sink ID

    try:
        api_response = api_instance.get_a_sink(subscription_id, sink_id)
        print("The response of SinkApi->get_a_sink:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SinkApi->get_a_sink: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 
 **sink_id** | **str**| Sink ID | 

### Return type

[**Sink**](Sink.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**400** | Bad Request |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**404** | Not Found |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**500** | Internal Server Error |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_sinks**
> ListSinks get_sinks(subscription_id)

Get All Sinks

Get Sinks

### Example

* Api Key Authentication (apiKey):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.korewireless.com/omnicore
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://api.korewireless.com/omnicore"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apiKey
configuration.api_key['apiKey'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apiKey'] = 'Bearer'

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.SinkApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID

    try:
        # Get All Sinks
        api_response = api_instance.get_sinks(subscription_id)
        print("The response of SinkApi->get_sinks:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SinkApi->get_sinks: %s\n" % e)
```

* Bearer (JWT) Authentication (bearerAuth):
```python
from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.korewireless.com/omnicore
# See configuration.py for a list of all supported configuration parameters.
configuration = OmniCore.Configuration(
    host = "https://api.korewireless.com/omnicore"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apiKey
configuration.api_key['apiKey'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apiKey'] = 'Bearer'

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.SinkApi(api_client)
    subscription_id = 'subscription_id_example' # str | Subscription ID

    try:
        # Get All Sinks
        api_response = api_instance.get_sinks(subscription_id)
        print("The response of SinkApi->get_sinks:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SinkApi->get_sinks: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription_id** | **str**| Subscription ID | 

### Return type

[**ListSinks**](ListSinks.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**400** | Bad Request |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**404** | Not Found |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |
**500** | Internal Server Error |  * Access-Control-Allow-Origin -  <br>  * Access-Control-Allow-Methods -  <br>  * Access-Control-Allow-Credentials -  <br>  * Access-Control-Allow-Headers -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

