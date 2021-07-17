# MsgVpnMqttSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clean** | Pointer to **bool** | Indicates whether the Client requested a clean (newly created) MQTT Session when connecting. If not clean (already existing), then previously stored messages for QoS 1 subscriptions are delivered. | [optional] 
**ClientName** | Pointer to **string** | The name of the MQTT Session Client. | [optional] 
**Counter** | Pointer to [**MsgVpnMqttSessionCounter**](MsgVpnMqttSessionCounter.md) |  | [optional] 
**CreatedByManagement** | Pointer to **bool** | Indicates whether the MQTT Session was created by a Management API. | [optional] 
**Durable** | Pointer to **bool** | Indicates whether the MQTT Session is durable. Disconnected durable MQTT Sessions are deleted when their expiry time is reached. Disconnected non-durable MQTT Sessions are deleted immediately. Available since 2.21. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the MQTT Session is enabled. | [optional] 
**ExpiryTime** | Pointer to **int64** | The timestamp of when the disconnected MQTT session expires and is deleted. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). A value of 0 indicates that the session is either connected, or will never expire. Available since 2.21. | [optional] 
**MaxPacketSize** | Pointer to **int64** | The maximum size of a packet, including all headers and payload, that the Client has signaled it is willing to accept. A value of zero indicates no limit. Note that there are other broker settings which may further limit packet size. Available since 2.21. | [optional] 
**MqttConnackErrorTxCount** | Pointer to **int64** | The number of MQTT connect acknowledgment (CONNACK) refused response packets transmitted to the Client. Available since 2.13. | [optional] 
**MqttConnackTxCount** | Pointer to **int64** | The number of MQTT connect acknowledgment (CONNACK) accepted response packets transmitted to the Client. Available since 2.13. | [optional] 
**MqttConnectRxCount** | Pointer to **int64** | The number of MQTT connect (CONNECT) request packets received from the Client. Available since 2.13. | [optional] 
**MqttDisconnectRxCount** | Pointer to **int64** | The number of MQTT disconnect (DISCONNECT) request packets received from the Client. Available since 2.13. | [optional] 
**MqttPubcompTxCount** | Pointer to **int64** | The number of MQTT publish complete (PUBCOMP) packets transmitted to the Client in response to a PUBREL packet. These packets are the fourth and final packet of a QoS 2 protocol exchange. Available since 2.13. | [optional] 
**MqttPublishQos0RxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 0 message delivery. Available since 2.13. | [optional] 
**MqttPublishQos0TxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 0 message delivery. Available since 2.13. | [optional] 
**MqttPublishQos1RxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 1 message delivery. Available since 2.13. | [optional] 
**MqttPublishQos1TxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 1 message delivery. Available since 2.13. | [optional] 
**MqttPublishQos2RxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 2 message delivery. Available since 2.13. | [optional] 
**MqttPubrecTxCount** | Pointer to **int64** | The number of MQTT publish received (PUBREC) packets transmitted to the Client in response to a PUBLISH packet with QoS 2. These packets are the second packet of a QoS 2 protocol exchange. Available since 2.13. | [optional] 
**MqttPubrelRxCount** | Pointer to **int64** | The number of MQTT publish release (PUBREL) packets received from the Client in response to a PUBREC packet. These packets are the third packet of a QoS 2 protocol exchange. Available since 2.13. | [optional] 
**MqttSessionClientId** | Pointer to **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | [optional] 
**MqttSessionVirtualRouter** | Pointer to **string** | The virtual router of the MQTT Session. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The MQTT Session belongs to the primary virtual router. \&quot;backup\&quot; - The MQTT Session belongs to the backup virtual router. &lt;/pre&gt;  | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Owner** | Pointer to **string** | The Client Username which owns the MQTT Session. | [optional] 
**QueueConsumerAckPropagationEnabled** | Pointer to **bool** | Indicates whether consumer acknowledgements (ACKs) received on the active replication Message VPN are propagated to the standby replication Message VPN. Available since 2.14. | [optional] 
**QueueDeadMsgQueue** | Pointer to **string** | The name of the Dead Message Queue (DMQ) used by the MQTT Session Queue. Available since 2.14. | [optional] 
**QueueEventBindCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**QueueEventMsgSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**QueueEventRejectLowPriorityMsgLimitThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**QueueMaxBindCount** | Pointer to **int64** | The maximum number of consumer flows that can bind to the MQTT Session Queue. Available since 2.14. | [optional] 
**QueueMaxDeliveredUnackedMsgsPerFlow** | Pointer to **int64** | The maximum number of messages delivered but not acknowledged per flow for the MQTT Session Queue. Available since 2.14. | [optional] 
**QueueMaxMsgSize** | Pointer to **int32** | The maximum message size allowed in the MQTT Session Queue, in bytes (B). Available since 2.14. | [optional] 
**QueueMaxMsgSpoolUsage** | Pointer to **int64** | The maximum message spool usage allowed by the MQTT Session Queue, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. Available since 2.14. | [optional] 
**QueueMaxRedeliveryCount** | Pointer to **int64** | The maximum number of times the MQTT Session Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. Available since 2.14. | [optional] 
**QueueMaxTtl** | Pointer to **int64** | The maximum time in seconds a message can stay in the MQTT Session Queue when &#x60;queueRespectTtlEnabled&#x60; is &#x60;\&quot;true\&quot;&#x60;. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the &#x60;queueMaxTtl&#x60; configured for the MQTT Session Queue, is exceeded. A value of 0 disables expiry. Available since 2.14. | [optional] 
**QueueName** | Pointer to **string** | The name of the MQTT Session Queue. | [optional] 
**QueueRejectLowPriorityMsgEnabled** | Pointer to **bool** | Indicates whether to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. Available since 2.14. | [optional] 
**QueueRejectLowPriorityMsgLimit** | Pointer to **int64** | The number of messages of any priority in the MQTT Session Queue above which low priority messages are not admitted but higher priority messages are allowed. Available since 2.14. | [optional] 
**QueueRejectMsgToSenderOnDiscardBehavior** | Pointer to **string** | Indicates whether negative acknowledgements (NACKs) are returned to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always return a negative acknowledgment (NACK) to the sending client on message discard. \&quot;when-queue-enabled\&quot; - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Queue is enabled. \&quot;never\&quot; - Never return a negative acknowledgment (NACK) to the sending client on message discard. &lt;/pre&gt;  Available since 2.14. | [optional] 
**QueueRespectTtlEnabled** | Pointer to **bool** | Indicates whether the time-to-live (TTL) for messages in the MQTT Session Queue is respected. When enabled, expired messages are discarded or moved to the DMQ. Available since 2.14. | [optional] 
**RxMax** | Pointer to **int64** | The maximum number of outstanding QoS1 and QoS2 messages that the Client has signaled it is willing to accept. Note that there are other broker settings which may further limit the number of outstanding messasges. Available since 2.21. | [optional] 
**Will** | Pointer to **bool** | Indicates whether the MQTT Session has the Will message specified by the Client. The Will message is published if the Client disconnects without sending the MQTT DISCONNECT packet. | [optional] 

