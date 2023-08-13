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

// checks if the DeleteStories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteStories{}

// DeleteStories struct for DeleteStories
type DeleteStories struct {
	// An array of IDs of Stories to delete.
	StoryIds []int64 `json:"story_ids"`
}

// NewDeleteStories instantiates a new DeleteStories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteStories(storyIds []int64) *DeleteStories {
	this := DeleteStories{}
	this.StoryIds = storyIds
	return &this
}

// NewDeleteStoriesWithDefaults instantiates a new DeleteStories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteStoriesWithDefaults() *DeleteStories {
	this := DeleteStories{}
	return &this
}

// GetStoryIds returns the StoryIds field value
func (o *DeleteStories) GetStoryIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.StoryIds
}

// GetStoryIdsOk returns a tuple with the StoryIds field value
// and a boolean to check if the value has been set.
func (o *DeleteStories) GetStoryIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoryIds, true
}

// SetStoryIds sets field value
func (o *DeleteStories) SetStoryIds(v []int64) {
	o.StoryIds = v
}

func (o DeleteStories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteStories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["story_ids"] = o.StoryIds
	return toSerialize, nil
}

type NullableDeleteStories struct {
	value *DeleteStories
	isSet bool
}

func (v NullableDeleteStories) Get() *DeleteStories {
	return v.value
}

func (v *NullableDeleteStories) Set(val *DeleteStories) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteStories) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteStories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteStories(val *DeleteStories) *NullableDeleteStories {
	return &NullableDeleteStories{value: val, isSet: true}
}

func (v NullableDeleteStories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteStories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


