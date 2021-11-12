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

// InlineResponse200103 struct for InlineResponse200103
type InlineResponse200103 struct {
	PoolId   int64    `json:"poolId"`
	PoolName string   `json:"poolName"`
	Assets   []string `json:"assets"`
}

// NewInlineResponse200103 instantiates a new InlineResponse200103 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200103(poolId int64, poolName string, assets []string) *InlineResponse200103 {
	this := InlineResponse200103{}
	this.PoolId = poolId
	this.PoolName = poolName
	this.Assets = assets
	return &this
}

// NewInlineResponse200103WithDefaults instantiates a new InlineResponse200103 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200103WithDefaults() *InlineResponse200103 {
	this := InlineResponse200103{}
	return &this
}

// GetPoolId returns the PoolId field value
func (o *InlineResponse200103) GetPoolId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetPoolIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *InlineResponse200103) SetPoolId(v int64) {
	o.PoolId = v
}

// GetPoolName returns the PoolName field value
func (o *InlineResponse200103) GetPoolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetPoolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolName, true
}

// SetPoolName sets field value
func (o *InlineResponse200103) SetPoolName(v string) {
	o.PoolName = v
}

// GetAssets returns the Assets field value
func (o *InlineResponse200103) GetAssets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200103) GetAssetsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assets, true
}

// SetAssets sets field value
func (o *InlineResponse200103) SetAssets(v []string) {
	o.Assets = v
}

func (o InlineResponse200103) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["poolId"] = o.PoolId
	}
	if true {
		toSerialize["poolName"] = o.PoolName
	}
	if true {
		toSerialize["assets"] = o.Assets
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200103 struct {
	value *InlineResponse200103
	isSet bool
}

func (v NullableInlineResponse200103) Get() *InlineResponse200103 {
	return v.value
}

func (v *NullableInlineResponse200103) Set(val *InlineResponse200103) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200103) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200103) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200103(val *InlineResponse200103) *NullableInlineResponse200103 {
	return &NullableInlineResponse200103{value: val, isSet: true}
}

func (v NullableInlineResponse200103) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200103) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
