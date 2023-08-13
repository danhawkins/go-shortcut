# LabelStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumRelatedDocuments** | **int64** | The total number of Documents associated this Label. | 
**NumEpics** | **int64** | The total number of Epics with this Label. | 
**NumStoriesUnstarted** | **int64** | The total number of stories unstarted Stories with this Label. | 
**NumStoriesTotal** | **int64** | The total number of Stories with this Label. | 
**NumEpicsUnstarted** | **int64** | The number of unstarted epics associated with this label. | 
**NumEpicsInProgress** | **int64** | The number of in progress epics associated with this label. | 
**NumPointsUnstarted** | **int64** | The total number of unstarted points with this Label. | 
**NumStoriesUnestimated** | **int64** | The total number of Stories with no point estimate with this Label. | 
**NumPointsInProgress** | **int64** | The total number of in-progress points with this Label. | 
**NumEpicsTotal** | **int64** | The total number of Epics associated with this Label. | 
**NumStoriesCompleted** | **int64** | The total number of completed Stories with this Label. | 
**NumPointsCompleted** | **int64** | The total number of completed points with this Label. | 
**NumStoriesBacklog** | **int64** | The total number of stories backlog Stories with this Label. | 
**NumPointsTotal** | **int64** | The total number of points with this Label. | 
**NumStoriesInProgress** | **int64** | The total number of in-progress Stories with this Label. | 
**NumPointsBacklog** | **int64** | The total number of backlog points with this Label. | 
**NumEpicsCompleted** | **int64** | The number of completed Epics associated with this Label. | 

## Methods

### NewLabelStats

`func NewLabelStats(numRelatedDocuments int64, numEpics int64, numStoriesUnstarted int64, numStoriesTotal int64, numEpicsUnstarted int64, numEpicsInProgress int64, numPointsUnstarted int64, numStoriesUnestimated int64, numPointsInProgress int64, numEpicsTotal int64, numStoriesCompleted int64, numPointsCompleted int64, numStoriesBacklog int64, numPointsTotal int64, numStoriesInProgress int64, numPointsBacklog int64, numEpicsCompleted int64, ) *LabelStats`

NewLabelStats instantiates a new LabelStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelStatsWithDefaults

`func NewLabelStatsWithDefaults() *LabelStats`

NewLabelStatsWithDefaults instantiates a new LabelStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumRelatedDocuments

`func (o *LabelStats) GetNumRelatedDocuments() int64`

GetNumRelatedDocuments returns the NumRelatedDocuments field if non-nil, zero value otherwise.

### GetNumRelatedDocumentsOk

`func (o *LabelStats) GetNumRelatedDocumentsOk() (*int64, bool)`

GetNumRelatedDocumentsOk returns a tuple with the NumRelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRelatedDocuments

`func (o *LabelStats) SetNumRelatedDocuments(v int64)`

SetNumRelatedDocuments sets NumRelatedDocuments field to given value.


### GetNumEpics

`func (o *LabelStats) GetNumEpics() int64`

GetNumEpics returns the NumEpics field if non-nil, zero value otherwise.

### GetNumEpicsOk

`func (o *LabelStats) GetNumEpicsOk() (*int64, bool)`

GetNumEpicsOk returns a tuple with the NumEpics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEpics

`func (o *LabelStats) SetNumEpics(v int64)`

SetNumEpics sets NumEpics field to given value.


### GetNumStoriesUnstarted

`func (o *LabelStats) GetNumStoriesUnstarted() int64`

GetNumStoriesUnstarted returns the NumStoriesUnstarted field if non-nil, zero value otherwise.

### GetNumStoriesUnstartedOk

`func (o *LabelStats) GetNumStoriesUnstartedOk() (*int64, bool)`

GetNumStoriesUnstartedOk returns a tuple with the NumStoriesUnstarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesUnstarted

`func (o *LabelStats) SetNumStoriesUnstarted(v int64)`

SetNumStoriesUnstarted sets NumStoriesUnstarted field to given value.


### GetNumStoriesTotal

