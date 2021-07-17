# MsgVpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | The name of another Message VPN which this Message VPN is an alias for. When this Message VPN is enabled, the alias has no effect. When this Message VPN is disabled, Clients (but not Bridges and routing Links) logging into this Message VPN are automatically logged in to the other Message VPN, and authentication and authorization take place in the context of the other Message VPN.  Aliases may form a non-circular chain, cascading one to the next. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.14. | [optional] 
**AuthenticationBasicEnabled** | Pointer to **bool** | Enable or disable basic authentication for clients connecting to the Message VPN. Basic authentication is authentication that involves the use of a username and password to prove identity. If a user provides credentials for a different authentication scheme, this setting is not applicable. The default value is &#x60;true&#x60;. | [optional] 
**AuthenticationBasicProfileName** | Pointer to **string** | The name of the RADIUS or LDAP Profile to use for basic authentication. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 
**AuthenticationBasicRadiusDomain** | Pointer to **string** | The RADIUS domain to use for basic authentication. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**AuthenticationBasicType** | Pointer to **string** | The type of basic authentication to use for clients connecting to the Message VPN. The default value is &#x60;\&quot;radius\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;internal\&quot; - Internal database. Authentication is against Client Usernames. \&quot;ldap\&quot; - LDAP authentication. An LDAP profile name must be provided. \&quot;radius\&quot; - RADIUS authentication. A RADIUS profile name must be provided. \&quot;none\&quot; - No authentication. Anonymous login allowed. &lt;/pre&gt;  | [optional] 
**AuthenticationClientCertAllowApiProvidedUsernameEnabled** | Pointer to **bool** | Enable or disable allowing a client to specify a Client Username via the API connect method. When disabled, the certificate CN (Common Name) is always used. The default value is &#x60;false&#x60;. | [optional] 
**AuthenticationClientCertEnabled** | Pointer to **bool** | Enable or disable client certificate authentication in the Message VPN. The default value is &#x60;false&#x60;. | [optional] 
**AuthenticationClientCertMaxChainDepth** | Pointer to **int64** | The maximum depth for a client certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. The default value is &#x60;3&#x60;. | [optional] 
**AuthenticationClientCertRevocationCheckMode** | Pointer to **string** | The desired behavior for client certificate revocation checking. The default value is &#x60;\&quot;allow-valid\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;allow-all\&quot; - Allow the client to authenticate, the result of client certificate revocation check is ignored. \&quot;allow-unknown\&quot; - Allow the client to authenticate even if the revocation status of his certificate cannot be determined. \&quot;allow-valid\&quot; - Allow the client to authenticate only when the revocation check returned an explicit positive response. &lt;/pre&gt;  Available since 2.6. | [optional] 
**AuthenticationClientCertUsernameSource** | Pointer to **string** | The field from the client certificate to use as the client username. The default value is &#x60;\&quot;common-name\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;certificate-thumbprint\&quot; - The username is computed as the SHA-1 hash over the entire DER-encoded contents of the client certificate. \&quot;common-name\&quot; - The username is extracted from the certificate&#39;s first instance of the Common Name attribute in the Subject DN. \&quot;common-name-last\&quot; - The username is extracted from the certificate&#39;s last instance of the Common Name attribute in the Subject DN. \&quot;subject-alternate-name-msupn\&quot; - The username is extracted from the certificate&#39;s Other Name type of the Subject Alternative Name and must have the msUPN signature. \&quot;uid\&quot; - The username is extracted from the certificate&#39;s first instance of the User Identifier attribute in the Subject DN. \&quot;uid-last\&quot; - The username is extracted from the certificate&#39;s last instance of the User Identifier attribute in the Subject DN. &lt;/pre&gt;  Available since 2.6. | [optional] 
**AuthenticationClientCertValidateDateEnabled** | Pointer to **bool** | Enable or disable validation of the \&quot;Not Before\&quot; and \&quot;Not After\&quot; validity dates in the client certificate. The default value is &#x60;true&#x60;. | [optional] 
**AuthenticationKerberosAllowApiProvidedUsernameEnabled** | Pointer to **bool** | Enable or disable allowing a client to specify a Client Username via the API connect method. When disabled, the Kerberos Principal name is always used. The default value is &#x60;false&#x60;. | [optional] 
**AuthenticationKerberosEnabled** | Pointer to **bool** | Enable or disable Kerberos authentication in the Message VPN. The default value is &#x60;false&#x60;. | [optional] 
**AuthenticationOauthDefaultProviderName** | Pointer to **string** | The name of the provider to use when the client does not supply a provider name. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.13. | [optional] 
**AuthenticationOauthEnabled** | Pointer to **bool** | Enable or disable OAuth authentication. The default value is &#x60;false&#x60;. Available since 2.13. | [optional] 
**AuthorizationLdapGroupMembershipAttributeName** | Pointer to **string** | The name of the attribute that is retrieved from the LDAP server as part of the LDAP search when authorizing a client connecting to the Message VPN. The default value is &#x60;\&quot;memberOf\&quot;&#x60;. | [optional] 
**AuthorizationLdapTrimClientUsernameDomainEnabled** | Pointer to **bool** | Enable or disable client-username domain trimming for LDAP lookups of client connections. When enabled, the value of $CLIENT_USERNAME (when used for searching) will be truncated at the first occurance of the @ character. For example, if the client-username is in the form of an email address, then the domain portion will be removed. The default value is &#x60;false&#x60;. Available since 2.13. | [optional] 
**AuthorizationProfileName** | Pointer to **string** | The name of the LDAP Profile to use for client authorization. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**AuthorizationType** | Pointer to **string** | The type of authorization to use for clients connecting to the Message VPN. The default value is &#x60;\&quot;internal\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;ldap\&quot; - LDAP authorization. \&quot;internal\&quot; - Internal authorization. &lt;/pre&gt;  | [optional] 
**BridgingTlsServerCertEnforceTrustedCommonNameEnabled** | Pointer to **bool** | Enable or disable validation of the Common Name (CN) in the server certificate from the remote broker. If enabled, the Common Name is checked against the list of Trusted Common Names configured for the Bridge. Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is enabled. The default value is &#x60;true&#x60;. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**BridgingTlsServerCertMaxChainDepth** | Pointer to **int64** | The maximum depth for a server certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. The default value is &#x60;3&#x60;. | [optional] 
**BridgingTlsServerCertValidateDateEnabled** | Pointer to **bool** | Enable or disable validation of the \&quot;Not Before\&quot; and \&quot;Not After\&quot; validity dates in the server certificate. When disabled, a certificate will be accepted even if the certificate is not valid based on these dates. The default value is &#x60;true&#x60;. | [optional] 
**BridgingTlsServerCertValidateNameEnabled** | Pointer to **bool** | Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the bridge. If enabled, the name used to connect to the bridge is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. The default value is &#x60;true&#x60;. Available since 2.18. | [optional] 
**DistributedCacheManagementEnabled** | Pointer to **bool** | Enable or disable managing of cache instances over the message bus. The default value is &#x60;true&#x60;. | [optional] 
**DmrEnabled** | Pointer to **bool** | Enable or disable Dynamic Message Routing (DMR) for the Message VPN. The default value is &#x60;false&#x60;. Available since 2.11. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the Message VPN. The default value is &#x60;false&#x60;. | [optional] 
**EventConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventEgressFlowCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventEgressMsgRateThreshold** | Pointer to [**EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] 
**EventEndpointCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventIngressFlowCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventIngressMsgRateThreshold** | Pointer to [**EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] 
**EventLargeMsgThreshold** | Pointer to **int64** | The threshold, in kilobytes, after which a message is considered to be large for the Message VPN. The default value is &#x60;1024&#x60;. | [optional] 
**EventLogTag** | Pointer to **string** | A prefix applied to all published Events in the Message VPN. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**EventMsgSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventPublishClientEnabled** | Pointer to **bool** | Enable or disable Client level Event message publishing. The default value is &#x60;false&#x60;. | [optional] 
**EventPublishMsgVpnEnabled** | Pointer to **bool** | Enable or disable Message VPN level Event message publishing. The default value is &#x60;false&#x60;. | [optional] 
**EventPublishSubscriptionMode** | Pointer to **string** | Subscription level Event message publishing mode. The default value is &#x60;\&quot;off\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;off\&quot; - Disable client level event message publishing. \&quot;on-with-format-v1\&quot; - Enable client level event message publishing with format v1. \&quot;on-with-no-unsubscribe-events-on-disconnect-format-v1\&quot; - As \&quot;on-with-format-v1\&quot;, but unsubscribe events are not generated when a client disconnects. Unsubscribe events are still raised when a client explicitly unsubscribes from its subscriptions. \&quot;on-with-format-v2\&quot; - Enable client level event message publishing with format v2. \&quot;on-with-no-unsubscribe-events-on-disconnect-format-v2\&quot; - As \&quot;on-with-format-v2\&quot;, but unsubscribe events are not generated when a client disconnects. Unsubscribe events are still raised when a client explicitly unsubscribes from its subscriptions. &lt;/pre&gt;  | [optional] 
**EventPublishTopicFormatMqttEnabled** | Pointer to **bool** | Enable or disable Event publish topics in MQTT format. The default value is &#x60;false&#x60;. | [optional] 
**EventPublishTopicFormatSmfEnabled** | Pointer to **bool** | Enable or disable Event publish topics in SMF format. The default value is &#x60;true&#x60;. | [optional] 
**EventServiceAmqpConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventServiceMqttConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventServiceRestIncomingConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventServiceSmfConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventServiceWebConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventSubscriptionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventTransactedSessionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventTransactionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**ExportSubscriptionsEnabled** | Pointer to **bool** | Enable or disable the export of subscriptions in the Message VPN to other routers in the network over Neighbor links. The default value is &#x60;false&#x60;. | [optional] 
**JndiEnabled** | Pointer to **bool** | Enable or disable JNDI access for clients in the Message VPN. The default value is &#x60;false&#x60;. Available since 2.2. | [optional] 
**MaxConnectionCount** | Pointer to **int64** | The maximum number of client connections to the Message VPN. The default is the maximum value supported by the platform. | [optional] 
**MaxEgressFlowCount** | Pointer to **int64** | The maximum number of transmit flows that can be created in the Message VPN. The default value is &#x60;1000&#x60;. | [optional] 
**MaxEndpointCount** | Pointer to **int64** | The maximum number of Queues and Topic Endpoints that can be created in the Message VPN. The default value is &#x60;1000&#x60;. | [optional] 
**MaxIngressFlowCount** | Pointer to **int64** | The maximum number of receive flows that can be created in the Message VPN. The default value is &#x60;1000&#x60;. | [optional] 
**MaxMsgSpoolUsage** | Pointer to **int64** | The maximum message spool usage by the Message VPN, in megabytes. The default value is &#x60;0&#x60;. | [optional] 
**MaxSubscriptionCount** | Pointer to **int64** | The maximum number of local client subscriptions that can be added to the Message VPN. This limit is not enforced when a subscription is added using a management interface, such as CLI or SEMP. The default varies by platform. | [optional] 
**MaxTransactedSessionCount** | Pointer to **int64** | The maximum number of transacted sessions that can be created in the Message VPN. The default varies by platform. | [optional] 
**MaxTransactionCount** | Pointer to **int64** | The maximum number of transactions that can be created in the Message VPN. The default varies by platform. | [optional] 
**MqttRetainMaxMemory** | Pointer to **int32** | The maximum total memory usage of the MQTT Retain feature for this Message VPN, in MB. If the maximum memory is reached, any arriving retain messages that require more memory are discarded. A value of -1 indicates that the memory is bounded only by the global max memory limit. A value of 0 prevents MQTT Retain from becoming operational. The default value is &#x60;-1&#x60;. Available since 2.11. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**ReplicationAckPropagationIntervalMsgCount** | Pointer to **int64** | The acknowledgement (ACK) propagation interval for the replication Bridge, in number of replicated messages. The default value is &#x60;20&#x60;. | [optional] 
**ReplicationBridgeAuthenticationBasicClientUsername** | Pointer to **string** | The Client Username the replication Bridge uses to login to the remote Message VPN. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**ReplicationBridgeAuthenticationBasicPassword** | Pointer to **string** | The password for the Client Username. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**ReplicationBridgeAuthenticationClientCertContent** | Pointer to **string** | The PEM formatted content for the client certificate used by this bridge to login to the Remote Message VPN. It must consist of a private key and between one and three certificates comprising the certificate trust chain. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.9. | [optional] 
**ReplicationBridgeAuthenticationClientCertPassword** | Pointer to **string** | The password for the client certificate. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.9. | [optional] 
**ReplicationBridgeAuthenticationScheme** | Pointer to **string** | The authentication scheme for the replication Bridge in the Message VPN. The default value is &#x60;\&quot;basic\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;basic\&quot; - Basic Authentication Scheme (via username and password). \&quot;client-certificate\&quot; - Client Certificate Authentication Scheme (via certificate file or content). &lt;/pre&gt;  | [optional] 
**ReplicationBridgeCompressedDataEnabled** | Pointer to **bool** | Enable or disable use of compression for the replication Bridge. The default value is &#x60;false&#x60;. | [optional] 
**ReplicationBridgeEgressFlowWindowSize** | Pointer to **int64** | The size of the window used for guaranteed messages published to the replication Bridge, in messages. The default value is &#x60;255&#x60;. | [optional] 
**ReplicationBridgeRetryDelay** | Pointer to **int64** | The number of seconds that must pass before retrying the replication Bridge connection. The default value is &#x60;3&#x60;. | [optional] 
**ReplicationBridgeTlsEnabled** | Pointer to **bool** | Enable or disable use of encryption (TLS) for the replication Bridge connection. The default value is &#x60;false&#x60;. | [optional] 
**ReplicationBridgeUnidirectionalClientProfileName** | Pointer to **string** | The Client Profile for the unidirectional replication Bridge in the Message VPN. It is used only for the TCP parameters. The default value is &#x60;\&quot;#client-profile\&quot;&#x60;. | [optional] 
**ReplicationEnabled** | Pointer to **bool** | Enable or disable replication for the Message VPN. The default value is &#x60;false&#x60;. | [optional] 
**ReplicationEnabledQueueBehavior** | Pointer to **string** | The behavior to take when enabling replication for the Message VPN, depending on the existence of the replication Queue. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;fail-on-existing-queue\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;fail-on-existing-queue\&quot; - The data replication queue must not already exist. \&quot;force-use-existing-queue\&quot; - The data replication queue must already exist. Any data messages on the Queue will be forwarded to interested applications. IMPORTANT: Before using this mode be certain that the messages are not stale or otherwise unsuitable to be forwarded. This mode can only be specified when the existing queue is configured the same as is currently specified under replication configuration otherwise the enabling of replication will fail. \&quot;force-recreate-queue\&quot; - The data replication queue must already exist. Any data messages on the Queue will be discarded. IMPORTANT: Before using this mode be certain that the messages on the existing data replication queue are not needed by interested applications. &lt;/pre&gt;  | [optional] 
**ReplicationQueueMaxMsgSpoolUsage** | Pointer to **int64** | The maximum message spool usage by the replication Bridge local Queue (quota), in megabytes. The default value is &#x60;60000&#x60;. | [optional] 
**ReplicationQueueRejectMsgToSenderOnDiscardEnabled** | Pointer to **bool** | Enable or disable whether messages discarded on the replication Bridge local Queue are rejected back to the sender. The default value is &#x60;true&#x60;. | [optional] 
**ReplicationRejectMsgWhenSyncIneligibleEnabled** | Pointer to **bool** | Enable or disable whether guaranteed messages published to synchronously replicated Topics are rejected back to the sender when synchronous replication becomes ineligible. The default value is &#x60;false&#x60;. | [optional] 
**ReplicationRole** | Pointer to **string** | The replication role for the Message VPN. The default value is &#x60;\&quot;standby\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;active\&quot; - Assume the Active role in replication for the Message VPN. \&quot;standby\&quot; - Assume the Standby role in replication for the Message VPN. &lt;/pre&gt;  | [optional] 
**ReplicationTransactionMode** | Pointer to **string** | The transaction replication mode for all transactions within the Message VPN. Changing this value during operation will not affect existing transactions; it is only used upon starting a transaction. The default value is &#x60;\&quot;async\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;sync\&quot; - Messages are acknowledged when replicated (spooled remotely). \&quot;async\&quot; - Messages are acknowledged when pending replication (spooled locally). &lt;/pre&gt;  | [optional] 
**RestTlsServerCertEnforceTrustedCommonNameEnabled** | Pointer to **bool** | Enable or disable validation of the Common Name (CN) in the server certificate from the remote REST Consumer. If enabled, the Common Name is checked against the list of Trusted Common Names configured for the REST Consumer. Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is enabled. The default value is &#x60;true&#x60;. Deprecated since 2.17. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**RestTlsServerCertMaxChainDepth** | Pointer to **int64** | The maximum depth for a REST Consumer server certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. The default value is &#x60;3&#x60;. | [optional] 
**RestTlsServerCertValidateDateEnabled** | Pointer to **bool** | Enable or disable validation of the \&quot;Not Before\&quot; and \&quot;Not After\&quot; validity dates in the REST Consumer server certificate. The default value is &#x60;true&#x60;. | [optional] 
**RestTlsServerCertValidateNameEnabled** | Pointer to **bool** | Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the remote REST Consumer. If enabled, the name used to connect to the remote REST Consumer is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. The default value is &#x60;true&#x60;. Available since 2.17. | [optional] 
**SempOverMsgBusAdminClientEnabled** | Pointer to **bool** | Enable or disable \&quot;admin client\&quot; SEMP over the message bus commands for the current Message VPN. The default value is &#x60;false&#x60;. | [optional] 
**SempOverMsgBusAdminDistributedCacheEnabled** | Pointer to **bool** | Enable or disable \&quot;admin distributed-cache\&quot; SEMP over the message bus commands for the current Message VPN. The default value is &#x60;false&#x60;. | [optional] 
**SempOverMsgBusAdminEnabled** | Pointer to **bool** | Enable or disable \&quot;admin\&quot; SEMP over the message bus commands for the current Message VPN. The default value is &#x60;false&#x60;. | [optional] 
**SempOverMsgBusEnabled** | Pointer to **bool** | Enable or disable SEMP over the message bus for the current Message VPN. The default value is &#x60;true&#x60;. | [optional] 
**SempOverMsgBusShowEnabled** | Pointer to **bool** | Enable or disable \&quot;show\&quot; SEMP over the message bus commands for the current Message VPN. The default value is &#x60;false&#x60;. | [optional] 
**ServiceAmqpMaxConnectionCount** | Pointer to **int64** | The maximum number of AMQP client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform. The default is the maximum value supported by the platform. Available since 2.7. | [optional] 
**ServiceAmqpPlainTextEnabled** | Pointer to **bool** | Enable or disable the plain-text AMQP service in the Message VPN. Disabling causes clients connected to the corresponding listen-port to be disconnected. The default value is &#x60;false&#x60;. Available since 2.7. | [optional] 
**ServiceAmqpPlainTextListenPort** | Pointer to **int64** | The port number for plain-text AMQP clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. The default value is &#x60;0&#x60;. Available since 2.7. | [optional] 
**ServiceAmqpTlsEnabled** | Pointer to **bool** | Enable or disable the use of encryption (TLS) for the AMQP service in the Message VPN. Disabling causes clients currently connected over TLS to be disconnected. The default value is &#x60;false&#x60;. Available since 2.7. | [optional] 
**ServiceAmqpTlsListenPort** | Pointer to **int64** | The port number for AMQP clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. The default value is &#x60;0&#x60;. Available since 2.7. | [optional] 
**ServiceMqttAuthenticationClientCertRequest** | Pointer to **string** | Determines when to request a client certificate from an incoming MQTT client connecting via a TLS port. The default value is &#x60;\&quot;when-enabled-in-message-vpn\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;never\&quot; - Never ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;when-enabled-in-message-vpn\&quot; - Only ask for a client-certificate if client certificate authentication is enabled under \&quot;message-vpn &gt;  authentication &gt; client-certificate &gt; shutdown\&quot;. &lt;/pre&gt;  Available since 2.21. | [optional] 
**ServiceMqttMaxConnectionCount** | Pointer to **int64** | The maximum number of MQTT client connections that can be simultaneously connected to the Message VPN. The default is the maximum value supported by the platform. Available since 2.1. | [optional] 
**ServiceMqttPlainTextEnabled** | Pointer to **bool** | Enable or disable the plain-text MQTT service in the Message VPN. Disabling causes clients currently connected to be disconnected. The default value is &#x60;false&#x60;. Available since 2.1. | [optional] 
**ServiceMqttPlainTextListenPort** | Pointer to **int64** | The port number for plain-text MQTT clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. The default value is &#x60;0&#x60;. Available since 2.1. | [optional] 
**ServiceMqttTlsEnabled** | Pointer to **bool** | Enable or disable the use of encryption (TLS) for the MQTT service in the Message VPN. Disabling causes clients currently connected over TLS to be disconnected. The default value is &#x60;false&#x60;. Available since 2.1. | [optional] 
**ServiceMqttTlsListenPort** | Pointer to **int64** | The port number for MQTT clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. The default value is &#x60;0&#x60;. Available since 2.1. | [optional] 
**ServiceMqttTlsWebSocketEnabled** | Pointer to **bool** | Enable or disable the use of encrypted WebSocket (WebSocket over TLS) for the MQTT service in the Message VPN. Disabling causes clients currently connected by encrypted WebSocket to be disconnected. The default value is &#x60;false&#x60;. Available since 2.1. | [optional] 
**ServiceMqttTlsWebSocketListenPort** | Pointer to **int64** | The port number for MQTT clients that connect to the Message VPN using WebSocket over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. The default value is &#x60;0&#x60;. Available since 2.1. | [optional] 
**ServiceMqttWebSocketEnabled** | Pointer to **bool** | Enable or disable the use of WebSocket for the MQTT service in the Message VPN. Disabling causes clients currently connected by WebSocket to be disconnected. The default value is &#x60;false&#x60;. Available since 2.1. | [optional] 
**ServiceMqttWebSocketListenPort** | Pointer to **int64** | The port number for plain-text MQTT clients that connect to the Message VPN using WebSocket. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. The default value is &#x60;0&#x60;. Available since 2.1. | [optional] 
**ServiceRestIncomingAuthenticationClientCertRequest** | Pointer to **string** | Determines when to request a client certificate from an incoming REST Producer connecting via a TLS port. The default value is &#x60;\&quot;when-enabled-in-message-vpn\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;never\&quot; - Never ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;when-enabled-in-message-vpn\&quot; - Only ask for a client-certificate if client certificate authentication is enabled under \&quot;message-vpn &gt;  authentication &gt; client-certificate &gt; shutdown\&quot;. &lt;/pre&gt;  Available since 2.21. | [optional] 
**ServiceRestIncomingAuthorizationHeaderHandling** | Pointer to **string** | The handling of Authorization headers for incoming REST connections. The default value is &#x60;\&quot;drop\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;drop\&quot; - Do not attach the Authorization header to the message as a user property. This configuration is most secure. \&quot;forward\&quot; - Forward the Authorization header, attaching it to the message as a user property in the same way as other headers. For best security, use the drop setting. \&quot;legacy\&quot; - If the Authorization header was used for authentication to the broker, do not attach it to the message. If the Authorization header was not used for authentication to the broker, attach it to the message as a user property in the same way as other headers. For best security, use the drop setting. &lt;/pre&gt;  Available since 2.19. | [optional] 
**ServiceRestIncomingMaxConnectionCount** | Pointer to **int64** | The maximum number of REST incoming client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform. The default is the maximum value supported by the platform. | [optional] 
**ServiceRestIncomingPlainTextEnabled** | Pointer to **bool** | Enable or disable the plain-text REST service for incoming clients in the Message VPN. Disabling causes clients currently connected to be disconnected. The default value is &#x60;false&#x60;. | [optional] 
**ServiceRestIncomingPlainTextListenPort** | Pointer to **int64** | The port number for incoming plain-text REST clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. The default value is &#x60;0&#x60;. | [optional] 
**ServiceRestIncomingTlsEnabled** | Pointer to **bool** | Enable or disable the use of encryption (TLS) for the REST service for incoming clients in the Message VPN. Disabling causes clients currently connected over TLS to be disconnected. The default value is &#x60;false&#x60;. | [optional] 
**ServiceRestIncomingTlsListenPort** | Pointer to **int64** | The port number for incoming REST clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. The default value is &#x60;0&#x60;. | [optional] 
**ServiceRestMode** | Pointer to **string** | The REST service mode for incoming REST clients that connect to the Message VPN. The default value is &#x60;\&quot;messaging\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;gateway\&quot; - Act as a message gateway through which REST messages are propagated. \&quot;messaging\&quot; - Act as a message broker on which REST messages are queued. &lt;/pre&gt;  Available since 2.6. | [optional] 
**ServiceRestOutgoingMaxConnectionCount** | Pointer to **int64** | The maximum number of REST Consumer (outgoing) client connections that can be simultaneously connected to the Message VPN. The default varies by platform. | [optional] 
**ServiceSmfMaxConnectionCount** | Pointer to **int64** | The maximum number of SMF client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform. The default varies by platform. | [optional] 
**ServiceSmfPlainTextEnabled** | Pointer to **bool** | Enable or disable the plain-text SMF service in the Message VPN. Disabling causes clients currently connected to be disconnected. The default value is &#x60;true&#x60;. | [optional] 
**ServiceSmfTlsEnabled** | Pointer to **bool** | Enable or disable the use of encryption (TLS) for the SMF service in the Message VPN. Disabling causes clients currently connected over TLS to be disconnected. The default value is &#x60;true&#x60;. | [optional] 
**ServiceWebAuthenticationClientCertRequest** | Pointer to **string** | Determines when to request a client certificate from a Web Transport client connecting via a TLS port. The default value is &#x60;\&quot;when-enabled-in-message-vpn\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;never\&quot; - Never ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;when-enabled-in-message-vpn\&quot; - Only ask for a client-certificate if client certificate authentication is enabled under \&quot;message-vpn &gt;  authentication &gt; client-certificate &gt; shutdown\&quot;. &lt;/pre&gt;  Available since 2.21. | [optional] 
**ServiceWebMaxConnectionCount** | Pointer to **int64** | The maximum number of Web Transport client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform. The default is the maximum value supported by the platform. | [optional] 
**ServiceWebPlainTextEnabled** | Pointer to **bool** | Enable or disable the plain-text Web Transport service in the Message VPN. Disabling causes clients currently connected to be disconnected. The default value is &#x60;true&#x60;. | [optional] 
**ServiceWebTlsEnabled** | Pointer to **bool** | Enable or disable the use of TLS for the Web Transport service in the Message VPN. Disabling causes clients currently connected over TLS to be disconnected. The default value is &#x60;true&#x60;. | [optional] 
**TlsAllowDowngradeToPlainTextEnabled** | Pointer to **bool** | Enable or disable the allowing of TLS SMF clients to downgrade their connections to plain-text connections. Changing this will not affect existing connections. The default value is &#x60;false&#x60;. | [optional] 

