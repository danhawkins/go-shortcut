# CreateMilestone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Milestone. | 
**Description** | Pointer to **string** | The Milestone&#39;s description. | [optional] 
**State** | Pointer to **string** | The workflow state that the Milestone is in. | [optional] 
**StartedAtOverride** | Pointer to **time.Time** | A manual override for the time/date the Milestone was started. | [optional] 
**CompletedAtOverride** | Pointer to **time.Time** | A manual override for the time/date the Milestone was completed. | [optional] 
**Categories** | Pointer to [**[]CreateCategoryParams**](CreateCategoryParams.md) | An array of IDs of Categories attached to the Milestone. | [optional] 

## Methods

### NewCreateMilestone

`func NewCreateMilestone(name string, ) *CreateMilestone`

NewCreateMilestone instantiates a new CreateMilestone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMilestoneWithDefaults

`func NewCreateMilestoneWithDefaults() *CreateMilestone`

NewCreateMilestoneWithDefaults instantiates a new CreateMilestone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMilestone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMilestone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMilestone) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateMilestone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMilestone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMilestone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMilestone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *CreateMilestone) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateMilestone) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateMilestone) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateMilestone) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStartedAtOverride

`func (o *CreateMilestone) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *CreateMilestone) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *CreateMilestone) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.

### HasStartedAtOverride

`func (o *CreateMilestone) HasStartedAtOverride() bool`

HasStartedAtOverride returns a boolean if a field has been set.

### GetCompletedAtOverride

`func (o *CreateMilestone) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *CreateMilestone) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *CreateMilestone) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.

### HasCompletedAtOverride

`func (o *CreateMilestone) HasCompletedAtOverride() bool`

HasCompletedAtOverride returns a boolean if a field has been set.

### GetCategories

`func (o *CreateMilestone) GetCategories() []CreateCategoryParams`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CreateMilestone) GetCategoriesOk() (*[]CreateCategoryParams, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CreateMilestone) SetCategories(v []CreateCategoryParams)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CreateMilestone) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


