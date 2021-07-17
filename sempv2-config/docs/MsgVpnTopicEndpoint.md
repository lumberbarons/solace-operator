# MsgVpnTopicEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** | The access type for delivering messages to consumer flows bound to the Topic Endpoint. The default value is &#x60;\&quot;exclusive\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;exclusive\&quot; - Exclusive delivery of messages to the first bound consumer flow. \&quot;non-exclusive\&quot; - Non-exclusive delivery of messages to all bound consumer flows in a round-robin fashion. &lt;/pre&gt;  Available since 2.4. | [optional] 
**ConsumerAckPropagationEnabled** | Pointer to **bool** | Enable or disable the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN. The default value is &#x60;true&#x60;. | [optional] 
**DeadMsgQueue** | Pointer to **string** | The name of the Dead Message Queue (DMQ) used by the Topic Endpoint. The default value is &#x60;\&quot;#DEAD_MSG_QUEUE\&quot;&#x60;. Available since 2.2. | [optional] 
**DeliveryCountEnabled** | Pointer to **bool** | Enable or disable the ability for client applications to query the message delivery count of messages received from the Topic Endpoint. This is a controlled availability feature. Please contact Solace to find out if this feature is supported for your use case. The default value is &#x60;false&#x60;. Available since 2.19. | [optional] 
**EgressEnabled** | Pointer to **bool** | Enable or disable the transmission of messages from the Topic Endpoint. The default value is &#x60;false&#x60;. | [optional] 
**EventBindCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventRejectLowPriorityMsgLimitThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**IngressEnabled** | Pointer to **bool** | Enable or disable the reception of messages to the Topic Endpoint. The default value is &#x60;false&#x60;. | [optional] 
**MaxBindCount** | Pointer to **int64** | The maximum number of consumer flows that can bind to the Topic Endpoint. The default value is &#x60;1&#x60;. Available since 2.4. | [optional] 
**MaxDeliveredUnackedMsgsPerFlow** | Pointer to **int64** | The maximum number of messages delivered but not acknowledged per flow for the Topic Endpoint. The default value is &#x60;10000&#x60;. | [optional] 
**MaxMsgSize** | Pointer to **int32** | The maximum message size allowed in the Topic Endpoint, in bytes (B). The default value is &#x60;10000000&#x60;. | [optional] 
**MaxRedeliveryCount** | Pointer to **int64** | The maximum number of times the Topic Endpoint will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. The default value is &#x60;0&#x60;. | [optional] 
**MaxSpoolUsage** | Pointer to **int64** | The maximum message spool usage allowed by the Topic Endpoint, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. The default value is &#x60;5000&#x60;. | [optional] 
**MaxTtl** | Pointer to **int64** | The maximum time in seconds a message can stay in the Topic Endpoint when &#x60;respectTtlEnabled&#x60; is &#x60;\&quot;true\&quot;&#x60;. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the &#x60;maxTtl&#x60; configured for the Topic Endpoint, is exceeded. A value of 0 disables expiry. The default value is &#x60;0&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Owner** | Pointer to **string** | The Client Username that owns the Topic Endpoint and has permission equivalent to &#x60;\&quot;delete\&quot;&#x60;. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**Permission** | Pointer to **string** | The permission level for all consumers of the Topic Endpoint, excluding the owner. The default value is &#x60;\&quot;no-access\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;no-access\&quot; - Disallows all access. \&quot;read-only\&quot; - Read-only access to the messages. \&quot;consume\&quot; - Consume (read and remove) messages. \&quot;modify-topic\&quot; - Consume messages or modify the topic/selector. \&quot;delete\&quot; - Consume messages, modify the topic/selector or delete the Client created endpoint altogether. &lt;/pre&gt;  | [optional] 
**RedeliveryEnabled** | Pointer to **bool** | Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the topic-endpoint more than once. The default value is &#x60;true&#x60;. Available since 2.18. | [optional] 
**RejectLowPriorityMsgEnabled** | Pointer to **bool** | Enable or disable the checking of low priority messages against the &#x60;rejectLowPriorityMsgLimit&#x60;. This may only be enabled if &#x60;rejectMsgToSenderOnDiscardBehavior&#x60; does not have a value of &#x60;\&quot;never\&quot;&#x60;. The default value is &#x60;false&#x60;. | [optional] 
**RejectLowPriorityMsgLimit** | Pointer to **int64** | The number of messages of any priority in the Topic Endpoint above which low priority messages are not admitted but higher priority messages are allowed. The default value is &#x60;0&#x60;. | [optional] 
**RejectMsgToSenderOnDiscardBehavior** | Pointer to **string** | Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The default value is &#x60;\&quot;never\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always return a negative acknowledgment (NACK) to the sending client on message discard. \&quot;when-topic-endpoint-enabled\&quot; - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Topic Endpoint is enabled. \&quot;never\&quot; - Never return a negative acknowledgment (NACK) to the sending client on message discard. &lt;/pre&gt;  | [optional] 
**RespectMsgPriorityEnabled** | Pointer to **bool** | Enable or disable the respecting of message priority. When enabled, messages contained in the Topic Endpoint are delivered in priority order, from 9 (highest) to 0 (lowest). The default value is &#x60;false&#x60;. Available since 2.8. | [optional] 
**RespectTtlEnabled** | Pointer to **bool** | Enable or disable the respecting of the time-to-live (TTL) for messages in the Topic Endpoint. When enabled, expired messages are discarded or moved to the DMQ. The default value is &#x60;false&#x60;. | [optional] 
**TopicEndpointName** | Pointer to **string** | The name of the Topic Endpoint. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


