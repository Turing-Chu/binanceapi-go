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

// InlineResponse2001Filters struct for InlineResponse2001Filters
type InlineResponse2001Filters struct {
	FilterType string `json:"filterType"`
	MinPrice   string `json:"minPrice"`
	MaxPrice   string `json:"maxPrice"`
	TickSize   string `json:"tickSize"`
}

// NewInlineResponse2001Filters instantiates a new InlineResponse2001Filters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001Filters(filterType string, minPrice string, maxPrice string, tickSize string) *InlineResponse2001Filters {
	this := InlineResponse2001Filters{}
	this.FilterType = filterType
	this.MinPrice = minPrice
	this.MaxPrice = maxPrice
	this.TickSize = tickSize
	return &this
}

// NewInlineResponse2001FiltersWithDefaults instantiates a new InlineResponse2001Filters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001FiltersWithDefaults() *InlineResponse2001Filters {
	this := InlineResponse2001Filters{}
	return &this
}

// GetFilterType returns the FilterType field value
func (o *InlineResponse2001Filters) GetFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Filters) GetFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value
func (o *InlineResponse2001Filters) SetFilterType(v string) {
	o.FilterType = v
}

// GetMinPrice returns the MinPrice field value
func (o *InlineResponse2001Filters) GetMinPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinPrice
}

// GetMinPriceOk returns a tuple with the MinPrice field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Filters) GetMinPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinPrice, true
}

// SetMinPrice sets field value
func (o *InlineResponse2001Filters) SetMinPrice(v string) {
	o.MinPrice = v
}

// GetMaxPrice returns the MaxPrice field value
func (o *InlineResponse2001Filters) GetMaxPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxPrice
}

// GetMaxPriceOk returns a tuple with the MaxPrice field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Filters) GetMaxPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPrice, true
}

// SetMaxPrice sets field value
func (o *InlineResponse2001Filters) SetMaxPrice(v string) {
	o.MaxPrice = v
}

// GetTickSize returns the TickSize field value
func (o *InlineResponse2001Filters) GetTickSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TickSize
}

// GetTickSizeOk returns a tuple with the TickSize field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Filters) GetTickSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TickSize, true
}

// SetTickSize sets field value
func (o *InlineResponse2001Filters) SetTickSize(v string) {
	o.TickSize = v
}

func (o InlineResponse2001Filters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filterType"] = o.FilterType
	}
	if true {
		toSerialize["minPrice"] = o.MinPrice
	}
	if true {
		toSerialize["maxPrice"] = o.MaxPrice
	}
	if true {
		toSerialize["tickSize"] = o.TickSize
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001Filters struct {
	value *InlineResponse2001Filters
	isSet bool
}

func (v NullableInlineResponse2001Filters) Get() *InlineResponse2001Filters {
	return v.value
}

func (v *NullableInlineResponse2001Filters) Set(val *InlineResponse2001Filters) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001Filters) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001Filters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001Filters(val *InlineResponse2001Filters) *NullableInlineResponse2001Filters {
	return &NullableInlineResponse2001Filters{value: val, isSet: true}
}

func (v NullableInlineResponse2001Filters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001Filters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
