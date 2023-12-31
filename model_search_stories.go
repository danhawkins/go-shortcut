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

// checks if the SearchStories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchStories{}

// SearchStories struct for SearchStories
type SearchStories struct {
	// A true/false boolean indicating whether the Story is in archived state.
	Archived *bool `json:"archived,omitempty"`
	// An array of UUIDs for any Users who may be Owners of the Stories.
	OwnerId NullableString `json:"owner_id,omitempty"`
	// The type of Stories that you want returned.
	StoryType *string `json:"story_type,omitempty"`
	// The Epic IDs that may be associated with the Stories.
	EpicIds []int64 `json:"epic_ids,omitempty"`
	// The IDs for the Projects the Stories may be assigned to.
	ProjectIds []*int64 `json:"project_ids,omitempty"`
	// Stories should have been updated before this date.
	UpdatedAtEnd *time.Time `json:"updated_at_end,omitempty"`
	// Stories should have been completed before this date.
	CompletedAtEnd *time.Time `json:"completed_at_end,omitempty"`
	// The type of Workflow State the Stories may be in.
	WorkflowStateTypes []string `json:"workflow_state_types,omitempty"`
	// Stories should have a deadline before this date.
	DeadlineEnd *time.Time `json:"deadline_end,omitempty"`
	// Stories should have been created after this date.
	CreatedAtStart *time.Time `json:"created_at_start,omitempty"`
	// The Epic IDs that may be associated with the Stories.
	EpicId NullableInt64 `json:"epic_id,omitempty"`
	// The name of any associated Labels.
	LabelName *string `json:"label_name,omitempty"`
	// The UUID of any Users who may have requested the Stories.
	RequestedById *string `json:"requested_by_id,omitempty"`
	// The Iteration ID that may be associated with the Stories.
	IterationId NullableInt64 `json:"iteration_id,omitempty"`
	// The Label IDs that may be associated with the Stories.
	LabelIds []int64 `json:"label_ids,omitempty"`
	// The Group ID that is associated with the Stories
	GroupId NullableString `json:"group_id,omitempty"`
	// The unique IDs of the specific Workflow States that the Stories should be in.
	WorkflowStateId *int64 `json:"workflow_state_id,omitempty"`
	// The Iteration IDs that may be associated with the Stories.
	IterationIds []int64 `json:"iteration_ids,omitempty"`
	// Stories should have been created before this date.
	CreatedAtEnd *time.Time `json:"created_at_end,omitempty"`
	// Stories should have a deadline after this date.
	DeadlineStart *time.Time `json:"deadline_start,omitempty"`
	// The Group IDs that are associated with the Stories
	GroupIds []string `json:"group_ids,omitempty"`
	// An array of UUIDs for any Users who may be Owners of the Stories.
	OwnerIds []string `json:"owner_ids,omitempty"`
	// An ID or URL that references an external resource. Useful during imports.
	ExternalId *string `json:"external_id,omitempty"`
	// Whether to include the story description in the response.
	IncludesDescription *bool `json:"includes_description,omitempty"`
	// The number of estimate points associate with the Stories.
	Estimate *int64 `json:"estimate,omitempty"`
	// The IDs for the Projects the Stories may be assigned to.
	ProjectId NullableInt64 `json:"project_id,omitempty"`
	// Stories should have been completed after this date.
	CompletedAtStart *time.Time `json:"completed_at_start,omitempty"`
	// Stories should have been updated after this date.
	UpdatedAtStart *time.Time `json:"updated_at_start,omitempty"`
}

// NewSearchStories instantiates a new SearchStories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchStories() *SearchStories {
	this := SearchStories{}
	return &this
}

// NewSearchStoriesWithDefaults instantiates a new SearchStories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchStoriesWithDefaults() *SearchStories {
	this := SearchStories{}
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *SearchStories) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *SearchStories) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *SearchStories) SetArchived(v bool) {
	o.Archived = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchStories) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchStories) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *SearchStories) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *SearchStories) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}
// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *SearchStories) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *SearchStories) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetStoryType returns the StoryType field value if set, zero value otherwise.
func (o *SearchStories) GetStoryType() string {
	if o == nil || IsNil(o.StoryType) {
		var ret string
		return ret
	}
	return *o.StoryType
}

// GetStoryTypeOk returns a tuple with the StoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetStoryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StoryType) {
		return nil, false
	}
	return o.StoryType, true
}

// HasStoryType returns a boolean if a field has been set.
func (o *SearchStories) HasStoryType() bool {
	if o != nil && !IsNil(o.StoryType) {
		return true
	}

	return false
}

