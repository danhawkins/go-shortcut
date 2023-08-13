# EpicSlim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Epic. | 
**Description** | Pointer to **string** | The Epic&#39;s description. | [optional] 
**Archived** | **bool** | True/false boolean that indicates whether the Epic is archived or not. | 
**Started** | **bool** | A true/false boolean indicating if the Epic has been started. | 
**EntityType** | **string** | A string description of this resource. | 
**Labels** | [**[]LabelSlim**](LabelSlim.md) | An array of Labels attached to the Epic. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Epic description. | 
**AssociatedGroups** | [**[]EpicAssociatedGroup**](EpicAssociatedGroup.md) | An array containing Group IDs and Group-owned story counts for the Epic&#39;s associated groups. | 
**ProjectIds** | **[]int64** | The IDs of Projects related to this Epic. | 
**StoriesWithoutProjects** | **int64** | The number of stories in this epic which are not associated with a project. | 
**CompletedAtOverride** | **NullableTime** | A manual override for the time/date the Epic was completed. | 
**ProductboardPluginId** | **NullableString** | The ID of the associated productboard integration. | 
**StartedAt** | **NullableTime** | The time/date the Epic was started. | 
**CompletedAt** | **NullableTime** | The time/date the Epic was completed. | 
**Name** | **string** | The name of the Epic. | 
**GlobalId** | **string** |  | 
**Completed** | **bool** | A true/false boolean indicating if the Epic has been completed. | 
**ProductboardUrl** | **NullableString** | The URL of the associated productboard feature. | 
**PlannedStartDate** | **NullableTime** | The Epic&#39;s planned start date. | 
**State** | **string** | &#x60;Deprecated&#x60; The workflow state that the Epic is in. | 
**MilestoneId** | **NullableInt64** | The ID of the Milestone this Epic is related to. | 
**RequestedById** | **string** | The ID of the Member that requested the epic. | 
**EpicStateId** | **int64** | The ID of the Epic State. | 
**LabelIds** | **[]int64** | An array of Label ids attached to the Epic. | 
**StartedAtOverride** | **NullableTime** | A manual override for the time/date the Epic was started. | 
**GroupId** | **NullableString** |  | 
**UpdatedAt** | **NullableTime** | The time/date the Epic was updated. | 
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Epic description. | 
**ProductboardId** | **NullableString** | The ID of the associated productboard feature. | 
**FollowerIds** | **[]string** | An array of UUIDs for any Members you want to add as Followers on this Epic. | 
**OwnerIds** | **[]string** | An array of UUIDs for any members you want to add as Owners on this new Epic. | 
**ExternalId** | **NullableString** | This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here. | 
**Id** | **int64** | The unique ID of the Epic. | 
**Position** | **int64** | The Epic&#39;s relative position in the Epic workflow state. | 
**ProductboardName** | **NullableString** | The name of the associated productboard feature. | 
**Deadline** | **NullableTime** | The Epic&#39;s deadline. | 
**Stats** | [**EpicStats**](EpicStats.md) |  | 
**CreatedAt** | **NullableTime** | The time/date the Epic was created. | 

## Methods

### NewEpicSlim

`func NewEpicSlim(appUrl string, archived bool, started bool, entityType string, labels []LabelSlim, mentionIds []string, memberMentionIds []string, associatedGroups []EpicAssociatedGroup, projectIds []int64, storiesWithoutProjects int64, completedAtOverride NullableTime, productboardPluginId NullableString, startedAt NullableTime, completedAt NullableTime, name string, globalId string, completed bool, productboardUrl NullableString, plannedStartDate NullableTime, state string, milestoneId NullableInt64, requestedById string, epicStateId int64, labelIds []int64, startedAtOverride NullableTime, groupId NullableString, updatedAt NullableTime, groupMentionIds []string, productboardId NullableString, followerIds []string, ownerIds []string, externalId NullableString, id int64, position int64, productboardName NullableString, deadline NullableTime, stats EpicStats, createdAt NullableTime, ) *EpicSlim`

NewEpicSlim instantiates a new EpicSlim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpicSlimWithDefaults

`func NewEpicSlimWithDefaults() *EpicSlim`

NewEpicSlimWithDefaults instantiates a new EpicSlim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *EpicSlim) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *EpicSlim) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *EpicSlim) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetDescription

`func (o *EpicSlim) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EpicSlim) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EpicSlim) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EpicSlim) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *EpicSlim) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *EpicSlim) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *EpicSlim) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetStarted

`func (o *EpicSlim) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *EpicSlim) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *EpicSlim) SetStarted(v bool)`

SetStarted sets Started field to given value.


### GetEntityType

`func (o *EpicSlim) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EpicSlim) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EpicSlim) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetLabels

`func (o *EpicSlim) GetLabels() []LabelSlim`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EpicSlim) GetLabelsOk() (*[]LabelSlim, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EpicSlim) SetLabels(v []LabelSlim)`

SetLabels sets Labels field to given value.


### GetMentionIds

`func (o *EpicSlim) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *EpicSlim) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *EpicSlim) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetMemberMentionIds

