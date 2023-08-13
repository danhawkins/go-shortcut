# EpicStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumPointsDone** | **int64** | The total number of completed points in this Epic. | 
**NumRelatedDocuments** | **int64** | The total number of documents associated with this Epic. | 
**AverageCycleTime** | Pointer to **int64** | The average cycle time (in seconds) of completed stories in this Epic. | [optional] 
**NumStoriesUnstarted** | **int64** | The total number of unstarted Stories in this Epic. | 
**NumStoriesTotal** | **int64** | The total number of Stories in this Epic. | 
**LastStoryUpdate** | **NullableTime** | The date of the last update of a Story in this Epic. | 
**NumPointsStarted** | **int64** | The total number of started points in this Epic. | 
**NumPointsUnstarted** | **int64** | The total number of unstarted points in this Epic. | 
**NumStoriesStarted** | **int64** | The total number of started Stories in this Epic. | 
**NumStoriesUnestimated** | **int64** | The total number of Stories with no point estimate. | 
**NumStoriesBacklog** | **int64** | The total number of backlog Stories in this Epic. | 
**AverageLeadTime** | Pointer to **int64** | The average lead time (in seconds) of completed stories in this Epic. | [optional] 
**NumPointsBacklog** | **int64** | The total number of backlog points in this Epic. | 
**NumPoints** | **int64** | The total number of points in this Epic. | 
**NumStoriesDone** | **int64** | The total number of done Stories in this Epic. | 

## Methods

### NewEpicStats

`func NewEpicStats(numPointsDone int64, numRelatedDocuments int64, numStoriesUnstarted int64, numStoriesTotal int64, lastStoryUpdate NullableTime, numPointsStarted int64, numPointsUnstarted int64, numStoriesStarted int64, numStoriesUnestimated int64, numStoriesBacklog int64, numPointsBacklog int64, numPoints int64, numStoriesDone int64, ) *EpicStats`

NewEpicStats instantiates a new EpicStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpicStatsWithDefaults

`func NewEpicStatsWithDefaults() *EpicStats`

NewEpicStatsWithDefaults instantiates a new EpicStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumPointsDone

`func (o *EpicStats) GetNumPointsDone() int64`

GetNumPointsDone returns the NumPointsDone field if non-nil, zero value otherwise.

### GetNumPointsDoneOk

`func (o *EpicStats) GetNumPointsDoneOk() (*int64, bool)`

GetNumPointsDoneOk returns a tuple with the NumPointsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsDone

`func (o *EpicStats) SetNumPointsDone(v int64)`

SetNumPointsDone sets NumPointsDone field to given value.


### GetNumRelatedDocuments

`func (o *EpicStats) GetNumRelatedDocuments() int64`

GetNumRelatedDocuments returns the NumRelatedDocuments field if non-nil, zero value otherwise.

### GetNumRelatedDocumentsOk

`func (o *EpicStats) GetNumRelatedDocumentsOk() (*int64, bool)`

GetNumRelatedDocumentsOk returns a tuple with the NumRelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRelatedDocuments

`func (o *EpicStats) SetNumRelatedDocuments(v int64)`

SetNumRelatedDocuments sets NumRelatedDocuments field to given value.


### GetAverageCycleTime

`func (o *EpicStats) GetAverageCycleTime() int64`

GetAverageCycleTime returns the AverageCycleTime field if non-nil, zero value otherwise.

### GetAverageCycleTimeOk

`func (o *EpicStats) GetAverageCycleTimeOk() (*int64, bool)`

GetAverageCycleTimeOk returns a tuple with the AverageCycleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCycleTime

`func (o *EpicStats) SetAverageCycleTime(v int64)`

SetAverageCycleTime sets AverageCycleTime field to given value.

### HasAverageCycleTime

`func (o *EpicStats) HasAverageCycleTime() bool`

HasAverageCycleTime returns a boolean if a field has been set.

### GetNumStoriesUnstarted

`func (o *EpicStats) GetNumStoriesUnstarted() int64`

GetNumStoriesUnstarted returns the NumStoriesUnstarted field if non-nil, zero value otherwise.

### GetNumStoriesUnstartedOk

`func (o *EpicStats) GetNumStoriesUnstartedOk() (*int64, bool)`

GetNumStoriesUnstartedOk returns a tuple with the NumStoriesUnstarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesUnstarted

`func (o *EpicStats) SetNumStoriesUnstarted(v int64)`

SetNumStoriesUnstarted sets NumStoriesUnstarted field to given value.


### GetNumStoriesTotal

`func (o *EpicStats) GetNumStoriesTotal() int64`

GetNumStoriesTotal returns the NumStoriesTotal field if non-nil, zero value otherwise.

### GetNumStoriesTotalOk

`func (o *EpicStats) GetNumStoriesTotalOk() (*int64, bool)`

GetNumStoriesTotalOk returns a tuple with the NumStoriesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesTotal

`func (o *EpicStats) SetNumStoriesTotal(v int64)`

SetNumStoriesTotal sets NumStoriesTotal field to given value.


### GetLastStoryUpdate

`func (o *EpicStats) GetLastStoryUpdate() time.Time`

GetLastStoryUpdate returns the LastStoryUpdate field if non-nil, zero value otherwise.

### GetLastStoryUpdateOk

`func (o *EpicStats) GetLastStoryUpdateOk() (*time.Time, bool)`

GetLastStoryUpdateOk returns a tuple with the LastStoryUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStoryUpdate

`func (o *EpicStats) SetLastStoryUpdate(v time.Time)`

