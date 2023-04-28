import jwt
import requests
from datetime import datetime

# This token generation is relevant to only Multitenant SaaS mode deployment.


def generateTokenHttp(clientId, clientSecret, url='https://api.korewireless.com/Api/token'):
    if clientId == "" and clientSecret == "":
        raise Exception("Client Id,Secret Invalid")
    # The token generated is by default set to 10 hours. To control the lifetime of token add "access_token_lifespan=seconds" in the token request body.
    data = "grant_type=client_credentials&client_id={clientId}&client_secret={clientSecret}".format(
        clientId=clientId, clientSecret=clientSecret)
    headers = {"content-type": "application/x-www-form-urlencoded",
               "cachecontrol": "no-cache", }
    response = requests.post(url=url, data=data, headers=headers)
    if response.status_code >= 200 and response.status_code < 300:
        return response.json()["access_token"]
    else:
        raise Exception("Error In Fetching Token"+response.text)


def isTokenExpired(access_token):
    decodedToken = jwt.decode(access_token, verify=False, algorithms=['RS256'])
    expTimestamp = decodedToken.get('exp')
    if not expTimestamp:
        return True
    # Token is invalid if it doesn't contain an expiration time
    expDatetime = datetime.utcfromtimestamp(expTimestamp)
    nowDatetime = datetime.utcnow()
    return nowDatetime >= expDatetime


def fetchToken(currentToken, clientId, clientSecret, url='https://api.korewireless.com/Api/token'):
    if currentToken != "":
        if isTokenExpired(currentToken):
            return generateTokenHttp(clientId, clientSecret, url)
        else:
            return currentToken
    else:
        return generateTokenHttp(clientId, clientSecret, url)
