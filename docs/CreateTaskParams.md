# CreateTaskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The Task description. | 
**Complete** | Pointer to **bool** | True/false boolean indicating whether the Task is completed. Defaults to false. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs for any members you want to add as Owners on this new Task. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Defaults to the time/date the Task is created but can be set to reflect another creation time/date. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Defaults to the time/date the Task is created in Shortcut but can be set to reflect another time/date. | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 

## Methods

### NewCreateTaskParams

`func NewCreateTaskParams(description string, ) *CreateTaskParams`

NewCreateTaskParams instantiates a new CreateTaskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaskParamsWithDefaults

`func NewCreateTaskParamsWithDefaults() *CreateTaskParams`

NewCreateTaskParamsWithDefaults instantiates a new CreateTaskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateTaskParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTaskParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTaskParams) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetComplete

`func (o *CreateTaskParams) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *CreateTaskParams) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *CreateTaskParams) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *CreateTaskParams) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetOwnerIds

`func (o *CreateTaskParams) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *CreateTaskParams) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *CreateTaskParams) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *CreateTaskParams) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateTaskParams) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateTaskParams) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateTaskParams) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateTaskParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateTaskParams) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateTaskParams) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateTaskParams) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateTaskParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateTaskParams) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateTaskParams) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateTaskParams) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateTaskParams) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


