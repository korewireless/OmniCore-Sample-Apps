package main

import (
	"context"
	"fmt"
	omnicore "github.com/korewireless/OmniCore-Go-SDK"
	"log"
	"os"
)

func CreateRegistry(registryId string, subscriptionId string, topicName string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	var certificate = "-----BEGIN CERTIFICATE-----\nMIIDLDCCAhSgAwIBAgITMNZHEBss/CrpIrK9eX8aMQtGNDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzIw\nOVoXDTMyMDgwMTEwMzIwOFowHjENMAsGA1UEChMEa29yZTENMAsGA1UEAxMEa29y\nZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL69ASS6Fz8EXBI1yPqm\nKASYIpyOVZN5VcMnV1spqn7uLtVHVnyeLuQon+aHTqb59TjafP+d7az8hyByL2Zm\n98juwo9vCnT4rKceskZM55olnfF07dbd+DpIT/3hbDh2RlU+rSRYwYoXZ3DOF7jd\nLsEJN9tFpExInAk5GK0ZaPXiZzCY+a9hjg+6RSPPqBhoTWuI6zPrPYeWSTKHvqNG\ngWiytwriss3hBp34Gdg4sO8COD+0uf9/Ia0hB/tpcr0Myyshgzw/SqCDK4mWVqNs\nLcsBeqZkPwNvbC3ZFvUI8NJMpeAyghW25qvswyQZWxIsvJgMI9SmSXgYCBf0DyFV\nGH8CAwEAAaNjMGEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYD\nVR0OBBYEFMo7tLXQIKicZWBlmWLs1UByKb8DMB8GA1UdIwQYMBaAFMo7tLXQIKic\nZWBlmWLs1UByKb8DMA0GCSqGSIb3DQEBCwUAA4IBAQCS7aFBnvAZ66HEC5ssutgC\nJ0Ak0Li7x8YThz0+AqWyU42/8woitwuxmyQOYP4g9CJty6bqM1LoKc3bOvb1GOMm\nZ0xhAe/+H3ZIKs6g5zXon7mZOEpGJSZQLMuyPI5searFIp+3mgIo4UpgcjYFrBTE\nYJezh93GxirAFVUQ2ZcOltvt13LiavjATlSNNwhSNKZ7lUzD2Y9d5VBkPIYBZw4U\neeJ4GnDPg3IX62rJWWpJ/unJQcmxTwjY8CT85P8C0oROjqCJc93dm8aNXpw7afVq\n+0R+VZwrzkHCmfJEV6ogrR0fUQODUWSvLmR8Z956s/OijpSqRmr88mzq1UZtbnwA\n-----END CERTIFICATE-----"
	var format = "X509_CERTIFICATE_PEM"
	var MqttConfig = "MQTT_DISABLED"
	var HttpConfig = "HTTP_DISABLED"
	registry := omnicore.DeviceRegistry{
		Id: registryId,
		EventNotificationConfigs: []omnicore.EventNotificationConfig{
			{
				PubsubTopicName: &topicName,
			},
		},
		StateNotificationConfig: &omnicore.NotificationConfig{
			PubsubTopicName: &topicName,
		},
		LogNotificationConfig: &omnicore.NotificationConfig{
			PubsubTopicName: &topicName,
		},
		MqttConfig: &omnicore.MqttConfig{
			MqttEnabledState: &MqttConfig,
		},
		HttpConfig: &omnicore.HttpConfig{
			HttpEnabledState: &HttpConfig,
		},
		LogLevel: omnicore.LogLevel("INFO").Ptr(),
		Credentials: []omnicore.RegistryCredential{
			{
				PublicKeyCertificate: &omnicore.PublicKeyCertificate{
					Certificate: &certificate,
					Format:      &format,
				},
			},
		},
	}

	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response, resp, err := apiClient.RegistryApi.CreateRegistry(ctx, subscriptionId).Registry(registry).Execute()
	if err != nil {
		log.Printf("Error %+v %+v", resp, err.Error())
		return
	}

	fmt.Printf("Created registry:")
	fmt.Printf("\tID: %s\n", response.GetId())
	fmt.Printf("\tName: %s\n", response.GetName())
	fmt.Printf("\tHTTP: %s\n", response.HttpConfig.GetHttpEnabledState())
	fmt.Printf("\tMQTT: %s\n", response.MqttConfig.GetMqttEnabledState())
	fmt.Printf("\tLogLevel: %s\n", response.GetLogLevel())
	fmt.Printf("\tEventNotificationConfigs: %+v\n", registry.EventNotificationConfigs[0].GetPubsubTopicName())
	fmt.Printf("\tStateNotificationConfig: %+v\n", registry.StateNotificationConfig.GetPubsubTopicName())
	fmt.Printf("\tLogNotificationConfig: %+v\n", registry.LogNotificationConfig.GetPubsubTopicName())

}

