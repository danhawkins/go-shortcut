# MaxSearchResultsExceededError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | The name for this type of error, &#x60;maximum-results-exceeded&#x60; | 
**Message** | **string** | An explanatory message: \&quot;A maximum of 1000 search results are supported.\&quot; | 
**MaximumResults** | **int64** | The maximum number of search results supported, &#x60;1000&#x60; | 

## Methods

### NewMaxSearchResultsExceededError

`func NewMaxSearchResultsExceededError(error_ string, message string, maximumResults int64, ) *MaxSearchResultsExceededError`

NewMaxSearchResultsExceededError instantiates a new MaxSearchResultsExceededError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxSearchResultsExceededErrorWithDefaults

`func NewMaxSearchResultsExceededErrorWithDefaults() *MaxSearchResultsExceededError`

NewMaxSearchResultsExceededErrorWithDefaults instantiates a new MaxSearchResultsExceededError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MaxSearchResultsExceededError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MaxSearchResultsExceededError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MaxSearchResultsExceededError) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *MaxSearchResultsExceededError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MaxSearchResultsExceededError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MaxSearchResultsExceededError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMaximumResults

`func (o *MaxSearchResultsExceededError) GetMaximumResults() int64`

GetMaximumResults returns the MaximumResults field if non-nil, zero value otherwise.

### GetMaximumResultsOk

`func (o *MaxSearchResultsExceededError) GetMaximumResultsOk() (*int64, bool)`

GetMaximumResultsOk returns a tuple with the MaximumResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResults

`func (o *MaxSearchResultsExceededError) SetMaximumResults(v int64)`

SetMaximumResults sets MaximumResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


