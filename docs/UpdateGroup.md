# UpdateGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of this Group. | [optional] 
**Archived** | Pointer to **NullableBool** | Whether or not this Group is archived. | [optional] 
**Color** | Pointer to **NullableString** | The color you wish to use for the Group in the system. | [optional] 
**DisplayIconId** | Pointer to **NullableString** | The Icon id for the avatar of this Group. | [optional] 
**MentionName** | Pointer to **string** | The mention name of this Group. | [optional] 
**Name** | Pointer to **string** | The name of this Group. | [optional] 
**ColorKey** | Pointer to **string** | The color key you wish to use for the Group in the system. | [optional] 
**MemberIds** | Pointer to **[]string** | The Member ids to add to this Group. | [optional] 
**WorkflowIds** | Pointer to **[]int64** | The Workflow ids to add to the Group. | [optional] 

## Methods

### NewUpdateGroup

`func NewUpdateGroup() *UpdateGroup`

NewUpdateGroup instantiates a new UpdateGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupWithDefaults

`func NewUpdateGroupWithDefaults() *UpdateGroup`

NewUpdateGroupWithDefaults instantiates a new UpdateGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *UpdateGroup) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateGroup) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateGroup) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateGroup) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### SetArchivedNil

`func (o *UpdateGroup) SetArchivedNil(b bool)`

 SetArchivedNil sets the value for Archived to be an explicit nil

### UnsetArchived
`func (o *UpdateGroup) UnsetArchived()`

UnsetArchived ensures that no value is present for Archived, not even an explicit nil
### GetColor

`func (o *UpdateGroup) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateGroup) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateGroup) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateGroup) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *UpdateGroup) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *UpdateGroup) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDisplayIconId

`func (o *UpdateGroup) GetDisplayIconId() string`

GetDisplayIconId returns the DisplayIconId field if non-nil, zero value otherwise.

### GetDisplayIconIdOk

`func (o *UpdateGroup) GetDisplayIconIdOk() (*string, bool)`

GetDisplayIconIdOk returns a tuple with the DisplayIconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIconId

`func (o *UpdateGroup) SetDisplayIconId(v string)`

SetDisplayIconId sets DisplayIconId field to given value.

### HasDisplayIconId

`func (o *UpdateGroup) HasDisplayIconId() bool`

HasDisplayIconId returns a boolean if a field has been set.

### SetDisplayIconIdNil

`func (o *UpdateGroup) SetDisplayIconIdNil(b bool)`

 SetDisplayIconIdNil sets the value for DisplayIconId to be an explicit nil

### UnsetDisplayIconId
`func (o *UpdateGroup) UnsetDisplayIconId()`

UnsetDisplayIconId ensures that no value is present for DisplayIconId, not even an explicit nil
### GetMentionName

`func (o *UpdateGroup) GetMentionName() string`

GetMentionName returns the MentionName field if non-nil, zero value otherwise.

### GetMentionNameOk

`func (o *UpdateGroup) GetMentionNameOk() (*string, bool)`

GetMentionNameOk returns a tuple with the MentionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionName

`func (o *UpdateGroup) SetMentionName(v string)`

SetMentionName sets MentionName field to given value.

### HasMentionName

`func (o *UpdateGroup) HasMentionName() bool`

HasMentionName returns a boolean if a field has been set.

### GetName

`func (o *UpdateGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColorKey

`func (o *UpdateGroup) GetColorKey() string`

GetColorKey returns the ColorKey field if non-nil, zero value otherwise.

### GetColorKeyOk

`func (o *UpdateGroup) GetColorKeyOk() (*string, bool)`

GetColorKeyOk returns a tuple with the ColorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorKey

`func (o *UpdateGroup) SetColorKey(v string)`

SetColorKey sets ColorKey field to given value.

### HasColorKey

`func (o *UpdateGroup) HasColorKey() bool`

HasColorKey returns a boolean if a field has been set.

### GetMemberIds

`func (o *UpdateGroup) GetMemberIds() []string`

GetMemberIds returns the MemberIds field if non-nil, zero value otherwise.

### GetMemberIdsOk

`func (o *UpdateGroup) GetMemberIdsOk() (*[]string, bool)`

GetMemberIdsOk returns a tuple with the MemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIds

`func (o *UpdateGroup) SetMemberIds(v []string)`

SetMemberIds sets MemberIds field to given value.

### HasMemberIds

`func (o *UpdateGroup) HasMemberIds() bool`

HasMemberIds returns a boolean if a field has been set.

### GetWorkflowIds

`func (o *UpdateGroup) GetWorkflowIds() []int64`

GetWorkflowIds returns the WorkflowIds field if non-nil, zero value otherwise.

### GetWorkflowIdsOk

`func (o *UpdateGroup) GetWorkflowIdsOk() (*[]int64, bool)`

GetWorkflowIdsOk returns a tuple with the WorkflowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowIds

`func (o *UpdateGroup) SetWorkflowIds(v []int64)`

SetWorkflowIds sets WorkflowIds field to given value.

### HasWorkflowIds

`func (o *UpdateGroup) HasWorkflowIds() bool`

HasWorkflowIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


