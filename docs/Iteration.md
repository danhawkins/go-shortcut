# Iteration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Iteration. | 
**Description** | **string** | The description of the iteration. | 
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

### NewIteration

`func NewIteration(appUrl string, description string, entityType string, labels []Label, mentionIds []string, memberMentionIds []string, associatedGroups []IterationAssociatedGroup, name string, globalId string, labelIds []int64, updatedAt time.Time, groupMentionIds []string, endDate time.Time, followerIds []string, groupIds []string, startDate time.Time, status string, id int64, stats IterationStats, createdAt time.Time, ) *Iteration`

NewIteration instantiates a new Iteration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationWithDefaults

`func NewIterationWithDefaults() *Iteration`

NewIterationWithDefaults instantiates a new Iteration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *Iteration) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *Iteration) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *Iteration) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetDescription

`func (o *Iteration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Iteration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Iteration) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntityType

`func (o *Iteration) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Iteration) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Iteration) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetLabels

`func (o *Iteration) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Iteration) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Iteration) SetLabels(v []Label)`

SetLabels sets Labels field to given value.


### GetMentionIds

`func (o *Iteration) GetMentionIds() []string`

GetMentionIds returns the MentionIds field if non-nil, zero value otherwise.

### GetMentionIdsOk

`func (o *Iteration) GetMentionIdsOk() (*[]string, bool)`

GetMentionIdsOk returns a tuple with the MentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionIds

`func (o *Iteration) SetMentionIds(v []string)`

SetMentionIds sets MentionIds field to given value.


### GetMemberMentionIds

`func (o *Iteration) GetMemberMentionIds() []string`

GetMemberMentionIds returns the MemberMentionIds field if non-nil, zero value otherwise.

### GetMemberMentionIdsOk

`func (o *Iteration) GetMemberMentionIdsOk() (*[]string, bool)`

GetMemberMentionIdsOk returns a tuple with the MemberMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberMentionIds

`func (o *Iteration) SetMemberMentionIds(v []string)`

SetMemberMentionIds sets MemberMentionIds field to given value.


### GetAssociatedGroups

`func (o *Iteration) GetAssociatedGroups() []IterationAssociatedGroup`

GetAssociatedGroups returns the AssociatedGroups field if non-nil, zero value otherwise.

### GetAssociatedGroupsOk

`func (o *Iteration) GetAssociatedGroupsOk() (*[]IterationAssociatedGroup, bool)`

GetAssociatedGroupsOk returns a tuple with the AssociatedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedGroups

`func (o *Iteration) SetAssociatedGroups(v []IterationAssociatedGroup)`

SetAssociatedGroups sets AssociatedGroups field to given value.


### GetName

`func (o *Iteration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Iteration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Iteration) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *Iteration) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *Iteration) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *Iteration) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetLabelIds

`func (o *Iteration) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *Iteration) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *Iteration) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.


### GetUpdatedAt

`func (o *Iteration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Iteration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Iteration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetGroupMentionIds

`func (o *Iteration) GetGroupMentionIds() []string`

GetGroupMentionIds returns the GroupMentionIds field if non-nil, zero value otherwise.

### GetGroupMentionIdsOk

`func (o *Iteration) GetGroupMentionIdsOk() (*[]string, bool)`

GetGroupMentionIdsOk returns a tuple with the GroupMentionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMentionIds

`func (o *Iteration) SetGroupMentionIds(v []string)`

SetGroupMentionIds sets GroupMentionIds field to given value.


### GetEndDate

`func (o *Iteration) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Iteration) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Iteration) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetFollowerIds

`func (o *Iteration) GetFollowerIds() []string`

GetFollowerIds returns the FollowerIds field if non-nil, zero value otherwise.

### GetFollowerIdsOk

`func (o *Iteration) GetFollowerIdsOk() (*[]string, bool)`

GetFollowerIdsOk returns a tuple with the FollowerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerIds

`func (o *Iteration) SetFollowerIds(v []string)`

SetFollowerIds sets FollowerIds field to given value.


### GetGroupIds

`func (o *Iteration) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *Iteration) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *Iteration) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.


### GetStartDate

`func (o *Iteration) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Iteration) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Iteration) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetStatus

`func (o *Iteration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Iteration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Iteration) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetId

`func (o *Iteration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Iteration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Iteration) SetId(v int64)`

SetId sets Id field to given value.


### GetStats

`func (o *Iteration) GetStats() IterationStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Iteration) GetStatsOk() (*IterationStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Iteration) SetStats(v IterationStats)`

SetStats sets Stats field to given value.


### GetCreatedAt

`func (o *Iteration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Iteration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Iteration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