`func (o *LabelStats) GetNumStoriesTotal() int64`

GetNumStoriesTotal returns the NumStoriesTotal field if non-nil, zero value otherwise.

### GetNumStoriesTotalOk

`func (o *LabelStats) GetNumStoriesTotalOk() (*int64, bool)`

GetNumStoriesTotalOk returns a tuple with the NumStoriesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesTotal

`func (o *LabelStats) SetNumStoriesTotal(v int64)`

SetNumStoriesTotal sets NumStoriesTotal field to given value.


### GetNumEpicsUnstarted

`func (o *LabelStats) GetNumEpicsUnstarted() int64`

GetNumEpicsUnstarted returns the NumEpicsUnstarted field if non-nil, zero value otherwise.

### GetNumEpicsUnstartedOk

`func (o *LabelStats) GetNumEpicsUnstartedOk() (*int64, bool)`

GetNumEpicsUnstartedOk returns a tuple with the NumEpicsUnstarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEpicsUnstarted

`func (o *LabelStats) SetNumEpicsUnstarted(v int64)`

SetNumEpicsUnstarted sets NumEpicsUnstarted field to given value.


### GetNumEpicsInProgress

`func (o *LabelStats) GetNumEpicsInProgress() int64`

GetNumEpicsInProgress returns the NumEpicsInProgress field if non-nil, zero value otherwise.

### GetNumEpicsInProgressOk

`func (o *LabelStats) GetNumEpicsInProgressOk() (*int64, bool)`

GetNumEpicsInProgressOk returns a tuple with the NumEpicsInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEpicsInProgress

`func (o *LabelStats) SetNumEpicsInProgress(v int64)`

SetNumEpicsInProgress sets NumEpicsInProgress field to given value.


### GetNumPointsUnstarted

`func (o *LabelStats) GetNumPointsUnstarted() int64`

GetNumPointsUnstarted returns the NumPointsUnstarted field if non-nil, zero value otherwise.

### GetNumPointsUnstartedOk

`func (o *LabelStats) GetNumPointsUnstartedOk() (*int64, bool)`

GetNumPointsUnstartedOk returns a tuple with the NumPointsUnstarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsUnstarted

`func (o *LabelStats) SetNumPointsUnstarted(v int64)`

SetNumPointsUnstarted sets NumPointsUnstarted field to given value.


### GetNumStoriesUnestimated

`func (o *LabelStats) GetNumStoriesUnestimated() int64`

GetNumStoriesUnestimated returns the NumStoriesUnestimated field if non-nil, zero value otherwise.

### GetNumStoriesUnestimatedOk

`func (o *LabelStats) GetNumStoriesUnestimatedOk() (*int64, bool)`

GetNumStoriesUnestimatedOk returns a tuple with the NumStoriesUnestimated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesUnestimated

`func (o *LabelStats) SetNumStoriesUnestimated(v int64)`

SetNumStoriesUnestimated sets NumStoriesUnestimated field to given value.


### GetNumPointsInProgress

`func (o *LabelStats) GetNumPointsInProgress() int64`

GetNumPointsInProgress returns the NumPointsInProgress field if non-nil, zero value otherwise.

### GetNumPointsInProgressOk

`func (o *LabelStats) GetNumPointsInProgressOk() (*int64, bool)`

GetNumPointsInProgressOk returns a tuple with the NumPointsInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsInProgress

`func (o *LabelStats) SetNumPointsInProgress(v int64)`

SetNumPointsInProgress sets NumPointsInProgress field to given value.


### GetNumEpicsTotal

`func (o *LabelStats) GetNumEpicsTotal() int64`

GetNumEpicsTotal returns the NumEpicsTotal field if non-nil, zero value otherwise.

### GetNumEpicsTotalOk

`func (o *LabelStats) GetNumEpicsTotalOk() (*int64, bool)`

GetNumEpicsTotalOk returns a tuple with the NumEpicsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEpicsTotal

`func (o *LabelStats) SetNumEpicsTotal(v int64)`

