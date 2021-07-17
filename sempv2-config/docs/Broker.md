# Broker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthClientCertRevocationCheckMode** | Pointer to **string** | The client certificate revocation checking mode used when a client authenticates with a client certificate. The default value is &#x60;\&quot;none\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;none\&quot; - Do not perform any certificate revocation checking. \&quot;ocsp\&quot; - Use the Open Certificate Status Protcol (OCSP) for certificate revocation checking. \&quot;crl\&quot; - Use Certificate Revocation Lists (CRL) for certificate revocation checking. \&quot;ocsp-crl\&quot; - Use OCSP first, but if OCSP fails to return an unambiguous result, then check via CRL. &lt;/pre&gt;  | [optional] 
**GuaranteedMsgingEnabled** | Pointer to **bool** | Enable or disable Guaranteed Messaging. The default value is &#x60;false&#x60;. Available since 2.18. | [optional] 
**GuaranteedMsgingEventCacheUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**GuaranteedMsgingEventDeliveredUnackedThreshold** | Pointer to [**EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] 
**GuaranteedMsgingEventDiskUsageThreshold** | Pointer to [**EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] 
**GuaranteedMsgingEventEgressFlowCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**GuaranteedMsgingEventEndpointCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**GuaranteedMsgingEventIngressFlowCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**GuaranteedMsgingEventMsgCountThreshold** | Pointer to [**EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] 
**GuaranteedMsgingEventMsgSpoolFileCountThreshold** | Pointer to [**EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] 
**GuaranteedMsgingEventMsgSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**GuaranteedMsgingEventTransactedSessionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**GuaranteedMsgingEventTransactedSessionResourceCountThreshold** | Pointer to [**EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] 
**GuaranteedMsgingEventTransactionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**GuaranteedMsgingMaxCacheUsage** | Pointer to **int32** | Guaranteed messaging cache usage limit. Expressed as a maximum percentage of the NAB&#39;s egress queueing. resources that the guaranteed message cache is allowed to use. The default value is &#x60;10&#x60;. Available since 2.18. | [optional] 
**GuaranteedMsgingMaxMsgSpoolUsage** | Pointer to **int64** | The maximum total message spool usage allowed across all VPNs on this broker, in megabytes. Recommendation: the maximum value should be less than 90% of the disk space allocated for the guaranteed message spool. The default value is &#x60;60000&#x60;. Available since 2.18. | [optional] 
**GuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout** | Pointer to **int64** | The maximum time, in milliseconds, that can be tolerated for remote acknowledgement of synchronization messages before which the remote system will be considered out of sync. The default value is &#x60;10000&#x60;. Available since 2.18. | [optional] 
**GuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout** | Pointer to **int64** | The maximum time, in milliseconds, that can be tolerated for remote disk writes before which the remote system will be considered out of sync. The default value is &#x60;10000&#x60;. Available since 2.18. | [optional] 
**GuaranteedMsgingTransactionReplicationCompatibilityMode** | Pointer to **string** | The replication compatibility mode for the router. The default value is &#x60;\&quot;legacy\&quot;&#x60;. The allowed values and their meaning are:\&quot;legacy\&quot; - All transactions originated by clients are replicated to the standby site without using transactions.\&quot;transacted\&quot; - All transactions originated by clients are replicated to the standby site using transactions. The default value is &#x60;\&quot;legacy\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;legacy\&quot; - All transactions originated by clients are replicated to the standby site without using transactions. \&quot;transacted\&quot; - All transactions originated by clients are replicated to the standby site using transactions. &lt;/pre&gt;  Available since 2.18. | [optional] 
**ServiceAmqpEnabled** | Pointer to **bool** | Enable or disable the AMQP service. When disabled new AMQP Clients may not connect through the global or per-VPN AMQP listen-ports, and all currently connected AMQP Clients are immediately disconnected. The default value is &#x60;false&#x60;. Available since 2.17. | [optional] 
**ServiceAmqpTlsListenPort** | Pointer to **int64** | TCP port number that AMQP clients can use to connect to the broker using raw TCP over TLS. The default value is &#x60;0&#x60;. Available since 2.17. | [optional] 
**ServiceEventConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**ServiceHealthCheckEnabled** | Pointer to **bool** | Enable or disable the health-check service. The default value is &#x60;false&#x60;. Available since 2.17. | [optional] 
**ServiceHealthCheckListenPort** | Pointer to **int64** | The port number for the health-check service. The port must be unique across the message backbone. The health-check service must be disabled to change the port. The default value is &#x60;5550&#x60;. Available since 2.17. | [optional] 
**ServiceMateLinkEnabled** | Pointer to **bool** | Enable or disable the mate-link service. The default value is &#x60;true&#x60;. Available since 2.17. | [optional] 
**ServiceMateLinkListenPort** | Pointer to **int64** | The port number for the mate-link service. The port must be unique across the message backbone. The mate-link service must be disabled to change the port. The default value is &#x60;8741&#x60;. Available since 2.17. | [optional] 
**ServiceMqttEnabled** | Pointer to **bool** | Enable or disable the MQTT service. When disabled new MQTT Clients may not connect through the per-VPN MQTT listen-ports, and all currently connected MQTT Clients are immediately disconnected. The default value is &#x60;false&#x60;. Available since 2.17. | [optional] 
**ServiceMsgBackboneEnabled** | Pointer to **bool** | Enable or disable the msg-backbone service. When disabled new Clients may not connect through global or per-VPN listen-ports, and all currently connected Clients are immediately disconnected. The default value is &#x60;true&#x60;. Available since 2.17. | [optional] 
**ServiceRedundancyEnabled** | Pointer to **bool** | Enable or disable the redundancy service. The default value is &#x60;true&#x60;. Available since 2.17. | [optional] 
**ServiceRedundancyFirstListenPort** | Pointer to **int64** | The first listen-port used for the redundancy service. Redundancy uses this port and the subsequent 2 ports. These port must be unique across the message backbone. The redundancy service must be disabled to change this port. The default value is &#x60;8300&#x60;. Available since 2.17. | [optional] 
**ServiceRestEventOutgoingConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**ServiceRestIncomingEnabled** | Pointer to **bool** | Enable or disable the REST service incoming connections on the router. The default value is &#x60;false&#x60;. Available since 2.17. | [optional] 
**ServiceRestOutgoingEnabled** | Pointer to **bool** | Enable or disable the REST service outgoing connections on the router. The default value is &#x60;false&#x60;. Available since 2.17. | [optional] 
**ServiceSempLegacyTimeoutEnabled** | Pointer to **bool** | Enable or disable extended SEMP timeouts for paged GETs. When a request times out, it returns the current page of content, even if the page is not full.  When enabled, the timeout is 60 seconds. When disabled, the timeout is 5 seconds.  The recommended setting is disabled (no legacy-timeout).  This parameter is intended as a temporary workaround to be used until SEMP clients can handle short pages.  This setting will be removed in a future release. The default value is &#x60;false&#x60;. Available since 2.18. | [optional] 
**ServiceSempPlainTextEnabled** | Pointer to **bool** | Enable or disable plain-text SEMP service. The default value is &#x60;true&#x60;. Available since 2.17. | [optional] 
**ServiceSempPlainTextListenPort** | Pointer to **int64** | The TCP port for plain-text SEMP client connections. The default value is &#x60;80&#x60;. Available since 2.17. | [optional] 
**ServiceSempSessionIdleTimeout** | Pointer to **int32** | The session idle timeout, in minutes. Sessions will be invalidated if there is no activity in this period of time. The default value is &#x60;15&#x60;. Available since 2.21. | [optional] 
**ServiceSempSessionMaxLifetime** | Pointer to **int32** | The maximum lifetime of a session, in minutes. Sessions will be invalidated after this period of time, regardless of activity. The default value is &#x60;43200&#x60;. Available since 2.21. | [optional] 
**ServiceSempTlsEnabled** | Pointer to **bool** | Enable or disable TLS SEMP service. The default value is &#x60;true&#x60;. Available since 2.17. | [optional] 
**ServiceSempTlsListenPort** | Pointer to **int64** | The TCP port for TLS SEMP client connections. The default value is &#x60;1943&#x60;. Available since 2.17. | [optional] 
**ServiceSmfCompressionListenPort** | Pointer to **int64** | TCP port number that SMF clients can use to connect to the broker using raw compression TCP. The default value is &#x60;55003&#x60;. Available since 2.17. | [optional] 
**ServiceSmfEnabled** | Pointer to **bool** | Enable or disable the SMF service. When disabled new SMF Clients may not connect through the global listen-ports, and all currently connected SMF Clients are immediately disconnected. The default value is &#x60;true&#x60;. Available since 2.17. | [optional] 
**ServiceSmfEventConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**ServiceSmfPlainTextListenPort** | Pointer to **int64** | TCP port number that SMF clients can use to connect to the broker using raw TCP. The default value is &#x60;55555&#x60;. Available since 2.17. | [optional] 
**ServiceSmfRoutingControlListenPort** | Pointer to **int64** | TCP port number that SMF clients can use to connect to the broker using raw routing control TCP. The default value is &#x60;55556&#x60;. Available since 2.17. | [optional] 
**ServiceSmfTlsListenPort** | Pointer to **int64** | TCP port number that SMF clients can use to connect to the broker using raw TCP over TLS. The default value is &#x60;55443&#x60;. Available since 2.17. | [optional] 
**ServiceTlsEventConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**ServiceWebTransportEnabled** | Pointer to **bool** | Enable or disable the web-transport service. When disabled new web-transport Clients may not connect through the global listen-ports, and all currently connected web-transport Clients are immediately disconnected. The default value is &#x60;false&#x60;. Available since 2.17. | [optional] 
**ServiceWebTransportPlainTextListenPort** | Pointer to **int64** | The TCP port for plain-text WEB client connections. The default value is &#x60;8008&#x60;. Available since 2.17. | [optional] 
**ServiceWebTransportTlsListenPort** | Pointer to **int64** | The TCP port for TLS WEB client connections. The default value is &#x60;1443&#x60;. Available since 2.17. | [optional] 
**ServiceWebTransportWebUrlSuffix** | Pointer to **string** | Used to specify the Web URL suffix that will be used by Web clients when communicating with the broker. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.17. | [optional] 
**TlsBlockVersion11Enabled** | Pointer to **bool** | Enable or disable the blocking of TLS version 1.1 connections. When blocked, all existing incoming and outgoing TLS 1.1 connections with Clients, SEMP users, and LDAP servers remain connected while new connections are blocked. Note that support for TLS 1.1 will eventually be discontinued, at which time TLS 1.1 connections will be blocked regardless of this setting. The default value is &#x60;false&#x60;. | [optional] 
**TlsCipherSuiteManagementList** | Pointer to **string** | The colon-separated list of cipher suites used for TLS management connections (e.g. SEMP, LDAP). The value \&quot;default\&quot; implies all supported suites ordered from most secure to least secure. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 
**TlsCipherSuiteMsgBackboneList** | Pointer to **string** | The colon-separated list of cipher suites used for TLS data connections (e.g. client pub/sub). The value \&quot;default\&quot; implies all supported suites ordered from most secure to least secure. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 
**TlsCipherSuiteSecureShellList** | Pointer to **string** | The colon-separated list of cipher suites used for TLS secure shell connections (e.g. SSH, SFTP, SCP). The value \&quot;default\&quot; implies all supported suites ordered from most secure to least secure. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 
**TlsCrimeExploitProtectionEnabled** | Pointer to **bool** | Enable or disable protection against the CRIME exploit. When enabled, TLS+compressed messaging performance is degraded. This protection should only be disabled if sufficient ACL and authentication features are being employed such that a potential attacker does not have sufficient access to trigger the exploit. The default value is &#x60;true&#x60;. | [optional] 
**TlsServerCertContent** | Pointer to **string** | The PEM formatted content for the server certificate used for TLS connections. It must consist of a private key and between one and three certificates comprising the certificate trust chain. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**TlsServerCertPassword** | Pointer to **string** | The password for the server certificate used for TLS connections. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**TlsStandardDomainCertificateAuthoritiesEnabled** | Pointer to **bool** | Enable or disable the standard domain certificate authority list. The default value is &#x60;true&#x60;. Available since 2.19. | [optional] 
**TlsTicketLifetime** | Pointer to **int32** | The TLS ticket lifetime in seconds. When a client connects with TLS, a session with a session ticket is created using the TLS ticket lifetime which determines how long the client has to resume the session. The default value is &#x60;86400&#x60;. | [optional] 

