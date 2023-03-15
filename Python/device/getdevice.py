from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com
# See configuration.py for a list of all supported configuration parameters.

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    host="https://demo-api.omnicore.cloud.korewireless.com",
    access_token="Insert Token Here"
)
# Configure API key authorization: apiKey
# configuration.api_key['apiKey'] = "Insert API Key Here"


# ALTERNATIVE
# Configure Bearer authorization (JWT): bearerAuth
# configuration = OmniCore.Configuration(
#     access_token="Insert Token Here"
# )

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    registry_id = 'Insert Registry Here'  # str | Registry ID
    subscription_id = 'Insert Subscription Here'  # str | Subscription ID
    device_id = 'shaizdev00'  # str | Device ID

    try:
        # Get Device
        api_response = api_instance.get_device(
            registry_id, subscription_id, device_id)
        print("The response of DeviceApi->get_device:\n")
        print("Device Id:%s" % api_response.id)
        print("Device Full Name:%s" % api_response.name)
        print("Certificate 1:%s" %
              api_response.credentials[0].public_key)
        print("Log Level:%s" % api_response.log_level)
        print("Created On:%s" % api_response.created_on)
        print("Updated On:%s" % api_response.updated_on)

    except Exception as e:
        print("Exception when calling DeviceApi->get_device: %s\n" % e)
