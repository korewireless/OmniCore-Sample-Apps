package main

import "fmt"

var apiKey = "<YOUR API KEY>"
var jwtToken string
var err error

func main() {
	jwtToken, err = FetchToken(jwtToken, "<Client-Id>", "<Client-Secret>")
	if err != nil {
		fmt.Print("Invalid client id or secret provided")
		return
	}
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
	SendBroadcastToDevices("korewireless-development", "sample-registry-sdk")
	CustomOnBoardDevice("korewireless-development", "sample-registry-sdk", "sample-custom-device-sdk")
	CreateGateway("korewireless-development", "sample-registry-sdk", "sample-gateway-sdk")
	GetGateways("korewireless-development", "sample-registry-sdk")
	GetGateway("korewireless-development", "sample-registry-sdk", "sample-gateway-sdk")
	UpdateGateway("korewireless-development", "sample-registry-sdk", "sample-gateway-sdk")
	bindDeviceToGateway("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	unbindDeviceFromGateway("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	bindDevicesToGateway("korewireless-development", "sample-registry-sdk", "sample-device-sdk")
	unbindDevicesFromGateway("korewireless-development", "sample-registry-sdk", "sample-device-sdk")

	DeleteDevice("korewireless-development", "sample-registry-sdk", "sample-custom-device-sdk")
	DeleteDevice("korewireless-development", "sample-registry-sdk", "sample-device-sdk")

	DeleteGateway("korewireless-development", "sample-registry-sdk", "sample-gateway-sdk")

	GetVaultAudit("korewireless-development")
	GetVaultStatus("korewireless-development")
	GetVaultMetrics("korewireless-development","1700028095723","1700632895723")
	GetFolders("korewireless-development")
	GetVaultFiles("korewireless-development","sample-registry-sdk")
	CreateVaultConfigurations("korewireless-development","gcs","{\"bucket\":\"the-vault-korewireless-development-1337\"}")
	configId := GetVaultConfigurations("korewireless-development")
	DeleteVaultConfiguration("korewireless-development",configId)
	GetReplays("korewireless-development")
	StartReplay("korewireless-development","sample-registry-sdk",1700159400,1700591400,"projects/gcp-iot-core-361019/topics/dumy-sub","sample-registry-sdk")
	GetExports("korewireless-development")
	StartExport("korewireless-development","the-vault-korewireless-development-1337","sample-registry-sdk")
    CreateVaultKey("korewireless-development","sdk-test-key")
	var keyId = GetVaultKeys("korewireless-development")
	DeleteVaulKey("korewireless-development",keyId)

	DeleteRegistry("korewireless-development", "sample-registry-sdk")

}
