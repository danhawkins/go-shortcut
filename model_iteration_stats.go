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

// checks if the IterationStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IterationStats{}

// IterationStats A group of calculated values for this Iteration.
type IterationStats struct {
	// The total number of completed points in this Iteration.
	NumPointsDone int64 `json:"num_points_done"`
	// The total number of documents related to an Iteration
	NumRelatedDocuments int64 `json:"num_related_documents"`
	// The average cycle time (in seconds) of completed stories in this Iteration.
	AverageCycleTime *int64 `json:"average_cycle_time,omitempty"`
	// The total number of unstarted Stories in this Iteration.
	NumStoriesUnstarted int64 `json:"num_stories_unstarted"`
	// The total number of started points in this Iteration.
	NumPointsStarted int64 `json:"num_points_started"`
	// The total number of unstarted points in this Iteration.
	NumPointsUnstarted int64 `json:"num_points_unstarted"`
	// The total number of started Stories in this Iteration.
	NumStoriesStarted int64 `json:"num_stories_started"`
	// The total number of Stories with no point estimate.
	NumStoriesUnestimated int64 `json:"num_stories_unestimated"`
	// The total number of backlog Stories in this Iteration.
	NumStoriesBacklog int64 `json:"num_stories_backlog"`
	// The average lead time (in seconds) of completed stories in this Iteration.
	AverageLeadTime *int64 `json:"average_lead_time,omitempty"`
	// The total number of backlog points in this Iteration.
	NumPointsBacklog int64 `json:"num_points_backlog"`
	// The total number of points in this Iteration.
	NumPoints int64 `json:"num_points"`
	// The total number of done Stories in this Iteration.
	NumStoriesDone int64 `json:"num_stories_done"`
}

// NewIterationStats instantiates a new IterationStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIterationStats(numPointsDone int64, numRelatedDocuments int64, numStoriesUnstarted int64, numPointsStarted int64, numPointsUnstarted int64, numStoriesStarted int64, numStoriesUnestimated int64, numStoriesBacklog int64, numPointsBacklog int64, numPoints int64, numStoriesDone int64) *IterationStats {
	this := IterationStats{}
	this.NumPointsDone = numPointsDone
	this.NumRelatedDocuments = numRelatedDocuments
	this.NumStoriesUnstarted = numStoriesUnstarted
	this.NumPointsStarted = numPointsStarted
	this.NumPointsUnstarted = numPointsUnstarted
	this.NumStoriesStarted = numStoriesStarted
	this.NumStoriesUnestimated = numStoriesUnestimated
	this.NumStoriesBacklog = numStoriesBacklog
	this.NumPointsBacklog = numPointsBacklog
	this.NumPoints = numPoints
	this.NumStoriesDone = numStoriesDone
	return &this
}

// NewIterationStatsWithDefaults instantiates a new IterationStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIterationStatsWithDefaults() *IterationStats {
	this := IterationStats{}
	return &this
}

// GetNumPointsDone returns the NumPointsDone field value
func (o *IterationStats) GetNumPointsDone() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumPointsDone
}

// GetNumPointsDoneOk returns a tuple with the NumPointsDone field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumPointsDoneOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPointsDone, true
}

// SetNumPointsDone sets field value
func (o *IterationStats) SetNumPointsDone(v int64) {
	o.NumPointsDone = v
}

// GetNumRelatedDocuments returns the NumRelatedDocuments field value
func (o *IterationStats) GetNumRelatedDocuments() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumRelatedDocuments
}

// GetNumRelatedDocumentsOk returns a tuple with the NumRelatedDocuments field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumRelatedDocumentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumRelatedDocuments, true
}

// SetNumRelatedDocuments sets field value
func (o *IterationStats) SetNumRelatedDocuments(v int64) {
	o.NumRelatedDocuments = v
}

// GetAverageCycleTime returns the AverageCycleTime field value if set, zero value otherwise.
func (o *IterationStats) GetAverageCycleTime() int64 {
	if o == nil || IsNil(o.AverageCycleTime) {
		var ret int64
		return ret
	}
	return *o.AverageCycleTime
}

// GetAverageCycleTimeOk returns a tuple with the AverageCycleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IterationStats) GetAverageCycleTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AverageCycleTime) {
		return nil, false
	}
	return o.AverageCycleTime, true
}

// HasAverageCycleTime returns a boolean if a field has been set.
func (o *IterationStats) HasAverageCycleTime() bool {
	if o != nil && !IsNil(o.AverageCycleTime) {
		return true
	}

	return false
}

// SetAverageCycleTime gets a reference to the given int64 and assigns it to the AverageCycleTime field.
func (o *IterationStats) SetAverageCycleTime(v int64) {
	o.AverageCycleTime = &v
}

// GetNumStoriesUnstarted returns the NumStoriesUnstarted field value
func (o *IterationStats) GetNumStoriesUnstarted() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumStoriesUnstarted
}

// GetNumStoriesUnstartedOk returns a tuple with the NumStoriesUnstarted field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumStoriesUnstartedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumStoriesUnstarted, true
}

// SetNumStoriesUnstarted sets field value
func (o *IterationStats) SetNumStoriesUnstarted(v int64) {
	o.NumStoriesUnstarted = v
}

// GetNumPointsStarted returns the NumPointsStarted field value
func (o *IterationStats) GetNumPointsStarted() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumPointsStarted
}

// GetNumPointsStartedOk returns a tuple with the NumPointsStarted field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumPointsStartedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPointsStarted, true
}

// SetNumPointsStarted sets field value
func (o *IterationStats) SetNumPointsStarted(v int64) {
	o.NumPointsStarted = v
}

