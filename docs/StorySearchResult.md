# StorySearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Story. | 
**Description** | Pointer to **string** | The description of the story. | [optional] 
**Archived** | **bool** | True if the story has been archived or not. | 
**Started** | **bool** | A true/false boolean indicating if the Story has been started. | 
**StoryLinks** | [**[]TypedStoryLink**](TypedStoryLink.md) | An array of story links attached to the Story. | 
**EntityType** | **string** | A string description of this resource. | 
**Labels** | [**[]LabelSlim**](LabelSlim.md) | An array of labels attached to the story. | 
**TaskIds** | Pointer to **[]int64** | An array of IDs of Tasks attached to the story. | [optional] 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**SyncedItem** | Pointer to [**SyncedItem**](SyncedItem.md) |  | [optional] 
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Story description. | 
**StoryType** | **string** | The type of story (feature, bug, chore). | 
**CustomFields** | Pointer to [**[]StoryCustomField**](StoryCustomField.md) | An array of CustomField value assertions for the story. | [optional] 
**LinkedFiles** | Pointer to [**[]LinkedFile**](LinkedFile.md) | An array of linked files attached to the story. | [optional] 
**FileIds** | Pointer to **[]int64** | An array of IDs of Files attached to the story. | [optional] 
**NumTasksCompleted** | Pointer to **int64** | The number of tasks on the story which are complete. | [optional] 
**WorkflowId** | **int64** | The ID of the workflow the story belongs to. | 
**CompletedAtOverride** | **NullableTime** | A manual override for the time/date the Story was completed. | 
**StartedAt** | **NullableTime** | The time/date the Story was started. | 
**CompletedAt** | **NullableTime** | The time/date the Story was completed. | 
**Name** | **string** | The name of the story. | 
**GlobalId** | **string** |  | 
**Completed** | **bool** | A true/false boolean indicating if the Story has been completed. | 
**Comments** | Pointer to [**[]StoryComment**](StoryComment.md) | An array of comments attached to the story. | [optional] 
**Blocker** | **bool** | A true/false boolean indicating if the Story is currently a blocker of another story. | 
**Branches** | Pointer to [**[]Branch**](Branch.md) | An array of Git branches attached to the story. | [optional] 
**EpicId** | **NullableInt64** | The ID of the epic the story belongs to. | 
**UnresolvedBlockerComments** | Pointer to **[]int64** | The IDs of any unresolved blocker comments on the Story. | [optional] 
**StoryTemplateId** | **NullableString** | The ID of the story template used to create this story, or null if not created using a template. | 
**ExternalLinks** | **[]string** | An array of external links (strings) associated with a Story | 
**PreviousIterationIds** | **[]int64** | The IDs of the iteration the story belongs to. | 
**RequestedById** | **string** | The ID of the Member that requested the story. | 
**IterationId** | **NullableInt64** | The ID of the iteration the story belongs to. | 
**Tasks** | Pointer to [**[]Task**](Task.md) | An array of tasks connected to the story. | [optional] 
**LabelIds** | **[]int64** | An array of label ids attached to the story. | 
**StartedAtOverride** | **NullableTime** | A manual override for the time/date the Story was started. | 
**GroupId** | **NullableString** | The ID of the group associated with the story. | 
**WorkflowStateId** | **int64** | The ID of the workflow state the story is currently in. | 
**UpdatedAt** | **NullableTime** | The time/date the Story was updated. | 
**PullRequests** | Pointer to [**[]PullRequest**](PullRequest.md) | An array of Pull/Merge Requests attached to the story. | [optional] 
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Story description. | 
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | 
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | 
**ExternalId** | **NullableString** | This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here. | 
**Id** | **int64** | The unique ID of the Story. | 
**LeadTime** | Pointer to **int64** | The lead time (in seconds) of this story when complete. | [optional] 
**Estimate** | **NullableInt64** | The numeric point estimate of the story. Can also be null, which means unestimated. | 
**Commits** | Pointer to [**[]Commit**](Commit.md) | An array of commits attached to the story. | [optional] 
**Files** | Pointer to [**[]UploadedFile**](UploadedFile.md) | An array of files attached to the story. | [optional] 
**Position** | **int64** | A number representing the position of the story in relation to every other story in the current project. | 
**Blocked** | **bool** | A true/false boolean indicating if the Story is currently blocked. | 
**ProjectId** | **NullableInt64** | The ID of the project the story belongs to. | 
**LinkedFileIds** | Pointer to **[]int64** | An array of IDs of LinkedFiles attached to the story. | [optional] 
**Deadline** | **NullableTime** | The due date of the story. | 
**Stats** | [**StoryStats**](StoryStats.md) |  | 
**CommentIds** | Pointer to **[]int64** | An array of IDs of Comments attached to the story. | [optional] 
**CycleTime** | Pointer to **int64** | The cycle time (in seconds) of this story when complete. | [optional] 
**CreatedAt** | **time.Time** | The time/date the Story was created. | 
**MovedAt** | **NullableTime** | The time/date the Story was last changed workflow-state. | 

