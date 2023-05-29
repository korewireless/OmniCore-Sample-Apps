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
    subscription_id = 'Insert Subscription Here'  # str | Subscription ID
    registry_id = 'Insert Registry Here'  # str | Registry ID
    device_id = 'shaizdev00'  # str | Device ID
    # str | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,logLevel, blocked,policy and metadata
    update_mask = 'credentials,logLevel,blocked,metadata,policy'
    device = OmniCore.Device(id=device_id)  # UpdateDevice | application/json
    device.blocked = False
    device.credentials = []
    device.credentials.append(OmniCore.DeviceCredential(
        publicKey=OmniCore.PublicKeyCredential(
            format="RSA_X509_PEM",
            key="-----BEGIN CERTIFICATE-----\nMIIDujCCAqKgAwIBAgITaOszYpBme+SRHZUkFWLgDs7EMDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzQy\nMloXDTMyMDgwMTEwMzIwOFowADCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAMFvTBHPdgH7+5wVlUnEdIS/0a4p9fkVzMdEMdDVr5s62VoGO7nZWxMCaxxU\nXqQiGuX3N7SINyD7h8LI8CxQsn5zyDda3QVNGU7I96iWjwzOYJmNHAN1nRI2hRDY\n8fJoQgTZI+IiRDBmgkmL9yjTY04qY7UP8zpofuMnKRuTwP6Ey1eFEMBqFfvgwrVl\niLNcq9At0bd/vlQ0VUnKV6oKqSTq9ZDPB6Cxu5amhejVwTeE6p5GGmiKw5vskmtB\ndGNgsom1K/pJdOMes8lODVp00tIVnsplL3jLgrWfbfCPALRnGz/C5XlKW8fNKEuW\nqFw2Lhnk51dtobw/oBo7vJcx2w0CAwEAAaOCAQ0wggEJMA4GA1UdDwEB/wQEAwIF\noDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQUiRS8X3OWpJNpjQoJ+22xGb3xXTAf\nBgNVHSMEGDAWgBTKO7S10CConGVgZZli7NVAcim/AzCBjQYIKwYBBQUHAQEEgYAw\nfjB8BggrBgEFBQcwAoZwaHR0cDovL3ByaXZhdGVjYS1jb250ZW50LTYyZTM5YmRh\nLTAwMDAtMjI0My1hNjFhLTNjMjg2ZDRlZWUwYS5zdG9yYWdlLmdvb2dsZWFwaXMu\nY29tL2MwZGQxZjg3ZDcwZGZhMDEwNGEwL2NhLmNydDAZBgNVHREBAf8EDzANggtn\nYWRnZW9uLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAK82b/xGn8B6Nfogw0myKjy3O\nWg53YPXuct3E04qRmD3JJtFzpSkjI2WyRmIkRX1b5SKF+ImOmGzvENZDkjT/Y2I/\nnsBL639OlXnz/+GYSq4rL6fVxXistP4LGA+khoBYSfHFZb7EYoVOYJFzZjnvJbtz\n7XG0jTMeHo8KhCBPxrNWkOERrcc7OWqREldQ36yg7zdbRLjDOjeD6FByTrpRhbDC\n0AeozF9ug9W/gPYtnkI++ksUqjJcV06uGd+9XLJPGcjH0Bai1alxROh+dkWx6TcB\nHC94el4KR6EJijMvylmnOyKHedmYaDvb52+B6zXTW9rQkh1UycmONlmAlA3OeQ==\n-----END CERTIFICATE-----"
        )))
    device.log_level = "ERROR"
    device.metadata = {
        "floor": "4"
    }
    device.policy = OmniCore.Policy(Connect=True, PublishEvents=True, PublishLoopback=True, PublishState=True, SubscribeBroadcast=True,
                                    SubscribeConfig=True, SubscribeCommand=True, PublishEventsRegex=".*", SubscribeCommandRegex=".*", SubscribeBroadcastRegex=".*")

    try:
        api_response = api_instance.update_device(
            subscription_id, registry_id, device_id, update_mask, device)
        print("The response of DeviceApi->update_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->update_device: %s\n" % e)
