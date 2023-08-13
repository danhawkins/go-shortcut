# CustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A string description of the CustomField | [optional] 
**IconSetIdentifier** | Pointer to **string** | A string that represents the icon that corresponds to this custom field. | [optional] 
**EntityType** | **string** | A string description of this resource. | 
**StoryTypes** | Pointer to **[]string** | The types of stories this CustomField is scoped to. | [optional] 
**Name** | **string** | The name of the Custom Field. | 
**FixedPosition** | Pointer to **bool** | When true, the CustomFieldEnumValues may not be reordered. | [optional] 
**UpdatedAt** | **time.Time** | The instant when this CustomField was last updated. | 
**Id** | **string** | The unique public ID for the CustomField. | 
**Values** | Pointer to [**[]CustomFieldEnumValue**](CustomFieldEnumValue.md) | A collection of legal values for a CustomField. | [optional] 
**FieldType** | **string** | The type of Custom Field, eg. &#39;enum&#39;. | 
**Position** | **int64** | An integer indicating the position of this Custom Field with respect to the other CustomField | 
**CanonicalName** | Pointer to **string** | The canonical name for a Shortcut-defined field. | [optional] 
**Enabled** | **bool** | When true, the CustomField can be applied to entities in the Workspace. | 
**CreatedAt** | **time.Time** | The instant when this CustomField was created. | 

## Methods

### NewCustomField

`func NewCustomField(entityType string, name string, updatedAt time.Time, id string, fieldType string, position int64, enabled bool, createdAt time.Time, ) *CustomField`

NewCustomField instantiates a new CustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldWithDefaults

`func NewCustomFieldWithDefaults() *CustomField`

NewCustomFieldWithDefaults instantiates a new CustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconSetIdentifier

`func (o *CustomField) GetIconSetIdentifier() string`

GetIconSetIdentifier returns the IconSetIdentifier field if non-nil, zero value otherwise.

### GetIconSetIdentifierOk

`func (o *CustomField) GetIconSetIdentifierOk() (*string, bool)`

GetIconSetIdentifierOk returns a tuple with the IconSetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconSetIdentifier

`func (o *CustomField) SetIconSetIdentifier(v string)`

SetIconSetIdentifier sets IconSetIdentifier field to given value.

### HasIconSetIdentifier

`func (o *CustomField) HasIconSetIdentifier() bool`

HasIconSetIdentifier returns a boolean if a field has been set.

### GetEntityType

`func (o *CustomField) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CustomField) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CustomField) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetStoryTypes

`func (o *CustomField) GetStoryTypes() []string`

GetStoryTypes returns the StoryTypes field if non-nil, zero value otherwise.

### GetStoryTypesOk

`func (o *CustomField) GetStoryTypesOk() (*[]string, bool)`

GetStoryTypesOk returns a tuple with the StoryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryTypes

`func (o *CustomField) SetStoryTypes(v []string)`

SetStoryTypes sets StoryTypes field to given value.

### HasStoryTypes

`func (o *CustomField) HasStoryTypes() bool`

HasStoryTypes returns a boolean if a field has been set.

### GetName

`func (o *CustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomField) SetName(v string)`

SetName sets Name field to given value.


### GetFixedPosition

`func (o *CustomField) GetFixedPosition() bool`

GetFixedPosition returns the FixedPosition field if non-nil, zero value otherwise.

### GetFixedPositionOk

`func (o *CustomField) GetFixedPositionOk() (*bool, bool)`

GetFixedPositionOk returns a tuple with the FixedPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedPosition

`func (o *CustomField) SetFixedPosition(v bool)`

SetFixedPosition sets FixedPosition field to given value.

### HasFixedPosition

`func (o *CustomField) HasFixedPosition() bool`

HasFixedPosition returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomField) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomField) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomField) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetId

`func (o *CustomField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomField) SetId(v string)`

SetId sets Id field to given value.


### GetValues

`func (o *CustomField) GetValues() []CustomFieldEnumValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CustomField) GetValuesOk() (*[]CustomFieldEnumValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CustomField) SetValues(v []CustomFieldEnumValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *CustomField) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetFieldType

`func (o *CustomField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *CustomField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *CustomField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetPosition

`func (o *CustomField) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CustomField) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CustomField) SetPosition(v int64)`

SetPosition sets Position field to given value.


### GetCanonicalName

`func (o *CustomField) GetCanonicalName() string`

GetCanonicalName returns the CanonicalName field if non-nil, zero value otherwise.

### GetCanonicalNameOk

`func (o *CustomField) GetCanonicalNameOk() (*string, bool)`

GetCanonicalNameOk returns a tuple with the CanonicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalName

`func (o *CustomField) SetCanonicalName(v string)`

SetCanonicalName sets CanonicalName field to given value.

### HasCanonicalName

`func (o *CustomField) HasCanonicalName() bool`

HasCanonicalName returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomField) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomField) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomField) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedAt

`func (o *CustomField) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomField) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomField) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


