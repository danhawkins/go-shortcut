# UpdateStory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the story. | [optional] 
**Archived** | Pointer to **bool** | True if the story is archived, otherwise false. | [optional] 
**Labels** | Pointer to [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels attached to the story. | [optional] 
**PullRequestIds** | Pointer to **[]int64** | An array of IDs of Pull/Merge Requests attached to the story. | [optional] 
**StoryType** | Pointer to **string** | The type of story (feature, bug, chore). | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValueParams**](CustomFieldValueParams.md) | A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. | [optional] 
**MoveTo** | Pointer to **string** | One of \&quot;first\&quot; or \&quot;last\&quot;. This can be used to move the given story to the first or last position in the workflow state. | [optional] 
**FileIds** | Pointer to **[]int64** | An array of IDs of files attached to the story. | [optional] 
**CompletedAtOverride** | Pointer to **NullableTime** | A manual override for the time/date the Story was completed. | [optional] 
**Name** | Pointer to **string** | The title of the story. | [optional] 
**EpicId** | Pointer to **NullableInt64** | The ID of the epic the story belongs to. | [optional] 
**ExternalLinks** | Pointer to **[]string** | An array of External Links associated with this story. | [optional] 
**BranchIds** | Pointer to **[]int64** | An array of IDs of Branches attached to the story. | [optional] 
**CommitIds** | Pointer to **[]int64** | An array of IDs of Commits attached to the story. | [optional] 
**RequestedById** | Pointer to **string** | The ID of the member that requested the story. | [optional] 
**IterationId** | Pointer to **NullableInt64** | The ID of the iteration the story belongs to. | [optional] 
**StartedAtOverride** | Pointer to **NullableTime** | A manual override for the time/date the Story was started. | [optional] 
**GroupId** | Pointer to **NullableString** | The ID of the group to associate with this story | [optional] 
**WorkflowStateId** | Pointer to **int64** | The ID of the workflow state to put the story in. | [optional] 
**FollowerIds** | Pointer to **[]string** | An array of UUIDs of the followers of this story. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs of the owners of this story. | [optional] 
**BeforeId** | Pointer to **int64** | The ID of the story we want to move this story before. | [optional] 
**Estimate** | Pointer to **NullableInt64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [optional] 
**AfterId** | Pointer to **int64** | The ID of the story we want to move this story after. | [optional] 
**ProjectId** | Pointer to **NullableInt64** | The ID of the project the story belongs to. | [optional] 
**LinkedFileIds** | Pointer to **[]int64** | An array of IDs of linked files attached to the story. | [optional] 
**Deadline** | Pointer to **NullableTime** | The due date of the story. | [optional] 

## Methods

### NewUpdateStory

`func NewUpdateStory() *UpdateStory`

NewUpdateStory instantiates a new UpdateStory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStoryWithDefaults

`func NewUpdateStoryWithDefaults() *UpdateStory`

NewUpdateStoryWithDefaults instantiates a new UpdateStory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateStory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateStory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateStory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateStory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *UpdateStory) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateStory) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateStory) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateStory) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateStory) GetLabels() []CreateLabelParams`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateStory) GetLabelsOk() (*[]CreateLabelParams, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateStory) SetLabels(v []CreateLabelParams)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateStory) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPullRequestIds

`func (o *UpdateStory) GetPullRequestIds() []int64`

GetPullRequestIds returns the PullRequestIds field if non-nil, zero value otherwise.

### GetPullRequestIdsOk

`func (o *UpdateStory) GetPullRequestIdsOk() (*[]int64, bool)`

GetPullRequestIdsOk returns a tuple with the PullRequestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestIds

`func (o *UpdateStory) SetPullRequestIds(v []int64)`

SetPullRequestIds sets PullRequestIds field to given value.

### HasPullRequestIds

`func (o *UpdateStory) HasPullRequestIds() bool`

HasPullRequestIds returns a boolean if a field has been set.

### GetStoryType

`func (o *UpdateStory) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *UpdateStory) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *UpdateStory) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.

