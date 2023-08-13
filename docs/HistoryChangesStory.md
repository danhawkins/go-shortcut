# HistoryChangesStory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**StoryHistoryChangeOldNewStr**](StoryHistoryChangeOldNewStr.md) |  | [optional] 
**Archived** | Pointer to [**StoryHistoryChangeOldNewBool**](StoryHistoryChangeOldNewBool.md) |  | [optional] 
**Started** | Pointer to [**StoryHistoryChangeOldNewBool**](StoryHistoryChangeOldNewBool.md) |  | [optional] 
**TaskIds** | Pointer to [**StoryHistoryChangeAddsRemovesInt**](StoryHistoryChangeAddsRemovesInt.md) |  | [optional] 
**MentionIds** | Pointer to [**StoryHistoryChangeAddsRemovesUuid**](StoryHistoryChangeAddsRemovesUuid.md) |  | [optional] 
**StoryType** | Pointer to [**StoryHistoryChangeOldNewStr**](StoryHistoryChangeOldNewStr.md) |  | [optional] 
**Name** | Pointer to [**StoryHistoryChangeOldNewStr**](StoryHistoryChangeOldNewStr.md) |  | [optional] 
**Completed** | Pointer to [**StoryHistoryChangeOldNewBool**](StoryHistoryChangeOldNewBool.md) |  | [optional] 
**Blocker** | Pointer to [**StoryHistoryChangeOldNewBool**](StoryHistoryChangeOldNewBool.md) |  | [optional] 
**EpicId** | Pointer to [**StoryHistoryChangeOldNewInt**](StoryHistoryChangeOldNewInt.md) |  | [optional] 
**BranchIds** | Pointer to [**StoryHistoryChangeAddsRemovesInt**](StoryHistoryChangeAddsRemovesInt.md) |  | [optional] 
**CommitIds** | Pointer to [**StoryHistoryChangeAddsRemovesInt**](StoryHistoryChangeAddsRemovesInt.md) |  | [optional] 
**RequestedById** | Pointer to [**StoryHistoryChangeOldNewUuid**](StoryHistoryChangeOldNewUuid.md) |  | [optional] 
**IterationId** | Pointer to [**StoryHistoryChangeOldNewInt**](StoryHistoryChangeOldNewInt.md) |  | [optional] 
**LabelIds** | Pointer to [**StoryHistoryChangeAddsRemovesInt**](StoryHistoryChangeAddsRemovesInt.md) |  | [optional] 
**GroupId** | Pointer to [**StoryHistoryChangeOldNewUuid**](StoryHistoryChangeOldNewUuid.md) |  | [optional] 
**WorkflowStateId** | Pointer to [**StoryHistoryChangeOldNewInt**](StoryHistoryChangeOldNewInt.md) |  | [optional] 
**ObjectStoryLinkIds** | Pointer to [**StoryHistoryChangeAddsRemovesInt**](StoryHistoryChangeAddsRemovesInt.md) |  | [optional] 
**FollowerIds** | Pointer to [**StoryHistoryChangeAddsRemovesUuid**](StoryHistoryChangeAddsRemovesUuid.md) |  | [optional] 
**OwnerIds** | Pointer to [**StoryHistoryChangeAddsRemovesUuid**](StoryHistoryChangeAddsRemovesUuid.md) |  | [optional] 
**CustomFieldValueIds** | Pointer to [**StoryHistoryChangeAddsRemovesUuid**](StoryHistoryChangeAddsRemovesUuid.md) |  | [optional] 
**Estimate** | Pointer to [**StoryHistoryChangeOldNewInt**](StoryHistoryChangeOldNewInt.md) |  | [optional] 
**SubjectStoryLinkIds** | Pointer to [**StoryHistoryChangeAddsRemovesInt**](StoryHistoryChangeAddsRemovesInt.md) |  | [optional] 
**Blocked** | Pointer to [**StoryHistoryChangeOldNewBool**](StoryHistoryChangeOldNewBool.md) |  | [optional] 
**ProjectId** | Pointer to [**StoryHistoryChangeOldNewInt**](StoryHistoryChangeOldNewInt.md) |  | [optional] 
**Deadline** | Pointer to [**StoryHistoryChangeOldNewStr**](StoryHistoryChangeOldNewStr.md) |  | [optional] 

## Methods

### NewHistoryChangesStory

