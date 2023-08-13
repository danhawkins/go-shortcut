# CreateStoryLinkParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | Pointer to **int64** | The unique ID of the Story defined as subject. | [optional] 
**Verb** | **string** | How the subject Story acts on the object Story. This can be \&quot;blocks\&quot;, \&quot;duplicates\&quot;, or \&quot;relates to\&quot;. | 
**ObjectId** | Pointer to **int64** | The unique ID of the Story defined as object. | [optional] 

## Methods

### NewCreateStoryLinkParams

`func NewCreateStoryLinkParams(verb string, ) *CreateStoryLinkParams`

NewCreateStoryLinkParams instantiates a new CreateStoryLinkParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoryLinkParamsWithDefaults

`func NewCreateStoryLinkParamsWithDefaults() *CreateStoryLinkParams`

NewCreateStoryLinkParamsWithDefaults instantiates a new CreateStoryLinkParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *CreateStoryLinkParams) GetSubjectId() int64`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *CreateStoryLinkParams) GetSubjectIdOk() (*int64, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *CreateStoryLinkParams) SetSubjectId(v int64)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *CreateStoryLinkParams) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetVerb

`func (o *CreateStoryLinkParams) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *CreateStoryLinkParams) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *CreateStoryLinkParams) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetObjectId

`func (o *CreateStoryLinkParams) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CreateStoryLinkParams) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CreateStoryLinkParams) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *CreateStoryLinkParams) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


