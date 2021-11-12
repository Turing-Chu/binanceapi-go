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

// InlineResponse20033 struct for InlineResponse20033
type InlineResponse20033 struct {
	Address   string `json:"address"`
	Amount    string `json:"amount"`
	ApplyTime string `json:"applyTime"`
	Coin      string `json:"coin"`
	Id        string `json:"id"`
	// will not be returned if there's no withdrawOrderId for this withdraw.
	WithdrawOrderId string `json:"withdrawOrderId"`
	Network         string `json:"network"`
	// 1 for internal transfer, 0 for external transfer
	TransferType   int32  `json:"transferType"`
	Status         int32  `json:"status"`
	TransactionFee string `json:"transactionFee"`
	ConfirmNo      int32  `json:"confirmNo"`
	TxId           string `json:"txId"`
}

// NewInlineResponse20033 instantiates a new InlineResponse20033 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20033(address string, amount string, applyTime string, coin string, id string, withdrawOrderId string, network string, transferType int32, status int32, transactionFee string, confirmNo int32, txId string) *InlineResponse20033 {
	this := InlineResponse20033{}
	this.Address = address
	this.Amount = amount
	this.ApplyTime = applyTime
	this.Coin = coin
	this.Id = id
	this.WithdrawOrderId = withdrawOrderId
	this.Network = network
	this.TransferType = transferType
	this.Status = status
	this.TransactionFee = transactionFee
	this.ConfirmNo = confirmNo
	this.TxId = txId
	return &this
}

// NewInlineResponse20033WithDefaults instantiates a new InlineResponse20033 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20033WithDefaults() *InlineResponse20033 {
	this := InlineResponse20033{}
	return &this
}

// GetAddress returns the Address field value
func (o *InlineResponse20033) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *InlineResponse20033) SetAddress(v string) {
	o.Address = v
}

// GetAmount returns the Amount field value
func (o *InlineResponse20033) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InlineResponse20033) SetAmount(v string) {
	o.Amount = v
}

// GetApplyTime returns the ApplyTime field value
func (o *InlineResponse20033) GetApplyTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplyTime
}

// GetApplyTimeOk returns a tuple with the ApplyTime field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetApplyTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplyTime, true
}

// SetApplyTime sets field value
func (o *InlineResponse20033) SetApplyTime(v string) {
	o.ApplyTime = v
}

// GetCoin returns the Coin field value
func (o *InlineResponse20033) GetCoin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coin
}

// GetCoinOk returns a tuple with the Coin field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetCoinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coin, true
}

// SetCoin sets field value
func (o *InlineResponse20033) SetCoin(v string) {
	o.Coin = v
}

// GetId returns the Id field value
func (o *InlineResponse20033) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse20033) SetId(v string) {
	o.Id = v
}

// GetWithdrawOrderId returns the WithdrawOrderId field value
func (o *InlineResponse20033) GetWithdrawOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WithdrawOrderId
}

// GetWithdrawOrderIdOk returns a tuple with the WithdrawOrderId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetWithdrawOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithdrawOrderId, true
}

// SetWithdrawOrderId sets field value
func (o *InlineResponse20033) SetWithdrawOrderId(v string) {
	o.WithdrawOrderId = v
}

// GetNetwork returns the Network field value
func (o *InlineResponse20033) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *InlineResponse20033) SetNetwork(v string) {
	o.Network = v
}

// GetTransferType returns the TransferType field value
func (o *InlineResponse20033) GetTransferType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetTransferTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferType, true
}

// SetTransferType sets field value
func (o *InlineResponse20033) SetTransferType(v int32) {
	o.TransferType = v
}

// GetStatus returns the Status field value
func (o *InlineResponse20033) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse20033) SetStatus(v int32) {
	o.Status = v
}

// GetTransactionFee returns the TransactionFee field value
func (o *InlineResponse20033) GetTransactionFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionFee
}

// GetTransactionFeeOk returns a tuple with the TransactionFee field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetTransactionFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionFee, true
}

// SetTransactionFee sets field value
func (o *InlineResponse20033) SetTransactionFee(v string) {
	o.TransactionFee = v
}

// GetConfirmNo returns the ConfirmNo field value
func (o *InlineResponse20033) GetConfirmNo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConfirmNo
}

// GetConfirmNoOk returns a tuple with the ConfirmNo field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetConfirmNoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmNo, true
}

// SetConfirmNo sets field value
func (o *InlineResponse20033) SetConfirmNo(v int32) {
	o.ConfirmNo = v
}

// GetTxId returns the TxId field value
func (o *InlineResponse20033) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20033) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *InlineResponse20033) SetTxId(v string) {
	o.TxId = v
}

func (o InlineResponse20033) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["applyTime"] = o.ApplyTime
	}
	if true {
		toSerialize["coin"] = o.Coin
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["withdrawOrderId"] = o.WithdrawOrderId
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["transferType"] = o.TransferType
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["transactionFee"] = o.TransactionFee
	}
	if true {
		toSerialize["confirmNo"] = o.ConfirmNo
	}
	if true {
		toSerialize["txId"] = o.TxId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20033 struct {
	value *InlineResponse20033
	isSet bool
}

func (v NullableInlineResponse20033) Get() *InlineResponse20033 {
	return v.value
}

func (v *NullableInlineResponse20033) Set(val *InlineResponse20033) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20033) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20033) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20033(val *InlineResponse20033) *NullableInlineResponse20033 {
	return &NullableInlineResponse20033{value: val, isSet: true}
}

func (v NullableInlineResponse20033) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20033) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