// GetNumPointsUnstarted returns the NumPointsUnstarted field value
func (o *IterationStats) GetNumPointsUnstarted() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumPointsUnstarted
}

// GetNumPointsUnstartedOk returns a tuple with the NumPointsUnstarted field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumPointsUnstartedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPointsUnstarted, true
}

// SetNumPointsUnstarted sets field value
func (o *IterationStats) SetNumPointsUnstarted(v int64) {
	o.NumPointsUnstarted = v
}

// GetNumStoriesStarted returns the NumStoriesStarted field value
func (o *IterationStats) GetNumStoriesStarted() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumStoriesStarted
}

// GetNumStoriesStartedOk returns a tuple with the NumStoriesStarted field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumStoriesStartedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumStoriesStarted, true
}

// SetNumStoriesStarted sets field value
func (o *IterationStats) SetNumStoriesStarted(v int64) {
	o.NumStoriesStarted = v
}

// GetNumStoriesUnestimated returns the NumStoriesUnestimated field value
func (o *IterationStats) GetNumStoriesUnestimated() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumStoriesUnestimated
}

// GetNumStoriesUnestimatedOk returns a tuple with the NumStoriesUnestimated field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumStoriesUnestimatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumStoriesUnestimated, true
}

// SetNumStoriesUnestimated sets field value
func (o *IterationStats) SetNumStoriesUnestimated(v int64) {
	o.NumStoriesUnestimated = v
}

// GetNumStoriesBacklog returns the NumStoriesBacklog field value
func (o *IterationStats) GetNumStoriesBacklog() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumStoriesBacklog
}

// GetNumStoriesBacklogOk returns a tuple with the NumStoriesBacklog field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumStoriesBacklogOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumStoriesBacklog, true
}

// SetNumStoriesBacklog sets field value
func (o *IterationStats) SetNumStoriesBacklog(v int64) {
	o.NumStoriesBacklog = v
}

// GetAverageLeadTime returns the AverageLeadTime field value if set, zero value otherwise.
func (o *IterationStats) GetAverageLeadTime() int64 {
	if o == nil || IsNil(o.AverageLeadTime) {
		var ret int64
		return ret
	}
	return *o.AverageLeadTime
}

// GetAverageLeadTimeOk returns a tuple with the AverageLeadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IterationStats) GetAverageLeadTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AverageLeadTime) {
		return nil, false
	}
	return o.AverageLeadTime, true
}

// HasAverageLeadTime returns a boolean if a field has been set.
func (o *IterationStats) HasAverageLeadTime() bool {
	if o != nil && !IsNil(o.AverageLeadTime) {
		return true
	}

	return false
}

// SetAverageLeadTime gets a reference to the given int64 and assigns it to the AverageLeadTime field.
func (o *IterationStats) SetAverageLeadTime(v int64) {
	o.AverageLeadTime = &v
}

// GetNumPointsBacklog returns the NumPointsBacklog field value
func (o *IterationStats) GetNumPointsBacklog() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumPointsBacklog
}

// GetNumPointsBacklogOk returns a tuple with the NumPointsBacklog field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumPointsBacklogOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPointsBacklog, true
}

// SetNumPointsBacklog sets field value
func (o *IterationStats) SetNumPointsBacklog(v int64) {
	o.NumPointsBacklog = v
}

// GetNumPoints returns the NumPoints field value
func (o *IterationStats) GetNumPoints() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumPoints
}

// GetNumPointsOk returns a tuple with the NumPoints field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumPointsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPoints, true
}

// SetNumPoints sets field value
func (o *IterationStats) SetNumPoints(v int64) {
	o.NumPoints = v
}

// GetNumStoriesDone returns the NumStoriesDone field value
func (o *IterationStats) GetNumStoriesDone() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumStoriesDone
}

// GetNumStoriesDoneOk returns a tuple with the NumStoriesDone field value
// and a boolean to check if the value has been set.
func (o *IterationStats) GetNumStoriesDoneOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumStoriesDone, true
}

// SetNumStoriesDone sets field value
func (o *IterationStats) SetNumStoriesDone(v int64) {
	o.NumStoriesDone = v
}

func (o IterationStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IterationStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["num_points_done"] = o.NumPointsDone
	toSerialize["num_related_documents"] = o.NumRelatedDocuments
	if !IsNil(o.AverageCycleTime) {
		toSerialize["average_cycle_time"] = o.AverageCycleTime
	}
	toSerialize["num_stories_unstarted"] = o.NumStoriesUnstarted
	toSerialize["num_points_started"] = o.NumPointsStarted
	toSerialize["num_points_unstarted"] = o.NumPointsUnstarted
	toSerialize["num_stories_started"] = o.NumStoriesStarted
	toSerialize["num_stories_unestimated"] = o.NumStoriesUnestimated
	toSerialize["num_stories_backlog"] = o.NumStoriesBacklog
	if !IsNil(o.AverageLeadTime) {
		toSerialize["average_lead_time"] = o.AverageLeadTime
	}
	toSerialize["num_points_backlog"] = o.NumPointsBacklog
	toSerialize["num_points"] = o.NumPoints
	toSerialize["num_stories_done"] = o.NumStoriesDone
	return toSerialize, nil
}

type NullableIterationStats struct {
	value *IterationStats
	isSet bool
}

func (v NullableIterationStats) Get() *IterationStats {
	return v.value
}

func (v *NullableIterationStats) Set(val *IterationStats) {
	v.value = val
	v.isSet = true
}

func (v NullableIterationStats) IsSet() bool {
	return v.isSet
}

func (v *NullableIterationStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIterationStats(val *IterationStats) *NullableIterationStats {
	return &NullableIterationStats{value: val, isSet: true}
}

func (v NullableIterationStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIterationStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


