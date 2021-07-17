# DmrClusterLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationBasicPassword** | Pointer to **string** | The password used to authenticate with the remote node when using basic internal authentication. If this per-Link password is not configured, the Cluster&#39;s password is used instead. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**AuthenticationScheme** | Pointer to **string** | The authentication scheme to be used by the Link which initiates connections to the remote node. The default value is &#x60;\&quot;basic\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;basic\&quot; - Basic Authentication Scheme (via username and password). \&quot;client-certificate\&quot; - Client Certificate Authentication Scheme (via certificate file or content). &lt;/pre&gt;  | [optional] 
**ClientProfileQueueControl1MaxDepth** | Pointer to **int32** | The maximum depth of the \&quot;Control 1\&quot; (C-1) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is &#x60;20000&#x60;. | [optional] 
**ClientProfileQueueControl1MinMsgBurst** | Pointer to **int32** | The number of messages that are always allowed entry into the \&quot;Control 1\&quot; (C-1) priority queue, regardless of the &#x60;clientProfileQueueControl1MaxDepth&#x60; value. The default value is &#x60;4&#x60;. | [optional] 
**ClientProfileQueueDirect1MaxDepth** | Pointer to **int32** | The maximum depth of the \&quot;Direct 1\&quot; (D-1) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is &#x60;20000&#x60;. | [optional] 
**ClientProfileQueueDirect1MinMsgBurst** | Pointer to **int32** | The number of messages that are always allowed entry into the \&quot;Direct 1\&quot; (D-1) priority queue, regardless of the &#x60;clientProfileQueueDirect1MaxDepth&#x60; value. The default value is &#x60;4&#x60;. | [optional] 
**ClientProfileQueueDirect2MaxDepth** | Pointer to **int32** | The maximum depth of the \&quot;Direct 2\&quot; (D-2) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is &#x60;20000&#x60;. | [optional] 
**ClientProfileQueueDirect2MinMsgBurst** | Pointer to **int32** | The number of messages that are always allowed entry into the \&quot;Direct 2\&quot; (D-2) priority queue, regardless of the &#x60;clientProfileQueueDirect2MaxDepth&#x60; value. The default value is &#x60;4&#x60;. | [optional] 
**ClientProfileQueueDirect3MaxDepth** | Pointer to **int32** | The maximum depth of the \&quot;Direct 3\&quot; (D-3) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is &#x60;20000&#x60;. | [optional] 
**ClientProfileQueueDirect3MinMsgBurst** | Pointer to **int32** | The number of messages that are always allowed entry into the \&quot;Direct 3\&quot; (D-3) priority queue, regardless of the &#x60;clientProfileQueueDirect3MaxDepth&#x60; value. The default value is &#x60;4&#x60;. | [optional] 
**ClientProfileQueueGuaranteed1MaxDepth** | Pointer to **int32** | The maximum depth of the \&quot;Guaranteed 1\&quot; (G-1) priority queue, in work units. Each work unit is 2048 bytes of message data. The default value is &#x60;20000&#x60;. | [optional] 
**ClientProfileQueueGuaranteed1MinMsgBurst** | Pointer to **int32** | The number of messages that are always allowed entry into the \&quot;Guaranteed 1\&quot; (G-3) priority queue, regardless of the &#x60;clientProfileQueueGuaranteed1MaxDepth&#x60; value. The default value is &#x60;255&#x60;. | [optional] 
**ClientProfileTcpCongestionWindowSize** | Pointer to **int64** | The TCP initial congestion window size, in multiples of the TCP Maximum Segment Size (MSS). Changing the value from its default of 2 results in non-compliance with RFC 2581. Contact Solace Support before changing this value. The default value is &#x60;2&#x60;. | [optional] 
**ClientProfileTcpKeepaliveCount** | Pointer to **int64** | The number of TCP keepalive retransmissions to be carried out before declaring that the remote end is not available. The default value is &#x60;5&#x60;. | [optional] 
**ClientProfileTcpKeepaliveIdleTime** | Pointer to **int64** | The amount of time a connection must remain idle before TCP begins sending keepalive probes, in seconds. The default value is &#x60;3&#x60;. | [optional] 
**ClientProfileTcpKeepaliveInterval** | Pointer to **int64** | The amount of time between TCP keepalive retransmissions when no acknowledgement is received, in seconds. The default value is &#x60;1&#x60;. | [optional] 
**ClientProfileTcpMaxSegmentSize** | Pointer to **int64** | The TCP maximum segment size, in bytes. Changes are applied to all existing connections. The default value is &#x60;1460&#x60;. | [optional] 
**ClientProfileTcpMaxWindowSize** | Pointer to **int64** | The TCP maximum window size, in kilobytes. Changes are applied to all existing connections. The default value is &#x60;256&#x60;. | [optional] 
**DmrClusterName** | Pointer to **string** | The name of the Cluster. | [optional] 
**EgressFlowWindowSize** | Pointer to **int64** | The number of outstanding guaranteed messages that can be sent over the Link before acknowledgement is received by the sender. The default value is &#x60;255&#x60;. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the Link. When disabled, subscription sets of this and the remote node are not kept up-to-date, and messages are not exchanged with the remote node. Published guaranteed messages will be queued up for future delivery based on current subscription sets. The default value is &#x60;false&#x60;. | [optional] 
**Initiator** | Pointer to **string** | The initiator of the Link&#39;s TCP connections. The default value is &#x60;\&quot;lexical\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;lexical\&quot; - The \&quot;higher\&quot; node-name initiates. \&quot;local\&quot; - The local node initiates. \&quot;remote\&quot; - The remote node initiates. &lt;/pre&gt;  | [optional] 
**QueueDeadMsgQueue** | Pointer to **string** | The name of the Dead Message Queue (DMQ) used by the Queue for discarded messages. The default value is &#x60;\&quot;#DEAD_MSG_QUEUE\&quot;&#x60;. | [optional] 
**QueueEventSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**QueueMaxDeliveredUnackedMsgsPerFlow** | Pointer to **int64** | The maximum number of messages delivered but not acknowledged per flow for the Queue. The default value is &#x60;1000000&#x60;. | [optional] 
**QueueMaxMsgSpoolUsage** | Pointer to **int64** | The maximum message spool usage by the Queue (quota), in megabytes (MB). The default varies by platform. | [optional] 
**QueueMaxRedeliveryCount** | Pointer to **int64** | The maximum number of times the Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. The default value is &#x60;0&#x60;. | [optional] 
**QueueMaxTtl** | Pointer to **int64** | The maximum time in seconds a message can stay in the Queue when &#x60;queueRespectTtlEnabled&#x60; is &#x60;true&#x60;. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the &#x60;queueMaxTtl&#x60; configured for the Queue, is exceeded. A value of 0 disables expiry. The default value is &#x60;0&#x60;. | [optional] 
**QueueRejectMsgToSenderOnDiscardBehavior** | Pointer to **string** | Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The default value is &#x60;\&quot;always\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always return a negative acknowledgment (NACK) to the sending client on message discard. \&quot;when-queue-enabled\&quot; - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Queue is enabled. \&quot;never\&quot; - Never return a negative acknowledgment (NACK) to the sending client on message discard. &lt;/pre&gt;  | [optional] 
**QueueRespectTtlEnabled** | Pointer to **bool** | Enable or disable the respecting of the time-to-live (TTL) for messages in the Queue. When enabled, expired messages are discarded or moved to the DMQ. The default value is &#x60;false&#x60;. | [optional] 
**RemoteNodeName** | Pointer to **string** | The name of the node at the remote end of the Link. | [optional] 
**Span** | Pointer to **string** | The span of the Link, either internal or external. Internal Links connect nodes within the same Cluster. External Links connect nodes within different Clusters. The default value is &#x60;\&quot;external\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;internal\&quot; - Link to same cluster. \&quot;external\&quot; - Link to other cluster. &lt;/pre&gt;  | [optional] 
**TransportCompressedEnabled** | Pointer to **bool** | Enable or disable compression on the Link. The default value is &#x60;false&#x60;. | [optional] 
**TransportTlsEnabled** | Pointer to **bool** | Enable or disable encryption (TLS) on the Link. The default value is &#x60;false&#x60;. | [optional] 

