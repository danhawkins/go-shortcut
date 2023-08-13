# UpdateStories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | If the Stories should be archived or not. | [optional] 
**StoryIds** | **[]int64** | The Ids of the Stories you wish to update. | 
**StoryType** | Pointer to **string** | The type of story (feature, bug, chore). | [optional] 
**MoveTo** | Pointer to **string** | One of \&quot;first\&quot; or \&quot;last\&quot;. This can be used to move the given story to the first or last position in the workflow state. | [optional] 
**FollowerIdsAdd** | Pointer to **[]string** | The UUIDs of the new followers to be added. | [optional] 
**EpicId** | Pointer to **NullableInt64** | The ID of the epic the story belongs to. | [optional] 
**ExternalLinks** | Pointer to **[]string** | An array of External Links associated with this story. | [optional] 
**FollowerIdsRemove** | Pointer to **[]string** | The UUIDs of the followers to be removed. | [optional] 
**RequestedById** | Pointer to **string** | The ID of the member that requested the story. | [optional] 
**IterationId** | Pointer to **NullableInt64** | The ID of the iteration the story belongs to. | [optional] 
**CustomFieldsRemove** | Pointer to [**[]CustomFieldValueParams**](CustomFieldValueParams.md) | A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. | [optional] 
**LabelsAdd** | Pointer to [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels to be added. | [optional] 
**GroupId** | Pointer to **NullableString** | The Id of the Group the Stories should belong to. | [optional] 
**WorkflowStateId** | Pointer to **int64** | The ID of the workflow state to put the stories in. | [optional] 
**BeforeId** | Pointer to **int64** | The ID of the story that the stories are to be moved before. | [optional] 
**Estimate** | Pointer to **NullableInt64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [optional] 
**AfterId** | Pointer to **int64** | The ID of the story that the stories are to be moved below. | [optional] 
**OwnerIdsRemove** | Pointer to **[]string** | The UUIDs of the owners to be removed. | [optional] 
**CustomFieldsAdd** | Pointer to [**[]CustomFieldValueParams**](CustomFieldValueParams.md) | A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. | [optional] 
**ProjectId** | Pointer to **NullableInt64** | The ID of the Project the Stories should belong to. | [optional] 
**LabelsRemove** | Pointer to [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels to be removed. | [optional] 
**Deadline** | Pointer to **NullableTime** | The due date of the story. | [optional] 
**OwnerIdsAdd** | Pointer to **[]string** | The UUIDs of the new owners to be added. | [optional] 

## Methods

### NewUpdateStories

`func NewUpdateStories(storyIds []int64, ) *UpdateStories`

NewUpdateStories instantiates a new UpdateStories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStoriesWithDefaults

`func NewUpdateStoriesWithDefaults() *UpdateStories`

NewUpdateStoriesWithDefaults instantiates a new UpdateStories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *UpdateStories) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateStories) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateStories) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateStories) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetStoryIds

`func (o *UpdateStories) GetStoryIds() []int64`

GetStoryIds returns the StoryIds field if non-nil, zero value otherwise.

### GetStoryIdsOk

`func (o *UpdateStories) GetStoryIdsOk() (*[]int64, bool)`

GetStoryIdsOk returns a tuple with the StoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryIds

`func (o *UpdateStories) SetStoryIds(v []int64)`

SetStoryIds sets StoryIds field to given value.


### GetStoryType

`func (o *UpdateStories) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *UpdateStories) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *UpdateStories) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.

### HasStoryType

`func (o *UpdateStories) HasStoryType() bool`

HasStoryType returns a boolean if a field has been set.

### GetMoveTo

`func (o *UpdateStories) GetMoveTo() string`

GetMoveTo returns the MoveTo field if non-nil, zero value otherwise.

### GetMoveToOk

`func (o *UpdateStories) GetMoveToOk() (*string, bool)`

GetMoveToOk returns a tuple with the MoveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTo

`func (o *UpdateStories) SetMoveTo(v string)`

SetMoveTo sets MoveTo field to given value.

### HasMoveTo

`func (o *UpdateStories) HasMoveTo() bool`

HasMoveTo returns a boolean if a field has been set.

### GetFollowerIdsAdd

`func (o *UpdateStories) GetFollowerIdsAdd() []string`

GetFollowerIdsAdd returns the FollowerIdsAdd field if non-nil, zero value otherwise.

### GetFollowerIdsAddOk

`func (o *UpdateStories) GetFollowerIdsAddOk() (*[]string, bool)`

GetFollowerIdsAddOk returns a tuple with the FollowerIdsAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIdsAdd

`func (o *UpdateStories) SetFollowerIdsAdd(v []string)`

SetFollowerIdsAdd sets FollowerIdsAdd field to given value.

### HasFollowerIdsAdd

`func (o *UpdateStories) HasFollowerIdsAdd() bool`

HasFollowerIdsAdd returns a boolean if a field has been set.

### GetEpicId

`func (o *UpdateStories) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *UpdateStories) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *UpdateStories) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.

