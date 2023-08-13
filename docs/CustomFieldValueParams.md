# CustomFieldValueParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** | The unique public ID for the CustomField. | 
**ValueId** | **string** | The unique public ID for the CustomFieldEnumValue. | 
**Value** | Pointer to **string** | A literal value for the CustomField. Currently ignored. | [optional] 

## Methods

### NewCustomFieldValueParams

`func NewCustomFieldValueParams(fieldId string, valueId string, ) *CustomFieldValueParams`

NewCustomFieldValueParams instantiates a new CustomFieldValueParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldValueParamsWithDefaults

`func NewCustomFieldValueParamsWithDefaults() *CustomFieldValueParams`

NewCustomFieldValueParamsWithDefaults instantiates a new CustomFieldValueParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *CustomFieldValueParams) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *CustomFieldValueParams) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *CustomFieldValueParams) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetValueId

`func (o *CustomFieldValueParams) GetValueId() string`

GetValueId returns the ValueId field if non-nil, zero value otherwise.

### GetValueIdOk

`func (o *CustomFieldValueParams) GetValueIdOk() (*string, bool)`

GetValueIdOk returns a tuple with the ValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueId

`func (o *CustomFieldValueParams) SetValueId(v string)`

SetValueId sets ValueId field to given value.


### GetValue

`func (o *CustomFieldValueParams) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldValueParams) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldValueParams) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldValueParams) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


