# MsgVpnMqttSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable the MQTT Session. When disabled, the client is disconnected, new messages matching QoS 0 subscriptions are discarded, and new messages matching QoS 1 subscriptions are stored for future delivery. The default value is &#x60;false&#x60;. | [optional] 
**MqttSessionClientId** | Pointer to **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | [optional] 
**MqttSessionVirtualRouter** | Pointer to **string** | The virtual router of the MQTT Session. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The MQTT Session belongs to the primary virtual router. \&quot;backup\&quot; - The MQTT Session belongs to the backup virtual router. &lt;/pre&gt;  | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Owner** | Pointer to **string** | The owner of the MQTT Session. For externally-created sessions this defaults to the Client Username of the connecting client. For management-created sessions this defaults to empty. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**QueueConsumerAckPropagationEnabled** | Pointer to **bool** | Enable or disable the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN. The default value is &#x60;true&#x60;. Available since 2.14. | [optional] 
**QueueDeadMsgQueue** | Pointer to **string** | The name of the Dead Message Queue (DMQ) used by the MQTT Session Queue. The default value is &#x60;\&quot;#DEAD_MSG_QUEUE\&quot;&#x60;. Available since 2.14. | [optional] 
**QueueEventBindCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**QueueEventMsgSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**QueueEventRejectLowPriorityMsgLimitThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**QueueMaxBindCount** | Pointer to **int64** | The maximum number of consumer flows that can bind to the MQTT Session Queue. The default value is &#x60;1000&#x60;. Available since 2.14. | [optional] 
**QueueMaxDeliveredUnackedMsgsPerFlow** | Pointer to **int64** | The maximum number of messages delivered but not acknowledged per flow for the MQTT Session Queue. The default value is &#x60;10000&#x60;. Available since 2.14. | [optional] 
**QueueMaxMsgSize** | Pointer to **int32** | The maximum message size allowed in the MQTT Session Queue, in bytes (B). The default value is &#x60;10000000&#x60;. Available since 2.14. | [optional] 
**QueueMaxMsgSpoolUsage** | Pointer to **int64** | The maximum message spool usage allowed by the MQTT Session Queue, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. The default value is &#x60;5000&#x60;. Available since 2.14. | [optional] 
**QueueMaxRedeliveryCount** | Pointer to **int64** | The maximum number of times the MQTT Session Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. The default value is &#x60;0&#x60;. Available since 2.14. | [optional] 
**QueueMaxTtl** | Pointer to **int64** | The maximum time in seconds a message can stay in the MQTT Session Queue when &#x60;queueRespectTtlEnabled&#x60; is &#x60;\&quot;true\&quot;&#x60;. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the &#x60;queueMaxTtl&#x60; configured for the MQTT Session Queue, is exceeded. A value of 0 disables expiry. The default value is &#x60;0&#x60;. Available since 2.14. | [optional] 
**QueueRejectLowPriorityMsgEnabled** | Pointer to **bool** | Enable or disable the checking of low priority messages against the &#x60;queueRejectLowPriorityMsgLimit&#x60;. This may only be enabled if &#x60;queueRejectMsgToSenderOnDiscardBehavior&#x60; does not have a value of &#x60;\&quot;never\&quot;&#x60;. The default value is &#x60;false&#x60;. Available since 2.14. | [optional] 
**QueueRejectLowPriorityMsgLimit** | Pointer to **int64** | The number of messages of any priority in the MQTT Session Queue above which low priority messages are not admitted but higher priority messages are allowed. The default value is &#x60;0&#x60;. Available since 2.14. | [optional] 
**QueueRejectMsgToSenderOnDiscardBehavior** | Pointer to **string** | Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The default value is &#x60;\&quot;when-queue-enabled\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always return a negative acknowledgment (NACK) to the sending client on message discard. \&quot;when-queue-enabled\&quot; - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Queue is enabled. \&quot;never\&quot; - Never return a negative acknowledgment (NACK) to the sending client on message discard. &lt;/pre&gt;  Available since 2.14. | [optional] 
**QueueRespectTtlEnabled** | Pointer to **bool** | Enable or disable the respecting of the time-to-live (TTL) for messages in the MQTT Session Queue. When enabled, expired messages are discarded or moved to the DMQ. The default value is &#x60;false&#x60;. Available since 2.14. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


