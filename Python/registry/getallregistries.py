from __future__ import print_function
import time
import os
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
    api_instance = OmniCore.RegistryApi(api_client)
    subscription_id = 'Insert Subscription Here'  # str | Subscription ID
    page_number = 1  # int | Page Number (optional)
    page_size = 10  # int | Page Size (optional)

    try:
        # Get All Registries
        api_response = api_instance.get_registries(
            subscription_id, page_number=page_number, page_size=page_size)
        print("The response of RegistryApi->get_registries:\n")
        print("Registry Id:%s" % api_response.device_registries[0].id)
        print("Registry Full Name:%s" % api_response.device_registries[0].name)
        print("Certificate 1:%s" %
              api_response.device_registries[0].credentials[0].public_key_certificate.certificate)
        print("Event Notification Config:%s" %
              api_response.device_registries[0].event_notification_configs)
        print("State Notification Config:%s" %
              api_response.device_registries[0].state_notification_config)
        print("Log Level:%s" % api_response.device_registries[0].log_level)
        print("Custom Onboard Status:%s" %
              api_response.device_registries[0].custom_onboard_enabled)
        print("Custom Onboard Notification Config:%s" %
              api_response.device_registries[0].custom_onboard_notification_config)
        print("Created On:%s" % api_response.device_registries[0].created_on)
        print("Updated On:%s" % api_response.device_registries[0].updated_on)
    except Exception as e:
        print("Exception when calling RegistryApi->get_registries: %s\n" % e)
