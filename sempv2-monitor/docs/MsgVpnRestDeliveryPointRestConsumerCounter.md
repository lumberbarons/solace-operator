# MsgVpnRestDeliveryPointRestConsumerCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpRequestConnectionCloseTxMsgCount** | Pointer to **int64** | The number of HTTP request messages transmitted to the REST Consumer to close the connection. Deprecated since 2.13. This attribute has been moved to the MsgVpnRestDeliveryPointRestConsumer object. | [optional] 
**HttpRequestOutstandingTxMsgCount** | Pointer to **int64** | The number of HTTP request messages transmitted to the REST Consumer that are waiting for a response. Deprecated since 2.13. This attribute has been moved to the MsgVpnRestDeliveryPointRestConsumer object. | [optional] 
**HttpRequestTimedOutTxMsgCount** | Pointer to **int64** | The number of HTTP request messages transmitted to the REST Consumer that have timed out. Deprecated since 2.13. This attribute has been moved to the MsgVpnRestDeliveryPointRestConsumer object. | [optional] 
**HttpRequestTxByteCount** | Pointer to **int64** | The amount of HTTP request messages transmitted to the REST Consumer, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnRestDeliveryPointRestConsumer object. | [optional] 
**HttpRequestTxMsgCount** | Pointer to **int64** | The number of HTTP request messages transmitted to the REST Consumer. Deprecated since 2.13. This attribute has been moved to the MsgVpnRestDeliveryPointRestConsumer object. | [optional] 
**HttpResponseErrorRxMsgCount** | Pointer to **int64** | The number of HTTP client/server error response messages received from the REST Consumer. Deprecated since 2.13. This attribute has been moved to the MsgVpnRestDeliveryPointRestConsumer object. | [optional] 
**HttpResponseRxByteCount** | Pointer to **int64** | The amount of HTTP response messages received from the REST Consumer, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnRestDeliveryPointRestConsumer object. | [optional] 
**HttpResponseRxMsgCount** | Pointer to **int64** | The number of HTTP response messages received from the REST Consumer. Deprecated since 2.13. This attribute has been moved to the MsgVpnRestDeliveryPointRestConsumer object. | [optional] 
**HttpResponseSuccessRxMsgCount** | Pointer to **int64** | The number of HTTP successful response messages received from the REST Consumer. Deprecated since 2.13. This attribute has been moved to the MsgVpnRestDeliveryPointRestConsumer object. | [optional] 

## Methods

### NewMsgVpnRestDeliveryPointRestConsumerCounter

`func NewMsgVpnRestDeliveryPointRestConsumerCounter() *MsgVpnRestDeliveryPointRestConsumerCounter`

NewMsgVpnRestDeliveryPointRestConsumerCounter instantiates a new MsgVpnRestDeliveryPointRestConsumerCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnRestDeliveryPointRestConsumerCounterWithDefaults

`func NewMsgVpnRestDeliveryPointRestConsumerCounterWithDefaults() *MsgVpnRestDeliveryPointRestConsumerCounter`

NewMsgVpnRestDeliveryPointRestConsumerCounterWithDefaults instantiates a new MsgVpnRestDeliveryPointRestConsumerCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpRequestConnectionCloseTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpRequestConnectionCloseTxMsgCount() int64`

GetHttpRequestConnectionCloseTxMsgCount returns the HttpRequestConnectionCloseTxMsgCount field if non-nil, zero value otherwise.

### GetHttpRequestConnectionCloseTxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpRequestConnectionCloseTxMsgCountOk() (*int64, bool)`

GetHttpRequestConnectionCloseTxMsgCountOk returns a tuple with the HttpRequestConnectionCloseTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestConnectionCloseTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) SetHttpRequestConnectionCloseTxMsgCount(v int64)`

SetHttpRequestConnectionCloseTxMsgCount sets HttpRequestConnectionCloseTxMsgCount field to given value.

### HasHttpRequestConnectionCloseTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) HasHttpRequestConnectionCloseTxMsgCount() bool`

HasHttpRequestConnectionCloseTxMsgCount returns a boolean if a field has been set.

### GetHttpRequestOutstandingTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpRequestOutstandingTxMsgCount() int64`

GetHttpRequestOutstandingTxMsgCount returns the HttpRequestOutstandingTxMsgCount field if non-nil, zero value otherwise.

### GetHttpRequestOutstandingTxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpRequestOutstandingTxMsgCountOk() (*int64, bool)`

GetHttpRequestOutstandingTxMsgCountOk returns a tuple with the HttpRequestOutstandingTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestOutstandingTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) SetHttpRequestOutstandingTxMsgCount(v int64)`

SetHttpRequestOutstandingTxMsgCount sets HttpRequestOutstandingTxMsgCount field to given value.

### HasHttpRequestOutstandingTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) HasHttpRequestOutstandingTxMsgCount() bool`

HasHttpRequestOutstandingTxMsgCount returns a boolean if a field has been set.

### GetHttpRequestTimedOutTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpRequestTimedOutTxMsgCount() int64`

GetHttpRequestTimedOutTxMsgCount returns the HttpRequestTimedOutTxMsgCount field if non-nil, zero value otherwise.

### GetHttpRequestTimedOutTxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpRequestTimedOutTxMsgCountOk() (*int64, bool)`

GetHttpRequestTimedOutTxMsgCountOk returns a tuple with the HttpRequestTimedOutTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestTimedOutTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) SetHttpRequestTimedOutTxMsgCount(v int64)`

SetHttpRequestTimedOutTxMsgCount sets HttpRequestTimedOutTxMsgCount field to given value.

### HasHttpRequestTimedOutTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) HasHttpRequestTimedOutTxMsgCount() bool`

HasHttpRequestTimedOutTxMsgCount returns a boolean if a field has been set.

### GetHttpRequestTxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpRequestTxByteCount() int64`

GetHttpRequestTxByteCount returns the HttpRequestTxByteCount field if non-nil, zero value otherwise.

### GetHttpRequestTxByteCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpRequestTxByteCountOk() (*int64, bool)`

GetHttpRequestTxByteCountOk returns a tuple with the HttpRequestTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestTxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) SetHttpRequestTxByteCount(v int64)`

SetHttpRequestTxByteCount sets HttpRequestTxByteCount field to given value.

### HasHttpRequestTxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) HasHttpRequestTxByteCount() bool`

HasHttpRequestTxByteCount returns a boolean if a field has been set.

### GetHttpRequestTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpRequestTxMsgCount() int64`

GetHttpRequestTxMsgCount returns the HttpRequestTxMsgCount field if non-nil, zero value otherwise.

### GetHttpRequestTxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpRequestTxMsgCountOk() (*int64, bool)`

GetHttpRequestTxMsgCountOk returns a tuple with the HttpRequestTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) SetHttpRequestTxMsgCount(v int64)`

SetHttpRequestTxMsgCount sets HttpRequestTxMsgCount field to given value.

### HasHttpRequestTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) HasHttpRequestTxMsgCount() bool`

HasHttpRequestTxMsgCount returns a boolean if a field has been set.

### GetHttpResponseErrorRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpResponseErrorRxMsgCount() int64`

GetHttpResponseErrorRxMsgCount returns the HttpResponseErrorRxMsgCount field if non-nil, zero value otherwise.

### GetHttpResponseErrorRxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpResponseErrorRxMsgCountOk() (*int64, bool)`

GetHttpResponseErrorRxMsgCountOk returns a tuple with the HttpResponseErrorRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseErrorRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) SetHttpResponseErrorRxMsgCount(v int64)`

SetHttpResponseErrorRxMsgCount sets HttpResponseErrorRxMsgCount field to given value.

### HasHttpResponseErrorRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) HasHttpResponseErrorRxMsgCount() bool`

HasHttpResponseErrorRxMsgCount returns a boolean if a field has been set.

### GetHttpResponseRxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpResponseRxByteCount() int64`

GetHttpResponseRxByteCount returns the HttpResponseRxByteCount field if non-nil, zero value otherwise.

### GetHttpResponseRxByteCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpResponseRxByteCountOk() (*int64, bool)`

GetHttpResponseRxByteCountOk returns a tuple with the HttpResponseRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseRxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) SetHttpResponseRxByteCount(v int64)`

SetHttpResponseRxByteCount sets HttpResponseRxByteCount field to given value.

### HasHttpResponseRxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) HasHttpResponseRxByteCount() bool`

HasHttpResponseRxByteCount returns a boolean if a field has been set.

### GetHttpResponseRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpResponseRxMsgCount() int64`

GetHttpResponseRxMsgCount returns the HttpResponseRxMsgCount field if non-nil, zero value otherwise.

### GetHttpResponseRxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpResponseRxMsgCountOk() (*int64, bool)`

GetHttpResponseRxMsgCountOk returns a tuple with the HttpResponseRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) SetHttpResponseRxMsgCount(v int64)`

SetHttpResponseRxMsgCount sets HttpResponseRxMsgCount field to given value.

### HasHttpResponseRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) HasHttpResponseRxMsgCount() bool`

HasHttpResponseRxMsgCount returns a boolean if a field has been set.

### GetHttpResponseSuccessRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpResponseSuccessRxMsgCount() int64`

GetHttpResponseSuccessRxMsgCount returns the HttpResponseSuccessRxMsgCount field if non-nil, zero value otherwise.

### GetHttpResponseSuccessRxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) GetHttpResponseSuccessRxMsgCountOk() (*int64, bool)`

GetHttpResponseSuccessRxMsgCountOk returns a tuple with the HttpResponseSuccessRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseSuccessRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) SetHttpResponseSuccessRxMsgCount(v int64)`

SetHttpResponseSuccessRxMsgCount sets HttpResponseSuccessRxMsgCount field to given value.

### HasHttpResponseSuccessRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumerCounter) HasHttpResponseSuccessRxMsgCount() bool`

HasHttpResponseSuccessRxMsgCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