## Methods

### NewDmrClusterLink

`func NewDmrClusterLink() *DmrClusterLink`

NewDmrClusterLink instantiates a new DmrClusterLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterLinkWithDefaults

`func NewDmrClusterLinkWithDefaults() *DmrClusterLink`

NewDmrClusterLinkWithDefaults instantiates a new DmrClusterLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationBasicPassword

`func (o *DmrClusterLink) GetAuthenticationBasicPassword() string`

GetAuthenticationBasicPassword returns the AuthenticationBasicPassword field if non-nil, zero value otherwise.

### GetAuthenticationBasicPasswordOk

`func (o *DmrClusterLink) GetAuthenticationBasicPasswordOk() (*string, bool)`

GetAuthenticationBasicPasswordOk returns a tuple with the AuthenticationBasicPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBasicPassword

`func (o *DmrClusterLink) SetAuthenticationBasicPassword(v string)`

SetAuthenticationBasicPassword sets AuthenticationBasicPassword field to given value.

### HasAuthenticationBasicPassword

`func (o *DmrClusterLink) HasAuthenticationBasicPassword() bool`

HasAuthenticationBasicPassword returns a boolean if a field has been set.

### GetAuthenticationScheme

`func (o *DmrClusterLink) GetAuthenticationScheme() string`

GetAuthenticationScheme returns the AuthenticationScheme field if non-nil, zero value otherwise.

### GetAuthenticationSchemeOk

`func (o *DmrClusterLink) GetAuthenticationSchemeOk() (*string, bool)`

GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationScheme

`func (o *DmrClusterLink) SetAuthenticationScheme(v string)`

SetAuthenticationScheme sets AuthenticationScheme field to given value.

### HasAuthenticationScheme