## Methods

### NewMsgVpn

`func NewMsgVpn() *MsgVpn`

NewMsgVpn instantiates a new MsgVpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnWithDefaults

`func NewMsgVpnWithDefaults() *MsgVpn`

NewMsgVpnWithDefaults instantiates a new MsgVpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *MsgVpn) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *MsgVpn) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *MsgVpn) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *MsgVpn) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAuthenticationBasicEnabled

`func (o *MsgVpn) GetAuthenticationBasicEnabled() bool`

GetAuthenticationBasicEnabled returns the AuthenticationBasicEnabled field if non-nil, zero value otherwise.

### GetAuthenticationBasicEnabledOk

`func (o *MsgVpn) GetAuthenticationBasicEnabledOk() (*bool, bool)`

GetAuthenticationBasicEnabledOk returns a tuple with the AuthenticationBasicEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBasicEnabled

`func (o *MsgVpn) SetAuthenticationBasicEnabled(v bool)`

SetAuthenticationBasicEnabled sets AuthenticationBasicEnabled field to given value.

### HasAuthenticationBasicEnabled

`func (o *MsgVpn) HasAuthenticationBasicEnabled() bool`

HasAuthenticationBasicEnabled returns a boolean if a field has been set.

### GetAuthenticationBasicProfileName

`func (o *MsgVpn) GetAuthenticationBasicProfileName() string`

GetAuthenticationBasicProfileName returns the AuthenticationBasicProfileName field if non-nil, zero value otherwise.

### GetAuthenticationBasicProfileNameOk

`func (o *MsgVpn) GetAuthenticationBasicProfileNameOk() (*string, bool)`

GetAuthenticationBasicProfileNameOk returns a tuple with the AuthenticationBasicProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBasicProfileName

`func (o *MsgVpn) SetAuthenticationBasicProfileName(v string)`

SetAuthenticationBasicProfileName sets AuthenticationBasicProfileName field to given value.

### HasAuthenticationBasicProfileName

`func (o *MsgVpn) HasAuthenticationBasicProfileName() bool`

HasAuthenticationBasicProfileName returns a boolean if a field has been set.

### GetAuthenticationBasicRadiusDomain

`func (o *MsgVpn) GetAuthenticationBasicRadiusDomain() string`

GetAuthenticationBasicRadiusDomain returns the AuthenticationBasicRadiusDomain field if non-nil, zero value otherwise.

### GetAuthenticationBasicRadiusDomainOk

`func (o *MsgVpn) GetAuthenticationBasicRadiusDomainOk() (*string, bool)`

GetAuthenticationBasicRadiusDomainOk returns a tuple with the AuthenticationBasicRadiusDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBasicRadiusDomain

`func (o *MsgVpn) SetAuthenticationBasicRadiusDomain(v string)`

SetAuthenticationBasicRadiusDomain sets AuthenticationBasicRadiusDomain field to given value.

### HasAuthenticationBasicRadiusDomain

`func (o *MsgVpn) HasAuthenticationBasicRadiusDomain() bool`

HasAuthenticationBasicRadiusDomain returns a boolean if a field has been set.

### GetAuthenticationBasicType

`func (o *MsgVpn) GetAuthenticationBasicType() string`

GetAuthenticationBasicType returns the AuthenticationBasicType field if non-nil, zero value otherwise.

### GetAuthenticationBasicTypeOk

`func (o *MsgVpn) GetAuthenticationBasicTypeOk() (*string, bool)`

GetAuthenticationBasicTypeOk returns a tuple with the AuthenticationBasicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationBasicType

`func (o *MsgVpn) SetAuthenticationBasicType(v string)`

SetAuthenticationBasicType sets AuthenticationBasicType field to given value.

### HasAuthenticationBasicType

`func (o *MsgVpn) HasAuthenticationBasicType() bool`

HasAuthenticationBasicType returns a boolean if a field has been set.

### GetAuthenticationClientCertAllowApiProvidedUsernameEnabled

`func (o *MsgVpn) GetAuthenticationClientCertAllowApiProvidedUsernameEnabled() bool`

GetAuthenticationClientCertAllowApiProvidedUsernameEnabled returns the AuthenticationClientCertAllowApiProvidedUsernameEnabled field if non-nil, zero value otherwise.

### GetAuthenticationClientCertAllowApiProvidedUsernameEnabledOk

`func (o *MsgVpn) GetAuthenticationClientCertAllowApiProvidedUsernameEnabledOk() (*bool, bool)`

GetAuthenticationClientCertAllowApiProvidedUsernameEnabledOk returns a tuple with the AuthenticationClientCertAllowApiProvidedUsernameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertAllowApiProvidedUsernameEnabled

`func (o *MsgVpn) SetAuthenticationClientCertAllowApiProvidedUsernameEnabled(v bool)`

SetAuthenticationClientCertAllowApiProvidedUsernameEnabled sets AuthenticationClientCertAllowApiProvidedUsernameEnabled field to given value.

### HasAuthenticationClientCertAllowApiProvidedUsernameEnabled

`func (o *MsgVpn) HasAuthenticationClientCertAllowApiProvidedUsernameEnabled() bool`

HasAuthenticationClientCertAllowApiProvidedUsernameEnabled returns a boolean if a field has been set.

### GetAuthenticationClientCertEnabled

`func (o *MsgVpn) GetAuthenticationClientCertEnabled() bool`

GetAuthenticationClientCertEnabled returns the AuthenticationClientCertEnabled field if non-nil, zero value otherwise.

### GetAuthenticationClientCertEnabledOk

`func (o *MsgVpn) GetAuthenticationClientCertEnabledOk() (*bool, bool)`

GetAuthenticationClientCertEnabledOk returns a tuple with the AuthenticationClientCertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertEnabled

`func (o *MsgVpn) SetAuthenticationClientCertEnabled(v bool)`

SetAuthenticationClientCertEnabled sets AuthenticationClientCertEnabled field to given value.

### HasAuthenticationClientCertEnabled

`func (o *MsgVpn) HasAuthenticationClientCertEnabled() bool`

HasAuthenticationClientCertEnabled returns a boolean if a field has been set.

### GetAuthenticationClientCertMaxChainDepth

`func (o *MsgVpn) GetAuthenticationClientCertMaxChainDepth() int64`

GetAuthenticationClientCertMaxChainDepth returns the AuthenticationClientCertMaxChainDepth field if non-nil, zero value otherwise.

### GetAuthenticationClientCertMaxChainDepthOk

`func (o *MsgVpn) GetAuthenticationClientCertMaxChainDepthOk() (*int64, bool)`

GetAuthenticationClientCertMaxChainDepthOk returns a tuple with the AuthenticationClientCertMaxChainDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertMaxChainDepth

`func (o *MsgVpn) SetAuthenticationClientCertMaxChainDepth(v int64)`

SetAuthenticationClientCertMaxChainDepth sets AuthenticationClientCertMaxChainDepth field to given value.

### HasAuthenticationClientCertMaxChainDepth

`func (o *MsgVpn) HasAuthenticationClientCertMaxChainDepth() bool`

HasAuthenticationClientCertMaxChainDepth returns a boolean if a field has been set.

### GetAuthenticationClientCertRevocationCheckMode

`func (o *MsgVpn) GetAuthenticationClientCertRevocationCheckMode() string`

GetAuthenticationClientCertRevocationCheckMode returns the AuthenticationClientCertRevocationCheckMode field if non-nil, zero value otherwise.

### GetAuthenticationClientCertRevocationCheckModeOk

`func (o *MsgVpn) GetAuthenticationClientCertRevocationCheckModeOk() (*string, bool)`

