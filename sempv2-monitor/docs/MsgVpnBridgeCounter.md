# MsgVpnBridgeCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlRxByteCount** | Pointer to **int64** | The amount of client control messages received from the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**ControlRxMsgCount** | Pointer to **int64** | The number of client control messages received from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**ControlTxByteCount** | Pointer to **int64** | The amount of client control messages transmitted to the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**ControlTxMsgCount** | Pointer to **int64** | The number of client control messages transmitted to the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**DataRxByteCount** | Pointer to **int64** | The amount of client data messages received from the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**DataRxMsgCount** | Pointer to **int64** | The number of client data messages received from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**DataTxByteCount** | Pointer to **int64** | The amount of client data messages transmitted to the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**DataTxMsgCount** | Pointer to **int64** | The number of client data messages transmitted to the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**DiscardedRxMsgCount** | Pointer to **int64** | The number of messages discarded during reception from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**DiscardedTxMsgCount** | Pointer to **int64** | The number of messages discarded during transmission to the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**LoginRxMsgCount** | Pointer to **int64** | The number of login request messages received from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**LoginTxMsgCount** | Pointer to **int64** | The number of login response messages transmitted to the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**MsgSpoolRxMsgCount** | Pointer to **int64** | The number of guaranteed messages received from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**RxByteCount** | Pointer to **int64** | The amount of messages received from the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**RxMsgCount** | Pointer to **int64** | The number of messages received from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**TxByteCount** | Pointer to **int64** | The amount of messages transmitted to the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 
**TxMsgCount** | Pointer to **int64** | The number of messages transmitted to the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object. | [optional] 

## Methods

### NewMsgVpnBridgeCounter

`func NewMsgVpnBridgeCounter() *MsgVpnBridgeCounter`

NewMsgVpnBridgeCounter instantiates a new MsgVpnBridgeCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeCounterWithDefaults

`func NewMsgVpnBridgeCounterWithDefaults() *MsgVpnBridgeCounter`

NewMsgVpnBridgeCounterWithDefaults instantiates a new MsgVpnBridgeCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlRxByteCount

`func (o *MsgVpnBridgeCounter) GetControlRxByteCount() int64`

GetControlRxByteCount returns the ControlRxByteCount field if non-nil, zero value otherwise.

### GetControlRxByteCountOk

`func (o *MsgVpnBridgeCounter) GetControlRxByteCountOk() (*int64, bool)`

GetControlRxByteCountOk returns a tuple with the ControlRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRxByteCount

`func (o *MsgVpnBridgeCounter) SetControlRxByteCount(v int64)`

SetControlRxByteCount sets ControlRxByteCount field to given value.

### HasControlRxByteCount

`func (o *MsgVpnBridgeCounter) HasControlRxByteCount() bool`

HasControlRxByteCount returns a boolean if a field has been set.

### GetControlRxMsgCount

`func (o *MsgVpnBridgeCounter) GetControlRxMsgCount() int64`

GetControlRxMsgCount returns the ControlRxMsgCount field if non-nil, zero value otherwise.

### GetControlRxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetControlRxMsgCountOk() (*int64, bool)`

GetControlRxMsgCountOk returns a tuple with the ControlRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRxMsgCount

`func (o *MsgVpnBridgeCounter) SetControlRxMsgCount(v int64)`

SetControlRxMsgCount sets ControlRxMsgCount field to given value.

### HasControlRxMsgCount

`func (o *MsgVpnBridgeCounter) HasControlRxMsgCount() bool`

HasControlRxMsgCount returns a boolean if a field has been set.

### GetControlTxByteCount

`func (o *MsgVpnBridgeCounter) GetControlTxByteCount() int64`

GetControlTxByteCount returns the ControlTxByteCount field if non-nil, zero value otherwise.

### GetControlTxByteCountOk

`func (o *MsgVpnBridgeCounter) GetControlTxByteCountOk() (*int64, bool)`

GetControlTxByteCountOk returns a tuple with the ControlTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTxByteCount

`func (o *MsgVpnBridgeCounter) SetControlTxByteCount(v int64)`

SetControlTxByteCount sets ControlTxByteCount field to given value.

### HasControlTxByteCount

`func (o *MsgVpnBridgeCounter) HasControlTxByteCount() bool`

HasControlTxByteCount returns a boolean if a field has been set.

### GetControlTxMsgCount

`func (o *MsgVpnBridgeCounter) GetControlTxMsgCount() int64`

GetControlTxMsgCount returns the ControlTxMsgCount field if non-nil, zero value otherwise.

### GetControlTxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetControlTxMsgCountOk() (*int64, bool)`

GetControlTxMsgCountOk returns a tuple with the ControlTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTxMsgCount

`func (o *MsgVpnBridgeCounter) SetControlTxMsgCount(v int64)`

SetControlTxMsgCount sets ControlTxMsgCount field to given value.

### HasControlTxMsgCount

`func (o *MsgVpnBridgeCounter) HasControlTxMsgCount() bool`

