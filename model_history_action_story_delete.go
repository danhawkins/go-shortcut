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

// checks if the HistoryActionStoryDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionStoryDelete{}

// HistoryActionStoryDelete An action representing a Story being deleted.
type HistoryActionStoryDelete struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	// The name of the Story.
	Name string `json:"name"`
	// The type of Story; either feature, bug, or chore.
	StoryType string `json:"story_type"`
}

// NewHistoryActionStoryDelete instantiates a new HistoryActionStoryDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionStoryDelete(id int64, entityType string, action string, name string, storyType string) *HistoryActionStoryDelete {
	this := HistoryActionStoryDelete{}
	this.Id = id
	this.EntityType = entityType
	this.Action = action
	this.Name = name
	this.StoryType = storyType
	return &this
}

// NewHistoryActionStoryDeleteWithDefaults instantiates a new HistoryActionStoryDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionStoryDeleteWithDefaults() *HistoryActionStoryDelete {
	this := HistoryActionStoryDelete{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionStoryDelete) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryDelete) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionStoryDelete) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionStoryDelete) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryDelete) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionStoryDelete) SetEntityType(v string) {
	o.EntityType = v
}

// GetAction returns the Action field value
func (o *HistoryActionStoryDelete) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryDelete) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionStoryDelete) SetAction(v string) {
	o.Action = v
}

// GetName returns the Name field value
func (o *HistoryActionStoryDelete) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryDelete) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HistoryActionStoryDelete) SetName(v string) {
	o.Name = v
}

// GetStoryType returns the StoryType field value
func (o *HistoryActionStoryDelete) GetStoryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StoryType
}

// GetStoryTypeOk returns a tuple with the StoryType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryDelete) GetStoryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoryType, true
}

// SetStoryType sets field value
func (o *HistoryActionStoryDelete) SetStoryType(v string) {
	o.StoryType = v
}

func (o HistoryActionStoryDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionStoryDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["action"] = o.Action
	toSerialize["name"] = o.Name
	toSerialize["story_type"] = o.StoryType
	return toSerialize, nil
}

type NullableHistoryActionStoryDelete struct {
	value *HistoryActionStoryDelete
	isSet bool
}

func (v NullableHistoryActionStoryDelete) Get() *HistoryActionStoryDelete {
	return v.value
}

func (v *NullableHistoryActionStoryDelete) Set(val *HistoryActionStoryDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionStoryDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionStoryDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionStoryDelete(val *HistoryActionStoryDelete) *NullableHistoryActionStoryDelete {
	return &NullableHistoryActionStoryDelete{value: val, isSet: true}
}

func (v NullableHistoryActionStoryDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionStoryDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


