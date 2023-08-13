# UpdateCustomFieldEnumValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of an existing EnumValue within the CustomField&#39;s domain. | [optional] 
**Value** | Pointer to **string** | A string value within the domain of this Custom Field. | [optional] 
**ColorKey** | Pointer to **NullableString** | A color key associated with this EnumValue within the CustomField&#39;s domain. | [optional] 
**Enabled** | Pointer to **bool** | Whether this EnumValue is enabled for its CustomField or not. Leaving this key out of the request leaves the current enabled state untouched. | [optional] 

## Methods

### NewUpdateCustomFieldEnumValue

`func NewUpdateCustomFieldEnumValue() *UpdateCustomFieldEnumValue`

NewUpdateCustomFieldEnumValue instantiates a new UpdateCustomFieldEnumValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomFieldEnumValueWithDefaults

`func NewUpdateCustomFieldEnumValueWithDefaults() *UpdateCustomFieldEnumValue`

NewUpdateCustomFieldEnumValueWithDefaults instantiates a new UpdateCustomFieldEnumValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCustomFieldEnumValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCustomFieldEnumValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCustomFieldEnumValue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCustomFieldEnumValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *UpdateCustomFieldEnumValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateCustomFieldEnumValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateCustomFieldEnumValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateCustomFieldEnumValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetColorKey

`func (o *UpdateCustomFieldEnumValue) GetColorKey() string`

GetColorKey returns the ColorKey field if non-nil, zero value otherwise.

### GetColorKeyOk

`func (o *UpdateCustomFieldEnumValue) GetColorKeyOk() (*string, bool)`

GetColorKeyOk returns a tuple with the ColorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorKey

`func (o *UpdateCustomFieldEnumValue) SetColorKey(v string)`

SetColorKey sets ColorKey field to given value.

### HasColorKey

`func (o *UpdateCustomFieldEnumValue) HasColorKey() bool`

HasColorKey returns a boolean if a field has been set.

### SetColorKeyNil

`func (o *UpdateCustomFieldEnumValue) SetColorKeyNil(b bool)`

 SetColorKeyNil sets the value for ColorKey to be an explicit nil

### UnsetColorKey
`func (o *UpdateCustomFieldEnumValue) UnsetColorKey()`

UnsetColorKey ensures that no value is present for ColorKey, not even an explicit nil
### GetEnabled

`func (o *UpdateCustomFieldEnumValue) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCustomFieldEnumValue) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCustomFieldEnumValue) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCustomFieldEnumValue) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


