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

// checks if the HistoryReferenceEpic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryReferenceEpic{}

// HistoryReferenceEpic A reference to an Epic.
type HistoryReferenceEpic struct {
	// The ID of the entity referenced.
	Id map[string]interface{} `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The application URL of the Epic.
	AppUrl string `json:"app_url"`
	// The name of the entity referenced.
	Name string `json:"name"`
}

// NewHistoryReferenceEpic instantiates a new HistoryReferenceEpic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryReferenceEpic(id map[string]interface{}, entityType string, appUrl string, name string) *HistoryReferenceEpic {
	this := HistoryReferenceEpic{}
	this.Id = id
	this.EntityType = entityType
	this.AppUrl = appUrl
	this.Name = name
	return &this
}

// NewHistoryReferenceEpicWithDefaults instantiates a new HistoryReferenceEpic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryReferenceEpicWithDefaults() *HistoryReferenceEpic {
	this := HistoryReferenceEpic{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryReferenceEpic) GetId() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryReferenceEpic) GetIdOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *HistoryReferenceEpic) SetId(v map[string]interface{}) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryReferenceEpic) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryReferenceEpic) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryReferenceEpic) SetEntityType(v string) {
	o.EntityType = v
}

// GetAppUrl returns the AppUrl field value
func (o *HistoryReferenceEpic) GetAppUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppUrl
}

// GetAppUrlOk returns a tuple with the AppUrl field value
// and a boolean to check if the value has been set.
func (o *HistoryReferenceEpic) GetAppUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUrl, true
}

// SetAppUrl sets field value
func (o *HistoryReferenceEpic) SetAppUrl(v string) {
	o.AppUrl = v
}

// GetName returns the Name field value
func (o *HistoryReferenceEpic) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HistoryReferenceEpic) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HistoryReferenceEpic) SetName(v string) {
	o.Name = v
}

func (o HistoryReferenceEpic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryReferenceEpic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["app_url"] = o.AppUrl
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableHistoryReferenceEpic struct {
	value *HistoryReferenceEpic
	isSet bool
}

func (v NullableHistoryReferenceEpic) Get() *HistoryReferenceEpic {
	return v.value
}

func (v *NullableHistoryReferenceEpic) Set(val *HistoryReferenceEpic) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryReferenceEpic) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryReferenceEpic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryReferenceEpic(val *HistoryReferenceEpic) *NullableHistoryReferenceEpic {
	return &NullableHistoryReferenceEpic{value: val, isSet: true}
}

func (v NullableHistoryReferenceEpic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryReferenceEpic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


