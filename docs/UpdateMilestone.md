# UpdateMilestone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The Milestone&#39;s description. | [optional] 
**Archived** | Pointer to **bool** | A boolean indicating whether the Milestone is archived or not | [optional] 
**CompletedAtOverride** | Pointer to **NullableTime** | A manual override for the time/date the Milestone was completed. | [optional] 
**Name** | Pointer to **string** | The name of the Milestone. | [optional] 
**State** | Pointer to **string** | The workflow state that the Milestone is in. | [optional] 
**StartedAtOverride** | Pointer to **NullableTime** | A manual override for the time/date the Milestone was started. | [optional] 
**Categories** | Pointer to [**[]CreateCategoryParams**](CreateCategoryParams.md) | An array of IDs of Categories attached to the Milestone. | [optional] 
**BeforeId** | Pointer to **int64** | The ID of the Milestone we want to move this Milestone before. | [optional] 
**AfterId** | Pointer to **int64** | The ID of the Milestone we want to move this Milestone after. | [optional] 

## Methods

### NewUpdateMilestone

`func NewUpdateMilestone() *UpdateMilestone`

NewUpdateMilestone instantiates a new UpdateMilestone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMilestoneWithDefaults

`func NewUpdateMilestoneWithDefaults() *UpdateMilestone`

NewUpdateMilestoneWithDefaults instantiates a new UpdateMilestone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateMilestone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMilestone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMilestone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMilestone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *UpdateMilestone) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateMilestone) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateMilestone) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateMilestone) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCompletedAtOverride

`func (o *UpdateMilestone) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *UpdateMilestone) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *UpdateMilestone) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.

### HasCompletedAtOverride

`func (o *UpdateMilestone) HasCompletedAtOverride() bool`

HasCompletedAtOverride returns a boolean if a field has been set.

### SetCompletedAtOverrideNil

`func (o *UpdateMilestone) SetCompletedAtOverrideNil(b bool)`

 SetCompletedAtOverrideNil sets the value for CompletedAtOverride to be an explicit nil

### UnsetCompletedAtOverride
`func (o *UpdateMilestone) UnsetCompletedAtOverride()`

UnsetCompletedAtOverride ensures that no value is present for CompletedAtOverride, not even an explicit nil
### GetName

`func (o *UpdateMilestone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMilestone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMilestone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMilestone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *UpdateMilestone) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateMilestone) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateMilestone) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateMilestone) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStartedAtOverride

`func (o *UpdateMilestone) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *UpdateMilestone) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *UpdateMilestone) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.

### HasStartedAtOverride

`func (o *UpdateMilestone) HasStartedAtOverride() bool`

HasStartedAtOverride returns a boolean if a field has been set.

### SetStartedAtOverrideNil

`func (o *UpdateMilestone) SetStartedAtOverrideNil(b bool)`

 SetStartedAtOverrideNil sets the value for StartedAtOverride to be an explicit nil

### UnsetStartedAtOverride
`func (o *UpdateMilestone) UnsetStartedAtOverride()`

UnsetStartedAtOverride ensures that no value is present for StartedAtOverride, not even an explicit nil
### GetCategories

`func (o *UpdateMilestone) GetCategories() []CreateCategoryParams`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *UpdateMilestone) GetCategoriesOk() (*[]CreateCategoryParams, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *UpdateMilestone) SetCategories(v []CreateCategoryParams)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *UpdateMilestone) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetBeforeId

`func (o *UpdateMilestone) GetBeforeId() int64`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *UpdateMilestone) GetBeforeIdOk() (*int64, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *UpdateMilestone) SetBeforeId(v int64)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *UpdateMilestone) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetAfterId

`func (o *UpdateMilestone) GetAfterId() int64`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *UpdateMilestone) GetAfterIdOk() (*int64, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *UpdateMilestone) SetAfterId(v int64)`

SetAfterId sets AfterId field to given value.

### HasAfterId

`func (o *UpdateMilestone) HasAfterId() bool`

HasAfterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


