package main

import (
	"context"
	"fmt"
	omnicore "github.com/korewireless/OmniCore-Go-SDK"
	"log"
	"os"
)

func CreateGateway(subscriptionId string, registryId string, deviceId string) {
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	var format = "RSA_X509_PEM"
	var key = "-----BEGIN CERTIFICATE-----\nMIIDujCCAqKgAwIBAgITaOszYpBme+SRHZUkFWLgDs7EMDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzQy\nMloXDTMyMDgwMTEwMzIwOFowADCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAMFvTBHPdgH7+5wVlUnEdIS/0a4p9fkVzMdEMdDVr5s62VoGO7nZWxMCaxxU\nXqQiGuX3N7SINyD7h8LI8CxQsn5zyDda3QVNGU7I96iWjwzOYJmNHAN1nRI2hRDY\n8fJoQgTZI+IiRDBmgkmL9yjTY04qY7UP8zpofuMnKRuTwP6Ey1eFEMBqFfvgwrVl\niLNcq9At0bd/vlQ0VUnKV6oKqSTq9ZDPB6Cxu5amhejVwTeE6p5GGmiKw5vskmtB\ndGNgsom1K/pJdOMes8lODVp00tIVnsplL3jLgrWfbfCPALRnGz/C5XlKW8fNKEuW\nqFw2Lhnk51dtobw/oBo7vJcx2w0CAwEAAaOCAQ0wggEJMA4GA1UdDwEB/wQEAwIF\noDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQUiRS8X3OWpJNpjQoJ+22xGb3xXTAf\nBgNVHSMEGDAWgBTKO7S10CConGVgZZli7NVAcim/AzCBjQYIKwYBBQUHAQEEgYAw\nfjB8BggrBgEFBQcwAoZwaHR0cDovL3ByaXZhdGVjYS1jb250ZW50LTYyZTM5YmRh\nLTAwMDAtMjI0My1hNjFhLTNjMjg2ZDRlZWUwYS5zdG9yYWdlLmdvb2dsZWFwaXMu\nY29tL2MwZGQxZjg3ZDcwZGZhMDEwNGEwL2NhLmNydDAZBgNVHREBAf8EDzANggtn\nYWRnZW9uLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAK82b/xGn8B6Nfogw0myKjy3O\nWg53YPXuct3E04qRmD3JJtFzpSkjI2WyRmIkRX1b5SKF+ImOmGzvENZDkjT/Y2I/\nnsBL639OlXnz/+GYSq4rL6fVxXistP4LGA+khoBYSfHFZb7EYoVOYJFzZjnvJbtz\n7XG0jTMeHo8KhCBPxrNWkOERrcc7OWqREldQ36yg7zdbRLjDOjeD6FByTrpRhbDC\n0AeozF9ug9W/gPYtnkI++ksUqjJcV06uGd+9XLJPGcjH0Bai1alxROh+dkWx6TcB\nHC94el4KR6EJijMvylmnOyKHedmYaDvb52+B6zXTW9rQkh1UycmONlmAlA3OeQ==\n-----END CERTIFICATE-----"
	var gatewayType = "GATEWAY"
	var gatewayAuth = "ASSOCIATION_ONLY"
	var value = true
	device := omnicore.Device{
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
	fmt.Printf("\tGatewayConfig: %+v\n", response.GatewayConfig.GetGatewayType())
	fmt.Printf("\tBlocked: %t\n", response.GetBlocked())
	fmt.Printf("\tLogLevel: %s\n", response.GetLogLevel())

}

func GetGateway(subscriptionId string, registryId string, deviceId string) {
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
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetGateway`: \n")
	fmt.Printf("\tID: %s\n", response.GetId())
	fmt.Printf("\tGatewayConfig: %+v\n", response.GatewayConfig.GetGatewayType())
	fmt.Printf("\tBlocked: %t\n", response.GetBlocked())
	fmt.Printf("\tName: %s\n", response.GetName())
	fmt.Printf("\tLogLevel: %s\n", response.GetLogLevel())
}

func GetGateways(subscriptionId string, registryId string) {
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.DeviceApi.GetDevices(ctx, registryId, subscriptionId).GatewayListOptionsGatewayType("GATEWAY").Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `GetGateways`
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetGateways`: \n")
	for _, devices := range resp.GetDevices() {
		fmt.Printf("\n\t Gateway Details: \t\n")
		fmt.Printf("\tID: %s\n", devices.GetId())
		fmt.Printf("\tName: %s\n", devices.GetName())
		fmt.Printf("\tGatewayConfig: %+v\n", devices.GatewayConfig.GetGatewayType())
		fmt.Printf("\tBlocked: %t\n", devices.GetBlocked())
		fmt.Printf("\tLogLevel: %s\n", devices.GetLogLevel())
	}
}

func UpdateGateway(subscriptionId string, registryId string, deviceId string) {
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	updateMask := "credentials,logLevel,blocked,metadata"
	var format = "RSA_X509_PEM"
	var key = "-----BEGIN CERTIFICATE-----\nMIIDujCCAqKgAwIBAgITaOszYpBme+SRHZUkFWLgDs7EMDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzQy\nMloXDTMyMDgwMTEwMzIwOFowADCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAMFvTBHPdgH7+5wVlUnEdIS/0a4p9fkVzMdEMdDVr5s62VoGO7nZWxMCaxxU\nXqQiGuX3N7SINyD7h8LI8CxQsn5zyDda3QVNGU7I96iWjwzOYJmNHAN1nRI2hRDY\n8fJoQgTZI+IiRDBmgkmL9yjTY04qY7UP8zpofuMnKRuTwP6Ey1eFEMBqFfvgwrVl\niLNcq9At0bd/vlQ0VUnKV6oKqSTq9ZDPB6Cxu5amhejVwTeE6p5GGmiKw5vskmtB\ndGNgsom1K/pJdOMes8lODVp00tIVnsplL3jLgrWfbfCPALRnGz/C5XlKW8fNKEuW\nqFw2Lhnk51dtobw/oBo7vJcx2w0CAwEAAaOCAQ0wggEJMA4GA1UdDwEB/wQEAwIF\noDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQUiRS8X3OWpJNpjQoJ+22xGb3xXTAf\nBgNVHSMEGDAWgBTKO7S10CConGVgZZli7NVAcim/AzCBjQYIKwYBBQUHAQEEgYAw\nfjB8BggrBgEFBQcwAoZwaHR0cDovL3ByaXZhdGVjYS1jb250ZW50LTYyZTM5YmRh\nLTAwMDAtMjI0My1hNjFhLTNjMjg2ZDRlZWUwYS5zdG9yYWdlLmdvb2dsZWFwaXMu\nY29tL2MwZGQxZjg3ZDcwZGZhMDEwNGEwL2NhLmNydDAZBgNVHREBAf8EDzANggtn\nYWRnZW9uLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAK82b/xGn8B6Nfogw0myKjy3O\nWg53YPXuct3E04qRmD3JJtFzpSkjI2WyRmIkRX1b5SKF+ImOmGzvENZDkjT/Y2I/\nnsBL639OlXnz/+GYSq4rL6fVxXistP4LGA+khoBYSfHFZb7EYoVOYJFzZjnvJbtz\n7XG0jTMeHo8KhCBPxrNWkOERrcc7OWqREldQ36yg7zdbRLjDOjeD6FByTrpRhbDC\n0AeozF9ug9W/gPYtnkI++ksUqjJcV06uGd+9XLJPGcjH0Bai1alxROh+dkWx6TcB\nHC94el4KR6EJijMvylmnOyKHedmYaDvb52+B6zXTW9rQkh1UycmONlmAlA3OeQ==\n-----END CERTIFICATE-----"
	var value = false
	registry := omnicore.Device{
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

	fmt.Printf("Updated Gateway:")
	fmt.Printf("\tID: %s\n", response.GetId())
	fmt.Printf("\tNumId: %s\n", response.GetNumId())
	fmt.Printf("\tName: %s\n", response.GetName())
	fmt.Printf("\tGatewayConfig: %+v\n", response.GatewayConfig.GetGatewayType())
	fmt.Printf("\tBlocked: %+v\n", response.GetBlocked())
	fmt.Printf("\tLogLevel: %+v\n", response.GetLogLevel())
	fmt.Printf("\tCredentials: %s\n", response.Credentials[0].PublicKey.GetKey())
	fmt.Printf("\tMetadata: %s\n", response.GetMetadata())

}

func DeleteGateway(subscriptionId string, registryId string, deviceId string) {
	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.DeviceApi.DeleteDevice(ctx, subscriptionId, registryId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `GetGateway`: OmnicoreGateway
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeleteGateway`: %v\n", resp.GetInfo())
}

func bindDeviceToGateway(subscriptionId string, registryId string, deviceId string) {
	device := *omnicore.NewBindRequest(deviceId, "sample-gateway-sdk") // BindRequest | application/json

	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.DeviceApi.BindDevice(ctx, subscriptionId, registryId).Device(device).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.BindDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `BindDevice`: Info
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.BindDevice`: %v\n", resp.GetInfo())
}

func unbindDeviceFromGateway(subscriptionId string, registryId string, deviceId string) {
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	device := *omnicore.NewBindRequest(deviceId, "sample-gateway-sdk") // BindRequest | application/json

	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceApi.UnBindDevice(ctx, subscriptionId, registryId).Device(device).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UnBindDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `UnBindDevice`: Info
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UnBindDevice`: %v\n", resp.GetInfo())
}

func bindDevicesToGateway(subscriptionId string, registryId string, deviceId string) {
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	device := *omnicore.NewBindRequestIdsGateway([]string{deviceId}, "sample-gateway-sdk") // BindRequestIdsGateway | application/json

	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceApi.BindDevices(ctx, subscriptionId, registryId).Device(device).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.BindDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `BindDevices`: Info
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.BindDevices`: %v\n", resp.GetInfo())
}

func unbindDevicesFromGateway(subscriptionId string, registryId string, deviceId string) {
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	device := *omnicore.NewBindRequestIdsGateway([]string{deviceId}, "sample-gateway-sdk") // BindRequestIdsGateway | application/json

	configuration := omnicore.NewConfiguration()
	apiClient := omnicore.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceApi.UnBindDevices(ctx, subscriptionId, registryId).Device(device).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UnBindDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `UnBindDevices`: Info
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UnBindDevices`: %v\n", resp.GetInfo())
}
