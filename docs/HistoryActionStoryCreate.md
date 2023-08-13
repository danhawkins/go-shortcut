# HistoryActionStoryCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The application URL of the Story. | 
**Description** | Pointer to **string** | The description of the Story. | [optional] 
**Started** | Pointer to **bool** | Whether or not the Story has been started. | [optional] 
**EntityType** | **string** | The type of entity referenced. | 
**TaskIds** | Pointer to **[]int64** | An array of Task IDs on this Story. | [optional] 
**StoryType** | **string** | The type of Story; either feature, bug, or chore. | 
**Name** | **string** | The name of the Story. | 
**Completed** | Pointer to **bool** | Whether or not the Story is completed. | [optional] 
**Blocker** | Pointer to **bool** | Whether or not the Story is blocking another Story. | [optional] 
**EpicId** | Pointer to **int64** | The Epic ID for this Story. | [optional] 
**RequestedById** | Pointer to **string** | The ID of the Member that requested the Story. | [optional] 
**IterationId** | Pointer to **NullableInt64** | The Iteration ID the Story is in. | [optional] 
**LabelIds** | Pointer to **[]int64** | An array of Labels IDs attached to the Story. | [optional] 
**GroupId** | Pointer to **string** | The Team IDs for the followers of the Story. | [optional] 
**WorkflowStateId** | Pointer to **int64** | An array of Workflow State IDs attached to the Story. | [optional] 
**ObjectStoryLinkIds** | Pointer to **[]int64** | An array of Story IDs that are the object of a Story Link relationship. | [optional] 
**FollowerIds** | Pointer to **[]string** | An array of Member IDs for the followers of the Story. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of Member IDs that are the owners of the Story. | [optional] 
**CustomFieldValueIds** | Pointer to **[]string** | An array of Custom Field Enum Value ids on this Story. | [optional] 
**Id** | **int64** | The ID of the entity referenced. | 
**Estimate** | Pointer to **int64** | The estimate (or point value) for the Story. | [optional] 
**SubjectStoryLinkIds** | Pointer to **[]int64** | An array of Story IDs that are the subject of a Story Link relationship. | [optional] 
**Action** | **string** | The action of the entity referenced. | 
**Blocked** | Pointer to **bool** | Whether or not the Story is blocked by another Story. | [optional] 
**ProjectId** | Pointer to **int64** | The Project ID of the Story is in. | [optional] 
**Deadline** | Pointer to **string** | The timestamp representing the Story&#39;s deadline. | [optional] 

## Methods

### NewHistoryActionStoryCreate

`func NewHistoryActionStoryCreate(appUrl string, entityType string, storyType string, name string, id int64, action string, ) *HistoryActionStoryCreate`

NewHistoryActionStoryCreate instantiates a new HistoryActionStoryCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionStoryCreateWithDefaults

`func NewHistoryActionStoryCreateWithDefaults() *HistoryActionStoryCreate`

NewHistoryActionStoryCreateWithDefaults instantiates a new HistoryActionStoryCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *HistoryActionStoryCreate) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *HistoryActionStoryCreate) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *HistoryActionStoryCreate) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetDescription

`func (o *HistoryActionStoryCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HistoryActionStoryCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HistoryActionStoryCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HistoryActionStoryCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStarted

`func (o *HistoryActionStoryCreate) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *HistoryActionStoryCreate) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *HistoryActionStoryCreate) SetStarted(v bool)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *HistoryActionStoryCreate) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetEntityType

`func (o *HistoryActionStoryCreate) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionStoryCreate) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionStoryCreate) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetTaskIds

`func (o *HistoryActionStoryCreate) GetTaskIds() []int64`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *HistoryActionStoryCreate) GetTaskIdsOk() (*[]int64, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *HistoryActionStoryCreate) SetTaskIds(v []int64)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *HistoryActionStoryCreate) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.

### GetStoryType

`func (o *HistoryActionStoryCreate) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *HistoryActionStoryCreate) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *HistoryActionStoryCreate) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.


### GetName

`func (o *HistoryActionStoryCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoryActionStoryCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HistoryActionStoryCreate) SetName(v string)`

SetName sets Name field to given value.


### GetCompleted

`func (o *HistoryActionStoryCreate) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *HistoryActionStoryCreate) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *HistoryActionStoryCreate) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *HistoryActionStoryCreate) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetBlocker

`func (o *HistoryActionStoryCreate) GetBlocker() bool`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *HistoryActionStoryCreate) GetBlockerOk() (*bool, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *HistoryActionStoryCreate) SetBlocker(v bool)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *HistoryActionStoryCreate) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetEpicId

`func (o *HistoryActionStoryCreate) GetEpicId() int64`

GetEpicId returns the EpicId field if non-nil, zero value otherwise.

### GetEpicIdOk

`func (o *HistoryActionStoryCreate) GetEpicIdOk() (*int64, bool)`

GetEpicIdOk returns a tuple with the EpicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicId

`func (o *HistoryActionStoryCreate) SetEpicId(v int64)`

SetEpicId sets EpicId field to given value.

### HasEpicId

`func (o *HistoryActionStoryCreate) HasEpicId() bool`

HasEpicId returns a boolean if a field has been set.

### GetRequestedById

`func (o *HistoryActionStoryCreate) GetRequestedById() string`

GetRequestedById returns the RequestedById field if non-nil, zero value otherwise.

### GetRequestedByIdOk

`func (o *HistoryActionStoryCreate) GetRequestedByIdOk() (*string, bool)`

GetRequestedByIdOk returns a tuple with the RequestedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedById

`func (o *HistoryActionStoryCreate) SetRequestedById(v string)`

SetRequestedById sets RequestedById field to given value.

### HasRequestedById

`func (o *HistoryActionStoryCreate) HasRequestedById() bool`

HasRequestedById returns a boolean if a field has been set.

### GetIterationId

`func (o *HistoryActionStoryCreate) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *HistoryActionStoryCreate) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *HistoryActionStoryCreate) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *HistoryActionStoryCreate) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### SetIterationIdNil

