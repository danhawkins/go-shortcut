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

// checks if the HistoryActionTaskUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionTaskUpdate{}

// HistoryActionTaskUpdate An action representing a Task being updated.
type HistoryActionTaskUpdate struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	Changes HistoryChangesTask `json:"changes"`
	// Whether or not the Task is complete.
	Complete *bool `json:"complete,omitempty"`
	// The description of the Task.
	Description string `json:"description"`
	// The Story ID that contains the Task.
	StoryId int64 `json:"story_id"`
}

// NewHistoryActionTaskUpdate instantiates a new HistoryActionTaskUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionTaskUpdate(id int64, entityType string, action string, changes HistoryChangesTask, description string, storyId int64) *HistoryActionTaskUpdate {
	this := HistoryActionTaskUpdate{}
	this.Id = id
	this.EntityType = entityType
	this.Action = action
	this.Changes = changes
	this.Description = description
	this.StoryId = storyId
	return &this
}

// NewHistoryActionTaskUpdateWithDefaults instantiates a new HistoryActionTaskUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionTaskUpdateWithDefaults() *HistoryActionTaskUpdate {
	this := HistoryActionTaskUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionTaskUpdate) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskUpdate) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionTaskUpdate) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionTaskUpdate) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskUpdate) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionTaskUpdate) SetEntityType(v string) {
	o.EntityType = v
}

// GetAction returns the Action field value
func (o *HistoryActionTaskUpdate) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskUpdate) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionTaskUpdate) SetAction(v string) {
	o.Action = v
}

// GetChanges returns the Changes field value
func (o *HistoryActionTaskUpdate) GetChanges() HistoryChangesTask {
	if o == nil {
		var ret HistoryChangesTask
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskUpdate) GetChangesOk() (*HistoryChangesTask, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Changes, true
}

// SetChanges sets field value
func (o *HistoryActionTaskUpdate) SetChanges(v HistoryChangesTask) {
	o.Changes = v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *HistoryActionTaskUpdate) GetComplete() bool {
	if o == nil || IsNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskUpdate) GetCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *HistoryActionTaskUpdate) HasComplete() bool {
	if o != nil && !IsNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *HistoryActionTaskUpdate) SetComplete(v bool) {
	o.Complete = &v
}

// GetDescription returns the Description field value
func (o *HistoryActionTaskUpdate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *HistoryActionTaskUpdate) SetDescription(v string) {
	o.Description = v
}

// GetStoryId returns the StoryId field value
func (o *HistoryActionTaskUpdate) GetStoryId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StoryId
}

// GetStoryIdOk returns a tuple with the StoryId field value
// and a boolean to check if the value has been set.
func (o *HistoryActionTaskUpdate) GetStoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoryId, true
}

// SetStoryId sets field value
func (o *HistoryActionTaskUpdate) SetStoryId(v int64) {
	o.StoryId = v
}

func (o HistoryActionTaskUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionTaskUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["action"] = o.Action
	toSerialize["changes"] = o.Changes
	if !IsNil(o.Complete) {
		toSerialize["complete"] = o.Complete
	}
	toSerialize["description"] = o.Description
	toSerialize["story_id"] = o.StoryId
	return toSerialize, nil
}

type NullableHistoryActionTaskUpdate struct {
	value *HistoryActionTaskUpdate
	isSet bool
}

func (v NullableHistoryActionTaskUpdate) Get() *HistoryActionTaskUpdate {
	return v.value
}

func (v *NullableHistoryActionTaskUpdate) Set(val *HistoryActionTaskUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionTaskUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionTaskUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionTaskUpdate(val *HistoryActionTaskUpdate) *NullableHistoryActionTaskUpdate {
	return &NullableHistoryActionTaskUpdate{value: val, isSet: true}
}

func (v NullableHistoryActionTaskUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionTaskUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

