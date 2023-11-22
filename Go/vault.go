package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	omnicore "github.com/korewireless/OmniCore-Go-SDK"
)

func GetVaultAudit(subscriptionId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	pageNumber := int32(1) // int32 | Page Number (optional)
    pageSize := int32(10) // int32 | Page Size (optional)
    ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response,resp,err := apiClient.VaultApi.GetVaultAudit(ctx,subscriptionId).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetVaultAudit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return
	}
	fmt.Println("Get Vault Audit Response:")
	for _,audit := range response.Audit{
		fmt.Printf("\tID: %s\n",audit.GetSubscriptionId())
		fmt.Printf("\tActor: %s\n",audit.GetActor())
		fmt.Printf("\tOperation: %s\n",audit.GetOperation())
		fmt.Printf("\tUpdatedOn: %s\n",audit.GetUpdatedon())
	}
	fmt.Printf("\tPageNumber: %d\n",response.GetPageNumber())
	fmt.Printf("\tPageSize: %d\n",response.GetPageSize())
	fmt.Printf("\tTotalCount: %d\n",response.GetTotalCount())
}

func GetVaultStatus(subscriptionId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response,resp,err := apiClient.VaultApi.GetVaultStatus(ctx,subscriptionId).Execute()
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetVaultStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return
    }
	fmt.Println("Get Vault Status Response:")
	fmt.Printf("\tID: %s\n",response.Details.GetSubscriptionId())
	fmt.Printf("\tMqttId: %s\n", response.Details.GetMqttId())
	fmt.Printf("\tCreatedOn: %s\n", response.Details.GetCreatedOn())
	fmt.Printf("\tIsCreated: %t\n", response.Details.GetIsCreated())
	fmt.Printf("\tIsEnabled: %t\n", response.Details.GetIsEnabled())
	fmt.Printf("\tStorageType: %s\n", response.Details.GetStorageType())
	fmt.Printf("\tUpdatedon: %s\n", response.Details.GetUpdatedon())
}

func GetVaultMetrics(subscriptionId string,startValue string,endValue string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	startTime := startValue // string | start time (optional) | (Unix epoch time in milliseconds)
    endTime := endValue // string | end time (optional) | (Unix epoch time in milliseconds)
	response,resp,err := apiClient.VaultApi.GetVaultMetrics(ctx,subscriptionId).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetVaultMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return
    }
	fmt.Println("Get Vault Metrics Response:")
	fmt.Printf("\tID: %s\n",response.Details.GetSubscriptionId())
	fmt.Printf("\tFileSize: %f\n",response.Details.GetFileSize())
    fmt.Printf("\tNoOfFiles: %d\n",response.Details.GetNoOfFiles())
	fmt.Printf("\tNoOfOperations: %d\n",response.Details.GetNoofoperations())
	fmt.Printf("\tDetailsForTimePeriod:\n")
	 for _,detail := range response.Details.DetailsForTimePeriod.Data {
		fmt.Printf("\t\tFileSize: %f\n", detail.GetFileSize())
        fmt.Printf("\t\tNoOfFiles: %d\n",detail.GetNoOfFiles())
	    fmt.Printf("\t\tNoOfOperations: %d\n",detail.GetNoofoperations())
	    fmt.Printf("\t\tUpdatedOn: %s\n",detail.GetUpdatedon())
		fmt.Printf("\t\tReplayFileSize: %f\n", detail.GetReplayFileSize())
	    fmt.Printf("\t\tExportFileSize: %f\n", detail.GetExportFileSize())
	 }
	 fmt.Printf("\t\tTotalReplaySize: %f\n", response.Details.DetailsForTimePeriod.GetTotalReplaySize())
	 fmt.Printf("\t\tTotalExportSize: %f\n", response.Details.DetailsForTimePeriod.GetTotalExportSize())
	fmt.Printf("\tOperations: %d\n",response.Details.GetOperations())
	fmt.Printf("\tNoOfReplays: %d\n",response.Details.GetNoOfReplays())
	fmt.Printf("\tNoOfExports: %d\n",response.Details.GetNoOfExports())
}

func GetFolders(subscriptionId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response,resp,err := apiClient.VaultApi.GetRegistryData(ctx,subscriptionId).Execute()
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return
    }
	fmt.Println("Get Vault Folders Response:")
	for _,folder := range response.Details{
		fmt.Printf("\tFolderName: %s\n",folder.GetRegistryid())
		fmt.Printf("\tFileSize: %f\n", folder.GetFileSize())
		fmt.Printf("\tNoOfFiles: %d\n",folder.GetNoOfFiles())
		fmt.Printf("\tNoOfOperations: %d\n", folder.GetNoofoperations())
		fmt.Printf("\tUpdatedon: %s\n", folder.GetUpdatedon())
	}
}

