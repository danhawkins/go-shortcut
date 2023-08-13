# EntityTemplateTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The Task description. | 
**Complete** | Pointer to **bool** | True/false boolean indicating whether the Task is completed. Defaults to false. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs for any members you want to add as Owners on this new Task. | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 

## Methods

### NewEntityTemplateTask

`func NewEntityTemplateTask(description string, ) *EntityTemplateTask`

NewEntityTemplateTask instantiates a new EntityTemplateTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTemplateTaskWithDefaults

`func NewEntityTemplateTaskWithDefaults() *EntityTemplateTask`

NewEntityTemplateTaskWithDefaults instantiates a new EntityTemplateTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EntityTemplateTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityTemplateTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityTemplateTask) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetComplete

`func (o *EntityTemplateTask) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *EntityTemplateTask) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *EntityTemplateTask) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *EntityTemplateTask) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetOwnerIds

`func (o *EntityTemplateTask) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *EntityTemplateTask) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *EntityTemplateTask) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *EntityTemplateTask) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetExternalId

`func (o *EntityTemplateTask) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EntityTemplateTask) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EntityTemplateTask) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EntityTemplateTask) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


