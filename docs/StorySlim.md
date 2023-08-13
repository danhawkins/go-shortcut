# StorySlim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Story. | 
**Description** | Pointer to **string** | The description of the Story. | [optional] 
**Archived** | **bool** | True if the story has been archived or not. | 
**Started** | **bool** | A true/false boolean indicating if the Story has been started. | 
**StoryLinks** | [**[]TypedStoryLink**](TypedStoryLink.md) | An array of story links attached to the Story. | 
**EntityType** | **string** | A string description of this resource. | 
**Labels** | [**[]LabelSlim**](LabelSlim.md) | An array of labels attached to the story. | 
**TaskIds** | **[]int64** | An array of IDs of Tasks attached to the story. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**SyncedItem** | Pointer to [**SyncedItem**](SyncedItem.md) |  | [optional] 
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Story description. | 
**StoryType** | **string** | The type of story (feature, bug, chore). | 
**CustomFields** | Pointer to [**[]StoryCustomField**](StoryCustomField.md) | An array of CustomField value assertions for the story. | [optional] 
**FileIds** | **[]int64** | An array of IDs of Files attached to the story. | 
**NumTasksCompleted** | **int64** | The number of tasks on the story which are complete. | 
**WorkflowId** | **int64** | The ID of the workflow the story belongs to. | 
**CompletedAtOverride** | **NullableTime** | A manual override for the time/date the Story was completed. | 
**StartedAt** | **NullableTime** | The time/date the Story was started. | 
**CompletedAt** | **NullableTime** | The time/date the Story was completed. | 
**Name** | **string** | The name of the story. | 
**GlobalId** | **string** |  | 
**Completed** | **bool** | A true/false boolean indicating if the Story has been completed. | 
**Blocker** | **bool** | A true/false boolean indicating if the Story is currently a blocker of another story. | 
**EpicId** | **NullableInt64** | The ID of the epic the story belongs to. | 
**UnresolvedBlockerComments** | Pointer to **[]int64** | The IDs of any unresolved blocker comments on the Story. | [optional] 
**StoryTemplateId** | **NullableString** | The ID of the story template used to create this story, or null if not created using a template. | 
**ExternalLinks** | **[]string** | An array of external links (strings) associated with a Story | 
**PreviousIterationIds** | **[]int64** | The IDs of the iteration the story belongs to. | 
**RequestedById** | **string** | The ID of the Member that requested the story. | 
**IterationId** | **NullableInt64** | The ID of the iteration the story belongs to. | 
**LabelIds** | **[]int64** | An array of label ids attached to the story. | 
**StartedAtOverride** | **NullableTime** | A manual override for the time/date the Story was started. | 
**GroupId** | **NullableString** | The ID of the group associated with the story. | 
**WorkflowStateId** | **int64** | The ID of the workflow state the story is currently in. | 
**UpdatedAt** | **NullableTime** | The time/date the Story was updated. | 
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Story description. | 
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | 
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | 
**ExternalId** | **NullableString** | This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here. | 
**Id** | **int64** | The unique ID of the Story. | 
**LeadTime** | Pointer to **int64** | The lead time (in seconds) of this story when complete. | [optional] 
**Estimate** | **NullableInt64** | The numeric point estimate of the story. Can also be null, which means unestimated. | 
**Position** | **int64** | A number representing the position of the story in relation to every other story in the current project. | 
**Blocked** | **bool** | A true/false boolean indicating if the Story is currently blocked. | 
**ProjectId** | **NullableInt64** | The ID of the project the story belongs to. | 
**LinkedFileIds** | **[]int64** | An array of IDs of LinkedFiles attached to the story. | 
**Deadline** | **NullableTime** | The due date of the story. | 
**Stats** | [**StoryStats**](StoryStats.md) |  | 
**CommentIds** | **[]int64** | An array of IDs of Comments attached to the story. | 
**CycleTime** | Pointer to **int64** | The cycle time (in seconds) of this story when complete. | [optional] 
**CreatedAt** | **time.Time** | The time/date the Story was created. | 
**MovedAt** | **NullableTime** | The time/date the Story was last changed workflow-state. | 

## Methods

### NewStorySlim

`func NewStorySlim(appUrl string, archived bool, started bool, storyLinks []TypedStoryLink, entityType string, labels []LabelSlim, taskIds []int64, mentionIds []string, memberMentionIds []string, storyType string, fileIds []int64, numTasksCompleted int64, workflowId int64, completedAtOverride NullableTime, startedAt NullableTime, completedAt NullableTime, name string, globalId string, completed bool, blocker bool, epicId NullableInt64, storyTemplateId NullableString, externalLinks []string, previousIterationIds []int64, requestedById string, iterationId NullableInt64, labelIds []int64, startedAtOverride NullableTime, groupId NullableString, workflowStateId int64, updatedAt NullableTime, groupMentionIds []string, followerIds []string, ownerIds []string, externalId NullableString, id int64, estimate NullableInt64, position int64, blocked bool, projectId NullableInt64, linkedFileIds []int64, deadline NullableTime, stats StoryStats, commentIds []int64, createdAt time.Time, movedAt NullableTime, ) *StorySlim`

