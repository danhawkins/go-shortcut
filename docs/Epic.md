# Epic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Epic. | 
**Description** | **string** | The Epic&#39;s description. | 
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
**Comments** | [**[]ThreadedComment**](ThreadedComment.md) | A nested array of threaded comments. | 
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

### NewEpic

`func NewEpic(appUrl string, description string, archived bool, started bool, entityType string, labels []LabelSlim, mentionIds []string, memberMentionIds []string, associatedGroups []EpicAssociatedGroup, projectIds []int64, storiesWithoutProjects int64, completedAtOverride NullableTime, productboardPluginId NullableString, startedAt NullableTime, completedAt NullableTime, name string, globalId string, completed bool, comments []ThreadedComment, productboardUrl NullableString, plannedStartDate NullableTime, state string, milestoneId NullableInt64, requestedById string, epicStateId int64, labelIds []int64, startedAtOverride NullableTime, groupId NullableString, updatedAt NullableTime, groupMentionIds []string, productboardId NullableString, followerIds []string, ownerIds []string, externalId NullableString, id int64, position int64, productboardName NullableString, deadline NullableTime, stats EpicStats, createdAt NullableTime, ) *Epic`

NewEpic instantiates a new Epic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpicWithDefaults

`func NewEpicWithDefaults() *Epic`

NewEpicWithDefaults instantiates a new Epic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *Epic) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *Epic) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *Epic) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetDescription

`func (o *Epic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Epic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Epic) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *Epic) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Epic) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Epic) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetStarted

`func (o *Epic) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *Epic) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *Epic) SetStarted(v bool)`

SetStarted sets Started field to given value.


### GetEntityType

`func (o *Epic) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Epic) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Epic) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetLabels

`func (o *Epic) GetLabels() []LabelSlim`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Epic) GetLabelsOk() (*[]LabelSlim, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Epic) SetLabels(v []LabelSlim)`

SetLabels sets Labels field to given value.


### GetMentionIds

`func (o *Epic) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *Epic) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *Epic) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetMemberMentionIds

`func (o *Epic) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *Epic) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *Epic) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetAssociatedGroups

`func (o *Epic) GetAssociatedGroups() []EpicAssociatedGroup`

GetAssociatedGroups returns the AssociatedGroups field if non-nil, zero value otherwise.

### GetAssociatedGroupsOk

`func (o *Epic) GetAssociatedGroupsOk() (*[]EpicAssociatedGroup, bool)`

GetAssociatedGroupsOk returns a tuple with the AssociatedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedGroups

`func (o *Epic) SetAssociatedGroups(v []EpicAssociatedGroup)`

SetAssociatedGroups sets AssociatedGroups field to given value.


### GetProjectIds

`func (o *Epic) GetProjectIds() []int64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *Epic) GetProjectIdsOk() (*[]int64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *Epic) SetProjectIds(v []int64)`

SetProjectIds sets ProjectIds field to given value.


### GetStoriesWithoutProjects

`func (o *Epic) GetStoriesWithoutProjects() int64`

GetStoriesWithoutProjects returns the StoriesWithoutProjects field if non-nil, zero value otherwise.

### GetStoriesWithoutProjectsOk

`func (o *Epic) GetStoriesWithoutProjectsOk() (*int64, bool)`

GetStoriesWithoutProjectsOk returns a tuple with the StoriesWithoutProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoriesWithoutProjects

`func (o *Epic) SetStoriesWithoutProjects(v int64)`

SetStoriesWithoutProjects sets StoriesWithoutProjects field to given value.


### GetCompletedAtOverride

`func (o *Epic) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *Epic) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *Epic) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.


### SetCompletedAtOverrideNil

`func (o *Epic) SetCompletedAtOverrideNil(b bool)`

 SetCompletedAtOverrideNil sets the value for CompletedAtOverride to be an explicit nil

### UnsetCompletedAtOverride
`func (o *Epic) UnsetCompletedAtOverride()`

