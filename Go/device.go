package main

import (
	"context"
	"encoding/base64"
	"fmt"
	omnicore "github.com/korewireless/OmniCore-Go-SDK"
	"log"
	"os"
)

func CreateDevice(subscriptionId string, registryId string, deviceId string) {
	configuration := omnicore.NewConfiguration()

	apiClient := omnicore.NewAPIClient(configuration)
	var format = "RSA_X509_PEM"
	var key = "-----BEGIN CERTIFICATE-----\nMIIDujCCAqKgAwIBAgITaOszYpBme+SRHZUkFWLgDs7EMDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzQy\nMloXDTMyMDgwMTEwMzIwOFowADCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAMFvTBHPdgH7+5wVlUnEdIS/0a4p9fkVzMdEMdDVr5s62VoGO7nZWxMCaxxU\nXqQiGuX3N7SINyD7h8LI8CxQsn5zyDda3QVNGU7I96iWjwzOYJmNHAN1nRI2hRDY\n8fJoQgTZI+IiRDBmgkmL9yjTY04qY7UP8zpofuMnKRuTwP6Ey1eFEMBqFfvgwrVl\niLNcq9At0bd/vlQ0VUnKV6oKqSTq9ZDPB6Cxu5amhejVwTeE6p5GGmiKw5vskmtB\ndGNgsom1K/pJdOMes8lODVp00tIVnsplL3jLgrWfbfCPALRnGz/C5XlKW8fNKEuW\nqFw2Lhnk51dtobw/oBo7vJcx2w0CAwEAAaOCAQ0wggEJMA4GA1UdDwEB/wQEAwIF\noDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQUiRS8X3OWpJNpjQoJ+22xGb3xXTAf\nBgNVHSMEGDAWgBTKO7S10CConGVgZZli7NVAcim/AzCBjQYIKwYBBQUHAQEEgYAw\nfjB8BggrBgEFBQcwAoZwaHR0cDovL3ByaXZhdGVjYS1jb250ZW50LTYyZTM5YmRh\nLTAwMDAtMjI0My1hNjFhLTNjMjg2ZDRlZWUwYS5zdG9yYWdlLmdvb2dsZWFwaXMu\nY29tL2MwZGQxZjg3ZDcwZGZhMDEwNGEwL2NhLmNydDAZBgNVHREBAf8EDzANggtn\nYWRnZW9uLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAK82b/xGn8B6Nfogw0myKjy3O\nWg53YPXuct3E04qRmD3JJtFzpSkjI2WyRmIkRX1b5SKF+ImOmGzvENZDkjT/Y2I/\nnsBL639OlXnz/+GYSq4rL6fVxXistP4LGA+khoBYSfHFZb7EYoVOYJFzZjnvJbtz\n7XG0jTMeHo8KhCBPxrNWkOERrcc7OWqREldQ36yg7zdbRLjDOjeD6FByTrpRhbDC\n0AeozF9ug9W/gPYtnkI++ksUqjJcV06uGd+9XLJPGcjH0Bai1alxROh+dkWx6TcB\nHC94el4KR6EJijMvylmnOyKHedmYaDvb52+B6zXTW9rQkh1UycmONlmAlA3OeQ==\n-----END CERTIFICATE-----"
	var gatewayType = "NON_GATEWAY"
	var gatewayAuth = "GATEWAY_AUTH_METHOD_UNSPECIFIED"
	var value = false
	device := omnicore.CreateNewDevice{
		Id:       deviceId,
		LogLevel: omnicore.LogLevel("INFO").Ptr(),
		Blocked:  &value,
		GatewayConfig: &omnicore.GatewayConfig{
			GatewayType:       &gatewayType,
			GatewayAuthMethod: &gatewayAuth,
		},
		Credentials: []omnicore.DeviceCredential{
			{
				PublicKey: &omnicore.PublicKeyCredential{
					Key:    &key,
					Format: format,
				},
			},
		},
		Metadata: &map[string]string{
			"Key":  "Value",
			"Data": "Sample Data",
		},
	}

	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response, _, err := apiClient.DeviceApi.CreateDevice(ctx, subscriptionId, registryId).Device(device).Execute()
	if err != nil {
		log.Print(err.Error())
		return
	}

	fmt.Printf("Created device:")
	fmt.Printf("\tID: %s\n", response.GetId())
	fmt.Printf("\tName: %s\n", response.GetName())
	fmt.Printf("\tGatewayConfig: %+v\n", response.GatewayConfig.GatewayType)
	fmt.Printf("\tBlocked: %t\n", response.GetBlocked())
	fmt.Printf("\tLogLevel: %s\n", response.GetLogLevel())

}

