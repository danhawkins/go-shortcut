# UpdateLinkedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the file. | [optional] 
**StoryId** | Pointer to **int64** | The ID of the linked story. | [optional] 
**Name** | Pointer to **string** | The name of the file. | [optional] 
**ThumbnailUrl** | Pointer to **string** | The URL of the thumbnail, if the integration provided it. | [optional] 
**Type** | Pointer to **string** | The integration type of the file (e.g. google, dropbox, box). | [optional] 
**Size** | Pointer to **int64** | The filesize, if the integration provided it. | [optional] 
**UploaderId** | Pointer to **string** | The UUID of the member that uploaded the file. | [optional] 
**Url** | Pointer to **string** | The URL of linked file. | [optional] 

## Methods

### NewUpdateLinkedFile

`func NewUpdateLinkedFile() *UpdateLinkedFile`

NewUpdateLinkedFile instantiates a new UpdateLinkedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLinkedFileWithDefaults

`func NewUpdateLinkedFileWithDefaults() *UpdateLinkedFile`

NewUpdateLinkedFileWithDefaults instantiates a new UpdateLinkedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateLinkedFile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLinkedFile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLinkedFile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLinkedFile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStoryId

`func (o *UpdateLinkedFile) GetStoryId() int64`

GetStoryId returns the StoryId field if non-nil, zero value otherwise.

### GetStoryIdOk

`func (o *UpdateLinkedFile) GetStoryIdOk() (*int64, bool)`

GetStoryIdOk returns a tuple with the StoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryId

`func (o *UpdateLinkedFile) SetStoryId(v int64)`

SetStoryId sets StoryId field to given value.

### HasStoryId

`func (o *UpdateLinkedFile) HasStoryId() bool`

HasStoryId returns a boolean if a field has been set.

### GetName

`func (o *UpdateLinkedFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLinkedFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLinkedFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLinkedFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *UpdateLinkedFile) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *UpdateLinkedFile) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *UpdateLinkedFile) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *UpdateLinkedFile) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetType

`func (o *UpdateLinkedFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateLinkedFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateLinkedFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateLinkedFile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *UpdateLinkedFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UpdateLinkedFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UpdateLinkedFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *UpdateLinkedFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUploaderId

`func (o *UpdateLinkedFile) GetUploaderId() string`

GetUploaderId returns the UploaderId field if non-nil, zero value otherwise.

### GetUploaderIdOk

`func (o *UpdateLinkedFile) GetUploaderIdOk() (*string, bool)`

GetUploaderIdOk returns a tuple with the UploaderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderId

`func (o *UpdateLinkedFile) SetUploaderId(v string)`

SetUploaderId sets UploaderId field to given value.

### HasUploaderId

`func (o *UpdateLinkedFile) HasUploaderId() bool`

HasUploaderId returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateLinkedFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateLinkedFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateLinkedFile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateLinkedFile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


