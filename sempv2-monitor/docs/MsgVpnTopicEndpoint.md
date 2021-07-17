# MsgVpnTopicEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** | The access type for delivering messages to consumer flows bound to the Topic Endpoint. The allowed values and their meaning are:  &lt;pre&gt; \&quot;exclusive\&quot; - Exclusive delivery of messages to the first bound consumer flow. \&quot;non-exclusive\&quot; - Non-exclusive delivery of messages to all bound consumer flows in a round-robin fashion. &lt;/pre&gt;  | [optional] 
**AlreadyBoundBindFailureCount** | Pointer to **int64** | The number of Topic Endpoint bind failures due to being already bound. | [optional] 
**AverageRxByteRate** | Pointer to **int64** | The one minute average of the message rate received by the Topic Endpoint, in bytes per second (B/sec). | [optional] 
**AverageRxMsgRate** | Pointer to **int64** | The one minute average of the message rate received by the Topic Endpoint, in messages per second (msg/sec). | [optional] 
**AverageTxByteRate** | Pointer to **int64** | The one minute average of the message rate transmitted by the Topic Endpoint, in bytes per second (B/sec). | [optional] 
**AverageTxMsgRate** | Pointer to **int64** | The one minute average of the message rate transmitted by the Topic Endpoint, in messages per second (msg/sec). | [optional] 
**BindRequestCount** | Pointer to **int64** | The number of consumer requests to bind to the Topic Endpoint. | [optional] 
**BindSuccessCount** | Pointer to **int64** | The number of successful consumer requests to bind to the Topic Endpoint. | [optional] 
**BindTimeForwardingMode** | Pointer to **string** | The forwarding mode of the Topic Endpoint at bind time. The allowed values and their meaning are:  &lt;pre&gt; \&quot;store-and-forward\&quot; - Deliver messages using the guaranteed data path. \&quot;cut-through\&quot; - Deliver messages using the direct and guaranteed data paths for lower latency. &lt;/pre&gt;  | [optional] 
**ClientProfileDeniedDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to being denied by the Client Profile. | [optional] 
**ConsumerAckPropagationEnabled** | Pointer to **bool** | Indicates whether the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN is enabled. | [optional] 
**CreatedByManagement** | Pointer to **bool** | Indicates whether the Topic Endpoint was created by a management API (CLI or SEMP). | [optional] 
**DeadMsgQueue** | Pointer to **string** | The name of the Dead Message Queue (DMQ) used by the Topic Endpoint. | [optional] 
**DeletedMsgCount** | Pointer to **int64** | The number of guaranteed messages deleted from the Topic Endpoint. | [optional] 
**DeliveryCountEnabled** | Pointer to **bool** | Enable or disable the ability for client applications to query the message delivery count of messages received from the Topic Endpoint. This is a controlled availability feature. Please contact Solace to find out if this feature is supported for your use case. Available since 2.19. | [optional] 
**DestinationGroupErrorDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to a destination group error. | [optional] 
**DestinationTopic** | Pointer to **string** | The destination topic of the Topic Endpoint. | [optional] 
**DisabledBindFailureCount** | Pointer to **int64** | The number of Topic Endpoint bind failures due to being disabled. | [optional] 
**DisabledDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to it being disabled. | [optional] 
**Durable** | Pointer to **bool** | Indicates whether the Topic Endpoint is durable and not temporary. | [optional] 
**EgressEnabled** | Pointer to **bool** | Indicates whether the transmission of messages from the Topic Endpoint is enabled. | [optional] 
**EventBindCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventRejectLowPriorityMsgLimitThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**HighestAckedMsgId** | Pointer to **int64** | The highest identifier (ID) of guaranteed messages in the Topic Endpoint that were acknowledged. | [optional] 
**HighestMsgId** | Pointer to **int64** | The highest identifier (ID) of guaranteed messages in the Topic Endpoint. | [optional] 
**InProgressAckMsgCount** | Pointer to **int64** | The number of acknowledgement messages received by the Topic Endpoint that are in the process of updating and deleting associated guaranteed messages. | [optional] 
**IngressEnabled** | Pointer to **bool** | Indicates whether the reception of messages to the Topic Endpoint is enabled. | [optional] 
**InvalidSelectorBindFailureCount** | Pointer to **int64** | The number of Topic Endpoint bind failures due to an invalid selector. | [optional] 
**LastReplayCompleteTime** | Pointer to **int32** | The timestamp of the last completed replay for the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastReplayFailureReason** | Pointer to **string** | The reason for the last replay failure for the Topic Endpoint. | [optional] 
**LastReplayFailureTime** | Pointer to **int32** | The timestamp of the last replay failure for the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastReplayStartTime** | Pointer to **int32** | The timestamp of the last replay started for the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastReplayedMsgTxTime** | Pointer to **int32** | The timestamp of the last replayed message transmitted by the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastSelectorExaminedMsgId** | Pointer to **int64** | The identifier (ID) of the last message examined by the Topic Endpoint selector. | [optional] 
**LastSpooledMsgId** | Pointer to **int64** | The identifier (ID) of the last guaranteed message spooled in the Topic Endpoint. | [optional] 
**LowPriorityMsgCongestionDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to low priority message congestion control. | [optional] 
**LowPriorityMsgCongestionState** | Pointer to **string** | The state of the low priority message congestion in the Topic Endpoint. The allowed values and their meaning are:  &lt;pre&gt; \&quot;disabled\&quot; - Messages are not being checked for priority. \&quot;not-congested\&quot; - Low priority messages are being stored and delivered. \&quot;congested\&quot; - Low priority messages are being discarded. &lt;/pre&gt;  | [optional] 
**LowestAckedMsgId** | Pointer to **int64** | The lowest identifier (ID) of guaranteed messages in the Topic Endpoint that were acknowledged. | [optional] 
**LowestMsgId** | Pointer to **int64** | The lowest identifier (ID) of guaranteed messages in the Topic Endpoint. | [optional] 
**MaxBindCount** | Pointer to **int64** | The maximum number of consumer flows that can bind to the Topic Endpoint. | [optional] 
**MaxBindCountExceededBindFailureCount** | Pointer to **int64** | The number of Topic Endpoint bind failures due to the maximum bind count being exceeded. | [optional] 
**MaxDeliveredUnackedMsgsPerFlow** | Pointer to **int64** | The maximum number of messages delivered but not acknowledged per flow for the Topic Endpoint. | [optional] 
**MaxEffectiveBindCount** | Pointer to **int32** | The effective maximum number of consumer flows that can bind to the Topic Endpoint. | [optional] 
**MaxMsgSize** | Pointer to **int32** | The maximum message size allowed in the Topic Endpoint, in bytes (B). | [optional] 
**MaxMsgSizeExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to the maximum message size being exceeded. | [optional] 
**MaxMsgSpoolUsageExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to the maximum message spool usage being exceeded. | [optional] 
**MaxRedeliveryCount** | Pointer to **int64** | The maximum number of times the Topic Endpoint will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. | [optional] 
**MaxRedeliveryExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to the maximum redelivery attempts being exceeded. | [optional] 
**MaxRedeliveryExceededToDmqFailedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to the maximum redelivery attempts being exceeded and failing to move to the Dead Message Queue (DMQ). | [optional] 
**MaxRedeliveryExceededToDmqMsgCount** | Pointer to **int64** | The number of guaranteed messages moved to the Dead Message Queue (DMQ) by the Topic Endpoint due to the maximum redelivery attempts being exceeded. | [optional] 
**MaxSpoolUsage** | Pointer to **int64** | The maximum message spool usage allowed by the Topic Endpoint, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. | [optional] 
**MaxTtl** | Pointer to **int64** | The maximum time in seconds a message can stay in the Topic Endpoint when &#x60;respectTtlEnabled&#x60; is &#x60;\&quot;true\&quot;&#x60;. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the &#x60;maxTtl&#x60; configured for the Topic Endpoint, is exceeded. A value of 0 disables expiry. | [optional] 
**MaxTtlExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to the maximum time-to-live (TTL) in hops being exceeded. The TTL hop count is incremented when the message crosses a bridge. | [optional] 
**MaxTtlExpiredDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to the maximum time-to-live (TTL) timestamp expiring. | [optional] 
**MaxTtlExpiredToDmqFailedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to the maximum time-to-live (TTL) timestamp expiring and failing to move to the Dead Message Queue (DMQ). | [optional] 
**MaxTtlExpiredToDmqMsgCount** | Pointer to **int64** | The number of guaranteed messages moved to the Dead Message Queue (DMQ) by the Topic Endpoint due to the maximum time-to-live (TTL) timestamp expiring. | [optional] 
**MsgSpoolPeakUsage** | Pointer to **int64** | The message spool peak usage by the Topic Endpoint, in bytes (B). | [optional] 
**MsgSpoolUsage** | Pointer to **int64** | The message spool usage by the Topic Endpoint, in bytes (B). | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**NetworkTopic** | Pointer to **string** | The name of the network topic for the Topic Endpoint. | [optional] 
**NoLocalDeliveryDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Topic Endpoint due to no local delivery being requested. | [optional] 
**OtherBindFailureCount** | Pointer to **int64** | The number of Topic Endpoint bind failures due to other reasons. | [optional] 
**Owner** | Pointer to **string** | The Client Username that owns the Topic Endpoint and has permission equivalent to &#x60;\&quot;delete\&quot;&#x60;. | [optional] 
**Permission** | Pointer to **string** | The permission level for all consumers of the Topic Endpoint, excluding the owner. The allowed values and their meaning are:  &lt;pre&gt; \&quot;no-access\&quot; - Disallows all access. \&quot;read-only\&quot; - Read-only access to the messages. \&quot;consume\&quot; - Consume (read and remove) messages. \&quot;modify-topic\&quot; - Consume messages or modify the topic/selector. \&quot;delete\&quot; - Consume messages, modify the topic/selector or delete the Client created endpoint altogether. &lt;/pre&gt;  | [optional] 
**RedeliveredMsgCount** | Pointer to **int64** | The number of guaranteed messages transmitted by the Topic Endpoint for redelivery. | [optional] 
**RedeliveryEnabled** | Pointer to **bool** | Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the topic-endpoint more than once. Available since 2.18. | [optional] 
**RejectLowPriorityMsgEnabled** | Pointer to **bool** | Indicates whether the checking of low priority messages against the &#x60;rejectLowPriorityMsgLimit&#x60; is enabled. | [optional] 
**RejectLowPriorityMsgLimit** | Pointer to **int64** | The number of messages of any priority in the Topic Endpoint above which low priority messages are not admitted but higher priority messages are allowed. | [optional] 
**RejectMsgToSenderOnDiscardBehavior** | Pointer to **string** | Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always return a negative acknowledgment (NACK) to the sending client on message discard. \&quot;when-topic-endpoint-enabled\&quot; - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Topic Endpoint is enabled. \&quot;never\&quot; - Never return a negative acknowledgment (NACK) to the sending client on message discard. &lt;/pre&gt;  | [optional] 
**ReplayFailureCount** | Pointer to **int64** | The number of replays that failed for the Topic Endpoint. | [optional] 
**ReplayStartCount** | Pointer to **int64** | The number of replays started for the Topic Endpoint. | [optional] 
**ReplayState** | Pointer to **string** | The state of replay for the Topic Endpoint. The allowed values and their meaning are:  &lt;pre&gt; \&quot;initializing\&quot; - All messages are being deleted from the endpoint before replay starts. \&quot;active\&quot; - Subscription matching logged messages are being replayed to the endpoint. \&quot;pending-complete\&quot; - Replay is complete, but final accounting is in progress. \&quot;complete\&quot; - Replay and all related activities are complete. \&quot;failed\&quot; - Replay has failed and is waiting for an unbind response. &lt;/pre&gt;  | [optional] 
**ReplaySuccessCount** | Pointer to **int64** | The number of replays that succeeded for the Topic Endpoint. | [optional] 
**ReplayedAckedMsgCount** | Pointer to **int64** | The number of replayed messages transmitted by the Topic Endpoint and acked by all consumers. | [optional] 
**ReplayedTxMsgCount** | Pointer to **int64** | The number of replayed messages transmitted by the Topic Endpoint. | [optional] 
**ReplicationActiveAckPropTxMsgCount** | Pointer to **int64** | The number of acknowledgement messages propagated by the Topic Endpoint to the replication standby remote Message VPN. | [optional] 
**ReplicationStandbyAckPropRxMsgCount** | Pointer to **int64** | The number of propagated acknowledgement messages received by the Topic Endpoint from the replication active remote Message VPN. | [optional] 
**ReplicationStandbyAckedByAckPropMsgCount** | Pointer to **int64** | The number of messages acknowledged in the Topic Endpoint by acknowledgement propagation from the replication active remote Message VPN. | [optional] 
**ReplicationStandbyRxMsgCount** | Pointer to **int64** | The number of messages received by the Topic Endpoint from the replication active remote Message VPN. | [optional] 
**RespectMsgPriorityEnabled** | Pointer to **bool** | Indicates whether message priorities are respected. When enabled, messages contained in the Topic Endpoint are delivered in priority order, from 9 (highest) to 0 (lowest). | [optional] 
**RespectTtlEnabled** | Pointer to **bool** | Indicates whether the time-to-live (TTL) for messages in the Topic Endpoint is respected. When enabled, expired messages are discarded or moved to the DMQ. | [optional] 
**RxByteRate** | Pointer to **int32** | The current message rate received by the Topic Endpoint, in bytes per second (B/sec). | [optional] 
**RxMsgRate** | Pointer to **int64** | The current message rate received by the Topic Endpoint, in messages per second (msg/sec). | [optional] 
**RxSelector** | Pointer to **bool** | Indicates whether the Topic Endpoint has a selector to filter received messages. | [optional] 
**Selector** | Pointer to **string** | The value of the receive selector for the Topic Endpoint. | [optional] 
**SelectorExaminedMsgCount** | Pointer to **int64** | The number of guaranteed messages examined by the Topic Endpoint selector. | [optional] 
**SelectorMatchedMsgCount** | Pointer to **int64** | The number of guaranteed messages for which the Topic Endpoint selector matched. | [optional] 
**SelectorNotMatchedMsgCount** | Pointer to **int64** | The number of guaranteed messages for which the Topic Endpoint selector did not match. | [optional] 
**SpooledByteCount** | Pointer to **int64** | The amount of guaranteed messages that were spooled in the Topic Endpoint, in bytes (B). | [optional] 
**SpooledMsgCount** | Pointer to **int64** | The number of guaranteed messages that were spooled in the Topic Endpoint. | [optional] 
**TopicEndpointName** | Pointer to **string** | The name of the Topic Endpoint. | [optional] 
**TransportRetransmitMsgCount** | Pointer to **int64** | The number of guaranteed messages that were retransmitted by the Topic Endpoint at the transport layer as part of a single delivery attempt. Available since 2.18. | [optional] 
**TxByteRate** | Pointer to **int64** | The current message rate transmitted by the Topic Endpoint, in bytes per second (B/sec). | [optional] 
**TxMsgRate** | Pointer to **int64** | The current message rate transmitted by the Topic Endpoint, in messages per second (msg/sec). | [optional] 
**TxUnackedMsgCount** | Pointer to **int64** | The number of guaranteed messages in the Topic Endpoint that have been transmitted but not acknowledged by all consumers. | [optional] 
**VirtualRouter** | Pointer to **string** | The virtual router used by the Topic Endpoint. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The endpoint belongs to the primary virtual router. \&quot;backup\&quot; - The endpoint belongs to the backup virtual router. &lt;/pre&gt;  | [optional] 

