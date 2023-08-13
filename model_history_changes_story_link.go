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

// checks if the HistoryChangesStoryLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryChangesStoryLink{}

// HistoryChangesStoryLink The changes that have occurred as a result of the action.
type HistoryChangesStoryLink struct {
	Verb *StoryHistoryChangeOldNewStr `json:"verb,omitempty"`
	ObjectId *StoryHistoryChangeOldNewInt `json:"object_id,omitempty"`
	SubjectId *StoryHistoryChangeOldNewInt `json:"subject_id,omitempty"`
}

// NewHistoryChangesStoryLink instantiates a new HistoryChangesStoryLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryChangesStoryLink() *HistoryChangesStoryLink {
	this := HistoryChangesStoryLink{}
	return &this
}

// NewHistoryChangesStoryLinkWithDefaults instantiates a new HistoryChangesStoryLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryChangesStoryLinkWithDefaults() *HistoryChangesStoryLink {
	this := HistoryChangesStoryLink{}
	return &this
}

// GetVerb returns the Verb field value if set, zero value otherwise.
func (o *HistoryChangesStoryLink) GetVerb() StoryHistoryChangeOldNewStr {
	if o == nil || IsNil(o.Verb) {
		var ret StoryHistoryChangeOldNewStr
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryChangesStoryLink) GetVerbOk() (*StoryHistoryChangeOldNewStr, bool) {
	if o == nil || IsNil(o.Verb) {
		return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *HistoryChangesStoryLink) HasVerb() bool {
	if o != nil && !IsNil(o.Verb) {
		return true
	}

	return false
}

// SetVerb gets a reference to the given StoryHistoryChangeOldNewStr and assigns it to the Verb field.
func (o *HistoryChangesStoryLink) SetVerb(v StoryHistoryChangeOldNewStr) {
	o.Verb = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *HistoryChangesStoryLink) GetObjectId() StoryHistoryChangeOldNewInt {
	if o == nil || IsNil(o.ObjectId) {
		var ret StoryHistoryChangeOldNewInt
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryChangesStoryLink) GetObjectIdOk() (*StoryHistoryChangeOldNewInt, bool) {
	if o == nil || IsNil(o.ObjectId) {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *HistoryChangesStoryLink) HasObjectId() bool {
	if o != nil && !IsNil(o.ObjectId) {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given StoryHistoryChangeOldNewInt and assigns it to the ObjectId field.
func (o *HistoryChangesStoryLink) SetObjectId(v StoryHistoryChangeOldNewInt) {
	o.ObjectId = &v
}

// GetSubjectId returns the SubjectId field value if set, zero value otherwise.
func (o *HistoryChangesStoryLink) GetSubjectId() StoryHistoryChangeOldNewInt {
	if o == nil || IsNil(o.SubjectId) {
		var ret StoryHistoryChangeOldNewInt
		return ret
	}
	return *o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryChangesStoryLink) GetSubjectIdOk() (*StoryHistoryChangeOldNewInt, bool) {
	if o == nil || IsNil(o.SubjectId) {
		return nil, false
	}
	return o.SubjectId, true
}

// HasSubjectId returns a boolean if a field has been set.
func (o *HistoryChangesStoryLink) HasSubjectId() bool {
	if o != nil && !IsNil(o.SubjectId) {
		return true
	}

	return false
}

// SetSubjectId gets a reference to the given StoryHistoryChangeOldNewInt and assigns it to the SubjectId field.
func (o *HistoryChangesStoryLink) SetSubjectId(v StoryHistoryChangeOldNewInt) {
	o.SubjectId = &v
}

func (o HistoryChangesStoryLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryChangesStoryLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Verb) {
		toSerialize["verb"] = o.Verb
	}
	if !IsNil(o.ObjectId) {
		toSerialize["object_id"] = o.ObjectId
	}
	if !IsNil(o.SubjectId) {
		toSerialize["subject_id"] = o.SubjectId
	}
	return toSerialize, nil
}

type NullableHistoryChangesStoryLink struct {
	value *HistoryChangesStoryLink
	isSet bool
}

func (v NullableHistoryChangesStoryLink) Get() *HistoryChangesStoryLink {
	return v.value
}

func (v *NullableHistoryChangesStoryLink) Set(val *HistoryChangesStoryLink) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryChangesStoryLink) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryChangesStoryLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryChangesStoryLink(val *HistoryChangesStoryLink) *NullableHistoryChangesStoryLink {
	return &NullableHistoryChangesStoryLink{value: val, isSet: true}
}

func (v NullableHistoryChangesStoryLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryChangesStoryLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

