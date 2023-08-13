# HistoryReferenceStory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **map[string]interface{}** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**AppUrl** | **string** | The application URL of the Story. | 
**Name** | **string** | The name of the entity referenced. | 
**StoryType** | **string** | If the referenced entity is a Story, either \&quot;bug\&quot;, \&quot;chore\&quot;, or \&quot;feature\&quot;. | 

## Methods

### NewHistoryReferenceStory

`func NewHistoryReferenceStory(id map[string]interface{}, entityType string, appUrl string, name string, storyType string, ) *HistoryReferenceStory`

NewHistoryReferenceStory instantiates a new HistoryReferenceStory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryReferenceStoryWithDefaults

`func NewHistoryReferenceStoryWithDefaults() *HistoryReferenceStory`

NewHistoryReferenceStoryWithDefaults instantiates a new HistoryReferenceStory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryReferenceStory) GetId() map[string]interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryReferenceStory) GetIdOk() (*map[string]interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryReferenceStory) SetId(v map[string]interface{})`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryReferenceStory) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryReferenceStory) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryReferenceStory) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetAppUrl

`func (o *HistoryReferenceStory) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *HistoryReferenceStory) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *HistoryReferenceStory) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetName

`func (o *HistoryReferenceStory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoryReferenceStory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HistoryReferenceStory) SetName(v string)`

SetName sets Name field to given value.


### GetStoryType

`func (o *HistoryReferenceStory) GetStoryType() string`

GetStoryType returns the StoryType field if non-nil, zero value otherwise.

### GetStoryTypeOk

`func (o *HistoryReferenceStory) GetStoryTypeOk() (*string, bool)`

GetStoryTypeOk returns a tuple with the StoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryType

`func (o *HistoryReferenceStory) SetStoryType(v string)`

SetStoryType sets StoryType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


