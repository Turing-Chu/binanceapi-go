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

// InlineResponse20030 struct for InlineResponse20030
type InlineResponse20030 struct {
	Coin              string                                 `json:"coin"`
	DepositAllEnable  bool                                   `json:"depositAllEnable"`
	Free              string                                 `json:"free"`
	Freeze            string                                 `json:"freeze"`
	Ipoable           string                                 `json:"ipoable"`
	Ipoing            string                                 `json:"ipoing"`
	IsLegalMoney      bool                                   `json:"isLegalMoney"`
	Locked            string                                 `json:"locked"`
	Name              string                                 `json:"name"`
	NetworkList       []SapiV1CapitalConfigGetallNetworkList `json:"networkList"`
	Storage           string                                 `json:"storage"`
	Trading           bool                                   `json:"trading"`
	WithdrawAllEnable bool                                   `json:"withdrawAllEnable"`
	Withdrawing       string                                 `json:"withdrawing"`
}

// NewInlineResponse20030 instantiates a new InlineResponse20030 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20030(coin string, depositAllEnable bool, free string, freeze string, ipoable string, ipoing string, isLegalMoney bool, locked string, name string, networkList []SapiV1CapitalConfigGetallNetworkList, storage string, trading bool, withdrawAllEnable bool, withdrawing string) *InlineResponse20030 {
	this := InlineResponse20030{}
	this.Coin = coin
	this.DepositAllEnable = depositAllEnable
	this.Free = free
	this.Freeze = freeze
	this.Ipoable = ipoable
	this.Ipoing = ipoing
	this.IsLegalMoney = isLegalMoney
	this.Locked = locked
	this.Name = name
	this.NetworkList = networkList
	this.Storage = storage
	this.Trading = trading
	this.WithdrawAllEnable = withdrawAllEnable
	this.Withdrawing = withdrawing
	return &this
}

// NewInlineResponse20030WithDefaults instantiates a new InlineResponse20030 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20030WithDefaults() *InlineResponse20030 {
	this := InlineResponse20030{}
	return &this
}

// GetCoin returns the Coin field value
func (o *InlineResponse20030) GetCoin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coin
}

// GetCoinOk returns a tuple with the Coin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetCoinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coin, true
}

// SetCoin sets field value
func (o *InlineResponse20030) SetCoin(v string) {
	o.Coin = v
}

// GetDepositAllEnable returns the DepositAllEnable field value
func (o *InlineResponse20030) GetDepositAllEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DepositAllEnable
}

// GetDepositAllEnableOk returns a tuple with the DepositAllEnable field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetDepositAllEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositAllEnable, true
}

// SetDepositAllEnable sets field value
func (o *InlineResponse20030) SetDepositAllEnable(v bool) {
	o.DepositAllEnable = v
}

// GetFree returns the Free field value
func (o *InlineResponse20030) GetFree() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetFreeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *InlineResponse20030) SetFree(v string) {
	o.Free = v
}

// GetFreeze returns the Freeze field value
func (o *InlineResponse20030) GetFreeze() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Freeze
}

// GetFreezeOk returns a tuple with the Freeze field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetFreezeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Freeze, true
}

// SetFreeze sets field value
func (o *InlineResponse20030) SetFreeze(v string) {
	o.Freeze = v
}

// GetIpoable returns the Ipoable field value
func (o *InlineResponse20030) GetIpoable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ipoable
}

// GetIpoableOk returns a tuple with the Ipoable field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetIpoableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ipoable, true
}

// SetIpoable sets field value
func (o *InlineResponse20030) SetIpoable(v string) {
	o.Ipoable = v
}

// GetIpoing returns the Ipoing field value
func (o *InlineResponse20030) GetIpoing() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ipoing
}

// GetIpoingOk returns a tuple with the Ipoing field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetIpoingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ipoing, true
}

// SetIpoing sets field value
func (o *InlineResponse20030) SetIpoing(v string) {
	o.Ipoing = v
}

// GetIsLegalMoney returns the IsLegalMoney field value
func (o *InlineResponse20030) GetIsLegalMoney() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLegalMoney
}

