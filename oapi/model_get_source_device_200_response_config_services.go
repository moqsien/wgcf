/*
 * untitled API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 536
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oapi

import (
	"encoding/json"
)

// GetSourceDevice200ResponseConfigServices struct for GetSourceDevice200ResponseConfigServices
type GetSourceDevice200ResponseConfigServices struct {
	HttpProxy string `json:"http_proxy"`
}

// NewGetSourceDevice200ResponseConfigServices instantiates a new GetSourceDevice200ResponseConfigServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSourceDevice200ResponseConfigServices(httpProxy string, ) *GetSourceDevice200ResponseConfigServices {
	this := GetSourceDevice200ResponseConfigServices{}
	this.HttpProxy = httpProxy
	return &this
}

// NewGetSourceDevice200ResponseConfigServicesWithDefaults instantiates a new GetSourceDevice200ResponseConfigServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSourceDevice200ResponseConfigServicesWithDefaults() *GetSourceDevice200ResponseConfigServices {
	this := GetSourceDevice200ResponseConfigServices{}
	return &this
}

// GetHttpProxy returns the HttpProxy field value
func (o *GetSourceDevice200ResponseConfigServices) GetHttpProxy() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.HttpProxy
}

// GetHttpProxyOk returns a tuple with the HttpProxy field value
// and a boolean to check if the value has been set.
func (o *GetSourceDevice200ResponseConfigServices) GetHttpProxyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HttpProxy, true
}

// SetHttpProxy sets field value
func (o *GetSourceDevice200ResponseConfigServices) SetHttpProxy(v string) {
	o.HttpProxy = v
}

func (o GetSourceDevice200ResponseConfigServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["http_proxy"] = o.HttpProxy
	}
	return json.Marshal(toSerialize)
}

type NullableGetSourceDevice200ResponseConfigServices struct {
	value *GetSourceDevice200ResponseConfigServices
	isSet bool
}

func (v NullableGetSourceDevice200ResponseConfigServices) Get() *GetSourceDevice200ResponseConfigServices {
	return v.value
}

func (v *NullableGetSourceDevice200ResponseConfigServices) Set(val *GetSourceDevice200ResponseConfigServices) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSourceDevice200ResponseConfigServices) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSourceDevice200ResponseConfigServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSourceDevice200ResponseConfigServices(val *GetSourceDevice200ResponseConfigServices) *NullableGetSourceDevice200ResponseConfigServices {
	return &NullableGetSourceDevice200ResponseConfigServices{value: val, isSet: true}
}

func (v NullableGetSourceDevice200ResponseConfigServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSourceDevice200ResponseConfigServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


