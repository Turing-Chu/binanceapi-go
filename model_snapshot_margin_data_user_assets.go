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

// SnapshotMarginDataUserAssets struct for SnapshotMarginDataUserAssets
type SnapshotMarginDataUserAssets struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// NewSnapshotMarginDataUserAssets instantiates a new SnapshotMarginDataUserAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotMarginDataUserAssets(asset string, borrowed string, free string, interest string, locked string, netAsset string) *SnapshotMarginDataUserAssets {
	this := SnapshotMarginDataUserAssets{}
	this.Asset = asset
	this.Borrowed = borrowed
	this.Free = free
	this.Interest = interest
	this.Locked = locked
	this.NetAsset = netAsset
	return &this
}

// NewSnapshotMarginDataUserAssetsWithDefaults instantiates a new SnapshotMarginDataUserAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotMarginDataUserAssetsWithDefaults() *SnapshotMarginDataUserAssets {
	this := SnapshotMarginDataUserAssets{}
	return &this
}

// GetAsset returns the Asset field value
func (o *SnapshotMarginDataUserAssets) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginDataUserAssets) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *SnapshotMarginDataUserAssets) SetAsset(v string) {
	o.Asset = v
}

// GetBorrowed returns the Borrowed field value
func (o *SnapshotMarginDataUserAssets) GetBorrowed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Borrowed
}

// GetBorrowedOk returns a tuple with the Borrowed field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginDataUserAssets) GetBorrowedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Borrowed, true
}

// SetBorrowed sets field value
func (o *SnapshotMarginDataUserAssets) SetBorrowed(v string) {
	o.Borrowed = v
}

// GetFree returns the Free field value
func (o *SnapshotMarginDataUserAssets) GetFree() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginDataUserAssets) GetFreeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *SnapshotMarginDataUserAssets) SetFree(v string) {
	o.Free = v
}

// GetInterest returns the Interest field value
func (o *SnapshotMarginDataUserAssets) GetInterest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interest
}

// GetInterestOk returns a tuple with the Interest field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginDataUserAssets) GetInterestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interest, true
}

// SetInterest sets field value
func (o *SnapshotMarginDataUserAssets) SetInterest(v string) {
	o.Interest = v
}

// GetLocked returns the Locked field value
func (o *SnapshotMarginDataUserAssets) GetLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginDataUserAssets) GetLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *SnapshotMarginDataUserAssets) SetLocked(v string) {
	o.Locked = v
}

// GetNetAsset returns the NetAsset field value
func (o *SnapshotMarginDataUserAssets) GetNetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetAsset
}

// GetNetAssetOk returns a tuple with the NetAsset field value
// and a boolean to check if the value has been set.
func (o *SnapshotMarginDataUserAssets) GetNetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetAsset, true
}

// SetNetAsset sets field value
func (o *SnapshotMarginDataUserAssets) SetNetAsset(v string) {
	o.NetAsset = v
}

func (o SnapshotMarginDataUserAssets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["borrowed"] = o.Borrowed
	}
	if true {
		toSerialize["free"] = o.Free
	}
	if true {
		toSerialize["interest"] = o.Interest
	}
	if true {
		toSerialize["locked"] = o.Locked
	}
	if true {
		toSerialize["netAsset"] = o.NetAsset
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotMarginDataUserAssets struct {
	value *SnapshotMarginDataUserAssets
	isSet bool
}

func (v NullableSnapshotMarginDataUserAssets) Get() *SnapshotMarginDataUserAssets {
	return v.value
}

func (v *NullableSnapshotMarginDataUserAssets) Set(val *SnapshotMarginDataUserAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotMarginDataUserAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotMarginDataUserAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotMarginDataUserAssets(val *SnapshotMarginDataUserAssets) *NullableSnapshotMarginDataUserAssets {
	return &NullableSnapshotMarginDataUserAssets{value: val, isSet: true}
}

func (v NullableSnapshotMarginDataUserAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotMarginDataUserAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
