# HistoryActionStoryLinkCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Action** | **string** | The action of the entity referenced. | 
**Verb** | **string** | The verb describing the link&#39;s relationship. | 
**SubjectId** | **int64** | The Story ID of the subject Story. | 
**ObjectId** | **int64** | The Story ID of the object Story. | 

## Methods

### NewHistoryActionStoryLinkCreate

`func NewHistoryActionStoryLinkCreate(id int64, entityType string, action string, verb string, subjectId int64, objectId int64, ) *HistoryActionStoryLinkCreate`

NewHistoryActionStoryLinkCreate instantiates a new HistoryActionStoryLinkCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionStoryLinkCreateWithDefaults

`func NewHistoryActionStoryLinkCreateWithDefaults() *HistoryActionStoryLinkCreate`

NewHistoryActionStoryLinkCreateWithDefaults instantiates a new HistoryActionStoryLinkCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryActionStoryLinkCreate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionStoryLinkCreate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionStoryLinkCreate) SetId(v int64)`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryActionStoryLinkCreate) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionStoryLinkCreate) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionStoryLinkCreate) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetAction

`func (o *HistoryActionStoryLinkCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionStoryLinkCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionStoryLinkCreate) SetAction(v string)`

SetAction sets Action field to given value.


### GetVerb

`func (o *HistoryActionStoryLinkCreate) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *HistoryActionStoryLinkCreate) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *HistoryActionStoryLinkCreate) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetSubjectId

`func (o *HistoryActionStoryLinkCreate) GetSubjectId() int64`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *HistoryActionStoryLinkCreate) GetSubjectIdOk() (*int64, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *HistoryActionStoryLinkCreate) SetSubjectId(v int64)`

SetSubjectId sets SubjectId field to given value.


### GetObjectId

`func (o *HistoryActionStoryLinkCreate) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *HistoryActionStoryLinkCreate) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *HistoryActionStoryLinkCreate) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


