# EpicWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | 
**Id** | **int64** | The unique ID of the Epic Workflow. | 
**CreatedAt** | **time.Time** | The date the Epic Workflow was created. | 
**UpdatedAt** | **time.Time** | The date the Epic Workflow was updated. | 
**DefaultEpicStateId** | **int64** | The unique ID of the default Epic State that new Epics are assigned by default. | 
**EpicStates** | [**[]EpicState**](EpicState.md) | A map of the Epic States in this Epic Workflow. | 

## Methods

### NewEpicWorkflow

`func NewEpicWorkflow(entityType string, id int64, createdAt time.Time, updatedAt time.Time, defaultEpicStateId int64, epicStates []EpicState, ) *EpicWorkflow`

NewEpicWorkflow instantiates a new EpicWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpicWorkflowWithDefaults

`func NewEpicWorkflowWithDefaults() *EpicWorkflow`

NewEpicWorkflowWithDefaults instantiates a new EpicWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *EpicWorkflow) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EpicWorkflow) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EpicWorkflow) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetId

`func (o *EpicWorkflow) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EpicWorkflow) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EpicWorkflow) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EpicWorkflow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EpicWorkflow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EpicWorkflow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EpicWorkflow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EpicWorkflow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EpicWorkflow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDefaultEpicStateId

`func (o *EpicWorkflow) GetDefaultEpicStateId() int64`

GetDefaultEpicStateId returns the DefaultEpicStateId field if non-nil, zero value otherwise.

### GetDefaultEpicStateIdOk

`func (o *EpicWorkflow) GetDefaultEpicStateIdOk() (*int64, bool)`

GetDefaultEpicStateIdOk returns a tuple with the DefaultEpicStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEpicStateId

`func (o *EpicWorkflow) SetDefaultEpicStateId(v int64)`

SetDefaultEpicStateId sets DefaultEpicStateId field to given value.


### GetEpicStates

`func (o *EpicWorkflow) GetEpicStates() []EpicState`

GetEpicStates returns the EpicStates field if non-nil, zero value otherwise.

### GetEpicStatesOk

`func (o *EpicWorkflow) GetEpicStatesOk() (*[]EpicState, bool)`

GetEpicStatesOk returns a tuple with the EpicStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicStates

`func (o *EpicWorkflow) SetEpicStates(v []EpicState)`

SetEpicStates sets EpicStates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


