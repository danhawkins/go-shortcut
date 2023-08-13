# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description of the workflow. | 
**EntityType** | **string** | A string description of this resource. | 
**ProjectIds** | **[]float64** | An array of IDs of projects within the Workflow. | 
**States** | [**[]WorkflowState**](WorkflowState.md) | A map of the states in this Workflow. | 
**Name** | **string** | The name of the workflow. | 
**UpdatedAt** | **time.Time** | The date the Workflow was updated. | 
**AutoAssignOwner** | **bool** | Indicates if an owner is automatically assigned when an unowned story is started. | 
**Id** | **int64** | The unique ID of the Workflow. | 
**TeamId** | **int64** | The ID of the team the workflow belongs to. | 
**CreatedAt** | **time.Time** | The date the Workflow was created. | 
**DefaultStateId** | **int64** | The unique ID of the default state that new Stories are entered into. | 

## Methods

### NewWorkflow

`func NewWorkflow(description string, entityType string, projectIds []float64, states []WorkflowState, name string, updatedAt time.Time, autoAssignOwner bool, id int64, teamId int64, createdAt time.Time, defaultStateId int64, ) *Workflow`

NewWorkflow instantiates a new Workflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithDefaults

`func NewWorkflowWithDefaults() *Workflow`

NewWorkflowWithDefaults instantiates a new Workflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Workflow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Workflow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Workflow) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntityType

`func (o *Workflow) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Workflow) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Workflow) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetProjectIds

`func (o *Workflow) GetProjectIds() []float64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *Workflow) GetProjectIdsOk() (*[]float64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *Workflow) SetProjectIds(v []float64)`

SetProjectIds sets ProjectIds field to given value.


### GetStates

`func (o *Workflow) GetStates() []WorkflowState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *Workflow) GetStatesOk() (*[]WorkflowState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *Workflow) SetStates(v []WorkflowState)`

SetStates sets States field to given value.


### GetName

`func (o *Workflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workflow) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *Workflow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Workflow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Workflow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAutoAssignOwner

`func (o *Workflow) GetAutoAssignOwner() bool`

GetAutoAssignOwner returns the AutoAssignOwner field if non-nil, zero value otherwise.

### GetAutoAssignOwnerOk

`func (o *Workflow) GetAutoAssignOwnerOk() (*bool, bool)`

GetAutoAssignOwnerOk returns a tuple with the AutoAssignOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssignOwner

`func (o *Workflow) SetAutoAssignOwner(v bool)`

SetAutoAssignOwner sets AutoAssignOwner field to given value.


### GetId

`func (o *Workflow) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workflow) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workflow) SetId(v int64)`

SetId sets Id field to given value.


### GetTeamId

`func (o *Workflow) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Workflow) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Workflow) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetCreatedAt

`func (o *Workflow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Workflow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Workflow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultStateId

`func (o *Workflow) GetDefaultStateId() int64`

GetDefaultStateId returns the DefaultStateId field if non-nil, zero value otherwise.

### GetDefaultStateIdOk

`func (o *Workflow) GetDefaultStateIdOk() (*int64, bool)`

GetDefaultStateIdOk returns a tuple with the DefaultStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStateId

`func (o *Workflow) SetDefaultStateId(v int64)`

SetDefaultStateId sets DefaultStateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


