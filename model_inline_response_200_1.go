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

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	Timezone        string                         `json:"timezone"`
	ServerTime      int64                          `json:"serverTime"`
	RateLimits      []InlineResponse2001RateLimits `json:"rateLimits"`
	ExchangeFilters []map[string]interface{}       `json:"exchangeFilters"`
	Symbols         []InlineResponse2001Symbols    `json:"symbols"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001(timezone string, serverTime int64, rateLimits []InlineResponse2001RateLimits, exchangeFilters []map[string]interface{}, symbols []InlineResponse2001Symbols) *InlineResponse2001 {
	this := InlineResponse2001{}
	this.Timezone = timezone
	this.ServerTime = serverTime
	this.RateLimits = rateLimits
	this.ExchangeFilters = exchangeFilters
	this.Symbols = symbols
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetTimezone returns the Timezone field value
func (o *InlineResponse2001) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *InlineResponse2001) SetTimezone(v string) {
	o.Timezone = v
}

// GetServerTime returns the ServerTime field value
func (o *InlineResponse2001) GetServerTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerTime
}

// GetServerTimeOk returns a tuple with the ServerTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetServerTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerTime, true
}

// SetServerTime sets field value
func (o *InlineResponse2001) SetServerTime(v int64) {
	o.ServerTime = v
}

// GetRateLimits returns the RateLimits field value
func (o *InlineResponse2001) GetRateLimits() []InlineResponse2001RateLimits {
	if o == nil {
		var ret []InlineResponse2001RateLimits
		return ret
	}

	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetRateLimitsOk() (*[]InlineResponse2001RateLimits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RateLimits, true
}

// SetRateLimits sets field value
func (o *InlineResponse2001) SetRateLimits(v []InlineResponse2001RateLimits) {
	o.RateLimits = v
}

// GetExchangeFilters returns the ExchangeFilters field value
func (o *InlineResponse2001) GetExchangeFilters() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.ExchangeFilters
}

// GetExchangeFiltersOk returns a tuple with the ExchangeFilters field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetExchangeFiltersOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeFilters, true
}

// SetExchangeFilters sets field value
func (o *InlineResponse2001) SetExchangeFilters(v []map[string]interface{}) {
	o.ExchangeFilters = v
}

// GetSymbols returns the Symbols field value
func (o *InlineResponse2001) GetSymbols() []InlineResponse2001Symbols {
	if o == nil {
		var ret []InlineResponse2001Symbols
		return ret
	}

	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetSymbolsOk() (*[]InlineResponse2001Symbols, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbols, true
}

// SetSymbols sets field value
func (o *InlineResponse2001) SetSymbols(v []InlineResponse2001Symbols) {
	o.Symbols = v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	if true {
		toSerialize["serverTime"] = o.ServerTime
	}
	if true {
		toSerialize["rateLimits"] = o.RateLimits
	}
	if true {
		toSerialize["exchangeFilters"] = o.ExchangeFilters
	}
	if true {
		toSerialize["symbols"] = o.Symbols
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
