# IterationSlim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Iteration. | 
**EntityType** | **string** | A string description of this resource | 
**Labels** | [**[]Label**](Label.md) | An array of labels attached to the iteration. | 
**MentionIds** | **[]string** | Deprecated: use member_mention_ids. | 
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Story description. | 
**AssociatedGroups** | [**[]IterationAssociatedGroup**](IterationAssociatedGroup.md) | An array containing Group IDs and Group-owned story counts for the Iteration&#39;s associated groups. | 
**Name** | **string** | The name of the iteration. | 
**GlobalId** | **string** |  | 
**LabelIds** | **[]int64** | An array of label ids attached to the iteration. | 
**UpdatedAt** | **time.Time** | The instant when this iteration was last updated. | 
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Story description. | 
**EndDate** | **time.Time** | The date this iteration begins. | 
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | 
**GroupIds** | **[]string** | An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI. | 
**StartDate** | **time.Time** | The date this iteration begins. | 
**Status** | **string** | The status of the iteration. Values are either \&quot;unstarted\&quot;, \&quot;started\&quot;, or \&quot;done\&quot;. | 
**Id** | **int64** | The ID of the iteration. | 
**Stats** | [**IterationStats**](IterationStats.md) |  | 
**CreatedAt** | **time.Time** | The instant when this iteration was created. | 

## Methods

### NewIterationSlim

`func NewIterationSlim(appUrl string, entityType string, labels []Label, mentionIds []string, memberMentionIds []string, associatedGroups []IterationAssociatedGroup, name string, globalId string, labelIds []int64, updatedAt time.Time, groupMentionIds []string, endDate time.Time, followerIds []string, groupIds []string, startDate time.Time, status string, id int64, stats IterationStats, createdAt time.Time, ) *IterationSlim`

NewIterationSlim instantiates a new IterationSlim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationSlimWithDefaults

`func NewIterationSlimWithDefaults() *IterationSlim`

NewIterationSlimWithDefaults instantiates a new IterationSlim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *IterationSlim) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *IterationSlim) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *IterationSlim) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetEntityType

`func (o *IterationSlim) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *IterationSlim) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *IterationSlim) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetLabels

`func (o *IterationSlim) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IterationSlim) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IterationSlim) SetLabels(v []Label)`

SetLabels sets Labels field to given value.


### GetMentionIds

`func (o *IterationSlim) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *IterationSlim) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *IterationSlim) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetMemberMentionIds

`func (o *IterationSlim) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *IterationSlim) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *IterationSlim) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetAssociatedGroups

`func (o *IterationSlim) GetAssociatedGroups() []IterationAssociatedGroup`

GetAssociatedGroups returns the AssociatedGroups field if non-nil, zero value otherwise.

### GetAssociatedGroupsOk

`func (o *IterationSlim) GetAssociatedGroupsOk() (*[]IterationAssociatedGroup, bool)`

GetAssociatedGroupsOk returns a tuple with the AssociatedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedGroups

`func (o *IterationSlim) SetAssociatedGroups(v []IterationAssociatedGroup)`

SetAssociatedGroups sets AssociatedGroups field to given value.


### GetName

`func (o *IterationSlim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IterationSlim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IterationSlim) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *IterationSlim) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *IterationSlim) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *IterationSlim) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetLabelIds

`func (o *IterationSlim) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *IterationSlim) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *IterationSlim) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.


### GetUpdatedAt

`func (o *IterationSlim) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IterationSlim) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IterationSlim) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetGroupMentionIds

`func (o *IterationSlim) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *IterationSlim) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *IterationSlim) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetEndDate

`func (o *IterationSlim) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *IterationSlim) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *IterationSlim) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetFollowerIds

`func (o *IterationSlim) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *IterationSlim) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *IterationSlim) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.


### GetGroupIds

`func (o *IterationSlim) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *IterationSlim) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *IterationSlim) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.


### GetStartDate

`func (o *IterationSlim) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *IterationSlim) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *IterationSlim) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetStatus

`func (o *IterationSlim) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IterationSlim) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IterationSlim) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetId

`func (o *IterationSlim) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IterationSlim) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IterationSlim) SetId(v int64)`

SetId sets Id field to given value.


### GetStats

`func (o *IterationSlim) GetStats() IterationStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *IterationSlim) GetStatsOk() (*IterationStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *IterationSlim) SetStats(v IterationStats)`

SetStats sets Stats field to given value.


### GetCreatedAt

`func (o *IterationSlim) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IterationSlim) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IterationSlim) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


