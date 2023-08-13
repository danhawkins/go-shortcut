# UpdateEpic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The Epic&#39;s description. | [optional] 
**Archived** | Pointer to **bool** | A true/false boolean indicating whether the Epic is in archived state. | [optional] 
**Labels** | Pointer to [**[]CreateLabelParams**](CreateLabelParams.md) | An array of Labels attached to the Epic. | [optional] 
**CompletedAtOverride** | Pointer to **NullableTime** | A manual override for the time/date the Epic was completed. | [optional] 
**Name** | Pointer to **string** | The Epic&#39;s name. | [optional] 
**PlannedStartDate** | Pointer to **NullableTime** | The Epic&#39;s planned start date. | [optional] 
**State** | Pointer to **string** | &#x60;Deprecated&#x60; The Epic&#39;s state (to do, in progress, or done); will be ignored when &#x60;epic_state_id&#x60; is set. | [optional] 
**MilestoneId** | Pointer to **NullableInt64** | The ID of the Milestone this Epic is related to. | [optional] 
**RequestedById** | Pointer to **string** | The ID of the member that requested the epic. | [optional] 
**EpicStateId** | Pointer to **int64** | The ID of the Epic State. | [optional] 
**StartedAtOverride** | Pointer to **NullableTime** | A manual override for the time/date the Epic was started. | [optional] 
**GroupId** | Pointer to **NullableString** | The ID of the group to associate with the epic. | [optional] 
**FollowerIds** | Pointer to **[]string** | An array of UUIDs for any Members you want to add as Followers on this Epic. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs for any members you want to add as Owners on this Epic. | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 
**BeforeId** | Pointer to **int64** | The ID of the Epic we want to move this Epic before. | [optional] 
**AfterId** | Pointer to **int64** | The ID of the Epic we want to move this Epic after. | [optional] 
**Deadline** | Pointer to **NullableTime** | The Epic&#39;s deadline. | [optional] 

## Methods

### NewUpdateEpic

`func NewUpdateEpic() *UpdateEpic`

NewUpdateEpic instantiates a new UpdateEpic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEpicWithDefaults

`func NewUpdateEpicWithDefaults() *UpdateEpic`

NewUpdateEpicWithDefaults instantiates a new UpdateEpic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateEpic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateEpic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateEpic) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateEpic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *UpdateEpic) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateEpic) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateEpic) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateEpic) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateEpic) GetLabels() []CreateLabelParams`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateEpic) GetLabelsOk() (*[]CreateLabelParams, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateEpic) SetLabels(v []CreateLabelParams)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateEpic) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCompletedAtOverride

`func (o *UpdateEpic) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *UpdateEpic) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *UpdateEpic) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.

### HasCompletedAtOverride

`func (o *UpdateEpic) HasCompletedAtOverride() bool`

HasCompletedAtOverride returns a boolean if a field has been set.

### SetCompletedAtOverrideNil

`func (o *UpdateEpic) SetCompletedAtOverrideNil(b bool)`

 SetCompletedAtOverrideNil sets the value for CompletedAtOverride to be an explicit nil

### UnsetCompletedAtOverride
`func (o *UpdateEpic) UnsetCompletedAtOverride()`

UnsetCompletedAtOverride ensures that no value is present for CompletedAtOverride, not even an explicit nil
### GetName

