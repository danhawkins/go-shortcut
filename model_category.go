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

// checks if the Category type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Category{}

// Category A Category can be used to associate Milestones.
type Category struct {
	// A true/false boolean indicating if the Category has been archived.
	Archived bool `json:"archived"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The hex color to be displayed with the Category (for example, \"#ff0000\").
	Color NullableString `json:"color"`
	// The name of the Category.
	Name string `json:"name"`
	// The Global ID of the Category.
	GlobalId string `json:"global_id"`
	// The type of entity this Category is associated with; currently Milestone is the only type of Category.
	Type string `json:"type"`
	// The time/date that the Category was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// This field can be set to another unique ID. In the case that the Category has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId NullableString `json:"external_id"`
	// The unique ID of the Category.
	Id int64 `json:"id"`
	// The time/date that the Category was created.
	CreatedAt time.Time `json:"created_at"`
}

// NewCategory instantiates a new Category object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategory(archived bool, entityType string, color NullableString, name string, globalId string, type_ string, updatedAt time.Time, externalId NullableString, id int64, createdAt time.Time) *Category {
	this := Category{}
	this.Archived = archived
	this.EntityType = entityType
	this.Color = color
	this.Name = name
	this.GlobalId = globalId
	this.Type = type_
	this.UpdatedAt = updatedAt
	this.ExternalId = externalId
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewCategoryWithDefaults instantiates a new Category object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryWithDefaults() *Category {
	this := Category{}
	return &this
}

// GetArchived returns the Archived field value
func (o *Category) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *Category) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *Category) SetArchived(v bool) {
	o.Archived = v
}

// GetEntityType returns the EntityType field value
func (o *Category) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *Category) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *Category) SetEntityType(v string) {
	o.EntityType = v
}

// GetColor returns the Color field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Category) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}

	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Category) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// SetColor sets field value
func (o *Category) SetColor(v string) {
	o.Color.Set(&v)
}

// GetName returns the Name field value
func (o *Category) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Category) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Category) SetName(v string) {
	o.Name = v
}

// GetGlobalId returns the GlobalId field value
func (o *Category) GetGlobalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *Category) GetGlobalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *Category) SetGlobalId(v string) {
	o.GlobalId = v
}

// GetType returns the Type field value
func (o *Category) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Category) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Category) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Category) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Category) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Category) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetExternalId returns the ExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Category) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Category) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// SetExternalId sets field value
func (o *Category) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// GetId returns the Id field value
func (o *Category) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Category) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Category) SetId(v int64) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Category) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Category) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Category) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o Category) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Category) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["archived"] = o.Archived
	toSerialize["entity_type"] = o.EntityType
	toSerialize["color"] = o.Color.Get()
	toSerialize["name"] = o.Name
	toSerialize["global_id"] = o.GlobalId
	toSerialize["type"] = o.Type
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["external_id"] = o.ExternalId.Get()
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

type NullableCategory struct {
	value *Category
	isSet bool
}

func (v NullableCategory) Get() *Category {
	return v.value
}

func (v *NullableCategory) Set(val *Category) {
	v.value = val
	v.isSet = true
}

func (v NullableCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategory(val *Category) *NullableCategory {
	return &NullableCategory{value: val, isSet: true}
}

func (v NullableCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


