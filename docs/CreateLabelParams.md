# CreateLabelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new Label. | 
**Description** | Pointer to **string** | The description of the new Label. | [optional] 
**Color** | Pointer to **string** | The hex color to be displayed with the Label (for example, \&quot;#ff0000\&quot;). | [optional] 
**ExternalId** | Pointer to **string** | This field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here. | [optional] 

## Methods

### NewCreateLabelParams

`func NewCreateLabelParams(name string, ) *CreateLabelParams`

NewCreateLabelParams instantiates a new CreateLabelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLabelParamsWithDefaults

`func NewCreateLabelParamsWithDefaults() *CreateLabelParams`

NewCreateLabelParamsWithDefaults instantiates a new CreateLabelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLabelParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLabelParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLabelParams) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateLabelParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLabelParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLabelParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLabelParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColor

`func (o *CreateLabelParams) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateLabelParams) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateLabelParams) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateLabelParams) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateLabelParams) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateLabelParams) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateLabelParams) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateLabelParams) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


