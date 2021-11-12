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

// InlineResponse20090Data struct for InlineResponse20090Data
type InlineResponse20090Data struct {
	OtherProfits []InlineResponse20090DataOtherProfits `json:"otherProfits"`
	// Total Rows
	TotalNum int64 `json:"totalNum"`
	// Rows per page
	PageSize int64 `json:"pageSize"`
}

// NewInlineResponse20090Data instantiates a new InlineResponse20090Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20090Data(otherProfits []InlineResponse20090DataOtherProfits, totalNum int64, pageSize int64) *InlineResponse20090Data {
	this := InlineResponse20090Data{}
	this.OtherProfits = otherProfits
	this.TotalNum = totalNum
	this.PageSize = pageSize
	return &this
}

// NewInlineResponse20090DataWithDefaults instantiates a new InlineResponse20090Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20090DataWithDefaults() *InlineResponse20090Data {
	this := InlineResponse20090Data{}
	return &this
}

// GetOtherProfits returns the OtherProfits field value
func (o *InlineResponse20090Data) GetOtherProfits() []InlineResponse20090DataOtherProfits {
	if o == nil {
		var ret []InlineResponse20090DataOtherProfits
		return ret
	}

	return o.OtherProfits
}

// GetOtherProfitsOk returns a tuple with the OtherProfits field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20090Data) GetOtherProfitsOk() (*[]InlineResponse20090DataOtherProfits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OtherProfits, true
}

// SetOtherProfits sets field value
func (o *InlineResponse20090Data) SetOtherProfits(v []InlineResponse20090DataOtherProfits) {
	o.OtherProfits = v
}

// GetTotalNum returns the TotalNum field value
func (o *InlineResponse20090Data) GetTotalNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20090Data) GetTotalNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNum, true
}

// SetTotalNum sets field value
func (o *InlineResponse20090Data) SetTotalNum(v int64) {
	o.TotalNum = v
}

// GetPageSize returns the PageSize field value
func (o *InlineResponse20090Data) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20090Data) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *InlineResponse20090Data) SetPageSize(v int64) {
	o.PageSize = v
}

func (o InlineResponse20090Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["otherProfits"] = o.OtherProfits
	}
	if true {
		toSerialize["totalNum"] = o.TotalNum
	}
	if true {
		toSerialize["pageSize"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20090Data struct {
	value *InlineResponse20090Data
	isSet bool
}

func (v NullableInlineResponse20090Data) Get() *InlineResponse20090Data {
	return v.value
}

func (v *NullableInlineResponse20090Data) Set(val *InlineResponse20090Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20090Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20090Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20090Data(val *InlineResponse20090Data) *NullableInlineResponse20090Data {
	return &NullableInlineResponse20090Data{value: val, isSet: true}
}

func (v NullableInlineResponse20090Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20090Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