## Methods

### NewMsgVpnMqttSession

`func NewMsgVpnMqttSession() *MsgVpnMqttSession`

NewMsgVpnMqttSession instantiates a new MsgVpnMqttSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnMqttSessionWithDefaults

`func NewMsgVpnMqttSessionWithDefaults() *MsgVpnMqttSession`

NewMsgVpnMqttSessionWithDefaults instantiates a new MsgVpnMqttSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClean

`func (o *MsgVpnMqttSession) GetClean() bool`

GetClean returns the Clean field if non-nil, zero value otherwise.

### GetCleanOk

`func (o *MsgVpnMqttSession) GetCleanOk() (*bool, bool)`

GetCleanOk returns a tuple with the Clean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClean

`func (o *MsgVpnMqttSession) SetClean(v bool)`

SetClean sets Clean field to given value.

### HasClean

`func (o *MsgVpnMqttSession) HasClean() bool`

HasClean returns a boolean if a field has been set.

### GetClientName

`func (o *MsgVpnMqttSession) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnMqttSession) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnMqttSession) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnMqttSession) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetCounter

`func (o *MsgVpnMqttSession) GetCounter() MsgVpnMqttSessionCounter`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *MsgVpnMqttSession) GetCounterOk() (*MsgVpnMqttSessionCounter, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *MsgVpnMqttSession) SetCounter(v MsgVpnMqttSessionCounter)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *MsgVpnMqttSession) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetCreatedByManagement

`func (o *MsgVpnMqttSession) GetCreatedByManagement() bool`

GetCreatedByManagement returns the CreatedByManagement field if non-nil, zero value otherwise.

### GetCreatedByManagementOk

`func (o *MsgVpnMqttSession) GetCreatedByManagementOk() (*bool, bool)`

GetCreatedByManagementOk returns a tuple with the CreatedByManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByManagement

`func (o *MsgVpnMqttSession) SetCreatedByManagement(v bool)`

SetCreatedByManagement sets CreatedByManagement field to given value.

### HasCreatedByManagement

`func (o *MsgVpnMqttSession) HasCreatedByManagement() bool`

HasCreatedByManagement returns a boolean if a field has been set.

### GetDurable

`func (o *MsgVpnMqttSession) GetDurable() bool`

GetDurable returns the Durable field if non-nil, zero value otherwise.

### GetDurableOk

`func (o *MsgVpnMqttSession) GetDurableOk() (*bool, bool)`

GetDurableOk returns a tuple with the Durable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurable

`func (o *MsgVpnMqttSession) SetDurable(v bool)`

SetDurable sets Durable field to given value.

### HasDurable

`func (o *MsgVpnMqttSession) HasDurable() bool`

HasDurable returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnMqttSession) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnMqttSession) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnMqttSession) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnMqttSession) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpiryTime

`func (o *MsgVpnMqttSession) GetExpiryTime() int64`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *MsgVpnMqttSession) GetExpiryTimeOk() (*int64, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *MsgVpnMqttSession) SetExpiryTime(v int64)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *MsgVpnMqttSession) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetMaxPacketSize

`func (o *MsgVpnMqttSession) GetMaxPacketSize() int64`

GetMaxPacketSize returns the MaxPacketSize field if non-nil, zero value otherwise.

### GetMaxPacketSizeOk

`func (o *MsgVpnMqttSession) GetMaxPacketSizeOk() (*int64, bool)`

GetMaxPacketSizeOk returns a tuple with the MaxPacketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketSize

`func (o *MsgVpnMqttSession) SetMaxPacketSize(v int64)`

SetMaxPacketSize sets MaxPacketSize field to given value.

### HasMaxPacketSize

`func (o *MsgVpnMqttSession) HasMaxPacketSize() bool`

HasMaxPacketSize returns a boolean if a field has been set.

### GetMqttConnackErrorTxCount

`func (o *MsgVpnMqttSession) GetMqttConnackErrorTxCount() int64`

GetMqttConnackErrorTxCount returns the MqttConnackErrorTxCount field if non-nil, zero value otherwise.

### GetMqttConnackErrorTxCountOk

`func (o *MsgVpnMqttSession) GetMqttConnackErrorTxCountOk() (*int64, bool)`

GetMqttConnackErrorTxCountOk returns a tuple with the MqttConnackErrorTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConnackErrorTxCount

`func (o *MsgVpnMqttSession) SetMqttConnackErrorTxCount(v int64)`

SetMqttConnackErrorTxCount sets MqttConnackErrorTxCount field to given value.

### HasMqttConnackErrorTxCount

`func (o *MsgVpnMqttSession) HasMqttConnackErrorTxCount() bool`

HasMqttConnackErrorTxCount returns a boolean if a field has been set.

### GetMqttConnackTxCount

`func (o *MsgVpnMqttSession) GetMqttConnackTxCount() int64`

GetMqttConnackTxCount returns the MqttConnackTxCount field if non-nil, zero value otherwise.

### GetMqttConnackTxCountOk

`func (o *MsgVpnMqttSession) GetMqttConnackTxCountOk() (*int64, bool)`

GetMqttConnackTxCountOk returns a tuple with the MqttConnackTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConnackTxCount

`func (o *MsgVpnMqttSession) SetMqttConnackTxCount(v int64)`

SetMqttConnackTxCount sets MqttConnackTxCount field to given value.

### HasMqttConnackTxCount

`func (o *MsgVpnMqttSession) HasMqttConnackTxCount() bool`

HasMqttConnackTxCount returns a boolean if a field has been set.

### GetMqttConnectRxCount

`func (o *MsgVpnMqttSession) GetMqttConnectRxCount() int64`

GetMqttConnectRxCount returns the MqttConnectRxCount field if non-nil, zero value otherwise.

### GetMqttConnectRxCountOk

`func (o *MsgVpnMqttSession) GetMqttConnectRxCountOk() (*int64, bool)`

GetMqttConnectRxCountOk returns a tuple with the MqttConnectRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConnectRxCount

`func (o *MsgVpnMqttSession) SetMqttConnectRxCount(v int64)`

SetMqttConnectRxCount sets MqttConnectRxCount field to given value.

### HasMqttConnectRxCount

`func (o *MsgVpnMqttSession) HasMqttConnectRxCount() bool`

HasMqttConnectRxCount returns a boolean if a field has been set.

### GetMqttDisconnectRxCount

`func (o *MsgVpnMqttSession) GetMqttDisconnectRxCount() int64`

GetMqttDisconnectRxCount returns the MqttDisconnectRxCount field if non-nil, zero value otherwise.

### GetMqttDisconnectRxCountOk

`func (o *MsgVpnMqttSession) GetMqttDisconnectRxCountOk() (*int64, bool)`

GetMqttDisconnectRxCountOk returns a tuple with the MqttDisconnectRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttDisconnectRxCount

`func (o *MsgVpnMqttSession) SetMqttDisconnectRxCount(v int64)`

SetMqttDisconnectRxCount sets MqttDisconnectRxCount field to given value.

### HasMqttDisconnectRxCount

`func (o *MsgVpnMqttSession) HasMqttDisconnectRxCount() bool`

HasMqttDisconnectRxCount returns a boolean if a field has been set.

### GetMqttPubcompTxCount

`func (o *MsgVpnMqttSession) GetMqttPubcompTxCount() int64`

GetMqttPubcompTxCount returns the MqttPubcompTxCount field if non-nil, zero value otherwise.

### GetMqttPubcompTxCountOk

`func (o *MsgVpnMqttSession) GetMqttPubcompTxCountOk() (*int64, bool)`

GetMqttPubcompTxCountOk returns a tuple with the MqttPubcompTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubcompTxCount

`func (o *MsgVpnMqttSession) SetMqttPubcompTxCount(v int64)`

SetMqttPubcompTxCount sets MqttPubcompTxCount field to given value.

### HasMqttPubcompTxCount

`func (o *MsgVpnMqttSession) HasMqttPubcompTxCount() bool`

HasMqttPubcompTxCount returns a boolean if a field has been set.

### GetMqttPublishQos0RxCount

`func (o *MsgVpnMqttSession) GetMqttPublishQos0RxCount() int64`

GetMqttPublishQos0RxCount returns the MqttPublishQos0RxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos0RxCountOk

`func (o *MsgVpnMqttSession) GetMqttPublishQos0RxCountOk() (*int64, bool)`

GetMqttPublishQos0RxCountOk returns a tuple with the MqttPublishQos0RxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos0RxCount

`func (o *MsgVpnMqttSession) SetMqttPublishQos0RxCount(v int64)`

SetMqttPublishQos0RxCount sets MqttPublishQos0RxCount field to given value.

### HasMqttPublishQos0RxCount

`func (o *MsgVpnMqttSession) HasMqttPublishQos0RxCount() bool`

HasMqttPublishQos0RxCount returns a boolean if a field has been set.

### GetMqttPublishQos0TxCount

`func (o *MsgVpnMqttSession) GetMqttPublishQos0TxCount() int64`

GetMqttPublishQos0TxCount returns the MqttPublishQos0TxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos0TxCountOk

`func (o *MsgVpnMqttSession) GetMqttPublishQos0TxCountOk() (*int64, bool)`

GetMqttPublishQos0TxCountOk returns a tuple with the MqttPublishQos0TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos0TxCount

`func (o *MsgVpnMqttSession) SetMqttPublishQos0TxCount(v int64)`

SetMqttPublishQos0TxCount sets MqttPublishQos0TxCount field to given value.

### HasMqttPublishQos0TxCount

`func (o *MsgVpnMqttSession) HasMqttPublishQos0TxCount() bool`

HasMqttPublishQos0TxCount returns a boolean if a field has been set.

### GetMqttPublishQos1RxCount

`func (o *MsgVpnMqttSession) GetMqttPublishQos1RxCount() int64`

GetMqttPublishQos1RxCount returns the MqttPublishQos1RxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos1RxCountOk

`func (o *MsgVpnMqttSession) GetMqttPublishQos1RxCountOk() (*int64, bool)`

GetMqttPublishQos1RxCountOk returns a tuple with the MqttPublishQos1RxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos1RxCount

`func (o *MsgVpnMqttSession) SetMqttPublishQos1RxCount(v int64)`

SetMqttPublishQos1RxCount sets MqttPublishQos1RxCount field to given value.

### HasMqttPublishQos1RxCount

`func (o *MsgVpnMqttSession) HasMqttPublishQos1RxCount() bool`

HasMqttPublishQos1RxCount returns a boolean if a field has been set.

### GetMqttPublishQos1TxCount

`func (o *MsgVpnMqttSession) GetMqttPublishQos1TxCount() int64`

GetMqttPublishQos1TxCount returns the MqttPublishQos1TxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos1TxCountOk

`func (o *MsgVpnMqttSession) GetMqttPublishQos1TxCountOk() (*int64, bool)`

GetMqttPublishQos1TxCountOk returns a tuple with the MqttPublishQos1TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos1TxCount

`func (o *MsgVpnMqttSession) SetMqttPublishQos1TxCount(v int64)`

SetMqttPublishQos1TxCount sets MqttPublishQos1TxCount field to given value.

### HasMqttPublishQos1TxCount

`func (o *MsgVpnMqttSession) HasMqttPublishQos1TxCount() bool`

HasMqttPublishQos1TxCount returns a boolean if a field has been set.

### GetMqttPublishQos2RxCount

`func (o *MsgVpnMqttSession) GetMqttPublishQos2RxCount() int64`

GetMqttPublishQos2RxCount returns the MqttPublishQos2RxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos2RxCountOk

`func (o *MsgVpnMqttSession) GetMqttPublishQos2RxCountOk() (*int64, bool)`

GetMqttPublishQos2RxCountOk returns a tuple with the MqttPublishQos2RxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos2RxCount

`func (o *MsgVpnMqttSession) SetMqttPublishQos2RxCount(v int64)`

SetMqttPublishQos2RxCount sets MqttPublishQos2RxCount field to given value.

### HasMqttPublishQos2RxCount

`func (o *MsgVpnMqttSession) HasMqttPublishQos2RxCount() bool`

HasMqttPublishQos2RxCount returns a boolean if a field has been set.

### GetMqttPubrecTxCount

`func (o *MsgVpnMqttSession) GetMqttPubrecTxCount() int64`

GetMqttPubrecTxCount returns the MqttPubrecTxCount field if non-nil, zero value otherwise.

### GetMqttPubrecTxCountOk

`func (o *MsgVpnMqttSession) GetMqttPubrecTxCountOk() (*int64, bool)`

GetMqttPubrecTxCountOk returns a tuple with the MqttPubrecTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubrecTxCount

`func (o *MsgVpnMqttSession) SetMqttPubrecTxCount(v int64)`

SetMqttPubrecTxCount sets MqttPubrecTxCount field to given value.

### HasMqttPubrecTxCount

`func (o *MsgVpnMqttSession) HasMqttPubrecTxCount() bool`

HasMqttPubrecTxCount returns a boolean if a field has been set.

### GetMqttPubrelRxCount

`func (o *MsgVpnMqttSession) GetMqttPubrelRxCount() int64`

GetMqttPubrelRxCount returns the MqttPubrelRxCount field if non-nil, zero value otherwise.

### GetMqttPubrelRxCountOk

`func (o *MsgVpnMqttSession) GetMqttPubrelRxCountOk() (*int64, bool)`

GetMqttPubrelRxCountOk returns a tuple with the MqttPubrelRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubrelRxCount

`func (o *MsgVpnMqttSession) SetMqttPubrelRxCount(v int64)`

SetMqttPubrelRxCount sets MqttPubrelRxCount field to given value.

### HasMqttPubrelRxCount

`func (o *MsgVpnMqttSession) HasMqttPubrelRxCount() bool`

HasMqttPubrelRxCount returns a boolean if a field has been set.

### GetMqttSessionClientId

`func (o *MsgVpnMqttSession) GetMqttSessionClientId() string`

GetMqttSessionClientId returns the MqttSessionClientId field if non-nil, zero value otherwise.

### GetMqttSessionClientIdOk

`func (o *MsgVpnMqttSession) GetMqttSessionClientIdOk() (*string, bool)`

GetMqttSessionClientIdOk returns a tuple with the MqttSessionClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttSessionClientId

`func (o *MsgVpnMqttSession) SetMqttSessionClientId(v string)`

SetMqttSessionClientId sets MqttSessionClientId field to given value.

### HasMqttSessionClientId

`func (o *MsgVpnMqttSession) HasMqttSessionClientId() bool`

HasMqttSessionClientId returns a boolean if a field has been set.

### GetMqttSessionVirtualRouter

`func (o *MsgVpnMqttSession) GetMqttSessionVirtualRouter() string`

GetMqttSessionVirtualRouter returns the MqttSessionVirtualRouter field if non-nil, zero value otherwise.

### GetMqttSessionVirtualRouterOk

`func (o *MsgVpnMqttSession) GetMqttSessionVirtualRouterOk() (*string, bool)`

GetMqttSessionVirtualRouterOk returns a tuple with the MqttSessionVirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttSessionVirtualRouter

`func (o *MsgVpnMqttSession) SetMqttSessionVirtualRouter(v string)`

SetMqttSessionVirtualRouter sets MqttSessionVirtualRouter field to given value.

### HasMqttSessionVirtualRouter

`func (o *MsgVpnMqttSession) HasMqttSessionVirtualRouter() bool`

HasMqttSessionVirtualRouter returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnMqttSession) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnMqttSession) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnMqttSession) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnMqttSession) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetOwner

`func (o *MsgVpnMqttSession) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MsgVpnMqttSession) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MsgVpnMqttSession) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MsgVpnMqttSession) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetQueueConsumerAckPropagationEnabled

