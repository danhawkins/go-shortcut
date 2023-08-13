# HistoryActionPullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Action** | **string** | The action of the entity referenced. | 
**Number** | **int64** | The VCS Repository-specific ID for the Pull Request. | 
**Title** | **string** | The title of the Pull Request. | 
**Url** | **string** | The URL from the provider of the VCS Pull Request. | 

## Methods

### NewHistoryActionPullRequest

`func NewHistoryActionPullRequest(id int64, entityType string, action string, number int64, title string, url string, ) *HistoryActionPullRequest`

NewHistoryActionPullRequest instantiates a new HistoryActionPullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionPullRequestWithDefaults

`func NewHistoryActionPullRequestWithDefaults() *HistoryActionPullRequest`

NewHistoryActionPullRequestWithDefaults instantiates a new HistoryActionPullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryActionPullRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionPullRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionPullRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryActionPullRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionPullRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionPullRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetAction

`func (o *HistoryActionPullRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionPullRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionPullRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetNumber

`func (o *HistoryActionPullRequest) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *HistoryActionPullRequest) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *HistoryActionPullRequest) SetNumber(v int64)`

SetNumber sets Number field to given value.


### GetTitle

`func (o *HistoryActionPullRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HistoryActionPullRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HistoryActionPullRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *HistoryActionPullRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HistoryActionPullRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HistoryActionPullRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


