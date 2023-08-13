# UpdateStoryContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the story. | [optional] 
**EntityType** | Pointer to **string** | A string description of this resource. | [optional] 
**Labels** | Pointer to [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels to be populated by the template. | [optional] 
**StoryType** | Pointer to **string** | The type of story (feature, bug, chore). | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValueParams**](CustomFieldValueParams.md) | An array of maps specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. | [optional] 
**LinkedFiles** | Pointer to [**[]LinkedFile**](LinkedFile.md) | An array of linked files attached to the story. | [optional] 
**FileIds** | Pointer to **[]int64** | An array of the attached file IDs to be populated. | [optional] 
**Name** | Pointer to **string** | The name of the story. | [optional] 
**EpicId** | Pointer to **NullableInt64** | The ID of the epic the to be populated. | [optional] 
**ExternalLinks** | Pointer to **[]string** | An array of external links to be populated. | [optional] 
**IterationId** | Pointer to **NullableInt64** | The ID of the iteration the to be populated. | [optional] 
**Tasks** | Pointer to [**[]EntityTemplateTask**](EntityTemplateTask.md) | An array of tasks to be populated by the template. | [optional] 
**LabelIds** | Pointer to **[]int64** | An array of label ids attached to the story. | [optional] 
**GroupId** | Pointer to **NullableString** | The ID of the group to be populated. | [optional] 
**WorkflowStateId** | Pointer to **int64** | The ID of the workflow state the story is currently in. | [optional] 
**FollowerIds** | Pointer to **[]string** | An array of UUIDs for any Members listed as Followers. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs of the owners of this story. | [optional] 
**Estimate** | Pointer to **NullableInt64** | The numeric point estimate to be populated. | [optional] 
**Files** | Pointer to [**[]UploadedFile**](UploadedFile.md) | An array of files attached to the story. | [optional] 
**ProjectId** | Pointer to **NullableInt64** | The ID of the project the story belongs to. | [optional] 
**LinkedFileIds** | Pointer to **[]int64** | An array of the linked file IDs to be populated. | [optional] 
**Deadline** | Pointer to **NullableTime** | The due date of the story. | [optional] 

## Methods

### NewUpdateStoryContents

`func NewUpdateStoryContents() *UpdateStoryContents`

NewUpdateStoryContents instantiates a new UpdateStoryContents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStoryContentsWithDefaults

`func NewUpdateStoryContentsWithDefaults() *UpdateStoryContents`

NewUpdateStoryContentsWithDefaults instantiates a new UpdateStoryContents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateStoryContents) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateStoryContents) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateStoryContents) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateStoryContents) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityType

`func (o *UpdateStoryContents) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *UpdateStoryContents) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *UpdateStoryContents) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *UpdateStoryContents) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateStoryContents) GetLabels() []CreateLabelParams`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateStoryContents) GetLabelsOk() (*[]CreateLabelParams, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateStoryContents) SetLabels(v []CreateLabelParams)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateStoryContents) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetStoryType

`func (o *UpdateStoryContents) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *UpdateStoryContents) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *UpdateStoryContents) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.

### HasStoryType

`func (o *UpdateStoryContents) HasStoryType() bool`

HasStoryType returns a boolean if a field has been set.

### GetCustomFields

`func (o *UpdateStoryContents) GetCustomFields() []CustomFieldValueParams`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateStoryContents) GetCustomFieldsOk() (*[]CustomFieldValueParams, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateStoryContents) SetCustomFields(v []CustomFieldValueParams)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateStoryContents) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetLinkedFiles

`func (o *UpdateStoryContents) GetLinkedFiles() []LinkedFile`

GetLinkedFiles returns the LinkedFiles field if non-nil, zero value otherwise.

### GetLinkedFilesOk

`func (o *UpdateStoryContents) GetLinkedFilesOk() (*[]LinkedFile, bool)`

GetLinkedFilesOk returns a tuple with the LinkedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFiles

`func (o *UpdateStoryContents) SetLinkedFiles(v []LinkedFile)`

SetLinkedFiles sets LinkedFiles field to given value.

### HasLinkedFiles

`func (o *UpdateStoryContents) HasLinkedFiles() bool`

HasLinkedFiles returns a boolean if a field has been set.

### GetFileIds

`func (o *UpdateStoryContents) GetFileIds() []int64`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *UpdateStoryContents) GetFileIdsOk() (*[]int64, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *UpdateStoryContents) SetFileIds(v []int64)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *UpdateStoryContents) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### GetName

`func (o *UpdateStoryContents) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStoryContents) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStoryContents) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStoryContents) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEpicId

`func (o *UpdateStoryContents) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *UpdateStoryContents) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *UpdateStoryContents) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.

### HasEpicId

`func (o *UpdateStoryContents) HasEpicId() bool`

HasEpicId returns a boolean if a field has been set.

### SetEpicIdNil

`func (o *UpdateStoryContents) SetEpicIdNil(b bool)`

 SetEpicIdNil sets the value for EpicId to be an explicit nil

### UnsetEpicId
`func (o *UpdateStoryContents) UnsetEpicId()`

UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
### GetExternalLinks

`func (o *UpdateStoryContents) GetExternalLinks() []string`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *UpdateStoryContents) GetExternalLinksOk() (*[]string, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *UpdateStoryContents) SetExternalLinks(v []string)`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinks

`func (o *UpdateStoryContents) HasExternalLinks() bool`

HasExternalLinks returns a boolean if a field has been set.

### GetIterationId

`func (o *UpdateStoryContents) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *UpdateStoryContents) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *UpdateStoryContents) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *UpdateStoryContents) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### SetIterationIdNil

`func (o *UpdateStoryContents) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *UpdateStoryContents) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetTasks

`func (o *UpdateStoryContents) GetTasks() []EntityTemplateTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *UpdateStoryContents) GetTasksOk() (*[]EntityTemplateTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *UpdateStoryContents) SetTasks(v []EntityTemplateTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *UpdateStoryContents) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetLabelIds

`func (o *UpdateStoryContents) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *UpdateStoryContents) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *UpdateStoryContents) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *UpdateStoryContents) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetGroupId

`func (o *UpdateStoryContents) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateStoryContents) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateStoryContents) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UpdateStoryContents) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *UpdateStoryContents) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *UpdateStoryContents) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetWorkflowStateId

`func (o *UpdateStoryContents) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *UpdateStoryContents) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *UpdateStoryContents) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.

### HasWorkflowStateId

`func (o *UpdateStoryContents) HasWorkflowStateId() bool`

HasWorkflowStateId returns a boolean if a field has been set.

### GetFollowerIds

`func (o *UpdateStoryContents) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *UpdateStoryContents) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *UpdateStoryContents) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *UpdateStoryContents) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *UpdateStoryContents) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *UpdateStoryContents) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *UpdateStoryContents) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *UpdateStoryContents) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetEstimate

`func (o *UpdateStoryContents) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *UpdateStoryContents) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *UpdateStoryContents) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *UpdateStoryContents) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### SetEstimateNil

`func (o *UpdateStoryContents) SetEstimateNil(b bool)`

 SetEstimateNil sets the value for Estimate to be an explicit nil

### UnsetEstimate
`func (o *UpdateStoryContents) UnsetEstimate()`

UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
### GetFiles

`func (o *UpdateStoryContents) GetFiles() []UploadedFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *UpdateStoryContents) GetFilesOk() (*[]UploadedFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *UpdateStoryContents) SetFiles(v []UploadedFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *UpdateStoryContents) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetProjectId

`func (o *UpdateStoryContents) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateStoryContents) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateStoryContents) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UpdateStoryContents) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *UpdateStoryContents) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *UpdateStoryContents) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLinkedFileIds

`func (o *UpdateStoryContents) GetLinkedFileIds() []int64`

GetLinkedFileIds returns the LinkedFileIds field if non-nil, zero value otherwise.

### GetLinkedFileIdsOk

`func (o *UpdateStoryContents) GetLinkedFileIdsOk() (*[]int64, bool)`

GetLinkedFileIdsOk returns a tuple with the LinkedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFileIds

`func (o *UpdateStoryContents) SetLinkedFileIds(v []int64)`

SetLinkedFileIds sets LinkedFileIds field to given value.

### HasLinkedFileIds

`func (o *UpdateStoryContents) HasLinkedFileIds() bool`

HasLinkedFileIds returns a boolean if a field has been set.

### GetDeadline

`func (o *UpdateStoryContents) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *UpdateStoryContents) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *UpdateStoryContents) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *UpdateStoryContents) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### SetDeadlineNil

`func (o *UpdateStoryContents) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *UpdateStoryContents) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


