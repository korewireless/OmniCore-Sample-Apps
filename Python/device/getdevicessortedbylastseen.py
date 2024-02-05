from __future__ import print_function
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
    ##host="https://api.korewireless.com",
    access_token=access_token,
    api_key={"apiKey": "Insert Api Key Here"}
)

# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    subscription_id = 'Insert Subscription Here'  # str | Subscription ID
    # List[str] | A list of device string IDs. For example, ['device0', 'device12']. If empty, this field is ignored. Maximum IDs: 10,000 (optional)
    device_ids = ['DeviceId1','DeviceId2'] #list of device ids
   

    try:
        # Get All Devices
        api_response = api_instance.get_devices_last_seen(subscription_id, device_ids=device_ids,sort_by_client_online=True)
        print("The response of DeviceApi->get_devices:\n")
        print("Device Id:%s" % api_response.devices[0].id)
        print("Device LastHeart :%s" % api_response.devices[0].last_heartbeat_time)
        print("Device Online:%s" % api_response.devices[0].client_online)
        print("Device Subscription:%s" % api_response.devices[0].subscription)
       
    except Exception as e:
        print("Exception when calling DeviceApi->get_devices: %s\n" % e)
