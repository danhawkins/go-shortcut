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

// checks if the UpdateCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCategory{}

// UpdateCategory struct for UpdateCategory
type UpdateCategory struct {
	// The new name of the Category.
	Name *string `json:"name,omitempty"`
	// The hex color to be displayed with the Category (for example, \"#ff0000\").
	Color NullableString `json:"color,omitempty"`
	// A true/false boolean indicating if the Category has been archived.
	Archived *bool `json:"archived,omitempty"`
}

// NewUpdateCategory instantiates a new UpdateCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCategory() *UpdateCategory {
	this := UpdateCategory{}
	return &this
}

// NewUpdateCategoryWithDefaults instantiates a new UpdateCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCategoryWithDefaults() *UpdateCategory {
	this := UpdateCategory{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateCategory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCategory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateCategory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateCategory) SetName(v string) {
	o.Name = &v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCategory) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCategory) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *UpdateCategory) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *UpdateCategory) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *UpdateCategory) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *UpdateCategory) UnsetColor() {
	o.Color.Unset()
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *UpdateCategory) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCategory) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *UpdateCategory) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *UpdateCategory) SetArchived(v bool) {
	o.Archived = &v
}

func (o UpdateCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	return toSerialize, nil
}

type NullableUpdateCategory struct {
	value *UpdateCategory
	isSet bool
}

func (v NullableUpdateCategory) Get() *UpdateCategory {
	return v.value
}

func (v *NullableUpdateCategory) Set(val *UpdateCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCategory(val *UpdateCategory) *NullableUpdateCategory {
	return &NullableUpdateCategory{value: val, isSet: true}
}

func (v NullableUpdateCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


