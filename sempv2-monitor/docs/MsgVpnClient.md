# MsgVpnClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | Pointer to **string** | The name of the access control list (ACL) profile of the Client. | [optional] 
**AliasedFromMsgVpnName** | Pointer to **string** | The name of the original MsgVpn which the client signaled in. Available since 2.14. | [optional] 
**AlreadyBoundBindFailureCount** | Pointer to **int64** | The number of Client bind failures due to endpoint being already bound. | [optional] 
**AuthorizationGroupName** | Pointer to **string** | The name of the authorization group of the Client. | [optional] 
**AverageRxByteRate** | Pointer to **int64** | The one minute average of the message rate received from the Client, in bytes per second (B/sec). | [optional] 
**AverageRxMsgRate** | Pointer to **int64** | The one minute average of the message rate received from the Client, in messages per second (msg/sec). | [optional] 
**AverageTxByteRate** | Pointer to **int64** | The one minute average of the message rate transmitted to the Client, in bytes per second (B/sec). | [optional] 
**AverageTxMsgRate** | Pointer to **int64** | The one minute average of the message rate transmitted to the Client, in messages per second (msg/sec). | [optional] 
**BindRequestCount** | Pointer to **int64** | The number of Client requests to bind to an endpoint. | [optional] 
**BindSuccessCount** | Pointer to **int64** | The number of successful Client requests to bind to an endpoint. | [optional] 
**ClientAddress** | Pointer to **string** | The IP address and port of the Client. | [optional] 
**ClientId** | Pointer to **int32** | The identifier (ID) of the Client. | [optional] 
**ClientName** | Pointer to **string** | The name of the Client. | [optional] 
**ClientProfileName** | Pointer to **string** | The name of the client profile of the Client. | [optional] 
**ClientUsername** | Pointer to **string** | The client username of the Client used for authorization. | [optional] 
**ControlRxByteCount** | Pointer to **int64** | The amount of client control messages received from the Client, in bytes (B). | [optional] 
**ControlRxMsgCount** | Pointer to **int64** | The number of client control messages received from the Client. | [optional] 
**ControlTxByteCount** | Pointer to **int64** | The amount of client control messages transmitted to the Client, in bytes (B). | [optional] 
**ControlTxMsgCount** | Pointer to **int64** | The number of client control messages transmitted to the Client. | [optional] 
**CutThroughDeniedBindFailureCount** | Pointer to **int64** | The number of Client bind failures due to being denied cut-through forwarding. | [optional] 
**DataRxByteCount** | Pointer to **int64** | The amount of client data messages received from the Client, in bytes (B). | [optional] 
**DataRxMsgCount** | Pointer to **int64** | The number of client data messages received from the Client. | [optional] 
**DataTxByteCount** | Pointer to **int64** | The amount of client data messages transmitted to the Client, in bytes (B). | [optional] 
**DataTxMsgCount** | Pointer to **int64** | The number of client data messages transmitted to the Client. | [optional] 
**Description** | Pointer to **string** | The description text of the Client. | [optional] 
**DisabledBindFailureCount** | Pointer to **int64** | The number of Client bind failures due to endpoint being disabled. | [optional] 
**DtoLocalPriority** | Pointer to **int32** | The priority of the Client&#39;s subscriptions for receiving deliver-to-one (DTO) messages published on the local broker. | [optional] 
**DtoNetworkPriority** | Pointer to **int32** | The priority of the Client&#39;s subscriptions for receiving deliver-to-one (DTO) messages published on a remote broker. | [optional] 
**Eliding** | Pointer to **bool** | Indicates whether message eliding is enabled for the Client. | [optional] 
**ElidingTopicCount** | Pointer to **int32** | The number of topics requiring message eliding for the Client. | [optional] 
**ElidingTopicPeakCount** | Pointer to **int32** | The peak number of topics requiring message eliding for the Client. | [optional] 
**GuaranteedDeniedBindFailureCount** | Pointer to **int64** | The number of Client bind failures due to being denied guaranteed messaging. | [optional] 
**InvalidSelectorBindFailureCount** | Pointer to **int64** | The number of Client bind failures due to an invalid selector. | [optional] 
**Keepalive** | Pointer to **bool** | Indicates whether keepalive messages from the Client are received by the broker. Applicable for SMF and MQTT clients only. Available since 2.19. | [optional] 
**KeepaliveEffectiveTimeout** | Pointer to **int32** | The maximum period of time the broker will accept inactivity from the Client before disconnecting, in seconds. Available since 2.19. | [optional] 
**LargeMsgEventRaised** | Pointer to **bool** | Indicates whether the large-message event has been raised for the Client. | [optional] 
**LoginRxMsgCount** | Pointer to **int64** | The number of login request messages received from the Client. | [optional] 
**LoginTxMsgCount** | Pointer to **int64** | The number of login response messages transmitted to the Client. | [optional] 
**MaxBindCountExceededBindFailureCount** | Pointer to **int64** | The number of Client bind failures due to the endpoint maximum bind count being exceeded. | [optional] 
**MaxElidingTopicCountEventRaised** | Pointer to **bool** | Indicates whether the max-eliding-topic-count event has been raised for the Client. | [optional] 
**MqttConnackErrorTxCount** | Pointer to **int64** | The number of MQTT connect acknowledgment (CONNACK) refused response packets transmitted to the Client. | [optional] 
**MqttConnackTxCount** | Pointer to **int64** | The number of MQTT connect acknowledgment (CONNACK) accepted response packets transmitted to the Client. | [optional] 
**MqttConnectRxCount** | Pointer to **int64** | The number of MQTT connect (CONNECT) request packets received from the Client. | [optional] 
**MqttDisconnectRxCount** | Pointer to **int64** | The number of MQTT disconnect (DISCONNECT) request packets received from the Client. | [optional] 
**MqttPingreqRxCount** | Pointer to **int64** | The number of MQTT ping request (PINGREQ) packets received from the Client. | [optional] 
**MqttPingrespTxCount** | Pointer to **int64** | The number of MQTT ping response (PINGRESP) packets transmitted to the Client. | [optional] 
**MqttPubackRxCount** | Pointer to **int64** | The number of MQTT publish acknowledgement (PUBACK) response packets received from the Client. | [optional] 
**MqttPubackTxCount** | Pointer to **int64** | The number of MQTT publish acknowledgement (PUBACK) response packets transmitted to the Client. | [optional] 
**MqttPubcompTxCount** | Pointer to **int64** | The number of MQTT publish complete (PUBCOMP) packets transmitted to the Client in response to a PUBREL packet. These packets are the fourth and final packet of a QoS 2 protocol exchange. | [optional] 
**MqttPublishQos0RxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 0 message delivery. | [optional] 
**MqttPublishQos0TxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 0 message delivery. | [optional] 
**MqttPublishQos1RxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 1 message delivery. | [optional] 
**MqttPublishQos1TxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 1 message delivery. | [optional] 
**MqttPublishQos2RxCount** | Pointer to **int64** | The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 2 message delivery. | [optional] 
**MqttPubrecTxCount** | Pointer to **int64** | The number of MQTT publish received (PUBREC) packets transmitted to the Client in response to a PUBLISH packet with QoS 2. These packets are the second packet of a QoS 2 protocol exchange. | [optional] 
**MqttPubrelRxCount** | Pointer to **int64** | The number of MQTT publish release (PUBREL) packets received from the Client in response to a PUBREC packet. These packets are the third packet of a QoS 2 protocol exchange. | [optional] 
**MqttSubackErrorTxCount** | Pointer to **int64** | The number of MQTT subscribe acknowledgement (SUBACK) failure response packets transmitted to the Client. | [optional] 
**MqttSubackTxCount** | Pointer to **int64** | The number of MQTT subscribe acknowledgement (SUBACK) response packets transmitted to the Client. | [optional] 
**MqttSubscribeRxCount** | Pointer to **int64** | The number of MQTT subscribe (SUBSCRIBE) request packets received from the Client to create one or more topic subscriptions. | [optional] 
**MqttUnsubackTxCount** | Pointer to **int64** | The number of MQTT unsubscribe acknowledgement (UNSUBACK) response packets transmitted to the Client. | [optional] 
**MqttUnsubscribeRxCount** | Pointer to **int64** | The number of MQTT unsubscribe (UNSUBSCRIBE) request packets received from the Client to remove one or more topic subscriptions. | [optional] 
**MsgSpoolCongestionRxDiscardedMsgCount** | Pointer to **int64** | The number of messages from the Client discarded due to message spool congestion primarily caused by message promotion. | [optional] 
**MsgSpoolRxDiscardedMsgCount** | Pointer to **int64** | The number of messages from the Client discarded by the message spool. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**NoLocalDelivery** | Pointer to **bool** | Indicates whether not to deliver messages to the Client if it published them. | [optional] 
**NoSubscriptionMatchRxDiscardedMsgCount** | Pointer to **int64** | The number of messages from the Client discarded due to no matching subscription found. | [optional] 
**OriginalClientUsername** | Pointer to **string** | The original value of the client username used for Client authentication. | [optional] 
**OtherBindFailureCount** | Pointer to **int64** | The number of Client bind failures due to other reasons. | [optional] 
**Platform** | Pointer to **string** | The platform the Client application software was built for, which may include the OS and API type. | [optional] 
**PublishTopicAclRxDiscardedMsgCount** | Pointer to **int64** | The number of messages from the Client discarded due to the publish topic being denied by the Access Control List (ACL) profile. | [optional] 
**RestHttpRequestRxByteCount** | Pointer to **int64** | The amount of HTTP request messages received from the Client, in bytes (B). | [optional] 
**RestHttpRequestRxMsgCount** | Pointer to **int64** | The number of HTTP request messages received from the Client. | [optional] 
**RestHttpRequestTxByteCount** | Pointer to **int64** | The amount of HTTP request messages transmitted to the Client, in bytes (B). | [optional] 
**RestHttpRequestTxMsgCount** | Pointer to **int64** | The number of HTTP request messages transmitted to the Client. | [optional] 
**RestHttpResponseErrorRxMsgCount** | Pointer to **int64** | The number of HTTP client/server error response messages received from the Client. | [optional] 
**RestHttpResponseErrorTxMsgCount** | Pointer to **int64** | The number of HTTP client/server error response messages transmitted to the Client. | [optional] 
**RestHttpResponseRxByteCount** | Pointer to **int64** | The amount of HTTP response messages received from the Client, in bytes (B). | [optional] 
**RestHttpResponseRxMsgCount** | Pointer to **int64** | The number of HTTP response messages received from the Client. | [optional] 
**RestHttpResponseSuccessRxMsgCount** | Pointer to **int64** | The number of HTTP successful response messages received from the Client. | [optional] 
**RestHttpResponseSuccessTxMsgCount** | Pointer to **int64** | The number of HTTP successful response messages transmitted to the Client. | [optional] 
**RestHttpResponseTimeoutRxMsgCount** | Pointer to **int64** | The number of HTTP wait for reply timeout response messages received from the Client. | [optional] 
**RestHttpResponseTimeoutTxMsgCount** | Pointer to **int64** | The number of HTTP wait for reply timeout response messages transmitted to the Client. | [optional] 
**RestHttpResponseTxByteCount** | Pointer to **int64** | The amount of HTTP response messages transmitted to the Client, in bytes (B). | [optional] 
**RestHttpResponseTxMsgCount** | Pointer to **int64** | The number of HTTP response messages transmitted to the Client. | [optional] 
**RxByteCount** | Pointer to **int64** | The amount of messages received from the Client, in bytes (B). | [optional] 
**RxByteRate** | Pointer to **int64** | The current message rate received from the Client, in bytes per second (B/sec). | [optional] 
**RxDiscardedMsgCount** | Pointer to **int64** | The number of messages discarded during reception from the Client. | [optional] 
**RxMsgCount** | Pointer to **int64** | The number of messages received from the Client. | [optional] 
**RxMsgRate** | Pointer to **int64** | The current message rate received from the Client, in messages per second (msg/sec). | [optional] 
**ScheduledDisconnectTime** | Pointer to **int32** | The timestamp of when the Client will be disconnected by the broker. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.13. | [optional] 
**SlowSubscriber** | Pointer to **bool** | Indicates whether the Client is a slow subscriber and blocks for a few seconds when receiving messages. | [optional] 
**SoftwareDate** | Pointer to **string** | The date the Client application software was built. | [optional] 
**SoftwareVersion** | Pointer to **string** | The version of the Client application software. | [optional] 
**TlsCipherDescription** | Pointer to **string** | The description of the TLS cipher used by the Client, which may include cipher name, key exchange and encryption algorithms. | [optional] 
**TlsDowngradedToPlainText** | Pointer to **bool** | Indicates whether the Client TLS connection was downgraded to plain-text to increase performance. | [optional] 
**TlsVersion** | Pointer to **string** | The version of TLS used by the Client. | [optional] 
**TopicParseErrorRxDiscardedMsgCount** | Pointer to **int64** | The number of messages from the Client discarded due to an error while parsing the publish topic. | [optional] 
**TxByteCount** | Pointer to **int64** | The amount of messages transmitted to the Client, in bytes (B). | [optional] 
**TxByteRate** | Pointer to **int64** | The current message rate transmitted to the Client, in bytes per second (B/sec). | [optional] 
**TxDiscardedMsgCount** | Pointer to **int64** | The number of messages discarded during transmission to the Client. | [optional] 
**TxMsgCount** | Pointer to **int64** | The number of messages transmitted to the Client. | [optional] 
**TxMsgRate** | Pointer to **int64** | The current message rate transmitted to the Client, in messages per second (msg/sec). | [optional] 
**Uptime** | Pointer to **int32** | The amount of time in seconds since the Client connected. | [optional] 
**User** | Pointer to **string** | The user description for the Client, which may include computer name and process ID. | [optional] 
**VirtualRouter** | Pointer to **string** | The virtual router used by the Client. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The Client is using the primary virtual router. \&quot;backup\&quot; - The Client is using the backup virtual router. \&quot;internal\&quot; - The Client is using the internal virtual router. \&quot;unknown\&quot; - The Client virtual router is unknown. &lt;/pre&gt;  | [optional] 
**WebInactiveTimeout** | Pointer to **int32** | The maximum web transport timeout for the Client being inactive, in seconds. | [optional] 
**WebMaxPayload** | Pointer to **int64** | The maximum web transport message payload size which excludes the size of the message header, in bytes. | [optional] 
**WebParseErrorRxDiscardedMsgCount** | Pointer to **int64** | The number of messages from the Client discarded due to an error while parsing the web message. | [optional] 
**WebRemainingTimeout** | Pointer to **int32** | The remaining web transport timeout for the Client being inactive, in seconds. | [optional] 
**WebRxByteCount** | Pointer to **int64** | The amount of web transport messages received from the Client, in bytes (B). | [optional] 
**WebRxEncoding** | Pointer to **string** | The type of encoding used during reception from the Client. The allowed values and their meaning are:  &lt;pre&gt; \&quot;binary\&quot; - The Client is using binary encoding. \&quot;base64\&quot; - The Client is using base64 encoding. \&quot;illegal\&quot; - The Client is using an illegal encoding type. &lt;/pre&gt;  | [optional] 
**WebRxMsgCount** | Pointer to **int64** | The number of web transport messages received from the Client. | [optional] 
**WebRxProtocol** | Pointer to **string** | The type of web transport used during reception from the Client. The allowed values and their meaning are:  &lt;pre&gt; \&quot;ws-binary\&quot; - The Client is using WebSocket binary transport. \&quot;http-binary-streaming\&quot; - The Client is using HTTP binary streaming transport. \&quot;http-binary\&quot; - The Client is using HTTP binary transport. \&quot;http-base64\&quot; - The Client is using HTTP base64 transport. &lt;/pre&gt;  | [optional] 
**WebRxRequestCount** | Pointer to **int64** | The number of web transport requests received from the Client (HTTP only). Not available for WebSockets. | [optional] 
**WebRxResponseCount** | Pointer to **int64** | The number of web transport responses transmitted to the Client on the receive connection (HTTP only). Not available for WebSockets. | [optional] 
**WebRxTcpState** | Pointer to **string** | The TCP state of the receive connection from the Client. When fully operational, should be: established. See RFC 793 for further details. The allowed values and their meaning are:  &lt;pre&gt; \&quot;closed\&quot; - No connection state at all. \&quot;listen\&quot; - Waiting for a connection request from any remote TCP and port. \&quot;syn-sent\&quot; - Waiting for a matching connection request after having sent a connection request. \&quot;syn-received\&quot; - Waiting for a confirming connection request acknowledgment after having both received and sent a connection request. \&quot;established\&quot; - An open connection, data received can be delivered to the user. \&quot;close-wait\&quot; - Waiting for a connection termination request from the local user. \&quot;fin-wait-1\&quot; - Waiting for a connection termination request from the remote TCP, or an acknowledgment of the connection termination request previously sent. \&quot;closing\&quot; - Waiting for a connection termination request acknowledgment from the remote TCP. \&quot;last-ack\&quot; - Waiting for an acknowledgment of the connection termination request previously sent to the remote TCP. \&quot;fin-wait-2\&quot; - Waiting for a connection termination request from the remote TCP. \&quot;time-wait\&quot; - Waiting for enough time to pass to be sure the remote TCP received the acknowledgment of its connection termination request. &lt;/pre&gt;  | [optional] 
**WebRxTlsCipherDescription** | Pointer to **string** | The description of the TLS cipher received from the Client, which may include cipher name, key exchange and encryption algorithms. | [optional] 
**WebRxTlsVersion** | Pointer to **string** | The version of TLS used during reception from the Client. | [optional] 
**WebSessionId** | Pointer to **string** | The identifier (ID) of the web transport session for the Client. | [optional] 
**WebTxByteCount** | Pointer to **int64** | The amount of web transport messages transmitted to the Client, in bytes (B). | [optional] 
**WebTxEncoding** | Pointer to **string** | The type of encoding used during transmission to the Client. The allowed values and their meaning are:  &lt;pre&gt; \&quot;binary\&quot; - The Client is using binary encoding. \&quot;base64\&quot; - The Client is using base64 encoding. \&quot;illegal\&quot; - The Client is using an illegal encoding type. &lt;/pre&gt;  | [optional] 
**WebTxMsgCount** | Pointer to **int64** | The number of web transport messages transmitted to the Client. | [optional] 
**WebTxProtocol** | Pointer to **string** | The type of web transport used during transmission to the Client. The allowed values and their meaning are:  &lt;pre&gt; \&quot;ws-binary\&quot; - The Client is using WebSocket binary transport. \&quot;http-binary-streaming\&quot; - The Client is using HTTP binary streaming transport. \&quot;http-binary\&quot; - The Client is using HTTP binary transport. \&quot;http-base64\&quot; - The Client is using HTTP base64 transport. &lt;/pre&gt;  | [optional] 
**WebTxRequestCount** | Pointer to **int64** | The number of web transport requests transmitted to the Client (HTTP only). Not available for WebSockets. | [optional] 
**WebTxResponseCount** | Pointer to **int64** | The number of web transport responses received from the Client on the transmit connection (HTTP only). Not available for WebSockets. | [optional] 
**WebTxTcpState** | Pointer to **string** | The TCP state of the transmit connection to the Client. When fully operational, should be: established. See RFC 793 for further details. The allowed values and their meaning are:  &lt;pre&gt; \&quot;closed\&quot; - No connection state at all. \&quot;listen\&quot; - Waiting for a connection request from any remote TCP and port. \&quot;syn-sent\&quot; - Waiting for a matching connection request after having sent a connection request. \&quot;syn-received\&quot; - Waiting for a confirming connection request acknowledgment after having both received and sent a connection request. \&quot;established\&quot; - An open connection, data received can be delivered to the user. \&quot;close-wait\&quot; - Waiting for a connection termination request from the local user. \&quot;fin-wait-1\&quot; - Waiting for a connection termination request from the remote TCP, or an acknowledgment of the connection termination request previously sent. \&quot;closing\&quot; - Waiting for a connection termination request acknowledgment from the remote TCP. \&quot;last-ack\&quot; - Waiting for an acknowledgment of the connection termination request previously sent to the remote TCP. \&quot;fin-wait-2\&quot; - Waiting for a connection termination request from the remote TCP. \&quot;time-wait\&quot; - Waiting for enough time to pass to be sure the remote TCP received the acknowledgment of its connection termination request. &lt;/pre&gt;  | [optional] 
**WebTxTlsCipherDescription** | Pointer to **string** | The description of the TLS cipher transmitted to the Client, which may include cipher name, key exchange and encryption algorithms. | [optional] 
**WebTxTlsVersion** | Pointer to **string** | The version of TLS used during transmission to the Client. | [optional] 

