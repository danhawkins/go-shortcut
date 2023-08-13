# UpdateEntityTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The updated template name. | [optional] 
**StoryContents** | Pointer to [**UpdateStoryContents**](UpdateStoryContents.md) |  | [optional] 

## Methods

### NewUpdateEntityTemplate

`func NewUpdateEntityTemplate() *UpdateEntityTemplate`

NewUpdateEntityTemplate instantiates a new UpdateEntityTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEntityTemplateWithDefaults

`func NewUpdateEntityTemplateWithDefaults() *UpdateEntityTemplate`

NewUpdateEntityTemplateWithDefaults instantiates a new UpdateEntityTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateEntityTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEntityTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEntityTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateEntityTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStoryContents

`func (o *UpdateEntityTemplate) GetStoryContents() UpdateStoryContents`

GetStoryContents returns the StoryContents field if non-nil, zero value otherwise.

### GetStoryContentsOk

`func (o *UpdateEntityTemplate) GetStoryContentsOk() (*UpdateStoryContents, bool)`

GetStoryContentsOk returns a tuple with the StoryContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryContents

`func (o *UpdateEntityTemplate) SetStoryContents(v UpdateStoryContents)`

SetStoryContents sets StoryContents field to given value.

### HasStoryContents

`func (o *UpdateEntityTemplate) HasStoryContents() bool`

HasStoryContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