`func (o *MsgVpnMqttSession) GetQueueConsumerAckPropagationEnabled() bool`

GetQueueConsumerAckPropagationEnabled returns the QueueConsumerAckPropagationEnabled field if non-nil, zero value otherwise.

### GetQueueConsumerAckPropagationEnabledOk

`func (o *MsgVpnMqttSession) GetQueueConsumerAckPropagationEnabledOk() (*bool, bool)`

GetQueueConsumerAckPropagationEnabledOk returns a tuple with the QueueConsumerAckPropagationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueConsumerAckPropagationEnabled

`func (o *MsgVpnMqttSession) SetQueueConsumerAckPropagationEnabled(v bool)`

SetQueueConsumerAckPropagationEnabled sets QueueConsumerAckPropagationEnabled field to given value.

### HasQueueConsumerAckPropagationEnabled

`func (o *MsgVpnMqttSession) HasQueueConsumerAckPropagationEnabled() bool`

HasQueueConsumerAckPropagationEnabled returns a boolean if a field has been set.

### GetQueueDeadMsgQueue

`func (o *MsgVpnMqttSession) GetQueueDeadMsgQueue() string`

GetQueueDeadMsgQueue returns the QueueDeadMsgQueue field if non-nil, zero value otherwise.

### GetQueueDeadMsgQueueOk