## Methods

### NewMsgVpnClient

`func NewMsgVpnClient() *MsgVpnClient`

NewMsgVpnClient instantiates a new MsgVpnClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientWithDefaults

`func NewMsgVpnClientWithDefaults() *MsgVpnClient`

NewMsgVpnClientWithDefaults instantiates a new MsgVpnClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfileName

`func (o *MsgVpnClient) GetAclProfileName() string`

GetAclProfileName returns the AclProfileName field if non-nil, zero value otherwise.

### GetAclProfileNameOk

`func (o *MsgVpnClient) GetAclProfileNameOk() (*string, bool)`

GetAclProfileNameOk returns a tuple with the AclProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfileName

`func (o *MsgVpnClient) SetAclProfileName(v string)`

SetAclProfileName sets AclProfileName field to given value.

### HasAclProfileName

`func (o *MsgVpnClient) HasAclProfileName() bool`

HasAclProfileName returns a boolean if a field has been set.

### GetAliasedFromMsgVpnName

`func (o *MsgVpnClient) GetAliasedFromMsgVpnName() string`

GetAliasedFromMsgVpnName returns the AliasedFromMsgVpnName field if non-nil, zero value otherwise.

### GetAliasedFromMsgVpnNameOk

`func (o *MsgVpnClient) GetAliasedFromMsgVpnNameOk() (*string, bool)`

GetAliasedFromMsgVpnNameOk returns a tuple with the AliasedFromMsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasedFromMsgVpnName

`func (o *MsgVpnClient) SetAliasedFromMsgVpnName(v string)`

SetAliasedFromMsgVpnName sets AliasedFromMsgVpnName field to given value.

### HasAliasedFromMsgVpnName

`func (o *MsgVpnClient) HasAliasedFromMsgVpnName() bool`

HasAliasedFromMsgVpnName returns a boolean if a field has been set.

### GetAlreadyBoundBindFailureCount

`func (o *MsgVpnClient) GetAlreadyBoundBindFailureCount() int64`

GetAlreadyBoundBindFailureCount returns the AlreadyBoundBindFailureCount field if non-nil, zero value otherwise.

### GetAlreadyBoundBindFailureCountOk

`func (o *MsgVpnClient) GetAlreadyBoundBindFailureCountOk() (*int64, bool)`

GetAlreadyBoundBindFailureCountOk returns a tuple with the AlreadyBoundBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyBoundBindFailureCount

`func (o *MsgVpnClient) SetAlreadyBoundBindFailureCount(v int64)`

SetAlreadyBoundBindFailureCount sets AlreadyBoundBindFailureCount field to given value.

### HasAlreadyBoundBindFailureCount

`func (o *MsgVpnClient) HasAlreadyBoundBindFailureCount() bool`

HasAlreadyBoundBindFailureCount returns a boolean if a field has been set.

### GetAuthorizationGroupName

`func (o *MsgVpnClient) GetAuthorizationGroupName() string`

GetAuthorizationGroupName returns the AuthorizationGroupName field if non-nil, zero value otherwise.

### GetAuthorizationGroupNameOk

`func (o *MsgVpnClient) GetAuthorizationGroupNameOk() (*string, bool)`

GetAuthorizationGroupNameOk returns a tuple with the AuthorizationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroupName

`func (o *MsgVpnClient) SetAuthorizationGroupName(v string)`

SetAuthorizationGroupName sets AuthorizationGroupName field to given value.

### HasAuthorizationGroupName

`func (o *MsgVpnClient) HasAuthorizationGroupName() bool`

HasAuthorizationGroupName returns a boolean if a field has been set.

### GetAverageRxByteRate

`func (o *MsgVpnClient) GetAverageRxByteRate() int64`

GetAverageRxByteRate returns the AverageRxByteRate field if non-nil, zero value otherwise.

### GetAverageRxByteRateOk

`func (o *MsgVpnClient) GetAverageRxByteRateOk() (*int64, bool)`

GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxByteRate

`func (o *MsgVpnClient) SetAverageRxByteRate(v int64)`

SetAverageRxByteRate sets AverageRxByteRate field to given value.

### HasAverageRxByteRate

`func (o *MsgVpnClient) HasAverageRxByteRate() bool`

HasAverageRxByteRate returns a boolean if a field has been set.

### GetAverageRxMsgRate

`func (o *MsgVpnClient) GetAverageRxMsgRate() int64`

GetAverageRxMsgRate returns the AverageRxMsgRate field if non-nil, zero value otherwise.

### GetAverageRxMsgRateOk

`func (o *MsgVpnClient) GetAverageRxMsgRateOk() (*int64, bool)`

GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxMsgRate

`func (o *MsgVpnClient) SetAverageRxMsgRate(v int64)`

SetAverageRxMsgRate sets AverageRxMsgRate field to given value.

### HasAverageRxMsgRate

`func (o *MsgVpnClient) HasAverageRxMsgRate() bool`

HasAverageRxMsgRate returns a boolean if a field has been set.

### GetAverageTxByteRate

`func (o *MsgVpnClient) GetAverageTxByteRate() int64`

GetAverageTxByteRate returns the AverageTxByteRate field if non-nil, zero value otherwise.

### GetAverageTxByteRateOk

`func (o *MsgVpnClient) GetAverageTxByteRateOk() (*int64, bool)`

GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxByteRate

`func (o *MsgVpnClient) SetAverageTxByteRate(v int64)`

SetAverageTxByteRate sets AverageTxByteRate field to given value.

### HasAverageTxByteRate

`func (o *MsgVpnClient) HasAverageTxByteRate() bool`

HasAverageTxByteRate returns a boolean if a field has been set.

### GetAverageTxMsgRate

`func (o *MsgVpnClient) GetAverageTxMsgRate() int64`

GetAverageTxMsgRate returns the AverageTxMsgRate field if non-nil, zero value otherwise.

### GetAverageTxMsgRateOk

`func (o *MsgVpnClient) GetAverageTxMsgRateOk() (*int64, bool)`

GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxMsgRate

`func (o *MsgVpnClient) SetAverageTxMsgRate(v int64)`

SetAverageTxMsgRate sets AverageTxMsgRate field to given value.

### HasAverageTxMsgRate

`func (o *MsgVpnClient) HasAverageTxMsgRate() bool`

HasAverageTxMsgRate returns a boolean if a field has been set.

### GetBindRequestCount

`func (o *MsgVpnClient) GetBindRequestCount() int64`

GetBindRequestCount returns the BindRequestCount field if non-nil, zero value otherwise.

### GetBindRequestCountOk

`func (o *MsgVpnClient) GetBindRequestCountOk() (*int64, bool)`

GetBindRequestCountOk returns a tuple with the BindRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindRequestCount

`func (o *MsgVpnClient) SetBindRequestCount(v int64)`

SetBindRequestCount sets BindRequestCount field to given value.

### HasBindRequestCount

`func (o *MsgVpnClient) HasBindRequestCount() bool`

HasBindRequestCount returns a boolean if a field has been set.

### GetBindSuccessCount

`func (o *MsgVpnClient) GetBindSuccessCount() int64`

GetBindSuccessCount returns the BindSuccessCount field if non-nil, zero value otherwise.

### GetBindSuccessCountOk

`func (o *MsgVpnClient) GetBindSuccessCountOk() (*int64, bool)`

GetBindSuccessCountOk returns a tuple with the BindSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindSuccessCount

`func (o *MsgVpnClient) SetBindSuccessCount(v int64)`

SetBindSuccessCount sets BindSuccessCount field to given value.

### HasBindSuccessCount

`func (o *MsgVpnClient) HasBindSuccessCount() bool`

HasBindSuccessCount returns a boolean if a field has been set.

### GetClientAddress

`func (o *MsgVpnClient) GetClientAddress() string`

GetClientAddress returns the ClientAddress field if non-nil, zero value otherwise.

### GetClientAddressOk

`func (o *MsgVpnClient) GetClientAddressOk() (*string, bool)`

GetClientAddressOk returns a tuple with the ClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAddress

`func (o *MsgVpnClient) SetClientAddress(v string)`

SetClientAddress sets ClientAddress field to given value.

### HasClientAddress

`func (o *MsgVpnClient) HasClientAddress() bool`

HasClientAddress returns a boolean if a field has been set.

### GetClientId

`func (o *MsgVpnClient) GetClientId() int32`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MsgVpnClient) GetClientIdOk() (*int32, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MsgVpnClient) SetClientId(v int32)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MsgVpnClient) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientName

`func (o *MsgVpnClient) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnClient) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnClient) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnClient) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientProfileName

`func (o *MsgVpnClient) GetClientProfileName() string`

GetClientProfileName returns the ClientProfileName field if non-nil, zero value otherwise.

### GetClientProfileNameOk

`func (o *MsgVpnClient) GetClientProfileNameOk() (*string, bool)`

GetClientProfileNameOk returns a tuple with the ClientProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileName

`func (o *MsgVpnClient) SetClientProfileName(v string)`

SetClientProfileName sets ClientProfileName field to given value.

### HasClientProfileName

`func (o *MsgVpnClient) HasClientProfileName() bool`

HasClientProfileName returns a boolean if a field has been set.

### GetClientUsername

`func (o *MsgVpnClient) GetClientUsername() string`

GetClientUsername returns the ClientUsername field if non-nil, zero value otherwise.

### GetClientUsernameOk

`func (o *MsgVpnClient) GetClientUsernameOk() (*string, bool)`

GetClientUsernameOk returns a tuple with the ClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUsername

`func (o *MsgVpnClient) SetClientUsername(v string)`

SetClientUsername sets ClientUsername field to given value.

### HasClientUsername

`func (o *MsgVpnClient) HasClientUsername() bool`

HasClientUsername returns a boolean if a field has been set.

### GetControlRxByteCount

`func (o *MsgVpnClient) GetControlRxByteCount() int64`

GetControlRxByteCount returns the ControlRxByteCount field if non-nil, zero value otherwise.

### GetControlRxByteCountOk

`func (o *MsgVpnClient) GetControlRxByteCountOk() (*int64, bool)`

GetControlRxByteCountOk returns a tuple with the ControlRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRxByteCount

`func (o *MsgVpnClient) SetControlRxByteCount(v int64)`

SetControlRxByteCount sets ControlRxByteCount field to given value.

### HasControlRxByteCount

`func (o *MsgVpnClient) HasControlRxByteCount() bool`

HasControlRxByteCount returns a boolean if a field has been set.

### GetControlRxMsgCount

`func (o *MsgVpnClient) GetControlRxMsgCount() int64`

GetControlRxMsgCount returns the ControlRxMsgCount field if non-nil, zero value otherwise.

### GetControlRxMsgCountOk

`func (o *MsgVpnClient) GetControlRxMsgCountOk() (*int64, bool)`

GetControlRxMsgCountOk returns a tuple with the ControlRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRxMsgCount

`func (o *MsgVpnClient) SetControlRxMsgCount(v int64)`

SetControlRxMsgCount sets ControlRxMsgCount field to given value.

### HasControlRxMsgCount

`func (o *MsgVpnClient) HasControlRxMsgCount() bool`

HasControlRxMsgCount returns a boolean if a field has been set.

### GetControlTxByteCount

`func (o *MsgVpnClient) GetControlTxByteCount() int64`

GetControlTxByteCount returns the ControlTxByteCount field if non-nil, zero value otherwise.

### GetControlTxByteCountOk

`func (o *MsgVpnClient) GetControlTxByteCountOk() (*int64, bool)`

GetControlTxByteCountOk returns a tuple with the ControlTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTxByteCount

`func (o *MsgVpnClient) SetControlTxByteCount(v int64)`

SetControlTxByteCount sets ControlTxByteCount field to given value.

### HasControlTxByteCount

`func (o *MsgVpnClient) HasControlTxByteCount() bool`

HasControlTxByteCount returns a boolean if a field has been set.

### GetControlTxMsgCount

`func (o *MsgVpnClient) GetControlTxMsgCount() int64`

GetControlTxMsgCount returns the ControlTxMsgCount field if non-nil, zero value otherwise.

### GetControlTxMsgCountOk

`func (o *MsgVpnClient) GetControlTxMsgCountOk() (*int64, bool)`

GetControlTxMsgCountOk returns a tuple with the ControlTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTxMsgCount

`func (o *MsgVpnClient) SetControlTxMsgCount(v int64)`

SetControlTxMsgCount sets ControlTxMsgCount field to given value.

### HasControlTxMsgCount

`func (o *MsgVpnClient) HasControlTxMsgCount() bool`

HasControlTxMsgCount returns a boolean if a field has been set.

### GetCutThroughDeniedBindFailureCount

`func (o *MsgVpnClient) GetCutThroughDeniedBindFailureCount() int64`

GetCutThroughDeniedBindFailureCount returns the CutThroughDeniedBindFailureCount field if non-nil, zero value otherwise.

### GetCutThroughDeniedBindFailureCountOk

`func (o *MsgVpnClient) GetCutThroughDeniedBindFailureCountOk() (*int64, bool)`

GetCutThroughDeniedBindFailureCountOk returns a tuple with the CutThroughDeniedBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutThroughDeniedBindFailureCount

`func (o *MsgVpnClient) SetCutThroughDeniedBindFailureCount(v int64)`

SetCutThroughDeniedBindFailureCount sets CutThroughDeniedBindFailureCount field to given value.

### HasCutThroughDeniedBindFailureCount

`func (o *MsgVpnClient) HasCutThroughDeniedBindFailureCount() bool`

HasCutThroughDeniedBindFailureCount returns a boolean if a field has been set.

### GetDataRxByteCount

`func (o *MsgVpnClient) GetDataRxByteCount() int64`

GetDataRxByteCount returns the DataRxByteCount field if non-nil, zero value otherwise.

### GetDataRxByteCountOk

`func (o *MsgVpnClient) GetDataRxByteCountOk() (*int64, bool)`

GetDataRxByteCountOk returns a tuple with the DataRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxByteCount

`func (o *MsgVpnClient) SetDataRxByteCount(v int64)`

SetDataRxByteCount sets DataRxByteCount field to given value.

### HasDataRxByteCount

`func (o *MsgVpnClient) HasDataRxByteCount() bool`

HasDataRxByteCount returns a boolean if a field has been set.

### GetDataRxMsgCount

`func (o *MsgVpnClient) GetDataRxMsgCount() int64`

GetDataRxMsgCount returns the DataRxMsgCount field if non-nil, zero value otherwise.

### GetDataRxMsgCountOk

`func (o *MsgVpnClient) GetDataRxMsgCountOk() (*int64, bool)`

GetDataRxMsgCountOk returns a tuple with the DataRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxMsgCount

`func (o *MsgVpnClient) SetDataRxMsgCount(v int64)`

SetDataRxMsgCount sets DataRxMsgCount field to given value.

### HasDataRxMsgCount

`func (o *MsgVpnClient) HasDataRxMsgCount() bool`

HasDataRxMsgCount returns a boolean if a field has been set.

### GetDataTxByteCount

`func (o *MsgVpnClient) GetDataTxByteCount() int64`

GetDataTxByteCount returns the DataTxByteCount field if non-nil, zero value otherwise.

### GetDataTxByteCountOk

`func (o *MsgVpnClient) GetDataTxByteCountOk() (*int64, bool)`

GetDataTxByteCountOk returns a tuple with the DataTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxByteCount

`func (o *MsgVpnClient) SetDataTxByteCount(v int64)`

SetDataTxByteCount sets DataTxByteCount field to given value.

### HasDataTxByteCount

`func (o *MsgVpnClient) HasDataTxByteCount() bool`

HasDataTxByteCount returns a boolean if a field has been set.

### GetDataTxMsgCount

`func (o *MsgVpnClient) GetDataTxMsgCount() int64`

GetDataTxMsgCount returns the DataTxMsgCount field if non-nil, zero value otherwise.

### GetDataTxMsgCountOk

`func (o *MsgVpnClient) GetDataTxMsgCountOk() (*int64, bool)`

GetDataTxMsgCountOk returns a tuple with the DataTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxMsgCount

`func (o *MsgVpnClient) SetDataTxMsgCount(v int64)`

SetDataTxMsgCount sets DataTxMsgCount field to given value.

### HasDataTxMsgCount

`func (o *MsgVpnClient) HasDataTxMsgCount() bool`

HasDataTxMsgCount returns a boolean if a field has been set.

### GetDescription

`func (o *MsgVpnClient) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MsgVpnClient) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MsgVpnClient) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MsgVpnClient) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabledBindFailureCount

`func (o *MsgVpnClient) GetDisabledBindFailureCount() int64`

GetDisabledBindFailureCount returns the DisabledBindFailureCount field if non-nil, zero value otherwise.

### GetDisabledBindFailureCountOk

`func (o *MsgVpnClient) GetDisabledBindFailureCountOk() (*int64, bool)`

GetDisabledBindFailureCountOk returns a tuple with the DisabledBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledBindFailureCount

`func (o *MsgVpnClient) SetDisabledBindFailureCount(v int64)`

SetDisabledBindFailureCount sets DisabledBindFailureCount field to given value.

### HasDisabledBindFailureCount

`func (o *MsgVpnClient) HasDisabledBindFailureCount() bool`

HasDisabledBindFailureCount returns a boolean if a field has been set.

### GetDtoLocalPriority

`func (o *MsgVpnClient) GetDtoLocalPriority() int32`

GetDtoLocalPriority returns the DtoLocalPriority field if non-nil, zero value otherwise.

### GetDtoLocalPriorityOk

`func (o *MsgVpnClient) GetDtoLocalPriorityOk() (*int32, bool)`

GetDtoLocalPriorityOk returns a tuple with the DtoLocalPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtoLocalPriority

`func (o *MsgVpnClient) SetDtoLocalPriority(v int32)`

SetDtoLocalPriority sets DtoLocalPriority field to given value.

### HasDtoLocalPriority

`func (o *MsgVpnClient) HasDtoLocalPriority() bool`

HasDtoLocalPriority returns a boolean if a field has been set.

### GetDtoNetworkPriority

`func (o *MsgVpnClient) GetDtoNetworkPriority() int32`

GetDtoNetworkPriority returns the DtoNetworkPriority field if non-nil, zero value otherwise.

### GetDtoNetworkPriorityOk

`func (o *MsgVpnClient) GetDtoNetworkPriorityOk() (*int32, bool)`

GetDtoNetworkPriorityOk returns a tuple with the DtoNetworkPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtoNetworkPriority

`func (o *MsgVpnClient) SetDtoNetworkPriority(v int32)`

SetDtoNetworkPriority sets DtoNetworkPriority field to given value.

### HasDtoNetworkPriority

`func (o *MsgVpnClient) HasDtoNetworkPriority() bool`

HasDtoNetworkPriority returns a boolean if a field has been set.

### GetEliding

`func (o *MsgVpnClient) GetEliding() bool`

GetEliding returns the Eliding field if non-nil, zero value otherwise.

### GetElidingOk

`func (o *MsgVpnClient) GetElidingOk() (*bool, bool)`

GetElidingOk returns a tuple with the Eliding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEliding

`func (o *MsgVpnClient) SetEliding(v bool)`

SetEliding sets Eliding field to given value.

### HasEliding

`func (o *MsgVpnClient) HasEliding() bool`

HasEliding returns a boolean if a field has been set.

### GetElidingTopicCount

`func (o *MsgVpnClient) GetElidingTopicCount() int32`

GetElidingTopicCount returns the ElidingTopicCount field if non-nil, zero value otherwise.

### GetElidingTopicCountOk

`func (o *MsgVpnClient) GetElidingTopicCountOk() (*int32, bool)`

GetElidingTopicCountOk returns a tuple with the ElidingTopicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElidingTopicCount

`func (o *MsgVpnClient) SetElidingTopicCount(v int32)`

SetElidingTopicCount sets ElidingTopicCount field to given value.

### HasElidingTopicCount

`func (o *MsgVpnClient) HasElidingTopicCount() bool`

HasElidingTopicCount returns a boolean if a field has been set.

### GetElidingTopicPeakCount

`func (o *MsgVpnClient) GetElidingTopicPeakCount() int32`

GetElidingTopicPeakCount returns the ElidingTopicPeakCount field if non-nil, zero value otherwise.

### GetElidingTopicPeakCountOk

`func (o *MsgVpnClient) GetElidingTopicPeakCountOk() (*int32, bool)`

GetElidingTopicPeakCountOk returns a tuple with the ElidingTopicPeakCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElidingTopicPeakCount

`func (o *MsgVpnClient) SetElidingTopicPeakCount(v int32)`

SetElidingTopicPeakCount sets ElidingTopicPeakCount field to given value.

### HasElidingTopicPeakCount

`func (o *MsgVpnClient) HasElidingTopicPeakCount() bool`

HasElidingTopicPeakCount returns a boolean if a field has been set.

### GetGuaranteedDeniedBindFailureCount

`func (o *MsgVpnClient) GetGuaranteedDeniedBindFailureCount() int64`

GetGuaranteedDeniedBindFailureCount returns the GuaranteedDeniedBindFailureCount field if non-nil, zero value otherwise.

### GetGuaranteedDeniedBindFailureCountOk

`func (o *MsgVpnClient) GetGuaranteedDeniedBindFailureCountOk() (*int64, bool)`

GetGuaranteedDeniedBindFailureCountOk returns a tuple with the GuaranteedDeniedBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedDeniedBindFailureCount

`func (o *MsgVpnClient) SetGuaranteedDeniedBindFailureCount(v int64)`

SetGuaranteedDeniedBindFailureCount sets GuaranteedDeniedBindFailureCount field to given value.

### HasGuaranteedDeniedBindFailureCount

`func (o *MsgVpnClient) HasGuaranteedDeniedBindFailureCount() bool`

HasGuaranteedDeniedBindFailureCount returns a boolean if a field has been set.

### GetInvalidSelectorBindFailureCount

`func (o *MsgVpnClient) GetInvalidSelectorBindFailureCount() int64`

GetInvalidSelectorBindFailureCount returns the InvalidSelectorBindFailureCount field if non-nil, zero value otherwise.

### GetInvalidSelectorBindFailureCountOk

`func (o *MsgVpnClient) GetInvalidSelectorBindFailureCountOk() (*int64, bool)`

GetInvalidSelectorBindFailureCountOk returns a tuple with the InvalidSelectorBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidSelectorBindFailureCount

`func (o *MsgVpnClient) SetInvalidSelectorBindFailureCount(v int64)`

SetInvalidSelectorBindFailureCount sets InvalidSelectorBindFailureCount field to given value.

### HasInvalidSelectorBindFailureCount

`func (o *MsgVpnClient) HasInvalidSelectorBindFailureCount() bool`

HasInvalidSelectorBindFailureCount returns a boolean if a field has been set.

### GetKeepalive

`func (o *MsgVpnClient) GetKeepalive() bool`

GetKeepalive returns the Keepalive field if non-nil, zero value otherwise.

### GetKeepaliveOk

`func (o *MsgVpnClient) GetKeepaliveOk() (*bool, bool)`

GetKeepaliveOk returns a tuple with the Keepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepalive

`func (o *MsgVpnClient) SetKeepalive(v bool)`

SetKeepalive sets Keepalive field to given value.

### HasKeepalive

`func (o *MsgVpnClient) HasKeepalive() bool`

HasKeepalive returns a boolean if a field has been set.

### GetKeepaliveEffectiveTimeout

