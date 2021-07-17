# MsgVpnJndiConnectionFactory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDuplicateClientIdEnabled** | Pointer to **bool** | Indicates whether new JMS connections can use the same Client identifier (ID) as an existing connection. | [optional] 
**ClientDescription** | Pointer to **string** | The description of the Client. | [optional] 
**ClientId** | Pointer to **string** | The Client identifier (ID). If not specified, a unique value for it will be generated. | [optional] 
**ConnectionFactoryName** | Pointer to **string** | The name of the JMS Connection Factory. | [optional] 
**DtoReceiveOverrideEnabled** | Pointer to **bool** | Indicates whether overriding by the Subscriber (Consumer) of the deliver-to-one (DTO) property on messages is enabled. When enabled, the Subscriber can receive all DTO tagged messages. | [optional] 
**DtoReceiveSubscriberLocalPriority** | Pointer to **int32** | The priority for receiving deliver-to-one (DTO) messages by the Subscriber (Consumer) if the messages are published on the local broker that the Subscriber is directly connected to. | [optional] 
**DtoReceiveSubscriberNetworkPriority** | Pointer to **int32** | The priority for receiving deliver-to-one (DTO) messages by the Subscriber (Consumer) if the messages are published on a remote broker. | [optional] 
**DtoSendEnabled** | Pointer to **bool** | Indicates whether the deliver-to-one (DTO) property is enabled on messages sent by the Publisher (Producer). | [optional] 
**DynamicEndpointCreateDurableEnabled** | Pointer to **bool** | Indicates whether a durable endpoint will be dynamically created on the broker when the client calls \&quot;Session.createDurableSubscriber()\&quot; or \&quot;Session.createQueue()\&quot;. The created endpoint respects the message time-to-live (TTL) according to the \&quot;dynamicEndpointRespectTtlEnabled\&quot; property. | [optional] 
**DynamicEndpointRespectTtlEnabled** | Pointer to **bool** | Indicates whether dynamically created durable and non-durable endpoints respect the message time-to-live (TTL) property. | [optional] 
**GuaranteedReceiveAckTimeout** | Pointer to **int32** | The timeout for sending the acknowledgement (ACK) for guaranteed messages received by the Subscriber (Consumer), in milliseconds. | [optional] 
**GuaranteedReceiveReconnectRetryCount** | Pointer to **int32** | The maximum number of attempts to reconnect to the host or list of hosts after the guaranteed  messaging connection has been lost. The value \&quot;-1\&quot; means to retry forever. Available since 2.14. | [optional] 
**GuaranteedReceiveReconnectRetryWait** | Pointer to **int32** | The amount of time to wait before making another attempt to connect or reconnect to the host after the guaranteed messaging connection has been lost, in milliseconds. Available since 2.14. | [optional] 
**GuaranteedReceiveWindowSize** | Pointer to **int32** | The size of the window for guaranteed messages received by the Subscriber (Consumer), in messages. | [optional] 
**GuaranteedReceiveWindowSizeAckThreshold** | Pointer to **int32** | The threshold for sending the acknowledgement (ACK) for guaranteed messages received by the Subscriber (Consumer) as a percentage of &#x60;guaranteedReceiveWindowSize&#x60;. | [optional] 
**GuaranteedSendAckTimeout** | Pointer to **int32** | The timeout for receiving the acknowledgement (ACK) for guaranteed messages sent by the Publisher (Producer), in milliseconds. | [optional] 
**GuaranteedSendWindowSize** | Pointer to **int32** | The size of the window for non-persistent guaranteed messages sent by the Publisher (Producer), in messages. For persistent messages the window size is fixed at 1. | [optional] 
**MessagingDefaultDeliveryMode** | Pointer to **string** | The default delivery mode for messages sent by the Publisher (Producer). The allowed values and their meaning are:  &lt;pre&gt; \&quot;persistent\&quot; - The broker spools messages (persists in the Message Spool) as part of the send operation. \&quot;non-persistent\&quot; - The broker does not spool messages (does not persist in the Message Spool) as part of the send operation. &lt;/pre&gt;  | [optional] 
**MessagingDefaultDmqEligibleEnabled** | Pointer to **bool** | Indicates whether messages sent by the Publisher (Producer) are Dead Message Queue (DMQ) eligible by default. | [optional] 
**MessagingDefaultElidingEligibleEnabled** | Pointer to **bool** | Indicates whether messages sent by the Publisher (Producer) are Eliding eligible by default. | [optional] 
**MessagingJmsxUserIdEnabled** | Pointer to **bool** | Indicates whether to include (add or replace) the JMSXUserID property in messages sent by the Publisher (Producer). | [optional] 
**MessagingTextInXmlPayloadEnabled** | Pointer to **bool** | Indicates whether encoding of JMS text messages in Publisher (Producer) messages is as XML payload. When disabled, JMS text messages are encoded as a binary attachment. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**TransportCompressionLevel** | Pointer to **int32** | The ZLIB compression level for the connection to the broker. The value \&quot;0\&quot; means no compression, and the value \&quot;-1\&quot; means the compression level is specified in the JNDI Properties file. | [optional] 
**TransportConnectRetryCount** | Pointer to **int32** | The maximum number of retry attempts to establish an initial connection to the host or list of hosts. The value \&quot;0\&quot; means a single attempt (no retries), and the value \&quot;-1\&quot; means to retry forever. | [optional] 
**TransportConnectRetryPerHostCount** | Pointer to **int32** | The maximum number of retry attempts to establish an initial connection to each host on the list of hosts. The value \&quot;0\&quot; means a single attempt (no retries), and the value \&quot;-1\&quot; means to retry forever. | [optional] 
**TransportConnectTimeout** | Pointer to **int32** | The timeout for establishing an initial connection to the broker, in milliseconds. | [optional] 
**TransportDirectTransportEnabled** | Pointer to **bool** | Indicates whether usage of the Direct Transport mode for sending non-persistent messages is enabled. When disabled, the Guaranteed Transport mode is used. | [optional] 
**TransportKeepaliveCount** | Pointer to **int32** | The maximum number of consecutive application-level keepalive messages sent without the broker response before the connection to the broker is closed. | [optional] 
**TransportKeepaliveEnabled** | Pointer to **bool** | Indicates whether application-level keepalive messages are used to maintain a connection with the Router. | [optional] 
**TransportKeepaliveInterval** | Pointer to **int32** | The interval between application-level keepalive messages, in milliseconds. | [optional] 
**TransportMsgCallbackOnIoThreadEnabled** | Pointer to **bool** | Indicates whether delivery of asynchronous messages is done directly from the I/O thread. | [optional] 
**TransportOptimizeDirectEnabled** | Pointer to **bool** | Indicates whether optimization for the Direct Transport delivery mode is enabled. If enabled, the client application is limited to one Publisher (Producer) and one non-durable Subscriber (Consumer). | [optional] 
**TransportPort** | Pointer to **int32** | The connection port number on the broker for SMF clients. The value \&quot;-1\&quot; means the port is specified in the JNDI Properties file. | [optional] 
**TransportReadTimeout** | Pointer to **int32** | The timeout for reading a reply from the broker, in milliseconds. | [optional] 
**TransportReceiveBufferSize** | Pointer to **int32** | The size of the receive socket buffer, in bytes. It corresponds to the SO_RCVBUF socket option. | [optional] 
**TransportReconnectRetryCount** | Pointer to **int32** | The maximum number of attempts to reconnect to the host or list of hosts after the connection has been lost. The value \&quot;-1\&quot; means to retry forever. | [optional] 
**TransportReconnectRetryWait** | Pointer to **int32** | The amount of time before making another attempt to connect or reconnect to the host after the connection has been lost, in milliseconds. | [optional] 
**TransportSendBufferSize** | Pointer to **int32** | The size of the send socket buffer, in bytes. It corresponds to the SO_SNDBUF socket option. | [optional] 
**TransportTcpNoDelayEnabled** | Pointer to **bool** | Indicates whether the TCP_NODELAY option is enabled, which disables Nagle&#39;s algorithm for TCP/IP congestion control (RFC 896). | [optional] 
**XaEnabled** | Pointer to **bool** | Indicates whether this is an XA Connection Factory. When enabled, the Connection Factory can be cast to \&quot;XAConnectionFactory\&quot;, \&quot;XAQueueConnectionFactory\&quot; or \&quot;XATopicConnectionFactory\&quot;. | [optional] 