`func (o *EpicSlim) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *EpicSlim) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *EpicSlim) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetAssociatedGroups

`func (o *EpicSlim) GetAssociatedGroups() []EpicAssociatedGroup`

GetAssociatedGroups returns the AssociatedGroups field if non-nil, zero value otherwise.

### GetAssociatedGroupsOk

`func (o *EpicSlim) GetAssociatedGroupsOk() (*[]EpicAssociatedGroup, bool)`

GetAssociatedGroupsOk returns a tuple with the AssociatedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedGroups

`func (o *EpicSlim) SetAssociatedGroups(v []EpicAssociatedGroup)`

SetAssociatedGroups sets AssociatedGroups field to given value.


### GetProjectIds

`func (o *EpicSlim) GetProjectIds() []int64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *EpicSlim) GetProjectIdsOk() (*[]int64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *EpicSlim) SetProjectIds(v []int64)`

SetProjectIds sets ProjectIds field to given value.


### GetStoriesWithoutProjects

`func (o *EpicSlim) GetStoriesWithoutProjects() int64`

GetStoriesWithoutProjects returns the StoriesWithoutProjects field if non-nil, zero value otherwise.

### GetStoriesWithoutProjectsOk

`func (o *EpicSlim) GetStoriesWithoutProjectsOk() (*int64, bool)`

GetStoriesWithoutProjectsOk returns a tuple with the StoriesWithoutProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoriesWithoutProjects

`func (o *EpicSlim) SetStoriesWithoutProjects(v int64)`

SetStoriesWithoutProjects sets StoriesWithoutProjects field to given value.


### GetCompletedAtOverride

`func (o *EpicSlim) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *EpicSlim) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *EpicSlim) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.


### SetCompletedAtOverrideNil

`func (o *EpicSlim) SetCompletedAtOverrideNil(b bool)`

 SetCompletedAtOverrideNil sets the value for CompletedAtOverride to be an explicit nil

### UnsetCompletedAtOverride
`func (o *EpicSlim) UnsetCompletedAtOverride()`

UnsetCompletedAtOverride ensures that no value is present for CompletedAtOverride, not even an explicit nil
### GetProductboardPluginId

`func (o *EpicSlim) GetProductboardPluginId() string`

GetProductboardPluginId returns the ProductboardPluginId field if non-nil, zero value otherwise.

### GetProductboardPluginIdOk

`func (o *EpicSlim) GetProductboardPluginIdOk() (*string, bool)`

GetProductboardPluginIdOk returns a tuple with the ProductboardPluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductboardPluginId

`func (o *EpicSlim) SetProductboardPluginId(v string)`

SetProductboardPluginId sets ProductboardPluginId field to given value.


### SetProductboardPluginIdNil

`func (o *EpicSlim) SetProductboardPluginIdNil(b bool)`

 SetProductboardPluginIdNil sets the value for ProductboardPluginId to be an explicit nil

### UnsetProductboardPluginId
`func (o *EpicSlim) UnsetProductboardPluginId()`

UnsetProductboardPluginId ensures that no value is present for ProductboardPluginId, not even an explicit nil
### GetStartedAt

`func (o *EpicSlim) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *EpicSlim) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *EpicSlim) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *EpicSlim) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *EpicSlim) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *EpicSlim) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *EpicSlim) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *EpicSlim) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *EpicSlim) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *EpicSlim) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetName

`func (o *EpicSlim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EpicSlim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EpicSlim) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *EpicSlim) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *EpicSlim) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *EpicSlim) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetCompleted

`func (o *EpicSlim) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *EpicSlim) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *EpicSlim) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetProductboardUrl

`func (o *EpicSlim) GetProductboardUrl() string`

GetProductboardUrl returns the ProductboardUrl field if non-nil, zero value otherwise.

### GetProductboardUrlOk

`func (o *EpicSlim) GetProductboardUrlOk() (*string, bool)`

GetProductboardUrlOk returns a tuple with the ProductboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductboardUrl

`func (o *EpicSlim) SetProductboardUrl(v string)`

SetProductboardUrl sets ProductboardUrl field to given value.


### SetProductboardUrlNil

`func (o *EpicSlim) SetProductboardUrlNil(b bool)`

 SetProductboardUrlNil sets the value for ProductboardUrl to be an explicit nil

### UnsetProductboardUrl
`func (o *EpicSlim) UnsetProductboardUrl()`

UnsetProductboardUrl ensures that no value is present for ProductboardUrl, not even an explicit nil
### GetPlannedStartDate

`func (o *EpicSlim) GetPlannedStartDate() time.Time`

GetPlannedStartDate returns the PlannedStartDate field if non-nil, zero value otherwise.

### GetPlannedStartDateOk

`func (o *EpicSlim) GetPlannedStartDateOk() (*time.Time, bool)`

GetPlannedStartDateOk returns a tuple with the PlannedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartDate

`func (o *EpicSlim) SetPlannedStartDate(v time.Time)`

SetPlannedStartDate sets PlannedStartDate field to given value.


### SetPlannedStartDateNil

`func (o *EpicSlim) SetPlannedStartDateNil(b bool)`

 SetPlannedStartDateNil sets the value for PlannedStartDate to be an explicit nil

### UnsetPlannedStartDate
`func (o *EpicSlim) UnsetPlannedStartDate()`

UnsetPlannedStartDate ensures that no value is present for PlannedStartDate, not even an explicit nil
### GetState

`func (o *EpicSlim) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EpicSlim) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EpicSlim) SetState(v string)`

SetState sets State field to given value.


### GetMilestoneId

`func (o *EpicSlim) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *EpicSlim) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *EpicSlim) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.


