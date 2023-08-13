# ListGroupStories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** | The maximum number of results to return. (Defaults to 1000, max 1000) | [optional] 
**Offset** | Pointer to **int64** | The offset at which to begin returning results. (Defaults to 0) | [optional] 

## Methods

### NewListGroupStories

`func NewListGroupStories() *ListGroupStories`

NewListGroupStories instantiates a new ListGroupStories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGroupStoriesWithDefaults

`func NewListGroupStoriesWithDefaults() *ListGroupStories`

NewListGroupStoriesWithDefaults instantiates a new ListGroupStories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListGroupStories) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListGroupStories) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListGroupStories) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListGroupStories) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListGroupStories) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListGroupStories) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListGroupStories) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListGroupStories) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


