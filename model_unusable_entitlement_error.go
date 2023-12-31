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

// checks if the UnusableEntitlementError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnusableEntitlementError{}

// UnusableEntitlementError struct for UnusableEntitlementError
type UnusableEntitlementError struct {
	// The tag for violating an entitlement action.
	ReasonTag string `json:"reason_tag"`
	// Short tag describing the unusable entitlement action taken by the user.
	EntitlementTag string `json:"entitlement_tag"`
	// Message displayed to the user on why their action failed.
	Message string `json:"message"`
}

// NewUnusableEntitlementError instantiates a new UnusableEntitlementError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnusableEntitlementError(reasonTag string, entitlementTag string, message string) *UnusableEntitlementError {
	this := UnusableEntitlementError{}
	this.ReasonTag = reasonTag
	this.EntitlementTag = entitlementTag
	this.Message = message
	return &this
}

// NewUnusableEntitlementErrorWithDefaults instantiates a new UnusableEntitlementError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnusableEntitlementErrorWithDefaults() *UnusableEntitlementError {
	this := UnusableEntitlementError{}
	return &this
}

// GetReasonTag returns the ReasonTag field value
func (o *UnusableEntitlementError) GetReasonTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReasonTag
}

// GetReasonTagOk returns a tuple with the ReasonTag field value
// and a boolean to check if the value has been set.
func (o *UnusableEntitlementError) GetReasonTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonTag, true
}

// SetReasonTag sets field value
func (o *UnusableEntitlementError) SetReasonTag(v string) {
	o.ReasonTag = v
}

// GetEntitlementTag returns the EntitlementTag field value
func (o *UnusableEntitlementError) GetEntitlementTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementTag
}

// GetEntitlementTagOk returns a tuple with the EntitlementTag field value
// and a boolean to check if the value has been set.
func (o *UnusableEntitlementError) GetEntitlementTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementTag, true
}

// SetEntitlementTag sets field value
func (o *UnusableEntitlementError) SetEntitlementTag(v string) {
	o.EntitlementTag = v
}

// GetMessage returns the Message field value
func (o *UnusableEntitlementError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UnusableEntitlementError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UnusableEntitlementError) SetMessage(v string) {
	o.Message = v
}

func (o UnusableEntitlementError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnusableEntitlementError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason_tag"] = o.ReasonTag
	toSerialize["entitlement_tag"] = o.EntitlementTag
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableUnusableEntitlementError struct {
	value *UnusableEntitlementError
	isSet bool
}

func (v NullableUnusableEntitlementError) Get() *UnusableEntitlementError {
	return v.value
}

func (v *NullableUnusableEntitlementError) Set(val *UnusableEntitlementError) {
	v.value = val
	v.isSet = true
}

func (v NullableUnusableEntitlementError) IsSet() bool {
	return v.isSet
}

func (v *NullableUnusableEntitlementError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnusableEntitlementError(val *UnusableEntitlementError) *NullableUnusableEntitlementError {
	return &NullableUnusableEntitlementError{value: val, isSet: true}
}

func (v NullableUnusableEntitlementError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnusableEntitlementError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