`func (o *MsgVpnMqttSession) GetQueueDeadMsgQueueOk() (*string, bool)`

GetQueueDeadMsgQueueOk returns a tuple with the QueueDeadMsgQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDeadMsgQueue

`func (o *MsgVpnMqttSession) SetQueueDeadMsgQueue(v string)`

SetQueueDeadMsgQueue sets QueueDeadMsgQueue field to given value.

### HasQueueDeadMsgQueue

`func (o *MsgVpnMqttSession) HasQueueDeadMsgQueue() bool`

HasQueueDeadMsgQueue returns a boolean if a field has been set.

### GetQueueEventBindCountThreshold

`func (o *MsgVpnMqttSession) GetQueueEventBindCountThreshold() EventThreshold`

GetQueueEventBindCountThreshold returns the QueueEventBindCountThreshold field if non-nil, zero value otherwise.

### GetQueueEventBindCountThresholdOk

`func (o *MsgVpnMqttSession) GetQueueEventBindCountThresholdOk() (*EventThreshold, bool)`

GetQueueEventBindCountThresholdOk returns a tuple with the QueueEventBindCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueEventBindCountThreshold

`func (o *MsgVpnMqttSession) SetQueueEventBindCountThreshold(v EventThreshold)`

SetQueueEventBindCountThreshold sets QueueEventBindCountThreshold field to given value.

