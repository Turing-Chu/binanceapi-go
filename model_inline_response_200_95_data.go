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

// InlineResponse20095Data struct for InlineResponse20095Data
type InlineResponse20095Data struct {
	FifteenMinHashRate string                             `json:"fifteenMinHashRate"`
	DayHashRate        string                             `json:"dayHashRate"`
	ValidNum           int64                              `json:"validNum"`
	InvalidNum         int64                              `json:"invalidNum"`
	ProfitToday        InlineResponse20095DataProfitToday `json:"profitToday"`
	ProfitYesterday    InlineResponse20095DataProfitToday `json:"profitYesterday"`
	UserName           string                             `json:"userName"`
	Unit               string                             `json:"unit"`
	Algo               string                             `json:"algo"`
}

// NewInlineResponse20095Data instantiates a new InlineResponse20095Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095Data(fifteenMinHashRate string, dayHashRate string, validNum int64, invalidNum int64, profitToday InlineResponse20095DataProfitToday, profitYesterday InlineResponse20095DataProfitToday, userName string, unit string, algo string) *InlineResponse20095Data {
	this := InlineResponse20095Data{}
	this.FifteenMinHashRate = fifteenMinHashRate
	this.DayHashRate = dayHashRate
	this.ValidNum = validNum
	this.InvalidNum = invalidNum
	this.ProfitToday = profitToday
	this.ProfitYesterday = profitYesterday
	this.UserName = userName
	this.Unit = unit
	this.Algo = algo
	return &this
}

// NewInlineResponse20095DataWithDefaults instantiates a new InlineResponse20095Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095DataWithDefaults() *InlineResponse20095Data {
	this := InlineResponse20095Data{}
	return &this
}

// GetFifteenMinHashRate returns the FifteenMinHashRate field value
func (o *InlineResponse20095Data) GetFifteenMinHashRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FifteenMinHashRate
}

// GetFifteenMinHashRateOk returns a tuple with the FifteenMinHashRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Data) GetFifteenMinHashRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FifteenMinHashRate, true
}

// SetFifteenMinHashRate sets field value
func (o *InlineResponse20095Data) SetFifteenMinHashRate(v string) {
	o.FifteenMinHashRate = v
}

// GetDayHashRate returns the DayHashRate field value
func (o *InlineResponse20095Data) GetDayHashRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayHashRate
}

// GetDayHashRateOk returns a tuple with the DayHashRate field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Data) GetDayHashRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayHashRate, true
}

// SetDayHashRate sets field value
func (o *InlineResponse20095Data) SetDayHashRate(v string) {
	o.DayHashRate = v
}

// GetValidNum returns the ValidNum field value
func (o *InlineResponse20095Data) GetValidNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ValidNum
}

// GetValidNumOk returns a tuple with the ValidNum field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Data) GetValidNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidNum, true
}

// SetValidNum sets field value
func (o *InlineResponse20095Data) SetValidNum(v int64) {
	o.ValidNum = v
}

// GetInvalidNum returns the InvalidNum field value
func (o *InlineResponse20095Data) GetInvalidNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InvalidNum
}

// GetInvalidNumOk returns a tuple with the InvalidNum field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Data) GetInvalidNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvalidNum, true
}

// SetInvalidNum sets field value
func (o *InlineResponse20095Data) SetInvalidNum(v int64) {
	o.InvalidNum = v
}

// GetProfitToday returns the ProfitToday field value
func (o *InlineResponse20095Data) GetProfitToday() InlineResponse20095DataProfitToday {
	if o == nil {
		var ret InlineResponse20095DataProfitToday
		return ret
	}

	return o.ProfitToday
}

// GetProfitTodayOk returns a tuple with the ProfitToday field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Data) GetProfitTodayOk() (*InlineResponse20095DataProfitToday, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfitToday, true
}

// SetProfitToday sets field value
func (o *InlineResponse20095Data) SetProfitToday(v InlineResponse20095DataProfitToday) {
	o.ProfitToday = v
}

// GetProfitYesterday returns the ProfitYesterday field value
func (o *InlineResponse20095Data) GetProfitYesterday() InlineResponse20095DataProfitToday {
	if o == nil {
		var ret InlineResponse20095DataProfitToday
		return ret
	}

	return o.ProfitYesterday
}

// GetProfitYesterdayOk returns a tuple with the ProfitYesterday field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Data) GetProfitYesterdayOk() (*InlineResponse20095DataProfitToday, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfitYesterday, true
}

// SetProfitYesterday sets field value
func (o *InlineResponse20095Data) SetProfitYesterday(v InlineResponse20095DataProfitToday) {
	o.ProfitYesterday = v
}

// GetUserName returns the UserName field value
func (o *InlineResponse20095Data) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Data) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *InlineResponse20095Data) SetUserName(v string) {
	o.UserName = v
}

// GetUnit returns the Unit field value
func (o *InlineResponse20095Data) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Data) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *InlineResponse20095Data) SetUnit(v string) {
	o.Unit = v
}

// GetAlgo returns the Algo field value
func (o *InlineResponse20095Data) GetAlgo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algo
}

// GetAlgoOk returns a tuple with the Algo field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20095Data) GetAlgoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algo, true
}

// SetAlgo sets field value
func (o *InlineResponse20095Data) SetAlgo(v string) {
	o.Algo = v
}

func (o InlineResponse20095Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fifteenMinHashRate"] = o.FifteenMinHashRate
	}
	if true {
		toSerialize["dayHashRate"] = o.DayHashRate
	}
	if true {
		toSerialize["validNum"] = o.ValidNum
	}
	if true {
		toSerialize["invalidNum"] = o.InvalidNum
	}
	if true {
		toSerialize["profitToday"] = o.ProfitToday
	}
	if true {
		toSerialize["profitYesterday"] = o.ProfitYesterday
	}
	if true {
		toSerialize["userName"] = o.UserName
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	if true {
		toSerialize["algo"] = o.Algo
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20095Data struct {
	value *InlineResponse20095Data
	isSet bool
}

func (v NullableInlineResponse20095Data) Get() *InlineResponse20095Data {
	return v.value
}

func (v *NullableInlineResponse20095Data) Set(val *InlineResponse20095Data) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095Data) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095Data(val *InlineResponse20095Data) *NullableInlineResponse20095Data {
	return &NullableInlineResponse20095Data{value: val, isSet: true}
}

func (v NullableInlineResponse20095Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