## Methods

### NewMsgVpnTopicEndpoint

`func NewMsgVpnTopicEndpoint() *MsgVpnTopicEndpoint`

NewMsgVpnTopicEndpoint instantiates a new MsgVpnTopicEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTopicEndpointWithDefaults

`func NewMsgVpnTopicEndpointWithDefaults() *MsgVpnTopicEndpoint`

NewMsgVpnTopicEndpointWithDefaults instantiates a new MsgVpnTopicEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *MsgVpnTopicEndpoint) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *MsgVpnTopicEndpoint) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *MsgVpnTopicEndpoint) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *MsgVpnTopicEndpoint) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetAlreadyBoundBindFailureCount

`func (o *MsgVpnTopicEndpoint) GetAlreadyBoundBindFailureCount() int64`

GetAlreadyBoundBindFailureCount returns the AlreadyBoundBindFailureCount field if non-nil, zero value otherwise.

### GetAlreadyBoundBindFailureCountOk

`func (o *MsgVpnTopicEndpoint) GetAlreadyBoundBindFailureCountOk() (*int64, bool)`

GetAlreadyBoundBindFailureCountOk returns a tuple with the AlreadyBoundBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyBoundBindFailureCount

`func (o *MsgVpnTopicEndpoint) SetAlreadyBoundBindFailureCount(v int64)`

SetAlreadyBoundBindFailureCount sets AlreadyBoundBindFailureCount field to given value.

### HasAlreadyBoundBindFailureCount

`func (o *MsgVpnTopicEndpoint) HasAlreadyBoundBindFailureCount() bool`

HasAlreadyBoundBindFailureCount returns a boolean if a field has been set.

### GetAverageRxByteRate

`func (o *MsgVpnTopicEndpoint) GetAverageRxByteRate() int64`

GetAverageRxByteRate returns the AverageRxByteRate field if non-nil, zero value otherwise.

### GetAverageRxByteRateOk

`func (o *MsgVpnTopicEndpoint) GetAverageRxByteRateOk() (*int64, bool)`

GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxByteRate

`func (o *MsgVpnTopicEndpoint) SetAverageRxByteRate(v int64)`

SetAverageRxByteRate sets AverageRxByteRate field to given value.

### HasAverageRxByteRate

`func (o *MsgVpnTopicEndpoint) HasAverageRxByteRate() bool`

HasAverageRxByteRate returns a boolean if a field has been set.

### GetAverageRxMsgRate

`func (o *MsgVpnTopicEndpoint) GetAverageRxMsgRate() int64`

GetAverageRxMsgRate returns the AverageRxMsgRate field if non-nil, zero value otherwise.

### GetAverageRxMsgRateOk

`func (o *MsgVpnTopicEndpoint) GetAverageRxMsgRateOk() (*int64, bool)`

GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxMsgRate

`func (o *MsgVpnTopicEndpoint) SetAverageRxMsgRate(v int64)`

SetAverageRxMsgRate sets AverageRxMsgRate field to given value.

### HasAverageRxMsgRate

`func (o *MsgVpnTopicEndpoint) HasAverageRxMsgRate() bool`

HasAverageRxMsgRate returns a boolean if a field has been set.

### GetAverageTxByteRate

`func (o *MsgVpnTopicEndpoint) GetAverageTxByteRate() int64`

GetAverageTxByteRate returns the AverageTxByteRate field if non-nil, zero value otherwise.

### GetAverageTxByteRateOk

`func (o *MsgVpnTopicEndpoint) GetAverageTxByteRateOk() (*int64, bool)`

GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxByteRate

`func (o *MsgVpnTopicEndpoint) SetAverageTxByteRate(v int64)`

SetAverageTxByteRate sets AverageTxByteRate field to given value.

### HasAverageTxByteRate

`func (o *MsgVpnTopicEndpoint) HasAverageTxByteRate() bool`

HasAverageTxByteRate returns a boolean if a field has been set.

### GetAverageTxMsgRate

`func (o *MsgVpnTopicEndpoint) GetAverageTxMsgRate() int64`

GetAverageTxMsgRate returns the AverageTxMsgRate field if non-nil, zero value otherwise.

### GetAverageTxMsgRateOk

`func (o *MsgVpnTopicEndpoint) GetAverageTxMsgRateOk() (*int64, bool)`

GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxMsgRate

`func (o *MsgVpnTopicEndpoint) SetAverageTxMsgRate(v int64)`

SetAverageTxMsgRate sets AverageTxMsgRate field to given value.

### HasAverageTxMsgRate

`func (o *MsgVpnTopicEndpoint) HasAverageTxMsgRate() bool`

HasAverageTxMsgRate returns a boolean if a field has been set.

### GetBindRequestCount

`func (o *MsgVpnTopicEndpoint) GetBindRequestCount() int64`

GetBindRequestCount returns the BindRequestCount field if non-nil, zero value otherwise.

### GetBindRequestCountOk

`func (o *MsgVpnTopicEndpoint) GetBindRequestCountOk() (*int64, bool)`

GetBindRequestCountOk returns a tuple with the BindRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindRequestCount

`func (o *MsgVpnTopicEndpoint) SetBindRequestCount(v int64)`

SetBindRequestCount sets BindRequestCount field to given value.

### HasBindRequestCount

`func (o *MsgVpnTopicEndpoint) HasBindRequestCount() bool`

HasBindRequestCount returns a boolean if a field has been set.

### GetBindSuccessCount

`func (o *MsgVpnTopicEndpoint) GetBindSuccessCount() int64`

GetBindSuccessCount returns the BindSuccessCount field if non-nil, zero value otherwise.

### GetBindSuccessCountOk

`func (o *MsgVpnTopicEndpoint) GetBindSuccessCountOk() (*int64, bool)`

GetBindSuccessCountOk returns a tuple with the BindSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindSuccessCount

`func (o *MsgVpnTopicEndpoint) SetBindSuccessCount(v int64)`

SetBindSuccessCount sets BindSuccessCount field to given value.

### HasBindSuccessCount

`func (o *MsgVpnTopicEndpoint) HasBindSuccessCount() bool`

HasBindSuccessCount returns a boolean if a field has been set.

### GetBindTimeForwardingMode

`func (o *MsgVpnTopicEndpoint) GetBindTimeForwardingMode() string`

GetBindTimeForwardingMode returns the BindTimeForwardingMode field if non-nil, zero value otherwise.

### GetBindTimeForwardingModeOk

`func (o *MsgVpnTopicEndpoint) GetBindTimeForwardingModeOk() (*string, bool)`

GetBindTimeForwardingModeOk returns a tuple with the BindTimeForwardingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindTimeForwardingMode

`func (o *MsgVpnTopicEndpoint) SetBindTimeForwardingMode(v string)`

SetBindTimeForwardingMode sets BindTimeForwardingMode field to given value.

### HasBindTimeForwardingMode

`func (o *MsgVpnTopicEndpoint) HasBindTimeForwardingMode() bool`

HasBindTimeForwardingMode returns a boolean if a field has been set.

### GetClientProfileDeniedDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) GetClientProfileDeniedDiscardedMsgCount() int64`

GetClientProfileDeniedDiscardedMsgCount returns the ClientProfileDeniedDiscardedMsgCount field if non-nil, zero value otherwise.

### GetClientProfileDeniedDiscardedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetClientProfileDeniedDiscardedMsgCountOk() (*int64, bool)`

GetClientProfileDeniedDiscardedMsgCountOk returns a tuple with the ClientProfileDeniedDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileDeniedDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) SetClientProfileDeniedDiscardedMsgCount(v int64)`

SetClientProfileDeniedDiscardedMsgCount sets ClientProfileDeniedDiscardedMsgCount field to given value.

### HasClientProfileDeniedDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) HasClientProfileDeniedDiscardedMsgCount() bool`

