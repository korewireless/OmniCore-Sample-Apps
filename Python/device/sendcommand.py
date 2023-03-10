from __future__ import print_function
import time
import os
import base64
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
    data_to_send = "Hi from omnicore"
    base64_string: str = encode_to_base64(data_to_send)
    # DeviceCommand | application/json
    device = OmniCore.DeviceCommand(binaryData=base64_string)
    device.subfolder = ""

    try:
        api_response = api_instance.send_command_to_device(
            subscriptionid, registry_id, device_id, device)
        print("The response of DeviceApi->send_command_to_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->send_command_to_device: %s\n" % e)
