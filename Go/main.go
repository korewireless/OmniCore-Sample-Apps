package main

var jwtToken = "<YOUR ACCESS TOKEN>"

func main() {
	CreateRegistry("sample-registry-sdk", "korewireless-development", "projects/gcp-iot-core-361019/topics/data")
	GetRegistry("korewireless-development", "sample-registry-sdk")
	GetRegistries("korewireless-development")
	UpdateRegistry("korewireless-development", "sample-registry-sdk")

	CreateDevice("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	GetDevices("korewireless-development", "sample-registry-sdk")
	GetDevice("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	UpdateDevice("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	sendCommandToDevice("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	updateConfigToDevice("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	getConfigurations("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	getStates("korewireless-development", "sample-registry-sdk", "sample-device-sdk")

	CreateGateway("korewireless-development", "sample-registry-sdk", "sample-gateway-sdk")
	GetGateways("korewireless-development", "sample-registry-sdk")
	GetGateway("korewireless-development", "sample-registry-sdk", "sample-gateway-sdk")
	UpdateGateway("korewireless-development", "sample-registry-sdk", "sample-gateway-sdk")
	bindDeviceToGateway("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	unbindDeviceFromGateway("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	bindDevicesToGateway("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	unbindDevicesFromGateway("korewireless-development", "sample-registry-sdk", "sample-device-sdk")

	DeleteDevice("korewireless-development", "sample-registry-sdk", "sample-device-sdk")

	DeleteGateway("korewireless-development", "sample-registry-sdk", "sample-gateway-sdk")

	DeleteRegistry("korewireless-development", "sample-registry-sdk")

}
