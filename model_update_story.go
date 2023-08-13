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

// checks if the UpdateStory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStory{}

// UpdateStory struct for UpdateStory
type UpdateStory struct {
	// The description of the story.
	Description *string `json:"description,omitempty"`
	// True if the story is archived, otherwise false.
	Archived *bool `json:"archived,omitempty"`
	// An array of labels attached to the story.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// An array of IDs of Pull/Merge Requests attached to the story.
	PullRequestIds []int64 `json:"pull_request_ids,omitempty"`
	// The type of story (feature, bug, chore).
	StoryType *string `json:"story_type,omitempty"`
	// A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField.
	CustomFields []CustomFieldValueParams `json:"custom_fields,omitempty"`
	// One of \"first\" or \"last\". This can be used to move the given story to the first or last position in the workflow state.
	MoveTo *string `json:"move_to,omitempty"`
	// An array of IDs of files attached to the story.
	FileIds []int64 `json:"file_ids,omitempty"`
	// A manual override for the time/date the Story was completed.
	CompletedAtOverride NullableTime `json:"completed_at_override,omitempty"`
	// The title of the story.
	Name *string `json:"name,omitempty"`
	// The ID of the epic the story belongs to.
	EpicId NullableInt64 `json:"epic_id,omitempty"`
	// An array of External Links associated with this story.
	ExternalLinks []string `json:"external_links,omitempty"`
	// An array of IDs of Branches attached to the story.
	BranchIds []int64 `json:"branch_ids,omitempty"`
	// An array of IDs of Commits attached to the story.
	CommitIds []int64 `json:"commit_ids,omitempty"`
	// The ID of the member that requested the story.
	RequestedById *string `json:"requested_by_id,omitempty"`
	// The ID of the iteration the story belongs to.
	IterationId NullableInt64 `json:"iteration_id,omitempty"`
	// A manual override for the time/date the Story was started.
	StartedAtOverride NullableTime `json:"started_at_override,omitempty"`
	// The ID of the group to associate with this story
	GroupId NullableString `json:"group_id,omitempty"`
	// The ID of the workflow state to put the story in.
	WorkflowStateId *int64 `json:"workflow_state_id,omitempty"`
	// An array of UUIDs of the followers of this story.
	FollowerIds []string `json:"follower_ids,omitempty"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids,omitempty"`
	// The ID of the story we want to move this story before.
	BeforeId *int64 `json:"before_id,omitempty"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate NullableInt64 `json:"estimate,omitempty"`
	// The ID of the story we want to move this story after.
	AfterId *int64 `json:"after_id,omitempty"`
	// The ID of the project the story belongs to.
	ProjectId NullableInt64 `json:"project_id,omitempty"`
	// An array of IDs of linked files attached to the story.
	LinkedFileIds []int64 `json:"linked_file_ids,omitempty"`
	// The due date of the story.
	Deadline NullableTime `json:"deadline,omitempty"`
}

// NewUpdateStory instantiates a new UpdateStory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStory() *UpdateStory {
	this := UpdateStory{}
	return &this
}

// NewUpdateStoryWithDefaults instantiates a new UpdateStory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStoryWithDefaults() *UpdateStory {
	this := UpdateStory{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateStory) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateStory) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateStory) SetDescription(v string) {
	o.Description = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *UpdateStory) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *UpdateStory) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *UpdateStory) SetArchived(v bool) {
	o.Archived = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateStory) GetLabels() []CreateLabelParams {
	if o == nil || IsNil(o.Labels) {
		var ret []CreateLabelParams
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetLabelsOk() ([]CreateLabelParams, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateStory) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []CreateLabelParams and assigns it to the Labels field.
func (o *UpdateStory) SetLabels(v []CreateLabelParams) {
	o.Labels = v
}

// GetPullRequestIds returns the PullRequestIds field value if set, zero value otherwise.
func (o *UpdateStory) GetPullRequestIds() []int64 {
	if o == nil || IsNil(o.PullRequestIds) {
		var ret []int64
		return ret
	}
	return o.PullRequestIds
}

// GetPullRequestIdsOk returns a tuple with the PullRequestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetPullRequestIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.PullRequestIds) {
		return nil, false
	}
	return o.PullRequestIds, true
}

// HasPullRequestIds returns a boolean if a field has been set.
func (o *UpdateStory) HasPullRequestIds() bool {
	if o != nil && !IsNil(o.PullRequestIds) {
		return true
	}

	return false
}

// SetPullRequestIds gets a reference to the given []int64 and assigns it to the PullRequestIds field.
func (o *UpdateStory) SetPullRequestIds(v []int64) {
	o.PullRequestIds = v
}

// GetStoryType returns the StoryType field value if set, zero value otherwise.
func (o *UpdateStory) GetStoryType() string {
	if o == nil || IsNil(o.StoryType) {
		var ret string
		return ret
	}
	return *o.StoryType
}

// GetStoryTypeOk returns a tuple with the StoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetStoryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StoryType) {
		return nil, false
	}
	return o.StoryType, true
}

// HasStoryType returns a boolean if a field has been set.
func (o *UpdateStory) HasStoryType() bool {
	if o != nil && !IsNil(o.StoryType) {
		return true
	}

	return false
}

// SetStoryType gets a reference to the given string and assigns it to the StoryType field.
func (o *UpdateStory) SetStoryType(v string) {
	o.StoryType = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *UpdateStory) GetCustomFields() []CustomFieldValueParams {
	if o == nil || IsNil(o.CustomFields) {
		var ret []CustomFieldValueParams
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetCustomFieldsOk() ([]CustomFieldValueParams, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *UpdateStory) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []CustomFieldValueParams and assigns it to the CustomFields field.
func (o *UpdateStory) SetCustomFields(v []CustomFieldValueParams) {
	o.CustomFields = v
}

// GetMoveTo returns the MoveTo field value if set, zero value otherwise.
func (o *UpdateStory) GetMoveTo() string {
	if o == nil || IsNil(o.MoveTo) {
		var ret string
		return ret
	}
	return *o.MoveTo
}

// GetMoveToOk returns a tuple with the MoveTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetMoveToOk() (*string, bool) {
	if o == nil || IsNil(o.MoveTo) {
		return nil, false
	}
	return o.MoveTo, true
}

// HasMoveTo returns a boolean if a field has been set.
func (o *UpdateStory) HasMoveTo() bool {
	if o != nil && !IsNil(o.MoveTo) {
		return true
	}

	return false
}

// SetMoveTo gets a reference to the given string and assigns it to the MoveTo field.
func (o *UpdateStory) SetMoveTo(v string) {
	o.MoveTo = &v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise.
func (o *UpdateStory) GetFileIds() []int64 {
	if o == nil || IsNil(o.FileIds) {
		var ret []int64
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetFileIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *UpdateStory) HasFileIds() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []int64 and assigns it to the FileIds field.
func (o *UpdateStory) SetFileIds(v []int64) {
	o.FileIds = v
}

// GetCompletedAtOverride returns the CompletedAtOverride field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStory) GetCompletedAtOverride() time.Time {
	if o == nil || IsNil(o.CompletedAtOverride.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAtOverride.Get()
}

// GetCompletedAtOverrideOk returns a tuple with the CompletedAtOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStory) GetCompletedAtOverrideOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAtOverride.Get(), o.CompletedAtOverride.IsSet()
}

// HasCompletedAtOverride returns a boolean if a field has been set.
func (o *UpdateStory) HasCompletedAtOverride() bool {
	if o != nil && o.CompletedAtOverride.IsSet() {
		return true
	}

	return false
}

// SetCompletedAtOverride gets a reference to the given NullableTime and assigns it to the CompletedAtOverride field.
func (o *UpdateStory) SetCompletedAtOverride(v time.Time) {
	o.CompletedAtOverride.Set(&v)
}
// SetCompletedAtOverrideNil sets the value for CompletedAtOverride to be an explicit nil
func (o *UpdateStory) SetCompletedAtOverrideNil() {
	o.CompletedAtOverride.Set(nil)
}

// UnsetCompletedAtOverride ensures that no value is present for CompletedAtOverride, not even an explicit nil
func (o *UpdateStory) UnsetCompletedAtOverride() {
	o.CompletedAtOverride.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateStory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateStory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateStory) SetName(v string) {
	o.Name = &v
}

// GetEpicId returns the EpicId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStory) GetEpicId() int64 {
	if o == nil || IsNil(o.EpicId.Get()) {
		var ret int64
		return ret
	}
	return *o.EpicId.Get()
}

// GetEpicIdOk returns a tuple with the EpicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStory) GetEpicIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EpicId.Get(), o.EpicId.IsSet()
}

// HasEpicId returns a boolean if a field has been set.
func (o *UpdateStory) HasEpicId() bool {
	if o != nil && o.EpicId.IsSet() {
		return true
	}

	return false
}

// SetEpicId gets a reference to the given NullableInt64 and assigns it to the EpicId field.
func (o *UpdateStory) SetEpicId(v int64) {
	o.EpicId.Set(&v)
}
// SetEpicIdNil sets the value for EpicId to be an explicit nil
func (o *UpdateStory) SetEpicIdNil() {
	o.EpicId.Set(nil)
}

// UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
func (o *UpdateStory) UnsetEpicId() {
	o.EpicId.Unset()
}

// GetExternalLinks returns the ExternalLinks field value if set, zero value otherwise.
func (o *UpdateStory) GetExternalLinks() []string {
	if o == nil || IsNil(o.ExternalLinks) {
		var ret []string
		return ret
	}
	return o.ExternalLinks
}

// GetExternalLinksOk returns a tuple with the ExternalLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetExternalLinksOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalLinks) {
		return nil, false
	}
	return o.ExternalLinks, true
}

// HasExternalLinks returns a boolean if a field has been set.
func (o *UpdateStory) HasExternalLinks() bool {
	if o != nil && !IsNil(o.ExternalLinks) {
		return true
	}

	return false
}

// SetExternalLinks gets a reference to the given []string and assigns it to the ExternalLinks field.
func (o *UpdateStory) SetExternalLinks(v []string) {
	o.ExternalLinks = v
}

// GetBranchIds returns the BranchIds field value if set, zero value otherwise.
func (o *UpdateStory) GetBranchIds() []int64 {
	if o == nil || IsNil(o.BranchIds) {
		var ret []int64
		return ret
	}
	return o.BranchIds
}

// GetBranchIdsOk returns a tuple with the BranchIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetBranchIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.BranchIds) {
		return nil, false
	}
	return o.BranchIds, true
}

// HasBranchIds returns a boolean if a field has been set.
func (o *UpdateStory) HasBranchIds() bool {
	if o != nil && !IsNil(o.BranchIds) {
		return true
	}

	return false
}

// SetBranchIds gets a reference to the given []int64 and assigns it to the BranchIds field.
func (o *UpdateStory) SetBranchIds(v []int64) {
	o.BranchIds = v
}

// GetCommitIds returns the CommitIds field value if set, zero value otherwise.
func (o *UpdateStory) GetCommitIds() []int64 {
	if o == nil || IsNil(o.CommitIds) {
		var ret []int64
		return ret
	}
	return o.CommitIds
}

// GetCommitIdsOk returns a tuple with the CommitIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetCommitIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.CommitIds) {
		return nil, false
	}
	return o.CommitIds, true
}

// HasCommitIds returns a boolean if a field has been set.
func (o *UpdateStory) HasCommitIds() bool {
	if o != nil && !IsNil(o.CommitIds) {
		return true
	}

	return false
}

// SetCommitIds gets a reference to the given []int64 and assigns it to the CommitIds field.
func (o *UpdateStory) SetCommitIds(v []int64) {
	o.CommitIds = v
}

// GetRequestedById returns the RequestedById field value if set, zero value otherwise.
func (o *UpdateStory) GetRequestedById() string {
	if o == nil || IsNil(o.RequestedById) {
		var ret string
		return ret
	}
	return *o.RequestedById
}

// GetRequestedByIdOk returns a tuple with the RequestedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetRequestedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedById) {
		return nil, false
	}
	return o.RequestedById, true
}

// HasRequestedById returns a boolean if a field has been set.
func (o *UpdateStory) HasRequestedById() bool {
	if o != nil && !IsNil(o.RequestedById) {
		return true
	}

	return false
}

// SetRequestedById gets a reference to the given string and assigns it to the RequestedById field.
func (o *UpdateStory) SetRequestedById(v string) {
	o.RequestedById = &v
}

// GetIterationId returns the IterationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStory) GetIterationId() int64 {
	if o == nil || IsNil(o.IterationId.Get()) {
		var ret int64
		return ret
	}
	return *o.IterationId.Get()
}

// GetIterationIdOk returns a tuple with the IterationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStory) GetIterationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IterationId.Get(), o.IterationId.IsSet()
}

// HasIterationId returns a boolean if a field has been set.
func (o *UpdateStory) HasIterationId() bool {
	if o != nil && o.IterationId.IsSet() {
		return true
	}

	return false
}

// SetIterationId gets a reference to the given NullableInt64 and assigns it to the IterationId field.
func (o *UpdateStory) SetIterationId(v int64) {
	o.IterationId.Set(&v)
}
// SetIterationIdNil sets the value for IterationId to be an explicit nil
func (o *UpdateStory) SetIterationIdNil() {
	o.IterationId.Set(nil)
}

// UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
func (o *UpdateStory) UnsetIterationId() {
	o.IterationId.Unset()
}

// GetStartedAtOverride returns the StartedAtOverride field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStory) GetStartedAtOverride() time.Time {
	if o == nil || IsNil(o.StartedAtOverride.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartedAtOverride.Get()
}

// GetStartedAtOverrideOk returns a tuple with the StartedAtOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStory) GetStartedAtOverrideOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAtOverride.Get(), o.StartedAtOverride.IsSet()
}

// HasStartedAtOverride returns a boolean if a field has been set.
func (o *UpdateStory) HasStartedAtOverride() bool {
	if o != nil && o.StartedAtOverride.IsSet() {
		return true
	}

	return false
}

// SetStartedAtOverride gets a reference to the given NullableTime and assigns it to the StartedAtOverride field.
func (o *UpdateStory) SetStartedAtOverride(v time.Time) {
	o.StartedAtOverride.Set(&v)
}
// SetStartedAtOverrideNil sets the value for StartedAtOverride to be an explicit nil
func (o *UpdateStory) SetStartedAtOverrideNil() {
	o.StartedAtOverride.Set(nil)
}

// UnsetStartedAtOverride ensures that no value is present for StartedAtOverride, not even an explicit nil
func (o *UpdateStory) UnsetStartedAtOverride() {
	o.StartedAtOverride.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStory) GetGroupId() string {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret string
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStory) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *UpdateStory) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableString and assigns it to the GroupId field.
func (o *UpdateStory) SetGroupId(v string) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *UpdateStory) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *UpdateStory) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetWorkflowStateId returns the WorkflowStateId field value if set, zero value otherwise.
func (o *UpdateStory) GetWorkflowStateId() int64 {
	if o == nil || IsNil(o.WorkflowStateId) {
		var ret int64
		return ret
	}
	return *o.WorkflowStateId
}

// GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetWorkflowStateIdOk() (*int64, bool) {
	if o == nil || IsNil(o.WorkflowStateId) {
		return nil, false
	}
	return o.WorkflowStateId, true
}

// HasWorkflowStateId returns a boolean if a field has been set.
func (o *UpdateStory) HasWorkflowStateId() bool {
	if o != nil && !IsNil(o.WorkflowStateId) {
		return true
	}

	return false
}

// SetWorkflowStateId gets a reference to the given int64 and assigns it to the WorkflowStateId field.
func (o *UpdateStory) SetWorkflowStateId(v int64) {
	o.WorkflowStateId = &v
}

// GetFollowerIds returns the FollowerIds field value if set, zero value otherwise.
func (o *UpdateStory) GetFollowerIds() []string {
	if o == nil || IsNil(o.FollowerIds) {
		var ret []string
		return ret
	}
	return o.FollowerIds
}

// GetFollowerIdsOk returns a tuple with the FollowerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetFollowerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FollowerIds) {
		return nil, false
	}
	return o.FollowerIds, true
}

// HasFollowerIds returns a boolean if a field has been set.
func (o *UpdateStory) HasFollowerIds() bool {
	if o != nil && !IsNil(o.FollowerIds) {
		return true
	}

	return false
}

// SetFollowerIds gets a reference to the given []string and assigns it to the FollowerIds field.
func (o *UpdateStory) SetFollowerIds(v []string) {
	o.FollowerIds = v
}

// GetOwnerIds returns the OwnerIds field value if set, zero value otherwise.
func (o *UpdateStory) GetOwnerIds() []string {
	if o == nil || IsNil(o.OwnerIds) {
		var ret []string
		return ret
	}
	return o.OwnerIds
}

// GetOwnerIdsOk returns a tuple with the OwnerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetOwnerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerIds) {
		return nil, false
	}
	return o.OwnerIds, true
}

// HasOwnerIds returns a boolean if a field has been set.
func (o *UpdateStory) HasOwnerIds() bool {
	if o != nil && !IsNil(o.OwnerIds) {
		return true
	}

	return false
}

// SetOwnerIds gets a reference to the given []string and assigns it to the OwnerIds field.
func (o *UpdateStory) SetOwnerIds(v []string) {
	o.OwnerIds = v
}

// GetBeforeId returns the BeforeId field value if set, zero value otherwise.
func (o *UpdateStory) GetBeforeId() int64 {
	if o == nil || IsNil(o.BeforeId) {
		var ret int64
		return ret
	}
	return *o.BeforeId
}

// GetBeforeIdOk returns a tuple with the BeforeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetBeforeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BeforeId) {
		return nil, false
	}
	return o.BeforeId, true
}

// HasBeforeId returns a boolean if a field has been set.
func (o *UpdateStory) HasBeforeId() bool {
	if o != nil && !IsNil(o.BeforeId) {
		return true
	}

	return false
}

// SetBeforeId gets a reference to the given int64 and assigns it to the BeforeId field.
func (o *UpdateStory) SetBeforeId(v int64) {
	o.BeforeId = &v
}

// GetEstimate returns the Estimate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStory) GetEstimate() int64 {
	if o == nil || IsNil(o.Estimate.Get()) {
		var ret int64
		return ret
	}
	return *o.Estimate.Get()
}

// GetEstimateOk returns a tuple with the Estimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStory) GetEstimateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Estimate.Get(), o.Estimate.IsSet()
}

// HasEstimate returns a boolean if a field has been set.
func (o *UpdateStory) HasEstimate() bool {
	if o != nil && o.Estimate.IsSet() {
		return true
	}

	return false
}

// SetEstimate gets a reference to the given NullableInt64 and assigns it to the Estimate field.
func (o *UpdateStory) SetEstimate(v int64) {
	o.Estimate.Set(&v)
}
// SetEstimateNil sets the value for Estimate to be an explicit nil
func (o *UpdateStory) SetEstimateNil() {
	o.Estimate.Set(nil)
}

// UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
func (o *UpdateStory) UnsetEstimate() {
	o.Estimate.Unset()
}

// GetAfterId returns the AfterId field value if set, zero value otherwise.
func (o *UpdateStory) GetAfterId() int64 {
	if o == nil || IsNil(o.AfterId) {
		var ret int64
		return ret
	}
	return *o.AfterId
}

// GetAfterIdOk returns a tuple with the AfterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetAfterIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AfterId) {
		return nil, false
	}
	return o.AfterId, true
}

// HasAfterId returns a boolean if a field has been set.
func (o *UpdateStory) HasAfterId() bool {
	if o != nil && !IsNil(o.AfterId) {
		return true
	}

	return false
}

// SetAfterId gets a reference to the given int64 and assigns it to the AfterId field.
func (o *UpdateStory) SetAfterId(v int64) {
	o.AfterId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStory) GetProjectId() int64 {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStory) GetProjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *UpdateStory) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableInt64 and assigns it to the ProjectId field.
func (o *UpdateStory) SetProjectId(v int64) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *UpdateStory) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *UpdateStory) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetLinkedFileIds returns the LinkedFileIds field value if set, zero value otherwise.
func (o *UpdateStory) GetLinkedFileIds() []int64 {
	if o == nil || IsNil(o.LinkedFileIds) {
		var ret []int64
		return ret
	}
	return o.LinkedFileIds
}

// GetLinkedFileIdsOk returns a tuple with the LinkedFileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStory) GetLinkedFileIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.LinkedFileIds) {
		return nil, false
	}
	return o.LinkedFileIds, true
}

// HasLinkedFileIds returns a boolean if a field has been set.
func (o *UpdateStory) HasLinkedFileIds() bool {
	if o != nil && !IsNil(o.LinkedFileIds) {
		return true
	}

	return false
}

// SetLinkedFileIds gets a reference to the given []int64 and assigns it to the LinkedFileIds field.
func (o *UpdateStory) SetLinkedFileIds(v []int64) {
	o.LinkedFileIds = v
}

// GetDeadline returns the Deadline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStory) GetDeadline() time.Time {
	if o == nil || IsNil(o.Deadline.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Deadline.Get()
}

// GetDeadlineOk returns a tuple with the Deadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStory) GetDeadlineOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deadline.Get(), o.Deadline.IsSet()
}

// HasDeadline returns a boolean if a field has been set.
func (o *UpdateStory) HasDeadline() bool {
	if o != nil && o.Deadline.IsSet() {
		return true
	}

	return false
}

// SetDeadline gets a reference to the given NullableTime and assigns it to the Deadline field.
func (o *UpdateStory) SetDeadline(v time.Time) {
	o.Deadline.Set(&v)
}
// SetDeadlineNil sets the value for Deadline to be an explicit nil
func (o *UpdateStory) SetDeadlineNil() {
	o.Deadline.Set(nil)
}

// UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
func (o *UpdateStory) UnsetDeadline() {
	o.Deadline.Unset()
}

func (o UpdateStory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.PullRequestIds) {
		toSerialize["pull_request_ids"] = o.PullRequestIds
	}
	if !IsNil(o.StoryType) {
		toSerialize["story_type"] = o.StoryType
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.MoveTo) {
		toSerialize["move_to"] = o.MoveTo
	}
	if !IsNil(o.FileIds) {
		toSerialize["file_ids"] = o.FileIds
	}
	if o.CompletedAtOverride.IsSet() {
		toSerialize["completed_at_override"] = o.CompletedAtOverride.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.EpicId.IsSet() {
		toSerialize["epic_id"] = o.EpicId.Get()
	}
	if !IsNil(o.ExternalLinks) {
		toSerialize["external_links"] = o.ExternalLinks
	}
	if !IsNil(o.BranchIds) {
		toSerialize["branch_ids"] = o.BranchIds
	}
	if !IsNil(o.CommitIds) {
		toSerialize["commit_ids"] = o.CommitIds
	}
	if !IsNil(o.RequestedById) {
		toSerialize["requested_by_id"] = o.RequestedById
	}
	if o.IterationId.IsSet() {
		toSerialize["iteration_id"] = o.IterationId.Get()
	}
	if o.StartedAtOverride.IsSet() {
		toSerialize["started_at_override"] = o.StartedAtOverride.Get()
	}
	if o.GroupId.IsSet() {
		toSerialize["group_id"] = o.GroupId.Get()
	}
	if !IsNil(o.WorkflowStateId) {
		toSerialize["workflow_state_id"] = o.WorkflowStateId
	}
	if !IsNil(o.FollowerIds) {
		toSerialize["follower_ids"] = o.FollowerIds
	}
	if !IsNil(o.OwnerIds) {
		toSerialize["owner_ids"] = o.OwnerIds
	}
	if !IsNil(o.BeforeId) {
		toSerialize["before_id"] = o.BeforeId
	}
	if o.Estimate.IsSet() {
		toSerialize["estimate"] = o.Estimate.Get()
	}
	if !IsNil(o.AfterId) {
		toSerialize["after_id"] = o.AfterId
	}
	if o.ProjectId.IsSet() {
		toSerialize["project_id"] = o.ProjectId.Get()
	}
	if !IsNil(o.LinkedFileIds) {
		toSerialize["linked_file_ids"] = o.LinkedFileIds
	}
	if o.Deadline.IsSet() {
		toSerialize["deadline"] = o.Deadline.Get()
	}
	return toSerialize, nil
}

type NullableUpdateStory struct {
	value *UpdateStory
	isSet bool
}

func (v NullableUpdateStory) Get() *UpdateStory {
	return v.value
}

func (v *NullableUpdateStory) Set(val *UpdateStory) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStory) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStory(val *UpdateStory) *NullableUpdateStory {
	return &NullableUpdateStory{value: val, isSet: true}
}

func (v NullableUpdateStory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


