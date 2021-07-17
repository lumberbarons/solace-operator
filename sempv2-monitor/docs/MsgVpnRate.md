# MsgVpnRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageRxByteRate** | Pointer to **int64** | The one minute average of the message rate received by the Message VPN, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**AverageRxMsgRate** | Pointer to **int64** | The one minute average of the message rate received by the Message VPN, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**AverageTxByteRate** | Pointer to **int64** | The one minute average of the message rate transmitted by the Message VPN, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**AverageTxMsgRate** | Pointer to **int64** | The one minute average of the message rate transmitted by the Message VPN, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**RxByteRate** | Pointer to **int64** | The current message rate received by the Message VPN, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**RxMsgRate** | Pointer to **int64** | The current message rate received by the Message VPN, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**TlsAverageRxByteRate** | Pointer to **int64** | The one minute average of the TLS message rate received by the Message VPN, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**TlsAverageTxByteRate** | Pointer to **int64** | The one minute average of the TLS message rate transmitted by the Message VPN, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**TlsRxByteRate** | Pointer to **int64** | The current TLS message rate received by the Message VPN, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**TlsTxByteRate** | Pointer to **int64** | The current TLS message rate transmitted by the Message VPN, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**TxByteRate** | Pointer to **int64** | The current message rate transmitted by the Message VPN, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**TxMsgRate** | Pointer to **int64** | The current message rate transmitted by the Message VPN, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 

## Methods

### NewMsgVpnRate

`func NewMsgVpnRate() *MsgVpnRate`

NewMsgVpnRate instantiates a new MsgVpnRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnRateWithDefaults

`func NewMsgVpnRateWithDefaults() *MsgVpnRate`

NewMsgVpnRateWithDefaults instantiates a new MsgVpnRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageRxByteRate

`func (o *MsgVpnRate) GetAverageRxByteRate() int64`

GetAverageRxByteRate returns the AverageRxByteRate field if non-nil, zero value otherwise.

### GetAverageRxByteRateOk

`func (o *MsgVpnRate) GetAverageRxByteRateOk() (*int64, bool)`

GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxByteRate

`func (o *MsgVpnRate) SetAverageRxByteRate(v int64)`

SetAverageRxByteRate sets AverageRxByteRate field to given value.

### HasAverageRxByteRate

`func (o *MsgVpnRate) HasAverageRxByteRate() bool`

HasAverageRxByteRate returns a boolean if a field has been set.

### GetAverageRxMsgRate

`func (o *MsgVpnRate) GetAverageRxMsgRate() int64`

GetAverageRxMsgRate returns the AverageRxMsgRate field if non-nil, zero value otherwise.

### GetAverageRxMsgRateOk

`func (o *MsgVpnRate) GetAverageRxMsgRateOk() (*int64, bool)`

GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxMsgRate

`func (o *MsgVpnRate) SetAverageRxMsgRate(v int64)`

SetAverageRxMsgRate sets AverageRxMsgRate field to given value.

### HasAverageRxMsgRate

`func (o *MsgVpnRate) HasAverageRxMsgRate() bool`

HasAverageRxMsgRate returns a boolean if a field has been set.

### GetAverageTxByteRate

`func (o *MsgVpnRate) GetAverageTxByteRate() int64`

GetAverageTxByteRate returns the AverageTxByteRate field if non-nil, zero value otherwise.

### GetAverageTxByteRateOk

`func (o *MsgVpnRate) GetAverageTxByteRateOk() (*int64, bool)`

GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxByteRate

`func (o *MsgVpnRate) SetAverageTxByteRate(v int64)`

SetAverageTxByteRate sets AverageTxByteRate field to given value.

### HasAverageTxByteRate

`func (o *MsgVpnRate) HasAverageTxByteRate() bool`

HasAverageTxByteRate returns a boolean if a field has been set.

### GetAverageTxMsgRate

`func (o *MsgVpnRate) GetAverageTxMsgRate() int64`

GetAverageTxMsgRate returns the AverageTxMsgRate field if non-nil, zero value otherwise.

### GetAverageTxMsgRateOk

`func (o *MsgVpnRate) GetAverageTxMsgRateOk() (*int64, bool)`

GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxMsgRate

`func (o *MsgVpnRate) SetAverageTxMsgRate(v int64)`

SetAverageTxMsgRate sets AverageTxMsgRate field to given value.

### HasAverageTxMsgRate

`func (o *MsgVpnRate) HasAverageTxMsgRate() bool`

HasAverageTxMsgRate returns a boolean if a field has been set.

### GetRxByteRate

`func (o *MsgVpnRate) GetRxByteRate() int64`

GetRxByteRate returns the RxByteRate field if non-nil, zero value otherwise.

### GetRxByteRateOk

`func (o *MsgVpnRate) GetRxByteRateOk() (*int64, bool)`

GetRxByteRateOk returns a tuple with the RxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteRate

`func (o *MsgVpnRate) SetRxByteRate(v int64)`

SetRxByteRate sets RxByteRate field to given value.

### HasRxByteRate

`func (o *MsgVpnRate) HasRxByteRate() bool`

HasRxByteRate returns a boolean if a field has been set.

