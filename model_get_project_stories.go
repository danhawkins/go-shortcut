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

// checks if the GetProjectStories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectStories{}

// GetProjectStories struct for GetProjectStories
type GetProjectStories struct {
	// A true/false boolean indicating whether to return Stories with their descriptions.
	IncludesDescription *bool `json:"includes_description,omitempty"`
}

// NewGetProjectStories instantiates a new GetProjectStories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectStories() *GetProjectStories {
	this := GetProjectStories{}
	return &this
}

// NewGetProjectStoriesWithDefaults instantiates a new GetProjectStories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectStoriesWithDefaults() *GetProjectStories {
	this := GetProjectStories{}
	return &this
}

// GetIncludesDescription returns the IncludesDescription field value if set, zero value otherwise.
func (o *GetProjectStories) GetIncludesDescription() bool {
	if o == nil || IsNil(o.IncludesDescription) {
		var ret bool
		return ret
	}
	return *o.IncludesDescription
}

// GetIncludesDescriptionOk returns a tuple with the IncludesDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectStories) GetIncludesDescriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludesDescription) {
		return nil, false
	}
	return o.IncludesDescription, true
}

// HasIncludesDescription returns a boolean if a field has been set.
func (o *GetProjectStories) HasIncludesDescription() bool {
	if o != nil && !IsNil(o.IncludesDescription) {
		return true
	}

	return false
}

// SetIncludesDescription gets a reference to the given bool and assigns it to the IncludesDescription field.
func (o *GetProjectStories) SetIncludesDescription(v bool) {
	o.IncludesDescription = &v
}

func (o GetProjectStories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectStories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludesDescription) {
		toSerialize["includes_description"] = o.IncludesDescription
	}
	return toSerialize, nil
}

type NullableGetProjectStories struct {
	value *GetProjectStories
	isSet bool
}

func (v NullableGetProjectStories) Get() *GetProjectStories {
	return v.value
}

func (v *NullableGetProjectStories) Set(val *GetProjectStories) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectStories) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectStories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectStories(val *GetProjectStories) *NullableGetProjectStories {
	return &NullableGetProjectStories{value: val, isSet: true}
}

func (v NullableGetProjectStories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectStories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