### HasEpicId

`func (o *UpdateStories) HasEpicId() bool`

HasEpicId returns a boolean if a field has been set.

### SetEpicIdNil

`func (o *UpdateStories) SetEpicIdNil(b bool)`

 SetEpicIdNil sets the value for EpicId to be an explicit nil

### UnsetEpicId
`func (o *UpdateStories) UnsetEpicId()`

UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
### GetExternalLinks

`func (o *UpdateStories) GetExternalLinks() []string`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *UpdateStories) GetExternalLinksOk() (*[]string, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *UpdateStories) SetExternalLinks(v []string)`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinks

`func (o *UpdateStories) HasExternalLinks() bool`

HasExternalLinks returns a boolean if a field has been set.

### GetFollowerIdsRemove

`func (o *UpdateStories) GetFollowerIdsRemove() []string`

GetFollowerIdsRemove returns the FollowerIdsRemove field if non-nil, zero value otherwise.

### GetFollowerIdsRemoveOk

`func (o *UpdateStories) GetFollowerIdsRemoveOk() (*[]string, bool)`

GetFollowerIdsRemoveOk returns a tuple with the FollowerIdsRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIdsRemove

`func (o *UpdateStories) SetFollowerIdsRemove(v []string)`

SetFollowerIdsRemove sets FollowerIdsRemove field to given value.

### HasFollowerIdsRemove

`func (o *UpdateStories) HasFollowerIdsRemove() bool`

HasFollowerIdsRemove returns a boolean if a field has been set.

### GetRequestedById

`func (o *UpdateStories) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *UpdateStories) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *UpdateStories) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.

### HasRequestedById

`func (o *UpdateStories) HasRequestedById() bool`

HasRequestedById returns a boolean if a field has been set.

### GetIterationId

`func (o *UpdateStories) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *UpdateStories) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *UpdateStories) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *UpdateStories) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### SetIterationIdNil

`func (o *UpdateStories) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *UpdateStories) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetCustomFieldsRemove

`func (o *UpdateStories) GetCustomFieldsRemove() []CustomFieldValueParams`

GetCustomFieldsRemove returns the CustomFieldsRemove field if non-nil, zero value otherwise.

### GetCustomFieldsRemoveOk

`func (o *UpdateStories) GetCustomFieldsRemoveOk() (*[]CustomFieldValueParams, bool)`

GetCustomFieldsRemoveOk returns a tuple with the CustomFieldsRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldsRemove

`func (o *UpdateStories) SetCustomFieldsRemove(v []CustomFieldValueParams)`

SetCustomFieldsRemove sets CustomFieldsRemove field to given value.

### HasCustomFieldsRemove

`func (o *UpdateStories) HasCustomFieldsRemove() bool`

HasCustomFieldsRemove returns a boolean if a field has been set.

### GetLabelsAdd

`func (o *UpdateStories) GetLabelsAdd() []CreateLabelParams`

GetLabelsAdd returns the LabelsAdd field if non-nil, zero value otherwise.

### GetLabelsAddOk

`func (o *UpdateStories) GetLabelsAddOk() (*[]CreateLabelParams, bool)`

GetLabelsAddOk returns a tuple with the LabelsAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsAdd

`func (o *UpdateStories) SetLabelsAdd(v []CreateLabelParams)`

SetLabelsAdd sets LabelsAdd field to given value.

### HasLabelsAdd

`func (o *UpdateStories) HasLabelsAdd() bool`

HasLabelsAdd returns a boolean if a field has been set.

### GetGroupId

`func (o *UpdateStories) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateStories) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateStories) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UpdateStories) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *UpdateStories) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *UpdateStories) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetWorkflowStateId

`func (o *UpdateStories) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *UpdateStories) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *UpdateStories) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.

### HasWorkflowStateId

`func (o *UpdateStories) HasWorkflowStateId() bool`

HasWorkflowStateId returns a boolean if a field has been set.

### GetBeforeId

