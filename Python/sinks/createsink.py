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
    api_instance = OmniCore.SinkApi(api_client)
    subscription_id = 'Insert Subscription Here'  # str | Subscription ID
    sink = OmniCore.Sink(
        subscription=subscription_id,
       

    )  
    sink.sink='pubsub' #type of sink
    sink.config=OmniCore.Config(connectionParameter='gcp-project-name')
    sink.status=True


    try:
        # Add New Sink
        api_response = api_instance.create_sink(
            subscription_id, sink)
        print("The response of SinkApi->create_sink:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SinkApi->create_sink: %s\n" % e)