// GetIsLegalMoneyOk returns a tuple with the IsLegalMoney field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetIsLegalMoneyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLegalMoney, true
}

// SetIsLegalMoney sets field value
func (o *InlineResponse20030) SetIsLegalMoney(v bool) {
	o.IsLegalMoney = v
}

// GetLocked returns the Locked field value
func (o *InlineResponse20030) GetLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *InlineResponse20030) SetLocked(v string) {
	o.Locked = v
}

// GetName returns the Name field value
func (o *InlineResponse20030) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineResponse20030) SetName(v string) {
	o.Name = v
}

// GetNetworkList returns the NetworkList field value
func (o *InlineResponse20030) GetNetworkList() []SapiV1CapitalConfigGetallNetworkList {
	if o == nil {
		var ret []SapiV1CapitalConfigGetallNetworkList
		return ret
	}

	return o.NetworkList
}

// GetNetworkListOk returns a tuple with the NetworkList field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetNetworkListOk() (*[]SapiV1CapitalConfigGetallNetworkList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkList, true
}

// SetNetworkList sets field value
func (o *InlineResponse20030) SetNetworkList(v []SapiV1CapitalConfigGetallNetworkList) {
	o.NetworkList = v
}

// GetStorage returns the Storage field value
func (o *InlineResponse20030) GetStorage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetStorageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *InlineResponse20030) SetStorage(v string) {
	o.Storage = v
}

// GetTrading returns the Trading field value
func (o *InlineResponse20030) GetTrading() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Trading
}

// GetTradingOk returns a tuple with the Trading field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetTradingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trading, true
}

// SetTrading sets field value
func (o *InlineResponse20030) SetTrading(v bool) {
	o.Trading = v
}

// GetWithdrawAllEnable returns the WithdrawAllEnable field value
func (o *InlineResponse20030) GetWithdrawAllEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WithdrawAllEnable
}

// GetWithdrawAllEnableOk returns a tuple with the WithdrawAllEnable field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetWithdrawAllEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithdrawAllEnable, true
}

// SetWithdrawAllEnable sets field value
func (o *InlineResponse20030) SetWithdrawAllEnable(v bool) {
	o.WithdrawAllEnable = v
}

// GetWithdrawing returns the Withdrawing field value
func (o *InlineResponse20030) GetWithdrawing() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Withdrawing
}

// GetWithdrawingOk returns a tuple with the Withdrawing field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetWithdrawingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Withdrawing, true
}

// SetWithdrawing sets field value
func (o *InlineResponse20030) SetWithdrawing(v string) {
	o.Withdrawing = v
}

func (o InlineResponse20030) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["coin"] = o.Coin
	}
	if true {
		toSerialize["depositAllEnable"] = o.DepositAllEnable
	}
	if true {
		toSerialize["free"] = o.Free
	}
	if true {
		toSerialize["freeze"] = o.Freeze
	}
	if true {
		toSerialize["ipoable"] = o.Ipoable
	}
	if true {
		toSerialize["ipoing"] = o.Ipoing
	}
	if true {
		toSerialize["isLegalMoney"] = o.IsLegalMoney
	}
	if true {
		toSerialize["locked"] = o.Locked
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["networkList"] = o.NetworkList
	}
	if true {
		toSerialize["storage"] = o.Storage
	}
	if true {
		toSerialize["trading"] = o.Trading
	}
	if true {
		toSerialize["withdrawAllEnable"] = o.WithdrawAllEnable
	}
	if true {
		toSerialize["withdrawing"] = o.Withdrawing
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20030 struct {
	value *InlineResponse20030
	isSet bool
}

func (v NullableInlineResponse20030) Get() *InlineResponse20030 {
	return v.value
}

func (v *NullableInlineResponse20030) Set(val *InlineResponse20030) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20030) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20030) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20030(val *InlineResponse20030) *NullableInlineResponse20030 {
	return &NullableInlineResponse20030{value: val, isSet: true}
}

func (v NullableInlineResponse20030) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20030) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