HasClientProfileDeniedDiscardedMsgCount returns a boolean if a field has been set.

### GetConsumerAckPropagationEnabled

`func (o *MsgVpnTopicEndpoint) GetConsumerAckPropagationEnabled() bool`

GetConsumerAckPropagationEnabled returns the ConsumerAckPropagationEnabled field if non-nil, zero value otherwise.

### GetConsumerAckPropagationEnabledOk

`func (o *MsgVpnTopicEndpoint) GetConsumerAckPropagationEnabledOk() (*bool, bool)`

GetConsumerAckPropagationEnabledOk returns a tuple with the ConsumerAckPropagationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerAckPropagationEnabled

`func (o *MsgVpnTopicEndpoint) SetConsumerAckPropagationEnabled(v bool)`

SetConsumerAckPropagationEnabled sets ConsumerAckPropagationEnabled field to given value.

### HasConsumerAckPropagationEnabled

`func (o *MsgVpnTopicEndpoint) HasConsumerAckPropagationEnabled() bool`

HasConsumerAckPropagationEnabled returns a boolean if a field has been set.

### GetCreatedByManagement

`func (o *MsgVpnTopicEndpoint) GetCreatedByManagement() bool`

GetCreatedByManagement returns the CreatedByManagement field if non-nil, zero value otherwise.

### GetCreatedByManagementOk

`func (o *MsgVpnTopicEndpoint) GetCreatedByManagementOk() (*bool, bool)`

GetCreatedByManagementOk returns a tuple with the CreatedByManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByManagement

`func (o *MsgVpnTopicEndpoint) SetCreatedByManagement(v bool)`

SetCreatedByManagement sets CreatedByManagement field to given value.

### HasCreatedByManagement

`func (o *MsgVpnTopicEndpoint) HasCreatedByManagement() bool`

HasCreatedByManagement returns a boolean if a field has been set.

### GetDeadMsgQueue

`func (o *MsgVpnTopicEndpoint) GetDeadMsgQueue() string`

GetDeadMsgQueue returns the DeadMsgQueue field if non-nil, zero value otherwise.

### GetDeadMsgQueueOk

`func (o *MsgVpnTopicEndpoint) GetDeadMsgQueueOk() (*string, bool)`

GetDeadMsgQueueOk returns a tuple with the DeadMsgQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadMsgQueue

`func (o *MsgVpnTopicEndpoint) SetDeadMsgQueue(v string)`

SetDeadMsgQueue sets DeadMsgQueue field to given value.

### HasDeadMsgQueue

`func (o *MsgVpnTopicEndpoint) HasDeadMsgQueue() bool`

HasDeadMsgQueue returns a boolean if a field has been set.

### GetDeletedMsgCount

`func (o *MsgVpnTopicEndpoint) GetDeletedMsgCount() int64`

GetDeletedMsgCount returns the DeletedMsgCount field if non-nil, zero value otherwise.

### GetDeletedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetDeletedMsgCountOk() (*int64, bool)`

GetDeletedMsgCountOk returns a tuple with the DeletedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedMsgCount

`func (o *MsgVpnTopicEndpoint) SetDeletedMsgCount(v int64)`

SetDeletedMsgCount sets DeletedMsgCount field to given value.

### HasDeletedMsgCount

`func (o *MsgVpnTopicEndpoint) HasDeletedMsgCount() bool`

HasDeletedMsgCount returns a boolean if a field has been set.

### GetDeliveryCountEnabled

`func (o *MsgVpnTopicEndpoint) GetDeliveryCountEnabled() bool`

GetDeliveryCountEnabled returns the DeliveryCountEnabled field if non-nil, zero value otherwise.

### GetDeliveryCountEnabledOk

`func (o *MsgVpnTopicEndpoint) GetDeliveryCountEnabledOk() (*bool, bool)`

GetDeliveryCountEnabledOk returns a tuple with the DeliveryCountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCountEnabled

`func (o *MsgVpnTopicEndpoint) SetDeliveryCountEnabled(v bool)`

SetDeliveryCountEnabled sets DeliveryCountEnabled field to given value.

### HasDeliveryCountEnabled

`func (o *MsgVpnTopicEndpoint) HasDeliveryCountEnabled() bool`

HasDeliveryCountEnabled returns a boolean if a field has been set.

### GetDestinationGroupErrorDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) GetDestinationGroupErrorDiscardedMsgCount() int64`

GetDestinationGroupErrorDiscardedMsgCount returns the DestinationGroupErrorDiscardedMsgCount field if non-nil, zero value otherwise.

### GetDestinationGroupErrorDiscardedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetDestinationGroupErrorDiscardedMsgCountOk() (*int64, bool)`

GetDestinationGroupErrorDiscardedMsgCountOk returns a tuple with the DestinationGroupErrorDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroupErrorDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) SetDestinationGroupErrorDiscardedMsgCount(v int64)`

SetDestinationGroupErrorDiscardedMsgCount sets DestinationGroupErrorDiscardedMsgCount field to given value.

### HasDestinationGroupErrorDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) HasDestinationGroupErrorDiscardedMsgCount() bool`

HasDestinationGroupErrorDiscardedMsgCount returns a boolean if a field has been set.

### GetDestinationTopic

`func (o *MsgVpnTopicEndpoint) GetDestinationTopic() string`

GetDestinationTopic returns the DestinationTopic field if non-nil, zero value otherwise.

### GetDestinationTopicOk

`func (o *MsgVpnTopicEndpoint) GetDestinationTopicOk() (*string, bool)`

GetDestinationTopicOk returns a tuple with the DestinationTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTopic

`func (o *MsgVpnTopicEndpoint) SetDestinationTopic(v string)`

SetDestinationTopic sets DestinationTopic field to given value.

### HasDestinationTopic

`func (o *MsgVpnTopicEndpoint) HasDestinationTopic() bool`

HasDestinationTopic returns a boolean if a field has been set.

### GetDisabledBindFailureCount

`func (o *MsgVpnTopicEndpoint) GetDisabledBindFailureCount() int64`

GetDisabledBindFailureCount returns the DisabledBindFailureCount field if non-nil, zero value otherwise.

### GetDisabledBindFailureCountOk

`func (o *MsgVpnTopicEndpoint) GetDisabledBindFailureCountOk() (*int64, bool)`

GetDisabledBindFailureCountOk returns a tuple with the DisabledBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledBindFailureCount

`func (o *MsgVpnTopicEndpoint) SetDisabledBindFailureCount(v int64)`

SetDisabledBindFailureCount sets DisabledBindFailureCount field to given value.

### HasDisabledBindFailureCount

`func (o *MsgVpnTopicEndpoint) HasDisabledBindFailureCount() bool`

HasDisabledBindFailureCount returns a boolean if a field has been set.

### GetDisabledDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) GetDisabledDiscardedMsgCount() int64`

GetDisabledDiscardedMsgCount returns the DisabledDiscardedMsgCount field if non-nil, zero value otherwise.

### GetDisabledDiscardedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetDisabledDiscardedMsgCountOk() (*int64, bool)`

GetDisabledDiscardedMsgCountOk returns a tuple with the DisabledDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) SetDisabledDiscardedMsgCount(v int64)`

SetDisabledDiscardedMsgCount sets DisabledDiscardedMsgCount field to given value.

### HasDisabledDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) HasDisabledDiscardedMsgCount() bool`

HasDisabledDiscardedMsgCount returns a boolean if a field has been set.

### GetDurable

`func (o *MsgVpnTopicEndpoint) GetDurable() bool`

GetDurable returns the Durable field if non-nil, zero value otherwise.

### GetDurableOk

`func (o *MsgVpnTopicEndpoint) GetDurableOk() (*bool, bool)`

GetDurableOk returns a tuple with the Durable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurable

`func (o *MsgVpnTopicEndpoint) SetDurable(v bool)`

SetDurable sets Durable field to given value.

### HasDurable

`func (o *MsgVpnTopicEndpoint) HasDurable() bool`

HasDurable returns a boolean if a field has been set.

### GetEgressEnabled

`func (o *MsgVpnTopicEndpoint) GetEgressEnabled() bool`

GetEgressEnabled returns the EgressEnabled field if non-nil, zero value otherwise.

### GetEgressEnabledOk

`func (o *MsgVpnTopicEndpoint) GetEgressEnabledOk() (*bool, bool)`

GetEgressEnabledOk returns a tuple with the EgressEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressEnabled

`func (o *MsgVpnTopicEndpoint) SetEgressEnabled(v bool)`

SetEgressEnabled sets EgressEnabled field to given value.

### HasEgressEnabled

`func (o *MsgVpnTopicEndpoint) HasEgressEnabled() bool`

HasEgressEnabled returns a boolean if a field has been set.

### GetEventBindCountThreshold

`func (o *MsgVpnTopicEndpoint) GetEventBindCountThreshold() EventThreshold`

GetEventBindCountThreshold returns the EventBindCountThreshold field if non-nil, zero value otherwise.

### GetEventBindCountThresholdOk

`func (o *MsgVpnTopicEndpoint) GetEventBindCountThresholdOk() (*EventThreshold, bool)`

GetEventBindCountThresholdOk returns a tuple with the EventBindCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventBindCountThreshold

`func (o *MsgVpnTopicEndpoint) SetEventBindCountThreshold(v EventThreshold)`

SetEventBindCountThreshold sets EventBindCountThreshold field to given value.

### HasEventBindCountThreshold

`func (o *MsgVpnTopicEndpoint) HasEventBindCountThreshold() bool`

HasEventBindCountThreshold returns a boolean if a field has been set.

### GetEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnTopicEndpoint) GetEventRejectLowPriorityMsgLimitThreshold() EventThreshold`

GetEventRejectLowPriorityMsgLimitThreshold returns the EventRejectLowPriorityMsgLimitThreshold field if non-nil, zero value otherwise.

### GetEventRejectLowPriorityMsgLimitThresholdOk

`func (o *MsgVpnTopicEndpoint) GetEventRejectLowPriorityMsgLimitThresholdOk() (*EventThreshold, bool)`

GetEventRejectLowPriorityMsgLimitThresholdOk returns a tuple with the EventRejectLowPriorityMsgLimitThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnTopicEndpoint) SetEventRejectLowPriorityMsgLimitThreshold(v EventThreshold)`

SetEventRejectLowPriorityMsgLimitThreshold sets EventRejectLowPriorityMsgLimitThreshold field to given value.

### HasEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnTopicEndpoint) HasEventRejectLowPriorityMsgLimitThreshold() bool`

HasEventRejectLowPriorityMsgLimitThreshold returns a boolean if a field has been set.

### GetEventSpoolUsageThreshold

`func (o *MsgVpnTopicEndpoint) GetEventSpoolUsageThreshold() EventThreshold`

GetEventSpoolUsageThreshold returns the EventSpoolUsageThreshold field if non-nil, zero value otherwise.

### GetEventSpoolUsageThresholdOk

`func (o *MsgVpnTopicEndpoint) GetEventSpoolUsageThresholdOk() (*EventThreshold, bool)`

GetEventSpoolUsageThresholdOk returns a tuple with the EventSpoolUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSpoolUsageThreshold

`func (o *MsgVpnTopicEndpoint) SetEventSpoolUsageThreshold(v EventThreshold)`

SetEventSpoolUsageThreshold sets EventSpoolUsageThreshold field to given value.

### HasEventSpoolUsageThreshold

`func (o *MsgVpnTopicEndpoint) HasEventSpoolUsageThreshold() bool`

HasEventSpoolUsageThreshold returns a boolean if a field has been set.

### GetHighestAckedMsgId

`func (o *MsgVpnTopicEndpoint) GetHighestAckedMsgId() int64`

GetHighestAckedMsgId returns the HighestAckedMsgId field if non-nil, zero value otherwise.