NewStorySlim instantiates a new StorySlim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorySlimWithDefaults

`func NewStorySlimWithDefaults() *StorySlim`

NewStorySlimWithDefaults instantiates a new StorySlim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *StorySlim) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *StorySlim) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *StorySlim) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetDescription

`func (o *StorySlim) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorySlim) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorySlim) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorySlim) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *StorySlim) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *StorySlim) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *StorySlim) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetStarted

`func (o *StorySlim) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *StorySlim) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *StorySlim) SetStarted(v bool)`

SetStarted sets Started field to given value.


### GetStoryLinks

`func (o *StorySlim) GetStoryLinks() []TypedStoryLink`

GetStoryLinks returns the StoryLinks field if non-nil, zero value otherwise.

### GetStoryLinksOk

`func (o *StorySlim) GetStoryLinksOk() (*[]TypedStoryLink, bool)`

GetStoryLinksOk returns a tuple with the StoryLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryLinks

`func (o *StorySlim) SetStoryLinks(v []TypedStoryLink)`

SetStoryLinks sets StoryLinks field to given value.


### GetEntityType

`func (o *StorySlim) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *StorySlim) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *StorySlim) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetLabels

`func (o *StorySlim) GetLabels() []LabelSlim`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StorySlim) GetLabelsOk() (*[]LabelSlim, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StorySlim) SetLabels(v []LabelSlim)`

SetLabels sets Labels field to given value.


### GetTaskIds

`func (o *StorySlim) GetTaskIds() []int64`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *StorySlim) GetTaskIdsOk() (*[]int64, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *StorySlim) SetTaskIds(v []int64)`

SetTaskIds sets TaskIds field to given value.


### GetMentionIds

`func (o *StorySlim) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *StorySlim) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *StorySlim) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetSyncedItem

`func (o *StorySlim) GetSyncedItem() SyncedItem`

GetSyncedItem returns the SyncedItem field if non-nil, zero value otherwise.

### GetSyncedItemOk

`func (o *StorySlim) GetSyncedItemOk() (*SyncedItem, bool)`

GetSyncedItemOk returns a tuple with the SyncedItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedItem

`func (o *StorySlim) SetSyncedItem(v SyncedItem)`

SetSyncedItem sets SyncedItem field to given value.

### HasSyncedItem

`func (o *StorySlim) HasSyncedItem() bool`

HasSyncedItem returns a boolean if a field has been set.

### GetMemberMentionIds

`func (o *StorySlim) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *StorySlim) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *StorySlim) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetStoryType

`func (o *StorySlim) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *StorySlim) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *StorySlim) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.


### GetCustomFields

`func (o *StorySlim) GetCustomFields() []StoryCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *StorySlim) GetCustomFieldsOk() (*[]StoryCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *StorySlim) SetCustomFields(v []StoryCustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *StorySlim) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetFileIds

`func (o *StorySlim) GetFileIds() []int64`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *StorySlim) GetFileIdsOk() (*[]int64, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *StorySlim) SetFileIds(v []int64)`

SetFileIds sets FileIds field to given value.


### GetNumTasksCompleted

`func (o *StorySlim) GetNumTasksCompleted() int64`

GetNumTasksCompleted returns the NumTasksCompleted field if non-nil, zero value otherwise.

### GetNumTasksCompletedOk

`func (o *StorySlim) GetNumTasksCompletedOk() (*int64, bool)`

GetNumTasksCompletedOk returns a tuple with the NumTasksCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTasksCompleted

`func (o *StorySlim) SetNumTasksCompleted(v int64)`

SetNumTasksCompleted sets NumTasksCompleted field to given value.


### GetWorkflowId

`func (o *StorySlim) GetWorkflowId() int64`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *StorySlim) GetWorkflowIdOk() (*int64, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *StorySlim) SetWorkflowId(v int64)`

SetWorkflowId sets WorkflowId field to given value.


### GetCompletedAtOverride

`func (o *StorySlim) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *StorySlim) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *StorySlim) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.


### SetCompletedAtOverrideNil

`func (o *StorySlim) SetCompletedAtOverrideNil(b bool)`

 SetCompletedAtOverrideNil sets the value for CompletedAtOverride to be an explicit nil

### UnsetCompletedAtOverride
`func (o *StorySlim) UnsetCompletedAtOverride()`

UnsetCompletedAtOverride ensures that no value is present for CompletedAtOverride, not even an explicit nil
### GetStartedAt

`func (o *StorySlim) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *StorySlim) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *StorySlim) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *StorySlim) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *StorySlim) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *StorySlim) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *StorySlim) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *StorySlim) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *StorySlim) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *StorySlim) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetName

`func (o *StorySlim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorySlim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorySlim) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *StorySlim) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *StorySlim) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *StorySlim) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetCompleted

`func (o *StorySlim) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *StorySlim) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *StorySlim) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetBlocker

`func (o *StorySlim) GetBlocker() bool`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *StorySlim) GetBlockerOk() (*bool, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *StorySlim) SetBlocker(v bool)`

SetBlocker sets Blocker field to given value.


### GetEpicId

`func (o *StorySlim) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *StorySlim) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *StorySlim) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.


### SetEpicIdNil

`func (o *StorySlim) SetEpicIdNil(b bool)`

 SetEpicIdNil sets the value for EpicId to be an explicit nil

### UnsetEpicId
`func (o *StorySlim) UnsetEpicId()`

UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
### GetUnresolvedBlockerComments

`func (o *StorySlim) GetUnresolvedBlockerComments() []int64`

GetUnresolvedBlockerComments returns the UnresolvedBlockerComments field if non-nil, zero value otherwise.

### GetUnresolvedBlockerCommentsOk

`func (o *StorySlim) GetUnresolvedBlockerCommentsOk() (*[]int64, bool)`

GetUnresolvedBlockerCommentsOk returns a tuple with the UnresolvedBlockerComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedBlockerComments

`func (o *StorySlim) SetUnresolvedBlockerComments(v []int64)`

SetUnresolvedBlockerComments sets UnresolvedBlockerComments field to given value.

### HasUnresolvedBlockerComments

`func (o *StorySlim) HasUnresolvedBlockerComments() bool`

HasUnresolvedBlockerComments returns a boolean if a field has been set.

### GetStoryTemplateId

`func (o *StorySlim) GetStoryTemplateId() string`

GetStoryTemplateId returns the StoryTemplateId field if non-nil, zero value otherwise.

### GetStoryTemplateIdOk

`func (o *StorySlim) GetStoryTemplateIdOk() (*string, bool)`

GetStoryTemplateIdOk returns a tuple with the StoryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryTemplateId

`func (o *StorySlim) SetStoryTemplateId(v string)`

SetStoryTemplateId sets StoryTemplateId field to given value.


### SetStoryTemplateIdNil

`func (o *StorySlim) SetStoryTemplateIdNil(b bool)`

 SetStoryTemplateIdNil sets the value for StoryTemplateId to be an explicit nil

### UnsetStoryTemplateId
`func (o *StorySlim) UnsetStoryTemplateId()`

UnsetStoryTemplateId ensures that no value is present for StoryTemplateId, not even an explicit nil
### GetExternalLinks

`func (o *StorySlim) GetExternalLinks() []string`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *StorySlim) GetExternalLinksOk() (*[]string, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *StorySlim) SetExternalLinks(v []string)`

SetExternalLinks sets ExternalLinks field to given value.


### GetPreviousIterationIds

`func (o *StorySlim) GetPreviousIterationIds() []int64`

GetPreviousIterationIds returns the PreviousIterationIds field if non-nil, zero value otherwise.

### GetPreviousIterationIdsOk

`func (o *StorySlim) GetPreviousIterationIdsOk() (*[]int64, bool)`

GetPreviousIterationIdsOk returns a tuple with the PreviousIterationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIterationIds

`func (o *StorySlim) SetPreviousIterationIds(v []int64)`

SetPreviousIterationIds sets PreviousIterationIds field to given value.


### GetRequestedById

`func (o *StorySlim) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *StorySlim) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *StorySlim) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.


### GetIterationId

`func (o *StorySlim) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *StorySlim) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *StorySlim) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.


### SetIterationIdNil

`func (o *StorySlim) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *StorySlim) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetLabelIds

`func (o *StorySlim) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *StorySlim) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *StorySlim) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.


### GetStartedAtOverride

`func (o *StorySlim) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *StorySlim) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *StorySlim) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.


### SetStartedAtOverrideNil

`func (o *StorySlim) SetStartedAtOverrideNil(b bool)`

 SetStartedAtOverrideNil sets the value for StartedAtOverride to be an explicit nil

### UnsetStartedAtOverride
`func (o *StorySlim) UnsetStartedAtOverride()`

UnsetStartedAtOverride ensures that no value is present for StartedAtOverride, not even an explicit nil
### GetGroupId

`func (o *StorySlim) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StorySlim) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StorySlim) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### SetGroupIdNil

`func (o *StorySlim) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *StorySlim) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetWorkflowStateId

