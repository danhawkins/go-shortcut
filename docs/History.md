# History

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangedAt** | **string** | The date when the change occurred. | 
**PrimaryId** | Pointer to **map[string]interface{}** | The ID of the primary entity that has changed, if applicable. | [optional] 
**References** | Pointer to **[]map[string]interface{}** | An array of objects affected by the change. Reference objects provide basic information for the entities reference in the history actions. Some have specific fields, but they always contain an id, entity_type, and a name. | [optional] 
**Actions** | **[]map[string]interface{}** | An array of actions that were performed for the change. | 
**MemberId** | Pointer to **string** | The ID of the member who performed the change. | [optional] 
**ExternalId** | Pointer to **string** | The ID of the webhook that handled the change. | [optional] 
**Id** | **string** | The ID representing the change for the story. | 
**Version** | **string** | The version of the change format. | 
**WebhookId** | Pointer to **NullableString** | The ID of the webhook that handled the change. | [optional] 

## Methods

### NewHistory

`func NewHistory(changedAt string, actions []map[string]interface{}, id string, version string, ) *History`

NewHistory instantiates a new History object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryWithDefaults

`func NewHistoryWithDefaults() *History`

NewHistoryWithDefaults instantiates a new History object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangedAt

`func (o *History) GetChangedAt() string`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *History) GetChangedAtOk() (*string, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *History) SetChangedAt(v string)`

SetChangedAt sets ChangedAt field to given value.


### GetPrimaryId

`func (o *History) GetPrimaryId() map[string]interface{}`

GetPrimaryId returns the PrimaryId field if non-nil, zero value otherwise.

### GetPrimaryIdOk

`func (o *History) GetPrimaryIdOk() (*map[string]interface{}, bool)`

GetPrimaryIdOk returns a tuple with the PrimaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryId

`func (o *History) SetPrimaryId(v map[string]interface{})`

SetPrimaryId sets PrimaryId field to given value.

### HasPrimaryId

`func (o *History) HasPrimaryId() bool`

HasPrimaryId returns a boolean if a field has been set.

### GetReferences

`func (o *History) GetReferences() []map[string]interface{}`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *History) GetReferencesOk() (*[]map[string]interface{}, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *History) SetReferences(v []map[string]interface{})`

SetReferences sets References field to given value.

### HasReferences

`func (o *History) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetActions

`func (o *History) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *History) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *History) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.


### GetMemberId

`func (o *History) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *History) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *History) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *History) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetExternalId

`func (o *History) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *History) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *History) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *History) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *History) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *History) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *History) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *History) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *History) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *History) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetWebhookId

`func (o *History) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *History) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *History) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *History) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### SetWebhookIdNil

`func (o *History) SetWebhookIdNil(b bool)`

 SetWebhookIdNil sets the value for WebhookId to be an explicit nil

### UnsetWebhookId
`func (o *History) UnsetWebhookId()`

UnsetWebhookId ensures that no value is present for WebhookId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


