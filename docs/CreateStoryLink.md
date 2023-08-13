# CreateStoryLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verb** | **string** | The type of link. | 
**SubjectId** | **int64** | The ID of the subject Story. | 
**ObjectId** | **int64** | The ID of the object Story. | 

## Methods

### NewCreateStoryLink

`func NewCreateStoryLink(verb string, subjectId int64, objectId int64, ) *CreateStoryLink`

NewCreateStoryLink instantiates a new CreateStoryLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoryLinkWithDefaults

`func NewCreateStoryLinkWithDefaults() *CreateStoryLink`

NewCreateStoryLinkWithDefaults instantiates a new CreateStoryLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerb

`func (o *CreateStoryLink) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *CreateStoryLink) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *CreateStoryLink) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetSubjectId

`func (o *CreateStoryLink) GetSubjectId() int64`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *CreateStoryLink) GetSubjectIdOk() (*int64, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *CreateStoryLink) SetSubjectId(v int64)`

SetSubjectId sets SubjectId field to given value.


### GetObjectId

`func (o *CreateStoryLink) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CreateStoryLink) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CreateStoryLink) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


