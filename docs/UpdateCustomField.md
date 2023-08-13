# UpdateCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether the Field is enabled for the Workspace. Only enabled fields can be applied to Stories. | [optional] 
**Name** | Pointer to **string** | A collection of objects representing reporting periods for years. | [optional] 
**Values** | Pointer to [**[]UpdateCustomFieldEnumValue**](UpdateCustomFieldEnumValue.md) | A collection of EnumValue objects representing the values in the domain of some Custom Field. | [optional] 
**IconSetIdentifier** | Pointer to **string** | A frontend-controlled string that represents the icon for this custom field. | [optional] 
**Description** | Pointer to **string** | A description of the purpose of this field. | [optional] 
**BeforeId** | Pointer to **string** | The ID of the CustomField we want to move this CustomField before. | [optional] 
**AfterId** | Pointer to **string** | The ID of the CustomField we want to move this CustomField after. | [optional] 

## Methods

### NewUpdateCustomField

`func NewUpdateCustomField() *UpdateCustomField`

NewUpdateCustomField instantiates a new UpdateCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomFieldWithDefaults

`func NewUpdateCustomFieldWithDefaults() *UpdateCustomField`

NewUpdateCustomFieldWithDefaults instantiates a new UpdateCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateCustomField) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCustomField) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCustomField) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCustomField) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *UpdateCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustomField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValues

`func (o *UpdateCustomField) GetValues() []UpdateCustomFieldEnumValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *UpdateCustomField) GetValuesOk() (*[]UpdateCustomFieldEnumValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *UpdateCustomField) SetValues(v []UpdateCustomFieldEnumValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *UpdateCustomField) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetIconSetIdentifier

`func (o *UpdateCustomField) GetIconSetIdentifier() string`

GetIconSetIdentifier returns the IconSetIdentifier field if non-nil, zero value otherwise.

### GetIconSetIdentifierOk

`func (o *UpdateCustomField) GetIconSetIdentifierOk() (*string, bool)`

GetIconSetIdentifierOk returns a tuple with the IconSetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconSetIdentifier

`func (o *UpdateCustomField) SetIconSetIdentifier(v string)`

SetIconSetIdentifier sets IconSetIdentifier field to given value.

### HasIconSetIdentifier

`func (o *UpdateCustomField) HasIconSetIdentifier() bool`

HasIconSetIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCustomField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCustomField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBeforeId

`func (o *UpdateCustomField) GetBeforeId() string`

GetBeforeId returns the BeforeId field if non-nil, zero value otherwise.

### GetBeforeIdOk

`func (o *UpdateCustomField) GetBeforeIdOk() (*string, bool)`

GetBeforeIdOk returns a tuple with the BeforeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeId

`func (o *UpdateCustomField) SetBeforeId(v string)`

SetBeforeId sets BeforeId field to given value.

### HasBeforeId

`func (o *UpdateCustomField) HasBeforeId() bool`

HasBeforeId returns a boolean if a field has been set.

### GetAfterId

`func (o *UpdateCustomField) GetAfterId() string`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *UpdateCustomField) GetAfterIdOk() (*string, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *UpdateCustomField) SetAfterId(v string)`

SetAfterId sets AfterId field to given value.

### HasAfterId

`func (o *UpdateCustomField) HasAfterId() bool`

HasAfterId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