`func (o *DmrClusterLink) HasAuthenticationScheme() bool`

HasAuthenticationScheme returns a boolean if a field has been set.

### GetClientProfileQueueControl1MaxDepth

`func (o *DmrClusterLink) GetClientProfileQueueControl1MaxDepth() int32`

GetClientProfileQueueControl1MaxDepth returns the ClientProfileQueueControl1MaxDepth field if non-nil, zero value otherwise.

### GetClientProfileQueueControl1MaxDepthOk

`func (o *DmrClusterLink) GetClientProfileQueueControl1MaxDepthOk() (*int32, bool)`

GetClientProfileQueueControl1MaxDepthOk returns a tuple with the ClientProfileQueueControl1MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileQueueControl1MaxDepth

`func (o *DmrClusterLink) SetClientProfileQueueControl1MaxDepth(v int32)`

SetClientProfileQueueControl1MaxDepth sets ClientProfileQueueControl1MaxDepth field to given value.

### HasClientProfileQueueControl1MaxDepth

`func (o *DmrClusterLink) HasClientProfileQueueControl1MaxDepth() bool`

HasClientProfileQueueControl1MaxDepth returns a boolean if a field has been set.

### GetClientProfileQueueControl1MinMsgBurst

`func (o *DmrClusterLink) GetClientProfileQueueControl1MinMsgBurst() int32`

GetClientProfileQueueControl1MinMsgBurst returns the ClientProfileQueueControl1MinMsgBurst field if non-nil, zero value otherwise.

### GetClientProfileQueueControl1MinMsgBurstOk

`func (o *DmrClusterLink) GetClientProfileQueueControl1MinMsgBurstOk() (*int32, bool)`

GetClientProfileQueueControl1MinMsgBurstOk returns a tuple with the ClientProfileQueueControl1MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileQueueControl1MinMsgBurst

`func (o *DmrClusterLink) SetClientProfileQueueControl1MinMsgBurst(v int32)`

SetClientProfileQueueControl1MinMsgBurst sets ClientProfileQueueControl1MinMsgBurst field to given value.

### HasClientProfileQueueControl1MinMsgBurst

`func (o *DmrClusterLink) HasClientProfileQueueControl1MinMsgBurst() bool`

HasClientProfileQueueControl1MinMsgBurst returns a boolean if a field has been set.

### GetClientProfileQueueDirect1MaxDepth

`func (o *DmrClusterLink) GetClientProfileQueueDirect1MaxDepth() int32`

GetClientProfileQueueDirect1MaxDepth returns the ClientProfileQueueDirect1MaxDepth field if non-nil, zero value otherwise.

### GetClientProfileQueueDirect1MaxDepthOk

`func (o *DmrClusterLink) GetClientProfileQueueDirect1MaxDepthOk() (*int32, bool)`

GetClientProfileQueueDirect1MaxDepthOk returns a tuple with the ClientProfileQueueDirect1MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileQueueDirect1MaxDepth

`func (o *DmrClusterLink) SetClientProfileQueueDirect1MaxDepth(v int32)`

SetClientProfileQueueDirect1MaxDepth sets ClientProfileQueueDirect1MaxDepth field to given value.

### HasClientProfileQueueDirect1MaxDepth

`func (o *DmrClusterLink) HasClientProfileQueueDirect1MaxDepth() bool`

HasClientProfileQueueDirect1MaxDepth returns a boolean if a field has been set.

### GetClientProfileQueueDirect1MinMsgBurst

`func (o *DmrClusterLink) GetClientProfileQueueDirect1MinMsgBurst() int32`

GetClientProfileQueueDirect1MinMsgBurst returns the ClientProfileQueueDirect1MinMsgBurst field if non-nil, zero value otherwise.

### GetClientProfileQueueDirect1MinMsgBurstOk

`func (o *DmrClusterLink) GetClientProfileQueueDirect1MinMsgBurstOk() (*int32, bool)`

GetClientProfileQueueDirect1MinMsgBurstOk returns a tuple with the ClientProfileQueueDirect1MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileQueueDirect1MinMsgBurst

`func (o *DmrClusterLink) SetClientProfileQueueDirect1MinMsgBurst(v int32)`

SetClientProfileQueueDirect1MinMsgBurst sets ClientProfileQueueDirect1MinMsgBurst field to given value.

### HasClientProfileQueueDirect1MinMsgBurst

`func (o *DmrClusterLink) HasClientProfileQueueDirect1MinMsgBurst() bool`

HasClientProfileQueueDirect1MinMsgBurst returns a boolean if a field has been set.

### GetClientProfileQueueDirect2MaxDepth

`func (o *DmrClusterLink) GetClientProfileQueueDirect2MaxDepth() int32`

GetClientProfileQueueDirect2MaxDepth returns the ClientProfileQueueDirect2MaxDepth field if non-nil, zero value otherwise.

### GetClientProfileQueueDirect2MaxDepthOk

`func (o *DmrClusterLink) GetClientProfileQueueDirect2MaxDepthOk() (*int32, bool)`

