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

// checks if the PullRequestLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PullRequestLabel{}

// PullRequestLabel Corresponds to a VCS Label associated with a Pull Request.
type PullRequestLabel struct {
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique ID of the VCS Label.
	Id int64 `json:"id"`
	// The color of the VCS label.
	Color string `json:"color"`
	// The description of the VCS label.
	Description NullableString `json:"description,omitempty"`
	// The name of the VCS label.
	Name string `json:"name"`
}

// NewPullRequestLabel instantiates a new PullRequestLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequestLabel(entityType string, id int64, color string, name string) *PullRequestLabel {
	this := PullRequestLabel{}
	this.EntityType = entityType
	this.Id = id
	this.Color = color
	this.Name = name
	return &this
}

// NewPullRequestLabelWithDefaults instantiates a new PullRequestLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestLabelWithDefaults() *PullRequestLabel {
	this := PullRequestLabel{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *PullRequestLabel) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *PullRequestLabel) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *PullRequestLabel) SetEntityType(v string) {
	o.EntityType = v
}

// GetId returns the Id field value
func (o *PullRequestLabel) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PullRequestLabel) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PullRequestLabel) SetId(v int64) {
	o.Id = v
}

// GetColor returns the Color field value
func (o *PullRequestLabel) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *PullRequestLabel) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *PullRequestLabel) SetColor(v string) {
	o.Color = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PullRequestLabel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestLabel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PullRequestLabel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PullRequestLabel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PullRequestLabel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PullRequestLabel) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value
func (o *PullRequestLabel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PullRequestLabel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PullRequestLabel) SetName(v string) {
	o.Name = v
}

func (o PullRequestLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PullRequestLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_type"] = o.EntityType
	toSerialize["id"] = o.Id
	toSerialize["color"] = o.Color
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullablePullRequestLabel struct {
	value *PullRequestLabel
	isSet bool
}

func (v NullablePullRequestLabel) Get() *PullRequestLabel {
	return v.value
}

func (v *NullablePullRequestLabel) Set(val *PullRequestLabel) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequestLabel) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequestLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequestLabel(val *PullRequestLabel) *NullablePullRequestLabel {
	return &NullablePullRequestLabel{value: val, isSet: true}
}

func (v NullablePullRequestLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequestLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


