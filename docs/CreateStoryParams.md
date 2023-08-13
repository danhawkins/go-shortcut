# CreateStoryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the story. | [optional] 
**Archived** | Pointer to **bool** | Controls the story&#39;s archived state. | [optional] 
**StoryLinks** | Pointer to [**[]CreateStoryLinkParams**](CreateStoryLinkParams.md) | An array of story links attached to the story. | [optional] 
**Labels** | Pointer to [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels attached to the story. | [optional] 
**StoryType** | Pointer to **string** | The type of story (feature, bug, chore). | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValueParams**](CustomFieldValueParams.md) | A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. | [optional] 
**MoveTo** | Pointer to **string** | One of \&quot;first\&quot; or \&quot;last\&quot;. This can be used to move the given story to the first or last position in the workflow state. | [optional] 
**FileIds** | Pointer to **[]int64** | An array of IDs of files attached to the story. | [optional] 
**CompletedAtOverride** | Pointer to **time.Time** | A manual override for the time/date the Story was completed. | [optional] 
**Name** | **string** | The name of the story. | 
**Comments** | Pointer to [**[]CreateStoryCommentParams**](CreateStoryCommentParams.md) | An array of comments to add to the story. | [optional] 
**EpicId** | Pointer to **NullableInt64** | The ID of the epic the story belongs to. | [optional] 
**StoryTemplateId** | Pointer to **NullableString** | The id of the story template used to create this story, if applicable. This is just an association; no content from the story template is inherited by the story simply by setting this field. | [optional] 
**ExternalLinks** | Pointer to **[]string** | An array of External Links associated with this story. | [optional] 
**RequestedById** | Pointer to **string** | The ID of the member that requested the story. | [optional] 
**IterationId** | Pointer to **NullableInt64** | The ID of the iteration the story belongs to. | [optional] 
**Tasks** | Pointer to [**[]CreateTaskParams**](CreateTaskParams.md) | An array of tasks connected to the story. | [optional] 
**StartedAtOverride** | Pointer to **time.Time** | A manual override for the time/date the Story was started. | [optional] 
**GroupId** | Pointer to **NullableString** | The id of the group to associate with this story. | [optional] 
**WorkflowStateId** | Pointer to **int64** | The ID of the workflow state the story will be in. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time/date the Story was updated. | [optional] 
**FollowerIds** | Pointer to **[]string** | An array of UUIDs of the followers of this story. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs of the owners of this story. | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 
**Estimate** | Pointer to **NullableInt64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [optional] 
**ProjectId** | Pointer to **NullableInt64** | The ID of the project the story belongs to. | [optional] 
**LinkedFileIds** | Pointer to **[]int64** | An array of IDs of linked files attached to the story. | [optional] 
**Deadline** | Pointer to **NullableTime** | The due date of the story. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time/date the Story was created. | [optional] 

## Methods

### NewCreateStoryParams

`func NewCreateStoryParams(name string, ) *CreateStoryParams`

NewCreateStoryParams instantiates a new CreateStoryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoryParamsWithDefaults

`func NewCreateStoryParamsWithDefaults() *CreateStoryParams`

NewCreateStoryParamsWithDefaults instantiates a new CreateStoryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateStoryParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStoryParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStoryParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStoryParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *CreateStoryParams) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CreateStoryParams) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CreateStoryParams) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *CreateStoryParams) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetStoryLinks

`func (o *CreateStoryParams) GetStoryLinks() []CreateStoryLinkParams`

GetStoryLinks returns the StoryLinks field if non-nil, zero value otherwise.

### GetStoryLinksOk

`func (o *CreateStoryParams) GetStoryLinksOk() (*[]CreateStoryLinkParams, bool)`

GetStoryLinksOk returns a tuple with the StoryLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryLinks

`func (o *CreateStoryParams) SetStoryLinks(v []CreateStoryLinkParams)`

SetStoryLinks sets StoryLinks field to given value.

### HasStoryLinks

`func (o *CreateStoryParams) HasStoryLinks() bool`

