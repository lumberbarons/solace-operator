# MsgVpnClientProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowBridgeConnectionsEnabled** | Pointer to **bool** | Enable or disable allowing Bridge clients using the Client Profile to connect. Changing this setting does not affect existing Bridge client connections. The default value is &#x60;false&#x60;. | [optional] 
**AllowCutThroughForwardingEnabled** | Pointer to **bool** | Enable or disable allowing clients using the Client Profile to bind to endpoints with the cut-through forwarding delivery mode. Changing this value does not affect existing client connections. The default value is &#x60;false&#x60;. | [optional] 
**AllowGuaranteedEndpointCreateDurability** | Pointer to **string** | The types of Queues and Topic Endpoints that clients using the client-profile can create. Changing this value does not affect existing client connections. The default value is &#x60;\&quot;all\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;all\&quot; - Client can create any type of endpoint. \&quot;durable\&quot; - Client can create only durable endpoints. \&quot;non-durable\&quot; - Client can create only non-durable endpoints. &lt;/pre&gt;  Available since 2.14. | [optional] 
**AllowGuaranteedEndpointCreateEnabled** | Pointer to **bool** | Enable or disable allowing clients using the Client Profile to create topic endponts or queues. Changing this value does not affect existing client connections. The default value is &#x60;false&#x60;. | [optional] 
**AllowGuaranteedMsgReceiveEnabled** | Pointer to **bool** | Enable or disable allowing clients using the Client Profile to receive guaranteed messages. Changing this setting does not affect existing client connections. The default value is &#x60;false&#x60;. | [optional] 
**AllowGuaranteedMsgSendEnabled** | Pointer to **bool** | Enable or disable allowing clients using the Client Profile to send guaranteed messages. Changing this setting does not affect existing client connections. The default value is &#x60;false&#x60;. | [optional] 
**AllowSharedSubscriptionsEnabled** | Pointer to **bool** | Enable or disable allowing shared subscriptions. Changing this setting does not affect existing subscriptions. The default value is &#x60;false&#x60;. Available since 2.11. | [optional] 
**AllowTransactedSessionsEnabled** | Pointer to **bool** | Enable or disable allowing clients using the Client Profile to establish transacted sessions. Changing this setting does not affect existing client connections. The default value is &#x60;false&#x60;. | [optional] 
**ApiQueueManagementCopyFromOnCreateName** | Pointer to **string** | The name of a queue to copy settings from when a new queue is created by a client using the Client Profile. The referenced queue must exist in the Message VPN. The default value is &#x60;\&quot;\&quot;&#x60;. Deprecated since 2.14. This attribute has been replaced with &#x60;apiQueueManagementCopyFromOnCreateTemplateName&#x60;. | [optional] 
**ApiQueueManagementCopyFromOnCreateTemplateName** | Pointer to **string** | The name of a queue template to copy settings from when a new queue is created by a client using the Client Profile. If the referenced queue template does not exist, queue creation will fail when it tries to resolve this template. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.14. | [optional] 
**ApiTopicEndpointManagementCopyFromOnCreateName** | Pointer to **string** | The name of a topic endpoint to copy settings from when a new topic endpoint is created by a client using the Client Profile. The referenced topic endpoint must exist in the Message VPN. The default value is &#x60;\&quot;\&quot;&#x60;. Deprecated since 2.14. This attribute has been replaced with &#x60;apiTopicEndpointManagementCopyFromOnCreateTemplateName&#x60;. | [optional] 
**ApiTopicEndpointManagementCopyFromOnCreateTemplateName** | Pointer to **string** | The name of a topic endpoint template to copy settings from when a new topic endpoint is created by a client using the Client Profile. If the referenced topic endpoint template does not exist, topic endpoint creation will fail when it tries to resolve this template. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.14. | [optional] 
**ClientProfileName** | Pointer to **string** | The name of the Client Profile. | [optional] 
**CompressionEnabled** | Pointer to **bool** | Enable or disable allowing clients using the Client Profile to use compression. The default value is &#x60;true&#x60;. Available since 2.10. | [optional] 
**ElidingDelay** | Pointer to **int64** | The amount of time to delay the delivery of messages to clients using the Client Profile after the initial message has been delivered (the eliding delay interval), in milliseconds. A value of 0 means there is no delay in delivering messages to clients. The default value is &#x60;0&#x60;. | [optional] 
**ElidingEnabled** | Pointer to **bool** | Enable or disable message eliding for clients using the Client Profile. The default value is &#x60;false&#x60;. | [optional] 
**ElidingMaxTopicCount** | Pointer to **int64** | The maximum number of topics tracked for message eliding per client connection using the Client Profile. The default value is &#x60;256&#x60;. | [optional] 
**EventClientProvisionedEndpointSpoolUsageThreshold** | Pointer to [**EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] 
**EventConnectionCountPerClientUsernameThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventEgressFlowCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventEndpointCountPerClientUsernameThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventIngressFlowCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventServiceSmfConnectionCountPerClientUsernameThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventServiceWebConnectionCountPerClientUsernameThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventSubscriptionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventTransactedSessionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventTransactionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**MaxConnectionCountPerClientUsername** | Pointer to **int64** | The maximum number of client connections per Client Username using the Client Profile. The default is the maximum value supported by the platform. | [optional] 
**MaxEgressFlowCount** | Pointer to **int64** | The maximum number of transmit flows that can be created by one client using the Client Profile. The default value is &#x60;1000&#x60;. | [optional] 
**MaxEndpointCountPerClientUsername** | Pointer to **int64** | The maximum number of queues and topic endpoints that can be created by clients with the same Client Username using the Client Profile. The default value is &#x60;1000&#x60;. | [optional] 
**MaxIngressFlowCount** | Pointer to **int64** | The maximum number of receive flows that can be created by one client using the Client Profile. The default value is &#x60;1000&#x60;. | [optional] 
**MaxMsgsPerTransaction** | Pointer to **int32** | The maximum number of publisher and consumer messages combined that is allowed within a transaction for each client associated with this client-profile. Exceeding this limit will result in a transaction prepare or commit failure. Changing this value during operation will not affect existing sessions. It is only validated at transaction creation time. Large transactions consume more resources and are more likely to require retrieving messages from the ADB or from disk to process the transaction prepare or commit requests. The transaction processing rate may diminish if a large number of messages must be retrieved from the ADB or from disk. Care should be taken to not use excessively large transactions needlessly to avoid exceeding resource limits and to avoid reducing the overall broker performance. The default value is &#x60;256&#x60;. Available since 2.20. | [optional] 
**MaxSubscriptionCount** | Pointer to **int64** | The maximum number of subscriptions per client using the Client Profile. This limit is not enforced when a client adds a subscription to an endpoint, except for MQTT QoS 1 subscriptions. In addition, this limit is not enforced when a subscription is added using a management interface, such as CLI or SEMP. The default varies by platform. | [optional] 
**MaxTransactedSessionCount** | Pointer to **int64** | The maximum number of transacted sessions that can be created by one client using the Client Profile. The default value is &#x60;10&#x60;. | [optional] 
**MaxTransactionCount** | Pointer to **int64** | The maximum number of transactions that can be created by one client using the Client Profile. The default varies by platform. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**QueueControl1MaxDepth** | Pointer to **int32** | The maximum depth of the \&quot;Control 1\&quot; (C-1) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is &#x60;20000&#x60;. | [optional] 
**QueueControl1MinMsgBurst** | Pointer to **int32** | The number of messages that are always allowed entry into the \&quot;Control 1\&quot; (C-1) priority queue, regardless of the &#x60;queueControl1MaxDepth&#x60; value. The default value is &#x60;4&#x60;. | [optional] 
**QueueDirect1MaxDepth** | Pointer to **int32** | The maximum depth of the \&quot;Direct 1\&quot; (D-1) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is &#x60;20000&#x60;. | [optional] 
**QueueDirect1MinMsgBurst** | Pointer to **int32** | The number of messages that are always allowed entry into the \&quot;Direct 1\&quot; (D-1) priority queue, regardless of the &#x60;queueDirect1MaxDepth&#x60; value. The default value is &#x60;4&#x60;. | [optional] 
**QueueDirect2MaxDepth** | Pointer to **int32** | The maximum depth of the \&quot;Direct 2\&quot; (D-2) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is &#x60;20000&#x60;. | [optional] 
**QueueDirect2MinMsgBurst** | Pointer to **int32** | The number of messages that are always allowed entry into the \&quot;Direct 2\&quot; (D-2) priority queue, regardless of the &#x60;queueDirect2MaxDepth&#x60; value. The default value is &#x60;4&#x60;. | [optional] 
**QueueDirect3MaxDepth** | Pointer to **int32** | The maximum depth of the \&quot;Direct 3\&quot; (D-3) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is &#x60;20000&#x60;. | [optional] 
**QueueDirect3MinMsgBurst** | Pointer to **int32** | The number of messages that are always allowed entry into the \&quot;Direct 3\&quot; (D-3) priority queue, regardless of the &#x60;queueDirect3MaxDepth&#x60; value. The default value is &#x60;4&#x60;. | [optional] 
**QueueGuaranteed1MaxDepth** | Pointer to **int32** | The maximum depth of the \&quot;Guaranteed 1\&quot; (G-1) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is &#x60;20000&#x60;. | [optional] 
**QueueGuaranteed1MinMsgBurst** | Pointer to **int32** | The number of messages that are always allowed entry into the \&quot;Guaranteed 1\&quot; (G-3) priority queue, regardless of the &#x60;queueGuaranteed1MaxDepth&#x60; value. The default value is &#x60;255&#x60;. | [optional] 
**RejectMsgToSenderOnNoSubscriptionMatchEnabled** | Pointer to **bool** | Enable or disable the sending of a negative acknowledgement (NACK) to a client using the Client Profile when discarding a guaranteed message due to no matching subscription found. The default value is &#x60;false&#x60;. Available since 2.2. | [optional] 
**ReplicationAllowClientConnectWhenStandbyEnabled** | Pointer to **bool** | Enable or disable allowing clients using the Client Profile to connect to the Message VPN when its replication state is standby. The default value is &#x60;false&#x60;. | [optional] 
**ServiceMinKeepaliveTimeout** | Pointer to **int32** | The minimum client keepalive timeout which will be enforced for client connections. The default value is &#x60;30&#x60;. Available since 2.19. | [optional] 
**ServiceSmfMaxConnectionCountPerClientUsername** | Pointer to **int64** | The maximum number of SMF client connections per Client Username using the Client Profile. The default is the maximum value supported by the platform. | [optional] 
**ServiceSmfMinKeepaliveEnabled** | Pointer to **bool** | Enable or disable the enforcement of a minimum keepalive timeout for SMF clients. The default value is &#x60;false&#x60;. Available since 2.19. | [optional] 
**ServiceWebInactiveTimeout** | Pointer to **int64** | The timeout for inactive Web Transport client sessions using the Client Profile, in seconds. The default value is &#x60;30&#x60;. | [optional] 
**ServiceWebMaxConnectionCountPerClientUsername** | Pointer to **int64** | The maximum number of Web Transport client connections per Client Username using the Client Profile. The default is the maximum value supported by the platform. | [optional] 
**ServiceWebMaxPayload** | Pointer to **int64** | The maximum Web Transport payload size before fragmentation occurs for clients using the Client Profile, in bytes. The size of the header is not included. The default value is &#x60;1000000&#x60;. | [optional] 
**TcpCongestionWindowSize** | Pointer to **int64** | The TCP initial congestion window size for clients using the Client Profile, in multiples of the TCP Maximum Segment Size (MSS). Changing the value from its default of 2 results in non-compliance with RFC 2581. Contact Solace Support before changing this value. The default value is &#x60;2&#x60;. | [optional] 
**TcpKeepaliveCount** | Pointer to **int64** | The number of TCP keepalive retransmissions to a client using the Client Profile before declaring that it is not available. The default value is &#x60;5&#x60;. | [optional] 
**TcpKeepaliveIdleTime** | Pointer to **int64** | The amount of time a client connection using the Client Profile must remain idle before TCP begins sending keepalive probes, in seconds. The default value is &#x60;3&#x60;. | [optional] 
**TcpKeepaliveInterval** | Pointer to **int64** | The amount of time between TCP keepalive retransmissions to a client using the Client Profile when no acknowledgement is received, in seconds. The default value is &#x60;1&#x60;. | [optional] 
**TcpMaxSegmentSize** | Pointer to **int64** | The TCP maximum segment size for clients using the Client Profile, in bytes. Changes are applied to all existing connections. The default value is &#x60;1460&#x60;. | [optional] 
**TcpMaxWindowSize** | Pointer to **int64** | The TCP maximum window size for clients using the Client Profile, in kilobytes. Changes are applied to all existing connections. The default value is &#x60;256&#x60;. | [optional] 
**TlsAllowDowngradeToPlainTextEnabled** | Pointer to **bool** | Enable or disable allowing a client using the Client Profile to downgrade an encrypted connection to plain text. The default value is &#x60;true&#x60;. Available since 2.8. | [optional] 