### GetRxMsgRate

`func (o *MsgVpnRate) GetRxMsgRate() int64`

GetRxMsgRate returns the RxMsgRate field if non-nil, zero value otherwise.

### GetRxMsgRateOk

`func (o *MsgVpnRate) GetRxMsgRateOk() (*int64, bool)`

GetRxMsgRateOk returns a tuple with the RxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgRate

`func (o *MsgVpnRate) SetRxMsgRate(v int64)`

SetRxMsgRate sets RxMsgRate field to given value.

### HasRxMsgRate

`func (o *MsgVpnRate) HasRxMsgRate() bool`

HasRxMsgRate returns a boolean if a field has been set.

### GetTlsAverageRxByteRate

`func (o *MsgVpnRate) GetTlsAverageRxByteRate() int64`

GetTlsAverageRxByteRate returns the TlsAverageRxByteRate field if non-nil, zero value otherwise.

### GetTlsAverageRxByteRateOk

`func (o *MsgVpnRate) GetTlsAverageRxByteRateOk() (*int64, bool)`

GetTlsAverageRxByteRateOk returns a tuple with the TlsAverageRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAverageRxByteRate

`func (o *MsgVpnRate) SetTlsAverageRxByteRate(v int64)`

SetTlsAverageRxByteRate sets TlsAverageRxByteRate field to given value.

### HasTlsAverageRxByteRate

`func (o *MsgVpnRate) HasTlsAverageRxByteRate() bool`

HasTlsAverageRxByteRate returns a boolean if a field has been set.

### GetTlsAverageTxByteRate

`func (o *MsgVpnRate) GetTlsAverageTxByteRate() int64`

GetTlsAverageTxByteRate returns the TlsAverageTxByteRate field if non-nil, zero value otherwise.

### GetTlsAverageTxByteRateOk

`func (o *MsgVpnRate) GetTlsAverageTxByteRateOk() (*int64, bool)`

GetTlsAverageTxByteRateOk returns a tuple with the TlsAverageTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAverageTxByteRate

`func (o *MsgVpnRate) SetTlsAverageTxByteRate(v int64)`

SetTlsAverageTxByteRate sets TlsAverageTxByteRate field to given value.

### HasTlsAverageTxByteRate

`func (o *MsgVpnRate) HasTlsAverageTxByteRate() bool`

HasTlsAverageTxByteRate returns a boolean if a field has been set.

### GetTlsRxByteRate

`func (o *MsgVpnRate) GetTlsRxByteRate() int64`

GetTlsRxByteRate returns the TlsRxByteRate field if non-nil, zero value otherwise.

### GetTlsRxByteRateOk

`func (o *MsgVpnRate) GetTlsRxByteRateOk() (*int64, bool)`

GetTlsRxByteRateOk returns a tuple with the TlsRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsRxByteRate

`func (o *MsgVpnRate) SetTlsRxByteRate(v int64)`

SetTlsRxByteRate sets TlsRxByteRate field to given value.

### HasTlsRxByteRate

`func (o *MsgVpnRate) HasTlsRxByteRate() bool`

HasTlsRxByteRate returns a boolean if a field has been set.

### GetTlsTxByteRate

`func (o *MsgVpnRate) GetTlsTxByteRate() int64`

GetTlsTxByteRate returns the TlsTxByteRate field if non-nil, zero value otherwise.

### GetTlsTxByteRateOk

`func (o *MsgVpnRate) GetTlsTxByteRateOk() (*int64, bool)`

GetTlsTxByteRateOk returns a tuple with the TlsTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTxByteRate

`func (o *MsgVpnRate) SetTlsTxByteRate(v int64)`

SetTlsTxByteRate sets TlsTxByteRate field to given value.

### HasTlsTxByteRate

`func (o *MsgVpnRate) HasTlsTxByteRate() bool`

HasTlsTxByteRate returns a boolean if a field has been set.

### GetTxByteRate

`func (o *MsgVpnRate) GetTxByteRate() int64`

GetTxByteRate returns the TxByteRate field if non-nil, zero value otherwise.

### GetTxByteRateOk

`func (o *MsgVpnRate) GetTxByteRateOk() (*int64, bool)`

GetTxByteRateOk returns a tuple with the TxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteRate

`func (o *MsgVpnRate) SetTxByteRate(v int64)`

SetTxByteRate sets TxByteRate field to given value.

### HasTxByteRate

`func (o *MsgVpnRate) HasTxByteRate() bool`

HasTxByteRate returns a boolean if a field has been set.

### GetTxMsgRate

`func (o *MsgVpnRate) GetTxMsgRate() int64`

GetTxMsgRate returns the TxMsgRate field if non-nil, zero value otherwise.

### GetTxMsgRateOk

`func (o *MsgVpnRate) GetTxMsgRateOk() (*int64, bool)`

GetTxMsgRateOk returns a tuple with the TxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgRate

`func (o *MsgVpnRate) SetTxMsgRate(v int64)`

SetTxMsgRate sets TxMsgRate field to given value.

### HasTxMsgRate

`func (o *MsgVpnRate) HasTxMsgRate() bool`

HasTxMsgRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