UnsetCompletedAtOverride ensures that no value is present for CompletedAtOverride, not even an explicit nil
### GetProductboardPluginId

`func (o *Epic) GetProductboardPluginId() string`

GetProductboardPluginId returns the ProductboardPluginId field if non-nil, zero value otherwise.

### GetProductboardPluginIdOk

`func (o *Epic) GetProductboardPluginIdOk() (*string, bool)`

GetProductboardPluginIdOk returns a tuple with the ProductboardPluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductboardPluginId

`func (o *Epic) SetProductboardPluginId(v string)`

SetProductboardPluginId sets ProductboardPluginId field to given value.


### SetProductboardPluginIdNil

`func (o *Epic) SetProductboardPluginIdNil(b bool)`

 SetProductboardPluginIdNil sets the value for ProductboardPluginId to be an explicit nil

### UnsetProductboardPluginId
`func (o *Epic) UnsetProductboardPluginId()`

UnsetProductboardPluginId ensures that no value is present for ProductboardPluginId, not even an explicit nil
### GetStartedAt

`func (o *Epic) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Epic) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Epic) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *Epic) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Epic) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *Epic) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Epic) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Epic) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *Epic) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Epic) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetName

`func (o *Epic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Epic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Epic) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *Epic) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *Epic) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *Epic) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetCompleted

`func (o *Epic) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *Epic) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *Epic) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetComments

`func (o *Epic) GetComments() []ThreadedComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Epic) GetCommentsOk() (*[]ThreadedComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Epic) SetComments(v []ThreadedComment)`

SetComments sets Comments field to given value.


### GetProductboardUrl

`func (o *Epic) GetProductboardUrl() string`

GetProductboardUrl returns the ProductboardUrl field if non-nil, zero value otherwise.

### GetProductboardUrlOk

`func (o *Epic) GetProductboardUrlOk() (*string, bool)`

GetProductboardUrlOk returns a tuple with the ProductboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductboardUrl

`func (o *Epic) SetProductboardUrl(v string)`

SetProductboardUrl sets ProductboardUrl field to given value.


### SetProductboardUrlNil

`func (o *Epic) SetProductboardUrlNil(b bool)`

 SetProductboardUrlNil sets the value for ProductboardUrl to be an explicit nil

### UnsetProductboardUrl
`func (o *Epic) UnsetProductboardUrl()`

UnsetProductboardUrl ensures that no value is present for ProductboardUrl, not even an explicit nil
### GetPlannedStartDate

`func (o *Epic) GetPlannedStartDate() time.Time`

GetPlannedStartDate returns the PlannedStartDate field if non-nil, zero value otherwise.

### GetPlannedStartDateOk

`func (o *Epic) GetPlannedStartDateOk() (*time.Time, bool)`

GetPlannedStartDateOk returns a tuple with the PlannedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartDate

`func (o *Epic) SetPlannedStartDate(v time.Time)`

SetPlannedStartDate sets PlannedStartDate field to given value.


### SetPlannedStartDateNil

`func (o *Epic) SetPlannedStartDateNil(b bool)`

 SetPlannedStartDateNil sets the value for PlannedStartDate to be an explicit nil

### UnsetPlannedStartDate
`func (o *Epic) UnsetPlannedStartDate()`

UnsetPlannedStartDate ensures that no value is present for PlannedStartDate, not even an explicit nil
### GetState

`func (o *Epic) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Epic) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Epic) SetState(v string)`

SetState sets State field to given value.


### GetMilestoneId

`func (o *Epic) GetMilestoneId() int64`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *Epic) GetMilestoneIdOk() (*int64, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *Epic) SetMilestoneId(v int64)`

SetMilestoneId sets MilestoneId field to given value.


### SetMilestoneIdNil

`func (o *Epic) SetMilestoneIdNil(b bool)`

 SetMilestoneIdNil sets the value for MilestoneId to be an explicit nil

### UnsetMilestoneId
`func (o *Epic) UnsetMilestoneId()`