### GetHighestAckedMsgIdOk

`func (o *MsgVpnTopicEndpoint) GetHighestAckedMsgIdOk() (*int64, bool)`

GetHighestAckedMsgIdOk returns a tuple with the HighestAckedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestAckedMsgId

`func (o *MsgVpnTopicEndpoint) SetHighestAckedMsgId(v int64)`

SetHighestAckedMsgId sets HighestAckedMsgId field to given value.

### HasHighestAckedMsgId

`func (o *MsgVpnTopicEndpoint) HasHighestAckedMsgId() bool`

HasHighestAckedMsgId returns a boolean if a field has been set.

### GetHighestMsgId

`func (o *MsgVpnTopicEndpoint) GetHighestMsgId() int64`

GetHighestMsgId returns the HighestMsgId field if non-nil, zero value otherwise.

### GetHighestMsgIdOk

`func (o *MsgVpnTopicEndpoint) GetHighestMsgIdOk() (*int64, bool)`

GetHighestMsgIdOk returns a tuple with the HighestMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestMsgId

`func (o *MsgVpnTopicEndpoint) SetHighestMsgId(v int64)`

SetHighestMsgId sets HighestMsgId field to given value.

### HasHighestMsgId

`func (o *MsgVpnTopicEndpoint) HasHighestMsgId() bool`

HasHighestMsgId returns a boolean if a field has been set.

### GetInProgressAckMsgCount

`func (o *MsgVpnTopicEndpoint) GetInProgressAckMsgCount() int64`

GetInProgressAckMsgCount returns the InProgressAckMsgCount field if non-nil, zero value otherwise.

### GetInProgressAckMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetInProgressAckMsgCountOk() (*int64, bool)`

GetInProgressAckMsgCountOk returns a tuple with the InProgressAckMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgressAckMsgCount

`func (o *MsgVpnTopicEndpoint) SetInProgressAckMsgCount(v int64)`

SetInProgressAckMsgCount sets InProgressAckMsgCount field to given value.

### HasInProgressAckMsgCount

`func (o *MsgVpnTopicEndpoint) HasInProgressAckMsgCount() bool`

HasInProgressAckMsgCount returns a boolean if a field has been set.

### GetIngressEnabled

`func (o *MsgVpnTopicEndpoint) GetIngressEnabled() bool`

GetIngressEnabled returns the IngressEnabled field if non-nil, zero value otherwise.

### GetIngressEnabledOk

`func (o *MsgVpnTopicEndpoint) GetIngressEnabledOk() (*bool, bool)`

GetIngressEnabledOk returns a tuple with the IngressEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressEnabled

`func (o *MsgVpnTopicEndpoint) SetIngressEnabled(v bool)`

SetIngressEnabled sets IngressEnabled field to given value.

### HasIngressEnabled

`func (o *MsgVpnTopicEndpoint) HasIngressEnabled() bool`

HasIngressEnabled returns a boolean if a field has been set.

### GetInvalidSelectorBindFailureCount

`func (o *MsgVpnTopicEndpoint) GetInvalidSelectorBindFailureCount() int64`

GetInvalidSelectorBindFailureCount returns the InvalidSelectorBindFailureCount field if non-nil, zero value otherwise.

### GetInvalidSelectorBindFailureCountOk

`func (o *MsgVpnTopicEndpoint) GetInvalidSelectorBindFailureCountOk() (*int64, bool)`

GetInvalidSelectorBindFailureCountOk returns a tuple with the InvalidSelectorBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidSelectorBindFailureCount

`func (o *MsgVpnTopicEndpoint) SetInvalidSelectorBindFailureCount(v int64)`

SetInvalidSelectorBindFailureCount sets InvalidSelectorBindFailureCount field to given value.

### HasInvalidSelectorBindFailureCount

`func (o *MsgVpnTopicEndpoint) HasInvalidSelectorBindFailureCount() bool`

HasInvalidSelectorBindFailureCount returns a boolean if a field has been set.

### GetLastReplayCompleteTime

`func (o *MsgVpnTopicEndpoint) GetLastReplayCompleteTime() int32`

GetLastReplayCompleteTime returns the LastReplayCompleteTime field if non-nil, zero value otherwise.

### GetLastReplayCompleteTimeOk

`func (o *MsgVpnTopicEndpoint) GetLastReplayCompleteTimeOk() (*int32, bool)`

GetLastReplayCompleteTimeOk returns a tuple with the LastReplayCompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplayCompleteTime

`func (o *MsgVpnTopicEndpoint) SetLastReplayCompleteTime(v int32)`

SetLastReplayCompleteTime sets LastReplayCompleteTime field to given value.

### HasLastReplayCompleteTime

`func (o *MsgVpnTopicEndpoint) HasLastReplayCompleteTime() bool`

HasLastReplayCompleteTime returns a boolean if a field has been set.

### GetLastReplayFailureReason

`func (o *MsgVpnTopicEndpoint) GetLastReplayFailureReason() string`

GetLastReplayFailureReason returns the LastReplayFailureReason field if non-nil, zero value otherwise.

### GetLastReplayFailureReasonOk

`func (o *MsgVpnTopicEndpoint) GetLastReplayFailureReasonOk() (*string, bool)`

GetLastReplayFailureReasonOk returns a tuple with the LastReplayFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplayFailureReason

`func (o *MsgVpnTopicEndpoint) SetLastReplayFailureReason(v string)`

SetLastReplayFailureReason sets LastReplayFailureReason field to given value.

### HasLastReplayFailureReason

`func (o *MsgVpnTopicEndpoint) HasLastReplayFailureReason() bool`

HasLastReplayFailureReason returns a boolean if a field has been set.

### GetLastReplayFailureTime

`func (o *MsgVpnTopicEndpoint) GetLastReplayFailureTime() int32`

GetLastReplayFailureTime returns the LastReplayFailureTime field if non-nil, zero value otherwise.

### GetLastReplayFailureTimeOk

`func (o *MsgVpnTopicEndpoint) GetLastReplayFailureTimeOk() (*int32, bool)`

GetLastReplayFailureTimeOk returns a tuple with the LastReplayFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplayFailureTime

`func (o *MsgVpnTopicEndpoint) SetLastReplayFailureTime(v int32)`

SetLastReplayFailureTime sets LastReplayFailureTime field to given value.

### HasLastReplayFailureTime

`func (o *MsgVpnTopicEndpoint) HasLastReplayFailureTime() bool`

HasLastReplayFailureTime returns a boolean if a field has been set.

### GetLastReplayStartTime

`func (o *MsgVpnTopicEndpoint) GetLastReplayStartTime() int32`

GetLastReplayStartTime returns the LastReplayStartTime field if non-nil, zero value otherwise.

### GetLastReplayStartTimeOk

`func (o *MsgVpnTopicEndpoint) GetLastReplayStartTimeOk() (*int32, bool)`

GetLastReplayStartTimeOk returns a tuple with the LastReplayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplayStartTime

`func (o *MsgVpnTopicEndpoint) SetLastReplayStartTime(v int32)`

SetLastReplayStartTime sets LastReplayStartTime field to given value.

### HasLastReplayStartTime

`func (o *MsgVpnTopicEndpoint) HasLastReplayStartTime() bool`

HasLastReplayStartTime returns a boolean if a field has been set.

### GetLastReplayedMsgTxTime

`func (o *MsgVpnTopicEndpoint) GetLastReplayedMsgTxTime() int32`

GetLastReplayedMsgTxTime returns the LastReplayedMsgTxTime field if non-nil, zero value otherwise.

### GetLastReplayedMsgTxTimeOk

`func (o *MsgVpnTopicEndpoint) GetLastReplayedMsgTxTimeOk() (*int32, bool)`

GetLastReplayedMsgTxTimeOk returns a tuple with the LastReplayedMsgTxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplayedMsgTxTime

`func (o *MsgVpnTopicEndpoint) SetLastReplayedMsgTxTime(v int32)`

SetLastReplayedMsgTxTime sets LastReplayedMsgTxTime field to given value.

### HasLastReplayedMsgTxTime

`func (o *MsgVpnTopicEndpoint) HasLastReplayedMsgTxTime() bool`

HasLastReplayedMsgTxTime returns a boolean if a field has been set.

### GetLastSelectorExaminedMsgId

`func (o *MsgVpnTopicEndpoint) GetLastSelectorExaminedMsgId() int64`

GetLastSelectorExaminedMsgId returns the LastSelectorExaminedMsgId field if non-nil, zero value otherwise.

### GetLastSelectorExaminedMsgIdOk

`func (o *MsgVpnTopicEndpoint) GetLastSelectorExaminedMsgIdOk() (*int64, bool)`

GetLastSelectorExaminedMsgIdOk returns a tuple with the LastSelectorExaminedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSelectorExaminedMsgId

`func (o *MsgVpnTopicEndpoint) SetLastSelectorExaminedMsgId(v int64)`

SetLastSelectorExaminedMsgId sets LastSelectorExaminedMsgId field to given value.

### HasLastSelectorExaminedMsgId

`func (o *MsgVpnTopicEndpoint) HasLastSelectorExaminedMsgId() bool`

HasLastSelectorExaminedMsgId returns a boolean if a field has been set.

### GetLastSpooledMsgId

`func (o *MsgVpnTopicEndpoint) GetLastSpooledMsgId() int64`

GetLastSpooledMsgId returns the LastSpooledMsgId field if non-nil, zero value otherwise.

### GetLastSpooledMsgIdOk

`func (o *MsgVpnTopicEndpoint) GetLastSpooledMsgIdOk() (*int64, bool)`

GetLastSpooledMsgIdOk returns a tuple with the LastSpooledMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSpooledMsgId

`func (o *MsgVpnTopicEndpoint) SetLastSpooledMsgId(v int64)`

SetLastSpooledMsgId sets LastSpooledMsgId field to given value.

### HasLastSpooledMsgId

`func (o *MsgVpnTopicEndpoint) HasLastSpooledMsgId() bool`

HasLastSpooledMsgId returns a boolean if a field has been set.

### GetLowPriorityMsgCongestionDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) GetLowPriorityMsgCongestionDiscardedMsgCount() int64`

GetLowPriorityMsgCongestionDiscardedMsgCount returns the LowPriorityMsgCongestionDiscardedMsgCount field if non-nil, zero value otherwise.

### GetLowPriorityMsgCongestionDiscardedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetLowPriorityMsgCongestionDiscardedMsgCountOk() (*int64, bool)`

GetLowPriorityMsgCongestionDiscardedMsgCountOk returns a tuple with the LowPriorityMsgCongestionDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPriorityMsgCongestionDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) SetLowPriorityMsgCongestionDiscardedMsgCount(v int64)`

SetLowPriorityMsgCongestionDiscardedMsgCount sets LowPriorityMsgCongestionDiscardedMsgCount field to given value.

### HasLowPriorityMsgCongestionDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) HasLowPriorityMsgCongestionDiscardedMsgCount() bool`

HasLowPriorityMsgCongestionDiscardedMsgCount returns a boolean if a field has been set.

### GetLowPriorityMsgCongestionState

`func (o *MsgVpnTopicEndpoint) GetLowPriorityMsgCongestionState() string`

GetLowPriorityMsgCongestionState returns the LowPriorityMsgCongestionState field if non-nil, zero value otherwise.

### GetLowPriorityMsgCongestionStateOk

`func (o *MsgVpnTopicEndpoint) GetLowPriorityMsgCongestionStateOk() (*string, bool)`

GetLowPriorityMsgCongestionStateOk returns a tuple with the LowPriorityMsgCongestionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPriorityMsgCongestionState

`func (o *MsgVpnTopicEndpoint) SetLowPriorityMsgCongestionState(v string)`

