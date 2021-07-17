# MsgVpnQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** | The access type for delivering messages to consumer flows bound to the Queue. The default value is &#x60;\&quot;exclusive\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;exclusive\&quot; - Exclusive delivery of messages to the first bound consumer flow. \&quot;non-exclusive\&quot; - Non-exclusive delivery of messages to all bound consumer flows in a round-robin fashion. &lt;/pre&gt;  | [optional] 
**ConsumerAckPropagationEnabled** | Pointer to **bool** | Enable or disable the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN. The default value is &#x60;true&#x60;. | [optional] 
**DeadMsgQueue** | Pointer to **string** | The name of the Dead Message Queue (DMQ) used by the Queue. The default value is &#x60;\&quot;#DEAD_MSG_QUEUE\&quot;&#x60;. Available since 2.2. | [optional] 
**DeliveryCountEnabled** | Pointer to **bool** | Enable or disable the ability for client applications to query the message delivery count of messages received from the Queue. This is a controlled availability feature. Please contact Solace to find out if this feature is supported for your use case. The default value is &#x60;false&#x60;. Available since 2.19. | [optional] 
**EgressEnabled** | Pointer to **bool** | Enable or disable the transmission of messages from the Queue. The default value is &#x60;false&#x60;. | [optional] 
**EventBindCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventMsgSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventRejectLowPriorityMsgLimitThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**IngressEnabled** | Pointer to **bool** | Enable or disable the reception of messages to the Queue. The default value is &#x60;false&#x60;. | [optional] 
**MaxBindCount** | Pointer to **int64** | The maximum number of consumer flows that can bind to the Queue. The default value is &#x60;1000&#x60;. | [optional] 
**MaxDeliveredUnackedMsgsPerFlow** | Pointer to **int64** | The maximum number of messages delivered but not acknowledged per flow for the Queue. The default value is &#x60;10000&#x60;. | [optional] 
**MaxMsgSize** | Pointer to **int32** | The maximum message size allowed in the Queue, in bytes (B). The default value is &#x60;10000000&#x60;. | [optional] 
**MaxMsgSpoolUsage** | Pointer to **int64** | The maximum message spool usage allowed by the Queue, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. The default value is &#x60;5000&#x60;. | [optional] 
**MaxRedeliveryCount** | Pointer to **int64** | The maximum number of times the Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. The default value is &#x60;0&#x60;. | [optional] 
**MaxTtl** | Pointer to **int64** | The maximum time in seconds a message can stay in the Queue when &#x60;respectTtlEnabled&#x60; is &#x60;\&quot;true\&quot;&#x60;. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the &#x60;maxTtl&#x60; configured for the Queue, is exceeded. A value of 0 disables expiry. The default value is &#x60;0&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Owner** | Pointer to **string** | The Client Username that owns the Queue and has permission equivalent to &#x60;\&quot;delete\&quot;&#x60;. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**Permission** | Pointer to **string** | The permission level for all consumers of the Queue, excluding the owner. The default value is &#x60;\&quot;no-access\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;no-access\&quot; - Disallows all access. \&quot;read-only\&quot; - Read-only access to the messages. \&quot;consume\&quot; - Consume (read and remove) messages. \&quot;modify-topic\&quot; - Consume messages or modify the topic/selector. \&quot;delete\&quot; - Consume messages, modify the topic/selector or delete the Client created endpoint altogether. &lt;/pre&gt;  | [optional] 
**QueueName** | Pointer to **string** | The name of the Queue. | [optional] 
**RedeliveryEnabled** | Pointer to **bool** | Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the queue more than once. The default value is &#x60;true&#x60;. Available since 2.18. | [optional] 
**RejectLowPriorityMsgEnabled** | Pointer to **bool** | Enable or disable the checking of low priority messages against the &#x60;rejectLowPriorityMsgLimit&#x60;. This may only be enabled if &#x60;rejectMsgToSenderOnDiscardBehavior&#x60; does not have a value of &#x60;\&quot;never\&quot;&#x60;. The default value is &#x60;false&#x60;. | [optional] 
**RejectLowPriorityMsgLimit** | Pointer to **int64** | The number of messages of any priority in the Queue above which low priority messages are not admitted but higher priority messages are allowed. The default value is &#x60;0&#x60;. | [optional] 
**RejectMsgToSenderOnDiscardBehavior** | Pointer to **string** | Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The default value is &#x60;\&quot;when-queue-enabled\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always return a negative acknowledgment (NACK) to the sending client on message discard. \&quot;when-queue-enabled\&quot; - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Queue is enabled. \&quot;never\&quot; - Never return a negative acknowledgment (NACK) to the sending client on message discard. &lt;/pre&gt;  Available since 2.1. | [optional] 
**RespectMsgPriorityEnabled** | Pointer to **bool** | Enable or disable the respecting of message priority. When enabled, messages contained in the Queue are delivered in priority order, from 9 (highest) to 0 (lowest). MQTT queues do not support enabling message priority. The default value is &#x60;false&#x60;. Available since 2.8. | [optional] 
**RespectTtlEnabled** | Pointer to **bool** | Enable or disable the respecting of the time-to-live (TTL) for messages in the Queue. When enabled, expired messages are discarded or moved to the DMQ. The default value is &#x60;false&#x60;. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