SetLastStoryUpdate sets LastStoryUpdate field to given value.


### SetLastStoryUpdateNil

`func (o *EpicStats) SetLastStoryUpdateNil(b bool)`

 SetLastStoryUpdateNil sets the value for LastStoryUpdate to be an explicit nil

### UnsetLastStoryUpdate
`func (o *EpicStats) UnsetLastStoryUpdate()`

UnsetLastStoryUpdate ensures that no value is present for LastStoryUpdate, not even an explicit nil
### GetNumPointsStarted

`func (o *EpicStats) GetNumPointsStarted() int64`

GetNumPointsStarted returns the NumPointsStarted field if non-nil, zero value otherwise.

### GetNumPointsStartedOk

`func (o *EpicStats) GetNumPointsStartedOk() (*int64, bool)`

GetNumPointsStartedOk returns a tuple with the NumPointsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsStarted

`func (o *EpicStats) SetNumPointsStarted(v int64)`

SetNumPointsStarted sets NumPointsStarted field to given value.


### GetNumPointsUnstarted

`func (o *EpicStats) GetNumPointsUnstarted() int64`

GetNumPointsUnstarted returns the NumPointsUnstarted field if non-nil, zero value otherwise.

### GetNumPointsUnstartedOk

`func (o *EpicStats) GetNumPointsUnstartedOk() (*int64, bool)`

GetNumPointsUnstartedOk returns a tuple with the NumPointsUnstarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsUnstarted

`func (o *EpicStats) SetNumPointsUnstarted(v int64)`

SetNumPointsUnstarted sets NumPointsUnstarted field to given value.


### GetNumStoriesStarted

`func (o *EpicStats) GetNumStoriesStarted() int64`

GetNumStoriesStarted returns the NumStoriesStarted field if non-nil, zero value otherwise.

### GetNumStoriesStartedOk

`func (o *EpicStats) GetNumStoriesStartedOk() (*int64, bool)`

GetNumStoriesStartedOk returns a tuple with the NumStoriesStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesStarted

`func (o *EpicStats) SetNumStoriesStarted(v int64)`

SetNumStoriesStarted sets NumStoriesStarted field to given value.


### GetNumStoriesUnestimated

`func (o *EpicStats) GetNumStoriesUnestimated() int64`

GetNumStoriesUnestimated returns the NumStoriesUnestimated field if non-nil, zero value otherwise.

### GetNumStoriesUnestimatedOk

`func (o *EpicStats) GetNumStoriesUnestimatedOk() (*int64, bool)`

GetNumStoriesUnestimatedOk returns a tuple with the NumStoriesUnestimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesUnestimated

`func (o *EpicStats) SetNumStoriesUnestimated(v int64)`

SetNumStoriesUnestimated sets NumStoriesUnestimated field to given value.


### GetNumStoriesBacklog

`func (o *EpicStats) GetNumStoriesBacklog() int64`

GetNumStoriesBacklog returns the NumStoriesBacklog field if non-nil, zero value otherwise.

### GetNumStoriesBacklogOk

`func (o *EpicStats) GetNumStoriesBacklogOk() (*int64, bool)`

GetNumStoriesBacklogOk returns a tuple with the NumStoriesBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesBacklog

`func (o *EpicStats) SetNumStoriesBacklog(v int64)`

SetNumStoriesBacklog sets NumStoriesBacklog field to given value.


### GetAverageLeadTime

`func (o *EpicStats) GetAverageLeadTime() int64`

GetAverageLeadTime returns the AverageLeadTime field if non-nil, zero value otherwise.

### GetAverageLeadTimeOk

`func (o *EpicStats) GetAverageLeadTimeOk() (*int64, bool)`

GetAverageLeadTimeOk returns a tuple with the AverageLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageLeadTime

`func (o *EpicStats) SetAverageLeadTime(v int64)`

SetAverageLeadTime sets AverageLeadTime field to given value.

### HasAverageLeadTime

`func (o *EpicStats) HasAverageLeadTime() bool`

HasAverageLeadTime returns a boolean if a field has been set.

### GetNumPointsBacklog

`func (o *EpicStats) GetNumPointsBacklog() int64`

GetNumPointsBacklog returns the NumPointsBacklog field if non-nil, zero value otherwise.

### GetNumPointsBacklogOk

`func (o *EpicStats) GetNumPointsBacklogOk() (*int64, bool)`

GetNumPointsBacklogOk returns a tuple with the NumPointsBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsBacklog

`func (o *EpicStats) SetNumPointsBacklog(v int64)`

SetNumPointsBacklog sets NumPointsBacklog field to given value.


### GetNumPoints

`func (o *EpicStats) GetNumPoints() int64`

GetNumPoints returns the NumPoints field if non-nil, zero value otherwise.

### GetNumPointsOk

`func (o *EpicStats) GetNumPointsOk() (*int64, bool)`

GetNumPointsOk returns a tuple with the NumPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPoints

`func (o *EpicStats) SetNumPoints(v int64)`

SetNumPoints sets NumPoints field to given value.


### GetNumStoriesDone

`func (o *EpicStats) GetNumStoriesDone() int64`

GetNumStoriesDone returns the NumStoriesDone field if non-nil, zero value otherwise.

### GetNumStoriesDoneOk

`func (o *EpicStats) GetNumStoriesDoneOk() (*int64, bool)`

GetNumStoriesDoneOk returns a tuple with the NumStoriesDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesDone

`func (o *EpicStats) SetNumStoriesDone(v int64)`

SetNumStoriesDone sets NumStoriesDone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


