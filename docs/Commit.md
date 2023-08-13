# Commit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | 
**AuthorId** | **NullableString** | The ID of the Member that authored the Commit, if known. | 
**Hash** | **string** | The Commit hash. | 
**UpdatedAt** | **NullableTime** | The time/date the Commit was updated. | 
**Id** | **NullableInt64** | The unique ID of the Commit. | 
**Url** | **string** | The URL of the Commit. | 
**AuthorEmail** | **string** | The email address of the VCS user that authored the Commit. | 
**Timestamp** | **time.Time** | The time/date the Commit was pushed. | 
**AuthorIdentity** | [**Identity**](Identity.md) |  | 
**RepositoryId** | **NullableInt64** | The ID of the Repository that contains the Commit. | 
**CreatedAt** | **time.Time** | The time/date the Commit was created. | 
**Message** | **string** | The Commit message. | 

## Methods

### NewCommit

`func NewCommit(entityType string, authorId NullableString, hash string, updatedAt NullableTime, id NullableInt64, url string, authorEmail string, timestamp time.Time, authorIdentity Identity, repositoryId NullableInt64, createdAt time.Time, message string, ) *Commit`

NewCommit instantiates a new Commit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitWithDefaults

`func NewCommitWithDefaults() *Commit`

NewCommitWithDefaults instantiates a new Commit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *Commit) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Commit) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Commit) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetAuthorId

`func (o *Commit) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Commit) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Commit) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### SetAuthorIdNil

`func (o *Commit) SetAuthorIdNil(b bool)`

 SetAuthorIdNil sets the value for AuthorId to be an explicit nil

### UnsetAuthorId
`func (o *Commit) UnsetAuthorId()`

UnsetAuthorId ensures that no value is present for AuthorId, not even an explicit nil
### GetHash

`func (o *Commit) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Commit) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Commit) SetHash(v string)`

SetHash sets Hash field to given value.


### GetUpdatedAt

`func (o *Commit) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Commit) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Commit) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Commit) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Commit) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetId

`func (o *Commit) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Commit) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Commit) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Commit) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Commit) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUrl

`func (o *Commit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Commit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Commit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthorEmail

`func (o *Commit) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *Commit) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *Commit) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.


### GetTimestamp

`func (o *Commit) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Commit) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Commit) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetAuthorIdentity

`func (o *Commit) GetAuthorIdentity() Identity`

GetAuthorIdentity returns the AuthorIdentity field if non-nil, zero value otherwise.

### GetAuthorIdentityOk

`func (o *Commit) GetAuthorIdentityOk() (*Identity, bool)`

GetAuthorIdentityOk returns a tuple with the AuthorIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIdentity

`func (o *Commit) SetAuthorIdentity(v Identity)`

SetAuthorIdentity sets AuthorIdentity field to given value.


### GetRepositoryId

`func (o *Commit) GetRepositoryId() int64`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *Commit) GetRepositoryIdOk() (*int64, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *Commit) SetRepositoryId(v int64)`

SetRepositoryId sets RepositoryId field to given value.


### SetRepositoryIdNil

`func (o *Commit) SetRepositoryIdNil(b bool)`

 SetRepositoryIdNil sets the value for RepositoryId to be an explicit nil

### UnsetRepositoryId
`func (o *Commit) UnsetRepositoryId()`

UnsetRepositoryId ensures that no value is present for RepositoryId, not even an explicit nil
### GetCreatedAt

`func (o *Commit) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Commit) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Commit) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMessage

`func (o *Commit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Commit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Commit) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


