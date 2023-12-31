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

// checks if the PullRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PullRequest{}

// PullRequest Corresponds to a VCS Pull Request attached to a Shortcut story.
type PullRequest struct {
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// True/False boolean indicating whether the VCS pull request has been closed.
	Closed bool `json:"closed"`
	// True/False boolean indicating whether the VCS pull request has been merged.
	Merged bool `json:"merged"`
	// Number of lines added in the pull request, according to VCS.
	NumAdded int64 `json:"num_added"`
	// The ID of the branch for the particular pull request.
	BranchId int64 `json:"branch_id"`
	// An array of Story ids that have Pull Requests that change at least one of the same lines this Pull Request changes.
	OverlappingStories []int64 `json:"overlapping_stories,omitempty"`
	// The pull request's unique number ID in VCS.
	Number int64 `json:"number"`
	// The name of the branch for the particular pull request.
	BranchName string `json:"branch_name"`
	// The name of the target branch for the particular pull request.
	TargetBranchName string `json:"target_branch_name"`
	// The number of commits on the pull request.
	NumCommits NullableInt64 `json:"num_commits"`
	// The title of the pull request.
	Title string `json:"title"`
	// The time/date the pull request was created.
	UpdatedAt time.Time `json:"updated_at"`

	// True/False boolean indicating whether the VCS pull request is in the draft state.
	Draft bool `json:"draft"`
	// The unique ID associated with the pull request in Shortcut.
	Id int64 `json:"id"`
	// An array of PullRequestLabels attached to the PullRequest.
	VcsLabels []PullRequestLabel `json:"vcs_labels,omitempty"`
	// The URL for the pull request.
	Url string `json:"url"`
	// Number of lines removed in the pull request, according to VCS.
	NumRemoved int64 `json:"num_removed"`
	// The status of the review for the pull request.
	ReviewStatus *string `json:"review_status,omitempty"`
	// Number of lines modified in the pull request, according to VCS.
	NumModified NullableInt64 `json:"num_modified"`
	// The status of the Continuous Integration workflow for the pull request.
	BuildStatus *string `json:"build_status,omitempty"`
	// The ID of the target branch for the particular pull request.
	TargetBranchId int64 `json:"target_branch_id"`
	// The ID of the repository for the particular pull request.
	RepositoryId int64 `json:"repository_id"`
	// The time/date the pull request was created.
	CreatedAt time.Time `json:"created_at"`
}

// NewPullRequest instantiates a new PullRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequest(entityType string, closed bool, merged bool, numAdded int64, branchId int64, number int64, branchName string, targetBranchName string, numCommits NullableInt64, title string, updatedAt time.Time,  draft bool, id int64, url string, numRemoved int64, numModified NullableInt64, targetBranchId int64, repositoryId int64, createdAt time.Time) *PullRequest {
	this := PullRequest{}
	this.EntityType = entityType
	this.Closed = closed
	this.Merged = merged
	this.NumAdded = numAdded
	this.BranchId = branchId
	this.Number = number
	this.BranchName = branchName
	this.TargetBranchName = targetBranchName
	this.NumCommits = numCommits
	this.Title = title
	this.UpdatedAt = updatedAt
	this.Draft = draft
	this.Id = id
	this.Url = url
	this.NumRemoved = numRemoved
	this.NumModified = numModified
	this.TargetBranchId = targetBranchId
	this.RepositoryId = repositoryId
	this.CreatedAt = createdAt
	return &this
}

// NewPullRequestWithDefaults instantiates a new PullRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestWithDefaults() *PullRequest {
	this := PullRequest{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *PullRequest) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *PullRequest) SetEntityType(v string) {
	o.EntityType = v
}

// GetClosed returns the Closed field value
func (o *PullRequest) GetClosed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Closed
}

// GetClosedOk returns a tuple with the Closed field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetClosedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Closed, true
}

// SetClosed sets field value
func (o *PullRequest) SetClosed(v bool) {
	o.Closed = v
}

// GetMerged returns the Merged field value
func (o *PullRequest) GetMerged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Merged
}

// GetMergedOk returns a tuple with the Merged field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetMergedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Merged, true
}

// SetMerged sets field value
func (o *PullRequest) SetMerged(v bool) {
	o.Merged = v
}

// GetNumAdded returns the NumAdded field value
func (o *PullRequest) GetNumAdded() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumAdded
}

// GetNumAddedOk returns a tuple with the NumAdded field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetNumAddedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumAdded, true
}

// SetNumAdded sets field value
func (o *PullRequest) SetNumAdded(v int64) {
	o.NumAdded = v
}

