# HistoryReferenceCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **map[string]interface{}** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Message** | **string** | The message from the Commit. | 
**Url** | **string** | The external URL for the Branch. | 

## Methods

### NewHistoryReferenceCommit

`func NewHistoryReferenceCommit(id map[string]interface{}, entityType string, message string, url string, ) *HistoryReferenceCommit`

NewHistoryReferenceCommit instantiates a new HistoryReferenceCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryReferenceCommitWithDefaults

`func NewHistoryReferenceCommitWithDefaults() *HistoryReferenceCommit`

NewHistoryReferenceCommitWithDefaults instantiates a new HistoryReferenceCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryReferenceCommit) GetId() map[string]interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryReferenceCommit) GetIdOk() (*map[string]interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryReferenceCommit) SetId(v map[string]interface{})`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryReferenceCommit) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryReferenceCommit) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryReferenceCommit) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetMessage

`func (o *HistoryReferenceCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HistoryReferenceCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HistoryReferenceCommit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetUrl

`func (o *HistoryReferenceCommit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HistoryReferenceCommit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HistoryReferenceCommit) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


