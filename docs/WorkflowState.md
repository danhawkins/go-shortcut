# WorkflowState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of what sort of Stories belong in that Workflow state. | 
**EntityType** | **string** | A string description of this resource. | 
**Color** | Pointer to **string** | The hex color for this Workflow State. | [optional] 
**Verb** | **NullableString** | The verb that triggers a move to that Workflow State when making VCS commits. | 
**Name** | **string** | The Workflow State&#39;s name. | 
**GlobalId** | **string** |  | 
**NumStories** | **int64** | The number of Stories currently in that Workflow State. | 
**Type** | **string** | The type of Workflow State (Unstarted, Started, or Finished) | 
**UpdatedAt** | **time.Time** | When the Workflow State was last updated. | 
**Id** | **int64** | The unique ID of the Workflow State. | 
**NumStoryTemplates** | **int64** | The number of Story Templates associated with that Workflow State. | 
**Position** | **int64** | The position that the Workflow State is in, starting with 0 at the left. | 
**CreatedAt** | **time.Time** | The time/date the Workflow State was created. | 

## Methods

### NewWorkflowState

`func NewWorkflowState(description string, entityType string, verb NullableString, name string, globalId string, numStories int64, type_ string, updatedAt time.Time, id int64, numStoryTemplates int64, position int64, createdAt time.Time, ) *WorkflowState`

NewWorkflowState instantiates a new WorkflowState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStateWithDefaults

`func NewWorkflowStateWithDefaults() *WorkflowState`

NewWorkflowStateWithDefaults instantiates a new WorkflowState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowState) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowState) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowState) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntityType

`func (o *WorkflowState) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *WorkflowState) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *WorkflowState) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetColor

`func (o *WorkflowState) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WorkflowState) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WorkflowState) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WorkflowState) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetVerb

`func (o *WorkflowState) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *WorkflowState) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *WorkflowState) SetVerb(v string)`

SetVerb sets Verb field to given value.


### SetVerbNil

`func (o *WorkflowState) SetVerbNil(b bool)`

 SetVerbNil sets the value for Verb to be an explicit nil

### UnsetVerb
`func (o *WorkflowState) UnsetVerb()`

UnsetVerb ensures that no value is present for Verb, not even an explicit nil
### GetName

`func (o *WorkflowState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowState) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *WorkflowState) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *WorkflowState) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *WorkflowState) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetNumStories

`func (o *WorkflowState) GetNumStories() int64`

GetNumStories returns the NumStories field if non-nil, zero value otherwise.

### GetNumStoriesOk

`func (o *WorkflowState) GetNumStoriesOk() (*int64, bool)`

GetNumStoriesOk returns a tuple with the NumStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStories

`func (o *WorkflowState) SetNumStories(v int64)`

SetNumStories sets NumStories field to given value.


### GetType

`func (o *WorkflowState) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowState) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowState) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *WorkflowState) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkflowState) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkflowState) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetId

`func (o *WorkflowState) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowState) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowState) SetId(v int64)`

SetId sets Id field to given value.


### GetNumStoryTemplates

`func (o *WorkflowState) GetNumStoryTemplates() int64`

GetNumStoryTemplates returns the NumStoryTemplates field if non-nil, zero value otherwise.

### GetNumStoryTemplatesOk

`func (o *WorkflowState) GetNumStoryTemplatesOk() (*int64, bool)`

GetNumStoryTemplatesOk returns a tuple with the NumStoryTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoryTemplates

`func (o *WorkflowState) SetNumStoryTemplates(v int64)`

SetNumStoryTemplates sets NumStoryTemplates field to given value.


### GetPosition

`func (o *WorkflowState) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WorkflowState) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WorkflowState) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetCreatedAt

`func (o *WorkflowState) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowState) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowState) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


