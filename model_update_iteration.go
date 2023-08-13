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

// checks if the UpdateIteration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIteration{}

// UpdateIteration struct for UpdateIteration
type UpdateIteration struct {
	// An array of UUIDs for any Members you want to add as Followers.
	FollowerIds []string `json:"follower_ids,omitempty"`
	// An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	GroupIds []string `json:"group_ids,omitempty"`
	// An array of Labels attached to the Iteration.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// The description of the Iteration.
	Description *string `json:"description,omitempty"`
	// The name of this Iteration
	Name *string `json:"name,omitempty"`
	// The date this Iteration begins, e.g. 2019-07-01
	StartDate *string `json:"start_date,omitempty"`
	// The date this Iteration ends, e.g. 2019-07-05.
	EndDate *string `json:"end_date,omitempty"`
}

// NewUpdateIteration instantiates a new UpdateIteration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIteration() *UpdateIteration {
	this := UpdateIteration{}
	return &this
}

// NewUpdateIterationWithDefaults instantiates a new UpdateIteration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIterationWithDefaults() *UpdateIteration {
	this := UpdateIteration{}
	return &this
}

// GetFollowerIds returns the FollowerIds field value if set, zero value otherwise.
func (o *UpdateIteration) GetFollowerIds() []string {
	if o == nil || IsNil(o.FollowerIds) {
		var ret []string
		return ret
	}
	return o.FollowerIds
}

// GetFollowerIdsOk returns a tuple with the FollowerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIteration) GetFollowerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FollowerIds) {
		return nil, false
	}
	return o.FollowerIds, true
}

// HasFollowerIds returns a boolean if a field has been set.
func (o *UpdateIteration) HasFollowerIds() bool {
	if o != nil && !IsNil(o.FollowerIds) {
		return true
	}

	return false
}

// SetFollowerIds gets a reference to the given []string and assigns it to the FollowerIds field.
func (o *UpdateIteration) SetFollowerIds(v []string) {
	o.FollowerIds = v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *UpdateIteration) GetGroupIds() []string {
	if o == nil || IsNil(o.GroupIds) {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIteration) GetGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupIds) {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *UpdateIteration) HasGroupIds() bool {
	if o != nil && !IsNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *UpdateIteration) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateIteration) GetLabels() []CreateLabelParams {
	if o == nil || IsNil(o.Labels) {
		var ret []CreateLabelParams
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIteration) GetLabelsOk() ([]CreateLabelParams, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateIteration) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []CreateLabelParams and assigns it to the Labels field.
func (o *UpdateIteration) SetLabels(v []CreateLabelParams) {
	o.Labels = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateIteration) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIteration) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateIteration) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateIteration) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateIteration) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIteration) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateIteration) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateIteration) SetName(v string) {
	o.Name = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UpdateIteration) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIteration) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UpdateIteration) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *UpdateIteration) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *UpdateIteration) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIteration) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *UpdateIteration) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *UpdateIteration) SetEndDate(v string) {
	o.EndDate = &v
}

func (o UpdateIteration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIteration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FollowerIds) {
		toSerialize["follower_ids"] = o.FollowerIds
	}
	if !IsNil(o.GroupIds) {
		toSerialize["group_ids"] = o.GroupIds
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableUpdateIteration struct {
	value *UpdateIteration
	isSet bool
}

func (v NullableUpdateIteration) Get() *UpdateIteration {
	return v.value
}

func (v *NullableUpdateIteration) Set(val *UpdateIteration) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIteration) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIteration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIteration(val *UpdateIteration) *NullableUpdateIteration {
	return &NullableUpdateIteration{value: val, isSet: true}
}

func (v NullableUpdateIteration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIteration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


