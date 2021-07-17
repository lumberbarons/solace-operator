# MsgVpnBridgeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageRxByteRate** | Pointer to **int64** | The one minute average of the message rate received from the Bridge, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**AverageRxMsgRate** | Pointer to **int64** | The one minute average of the message rate received from the Bridge, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**AverageTxByteRate** | Pointer to **int64** | The one minute average of the message rate transmitted to the Bridge, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**AverageTxMsgRate** | Pointer to **int64** | The one minute average of the message rate transmitted to the Bridge, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**RxByteRate** | Pointer to **int64** | The current message rate received from the Bridge, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**RxMsgRate** | Pointer to **int64** | The current message rate received from the Bridge, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**TxByteRate** | Pointer to **int64** | The current message rate transmitted to the Bridge, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**TxMsgRate** | Pointer to **int64** | The current message rate transmitted to the Bridge, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 

## Methods

### NewMsgVpnBridgeRate

`func NewMsgVpnBridgeRate() *MsgVpnBridgeRate`

NewMsgVpnBridgeRate instantiates a new MsgVpnBridgeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeRateWithDefaults

`func NewMsgVpnBridgeRateWithDefaults() *MsgVpnBridgeRate`

NewMsgVpnBridgeRateWithDefaults instantiates a new MsgVpnBridgeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageRxByteRate

`func (o *MsgVpnBridgeRate) GetAverageRxByteRate() int64`

GetAverageRxByteRate returns the AverageRxByteRate field if non-nil, zero value otherwise.

### GetAverageRxByteRateOk

`func (o *MsgVpnBridgeRate) GetAverageRxByteRateOk() (*int64, bool)`

GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxByteRate

`func (o *MsgVpnBridgeRate) SetAverageRxByteRate(v int64)`

SetAverageRxByteRate sets AverageRxByteRate field to given value.

### HasAverageRxByteRate

`func (o *MsgVpnBridgeRate) HasAverageRxByteRate() bool`

HasAverageRxByteRate returns a boolean if a field has been set.

### GetAverageRxMsgRate

`func (o *MsgVpnBridgeRate) GetAverageRxMsgRate() int64`

GetAverageRxMsgRate returns the AverageRxMsgRate field if non-nil, zero value otherwise.

### GetAverageRxMsgRateOk

`func (o *MsgVpnBridgeRate) GetAverageRxMsgRateOk() (*int64, bool)`

GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxMsgRate

`func (o *MsgVpnBridgeRate) SetAverageRxMsgRate(v int64)`

SetAverageRxMsgRate sets AverageRxMsgRate field to given value.

### HasAverageRxMsgRate

`func (o *MsgVpnBridgeRate) HasAverageRxMsgRate() bool`

HasAverageRxMsgRate returns a boolean if a field has been set.

### GetAverageTxByteRate

`func (o *MsgVpnBridgeRate) GetAverageTxByteRate() int64`

GetAverageTxByteRate returns the AverageTxByteRate field if non-nil, zero value otherwise.

### GetAverageTxByteRateOk

`func (o *MsgVpnBridgeRate) GetAverageTxByteRateOk() (*int64, bool)`

GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxByteRate

`func (o *MsgVpnBridgeRate) SetAverageTxByteRate(v int64)`

SetAverageTxByteRate sets AverageTxByteRate field to given value.

### HasAverageTxByteRate

`func (o *MsgVpnBridgeRate) HasAverageTxByteRate() bool`

HasAverageTxByteRate returns a boolean if a field has been set.

### GetAverageTxMsgRate

`func (o *MsgVpnBridgeRate) GetAverageTxMsgRate() int64`

GetAverageTxMsgRate returns the AverageTxMsgRate field if non-nil, zero value otherwise.

### GetAverageTxMsgRateOk

`func (o *MsgVpnBridgeRate) GetAverageTxMsgRateOk() (*int64, bool)`

GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxMsgRate

`func (o *MsgVpnBridgeRate) SetAverageTxMsgRate(v int64)`

SetAverageTxMsgRate sets AverageTxMsgRate field to given value.

### HasAverageTxMsgRate

`func (o *MsgVpnBridgeRate) HasAverageTxMsgRate() bool`

HasAverageTxMsgRate returns a boolean if a field has been set.

### GetRxByteRate

`func (o *MsgVpnBridgeRate) GetRxByteRate() int64`

GetRxByteRate returns the RxByteRate field if non-nil, zero value otherwise.

### GetRxByteRateOk

`func (o *MsgVpnBridgeRate) GetRxByteRateOk() (*int64, bool)`

GetRxByteRateOk returns a tuple with the RxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteRate

`func (o *MsgVpnBridgeRate) SetRxByteRate(v int64)`

SetRxByteRate sets RxByteRate field to given value.

### HasRxByteRate

`func (o *MsgVpnBridgeRate) HasRxByteRate() bool`

HasRxByteRate returns a boolean if a field has been set.

### GetRxMsgRate

`func (o *MsgVpnBridgeRate) GetRxMsgRate() int64`

GetRxMsgRate returns the RxMsgRate field if non-nil, zero value otherwise.

### GetRxMsgRateOk

`func (o *MsgVpnBridgeRate) GetRxMsgRateOk() (*int64, bool)`

GetRxMsgRateOk returns a tuple with the RxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgRate

`func (o *MsgVpnBridgeRate) SetRxMsgRate(v int64)`

SetRxMsgRate sets RxMsgRate field to given value.

### HasRxMsgRate

`func (o *MsgVpnBridgeRate) HasRxMsgRate() bool`

HasRxMsgRate returns a boolean if a field has been set.

### GetTxByteRate

`func (o *MsgVpnBridgeRate) GetTxByteRate() int64`

GetTxByteRate returns the TxByteRate field if non-nil, zero value otherwise.

### GetTxByteRateOk

`func (o *MsgVpnBridgeRate) GetTxByteRateOk() (*int64, bool)`

GetTxByteRateOk returns a tuple with the TxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteRate

`func (o *MsgVpnBridgeRate) SetTxByteRate(v int64)`

SetTxByteRate sets TxByteRate field to given value.

### HasTxByteRate

`func (o *MsgVpnBridgeRate) HasTxByteRate() bool`

HasTxByteRate returns a boolean if a field has been set.

### GetTxMsgRate

`func (o *MsgVpnBridgeRate) GetTxMsgRate() int64`

GetTxMsgRate returns the TxMsgRate field if non-nil, zero value otherwise.

### GetTxMsgRateOk

`func (o *MsgVpnBridgeRate) GetTxMsgRateOk() (*int64, bool)`

GetTxMsgRateOk returns a tuple with the TxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgRate

`func (o *MsgVpnBridgeRate) SetTxMsgRate(v int64)`

SetTxMsgRate sets TxMsgRate field to given value.

### HasTxMsgRate

`func (o *MsgVpnBridgeRate) HasTxMsgRate() bool`

HasTxMsgRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


