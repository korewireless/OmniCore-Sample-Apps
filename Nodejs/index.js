const dotenv = require('dotenv');
dotenv.config();

const deviceApi = require('./device/device')
const gatewayApi = require('./gateway/gateway')
const registryApi = require('./registry/registry')

token = process.env.FIREBASE_TOKEN ? process.env.FIREBASE_TOKEN:"Token example"
subscriptionId = process.env.SUBSCRIPTION_ID ? process.env.SUBSCRIPTION_ID:"Subscription_id example"
registryId = process.env.REGISTRY_ID ? process.env.REGISTRY_ID:"registry_id example"
deviceID = process.env.DEVICE_ID ? process.env.DEVICE_ID:"deviceId_example"
gatewayId = process.env.GATEWAY_ID ? process.env.GATEWAY_ID:"gatewayId_example"
newRegistryId = process.env.NEW_REGISTRY_ID ? process.env.NEW_REGISTRY_ID: "registryId_example"
hostUrl = process.env.HOST_URL ? process.env.HOST_URL: "hosturl_example"

// Device Api calls
deviceApi.createDevice(token,subscriptionId,registryId,deviceID,hostUrl)
deviceApi.getDevice(token,subscriptionId,registryId,deviceID,hostUrl);
deviceApi.updateDevice(token,subscriptionId,registryId,deviceID,hostUrl)
deviceApi.getDevices(token,subscriptionId,registryId,deviceID,hostUrl)
deviceApi.sendCommandToDevice(token,subscriptionId,registryId,deviceID,hostUrl)
deviceApi.updateConfigurationToDevice(token,subscriptionId,registryId,deviceID,hostUrl)
deviceApi.getState(token,subscriptionId,registryId,deviceID,hostUrl)
deviceApi.getConfig(token,subscriptionId,registryId,deviceID,hostUrl)
deviceApi.blockDeviceCommunication(token,subscriptionId,registryId,deviceID,hostUrl)
deviceApi.deleteDevice(token,subscriptionId,registryId,deviceID,hostUrl)

// Gateway Api calls
gatewayApi.bindDevice(token,subscriptionId,registryId,deviceID,gatewayId,hostUrl)
gatewayApi.unbindDevice(token,subscriptionId,registryId,deviceID,gatewayId,hostUrl)
gatewayApi.bindDevices(token,subscriptionId,registryId,deviceID,gatewayId,hostUrl)
gatewayApi.unbindDevices(token,subscriptionId,registryId,deviceID,gatewayId,hostUrl)

// Registry Api calls
registryApi.createRegistry(token,subscriptionId,newRegistryId,hostUrl)
registryApi.getRegistry(token,subscriptionId,newRegistryId,hostUrl)
registryApi.getRegistries(token,subscriptionId,newRegistryId,hostUrl)
registryApi.updateRegistry(token,subscriptionId,newRegistryId,hostUrl)
registryApi.deleteRegistry(token,subscriptionId,newRegistryId,hostUrl) 