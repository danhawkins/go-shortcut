# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Group. | 
**Description** | **string** | The description of the Group. | 
**Archived** | **bool** | Whether or not the Group is archived. | 
**EntityType** | **string** | A string description of this resource. | 
**Color** | **NullableString** | The hex color to be displayed with the Group (for example, \&quot;#ff0000\&quot;). | 
**NumStoriesStarted** | **int64** | The number of stories assigned to the group which are in a started workflow state. | 
**MentionName** | **string** | The mention name of the Group. | 
**Name** | **string** | The name of the Group. | 
**GlobalId** | **string** |  | 
**ColorKey** | **NullableString** | The color key to be displayed with the Group. | 
**NumStories** | **int64** | The total number of stories assigned ot the group. | 
**NumEpicsStarted** | **int64** | The number of epics assigned to the group which are in the started workflow state. | 
**NumStoriesBacklog** | **int64** | The number of stories assigned to the group which are in a backlog workflow state. | 
**Id** | **string** | The id of the Group. | 
**DisplayIcon** | [**Icon**](Icon.md) |  | 
**MemberIds** | **[]string** | The Member IDs contain within the Group. | 
**WorkflowIds** | **[]int64** | The Workflow IDs contained within the Group. | 

## Methods

### NewGroup

`func NewGroup(appUrl string, description string, archived bool, entityType string, color NullableString, numStoriesStarted int64, mentionName string, name string, globalId string, colorKey NullableString, numStories int64, numEpicsStarted int64, numStoriesBacklog int64, id string, displayIcon Icon, memberIds []string, workflowIds []int64, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUrl

`func (o *Group) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *Group) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *Group) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArchived

`func (o *Group) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Group) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Group) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetEntityType

`func (o *Group) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Group) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Group) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetColor

`func (o *Group) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Group) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Group) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *Group) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *Group) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetNumStoriesStarted

`func (o *Group) GetNumStoriesStarted() int64`

GetNumStoriesStarted returns the NumStoriesStarted field if non-nil, zero value otherwise.

### GetNumStoriesStartedOk

`func (o *Group) GetNumStoriesStartedOk() (*int64, bool)`

GetNumStoriesStartedOk returns a tuple with the NumStoriesStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesStarted

`func (o *Group) SetNumStoriesStarted(v int64)`

SetNumStoriesStarted sets NumStoriesStarted field to given value.


### GetMentionName

`func (o *Group) GetMentionName() string`

GetMentionName returns the MentionName field if non-nil, zero value otherwise.

### GetMentionNameOk

`func (o *Group) GetMentionNameOk() (*string, bool)`

GetMentionNameOk returns a tuple with the MentionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionName

`func (o *Group) SetMentionName(v string)`

SetMentionName sets MentionName field to given value.


### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetGlobalId

`func (o *Group) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *Group) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *Group) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.


### GetColorKey

`func (o *Group) GetColorKey() string`

GetColorKey returns the ColorKey field if non-nil, zero value otherwise.

### GetColorKeyOk

`func (o *Group) GetColorKeyOk() (*string, bool)`

GetColorKeyOk returns a tuple with the ColorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorKey

`func (o *Group) SetColorKey(v string)`

SetColorKey sets ColorKey field to given value.


### SetColorKeyNil

`func (o *Group) SetColorKeyNil(b bool)`

 SetColorKeyNil sets the value for ColorKey to be an explicit nil

### UnsetColorKey
`func (o *Group) UnsetColorKey()`

UnsetColorKey ensures that no value is present for ColorKey, not even an explicit nil
### GetNumStories

`func (o *Group) GetNumStories() int64`

GetNumStories returns the NumStories field if non-nil, zero value otherwise.

### GetNumStoriesOk

`func (o *Group) GetNumStoriesOk() (*int64, bool)`

GetNumStoriesOk returns a tuple with the NumStories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStories

`func (o *Group) SetNumStories(v int64)`

SetNumStories sets NumStories field to given value.


### GetNumEpicsStarted

`func (o *Group) GetNumEpicsStarted() int64`

GetNumEpicsStarted returns the NumEpicsStarted field if non-nil, zero value otherwise.

### GetNumEpicsStartedOk

`func (o *Group) GetNumEpicsStartedOk() (*int64, bool)`

GetNumEpicsStartedOk returns a tuple with the NumEpicsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEpicsStarted

`func (o *Group) SetNumEpicsStarted(v int64)`

SetNumEpicsStarted sets NumEpicsStarted field to given value.


### GetNumStoriesBacklog

`func (o *Group) GetNumStoriesBacklog() int64`

GetNumStoriesBacklog returns the NumStoriesBacklog field if non-nil, zero value otherwise.

### GetNumStoriesBacklogOk

`func (o *Group) GetNumStoriesBacklogOk() (*int64, bool)`

GetNumStoriesBacklogOk returns a tuple with the NumStoriesBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStoriesBacklog

`func (o *Group) SetNumStoriesBacklog(v int64)`

SetNumStoriesBacklog sets NumStoriesBacklog field to given value.


### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayIcon

`func (o *Group) GetDisplayIcon() Icon`

GetDisplayIcon returns the DisplayIcon field if non-nil, zero value otherwise.

### GetDisplayIconOk

`func (o *Group) GetDisplayIconOk() (*Icon, bool)`

GetDisplayIconOk returns a tuple with the DisplayIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIcon

`func (o *Group) SetDisplayIcon(v Icon)`

SetDisplayIcon sets DisplayIcon field to given value.


### GetMemberIds

`func (o *Group) GetMemberIds() []string`

GetMemberIds returns the MemberIds field if non-nil, zero value otherwise.

### GetMemberIdsOk

`func (o *Group) GetMemberIdsOk() (*[]string, bool)`

GetMemberIdsOk returns a tuple with the MemberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIds

`func (o *Group) SetMemberIds(v []string)`

SetMemberIds sets MemberIds field to given value.


### GetWorkflowIds

`func (o *Group) GetWorkflowIds() []int64`

GetWorkflowIds returns the WorkflowIds field if non-nil, zero value otherwise.

### GetWorkflowIdsOk

`func (o *Group) GetWorkflowIdsOk() (*[]int64, bool)`

GetWorkflowIdsOk returns a tuple with the WorkflowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowIds

`func (o *Group) SetWorkflowIds(v []int64)`

SetWorkflowIds sets WorkflowIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


