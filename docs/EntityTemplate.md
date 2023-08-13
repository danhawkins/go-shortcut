# EntityTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | 
**Id** | **string** | The unique identifier for the entity template. | 
**CreatedAt** | **time.Time** | The time/date when the entity template was created. | 
**UpdatedAt** | **time.Time** | The time/date when the entity template was last updated. | 
**Name** | **string** | The template&#39;s name. | 
**AuthorId** | **string** | The unique ID of the member who created the template. | 
**LastUsedAt** | **time.Time** | The last time that someone created an entity using this template. | 
**StoryContents** | [**StoryContents**](StoryContents.md) |  | 

## Methods

### NewEntityTemplate

`func NewEntityTemplate(entityType string, id string, createdAt time.Time, updatedAt time.Time, name string, authorId string, lastUsedAt time.Time, storyContents StoryContents, ) *EntityTemplate`

NewEntityTemplate instantiates a new EntityTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTemplateWithDefaults

`func NewEntityTemplateWithDefaults() *EntityTemplate`

NewEntityTemplateWithDefaults instantiates a new EntityTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *EntityTemplate) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EntityTemplate) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EntityTemplate) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetId

`func (o *EntityTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EntityTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EntityTemplate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityTemplate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityTemplate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetName

`func (o *EntityTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetAuthorId

`func (o *EntityTemplate) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *EntityTemplate) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *EntityTemplate) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetLastUsedAt

`func (o *EntityTemplate) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *EntityTemplate) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *EntityTemplate) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### GetStoryContents

`func (o *EntityTemplate) GetStoryContents() StoryContents`

GetStoryContents returns the StoryContents field if non-nil, zero value otherwise.

### GetStoryContentsOk

`func (o *EntityTemplate) GetStoryContentsOk() (*StoryContents, bool)`

GetStoryContentsOk returns a tuple with the StoryContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryContents

`func (o *EntityTemplate) SetStoryContents(v StoryContents)`

SetStoryContents sets StoryContents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


