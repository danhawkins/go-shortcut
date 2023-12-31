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

// checks if the HistoryReferenceWorkflowState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryReferenceWorkflowState{}

// HistoryReferenceWorkflowState A references to a Story Workflow State.
type HistoryReferenceWorkflowState struct {
	// The ID of the entity referenced.
	Id map[string]interface{} `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// Either \"unstarted\", \"started\", or \"done\".
	Type string `json:"type"`
	// The name of the entity referenced.
	Name string `json:"name"`
}

// NewHistoryReferenceWorkflowState instantiates a new HistoryReferenceWorkflowState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryReferenceWorkflowState(id map[string]interface{}, entityType string, type_ string, name string) *HistoryReferenceWorkflowState {
	this := HistoryReferenceWorkflowState{}
	this.Id = id
	this.EntityType = entityType
	this.Type = type_
	this.Name = name
	return &this
}

// NewHistoryReferenceWorkflowStateWithDefaults instantiates a new HistoryReferenceWorkflowState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryReferenceWorkflowStateWithDefaults() *HistoryReferenceWorkflowState {
	this := HistoryReferenceWorkflowState{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryReferenceWorkflowState) GetId() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryReferenceWorkflowState) GetIdOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *HistoryReferenceWorkflowState) SetId(v map[string]interface{}) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryReferenceWorkflowState) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryReferenceWorkflowState) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryReferenceWorkflowState) SetEntityType(v string) {
	o.EntityType = v
}

// GetType returns the Type field value
func (o *HistoryReferenceWorkflowState) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HistoryReferenceWorkflowState) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HistoryReferenceWorkflowState) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *HistoryReferenceWorkflowState) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HistoryReferenceWorkflowState) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HistoryReferenceWorkflowState) SetName(v string) {
	o.Name = v
}

func (o HistoryReferenceWorkflowState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryReferenceWorkflowState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableHistoryReferenceWorkflowState struct {
	value *HistoryReferenceWorkflowState
	isSet bool
}

func (v NullableHistoryReferenceWorkflowState) Get() *HistoryReferenceWorkflowState {
	return v.value
}

func (v *NullableHistoryReferenceWorkflowState) Set(val *HistoryReferenceWorkflowState) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryReferenceWorkflowState) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryReferenceWorkflowState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryReferenceWorkflowState(val *HistoryReferenceWorkflowState) *NullableHistoryReferenceWorkflowState {
	return &NullableHistoryReferenceWorkflowState{value: val, isSet: true}
}

func (v NullableHistoryReferenceWorkflowState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryReferenceWorkflowState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


