# TypedStoryLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | 
**ObjectId** | **int64** | The ID of the object Story. | 
**Verb** | **string** | How the subject Story acts on the object Story. This can be \&quot;blocks\&quot;, \&quot;duplicates\&quot;, or \&quot;relates to\&quot;. | 
**Type** | **string** | This indicates whether the Story is the subject or object in the Story Link. | 
**UpdatedAt** | **time.Time** | The time/date when the Story Link was last updated. | 
**Id** | **int64** | The unique identifier of the Story Link. | 
**SubjectId** | **int64** | The ID of the subject Story. | 
**SubjectWorkflowStateId** | **int64** | The workflow state of the \&quot;subject\&quot; story. | 
**CreatedAt** | **time.Time** | The time/date when the Story Link was created. | 

## Methods

### NewTypedStoryLink

`func NewTypedStoryLink(entityType string, objectId int64, verb string, type_ string, updatedAt time.Time, id int64, subjectId int64, subjectWorkflowStateId int64, createdAt time.Time, ) *TypedStoryLink`

NewTypedStoryLink instantiates a new TypedStoryLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypedStoryLinkWithDefaults

`func NewTypedStoryLinkWithDefaults() *TypedStoryLink`

NewTypedStoryLinkWithDefaults instantiates a new TypedStoryLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *TypedStoryLink) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TypedStoryLink) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TypedStoryLink) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetObjectId

`func (o *TypedStoryLink) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *TypedStoryLink) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *TypedStoryLink) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetVerb

`func (o *TypedStoryLink) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *TypedStoryLink) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *TypedStoryLink) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetType

`func (o *TypedStoryLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypedStoryLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypedStoryLink) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *TypedStoryLink) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TypedStoryLink) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TypedStoryLink) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetId

`func (o *TypedStoryLink) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypedStoryLink) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypedStoryLink) SetId(v int64)`

SetId sets Id field to given value.


### GetSubjectId

`func (o *TypedStoryLink) GetSubjectId() int64`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *TypedStoryLink) GetSubjectIdOk() (*int64, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *TypedStoryLink) SetSubjectId(v int64)`

SetSubjectId sets SubjectId field to given value.


### GetSubjectWorkflowStateId

`func (o *TypedStoryLink) GetSubjectWorkflowStateId() int64`

GetSubjectWorkflowStateId returns the SubjectWorkflowStateId field if non-nil, zero value otherwise.

### GetSubjectWorkflowStateIdOk

`func (o *TypedStoryLink) GetSubjectWorkflowStateIdOk() (*int64, bool)`

GetSubjectWorkflowStateIdOk returns a tuple with the SubjectWorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectWorkflowStateId

`func (o *TypedStoryLink) SetSubjectWorkflowStateId(v int64)`

SetSubjectWorkflowStateId sets SubjectWorkflowStateId field to given value.


### GetCreatedAt

`func (o *TypedStoryLink) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TypedStoryLink) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TypedStoryLink) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


