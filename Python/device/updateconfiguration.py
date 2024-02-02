from __future__ import print_function
import base64
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
from generatetoken import fetchToken
# Defining the host is optional and defaults to https://api.korewireless.com
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
    #host="https://api.korewireless.com",
    access_token=access_token,
    api_key={"apiKey": "Insert Api Key Here"}
)


def encode_to_base64(input: str):
    try:
        sample_string_bytes = input.encode("ascii")
        base64_bytes = base64.b64encode(sample_string_bytes)
        base64_string: str = base64_bytes.decode("ascii")
        return base64_string
    except Exception as e:
        print("Exception when encoding as base64 %s\n" % e)


# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscriptionid = 'Insert Subscription Here'  # str | Subscription ID
    registry_id = 'Insert Registry Here'  # str | Registry ID
    device_id = 'shaizdev00'  # str | Device ID
    # DeviceConfiguration | application/json
    data_to_send: str = "Hi from omnicore"
    base64_string: str = encode_to_base64(data_to_send)
    device = OmniCore.DeviceConfiguration(binaryData=base64_string)
    device.subfolder = ""

    try:
        api_response = api_instance.update_configuration_to_device(
            subscriptionid, registry_id, device_id, device)
        print("The response of DeviceApi->update_configuration_to_device:\n")
        print("Cloud Update Time", api_response.cloud_update_time)
    except Exception as e:
        print("Exception when calling DeviceApi->update_configuration_to_device: %s\n" % e)
