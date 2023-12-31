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

// checks if the CreateCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCategory{}

// CreateCategory struct for CreateCategory
type CreateCategory struct {
	// The name of the new Category.
	Name string `json:"name"`
	// The hex color to be displayed with the Category (for example, \"#ff0000\").
	Color *string `json:"color,omitempty"`
	// This field can be set to another unique ID. In the case that the Category has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id,omitempty"`
	// The type of entity this Category is associated with; currently Milestone is the only type of Category.
	Type string `json:"type"`
}

// NewCreateCategory instantiates a new CreateCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCategory(name string, type_ string) *CreateCategory {
	this := CreateCategory{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewCreateCategoryWithDefaults instantiates a new CreateCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCategoryWithDefaults() *CreateCategory {
	this := CreateCategory{}
	return &this
}

// GetName returns the Name field value
func (o *CreateCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCategory) SetName(v string) {
	o.Name = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *CreateCategory) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCategory) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *CreateCategory) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *CreateCategory) SetColor(v string) {
	o.Color = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *CreateCategory) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCategory) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *CreateCategory) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *CreateCategory) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetType returns the Type field value
func (o *CreateCategory) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateCategory) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateCategory) SetType(v string) {
	o.Type = v
}

func (o CreateCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCreateCategory struct {
	value *CreateCategory
	isSet bool
}

func (v NullableCreateCategory) Get() *CreateCategory {
	return v.value
}

func (v *NullableCreateCategory) Set(val *CreateCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCategory(val *CreateCategory) *NullableCreateCategory {
	return &NullableCreateCategory{value: val, isSet: true}
}

func (v NullableCreateCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