`func (o *HistoryActionStoryCreate) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *HistoryActionStoryCreate) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetLabelIds

`func (o *HistoryActionStoryCreate) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *HistoryActionStoryCreate) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *HistoryActionStoryCreate) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *HistoryActionStoryCreate) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetGroupId

`func (o *HistoryActionStoryCreate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *HistoryActionStoryCreate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *HistoryActionStoryCreate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *HistoryActionStoryCreate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetWorkflowStateId

`func (o *HistoryActionStoryCreate) GetWorkflowStateId() int64`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *HistoryActionStoryCreate) GetWorkflowStateIdOk() (*int64, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *HistoryActionStoryCreate) SetWorkflowStateId(v int64)`

SetWorkflowStateId sets WorkflowStateId field to given value.

### HasWorkflowStateId

`func (o *HistoryActionStoryCreate) HasWorkflowStateId() bool`

HasWorkflowStateId returns a boolean if a field has been set.

### GetObjectStoryLinkIds

`func (o *HistoryActionStoryCreate) GetObjectStoryLinkIds() []int64`

GetObjectStoryLinkIds returns the ObjectStoryLinkIds field if non-nil, zero value otherwise.

### GetObjectStoryLinkIdsOk

`func (o *HistoryActionStoryCreate) GetObjectStoryLinkIdsOk() (*[]int64, bool)`

GetObjectStoryLinkIdsOk returns a tuple with the ObjectStoryLinkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoryLinkIds

`func (o *HistoryActionStoryCreate) SetObjectStoryLinkIds(v []int64)`

SetObjectStoryLinkIds sets ObjectStoryLinkIds field to given value.

### HasObjectStoryLinkIds

`func (o *HistoryActionStoryCreate) HasObjectStoryLinkIds() bool`

HasObjectStoryLinkIds returns a boolean if a field has been set.

### GetFollowerIds

`func (o *HistoryActionStoryCreate) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *HistoryActionStoryCreate) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *HistoryActionStoryCreate) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *HistoryActionStoryCreate) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *HistoryActionStoryCreate) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *HistoryActionStoryCreate) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *HistoryActionStoryCreate) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *HistoryActionStoryCreate) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetCustomFieldValueIds

`func (o *HistoryActionStoryCreate) GetCustomFieldValueIds() []string`

GetCustomFieldValueIds returns the CustomFieldValueIds field if non-nil, zero value otherwise.

### GetCustomFieldValueIdsOk

`func (o *HistoryActionStoryCreate) GetCustomFieldValueIdsOk() (*[]string, bool)`

GetCustomFieldValueIdsOk returns a tuple with the CustomFieldValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldValueIds

`func (o *HistoryActionStoryCreate) SetCustomFieldValueIds(v []string)`

SetCustomFieldValueIds sets CustomFieldValueIds field to given value.

### HasCustomFieldValueIds

`func (o *HistoryActionStoryCreate) HasCustomFieldValueIds() bool`

HasCustomFieldValueIds returns a boolean if a field has been set.

### GetId

`func (o *HistoryActionStoryCreate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionStoryCreate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionStoryCreate) SetId(v int64)`

SetId sets Id field to given value.


### GetEstimate

`func (o *HistoryActionStoryCreate) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *HistoryActionStoryCreate) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *HistoryActionStoryCreate) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *HistoryActionStoryCreate) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### GetSubjectStoryLinkIds

`func (o *HistoryActionStoryCreate) GetSubjectStoryLinkIds() []int64`

GetSubjectStoryLinkIds returns the SubjectStoryLinkIds field if non-nil, zero value otherwise.

### GetSubjectStoryLinkIdsOk

`func (o *HistoryActionStoryCreate) GetSubjectStoryLinkIdsOk() (*[]int64, bool)`

GetSubjectStoryLinkIdsOk returns a tuple with the SubjectStoryLinkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectStoryLinkIds

`func (o *HistoryActionStoryCreate) SetSubjectStoryLinkIds(v []int64)`

SetSubjectStoryLinkIds sets SubjectStoryLinkIds field to given value.

### HasSubjectStoryLinkIds

`func (o *HistoryActionStoryCreate) HasSubjectStoryLinkIds() bool`

HasSubjectStoryLinkIds returns a boolean if a field has been set.

### GetAction

`func (o *HistoryActionStoryCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionStoryCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionStoryCreate) SetAction(v string)`

SetAction sets Action field to given value.


### GetBlocked

`func (o *HistoryActionStoryCreate) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *HistoryActionStoryCreate) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *HistoryActionStoryCreate) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *HistoryActionStoryCreate) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetProjectId

`func (o *HistoryActionStoryCreate) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HistoryActionStoryCreate) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HistoryActionStoryCreate) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *HistoryActionStoryCreate) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetDeadline

`func (o *HistoryActionStoryCreate) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *HistoryActionStoryCreate) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *HistoryActionStoryCreate) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *HistoryActionStoryCreate) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