HasStoryLinks returns a boolean if a field has been set.

### GetLabels

`func (o *CreateStoryParams) GetLabels() []CreateLabelParams`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateStoryParams) GetLabelsOk() (*[]CreateLabelParams, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateStoryParams) SetLabels(v []CreateLabelParams)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateStoryParams) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetStoryType

`func (o *CreateStoryParams) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *CreateStoryParams) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *CreateStoryParams) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.

### HasStoryType

`func (o *CreateStoryParams) HasStoryType() bool`

HasStoryType returns a boolean if a field has been set.

### GetCustomFields

`func (o *CreateStoryParams) GetCustomFields() []CustomFieldValueParams`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateStoryParams) GetCustomFieldsOk() (*[]CustomFieldValueParams, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateStoryParams) SetCustomFields(v []CustomFieldValueParams)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateStoryParams) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMoveTo

`func (o *CreateStoryParams) GetMoveTo() string`

GetMoveTo returns the MoveTo field if non-nil, zero value otherwise.

### GetMoveToOk

`func (o *CreateStoryParams) GetMoveToOk() (*string, bool)`

GetMoveToOk returns a tuple with the MoveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTo

`func (o *CreateStoryParams) SetMoveTo(v string)`

SetMoveTo sets MoveTo field to given value.

### HasMoveTo

`func (o *CreateStoryParams) HasMoveTo() bool`

HasMoveTo returns a boolean if a field has been set.

### GetFileIds

`func (o *CreateStoryParams) GetFileIds() []int64`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *CreateStoryParams) GetFileIdsOk() (*[]int64, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *CreateStoryParams) SetFileIds(v []int64)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *CreateStoryParams) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### GetCompletedAtOverride

`func (o *CreateStoryParams) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *CreateStoryParams) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *CreateStoryParams) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.

### HasCompletedAtOverride

`func (o *CreateStoryParams) HasCompletedAtOverride() bool`

HasCompletedAtOverride returns a boolean if a field has been set.

### GetName

`func (o *CreateStoryParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStoryParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStoryParams) SetName(v string)`

SetName sets Name field to given value.


### GetComments

`func (o *CreateStoryParams) GetComments() []CreateStoryCommentParams`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CreateStoryParams) GetCommentsOk() (*[]CreateStoryCommentParams, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CreateStoryParams) SetComments(v []CreateStoryCommentParams)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CreateStoryParams) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetEpicId

`func (o *CreateStoryParams) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *CreateStoryParams) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *CreateStoryParams) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.

### HasEpicId

`func (o *CreateStoryParams) HasEpicId() bool`

HasEpicId returns a boolean if a field has been set.

### SetEpicIdNil

`func (o *CreateStoryParams) SetEpicIdNil(b bool)`

 SetEpicIdNil sets the value for EpicId to be an explicit nil

### UnsetEpicId
`func (o *CreateStoryParams) UnsetEpicId()`

UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
### GetStoryTemplateId

`func (o *CreateStoryParams) GetStoryTemplateId() string`

GetStoryTemplateId returns the StoryTemplateId field if non-nil, zero value otherwise.

### GetStoryTemplateIdOk

`func (o *CreateStoryParams) GetStoryTemplateIdOk() (*string, bool)`

GetStoryTemplateIdOk returns a tuple with the StoryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryTemplateId

`func (o *CreateStoryParams) SetStoryTemplateId(v string)`

SetStoryTemplateId sets StoryTemplateId field to given value.

### HasStoryTemplateId

`func (o *CreateStoryParams) HasStoryTemplateId() bool`

HasStoryTemplateId returns a boolean if a field has been set.

### SetStoryTemplateIdNil

`func (o *CreateStoryParams) SetStoryTemplateIdNil(b bool)`

 SetStoryTemplateIdNil sets the value for StoryTemplateId to be an explicit nil

### UnsetStoryTemplateId
`func (o *CreateStoryParams) UnsetStoryTemplateId()`

UnsetStoryTemplateId ensures that no value is present for StoryTemplateId, not even an explicit nil
### GetExternalLinks

`func (o *CreateStoryParams) GetExternalLinks() []string`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *CreateStoryParams) GetExternalLinksOk() (*[]string, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *CreateStoryParams) SetExternalLinks(v []string)`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinks

`func (o *CreateStoryParams) HasExternalLinks() bool`

HasExternalLinks returns a boolean if a field has been set.

### GetRequestedById

`func (o *CreateStoryParams) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *CreateStoryParams) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *CreateStoryParams) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.

### HasRequestedById

`func (o *CreateStoryParams) HasRequestedById() bool`

HasRequestedById returns a boolean if a field has been set.

### GetIterationId

`func (o *CreateStoryParams) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *CreateStoryParams) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *CreateStoryParams) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *CreateStoryParams) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### SetIterationIdNil

`func (o *CreateStoryParams) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *CreateStoryParams) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetTasks

`func (o *CreateStoryParams) GetTasks() []CreateTaskParams`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *CreateStoryParams) GetTasksOk() (*[]CreateTaskParams, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *CreateStoryParams) SetTasks(v []CreateTaskParams)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *CreateStoryParams) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetStartedAtOverride

`func (o *CreateStoryParams) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *CreateStoryParams) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *CreateStoryParams) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.

### HasStartedAtOverride

`func (o *CreateStoryParams) HasStartedAtOverride() bool`

HasStartedAtOverride returns a boolean if a field has been set.

### GetGroupId

`func (o *CreateStoryParams) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateStoryParams) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateStoryParams) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CreateStoryParams) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CreateStoryParams) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CreateStoryParams) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetWorkflowStateId

`func (o *CreateStoryParams) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *CreateStoryParams) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *CreateStoryParams) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.

### HasWorkflowStateId

`func (o *CreateStoryParams) HasWorkflowStateId() bool`

HasWorkflowStateId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateStoryParams) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateStoryParams) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateStoryParams) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateStoryParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFollowerIds

`func (o *CreateStoryParams) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *CreateStoryParams) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *CreateStoryParams) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *CreateStoryParams) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *CreateStoryParams) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *CreateStoryParams) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *CreateStoryParams) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *CreateStoryParams) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateStoryParams) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateStoryParams) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateStoryParams) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateStoryParams) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEstimate

`func (o *CreateStoryParams) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *CreateStoryParams) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *CreateStoryParams) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *CreateStoryParams) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### SetEstimateNil

`func (o *CreateStoryParams) SetEstimateNil(b bool)`

 SetEstimateNil sets the value for Estimate to be an explicit nil

### UnsetEstimate
`func (o *CreateStoryParams) UnsetEstimate()`

UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
### GetProjectId

`func (o *CreateStoryParams) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateStoryParams) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateStoryParams) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateStoryParams) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *CreateStoryParams) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *CreateStoryParams) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLinkedFileIds

`func (o *CreateStoryParams) GetLinkedFileIds() []int64`

GetLinkedFileIds returns the LinkedFileIds field if non-nil, zero value otherwise.

### GetLinkedFileIdsOk

`func (o *CreateStoryParams) GetLinkedFileIdsOk() (*[]int64, bool)`

GetLinkedFileIdsOk returns a tuple with the LinkedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFileIds

`func (o *CreateStoryParams) SetLinkedFileIds(v []int64)`

SetLinkedFileIds sets LinkedFileIds field to given value.

### HasLinkedFileIds

`func (o *CreateStoryParams) HasLinkedFileIds() bool`

HasLinkedFileIds returns a boolean if a field has been set.

### GetDeadline

`func (o *CreateStoryParams) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *CreateStoryParams) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *CreateStoryParams) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *CreateStoryParams) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### SetDeadlineNil

`func (o *CreateStoryParams) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *CreateStoryParams) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetCreatedAt

`func (o *CreateStoryParams) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateStoryParams) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateStoryParams) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateStoryParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