// SetStoryType gets a reference to the given string and assigns it to the StoryType field.
func (o *SearchStories) SetStoryType(v string) {
	o.StoryType = &v
}

// GetEpicIds returns the EpicIds field value if set, zero value otherwise.
func (o *SearchStories) GetEpicIds() []int64 {
	if o == nil || IsNil(o.EpicIds) {
		var ret []int64
		return ret
	}
	return o.EpicIds
}

// GetEpicIdsOk returns a tuple with the EpicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetEpicIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.EpicIds) {
		return nil, false
	}
	return o.EpicIds, true
}

// HasEpicIds returns a boolean if a field has been set.
func (o *SearchStories) HasEpicIds() bool {
	if o != nil && !IsNil(o.EpicIds) {
		return true
	}

	return false
}

// SetEpicIds gets a reference to the given []int64 and assigns it to the EpicIds field.
func (o *SearchStories) SetEpicIds(v []int64) {
	o.EpicIds = v
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *SearchStories) GetProjectIds() []*int64 {
	if o == nil || IsNil(o.ProjectIds) {
		var ret []*int64
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetProjectIdsOk() ([]*int64, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *SearchStories) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []*int64 and assigns it to the ProjectIds field.
func (o *SearchStories) SetProjectIds(v []*int64) {
	o.ProjectIds = v
}

// GetUpdatedAtEnd returns the UpdatedAtEnd field value if set, zero value otherwise.
func (o *SearchStories) GetUpdatedAtEnd() time.Time {
	if o == nil || IsNil(o.UpdatedAtEnd) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtEnd
}

// GetUpdatedAtEndOk returns a tuple with the UpdatedAtEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetUpdatedAtEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAtEnd) {
		return nil, false
	}
	return o.UpdatedAtEnd, true
}

// HasUpdatedAtEnd returns a boolean if a field has been set.
func (o *SearchStories) HasUpdatedAtEnd() bool {
	if o != nil && !IsNil(o.UpdatedAtEnd) {
		return true
	}

	return false
}

// SetUpdatedAtEnd gets a reference to the given time.Time and assigns it to the UpdatedAtEnd field.
func (o *SearchStories) SetUpdatedAtEnd(v time.Time) {
	o.UpdatedAtEnd = &v
}

// GetCompletedAtEnd returns the CompletedAtEnd field value if set, zero value otherwise.
func (o *SearchStories) GetCompletedAtEnd() time.Time {
	if o == nil || IsNil(o.CompletedAtEnd) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAtEnd
}

// GetCompletedAtEndOk returns a tuple with the CompletedAtEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetCompletedAtEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletedAtEnd) {
		return nil, false
	}
	return o.CompletedAtEnd, true
}

// HasCompletedAtEnd returns a boolean if a field has been set.
func (o *SearchStories) HasCompletedAtEnd() bool {
	if o != nil && !IsNil(o.CompletedAtEnd) {
		return true
	}

	return false
}

// SetCompletedAtEnd gets a reference to the given time.Time and assigns it to the CompletedAtEnd field.
func (o *SearchStories) SetCompletedAtEnd(v time.Time) {
	o.CompletedAtEnd = &v
}

// GetWorkflowStateTypes returns the WorkflowStateTypes field value if set, zero value otherwise.
func (o *SearchStories) GetWorkflowStateTypes() []string {
	if o == nil || IsNil(o.WorkflowStateTypes) {
		var ret []string
		return ret
	}
	return o.WorkflowStateTypes
}

// GetWorkflowStateTypesOk returns a tuple with the WorkflowStateTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetWorkflowStateTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.WorkflowStateTypes) {
		return nil, false
	}
	return o.WorkflowStateTypes, true
}

// HasWorkflowStateTypes returns a boolean if a field has been set.
func (o *SearchStories) HasWorkflowStateTypes() bool {
	if o != nil && !IsNil(o.WorkflowStateTypes) {
		return true
	}

	return false
}

// SetWorkflowStateTypes gets a reference to the given []string and assigns it to the WorkflowStateTypes field.
func (o *SearchStories) SetWorkflowStateTypes(v []string) {
	o.WorkflowStateTypes = v
}

// GetDeadlineEnd returns the DeadlineEnd field value if set, zero value otherwise.
func (o *SearchStories) GetDeadlineEnd() time.Time {
	if o == nil || IsNil(o.DeadlineEnd) {
		var ret time.Time
		return ret
	}
	return *o.DeadlineEnd
}

