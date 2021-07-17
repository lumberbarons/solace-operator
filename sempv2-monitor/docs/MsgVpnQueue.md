# MsgVpnQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** | The access type for delivering messages to consumer flows bound to the Queue. The allowed values and their meaning are:  &lt;pre&gt; \&quot;exclusive\&quot; - Exclusive delivery of messages to the first bound consumer flow. \&quot;non-exclusive\&quot; - Non-exclusive delivery of messages to all bound consumer flows in a round-robin fashion. &lt;/pre&gt;  | [optional] 
**AlreadyBoundBindFailureCount** | Pointer to **int64** | The number of Queue bind failures due to being already bound. | [optional] 
**AverageRxByteRate** | Pointer to **int64** | The one minute average of the message rate received by the Queue, in bytes per second (B/sec). | [optional] 
**AverageRxMsgRate** | Pointer to **int64** | The one minute average of the message rate received by the Queue, in messages per second (msg/sec). | [optional] 
**AverageTxByteRate** | Pointer to **int64** | The one minute average of the message rate transmitted by the Queue, in bytes per second (B/sec). | [optional] 
**AverageTxMsgRate** | Pointer to **int64** | The one minute average of the message rate transmitted by the Queue, in messages per second (msg/sec). | [optional] 
**BindRequestCount** | Pointer to **int64** | The number of consumer requests to bind to the Queue. | [optional] 
**BindSuccessCount** | Pointer to **int64** | The number of successful consumer requests to bind to the Queue. | [optional] 
**BindTimeForwardingMode** | Pointer to **string** | The forwarding mode of the Queue at bind time. The allowed values and their meaning are:  &lt;pre&gt; \&quot;store-and-forward\&quot; - Deliver messages using the guaranteed data path. \&quot;cut-through\&quot; - Deliver messages using the direct and guaranteed data paths for lower latency. &lt;/pre&gt;  | [optional] 
**ClientProfileDeniedDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to being denied by the Client Profile. | [optional] 
**ConsumerAckPropagationEnabled** | Pointer to **bool** | Indicates whether the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN is enabled. | [optional] 
**CreatedByManagement** | Pointer to **bool** | Indicates whether the Queue was created by a management API (CLI or SEMP). | [optional] 
**DeadMsgQueue** | Pointer to **string** | The name of the Dead Message Queue (DMQ) used by the Queue. | [optional] 
**DeletedMsgCount** | Pointer to **int64** | The number of guaranteed messages deleted from the Queue. | [optional] 
**DeliveryCountEnabled** | Pointer to **bool** | Enable or disable the ability for client applications to query the message delivery count of messages received from the Queue. This is a controlled availability feature. Please contact Solace to find out if this feature is supported for your use case. Available since 2.19. | [optional] 
**DestinationGroupErrorDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to a destination group error. | [optional] 
**DisabledBindFailureCount** | Pointer to **int64** | The number of Queue bind failures due to being disabled. | [optional] 
**DisabledDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to it being disabled. | [optional] 
**Durable** | Pointer to **bool** | Indicates whether the Queue is durable and not temporary. | [optional] 
**EgressEnabled** | Pointer to **bool** | Indicates whether the transmission of messages from the Queue is enabled. | [optional] 
**EventBindCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventMsgSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventRejectLowPriorityMsgLimitThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**HighestAckedMsgId** | Pointer to **int64** | The highest identifier (ID) of guaranteed messages in the Queue that were acknowledged. | [optional] 
**HighestMsgId** | Pointer to **int64** | The highest identifier (ID) of guaranteed messages in the Queue. | [optional] 
**InProgressAckMsgCount** | Pointer to **int64** | The number of acknowledgement messages received by the Queue that are in the process of updating and deleting associated guaranteed messages. | [optional] 
**IngressEnabled** | Pointer to **bool** | Indicates whether the reception of messages to the Queue is enabled. | [optional] 
**InvalidSelectorBindFailureCount** | Pointer to **int64** | The number of Queue bind failures due to an invalid selector. | [optional] 
**LastReplayCompleteTime** | Pointer to **int32** | The timestamp of the last completed replay for the Queue. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastReplayFailureReason** | Pointer to **string** | The reason for the last replay failure for the Queue. | [optional] 
**LastReplayFailureTime** | Pointer to **int32** | The timestamp of the last replay failure for the Queue. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastReplayStartTime** | Pointer to **int32** | The timestamp of the last replay started for the Queue. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastReplayedMsgTxTime** | Pointer to **int32** | The timestamp of the last replayed message transmitted by the Queue. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastSpooledMsgId** | Pointer to **int64** | The identifier (ID) of the last guaranteed message spooled in the Queue. | [optional] 
**LowPriorityMsgCongestionDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to low priority message congestion control. | [optional] 
**LowPriorityMsgCongestionState** | Pointer to **string** | The state of the low priority message congestion in the Queue. The allowed values and their meaning are:  &lt;pre&gt; \&quot;disabled\&quot; - Messages are not being checked for priority. \&quot;not-congested\&quot; - Low priority messages are being stored and delivered. \&quot;congested\&quot; - Low priority messages are being discarded. &lt;/pre&gt;  | [optional] 
**LowestAckedMsgId** | Pointer to **int64** | The lowest identifier (ID) of guaranteed messages in the Queue that were acknowledged. | [optional] 
**LowestMsgId** | Pointer to **int64** | The lowest identifier (ID) of guaranteed messages in the Queue. | [optional] 
**MaxBindCount** | Pointer to **int64** | The maximum number of consumer flows that can bind to the Queue. | [optional] 
**MaxBindCountExceededBindFailureCount** | Pointer to **int64** | The number of Queue bind failures due to the maximum bind count being exceeded. | [optional] 
**MaxDeliveredUnackedMsgsPerFlow** | Pointer to **int64** | The maximum number of messages delivered but not acknowledged per flow for the Queue. | [optional] 
**MaxMsgSize** | Pointer to **int32** | The maximum message size allowed in the Queue, in bytes (B). | [optional] 
**MaxMsgSizeExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to the maximum message size being exceeded. | [optional] 
**MaxMsgSpoolUsage** | Pointer to **int64** | The maximum message spool usage allowed by the Queue, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. | [optional] 
**MaxMsgSpoolUsageExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to the maximum message spool usage being exceeded. | [optional] 
**MaxRedeliveryCount** | Pointer to **int64** | The maximum number of times the Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. | [optional] 
**MaxRedeliveryExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to the maximum redelivery attempts being exceeded. | [optional] 
**MaxRedeliveryExceededToDmqFailedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to the maximum redelivery attempts being exceeded and failing to move to the Dead Message Queue (DMQ). | [optional] 
**MaxRedeliveryExceededToDmqMsgCount** | Pointer to **int64** | The number of guaranteed messages moved to the Dead Message Queue (DMQ) by the Queue due to the maximum redelivery attempts being exceeded. | [optional] 
**MaxTtl** | Pointer to **int64** | The maximum time in seconds a message can stay in the Queue when &#x60;respectTtlEnabled&#x60; is &#x60;\&quot;true\&quot;&#x60;. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the &#x60;maxTtl&#x60; configured for the Queue, is exceeded. A value of 0 disables expiry. | [optional] 
**MaxTtlExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to the maximum time-to-live (TTL) in hops being exceeded. The TTL hop count is incremented when the message crosses a bridge. | [optional] 
**MaxTtlExpiredDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to the maximum time-to-live (TTL) timestamp expiring. | [optional] 
**MaxTtlExpiredToDmqFailedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to the maximum time-to-live (TTL) timestamp expiring and failing to move to the Dead Message Queue (DMQ). | [optional] 
**MaxTtlExpiredToDmqMsgCount** | Pointer to **int64** | The number of guaranteed messages moved to the Dead Message Queue (DMQ) by the Queue due to the maximum time-to-live (TTL) timestamp expiring. | [optional] 
**MsgSpoolPeakUsage** | Pointer to **int64** | The message spool peak usage by the Queue, in bytes (B). | [optional] 
**MsgSpoolUsage** | Pointer to **int64** | The message spool usage by the Queue, in bytes (B). | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**NetworkTopic** | Pointer to **string** | The name of the network topic for the Queue. | [optional] 
**NoLocalDeliveryDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages discarded by the Queue due to no local delivery being requested. | [optional] 
**OtherBindFailureCount** | Pointer to **int64** | The number of Queue bind failures due to other reasons. | [optional] 
**Owner** | Pointer to **string** | The Client Username that owns the Queue and has permission equivalent to &#x60;\&quot;delete\&quot;&#x60;. | [optional] 
**Permission** | Pointer to **string** | The permission level for all consumers of the Queue, excluding the owner. The allowed values and their meaning are:  &lt;pre&gt; \&quot;no-access\&quot; - Disallows all access. \&quot;read-only\&quot; - Read-only access to the messages. \&quot;consume\&quot; - Consume (read and remove) messages. \&quot;modify-topic\&quot; - Consume messages or modify the topic/selector. \&quot;delete\&quot; - Consume messages, modify the topic/selector or delete the Client created endpoint altogether. &lt;/pre&gt;  | [optional] 
**QueueName** | Pointer to **string** | The name of the Queue. | [optional] 
**RedeliveredMsgCount** | Pointer to **int64** | The number of guaranteed messages transmitted by the Queue for redelivery. | [optional] 
**RedeliveryEnabled** | Pointer to **bool** | Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the queue more than once. Available since 2.18. | [optional] 
**RejectLowPriorityMsgEnabled** | Pointer to **bool** | Indicates whether the checking of low priority messages against the &#x60;rejectLowPriorityMsgLimit&#x60; is enabled. | [optional] 
**RejectLowPriorityMsgLimit** | Pointer to **int64** | The number of messages of any priority in the Queue above which low priority messages are not admitted but higher priority messages are allowed. | [optional] 
**RejectMsgToSenderOnDiscardBehavior** | Pointer to **string** | Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always return a negative acknowledgment (NACK) to the sending client on message discard. \&quot;when-queue-enabled\&quot; - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Queue is enabled. \&quot;never\&quot; - Never return a negative acknowledgment (NACK) to the sending client on message discard. &lt;/pre&gt;  | [optional] 
**ReplayFailureCount** | Pointer to **int64** | The number of replays that failed for the Queue. | [optional] 
**ReplayStartCount** | Pointer to **int64** | The number of replays started for the Queue. | [optional] 
**ReplayState** | Pointer to **string** | The state of replay for the Queue. The allowed values and their meaning are:  &lt;pre&gt; \&quot;initializing\&quot; - All messages are being deleted from the endpoint before replay starts. \&quot;active\&quot; - Subscription matching logged messages are being replayed to the endpoint. \&quot;pending-complete\&quot; - Replay is complete, but final accounting is in progress. \&quot;complete\&quot; - Replay and all related activities are complete. \&quot;failed\&quot; - Replay has failed and is waiting for an unbind response. &lt;/pre&gt;  | [optional] 
**ReplaySuccessCount** | Pointer to **int64** | The number of replays that succeeded for the Queue. | [optional] 
**ReplayedAckedMsgCount** | Pointer to **int64** | The number of replayed messages transmitted by the Queue and acked by all consumers. | [optional] 
**ReplayedTxMsgCount** | Pointer to **int64** | The number of replayed messages transmitted by the Queue. | [optional] 
**ReplicationActiveAckPropTxMsgCount** | Pointer to **int64** | The number of acknowledgement messages propagated by the Queue to the replication standby remote Message VPN. | [optional] 
**ReplicationStandbyAckPropRxMsgCount** | Pointer to **int64** | The number of propagated acknowledgement messages received by the Queue from the replication active remote Message VPN. | [optional] 
**ReplicationStandbyAckedByAckPropMsgCount** | Pointer to **int64** | The number of messages acknowledged in the Queue by acknowledgement propagation from the replication active remote Message VPN. | [optional] 
**ReplicationStandbyRxMsgCount** | Pointer to **int64** | The number of messages received by the Queue from the replication active remote Message VPN. | [optional] 
**RespectMsgPriorityEnabled** | Pointer to **bool** | Indicates whether message priorities are respected. When enabled, messages contained in the Queue are delivered in priority order, from 9 (highest) to 0 (lowest). | [optional] 
**RespectTtlEnabled** | Pointer to **bool** | Indicates whether the the time-to-live (TTL) for messages in the Queue is respected. When enabled, expired messages are discarded or moved to the DMQ. | [optional] 
**RxByteRate** | Pointer to **int64** | The current message rate received by the Queue, in bytes per second (B/sec). | [optional] 
**RxMsgRate** | Pointer to **int64** | The current message rate received by the Queue, in messages per second (msg/sec). | [optional] 
**SpooledByteCount** | Pointer to **int64** | The amount of guaranteed messages that were spooled in the Queue, in bytes (B). | [optional] 
**SpooledMsgCount** | Pointer to **int64** | The number of guaranteed messages that were spooled in the Queue. | [optional] 
**TransportRetransmitMsgCount** | Pointer to **int64** | The number of guaranteed messages that were retransmitted by the Queue at the transport layer as part of a single delivery attempt. Available since 2.18. | [optional] 
**TxByteRate** | Pointer to **int64** | The current message rate transmitted by the Queue, in bytes per second (B/sec). | [optional] 
**TxMsgRate** | Pointer to **int64** | The current message rate transmitted by the Queue, in messages per second (msg/sec). | [optional] 
**TxSelector** | Pointer to **bool** | Indicates whether the Queue has consumers with selectors to filter transmitted messages. | [optional] 
**TxUnackedMsgCount** | Pointer to **int64** | The number of guaranteed messages in the Queue that have been transmitted but not acknowledged by all consumers. | [optional] 
**VirtualRouter** | Pointer to **string** | The virtual router of the Queue. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The endpoint belongs to the primary virtual router. \&quot;backup\&quot; - The endpoint belongs to the backup virtual router. &lt;/pre&gt;  | [optional] 

