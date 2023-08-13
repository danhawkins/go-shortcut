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

// checks if the UpdateStories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStories{}

// UpdateStories struct for UpdateStories
type UpdateStories struct {
	// If the Stories should be archived or not.
	Archived *bool `json:"archived,omitempty"`
	// The Ids of the Stories you wish to update.
	StoryIds []int64 `json:"story_ids"`
	// The type of story (feature, bug, chore).
	StoryType *string `json:"story_type,omitempty"`
	// One of \"first\" or \"last\". This can be used to move the given story to the first or last position in the workflow state.
	MoveTo *string `json:"move_to,omitempty"`
	// The UUIDs of the new followers to be added.
	FollowerIdsAdd []string `json:"follower_ids_add,omitempty"`
	// The ID of the epic the story belongs to.
	EpicId NullableInt64 `json:"epic_id,omitempty"`
	// An array of External Links associated with this story.
	ExternalLinks []string `json:"external_links,omitempty"`
	// The UUIDs of the followers to be removed.
	FollowerIdsRemove []string `json:"follower_ids_remove,omitempty"`
	// The ID of the member that requested the story.
	RequestedById *string `json:"requested_by_id,omitempty"`
	// The ID of the iteration the story belongs to.
	IterationId NullableInt64 `json:"iteration_id,omitempty"`
	// A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField.
	CustomFieldsRemove []CustomFieldValueParams `json:"custom_fields_remove,omitempty"`
	// An array of labels to be added.
	LabelsAdd []CreateLabelParams `json:"labels_add,omitempty"`
	// The Id of the Group the Stories should belong to.
	GroupId NullableString `json:"group_id,omitempty"`
	// The ID of the workflow state to put the stories in.
	WorkflowStateId *int64 `json:"workflow_state_id,omitempty"`
	// The ID of the story that the stories are to be moved before.
	BeforeId *int64 `json:"before_id,omitempty"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate NullableInt64 `json:"estimate,omitempty"`
	// The ID of the story that the stories are to be moved below.
	AfterId *int64 `json:"after_id,omitempty"`
	// The UUIDs of the owners to be removed.
	OwnerIdsRemove []string `json:"owner_ids_remove,omitempty"`
	// A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField.
	CustomFieldsAdd []CustomFieldValueParams `json:"custom_fields_add,omitempty"`
	// The ID of the Project the Stories should belong to.
	ProjectId NullableInt64 `json:"project_id,omitempty"`
	// An array of labels to be removed.
	LabelsRemove []CreateLabelParams `json:"labels_remove,omitempty"`
	// The due date of the story.
	Deadline NullableTime `json:"deadline,omitempty"`
	// The UUIDs of the new owners to be added.
	OwnerIdsAdd []string `json:"owner_ids_add,omitempty"`
}

// NewUpdateStories instantiates a new UpdateStories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStories(storyIds []int64) *UpdateStories {
	this := UpdateStories{}
	this.StoryIds = storyIds
	return &this
}

// NewUpdateStoriesWithDefaults instantiates a new UpdateStories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStoriesWithDefaults() *UpdateStories {
	this := UpdateStories{}
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *UpdateStories) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *UpdateStories) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *UpdateStories) SetArchived(v bool) {
	o.Archived = &v
}

// GetStoryIds returns the StoryIds field value
func (o *UpdateStories) GetStoryIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.StoryIds
}

// GetStoryIdsOk returns a tuple with the StoryIds field value
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetStoryIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoryIds, true
}

// SetStoryIds sets field value
func (o *UpdateStories) SetStoryIds(v []int64) {
	o.StoryIds = v
}

// GetStoryType returns the StoryType field value if set, zero value otherwise.
func (o *UpdateStories) GetStoryType() string {
	if o == nil || IsNil(o.StoryType) {
		var ret string
		return ret
	}
	return *o.StoryType
}

// GetStoryTypeOk returns a tuple with the StoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetStoryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StoryType) {
		return nil, false
	}
	return o.StoryType, true
}

// HasStoryType returns a boolean if a field has been set.
func (o *UpdateStories) HasStoryType() bool {
	if o != nil && !IsNil(o.StoryType) {
		return true
	}

	return false
}

// SetStoryType gets a reference to the given string and assigns it to the StoryType field.
func (o *UpdateStories) SetStoryType(v string) {
	o.StoryType = &v
}

// GetMoveTo returns the MoveTo field value if set, zero value otherwise.
func (o *UpdateStories) GetMoveTo() string {
	if o == nil || IsNil(o.MoveTo) {
		var ret string
		return ret
	}
	return *o.MoveTo
}

// GetMoveToOk returns a tuple with the MoveTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetMoveToOk() (*string, bool) {
	if o == nil || IsNil(o.MoveTo) {
		return nil, false
	}
	return o.MoveTo, true
}

// HasMoveTo returns a boolean if a field has been set.
func (o *UpdateStories) HasMoveTo() bool {
	if o != nil && !IsNil(o.MoveTo) {
		return true
	}

	return false
}

// SetMoveTo gets a reference to the given string and assigns it to the MoveTo field.
func (o *UpdateStories) SetMoveTo(v string) {
	o.MoveTo = &v
}

// GetFollowerIdsAdd returns the FollowerIdsAdd field value if set, zero value otherwise.
func (o *UpdateStories) GetFollowerIdsAdd() []string {
	if o == nil || IsNil(o.FollowerIdsAdd) {
		var ret []string
		return ret
	}
	return o.FollowerIdsAdd
}

// GetFollowerIdsAddOk returns a tuple with the FollowerIdsAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetFollowerIdsAddOk() ([]string, bool) {
	if o == nil || IsNil(o.FollowerIdsAdd) {
		return nil, false
	}
	return o.FollowerIdsAdd, true
}

// HasFollowerIdsAdd returns a boolean if a field has been set.
func (o *UpdateStories) HasFollowerIdsAdd() bool {
	if o != nil && !IsNil(o.FollowerIdsAdd) {
		return true
	}

	return false
}

// SetFollowerIdsAdd gets a reference to the given []string and assigns it to the FollowerIdsAdd field.
func (o *UpdateStories) SetFollowerIdsAdd(v []string) {
	o.FollowerIdsAdd = v
}

// GetEpicId returns the EpicId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStories) GetEpicId() int64 {
	if o == nil || IsNil(o.EpicId.Get()) {
		var ret int64
		return ret
	}
	return *o.EpicId.Get()
}

// GetEpicIdOk returns a tuple with the EpicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStories) GetEpicIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EpicId.Get(), o.EpicId.IsSet()
}

// HasEpicId returns a boolean if a field has been set.
func (o *UpdateStories) HasEpicId() bool {
	if o != nil && o.EpicId.IsSet() {
		return true
	}

	return false
}

// SetEpicId gets a reference to the given NullableInt64 and assigns it to the EpicId field.
func (o *UpdateStories) SetEpicId(v int64) {
	o.EpicId.Set(&v)
}
// SetEpicIdNil sets the value for EpicId to be an explicit nil
func (o *UpdateStories) SetEpicIdNil() {
	o.EpicId.Set(nil)
}

// UnsetEpicId ensures that no value is present for EpicId, not even an explicit nil
func (o *UpdateStories) UnsetEpicId() {
	o.EpicId.Unset()
}

// GetExternalLinks returns the ExternalLinks field value if set, zero value otherwise.
func (o *UpdateStories) GetExternalLinks() []string {
	if o == nil || IsNil(o.ExternalLinks) {
		var ret []string
		return ret
	}
	return o.ExternalLinks
}

// GetExternalLinksOk returns a tuple with the ExternalLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetExternalLinksOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalLinks) {
		return nil, false
	}
	return o.ExternalLinks, true
}

// HasExternalLinks returns a boolean if a field has been set.
func (o *UpdateStories) HasExternalLinks() bool {
	if o != nil && !IsNil(o.ExternalLinks) {
		return true
	}

	return false
}

// SetExternalLinks gets a reference to the given []string and assigns it to the ExternalLinks field.
func (o *UpdateStories) SetExternalLinks(v []string) {
	o.ExternalLinks = v
}

// GetFollowerIdsRemove returns the FollowerIdsRemove field value if set, zero value otherwise.
func (o *UpdateStories) GetFollowerIdsRemove() []string {
	if o == nil || IsNil(o.FollowerIdsRemove) {
		var ret []string
		return ret
	}
	return o.FollowerIdsRemove
}

// GetFollowerIdsRemoveOk returns a tuple with the FollowerIdsRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetFollowerIdsRemoveOk() ([]string, bool) {
	if o == nil || IsNil(o.FollowerIdsRemove) {
		return nil, false
	}
	return o.FollowerIdsRemove, true
}

// HasFollowerIdsRemove returns a boolean if a field has been set.
func (o *UpdateStories) HasFollowerIdsRemove() bool {
	if o != nil && !IsNil(o.FollowerIdsRemove) {
		return true
	}

	return false
}

// SetFollowerIdsRemove gets a reference to the given []string and assigns it to the FollowerIdsRemove field.
func (o *UpdateStories) SetFollowerIdsRemove(v []string) {
	o.FollowerIdsRemove = v
}

// GetRequestedById returns the RequestedById field value if set, zero value otherwise.
func (o *UpdateStories) GetRequestedById() string {
	if o == nil || IsNil(o.RequestedById) {
		var ret string
		return ret
	}
	return *o.RequestedById
}

// GetRequestedByIdOk returns a tuple with the RequestedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetRequestedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedById) {
		return nil, false
	}
	return o.RequestedById, true
}

// HasRequestedById returns a boolean if a field has been set.
func (o *UpdateStories) HasRequestedById() bool {
	if o != nil && !IsNil(o.RequestedById) {
		return true
	}

	return false
}

// SetRequestedById gets a reference to the given string and assigns it to the RequestedById field.
func (o *UpdateStories) SetRequestedById(v string) {
	o.RequestedById = &v
}

// GetIterationId returns the IterationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStories) GetIterationId() int64 {
	if o == nil || IsNil(o.IterationId.Get()) {
		var ret int64
		return ret
	}
	return *o.IterationId.Get()
}

// GetIterationIdOk returns a tuple with the IterationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStories) GetIterationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IterationId.Get(), o.IterationId.IsSet()
}

// HasIterationId returns a boolean if a field has been set.
func (o *UpdateStories) HasIterationId() bool {
	if o != nil && o.IterationId.IsSet() {
		return true
	}

	return false
}

// SetIterationId gets a reference to the given NullableInt64 and assigns it to the IterationId field.
func (o *UpdateStories) SetIterationId(v int64) {
	o.IterationId.Set(&v)
}
// SetIterationIdNil sets the value for IterationId to be an explicit nil
func (o *UpdateStories) SetIterationIdNil() {
	o.IterationId.Set(nil)
}

// UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
func (o *UpdateStories) UnsetIterationId() {
	o.IterationId.Unset()
}

// GetCustomFieldsRemove returns the CustomFieldsRemove field value if set, zero value otherwise.
func (o *UpdateStories) GetCustomFieldsRemove() []CustomFieldValueParams {
	if o == nil || IsNil(o.CustomFieldsRemove) {
		var ret []CustomFieldValueParams
		return ret
	}
	return o.CustomFieldsRemove
}

// GetCustomFieldsRemoveOk returns a tuple with the CustomFieldsRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetCustomFieldsRemoveOk() ([]CustomFieldValueParams, bool) {
	if o == nil || IsNil(o.CustomFieldsRemove) {
		return nil, false
	}
	return o.CustomFieldsRemove, true
}

// HasCustomFieldsRemove returns a boolean if a field has been set.
func (o *UpdateStories) HasCustomFieldsRemove() bool {
	if o != nil && !IsNil(o.CustomFieldsRemove) {
		return true
	}

	return false
}

// SetCustomFieldsRemove gets a reference to the given []CustomFieldValueParams and assigns it to the CustomFieldsRemove field.
func (o *UpdateStories) SetCustomFieldsRemove(v []CustomFieldValueParams) {
	o.CustomFieldsRemove = v
}

// GetLabelsAdd returns the LabelsAdd field value if set, zero value otherwise.
func (o *UpdateStories) GetLabelsAdd() []CreateLabelParams {
	if o == nil || IsNil(o.LabelsAdd) {
		var ret []CreateLabelParams
		return ret
	}
	return o.LabelsAdd
}

// GetLabelsAddOk returns a tuple with the LabelsAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetLabelsAddOk() ([]CreateLabelParams, bool) {
	if o == nil || IsNil(o.LabelsAdd) {
		return nil, false
	}
	return o.LabelsAdd, true
}

// HasLabelsAdd returns a boolean if a field has been set.
func (o *UpdateStories) HasLabelsAdd() bool {
	if o != nil && !IsNil(o.LabelsAdd) {
		return true
	}

	return false
}

// SetLabelsAdd gets a reference to the given []CreateLabelParams and assigns it to the LabelsAdd field.
func (o *UpdateStories) SetLabelsAdd(v []CreateLabelParams) {
	o.LabelsAdd = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStories) GetGroupId() string {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret string
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStories) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *UpdateStories) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableString and assigns it to the GroupId field.
func (o *UpdateStories) SetGroupId(v string) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *UpdateStories) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *UpdateStories) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetWorkflowStateId returns the WorkflowStateId field value if set, zero value otherwise.
func (o *UpdateStories) GetWorkflowStateId() int64 {
	if o == nil || IsNil(o.WorkflowStateId) {
		var ret int64
		return ret
	}
	return *o.WorkflowStateId
}

// GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetWorkflowStateIdOk() (*int64, bool) {
	if o == nil || IsNil(o.WorkflowStateId) {
		return nil, false
	}
	return o.WorkflowStateId, true
}

// HasWorkflowStateId returns a boolean if a field has been set.
func (o *UpdateStories) HasWorkflowStateId() bool {
	if o != nil && !IsNil(o.WorkflowStateId) {
		return true
	}

	return false
}

// SetWorkflowStateId gets a reference to the given int64 and assigns it to the WorkflowStateId field.
func (o *UpdateStories) SetWorkflowStateId(v int64) {
	o.WorkflowStateId = &v
}

// GetBeforeId returns the BeforeId field value if set, zero value otherwise.
func (o *UpdateStories) GetBeforeId() int64 {
	if o == nil || IsNil(o.BeforeId) {
		var ret int64
		return ret
	}
	return *o.BeforeId
}

// GetBeforeIdOk returns a tuple with the BeforeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetBeforeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BeforeId) {
		return nil, false
	}
	return o.BeforeId, true
}

// HasBeforeId returns a boolean if a field has been set.
func (o *UpdateStories) HasBeforeId() bool {
	if o != nil && !IsNil(o.BeforeId) {
		return true
	}

	return false
}

// SetBeforeId gets a reference to the given int64 and assigns it to the BeforeId field.
func (o *UpdateStories) SetBeforeId(v int64) {
	o.BeforeId = &v
}

// GetEstimate returns the Estimate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStories) GetEstimate() int64 {
	if o == nil || IsNil(o.Estimate.Get()) {
		var ret int64
		return ret
	}
	return *o.Estimate.Get()
}

// GetEstimateOk returns a tuple with the Estimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStories) GetEstimateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Estimate.Get(), o.Estimate.IsSet()
}

// HasEstimate returns a boolean if a field has been set.
func (o *UpdateStories) HasEstimate() bool {
	if o != nil && o.Estimate.IsSet() {
		return true
	}

	return false
}

// SetEstimate gets a reference to the given NullableInt64 and assigns it to the Estimate field.
func (o *UpdateStories) SetEstimate(v int64) {
	o.Estimate.Set(&v)
}
// SetEstimateNil sets the value for Estimate to be an explicit nil
func (o *UpdateStories) SetEstimateNil() {
	o.Estimate.Set(nil)
}

// UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
func (o *UpdateStories) UnsetEstimate() {
	o.Estimate.Unset()
}

// GetAfterId returns the AfterId field value if set, zero value otherwise.
func (o *UpdateStories) GetAfterId() int64 {
	if o == nil || IsNil(o.AfterId) {
		var ret int64
		return ret
	}
	return *o.AfterId
}

// GetAfterIdOk returns a tuple with the AfterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetAfterIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AfterId) {
		return nil, false
	}
	return o.AfterId, true
}

// HasAfterId returns a boolean if a field has been set.
func (o *UpdateStories) HasAfterId() bool {
	if o != nil && !IsNil(o.AfterId) {
		return true
	}

	return false
}

// SetAfterId gets a reference to the given int64 and assigns it to the AfterId field.
func (o *UpdateStories) SetAfterId(v int64) {
	o.AfterId = &v
}

// GetOwnerIdsRemove returns the OwnerIdsRemove field value if set, zero value otherwise.
func (o *UpdateStories) GetOwnerIdsRemove() []string {
	if o == nil || IsNil(o.OwnerIdsRemove) {
		var ret []string
		return ret
	}
	return o.OwnerIdsRemove
}

// GetOwnerIdsRemoveOk returns a tuple with the OwnerIdsRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetOwnerIdsRemoveOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerIdsRemove) {
		return nil, false
	}
	return o.OwnerIdsRemove, true
}

// HasOwnerIdsRemove returns a boolean if a field has been set.
func (o *UpdateStories) HasOwnerIdsRemove() bool {
	if o != nil && !IsNil(o.OwnerIdsRemove) {
		return true
	}

	return false
}

// SetOwnerIdsRemove gets a reference to the given []string and assigns it to the OwnerIdsRemove field.
func (o *UpdateStories) SetOwnerIdsRemove(v []string) {
	o.OwnerIdsRemove = v
}

// GetCustomFieldsAdd returns the CustomFieldsAdd field value if set, zero value otherwise.
func (o *UpdateStories) GetCustomFieldsAdd() []CustomFieldValueParams {
	if o == nil || IsNil(o.CustomFieldsAdd) {
		var ret []CustomFieldValueParams
		return ret
	}
	return o.CustomFieldsAdd
}

// GetCustomFieldsAddOk returns a tuple with the CustomFieldsAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetCustomFieldsAddOk() ([]CustomFieldValueParams, bool) {
	if o == nil || IsNil(o.CustomFieldsAdd) {
		return nil, false
	}
	return o.CustomFieldsAdd, true
}

// HasCustomFieldsAdd returns a boolean if a field has been set.
func (o *UpdateStories) HasCustomFieldsAdd() bool {
	if o != nil && !IsNil(o.CustomFieldsAdd) {
		return true
	}

	return false
}

// SetCustomFieldsAdd gets a reference to the given []CustomFieldValueParams and assigns it to the CustomFieldsAdd field.
func (o *UpdateStories) SetCustomFieldsAdd(v []CustomFieldValueParams) {
	o.CustomFieldsAdd = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStories) GetProjectId() int64 {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStories) GetProjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *UpdateStories) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableInt64 and assigns it to the ProjectId field.
func (o *UpdateStories) SetProjectId(v int64) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *UpdateStories) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *UpdateStories) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetLabelsRemove returns the LabelsRemove field value if set, zero value otherwise.
func (o *UpdateStories) GetLabelsRemove() []CreateLabelParams {
	if o == nil || IsNil(o.LabelsRemove) {
		var ret []CreateLabelParams
		return ret
	}
	return o.LabelsRemove
}

// GetLabelsRemoveOk returns a tuple with the LabelsRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetLabelsRemoveOk() ([]CreateLabelParams, bool) {
	if o == nil || IsNil(o.LabelsRemove) {
		return nil, false
	}
	return o.LabelsRemove, true
}

// HasLabelsRemove returns a boolean if a field has been set.
func (o *UpdateStories) HasLabelsRemove() bool {
	if o != nil && !IsNil(o.LabelsRemove) {
		return true
	}

	return false
}

// SetLabelsRemove gets a reference to the given []CreateLabelParams and assigns it to the LabelsRemove field.
func (o *UpdateStories) SetLabelsRemove(v []CreateLabelParams) {
	o.LabelsRemove = v
}

// GetDeadline returns the Deadline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateStories) GetDeadline() time.Time {
	if o == nil || IsNil(o.Deadline.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Deadline.Get()
}

// GetDeadlineOk returns a tuple with the Deadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateStories) GetDeadlineOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deadline.Get(), o.Deadline.IsSet()
}

// HasDeadline returns a boolean if a field has been set.
func (o *UpdateStories) HasDeadline() bool {
	if o != nil && o.Deadline.IsSet() {
		return true
	}

	return false
}

// SetDeadline gets a reference to the given NullableTime and assigns it to the Deadline field.
func (o *UpdateStories) SetDeadline(v time.Time) {
	o.Deadline.Set(&v)
}
// SetDeadlineNil sets the value for Deadline to be an explicit nil
func (o *UpdateStories) SetDeadlineNil() {
	o.Deadline.Set(nil)
}

// UnsetDeadline ensures that no value is present for Deadline, not even an explicit nil
func (o *UpdateStories) UnsetDeadline() {
	o.Deadline.Unset()
}

// GetOwnerIdsAdd returns the OwnerIdsAdd field value if set, zero value otherwise.
func (o *UpdateStories) GetOwnerIdsAdd() []string {
	if o == nil || IsNil(o.OwnerIdsAdd) {
		var ret []string
		return ret
	}
	return o.OwnerIdsAdd
}

// GetOwnerIdsAddOk returns a tuple with the OwnerIdsAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStories) GetOwnerIdsAddOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerIdsAdd) {
		return nil, false
	}
	return o.OwnerIdsAdd, true
}

// HasOwnerIdsAdd returns a boolean if a field has been set.
func (o *UpdateStories) HasOwnerIdsAdd() bool {
	if o != nil && !IsNil(o.OwnerIdsAdd) {
		return true
	}

	return false
}

// SetOwnerIdsAdd gets a reference to the given []string and assigns it to the OwnerIdsAdd field.
func (o *UpdateStories) SetOwnerIdsAdd(v []string) {
	o.OwnerIdsAdd = v
}

func (o UpdateStories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	toSerialize["story_ids"] = o.StoryIds
	if !IsNil(o.StoryType) {
		toSerialize["story_type"] = o.StoryType
	}
	if !IsNil(o.MoveTo) {
		toSerialize["move_to"] = o.MoveTo
	}
	if !IsNil(o.FollowerIdsAdd) {
		toSerialize["follower_ids_add"] = o.FollowerIdsAdd
	}
	if o.EpicId.IsSet() {
		toSerialize["epic_id"] = o.EpicId.Get()
	}
	if !IsNil(o.ExternalLinks) {
		toSerialize["external_links"] = o.ExternalLinks
	}
	if !IsNil(o.FollowerIdsRemove) {
		toSerialize["follower_ids_remove"] = o.FollowerIdsRemove
	}
	if !IsNil(o.RequestedById) {
		toSerialize["requested_by_id"] = o.RequestedById
	}
	if o.IterationId.IsSet() {
		toSerialize["iteration_id"] = o.IterationId.Get()
	}
	if !IsNil(o.CustomFieldsRemove) {
		toSerialize["custom_fields_remove"] = o.CustomFieldsRemove
	}
	if !IsNil(o.LabelsAdd) {
		toSerialize["labels_add"] = o.LabelsAdd
	}
	if o.GroupId.IsSet() {
		toSerialize["group_id"] = o.GroupId.Get()
	}
	if !IsNil(o.WorkflowStateId) {
		toSerialize["workflow_state_id"] = o.WorkflowStateId
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
	if !IsNil(o.OwnerIdsRemove) {
		toSerialize["owner_ids_remove"] = o.OwnerIdsRemove
	}
	if !IsNil(o.CustomFieldsAdd) {
		toSerialize["custom_fields_add"] = o.CustomFieldsAdd
	}
	if o.ProjectId.IsSet() {
		toSerialize["project_id"] = o.ProjectId.Get()
	}
	if !IsNil(o.LabelsRemove) {
		toSerialize["labels_remove"] = o.LabelsRemove
	}
	if o.Deadline.IsSet() {
		toSerialize["deadline"] = o.Deadline.Get()
	}
	if !IsNil(o.OwnerIdsAdd) {
		toSerialize["owner_ids_add"] = o.OwnerIdsAdd
	}
	return toSerialize, nil
}

type NullableUpdateStories struct {
	value *UpdateStories
	isSet bool
}

func (v NullableUpdateStories) Get() *UpdateStories {
	return v.value
}

func (v *NullableUpdateStories) Set(val *UpdateStories) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStories) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStories(val *UpdateStories) *NullableUpdateStories {
	return &NullableUpdateStories{value: val, isSet: true}
}

func (v NullableUpdateStories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


