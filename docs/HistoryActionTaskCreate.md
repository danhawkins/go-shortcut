# HistoryActionTaskCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the Task. | 
**EntityType** | **string** | The type of entity referenced. | 
**MentionIds** | Pointer to **[]string** | An array of Member IDs that represent who has been mentioned in the Task. | [optional] 
**GroupMentionIds** | Pointer to **[]string** | An array of Groups IDs that represent which have been mentioned in the Task. | [optional] 
**OwnerIds** | Pointer to **[]string** | An array of Member IDs that represent the Task&#39;s owners. | [optional] 
**Id** | **int64** | The ID of the entity referenced. | 
**Action** | **string** | The action of the entity referenced. | 
**Complete** | **bool** | Whether or not the Task is complete. | 
**Deadline** | Pointer to **string** | A timestamp that represent&#39;s the Task&#39;s deadline. | [optional] 

## Methods

### NewHistoryActionTaskCreate

`func NewHistoryActionTaskCreate(description string, entityType string, id int64, action string, complete bool, ) *HistoryActionTaskCreate`

NewHistoryActionTaskCreate instantiates a new HistoryActionTaskCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionTaskCreateWithDefaults

`func NewHistoryActionTaskCreateWithDefaults() *HistoryActionTaskCreate`

NewHistoryActionTaskCreateWithDefaults instantiates a new HistoryActionTaskCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HistoryActionTaskCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HistoryActionTaskCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HistoryActionTaskCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntityType

`func (o *HistoryActionTaskCreate) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionTaskCreate) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionTaskCreate) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetMentionIds

`func (o *HistoryActionTaskCreate) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *HistoryActionTaskCreate) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *HistoryActionTaskCreate) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.

### HasMentionIds

`func (o *HistoryActionTaskCreate) HasMentionIds() bool`

HasMentionIds returns a boolean if a field has been set.

### GetGroupMentionIds

`func (o *HistoryActionTaskCreate) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *HistoryActionTaskCreate) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *HistoryActionTaskCreate) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.

### HasGroupMentionIds

`func (o *HistoryActionTaskCreate) HasGroupMentionIds() bool`

HasGroupMentionIds returns a boolean if a field has been set.

### GetOwnerIds

`func (o *HistoryActionTaskCreate) GetOwnerIds() []string`

GetOwnerIds returns the OwnerIds field if non-nil, zero value otherwise.

### GetOwnerIdsOk

`func (o *HistoryActionTaskCreate) GetOwnerIdsOk() (*[]string, bool)`

GetOwnerIdsOk returns a tuple with the OwnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerIds

`func (o *HistoryActionTaskCreate) SetOwnerIds(v []string)`

SetOwnerIds sets OwnerIds field to given value.

### HasOwnerIds

`func (o *HistoryActionTaskCreate) HasOwnerIds() bool`

HasOwnerIds returns a boolean if a field has been set.

### GetId

`func (o *HistoryActionTaskCreate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionTaskCreate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionTaskCreate) SetId(v int64)`

SetId sets Id field to given value.


### GetAction

`func (o *HistoryActionTaskCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionTaskCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionTaskCreate) SetAction(v string)`

SetAction sets Action field to given value.


### GetComplete

`func (o *HistoryActionTaskCreate) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *HistoryActionTaskCreate) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *HistoryActionTaskCreate) SetComplete(v bool)`

SetComplete sets Complete field to given value.


### GetDeadline

`func (o *HistoryActionTaskCreate) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *HistoryActionTaskCreate) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *HistoryActionTaskCreate) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *HistoryActionTaskCreate) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


