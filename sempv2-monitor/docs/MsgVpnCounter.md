# MsgVpnCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlRxByteCount** | Pointer to **int64** | The amount of client control messages received from clients by the Message VPN, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**ControlRxMsgCount** | Pointer to **int64** | The number of client control messages received from clients by the Message VPN. Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**ControlTxByteCount** | Pointer to **int64** | The amount of client control messages transmitted to clients by the Message VPN, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**ControlTxMsgCount** | Pointer to **int64** | The number of client control messages transmitted to clients by the Message VPN. Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**DataRxByteCount** | Pointer to **int64** | The amount of client data messages received from clients by the Message VPN, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**DataRxMsgCount** | Pointer to **int64** | The number of client data messages received from clients by the Message VPN. Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**DataTxByteCount** | Pointer to **int64** | The amount of client data messages transmitted to clients by the Message VPN, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**DataTxMsgCount** | Pointer to **int64** | The number of client data messages transmitted to clients by the Message VPN. Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**DiscardedRxMsgCount** | Pointer to **int64** | The number of messages discarded during reception by the Message VPN. Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**DiscardedTxMsgCount** | Pointer to **int64** | The number of messages discarded during transmission by the Message VPN. Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**LoginRxMsgCount** | Pointer to **int64** | The number of login request messages received by the Message VPN. Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**LoginTxMsgCount** | Pointer to **int64** | The number of login response messages transmitted by the Message VPN. Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**MsgSpoolRxMsgCount** | Pointer to **int64** | The number of guaranteed messages received by the Message VPN. Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**MsgSpoolTxMsgCount** | Pointer to **int64** | The number of guaranteed messages transmitted by the Message VPN. One message to multiple clients is counted as one message. Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**TlsRxByteCount** | Pointer to **int64** | The amount of TLS messages received by the Message VPN, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 
**TlsTxByteCount** | Pointer to **int64** | The amount of TLS messages transmitted by the Message VPN, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpn object. | [optional] 

## Methods

### NewMsgVpnCounter

`func NewMsgVpnCounter() *MsgVpnCounter`

NewMsgVpnCounter instantiates a new MsgVpnCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnCounterWithDefaults

`func NewMsgVpnCounterWithDefaults() *MsgVpnCounter`

NewMsgVpnCounterWithDefaults instantiates a new MsgVpnCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlRxByteCount

`func (o *MsgVpnCounter) GetControlRxByteCount() int64`

GetControlRxByteCount returns the ControlRxByteCount field if non-nil, zero value otherwise.

### GetControlRxByteCountOk

`func (o *MsgVpnCounter) GetControlRxByteCountOk() (*int64, bool)`

GetControlRxByteCountOk returns a tuple with the ControlRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRxByteCount

`func (o *MsgVpnCounter) SetControlRxByteCount(v int64)`

SetControlRxByteCount sets ControlRxByteCount field to given value.

### HasControlRxByteCount

`func (o *MsgVpnCounter) HasControlRxByteCount() bool`

HasControlRxByteCount returns a boolean if a field has been set.

### GetControlRxMsgCount

`func (o *MsgVpnCounter) GetControlRxMsgCount() int64`

GetControlRxMsgCount returns the ControlRxMsgCount field if non-nil, zero value otherwise.

### GetControlRxMsgCountOk

`func (o *MsgVpnCounter) GetControlRxMsgCountOk() (*int64, bool)`

GetControlRxMsgCountOk returns a tuple with the ControlRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRxMsgCount

`func (o *MsgVpnCounter) SetControlRxMsgCount(v int64)`

SetControlRxMsgCount sets ControlRxMsgCount field to given value.

### HasControlRxMsgCount

`func (o *MsgVpnCounter) HasControlRxMsgCount() bool`

HasControlRxMsgCount returns a boolean if a field has been set.

### GetControlTxByteCount

`func (o *MsgVpnCounter) GetControlTxByteCount() int64`

GetControlTxByteCount returns the ControlTxByteCount field if non-nil, zero value otherwise.

### GetControlTxByteCountOk

`func (o *MsgVpnCounter) GetControlTxByteCountOk() (*int64, bool)`

GetControlTxByteCountOk returns a tuple with the ControlTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTxByteCount

`func (o *MsgVpnCounter) SetControlTxByteCount(v int64)`

SetControlTxByteCount sets ControlTxByteCount field to given value.

### HasControlTxByteCount

`func (o *MsgVpnCounter) HasControlTxByteCount() bool`

HasControlTxByteCount returns a boolean if a field has been set.

### GetControlTxMsgCount

`func (o *MsgVpnCounter) GetControlTxMsgCount() int64`

GetControlTxMsgCount returns the ControlTxMsgCount field if non-nil, zero value otherwise.

### GetControlTxMsgCountOk

`func (o *MsgVpnCounter) GetControlTxMsgCountOk() (*int64, bool)`

GetControlTxMsgCountOk returns a tuple with the ControlTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTxMsgCount

`func (o *MsgVpnCounter) SetControlTxMsgCount(v int64)`

SetControlTxMsgCount sets ControlTxMsgCount field to given value.

### HasControlTxMsgCount

`func (o *MsgVpnCounter) HasControlTxMsgCount() bool`

HasControlTxMsgCount returns a boolean if a field has been set.

### GetDataRxByteCount

`func (o *MsgVpnCounter) GetDataRxByteCount() int64`

GetDataRxByteCount returns the DataRxByteCount field if non-nil, zero value otherwise.

### GetDataRxByteCountOk

`func (o *MsgVpnCounter) GetDataRxByteCountOk() (*int64, bool)`

GetDataRxByteCountOk returns a tuple with the DataRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxByteCount

`func (o *MsgVpnCounter) SetDataRxByteCount(v int64)`

SetDataRxByteCount sets DataRxByteCount field to given value.

### HasDataRxByteCount

`func (o *MsgVpnCounter) HasDataRxByteCount() bool`

HasDataRxByteCount returns a boolean if a field has been set.

### GetDataRxMsgCount

`func (o *MsgVpnCounter) GetDataRxMsgCount() int64`

GetDataRxMsgCount returns the DataRxMsgCount field if non-nil, zero value otherwise.

### GetDataRxMsgCountOk

`func (o *MsgVpnCounter) GetDataRxMsgCountOk() (*int64, bool)`

GetDataRxMsgCountOk returns a tuple with the DataRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxMsgCount

`func (o *MsgVpnCounter) SetDataRxMsgCount(v int64)`

SetDataRxMsgCount sets DataRxMsgCount field to given value.

### HasDataRxMsgCount

`func (o *MsgVpnCounter) HasDataRxMsgCount() bool`

HasDataRxMsgCount returns a boolean if a field has been set.

### GetDataTxByteCount

`func (o *MsgVpnCounter) GetDataTxByteCount() int64`

GetDataTxByteCount returns the DataTxByteCount field if non-nil, zero value otherwise.

### GetDataTxByteCountOk

`func (o *MsgVpnCounter) GetDataTxByteCountOk() (*int64, bool)`

GetDataTxByteCountOk returns a tuple with the DataTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxByteCount

`func (o *MsgVpnCounter) SetDataTxByteCount(v int64)`

SetDataTxByteCount sets DataTxByteCount field to given value.

### HasDataTxByteCount

`func (o *MsgVpnCounter) HasDataTxByteCount() bool`

HasDataTxByteCount returns a boolean if a field has been set.

### GetDataTxMsgCount

`func (o *MsgVpnCounter) GetDataTxMsgCount() int64`

GetDataTxMsgCount returns the DataTxMsgCount field if non-nil, zero value otherwise.

### GetDataTxMsgCountOk

`func (o *MsgVpnCounter) GetDataTxMsgCountOk() (*int64, bool)`

GetDataTxMsgCountOk returns a tuple with the DataTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxMsgCount

`func (o *MsgVpnCounter) SetDataTxMsgCount(v int64)`

SetDataTxMsgCount sets DataTxMsgCount field to given value.

### HasDataTxMsgCount

`func (o *MsgVpnCounter) HasDataTxMsgCount() bool`

HasDataTxMsgCount returns a boolean if a field has been set.

### GetDiscardedRxMsgCount

`func (o *MsgVpnCounter) GetDiscardedRxMsgCount() int64`

GetDiscardedRxMsgCount returns the DiscardedRxMsgCount field if non-nil, zero value otherwise.

### GetDiscardedRxMsgCountOk

`func (o *MsgVpnCounter) GetDiscardedRxMsgCountOk() (*int64, bool)`

GetDiscardedRxMsgCountOk returns a tuple with the DiscardedRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedRxMsgCount

`func (o *MsgVpnCounter) SetDiscardedRxMsgCount(v int64)`

SetDiscardedRxMsgCount sets DiscardedRxMsgCount field to given value.

### HasDiscardedRxMsgCount

`func (o *MsgVpnCounter) HasDiscardedRxMsgCount() bool`

HasDiscardedRxMsgCount returns a boolean if a field has been set.

### GetDiscardedTxMsgCount

`func (o *MsgVpnCounter) GetDiscardedTxMsgCount() int64`

GetDiscardedTxMsgCount returns the DiscardedTxMsgCount field if non-nil, zero value otherwise.

### GetDiscardedTxMsgCountOk

`func (o *MsgVpnCounter) GetDiscardedTxMsgCountOk() (*int64, bool)`

GetDiscardedTxMsgCountOk returns a tuple with the DiscardedTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedTxMsgCount

`func (o *MsgVpnCounter) SetDiscardedTxMsgCount(v int64)`

SetDiscardedTxMsgCount sets DiscardedTxMsgCount field to given value.

### HasDiscardedTxMsgCount

`func (o *MsgVpnCounter) HasDiscardedTxMsgCount() bool`

HasDiscardedTxMsgCount returns a boolean if a field has been set.

