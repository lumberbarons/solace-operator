# MsgVpnQueueTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to **string** | The access type for delivering messages to consumer flows. The default value is &#x60;\&quot;exclusive\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;exclusive\&quot; - Exclusive delivery of messages to the first bound consumer flow. \&quot;non-exclusive\&quot; - Non-exclusive delivery of messages to all bound consumer flows in a round-robin fashion. &lt;/pre&gt;  | [optional] 
**ConsumerAckPropagationEnabled** | Pointer to **bool** | Enable or disable the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN. The default value is &#x60;true&#x60;. | [optional] 
**DeadMsgQueue** | Pointer to **string** | The name of the Dead Message Queue (DMQ). The default value is &#x60;\&quot;#DEAD_MSG_QUEUE\&quot;&#x60;. | [optional] 
**DurabilityOverride** | Pointer to **string** | Controls the durability of queues created from this template. If non-durable, the created queue will be non-durable, regardless of the specified durability. If none, the created queue will have the requested durability. The default value is &#x60;\&quot;none\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;none\&quot; - The durability of the endpoint will be as requested on create. \&quot;non-durable\&quot; - The durability of the created queue will be non-durable, regardless of what was requested. &lt;/pre&gt;  | [optional] 
**EventBindCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventMsgSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventRejectLowPriorityMsgLimitThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**MaxBindCount** | Pointer to **int64** | The maximum number of consumer flows that can bind. The default value is &#x60;1000&#x60;. | [optional] 
**MaxDeliveredUnackedMsgsPerFlow** | Pointer to **int64** | The maximum number of messages delivered but not acknowledged per flow. The default value is &#x60;10000&#x60;. | [optional] 
**MaxMsgSize** | Pointer to **int32** | The maximum message size allowed, in bytes (B). The default value is &#x60;10000000&#x60;. | [optional] 
**MaxMsgSpoolUsage** | Pointer to **int64** | The maximum message spool usage allowed, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. The default value is &#x60;5000&#x60;. | [optional] 
**MaxRedeliveryCount** | Pointer to **int64** | The maximum number of message redelivery attempts that will occur prior to the message being discarded or moved to the DMQ. A value of 0 means to retry forever. The default value is &#x60;0&#x60;. | [optional] 
**MaxTtl** | Pointer to **int64** | The maximum time in seconds a message can stay in a Queue when &#x60;respectTtlEnabled&#x60; is &#x60;\&quot;true\&quot;&#x60;. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the &#x60;maxTtl&#x60; configured for the Queue, is exceeded. A value of 0 disables expiry. The default value is &#x60;0&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Permission** | Pointer to **string** | The permission level for all consumers, excluding the owner. The default value is &#x60;\&quot;no-access\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;no-access\&quot; - Disallows all access. \&quot;read-only\&quot; - Read-only access to the messages. \&quot;consume\&quot; - Consume (read and remove) messages. \&quot;modify-topic\&quot; - Consume messages or modify the topic/selector. \&quot;delete\&quot; - Consume messages, modify the topic/selector or delete the Client created endpoint altogether. &lt;/pre&gt;  | [optional] 
**QueueNameFilter** | Pointer to **string** | A wildcardable pattern used to determine which Queues use settings from this Template. Two different wildcards are supported: * and &gt;. Similar to topic filters or subscription patterns, a &gt; matches anything (but only when used at the end), and a * matches zero or more characters but never a slash (/). A &gt; is only a wildcard when used at the end, after a /. A * is only allowed at the end, after a slash (/). The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**QueueTemplateName** | Pointer to **string** | The name of the Queue Template. | [optional] 
**RedeliveryEnabled** | Pointer to **bool** | Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the queue more than once. The default value is &#x60;true&#x60;. Available since 2.18. | [optional] 
**RejectLowPriorityMsgEnabled** | Pointer to **bool** | Enable or disable the checking of low priority messages against the &#x60;rejectLowPriorityMsgLimit&#x60;. This may only be enabled if &#x60;rejectMsgToSenderOnDiscardBehavior&#x60; does not have a value of &#x60;\&quot;never\&quot;&#x60;. The default value is &#x60;false&#x60;. | [optional] 
**RejectLowPriorityMsgLimit** | Pointer to **int64** | The number of messages of any priority above which low priority messages are not admitted but higher priority messages are allowed. The default value is &#x60;0&#x60;. | [optional] 
**RejectMsgToSenderOnDiscardBehavior** | Pointer to **string** | Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs prevent the message from being delivered to any destination and Transacted Session commits to fail. The default value is &#x60;\&quot;when-queue-enabled\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always return a negative acknowledgment (NACK) to the sending client on message discard. \&quot;when-queue-enabled\&quot; - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Queue is enabled. \&quot;never\&quot; - Never return a negative acknowledgment (NACK) to the sending client on message discard. &lt;/pre&gt;  | [optional] 
**RespectMsgPriorityEnabled** | Pointer to **bool** | Enable or disable the respecting of message priority. When enabled, messages are delivered in priority order, from 9 (highest) to 0 (lowest). The default value is &#x60;false&#x60;. | [optional] 
**RespectTtlEnabled** | Pointer to **bool** | Enable or disable the respecting of the time-to-live (TTL) for messages. When enabled, expired messages are discarded or moved to the DMQ. The default value is &#x60;false&#x60;. | [optional] 

