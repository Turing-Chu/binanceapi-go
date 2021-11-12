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

// InlineResponse20058SubAccountList struct for InlineResponse20058SubAccountList
type InlineResponse20058SubAccountList struct {
	Email               string `json:"email"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`
}

// NewInlineResponse20058SubAccountList instantiates a new InlineResponse20058SubAccountList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20058SubAccountList(email string, totalAssetOfBtc string, totalLiabilityOfBtc string, totalNetAssetOfBtc string) *InlineResponse20058SubAccountList {
	this := InlineResponse20058SubAccountList{}
	this.Email = email
	this.TotalAssetOfBtc = totalAssetOfBtc
	this.TotalLiabilityOfBtc = totalLiabilityOfBtc
	this.TotalNetAssetOfBtc = totalNetAssetOfBtc
	return &this
}

// NewInlineResponse20058SubAccountListWithDefaults instantiates a new InlineResponse20058SubAccountList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20058SubAccountListWithDefaults() *InlineResponse20058SubAccountList {
	this := InlineResponse20058SubAccountList{}
	return &this
}

// GetEmail returns the Email field value
func (o *InlineResponse20058SubAccountList) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20058SubAccountList) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InlineResponse20058SubAccountList) SetEmail(v string) {
	o.Email = v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value
func (o *InlineResponse20058SubAccountList) GetTotalAssetOfBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20058SubAccountList) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAssetOfBtc, true
}

// SetTotalAssetOfBtc sets field value
func (o *InlineResponse20058SubAccountList) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value
func (o *InlineResponse20058SubAccountList) GetTotalLiabilityOfBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20058SubAccountList) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalLiabilityOfBtc, true
}

// SetTotalLiabilityOfBtc sets field value
func (o *InlineResponse20058SubAccountList) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value
func (o *InlineResponse20058SubAccountList) GetTotalNetAssetOfBtc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20058SubAccountList) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNetAssetOfBtc, true
}

// SetTotalNetAssetOfBtc sets field value
func (o *InlineResponse20058SubAccountList) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = v
}

func (o InlineResponse20058SubAccountList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	if true {
		toSerialize["totalLiabilityOfBtc"] = o.TotalLiabilityOfBtc
	}
	if true {
		toSerialize["totalNetAssetOfBtc"] = o.TotalNetAssetOfBtc
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20058SubAccountList struct {
	value *InlineResponse20058SubAccountList
	isSet bool
}

func (v NullableInlineResponse20058SubAccountList) Get() *InlineResponse20058SubAccountList {
	return v.value
}

func (v *NullableInlineResponse20058SubAccountList) Set(val *InlineResponse20058SubAccountList) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20058SubAccountList) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20058SubAccountList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20058SubAccountList(val *InlineResponse20058SubAccountList) *NullableInlineResponse20058SubAccountList {
	return &NullableInlineResponse20058SubAccountList{value: val, isSet: true}
}

func (v NullableInlineResponse20058SubAccountList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20058SubAccountList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
