# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connect** | **bool** |  | 
**PublishState** | **bool** |  | 
**PublishEvents** | **bool** |  | 
**PublishEventsRegex** | **string** |  | 
**PublishLoopback** | **bool** |  | 
**SubscribeCommand** | **bool** |  | 
**SubscribeCommandRegex** | **string** |  | 
**SubscribeBroadcast** | **bool** |  | 
**SubscribeBroadcastRegex** | **string** |  | 
**SubscribeConfig** | **bool** |  | 

## Methods

### NewPolicy

`func NewPolicy(connect bool, publishState bool, publishEvents bool, publishEventsRegex string, publishLoopback bool, subscribeCommand bool, subscribeCommandRegex string, subscribeBroadcast bool, subscribeBroadcastRegex string, subscribeConfig bool, ) *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnect

`func (o *Policy) GetConnect() bool`

GetConnect returns the Connect field if non-nil, zero value otherwise.

### GetConnectOk

`func (o *Policy) GetConnectOk() (*bool, bool)`

GetConnectOk returns a tuple with the Connect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnect

`func (o *Policy) SetConnect(v bool)`

SetConnect sets Connect field to given value.


### GetPublishState

`func (o *Policy) GetPublishState() bool`

GetPublishState returns the PublishState field if non-nil, zero value otherwise.

### GetPublishStateOk

`func (o *Policy) GetPublishStateOk() (*bool, bool)`

GetPublishStateOk returns a tuple with the PublishState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishState

`func (o *Policy) SetPublishState(v bool)`

SetPublishState sets PublishState field to given value.


### GetPublishEvents

`func (o *Policy) GetPublishEvents() bool`

GetPublishEvents returns the PublishEvents field if non-nil, zero value otherwise.

### GetPublishEventsOk

`func (o *Policy) GetPublishEventsOk() (*bool, bool)`

GetPublishEventsOk returns a tuple with the PublishEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishEvents

`func (o *Policy) SetPublishEvents(v bool)`

SetPublishEvents sets PublishEvents field to given value.


### GetPublishEventsRegex

`func (o *Policy) GetPublishEventsRegex() string`

GetPublishEventsRegex returns the PublishEventsRegex field if non-nil, zero value otherwise.

### GetPublishEventsRegexOk

`func (o *Policy) GetPublishEventsRegexOk() (*string, bool)`

GetPublishEventsRegexOk returns a tuple with the PublishEventsRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishEventsRegex

`func (o *Policy) SetPublishEventsRegex(v string)`

SetPublishEventsRegex sets PublishEventsRegex field to given value.


### GetPublishLoopback

`func (o *Policy) GetPublishLoopback() bool`

GetPublishLoopback returns the PublishLoopback field if non-nil, zero value otherwise.

### GetPublishLoopbackOk

`func (o *Policy) GetPublishLoopbackOk() (*bool, bool)`

GetPublishLoopbackOk returns a tuple with the PublishLoopback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishLoopback

`func (o *Policy) SetPublishLoopback(v bool)`

SetPublishLoopback sets PublishLoopback field to given value.


### GetSubscribeCommand

`func (o *Policy) GetSubscribeCommand() bool`

GetSubscribeCommand returns the SubscribeCommand field if non-nil, zero value otherwise.

### GetSubscribeCommandOk

`func (o *Policy) GetSubscribeCommandOk() (*bool, bool)`

GetSubscribeCommandOk returns a tuple with the SubscribeCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeCommand

`func (o *Policy) SetSubscribeCommand(v bool)`

SetSubscribeCommand sets SubscribeCommand field to given value.


### GetSubscribeCommandRegex

`func (o *Policy) GetSubscribeCommandRegex() string`

GetSubscribeCommandRegex returns the SubscribeCommandRegex field if non-nil, zero value otherwise.

### GetSubscribeCommandRegexOk

`func (o *Policy) GetSubscribeCommandRegexOk() (*string, bool)`

GetSubscribeCommandRegexOk returns a tuple with the SubscribeCommandRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeCommandRegex

`func (o *Policy) SetSubscribeCommandRegex(v string)`

SetSubscribeCommandRegex sets SubscribeCommandRegex field to given value.


### GetSubscribeBroadcast

`func (o *Policy) GetSubscribeBroadcast() bool`

GetSubscribeBroadcast returns the SubscribeBroadcast field if non-nil, zero value otherwise.

### GetSubscribeBroadcastOk

`func (o *Policy) GetSubscribeBroadcastOk() (*bool, bool)`

GetSubscribeBroadcastOk returns a tuple with the SubscribeBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeBroadcast

`func (o *Policy) SetSubscribeBroadcast(v bool)`

SetSubscribeBroadcast sets SubscribeBroadcast field to given value.


### GetSubscribeBroadcastRegex

`func (o *Policy) GetSubscribeBroadcastRegex() string`

GetSubscribeBroadcastRegex returns the SubscribeBroadcastRegex field if non-nil, zero value otherwise.

### GetSubscribeBroadcastRegexOk

`func (o *Policy) GetSubscribeBroadcastRegexOk() (*string, bool)`

GetSubscribeBroadcastRegexOk returns a tuple with the SubscribeBroadcastRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeBroadcastRegex

`func (o *Policy) SetSubscribeBroadcastRegex(v string)`

SetSubscribeBroadcastRegex sets SubscribeBroadcastRegex field to given value.


### GetSubscribeConfig

`func (o *Policy) GetSubscribeConfig() bool`

GetSubscribeConfig returns the SubscribeConfig field if non-nil, zero value otherwise.

### GetSubscribeConfigOk

`func (o *Policy) GetSubscribeConfigOk() (*bool, bool)`

GetSubscribeConfigOk returns a tuple with the SubscribeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeConfig

`func (o *Policy) SetSubscribeConfig(v bool)`

SetSubscribeConfig sets SubscribeConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


