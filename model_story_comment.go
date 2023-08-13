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

// checks if the StoryComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoryComment{}

// StoryComment A Comment is any note added within the Comment field of a Story.
type StoryComment struct {
	// The Shortcut application url for the Comment.
	AppUrl string `json:"app_url"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// True/false boolean indicating whether the Comment has been deleted.
	Deleted bool `json:"deleted"`
	// The ID of the Story on which the Comment appears.
	StoryId int64 `json:"story_id"`
	// Deprecated: use member_mention_ids.
	MentionIds []string `json:"mention_ids"`
	// The unique ID of the Member who is the Comment's author.
	AuthorId NullableString `json:"author_id"`
	// The unique IDs of the Member who are mentioned in the Comment.
	MemberMentionIds []string `json:"member_mention_ids"`
	// Marks the comment as a blocker that can be surfaced to permissions or teams mentioned in the comment. Can only be used on a top-level comment.
	Blocker *bool `json:"blocker,omitempty"`
	// The time/date when the Comment was updated.
	UpdatedAt NullableTime `json:"updated_at"`
	// The unique IDs of the Group who are mentioned in the Comment.
	GroupMentionIds []string `json:"group_mention_ids"`
	// This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId NullableString `json:"external_id"`
	// The ID of the parent Comment this Comment is threaded under.
	ParentId NullableInt64 `json:"parent_id,omitempty"`
	// The unique ID of the Comment.
	Id int64 `json:"id"`
	// The Comments numerical position in the list from oldest to newest.
	Position int64 `json:"position"`
	// Marks the comment as an unblocker to its  blocker parent. Can only be set on a threaded comment who has a parent with `blocker` set.
	UnblocksParent *bool `json:"unblocks_parent,omitempty"`
	// A set of Reactions to this Comment.
	Reactions []StoryReaction `json:"reactions"`
	// The time/date when the Comment was created.
	CreatedAt time.Time `json:"created_at"`
	// The text of the Comment. In the case that the Comment has been deleted, this field can be set to nil.
	Text NullableString `json:"text"`
}

// NewStoryComment instantiates a new StoryComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoryComment(appUrl string, entityType string, deleted bool, storyId int64, mentionIds []string, authorId NullableString, memberMentionIds []string, updatedAt NullableTime, groupMentionIds []string, externalId NullableString, id int64, position int64, reactions []StoryReaction, createdAt time.Time, text NullableString) *StoryComment {
	this := StoryComment{}
	this.AppUrl = appUrl
	this.EntityType = entityType
	this.Deleted = deleted
	this.StoryId = storyId
	this.MentionIds = mentionIds
	this.AuthorId = authorId
	this.MemberMentionIds = memberMentionIds
	this.UpdatedAt = updatedAt
	this.GroupMentionIds = groupMentionIds
	this.ExternalId = externalId
	this.Id = id
	this.Position = position
	this.Reactions = reactions
	this.CreatedAt = createdAt
	this.Text = text
	return &this
}

// NewStoryCommentWithDefaults instantiates a new StoryComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoryCommentWithDefaults() *StoryComment {
	this := StoryComment{}
	return &this
}

// GetAppUrl returns the AppUrl field value
func (o *StoryComment) GetAppUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppUrl
}

// GetAppUrlOk returns a tuple with the AppUrl field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetAppUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUrl, true
}

// SetAppUrl sets field value
func (o *StoryComment) SetAppUrl(v string) {
	o.AppUrl = v
}

// GetEntityType returns the EntityType field value
func (o *StoryComment) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *StoryComment) SetEntityType(v string) {
	o.EntityType = v
}

// GetDeleted returns the Deleted field value
func (o *StoryComment) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *StoryComment) SetDeleted(v bool) {
	o.Deleted = v
}

// GetStoryId returns the StoryId field value
func (o *StoryComment) GetStoryId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StoryId
}

// GetStoryIdOk returns a tuple with the StoryId field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetStoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoryId, true
}

// SetStoryId sets field value
func (o *StoryComment) SetStoryId(v int64) {
	o.StoryId = v
}

// GetMentionIds returns the MentionIds field value
func (o *StoryComment) GetMentionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MentionIds
}

// GetMentionIdsOk returns a tuple with the MentionIds field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetMentionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionIds, true
}

// SetMentionIds sets field value
func (o *StoryComment) SetMentionIds(v []string) {
	o.MentionIds = v
}

// GetAuthorId returns the AuthorId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StoryComment) GetAuthorId() string {
	if o == nil || o.AuthorId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthorId.Get()
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoryComment) GetAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorId.Get(), o.AuthorId.IsSet()
}

// SetAuthorId sets field value
func (o *StoryComment) SetAuthorId(v string) {
	o.AuthorId.Set(&v)
}

// GetMemberMentionIds returns the MemberMentionIds field value
func (o *StoryComment) GetMemberMentionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MemberMentionIds
}

// GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetMemberMentionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberMentionIds, true
}

// SetMemberMentionIds sets field value
func (o *StoryComment) SetMemberMentionIds(v []string) {
	o.MemberMentionIds = v
}

// GetBlocker returns the Blocker field value if set, zero value otherwise.
func (o *StoryComment) GetBlocker() bool {
	if o == nil || IsNil(o.Blocker) {
		var ret bool
		return ret
	}
	return *o.Blocker
}

// GetBlockerOk returns a tuple with the Blocker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryComment) GetBlockerOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocker) {
		return nil, false
	}
	return o.Blocker, true
}

// HasBlocker returns a boolean if a field has been set.
func (o *StoryComment) HasBlocker() bool {
	if o != nil && !IsNil(o.Blocker) {
		return true
	}

	return false
}

// SetBlocker gets a reference to the given bool and assigns it to the Blocker field.
func (o *StoryComment) SetBlocker(v bool) {
	o.Blocker = &v
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *StoryComment) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoryComment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *StoryComment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// GetGroupMentionIds returns the GroupMentionIds field value
func (o *StoryComment) GetGroupMentionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupMentionIds
}

// GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetGroupMentionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupMentionIds, true
}

// SetGroupMentionIds sets field value
func (o *StoryComment) SetGroupMentionIds(v []string) {
	o.GroupMentionIds = v
}

// GetExternalId returns the ExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StoryComment) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoryComment) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// SetExternalId sets field value
func (o *StoryComment) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoryComment) GetParentId() int64 {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret int64
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoryComment) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *StoryComment) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableInt64 and assigns it to the ParentId field.
func (o *StoryComment) SetParentId(v int64) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *StoryComment) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *StoryComment) UnsetParentId() {
	o.ParentId.Unset()
}

// GetId returns the Id field value
func (o *StoryComment) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StoryComment) SetId(v int64) {
	o.Id = v
}

// GetPosition returns the Position field value
func (o *StoryComment) GetPosition() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetPositionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *StoryComment) SetPosition(v int64) {
	o.Position = v
}

// GetUnblocksParent returns the UnblocksParent field value if set, zero value otherwise.
func (o *StoryComment) GetUnblocksParent() bool {
	if o == nil || IsNil(o.UnblocksParent) {
		var ret bool
		return ret
	}
	return *o.UnblocksParent
}

// GetUnblocksParentOk returns a tuple with the UnblocksParent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoryComment) GetUnblocksParentOk() (*bool, bool) {
	if o == nil || IsNil(o.UnblocksParent) {
		return nil, false
	}
	return o.UnblocksParent, true
}

// HasUnblocksParent returns a boolean if a field has been set.
func (o *StoryComment) HasUnblocksParent() bool {
	if o != nil && !IsNil(o.UnblocksParent) {
		return true
	}

	return false
}

// SetUnblocksParent gets a reference to the given bool and assigns it to the UnblocksParent field.
func (o *StoryComment) SetUnblocksParent(v bool) {
	o.UnblocksParent = &v
}

// GetReactions returns the Reactions field value
func (o *StoryComment) GetReactions() []StoryReaction {
	if o == nil {
		var ret []StoryReaction
		return ret
	}

	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetReactionsOk() ([]StoryReaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reactions, true
}

// SetReactions sets field value
func (o *StoryComment) SetReactions(v []StoryReaction) {
	o.Reactions = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *StoryComment) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *StoryComment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *StoryComment) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetText returns the Text field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StoryComment) GetText() string {
	if o == nil || o.Text.Get() == nil {
		var ret string
		return ret
	}

	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoryComment) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// SetText sets field value
func (o *StoryComment) SetText(v string) {
	o.Text.Set(&v)
}

func (o StoryComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoryComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_url"] = o.AppUrl
	toSerialize["entity_type"] = o.EntityType
	toSerialize["deleted"] = o.Deleted
	toSerialize["story_id"] = o.StoryId
	toSerialize["mention_ids"] = o.MentionIds
	toSerialize["author_id"] = o.AuthorId.Get()
	toSerialize["member_mention_ids"] = o.MemberMentionIds
	if !IsNil(o.Blocker) {
		toSerialize["blocker"] = o.Blocker
	}
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	toSerialize["group_mention_ids"] = o.GroupMentionIds
	toSerialize["external_id"] = o.ExternalId.Get()
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["position"] = o.Position
	if !IsNil(o.UnblocksParent) {
		toSerialize["unblocks_parent"] = o.UnblocksParent
	}
	toSerialize["reactions"] = o.Reactions
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["text"] = o.Text.Get()
	return toSerialize, nil
}

type NullableStoryComment struct {
	value *StoryComment
	isSet bool
}

func (v NullableStoryComment) Get() *StoryComment {
	return v.value
}

func (v *NullableStoryComment) Set(val *StoryComment) {
	v.value = val
	v.isSet = true
}

func (v NullableStoryComment) IsSet() bool {
	return v.isSet
}

func (v *NullableStoryComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoryComment(val *StoryComment) *NullableStoryComment {
	return &NullableStoryComment{value: val, isSet: true}
}

func (v NullableStoryComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoryComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

