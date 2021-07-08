# MsgVpnJndiConnectionFactory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDuplicateClientIdEnabled** | **bool** | Indicates whether new JMS connections can use the same Client identifier (ID) as an existing connection. | [optional] [default to null]
**ClientDescription** | **string** | The description of the Client. | [optional] [default to null]
**ClientId** | **string** | The Client identifier (ID). If not specified, a unique value for it will be generated. | [optional] [default to null]
**ConnectionFactoryName** | **string** | The name of the JMS Connection Factory. | [optional] [default to null]
**DtoReceiveOverrideEnabled** | **bool** | Indicates whether overriding by the Subscriber (Consumer) of the deliver-to-one (DTO) property on messages is enabled. When enabled, the Subscriber can receive all DTO tagged messages. | [optional] [default to null]
**DtoReceiveSubscriberLocalPriority** | **int32** | The priority for receiving deliver-to-one (DTO) messages by the Subscriber (Consumer) if the messages are published on the local broker that the Subscriber is directly connected to. | [optional] [default to null]
**DtoReceiveSubscriberNetworkPriority** | **int32** | The priority for receiving deliver-to-one (DTO) messages by the Subscriber (Consumer) if the messages are published on a remote broker. | [optional] [default to null]
**DtoSendEnabled** | **bool** | Indicates whether the deliver-to-one (DTO) property is enabled on messages sent by the Publisher (Producer). | [optional] [default to null]
**DynamicEndpointCreateDurableEnabled** | **bool** | Indicates whether a durable endpoint will be dynamically created on the broker when the client calls \&quot;Session.createDurableSubscriber()\&quot; or \&quot;Session.createQueue()\&quot;. The created endpoint respects the message time-to-live (TTL) according to the \&quot;dynamicEndpointRespectTtlEnabled\&quot; property. | [optional] [default to null]
**DynamicEndpointRespectTtlEnabled** | **bool** | Indicates whether dynamically created durable and non-durable endpoints respect the message time-to-live (TTL) property. | [optional] [default to null]
**GuaranteedReceiveAckTimeout** | **int32** | The timeout for sending the acknowledgement (ACK) for guaranteed messages received by the Subscriber (Consumer), in milliseconds. | [optional] [default to null]
**GuaranteedReceiveReconnectRetryCount** | **int32** | The maximum number of attempts to reconnect to the host or list of hosts after the guaranteed  messaging connection has been lost. The value \&quot;-1\&quot; means to retry forever. Available since 2.14. | [optional] [default to null]
**GuaranteedReceiveReconnectRetryWait** | **int32** | The amount of time to wait before making another attempt to connect or reconnect to the host after the guaranteed messaging connection has been lost, in milliseconds. Available since 2.14. | [optional] [default to null]
**GuaranteedReceiveWindowSize** | **int32** | The size of the window for guaranteed messages received by the Subscriber (Consumer), in messages. | [optional] [default to null]
**GuaranteedReceiveWindowSizeAckThreshold** | **int32** | The threshold for sending the acknowledgement (ACK) for guaranteed messages received by the Subscriber (Consumer) as a percentage of &#x60;guaranteedReceiveWindowSize&#x60;. | [optional] [default to null]
**GuaranteedSendAckTimeout** | **int32** | The timeout for receiving the acknowledgement (ACK) for guaranteed messages sent by the Publisher (Producer), in milliseconds. | [optional] [default to null]
**GuaranteedSendWindowSize** | **int32** | The size of the window for non-persistent guaranteed messages sent by the Publisher (Producer), in messages. For persistent messages the window size is fixed at 1. | [optional] [default to null]
**MessagingDefaultDeliveryMode** | **string** | The default delivery mode for messages sent by the Publisher (Producer). The allowed values and their meaning are:  &lt;pre&gt; \&quot;persistent\&quot; - The broker spools messages (persists in the Message Spool) as part of the send operation. \&quot;non-persistent\&quot; - The broker does not spool messages (does not persist in the Message Spool) as part of the send operation. &lt;/pre&gt;  | [optional] [default to null]
**MessagingDefaultDmqEligibleEnabled** | **bool** | Indicates whether messages sent by the Publisher (Producer) are Dead Message Queue (DMQ) eligible by default. | [optional] [default to null]
**MessagingDefaultElidingEligibleEnabled** | **bool** | Indicates whether messages sent by the Publisher (Producer) are Eliding eligible by default. | [optional] [default to null]
**MessagingJmsxUserIdEnabled** | **bool** | Indicates whether to include (add or replace) the JMSXUserID property in messages sent by the Publisher (Producer). | [optional] [default to null]
**MessagingTextInXmlPayloadEnabled** | **bool** | Indicates whether encoding of JMS text messages in Publisher (Producer) messages is as XML payload. When disabled, JMS text messages are encoded as a binary attachment. | [optional] [default to null]
**MsgVpnName** | **string** | The name of the Message VPN. | [optional] [default to null]
**TransportCompressionLevel** | **int32** | The ZLIB compression level for the connection to the broker. The value \&quot;0\&quot; means no compression, and the value \&quot;-1\&quot; means the compression level is specified in the JNDI Properties file. | [optional] [default to null]
**TransportConnectRetryCount** | **int32** | The maximum number of retry attempts to establish an initial connection to the host or list of hosts. The value \&quot;0\&quot; means a single attempt (no retries), and the value \&quot;-1\&quot; means to retry forever. | [optional] [default to null]
**TransportConnectRetryPerHostCount** | **int32** | The maximum number of retry attempts to establish an initial connection to each host on the list of hosts. The value \&quot;0\&quot; means a single attempt (no retries), and the value \&quot;-1\&quot; means to retry forever. | [optional] [default to null]
**TransportConnectTimeout** | **int32** | The timeout for establishing an initial connection to the broker, in milliseconds. | [optional] [default to null]
**TransportDirectTransportEnabled** | **bool** | Indicates whether usage of the Direct Transport mode for sending non-persistent messages is enabled. When disabled, the Guaranteed Transport mode is used. | [optional] [default to null]
**TransportKeepaliveCount** | **int32** | The maximum number of consecutive application-level keepalive messages sent without the broker response before the connection to the broker is closed. | [optional] [default to null]
**TransportKeepaliveEnabled** | **bool** | Indicates whether application-level keepalive messages are used to maintain a connection with the Router. | [optional] [default to null]
**TransportKeepaliveInterval** | **int32** | The interval between application-level keepalive messages, in milliseconds. | [optional] [default to null]
**TransportMsgCallbackOnIoThreadEnabled** | **bool** | Indicates whether delivery of asynchronous messages is done directly from the I/O thread. | [optional] [default to null]
**TransportOptimizeDirectEnabled** | **bool** | Indicates whether optimization for the Direct Transport delivery mode is enabled. If enabled, the client application is limited to one Publisher (Producer) and one non-durable Subscriber (Consumer). | [optional] [default to null]
**TransportPort** | **int32** | The connection port number on the broker for SMF clients. The value \&quot;-1\&quot; means the port is specified in the JNDI Properties file. | [optional] [default to null]
**TransportReadTimeout** | **int32** | The timeout for reading a reply from the broker, in milliseconds. | [optional] [default to null]
**TransportReceiveBufferSize** | **int32** | The size of the receive socket buffer, in bytes. It corresponds to the SO_RCVBUF socket option. | [optional] [default to null]
**TransportReconnectRetryCount** | **int32** | The maximum number of attempts to reconnect to the host or list of hosts after the connection has been lost. The value \&quot;-1\&quot; means to retry forever. | [optional] [default to null]
**TransportReconnectRetryWait** | **int32** | The amount of time before making another attempt to connect or reconnect to the host after the connection has been lost, in milliseconds. | [optional] [default to null]
**TransportSendBufferSize** | **int32** | The size of the send socket buffer, in bytes. It corresponds to the SO_SNDBUF socket option. | [optional] [default to null]
**TransportTcpNoDelayEnabled** | **bool** | Indicates whether the TCP_NODELAY option is enabled, which disables Nagle&#x27;s algorithm for TCP/IP congestion control (RFC 896). | [optional] [default to null]
**XaEnabled** | **bool** | Indicates whether this is an XA Connection Factory. When enabled, the Connection Factory can be cast to \&quot;XAConnectionFactory\&quot;, \&quot;XAQueueConnectionFactory\&quot; or \&quot;XATopicConnectionFactory\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
