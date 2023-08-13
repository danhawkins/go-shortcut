# SearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epics** | Pointer to [**EpicSearchResults**](EpicSearchResults.md) |  | [optional] 
**Stories** | Pointer to [**StorySearchResults**](StorySearchResults.md) |  | [optional] 
**Iterations** | Pointer to [**IterationSearchResults**](IterationSearchResults.md) |  | [optional] 
**Milestones** | Pointer to [**MilestoneSearchResults**](MilestoneSearchResults.md) |  | [optional] 

## Methods

### NewSearchResults

`func NewSearchResults() *SearchResults`

NewSearchResults instantiates a new SearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultsWithDefaults

`func NewSearchResultsWithDefaults() *SearchResults`

NewSearchResultsWithDefaults instantiates a new SearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpics

`func (o *SearchResults) GetEpics() EpicSearchResults`

GetEpics returns the Epics field if non-nil, zero value otherwise.

### GetEpicsOk

`func (o *SearchResults) GetEpicsOk() (*EpicSearchResults, bool)`

GetEpicsOk returns a tuple with the Epics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpics

`func (o *SearchResults) SetEpics(v EpicSearchResults)`

SetEpics sets Epics field to given value.

### HasEpics

`func (o *SearchResults) HasEpics() bool`

HasEpics returns a boolean if a field has been set.

### GetStories

`func (o *SearchResults) GetStories() StorySearchResults`

GetStories returns the Stories field if non-nil, zero value otherwise.

### GetStoriesOk

`func (o *SearchResults) GetStoriesOk() (*StorySearchResults, bool)`

GetStoriesOk returns a tuple with the Stories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStories

`func (o *SearchResults) SetStories(v StorySearchResults)`

SetStories sets Stories field to given value.

### HasStories

`func (o *SearchResults) HasStories() bool`

HasStories returns a boolean if a field has been set.

### GetIterations

`func (o *SearchResults) GetIterations() IterationSearchResults`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *SearchResults) GetIterationsOk() (*IterationSearchResults, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *SearchResults) SetIterations(v IterationSearchResults)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *SearchResults) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### GetMilestones

`func (o *SearchResults) GetMilestones() MilestoneSearchResults`

GetMilestones returns the Milestones field if non-nil, zero value otherwise.

### GetMilestonesOk

`func (o *SearchResults) GetMilestonesOk() (*MilestoneSearchResults, bool)`

GetMilestonesOk returns a tuple with the Milestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestones

`func (o *SearchResults) SetMilestones(v MilestoneSearchResults)`

SetMilestones sets Milestones field to given value.

### HasMilestones

`func (o *SearchResults) HasMilestones() bool`

HasMilestones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


