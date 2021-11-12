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

// SapiV1BswapLiquidityLiquidity struct for SapiV1BswapLiquidityLiquidity
type SapiV1BswapLiquidityLiquidity struct {
	BUSD float64 `json:"BUSD"`
	USDT float64 `json:"USDT"`
}

// NewSapiV1BswapLiquidityLiquidity instantiates a new SapiV1BswapLiquidityLiquidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSapiV1BswapLiquidityLiquidity(bUSD float64, uSDT float64) *SapiV1BswapLiquidityLiquidity {
	this := SapiV1BswapLiquidityLiquidity{}
	this.BUSD = bUSD
	this.USDT = uSDT
	return &this
}

// NewSapiV1BswapLiquidityLiquidityWithDefaults instantiates a new SapiV1BswapLiquidityLiquidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSapiV1BswapLiquidityLiquidityWithDefaults() *SapiV1BswapLiquidityLiquidity {
	this := SapiV1BswapLiquidityLiquidity{}
	return &this
}

// GetBUSD returns the BUSD field value
func (o *SapiV1BswapLiquidityLiquidity) GetBUSD() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.BUSD
}

// GetBUSDOk returns a tuple with the BUSD field value
// and a boolean to check if the value has been set.
func (o *SapiV1BswapLiquidityLiquidity) GetBUSDOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BUSD, true
}

// SetBUSD sets field value
func (o *SapiV1BswapLiquidityLiquidity) SetBUSD(v float64) {
	o.BUSD = v
}

// GetUSDT returns the USDT field value
func (o *SapiV1BswapLiquidityLiquidity) GetUSDT() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.USDT
}

// GetUSDTOk returns a tuple with the USDT field value
// and a boolean to check if the value has been set.
func (o *SapiV1BswapLiquidityLiquidity) GetUSDTOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.USDT, true
}

// SetUSDT sets field value
func (o *SapiV1BswapLiquidityLiquidity) SetUSDT(v float64) {
	o.USDT = v
}

func (o SapiV1BswapLiquidityLiquidity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["BUSD"] = o.BUSD
	}
	if true {
		toSerialize["USDT"] = o.USDT
	}
	return json.Marshal(toSerialize)
}

type NullableSapiV1BswapLiquidityLiquidity struct {
	value *SapiV1BswapLiquidityLiquidity
	isSet bool
}

func (v NullableSapiV1BswapLiquidityLiquidity) Get() *SapiV1BswapLiquidityLiquidity {
	return v.value
}

func (v *NullableSapiV1BswapLiquidityLiquidity) Set(val *SapiV1BswapLiquidityLiquidity) {
	v.value = val
	v.isSet = true
}

func (v NullableSapiV1BswapLiquidityLiquidity) IsSet() bool {
	return v.isSet
}

func (v *NullableSapiV1BswapLiquidityLiquidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSapiV1BswapLiquidityLiquidity(val *SapiV1BswapLiquidityLiquidity) *NullableSapiV1BswapLiquidityLiquidity {
	return &NullableSapiV1BswapLiquidityLiquidity{value: val, isSet: true}
}

func (v NullableSapiV1BswapLiquidityLiquidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSapiV1BswapLiquidityLiquidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