func GetRegistry(subscriptionId string, registryId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response, r, err := apiClient.RegistryApi.GetRegistry(ctx, subscriptionId, registryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r.Body)
		return
	}
	// response from `GetRegistry`: OmnicoreDeviceRegistry
	fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetRegistry`:")
	fmt.Printf("\tID: %s\n", response.GetId())
	fmt.Printf("\tName: %s\n", response.GetName())
	fmt.Printf("\tHTTP: %s\n", response.HttpConfig.GetHttpEnabledState())
	fmt.Printf("\tMQTT: %s\n", response.MqttConfig.GetMqttEnabledState())
	fmt.Printf("\tLogLevel: %s\n", response.GetLogLevel())
	fmt.Printf("\tEventNotificationConfigs: %+v\n", response.EventNotificationConfigs[0].GetPubsubTopicName())
	fmt.Printf("\tStateNotificationConfig: %+v\n", response.StateNotificationConfig.GetPubsubTopicName())
	fmt.Printf("\tLogNotificationConfig: %+v\n", response.LogNotificationConfig.GetPubsubTopicName())
}

func GetRegistries(subscriptionId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.RegistryApi.GetRegistries(ctx, subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetRegistries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `GetRegistries`: OmnicoreDeviceRegistries
	fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetRegistries`: \n")
	for _, registry := range resp.DeviceRegistries {
		fmt.Printf("\n\tRegistry Details:\n")
		fmt.Printf("\tID: %s\n", registry.GetId())
		fmt.Printf("\tName: %s\n", registry.GetName())
		fmt.Printf("\tLogLevel: %s\n", registry.GetLogLevel())
		fmt.Printf("\tHTTP: %s\n", registry.HttpConfig.GetHttpEnabledState())
		fmt.Printf("\tMQTT: %s\n", registry.MqttConfig.GetMqttEnabledState())
		fmt.Printf("\tStateNotificationConfig: %+v\n", registry.StateNotificationConfig.GetPubsubTopicName())
		fmt.Printf("\tLogNotificationConfig: %+v\n", registry.LogNotificationConfig.GetPubsubTopicName())
	}
}