// GetBranchId returns the BranchId field value
func (o *PullRequest) GetBranchId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BranchId
}

// GetBranchIdOk returns a tuple with the BranchId field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetBranchIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchId, true
}

// SetBranchId sets field value
func (o *PullRequest) SetBranchId(v int64) {
	o.BranchId = v
}

// GetOverlappingStories returns the OverlappingStories field value if set, zero value otherwise.
func (o *PullRequest) GetOverlappingStories() []int64 {
	if o == nil || IsNil(o.OverlappingStories) {
		var ret []int64
		return ret
	}
	return o.OverlappingStories
}

// GetOverlappingStoriesOk returns a tuple with the OverlappingStories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetOverlappingStoriesOk() ([]int64, bool) {
	if o == nil || IsNil(o.OverlappingStories) {
		return nil, false
	}
	return o.OverlappingStories, true
}

// HasOverlappingStories returns a boolean if a field has been set.
func (o *PullRequest) HasOverlappingStories() bool {
	if o != nil && !IsNil(o.OverlappingStories) {
		return true
	}

	return false
}

// SetOverlappingStories gets a reference to the given []int64 and assigns it to the OverlappingStories field.
func (o *PullRequest) SetOverlappingStories(v []int64) {
	o.OverlappingStories = v
}

// GetNumber returns the Number field value
func (o *PullRequest) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *PullRequest) SetNumber(v int64) {
	o.Number = v
}

// GetBranchName returns the BranchName field value
func (o *PullRequest) GetBranchName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchName, true
}

// SetBranchName sets field value
func (o *PullRequest) SetBranchName(v string) {
	o.BranchName = v
}

// GetTargetBranchName returns the TargetBranchName field value
func (o *PullRequest) GetTargetBranchName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetBranchName
}

// GetTargetBranchNameOk returns a tuple with the TargetBranchName field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetTargetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetBranchName, true
}

// SetTargetBranchName sets field value
func (o *PullRequest) SetTargetBranchName(v string) {
	o.TargetBranchName = v
}

// GetNumCommits returns the NumCommits field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *PullRequest) GetNumCommits() int64 {
	if o == nil || o.NumCommits.Get() == nil {
		var ret int64
		return ret
	}

	return *o.NumCommits.Get()
}

// GetNumCommitsOk returns a tuple with the NumCommits field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequest) GetNumCommitsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumCommits.Get(), o.NumCommits.IsSet()
}

// SetNumCommits sets field value
func (o *PullRequest) SetNumCommits(v int64) {
	o.NumCommits.Set(&v)
}

// GetTitle returns the Title field value
func (o *PullRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PullRequest) SetTitle(v string) {
	o.Title = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PullRequest) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PullRequest) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetHasOverlappingStories returns the HasOverlappingStories field value
func (o *PullRequest) GetHasOverlappingStories() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasOverlappingStories()
}


// GetDraft returns the Draft field value
func (o *PullRequest) GetDraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Draft
}

// GetDraftOk returns a tuple with the Draft field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetDraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draft, true
}

// SetDraft sets field value
func (o *PullRequest) SetDraft(v bool) {
	o.Draft = v
}

// GetId returns the Id field value
func (o *PullRequest) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PullRequest) SetId(v int64) {
	o.Id = v
}

// GetVcsLabels returns the VcsLabels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PullRequest) GetVcsLabels() []PullRequestLabel {
	if o == nil {
		var ret []PullRequestLabel
		return ret
	}
	return o.VcsLabels
}

// GetVcsLabelsOk returns a tuple with the VcsLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequest) GetVcsLabelsOk() ([]PullRequestLabel, bool) {
	if o == nil || IsNil(o.VcsLabels) {
		return nil, false
	}
	return o.VcsLabels, true
}

// HasVcsLabels returns a boolean if a field has been set.
func (o *PullRequest) HasVcsLabels() bool {
	if o != nil && IsNil(o.VcsLabels) {
		return true
	}

	return false
}

// SetVcsLabels gets a reference to the given []PullRequestLabel and assigns it to the VcsLabels field.
func (o *PullRequest) SetVcsLabels(v []PullRequestLabel) {
	o.VcsLabels = v
}

// GetUrl returns the Url field value
func (o *PullRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PullRequest) SetUrl(v string) {
	o.Url = v
}

// GetNumRemoved returns the NumRemoved field value
func (o *PullRequest) GetNumRemoved() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumRemoved
}

// GetNumRemovedOk returns a tuple with the NumRemoved field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetNumRemovedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumRemoved, true
}

// SetNumRemoved sets field value
func (o *PullRequest) SetNumRemoved(v int64) {
	o.NumRemoved = v
}

