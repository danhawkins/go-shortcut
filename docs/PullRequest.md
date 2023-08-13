# PullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | 
**Closed** | **bool** | True/False boolean indicating whether the VCS pull request has been closed. | 
**Merged** | **bool** | True/False boolean indicating whether the VCS pull request has been merged. | 
**NumAdded** | **int64** | Number of lines added in the pull request, according to VCS. | 
**BranchId** | **int64** | The ID of the branch for the particular pull request. | 
**OverlappingStories** | Pointer to **[]int64** | An array of Story ids that have Pull Requests that change at least one of the same lines this Pull Request changes. | [optional] 
**Number** | **int64** | The pull request&#39;s unique number ID in VCS. | 
**BranchName** | **string** | The name of the branch for the particular pull request. | 
**TargetBranchName** | **string** | The name of the target branch for the particular pull request. | 
**NumCommits** | **NullableInt64** | The number of commits on the pull request. | 
**Title** | **string** | The title of the pull request. | 
**UpdatedAt** | **time.Time** | The time/date the pull request was created. | 
**HasOverlappingStories** | **bool** | Boolean indicating that the Pull Request has Stories that have Pull Requests that change at least one of the same lines this Pull Request changes. | 
**Draft** | **bool** | True/False boolean indicating whether the VCS pull request is in the draft state. | 
**Id** | **int64** | The unique ID associated with the pull request in Shortcut. | 
**VcsLabels** | Pointer to [**[]PullRequestLabel**](PullRequestLabel.md) | An array of PullRequestLabels attached to the PullRequest. | [optional] 
**Url** | **string** | The URL for the pull request. | 
**NumRemoved** | **int64** | Number of lines removed in the pull request, according to VCS. | 
**ReviewStatus** | Pointer to **string** | The status of the review for the pull request. | [optional] 
**NumModified** | **NullableInt64** | Number of lines modified in the pull request, according to VCS. | 
**BuildStatus** | Pointer to **string** | The status of the Continuous Integration workflow for the pull request. | [optional] 
**TargetBranchId** | **int64** | The ID of the target branch for the particular pull request. | 
**RepositoryId** | **int64** | The ID of the repository for the particular pull request. | 
**CreatedAt** | **time.Time** | The time/date the pull request was created. | 

## Methods

### NewPullRequest

`func NewPullRequest(entityType string, closed bool, merged bool, numAdded int64, branchId int64, number int64, branchName string, targetBranchName string, numCommits NullableInt64, title string, updatedAt time.Time, hasOverlappingStories bool, draft bool, id int64, url string, numRemoved int64, numModified NullableInt64, targetBranchId int64, repositoryId int64, createdAt time.Time, ) *PullRequest`

NewPullRequest instantiates a new PullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestWithDefaults

`func NewPullRequestWithDefaults() *PullRequest`

NewPullRequestWithDefaults instantiates a new PullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *PullRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *PullRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *PullRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetClosed

`func (o *PullRequest) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *PullRequest) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *PullRequest) SetClosed(v bool)`

SetClosed sets Closed field to given value.


### GetMerged

`func (o *PullRequest) GetMerged() bool`

GetMerged returns the Merged field if non-nil, zero value otherwise.

### GetMergedOk

`func (o *PullRequest) GetMergedOk() (*bool, bool)`

GetMergedOk returns a tuple with the Merged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerged

`func (o *PullRequest) SetMerged(v bool)`

SetMerged sets Merged field to given value.


### GetNumAdded

`func (o *PullRequest) GetNumAdded() int64`

GetNumAdded returns the NumAdded field if non-nil, zero value otherwise.

### GetNumAddedOk

`func (o *PullRequest) GetNumAddedOk() (*int64, bool)`

GetNumAddedOk returns a tuple with the NumAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdded

`func (o *PullRequest) SetNumAdded(v int64)`

SetNumAdded sets NumAdded field to given value.


### GetBranchId

`func (o *PullRequest) GetBranchId() int64`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *PullRequest) GetBranchIdOk() (*int64, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *PullRequest) SetBranchId(v int64)`

SetBranchId sets BranchId field to given value.


### GetOverlappingStories

`func (o *PullRequest) GetOverlappingStories() []int64`

GetOverlappingStories returns the OverlappingStories field if non-nil, zero value otherwise.

### GetOverlappingStoriesOk

`func (o *PullRequest) GetOverlappingStoriesOk() (*[]int64, bool)`

GetOverlappingStoriesOk returns a tuple with the OverlappingStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlappingStories

`func (o *PullRequest) SetOverlappingStories(v []int64)`

SetOverlappingStories sets OverlappingStories field to given value.

### HasOverlappingStories

`func (o *PullRequest) HasOverlappingStories() bool`

HasOverlappingStories returns a boolean if a field has been set.

### GetNumber

`func (o *PullRequest) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PullRequest) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PullRequest) SetNumber(v int64)`

SetNumber sets Number field to given value.


### GetBranchName

`func (o *PullRequest) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *PullRequest) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *PullRequest) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetTargetBranchName

`func (o *PullRequest) GetTargetBranchName() string`

GetTargetBranchName returns the TargetBranchName field if non-nil, zero value otherwise.

### GetTargetBranchNameOk

`func (o *PullRequest) GetTargetBranchNameOk() (*string, bool)`

GetTargetBranchNameOk returns a tuple with the TargetBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranchName

`func (o *PullRequest) SetTargetBranchName(v string)`

SetTargetBranchName sets TargetBranchName field to given value.


