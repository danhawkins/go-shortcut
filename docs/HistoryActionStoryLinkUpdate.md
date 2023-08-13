# HistoryActionStoryLinkUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Action** | **string** | The action of the entity referenced. | 
**Verb** | **string** | The verb describing the link&#39;s relationship. | 
**SubjectId** | **int64** | The Story ID of the subject Story. | 
**ObjectId** | **int64** | The Story ID of the object Story. | 
**Changes** | [**HistoryChangesStoryLink**](HistoryChangesStoryLink.md) |  | 

## Methods

### NewHistoryActionStoryLinkUpdate

`func NewHistoryActionStoryLinkUpdate(id int64, entityType string, action string, verb string, subjectId int64, objectId int64, changes HistoryChangesStoryLink, ) *HistoryActionStoryLinkUpdate`

NewHistoryActionStoryLinkUpdate instantiates a new HistoryActionStoryLinkUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionStoryLinkUpdateWithDefaults

`func NewHistoryActionStoryLinkUpdateWithDefaults() *HistoryActionStoryLinkUpdate`

NewHistoryActionStoryLinkUpdateWithDefaults instantiates a new HistoryActionStoryLinkUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryActionStoryLinkUpdate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionStoryLinkUpdate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionStoryLinkUpdate) SetId(v int64)`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryActionStoryLinkUpdate) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionStoryLinkUpdate) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionStoryLinkUpdate) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetAction

`func (o *HistoryActionStoryLinkUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionStoryLinkUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionStoryLinkUpdate) SetAction(v string)`

SetAction sets Action field to given value.


### GetVerb

`func (o *HistoryActionStoryLinkUpdate) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *HistoryActionStoryLinkUpdate) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *HistoryActionStoryLinkUpdate) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetSubjectId

`func (o *HistoryActionStoryLinkUpdate) GetSubjectId() int64`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *HistoryActionStoryLinkUpdate) GetSubjectIdOk() (*int64, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *HistoryActionStoryLinkUpdate) SetSubjectId(v int64)`

SetSubjectId sets SubjectId field to given value.


### GetObjectId

`func (o *HistoryActionStoryLinkUpdate) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *HistoryActionStoryLinkUpdate) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *HistoryActionStoryLinkUpdate) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetChanges

`func (o *HistoryActionStoryLinkUpdate) GetChanges() HistoryChangesStoryLink`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *HistoryActionStoryLinkUpdate) GetChangesOk() (*HistoryChangesStoryLink, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *HistoryActionStoryLinkUpdate) SetChanges(v HistoryChangesStoryLink)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


