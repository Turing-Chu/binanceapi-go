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

// SavingsFlexiblePurchaseRecord struct for SavingsFlexiblePurchaseRecord
type SavingsFlexiblePurchaseRecord struct {
	Items []map[string]interface{}
}

// NewSavingsFlexiblePurchaseRecord instantiates a new SavingsFlexiblePurchaseRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavingsFlexiblePurchaseRecord() *SavingsFlexiblePurchaseRecord {
	this := SavingsFlexiblePurchaseRecord{}
	return &this
}

// NewSavingsFlexiblePurchaseRecordWithDefaults instantiates a new SavingsFlexiblePurchaseRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavingsFlexiblePurchaseRecordWithDefaults() *SavingsFlexiblePurchaseRecord {
	this := SavingsFlexiblePurchaseRecord{}
	return &this
}

func (o SavingsFlexiblePurchaseRecord) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return json.Marshal(toSerialize)
}

func (o *SavingsFlexiblePurchaseRecord) UnmarshalJSON(bytes []byte) (err error) {
	return json.Unmarshal(bytes, &o.Items)
}

type NullableSavingsFlexiblePurchaseRecord struct {
	value *SavingsFlexiblePurchaseRecord
	isSet bool
}

func (v NullableSavingsFlexiblePurchaseRecord) Get() *SavingsFlexiblePurchaseRecord {
	return v.value
}

func (v *NullableSavingsFlexiblePurchaseRecord) Set(val *SavingsFlexiblePurchaseRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableSavingsFlexiblePurchaseRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableSavingsFlexiblePurchaseRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavingsFlexiblePurchaseRecord(val *SavingsFlexiblePurchaseRecord) *NullableSavingsFlexiblePurchaseRecord {
	return &NullableSavingsFlexiblePurchaseRecord{value: val, isSet: true}
}

func (v NullableSavingsFlexiblePurchaseRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavingsFlexiblePurchaseRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
