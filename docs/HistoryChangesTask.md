# HistoryChangesTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | Pointer to [**StoryHistoryChangeOldNewBool**](StoryHistoryChangeOldNewBool.md) |  | [optional] 
**Description** | Pointer to [**StoryHistoryChangeOldNewStr**](StoryHistoryChangeOldNewStr.md) |  | [optional] 
**MentionIds** | Pointer to [**StoryHistoryChangeAddsRemovesUuid**](StoryHistoryChangeAddsRemovesUuid.md) |  | [optional] 
**OwnerIds** | Pointer to [**StoryHistoryChangeAddsRemovesUuid**](StoryHistoryChangeAddsRemovesUuid.md) |  | [optional] 

## Methods

### NewHistoryChangesTask

`func NewHistoryChangesTask() *HistoryChangesTask`

NewHistoryChangesTask instantiates a new HistoryChangesTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryChangesTaskWithDefaults

`func NewHistoryChangesTaskWithDefaults() *HistoryChangesTask`

NewHistoryChangesTaskWithDefaults instantiates a new HistoryChangesTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplete

`func (o *HistoryChangesTask) GetComplete() StoryHistoryChangeOldNewBool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *HistoryChangesTask) GetCompleteOk() (*StoryHistoryChangeOldNewBool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *HistoryChangesTask) SetComplete(v StoryHistoryChangeOldNewBool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *HistoryChangesTask) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetDescription

`func (o *HistoryChangesTask) GetDescription() StoryHistoryChangeOldNewStr`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HistoryChangesTask) GetDescriptionOk() (*StoryHistoryChangeOldNewStr, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HistoryChangesTask) SetDescription(v StoryHistoryChangeOldNewStr)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HistoryChangesTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMentionIds

`func (o *HistoryChangesTask) GetMentionIds() StoryHistoryChangeAddsRemovesUuid`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *HistoryChangesTask) GetMentionIdsOk() (*StoryHistoryChangeAddsRemovesUuid, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *HistoryChangesTask) SetMentionIds(v StoryHistoryChangeAddsRemovesUuid)`

SetMentionIds sets MentionIds field to given value.

### HasMentionIds

`func (o *HistoryChangesTask) HasMentionIds() bool`

HasMentionIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *HistoryChangesTask) GetOwnerIds() StoryHistoryChangeAddsRemovesUuid`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *HistoryChangesTask) GetOwnerIdsOk() (*StoryHistoryChangeAddsRemovesUuid, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *HistoryChangesTask) SetOwnerIds(v StoryHistoryChangeAddsRemovesUuid)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *HistoryChangesTask) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


