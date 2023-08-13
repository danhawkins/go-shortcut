# IterationAssociatedGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The Group ID of the associated group. | 
**AssociatedStoriesCount** | Pointer to **int64** | The number of stories this Group owns in the Iteration. | [optional] 

## Methods

### NewIterationAssociatedGroup

`func NewIterationAssociatedGroup(groupId string, ) *IterationAssociatedGroup`

NewIterationAssociatedGroup instantiates a new IterationAssociatedGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationAssociatedGroupWithDefaults

`func NewIterationAssociatedGroupWithDefaults() *IterationAssociatedGroup`

NewIterationAssociatedGroupWithDefaults instantiates a new IterationAssociatedGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *IterationAssociatedGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *IterationAssociatedGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *IterationAssociatedGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetAssociatedStoriesCount

`func (o *IterationAssociatedGroup) GetAssociatedStoriesCount() int64`

GetAssociatedStoriesCount returns the AssociatedStoriesCount field if non-nil, zero value otherwise.

### GetAssociatedStoriesCountOk

`func (o *IterationAssociatedGroup) GetAssociatedStoriesCountOk() (*int64, bool)`

GetAssociatedStoriesCountOk returns a tuple with the AssociatedStoriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedStoriesCount

`func (o *IterationAssociatedGroup) SetAssociatedStoriesCount(v int64)`

SetAssociatedStoriesCount sets AssociatedStoriesCount field to given value.

### HasAssociatedStoriesCount

`func (o *IterationAssociatedGroup) HasAssociatedStoriesCount() bool`

HasAssociatedStoriesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


