/*
Shortcut API

Shortcut API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Member type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Member{}

// Member Details about an individual user within the Workspace.
type Member struct {
	// The Member's role in the Workspace.
	Role string `json:"role"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// True/false boolean indicating whether the Member has been disabled within the Workspace.
	Disabled bool `json:"disabled"`
	GlobalId string `json:"global_id"`
	// The user state, one of partial, full, disabled, or imported.  A partial user is disabled, has no means to log in, and is not an import user.  A full user is enabled and has a means to log in.  A disabled user is disabled and has a means to log in.  An import user is disabled, has no means to log in, and is marked as an import user.
	State string `json:"state"`
	// The time/date the Member was last updated.
	UpdatedAt NullableTime `json:"updated_at"`
	// Whether this member was created as a placeholder entity.
	CreatedWithoutInvite bool `json:"created_without_invite"`
	// The Member's group ids
	GroupIds []string `json:"group_ids"`
	// The Member's ID in Shortcut.
	Id string `json:"id"`
	Profile Profile `json:"profile"`
	// The time/date the Member was created.
	CreatedAt NullableTime `json:"created_at"`
	// The id of the member that replaces this one when merged.
	ReplacedBy *string `json:"replaced_by,omitempty"`
}

// NewMember instantiates a new Member object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMember(role string, entityType string, disabled bool, globalId string, state string, updatedAt NullableTime, createdWithoutInvite bool, groupIds []string, id string, profile Profile, createdAt NullableTime) *Member {
	this := Member{}
	this.Role = role
	this.EntityType = entityType
	this.Disabled = disabled
	this.GlobalId = globalId
	this.State = state
	this.UpdatedAt = updatedAt
	this.CreatedWithoutInvite = createdWithoutInvite
	this.GroupIds = groupIds
	this.Id = id
	this.Profile = profile
	this.CreatedAt = createdAt
	return &this
}

// NewMemberWithDefaults instantiates a new Member object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberWithDefaults() *Member {
	this := Member{}
	return &this
}

// GetRole returns the Role field value
func (o *Member) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Member) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *Member) SetRole(v string) {
	o.Role = v
}

// GetEntityType returns the EntityType field value
func (o *Member) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *Member) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *Member) SetEntityType(v string) {
	o.EntityType = v
}

// GetDisabled returns the Disabled field value
func (o *Member) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *Member) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *Member) SetDisabled(v bool) {
	o.Disabled = v
}

// GetGlobalId returns the GlobalId field value
func (o *Member) GetGlobalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *Member) GetGlobalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *Member) SetGlobalId(v string) {
	o.GlobalId = v
}

// GetState returns the State field value
func (o *Member) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Member) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Member) SetState(v string) {
	o.State = v
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Member) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Member) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *Member) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// GetCreatedWithoutInvite returns the CreatedWithoutInvite field value
func (o *Member) GetCreatedWithoutInvite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CreatedWithoutInvite
}

// GetCreatedWithoutInviteOk returns a tuple with the CreatedWithoutInvite field value
// and a boolean to check if the value has been set.
func (o *Member) GetCreatedWithoutInviteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedWithoutInvite, true
}

// SetCreatedWithoutInvite sets field value
func (o *Member) SetCreatedWithoutInvite(v bool) {
	o.CreatedWithoutInvite = v
}

// GetGroupIds returns the GroupIds field value
func (o *Member) GetGroupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value
// and a boolean to check if the value has been set.
func (o *Member) GetGroupIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// SetGroupIds sets field value
func (o *Member) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetId returns the Id field value
func (o *Member) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Member) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Member) SetId(v string) {
	o.Id = v
}

// GetProfile returns the Profile field value
func (o *Member) GetProfile() Profile {
	if o == nil {
		var ret Profile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *Member) GetProfileOk() (*Profile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *Member) SetProfile(v Profile) {
	o.Profile = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Member) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Member) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *Member) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// GetReplacedBy returns the ReplacedBy field value if set, zero value otherwise.
func (o *Member) GetReplacedBy() string {
	if o == nil || IsNil(o.ReplacedBy) {
		var ret string
		return ret
	}
	return *o.ReplacedBy
}

// GetReplacedByOk returns a tuple with the ReplacedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Member) GetReplacedByOk() (*string, bool) {
	if o == nil || IsNil(o.ReplacedBy) {
		return nil, false
	}
	return o.ReplacedBy, true
}

// HasReplacedBy returns a boolean if a field has been set.
func (o *Member) HasReplacedBy() bool {
	if o != nil && !IsNil(o.ReplacedBy) {
		return true
	}

	return false
}

// SetReplacedBy gets a reference to the given string and assigns it to the ReplacedBy field.
func (o *Member) SetReplacedBy(v string) {
	o.ReplacedBy = &v
}

func (o Member) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Member) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["entity_type"] = o.EntityType
	toSerialize["disabled"] = o.Disabled
	toSerialize["global_id"] = o.GlobalId
	toSerialize["state"] = o.State
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	toSerialize["created_without_invite"] = o.CreatedWithoutInvite
	toSerialize["group_ids"] = o.GroupIds
	toSerialize["id"] = o.Id
	toSerialize["profile"] = o.Profile
	toSerialize["created_at"] = o.CreatedAt.Get()
	if !IsNil(o.ReplacedBy) {
		toSerialize["replaced_by"] = o.ReplacedBy
	}
	return toSerialize, nil
}

type NullableMember struct {
	value *Member
	isSet bool
}

func (v NullableMember) Get() *Member {
	return v.value
}

func (v *NullableMember) Set(val *Member) {
	v.value = val
	v.isSet = true
}

func (v NullableMember) IsSet() bool {
	return v.isSet
}

func (v *NullableMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMember(val *Member) *NullableMember {
	return &NullableMember{value: val, isSet: true}
}

func (v NullableMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