## Methods

### NewMsgVpnClientProfile

`func NewMsgVpnClientProfile() *MsgVpnClientProfile`

NewMsgVpnClientProfile instantiates a new MsgVpnClientProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientProfileWithDefaults

`func NewMsgVpnClientProfileWithDefaults() *MsgVpnClientProfile`

NewMsgVpnClientProfileWithDefaults instantiates a new MsgVpnClientProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowBridgeConnectionsEnabled

`func (o *MsgVpnClientProfile) GetAllowBridgeConnectionsEnabled() bool`

GetAllowBridgeConnectionsEnabled returns the AllowBridgeConnectionsEnabled field if non-nil, zero value otherwise.

### GetAllowBridgeConnectionsEnabledOk

`func (o *MsgVpnClientProfile) GetAllowBridgeConnectionsEnabledOk() (*bool, bool)`

GetAllowBridgeConnectionsEnabledOk returns a tuple with the AllowBridgeConnectionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBridgeConnectionsEnabled

`func (o *MsgVpnClientProfile) SetAllowBridgeConnectionsEnabled(v bool)`

SetAllowBridgeConnectionsEnabled sets AllowBridgeConnectionsEnabled field to given value.

### HasAllowBridgeConnectionsEnabled

`func (o *MsgVpnClientProfile) HasAllowBridgeConnectionsEnabled() bool`

HasAllowBridgeConnectionsEnabled returns a boolean if a field has been set.

### GetAllowCutThroughForwardingEnabled

`func (o *MsgVpnClientProfile) GetAllowCutThroughForwardingEnabled() bool`

GetAllowCutThroughForwardingEnabled returns the AllowCutThroughForwardingEnabled field if non-nil, zero value otherwise.

### GetAllowCutThroughForwardingEnabledOk

`func (o *MsgVpnClientProfile) GetAllowCutThroughForwardingEnabledOk() (*bool, bool)`

GetAllowCutThroughForwardingEnabledOk returns a tuple with the AllowCutThroughForwardingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCutThroughForwardingEnabled

`func (o *MsgVpnClientProfile) SetAllowCutThroughForwardingEnabled(v bool)`

SetAllowCutThroughForwardingEnabled sets AllowCutThroughForwardingEnabled field to given value.

### HasAllowCutThroughForwardingEnabled

`func (o *MsgVpnClientProfile) HasAllowCutThroughForwardingEnabled() bool`

HasAllowCutThroughForwardingEnabled returns a boolean if a field has been set.

### GetAllowGuaranteedEndpointCreateDurability

`func (o *MsgVpnClientProfile) GetAllowGuaranteedEndpointCreateDurability() string`

GetAllowGuaranteedEndpointCreateDurability returns the AllowGuaranteedEndpointCreateDurability field if non-nil, zero value otherwise.

### GetAllowGuaranteedEndpointCreateDurabilityOk

`func (o *MsgVpnClientProfile) GetAllowGuaranteedEndpointCreateDurabilityOk() (*string, bool)`

GetAllowGuaranteedEndpointCreateDurabilityOk returns a tuple with the AllowGuaranteedEndpointCreateDurability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuaranteedEndpointCreateDurability

`func (o *MsgVpnClientProfile) SetAllowGuaranteedEndpointCreateDurability(v string)`

SetAllowGuaranteedEndpointCreateDurability sets AllowGuaranteedEndpointCreateDurability field to given value.

### HasAllowGuaranteedEndpointCreateDurability

`func (o *MsgVpnClientProfile) HasAllowGuaranteedEndpointCreateDurability() bool`

HasAllowGuaranteedEndpointCreateDurability returns a boolean if a field has been set.

### GetAllowGuaranteedEndpointCreateEnabled

`func (o *MsgVpnClientProfile) GetAllowGuaranteedEndpointCreateEnabled() bool`

GetAllowGuaranteedEndpointCreateEnabled returns the AllowGuaranteedEndpointCreateEnabled field if non-nil, zero value otherwise.

### GetAllowGuaranteedEndpointCreateEnabledOk

`func (o *MsgVpnClientProfile) GetAllowGuaranteedEndpointCreateEnabledOk() (*bool, bool)`

GetAllowGuaranteedEndpointCreateEnabledOk returns a tuple with the AllowGuaranteedEndpointCreateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuaranteedEndpointCreateEnabled

`func (o *MsgVpnClientProfile) SetAllowGuaranteedEndpointCreateEnabled(v bool)`

SetAllowGuaranteedEndpointCreateEnabled sets AllowGuaranteedEndpointCreateEnabled field to given value.

### HasAllowGuaranteedEndpointCreateEnabled

`func (o *MsgVpnClientProfile) HasAllowGuaranteedEndpointCreateEnabled() bool`

HasAllowGuaranteedEndpointCreateEnabled returns a boolean if a field has been set.

### GetAllowGuaranteedMsgReceiveEnabled

`func (o *MsgVpnClientProfile) GetAllowGuaranteedMsgReceiveEnabled() bool`

GetAllowGuaranteedMsgReceiveEnabled returns the AllowGuaranteedMsgReceiveEnabled field if non-nil, zero value otherwise.

### GetAllowGuaranteedMsgReceiveEnabledOk

`func (o *MsgVpnClientProfile) GetAllowGuaranteedMsgReceiveEnabledOk() (*bool, bool)`

