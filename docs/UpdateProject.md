# UpdateProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The Project&#39;s description. | [optional] 
**Archived** | Pointer to **bool** | A true/false boolean indicating whether the Story is in archived state. | [optional] 
**DaysToThermometer** | Pointer to **int64** | The number of days before the thermometer appears in the Story summary. | [optional] 
**Color** | Pointer to **string** | The color that represents the Project in the UI. | [optional] 
**Name** | Pointer to **string** | The Project&#39;s name. | [optional] 
**FollowerIds** | Pointer to **[]string** | An array of UUIDs for any Members you want to add as Followers. | [optional] 
**ShowThermometer** | Pointer to **bool** | Configuration to enable or disable thermometers in the Story summary. | [optional] 
**TeamId** | Pointer to **int64** | The ID of the team the project belongs to. | [optional] 
**Abbreviation** | Pointer to **string** | The Project abbreviation used in Story summaries. Should be kept to 3 characters at most. | [optional] 

## Methods

### NewUpdateProject

`func NewUpdateProject() *UpdateProject`

NewUpdateProject instantiates a new UpdateProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectWithDefaults

`func NewUpdateProjectWithDefaults() *UpdateProject`

NewUpdateProjectWithDefaults instantiates a new UpdateProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArchived

`func (o *UpdateProject) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateProject) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateProject) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateProject) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDaysToThermometer

`func (o *UpdateProject) GetDaysToThermometer() int64`

GetDaysToThermometer returns the DaysToThermometer field if non-nil, zero value otherwise.

### GetDaysToThermometerOk

`func (o *UpdateProject) GetDaysToThermometerOk() (*int64, bool)`

GetDaysToThermometerOk returns a tuple with the DaysToThermometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToThermometer

`func (o *UpdateProject) SetDaysToThermometer(v int64)`

SetDaysToThermometer sets DaysToThermometer field to given value.

### HasDaysToThermometer

`func (o *UpdateProject) HasDaysToThermometer() bool`

HasDaysToThermometer returns a boolean if a field has been set.

### GetColor

`func (o *UpdateProject) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateProject) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateProject) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateProject) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetName

`func (o *UpdateProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFollowerIds

`func (o *UpdateProject) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *UpdateProject) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *UpdateProject) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.

### HasFollowerIds

`func (o *UpdateProject) HasFollowerIds() bool`

HasFollowerIds returns a boolean if a field has been set.

### GetShowThermometer

`func (o *UpdateProject) GetShowThermometer() bool`

GetShowThermometer returns the ShowThermometer field if non-nil, zero value otherwise.

### GetShowThermometerOk

`func (o *UpdateProject) GetShowThermometerOk() (*bool, bool)`

GetShowThermometerOk returns a tuple with the ShowThermometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowThermometer

`func (o *UpdateProject) SetShowThermometer(v bool)`

SetShowThermometer sets ShowThermometer field to given value.

### HasShowThermometer

`func (o *UpdateProject) HasShowThermometer() bool`

HasShowThermometer returns a boolean if a field has been set.

### GetTeamId

`func (o *UpdateProject) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *UpdateProject) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *UpdateProject) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *UpdateProject) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetAbbreviation

`func (o *UpdateProject) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *UpdateProject) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *UpdateProject) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.

### HasAbbreviation

`func (o *UpdateProject) HasAbbreviation() bool`

HasAbbreviation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