`func (o *UpdateEpic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEpic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEpic) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateEpic) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlannedStartDate

`func (o *UpdateEpic) GetPlannedStartDate() time.Time`

GetPlannedStartDate returns the PlannedStartDate field if non-nil, zero value otherwise.

### GetPlannedStartDateOk

`func (o *UpdateEpic) GetPlannedStartDateOk() (*time.Time, bool)`

GetPlannedStartDateOk returns a tuple with the PlannedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartDate

`func (o *UpdateEpic) SetPlannedStartDate(v time.Time)`

SetPlannedStartDate sets PlannedStartDate field to given value.

### HasPlannedStartDate

`func (o *UpdateEpic) HasPlannedStartDate() bool`

HasPlannedStartDate returns a boolean if a field has been set.

### SetPlannedStartDateNil

`func (o *UpdateEpic) SetPlannedStartDateNil(b bool)`

 SetPlannedStartDateNil sets the value for PlannedStartDate to be an explicit nil

### UnsetPlannedStartDate
`func (o *UpdateEpic) UnsetPlannedStartDate()`

UnsetPlannedStartDate ensures that no value is present for PlannedStartDate, not even an explicit nil
### GetState

`func (o *UpdateEpic) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateEpic) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateEpic) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateEpic) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMilestoneId

`func (o *UpdateEpic) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *UpdateEpic) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *UpdateEpic) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.

### HasMilestoneId

`func (o *UpdateEpic) HasMilestoneId() bool`

HasMilestoneId returns a boolean if a field has been set.

### SetMilestoneIdNil

`func (o *UpdateEpic) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *UpdateEpic) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetRequestedById

`func (o *UpdateEpic) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *UpdateEpic) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *UpdateEpic) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.

### HasRequestedById

`func (o *UpdateEpic) HasRequestedById() bool`

HasRequestedById returns a boolean if a field has been set.

### GetEpicStateId

`func (o *UpdateEpic) GetEpicStateId() int64`

GetEpicStateId returns the EpicStateId field if non-nil, zero value otherwise.

### GetEpicStateIdOk

`func (o *UpdateEpic) GetEpicStateIdOk() (*int64, bool)`

GetEpicStateIdOk returns a tuple with the EpicStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicStateId

`func (o *UpdateEpic) SetEpicStateId(v int64)`

SetEpicStateId sets EpicStateId field to given value.

### HasEpicStateId

`func (o *UpdateEpic) HasEpicStateId() bool`

HasEpicStateId returns a boolean if a field has been set.

### GetStartedAtOverride

`func (o *UpdateEpic) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *UpdateEpic) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *UpdateEpic) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.

### HasStartedAtOverride

`func (o *UpdateEpic) HasStartedAtOverride() bool`

HasStartedAtOverride returns a boolean if a field has been set.

### SetStartedAtOverrideNil

`func (o *UpdateEpic) SetStartedAtOverrideNil(b bool)`

 SetStartedAtOverrideNil sets the value for StartedAtOverride to be an explicit nil

### UnsetStartedAtOverride
`func (o *UpdateEpic) UnsetStartedAtOverride()`

UnsetStartedAtOverride ensures that no value is present for StartedAtOverride, not even an explicit nil
### GetGroupId

`func (o *UpdateEpic) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateEpic) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateEpic) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UpdateEpic) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *UpdateEpic) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *UpdateEpic) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetFollowerIds

`func (o *UpdateEpic) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *UpdateEpic) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *UpdateEpic) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *UpdateEpic) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *UpdateEpic) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *UpdateEpic) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *UpdateEpic) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *UpdateEpic) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateEpic) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateEpic) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateEpic) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateEpic) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetBeforeId

`func (o *UpdateEpic) GetBeforeId() int64`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *UpdateEpic) GetBeforeIdOk() (*int64, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *UpdateEpic) SetBeforeId(v int64)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *UpdateEpic) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetAfterId

`func (o *UpdateEpic) GetAfterId() int64`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *UpdateEpic) GetAfterIdOk() (*int64, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *UpdateEpic) SetAfterId(v int64)`

SetAfterId sets AfterId field to given value.

### HasAfterId

`func (o *UpdateEpic) HasAfterId() bool`

HasAfterId returns a boolean if a field has been set.

### GetDeadline

`func (o *UpdateEpic) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *UpdateEpic) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *UpdateEpic) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *UpdateEpic) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### SetDeadlineNil

`func (o *UpdateEpic) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *UpdateEpic) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


