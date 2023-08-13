# UpdateIteration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FollowerIds** | Pointer to **[]string** | An array of UUIDs for any Members you want to add as Followers. | [optional] 
**GroupIds** | Pointer to **[]string** | An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI. | [optional] 
**Labels** | Pointer to [**[]CreateLabelParams**](CreateLabelParams.md) | An array of Labels attached to the Iteration. | [optional] 
**Description** | Pointer to **string** | The description of the Iteration. | [optional] 
**Name** | Pointer to **string** | The name of this Iteration | [optional] 
**StartDate** | Pointer to **string** | The date this Iteration begins, e.g. 2019-07-01 | [optional] 
**EndDate** | Pointer to **string** | The date this Iteration ends, e.g. 2019-07-05. | [optional] 

## Methods

### NewUpdateIteration

`func NewUpdateIteration() *UpdateIteration`

NewUpdateIteration instantiates a new UpdateIteration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIterationWithDefaults

`func NewUpdateIterationWithDefaults() *UpdateIteration`

NewUpdateIterationWithDefaults instantiates a new UpdateIteration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowerIds

`func (o *UpdateIteration) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *UpdateIteration) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *UpdateIteration) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *UpdateIteration) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetGroupIds

`func (o *UpdateIteration) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *UpdateIteration) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *UpdateIteration) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *UpdateIteration) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateIteration) GetLabels() []CreateLabelParams`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateIteration) GetLabelsOk() (*[]CreateLabelParams, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateIteration) SetLabels(v []CreateLabelParams)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateIteration) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIteration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIteration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIteration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIteration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateIteration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIteration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIteration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIteration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateIteration) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateIteration) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateIteration) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateIteration) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateIteration) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateIteration) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateIteration) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateIteration) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


