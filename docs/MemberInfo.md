# MemberInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**MentionName** | **string** |  | 
**Workspace2** | [**BasicWorkspaceInfo**](BasicWorkspaceInfo.md) |  | 

## Methods

### NewMemberInfo

`func NewMemberInfo(id string, name string, mentionName string, workspace2 BasicWorkspaceInfo, ) *MemberInfo`

NewMemberInfo instantiates a new MemberInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberInfoWithDefaults

`func NewMemberInfoWithDefaults() *MemberInfo`

NewMemberInfoWithDefaults instantiates a new MemberInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MemberInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberInfo) SetName(v string)`

SetName sets Name field to given value.


### GetMentionName

`func (o *MemberInfo) GetMentionName() string`

GetMentionName returns the MentionName field if non-nil, zero value otherwise.

### GetMentionNameOk

`func (o *MemberInfo) GetMentionNameOk() (*string, bool)`

GetMentionNameOk returns a tuple with the MentionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionName

`func (o *MemberInfo) SetMentionName(v string)`

SetMentionName sets MentionName field to given value.


### GetWorkspace2

`func (o *MemberInfo) GetWorkspace2() BasicWorkspaceInfo`

GetWorkspace2 returns the Workspace2 field if non-nil, zero value otherwise.

### GetWorkspace2Ok

`func (o *MemberInfo) GetWorkspace2Ok() (*BasicWorkspaceInfo, bool)`

GetWorkspace2Ok returns a tuple with the Workspace2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace2

`func (o *MemberInfo) SetWorkspace2(v BasicWorkspaceInfo)`

SetWorkspace2 sets Workspace2 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