`func (o *MsgVpnClient) GetKeepaliveEffectiveTimeout() int32`

GetKeepaliveEffectiveTimeout returns the KeepaliveEffectiveTimeout field if non-nil, zero value otherwise.

### GetKeepaliveEffectiveTimeoutOk

`func (o *MsgVpnClient) GetKeepaliveEffectiveTimeoutOk() (*int32, bool)`

GetKeepaliveEffectiveTimeoutOk returns a tuple with the KeepaliveEffectiveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveEffectiveTimeout

`func (o *MsgVpnClient) SetKeepaliveEffectiveTimeout(v int32)`

SetKeepaliveEffectiveTimeout sets KeepaliveEffectiveTimeout field to given value.

### HasKeepaliveEffectiveTimeout

`func (o *MsgVpnClient) HasKeepaliveEffectiveTimeout() bool`

HasKeepaliveEffectiveTimeout returns a boolean if a field has been set.

### GetLargeMsgEventRaised

`func (o *MsgVpnClient) GetLargeMsgEventRaised() bool`

GetLargeMsgEventRaised returns the LargeMsgEventRaised field if non-nil, zero value otherwise.

### GetLargeMsgEventRaisedOk

`func (o *MsgVpnClient) GetLargeMsgEventRaisedOk() (*bool, bool)`

GetLargeMsgEventRaisedOk returns a tuple with the LargeMsgEventRaised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeMsgEventRaised

`func (o *MsgVpnClient) SetLargeMsgEventRaised(v bool)`

SetLargeMsgEventRaised sets LargeMsgEventRaised field to given value.

### HasLargeMsgEventRaised

`func (o *MsgVpnClient) HasLargeMsgEventRaised() bool`

HasLargeMsgEventRaised returns a boolean if a field has been set.

### GetLoginRxMsgCount

`func (o *MsgVpnClient) GetLoginRxMsgCount() int64`

GetLoginRxMsgCount returns the LoginRxMsgCount field if non-nil, zero value otherwise.

### GetLoginRxMsgCountOk

`func (o *MsgVpnClient) GetLoginRxMsgCountOk() (*int64, bool)`

GetLoginRxMsgCountOk returns a tuple with the LoginRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginRxMsgCount

`func (o *MsgVpnClient) SetLoginRxMsgCount(v int64)`

SetLoginRxMsgCount sets LoginRxMsgCount field to given value.

### HasLoginRxMsgCount

`func (o *MsgVpnClient) HasLoginRxMsgCount() bool`

HasLoginRxMsgCount returns a boolean if a field has been set.

### GetLoginTxMsgCount

`func (o *MsgVpnClient) GetLoginTxMsgCount() int64`

GetLoginTxMsgCount returns the LoginTxMsgCount field if non-nil, zero value otherwise.

### GetLoginTxMsgCountOk

`func (o *MsgVpnClient) GetLoginTxMsgCountOk() (*int64, bool)`

GetLoginTxMsgCountOk returns a tuple with the LoginTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginTxMsgCount

`func (o *MsgVpnClient) SetLoginTxMsgCount(v int64)`

SetLoginTxMsgCount sets LoginTxMsgCount field to given value.

### HasLoginTxMsgCount

`func (o *MsgVpnClient) HasLoginTxMsgCount() bool`

HasLoginTxMsgCount returns a boolean if a field has been set.

### GetMaxBindCountExceededBindFailureCount

`func (o *MsgVpnClient) GetMaxBindCountExceededBindFailureCount() int64`

GetMaxBindCountExceededBindFailureCount returns the MaxBindCountExceededBindFailureCount field if non-nil, zero value otherwise.

### GetMaxBindCountExceededBindFailureCountOk

`func (o *MsgVpnClient) GetMaxBindCountExceededBindFailureCountOk() (*int64, bool)`

GetMaxBindCountExceededBindFailureCountOk returns a tuple with the MaxBindCountExceededBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBindCountExceededBindFailureCount

`func (o *MsgVpnClient) SetMaxBindCountExceededBindFailureCount(v int64)`

SetMaxBindCountExceededBindFailureCount sets MaxBindCountExceededBindFailureCount field to given value.

### HasMaxBindCountExceededBindFailureCount

`func (o *MsgVpnClient) HasMaxBindCountExceededBindFailureCount() bool`

HasMaxBindCountExceededBindFailureCount returns a boolean if a field has been set.

### GetMaxElidingTopicCountEventRaised

`func (o *MsgVpnClient) GetMaxElidingTopicCountEventRaised() bool`

GetMaxElidingTopicCountEventRaised returns the MaxElidingTopicCountEventRaised field if non-nil, zero value otherwise.

### GetMaxElidingTopicCountEventRaisedOk

`func (o *MsgVpnClient) GetMaxElidingTopicCountEventRaisedOk() (*bool, bool)`

GetMaxElidingTopicCountEventRaisedOk returns a tuple with the MaxElidingTopicCountEventRaised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElidingTopicCountEventRaised

`func (o *MsgVpnClient) SetMaxElidingTopicCountEventRaised(v bool)`

SetMaxElidingTopicCountEventRaised sets MaxElidingTopicCountEventRaised field to given value.

### HasMaxElidingTopicCountEventRaised

`func (o *MsgVpnClient) HasMaxElidingTopicCountEventRaised() bool`

HasMaxElidingTopicCountEventRaised returns a boolean if a field has been set.

### GetMqttConnackErrorTxCount

`func (o *MsgVpnClient) GetMqttConnackErrorTxCount() int64`

GetMqttConnackErrorTxCount returns the MqttConnackErrorTxCount field if non-nil, zero value otherwise.

### GetMqttConnackErrorTxCountOk

`func (o *MsgVpnClient) GetMqttConnackErrorTxCountOk() (*int64, bool)`

GetMqttConnackErrorTxCountOk returns a tuple with the MqttConnackErrorTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConnackErrorTxCount

`func (o *MsgVpnClient) SetMqttConnackErrorTxCount(v int64)`

SetMqttConnackErrorTxCount sets MqttConnackErrorTxCount field to given value.

### HasMqttConnackErrorTxCount

`func (o *MsgVpnClient) HasMqttConnackErrorTxCount() bool`

HasMqttConnackErrorTxCount returns a boolean if a field has been set.

### GetMqttConnackTxCount

`func (o *MsgVpnClient) GetMqttConnackTxCount() int64`

GetMqttConnackTxCount returns the MqttConnackTxCount field if non-nil, zero value otherwise.

### GetMqttConnackTxCountOk

`func (o *MsgVpnClient) GetMqttConnackTxCountOk() (*int64, bool)`

GetMqttConnackTxCountOk returns a tuple with the MqttConnackTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConnackTxCount

`func (o *MsgVpnClient) SetMqttConnackTxCount(v int64)`

SetMqttConnackTxCount sets MqttConnackTxCount field to given value.

### HasMqttConnackTxCount

`func (o *MsgVpnClient) HasMqttConnackTxCount() bool`

HasMqttConnackTxCount returns a boolean if a field has been set.

### GetMqttConnectRxCount

`func (o *MsgVpnClient) GetMqttConnectRxCount() int64`

GetMqttConnectRxCount returns the MqttConnectRxCount field if non-nil, zero value otherwise.

### GetMqttConnectRxCountOk

`func (o *MsgVpnClient) GetMqttConnectRxCountOk() (*int64, bool)`

GetMqttConnectRxCountOk returns a tuple with the MqttConnectRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttConnectRxCount

`func (o *MsgVpnClient) SetMqttConnectRxCount(v int64)`

SetMqttConnectRxCount sets MqttConnectRxCount field to given value.

### HasMqttConnectRxCount

`func (o *MsgVpnClient) HasMqttConnectRxCount() bool`

HasMqttConnectRxCount returns a boolean if a field has been set.

### GetMqttDisconnectRxCount

`func (o *MsgVpnClient) GetMqttDisconnectRxCount() int64`

GetMqttDisconnectRxCount returns the MqttDisconnectRxCount field if non-nil, zero value otherwise.

### GetMqttDisconnectRxCountOk

`func (o *MsgVpnClient) GetMqttDisconnectRxCountOk() (*int64, bool)`

GetMqttDisconnectRxCountOk returns a tuple with the MqttDisconnectRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttDisconnectRxCount

`func (o *MsgVpnClient) SetMqttDisconnectRxCount(v int64)`

SetMqttDisconnectRxCount sets MqttDisconnectRxCount field to given value.

### HasMqttDisconnectRxCount

`func (o *MsgVpnClient) HasMqttDisconnectRxCount() bool`

HasMqttDisconnectRxCount returns a boolean if a field has been set.

### GetMqttPingreqRxCount

`func (o *MsgVpnClient) GetMqttPingreqRxCount() int64`

GetMqttPingreqRxCount returns the MqttPingreqRxCount field if non-nil, zero value otherwise.

### GetMqttPingreqRxCountOk

`func (o *MsgVpnClient) GetMqttPingreqRxCountOk() (*int64, bool)`

GetMqttPingreqRxCountOk returns a tuple with the MqttPingreqRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPingreqRxCount

`func (o *MsgVpnClient) SetMqttPingreqRxCount(v int64)`

SetMqttPingreqRxCount sets MqttPingreqRxCount field to given value.

### HasMqttPingreqRxCount

`func (o *MsgVpnClient) HasMqttPingreqRxCount() bool`

HasMqttPingreqRxCount returns a boolean if a field has been set.

### GetMqttPingrespTxCount

`func (o *MsgVpnClient) GetMqttPingrespTxCount() int64`

GetMqttPingrespTxCount returns the MqttPingrespTxCount field if non-nil, zero value otherwise.

### GetMqttPingrespTxCountOk

`func (o *MsgVpnClient) GetMqttPingrespTxCountOk() (*int64, bool)`

GetMqttPingrespTxCountOk returns a tuple with the MqttPingrespTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPingrespTxCount

`func (o *MsgVpnClient) SetMqttPingrespTxCount(v int64)`

SetMqttPingrespTxCount sets MqttPingrespTxCount field to given value.

### HasMqttPingrespTxCount

`func (o *MsgVpnClient) HasMqttPingrespTxCount() bool`

HasMqttPingrespTxCount returns a boolean if a field has been set.

### GetMqttPubackRxCount

`func (o *MsgVpnClient) GetMqttPubackRxCount() int64`

GetMqttPubackRxCount returns the MqttPubackRxCount field if non-nil, zero value otherwise.

### GetMqttPubackRxCountOk

`func (o *MsgVpnClient) GetMqttPubackRxCountOk() (*int64, bool)`

GetMqttPubackRxCountOk returns a tuple with the MqttPubackRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubackRxCount

`func (o *MsgVpnClient) SetMqttPubackRxCount(v int64)`

SetMqttPubackRxCount sets MqttPubackRxCount field to given value.

### HasMqttPubackRxCount

`func (o *MsgVpnClient) HasMqttPubackRxCount() bool`

HasMqttPubackRxCount returns a boolean if a field has been set.

### GetMqttPubackTxCount

`func (o *MsgVpnClient) GetMqttPubackTxCount() int64`

GetMqttPubackTxCount returns the MqttPubackTxCount field if non-nil, zero value otherwise.

### GetMqttPubackTxCountOk

`func (o *MsgVpnClient) GetMqttPubackTxCountOk() (*int64, bool)`

GetMqttPubackTxCountOk returns a tuple with the MqttPubackTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubackTxCount

`func (o *MsgVpnClient) SetMqttPubackTxCount(v int64)`

SetMqttPubackTxCount sets MqttPubackTxCount field to given value.

### HasMqttPubackTxCount

`func (o *MsgVpnClient) HasMqttPubackTxCount() bool`

HasMqttPubackTxCount returns a boolean if a field has been set.

### GetMqttPubcompTxCount

`func (o *MsgVpnClient) GetMqttPubcompTxCount() int64`

GetMqttPubcompTxCount returns the MqttPubcompTxCount field if non-nil, zero value otherwise.

### GetMqttPubcompTxCountOk

`func (o *MsgVpnClient) GetMqttPubcompTxCountOk() (*int64, bool)`

GetMqttPubcompTxCountOk returns a tuple with the MqttPubcompTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubcompTxCount

`func (o *MsgVpnClient) SetMqttPubcompTxCount(v int64)`

SetMqttPubcompTxCount sets MqttPubcompTxCount field to given value.

### HasMqttPubcompTxCount

`func (o *MsgVpnClient) HasMqttPubcompTxCount() bool`

HasMqttPubcompTxCount returns a boolean if a field has been set.

### GetMqttPublishQos0RxCount

`func (o *MsgVpnClient) GetMqttPublishQos0RxCount() int64`

GetMqttPublishQos0RxCount returns the MqttPublishQos0RxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos0RxCountOk

`func (o *MsgVpnClient) GetMqttPublishQos0RxCountOk() (*int64, bool)`

GetMqttPublishQos0RxCountOk returns a tuple with the MqttPublishQos0RxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos0RxCount

`func (o *MsgVpnClient) SetMqttPublishQos0RxCount(v int64)`

SetMqttPublishQos0RxCount sets MqttPublishQos0RxCount field to given value.

### HasMqttPublishQos0RxCount

`func (o *MsgVpnClient) HasMqttPublishQos0RxCount() bool`

HasMqttPublishQos0RxCount returns a boolean if a field has been set.

### GetMqttPublishQos0TxCount

`func (o *MsgVpnClient) GetMqttPublishQos0TxCount() int64`

GetMqttPublishQos0TxCount returns the MqttPublishQos0TxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos0TxCountOk

`func (o *MsgVpnClient) GetMqttPublishQos0TxCountOk() (*int64, bool)`

GetMqttPublishQos0TxCountOk returns a tuple with the MqttPublishQos0TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos0TxCount

`func (o *MsgVpnClient) SetMqttPublishQos0TxCount(v int64)`

SetMqttPublishQos0TxCount sets MqttPublishQos0TxCount field to given value.

### HasMqttPublishQos0TxCount

`func (o *MsgVpnClient) HasMqttPublishQos0TxCount() bool`

HasMqttPublishQos0TxCount returns a boolean if a field has been set.

### GetMqttPublishQos1RxCount

`func (o *MsgVpnClient) GetMqttPublishQos1RxCount() int64`

GetMqttPublishQos1RxCount returns the MqttPublishQos1RxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos1RxCountOk

`func (o *MsgVpnClient) GetMqttPublishQos1RxCountOk() (*int64, bool)`

GetMqttPublishQos1RxCountOk returns a tuple with the MqttPublishQos1RxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos1RxCount

`func (o *MsgVpnClient) SetMqttPublishQos1RxCount(v int64)`

SetMqttPublishQos1RxCount sets MqttPublishQos1RxCount field to given value.

### HasMqttPublishQos1RxCount