GetAllowGuaranteedMsgReceiveEnabledOk returns a tuple with the AllowGuaranteedMsgReceiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuaranteedMsgReceiveEnabled

`func (o *MsgVpnClientProfile) SetAllowGuaranteedMsgReceiveEnabled(v bool)`

SetAllowGuaranteedMsgReceiveEnabled sets AllowGuaranteedMsgReceiveEnabled field to given value.

### HasAllowGuaranteedMsgReceiveEnabled

`func (o *MsgVpnClientProfile) HasAllowGuaranteedMsgReceiveEnabled() bool`

HasAllowGuaranteedMsgReceiveEnabled returns a boolean if a field has been set.

### GetAllowGuaranteedMsgSendEnabled

`func (o *MsgVpnClientProfile) GetAllowGuaranteedMsgSendEnabled() bool`

GetAllowGuaranteedMsgSendEnabled returns the AllowGuaranteedMsgSendEnabled field if non-nil, zero value otherwise.

### GetAllowGuaranteedMsgSendEnabledOk

`func (o *MsgVpnClientProfile) GetAllowGuaranteedMsgSendEnabledOk() (*bool, bool)`

GetAllowGuaranteedMsgSendEnabledOk returns a tuple with the AllowGuaranteedMsgSendEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGuaranteedMsgSendEnabled

`func (o *MsgVpnClientProfile) SetAllowGuaranteedMsgSendEnabled(v bool)`

SetAllowGuaranteedMsgSendEnabled sets AllowGuaranteedMsgSendEnabled field to given value.

### HasAllowGuaranteedMsgSendEnabled

`func (o *MsgVpnClientProfile) HasAllowGuaranteedMsgSendEnabled() bool`

HasAllowGuaranteedMsgSendEnabled returns a boolean if a field has been set.

### GetAllowSharedSubscriptionsEnabled

`func (o *MsgVpnClientProfile) GetAllowSharedSubscriptionsEnabled() bool`

GetAllowSharedSubscriptionsEnabled returns the AllowSharedSubscriptionsEnabled field if non-nil, zero value otherwise.

### GetAllowSharedSubscriptionsEnabledOk

`func (o *MsgVpnClientProfile) GetAllowSharedSubscriptionsEnabledOk() (*bool, bool)`

GetAllowSharedSubscriptionsEnabledOk returns a tuple with the AllowSharedSubscriptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSharedSubscriptionsEnabled

`func (o *MsgVpnClientProfile) SetAllowSharedSubscriptionsEnabled(v bool)`

SetAllowSharedSubscriptionsEnabled sets AllowSharedSubscriptionsEnabled field to given value.

### HasAllowSharedSubscriptionsEnabled

`func (o *MsgVpnClientProfile) HasAllowSharedSubscriptionsEnabled() bool`

HasAllowSharedSubscriptionsEnabled returns a boolean if a field has been set.

### GetAllowTransactedSessionsEnabled

`func (o *MsgVpnClientProfile) GetAllowTransactedSessionsEnabled() bool`

GetAllowTransactedSessionsEnabled returns the AllowTransactedSessionsEnabled field if non-nil, zero value otherwise.

### GetAllowTransactedSessionsEnabledOk

`func (o *MsgVpnClientProfile) GetAllowTransactedSessionsEnabledOk() (*bool, bool)`

GetAllowTransactedSessionsEnabledOk returns a tuple with the AllowTransactedSessionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransactedSessionsEnabled

`func (o *MsgVpnClientProfile) SetAllowTransactedSessionsEnabled(v bool)`

SetAllowTransactedSessionsEnabled sets AllowTransactedSessionsEnabled field to given value.

### HasAllowTransactedSessionsEnabled

`func (o *MsgVpnClientProfile) HasAllowTransactedSessionsEnabled() bool`

HasAllowTransactedSessionsEnabled returns a boolean if a field has been set.

### GetApiQueueManagementCopyFromOnCreateName

`func (o *MsgVpnClientProfile) GetApiQueueManagementCopyFromOnCreateName() string`

GetApiQueueManagementCopyFromOnCreateName returns the ApiQueueManagementCopyFromOnCreateName field if non-nil, zero value otherwise.

### GetApiQueueManagementCopyFromOnCreateNameOk

`func (o *MsgVpnClientProfile) GetApiQueueManagementCopyFromOnCreateNameOk() (*string, bool)`

GetApiQueueManagementCopyFromOnCreateNameOk returns a tuple with the ApiQueueManagementCopyFromOnCreateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiQueueManagementCopyFromOnCreateName

`func (o *MsgVpnClientProfile) SetApiQueueManagementCopyFromOnCreateName(v string)`

SetApiQueueManagementCopyFromOnCreateName sets ApiQueueManagementCopyFromOnCreateName field to given value.

### HasApiQueueManagementCopyFromOnCreateName

`func (o *MsgVpnClientProfile) HasApiQueueManagementCopyFromOnCreateName() bool`

HasApiQueueManagementCopyFromOnCreateName returns a boolean if a field has been set.

### GetApiQueueManagementCopyFromOnCreateTemplateName

`func (o *MsgVpnClientProfile) GetApiQueueManagementCopyFromOnCreateTemplateName() string`

GetApiQueueManagementCopyFromOnCreateTemplateName returns the ApiQueueManagementCopyFromOnCreateTemplateName field if non-nil, zero value otherwise.

### GetApiQueueManagementCopyFromOnCreateTemplateNameOk

`func (o *MsgVpnClientProfile) GetApiQueueManagementCopyFromOnCreateTemplateNameOk() (*string, bool)`

GetApiQueueManagementCopyFromOnCreateTemplateNameOk returns a tuple with the ApiQueueManagementCopyFromOnCreateTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiQueueManagementCopyFromOnCreateTemplateName

`func (o *MsgVpnClientProfile) SetApiQueueManagementCopyFromOnCreateTemplateName(v string)`

SetApiQueueManagementCopyFromOnCreateTemplateName sets ApiQueueManagementCopyFromOnCreateTemplateName field to given value.

### HasApiQueueManagementCopyFromOnCreateTemplateName

`func (o *MsgVpnClientProfile) HasApiQueueManagementCopyFromOnCreateTemplateName() bool`

HasApiQueueManagementCopyFromOnCreateTemplateName returns a boolean if a field has been set.

### GetApiTopicEndpointManagementCopyFromOnCreateName

`func (o *MsgVpnClientProfile) GetApiTopicEndpointManagementCopyFromOnCreateName() string`

GetApiTopicEndpointManagementCopyFromOnCreateName returns the ApiTopicEndpointManagementCopyFromOnCreateName field if non-nil, zero value otherwise.

### GetApiTopicEndpointManagementCopyFromOnCreateNameOk

`func (o *MsgVpnClientProfile) GetApiTopicEndpointManagementCopyFromOnCreateNameOk() (*string, bool)`

GetApiTopicEndpointManagementCopyFromOnCreateNameOk returns a tuple with the ApiTopicEndpointManagementCopyFromOnCreateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTopicEndpointManagementCopyFromOnCreateName

`func (o *MsgVpnClientProfile) SetApiTopicEndpointManagementCopyFromOnCreateName(v string)`

SetApiTopicEndpointManagementCopyFromOnCreateName sets ApiTopicEndpointManagementCopyFromOnCreateName field to given value.

### HasApiTopicEndpointManagementCopyFromOnCreateName

`func (o *MsgVpnClientProfile) HasApiTopicEndpointManagementCopyFromOnCreateName() bool`

HasApiTopicEndpointManagementCopyFromOnCreateName returns a boolean if a field has been set.

### GetApiTopicEndpointManagementCopyFromOnCreateTemplateName

`func (o *MsgVpnClientProfile) GetApiTopicEndpointManagementCopyFromOnCreateTemplateName() string`

GetApiTopicEndpointManagementCopyFromOnCreateTemplateName returns the ApiTopicEndpointManagementCopyFromOnCreateTemplateName field if non-nil, zero value otherwise.

### GetApiTopicEndpointManagementCopyFromOnCreateTemplateNameOk

`func (o *MsgVpnClientProfile) GetApiTopicEndpointManagementCopyFromOnCreateTemplateNameOk() (*string, bool)`

GetApiTopicEndpointManagementCopyFromOnCreateTemplateNameOk returns a tuple with the ApiTopicEndpointManagementCopyFromOnCreateTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTopicEndpointManagementCopyFromOnCreateTemplateName

`func (o *MsgVpnClientProfile) SetApiTopicEndpointManagementCopyFromOnCreateTemplateName(v string)`

SetApiTopicEndpointManagementCopyFromOnCreateTemplateName sets ApiTopicEndpointManagementCopyFromOnCreateTemplateName field to given value.

### HasApiTopicEndpointManagementCopyFromOnCreateTemplateName

`func (o *MsgVpnClientProfile) HasApiTopicEndpointManagementCopyFromOnCreateTemplateName() bool`

HasApiTopicEndpointManagementCopyFromOnCreateTemplateName returns a boolean if a field has been set.

### GetClientProfileName

`func (o *MsgVpnClientProfile) GetClientProfileName() string`

GetClientProfileName returns the ClientProfileName field if non-nil, zero value otherwise.