## Methods

### NewMsgVpnJndiConnectionFactory

`func NewMsgVpnJndiConnectionFactory() *MsgVpnJndiConnectionFactory`

NewMsgVpnJndiConnectionFactory instantiates a new MsgVpnJndiConnectionFactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnJndiConnectionFactoryWithDefaults

`func NewMsgVpnJndiConnectionFactoryWithDefaults() *MsgVpnJndiConnectionFactory`

NewMsgVpnJndiConnectionFactoryWithDefaults instantiates a new MsgVpnJndiConnectionFactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDuplicateClientIdEnabled

`func (o *MsgVpnJndiConnectionFactory) GetAllowDuplicateClientIdEnabled() bool`

GetAllowDuplicateClientIdEnabled returns the AllowDuplicateClientIdEnabled field if non-nil, zero value otherwise.

### GetAllowDuplicateClientIdEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetAllowDuplicateClientIdEnabledOk() (*bool, bool)`

GetAllowDuplicateClientIdEnabledOk returns a tuple with the AllowDuplicateClientIdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicateClientIdEnabled

`func (o *MsgVpnJndiConnectionFactory) SetAllowDuplicateClientIdEnabled(v bool)`

SetAllowDuplicateClientIdEnabled sets AllowDuplicateClientIdEnabled field to given value.

### HasAllowDuplicateClientIdEnabled

`func (o *MsgVpnJndiConnectionFactory) HasAllowDuplicateClientIdEnabled() bool`

HasAllowDuplicateClientIdEnabled returns a boolean if a field has been set.

### GetClientDescription

`func (o *MsgVpnJndiConnectionFactory) GetClientDescription() string`

GetClientDescription returns the ClientDescription field if non-nil, zero value otherwise.

### GetClientDescriptionOk

`func (o *MsgVpnJndiConnectionFactory) GetClientDescriptionOk() (*string, bool)`

GetClientDescriptionOk returns a tuple with the ClientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDescription

`func (o *MsgVpnJndiConnectionFactory) SetClientDescription(v string)`

SetClientDescription sets ClientDescription field to given value.

### HasClientDescription

`func (o *MsgVpnJndiConnectionFactory) HasClientDescription() bool`

HasClientDescription returns a boolean if a field has been set.

### GetClientId

`func (o *MsgVpnJndiConnectionFactory) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MsgVpnJndiConnectionFactory) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MsgVpnJndiConnectionFactory) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MsgVpnJndiConnectionFactory) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetConnectionFactoryName

`func (o *MsgVpnJndiConnectionFactory) GetConnectionFactoryName() string`

GetConnectionFactoryName returns the ConnectionFactoryName field if non-nil, zero value otherwise.

### GetConnectionFactoryNameOk

`func (o *MsgVpnJndiConnectionFactory) GetConnectionFactoryNameOk() (*string, bool)`

GetConnectionFactoryNameOk returns a tuple with the ConnectionFactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionFactoryName

`func (o *MsgVpnJndiConnectionFactory) SetConnectionFactoryName(v string)`

SetConnectionFactoryName sets ConnectionFactoryName field to given value.

### HasConnectionFactoryName

`func (o *MsgVpnJndiConnectionFactory) HasConnectionFactoryName() bool`

HasConnectionFactoryName returns a boolean if a field has been set.

### GetDtoReceiveOverrideEnabled

`func (o *MsgVpnJndiConnectionFactory) GetDtoReceiveOverrideEnabled() bool`

GetDtoReceiveOverrideEnabled returns the DtoReceiveOverrideEnabled field if non-nil, zero value otherwise.

### GetDtoReceiveOverrideEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetDtoReceiveOverrideEnabledOk() (*bool, bool)`