## Methods

### NewMsgVpnQueueTemplate

`func NewMsgVpnQueueTemplate() *MsgVpnQueueTemplate`

NewMsgVpnQueueTemplate instantiates a new MsgVpnQueueTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueueTemplateWithDefaults

`func NewMsgVpnQueueTemplateWithDefaults() *MsgVpnQueueTemplate`

NewMsgVpnQueueTemplateWithDefaults instantiates a new MsgVpnQueueTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *MsgVpnQueueTemplate) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *MsgVpnQueueTemplate) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *MsgVpnQueueTemplate) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *MsgVpnQueueTemplate) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetConsumerAckPropagationEnabled

`func (o *MsgVpnQueueTemplate) GetConsumerAckPropagationEnabled() bool`

GetConsumerAckPropagationEnabled returns the ConsumerAckPropagationEnabled field if non-nil, zero value otherwise.

### GetConsumerAckPropagationEnabledOk

`func (o *MsgVpnQueueTemplate) GetConsumerAckPropagationEnabledOk() (*bool, bool)`

GetConsumerAckPropagationEnabledOk returns a tuple with the ConsumerAckPropagationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerAckPropagationEnabled

`func (o *MsgVpnQueueTemplate) SetConsumerAckPropagationEnabled(v bool)`

SetConsumerAckPropagationEnabled sets ConsumerAckPropagationEnabled field to given value.

### HasConsumerAckPropagationEnabled

`func (o *MsgVpnQueueTemplate) HasConsumerAckPropagationEnabled() bool`

HasConsumerAckPropagationEnabled returns a boolean if a field has been set.

### GetDeadMsgQueue

`func (o *MsgVpnQueueTemplate) GetDeadMsgQueue() string`

GetDeadMsgQueue returns the DeadMsgQueue field if non-nil, zero value otherwise.

### GetDeadMsgQueueOk

`func (o *MsgVpnQueueTemplate) GetDeadMsgQueueOk() (*string, bool)`

GetDeadMsgQueueOk returns a tuple with the DeadMsgQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadMsgQueue

`func (o *MsgVpnQueueTemplate) SetDeadMsgQueue(v string)`

SetDeadMsgQueue sets DeadMsgQueue field to given value.

### HasDeadMsgQueue

`func (o *MsgVpnQueueTemplate) HasDeadMsgQueue() bool`

HasDeadMsgQueue returns a boolean if a field has been set.

### GetDurabilityOverride

`func (o *MsgVpnQueueTemplate) GetDurabilityOverride() string`

GetDurabilityOverride returns the DurabilityOverride field if non-nil, zero value otherwise.