### HasStoryType

`func (o *UpdateStory) HasStoryType() bool`

HasStoryType returns a boolean if a field has been set.

### GetCustomFields

`func (o *UpdateStory) GetCustomFields() []CustomFieldValueParams`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateStory) GetCustomFieldsOk() (*[]CustomFieldValueParams, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateStory) SetCustomFields(v []CustomFieldValueParams)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateStory) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMoveTo

`func (o *UpdateStory) GetMoveTo() string`

GetMoveTo returns the MoveTo field if non-nil, zero value otherwise.

### GetMoveToOk

`func (o *UpdateStory) GetMoveToOk() (*string, bool)`

GetMoveToOk returns a tuple with the MoveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTo

`func (o *UpdateStory) SetMoveTo(v string)`

SetMoveTo sets MoveTo field to given value.

### HasMoveTo

`func (o *UpdateStory) HasMoveTo() bool`

HasMoveTo returns a boolean if a field has been set.

### GetFileIds

`func (o *UpdateStory) GetFileIds() []int64`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *UpdateStory) GetFileIdsOk() (*[]int64, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *UpdateStory) SetFileIds(v []int64)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *UpdateStory) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### GetCompletedAtOverride

`func (o *UpdateStory) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *UpdateStory) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *UpdateStory) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.

### HasCompletedAtOverride

`func (o *UpdateStory) HasCompletedAtOverride() bool`

HasCompletedAtOverride returns a boolean if a field has been set.

### SetCompletedAtOverrideNil

`func (o *UpdateStory) SetCompletedAtOverrideNil(b bool)`

 SetCompletedAtOverrideNil sets the value for CompletedAtOverride to be an explicit nil

### UnsetCompletedAtOverride
`func (o *UpdateStory) UnsetCompletedAtOverride()`

UnsetCompletedAtOverride ensures that no value is present for CompletedAtOverride, not even an explicit nil
### GetName

`func (o *UpdateStory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEpicId

`func (o *UpdateStory) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *UpdateStory) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *UpdateStory) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.

### HasEpicId

`func (o *UpdateStory) HasEpicId() bool`

HasEpicId returns a boolean if a field has been set.

### SetEpicIdNil

`func (o *UpdateStory) SetEpicIdNil(b bool)`

 SetEpicIdNil sets the value for EpicId to be an explicit nil

### UnsetEpicId
`func (o *UpdateStory) UnsetEpicId()`

UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
### GetExternalLinks

`func (o *UpdateStory) GetExternalLinks() []string`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *UpdateStory) GetExternalLinksOk() (*[]string, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *UpdateStory) SetExternalLinks(v []string)`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinks

`func (o *UpdateStory) HasExternalLinks() bool`

HasExternalLinks returns a boolean if a field has been set.

### GetBranchIds

`func (o *UpdateStory) GetBranchIds() []int64`

GetBranchIds returns the BranchIds field if non-nil, zero value otherwise.

### GetBranchIdsOk

`func (o *UpdateStory) GetBranchIdsOk() (*[]int64, bool)`

GetBranchIdsOk returns a tuple with the BranchIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIds

`func (o *UpdateStory) SetBranchIds(v []int64)`

SetBranchIds sets BranchIds field to given value.

### HasBranchIds

`func (o *UpdateStory) HasBranchIds() bool`

HasBranchIds returns a boolean if a field has been set.

### GetCommitIds

`func (o *UpdateStory) GetCommitIds() []int64`

GetCommitIds returns the CommitIds field if non-nil, zero value otherwise.

### GetCommitIdsOk

`func (o *UpdateStory) GetCommitIdsOk() (*[]int64, bool)`

GetCommitIdsOk returns a tuple with the CommitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitIds

`func (o *UpdateStory) SetCommitIds(v []int64)`

SetCommitIds sets CommitIds field to given value.

### HasCommitIds

`func (o *UpdateStory) HasCommitIds() bool`

HasCommitIds returns a boolean if a field has been set.

### GetRequestedById

`func (o *UpdateStory) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *UpdateStory) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *UpdateStory) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.

### HasRequestedById

`func (o *UpdateStory) HasRequestedById() bool`

HasRequestedById returns a boolean if a field has been set.

### GetIterationId

`func (o *UpdateStory) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *UpdateStory) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *UpdateStory) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *UpdateStory) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### SetIterationIdNil

