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

// checks if the Branch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Branch{}

// Branch Branch refers to a VCS branch. Branches are feature branches associated with Shortcut Stories.
type Branch struct {
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// A true/false boolean indicating if the Branch has been deleted.
	Deleted bool `json:"deleted"`
	// The name of the Branch.
	Name string `json:"name"`
	// A true/false boolean indicating if the Branch is persistent; e.g. master.
	Persistent bool `json:"persistent"`
	// The time/date the Branch was updated.
	UpdatedAt NullableTime `json:"updated_at"`
	// An array of PullRequests attached to the Branch (there is usually only one).
	PullRequests []PullRequest `json:"pull_requests"`
	// The IDs of the Branches the Branch has been merged into.
	MergedBranchIds []int64 `json:"merged_branch_ids"`
	// The unique ID of the Branch.
	Id NullableInt64 `json:"id"`
	// The URL of the Branch.
	Url string `json:"url"`
	// The ID of the Repository that contains the Branch.
	RepositoryId int64 `json:"repository_id"`
	// The time/date the Branch was created.
	CreatedAt NullableTime `json:"created_at"`
}

// NewBranch instantiates a new Branch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranch(entityType string, deleted bool, name string, persistent bool, updatedAt NullableTime, pullRequests []PullRequest, mergedBranchIds []int64, id NullableInt64, url string, repositoryId int64, createdAt NullableTime) *Branch {
	this := Branch{}
	this.EntityType = entityType
	this.Deleted = deleted
	this.Name = name
	this.Persistent = persistent
	this.UpdatedAt = updatedAt
	this.PullRequests = pullRequests
	this.MergedBranchIds = mergedBranchIds
	this.Id = id
	this.Url = url
	this.RepositoryId = repositoryId
	this.CreatedAt = createdAt
	return &this
}

// NewBranchWithDefaults instantiates a new Branch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchWithDefaults() *Branch {
	this := Branch{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *Branch) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *Branch) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *Branch) SetEntityType(v string) {
	o.EntityType = v
}

// GetDeleted returns the Deleted field value
func (o *Branch) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *Branch) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *Branch) SetDeleted(v bool) {
	o.Deleted = v
}

// GetName returns the Name field value
func (o *Branch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Branch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Branch) SetName(v string) {
	o.Name = v
}

// GetPersistent returns the Persistent field value
func (o *Branch) GetPersistent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Persistent
}

// GetPersistentOk returns a tuple with the Persistent field value
// and a boolean to check if the value has been set.
func (o *Branch) GetPersistentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Persistent, true
}

// SetPersistent sets field value
func (o *Branch) SetPersistent(v bool) {
	o.Persistent = v
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Branch) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Branch) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *Branch) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// GetPullRequests returns the PullRequests field value
func (o *Branch) GetPullRequests() []PullRequest {
	if o == nil {
		var ret []PullRequest
		return ret
	}

	return o.PullRequests
}

// GetPullRequestsOk returns a tuple with the PullRequests field value
// and a boolean to check if the value has been set.
func (o *Branch) GetPullRequestsOk() ([]PullRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.PullRequests, true
}

// SetPullRequests sets field value
func (o *Branch) SetPullRequests(v []PullRequest) {
	o.PullRequests = v
}

// GetMergedBranchIds returns the MergedBranchIds field value
func (o *Branch) GetMergedBranchIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.MergedBranchIds
}

// GetMergedBranchIdsOk returns a tuple with the MergedBranchIds field value
// and a boolean to check if the value has been set.
func (o *Branch) GetMergedBranchIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MergedBranchIds, true
}

// SetMergedBranchIds sets field value
func (o *Branch) SetMergedBranchIds(v []int64) {
	o.MergedBranchIds = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *Branch) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Branch) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *Branch) SetId(v int64) {
	o.Id.Set(&v)
}

// GetUrl returns the Url field value
func (o *Branch) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Branch) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Branch) SetUrl(v string) {
	o.Url = v
}

// GetRepositoryId returns the RepositoryId field value
func (o *Branch) GetRepositoryId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
func (o *Branch) GetRepositoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryId, true
}

// SetRepositoryId sets field value
func (o *Branch) SetRepositoryId(v int64) {
	o.RepositoryId = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Branch) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Branch) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *Branch) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

func (o Branch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Branch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_type"] = o.EntityType
	toSerialize["deleted"] = o.Deleted
	toSerialize["name"] = o.Name
	toSerialize["persistent"] = o.Persistent
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	toSerialize["pull_requests"] = o.PullRequests
	toSerialize["merged_branch_ids"] = o.MergedBranchIds
	toSerialize["id"] = o.Id.Get()
	toSerialize["url"] = o.Url
	toSerialize["repository_id"] = o.RepositoryId
	toSerialize["created_at"] = o.CreatedAt.Get()
	return toSerialize, nil
}

type NullableBranch struct {
	value *Branch
	isSet bool
}

func (v NullableBranch) Get() *Branch {
	return v.value
}

func (v *NullableBranch) Set(val *Branch) {
	v.value = val
	v.isSet = true
}

func (v NullableBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranch(val *Branch) *NullableBranch {
	return &NullableBranch{value: val, isSet: true}
}

func (v NullableBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