## Methods

### NewMsgVpnQueue

`func NewMsgVpnQueue() *MsgVpnQueue`

NewMsgVpnQueue instantiates a new MsgVpnQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueueWithDefaults

`func NewMsgVpnQueueWithDefaults() *MsgVpnQueue`

NewMsgVpnQueueWithDefaults instantiates a new MsgVpnQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *MsgVpnQueue) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *MsgVpnQueue) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *MsgVpnQueue) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *MsgVpnQueue) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetAlreadyBoundBindFailureCount

`func (o *MsgVpnQueue) GetAlreadyBoundBindFailureCount() int64`

GetAlreadyBoundBindFailureCount returns the AlreadyBoundBindFailureCount field if non-nil, zero value otherwise.

### GetAlreadyBoundBindFailureCountOk

`func (o *MsgVpnQueue) GetAlreadyBoundBindFailureCountOk() (*int64, bool)`

GetAlreadyBoundBindFailureCountOk returns a tuple with the AlreadyBoundBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyBoundBindFailureCount

`func (o *MsgVpnQueue) SetAlreadyBoundBindFailureCount(v int64)`

SetAlreadyBoundBindFailureCount sets AlreadyBoundBindFailureCount field to given value.

### HasAlreadyBoundBindFailureCount

`func (o *MsgVpnQueue) HasAlreadyBoundBindFailureCount() bool`

HasAlreadyBoundBindFailureCount returns a boolean if a field has been set.

### GetAverageRxByteRate

`func (o *MsgVpnQueue) GetAverageRxByteRate() int64`

GetAverageRxByteRate returns the AverageRxByteRate field if non-nil, zero value otherwise.

### GetAverageRxByteRateOk

`func (o *MsgVpnQueue) GetAverageRxByteRateOk() (*int64, bool)`

GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxByteRate

`func (o *MsgVpnQueue) SetAverageRxByteRate(v int64)`

SetAverageRxByteRate sets AverageRxByteRate field to given value.

### HasAverageRxByteRate

`func (o *MsgVpnQueue) HasAverageRxByteRate() bool`

HasAverageRxByteRate returns a boolean if a field has been set.

### GetAverageRxMsgRate

`func (o *MsgVpnQueue) GetAverageRxMsgRate() int64`

GetAverageRxMsgRate returns the AverageRxMsgRate field if non-nil, zero value otherwise.

### GetAverageRxMsgRateOk

`func (o *MsgVpnQueue) GetAverageRxMsgRateOk() (*int64, bool)`

GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxMsgRate

`func (o *MsgVpnQueue) SetAverageRxMsgRate(v int64)`

SetAverageRxMsgRate sets AverageRxMsgRate field to given value.

### HasAverageRxMsgRate

`func (o *MsgVpnQueue) HasAverageRxMsgRate() bool`

HasAverageRxMsgRate returns a boolean if a field has been set.

### GetAverageTxByteRate

`func (o *MsgVpnQueue) GetAverageTxByteRate() int64`

GetAverageTxByteRate returns the AverageTxByteRate field if non-nil, zero value otherwise.

### GetAverageTxByteRateOk

`func (o *MsgVpnQueue) GetAverageTxByteRateOk() (*int64, bool)`

GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxByteRate

`func (o *MsgVpnQueue) SetAverageTxByteRate(v int64)`

SetAverageTxByteRate sets AverageTxByteRate field to given value.

### HasAverageTxByteRate

`func (o *MsgVpnQueue) HasAverageTxByteRate() bool`

HasAverageTxByteRate returns a boolean if a field has been set.

### GetAverageTxMsgRate

`func (o *MsgVpnQueue) GetAverageTxMsgRate() int64`

GetAverageTxMsgRate returns the AverageTxMsgRate field if non-nil, zero value otherwise.

### GetAverageTxMsgRateOk

`func (o *MsgVpnQueue) GetAverageTxMsgRateOk() (*int64, bool)`

GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxMsgRate

`func (o *MsgVpnQueue) SetAverageTxMsgRate(v int64)`

SetAverageTxMsgRate sets AverageTxMsgRate field to given value.

### HasAverageTxMsgRate

`func (o *MsgVpnQueue) HasAverageTxMsgRate() bool`

HasAverageTxMsgRate returns a boolean if a field has been set.

### GetBindRequestCount

`func (o *MsgVpnQueue) GetBindRequestCount() int64`

GetBindRequestCount returns the BindRequestCount field if non-nil, zero value otherwise.

### GetBindRequestCountOk

`func (o *MsgVpnQueue) GetBindRequestCountOk() (*int64, bool)`

GetBindRequestCountOk returns a tuple with the BindRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindRequestCount

`func (o *MsgVpnQueue) SetBindRequestCount(v int64)`

SetBindRequestCount sets BindRequestCount field to given value.

### HasBindRequestCount

`func (o *MsgVpnQueue) HasBindRequestCount() bool`

HasBindRequestCount returns a boolean if a field has been set.

### GetBindSuccessCount

`func (o *MsgVpnQueue) GetBindSuccessCount() int64`

GetBindSuccessCount returns the BindSuccessCount field if non-nil, zero value otherwise.

### GetBindSuccessCountOk

`func (o *MsgVpnQueue) GetBindSuccessCountOk() (*int64, bool)`

GetBindSuccessCountOk returns a tuple with the BindSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindSuccessCount

`func (o *MsgVpnQueue) SetBindSuccessCount(v int64)`

SetBindSuccessCount sets BindSuccessCount field to given value.

### HasBindSuccessCount

`func (o *MsgVpnQueue) HasBindSuccessCount() bool`

HasBindSuccessCount returns a boolean if a field has been set.

### GetBindTimeForwardingMode

`func (o *MsgVpnQueue) GetBindTimeForwardingMode() string`

GetBindTimeForwardingMode returns the BindTimeForwardingMode field if non-nil, zero value otherwise.

### GetBindTimeForwardingModeOk

`func (o *MsgVpnQueue) GetBindTimeForwardingModeOk() (*string, bool)`

GetBindTimeForwardingModeOk returns a tuple with the BindTimeForwardingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindTimeForwardingMode

`func (o *MsgVpnQueue) SetBindTimeForwardingMode(v string)`

SetBindTimeForwardingMode sets BindTimeForwardingMode field to given value.

### HasBindTimeForwardingMode

`func (o *MsgVpnQueue) HasBindTimeForwardingMode() bool`

HasBindTimeForwardingMode returns a boolean if a field has been set.

### GetClientProfileDeniedDiscardedMsgCount

`func (o *MsgVpnQueue) GetClientProfileDeniedDiscardedMsgCount() int64`

GetClientProfileDeniedDiscardedMsgCount returns the ClientProfileDeniedDiscardedMsgCount field if non-nil, zero value otherwise.

### GetClientProfileDeniedDiscardedMsgCountOk

`func (o *MsgVpnQueue) GetClientProfileDeniedDiscardedMsgCountOk() (*int64, bool)`

GetClientProfileDeniedDiscardedMsgCountOk returns a tuple with the ClientProfileDeniedDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileDeniedDiscardedMsgCount

`func (o *MsgVpnQueue) SetClientProfileDeniedDiscardedMsgCount(v int64)`

SetClientProfileDeniedDiscardedMsgCount sets ClientProfileDeniedDiscardedMsgCount field to given value.

### HasClientProfileDeniedDiscardedMsgCount

`func (o *MsgVpnQueue) HasClientProfileDeniedDiscardedMsgCount() bool`

HasClientProfileDeniedDiscardedMsgCount returns a boolean if a field has been set.

### GetConsumerAckPropagationEnabled

`func (o *MsgVpnQueue) GetConsumerAckPropagationEnabled() bool`

