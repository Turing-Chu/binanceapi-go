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

// BswapLiquidityLiquidity struct for BswapLiquidityLiquidity
type BswapLiquidityLiquidity struct {
	BUSD float64 `json:"BUSD"`
	USDT float64 `json:"USDT"`
}

// NewBswapLiquidityLiquidity instantiates a new BswapLiquidityLiquidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBswapLiquidityLiquidity(bUSD float64, uSDT float64) *BswapLiquidityLiquidity {
	this := BswapLiquidityLiquidity{}
	this.BUSD = bUSD
	this.USDT = uSDT
	return &this
}

// NewBswapLiquidityLiquidityWithDefaults instantiates a new BswapLiquidityLiquidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBswapLiquidityLiquidityWithDefaults() *BswapLiquidityLiquidity {
	this := BswapLiquidityLiquidity{}
	return &this
}

// GetBUSD returns the BUSD field value
func (o *BswapLiquidityLiquidity) GetBUSD() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.BUSD
}

// GetBUSDOk returns a tuple with the BUSD field value
// and a boolean to check if the value has been set.
func (o *BswapLiquidityLiquidity) GetBUSDOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BUSD, true
}

// SetBUSD sets field value
func (o *BswapLiquidityLiquidity) SetBUSD(v float64) {
	o.BUSD = v
}

// GetUSDT returns the USDT field value
func (o *BswapLiquidityLiquidity) GetUSDT() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.USDT
}

// GetUSDTOk returns a tuple with the USDT field value
// and a boolean to check if the value has been set.
func (o *BswapLiquidityLiquidity) GetUSDTOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.USDT, true
}

// SetUSDT sets field value
func (o *BswapLiquidityLiquidity) SetUSDT(v float64) {
	o.USDT = v
}

func (o BswapLiquidityLiquidity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["BUSD"] = o.BUSD
	}
	if true {
		toSerialize["USDT"] = o.USDT
	}
	return json.Marshal(toSerialize)
}

type NullableBswapLiquidityLiquidity struct {
	value *BswapLiquidityLiquidity
	isSet bool
}

func (v NullableBswapLiquidityLiquidity) Get() *BswapLiquidityLiquidity {
	return v.value
}

func (v *NullableBswapLiquidityLiquidity) Set(val *BswapLiquidityLiquidity) {
	v.value = val
	v.isSet = true
}

func (v NullableBswapLiquidityLiquidity) IsSet() bool {
	return v.isSet
}

func (v *NullableBswapLiquidityLiquidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBswapLiquidityLiquidity(val *BswapLiquidityLiquidity) *NullableBswapLiquidityLiquidity {
	return &NullableBswapLiquidityLiquidity{value: val, isSet: true}
}

func (v NullableBswapLiquidityLiquidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBswapLiquidityLiquidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}