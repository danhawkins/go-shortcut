# HistoryReferenceWorkflowState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **map[string]interface{}** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Type** | **string** | Either \&quot;unstarted\&quot;, \&quot;started\&quot;, or \&quot;done\&quot;. | 
**Name** | **string** | The name of the entity referenced. | 

## Methods

### NewHistoryReferenceWorkflowState

`func NewHistoryReferenceWorkflowState(id map[string]interface{}, entityType string, type_ string, name string, ) *HistoryReferenceWorkflowState`

NewHistoryReferenceWorkflowState instantiates a new HistoryReferenceWorkflowState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryReferenceWorkflowStateWithDefaults

`func NewHistoryReferenceWorkflowStateWithDefaults() *HistoryReferenceWorkflowState`

NewHistoryReferenceWorkflowStateWithDefaults instantiates a new HistoryReferenceWorkflowState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryReferenceWorkflowState) GetId() map[string]interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryReferenceWorkflowState) GetIdOk() (*map[string]interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryReferenceWorkflowState) SetId(v map[string]interface{})`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryReferenceWorkflowState) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryReferenceWorkflowState) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryReferenceWorkflowState) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetType

`func (o *HistoryReferenceWorkflowState) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoryReferenceWorkflowState) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoryReferenceWorkflowState) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *HistoryReferenceWorkflowState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoryReferenceWorkflowState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HistoryReferenceWorkflowState) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


