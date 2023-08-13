# PullRequestLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | 
**Id** | **int64** | The unique ID of the VCS Label. | 
**Color** | **string** | The color of the VCS label. | 
**Description** | Pointer to **NullableString** | The description of the VCS label. | [optional] 
**Name** | **string** | The name of the VCS label. | 

## Methods

### NewPullRequestLabel

`func NewPullRequestLabel(entityType string, id int64, color string, name string, ) *PullRequestLabel`

NewPullRequestLabel instantiates a new PullRequestLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestLabelWithDefaults

`func NewPullRequestLabelWithDefaults() *PullRequestLabel`

NewPullRequestLabelWithDefaults instantiates a new PullRequestLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *PullRequestLabel) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *PullRequestLabel) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *PullRequestLabel) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetId

`func (o *PullRequestLabel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequestLabel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequestLabel) SetId(v int64)`

SetId sets Id field to given value.


### GetColor

`func (o *PullRequestLabel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PullRequestLabel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PullRequestLabel) SetColor(v string)`

SetColor sets Color field to given value.


### GetDescription

`func (o *PullRequestLabel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PullRequestLabel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PullRequestLabel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PullRequestLabel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PullRequestLabel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PullRequestLabel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *PullRequestLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PullRequestLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PullRequestLabel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


