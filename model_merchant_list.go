/*
Merchant APIs

universalmacro

API version: 0.0.3
Contact: chenyunda218@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MerchantList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantList{}

// MerchantList struct for MerchantList
type MerchantList struct {
	Items []Merchant `json:"items"`
	Pagination Pagination `json:"pagination"`
}

type _MerchantList MerchantList

// NewMerchantList instantiates a new MerchantList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantList(items []Merchant, pagination Pagination) *MerchantList {
	this := MerchantList{}
	this.Items = items
	this.Pagination = pagination
	return &this
}

// NewMerchantListWithDefaults instantiates a new MerchantList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantListWithDefaults() *MerchantList {
	this := MerchantList{}
	return &this
}

// GetItems returns the Items field value
func (o *MerchantList) GetItems() []Merchant {
	if o == nil {
		var ret []Merchant
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *MerchantList) GetItemsOk() ([]Merchant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *MerchantList) SetItems(v []Merchant) {
	o.Items = v
}

// GetPagination returns the Pagination field value
func (o *MerchantList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *MerchantList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *MerchantList) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o MerchantList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

func (o *MerchantList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"pagination",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMerchantList := _MerchantList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMerchantList)

	if err != nil {
		return err
	}

	*o = MerchantList(varMerchantList)

	return err
}

type NullableMerchantList struct {
	value *MerchantList
	isSet bool
}

func (v NullableMerchantList) Get() *MerchantList {
	return v.value
}

func (v *NullableMerchantList) Set(val *MerchantList) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantList) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantList(val *MerchantList) *NullableMerchantList {
	return &NullableMerchantList{value: val, isSet: true}
}

func (v NullableMerchantList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


