from __future__ import print_function
import time
import os
import OmniCore
from OmniCore.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://demo-api.omnicore.cloud.korewireless.com/model-state-management
# See configuration.py for a list of all supported configuration parameters.

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = OmniCore.Configuration(
    host="https://demo-api.omnicore.cloud.korewireless.com/model-state-management",
    access_token="eyJhbGciOiJSUzI1NiIsImtpZCI6ImY4NzZiNzIxNDAwYmZhZmEyOWQ0MTFmZTYwODE2YmRhZWMyM2IzODIiLCJ0eXAiOiJKV1QifQ.eyJyb2xlIjoiVGVuYW50QWRtaW4iLCJzdWJzY3JpcHRpb25JZCI6ImtvcmV3aXJlbGVzcy1kZXZlbG9wbWVudCIsImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS9nY3AtaW90LWNvcmUtZGVtby0zNjE2MTQiLCJhdWQiOiJnY3AtaW90LWNvcmUtZGVtby0zNjE2MTQiLCJhdXRoX3RpbWUiOjE2NzY1MzQ2MzAsInVzZXJfaWQiOiJCb2RaZEQ0YVhpWHRSMVNUR0pKWTlTaGZuTW0yIiwic3ViIjoiQm9kWmRENGFYaVh0UjFTVEdKSlk5U2hmbk1tMiIsImlhdCI6MTY3ODE2MDk0MiwiZXhwIjoxNjc4MTY0NTQyLCJlbWFpbCI6InNoZWh6YWRAa29yZXdpcmVsZXNzLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJzaGVoemFkQGtvcmV3aXJlbGVzcy5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJwYXNzd29yZCIsInRlbmFudCI6ImtvcmV3aXJlbGVzcy1jb20taWUzZ3IifX0.ZqQYWmZ7JWvKOTwAt-OEon7RhPimEpPU9fxUrpJPrBc4mp682nZ1QYlGqjg9MzgxYQ5OMp5TeQbhSv0yyjnfQheOPcgweS4_hIUUV9PbXEOiyhD8UMjp3GNXDSNST6HHzZwqtAOBUG33LQNcJeEhNmgu6Skx_1GrFbv4A3GLDE0tkJXfI3OqYJbtXUkoHtt6mQ0tftY2M5FQ2RIjFvpmvnqneEz2M9Anavh9j_Uix7Ay-sXACz6S8YRAlV8VJ9I50jYV4M37ZcOksc0mo6n0wT9eg-UgMrEIGAZvO2ID-b0TOY5iIaQxKhEEENeHTwywOC2K0JCCWp6LKKOvZA6BSA"
)


# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.DeviceApi(api_client)
    registry_id = 'shaiz-registry-sdk'  # str | Registry ID
    subscription_id = 'korewireless-development'  # str | Subscription ID
    page_number = 1  # int | Page Number (optional)
    # int | The maximum number of devices to return in the response. If this value is zero, the service will select a default size.  (optional)
    page_size = 10
    # str | The fields of the Device resource to be returned to the response. The fields id and numId are always returned, along with any other fields specified. A comma-separated list of fully qualified names of fields. Example:  (optional)
    field_mask = ''
    # List[str] | A list of device string IDs. For example, ['device0', 'device12']. If empty, this field is ignored. Maximum IDs: 10,000 (optional)
    device_ids = []
    # List[str] | A list of device numeric IDs. If empty, this field is ignored. Maximum IDs: 10,000. (optional)
    device_num_ids = []
    # str | If set, returns only the gateways with which the specified device is associated. The device ID can be numeric (num_id) or the user-defined string (id). For example, if 456 is specified, returns only the gateways to which the device with num_id 456 is bound. (optional)
    gateway_list_options_associations_device_id = ''
    # str | If set, only devices associated with the specified gateway are returned. The gateway ID can be numeric (num_id) or the user-defined string (id). For example, if 123 is specified, only devices bound to the gateway with num_id 123 are returned (optional)
    gateway_list_options_associations_gateway_id = ''
    # str | If GATEWAY is specified, only gateways are returned. If NON_GATEWAY is specified, only non-gateway devices are returned. If GATEWAY_TYPE_UNSPECIFIED is specified, all devices are returned. (optional)
    gateway_list_options_gateway_type = ''

    try:
        # Get All Devices
        api_response = api_instance.get_devices(registry_id, subscription_id, page_number=page_number, page_size=page_size, field_mask=field_mask, device_ids=device_ids, device_num_ids=device_num_ids,
                                                gateway_list_options_associations_device_id=gateway_list_options_associations_device_id, gateway_list_options_associations_gateway_id=gateway_list_options_associations_gateway_id, gateway_list_options_gateway_type=gateway_list_options_gateway_type)
        print("The response of DeviceApi->get_devices:\n")
        print("The response of DeviceApi->get_device:\n")
        print("Device Id:%s" % api_response.devices[0].id)
        print("Device Full Name:%s" % api_response.devices[0].name)
        print("Certificate 1:%s" %
              api_response.devices[0].credentials[0].public_key)
        print("Log Level:%s" % api_response.devices[0].log_level)
        print("Created On:%s" % api_response.devices[0].created_on)
        print("Updated On:%s" % api_response.devices[0].updated_on)
    except Exception as e:
        print("Exception when calling DeviceApi->get_devices: %s\n" % e)
