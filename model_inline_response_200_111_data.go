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

// InlineResponse200111Data struct for InlineResponse200111Data
type InlineResponse200111Data struct {
	OrderNumber string `json:"orderNumber"`
	AdvNo       string `json:"advNo"`
	TradeType   string `json:"tradeType"`
	Asset       string `json:"asset"`
	Fiat        string `json:"fiat"`
	FiatSymbol  string `json:"fiatSymbol"`
	// Quantity (in Crypto)
	Amount     string `json:"amount"`
	TotalPrice string `json:"totalPrice"`
	// Unit Price (in Fiat)
	UnitPrice string `json:"unitPrice"`
	// PENDING, TRADING, BUYER_PAYED, DISTRIBUTING, COMPLETED, IN_APPEAL, CANCELLED, CANCELLED_BY_SYSTEM
	OrderStatus string `json:"orderStatus"`
	CreateTime  int64  `json:"createTime"`
	// Transaction Fee (in Crypto)
	Commission          string `json:"commission"`
	CounterPartNickName string `json:"counterPartNickName"`
	AdvertisementRole   string `json:"advertisementRole"`
}

// NewInlineResponse200111Data instantiates a new InlineResponse200111Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200111Data(orderNumber string, advNo string, tradeType string, asset string, fiat string, fiatSymbol string, amount string, totalPrice string, unitPrice string, orderStatus string, createTime int64, commission string, counterPartNickName string, advertisementRole string) *InlineResponse200111Data {
	this := InlineResponse200111Data{}
	this.OrderNumber = orderNumber
	this.AdvNo = advNo
	this.TradeType = tradeType
	this.Asset = asset
	this.Fiat = fiat
	this.FiatSymbol = fiatSymbol
	this.Amount = amount
	this.TotalPrice = totalPrice
	this.UnitPrice = unitPrice
	this.OrderStatus = orderStatus
	this.CreateTime = createTime
	this.Commission = commission
	this.CounterPartNickName = counterPartNickName
	this.AdvertisementRole = advertisementRole
	return &this
}

// NewInlineResponse200111DataWithDefaults instantiates a new InlineResponse200111Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200111DataWithDefaults() *InlineResponse200111Data {
	this := InlineResponse200111Data{}
	return &this
}

// GetOrderNumber returns the OrderNumber field value
func (o *InlineResponse200111Data) GetOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderNumber, true
}

// SetOrderNumber sets field value
func (o *InlineResponse200111Data) SetOrderNumber(v string) {
	o.OrderNumber = v
}

// GetAdvNo returns the AdvNo field value
func (o *InlineResponse200111Data) GetAdvNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdvNo
}

// GetAdvNoOk returns a tuple with the AdvNo field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetAdvNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdvNo, true
}

// SetAdvNo sets field value
func (o *InlineResponse200111Data) SetAdvNo(v string) {
	o.AdvNo = v
}

// GetTradeType returns the TradeType field value
func (o *InlineResponse200111Data) GetTradeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetTradeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeType, true
}

// SetTradeType sets field value
func (o *InlineResponse200111Data) SetTradeType(v string) {
	o.TradeType = v
}

// GetAsset returns the Asset field value
func (o *InlineResponse200111Data) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse200111Data) SetAsset(v string) {
	o.Asset = v
}

// GetFiat returns the Fiat field value
func (o *InlineResponse200111Data) GetFiat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fiat
}

// GetFiatOk returns a tuple with the Fiat field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetFiatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fiat, true
}

// SetFiat sets field value
func (o *InlineResponse200111Data) SetFiat(v string) {
	o.Fiat = v
}

// GetFiatSymbol returns the FiatSymbol field value
func (o *InlineResponse200111Data) GetFiatSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FiatSymbol
}

// GetFiatSymbolOk returns a tuple with the FiatSymbol field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetFiatSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiatSymbol, true
}

// SetFiatSymbol sets field value
func (o *InlineResponse200111Data) SetFiatSymbol(v string) {
	o.FiatSymbol = v
}

// GetAmount returns the Amount field value
func (o *InlineResponse200111Data) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InlineResponse200111Data) SetAmount(v string) {
	o.Amount = v
}

// GetTotalPrice returns the TotalPrice field value
func (o *InlineResponse200111Data) GetTotalPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetTotalPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPrice, true
}

// SetTotalPrice sets field value
func (o *InlineResponse200111Data) SetTotalPrice(v string) {
	o.TotalPrice = v
}

// GetUnitPrice returns the UnitPrice field value
func (o *InlineResponse200111Data) GetUnitPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetUnitPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *InlineResponse200111Data) SetUnitPrice(v string) {
	o.UnitPrice = v
}

// GetOrderStatus returns the OrderStatus field value
func (o *InlineResponse200111Data) GetOrderStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetOrderStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderStatus, true
}

// SetOrderStatus sets field value
func (o *InlineResponse200111Data) SetOrderStatus(v string) {
	o.OrderStatus = v
}

// GetCreateTime returns the CreateTime field value
func (o *InlineResponse200111Data) GetCreateTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetCreateTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateTime, true
}

// SetCreateTime sets field value
func (o *InlineResponse200111Data) SetCreateTime(v int64) {
	o.CreateTime = v
}

// GetCommission returns the Commission field value
func (o *InlineResponse200111Data) GetCommission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetCommissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commission, true
}

// SetCommission sets field value
func (o *InlineResponse200111Data) SetCommission(v string) {
	o.Commission = v
}

// GetCounterPartNickName returns the CounterPartNickName field value
func (o *InlineResponse200111Data) GetCounterPartNickName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CounterPartNickName
}

// GetCounterPartNickNameOk returns a tuple with the CounterPartNickName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetCounterPartNickNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CounterPartNickName, true
}

// SetCounterPartNickName sets field value
func (o *InlineResponse200111Data) SetCounterPartNickName(v string) {
	o.CounterPartNickName = v
}

// GetAdvertisementRole returns the AdvertisementRole field value
func (o *InlineResponse200111Data) GetAdvertisementRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdvertisementRole
}

// GetAdvertisementRoleOk returns a tuple with the AdvertisementRole field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200111Data) GetAdvertisementRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdvertisementRole, true
}

// SetAdvertisementRole sets field value
func (o *InlineResponse200111Data) SetAdvertisementRole(v string) {
	o.AdvertisementRole = v
}

func (o InlineResponse200111Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if true {
		toSerialize["advNo"] = o.AdvNo
	}
	if true {
		toSerialize["tradeType"] = o.TradeType
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["fiat"] = o.Fiat
	}
	if true {
		toSerialize["fiatSymbol"] = o.FiatSymbol
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["totalPrice"] = o.TotalPrice
	}
	if true {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	if true {
		toSerialize["orderStatus"] = o.OrderStatus
	}
	if true {
		toSerialize["createTime"] = o.CreateTime
	}
	if true {
		toSerialize["commission"] = o.Commission
	}
	if true {
		toSerialize["counterPartNickName"] = o.CounterPartNickName
	}
	if true {
		toSerialize["advertisementRole"] = o.AdvertisementRole
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200111Data struct {
	value *InlineResponse200111Data
	isSet bool
}

func (v NullableInlineResponse200111Data) Get() *InlineResponse200111Data {
	return v.value
}

func (v *NullableInlineResponse200111Data) Set(val *InlineResponse200111Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200111Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200111Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200111Data(val *InlineResponse200111Data) *NullableInlineResponse200111Data {
	return &NullableInlineResponse200111Data{value: val, isSet: true}
}

func (v NullableInlineResponse200111Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200111Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
