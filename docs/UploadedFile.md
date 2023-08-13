# UploadedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** | The description of the file. | 
**EntityType** | **string** | A string description of this resource. | 
**StoryIds** | **[]int64** | The unique IDs of the Stories associated with this file. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**MemberMentionIds** | **[]string** | The unique IDs of the Members who are mentioned in the file description. | 
**Name** | **string** | The optional User-specified name of the file. | 
**ThumbnailUrl** | **NullableString** | The url where the thumbnail of the file can be found in Shortcut. | 
**Size** | **int64** | The size of the file. | 
**UploaderId** | **string** | The unique ID of the Member who uploaded the file. | 
**ContentType** | **string** | Free form string corresponding to a text or image file. | 
**UpdatedAt** | **NullableTime** | The time/date that the file was updated. | 
**Filename** | **string** | The name assigned to the file in Shortcut upon upload. | 
**GroupMentionIds** | **[]string** | The unique IDs of the Groups who are mentioned in the file description. | 
**ExternalId** | **NullableString** | This field can be set to another unique ID. In the case that the File has been imported from another tool, the ID in the other tool can be indicated here. | 
**Id** | **int64** | The unique ID for the file. | 
**Url** | **NullableString** | The URL for the file. | 
**CreatedAt** | **time.Time** | The time/date that the file was created. | 

## Methods

### NewUploadedFile

`func NewUploadedFile(description NullableString, entityType string, storyIds []int64, mentionIds []string, memberMentionIds []string, name string, thumbnailUrl NullableString, size int64, uploaderId string, contentType string, updatedAt NullableTime, filename string, groupMentionIds []string, externalId NullableString, id int64, url NullableString, createdAt time.Time, ) *UploadedFile`

NewUploadedFile instantiates a new UploadedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadedFileWithDefaults

`func NewUploadedFileWithDefaults() *UploadedFile`

NewUploadedFileWithDefaults instantiates a new UploadedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UploadedFile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UploadedFile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UploadedFile) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *UploadedFile) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UploadedFile) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEntityType

`func (o *UploadedFile) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *UploadedFile) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *UploadedFile) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetStoryIds

`func (o *UploadedFile) GetStoryIds() []int64`

GetStoryIds returns the StoryIds field if non-nil, zero value otherwise.

### GetStoryIdsOk

`func (o *UploadedFile) GetStoryIdsOk() (*[]int64, bool)`

GetStoryIdsOk returns a tuple with the StoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryIds

`func (o *UploadedFile) SetStoryIds(v []int64)`

SetStoryIds sets StoryIds field to given value.


### GetMentionIds

`func (o *UploadedFile) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *UploadedFile) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *UploadedFile) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetMemberMentionIds

`func (o *UploadedFile) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *UploadedFile) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *UploadedFile) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetName

`func (o *UploadedFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadedFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadedFile) SetName(v string)`

SetName sets Name field to given value.


### GetThumbnailUrl

`func (o *UploadedFile) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *UploadedFile) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *UploadedFile) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### SetThumbnailUrlNil

`func (o *UploadedFile) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *UploadedFile) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetSize

`func (o *UploadedFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UploadedFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UploadedFile) SetSize(v int64)`

SetSize sets Size field to given value.


### GetUploaderId

`func (o *UploadedFile) GetUploaderId() string`

GetUploaderId returns the UploaderId field if non-nil, zero value otherwise.

### GetUploaderIdOk

`func (o *UploadedFile) GetUploaderIdOk() (*string, bool)`

GetUploaderIdOk returns a tuple with the UploaderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderId

`func (o *UploadedFile) SetUploaderId(v string)`

SetUploaderId sets UploaderId field to given value.


### GetContentType

`func (o *UploadedFile) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *UploadedFile) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *UploadedFile) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetUpdatedAt

`func (o *UploadedFile) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UploadedFile) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UploadedFile) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *UploadedFile) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *UploadedFile) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetFilename

`func (o *UploadedFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UploadedFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UploadedFile) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetGroupMentionIds

`func (o *UploadedFile) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *UploadedFile) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *UploadedFile) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetExternalId

`func (o *UploadedFile) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UploadedFile) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UploadedFile) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *UploadedFile) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UploadedFile) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetId

`func (o *UploadedFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadedFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadedFile) SetId(v int64)`

SetId sets Id field to given value.


### GetUrl

`func (o *UploadedFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadedFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadedFile) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *UploadedFile) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *UploadedFile) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCreatedAt

`func (o *UploadedFile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UploadedFile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UploadedFile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