`func (o *StorySlim) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *StorySlim) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *StorySlim) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.


### GetUpdatedAt

`func (o *StorySlim) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StorySlim) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StorySlim) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *StorySlim) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *StorySlim) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetGroupMentionIds

`func (o *StorySlim) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *StorySlim) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *StorySlim) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetFollowerIds

`func (o *StorySlim) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *StorySlim) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *StorySlim) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.


### GetOwnerIds

`func (o *StorySlim) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *StorySlim) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *StorySlim) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.


### GetExternalId

`func (o *StorySlim) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *StorySlim) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *StorySlim) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *StorySlim) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *StorySlim) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetId

`func (o *StorySlim) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorySlim) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorySlim) SetId(v int64)`

SetId sets Id field to given value.


### GetLeadTime

`func (o *StorySlim) GetLeadTime() int64`

GetLeadTime returns the LeadTime field if non-nil, zero value otherwise.

### GetLeadTimeOk

`func (o *StorySlim) GetLeadTimeOk() (*int64, bool)`

GetLeadTimeOk returns a tuple with the LeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadTime

`func (o *StorySlim) SetLeadTime(v int64)`

SetLeadTime sets LeadTime field to given value.

### HasLeadTime

`func (o *StorySlim) HasLeadTime() bool`

HasLeadTime returns a boolean if a field has been set.

### GetEstimate

`func (o *StorySlim) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *StorySlim) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *StorySlim) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.


### SetEstimateNil

`func (o *StorySlim) SetEstimateNil(b bool)`

 SetEstimateNil sets the value for Estimate to be an explicit nil

### UnsetEstimate
`func (o *StorySlim) UnsetEstimate()`

UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
### GetPosition

`func (o *StorySlim) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *StorySlim) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *StorySlim) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetBlocked

`func (o *StorySlim) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *StorySlim) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *StorySlim) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetProjectId

`func (o *StorySlim) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *StorySlim) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *StorySlim) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *StorySlim) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *StorySlim) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLinkedFileIds

`func (o *StorySlim) GetLinkedFileIds() []int64`

GetLinkedFileIds returns the LinkedFileIds field if non-nil, zero value otherwise.

### GetLinkedFileIdsOk

`func (o *StorySlim) GetLinkedFileIdsOk() (*[]int64, bool)`

GetLinkedFileIdsOk returns a tuple with the LinkedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFileIds

`func (o *StorySlim) SetLinkedFileIds(v []int64)`

SetLinkedFileIds sets LinkedFileIds field to given value.


### GetDeadline

`func (o *StorySlim) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *StorySlim) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *StorySlim) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### SetDeadlineNil

`func (o *StorySlim) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *StorySlim) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetStats

`func (o *StorySlim) GetStats() StoryStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StorySlim) GetStatsOk() (*StoryStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StorySlim) SetStats(v StoryStats)`

SetStats sets Stats field to given value.


### GetCommentIds

`func (o *StorySlim) GetCommentIds() []int64`

GetCommentIds returns the CommentIds field if non-nil, zero value otherwise.

### GetCommentIdsOk

`func (o *StorySlim) GetCommentIdsOk() (*[]int64, bool)`

GetCommentIdsOk returns a tuple with the CommentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentIds

`func (o *StorySlim) SetCommentIds(v []int64)`

SetCommentIds sets CommentIds field to given value.


### GetCycleTime

`func (o *StorySlim) GetCycleTime() int64`

GetCycleTime returns the CycleTime field if non-nil, zero value otherwise.

### GetCycleTimeOk

`func (o *StorySlim) GetCycleTimeOk() (*int64, bool)`

GetCycleTimeOk returns a tuple with the CycleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleTime

`func (o *StorySlim) SetCycleTime(v int64)`

SetCycleTime sets CycleTime field to given value.

### HasCycleTime

`func (o *StorySlim) HasCycleTime() bool`

HasCycleTime returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StorySlim) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StorySlim) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StorySlim) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMovedAt

`func (o *StorySlim) GetMovedAt() time.Time`

GetMovedAt returns the MovedAt field if non-nil, zero value otherwise.

### GetMovedAtOk

`func (o *StorySlim) GetMovedAtOk() (*time.Time, bool)`

GetMovedAtOk returns a tuple with the MovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedAt

`func (o *StorySlim) SetMovedAt(v time.Time)`

SetMovedAt sets MovedAt field to given value.


### SetMovedAtNil

`func (o *StorySlim) SetMovedAtNil(b bool)`

 SetMovedAtNil sets the value for MovedAt to be an explicit nil

### UnsetMovedAt
`func (o *StorySlim) UnsetMovedAt()`

UnsetMovedAt ensures that no value is present for MovedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


