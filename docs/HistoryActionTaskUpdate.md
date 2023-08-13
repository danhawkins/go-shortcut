# HistoryActionTaskUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Action** | **string** | The action of the entity referenced. | 
**Changes** | [**HistoryChangesTask**](HistoryChangesTask.md) |  | 
**Complete** | Pointer to **bool** | Whether or not the Task is complete. | [optional] 
**Description** | **string** | The description of the Task. | 
**StoryId** | **int64** | The Story ID that contains the Task. | 

## Methods

### NewHistoryActionTaskUpdate

`func NewHistoryActionTaskUpdate(id int64, entityType string, action string, changes HistoryChangesTask, description string, storyId int64, ) *HistoryActionTaskUpdate`

NewHistoryActionTaskUpdate instantiates a new HistoryActionTaskUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionTaskUpdateWithDefaults

`func NewHistoryActionTaskUpdateWithDefaults() *HistoryActionTaskUpdate`

NewHistoryActionTaskUpdateWithDefaults instantiates a new HistoryActionTaskUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryActionTaskUpdate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionTaskUpdate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionTaskUpdate) SetId(v int64)`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryActionTaskUpdate) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionTaskUpdate) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionTaskUpdate) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetAction

`func (o *HistoryActionTaskUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionTaskUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionTaskUpdate) SetAction(v string)`

SetAction sets Action field to given value.


### GetChanges

`func (o *HistoryActionTaskUpdate) GetChanges() HistoryChangesTask`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *HistoryActionTaskUpdate) GetChangesOk() (*HistoryChangesTask, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *HistoryActionTaskUpdate) SetChanges(v HistoryChangesTask)`

SetChanges sets Changes field to given value.


### GetComplete

`func (o *HistoryActionTaskUpdate) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *HistoryActionTaskUpdate) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *HistoryActionTaskUpdate) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *HistoryActionTaskUpdate) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetDescription

`func (o *HistoryActionTaskUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HistoryActionTaskUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HistoryActionTaskUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStoryId

`func (o *HistoryActionTaskUpdate) GetStoryId() int64`

GetStoryId returns the StoryId field if non-nil, zero value otherwise.

### GetStoryIdOk

`func (o *HistoryActionTaskUpdate) GetStoryIdOk() (*int64, bool)`

GetStoryIdOk returns a tuple with the StoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryId

`func (o *HistoryActionTaskUpdate) SetStoryId(v int64)`

SetStoryId sets StoryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


