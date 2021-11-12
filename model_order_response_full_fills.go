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

// OrderResponseFullFills struct for OrderResponseFullFills
type OrderResponseFullFills struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
}

// NewOrderResponseFullFills instantiates a new OrderResponseFullFills object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseFullFills(price string, qty string, commission string, commissionAsset string) *OrderResponseFullFills {
	this := OrderResponseFullFills{}
	this.Price = price
	this.Qty = qty
	this.Commission = commission
	this.CommissionAsset = commissionAsset
	return &this
}

// NewOrderResponseFullFillsWithDefaults instantiates a new OrderResponseFullFills object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseFullFillsWithDefaults() *OrderResponseFullFills {
	this := OrderResponseFullFills{}
	return &this
}

// GetPrice returns the Price field value
func (o *OrderResponseFullFills) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *OrderResponseFullFills) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *OrderResponseFullFills) SetPrice(v string) {
	o.Price = v
}

// GetQty returns the Qty field value
func (o *OrderResponseFullFills) GetQty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Qty
}

// GetQtyOk returns a tuple with the Qty field value
// and a boolean to check if the value has been set.
func (o *OrderResponseFullFills) GetQtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qty, true
}

// SetQty sets field value
func (o *OrderResponseFullFills) SetQty(v string) {
	o.Qty = v
}

// GetCommission returns the Commission field value
func (o *OrderResponseFullFills) GetCommission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value
// and a boolean to check if the value has been set.
func (o *OrderResponseFullFills) GetCommissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commission, true
}

// SetCommission sets field value
func (o *OrderResponseFullFills) SetCommission(v string) {
	o.Commission = v
}

// GetCommissionAsset returns the CommissionAsset field value
func (o *OrderResponseFullFills) GetCommissionAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value
// and a boolean to check if the value has been set.
func (o *OrderResponseFullFills) GetCommissionAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommissionAsset, true
}

// SetCommissionAsset sets field value
func (o *OrderResponseFullFills) SetCommissionAsset(v string) {
	o.CommissionAsset = v
}

func (o OrderResponseFullFills) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["qty"] = o.Qty
	}
	if true {
		toSerialize["commission"] = o.Commission
	}
	if true {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}
	return json.Marshal(toSerialize)
}

type NullableOrderResponseFullFills struct {
	value *OrderResponseFullFills
	isSet bool
}

func (v NullableOrderResponseFullFills) Get() *OrderResponseFullFills {
	return v.value
}

func (v *NullableOrderResponseFullFills) Set(val *OrderResponseFullFills) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseFullFills) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseFullFills) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseFullFills(val *OrderResponseFullFills) *NullableOrderResponseFullFills {
	return &NullableOrderResponseFullFills{value: val, isSet: true}
}

func (v NullableOrderResponseFullFills) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseFullFills) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
