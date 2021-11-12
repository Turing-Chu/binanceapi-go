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

// InlineResponse20060Assets struct for InlineResponse20060Assets
type InlineResponse20060Assets struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintenanceMargin      string `json:"maintenanceMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}

// NewInlineResponse20060Assets instantiates a new InlineResponse20060Assets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20060Assets(asset string, initialMargin string, maintenanceMargin string, marginBalance string, maxWithdrawAmount string, openOrderInitialMargin string, positionInitialMargin string, unrealizedProfit string, walletBalance string) *InlineResponse20060Assets {
	this := InlineResponse20060Assets{}
	this.Asset = asset
	this.InitialMargin = initialMargin
	this.MaintenanceMargin = maintenanceMargin
	this.MarginBalance = marginBalance
	this.MaxWithdrawAmount = maxWithdrawAmount
	this.OpenOrderInitialMargin = openOrderInitialMargin
	this.PositionInitialMargin = positionInitialMargin
	this.UnrealizedProfit = unrealizedProfit
	this.WalletBalance = walletBalance
	return &this
}

// NewInlineResponse20060AssetsWithDefaults instantiates a new InlineResponse20060Assets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20060AssetsWithDefaults() *InlineResponse20060Assets {
	this := InlineResponse20060Assets{}
	return &this
}

// GetAsset returns the Asset field value
func (o *InlineResponse20060Assets) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assets) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *InlineResponse20060Assets) SetAsset(v string) {
	o.Asset = v
}

// GetInitialMargin returns the InitialMargin field value
func (o *InlineResponse20060Assets) GetInitialMargin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assets) GetInitialMarginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialMargin, true
}

// SetInitialMargin sets field value
func (o *InlineResponse20060Assets) SetInitialMargin(v string) {
	o.InitialMargin = v
}

// GetMaintenanceMargin returns the MaintenanceMargin field value
func (o *InlineResponse20060Assets) GetMaintenanceMargin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaintenanceMargin
}

// GetMaintenanceMarginOk returns a tuple with the MaintenanceMargin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assets) GetMaintenanceMarginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaintenanceMargin, true
}

// SetMaintenanceMargin sets field value
func (o *InlineResponse20060Assets) SetMaintenanceMargin(v string) {
	o.MaintenanceMargin = v
}

// GetMarginBalance returns the MarginBalance field value
func (o *InlineResponse20060Assets) GetMarginBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assets) GetMarginBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarginBalance, true
}

// SetMarginBalance sets field value
func (o *InlineResponse20060Assets) SetMarginBalance(v string) {
	o.MarginBalance = v
}

// GetMaxWithdrawAmount returns the MaxWithdrawAmount field value
func (o *InlineResponse20060Assets) GetMaxWithdrawAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxWithdrawAmount
}

// GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assets) GetMaxWithdrawAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxWithdrawAmount, true
}

// SetMaxWithdrawAmount sets field value
func (o *InlineResponse20060Assets) SetMaxWithdrawAmount(v string) {
	o.MaxWithdrawAmount = v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value
func (o *InlineResponse20060Assets) GetOpenOrderInitialMargin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assets) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenOrderInitialMargin, true
}

// SetOpenOrderInitialMargin sets field value
func (o *InlineResponse20060Assets) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value
func (o *InlineResponse20060Assets) GetPositionInitialMargin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assets) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PositionInitialMargin, true
}

// SetPositionInitialMargin sets field value
func (o *InlineResponse20060Assets) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value
func (o *InlineResponse20060Assets) GetUnrealizedProfit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assets) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnrealizedProfit, true
}

// SetUnrealizedProfit sets field value
func (o *InlineResponse20060Assets) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = v
}

// GetWalletBalance returns the WalletBalance field value
func (o *InlineResponse20060Assets) GetWalletBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletBalance
}

// GetWalletBalanceOk returns a tuple with the WalletBalance field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20060Assets) GetWalletBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletBalance, true
}

// SetWalletBalance sets field value
func (o *InlineResponse20060Assets) SetWalletBalance(v string) {
	o.WalletBalance = v
}

func (o InlineResponse20060Assets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if true {
		toSerialize["maintenanceMargin"] = o.MaintenanceMargin
	}
	if true {
		toSerialize["marginBalance"] = o.MarginBalance
	}
	if true {
		toSerialize["maxWithdrawAmount"] = o.MaxWithdrawAmount
	}
	if true {
		toSerialize["openOrderInitialMargin"] = o.OpenOrderInitialMargin
	}
	if true {
		toSerialize["positionInitialMargin"] = o.PositionInitialMargin
	}
	if true {
		toSerialize["unrealizedProfit"] = o.UnrealizedProfit
	}
	if true {
		toSerialize["walletBalance"] = o.WalletBalance
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20060Assets struct {
	value *InlineResponse20060Assets
	isSet bool
}

func (v NullableInlineResponse20060Assets) Get() *InlineResponse20060Assets {
	return v.value
}

func (v *NullableInlineResponse20060Assets) Set(val *InlineResponse20060Assets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20060Assets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20060Assets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20060Assets(val *InlineResponse20060Assets) *NullableInlineResponse20060Assets {
	return &NullableInlineResponse20060Assets{value: val, isSet: true}
}

func (v NullableInlineResponse20060Assets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20060Assets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
