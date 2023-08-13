# StoryReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emoji** | **string** | Emoji text of the reaction. | 
**PermissionIds** | **[]string** | Permissions who have reacted with this. | 

## Methods

### NewStoryReaction

`func NewStoryReaction(emoji string, permissionIds []string, ) *StoryReaction`

NewStoryReaction instantiates a new StoryReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryReactionWithDefaults

`func NewStoryReactionWithDefaults() *StoryReaction`

NewStoryReactionWithDefaults instantiates a new StoryReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmoji

`func (o *StoryReaction) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *StoryReaction) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *StoryReaction) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.


### GetPermissionIds

`func (o *StoryReaction) GetPermissionIds() []string`

GetPermissionIds returns the PermissionIds field if non-nil, zero value otherwise.

### GetPermissionIdsOk

`func (o *StoryReaction) GetPermissionIdsOk() (*[]string, bool)`

GetPermissionIdsOk returns a tuple with the PermissionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionIds

`func (o *StoryReaction) SetPermissionIds(v []string)`

SetPermissionIds sets PermissionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


