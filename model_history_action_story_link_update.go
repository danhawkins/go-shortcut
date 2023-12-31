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

// checks if the HistoryActionStoryLinkUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionStoryLinkUpdate{}

// HistoryActionStoryLinkUpdate An action representing a Story Link being updated.
type HistoryActionStoryLinkUpdate struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	// The verb describing the link's relationship.
	Verb string `json:"verb"`
	// The Story ID of the subject Story.
	SubjectId int64 `json:"subject_id"`
	// The Story ID of the object Story.
	ObjectId int64 `json:"object_id"`
	Changes HistoryChangesStoryLink `json:"changes"`
}

// NewHistoryActionStoryLinkUpdate instantiates a new HistoryActionStoryLinkUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionStoryLinkUpdate(id int64, entityType string, action string, verb string, subjectId int64, objectId int64, changes HistoryChangesStoryLink) *HistoryActionStoryLinkUpdate {
	this := HistoryActionStoryLinkUpdate{}
	this.Id = id
	this.EntityType = entityType
	this.Action = action
	this.Verb = verb
	this.SubjectId = subjectId
	this.ObjectId = objectId
	this.Changes = changes
	return &this
}

// NewHistoryActionStoryLinkUpdateWithDefaults instantiates a new HistoryActionStoryLinkUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionStoryLinkUpdateWithDefaults() *HistoryActionStoryLinkUpdate {
	this := HistoryActionStoryLinkUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionStoryLinkUpdate) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryLinkUpdate) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionStoryLinkUpdate) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionStoryLinkUpdate) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryLinkUpdate) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionStoryLinkUpdate) SetEntityType(v string) {
	o.EntityType = v
}

// GetAction returns the Action field value
func (o *HistoryActionStoryLinkUpdate) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryLinkUpdate) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionStoryLinkUpdate) SetAction(v string) {
	o.Action = v
}

// GetVerb returns the Verb field value
func (o *HistoryActionStoryLinkUpdate) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// GetVerbOk returns a tuple with the Verb field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryLinkUpdate) GetVerbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verb, true
}

// SetVerb sets field value
func (o *HistoryActionStoryLinkUpdate) SetVerb(v string) {
	o.Verb = v
}

// GetSubjectId returns the SubjectId field value
func (o *HistoryActionStoryLinkUpdate) GetSubjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryLinkUpdate) GetSubjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectId, true
}

// SetSubjectId sets field value
func (o *HistoryActionStoryLinkUpdate) SetSubjectId(v int64) {
	o.SubjectId = v
}

// GetObjectId returns the ObjectId field value
func (o *HistoryActionStoryLinkUpdate) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryLinkUpdate) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *HistoryActionStoryLinkUpdate) SetObjectId(v int64) {
	o.ObjectId = v
}

// GetChanges returns the Changes field value
func (o *HistoryActionStoryLinkUpdate) GetChanges() HistoryChangesStoryLink {
	if o == nil {
		var ret HistoryChangesStoryLink
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryLinkUpdate) GetChangesOk() (*HistoryChangesStoryLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Changes, true
}

// SetChanges sets field value
func (o *HistoryActionStoryLinkUpdate) SetChanges(v HistoryChangesStoryLink) {
	o.Changes = v
}

func (o HistoryActionStoryLinkUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionStoryLinkUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["action"] = o.Action
	toSerialize["verb"] = o.Verb
	toSerialize["subject_id"] = o.SubjectId
	toSerialize["object_id"] = o.ObjectId
	toSerialize["changes"] = o.Changes
	return toSerialize, nil
}

type NullableHistoryActionStoryLinkUpdate struct {
	value *HistoryActionStoryLinkUpdate
	isSet bool
}

func (v NullableHistoryActionStoryLinkUpdate) Get() *HistoryActionStoryLinkUpdate {
	return v.value
}

func (v *NullableHistoryActionStoryLinkUpdate) Set(val *HistoryActionStoryLinkUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionStoryLinkUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionStoryLinkUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionStoryLinkUpdate(val *HistoryActionStoryLinkUpdate) *NullableHistoryActionStoryLinkUpdate {
	return &NullableHistoryActionStoryLinkUpdate{value: val, isSet: true}
}

func (v NullableHistoryActionStoryLinkUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionStoryLinkUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