`func (o *UpdateStory) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *UpdateStory) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetStartedAtOverride

`func (o *UpdateStory) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *UpdateStory) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *UpdateStory) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.

### HasStartedAtOverride

`func (o *UpdateStory) HasStartedAtOverride() bool`

HasStartedAtOverride returns a boolean if a field has been set.

### SetStartedAtOverrideNil

`func (o *UpdateStory) SetStartedAtOverrideNil(b bool)`

 SetStartedAtOverrideNil sets the value for StartedAtOverride to be an explicit nil

### UnsetStartedAtOverride
`func (o *UpdateStory) UnsetStartedAtOverride()`

UnsetStartedAtOverride ensures that no value is present for StartedAtOverride, not even an explicit nil
### GetGroupId

`func (o *UpdateStory) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateStory) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateStory) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UpdateStory) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *UpdateStory) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *UpdateStory) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetWorkflowStateId

`func (o *UpdateStory) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *UpdateStory) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *UpdateStory) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.

### HasWorkflowStateId

`func (o *UpdateStory) HasWorkflowStateId() bool`

HasWorkflowStateId returns a boolean if a field has been set.

### GetFollowerIds

`func (o *UpdateStory) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *UpdateStory) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *UpdateStory) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *UpdateStory) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *UpdateStory) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *UpdateStory) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *UpdateStory) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *UpdateStory) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetBeforeId

`func (o *UpdateStory) GetBeforeId() int64`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *UpdateStory) GetBeforeIdOk() (*int64, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *UpdateStory) SetBeforeId(v int64)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *UpdateStory) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetEstimate

`func (o *UpdateStory) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *UpdateStory) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *UpdateStory) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *UpdateStory) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### SetEstimateNil

`func (o *UpdateStory) SetEstimateNil(b bool)`

 SetEstimateNil sets the value for Estimate to be an explicit nil

### UnsetEstimate
`func (o *UpdateStory) UnsetEstimate()`

UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
### GetAfterId

`func (o *UpdateStory) GetAfterId() int64`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *UpdateStory) GetAfterIdOk() (*int64, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *UpdateStory) SetAfterId(v int64)`

SetAfterId sets AfterId field to given value.

### HasAfterId

`func (o *UpdateStory) HasAfterId() bool`

HasAfterId returns a boolean if a field has been set.

### GetProjectId

`func (o *UpdateStory) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateStory) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateStory) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UpdateStory) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *UpdateStory) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *UpdateStory) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLinkedFileIds

`func (o *UpdateStory) GetLinkedFileIds() []int64`

GetLinkedFileIds returns the LinkedFileIds field if non-nil, zero value otherwise.

### GetLinkedFileIdsOk

`func (o *UpdateStory) GetLinkedFileIdsOk() (*[]int64, bool)`

GetLinkedFileIdsOk returns a tuple with the LinkedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFileIds

`func (o *UpdateStory) SetLinkedFileIds(v []int64)`

SetLinkedFileIds sets LinkedFileIds field to given value.

### HasLinkedFileIds

`func (o *UpdateStory) HasLinkedFileIds() bool`

HasLinkedFileIds returns a boolean if a field has been set.

### GetDeadline

`func (o *UpdateStory) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *UpdateStory) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *UpdateStory) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *UpdateStory) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### SetDeadlineNil

`func (o *UpdateStory) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *UpdateStory) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


