# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | The Member&#39;s role in the Workspace. | 
**EntityType** | **string** | A string description of this resource. | 
**Disabled** | **bool** | True/false boolean indicating whether the Member has been disabled within the Workspace. | 
**GlobalId** | **string** |  | 
**State** | **string** | The user state, one of partial, full, disabled, or imported.  A partial user is disabled, has no means to log in, and is not an import user.  A full user is enabled and has a means to log in.  A disabled user is disabled and has a means to log in.  An import user is disabled, has no means to log in, and is marked as an import user. | 
**UpdatedAt** | **NullableTime** | The time/date the Member was last updated. | 
**CreatedWithoutInvite** | **bool** | Whether this member was created as a placeholder entity. | 
**GroupIds** | **[]string** | The Member&#39;s group ids | 
**Id** | **string** | The Member&#39;s ID in Shortcut. | 
**Profile** | [**Profile**](Profile.md) |  | 
**CreatedAt** | **NullableTime** | The time/date the Member was created. | 
**ReplacedBy** | Pointer to **string** | The id of the member that replaces this one when merged. | [optional] 

## Methods

### NewMember

`func NewMember(role string, entityType string, disabled bool, globalId string, state string, updatedAt NullableTime, createdWithoutInvite bool, groupIds []string, id string, profile Profile, createdAt NullableTime, ) *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *Member) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Member) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Member) SetRole(v string)`

SetRole sets Role field to given value.


### GetEntityType

`func (o *Member) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Member) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Member) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetDisabled

`func (o *Member) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Member) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Member) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetGlobalId

`func (o *Member) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *Member) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *Member) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetState

`func (o *Member) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Member) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Member) SetState(v string)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *Member) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Member) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Member) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Member) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Member) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetCreatedWithoutInvite

`func (o *Member) GetCreatedWithoutInvite() bool`

GetCreatedWithoutInvite returns the CreatedWithoutInvite field if non-nil, zero value otherwise.

### GetCreatedWithoutInviteOk

`func (o *Member) GetCreatedWithoutInviteOk() (*bool, bool)`

GetCreatedWithoutInviteOk returns a tuple with the CreatedWithoutInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedWithoutInvite

`func (o *Member) SetCreatedWithoutInvite(v bool)`

SetCreatedWithoutInvite sets CreatedWithoutInvite field to given value.


### GetGroupIds

`func (o *Member) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *Member) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *Member) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.


### GetId

`func (o *Member) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v string)`

SetId sets Id field to given value.


### GetProfile

`func (o *Member) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Member) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Member) SetProfile(v Profile)`

SetProfile sets Profile field to given value.


### GetCreatedAt

`func (o *Member) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Member) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Member) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Member) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Member) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetReplacedBy

`func (o *Member) GetReplacedBy() string`

GetReplacedBy returns the ReplacedBy field if non-nil, zero value otherwise.

### GetReplacedByOk

`func (o *Member) GetReplacedByOk() (*string, bool)`

GetReplacedByOk returns a tuple with the ReplacedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedBy

`func (o *Member) SetReplacedBy(v string)`

SetReplacedBy sets ReplacedBy field to given value.

### HasReplacedBy

`func (o *Member) HasReplacedBy() bool`

HasReplacedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