GetConsumerAckPropagationEnabled returns the ConsumerAckPropagationEnabled field if non-nil, zero value otherwise.

### GetConsumerAckPropagationEnabledOk

`func (o *MsgVpnQueue) GetConsumerAckPropagationEnabledOk() (*bool, bool)`

GetConsumerAckPropagationEnabledOk returns a tuple with the ConsumerAckPropagationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerAckPropagationEnabled

`func (o *MsgVpnQueue) SetConsumerAckPropagationEnabled(v bool)`

SetConsumerAckPropagationEnabled sets ConsumerAckPropagationEnabled field to given value.

### HasConsumerAckPropagationEnabled

`func (o *MsgVpnQueue) HasConsumerAckPropagationEnabled() bool`

HasConsumerAckPropagationEnabled returns a boolean if a field has been set.

### GetCreatedByManagement

`func (o *MsgVpnQueue) GetCreatedByManagement() bool`

GetCreatedByManagement returns the CreatedByManagement field if non-nil, zero value otherwise.

### GetCreatedByManagementOk

`func (o *MsgVpnQueue) GetCreatedByManagementOk() (*bool, bool)`

GetCreatedByManagementOk returns a tuple with the CreatedByManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByManagement

`func (o *MsgVpnQueue) SetCreatedByManagement(v bool)`

SetCreatedByManagement sets CreatedByManagement field to given value.

### HasCreatedByManagement

`func (o *MsgVpnQueue) HasCreatedByManagement() bool`

HasCreatedByManagement returns a boolean if a field has been set.

### GetDeadMsgQueue

`func (o *MsgVpnQueue) GetDeadMsgQueue() string`

GetDeadMsgQueue returns the DeadMsgQueue field if non-nil, zero value otherwise.

### GetDeadMsgQueueOk

`func (o *MsgVpnQueue) GetDeadMsgQueueOk() (*string, bool)`

GetDeadMsgQueueOk returns a tuple with the DeadMsgQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadMsgQueue

`func (o *MsgVpnQueue) SetDeadMsgQueue(v string)`

SetDeadMsgQueue sets DeadMsgQueue field to given value.

### HasDeadMsgQueue

`func (o *MsgVpnQueue) HasDeadMsgQueue() bool`

HasDeadMsgQueue returns a boolean if a field has been set.

### GetDeletedMsgCount

`func (o *MsgVpnQueue) GetDeletedMsgCount() int64`

GetDeletedMsgCount returns the DeletedMsgCount field if non-nil, zero value otherwise.

### GetDeletedMsgCountOk

`func (o *MsgVpnQueue) GetDeletedMsgCountOk() (*int64, bool)`

GetDeletedMsgCountOk returns a tuple with the DeletedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedMsgCount

`func (o *MsgVpnQueue) SetDeletedMsgCount(v int64)`

SetDeletedMsgCount sets DeletedMsgCount field to given value.

### HasDeletedMsgCount

`func (o *MsgVpnQueue) HasDeletedMsgCount() bool`

HasDeletedMsgCount returns a boolean if a field has been set.

### GetDeliveryCountEnabled

`func (o *MsgVpnQueue) GetDeliveryCountEnabled() bool`

GetDeliveryCountEnabled returns the DeliveryCountEnabled field if non-nil, zero value otherwise.

### GetDeliveryCountEnabledOk

`func (o *MsgVpnQueue) GetDeliveryCountEnabledOk() (*bool, bool)`

GetDeliveryCountEnabledOk returns a tuple with the DeliveryCountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCountEnabled

`func (o *MsgVpnQueue) SetDeliveryCountEnabled(v bool)`

SetDeliveryCountEnabled sets DeliveryCountEnabled field to given value.

### HasDeliveryCountEnabled

`func (o *MsgVpnQueue) HasDeliveryCountEnabled() bool`

HasDeliveryCountEnabled returns a boolean if a field has been set.

### GetDestinationGroupErrorDiscardedMsgCount

`func (o *MsgVpnQueue) GetDestinationGroupErrorDiscardedMsgCount() int64`

GetDestinationGroupErrorDiscardedMsgCount returns the DestinationGroupErrorDiscardedMsgCount field if non-nil, zero value otherwise.

### GetDestinationGroupErrorDiscardedMsgCountOk

`func (o *MsgVpnQueue) GetDestinationGroupErrorDiscardedMsgCountOk() (*int64, bool)`

GetDestinationGroupErrorDiscardedMsgCountOk returns a tuple with the DestinationGroupErrorDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroupErrorDiscardedMsgCount

`func (o *MsgVpnQueue) SetDestinationGroupErrorDiscardedMsgCount(v int64)`

SetDestinationGroupErrorDiscardedMsgCount sets DestinationGroupErrorDiscardedMsgCount field to given value.

### HasDestinationGroupErrorDiscardedMsgCount

`func (o *MsgVpnQueue) HasDestinationGroupErrorDiscardedMsgCount() bool`

HasDestinationGroupErrorDiscardedMsgCount returns a boolean if a field has been set.

### GetDisabledBindFailureCount

`func (o *MsgVpnQueue) GetDisabledBindFailureCount() int64`

GetDisabledBindFailureCount returns the DisabledBindFailureCount field if non-nil, zero value otherwise.

### GetDisabledBindFailureCountOk

`func (o *MsgVpnQueue) GetDisabledBindFailureCountOk() (*int64, bool)`

GetDisabledBindFailureCountOk returns a tuple with the DisabledBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledBindFailureCount

`func (o *MsgVpnQueue) SetDisabledBindFailureCount(v int64)`

SetDisabledBindFailureCount sets DisabledBindFailureCount field to given value.

### HasDisabledBindFailureCount

`func (o *MsgVpnQueue) HasDisabledBindFailureCount() bool`

HasDisabledBindFailureCount returns a boolean if a field has been set.

### GetDisabledDiscardedMsgCount

`func (o *MsgVpnQueue) GetDisabledDiscardedMsgCount() int64`

GetDisabledDiscardedMsgCount returns the DisabledDiscardedMsgCount field if non-nil, zero value otherwise.

### GetDisabledDiscardedMsgCountOk

`func (o *MsgVpnQueue) GetDisabledDiscardedMsgCountOk() (*int64, bool)`

GetDisabledDiscardedMsgCountOk returns a tuple with the DisabledDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledDiscardedMsgCount

`func (o *MsgVpnQueue) SetDisabledDiscardedMsgCount(v int64)`

SetDisabledDiscardedMsgCount sets DisabledDiscardedMsgCount field to given value.

### HasDisabledDiscardedMsgCount

`func (o *MsgVpnQueue) HasDisabledDiscardedMsgCount() bool`

HasDisabledDiscardedMsgCount returns a boolean if a field has been set.

### GetDurable

`func (o *MsgVpnQueue) GetDurable() bool`

GetDurable returns the Durable field if non-nil, zero value otherwise.

### GetDurableOk

`func (o *MsgVpnQueue) GetDurableOk() (*bool, bool)`

GetDurableOk returns a tuple with the Durable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurable

`func (o *MsgVpnQueue) SetDurable(v bool)`

SetDurable sets Durable field to given value.

### HasDurable

`func (o *MsgVpnQueue) HasDurable() bool`

HasDurable returns a boolean if a field has been set.

### GetEgressEnabled

`func (o *MsgVpnQueue) GetEgressEnabled() bool`

GetEgressEnabled returns the EgressEnabled field if non-nil, zero value otherwise.

### GetEgressEnabledOk

`func (o *MsgVpnQueue) GetEgressEnabledOk() (*bool, bool)`

GetEgressEnabledOk returns a tuple with the EgressEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressEnabled

`func (o *MsgVpnQueue) SetEgressEnabled(v bool)`

SetEgressEnabled sets EgressEnabled field to given value.

### HasEgressEnabled

`func (o *MsgVpnQueue) HasEgressEnabled() bool`

HasEgressEnabled returns a boolean if a field has been set.

### GetEventBindCountThreshold

`func (o *MsgVpnQueue) GetEventBindCountThreshold() EventThreshold`

GetEventBindCountThreshold returns the EventBindCountThreshold field if non-nil, zero value otherwise.

### GetEventBindCountThresholdOk

`func (o *MsgVpnQueue) GetEventBindCountThresholdOk() (*EventThreshold, bool)`

GetEventBindCountThresholdOk returns a tuple with the EventBindCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventBindCountThreshold

`func (o *MsgVpnQueue) SetEventBindCountThreshold(v EventThreshold)`

SetEventBindCountThreshold sets EventBindCountThreshold field to given value.

### HasEventBindCountThreshold

`func (o *MsgVpnQueue) HasEventBindCountThreshold() bool`

HasEventBindCountThreshold returns a boolean if a field has been set.

### GetEventMsgSpoolUsageThreshold

`func (o *MsgVpnQueue) GetEventMsgSpoolUsageThreshold() EventThreshold`

GetEventMsgSpoolUsageThreshold returns the EventMsgSpoolUsageThreshold field if non-nil, zero value otherwise.

### GetEventMsgSpoolUsageThresholdOk

`func (o *MsgVpnQueue) GetEventMsgSpoolUsageThresholdOk() (*EventThreshold, bool)`

GetEventMsgSpoolUsageThresholdOk returns a tuple with the EventMsgSpoolUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMsgSpoolUsageThreshold

`func (o *MsgVpnQueue) SetEventMsgSpoolUsageThreshold(v EventThreshold)`

SetEventMsgSpoolUsageThreshold sets EventMsgSpoolUsageThreshold field to given value.

### HasEventMsgSpoolUsageThreshold

`func (o *MsgVpnQueue) HasEventMsgSpoolUsageThreshold() bool`

HasEventMsgSpoolUsageThreshold returns a boolean if a field has been set.

### GetEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnQueue) GetEventRejectLowPriorityMsgLimitThreshold() EventThreshold`

GetEventRejectLowPriorityMsgLimitThreshold returns the EventRejectLowPriorityMsgLimitThreshold field if non-nil, zero value otherwise.

### GetEventRejectLowPriorityMsgLimitThresholdOk

`func (o *MsgVpnQueue) GetEventRejectLowPriorityMsgLimitThresholdOk() (*EventThreshold, bool)`

GetEventRejectLowPriorityMsgLimitThresholdOk returns a tuple with the EventRejectLowPriorityMsgLimitThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnQueue) SetEventRejectLowPriorityMsgLimitThreshold(v EventThreshold)`

SetEventRejectLowPriorityMsgLimitThreshold sets EventRejectLowPriorityMsgLimitThreshold field to given value.

### HasEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnQueue) HasEventRejectLowPriorityMsgLimitThreshold() bool`

HasEventRejectLowPriorityMsgLimitThreshold returns a boolean if a field has been set.

### GetHighestAckedMsgId