### GetDurabilityOverrideOk

`func (o *MsgVpnQueueTemplate) GetDurabilityOverrideOk() (*string, bool)`

GetDurabilityOverrideOk returns a tuple with the DurabilityOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurabilityOverride

`func (o *MsgVpnQueueTemplate) SetDurabilityOverride(v string)`

SetDurabilityOverride sets DurabilityOverride field to given value.

### HasDurabilityOverride

`func (o *MsgVpnQueueTemplate) HasDurabilityOverride() bool`

HasDurabilityOverride returns a boolean if a field has been set.

### GetEventBindCountThreshold

`func (o *MsgVpnQueueTemplate) GetEventBindCountThreshold() EventThreshold`

GetEventBindCountThreshold returns the EventBindCountThreshold field if non-nil, zero value otherwise.

### GetEventBindCountThresholdOk

`func (o *MsgVpnQueueTemplate) GetEventBindCountThresholdOk() (*EventThreshold, bool)`

GetEventBindCountThresholdOk returns a tuple with the EventBindCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventBindCountThreshold

`func (o *MsgVpnQueueTemplate) SetEventBindCountThreshold(v EventThreshold)`

SetEventBindCountThreshold sets EventBindCountThreshold field to given value.

### HasEventBindCountThreshold

`func (o *MsgVpnQueueTemplate) HasEventBindCountThreshold() bool`

HasEventBindCountThreshold returns a boolean if a field has been set.

### GetEventMsgSpoolUsageThreshold

`func (o *MsgVpnQueueTemplate) GetEventMsgSpoolUsageThreshold() EventThreshold`

GetEventMsgSpoolUsageThreshold returns the EventMsgSpoolUsageThreshold field if non-nil, zero value otherwise.

### GetEventMsgSpoolUsageThresholdOk

`func (o *MsgVpnQueueTemplate) GetEventMsgSpoolUsageThresholdOk() (*EventThreshold, bool)`

GetEventMsgSpoolUsageThresholdOk returns a tuple with the EventMsgSpoolUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMsgSpoolUsageThreshold

`func (o *MsgVpnQueueTemplate) SetEventMsgSpoolUsageThreshold(v EventThreshold)`

SetEventMsgSpoolUsageThreshold sets EventMsgSpoolUsageThreshold field to given value.

### HasEventMsgSpoolUsageThreshold

`func (o *MsgVpnQueueTemplate) HasEventMsgSpoolUsageThreshold() bool`

HasEventMsgSpoolUsageThreshold returns a boolean if a field has been set.

### GetEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnQueueTemplate) GetEventRejectLowPriorityMsgLimitThreshold() EventThreshold`

GetEventRejectLowPriorityMsgLimitThreshold returns the EventRejectLowPriorityMsgLimitThreshold field if non-nil, zero value otherwise.

### GetEventRejectLowPriorityMsgLimitThresholdOk

`func (o *MsgVpnQueueTemplate) GetEventRejectLowPriorityMsgLimitThresholdOk() (*EventThreshold, bool)`

GetEventRejectLowPriorityMsgLimitThresholdOk returns a tuple with the EventRejectLowPriorityMsgLimitThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnQueueTemplate) SetEventRejectLowPriorityMsgLimitThreshold(v EventThreshold)`

SetEventRejectLowPriorityMsgLimitThreshold sets EventRejectLowPriorityMsgLimitThreshold field to given value.

### HasEventRejectLowPriorityMsgLimitThreshold

`func (o *MsgVpnQueueTemplate) HasEventRejectLowPriorityMsgLimitThreshold() bool`

HasEventRejectLowPriorityMsgLimitThreshold returns a boolean if a field has been set.

### GetMaxBindCount

`func (o *MsgVpnQueueTemplate) GetMaxBindCount() int64`

GetMaxBindCount returns the MaxBindCount field if non-nil, zero value otherwise.

### GetMaxBindCountOk

`func (o *MsgVpnQueueTemplate) GetMaxBindCountOk() (*int64, bool)`

GetMaxBindCountOk returns a tuple with the MaxBindCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBindCount

`func (o *MsgVpnQueueTemplate) SetMaxBindCount(v int64)`

SetMaxBindCount sets MaxBindCount field to given value.

### HasMaxBindCount

`func (o *MsgVpnQueueTemplate) HasMaxBindCount() bool`

HasMaxBindCount returns a boolean if a field has been set.

### GetMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnQueueTemplate) GetMaxDeliveredUnackedMsgsPerFlow() int64`