func GetDevice(subscriptionId string, registryId string, deviceId string) {
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response, r, err := apiClient.DeviceApi.GetDevice(ctx, registryId, subscriptionId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `GetDevice`: OmnicoreDevice
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevice`: \n")
	fmt.Printf("Device Details:")
	fmt.Printf("\tID: %s\n", response.GetId())
	fmt.Printf("\tName: %s\n", response.GetName())
	fmt.Printf("\tGatewayConfig: %+v\n", response.GatewayConfig.GatewayType)
	fmt.Printf("\tBlocked: %t\n", response.GetBlocked())
	fmt.Printf("\tLogLevel: %s\n", response.GetLogLevel())
}

func GetDevices(subscriptionId string, registryId string) {
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.DeviceApi.GetDevices(ctx, registryId, subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `GetDevices`: OmnicoreDevices
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevices`: \n")
	for _, device := range resp.GetDevices() {
		fmt.Printf("Devide Details:")
		fmt.Printf("\tID: %s\n", device.GetId())
		fmt.Printf("\tName: %s\n", device.GetName())
		fmt.Printf("\tGatewayConfig: %+v\n", device.GatewayConfig.GatewayType)
		fmt.Printf("\tBlocked: %t\n", device.GetBlocked())
		fmt.Printf("\tLogLevel: %s\n", device.GetLogLevel())
	}
}

func UpdateDevice(subscriptionId string, registryId string, deviceId string) {
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	updateMask := "credentials,logLevel,blocked,metadata"
	var format = "RSA_X509_PEM"
	var key = "-----BEGIN CERTIFICATE-----\nMIIDujCCAqKgAwIBAgITaOszYpBme+SRHZUkFWLgDs7EMDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzQy\nMloXDTMyMDgwMTEwMzIwOFowADCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAMFvTBHPdgH7+5wVlUnEdIS/0a4p9fkVzMdEMdDVr5s62VoGO7nZWxMCaxxU\nXqQiGuX3N7SINyD7h8LI8CxQsn5zyDda3QVNGU7I96iWjwzOYJmNHAN1nRI2hRDY\n8fJoQgTZI+IiRDBmgkmL9yjTY04qY7UP8zpofuMnKRuTwP6Ey1eFEMBqFfvgwrVl\niLNcq9At0bd/vlQ0VUnKV6oKqSTq9ZDPB6Cxu5amhejVwTeE6p5GGmiKw5vskmtB\ndGNgsom1K/pJdOMes8lODVp00tIVnsplL3jLgrWfbfCPALRnGz/C5XlKW8fNKEuW\nqFw2Lhnk51dtobw/oBo7vJcx2w0CAwEAAaOCAQ0wggEJMA4GA1UdDwEB/wQEAwIF\noDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQUiRS8X3OWpJNpjQoJ+22xGb3xXTAf\nBgNVHSMEGDAWgBTKO7S10CConGVgZZli7NVAcim/AzCBjQYIKwYBBQUHAQEEgYAw\nfjB8BggrBgEFBQcwAoZwaHR0cDovL3ByaXZhdGVjYS1jb250ZW50LTYyZTM5YmRh\nLTAwMDAtMjI0My1hNjFhLTNjMjg2ZDRlZWUwYS5zdG9yYWdlLmdvb2dsZWFwaXMu\nY29tL2MwZGQxZjg3ZDcwZGZhMDEwNGEwL2NhLmNydDAZBgNVHREBAf8EDzANggtn\nYWRnZW9uLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAK82b/xGn8B6Nfogw0myKjy3O\nWg53YPXuct3E04qRmD3JJtFzpSkjI2WyRmIkRX1b5SKF+ImOmGzvENZDkjT/Y2I/\nnsBL639OlXnz/+GYSq4rL6fVxXistP4LGA+khoBYSfHFZb7EYoVOYJFzZjnvJbtz\n7XG0jTMeHo8KhCBPxrNWkOERrcc7OWqREldQ36yg7zdbRLjDOjeD6FByTrpRhbDC\n0AeozF9ug9W/gPYtnkI++ksUqjJcV06uGd+9XLJPGcjH0Bai1alxROh+dkWx6TcB\nHC94el4KR6EJijMvylmnOyKHedmYaDvb52+B6zXTW9rQkh1UycmONlmAlA3OeQ==\n-----END CERTIFICATE-----"
	var value = false
	registry := omnicore.UpdateDevice{
		Blocked: &value,
		Metadata: &map[string]string{
			"Key":  "Value",
			"Data": "Sample Data",
		},
		LogLevel: omnicore.LogLevel("INFO").Ptr(),
		Credentials: []omnicore.DeviceCredential{
			{
				PublicKey: &omnicore.PublicKeyCredential{
					Key:    &key,
					Format: format,
				},
			},
		},
	}
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response, _, err := apiClient.DeviceApi.UpdateDevice(ctx, subscriptionId, registryId, deviceId).Device(registry).UpdateMask(updateMask).Execute()
	if err != nil {
		log.Print(err.Error())
		return
	}

	fmt.Printf("Updated device:")
	fmt.Printf("\tID: %s\n", response.GetId())
	fmt.Printf("\tNumId: %s\n", response.GetNumId())
	fmt.Printf("\tName: %s\n", response.GetName())
	fmt.Printf("\tGatewayConfig: %+v\n", response.GatewayConfig.GatewayType)
	fmt.Printf("\tBlocked: %t\n", response.GetBlocked())
	fmt.Printf("\tName: %s\n", response.GetName())
	fmt.Printf("\tLogLevel: %s\n", response.GetLogLevel())
	fmt.Printf("\tCredentials: %s\n", response.Credentials[0].PublicKey.GetKey())
	fmt.Printf("\tMetadata: %s\n", response.GetMetadata())
}

func DeleteDevice(subscriptionId string, registryId string, deviceId string) {
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.DeviceApi.DeleteDevice(ctx, subscriptionId, registryId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeleteDevice`: %v\n", resp.GetInfo())
}

func sendCommandToDevice(subscriptionId string, registryId string, deviceId string) {
	var data = base64.StdEncoding.EncodeToString([]byte("BinaryData_example"))
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	device := *omnicore.NewDeviceCommand(data) // DeviceCommand | application/json
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.DeviceApi.SendCommandToDevice(ctx, subscriptionId, registryId, deviceId).Device(device).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.SendCommandToDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r.Body)
		return
	}
	// response from `SendCommandToDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.SendCommandToDevice`: %v\n", resp)
}

func updateConfigToDevice(subscriptionId string, registryId string, deviceId string) {
	var data = base64.StdEncoding.EncodeToString([]byte("BinaryData_example"))
	device := *omnicore.NewDeviceConfiguration(data) // DeviceConfiguration | application/json

	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.DeviceApi.UpdateConfigurationToDevice(ctx, subscriptionId, registryId, deviceId).Device(device).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.SendConfigurationToDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r.Body)
		return
	}
	// response from `SendConfigurationToDevice`: DeviceConfig
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.SendConfigurationToDevice`: %v\n", resp.GetBinaryData())
}

func getConfigurations(subscriptionId string, registryId string, deviceId string) {
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	numVersions := int32(0) // int32 | Device ID

	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.DeviceApi.GetConfig(ctx, subscriptionId, registryId, deviceId).NumVersions(numVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r.Body)
		return
	}
	// response from `GetConfig`: ListDeviceConfigVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetConfig`: %v\n", resp.GetDeviceConfigs())
}

func getStates(subscriptionId string, registryId string, deviceId string) {
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	numStates := int32(0) // int32 | The number of states to list. States are listed in descending order of update time. The maximum number of states retained is 10. If this value is zero, it will return all the states available. (optional)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.DeviceApi.GetStates(ctx, subscriptionId, registryId, deviceId).NumStates(numStates).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetStates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r.Body)
		return
	}
	// response from `GetStates`: ListDeviceStatesResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetStates`: %v\n", resp.GetDeviceStates())
}