SetLowPriorityMsgCongestionState sets LowPriorityMsgCongestionState field to given value.

### HasLowPriorityMsgCongestionState

`func (o *MsgVpnTopicEndpoint) HasLowPriorityMsgCongestionState() bool`

HasLowPriorityMsgCongestionState returns a boolean if a field has been set.

### GetLowestAckedMsgId

`func (o *MsgVpnTopicEndpoint) GetLowestAckedMsgId() int64`

GetLowestAckedMsgId returns the LowestAckedMsgId field if non-nil, zero value otherwise.

### GetLowestAckedMsgIdOk

`func (o *MsgVpnTopicEndpoint) GetLowestAckedMsgIdOk() (*int64, bool)`

GetLowestAckedMsgIdOk returns a tuple with the LowestAckedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestAckedMsgId

`func (o *MsgVpnTopicEndpoint) SetLowestAckedMsgId(v int64)`

SetLowestAckedMsgId sets LowestAckedMsgId field to given value.

### HasLowestAckedMsgId

`func (o *MsgVpnTopicEndpoint) HasLowestAckedMsgId() bool`

HasLowestAckedMsgId returns a boolean if a field has been set.

### GetLowestMsgId

`func (o *MsgVpnTopicEndpoint) GetLowestMsgId() int64`

GetLowestMsgId returns the LowestMsgId field if non-nil, zero value otherwise.

### GetLowestMsgIdOk

`func (o *MsgVpnTopicEndpoint) GetLowestMsgIdOk() (*int64, bool)`

GetLowestMsgIdOk returns a tuple with the LowestMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestMsgId

`func (o *MsgVpnTopicEndpoint) SetLowestMsgId(v int64)`

SetLowestMsgId sets LowestMsgId field to given value.

### HasLowestMsgId

`func (o *MsgVpnTopicEndpoint) HasLowestMsgId() bool`

HasLowestMsgId returns a boolean if a field has been set.

### GetMaxBindCount

`func (o *MsgVpnTopicEndpoint) GetMaxBindCount() int64`

GetMaxBindCount returns the MaxBindCount field if non-nil, zero value otherwise.

### GetMaxBindCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxBindCountOk() (*int64, bool)`

GetMaxBindCountOk returns a tuple with the MaxBindCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBindCount

`func (o *MsgVpnTopicEndpoint) SetMaxBindCount(v int64)`

SetMaxBindCount sets MaxBindCount field to given value.

### HasMaxBindCount

`func (o *MsgVpnTopicEndpoint) HasMaxBindCount() bool`

HasMaxBindCount returns a boolean if a field has been set.

### GetMaxBindCountExceededBindFailureCount

`func (o *MsgVpnTopicEndpoint) GetMaxBindCountExceededBindFailureCount() int64`

GetMaxBindCountExceededBindFailureCount returns the MaxBindCountExceededBindFailureCount field if non-nil, zero value otherwise.

### GetMaxBindCountExceededBindFailureCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxBindCountExceededBindFailureCountOk() (*int64, bool)`

GetMaxBindCountExceededBindFailureCountOk returns a tuple with the MaxBindCountExceededBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBindCountExceededBindFailureCount

`func (o *MsgVpnTopicEndpoint) SetMaxBindCountExceededBindFailureCount(v int64)`

SetMaxBindCountExceededBindFailureCount sets MaxBindCountExceededBindFailureCount field to given value.

### HasMaxBindCountExceededBindFailureCount

`func (o *MsgVpnTopicEndpoint) HasMaxBindCountExceededBindFailureCount() bool`

HasMaxBindCountExceededBindFailureCount returns a boolean if a field has been set.

### GetMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnTopicEndpoint) GetMaxDeliveredUnackedMsgsPerFlow() int64`

GetMaxDeliveredUnackedMsgsPerFlow returns the MaxDeliveredUnackedMsgsPerFlow field if non-nil, zero value otherwise.

### GetMaxDeliveredUnackedMsgsPerFlowOk

`func (o *MsgVpnTopicEndpoint) GetMaxDeliveredUnackedMsgsPerFlowOk() (*int64, bool)`

GetMaxDeliveredUnackedMsgsPerFlowOk returns a tuple with the MaxDeliveredUnackedMsgsPerFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnTopicEndpoint) SetMaxDeliveredUnackedMsgsPerFlow(v int64)`

SetMaxDeliveredUnackedMsgsPerFlow sets MaxDeliveredUnackedMsgsPerFlow field to given value.

### HasMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnTopicEndpoint) HasMaxDeliveredUnackedMsgsPerFlow() bool`

HasMaxDeliveredUnackedMsgsPerFlow returns a boolean if a field has been set.

### GetMaxEffectiveBindCount

`func (o *MsgVpnTopicEndpoint) GetMaxEffectiveBindCount() int32`

GetMaxEffectiveBindCount returns the MaxEffectiveBindCount field if non-nil, zero value otherwise.

### GetMaxEffectiveBindCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxEffectiveBindCountOk() (*int32, bool)`

GetMaxEffectiveBindCountOk returns a tuple with the MaxEffectiveBindCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEffectiveBindCount

`func (o *MsgVpnTopicEndpoint) SetMaxEffectiveBindCount(v int32)`

SetMaxEffectiveBindCount sets MaxEffectiveBindCount field to given value.

### HasMaxEffectiveBindCount

`func (o *MsgVpnTopicEndpoint) HasMaxEffectiveBindCount() bool`

HasMaxEffectiveBindCount returns a boolean if a field has been set.

### GetMaxMsgSize

`func (o *MsgVpnTopicEndpoint) GetMaxMsgSize() int32`

GetMaxMsgSize returns the MaxMsgSize field if non-nil, zero value otherwise.

### GetMaxMsgSizeOk

`func (o *MsgVpnTopicEndpoint) GetMaxMsgSizeOk() (*int32, bool)`

GetMaxMsgSizeOk returns a tuple with the MaxMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSize

`func (o *MsgVpnTopicEndpoint) SetMaxMsgSize(v int32)`

SetMaxMsgSize sets MaxMsgSize field to given value.

### HasMaxMsgSize

`func (o *MsgVpnTopicEndpoint) HasMaxMsgSize() bool`

HasMaxMsgSize returns a boolean if a field has been set.

### GetMaxMsgSizeExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) GetMaxMsgSizeExceededDiscardedMsgCount() int64`

GetMaxMsgSizeExceededDiscardedMsgCount returns the MaxMsgSizeExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxMsgSizeExceededDiscardedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxMsgSizeExceededDiscardedMsgCountOk() (*int64, bool)`

GetMaxMsgSizeExceededDiscardedMsgCountOk returns a tuple with the MaxMsgSizeExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSizeExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) SetMaxMsgSizeExceededDiscardedMsgCount(v int64)`

SetMaxMsgSizeExceededDiscardedMsgCount sets MaxMsgSizeExceededDiscardedMsgCount field to given value.

### HasMaxMsgSizeExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) HasMaxMsgSizeExceededDiscardedMsgCount() bool`

HasMaxMsgSizeExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxMsgSpoolUsageExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) GetMaxMsgSpoolUsageExceededDiscardedMsgCount() int64`

GetMaxMsgSpoolUsageExceededDiscardedMsgCount returns the MaxMsgSpoolUsageExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxMsgSpoolUsageExceededDiscardedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxMsgSpoolUsageExceededDiscardedMsgCountOk() (*int64, bool)`

GetMaxMsgSpoolUsageExceededDiscardedMsgCountOk returns a tuple with the MaxMsgSpoolUsageExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSpoolUsageExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) SetMaxMsgSpoolUsageExceededDiscardedMsgCount(v int64)`

SetMaxMsgSpoolUsageExceededDiscardedMsgCount sets MaxMsgSpoolUsageExceededDiscardedMsgCount field to given value.

### HasMaxMsgSpoolUsageExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) HasMaxMsgSpoolUsageExceededDiscardedMsgCount() bool`

HasMaxMsgSpoolUsageExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxRedeliveryCount

`func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryCount() int64`

GetMaxRedeliveryCount returns the MaxRedeliveryCount field if non-nil, zero value otherwise.

### GetMaxRedeliveryCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryCountOk() (*int64, bool)`

GetMaxRedeliveryCountOk returns a tuple with the MaxRedeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedeliveryCount

`func (o *MsgVpnTopicEndpoint) SetMaxRedeliveryCount(v int64)`

SetMaxRedeliveryCount sets MaxRedeliveryCount field to given value.

### HasMaxRedeliveryCount

`func (o *MsgVpnTopicEndpoint) HasMaxRedeliveryCount() bool`

HasMaxRedeliveryCount returns a boolean if a field has been set.

### GetMaxRedeliveryExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededDiscardedMsgCount() int64`

GetMaxRedeliveryExceededDiscardedMsgCount returns the MaxRedeliveryExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxRedeliveryExceededDiscardedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededDiscardedMsgCountOk() (*int64, bool)`

GetMaxRedeliveryExceededDiscardedMsgCountOk returns a tuple with the MaxRedeliveryExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedeliveryExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) SetMaxRedeliveryExceededDiscardedMsgCount(v int64)`

SetMaxRedeliveryExceededDiscardedMsgCount sets MaxRedeliveryExceededDiscardedMsgCount field to given value.

### HasMaxRedeliveryExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) HasMaxRedeliveryExceededDiscardedMsgCount() bool`

HasMaxRedeliveryExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxRedeliveryExceededToDmqFailedMsgCount

`func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededToDmqFailedMsgCount() int64`

GetMaxRedeliveryExceededToDmqFailedMsgCount returns the MaxRedeliveryExceededToDmqFailedMsgCount field if non-nil, zero value otherwise.

### GetMaxRedeliveryExceededToDmqFailedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededToDmqFailedMsgCountOk() (*int64, bool)`

GetMaxRedeliveryExceededToDmqFailedMsgCountOk returns a tuple with the MaxRedeliveryExceededToDmqFailedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedeliveryExceededToDmqFailedMsgCount

`func (o *MsgVpnTopicEndpoint) SetMaxRedeliveryExceededToDmqFailedMsgCount(v int64)`

SetMaxRedeliveryExceededToDmqFailedMsgCount sets MaxRedeliveryExceededToDmqFailedMsgCount field to given value.

### HasMaxRedeliveryExceededToDmqFailedMsgCount

`func (o *MsgVpnTopicEndpoint) HasMaxRedeliveryExceededToDmqFailedMsgCount() bool`

HasMaxRedeliveryExceededToDmqFailedMsgCount returns a boolean if a field has been set.

### GetMaxRedeliveryExceededToDmqMsgCount

`func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededToDmqMsgCount() int64`

GetMaxRedeliveryExceededToDmqMsgCount returns the MaxRedeliveryExceededToDmqMsgCount field if non-nil, zero value otherwise.

### GetMaxRedeliveryExceededToDmqMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededToDmqMsgCountOk() (*int64, bool)`

GetMaxRedeliveryExceededToDmqMsgCountOk returns a tuple with the MaxRedeliveryExceededToDmqMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedeliveryExceededToDmqMsgCount

`func (o *MsgVpnTopicEndpoint) SetMaxRedeliveryExceededToDmqMsgCount(v int64)`

SetMaxRedeliveryExceededToDmqMsgCount sets MaxRedeliveryExceededToDmqMsgCount field to given value.

### HasMaxRedeliveryExceededToDmqMsgCount

`func (o *MsgVpnTopicEndpoint) HasMaxRedeliveryExceededToDmqMsgCount() bool`

HasMaxRedeliveryExceededToDmqMsgCount returns a boolean if a field has been set.

### GetMaxSpoolUsage

`func (o *MsgVpnTopicEndpoint) GetMaxSpoolUsage() int64`

