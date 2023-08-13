# Story

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Story. | 
**Description** | **string** | The description of the story. | 
**Archived** | **bool** | True if the story has been archived or not. | 
**Started** | **bool** | A true/false boolean indicating if the Story has been started. | 
**StoryLinks** | [**[]TypedStoryLink**](TypedStoryLink.md) | An array of story links attached to the Story. | 
**EntityType** | **string** | A string description of this resource. | 
**Labels** | [**[]LabelSlim**](LabelSlim.md) | An array of labels attached to the story. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**SyncedItem** | Pointer to [**SyncedItem**](SyncedItem.md) |  | [optional] 
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Story description. | 
**StoryType** | **string** | The type of story (feature, bug, chore). | 
**CustomFields** | Pointer to [**[]StoryCustomField**](StoryCustomField.md) | An array of CustomField value assertions for the story. | [optional] 
**LinkedFiles** | [**[]LinkedFile**](LinkedFile.md) | An array of linked files attached to the story. | 
**WorkflowId** | **int64** | The ID of the workflow the story belongs to. | 
**CompletedAtOverride** | **NullableTime** | A manual override for the time/date the Story was completed. | 
**StartedAt** | **NullableTime** | The time/date the Story was started. | 
**CompletedAt** | **NullableTime** | The time/date the Story was completed. | 
**Name** | **string** | The name of the story. | 
**GlobalId** | **string** |  | 
**Completed** | **bool** | A true/false boolean indicating if the Story has been completed. | 
**Comments** | [**[]StoryComment**](StoryComment.md) | An array of comments attached to the story. | 
**Blocker** | **bool** | A true/false boolean indicating if the Story is currently a blocker of another story. | 
**Branches** | [**[]Branch**](Branch.md) | An array of Git branches attached to the story. | 
**EpicId** | **NullableInt64** | The ID of the epic the story belongs to. | 
**UnresolvedBlockerComments** | Pointer to **[]int64** | The IDs of any unresolved blocker comments on the Story. | [optional] 
**StoryTemplateId** | **NullableString** | The ID of the story template used to create this story, or null if not created using a template. | 
**ExternalLinks** | **[]string** | An array of external links (strings) associated with a Story | 
**PreviousIterationIds** | **[]int64** | The IDs of the iteration the story belongs to. | 
**RequestedById** | **string** | The ID of the Member that requested the story. | 
**IterationId** | **NullableInt64** | The ID of the iteration the story belongs to. | 
**Tasks** | [**[]Task**](Task.md) | An array of tasks connected to the story. | 
**LabelIds** | **[]int64** | An array of label ids attached to the story. | 
**StartedAtOverride** | **NullableTime** | A manual override for the time/date the Story was started. | 
**GroupId** | **NullableString** | The ID of the group associated with the story. | 
**WorkflowStateId** | **int64** | The ID of the workflow state the story is currently in. | 
**UpdatedAt** | **NullableTime** | The time/date the Story was updated. | 
**PullRequests** | [**[]PullRequest**](PullRequest.md) | An array of Pull/Merge Requests attached to the story. | 
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Story description. | 
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | 
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | 
**ExternalId** | **NullableString** | This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here. | 
**Id** | **int64** | The unique ID of the Story. | 
**LeadTime** | Pointer to **int64** | The lead time (in seconds) of this story when complete. | [optional] 
**Estimate** | **NullableInt64** | The numeric point estimate of the story. Can also be null, which means unestimated. | 
**Commits** | [**[]Commit**](Commit.md) | An array of commits attached to the story. | 
**Files** | [**[]UploadedFile**](UploadedFile.md) | An array of files attached to the story. | 
**Position** | **int64** | A number representing the position of the story in relation to every other story in the current project. | 
**Blocked** | **bool** | A true/false boolean indicating if the Story is currently blocked. | 
**ProjectId** | **NullableInt64** | The ID of the project the story belongs to. | 
**Deadline** | **NullableTime** | The due date of the story. | 
**Stats** | [**StoryStats**](StoryStats.md) |  | 
**CycleTime** | Pointer to **int64** | The cycle time (in seconds) of this story when complete. | [optional] 
**CreatedAt** | **time.Time** | The time/date the Story was created. | 
**MovedAt** | **NullableTime** | The time/date the Story was last changed workflow-state. | 

## Methods

### NewStory