`func NewHistoryChangesStory() *HistoryChangesStory`

NewHistoryChangesStory instantiates a new HistoryChangesStory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryChangesStoryWithDefaults

`func NewHistoryChangesStoryWithDefaults() *HistoryChangesStory`

NewHistoryChangesStoryWithDefaults instantiates a new HistoryChangesStory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HistoryChangesStory) GetDescription() StoryHistoryChangeOldNewStr`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HistoryChangesStory) GetDescriptionOk() (*StoryHistoryChangeOldNewStr, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HistoryChangesStory) SetDescription(v StoryHistoryChangeOldNewStr)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HistoryChangesStory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *HistoryChangesStory) GetArchived() StoryHistoryChangeOldNewBool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *HistoryChangesStory) GetArchivedOk() (*StoryHistoryChangeOldNewBool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *HistoryChangesStory) SetArchived(v StoryHistoryChangeOldNewBool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *HistoryChangesStory) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetStarted

`func (o *HistoryChangesStory) GetStarted() StoryHistoryChangeOldNewBool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *HistoryChangesStory) GetStartedOk() (*StoryHistoryChangeOldNewBool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *HistoryChangesStory) SetStarted(v StoryHistoryChangeOldNewBool)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *HistoryChangesStory) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetTaskIds

`func (o *HistoryChangesStory) GetTaskIds() StoryHistoryChangeAddsRemovesInt`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *HistoryChangesStory) GetTaskIdsOk() (*StoryHistoryChangeAddsRemovesInt, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *HistoryChangesStory) SetTaskIds(v StoryHistoryChangeAddsRemovesInt)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *HistoryChangesStory) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### GetMentionIds

`func (o *HistoryChangesStory) GetMentionIds() StoryHistoryChangeAddsRemovesUuid`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *HistoryChangesStory) GetMentionIdsOk() (*StoryHistoryChangeAddsRemovesUuid, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *HistoryChangesStory) SetMentionIds(v StoryHistoryChangeAddsRemovesUuid)`

SetMentionIds sets MentionIds field to given value.

### HasMentionIds

`func (o *HistoryChangesStory) HasMentionIds() bool`

HasMentionIds returns a boolean if a field has been set.

### GetStoryType

`func (o *HistoryChangesStory) GetStoryType() StoryHistoryChangeOldNewStr`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *HistoryChangesStory) GetStoryTypeOk() (*StoryHistoryChangeOldNewStr, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *HistoryChangesStory) SetStoryType(v StoryHistoryChangeOldNewStr)`

SetStoryType sets StoryType field to given value.

### HasStoryType

`func (o *HistoryChangesStory) HasStoryType() bool`

HasStoryType returns a boolean if a field has been set.

### GetName