GetMaxSpoolUsage returns the MaxSpoolUsage field if non-nil, zero value otherwise.

### GetMaxSpoolUsageOk

`func (o *MsgVpnTopicEndpoint) GetMaxSpoolUsageOk() (*int64, bool)`

GetMaxSpoolUsageOk returns a tuple with the MaxSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpoolUsage

`func (o *MsgVpnTopicEndpoint) SetMaxSpoolUsage(v int64)`

SetMaxSpoolUsage sets MaxSpoolUsage field to given value.

### HasMaxSpoolUsage

`func (o *MsgVpnTopicEndpoint) HasMaxSpoolUsage() bool`

HasMaxSpoolUsage returns a boolean if a field has been set.

### GetMaxTtl

`func (o *MsgVpnTopicEndpoint) GetMaxTtl() int64`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *MsgVpnTopicEndpoint) GetMaxTtlOk() (*int64, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *MsgVpnTopicEndpoint) SetMaxTtl(v int64)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *MsgVpnTopicEndpoint) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetMaxTtlExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) GetMaxTtlExceededDiscardedMsgCount() int64`

GetMaxTtlExceededDiscardedMsgCount returns the MaxTtlExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxTtlExceededDiscardedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxTtlExceededDiscardedMsgCountOk() (*int64, bool)`

GetMaxTtlExceededDiscardedMsgCountOk returns a tuple with the MaxTtlExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtlExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) SetMaxTtlExceededDiscardedMsgCount(v int64)`

SetMaxTtlExceededDiscardedMsgCount sets MaxTtlExceededDiscardedMsgCount field to given value.

### HasMaxTtlExceededDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) HasMaxTtlExceededDiscardedMsgCount() bool`

HasMaxTtlExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxTtlExpiredDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredDiscardedMsgCount() int64`

GetMaxTtlExpiredDiscardedMsgCount returns the MaxTtlExpiredDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxTtlExpiredDiscardedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredDiscardedMsgCountOk() (*int64, bool)`

GetMaxTtlExpiredDiscardedMsgCountOk returns a tuple with the MaxTtlExpiredDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtlExpiredDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) SetMaxTtlExpiredDiscardedMsgCount(v int64)`

SetMaxTtlExpiredDiscardedMsgCount sets MaxTtlExpiredDiscardedMsgCount field to given value.

### HasMaxTtlExpiredDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) HasMaxTtlExpiredDiscardedMsgCount() bool`

HasMaxTtlExpiredDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxTtlExpiredToDmqFailedMsgCount

`func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredToDmqFailedMsgCount() int64`

GetMaxTtlExpiredToDmqFailedMsgCount returns the MaxTtlExpiredToDmqFailedMsgCount field if non-nil, zero value otherwise.

### GetMaxTtlExpiredToDmqFailedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredToDmqFailedMsgCountOk() (*int64, bool)`

GetMaxTtlExpiredToDmqFailedMsgCountOk returns a tuple with the MaxTtlExpiredToDmqFailedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtlExpiredToDmqFailedMsgCount

`func (o *MsgVpnTopicEndpoint) SetMaxTtlExpiredToDmqFailedMsgCount(v int64)`

SetMaxTtlExpiredToDmqFailedMsgCount sets MaxTtlExpiredToDmqFailedMsgCount field to given value.

### HasMaxTtlExpiredToDmqFailedMsgCount

`func (o *MsgVpnTopicEndpoint) HasMaxTtlExpiredToDmqFailedMsgCount() bool`

HasMaxTtlExpiredToDmqFailedMsgCount returns a boolean if a field has been set.

### GetMaxTtlExpiredToDmqMsgCount

`func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredToDmqMsgCount() int64`

GetMaxTtlExpiredToDmqMsgCount returns the MaxTtlExpiredToDmqMsgCount field if non-nil, zero value otherwise.

### GetMaxTtlExpiredToDmqMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredToDmqMsgCountOk() (*int64, bool)`

GetMaxTtlExpiredToDmqMsgCountOk returns a tuple with the MaxTtlExpiredToDmqMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtlExpiredToDmqMsgCount

`func (o *MsgVpnTopicEndpoint) SetMaxTtlExpiredToDmqMsgCount(v int64)`

SetMaxTtlExpiredToDmqMsgCount sets MaxTtlExpiredToDmqMsgCount field to given value.

### HasMaxTtlExpiredToDmqMsgCount

`func (o *MsgVpnTopicEndpoint) HasMaxTtlExpiredToDmqMsgCount() bool`

HasMaxTtlExpiredToDmqMsgCount returns a boolean if a field has been set.

### GetMsgSpoolPeakUsage

`func (o *MsgVpnTopicEndpoint) GetMsgSpoolPeakUsage() int64`

GetMsgSpoolPeakUsage returns the MsgSpoolPeakUsage field if non-nil, zero value otherwise.

### GetMsgSpoolPeakUsageOk

`func (o *MsgVpnTopicEndpoint) GetMsgSpoolPeakUsageOk() (*int64, bool)`

GetMsgSpoolPeakUsageOk returns a tuple with the MsgSpoolPeakUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolPeakUsage

`func (o *MsgVpnTopicEndpoint) SetMsgSpoolPeakUsage(v int64)`

SetMsgSpoolPeakUsage sets MsgSpoolPeakUsage field to given value.

### HasMsgSpoolPeakUsage

`func (o *MsgVpnTopicEndpoint) HasMsgSpoolPeakUsage() bool`

HasMsgSpoolPeakUsage returns a boolean if a field has been set.

### GetMsgSpoolUsage

`func (o *MsgVpnTopicEndpoint) GetMsgSpoolUsage() int64`

GetMsgSpoolUsage returns the MsgSpoolUsage field if non-nil, zero value otherwise.

### GetMsgSpoolUsageOk

`func (o *MsgVpnTopicEndpoint) GetMsgSpoolUsageOk() (*int64, bool)`

GetMsgSpoolUsageOk returns a tuple with the MsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolUsage

`func (o *MsgVpnTopicEndpoint) SetMsgSpoolUsage(v int64)`

SetMsgSpoolUsage sets MsgSpoolUsage field to given value.

### HasMsgSpoolUsage

`func (o *MsgVpnTopicEndpoint) HasMsgSpoolUsage() bool`

HasMsgSpoolUsage returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnTopicEndpoint) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnTopicEndpoint) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnTopicEndpoint) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnTopicEndpoint) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetNetworkTopic

`func (o *MsgVpnTopicEndpoint) GetNetworkTopic() string`

GetNetworkTopic returns the NetworkTopic field if non-nil, zero value otherwise.

### GetNetworkTopicOk

`func (o *MsgVpnTopicEndpoint) GetNetworkTopicOk() (*string, bool)`

GetNetworkTopicOk returns a tuple with the NetworkTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTopic

`func (o *MsgVpnTopicEndpoint) SetNetworkTopic(v string)`

SetNetworkTopic sets NetworkTopic field to given value.

### HasNetworkTopic

`func (o *MsgVpnTopicEndpoint) HasNetworkTopic() bool`

HasNetworkTopic returns a boolean if a field has been set.

### GetNoLocalDeliveryDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) GetNoLocalDeliveryDiscardedMsgCount() int64`

GetNoLocalDeliveryDiscardedMsgCount returns the NoLocalDeliveryDiscardedMsgCount field if non-nil, zero value otherwise.

### GetNoLocalDeliveryDiscardedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetNoLocalDeliveryDiscardedMsgCountOk() (*int64, bool)`

GetNoLocalDeliveryDiscardedMsgCountOk returns a tuple with the NoLocalDeliveryDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLocalDeliveryDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) SetNoLocalDeliveryDiscardedMsgCount(v int64)`

SetNoLocalDeliveryDiscardedMsgCount sets NoLocalDeliveryDiscardedMsgCount field to given value.

### HasNoLocalDeliveryDiscardedMsgCount

`func (o *MsgVpnTopicEndpoint) HasNoLocalDeliveryDiscardedMsgCount() bool`

HasNoLocalDeliveryDiscardedMsgCount returns a boolean if a field has been set.

### GetOtherBindFailureCount

`func (o *MsgVpnTopicEndpoint) GetOtherBindFailureCount() int64`

GetOtherBindFailureCount returns the OtherBindFailureCount field if non-nil, zero value otherwise.

### GetOtherBindFailureCountOk

`func (o *MsgVpnTopicEndpoint) GetOtherBindFailureCountOk() (*int64, bool)`

GetOtherBindFailureCountOk returns a tuple with the OtherBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherBindFailureCount

`func (o *MsgVpnTopicEndpoint) SetOtherBindFailureCount(v int64)`

SetOtherBindFailureCount sets OtherBindFailureCount field to given value.

### HasOtherBindFailureCount

`func (o *MsgVpnTopicEndpoint) HasOtherBindFailureCount() bool`

HasOtherBindFailureCount returns a boolean if a field has been set.

### GetOwner

`func (o *MsgVpnTopicEndpoint) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MsgVpnTopicEndpoint) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MsgVpnTopicEndpoint) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MsgVpnTopicEndpoint) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPermission

`func (o *MsgVpnTopicEndpoint) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *MsgVpnTopicEndpoint) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *MsgVpnTopicEndpoint) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *MsgVpnTopicEndpoint) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetRedeliveredMsgCount

`func (o *MsgVpnTopicEndpoint) GetRedeliveredMsgCount() int64`

GetRedeliveredMsgCount returns the RedeliveredMsgCount field if non-nil, zero value otherwise.

### GetRedeliveredMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetRedeliveredMsgCountOk() (*int64, bool)`

GetRedeliveredMsgCountOk returns a tuple with the RedeliveredMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveredMsgCount

`func (o *MsgVpnTopicEndpoint) SetRedeliveredMsgCount(v int64)`

SetRedeliveredMsgCount sets RedeliveredMsgCount field to given value.

### HasRedeliveredMsgCount

`func (o *MsgVpnTopicEndpoint) HasRedeliveredMsgCount() bool`

HasRedeliveredMsgCount returns a boolean if a field has been set.

### GetRedeliveryEnabled

`func (o *MsgVpnTopicEndpoint) GetRedeliveryEnabled() bool`

GetRedeliveryEnabled returns the RedeliveryEnabled field if non-nil, zero value otherwise.

### GetRedeliveryEnabledOk

`func (o *MsgVpnTopicEndpoint) GetRedeliveryEnabledOk() (*bool, bool)`

GetRedeliveryEnabledOk returns a tuple with the RedeliveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveryEnabled

`func (o *MsgVpnTopicEndpoint) SetRedeliveryEnabled(v bool)`

SetRedeliveryEnabled sets RedeliveryEnabled field to given value.

### HasRedeliveryEnabled

`func (o *MsgVpnTopicEndpoint) HasRedeliveryEnabled() bool`

HasRedeliveryEnabled returns a boolean if a field has been set.

### GetRejectLowPriorityMsgEnabled

`func (o *MsgVpnTopicEndpoint) GetRejectLowPriorityMsgEnabled() bool`

GetRejectLowPriorityMsgEnabled returns the RejectLowPriorityMsgEnabled field if non-nil, zero value otherwise.

### GetRejectLowPriorityMsgEnabledOk

`func (o *MsgVpnTopicEndpoint) GetRejectLowPriorityMsgEnabledOk() (*bool, bool)`

GetRejectLowPriorityMsgEnabledOk returns a tuple with the RejectLowPriorityMsgEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectLowPriorityMsgEnabled

`func (o *MsgVpnTopicEndpoint) SetRejectLowPriorityMsgEnabled(v bool)`

SetRejectLowPriorityMsgEnabled sets RejectLowPriorityMsgEnabled field to given value.

