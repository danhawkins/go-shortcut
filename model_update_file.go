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

// checks if the UpdateFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFile{}

// UpdateFile struct for UpdateFile
type UpdateFile struct {
	// The description of the file.
	Description *string `json:"description,omitempty"`
	// The time/date that the file was uploaded.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The time/date that the file was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The name of the file.
	Name *string `json:"name,omitempty"`
	// The unique ID assigned to the Member who uploaded the file to Shortcut.
	UploaderId *string `json:"uploader_id,omitempty"`
	// An additional ID that you may wish to assign to the file.
	ExternalId *string `json:"external_id,omitempty"`
}

// NewUpdateFile instantiates a new UpdateFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFile() *UpdateFile {
	this := UpdateFile{}
	return &this
}

// NewUpdateFileWithDefaults instantiates a new UpdateFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFileWithDefaults() *UpdateFile {
	this := UpdateFile{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateFile) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFile) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateFile) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateFile) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UpdateFile) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFile) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UpdateFile) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UpdateFile) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UpdateFile) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFile) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UpdateFile) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UpdateFile) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateFile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateFile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateFile) SetName(v string) {
	o.Name = &v
}

// GetUploaderId returns the UploaderId field value if set, zero value otherwise.
func (o *UpdateFile) GetUploaderId() string {
	if o == nil || IsNil(o.UploaderId) {
		var ret string
		return ret
	}
	return *o.UploaderId
}

// GetUploaderIdOk returns a tuple with the UploaderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFile) GetUploaderIdOk() (*string, bool) {
	if o == nil || IsNil(o.UploaderId) {
		return nil, false
	}
	return o.UploaderId, true
}

// HasUploaderId returns a boolean if a field has been set.
func (o *UpdateFile) HasUploaderId() bool {
	if o != nil && !IsNil(o.UploaderId) {
		return true
	}

	return false
}

// SetUploaderId gets a reference to the given string and assigns it to the UploaderId field.
func (o *UpdateFile) SetUploaderId(v string) {
	o.UploaderId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *UpdateFile) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFile) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *UpdateFile) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *UpdateFile) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o UpdateFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UploaderId) {
		toSerialize["uploader_id"] = o.UploaderId
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	return toSerialize, nil
}

type NullableUpdateFile struct {
	value *UpdateFile
	isSet bool
}

func (v NullableUpdateFile) Get() *UpdateFile {
	return v.value
}

func (v *NullableUpdateFile) Set(val *UpdateFile) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFile) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFile(val *UpdateFile) *NullableUpdateFile {
	return &NullableUpdateFile{value: val, isSet: true}
}

func (v NullableUpdateFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

