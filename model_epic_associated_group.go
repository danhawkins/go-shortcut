/*
Shortcut API

Shortcut API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EpicAssociatedGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EpicAssociatedGroup{}

// EpicAssociatedGroup struct for EpicAssociatedGroup
type EpicAssociatedGroup struct {
	// The Group ID of the associated group.
	GroupId string `json:"group_id"`
	// The number of stories this Group owns in the Epic.
	AssociatedStoriesCount *int64 `json:"associated_stories_count,omitempty"`
}

// NewEpicAssociatedGroup instantiates a new EpicAssociatedGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpicAssociatedGroup(groupId string) *EpicAssociatedGroup {
	this := EpicAssociatedGroup{}
	this.GroupId = groupId
	return &this
}

// NewEpicAssociatedGroupWithDefaults instantiates a new EpicAssociatedGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpicAssociatedGroupWithDefaults() *EpicAssociatedGroup {
	this := EpicAssociatedGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *EpicAssociatedGroup) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *EpicAssociatedGroup) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *EpicAssociatedGroup) SetGroupId(v string) {
	o.GroupId = v
}

// GetAssociatedStoriesCount returns the AssociatedStoriesCount field value if set, zero value otherwise.
func (o *EpicAssociatedGroup) GetAssociatedStoriesCount() int64 {
	if o == nil || IsNil(o.AssociatedStoriesCount) {
		var ret int64
		return ret
	}
	return *o.AssociatedStoriesCount
}

// GetAssociatedStoriesCountOk returns a tuple with the AssociatedStoriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpicAssociatedGroup) GetAssociatedStoriesCountOk() (*int64, bool) {
	if o == nil || IsNil(o.AssociatedStoriesCount) {
		return nil, false
	}
	return o.AssociatedStoriesCount, true
}

// HasAssociatedStoriesCount returns a boolean if a field has been set.
func (o *EpicAssociatedGroup) HasAssociatedStoriesCount() bool {
	if o != nil && !IsNil(o.AssociatedStoriesCount) {
		return true
	}

	return false
}

// SetAssociatedStoriesCount gets a reference to the given int64 and assigns it to the AssociatedStoriesCount field.
func (o *EpicAssociatedGroup) SetAssociatedStoriesCount(v int64) {
	o.AssociatedStoriesCount = &v
}

func (o EpicAssociatedGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpicAssociatedGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group_id"] = o.GroupId
	if !IsNil(o.AssociatedStoriesCount) {
		toSerialize["associated_stories_count"] = o.AssociatedStoriesCount
	}
	return toSerialize, nil
}

type NullableEpicAssociatedGroup struct {
	value *EpicAssociatedGroup
	isSet bool
}

func (v NullableEpicAssociatedGroup) Get() *EpicAssociatedGroup {
	return v.value
}

func (v *NullableEpicAssociatedGroup) Set(val *EpicAssociatedGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableEpicAssociatedGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableEpicAssociatedGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpicAssociatedGroup(val *EpicAssociatedGroup) *NullableEpicAssociatedGroup {
	return &NullableEpicAssociatedGroup{value: val, isSet: true}
}

func (v NullableEpicAssociatedGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpicAssociatedGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