`func (o *UpdateStories) GetBeforeId() int64`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *UpdateStories) GetBeforeIdOk() (*int64, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *UpdateStories) SetBeforeId(v int64)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *UpdateStories) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetEstimate

`func (o *UpdateStories) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *UpdateStories) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *UpdateStories) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *UpdateStories) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### SetEstimateNil

`func (o *UpdateStories) SetEstimateNil(b bool)`

 SetEstimateNil sets the value for Estimate to be an explicit nil

### UnsetEstimate
`func (o *UpdateStories) UnsetEstimate()`

UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
### GetAfterId

`func (o *UpdateStories) GetAfterId() int64`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *UpdateStories) GetAfterIdOk() (*int64, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *UpdateStories) SetAfterId(v int64)`

SetAfterId sets AfterId field to given value.

### HasAfterId

`func (o *UpdateStories) HasAfterId() bool`

HasAfterId returns a boolean if a field has been set.

### GetOwnerIdsRemove

`func (o *UpdateStories) GetOwnerIdsRemove() []string`

GetOwnerIdsRemove returns the OwnerIdsRemove field if non-nil, zero value otherwise.

### GetOwnerIdsRemoveOk

`func (o *UpdateStories) GetOwnerIdsRemoveOk() (*[]string, bool)`

GetOwnerIdsRemoveOk returns a tuple with the OwnerIdsRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdsRemove

`func (o *UpdateStories) SetOwnerIdsRemove(v []string)`

SetOwnerIdsRemove sets OwnerIdsRemove field to given value.

### HasOwnerIdsRemove

`func (o *UpdateStories) HasOwnerIdsRemove() bool`

HasOwnerIdsRemove returns a boolean if a field has been set.

### GetCustomFieldsAdd

`func (o *UpdateStories) GetCustomFieldsAdd() []CustomFieldValueParams`

GetCustomFieldsAdd returns the CustomFieldsAdd field if non-nil, zero value otherwise.

### GetCustomFieldsAddOk

`func (o *UpdateStories) GetCustomFieldsAddOk() (*[]CustomFieldValueParams, bool)`

GetCustomFieldsAddOk returns a tuple with the CustomFieldsAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldsAdd

`func (o *UpdateStories) SetCustomFieldsAdd(v []CustomFieldValueParams)`

SetCustomFieldsAdd sets CustomFieldsAdd field to given value.

### HasCustomFieldsAdd

`func (o *UpdateStories) HasCustomFieldsAdd() bool`

HasCustomFieldsAdd returns a boolean if a field has been set.

### GetProjectId

`func (o *UpdateStories) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateStories) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateStories) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UpdateStories) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *UpdateStories) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *UpdateStories) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLabelsRemove

`func (o *UpdateStories) GetLabelsRemove() []CreateLabelParams`

GetLabelsRemove returns the LabelsRemove field if non-nil, zero value otherwise.

### GetLabelsRemoveOk

`func (o *UpdateStories) GetLabelsRemoveOk() (*[]CreateLabelParams, bool)`

GetLabelsRemoveOk returns a tuple with the LabelsRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsRemove

`func (o *UpdateStories) SetLabelsRemove(v []CreateLabelParams)`

SetLabelsRemove sets LabelsRemove field to given value.

### HasLabelsRemove

`func (o *UpdateStories) HasLabelsRemove() bool`

HasLabelsRemove returns a boolean if a field has been set.

### GetDeadline

`func (o *UpdateStories) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *UpdateStories) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *UpdateStories) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *UpdateStories) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### SetDeadlineNil

`func (o *UpdateStories) SetDeadlineNil(b bool)`

 SetDeadlineNil sets the value for Deadline to be an explicit nil

### UnsetDeadline
`func (o *UpdateStories) UnsetDeadline()`

UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
### GetOwnerIdsAdd

`func (o *UpdateStories) GetOwnerIdsAdd() []string`

GetOwnerIdsAdd returns the OwnerIdsAdd field if non-nil, zero value otherwise.

### GetOwnerIdsAddOk

`func (o *UpdateStories) GetOwnerIdsAddOk() (*[]string, bool)`

GetOwnerIdsAddOk returns a tuple with the OwnerIdsAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIdsAdd

`func (o *UpdateStories) SetOwnerIdsAdd(v []string)`

SetOwnerIdsAdd sets OwnerIdsAdd field to given value.

### HasOwnerIdsAdd

`func (o *UpdateStories) HasOwnerIdsAdd() bool`

HasOwnerIdsAdd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