GetMaxDeliveredUnackedMsgsPerFlow returns the MaxDeliveredUnackedMsgsPerFlow field if non-nil, zero value otherwise.

### GetMaxDeliveredUnackedMsgsPerFlowOk

`func (o *MsgVpnQueueTemplate) GetMaxDeliveredUnackedMsgsPerFlowOk() (*int64, bool)`

GetMaxDeliveredUnackedMsgsPerFlowOk returns a tuple with the MaxDeliveredUnackedMsgsPerFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnQueueTemplate) SetMaxDeliveredUnackedMsgsPerFlow(v int64)`

SetMaxDeliveredUnackedMsgsPerFlow sets MaxDeliveredUnackedMsgsPerFlow field to given value.

### HasMaxDeliveredUnackedMsgsPerFlow

`func (o *MsgVpnQueueTemplate) HasMaxDeliveredUnackedMsgsPerFlow() bool`

HasMaxDeliveredUnackedMsgsPerFlow returns a boolean if a field has been set.

### GetMaxMsgSize

`func (o *MsgVpnQueueTemplate) GetMaxMsgSize() int32`

GetMaxMsgSize returns the MaxMsgSize field if non-nil, zero value otherwise.

### GetMaxMsgSizeOk

`func (o *MsgVpnQueueTemplate) GetMaxMsgSizeOk() (*int32, bool)`

GetMaxMsgSizeOk returns a tuple with the MaxMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSize

`func (o *MsgVpnQueueTemplate) SetMaxMsgSize(v int32)`

SetMaxMsgSize sets MaxMsgSize field to given value.

### HasMaxMsgSize

`func (o *MsgVpnQueueTemplate) HasMaxMsgSize() bool`

HasMaxMsgSize returns a boolean if a field has been set.

### GetMaxMsgSpoolUsage

`func (o *MsgVpnQueueTemplate) GetMaxMsgSpoolUsage() int64`

GetMaxMsgSpoolUsage returns the MaxMsgSpoolUsage field if non-nil, zero value otherwise.

### GetMaxMsgSpoolUsageOk

`func (o *MsgVpnQueueTemplate) GetMaxMsgSpoolUsageOk() (*int64, bool)`

GetMaxMsgSpoolUsageOk returns a tuple with the MaxMsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSpoolUsage

`func (o *MsgVpnQueueTemplate) SetMaxMsgSpoolUsage(v int64)`

SetMaxMsgSpoolUsage sets MaxMsgSpoolUsage field to given value.

### HasMaxMsgSpoolUsage

`func (o *MsgVpnQueueTemplate) HasMaxMsgSpoolUsage() bool`

HasMaxMsgSpoolUsage returns a boolean if a field has been set.

### GetMaxRedeliveryCount

`func (o *MsgVpnQueueTemplate) GetMaxRedeliveryCount() int64`

GetMaxRedeliveryCount returns the MaxRedeliveryCount field if non-nil, zero value otherwise.

### GetMaxRedeliveryCountOk

`func (o *MsgVpnQueueTemplate) GetMaxRedeliveryCountOk() (*int64, bool)`

GetMaxRedeliveryCountOk returns a tuple with the MaxRedeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedeliveryCount

`func (o *MsgVpnQueueTemplate) SetMaxRedeliveryCount(v int64)`

SetMaxRedeliveryCount sets MaxRedeliveryCount field to given value.

### HasMaxRedeliveryCount

`func (o *MsgVpnQueueTemplate) HasMaxRedeliveryCount() bool`

HasMaxRedeliveryCount returns a boolean if a field has been set.

### GetMaxTtl

`func (o *MsgVpnQueueTemplate) GetMaxTtl() int64`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *MsgVpnQueueTemplate) GetMaxTtlOk() (*int64, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *MsgVpnQueueTemplate) SetMaxTtl(v int64)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *MsgVpnQueueTemplate) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnQueueTemplate) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnQueueTemplate) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnQueueTemplate) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnQueueTemplate) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPermission

`func (o *MsgVpnQueueTemplate) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *MsgVpnQueueTemplate) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *MsgVpnQueueTemplate) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *MsgVpnQueueTemplate) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetQueueNameFilter

`func (o *MsgVpnQueueTemplate) GetQueueNameFilter() string`

GetQueueNameFilter returns the QueueNameFilter field if non-nil, zero value otherwise.

### GetQueueNameFilterOk

`func (o *MsgVpnQueueTemplate) GetQueueNameFilterOk() (*string, bool)`

GetQueueNameFilterOk returns a tuple with the QueueNameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueNameFilter

`func (o *MsgVpnQueueTemplate) SetQueueNameFilter(v string)`

SetQueueNameFilter sets QueueNameFilter field to given value.

### HasQueueNameFilter

`func (o *MsgVpnQueueTemplate) HasQueueNameFilter() bool`

HasQueueNameFilter returns a boolean if a field has been set.

### GetQueueTemplateName

`func (o *MsgVpnQueueTemplate) GetQueueTemplateName() string`

GetQueueTemplateName returns the QueueTemplateName field if non-nil, zero value otherwise.

### GetQueueTemplateNameOk

`func (o *MsgVpnQueueTemplate) GetQueueTemplateNameOk() (*string, bool)`

GetQueueTemplateNameOk returns a tuple with the QueueTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTemplateName

`func (o *MsgVpnQueueTemplate) SetQueueTemplateName(v string)`

SetQueueTemplateName sets QueueTemplateName field to given value.

### HasQueueTemplateName

`func (o *MsgVpnQueueTemplate) HasQueueTemplateName() bool`

HasQueueTemplateName returns a boolean if a field has been set.

### GetRedeliveryEnabled

`func (o *MsgVpnQueueTemplate) GetRedeliveryEnabled() bool`

GetRedeliveryEnabled returns the RedeliveryEnabled field if non-nil, zero value otherwise.

### GetRedeliveryEnabledOk

`func (o *MsgVpnQueueTemplate) GetRedeliveryEnabledOk() (*bool, bool)`

GetRedeliveryEnabledOk returns a tuple with the RedeliveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveryEnabled

`func (o *MsgVpnQueueTemplate) SetRedeliveryEnabled(v bool)`

SetRedeliveryEnabled sets RedeliveryEnabled field to given value.

### HasRedeliveryEnabled

`func (o *MsgVpnQueueTemplate) HasRedeliveryEnabled() bool`

HasRedeliveryEnabled returns a boolean if a field has been set.

### GetRejectLowPriorityMsgEnabled

`func (o *MsgVpnQueueTemplate) GetRejectLowPriorityMsgEnabled() bool`

GetRejectLowPriorityMsgEnabled returns the RejectLowPriorityMsgEnabled field if non-nil, zero value otherwise.

### GetRejectLowPriorityMsgEnabledOk

`func (o *MsgVpnQueueTemplate) GetRejectLowPriorityMsgEnabledOk() (*bool, bool)`

GetRejectLowPriorityMsgEnabledOk returns a tuple with the RejectLowPriorityMsgEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectLowPriorityMsgEnabled

`func (o *MsgVpnQueueTemplate) SetRejectLowPriorityMsgEnabled(v bool)`

SetRejectLowPriorityMsgEnabled sets RejectLowPriorityMsgEnabled field to given value.

### HasRejectLowPriorityMsgEnabled

`func (o *MsgVpnQueueTemplate) HasRejectLowPriorityMsgEnabled() bool`

