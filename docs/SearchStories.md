# SearchStories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** | A true/false boolean indicating whether the Story is in archived state. | [optional] 
**OwnerId** | Pointer to **NullableString** | An array of UUIDs for any Users who may be Owners of the Stories. | [optional] 
**StoryType** | Pointer to **string** | The type of Stories that you want returned. | [optional] 
**EpicIds** | Pointer to **[]int64** | The Epic IDs that may be associated with the Stories. | [optional] 
**ProjectIds** | Pointer to **[]int64** | The IDs for the Projects the Stories may be assigned to. | [optional] 
**UpdatedAtEnd** | Pointer to **time.Time** | Stories should have been updated before this date. | [optional] 
**CompletedAtEnd** | Pointer to **time.Time** | Stories should have been completed before this date. | [optional] 
**WorkflowStateTypes** | Pointer to **[]string** | The type of Workflow State the Stories may be in. | [optional] 
**DeadlineEnd** | Pointer to **time.Time** | Stories should have a deadline before this date. | [optional] 
**CreatedAtStart** | Pointer to **time.Time** | Stories should have been created after this date. | [optional] 
**EpicId** | Pointer to **NullableInt64** | The Epic IDs that may be associated with the Stories. | [optional] 
**LabelName** | Pointer to **string** | The name of any associated Labels. | [optional] 
**RequestedById** | Pointer to **string** | The UUID of any Users who may have requested the Stories. | [optional] 
**IterationId** | Pointer to **NullableInt64** | The Iteration ID that may be associated with the Stories. | [optional] 
**LabelIds** | Pointer to **[]int64** | The Label IDs that may be associated with the Stories. | [optional] 
**GroupId** | Pointer to **NullableString** | The Group ID that is associated with the Stories | [optional] 
**WorkflowStateId** | Pointer to **int64** | The unique IDs of the specific Workflow States that the Stories should be in. | [optional] 
**IterationIds** | Pointer to **[]int64** | The Iteration IDs that may be associated with the Stories. | [optional] 
**CreatedAtEnd** | Pointer to **time.Time** | Stories should have been created before this date. | [optional] 
**DeadlineStart** | Pointer to **time.Time** | Stories should have a deadline after this date. | [optional] 
**GroupIds** | Pointer to **[]string** | The Group IDs that are associated with the Stories | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of UUIDs for any Users who may be Owners of the Stories. | [optional] 
**ExternalId** | Pointer to **string** | An ID or URL that references an external resource. Useful during imports. | [optional] 
**IncludesDescription** | Pointer to **bool** | Whether to include the story description in the response. | [optional] 
**Estimate** | Pointer to **int64** | The number of estimate points associate with the Stories. | [optional] 
**ProjectId** | Pointer to **NullableInt64** | The IDs for the Projects the Stories may be assigned to. | [optional] 
**CompletedAtStart** | Pointer to **time.Time** | Stories should have been completed after this date. | [optional] 
**UpdatedAtStart** | Pointer to **time.Time** | Stories should have been updated after this date. | [optional] 

## Methods

### NewSearchStories

`func NewSearchStories() *SearchStories`

NewSearchStories instantiates a new SearchStories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStoriesWithDefaults

`func NewSearchStoriesWithDefaults() *SearchStories`

NewSearchStoriesWithDefaults instantiates a new SearchStories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *SearchStories) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *SearchStories) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *SearchStories) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *SearchStories) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetOwnerId

`func (o *SearchStories) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *SearchStories) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *SearchStories) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *SearchStories) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *SearchStories) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *SearchStories) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetStoryType

`func (o *SearchStories) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *SearchStories) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *SearchStories) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.

### HasStoryType

`func (o *SearchStories) HasStoryType() bool`

HasStoryType returns a boolean if a field has been set.

### GetEpicIds

`func (o *SearchStories) GetEpicIds() []int64`

GetEpicIds returns the EpicIds field if non-nil, zero value otherwise.

### GetEpicIdsOk

`func (o *SearchStories) GetEpicIdsOk() (*[]int64, bool)`

GetEpicIdsOk returns a tuple with the EpicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicIds

`func (o *SearchStories) SetEpicIds(v []int64)`

SetEpicIds sets EpicIds field to given value.

### HasEpicIds

`func (o *SearchStories) HasEpicIds() bool`

HasEpicIds returns a boolean if a field has been set.

### GetProjectIds