GetClientProfileQueueDirect2MaxDepthOk returns a tuple with the ClientProfileQueueDirect2MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileQueueDirect2MaxDepth

`func (o *DmrClusterLink) SetClientProfileQueueDirect2MaxDepth(v int32)`

SetClientProfileQueueDirect2MaxDepth sets ClientProfileQueueDirect2MaxDepth field to given value.

### HasClientProfileQueueDirect2MaxDepth

`func (o *DmrClusterLink) HasClientProfileQueueDirect2MaxDepth() bool`

HasClientProfileQueueDirect2MaxDepth returns a boolean if a field has been set.

### GetClientProfileQueueDirect2MinMsgBurst

`func (o *DmrClusterLink) GetClientProfileQueueDirect2MinMsgBurst() int32`

GetClientProfileQueueDirect2MinMsgBurst returns the ClientProfileQueueDirect2MinMsgBurst field if non-nil, zero value otherwise.

### GetClientProfileQueueDirect2MinMsgBurstOk

`func (o *DmrClusterLink) GetClientProfileQueueDirect2MinMsgBurstOk() (*int32, bool)`

GetClientProfileQueueDirect2MinMsgBurstOk returns a tuple with the ClientProfileQueueDirect2MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileQueueDirect2MinMsgBurst

`func (o *DmrClusterLink) SetClientProfileQueueDirect2MinMsgBurst(v int32)`

SetClientProfileQueueDirect2MinMsgBurst sets ClientProfileQueueDirect2MinMsgBurst field to given value.

### HasClientProfileQueueDirect2MinMsgBurst

`func (o *DmrClusterLink) HasClientProfileQueueDirect2MinMsgBurst() bool`

HasClientProfileQueueDirect2MinMsgBurst returns a boolean if a field has been set.

### GetClientProfileQueueDirect3MaxDepth

`func (o *DmrClusterLink) GetClientProfileQueueDirect3MaxDepth() int32`

GetClientProfileQueueDirect3MaxDepth returns the ClientProfileQueueDirect3MaxDepth field if non-nil, zero value otherwise.

### GetClientProfileQueueDirect3MaxDepthOk

`func (o *DmrClusterLink) GetClientProfileQueueDirect3MaxDepthOk() (*int32, bool)`

GetClientProfileQueueDirect3MaxDepthOk returns a tuple with the ClientProfileQueueDirect3MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileQueueDirect3MaxDepth

`func (o *DmrClusterLink) SetClientProfileQueueDirect3MaxDepth(v int32)`

SetClientProfileQueueDirect3MaxDepth sets ClientProfileQueueDirect3MaxDepth field to given value.

### HasClientProfileQueueDirect3MaxDepth

`func (o *DmrClusterLink) HasClientProfileQueueDirect3MaxDepth() bool`

HasClientProfileQueueDirect3MaxDepth returns a boolean if a field has been set.

### GetClientProfileQueueDirect3MinMsgBurst

`func (o *DmrClusterLink) GetClientProfileQueueDirect3MinMsgBurst() int32`

GetClientProfileQueueDirect3MinMsgBurst returns the ClientProfileQueueDirect3MinMsgBurst field if non-nil, zero value otherwise.

### GetClientProfileQueueDirect3MinMsgBurstOk

`func (o *DmrClusterLink) GetClientProfileQueueDirect3MinMsgBurstOk() (*int32, bool)`

GetClientProfileQueueDirect3MinMsgBurstOk returns a tuple with the ClientProfileQueueDirect3MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileQueueDirect3MinMsgBurst

`func (o *DmrClusterLink) SetClientProfileQueueDirect3MinMsgBurst(v int32)`

SetClientProfileQueueDirect3MinMsgBurst sets ClientProfileQueueDirect3MinMsgBurst field to given value.

### HasClientProfileQueueDirect3MinMsgBurst

`func (o *DmrClusterLink) HasClientProfileQueueDirect3MinMsgBurst() bool`

HasClientProfileQueueDirect3MinMsgBurst returns a boolean if a field has been set.

### GetClientProfileQueueGuaranteed1MaxDepth

`func (o *DmrClusterLink) GetClientProfileQueueGuaranteed1MaxDepth() int32`

GetClientProfileQueueGuaranteed1MaxDepth returns the ClientProfileQueueGuaranteed1MaxDepth field if non-nil, zero value otherwise.

### GetClientProfileQueueGuaranteed1MaxDepthOk

`func (o *DmrClusterLink) GetClientProfileQueueGuaranteed1MaxDepthOk() (*int32, bool)`

GetClientProfileQueueGuaranteed1MaxDepthOk returns a tuple with the ClientProfileQueueGuaranteed1MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileQueueGuaranteed1MaxDepth

`func (o *DmrClusterLink) SetClientProfileQueueGuaranteed1MaxDepth(v int32)`

SetClientProfileQueueGuaranteed1MaxDepth sets ClientProfileQueueGuaranteed1MaxDepth field to given value.

### HasClientProfileQueueGuaranteed1MaxDepth

