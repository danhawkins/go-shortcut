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

// checks if the ThreadedComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadedComment{}

// ThreadedComment Comments associated with Epic Discussions.
type ThreadedComment struct {
	// The Shortcut application url for the Comment.
	AppUrl string `json:"app_url"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// True/false boolean indicating whether the Comment is deleted.
	Deleted bool `json:"deleted"`
	// Deprecated: use member_mention_ids.
	MentionIds []string `json:"mention_ids"`
	// The unique ID of the Member that authored the Comment.
	AuthorId string `json:"author_id"`
	// An array of Member IDs that have been mentioned in this Comment.
	MemberMentionIds []string `json:"member_mention_ids"`
	// A nested array of threaded comments.
	Comments []ThreadedComment `json:"comments"`
	// The time/date the Comment was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// An array of Group IDs that have been mentioned in this Comment.
	GroupMentionIds []string `json:"group_mention_ids"`
	// This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId NullableString `json:"external_id"`
	// The unique ID of the Comment.
	Id int64 `json:"id"`
	// The time/date the Comment was created.
	CreatedAt time.Time `json:"created_at"`
	// The text of the Comment.
	Text string `json:"text"`
}

// NewThreadedComment instantiates a new ThreadedComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadedComment(appUrl string, entityType string, deleted bool, mentionIds []string, authorId string, memberMentionIds []string, comments []ThreadedComment, updatedAt time.Time, groupMentionIds []string, externalId NullableString, id int64, createdAt time.Time, text string) *ThreadedComment {
	this := ThreadedComment{}
	this.AppUrl = appUrl
	this.EntityType = entityType
	this.Deleted = deleted
	this.MentionIds = mentionIds
	this.AuthorId = authorId
	this.MemberMentionIds = memberMentionIds
	this.Comments = comments
	this.UpdatedAt = updatedAt
	this.GroupMentionIds = groupMentionIds
	this.ExternalId = externalId
	this.Id = id
	this.CreatedAt = createdAt
	this.Text = text
	return &this
}

// NewThreadedCommentWithDefaults instantiates a new ThreadedComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadedCommentWithDefaults() *ThreadedComment {
	this := ThreadedComment{}
	return &this
}

// GetAppUrl returns the AppUrl field value
func (o *ThreadedComment) GetAppUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppUrl
}

// GetAppUrlOk returns a tuple with the AppUrl field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetAppUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUrl, true
}

// SetAppUrl sets field value
func (o *ThreadedComment) SetAppUrl(v string) {
	o.AppUrl = v
}

// GetEntityType returns the EntityType field value
func (o *ThreadedComment) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *ThreadedComment) SetEntityType(v string) {
	o.EntityType = v
}

// GetDeleted returns the Deleted field value
func (o *ThreadedComment) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *ThreadedComment) SetDeleted(v bool) {
	o.Deleted = v
}

// GetMentionIds returns the MentionIds field value
func (o *ThreadedComment) GetMentionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MentionIds
}

// GetMentionIdsOk returns a tuple with the MentionIds field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetMentionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionIds, true
}

// SetMentionIds sets field value
func (o *ThreadedComment) SetMentionIds(v []string) {
	o.MentionIds = v
}

// GetAuthorId returns the AuthorId field value
func (o *ThreadedComment) GetAuthorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorId, true
}

// SetAuthorId sets field value
func (o *ThreadedComment) SetAuthorId(v string) {
	o.AuthorId = v
}

// GetMemberMentionIds returns the MemberMentionIds field value
func (o *ThreadedComment) GetMemberMentionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MemberMentionIds
}

// GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetMemberMentionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberMentionIds, true
}

// SetMemberMentionIds sets field value
func (o *ThreadedComment) SetMemberMentionIds(v []string) {
	o.MemberMentionIds = v
}

// GetComments returns the Comments field value
func (o *ThreadedComment) GetComments() []ThreadedComment {
	if o == nil {
		var ret []ThreadedComment
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetCommentsOk() ([]ThreadedComment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments, true
}

// SetComments sets field value
func (o *ThreadedComment) SetComments(v []ThreadedComment) {
	o.Comments = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ThreadedComment) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ThreadedComment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetGroupMentionIds returns the GroupMentionIds field value
func (o *ThreadedComment) GetGroupMentionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupMentionIds
}

// GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetGroupMentionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupMentionIds, true
}

// SetGroupMentionIds sets field value
func (o *ThreadedComment) SetGroupMentionIds(v []string) {
	o.GroupMentionIds = v
}

// GetExternalId returns the ExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ThreadedComment) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadedComment) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// SetExternalId sets field value
func (o *ThreadedComment) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// GetId returns the Id field value
func (o *ThreadedComment) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThreadedComment) SetId(v int64) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ThreadedComment) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ThreadedComment) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetText returns the Text field value
func (o *ThreadedComment) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *ThreadedComment) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *ThreadedComment) SetText(v string) {
	o.Text = v
}

func (o ThreadedComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadedComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_url"] = o.AppUrl
	toSerialize["entity_type"] = o.EntityType
	toSerialize["deleted"] = o.Deleted
	toSerialize["mention_ids"] = o.MentionIds
	toSerialize["author_id"] = o.AuthorId
	toSerialize["member_mention_ids"] = o.MemberMentionIds
	toSerialize["comments"] = o.Comments
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["group_mention_ids"] = o.GroupMentionIds
	toSerialize["external_id"] = o.ExternalId.Get()
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

type NullableThreadedComment struct {
	value *ThreadedComment
	isSet bool
}

func (v NullableThreadedComment) Get() *ThreadedComment {
	return v.value
}

func (v *NullableThreadedComment) Set(val *ThreadedComment) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadedComment) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadedComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadedComment(val *ThreadedComment) *NullableThreadedComment {
	return &NullableThreadedComment{value: val, isSet: true}
}

func (v NullableThreadedComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadedComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


