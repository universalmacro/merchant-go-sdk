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

// checks if the UpdatePasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePasswordRequest{}

// UpdatePasswordRequest struct for UpdatePasswordRequest
type UpdatePasswordRequest struct {
	Password string `json:"password"`
	OldPassword *string `json:"oldPassword,omitempty"`
}

type _UpdatePasswordRequest UpdatePasswordRequest

// NewUpdatePasswordRequest instantiates a new UpdatePasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePasswordRequest(password string) *UpdatePasswordRequest {
	this := UpdatePasswordRequest{}
	this.Password = password
	return &this
}

// NewUpdatePasswordRequestWithDefaults instantiates a new UpdatePasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePasswordRequestWithDefaults() *UpdatePasswordRequest {
	this := UpdatePasswordRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *UpdatePasswordRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UpdatePasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UpdatePasswordRequest) SetPassword(v string) {
	o.Password = v
}

// GetOldPassword returns the OldPassword field value if set, zero value otherwise.
func (o *UpdatePasswordRequest) GetOldPassword() string {
	if o == nil || IsNil(o.OldPassword) {
		var ret string
		return ret
	}
	return *o.OldPassword
}

// GetOldPasswordOk returns a tuple with the OldPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePasswordRequest) GetOldPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.OldPassword) {
		return nil, false
	}
	return o.OldPassword, true
}

// HasOldPassword returns a boolean if a field has been set.
func (o *UpdatePasswordRequest) HasOldPassword() bool {
	if o != nil && !IsNil(o.OldPassword) {
		return true
	}

	return false
}

// SetOldPassword gets a reference to the given string and assigns it to the OldPassword field.
func (o *UpdatePasswordRequest) SetOldPassword(v string) {
	o.OldPassword = &v
}

func (o UpdatePasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	if !IsNil(o.OldPassword) {
		toSerialize["oldPassword"] = o.OldPassword
	}
	return toSerialize, nil
}

func (o *UpdatePasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
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

	varUpdatePasswordRequest := _UpdatePasswordRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdatePasswordRequest)

	if err != nil {
		return err
	}

	*o = UpdatePasswordRequest(varUpdatePasswordRequest)

	return err
}

type NullableUpdatePasswordRequest struct {
	value *UpdatePasswordRequest
	isSet bool
}

func (v NullableUpdatePasswordRequest) Get() *UpdatePasswordRequest {
	return v.value
}

func (v *NullableUpdatePasswordRequest) Set(val *UpdatePasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePasswordRequest(val *UpdatePasswordRequest) *NullableUpdatePasswordRequest {
	return &NullableUpdatePasswordRequest{value: val, isSet: true}
}

func (v NullableUpdatePasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


