# EventNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PubsubTopicName** | Pointer to **string** | PubsubTopicName: A Cloud Pub/Sub topic name. For example, &#x60;projects/myProject/topics/deviceEvents&#x60;. | [optional] 
**SubfolderMatches** | Pointer to **string** | SubfolderMatches: If the subfolder name matches this string exactly, this configuration will be used. The string must not include the leading &#39;/&#39; character. If empty, all strings are matched. This field is used only for telemetry events; subfolders are not supported for state changes. | [optional] 

## Methods

### NewEventNotificationConfig

`func NewEventNotificationConfig() *EventNotificationConfig`

NewEventNotificationConfig instantiates a new EventNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationConfigWithDefaults

`func NewEventNotificationConfigWithDefaults() *EventNotificationConfig`

NewEventNotificationConfigWithDefaults instantiates a new EventNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPubsubTopicName

`func (o *EventNotificationConfig) GetPubsubTopicName() string`

GetPubsubTopicName returns the PubsubTopicName field if non-nil, zero value otherwise.

### GetPubsubTopicNameOk

`func (o *EventNotificationConfig) GetPubsubTopicNameOk() (*string, bool)`

GetPubsubTopicNameOk returns a tuple with the PubsubTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubsubTopicName

`func (o *EventNotificationConfig) SetPubsubTopicName(v string)`

SetPubsubTopicName sets PubsubTopicName field to given value.

### HasPubsubTopicName

`func (o *EventNotificationConfig) HasPubsubTopicName() bool`

HasPubsubTopicName returns a boolean if a field has been set.

### GetSubfolderMatches

`func (o *EventNotificationConfig) GetSubfolderMatches() string`

GetSubfolderMatches returns the SubfolderMatches field if non-nil, zero value otherwise.

### GetSubfolderMatchesOk

`func (o *EventNotificationConfig) GetSubfolderMatchesOk() (*string, bool)`

GetSubfolderMatchesOk returns a tuple with the SubfolderMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubfolderMatches

`func (o *EventNotificationConfig) SetSubfolderMatches(v string)`

SetSubfolderMatches sets SubfolderMatches field to given value.

### HasSubfolderMatches

`func (o *EventNotificationConfig) HasSubfolderMatches() bool`

HasSubfolderMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


