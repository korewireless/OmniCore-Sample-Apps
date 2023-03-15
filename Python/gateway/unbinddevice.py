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

# Configure Either access token or apikey
configuration = OmniCore.Configuration(
    host="https://demo-api.omnicore.cloud.korewireless.com",
    access_token="Insert Token Here",
    api_key={"apiKey": "Insert API Key Here"}
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
    subscription_id = 'Insert Subscription Here'  # str | Subscription ID
    registry_id = 'Insert Registry Here'  # str | Registry ID
    device = OmniCore.BindRequest(deviceId="Insert Device Id Here",
                                  gatewayId="Insert Gateway Id here")  # BindRequest | application/json

    try:
        api_response = api_instance.un_bind_device(
            subscription_id, registry_id, device)
        print("The response of DeviceApi->un_bind_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->un_bind_device: %s\n" % e)
