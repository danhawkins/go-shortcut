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

// checks if the CreateTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTask{}

// CreateTask struct for CreateTask
type CreateTask struct {
	// The Task description.
	Description string `json:"description"`
	// True/false boolean indicating whether the Task is completed. Defaults to false.
	Complete *bool `json:"complete,omitempty"`
	// An array of UUIDs for any members you want to add as Owners on this new Task.
	OwnerIds []string `json:"owner_ids,omitempty"`
	// Defaults to the time/date the Task is created but can be set to reflect another creation time/date.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Defaults to the time/date the Task is created in Shortcut but can be set to reflect another time/date.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id,omitempty"`
}

// NewCreateTask instantiates a new CreateTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTask(description string) *CreateTask {
	this := CreateTask{}
	this.Description = description
	return &this
}

// NewCreateTaskWithDefaults instantiates a new CreateTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTaskWithDefaults() *CreateTask {
	this := CreateTask{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateTask) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateTask) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateTask) SetDescription(v string) {
	o.Description = v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *CreateTask) GetComplete() bool {
	if o == nil || IsNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTask) GetCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *CreateTask) HasComplete() bool {
	if o != nil && !IsNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *CreateTask) SetComplete(v bool) {
	o.Complete = &v
}

// GetOwnerIds returns the OwnerIds field value if set, zero value otherwise.
func (o *CreateTask) GetOwnerIds() []string {
	if o == nil || IsNil(o.OwnerIds) {
		var ret []string
		return ret
	}
	return o.OwnerIds
}

// GetOwnerIdsOk returns a tuple with the OwnerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTask) GetOwnerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerIds) {
		return nil, false
	}
	return o.OwnerIds, true
}

// HasOwnerIds returns a boolean if a field has been set.
func (o *CreateTask) HasOwnerIds() bool {
	if o != nil && !IsNil(o.OwnerIds) {
		return true
	}

	return false
}

// SetOwnerIds gets a reference to the given []string and assigns it to the OwnerIds field.
func (o *CreateTask) SetOwnerIds(v []string) {
	o.OwnerIds = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateTask) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTask) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateTask) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateTask) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreateTask) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTask) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreateTask) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CreateTask) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *CreateTask) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTask) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *CreateTask) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *CreateTask) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o CreateTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if !IsNil(o.Complete) {
		toSerialize["complete"] = o.Complete
	}
	if !IsNil(o.OwnerIds) {
		toSerialize["owner_ids"] = o.OwnerIds
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	return toSerialize, nil
}

type NullableCreateTask struct {
	value *CreateTask
	isSet bool
}

func (v NullableCreateTask) Get() *CreateTask {
	return v.value
}

func (v *NullableCreateTask) Set(val *CreateTask) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTask) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTask(val *CreateTask) *NullableCreateTask {
	return &NullableCreateTask{value: val, isSet: true}
}

func (v NullableCreateTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

