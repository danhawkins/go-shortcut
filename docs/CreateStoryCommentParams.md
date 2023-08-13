# CreateStoryCommentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The comment text. | 
**AuthorId** | Pointer to **string** | The Member ID of the Comment&#39;s author. Defaults to the user identified by the API token. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Defaults to the time/date the comment is created, but can be set to reflect another date. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Defaults to the time/date the comment is last updated, but can be set to reflect another date. | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the comment has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 
**ParentId** | Pointer to **NullableInt64** | The ID of the Comment that this comment is threaded under. | [optional] 
**Blocker** | Pointer to **bool** | Marks the comment as a blocker that can be surfaced to permissions or teams mentioned in the comment. Can only be used on a top-level comment. | [optional] 
**UnblocksParent** | Pointer to **bool** | Marks the comment as an unblocker to its  blocker parent. Can only be set on a threaded comment who has a parent with &#x60;blocker&#x60; set. | [optional] 

## Methods

### NewCreateStoryCommentParams

`func NewCreateStoryCommentParams(text string, ) *CreateStoryCommentParams`

NewCreateStoryCommentParams instantiates a new CreateStoryCommentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoryCommentParamsWithDefaults

`func NewCreateStoryCommentParamsWithDefaults() *CreateStoryCommentParams`

NewCreateStoryCommentParamsWithDefaults instantiates a new CreateStoryCommentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *CreateStoryCommentParams) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateStoryCommentParams) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateStoryCommentParams) SetText(v string)`

SetText sets Text field to given value.


### GetAuthorId

`func (o *CreateStoryCommentParams) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CreateStoryCommentParams) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CreateStoryCommentParams) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CreateStoryCommentParams) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateStoryCommentParams) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateStoryCommentParams) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateStoryCommentParams) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateStoryCommentParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateStoryCommentParams) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateStoryCommentParams) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateStoryCommentParams) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateStoryCommentParams) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateStoryCommentParams) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateStoryCommentParams) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateStoryCommentParams) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateStoryCommentParams) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetParentId

`func (o *CreateStoryCommentParams) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateStoryCommentParams) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateStoryCommentParams) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateStoryCommentParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *CreateStoryCommentParams) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *CreateStoryCommentParams) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetBlocker

`func (o *CreateStoryCommentParams) GetBlocker() bool`

GetBlocker returns the Blocker field if non-nil, zero value otherwise.

### GetBlockerOk

`func (o *CreateStoryCommentParams) GetBlockerOk() (*bool, bool)`

GetBlockerOk returns a tuple with the Blocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocker

`func (o *CreateStoryCommentParams) SetBlocker(v bool)`

SetBlocker sets Blocker field to given value.

### HasBlocker

`func (o *CreateStoryCommentParams) HasBlocker() bool`

HasBlocker returns a boolean if a field has been set.

### GetUnblocksParent

`func (o *CreateStoryCommentParams) GetUnblocksParent() bool`

GetUnblocksParent returns the UnblocksParent field if non-nil, zero value otherwise.

### GetUnblocksParentOk

`func (o *CreateStoryCommentParams) GetUnblocksParentOk() (*bool, bool)`

GetUnblocksParentOk returns a tuple with the UnblocksParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnblocksParent

`func (o *CreateStoryCommentParams) SetUnblocksParent(v bool)`

SetUnblocksParent sets UnblocksParent field to given value.

### HasUnblocksParent

`func (o *CreateStoryCommentParams) HasUnblocksParent() bool`

HasUnblocksParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