`func (o *MsgVpnQueue) GetHighestAckedMsgId() int64`

GetHighestAckedMsgId returns the HighestAckedMsgId field if non-nil, zero value otherwise.

### GetHighestAckedMsgIdOk

`func (o *MsgVpnQueue) GetHighestAckedMsgIdOk() (*int64, bool)`

GetHighestAckedMsgIdOk returns a tuple with the HighestAckedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestAckedMsgId

`func (o *MsgVpnQueue) SetHighestAckedMsgId(v int64)`

SetHighestAckedMsgId sets HighestAckedMsgId field to given value.

### HasHighestAckedMsgId

`func (o *MsgVpnQueue) HasHighestAckedMsgId() bool`

HasHighestAckedMsgId returns a boolean if a field has been set.

### GetHighestMsgId

`func (o *MsgVpnQueue) GetHighestMsgId() int64`

GetHighestMsgId returns the HighestMsgId field if non-nil, zero value otherwise.

### GetHighestMsgIdOk

`func (o *MsgVpnQueue) GetHighestMsgIdOk() (*int64, bool)`

GetHighestMsgIdOk returns a tuple with the HighestMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestMsgId

`func (o *MsgVpnQueue) SetHighestMsgId(v int64)`

SetHighestMsgId sets HighestMsgId field to given value.

### HasHighestMsgId

`func (o *MsgVpnQueue) HasHighestMsgId() bool`

HasHighestMsgId returns a boolean if a field has been set.

### GetInProgressAckMsgCount

`func (o *MsgVpnQueue) GetInProgressAckMsgCount() int64`

GetInProgressAckMsgCount returns the InProgressAckMsgCount field if non-nil, zero value otherwise.

### GetInProgressAckMsgCountOk

`func (o *MsgVpnQueue) GetInProgressAckMsgCountOk() (*int64, bool)`

GetInProgressAckMsgCountOk returns a tuple with the InProgressAckMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgressAckMsgCount

`func (o *MsgVpnQueue) SetInProgressAckMsgCount(v int64)`

SetInProgressAckMsgCount sets InProgressAckMsgCount field to given value.

### HasInProgressAckMsgCount

`func (o *MsgVpnQueue) HasInProgressAckMsgCount() bool`

HasInProgressAckMsgCount returns a boolean if a field has been set.

### GetIngressEnabled

`func (o *MsgVpnQueue) GetIngressEnabled() bool`

GetIngressEnabled returns the IngressEnabled field if non-nil, zero value otherwise.

### GetIngressEnabledOk

`func (o *MsgVpnQueue) GetIngressEnabledOk() (*bool, bool)`

GetIngressEnabledOk returns a tuple with the IngressEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressEnabled

`func (o *MsgVpnQueue) SetIngressEnabled(v bool)`

SetIngressEnabled sets IngressEnabled field to given value.

### HasIngressEnabled

`func (o *MsgVpnQueue) HasIngressEnabled() bool`

HasIngressEnabled returns a boolean if a field has been set.

### GetInvalidSelectorBindFailureCount

`func (o *MsgVpnQueue) GetInvalidSelectorBindFailureCount() int64`

GetInvalidSelectorBindFailureCount returns the InvalidSelectorBindFailureCount field if non-nil, zero value otherwise.

### GetInvalidSelectorBindFailureCountOk

`func (o *MsgVpnQueue) GetInvalidSelectorBindFailureCountOk() (*int64, bool)`

GetInvalidSelectorBindFailureCountOk returns a tuple with the InvalidSelectorBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidSelectorBindFailureCount

`func (o *MsgVpnQueue) SetInvalidSelectorBindFailureCount(v int64)`

SetInvalidSelectorBindFailureCount sets InvalidSelectorBindFailureCount field to given value.

### HasInvalidSelectorBindFailureCount

`func (o *MsgVpnQueue) HasInvalidSelectorBindFailureCount() bool`

HasInvalidSelectorBindFailureCount returns a boolean if a field has been set.

### GetLastReplayCompleteTime

`func (o *MsgVpnQueue) GetLastReplayCompleteTime() int32`

GetLastReplayCompleteTime returns the LastReplayCompleteTime field if non-nil, zero value otherwise.

### GetLastReplayCompleteTimeOk

`func (o *MsgVpnQueue) GetLastReplayCompleteTimeOk() (*int32, bool)`

GetLastReplayCompleteTimeOk returns a tuple with the LastReplayCompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplayCompleteTime

`func (o *MsgVpnQueue) SetLastReplayCompleteTime(v int32)`

SetLastReplayCompleteTime sets LastReplayCompleteTime field to given value.

### HasLastReplayCompleteTime

`func (o *MsgVpnQueue) HasLastReplayCompleteTime() bool`

HasLastReplayCompleteTime returns a boolean if a field has been set.

### GetLastReplayFailureReason

`func (o *MsgVpnQueue) GetLastReplayFailureReason() string`

GetLastReplayFailureReason returns the LastReplayFailureReason field if non-nil, zero value otherwise.

### GetLastReplayFailureReasonOk

`func (o *MsgVpnQueue) GetLastReplayFailureReasonOk() (*string, bool)`

GetLastReplayFailureReasonOk returns a tuple with the LastReplayFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplayFailureReason

`func (o *MsgVpnQueue) SetLastReplayFailureReason(v string)`

SetLastReplayFailureReason sets LastReplayFailureReason field to given value.

### HasLastReplayFailureReason

`func (o *MsgVpnQueue) HasLastReplayFailureReason() bool`

HasLastReplayFailureReason returns a boolean if a field has been set.

### GetLastReplayFailureTime

`func (o *MsgVpnQueue) GetLastReplayFailureTime() int32`

GetLastReplayFailureTime returns the LastReplayFailureTime field if non-nil, zero value otherwise.

### GetLastReplayFailureTimeOk

`func (o *MsgVpnQueue) GetLastReplayFailureTimeOk() (*int32, bool)`

GetLastReplayFailureTimeOk returns a tuple with the LastReplayFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplayFailureTime

`func (o *MsgVpnQueue) SetLastReplayFailureTime(v int32)`

SetLastReplayFailureTime sets LastReplayFailureTime field to given value.

### HasLastReplayFailureTime

`func (o *MsgVpnQueue) HasLastReplayFailureTime() bool`

HasLastReplayFailureTime returns a boolean if a field has been set.

### GetLastReplayStartTime

`func (o *MsgVpnQueue) GetLastReplayStartTime() int32`

GetLastReplayStartTime returns the LastReplayStartTime field if non-nil, zero value otherwise.

### GetLastReplayStartTimeOk

`func (o *MsgVpnQueue) GetLastReplayStartTimeOk() (*int32, bool)`

GetLastReplayStartTimeOk returns a tuple with the LastReplayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplayStartTime

`func (o *MsgVpnQueue) SetLastReplayStartTime(v int32)`

SetLastReplayStartTime sets LastReplayStartTime field to given value.

### HasLastReplayStartTime

`func (o *MsgVpnQueue) HasLastReplayStartTime() bool`

HasLastReplayStartTime returns a boolean if a field has been set.

### GetLastReplayedMsgTxTime

`func (o *MsgVpnQueue) GetLastReplayedMsgTxTime() int32`

GetLastReplayedMsgTxTime returns the LastReplayedMsgTxTime field if non-nil, zero value otherwise.

### GetLastReplayedMsgTxTimeOk

`func (o *MsgVpnQueue) GetLastReplayedMsgTxTimeOk() (*int32, bool)`

GetLastReplayedMsgTxTimeOk returns a tuple with the LastReplayedMsgTxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplayedMsgTxTime

`func (o *MsgVpnQueue) SetLastReplayedMsgTxTime(v int32)`

SetLastReplayedMsgTxTime sets LastReplayedMsgTxTime field to given value.

### HasLastReplayedMsgTxTime

`func (o *MsgVpnQueue) HasLastReplayedMsgTxTime() bool`

HasLastReplayedMsgTxTime returns a boolean if a field has been set.

### GetLastSpooledMsgId

`func (o *MsgVpnQueue) GetLastSpooledMsgId() int64`

GetLastSpooledMsgId returns the LastSpooledMsgId field if non-nil, zero value otherwise.

### GetLastSpooledMsgIdOk

`func (o *MsgVpnQueue) GetLastSpooledMsgIdOk() (*int64, bool)`

GetLastSpooledMsgIdOk returns a tuple with the LastSpooledMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSpooledMsgId

`func (o *MsgVpnQueue) SetLastSpooledMsgId(v int64)`

SetLastSpooledMsgId sets LastSpooledMsgId field to given value.

### HasLastSpooledMsgId

`func (o *MsgVpnQueue) HasLastSpooledMsgId() bool`

HasLastSpooledMsgId returns a boolean if a field has been set.

### GetLowPriorityMsgCongestionDiscardedMsgCount

`func (o *MsgVpnQueue) GetLowPriorityMsgCongestionDiscardedMsgCount() int64`

GetLowPriorityMsgCongestionDiscardedMsgCount returns the LowPriorityMsgCongestionDiscardedMsgCount field if non-nil, zero value otherwise.

### GetLowPriorityMsgCongestionDiscardedMsgCountOk

`func (o *MsgVpnQueue) GetLowPriorityMsgCongestionDiscardedMsgCountOk() (*int64, bool)`

GetLowPriorityMsgCongestionDiscardedMsgCountOk returns a tuple with the LowPriorityMsgCongestionDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPriorityMsgCongestionDiscardedMsgCount

`func (o *MsgVpnQueue) SetLowPriorityMsgCongestionDiscardedMsgCount(v int64)`

SetLowPriorityMsgCongestionDiscardedMsgCount sets LowPriorityMsgCongestionDiscardedMsgCount field to given value.

### HasLowPriorityMsgCongestionDiscardedMsgCount

`func (o *MsgVpnQueue) HasLowPriorityMsgCongestionDiscardedMsgCount() bool`

HasLowPriorityMsgCongestionDiscardedMsgCount returns a boolean if a field has been set.

### GetLowPriorityMsgCongestionState

`func (o *MsgVpnQueue) GetLowPriorityMsgCongestionState() string`

GetLowPriorityMsgCongestionState returns the LowPriorityMsgCongestionState field if non-nil, zero value otherwise.

### GetLowPriorityMsgCongestionStateOk

`func (o *MsgVpnQueue) GetLowPriorityMsgCongestionStateOk() (*string, bool)`

GetLowPriorityMsgCongestionStateOk returns a tuple with the LowPriorityMsgCongestionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPriorityMsgCongestionState

`func (o *MsgVpnQueue) SetLowPriorityMsgCongestionState(v string)`

SetLowPriorityMsgCongestionState sets LowPriorityMsgCongestionState field to given value.

### HasLowPriorityMsgCongestionState

`func (o *MsgVpnQueue) HasLowPriorityMsgCongestionState() bool`

