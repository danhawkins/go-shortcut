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

// checks if the UpdateCustomField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCustomField{}

// UpdateCustomField struct for UpdateCustomField
type UpdateCustomField struct {
	// Indicates whether the Field is enabled for the Workspace. Only enabled fields can be applied to Stories.
	Enabled *bool `json:"enabled,omitempty"`
	// A collection of objects representing reporting periods for years.
	Name *string `json:"name,omitempty"`
	// A collection of EnumValue objects representing the values in the domain of some Custom Field.
	Values []UpdateCustomFieldEnumValue `json:"values,omitempty"`
	// A frontend-controlled string that represents the icon for this custom field.
	IconSetIdentifier *string `json:"icon_set_identifier,omitempty"`
	// A description of the purpose of this field.
	Description *string `json:"description,omitempty"`
	// The ID of the CustomField we want to move this CustomField before.
	BeforeId *string `json:"before_id,omitempty"`
	// The ID of the CustomField we want to move this CustomField after.
	AfterId *string `json:"after_id,omitempty"`
}

// NewUpdateCustomField instantiates a new UpdateCustomField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomField() *UpdateCustomField {
	this := UpdateCustomField{}
	return &this
}

// NewUpdateCustomFieldWithDefaults instantiates a new UpdateCustomField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomFieldWithDefaults() *UpdateCustomField {
	this := UpdateCustomField{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateCustomField) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomField) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateCustomField) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateCustomField) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateCustomField) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomField) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateCustomField) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateCustomField) SetName(v string) {
	o.Name = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *UpdateCustomField) GetValues() []UpdateCustomFieldEnumValue {
	if o == nil || IsNil(o.Values) {
		var ret []UpdateCustomFieldEnumValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomField) GetValuesOk() ([]UpdateCustomFieldEnumValue, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *UpdateCustomField) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []UpdateCustomFieldEnumValue and assigns it to the Values field.
func (o *UpdateCustomField) SetValues(v []UpdateCustomFieldEnumValue) {
	o.Values = v
}

// GetIconSetIdentifier returns the IconSetIdentifier field value if set, zero value otherwise.
func (o *UpdateCustomField) GetIconSetIdentifier() string {
	if o == nil || IsNil(o.IconSetIdentifier) {
		var ret string
		return ret
	}
	return *o.IconSetIdentifier
}

// GetIconSetIdentifierOk returns a tuple with the IconSetIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomField) GetIconSetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.IconSetIdentifier) {
		return nil, false
	}
	return o.IconSetIdentifier, true
}

// HasIconSetIdentifier returns a boolean if a field has been set.
func (o *UpdateCustomField) HasIconSetIdentifier() bool {
	if o != nil && !IsNil(o.IconSetIdentifier) {
		return true
	}

	return false
}

// SetIconSetIdentifier gets a reference to the given string and assigns it to the IconSetIdentifier field.
func (o *UpdateCustomField) SetIconSetIdentifier(v string) {
	o.IconSetIdentifier = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCustomField) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomField) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateCustomField) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCustomField) SetDescription(v string) {
	o.Description = &v
}

// GetBeforeId returns the BeforeId field value if set, zero value otherwise.
func (o *UpdateCustomField) GetBeforeId() string {
	if o == nil || IsNil(o.BeforeId) {
		var ret string
		return ret
	}
	return *o.BeforeId
}

// GetBeforeIdOk returns a tuple with the BeforeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomField) GetBeforeIdOk() (*string, bool) {
	if o == nil || IsNil(o.BeforeId) {
		return nil, false
	}
	return o.BeforeId, true
}

// HasBeforeId returns a boolean if a field has been set.
func (o *UpdateCustomField) HasBeforeId() bool {
	if o != nil && !IsNil(o.BeforeId) {
		return true
	}

	return false
}

// SetBeforeId gets a reference to the given string and assigns it to the BeforeId field.
func (o *UpdateCustomField) SetBeforeId(v string) {
	o.BeforeId = &v
}

// GetAfterId returns the AfterId field value if set, zero value otherwise.
func (o *UpdateCustomField) GetAfterId() string {
	if o == nil || IsNil(o.AfterId) {
		var ret string
		return ret
	}
	return *o.AfterId
}

// GetAfterIdOk returns a tuple with the AfterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomField) GetAfterIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfterId) {
		return nil, false
	}
	return o.AfterId, true
}

// HasAfterId returns a boolean if a field has been set.
func (o *UpdateCustomField) HasAfterId() bool {
	if o != nil && !IsNil(o.AfterId) {
		return true
	}

	return false
}

// SetAfterId gets a reference to the given string and assigns it to the AfterId field.
func (o *UpdateCustomField) SetAfterId(v string) {
	o.AfterId = &v
}

func (o UpdateCustomField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.IconSetIdentifier) {
		toSerialize["icon_set_identifier"] = o.IconSetIdentifier
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.BeforeId) {
		toSerialize["before_id"] = o.BeforeId
	}
	if !IsNil(o.AfterId) {
		toSerialize["after_id"] = o.AfterId
	}
	return toSerialize, nil
}

type NullableUpdateCustomField struct {
	value *UpdateCustomField
	isSet bool
}

func (v NullableUpdateCustomField) Get() *UpdateCustomField {
	return v.value
}

func (v *NullableUpdateCustomField) Set(val *UpdateCustomField) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomField) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomField(val *UpdateCustomField) *NullableUpdateCustomField {
	return &NullableUpdateCustomField{value: val, isSet: true}
}

func (v NullableUpdateCustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