## Methods

### NewBroker

`func NewBroker() *Broker`

NewBroker instantiates a new Broker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerWithDefaults

`func NewBrokerWithDefaults() *Broker`

NewBrokerWithDefaults instantiates a new Broker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthClientCertRevocationCheckMode

`func (o *Broker) GetAuthClientCertRevocationCheckMode() string`

GetAuthClientCertRevocationCheckMode returns the AuthClientCertRevocationCheckMode field if non-nil, zero value otherwise.

### GetAuthClientCertRevocationCheckModeOk

`func (o *Broker) GetAuthClientCertRevocationCheckModeOk() (*string, bool)`

GetAuthClientCertRevocationCheckModeOk returns a tuple with the AuthClientCertRevocationCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthClientCertRevocationCheckMode

`func (o *Broker) SetAuthClientCertRevocationCheckMode(v string)`

SetAuthClientCertRevocationCheckMode sets AuthClientCertRevocationCheckMode field to given value.

### HasAuthClientCertRevocationCheckMode

`func (o *Broker) HasAuthClientCertRevocationCheckMode() bool`

HasAuthClientCertRevocationCheckMode returns a boolean if a field has been set.

### GetGuaranteedMsgingEnabled

`func (o *Broker) GetGuaranteedMsgingEnabled() bool`

GetGuaranteedMsgingEnabled returns the GuaranteedMsgingEnabled field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEnabledOk

`func (o *Broker) GetGuaranteedMsgingEnabledOk() (*bool, bool)`

GetGuaranteedMsgingEnabledOk returns a tuple with the GuaranteedMsgingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEnabled

`func (o *Broker) SetGuaranteedMsgingEnabled(v bool)`

SetGuaranteedMsgingEnabled sets GuaranteedMsgingEnabled field to given value.

### HasGuaranteedMsgingEnabled

`func (o *Broker) HasGuaranteedMsgingEnabled() bool`

HasGuaranteedMsgingEnabled returns a boolean if a field has been set.

### GetGuaranteedMsgingEventCacheUsageThreshold

`func (o *Broker) GetGuaranteedMsgingEventCacheUsageThreshold() EventThreshold`

GetGuaranteedMsgingEventCacheUsageThreshold returns the GuaranteedMsgingEventCacheUsageThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventCacheUsageThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventCacheUsageThresholdOk() (*EventThreshold, bool)`

GetGuaranteedMsgingEventCacheUsageThresholdOk returns a tuple with the GuaranteedMsgingEventCacheUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventCacheUsageThreshold

`func (o *Broker) SetGuaranteedMsgingEventCacheUsageThreshold(v EventThreshold)`

SetGuaranteedMsgingEventCacheUsageThreshold sets GuaranteedMsgingEventCacheUsageThreshold field to given value.

### HasGuaranteedMsgingEventCacheUsageThreshold

`func (o *Broker) HasGuaranteedMsgingEventCacheUsageThreshold() bool`

HasGuaranteedMsgingEventCacheUsageThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventDeliveredUnackedThreshold

`func (o *Broker) GetGuaranteedMsgingEventDeliveredUnackedThreshold() EventThresholdByPercent`

GetGuaranteedMsgingEventDeliveredUnackedThreshold returns the GuaranteedMsgingEventDeliveredUnackedThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventDeliveredUnackedThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventDeliveredUnackedThresholdOk() (*EventThresholdByPercent, bool)`

GetGuaranteedMsgingEventDeliveredUnackedThresholdOk returns a tuple with the GuaranteedMsgingEventDeliveredUnackedThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventDeliveredUnackedThreshold

`func (o *Broker) SetGuaranteedMsgingEventDeliveredUnackedThreshold(v EventThresholdByPercent)`

SetGuaranteedMsgingEventDeliveredUnackedThreshold sets GuaranteedMsgingEventDeliveredUnackedThreshold field to given value.

### HasGuaranteedMsgingEventDeliveredUnackedThreshold

`func (o *Broker) HasGuaranteedMsgingEventDeliveredUnackedThreshold() bool`

HasGuaranteedMsgingEventDeliveredUnackedThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventDiskUsageThreshold

`func (o *Broker) GetGuaranteedMsgingEventDiskUsageThreshold() EventThresholdByPercent`

