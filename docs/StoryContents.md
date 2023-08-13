# StoryContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the story. | [optional] 
**EntityType** | Pointer to **string** | A string description of this resource. | [optional] 
**Labels** | Pointer to [**[]LabelSlim**](LabelSlim.md) | An array of labels attached to the story. | [optional] 
**StoryType** | Pointer to **string** | The type of story (feature, bug, chore). | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValueParams**](CustomFieldValueParams.md) | An array of maps specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. | [optional] 
**LinkedFiles** | Pointer to [**[]LinkedFile**](LinkedFile.md) | An array of linked files attached to the story. | [optional] 
**Name** | Pointer to **string** | The name of the story. | [optional] 
**EpicId** | Pointer to **int64** | The ID of the epic the story belongs to. | [optional] 
**ExternalLinks** | Pointer to **[]string** | An array of external links connected to the story. | [optional] 
**IterationId** | Pointer to **int64** | The ID of the iteration the story belongs to. | [optional] 
**Tasks** | Pointer to [**[]StoryContentsTask**](StoryContentsTask.md) | An array of tasks connected to the story. | [optional] 
**LabelIds** | Pointer to **[]int64** | An array of label ids attached to the story. | [optional] 
**GroupId** | Pointer to **string** | The ID of the group to which the story is assigned. | [optional] 
**WorkflowStateId** | Pointer to **int64** | The ID of the workflow state the story is currently in. | [optional] 
**FollowerIds** | Pointer to **[]string** | An array of UUIDs for any Members listed as Followers. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs of the owners of this story. | [optional] 
**Estimate** | Pointer to **int64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [optional] 
**Files** | Pointer to [**[]UploadedFile**](UploadedFile.md) | An array of files attached to the story. | [optional] 
**ProjectId** | Pointer to **int64** | The ID of the project the story belongs to. | [optional] 
**Deadline** | Pointer to **time.Time** | The due date of the story. | [optional] 

## Methods

### NewStoryContents

`func NewStoryContents() *StoryContents`

NewStoryContents instantiates a new StoryContents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryContentsWithDefaults

`func NewStoryContentsWithDefaults() *StoryContents`

NewStoryContentsWithDefaults instantiates a new StoryContents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StoryContents) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoryContents) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoryContents) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoryContents) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityType

`func (o *StoryContents) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *StoryContents) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *StoryContents) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *StoryContents) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetLabels

`func (o *StoryContents) GetLabels() []LabelSlim`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StoryContents) GetLabelsOk() (*[]LabelSlim, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StoryContents) SetLabels(v []LabelSlim)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *StoryContents) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetStoryType

`func (o *StoryContents) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *StoryContents) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *StoryContents) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.

### HasStoryType

`func (o *StoryContents) HasStoryType() bool`

HasStoryType returns a boolean if a field has been set.

### GetCustomFields

`func (o *StoryContents) GetCustomFields() []CustomFieldValueParams`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *StoryContents) GetCustomFieldsOk() (*[]CustomFieldValueParams, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *StoryContents) SetCustomFields(v []CustomFieldValueParams)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *StoryContents) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetLinkedFiles

`func (o *StoryContents) GetLinkedFiles() []LinkedFile`

GetLinkedFiles returns the LinkedFiles field if non-nil, zero value otherwise.

### GetLinkedFilesOk

`func (o *StoryContents) GetLinkedFilesOk() (*[]LinkedFile, bool)`

GetLinkedFilesOk returns a tuple with the LinkedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFiles

`func (o *StoryContents) SetLinkedFiles(v []LinkedFile)`

SetLinkedFiles sets LinkedFiles field to given value.

### HasLinkedFiles

`func (o *StoryContents) HasLinkedFiles() bool`

HasLinkedFiles returns a boolean if a field has been set.

### GetName

`func (o *StoryContents) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoryContents) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoryContents) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoryContents) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEpicId

`func (o *StoryContents) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *StoryContents) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *StoryContents) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.

### HasEpicId

`func (o *StoryContents) HasEpicId() bool`

HasEpicId returns a boolean if a field has been set.

### GetExternalLinks

`func (o *StoryContents) GetExternalLinks() []string`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *StoryContents) GetExternalLinksOk() (*[]string, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *StoryContents) SetExternalLinks(v []string)`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinks

`func (o *StoryContents) HasExternalLinks() bool`

HasExternalLinks returns a boolean if a field has been set.

### GetIterationId

`func (o *StoryContents) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *StoryContents) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *StoryContents) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *StoryContents) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### GetTasks

`func (o *StoryContents) GetTasks() []StoryContentsTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *StoryContents) GetTasksOk() (*[]StoryContentsTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *StoryContents) SetTasks(v []StoryContentsTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *StoryContents) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetLabelIds

`func (o *StoryContents) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *StoryContents) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *StoryContents) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *StoryContents) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetGroupId

`func (o *StoryContents) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StoryContents) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StoryContents) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *StoryContents) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetWorkflowStateId

`func (o *StoryContents) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *StoryContents) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *StoryContents) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.

### HasWorkflowStateId

`func (o *StoryContents) HasWorkflowStateId() bool`

HasWorkflowStateId returns a boolean if a field has been set.

### GetFollowerIds

`func (o *StoryContents) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *StoryContents) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *StoryContents) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *StoryContents) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *StoryContents) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *StoryContents) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *StoryContents) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *StoryContents) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetEstimate

`func (o *StoryContents) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *StoryContents) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *StoryContents) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *StoryContents) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### GetFiles

`func (o *StoryContents) GetFiles() []UploadedFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *StoryContents) GetFilesOk() (*[]UploadedFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *StoryContents) SetFiles(v []UploadedFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *StoryContents) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetProjectId

`func (o *StoryContents) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *StoryContents) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *StoryContents) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *StoryContents) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetDeadline

`func (o *StoryContents) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *StoryContents) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *StoryContents) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *StoryContents) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