HasLowPriorityMsgCongestionState returns a boolean if a field has been set.

### GetLowestAckedMsgId

`func (o *MsgVpnQueue) GetLowestAckedMsgId() int64`

GetLowestAckedMsgId returns the LowestAckedMsgId field if non-nil, zero value otherwise.

### GetLowestAckedMsgIdOk

`func (o *MsgVpnQueue) GetLowestAckedMsgIdOk() (*int64, bool)`

GetLowestAckedMsgIdOk returns a tuple with the LowestAckedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestAckedMsgId

`func (o *MsgVpnQueue) SetLowestAckedMsgId(v int64)`

SetLowestAckedMsgId sets LowestAckedMsgId field to given value.

### HasLowestAckedMsgId

`func (o *MsgVpnQueue) HasLowestAckedMsgId() bool`

HasLowestAckedMsgId returns a boolean if a field has been set.

### GetLowestMsgId

`func (o *MsgVpnQueue) GetLowestMsgId() int64`

GetLowestMsgId returns the LowestMsgId field if non-nil, zero value otherwise.

### GetLowestMsgIdOk

`func (o *MsgVpnQueue) GetLowestMsgIdOk() (*int64, bool)`

GetLowestMsgIdOk returns a tuple with the LowestMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestMsgId

`func (o *MsgVpnQueue) SetLowestMsgId(v int64)`

SetLowestMsgId sets LowestMsgId field to given value.

### HasLowestMsgId

`func (o *MsgVpnQueue) HasLowestMsgId() bool`

HasLowestMsgId returns a boolean if a field has been set.

### GetMaxBindCount

`func (o *MsgVpnQueue) GetMaxBindCount() int64`

GetMaxBindCount returns the MaxBindCount field if non-nil, zero value otherwise.

### GetMaxBindCountOk

`func (o *MsgVpnQueue) GetMaxBindCountOk() (*int64, bool)`

GetMaxBindCountOk returns a tuple with the MaxBindCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBindCount

`func (o *MsgVpnQueue) SetMaxBindCount(v int64)`

SetMaxBindCount sets MaxBindCount field to given value.

### HasMaxBindCount

`func (o *MsgVpnQueue) HasMaxBindCount() bool`

HasMaxBindCount returns a boolean if a field has been set.

### GetMaxBindCountExceededBindFailureCount

`func (o *MsgVpnQueue) GetMaxBindCountExceededBindFailureCount() int64`

GetMaxBindCountExceededBindFailureCount returns the MaxBindCountExceededBindFailureCount field if non-nil, zero value otherwise.

### GetMaxBindCountExceededBindFailureCountOk

`func (o *MsgVpnQueue) GetMaxBindCountExceededBindFailureCountOk() (*int64, bool)`

GetMaxBindCountExceededBindFailureCountOk returns a tuple with the MaxBindCountExceededBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBindCountExceededBindFailureCount

`func (o *MsgVpnQueue) SetMaxBindCountExceededBindFailureCount(v int64)`

SetMaxBindCountExceededBindFailureCount sets MaxBindCountExceededBindFailureCount field to given value.

### HasMaxBindCountExceededBindFailureCount

`func (o *MsgVpnQueue) HasMaxBindCountExceededBindFailureCount() bool`

HasMaxBindCountExceededBindFailureCount returns a boolean if a field has been set.

### GetMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnQueue) GetMaxDeliveredUnackedMsgsPerFlow() int64`

GetMaxDeliveredUnackedMsgsPerFlow returns the MaxDeliveredUnackedMsgsPerFlow field if non-nil, zero value otherwise.

### GetMaxDeliveredUnackedMsgsPerFlowOk

`func (o *MsgVpnQueue) GetMaxDeliveredUnackedMsgsPerFlowOk() (*int64, bool)`

GetMaxDeliveredUnackedMsgsPerFlowOk returns a tuple with the MaxDeliveredUnackedMsgsPerFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnQueue) SetMaxDeliveredUnackedMsgsPerFlow(v int64)`

SetMaxDeliveredUnackedMsgsPerFlow sets MaxDeliveredUnackedMsgsPerFlow field to given value.

### HasMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnQueue) HasMaxDeliveredUnackedMsgsPerFlow() bool`

HasMaxDeliveredUnackedMsgsPerFlow returns a boolean if a field has been set.

### GetMaxMsgSize

`func (o *MsgVpnQueue) GetMaxMsgSize() int32`

GetMaxMsgSize returns the MaxMsgSize field if non-nil, zero value otherwise.

### GetMaxMsgSizeOk

`func (o *MsgVpnQueue) GetMaxMsgSizeOk() (*int32, bool)`

GetMaxMsgSizeOk returns a tuple with the MaxMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSize

`func (o *MsgVpnQueue) SetMaxMsgSize(v int32)`

SetMaxMsgSize sets MaxMsgSize field to given value.

### HasMaxMsgSize

`func (o *MsgVpnQueue) HasMaxMsgSize() bool`

HasMaxMsgSize returns a boolean if a field has been set.

### GetMaxMsgSizeExceededDiscardedMsgCount

`func (o *MsgVpnQueue) GetMaxMsgSizeExceededDiscardedMsgCount() int64`

GetMaxMsgSizeExceededDiscardedMsgCount returns the MaxMsgSizeExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxMsgSizeExceededDiscardedMsgCountOk

`func (o *MsgVpnQueue) GetMaxMsgSizeExceededDiscardedMsgCountOk() (*int64, bool)`

GetMaxMsgSizeExceededDiscardedMsgCountOk returns a tuple with the MaxMsgSizeExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSizeExceededDiscardedMsgCount

`func (o *MsgVpnQueue) SetMaxMsgSizeExceededDiscardedMsgCount(v int64)`

SetMaxMsgSizeExceededDiscardedMsgCount sets MaxMsgSizeExceededDiscardedMsgCount field to given value.

### HasMaxMsgSizeExceededDiscardedMsgCount

`func (o *MsgVpnQueue) HasMaxMsgSizeExceededDiscardedMsgCount() bool`

HasMaxMsgSizeExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxMsgSpoolUsage

`func (o *MsgVpnQueue) GetMaxMsgSpoolUsage() int64`

GetMaxMsgSpoolUsage returns the MaxMsgSpoolUsage field if non-nil, zero value otherwise.

### GetMaxMsgSpoolUsageOk

`func (o *MsgVpnQueue) GetMaxMsgSpoolUsageOk() (*int64, bool)`

GetMaxMsgSpoolUsageOk returns a tuple with the MaxMsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSpoolUsage

`func (o *MsgVpnQueue) SetMaxMsgSpoolUsage(v int64)`

SetMaxMsgSpoolUsage sets MaxMsgSpoolUsage field to given value.

### HasMaxMsgSpoolUsage

`func (o *MsgVpnQueue) HasMaxMsgSpoolUsage() bool`

HasMaxMsgSpoolUsage returns a boolean if a field has been set.

### GetMaxMsgSpoolUsageExceededDiscardedMsgCount

`func (o *MsgVpnQueue) GetMaxMsgSpoolUsageExceededDiscardedMsgCount() int64`

GetMaxMsgSpoolUsageExceededDiscardedMsgCount returns the MaxMsgSpoolUsageExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxMsgSpoolUsageExceededDiscardedMsgCountOk

`func (o *MsgVpnQueue) GetMaxMsgSpoolUsageExceededDiscardedMsgCountOk() (*int64, bool)`

GetMaxMsgSpoolUsageExceededDiscardedMsgCountOk returns a tuple with the MaxMsgSpoolUsageExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSpoolUsageExceededDiscardedMsgCount

`func (o *MsgVpnQueue) SetMaxMsgSpoolUsageExceededDiscardedMsgCount(v int64)`

SetMaxMsgSpoolUsageExceededDiscardedMsgCount sets MaxMsgSpoolUsageExceededDiscardedMsgCount field to given value.

### HasMaxMsgSpoolUsageExceededDiscardedMsgCount

`func (o *MsgVpnQueue) HasMaxMsgSpoolUsageExceededDiscardedMsgCount() bool`

HasMaxMsgSpoolUsageExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxRedeliveryCount

`func (o *MsgVpnQueue) GetMaxRedeliveryCount() int64`

GetMaxRedeliveryCount returns the MaxRedeliveryCount field if non-nil, zero value otherwise.

### GetMaxRedeliveryCountOk

`func (o *MsgVpnQueue) GetMaxRedeliveryCountOk() (*int64, bool)`

GetMaxRedeliveryCountOk returns a tuple with the MaxRedeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedeliveryCount

`func (o *MsgVpnQueue) SetMaxRedeliveryCount(v int64)`

SetMaxRedeliveryCount sets MaxRedeliveryCount field to given value.

### HasMaxRedeliveryCount

`func (o *MsgVpnQueue) HasMaxRedeliveryCount() bool`

HasMaxRedeliveryCount returns a boolean if a field has been set.

### GetMaxRedeliveryExceededDiscardedMsgCount

`func (o *MsgVpnQueue) GetMaxRedeliveryExceededDiscardedMsgCount() int64`

GetMaxRedeliveryExceededDiscardedMsgCount returns the MaxRedeliveryExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxRedeliveryExceededDiscardedMsgCountOk

`func (o *MsgVpnQueue) GetMaxRedeliveryExceededDiscardedMsgCountOk() (*int64, bool)`

GetMaxRedeliveryExceededDiscardedMsgCountOk returns a tuple with the MaxRedeliveryExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedeliveryExceededDiscardedMsgCount

`func (o *MsgVpnQueue) SetMaxRedeliveryExceededDiscardedMsgCount(v int64)`

SetMaxRedeliveryExceededDiscardedMsgCount sets MaxRedeliveryExceededDiscardedMsgCount field to given value.

### HasMaxRedeliveryExceededDiscardedMsgCount

`func (o *MsgVpnQueue) HasMaxRedeliveryExceededDiscardedMsgCount() bool`

HasMaxRedeliveryExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxRedeliveryExceededToDmqFailedMsgCount

`func (o *MsgVpnQueue) GetMaxRedeliveryExceededToDmqFailedMsgCount() int64`

GetMaxRedeliveryExceededToDmqFailedMsgCount returns the MaxRedeliveryExceededToDmqFailedMsgCount field if non-nil, zero value otherwise.

### GetMaxRedeliveryExceededToDmqFailedMsgCountOk

`func (o *MsgVpnQueue) GetMaxRedeliveryExceededToDmqFailedMsgCountOk() (*int64, bool)`

GetMaxRedeliveryExceededToDmqFailedMsgCountOk returns a tuple with the MaxRedeliveryExceededToDmqFailedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedeliveryExceededToDmqFailedMsgCount

`func (o *MsgVpnQueue) SetMaxRedeliveryExceededToDmqFailedMsgCount(v int64)`