func GetVaultFiles(subscriptionId string,registryId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response,resp,err := apiClient.VaultApi.GetVaultFiles(ctx,subscriptionId,registryId).FileType("state").Execute()
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return
    }
	fmt.Println("Get Vault Files Response:")
	for _,file := range response.Details{
	fmt.Printf("\tFilerName: %s\n",file.GetName())
	fmt.Printf("\tUrl: %s\n",file.GetUrl())
	}
}

func CreateVaultConfigurations(subscriptionId string,typeValue string, dataValue string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	createConfiguration := omnicore.CreateConfiguration{
		Type: &typeValue,
		Data: &dataValue,
	}
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.VaultApi.CreateVaultConfiguration(ctx, subscriptionId).CreateConfiguration(createConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.CreateVaultConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
    }
	fmt.Println("Create Vault Configurations Response:")
	fmt.Printf("\tDetails: %s\n", resp.GetDetails())
}

func GetVaultConfigurations(subscriptionId string) string {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	var configId string
	response,resp,err := apiClient.VaultApi.GetVaultConfigurations(ctx,subscriptionId).Execute()
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return configId
    }
	fmt.Println("Get Vault Configurations Response:")
	for _,config := range response.Details{ 
		fmt.Printf("\tID: %d\n", config.GetId())
		fmt.Printf("\tSubscription: %s\n", config.GetSubscription())
		fmt.Printf("\tType: %s\n", config.GetType())
	    fmt.Printf("\tData: %s\n", config.GetData())
		if config.Data != nil {
			if config.GetData() == "{\"bucket\":\"the-vault-korewireless-development-1337\"}"{
				id := config.GetId()
				configId = strconv.Itoa(int(id))
			}
		}
	}
	return configId
}

func DeleteVaultConfiguration(subscriptionId string, configId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.VaultApi.DeleteConfiguration(ctx, subscriptionId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.DeleteConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
    }
	fmt.Println("Delete Vault Configurations Response:")
	fmt.Printf("\tDetails: %s\n", resp.GetDetails())
}