GetDtoReceiveOverrideEnabledOk returns a tuple with the DtoReceiveOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtoReceiveOverrideEnabled

`func (o *MsgVpnJndiConnectionFactory) SetDtoReceiveOverrideEnabled(v bool)`

SetDtoReceiveOverrideEnabled sets DtoReceiveOverrideEnabled field to given value.

### HasDtoReceiveOverrideEnabled

`func (o *MsgVpnJndiConnectionFactory) HasDtoReceiveOverrideEnabled() bool`

HasDtoReceiveOverrideEnabled returns a boolean if a field has been set.

### GetDtoReceiveSubscriberLocalPriority

`func (o *MsgVpnJndiConnectionFactory) GetDtoReceiveSubscriberLocalPriority() int32`

GetDtoReceiveSubscriberLocalPriority returns the DtoReceiveSubscriberLocalPriority field if non-nil, zero value otherwise.

### GetDtoReceiveSubscriberLocalPriorityOk

`func (o *MsgVpnJndiConnectionFactory) GetDtoReceiveSubscriberLocalPriorityOk() (*int32, bool)`

GetDtoReceiveSubscriberLocalPriorityOk returns a tuple with the DtoReceiveSubscriberLocalPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtoReceiveSubscriberLocalPriority

`func (o *MsgVpnJndiConnectionFactory) SetDtoReceiveSubscriberLocalPriority(v int32)`

SetDtoReceiveSubscriberLocalPriority sets DtoReceiveSubscriberLocalPriority field to given value.

### HasDtoReceiveSubscriberLocalPriority

`func (o *MsgVpnJndiConnectionFactory) HasDtoReceiveSubscriberLocalPriority() bool`

HasDtoReceiveSubscriberLocalPriority returns a boolean if a field has been set.

### GetDtoReceiveSubscriberNetworkPriority

`func (o *MsgVpnJndiConnectionFactory) GetDtoReceiveSubscriberNetworkPriority() int32`

GetDtoReceiveSubscriberNetworkPriority returns the DtoReceiveSubscriberNetworkPriority field if non-nil, zero value otherwise.

### GetDtoReceiveSubscriberNetworkPriorityOk

`func (o *MsgVpnJndiConnectionFactory) GetDtoReceiveSubscriberNetworkPriorityOk() (*int32, bool)`

GetDtoReceiveSubscriberNetworkPriorityOk returns a tuple with the DtoReceiveSubscriberNetworkPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtoReceiveSubscriberNetworkPriority

`func (o *MsgVpnJndiConnectionFactory) SetDtoReceiveSubscriberNetworkPriority(v int32)`

SetDtoReceiveSubscriberNetworkPriority sets DtoReceiveSubscriberNetworkPriority field to given value.

### HasDtoReceiveSubscriberNetworkPriority

`func (o *MsgVpnJndiConnectionFactory) HasDtoReceiveSubscriberNetworkPriority() bool`

HasDtoReceiveSubscriberNetworkPriority returns a boolean if a field has been set.

### GetDtoSendEnabled

`func (o *MsgVpnJndiConnectionFactory) GetDtoSendEnabled() bool`

GetDtoSendEnabled returns the DtoSendEnabled field if non-nil, zero value otherwise.

### GetDtoSendEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetDtoSendEnabledOk() (*bool, bool)`

GetDtoSendEnabledOk returns a tuple with the DtoSendEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtoSendEnabled

`func (o *MsgVpnJndiConnectionFactory) SetDtoSendEnabled(v bool)`

SetDtoSendEnabled sets DtoSendEnabled field to given value.

### HasDtoSendEnabled

`func (o *MsgVpnJndiConnectionFactory) HasDtoSendEnabled() bool`

HasDtoSendEnabled returns a boolean if a field has been set.

### GetDynamicEndpointCreateDurableEnabled

`func (o *MsgVpnJndiConnectionFactory) GetDynamicEndpointCreateDurableEnabled() bool`

GetDynamicEndpointCreateDurableEnabled returns the DynamicEndpointCreateDurableEnabled field if non-nil, zero value otherwise.

### GetDynamicEndpointCreateDurableEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetDynamicEndpointCreateDurableEnabledOk() (*bool, bool)`

GetDynamicEndpointCreateDurableEnabledOk returns a tuple with the DynamicEndpointCreateDurableEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicEndpointCreateDurableEnabled

`func (o *MsgVpnJndiConnectionFactory) SetDynamicEndpointCreateDurableEnabled(v bool)`

SetDynamicEndpointCreateDurableEnabled sets DynamicEndpointCreateDurableEnabled field to given value.

### HasDynamicEndpointCreateDurableEnabled

`func (o *MsgVpnJndiConnectionFactory) HasDynamicEndpointCreateDurableEnabled() bool`

HasDynamicEndpointCreateDurableEnabled returns a boolean if a field has been set.

### GetDynamicEndpointRespectTtlEnabled

`func (o *MsgVpnJndiConnectionFactory) GetDynamicEndpointRespectTtlEnabled() bool`

GetDynamicEndpointRespectTtlEnabled returns the DynamicEndpointRespectTtlEnabled field if non-nil, zero value otherwise.

### GetDynamicEndpointRespectTtlEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetDynamicEndpointRespectTtlEnabledOk() (*bool, bool)`

GetDynamicEndpointRespectTtlEnabledOk returns a tuple with the DynamicEndpointRespectTtlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicEndpointRespectTtlEnabled

`func (o *MsgVpnJndiConnectionFactory) SetDynamicEndpointRespectTtlEnabled(v bool)`

SetDynamicEndpointRespectTtlEnabled sets DynamicEndpointRespectTtlEnabled field to given value.

### HasDynamicEndpointRespectTtlEnabled

`func (o *MsgVpnJndiConnectionFactory) HasDynamicEndpointRespectTtlEnabled() bool`

HasDynamicEndpointRespectTtlEnabled returns a boolean if a field has been set.

### GetGuaranteedReceiveAckTimeout

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedReceiveAckTimeout() int32`

GetGuaranteedReceiveAckTimeout returns the GuaranteedReceiveAckTimeout field if non-nil, zero value otherwise.

### GetGuaranteedReceiveAckTimeoutOk

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedReceiveAckTimeoutOk() (*int32, bool)`

GetGuaranteedReceiveAckTimeoutOk returns a tuple with the GuaranteedReceiveAckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedReceiveAckTimeout

`func (o *MsgVpnJndiConnectionFactory) SetGuaranteedReceiveAckTimeout(v int32)`

SetGuaranteedReceiveAckTimeout sets GuaranteedReceiveAckTimeout field to given value.

### HasGuaranteedReceiveAckTimeout

`func (o *MsgVpnJndiConnectionFactory) HasGuaranteedReceiveAckTimeout() bool`

HasGuaranteedReceiveAckTimeout returns a boolean if a field has been set.

### GetGuaranteedReceiveReconnectRetryCount

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedReceiveReconnectRetryCount() int32`

GetGuaranteedReceiveReconnectRetryCount returns the GuaranteedReceiveReconnectRetryCount field if non-nil, zero value otherwise.

### GetGuaranteedReceiveReconnectRetryCountOk

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedReceiveReconnectRetryCountOk() (*int32, bool)`

GetGuaranteedReceiveReconnectRetryCountOk returns a tuple with the GuaranteedReceiveReconnectRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedReceiveReconnectRetryCount

`func (o *MsgVpnJndiConnectionFactory) SetGuaranteedReceiveReconnectRetryCount(v int32)`

SetGuaranteedReceiveReconnectRetryCount sets GuaranteedReceiveReconnectRetryCount field to given value.

### HasGuaranteedReceiveReconnectRetryCount

`func (o *MsgVpnJndiConnectionFactory) HasGuaranteedReceiveReconnectRetryCount() bool`

HasGuaranteedReceiveReconnectRetryCount returns a boolean if a field has been set.

### GetGuaranteedReceiveReconnectRetryWait

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedReceiveReconnectRetryWait() int32`

GetGuaranteedReceiveReconnectRetryWait returns the GuaranteedReceiveReconnectRetryWait field if non-nil, zero value otherwise.

### GetGuaranteedReceiveReconnectRetryWaitOk

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedReceiveReconnectRetryWaitOk() (*int32, bool)`

GetGuaranteedReceiveReconnectRetryWaitOk returns a tuple with the GuaranteedReceiveReconnectRetryWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedReceiveReconnectRetryWait

`func (o *MsgVpnJndiConnectionFactory) SetGuaranteedReceiveReconnectRetryWait(v int32)`

SetGuaranteedReceiveReconnectRetryWait sets GuaranteedReceiveReconnectRetryWait field to given value.

### HasGuaranteedReceiveReconnectRetryWait

`func (o *MsgVpnJndiConnectionFactory) HasGuaranteedReceiveReconnectRetryWait() bool`

HasGuaranteedReceiveReconnectRetryWait returns a boolean if a field has been set.

### GetGuaranteedReceiveWindowSize

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedReceiveWindowSize() int32`

GetGuaranteedReceiveWindowSize returns the GuaranteedReceiveWindowSize field if non-nil, zero value otherwise.

### GetGuaranteedReceiveWindowSizeOk

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedReceiveWindowSizeOk() (*int32, bool)`

GetGuaranteedReceiveWindowSizeOk returns a tuple with the GuaranteedReceiveWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedReceiveWindowSize

`func (o *MsgVpnJndiConnectionFactory) SetGuaranteedReceiveWindowSize(v int32)`

SetGuaranteedReceiveWindowSize sets GuaranteedReceiveWindowSize field to given value.

### HasGuaranteedReceiveWindowSize

`func (o *MsgVpnJndiConnectionFactory) HasGuaranteedReceiveWindowSize() bool`

HasGuaranteedReceiveWindowSize returns a boolean if a field has been set.

### GetGuaranteedReceiveWindowSizeAckThreshold

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedReceiveWindowSizeAckThreshold() int32`

GetGuaranteedReceiveWindowSizeAckThreshold returns the GuaranteedReceiveWindowSizeAckThreshold field if non-nil, zero value otherwise.

### GetGuaranteedReceiveWindowSizeAckThresholdOk

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedReceiveWindowSizeAckThresholdOk() (*int32, bool)`

GetGuaranteedReceiveWindowSizeAckThresholdOk returns a tuple with the GuaranteedReceiveWindowSizeAckThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedReceiveWindowSizeAckThreshold

`func (o *MsgVpnJndiConnectionFactory) SetGuaranteedReceiveWindowSizeAckThreshold(v int32)`

SetGuaranteedReceiveWindowSizeAckThreshold sets GuaranteedReceiveWindowSizeAckThreshold field to given value.

### HasGuaranteedReceiveWindowSizeAckThreshold