## Methods

### NewStorySearchResult

`func NewStorySearchResult(appUrl string, archived bool, started bool, storyLinks []TypedStoryLink, entityType string, labels []LabelSlim, mentionIds []string, memberMentionIds []string, storyType string, workflowId int64, completedAtOverride NullableTime, startedAt NullableTime, completedAt NullableTime, name string, globalId string, completed bool, blocker bool, epicId NullableInt64, storyTemplateId NullableString, externalLinks []string, previousIterationIds []int64, requestedById string, iterationId NullableInt64, labelIds []int64, startedAtOverride NullableTime, groupId NullableString, workflowStateId int64, updatedAt NullableTime, groupMentionIds []string, followerIds []string, ownerIds []string, externalId NullableString, id int64, estimate NullableInt64, position int64, blocked bool, projectId NullableInt64, deadline NullableTime, stats StoryStats, createdAt time.Time, movedAt NullableTime, ) *StorySearchResult`

NewStorySearchResult instantiates a new StorySearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorySearchResultWithDefaults

`func NewStorySearchResultWithDefaults() *StorySearchResult`

NewStorySearchResultWithDefaults instantiates a new StorySearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *StorySearchResult) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *StorySearchResult) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *StorySearchResult) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetDescription

`func (o *StorySearchResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorySearchResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorySearchResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorySearchResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *StorySearchResult) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *StorySearchResult) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *StorySearchResult) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetStarted

`func (o *StorySearchResult) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *StorySearchResult) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *StorySearchResult) SetStarted(v bool)`

SetStarted sets Started field to given value.


### GetStoryLinks

`func (o *StorySearchResult) GetStoryLinks() []TypedStoryLink`

GetStoryLinks returns the StoryLinks field if non-nil, zero value otherwise.

### GetStoryLinksOk

`func (o *StorySearchResult) GetStoryLinksOk() (*[]TypedStoryLink, bool)`

GetStoryLinksOk returns a tuple with the StoryLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryLinks

`func (o *StorySearchResult) SetStoryLinks(v []TypedStoryLink)`

SetStoryLinks sets StoryLinks field to given value.


### GetEntityType

`func (o *StorySearchResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *StorySearchResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *StorySearchResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetLabels

`func (o *StorySearchResult) GetLabels() []LabelSlim`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StorySearchResult) GetLabelsOk() (*[]LabelSlim, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StorySearchResult) SetLabels(v []LabelSlim)`

SetLabels sets Labels field to given value.


### GetTaskIds

`func (o *StorySearchResult) GetTaskIds() []int64`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *StorySearchResult) GetTaskIdsOk() (*[]int64, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *StorySearchResult) SetTaskIds(v []int64)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *StorySearchResult) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### GetMentionIds

`func (o *StorySearchResult) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *StorySearchResult) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *StorySearchResult) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetSyncedItem

`func (o *StorySearchResult) GetSyncedItem() SyncedItem`

GetSyncedItem returns the SyncedItem field if non-nil, zero value otherwise.

### GetSyncedItemOk

`func (o *StorySearchResult) GetSyncedItemOk() (*SyncedItem, bool)`

GetSyncedItemOk returns a tuple with the SyncedItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedItem

`func (o *StorySearchResult) SetSyncedItem(v SyncedItem)`

SetSyncedItem sets SyncedItem field to given value.

### HasSyncedItem

`func (o *StorySearchResult) HasSyncedItem() bool`

HasSyncedItem returns a boolean if a field has been set.

### GetMemberMentionIds

`func (o *StorySearchResult) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *StorySearchResult) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *StorySearchResult) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetStoryType

`func (o *StorySearchResult) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *StorySearchResult) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *StorySearchResult) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.


### GetCustomFields

`func (o *StorySearchResult) GetCustomFields() []StoryCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *StorySearchResult) GetCustomFieldsOk() (*[]StoryCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *StorySearchResult) SetCustomFields(v []StoryCustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *StorySearchResult) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetLinkedFiles

`func (o *StorySearchResult) GetLinkedFiles() []LinkedFile`

GetLinkedFiles returns the LinkedFiles field if non-nil, zero value otherwise.

### GetLinkedFilesOk

`func (o *StorySearchResult) GetLinkedFilesOk() (*[]LinkedFile, bool)`

GetLinkedFilesOk returns a tuple with the LinkedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFiles

`func (o *StorySearchResult) SetLinkedFiles(v []LinkedFile)`

SetLinkedFiles sets LinkedFiles field to given value.

### HasLinkedFiles

`func (o *StorySearchResult) HasLinkedFiles() bool`

HasLinkedFiles returns a boolean if a field has been set.

### GetFileIds

`func (o *StorySearchResult) GetFileIds() []int64`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *StorySearchResult) GetFileIdsOk() (*[]int64, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *StorySearchResult) SetFileIds(v []int64)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *StorySearchResult) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### GetNumTasksCompleted

`func (o *StorySearchResult) GetNumTasksCompleted() int64`

GetNumTasksCompleted returns the NumTasksCompleted field if non-nil, zero value otherwise.

### GetNumTasksCompletedOk

`func (o *StorySearchResult) GetNumTasksCompletedOk() (*int64, bool)`

GetNumTasksCompletedOk returns a tuple with the NumTasksCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTasksCompleted

`func (o *StorySearchResult) SetNumTasksCompleted(v int64)`

SetNumTasksCompleted sets NumTasksCompleted field to given value.

### HasNumTasksCompleted

`func (o *StorySearchResult) HasNumTasksCompleted() bool`

HasNumTasksCompleted returns a boolean if a field has been set.

### GetWorkflowId

`func (o *StorySearchResult) GetWorkflowId() int64`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *StorySearchResult) GetWorkflowIdOk() (*int64, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *StorySearchResult) SetWorkflowId(v int64)`

SetWorkflowId sets WorkflowId field to given value.


### GetCompletedAtOverride

`func (o *StorySearchResult) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *StorySearchResult) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *StorySearchResult) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.


### SetCompletedAtOverrideNil

`func (o *StorySearchResult) SetCompletedAtOverrideNil(b bool)`

 SetCompletedAtOverrideNil sets the value for CompletedAtOverride to be an explicit nil

### UnsetCompletedAtOverride
`func (o *StorySearchResult) UnsetCompletedAtOverride()`

UnsetCompletedAtOverride ensures that no value is present for CompletedAtOverride, not even an explicit nil
### GetStartedAt

`func (o *StorySearchResult) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *StorySearchResult) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *StorySearchResult) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *StorySearchResult) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *StorySearchResult) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *StorySearchResult) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *StorySearchResult) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *StorySearchResult) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *StorySearchResult) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *StorySearchResult) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetName

`func (o *StorySearchResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorySearchResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorySearchResult) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *StorySearchResult) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *StorySearchResult) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *StorySearchResult) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetCompleted

`func (o *StorySearchResult) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *StorySearchResult) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *StorySearchResult) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetComments

`func (o *StorySearchResult) GetComments() []StoryComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *StorySearchResult) GetCommentsOk() (*[]StoryComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *StorySearchResult) SetComments(v []StoryComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *StorySearchResult) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetBlocker

`func (o *StorySearchResult) GetBlocker() bool`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *StorySearchResult) GetBlockerOk() (*bool, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *StorySearchResult) SetBlocker(v bool)`

SetBlocker sets Blocker field to given value.


### GetBranches

`func (o *StorySearchResult) GetBranches() []Branch`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *StorySearchResult) GetBranchesOk() (*[]Branch, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *StorySearchResult) SetBranches(v []Branch)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *StorySearchResult) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetEpicId

`func (o *StorySearchResult) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *StorySearchResult) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *StorySearchResult) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.


### SetEpicIdNil

`func (o *StorySearchResult) SetEpicIdNil(b bool)`

 SetEpicIdNil sets the value for EpicId to be an explicit nil

### UnsetEpicId
`func (o *StorySearchResult) UnsetEpicId()`

UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
### GetUnresolvedBlockerComments

`func (o *StorySearchResult) GetUnresolvedBlockerComments() []int64`

GetUnresolvedBlockerComments returns the UnresolvedBlockerComments field if non-nil, zero value otherwise.

### GetUnresolvedBlockerCommentsOk

`func (o *StorySearchResult) GetUnresolvedBlockerCommentsOk() (*[]int64, bool)`

GetUnresolvedBlockerCommentsOk returns a tuple with the UnresolvedBlockerComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedBlockerComments

`func (o *StorySearchResult) SetUnresolvedBlockerComments(v []int64)`

SetUnresolvedBlockerComments sets UnresolvedBlockerComments field to given value.

### HasUnresolvedBlockerComments

`func (o *StorySearchResult) HasUnresolvedBlockerComments() bool`

HasUnresolvedBlockerComments returns a boolean if a field has been set.

### GetStoryTemplateId

`func (o *StorySearchResult) GetStoryTemplateId() string`

GetStoryTemplateId returns the StoryTemplateId field if non-nil, zero value otherwise.

### GetStoryTemplateIdOk

`func (o *StorySearchResult) GetStoryTemplateIdOk() (*string, bool)`

GetStoryTemplateIdOk returns a tuple with the StoryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryTemplateId

`func (o *StorySearchResult) SetStoryTemplateId(v string)`

SetStoryTemplateId sets StoryTemplateId field to given value.


### SetStoryTemplateIdNil

`func (o *StorySearchResult) SetStoryTemplateIdNil(b bool)`

 SetStoryTemplateIdNil sets the value for StoryTemplateId to be an explicit nil

### UnsetStoryTemplateId
`func (o *StorySearchResult) UnsetStoryTemplateId()`

UnsetStoryTemplateId ensures that no value is present for StoryTemplateId, not even an explicit nil
### GetExternalLinks

`func (o *StorySearchResult) GetExternalLinks() []string`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *StorySearchResult) GetExternalLinksOk() (*[]string, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *StorySearchResult) SetExternalLinks(v []string)`

SetExternalLinks sets ExternalLinks field to given value.


### GetPreviousIterationIds

`func (o *StorySearchResult) GetPreviousIterationIds() []int64`

GetPreviousIterationIds returns the PreviousIterationIds field if non-nil, zero value otherwise.

### GetPreviousIterationIdsOk

`func (o *StorySearchResult) GetPreviousIterationIdsOk() (*[]int64, bool)`