GetGuaranteedMsgingEventDiskUsageThreshold returns the GuaranteedMsgingEventDiskUsageThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventDiskUsageThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventDiskUsageThresholdOk() (*EventThresholdByPercent, bool)`

GetGuaranteedMsgingEventDiskUsageThresholdOk returns a tuple with the GuaranteedMsgingEventDiskUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventDiskUsageThreshold

`func (o *Broker) SetGuaranteedMsgingEventDiskUsageThreshold(v EventThresholdByPercent)`

SetGuaranteedMsgingEventDiskUsageThreshold sets GuaranteedMsgingEventDiskUsageThreshold field to given value.

### HasGuaranteedMsgingEventDiskUsageThreshold

`func (o *Broker) HasGuaranteedMsgingEventDiskUsageThreshold() bool`

HasGuaranteedMsgingEventDiskUsageThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventEgressFlowCountThreshold

`func (o *Broker) GetGuaranteedMsgingEventEgressFlowCountThreshold() EventThreshold`

GetGuaranteedMsgingEventEgressFlowCountThreshold returns the GuaranteedMsgingEventEgressFlowCountThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventEgressFlowCountThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventEgressFlowCountThresholdOk() (*EventThreshold, bool)`

GetGuaranteedMsgingEventEgressFlowCountThresholdOk returns a tuple with the GuaranteedMsgingEventEgressFlowCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventEgressFlowCountThreshold

`func (o *Broker) SetGuaranteedMsgingEventEgressFlowCountThreshold(v EventThreshold)`

SetGuaranteedMsgingEventEgressFlowCountThreshold sets GuaranteedMsgingEventEgressFlowCountThreshold field to given value.

### HasGuaranteedMsgingEventEgressFlowCountThreshold

`func (o *Broker) HasGuaranteedMsgingEventEgressFlowCountThreshold() bool`

HasGuaranteedMsgingEventEgressFlowCountThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventEndpointCountThreshold

`func (o *Broker) GetGuaranteedMsgingEventEndpointCountThreshold() EventThreshold`

GetGuaranteedMsgingEventEndpointCountThreshold returns the GuaranteedMsgingEventEndpointCountThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventEndpointCountThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventEndpointCountThresholdOk() (*EventThreshold, bool)`

GetGuaranteedMsgingEventEndpointCountThresholdOk returns a tuple with the GuaranteedMsgingEventEndpointCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventEndpointCountThreshold

`func (o *Broker) SetGuaranteedMsgingEventEndpointCountThreshold(v EventThreshold)`

SetGuaranteedMsgingEventEndpointCountThreshold sets GuaranteedMsgingEventEndpointCountThreshold field to given value.

### HasGuaranteedMsgingEventEndpointCountThreshold

`func (o *Broker) HasGuaranteedMsgingEventEndpointCountThreshold() bool`

HasGuaranteedMsgingEventEndpointCountThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventIngressFlowCountThreshold

`func (o *Broker) GetGuaranteedMsgingEventIngressFlowCountThreshold() EventThreshold`

GetGuaranteedMsgingEventIngressFlowCountThreshold returns the GuaranteedMsgingEventIngressFlowCountThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventIngressFlowCountThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventIngressFlowCountThresholdOk() (*EventThreshold, bool)`

GetGuaranteedMsgingEventIngressFlowCountThresholdOk returns a tuple with the GuaranteedMsgingEventIngressFlowCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventIngressFlowCountThreshold

`func (o *Broker) SetGuaranteedMsgingEventIngressFlowCountThreshold(v EventThreshold)`

SetGuaranteedMsgingEventIngressFlowCountThreshold sets GuaranteedMsgingEventIngressFlowCountThreshold field to given value.

### HasGuaranteedMsgingEventIngressFlowCountThreshold

`func (o *Broker) HasGuaranteedMsgingEventIngressFlowCountThreshold() bool`

HasGuaranteedMsgingEventIngressFlowCountThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventMsgCountThreshold

`func (o *Broker) GetGuaranteedMsgingEventMsgCountThreshold() EventThresholdByPercent`

GetGuaranteedMsgingEventMsgCountThreshold returns the GuaranteedMsgingEventMsgCountThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventMsgCountThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventMsgCountThresholdOk() (*EventThresholdByPercent, bool)`

GetGuaranteedMsgingEventMsgCountThresholdOk returns a tuple with the GuaranteedMsgingEventMsgCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventMsgCountThreshold

`func (o *Broker) SetGuaranteedMsgingEventMsgCountThreshold(v EventThresholdByPercent)`

SetGuaranteedMsgingEventMsgCountThreshold sets GuaranteedMsgingEventMsgCountThreshold field to given value.

### HasGuaranteedMsgingEventMsgCountThreshold

`func (o *Broker) HasGuaranteedMsgingEventMsgCountThreshold() bool`

HasGuaranteedMsgingEventMsgCountThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventMsgSpoolFileCountThreshold

`func (o *Broker) GetGuaranteedMsgingEventMsgSpoolFileCountThreshold() EventThresholdByPercent`

GetGuaranteedMsgingEventMsgSpoolFileCountThreshold returns the GuaranteedMsgingEventMsgSpoolFileCountThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventMsgSpoolFileCountThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventMsgSpoolFileCountThresholdOk() (*EventThresholdByPercent, bool)`

GetGuaranteedMsgingEventMsgSpoolFileCountThresholdOk returns a tuple with the GuaranteedMsgingEventMsgSpoolFileCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventMsgSpoolFileCountThreshold

`func (o *Broker) SetGuaranteedMsgingEventMsgSpoolFileCountThreshold(v EventThresholdByPercent)`

SetGuaranteedMsgingEventMsgSpoolFileCountThreshold sets GuaranteedMsgingEventMsgSpoolFileCountThreshold field to given value.

### HasGuaranteedMsgingEventMsgSpoolFileCountThreshold

`func (o *Broker) HasGuaranteedMsgingEventMsgSpoolFileCountThreshold() bool`

HasGuaranteedMsgingEventMsgSpoolFileCountThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventMsgSpoolUsageThreshold

`func (o *Broker) GetGuaranteedMsgingEventMsgSpoolUsageThreshold() EventThreshold`

GetGuaranteedMsgingEventMsgSpoolUsageThreshold returns the GuaranteedMsgingEventMsgSpoolUsageThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventMsgSpoolUsageThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventMsgSpoolUsageThresholdOk() (*EventThreshold, bool)`

GetGuaranteedMsgingEventMsgSpoolUsageThresholdOk returns a tuple with the GuaranteedMsgingEventMsgSpoolUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventMsgSpoolUsageThreshold

`func (o *Broker) SetGuaranteedMsgingEventMsgSpoolUsageThreshold(v EventThreshold)`

SetGuaranteedMsgingEventMsgSpoolUsageThreshold sets GuaranteedMsgingEventMsgSpoolUsageThreshold field to given value.

### HasGuaranteedMsgingEventMsgSpoolUsageThreshold

`func (o *Broker) HasGuaranteedMsgingEventMsgSpoolUsageThreshold() bool`

HasGuaranteedMsgingEventMsgSpoolUsageThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventTransactedSessionCountThreshold

`func (o *Broker) GetGuaranteedMsgingEventTransactedSessionCountThreshold() EventThreshold`

GetGuaranteedMsgingEventTransactedSessionCountThreshold returns the GuaranteedMsgingEventTransactedSessionCountThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventTransactedSessionCountThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventTransactedSessionCountThresholdOk() (*EventThreshold, bool)`

GetGuaranteedMsgingEventTransactedSessionCountThresholdOk returns a tuple with the GuaranteedMsgingEventTransactedSessionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventTransactedSessionCountThreshold

`func (o *Broker) SetGuaranteedMsgingEventTransactedSessionCountThreshold(v EventThreshold)`

SetGuaranteedMsgingEventTransactedSessionCountThreshold sets GuaranteedMsgingEventTransactedSessionCountThreshold field to given value.

### HasGuaranteedMsgingEventTransactedSessionCountThreshold

`func (o *Broker) HasGuaranteedMsgingEventTransactedSessionCountThreshold() bool`

HasGuaranteedMsgingEventTransactedSessionCountThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventTransactedSessionResourceCountThreshold

`func (o *Broker) GetGuaranteedMsgingEventTransactedSessionResourceCountThreshold() EventThresholdByPercent`