### HasQueueEventBindCountThreshold

`func (o *MsgVpnMqttSession) HasQueueEventBindCountThreshold() bool`

HasQueueEventBindCountThreshold returns a boolean if a field has been set.

### GetQueueEventMsgSpoolUsageThreshold

`func (o *MsgVpnMqttSession) GetQueueEventMsgSpoolUsageThreshold() EventThreshold`

GetQueueEventMsgSpoolUsageThreshold returns the QueueEventMsgSpoolUsageThreshold field if non-nil, zero value otherwise.

### GetQueueEventMsgSpoolUsageThresholdOk

`func (o *MsgVpnMqttSession) GetQueueEventMsgSpoolUsageThresholdOk() (*EventThreshold, bool)`

GetQueueEventMsgSpoolUsageThresholdOk returns a tuple with the QueueEventMsgSpoolUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueEventMsgSpoolUsageThreshold

`func (o *MsgVpnMqttSession) SetQueueEventMsgSpoolUsageThreshold(v EventThreshold)`

SetQueueEventMsgSpoolUsageThreshold sets QueueEventMsgSpoolUsageThreshold field to given value.

### HasQueueEventMsgSpoolUsageThreshold

`func (o *MsgVpnMqttSession) HasQueueEventMsgSpoolUsageThreshold() bool`

