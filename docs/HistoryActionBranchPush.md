# HistoryActionBranchPush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the entity referenced. | 
**EntityType** | **string** | The type of entity referenced. | 
**Name** | **string** | The name of the VCS Branch that was pushed | 
**Url** | **string** | The URL from the provider of the VCS Branch that was pushed | 
**Action** | **string** | The action of the entity referenced. | 

## Methods

### NewHistoryActionBranchPush

`func NewHistoryActionBranchPush(id int64, entityType string, name string, url string, action string, ) *HistoryActionBranchPush`

NewHistoryActionBranchPush instantiates a new HistoryActionBranchPush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionBranchPushWithDefaults

`func NewHistoryActionBranchPushWithDefaults() *HistoryActionBranchPush`

NewHistoryActionBranchPushWithDefaults instantiates a new HistoryActionBranchPush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryActionBranchPush) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionBranchPush) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionBranchPush) SetId(v int64)`

SetId sets Id field to given value.


### GetEntityType

`func (o *HistoryActionBranchPush) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryActionBranchPush) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryActionBranchPush) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetName

`func (o *HistoryActionBranchPush) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HistoryActionBranchPush) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HistoryActionBranchPush) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *HistoryActionBranchPush) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HistoryActionBranchPush) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HistoryActionBranchPush) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAction

`func (o *HistoryActionBranchPush) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryActionBranchPush) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryActionBranchPush) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