func UpdateRegistry(subscriptionId string, registryId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	updateMask := "eventNotificationConfigs,stateNotificationConfig.pubsub_topic_name,logNotificationConfig.pubsub_topic_name,mqttConfig.mqtt_enabled_state,httpConfig.http_enabled_state,logLevel,credentials"
	var eventTopic = "projects/gcp-iot-core-361019/topics/data"
	var stateTopic = "projects/gcp-iot-core-361019/topics/control"
	var logTopic = "projects/gcp-iot-core-361019/topics/control"
	var certificate = "-----BEGIN CERTIFICATE-----\nMIIDLDCCAhSgAwIBAgITMNZHEBss/CrpIrK9eX8aMQtGNDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzIw\nOVoXDTMyMDgwMTEwMzIwOFowHjENMAsGA1UEChMEa29yZTENMAsGA1UEAxMEa29y\nZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL69ASS6Fz8EXBI1yPqm\nKASYIpyOVZN5VcMnV1spqn7uLtVHVnyeLuQon+aHTqb59TjafP+d7az8hyByL2Zm\n98juwo9vCnT4rKceskZM55olnfF07dbd+DpIT/3hbDh2RlU+rSRYwYoXZ3DOF7jd\nLsEJN9tFpExInAk5GK0ZaPXiZzCY+a9hjg+6RSPPqBhoTWuI6zPrPYeWSTKHvqNG\ngWiytwriss3hBp34Gdg4sO8COD+0uf9/Ia0hB/tpcr0Myyshgzw/SqCDK4mWVqNs\nLcsBeqZkPwNvbC3ZFvUI8NJMpeAyghW25qvswyQZWxIsvJgMI9SmSXgYCBf0DyFV\nGH8CAwEAAaNjMGEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYD\nVR0OBBYEFMo7tLXQIKicZWBlmWLs1UByKb8DMB8GA1UdIwQYMBaAFMo7tLXQIKic\nZWBlmWLs1UByKb8DMA0GCSqGSIb3DQEBCwUAA4IBAQCS7aFBnvAZ66HEC5ssutgC\nJ0Ak0Li7x8YThz0+AqWyU42/8woitwuxmyQOYP4g9CJty6bqM1LoKc3bOvb1GOMm\nZ0xhAe/+H3ZIKs6g5zXon7mZOEpGJSZQLMuyPI5searFIp+3mgIo4UpgcjYFrBTE\nYJezh93GxirAFVUQ2ZcOltvt13LiavjATlSNNwhSNKZ7lUzD2Y9d5VBkPIYBZw4U\neeJ4GnDPg3IX62rJWWpJ/unJQcmxTwjY8CT85P8C0oROjqCJc93dm8aNXpw7afVq\n+0R+VZwrzkHCmfJEV6ogrR0fUQODUWSvLmR8Z956s/OijpSqRmr88mzq1UZtbnwA\n-----END CERTIFICATE-----"
	var format = "X509_CERTIFICATE_PEM"
	var MqttConfig = "MQTT_ENABLED"
	var HttpConfig = "HTTP_ENABLED"
	registry := omnicore.DeviceRegistry{
		EventNotificationConfigs: []omnicore.EventNotificationConfig{
			{
				PubsubTopicName: &eventTopic,
			},
		},
		StateNotificationConfig: &omnicore.NotificationConfig{
			PubsubTopicName: &stateTopic,
		},
		LogNotificationConfig: &omnicore.NotificationConfig{
			PubsubTopicName: &logTopic,
		},
		MqttConfig: &omnicore.MqttConfig{
			MqttEnabledState: &MqttConfig,
		},
		HttpConfig: &omnicore.HttpConfig{
			HttpEnabledState: &HttpConfig,
		},
		LogLevel: omnicore.LogLevel("INFO").Ptr(),
		Credentials: []omnicore.RegistryCredential{
			{
				PublicKeyCertificate: &omnicore.PublicKeyCertificate{
					Certificate: &certificate,
					Format:      &format,
				},
			},
		},
	}
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response, _, err := apiClient.RegistryApi.UpdateRegistry(ctx, subscriptionId, registryId).Registry(registry).UpdateMask(updateMask).Execute()
	if err != nil {
		log.Print(err.Error())
		return
	}

	fmt.Printf("Updated registry:")
	fmt.Printf("\tID: %s\n", response.GetId())
	fmt.Printf("\tName: %s\n", response.GetName())
	fmt.Printf("\tHTTP: %s\n", response.HttpConfig.GetHttpEnabledState())
	fmt.Printf("\tMQTT: %s\n", response.MqttConfig.GetMqttEnabledState())
	fmt.Printf("\tLogLevel: %s\n", registry.GetLogLevel())
	fmt.Printf("\tEventNotificationConfigs: %+v\n", registry.EventNotificationConfigs[0].GetPubsubTopicName())
	fmt.Printf("\tStateNotificationConfig: %+v\n", registry.StateNotificationConfig.GetPubsubTopicName())
	fmt.Printf("\tLogNotificationConfig: %+v\n", registry.LogNotificationConfig.GetPubsubTopicName())

}

func DeleteRegistry(subscriptionId string, registryId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.RegistryApi.DeleteRegistry(ctx, subscriptionId, registryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.DeleteRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `GetRegistry`: OmnicoreDeviceRegistry
	fmt.Fprintf(os.Stdout, "Response from `RegistryApi.DeleteRegistry`: %v\n", resp.GetInfo())
}
