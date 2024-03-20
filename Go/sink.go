package main

import (
	"context"
	"fmt"
	omnicore "github.com/korewireless/OmniCore-Go-SDK"
	"log"
	"os"
)

func CreateSink(subscriptionId string, sinkId string) string {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	
	apiClient := omnicore.NewAPIClient(configuration)
	region := "us-east-1"
	externalId:= "<external Id>"
    sinkType:= "kinesis"

	sink := omnicore.Sink{
		Id:       &sinkId,
		Config: &omnicore.Config{
			ConnectionParameter: "<roleArn>",
			Region: &region,
			ExternalId: &externalId,
		},
		Sink: &sinkType,
	}

	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response, _, err := apiClient.SinkApi.CreateSink(ctx, subscriptionId).Sink(sink).Execute()
	if err != nil {
		log.Print(err.Error())
		return ""
	}

	fmt.Printf("Created sink:")
	fmt.Printf("\tID: %s\n", response.GetId())
	fmt.Printf("\tName: %s\n", response.GetSink())
	return response.GetId()
}

func GetSink(subscriptionId string, sinkId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response, r, err := apiClient.SinkApi.GetASink(ctx, subscriptionId, sinkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinkApi.GetASink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `GetDevice`: OmnicoreDevice
	fmt.Fprintf(os.Stdout, "Response from `SinkApi.GetASink`: \n")
	fmt.Printf("\tID: %s\n", response.GetId())
	fmt.Printf("\tSinkConfig: %+v\n", response.GetConfig())
}

func GetSinks(subscriptionId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.SinkApi.GetSinks(ctx, subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SinkApi.GetSinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}
	// response from `GetSinks`
	fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetSinks`: \n")
	for _, devices := range resp.GetSinks() {
		fmt.Printf("\n\t Sink Details: \t\n")
		fmt.Printf("\tID: %s\n", devices.GetId())
		fmt.Printf("\tName: %s\n", devices.GetSink())
	}
}

func DeleteSink(subscriptionId string, sinkId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	
	apiClient := omnicore.NewAPIClient(configuration)
	
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response, _, err := apiClient.SinkApi.DeleteSink(ctx, subscriptionId, sinkId).Execute()
	if err != nil {
		log.Print(err.Error())
		return
	}

	fmt.Printf("Deleted Sink:")
	fmt.Printf("\tID: %s\n", response)
}
