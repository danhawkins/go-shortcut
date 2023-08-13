# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Project. | 
**Description** | **NullableString** | The description of the Project. | 
**Archived** | **bool** | True/false boolean indicating whether the Project is in an Archived state. | 
**EntityType** | **string** | A string description of this resource. | 
**DaysToThermometer** | **int64** | The number of days before the thermometer appears in the Story summary. | 
**Color** | **NullableString** | The color associated with the Project in the Shortcut member interface. | 
**WorkflowId** | **int64** | The ID of the workflow the project belongs to. | 
**Name** | **string** | The name of the Project | 
**GlobalId** | **string** | The Global ID of the Project. | 
**StartTime** | **time.Time** | The date at which the Project was started. | 
**UpdatedAt** | **NullableTime** | The time/date that the Project was last updated. | 
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | 
**ExternalId** | **NullableString** | This field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here. | 
**Id** | **int64** | The unique ID of the Project. | 
**ShowThermometer** | **bool** | Configuration to enable or disable thermometers in the Story summary. | 
**TeamId** | **int64** | The ID of the team the project belongs to. | 
**IterationLength** | **int64** | The number of weeks per iteration in this Project. | 
**Abbreviation** | **NullableString** | The Project abbreviation used in Story summaries. Should be kept to 3 characters at most. | 
**Stats** | [**ProjectStats**](ProjectStats.md) |  | 
**CreatedAt** | **NullableTime** | The time/date that the Project was created. | 

## Methods

### NewProject

`func NewProject(appUrl string, description NullableString, archived bool, entityType string, daysToThermometer int64, color NullableString, workflowId int64, name string, globalId string, startTime time.Time, updatedAt NullableTime, followerIds []string, externalId NullableString, id int64, showThermometer bool, teamId int64, iterationLength int64, abbreviation NullableString, stats ProjectStats, createdAt NullableTime, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *Project) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *Project) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *Project) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Project) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Project) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetArchived

`func (o *Project) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Project) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Project) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetEntityType

`func (o *Project) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Project) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Project) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetDaysToThermometer

`func (o *Project) GetDaysToThermometer() int64`

GetDaysToThermometer returns the DaysToThermometer field if non-nil, zero value otherwise.

### GetDaysToThermometerOk

`func (o *Project) GetDaysToThermometerOk() (*int64, bool)`

GetDaysToThermometerOk returns a tuple with the DaysToThermometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToThermometer

`func (o *Project) SetDaysToThermometer(v int64)`

SetDaysToThermometer sets DaysToThermometer field to given value.


### GetColor

`func (o *Project) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Project) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Project) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *Project) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *Project) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetWorkflowId

`func (o *Project) GetWorkflowId() int64`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *Project) GetWorkflowIdOk() (*int64, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *Project) SetWorkflowId(v int64)`

SetWorkflowId sets WorkflowId field to given value.


### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *Project) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *Project) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *Project) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetStartTime

`func (o *Project) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Project) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Project) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetUpdatedAt

`func (o *Project) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Project) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Project) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Project) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Project) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetFollowerIds

`func (o *Project) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *Project) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *Project) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.


### GetExternalId

`func (o *Project) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Project) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Project) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *Project) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Project) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetId

`func (o *Project) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v int64)`

SetId sets Id field to given value.


### GetShowThermometer

`func (o *Project) GetShowThermometer() bool`

GetShowThermometer returns the ShowThermometer field if non-nil, zero value otherwise.

### GetShowThermometerOk

`func (o *Project) GetShowThermometerOk() (*bool, bool)`

GetShowThermometerOk returns a tuple with the ShowThermometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowThermometer

`func (o *Project) SetShowThermometer(v bool)`

SetShowThermometer sets ShowThermometer field to given value.


### GetTeamId

`func (o *Project) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Project) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Project) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetIterationLength

`func (o *Project) GetIterationLength() int64`

GetIterationLength returns the IterationLength field if non-nil, zero value otherwise.

### GetIterationLengthOk

`func (o *Project) GetIterationLengthOk() (*int64, bool)`

GetIterationLengthOk returns a tuple with the IterationLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationLength

`func (o *Project) SetIterationLength(v int64)`

SetIterationLength sets IterationLength field to given value.


### GetAbbreviation

`func (o *Project) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *Project) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *Project) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.


### SetAbbreviationNil

`func (o *Project) SetAbbreviationNil(b bool)`

 SetAbbreviationNil sets the value for Abbreviation to be an explicit nil

### UnsetAbbreviation
`func (o *Project) UnsetAbbreviation()`

UnsetAbbreviation ensures that no value is present for Abbreviation, not even an explicit nil
### GetStats

`func (o *Project) GetStats() ProjectStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Project) GetStatsOk() (*ProjectStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Project) SetStats(v ProjectStats)`

SetStats sets Stats field to given value.


### GetCreatedAt

`func (o *Project) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Project) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Project) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Project) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Project) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


