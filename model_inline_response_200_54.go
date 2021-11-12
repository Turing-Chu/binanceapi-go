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

// InlineResponse20054 struct for InlineResponse20054
type InlineResponse20054 struct {
	Amount       string `json:"amount"`
	Coin         string `json:"coin"`
	Network      string `json:"network"`
	Status       int32  `json:"status"`
	Address      string `json:"address"`
	AddressTag   string `json:"addressTag"`
	TxId         string `json:"txId"`
	InsertTime   int64  `json:"insertTime"`
	TransferType int32  `json:"transferType"`
	ConfirmTimes string `json:"confirmTimes"`
}

// NewInlineResponse20054 instantiates a new InlineResponse20054 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20054(amount string, coin string, network string, status int32, address string, addressTag string, txId string, insertTime int64, transferType int32, confirmTimes string) *InlineResponse20054 {
	this := InlineResponse20054{}
	this.Amount = amount
	this.Coin = coin
	this.Network = network
	this.Status = status
	this.Address = address
	this.AddressTag = addressTag
	this.TxId = txId
	this.InsertTime = insertTime
	this.TransferType = transferType
	this.ConfirmTimes = confirmTimes
	return &this
}

// NewInlineResponse20054WithDefaults instantiates a new InlineResponse20054 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20054WithDefaults() *InlineResponse20054 {
	this := InlineResponse20054{}
	return &this
}

// GetAmount returns the Amount field value
func (o *InlineResponse20054) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InlineResponse20054) SetAmount(v string) {
	o.Amount = v
}

// GetCoin returns the Coin field value
func (o *InlineResponse20054) GetCoin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coin
}

// GetCoinOk returns a tuple with the Coin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetCoinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coin, true
}

// SetCoin sets field value
func (o *InlineResponse20054) SetCoin(v string) {
	o.Coin = v
}

// GetNetwork returns the Network field value
func (o *InlineResponse20054) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *InlineResponse20054) SetNetwork(v string) {
	o.Network = v
}

// GetStatus returns the Status field value
func (o *InlineResponse20054) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20054) SetStatus(v int32) {
	o.Status = v
}

// GetAddress returns the Address field value
func (o *InlineResponse20054) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *InlineResponse20054) SetAddress(v string) {
	o.Address = v
}

// GetAddressTag returns the AddressTag field value
func (o *InlineResponse20054) GetAddressTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressTag
}

// GetAddressTagOk returns a tuple with the AddressTag field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetAddressTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressTag, true
}

// SetAddressTag sets field value
func (o *InlineResponse20054) SetAddressTag(v string) {
	o.AddressTag = v
}

// GetTxId returns the TxId field value
func (o *InlineResponse20054) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *InlineResponse20054) SetTxId(v string) {
	o.TxId = v
}

// GetInsertTime returns the InsertTime field value
func (o *InlineResponse20054) GetInsertTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InsertTime
}

// GetInsertTimeOk returns a tuple with the InsertTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetInsertTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InsertTime, true
}

// SetInsertTime sets field value
func (o *InlineResponse20054) SetInsertTime(v int64) {
	o.InsertTime = v
}

// GetTransferType returns the TransferType field value
func (o *InlineResponse20054) GetTransferType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetTransferTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferType, true
}

// SetTransferType sets field value
func (o *InlineResponse20054) SetTransferType(v int32) {
	o.TransferType = v
}

// GetConfirmTimes returns the ConfirmTimes field value
func (o *InlineResponse20054) GetConfirmTimes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfirmTimes
}

// GetConfirmTimesOk returns a tuple with the ConfirmTimes field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetConfirmTimesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmTimes, true
}

// SetConfirmTimes sets field value
func (o *InlineResponse20054) SetConfirmTimes(v string) {
	o.ConfirmTimes = v
}

func (o InlineResponse20054) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["coin"] = o.Coin
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["addressTag"] = o.AddressTag
	}
	if true {
		toSerialize["txId"] = o.TxId
	}
	if true {
		toSerialize["insertTime"] = o.InsertTime
	}
	if true {
		toSerialize["transferType"] = o.TransferType
	}
	if true {
		toSerialize["confirmTimes"] = o.ConfirmTimes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20054 struct {
	value *InlineResponse20054
	isSet bool
}

func (v NullableInlineResponse20054) Get() *InlineResponse20054 {
	return v.value
}

func (v *NullableInlineResponse20054) Set(val *InlineResponse20054) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20054) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20054) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20054(val *InlineResponse20054) *NullableInlineResponse20054 {
	return &NullableInlineResponse20054{value: val, isSet: true}
}

func (v NullableInlineResponse20054) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20054) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