// GetDeadlineEndOk returns a tuple with the DeadlineEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetDeadlineEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeadlineEnd) {
		return nil, false
	}
	return o.DeadlineEnd, true
}

// HasDeadlineEnd returns a boolean if a field has been set.
func (o *SearchStories) HasDeadlineEnd() bool {
	if o != nil && !IsNil(o.DeadlineEnd) {
		return true
	}

	return false
}

// SetDeadlineEnd gets a reference to the given time.Time and assigns it to the DeadlineEnd field.
func (o *SearchStories) SetDeadlineEnd(v time.Time) {
	o.DeadlineEnd = &v
}

// GetCreatedAtStart returns the CreatedAtStart field value if set, zero value otherwise.
func (o *SearchStories) GetCreatedAtStart() time.Time {
	if o == nil || IsNil(o.CreatedAtStart) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtStart
}

// GetCreatedAtStartOk returns a tuple with the CreatedAtStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetCreatedAtStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAtStart) {
		return nil, false
	}
	return o.CreatedAtStart, true
}

// HasCreatedAtStart returns a boolean if a field has been set.
func (o *SearchStories) HasCreatedAtStart() bool {
	if o != nil && !IsNil(o.CreatedAtStart) {
		return true
	}

	return false
}

// SetCreatedAtStart gets a reference to the given time.Time and assigns it to the CreatedAtStart field.
func (o *SearchStories) SetCreatedAtStart(v time.Time) {
	o.CreatedAtStart = &v
}

// GetEpicId returns the EpicId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchStories) GetEpicId() int64 {
	if o == nil || IsNil(o.EpicId.Get()) {
		var ret int64
		return ret
	}
	return *o.EpicId.Get()
}

// GetEpicIdOk returns a tuple with the EpicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchStories) GetEpicIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EpicId.Get(), o.EpicId.IsSet()
}

// HasEpicId returns a boolean if a field has been set.
func (o *SearchStories) HasEpicId() bool {
	if o != nil && o.EpicId.IsSet() {
		return true
	}

	return false
}

// SetEpicId gets a reference to the given NullableInt64 and assigns it to the EpicId field.
func (o *SearchStories) SetEpicId(v int64) {
	o.EpicId.Set(&v)
}
// SetEpicIdNil sets the value for EpicId to be an explicit nil
func (o *SearchStories) SetEpicIdNil() {
	o.EpicId.Set(nil)
}

// UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
func (o *SearchStories) UnsetEpicId() {
	o.EpicId.Unset()
}

// GetLabelName returns the LabelName field value if set, zero value otherwise.
func (o *SearchStories) GetLabelName() string {
	if o == nil || IsNil(o.LabelName) {
		var ret string
		return ret
	}
	return *o.LabelName
}

// GetLabelNameOk returns a tuple with the LabelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetLabelNameOk() (*string, bool) {
	if o == nil || IsNil(o.LabelName) {
		return nil, false
	}
	return o.LabelName, true
}

// HasLabelName returns a boolean if a field has been set.
func (o *SearchStories) HasLabelName() bool {
	if o != nil && !IsNil(o.LabelName) {
		return true
	}

	return false
}

// SetLabelName gets a reference to the given string and assigns it to the LabelName field.
func (o *SearchStories) SetLabelName(v string) {
	o.LabelName = &v
}

// GetRequestedById returns the RequestedById field value if set, zero value otherwise.
func (o *SearchStories) GetRequestedById() string {
	if o == nil || IsNil(o.RequestedById) {
		var ret string
		return ret
	}
	return *o.RequestedById
}

// GetRequestedByIdOk returns a tuple with the RequestedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetRequestedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedById) {
		return nil, false
	}
	return o.RequestedById, true
}

// HasRequestedById returns a boolean if a field has been set.
func (o *SearchStories) HasRequestedById() bool {
	if o != nil && !IsNil(o.RequestedById) {
		return true
	}

	return false
}

// SetRequestedById gets a reference to the given string and assigns it to the RequestedById field.
func (o *SearchStories) SetRequestedById(v string) {
	o.RequestedById = &v
}

// GetIterationId returns the IterationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchStories) GetIterationId() int64 {
	if o == nil || IsNil(o.IterationId.Get()) {
		var ret int64
		return ret
	}
	return *o.IterationId.Get()
}

// GetIterationIdOk returns a tuple with the IterationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchStories) GetIterationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IterationId.Get(), o.IterationId.IsSet()
}

