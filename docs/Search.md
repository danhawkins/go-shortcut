# Search

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | See our help center article on [search operators](https://help.shortcut.com/hc/en-us/articles/360000046646-Search-Operators) | 
**PageSize** | Pointer to **int64** | The number of search results to include in a page. Minimum of 1 and maximum of 25. | [optional] 
**Detail** | Pointer to **string** | The amount of detail included in each result item.    \&quot;full\&quot; will include all descriptions and comments and more fields on    related items such as pull requests, branches and tasks.    \&quot;slim\&quot; omits larger fulltext fields such as descriptions and comments    and only references related items by id.    The default is \&quot;full\&quot;. | [optional] 
**Next** | Pointer to **string** | The next page token. | [optional] 
**Include** | Pointer to **string** |  | [optional] 
**EntityTypes** | Pointer to **[]string** | A collection of entity_types to search. Defaults to story and epic. Supports: epic, iteration, milestone, story. | [optional] 

## Methods

### NewSearch

`func NewSearch(query string, ) *Search`

NewSearch instantiates a new Search object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchWithDefaults

`func NewSearchWithDefaults() *Search`

NewSearchWithDefaults instantiates a new Search object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *Search) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Search) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Search) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetPageSize

`func (o *Search) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Search) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Search) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Search) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetDetail

`func (o *Search) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Search) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Search) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Search) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetNext

`func (o *Search) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Search) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Search) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Search) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetInclude

`func (o *Search) GetInclude() string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *Search) GetIncludeOk() (*string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *Search) SetInclude(v string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *Search) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetEntityTypes

`func (o *Search) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *Search) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *Search) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.

### HasEntityTypes

`func (o *Search) HasEntityTypes() bool`

HasEntityTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


