/*
Binance Public Spot API

OpenAPI Specifications for the Binance Public Spot API generated with [binance/binance-api-swagger/blob/master/spot_api.yaml](https://github.com/binance/binance-api-swagger/blob/master/spot_api.yaml) with commit [v1.2.0 release](https://github.com/binance/binance-api-swagger/commit/60d14be031c031600c853d5cdab86db5ab73603e)  API documents:   - [https://github.com/binance/binance-spot-api-docs](https://github.com/binance/binance-spot-api-docs)   - [https://binance-docs.github.io/apidocs/spot/en](https://binance-docs.github.io/apidocs/spot/en)

API version: 1.0
Contact: qishiwenjun@163.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binanceapi

import (
	"encoding/json"
)

// SnapshotMargin struct for SnapshotMargin
type SnapshotMargin struct {
	Code        int64                       `json:"code"`
	Msg         string                      `json:"msg"`
	SnapshotVos []SnapshotMarginSnapshotVos `json:"snapshotVos"`
}

// NewSnapshotMargin instantiates a new SnapshotMargin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotMargin(code int64, msg string, snapshotVos []SnapshotMarginSnapshotVos) *SnapshotMargin {
	this := SnapshotMargin{}
	this.Code = code
	this.Msg = msg
	this.SnapshotVos = snapshotVos
	return &this
}

// NewSnapshotMarginWithDefaults instantiates a new SnapshotMargin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotMarginWithDefaults() *SnapshotMargin {
	this := SnapshotMargin{}
	return &this
}

// GetCode returns the Code field value
func (o *SnapshotMargin) GetCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SnapshotMargin) GetCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SnapshotMargin) SetCode(v int64) {
	o.Code = v
}

// GetMsg returns the Msg field value
func (o *SnapshotMargin) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *SnapshotMargin) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *SnapshotMargin) SetMsg(v string) {
	o.Msg = v
}

// GetSnapshotVos returns the SnapshotVos field value
func (o *SnapshotMargin) GetSnapshotVos() []SnapshotMarginSnapshotVos {
	if o == nil {
		var ret []SnapshotMarginSnapshotVos
		return ret
	}

	return o.SnapshotVos
}

// GetSnapshotVosOk returns a tuple with the SnapshotVos field value
// and a boolean to check if the value has been set.
func (o *SnapshotMargin) GetSnapshotVosOk() (*[]SnapshotMarginSnapshotVos, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotVos, true
}

// SetSnapshotVos sets field value
func (o *SnapshotMargin) SetSnapshotVos(v []SnapshotMarginSnapshotVos) {
	o.SnapshotVos = v
}

func (o SnapshotMargin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["msg"] = o.Msg
	}
	if true {
		toSerialize["snapshotVos"] = o.SnapshotVos
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotMargin struct {
	value *SnapshotMargin
	isSet bool
}

func (v NullableSnapshotMargin) Get() *SnapshotMargin {
	return v.value
}

func (v *NullableSnapshotMargin) Set(val *SnapshotMargin) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotMargin) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotMargin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotMargin(val *SnapshotMargin) *NullableSnapshotMargin {
	return &NullableSnapshotMargin{value: val, isSet: true}
}

func (v NullableSnapshotMargin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotMargin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
