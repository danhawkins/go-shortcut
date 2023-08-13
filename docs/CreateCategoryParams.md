# CreateCategoryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new Category. | 
**Color** | Pointer to **string** | The hex color to be displayed with the Category (for example, \&quot;#ff0000\&quot;). | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Category has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 

## Methods

### NewCreateCategoryParams

`func NewCreateCategoryParams(name string, ) *CreateCategoryParams`

NewCreateCategoryParams instantiates a new CreateCategoryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCategoryParamsWithDefaults

`func NewCreateCategoryParamsWithDefaults() *CreateCategoryParams`

NewCreateCategoryParamsWithDefaults instantiates a new CreateCategoryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCategoryParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCategoryParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCategoryParams) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *CreateCategoryParams) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateCategoryParams) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateCategoryParams) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateCategoryParams) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateCategoryParams) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateCategoryParams) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateCategoryParams) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateCategoryParams) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