### GetLoginRxMsgCount

`func (o *MsgVpnCounter) GetLoginRxMsgCount() int64`

GetLoginRxMsgCount returns the LoginRxMsgCount field if non-nil, zero value otherwise.

### GetLoginRxMsgCountOk

`func (o *MsgVpnCounter) GetLoginRxMsgCountOk() (*int64, bool)`

GetLoginRxMsgCountOk returns a tuple with the LoginRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginRxMsgCount

`func (o *MsgVpnCounter) SetLoginRxMsgCount(v int64)`

SetLoginRxMsgCount sets LoginRxMsgCount field to given value.

### HasLoginRxMsgCount

`func (o *MsgVpnCounter) HasLoginRxMsgCount() bool`

HasLoginRxMsgCount returns a boolean if a field has been set.

### GetLoginTxMsgCount

`func (o *MsgVpnCounter) GetLoginTxMsgCount() int64`

GetLoginTxMsgCount returns the LoginTxMsgCount field if non-nil, zero value otherwise.

### GetLoginTxMsgCountOk

`func (o *MsgVpnCounter) GetLoginTxMsgCountOk() (*int64, bool)`

GetLoginTxMsgCountOk returns a tuple with the LoginTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginTxMsgCount

`func (o *MsgVpnCounter) SetLoginTxMsgCount(v int64)`

SetLoginTxMsgCount sets LoginTxMsgCount field to given value.

### HasLoginTxMsgCount

`func (o *MsgVpnCounter) HasLoginTxMsgCount() bool`

HasLoginTxMsgCount returns a boolean if a field has been set.

### GetMsgSpoolRxMsgCount

`func (o *MsgVpnCounter) GetMsgSpoolRxMsgCount() int64`

GetMsgSpoolRxMsgCount returns the MsgSpoolRxMsgCount field if non-nil, zero value otherwise.

### GetMsgSpoolRxMsgCountOk

`func (o *MsgVpnCounter) GetMsgSpoolRxMsgCountOk() (*int64, bool)`

GetMsgSpoolRxMsgCountOk returns a tuple with the MsgSpoolRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolRxMsgCount

`func (o *MsgVpnCounter) SetMsgSpoolRxMsgCount(v int64)`

SetMsgSpoolRxMsgCount sets MsgSpoolRxMsgCount field to given value.

### HasMsgSpoolRxMsgCount

`func (o *MsgVpnCounter) HasMsgSpoolRxMsgCount() bool`

HasMsgSpoolRxMsgCount returns a boolean if a field has been set.

### GetMsgSpoolTxMsgCount

`func (o *MsgVpnCounter) GetMsgSpoolTxMsgCount() int64`

GetMsgSpoolTxMsgCount returns the MsgSpoolTxMsgCount field if non-nil, zero value otherwise.

### GetMsgSpoolTxMsgCountOk

`func (o *MsgVpnCounter) GetMsgSpoolTxMsgCountOk() (*int64, bool)`

GetMsgSpoolTxMsgCountOk returns a tuple with the MsgSpoolTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolTxMsgCount

`func (o *MsgVpnCounter) SetMsgSpoolTxMsgCount(v int64)`

SetMsgSpoolTxMsgCount sets MsgSpoolTxMsgCount field to given value.

### HasMsgSpoolTxMsgCount

`func (o *MsgVpnCounter) HasMsgSpoolTxMsgCount() bool`

HasMsgSpoolTxMsgCount returns a boolean if a field has been set.

### GetTlsRxByteCount

`func (o *MsgVpnCounter) GetTlsRxByteCount() int64`

GetTlsRxByteCount returns the TlsRxByteCount field if non-nil, zero value otherwise.

### GetTlsRxByteCountOk

`func (o *MsgVpnCounter) GetTlsRxByteCountOk() (*int64, bool)`

GetTlsRxByteCountOk returns a tuple with the TlsRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsRxByteCount

`func (o *MsgVpnCounter) SetTlsRxByteCount(v int64)`

SetTlsRxByteCount sets TlsRxByteCount field to given value.

### HasTlsRxByteCount

`func (o *MsgVpnCounter) HasTlsRxByteCount() bool`

HasTlsRxByteCount returns a boolean if a field has been set.

### GetTlsTxByteCount

`func (o *MsgVpnCounter) GetTlsTxByteCount() int64`

GetTlsTxByteCount returns the TlsTxByteCount field if non-nil, zero value otherwise.

### GetTlsTxByteCountOk

`func (o *MsgVpnCounter) GetTlsTxByteCountOk() (*int64, bool)`

GetTlsTxByteCountOk returns a tuple with the TlsTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTxByteCount

`func (o *MsgVpnCounter) SetTlsTxByteCount(v int64)`

SetTlsTxByteCount sets TlsTxByteCount field to given value.

### HasTlsTxByteCount

`func (o *MsgVpnCounter) HasTlsTxByteCount() bool`

HasTlsTxByteCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


