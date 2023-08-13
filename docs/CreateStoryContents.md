# CreateStoryContents

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
**WorkflowId** | Pointer to **NullableInt64** | The ID of the workflow. | [optional] 
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
**ProjectId** | Pointer to **int64** | The ID of the project the story belongs to. | [optional] 
**LinkedFileIds** | Pointer to **[]int64** | An array of the linked file IDs to be populated. | [optional] 
**Deadline** | Pointer to **NullableTime** | The due date of the story. | [optional] 

## Methods

### NewCreateStoryContents

`func NewCreateStoryContents() *CreateStoryContents`

NewCreateStoryContents instantiates a new CreateStoryContents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoryContentsWithDefaults

`func NewCreateStoryContentsWithDefaults() *CreateStoryContents`

NewCreateStoryContentsWithDefaults instantiates a new CreateStoryContents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateStoryContents) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStoryContents) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStoryContents) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStoryContents) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityType

`func (o *CreateStoryContents) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateStoryContents) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateStoryContents) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CreateStoryContents) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetLabels

`func (o *CreateStoryContents) GetLabels() []CreateLabelParams`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateStoryContents) GetLabelsOk() (*[]CreateLabelParams, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateStoryContents) SetLabels(v []CreateLabelParams)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateStoryContents) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetStoryType

`func (o *CreateStoryContents) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *CreateStoryContents) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *CreateStoryContents) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.

### HasStoryType

`func (o *CreateStoryContents) HasStoryType() bool`

HasStoryType returns a boolean if a field has been set.

### GetCustomFields

`func (o *CreateStoryContents) GetCustomFields() []CustomFieldValueParams`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateStoryContents) GetCustomFieldsOk() (*[]CustomFieldValueParams, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateStoryContents) SetCustomFields(v []CustomFieldValueParams)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateStoryContents) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetLinkedFiles

`func (o *CreateStoryContents) GetLinkedFiles() []LinkedFile`

GetLinkedFiles returns the LinkedFiles field if non-nil, zero value otherwise.

### GetLinkedFilesOk

`func (o *CreateStoryContents) GetLinkedFilesOk() (*[]LinkedFile, bool)`

GetLinkedFilesOk returns a tuple with the LinkedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFiles

`func (o *CreateStoryContents) SetLinkedFiles(v []LinkedFile)`

SetLinkedFiles sets LinkedFiles field to given value.

### HasLinkedFiles

`func (o *CreateStoryContents) HasLinkedFiles() bool`

HasLinkedFiles returns a boolean if a field has been set.

### GetFileIds

`func (o *CreateStoryContents) GetFileIds() []int64`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *CreateStoryContents) GetFileIdsOk() (*[]int64, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *CreateStoryContents) SetFileIds(v []int64)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *CreateStoryContents) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### GetWorkflowId

`func (o *CreateStoryContents) GetWorkflowId() int64`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *CreateStoryContents) GetWorkflowIdOk() (*int64, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *CreateStoryContents) SetWorkflowId(v int64)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *CreateStoryContents) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### SetWorkflowIdNil

`func (o *CreateStoryContents) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *CreateStoryContents) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetName

`func (o *CreateStoryContents) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStoryContents) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStoryContents) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateStoryContents) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEpicId

`func (o *CreateStoryContents) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *CreateStoryContents) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *CreateStoryContents) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.

### HasEpicId

`func (o *CreateStoryContents) HasEpicId() bool`

HasEpicId returns a boolean if a field has been set.

### SetEpicIdNil

`func (o *CreateStoryContents) SetEpicIdNil(b bool)`

 SetEpicIdNil sets the value for EpicId to be an explicit nil

### UnsetEpicId
`func (o *CreateStoryContents) UnsetEpicId()`

UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
### GetExternalLinks

`func (o *CreateStoryContents) GetExternalLinks() []string`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *CreateStoryContents) GetExternalLinksOk() (*[]string, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *CreateStoryContents) SetExternalLinks(v []string)`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinks

`func (o *CreateStoryContents) HasExternalLinks() bool`

HasExternalLinks returns a boolean if a field has been set.

### GetIterationId

`func (o *CreateStoryContents) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *CreateStoryContents) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *CreateStoryContents) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *CreateStoryContents) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### SetIterationIdNil

`func (o *CreateStoryContents) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *CreateStoryContents) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetTasks

`func (o *CreateStoryContents) GetTasks() []EntityTemplateTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *CreateStoryContents) GetTasksOk() (*[]EntityTemplateTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *CreateStoryContents) SetTasks(v []EntityTemplateTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *CreateStoryContents) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetLabelIds

`func (o *CreateStoryContents) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CreateStoryContents) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CreateStoryContents) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *CreateStoryContents) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetGroupId

`func (o *CreateStoryContents) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateStoryContents) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateStoryContents) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CreateStoryContents) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *CreateStoryContents) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *CreateStoryContents) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetWorkflowStateId

`func (o *CreateStoryContents) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *CreateStoryContents) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *CreateStoryContents) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.

### HasWorkflowStateId

`func (o *CreateStoryContents) HasWorkflowStateId() bool`

HasWorkflowStateId returns a boolean if a field has been set.

### GetFollowerIds

`func (o *CreateStoryContents) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *CreateStoryContents) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *CreateStoryContents) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *CreateStoryContents) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *CreateStoryContents) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *CreateStoryContents) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *CreateStoryContents) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *CreateStoryContents) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetEstimate

`func (o *CreateStoryContents) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *CreateStoryContents) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *CreateStoryContents) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *CreateStoryContents) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### SetEstimateNil

`func (o *CreateStoryContents) SetEstimateNil(b bool)`

 SetEstimateNil sets the value for Estimate to be an explicit nil

### UnsetEstimate
`func (o *CreateStoryContents) UnsetEstimate()`

UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
### GetFiles

`func (o *CreateStoryContents) GetFiles() []UploadedFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CreateStoryContents) GetFilesOk() (*[]UploadedFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CreateStoryContents) SetFiles(v []UploadedFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CreateStoryContents) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateStoryContents) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateStoryContents) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateStoryContents) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateStoryContents) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetLinkedFileIds

`func (o *CreateStoryContents) GetLinkedFileIds() []int64`

GetLinkedFileIds returns the LinkedFileIds field if non-nil, zero value otherwise.

### GetLinkedFileIdsOk

`func (o *CreateStoryContents) GetLinkedFileIdsOk() (*[]int64, bool)`

GetLinkedFileIdsOk returns a tuple with the LinkedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFileIds

`func (o *CreateStoryContents) SetLinkedFileIds(v []int64)`

SetLinkedFileIds sets LinkedFileIds field to given value.

### HasLinkedFileIds

`func (o *CreateStoryContents) HasLinkedFileIds() bool`

HasLinkedFileIds returns a boolean if a field has been set.

### GetDeadline

`func (o *CreateStoryContents) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *CreateStoryContents) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *CreateStoryContents) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *CreateStoryContents) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### SetDeadlineNil

`func (o *CreateStoryContents) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *CreateStoryContents) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


