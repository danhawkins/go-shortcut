# CustomFieldEnumValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique public ID for the Custom Field. | 
**Value** | **string** | A string value within the domain of this Custom Field. | 
**Position** | **int64** | An integer indicating the position of this Value with respect to the other CustomFieldEnumValues in the enumeration. | 
**ColorKey** | **NullableString** | A color key associated with this CustomFieldEnumValue. | 
**EntityType** | **string** | A string description of this resource. | 
**Enabled** | **bool** | When true, the CustomFieldEnumValue can be selected for the CustomField. | 

## Methods

### NewCustomFieldEnumValue

`func NewCustomFieldEnumValue(id string, value string, position int64, colorKey NullableString, entityType string, enabled bool, ) *CustomFieldEnumValue`

NewCustomFieldEnumValue instantiates a new CustomFieldEnumValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldEnumValueWithDefaults

`func NewCustomFieldEnumValueWithDefaults() *CustomFieldEnumValue`

NewCustomFieldEnumValueWithDefaults instantiates a new CustomFieldEnumValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldEnumValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldEnumValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldEnumValue) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *CustomFieldEnumValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldEnumValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldEnumValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetPosition

`func (o *CustomFieldEnumValue) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CustomFieldEnumValue) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CustomFieldEnumValue) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetColorKey

`func (o *CustomFieldEnumValue) GetColorKey() string`

GetColorKey returns the ColorKey field if non-nil, zero value otherwise.

### GetColorKeyOk

`func (o *CustomFieldEnumValue) GetColorKeyOk() (*string, bool)`

GetColorKeyOk returns a tuple with the ColorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorKey

`func (o *CustomFieldEnumValue) SetColorKey(v string)`

SetColorKey sets ColorKey field to given value.


### SetColorKeyNil

`func (o *CustomFieldEnumValue) SetColorKeyNil(b bool)`

 SetColorKeyNil sets the value for ColorKey to be an explicit nil

### UnsetColorKey
`func (o *CustomFieldEnumValue) UnsetColorKey()`

UnsetColorKey ensures that no value is present for ColorKey, not even an explicit nil
### GetEntityType

`func (o *CustomFieldEnumValue) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CustomFieldEnumValue) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CustomFieldEnumValue) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetEnabled

`func (o *CustomFieldEnumValue) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomFieldEnumValue) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomFieldEnumValue) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


