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

// checks if the PhoneNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneNumber{}

// PhoneNumber struct for PhoneNumber
type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	Number string `json:"number"`
}

type _PhoneNumber PhoneNumber

// NewPhoneNumber instantiates a new PhoneNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneNumber(countryCode string, number string) *PhoneNumber {
	this := PhoneNumber{}
	this.CountryCode = countryCode
	this.Number = number
	return &this
}

// NewPhoneNumberWithDefaults instantiates a new PhoneNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneNumberWithDefaults() *PhoneNumber {
	this := PhoneNumber{}
	return &this
}

// GetCountryCode returns the CountryCode field value
func (o *PhoneNumber) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *PhoneNumber) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *PhoneNumber) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetNumber returns the Number field value
func (o *PhoneNumber) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *PhoneNumber) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *PhoneNumber) SetNumber(v string) {
	o.Number = v
}

func (o PhoneNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryCode"] = o.CountryCode
	toSerialize["number"] = o.Number
	return toSerialize, nil
}

func (o *PhoneNumber) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"countryCode",
		"number",
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

	varPhoneNumber := _PhoneNumber{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPhoneNumber)

	if err != nil {
		return err
	}

	*o = PhoneNumber(varPhoneNumber)

	return err
}

type NullablePhoneNumber struct {
	value *PhoneNumber
	isSet bool
}

func (v NullablePhoneNumber) Get() *PhoneNumber {
	return v.value
}

func (v *NullablePhoneNumber) Set(val *PhoneNumber) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneNumber) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneNumber(val *PhoneNumber) *NullablePhoneNumber {
	return &NullablePhoneNumber{value: val, isSet: true}
}

func (v NullablePhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