### HasRejectLowPriorityMsgEnabled

`func (o *MsgVpnTopicEndpoint) HasRejectLowPriorityMsgEnabled() bool`

HasRejectLowPriorityMsgEnabled returns a boolean if a field has been set.

### GetRejectLowPriorityMsgLimit

`func (o *MsgVpnTopicEndpoint) GetRejectLowPriorityMsgLimit() int64`

GetRejectLowPriorityMsgLimit returns the RejectLowPriorityMsgLimit field if non-nil, zero value otherwise.

### GetRejectLowPriorityMsgLimitOk

`func (o *MsgVpnTopicEndpoint) GetRejectLowPriorityMsgLimitOk() (*int64, bool)`

GetRejectLowPriorityMsgLimitOk returns a tuple with the RejectLowPriorityMsgLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectLowPriorityMsgLimit

`func (o *MsgVpnTopicEndpoint) SetRejectLowPriorityMsgLimit(v int64)`

SetRejectLowPriorityMsgLimit sets RejectLowPriorityMsgLimit field to given value.

### HasRejectLowPriorityMsgLimit

`func (o *MsgVpnTopicEndpoint) HasRejectLowPriorityMsgLimit() bool`

HasRejectLowPriorityMsgLimit returns a boolean if a field has been set.

### GetRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnTopicEndpoint) GetRejectMsgToSenderOnDiscardBehavior() string`

GetRejectMsgToSenderOnDiscardBehavior returns the RejectMsgToSenderOnDiscardBehavior field if non-nil, zero value otherwise.

### GetRejectMsgToSenderOnDiscardBehaviorOk

`func (o *MsgVpnTopicEndpoint) GetRejectMsgToSenderOnDiscardBehaviorOk() (*string, bool)`

GetRejectMsgToSenderOnDiscardBehaviorOk returns a tuple with the RejectMsgToSenderOnDiscardBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnTopicEndpoint) SetRejectMsgToSenderOnDiscardBehavior(v string)`

SetRejectMsgToSenderOnDiscardBehavior sets RejectMsgToSenderOnDiscardBehavior field to given value.

### HasRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnTopicEndpoint) HasRejectMsgToSenderOnDiscardBehavior() bool`

HasRejectMsgToSenderOnDiscardBehavior returns a boolean if a field has been set.

### GetReplayFailureCount

`func (o *MsgVpnTopicEndpoint) GetReplayFailureCount() int64`

GetReplayFailureCount returns the ReplayFailureCount field if non-nil, zero value otherwise.

### GetReplayFailureCountOk

`func (o *MsgVpnTopicEndpoint) GetReplayFailureCountOk() (*int64, bool)`

GetReplayFailureCountOk returns a tuple with the ReplayFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayFailureCount

`func (o *MsgVpnTopicEndpoint) SetReplayFailureCount(v int64)`

SetReplayFailureCount sets ReplayFailureCount field to given value.

### HasReplayFailureCount

`func (o *MsgVpnTopicEndpoint) HasReplayFailureCount() bool`

HasReplayFailureCount returns a boolean if a field has been set.

### GetReplayStartCount

`func (o *MsgVpnTopicEndpoint) GetReplayStartCount() int64`

GetReplayStartCount returns the ReplayStartCount field if non-nil, zero value otherwise.

### GetReplayStartCountOk

`func (o *MsgVpnTopicEndpoint) GetReplayStartCountOk() (*int64, bool)`

GetReplayStartCountOk returns a tuple with the ReplayStartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayStartCount

`func (o *MsgVpnTopicEndpoint) SetReplayStartCount(v int64)`

SetReplayStartCount sets ReplayStartCount field to given value.

### HasReplayStartCount

`func (o *MsgVpnTopicEndpoint) HasReplayStartCount() bool`

HasReplayStartCount returns a boolean if a field has been set.

### GetReplayState

`func (o *MsgVpnTopicEndpoint) GetReplayState() string`

GetReplayState returns the ReplayState field if non-nil, zero value otherwise.

### GetReplayStateOk

`func (o *MsgVpnTopicEndpoint) GetReplayStateOk() (*string, bool)`

GetReplayStateOk returns a tuple with the ReplayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayState

`func (o *MsgVpnTopicEndpoint) SetReplayState(v string)`

SetReplayState sets ReplayState field to given value.

### HasReplayState

`func (o *MsgVpnTopicEndpoint) HasReplayState() bool`

HasReplayState returns a boolean if a field has been set.

### GetReplaySuccessCount

`func (o *MsgVpnTopicEndpoint) GetReplaySuccessCount() int64`

GetReplaySuccessCount returns the ReplaySuccessCount field if non-nil, zero value otherwise.

### GetReplaySuccessCountOk

`func (o *MsgVpnTopicEndpoint) GetReplaySuccessCountOk() (*int64, bool)`

GetReplaySuccessCountOk returns a tuple with the ReplaySuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaySuccessCount

`func (o *MsgVpnTopicEndpoint) SetReplaySuccessCount(v int64)`

SetReplaySuccessCount sets ReplaySuccessCount field to given value.

### HasReplaySuccessCount

`func (o *MsgVpnTopicEndpoint) HasReplaySuccessCount() bool`

HasReplaySuccessCount returns a boolean if a field has been set.

### GetReplayedAckedMsgCount

`func (o *MsgVpnTopicEndpoint) GetReplayedAckedMsgCount() int64`

GetReplayedAckedMsgCount returns the ReplayedAckedMsgCount field if non-nil, zero value otherwise.

### GetReplayedAckedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetReplayedAckedMsgCountOk() (*int64, bool)`

GetReplayedAckedMsgCountOk returns a tuple with the ReplayedAckedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayedAckedMsgCount

`func (o *MsgVpnTopicEndpoint) SetReplayedAckedMsgCount(v int64)`

SetReplayedAckedMsgCount sets ReplayedAckedMsgCount field to given value.

### HasReplayedAckedMsgCount

`func (o *MsgVpnTopicEndpoint) HasReplayedAckedMsgCount() bool`

HasReplayedAckedMsgCount returns a boolean if a field has been set.

### GetReplayedTxMsgCount

`func (o *MsgVpnTopicEndpoint) GetReplayedTxMsgCount() int64`

GetReplayedTxMsgCount returns the ReplayedTxMsgCount field if non-nil, zero value otherwise.

### GetReplayedTxMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetReplayedTxMsgCountOk() (*int64, bool)`

GetReplayedTxMsgCountOk returns a tuple with the ReplayedTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayedTxMsgCount

`func (o *MsgVpnTopicEndpoint) SetReplayedTxMsgCount(v int64)`

SetReplayedTxMsgCount sets ReplayedTxMsgCount field to given value.

### HasReplayedTxMsgCount

`func (o *MsgVpnTopicEndpoint) HasReplayedTxMsgCount() bool`

HasReplayedTxMsgCount returns a boolean if a field has been set.

### GetReplicationActiveAckPropTxMsgCount

`func (o *MsgVpnTopicEndpoint) GetReplicationActiveAckPropTxMsgCount() int64`

GetReplicationActiveAckPropTxMsgCount returns the ReplicationActiveAckPropTxMsgCount field if non-nil, zero value otherwise.

### GetReplicationActiveAckPropTxMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetReplicationActiveAckPropTxMsgCountOk() (*int64, bool)`

GetReplicationActiveAckPropTxMsgCountOk returns a tuple with the ReplicationActiveAckPropTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveAckPropTxMsgCount

`func (o *MsgVpnTopicEndpoint) SetReplicationActiveAckPropTxMsgCount(v int64)`

SetReplicationActiveAckPropTxMsgCount sets ReplicationActiveAckPropTxMsgCount field to given value.

### HasReplicationActiveAckPropTxMsgCount

`func (o *MsgVpnTopicEndpoint) HasReplicationActiveAckPropTxMsgCount() bool`

HasReplicationActiveAckPropTxMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyAckPropRxMsgCount

`func (o *MsgVpnTopicEndpoint) GetReplicationStandbyAckPropRxMsgCount() int64`

GetReplicationStandbyAckPropRxMsgCount returns the ReplicationStandbyAckPropRxMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyAckPropRxMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetReplicationStandbyAckPropRxMsgCountOk() (*int64, bool)`

GetReplicationStandbyAckPropRxMsgCountOk returns a tuple with the ReplicationStandbyAckPropRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyAckPropRxMsgCount

`func (o *MsgVpnTopicEndpoint) SetReplicationStandbyAckPropRxMsgCount(v int64)`

SetReplicationStandbyAckPropRxMsgCount sets ReplicationStandbyAckPropRxMsgCount field to given value.

### HasReplicationStandbyAckPropRxMsgCount

`func (o *MsgVpnTopicEndpoint) HasReplicationStandbyAckPropRxMsgCount() bool`

HasReplicationStandbyAckPropRxMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyAckedByAckPropMsgCount

`func (o *MsgVpnTopicEndpoint) GetReplicationStandbyAckedByAckPropMsgCount() int64`

GetReplicationStandbyAckedByAckPropMsgCount returns the ReplicationStandbyAckedByAckPropMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyAckedByAckPropMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetReplicationStandbyAckedByAckPropMsgCountOk() (*int64, bool)`

GetReplicationStandbyAckedByAckPropMsgCountOk returns a tuple with the ReplicationStandbyAckedByAckPropMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyAckedByAckPropMsgCount

`func (o *MsgVpnTopicEndpoint) SetReplicationStandbyAckedByAckPropMsgCount(v int64)`

SetReplicationStandbyAckedByAckPropMsgCount sets ReplicationStandbyAckedByAckPropMsgCount field to given value.

### HasReplicationStandbyAckedByAckPropMsgCount

`func (o *MsgVpnTopicEndpoint) HasReplicationStandbyAckedByAckPropMsgCount() bool`

HasReplicationStandbyAckedByAckPropMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyRxMsgCount

`func (o *MsgVpnTopicEndpoint) GetReplicationStandbyRxMsgCount() int64`

