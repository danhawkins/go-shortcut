# HistoryActionStoryDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Action** | **string** | The action of the entity referenced. | 
**Name** | **string** | The name of the Story. | 
**StoryType** | **string** | The type of Story; either feature, bug, or chore. | 

## Methods

### NewHistoryActionStoryDelete

`func NewHistoryActionStoryDelete(id int64, entityType string, action string, name string, storyType string, ) *HistoryActionStoryDelete`

NewHistoryActionStoryDelete instantiates a new HistoryActionStoryDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionStoryDeleteWithDefaults

`func NewHistoryActionStoryDeleteWithDefaults() *HistoryActionStoryDelete`

NewHistoryActionStoryDeleteWithDefaults instantiates a new HistoryActionStoryDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryActionStoryDelete) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionStoryDelete) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionStoryDelete) SetId(v int64)`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryActionStoryDelete) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionStoryDelete) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionStoryDelete) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetAction

`func (o *HistoryActionStoryDelete) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionStoryDelete) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionStoryDelete) SetAction(v string)`

SetAction sets Action field to given value.


### GetName

`func (o *HistoryActionStoryDelete) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoryActionStoryDelete) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HistoryActionStoryDelete) SetName(v string)`

SetName sets Name field to given value.


### GetStoryType

`func (o *HistoryActionStoryDelete) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *HistoryActionStoryDelete) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *HistoryActionStoryDelete) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


