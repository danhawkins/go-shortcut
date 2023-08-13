# UnusableEntitlementError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonTag** | **string** | The tag for violating an entitlement action. | 
**EntitlementTag** | **string** | Short tag describing the unusable entitlement action taken by the user. | 
**Message** | **string** | Message displayed to the user on why their action failed. | 

## Methods

### NewUnusableEntitlementError

`func NewUnusableEntitlementError(reasonTag string, entitlementTag string, message string, ) *UnusableEntitlementError`

NewUnusableEntitlementError instantiates a new UnusableEntitlementError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnusableEntitlementErrorWithDefaults

`func NewUnusableEntitlementErrorWithDefaults() *UnusableEntitlementError`

NewUnusableEntitlementErrorWithDefaults instantiates a new UnusableEntitlementError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonTag

`func (o *UnusableEntitlementError) GetReasonTag() string`

GetReasonTag returns the ReasonTag field if non-nil, zero value otherwise.

### GetReasonTagOk

`func (o *UnusableEntitlementError) GetReasonTagOk() (*string, bool)`

GetReasonTagOk returns a tuple with the ReasonTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonTag

`func (o *UnusableEntitlementError) SetReasonTag(v string)`

SetReasonTag sets ReasonTag field to given value.


### GetEntitlementTag

`func (o *UnusableEntitlementError) GetEntitlementTag() string`

GetEntitlementTag returns the EntitlementTag field if non-nil, zero value otherwise.

### GetEntitlementTagOk

`func (o *UnusableEntitlementError) GetEntitlementTagOk() (*string, bool)`

GetEntitlementTagOk returns a tuple with the EntitlementTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementTag

`func (o *UnusableEntitlementError) SetEntitlementTag(v string)`

SetEntitlementTag sets EntitlementTag field to given value.


### GetMessage

`func (o *UnusableEntitlementError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnusableEntitlementError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnusableEntitlementError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


