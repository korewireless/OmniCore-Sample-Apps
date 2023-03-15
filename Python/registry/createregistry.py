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
    api_instance = OmniCore.RegistryApi(api_client)
    subscription_id = 'Insert Subscription Here'  # str | Subscription ID
    registry = OmniCore.DeviceRegistry(
        id="Insert Registry Here",

    )
    registry.event_notification_configs = [
        OmniCore.EventNotificationConfig(
            pubsubTopicName="projects/gcp project id /topics/pubsubtopic name",
            subfolderMatches=""
        )
    ]

    registry.state_notification_config = OmniCore.NotificationConfig(
        pubsubTopicName="projects/gcp project id /topics/pubsubtopic name"
    )
    registry.credentials = [
        OmniCore.RegistryCredential(
            publicKeyCertificate=OmniCore.PublicKeyCertificate(
                format="X509_CERTIFICATE_PEM",
                certificate="-----BEGIN CERTIFICATE-----\nMIIDLDCCAhSgAwIBAgITMNZHEBss/CrpIrK9eX8aMQtGNDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzIw\nOVoXDTMyMDgwMTEwMzIwOFowHjENMAsGA1UEChMEa29yZTENMAsGA1UEAxMEa29y\nZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL69ASS6Fz8EXBI1yPqm\nKASYIpyOVZN5VcMnV1spqn7uLtVHVnyeLuQon+aHTqb59TjafP+d7az8hyByL2Zm\n98juwo9vCnT4rKceskZM55olnfF07dbd+DpIT/3hbDh2RlU+rSRYwYoXZ3DOF7jd\nLsEJN9tFpExInAk5GK0ZaPXiZzCY+a9hjg+6RSPPqBhoTWuI6zPrPYeWSTKHvqNG\ngWiytwriss3hBp34Gdg4sO8COD+0uf9/Ia0hB/tpcr0Myyshgzw/SqCDK4mWVqNs\nLcsBeqZkPwNvbC3ZFvUI8NJMpeAyghW25qvswyQZWxIsvJgMI9SmSXgYCBf0DyFV\nGH8CAwEAAaNjMGEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYD\nVR0OBBYEFMo7tLXQIKicZWBlmWLs1UByKb8DMB8GA1UdIwQYMBaAFMo7tLXQIKic\nZWBlmWLs1UByKb8DMA0GCSqGSIb3DQEBCwUAA4IBAQCS7aFBnvAZ66HEC5ssutgC\nJ0Ak0Li7x8YThz0+AqWyU42/8woitwuxmyQOYP4g9CJty6bqM1LoKc3bOvb1GOMm\nZ0xhAe/+H3ZIKs6g5zXon7mZOEpGJSZQLMuyPI5searFIp+3mgIo4UpgcjYFrBTE\nYJezh93GxirAFVUQ2ZcOltvt13LiavjATlSNNwhSNKZ7lUzD2Y9d5VBkPIYBZw4U\neeJ4GnDPg3IX62rJWWpJ/unJQcmxTwjY8CT85P8C0oROjqCJc93dm8aNXpw7afVq\n+0R+VZwrzkHCmfJEV6ogrR0fUQODUWSvLmR8Z956s/OijpSqRmr88mzq1UZtbnwA\n-----END CERTIFICATE-----"
            )
        )
    ]
    registry.mqtt_config = OmniCore.MqttConfig(
        mqttEnabledState="MQTT_ENABLED"
    )
    registry.log_level = "INFO"
    try:
        # Add New Registry
        api_response = api_instance.create_registry(subscription_id, registry)
        print("The response of RegistryApi->create_registry:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RegistryApi->create_registry: %s\n" % e)
