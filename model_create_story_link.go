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

// checks if the CreateStoryLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStoryLink{}

// CreateStoryLink struct for CreateStoryLink
type CreateStoryLink struct {
	// The type of link.
	Verb string `json:"verb"`
	// The ID of the subject Story.
	SubjectId int64 `json:"subject_id"`
	// The ID of the object Story.
	ObjectId int64 `json:"object_id"`
}

// NewCreateStoryLink instantiates a new CreateStoryLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStoryLink(verb string, subjectId int64, objectId int64) *CreateStoryLink {
	this := CreateStoryLink{}
	this.Verb = verb
	this.SubjectId = subjectId
	this.ObjectId = objectId
	return &this
}

// NewCreateStoryLinkWithDefaults instantiates a new CreateStoryLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStoryLinkWithDefaults() *CreateStoryLink {
	this := CreateStoryLink{}
	return &this
}

// GetVerb returns the Verb field value
func (o *CreateStoryLink) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// GetVerbOk returns a tuple with the Verb field value
// and a boolean to check if the value has been set.
func (o *CreateStoryLink) GetVerbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verb, true
}

// SetVerb sets field value
func (o *CreateStoryLink) SetVerb(v string) {
	o.Verb = v
}

// GetSubjectId returns the SubjectId field value
func (o *CreateStoryLink) GetSubjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value
// and a boolean to check if the value has been set.
func (o *CreateStoryLink) GetSubjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectId, true
}

// SetSubjectId sets field value
func (o *CreateStoryLink) SetSubjectId(v int64) {
	o.SubjectId = v
}

// GetObjectId returns the ObjectId field value
func (o *CreateStoryLink) GetObjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *CreateStoryLink) GetObjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *CreateStoryLink) SetObjectId(v int64) {
	o.ObjectId = v
}

func (o CreateStoryLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStoryLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["verb"] = o.Verb
	toSerialize["subject_id"] = o.SubjectId
	toSerialize["object_id"] = o.ObjectId
	return toSerialize, nil
}

type NullableCreateStoryLink struct {
	value *CreateStoryLink
	isSet bool
}

func (v NullableCreateStoryLink) Get() *CreateStoryLink {
	return v.value
}

func (v *NullableCreateStoryLink) Set(val *CreateStoryLink) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStoryLink) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStoryLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStoryLink(val *CreateStoryLink) *NullableCreateStoryLink {
	return &NullableCreateStoryLink{value: val, isSet: true}
}

func (v NullableCreateStoryLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStoryLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