UnsetMilestoneId ensures that no value is present for MilestoneId, not even an explicit nil
### GetRequestedById

`func (o *Epic) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *Epic) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *Epic) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.


### GetEpicStateId

`func (o *Epic) GetEpicStateId() int64`

GetEpicStateId returns the EpicStateId field if non-nil, zero value otherwise.

### GetEpicStateIdOk

`func (o *Epic) GetEpicStateIdOk() (*int64, bool)`

GetEpicStateIdOk returns a tuple with the EpicStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicStateId

`func (o *Epic) SetEpicStateId(v int64)`

SetEpicStateId sets EpicStateId field to given value.


### GetLabelIds

`func (o *Epic) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *Epic) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *Epic) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.


### GetStartedAtOverride

`func (o *Epic) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *Epic) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *Epic) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.


### SetStartedAtOverrideNil

`func (o *Epic) SetStartedAtOverrideNil(b bool)`

 SetStartedAtOverrideNil sets the value for StartedAtOverride to be an explicit nil

### UnsetStartedAtOverride
`func (o *Epic) UnsetStartedAtOverride()`

UnsetStartedAtOverride ensures that no value is present for StartedAtOverride, not even an explicit nil
### GetGroupId

`func (o *Epic) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Epic) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Epic) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### SetGroupIdNil

`func (o *Epic) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *Epic) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetUpdatedAt

`func (o *Epic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Epic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Epic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Epic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Epic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetGroupMentionIds

`func (o *Epic) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *Epic) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *Epic) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetProductboardId

`func (o *Epic) GetProductboardId() string`

GetProductboardId returns the ProductboardId field if non-nil, zero value otherwise.

### GetProductboardIdOk

`func (o *Epic) GetProductboardIdOk() (*string, bool)`

GetProductboardIdOk returns a tuple with the ProductboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductboardId

`func (o *Epic) SetProductboardId(v string)`

SetProductboardId sets ProductboardId field to given value.


### SetProductboardIdNil

`func (o *Epic) SetProductboardIdNil(b bool)`

 SetProductboardIdNil sets the value for ProductboardId to be an explicit nil

### UnsetProductboardId
`func (o *Epic) UnsetProductboardId()`

UnsetProductboardId ensures that no value is present for ProductboardId, not even an explicit nil
### GetFollowerIds

`func (o *Epic) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *Epic) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *Epic) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.


### GetOwnerIds

`func (o *Epic) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *Epic) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *Epic) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.


### GetExternalId

`func (o *Epic) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Epic) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Epic) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *Epic) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Epic) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetId

`func (o *Epic) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Epic) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Epic) SetId(v int64)`

SetId sets Id field to given value.


### GetPosition

`func (o *Epic) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Epic) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Epic) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetProductboardName

`func (o *Epic) GetProductboardName() string`

GetProductboardName returns the ProductboardName field if non-nil, zero value otherwise.

### GetProductboardNameOk

`func (o *Epic) GetProductboardNameOk() (*string, bool)`

GetProductboardNameOk returns a tuple with the ProductboardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductboardName

`func (o *Epic) SetProductboardName(v string)`

SetProductboardName sets ProductboardName field to given value.


### SetProductboardNameNil

`func (o *Epic) SetProductboardNameNil(b bool)`

 SetProductboardNameNil sets the value for ProductboardName to be an explicit nil

### UnsetProductboardName
`func (o *Epic) UnsetProductboardName()`

UnsetProductboardName ensures that no value is present for ProductboardName, not even an explicit nil
### GetDeadline

`func (o *Epic) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *Epic) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *Epic) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### SetDeadlineNil

`func (o *Epic) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *Epic) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetStats

`func (o *Epic) GetStats() EpicStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Epic) GetStatsOk() (*EpicStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Epic) SetStats(v EpicStats)`

SetStats sets Stats field to given value.


### GetCreatedAt

`func (o *Epic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Epic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Epic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Epic) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Epic) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