// HasIterationId returns a boolean if a field has been set.
func (o *SearchStories) HasIterationId() bool {
	if o != nil && o.IterationId.IsSet() {
		return true
	}

	return false
}

// SetIterationId gets a reference to the given NullableInt64 and assigns it to the IterationId field.
func (o *SearchStories) SetIterationId(v int64) {
	o.IterationId.Set(&v)
}
// SetIterationIdNil sets the value for IterationId to be an explicit nil
func (o *SearchStories) SetIterationIdNil() {
	o.IterationId.Set(nil)
}

// UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
func (o *SearchStories) UnsetIterationId() {
	o.IterationId.Unset()
}

// GetLabelIds returns the LabelIds field value if set, zero value otherwise.
func (o *SearchStories) GetLabelIds() []int64 {
	if o == nil || IsNil(o.LabelIds) {
		var ret []int64
		return ret
	}
	return o.LabelIds
}

// GetLabelIdsOk returns a tuple with the LabelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetLabelIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.LabelIds) {
		return nil, false
	}
	return o.LabelIds, true
}

// HasLabelIds returns a boolean if a field has been set.
func (o *SearchStories) HasLabelIds() bool {
	if o != nil && !IsNil(o.LabelIds) {
		return true
	}

	return false
}

// SetLabelIds gets a reference to the given []int64 and assigns it to the LabelIds field.
func (o *SearchStories) SetLabelIds(v []int64) {
	o.LabelIds = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchStories) GetGroupId() string {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret string
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchStories) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *SearchStories) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableString and assigns it to the GroupId field.
func (o *SearchStories) SetGroupId(v string) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *SearchStories) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *SearchStories) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetWorkflowStateId returns the WorkflowStateId field value if set, zero value otherwise.
func (o *SearchStories) GetWorkflowStateId() int64 {
	if o == nil || IsNil(o.WorkflowStateId) {
		var ret int64
		return ret
	}
	return *o.WorkflowStateId
}

// GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetWorkflowStateIdOk() (*int64, bool) {
	if o == nil || IsNil(o.WorkflowStateId) {
		return nil, false
	}
	return o.WorkflowStateId, true
}

// HasWorkflowStateId returns a boolean if a field has been set.
func (o *SearchStories) HasWorkflowStateId() bool {
	if o != nil && !IsNil(o.WorkflowStateId) {
		return true
	}

	return false
}

// SetWorkflowStateId gets a reference to the given int64 and assigns it to the WorkflowStateId field.
func (o *SearchStories) SetWorkflowStateId(v int64) {
	o.WorkflowStateId = &v
}

// GetIterationIds returns the IterationIds field value if set, zero value otherwise.
func (o *SearchStories) GetIterationIds() []int64 {
	if o == nil || IsNil(o.IterationIds) {
		var ret []int64
		return ret
	}
	return o.IterationIds
}

// GetIterationIdsOk returns a tuple with the IterationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetIterationIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.IterationIds) {
		return nil, false
	}
	return o.IterationIds, true
}

// HasIterationIds returns a boolean if a field has been set.
func (o *SearchStories) HasIterationIds() bool {
	if o != nil && !IsNil(o.IterationIds) {
		return true
	}

	return false
}

// SetIterationIds gets a reference to the given []int64 and assigns it to the IterationIds field.
func (o *SearchStories) SetIterationIds(v []int64) {
	o.IterationIds = v
}

// GetCreatedAtEnd returns the CreatedAtEnd field value if set, zero value otherwise.
func (o *SearchStories) GetCreatedAtEnd() time.Time {
	if o == nil || IsNil(o.CreatedAtEnd) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtEnd
}

// GetCreatedAtEndOk returns a tuple with the CreatedAtEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetCreatedAtEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAtEnd) {
		return nil, false
	}
	return o.CreatedAtEnd, true
}

// HasCreatedAtEnd returns a boolean if a field has been set.
func (o *SearchStories) HasCreatedAtEnd() bool {
	if o != nil && !IsNil(o.CreatedAtEnd) {
		return true
	}

	return false
}

// SetCreatedAtEnd gets a reference to the given time.Time and assigns it to the CreatedAtEnd field.
func (o *SearchStories) SetCreatedAtEnd(v time.Time) {
	o.CreatedAtEnd = &v
}

// GetDeadlineStart returns the DeadlineStart field value if set, zero value otherwise.
func (o *SearchStories) GetDeadlineStart() time.Time {
	if o == nil || IsNil(o.DeadlineStart) {
		var ret time.Time
		return ret
	}
	return *o.DeadlineStart
}