`func (o *MsgVpnJndiConnectionFactory) HasGuaranteedReceiveWindowSizeAckThreshold() bool`

HasGuaranteedReceiveWindowSizeAckThreshold returns a boolean if a field has been set.

### GetGuaranteedSendAckTimeout

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedSendAckTimeout() int32`

GetGuaranteedSendAckTimeout returns the GuaranteedSendAckTimeout field if non-nil, zero value otherwise.

### GetGuaranteedSendAckTimeoutOk

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedSendAckTimeoutOk() (*int32, bool)`

GetGuaranteedSendAckTimeoutOk returns a tuple with the GuaranteedSendAckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedSendAckTimeout

`func (o *MsgVpnJndiConnectionFactory) SetGuaranteedSendAckTimeout(v int32)`

SetGuaranteedSendAckTimeout sets GuaranteedSendAckTimeout field to given value.

### HasGuaranteedSendAckTimeout

`func (o *MsgVpnJndiConnectionFactory) HasGuaranteedSendAckTimeout() bool`

HasGuaranteedSendAckTimeout returns a boolean if a field has been set.

### GetGuaranteedSendWindowSize

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedSendWindowSize() int32`

GetGuaranteedSendWindowSize returns the GuaranteedSendWindowSize field if non-nil, zero value otherwise.

### GetGuaranteedSendWindowSizeOk

`func (o *MsgVpnJndiConnectionFactory) GetGuaranteedSendWindowSizeOk() (*int32, bool)`

GetGuaranteedSendWindowSizeOk returns a tuple with the GuaranteedSendWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedSendWindowSize

`func (o *MsgVpnJndiConnectionFactory) SetGuaranteedSendWindowSize(v int32)`

SetGuaranteedSendWindowSize sets GuaranteedSendWindowSize field to given value.

### HasGuaranteedSendWindowSize

`func (o *MsgVpnJndiConnectionFactory) HasGuaranteedSendWindowSize() bool`

HasGuaranteedSendWindowSize returns a boolean if a field has been set.

### GetMessagingDefaultDeliveryMode

`func (o *MsgVpnJndiConnectionFactory) GetMessagingDefaultDeliveryMode() string`

GetMessagingDefaultDeliveryMode returns the MessagingDefaultDeliveryMode field if non-nil, zero value otherwise.

### GetMessagingDefaultDeliveryModeOk

`func (o *MsgVpnJndiConnectionFactory) GetMessagingDefaultDeliveryModeOk() (*string, bool)`

GetMessagingDefaultDeliveryModeOk returns a tuple with the MessagingDefaultDeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingDefaultDeliveryMode

`func (o *MsgVpnJndiConnectionFactory) SetMessagingDefaultDeliveryMode(v string)`

SetMessagingDefaultDeliveryMode sets MessagingDefaultDeliveryMode field to given value.

### HasMessagingDefaultDeliveryMode

`func (o *MsgVpnJndiConnectionFactory) HasMessagingDefaultDeliveryMode() bool`

HasMessagingDefaultDeliveryMode returns a boolean if a field has been set.

### GetMessagingDefaultDmqEligibleEnabled

`func (o *MsgVpnJndiConnectionFactory) GetMessagingDefaultDmqEligibleEnabled() bool`

GetMessagingDefaultDmqEligibleEnabled returns the MessagingDefaultDmqEligibleEnabled field if non-nil, zero value otherwise.

### GetMessagingDefaultDmqEligibleEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetMessagingDefaultDmqEligibleEnabledOk() (*bool, bool)`

GetMessagingDefaultDmqEligibleEnabledOk returns a tuple with the MessagingDefaultDmqEligibleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingDefaultDmqEligibleEnabled

`func (o *MsgVpnJndiConnectionFactory) SetMessagingDefaultDmqEligibleEnabled(v bool)`

SetMessagingDefaultDmqEligibleEnabled sets MessagingDefaultDmqEligibleEnabled field to given value.

### HasMessagingDefaultDmqEligibleEnabled

`func (o *MsgVpnJndiConnectionFactory) HasMessagingDefaultDmqEligibleEnabled() bool`

HasMessagingDefaultDmqEligibleEnabled returns a boolean if a field has been set.

### GetMessagingDefaultElidingEligibleEnabled

`func (o *MsgVpnJndiConnectionFactory) GetMessagingDefaultElidingEligibleEnabled() bool`

GetMessagingDefaultElidingEligibleEnabled returns the MessagingDefaultElidingEligibleEnabled field if non-nil, zero value otherwise.

### GetMessagingDefaultElidingEligibleEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetMessagingDefaultElidingEligibleEnabledOk() (*bool, bool)`

GetMessagingDefaultElidingEligibleEnabledOk returns a tuple with the MessagingDefaultElidingEligibleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingDefaultElidingEligibleEnabled

`func (o *MsgVpnJndiConnectionFactory) SetMessagingDefaultElidingEligibleEnabled(v bool)`

SetMessagingDefaultElidingEligibleEnabled sets MessagingDefaultElidingEligibleEnabled field to given value.

### HasMessagingDefaultElidingEligibleEnabled

`func (o *MsgVpnJndiConnectionFactory) HasMessagingDefaultElidingEligibleEnabled() bool`

HasMessagingDefaultElidingEligibleEnabled returns a boolean if a field has been set.

### GetMessagingJmsxUserIdEnabled

`func (o *MsgVpnJndiConnectionFactory) GetMessagingJmsxUserIdEnabled() bool`

GetMessagingJmsxUserIdEnabled returns the MessagingJmsxUserIdEnabled field if non-nil, zero value otherwise.

### GetMessagingJmsxUserIdEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetMessagingJmsxUserIdEnabledOk() (*bool, bool)`

GetMessagingJmsxUserIdEnabledOk returns a tuple with the MessagingJmsxUserIdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingJmsxUserIdEnabled

`func (o *MsgVpnJndiConnectionFactory) SetMessagingJmsxUserIdEnabled(v bool)`

SetMessagingJmsxUserIdEnabled sets MessagingJmsxUserIdEnabled field to given value.

### HasMessagingJmsxUserIdEnabled

`func (o *MsgVpnJndiConnectionFactory) HasMessagingJmsxUserIdEnabled() bool`

HasMessagingJmsxUserIdEnabled returns a boolean if a field has been set.

### GetMessagingTextInXmlPayloadEnabled

`func (o *MsgVpnJndiConnectionFactory) GetMessagingTextInXmlPayloadEnabled() bool`

GetMessagingTextInXmlPayloadEnabled returns the MessagingTextInXmlPayloadEnabled field if non-nil, zero value otherwise.

### GetMessagingTextInXmlPayloadEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetMessagingTextInXmlPayloadEnabledOk() (*bool, bool)`

GetMessagingTextInXmlPayloadEnabledOk returns a tuple with the MessagingTextInXmlPayloadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingTextInXmlPayloadEnabled

`func (o *MsgVpnJndiConnectionFactory) SetMessagingTextInXmlPayloadEnabled(v bool)`

SetMessagingTextInXmlPayloadEnabled sets MessagingTextInXmlPayloadEnabled field to given value.

### HasMessagingTextInXmlPayloadEnabled

`func (o *MsgVpnJndiConnectionFactory) HasMessagingTextInXmlPayloadEnabled() bool`

HasMessagingTextInXmlPayloadEnabled returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnJndiConnectionFactory) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnJndiConnectionFactory) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnJndiConnectionFactory) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnJndiConnectionFactory) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetTransportCompressionLevel

`func (o *MsgVpnJndiConnectionFactory) GetTransportCompressionLevel() int32`

GetTransportCompressionLevel returns the TransportCompressionLevel field if non-nil, zero value otherwise.

### GetTransportCompressionLevelOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportCompressionLevelOk() (*int32, bool)`

GetTransportCompressionLevelOk returns a tuple with the TransportCompressionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCompressionLevel

`func (o *MsgVpnJndiConnectionFactory) SetTransportCompressionLevel(v int32)`

SetTransportCompressionLevel sets TransportCompressionLevel field to given value.

### HasTransportCompressionLevel

`func (o *MsgVpnJndiConnectionFactory) HasTransportCompressionLevel() bool`

HasTransportCompressionLevel returns a boolean if a field has been set.

### GetTransportConnectRetryCount

`func (o *MsgVpnJndiConnectionFactory) GetTransportConnectRetryCount() int32`

GetTransportConnectRetryCount returns the TransportConnectRetryCount field if non-nil, zero value otherwise.

### GetTransportConnectRetryCountOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportConnectRetryCountOk() (*int32, bool)`

GetTransportConnectRetryCountOk returns a tuple with the TransportConnectRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportConnectRetryCount

`func (o *MsgVpnJndiConnectionFactory) SetTransportConnectRetryCount(v int32)`

SetTransportConnectRetryCount sets TransportConnectRetryCount field to given value.

### HasTransportConnectRetryCount

`func (o *MsgVpnJndiConnectionFactory) HasTransportConnectRetryCount() bool`

HasTransportConnectRetryCount returns a boolean if a field has been set.

### GetTransportConnectRetryPerHostCount

`func (o *MsgVpnJndiConnectionFactory) GetTransportConnectRetryPerHostCount() int32`

GetTransportConnectRetryPerHostCount returns the TransportConnectRetryPerHostCount field if non-nil, zero value otherwise.

### GetTransportConnectRetryPerHostCountOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportConnectRetryPerHostCountOk() (*int32, bool)`

GetTransportConnectRetryPerHostCountOk returns a tuple with the TransportConnectRetryPerHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportConnectRetryPerHostCount

`func (o *MsgVpnJndiConnectionFactory) SetTransportConnectRetryPerHostCount(v int32)`

SetTransportConnectRetryPerHostCount sets TransportConnectRetryPerHostCount field to given value.

### HasTransportConnectRetryPerHostCount

`func (o *MsgVpnJndiConnectionFactory) HasTransportConnectRetryPerHostCount() bool`

HasTransportConnectRetryPerHostCount returns a boolean if a field has been set.

### GetTransportConnectTimeout

`func (o *MsgVpnJndiConnectionFactory) GetTransportConnectTimeout() int32`

GetTransportConnectTimeout returns the TransportConnectTimeout field if non-nil, zero value otherwise.

### GetTransportConnectTimeoutOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportConnectTimeoutOk() (*int32, bool)`

GetTransportConnectTimeoutOk returns a tuple with the TransportConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportConnectTimeout

`func (o *MsgVpnJndiConnectionFactory) SetTransportConnectTimeout(v int32)`

SetTransportConnectTimeout sets TransportConnectTimeout field to given value.

### HasTransportConnectTimeout

`func (o *MsgVpnJndiConnectionFactory) HasTransportConnectTimeout() bool`

HasTransportConnectTimeout returns a boolean if a field has been set.

### GetTransportDirectTransportEnabled

`func (o *MsgVpnJndiConnectionFactory) GetTransportDirectTransportEnabled() bool`

GetTransportDirectTransportEnabled returns the TransportDirectTransportEnabled field if non-nil, zero value otherwise.

### GetTransportDirectTransportEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportDirectTransportEnabledOk() (*bool, bool)`

GetTransportDirectTransportEnabledOk returns a tuple with the TransportDirectTransportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportDirectTransportEnabled

`func (o *MsgVpnJndiConnectionFactory) SetTransportDirectTransportEnabled(v bool)`

SetTransportDirectTransportEnabled sets TransportDirectTransportEnabled field to given value.

