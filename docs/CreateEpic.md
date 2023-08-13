# CreateEpic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The Epic&#39;s description. | [optional] 
**Labels** | Pointer to [**[]CreateLabelParams**](CreateLabelParams.md) | An array of Labels attached to the Epic. | [optional] 
**CompletedAtOverride** | Pointer to **time.Time** | A manual override for the time/date the Epic was completed. | [optional] 
**Name** | **string** | The Epic&#39;s name. | 
**PlannedStartDate** | Pointer to **NullableTime** | The Epic&#39;s planned start date. | [optional] 
**State** | Pointer to **string** | &#x60;Deprecated&#x60; The Epic&#39;s state (to do, in progress, or done); will be ignored when &#x60;epic_state_id&#x60; is set. | [optional] 
**MilestoneId** | Pointer to **NullableInt64** | The ID of the Milestone this Epic is related to. | [optional] 
**RequestedById** | Pointer to **string** | The ID of the member that requested the epic. | [optional] 
**EpicStateId** | Pointer to **int64** | The ID of the Epic State. | [optional] 
**StartedAtOverride** | Pointer to **time.Time** | A manual override for the time/date the Epic was started. | [optional] 
**GroupId** | Pointer to **NullableString** | The ID of the group to associate with the epic. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Defaults to the time/date it is created but can be set to reflect another date. | [optional] 
**FollowerIds** | Pointer to **[]string** | An array of UUIDs for any Members you want to add as Followers on this new Epic. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs for any members you want to add as Owners on this new Epic. | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 
**Deadline** | Pointer to **NullableTime** | The Epic&#39;s deadline. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Defaults to the time/date it is created but can be set to reflect another date. | [optional] 

## Methods

### NewCreateEpic

`func NewCreateEpic(name string, ) *CreateEpic`

NewCreateEpic instantiates a new CreateEpic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEpicWithDefaults

`func NewCreateEpicWithDefaults() *CreateEpic`

NewCreateEpicWithDefaults instantiates a new CreateEpic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateEpic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEpic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEpic) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEpic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *CreateEpic) GetLabels() []CreateLabelParams`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateEpic) GetLabelsOk() (*[]CreateLabelParams, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateEpic) SetLabels(v []CreateLabelParams)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateEpic) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCompletedAtOverride

`func (o *CreateEpic) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *CreateEpic) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *CreateEpic) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.

### HasCompletedAtOverride

`func (o *CreateEpic) HasCompletedAtOverride() bool`

HasCompletedAtOverride returns a boolean if a field has been set.

### GetName

`func (o *CreateEpic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEpic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEpic) SetName(v string)`

SetName sets Name field to given value.


### GetPlannedStartDate

`func (o *CreateEpic) GetPlannedStartDate() time.Time`

GetPlannedStartDate returns the PlannedStartDate field if non-nil, zero value otherwise.

### GetPlannedStartDateOk

`func (o *CreateEpic) GetPlannedStartDateOk() (*time.Time, bool)`

GetPlannedStartDateOk returns a tuple with the PlannedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartDate

`func (o *CreateEpic) SetPlannedStartDate(v time.Time)`

SetPlannedStartDate sets PlannedStartDate field to given value.

### HasPlannedStartDate

`func (o *CreateEpic) HasPlannedStartDate() bool`

HasPlannedStartDate returns a boolean if a field has been set.

### SetPlannedStartDateNil

`func (o *CreateEpic) SetPlannedStartDateNil(b bool)`

 SetPlannedStartDateNil sets the value for PlannedStartDate to be an explicit nil

### UnsetPlannedStartDate
`func (o *CreateEpic) UnsetPlannedStartDate()`

UnsetPlannedStartDate ensures that no value is present for PlannedStartDate, not even an explicit nil
### GetState

`func (o *CreateEpic) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateEpic) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateEpic) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateEpic) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMilestoneId

`func (o *CreateEpic) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *CreateEpic) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *CreateEpic) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *CreateEpic) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *CreateEpic) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *CreateEpic) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetRequestedById

`func (o *CreateEpic) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *CreateEpic) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *CreateEpic) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.

### HasRequestedById

`func (o *CreateEpic) HasRequestedById() bool`

HasRequestedById returns a boolean if a field has been set.

### GetEpicStateId

`func (o *CreateEpic) GetEpicStateId() int64`

GetEpicStateId returns the EpicStateId field if non-nil, zero value otherwise.

### GetEpicStateIdOk

`func (o *CreateEpic) GetEpicStateIdOk() (*int64, bool)`

GetEpicStateIdOk returns a tuple with the EpicStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicStateId

`func (o *CreateEpic) SetEpicStateId(v int64)`

SetEpicStateId sets EpicStateId field to given value.

### HasEpicStateId

`func (o *CreateEpic) HasEpicStateId() bool`

HasEpicStateId returns a boolean if a field has been set.

### GetStartedAtOverride

`func (o *CreateEpic) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *CreateEpic) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *CreateEpic) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.

### HasStartedAtOverride

`func (o *CreateEpic) HasStartedAtOverride() bool`

HasStartedAtOverride returns a boolean if a field has been set.

### GetGroupId

`func (o *CreateEpic) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateEpic) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateEpic) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CreateEpic) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CreateEpic) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CreateEpic) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetUpdatedAt

`func (o *CreateEpic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateEpic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateEpic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateEpic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFollowerIds

`func (o *CreateEpic) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *CreateEpic) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *CreateEpic) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *CreateEpic) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *CreateEpic) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *CreateEpic) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *CreateEpic) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *CreateEpic) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateEpic) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateEpic) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateEpic) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateEpic) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDeadline

`func (o *CreateEpic) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *CreateEpic) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *CreateEpic) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *CreateEpic) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### SetDeadlineNil

`func (o *CreateEpic) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *CreateEpic) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetCreatedAt

`func (o *CreateEpic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateEpic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateEpic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateEpic) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


