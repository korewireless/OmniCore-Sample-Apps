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
    subscription_id = 'korewireless-development'  # str | Subscription ID
    registry_id = 'shaiz-registry-sdk'  # str | Registry ID
    device_id = 'shaizgw00'  # str | Device ID
    # str | Required. Only updates the device fields indicated by this mask. The field mask must not be empty, and it must not contain fields that are immutable or only set by the server. Mutable top-level fields: credentials,logLevel, blocked, and metadata
    update_mask = 'crdentials,logLevel,blocked,metadata'
    device = OmniCore.UpdateDevice()  # UpdateDevice | application/json
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

    try:
        api_response = api_instance.update_device(
            subscription_id, registry_id, device_id, update_mask, device)
        print("The response of DeviceApi->update_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceApi->update_device: %s\n" % e)