SetMaxRedeliveryExceededToDmqFailedMsgCount sets MaxRedeliveryExceededToDmqFailedMsgCount field to given value.

### HasMaxRedeliveryExceededToDmqFailedMsgCount

`func (o *MsgVpnQueue) HasMaxRedeliveryExceededToDmqFailedMsgCount() bool`

HasMaxRedeliveryExceededToDmqFailedMsgCount returns a boolean if a field has been set.

### GetMaxRedeliveryExceededToDmqMsgCount

`func (o *MsgVpnQueue) GetMaxRedeliveryExceededToDmqMsgCount() int64`

GetMaxRedeliveryExceededToDmqMsgCount returns the MaxRedeliveryExceededToDmqMsgCount field if non-nil, zero value otherwise.

### GetMaxRedeliveryExceededToDmqMsgCountOk

`func (o *MsgVpnQueue) GetMaxRedeliveryExceededToDmqMsgCountOk() (*int64, bool)`

GetMaxRedeliveryExceededToDmqMsgCountOk returns a tuple with the MaxRedeliveryExceededToDmqMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedeliveryExceededToDmqMsgCount

`func (o *MsgVpnQueue) SetMaxRedeliveryExceededToDmqMsgCount(v int64)`

SetMaxRedeliveryExceededToDmqMsgCount sets MaxRedeliveryExceededToDmqMsgCount field to given value.

### HasMaxRedeliveryExceededToDmqMsgCount

`func (o *MsgVpnQueue) HasMaxRedeliveryExceededToDmqMsgCount() bool`

HasMaxRedeliveryExceededToDmqMsgCount returns a boolean if a field has been set.

### GetMaxTtl

`func (o *MsgVpnQueue) GetMaxTtl() int64`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *MsgVpnQueue) GetMaxTtlOk() (*int64, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *MsgVpnQueue) SetMaxTtl(v int64)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *MsgVpnQueue) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetMaxTtlExceededDiscardedMsgCount

`func (o *MsgVpnQueue) GetMaxTtlExceededDiscardedMsgCount() int64`

GetMaxTtlExceededDiscardedMsgCount returns the MaxTtlExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxTtlExceededDiscardedMsgCountOk

`func (o *MsgVpnQueue) GetMaxTtlExceededDiscardedMsgCountOk() (*int64, bool)`

GetMaxTtlExceededDiscardedMsgCountOk returns a tuple with the MaxTtlExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtlExceededDiscardedMsgCount

`func (o *MsgVpnQueue) SetMaxTtlExceededDiscardedMsgCount(v int64)`

SetMaxTtlExceededDiscardedMsgCount sets MaxTtlExceededDiscardedMsgCount field to given value.

### HasMaxTtlExceededDiscardedMsgCount

`func (o *MsgVpnQueue) HasMaxTtlExceededDiscardedMsgCount() bool`

HasMaxTtlExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxTtlExpiredDiscardedMsgCount

`func (o *MsgVpnQueue) GetMaxTtlExpiredDiscardedMsgCount() int64`

GetMaxTtlExpiredDiscardedMsgCount returns the MaxTtlExpiredDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxTtlExpiredDiscardedMsgCountOk

`func (o *MsgVpnQueue) GetMaxTtlExpiredDiscardedMsgCountOk() (*int64, bool)`

GetMaxTtlExpiredDiscardedMsgCountOk returns a tuple with the MaxTtlExpiredDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtlExpiredDiscardedMsgCount

`func (o *MsgVpnQueue) SetMaxTtlExpiredDiscardedMsgCount(v int64)`

SetMaxTtlExpiredDiscardedMsgCount sets MaxTtlExpiredDiscardedMsgCount field to given value.

### HasMaxTtlExpiredDiscardedMsgCount

`func (o *MsgVpnQueue) HasMaxTtlExpiredDiscardedMsgCount() bool`

HasMaxTtlExpiredDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxTtlExpiredToDmqFailedMsgCount

`func (o *MsgVpnQueue) GetMaxTtlExpiredToDmqFailedMsgCount() int64`

GetMaxTtlExpiredToDmqFailedMsgCount returns the MaxTtlExpiredToDmqFailedMsgCount field if non-nil, zero value otherwise.

### GetMaxTtlExpiredToDmqFailedMsgCountOk

`func (o *MsgVpnQueue) GetMaxTtlExpiredToDmqFailedMsgCountOk() (*int64, bool)`

GetMaxTtlExpiredToDmqFailedMsgCountOk returns a tuple with the MaxTtlExpiredToDmqFailedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtlExpiredToDmqFailedMsgCount

`func (o *MsgVpnQueue) SetMaxTtlExpiredToDmqFailedMsgCount(v int64)`

SetMaxTtlExpiredToDmqFailedMsgCount sets MaxTtlExpiredToDmqFailedMsgCount field to given value.

### HasMaxTtlExpiredToDmqFailedMsgCount

`func (o *MsgVpnQueue) HasMaxTtlExpiredToDmqFailedMsgCount() bool`

HasMaxTtlExpiredToDmqFailedMsgCount returns a boolean if a field has been set.

### GetMaxTtlExpiredToDmqMsgCount

`func (o *MsgVpnQueue) GetMaxTtlExpiredToDmqMsgCount() int64`

GetMaxTtlExpiredToDmqMsgCount returns the MaxTtlExpiredToDmqMsgCount field if non-nil, zero value otherwise.

### GetMaxTtlExpiredToDmqMsgCountOk

`func (o *MsgVpnQueue) GetMaxTtlExpiredToDmqMsgCountOk() (*int64, bool)`

GetMaxTtlExpiredToDmqMsgCountOk returns a tuple with the MaxTtlExpiredToDmqMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtlExpiredToDmqMsgCount

`func (o *MsgVpnQueue) SetMaxTtlExpiredToDmqMsgCount(v int64)`

SetMaxTtlExpiredToDmqMsgCount sets MaxTtlExpiredToDmqMsgCount field to given value.

### HasMaxTtlExpiredToDmqMsgCount

`func (o *MsgVpnQueue) HasMaxTtlExpiredToDmqMsgCount() bool`

HasMaxTtlExpiredToDmqMsgCount returns a boolean if a field has been set.

### GetMsgSpoolPeakUsage

`func (o *MsgVpnQueue) GetMsgSpoolPeakUsage() int64`

GetMsgSpoolPeakUsage returns the MsgSpoolPeakUsage field if non-nil, zero value otherwise.

### GetMsgSpoolPeakUsageOk

`func (o *MsgVpnQueue) GetMsgSpoolPeakUsageOk() (*int64, bool)`

GetMsgSpoolPeakUsageOk returns a tuple with the MsgSpoolPeakUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolPeakUsage

`func (o *MsgVpnQueue) SetMsgSpoolPeakUsage(v int64)`

SetMsgSpoolPeakUsage sets MsgSpoolPeakUsage field to given value.

### HasMsgSpoolPeakUsage

`func (o *MsgVpnQueue) HasMsgSpoolPeakUsage() bool`

HasMsgSpoolPeakUsage returns a boolean if a field has been set.

### GetMsgSpoolUsage

`func (o *MsgVpnQueue) GetMsgSpoolUsage() int64`

GetMsgSpoolUsage returns the MsgSpoolUsage field if non-nil, zero value otherwise.

### GetMsgSpoolUsageOk

`func (o *MsgVpnQueue) GetMsgSpoolUsageOk() (*int64, bool)`

GetMsgSpoolUsageOk returns a tuple with the MsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolUsage

`func (o *MsgVpnQueue) SetMsgSpoolUsage(v int64)`

SetMsgSpoolUsage sets MsgSpoolUsage field to given value.

### HasMsgSpoolUsage

`func (o *MsgVpnQueue) HasMsgSpoolUsage() bool`

HasMsgSpoolUsage returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnQueue) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnQueue) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnQueue) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnQueue) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetNetworkTopic

`func (o *MsgVpnQueue) GetNetworkTopic() string`

GetNetworkTopic returns the NetworkTopic field if non-nil, zero value otherwise.

### GetNetworkTopicOk

`func (o *MsgVpnQueue) GetNetworkTopicOk() (*string, bool)`

GetNetworkTopicOk returns a tuple with the NetworkTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTopic

`func (o *MsgVpnQueue) SetNetworkTopic(v string)`

SetNetworkTopic sets NetworkTopic field to given value.

### HasNetworkTopic

`func (o *MsgVpnQueue) HasNetworkTopic() bool`

HasNetworkTopic returns a boolean if a field has been set.

### GetNoLocalDeliveryDiscardedMsgCount

`func (o *MsgVpnQueue) GetNoLocalDeliveryDiscardedMsgCount() int64`

GetNoLocalDeliveryDiscardedMsgCount returns the NoLocalDeliveryDiscardedMsgCount field if non-nil, zero value otherwise.

### GetNoLocalDeliveryDiscardedMsgCountOk

`func (o *MsgVpnQueue) GetNoLocalDeliveryDiscardedMsgCountOk() (*int64, bool)`

GetNoLocalDeliveryDiscardedMsgCountOk returns a tuple with the NoLocalDeliveryDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLocalDeliveryDiscardedMsgCount

`func (o *MsgVpnQueue) SetNoLocalDeliveryDiscardedMsgCount(v int64)`

SetNoLocalDeliveryDiscardedMsgCount sets NoLocalDeliveryDiscardedMsgCount field to given value.

### HasNoLocalDeliveryDiscardedMsgCount

`func (o *MsgVpnQueue) HasNoLocalDeliveryDiscardedMsgCount() bool`

HasNoLocalDeliveryDiscardedMsgCount returns a boolean if a field has been set.

### GetOtherBindFailureCount

`func (o *MsgVpnQueue) GetOtherBindFailureCount() int64`

GetOtherBindFailureCount returns the OtherBindFailureCount field if non-nil, zero value otherwise.

### GetOtherBindFailureCountOk

`func (o *MsgVpnQueue) GetOtherBindFailureCountOk() (*int64, bool)`

GetOtherBindFailureCountOk returns a tuple with the OtherBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherBindFailureCount

`func (o *MsgVpnQueue) SetOtherBindFailureCount(v int64)`

SetOtherBindFailureCount sets OtherBindFailureCount field to given value.

### HasOtherBindFailureCount

`func (o *MsgVpnQueue) HasOtherBindFailureCount() bool`

HasOtherBindFailureCount returns a boolean if a field has been set.

### GetOwner

`func (o *MsgVpnQueue) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MsgVpnQueue) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MsgVpnQueue) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *MsgVpnQueue) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPermission

`func (o *MsgVpnQueue) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *MsgVpnQueue) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *MsgVpnQueue) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *MsgVpnQueue) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetQueueName

