# HistoryReferenceCustomFieldEnumValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | The type of entity referenced. | 
**FieldName** | **string** | The name as it is displayed to the user of the parent custom-field of this enum value. | 
**IntegerValue** | **NullableInt64** | The custom-field enum value as a string. | 
**FieldEnabled** | **bool** | Whether or not the custom-field is enabled. | 
**Id** | **map[string]interface{}** | The ID of the entity referenced. | 
**FieldType** | **string** | The type variety of the parent custom-field of this enum value. | 
**FieldId** | **string** | The public-id of the parent custom-field of this enum value. | 
**StringValue** | **NullableString** | The custom-field enum value as a string. | 
**EnumValueEnabled** | **NullableBool** | Whether or not the custom-field enum value is enabled. | 

## Methods

### NewHistoryReferenceCustomFieldEnumValue

`func NewHistoryReferenceCustomFieldEnumValue(entityType string, fieldName string, integerValue NullableInt64, fieldEnabled bool, id map[string]interface{}, fieldType string, fieldId string, stringValue NullableString, enumValueEnabled NullableBool, ) *HistoryReferenceCustomFieldEnumValue`

NewHistoryReferenceCustomFieldEnumValue instantiates a new HistoryReferenceCustomFieldEnumValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryReferenceCustomFieldEnumValueWithDefaults

`func NewHistoryReferenceCustomFieldEnumValueWithDefaults() *HistoryReferenceCustomFieldEnumValue`

NewHistoryReferenceCustomFieldEnumValueWithDefaults instantiates a new HistoryReferenceCustomFieldEnumValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *HistoryReferenceCustomFieldEnumValue) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *HistoryReferenceCustomFieldEnumValue) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *HistoryReferenceCustomFieldEnumValue) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetFieldName

`func (o *HistoryReferenceCustomFieldEnumValue) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *HistoryReferenceCustomFieldEnumValue) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *HistoryReferenceCustomFieldEnumValue) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetIntegerValue

`func (o *HistoryReferenceCustomFieldEnumValue) GetIntegerValue() int64`

GetIntegerValue returns the IntegerValue field if non-nil, zero value otherwise.

### GetIntegerValueOk

`func (o *HistoryReferenceCustomFieldEnumValue) GetIntegerValueOk() (*int64, bool)`

GetIntegerValueOk returns a tuple with the IntegerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegerValue

`func (o *HistoryReferenceCustomFieldEnumValue) SetIntegerValue(v int64)`

SetIntegerValue sets IntegerValue field to given value.


### SetIntegerValueNil

`func (o *HistoryReferenceCustomFieldEnumValue) SetIntegerValueNil(b bool)`

 SetIntegerValueNil sets the value for IntegerValue to be an explicit nil

### UnsetIntegerValue
`func (o *HistoryReferenceCustomFieldEnumValue) UnsetIntegerValue()`

UnsetIntegerValue ensures that no value is present for IntegerValue, not even an explicit nil
### GetFieldEnabled

`func (o *HistoryReferenceCustomFieldEnumValue) GetFieldEnabled() bool`

GetFieldEnabled returns the FieldEnabled field if non-nil, zero value otherwise.

### GetFieldEnabledOk

`func (o *HistoryReferenceCustomFieldEnumValue) GetFieldEnabledOk() (*bool, bool)`

GetFieldEnabledOk returns a tuple with the FieldEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldEnabled

`func (o *HistoryReferenceCustomFieldEnumValue) SetFieldEnabled(v bool)`

SetFieldEnabled sets FieldEnabled field to given value.


### GetId

`func (o *HistoryReferenceCustomFieldEnumValue) GetId() map[string]interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryReferenceCustomFieldEnumValue) GetIdOk() (*map[string]interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryReferenceCustomFieldEnumValue) SetId(v map[string]interface{})`

SetId sets Id field to given value.


### GetFieldType

`func (o *HistoryReferenceCustomFieldEnumValue) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *HistoryReferenceCustomFieldEnumValue) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *HistoryReferenceCustomFieldEnumValue) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetFieldId

`func (o *HistoryReferenceCustomFieldEnumValue) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *HistoryReferenceCustomFieldEnumValue) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *HistoryReferenceCustomFieldEnumValue) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetStringValue

`func (o *HistoryReferenceCustomFieldEnumValue) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *HistoryReferenceCustomFieldEnumValue) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *HistoryReferenceCustomFieldEnumValue) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.


### SetStringValueNil

`func (o *HistoryReferenceCustomFieldEnumValue) SetStringValueNil(b bool)`

 SetStringValueNil sets the value for StringValue to be an explicit nil

### UnsetStringValue
`func (o *HistoryReferenceCustomFieldEnumValue) UnsetStringValue()`

UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil
### GetEnumValueEnabled

`func (o *HistoryReferenceCustomFieldEnumValue) GetEnumValueEnabled() bool`

GetEnumValueEnabled returns the EnumValueEnabled field if non-nil, zero value otherwise.

### GetEnumValueEnabledOk

`func (o *HistoryReferenceCustomFieldEnumValue) GetEnumValueEnabledOk() (*bool, bool)`

GetEnumValueEnabledOk returns a tuple with the EnumValueEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValueEnabled

`func (o *HistoryReferenceCustomFieldEnumValue) SetEnumValueEnabled(v bool)`

SetEnumValueEnabled sets EnumValueEnabled field to given value.


### SetEnumValueEnabledNil

`func (o *HistoryReferenceCustomFieldEnumValue) SetEnumValueEnabledNil(b bool)`

 SetEnumValueEnabledNil sets the value for EnumValueEnabled to be an explicit nil

### UnsetEnumValueEnabled
`func (o *HistoryReferenceCustomFieldEnumValue) UnsetEnumValueEnabled()`

UnsetEnumValueEnabled ensures that no value is present for EnumValueEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


