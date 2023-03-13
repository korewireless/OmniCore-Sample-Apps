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

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.RegistryApi(api_client)
    subscription_id = 'Insert Subscription Here'  # str | Subscription ID
    registry_id = 'Insert Registry Here'  # str | Registry ID

    try:
        # Get Registry
        api_response = api_instance.get_registry(subscription_id, registry_id)
        print("The response of RegistryApi->get_registry:\n")
        print("Registry Id:%s" % api_response.id)
        print("Registry Full Name:%s" % api_response.name)
        print("Certificate 1:%s" %
              api_response.credentials[0].public_key_certificate.certificate)
        print("Event Notification Config:%s" %
              api_response.event_notification_configs)
        print("State Notification Config:%s" %
              api_response.state_notification_config)
        print("Log Level:%s" % api_response.log_level)
        print("Created On:%s" % api_response.created_on)
        print("Updated On:%s" % api_response.updated_on)
    except Exception as e:
        print("Exception when calling RegistryApi->get_registry: %s\n" % e)
