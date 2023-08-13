# CreateStoryComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorId** | Pointer to **string** | The Member ID of the Comment&#39;s author. Defaults to the user identified by the API token. | [optional] 
**Blocker** | Pointer to **bool** | Marks the comment as a blocker that can be surfaced to permissions or teams mentioned in the comment. Can only be used on a top-level comment. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Defaults to the time/date the comment is last updated, but can be set to reflect another date. | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the comment has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 
**ParentId** | Pointer to **NullableInt64** | The ID of the Comment that this comment is threaded under. | [optional] 
**UnblocksParent** | Pointer to **bool** | Marks the comment as an unblocker to its  blocker parent. Can only be set on a threaded comment who has a parent with &#x60;blocker&#x60; set. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Defaults to the time/date the comment is created, but can be set to reflect another date. | [optional] 
**Text** | **string** | The comment text. | 

## Methods

### NewCreateStoryComment

`func NewCreateStoryComment(text string, ) *CreateStoryComment`

NewCreateStoryComment instantiates a new CreateStoryComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoryCommentWithDefaults

`func NewCreateStoryCommentWithDefaults() *CreateStoryComment`

NewCreateStoryCommentWithDefaults instantiates a new CreateStoryComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorId

`func (o *CreateStoryComment) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CreateStoryComment) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CreateStoryComment) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CreateStoryComment) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetBlocker

`func (o *CreateStoryComment) GetBlocker() bool`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *CreateStoryComment) GetBlockerOk() (*bool, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *CreateStoryComment) SetBlocker(v bool)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *CreateStoryComment) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateStoryComment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateStoryComment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateStoryComment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateStoryComment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateStoryComment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateStoryComment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateStoryComment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateStoryComment) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetParentId

`func (o *CreateStoryComment) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateStoryComment) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateStoryComment) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateStoryComment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *CreateStoryComment) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *CreateStoryComment) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetUnblocksParent

`func (o *CreateStoryComment) GetUnblocksParent() bool`

GetUnblocksParent returns the UnblocksParent field if non-nil, zero value otherwise.

### GetUnblocksParentOk

`func (o *CreateStoryComment) GetUnblocksParentOk() (*bool, bool)`

GetUnblocksParentOk returns a tuple with the UnblocksParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnblocksParent

`func (o *CreateStoryComment) SetUnblocksParent(v bool)`

SetUnblocksParent sets UnblocksParent field to given value.

### HasUnblocksParent

`func (o *CreateStoryComment) HasUnblocksParent() bool`

HasUnblocksParent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateStoryComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateStoryComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateStoryComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateStoryComment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetText

`func (o *CreateStoryComment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateStoryComment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateStoryComment) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


