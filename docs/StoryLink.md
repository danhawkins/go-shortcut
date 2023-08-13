# StoryLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | 
**Id** | **int64** | The unique identifier of the Story Link. | 
**SubjectId** | **int64** | The ID of the subject Story. | 
**SubjectWorkflowStateId** | **int64** | The workflow state of the \&quot;subject\&quot; story. | 
**Verb** | **string** | How the subject Story acts on the object Story. This can be \&quot;blocks\&quot;, \&quot;duplicates\&quot;, or \&quot;relates to\&quot;. | 
**ObjectId** | **int64** | The ID of the object Story. | 
**CreatedAt** | **time.Time** | The time/date when the Story Link was created. | 
**UpdatedAt** | **time.Time** | The time/date when the Story Link was last updated. | 

## Methods

### NewStoryLink

`func NewStoryLink(entityType string, id int64, subjectId int64, subjectWorkflowStateId int64, verb string, objectId int64, createdAt time.Time, updatedAt time.Time, ) *StoryLink`

NewStoryLink instantiates a new StoryLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryLinkWithDefaults

`func NewStoryLinkWithDefaults() *StoryLink`

NewStoryLinkWithDefaults instantiates a new StoryLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *StoryLink) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *StoryLink) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *StoryLink) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetId

`func (o *StoryLink) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoryLink) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoryLink) SetId(v int64)`

SetId sets Id field to given value.


### GetSubjectId

`func (o *StoryLink) GetSubjectId() int64`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *StoryLink) GetSubjectIdOk() (*int64, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *StoryLink) SetSubjectId(v int64)`

SetSubjectId sets SubjectId field to given value.


### GetSubjectWorkflowStateId

`func (o *StoryLink) GetSubjectWorkflowStateId() int64`

GetSubjectWorkflowStateId returns the SubjectWorkflowStateId field if non-nil, zero value otherwise.

### GetSubjectWorkflowStateIdOk

`func (o *StoryLink) GetSubjectWorkflowStateIdOk() (*int64, bool)`

GetSubjectWorkflowStateIdOk returns a tuple with the SubjectWorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectWorkflowStateId

`func (o *StoryLink) SetSubjectWorkflowStateId(v int64)`

SetSubjectWorkflowStateId sets SubjectWorkflowStateId field to given value.


### GetVerb

`func (o *StoryLink) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *StoryLink) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *StoryLink) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetObjectId

`func (o *StoryLink) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *StoryLink) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *StoryLink) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetCreatedAt

`func (o *StoryLink) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StoryLink) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StoryLink) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *StoryLink) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StoryLink) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StoryLink) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


