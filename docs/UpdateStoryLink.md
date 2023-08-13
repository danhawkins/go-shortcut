# UpdateStoryLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verb** | Pointer to **string** | The type of link. | [optional] 
**SubjectId** | Pointer to **int64** | The ID of the subject Story. | [optional] 
**ObjectId** | Pointer to **int64** | The ID of the object Story. | [optional] 

## Methods

### NewUpdateStoryLink

`func NewUpdateStoryLink() *UpdateStoryLink`

NewUpdateStoryLink instantiates a new UpdateStoryLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStoryLinkWithDefaults

`func NewUpdateStoryLinkWithDefaults() *UpdateStoryLink`

NewUpdateStoryLinkWithDefaults instantiates a new UpdateStoryLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerb

`func (o *UpdateStoryLink) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *UpdateStoryLink) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *UpdateStoryLink) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *UpdateStoryLink) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### GetSubjectId

`func (o *UpdateStoryLink) GetSubjectId() int64`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *UpdateStoryLink) GetSubjectIdOk() (*int64, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *UpdateStoryLink) SetSubjectId(v int64)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *UpdateStoryLink) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetObjectId

`func (o *UpdateStoryLink) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *UpdateStoryLink) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *UpdateStoryLink) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *UpdateStoryLink) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


