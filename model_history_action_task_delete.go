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

// checks if the HistoryActionTaskDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionTaskDelete{}

// HistoryActionTaskDelete An action representing a Task being deleted.
type HistoryActionTaskDelete struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	// The description of the Task being deleted.
	Description string `json:"description"`
}

// NewHistoryActionTaskDelete instantiates a new HistoryActionTaskDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionTaskDelete(id int64, entityType string, action string, description string) *HistoryActionTaskDelete {
	this := HistoryActionTaskDelete{}
	this.Id = id
	this.EntityType = entityType
	this.Action = action
	this.Description = description
	return &this
}

// NewHistoryActionTaskDeleteWithDefaults instantiates a new HistoryActionTaskDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionTaskDeleteWithDefaults() *HistoryActionTaskDelete {
	this := HistoryActionTaskDelete{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionTaskDelete) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskDelete) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionTaskDelete) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionTaskDelete) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskDelete) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionTaskDelete) SetEntityType(v string) {
	o.EntityType = v
}

// GetAction returns the Action field value
func (o *HistoryActionTaskDelete) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskDelete) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionTaskDelete) SetAction(v string) {
	o.Action = v
}

// GetDescription returns the Description field value
func (o *HistoryActionTaskDelete) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskDelete) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *HistoryActionTaskDelete) SetDescription(v string) {
	o.Description = v
}

func (o HistoryActionTaskDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionTaskDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["action"] = o.Action
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableHistoryActionTaskDelete struct {
	value *HistoryActionTaskDelete
	isSet bool
}

func (v NullableHistoryActionTaskDelete) Get() *HistoryActionTaskDelete {
	return v.value
}

func (v *NullableHistoryActionTaskDelete) Set(val *HistoryActionTaskDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionTaskDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionTaskDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionTaskDelete(val *HistoryActionTaskDelete) *NullableHistoryActionTaskDelete {
	return &NullableHistoryActionTaskDelete{value: val, isSet: true}
}

func (v NullableHistoryActionTaskDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionTaskDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