GetPreviousIterationIdsOk returns a tuple with the PreviousIterationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIterationIds

`func (o *StorySearchResult) SetPreviousIterationIds(v []int64)`

SetPreviousIterationIds sets PreviousIterationIds field to given value.


### GetRequestedById

`func (o *StorySearchResult) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *StorySearchResult) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *StorySearchResult) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.


### GetIterationId

`func (o *StorySearchResult) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *StorySearchResult) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *StorySearchResult) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.


### SetIterationIdNil

`func (o *StorySearchResult) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *StorySearchResult) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetTasks

`func (o *StorySearchResult) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *StorySearchResult) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *StorySearchResult) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *StorySearchResult) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetLabelIds

`func (o *StorySearchResult) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *StorySearchResult) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *StorySearchResult) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.


### GetStartedAtOverride

`func (o *StorySearchResult) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *StorySearchResult) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *StorySearchResult) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.


### SetStartedAtOverrideNil

`func (o *StorySearchResult) SetStartedAtOverrideNil(b bool)`

 SetStartedAtOverrideNil sets the value for StartedAtOverride to be an explicit nil

### UnsetStartedAtOverride
`func (o *StorySearchResult) UnsetStartedAtOverride()`

UnsetStartedAtOverride ensures that no value is present for StartedAtOverride, not even an explicit nil
### GetGroupId

`func (o *StorySearchResult) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StorySearchResult) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StorySearchResult) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### SetGroupIdNil

`func (o *StorySearchResult) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *StorySearchResult) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetWorkflowStateId

`func (o *StorySearchResult) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *StorySearchResult) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *StorySearchResult) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.


### GetUpdatedAt

`func (o *StorySearchResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StorySearchResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StorySearchResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *StorySearchResult) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *StorySearchResult) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetPullRequests

`func (o *StorySearchResult) GetPullRequests() []PullRequest`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *StorySearchResult) GetPullRequestsOk() (*[]PullRequest, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *StorySearchResult) SetPullRequests(v []PullRequest)`

SetPullRequests sets PullRequests field to given value.

### HasPullRequests

`func (o *StorySearchResult) HasPullRequests() bool`

HasPullRequests returns a boolean if a field has been set.

### GetGroupMentionIds

`func (o *StorySearchResult) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *StorySearchResult) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *StorySearchResult) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetFollowerIds

`func (o *StorySearchResult) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *StorySearchResult) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *StorySearchResult) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.


### GetOwnerIds

`func (o *StorySearchResult) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *StorySearchResult) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *StorySearchResult) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.


### GetExternalId

`func (o *StorySearchResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *StorySearchResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *StorySearchResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *StorySearchResult) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *StorySearchResult) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetId

`func (o *StorySearchResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorySearchResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorySearchResult) SetId(v int64)`

SetId sets Id field to given value.


### GetLeadTime

`func (o *StorySearchResult) GetLeadTime() int64`

GetLeadTime returns the LeadTime field if non-nil, zero value otherwise.

### GetLeadTimeOk

`func (o *StorySearchResult) GetLeadTimeOk() (*int64, bool)`

GetLeadTimeOk returns a tuple with the LeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadTime

`func (o *StorySearchResult) SetLeadTime(v int64)`

SetLeadTime sets LeadTime field to given value.

### HasLeadTime

`func (o *StorySearchResult) HasLeadTime() bool`

HasLeadTime returns a boolean if a field has been set.

### GetEstimate

`func (o *StorySearchResult) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *StorySearchResult) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *StorySearchResult) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.


### SetEstimateNil

`func (o *StorySearchResult) SetEstimateNil(b bool)`

 SetEstimateNil sets the value for Estimate to be an explicit nil

### UnsetEstimate
`func (o *StorySearchResult) UnsetEstimate()`

UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
### GetCommits