`func NewStory(appUrl string, description string, archived bool, started bool, storyLinks []TypedStoryLink, entityType string, labels []LabelSlim, mentionIds []string, memberMentionIds []string, storyType string, linkedFiles []LinkedFile, workflowId int64, completedAtOverride NullableTime, startedAt NullableTime, completedAt NullableTime, name string, globalId string, completed bool, comments []StoryComment, blocker bool, branches []Branch, epicId NullableInt64, storyTemplateId NullableString, externalLinks []string, previousIterationIds []int64, requestedById string, iterationId NullableInt64, tasks []Task, labelIds []int64, startedAtOverride NullableTime, groupId NullableString, workflowStateId int64, updatedAt NullableTime, pullRequests []PullRequest, groupMentionIds []string, followerIds []string, ownerIds []string, externalId NullableString, id int64, estimate NullableInt64, commits []Commit, files []UploadedFile, position int64, blocked bool, projectId NullableInt64, deadline NullableTime, stats StoryStats, createdAt time.Time, movedAt NullableTime, ) *Story`

NewStory instantiates a new Story object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryWithDefaults

`func NewStoryWithDefaults() *Story`

NewStoryWithDefaults instantiates a new Story object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *Story) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *Story) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *Story) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetDescription

`func (o *Story) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Story) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Story) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *Story) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Story) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Story) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetStarted

`func (o *Story) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *Story) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *Story) SetStarted(v bool)`

SetStarted sets Started field to given value.


### GetStoryLinks

`func (o *Story) GetStoryLinks() []TypedStoryLink`

GetStoryLinks returns the StoryLinks field if non-nil, zero value otherwise.

### GetStoryLinksOk

`func (o *Story) GetStoryLinksOk() (*[]TypedStoryLink, bool)`

GetStoryLinksOk returns a tuple with the StoryLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryLinks

`func (o *Story) SetStoryLinks(v []TypedStoryLink)`

SetStoryLinks sets StoryLinks field to given value.


### GetEntityType

`func (o *Story) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Story) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Story) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetLabels

`func (o *Story) GetLabels() []LabelSlim`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Story) GetLabelsOk() (*[]LabelSlim, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Story) SetLabels(v []LabelSlim)`

SetLabels sets Labels field to given value.


### GetMentionIds

`func (o *Story) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *Story) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *Story) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetSyncedItem

`func (o *Story) GetSyncedItem() SyncedItem`

GetSyncedItem returns the SyncedItem field if non-nil, zero value otherwise.

### GetSyncedItemOk

`func (o *Story) GetSyncedItemOk() (*SyncedItem, bool)`

GetSyncedItemOk returns a tuple with the SyncedItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedItem

`func (o *Story) SetSyncedItem(v SyncedItem)`

SetSyncedItem sets SyncedItem field to given value.

### HasSyncedItem

`func (o *Story) HasSyncedItem() bool`

HasSyncedItem returns a boolean if a field has been set.

### GetMemberMentionIds

`func (o *Story) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *Story) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *Story) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetStoryType

`func (o *Story) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *Story) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *Story) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.


### GetCustomFields

`func (o *Story) GetCustomFields() []StoryCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Story) GetCustomFieldsOk() (*[]StoryCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Story) SetCustomFields(v []StoryCustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Story) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetLinkedFiles

`func (o *Story) GetLinkedFiles() []LinkedFile`

GetLinkedFiles returns the LinkedFiles field if non-nil, zero value otherwise.

### GetLinkedFilesOk

`func (o *Story) GetLinkedFilesOk() (*[]LinkedFile, bool)`

GetLinkedFilesOk returns a tuple with the LinkedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFiles

`func (o *Story) SetLinkedFiles(v []LinkedFile)`

SetLinkedFiles sets LinkedFiles field to given value.


### GetWorkflowId

`func (o *Story) GetWorkflowId() int64`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *Story) GetWorkflowIdOk() (*int64, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *Story) SetWorkflowId(v int64)`

SetWorkflowId sets WorkflowId field to given value.


### GetCompletedAtOverride

`func (o *Story) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *Story) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *Story) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.


### SetCompletedAtOverrideNil

`func (o *Story) SetCompletedAtOverrideNil(b bool)`

 SetCompletedAtOverrideNil sets the value for CompletedAtOverride to be an explicit nil

### UnsetCompletedAtOverride
`func (o *Story) UnsetCompletedAtOverride()`

UnsetCompletedAtOverride ensures that no value is present for CompletedAtOverride, not even an explicit nil
### GetStartedAt

`func (o *Story) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Story) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Story) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *Story) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Story) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *Story) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Story) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Story) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *Story) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Story) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetName

`func (o *Story) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Story) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Story) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *Story) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *Story) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *Story) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetCompleted

`func (o *Story) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *Story) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *Story) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetComments

`func (o *Story) GetComments() []StoryComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Story) GetCommentsOk() (*[]StoryComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Story) SetComments(v []StoryComment)`

SetComments sets Comments field to given value.


### GetBlocker

`func (o *Story) GetBlocker() bool`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *Story) GetBlockerOk() (*bool, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *Story) SetBlocker(v bool)`

SetBlocker sets Blocker field to given value.


### GetBranches

`func (o *Story) GetBranches() []Branch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *Story) GetBranchesOk() (*[]Branch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *Story) SetBranches(v []Branch)`

SetBranches sets Branches field to given value.


### GetEpicId

`func (o *Story) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *Story) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *Story) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.


### SetEpicIdNil

`func (o *Story) SetEpicIdNil(b bool)`

 SetEpicIdNil sets the value for EpicId to be an explicit nil

### UnsetEpicId
`func (o *Story) UnsetEpicId()`

UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
### GetUnresolvedBlockerComments

`func (o *Story) GetUnresolvedBlockerComments() []int64`

GetUnresolvedBlockerComments returns the UnresolvedBlockerComments field if non-nil, zero value otherwise.

### GetUnresolvedBlockerCommentsOk

`func (o *Story) GetUnresolvedBlockerCommentsOk() (*[]int64, bool)`

GetUnresolvedBlockerCommentsOk returns a tuple with the UnresolvedBlockerComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedBlockerComments

`func (o *Story) SetUnresolvedBlockerComments(v []int64)`

SetUnresolvedBlockerComments sets UnresolvedBlockerComments field to given value.

### HasUnresolvedBlockerComments

`func (o *Story) HasUnresolvedBlockerComments() bool`

HasUnresolvedBlockerComments returns a boolean if a field has been set.

### GetStoryTemplateId

`func (o *Story) GetStoryTemplateId() string`

GetStoryTemplateId returns the StoryTemplateId field if non-nil, zero value otherwise.

### GetStoryTemplateIdOk

`func (o *Story) GetStoryTemplateIdOk() (*string, bool)`

GetStoryTemplateIdOk returns a tuple with the StoryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryTemplateId

`func (o *Story) SetStoryTemplateId(v string)`

SetStoryTemplateId sets StoryTemplateId field to given value.


### SetStoryTemplateIdNil

`func (o *Story) SetStoryTemplateIdNil(b bool)`

 SetStoryTemplateIdNil sets the value for StoryTemplateId to be an explicit nil

### UnsetStoryTemplateId
`func (o *Story) UnsetStoryTemplateId()`

UnsetStoryTemplateId ensures that no value is present for StoryTemplateId, not even an explicit nil
### GetExternalLinks

`func (o *Story) GetExternalLinks() []string`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *Story) GetExternalLinksOk() (*[]string, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *Story) SetExternalLinks(v []string)`

SetExternalLinks sets ExternalLinks field to given value.


### GetPreviousIterationIds

`func (o *Story) GetPreviousIterationIds() []int64`

GetPreviousIterationIds returns the PreviousIterationIds field if non-nil, zero value otherwise.

### GetPreviousIterationIdsOk

`func (o *Story) GetPreviousIterationIdsOk() (*[]int64, bool)`

GetPreviousIterationIdsOk returns a tuple with the PreviousIterationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIterationIds

`func (o *Story) SetPreviousIterationIds(v []int64)`

SetPreviousIterationIds sets PreviousIterationIds field to given value.


### GetRequestedById

`func (o *Story) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *Story) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *Story) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.


### GetIterationId

`func (o *Story) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *Story) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *Story) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.


### SetIterationIdNil

`func (o *Story) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *Story) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetTasks

`func (o *Story) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *Story) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *Story) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.


### GetLabelIds

`func (o *Story) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *Story) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *Story) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.


### GetStartedAtOverride