GetAuthenticationClientCertRevocationCheckModeOk returns a tuple with the AuthenticationClientCertRevocationCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertRevocationCheckMode

`func (o *MsgVpn) SetAuthenticationClientCertRevocationCheckMode(v string)`

SetAuthenticationClientCertRevocationCheckMode sets AuthenticationClientCertRevocationCheckMode field to given value.

### HasAuthenticationClientCertRevocationCheckMode

`func (o *MsgVpn) HasAuthenticationClientCertRevocationCheckMode() bool`

HasAuthenticationClientCertRevocationCheckMode returns a boolean if a field has been set.

### GetAuthenticationClientCertUsernameSource

`func (o *MsgVpn) GetAuthenticationClientCertUsernameSource() string`

GetAuthenticationClientCertUsernameSource returns the AuthenticationClientCertUsernameSource field if non-nil, zero value otherwise.

### GetAuthenticationClientCertUsernameSourceOk

`func (o *MsgVpn) GetAuthenticationClientCertUsernameSourceOk() (*string, bool)`

GetAuthenticationClientCertUsernameSourceOk returns a tuple with the AuthenticationClientCertUsernameSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertUsernameSource

`func (o *MsgVpn) SetAuthenticationClientCertUsernameSource(v string)`

SetAuthenticationClientCertUsernameSource sets AuthenticationClientCertUsernameSource field to given value.

### HasAuthenticationClientCertUsernameSource

`func (o *MsgVpn) HasAuthenticationClientCertUsernameSource() bool`

HasAuthenticationClientCertUsernameSource returns a boolean if a field has been set.

### GetAuthenticationClientCertValidateDateEnabled

`func (o *MsgVpn) GetAuthenticationClientCertValidateDateEnabled() bool`

GetAuthenticationClientCertValidateDateEnabled returns the AuthenticationClientCertValidateDateEnabled field if non-nil, zero value otherwise.

### GetAuthenticationClientCertValidateDateEnabledOk

`func (o *MsgVpn) GetAuthenticationClientCertValidateDateEnabledOk() (*bool, bool)`

GetAuthenticationClientCertValidateDateEnabledOk returns a tuple with the AuthenticationClientCertValidateDateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertValidateDateEnabled

`func (o *MsgVpn) SetAuthenticationClientCertValidateDateEnabled(v bool)`

SetAuthenticationClientCertValidateDateEnabled sets AuthenticationClientCertValidateDateEnabled field to given value.

### HasAuthenticationClientCertValidateDateEnabled

`func (o *MsgVpn) HasAuthenticationClientCertValidateDateEnabled() bool`

HasAuthenticationClientCertValidateDateEnabled returns a boolean if a field has been set.

### GetAuthenticationKerberosAllowApiProvidedUsernameEnabled

`func (o *MsgVpn) GetAuthenticationKerberosAllowApiProvidedUsernameEnabled() bool`

GetAuthenticationKerberosAllowApiProvidedUsernameEnabled returns the AuthenticationKerberosAllowApiProvidedUsernameEnabled field if non-nil, zero value otherwise.

### GetAuthenticationKerberosAllowApiProvidedUsernameEnabledOk

`func (o *MsgVpn) GetAuthenticationKerberosAllowApiProvidedUsernameEnabledOk() (*bool, bool)`

GetAuthenticationKerberosAllowApiProvidedUsernameEnabledOk returns a tuple with the AuthenticationKerberosAllowApiProvidedUsernameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKerberosAllowApiProvidedUsernameEnabled

`func (o *MsgVpn) SetAuthenticationKerberosAllowApiProvidedUsernameEnabled(v bool)`

SetAuthenticationKerberosAllowApiProvidedUsernameEnabled sets AuthenticationKerberosAllowApiProvidedUsernameEnabled field to given value.

### HasAuthenticationKerberosAllowApiProvidedUsernameEnabled

`func (o *MsgVpn) HasAuthenticationKerberosAllowApiProvidedUsernameEnabled() bool`

HasAuthenticationKerberosAllowApiProvidedUsernameEnabled returns a boolean if a field has been set.

### GetAuthenticationKerberosEnabled

`func (o *MsgVpn) GetAuthenticationKerberosEnabled() bool`

GetAuthenticationKerberosEnabled returns the AuthenticationKerberosEnabled field if non-nil, zero value otherwise.

### GetAuthenticationKerberosEnabledOk

`func (o *MsgVpn) GetAuthenticationKerberosEnabledOk() (*bool, bool)`

GetAuthenticationKerberosEnabledOk returns a tuple with the AuthenticationKerberosEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKerberosEnabled

`func (o *MsgVpn) SetAuthenticationKerberosEnabled(v bool)`

SetAuthenticationKerberosEnabled sets AuthenticationKerberosEnabled field to given value.

### HasAuthenticationKerberosEnabled

`func (o *MsgVpn) HasAuthenticationKerberosEnabled() bool`

HasAuthenticationKerberosEnabled returns a boolean if a field has been set.

### GetAuthenticationOauthDefaultProviderName

`func (o *MsgVpn) GetAuthenticationOauthDefaultProviderName() string`

GetAuthenticationOauthDefaultProviderName returns the AuthenticationOauthDefaultProviderName field if non-nil, zero value otherwise.

### GetAuthenticationOauthDefaultProviderNameOk

`func (o *MsgVpn) GetAuthenticationOauthDefaultProviderNameOk() (*string, bool)`

GetAuthenticationOauthDefaultProviderNameOk returns a tuple with the AuthenticationOauthDefaultProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthDefaultProviderName

`func (o *MsgVpn) SetAuthenticationOauthDefaultProviderName(v string)`

SetAuthenticationOauthDefaultProviderName sets AuthenticationOauthDefaultProviderName field to given value.

### HasAuthenticationOauthDefaultProviderName

`func (o *MsgVpn) HasAuthenticationOauthDefaultProviderName() bool`

HasAuthenticationOauthDefaultProviderName returns a boolean if a field has been set.

### GetAuthenticationOauthEnabled

`func (o *MsgVpn) GetAuthenticationOauthEnabled() bool`

GetAuthenticationOauthEnabled returns the AuthenticationOauthEnabled field if non-nil, zero value otherwise.

### GetAuthenticationOauthEnabledOk

`func (o *MsgVpn) GetAuthenticationOauthEnabledOk() (*bool, bool)`

GetAuthenticationOauthEnabledOk returns a tuple with the AuthenticationOauthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthEnabled

`func (o *MsgVpn) SetAuthenticationOauthEnabled(v bool)`

SetAuthenticationOauthEnabled sets AuthenticationOauthEnabled field to given value.

### HasAuthenticationOauthEnabled

`func (o *MsgVpn) HasAuthenticationOauthEnabled() bool`

HasAuthenticationOauthEnabled returns a boolean if a field has been set.

### GetAuthorizationLdapGroupMembershipAttributeName

`func (o *MsgVpn) GetAuthorizationLdapGroupMembershipAttributeName() string`

GetAuthorizationLdapGroupMembershipAttributeName returns the AuthorizationLdapGroupMembershipAttributeName field if non-nil, zero value otherwise.

### GetAuthorizationLdapGroupMembershipAttributeNameOk

`func (o *MsgVpn) GetAuthorizationLdapGroupMembershipAttributeNameOk() (*string, bool)`

GetAuthorizationLdapGroupMembershipAttributeNameOk returns a tuple with the AuthorizationLdapGroupMembershipAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLdapGroupMembershipAttributeName

`func (o *MsgVpn) SetAuthorizationLdapGroupMembershipAttributeName(v string)`

SetAuthorizationLdapGroupMembershipAttributeName sets AuthorizationLdapGroupMembershipAttributeName field to given value.

### HasAuthorizationLdapGroupMembershipAttributeName

`func (o *MsgVpn) HasAuthorizationLdapGroupMembershipAttributeName() bool`

HasAuthorizationLdapGroupMembershipAttributeName returns a boolean if a field has been set.

### GetAuthorizationLdapTrimClientUsernameDomainEnabled

`func (o *MsgVpn) GetAuthorizationLdapTrimClientUsernameDomainEnabled() bool`

GetAuthorizationLdapTrimClientUsernameDomainEnabled returns the AuthorizationLdapTrimClientUsernameDomainEnabled field if non-nil, zero value otherwise.

### GetAuthorizationLdapTrimClientUsernameDomainEnabledOk

`func (o *MsgVpn) GetAuthorizationLdapTrimClientUsernameDomainEnabledOk() (*bool, bool)`

GetAuthorizationLdapTrimClientUsernameDomainEnabledOk returns a tuple with the AuthorizationLdapTrimClientUsernameDomainEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationLdapTrimClientUsernameDomainEnabled

`func (o *MsgVpn) SetAuthorizationLdapTrimClientUsernameDomainEnabled(v bool)`

SetAuthorizationLdapTrimClientUsernameDomainEnabled sets AuthorizationLdapTrimClientUsernameDomainEnabled field to given value.

### HasAuthorizationLdapTrimClientUsernameDomainEnabled

`func (o *MsgVpn) HasAuthorizationLdapTrimClientUsernameDomainEnabled() bool`

HasAuthorizationLdapTrimClientUsernameDomainEnabled returns a boolean if a field has been set.

### GetAuthorizationProfileName

`func (o *MsgVpn) GetAuthorizationProfileName() string`

GetAuthorizationProfileName returns the AuthorizationProfileName field if non-nil, zero value otherwise.

### GetAuthorizationProfileNameOk

`func (o *MsgVpn) GetAuthorizationProfileNameOk() (*string, bool)`

GetAuthorizationProfileNameOk returns a tuple with the AuthorizationProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationProfileName

`func (o *MsgVpn) SetAuthorizationProfileName(v string)`

SetAuthorizationProfileName sets AuthorizationProfileName field to given value.

### HasAuthorizationProfileName

`func (o *MsgVpn) HasAuthorizationProfileName() bool`

HasAuthorizationProfileName returns a boolean if a field has been set.

### GetAuthorizationType