`func (o *HistoryChangesStory) GetName() StoryHistoryChangeOldNewStr`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoryChangesStory) GetNameOk() (*StoryHistoryChangeOldNewStr, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HistoryChangesStory) SetName(v StoryHistoryChangeOldNewStr)`

SetName sets Name field to given value.

### HasName

`func (o *HistoryChangesStory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCompleted

`func (o *HistoryChangesStory) GetCompleted() StoryHistoryChangeOldNewBool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *HistoryChangesStory) GetCompletedOk() (*StoryHistoryChangeOldNewBool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *HistoryChangesStory) SetCompleted(v StoryHistoryChangeOldNewBool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *HistoryChangesStory) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetBlocker

`func (o *HistoryChangesStory) GetBlocker() StoryHistoryChangeOldNewBool`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *HistoryChangesStory) GetBlockerOk() (*StoryHistoryChangeOldNewBool, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *HistoryChangesStory) SetBlocker(v StoryHistoryChangeOldNewBool)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *HistoryChangesStory) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetEpicId

`func (o *HistoryChangesStory) GetEpicId() StoryHistoryChangeOldNewInt`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *HistoryChangesStory) GetEpicIdOk() (*StoryHistoryChangeOldNewInt, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *HistoryChangesStory) SetEpicId(v StoryHistoryChangeOldNewInt)`

SetEpicId sets EpicId field to given value.

### HasEpicId

`func (o *HistoryChangesStory) HasEpicId() bool`

HasEpicId returns a boolean if a field has been set.

### GetBranchIds

`func (o *HistoryChangesStory) GetBranchIds() StoryHistoryChangeAddsRemovesInt`

GetBranchIds returns the BranchIds field if non-nil, zero value otherwise.

### GetBranchIdsOk

`func (o *HistoryChangesStory) GetBranchIdsOk() (*StoryHistoryChangeAddsRemovesInt, bool)`

GetBranchIdsOk returns a tuple with the BranchIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchIds

`func (o *HistoryChangesStory) SetBranchIds(v StoryHistoryChangeAddsRemovesInt)`

SetBranchIds sets BranchIds field to given value.

### HasBranchIds

`func (o *HistoryChangesStory) HasBranchIds() bool`

HasBranchIds returns a boolean if a field has been set.

### GetCommitIds

`func (o *HistoryChangesStory) GetCommitIds() StoryHistoryChangeAddsRemovesInt`

GetCommitIds returns the CommitIds field if non-nil, zero value otherwise.

### GetCommitIdsOk

`func (o *HistoryChangesStory) GetCommitIdsOk() (*StoryHistoryChangeAddsRemovesInt, bool)`

GetCommitIdsOk returns a tuple with the CommitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitIds

`func (o *HistoryChangesStory) SetCommitIds(v StoryHistoryChangeAddsRemovesInt)`

SetCommitIds sets CommitIds field to given value.

### HasCommitIds

`func (o *HistoryChangesStory) HasCommitIds() bool`

HasCommitIds returns a boolean if a field has been set.

### GetRequestedById

`func (o *HistoryChangesStory) GetRequestedById() StoryHistoryChangeOldNewUuid`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *HistoryChangesStory) GetRequestedByIdOk() (*StoryHistoryChangeOldNewUuid, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *HistoryChangesStory) SetRequestedById(v StoryHistoryChangeOldNewUuid)`

SetRequestedById sets RequestedById field to given value.

### HasRequestedById

`func (o *HistoryChangesStory) HasRequestedById() bool`

HasRequestedById returns a boolean if a field has been set.

### GetIterationId

`func (o *HistoryChangesStory) GetIterationId() StoryHistoryChangeOldNewInt`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *HistoryChangesStory) GetIterationIdOk() (*StoryHistoryChangeOldNewInt, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *HistoryChangesStory) SetIterationId(v StoryHistoryChangeOldNewInt)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *HistoryChangesStory) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### GetLabelIds

`func (o *HistoryChangesStory) GetLabelIds() StoryHistoryChangeAddsRemovesInt`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *HistoryChangesStory) GetLabelIdsOk() (*StoryHistoryChangeAddsRemovesInt, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *HistoryChangesStory) SetLabelIds(v StoryHistoryChangeAddsRemovesInt)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *HistoryChangesStory) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetGroupId

`func (o *HistoryChangesStory) GetGroupId() StoryHistoryChangeOldNewUuid`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *HistoryChangesStory) GetGroupIdOk() (*StoryHistoryChangeOldNewUuid, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *HistoryChangesStory) SetGroupId(v StoryHistoryChangeOldNewUuid)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *HistoryChangesStory) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetWorkflowStateId

`func (o *HistoryChangesStory) GetWorkflowStateId() StoryHistoryChangeOldNewInt`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *HistoryChangesStory) GetWorkflowStateIdOk() (*StoryHistoryChangeOldNewInt, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *HistoryChangesStory) SetWorkflowStateId(v StoryHistoryChangeOldNewInt)`

SetWorkflowStateId sets WorkflowStateId field to given value.

### HasWorkflowStateId

`func (o *HistoryChangesStory) HasWorkflowStateId() bool`

HasWorkflowStateId returns a boolean if a field has been set.

### GetObjectStoryLinkIds

`func (o *HistoryChangesStory) GetObjectStoryLinkIds() StoryHistoryChangeAddsRemovesInt`

GetObjectStoryLinkIds returns the ObjectStoryLinkIds field if non-nil, zero value otherwise.

### GetObjectStoryLinkIdsOk

`func (o *HistoryChangesStory) GetObjectStoryLinkIdsOk() (*StoryHistoryChangeAddsRemovesInt, bool)`

GetObjectStoryLinkIdsOk returns a tuple with the ObjectStoryLinkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoryLinkIds

`func (o *HistoryChangesStory) SetObjectStoryLinkIds(v StoryHistoryChangeAddsRemovesInt)`

