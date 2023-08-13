# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | 
**Deactivated** | **bool** | A true/false boolean indicating whether the Member has been deactivated within Shortcut. | 
**TwoFactorAuthActivated** | Pointer to **bool** | If Two Factor Authentication is activated for this User. | [optional] 
**MentionName** | **string** | The Member&#39;s username within the Organization. | 
**Name** | **NullableString** | The Member&#39;s name within the Organization. | 
**GravatarHash** | **NullableString** | This is the gravatar hash associated with email_address. | 
**Id** | **string** | The unique identifier of the profile. | 
**DisplayIcon** | [**Icon**](Icon.md) |  | 
**IsOwner** | **bool** | A boolean indicating whether this profile is an owner at their associated organization. | 
**EmailAddress** | **NullableString** | The primary email address of the Member with the Organization. | 

## Methods

### NewProfile

`func NewProfile(entityType string, deactivated bool, mentionName string, name NullableString, gravatarHash NullableString, id string, displayIcon Icon, isOwner bool, emailAddress NullableString, ) *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *Profile) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Profile) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Profile) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetDeactivated

`func (o *Profile) GetDeactivated() bool`

GetDeactivated returns the Deactivated field if non-nil, zero value otherwise.

### GetDeactivatedOk

`func (o *Profile) GetDeactivatedOk() (*bool, bool)`

GetDeactivatedOk returns a tuple with the Deactivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivated

`func (o *Profile) SetDeactivated(v bool)`

SetDeactivated sets Deactivated field to given value.


### GetTwoFactorAuthActivated

`func (o *Profile) GetTwoFactorAuthActivated() bool`

GetTwoFactorAuthActivated returns the TwoFactorAuthActivated field if non-nil, zero value otherwise.

### GetTwoFactorAuthActivatedOk

`func (o *Profile) GetTwoFactorAuthActivatedOk() (*bool, bool)`

GetTwoFactorAuthActivatedOk returns a tuple with the TwoFactorAuthActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthActivated

`func (o *Profile) SetTwoFactorAuthActivated(v bool)`

SetTwoFactorAuthActivated sets TwoFactorAuthActivated field to given value.

### HasTwoFactorAuthActivated

`func (o *Profile) HasTwoFactorAuthActivated() bool`

HasTwoFactorAuthActivated returns a boolean if a field has been set.

### GetMentionName

`func (o *Profile) GetMentionName() string`

GetMentionName returns the MentionName field if non-nil, zero value otherwise.

### GetMentionNameOk

`func (o *Profile) GetMentionNameOk() (*string, bool)`

GetMentionNameOk returns a tuple with the MentionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionName

`func (o *Profile) SetMentionName(v string)`

SetMentionName sets MentionName field to given value.


### GetName

`func (o *Profile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Profile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Profile) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Profile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Profile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetGravatarHash

`func (o *Profile) GetGravatarHash() string`

GetGravatarHash returns the GravatarHash field if non-nil, zero value otherwise.

### GetGravatarHashOk

`func (o *Profile) GetGravatarHashOk() (*string, bool)`

GetGravatarHashOk returns a tuple with the GravatarHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarHash

`func (o *Profile) SetGravatarHash(v string)`

SetGravatarHash sets GravatarHash field to given value.


### SetGravatarHashNil

`func (o *Profile) SetGravatarHashNil(b bool)`

 SetGravatarHashNil sets the value for GravatarHash to be an explicit nil

### UnsetGravatarHash
`func (o *Profile) UnsetGravatarHash()`

UnsetGravatarHash ensures that no value is present for GravatarHash, not even an explicit nil
### GetId

`func (o *Profile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Profile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Profile) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayIcon

`func (o *Profile) GetDisplayIcon() Icon`

GetDisplayIcon returns the DisplayIcon field if non-nil, zero value otherwise.

### GetDisplayIconOk

`func (o *Profile) GetDisplayIconOk() (*Icon, bool)`

GetDisplayIconOk returns a tuple with the DisplayIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIcon

`func (o *Profile) SetDisplayIcon(v Icon)`

SetDisplayIcon sets DisplayIcon field to given value.


### GetIsOwner

`func (o *Profile) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *Profile) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *Profile) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.


### GetEmailAddress

`func (o *Profile) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Profile) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Profile) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### SetEmailAddressNil

`func (o *Profile) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *Profile) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