`func (o *DmrClusterLink) HasClientProfileQueueGuaranteed1MaxDepth() bool`

HasClientProfileQueueGuaranteed1MaxDepth returns a boolean if a field has been set.

### GetClientProfileQueueGuaranteed1MinMsgBurst

`func (o *DmrClusterLink) GetClientProfileQueueGuaranteed1MinMsgBurst() int32`

GetClientProfileQueueGuaranteed1MinMsgBurst returns the ClientProfileQueueGuaranteed1MinMsgBurst field if non-nil, zero value otherwise.

### GetClientProfileQueueGuaranteed1MinMsgBurstOk

`func (o *DmrClusterLink) GetClientProfileQueueGuaranteed1MinMsgBurstOk() (*int32, bool)`

GetClientProfileQueueGuaranteed1MinMsgBurstOk returns a tuple with the ClientProfileQueueGuaranteed1MinMsgBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileQueueGuaranteed1MinMsgBurst

`func (o *DmrClusterLink) SetClientProfileQueueGuaranteed1MinMsgBurst(v int32)`

SetClientProfileQueueGuaranteed1MinMsgBurst sets ClientProfileQueueGuaranteed1MinMsgBurst field to given value.

### HasClientProfileQueueGuaranteed1MinMsgBurst

`func (o *DmrClusterLink) HasClientProfileQueueGuaranteed1MinMsgBurst() bool`

HasClientProfileQueueGuaranteed1MinMsgBurst returns a boolean if a field has been set.

### GetClientProfileTcpCongestionWindowSize

`func (o *DmrClusterLink) GetClientProfileTcpCongestionWindowSize() int64`

GetClientProfileTcpCongestionWindowSize returns the ClientProfileTcpCongestionWindowSize field if non-nil, zero value otherwise.

### GetClientProfileTcpCongestionWindowSizeOk

`func (o *DmrClusterLink) GetClientProfileTcpCongestionWindowSizeOk() (*int64, bool)`

GetClientProfileTcpCongestionWindowSizeOk returns a tuple with the ClientProfileTcpCongestionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileTcpCongestionWindowSize

`func (o *DmrClusterLink) SetClientProfileTcpCongestionWindowSize(v int64)`

SetClientProfileTcpCongestionWindowSize sets ClientProfileTcpCongestionWindowSize field to given value.

### HasClientProfileTcpCongestionWindowSize

`func (o *DmrClusterLink) HasClientProfileTcpCongestionWindowSize() bool`

HasClientProfileTcpCongestionWindowSize returns a boolean if a field has been set.

### GetClientProfileTcpKeepaliveCount

`func (o *DmrClusterLink) GetClientProfileTcpKeepaliveCount() int64`

GetClientProfileTcpKeepaliveCount returns the ClientProfileTcpKeepaliveCount field if non-nil, zero value otherwise.

### GetClientProfileTcpKeepaliveCountOk

`func (o *DmrClusterLink) GetClientProfileTcpKeepaliveCountOk() (*int64, bool)`

GetClientProfileTcpKeepaliveCountOk returns a tuple with the ClientProfileTcpKeepaliveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileTcpKeepaliveCount

`func (o *DmrClusterLink) SetClientProfileTcpKeepaliveCount(v int64)`

SetClientProfileTcpKeepaliveCount sets ClientProfileTcpKeepaliveCount field to given value.

### HasClientProfileTcpKeepaliveCount

`func (o *DmrClusterLink) HasClientProfileTcpKeepaliveCount() bool`

HasClientProfileTcpKeepaliveCount returns a boolean if a field has been set.

### GetClientProfileTcpKeepaliveIdleTime

`func (o *DmrClusterLink) GetClientProfileTcpKeepaliveIdleTime() int64`

GetClientProfileTcpKeepaliveIdleTime returns the ClientProfileTcpKeepaliveIdleTime field if non-nil, zero value otherwise.

### GetClientProfileTcpKeepaliveIdleTimeOk

`func (o *DmrClusterLink) GetClientProfileTcpKeepaliveIdleTimeOk() (*int64, bool)`

GetClientProfileTcpKeepaliveIdleTimeOk returns a tuple with the ClientProfileTcpKeepaliveIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileTcpKeepaliveIdleTime

`func (o *DmrClusterLink) SetClientProfileTcpKeepaliveIdleTime(v int64)`

SetClientProfileTcpKeepaliveIdleTime sets ClientProfileTcpKeepaliveIdleTime field to given value.

### HasClientProfileTcpKeepaliveIdleTime

`func (o *DmrClusterLink) HasClientProfileTcpKeepaliveIdleTime() bool`

HasClientProfileTcpKeepaliveIdleTime returns a boolean if a field has been set.

### GetClientProfileTcpKeepaliveInterval

`func (o *DmrClusterLink) GetClientProfileTcpKeepaliveInterval() int64`

GetClientProfileTcpKeepaliveInterval returns the ClientProfileTcpKeepaliveInterval field if non-nil, zero value otherwise.

### GetClientProfileTcpKeepaliveIntervalOk

