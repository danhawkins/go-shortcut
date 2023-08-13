# HistoryActionStoryUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Action** | **string** | The action of the entity referenced. | 
**AppUrl** | **string** | The application URL of the Story. | 
**Changes** | Pointer to [**HistoryChangesStory**](HistoryChangesStory.md) |  | [optional] 
**Name** | **string** | The name of the Story. | 
**StoryType** | **string** | The type of Story; either feature, bug, or chore. | 

## Methods

### NewHistoryActionStoryUpdate

`func NewHistoryActionStoryUpdate(id int64, entityType string, action string, appUrl string, name string, storyType string, ) *HistoryActionStoryUpdate`

NewHistoryActionStoryUpdate instantiates a new HistoryActionStoryUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionStoryUpdateWithDefaults

`func NewHistoryActionStoryUpdateWithDefaults() *HistoryActionStoryUpdate`

NewHistoryActionStoryUpdateWithDefaults instantiates a new HistoryActionStoryUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryActionStoryUpdate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionStoryUpdate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionStoryUpdate) SetId(v int64)`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryActionStoryUpdate) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionStoryUpdate) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionStoryUpdate) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetAction

`func (o *HistoryActionStoryUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionStoryUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionStoryUpdate) SetAction(v string)`

SetAction sets Action field to given value.


### GetAppUrl

`func (o *HistoryActionStoryUpdate) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *HistoryActionStoryUpdate) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *HistoryActionStoryUpdate) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetChanges

`func (o *HistoryActionStoryUpdate) GetChanges() HistoryChangesStory`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *HistoryActionStoryUpdate) GetChangesOk() (*HistoryChangesStory, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *HistoryActionStoryUpdate) SetChanges(v HistoryChangesStory)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *HistoryActionStoryUpdate) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetName

`func (o *HistoryActionStoryUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoryActionStoryUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HistoryActionStoryUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetStoryType

`func (o *HistoryActionStoryUpdate) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *HistoryActionStoryUpdate) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *HistoryActionStoryUpdate) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