### GetNumCommits

`func (o *PullRequest) GetNumCommits() int64`

GetNumCommits returns the NumCommits field if non-nil, zero value otherwise.

### GetNumCommitsOk

`func (o *PullRequest) GetNumCommitsOk() (*int64, bool)`

GetNumCommitsOk returns a tuple with the NumCommits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCommits

`func (o *PullRequest) SetNumCommits(v int64)`

SetNumCommits sets NumCommits field to given value.


### SetNumCommitsNil

`func (o *PullRequest) SetNumCommitsNil(b bool)`

 SetNumCommitsNil sets the value for NumCommits to be an explicit nil

### UnsetNumCommits
`func (o *PullRequest) UnsetNumCommits()`

UnsetNumCommits ensures that no value is present for NumCommits, not even an explicit nil
### GetTitle

`func (o *PullRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *PullRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PullRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PullRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetHasOverlappingStories

`func (o *PullRequest) GetHasOverlappingStories() bool`

GetHasOverlappingStories returns the HasOverlappingStories field if non-nil, zero value otherwise.

### GetHasOverlappingStoriesOk

`func (o *PullRequest) GetHasOverlappingStoriesOk() (*bool, bool)`

GetHasOverlappingStoriesOk returns a tuple with the HasOverlappingStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOverlappingStories

`func (o *PullRequest) SetHasOverlappingStories(v bool)`

SetHasOverlappingStories sets HasOverlappingStories field to given value.


### GetDraft

`func (o *PullRequest) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *PullRequest) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *PullRequest) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetId

`func (o *PullRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetVcsLabels

`func (o *PullRequest) GetVcsLabels() []PullRequestLabel`

GetVcsLabels returns the VcsLabels field if non-nil, zero value otherwise.

### GetVcsLabelsOk

`func (o *PullRequest) GetVcsLabelsOk() (*[]PullRequestLabel, bool)`

GetVcsLabelsOk returns a tuple with the VcsLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsLabels

`func (o *PullRequest) SetVcsLabels(v []PullRequestLabel)`

SetVcsLabels sets VcsLabels field to given value.

### HasVcsLabels

`func (o *PullRequest) HasVcsLabels() bool`

HasVcsLabels returns a boolean if a field has been set.

### SetVcsLabelsNil

`func (o *PullRequest) SetVcsLabelsNil(b bool)`

 SetVcsLabelsNil sets the value for VcsLabels to be an explicit nil

### UnsetVcsLabels
`func (o *PullRequest) UnsetVcsLabels()`

UnsetVcsLabels ensures that no value is present for VcsLabels, not even an explicit nil
### GetUrl

`func (o *PullRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNumRemoved

`func (o *PullRequest) GetNumRemoved() int64`

GetNumRemoved returns the NumRemoved field if non-nil, zero value otherwise.

### GetNumRemovedOk

`func (o *PullRequest) GetNumRemovedOk() (*int64, bool)`

GetNumRemovedOk returns a tuple with the NumRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRemoved

`func (o *PullRequest) SetNumRemoved(v int64)`

SetNumRemoved sets NumRemoved field to given value.


### GetReviewStatus

`func (o *PullRequest) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *PullRequest) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *PullRequest) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.

### HasReviewStatus

`func (o *PullRequest) HasReviewStatus() bool`

HasReviewStatus returns a boolean if a field has been set.

### GetNumModified

`func (o *PullRequest) GetNumModified() int64`

GetNumModified returns the NumModified field if non-nil, zero value otherwise.

### GetNumModifiedOk

`func (o *PullRequest) GetNumModifiedOk() (*int64, bool)`

GetNumModifiedOk returns a tuple with the NumModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumModified

`func (o *PullRequest) SetNumModified(v int64)`

SetNumModified sets NumModified field to given value.


### SetNumModifiedNil

`func (o *PullRequest) SetNumModifiedNil(b bool)`

 SetNumModifiedNil sets the value for NumModified to be an explicit nil

### UnsetNumModified
`func (o *PullRequest) UnsetNumModified()`

UnsetNumModified ensures that no value is present for NumModified, not even an explicit nil
### GetBuildStatus

`func (o *PullRequest) GetBuildStatus() string`

GetBuildStatus returns the BuildStatus field if non-nil, zero value otherwise.

### GetBuildStatusOk

`func (o *PullRequest) GetBuildStatusOk() (*string, bool)`

GetBuildStatusOk returns a tuple with the BuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildStatus

`func (o *PullRequest) SetBuildStatus(v string)`

SetBuildStatus sets BuildStatus field to given value.

### HasBuildStatus

`func (o *PullRequest) HasBuildStatus() bool`

HasBuildStatus returns a boolean if a field has been set.

### GetTargetBranchId

`func (o *PullRequest) GetTargetBranchId() int64`

GetTargetBranchId returns the TargetBranchId field if non-nil, zero value otherwise.

### GetTargetBranchIdOk

`func (o *PullRequest) GetTargetBranchIdOk() (*int64, bool)`

GetTargetBranchIdOk returns a tuple with the TargetBranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranchId

`func (o *PullRequest) SetTargetBranchId(v int64)`

SetTargetBranchId sets TargetBranchId field to given value.


### GetRepositoryId

`func (o *PullRequest) GetRepositoryId() int64`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *PullRequest) GetRepositoryIdOk() (*int64, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *PullRequest) SetRepositoryId(v int64)`

SetRepositoryId sets RepositoryId field to given value.


### GetCreatedAt

`func (o *PullRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PullRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PullRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


