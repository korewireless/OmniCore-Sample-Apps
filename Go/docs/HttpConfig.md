# HttpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpEnabledState** | Pointer to **string** | HttpEnabledState: If enabled, allows devices to use DeviceService via the HTTP protocol. Otherwise, any requests to DeviceService will fail for this registry.  Possible values:   \&quot;HTTP_STATE_UNSPECIFIED\&quot; - No HTTP state specified. If not specified, DeviceService will be enabled by default.   \&quot;HTTP_ENABLED\&quot; - Enables DeviceService (HTTP) service for the registry.   \&quot;HTTP_DISABLED\&quot; - Disables DeviceService (HTTP) service for the registry. | [optional] 

## Methods

### NewHttpConfig

`func NewHttpConfig() *HttpConfig`

NewHttpConfig instantiates a new HttpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpConfigWithDefaults

`func NewHttpConfigWithDefaults() *HttpConfig`

NewHttpConfigWithDefaults instantiates a new HttpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpEnabledState

`func (o *HttpConfig) GetHttpEnabledState() string`

GetHttpEnabledState returns the HttpEnabledState field if non-nil, zero value otherwise.

### GetHttpEnabledStateOk

`func (o *HttpConfig) GetHttpEnabledStateOk() (*string, bool)`

GetHttpEnabledStateOk returns a tuple with the HttpEnabledState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEnabledState

`func (o *HttpConfig) SetHttpEnabledState(v string)`

SetHttpEnabledState sets HttpEnabledState field to given value.

### HasHttpEnabledState

`func (o *HttpConfig) HasHttpEnabledState() bool`

HasHttpEnabledState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


