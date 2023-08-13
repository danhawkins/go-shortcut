/*
Shortcut API

Shortcut API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the IterationSlim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IterationSlim{}

// IterationSlim IterationSlim represents the same resource as an Iteration, but is more light-weight. Use the [Get Iteration](#Get-Iteration) endpoint to fetch the unabridged payload for an Iteration. 
type IterationSlim struct {
	// The Shortcut application url for the Iteration.
	AppUrl string `json:"app_url"`
	// A string description of this resource
	EntityType string `json:"entity_type"`
	// An array of labels attached to the iteration.
	Labels []Label `json:"labels"`
	// Deprecated: use member_mention_ids.
	MentionIds []string `json:"mention_ids"`
	// An array of Member IDs that have been mentioned in the Story description.
	MemberMentionIds []string `json:"member_mention_ids"`
	// An array containing Group IDs and Group-owned story counts for the Iteration's associated groups.
	AssociatedGroups []IterationAssociatedGroup `json:"associated_groups"`
	// The name of the iteration.
	Name string `json:"name"`
	GlobalId string `json:"global_id"`
	// An array of label ids attached to the iteration.
	LabelIds []int64 `json:"label_ids"`
	// The instant when this iteration was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// An array of Group IDs that have been mentioned in the Story description.
	GroupMentionIds []string `json:"group_mention_ids"`
	// The date this iteration begins.
	EndDate time.Time `json:"end_date"`
	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []string `json:"follower_ids"`
	// An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	GroupIds []string `json:"group_ids"`
	// The date this iteration begins.
	StartDate time.Time `json:"start_date"`
	// The status of the iteration. Values are either \"unstarted\", \"started\", or \"done\".
	Status string `json:"status"`
	// The ID of the iteration.
	Id int64 `json:"id"`
	Stats IterationStats `json:"stats"`
	// The instant when this iteration was created.
	CreatedAt time.Time `json:"created_at"`
}

// NewIterationSlim instantiates a new IterationSlim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIterationSlim(appUrl string, entityType string, labels []Label, mentionIds []string, memberMentionIds []string, associatedGroups []IterationAssociatedGroup, name string, globalId string, labelIds []int64, updatedAt time.Time, groupMentionIds []string, endDate time.Time, followerIds []string, groupIds []string, startDate time.Time, status string, id int64, stats IterationStats, createdAt time.Time) *IterationSlim {
	this := IterationSlim{}
	this.AppUrl = appUrl
	this.EntityType = entityType
	this.Labels = labels
	this.MentionIds = mentionIds
	this.MemberMentionIds = memberMentionIds
	this.AssociatedGroups = associatedGroups
	this.Name = name
	this.GlobalId = globalId
	this.LabelIds = labelIds
	this.UpdatedAt = updatedAt
	this.GroupMentionIds = groupMentionIds
	this.EndDate = endDate
	this.FollowerIds = followerIds
	this.GroupIds = groupIds
	this.StartDate = startDate
	this.Status = status
	this.Id = id
	this.Stats = stats
	this.CreatedAt = createdAt
	return &this
}

// NewIterationSlimWithDefaults instantiates a new IterationSlim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIterationSlimWithDefaults() *IterationSlim {
	this := IterationSlim{}
	return &this
}

// GetAppUrl returns the AppUrl field value
func (o *IterationSlim) GetAppUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppUrl
}

// GetAppUrlOk returns a tuple with the AppUrl field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetAppUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUrl, true
}

// SetAppUrl sets field value
func (o *IterationSlim) SetAppUrl(v string) {
	o.AppUrl = v
}

// GetEntityType returns the EntityType field value
func (o *IterationSlim) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *IterationSlim) SetEntityType(v string) {
	o.EntityType = v
}

// GetLabels returns the Labels field value
func (o *IterationSlim) GetLabels() []Label {
	if o == nil {
		var ret []Label
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetLabelsOk() ([]Label, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *IterationSlim) SetLabels(v []Label) {
	o.Labels = v
}

// GetMentionIds returns the MentionIds field value
func (o *IterationSlim) GetMentionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MentionIds
}

// GetMentionIdsOk returns a tuple with the MentionIds field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetMentionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionIds, true
}

// SetMentionIds sets field value
func (o *IterationSlim) SetMentionIds(v []string) {
	o.MentionIds = v
}

// GetMemberMentionIds returns the MemberMentionIds field value
func (o *IterationSlim) GetMemberMentionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MemberMentionIds
}

// GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetMemberMentionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberMentionIds, true
}

// SetMemberMentionIds sets field value
func (o *IterationSlim) SetMemberMentionIds(v []string) {
	o.MemberMentionIds = v
}

// GetAssociatedGroups returns the AssociatedGroups field value
func (o *IterationSlim) GetAssociatedGroups() []IterationAssociatedGroup {
	if o == nil {
		var ret []IterationAssociatedGroup
		return ret
	}

	return o.AssociatedGroups
}

// GetAssociatedGroupsOk returns a tuple with the AssociatedGroups field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetAssociatedGroupsOk() ([]IterationAssociatedGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssociatedGroups, true
}

// SetAssociatedGroups sets field value
func (o *IterationSlim) SetAssociatedGroups(v []IterationAssociatedGroup) {
	o.AssociatedGroups = v
}

// GetName returns the Name field value
func (o *IterationSlim) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IterationSlim) SetName(v string) {
	o.Name = v
}

// GetGlobalId returns the GlobalId field value
func (o *IterationSlim) GetGlobalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalId
}

// GetGlobalIdOk returns a tuple with the GlobalId field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetGlobalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalId, true
}

// SetGlobalId sets field value
func (o *IterationSlim) SetGlobalId(v string) {
	o.GlobalId = v
}

// GetLabelIds returns the LabelIds field value
func (o *IterationSlim) GetLabelIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.LabelIds
}

// GetLabelIdsOk returns a tuple with the LabelIds field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetLabelIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LabelIds, true
}

// SetLabelIds sets field value
func (o *IterationSlim) SetLabelIds(v []int64) {
	o.LabelIds = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *IterationSlim) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *IterationSlim) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetGroupMentionIds returns the GroupMentionIds field value
func (o *IterationSlim) GetGroupMentionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupMentionIds
}

// GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetGroupMentionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupMentionIds, true
}

// SetGroupMentionIds sets field value
func (o *IterationSlim) SetGroupMentionIds(v []string) {
	o.GroupMentionIds = v
}

// GetEndDate returns the EndDate field value
func (o *IterationSlim) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *IterationSlim) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetFollowerIds returns the FollowerIds field value
func (o *IterationSlim) GetFollowerIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FollowerIds
}

// GetFollowerIdsOk returns a tuple with the FollowerIds field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetFollowerIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowerIds, true
}

// SetFollowerIds sets field value
func (o *IterationSlim) SetFollowerIds(v []string) {
	o.FollowerIds = v
}

// GetGroupIds returns the GroupIds field value
func (o *IterationSlim) GetGroupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetGroupIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// SetGroupIds sets field value
func (o *IterationSlim) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetStartDate returns the StartDate field value
func (o *IterationSlim) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *IterationSlim) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetStatus returns the Status field value
func (o *IterationSlim) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *IterationSlim) SetStatus(v string) {
	o.Status = v
}

// GetId returns the Id field value
func (o *IterationSlim) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IterationSlim) SetId(v int64) {
	o.Id = v
}

// GetStats returns the Stats field value
func (o *IterationSlim) GetStats() IterationStats {
	if o == nil {
		var ret IterationStats
		return ret
	}

	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetStatsOk() (*IterationStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value
func (o *IterationSlim) SetStats(v IterationStats) {
	o.Stats = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *IterationSlim) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IterationSlim) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *IterationSlim) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o IterationSlim) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IterationSlim) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_url"] = o.AppUrl
	toSerialize["entity_type"] = o.EntityType
	toSerialize["labels"] = o.Labels
	toSerialize["mention_ids"] = o.MentionIds
	toSerialize["member_mention_ids"] = o.MemberMentionIds
	toSerialize["associated_groups"] = o.AssociatedGroups
	toSerialize["name"] = o.Name
	toSerialize["global_id"] = o.GlobalId
	toSerialize["label_ids"] = o.LabelIds
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["group_mention_ids"] = o.GroupMentionIds
	toSerialize["end_date"] = o.EndDate
	toSerialize["follower_ids"] = o.FollowerIds
	toSerialize["group_ids"] = o.GroupIds
	toSerialize["start_date"] = o.StartDate
	toSerialize["status"] = o.Status
	toSerialize["id"] = o.Id
	toSerialize["stats"] = o.Stats
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

type NullableIterationSlim struct {
	value *IterationSlim
	isSet bool
}

func (v NullableIterationSlim) Get() *IterationSlim {
	return v.value
}

func (v *NullableIterationSlim) Set(val *IterationSlim) {
	v.value = val
	v.isSet = true
}

func (v NullableIterationSlim) IsSet() bool {
	return v.isSet
}

func (v *NullableIterationSlim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIterationSlim(val *IterationSlim) *NullableIterationSlim {
	return &NullableIterationSlim{value: val, isSet: true}
}

func (v NullableIterationSlim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIterationSlim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


