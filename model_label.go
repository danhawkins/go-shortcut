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

// checks if the Label type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Label{}

// Label A Label can be used to associate and filter Stories and Epics, and also create new Workspaces.
type Label struct {
	// The Shortcut application url for the Label.
	AppUrl string `json:"app_url"`
	// The description of the Label.
	Description NullableString `json:"description"`
	// A true/false boolean indicating if the Label has been archived.
	Archived bool `json:"archived"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The hex color to be displayed with the Label (for example, \"#ff0000\").
	Color NullableString `json:"color"`
	// The name of the Label.
	Name string `json:"name"`
	GlobalId string `json:"global_id"`
	// The time/date that the Label was updated.
	UpdatedAt NullableTime `json:"updated_at"`
	// This field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId NullableString `json:"external_id"`
	// The unique ID of the Label.
	Id int64 `json:"id"`
	Stats *LabelStats `json:"stats,omitempty"`
	// The time/date that the Label was created.
	CreatedAt NullableTime `json:"created_at"`
}

// NewLabel instantiates a new Label object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabel(appUrl string, description NullableString, archived bool, entityType string, color NullableString, name string, globalId string, updatedAt NullableTime, externalId NullableString, id int64, createdAt NullableTime) *Label {
	this := Label{}
	this.AppUrl = appUrl
	this.Description = description
	this.Archived = archived
	this.EntityType = entityType
	this.Color = color
	this.Name = name
	this.GlobalId = globalId
	this.UpdatedAt = updatedAt
	this.ExternalId = externalId
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewLabelWithDefaults instantiates a new Label object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelWithDefaults() *Label {
	this := Label{}
	return &this
}

// GetAppUrl returns the AppUrl field value
func (o *Label) GetAppUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppUrl
}

// GetAppUrlOk returns a tuple with the AppUrl field value
// and a boolean to check if the value has been set.
func (o *Label) GetAppUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUrl, true
}

// SetAppUrl sets field value
func (o *Label) SetAppUrl(v string) {
	o.AppUrl = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Label) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Label) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *Label) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetArchived returns the Archived field value
func (o *Label) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *Label) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *Label) SetArchived(v bool) {
	o.Archived = v
}

// GetEntityType returns the EntityType field value
func (o *Label) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *Label) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *Label) SetEntityType(v string) {
	o.EntityType = v
}

// GetColor returns the Color field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Label) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}

	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Label) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// SetColor sets field value
func (o *Label) SetColor(v string) {
	o.Color.Set(&v)
}

// GetName returns the Name field value
func (o *Label) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Label) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Label) SetName(v string) {
	o.Name = v
}

// GetGlobalId returns the GlobalId field value
func (o *Label) GetGlobalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *Label) GetGlobalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *Label) SetGlobalId(v string) {
	o.GlobalId = v
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Label) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Label) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *Label) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// GetExternalId returns the ExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Label) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Label) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// SetExternalId sets field value
func (o *Label) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// GetId returns the Id field value
func (o *Label) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Label) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Label) SetId(v int64) {
	o.Id = v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *Label) GetStats() LabelStats {
	if o == nil || IsNil(o.Stats) {
		var ret LabelStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Label) GetStatsOk() (*LabelStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *Label) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given LabelStats and assigns it to the Stats field.
func (o *Label) SetStats(v LabelStats) {
	o.Stats = &v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Label) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Label) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *Label) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

func (o Label) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Label) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_url"] = o.AppUrl
	toSerialize["description"] = o.Description.Get()
	toSerialize["archived"] = o.Archived
	toSerialize["entity_type"] = o.EntityType
	toSerialize["color"] = o.Color.Get()
	toSerialize["name"] = o.Name
	toSerialize["global_id"] = o.GlobalId
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	toSerialize["external_id"] = o.ExternalId.Get()
	toSerialize["id"] = o.Id
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	toSerialize["created_at"] = o.CreatedAt.Get()
	return toSerialize, nil
}

type NullableLabel struct {
	value *Label
	isSet bool
}

func (v NullableLabel) Get() *Label {
	return v.value
}

func (v *NullableLabel) Set(val *Label) {
	v.value = val
	v.isSet = true
}

func (v NullableLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabel(val *Label) *NullableLabel {
	return &NullableLabel{value: val, isSet: true}
}

func (v NullableLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

