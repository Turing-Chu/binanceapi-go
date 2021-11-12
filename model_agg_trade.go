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

// AggTrade struct for AggTrade
type AggTrade struct {
	// Aggregate tradeId
	A int64 `json:"a"`
	// Price
	P string `json:"p"`
	// Quantity
	Q string `json:"q"`
	// First tradeId
	F int64 `json:"f"`
	// Last tradeId
	L int64 `json:"l"`
	// Timestamp
	T bool `json:"T"`
	// Was the buyer the maker?
	M bool `json:"m"`
	// Was the trade the best price match?
	M bool `json:"M"`
}

// NewAggTrade instantiates a new AggTrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggTrade(a int64, p string, q string, f int64, l int64, t bool, m bool, m bool) *AggTrade {
	this := AggTrade{}
	this.A = a
	this.P = p
	this.Q = q
	this.F = f
	this.L = l
	this.T = t
	this.M = m
	this.M = m
	return &this
}

// NewAggTradeWithDefaults instantiates a new AggTrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggTradeWithDefaults() *AggTrade {
	this := AggTrade{}
	return &this
}

// GetA returns the A field value
func (o *AggTrade) GetA() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.A
}

// GetAOk returns a tuple with the A field value
// and a boolean to check if the value has been set.
func (o *AggTrade) GetAOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.A, true
}

// SetA sets field value
func (o *AggTrade) SetA(v int64) {
	o.A = v
}

// GetP returns the P field value
func (o *AggTrade) GetP() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.P
}

// GetPOk returns a tuple with the P field value
// and a boolean to check if the value has been set.
func (o *AggTrade) GetPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.P, true
}

// SetP sets field value
func (o *AggTrade) SetP(v string) {
	o.P = v
}

// GetQ returns the Q field value
func (o *AggTrade) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *AggTrade) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value
func (o *AggTrade) SetQ(v string) {
	o.Q = v
}

// GetF returns the F field value
func (o *AggTrade) GetF() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.F
}

// GetFOk returns a tuple with the F field value
// and a boolean to check if the value has been set.
func (o *AggTrade) GetFOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.F, true
}

// SetF sets field value
func (o *AggTrade) SetF(v int64) {
	o.F = v
}

// GetL returns the L field value
func (o *AggTrade) GetL() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.L
}

// GetLOk returns a tuple with the L field value
// and a boolean to check if the value has been set.
func (o *AggTrade) GetLOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L, true
}

// SetL sets field value
func (o *AggTrade) SetL(v int64) {
	o.L = v
}

// GetT returns the T field value
func (o *AggTrade) GetT() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *AggTrade) GetTOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *AggTrade) SetT(v bool) {
	o.T = v
}

// GetM returns the M field value
func (o *AggTrade) GetM() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.M
}

// GetMOk returns a tuple with the M field value
// and a boolean to check if the value has been set.
func (o *AggTrade) GetMOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.M, true
}

// SetM sets field value
func (o *AggTrade) SetM(v bool) {
	o.M = v
}

// GetM returns the M field value
func (o *AggTrade) GetM() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.M
}

// GetMOk returns a tuple with the M field value
// and a boolean to check if the value has been set.
func (o *AggTrade) GetMOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.M, true
}

// SetM sets field value
func (o *AggTrade) SetM(v bool) {
	o.M = v
}

func (o AggTrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a"] = o.A
	}
	if true {
		toSerialize["p"] = o.P
	}
	if true {
		toSerialize["q"] = o.Q
	}
	if true {
		toSerialize["f"] = o.F
	}
	if true {
		toSerialize["l"] = o.L
	}
	if true {
		toSerialize["T"] = o.T
	}
	if true {
		toSerialize["m"] = o.M
	}
	if true {
		toSerialize["M"] = o.M
	}
	return json.Marshal(toSerialize)
}

type NullableAggTrade struct {
	value *AggTrade
	isSet bool
}

func (v NullableAggTrade) Get() *AggTrade {
	return v.value
}

func (v *NullableAggTrade) Set(val *AggTrade) {
	v.value = val
	v.isSet = true
}

func (v NullableAggTrade) IsSet() bool {
	return v.isSet
}

func (v *NullableAggTrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggTrade(val *AggTrade) *NullableAggTrade {
	return &NullableAggTrade{value: val, isSet: true}
}

func (v NullableAggTrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggTrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
