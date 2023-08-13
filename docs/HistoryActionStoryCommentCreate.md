# HistoryActionStoryCommentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Action** | **string** | The action of the entity referenced. | 
**AppUrl** | **string** | The application URL of the Story Comment. | 
**Text** | **string** | The text of the Story Comment. | 
**AuthorId** | **string** | The Member ID of who created the Story Comment. | 

## Methods

### NewHistoryActionStoryCommentCreate

`func NewHistoryActionStoryCommentCreate(id int64, entityType string, action string, appUrl string, text string, authorId string, ) *HistoryActionStoryCommentCreate`

NewHistoryActionStoryCommentCreate instantiates a new HistoryActionStoryCommentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionStoryCommentCreateWithDefaults

`func NewHistoryActionStoryCommentCreateWithDefaults() *HistoryActionStoryCommentCreate`

NewHistoryActionStoryCommentCreateWithDefaults instantiates a new HistoryActionStoryCommentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryActionStoryCommentCreate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionStoryCommentCreate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionStoryCommentCreate) SetId(v int64)`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryActionStoryCommentCreate) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionStoryCommentCreate) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionStoryCommentCreate) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetAction

`func (o *HistoryActionStoryCommentCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionStoryCommentCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionStoryCommentCreate) SetAction(v string)`

SetAction sets Action field to given value.


### GetAppUrl

`func (o *HistoryActionStoryCommentCreate) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *HistoryActionStoryCommentCreate) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *HistoryActionStoryCommentCreate) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetText

`func (o *HistoryActionStoryCommentCreate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *HistoryActionStoryCommentCreate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *HistoryActionStoryCommentCreate) SetText(v string)`

SetText sets Text field to given value.


### GetAuthorId

`func (o *HistoryActionStoryCommentCreate) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *HistoryActionStoryCommentCreate) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *HistoryActionStoryCommentCreate) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