`func (o *MsgVpnClient) HasMqttPublishQos1RxCount() bool`

HasMqttPublishQos1RxCount returns a boolean if a field has been set.

### GetMqttPublishQos1TxCount

`func (o *MsgVpnClient) GetMqttPublishQos1TxCount() int64`

GetMqttPublishQos1TxCount returns the MqttPublishQos1TxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos1TxCountOk

`func (o *MsgVpnClient) GetMqttPublishQos1TxCountOk() (*int64, bool)`

GetMqttPublishQos1TxCountOk returns a tuple with the MqttPublishQos1TxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos1TxCount

`func (o *MsgVpnClient) SetMqttPublishQos1TxCount(v int64)`

SetMqttPublishQos1TxCount sets MqttPublishQos1TxCount field to given value.

### HasMqttPublishQos1TxCount

`func (o *MsgVpnClient) HasMqttPublishQos1TxCount() bool`

HasMqttPublishQos1TxCount returns a boolean if a field has been set.

### GetMqttPublishQos2RxCount

`func (o *MsgVpnClient) GetMqttPublishQos2RxCount() int64`

GetMqttPublishQos2RxCount returns the MqttPublishQos2RxCount field if non-nil, zero value otherwise.

### GetMqttPublishQos2RxCountOk

`func (o *MsgVpnClient) GetMqttPublishQos2RxCountOk() (*int64, bool)`

GetMqttPublishQos2RxCountOk returns a tuple with the MqttPublishQos2RxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPublishQos2RxCount

`func (o *MsgVpnClient) SetMqttPublishQos2RxCount(v int64)`

SetMqttPublishQos2RxCount sets MqttPublishQos2RxCount field to given value.

### HasMqttPublishQos2RxCount

`func (o *MsgVpnClient) HasMqttPublishQos2RxCount() bool`

HasMqttPublishQos2RxCount returns a boolean if a field has been set.

### GetMqttPubrecTxCount

`func (o *MsgVpnClient) GetMqttPubrecTxCount() int64`

GetMqttPubrecTxCount returns the MqttPubrecTxCount field if non-nil, zero value otherwise.

### GetMqttPubrecTxCountOk

`func (o *MsgVpnClient) GetMqttPubrecTxCountOk() (*int64, bool)`

GetMqttPubrecTxCountOk returns a tuple with the MqttPubrecTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubrecTxCount

`func (o *MsgVpnClient) SetMqttPubrecTxCount(v int64)`

SetMqttPubrecTxCount sets MqttPubrecTxCount field to given value.

### HasMqttPubrecTxCount

`func (o *MsgVpnClient) HasMqttPubrecTxCount() bool`

HasMqttPubrecTxCount returns a boolean if a field has been set.

### GetMqttPubrelRxCount

`func (o *MsgVpnClient) GetMqttPubrelRxCount() int64`

GetMqttPubrelRxCount returns the MqttPubrelRxCount field if non-nil, zero value otherwise.

### GetMqttPubrelRxCountOk

`func (o *MsgVpnClient) GetMqttPubrelRxCountOk() (*int64, bool)`

GetMqttPubrelRxCountOk returns a tuple with the MqttPubrelRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttPubrelRxCount

`func (o *MsgVpnClient) SetMqttPubrelRxCount(v int64)`

SetMqttPubrelRxCount sets MqttPubrelRxCount field to given value.

### HasMqttPubrelRxCount

`func (o *MsgVpnClient) HasMqttPubrelRxCount() bool`

HasMqttPubrelRxCount returns a boolean if a field has been set.

### GetMqttSubackErrorTxCount

`func (o *MsgVpnClient) GetMqttSubackErrorTxCount() int64`

GetMqttSubackErrorTxCount returns the MqttSubackErrorTxCount field if non-nil, zero value otherwise.

### GetMqttSubackErrorTxCountOk

`func (o *MsgVpnClient) GetMqttSubackErrorTxCountOk() (*int64, bool)`

GetMqttSubackErrorTxCountOk returns a tuple with the MqttSubackErrorTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttSubackErrorTxCount

`func (o *MsgVpnClient) SetMqttSubackErrorTxCount(v int64)`

SetMqttSubackErrorTxCount sets MqttSubackErrorTxCount field to given value.

### HasMqttSubackErrorTxCount

`func (o *MsgVpnClient) HasMqttSubackErrorTxCount() bool`

HasMqttSubackErrorTxCount returns a boolean if a field has been set.

### GetMqttSubackTxCount

`func (o *MsgVpnClient) GetMqttSubackTxCount() int64`

GetMqttSubackTxCount returns the MqttSubackTxCount field if non-nil, zero value otherwise.

### GetMqttSubackTxCountOk

`func (o *MsgVpnClient) GetMqttSubackTxCountOk() (*int64, bool)`

GetMqttSubackTxCountOk returns a tuple with the MqttSubackTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttSubackTxCount

`func (o *MsgVpnClient) SetMqttSubackTxCount(v int64)`

SetMqttSubackTxCount sets MqttSubackTxCount field to given value.

### HasMqttSubackTxCount

`func (o *MsgVpnClient) HasMqttSubackTxCount() bool`

HasMqttSubackTxCount returns a boolean if a field has been set.

### GetMqttSubscribeRxCount

`func (o *MsgVpnClient) GetMqttSubscribeRxCount() int64`

GetMqttSubscribeRxCount returns the MqttSubscribeRxCount field if non-nil, zero value otherwise.

### GetMqttSubscribeRxCountOk

`func (o *MsgVpnClient) GetMqttSubscribeRxCountOk() (*int64, bool)`

GetMqttSubscribeRxCountOk returns a tuple with the MqttSubscribeRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttSubscribeRxCount

`func (o *MsgVpnClient) SetMqttSubscribeRxCount(v int64)`

SetMqttSubscribeRxCount sets MqttSubscribeRxCount field to given value.

### HasMqttSubscribeRxCount

`func (o *MsgVpnClient) HasMqttSubscribeRxCount() bool`

HasMqttSubscribeRxCount returns a boolean if a field has been set.

### GetMqttUnsubackTxCount

`func (o *MsgVpnClient) GetMqttUnsubackTxCount() int64`

GetMqttUnsubackTxCount returns the MqttUnsubackTxCount field if non-nil, zero value otherwise.

### GetMqttUnsubackTxCountOk

`func (o *MsgVpnClient) GetMqttUnsubackTxCountOk() (*int64, bool)`

GetMqttUnsubackTxCountOk returns a tuple with the MqttUnsubackTxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttUnsubackTxCount

`func (o *MsgVpnClient) SetMqttUnsubackTxCount(v int64)`

SetMqttUnsubackTxCount sets MqttUnsubackTxCount field to given value.

### HasMqttUnsubackTxCount

`func (o *MsgVpnClient) HasMqttUnsubackTxCount() bool`

HasMqttUnsubackTxCount returns a boolean if a field has been set.

### GetMqttUnsubscribeRxCount

`func (o *MsgVpnClient) GetMqttUnsubscribeRxCount() int64`

GetMqttUnsubscribeRxCount returns the MqttUnsubscribeRxCount field if non-nil, zero value otherwise.

### GetMqttUnsubscribeRxCountOk

`func (o *MsgVpnClient) GetMqttUnsubscribeRxCountOk() (*int64, bool)`

GetMqttUnsubscribeRxCountOk returns a tuple with the MqttUnsubscribeRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttUnsubscribeRxCount

`func (o *MsgVpnClient) SetMqttUnsubscribeRxCount(v int64)`

SetMqttUnsubscribeRxCount sets MqttUnsubscribeRxCount field to given value.

### HasMqttUnsubscribeRxCount

`func (o *MsgVpnClient) HasMqttUnsubscribeRxCount() bool`

HasMqttUnsubscribeRxCount returns a boolean if a field has been set.

### GetMsgSpoolCongestionRxDiscardedMsgCount

`func (o *MsgVpnClient) GetMsgSpoolCongestionRxDiscardedMsgCount() int64`

GetMsgSpoolCongestionRxDiscardedMsgCount returns the MsgSpoolCongestionRxDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMsgSpoolCongestionRxDiscardedMsgCountOk

`func (o *MsgVpnClient) GetMsgSpoolCongestionRxDiscardedMsgCountOk() (*int64, bool)`

GetMsgSpoolCongestionRxDiscardedMsgCountOk returns a tuple with the MsgSpoolCongestionRxDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolCongestionRxDiscardedMsgCount

`func (o *MsgVpnClient) SetMsgSpoolCongestionRxDiscardedMsgCount(v int64)`

SetMsgSpoolCongestionRxDiscardedMsgCount sets MsgSpoolCongestionRxDiscardedMsgCount field to given value.

### HasMsgSpoolCongestionRxDiscardedMsgCount

`func (o *MsgVpnClient) HasMsgSpoolCongestionRxDiscardedMsgCount() bool`

HasMsgSpoolCongestionRxDiscardedMsgCount returns a boolean if a field has been set.

### GetMsgSpoolRxDiscardedMsgCount

`func (o *MsgVpnClient) GetMsgSpoolRxDiscardedMsgCount() int64`

GetMsgSpoolRxDiscardedMsgCount returns the MsgSpoolRxDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMsgSpoolRxDiscardedMsgCountOk

`func (o *MsgVpnClient) GetMsgSpoolRxDiscardedMsgCountOk() (*int64, bool)`

GetMsgSpoolRxDiscardedMsgCountOk returns a tuple with the MsgSpoolRxDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolRxDiscardedMsgCount

`func (o *MsgVpnClient) SetMsgSpoolRxDiscardedMsgCount(v int64)`

SetMsgSpoolRxDiscardedMsgCount sets MsgSpoolRxDiscardedMsgCount field to given value.

### HasMsgSpoolRxDiscardedMsgCount

`func (o *MsgVpnClient) HasMsgSpoolRxDiscardedMsgCount() bool`

HasMsgSpoolRxDiscardedMsgCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnClient) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnClient) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnClient) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnClient) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetNoLocalDelivery

`func (o *MsgVpnClient) GetNoLocalDelivery() bool`

GetNoLocalDelivery returns the NoLocalDelivery field if non-nil, zero value otherwise.

### GetNoLocalDeliveryOk

`func (o *MsgVpnClient) GetNoLocalDeliveryOk() (*bool, bool)`

GetNoLocalDeliveryOk returns a tuple with the NoLocalDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLocalDelivery

`func (o *MsgVpnClient) SetNoLocalDelivery(v bool)`

SetNoLocalDelivery sets NoLocalDelivery field to given value.

### HasNoLocalDelivery

`func (o *MsgVpnClient) HasNoLocalDelivery() bool`

HasNoLocalDelivery returns a boolean if a field has been set.

### GetNoSubscriptionMatchRxDiscardedMsgCount

`func (o *MsgVpnClient) GetNoSubscriptionMatchRxDiscardedMsgCount() int64`

GetNoSubscriptionMatchRxDiscardedMsgCount returns the NoSubscriptionMatchRxDiscardedMsgCount field if non-nil, zero value otherwise.

### GetNoSubscriptionMatchRxDiscardedMsgCountOk

`func (o *MsgVpnClient) GetNoSubscriptionMatchRxDiscardedMsgCountOk() (*int64, bool)`

GetNoSubscriptionMatchRxDiscardedMsgCountOk returns a tuple with the NoSubscriptionMatchRxDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSubscriptionMatchRxDiscardedMsgCount

`func (o *MsgVpnClient) SetNoSubscriptionMatchRxDiscardedMsgCount(v int64)`

SetNoSubscriptionMatchRxDiscardedMsgCount sets NoSubscriptionMatchRxDiscardedMsgCount field to given value.

### HasNoSubscriptionMatchRxDiscardedMsgCount

`func (o *MsgVpnClient) HasNoSubscriptionMatchRxDiscardedMsgCount() bool`

HasNoSubscriptionMatchRxDiscardedMsgCount returns a boolean if a field has been set.

### GetOriginalClientUsername

`func (o *MsgVpnClient) GetOriginalClientUsername() string`

GetOriginalClientUsername returns the OriginalClientUsername field if non-nil, zero value otherwise.

### GetOriginalClientUsernameOk

`func (o *MsgVpnClient) GetOriginalClientUsernameOk() (*string, bool)`

GetOriginalClientUsernameOk returns a tuple with the OriginalClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalClientUsername

`func (o *MsgVpnClient) SetOriginalClientUsername(v string)`

SetOriginalClientUsername sets OriginalClientUsername field to given value.

### HasOriginalClientUsername

`func (o *MsgVpnClient) HasOriginalClientUsername() bool`

HasOriginalClientUsername returns a boolean if a field has been set.

### GetOtherBindFailureCount

`func (o *MsgVpnClient) GetOtherBindFailureCount() int64`

GetOtherBindFailureCount returns the OtherBindFailureCount field if non-nil, zero value otherwise.

### GetOtherBindFailureCountOk

`func (o *MsgVpnClient) GetOtherBindFailureCountOk() (*int64, bool)`

GetOtherBindFailureCountOk returns a tuple with the OtherBindFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherBindFailureCount

`func (o *MsgVpnClient) SetOtherBindFailureCount(v int64)`

SetOtherBindFailureCount sets OtherBindFailureCount field to given value.

### HasOtherBindFailureCount

`func (o *MsgVpnClient) HasOtherBindFailureCount() bool`

HasOtherBindFailureCount returns a boolean if a field has been set.

### GetPlatform

