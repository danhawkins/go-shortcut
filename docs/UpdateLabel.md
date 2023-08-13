# UpdateLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new name of the label. | [optional] 
**Description** | Pointer to **string** | The new description of the label. | [optional] 
**Color** | Pointer to **NullableString** | The hex color to be displayed with the Label (for example, \&quot;#ff0000\&quot;). | [optional] 
**Archived** | Pointer to **bool** | A true/false boolean indicating if the Label has been archived. | [optional] 

## Methods

### NewUpdateLabel

`func NewUpdateLabel() *UpdateLabel`

NewUpdateLabel instantiates a new UpdateLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLabelWithDefaults

`func NewUpdateLabelWithDefaults() *UpdateLabel`

NewUpdateLabelWithDefaults instantiates a new UpdateLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLabel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLabel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLabel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLabel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLabel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLabel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColor

`func (o *UpdateLabel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateLabel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateLabel) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateLabel) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *UpdateLabel) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *UpdateLabel) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetArchived

`func (o *UpdateLabel) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateLabel) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateLabel) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateLabel) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


