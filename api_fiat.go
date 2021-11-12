/*
Binance Public Spot API

OpenAPI Specifications for the Binance Public Spot API generated with [binance/binance-api-swagger/blob/master/spot_api.yaml](https://github.com/binance/binance-api-swagger/blob/master/spot_api.yaml) with commit [v1.2.0 release](https://github.com/binance/binance-api-swagger/commit/60d14be031c031600c853d5cdab86db5ab73603e)  API documents:   - [https://github.com/binance/binance-spot-api-docs](https://github.com/binance/binance-spot-api-docs)   - [https://binance-docs.github.io/apidocs/spot/en](https://binance-docs.github.io/apidocs/spot/en)

API version: 1.0
Contact: qishiwenjun@163.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binanceapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// FiatApiService FiatApi service
type FiatApiService service

type ApiFiatGetHistoryOrdersRequest struct {
	ctx             _context.Context
	ApiService      *FiatApiService
	transactionType *string
	timestamp       *int64
	signature       *string
	beginTime       *int64
	endTime         *int64
	page            *int32
	rows            *int32
	recvWindow      *int64
}

// 0-deposit, 1-withdraw
func (r ApiFiatGetHistoryOrdersRequest) TransactionType(transactionType string) ApiFiatGetHistoryOrdersRequest {
	r.transactionType = &transactionType
	return r
}

// UTC timestamp in ms
func (r ApiFiatGetHistoryOrdersRequest) Timestamp(timestamp int64) ApiFiatGetHistoryOrdersRequest {
	r.timestamp = &timestamp
	return r
}

// Signature
func (r ApiFiatGetHistoryOrdersRequest) Signature(signature string) ApiFiatGetHistoryOrdersRequest {
	r.signature = &signature
	return r
}
func (r ApiFiatGetHistoryOrdersRequest) BeginTime(beginTime int64) ApiFiatGetHistoryOrdersRequest {
	r.beginTime = &beginTime
	return r
}

// UTC timestamp in ms
func (r ApiFiatGetHistoryOrdersRequest) EndTime(endTime int64) ApiFiatGetHistoryOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default 1
func (r ApiFiatGetHistoryOrdersRequest) Page(page int32) ApiFiatGetHistoryOrdersRequest {
	r.page = &page
	return r
}

// Default 100, max 500
func (r ApiFiatGetHistoryOrdersRequest) Rows(rows int32) ApiFiatGetHistoryOrdersRequest {
	r.rows = &rows
	return r
}

// The value cannot be greater than 60000
func (r ApiFiatGetHistoryOrdersRequest) RecvWindow(recvWindow int64) ApiFiatGetHistoryOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFiatGetHistoryOrdersRequest) Execute() (InlineResponse20072, *_nethttp.Response, error) {
	return r.ApiService.FiatGetHistoryOrdersExecute(r)
}

/*
FiatGetHistoryOrders Fiat Deposit/Withdraw History (USER_DATA)

- If beginTime and endTime are not sent, the recent 30-day data will be returned.

Weight(IP): 1

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFiatGetHistoryOrdersRequest
*/
func (a *FiatApiService) FiatGetHistoryOrders(ctx _context.Context) ApiFiatGetHistoryOrdersRequest {
	return ApiFiatGetHistoryOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return InlineResponse20072
func (a *FiatApiService) FiatGetHistoryOrdersExecute(r ApiFiatGetHistoryOrdersRequest) (InlineResponse20072, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20072
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FiatApiService.FiatGetHistoryOrders")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sapi/v1/fiat/orders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.transactionType == nil {
		return localVarReturnValue, nil, reportError("transactionType is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}
	if r.signature == nil {
		return localVarReturnValue, nil, reportError("signature is required and must be specified")
	}

	localVarQueryParams.Add("transactionType", parameterToString(*r.transactionType, ""))
	if r.beginTime != nil {
		localVarQueryParams.Add("beginTime", parameterToString(*r.beginTime, ""))
	}
	if r.endTime != nil {
		localVarQueryParams.Add("endTime", parameterToString(*r.endTime, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.rows != nil {
		localVarQueryParams.Add("rows", parameterToString(*r.rows, ""))
	}
	if r.recvWindow != nil {
		localVarQueryParams.Add("recvWindow", parameterToString(*r.recvWindow, ""))
	}
	localVarQueryParams.Add("timestamp", parameterToString(*r.timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(*r.signature, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-MBX-APIKEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFiatGetPaymentsHistoryRequest struct {
	ctx             _context.Context
	ApiService      *FiatApiService
	transactionType *string
	timestamp       *int64
	signature       *string
	beginTime       *int64
	endTime         *int64
	page            *int32
	rows            *int32
	recvWindow      *int64
}

// 0-deposit, 1-withdraw
func (r ApiFiatGetPaymentsHistoryRequest) TransactionType(transactionType string) ApiFiatGetPaymentsHistoryRequest {
	r.transactionType = &transactionType
	return r
}

// UTC timestamp in ms
func (r ApiFiatGetPaymentsHistoryRequest) Timestamp(timestamp int64) ApiFiatGetPaymentsHistoryRequest {
	r.timestamp = &timestamp
	return r
}

// Signature
func (r ApiFiatGetPaymentsHistoryRequest) Signature(signature string) ApiFiatGetPaymentsHistoryRequest {
	r.signature = &signature
	return r
}
func (r ApiFiatGetPaymentsHistoryRequest) BeginTime(beginTime int64) ApiFiatGetPaymentsHistoryRequest {
	r.beginTime = &beginTime
	return r
}

// UTC timestamp in ms
func (r ApiFiatGetPaymentsHistoryRequest) EndTime(endTime int64) ApiFiatGetPaymentsHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default 1
func (r ApiFiatGetPaymentsHistoryRequest) Page(page int32) ApiFiatGetPaymentsHistoryRequest {
	r.page = &page
	return r
}

// Default 100, max 500
func (r ApiFiatGetPaymentsHistoryRequest) Rows(rows int32) ApiFiatGetPaymentsHistoryRequest {
	r.rows = &rows
	return r
}

// The value cannot be greater than 60000
func (r ApiFiatGetPaymentsHistoryRequest) RecvWindow(recvWindow int64) ApiFiatGetPaymentsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFiatGetPaymentsHistoryRequest) Execute() (InlineResponse20073, *_nethttp.Response, error) {
	return r.ApiService.FiatGetPaymentsHistoryExecute(r)
}

/*
FiatGetPaymentsHistory Fiat Payments History (USER_DATA)

- If beginTime and endTime are not sent, the recent 30-day data will be returned.

Weight(IP): 1

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFiatGetPaymentsHistoryRequest
*/
func (a *FiatApiService) FiatGetPaymentsHistory(ctx _context.Context) ApiFiatGetPaymentsHistoryRequest {
	return ApiFiatGetPaymentsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return InlineResponse20073
func (a *FiatApiService) FiatGetPaymentsHistoryExecute(r ApiFiatGetPaymentsHistoryRequest) (InlineResponse20073, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20073
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FiatApiService.FiatGetPaymentsHistory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sapi/v1/fiat/payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.transactionType == nil {
		return localVarReturnValue, nil, reportError("transactionType is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}
	if r.signature == nil {
		return localVarReturnValue, nil, reportError("signature is required and must be specified")
	}

	localVarQueryParams.Add("transactionType", parameterToString(*r.transactionType, ""))
	if r.beginTime != nil {
		localVarQueryParams.Add("beginTime", parameterToString(*r.beginTime, ""))
	}
	if r.endTime != nil {
		localVarQueryParams.Add("endTime", parameterToString(*r.endTime, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.rows != nil {
		localVarQueryParams.Add("rows", parameterToString(*r.rows, ""))
	}
	if r.recvWindow != nil {
		localVarQueryParams.Add("recvWindow", parameterToString(*r.recvWindow, ""))
	}
	localVarQueryParams.Add("timestamp", parameterToString(*r.timestamp, ""))
	localVarQueryParams.Add("signature", parameterToString(*r.signature, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-MBX-APIKEY"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