`func (o *SearchStories) GetProjectIds() []*int64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *SearchStories) GetProjectIdsOk() (*[]*int64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *SearchStories) SetProjectIds(v []*int64)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *SearchStories) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetUpdatedAtEnd

`func (o *SearchStories) GetUpdatedAtEnd() time.Time`

GetUpdatedAtEnd returns the UpdatedAtEnd field if non-nil, zero value otherwise.

### GetUpdatedAtEndOk

`func (o *SearchStories) GetUpdatedAtEndOk() (*time.Time, bool)`

GetUpdatedAtEndOk returns a tuple with the UpdatedAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtEnd

`func (o *SearchStories) SetUpdatedAtEnd(v time.Time)`

SetUpdatedAtEnd sets UpdatedAtEnd field to given value.

### HasUpdatedAtEnd

`func (o *SearchStories) HasUpdatedAtEnd() bool`

HasUpdatedAtEnd returns a boolean if a field has been set.

### GetCompletedAtEnd

`func (o *SearchStories) GetCompletedAtEnd() time.Time`

GetCompletedAtEnd returns the CompletedAtEnd field if non-nil, zero value otherwise.

### GetCompletedAtEndOk

`func (o *SearchStories) GetCompletedAtEndOk() (*time.Time, bool)`

GetCompletedAtEndOk returns a tuple with the CompletedAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtEnd

`func (o *SearchStories) SetCompletedAtEnd(v time.Time)`

SetCompletedAtEnd sets CompletedAtEnd field to given value.

### HasCompletedAtEnd

`func (o *SearchStories) HasCompletedAtEnd() bool`

HasCompletedAtEnd returns a boolean if a field has been set.

### GetWorkflowStateTypes

`func (o *SearchStories) GetWorkflowStateTypes() []string`

GetWorkflowStateTypes returns the WorkflowStateTypes field if non-nil, zero value otherwise.

### GetWorkflowStateTypesOk

`func (o *SearchStories) GetWorkflowStateTypesOk() (*[]string, bool)`

GetWorkflowStateTypesOk returns a tuple with the WorkflowStateTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateTypes

`func (o *SearchStories) SetWorkflowStateTypes(v []string)`

SetWorkflowStateTypes sets WorkflowStateTypes field to given value.

### HasWorkflowStateTypes

`func (o *SearchStories) HasWorkflowStateTypes() bool`

HasWorkflowStateTypes returns a boolean if a field has been set.

### GetDeadlineEnd

`func (o *SearchStories) GetDeadlineEnd() time.Time`

GetDeadlineEnd returns the DeadlineEnd field if non-nil, zero value otherwise.

### GetDeadlineEndOk

`func (o *SearchStories) GetDeadlineEndOk() (*time.Time, bool)`

GetDeadlineEndOk returns a tuple with the DeadlineEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineEnd

`func (o *SearchStories) SetDeadlineEnd(v time.Time)`

SetDeadlineEnd sets DeadlineEnd field to given value.

### HasDeadlineEnd

`func (o *SearchStories) HasDeadlineEnd() bool`

HasDeadlineEnd returns a boolean if a field has been set.

### GetCreatedAtStart

`func (o *SearchStories) GetCreatedAtStart() time.Time`

GetCreatedAtStart returns the CreatedAtStart field if non-nil, zero value otherwise.

### GetCreatedAtStartOk

`func (o *SearchStories) GetCreatedAtStartOk() (*time.Time, bool)`

GetCreatedAtStartOk returns a tuple with the CreatedAtStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtStart

`func (o *SearchStories) SetCreatedAtStart(v time.Time)`

SetCreatedAtStart sets CreatedAtStart field to given value.

### HasCreatedAtStart

`func (o *SearchStories) HasCreatedAtStart() bool`

HasCreatedAtStart returns a boolean if a field has been set.

### GetEpicId

`func (o *SearchStories) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *SearchStories) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *SearchStories) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.

### HasEpicId

`func (o *SearchStories) HasEpicId() bool`

HasEpicId returns a boolean if a field has been set.

### SetEpicIdNil

`func (o *SearchStories) SetEpicIdNil(b bool)`

 SetEpicIdNil sets the value for EpicId to be an explicit nil

### UnsetEpicId
`func (o *SearchStories) UnsetEpicId()`

UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
### GetLabelName

`func (o *SearchStories) GetLabelName() string`

GetLabelName returns the LabelName field if non-nil, zero value otherwise.

### GetLabelNameOk

`func (o *SearchStories) GetLabelNameOk() (*string, bool)`

GetLabelNameOk returns a tuple with the LabelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelName

`func (o *SearchStories) SetLabelName(v string)`

SetLabelName sets LabelName field to given value.

### HasLabelName

`func (o *SearchStories) HasLabelName() bool`

HasLabelName returns a boolean if a field has been set.

### GetRequestedById

`func (o *SearchStories) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *SearchStories) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *SearchStories) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.

### HasRequestedById

`func (o *SearchStories) HasRequestedById() bool`

HasRequestedById returns a boolean if a field has been set.

### GetIterationId

`func (o *SearchStories) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *SearchStories) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *SearchStories) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *SearchStories) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### SetIterationIdNil

`func (o *SearchStories) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *SearchStories) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetLabelIds

`func (o *SearchStories) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *SearchStories) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *SearchStories) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *SearchStories) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetGroupId

`func (o *SearchStories) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SearchStories) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SearchStories) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SearchStories) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *SearchStories) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *SearchStories) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetWorkflowStateId

`func (o *SearchStories) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *SearchStories) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *SearchStories) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.

### HasWorkflowStateId

`func (o *SearchStories) HasWorkflowStateId() bool`

HasWorkflowStateId returns a boolean if a field has been set.

### GetIterationIds

`func (o *SearchStories) GetIterationIds() []int64`

GetIterationIds returns the IterationIds field if non-nil, zero value otherwise.

### GetIterationIdsOk

`func (o *SearchStories) GetIterationIdsOk() (*[]int64, bool)`

GetIterationIdsOk returns a tuple with the IterationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationIds

`func (o *SearchStories) SetIterationIds(v []int64)`

SetIterationIds sets IterationIds field to given value.

### HasIterationIds

`func (o *SearchStories) HasIterationIds() bool`

HasIterationIds returns a boolean if a field has been set.

### GetCreatedAtEnd

`func (o *SearchStories) GetCreatedAtEnd() time.Time`

GetCreatedAtEnd returns the CreatedAtEnd field if non-nil, zero value otherwise.

### GetCreatedAtEndOk

`func (o *SearchStories) GetCreatedAtEndOk() (*time.Time, bool)`

GetCreatedAtEndOk returns a tuple with the CreatedAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtEnd

`func (o *SearchStories) SetCreatedAtEnd(v time.Time)`

SetCreatedAtEnd sets CreatedAtEnd field to given value.

### HasCreatedAtEnd

`func (o *SearchStories) HasCreatedAtEnd() bool`

HasCreatedAtEnd returns a boolean if a field has been set.

### GetDeadlineStart

`func (o *SearchStories) GetDeadlineStart() time.Time`

GetDeadlineStart returns the DeadlineStart field if non-nil, zero value otherwise.

### GetDeadlineStartOk

`func (o *SearchStories) GetDeadlineStartOk() (*time.Time, bool)`

GetDeadlineStartOk returns a tuple with the DeadlineStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineStart

`func (o *SearchStories) SetDeadlineStart(v time.Time)`

SetDeadlineStart sets DeadlineStart field to given value.

### HasDeadlineStart

`func (o *SearchStories) HasDeadlineStart() bool`

HasDeadlineStart returns a boolean if a field has been set.

### GetGroupIds

`func (o *SearchStories) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *SearchStories) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *SearchStories) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *SearchStories) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *SearchStories) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *SearchStories) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *SearchStories) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *SearchStories) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetExternalId

`func (o *SearchStories) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SearchStories) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SearchStories) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SearchStories) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIncludesDescription

`func (o *SearchStories) GetIncludesDescription() bool`

GetIncludesDescription returns the IncludesDescription field if non-nil, zero value otherwise.

### GetIncludesDescriptionOk

`func (o *SearchStories) GetIncludesDescriptionOk() (*bool, bool)`

GetIncludesDescriptionOk returns a tuple with the IncludesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesDescription

`func (o *SearchStories) SetIncludesDescription(v bool)`

SetIncludesDescription sets IncludesDescription field to given value.

### HasIncludesDescription

`func (o *SearchStories) HasIncludesDescription() bool`

HasIncludesDescription returns a boolean if a field has been set.

### GetEstimate

`func (o *SearchStories) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *SearchStories) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *SearchStories) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *SearchStories) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### GetProjectId

`func (o *SearchStories) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SearchStories) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SearchStories) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SearchStories) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *SearchStories) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *SearchStories) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetCompletedAtStart

`func (o *SearchStories) GetCompletedAtStart() time.Time`

GetCompletedAtStart returns the CompletedAtStart field if non-nil, zero value otherwise.

### GetCompletedAtStartOk

`func (o *SearchStories) GetCompletedAtStartOk() (*time.Time, bool)`

GetCompletedAtStartOk returns a tuple with the CompletedAtStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtStart

`func (o *SearchStories) SetCompletedAtStart(v time.Time)`

SetCompletedAtStart sets CompletedAtStart field to given value.

### HasCompletedAtStart

`func (o *SearchStories) HasCompletedAtStart() bool`

HasCompletedAtStart returns a boolean if a field has been set.

### GetUpdatedAtStart

`func (o *SearchStories) GetUpdatedAtStart() time.Time`

GetUpdatedAtStart returns the UpdatedAtStart field if non-nil, zero value otherwise.

### GetUpdatedAtStartOk

`func (o *SearchStories) GetUpdatedAtStartOk() (*time.Time, bool)`

GetUpdatedAtStartOk returns a tuple with the UpdatedAtStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtStart

`func (o *SearchStories) SetUpdatedAtStart(v time.Time)`

SetUpdatedAtStart sets UpdatedAtStart field to given value.

### HasUpdatedAtStart

`func (o *SearchStories) HasUpdatedAtStart() bool`

HasUpdatedAtStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


