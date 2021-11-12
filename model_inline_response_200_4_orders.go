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

// InlineResponse2004Orders struct for InlineResponse2004Orders
type InlineResponse2004Orders struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

// NewInlineResponse2004Orders instantiates a new InlineResponse2004Orders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2004Orders(symbol string, orderId int64, clientOrderId string) *InlineResponse2004Orders {
	this := InlineResponse2004Orders{}
	this.Symbol = symbol
	this.OrderId = orderId
	this.ClientOrderId = clientOrderId
	return &this
}

// NewInlineResponse2004OrdersWithDefaults instantiates a new InlineResponse2004Orders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2004OrdersWithDefaults() *InlineResponse2004Orders {
	this := InlineResponse2004Orders{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *InlineResponse2004Orders) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004Orders) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *InlineResponse2004Orders) SetSymbol(v string) {
	o.Symbol = v
}

// GetOrderId returns the OrderId field value
func (o *InlineResponse2004Orders) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004Orders) GetOrderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *InlineResponse2004Orders) SetOrderId(v int64) {
	o.OrderId = v
}

// GetClientOrderId returns the ClientOrderId field value
func (o *InlineResponse2004Orders) GetClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2004Orders) GetClientOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientOrderId, true
}

// SetClientOrderId sets field value
func (o *InlineResponse2004Orders) SetClientOrderId(v string) {
	o.ClientOrderId = v
}

func (o InlineResponse2004Orders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["orderId"] = o.OrderId
	}
	if true {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2004Orders struct {
	value *InlineResponse2004Orders
	isSet bool
}

func (v NullableInlineResponse2004Orders) Get() *InlineResponse2004Orders {
	return v.value
}

func (v *NullableInlineResponse2004Orders) Set(val *InlineResponse2004Orders) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2004Orders) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2004Orders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2004Orders(val *InlineResponse2004Orders) *NullableInlineResponse2004Orders {
	return &NullableInlineResponse2004Orders{value: val, isSet: true}
}

func (v NullableInlineResponse2004Orders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2004Orders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
