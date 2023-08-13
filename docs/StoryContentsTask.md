# StoryContentsTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Full text of the Task. | 
**Position** | Pointer to **int64** | The number corresponding to the Task&#39;s position within a list of Tasks on a Story. | [optional] 
**Complete** | Pointer to **bool** | True/false boolean indicating whether the Task has been completed. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs of the Owners of this Task. | [optional] 
**ExternalId** | Pointer to **NullableString** | This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 

## Methods

### NewStoryContentsTask

`func NewStoryContentsTask(description string, ) *StoryContentsTask`

NewStoryContentsTask instantiates a new StoryContentsTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryContentsTaskWithDefaults

`func NewStoryContentsTaskWithDefaults() *StoryContentsTask`

NewStoryContentsTaskWithDefaults instantiates a new StoryContentsTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StoryContentsTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoryContentsTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoryContentsTask) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPosition

`func (o *StoryContentsTask) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *StoryContentsTask) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *StoryContentsTask) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *StoryContentsTask) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetComplete

`func (o *StoryContentsTask) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *StoryContentsTask) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *StoryContentsTask) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *StoryContentsTask) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetOwnerIds

`func (o *StoryContentsTask) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *StoryContentsTask) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *StoryContentsTask) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *StoryContentsTask) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetExternalId

`func (o *StoryContentsTask) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *StoryContentsTask) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *StoryContentsTask) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *StoryContentsTask) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *StoryContentsTask) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *StoryContentsTask) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


