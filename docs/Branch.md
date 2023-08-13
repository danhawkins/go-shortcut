# Branch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | 
**Deleted** | **bool** | A true/false boolean indicating if the Branch has been deleted. | 
**Name** | **string** | The name of the Branch. | 
**Persistent** | **bool** | A true/false boolean indicating if the Branch is persistent; e.g. master. | 
**UpdatedAt** | **NullableTime** | The time/date the Branch was updated. | 
**PullRequests** | [**[]PullRequest**](PullRequest.md) | An array of PullRequests attached to the Branch (there is usually only one). | 
**MergedBranchIds** | **[]int64** | The IDs of the Branches the Branch has been merged into. | 
**Id** | **NullableInt64** | The unique ID of the Branch. | 
**Url** | **string** | The URL of the Branch. | 
**RepositoryId** | **int64** | The ID of the Repository that contains the Branch. | 
**CreatedAt** | **NullableTime** | The time/date the Branch was created. | 

## Methods

### NewBranch

`func NewBranch(entityType string, deleted bool, name string, persistent bool, updatedAt NullableTime, pullRequests []PullRequest, mergedBranchIds []int64, id NullableInt64, url string, repositoryId int64, createdAt NullableTime, ) *Branch`

NewBranch instantiates a new Branch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchWithDefaults

`func NewBranchWithDefaults() *Branch`

NewBranchWithDefaults instantiates a new Branch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *Branch) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Branch) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Branch) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetDeleted

`func (o *Branch) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Branch) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Branch) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetName

`func (o *Branch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Branch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Branch) SetName(v string)`

SetName sets Name field to given value.


### GetPersistent

`func (o *Branch) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *Branch) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *Branch) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.


### GetUpdatedAt

`func (o *Branch) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Branch) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Branch) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Branch) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Branch) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetPullRequests

`func (o *Branch) GetPullRequests() []PullRequest`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *Branch) GetPullRequestsOk() (*[]PullRequest, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *Branch) SetPullRequests(v []PullRequest)`

SetPullRequests sets PullRequests field to given value.


### GetMergedBranchIds

`func (o *Branch) GetMergedBranchIds() []int64`

GetMergedBranchIds returns the MergedBranchIds field if non-nil, zero value otherwise.

### GetMergedBranchIdsOk

`func (o *Branch) GetMergedBranchIdsOk() (*[]int64, bool)`

GetMergedBranchIdsOk returns a tuple with the MergedBranchIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedBranchIds

`func (o *Branch) SetMergedBranchIds(v []int64)`

SetMergedBranchIds sets MergedBranchIds field to given value.


### GetId

`func (o *Branch) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Branch) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Branch) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Branch) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Branch) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUrl

`func (o *Branch) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Branch) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Branch) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRepositoryId

`func (o *Branch) GetRepositoryId() int64`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *Branch) GetRepositoryIdOk() (*int64, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *Branch) SetRepositoryId(v int64)`

SetRepositoryId sets RepositoryId field to given value.


### GetCreatedAt

`func (o *Branch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Branch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Branch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Branch) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Branch) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