### GetClientProfileNameOk

`func (o *MsgVpnClientProfile) GetClientProfileNameOk() (*string, bool)`

GetClientProfileNameOk returns a tuple with the ClientProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileName

`func (o *MsgVpnClientProfile) SetClientProfileName(v string)`

SetClientProfileName sets ClientProfileName field to given value.

### HasClientProfileName

`func (o *MsgVpnClientProfile) HasClientProfileName() bool`

HasClientProfileName returns a boolean if a field has been set.

### GetCompressionEnabled

`func (o *MsgVpnClientProfile) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *MsgVpnClientProfile) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *MsgVpnClientProfile) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *MsgVpnClientProfile) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### GetElidingDelay

`func (o *MsgVpnClientProfile) GetElidingDelay() int64`

GetElidingDelay returns the ElidingDelay field if non-nil, zero value otherwise.

### GetElidingDelayOk

`func (o *MsgVpnClientProfile) GetElidingDelayOk() (*int64, bool)`

GetElidingDelayOk returns a tuple with the ElidingDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElidingDelay

`func (o *MsgVpnClientProfile) SetElidingDelay(v int64)`

SetElidingDelay sets ElidingDelay field to given value.

### HasElidingDelay

`func (o *MsgVpnClientProfile) HasElidingDelay() bool`

HasElidingDelay returns a boolean if a field has been set.

### GetElidingEnabled

`func (o *MsgVpnClientProfile) GetElidingEnabled() bool`

GetElidingEnabled returns the ElidingEnabled field if non-nil, zero value otherwise.

### GetElidingEnabledOk

`func (o *MsgVpnClientProfile) GetElidingEnabledOk() (*bool, bool)`

GetElidingEnabledOk returns a tuple with the ElidingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElidingEnabled

`func (o *MsgVpnClientProfile) SetElidingEnabled(v bool)`

SetElidingEnabled sets ElidingEnabled field to given value.

### HasElidingEnabled

`func (o *MsgVpnClientProfile) HasElidingEnabled() bool`

HasElidingEnabled returns a boolean if a field has been set.

### GetElidingMaxTopicCount

`func (o *MsgVpnClientProfile) GetElidingMaxTopicCount() int64`

GetElidingMaxTopicCount returns the ElidingMaxTopicCount field if non-nil, zero value otherwise.

### GetElidingMaxTopicCountOk

`func (o *MsgVpnClientProfile) GetElidingMaxTopicCountOk() (*int64, bool)`

GetElidingMaxTopicCountOk returns a tuple with the ElidingMaxTopicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElidingMaxTopicCount

`func (o *MsgVpnClientProfile) SetElidingMaxTopicCount(v int64)`

SetElidingMaxTopicCount sets ElidingMaxTopicCount field to given value.

### HasElidingMaxTopicCount

`func (o *MsgVpnClientProfile) HasElidingMaxTopicCount() bool`

HasElidingMaxTopicCount returns a boolean if a field has been set.

### GetEventClientProvisionedEndpointSpoolUsageThreshold

`func (o *MsgVpnClientProfile) GetEventClientProvisionedEndpointSpoolUsageThreshold() EventThresholdByPercent`

GetEventClientProvisionedEndpointSpoolUsageThreshold returns the EventClientProvisionedEndpointSpoolUsageThreshold field if non-nil, zero value otherwise.

### GetEventClientProvisionedEndpointSpoolUsageThresholdOk

`func (o *MsgVpnClientProfile) GetEventClientProvisionedEndpointSpoolUsageThresholdOk() (*EventThresholdByPercent, bool)`

GetEventClientProvisionedEndpointSpoolUsageThresholdOk returns a tuple with the EventClientProvisionedEndpointSpoolUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClientProvisionedEndpointSpoolUsageThreshold

`func (o *MsgVpnClientProfile) SetEventClientProvisionedEndpointSpoolUsageThreshold(v EventThresholdByPercent)`

SetEventClientProvisionedEndpointSpoolUsageThreshold sets EventClientProvisionedEndpointSpoolUsageThreshold field to given value.

### HasEventClientProvisionedEndpointSpoolUsageThreshold

`func (o *MsgVpnClientProfile) HasEventClientProvisionedEndpointSpoolUsageThreshold() bool`

HasEventClientProvisionedEndpointSpoolUsageThreshold returns a boolean if a field has been set.

### GetEventConnectionCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) GetEventConnectionCountPerClientUsernameThreshold() EventThreshold`

GetEventConnectionCountPerClientUsernameThreshold returns the EventConnectionCountPerClientUsernameThreshold field if non-nil, zero value otherwise.

### GetEventConnectionCountPerClientUsernameThresholdOk

`func (o *MsgVpnClientProfile) GetEventConnectionCountPerClientUsernameThresholdOk() (*EventThreshold, bool)`

GetEventConnectionCountPerClientUsernameThresholdOk returns a tuple with the EventConnectionCountPerClientUsernameThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventConnectionCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) SetEventConnectionCountPerClientUsernameThreshold(v EventThreshold)`

SetEventConnectionCountPerClientUsernameThreshold sets EventConnectionCountPerClientUsernameThreshold field to given value.

### HasEventConnectionCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) HasEventConnectionCountPerClientUsernameThreshold() bool`

HasEventConnectionCountPerClientUsernameThreshold returns a boolean if a field has been set.

### GetEventEgressFlowCountThreshold

`func (o *MsgVpnClientProfile) GetEventEgressFlowCountThreshold() EventThreshold`

GetEventEgressFlowCountThreshold returns the EventEgressFlowCountThreshold field if non-nil, zero value otherwise.

### GetEventEgressFlowCountThresholdOk

`func (o *MsgVpnClientProfile) GetEventEgressFlowCountThresholdOk() (*EventThreshold, bool)`

GetEventEgressFlowCountThresholdOk returns a tuple with the EventEgressFlowCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventEgressFlowCountThreshold

`func (o *MsgVpnClientProfile) SetEventEgressFlowCountThreshold(v EventThreshold)`

SetEventEgressFlowCountThreshold sets EventEgressFlowCountThreshold field to given value.

### HasEventEgressFlowCountThreshold

`func (o *MsgVpnClientProfile) HasEventEgressFlowCountThreshold() bool`

HasEventEgressFlowCountThreshold returns a boolean if a field has been set.

### GetEventEndpointCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) GetEventEndpointCountPerClientUsernameThreshold() EventThreshold`

GetEventEndpointCountPerClientUsernameThreshold returns the EventEndpointCountPerClientUsernameThreshold field if non-nil, zero value otherwise.

### GetEventEndpointCountPerClientUsernameThresholdOk

`func (o *MsgVpnClientProfile) GetEventEndpointCountPerClientUsernameThresholdOk() (*EventThreshold, bool)`

GetEventEndpointCountPerClientUsernameThresholdOk returns a tuple with the EventEndpointCountPerClientUsernameThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventEndpointCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) SetEventEndpointCountPerClientUsernameThreshold(v EventThreshold)`

SetEventEndpointCountPerClientUsernameThreshold sets EventEndpointCountPerClientUsernameThreshold field to given value.

### HasEventEndpointCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) HasEventEndpointCountPerClientUsernameThreshold() bool`

HasEventEndpointCountPerClientUsernameThreshold returns a boolean if a field has been set.

### GetEventIngressFlowCountThreshold

`func (o *MsgVpnClientProfile) GetEventIngressFlowCountThreshold() EventThreshold`

GetEventIngressFlowCountThreshold returns the EventIngressFlowCountThreshold field if non-nil, zero value otherwise.

### GetEventIngressFlowCountThresholdOk

`func (o *MsgVpnClientProfile) GetEventIngressFlowCountThresholdOk() (*EventThreshold, bool)`

GetEventIngressFlowCountThresholdOk returns a tuple with the EventIngressFlowCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIngressFlowCountThreshold

`func (o *MsgVpnClientProfile) SetEventIngressFlowCountThreshold(v EventThreshold)`

SetEventIngressFlowCountThreshold sets EventIngressFlowCountThreshold field to given value.

### HasEventIngressFlowCountThreshold

`func (o *MsgVpnClientProfile) HasEventIngressFlowCountThreshold() bool`

HasEventIngressFlowCountThreshold returns a boolean if a field has been set.

### GetEventServiceSmfConnectionCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) GetEventServiceSmfConnectionCountPerClientUsernameThreshold() EventThreshold`

GetEventServiceSmfConnectionCountPerClientUsernameThreshold returns the EventServiceSmfConnectionCountPerClientUsernameThreshold field if non-nil, zero value otherwise.

### GetEventServiceSmfConnectionCountPerClientUsernameThresholdOk

`func (o *MsgVpnClientProfile) GetEventServiceSmfConnectionCountPerClientUsernameThresholdOk() (*EventThreshold, bool)`

GetEventServiceSmfConnectionCountPerClientUsernameThresholdOk returns a tuple with the EventServiceSmfConnectionCountPerClientUsernameThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceSmfConnectionCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) SetEventServiceSmfConnectionCountPerClientUsernameThreshold(v EventThreshold)`

SetEventServiceSmfConnectionCountPerClientUsernameThreshold sets EventServiceSmfConnectionCountPerClientUsernameThreshold field to given value.

### HasEventServiceSmfConnectionCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) HasEventServiceSmfConnectionCountPerClientUsernameThreshold() bool`

