/*
Shortcut API

Shortcut API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the EpicState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EpicState{}

// EpicState Epic State is any of the at least 3 columns. Epic States correspond to one of 3 types: Unstarted, Started, or Done.
type EpicState struct {
	// The description of what sort of Epics belong in that Epic State.
	Description string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The hex color for this Epic State.
	Color *string `json:"color,omitempty"`
	// The Epic State's name.
	Name string `json:"name"`
	GlobalId string `json:"global_id"`
	// The type of Epic State (Unstarted, Started, or Done)
	Type string `json:"type"`
	// When the Epic State was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The unique ID of the Epic State.
	Id int64 `json:"id"`
	// The position that the Epic State is in, starting with 0 at the left.
	Position int64 `json:"position"`
	// The time/date the Epic State was created.
	CreatedAt time.Time `json:"created_at"`
}

// NewEpicState instantiates a new EpicState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpicState(description string, entityType string, name string, globalId string, type_ string, updatedAt time.Time, id int64, position int64, createdAt time.Time) *EpicState {
	this := EpicState{}
	this.Description = description
	this.EntityType = entityType
	this.Name = name
	this.GlobalId = globalId
	this.Type = type_
	this.UpdatedAt = updatedAt
	this.Id = id
	this.Position = position
	this.CreatedAt = createdAt
	return &this
}

// NewEpicStateWithDefaults instantiates a new EpicState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpicStateWithDefaults() *EpicState {
	this := EpicState{}
	return &this
}

// GetDescription returns the Description field value
func (o *EpicState) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EpicState) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EpicState) SetDescription(v string) {
	o.Description = v
}

// GetEntityType returns the EntityType field value
func (o *EpicState) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *EpicState) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *EpicState) SetEntityType(v string) {
	o.EntityType = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *EpicState) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpicState) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *EpicState) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *EpicState) SetColor(v string) {
	o.Color = &v
}

// GetName returns the Name field value
func (o *EpicState) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EpicState) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EpicState) SetName(v string) {
	o.Name = v
}

// GetGlobalId returns the GlobalId field value
func (o *EpicState) GetGlobalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *EpicState) GetGlobalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *EpicState) SetGlobalId(v string) {
	o.GlobalId = v
}

// GetType returns the Type field value
func (o *EpicState) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EpicState) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EpicState) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EpicState) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EpicState) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EpicState) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetId returns the Id field value
func (o *EpicState) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EpicState) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EpicState) SetId(v int64) {
	o.Id = v
}

// GetPosition returns the Position field value
func (o *EpicState) GetPosition() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *EpicState) GetPositionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *EpicState) SetPosition(v int64) {
	o.Position = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EpicState) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EpicState) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EpicState) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o EpicState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpicState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["entity_type"] = o.EntityType
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	toSerialize["name"] = o.Name
	toSerialize["global_id"] = o.GlobalId
	toSerialize["type"] = o.Type
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["id"] = o.Id
	toSerialize["position"] = o.Position
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

type NullableEpicState struct {
	value *EpicState
	isSet bool
}

func (v NullableEpicState) Get() *EpicState {
	return v.value
}

func (v *NullableEpicState) Set(val *EpicState) {
	v.value = val
	v.isSet = true
}

func (v NullableEpicState) IsSet() bool {
	return v.isSet
}

func (v *NullableEpicState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpicState(val *EpicState) *NullableEpicState {
	return &NullableEpicState{value: val, isSet: true}
}

func (v NullableEpicState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpicState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