`func (o *DmrClusterLink) GetClientProfileTcpKeepaliveIntervalOk() (*int64, bool)`

GetClientProfileTcpKeepaliveIntervalOk returns a tuple with the ClientProfileTcpKeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileTcpKeepaliveInterval

`func (o *DmrClusterLink) SetClientProfileTcpKeepaliveInterval(v int64)`

SetClientProfileTcpKeepaliveInterval sets ClientProfileTcpKeepaliveInterval field to given value.

### HasClientProfileTcpKeepaliveInterval

`func (o *DmrClusterLink) HasClientProfileTcpKeepaliveInterval() bool`

HasClientProfileTcpKeepaliveInterval returns a boolean if a field has been set.

### GetClientProfileTcpMaxSegmentSize

`func (o *DmrClusterLink) GetClientProfileTcpMaxSegmentSize() int64`

GetClientProfileTcpMaxSegmentSize returns the ClientProfileTcpMaxSegmentSize field if non-nil, zero value otherwise.

### GetClientProfileTcpMaxSegmentSizeOk

`func (o *DmrClusterLink) GetClientProfileTcpMaxSegmentSizeOk() (*int64, bool)`

GetClientProfileTcpMaxSegmentSizeOk returns a tuple with the ClientProfileTcpMaxSegmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileTcpMaxSegmentSize

`func (o *DmrClusterLink) SetClientProfileTcpMaxSegmentSize(v int64)`

SetClientProfileTcpMaxSegmentSize sets ClientProfileTcpMaxSegmentSize field to given value.

### HasClientProfileTcpMaxSegmentSize

`func (o *DmrClusterLink) HasClientProfileTcpMaxSegmentSize() bool`

HasClientProfileTcpMaxSegmentSize returns a boolean if a field has been set.

### GetClientProfileTcpMaxWindowSize

`func (o *DmrClusterLink) GetClientProfileTcpMaxWindowSize() int64`

GetClientProfileTcpMaxWindowSize returns the ClientProfileTcpMaxWindowSize field if non-nil, zero value otherwise.

### GetClientProfileTcpMaxWindowSizeOk

`func (o *DmrClusterLink) GetClientProfileTcpMaxWindowSizeOk() (*int64, bool)`

GetClientProfileTcpMaxWindowSizeOk returns a tuple with the ClientProfileTcpMaxWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileTcpMaxWindowSize

`func (o *DmrClusterLink) SetClientProfileTcpMaxWindowSize(v int64)`

SetClientProfileTcpMaxWindowSize sets ClientProfileTcpMaxWindowSize field to given value.

### HasClientProfileTcpMaxWindowSize

`func (o *DmrClusterLink) HasClientProfileTcpMaxWindowSize() bool`

HasClientProfileTcpMaxWindowSize returns a boolean if a field has been set.

### GetDmrClusterName

`func (o *DmrClusterLink) GetDmrClusterName() string`

GetDmrClusterName returns the DmrClusterName field if non-nil, zero value otherwise.

### GetDmrClusterNameOk

`func (o *DmrClusterLink) GetDmrClusterNameOk() (*string, bool)`

GetDmrClusterNameOk returns a tuple with the DmrClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrClusterName

`func (o *DmrClusterLink) SetDmrClusterName(v string)`

SetDmrClusterName sets DmrClusterName field to given value.

### HasDmrClusterName

`func (o *DmrClusterLink) HasDmrClusterName() bool`

HasDmrClusterName returns a boolean if a field has been set.

### GetEgressFlowWindowSize

`func (o *DmrClusterLink) GetEgressFlowWindowSize() int64`

GetEgressFlowWindowSize returns the EgressFlowWindowSize field if non-nil, zero value otherwise.

### GetEgressFlowWindowSizeOk

`func (o *DmrClusterLink) GetEgressFlowWindowSizeOk() (*int64, bool)`

GetEgressFlowWindowSizeOk returns a tuple with the EgressFlowWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressFlowWindowSize

`func (o *DmrClusterLink) SetEgressFlowWindowSize(v int64)`

SetEgressFlowWindowSize sets EgressFlowWindowSize field to given value.

### HasEgressFlowWindowSize

`func (o *DmrClusterLink) HasEgressFlowWindowSize() bool`

HasEgressFlowWindowSize returns a boolean if a field has been set.

### GetEnabled

`func (o *DmrClusterLink) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DmrClusterLink) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DmrClusterLink) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DmrClusterLink) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInitiator

`func (o *DmrClusterLink) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *DmrClusterLink) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *DmrClusterLink) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *DmrClusterLink) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetQueueDeadMsgQueue

`func (o *DmrClusterLink) GetQueueDeadMsgQueue() string`

GetQueueDeadMsgQueue returns the QueueDeadMsgQueue field if non-nil, zero value otherwise.

### GetQueueDeadMsgQueueOk

`func (o *DmrClusterLink) GetQueueDeadMsgQueueOk() (*string, bool)`

GetQueueDeadMsgQueueOk returns a tuple with the QueueDeadMsgQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDeadMsgQueue

`func (o *DmrClusterLink) SetQueueDeadMsgQueue(v string)`

SetQueueDeadMsgQueue sets QueueDeadMsgQueue field to given value.

### HasQueueDeadMsgQueue

`func (o *DmrClusterLink) HasQueueDeadMsgQueue() bool`

HasQueueDeadMsgQueue returns a boolean if a field has been set.

### GetQueueEventSpoolUsageThreshold

`func (o *DmrClusterLink) GetQueueEventSpoolUsageThreshold() EventThreshold`

GetQueueEventSpoolUsageThreshold returns the QueueEventSpoolUsageThreshold field if non-nil, zero value otherwise.

### GetQueueEventSpoolUsageThresholdOk

`func (o *DmrClusterLink) GetQueueEventSpoolUsageThresholdOk() (*EventThreshold, bool)`

GetQueueEventSpoolUsageThresholdOk returns a tuple with the QueueEventSpoolUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueEventSpoolUsageThreshold

`func (o *DmrClusterLink) SetQueueEventSpoolUsageThreshold(v EventThreshold)`

SetQueueEventSpoolUsageThreshold sets QueueEventSpoolUsageThreshold field to given value.

### HasQueueEventSpoolUsageThreshold

`func (o *DmrClusterLink) HasQueueEventSpoolUsageThreshold() bool`

HasQueueEventSpoolUsageThreshold returns a boolean if a field has been set.

### GetQueueMaxDeliveredUnackedMsgsPerFlow

`func (o *DmrClusterLink) GetQueueMaxDeliveredUnackedMsgsPerFlow() int64`

GetQueueMaxDeliveredUnackedMsgsPerFlow returns the QueueMaxDeliveredUnackedMsgsPerFlow field if non-nil, zero value otherwise.

### GetQueueMaxDeliveredUnackedMsgsPerFlowOk

`func (o *DmrClusterLink) GetQueueMaxDeliveredUnackedMsgsPerFlowOk() (*int64, bool)`

GetQueueMaxDeliveredUnackedMsgsPerFlowOk returns a tuple with the QueueMaxDeliveredUnackedMsgsPerFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxDeliveredUnackedMsgsPerFlow

`func (o *DmrClusterLink) SetQueueMaxDeliveredUnackedMsgsPerFlow(v int64)`

SetQueueMaxDeliveredUnackedMsgsPerFlow sets QueueMaxDeliveredUnackedMsgsPerFlow field to given value.

### HasQueueMaxDeliveredUnackedMsgsPerFlow

`func (o *DmrClusterLink) HasQueueMaxDeliveredUnackedMsgsPerFlow() bool`

HasQueueMaxDeliveredUnackedMsgsPerFlow returns a boolean if a field has been set.

### GetQueueMaxMsgSpoolUsage

`func (o *DmrClusterLink) GetQueueMaxMsgSpoolUsage() int64`

GetQueueMaxMsgSpoolUsage returns the QueueMaxMsgSpoolUsage field if non-nil, zero value otherwise.

### GetQueueMaxMsgSpoolUsageOk

`func (o *DmrClusterLink) GetQueueMaxMsgSpoolUsageOk() (*int64, bool)`

GetQueueMaxMsgSpoolUsageOk returns a tuple with the QueueMaxMsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxMsgSpoolUsage

`func (o *DmrClusterLink) SetQueueMaxMsgSpoolUsage(v int64)`

SetQueueMaxMsgSpoolUsage sets QueueMaxMsgSpoolUsage field to given value.

### HasQueueMaxMsgSpoolUsage

`func (o *DmrClusterLink) HasQueueMaxMsgSpoolUsage() bool`

HasQueueMaxMsgSpoolUsage returns a boolean if a field has been set.

### GetQueueMaxRedeliveryCount

`func (o *DmrClusterLink) GetQueueMaxRedeliveryCount() int64`

GetQueueMaxRedeliveryCount returns the QueueMaxRedeliveryCount field if non-nil, zero value otherwise.

### GetQueueMaxRedeliveryCountOk

`func (o *DmrClusterLink) GetQueueMaxRedeliveryCountOk() (*int64, bool)`

GetQueueMaxRedeliveryCountOk returns a tuple with the QueueMaxRedeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxRedeliveryCount

`func (o *DmrClusterLink) SetQueueMaxRedeliveryCount(v int64)`

SetQueueMaxRedeliveryCount sets QueueMaxRedeliveryCount field to given value.

### HasQueueMaxRedeliveryCount

`func (o *DmrClusterLink) HasQueueMaxRedeliveryCount() bool`

HasQueueMaxRedeliveryCount returns a boolean if a field has been set.

### GetQueueMaxTtl

`func (o *DmrClusterLink) GetQueueMaxTtl() int64`

GetQueueMaxTtl returns the QueueMaxTtl field if non-nil, zero value otherwise.

### GetQueueMaxTtlOk

`func (o *DmrClusterLink) GetQueueMaxTtlOk() (*int64, bool)`

