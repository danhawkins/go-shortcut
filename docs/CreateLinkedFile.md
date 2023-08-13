# CreateLinkedFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the file. | [optional] 
**StoryId** | Pointer to **int64** | The ID of the linked story. | [optional] 
**Name** | **string** | The name of the file. | 
**ThumbnailUrl** | Pointer to **string** | The URL of the thumbnail, if the integration provided it. | [optional] 
**Type** | **string** | The integration type of the file (e.g. google, dropbox, box). | 
**Size** | Pointer to **int64** | The filesize, if the integration provided it. | [optional] 
**UploaderId** | Pointer to **string** | The UUID of the member that uploaded the file. | [optional] 
**ContentType** | Pointer to **string** | The content type of the image (e.g. txt/plain). | [optional] 
**Url** | **string** | The URL of linked file. | 

## Methods

### NewCreateLinkedFile

`func NewCreateLinkedFile(name string, type_ string, url string, ) *CreateLinkedFile`

NewCreateLinkedFile instantiates a new CreateLinkedFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLinkedFileWithDefaults

`func NewCreateLinkedFileWithDefaults() *CreateLinkedFile`

NewCreateLinkedFileWithDefaults instantiates a new CreateLinkedFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateLinkedFile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLinkedFile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLinkedFile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLinkedFile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStoryId

`func (o *CreateLinkedFile) GetStoryId() int64`

GetStoryId returns the StoryId field if non-nil, zero value otherwise.

### GetStoryIdOk

`func (o *CreateLinkedFile) GetStoryIdOk() (*int64, bool)`

GetStoryIdOk returns a tuple with the StoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryId

`func (o *CreateLinkedFile) SetStoryId(v int64)`

SetStoryId sets StoryId field to given value.

### HasStoryId

`func (o *CreateLinkedFile) HasStoryId() bool`

HasStoryId returns a boolean if a field has been set.

### GetName

`func (o *CreateLinkedFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLinkedFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLinkedFile) SetName(v string)`

SetName sets Name field to given value.


### GetThumbnailUrl

`func (o *CreateLinkedFile) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *CreateLinkedFile) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *CreateLinkedFile) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *CreateLinkedFile) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetType

`func (o *CreateLinkedFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateLinkedFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateLinkedFile) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *CreateLinkedFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateLinkedFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateLinkedFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateLinkedFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUploaderId

`func (o *CreateLinkedFile) GetUploaderId() string`

GetUploaderId returns the UploaderId field if non-nil, zero value otherwise.

### GetUploaderIdOk

`func (o *CreateLinkedFile) GetUploaderIdOk() (*string, bool)`

GetUploaderIdOk returns a tuple with the UploaderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderId

`func (o *CreateLinkedFile) SetUploaderId(v string)`

SetUploaderId sets UploaderId field to given value.

### HasUploaderId

`func (o *CreateLinkedFile) HasUploaderId() bool`

HasUploaderId returns a boolean if a field has been set.

### GetContentType

`func (o *CreateLinkedFile) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CreateLinkedFile) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CreateLinkedFile) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CreateLinkedFile) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetUrl

`func (o *CreateLinkedFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateLinkedFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateLinkedFile) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


