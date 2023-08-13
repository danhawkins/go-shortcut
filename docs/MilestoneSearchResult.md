# MilestoneSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Milestone. | 
**Description** | Pointer to **string** | The Milestone&#39;s description. | [optional] 
**Archived** | **bool** | A boolean indicating whether the Milestone has been archived or not. | 
**Started** | **bool** | A true/false boolean indicating if the Milestone has been started. | 
**EntityType** | **string** | A string description of this resource. | 
**CompletedAtOverride** | **NullableTime** | A manual override for the time/date the Milestone was completed. | 
**StartedAt** | **NullableTime** | The time/date the Milestone was started. | 
**CompletedAt** | **NullableTime** | The time/date the Milestone was completed. | 
**Name** | **string** | The name of the Milestone. | 
**GlobalId** | **string** |  | 
**Completed** | **bool** | A true/false boolean indicating if the Milestone has been completed. | 
**State** | **string** | The workflow state that the Milestone is in. | 
**StartedAtOverride** | **NullableTime** | A manual override for the time/date the Milestone was started. | 
**UpdatedAt** | **time.Time** | The time/date the Milestone was updated. | 
**Categories** | [**[]Category**](Category.md) | An array of Categories attached to the Milestone. | 
**Id** | **int64** | The unique ID of the Milestone. | 
**Position** | **int64** | A number representing the position of the Milestone in relation to every other Milestone within the Workspace. | 
**Stats** | [**MilestoneStats**](MilestoneStats.md) |  | 
**CreatedAt** | **time.Time** | The time/date the Milestone was created. | 

## Methods

### NewMilestoneSearchResult

`func NewMilestoneSearchResult(appUrl string, archived bool, started bool, entityType string, completedAtOverride NullableTime, startedAt NullableTime, completedAt NullableTime, name string, globalId string, completed bool, state string, startedAtOverride NullableTime, updatedAt time.Time, categories []Category, id int64, position int64, stats MilestoneStats, createdAt time.Time, ) *MilestoneSearchResult`

NewMilestoneSearchResult instantiates a new MilestoneSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilestoneSearchResultWithDefaults

`func NewMilestoneSearchResultWithDefaults() *MilestoneSearchResult`

NewMilestoneSearchResultWithDefaults instantiates a new MilestoneSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *MilestoneSearchResult) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *MilestoneSearchResult) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *MilestoneSearchResult) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetDescription

`func (o *MilestoneSearchResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MilestoneSearchResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MilestoneSearchResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MilestoneSearchResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *MilestoneSearchResult) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *MilestoneSearchResult) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *MilestoneSearchResult) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetStarted

`func (o *MilestoneSearchResult) GetStarted() bool`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *MilestoneSearchResult) GetStartedOk() (*bool, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *MilestoneSearchResult) SetStarted(v bool)`

SetStarted sets Started field to given value.


### GetEntityType

`func (o *MilestoneSearchResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *MilestoneSearchResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *MilestoneSearchResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetCompletedAtOverride

`func (o *MilestoneSearchResult) GetCompletedAtOverride() time.Time`

GetCompletedAtOverride returns the CompletedAtOverride field if non-nil, zero value otherwise.

### GetCompletedAtOverrideOk

`func (o *MilestoneSearchResult) GetCompletedAtOverrideOk() (*time.Time, bool)`

GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtOverride

`func (o *MilestoneSearchResult) SetCompletedAtOverride(v time.Time)`

SetCompletedAtOverride sets CompletedAtOverride field to given value.


### SetCompletedAtOverrideNil

`func (o *MilestoneSearchResult) SetCompletedAtOverrideNil(b bool)`

 SetCompletedAtOverrideNil sets the value for CompletedAtOverride to be an explicit nil

### UnsetCompletedAtOverride
`func (o *MilestoneSearchResult) UnsetCompletedAtOverride()`

UnsetCompletedAtOverride ensures that no value is present for CompletedAtOverride, not even an explicit nil
### GetStartedAt

`func (o *MilestoneSearchResult) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MilestoneSearchResult) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MilestoneSearchResult) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *MilestoneSearchResult) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *MilestoneSearchResult) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *MilestoneSearchResult) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *MilestoneSearchResult) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *MilestoneSearchResult) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *MilestoneSearchResult) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *MilestoneSearchResult) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetName

`func (o *MilestoneSearchResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MilestoneSearchResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MilestoneSearchResult) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *MilestoneSearchResult) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *MilestoneSearchResult) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *MilestoneSearchResult) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetCompleted

`func (o *MilestoneSearchResult) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *MilestoneSearchResult) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *MilestoneSearchResult) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.


### GetState

`func (o *MilestoneSearchResult) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MilestoneSearchResult) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MilestoneSearchResult) SetState(v string)`

SetState sets State field to given value.


### GetStartedAtOverride

`func (o *MilestoneSearchResult) GetStartedAtOverride() time.Time`

GetStartedAtOverride returns the StartedAtOverride field if non-nil, zero value otherwise.

### GetStartedAtOverrideOk

`func (o *MilestoneSearchResult) GetStartedAtOverrideOk() (*time.Time, bool)`

GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtOverride

`func (o *MilestoneSearchResult) SetStartedAtOverride(v time.Time)`

SetStartedAtOverride sets StartedAtOverride field to given value.


### SetStartedAtOverrideNil

`func (o *MilestoneSearchResult) SetStartedAtOverrideNil(b bool)`

 SetStartedAtOverrideNil sets the value for StartedAtOverride to be an explicit nil

### UnsetStartedAtOverride
`func (o *MilestoneSearchResult) UnsetStartedAtOverride()`

UnsetStartedAtOverride ensures that no value is present for StartedAtOverride, not even an explicit nil
### GetUpdatedAt

`func (o *MilestoneSearchResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MilestoneSearchResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MilestoneSearchResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCategories

`func (o *MilestoneSearchResult) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MilestoneSearchResult) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MilestoneSearchResult) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetId

`func (o *MilestoneSearchResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MilestoneSearchResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MilestoneSearchResult) SetId(v int64)`

SetId sets Id field to given value.


### GetPosition

`func (o *MilestoneSearchResult) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MilestoneSearchResult) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *MilestoneSearchResult) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetStats

`func (o *MilestoneSearchResult) GetStats() MilestoneStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *MilestoneSearchResult) GetStatsOk() (*MilestoneStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *MilestoneSearchResult) SetStats(v MilestoneStats)`

SetStats sets Stats field to given value.


### GetCreatedAt

`func (o *MilestoneSearchResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MilestoneSearchResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MilestoneSearchResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


