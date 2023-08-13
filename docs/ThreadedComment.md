# ThreadedComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Comment. | 
**EntityType** | **string** | A string description of this resource. | 
**Deleted** | **bool** | True/false boolean indicating whether the Comment is deleted. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**AuthorId** | **string** | The unique ID of the Member that authored the Comment. | 
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in this Comment. | 
**Comments** | [**[]ThreadedComment**](ThreadedComment.md) | A nested array of threaded comments. | 
**UpdatedAt** | **time.Time** | The time/date the Comment was updated. | 
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in this Comment. | 
**ExternalId** | **NullableString** | This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here. | 
**Id** | **int64** | The unique ID of the Comment. | 
**CreatedAt** | **time.Time** | The time/date the Comment was created. | 
**Text** | **string** | The text of the Comment. | 

## Methods

### NewThreadedComment

`func NewThreadedComment(appUrl string, entityType string, deleted bool, mentionIds []string, authorId string, memberMentionIds []string, comments []ThreadedComment, updatedAt time.Time, groupMentionIds []string, externalId NullableString, id int64, createdAt time.Time, text string, ) *ThreadedComment`

NewThreadedComment instantiates a new ThreadedComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadedCommentWithDefaults

`func NewThreadedCommentWithDefaults() *ThreadedComment`

NewThreadedCommentWithDefaults instantiates a new ThreadedComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *ThreadedComment) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *ThreadedComment) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *ThreadedComment) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetEntityType

`func (o *ThreadedComment) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ThreadedComment) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ThreadedComment) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetDeleted

`func (o *ThreadedComment) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ThreadedComment) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ThreadedComment) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetMentionIds

`func (o *ThreadedComment) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *ThreadedComment) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *ThreadedComment) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetAuthorId

`func (o *ThreadedComment) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *ThreadedComment) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *ThreadedComment) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetMemberMentionIds

`func (o *ThreadedComment) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *ThreadedComment) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *ThreadedComment) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetComments

`func (o *ThreadedComment) GetComments() []ThreadedComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ThreadedComment) GetCommentsOk() (*[]ThreadedComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ThreadedComment) SetComments(v []ThreadedComment)`

SetComments sets Comments field to given value.


### GetUpdatedAt

`func (o *ThreadedComment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ThreadedComment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ThreadedComment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetGroupMentionIds

`func (o *ThreadedComment) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *ThreadedComment) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *ThreadedComment) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetExternalId

`func (o *ThreadedComment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ThreadedComment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ThreadedComment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *ThreadedComment) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ThreadedComment) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetId

`func (o *ThreadedComment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreadedComment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreadedComment) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ThreadedComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ThreadedComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ThreadedComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetText

`func (o *ThreadedComment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ThreadedComment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ThreadedComment) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


