# CreateCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new Category. | 
**Color** | Pointer to **string** | The hex color to be displayed with the Category (for example, \&quot;#ff0000\&quot;). | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Category has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 
**Type** | **string** | The type of entity this Category is associated with; currently Milestone is the only type of Category. | 

## Methods

### NewCreateCategory

`func NewCreateCategory(name string, type_ string, ) *CreateCategory`

NewCreateCategory instantiates a new CreateCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCategoryWithDefaults

`func NewCreateCategoryWithDefaults() *CreateCategory`

NewCreateCategoryWithDefaults instantiates a new CreateCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCategory) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *CreateCategory) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateCategory) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateCategory) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateCategory) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateCategory) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateCategory) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateCategory) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateCategory) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetType

`func (o *CreateCategory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCategory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCategory) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


