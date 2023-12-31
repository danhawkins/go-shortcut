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

// checks if the UpdateEntityTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntityTemplate{}

// UpdateEntityTemplate Request parameters for changing either a template's name or any of   the attributes it is designed to pre-populate.
type UpdateEntityTemplate struct {
	// The updated template name.
	Name *string `json:"name,omitempty"`
	StoryContents *UpdateStoryContents `json:"story_contents,omitempty"`
}

// NewUpdateEntityTemplate instantiates a new UpdateEntityTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntityTemplate() *UpdateEntityTemplate {
	this := UpdateEntityTemplate{}
	return &this
}

// NewUpdateEntityTemplateWithDefaults instantiates a new UpdateEntityTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntityTemplateWithDefaults() *UpdateEntityTemplate {
	this := UpdateEntityTemplate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateEntityTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateEntityTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateEntityTemplate) SetName(v string) {
	o.Name = &v
}

// GetStoryContents returns the StoryContents field value if set, zero value otherwise.
func (o *UpdateEntityTemplate) GetStoryContents() UpdateStoryContents {
	if o == nil || IsNil(o.StoryContents) {
		var ret UpdateStoryContents
		return ret
	}
	return *o.StoryContents
}

// GetStoryContentsOk returns a tuple with the StoryContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEntityTemplate) GetStoryContentsOk() (*UpdateStoryContents, bool) {
	if o == nil || IsNil(o.StoryContents) {
		return nil, false
	}
	return o.StoryContents, true
}

// HasStoryContents returns a boolean if a field has been set.
func (o *UpdateEntityTemplate) HasStoryContents() bool {
	if o != nil && !IsNil(o.StoryContents) {
		return true
	}

	return false
}

// SetStoryContents gets a reference to the given UpdateStoryContents and assigns it to the StoryContents field.
func (o *UpdateEntityTemplate) SetStoryContents(v UpdateStoryContents) {
	o.StoryContents = &v
}

func (o UpdateEntityTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntityTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StoryContents) {
		toSerialize["story_contents"] = o.StoryContents
	}
	return toSerialize, nil
}

type NullableUpdateEntityTemplate struct {
	value *UpdateEntityTemplate
	isSet bool
}

func (v NullableUpdateEntityTemplate) Get() *UpdateEntityTemplate {
	return v.value
}

func (v *NullableUpdateEntityTemplate) Set(val *UpdateEntityTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntityTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntityTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntityTemplate(val *UpdateEntityTemplate) *NullableUpdateEntityTemplate {
	return &NullableUpdateEntityTemplate{value: val, isSet: true}
}

func (v NullableUpdateEntityTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntityTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