HasQueueEventMsgSpoolUsageThreshold returns a boolean if a field has been set.

### GetQueueEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnMqttSession) GetQueueEventRejectLowPriorityMsgLimitThreshold() EventThreshold`

GetQueueEventRejectLowPriorityMsgLimitThreshold returns the QueueEventRejectLowPriorityMsgLimitThreshold field if non-nil, zero value otherwise.

### GetQueueEventRejectLowPriorityMsgLimitThresholdOk

`func (o *MsgVpnMqttSession) GetQueueEventRejectLowPriorityMsgLimitThresholdOk() (*EventThreshold, bool)`

GetQueueEventRejectLowPriorityMsgLimitThresholdOk returns a tuple with the QueueEventRejectLowPriorityMsgLimitThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnMqttSession) SetQueueEventRejectLowPriorityMsgLimitThreshold(v EventThreshold)`

SetQueueEventRejectLowPriorityMsgLimitThreshold sets QueueEventRejectLowPriorityMsgLimitThreshold field to given value.

### HasQueueEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnMqttSession) HasQueueEventRejectLowPriorityMsgLimitThreshold() bool`

HasQueueEventRejectLowPriorityMsgLimitThreshold returns a boolean if a field has been set.

### GetQueueMaxBindCount

`func (o *MsgVpnMqttSession) GetQueueMaxBindCount() int64`

GetQueueMaxBindCount returns the QueueMaxBindCount field if non-nil, zero value otherwise.

### GetQueueMaxBindCountOk

`func (o *MsgVpnMqttSession) GetQueueMaxBindCountOk() (*int64, bool)`

GetQueueMaxBindCountOk returns a tuple with the QueueMaxBindCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxBindCount

`func (o *MsgVpnMqttSession) SetQueueMaxBindCount(v int64)`

SetQueueMaxBindCount sets QueueMaxBindCount field to given value.

### HasQueueMaxBindCount

`func (o *MsgVpnMqttSession) HasQueueMaxBindCount() bool`

HasQueueMaxBindCount returns a boolean if a field has been set.

### GetQueueMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnMqttSession) GetQueueMaxDeliveredUnackedMsgsPerFlow() int64`

GetQueueMaxDeliveredUnackedMsgsPerFlow returns the QueueMaxDeliveredUnackedMsgsPerFlow field if non-nil, zero value otherwise.

### GetQueueMaxDeliveredUnackedMsgsPerFlowOk