GetGuaranteedMsgingEventTransactedSessionResourceCountThreshold returns the GuaranteedMsgingEventTransactedSessionResourceCountThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventTransactedSessionResourceCountThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventTransactedSessionResourceCountThresholdOk() (*EventThresholdByPercent, bool)`

GetGuaranteedMsgingEventTransactedSessionResourceCountThresholdOk returns a tuple with the GuaranteedMsgingEventTransactedSessionResourceCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventTransactedSessionResourceCountThreshold

`func (o *Broker) SetGuaranteedMsgingEventTransactedSessionResourceCountThreshold(v EventThresholdByPercent)`

SetGuaranteedMsgingEventTransactedSessionResourceCountThreshold sets GuaranteedMsgingEventTransactedSessionResourceCountThreshold field to given value.

### HasGuaranteedMsgingEventTransactedSessionResourceCountThreshold

`func (o *Broker) HasGuaranteedMsgingEventTransactedSessionResourceCountThreshold() bool`

HasGuaranteedMsgingEventTransactedSessionResourceCountThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingEventTransactionCountThreshold

`func (o *Broker) GetGuaranteedMsgingEventTransactionCountThreshold() EventThreshold`

GetGuaranteedMsgingEventTransactionCountThreshold returns the GuaranteedMsgingEventTransactionCountThreshold field if non-nil, zero value otherwise.

### GetGuaranteedMsgingEventTransactionCountThresholdOk

`func (o *Broker) GetGuaranteedMsgingEventTransactionCountThresholdOk() (*EventThreshold, bool)`

GetGuaranteedMsgingEventTransactionCountThresholdOk returns a tuple with the GuaranteedMsgingEventTransactionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingEventTransactionCountThreshold

`func (o *Broker) SetGuaranteedMsgingEventTransactionCountThreshold(v EventThreshold)`

SetGuaranteedMsgingEventTransactionCountThreshold sets GuaranteedMsgingEventTransactionCountThreshold field to given value.

### HasGuaranteedMsgingEventTransactionCountThreshold

`func (o *Broker) HasGuaranteedMsgingEventTransactionCountThreshold() bool`

HasGuaranteedMsgingEventTransactionCountThreshold returns a boolean if a field has been set.

### GetGuaranteedMsgingMaxCacheUsage

`func (o *Broker) GetGuaranteedMsgingMaxCacheUsage() int32`

GetGuaranteedMsgingMaxCacheUsage returns the GuaranteedMsgingMaxCacheUsage field if non-nil, zero value otherwise.

### GetGuaranteedMsgingMaxCacheUsageOk

`func (o *Broker) GetGuaranteedMsgingMaxCacheUsageOk() (*int32, bool)`

GetGuaranteedMsgingMaxCacheUsageOk returns a tuple with the GuaranteedMsgingMaxCacheUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingMaxCacheUsage

`func (o *Broker) SetGuaranteedMsgingMaxCacheUsage(v int32)`

SetGuaranteedMsgingMaxCacheUsage sets GuaranteedMsgingMaxCacheUsage field to given value.

### HasGuaranteedMsgingMaxCacheUsage

`func (o *Broker) HasGuaranteedMsgingMaxCacheUsage() bool`

HasGuaranteedMsgingMaxCacheUsage returns a boolean if a field has been set.

### GetGuaranteedMsgingMaxMsgSpoolUsage

`func (o *Broker) GetGuaranteedMsgingMaxMsgSpoolUsage() int64`

GetGuaranteedMsgingMaxMsgSpoolUsage returns the GuaranteedMsgingMaxMsgSpoolUsage field if non-nil, zero value otherwise.

### GetGuaranteedMsgingMaxMsgSpoolUsageOk

`func (o *Broker) GetGuaranteedMsgingMaxMsgSpoolUsageOk() (*int64, bool)`

GetGuaranteedMsgingMaxMsgSpoolUsageOk returns a tuple with the GuaranteedMsgingMaxMsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingMaxMsgSpoolUsage

`func (o *Broker) SetGuaranteedMsgingMaxMsgSpoolUsage(v int64)`

SetGuaranteedMsgingMaxMsgSpoolUsage sets GuaranteedMsgingMaxMsgSpoolUsage field to given value.

### HasGuaranteedMsgingMaxMsgSpoolUsage

`func (o *Broker) HasGuaranteedMsgingMaxMsgSpoolUsage() bool`

HasGuaranteedMsgingMaxMsgSpoolUsage returns a boolean if a field has been set.

### GetGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout

`func (o *Broker) GetGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout() int64`

GetGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout returns the GuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout field if non-nil, zero value otherwise.

### GetGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeoutOk

`func (o *Broker) GetGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeoutOk() (*int64, bool)`

GetGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeoutOk returns a tuple with the GuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout

`func (o *Broker) SetGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout(v int64)`

SetGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout sets GuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout field to given value.

### HasGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout

`func (o *Broker) HasGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout() bool`

HasGuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout returns a boolean if a field has been set.

### GetGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout

`func (o *Broker) GetGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout() int64`

GetGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout returns the GuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout field if non-nil, zero value otherwise.

### GetGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeoutOk

`func (o *Broker) GetGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeoutOk() (*int64, bool)`

GetGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeoutOk returns a tuple with the GuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout

`func (o *Broker) SetGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout(v int64)`

SetGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout sets GuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout field to given value.

### HasGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout

`func (o *Broker) HasGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout() bool`

HasGuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout returns a boolean if a field has been set.

### GetGuaranteedMsgingTransactionReplicationCompatibilityMode

`func (o *Broker) GetGuaranteedMsgingTransactionReplicationCompatibilityMode() string`

GetGuaranteedMsgingTransactionReplicationCompatibilityMode returns the GuaranteedMsgingTransactionReplicationCompatibilityMode field if non-nil, zero value otherwise.

### GetGuaranteedMsgingTransactionReplicationCompatibilityModeOk

`func (o *Broker) GetGuaranteedMsgingTransactionReplicationCompatibilityModeOk() (*string, bool)`

GetGuaranteedMsgingTransactionReplicationCompatibilityModeOk returns a tuple with the GuaranteedMsgingTransactionReplicationCompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgingTransactionReplicationCompatibilityMode

`func (o *Broker) SetGuaranteedMsgingTransactionReplicationCompatibilityMode(v string)`

SetGuaranteedMsgingTransactionReplicationCompatibilityMode sets GuaranteedMsgingTransactionReplicationCompatibilityMode field to given value.

### HasGuaranteedMsgingTransactionReplicationCompatibilityMode

`func (o *Broker) HasGuaranteedMsgingTransactionReplicationCompatibilityMode() bool`

HasGuaranteedMsgingTransactionReplicationCompatibilityMode returns a boolean if a field has been set.

### GetServiceAmqpEnabled

`func (o *Broker) GetServiceAmqpEnabled() bool`

GetServiceAmqpEnabled returns the ServiceAmqpEnabled field if non-nil, zero value otherwise.

### GetServiceAmqpEnabledOk

`func (o *Broker) GetServiceAmqpEnabledOk() (*bool, bool)`

GetServiceAmqpEnabledOk returns a tuple with the ServiceAmqpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpEnabled

`func (o *Broker) SetServiceAmqpEnabled(v bool)`

SetServiceAmqpEnabled sets ServiceAmqpEnabled field to given value.

### HasServiceAmqpEnabled

`func (o *Broker) HasServiceAmqpEnabled() bool`

HasServiceAmqpEnabled returns a boolean if a field has been set.

### GetServiceAmqpTlsListenPort

`func (o *Broker) GetServiceAmqpTlsListenPort() int64`

GetServiceAmqpTlsListenPort returns the ServiceAmqpTlsListenPort field if non-nil, zero value otherwise.

### GetServiceAmqpTlsListenPortOk

`func (o *Broker) GetServiceAmqpTlsListenPortOk() (*int64, bool)`

GetServiceAmqpTlsListenPortOk returns a tuple with the ServiceAmqpTlsListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpTlsListenPort

`func (o *Broker) SetServiceAmqpTlsListenPort(v int64)`

SetServiceAmqpTlsListenPort sets ServiceAmqpTlsListenPort field to given value.

### HasServiceAmqpTlsListenPort

`func (o *Broker) HasServiceAmqpTlsListenPort() bool`

HasServiceAmqpTlsListenPort returns a boolean if a field has been set.

### GetServiceEventConnectionCountThreshold

`func (o *Broker) GetServiceEventConnectionCountThreshold() EventThreshold`

GetServiceEventConnectionCountThreshold returns the ServiceEventConnectionCountThreshold field if non-nil, zero value otherwise.

### GetServiceEventConnectionCountThresholdOk

`func (o *Broker) GetServiceEventConnectionCountThresholdOk() (*EventThreshold, bool)`

GetServiceEventConnectionCountThresholdOk returns a tuple with the ServiceEventConnectionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEventConnectionCountThreshold

`func (o *Broker) SetServiceEventConnectionCountThreshold(v EventThreshold)`

SetServiceEventConnectionCountThreshold sets ServiceEventConnectionCountThreshold field to given value.

### HasServiceEventConnectionCountThreshold

`func (o *Broker) HasServiceEventConnectionCountThreshold() bool`

HasServiceEventConnectionCountThreshold returns a boolean if a field has been set.

### GetServiceHealthCheckEnabled

`func (o *Broker) GetServiceHealthCheckEnabled() bool`

GetServiceHealthCheckEnabled returns the ServiceHealthCheckEnabled field if non-nil, zero value otherwise.

### GetServiceHealthCheckEnabledOk

`func (o *Broker) GetServiceHealthCheckEnabledOk() (*bool, bool)`

GetServiceHealthCheckEnabledOk returns a tuple with the ServiceHealthCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHealthCheckEnabled

`func (o *Broker) SetServiceHealthCheckEnabled(v bool)`

SetServiceHealthCheckEnabled sets ServiceHealthCheckEnabled field to given value.

### HasServiceHealthCheckEnabled

`func (o *Broker) HasServiceHealthCheckEnabled() bool`

HasServiceHealthCheckEnabled returns a boolean if a field has been set.

### GetServiceHealthCheckListenPort

`func (o *Broker) GetServiceHealthCheckListenPort() int64`

GetServiceHealthCheckListenPort returns the ServiceHealthCheckListenPort field if non-nil, zero value otherwise.

### GetServiceHealthCheckListenPortOk

`func (o *Broker) GetServiceHealthCheckListenPortOk() (*int64, bool)`

GetServiceHealthCheckListenPortOk returns a tuple with the ServiceHealthCheckListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHealthCheckListenPort

`func (o *Broker) SetServiceHealthCheckListenPort(v int64)`

SetServiceHealthCheckListenPort sets ServiceHealthCheckListenPort field to given value.

### HasServiceHealthCheckListenPort

`func (o *Broker) HasServiceHealthCheckListenPort() bool`

HasServiceHealthCheckListenPort returns a boolean if a field has been set.

### GetServiceMateLinkEnabled

`func (o *Broker) GetServiceMateLinkEnabled() bool`

GetServiceMateLinkEnabled returns the ServiceMateLinkEnabled field if non-nil, zero value otherwise.

### GetServiceMateLinkEnabledOk

`func (o *Broker) GetServiceMateLinkEnabledOk() (*bool, bool)`

GetServiceMateLinkEnabledOk returns a tuple with the ServiceMateLinkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMateLinkEnabled

`func (o *Broker) SetServiceMateLinkEnabled(v bool)`

SetServiceMateLinkEnabled sets ServiceMateLinkEnabled field to given value.

### HasServiceMateLinkEnabled

`func (o *Broker) HasServiceMateLinkEnabled() bool`

HasServiceMateLinkEnabled returns a boolean if a field has been set.

### GetServiceMateLinkListenPort

`func (o *Broker) GetServiceMateLinkListenPort() int64`

GetServiceMateLinkListenPort returns the ServiceMateLinkListenPort field if non-nil, zero value otherwise.

### GetServiceMateLinkListenPortOk

`func (o *Broker) GetServiceMateLinkListenPortOk() (*int64, bool)`

GetServiceMateLinkListenPortOk returns a tuple with the ServiceMateLinkListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMateLinkListenPort

`func (o *Broker) SetServiceMateLinkListenPort(v int64)`

SetServiceMateLinkListenPort sets ServiceMateLinkListenPort field to given value.

### HasServiceMateLinkListenPort

`func (o *Broker) HasServiceMateLinkListenPort() bool`

HasServiceMateLinkListenPort returns a boolean if a field has been set.

### GetServiceMqttEnabled

`func (o *Broker) GetServiceMqttEnabled() bool`

GetServiceMqttEnabled returns the ServiceMqttEnabled field if non-nil, zero value otherwise.

### GetServiceMqttEnabledOk

`func (o *Broker) GetServiceMqttEnabledOk() (*bool, bool)`

GetServiceMqttEnabledOk returns a tuple with the ServiceMqttEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttEnabled

`func (o *Broker) SetServiceMqttEnabled(v bool)`

SetServiceMqttEnabled sets ServiceMqttEnabled field to given value.

### HasServiceMqttEnabled

`func (o *Broker) HasServiceMqttEnabled() bool`

HasServiceMqttEnabled returns a boolean if a field has been set.

### GetServiceMsgBackboneEnabled

`func (o *Broker) GetServiceMsgBackboneEnabled() bool`

GetServiceMsgBackboneEnabled returns the ServiceMsgBackboneEnabled field if non-nil, zero value otherwise.

### GetServiceMsgBackboneEnabledOk

`func (o *Broker) GetServiceMsgBackboneEnabledOk() (*bool, bool)`

GetServiceMsgBackboneEnabledOk returns a tuple with the ServiceMsgBackboneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMsgBackboneEnabled

`func (o *Broker) SetServiceMsgBackboneEnabled(v bool)`

SetServiceMsgBackboneEnabled sets ServiceMsgBackboneEnabled field to given value.

### HasServiceMsgBackboneEnabled

`func (o *Broker) HasServiceMsgBackboneEnabled() bool`

HasServiceMsgBackboneEnabled returns a boolean if a field has been set.

### GetServiceRedundancyEnabled

`func (o *Broker) GetServiceRedundancyEnabled() bool`

GetServiceRedundancyEnabled returns the ServiceRedundancyEnabled field if non-nil, zero value otherwise.

### GetServiceRedundancyEnabledOk

`func (o *Broker) GetServiceRedundancyEnabledOk() (*bool, bool)`

GetServiceRedundancyEnabledOk returns a tuple with the ServiceRedundancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRedundancyEnabled

`func (o *Broker) SetServiceRedundancyEnabled(v bool)`

SetServiceRedundancyEnabled sets ServiceRedundancyEnabled field to given value.

### HasServiceRedundancyEnabled

`func (o *Broker) HasServiceRedundancyEnabled() bool`

HasServiceRedundancyEnabled returns a boolean if a field has been set.

### GetServiceRedundancyFirstListenPort

`func (o *Broker) GetServiceRedundancyFirstListenPort() int64`

GetServiceRedundancyFirstListenPort returns the ServiceRedundancyFirstListenPort field if non-nil, zero value otherwise.

### GetServiceRedundancyFirstListenPortOk

`func (o *Broker) GetServiceRedundancyFirstListenPortOk() (*int64, bool)`

GetServiceRedundancyFirstListenPortOk returns a tuple with the ServiceRedundancyFirstListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRedundancyFirstListenPort

`func (o *Broker) SetServiceRedundancyFirstListenPort(v int64)`

SetServiceRedundancyFirstListenPort sets ServiceRedundancyFirstListenPort field to given value.

### HasServiceRedundancyFirstListenPort

`func (o *Broker) HasServiceRedundancyFirstListenPort() bool`

HasServiceRedundancyFirstListenPort returns a boolean if a field has been set.

### GetServiceRestEventOutgoingConnectionCountThreshold

`func (o *Broker) GetServiceRestEventOutgoingConnectionCountThreshold() EventThreshold`

GetServiceRestEventOutgoingConnectionCountThreshold returns the ServiceRestEventOutgoingConnectionCountThreshold field if non-nil, zero value otherwise.

### GetServiceRestEventOutgoingConnectionCountThresholdOk

`func (o *Broker) GetServiceRestEventOutgoingConnectionCountThresholdOk() (*EventThreshold, bool)`

GetServiceRestEventOutgoingConnectionCountThresholdOk returns a tuple with the ServiceRestEventOutgoingConnectionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestEventOutgoingConnectionCountThreshold

`func (o *Broker) SetServiceRestEventOutgoingConnectionCountThreshold(v EventThreshold)`

SetServiceRestEventOutgoingConnectionCountThreshold sets ServiceRestEventOutgoingConnectionCountThreshold field to given value.

### HasServiceRestEventOutgoingConnectionCountThreshold

`func (o *Broker) HasServiceRestEventOutgoingConnectionCountThreshold() bool`

HasServiceRestEventOutgoingConnectionCountThreshold returns a boolean if a field has been set.

### GetServiceRestIncomingEnabled

`func (o *Broker) GetServiceRestIncomingEnabled() bool`

GetServiceRestIncomingEnabled returns the ServiceRestIncomingEnabled field if non-nil, zero value otherwise.

### GetServiceRestIncomingEnabledOk

`func (o *Broker) GetServiceRestIncomingEnabledOk() (*bool, bool)`

GetServiceRestIncomingEnabledOk returns a tuple with the ServiceRestIncomingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingEnabled

`func (o *Broker) SetServiceRestIncomingEnabled(v bool)`

SetServiceRestIncomingEnabled sets ServiceRestIncomingEnabled field to given value.

### HasServiceRestIncomingEnabled

`func (o *Broker) HasServiceRestIncomingEnabled() bool`

HasServiceRestIncomingEnabled returns a boolean if a field has been set.

### GetServiceRestOutgoingEnabled

`func (o *Broker) GetServiceRestOutgoingEnabled() bool`

GetServiceRestOutgoingEnabled returns the ServiceRestOutgoingEnabled field if non-nil, zero value otherwise.

### GetServiceRestOutgoingEnabledOk

`func (o *Broker) GetServiceRestOutgoingEnabledOk() (*bool, bool)`

GetServiceRestOutgoingEnabledOk returns a tuple with the ServiceRestOutgoingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestOutgoingEnabled

`func (o *Broker) SetServiceRestOutgoingEnabled(v bool)`

SetServiceRestOutgoingEnabled sets ServiceRestOutgoingEnabled field to given value.

### HasServiceRestOutgoingEnabled

`func (o *Broker) HasServiceRestOutgoingEnabled() bool`

HasServiceRestOutgoingEnabled returns a boolean if a field has been set.

### GetServiceSempLegacyTimeoutEnabled

`func (o *Broker) GetServiceSempLegacyTimeoutEnabled() bool`

GetServiceSempLegacyTimeoutEnabled returns the ServiceSempLegacyTimeoutEnabled field if non-nil, zero value otherwise.

### GetServiceSempLegacyTimeoutEnabledOk

`func (o *Broker) GetServiceSempLegacyTimeoutEnabledOk() (*bool, bool)`

GetServiceSempLegacyTimeoutEnabledOk returns a tuple with the ServiceSempLegacyTimeoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSempLegacyTimeoutEnabled

`func (o *Broker) SetServiceSempLegacyTimeoutEnabled(v bool)`

SetServiceSempLegacyTimeoutEnabled sets ServiceSempLegacyTimeoutEnabled field to given value.

### HasServiceSempLegacyTimeoutEnabled

`func (o *Broker) HasServiceSempLegacyTimeoutEnabled() bool`

HasServiceSempLegacyTimeoutEnabled returns a boolean if a field has been set.

### GetServiceSempPlainTextEnabled

`func (o *Broker) GetServiceSempPlainTextEnabled() bool`

GetServiceSempPlainTextEnabled returns the ServiceSempPlainTextEnabled field if non-nil, zero value otherwise.

### GetServiceSempPlainTextEnabledOk

`func (o *Broker) GetServiceSempPlainTextEnabledOk() (*bool, bool)`

GetServiceSempPlainTextEnabledOk returns a tuple with the ServiceSempPlainTextEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSempPlainTextEnabled

`func (o *Broker) SetServiceSempPlainTextEnabled(v bool)`

SetServiceSempPlainTextEnabled sets ServiceSempPlainTextEnabled field to given value.

### HasServiceSempPlainTextEnabled

`func (o *Broker) HasServiceSempPlainTextEnabled() bool`

HasServiceSempPlainTextEnabled returns a boolean if a field has been set.

### GetServiceSempPlainTextListenPort

`func (o *Broker) GetServiceSempPlainTextListenPort() int64`

GetServiceSempPlainTextListenPort returns the ServiceSempPlainTextListenPort field if non-nil, zero value otherwise.

### GetServiceSempPlainTextListenPortOk

`func (o *Broker) GetServiceSempPlainTextListenPortOk() (*int64, bool)`

GetServiceSempPlainTextListenPortOk returns a tuple with the ServiceSempPlainTextListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSempPlainTextListenPort

`func (o *Broker) SetServiceSempPlainTextListenPort(v int64)`

SetServiceSempPlainTextListenPort sets ServiceSempPlainTextListenPort field to given value.

### HasServiceSempPlainTextListenPort

`func (o *Broker) HasServiceSempPlainTextListenPort() bool`

HasServiceSempPlainTextListenPort returns a boolean if a field has been set.

### GetServiceSempSessionIdleTimeout

`func (o *Broker) GetServiceSempSessionIdleTimeout() int32`

GetServiceSempSessionIdleTimeout returns the ServiceSempSessionIdleTimeout field if non-nil, zero value otherwise.

### GetServiceSempSessionIdleTimeoutOk

`func (o *Broker) GetServiceSempSessionIdleTimeoutOk() (*int32, bool)`

GetServiceSempSessionIdleTimeoutOk returns a tuple with the ServiceSempSessionIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSempSessionIdleTimeout

`func (o *Broker) SetServiceSempSessionIdleTimeout(v int32)`

SetServiceSempSessionIdleTimeout sets ServiceSempSessionIdleTimeout field to given value.

### HasServiceSempSessionIdleTimeout

`func (o *Broker) HasServiceSempSessionIdleTimeout() bool`

HasServiceSempSessionIdleTimeout returns a boolean if a field has been set.

### GetServiceSempSessionMaxLifetime

`func (o *Broker) GetServiceSempSessionMaxLifetime() int32`

GetServiceSempSessionMaxLifetime returns the ServiceSempSessionMaxLifetime field if non-nil, zero value otherwise.

### GetServiceSempSessionMaxLifetimeOk

`func (o *Broker) GetServiceSempSessionMaxLifetimeOk() (*int32, bool)`

GetServiceSempSessionMaxLifetimeOk returns a tuple with the ServiceSempSessionMaxLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSempSessionMaxLifetime

`func (o *Broker) SetServiceSempSessionMaxLifetime(v int32)`

SetServiceSempSessionMaxLifetime sets ServiceSempSessionMaxLifetime field to given value.

### HasServiceSempSessionMaxLifetime

`func (o *Broker) HasServiceSempSessionMaxLifetime() bool`

HasServiceSempSessionMaxLifetime returns a boolean if a field has been set.

### GetServiceSempTlsEnabled

`func (o *Broker) GetServiceSempTlsEnabled() bool`

GetServiceSempTlsEnabled returns the ServiceSempTlsEnabled field if non-nil, zero value otherwise.

### GetServiceSempTlsEnabledOk

`func (o *Broker) GetServiceSempTlsEnabledOk() (*bool, bool)`

GetServiceSempTlsEnabledOk returns a tuple with the ServiceSempTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSempTlsEnabled

`func (o *Broker) SetServiceSempTlsEnabled(v bool)`

SetServiceSempTlsEnabled sets ServiceSempTlsEnabled field to given value.

### HasServiceSempTlsEnabled

`func (o *Broker) HasServiceSempTlsEnabled() bool`

HasServiceSempTlsEnabled returns a boolean if a field has been set.

### GetServiceSempTlsListenPort

`func (o *Broker) GetServiceSempTlsListenPort() int64`

GetServiceSempTlsListenPort returns the ServiceSempTlsListenPort field if non-nil, zero value otherwise.

### GetServiceSempTlsListenPortOk

`func (o *Broker) GetServiceSempTlsListenPortOk() (*int64, bool)`

GetServiceSempTlsListenPortOk returns a tuple with the ServiceSempTlsListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSempTlsListenPort

`func (o *Broker) SetServiceSempTlsListenPort(v int64)`

SetServiceSempTlsListenPort sets ServiceSempTlsListenPort field to given value.

### HasServiceSempTlsListenPort

`func (o *Broker) HasServiceSempTlsListenPort() bool`

HasServiceSempTlsListenPort returns a boolean if a field has been set.

### GetServiceSmfCompressionListenPort

`func (o *Broker) GetServiceSmfCompressionListenPort() int64`

GetServiceSmfCompressionListenPort returns the ServiceSmfCompressionListenPort field if non-nil, zero value otherwise.

### GetServiceSmfCompressionListenPortOk

`func (o *Broker) GetServiceSmfCompressionListenPortOk() (*int64, bool)`

GetServiceSmfCompressionListenPortOk returns a tuple with the ServiceSmfCompressionListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfCompressionListenPort

`func (o *Broker) SetServiceSmfCompressionListenPort(v int64)`

SetServiceSmfCompressionListenPort sets ServiceSmfCompressionListenPort field to given value.

### HasServiceSmfCompressionListenPort

`func (o *Broker) HasServiceSmfCompressionListenPort() bool`

HasServiceSmfCompressionListenPort returns a boolean if a field has been set.

### GetServiceSmfEnabled

`func (o *Broker) GetServiceSmfEnabled() bool`

GetServiceSmfEnabled returns the ServiceSmfEnabled field if non-nil, zero value otherwise.

### GetServiceSmfEnabledOk

`func (o *Broker) GetServiceSmfEnabledOk() (*bool, bool)`

GetServiceSmfEnabledOk returns a tuple with the ServiceSmfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfEnabled

`func (o *Broker) SetServiceSmfEnabled(v bool)`

SetServiceSmfEnabled sets ServiceSmfEnabled field to given value.

### HasServiceSmfEnabled

`func (o *Broker) HasServiceSmfEnabled() bool`

HasServiceSmfEnabled returns a boolean if a field has been set.

### GetServiceSmfEventConnectionCountThreshold

`func (o *Broker) GetServiceSmfEventConnectionCountThreshold() EventThreshold`

GetServiceSmfEventConnectionCountThreshold returns the ServiceSmfEventConnectionCountThreshold field if non-nil, zero value otherwise.

### GetServiceSmfEventConnectionCountThresholdOk

`func (o *Broker) GetServiceSmfEventConnectionCountThresholdOk() (*EventThreshold, bool)`

GetServiceSmfEventConnectionCountThresholdOk returns a tuple with the ServiceSmfEventConnectionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfEventConnectionCountThreshold

`func (o *Broker) SetServiceSmfEventConnectionCountThreshold(v EventThreshold)`

SetServiceSmfEventConnectionCountThreshold sets ServiceSmfEventConnectionCountThreshold field to given value.

### HasServiceSmfEventConnectionCountThreshold

`func (o *Broker) HasServiceSmfEventConnectionCountThreshold() bool`

HasServiceSmfEventConnectionCountThreshold returns a boolean if a field has been set.

### GetServiceSmfPlainTextListenPort

`func (o *Broker) GetServiceSmfPlainTextListenPort() int64`

GetServiceSmfPlainTextListenPort returns the ServiceSmfPlainTextListenPort field if non-nil, zero value otherwise.

### GetServiceSmfPlainTextListenPortOk

`func (o *Broker) GetServiceSmfPlainTextListenPortOk() (*int64, bool)`

GetServiceSmfPlainTextListenPortOk returns a tuple with the ServiceSmfPlainTextListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfPlainTextListenPort

`func (o *Broker) SetServiceSmfPlainTextListenPort(v int64)`

SetServiceSmfPlainTextListenPort sets ServiceSmfPlainTextListenPort field to given value.

### HasServiceSmfPlainTextListenPort

`func (o *Broker) HasServiceSmfPlainTextListenPort() bool`

HasServiceSmfPlainTextListenPort returns a boolean if a field has been set.

### GetServiceSmfRoutingControlListenPort

`func (o *Broker) GetServiceSmfRoutingControlListenPort() int64`

GetServiceSmfRoutingControlListenPort returns the ServiceSmfRoutingControlListenPort field if non-nil, zero value otherwise.

### GetServiceSmfRoutingControlListenPortOk

`func (o *Broker) GetServiceSmfRoutingControlListenPortOk() (*int64, bool)`

GetServiceSmfRoutingControlListenPortOk returns a tuple with the ServiceSmfRoutingControlListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfRoutingControlListenPort

`func (o *Broker) SetServiceSmfRoutingControlListenPort(v int64)`

SetServiceSmfRoutingControlListenPort sets ServiceSmfRoutingControlListenPort field to given value.

### HasServiceSmfRoutingControlListenPort

`func (o *Broker) HasServiceSmfRoutingControlListenPort() bool`

HasServiceSmfRoutingControlListenPort returns a boolean if a field has been set.

### GetServiceSmfTlsListenPort

`func (o *Broker) GetServiceSmfTlsListenPort() int64`

GetServiceSmfTlsListenPort returns the ServiceSmfTlsListenPort field if non-nil, zero value otherwise.

### GetServiceSmfTlsListenPortOk

`func (o *Broker) GetServiceSmfTlsListenPortOk() (*int64, bool)`

GetServiceSmfTlsListenPortOk returns a tuple with the ServiceSmfTlsListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfTlsListenPort

`func (o *Broker) SetServiceSmfTlsListenPort(v int64)`

SetServiceSmfTlsListenPort sets ServiceSmfTlsListenPort field to given value.

### HasServiceSmfTlsListenPort

`func (o *Broker) HasServiceSmfTlsListenPort() bool`

HasServiceSmfTlsListenPort returns a boolean if a field has been set.

### GetServiceTlsEventConnectionCountThreshold

`func (o *Broker) GetServiceTlsEventConnectionCountThreshold() EventThreshold`

GetServiceTlsEventConnectionCountThreshold returns the ServiceTlsEventConnectionCountThreshold field if non-nil, zero value otherwise.

### GetServiceTlsEventConnectionCountThresholdOk

`func (o *Broker) GetServiceTlsEventConnectionCountThresholdOk() (*EventThreshold, bool)`

GetServiceTlsEventConnectionCountThresholdOk returns a tuple with the ServiceTlsEventConnectionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTlsEventConnectionCountThreshold

`func (o *Broker) SetServiceTlsEventConnectionCountThreshold(v EventThreshold)`

SetServiceTlsEventConnectionCountThreshold sets ServiceTlsEventConnectionCountThreshold field to given value.

### HasServiceTlsEventConnectionCountThreshold

`func (o *Broker) HasServiceTlsEventConnectionCountThreshold() bool`

HasServiceTlsEventConnectionCountThreshold returns a boolean if a field has been set.

### GetServiceWebTransportEnabled

`func (o *Broker) GetServiceWebTransportEnabled() bool`

GetServiceWebTransportEnabled returns the ServiceWebTransportEnabled field if non-nil, zero value otherwise.

### GetServiceWebTransportEnabledOk

`func (o *Broker) GetServiceWebTransportEnabledOk() (*bool, bool)`

GetServiceWebTransportEnabledOk returns a tuple with the ServiceWebTransportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebTransportEnabled

`func (o *Broker) SetServiceWebTransportEnabled(v bool)`

SetServiceWebTransportEnabled sets ServiceWebTransportEnabled field to given value.

### HasServiceWebTransportEnabled

`func (o *Broker) HasServiceWebTransportEnabled() bool`

HasServiceWebTransportEnabled returns a boolean if a field has been set.

### GetServiceWebTransportPlainTextListenPort

`func (o *Broker) GetServiceWebTransportPlainTextListenPort() int64`

GetServiceWebTransportPlainTextListenPort returns the ServiceWebTransportPlainTextListenPort field if non-nil, zero value otherwise.

### GetServiceWebTransportPlainTextListenPortOk

`func (o *Broker) GetServiceWebTransportPlainTextListenPortOk() (*int64, bool)`

GetServiceWebTransportPlainTextListenPortOk returns a tuple with the ServiceWebTransportPlainTextListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebTransportPlainTextListenPort

`func (o *Broker) SetServiceWebTransportPlainTextListenPort(v int64)`

SetServiceWebTransportPlainTextListenPort sets ServiceWebTransportPlainTextListenPort field to given value.

### HasServiceWebTransportPlainTextListenPort

`func (o *Broker) HasServiceWebTransportPlainTextListenPort() bool`

HasServiceWebTransportPlainTextListenPort returns a boolean if a field has been set.

### GetServiceWebTransportTlsListenPort

`func (o *Broker) GetServiceWebTransportTlsListenPort() int64`

GetServiceWebTransportTlsListenPort returns the ServiceWebTransportTlsListenPort field if non-nil, zero value otherwise.

### GetServiceWebTransportTlsListenPortOk

`func (o *Broker) GetServiceWebTransportTlsListenPortOk() (*int64, bool)`

GetServiceWebTransportTlsListenPortOk returns a tuple with the ServiceWebTransportTlsListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebTransportTlsListenPort

`func (o *Broker) SetServiceWebTransportTlsListenPort(v int64)`

SetServiceWebTransportTlsListenPort sets ServiceWebTransportTlsListenPort field to given value.

### HasServiceWebTransportTlsListenPort

`func (o *Broker) HasServiceWebTransportTlsListenPort() bool`

HasServiceWebTransportTlsListenPort returns a boolean if a field has been set.

### GetServiceWebTransportWebUrlSuffix

`func (o *Broker) GetServiceWebTransportWebUrlSuffix() string`

GetServiceWebTransportWebUrlSuffix returns the ServiceWebTransportWebUrlSuffix field if non-nil, zero value otherwise.

### GetServiceWebTransportWebUrlSuffixOk

`func (o *Broker) GetServiceWebTransportWebUrlSuffixOk() (*string, bool)`

GetServiceWebTransportWebUrlSuffixOk returns a tuple with the ServiceWebTransportWebUrlSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebTransportWebUrlSuffix

`func (o *Broker) SetServiceWebTransportWebUrlSuffix(v string)`

SetServiceWebTransportWebUrlSuffix sets ServiceWebTransportWebUrlSuffix field to given value.

### HasServiceWebTransportWebUrlSuffix

`func (o *Broker) HasServiceWebTransportWebUrlSuffix() bool`

HasServiceWebTransportWebUrlSuffix returns a boolean if a field has been set.

### GetTlsBlockVersion11Enabled

`func (o *Broker) GetTlsBlockVersion11Enabled() bool`

GetTlsBlockVersion11Enabled returns the TlsBlockVersion11Enabled field if non-nil, zero value otherwise.

### GetTlsBlockVersion11EnabledOk

`func (o *Broker) GetTlsBlockVersion11EnabledOk() (*bool, bool)`

GetTlsBlockVersion11EnabledOk returns a tuple with the TlsBlockVersion11Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsBlockVersion11Enabled

`func (o *Broker) SetTlsBlockVersion11Enabled(v bool)`

SetTlsBlockVersion11Enabled sets TlsBlockVersion11Enabled field to given value.

### HasTlsBlockVersion11Enabled

`func (o *Broker) HasTlsBlockVersion11Enabled() bool`

HasTlsBlockVersion11Enabled returns a boolean if a field has been set.

### GetTlsCipherSuiteManagementList

`func (o *Broker) GetTlsCipherSuiteManagementList() string`

GetTlsCipherSuiteManagementList returns the TlsCipherSuiteManagementList field if non-nil, zero value otherwise.

### GetTlsCipherSuiteManagementListOk

`func (o *Broker) GetTlsCipherSuiteManagementListOk() (*string, bool)`

GetTlsCipherSuiteManagementListOk returns a tuple with the TlsCipherSuiteManagementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCipherSuiteManagementList

`func (o *Broker) SetTlsCipherSuiteManagementList(v string)`

SetTlsCipherSuiteManagementList sets TlsCipherSuiteManagementList field to given value.

### HasTlsCipherSuiteManagementList

`func (o *Broker) HasTlsCipherSuiteManagementList() bool`

HasTlsCipherSuiteManagementList returns a boolean if a field has been set.

### GetTlsCipherSuiteMsgBackboneList

`func (o *Broker) GetTlsCipherSuiteMsgBackboneList() string`

GetTlsCipherSuiteMsgBackboneList returns the TlsCipherSuiteMsgBackboneList field if non-nil, zero value otherwise.

### GetTlsCipherSuiteMsgBackboneListOk

`func (o *Broker) GetTlsCipherSuiteMsgBackboneListOk() (*string, bool)`

GetTlsCipherSuiteMsgBackboneListOk returns a tuple with the TlsCipherSuiteMsgBackboneList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCipherSuiteMsgBackboneList

`func (o *Broker) SetTlsCipherSuiteMsgBackboneList(v string)`

SetTlsCipherSuiteMsgBackboneList sets TlsCipherSuiteMsgBackboneList field to given value.

### HasTlsCipherSuiteMsgBackboneList

`func (o *Broker) HasTlsCipherSuiteMsgBackboneList() bool`

HasTlsCipherSuiteMsgBackboneList returns a boolean if a field has been set.

### GetTlsCipherSuiteSecureShellList

`func (o *Broker) GetTlsCipherSuiteSecureShellList() string`

GetTlsCipherSuiteSecureShellList returns the TlsCipherSuiteSecureShellList field if non-nil, zero value otherwise.

### GetTlsCipherSuiteSecureShellListOk

`func (o *Broker) GetTlsCipherSuiteSecureShellListOk() (*string, bool)`

GetTlsCipherSuiteSecureShellListOk returns a tuple with the TlsCipherSuiteSecureShellList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCipherSuiteSecureShellList

`func (o *Broker) SetTlsCipherSuiteSecureShellList(v string)`

SetTlsCipherSuiteSecureShellList sets TlsCipherSuiteSecureShellList field to given value.

### HasTlsCipherSuiteSecureShellList

`func (o *Broker) HasTlsCipherSuiteSecureShellList() bool`

HasTlsCipherSuiteSecureShellList returns a boolean if a field has been set.

### GetTlsCrimeExploitProtectionEnabled

`func (o *Broker) GetTlsCrimeExploitProtectionEnabled() bool`

GetTlsCrimeExploitProtectionEnabled returns the TlsCrimeExploitProtectionEnabled field if non-nil, zero value otherwise.

### GetTlsCrimeExploitProtectionEnabledOk

`func (o *Broker) GetTlsCrimeExploitProtectionEnabledOk() (*bool, bool)`

GetTlsCrimeExploitProtectionEnabledOk returns a tuple with the TlsCrimeExploitProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCrimeExploitProtectionEnabled

`func (o *Broker) SetTlsCrimeExploitProtectionEnabled(v bool)`

SetTlsCrimeExploitProtectionEnabled sets TlsCrimeExploitProtectionEnabled field to given value.

### HasTlsCrimeExploitProtectionEnabled

`func (o *Broker) HasTlsCrimeExploitProtectionEnabled() bool`

HasTlsCrimeExploitProtectionEnabled returns a boolean if a field has been set.

### GetTlsServerCertContent

`func (o *Broker) GetTlsServerCertContent() string`

GetTlsServerCertContent returns the TlsServerCertContent field if non-nil, zero value otherwise.

### GetTlsServerCertContentOk

`func (o *Broker) GetTlsServerCertContentOk() (*string, bool)`

GetTlsServerCertContentOk returns a tuple with the TlsServerCertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerCertContent

`func (o *Broker) SetTlsServerCertContent(v string)`

SetTlsServerCertContent sets TlsServerCertContent field to given value.

### HasTlsServerCertContent

`func (o *Broker) HasTlsServerCertContent() bool`

HasTlsServerCertContent returns a boolean if a field has been set.

### GetTlsServerCertPassword

`func (o *Broker) GetTlsServerCertPassword() string`

GetTlsServerCertPassword returns the TlsServerCertPassword field if non-nil, zero value otherwise.

### GetTlsServerCertPasswordOk

`func (o *Broker) GetTlsServerCertPasswordOk() (*string, bool)`

GetTlsServerCertPasswordOk returns a tuple with the TlsServerCertPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerCertPassword

`func (o *Broker) SetTlsServerCertPassword(v string)`

SetTlsServerCertPassword sets TlsServerCertPassword field to given value.

### HasTlsServerCertPassword

`func (o *Broker) HasTlsServerCertPassword() bool`

HasTlsServerCertPassword returns a boolean if a field has been set.

### GetTlsStandardDomainCertificateAuthoritiesEnabled

`func (o *Broker) GetTlsStandardDomainCertificateAuthoritiesEnabled() bool`

GetTlsStandardDomainCertificateAuthoritiesEnabled returns the TlsStandardDomainCertificateAuthoritiesEnabled field if non-nil, zero value otherwise.

### GetTlsStandardDomainCertificateAuthoritiesEnabledOk

`func (o *Broker) GetTlsStandardDomainCertificateAuthoritiesEnabledOk() (*bool, bool)`

GetTlsStandardDomainCertificateAuthoritiesEnabledOk returns a tuple with the TlsStandardDomainCertificateAuthoritiesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsStandardDomainCertificateAuthoritiesEnabled

`func (o *Broker) SetTlsStandardDomainCertificateAuthoritiesEnabled(v bool)`

SetTlsStandardDomainCertificateAuthoritiesEnabled sets TlsStandardDomainCertificateAuthoritiesEnabled field to given value.

### HasTlsStandardDomainCertificateAuthoritiesEnabled

`func (o *Broker) HasTlsStandardDomainCertificateAuthoritiesEnabled() bool`

HasTlsStandardDomainCertificateAuthoritiesEnabled returns a boolean if a field has been set.

### GetTlsTicketLifetime

`func (o *Broker) GetTlsTicketLifetime() int32`

GetTlsTicketLifetime returns the TlsTicketLifetime field if non-nil, zero value otherwise.

### GetTlsTicketLifetimeOk

`func (o *Broker) GetTlsTicketLifetimeOk() (*int32, bool)`

GetTlsTicketLifetimeOk returns a tuple with the TlsTicketLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTicketLifetime

`func (o *Broker) SetTlsTicketLifetime(v int32)`

SetTlsTicketLifetime sets TlsTicketLifetime field to given value.

### HasTlsTicketLifetime

`func (o *Broker) HasTlsTicketLifetime() bool`

HasTlsTicketLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


