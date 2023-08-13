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

// checks if the HistoryActionStoryCommentCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryActionStoryCommentCreate{}

// HistoryActionStoryCommentCreate An action representing a Story Comment being created.
type HistoryActionStoryCommentCreate struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	// The application URL of the Story Comment.
	AppUrl string `json:"app_url"`
	// The text of the Story Comment.
	Text string `json:"text"`
	// The Member ID of who created the Story Comment.
	AuthorId string `json:"author_id"`
}

// NewHistoryActionStoryCommentCreate instantiates a new HistoryActionStoryCommentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryActionStoryCommentCreate(id int64, entityType string, action string, appUrl string, text string, authorId string) *HistoryActionStoryCommentCreate {
	this := HistoryActionStoryCommentCreate{}
	this.Id = id
	this.EntityType = entityType
	this.Action = action
	this.AppUrl = appUrl
	this.Text = text
	this.AuthorId = authorId
	return &this
}

// NewHistoryActionStoryCommentCreateWithDefaults instantiates a new HistoryActionStoryCommentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryActionStoryCommentCreateWithDefaults() *HistoryActionStoryCommentCreate {
	this := HistoryActionStoryCommentCreate{}
	return &this
}

// GetId returns the Id field value
func (o *HistoryActionStoryCommentCreate) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryCommentCreate) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HistoryActionStoryCommentCreate) SetId(v int64) {
	o.Id = v
}

// GetEntityType returns the EntityType field value
func (o *HistoryActionStoryCommentCreate) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryCommentCreate) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *HistoryActionStoryCommentCreate) SetEntityType(v string) {
	o.EntityType = v
}

// GetAction returns the Action field value
func (o *HistoryActionStoryCommentCreate) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryCommentCreate) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *HistoryActionStoryCommentCreate) SetAction(v string) {
	o.Action = v
}

// GetAppUrl returns the AppUrl field value
func (o *HistoryActionStoryCommentCreate) GetAppUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppUrl
}

// GetAppUrlOk returns a tuple with the AppUrl field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryCommentCreate) GetAppUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUrl, true
}

// SetAppUrl sets field value
func (o *HistoryActionStoryCommentCreate) SetAppUrl(v string) {
	o.AppUrl = v
}

// GetText returns the Text field value
func (o *HistoryActionStoryCommentCreate) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryCommentCreate) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *HistoryActionStoryCommentCreate) SetText(v string) {
	o.Text = v
}

// GetAuthorId returns the AuthorId field value
func (o *HistoryActionStoryCommentCreate) GetAuthorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
func (o *HistoryActionStoryCommentCreate) GetAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorId, true
}

// SetAuthorId sets field value
func (o *HistoryActionStoryCommentCreate) SetAuthorId(v string) {
	o.AuthorId = v
}

func (o HistoryActionStoryCommentCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryActionStoryCommentCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["entity_type"] = o.EntityType
	toSerialize["action"] = o.Action
	toSerialize["app_url"] = o.AppUrl
	toSerialize["text"] = o.Text
	toSerialize["author_id"] = o.AuthorId
	return toSerialize, nil
}

type NullableHistoryActionStoryCommentCreate struct {
	value *HistoryActionStoryCommentCreate
	isSet bool
}

func (v NullableHistoryActionStoryCommentCreate) Get() *HistoryActionStoryCommentCreate {
	return v.value
}

func (v *NullableHistoryActionStoryCommentCreate) Set(val *HistoryActionStoryCommentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryActionStoryCommentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryActionStoryCommentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryActionStoryCommentCreate(val *HistoryActionStoryCommentCreate) *NullableHistoryActionStoryCommentCreate {
	return &NullableHistoryActionStoryCommentCreate{value: val, isSet: true}
}

func (v NullableHistoryActionStoryCommentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryActionStoryCommentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


