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


# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscriptionid = 'Insert Subscription Id Here'  # str | Subscription ID
    registry_id = 'Insert Registry Id Here'  # str | Registry ID
    device_id = 'Insert Device Id Here'  # str | Device ID
    # int | The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available. (optional)
    num_states = 0

    try:
        api_response = api_instance.get_states(
            subscriptionid, registry_id, device_id, num_states=num_states)
        print("The response of DeviceApi->get_states:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->get_states: %s\n" % e)