HasEventServiceSmfConnectionCountPerClientUsernameThreshold returns a boolean if a field has been set.

### GetEventServiceWebConnectionCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) GetEventServiceWebConnectionCountPerClientUsernameThreshold() EventThreshold`

GetEventServiceWebConnectionCountPerClientUsernameThreshold returns the EventServiceWebConnectionCountPerClientUsernameThreshold field if non-nil, zero value otherwise.

### GetEventServiceWebConnectionCountPerClientUsernameThresholdOk

`func (o *MsgVpnClientProfile) GetEventServiceWebConnectionCountPerClientUsernameThresholdOk() (*EventThreshold, bool)`

GetEventServiceWebConnectionCountPerClientUsernameThresholdOk returns a tuple with the EventServiceWebConnectionCountPerClientUsernameThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceWebConnectionCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) SetEventServiceWebConnectionCountPerClientUsernameThreshold(v EventThreshold)`

SetEventServiceWebConnectionCountPerClientUsernameThreshold sets EventServiceWebConnectionCountPerClientUsernameThreshold field to given value.

### HasEventServiceWebConnectionCountPerClientUsernameThreshold

`func (o *MsgVpnClientProfile) HasEventServiceWebConnectionCountPerClientUsernameThreshold() bool`

HasEventServiceWebConnectionCountPerClientUsernameThreshold returns a boolean if a field has been set.

### GetEventSubscriptionCountThreshold

`func (o *MsgVpnClientProfile) GetEventSubscriptionCountThreshold() EventThreshold`

GetEventSubscriptionCountThreshold returns the EventSubscriptionCountThreshold field if non-nil, zero value otherwise.

### GetEventSubscriptionCountThresholdOk

`func (o *MsgVpnClientProfile) GetEventSubscriptionCountThresholdOk() (*EventThreshold, bool)`

GetEventSubscriptionCountThresholdOk returns a tuple with the EventSubscriptionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubscriptionCountThreshold

`func (o *MsgVpnClientProfile) SetEventSubscriptionCountThreshold(v EventThreshold)`

SetEventSubscriptionCountThreshold sets EventSubscriptionCountThreshold field to given value.

### HasEventSubscriptionCountThreshold

`func (o *MsgVpnClientProfile) HasEventSubscriptionCountThreshold() bool`

HasEventSubscriptionCountThreshold returns a boolean if a field has been set.

### GetEventTransactedSessionCountThreshold

`func (o *MsgVpnClientProfile) GetEventTransactedSessionCountThreshold() EventThreshold`

GetEventTransactedSessionCountThreshold returns the EventTransactedSessionCountThreshold field if non-nil, zero value otherwise.

### GetEventTransactedSessionCountThresholdOk

`func (o *MsgVpnClientProfile) GetEventTransactedSessionCountThresholdOk() (*EventThreshold, bool)`

GetEventTransactedSessionCountThresholdOk returns a tuple with the EventTransactedSessionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTransactedSessionCountThreshold

`func (o *MsgVpnClientProfile) SetEventTransactedSessionCountThreshold(v EventThreshold)`

SetEventTransactedSessionCountThreshold sets EventTransactedSessionCountThreshold field to given value.

### HasEventTransactedSessionCountThreshold

`func (o *MsgVpnClientProfile) HasEventTransactedSessionCountThreshold() bool`

HasEventTransactedSessionCountThreshold returns a boolean if a field has been set.

### GetEventTransactionCountThreshold

`func (o *MsgVpnClientProfile) GetEventTransactionCountThreshold() EventThreshold`

GetEventTransactionCountThreshold returns the EventTransactionCountThreshold field if non-nil, zero value otherwise.

### GetEventTransactionCountThresholdOk

`func (o *MsgVpnClientProfile) GetEventTransactionCountThresholdOk() (*EventThreshold, bool)`

GetEventTransactionCountThresholdOk returns a tuple with the EventTransactionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTransactionCountThreshold

`func (o *MsgVpnClientProfile) SetEventTransactionCountThreshold(v EventThreshold)`

SetEventTransactionCountThreshold sets EventTransactionCountThreshold field to given value.

### HasEventTransactionCountThreshold

`func (o *MsgVpnClientProfile) HasEventTransactionCountThreshold() bool`

HasEventTransactionCountThreshold returns a boolean if a field has been set.

### GetMaxConnectionCountPerClientUsername

`func (o *MsgVpnClientProfile) GetMaxConnectionCountPerClientUsername() int64`

GetMaxConnectionCountPerClientUsername returns the MaxConnectionCountPerClientUsername field if non-nil, zero value otherwise.

### GetMaxConnectionCountPerClientUsernameOk

`func (o *MsgVpnClientProfile) GetMaxConnectionCountPerClientUsernameOk() (*int64, bool)`

GetMaxConnectionCountPerClientUsernameOk returns a tuple with the MaxConnectionCountPerClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionCountPerClientUsername

`func (o *MsgVpnClientProfile) SetMaxConnectionCountPerClientUsername(v int64)`

SetMaxConnectionCountPerClientUsername sets MaxConnectionCountPerClientUsername field to given value.

### HasMaxConnectionCountPerClientUsername

`func (o *MsgVpnClientProfile) HasMaxConnectionCountPerClientUsername() bool`

HasMaxConnectionCountPerClientUsername returns a boolean if a field has been set.

### GetMaxEgressFlowCount

`func (o *MsgVpnClientProfile) GetMaxEgressFlowCount() int64`

GetMaxEgressFlowCount returns the MaxEgressFlowCount field if non-nil, zero value otherwise.

### GetMaxEgressFlowCountOk

`func (o *MsgVpnClientProfile) GetMaxEgressFlowCountOk() (*int64, bool)`

GetMaxEgressFlowCountOk returns a tuple with the MaxEgressFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEgressFlowCount

`func (o *MsgVpnClientProfile) SetMaxEgressFlowCount(v int64)`

SetMaxEgressFlowCount sets MaxEgressFlowCount field to given value.

### HasMaxEgressFlowCount

`func (o *MsgVpnClientProfile) HasMaxEgressFlowCount() bool`

HasMaxEgressFlowCount returns a boolean if a field has been set.

### GetMaxEndpointCountPerClientUsername

`func (o *MsgVpnClientProfile) GetMaxEndpointCountPerClientUsername() int64`

GetMaxEndpointCountPerClientUsername returns the MaxEndpointCountPerClientUsername field if non-nil, zero value otherwise.

### GetMaxEndpointCountPerClientUsernameOk

`func (o *MsgVpnClientProfile) GetMaxEndpointCountPerClientUsernameOk() (*int64, bool)`

GetMaxEndpointCountPerClientUsernameOk returns a tuple with the MaxEndpointCountPerClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndpointCountPerClientUsername

`func (o *MsgVpnClientProfile) SetMaxEndpointCountPerClientUsername(v int64)`

SetMaxEndpointCountPerClientUsername sets MaxEndpointCountPerClientUsername field to given value.

### HasMaxEndpointCountPerClientUsername

`func (o *MsgVpnClientProfile) HasMaxEndpointCountPerClientUsername() bool`

HasMaxEndpointCountPerClientUsername returns a boolean if a field has been set.

### GetMaxIngressFlowCount

`func (o *MsgVpnClientProfile) GetMaxIngressFlowCount() int64`

GetMaxIngressFlowCount returns the MaxIngressFlowCount field if non-nil, zero value otherwise.

### GetMaxIngressFlowCountOk

`func (o *MsgVpnClientProfile) GetMaxIngressFlowCountOk() (*int64, bool)`

GetMaxIngressFlowCountOk returns a tuple with the MaxIngressFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIngressFlowCount

`func (o *MsgVpnClientProfile) SetMaxIngressFlowCount(v int64)`

SetMaxIngressFlowCount sets MaxIngressFlowCount field to given value.

### HasMaxIngressFlowCount

`func (o *MsgVpnClientProfile) HasMaxIngressFlowCount() bool`

HasMaxIngressFlowCount returns a boolean if a field has been set.

### GetMaxMsgsPerTransaction

`func (o *MsgVpnClientProfile) GetMaxMsgsPerTransaction() int32`

GetMaxMsgsPerTransaction returns the MaxMsgsPerTransaction field if non-nil, zero value otherwise.

### GetMaxMsgsPerTransactionOk

`func (o *MsgVpnClientProfile) GetMaxMsgsPerTransactionOk() (*int32, bool)`

GetMaxMsgsPerTransactionOk returns a tuple with the MaxMsgsPerTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgsPerTransaction

`func (o *MsgVpnClientProfile) SetMaxMsgsPerTransaction(v int32)`

SetMaxMsgsPerTransaction sets MaxMsgsPerTransaction field to given value.

### HasMaxMsgsPerTransaction

`func (o *MsgVpnClientProfile) HasMaxMsgsPerTransaction() bool`

HasMaxMsgsPerTransaction returns a boolean if a field has been set.

### GetMaxSubscriptionCount

`func (o *MsgVpnClientProfile) GetMaxSubscriptionCount() int64`

GetMaxSubscriptionCount returns the MaxSubscriptionCount field if non-nil, zero value otherwise.

### GetMaxSubscriptionCountOk

`func (o *MsgVpnClientProfile) GetMaxSubscriptionCountOk() (*int64, bool)`

GetMaxSubscriptionCountOk returns a tuple with the MaxSubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSubscriptionCount

`func (o *MsgVpnClientProfile) SetMaxSubscriptionCount(v int64)`

SetMaxSubscriptionCount sets MaxSubscriptionCount field to given value.

### HasMaxSubscriptionCount

`func (o *MsgVpnClientProfile) HasMaxSubscriptionCount() bool`

HasMaxSubscriptionCount returns a boolean if a field has been set.

### GetMaxTransactedSessionCount

`func (o *MsgVpnClientProfile) GetMaxTransactedSessionCount() int64`

GetMaxTransactedSessionCount returns the MaxTransactedSessionCount field if non-nil, zero value otherwise.

### GetMaxTransactedSessionCountOk

`func (o *MsgVpnClientProfile) GetMaxTransactedSessionCountOk() (*int64, bool)`

GetMaxTransactedSessionCountOk returns a tuple with the MaxTransactedSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactedSessionCount

`func (o *MsgVpnClientProfile) SetMaxTransactedSessionCount(v int64)`

SetMaxTransactedSessionCount sets MaxTransactedSessionCount field to given value.

### HasMaxTransactedSessionCount

`func (o *MsgVpnClientProfile) HasMaxTransactedSessionCount() bool`

HasMaxTransactedSessionCount returns a boolean if a field has been set.

### GetMaxTransactionCount

`func (o *MsgVpnClientProfile) GetMaxTransactionCount() int64`

GetMaxTransactionCount returns the MaxTransactionCount field if non-nil, zero value otherwise.

### GetMaxTransactionCountOk

`func (o *MsgVpnClientProfile) GetMaxTransactionCountOk() (*int64, bool)`

GetMaxTransactionCountOk returns a tuple with the MaxTransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactionCount

`func (o *MsgVpnClientProfile) SetMaxTransactionCount(v int64)`

SetMaxTransactionCount sets MaxTransactionCount field to given value.

### HasMaxTransactionCount

`func (o *MsgVpnClientProfile) HasMaxTransactionCount() bool`

HasMaxTransactionCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnClientProfile) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnClientProfile) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnClientProfile) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnClientProfile) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetQueueControl1MaxDepth

`func (o *MsgVpnClientProfile) GetQueueControl1MaxDepth() int32`

GetQueueControl1MaxDepth returns the QueueControl1MaxDepth field if non-nil, zero value otherwise.

### GetQueueControl1MaxDepthOk

`func (o *MsgVpnClientProfile) GetQueueControl1MaxDepthOk() (*int32, bool)`

GetQueueControl1MaxDepthOk returns a tuple with the QueueControl1MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueControl1MaxDepth

`func (o *MsgVpnClientProfile) SetQueueControl1MaxDepth(v int32)`

SetQueueControl1MaxDepth sets QueueControl1MaxDepth field to given value.

### HasQueueControl1MaxDepth

`func (o *MsgVpnClientProfile) HasQueueControl1MaxDepth() bool`

HasQueueControl1MaxDepth returns a boolean if a field has been set.

### GetQueueControl1MinMsgBurst

`func (o *MsgVpnClientProfile) GetQueueControl1MinMsgBurst() int32`

GetQueueControl1MinMsgBurst returns the QueueControl1MinMsgBurst field if non-nil, zero value otherwise.

### GetQueueControl1MinMsgBurstOk

`func (o *MsgVpnClientProfile) GetQueueControl1MinMsgBurstOk() (*int32, bool)`

GetQueueControl1MinMsgBurstOk returns a tuple with the QueueControl1MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueControl1MinMsgBurst

`func (o *MsgVpnClientProfile) SetQueueControl1MinMsgBurst(v int32)`

SetQueueControl1MinMsgBurst sets QueueControl1MinMsgBurst field to given value.

### HasQueueControl1MinMsgBurst

`func (o *MsgVpnClientProfile) HasQueueControl1MinMsgBurst() bool`

HasQueueControl1MinMsgBurst returns a boolean if a field has been set.

### GetQueueDirect1MaxDepth

`func (o *MsgVpnClientProfile) GetQueueDirect1MaxDepth() int32`

GetQueueDirect1MaxDepth returns the QueueDirect1MaxDepth field if non-nil, zero value otherwise.

### GetQueueDirect1MaxDepthOk

`func (o *MsgVpnClientProfile) GetQueueDirect1MaxDepthOk() (*int32, bool)`

GetQueueDirect1MaxDepthOk returns a tuple with the QueueDirect1MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect1MaxDepth

`func (o *MsgVpnClientProfile) SetQueueDirect1MaxDepth(v int32)`

SetQueueDirect1MaxDepth sets QueueDirect1MaxDepth field to given value.

### HasQueueDirect1MaxDepth

`func (o *MsgVpnClientProfile) HasQueueDirect1MaxDepth() bool`

HasQueueDirect1MaxDepth returns a boolean if a field has been set.

### GetQueueDirect1MinMsgBurst

`func (o *MsgVpnClientProfile) GetQueueDirect1MinMsgBurst() int32`

GetQueueDirect1MinMsgBurst returns the QueueDirect1MinMsgBurst field if non-nil, zero value otherwise.

### GetQueueDirect1MinMsgBurstOk

`func (o *MsgVpnClientProfile) GetQueueDirect1MinMsgBurstOk() (*int32, bool)`

GetQueueDirect1MinMsgBurstOk returns a tuple with the QueueDirect1MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect1MinMsgBurst

`func (o *MsgVpnClientProfile) SetQueueDirect1MinMsgBurst(v int32)`

SetQueueDirect1MinMsgBurst sets QueueDirect1MinMsgBurst field to given value.

### HasQueueDirect1MinMsgBurst

`func (o *MsgVpnClientProfile) HasQueueDirect1MinMsgBurst() bool`

HasQueueDirect1MinMsgBurst returns a boolean if a field has been set.

### GetQueueDirect2MaxDepth

`func (o *MsgVpnClientProfile) GetQueueDirect2MaxDepth() int32`

GetQueueDirect2MaxDepth returns the QueueDirect2MaxDepth field if non-nil, zero value otherwise.

### GetQueueDirect2MaxDepthOk

`func (o *MsgVpnClientProfile) GetQueueDirect2MaxDepthOk() (*int32, bool)`

GetQueueDirect2MaxDepthOk returns a tuple with the QueueDirect2MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect2MaxDepth

`func (o *MsgVpnClientProfile) SetQueueDirect2MaxDepth(v int32)`

SetQueueDirect2MaxDepth sets QueueDirect2MaxDepth field to given value.

### HasQueueDirect2MaxDepth

`func (o *MsgVpnClientProfile) HasQueueDirect2MaxDepth() bool`

HasQueueDirect2MaxDepth returns a boolean if a field has been set.

### GetQueueDirect2MinMsgBurst

`func (o *MsgVpnClientProfile) GetQueueDirect2MinMsgBurst() int32`

GetQueueDirect2MinMsgBurst returns the QueueDirect2MinMsgBurst field if non-nil, zero value otherwise.

### GetQueueDirect2MinMsgBurstOk

`func (o *MsgVpnClientProfile) GetQueueDirect2MinMsgBurstOk() (*int32, bool)`

GetQueueDirect2MinMsgBurstOk returns a tuple with the QueueDirect2MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect2MinMsgBurst

`func (o *MsgVpnClientProfile) SetQueueDirect2MinMsgBurst(v int32)`

SetQueueDirect2MinMsgBurst sets QueueDirect2MinMsgBurst field to given value.

### HasQueueDirect2MinMsgBurst

`func (o *MsgVpnClientProfile) HasQueueDirect2MinMsgBurst() bool`

HasQueueDirect2MinMsgBurst returns a boolean if a field has been set.

### GetQueueDirect3MaxDepth

`func (o *MsgVpnClientProfile) GetQueueDirect3MaxDepth() int32`

GetQueueDirect3MaxDepth returns the QueueDirect3MaxDepth field if non-nil, zero value otherwise.

### GetQueueDirect3MaxDepthOk

`func (o *MsgVpnClientProfile) GetQueueDirect3MaxDepthOk() (*int32, bool)`

GetQueueDirect3MaxDepthOk returns a tuple with the QueueDirect3MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect3MaxDepth

`func (o *MsgVpnClientProfile) SetQueueDirect3MaxDepth(v int32)`

SetQueueDirect3MaxDepth sets QueueDirect3MaxDepth field to given value.

### HasQueueDirect3MaxDepth

`func (o *MsgVpnClientProfile) HasQueueDirect3MaxDepth() bool`

HasQueueDirect3MaxDepth returns a boolean if a field has been set.

### GetQueueDirect3MinMsgBurst

`func (o *MsgVpnClientProfile) GetQueueDirect3MinMsgBurst() int32`

GetQueueDirect3MinMsgBurst returns the QueueDirect3MinMsgBurst field if non-nil, zero value otherwise.

### GetQueueDirect3MinMsgBurstOk

`func (o *MsgVpnClientProfile) GetQueueDirect3MinMsgBurstOk() (*int32, bool)`

GetQueueDirect3MinMsgBurstOk returns a tuple with the QueueDirect3MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirect3MinMsgBurst

`func (o *MsgVpnClientProfile) SetQueueDirect3MinMsgBurst(v int32)`

SetQueueDirect3MinMsgBurst sets QueueDirect3MinMsgBurst field to given value.

### HasQueueDirect3MinMsgBurst

`func (o *MsgVpnClientProfile) HasQueueDirect3MinMsgBurst() bool`

HasQueueDirect3MinMsgBurst returns a boolean if a field has been set.

### GetQueueGuaranteed1MaxDepth

`func (o *MsgVpnClientProfile) GetQueueGuaranteed1MaxDepth() int32`

GetQueueGuaranteed1MaxDepth returns the QueueGuaranteed1MaxDepth field if non-nil, zero value otherwise.

### GetQueueGuaranteed1MaxDepthOk

`func (o *MsgVpnClientProfile) GetQueueGuaranteed1MaxDepthOk() (*int32, bool)`

GetQueueGuaranteed1MaxDepthOk returns a tuple with the QueueGuaranteed1MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueGuaranteed1MaxDepth

`func (o *MsgVpnClientProfile) SetQueueGuaranteed1MaxDepth(v int32)`

SetQueueGuaranteed1MaxDepth sets QueueGuaranteed1MaxDepth field to given value.

### HasQueueGuaranteed1MaxDepth

`func (o *MsgVpnClientProfile) HasQueueGuaranteed1MaxDepth() bool`

HasQueueGuaranteed1MaxDepth returns a boolean if a field has been set.

### GetQueueGuaranteed1MinMsgBurst

`func (o *MsgVpnClientProfile) GetQueueGuaranteed1MinMsgBurst() int32`

GetQueueGuaranteed1MinMsgBurst returns the QueueGuaranteed1MinMsgBurst field if non-nil, zero value otherwise.

### GetQueueGuaranteed1MinMsgBurstOk

`func (o *MsgVpnClientProfile) GetQueueGuaranteed1MinMsgBurstOk() (*int32, bool)`

GetQueueGuaranteed1MinMsgBurstOk returns a tuple with the QueueGuaranteed1MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueGuaranteed1MinMsgBurst

`func (o *MsgVpnClientProfile) SetQueueGuaranteed1MinMsgBurst(v int32)`

SetQueueGuaranteed1MinMsgBurst sets QueueGuaranteed1MinMsgBurst field to given value.

### HasQueueGuaranteed1MinMsgBurst

`func (o *MsgVpnClientProfile) HasQueueGuaranteed1MinMsgBurst() bool`

HasQueueGuaranteed1MinMsgBurst returns a boolean if a field has been set.

### GetRejectMsgToSenderOnNoSubscriptionMatchEnabled

`func (o *MsgVpnClientProfile) GetRejectMsgToSenderOnNoSubscriptionMatchEnabled() bool`

GetRejectMsgToSenderOnNoSubscriptionMatchEnabled returns the RejectMsgToSenderOnNoSubscriptionMatchEnabled field if non-nil, zero value otherwise.

### GetRejectMsgToSenderOnNoSubscriptionMatchEnabledOk

`func (o *MsgVpnClientProfile) GetRejectMsgToSenderOnNoSubscriptionMatchEnabledOk() (*bool, bool)`

GetRejectMsgToSenderOnNoSubscriptionMatchEnabledOk returns a tuple with the RejectMsgToSenderOnNoSubscriptionMatchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectMsgToSenderOnNoSubscriptionMatchEnabled

`func (o *MsgVpnClientProfile) SetRejectMsgToSenderOnNoSubscriptionMatchEnabled(v bool)`

SetRejectMsgToSenderOnNoSubscriptionMatchEnabled sets RejectMsgToSenderOnNoSubscriptionMatchEnabled field to given value.

### HasRejectMsgToSenderOnNoSubscriptionMatchEnabled

`func (o *MsgVpnClientProfile) HasRejectMsgToSenderOnNoSubscriptionMatchEnabled() bool`

HasRejectMsgToSenderOnNoSubscriptionMatchEnabled returns a boolean if a field has been set.

### GetReplicationAllowClientConnectWhenStandbyEnabled

`func (o *MsgVpnClientProfile) GetReplicationAllowClientConnectWhenStandbyEnabled() bool`

GetReplicationAllowClientConnectWhenStandbyEnabled returns the ReplicationAllowClientConnectWhenStandbyEnabled field if non-nil, zero value otherwise.

### GetReplicationAllowClientConnectWhenStandbyEnabledOk

`func (o *MsgVpnClientProfile) GetReplicationAllowClientConnectWhenStandbyEnabledOk() (*bool, bool)`

GetReplicationAllowClientConnectWhenStandbyEnabledOk returns a tuple with the ReplicationAllowClientConnectWhenStandbyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationAllowClientConnectWhenStandbyEnabled

`func (o *MsgVpnClientProfile) SetReplicationAllowClientConnectWhenStandbyEnabled(v bool)`

SetReplicationAllowClientConnectWhenStandbyEnabled sets ReplicationAllowClientConnectWhenStandbyEnabled field to given value.

### HasReplicationAllowClientConnectWhenStandbyEnabled

`func (o *MsgVpnClientProfile) HasReplicationAllowClientConnectWhenStandbyEnabled() bool`

HasReplicationAllowClientConnectWhenStandbyEnabled returns a boolean if a field has been set.

### GetServiceMinKeepaliveTimeout

`func (o *MsgVpnClientProfile) GetServiceMinKeepaliveTimeout() int32`

GetServiceMinKeepaliveTimeout returns the ServiceMinKeepaliveTimeout field if non-nil, zero value otherwise.

### GetServiceMinKeepaliveTimeoutOk

`func (o *MsgVpnClientProfile) GetServiceMinKeepaliveTimeoutOk() (*int32, bool)`

GetServiceMinKeepaliveTimeoutOk returns a tuple with the ServiceMinKeepaliveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMinKeepaliveTimeout

`func (o *MsgVpnClientProfile) SetServiceMinKeepaliveTimeout(v int32)`

SetServiceMinKeepaliveTimeout sets ServiceMinKeepaliveTimeout field to given value.

### HasServiceMinKeepaliveTimeout

`func (o *MsgVpnClientProfile) HasServiceMinKeepaliveTimeout() bool`

HasServiceMinKeepaliveTimeout returns a boolean if a field has been set.

### GetServiceSmfMaxConnectionCountPerClientUsername

`func (o *MsgVpnClientProfile) GetServiceSmfMaxConnectionCountPerClientUsername() int64`

GetServiceSmfMaxConnectionCountPerClientUsername returns the ServiceSmfMaxConnectionCountPerClientUsername field if non-nil, zero value otherwise.

### GetServiceSmfMaxConnectionCountPerClientUsernameOk

`func (o *MsgVpnClientProfile) GetServiceSmfMaxConnectionCountPerClientUsernameOk() (*int64, bool)`

GetServiceSmfMaxConnectionCountPerClientUsernameOk returns a tuple with the ServiceSmfMaxConnectionCountPerClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfMaxConnectionCountPerClientUsername

`func (o *MsgVpnClientProfile) SetServiceSmfMaxConnectionCountPerClientUsername(v int64)`

SetServiceSmfMaxConnectionCountPerClientUsername sets ServiceSmfMaxConnectionCountPerClientUsername field to given value.

### HasServiceSmfMaxConnectionCountPerClientUsername

`func (o *MsgVpnClientProfile) HasServiceSmfMaxConnectionCountPerClientUsername() bool`

HasServiceSmfMaxConnectionCountPerClientUsername returns a boolean if a field has been set.

### GetServiceSmfMinKeepaliveEnabled

`func (o *MsgVpnClientProfile) GetServiceSmfMinKeepaliveEnabled() bool`

GetServiceSmfMinKeepaliveEnabled returns the ServiceSmfMinKeepaliveEnabled field if non-nil, zero value otherwise.

### GetServiceSmfMinKeepaliveEnabledOk

`func (o *MsgVpnClientProfile) GetServiceSmfMinKeepaliveEnabledOk() (*bool, bool)`

GetServiceSmfMinKeepaliveEnabledOk returns a tuple with the ServiceSmfMinKeepaliveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfMinKeepaliveEnabled

`func (o *MsgVpnClientProfile) SetServiceSmfMinKeepaliveEnabled(v bool)`

SetServiceSmfMinKeepaliveEnabled sets ServiceSmfMinKeepaliveEnabled field to given value.

### HasServiceSmfMinKeepaliveEnabled

`func (o *MsgVpnClientProfile) HasServiceSmfMinKeepaliveEnabled() bool`

HasServiceSmfMinKeepaliveEnabled returns a boolean if a field has been set.

### GetServiceWebInactiveTimeout

`func (o *MsgVpnClientProfile) GetServiceWebInactiveTimeout() int64`

GetServiceWebInactiveTimeout returns the ServiceWebInactiveTimeout field if non-nil, zero value otherwise.

### GetServiceWebInactiveTimeoutOk

`func (o *MsgVpnClientProfile) GetServiceWebInactiveTimeoutOk() (*int64, bool)`

GetServiceWebInactiveTimeoutOk returns a tuple with the ServiceWebInactiveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebInactiveTimeout

`func (o *MsgVpnClientProfile) SetServiceWebInactiveTimeout(v int64)`

SetServiceWebInactiveTimeout sets ServiceWebInactiveTimeout field to given value.

### HasServiceWebInactiveTimeout

`func (o *MsgVpnClientProfile) HasServiceWebInactiveTimeout() bool`

HasServiceWebInactiveTimeout returns a boolean if a field has been set.

### GetServiceWebMaxConnectionCountPerClientUsername

`func (o *MsgVpnClientProfile) GetServiceWebMaxConnectionCountPerClientUsername() int64`

GetServiceWebMaxConnectionCountPerClientUsername returns the ServiceWebMaxConnectionCountPerClientUsername field if non-nil, zero value otherwise.

### GetServiceWebMaxConnectionCountPerClientUsernameOk

`func (o *MsgVpnClientProfile) GetServiceWebMaxConnectionCountPerClientUsernameOk() (*int64, bool)`

GetServiceWebMaxConnectionCountPerClientUsernameOk returns a tuple with the ServiceWebMaxConnectionCountPerClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebMaxConnectionCountPerClientUsername

`func (o *MsgVpnClientProfile) SetServiceWebMaxConnectionCountPerClientUsername(v int64)`

SetServiceWebMaxConnectionCountPerClientUsername sets ServiceWebMaxConnectionCountPerClientUsername field to given value.

### HasServiceWebMaxConnectionCountPerClientUsername

`func (o *MsgVpnClientProfile) HasServiceWebMaxConnectionCountPerClientUsername() bool`

HasServiceWebMaxConnectionCountPerClientUsername returns a boolean if a field has been set.

### GetServiceWebMaxPayload

`func (o *MsgVpnClientProfile) GetServiceWebMaxPayload() int64`

GetServiceWebMaxPayload returns the ServiceWebMaxPayload field if non-nil, zero value otherwise.

### GetServiceWebMaxPayloadOk

`func (o *MsgVpnClientProfile) GetServiceWebMaxPayloadOk() (*int64, bool)`

GetServiceWebMaxPayloadOk returns a tuple with the ServiceWebMaxPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebMaxPayload

`func (o *MsgVpnClientProfile) SetServiceWebMaxPayload(v int64)`

SetServiceWebMaxPayload sets ServiceWebMaxPayload field to given value.

### HasServiceWebMaxPayload

`func (o *MsgVpnClientProfile) HasServiceWebMaxPayload() bool`

HasServiceWebMaxPayload returns a boolean if a field has been set.

### GetTcpCongestionWindowSize

`func (o *MsgVpnClientProfile) GetTcpCongestionWindowSize() int64`

GetTcpCongestionWindowSize returns the TcpCongestionWindowSize field if non-nil, zero value otherwise.

### GetTcpCongestionWindowSizeOk

`func (o *MsgVpnClientProfile) GetTcpCongestionWindowSizeOk() (*int64, bool)`

GetTcpCongestionWindowSizeOk returns a tuple with the TcpCongestionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpCongestionWindowSize

`func (o *MsgVpnClientProfile) SetTcpCongestionWindowSize(v int64)`

SetTcpCongestionWindowSize sets TcpCongestionWindowSize field to given value.

### HasTcpCongestionWindowSize

`func (o *MsgVpnClientProfile) HasTcpCongestionWindowSize() bool`

HasTcpCongestionWindowSize returns a boolean if a field has been set.

### GetTcpKeepaliveCount

`func (o *MsgVpnClientProfile) GetTcpKeepaliveCount() int64`

GetTcpKeepaliveCount returns the TcpKeepaliveCount field if non-nil, zero value otherwise.

### GetTcpKeepaliveCountOk

`func (o *MsgVpnClientProfile) GetTcpKeepaliveCountOk() (*int64, bool)`

GetTcpKeepaliveCountOk returns a tuple with the TcpKeepaliveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpKeepaliveCount

`func (o *MsgVpnClientProfile) SetTcpKeepaliveCount(v int64)`

SetTcpKeepaliveCount sets TcpKeepaliveCount field to given value.

### HasTcpKeepaliveCount

`func (o *MsgVpnClientProfile) HasTcpKeepaliveCount() bool`

HasTcpKeepaliveCount returns a boolean if a field has been set.

### GetTcpKeepaliveIdleTime

`func (o *MsgVpnClientProfile) GetTcpKeepaliveIdleTime() int64`

GetTcpKeepaliveIdleTime returns the TcpKeepaliveIdleTime field if non-nil, zero value otherwise.

### GetTcpKeepaliveIdleTimeOk

`func (o *MsgVpnClientProfile) GetTcpKeepaliveIdleTimeOk() (*int64, bool)`

GetTcpKeepaliveIdleTimeOk returns a tuple with the TcpKeepaliveIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpKeepaliveIdleTime

`func (o *MsgVpnClientProfile) SetTcpKeepaliveIdleTime(v int64)`

SetTcpKeepaliveIdleTime sets TcpKeepaliveIdleTime field to given value.

### HasTcpKeepaliveIdleTime

`func (o *MsgVpnClientProfile) HasTcpKeepaliveIdleTime() bool`

HasTcpKeepaliveIdleTime returns a boolean if a field has been set.

### GetTcpKeepaliveInterval

`func (o *MsgVpnClientProfile) GetTcpKeepaliveInterval() int64`

GetTcpKeepaliveInterval returns the TcpKeepaliveInterval field if non-nil, zero value otherwise.

### GetTcpKeepaliveIntervalOk

`func (o *MsgVpnClientProfile) GetTcpKeepaliveIntervalOk() (*int64, bool)`

GetTcpKeepaliveIntervalOk returns a tuple with the TcpKeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpKeepaliveInterval

`func (o *MsgVpnClientProfile) SetTcpKeepaliveInterval(v int64)`

SetTcpKeepaliveInterval sets TcpKeepaliveInterval field to given value.

### HasTcpKeepaliveInterval

`func (o *MsgVpnClientProfile) HasTcpKeepaliveInterval() bool`

HasTcpKeepaliveInterval returns a boolean if a field has been set.

### GetTcpMaxSegmentSize

`func (o *MsgVpnClientProfile) GetTcpMaxSegmentSize() int64`

GetTcpMaxSegmentSize returns the TcpMaxSegmentSize field if non-nil, zero value otherwise.

### GetTcpMaxSegmentSizeOk

`func (o *MsgVpnClientProfile) GetTcpMaxSegmentSizeOk() (*int64, bool)`

GetTcpMaxSegmentSizeOk returns a tuple with the TcpMaxSegmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMaxSegmentSize

`func (o *MsgVpnClientProfile) SetTcpMaxSegmentSize(v int64)`

SetTcpMaxSegmentSize sets TcpMaxSegmentSize field to given value.

### HasTcpMaxSegmentSize

`func (o *MsgVpnClientProfile) HasTcpMaxSegmentSize() bool`

HasTcpMaxSegmentSize returns a boolean if a field has been set.

### GetTcpMaxWindowSize

`func (o *MsgVpnClientProfile) GetTcpMaxWindowSize() int64`

GetTcpMaxWindowSize returns the TcpMaxWindowSize field if non-nil, zero value otherwise.

### GetTcpMaxWindowSizeOk

`func (o *MsgVpnClientProfile) GetTcpMaxWindowSizeOk() (*int64, bool)`

GetTcpMaxWindowSizeOk returns a tuple with the TcpMaxWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMaxWindowSize

`func (o *MsgVpnClientProfile) SetTcpMaxWindowSize(v int64)`

SetTcpMaxWindowSize sets TcpMaxWindowSize field to given value.

### HasTcpMaxWindowSize

`func (o *MsgVpnClientProfile) HasTcpMaxWindowSize() bool`

HasTcpMaxWindowSize returns a boolean if a field has been set.

### GetTlsAllowDowngradeToPlainTextEnabled

`func (o *MsgVpnClientProfile) GetTlsAllowDowngradeToPlainTextEnabled() bool`

GetTlsAllowDowngradeToPlainTextEnabled returns the TlsAllowDowngradeToPlainTextEnabled field if non-nil, zero value otherwise.

### GetTlsAllowDowngradeToPlainTextEnabledOk

`func (o *MsgVpnClientProfile) GetTlsAllowDowngradeToPlainTextEnabledOk() (*bool, bool)`

GetTlsAllowDowngradeToPlainTextEnabledOk returns a tuple with the TlsAllowDowngradeToPlainTextEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAllowDowngradeToPlainTextEnabled

`func (o *MsgVpnClientProfile) SetTlsAllowDowngradeToPlainTextEnabled(v bool)`

SetTlsAllowDowngradeToPlainTextEnabled sets TlsAllowDowngradeToPlainTextEnabled field to given value.

### HasTlsAllowDowngradeToPlainTextEnabled

`func (o *MsgVpnClientProfile) HasTlsAllowDowngradeToPlainTextEnabled() bool`

HasTlsAllowDowngradeToPlainTextEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


