# UpdateCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new name of the Category. | [optional] 
**Color** | Pointer to **NullableString** | The hex color to be displayed with the Category (for example, \&quot;#ff0000\&quot;). | [optional] 
**Archived** | Pointer to **bool** | A true/false boolean indicating if the Category has been archived. | [optional] 

## Methods

### NewUpdateCategory

`func NewUpdateCategory() *UpdateCategory`

NewUpdateCategory instantiates a new UpdateCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCategoryWithDefaults

`func NewUpdateCategoryWithDefaults() *UpdateCategory`

NewUpdateCategoryWithDefaults instantiates a new UpdateCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *UpdateCategory) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateCategory) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateCategory) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateCategory) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *UpdateCategory) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *UpdateCategory) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetArchived

`func (o *UpdateCategory) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UpdateCategory) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UpdateCategory) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UpdateCategory) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


