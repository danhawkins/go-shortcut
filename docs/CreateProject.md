# CreateProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The Project description. | [optional] 
**Color** | Pointer to **string** | The color you wish to use for the Project in the system. | [optional] 
**Name** | **string** | The name of the Project. | 
**StartTime** | Pointer to **time.Time** | The date at which the Project was started. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Defaults to the time/date it is created but can be set to reflect another date. | [optional] 
**FollowerIds** | Pointer to **[]string** | An array of UUIDs for any members you want to add as Owners on this new Epic. | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 
**TeamId** | **int64** | The ID of the team the project belongs to. | 
**IterationLength** | Pointer to **int64** | The number of weeks per iteration in this Project. | [optional] 
**Abbreviation** | Pointer to **string** | The Project abbreviation used in Story summaries. Should be kept to 3 characters at most. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Defaults to the time/date it is created but can be set to reflect another date. | [optional] 

## Methods

### NewCreateProject

`func NewCreateProject(name string, teamId int64, ) *CreateProject`

NewCreateProject instantiates a new CreateProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectWithDefaults

`func NewCreateProjectWithDefaults() *CreateProject`

NewCreateProjectWithDefaults instantiates a new CreateProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColor

`func (o *CreateProject) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateProject) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateProject) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateProject) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetName

`func (o *CreateProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProject) SetName(v string)`

SetName sets Name field to given value.


### GetStartTime

`func (o *CreateProject) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateProject) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateProject) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CreateProject) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateProject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateProject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateProject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFollowerIds

`func (o *CreateProject) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *CreateProject) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *CreateProject) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *CreateProject) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateProject) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateProject) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateProject) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateProject) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTeamId

`func (o *CreateProject) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CreateProject) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CreateProject) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetIterationLength

`func (o *CreateProject) GetIterationLength() int64`

GetIterationLength returns the IterationLength field if non-nil, zero value otherwise.

### GetIterationLengthOk

`func (o *CreateProject) GetIterationLengthOk() (*int64, bool)`

GetIterationLengthOk returns a tuple with the IterationLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationLength

`func (o *CreateProject) SetIterationLength(v int64)`

SetIterationLength sets IterationLength field to given value.

### HasIterationLength

`func (o *CreateProject) HasIterationLength() bool`

HasIterationLength returns a boolean if a field has been set.

### GetAbbreviation

`func (o *CreateProject) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *CreateProject) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *CreateProject) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.

### HasAbbreviation

`func (o *CreateProject) HasAbbreviation() bool`

HasAbbreviation returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


