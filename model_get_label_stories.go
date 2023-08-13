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

// checks if the GetLabelStories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLabelStories{}

// GetLabelStories struct for GetLabelStories
type GetLabelStories struct {
	// A true/false boolean indicating whether to return Stories with their descriptions.
	IncludesDescription *bool `json:"includes_description,omitempty"`
}

// NewGetLabelStories instantiates a new GetLabelStories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLabelStories() *GetLabelStories {
	this := GetLabelStories{}
	return &this
}

// NewGetLabelStoriesWithDefaults instantiates a new GetLabelStories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLabelStoriesWithDefaults() *GetLabelStories {
	this := GetLabelStories{}
	return &this
}

// GetIncludesDescription returns the IncludesDescription field value if set, zero value otherwise.
func (o *GetLabelStories) GetIncludesDescription() bool {
	if o == nil || IsNil(o.IncludesDescription) {
		var ret bool
		return ret
	}
	return *o.IncludesDescription
}

// GetIncludesDescriptionOk returns a tuple with the IncludesDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLabelStories) GetIncludesDescriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludesDescription) {
		return nil, false
	}
	return o.IncludesDescription, true
}

// HasIncludesDescription returns a boolean if a field has been set.
func (o *GetLabelStories) HasIncludesDescription() bool {
	if o != nil && !IsNil(o.IncludesDescription) {
		return true
	}

	return false
}

// SetIncludesDescription gets a reference to the given bool and assigns it to the IncludesDescription field.
func (o *GetLabelStories) SetIncludesDescription(v bool) {
	o.IncludesDescription = &v
}

func (o GetLabelStories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLabelStories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludesDescription) {
		toSerialize["includes_description"] = o.IncludesDescription
	}
	return toSerialize, nil
}

type NullableGetLabelStories struct {
	value *GetLabelStories
	isSet bool
}

func (v NullableGetLabelStories) Get() *GetLabelStories {
	return v.value
}

func (v *NullableGetLabelStories) Set(val *GetLabelStories) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLabelStories) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLabelStories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLabelStories(val *GetLabelStories) *NullableGetLabelStories {
	return &NullableGetLabelStories{value: val, isSet: true}
}

func (v NullableGetLabelStories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLabelStories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


