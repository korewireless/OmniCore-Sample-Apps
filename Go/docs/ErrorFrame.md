# ErrorFrame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Details** | **string** |  | 
**Message** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewErrorFrame

`func NewErrorFrame(code int32, details string, message string, status string, ) *ErrorFrame`

NewErrorFrame instantiates a new ErrorFrame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorFrameWithDefaults

`func NewErrorFrameWithDefaults() *ErrorFrame`

NewErrorFrameWithDefaults instantiates a new ErrorFrame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorFrame) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorFrame) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorFrame) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDetails

`func (o *ErrorFrame) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorFrame) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorFrame) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetMessage

`func (o *ErrorFrame) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorFrame) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorFrame) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatus

`func (o *ErrorFrame) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorFrame) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorFrame) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