SetNumEpicsTotal sets NumEpicsTotal field to given value.


### GetNumStoriesCompleted

`func (o *LabelStats) GetNumStoriesCompleted() int64`

GetNumStoriesCompleted returns the NumStoriesCompleted field if non-nil, zero value otherwise.

### GetNumStoriesCompletedOk

`func (o *LabelStats) GetNumStoriesCompletedOk() (*int64, bool)`

GetNumStoriesCompletedOk returns a tuple with the NumStoriesCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesCompleted

`func (o *LabelStats) SetNumStoriesCompleted(v int64)`

SetNumStoriesCompleted sets NumStoriesCompleted field to given value.


### GetNumPointsCompleted

`func (o *LabelStats) GetNumPointsCompleted() int64`

GetNumPointsCompleted returns the NumPointsCompleted field if non-nil, zero value otherwise.

### GetNumPointsCompletedOk

`func (o *LabelStats) GetNumPointsCompletedOk() (*int64, bool)`

GetNumPointsCompletedOk returns a tuple with the NumPointsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsCompleted

`func (o *LabelStats) SetNumPointsCompleted(v int64)`

SetNumPointsCompleted sets NumPointsCompleted field to given value.


### GetNumStoriesBacklog

`func (o *LabelStats) GetNumStoriesBacklog() int64`

GetNumStoriesBacklog returns the NumStoriesBacklog field if non-nil, zero value otherwise.

### GetNumStoriesBacklogOk

`func (o *LabelStats) GetNumStoriesBacklogOk() (*int64, bool)`

GetNumStoriesBacklogOk returns a tuple with the NumStoriesBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesBacklog

`func (o *LabelStats) SetNumStoriesBacklog(v int64)`

SetNumStoriesBacklog sets NumStoriesBacklog field to given value.


### GetNumPointsTotal

`func (o *LabelStats) GetNumPointsTotal() int64`

GetNumPointsTotal returns the NumPointsTotal field if non-nil, zero value otherwise.

### GetNumPointsTotalOk

`func (o *LabelStats) GetNumPointsTotalOk() (*int64, bool)`

GetNumPointsTotalOk returns a tuple with the NumPointsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsTotal

`func (o *LabelStats) SetNumPointsTotal(v int64)`

SetNumPointsTotal sets NumPointsTotal field to given value.


### GetNumStoriesInProgress

`func (o *LabelStats) GetNumStoriesInProgress() int64`

GetNumStoriesInProgress returns the NumStoriesInProgress field if non-nil, zero value otherwise.

### GetNumStoriesInProgressOk

`func (o *LabelStats) GetNumStoriesInProgressOk() (*int64, bool)`

GetNumStoriesInProgressOk returns a tuple with the NumStoriesInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesInProgress

`func (o *LabelStats) SetNumStoriesInProgress(v int64)`

SetNumStoriesInProgress sets NumStoriesInProgress field to given value.


### GetNumPointsBacklog

`func (o *LabelStats) GetNumPointsBacklog() int64`

GetNumPointsBacklog returns the NumPointsBacklog field if non-nil, zero value otherwise.

### GetNumPointsBacklogOk

`func (o *LabelStats) GetNumPointsBacklogOk() (*int64, bool)`

GetNumPointsBacklogOk returns a tuple with the NumPointsBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPointsBacklog

`func (o *LabelStats) SetNumPointsBacklog(v int64)`

SetNumPointsBacklog sets NumPointsBacklog field to given value.


### GetNumEpicsCompleted

`func (o *LabelStats) GetNumEpicsCompleted() int64`

GetNumEpicsCompleted returns the NumEpicsCompleted field if non-nil, zero value otherwise.

### GetNumEpicsCompletedOk

`func (o *LabelStats) GetNumEpicsCompletedOk() (*int64, bool)`

GetNumEpicsCompletedOk returns a tuple with the NumEpicsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEpicsCompleted

`func (o *LabelStats) SetNumEpicsCompleted(v int64)`

SetNumEpicsCompleted sets NumEpicsCompleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


