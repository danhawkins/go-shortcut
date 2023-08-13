/*
Shortcut API

Shortcut API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProject{}

// UpdateProject struct for UpdateProject
type UpdateProject struct {
	// The Project's description.
	Description *string `json:"description,omitempty"`
	// A true/false boolean indicating whether the Story is in archived state.
	Archived *bool `json:"archived,omitempty"`
	// The number of days before the thermometer appears in the Story summary.
	DaysToThermometer *int64 `json:"days_to_thermometer,omitempty"`
	// The color that represents the Project in the UI.
	Color *string `json:"color,omitempty"`
	// The Project's name.
	Name *string `json:"name,omitempty"`
	// An array of UUIDs for any Members you want to add as Followers.
	FollowerIds []string `json:"follower_ids,omitempty"`
	// Configuration to enable or disable thermometers in the Story summary.
	ShowThermometer *bool `json:"show_thermometer,omitempty"`
	// The ID of the team the project belongs to.
	TeamId *int64 `json:"team_id,omitempty"`
	// The Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation *string `json:"abbreviation,omitempty"`
}

// NewUpdateProject instantiates a new UpdateProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProject() *UpdateProject {
	this := UpdateProject{}
	return &this
}

// NewUpdateProjectWithDefaults instantiates a new UpdateProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProjectWithDefaults() *UpdateProject {
	this := UpdateProject{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateProject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateProject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateProject) SetDescription(v string) {
	o.Description = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *UpdateProject) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *UpdateProject) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *UpdateProject) SetArchived(v bool) {
	o.Archived = &v
}

// GetDaysToThermometer returns the DaysToThermometer field value if set, zero value otherwise.
func (o *UpdateProject) GetDaysToThermometer() int64 {
	if o == nil || IsNil(o.DaysToThermometer) {
		var ret int64
		return ret
	}
	return *o.DaysToThermometer
}

// GetDaysToThermometerOk returns a tuple with the DaysToThermometer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetDaysToThermometerOk() (*int64, bool) {
	if o == nil || IsNil(o.DaysToThermometer) {
		return nil, false
	}
	return o.DaysToThermometer, true
}

// HasDaysToThermometer returns a boolean if a field has been set.
func (o *UpdateProject) HasDaysToThermometer() bool {
	if o != nil && !IsNil(o.DaysToThermometer) {
		return true
	}

	return false
}

// SetDaysToThermometer gets a reference to the given int64 and assigns it to the DaysToThermometer field.
func (o *UpdateProject) SetDaysToThermometer(v int64) {
	o.DaysToThermometer = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *UpdateProject) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *UpdateProject) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *UpdateProject) SetColor(v string) {
	o.Color = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateProject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateProject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateProject) SetName(v string) {
	o.Name = &v
}

// GetFollowerIds returns the FollowerIds field value if set, zero value otherwise.
func (o *UpdateProject) GetFollowerIds() []string {
	if o == nil || IsNil(o.FollowerIds) {
		var ret []string
		return ret
	}
	return o.FollowerIds
}

// GetFollowerIdsOk returns a tuple with the FollowerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetFollowerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FollowerIds) {
		return nil, false
	}
	return o.FollowerIds, true
}

// HasFollowerIds returns a boolean if a field has been set.
func (o *UpdateProject) HasFollowerIds() bool {
	if o != nil && !IsNil(o.FollowerIds) {
		return true
	}

	return false
}

// SetFollowerIds gets a reference to the given []string and assigns it to the FollowerIds field.
func (o *UpdateProject) SetFollowerIds(v []string) {
	o.FollowerIds = v
}

// GetShowThermometer returns the ShowThermometer field value if set, zero value otherwise.
func (o *UpdateProject) GetShowThermometer() bool {
	if o == nil || IsNil(o.ShowThermometer) {
		var ret bool
		return ret
	}
	return *o.ShowThermometer
}

// GetShowThermometerOk returns a tuple with the ShowThermometer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetShowThermometerOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowThermometer) {
		return nil, false
	}
	return o.ShowThermometer, true
}

// HasShowThermometer returns a boolean if a field has been set.
func (o *UpdateProject) HasShowThermometer() bool {
	if o != nil && !IsNil(o.ShowThermometer) {
		return true
	}

	return false
}

// SetShowThermometer gets a reference to the given bool and assigns it to the ShowThermometer field.
func (o *UpdateProject) SetShowThermometer(v bool) {
	o.ShowThermometer = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *UpdateProject) GetTeamId() int64 {
	if o == nil || IsNil(o.TeamId) {
		var ret int64
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetTeamIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *UpdateProject) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given int64 and assigns it to the TeamId field.
func (o *UpdateProject) SetTeamId(v int64) {
	o.TeamId = &v
}

// GetAbbreviation returns the Abbreviation field value if set, zero value otherwise.
func (o *UpdateProject) GetAbbreviation() string {
	if o == nil || IsNil(o.Abbreviation) {
		var ret string
		return ret
	}
	return *o.Abbreviation
}

// GetAbbreviationOk returns a tuple with the Abbreviation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetAbbreviationOk() (*string, bool) {
	if o == nil || IsNil(o.Abbreviation) {
		return nil, false
	}
	return o.Abbreviation, true
}

// HasAbbreviation returns a boolean if a field has been set.
func (o *UpdateProject) HasAbbreviation() bool {
	if o != nil && !IsNil(o.Abbreviation) {
		return true
	}

	return false
}

// SetAbbreviation gets a reference to the given string and assigns it to the Abbreviation field.
func (o *UpdateProject) SetAbbreviation(v string) {
	o.Abbreviation = &v
}

func (o UpdateProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.DaysToThermometer) {
		toSerialize["days_to_thermometer"] = o.DaysToThermometer
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FollowerIds) {
		toSerialize["follower_ids"] = o.FollowerIds
	}
	if !IsNil(o.ShowThermometer) {
		toSerialize["show_thermometer"] = o.ShowThermometer
	}
	if !IsNil(o.TeamId) {
		toSerialize["team_id"] = o.TeamId
	}
	if !IsNil(o.Abbreviation) {
		toSerialize["abbreviation"] = o.Abbreviation
	}
	return toSerialize, nil
}

type NullableUpdateProject struct {
	value *UpdateProject
	isSet bool
}

func (v NullableUpdateProject) Get() *UpdateProject {
	return v.value
}

func (v *NullableUpdateProject) Set(val *UpdateProject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProject(val *UpdateProject) *NullableUpdateProject {
	return &NullableUpdateProject{value: val, isSet: true}
}

func (v NullableUpdateProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

