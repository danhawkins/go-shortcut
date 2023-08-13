# LinkedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** | The description of the file. | 
**EntityType** | **string** | A string description of this resource. | 
**StoryIds** | **[]int64** | The IDs of the stories this file is attached to. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**MemberMentionIds** | **[]string** | The members that are mentioned in the description of the file. | 
**Name** | **string** | The name of the linked file. | 
**ThumbnailUrl** | **NullableString** | The URL of the file thumbnail, if the integration provided it. | 
**Type** | **string** | The integration type (e.g. google, dropbox, box). | 
**Size** | **NullableInt64** | The filesize, if the integration provided it. | 
**UploaderId** | **string** | The UUID of the member that uploaded the file. | 
**ContentType** | **NullableString** | The content type of the image (e.g. txt/plain). | 
**UpdatedAt** | **time.Time** | The time/date the LinkedFile was updated. | 
**GroupMentionIds** | **[]string** | The groups that are mentioned in the description of the file. | 
**Id** | **int64** | The unique identifier for the file. | 
**Url** | **string** | The URL of the file. | 
**CreatedAt** | **time.Time** | The time/date the LinkedFile was created. | 

## Methods

### NewLinkedFile

`func NewLinkedFile(description NullableString, entityType string, storyIds []int64, mentionIds []string, memberMentionIds []string, name string, thumbnailUrl NullableString, type_ string, size NullableInt64, uploaderId string, contentType NullableString, updatedAt time.Time, groupMentionIds []string, id int64, url string, createdAt time.Time, ) *LinkedFile`

NewLinkedFile instantiates a new LinkedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedFileWithDefaults

`func NewLinkedFileWithDefaults() *LinkedFile`

NewLinkedFileWithDefaults instantiates a new LinkedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LinkedFile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinkedFile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinkedFile) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *LinkedFile) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LinkedFile) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEntityType

`func (o *LinkedFile) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *LinkedFile) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *LinkedFile) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetStoryIds

`func (o *LinkedFile) GetStoryIds() []int64`

GetStoryIds returns the StoryIds field if non-nil, zero value otherwise.

### GetStoryIdsOk

`func (o *LinkedFile) GetStoryIdsOk() (*[]int64, bool)`

GetStoryIdsOk returns a tuple with the StoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryIds

`func (o *LinkedFile) SetStoryIds(v []int64)`

SetStoryIds sets StoryIds field to given value.


### GetMentionIds

`func (o *LinkedFile) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *LinkedFile) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *LinkedFile) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetMemberMentionIds

`func (o *LinkedFile) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *LinkedFile) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *LinkedFile) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetName

`func (o *LinkedFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkedFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkedFile) SetName(v string)`

SetName sets Name field to given value.


### GetThumbnailUrl

`func (o *LinkedFile) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *LinkedFile) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *LinkedFile) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### SetThumbnailUrlNil

`func (o *LinkedFile) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *LinkedFile) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetType

`func (o *LinkedFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedFile) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *LinkedFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LinkedFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LinkedFile) SetSize(v int64)`

SetSize sets Size field to given value.


### SetSizeNil

`func (o *LinkedFile) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *LinkedFile) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetUploaderId

`func (o *LinkedFile) GetUploaderId() string`

GetUploaderId returns the UploaderId field if non-nil, zero value otherwise.

### GetUploaderIdOk

`func (o *LinkedFile) GetUploaderIdOk() (*string, bool)`

GetUploaderIdOk returns a tuple with the UploaderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderId

`func (o *LinkedFile) SetUploaderId(v string)`

SetUploaderId sets UploaderId field to given value.


### GetContentType

`func (o *LinkedFile) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LinkedFile) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LinkedFile) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### SetContentTypeNil

`func (o *LinkedFile) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *LinkedFile) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetUpdatedAt

`func (o *LinkedFile) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LinkedFile) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LinkedFile) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetGroupMentionIds

`func (o *LinkedFile) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *LinkedFile) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *LinkedFile) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetId

`func (o *LinkedFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedFile) SetId(v int64)`

SetId sets Id field to given value.


### GetUrl

`func (o *LinkedFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LinkedFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LinkedFile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCreatedAt

`func (o *LinkedFile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LinkedFile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LinkedFile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


