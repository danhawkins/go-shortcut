# CreateGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the Group. | [optional] 
**MemberIds** | Pointer to **[]string** | The Member ids to add to this Group. | [optional] 
**WorkflowIds** | Pointer to **[]int64** | The Workflow ids to add to the Group. | [optional] 
**Name** | **string** | The name of this Group. | 
**MentionName** | **string** | The mention name of this Group. | 
**Color** | Pointer to **string** | The color you wish to use for the Group in the system. | [optional] 
**ColorKey** | Pointer to **string** | The color key you wish to use for the Group in the system. | [optional] 
**DisplayIconId** | Pointer to **string** | The Icon id for the avatar of this Group. | [optional] 

## Methods

### NewCreateGroup

`func NewCreateGroup(name string, mentionName string, ) *CreateGroup`

NewCreateGroup instantiates a new CreateGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupWithDefaults

`func NewCreateGroupWithDefaults() *CreateGroup`

NewCreateGroupWithDefaults instantiates a new CreateGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMemberIds

`func (o *CreateGroup) GetMemberIds() []string`

GetMemberIds returns the MemberIds field if non-nil, zero value otherwise.

### GetMemberIdsOk

`func (o *CreateGroup) GetMemberIdsOk() (*[]string, bool)`

GetMemberIdsOk returns a tuple with the MemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIds

`func (o *CreateGroup) SetMemberIds(v []string)`

SetMemberIds sets MemberIds field to given value.

### HasMemberIds

`func (o *CreateGroup) HasMemberIds() bool`

HasMemberIds returns a boolean if a field has been set.

### GetWorkflowIds

`func (o *CreateGroup) GetWorkflowIds() []int64`

GetWorkflowIds returns the WorkflowIds field if non-nil, zero value otherwise.

### GetWorkflowIdsOk

`func (o *CreateGroup) GetWorkflowIdsOk() (*[]int64, bool)`

GetWorkflowIdsOk returns a tuple with the WorkflowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowIds

`func (o *CreateGroup) SetWorkflowIds(v []int64)`

SetWorkflowIds sets WorkflowIds field to given value.

### HasWorkflowIds

`func (o *CreateGroup) HasWorkflowIds() bool`

HasWorkflowIds returns a boolean if a field has been set.

### GetName

`func (o *CreateGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroup) SetName(v string)`

SetName sets Name field to given value.


### GetMentionName

`func (o *CreateGroup) GetMentionName() string`

GetMentionName returns the MentionName field if non-nil, zero value otherwise.

### GetMentionNameOk

`func (o *CreateGroup) GetMentionNameOk() (*string, bool)`

GetMentionNameOk returns a tuple with the MentionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionName

`func (o *CreateGroup) SetMentionName(v string)`

SetMentionName sets MentionName field to given value.


### GetColor

`func (o *CreateGroup) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateGroup) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateGroup) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateGroup) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetColorKey

`func (o *CreateGroup) GetColorKey() string`

GetColorKey returns the ColorKey field if non-nil, zero value otherwise.

### GetColorKeyOk

`func (o *CreateGroup) GetColorKeyOk() (*string, bool)`

GetColorKeyOk returns a tuple with the ColorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorKey

`func (o *CreateGroup) SetColorKey(v string)`

SetColorKey sets ColorKey field to given value.

### HasColorKey

`func (o *CreateGroup) HasColorKey() bool`

HasColorKey returns a boolean if a field has been set.

### GetDisplayIconId

`func (o *CreateGroup) GetDisplayIconId() string`

GetDisplayIconId returns the DisplayIconId field if non-nil, zero value otherwise.

### GetDisplayIconIdOk

`func (o *CreateGroup) GetDisplayIconIdOk() (*string, bool)`

GetDisplayIconIdOk returns a tuple with the DisplayIconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIconId

`func (o *CreateGroup) SetDisplayIconId(v string)`

SetDisplayIconId sets DisplayIconId field to given value.

### HasDisplayIconId

`func (o *CreateGroup) HasDisplayIconId() bool`

HasDisplayIconId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


