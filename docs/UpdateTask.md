# UpdateTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The Task&#39;s description. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs of the owners of this story. | [optional] 
**Complete** | Pointer to **bool** | A true/false boolean indicating whether the task is complete. | [optional] 
**BeforeId** | Pointer to **int64** | Move task before this task ID. | [optional] 
**AfterId** | Pointer to **int64** | Move task after this task ID. | [optional] 

## Methods

### NewUpdateTask

`func NewUpdateTask() *UpdateTask`

NewUpdateTask instantiates a new UpdateTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskWithDefaults

`func NewUpdateTaskWithDefaults() *UpdateTask`

NewUpdateTaskWithDefaults instantiates a new UpdateTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwnerIds

`func (o *UpdateTask) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *UpdateTask) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *UpdateTask) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *UpdateTask) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetComplete

`func (o *UpdateTask) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *UpdateTask) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *UpdateTask) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *UpdateTask) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetBeforeId

`func (o *UpdateTask) GetBeforeId() int64`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *UpdateTask) GetBeforeIdOk() (*int64, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *UpdateTask) SetBeforeId(v int64)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *UpdateTask) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetAfterId

`func (o *UpdateTask) GetAfterId() int64`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *UpdateTask) GetAfterIdOk() (*int64, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *UpdateTask) SetAfterId(v int64)`

SetAfterId sets AfterId field to given value.

### HasAfterId

`func (o *UpdateTask) HasAfterId() bool`

HasAfterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


