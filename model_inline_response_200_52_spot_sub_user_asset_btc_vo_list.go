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

// InlineResponse20052SpotSubUserAssetBtcVoList struct for InlineResponse20052SpotSubUserAssetBtcVoList
type InlineResponse20052SpotSubUserAssetBtcVoList struct {
	Email      string `json:"email"`
	TotalAsset string `json:"totalAsset"`
}

// NewInlineResponse20052SpotSubUserAssetBtcVoList instantiates a new InlineResponse20052SpotSubUserAssetBtcVoList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20052SpotSubUserAssetBtcVoList(email string, totalAsset string) *InlineResponse20052SpotSubUserAssetBtcVoList {
	this := InlineResponse20052SpotSubUserAssetBtcVoList{}
	this.Email = email
	this.TotalAsset = totalAsset
	return &this
}

// NewInlineResponse20052SpotSubUserAssetBtcVoListWithDefaults instantiates a new InlineResponse20052SpotSubUserAssetBtcVoList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20052SpotSubUserAssetBtcVoListWithDefaults() *InlineResponse20052SpotSubUserAssetBtcVoList {
	this := InlineResponse20052SpotSubUserAssetBtcVoList{}
	return &this
}

// GetEmail returns the Email field value
func (o *InlineResponse20052SpotSubUserAssetBtcVoList) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20052SpotSubUserAssetBtcVoList) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *InlineResponse20052SpotSubUserAssetBtcVoList) SetEmail(v string) {
	o.Email = v
}

// GetTotalAsset returns the TotalAsset field value
func (o *InlineResponse20052SpotSubUserAssetBtcVoList) GetTotalAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TotalAsset
}

// GetTotalAssetOk returns a tuple with the TotalAsset field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20052SpotSubUserAssetBtcVoList) GetTotalAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAsset, true
}

// SetTotalAsset sets field value
func (o *InlineResponse20052SpotSubUserAssetBtcVoList) SetTotalAsset(v string) {
	o.TotalAsset = v
}

func (o InlineResponse20052SpotSubUserAssetBtcVoList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["totalAsset"] = o.TotalAsset
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20052SpotSubUserAssetBtcVoList struct {
	value *InlineResponse20052SpotSubUserAssetBtcVoList
	isSet bool
}

func (v NullableInlineResponse20052SpotSubUserAssetBtcVoList) Get() *InlineResponse20052SpotSubUserAssetBtcVoList {
	return v.value
}

func (v *NullableInlineResponse20052SpotSubUserAssetBtcVoList) Set(val *InlineResponse20052SpotSubUserAssetBtcVoList) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20052SpotSubUserAssetBtcVoList) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20052SpotSubUserAssetBtcVoList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20052SpotSubUserAssetBtcVoList(val *InlineResponse20052SpotSubUserAssetBtcVoList) *NullableInlineResponse20052SpotSubUserAssetBtcVoList {
	return &NullableInlineResponse20052SpotSubUserAssetBtcVoList{value: val, isSet: true}
}

func (v NullableInlineResponse20052SpotSubUserAssetBtcVoList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20052SpotSubUserAssetBtcVoList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