`func (o *MsgVpnQueue) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *MsgVpnQueue) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *MsgVpnQueue) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *MsgVpnQueue) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetRedeliveredMsgCount

`func (o *MsgVpnQueue) GetRedeliveredMsgCount() int64`

GetRedeliveredMsgCount returns the RedeliveredMsgCount field if non-nil, zero value otherwise.

### GetRedeliveredMsgCountOk

`func (o *MsgVpnQueue) GetRedeliveredMsgCountOk() (*int64, bool)`

GetRedeliveredMsgCountOk returns a tuple with the RedeliveredMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveredMsgCount

`func (o *MsgVpnQueue) SetRedeliveredMsgCount(v int64)`

SetRedeliveredMsgCount sets RedeliveredMsgCount field to given value.

### HasRedeliveredMsgCount

`func (o *MsgVpnQueue) HasRedeliveredMsgCount() bool`

HasRedeliveredMsgCount returns a boolean if a field has been set.

### GetRedeliveryEnabled

`func (o *MsgVpnQueue) GetRedeliveryEnabled() bool`

GetRedeliveryEnabled returns the RedeliveryEnabled field if non-nil, zero value otherwise.

### GetRedeliveryEnabledOk

`func (o *MsgVpnQueue) GetRedeliveryEnabledOk() (*bool, bool)`

GetRedeliveryEnabledOk returns a tuple with the RedeliveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveryEnabled

`func (o *MsgVpnQueue) SetRedeliveryEnabled(v bool)`

SetRedeliveryEnabled sets RedeliveryEnabled field to given value.

### HasRedeliveryEnabled

`func (o *MsgVpnQueue) HasRedeliveryEnabled() bool`

HasRedeliveryEnabled returns a boolean if a field has been set.

### GetRejectLowPriorityMsgEnabled

`func (o *MsgVpnQueue) GetRejectLowPriorityMsgEnabled() bool`

GetRejectLowPriorityMsgEnabled returns the RejectLowPriorityMsgEnabled field if non-nil, zero value otherwise.

### GetRejectLowPriorityMsgEnabledOk

`func (o *MsgVpnQueue) GetRejectLowPriorityMsgEnabledOk() (*bool, bool)`

GetRejectLowPriorityMsgEnabledOk returns a tuple with the RejectLowPriorityMsgEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectLowPriorityMsgEnabled

`func (o *MsgVpnQueue) SetRejectLowPriorityMsgEnabled(v bool)`

SetRejectLowPriorityMsgEnabled sets RejectLowPriorityMsgEnabled field to given value.

### HasRejectLowPriorityMsgEnabled

`func (o *MsgVpnQueue) HasRejectLowPriorityMsgEnabled() bool`

HasRejectLowPriorityMsgEnabled returns a boolean if a field has been set.

### GetRejectLowPriorityMsgLimit

`func (o *MsgVpnQueue) GetRejectLowPriorityMsgLimit() int64`

GetRejectLowPriorityMsgLimit returns the RejectLowPriorityMsgLimit field if non-nil, zero value otherwise.

### GetRejectLowPriorityMsgLimitOk

`func (o *MsgVpnQueue) GetRejectLowPriorityMsgLimitOk() (*int64, bool)`

GetRejectLowPriorityMsgLimitOk returns a tuple with the RejectLowPriorityMsgLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectLowPriorityMsgLimit

`func (o *MsgVpnQueue) SetRejectLowPriorityMsgLimit(v int64)`

SetRejectLowPriorityMsgLimit sets RejectLowPriorityMsgLimit field to given value.

### HasRejectLowPriorityMsgLimit

`func (o *MsgVpnQueue) HasRejectLowPriorityMsgLimit() bool`

HasRejectLowPriorityMsgLimit returns a boolean if a field has been set.

### GetRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnQueue) GetRejectMsgToSenderOnDiscardBehavior() string`

GetRejectMsgToSenderOnDiscardBehavior returns the RejectMsgToSenderOnDiscardBehavior field if non-nil, zero value otherwise.

### GetRejectMsgToSenderOnDiscardBehaviorOk

`func (o *MsgVpnQueue) GetRejectMsgToSenderOnDiscardBehaviorOk() (*string, bool)`

GetRejectMsgToSenderOnDiscardBehaviorOk returns a tuple with the RejectMsgToSenderOnDiscardBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnQueue) SetRejectMsgToSenderOnDiscardBehavior(v string)`

SetRejectMsgToSenderOnDiscardBehavior sets RejectMsgToSenderOnDiscardBehavior field to given value.

### HasRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnQueue) HasRejectMsgToSenderOnDiscardBehavior() bool`

HasRejectMsgToSenderOnDiscardBehavior returns a boolean if a field has been set.

### GetReplayFailureCount

`func (o *MsgVpnQueue) GetReplayFailureCount() int64`

GetReplayFailureCount returns the ReplayFailureCount field if non-nil, zero value otherwise.

### GetReplayFailureCountOk

`func (o *MsgVpnQueue) GetReplayFailureCountOk() (*int64, bool)`

GetReplayFailureCountOk returns a tuple with the ReplayFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayFailureCount

`func (o *MsgVpnQueue) SetReplayFailureCount(v int64)`

SetReplayFailureCount sets ReplayFailureCount field to given value.

### HasReplayFailureCount

`func (o *MsgVpnQueue) HasReplayFailureCount() bool`

HasReplayFailureCount returns a boolean if a field has been set.

### GetReplayStartCount

`func (o *MsgVpnQueue) GetReplayStartCount() int64`

GetReplayStartCount returns the ReplayStartCount field if non-nil, zero value otherwise.

### GetReplayStartCountOk

`func (o *MsgVpnQueue) GetReplayStartCountOk() (*int64, bool)`

GetReplayStartCountOk returns a tuple with the ReplayStartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayStartCount

`func (o *MsgVpnQueue) SetReplayStartCount(v int64)`

SetReplayStartCount sets ReplayStartCount field to given value.

### HasReplayStartCount

`func (o *MsgVpnQueue) HasReplayStartCount() bool`

HasReplayStartCount returns a boolean if a field has been set.

### GetReplayState

`func (o *MsgVpnQueue) GetReplayState() string`

GetReplayState returns the ReplayState field if non-nil, zero value otherwise.

### GetReplayStateOk

`func (o *MsgVpnQueue) GetReplayStateOk() (*string, bool)`

GetReplayStateOk returns a tuple with the ReplayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayState

`func (o *MsgVpnQueue) SetReplayState(v string)`

SetReplayState sets ReplayState field to given value.

### HasReplayState

`func (o *MsgVpnQueue) HasReplayState() bool`

HasReplayState returns a boolean if a field has been set.

### GetReplaySuccessCount

`func (o *MsgVpnQueue) GetReplaySuccessCount() int64`

GetReplaySuccessCount returns the ReplaySuccessCount field if non-nil, zero value otherwise.

### GetReplaySuccessCountOk

`func (o *MsgVpnQueue) GetReplaySuccessCountOk() (*int64, bool)`

GetReplaySuccessCountOk returns a tuple with the ReplaySuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaySuccessCount

`func (o *MsgVpnQueue) SetReplaySuccessCount(v int64)`

SetReplaySuccessCount sets ReplaySuccessCount field to given value.

### HasReplaySuccessCount

`func (o *MsgVpnQueue) HasReplaySuccessCount() bool`

HasReplaySuccessCount returns a boolean if a field has been set.

### GetReplayedAckedMsgCount

`func (o *MsgVpnQueue) GetReplayedAckedMsgCount() int64`

GetReplayedAckedMsgCount returns the ReplayedAckedMsgCount field if non-nil, zero value otherwise.

### GetReplayedAckedMsgCountOk

`func (o *MsgVpnQueue) GetReplayedAckedMsgCountOk() (*int64, bool)`

GetReplayedAckedMsgCountOk returns a tuple with the ReplayedAckedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayedAckedMsgCount

`func (o *MsgVpnQueue) SetReplayedAckedMsgCount(v int64)`

SetReplayedAckedMsgCount sets ReplayedAckedMsgCount field to given value.

### HasReplayedAckedMsgCount

`func (o *MsgVpnQueue) HasReplayedAckedMsgCount() bool`

HasReplayedAckedMsgCount returns a boolean if a field has been set.

### GetReplayedTxMsgCount

`func (o *MsgVpnQueue) GetReplayedTxMsgCount() int64`

GetReplayedTxMsgCount returns the ReplayedTxMsgCount field if non-nil, zero value otherwise.

### GetReplayedTxMsgCountOk

`func (o *MsgVpnQueue) GetReplayedTxMsgCountOk() (*int64, bool)`

GetReplayedTxMsgCountOk returns a tuple with the ReplayedTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayedTxMsgCount

`func (o *MsgVpnQueue) SetReplayedTxMsgCount(v int64)`

SetReplayedTxMsgCount sets ReplayedTxMsgCount field to given value.

### HasReplayedTxMsgCount

`func (o *MsgVpnQueue) HasReplayedTxMsgCount() bool`

HasReplayedTxMsgCount returns a boolean if a field has been set.

### GetReplicationActiveAckPropTxMsgCount

`func (o *MsgVpnQueue) GetReplicationActiveAckPropTxMsgCount() int64`

GetReplicationActiveAckPropTxMsgCount returns the ReplicationActiveAckPropTxMsgCount field if non-nil, zero value otherwise.

### GetReplicationActiveAckPropTxMsgCountOk

`func (o *MsgVpnQueue) GetReplicationActiveAckPropTxMsgCountOk() (*int64, bool)`

GetReplicationActiveAckPropTxMsgCountOk returns a tuple with the ReplicationActiveAckPropTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveAckPropTxMsgCount

`func (o *MsgVpnQueue) SetReplicationActiveAckPropTxMsgCount(v int64)`

SetReplicationActiveAckPropTxMsgCount sets ReplicationActiveAckPropTxMsgCount field to given value.

### HasReplicationActiveAckPropTxMsgCount

`func (o *MsgVpnQueue) HasReplicationActiveAckPropTxMsgCount() bool`

HasReplicationActiveAckPropTxMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyAckPropRxMsgCount

`func (o *MsgVpnQueue) GetReplicationStandbyAckPropRxMsgCount() int64`

GetReplicationStandbyAckPropRxMsgCount returns the ReplicationStandbyAckPropRxMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyAckPropRxMsgCountOk

`func (o *MsgVpnQueue) GetReplicationStandbyAckPropRxMsgCountOk() (*int64, bool)`

GetReplicationStandbyAckPropRxMsgCountOk returns a tuple with the ReplicationStandbyAckPropRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyAckPropRxMsgCount

`func (o *MsgVpnQueue) SetReplicationStandbyAckPropRxMsgCount(v int64)`

SetReplicationStandbyAckPropRxMsgCount sets ReplicationStandbyAckPropRxMsgCount field to given value.

### HasReplicationStandbyAckPropRxMsgCount

