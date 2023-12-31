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

// checks if the HistoryActionBranchCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionBranchCreate{}

// HistoryActionBranchCreate An action representing a VCS Branch being created.
type HistoryActionBranchCreate struct {
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

// NewHistoryActionBranchCreate instantiates a new HistoryActionBranchCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionBranchCreate(id int64, entityType string, name string, url string, action string) *HistoryActionBranchCreate {
	this := HistoryActionBranchCreate{}
	this.Id = id
	this.EntityType = entityType
	this.Name = name
	this.Url = url
	this.Action = action
	return &this
}

// NewHistoryActionBranchCreateWithDefaults instantiates a new HistoryActionBranchCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionBranchCreateWithDefaults() *HistoryActionBranchCreate {
	this := HistoryActionBranchCreate{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionBranchCreate) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionBranchCreate) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionBranchCreate) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionBranchCreate) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionBranchCreate) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionBranchCreate) SetEntityType(v string) {
	o.EntityType = v
}

// GetName returns the Name field value
func (o *HistoryActionBranchCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HistoryActionBranchCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HistoryActionBranchCreate) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *HistoryActionBranchCreate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *HistoryActionBranchCreate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *HistoryActionBranchCreate) SetUrl(v string) {
	o.Url = v
}

// GetAction returns the Action field value
func (o *HistoryActionBranchCreate) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionBranchCreate) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionBranchCreate) SetAction(v string) {
	o.Action = v
}

func (o HistoryActionBranchCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionBranchCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableHistoryActionBranchCreate struct {
	value *HistoryActionBranchCreate
	isSet bool
}

func (v NullableHistoryActionBranchCreate) Get() *HistoryActionBranchCreate {
	return v.value
}

func (v *NullableHistoryActionBranchCreate) Set(val *HistoryActionBranchCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionBranchCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionBranchCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionBranchCreate(val *HistoryActionBranchCreate) *NullableHistoryActionBranchCreate {
	return &NullableHistoryActionBranchCreate{value: val, isSet: true}
}

func (v NullableHistoryActionBranchCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionBranchCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