SetObjectStoryLinkIds sets ObjectStoryLinkIds field to given value.

### HasObjectStoryLinkIds

`func (o *HistoryChangesStory) HasObjectStoryLinkIds() bool`

HasObjectStoryLinkIds returns a boolean if a field has been set.

### GetFollowerIds

`func (o *HistoryChangesStory) GetFollowerIds() StoryHistoryChangeAddsRemovesUuid`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *HistoryChangesStory) GetFollowerIdsOk() (*StoryHistoryChangeAddsRemovesUuid, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *HistoryChangesStory) SetFollowerIds(v StoryHistoryChangeAddsRemovesUuid)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *HistoryChangesStory) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *HistoryChangesStory) GetOwnerIds() StoryHistoryChangeAddsRemovesUuid`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *HistoryChangesStory) GetOwnerIdsOk() (*StoryHistoryChangeAddsRemovesUuid, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *HistoryChangesStory) SetOwnerIds(v StoryHistoryChangeAddsRemovesUuid)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *HistoryChangesStory) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetCustomFieldValueIds

`func (o *HistoryChangesStory) GetCustomFieldValueIds() StoryHistoryChangeAddsRemovesUuid`

GetCustomFieldValueIds returns the CustomFieldValueIds field if non-nil, zero value otherwise.

### GetCustomFieldValueIdsOk

`func (o *HistoryChangesStory) GetCustomFieldValueIdsOk() (*StoryHistoryChangeAddsRemovesUuid, bool)`

GetCustomFieldValueIdsOk returns a tuple with the CustomFieldValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldValueIds

`func (o *HistoryChangesStory) SetCustomFieldValueIds(v StoryHistoryChangeAddsRemovesUuid)`

SetCustomFieldValueIds sets CustomFieldValueIds field to given value.

### HasCustomFieldValueIds

`func (o *HistoryChangesStory) HasCustomFieldValueIds() bool`

HasCustomFieldValueIds returns a boolean if a field has been set.

### GetEstimate

`func (o *HistoryChangesStory) GetEstimate() StoryHistoryChangeOldNewInt`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *HistoryChangesStory) GetEstimateOk() (*StoryHistoryChangeOldNewInt, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *HistoryChangesStory) SetEstimate(v StoryHistoryChangeOldNewInt)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *HistoryChangesStory) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### GetSubjectStoryLinkIds

`func (o *HistoryChangesStory) GetSubjectStoryLinkIds() StoryHistoryChangeAddsRemovesInt`

GetSubjectStoryLinkIds returns the SubjectStoryLinkIds field if non-nil, zero value otherwise.

### GetSubjectStoryLinkIdsOk

`func (o *HistoryChangesStory) GetSubjectStoryLinkIdsOk() (*StoryHistoryChangeAddsRemovesInt, bool)`

GetSubjectStoryLinkIdsOk returns a tuple with the SubjectStoryLinkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectStoryLinkIds

`func (o *HistoryChangesStory) SetSubjectStoryLinkIds(v StoryHistoryChangeAddsRemovesInt)`

SetSubjectStoryLinkIds sets SubjectStoryLinkIds field to given value.

### HasSubjectStoryLinkIds

`func (o *HistoryChangesStory) HasSubjectStoryLinkIds() bool`

HasSubjectStoryLinkIds returns a boolean if a field has been set.

### GetBlocked

`func (o *HistoryChangesStory) GetBlocked() StoryHistoryChangeOldNewBool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *HistoryChangesStory) GetBlockedOk() (*StoryHistoryChangeOldNewBool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *HistoryChangesStory) SetBlocked(v StoryHistoryChangeOldNewBool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *HistoryChangesStory) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetProjectId

`func (o *HistoryChangesStory) GetProjectId() StoryHistoryChangeOldNewInt`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HistoryChangesStory) GetProjectIdOk() (*StoryHistoryChangeOldNewInt, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HistoryChangesStory) SetProjectId(v StoryHistoryChangeOldNewInt)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *HistoryChangesStory) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetDeadline

`func (o *HistoryChangesStory) GetDeadline() StoryHistoryChangeOldNewStr`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *HistoryChangesStory) GetDeadlineOk() (*StoryHistoryChangeOldNewStr, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *HistoryChangesStory) SetDeadline(v StoryHistoryChangeOldNewStr)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *HistoryChangesStory) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


