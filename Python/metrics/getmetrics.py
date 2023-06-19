from __future__ import print_function
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
from generatetoken import fetchToken
# Defining the host is optional and defaults to https://api.korewireless.com/omnicore
# See configuration.py for a list of all supported configuration parameters.

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Step 1 Generate Client Token,Inputs are clientId And Client Secret Which can be fetched from ui keys console,
# If Token is not expired new token is not used, if token is expired new token is generated.

access_token = fetchToken("Insert previous token if used,leave empty if no previous token",
                          'Insert Client Id Here', 'Insert Client Secret Here')

# API key is a static key which can be fetched from keys console
# Configure access token and apikey
configuration = OmniCore.Configuration(
    host="https://api.korewireless.com/omnicore",
    access_token=access_token,
    api_key={"apiKey": "Insert Api Key Here"}
)

## Properties Of Metrics Response
# Name | Type | Description | Notes
# ------------ | ------------- | ------------- | -------------
# **no_of_messages_for30_minutes** | **List[object]** |  | [optional] 
# **no_of_messages_for48_hours** | **List[object]** |  | [optional] 
# **billable_bytes_received** | **int** |  | [optional] 
# **billable_bytes_sent** | **int** |  | [optional] 
# **billable_message_size** | **int** |  | [optional] 
# **bytes_received** | **int** |  | [optional] 
# **bytes_sent** | **int** |  | [optional] 
# **message_size** | **int** |  | [optional] 
# **no_of_ack_messages** | **int** |  | [optional] 
# **no_of_command_messages** | **int** |  | [optional] 
# **no_of_config_messages** | **int** |  | [optional] 
# **no_of_device_connections_failed** | **int** |  | [optional] 
# **no_of_devices** | **int** |  | [optional] 
# **no_of_dis_connections** | **int** |  | [optional] 
# **no_of_event_messages** | **int** |  | [optional] 
# **no_of_gateway_connections_failed** | **int** |  | [optional] 
# **no_of_gateways** | **int** |  | [optional] 
# **no_of_loop_back_messages** | **int** |  | [optional] 
# **no_of_messages** | **int** |  | [optional] 
# **no_of_publish_errors** | **int** |  | [optional] 
# **no_of_registries** | **int** |  | [optional] 
# **no_of_state_messages** | **int** |  | [optional] 
# **no_of_subscribe** | **int** |  | [optional] 
# **no_of_successful_connections** | **int** |  | [optional] 
# **no_of_un_subscribe** | **int** |  | [optional] 
# **subscription_id** | **str** |  | [optional] 
# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.MetricsApi(api_client)
    subscription_id = 'Insert Subscription Here'  # str | Subscription ID

    try:
        # Get Device
        api_response = api_instance.get_metrics(
            subscription_id)
        print("The response of MetricsApi->get_metricse:\n")
        print("No Of Devices:%s" % api_response.details.no_of_devices)
        print("Device Connections:%s" % api_response.details.no_of_successful_connections)
        print("Device DisConnections:%s" % api_response.details.no_of_dis_connections)
        print("Device Messages For 30 Minutes:%s" % api_response.details.no_of_messages_for30_minutes)


    except Exception as e:
        print("Exception when calling MetricsApi->get_metrics: %s\n" % e)
