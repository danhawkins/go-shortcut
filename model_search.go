/*
Shortcut API

Shortcut API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Search type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Search{}

// Search struct for Search
type Search struct {
	// See our help center article on [search operators](https://help.shortcut.com/hc/en-us/articles/360000046646-Search-Operators)
	Query string `json:"query"`
	// The number of search results to include in a page. Minimum of 1 and maximum of 25.
	PageSize *int64 `json:"page_size,omitempty"`
	// The amount of detail included in each result item.    \"full\" will include all descriptions and comments and more fields on    related items such as pull requests, branches and tasks.    \"slim\" omits larger fulltext fields such as descriptions and comments    and only references related items by id.    The default is \"full\".
	Detail *string `json:"detail,omitempty"`
	// The next page token.
	Next *string `json:"next,omitempty"`
	Include *string `json:"include,omitempty"`
	// A collection of entity_types to search. Defaults to story and epic. Supports: epic, iteration, milestone, story.
	EntityTypes []string `json:"entity_types,omitempty"`
}

// NewSearch instantiates a new Search object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearch(query string) *Search {
	this := Search{}
	this.Query = query
	return &this
}

// NewSearchWithDefaults instantiates a new Search object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchWithDefaults() *Search {
	this := Search{}
	return &this
}

// GetQuery returns the Query field value
func (o *Search) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *Search) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *Search) SetQuery(v string) {
	o.Query = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *Search) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *Search) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *Search) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *Search) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *Search) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *Search) SetDetail(v string) {
	o.Detail = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *Search) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *Search) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *Search) SetNext(v string) {
	o.Next = &v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *Search) GetInclude() string {
	if o == nil || IsNil(o.Include) {
		var ret string
		return ret
	}
	return *o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetIncludeOk() (*string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *Search) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given string and assigns it to the Include field.
func (o *Search) SetInclude(v string) {
	o.Include = &v
}

// GetEntityTypes returns the EntityTypes field value if set, zero value otherwise.
func (o *Search) GetEntityTypes() []string {
	if o == nil || IsNil(o.EntityTypes) {
		var ret []string
		return ret
	}
	return o.EntityTypes
}

// GetEntityTypesOk returns a tuple with the EntityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetEntityTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.EntityTypes) {
		return nil, false
	}
	return o.EntityTypes, true
}

// HasEntityTypes returns a boolean if a field has been set.
func (o *Search) HasEntityTypes() bool {
	if o != nil && !IsNil(o.EntityTypes) {
		return true
	}

	return false
}

// SetEntityTypes gets a reference to the given []string and assigns it to the EntityTypes field.
func (o *Search) SetEntityTypes(v []string) {
	o.EntityTypes = v
}

func (o Search) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Search) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	if !IsNil(o.EntityTypes) {
		toSerialize["entity_types"] = o.EntityTypes
	}
	return toSerialize, nil
}

type NullableSearch struct {
	value *Search
	isSet bool
}

func (v NullableSearch) Get() *Search {
	return v.value
}

func (v *NullableSearch) Set(val *Search) {
	v.value = val
	v.isSet = true
}

func (v NullableSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearch(val *Search) *NullableSearch {
	return &NullableSearch{value: val, isSet: true}
}

func (v NullableSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