`func (o *Story) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *Story) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *Story) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.


### SetStartedAtOverrideNil

`func (o *Story) SetStartedAtOverrideNil(b bool)`

 SetStartedAtOverrideNil sets the value for StartedAtOverride to be an explicit nil

### UnsetStartedAtOverride
`func (o *Story) UnsetStartedAtOverride()`

UnsetStartedAtOverride ensures that no value is present for StartedAtOverride, not even an explicit nil
### GetGroupId

`func (o *Story) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Story) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Story) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### SetGroupIdNil

`func (o *Story) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *Story) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetWorkflowStateId

`func (o *Story) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *Story) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *Story) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.


### GetUpdatedAt

`func (o *Story) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Story) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Story) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Story) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Story) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetPullRequests

`func (o *Story) GetPullRequests() []PullRequest`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *Story) GetPullRequestsOk() (*[]PullRequest, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *Story) SetPullRequests(v []PullRequest)`

SetPullRequests sets PullRequests field to given value.


### GetGroupMentionIds

`func (o *Story) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *Story) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *Story) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetFollowerIds

`func (o *Story) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *Story) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *Story) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.


### GetOwnerIds

`func (o *Story) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *Story) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *Story) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.


### GetExternalId

`func (o *Story) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Story) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Story) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *Story) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Story) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetId

`func (o *Story) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Story) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Story) SetId(v int64)`

SetId sets Id field to given value.


### GetLeadTime

`func (o *Story) GetLeadTime() int64`

GetLeadTime returns the LeadTime field if non-nil, zero value otherwise.

### GetLeadTimeOk

`func (o *Story) GetLeadTimeOk() (*int64, bool)`

GetLeadTimeOk returns a tuple with the LeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadTime

`func (o *Story) SetLeadTime(v int64)`

SetLeadTime sets LeadTime field to given value.

### HasLeadTime

`func (o *Story) HasLeadTime() bool`

HasLeadTime returns a boolean if a field has been set.

### GetEstimate

`func (o *Story) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *Story) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *Story) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.


### SetEstimateNil

`func (o *Story) SetEstimateNil(b bool)`

 SetEstimateNil sets the value for Estimate to be an explicit nil

### UnsetEstimate
`func (o *Story) UnsetEstimate()`

UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
### GetCommits

`func (o *Story) GetCommits() []Commit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *Story) GetCommitsOk() (*[]Commit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *Story) SetCommits(v []Commit)`

SetCommits sets Commits field to given value.


### GetFiles

`func (o *Story) GetFiles() []UploadedFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Story) GetFilesOk() (*[]UploadedFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Story) SetFiles(v []UploadedFile)`

SetFiles sets Files field to given value.


### GetPosition

`func (o *Story) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Story) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Story) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetBlocked

`func (o *Story) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *Story) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *Story) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetProjectId

`func (o *Story) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Story) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Story) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *Story) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Story) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetDeadline

`func (o *Story) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *Story) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *Story) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### SetDeadlineNil

`func (o *Story) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *Story) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetStats

`func (o *Story) GetStats() StoryStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Story) GetStatsOk() (*StoryStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Story) SetStats(v StoryStats)`

SetStats sets Stats field to given value.


### GetCycleTime

`func (o *Story) GetCycleTime() int64`

GetCycleTime returns the CycleTime field if non-nil, zero value otherwise.

### GetCycleTimeOk

`func (o *Story) GetCycleTimeOk() (*int64, bool)`

GetCycleTimeOk returns a tuple with the CycleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleTime

`func (o *Story) SetCycleTime(v int64)`

SetCycleTime sets CycleTime field to given value.

### HasCycleTime

`func (o *Story) HasCycleTime() bool`

HasCycleTime returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Story) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Story) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Story) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMovedAt

`func (o *Story) GetMovedAt() time.Time`

GetMovedAt returns the MovedAt field if non-nil, zero value otherwise.

### GetMovedAtOk

`func (o *Story) GetMovedAtOk() (*time.Time, bool)`

GetMovedAtOk returns a tuple with the MovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedAt

`func (o *Story) SetMovedAt(v time.Time)`

SetMovedAt sets MovedAt field to given value.


### SetMovedAtNil

`func (o *Story) SetMovedAtNil(b bool)`

 SetMovedAtNil sets the value for MovedAt to be an explicit nil

### UnsetMovedAt
`func (o *Story) UnsetMovedAt()`

UnsetMovedAt ensures that no value is present for MovedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


