# DmrClusterLink

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationScheme** | **string** | The authentication scheme to be used by the Link which initiates connections to the remote node. The allowed values and their meaning are:  &lt;pre&gt; \&quot;basic\&quot; - Basic Authentication Scheme (via username and password). \&quot;client-certificate\&quot; - Client Certificate Authentication Scheme (via certificate file or content). &lt;/pre&gt;  | [optional] [default to null]
**ClientProfileName** | **string** | The name of the Client Profile used by the Link. | [optional] [default to null]
**ClientProfileQueueControl1MaxDepth** | **int32** | The maximum depth of the \&quot;Control 1\&quot; (C-1) priority queue, in work units. Each work unit is 2048 bytes of message data. | [optional] [default to null]
**ClientProfileQueueControl1MinMsgBurst** | **int32** | The number of messages that are always allowed entry into the \&quot;Control 1\&quot; (C-1) priority queue, regardless of the &#x60;clientProfileQueueControl1MaxDepth&#x60; value. | [optional] [default to null]
**ClientProfileQueueDirect1MaxDepth** | **int32** | The maximum depth of the \&quot;Direct 1\&quot; (D-1) priority queue, in work units. Each work unit is 2048 bytes of message data. | [optional] [default to null]
**ClientProfileQueueDirect1MinMsgBurst** | **int32** | The number of messages that are always allowed entry into the \&quot;Direct 1\&quot; (D-1) priority queue, regardless of the &#x60;clientProfileQueueDirect1MaxDepth&#x60; value. | [optional] [default to null]
**ClientProfileQueueDirect2MaxDepth** | **int32** | The maximum depth of the \&quot;Direct 2\&quot; (D-2) priority queue, in work units. Each work unit is 2048 bytes of message data. | [optional] [default to null]
**ClientProfileQueueDirect2MinMsgBurst** | **int32** | The number of messages that are always allowed entry into the \&quot;Direct 2\&quot; (D-2) priority queue, regardless of the &#x60;clientProfileQueueDirect2MaxDepth&#x60; value. | [optional] [default to null]
**ClientProfileQueueDirect3MaxDepth** | **int32** | The maximum depth of the \&quot;Direct 3\&quot; (D-3) priority queue, in work units. Each work unit is 2048 bytes of message data. | [optional] [default to null]
**ClientProfileQueueDirect3MinMsgBurst** | **int32** | The number of messages that are always allowed entry into the \&quot;Direct 3\&quot; (D-3) priority queue, regardless of the &#x60;clientProfileQueueDirect3MaxDepth&#x60; value. | [optional] [default to null]
**ClientProfileQueueGuaranteed1MaxDepth** | **int32** | The maximum depth of the \&quot;Guaranteed 1\&quot; (G-1) priority queue, in work units. Each work unit is 2048 bytes of message data. | [optional] [default to null]
**ClientProfileQueueGuaranteed1MinMsgBurst** | **int32** | The number of messages that are always allowed entry into the \&quot;Guaranteed 1\&quot; (G-3) priority queue, regardless of the &#x60;clientProfileQueueGuaranteed1MaxDepth&#x60; value. | [optional] [default to null]
**ClientProfileTcpCongestionWindowSize** | **int64** | The TCP initial congestion window size, in multiples of the TCP Maximum Segment Size (MSS). Changing the value from its default of 2 results in non-compliance with RFC 2581. Contact Solace Support before changing this value. | [optional] [default to null]
**ClientProfileTcpKeepaliveCount** | **int64** | The number of TCP keepalive retransmissions to be carried out before declaring that the remote end is not available. | [optional] [default to null]
**ClientProfileTcpKeepaliveIdleTime** | **int64** | The amount of time a connection must remain idle before TCP begins sending keepalive probes, in seconds. | [optional] [default to null]
**ClientProfileTcpKeepaliveInterval** | **int64** | The amount of time between TCP keepalive retransmissions when no acknowledgement is received, in seconds. | [optional] [default to null]
**ClientProfileTcpMaxSegmentSize** | **int64** | The TCP maximum segment size, in bytes. Changes are applied to all existing connections. | [optional] [default to null]
**ClientProfileTcpMaxWindowSize** | **int64** | The TCP maximum window size, in kilobytes. Changes are applied to all existing connections. | [optional] [default to null]
**DmrClusterName** | **string** | The name of the Cluster. | [optional] [default to null]
**EgressFlowWindowSize** | **int64** | The number of outstanding guaranteed messages that can be sent over the Link before acknowledgement is received by the sender. | [optional] [default to null]
**Enabled** | **bool** | Indicates whether the Link is enabled. When disabled, subscription sets of this and the remote node are not kept up-to-date, and messages are not exchanged with the remote node. Published guaranteed messages will be queued up for future delivery based on current subscription sets. | [optional] [default to null]
**FailureReason** | **string** | The failure reason for the Link being down. | [optional] [default to null]
**Initiator** | **string** | The initiator of the Link&#x27;s TCP connections. The allowed values and their meaning are:  &lt;pre&gt; \&quot;lexical\&quot; - The \&quot;higher\&quot; node-name initiates. \&quot;local\&quot; - The local node initiates. \&quot;remote\&quot; - The remote node initiates. &lt;/pre&gt;  | [optional] [default to null]
**QueueDeadMsgQueue** | **string** | The name of the Dead Message Queue (DMQ) used by the Queue for discarded messages. | [optional] [default to null]
**QueueEventSpoolUsageThreshold** | [***EventThreshold**](EventThreshold.md) |  | [optional] [default to null]
**QueueMaxDeliveredUnackedMsgsPerFlow** | **int64** | The maximum number of messages delivered but not acknowledged per flow for the Queue. | [optional] [default to null]
**QueueMaxMsgSpoolUsage** | **int64** | The maximum message spool usage by the Queue (quota), in megabytes (MB). | [optional] [default to null]
**QueueMaxRedeliveryCount** | **int64** | The maximum number of times the Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. | [optional] [default to null]
**QueueMaxTtl** | **int64** | The maximum time in seconds a message can stay in the Queue when &#x60;queueRespectTtlEnabled&#x60; is &#x60;true&#x60;. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the &#x60;queueMaxTtl&#x60; configured for the Queue, is exceeded. A value of 0 disables expiry. | [optional] [default to null]
**QueueRejectMsgToSenderOnDiscardBehavior** | **string** | Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always return a negative acknowledgment (NACK) to the sending client on message discard. \&quot;when-queue-enabled\&quot; - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Queue is enabled. \&quot;never\&quot; - Never return a negative acknowledgment (NACK) to the sending client on message discard. &lt;/pre&gt;  | [optional] [default to null]
**QueueRespectTtlEnabled** | **bool** | Indicates whether the the time-to-live (TTL) for messages in the Queue is respected. When enabled, expired messages are discarded or moved to the DMQ. | [optional] [default to null]
**RemoteClusterName** | **string** | The cluster name of the remote node. Available since 2.17. | [optional] [default to null]
**RemoteNodeName** | **string** | The name of the node at the remote end of the Link. | [optional] [default to null]
**Span** | **string** | The span of the Link, either internal or external. Internal Links connect nodes within the same Cluster. External Links connect nodes within different Clusters. The allowed values and their meaning are:  &lt;pre&gt; \&quot;internal\&quot; - Link to same cluster. \&quot;external\&quot; - Link to other cluster. &lt;/pre&gt;  | [optional] [default to null]
**TransportCompressedEnabled** | **bool** | Indicates whether compression is enabled on the Link. | [optional] [default to null]
**TransportTlsEnabled** | **bool** | Indicates whether encryption (TLS) is enabled on the Link. | [optional] [default to null]
**Up** | **bool** | Indicates whether the Link is operationally up. | [optional] [default to null]
**Uptime** | **int64** | The amount of time in seconds since the Link was up. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
