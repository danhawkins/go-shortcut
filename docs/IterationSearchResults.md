# IterationSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** | The total number of matches for the search query. The first 1000 matches can be paged through via the API. | 
**Data** | [**[]IterationSlim**](IterationSlim.md) | A list of search results. | 
**Next** | **NullableString** | The URL path and query string for the next page of search results. | 
**Cursors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIterationSearchResults

`func NewIterationSearchResults(total int64, data []IterationSlim, next NullableString, ) *IterationSearchResults`

NewIterationSearchResults instantiates a new IterationSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationSearchResultsWithDefaults

`func NewIterationSearchResultsWithDefaults() *IterationSearchResults`

NewIterationSearchResultsWithDefaults instantiates a new IterationSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *IterationSearchResults) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IterationSearchResults) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IterationSearchResults) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetData

`func (o *IterationSearchResults) GetData() []IterationSlim`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IterationSearchResults) GetDataOk() (*[]IterationSlim, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IterationSearchResults) SetData(v []IterationSlim)`

SetData sets Data field to given value.


### GetNext

`func (o *IterationSearchResults) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *IterationSearchResults) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *IterationSearchResults) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *IterationSearchResults) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *IterationSearchResults) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetCursors

`func (o *IterationSearchResults) GetCursors() []string`

GetCursors returns the Cursors field if non-nil, zero value otherwise.

### GetCursorsOk

`func (o *IterationSearchResults) GetCursorsOk() (*[]string, bool)`

GetCursorsOk returns a tuple with the Cursors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursors

`func (o *IterationSearchResults) SetCursors(v []string)`

SetCursors sets Cursors field to given value.

### HasCursors

`func (o *IterationSearchResults) HasCursors() bool`

HasCursors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


