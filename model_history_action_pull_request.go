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

// checks if the HistoryActionPullRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionPullRequest{}

// HistoryActionPullRequest An action representing various operations for a Pull Request.
type HistoryActionPullRequest struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	// The VCS Repository-specific ID for the Pull Request.
	Number int64 `json:"number"`
	// The title of the Pull Request.
	Title string `json:"title"`
	// The URL from the provider of the VCS Pull Request.
	Url string `json:"url"`
}

// NewHistoryActionPullRequest instantiates a new HistoryActionPullRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionPullRequest(id int64, entityType string, action string, number int64, title string, url string) *HistoryActionPullRequest {
	this := HistoryActionPullRequest{}
	this.Id = id
	this.EntityType = entityType
	this.Action = action
	this.Number = number
	this.Title = title
	this.Url = url
	return &this
}

// NewHistoryActionPullRequestWithDefaults instantiates a new HistoryActionPullRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionPullRequestWithDefaults() *HistoryActionPullRequest {
	this := HistoryActionPullRequest{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionPullRequest) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionPullRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionPullRequest) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionPullRequest) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionPullRequest) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionPullRequest) SetEntityType(v string) {
	o.EntityType = v
}

// GetAction returns the Action field value
func (o *HistoryActionPullRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionPullRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionPullRequest) SetAction(v string) {
	o.Action = v
}

// GetNumber returns the Number field value
func (o *HistoryActionPullRequest) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *HistoryActionPullRequest) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *HistoryActionPullRequest) SetNumber(v int64) {
	o.Number = v
}

// GetTitle returns the Title field value
func (o *HistoryActionPullRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *HistoryActionPullRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *HistoryActionPullRequest) SetTitle(v string) {
	o.Title = v
}

// GetUrl returns the Url field value
func (o *HistoryActionPullRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *HistoryActionPullRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *HistoryActionPullRequest) SetUrl(v string) {
	o.Url = v
}

func (o HistoryActionPullRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionPullRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["action"] = o.Action
	toSerialize["number"] = o.Number
	toSerialize["title"] = o.Title
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableHistoryActionPullRequest struct {
	value *HistoryActionPullRequest
	isSet bool
}

func (v NullableHistoryActionPullRequest) Get() *HistoryActionPullRequest {
	return v.value
}

func (v *NullableHistoryActionPullRequest) Set(val *HistoryActionPullRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionPullRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionPullRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionPullRequest(val *HistoryActionPullRequest) *NullableHistoryActionPullRequest {
	return &NullableHistoryActionPullRequest{value: val, isSet: true}
}

func (v NullableHistoryActionPullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionPullRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