HasRejectLowPriorityMsgEnabled returns a boolean if a field has been set.

### GetRejectLowPriorityMsgLimit

`func (o *MsgVpnQueueTemplate) GetRejectLowPriorityMsgLimit() int64`

GetRejectLowPriorityMsgLimit returns the RejectLowPriorityMsgLimit field if non-nil, zero value otherwise.

### GetRejectLowPriorityMsgLimitOk

`func (o *MsgVpnQueueTemplate) GetRejectLowPriorityMsgLimitOk() (*int64, bool)`

GetRejectLowPriorityMsgLimitOk returns a tuple with the RejectLowPriorityMsgLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectLowPriorityMsgLimit

`func (o *MsgVpnQueueTemplate) SetRejectLowPriorityMsgLimit(v int64)`

SetRejectLowPriorityMsgLimit sets RejectLowPriorityMsgLimit field to given value.

### HasRejectLowPriorityMsgLimit

`func (o *MsgVpnQueueTemplate) HasRejectLowPriorityMsgLimit() bool`

HasRejectLowPriorityMsgLimit returns a boolean if a field has been set.

### GetRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnQueueTemplate) GetRejectMsgToSenderOnDiscardBehavior() string`

GetRejectMsgToSenderOnDiscardBehavior returns the RejectMsgToSenderOnDiscardBehavior field if non-nil, zero value otherwise.

### GetRejectMsgToSenderOnDiscardBehaviorOk

`func (o *MsgVpnQueueTemplate) GetRejectMsgToSenderOnDiscardBehaviorOk() (*string, bool)`

GetRejectMsgToSenderOnDiscardBehaviorOk returns a tuple with the RejectMsgToSenderOnDiscardBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnQueueTemplate) SetRejectMsgToSenderOnDiscardBehavior(v string)`

SetRejectMsgToSenderOnDiscardBehavior sets RejectMsgToSenderOnDiscardBehavior field to given value.

### HasRejectMsgToSenderOnDiscardBehavior

`func (o *MsgVpnQueueTemplate) HasRejectMsgToSenderOnDiscardBehavior() bool`

HasRejectMsgToSenderOnDiscardBehavior returns a boolean if a field has been set.

### GetRespectMsgPriorityEnabled

`func (o *MsgVpnQueueTemplate) GetRespectMsgPriorityEnabled() bool`

GetRespectMsgPriorityEnabled returns the RespectMsgPriorityEnabled field if non-nil, zero value otherwise.

### GetRespectMsgPriorityEnabledOk

`func (o *MsgVpnQueueTemplate) GetRespectMsgPriorityEnabledOk() (*bool, bool)`

GetRespectMsgPriorityEnabledOk returns a tuple with the RespectMsgPriorityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespectMsgPriorityEnabled

`func (o *MsgVpnQueueTemplate) SetRespectMsgPriorityEnabled(v bool)`

SetRespectMsgPriorityEnabled sets RespectMsgPriorityEnabled field to given value.

### HasRespectMsgPriorityEnabled

`func (o *MsgVpnQueueTemplate) HasRespectMsgPriorityEnabled() bool`

HasRespectMsgPriorityEnabled returns a boolean if a field has been set.

### GetRespectTtlEnabled

`func (o *MsgVpnQueueTemplate) GetRespectTtlEnabled() bool`

GetRespectTtlEnabled returns the RespectTtlEnabled field if non-nil, zero value otherwise.

### GetRespectTtlEnabledOk

`func (o *MsgVpnQueueTemplate) GetRespectTtlEnabledOk() (*bool, bool)`

GetRespectTtlEnabledOk returns a tuple with the RespectTtlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespectTtlEnabled

`func (o *MsgVpnQueueTemplate) SetRespectTtlEnabled(v bool)`

SetRespectTtlEnabled sets RespectTtlEnabled field to given value.

### HasRespectTtlEnabled

`func (o *MsgVpnQueueTemplate) HasRespectTtlEnabled() bool`

HasRespectTtlEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