// GetReviewStatus returns the ReviewStatus field value if set, zero value otherwise.
func (o *PullRequest) GetReviewStatus() string {
	if o == nil || IsNil(o.ReviewStatus) {
		var ret string
		return ret
	}
	return *o.ReviewStatus
}

// GetReviewStatusOk returns a tuple with the ReviewStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetReviewStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewStatus) {
		return nil, false
	}
	return o.ReviewStatus, true
}

// HasReviewStatus returns a boolean if a field has been set.
func (o *PullRequest) HasReviewStatus() bool {
	if o != nil && !IsNil(o.ReviewStatus) {
		return true
	}

	return false
}

// SetReviewStatus gets a reference to the given string and assigns it to the ReviewStatus field.
func (o *PullRequest) SetReviewStatus(v string) {
	o.ReviewStatus = &v
}

// GetNumModified returns the NumModified field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *PullRequest) GetNumModified() int64 {
	if o == nil || o.NumModified.Get() == nil {
		var ret int64
		return ret
	}

	return *o.NumModified.Get()
}

// GetNumModifiedOk returns a tuple with the NumModified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequest) GetNumModifiedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumModified.Get(), o.NumModified.IsSet()
}

// SetNumModified sets field value
func (o *PullRequest) SetNumModified(v int64) {
	o.NumModified.Set(&v)
}

// GetBuildStatus returns the BuildStatus field value if set, zero value otherwise.
func (o *PullRequest) GetBuildStatus() string {
	if o == nil || IsNil(o.BuildStatus) {
		var ret string
		return ret
	}
	return *o.BuildStatus
}

// GetBuildStatusOk returns a tuple with the BuildStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequest) GetBuildStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BuildStatus) {
		return nil, false
	}
	return o.BuildStatus, true
}

// HasBuildStatus returns a boolean if a field has been set.
func (o *PullRequest) HasBuildStatus() bool {
	if o != nil && !IsNil(o.BuildStatus) {
		return true
	}

	return false
}

// SetBuildStatus gets a reference to the given string and assigns it to the BuildStatus field.
func (o *PullRequest) SetBuildStatus(v string) {
	o.BuildStatus = &v
}

// GetTargetBranchId returns the TargetBranchId field value
func (o *PullRequest) GetTargetBranchId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TargetBranchId
}

// GetTargetBranchIdOk returns a tuple with the TargetBranchId field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetTargetBranchIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetBranchId, true
}

// SetTargetBranchId sets field value
func (o *PullRequest) SetTargetBranchId(v int64) {
	o.TargetBranchId = v
}

// GetRepositoryId returns the RepositoryId field value
func (o *PullRequest) GetRepositoryId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetRepositoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryId, true
}

// SetRepositoryId sets field value
func (o *PullRequest) SetRepositoryId(v int64) {
	o.RepositoryId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PullRequest) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PullRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PullRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o PullRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PullRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_type"] = o.EntityType
	toSerialize["closed"] = o.Closed
	toSerialize["merged"] = o.Merged
	toSerialize["num_added"] = o.NumAdded
	toSerialize["branch_id"] = o.BranchId
	if !IsNil(o.OverlappingStories) {
		toSerialize["overlapping_stories"] = o.OverlappingStories
	}
	toSerialize["number"] = o.Number
	toSerialize["branch_name"] = o.BranchName
	toSerialize["target_branch_name"] = o.TargetBranchName
	toSerialize["num_commits"] = o.NumCommits.Get()
	toSerialize["title"] = o.Title
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["has_overlapping_stories"] = o.HasOverlappingStories
	toSerialize["draft"] = o.Draft
	toSerialize["id"] = o.Id
	if o.VcsLabels != nil {
		toSerialize["vcs_labels"] = o.VcsLabels
	}
	toSerialize["url"] = o.Url
	toSerialize["num_removed"] = o.NumRemoved
	if !IsNil(o.ReviewStatus) {
		toSerialize["review_status"] = o.ReviewStatus
	}
	toSerialize["num_modified"] = o.NumModified.Get()
	if !IsNil(o.BuildStatus) {
		toSerialize["build_status"] = o.BuildStatus
	}
	toSerialize["target_branch_id"] = o.TargetBranchId
	toSerialize["repository_id"] = o.RepositoryId
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

type NullablePullRequest struct {
	value *PullRequest
	isSet bool
}

func (v NullablePullRequest) Get() *PullRequest {
	return v.value
}

func (v *NullablePullRequest) Set(val *PullRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequest(val *PullRequest) *NullablePullRequest {
	return &NullablePullRequest{value: val, isSet: true}
}

func (v NullablePullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