// GetDeadlineStartOk returns a tuple with the DeadlineStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetDeadlineStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeadlineStart) {
		return nil, false
	}
	return o.DeadlineStart, true
}

// HasDeadlineStart returns a boolean if a field has been set.
func (o *SearchStories) HasDeadlineStart() bool {
	if o != nil && !IsNil(o.DeadlineStart) {
		return true
	}

	return false
}

// SetDeadlineStart gets a reference to the given time.Time and assigns it to the DeadlineStart field.
func (o *SearchStories) SetDeadlineStart(v time.Time) {
	o.DeadlineStart = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *SearchStories) GetGroupIds() []string {
	if o == nil || IsNil(o.GroupIds) {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupIds) {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *SearchStories) HasGroupIds() bool {
	if o != nil && !IsNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *SearchStories) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetOwnerIds returns the OwnerIds field value if set, zero value otherwise.
func (o *SearchStories) GetOwnerIds() []string {
	if o == nil || IsNil(o.OwnerIds) {
		var ret []string
		return ret
	}
	return o.OwnerIds
}

// GetOwnerIdsOk returns a tuple with the OwnerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetOwnerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerIds) {
		return nil, false
	}
	return o.OwnerIds, true
}

// HasOwnerIds returns a boolean if a field has been set.
func (o *SearchStories) HasOwnerIds() bool {
	if o != nil && !IsNil(o.OwnerIds) {
		return true
	}

	return false
}

// SetOwnerIds gets a reference to the given []string and assigns it to the OwnerIds field.
func (o *SearchStories) SetOwnerIds(v []string) {
	o.OwnerIds = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *SearchStories) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *SearchStories) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *SearchStories) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetIncludesDescription returns the IncludesDescription field value if set, zero value otherwise.
func (o *SearchStories) GetIncludesDescription() bool {
	if o == nil || IsNil(o.IncludesDescription) {
		var ret bool
		return ret
	}
	return *o.IncludesDescription
}

// GetIncludesDescriptionOk returns a tuple with the IncludesDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetIncludesDescriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludesDescription) {
		return nil, false
	}
	return o.IncludesDescription, true
}

// HasIncludesDescription returns a boolean if a field has been set.
func (o *SearchStories) HasIncludesDescription() bool {
	if o != nil && !IsNil(o.IncludesDescription) {
		return true
	}

	return false
}

// SetIncludesDescription gets a reference to the given bool and assigns it to the IncludesDescription field.
func (o *SearchStories) SetIncludesDescription(v bool) {
	o.IncludesDescription = &v
}

// GetEstimate returns the Estimate field value if set, zero value otherwise.
func (o *SearchStories) GetEstimate() int64 {
	if o == nil || IsNil(o.Estimate) {
		var ret int64
		return ret
	}
	return *o.Estimate
}

// GetEstimateOk returns a tuple with the Estimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetEstimateOk() (*int64, bool) {
	if o == nil || IsNil(o.Estimate) {
		return nil, false
	}
	return o.Estimate, true
}

// HasEstimate returns a boolean if a field has been set.
func (o *SearchStories) HasEstimate() bool {
	if o != nil && !IsNil(o.Estimate) {
		return true
	}

	return false
}

// SetEstimate gets a reference to the given int64 and assigns it to the Estimate field.
func (o *SearchStories) SetEstimate(v int64) {
	o.Estimate = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchStories) GetProjectId() int64 {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchStories) GetProjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *SearchStories) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableInt64 and assigns it to the ProjectId field.
func (o *SearchStories) SetProjectId(v int64) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *SearchStories) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *SearchStories) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetCompletedAtStart returns the CompletedAtStart field value if set, zero value otherwise.
func (o *SearchStories) GetCompletedAtStart() time.Time {
	if o == nil || IsNil(o.CompletedAtStart) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAtStart
}

// GetCompletedAtStartOk returns a tuple with the CompletedAtStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetCompletedAtStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletedAtStart) {
		return nil, false
	}
	return o.CompletedAtStart, true
}

// HasCompletedAtStart returns a boolean if a field has been set.
func (o *SearchStories) HasCompletedAtStart() bool {
	if o != nil && !IsNil(o.CompletedAtStart) {
		return true
	}

	return false
}

// SetCompletedAtStart gets a reference to the given time.Time and assigns it to the CompletedAtStart field.
func (o *SearchStories) SetCompletedAtStart(v time.Time) {
	o.CompletedAtStart = &v
}

