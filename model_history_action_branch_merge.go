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

// checks if the HistoryActionBranchMerge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionBranchMerge{}

// HistoryActionBranchMerge An action representing a VCS Branch being merged.
type HistoryActionBranchMerge struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The name of the VCS Branch that was pushed
	Name string `json:"name"`
	// The URL from the provider of the VCS Branch that was pushed
	Url string `json:"url"`
	// The action of the entity referenced.
	Action string `json:"action"`
}

// NewHistoryActionBranchMerge instantiates a new HistoryActionBranchMerge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionBranchMerge(id int64, entityType string, name string, url string, action string) *HistoryActionBranchMerge {
	this := HistoryActionBranchMerge{}
	this.Id = id
	this.EntityType = entityType
	this.Name = name
	this.Url = url
	this.Action = action
	return &this
}

// NewHistoryActionBranchMergeWithDefaults instantiates a new HistoryActionBranchMerge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionBranchMergeWithDefaults() *HistoryActionBranchMerge {
	this := HistoryActionBranchMerge{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionBranchMerge) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionBranchMerge) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionBranchMerge) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionBranchMerge) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionBranchMerge) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionBranchMerge) SetEntityType(v string) {
	o.EntityType = v
}

// GetName returns the Name field value
func (o *HistoryActionBranchMerge) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HistoryActionBranchMerge) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HistoryActionBranchMerge) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *HistoryActionBranchMerge) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *HistoryActionBranchMerge) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *HistoryActionBranchMerge) SetUrl(v string) {
	o.Url = v
}

// GetAction returns the Action field value
func (o *HistoryActionBranchMerge) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionBranchMerge) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionBranchMerge) SetAction(v string) {
	o.Action = v
}

func (o HistoryActionBranchMerge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionBranchMerge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableHistoryActionBranchMerge struct {
	value *HistoryActionBranchMerge
	isSet bool
}

func (v NullableHistoryActionBranchMerge) Get() *HistoryActionBranchMerge {
	return v.value
}

func (v *NullableHistoryActionBranchMerge) Set(val *HistoryActionBranchMerge) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionBranchMerge) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionBranchMerge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionBranchMerge(val *HistoryActionBranchMerge) *NullableHistoryActionBranchMerge {
	return &NullableHistoryActionBranchMerge{value: val, isSet: true}
}

func (v NullableHistoryActionBranchMerge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionBranchMerge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


