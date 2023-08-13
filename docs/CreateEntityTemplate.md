# CreateEntityTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new entity template | 
**AuthorId** | Pointer to **string** | The id of the user creating this template. | [optional] 
**StoryContents** | [**CreateStoryContents**](CreateStoryContents.md) |  | 

## Methods

### NewCreateEntityTemplate

`func NewCreateEntityTemplate(name string, storyContents CreateStoryContents, ) *CreateEntityTemplate`

NewCreateEntityTemplate instantiates a new CreateEntityTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntityTemplateWithDefaults

`func NewCreateEntityTemplateWithDefaults() *CreateEntityTemplate`

NewCreateEntityTemplateWithDefaults instantiates a new CreateEntityTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEntityTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEntityTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEntityTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetAuthorId

`func (o *CreateEntityTemplate) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *CreateEntityTemplate) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *CreateEntityTemplate) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *CreateEntityTemplate) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetStoryContents

`func (o *CreateEntityTemplate) GetStoryContents() CreateStoryContents`

GetStoryContents returns the StoryContents field if non-nil, zero value otherwise.

### GetStoryContentsOk

`func (o *CreateEntityTemplate) GetStoryContentsOk() (*CreateStoryContents, bool)`

GetStoryContentsOk returns a tuple with the StoryContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryContents

`func (o *CreateEntityTemplate) SetStoryContents(v CreateStoryContents)`

SetStoryContents sets StoryContents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


