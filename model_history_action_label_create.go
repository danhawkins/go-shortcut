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

// checks if the HistoryActionLabelCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionLabelCreate{}

// HistoryActionLabelCreate An action representing a Label being created.
type HistoryActionLabelCreate struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	// The application URL of the Label.
	AppUrl string `json:"app_url"`
	// The name of the Label.
	Name string `json:"name"`
}

// NewHistoryActionLabelCreate instantiates a new HistoryActionLabelCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionLabelCreate(id int64, entityType string, action string, appUrl string, name string) *HistoryActionLabelCreate {
	this := HistoryActionLabelCreate{}
	this.Id = id
	this.EntityType = entityType
	this.Action = action
	this.AppUrl = appUrl
	this.Name = name
	return &this
}

// NewHistoryActionLabelCreateWithDefaults instantiates a new HistoryActionLabelCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionLabelCreateWithDefaults() *HistoryActionLabelCreate {
	this := HistoryActionLabelCreate{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionLabelCreate) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelCreate) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionLabelCreate) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionLabelCreate) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelCreate) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionLabelCreate) SetEntityType(v string) {
	o.EntityType = v
}

// GetAction returns the Action field value
func (o *HistoryActionLabelCreate) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelCreate) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionLabelCreate) SetAction(v string) {
	o.Action = v
}

// GetAppUrl returns the AppUrl field value
func (o *HistoryActionLabelCreate) GetAppUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppUrl
}

// GetAppUrlOk returns a tuple with the AppUrl field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelCreate) GetAppUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUrl, true
}

// SetAppUrl sets field value
func (o *HistoryActionLabelCreate) SetAppUrl(v string) {
	o.AppUrl = v
}

// GetName returns the Name field value
func (o *HistoryActionLabelCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HistoryActionLabelCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HistoryActionLabelCreate) SetName(v string) {
	o.Name = v
}

func (o HistoryActionLabelCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionLabelCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["action"] = o.Action
	toSerialize["app_url"] = o.AppUrl
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableHistoryActionLabelCreate struct {
	value *HistoryActionLabelCreate
	isSet bool
}

func (v NullableHistoryActionLabelCreate) Get() *HistoryActionLabelCreate {
	return v.value
}

func (v *NullableHistoryActionLabelCreate) Set(val *HistoryActionLabelCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionLabelCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionLabelCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionLabelCreate(val *HistoryActionLabelCreate) *NullableHistoryActionLabelCreate {
	return &NullableHistoryActionLabelCreate{value: val, isSet: true}
}

func (v NullableHistoryActionLabelCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionLabelCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