### SetMilestoneIdNil

`func (o *EpicSlim) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *EpicSlim) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetRequestedById

`func (o *EpicSlim) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *EpicSlim) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *EpicSlim) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.


### GetEpicStateId

`func (o *EpicSlim) GetEpicStateId() int64`

GetEpicStateId returns the EpicStateId field if non-nil, zero value otherwise.

### GetEpicStateIdOk

`func (o *EpicSlim) GetEpicStateIdOk() (*int64, bool)`

GetEpicStateIdOk returns a tuple with the EpicStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicStateId

`func (o *EpicSlim) SetEpicStateId(v int64)`

SetEpicStateId sets EpicStateId field to given value.


### GetLabelIds

`func (o *EpicSlim) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *EpicSlim) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *EpicSlim) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.


### GetStartedAtOverride

`func (o *EpicSlim) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *EpicSlim) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *EpicSlim) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.


### SetStartedAtOverrideNil

`func (o *EpicSlim) SetStartedAtOverrideNil(b bool)`

 SetStartedAtOverrideNil sets the value for StartedAtOverride to be an explicit nil

### UnsetStartedAtOverride
`func (o *EpicSlim) UnsetStartedAtOverride()`

UnsetStartedAtOverride ensures that no value is present for StartedAtOverride, not even an explicit nil
### GetGroupId

`func (o *EpicSlim) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *EpicSlim) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *EpicSlim) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### SetGroupIdNil

`func (o *EpicSlim) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *EpicSlim) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetUpdatedAt

`func (o *EpicSlim) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EpicSlim) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EpicSlim) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *EpicSlim) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *EpicSlim) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetGroupMentionIds

`func (o *EpicSlim) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *EpicSlim) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *EpicSlim) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetProductboardId

`func (o *EpicSlim) GetProductboardId() string`

GetProductboardId returns the ProductboardId field if non-nil, zero value otherwise.

### GetProductboardIdOk

`func (o *EpicSlim) GetProductboardIdOk() (*string, bool)`

GetProductboardIdOk returns a tuple with the ProductboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductboardId

`func (o *EpicSlim) SetProductboardId(v string)`

SetProductboardId sets ProductboardId field to given value.


### SetProductboardIdNil

`func (o *EpicSlim) SetProductboardIdNil(b bool)`

 SetProductboardIdNil sets the value for ProductboardId to be an explicit nil

### UnsetProductboardId
`func (o *EpicSlim) UnsetProductboardId()`

UnsetProductboardId ensures that no value is present for ProductboardId, not even an explicit nil
### GetFollowerIds

`func (o *EpicSlim) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *EpicSlim) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *EpicSlim) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.


### GetOwnerIds

`func (o *EpicSlim) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *EpicSlim) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *EpicSlim) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.


### GetExternalId

`func (o *EpicSlim) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EpicSlim) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EpicSlim) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *EpicSlim) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *EpicSlim) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetId

`func (o *EpicSlim) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EpicSlim) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EpicSlim) SetId(v int64)`

SetId sets Id field to given value.


### GetPosition

`func (o *EpicSlim) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EpicSlim) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EpicSlim) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetProductboardName

`func (o *EpicSlim) GetProductboardName() string`

GetProductboardName returns the ProductboardName field if non-nil, zero value otherwise.

### GetProductboardNameOk

`func (o *EpicSlim) GetProductboardNameOk() (*string, bool)`

GetProductboardNameOk returns a tuple with the ProductboardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductboardName

`func (o *EpicSlim) SetProductboardName(v string)`

SetProductboardName sets ProductboardName field to given value.


### SetProductboardNameNil

`func (o *EpicSlim) SetProductboardNameNil(b bool)`

 SetProductboardNameNil sets the value for ProductboardName to be an explicit nil

### UnsetProductboardName
`func (o *EpicSlim) UnsetProductboardName()`

UnsetProductboardName ensures that no value is present for ProductboardName, not even an explicit nil
### GetDeadline

`func (o *EpicSlim) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *EpicSlim) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *EpicSlim) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### SetDeadlineNil

`func (o *EpicSlim) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *EpicSlim) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetStats

`func (o *EpicSlim) GetStats() EpicStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *EpicSlim) GetStatsOk() (*EpicStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *EpicSlim) SetStats(v EpicStats)`

SetStats sets Stats field to given value.


### GetCreatedAt

`func (o *EpicSlim) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EpicSlim) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EpicSlim) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *EpicSlim) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *EpicSlim) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