### HasTransportDirectTransportEnabled

`func (o *MsgVpnJndiConnectionFactory) HasTransportDirectTransportEnabled() bool`

HasTransportDirectTransportEnabled returns a boolean if a field has been set.

### GetTransportKeepaliveCount

`func (o *MsgVpnJndiConnectionFactory) GetTransportKeepaliveCount() int32`

GetTransportKeepaliveCount returns the TransportKeepaliveCount field if non-nil, zero value otherwise.

### GetTransportKeepaliveCountOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportKeepaliveCountOk() (*int32, bool)`

GetTransportKeepaliveCountOk returns a tuple with the TransportKeepaliveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportKeepaliveCount

`func (o *MsgVpnJndiConnectionFactory) SetTransportKeepaliveCount(v int32)`

SetTransportKeepaliveCount sets TransportKeepaliveCount field to given value.

### HasTransportKeepaliveCount

`func (o *MsgVpnJndiConnectionFactory) HasTransportKeepaliveCount() bool`

HasTransportKeepaliveCount returns a boolean if a field has been set.

### GetTransportKeepaliveEnabled

`func (o *MsgVpnJndiConnectionFactory) GetTransportKeepaliveEnabled() bool`

GetTransportKeepaliveEnabled returns the TransportKeepaliveEnabled field if non-nil, zero value otherwise.

### GetTransportKeepaliveEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportKeepaliveEnabledOk() (*bool, bool)`

GetTransportKeepaliveEnabledOk returns a tuple with the TransportKeepaliveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportKeepaliveEnabled

`func (o *MsgVpnJndiConnectionFactory) SetTransportKeepaliveEnabled(v bool)`

SetTransportKeepaliveEnabled sets TransportKeepaliveEnabled field to given value.

### HasTransportKeepaliveEnabled

`func (o *MsgVpnJndiConnectionFactory) HasTransportKeepaliveEnabled() bool`

HasTransportKeepaliveEnabled returns a boolean if a field has been set.

### GetTransportKeepaliveInterval

`func (o *MsgVpnJndiConnectionFactory) GetTransportKeepaliveInterval() int32`

GetTransportKeepaliveInterval returns the TransportKeepaliveInterval field if non-nil, zero value otherwise.

### GetTransportKeepaliveIntervalOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportKeepaliveIntervalOk() (*int32, bool)`

GetTransportKeepaliveIntervalOk returns a tuple with the TransportKeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportKeepaliveInterval

`func (o *MsgVpnJndiConnectionFactory) SetTransportKeepaliveInterval(v int32)`

SetTransportKeepaliveInterval sets TransportKeepaliveInterval field to given value.

### HasTransportKeepaliveInterval

`func (o *MsgVpnJndiConnectionFactory) HasTransportKeepaliveInterval() bool`

HasTransportKeepaliveInterval returns a boolean if a field has been set.

### GetTransportMsgCallbackOnIoThreadEnabled

`func (o *MsgVpnJndiConnectionFactory) GetTransportMsgCallbackOnIoThreadEnabled() bool`

GetTransportMsgCallbackOnIoThreadEnabled returns the TransportMsgCallbackOnIoThreadEnabled field if non-nil, zero value otherwise.

### GetTransportMsgCallbackOnIoThreadEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportMsgCallbackOnIoThreadEnabledOk() (*bool, bool)`

GetTransportMsgCallbackOnIoThreadEnabledOk returns a tuple with the TransportMsgCallbackOnIoThreadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMsgCallbackOnIoThreadEnabled

`func (o *MsgVpnJndiConnectionFactory) SetTransportMsgCallbackOnIoThreadEnabled(v bool)`

SetTransportMsgCallbackOnIoThreadEnabled sets TransportMsgCallbackOnIoThreadEnabled field to given value.

### HasTransportMsgCallbackOnIoThreadEnabled

`func (o *MsgVpnJndiConnectionFactory) HasTransportMsgCallbackOnIoThreadEnabled() bool`

HasTransportMsgCallbackOnIoThreadEnabled returns a boolean if a field has been set.

### GetTransportOptimizeDirectEnabled

`func (o *MsgVpnJndiConnectionFactory) GetTransportOptimizeDirectEnabled() bool`

GetTransportOptimizeDirectEnabled returns the TransportOptimizeDirectEnabled field if non-nil, zero value otherwise.

### GetTransportOptimizeDirectEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportOptimizeDirectEnabledOk() (*bool, bool)`

GetTransportOptimizeDirectEnabledOk returns a tuple with the TransportOptimizeDirectEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportOptimizeDirectEnabled

`func (o *MsgVpnJndiConnectionFactory) SetTransportOptimizeDirectEnabled(v bool)`

SetTransportOptimizeDirectEnabled sets TransportOptimizeDirectEnabled field to given value.

### HasTransportOptimizeDirectEnabled

`func (o *MsgVpnJndiConnectionFactory) HasTransportOptimizeDirectEnabled() bool`

HasTransportOptimizeDirectEnabled returns a boolean if a field has been set.

### GetTransportPort

`func (o *MsgVpnJndiConnectionFactory) GetTransportPort() int32`

GetTransportPort returns the TransportPort field if non-nil, zero value otherwise.

### GetTransportPortOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportPortOk() (*int32, bool)`

GetTransportPortOk returns a tuple with the TransportPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportPort

`func (o *MsgVpnJndiConnectionFactory) SetTransportPort(v int32)`

SetTransportPort sets TransportPort field to given value.

### HasTransportPort

`func (o *MsgVpnJndiConnectionFactory) HasTransportPort() bool`

HasTransportPort returns a boolean if a field has been set.

### GetTransportReadTimeout

`func (o *MsgVpnJndiConnectionFactory) GetTransportReadTimeout() int32`