`func (o *StorySearchResult) GetCommits() []Commit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *StorySearchResult) GetCommitsOk() (*[]Commit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *StorySearchResult) SetCommits(v []Commit)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *StorySearchResult) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetFiles

`func (o *StorySearchResult) GetFiles() []UploadedFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *StorySearchResult) GetFilesOk() (*[]UploadedFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *StorySearchResult) SetFiles(v []UploadedFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *StorySearchResult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetPosition

`func (o *StorySearchResult) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *StorySearchResult) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *StorySearchResult) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetBlocked

`func (o *StorySearchResult) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *StorySearchResult) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *StorySearchResult) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetProjectId

`func (o *StorySearchResult) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *StorySearchResult) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *StorySearchResult) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *StorySearchResult) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *StorySearchResult) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLinkedFileIds

`func (o *StorySearchResult) GetLinkedFileIds() []int64`

GetLinkedFileIds returns the LinkedFileIds field if non-nil, zero value otherwise.

### GetLinkedFileIdsOk

`func (o *StorySearchResult) GetLinkedFileIdsOk() (*[]int64, bool)`

GetLinkedFileIdsOk returns a tuple with the LinkedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFileIds

`func (o *StorySearchResult) SetLinkedFileIds(v []int64)`

SetLinkedFileIds sets LinkedFileIds field to given value.

### HasLinkedFileIds

`func (o *StorySearchResult) HasLinkedFileIds() bool`

HasLinkedFileIds returns a boolean if a field has been set.

### GetDeadline

`func (o *StorySearchResult) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *StorySearchResult) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *StorySearchResult) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### SetDeadlineNil

`func (o *StorySearchResult) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *StorySearchResult) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetStats

`func (o *StorySearchResult) GetStats() StoryStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StorySearchResult) GetStatsOk() (*StoryStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StorySearchResult) SetStats(v StoryStats)`

SetStats sets Stats field to given value.


### GetCommentIds

`func (o *StorySearchResult) GetCommentIds() []int64`

GetCommentIds returns the CommentIds field if non-nil, zero value otherwise.

### GetCommentIdsOk

`func (o *StorySearchResult) GetCommentIdsOk() (*[]int64, bool)`

GetCommentIdsOk returns a tuple with the CommentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentIds

`func (o *StorySearchResult) SetCommentIds(v []int64)`

SetCommentIds sets CommentIds field to given value.

### HasCommentIds

`func (o *StorySearchResult) HasCommentIds() bool`

HasCommentIds returns a boolean if a field has been set.

### GetCycleTime

`func (o *StorySearchResult) GetCycleTime() int64`

GetCycleTime returns the CycleTime field if non-nil, zero value otherwise.

### GetCycleTimeOk

`func (o *StorySearchResult) GetCycleTimeOk() (*int64, bool)`

GetCycleTimeOk returns a tuple with the CycleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleTime

`func (o *StorySearchResult) SetCycleTime(v int64)`

SetCycleTime sets CycleTime field to given value.

### HasCycleTime

`func (o *StorySearchResult) HasCycleTime() bool`

HasCycleTime returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StorySearchResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StorySearchResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StorySearchResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMovedAt

`func (o *StorySearchResult) GetMovedAt() time.Time`

GetMovedAt returns the MovedAt field if non-nil, zero value otherwise.

### GetMovedAtOk

`func (o *StorySearchResult) GetMovedAtOk() (*time.Time, bool)`

GetMovedAtOk returns a tuple with the MovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovedAt

`func (o *StorySearchResult) SetMovedAt(v time.Time)`

SetMovedAt sets MovedAt field to given value.


### SetMovedAtNil

`func (o *StorySearchResult) SetMovedAtNil(b bool)`

 SetMovedAtNil sets the value for MovedAt to be an explicit nil

### UnsetMovedAt
`func (o *StorySearchResult) UnsetMovedAt()`

UnsetMovedAt ensures that no value is present for MovedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