`func (o *MsgVpnClient) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *MsgVpnClient) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *MsgVpnClient) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *MsgVpnClient) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPublishTopicAclRxDiscardedMsgCount

`func (o *MsgVpnClient) GetPublishTopicAclRxDiscardedMsgCount() int64`

GetPublishTopicAclRxDiscardedMsgCount returns the PublishTopicAclRxDiscardedMsgCount field if non-nil, zero value otherwise.

### GetPublishTopicAclRxDiscardedMsgCountOk

`func (o *MsgVpnClient) GetPublishTopicAclRxDiscardedMsgCountOk() (*int64, bool)`

GetPublishTopicAclRxDiscardedMsgCountOk returns a tuple with the PublishTopicAclRxDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTopicAclRxDiscardedMsgCount

`func (o *MsgVpnClient) SetPublishTopicAclRxDiscardedMsgCount(v int64)`

SetPublishTopicAclRxDiscardedMsgCount sets PublishTopicAclRxDiscardedMsgCount field to given value.

### HasPublishTopicAclRxDiscardedMsgCount

`func (o *MsgVpnClient) HasPublishTopicAclRxDiscardedMsgCount() bool`

HasPublishTopicAclRxDiscardedMsgCount returns a boolean if a field has been set.

### GetRestHttpRequestRxByteCount

`func (o *MsgVpnClient) GetRestHttpRequestRxByteCount() int64`

GetRestHttpRequestRxByteCount returns the RestHttpRequestRxByteCount field if non-nil, zero value otherwise.

### GetRestHttpRequestRxByteCountOk

`func (o *MsgVpnClient) GetRestHttpRequestRxByteCountOk() (*int64, bool)`

GetRestHttpRequestRxByteCountOk returns a tuple with the RestHttpRequestRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpRequestRxByteCount

`func (o *MsgVpnClient) SetRestHttpRequestRxByteCount(v int64)`

SetRestHttpRequestRxByteCount sets RestHttpRequestRxByteCount field to given value.

### HasRestHttpRequestRxByteCount

`func (o *MsgVpnClient) HasRestHttpRequestRxByteCount() bool`

HasRestHttpRequestRxByteCount returns a boolean if a field has been set.

### GetRestHttpRequestRxMsgCount

`func (o *MsgVpnClient) GetRestHttpRequestRxMsgCount() int64`

GetRestHttpRequestRxMsgCount returns the RestHttpRequestRxMsgCount field if non-nil, zero value otherwise.

### GetRestHttpRequestRxMsgCountOk

`func (o *MsgVpnClient) GetRestHttpRequestRxMsgCountOk() (*int64, bool)`

GetRestHttpRequestRxMsgCountOk returns a tuple with the RestHttpRequestRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpRequestRxMsgCount

`func (o *MsgVpnClient) SetRestHttpRequestRxMsgCount(v int64)`

SetRestHttpRequestRxMsgCount sets RestHttpRequestRxMsgCount field to given value.

### HasRestHttpRequestRxMsgCount

`func (o *MsgVpnClient) HasRestHttpRequestRxMsgCount() bool`

HasRestHttpRequestRxMsgCount returns a boolean if a field has been set.

### GetRestHttpRequestTxByteCount

`func (o *MsgVpnClient) GetRestHttpRequestTxByteCount() int64`

GetRestHttpRequestTxByteCount returns the RestHttpRequestTxByteCount field if non-nil, zero value otherwise.

### GetRestHttpRequestTxByteCountOk

`func (o *MsgVpnClient) GetRestHttpRequestTxByteCountOk() (*int64, bool)`

GetRestHttpRequestTxByteCountOk returns a tuple with the RestHttpRequestTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpRequestTxByteCount

`func (o *MsgVpnClient) SetRestHttpRequestTxByteCount(v int64)`

SetRestHttpRequestTxByteCount sets RestHttpRequestTxByteCount field to given value.

### HasRestHttpRequestTxByteCount

`func (o *MsgVpnClient) HasRestHttpRequestTxByteCount() bool`

HasRestHttpRequestTxByteCount returns a boolean if a field has been set.

### GetRestHttpRequestTxMsgCount

`func (o *MsgVpnClient) GetRestHttpRequestTxMsgCount() int64`

GetRestHttpRequestTxMsgCount returns the RestHttpRequestTxMsgCount field if non-nil, zero value otherwise.

### GetRestHttpRequestTxMsgCountOk

`func (o *MsgVpnClient) GetRestHttpRequestTxMsgCountOk() (*int64, bool)`

GetRestHttpRequestTxMsgCountOk returns a tuple with the RestHttpRequestTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpRequestTxMsgCount

`func (o *MsgVpnClient) SetRestHttpRequestTxMsgCount(v int64)`

SetRestHttpRequestTxMsgCount sets RestHttpRequestTxMsgCount field to given value.

### HasRestHttpRequestTxMsgCount

`func (o *MsgVpnClient) HasRestHttpRequestTxMsgCount() bool`

HasRestHttpRequestTxMsgCount returns a boolean if a field has been set.

### GetRestHttpResponseErrorRxMsgCount

`func (o *MsgVpnClient) GetRestHttpResponseErrorRxMsgCount() int64`

GetRestHttpResponseErrorRxMsgCount returns the RestHttpResponseErrorRxMsgCount field if non-nil, zero value otherwise.

### GetRestHttpResponseErrorRxMsgCountOk

`func (o *MsgVpnClient) GetRestHttpResponseErrorRxMsgCountOk() (*int64, bool)`

GetRestHttpResponseErrorRxMsgCountOk returns a tuple with the RestHttpResponseErrorRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpResponseErrorRxMsgCount

`func (o *MsgVpnClient) SetRestHttpResponseErrorRxMsgCount(v int64)`

SetRestHttpResponseErrorRxMsgCount sets RestHttpResponseErrorRxMsgCount field to given value.

### HasRestHttpResponseErrorRxMsgCount

`func (o *MsgVpnClient) HasRestHttpResponseErrorRxMsgCount() bool`

HasRestHttpResponseErrorRxMsgCount returns a boolean if a field has been set.

### GetRestHttpResponseErrorTxMsgCount

`func (o *MsgVpnClient) GetRestHttpResponseErrorTxMsgCount() int64`

GetRestHttpResponseErrorTxMsgCount returns the RestHttpResponseErrorTxMsgCount field if non-nil, zero value otherwise.

### GetRestHttpResponseErrorTxMsgCountOk

`func (o *MsgVpnClient) GetRestHttpResponseErrorTxMsgCountOk() (*int64, bool)`

GetRestHttpResponseErrorTxMsgCountOk returns a tuple with the RestHttpResponseErrorTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpResponseErrorTxMsgCount

`func (o *MsgVpnClient) SetRestHttpResponseErrorTxMsgCount(v int64)`

SetRestHttpResponseErrorTxMsgCount sets RestHttpResponseErrorTxMsgCount field to given value.

### HasRestHttpResponseErrorTxMsgCount

`func (o *MsgVpnClient) HasRestHttpResponseErrorTxMsgCount() bool`

HasRestHttpResponseErrorTxMsgCount returns a boolean if a field has been set.

### GetRestHttpResponseRxByteCount

`func (o *MsgVpnClient) GetRestHttpResponseRxByteCount() int64`

GetRestHttpResponseRxByteCount returns the RestHttpResponseRxByteCount field if non-nil, zero value otherwise.

### GetRestHttpResponseRxByteCountOk

`func (o *MsgVpnClient) GetRestHttpResponseRxByteCountOk() (*int64, bool)`

GetRestHttpResponseRxByteCountOk returns a tuple with the RestHttpResponseRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpResponseRxByteCount

`func (o *MsgVpnClient) SetRestHttpResponseRxByteCount(v int64)`

SetRestHttpResponseRxByteCount sets RestHttpResponseRxByteCount field to given value.

### HasRestHttpResponseRxByteCount

`func (o *MsgVpnClient) HasRestHttpResponseRxByteCount() bool`

HasRestHttpResponseRxByteCount returns a boolean if a field has been set.

### GetRestHttpResponseRxMsgCount

`func (o *MsgVpnClient) GetRestHttpResponseRxMsgCount() int64`

GetRestHttpResponseRxMsgCount returns the RestHttpResponseRxMsgCount field if non-nil, zero value otherwise.

### GetRestHttpResponseRxMsgCountOk

`func (o *MsgVpnClient) GetRestHttpResponseRxMsgCountOk() (*int64, bool)`

GetRestHttpResponseRxMsgCountOk returns a tuple with the RestHttpResponseRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpResponseRxMsgCount

`func (o *MsgVpnClient) SetRestHttpResponseRxMsgCount(v int64)`

SetRestHttpResponseRxMsgCount sets RestHttpResponseRxMsgCount field to given value.

### HasRestHttpResponseRxMsgCount

`func (o *MsgVpnClient) HasRestHttpResponseRxMsgCount() bool`

HasRestHttpResponseRxMsgCount returns a boolean if a field has been set.

### GetRestHttpResponseSuccessRxMsgCount

`func (o *MsgVpnClient) GetRestHttpResponseSuccessRxMsgCount() int64`

GetRestHttpResponseSuccessRxMsgCount returns the RestHttpResponseSuccessRxMsgCount field if non-nil, zero value otherwise.

### GetRestHttpResponseSuccessRxMsgCountOk

`func (o *MsgVpnClient) GetRestHttpResponseSuccessRxMsgCountOk() (*int64, bool)`

GetRestHttpResponseSuccessRxMsgCountOk returns a tuple with the RestHttpResponseSuccessRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpResponseSuccessRxMsgCount

`func (o *MsgVpnClient) SetRestHttpResponseSuccessRxMsgCount(v int64)`

SetRestHttpResponseSuccessRxMsgCount sets RestHttpResponseSuccessRxMsgCount field to given value.

### HasRestHttpResponseSuccessRxMsgCount

`func (o *MsgVpnClient) HasRestHttpResponseSuccessRxMsgCount() bool`

HasRestHttpResponseSuccessRxMsgCount returns a boolean if a field has been set.

### GetRestHttpResponseSuccessTxMsgCount

`func (o *MsgVpnClient) GetRestHttpResponseSuccessTxMsgCount() int64`

GetRestHttpResponseSuccessTxMsgCount returns the RestHttpResponseSuccessTxMsgCount field if non-nil, zero value otherwise.

### GetRestHttpResponseSuccessTxMsgCountOk

`func (o *MsgVpnClient) GetRestHttpResponseSuccessTxMsgCountOk() (*int64, bool)`

GetRestHttpResponseSuccessTxMsgCountOk returns a tuple with the RestHttpResponseSuccessTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpResponseSuccessTxMsgCount

`func (o *MsgVpnClient) SetRestHttpResponseSuccessTxMsgCount(v int64)`

SetRestHttpResponseSuccessTxMsgCount sets RestHttpResponseSuccessTxMsgCount field to given value.

### HasRestHttpResponseSuccessTxMsgCount

`func (o *MsgVpnClient) HasRestHttpResponseSuccessTxMsgCount() bool`

HasRestHttpResponseSuccessTxMsgCount returns a boolean if a field has been set.

### GetRestHttpResponseTimeoutRxMsgCount

`func (o *MsgVpnClient) GetRestHttpResponseTimeoutRxMsgCount() int64`

GetRestHttpResponseTimeoutRxMsgCount returns the RestHttpResponseTimeoutRxMsgCount field if non-nil, zero value otherwise.

### GetRestHttpResponseTimeoutRxMsgCountOk

`func (o *MsgVpnClient) GetRestHttpResponseTimeoutRxMsgCountOk() (*int64, bool)`

GetRestHttpResponseTimeoutRxMsgCountOk returns a tuple with the RestHttpResponseTimeoutRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpResponseTimeoutRxMsgCount

`func (o *MsgVpnClient) SetRestHttpResponseTimeoutRxMsgCount(v int64)`

SetRestHttpResponseTimeoutRxMsgCount sets RestHttpResponseTimeoutRxMsgCount field to given value.

### HasRestHttpResponseTimeoutRxMsgCount

`func (o *MsgVpnClient) HasRestHttpResponseTimeoutRxMsgCount() bool`

HasRestHttpResponseTimeoutRxMsgCount returns a boolean if a field has been set.

### GetRestHttpResponseTimeoutTxMsgCount

`func (o *MsgVpnClient) GetRestHttpResponseTimeoutTxMsgCount() int64`

GetRestHttpResponseTimeoutTxMsgCount returns the RestHttpResponseTimeoutTxMsgCount field if non-nil, zero value otherwise.

### GetRestHttpResponseTimeoutTxMsgCountOk

`func (o *MsgVpnClient) GetRestHttpResponseTimeoutTxMsgCountOk() (*int64, bool)`

GetRestHttpResponseTimeoutTxMsgCountOk returns a tuple with the RestHttpResponseTimeoutTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpResponseTimeoutTxMsgCount

`func (o *MsgVpnClient) SetRestHttpResponseTimeoutTxMsgCount(v int64)`

SetRestHttpResponseTimeoutTxMsgCount sets RestHttpResponseTimeoutTxMsgCount field to given value.

### HasRestHttpResponseTimeoutTxMsgCount

`func (o *MsgVpnClient) HasRestHttpResponseTimeoutTxMsgCount() bool`

HasRestHttpResponseTimeoutTxMsgCount returns a boolean if a field has been set.

### GetRestHttpResponseTxByteCount

`func (o *MsgVpnClient) GetRestHttpResponseTxByteCount() int64`

GetRestHttpResponseTxByteCount returns the RestHttpResponseTxByteCount field if non-nil, zero value otherwise.

### GetRestHttpResponseTxByteCountOk

`func (o *MsgVpnClient) GetRestHttpResponseTxByteCountOk() (*int64, bool)`

GetRestHttpResponseTxByteCountOk returns a tuple with the RestHttpResponseTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpResponseTxByteCount

`func (o *MsgVpnClient) SetRestHttpResponseTxByteCount(v int64)`

SetRestHttpResponseTxByteCount sets RestHttpResponseTxByteCount field to given value.

### HasRestHttpResponseTxByteCount

`func (o *MsgVpnClient) HasRestHttpResponseTxByteCount() bool`

HasRestHttpResponseTxByteCount returns a boolean if a field has been set.

### GetRestHttpResponseTxMsgCount

`func (o *MsgVpnClient) GetRestHttpResponseTxMsgCount() int64`

GetRestHttpResponseTxMsgCount returns the RestHttpResponseTxMsgCount field if non-nil, zero value otherwise.

### GetRestHttpResponseTxMsgCountOk

`func (o *MsgVpnClient) GetRestHttpResponseTxMsgCountOk() (*int64, bool)`

GetRestHttpResponseTxMsgCountOk returns a tuple with the RestHttpResponseTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestHttpResponseTxMsgCount

`func (o *MsgVpnClient) SetRestHttpResponseTxMsgCount(v int64)`

SetRestHttpResponseTxMsgCount sets RestHttpResponseTxMsgCount field to given value.

### HasRestHttpResponseTxMsgCount

`func (o *MsgVpnClient) HasRestHttpResponseTxMsgCount() bool`

HasRestHttpResponseTxMsgCount returns a boolean if a field has been set.

### GetRxByteCount

`func (o *MsgVpnClient) GetRxByteCount() int64`

GetRxByteCount returns the RxByteCount field if non-nil, zero value otherwise.

### GetRxByteCountOk

`func (o *MsgVpnClient) GetRxByteCountOk() (*int64, bool)`

GetRxByteCountOk returns a tuple with the RxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteCount

`func (o *MsgVpnClient) SetRxByteCount(v int64)`

SetRxByteCount sets RxByteCount field to given value.

### HasRxByteCount

`func (o *MsgVpnClient) HasRxByteCount() bool`

HasRxByteCount returns a boolean if a field has been set.

### GetRxByteRate

`func (o *MsgVpnClient) GetRxByteRate() int64`

GetRxByteRate returns the RxByteRate field if non-nil, zero value otherwise.

### GetRxByteRateOk

`func (o *MsgVpnClient) GetRxByteRateOk() (*int64, bool)`

GetRxByteRateOk returns a tuple with the RxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteRate

`func (o *MsgVpnClient) SetRxByteRate(v int64)`

SetRxByteRate sets RxByteRate field to given value.

### HasRxByteRate

`func (o *MsgVpnClient) HasRxByteRate() bool`

HasRxByteRate returns a boolean if a field has been set.

### GetRxDiscardedMsgCount

`func (o *MsgVpnClient) GetRxDiscardedMsgCount() int64`

GetRxDiscardedMsgCount returns the RxDiscardedMsgCount field if non-nil, zero value otherwise.

### GetRxDiscardedMsgCountOk

`func (o *MsgVpnClient) GetRxDiscardedMsgCountOk() (*int64, bool)`

GetRxDiscardedMsgCountOk returns a tuple with the RxDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDiscardedMsgCount

`func (o *MsgVpnClient) SetRxDiscardedMsgCount(v int64)`

SetRxDiscardedMsgCount sets RxDiscardedMsgCount field to given value.

### HasRxDiscardedMsgCount

`func (o *MsgVpnClient) HasRxDiscardedMsgCount() bool`

HasRxDiscardedMsgCount returns a boolean if a field has been set.

### GetRxMsgCount

`func (o *MsgVpnClient) GetRxMsgCount() int64`

GetRxMsgCount returns the RxMsgCount field if non-nil, zero value otherwise.

### GetRxMsgCountOk

`func (o *MsgVpnClient) GetRxMsgCountOk() (*int64, bool)`

GetRxMsgCountOk returns a tuple with the RxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgCount

`func (o *MsgVpnClient) SetRxMsgCount(v int64)`

SetRxMsgCount sets RxMsgCount field to given value.

### HasRxMsgCount

`func (o *MsgVpnClient) HasRxMsgCount() bool`

HasRxMsgCount returns a boolean if a field has been set.

### GetRxMsgRate

`func (o *MsgVpnClient) GetRxMsgRate() int64`

GetRxMsgRate returns the RxMsgRate field if non-nil, zero value otherwise.

### GetRxMsgRateOk

`func (o *MsgVpnClient) GetRxMsgRateOk() (*int64, bool)`

GetRxMsgRateOk returns a tuple with the RxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgRate

`func (o *MsgVpnClient) SetRxMsgRate(v int64)`

SetRxMsgRate sets RxMsgRate field to given value.

### HasRxMsgRate

`func (o *MsgVpnClient) HasRxMsgRate() bool`

HasRxMsgRate returns a boolean if a field has been set.

### GetScheduledDisconnectTime

`func (o *MsgVpnClient) GetScheduledDisconnectTime() int32`

GetScheduledDisconnectTime returns the ScheduledDisconnectTime field if non-nil, zero value otherwise.

### GetScheduledDisconnectTimeOk

`func (o *MsgVpnClient) GetScheduledDisconnectTimeOk() (*int32, bool)`

GetScheduledDisconnectTimeOk returns a tuple with the ScheduledDisconnectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDisconnectTime

`func (o *MsgVpnClient) SetScheduledDisconnectTime(v int32)`

SetScheduledDisconnectTime sets ScheduledDisconnectTime field to given value.

### HasScheduledDisconnectTime

`func (o *MsgVpnClient) HasScheduledDisconnectTime() bool`

HasScheduledDisconnectTime returns a boolean if a field has been set.

### GetSlowSubscriber

`func (o *MsgVpnClient) GetSlowSubscriber() bool`

GetSlowSubscriber returns the SlowSubscriber field if non-nil, zero value otherwise.

### GetSlowSubscriberOk

`func (o *MsgVpnClient) GetSlowSubscriberOk() (*bool, bool)`

GetSlowSubscriberOk returns a tuple with the SlowSubscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowSubscriber

`func (o *MsgVpnClient) SetSlowSubscriber(v bool)`

SetSlowSubscriber sets SlowSubscriber field to given value.

### HasSlowSubscriber

`func (o *MsgVpnClient) HasSlowSubscriber() bool`

HasSlowSubscriber returns a boolean if a field has been set.

### GetSoftwareDate

`func (o *MsgVpnClient) GetSoftwareDate() string`

GetSoftwareDate returns the SoftwareDate field if non-nil, zero value otherwise.

### GetSoftwareDateOk

`func (o *MsgVpnClient) GetSoftwareDateOk() (*string, bool)`

GetSoftwareDateOk returns a tuple with the SoftwareDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDate

`func (o *MsgVpnClient) SetSoftwareDate(v string)`

SetSoftwareDate sets SoftwareDate field to given value.

### HasSoftwareDate

`func (o *MsgVpnClient) HasSoftwareDate() bool`

HasSoftwareDate returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *MsgVpnClient) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *MsgVpnClient) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *MsgVpnClient) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *MsgVpnClient) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetTlsCipherDescription

`func (o *MsgVpnClient) GetTlsCipherDescription() string`

GetTlsCipherDescription returns the TlsCipherDescription field if non-nil, zero value otherwise.

### GetTlsCipherDescriptionOk

`func (o *MsgVpnClient) GetTlsCipherDescriptionOk() (*string, bool)`

GetTlsCipherDescriptionOk returns a tuple with the TlsCipherDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCipherDescription

`func (o *MsgVpnClient) SetTlsCipherDescription(v string)`

SetTlsCipherDescription sets TlsCipherDescription field to given value.

### HasTlsCipherDescription

`func (o *MsgVpnClient) HasTlsCipherDescription() bool`

HasTlsCipherDescription returns a boolean if a field has been set.

### GetTlsDowngradedToPlainText

`func (o *MsgVpnClient) GetTlsDowngradedToPlainText() bool`

GetTlsDowngradedToPlainText returns the TlsDowngradedToPlainText field if non-nil, zero value otherwise.

### GetTlsDowngradedToPlainTextOk

`func (o *MsgVpnClient) GetTlsDowngradedToPlainTextOk() (*bool, bool)`

GetTlsDowngradedToPlainTextOk returns a tuple with the TlsDowngradedToPlainText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsDowngradedToPlainText

`func (o *MsgVpnClient) SetTlsDowngradedToPlainText(v bool)`

SetTlsDowngradedToPlainText sets TlsDowngradedToPlainText field to given value.

### HasTlsDowngradedToPlainText

`func (o *MsgVpnClient) HasTlsDowngradedToPlainText() bool`

HasTlsDowngradedToPlainText returns a boolean if a field has been set.

### GetTlsVersion

`func (o *MsgVpnClient) GetTlsVersion() string`

GetTlsVersion returns the TlsVersion field if non-nil, zero value otherwise.

### GetTlsVersionOk

`func (o *MsgVpnClient) GetTlsVersionOk() (*string, bool)`

GetTlsVersionOk returns a tuple with the TlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersion

`func (o *MsgVpnClient) SetTlsVersion(v string)`

SetTlsVersion sets TlsVersion field to given value.

### HasTlsVersion

`func (o *MsgVpnClient) HasTlsVersion() bool`

HasTlsVersion returns a boolean if a field has been set.

### GetTopicParseErrorRxDiscardedMsgCount

`func (o *MsgVpnClient) GetTopicParseErrorRxDiscardedMsgCount() int64`

GetTopicParseErrorRxDiscardedMsgCount returns the TopicParseErrorRxDiscardedMsgCount field if non-nil, zero value otherwise.

### GetTopicParseErrorRxDiscardedMsgCountOk

`func (o *MsgVpnClient) GetTopicParseErrorRxDiscardedMsgCountOk() (*int64, bool)`

GetTopicParseErrorRxDiscardedMsgCountOk returns a tuple with the TopicParseErrorRxDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicParseErrorRxDiscardedMsgCount

`func (o *MsgVpnClient) SetTopicParseErrorRxDiscardedMsgCount(v int64)`

SetTopicParseErrorRxDiscardedMsgCount sets TopicParseErrorRxDiscardedMsgCount field to given value.

### HasTopicParseErrorRxDiscardedMsgCount

`func (o *MsgVpnClient) HasTopicParseErrorRxDiscardedMsgCount() bool`

HasTopicParseErrorRxDiscardedMsgCount returns a boolean if a field has been set.

### GetTxByteCount

`func (o *MsgVpnClient) GetTxByteCount() int64`

GetTxByteCount returns the TxByteCount field if non-nil, zero value otherwise.

### GetTxByteCountOk

`func (o *MsgVpnClient) GetTxByteCountOk() (*int64, bool)`

GetTxByteCountOk returns a tuple with the TxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteCount

`func (o *MsgVpnClient) SetTxByteCount(v int64)`

SetTxByteCount sets TxByteCount field to given value.

### HasTxByteCount

`func (o *MsgVpnClient) HasTxByteCount() bool`

HasTxByteCount returns a boolean if a field has been set.

### GetTxByteRate

`func (o *MsgVpnClient) GetTxByteRate() int64`

GetTxByteRate returns the TxByteRate field if non-nil, zero value otherwise.

### GetTxByteRateOk

`func (o *MsgVpnClient) GetTxByteRateOk() (*int64, bool)`

GetTxByteRateOk returns a tuple with the TxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteRate

`func (o *MsgVpnClient) SetTxByteRate(v int64)`

SetTxByteRate sets TxByteRate field to given value.

### HasTxByteRate

`func (o *MsgVpnClient) HasTxByteRate() bool`

HasTxByteRate returns a boolean if a field has been set.

### GetTxDiscardedMsgCount

`func (o *MsgVpnClient) GetTxDiscardedMsgCount() int64`

GetTxDiscardedMsgCount returns the TxDiscardedMsgCount field if non-nil, zero value otherwise.

### GetTxDiscardedMsgCountOk

`func (o *MsgVpnClient) GetTxDiscardedMsgCountOk() (*int64, bool)`

GetTxDiscardedMsgCountOk returns a tuple with the TxDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDiscardedMsgCount

`func (o *MsgVpnClient) SetTxDiscardedMsgCount(v int64)`

SetTxDiscardedMsgCount sets TxDiscardedMsgCount field to given value.

### HasTxDiscardedMsgCount

`func (o *MsgVpnClient) HasTxDiscardedMsgCount() bool`

HasTxDiscardedMsgCount returns a boolean if a field has been set.

### GetTxMsgCount

`func (o *MsgVpnClient) GetTxMsgCount() int64`

GetTxMsgCount returns the TxMsgCount field if non-nil, zero value otherwise.

### GetTxMsgCountOk

`func (o *MsgVpnClient) GetTxMsgCountOk() (*int64, bool)`

GetTxMsgCountOk returns a tuple with the TxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgCount

`func (o *MsgVpnClient) SetTxMsgCount(v int64)`

SetTxMsgCount sets TxMsgCount field to given value.

### HasTxMsgCount

`func (o *MsgVpnClient) HasTxMsgCount() bool`

HasTxMsgCount returns a boolean if a field has been set.

### GetTxMsgRate

`func (o *MsgVpnClient) GetTxMsgRate() int64`

GetTxMsgRate returns the TxMsgRate field if non-nil, zero value otherwise.

### GetTxMsgRateOk

`func (o *MsgVpnClient) GetTxMsgRateOk() (*int64, bool)`

GetTxMsgRateOk returns a tuple with the TxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgRate

`func (o *MsgVpnClient) SetTxMsgRate(v int64)`

SetTxMsgRate sets TxMsgRate field to given value.

### HasTxMsgRate

`func (o *MsgVpnClient) HasTxMsgRate() bool`

HasTxMsgRate returns a boolean if a field has been set.

### GetUptime

`func (o *MsgVpnClient) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MsgVpnClient) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MsgVpnClient) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MsgVpnClient) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUser

`func (o *MsgVpnClient) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MsgVpnClient) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MsgVpnClient) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *MsgVpnClient) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVirtualRouter

`func (o *MsgVpnClient) GetVirtualRouter() string`

GetVirtualRouter returns the VirtualRouter field if non-nil, zero value otherwise.

### GetVirtualRouterOk

`func (o *MsgVpnClient) GetVirtualRouterOk() (*string, bool)`

GetVirtualRouterOk returns a tuple with the VirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouter

`func (o *MsgVpnClient) SetVirtualRouter(v string)`

SetVirtualRouter sets VirtualRouter field to given value.

### HasVirtualRouter

`func (o *MsgVpnClient) HasVirtualRouter() bool`

HasVirtualRouter returns a boolean if a field has been set.

### GetWebInactiveTimeout

`func (o *MsgVpnClient) GetWebInactiveTimeout() int32`

GetWebInactiveTimeout returns the WebInactiveTimeout field if non-nil, zero value otherwise.

### GetWebInactiveTimeoutOk

`func (o *MsgVpnClient) GetWebInactiveTimeoutOk() (*int32, bool)`

GetWebInactiveTimeoutOk returns a tuple with the WebInactiveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebInactiveTimeout

`func (o *MsgVpnClient) SetWebInactiveTimeout(v int32)`

SetWebInactiveTimeout sets WebInactiveTimeout field to given value.

### HasWebInactiveTimeout

`func (o *MsgVpnClient) HasWebInactiveTimeout() bool`

HasWebInactiveTimeout returns a boolean if a field has been set.

### GetWebMaxPayload

`func (o *MsgVpnClient) GetWebMaxPayload() int64`

GetWebMaxPayload returns the WebMaxPayload field if non-nil, zero value otherwise.

### GetWebMaxPayloadOk

`func (o *MsgVpnClient) GetWebMaxPayloadOk() (*int64, bool)`

GetWebMaxPayloadOk returns a tuple with the WebMaxPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMaxPayload

`func (o *MsgVpnClient) SetWebMaxPayload(v int64)`

SetWebMaxPayload sets WebMaxPayload field to given value.

### HasWebMaxPayload

`func (o *MsgVpnClient) HasWebMaxPayload() bool`

HasWebMaxPayload returns a boolean if a field has been set.

### GetWebParseErrorRxDiscardedMsgCount

`func (o *MsgVpnClient) GetWebParseErrorRxDiscardedMsgCount() int64`

GetWebParseErrorRxDiscardedMsgCount returns the WebParseErrorRxDiscardedMsgCount field if non-nil, zero value otherwise.

### GetWebParseErrorRxDiscardedMsgCountOk

`func (o *MsgVpnClient) GetWebParseErrorRxDiscardedMsgCountOk() (*int64, bool)`

GetWebParseErrorRxDiscardedMsgCountOk returns a tuple with the WebParseErrorRxDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebParseErrorRxDiscardedMsgCount

`func (o *MsgVpnClient) SetWebParseErrorRxDiscardedMsgCount(v int64)`

SetWebParseErrorRxDiscardedMsgCount sets WebParseErrorRxDiscardedMsgCount field to given value.

### HasWebParseErrorRxDiscardedMsgCount

`func (o *MsgVpnClient) HasWebParseErrorRxDiscardedMsgCount() bool`

HasWebParseErrorRxDiscardedMsgCount returns a boolean if a field has been set.

### GetWebRemainingTimeout

`func (o *MsgVpnClient) GetWebRemainingTimeout() int32`

GetWebRemainingTimeout returns the WebRemainingTimeout field if non-nil, zero value otherwise.

### GetWebRemainingTimeoutOk

`func (o *MsgVpnClient) GetWebRemainingTimeoutOk() (*int32, bool)`

GetWebRemainingTimeoutOk returns a tuple with the WebRemainingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRemainingTimeout

`func (o *MsgVpnClient) SetWebRemainingTimeout(v int32)`

SetWebRemainingTimeout sets WebRemainingTimeout field to given value.

### HasWebRemainingTimeout

`func (o *MsgVpnClient) HasWebRemainingTimeout() bool`

HasWebRemainingTimeout returns a boolean if a field has been set.

### GetWebRxByteCount

`func (o *MsgVpnClient) GetWebRxByteCount() int64`

GetWebRxByteCount returns the WebRxByteCount field if non-nil, zero value otherwise.

### GetWebRxByteCountOk

`func (o *MsgVpnClient) GetWebRxByteCountOk() (*int64, bool)`

GetWebRxByteCountOk returns a tuple with the WebRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRxByteCount

`func (o *MsgVpnClient) SetWebRxByteCount(v int64)`

SetWebRxByteCount sets WebRxByteCount field to given value.

### HasWebRxByteCount

`func (o *MsgVpnClient) HasWebRxByteCount() bool`

HasWebRxByteCount returns a boolean if a field has been set.

### GetWebRxEncoding

`func (o *MsgVpnClient) GetWebRxEncoding() string`

GetWebRxEncoding returns the WebRxEncoding field if non-nil, zero value otherwise.

### GetWebRxEncodingOk

`func (o *MsgVpnClient) GetWebRxEncodingOk() (*string, bool)`

GetWebRxEncodingOk returns a tuple with the WebRxEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRxEncoding

`func (o *MsgVpnClient) SetWebRxEncoding(v string)`

SetWebRxEncoding sets WebRxEncoding field to given value.

### HasWebRxEncoding

`func (o *MsgVpnClient) HasWebRxEncoding() bool`

HasWebRxEncoding returns a boolean if a field has been set.

### GetWebRxMsgCount

`func (o *MsgVpnClient) GetWebRxMsgCount() int64`

GetWebRxMsgCount returns the WebRxMsgCount field if non-nil, zero value otherwise.

### GetWebRxMsgCountOk

`func (o *MsgVpnClient) GetWebRxMsgCountOk() (*int64, bool)`

GetWebRxMsgCountOk returns a tuple with the WebRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRxMsgCount

`func (o *MsgVpnClient) SetWebRxMsgCount(v int64)`

SetWebRxMsgCount sets WebRxMsgCount field to given value.

### HasWebRxMsgCount

`func (o *MsgVpnClient) HasWebRxMsgCount() bool`

HasWebRxMsgCount returns a boolean if a field has been set.

### GetWebRxProtocol

`func (o *MsgVpnClient) GetWebRxProtocol() string`

GetWebRxProtocol returns the WebRxProtocol field if non-nil, zero value otherwise.

### GetWebRxProtocolOk

`func (o *MsgVpnClient) GetWebRxProtocolOk() (*string, bool)`

GetWebRxProtocolOk returns a tuple with the WebRxProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRxProtocol

`func (o *MsgVpnClient) SetWebRxProtocol(v string)`

SetWebRxProtocol sets WebRxProtocol field to given value.

### HasWebRxProtocol

`func (o *MsgVpnClient) HasWebRxProtocol() bool`

HasWebRxProtocol returns a boolean if a field has been set.

### GetWebRxRequestCount

`func (o *MsgVpnClient) GetWebRxRequestCount() int64`

GetWebRxRequestCount returns the WebRxRequestCount field if non-nil, zero value otherwise.

### GetWebRxRequestCountOk

`func (o *MsgVpnClient) GetWebRxRequestCountOk() (*int64, bool)`

GetWebRxRequestCountOk returns a tuple with the WebRxRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRxRequestCount

`func (o *MsgVpnClient) SetWebRxRequestCount(v int64)`

SetWebRxRequestCount sets WebRxRequestCount field to given value.

### HasWebRxRequestCount

`func (o *MsgVpnClient) HasWebRxRequestCount() bool`

HasWebRxRequestCount returns a boolean if a field has been set.

### GetWebRxResponseCount

`func (o *MsgVpnClient) GetWebRxResponseCount() int64`

GetWebRxResponseCount returns the WebRxResponseCount field if non-nil, zero value otherwise.

### GetWebRxResponseCountOk

`func (o *MsgVpnClient) GetWebRxResponseCountOk() (*int64, bool)`

GetWebRxResponseCountOk returns a tuple with the WebRxResponseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRxResponseCount

`func (o *MsgVpnClient) SetWebRxResponseCount(v int64)`

SetWebRxResponseCount sets WebRxResponseCount field to given value.

### HasWebRxResponseCount

`func (o *MsgVpnClient) HasWebRxResponseCount() bool`

HasWebRxResponseCount returns a boolean if a field has been set.

### GetWebRxTcpState

`func (o *MsgVpnClient) GetWebRxTcpState() string`

GetWebRxTcpState returns the WebRxTcpState field if non-nil, zero value otherwise.

### GetWebRxTcpStateOk

`func (o *MsgVpnClient) GetWebRxTcpStateOk() (*string, bool)`

GetWebRxTcpStateOk returns a tuple with the WebRxTcpState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRxTcpState

`func (o *MsgVpnClient) SetWebRxTcpState(v string)`

SetWebRxTcpState sets WebRxTcpState field to given value.

### HasWebRxTcpState

`func (o *MsgVpnClient) HasWebRxTcpState() bool`

HasWebRxTcpState returns a boolean if a field has been set.

### GetWebRxTlsCipherDescription

`func (o *MsgVpnClient) GetWebRxTlsCipherDescription() string`

GetWebRxTlsCipherDescription returns the WebRxTlsCipherDescription field if non-nil, zero value otherwise.

### GetWebRxTlsCipherDescriptionOk

`func (o *MsgVpnClient) GetWebRxTlsCipherDescriptionOk() (*string, bool)`

GetWebRxTlsCipherDescriptionOk returns a tuple with the WebRxTlsCipherDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRxTlsCipherDescription

`func (o *MsgVpnClient) SetWebRxTlsCipherDescription(v string)`

SetWebRxTlsCipherDescription sets WebRxTlsCipherDescription field to given value.

### HasWebRxTlsCipherDescription

`func (o *MsgVpnClient) HasWebRxTlsCipherDescription() bool`

HasWebRxTlsCipherDescription returns a boolean if a field has been set.

### GetWebRxTlsVersion

`func (o *MsgVpnClient) GetWebRxTlsVersion() string`

GetWebRxTlsVersion returns the WebRxTlsVersion field if non-nil, zero value otherwise.

### GetWebRxTlsVersionOk

`func (o *MsgVpnClient) GetWebRxTlsVersionOk() (*string, bool)`

GetWebRxTlsVersionOk returns a tuple with the WebRxTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRxTlsVersion

`func (o *MsgVpnClient) SetWebRxTlsVersion(v string)`

SetWebRxTlsVersion sets WebRxTlsVersion field to given value.

### HasWebRxTlsVersion

`func (o *MsgVpnClient) HasWebRxTlsVersion() bool`

HasWebRxTlsVersion returns a boolean if a field has been set.

### GetWebSessionId

`func (o *MsgVpnClient) GetWebSessionId() string`

GetWebSessionId returns the WebSessionId field if non-nil, zero value otherwise.

### GetWebSessionIdOk

`func (o *MsgVpnClient) GetWebSessionIdOk() (*string, bool)`

GetWebSessionIdOk returns a tuple with the WebSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSessionId

`func (o *MsgVpnClient) SetWebSessionId(v string)`

SetWebSessionId sets WebSessionId field to given value.

### HasWebSessionId

`func (o *MsgVpnClient) HasWebSessionId() bool`

HasWebSessionId returns a boolean if a field has been set.

### GetWebTxByteCount

`func (o *MsgVpnClient) GetWebTxByteCount() int64`

GetWebTxByteCount returns the WebTxByteCount field if non-nil, zero value otherwise.

### GetWebTxByteCountOk

`func (o *MsgVpnClient) GetWebTxByteCountOk() (*int64, bool)`

GetWebTxByteCountOk returns a tuple with the WebTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTxByteCount

`func (o *MsgVpnClient) SetWebTxByteCount(v int64)`

SetWebTxByteCount sets WebTxByteCount field to given value.

### HasWebTxByteCount

`func (o *MsgVpnClient) HasWebTxByteCount() bool`

HasWebTxByteCount returns a boolean if a field has been set.

### GetWebTxEncoding

`func (o *MsgVpnClient) GetWebTxEncoding() string`

GetWebTxEncoding returns the WebTxEncoding field if non-nil, zero value otherwise.

### GetWebTxEncodingOk

`func (o *MsgVpnClient) GetWebTxEncodingOk() (*string, bool)`

GetWebTxEncodingOk returns a tuple with the WebTxEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTxEncoding

`func (o *MsgVpnClient) SetWebTxEncoding(v string)`

SetWebTxEncoding sets WebTxEncoding field to given value.

### HasWebTxEncoding

`func (o *MsgVpnClient) HasWebTxEncoding() bool`

HasWebTxEncoding returns a boolean if a field has been set.

### GetWebTxMsgCount

`func (o *MsgVpnClient) GetWebTxMsgCount() int64`

GetWebTxMsgCount returns the WebTxMsgCount field if non-nil, zero value otherwise.

### GetWebTxMsgCountOk

`func (o *MsgVpnClient) GetWebTxMsgCountOk() (*int64, bool)`

GetWebTxMsgCountOk returns a tuple with the WebTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTxMsgCount

`func (o *MsgVpnClient) SetWebTxMsgCount(v int64)`

SetWebTxMsgCount sets WebTxMsgCount field to given value.

### HasWebTxMsgCount

`func (o *MsgVpnClient) HasWebTxMsgCount() bool`

HasWebTxMsgCount returns a boolean if a field has been set.

### GetWebTxProtocol

`func (o *MsgVpnClient) GetWebTxProtocol() string`

GetWebTxProtocol returns the WebTxProtocol field if non-nil, zero value otherwise.

### GetWebTxProtocolOk

`func (o *MsgVpnClient) GetWebTxProtocolOk() (*string, bool)`

GetWebTxProtocolOk returns a tuple with the WebTxProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTxProtocol

`func (o *MsgVpnClient) SetWebTxProtocol(v string)`

SetWebTxProtocol sets WebTxProtocol field to given value.

### HasWebTxProtocol

`func (o *MsgVpnClient) HasWebTxProtocol() bool`

HasWebTxProtocol returns a boolean if a field has been set.

### GetWebTxRequestCount

`func (o *MsgVpnClient) GetWebTxRequestCount() int64`

GetWebTxRequestCount returns the WebTxRequestCount field if non-nil, zero value otherwise.

### GetWebTxRequestCountOk

`func (o *MsgVpnClient) GetWebTxRequestCountOk() (*int64, bool)`

GetWebTxRequestCountOk returns a tuple with the WebTxRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTxRequestCount

`func (o *MsgVpnClient) SetWebTxRequestCount(v int64)`

SetWebTxRequestCount sets WebTxRequestCount field to given value.

### HasWebTxRequestCount

`func (o *MsgVpnClient) HasWebTxRequestCount() bool`

HasWebTxRequestCount returns a boolean if a field has been set.

### GetWebTxResponseCount

`func (o *MsgVpnClient) GetWebTxResponseCount() int64`

GetWebTxResponseCount returns the WebTxResponseCount field if non-nil, zero value otherwise.

### GetWebTxResponseCountOk

`func (o *MsgVpnClient) GetWebTxResponseCountOk() (*int64, bool)`

GetWebTxResponseCountOk returns a tuple with the WebTxResponseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTxResponseCount

`func (o *MsgVpnClient) SetWebTxResponseCount(v int64)`

SetWebTxResponseCount sets WebTxResponseCount field to given value.

### HasWebTxResponseCount

`func (o *MsgVpnClient) HasWebTxResponseCount() bool`

HasWebTxResponseCount returns a boolean if a field has been set.

### GetWebTxTcpState

`func (o *MsgVpnClient) GetWebTxTcpState() string`

GetWebTxTcpState returns the WebTxTcpState field if non-nil, zero value otherwise.

### GetWebTxTcpStateOk

`func (o *MsgVpnClient) GetWebTxTcpStateOk() (*string, bool)`

GetWebTxTcpStateOk returns a tuple with the WebTxTcpState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTxTcpState

`func (o *MsgVpnClient) SetWebTxTcpState(v string)`

SetWebTxTcpState sets WebTxTcpState field to given value.

### HasWebTxTcpState

`func (o *MsgVpnClient) HasWebTxTcpState() bool`

HasWebTxTcpState returns a boolean if a field has been set.

### GetWebTxTlsCipherDescription

`func (o *MsgVpnClient) GetWebTxTlsCipherDescription() string`

GetWebTxTlsCipherDescription returns the WebTxTlsCipherDescription field if non-nil, zero value otherwise.

### GetWebTxTlsCipherDescriptionOk

`func (o *MsgVpnClient) GetWebTxTlsCipherDescriptionOk() (*string, bool)`

GetWebTxTlsCipherDescriptionOk returns a tuple with the WebTxTlsCipherDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTxTlsCipherDescription

`func (o *MsgVpnClient) SetWebTxTlsCipherDescription(v string)`

SetWebTxTlsCipherDescription sets WebTxTlsCipherDescription field to given value.

### HasWebTxTlsCipherDescription

`func (o *MsgVpnClient) HasWebTxTlsCipherDescription() bool`

HasWebTxTlsCipherDescription returns a boolean if a field has been set.

### GetWebTxTlsVersion

`func (o *MsgVpnClient) GetWebTxTlsVersion() string`

GetWebTxTlsVersion returns the WebTxTlsVersion field if non-nil, zero value otherwise.

### GetWebTxTlsVersionOk

`func (o *MsgVpnClient) GetWebTxTlsVersionOk() (*string, bool)`

GetWebTxTlsVersionOk returns a tuple with the WebTxTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTxTlsVersion

`func (o *MsgVpnClient) SetWebTxTlsVersion(v string)`

SetWebTxTlsVersion sets WebTxTlsVersion field to given value.

### HasWebTxTlsVersion

`func (o *MsgVpnClient) HasWebTxTlsVersion() bool`

HasWebTxTlsVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


