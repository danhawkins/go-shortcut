# MilestoneStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageCycleTime** | Pointer to **int64** | The average cycle time (in seconds) of completed stories in this Milestone. | [optional] 
**AverageLeadTime** | Pointer to **int64** | The average lead time (in seconds) of completed stories in this Milestone. | [optional] 
**NumRelatedDocuments** | **int64** | The number of related documents tp this Milestone. | 

## Methods

### NewMilestoneStats

`func NewMilestoneStats(numRelatedDocuments int64, ) *MilestoneStats`

NewMilestoneStats instantiates a new MilestoneStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilestoneStatsWithDefaults

`func NewMilestoneStatsWithDefaults() *MilestoneStats`

NewMilestoneStatsWithDefaults instantiates a new MilestoneStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageCycleTime

`func (o *MilestoneStats) GetAverageCycleTime() int64`

GetAverageCycleTime returns the AverageCycleTime field if non-nil, zero value otherwise.

### GetAverageCycleTimeOk

`func (o *MilestoneStats) GetAverageCycleTimeOk() (*int64, bool)`

GetAverageCycleTimeOk returns a tuple with the AverageCycleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCycleTime

`func (o *MilestoneStats) SetAverageCycleTime(v int64)`

SetAverageCycleTime sets AverageCycleTime field to given value.

### HasAverageCycleTime

`func (o *MilestoneStats) HasAverageCycleTime() bool`

HasAverageCycleTime returns a boolean if a field has been set.

### GetAverageLeadTime

`func (o *MilestoneStats) GetAverageLeadTime() int64`

GetAverageLeadTime returns the AverageLeadTime field if non-nil, zero value otherwise.

### GetAverageLeadTimeOk

`func (o *MilestoneStats) GetAverageLeadTimeOk() (*int64, bool)`

GetAverageLeadTimeOk returns a tuple with the AverageLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageLeadTime

`func (o *MilestoneStats) SetAverageLeadTime(v int64)`

SetAverageLeadTime sets AverageLeadTime field to given value.

### HasAverageLeadTime

`func (o *MilestoneStats) HasAverageLeadTime() bool`

HasAverageLeadTime returns a boolean if a field has been set.

### GetNumRelatedDocuments

`func (o *MilestoneStats) GetNumRelatedDocuments() int64`

GetNumRelatedDocuments returns the NumRelatedDocuments field if non-nil, zero value otherwise.

### GetNumRelatedDocumentsOk

`func (o *MilestoneStats) GetNumRelatedDocumentsOk() (*int64, bool)`

GetNumRelatedDocumentsOk returns a tuple with the NumRelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRelatedDocuments

`func (o *MilestoneStats) SetNumRelatedDocuments(v int64)`

SetNumRelatedDocuments sets NumRelatedDocuments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