HasControlTxMsgCount returns a boolean if a field has been set.

### GetDataRxByteCount

`func (o *MsgVpnBridgeCounter) GetDataRxByteCount() int64`

GetDataRxByteCount returns the DataRxByteCount field if non-nil, zero value otherwise.

### GetDataRxByteCountOk

`func (o *MsgVpnBridgeCounter) GetDataRxByteCountOk() (*int64, bool)`

GetDataRxByteCountOk returns a tuple with the DataRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxByteCount

`func (o *MsgVpnBridgeCounter) SetDataRxByteCount(v int64)`

SetDataRxByteCount sets DataRxByteCount field to given value.

### HasDataRxByteCount

`func (o *MsgVpnBridgeCounter) HasDataRxByteCount() bool`

HasDataRxByteCount returns a boolean if a field has been set.

### GetDataRxMsgCount

`func (o *MsgVpnBridgeCounter) GetDataRxMsgCount() int64`

GetDataRxMsgCount returns the DataRxMsgCount field if non-nil, zero value otherwise.

### GetDataRxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetDataRxMsgCountOk() (*int64, bool)`

GetDataRxMsgCountOk returns a tuple with the DataRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxMsgCount

`func (o *MsgVpnBridgeCounter) SetDataRxMsgCount(v int64)`

SetDataRxMsgCount sets DataRxMsgCount field to given value.

### HasDataRxMsgCount

`func (o *MsgVpnBridgeCounter) HasDataRxMsgCount() bool`

HasDataRxMsgCount returns a boolean if a field has been set.

### GetDataTxByteCount

`func (o *MsgVpnBridgeCounter) GetDataTxByteCount() int64`

GetDataTxByteCount returns the DataTxByteCount field if non-nil, zero value otherwise.

### GetDataTxByteCountOk

`func (o *MsgVpnBridgeCounter) GetDataTxByteCountOk() (*int64, bool)`

GetDataTxByteCountOk returns a tuple with the DataTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxByteCount

`func (o *MsgVpnBridgeCounter) SetDataTxByteCount(v int64)`

SetDataTxByteCount sets DataTxByteCount field to given value.

### HasDataTxByteCount

`func (o *MsgVpnBridgeCounter) HasDataTxByteCount() bool`

HasDataTxByteCount returns a boolean if a field has been set.

### GetDataTxMsgCount

`func (o *MsgVpnBridgeCounter) GetDataTxMsgCount() int64`

GetDataTxMsgCount returns the DataTxMsgCount field if non-nil, zero value otherwise.

### GetDataTxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetDataTxMsgCountOk() (*int64, bool)`

GetDataTxMsgCountOk returns a tuple with the DataTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxMsgCount

`func (o *MsgVpnBridgeCounter) SetDataTxMsgCount(v int64)`

SetDataTxMsgCount sets DataTxMsgCount field to given value.

### HasDataTxMsgCount

`func (o *MsgVpnBridgeCounter) HasDataTxMsgCount() bool`

HasDataTxMsgCount returns a boolean if a field has been set.

### GetDiscardedRxMsgCount

`func (o *MsgVpnBridgeCounter) GetDiscardedRxMsgCount() int64`

GetDiscardedRxMsgCount returns the DiscardedRxMsgCount field if non-nil, zero value otherwise.

### GetDiscardedRxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetDiscardedRxMsgCountOk() (*int64, bool)`

GetDiscardedRxMsgCountOk returns a tuple with the DiscardedRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedRxMsgCount

`func (o *MsgVpnBridgeCounter) SetDiscardedRxMsgCount(v int64)`

SetDiscardedRxMsgCount sets DiscardedRxMsgCount field to given value.

### HasDiscardedRxMsgCount

`func (o *MsgVpnBridgeCounter) HasDiscardedRxMsgCount() bool`

HasDiscardedRxMsgCount returns a boolean if a field has been set.

### GetDiscardedTxMsgCount

`func (o *MsgVpnBridgeCounter) GetDiscardedTxMsgCount() int64`

GetDiscardedTxMsgCount returns the DiscardedTxMsgCount field if non-nil, zero value otherwise.

### GetDiscardedTxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetDiscardedTxMsgCountOk() (*int64, bool)`

GetDiscardedTxMsgCountOk returns a tuple with the DiscardedTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedTxMsgCount

`func (o *MsgVpnBridgeCounter) SetDiscardedTxMsgCount(v int64)`

SetDiscardedTxMsgCount sets DiscardedTxMsgCount field to given value.

### HasDiscardedTxMsgCount

`func (o *MsgVpnBridgeCounter) HasDiscardedTxMsgCount() bool`

HasDiscardedTxMsgCount returns a boolean if a field has been set.

### GetLoginRxMsgCount

`func (o *MsgVpnBridgeCounter) GetLoginRxMsgCount() int64`

GetLoginRxMsgCount returns the LoginRxMsgCount field if non-nil, zero value otherwise.

### GetLoginRxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetLoginRxMsgCountOk() (*int64, bool)`

GetLoginRxMsgCountOk returns a tuple with the LoginRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginRxMsgCount

`func (o *MsgVpnBridgeCounter) SetLoginRxMsgCount(v int64)`

SetLoginRxMsgCount sets LoginRxMsgCount field to given value.

### HasLoginRxMsgCount

`func (o *MsgVpnBridgeCounter) HasLoginRxMsgCount() bool`

HasLoginRxMsgCount returns a boolean if a field has been set.

### GetLoginTxMsgCount

`func (o *MsgVpnBridgeCounter) GetLoginTxMsgCount() int64`

GetLoginTxMsgCount returns the LoginTxMsgCount field if non-nil, zero value otherwise.

### GetLoginTxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetLoginTxMsgCountOk() (*int64, bool)`

GetLoginTxMsgCountOk returns a tuple with the LoginTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginTxMsgCount

`func (o *MsgVpnBridgeCounter) SetLoginTxMsgCount(v int64)`

SetLoginTxMsgCount sets LoginTxMsgCount field to given value.

### HasLoginTxMsgCount

`func (o *MsgVpnBridgeCounter) HasLoginTxMsgCount() bool`

HasLoginTxMsgCount returns a boolean if a field has been set.

### GetMsgSpoolRxMsgCount

`func (o *MsgVpnBridgeCounter) GetMsgSpoolRxMsgCount() int64`

GetMsgSpoolRxMsgCount returns the MsgSpoolRxMsgCount field if non-nil, zero value otherwise.

### GetMsgSpoolRxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetMsgSpoolRxMsgCountOk() (*int64, bool)`

GetMsgSpoolRxMsgCountOk returns a tuple with the MsgSpoolRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolRxMsgCount

`func (o *MsgVpnBridgeCounter) SetMsgSpoolRxMsgCount(v int64)`

SetMsgSpoolRxMsgCount sets MsgSpoolRxMsgCount field to given value.

### HasMsgSpoolRxMsgCount

`func (o *MsgVpnBridgeCounter) HasMsgSpoolRxMsgCount() bool`

HasMsgSpoolRxMsgCount returns a boolean if a field has been set.

### GetRxByteCount

`func (o *MsgVpnBridgeCounter) GetRxByteCount() int64`

GetRxByteCount returns the RxByteCount field if non-nil, zero value otherwise.

### GetRxByteCountOk

`func (o *MsgVpnBridgeCounter) GetRxByteCountOk() (*int64, bool)`

GetRxByteCountOk returns a tuple with the RxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteCount

`func (o *MsgVpnBridgeCounter) SetRxByteCount(v int64)`

SetRxByteCount sets RxByteCount field to given value.

### HasRxByteCount

`func (o *MsgVpnBridgeCounter) HasRxByteCount() bool`

HasRxByteCount returns a boolean if a field has been set.

### GetRxMsgCount

`func (o *MsgVpnBridgeCounter) GetRxMsgCount() int64`

GetRxMsgCount returns the RxMsgCount field if non-nil, zero value otherwise.

### GetRxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetRxMsgCountOk() (*int64, bool)`

GetRxMsgCountOk returns a tuple with the RxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgCount

`func (o *MsgVpnBridgeCounter) SetRxMsgCount(v int64)`

SetRxMsgCount sets RxMsgCount field to given value.

### HasRxMsgCount

`func (o *MsgVpnBridgeCounter) HasRxMsgCount() bool`

HasRxMsgCount returns a boolean if a field has been set.

### GetTxByteCount

`func (o *MsgVpnBridgeCounter) GetTxByteCount() int64`

GetTxByteCount returns the TxByteCount field if non-nil, zero value otherwise.

### GetTxByteCountOk

`func (o *MsgVpnBridgeCounter) GetTxByteCountOk() (*int64, bool)`

GetTxByteCountOk returns a tuple with the TxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteCount

`func (o *MsgVpnBridgeCounter) SetTxByteCount(v int64)`

SetTxByteCount sets TxByteCount field to given value.

### HasTxByteCount

`func (o *MsgVpnBridgeCounter) HasTxByteCount() bool`

HasTxByteCount returns a boolean if a field has been set.

### GetTxMsgCount

`func (o *MsgVpnBridgeCounter) GetTxMsgCount() int64`

GetTxMsgCount returns the TxMsgCount field if non-nil, zero value otherwise.

### GetTxMsgCountOk

`func (o *MsgVpnBridgeCounter) GetTxMsgCountOk() (*int64, bool)`

GetTxMsgCountOk returns a tuple with the TxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgCount

`func (o *MsgVpnBridgeCounter) SetTxMsgCount(v int64)`

SetTxMsgCount sets TxMsgCount field to given value.

### HasTxMsgCount

`func (o *MsgVpnBridgeCounter) HasTxMsgCount() bool`

HasTxMsgCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


