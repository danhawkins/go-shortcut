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

// checks if the HistoryActionLabelUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionLabelUpdate{}

// HistoryActionLabelUpdate An action representing a Label being updated.
type HistoryActionLabelUpdate struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
}

// NewHistoryActionLabelUpdate instantiates a new HistoryActionLabelUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionLabelUpdate(id int64, entityType string, action string) *HistoryActionLabelUpdate {
	this := HistoryActionLabelUpdate{}
	this.Id = id
	this.EntityType = entityType
	this.Action = action
	return &this
}

// NewHistoryActionLabelUpdateWithDefaults instantiates a new HistoryActionLabelUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionLabelUpdateWithDefaults() *HistoryActionLabelUpdate {
	this := HistoryActionLabelUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionLabelUpdate) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelUpdate) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionLabelUpdate) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionLabelUpdate) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelUpdate) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionLabelUpdate) SetEntityType(v string) {
	o.EntityType = v
}

// GetAction returns the Action field value
func (o *HistoryActionLabelUpdate) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelUpdate) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionLabelUpdate) SetAction(v string) {
	o.Action = v
}

func (o HistoryActionLabelUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionLabelUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableHistoryActionLabelUpdate struct {
	value *HistoryActionLabelUpdate
	isSet bool
}

func (v NullableHistoryActionLabelUpdate) Get() *HistoryActionLabelUpdate {
	return v.value
}

func (v *NullableHistoryActionLabelUpdate) Set(val *HistoryActionLabelUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionLabelUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionLabelUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionLabelUpdate(val *HistoryActionLabelUpdate) *NullableHistoryActionLabelUpdate {
	return &NullableHistoryActionLabelUpdate{value: val, isSet: true}
}

func (v NullableHistoryActionLabelUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionLabelUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

