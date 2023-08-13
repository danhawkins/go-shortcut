# EpicState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of what sort of Epics belong in that Epic State. | 
**EntityType** | **string** | A string description of this resource. | 
**Color** | Pointer to **string** | The hex color for this Epic State. | [optional] 
**Name** | **string** | The Epic State&#39;s name. | 
**GlobalId** | **string** |  | 
**Type** | **string** | The type of Epic State (Unstarted, Started, or Done) | 
**UpdatedAt** | **time.Time** | When the Epic State was last updated. | 
**Id** | **int64** | The unique ID of the Epic State. | 
**Position** | **int64** | The position that the Epic State is in, starting with 0 at the left. | 
**CreatedAt** | **time.Time** | The time/date the Epic State was created. | 

## Methods

### NewEpicState

`func NewEpicState(description string, entityType string, name string, globalId string, type_ string, updatedAt time.Time, id int64, position int64, createdAt time.Time, ) *EpicState`

NewEpicState instantiates a new EpicState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpicStateWithDefaults

`func NewEpicStateWithDefaults() *EpicState`

NewEpicStateWithDefaults instantiates a new EpicState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EpicState) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EpicState) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EpicState) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntityType

`func (o *EpicState) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EpicState) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EpicState) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetColor

`func (o *EpicState) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *EpicState) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *EpicState) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *EpicState) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetName

`func (o *EpicState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EpicState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EpicState) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *EpicState) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *EpicState) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *EpicState) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetType

`func (o *EpicState) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EpicState) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EpicState) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *EpicState) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EpicState) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EpicState) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetId

`func (o *EpicState) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EpicState) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EpicState) SetId(v int64)`

SetId sets Id field to given value.


### GetPosition

`func (o *EpicState) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EpicState) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EpicState) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetCreatedAt

`func (o *EpicState) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EpicState) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EpicState) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


