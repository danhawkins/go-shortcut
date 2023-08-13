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

// checks if the HistoryActionLabelDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionLabelDelete{}

// HistoryActionLabelDelete An action representing a Label being deleted.
type HistoryActionLabelDelete struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	// The name of the Label.
	Name string `json:"name"`
}

// NewHistoryActionLabelDelete instantiates a new HistoryActionLabelDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionLabelDelete(id int64, entityType string, action string, name string) *HistoryActionLabelDelete {
	this := HistoryActionLabelDelete{}
	this.Id = id
	this.EntityType = entityType
	this.Action = action
	this.Name = name
	return &this
}

// NewHistoryActionLabelDeleteWithDefaults instantiates a new HistoryActionLabelDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionLabelDeleteWithDefaults() *HistoryActionLabelDelete {
	this := HistoryActionLabelDelete{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionLabelDelete) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelDelete) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionLabelDelete) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionLabelDelete) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelDelete) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionLabelDelete) SetEntityType(v string) {
	o.EntityType = v
}

// GetAction returns the Action field value
func (o *HistoryActionLabelDelete) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelDelete) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionLabelDelete) SetAction(v string) {
	o.Action = v
}

// GetName returns the Name field value
func (o *HistoryActionLabelDelete) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelDelete) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HistoryActionLabelDelete) SetName(v string) {
	o.Name = v
}

func (o HistoryActionLabelDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionLabelDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["action"] = o.Action
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableHistoryActionLabelDelete struct {
	value *HistoryActionLabelDelete
	isSet bool
}

func (v NullableHistoryActionLabelDelete) Get() *HistoryActionLabelDelete {
	return v.value
}

func (v *NullableHistoryActionLabelDelete) Set(val *HistoryActionLabelDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionLabelDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionLabelDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionLabelDelete(val *HistoryActionLabelDelete) *NullableHistoryActionLabelDelete {
	return &NullableHistoryActionLabelDelete{value: val, isSet: true}
}

func (v NullableHistoryActionLabelDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionLabelDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