GetTransportReadTimeout returns the TransportReadTimeout field if non-nil, zero value otherwise.

### GetTransportReadTimeoutOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportReadTimeoutOk() (*int32, bool)`

GetTransportReadTimeoutOk returns a tuple with the TransportReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportReadTimeout

`func (o *MsgVpnJndiConnectionFactory) SetTransportReadTimeout(v int32)`

SetTransportReadTimeout sets TransportReadTimeout field to given value.

### HasTransportReadTimeout

`func (o *MsgVpnJndiConnectionFactory) HasTransportReadTimeout() bool`

HasTransportReadTimeout returns a boolean if a field has been set.

### GetTransportReceiveBufferSize

`func (o *MsgVpnJndiConnectionFactory) GetTransportReceiveBufferSize() int32`

GetTransportReceiveBufferSize returns the TransportReceiveBufferSize field if non-nil, zero value otherwise.

### GetTransportReceiveBufferSizeOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportReceiveBufferSizeOk() (*int32, bool)`

GetTransportReceiveBufferSizeOk returns a tuple with the TransportReceiveBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportReceiveBufferSize

`func (o *MsgVpnJndiConnectionFactory) SetTransportReceiveBufferSize(v int32)`

SetTransportReceiveBufferSize sets TransportReceiveBufferSize field to given value.

### HasTransportReceiveBufferSize

`func (o *MsgVpnJndiConnectionFactory) HasTransportReceiveBufferSize() bool`

HasTransportReceiveBufferSize returns a boolean if a field has been set.

### GetTransportReconnectRetryCount

`func (o *MsgVpnJndiConnectionFactory) GetTransportReconnectRetryCount() int32`

GetTransportReconnectRetryCount returns the TransportReconnectRetryCount field if non-nil, zero value otherwise.

### GetTransportReconnectRetryCountOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportReconnectRetryCountOk() (*int32, bool)`

GetTransportReconnectRetryCountOk returns a tuple with the TransportReconnectRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportReconnectRetryCount

`func (o *MsgVpnJndiConnectionFactory) SetTransportReconnectRetryCount(v int32)`

SetTransportReconnectRetryCount sets TransportReconnectRetryCount field to given value.

### HasTransportReconnectRetryCount

`func (o *MsgVpnJndiConnectionFactory) HasTransportReconnectRetryCount() bool`

HasTransportReconnectRetryCount returns a boolean if a field has been set.

### GetTransportReconnectRetryWait

`func (o *MsgVpnJndiConnectionFactory) GetTransportReconnectRetryWait() int32`

GetTransportReconnectRetryWait returns the TransportReconnectRetryWait field if non-nil, zero value otherwise.

### GetTransportReconnectRetryWaitOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportReconnectRetryWaitOk() (*int32, bool)`

GetTransportReconnectRetryWaitOk returns a tuple with the TransportReconnectRetryWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportReconnectRetryWait

`func (o *MsgVpnJndiConnectionFactory) SetTransportReconnectRetryWait(v int32)`

SetTransportReconnectRetryWait sets TransportReconnectRetryWait field to given value.

### HasTransportReconnectRetryWait

`func (o *MsgVpnJndiConnectionFactory) HasTransportReconnectRetryWait() bool`

HasTransportReconnectRetryWait returns a boolean if a field has been set.

### GetTransportSendBufferSize

`func (o *MsgVpnJndiConnectionFactory) GetTransportSendBufferSize() int32`

GetTransportSendBufferSize returns the TransportSendBufferSize field if non-nil, zero value otherwise.

### GetTransportSendBufferSizeOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportSendBufferSizeOk() (*int32, bool)`

GetTransportSendBufferSizeOk returns a tuple with the TransportSendBufferSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportSendBufferSize

`func (o *MsgVpnJndiConnectionFactory) SetTransportSendBufferSize(v int32)`

SetTransportSendBufferSize sets TransportSendBufferSize field to given value.

### HasTransportSendBufferSize

`func (o *MsgVpnJndiConnectionFactory) HasTransportSendBufferSize() bool`

HasTransportSendBufferSize returns a boolean if a field has been set.

### GetTransportTcpNoDelayEnabled

`func (o *MsgVpnJndiConnectionFactory) GetTransportTcpNoDelayEnabled() bool`

GetTransportTcpNoDelayEnabled returns the TransportTcpNoDelayEnabled field if non-nil, zero value otherwise.

### GetTransportTcpNoDelayEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetTransportTcpNoDelayEnabledOk() (*bool, bool)`

GetTransportTcpNoDelayEnabledOk returns a tuple with the TransportTcpNoDelayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportTcpNoDelayEnabled

`func (o *MsgVpnJndiConnectionFactory) SetTransportTcpNoDelayEnabled(v bool)`

SetTransportTcpNoDelayEnabled sets TransportTcpNoDelayEnabled field to given value.

### HasTransportTcpNoDelayEnabled

`func (o *MsgVpnJndiConnectionFactory) HasTransportTcpNoDelayEnabled() bool`

HasTransportTcpNoDelayEnabled returns a boolean if a field has been set.

### GetXaEnabled

`func (o *MsgVpnJndiConnectionFactory) GetXaEnabled() bool`

GetXaEnabled returns the XaEnabled field if non-nil, zero value otherwise.

### GetXaEnabledOk

`func (o *MsgVpnJndiConnectionFactory) GetXaEnabledOk() (*bool, bool)`

GetXaEnabledOk returns a tuple with the XaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXaEnabled

`func (o *MsgVpnJndiConnectionFactory) SetXaEnabled(v bool)`

SetXaEnabled sets XaEnabled field to given value.

### HasXaEnabled

`func (o *MsgVpnJndiConnectionFactory) HasXaEnabled() bool`

HasXaEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


