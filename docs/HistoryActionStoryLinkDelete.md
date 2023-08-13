# HistoryActionStoryLinkDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Action** | **string** | The action of the entity referenced. | 
**Verb** | **string** | The verb describing the link&#39;s relationship. | 
**SubjectId** | **NullableInt64** | The Story ID of the subject Story. | 
**ObjectId** | **NullableInt64** | The Story ID of the object Story. | 

## Methods

### NewHistoryActionStoryLinkDelete

`func NewHistoryActionStoryLinkDelete(id int64, entityType string, action string, verb string, subjectId NullableInt64, objectId NullableInt64, ) *HistoryActionStoryLinkDelete`

NewHistoryActionStoryLinkDelete instantiates a new HistoryActionStoryLinkDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionStoryLinkDeleteWithDefaults

`func NewHistoryActionStoryLinkDeleteWithDefaults() *HistoryActionStoryLinkDelete`

NewHistoryActionStoryLinkDeleteWithDefaults instantiates a new HistoryActionStoryLinkDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryActionStoryLinkDelete) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionStoryLinkDelete) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionStoryLinkDelete) SetId(v int64)`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryActionStoryLinkDelete) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionStoryLinkDelete) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionStoryLinkDelete) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetAction

`func (o *HistoryActionStoryLinkDelete) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionStoryLinkDelete) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionStoryLinkDelete) SetAction(v string)`

SetAction sets Action field to given value.


### GetVerb

`func (o *HistoryActionStoryLinkDelete) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *HistoryActionStoryLinkDelete) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *HistoryActionStoryLinkDelete) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetSubjectId

`func (o *HistoryActionStoryLinkDelete) GetSubjectId() int64`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *HistoryActionStoryLinkDelete) GetSubjectIdOk() (*int64, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *HistoryActionStoryLinkDelete) SetSubjectId(v int64)`

SetSubjectId sets SubjectId field to given value.


### SetSubjectIdNil

`func (o *HistoryActionStoryLinkDelete) SetSubjectIdNil(b bool)`

 SetSubjectIdNil sets the value for SubjectId to be an explicit nil

### UnsetSubjectId
`func (o *HistoryActionStoryLinkDelete) UnsetSubjectId()`

UnsetSubjectId ensures that no value is present for SubjectId, not even an explicit nil
### GetObjectId

`func (o *HistoryActionStoryLinkDelete) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *HistoryActionStoryLinkDelete) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *HistoryActionStoryLinkDelete) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### SetObjectIdNil

`func (o *HistoryActionStoryLinkDelete) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *HistoryActionStoryLinkDelete) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