`func (o *MsgVpnQueue) HasReplicationStandbyAckPropRxMsgCount() bool`

HasReplicationStandbyAckPropRxMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyAckedByAckPropMsgCount

`func (o *MsgVpnQueue) GetReplicationStandbyAckedByAckPropMsgCount() int64`

GetReplicationStandbyAckedByAckPropMsgCount returns the ReplicationStandbyAckedByAckPropMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyAckedByAckPropMsgCountOk

`func (o *MsgVpnQueue) GetReplicationStandbyAckedByAckPropMsgCountOk() (*int64, bool)`

GetReplicationStandbyAckedByAckPropMsgCountOk returns a tuple with the ReplicationStandbyAckedByAckPropMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyAckedByAckPropMsgCount

`func (o *MsgVpnQueue) SetReplicationStandbyAckedByAckPropMsgCount(v int64)`

SetReplicationStandbyAckedByAckPropMsgCount sets ReplicationStandbyAckedByAckPropMsgCount field to given value.

### HasReplicationStandbyAckedByAckPropMsgCount

`func (o *MsgVpnQueue) HasReplicationStandbyAckedByAckPropMsgCount() bool`

HasReplicationStandbyAckedByAckPropMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyRxMsgCount

`func (o *MsgVpnQueue) GetReplicationStandbyRxMsgCount() int64`

GetReplicationStandbyRxMsgCount returns the ReplicationStandbyRxMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyRxMsgCountOk

`func (o *MsgVpnQueue) GetReplicationStandbyRxMsgCountOk() (*int64, bool)`

GetReplicationStandbyRxMsgCountOk returns a tuple with the ReplicationStandbyRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyRxMsgCount

`func (o *MsgVpnQueue) SetReplicationStandbyRxMsgCount(v int64)`

SetReplicationStandbyRxMsgCount sets ReplicationStandbyRxMsgCount field to given value.

### HasReplicationStandbyRxMsgCount

`func (o *MsgVpnQueue) HasReplicationStandbyRxMsgCount() bool`

HasReplicationStandbyRxMsgCount returns a boolean if a field has been set.

### GetRespectMsgPriorityEnabled

`func (o *MsgVpnQueue) GetRespectMsgPriorityEnabled() bool`

GetRespectMsgPriorityEnabled returns the RespectMsgPriorityEnabled field if non-nil, zero value otherwise.

### GetRespectMsgPriorityEnabledOk

`func (o *MsgVpnQueue) GetRespectMsgPriorityEnabledOk() (*bool, bool)`

GetRespectMsgPriorityEnabledOk returns a tuple with the RespectMsgPriorityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespectMsgPriorityEnabled

`func (o *MsgVpnQueue) SetRespectMsgPriorityEnabled(v bool)`

SetRespectMsgPriorityEnabled sets RespectMsgPriorityEnabled field to given value.

### HasRespectMsgPriorityEnabled

`func (o *MsgVpnQueue) HasRespectMsgPriorityEnabled() bool`

HasRespectMsgPriorityEnabled returns a boolean if a field has been set.

### GetRespectTtlEnabled

`func (o *MsgVpnQueue) GetRespectTtlEnabled() bool`

GetRespectTtlEnabled returns the RespectTtlEnabled field if non-nil, zero value otherwise.

### GetRespectTtlEnabledOk

`func (o *MsgVpnQueue) GetRespectTtlEnabledOk() (*bool, bool)`

GetRespectTtlEnabledOk returns a tuple with the RespectTtlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespectTtlEnabled

`func (o *MsgVpnQueue) SetRespectTtlEnabled(v bool)`

SetRespectTtlEnabled sets RespectTtlEnabled field to given value.

### HasRespectTtlEnabled

`func (o *MsgVpnQueue) HasRespectTtlEnabled() bool`

HasRespectTtlEnabled returns a boolean if a field has been set.

### GetRxByteRate

`func (o *MsgVpnQueue) GetRxByteRate() int64`

GetRxByteRate returns the RxByteRate field if non-nil, zero value otherwise.

### GetRxByteRateOk

`func (o *MsgVpnQueue) GetRxByteRateOk() (*int64, bool)`

GetRxByteRateOk returns a tuple with the RxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteRate

`func (o *MsgVpnQueue) SetRxByteRate(v int64)`

SetRxByteRate sets RxByteRate field to given value.

### HasRxByteRate

`func (o *MsgVpnQueue) HasRxByteRate() bool`

HasRxByteRate returns a boolean if a field has been set.

### GetRxMsgRate

`func (o *MsgVpnQueue) GetRxMsgRate() int64`

GetRxMsgRate returns the RxMsgRate field if non-nil, zero value otherwise.

### GetRxMsgRateOk

`func (o *MsgVpnQueue) GetRxMsgRateOk() (*int64, bool)`

GetRxMsgRateOk returns a tuple with the RxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgRate

`func (o *MsgVpnQueue) SetRxMsgRate(v int64)`

SetRxMsgRate sets RxMsgRate field to given value.

### HasRxMsgRate

`func (o *MsgVpnQueue) HasRxMsgRate() bool`

HasRxMsgRate returns a boolean if a field has been set.

### GetSpooledByteCount

`func (o *MsgVpnQueue) GetSpooledByteCount() int64`

GetSpooledByteCount returns the SpooledByteCount field if non-nil, zero value otherwise.

### GetSpooledByteCountOk

`func (o *MsgVpnQueue) GetSpooledByteCountOk() (*int64, bool)`

GetSpooledByteCountOk returns a tuple with the SpooledByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledByteCount

`func (o *MsgVpnQueue) SetSpooledByteCount(v int64)`

SetSpooledByteCount sets SpooledByteCount field to given value.

### HasSpooledByteCount

`func (o *MsgVpnQueue) HasSpooledByteCount() bool`

HasSpooledByteCount returns a boolean if a field has been set.

### GetSpooledMsgCount

`func (o *MsgVpnQueue) GetSpooledMsgCount() int64`

GetSpooledMsgCount returns the SpooledMsgCount field if non-nil, zero value otherwise.

### GetSpooledMsgCountOk

`func (o *MsgVpnQueue) GetSpooledMsgCountOk() (*int64, bool)`

GetSpooledMsgCountOk returns a tuple with the SpooledMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledMsgCount

`func (o *MsgVpnQueue) SetSpooledMsgCount(v int64)`

SetSpooledMsgCount sets SpooledMsgCount field to given value.

### HasSpooledMsgCount

`func (o *MsgVpnQueue) HasSpooledMsgCount() bool`

HasSpooledMsgCount returns a boolean if a field has been set.

### GetTransportRetransmitMsgCount

`func (o *MsgVpnQueue) GetTransportRetransmitMsgCount() int64`

GetTransportRetransmitMsgCount returns the TransportRetransmitMsgCount field if non-nil, zero value otherwise.

### GetTransportRetransmitMsgCountOk

`func (o *MsgVpnQueue) GetTransportRetransmitMsgCountOk() (*int64, bool)`

GetTransportRetransmitMsgCountOk returns a tuple with the TransportRetransmitMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportRetransmitMsgCount

`func (o *MsgVpnQueue) SetTransportRetransmitMsgCount(v int64)`

SetTransportRetransmitMsgCount sets TransportRetransmitMsgCount field to given value.

### HasTransportRetransmitMsgCount

`func (o *MsgVpnQueue) HasTransportRetransmitMsgCount() bool`

HasTransportRetransmitMsgCount returns a boolean if a field has been set.

### GetTxByteRate

`func (o *MsgVpnQueue) GetTxByteRate() int64`

GetTxByteRate returns the TxByteRate field if non-nil, zero value otherwise.

### GetTxByteRateOk

`func (o *MsgVpnQueue) GetTxByteRateOk() (*int64, bool)`

GetTxByteRateOk returns a tuple with the TxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteRate

`func (o *MsgVpnQueue) SetTxByteRate(v int64)`

SetTxByteRate sets TxByteRate field to given value.

### HasTxByteRate

`func (o *MsgVpnQueue) HasTxByteRate() bool`

HasTxByteRate returns a boolean if a field has been set.

### GetTxMsgRate

`func (o *MsgVpnQueue) GetTxMsgRate() int64`

GetTxMsgRate returns the TxMsgRate field if non-nil, zero value otherwise.

### GetTxMsgRateOk

`func (o *MsgVpnQueue) GetTxMsgRateOk() (*int64, bool)`

GetTxMsgRateOk returns a tuple with the TxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgRate

`func (o *MsgVpnQueue) SetTxMsgRate(v int64)`

SetTxMsgRate sets TxMsgRate field to given value.

### HasTxMsgRate

`func (o *MsgVpnQueue) HasTxMsgRate() bool`

HasTxMsgRate returns a boolean if a field has been set.

### GetTxSelector

`func (o *MsgVpnQueue) GetTxSelector() bool`

GetTxSelector returns the TxSelector field if non-nil, zero value otherwise.

### GetTxSelectorOk

`func (o *MsgVpnQueue) GetTxSelectorOk() (*bool, bool)`

GetTxSelectorOk returns a tuple with the TxSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxSelector

`func (o *MsgVpnQueue) SetTxSelector(v bool)`

SetTxSelector sets TxSelector field to given value.

### HasTxSelector

`func (o *MsgVpnQueue) HasTxSelector() bool`

HasTxSelector returns a boolean if a field has been set.

### GetTxUnackedMsgCount

`func (o *MsgVpnQueue) GetTxUnackedMsgCount() int64`

GetTxUnackedMsgCount returns the TxUnackedMsgCount field if non-nil, zero value otherwise.

### GetTxUnackedMsgCountOk

`func (o *MsgVpnQueue) GetTxUnackedMsgCountOk() (*int64, bool)`

GetTxUnackedMsgCountOk returns a tuple with the TxUnackedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxUnackedMsgCount

`func (o *MsgVpnQueue) SetTxUnackedMsgCount(v int64)`

SetTxUnackedMsgCount sets TxUnackedMsgCount field to given value.

### HasTxUnackedMsgCount

`func (o *MsgVpnQueue) HasTxUnackedMsgCount() bool`

HasTxUnackedMsgCount returns a boolean if a field has been set.

### GetVirtualRouter

`func (o *MsgVpnQueue) GetVirtualRouter() string`

GetVirtualRouter returns the VirtualRouter field if non-nil, zero value otherwise.

### GetVirtualRouterOk

`func (o *MsgVpnQueue) GetVirtualRouterOk() (*string, bool)`

GetVirtualRouterOk returns a tuple with the VirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouter

`func (o *MsgVpnQueue) SetVirtualRouter(v string)`

SetVirtualRouter sets VirtualRouter field to given value.

### HasVirtualRouter

`func (o *MsgVpnQueue) HasVirtualRouter() bool`

HasVirtualRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