func GetReplays(subscriptionId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.VaultApi.GetReplays(ctx, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetReplays``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
	fmt.Println("Get Replays Response:")
	for _,replay := range resp.Details {
		fmt.Printf("\tID: %d\n", replay.GetId())
		fmt.Printf("\tSubscription: %s\n", replay.GetSubscription())
		fmt.Printf("\tRegistry: %s\n", replay.GetRegistry())
	    fmt.Printf("\tStartTime: %d\n", replay.GetStartTime())
		fmt.Printf("\tStartTime: %d\n", replay.GetEndTime())
		fmt.Printf("\tDestination: %s\n", replay.GetDestination())
		fmt.Printf("\tSource: %s\n", replay.GetSource())
	    fmt.Printf("\tSize: %d\n", replay.GetSize())
		fmt.Printf("\tStatus: %s\n", replay.GetStatus())
		fmt.Printf("\tCount: %d\n", replay.GetCount())
	}
}

func StartReplay(subscriptionId string,registryId string,startTime int32,endTime int32,destination string,source string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	replayBody := omnicore.ReplayBody{
		Registry: &registryId,
		StartTime: &startTime,
		EndTime: &endTime,
		Destination: &destination,
		Source: &source,
	}
	resp, r, err := apiClient.VaultApi.StartReplay(ctx, subscriptionId).ReplayBody(replayBody).Execute()
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.StartReplay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
	fmt.Println("Start Replays Response:")
    fmt.Printf("\tDetails: %s\n", resp)
}

func GetExports(subscriptionId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	resp, r, err := apiClient.VaultApi.GetExports(ctx, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
	fmt.Println("Get Export Response:")
	for _,export := range resp.Details {
		fmt.Printf("\tName: %s\n", export.GetName())
		fmt.Printf("\tSubscription: %s\n", export.GetSubscription())
		fmt.Printf("\tNooffiles: %d\n", export.GetNooffiles())
	    fmt.Printf("\tCreatedon: %s\n", export.GetCreatedOn())
		fmt.Printf("\tFileSize: %f\n", export.GetFileSize())
		fmt.Printf("\tDestination: %s\n", export.GetDestination())
		fmt.Printf("\tSource: %s\n", export.GetSource())
		fmt.Printf("\tStatus: %s\n", export.GetStatus())
		fmt.Printf("\tDone: %t\n", export.GetDone())
	}
}

func StartExport(subscriptionId string, destination string, source string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	exportBody := omnicore.StartExportBody{
		Destination: &destination,
		Source: &source,
	}
	resp, r, err := apiClient.VaultApi.StartExport(ctx, subscriptionId).StartExportBody(exportBody).Execute()
	if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.StartReplay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
	fmt.Println("Start Exports Response:")
    fmt.Printf("\tDetails: %s\n", resp.GetDetails())
}

func GetVaultKeys(subscriptionId string) string {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response,resp,err := apiClient.VaultApi.GetVaultKeys(ctx,subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.GetVaultKeys`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return ""
	}
	fmt.Println("Get Vault Keys Response:")
	var keyId = ""
	for _,key := range response.Details {
		fmt.Printf("\tID: %d\n", key.GetId())
		fmt.Printf("\tName: %s\n", key.GetName())
		fmt.Printf("\tKey: %s\n", key.GetKey())
		fmt.Printf("\tSubscription: %s\n", key.GetSubscription())
		if key.Name != nil && *key.Name == "sdk-test-key" {
			id := key.GetId()
			keyId = strconv.Itoa(int(id))
		}
	}
	return keyId
}

func CreateVaultKey(subscriptionId string,name string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	key := "-----BEGIN CERTIFICATE-----\nMIIDLDCCAhSgAwIBAgITMNZHEBss/CrpIrK9eX8aMQtGNDANBgkqhkiG9w0BAQsF\nADAeMQ0wCwYDVQQKEwRrb3JlMQ0wCwYDVQQDEwRrb3JlMB4XDTIyMDgwNDEwMzIw\nOVoXDTMyMDgwMTEwMzIwOFowHjENMAsGA1UEChMEa29yZTENMAsGA1UEAxMEa29y\nZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL69ASS6Fz8EXBI1yPqm\nKASYIpyOVZN5VcMnV1spqn7uLtVHVnyeLuQon+aHTqb59TjafP+d7az8hyByL2Zm\n98juwo9vCnT4rKceskZM55olnfF07dbd+DpIT/3hbDh2RlU+rSRYwYoXZ3DOF7jd\nLsEJN9tFpExInAk5GK0ZaPXiZzCY+a9hjg+6RSPPqBhoTWuI6zPrPYeWSTKHvqNG\ngWiytwriss3hBp34Gdg4sO8COD+0uf9/Ia0hB/tpcr0Myyshgzw/SqCDK4mWVqNs\nLcsBeqZkPwNvbC3ZFvUI8NJMpeAyghW25qvswyQZWxIsvJgMI9SmSXgYCBf0DyFV\nGH8CAwEAAaNjMGEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYD\nVR0OBBYEFMo7tLXQIKicZWBlmWLs1UByKb8DMB8GA1UdIwQYMBaAFMo7tLXQIKic\nZWBlmWLs1UByKb8DMA0GCSqGSIb3DQEBCwUAA4IBAQCS7aFBnvAZ66HEC5ssutgC\nJ0Ak0Li7x8YThz0+AqWyU42/8woitwuxmyQOYP4g9CJty6bqM1LoKc3bOvb1GOMm\nZ0xhAe/+H3ZIKs6g5zXon7mZOEpGJSZQLMuyPI5searFIp+3mgIo4UpgcjYFrBTE\nYJezh93GxirAFVUQ2ZcOltvt13LiavjATlSNNwhSNKZ7lUzD2Y9d5VBkPIYBZw4U\neeJ4GnDPg3IX62rJWWpJ/unJQcmxTwjY8CT85P8C0oROjqCJc93dm8aNXpw7afVq\n+0R+VZwrzkHCmfJEV6ogrR0fUQODUWSvLmR8Z956s/OijpSqRmr88mzq1UZtbnwA\n-----END CERTIFICATE-----"
	keyBody := omnicore.CreateVaultKeyBody{
		Name: &name,
		Key: &key,
	}
	response,resp,err := apiClient.VaultApi.CreateVaultKey(ctx,subscriptionId).CreateVaultKeyBody(keyBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.CreateVaultKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return
	}
	fmt.Println("Create Vault Key Response:")
	fmt.Printf("\tDetails: %s\n", response.GetDetails())
}

func DeleteVaulKey(subscriptionId string,keyId string) {
	configuration := omnicore.NewConfiguration()
	configuration.AddDefaultHeader("x-api-key", apiKey)
	apiClient := omnicore.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), omnicore.ContextAccessToken, jwtToken)
	response,resp,err := apiClient.VaultApi.DeleteVaultKey(ctx,subscriptionId,keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultApi.DeleteVaultKey`: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		return
	}
	fmt.Println("Delete Vault Key Response:")
	fmt.Printf("\tDetails: %s\n", response.GetDetails())
}