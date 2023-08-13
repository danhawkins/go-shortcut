# StoryCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** | The unique public ID for a CustomField. | 
**ValueId** | **string** | The unique public ID for a CustomFieldEnumValue. | 
**Value** | **string** | A string representation of the value, if applicable. | 

## Methods

### NewStoryCustomField

`func NewStoryCustomField(fieldId string, valueId string, value string, ) *StoryCustomField`

NewStoryCustomField instantiates a new StoryCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoryCustomFieldWithDefaults

`func NewStoryCustomFieldWithDefaults() *StoryCustomField`

NewStoryCustomFieldWithDefaults instantiates a new StoryCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *StoryCustomField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *StoryCustomField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *StoryCustomField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetValueId

`func (o *StoryCustomField) GetValueId() string`

GetValueId returns the ValueId field if non-nil, zero value otherwise.

### GetValueIdOk

`func (o *StoryCustomField) GetValueIdOk() (*string, bool)`

GetValueIdOk returns a tuple with the ValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueId

`func (o *StoryCustomField) SetValueId(v string)`

SetValueId sets ValueId field to given value.


### GetValue

`func (o *StoryCustomField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StoryCustomField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StoryCustomField) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