GetReplicationStandbyRxMsgCount returns the ReplicationStandbyRxMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyRxMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetReplicationStandbyRxMsgCountOk() (*int64, bool)`

GetReplicationStandbyRxMsgCountOk returns a tuple with the ReplicationStandbyRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyRxMsgCount

`func (o *MsgVpnTopicEndpoint) SetReplicationStandbyRxMsgCount(v int64)`

SetReplicationStandbyRxMsgCount sets ReplicationStandbyRxMsgCount field to given value.

### HasReplicationStandbyRxMsgCount

`func (o *MsgVpnTopicEndpoint) HasReplicationStandbyRxMsgCount() bool`

HasReplicationStandbyRxMsgCount returns a boolean if a field has been set.

### GetRespectMsgPriorityEnabled

`func (o *MsgVpnTopicEndpoint) GetRespectMsgPriorityEnabled() bool`

GetRespectMsgPriorityEnabled returns the RespectMsgPriorityEnabled field if non-nil, zero value otherwise.

### GetRespectMsgPriorityEnabledOk

`func (o *MsgVpnTopicEndpoint) GetRespectMsgPriorityEnabledOk() (*bool, bool)`

GetRespectMsgPriorityEnabledOk returns a tuple with the RespectMsgPriorityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespectMsgPriorityEnabled

`func (o *MsgVpnTopicEndpoint) SetRespectMsgPriorityEnabled(v bool)`

SetRespectMsgPriorityEnabled sets RespectMsgPriorityEnabled field to given value.

### HasRespectMsgPriorityEnabled

`func (o *MsgVpnTopicEndpoint) HasRespectMsgPriorityEnabled() bool`

HasRespectMsgPriorityEnabled returns a boolean if a field has been set.

### GetRespectTtlEnabled

`func (o *MsgVpnTopicEndpoint) GetRespectTtlEnabled() bool`

GetRespectTtlEnabled returns the RespectTtlEnabled field if non-nil, zero value otherwise.

### GetRespectTtlEnabledOk

`func (o *MsgVpnTopicEndpoint) GetRespectTtlEnabledOk() (*bool, bool)`

GetRespectTtlEnabledOk returns a tuple with the RespectTtlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespectTtlEnabled

`func (o *MsgVpnTopicEndpoint) SetRespectTtlEnabled(v bool)`

SetRespectTtlEnabled sets RespectTtlEnabled field to given value.

### HasRespectTtlEnabled

`func (o *MsgVpnTopicEndpoint) HasRespectTtlEnabled() bool`

HasRespectTtlEnabled returns a boolean if a field has been set.

### GetRxByteRate

`func (o *MsgVpnTopicEndpoint) GetRxByteRate() int32`

GetRxByteRate returns the RxByteRate field if non-nil, zero value otherwise.

### GetRxByteRateOk

`func (o *MsgVpnTopicEndpoint) GetRxByteRateOk() (*int32, bool)`

GetRxByteRateOk returns a tuple with the RxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteRate

`func (o *MsgVpnTopicEndpoint) SetRxByteRate(v int32)`

SetRxByteRate sets RxByteRate field to given value.

### HasRxByteRate

`func (o *MsgVpnTopicEndpoint) HasRxByteRate() bool`

HasRxByteRate returns a boolean if a field has been set.

### GetRxMsgRate

`func (o *MsgVpnTopicEndpoint) GetRxMsgRate() int64`

GetRxMsgRate returns the RxMsgRate field if non-nil, zero value otherwise.

### GetRxMsgRateOk

`func (o *MsgVpnTopicEndpoint) GetRxMsgRateOk() (*int64, bool)`

GetRxMsgRateOk returns a tuple with the RxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgRate

`func (o *MsgVpnTopicEndpoint) SetRxMsgRate(v int64)`

SetRxMsgRate sets RxMsgRate field to given value.

### HasRxMsgRate

`func (o *MsgVpnTopicEndpoint) HasRxMsgRate() bool`

HasRxMsgRate returns a boolean if a field has been set.

### GetRxSelector

`func (o *MsgVpnTopicEndpoint) GetRxSelector() bool`

GetRxSelector returns the RxSelector field if non-nil, zero value otherwise.

### GetRxSelectorOk

`func (o *MsgVpnTopicEndpoint) GetRxSelectorOk() (*bool, bool)`

GetRxSelectorOk returns a tuple with the RxSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxSelector

`func (o *MsgVpnTopicEndpoint) SetRxSelector(v bool)`

SetRxSelector sets RxSelector field to given value.

### HasRxSelector

`func (o *MsgVpnTopicEndpoint) HasRxSelector() bool`

HasRxSelector returns a boolean if a field has been set.

### GetSelector

`func (o *MsgVpnTopicEndpoint) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *MsgVpnTopicEndpoint) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *MsgVpnTopicEndpoint) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *MsgVpnTopicEndpoint) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSelectorExaminedMsgCount

`func (o *MsgVpnTopicEndpoint) GetSelectorExaminedMsgCount() int64`

GetSelectorExaminedMsgCount returns the SelectorExaminedMsgCount field if non-nil, zero value otherwise.

### GetSelectorExaminedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetSelectorExaminedMsgCountOk() (*int64, bool)`

GetSelectorExaminedMsgCountOk returns a tuple with the SelectorExaminedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorExaminedMsgCount

`func (o *MsgVpnTopicEndpoint) SetSelectorExaminedMsgCount(v int64)`

SetSelectorExaminedMsgCount sets SelectorExaminedMsgCount field to given value.

### HasSelectorExaminedMsgCount

`func (o *MsgVpnTopicEndpoint) HasSelectorExaminedMsgCount() bool`

HasSelectorExaminedMsgCount returns a boolean if a field has been set.

### GetSelectorMatchedMsgCount

`func (o *MsgVpnTopicEndpoint) GetSelectorMatchedMsgCount() int64`

GetSelectorMatchedMsgCount returns the SelectorMatchedMsgCount field if non-nil, zero value otherwise.

### GetSelectorMatchedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetSelectorMatchedMsgCountOk() (*int64, bool)`

GetSelectorMatchedMsgCountOk returns a tuple with the SelectorMatchedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorMatchedMsgCount

`func (o *MsgVpnTopicEndpoint) SetSelectorMatchedMsgCount(v int64)`

SetSelectorMatchedMsgCount sets SelectorMatchedMsgCount field to given value.

### HasSelectorMatchedMsgCount

`func (o *MsgVpnTopicEndpoint) HasSelectorMatchedMsgCount() bool`

HasSelectorMatchedMsgCount returns a boolean if a field has been set.

### GetSelectorNotMatchedMsgCount

`func (o *MsgVpnTopicEndpoint) GetSelectorNotMatchedMsgCount() int64`

GetSelectorNotMatchedMsgCount returns the SelectorNotMatchedMsgCount field if non-nil, zero value otherwise.

### GetSelectorNotMatchedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetSelectorNotMatchedMsgCountOk() (*int64, bool)`

GetSelectorNotMatchedMsgCountOk returns a tuple with the SelectorNotMatchedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorNotMatchedMsgCount

`func (o *MsgVpnTopicEndpoint) SetSelectorNotMatchedMsgCount(v int64)`

SetSelectorNotMatchedMsgCount sets SelectorNotMatchedMsgCount field to given value.

### HasSelectorNotMatchedMsgCount

`func (o *MsgVpnTopicEndpoint) HasSelectorNotMatchedMsgCount() bool`

HasSelectorNotMatchedMsgCount returns a boolean if a field has been set.

### GetSpooledByteCount

`func (o *MsgVpnTopicEndpoint) GetSpooledByteCount() int64`

GetSpooledByteCount returns the SpooledByteCount field if non-nil, zero value otherwise.

### GetSpooledByteCountOk

`func (o *MsgVpnTopicEndpoint) GetSpooledByteCountOk() (*int64, bool)`

GetSpooledByteCountOk returns a tuple with the SpooledByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledByteCount

`func (o *MsgVpnTopicEndpoint) SetSpooledByteCount(v int64)`

SetSpooledByteCount sets SpooledByteCount field to given value.

### HasSpooledByteCount

`func (o *MsgVpnTopicEndpoint) HasSpooledByteCount() bool`

HasSpooledByteCount returns a boolean if a field has been set.

### GetSpooledMsgCount

`func (o *MsgVpnTopicEndpoint) GetSpooledMsgCount() int64`

GetSpooledMsgCount returns the SpooledMsgCount field if non-nil, zero value otherwise.

### GetSpooledMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetSpooledMsgCountOk() (*int64, bool)`

GetSpooledMsgCountOk returns a tuple with the SpooledMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledMsgCount

`func (o *MsgVpnTopicEndpoint) SetSpooledMsgCount(v int64)`

SetSpooledMsgCount sets SpooledMsgCount field to given value.

### HasSpooledMsgCount

`func (o *MsgVpnTopicEndpoint) HasSpooledMsgCount() bool`

HasSpooledMsgCount returns a boolean if a field has been set.

### GetTopicEndpointName

`func (o *MsgVpnTopicEndpoint) GetTopicEndpointName() string`

GetTopicEndpointName returns the TopicEndpointName field if non-nil, zero value otherwise.

### GetTopicEndpointNameOk

`func (o *MsgVpnTopicEndpoint) GetTopicEndpointNameOk() (*string, bool)`

GetTopicEndpointNameOk returns a tuple with the TopicEndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicEndpointName

`func (o *MsgVpnTopicEndpoint) SetTopicEndpointName(v string)`

SetTopicEndpointName sets TopicEndpointName field to given value.

### HasTopicEndpointName

`func (o *MsgVpnTopicEndpoint) HasTopicEndpointName() bool`

HasTopicEndpointName returns a boolean if a field has been set.

### GetTransportRetransmitMsgCount

`func (o *MsgVpnTopicEndpoint) GetTransportRetransmitMsgCount() int64`

GetTransportRetransmitMsgCount returns the TransportRetransmitMsgCount field if non-nil, zero value otherwise.

### GetTransportRetransmitMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetTransportRetransmitMsgCountOk() (*int64, bool)`

GetTransportRetransmitMsgCountOk returns a tuple with the TransportRetransmitMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportRetransmitMsgCount

`func (o *MsgVpnTopicEndpoint) SetTransportRetransmitMsgCount(v int64)`

SetTransportRetransmitMsgCount sets TransportRetransmitMsgCount field to given value.

### HasTransportRetransmitMsgCount

`func (o *MsgVpnTopicEndpoint) HasTransportRetransmitMsgCount() bool`

HasTransportRetransmitMsgCount returns a boolean if a field has been set.

### GetTxByteRate

`func (o *MsgVpnTopicEndpoint) GetTxByteRate() int64`

GetTxByteRate returns the TxByteRate field if non-nil, zero value otherwise.

### GetTxByteRateOk

`func (o *MsgVpnTopicEndpoint) GetTxByteRateOk() (*int64, bool)`

GetTxByteRateOk returns a tuple with the TxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteRate

`func (o *MsgVpnTopicEndpoint) SetTxByteRate(v int64)`

SetTxByteRate sets TxByteRate field to given value.

### HasTxByteRate

`func (o *MsgVpnTopicEndpoint) HasTxByteRate() bool`

HasTxByteRate returns a boolean if a field has been set.

### GetTxMsgRate

`func (o *MsgVpnTopicEndpoint) GetTxMsgRate() int64`

GetTxMsgRate returns the TxMsgRate field if non-nil, zero value otherwise.

### GetTxMsgRateOk

`func (o *MsgVpnTopicEndpoint) GetTxMsgRateOk() (*int64, bool)`

GetTxMsgRateOk returns a tuple with the TxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgRate

`func (o *MsgVpnTopicEndpoint) SetTxMsgRate(v int64)`

SetTxMsgRate sets TxMsgRate field to given value.

### HasTxMsgRate

`func (o *MsgVpnTopicEndpoint) HasTxMsgRate() bool`

HasTxMsgRate returns a boolean if a field has been set.

### GetTxUnackedMsgCount

`func (o *MsgVpnTopicEndpoint) GetTxUnackedMsgCount() int64`

GetTxUnackedMsgCount returns the TxUnackedMsgCount field if non-nil, zero value otherwise.

### GetTxUnackedMsgCountOk

`func (o *MsgVpnTopicEndpoint) GetTxUnackedMsgCountOk() (*int64, bool)`

GetTxUnackedMsgCountOk returns a tuple with the TxUnackedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxUnackedMsgCount

`func (o *MsgVpnTopicEndpoint) SetTxUnackedMsgCount(v int64)`

SetTxUnackedMsgCount sets TxUnackedMsgCount field to given value.

### HasTxUnackedMsgCount

`func (o *MsgVpnTopicEndpoint) HasTxUnackedMsgCount() bool`

HasTxUnackedMsgCount returns a boolean if a field has been set.

### GetVirtualRouter

`func (o *MsgVpnTopicEndpoint) GetVirtualRouter() string`

GetVirtualRouter returns the VirtualRouter field if non-nil, zero value otherwise.

### GetVirtualRouterOk

`func (o *MsgVpnTopicEndpoint) GetVirtualRouterOk() (*string, bool)`

GetVirtualRouterOk returns a tuple with the VirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouter

`func (o *MsgVpnTopicEndpoint) SetVirtualRouter(v string)`

SetVirtualRouter sets VirtualRouter field to given value.

### HasVirtualRouter

`func (o *MsgVpnTopicEndpoint) HasVirtualRouter() bool`

HasVirtualRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


