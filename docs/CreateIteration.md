# CreateIteration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FollowerIds** | Pointer to **[]string** | An array of UUIDs for any Members you want to add as Followers. | [optional] 
**GroupIds** | Pointer to **[]string** | An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI. | [optional] 
**Labels** | Pointer to [**[]CreateLabelParams**](CreateLabelParams.md) | An array of Labels attached to the Iteration. | [optional] 
**Description** | Pointer to **string** | The description of the Iteration. | [optional] 
**Name** | **string** | The name of this Iteration. | 
**StartDate** | **string** | The date this Iteration begins, e.g. 2019-07-01. | 
**EndDate** | **string** | The date this Iteration ends, e.g. 2019-07-01. | 

## Methods

### NewCreateIteration

`func NewCreateIteration(name string, startDate string, endDate string, ) *CreateIteration`

NewCreateIteration instantiates a new CreateIteration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIterationWithDefaults

`func NewCreateIterationWithDefaults() *CreateIteration`

NewCreateIterationWithDefaults instantiates a new CreateIteration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowerIds

`func (o *CreateIteration) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *CreateIteration) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *CreateIteration) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *CreateIteration) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetGroupIds

`func (o *CreateIteration) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *CreateIteration) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *CreateIteration) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *CreateIteration) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetLabels

`func (o *CreateIteration) GetLabels() []CreateLabelParams`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateIteration) GetLabelsOk() (*[]CreateLabelParams, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateIteration) SetLabels(v []CreateLabelParams)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateIteration) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *CreateIteration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIteration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIteration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIteration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateIteration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIteration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIteration) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *CreateIteration) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateIteration) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateIteration) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *CreateIteration) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateIteration) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateIteration) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