`func (o *MsgVpn) GetAuthorizationType() string`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *MsgVpn) GetAuthorizationTypeOk() (*string, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *MsgVpn) SetAuthorizationType(v string)`

SetAuthorizationType sets AuthorizationType field to given value.

### HasAuthorizationType

`func (o *MsgVpn) HasAuthorizationType() bool`

HasAuthorizationType returns a boolean if a field has been set.

### GetBridgingTlsServerCertEnforceTrustedCommonNameEnabled

`func (o *MsgVpn) GetBridgingTlsServerCertEnforceTrustedCommonNameEnabled() bool`

GetBridgingTlsServerCertEnforceTrustedCommonNameEnabled returns the BridgingTlsServerCertEnforceTrustedCommonNameEnabled field if non-nil, zero value otherwise.

### GetBridgingTlsServerCertEnforceTrustedCommonNameEnabledOk

`func (o *MsgVpn) GetBridgingTlsServerCertEnforceTrustedCommonNameEnabledOk() (*bool, bool)`

GetBridgingTlsServerCertEnforceTrustedCommonNameEnabledOk returns a tuple with the BridgingTlsServerCertEnforceTrustedCommonNameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgingTlsServerCertEnforceTrustedCommonNameEnabled

`func (o *MsgVpn) SetBridgingTlsServerCertEnforceTrustedCommonNameEnabled(v bool)`

SetBridgingTlsServerCertEnforceTrustedCommonNameEnabled sets BridgingTlsServerCertEnforceTrustedCommonNameEnabled field to given value.

### HasBridgingTlsServerCertEnforceTrustedCommonNameEnabled

`func (o *MsgVpn) HasBridgingTlsServerCertEnforceTrustedCommonNameEnabled() bool`

HasBridgingTlsServerCertEnforceTrustedCommonNameEnabled returns a boolean if a field has been set.

### GetBridgingTlsServerCertMaxChainDepth

`func (o *MsgVpn) GetBridgingTlsServerCertMaxChainDepth() int64`

GetBridgingTlsServerCertMaxChainDepth returns the BridgingTlsServerCertMaxChainDepth field if non-nil, zero value otherwise.

### GetBridgingTlsServerCertMaxChainDepthOk

`func (o *MsgVpn) GetBridgingTlsServerCertMaxChainDepthOk() (*int64, bool)`

GetBridgingTlsServerCertMaxChainDepthOk returns a tuple with the BridgingTlsServerCertMaxChainDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgingTlsServerCertMaxChainDepth

`func (o *MsgVpn) SetBridgingTlsServerCertMaxChainDepth(v int64)`

SetBridgingTlsServerCertMaxChainDepth sets BridgingTlsServerCertMaxChainDepth field to given value.

### HasBridgingTlsServerCertMaxChainDepth

`func (o *MsgVpn) HasBridgingTlsServerCertMaxChainDepth() bool`

HasBridgingTlsServerCertMaxChainDepth returns a boolean if a field has been set.

### GetBridgingTlsServerCertValidateDateEnabled

`func (o *MsgVpn) GetBridgingTlsServerCertValidateDateEnabled() bool`

GetBridgingTlsServerCertValidateDateEnabled returns the BridgingTlsServerCertValidateDateEnabled field if non-nil, zero value otherwise.

### GetBridgingTlsServerCertValidateDateEnabledOk

`func (o *MsgVpn) GetBridgingTlsServerCertValidateDateEnabledOk() (*bool, bool)`

GetBridgingTlsServerCertValidateDateEnabledOk returns a tuple with the BridgingTlsServerCertValidateDateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgingTlsServerCertValidateDateEnabled

`func (o *MsgVpn) SetBridgingTlsServerCertValidateDateEnabled(v bool)`

SetBridgingTlsServerCertValidateDateEnabled sets BridgingTlsServerCertValidateDateEnabled field to given value.

### HasBridgingTlsServerCertValidateDateEnabled

`func (o *MsgVpn) HasBridgingTlsServerCertValidateDateEnabled() bool`

HasBridgingTlsServerCertValidateDateEnabled returns a boolean if a field has been set.

### GetBridgingTlsServerCertValidateNameEnabled

`func (o *MsgVpn) GetBridgingTlsServerCertValidateNameEnabled() bool`

GetBridgingTlsServerCertValidateNameEnabled returns the BridgingTlsServerCertValidateNameEnabled field if non-nil, zero value otherwise.

### GetBridgingTlsServerCertValidateNameEnabledOk

`func (o *MsgVpn) GetBridgingTlsServerCertValidateNameEnabledOk() (*bool, bool)`

GetBridgingTlsServerCertValidateNameEnabledOk returns a tuple with the BridgingTlsServerCertValidateNameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgingTlsServerCertValidateNameEnabled

`func (o *MsgVpn) SetBridgingTlsServerCertValidateNameEnabled(v bool)`

SetBridgingTlsServerCertValidateNameEnabled sets BridgingTlsServerCertValidateNameEnabled field to given value.

### HasBridgingTlsServerCertValidateNameEnabled

`func (o *MsgVpn) HasBridgingTlsServerCertValidateNameEnabled() bool`

HasBridgingTlsServerCertValidateNameEnabled returns a boolean if a field has been set.

### GetDistributedCacheManagementEnabled

`func (o *MsgVpn) GetDistributedCacheManagementEnabled() bool`

GetDistributedCacheManagementEnabled returns the DistributedCacheManagementEnabled field if non-nil, zero value otherwise.

### GetDistributedCacheManagementEnabledOk

`func (o *MsgVpn) GetDistributedCacheManagementEnabledOk() (*bool, bool)`

GetDistributedCacheManagementEnabledOk returns a tuple with the DistributedCacheManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedCacheManagementEnabled

`func (o *MsgVpn) SetDistributedCacheManagementEnabled(v bool)`

SetDistributedCacheManagementEnabled sets DistributedCacheManagementEnabled field to given value.

### HasDistributedCacheManagementEnabled

`func (o *MsgVpn) HasDistributedCacheManagementEnabled() bool`

HasDistributedCacheManagementEnabled returns a boolean if a field has been set.

### GetDmrEnabled

`func (o *MsgVpn) GetDmrEnabled() bool`

GetDmrEnabled returns the DmrEnabled field if non-nil, zero value otherwise.

### GetDmrEnabledOk

`func (o *MsgVpn) GetDmrEnabledOk() (*bool, bool)`

GetDmrEnabledOk returns a tuple with the DmrEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrEnabled

`func (o *MsgVpn) SetDmrEnabled(v bool)`

SetDmrEnabled sets DmrEnabled field to given value.

### HasDmrEnabled

`func (o *MsgVpn) HasDmrEnabled() bool`

HasDmrEnabled returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpn) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpn) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpn) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpn) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventConnectionCountThreshold

`func (o *MsgVpn) GetEventConnectionCountThreshold() EventThreshold`

GetEventConnectionCountThreshold returns the EventConnectionCountThreshold field if non-nil, zero value otherwise.

### GetEventConnectionCountThresholdOk

`func (o *MsgVpn) GetEventConnectionCountThresholdOk() (*EventThreshold, bool)`

GetEventConnectionCountThresholdOk returns a tuple with the EventConnectionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventConnectionCountThreshold

`func (o *MsgVpn) SetEventConnectionCountThreshold(v EventThreshold)`

SetEventConnectionCountThreshold sets EventConnectionCountThreshold field to given value.

### HasEventConnectionCountThreshold

`func (o *MsgVpn) HasEventConnectionCountThreshold() bool`

HasEventConnectionCountThreshold returns a boolean if a field has been set.

### GetEventEgressFlowCountThreshold

`func (o *MsgVpn) GetEventEgressFlowCountThreshold() EventThreshold`

GetEventEgressFlowCountThreshold returns the EventEgressFlowCountThreshold field if non-nil, zero value otherwise.

### GetEventEgressFlowCountThresholdOk

`func (o *MsgVpn) GetEventEgressFlowCountThresholdOk() (*EventThreshold, bool)`

GetEventEgressFlowCountThresholdOk returns a tuple with the EventEgressFlowCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventEgressFlowCountThreshold

`func (o *MsgVpn) SetEventEgressFlowCountThreshold(v EventThreshold)`

SetEventEgressFlowCountThreshold sets EventEgressFlowCountThreshold field to given value.

### HasEventEgressFlowCountThreshold

`func (o *MsgVpn) HasEventEgressFlowCountThreshold() bool`

HasEventEgressFlowCountThreshold returns a boolean if a field has been set.

### GetEventEgressMsgRateThreshold

`func (o *MsgVpn) GetEventEgressMsgRateThreshold() EventThresholdByValue`

GetEventEgressMsgRateThreshold returns the EventEgressMsgRateThreshold field if non-nil, zero value otherwise.

### GetEventEgressMsgRateThresholdOk

`func (o *MsgVpn) GetEventEgressMsgRateThresholdOk() (*EventThresholdByValue, bool)`

GetEventEgressMsgRateThresholdOk returns a tuple with the EventEgressMsgRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventEgressMsgRateThreshold

`func (o *MsgVpn) SetEventEgressMsgRateThreshold(v EventThresholdByValue)`

SetEventEgressMsgRateThreshold sets EventEgressMsgRateThreshold field to given value.

### HasEventEgressMsgRateThreshold

`func (o *MsgVpn) HasEventEgressMsgRateThreshold() bool`

HasEventEgressMsgRateThreshold returns a boolean if a field has been set.

### GetEventEndpointCountThreshold

`func (o *MsgVpn) GetEventEndpointCountThreshold() EventThreshold`

GetEventEndpointCountThreshold returns the EventEndpointCountThreshold field if non-nil, zero value otherwise.

### GetEventEndpointCountThresholdOk

`func (o *MsgVpn) GetEventEndpointCountThresholdOk() (*EventThreshold, bool)`

GetEventEndpointCountThresholdOk returns a tuple with the EventEndpointCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventEndpointCountThreshold

`func (o *MsgVpn) SetEventEndpointCountThreshold(v EventThreshold)`

SetEventEndpointCountThreshold sets EventEndpointCountThreshold field to given value.

### HasEventEndpointCountThreshold

`func (o *MsgVpn) HasEventEndpointCountThreshold() bool`

HasEventEndpointCountThreshold returns a boolean if a field has been set.

### GetEventIngressFlowCountThreshold

`func (o *MsgVpn) GetEventIngressFlowCountThreshold() EventThreshold`

GetEventIngressFlowCountThreshold returns the EventIngressFlowCountThreshold field if non-nil, zero value otherwise.

### GetEventIngressFlowCountThresholdOk

`func (o *MsgVpn) GetEventIngressFlowCountThresholdOk() (*EventThreshold, bool)`

GetEventIngressFlowCountThresholdOk returns a tuple with the EventIngressFlowCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIngressFlowCountThreshold

`func (o *MsgVpn) SetEventIngressFlowCountThreshold(v EventThreshold)`

SetEventIngressFlowCountThreshold sets EventIngressFlowCountThreshold field to given value.

### HasEventIngressFlowCountThreshold

`func (o *MsgVpn) HasEventIngressFlowCountThreshold() bool`

HasEventIngressFlowCountThreshold returns a boolean if a field has been set.

### GetEventIngressMsgRateThreshold

`func (o *MsgVpn) GetEventIngressMsgRateThreshold() EventThresholdByValue`

GetEventIngressMsgRateThreshold returns the EventIngressMsgRateThreshold field if non-nil, zero value otherwise.

### GetEventIngressMsgRateThresholdOk

`func (o *MsgVpn) GetEventIngressMsgRateThresholdOk() (*EventThresholdByValue, bool)`

GetEventIngressMsgRateThresholdOk returns a tuple with the EventIngressMsgRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIngressMsgRateThreshold

`func (o *MsgVpn) SetEventIngressMsgRateThreshold(v EventThresholdByValue)`

SetEventIngressMsgRateThreshold sets EventIngressMsgRateThreshold field to given value.

### HasEventIngressMsgRateThreshold

`func (o *MsgVpn) HasEventIngressMsgRateThreshold() bool`

HasEventIngressMsgRateThreshold returns a boolean if a field has been set.

### GetEventLargeMsgThreshold

`func (o *MsgVpn) GetEventLargeMsgThreshold() int64`

GetEventLargeMsgThreshold returns the EventLargeMsgThreshold field if non-nil, zero value otherwise.

### GetEventLargeMsgThresholdOk

`func (o *MsgVpn) GetEventLargeMsgThresholdOk() (*int64, bool)`

GetEventLargeMsgThresholdOk returns a tuple with the EventLargeMsgThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLargeMsgThreshold

`func (o *MsgVpn) SetEventLargeMsgThreshold(v int64)`

SetEventLargeMsgThreshold sets EventLargeMsgThreshold field to given value.

### HasEventLargeMsgThreshold

`func (o *MsgVpn) HasEventLargeMsgThreshold() bool`

HasEventLargeMsgThreshold returns a boolean if a field has been set.

### GetEventLogTag

`func (o *MsgVpn) GetEventLogTag() string`

GetEventLogTag returns the EventLogTag field if non-nil, zero value otherwise.

### GetEventLogTagOk

`func (o *MsgVpn) GetEventLogTagOk() (*string, bool)`

GetEventLogTagOk returns a tuple with the EventLogTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLogTag

`func (o *MsgVpn) SetEventLogTag(v string)`

SetEventLogTag sets EventLogTag field to given value.

### HasEventLogTag

`func (o *MsgVpn) HasEventLogTag() bool`

HasEventLogTag returns a boolean if a field has been set.

### GetEventMsgSpoolUsageThreshold

`func (o *MsgVpn) GetEventMsgSpoolUsageThreshold() EventThreshold`

GetEventMsgSpoolUsageThreshold returns the EventMsgSpoolUsageThreshold field if non-nil, zero value otherwise.

### GetEventMsgSpoolUsageThresholdOk

`func (o *MsgVpn) GetEventMsgSpoolUsageThresholdOk() (*EventThreshold, bool)`

GetEventMsgSpoolUsageThresholdOk returns a tuple with the EventMsgSpoolUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMsgSpoolUsageThreshold

`func (o *MsgVpn) SetEventMsgSpoolUsageThreshold(v EventThreshold)`

SetEventMsgSpoolUsageThreshold sets EventMsgSpoolUsageThreshold field to given value.

### HasEventMsgSpoolUsageThreshold

`func (o *MsgVpn) HasEventMsgSpoolUsageThreshold() bool`

HasEventMsgSpoolUsageThreshold returns a boolean if a field has been set.

### GetEventPublishClientEnabled

`func (o *MsgVpn) GetEventPublishClientEnabled() bool`

GetEventPublishClientEnabled returns the EventPublishClientEnabled field if non-nil, zero value otherwise.

### GetEventPublishClientEnabledOk

`func (o *MsgVpn) GetEventPublishClientEnabledOk() (*bool, bool)`

GetEventPublishClientEnabledOk returns a tuple with the EventPublishClientEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPublishClientEnabled

`func (o *MsgVpn) SetEventPublishClientEnabled(v bool)`

SetEventPublishClientEnabled sets EventPublishClientEnabled field to given value.

### HasEventPublishClientEnabled

`func (o *MsgVpn) HasEventPublishClientEnabled() bool`

HasEventPublishClientEnabled returns a boolean if a field has been set.

### GetEventPublishMsgVpnEnabled

`func (o *MsgVpn) GetEventPublishMsgVpnEnabled() bool`

GetEventPublishMsgVpnEnabled returns the EventPublishMsgVpnEnabled field if non-nil, zero value otherwise.

### GetEventPublishMsgVpnEnabledOk

`func (o *MsgVpn) GetEventPublishMsgVpnEnabledOk() (*bool, bool)`

GetEventPublishMsgVpnEnabledOk returns a tuple with the EventPublishMsgVpnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPublishMsgVpnEnabled

`func (o *MsgVpn) SetEventPublishMsgVpnEnabled(v bool)`

SetEventPublishMsgVpnEnabled sets EventPublishMsgVpnEnabled field to given value.

### HasEventPublishMsgVpnEnabled

`func (o *MsgVpn) HasEventPublishMsgVpnEnabled() bool`

HasEventPublishMsgVpnEnabled returns a boolean if a field has been set.

### GetEventPublishSubscriptionMode

`func (o *MsgVpn) GetEventPublishSubscriptionMode() string`

GetEventPublishSubscriptionMode returns the EventPublishSubscriptionMode field if non-nil, zero value otherwise.

### GetEventPublishSubscriptionModeOk

`func (o *MsgVpn) GetEventPublishSubscriptionModeOk() (*string, bool)`

GetEventPublishSubscriptionModeOk returns a tuple with the EventPublishSubscriptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPublishSubscriptionMode

`func (o *MsgVpn) SetEventPublishSubscriptionMode(v string)`

SetEventPublishSubscriptionMode sets EventPublishSubscriptionMode field to given value.

### HasEventPublishSubscriptionMode

`func (o *MsgVpn) HasEventPublishSubscriptionMode() bool`

HasEventPublishSubscriptionMode returns a boolean if a field has been set.

### GetEventPublishTopicFormatMqttEnabled

`func (o *MsgVpn) GetEventPublishTopicFormatMqttEnabled() bool`

GetEventPublishTopicFormatMqttEnabled returns the EventPublishTopicFormatMqttEnabled field if non-nil, zero value otherwise.

### GetEventPublishTopicFormatMqttEnabledOk

`func (o *MsgVpn) GetEventPublishTopicFormatMqttEnabledOk() (*bool, bool)`

GetEventPublishTopicFormatMqttEnabledOk returns a tuple with the EventPublishTopicFormatMqttEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPublishTopicFormatMqttEnabled

`func (o *MsgVpn) SetEventPublishTopicFormatMqttEnabled(v bool)`

SetEventPublishTopicFormatMqttEnabled sets EventPublishTopicFormatMqttEnabled field to given value.

### HasEventPublishTopicFormatMqttEnabled

`func (o *MsgVpn) HasEventPublishTopicFormatMqttEnabled() bool`

HasEventPublishTopicFormatMqttEnabled returns a boolean if a field has been set.

### GetEventPublishTopicFormatSmfEnabled

`func (o *MsgVpn) GetEventPublishTopicFormatSmfEnabled() bool`

GetEventPublishTopicFormatSmfEnabled returns the EventPublishTopicFormatSmfEnabled field if non-nil, zero value otherwise.

### GetEventPublishTopicFormatSmfEnabledOk

`func (o *MsgVpn) GetEventPublishTopicFormatSmfEnabledOk() (*bool, bool)`

GetEventPublishTopicFormatSmfEnabledOk returns a tuple with the EventPublishTopicFormatSmfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPublishTopicFormatSmfEnabled

`func (o *MsgVpn) SetEventPublishTopicFormatSmfEnabled(v bool)`

SetEventPublishTopicFormatSmfEnabled sets EventPublishTopicFormatSmfEnabled field to given value.

### HasEventPublishTopicFormatSmfEnabled

`func (o *MsgVpn) HasEventPublishTopicFormatSmfEnabled() bool`

HasEventPublishTopicFormatSmfEnabled returns a boolean if a field has been set.

### GetEventServiceAmqpConnectionCountThreshold

`func (o *MsgVpn) GetEventServiceAmqpConnectionCountThreshold() EventThreshold`

GetEventServiceAmqpConnectionCountThreshold returns the EventServiceAmqpConnectionCountThreshold field if non-nil, zero value otherwise.

### GetEventServiceAmqpConnectionCountThresholdOk

`func (o *MsgVpn) GetEventServiceAmqpConnectionCountThresholdOk() (*EventThreshold, bool)`

GetEventServiceAmqpConnectionCountThresholdOk returns a tuple with the EventServiceAmqpConnectionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceAmqpConnectionCountThreshold

`func (o *MsgVpn) SetEventServiceAmqpConnectionCountThreshold(v EventThreshold)`

SetEventServiceAmqpConnectionCountThreshold sets EventServiceAmqpConnectionCountThreshold field to given value.

### HasEventServiceAmqpConnectionCountThreshold

`func (o *MsgVpn) HasEventServiceAmqpConnectionCountThreshold() bool`

HasEventServiceAmqpConnectionCountThreshold returns a boolean if a field has been set.

### GetEventServiceMqttConnectionCountThreshold

`func (o *MsgVpn) GetEventServiceMqttConnectionCountThreshold() EventThreshold`

GetEventServiceMqttConnectionCountThreshold returns the EventServiceMqttConnectionCountThreshold field if non-nil, zero value otherwise.

### GetEventServiceMqttConnectionCountThresholdOk

`func (o *MsgVpn) GetEventServiceMqttConnectionCountThresholdOk() (*EventThreshold, bool)`

GetEventServiceMqttConnectionCountThresholdOk returns a tuple with the EventServiceMqttConnectionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceMqttConnectionCountThreshold

`func (o *MsgVpn) SetEventServiceMqttConnectionCountThreshold(v EventThreshold)`

SetEventServiceMqttConnectionCountThreshold sets EventServiceMqttConnectionCountThreshold field to given value.

### HasEventServiceMqttConnectionCountThreshold

`func (o *MsgVpn) HasEventServiceMqttConnectionCountThreshold() bool`

HasEventServiceMqttConnectionCountThreshold returns a boolean if a field has been set.

### GetEventServiceRestIncomingConnectionCountThreshold

`func (o *MsgVpn) GetEventServiceRestIncomingConnectionCountThreshold() EventThreshold`

GetEventServiceRestIncomingConnectionCountThreshold returns the EventServiceRestIncomingConnectionCountThreshold field if non-nil, zero value otherwise.

### GetEventServiceRestIncomingConnectionCountThresholdOk

`func (o *MsgVpn) GetEventServiceRestIncomingConnectionCountThresholdOk() (*EventThreshold, bool)`

GetEventServiceRestIncomingConnectionCountThresholdOk returns a tuple with the EventServiceRestIncomingConnectionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceRestIncomingConnectionCountThreshold

`func (o *MsgVpn) SetEventServiceRestIncomingConnectionCountThreshold(v EventThreshold)`

SetEventServiceRestIncomingConnectionCountThreshold sets EventServiceRestIncomingConnectionCountThreshold field to given value.

### HasEventServiceRestIncomingConnectionCountThreshold

`func (o *MsgVpn) HasEventServiceRestIncomingConnectionCountThreshold() bool`

HasEventServiceRestIncomingConnectionCountThreshold returns a boolean if a field has been set.

### GetEventServiceSmfConnectionCountThreshold

`func (o *MsgVpn) GetEventServiceSmfConnectionCountThreshold() EventThreshold`

GetEventServiceSmfConnectionCountThreshold returns the EventServiceSmfConnectionCountThreshold field if non-nil, zero value otherwise.

### GetEventServiceSmfConnectionCountThresholdOk

`func (o *MsgVpn) GetEventServiceSmfConnectionCountThresholdOk() (*EventThreshold, bool)`

GetEventServiceSmfConnectionCountThresholdOk returns a tuple with the EventServiceSmfConnectionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceSmfConnectionCountThreshold

`func (o *MsgVpn) SetEventServiceSmfConnectionCountThreshold(v EventThreshold)`

SetEventServiceSmfConnectionCountThreshold sets EventServiceSmfConnectionCountThreshold field to given value.

### HasEventServiceSmfConnectionCountThreshold

`func (o *MsgVpn) HasEventServiceSmfConnectionCountThreshold() bool`

HasEventServiceSmfConnectionCountThreshold returns a boolean if a field has been set.

### GetEventServiceWebConnectionCountThreshold

`func (o *MsgVpn) GetEventServiceWebConnectionCountThreshold() EventThreshold`

GetEventServiceWebConnectionCountThreshold returns the EventServiceWebConnectionCountThreshold field if non-nil, zero value otherwise.

### GetEventServiceWebConnectionCountThresholdOk

`func (o *MsgVpn) GetEventServiceWebConnectionCountThresholdOk() (*EventThreshold, bool)`

GetEventServiceWebConnectionCountThresholdOk returns a tuple with the EventServiceWebConnectionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventServiceWebConnectionCountThreshold

`func (o *MsgVpn) SetEventServiceWebConnectionCountThreshold(v EventThreshold)`

SetEventServiceWebConnectionCountThreshold sets EventServiceWebConnectionCountThreshold field to given value.

### HasEventServiceWebConnectionCountThreshold

`func (o *MsgVpn) HasEventServiceWebConnectionCountThreshold() bool`

HasEventServiceWebConnectionCountThreshold returns a boolean if a field has been set.

### GetEventSubscriptionCountThreshold

`func (o *MsgVpn) GetEventSubscriptionCountThreshold() EventThreshold`

GetEventSubscriptionCountThreshold returns the EventSubscriptionCountThreshold field if non-nil, zero value otherwise.

### GetEventSubscriptionCountThresholdOk

`func (o *MsgVpn) GetEventSubscriptionCountThresholdOk() (*EventThreshold, bool)`

GetEventSubscriptionCountThresholdOk returns a tuple with the EventSubscriptionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubscriptionCountThreshold

`func (o *MsgVpn) SetEventSubscriptionCountThreshold(v EventThreshold)`

SetEventSubscriptionCountThreshold sets EventSubscriptionCountThreshold field to given value.

### HasEventSubscriptionCountThreshold

`func (o *MsgVpn) HasEventSubscriptionCountThreshold() bool`

HasEventSubscriptionCountThreshold returns a boolean if a field has been set.

### GetEventTransactedSessionCountThreshold

`func (o *MsgVpn) GetEventTransactedSessionCountThreshold() EventThreshold`

GetEventTransactedSessionCountThreshold returns the EventTransactedSessionCountThreshold field if non-nil, zero value otherwise.

### GetEventTransactedSessionCountThresholdOk

`func (o *MsgVpn) GetEventTransactedSessionCountThresholdOk() (*EventThreshold, bool)`

GetEventTransactedSessionCountThresholdOk returns a tuple with the EventTransactedSessionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTransactedSessionCountThreshold

`func (o *MsgVpn) SetEventTransactedSessionCountThreshold(v EventThreshold)`

SetEventTransactedSessionCountThreshold sets EventTransactedSessionCountThreshold field to given value.

### HasEventTransactedSessionCountThreshold

`func (o *MsgVpn) HasEventTransactedSessionCountThreshold() bool`

HasEventTransactedSessionCountThreshold returns a boolean if a field has been set.

### GetEventTransactionCountThreshold

`func (o *MsgVpn) GetEventTransactionCountThreshold() EventThreshold`

GetEventTransactionCountThreshold returns the EventTransactionCountThreshold field if non-nil, zero value otherwise.

### GetEventTransactionCountThresholdOk

`func (o *MsgVpn) GetEventTransactionCountThresholdOk() (*EventThreshold, bool)`

GetEventTransactionCountThresholdOk returns a tuple with the EventTransactionCountThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTransactionCountThreshold

`func (o *MsgVpn) SetEventTransactionCountThreshold(v EventThreshold)`

SetEventTransactionCountThreshold sets EventTransactionCountThreshold field to given value.

### HasEventTransactionCountThreshold

`func (o *MsgVpn) HasEventTransactionCountThreshold() bool`

HasEventTransactionCountThreshold returns a boolean if a field has been set.

### GetExportSubscriptionsEnabled

`func (o *MsgVpn) GetExportSubscriptionsEnabled() bool`

GetExportSubscriptionsEnabled returns the ExportSubscriptionsEnabled field if non-nil, zero value otherwise.

### GetExportSubscriptionsEnabledOk

`func (o *MsgVpn) GetExportSubscriptionsEnabledOk() (*bool, bool)`

GetExportSubscriptionsEnabledOk returns a tuple with the ExportSubscriptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportSubscriptionsEnabled

`func (o *MsgVpn) SetExportSubscriptionsEnabled(v bool)`

SetExportSubscriptionsEnabled sets ExportSubscriptionsEnabled field to given value.

### HasExportSubscriptionsEnabled

`func (o *MsgVpn) HasExportSubscriptionsEnabled() bool`

HasExportSubscriptionsEnabled returns a boolean if a field has been set.

### GetJndiEnabled

`func (o *MsgVpn) GetJndiEnabled() bool`

GetJndiEnabled returns the JndiEnabled field if non-nil, zero value otherwise.

### GetJndiEnabledOk

`func (o *MsgVpn) GetJndiEnabledOk() (*bool, bool)`

GetJndiEnabledOk returns a tuple with the JndiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJndiEnabled

`func (o *MsgVpn) SetJndiEnabled(v bool)`

SetJndiEnabled sets JndiEnabled field to given value.

### HasJndiEnabled

`func (o *MsgVpn) HasJndiEnabled() bool`

HasJndiEnabled returns a boolean if a field has been set.

### GetMaxConnectionCount

`func (o *MsgVpn) GetMaxConnectionCount() int64`

GetMaxConnectionCount returns the MaxConnectionCount field if non-nil, zero value otherwise.

### GetMaxConnectionCountOk

`func (o *MsgVpn) GetMaxConnectionCountOk() (*int64, bool)`

GetMaxConnectionCountOk returns a tuple with the MaxConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnectionCount

`func (o *MsgVpn) SetMaxConnectionCount(v int64)`

SetMaxConnectionCount sets MaxConnectionCount field to given value.

### HasMaxConnectionCount

`func (o *MsgVpn) HasMaxConnectionCount() bool`

HasMaxConnectionCount returns a boolean if a field has been set.

### GetMaxEgressFlowCount

`func (o *MsgVpn) GetMaxEgressFlowCount() int64`

GetMaxEgressFlowCount returns the MaxEgressFlowCount field if non-nil, zero value otherwise.

### GetMaxEgressFlowCountOk

`func (o *MsgVpn) GetMaxEgressFlowCountOk() (*int64, bool)`

GetMaxEgressFlowCountOk returns a tuple with the MaxEgressFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEgressFlowCount

`func (o *MsgVpn) SetMaxEgressFlowCount(v int64)`

SetMaxEgressFlowCount sets MaxEgressFlowCount field to given value.

### HasMaxEgressFlowCount

`func (o *MsgVpn) HasMaxEgressFlowCount() bool`

HasMaxEgressFlowCount returns a boolean if a field has been set.

### GetMaxEndpointCount

`func (o *MsgVpn) GetMaxEndpointCount() int64`

GetMaxEndpointCount returns the MaxEndpointCount field if non-nil, zero value otherwise.

### GetMaxEndpointCountOk

`func (o *MsgVpn) GetMaxEndpointCountOk() (*int64, bool)`

GetMaxEndpointCountOk returns a tuple with the MaxEndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndpointCount

`func (o *MsgVpn) SetMaxEndpointCount(v int64)`

SetMaxEndpointCount sets MaxEndpointCount field to given value.

### HasMaxEndpointCount

`func (o *MsgVpn) HasMaxEndpointCount() bool`

HasMaxEndpointCount returns a boolean if a field has been set.

### GetMaxIngressFlowCount

`func (o *MsgVpn) GetMaxIngressFlowCount() int64`

GetMaxIngressFlowCount returns the MaxIngressFlowCount field if non-nil, zero value otherwise.

### GetMaxIngressFlowCountOk

`func (o *MsgVpn) GetMaxIngressFlowCountOk() (*int64, bool)`

GetMaxIngressFlowCountOk returns a tuple with the MaxIngressFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIngressFlowCount

`func (o *MsgVpn) SetMaxIngressFlowCount(v int64)`

SetMaxIngressFlowCount sets MaxIngressFlowCount field to given value.

### HasMaxIngressFlowCount

`func (o *MsgVpn) HasMaxIngressFlowCount() bool`

HasMaxIngressFlowCount returns a boolean if a field has been set.

### GetMaxMsgSpoolUsage

`func (o *MsgVpn) GetMaxMsgSpoolUsage() int64`

GetMaxMsgSpoolUsage returns the MaxMsgSpoolUsage field if non-nil, zero value otherwise.

### GetMaxMsgSpoolUsageOk

`func (o *MsgVpn) GetMaxMsgSpoolUsageOk() (*int64, bool)`

GetMaxMsgSpoolUsageOk returns a tuple with the MaxMsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSpoolUsage

`func (o *MsgVpn) SetMaxMsgSpoolUsage(v int64)`

SetMaxMsgSpoolUsage sets MaxMsgSpoolUsage field to given value.

### HasMaxMsgSpoolUsage

`func (o *MsgVpn) HasMaxMsgSpoolUsage() bool`

HasMaxMsgSpoolUsage returns a boolean if a field has been set.

### GetMaxSubscriptionCount

`func (o *MsgVpn) GetMaxSubscriptionCount() int64`

GetMaxSubscriptionCount returns the MaxSubscriptionCount field if non-nil, zero value otherwise.

### GetMaxSubscriptionCountOk

`func (o *MsgVpn) GetMaxSubscriptionCountOk() (*int64, bool)`

GetMaxSubscriptionCountOk returns a tuple with the MaxSubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSubscriptionCount

`func (o *MsgVpn) SetMaxSubscriptionCount(v int64)`

SetMaxSubscriptionCount sets MaxSubscriptionCount field to given value.

### HasMaxSubscriptionCount

`func (o *MsgVpn) HasMaxSubscriptionCount() bool`

HasMaxSubscriptionCount returns a boolean if a field has been set.

### GetMaxTransactedSessionCount

`func (o *MsgVpn) GetMaxTransactedSessionCount() int64`

GetMaxTransactedSessionCount returns the MaxTransactedSessionCount field if non-nil, zero value otherwise.

### GetMaxTransactedSessionCountOk

`func (o *MsgVpn) GetMaxTransactedSessionCountOk() (*int64, bool)`

GetMaxTransactedSessionCountOk returns a tuple with the MaxTransactedSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactedSessionCount

`func (o *MsgVpn) SetMaxTransactedSessionCount(v int64)`

SetMaxTransactedSessionCount sets MaxTransactedSessionCount field to given value.

### HasMaxTransactedSessionCount

`func (o *MsgVpn) HasMaxTransactedSessionCount() bool`

HasMaxTransactedSessionCount returns a boolean if a field has been set.

### GetMaxTransactionCount

`func (o *MsgVpn) GetMaxTransactionCount() int64`

GetMaxTransactionCount returns the MaxTransactionCount field if non-nil, zero value otherwise.

### GetMaxTransactionCountOk

`func (o *MsgVpn) GetMaxTransactionCountOk() (*int64, bool)`

GetMaxTransactionCountOk returns a tuple with the MaxTransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactionCount

`func (o *MsgVpn) SetMaxTransactionCount(v int64)`

SetMaxTransactionCount sets MaxTransactionCount field to given value.

### HasMaxTransactionCount

`func (o *MsgVpn) HasMaxTransactionCount() bool`

HasMaxTransactionCount returns a boolean if a field has been set.

### GetMqttRetainMaxMemory

`func (o *MsgVpn) GetMqttRetainMaxMemory() int32`

GetMqttRetainMaxMemory returns the MqttRetainMaxMemory field if non-nil, zero value otherwise.

### GetMqttRetainMaxMemoryOk

`func (o *MsgVpn) GetMqttRetainMaxMemoryOk() (*int32, bool)`

GetMqttRetainMaxMemoryOk returns a tuple with the MqttRetainMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttRetainMaxMemory

`func (o *MsgVpn) SetMqttRetainMaxMemory(v int32)`

SetMqttRetainMaxMemory sets MqttRetainMaxMemory field to given value.

### HasMqttRetainMaxMemory

`func (o *MsgVpn) HasMqttRetainMaxMemory() bool`

HasMqttRetainMaxMemory returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpn) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpn) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpn) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpn) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetReplicationAckPropagationIntervalMsgCount

`func (o *MsgVpn) GetReplicationAckPropagationIntervalMsgCount() int64`

GetReplicationAckPropagationIntervalMsgCount returns the ReplicationAckPropagationIntervalMsgCount field if non-nil, zero value otherwise.

### GetReplicationAckPropagationIntervalMsgCountOk

`func (o *MsgVpn) GetReplicationAckPropagationIntervalMsgCountOk() (*int64, bool)`

GetReplicationAckPropagationIntervalMsgCountOk returns a tuple with the ReplicationAckPropagationIntervalMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationAckPropagationIntervalMsgCount

`func (o *MsgVpn) SetReplicationAckPropagationIntervalMsgCount(v int64)`

SetReplicationAckPropagationIntervalMsgCount sets ReplicationAckPropagationIntervalMsgCount field to given value.

### HasReplicationAckPropagationIntervalMsgCount

`func (o *MsgVpn) HasReplicationAckPropagationIntervalMsgCount() bool`

HasReplicationAckPropagationIntervalMsgCount returns a boolean if a field has been set.

### GetReplicationBridgeAuthenticationBasicClientUsername

`func (o *MsgVpn) GetReplicationBridgeAuthenticationBasicClientUsername() string`

GetReplicationBridgeAuthenticationBasicClientUsername returns the ReplicationBridgeAuthenticationBasicClientUsername field if non-nil, zero value otherwise.

### GetReplicationBridgeAuthenticationBasicClientUsernameOk

`func (o *MsgVpn) GetReplicationBridgeAuthenticationBasicClientUsernameOk() (*string, bool)`

GetReplicationBridgeAuthenticationBasicClientUsernameOk returns a tuple with the ReplicationBridgeAuthenticationBasicClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeAuthenticationBasicClientUsername

`func (o *MsgVpn) SetReplicationBridgeAuthenticationBasicClientUsername(v string)`

SetReplicationBridgeAuthenticationBasicClientUsername sets ReplicationBridgeAuthenticationBasicClientUsername field to given value.

### HasReplicationBridgeAuthenticationBasicClientUsername

`func (o *MsgVpn) HasReplicationBridgeAuthenticationBasicClientUsername() bool`

HasReplicationBridgeAuthenticationBasicClientUsername returns a boolean if a field has been set.

### GetReplicationBridgeAuthenticationBasicPassword

`func (o *MsgVpn) GetReplicationBridgeAuthenticationBasicPassword() string`

GetReplicationBridgeAuthenticationBasicPassword returns the ReplicationBridgeAuthenticationBasicPassword field if non-nil, zero value otherwise.

### GetReplicationBridgeAuthenticationBasicPasswordOk

`func (o *MsgVpn) GetReplicationBridgeAuthenticationBasicPasswordOk() (*string, bool)`

GetReplicationBridgeAuthenticationBasicPasswordOk returns a tuple with the ReplicationBridgeAuthenticationBasicPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeAuthenticationBasicPassword

`func (o *MsgVpn) SetReplicationBridgeAuthenticationBasicPassword(v string)`

SetReplicationBridgeAuthenticationBasicPassword sets ReplicationBridgeAuthenticationBasicPassword field to given value.

### HasReplicationBridgeAuthenticationBasicPassword

`func (o *MsgVpn) HasReplicationBridgeAuthenticationBasicPassword() bool`

HasReplicationBridgeAuthenticationBasicPassword returns a boolean if a field has been set.

### GetReplicationBridgeAuthenticationClientCertContent

`func (o *MsgVpn) GetReplicationBridgeAuthenticationClientCertContent() string`

GetReplicationBridgeAuthenticationClientCertContent returns the ReplicationBridgeAuthenticationClientCertContent field if non-nil, zero value otherwise.

### GetReplicationBridgeAuthenticationClientCertContentOk

`func (o *MsgVpn) GetReplicationBridgeAuthenticationClientCertContentOk() (*string, bool)`

GetReplicationBridgeAuthenticationClientCertContentOk returns a tuple with the ReplicationBridgeAuthenticationClientCertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeAuthenticationClientCertContent

`func (o *MsgVpn) SetReplicationBridgeAuthenticationClientCertContent(v string)`

SetReplicationBridgeAuthenticationClientCertContent sets ReplicationBridgeAuthenticationClientCertContent field to given value.

### HasReplicationBridgeAuthenticationClientCertContent

`func (o *MsgVpn) HasReplicationBridgeAuthenticationClientCertContent() bool`

HasReplicationBridgeAuthenticationClientCertContent returns a boolean if a field has been set.

### GetReplicationBridgeAuthenticationClientCertPassword

`func (o *MsgVpn) GetReplicationBridgeAuthenticationClientCertPassword() string`

GetReplicationBridgeAuthenticationClientCertPassword returns the ReplicationBridgeAuthenticationClientCertPassword field if non-nil, zero value otherwise.

### GetReplicationBridgeAuthenticationClientCertPasswordOk

`func (o *MsgVpn) GetReplicationBridgeAuthenticationClientCertPasswordOk() (*string, bool)`

GetReplicationBridgeAuthenticationClientCertPasswordOk returns a tuple with the ReplicationBridgeAuthenticationClientCertPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeAuthenticationClientCertPassword

`func (o *MsgVpn) SetReplicationBridgeAuthenticationClientCertPassword(v string)`

SetReplicationBridgeAuthenticationClientCertPassword sets ReplicationBridgeAuthenticationClientCertPassword field to given value.

### HasReplicationBridgeAuthenticationClientCertPassword

`func (o *MsgVpn) HasReplicationBridgeAuthenticationClientCertPassword() bool`

HasReplicationBridgeAuthenticationClientCertPassword returns a boolean if a field has been set.

### GetReplicationBridgeAuthenticationScheme

`func (o *MsgVpn) GetReplicationBridgeAuthenticationScheme() string`

GetReplicationBridgeAuthenticationScheme returns the ReplicationBridgeAuthenticationScheme field if non-nil, zero value otherwise.

### GetReplicationBridgeAuthenticationSchemeOk

`func (o *MsgVpn) GetReplicationBridgeAuthenticationSchemeOk() (*string, bool)`

GetReplicationBridgeAuthenticationSchemeOk returns a tuple with the ReplicationBridgeAuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeAuthenticationScheme

`func (o *MsgVpn) SetReplicationBridgeAuthenticationScheme(v string)`

SetReplicationBridgeAuthenticationScheme sets ReplicationBridgeAuthenticationScheme field to given value.

### HasReplicationBridgeAuthenticationScheme

`func (o *MsgVpn) HasReplicationBridgeAuthenticationScheme() bool`

HasReplicationBridgeAuthenticationScheme returns a boolean if a field has been set.

### GetReplicationBridgeCompressedDataEnabled

`func (o *MsgVpn) GetReplicationBridgeCompressedDataEnabled() bool`

GetReplicationBridgeCompressedDataEnabled returns the ReplicationBridgeCompressedDataEnabled field if non-nil, zero value otherwise.

### GetReplicationBridgeCompressedDataEnabledOk

`func (o *MsgVpn) GetReplicationBridgeCompressedDataEnabledOk() (*bool, bool)`

GetReplicationBridgeCompressedDataEnabledOk returns a tuple with the ReplicationBridgeCompressedDataEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeCompressedDataEnabled

`func (o *MsgVpn) SetReplicationBridgeCompressedDataEnabled(v bool)`

SetReplicationBridgeCompressedDataEnabled sets ReplicationBridgeCompressedDataEnabled field to given value.

### HasReplicationBridgeCompressedDataEnabled

`func (o *MsgVpn) HasReplicationBridgeCompressedDataEnabled() bool`

HasReplicationBridgeCompressedDataEnabled returns a boolean if a field has been set.

### GetReplicationBridgeEgressFlowWindowSize

`func (o *MsgVpn) GetReplicationBridgeEgressFlowWindowSize() int64`

GetReplicationBridgeEgressFlowWindowSize returns the ReplicationBridgeEgressFlowWindowSize field if non-nil, zero value otherwise.

### GetReplicationBridgeEgressFlowWindowSizeOk

`func (o *MsgVpn) GetReplicationBridgeEgressFlowWindowSizeOk() (*int64, bool)`

GetReplicationBridgeEgressFlowWindowSizeOk returns a tuple with the ReplicationBridgeEgressFlowWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeEgressFlowWindowSize

`func (o *MsgVpn) SetReplicationBridgeEgressFlowWindowSize(v int64)`

SetReplicationBridgeEgressFlowWindowSize sets ReplicationBridgeEgressFlowWindowSize field to given value.

### HasReplicationBridgeEgressFlowWindowSize

`func (o *MsgVpn) HasReplicationBridgeEgressFlowWindowSize() bool`

HasReplicationBridgeEgressFlowWindowSize returns a boolean if a field has been set.

### GetReplicationBridgeRetryDelay

`func (o *MsgVpn) GetReplicationBridgeRetryDelay() int64`

GetReplicationBridgeRetryDelay returns the ReplicationBridgeRetryDelay field if non-nil, zero value otherwise.

### GetReplicationBridgeRetryDelayOk

`func (o *MsgVpn) GetReplicationBridgeRetryDelayOk() (*int64, bool)`

GetReplicationBridgeRetryDelayOk returns a tuple with the ReplicationBridgeRetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeRetryDelay

`func (o *MsgVpn) SetReplicationBridgeRetryDelay(v int64)`

SetReplicationBridgeRetryDelay sets ReplicationBridgeRetryDelay field to given value.

### HasReplicationBridgeRetryDelay

`func (o *MsgVpn) HasReplicationBridgeRetryDelay() bool`

HasReplicationBridgeRetryDelay returns a boolean if a field has been set.

### GetReplicationBridgeTlsEnabled

`func (o *MsgVpn) GetReplicationBridgeTlsEnabled() bool`

GetReplicationBridgeTlsEnabled returns the ReplicationBridgeTlsEnabled field if non-nil, zero value otherwise.

### GetReplicationBridgeTlsEnabledOk

`func (o *MsgVpn) GetReplicationBridgeTlsEnabledOk() (*bool, bool)`

GetReplicationBridgeTlsEnabledOk returns a tuple with the ReplicationBridgeTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeTlsEnabled

`func (o *MsgVpn) SetReplicationBridgeTlsEnabled(v bool)`

SetReplicationBridgeTlsEnabled sets ReplicationBridgeTlsEnabled field to given value.

### HasReplicationBridgeTlsEnabled

`func (o *MsgVpn) HasReplicationBridgeTlsEnabled() bool`

HasReplicationBridgeTlsEnabled returns a boolean if a field has been set.

### GetReplicationBridgeUnidirectionalClientProfileName

`func (o *MsgVpn) GetReplicationBridgeUnidirectionalClientProfileName() string`

GetReplicationBridgeUnidirectionalClientProfileName returns the ReplicationBridgeUnidirectionalClientProfileName field if non-nil, zero value otherwise.

### GetReplicationBridgeUnidirectionalClientProfileNameOk

`func (o *MsgVpn) GetReplicationBridgeUnidirectionalClientProfileNameOk() (*string, bool)`

GetReplicationBridgeUnidirectionalClientProfileNameOk returns a tuple with the ReplicationBridgeUnidirectionalClientProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeUnidirectionalClientProfileName

`func (o *MsgVpn) SetReplicationBridgeUnidirectionalClientProfileName(v string)`

SetReplicationBridgeUnidirectionalClientProfileName sets ReplicationBridgeUnidirectionalClientProfileName field to given value.

### HasReplicationBridgeUnidirectionalClientProfileName

`func (o *MsgVpn) HasReplicationBridgeUnidirectionalClientProfileName() bool`

HasReplicationBridgeUnidirectionalClientProfileName returns a boolean if a field has been set.

### GetReplicationEnabled

`func (o *MsgVpn) GetReplicationEnabled() bool`

GetReplicationEnabled returns the ReplicationEnabled field if non-nil, zero value otherwise.

### GetReplicationEnabledOk

`func (o *MsgVpn) GetReplicationEnabledOk() (*bool, bool)`

GetReplicationEnabledOk returns a tuple with the ReplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationEnabled

`func (o *MsgVpn) SetReplicationEnabled(v bool)`

SetReplicationEnabled sets ReplicationEnabled field to given value.

### HasReplicationEnabled

`func (o *MsgVpn) HasReplicationEnabled() bool`

HasReplicationEnabled returns a boolean if a field has been set.

### GetReplicationEnabledQueueBehavior

`func (o *MsgVpn) GetReplicationEnabledQueueBehavior() string`

GetReplicationEnabledQueueBehavior returns the ReplicationEnabledQueueBehavior field if non-nil, zero value otherwise.

### GetReplicationEnabledQueueBehaviorOk

`func (o *MsgVpn) GetReplicationEnabledQueueBehaviorOk() (*string, bool)`

GetReplicationEnabledQueueBehaviorOk returns a tuple with the ReplicationEnabledQueueBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationEnabledQueueBehavior

`func (o *MsgVpn) SetReplicationEnabledQueueBehavior(v string)`

SetReplicationEnabledQueueBehavior sets ReplicationEnabledQueueBehavior field to given value.

### HasReplicationEnabledQueueBehavior

`func (o *MsgVpn) HasReplicationEnabledQueueBehavior() bool`

HasReplicationEnabledQueueBehavior returns a boolean if a field has been set.

### GetReplicationQueueMaxMsgSpoolUsage

`func (o *MsgVpn) GetReplicationQueueMaxMsgSpoolUsage() int64`

GetReplicationQueueMaxMsgSpoolUsage returns the ReplicationQueueMaxMsgSpoolUsage field if non-nil, zero value otherwise.

### GetReplicationQueueMaxMsgSpoolUsageOk

`func (o *MsgVpn) GetReplicationQueueMaxMsgSpoolUsageOk() (*int64, bool)`

GetReplicationQueueMaxMsgSpoolUsageOk returns a tuple with the ReplicationQueueMaxMsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationQueueMaxMsgSpoolUsage

`func (o *MsgVpn) SetReplicationQueueMaxMsgSpoolUsage(v int64)`

SetReplicationQueueMaxMsgSpoolUsage sets ReplicationQueueMaxMsgSpoolUsage field to given value.

### HasReplicationQueueMaxMsgSpoolUsage

`func (o *MsgVpn) HasReplicationQueueMaxMsgSpoolUsage() bool`

HasReplicationQueueMaxMsgSpoolUsage returns a boolean if a field has been set.

### GetReplicationQueueRejectMsgToSenderOnDiscardEnabled

`func (o *MsgVpn) GetReplicationQueueRejectMsgToSenderOnDiscardEnabled() bool`

GetReplicationQueueRejectMsgToSenderOnDiscardEnabled returns the ReplicationQueueRejectMsgToSenderOnDiscardEnabled field if non-nil, zero value otherwise.

### GetReplicationQueueRejectMsgToSenderOnDiscardEnabledOk

`func (o *MsgVpn) GetReplicationQueueRejectMsgToSenderOnDiscardEnabledOk() (*bool, bool)`

GetReplicationQueueRejectMsgToSenderOnDiscardEnabledOk returns a tuple with the ReplicationQueueRejectMsgToSenderOnDiscardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationQueueRejectMsgToSenderOnDiscardEnabled

`func (o *MsgVpn) SetReplicationQueueRejectMsgToSenderOnDiscardEnabled(v bool)`

SetReplicationQueueRejectMsgToSenderOnDiscardEnabled sets ReplicationQueueRejectMsgToSenderOnDiscardEnabled field to given value.

### HasReplicationQueueRejectMsgToSenderOnDiscardEnabled

`func (o *MsgVpn) HasReplicationQueueRejectMsgToSenderOnDiscardEnabled() bool`

HasReplicationQueueRejectMsgToSenderOnDiscardEnabled returns a boolean if a field has been set.

### GetReplicationRejectMsgWhenSyncIneligibleEnabled

`func (o *MsgVpn) GetReplicationRejectMsgWhenSyncIneligibleEnabled() bool`

GetReplicationRejectMsgWhenSyncIneligibleEnabled returns the ReplicationRejectMsgWhenSyncIneligibleEnabled field if non-nil, zero value otherwise.

### GetReplicationRejectMsgWhenSyncIneligibleEnabledOk

`func (o *MsgVpn) GetReplicationRejectMsgWhenSyncIneligibleEnabledOk() (*bool, bool)`

GetReplicationRejectMsgWhenSyncIneligibleEnabledOk returns a tuple with the ReplicationRejectMsgWhenSyncIneligibleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRejectMsgWhenSyncIneligibleEnabled

`func (o *MsgVpn) SetReplicationRejectMsgWhenSyncIneligibleEnabled(v bool)`

SetReplicationRejectMsgWhenSyncIneligibleEnabled sets ReplicationRejectMsgWhenSyncIneligibleEnabled field to given value.

### HasReplicationRejectMsgWhenSyncIneligibleEnabled

`func (o *MsgVpn) HasReplicationRejectMsgWhenSyncIneligibleEnabled() bool`

HasReplicationRejectMsgWhenSyncIneligibleEnabled returns a boolean if a field has been set.

### GetReplicationRole

`func (o *MsgVpn) GetReplicationRole() string`

GetReplicationRole returns the ReplicationRole field if non-nil, zero value otherwise.

### GetReplicationRoleOk

`func (o *MsgVpn) GetReplicationRoleOk() (*string, bool)`

GetReplicationRoleOk returns a tuple with the ReplicationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRole

`func (o *MsgVpn) SetReplicationRole(v string)`

SetReplicationRole sets ReplicationRole field to given value.

### HasReplicationRole

`func (o *MsgVpn) HasReplicationRole() bool`

HasReplicationRole returns a boolean if a field has been set.

### GetReplicationTransactionMode

`func (o *MsgVpn) GetReplicationTransactionMode() string`

GetReplicationTransactionMode returns the ReplicationTransactionMode field if non-nil, zero value otherwise.

### GetReplicationTransactionModeOk

`func (o *MsgVpn) GetReplicationTransactionModeOk() (*string, bool)`

GetReplicationTransactionModeOk returns a tuple with the ReplicationTransactionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTransactionMode

`func (o *MsgVpn) SetReplicationTransactionMode(v string)`

SetReplicationTransactionMode sets ReplicationTransactionMode field to given value.

### HasReplicationTransactionMode

`func (o *MsgVpn) HasReplicationTransactionMode() bool`

HasReplicationTransactionMode returns a boolean if a field has been set.

### GetRestTlsServerCertEnforceTrustedCommonNameEnabled

`func (o *MsgVpn) GetRestTlsServerCertEnforceTrustedCommonNameEnabled() bool`

GetRestTlsServerCertEnforceTrustedCommonNameEnabled returns the RestTlsServerCertEnforceTrustedCommonNameEnabled field if non-nil, zero value otherwise.

### GetRestTlsServerCertEnforceTrustedCommonNameEnabledOk

`func (o *MsgVpn) GetRestTlsServerCertEnforceTrustedCommonNameEnabledOk() (*bool, bool)`

GetRestTlsServerCertEnforceTrustedCommonNameEnabledOk returns a tuple with the RestTlsServerCertEnforceTrustedCommonNameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestTlsServerCertEnforceTrustedCommonNameEnabled

`func (o *MsgVpn) SetRestTlsServerCertEnforceTrustedCommonNameEnabled(v bool)`

SetRestTlsServerCertEnforceTrustedCommonNameEnabled sets RestTlsServerCertEnforceTrustedCommonNameEnabled field to given value.

### HasRestTlsServerCertEnforceTrustedCommonNameEnabled

`func (o *MsgVpn) HasRestTlsServerCertEnforceTrustedCommonNameEnabled() bool`

HasRestTlsServerCertEnforceTrustedCommonNameEnabled returns a boolean if a field has been set.

### GetRestTlsServerCertMaxChainDepth

`func (o *MsgVpn) GetRestTlsServerCertMaxChainDepth() int64`

GetRestTlsServerCertMaxChainDepth returns the RestTlsServerCertMaxChainDepth field if non-nil, zero value otherwise.

### GetRestTlsServerCertMaxChainDepthOk

`func (o *MsgVpn) GetRestTlsServerCertMaxChainDepthOk() (*int64, bool)`

GetRestTlsServerCertMaxChainDepthOk returns a tuple with the RestTlsServerCertMaxChainDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestTlsServerCertMaxChainDepth

`func (o *MsgVpn) SetRestTlsServerCertMaxChainDepth(v int64)`

SetRestTlsServerCertMaxChainDepth sets RestTlsServerCertMaxChainDepth field to given value.

### HasRestTlsServerCertMaxChainDepth

`func (o *MsgVpn) HasRestTlsServerCertMaxChainDepth() bool`

HasRestTlsServerCertMaxChainDepth returns a boolean if a field has been set.

### GetRestTlsServerCertValidateDateEnabled

`func (o *MsgVpn) GetRestTlsServerCertValidateDateEnabled() bool`

GetRestTlsServerCertValidateDateEnabled returns the RestTlsServerCertValidateDateEnabled field if non-nil, zero value otherwise.

### GetRestTlsServerCertValidateDateEnabledOk

`func (o *MsgVpn) GetRestTlsServerCertValidateDateEnabledOk() (*bool, bool)`

GetRestTlsServerCertValidateDateEnabledOk returns a tuple with the RestTlsServerCertValidateDateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestTlsServerCertValidateDateEnabled

`func (o *MsgVpn) SetRestTlsServerCertValidateDateEnabled(v bool)`

SetRestTlsServerCertValidateDateEnabled sets RestTlsServerCertValidateDateEnabled field to given value.

### HasRestTlsServerCertValidateDateEnabled

`func (o *MsgVpn) HasRestTlsServerCertValidateDateEnabled() bool`

HasRestTlsServerCertValidateDateEnabled returns a boolean if a field has been set.

### GetRestTlsServerCertValidateNameEnabled

`func (o *MsgVpn) GetRestTlsServerCertValidateNameEnabled() bool`

GetRestTlsServerCertValidateNameEnabled returns the RestTlsServerCertValidateNameEnabled field if non-nil, zero value otherwise.

### GetRestTlsServerCertValidateNameEnabledOk

`func (o *MsgVpn) GetRestTlsServerCertValidateNameEnabledOk() (*bool, bool)`

GetRestTlsServerCertValidateNameEnabledOk returns a tuple with the RestTlsServerCertValidateNameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestTlsServerCertValidateNameEnabled

`func (o *MsgVpn) SetRestTlsServerCertValidateNameEnabled(v bool)`

SetRestTlsServerCertValidateNameEnabled sets RestTlsServerCertValidateNameEnabled field to given value.

### HasRestTlsServerCertValidateNameEnabled

`func (o *MsgVpn) HasRestTlsServerCertValidateNameEnabled() bool`

HasRestTlsServerCertValidateNameEnabled returns a boolean if a field has been set.

### GetSempOverMsgBusAdminClientEnabled

`func (o *MsgVpn) GetSempOverMsgBusAdminClientEnabled() bool`

GetSempOverMsgBusAdminClientEnabled returns the SempOverMsgBusAdminClientEnabled field if non-nil, zero value otherwise.

### GetSempOverMsgBusAdminClientEnabledOk

`func (o *MsgVpn) GetSempOverMsgBusAdminClientEnabledOk() (*bool, bool)`

GetSempOverMsgBusAdminClientEnabledOk returns a tuple with the SempOverMsgBusAdminClientEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempOverMsgBusAdminClientEnabled

`func (o *MsgVpn) SetSempOverMsgBusAdminClientEnabled(v bool)`

SetSempOverMsgBusAdminClientEnabled sets SempOverMsgBusAdminClientEnabled field to given value.

### HasSempOverMsgBusAdminClientEnabled

`func (o *MsgVpn) HasSempOverMsgBusAdminClientEnabled() bool`

HasSempOverMsgBusAdminClientEnabled returns a boolean if a field has been set.

### GetSempOverMsgBusAdminDistributedCacheEnabled

`func (o *MsgVpn) GetSempOverMsgBusAdminDistributedCacheEnabled() bool`

GetSempOverMsgBusAdminDistributedCacheEnabled returns the SempOverMsgBusAdminDistributedCacheEnabled field if non-nil, zero value otherwise.

### GetSempOverMsgBusAdminDistributedCacheEnabledOk

`func (o *MsgVpn) GetSempOverMsgBusAdminDistributedCacheEnabledOk() (*bool, bool)`

GetSempOverMsgBusAdminDistributedCacheEnabledOk returns a tuple with the SempOverMsgBusAdminDistributedCacheEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempOverMsgBusAdminDistributedCacheEnabled

`func (o *MsgVpn) SetSempOverMsgBusAdminDistributedCacheEnabled(v bool)`

SetSempOverMsgBusAdminDistributedCacheEnabled sets SempOverMsgBusAdminDistributedCacheEnabled field to given value.

### HasSempOverMsgBusAdminDistributedCacheEnabled

`func (o *MsgVpn) HasSempOverMsgBusAdminDistributedCacheEnabled() bool`

HasSempOverMsgBusAdminDistributedCacheEnabled returns a boolean if a field has been set.

### GetSempOverMsgBusAdminEnabled

`func (o *MsgVpn) GetSempOverMsgBusAdminEnabled() bool`

GetSempOverMsgBusAdminEnabled returns the SempOverMsgBusAdminEnabled field if non-nil, zero value otherwise.

### GetSempOverMsgBusAdminEnabledOk

`func (o *MsgVpn) GetSempOverMsgBusAdminEnabledOk() (*bool, bool)`

GetSempOverMsgBusAdminEnabledOk returns a tuple with the SempOverMsgBusAdminEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempOverMsgBusAdminEnabled

`func (o *MsgVpn) SetSempOverMsgBusAdminEnabled(v bool)`

SetSempOverMsgBusAdminEnabled sets SempOverMsgBusAdminEnabled field to given value.

### HasSempOverMsgBusAdminEnabled

`func (o *MsgVpn) HasSempOverMsgBusAdminEnabled() bool`

HasSempOverMsgBusAdminEnabled returns a boolean if a field has been set.

### GetSempOverMsgBusEnabled

`func (o *MsgVpn) GetSempOverMsgBusEnabled() bool`

GetSempOverMsgBusEnabled returns the SempOverMsgBusEnabled field if non-nil, zero value otherwise.

### GetSempOverMsgBusEnabledOk

`func (o *MsgVpn) GetSempOverMsgBusEnabledOk() (*bool, bool)`

GetSempOverMsgBusEnabledOk returns a tuple with the SempOverMsgBusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempOverMsgBusEnabled

`func (o *MsgVpn) SetSempOverMsgBusEnabled(v bool)`

SetSempOverMsgBusEnabled sets SempOverMsgBusEnabled field to given value.

### HasSempOverMsgBusEnabled

`func (o *MsgVpn) HasSempOverMsgBusEnabled() bool`

HasSempOverMsgBusEnabled returns a boolean if a field has been set.

### GetSempOverMsgBusShowEnabled

`func (o *MsgVpn) GetSempOverMsgBusShowEnabled() bool`

GetSempOverMsgBusShowEnabled returns the SempOverMsgBusShowEnabled field if non-nil, zero value otherwise.

### GetSempOverMsgBusShowEnabledOk

`func (o *MsgVpn) GetSempOverMsgBusShowEnabledOk() (*bool, bool)`

GetSempOverMsgBusShowEnabledOk returns a tuple with the SempOverMsgBusShowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSempOverMsgBusShowEnabled

`func (o *MsgVpn) SetSempOverMsgBusShowEnabled(v bool)`

SetSempOverMsgBusShowEnabled sets SempOverMsgBusShowEnabled field to given value.

### HasSempOverMsgBusShowEnabled

`func (o *MsgVpn) HasSempOverMsgBusShowEnabled() bool`

HasSempOverMsgBusShowEnabled returns a boolean if a field has been set.

### GetServiceAmqpMaxConnectionCount

`func (o *MsgVpn) GetServiceAmqpMaxConnectionCount() int64`

GetServiceAmqpMaxConnectionCount returns the ServiceAmqpMaxConnectionCount field if non-nil, zero value otherwise.

### GetServiceAmqpMaxConnectionCountOk

`func (o *MsgVpn) GetServiceAmqpMaxConnectionCountOk() (*int64, bool)`

GetServiceAmqpMaxConnectionCountOk returns a tuple with the ServiceAmqpMaxConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpMaxConnectionCount

`func (o *MsgVpn) SetServiceAmqpMaxConnectionCount(v int64)`

SetServiceAmqpMaxConnectionCount sets ServiceAmqpMaxConnectionCount field to given value.

### HasServiceAmqpMaxConnectionCount

`func (o *MsgVpn) HasServiceAmqpMaxConnectionCount() bool`

HasServiceAmqpMaxConnectionCount returns a boolean if a field has been set.

### GetServiceAmqpPlainTextEnabled

`func (o *MsgVpn) GetServiceAmqpPlainTextEnabled() bool`

GetServiceAmqpPlainTextEnabled returns the ServiceAmqpPlainTextEnabled field if non-nil, zero value otherwise.

### GetServiceAmqpPlainTextEnabledOk

`func (o *MsgVpn) GetServiceAmqpPlainTextEnabledOk() (*bool, bool)`

GetServiceAmqpPlainTextEnabledOk returns a tuple with the ServiceAmqpPlainTextEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpPlainTextEnabled

`func (o *MsgVpn) SetServiceAmqpPlainTextEnabled(v bool)`

SetServiceAmqpPlainTextEnabled sets ServiceAmqpPlainTextEnabled field to given value.

### HasServiceAmqpPlainTextEnabled

`func (o *MsgVpn) HasServiceAmqpPlainTextEnabled() bool`

HasServiceAmqpPlainTextEnabled returns a boolean if a field has been set.

### GetServiceAmqpPlainTextListenPort

`func (o *MsgVpn) GetServiceAmqpPlainTextListenPort() int64`

GetServiceAmqpPlainTextListenPort returns the ServiceAmqpPlainTextListenPort field if non-nil, zero value otherwise.

### GetServiceAmqpPlainTextListenPortOk

`func (o *MsgVpn) GetServiceAmqpPlainTextListenPortOk() (*int64, bool)`

GetServiceAmqpPlainTextListenPortOk returns a tuple with the ServiceAmqpPlainTextListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpPlainTextListenPort

`func (o *MsgVpn) SetServiceAmqpPlainTextListenPort(v int64)`

SetServiceAmqpPlainTextListenPort sets ServiceAmqpPlainTextListenPort field to given value.

### HasServiceAmqpPlainTextListenPort

`func (o *MsgVpn) HasServiceAmqpPlainTextListenPort() bool`

HasServiceAmqpPlainTextListenPort returns a boolean if a field has been set.

### GetServiceAmqpTlsEnabled

`func (o *MsgVpn) GetServiceAmqpTlsEnabled() bool`

GetServiceAmqpTlsEnabled returns the ServiceAmqpTlsEnabled field if non-nil, zero value otherwise.

### GetServiceAmqpTlsEnabledOk

`func (o *MsgVpn) GetServiceAmqpTlsEnabledOk() (*bool, bool)`

GetServiceAmqpTlsEnabledOk returns a tuple with the ServiceAmqpTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpTlsEnabled

`func (o *MsgVpn) SetServiceAmqpTlsEnabled(v bool)`

SetServiceAmqpTlsEnabled sets ServiceAmqpTlsEnabled field to given value.

### HasServiceAmqpTlsEnabled

`func (o *MsgVpn) HasServiceAmqpTlsEnabled() bool`

HasServiceAmqpTlsEnabled returns a boolean if a field has been set.

### GetServiceAmqpTlsListenPort

`func (o *MsgVpn) GetServiceAmqpTlsListenPort() int64`

GetServiceAmqpTlsListenPort returns the ServiceAmqpTlsListenPort field if non-nil, zero value otherwise.

### GetServiceAmqpTlsListenPortOk

`func (o *MsgVpn) GetServiceAmqpTlsListenPortOk() (*int64, bool)`

GetServiceAmqpTlsListenPortOk returns a tuple with the ServiceAmqpTlsListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpTlsListenPort

`func (o *MsgVpn) SetServiceAmqpTlsListenPort(v int64)`

SetServiceAmqpTlsListenPort sets ServiceAmqpTlsListenPort field to given value.

### HasServiceAmqpTlsListenPort

`func (o *MsgVpn) HasServiceAmqpTlsListenPort() bool`

HasServiceAmqpTlsListenPort returns a boolean if a field has been set.

### GetServiceMqttAuthenticationClientCertRequest

`func (o *MsgVpn) GetServiceMqttAuthenticationClientCertRequest() string`

GetServiceMqttAuthenticationClientCertRequest returns the ServiceMqttAuthenticationClientCertRequest field if non-nil, zero value otherwise.

### GetServiceMqttAuthenticationClientCertRequestOk

`func (o *MsgVpn) GetServiceMqttAuthenticationClientCertRequestOk() (*string, bool)`

GetServiceMqttAuthenticationClientCertRequestOk returns a tuple with the ServiceMqttAuthenticationClientCertRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttAuthenticationClientCertRequest

`func (o *MsgVpn) SetServiceMqttAuthenticationClientCertRequest(v string)`

SetServiceMqttAuthenticationClientCertRequest sets ServiceMqttAuthenticationClientCertRequest field to given value.

### HasServiceMqttAuthenticationClientCertRequest

`func (o *MsgVpn) HasServiceMqttAuthenticationClientCertRequest() bool`

HasServiceMqttAuthenticationClientCertRequest returns a boolean if a field has been set.

### GetServiceMqttMaxConnectionCount

`func (o *MsgVpn) GetServiceMqttMaxConnectionCount() int64`

GetServiceMqttMaxConnectionCount returns the ServiceMqttMaxConnectionCount field if non-nil, zero value otherwise.

### GetServiceMqttMaxConnectionCountOk

`func (o *MsgVpn) GetServiceMqttMaxConnectionCountOk() (*int64, bool)`

GetServiceMqttMaxConnectionCountOk returns a tuple with the ServiceMqttMaxConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttMaxConnectionCount

`func (o *MsgVpn) SetServiceMqttMaxConnectionCount(v int64)`

SetServiceMqttMaxConnectionCount sets ServiceMqttMaxConnectionCount field to given value.

### HasServiceMqttMaxConnectionCount

`func (o *MsgVpn) HasServiceMqttMaxConnectionCount() bool`

HasServiceMqttMaxConnectionCount returns a boolean if a field has been set.

### GetServiceMqttPlainTextEnabled

`func (o *MsgVpn) GetServiceMqttPlainTextEnabled() bool`

GetServiceMqttPlainTextEnabled returns the ServiceMqttPlainTextEnabled field if non-nil, zero value otherwise.

### GetServiceMqttPlainTextEnabledOk

`func (o *MsgVpn) GetServiceMqttPlainTextEnabledOk() (*bool, bool)`

GetServiceMqttPlainTextEnabledOk returns a tuple with the ServiceMqttPlainTextEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttPlainTextEnabled

`func (o *MsgVpn) SetServiceMqttPlainTextEnabled(v bool)`

SetServiceMqttPlainTextEnabled sets ServiceMqttPlainTextEnabled field to given value.

### HasServiceMqttPlainTextEnabled

`func (o *MsgVpn) HasServiceMqttPlainTextEnabled() bool`

HasServiceMqttPlainTextEnabled returns a boolean if a field has been set.

### GetServiceMqttPlainTextListenPort

`func (o *MsgVpn) GetServiceMqttPlainTextListenPort() int64`

GetServiceMqttPlainTextListenPort returns the ServiceMqttPlainTextListenPort field if non-nil, zero value otherwise.

### GetServiceMqttPlainTextListenPortOk

`func (o *MsgVpn) GetServiceMqttPlainTextListenPortOk() (*int64, bool)`

GetServiceMqttPlainTextListenPortOk returns a tuple with the ServiceMqttPlainTextListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttPlainTextListenPort

`func (o *MsgVpn) SetServiceMqttPlainTextListenPort(v int64)`

SetServiceMqttPlainTextListenPort sets ServiceMqttPlainTextListenPort field to given value.

### HasServiceMqttPlainTextListenPort

`func (o *MsgVpn) HasServiceMqttPlainTextListenPort() bool`

HasServiceMqttPlainTextListenPort returns a boolean if a field has been set.

### GetServiceMqttTlsEnabled

`func (o *MsgVpn) GetServiceMqttTlsEnabled() bool`

GetServiceMqttTlsEnabled returns the ServiceMqttTlsEnabled field if non-nil, zero value otherwise.

### GetServiceMqttTlsEnabledOk

`func (o *MsgVpn) GetServiceMqttTlsEnabledOk() (*bool, bool)`

GetServiceMqttTlsEnabledOk returns a tuple with the ServiceMqttTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttTlsEnabled

`func (o *MsgVpn) SetServiceMqttTlsEnabled(v bool)`

SetServiceMqttTlsEnabled sets ServiceMqttTlsEnabled field to given value.

### HasServiceMqttTlsEnabled

`func (o *MsgVpn) HasServiceMqttTlsEnabled() bool`

HasServiceMqttTlsEnabled returns a boolean if a field has been set.

### GetServiceMqttTlsListenPort

`func (o *MsgVpn) GetServiceMqttTlsListenPort() int64`

GetServiceMqttTlsListenPort returns the ServiceMqttTlsListenPort field if non-nil, zero value otherwise.

### GetServiceMqttTlsListenPortOk

`func (o *MsgVpn) GetServiceMqttTlsListenPortOk() (*int64, bool)`

GetServiceMqttTlsListenPortOk returns a tuple with the ServiceMqttTlsListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttTlsListenPort

`func (o *MsgVpn) SetServiceMqttTlsListenPort(v int64)`

SetServiceMqttTlsListenPort sets ServiceMqttTlsListenPort field to given value.

### HasServiceMqttTlsListenPort

`func (o *MsgVpn) HasServiceMqttTlsListenPort() bool`

HasServiceMqttTlsListenPort returns a boolean if a field has been set.

### GetServiceMqttTlsWebSocketEnabled

`func (o *MsgVpn) GetServiceMqttTlsWebSocketEnabled() bool`

GetServiceMqttTlsWebSocketEnabled returns the ServiceMqttTlsWebSocketEnabled field if non-nil, zero value otherwise.

### GetServiceMqttTlsWebSocketEnabledOk

`func (o *MsgVpn) GetServiceMqttTlsWebSocketEnabledOk() (*bool, bool)`

GetServiceMqttTlsWebSocketEnabledOk returns a tuple with the ServiceMqttTlsWebSocketEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttTlsWebSocketEnabled

`func (o *MsgVpn) SetServiceMqttTlsWebSocketEnabled(v bool)`

SetServiceMqttTlsWebSocketEnabled sets ServiceMqttTlsWebSocketEnabled field to given value.

### HasServiceMqttTlsWebSocketEnabled

`func (o *MsgVpn) HasServiceMqttTlsWebSocketEnabled() bool`

HasServiceMqttTlsWebSocketEnabled returns a boolean if a field has been set.

### GetServiceMqttTlsWebSocketListenPort

`func (o *MsgVpn) GetServiceMqttTlsWebSocketListenPort() int64`

GetServiceMqttTlsWebSocketListenPort returns the ServiceMqttTlsWebSocketListenPort field if non-nil, zero value otherwise.

### GetServiceMqttTlsWebSocketListenPortOk

`func (o *MsgVpn) GetServiceMqttTlsWebSocketListenPortOk() (*int64, bool)`

GetServiceMqttTlsWebSocketListenPortOk returns a tuple with the ServiceMqttTlsWebSocketListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttTlsWebSocketListenPort

`func (o *MsgVpn) SetServiceMqttTlsWebSocketListenPort(v int64)`

SetServiceMqttTlsWebSocketListenPort sets ServiceMqttTlsWebSocketListenPort field to given value.

### HasServiceMqttTlsWebSocketListenPort

`func (o *MsgVpn) HasServiceMqttTlsWebSocketListenPort() bool`

HasServiceMqttTlsWebSocketListenPort returns a boolean if a field has been set.

### GetServiceMqttWebSocketEnabled

`func (o *MsgVpn) GetServiceMqttWebSocketEnabled() bool`

GetServiceMqttWebSocketEnabled returns the ServiceMqttWebSocketEnabled field if non-nil, zero value otherwise.

### GetServiceMqttWebSocketEnabledOk

`func (o *MsgVpn) GetServiceMqttWebSocketEnabledOk() (*bool, bool)`

GetServiceMqttWebSocketEnabledOk returns a tuple with the ServiceMqttWebSocketEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttWebSocketEnabled

`func (o *MsgVpn) SetServiceMqttWebSocketEnabled(v bool)`

SetServiceMqttWebSocketEnabled sets ServiceMqttWebSocketEnabled field to given value.

### HasServiceMqttWebSocketEnabled

`func (o *MsgVpn) HasServiceMqttWebSocketEnabled() bool`

HasServiceMqttWebSocketEnabled returns a boolean if a field has been set.

### GetServiceMqttWebSocketListenPort

`func (o *MsgVpn) GetServiceMqttWebSocketListenPort() int64`

GetServiceMqttWebSocketListenPort returns the ServiceMqttWebSocketListenPort field if non-nil, zero value otherwise.

### GetServiceMqttWebSocketListenPortOk

`func (o *MsgVpn) GetServiceMqttWebSocketListenPortOk() (*int64, bool)`

GetServiceMqttWebSocketListenPortOk returns a tuple with the ServiceMqttWebSocketListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttWebSocketListenPort

`func (o *MsgVpn) SetServiceMqttWebSocketListenPort(v int64)`

SetServiceMqttWebSocketListenPort sets ServiceMqttWebSocketListenPort field to given value.

### HasServiceMqttWebSocketListenPort

`func (o *MsgVpn) HasServiceMqttWebSocketListenPort() bool`

HasServiceMqttWebSocketListenPort returns a boolean if a field has been set.

### GetServiceRestIncomingAuthenticationClientCertRequest

`func (o *MsgVpn) GetServiceRestIncomingAuthenticationClientCertRequest() string`

GetServiceRestIncomingAuthenticationClientCertRequest returns the ServiceRestIncomingAuthenticationClientCertRequest field if non-nil, zero value otherwise.

### GetServiceRestIncomingAuthenticationClientCertRequestOk

`func (o *MsgVpn) GetServiceRestIncomingAuthenticationClientCertRequestOk() (*string, bool)`

GetServiceRestIncomingAuthenticationClientCertRequestOk returns a tuple with the ServiceRestIncomingAuthenticationClientCertRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingAuthenticationClientCertRequest

`func (o *MsgVpn) SetServiceRestIncomingAuthenticationClientCertRequest(v string)`

SetServiceRestIncomingAuthenticationClientCertRequest sets ServiceRestIncomingAuthenticationClientCertRequest field to given value.

### HasServiceRestIncomingAuthenticationClientCertRequest

`func (o *MsgVpn) HasServiceRestIncomingAuthenticationClientCertRequest() bool`

HasServiceRestIncomingAuthenticationClientCertRequest returns a boolean if a field has been set.

### GetServiceRestIncomingAuthorizationHeaderHandling

`func (o *MsgVpn) GetServiceRestIncomingAuthorizationHeaderHandling() string`

GetServiceRestIncomingAuthorizationHeaderHandling returns the ServiceRestIncomingAuthorizationHeaderHandling field if non-nil, zero value otherwise.

### GetServiceRestIncomingAuthorizationHeaderHandlingOk

`func (o *MsgVpn) GetServiceRestIncomingAuthorizationHeaderHandlingOk() (*string, bool)`

GetServiceRestIncomingAuthorizationHeaderHandlingOk returns a tuple with the ServiceRestIncomingAuthorizationHeaderHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingAuthorizationHeaderHandling

`func (o *MsgVpn) SetServiceRestIncomingAuthorizationHeaderHandling(v string)`

SetServiceRestIncomingAuthorizationHeaderHandling sets ServiceRestIncomingAuthorizationHeaderHandling field to given value.

### HasServiceRestIncomingAuthorizationHeaderHandling

`func (o *MsgVpn) HasServiceRestIncomingAuthorizationHeaderHandling() bool`

HasServiceRestIncomingAuthorizationHeaderHandling returns a boolean if a field has been set.

### GetServiceRestIncomingMaxConnectionCount

`func (o *MsgVpn) GetServiceRestIncomingMaxConnectionCount() int64`

GetServiceRestIncomingMaxConnectionCount returns the ServiceRestIncomingMaxConnectionCount field if non-nil, zero value otherwise.

### GetServiceRestIncomingMaxConnectionCountOk

`func (o *MsgVpn) GetServiceRestIncomingMaxConnectionCountOk() (*int64, bool)`

GetServiceRestIncomingMaxConnectionCountOk returns a tuple with the ServiceRestIncomingMaxConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingMaxConnectionCount

`func (o *MsgVpn) SetServiceRestIncomingMaxConnectionCount(v int64)`

SetServiceRestIncomingMaxConnectionCount sets ServiceRestIncomingMaxConnectionCount field to given value.

### HasServiceRestIncomingMaxConnectionCount

`func (o *MsgVpn) HasServiceRestIncomingMaxConnectionCount() bool`

HasServiceRestIncomingMaxConnectionCount returns a boolean if a field has been set.

### GetServiceRestIncomingPlainTextEnabled

`func (o *MsgVpn) GetServiceRestIncomingPlainTextEnabled() bool`

GetServiceRestIncomingPlainTextEnabled returns the ServiceRestIncomingPlainTextEnabled field if non-nil, zero value otherwise.

### GetServiceRestIncomingPlainTextEnabledOk

`func (o *MsgVpn) GetServiceRestIncomingPlainTextEnabledOk() (*bool, bool)`

GetServiceRestIncomingPlainTextEnabledOk returns a tuple with the ServiceRestIncomingPlainTextEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingPlainTextEnabled

`func (o *MsgVpn) SetServiceRestIncomingPlainTextEnabled(v bool)`

SetServiceRestIncomingPlainTextEnabled sets ServiceRestIncomingPlainTextEnabled field to given value.

### HasServiceRestIncomingPlainTextEnabled

`func (o *MsgVpn) HasServiceRestIncomingPlainTextEnabled() bool`

HasServiceRestIncomingPlainTextEnabled returns a boolean if a field has been set.

### GetServiceRestIncomingPlainTextListenPort

`func (o *MsgVpn) GetServiceRestIncomingPlainTextListenPort() int64`

GetServiceRestIncomingPlainTextListenPort returns the ServiceRestIncomingPlainTextListenPort field if non-nil, zero value otherwise.

### GetServiceRestIncomingPlainTextListenPortOk

`func (o *MsgVpn) GetServiceRestIncomingPlainTextListenPortOk() (*int64, bool)`

GetServiceRestIncomingPlainTextListenPortOk returns a tuple with the ServiceRestIncomingPlainTextListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingPlainTextListenPort

`func (o *MsgVpn) SetServiceRestIncomingPlainTextListenPort(v int64)`

SetServiceRestIncomingPlainTextListenPort sets ServiceRestIncomingPlainTextListenPort field to given value.

### HasServiceRestIncomingPlainTextListenPort

`func (o *MsgVpn) HasServiceRestIncomingPlainTextListenPort() bool`

HasServiceRestIncomingPlainTextListenPort returns a boolean if a field has been set.

### GetServiceRestIncomingTlsEnabled

`func (o *MsgVpn) GetServiceRestIncomingTlsEnabled() bool`

GetServiceRestIncomingTlsEnabled returns the ServiceRestIncomingTlsEnabled field if non-nil, zero value otherwise.

### GetServiceRestIncomingTlsEnabledOk

`func (o *MsgVpn) GetServiceRestIncomingTlsEnabledOk() (*bool, bool)`

GetServiceRestIncomingTlsEnabledOk returns a tuple with the ServiceRestIncomingTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingTlsEnabled

`func (o *MsgVpn) SetServiceRestIncomingTlsEnabled(v bool)`

SetServiceRestIncomingTlsEnabled sets ServiceRestIncomingTlsEnabled field to given value.

### HasServiceRestIncomingTlsEnabled

`func (o *MsgVpn) HasServiceRestIncomingTlsEnabled() bool`

HasServiceRestIncomingTlsEnabled returns a boolean if a field has been set.

### GetServiceRestIncomingTlsListenPort

`func (o *MsgVpn) GetServiceRestIncomingTlsListenPort() int64`

GetServiceRestIncomingTlsListenPort returns the ServiceRestIncomingTlsListenPort field if non-nil, zero value otherwise.

### GetServiceRestIncomingTlsListenPortOk

`func (o *MsgVpn) GetServiceRestIncomingTlsListenPortOk() (*int64, bool)`

GetServiceRestIncomingTlsListenPortOk returns a tuple with the ServiceRestIncomingTlsListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingTlsListenPort

`func (o *MsgVpn) SetServiceRestIncomingTlsListenPort(v int64)`

SetServiceRestIncomingTlsListenPort sets ServiceRestIncomingTlsListenPort field to given value.

### HasServiceRestIncomingTlsListenPort

`func (o *MsgVpn) HasServiceRestIncomingTlsListenPort() bool`

HasServiceRestIncomingTlsListenPort returns a boolean if a field has been set.

### GetServiceRestMode

`func (o *MsgVpn) GetServiceRestMode() string`

GetServiceRestMode returns the ServiceRestMode field if non-nil, zero value otherwise.

### GetServiceRestModeOk

`func (o *MsgVpn) GetServiceRestModeOk() (*string, bool)`

GetServiceRestModeOk returns a tuple with the ServiceRestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestMode

`func (o *MsgVpn) SetServiceRestMode(v string)`

SetServiceRestMode sets ServiceRestMode field to given value.

### HasServiceRestMode

`func (o *MsgVpn) HasServiceRestMode() bool`

HasServiceRestMode returns a boolean if a field has been set.

### GetServiceRestOutgoingMaxConnectionCount

`func (o *MsgVpn) GetServiceRestOutgoingMaxConnectionCount() int64`

GetServiceRestOutgoingMaxConnectionCount returns the ServiceRestOutgoingMaxConnectionCount field if non-nil, zero value otherwise.

### GetServiceRestOutgoingMaxConnectionCountOk

`func (o *MsgVpn) GetServiceRestOutgoingMaxConnectionCountOk() (*int64, bool)`

GetServiceRestOutgoingMaxConnectionCountOk returns a tuple with the ServiceRestOutgoingMaxConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestOutgoingMaxConnectionCount

`func (o *MsgVpn) SetServiceRestOutgoingMaxConnectionCount(v int64)`

SetServiceRestOutgoingMaxConnectionCount sets ServiceRestOutgoingMaxConnectionCount field to given value.

### HasServiceRestOutgoingMaxConnectionCount

`func (o *MsgVpn) HasServiceRestOutgoingMaxConnectionCount() bool`

HasServiceRestOutgoingMaxConnectionCount returns a boolean if a field has been set.

### GetServiceSmfMaxConnectionCount

`func (o *MsgVpn) GetServiceSmfMaxConnectionCount() int64`

GetServiceSmfMaxConnectionCount returns the ServiceSmfMaxConnectionCount field if non-nil, zero value otherwise.

### GetServiceSmfMaxConnectionCountOk

`func (o *MsgVpn) GetServiceSmfMaxConnectionCountOk() (*int64, bool)`

GetServiceSmfMaxConnectionCountOk returns a tuple with the ServiceSmfMaxConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfMaxConnectionCount

`func (o *MsgVpn) SetServiceSmfMaxConnectionCount(v int64)`

SetServiceSmfMaxConnectionCount sets ServiceSmfMaxConnectionCount field to given value.

### HasServiceSmfMaxConnectionCount

`func (o *MsgVpn) HasServiceSmfMaxConnectionCount() bool`

HasServiceSmfMaxConnectionCount returns a boolean if a field has been set.

### GetServiceSmfPlainTextEnabled

`func (o *MsgVpn) GetServiceSmfPlainTextEnabled() bool`

GetServiceSmfPlainTextEnabled returns the ServiceSmfPlainTextEnabled field if non-nil, zero value otherwise.

### GetServiceSmfPlainTextEnabledOk

`func (o *MsgVpn) GetServiceSmfPlainTextEnabledOk() (*bool, bool)`

GetServiceSmfPlainTextEnabledOk returns a tuple with the ServiceSmfPlainTextEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfPlainTextEnabled

`func (o *MsgVpn) SetServiceSmfPlainTextEnabled(v bool)`

SetServiceSmfPlainTextEnabled sets ServiceSmfPlainTextEnabled field to given value.

### HasServiceSmfPlainTextEnabled

`func (o *MsgVpn) HasServiceSmfPlainTextEnabled() bool`

HasServiceSmfPlainTextEnabled returns a boolean if a field has been set.

### GetServiceSmfTlsEnabled

`func (o *MsgVpn) GetServiceSmfTlsEnabled() bool`

GetServiceSmfTlsEnabled returns the ServiceSmfTlsEnabled field if non-nil, zero value otherwise.

### GetServiceSmfTlsEnabledOk

`func (o *MsgVpn) GetServiceSmfTlsEnabledOk() (*bool, bool)`

GetServiceSmfTlsEnabledOk returns a tuple with the ServiceSmfTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfTlsEnabled

`func (o *MsgVpn) SetServiceSmfTlsEnabled(v bool)`

SetServiceSmfTlsEnabled sets ServiceSmfTlsEnabled field to given value.

### HasServiceSmfTlsEnabled

`func (o *MsgVpn) HasServiceSmfTlsEnabled() bool`

HasServiceSmfTlsEnabled returns a boolean if a field has been set.

### GetServiceWebAuthenticationClientCertRequest

`func (o *MsgVpn) GetServiceWebAuthenticationClientCertRequest() string`

GetServiceWebAuthenticationClientCertRequest returns the ServiceWebAuthenticationClientCertRequest field if non-nil, zero value otherwise.

### GetServiceWebAuthenticationClientCertRequestOk

`func (o *MsgVpn) GetServiceWebAuthenticationClientCertRequestOk() (*string, bool)`

GetServiceWebAuthenticationClientCertRequestOk returns a tuple with the ServiceWebAuthenticationClientCertRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebAuthenticationClientCertRequest

`func (o *MsgVpn) SetServiceWebAuthenticationClientCertRequest(v string)`

SetServiceWebAuthenticationClientCertRequest sets ServiceWebAuthenticationClientCertRequest field to given value.

### HasServiceWebAuthenticationClientCertRequest

`func (o *MsgVpn) HasServiceWebAuthenticationClientCertRequest() bool`

HasServiceWebAuthenticationClientCertRequest returns a boolean if a field has been set.

### GetServiceWebMaxConnectionCount

`func (o *MsgVpn) GetServiceWebMaxConnectionCount() int64`

GetServiceWebMaxConnectionCount returns the ServiceWebMaxConnectionCount field if non-nil, zero value otherwise.

### GetServiceWebMaxConnectionCountOk

`func (o *MsgVpn) GetServiceWebMaxConnectionCountOk() (*int64, bool)`

GetServiceWebMaxConnectionCountOk returns a tuple with the ServiceWebMaxConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebMaxConnectionCount

`func (o *MsgVpn) SetServiceWebMaxConnectionCount(v int64)`

SetServiceWebMaxConnectionCount sets ServiceWebMaxConnectionCount field to given value.

### HasServiceWebMaxConnectionCount

`func (o *MsgVpn) HasServiceWebMaxConnectionCount() bool`

HasServiceWebMaxConnectionCount returns a boolean if a field has been set.

### GetServiceWebPlainTextEnabled

`func (o *MsgVpn) GetServiceWebPlainTextEnabled() bool`

GetServiceWebPlainTextEnabled returns the ServiceWebPlainTextEnabled field if non-nil, zero value otherwise.

### GetServiceWebPlainTextEnabledOk

`func (o *MsgVpn) GetServiceWebPlainTextEnabledOk() (*bool, bool)`

GetServiceWebPlainTextEnabledOk returns a tuple with the ServiceWebPlainTextEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebPlainTextEnabled

`func (o *MsgVpn) SetServiceWebPlainTextEnabled(v bool)`

SetServiceWebPlainTextEnabled sets ServiceWebPlainTextEnabled field to given value.

### HasServiceWebPlainTextEnabled

`func (o *MsgVpn) HasServiceWebPlainTextEnabled() bool`

HasServiceWebPlainTextEnabled returns a boolean if a field has been set.

### GetServiceWebTlsEnabled

`func (o *MsgVpn) GetServiceWebTlsEnabled() bool`

GetServiceWebTlsEnabled returns the ServiceWebTlsEnabled field if non-nil, zero value otherwise.

### GetServiceWebTlsEnabledOk

`func (o *MsgVpn) GetServiceWebTlsEnabledOk() (*bool, bool)`

GetServiceWebTlsEnabledOk returns a tuple with the ServiceWebTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebTlsEnabled

`func (o *MsgVpn) SetServiceWebTlsEnabled(v bool)`

SetServiceWebTlsEnabled sets ServiceWebTlsEnabled field to given value.

### HasServiceWebTlsEnabled

`func (o *MsgVpn) HasServiceWebTlsEnabled() bool`

HasServiceWebTlsEnabled returns a boolean if a field has been set.

### GetTlsAllowDowngradeToPlainTextEnabled

`func (o *MsgVpn) GetTlsAllowDowngradeToPlainTextEnabled() bool`

GetTlsAllowDowngradeToPlainTextEnabled returns the TlsAllowDowngradeToPlainTextEnabled field if non-nil, zero value otherwise.

### GetTlsAllowDowngradeToPlainTextEnabledOk

`func (o *MsgVpn) GetTlsAllowDowngradeToPlainTextEnabledOk() (*bool, bool)`

GetTlsAllowDowngradeToPlainTextEnabledOk returns a tuple with the TlsAllowDowngradeToPlainTextEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAllowDowngradeToPlainTextEnabled

`func (o *MsgVpn) SetTlsAllowDowngradeToPlainTextEnabled(v bool)`

SetTlsAllowDowngradeToPlainTextEnabled sets TlsAllowDowngradeToPlainTextEnabled field to given value.

### HasTlsAllowDowngradeToPlainTextEnabled

`func (o *MsgVpn) HasTlsAllowDowngradeToPlainTextEnabled() bool`

HasTlsAllowDowngradeToPlainTextEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