`func (o *MsgVpnMqttSession) GetQueueMaxDeliveredUnackedMsgsPerFlowOk() (*int64, bool)`

GetQueueMaxDeliveredUnackedMsgsPerFlowOk returns a tuple with the QueueMaxDeliveredUnackedMsgsPerFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnMqttSession) SetQueueMaxDeliveredUnackedMsgsPerFlow(v int64)`

SetQueueMaxDeliveredUnackedMsgsPerFlow sets QueueMaxDeliveredUnackedMsgsPerFlow field to given value.

### HasQueueMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnMqttSession) HasQueueMaxDeliveredUnackedMsgsPerFlow() bool`

HasQueueMaxDeliveredUnackedMsgsPerFlow returns a boolean if a field has been set.

### GetQueueMaxMsgSize

`func (o *MsgVpnMqttSession) GetQueueMaxMsgSize() int32`

GetQueueMaxMsgSize returns the QueueMaxMsgSize field if non-nil, zero value otherwise.

### GetQueueMaxMsgSizeOk

`func (o *MsgVpnMqttSession) GetQueueMaxMsgSizeOk() (*int32, bool)`

GetQueueMaxMsgSizeOk returns a tuple with the QueueMaxMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxMsgSize

`func (o *MsgVpnMqttSession) SetQueueMaxMsgSize(v int32)`

SetQueueMaxMsgSize sets QueueMaxMsgSize field to given value.

### HasQueueMaxMsgSize

`func (o *MsgVpnMqttSession) HasQueueMaxMsgSize() bool`

HasQueueMaxMsgSize returns a boolean if a field has been set.

### GetQueueMaxMsgSpoolUsage

`func (o *MsgVpnMqttSession) GetQueueMaxMsgSpoolUsage() int64`

GetQueueMaxMsgSpoolUsage returns the QueueMaxMsgSpoolUsage field if non-nil, zero value otherwise.

### GetQueueMaxMsgSpoolUsageOk

`func (o *MsgVpnMqttSession) GetQueueMaxMsgSpoolUsageOk() (*int64, bool)`

GetQueueMaxMsgSpoolUsageOk returns a tuple with the QueueMaxMsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxMsgSpoolUsage

`func (o *MsgVpnMqttSession) SetQueueMaxMsgSpoolUsage(v int64)`

SetQueueMaxMsgSpoolUsage sets QueueMaxMsgSpoolUsage field to given value.

### HasQueueMaxMsgSpoolUsage

`func (o *MsgVpnMqttSession) HasQueueMaxMsgSpoolUsage() bool`

HasQueueMaxMsgSpoolUsage returns a boolean if a field has been set.

### GetQueueMaxRedeliveryCount

`func (o *MsgVpnMqttSession) GetQueueMaxRedeliveryCount() int64`

GetQueueMaxRedeliveryCount returns the QueueMaxRedeliveryCount field if non-nil, zero value otherwise.

### GetQueueMaxRedeliveryCountOk

`func (o *MsgVpnMqttSession) GetQueueMaxRedeliveryCountOk() (*int64, bool)`

GetQueueMaxRedeliveryCountOk returns a tuple with the QueueMaxRedeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxRedeliveryCount

`func (o *MsgVpnMqttSession) SetQueueMaxRedeliveryCount(v int64)`

SetQueueMaxRedeliveryCount sets QueueMaxRedeliveryCount field to given value.

### HasQueueMaxRedeliveryCount

`func (o *MsgVpnMqttSession) HasQueueMaxRedeliveryCount() bool`

HasQueueMaxRedeliveryCount returns a boolean if a field has been set.

### GetQueueMaxTtl

`func (o *MsgVpnMqttSession) GetQueueMaxTtl() int64`

GetQueueMaxTtl returns the QueueMaxTtl field if non-nil, zero value otherwise.

### GetQueueMaxTtlOk

`func (o *MsgVpnMqttSession) GetQueueMaxTtlOk() (*int64, bool)`

GetQueueMaxTtlOk returns a tuple with the QueueMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxTtl

`func (o *MsgVpnMqttSession) SetQueueMaxTtl(v int64)`

SetQueueMaxTtl sets QueueMaxTtl field to given value.

### HasQueueMaxTtl

`func (o *MsgVpnMqttSession) HasQueueMaxTtl() bool`

HasQueueMaxTtl returns a boolean if a field has been set.

### GetQueueName

