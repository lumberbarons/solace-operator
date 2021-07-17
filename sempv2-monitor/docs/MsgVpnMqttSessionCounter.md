# MsgVpnMqttSessionCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MqttConnackErrorTxCount** | Pointer to **int64** | The number of MQTT connect acknowledgment (CONNACK) refused response packets transmitted to the Client. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttConnackTxCount** | Pointer to **int64** | The number of MQTT connect acknowledgment (CONNACK) accepted response packets transmitted to the Client. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttConnectRxCount** | Pointer to **int64** | The number of MQTT connect (CONNECT) request packets received from the Client. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttDisconnectRxCount** | Pointer to **int64** | The number of MQTT disconnect (DISCONNECT) request packets received from the Client. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttPubcompTxCount** | Pointer to **int64** | The number of MQTT publish complete (PUBCOMP) packets transmitted to the Client in response to a PUBREL packet. These packets are the fourth and final packet of a QoS 2 protocol exchange. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttPublishQos0RxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 0 message delivery. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttPublishQos0TxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 0 message delivery. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttPublishQos1RxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 1 message delivery. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttPublishQos1TxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 1 message delivery. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttPublishQos2RxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 2 message delivery. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttPubrecTxCount** | Pointer to **int64** | The number of MQTT publish received (PUBREC) packets transmitted to the Client in response to a PUBLISH packet with QoS 2. These packets are the second packet of a QoS 2 protocol exchange. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 
**MqttPubrelRxCount** | Pointer to **int64** | The number of MQTT publish release (PUBREL) packets received from the Client in response to a PUBREC packet. These packets are the third packet of a QoS 2 protocol exchange. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object. | [optional] 

## Methods

### NewMsgVpnMqttSessionCounter

`func NewMsgVpnMqttSessionCounter() *MsgVpnMqttSessionCounter`

NewMsgVpnMqttSessionCounter instantiates a new MsgVpnMqttSessionCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnMqttSessionCounterWithDefaults

`func NewMsgVpnMqttSessionCounterWithDefaults() *MsgVpnMqttSessionCounter`

NewMsgVpnMqttSessionCounterWithDefaults instantiates a new MsgVpnMqttSessionCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMqttConnackErrorTxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttConnackErrorTxCount() int64`

GetMqttConnackErrorTxCount returns the MqttConnackErrorTxCount field if non-nil, zero value otherwise.

### GetMqttConnackErrorTxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttConnackErrorTxCountOk() (*int64, bool)`

GetMqttConnackErrorTxCountOk returns a tuple with the MqttConnackErrorTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConnackErrorTxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttConnackErrorTxCount(v int64)`

SetMqttConnackErrorTxCount sets MqttConnackErrorTxCount field to given value.

### HasMqttConnackErrorTxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttConnackErrorTxCount() bool`

HasMqttConnackErrorTxCount returns a boolean if a field has been set.

### GetMqttConnackTxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttConnackTxCount() int64`

GetMqttConnackTxCount returns the MqttConnackTxCount field if non-nil, zero value otherwise.

### GetMqttConnackTxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttConnackTxCountOk() (*int64, bool)`

GetMqttConnackTxCountOk returns a tuple with the MqttConnackTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConnackTxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttConnackTxCount(v int64)`

SetMqttConnackTxCount sets MqttConnackTxCount field to given value.

### HasMqttConnackTxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttConnackTxCount() bool`

HasMqttConnackTxCount returns a boolean if a field has been set.

### GetMqttConnectRxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttConnectRxCount() int64`

GetMqttConnectRxCount returns the MqttConnectRxCount field if non-nil, zero value otherwise.

### GetMqttConnectRxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttConnectRxCountOk() (*int64, bool)`

GetMqttConnectRxCountOk returns a tuple with the MqttConnectRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConnectRxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttConnectRxCount(v int64)`

SetMqttConnectRxCount sets MqttConnectRxCount field to given value.

### HasMqttConnectRxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttConnectRxCount() bool`

HasMqttConnectRxCount returns a boolean if a field has been set.

### GetMqttDisconnectRxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttDisconnectRxCount() int64`

GetMqttDisconnectRxCount returns the MqttDisconnectRxCount field if non-nil, zero value otherwise.

### GetMqttDisconnectRxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttDisconnectRxCountOk() (*int64, bool)`

GetMqttDisconnectRxCountOk returns a tuple with the MqttDisconnectRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttDisconnectRxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttDisconnectRxCount(v int64)`

SetMqttDisconnectRxCount sets MqttDisconnectRxCount field to given value.

### HasMqttDisconnectRxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttDisconnectRxCount() bool`

HasMqttDisconnectRxCount returns a boolean if a field has been set.

### GetMqttPubcompTxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttPubcompTxCount() int64`

GetMqttPubcompTxCount returns the MqttPubcompTxCount field if non-nil, zero value otherwise.

### GetMqttPubcompTxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttPubcompTxCountOk() (*int64, bool)`

GetMqttPubcompTxCountOk returns a tuple with the MqttPubcompTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubcompTxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttPubcompTxCount(v int64)`

SetMqttPubcompTxCount sets MqttPubcompTxCount field to given value.

### HasMqttPubcompTxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttPubcompTxCount() bool`

HasMqttPubcompTxCount returns a boolean if a field has been set.

### GetMqttPublishQos0RxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos0RxCount() int64`

GetMqttPublishQos0RxCount returns the MqttPublishQos0RxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos0RxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos0RxCountOk() (*int64, bool)`

GetMqttPublishQos0RxCountOk returns a tuple with the MqttPublishQos0RxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos0RxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttPublishQos0RxCount(v int64)`

SetMqttPublishQos0RxCount sets MqttPublishQos0RxCount field to given value.

### HasMqttPublishQos0RxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttPublishQos0RxCount() bool`

HasMqttPublishQos0RxCount returns a boolean if a field has been set.

### GetMqttPublishQos0TxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos0TxCount() int64`

GetMqttPublishQos0TxCount returns the MqttPublishQos0TxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos0TxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos0TxCountOk() (*int64, bool)`

GetMqttPublishQos0TxCountOk returns a tuple with the MqttPublishQos0TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos0TxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttPublishQos0TxCount(v int64)`

SetMqttPublishQos0TxCount sets MqttPublishQos0TxCount field to given value.

### HasMqttPublishQos0TxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttPublishQos0TxCount() bool`

HasMqttPublishQos0TxCount returns a boolean if a field has been set.

### GetMqttPublishQos1RxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos1RxCount() int64`

GetMqttPublishQos1RxCount returns the MqttPublishQos1RxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos1RxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos1RxCountOk() (*int64, bool)`

GetMqttPublishQos1RxCountOk returns a tuple with the MqttPublishQos1RxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos1RxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttPublishQos1RxCount(v int64)`

SetMqttPublishQos1RxCount sets MqttPublishQos1RxCount field to given value.

### HasMqttPublishQos1RxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttPublishQos1RxCount() bool`

HasMqttPublishQos1RxCount returns a boolean if a field has been set.

### GetMqttPublishQos1TxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos1TxCount() int64`

GetMqttPublishQos1TxCount returns the MqttPublishQos1TxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos1TxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos1TxCountOk() (*int64, bool)`

GetMqttPublishQos1TxCountOk returns a tuple with the MqttPublishQos1TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos1TxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttPublishQos1TxCount(v int64)`

SetMqttPublishQos1TxCount sets MqttPublishQos1TxCount field to given value.

### HasMqttPublishQos1TxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttPublishQos1TxCount() bool`

HasMqttPublishQos1TxCount returns a boolean if a field has been set.

### GetMqttPublishQos2RxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos2RxCount() int64`

GetMqttPublishQos2RxCount returns the MqttPublishQos2RxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos2RxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos2RxCountOk() (*int64, bool)`

GetMqttPublishQos2RxCountOk returns a tuple with the MqttPublishQos2RxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos2RxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttPublishQos2RxCount(v int64)`

SetMqttPublishQos2RxCount sets MqttPublishQos2RxCount field to given value.

### HasMqttPublishQos2RxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttPublishQos2RxCount() bool`

HasMqttPublishQos2RxCount returns a boolean if a field has been set.

### GetMqttPubrecTxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttPubrecTxCount() int64`

GetMqttPubrecTxCount returns the MqttPubrecTxCount field if non-nil, zero value otherwise.

### GetMqttPubrecTxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttPubrecTxCountOk() (*int64, bool)`

GetMqttPubrecTxCountOk returns a tuple with the MqttPubrecTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubrecTxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttPubrecTxCount(v int64)`

SetMqttPubrecTxCount sets MqttPubrecTxCount field to given value.

### HasMqttPubrecTxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttPubrecTxCount() bool`

HasMqttPubrecTxCount returns a boolean if a field has been set.

### GetMqttPubrelRxCount

`func (o *MsgVpnMqttSessionCounter) GetMqttPubrelRxCount() int64`

GetMqttPubrelRxCount returns the MqttPubrelRxCount field if non-nil, zero value otherwise.

### GetMqttPubrelRxCountOk

`func (o *MsgVpnMqttSessionCounter) GetMqttPubrelRxCountOk() (*int64, bool)`

GetMqttPubrelRxCountOk returns a tuple with the MqttPubrelRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubrelRxCount

`func (o *MsgVpnMqttSessionCounter) SetMqttPubrelRxCount(v int64)`

SetMqttPubrelRxCount sets MqttPubrelRxCount field to given value.

### HasMqttPubrelRxCount

`func (o *MsgVpnMqttSessionCounter) HasMqttPubrelRxCount() bool`

HasMqttPubrelRxCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


