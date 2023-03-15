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
    subscriptionid = 'Insert Subscription Here'  # str | Subscription ID
    registry_id = 'Insert Registry Here'  # str | Registry ID
    device_id = 'Insert Device ID Here'  # str | Device ID
    # int | The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available. (optional)
    num_states = 56

    try:
        api_response = api_instance.get_states(
            subscriptionid, registry_id, device_id, num_states=num_states)
        print("The response of DeviceApi->get_states:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->get_states: %s\n" % e)