`func (o *MsgVpnMqttSession) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *MsgVpnMqttSession) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *MsgVpnMqttSession) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *MsgVpnMqttSession) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetQueueRejectLowPriorityMsgEnabled

`func (o *MsgVpnMqttSession) GetQueueRejectLowPriorityMsgEnabled() bool`

GetQueueRejectLowPriorityMsgEnabled returns the QueueRejectLowPriorityMsgEnabled field if non-nil, zero value otherwise.

### GetQueueRejectLowPriorityMsgEnabledOk

`func (o *MsgVpnMqttSession) GetQueueRejectLowPriorityMsgEnabledOk() (*bool, bool)`

GetQueueRejectLowPriorityMsgEnabledOk returns a tuple with the QueueRejectLowPriorityMsgEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueRejectLowPriorityMsgEnabled

`func (o *MsgVpnMqttSession) SetQueueRejectLowPriorityMsgEnabled(v bool)`

SetQueueRejectLowPriorityMsgEnabled sets QueueRejectLowPriorityMsgEnabled field to given value.

### HasQueueRejectLowPriorityMsgEnabled

`func (o *MsgVpnMqttSession) HasQueueRejectLowPriorityMsgEnabled() bool`

HasQueueRejectLowPriorityMsgEnabled returns a boolean if a field has been set.

### GetQueueRejectLowPriorityMsgLimit

`func (o *MsgVpnMqttSession) GetQueueRejectLowPriorityMsgLimit() int64`

GetQueueRejectLowPriorityMsgLimit returns the QueueRejectLowPriorityMsgLimit field if non-nil, zero value otherwise.

### GetQueueRejectLowPriorityMsgLimitOk

`func (o *MsgVpnMqttSession) GetQueueRejectLowPriorityMsgLimitOk() (*int64, bool)`

GetQueueRejectLowPriorityMsgLimitOk returns a tuple with the QueueRejectLowPriorityMsgLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueRejectLowPriorityMsgLimit

`func (o *MsgVpnMqttSession) SetQueueRejectLowPriorityMsgLimit(v int64)`

SetQueueRejectLowPriorityMsgLimit sets QueueRejectLowPriorityMsgLimit field to given value.

### HasQueueRejectLowPriorityMsgLimit

`func (o *MsgVpnMqttSession) HasQueueRejectLowPriorityMsgLimit() bool`

HasQueueRejectLowPriorityMsgLimit returns a boolean if a field has been set.

### GetQueueRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnMqttSession) GetQueueRejectMsgToSenderOnDiscardBehavior() string`

GetQueueRejectMsgToSenderOnDiscardBehavior returns the QueueRejectMsgToSenderOnDiscardBehavior field if non-nil, zero value otherwise.

### GetQueueRejectMsgToSenderOnDiscardBehaviorOk

`func (o *MsgVpnMqttSession) GetQueueRejectMsgToSenderOnDiscardBehaviorOk() (*string, bool)`

GetQueueRejectMsgToSenderOnDiscardBehaviorOk returns a tuple with the QueueRejectMsgToSenderOnDiscardBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnMqttSession) SetQueueRejectMsgToSenderOnDiscardBehavior(v string)`

SetQueueRejectMsgToSenderOnDiscardBehavior sets QueueRejectMsgToSenderOnDiscardBehavior field to given value.

### HasQueueRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnMqttSession) HasQueueRejectMsgToSenderOnDiscardBehavior() bool`

HasQueueRejectMsgToSenderOnDiscardBehavior returns a boolean if a field has been set.

### GetQueueRespectTtlEnabled

`func (o *MsgVpnMqttSession) GetQueueRespectTtlEnabled() bool`

GetQueueRespectTtlEnabled returns the QueueRespectTtlEnabled field if non-nil, zero value otherwise.

### GetQueueRespectTtlEnabledOk

`func (o *MsgVpnMqttSession) GetQueueRespectTtlEnabledOk() (*bool, bool)`

GetQueueRespectTtlEnabledOk returns a tuple with the QueueRespectTtlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueRespectTtlEnabled

`func (o *MsgVpnMqttSession) SetQueueRespectTtlEnabled(v bool)`

SetQueueRespectTtlEnabled sets QueueRespectTtlEnabled field to given value.

### HasQueueRespectTtlEnabled

`func (o *MsgVpnMqttSession) HasQueueRespectTtlEnabled() bool`

HasQueueRespectTtlEnabled returns a boolean if a field has been set.

### GetRxMax

`func (o *MsgVpnMqttSession) GetRxMax() int64`

GetRxMax returns the RxMax field if non-nil, zero value otherwise.

### GetRxMaxOk

`func (o *MsgVpnMqttSession) GetRxMaxOk() (*int64, bool)`

GetRxMaxOk returns a tuple with the RxMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMax

`func (o *MsgVpnMqttSession) SetRxMax(v int64)`

SetRxMax sets RxMax field to given value.

### HasRxMax

`func (o *MsgVpnMqttSession) HasRxMax() bool`

HasRxMax returns a boolean if a field has been set.

### GetWill

`func (o *MsgVpnMqttSession) GetWill() bool`

GetWill returns the Will field if non-nil, zero value otherwise.

### GetWillOk

`func (o *MsgVpnMqttSession) GetWillOk() (*bool, bool)`

GetWillOk returns a tuple with the Will field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWill

`func (o *MsgVpnMqttSession) SetWill(v bool)`

SetWill sets Will field to given value.

### HasWill

`func (o *MsgVpnMqttSession) HasWill() bool`

HasWill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


