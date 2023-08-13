# IterationStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumPointsDone** | **int64** | The total number of completed points in this Iteration. | 
**NumRelatedDocuments** | **int64** | The total number of documents related to an Iteration | 
**AverageCycleTime** | Pointer to **int64** | The average cycle time (in seconds) of completed stories in this Iteration. | [optional] 
**NumStoriesUnstarted** | **int64** | The total number of unstarted Stories in this Iteration. | 
**NumPointsStarted** | **int64** | The total number of started points in this Iteration. | 
**NumPointsUnstarted** | **int64** | The total number of unstarted points in this Iteration. | 
**NumStoriesStarted** | **int64** | The total number of started Stories in this Iteration. | 
**NumStoriesUnestimated** | **int64** | The total number of Stories with no point estimate. | 
**NumStoriesBacklog** | **int64** | The total number of backlog Stories in this Iteration. | 
**AverageLeadTime** | Pointer to **int64** | The average lead time (in seconds) of completed stories in this Iteration. | [optional] 
**NumPointsBacklog** | **int64** | The total number of backlog points in this Iteration. | 
**NumPoints** | **int64** | The total number of points in this Iteration. | 
**NumStoriesDone** | **int64** | The total number of done Stories in this Iteration. | 

## Methods

### NewIterationStats

`func NewIterationStats(numPointsDone int64, numRelatedDocuments int64, numStoriesUnstarted int64, numPointsStarted int64, numPointsUnstarted int64, numStoriesStarted int64, numStoriesUnestimated int64, numStoriesBacklog int64, numPointsBacklog int64, numPoints int64, numStoriesDone int64, ) *IterationStats`

NewIterationStats instantiates a new IterationStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationStatsWithDefaults

`func NewIterationStatsWithDefaults() *IterationStats`

NewIterationStatsWithDefaults instantiates a new IterationStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumPointsDone

`func (o *IterationStats) GetNumPointsDone() int64`

GetNumPointsDone returns the NumPointsDone field if non-nil, zero value otherwise.

### GetNumPointsDoneOk

`func (o *IterationStats) GetNumPointsDoneOk() (*int64, bool)`

GetNumPointsDoneOk returns a tuple with the NumPointsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsDone

`func (o *IterationStats) SetNumPointsDone(v int64)`

SetNumPointsDone sets NumPointsDone field to given value.


### GetNumRelatedDocuments

`func (o *IterationStats) GetNumRelatedDocuments() int64`

GetNumRelatedDocuments returns the NumRelatedDocuments field if non-nil, zero value otherwise.

### GetNumRelatedDocumentsOk

`func (o *IterationStats) GetNumRelatedDocumentsOk() (*int64, bool)`

GetNumRelatedDocumentsOk returns a tuple with the NumRelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRelatedDocuments

`func (o *IterationStats) SetNumRelatedDocuments(v int64)`

SetNumRelatedDocuments sets NumRelatedDocuments field to given value.


### GetAverageCycleTime

`func (o *IterationStats) GetAverageCycleTime() int64`

GetAverageCycleTime returns the AverageCycleTime field if non-nil, zero value otherwise.

### GetAverageCycleTimeOk

`func (o *IterationStats) GetAverageCycleTimeOk() (*int64, bool)`

GetAverageCycleTimeOk returns a tuple with the AverageCycleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCycleTime

`func (o *IterationStats) SetAverageCycleTime(v int64)`

SetAverageCycleTime sets AverageCycleTime field to given value.

### HasAverageCycleTime

`func (o *IterationStats) HasAverageCycleTime() bool`

HasAverageCycleTime returns a boolean if a field has been set.

### GetNumStoriesUnstarted

`func (o *IterationStats) GetNumStoriesUnstarted() int64`

GetNumStoriesUnstarted returns the NumStoriesUnstarted field if non-nil, zero value otherwise.

### GetNumStoriesUnstartedOk

`func (o *IterationStats) GetNumStoriesUnstartedOk() (*int64, bool)`

GetNumStoriesUnstartedOk returns a tuple with the NumStoriesUnstarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesUnstarted

`func (o *IterationStats) SetNumStoriesUnstarted(v int64)`

SetNumStoriesUnstarted sets NumStoriesUnstarted field to given value.


### GetNumPointsStarted

`func (o *IterationStats) GetNumPointsStarted() int64`

GetNumPointsStarted returns the NumPointsStarted field if non-nil, zero value otherwise.

### GetNumPointsStartedOk

`func (o *IterationStats) GetNumPointsStartedOk() (*int64, bool)`

GetNumPointsStartedOk returns a tuple with the NumPointsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsStarted

`func (o *IterationStats) SetNumPointsStarted(v int64)`

SetNumPointsStarted sets NumPointsStarted field to given value.


### GetNumPointsUnstarted

`func (o *IterationStats) GetNumPointsUnstarted() int64`

GetNumPointsUnstarted returns the NumPointsUnstarted field if non-nil, zero value otherwise.

### GetNumPointsUnstartedOk

`func (o *IterationStats) GetNumPointsUnstartedOk() (*int64, bool)`

GetNumPointsUnstartedOk returns a tuple with the NumPointsUnstarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsUnstarted

`func (o *IterationStats) SetNumPointsUnstarted(v int64)`

SetNumPointsUnstarted sets NumPointsUnstarted field to given value.


### GetNumStoriesStarted

`func (o *IterationStats) GetNumStoriesStarted() int64`

GetNumStoriesStarted returns the NumStoriesStarted field if non-nil, zero value otherwise.

### GetNumStoriesStartedOk

`func (o *IterationStats) GetNumStoriesStartedOk() (*int64, bool)`

GetNumStoriesStartedOk returns a tuple with the NumStoriesStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesStarted

`func (o *IterationStats) SetNumStoriesStarted(v int64)`

SetNumStoriesStarted sets NumStoriesStarted field to given value.


### GetNumStoriesUnestimated

`func (o *IterationStats) GetNumStoriesUnestimated() int64`

GetNumStoriesUnestimated returns the NumStoriesUnestimated field if non-nil, zero value otherwise.

### GetNumStoriesUnestimatedOk

`func (o *IterationStats) GetNumStoriesUnestimatedOk() (*int64, bool)`

GetNumStoriesUnestimatedOk returns a tuple with the NumStoriesUnestimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesUnestimated

`func (o *IterationStats) SetNumStoriesUnestimated(v int64)`

SetNumStoriesUnestimated sets NumStoriesUnestimated field to given value.


### GetNumStoriesBacklog

`func (o *IterationStats) GetNumStoriesBacklog() int64`

GetNumStoriesBacklog returns the NumStoriesBacklog field if non-nil, zero value otherwise.

### GetNumStoriesBacklogOk

`func (o *IterationStats) GetNumStoriesBacklogOk() (*int64, bool)`

GetNumStoriesBacklogOk returns a tuple with the NumStoriesBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesBacklog

`func (o *IterationStats) SetNumStoriesBacklog(v int64)`

SetNumStoriesBacklog sets NumStoriesBacklog field to given value.


### GetAverageLeadTime

`func (o *IterationStats) GetAverageLeadTime() int64`

GetAverageLeadTime returns the AverageLeadTime field if non-nil, zero value otherwise.

### GetAverageLeadTimeOk

`func (o *IterationStats) GetAverageLeadTimeOk() (*int64, bool)`

GetAverageLeadTimeOk returns a tuple with the AverageLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageLeadTime

`func (o *IterationStats) SetAverageLeadTime(v int64)`

SetAverageLeadTime sets AverageLeadTime field to given value.

### HasAverageLeadTime

`func (o *IterationStats) HasAverageLeadTime() bool`

HasAverageLeadTime returns a boolean if a field has been set.

### GetNumPointsBacklog

`func (o *IterationStats) GetNumPointsBacklog() int64`

GetNumPointsBacklog returns the NumPointsBacklog field if non-nil, zero value otherwise.

### GetNumPointsBacklogOk

`func (o *IterationStats) GetNumPointsBacklogOk() (*int64, bool)`

GetNumPointsBacklogOk returns a tuple with the NumPointsBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsBacklog

`func (o *IterationStats) SetNumPointsBacklog(v int64)`

SetNumPointsBacklog sets NumPointsBacklog field to given value.


### GetNumPoints

`func (o *IterationStats) GetNumPoints() int64`

GetNumPoints returns the NumPoints field if non-nil, zero value otherwise.

### GetNumPointsOk

`func (o *IterationStats) GetNumPointsOk() (*int64, bool)`

GetNumPointsOk returns a tuple with the NumPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPoints

`func (o *IterationStats) SetNumPoints(v int64)`

SetNumPoints sets NumPoints field to given value.


### GetNumStoriesDone

`func (o *IterationStats) GetNumStoriesDone() int64`

GetNumStoriesDone returns the NumStoriesDone field if non-nil, zero value otherwise.

### GetNumStoriesDoneOk

`func (o *IterationStats) GetNumStoriesDoneOk() (*int64, bool)`

GetNumStoriesDoneOk returns a tuple with the NumStoriesDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesDone

`func (o *IterationStats) SetNumStoriesDone(v int64)`

SetNumStoriesDone sets NumStoriesDone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