GetQueueMaxTtlOk returns a tuple with the QueueMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxTtl

`func (o *DmrClusterLink) SetQueueMaxTtl(v int64)`

SetQueueMaxTtl sets QueueMaxTtl field to given value.

### HasQueueMaxTtl

`func (o *DmrClusterLink) HasQueueMaxTtl() bool`

HasQueueMaxTtl returns a boolean if a field has been set.

### GetQueueRejectMsgToSenderOnDiscardBehavior

`func (o *DmrClusterLink) GetQueueRejectMsgToSenderOnDiscardBehavior() string`

GetQueueRejectMsgToSenderOnDiscardBehavior returns the QueueRejectMsgToSenderOnDiscardBehavior field if non-nil, zero value otherwise.

### GetQueueRejectMsgToSenderOnDiscardBehaviorOk

`func (o *DmrClusterLink) GetQueueRejectMsgToSenderOnDiscardBehaviorOk() (*string, bool)`

GetQueueRejectMsgToSenderOnDiscardBehaviorOk returns a tuple with the QueueRejectMsgToSenderOnDiscardBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueRejectMsgToSenderOnDiscardBehavior

`func (o *DmrClusterLink) SetQueueRejectMsgToSenderOnDiscardBehavior(v string)`

SetQueueRejectMsgToSenderOnDiscardBehavior sets QueueRejectMsgToSenderOnDiscardBehavior field to given value.

### HasQueueRejectMsgToSenderOnDiscardBehavior

`func (o *DmrClusterLink) HasQueueRejectMsgToSenderOnDiscardBehavior() bool`

HasQueueRejectMsgToSenderOnDiscardBehavior returns a boolean if a field has been set.

### GetQueueRespectTtlEnabled

`func (o *DmrClusterLink) GetQueueRespectTtlEnabled() bool`

GetQueueRespectTtlEnabled returns the QueueRespectTtlEnabled field if non-nil, zero value otherwise.

### GetQueueRespectTtlEnabledOk

`func (o *DmrClusterLink) GetQueueRespectTtlEnabledOk() (*bool, bool)`

GetQueueRespectTtlEnabledOk returns a tuple with the QueueRespectTtlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueRespectTtlEnabled

`func (o *DmrClusterLink) SetQueueRespectTtlEnabled(v bool)`

SetQueueRespectTtlEnabled sets QueueRespectTtlEnabled field to given value.

### HasQueueRespectTtlEnabled

`func (o *DmrClusterLink) HasQueueRespectTtlEnabled() bool`

HasQueueRespectTtlEnabled returns a boolean if a field has been set.

### GetRemoteNodeName

`func (o *DmrClusterLink) GetRemoteNodeName() string`

GetRemoteNodeName returns the RemoteNodeName field if non-nil, zero value otherwise.

### GetRemoteNodeNameOk

`func (o *DmrClusterLink) GetRemoteNodeNameOk() (*string, bool)`

GetRemoteNodeNameOk returns a tuple with the RemoteNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNodeName

`func (o *DmrClusterLink) SetRemoteNodeName(v string)`

SetRemoteNodeName sets RemoteNodeName field to given value.

### HasRemoteNodeName

`func (o *DmrClusterLink) HasRemoteNodeName() bool`

HasRemoteNodeName returns a boolean if a field has been set.

### GetSpan

`func (o *DmrClusterLink) GetSpan() string`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *DmrClusterLink) GetSpanOk() (*string, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *DmrClusterLink) SetSpan(v string)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *DmrClusterLink) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetTransportCompressedEnabled

`func (o *DmrClusterLink) GetTransportCompressedEnabled() bool`

GetTransportCompressedEnabled returns the TransportCompressedEnabled field if non-nil, zero value otherwise.

### GetTransportCompressedEnabledOk

`func (o *DmrClusterLink) GetTransportCompressedEnabledOk() (*bool, bool)`

GetTransportCompressedEnabledOk returns a tuple with the TransportCompressedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCompressedEnabled

`func (o *DmrClusterLink) SetTransportCompressedEnabled(v bool)`

SetTransportCompressedEnabled sets TransportCompressedEnabled field to given value.

### HasTransportCompressedEnabled

`func (o *DmrClusterLink) HasTransportCompressedEnabled() bool`

HasTransportCompressedEnabled returns a boolean if a field has been set.

### GetTransportTlsEnabled

`func (o *DmrClusterLink) GetTransportTlsEnabled() bool`

GetTransportTlsEnabled returns the TransportTlsEnabled field if non-nil, zero value otherwise.

### GetTransportTlsEnabledOk

`func (o *DmrClusterLink) GetTransportTlsEnabledOk() (*bool, bool)`

GetTransportTlsEnabledOk returns a tuple with the TransportTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportTlsEnabled

`func (o *DmrClusterLink) SetTransportTlsEnabled(v bool)`

SetTransportTlsEnabled sets TransportTlsEnabled field to given value.

### HasTransportTlsEnabled

`func (o *DmrClusterLink) HasTransportTlsEnabled() bool`

HasTransportTlsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


