# StoryComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Comment. | 
**EntityType** | **string** | A string description of this resource. | 
**Deleted** | **bool** | True/false boolean indicating whether the Comment has been deleted. | 
**StoryId** | **int64** | The ID of the Story on which the Comment appears. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**AuthorId** | **NullableString** | The unique ID of the Member who is the Comment&#39;s author. | 
**MemberMentionIds** | **[]string** | The unique IDs of the Member who are mentioned in the Comment. | 
**Blocker** | Pointer to **bool** | Marks the comment as a blocker that can be surfaced to permissions or teams mentioned in the comment. Can only be used on a top-level comment. | [optional] 
**UpdatedAt** | **NullableTime** | The time/date when the Comment was updated. | 
**GroupMentionIds** | **[]string** | The unique IDs of the Group who are mentioned in the Comment. | 
**ExternalId** | **NullableString** | This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here. | 
**ParentId** | Pointer to **NullableInt64** | The ID of the parent Comment this Comment is threaded under. | [optional] 
**Id** | **int64** | The unique ID of the Comment. | 
**Position** | **int64** | The Comments numerical position in the list from oldest to newest. | 
**UnblocksParent** | Pointer to **bool** | Marks the comment as an unblocker to its  blocker parent. Can only be set on a threaded comment who has a parent with &#x60;blocker&#x60; set. | [optional] 
**Reactions** | [**[]StoryReaction**](StoryReaction.md) | A set of Reactions to this Comment. | 
**CreatedAt** | **time.Time** | The time/date when the Comment was created. | 
**Text** | **NullableString** | The text of the Comment. In the case that the Comment has been deleted, this field can be set to nil. | 

## Methods

### NewStoryComment

`func NewStoryComment(appUrl string, entityType string, deleted bool, storyId int64, mentionIds []string, authorId NullableString, memberMentionIds []string, updatedAt NullableTime, groupMentionIds []string, externalId NullableString, id int64, position int64, reactions []StoryReaction, createdAt time.Time, text NullableString, ) *StoryComment`

NewStoryComment instantiates a new StoryComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryCommentWithDefaults

`func NewStoryCommentWithDefaults() *StoryComment`

NewStoryCommentWithDefaults instantiates a new StoryComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *StoryComment) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *StoryComment) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *StoryComment) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetEntityType

`func (o *StoryComment) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *StoryComment) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *StoryComment) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetDeleted

`func (o *StoryComment) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *StoryComment) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *StoryComment) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetStoryId

`func (o *StoryComment) GetStoryId() int64`

GetStoryId returns the StoryId field if non-nil, zero value otherwise.

### GetStoryIdOk

`func (o *StoryComment) GetStoryIdOk() (*int64, bool)`

GetStoryIdOk returns a tuple with the StoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryId

`func (o *StoryComment) SetStoryId(v int64)`

SetStoryId sets StoryId field to given value.


### GetMentionIds

`func (o *StoryComment) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *StoryComment) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *StoryComment) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetAuthorId

`func (o *StoryComment) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *StoryComment) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *StoryComment) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### SetAuthorIdNil

`func (o *StoryComment) SetAuthorIdNil(b bool)`

 SetAuthorIdNil sets the value for AuthorId to be an explicit nil

### UnsetAuthorId
`func (o *StoryComment) UnsetAuthorId()`

UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil
### GetMemberMentionIds

`func (o *StoryComment) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *StoryComment) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *StoryComment) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetBlocker

`func (o *StoryComment) GetBlocker() bool`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *StoryComment) GetBlockerOk() (*bool, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *StoryComment) SetBlocker(v bool)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *StoryComment) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StoryComment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StoryComment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StoryComment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *StoryComment) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *StoryComment) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetGroupMentionIds

`func (o *StoryComment) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *StoryComment) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *StoryComment) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetExternalId

`func (o *StoryComment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *StoryComment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *StoryComment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *StoryComment) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *StoryComment) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetParentId

`func (o *StoryComment) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *StoryComment) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *StoryComment) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *StoryComment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *StoryComment) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *StoryComment) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetId

`func (o *StoryComment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoryComment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoryComment) SetId(v int64)`

SetId sets Id field to given value.


### GetPosition

`func (o *StoryComment) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *StoryComment) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *StoryComment) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetUnblocksParent

`func (o *StoryComment) GetUnblocksParent() bool`

GetUnblocksParent returns the UnblocksParent field if non-nil, zero value otherwise.

### GetUnblocksParentOk

`func (o *StoryComment) GetUnblocksParentOk() (*bool, bool)`

GetUnblocksParentOk returns a tuple with the UnblocksParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnblocksParent

`func (o *StoryComment) SetUnblocksParent(v bool)`

SetUnblocksParent sets UnblocksParent field to given value.

### HasUnblocksParent

`func (o *StoryComment) HasUnblocksParent() bool`

HasUnblocksParent returns a boolean if a field has been set.

### GetReactions

`func (o *StoryComment) GetReactions() []StoryReaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *StoryComment) GetReactionsOk() (*[]StoryReaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *StoryComment) SetReactions(v []StoryReaction)`

SetReactions sets Reactions field to given value.


### GetCreatedAt

`func (o *StoryComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StoryComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StoryComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetText

`func (o *StoryComment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StoryComment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StoryComment) SetText(v string)`

SetText sets Text field to given value.


### SetTextNil

`func (o *StoryComment) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *StoryComment) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