// GetUpdatedAtStart returns the UpdatedAtStart field value if set, zero value otherwise.
func (o *SearchStories) GetUpdatedAtStart() time.Time {
	if o == nil || IsNil(o.UpdatedAtStart) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtStart
}

// GetUpdatedAtStartOk returns a tuple with the UpdatedAtStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchStories) GetUpdatedAtStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAtStart) {
		return nil, false
	}
	return o.UpdatedAtStart, true
}

// HasUpdatedAtStart returns a boolean if a field has been set.
func (o *SearchStories) HasUpdatedAtStart() bool {
	if o != nil && !IsNil(o.UpdatedAtStart) {
		return true
	}

	return false
}

// SetUpdatedAtStart gets a reference to the given time.Time and assigns it to the UpdatedAtStart field.
func (o *SearchStories) SetUpdatedAtStart(v time.Time) {
	o.UpdatedAtStart = &v
}

func (o SearchStories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchStories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if o.OwnerId.IsSet() {
		toSerialize["owner_id"] = o.OwnerId.Get()
	}
	if !IsNil(o.StoryType) {
		toSerialize["story_type"] = o.StoryType
	}
	if !IsNil(o.EpicIds) {
		toSerialize["epic_ids"] = o.EpicIds
	}
	if !IsNil(o.ProjectIds) {
		toSerialize["project_ids"] = o.ProjectIds
	}
	if !IsNil(o.UpdatedAtEnd) {
		toSerialize["updated_at_end"] = o.UpdatedAtEnd
	}
	if !IsNil(o.CompletedAtEnd) {
		toSerialize["completed_at_end"] = o.CompletedAtEnd
	}
	if !IsNil(o.WorkflowStateTypes) {
		toSerialize["workflow_state_types"] = o.WorkflowStateTypes
	}
	if !IsNil(o.DeadlineEnd) {
		toSerialize["deadline_end"] = o.DeadlineEnd
	}
	if !IsNil(o.CreatedAtStart) {
		toSerialize["created_at_start"] = o.CreatedAtStart
	}
	if o.EpicId.IsSet() {
		toSerialize["epic_id"] = o.EpicId.Get()
	}
	if !IsNil(o.LabelName) {
		toSerialize["label_name"] = o.LabelName
	}
	if !IsNil(o.RequestedById) {
		toSerialize["requested_by_id"] = o.RequestedById
	}
	if o.IterationId.IsSet() {
		toSerialize["iteration_id"] = o.IterationId.Get()
	}
	if !IsNil(o.LabelIds) {
		toSerialize["label_ids"] = o.LabelIds
	}
	if o.GroupId.IsSet() {
		toSerialize["group_id"] = o.GroupId.Get()
	}
	if !IsNil(o.WorkflowStateId) {
		toSerialize["workflow_state_id"] = o.WorkflowStateId
	}
	if !IsNil(o.IterationIds) {
		toSerialize["iteration_ids"] = o.IterationIds
	}
	if !IsNil(o.CreatedAtEnd) {
		toSerialize["created_at_end"] = o.CreatedAtEnd
	}
	if !IsNil(o.DeadlineStart) {
		toSerialize["deadline_start"] = o.DeadlineStart
	}
	if !IsNil(o.GroupIds) {
		toSerialize["group_ids"] = o.GroupIds
	}
	if !IsNil(o.OwnerIds) {
		toSerialize["owner_ids"] = o.OwnerIds
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.IncludesDescription) {
		toSerialize["includes_description"] = o.IncludesDescription
	}
	if !IsNil(o.Estimate) {
		toSerialize["estimate"] = o.Estimate
	}
	if o.ProjectId.IsSet() {
		toSerialize["project_id"] = o.ProjectId.Get()
	}
	if !IsNil(o.CompletedAtStart) {
		toSerialize["completed_at_start"] = o.CompletedAtStart
	}
	if !IsNil(o.UpdatedAtStart) {
		toSerialize["updated_at_start"] = o.UpdatedAtStart
	}
	return toSerialize, nil
}

type NullableSearchStories struct {
	value *SearchStories
	isSet bool
}

func (v NullableSearchStories) Get() *SearchStories {
	return v.value
}

func (v *NullableSearchStories) Set(val *SearchStories) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchStories) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchStories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchStories(val *SearchStories) *NullableSearchStories {
	return &NullableSearchStories{value: val, isSet: true}
}

func (v NullableSearchStories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchStories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


