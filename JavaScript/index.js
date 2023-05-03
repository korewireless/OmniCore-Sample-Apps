const dotenv = require('dotenv');
dotenv.config();

const TokenHelper = require('./token-helper');

const deviceApi = require('./device/device')
const gatewayApi = require('./gateway/gateway')
const registryApi = require('./registry/registry')

tokenUrl = process.env.TOKEN_URL ? process.env.TOKEN_URL : "https://api.korewireless.com/Api/token"
clientId = process.env.CLIENT_ID ? process.env.CLIENT_ID : "your-api-client-id"
clientSecret = process.env.CLIENT_SECRET ? process.env.CLIENT_SECRET : "your-api-client-secret"

subscriptionId = process.env.SUBSCRIPTION_ID ? process.env.SUBSCRIPTION_ID : "Subscription_id example"
registryId = process.env.REGISTRY_ID ? process.env.REGISTRY_ID : "registry_id example"
deviceID = process.env.DEVICE_ID ? process.env.DEVICE_ID : "deviceId_example"
gatewayId = process.env.GATEWAY_ID ? process.env.GATEWAY_ID : "gatewayId_example"
newRegistryId = process.env.NEW_REGISTRY_ID ? process.env.NEW_REGISTRY_ID : "registryId_example"
hostUrl = process.env.HOST_URL ? process.env.HOST_URL : "https://api.korewireless.com/omnicore"
apikey = process.env.API_KEY ? process.env.API_KEY : "apikey_example"

//Get token
let token = '';
token = TokenHelper.fetchToken(token, clientId, clientSecret, tokenUrl)

// Device Api calls
deviceApi.createDevice(token, subscriptionId, registryId, deviceID, hostUrl,apikey)
deviceApi.getDevice(token, subscriptionId, registryId, deviceID, hostUrl,apikey);
deviceApi.updateDevice(token, subscriptionId, registryId, deviceID, hostUrl,apikey)
deviceApi.getDevices(token, subscriptionId, registryId, deviceID, hostUrl,apikey)
deviceApi.sendCommandToDevice(token, subscriptionId, registryId, deviceID, hostUrl,apikey)
deviceApi.updateConfigurationToDevice(token, subscriptionId, registryId, deviceID, hostUrl,apikey)
deviceApi.getState(token, subscriptionId, registryId, deviceID, hostUrl,apikey)
deviceApi.getConfig(token, subscriptionId, registryId, deviceID, hostUrl,apikey)
deviceApi.blockDeviceCommunication(token, subscriptionId, registryId, deviceID, hostUrl,apikey)
deviceApi.deleteDevice(token, subscriptionId, registryId, deviceID, hostUrl,apikey)

// Gateway Api calls
gatewayApi.bindDevice(token, subscriptionId, registryId, deviceID, gatewayId, hostUrl, apikey)
gatewayApi.unbindDevice(token, subscriptionId, registryId, deviceID, gatewayId, hostUrl, apikey)
gatewayApi.bindDevices(token, subscriptionId, registryId, deviceID, gatewayId, hostUrl, apikey)
gatewayApi.unbindDevices(token, subscriptionId, registryId, deviceID, gatewayId, hostUrl, apikey)

// Registry Api calls
registryApi.createRegistry(token, subscriptionId, newRegistryId, hostUrl, apikey)
registryApi.getRegistry(token, subscriptionId, newRegistryId, hostUrl, apikey)
registryApi.getRegistries(token, subscriptionId, newRegistryId, hostUrl, apikey)
registryApi.updateRegistry(token, subscriptionId, newRegistryId, hostUrl, apikey)
registryApi.deleteRegistry(token, subscriptionId, newRegistryId, hostUrl, apikey)
registryApi.sendBroadcast(token, subscriptionId, newRegistryId, hostUrl, apikey)