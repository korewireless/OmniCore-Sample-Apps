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
    #host="https://api.korewireless.com",
    access_token=access_token,
    api_key={"apiKey": "Insert Api Key Here"}
)


# Enter a context with an instance of the API client
with OmniCore.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = OmniCore.RegistryApi(api_client)
    subscription_id = 'Insert Subscription Here'  # str | Subscription ID
    registry_id = 'Insert Registry Here'  # str | Registry ID
    # str | values to be updated: eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials,customOnboardNotificationConfig.pubsub_topic_name,customOnboardEnabled
    update_mask = 'eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials,customOnboardNotificationConfig.pubsub_topic_name,customOnboardEnabled'
    # UpdateRegistry | application/json
    registry = OmniCore.DeviceRegistry(id=registry_id)
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
    registry.log_level = "ERROR"
    registry.custom_onboard_enabled = True
    registry.custom_onboard_notification_config = OmniCore.NotificationConfig(
        pubsubTopicName="projects/gcp project id /topics/pubsubtopic name"
    )
    try:
        api_response = api_instance.update_registry(
            subscription_id, registry_id, update_mask, registry)
        print("The response of RegistryApi->update_registry:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RegistryApi->update_registry: %s\n" % e)
