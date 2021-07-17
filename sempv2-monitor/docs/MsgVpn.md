# MsgVpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | The name of another Message VPN which this Message VPN is an alias for. Available since 2.14. | [optional] 
**AuthenticationBasicEnabled** | Pointer to **bool** | Indicates whether basic authentication is enabled for clients connecting to the Message VPN. | [optional] 
**AuthenticationBasicProfileName** | Pointer to **string** | The name of the RADIUS or LDAP Profile to use for basic authentication. | [optional] 
**AuthenticationBasicRadiusDomain** | Pointer to **string** | The RADIUS domain to use for basic authentication. | [optional] 
**AuthenticationBasicType** | Pointer to **string** | The type of basic authentication to use for clients connecting to the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;internal\&quot; - Internal database. Authentication is against Client Usernames. \&quot;ldap\&quot; - LDAP authentication. An LDAP profile name must be provided. \&quot;radius\&quot; - RADIUS authentication. A RADIUS profile name must be provided. \&quot;none\&quot; - No authentication. Anonymous login allowed. &lt;/pre&gt;  | [optional] 
**AuthenticationClientCertAllowApiProvidedUsernameEnabled** | Pointer to **bool** | Indicates whether a client is allowed to specify a Client Username via the API connect method. When disabled, the certificate CN (Common Name) is always used. | [optional] 
**AuthenticationClientCertEnabled** | Pointer to **bool** | Indicates whether client certificate authentication is enabled in the Message VPN. | [optional] 
**AuthenticationClientCertMaxChainDepth** | Pointer to **int64** | The maximum depth for a client certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. | [optional] 
**AuthenticationClientCertRevocationCheckMode** | Pointer to **string** | The desired behavior for client certificate revocation checking. The allowed values and their meaning are:  &lt;pre&gt; \&quot;allow-all\&quot; - Allow the client to authenticate, the result of client certificate revocation check is ignored. \&quot;allow-unknown\&quot; - Allow the client to authenticate even if the revocation status of his certificate cannot be determined. \&quot;allow-valid\&quot; - Allow the client to authenticate only when the revocation check returned an explicit positive response. &lt;/pre&gt;  | [optional] 
**AuthenticationClientCertUsernameSource** | Pointer to **string** | The field from the client certificate to use as the client username. The allowed values and their meaning are:  &lt;pre&gt; \&quot;certificate-thumbprint\&quot; - The username is computed as the SHA-1 hash over the entire DER-encoded contents of the client certificate. \&quot;common-name\&quot; - The username is extracted from the certificate&#39;s first instance of the Common Name attribute in the Subject DN. \&quot;common-name-last\&quot; - The username is extracted from the certificate&#39;s last instance of the Common Name attribute in the Subject DN. \&quot;subject-alternate-name-msupn\&quot; - The username is extracted from the certificate&#39;s Other Name type of the Subject Alternative Name and must have the msUPN signature. \&quot;uid\&quot; - The username is extracted from the certificate&#39;s first instance of the User Identifier attribute in the Subject DN. \&quot;uid-last\&quot; - The username is extracted from the certificate&#39;s last instance of the User Identifier attribute in the Subject DN. &lt;/pre&gt;  | [optional] 
**AuthenticationClientCertValidateDateEnabled** | Pointer to **bool** | Indicates whether the \&quot;Not Before\&quot; and \&quot;Not After\&quot; validity dates in the client certificate are checked. | [optional] 
**AuthenticationKerberosAllowApiProvidedUsernameEnabled** | Pointer to **bool** | Indicates whether a client is allowed to specify a Client Username via the API connect method. When disabled, the Kerberos Principal name is always used. | [optional] 
**AuthenticationKerberosEnabled** | Pointer to **bool** | Indicates whether Kerberos authentication is enabled in the Message VPN. | [optional] 
**AuthenticationOauthDefaultProviderName** | Pointer to **string** | The name of the provider to use when the client does not supply a provider name. Available since 2.13. | [optional] 
**AuthenticationOauthEnabled** | Pointer to **bool** | Indicates whether OAuth authentication is enabled. Available since 2.13. | [optional] 
**AuthorizationLdapGroupMembershipAttributeName** | Pointer to **string** | The name of the attribute that is retrieved from the LDAP server as part of the LDAP search when authorizing a client connecting to the Message VPN. | [optional] 
**AuthorizationLdapTrimClientUsernameDomainEnabled** | Pointer to **bool** | Indicates whether client-username domain trimming for LDAP lookups of client connections is enabled. Available since 2.13. | [optional] 
**AuthorizationProfileName** | Pointer to **string** | The name of the LDAP Profile to use for client authorization. | [optional] 
**AuthorizationType** | Pointer to **string** | The type of authorization to use for clients connecting to the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;ldap\&quot; - LDAP authorization. \&quot;internal\&quot; - Internal authorization. &lt;/pre&gt;  | [optional] 
**AverageRxByteRate** | Pointer to **int64** | The one minute average of the message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13. | [optional] 
**AverageRxCompressedByteRate** | Pointer to **int64** | The one minute average of the compressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12. | [optional] 
**AverageRxMsgRate** | Pointer to **int64** | The one minute average of the message rate received by the Message VPN, in messages per second (msg/sec). Available since 2.13. | [optional] 
**AverageRxUncompressedByteRate** | Pointer to **int64** | The one minute average of the uncompressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12. | [optional] 
**AverageTxByteRate** | Pointer to **int64** | The one minute average of the message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13. | [optional] 
**AverageTxCompressedByteRate** | Pointer to **int64** | The one minute average of the compressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12. | [optional] 
**AverageTxMsgRate** | Pointer to **int64** | The one minute average of the message rate transmitted by the Message VPN, in messages per second (msg/sec). Available since 2.13. | [optional] 
**AverageTxUncompressedByteRate** | Pointer to **int64** | The one minute average of the uncompressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12. | [optional] 
**BridgingTlsServerCertEnforceTrustedCommonNameEnabled** | Pointer to **bool** | Indicates whether the Common Name (CN) in the server certificate from the remote broker is validated for the Bridge. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**BridgingTlsServerCertMaxChainDepth** | Pointer to **int64** | The maximum depth for a server certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. | [optional] 
**BridgingTlsServerCertValidateDateEnabled** | Pointer to **bool** | Indicates whether the \&quot;Not Before\&quot; and \&quot;Not After\&quot; validity dates in the server certificate are checked. | [optional] 
**BridgingTlsServerCertValidateNameEnabled** | Pointer to **bool** | Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the bridge. If enabled, the name used to connect to the bridge is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. Available since 2.18. | [optional] 
**ConfigSyncLocalKey** | Pointer to **string** | The key for the config sync table of the local Message VPN. Available since 2.12. | [optional] 
**ConfigSyncLocalLastResult** | Pointer to **string** | The result of the last operation on the config sync table of the local Message VPN. Available since 2.12. | [optional] 
**ConfigSyncLocalRole** | Pointer to **string** | The role of the config sync table of the local Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;unknown\&quot; - The role is unknown. \&quot;primary\&quot; - Acts as the primary source of config data. \&quot;replica\&quot; - Acts as a replica of the primary config data. &lt;/pre&gt;  Available since 2.12. | [optional] 
**ConfigSyncLocalState** | Pointer to **string** | The state of the config sync table of the local Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;unknown\&quot; - The state is unknown. \&quot;in-sync\&quot; - The config data is synchronized between Message VPNs. \&quot;reconciling\&quot; - The config data is reconciling between Message VPNs. \&quot;blocked\&quot; - The config data is blocked from reconciling due to an error. \&quot;out-of-sync\&quot; - The config data is out of sync between Message VPNs. \&quot;down\&quot; - The state is down due to configuration. &lt;/pre&gt;  Available since 2.12. | [optional] 
**ConfigSyncLocalTimeInState** | Pointer to **int32** | The amount of time in seconds the config sync table of the local Message VPN has been in the current state. Available since 2.12. | [optional] 
**ControlRxByteCount** | Pointer to **int64** | The amount of client control messages received from clients by the Message VPN, in bytes (B). Available since 2.13. | [optional] 
**ControlRxMsgCount** | Pointer to **int64** | The number of client control messages received from clients by the Message VPN. Available since 2.13. | [optional] 
**ControlTxByteCount** | Pointer to **int64** | The amount of client control messages transmitted to clients by the Message VPN, in bytes (B). Available since 2.13. | [optional] 
**ControlTxMsgCount** | Pointer to **int64** | The number of client control messages transmitted to clients by the Message VPN. Available since 2.13. | [optional] 
**Counter** | Pointer to [**MsgVpnCounter**](MsgVpnCounter.md) |  | [optional] 
**DataRxByteCount** | Pointer to **int64** | The amount of client data messages received from clients by the Message VPN, in bytes (B). Available since 2.13. | [optional] 
**DataRxMsgCount** | Pointer to **int64** | The number of client data messages received from clients by the Message VPN. Available since 2.13. | [optional] 
**DataTxByteCount** | Pointer to **int64** | The amount of client data messages transmitted to clients by the Message VPN, in bytes (B). Available since 2.13. | [optional] 
**DataTxMsgCount** | Pointer to **int64** | The number of client data messages transmitted to clients by the Message VPN. Available since 2.13. | [optional] 
**DiscardedRxMsgCount** | Pointer to **int64** | The number of messages discarded during reception by the Message VPN. Available since 2.13. | [optional] 
**DiscardedTxMsgCount** | Pointer to **int64** | The number of messages discarded during transmission by the Message VPN. Available since 2.13. | [optional] 
**DistributedCacheManagementEnabled** | Pointer to **bool** | Indicates whether managing of cache instances over the message bus is enabled in the Message VPN. | [optional] 
**DmrEnabled** | Pointer to **bool** | Indicates whether Dynamic Message Routing (DMR) is enabled for the Message VPN. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Message VPN is enabled. | [optional] 
**EventConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventEgressFlowCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventEgressMsgRateThreshold** | Pointer to [**EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] 
**EventEndpointCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventIngressFlowCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventIngressMsgRateThreshold** | Pointer to [**EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] 
**EventLargeMsgThreshold** | Pointer to **int64** | Exceeding this message size in kilobytes (KB) triggers a corresponding Event in the Message VPN. | [optional] 
**EventLogTag** | Pointer to **string** | The value of the prefix applied to all published Events in the Message VPN. | [optional] 
**EventMsgSpoolUsageThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventPublishClientEnabled** | Pointer to **bool** | Indicates whether client Events are published in the Message VPN. | [optional] 
**EventPublishMsgVpnEnabled** | Pointer to **bool** | Indicates whether Message VPN Events are published in the Message VPN. | [optional] 
**EventPublishSubscriptionMode** | Pointer to **string** | The mode of subscription Events published in the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;off\&quot; - Disable client level event message publishing. \&quot;on-with-format-v1\&quot; - Enable client level event message publishing with format v1. \&quot;on-with-no-unsubscribe-events-on-disconnect-format-v1\&quot; - As \&quot;on-with-format-v1\&quot;, but unsubscribe events are not generated when a client disconnects. Unsubscribe events are still raised when a client explicitly unsubscribes from its subscriptions. \&quot;on-with-format-v2\&quot; - Enable client level event message publishing with format v2. \&quot;on-with-no-unsubscribe-events-on-disconnect-format-v2\&quot; - As \&quot;on-with-format-v2\&quot;, but unsubscribe events are not generated when a client disconnects. Unsubscribe events are still raised when a client explicitly unsubscribes from its subscriptions. &lt;/pre&gt;  | [optional] 
**EventPublishTopicFormatMqttEnabled** | Pointer to **bool** | Indicates whether Message VPN Events are published in the MQTT format. | [optional] 
**EventPublishTopicFormatSmfEnabled** | Pointer to **bool** | Indicates whether Message VPN Events are published in the SMF format. | [optional] 
**EventServiceAmqpConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventServiceMqttConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventServiceRestIncomingConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventServiceSmfConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventServiceWebConnectionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventSubscriptionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventTransactedSessionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**EventTransactionCountThreshold** | Pointer to [**EventThreshold**](EventThreshold.md) |  | [optional] 
**ExportSubscriptionsEnabled** | Pointer to **bool** | Indicates whether exports of subscriptions to other routers in the network over neighbour links is enabled in the Message VPN. | [optional] 
**FailureReason** | Pointer to **string** | The reason for the Message VPN failure. | [optional] 
**JndiEnabled** | Pointer to **bool** | Indicates whether the JNDI access for clients is enabled in the Message VPN. | [optional] 
**LoginRxMsgCount** | Pointer to **int64** | The number of login request messages received by the Message VPN. Available since 2.13. | [optional] 
**LoginTxMsgCount** | Pointer to **int64** | The number of login response messages transmitted by the Message VPN. Available since 2.13. | [optional] 
**MaxConnectionCount** | Pointer to **int64** | The maximum number of client connections to the Message VPN. | [optional] 
**MaxEffectiveEndpointCount** | Pointer to **int32** | The effective maximum number of Queues and Topic Endpoints allowed in the Message VPN. | [optional] 
**MaxEffectiveRxFlowCount** | Pointer to **int32** | The effective maximum number of receive flows allowed in the Message VPN. | [optional] 
**MaxEffectiveSubscriptionCount** | Pointer to **int64** | The effective maximum number of subscriptions allowed in the Message VPN. | [optional] 
**MaxEffectiveTransactedSessionCount** | Pointer to **int32** | The effective maximum number of transacted sessions allowed in the Message VPN. | [optional] 
**MaxEffectiveTransactionCount** | Pointer to **int32** | The effective maximum number of transactions allowed in the Message VPN. | [optional] 
**MaxEffectiveTxFlowCount** | Pointer to **int32** | The effective maximum number of transmit flows allowed in the Message VPN. | [optional] 
**MaxEgressFlowCount** | Pointer to **int64** | The maximum number of transmit flows that can be created in the Message VPN. | [optional] 
**MaxEndpointCount** | Pointer to **int64** | The maximum number of Queues and Topic Endpoints that can be created in the Message VPN. | [optional] 
**MaxIngressFlowCount** | Pointer to **int64** | The maximum number of receive flows that can be created in the Message VPN. | [optional] 
**MaxMsgSpoolUsage** | Pointer to **int64** | The maximum message spool usage by the Message VPN, in megabytes. | [optional] 
**MaxSubscriptionCount** | Pointer to **int64** | The maximum number of local client subscriptions that can be added to the Message VPN. This limit is not enforced when a subscription is added using a management interface, such as CLI or SEMP. | [optional] 
**MaxTransactedSessionCount** | Pointer to **int64** | The maximum number of transacted sessions that can be created in the Message VPN. | [optional] 
**MaxTransactionCount** | Pointer to **int64** | The maximum number of transactions that can be created in the Message VPN. | [optional] 
**MqttRetainMaxMemory** | Pointer to **int32** | The maximum total memory usage of the MQTT Retain feature for this Message VPN, in MB. If the maximum memory is reached, any arriving retain messages that require more memory are discarded. A value of -1 indicates that the memory is bounded only by the global max memory limit. A value of 0 prevents MQTT Retain from becoming operational. | [optional] 
**MsgReplayActiveCount** | Pointer to **int32** | The number of message replays that are currently active in the Message VPN. | [optional] 
**MsgReplayFailedCount** | Pointer to **int32** | The number of message replays that are currently failed in the Message VPN. | [optional] 
**MsgReplayInitializingCount** | Pointer to **int32** | The number of message replays that are currently initializing in the Message VPN. | [optional] 
**MsgReplayPendingCompleteCount** | Pointer to **int32** | The number of message replays that are pending complete in the Message VPN. | [optional] 
**MsgSpoolMsgCount** | Pointer to **int64** | The current number of messages spooled (persisted in the Message Spool) in the Message VPN. Available since 2.14. | [optional] 
**MsgSpoolRxMsgCount** | Pointer to **int64** | The number of guaranteed messages received by the Message VPN. Available since 2.13. | [optional] 
**MsgSpoolTxMsgCount** | Pointer to **int64** | The number of guaranteed messages transmitted by the Message VPN. One message to multiple clients is counted as one message. Available since 2.13. | [optional] 
**MsgSpoolUsage** | Pointer to **int64** | The current message spool usage by the Message VPN, in bytes (B). | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Rate** | Pointer to [**MsgVpnRate**](MsgVpnRate.md) |  | [optional] 
**ReplicationAckPropagationIntervalMsgCount** | Pointer to **int64** | The acknowledgement (ACK) propagation interval for the replication Bridge, in number of replicated messages. Available since 2.12. | [optional] 
**ReplicationActiveAckPropTxMsgCount** | Pointer to **int64** | The number of acknowledgement messages propagated to the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationActiveAsyncQueuedMsgCount** | Pointer to **int64** | The number of async messages queued to the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationActiveLocallyConsumedMsgCount** | Pointer to **int64** | The number of messages consumed in the replication active local Message VPN. Available since 2.12. | [optional] 
**ReplicationActiveMateFlowCongestedPeakTime** | Pointer to **int32** | The peak amount of time in seconds the message flow has been congested to the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationActiveMateFlowNotCongestedPeakTime** | Pointer to **int32** | The peak amount of time in seconds the message flow has not been congested to the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationActivePromotedQueuedMsgCount** | Pointer to **int64** | The number of promoted messages queued to the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationActiveReconcileRequestRxMsgCount** | Pointer to **int64** | The number of reconcile request messages received from the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationActiveSyncEligiblePeakTime** | Pointer to **int32** | The peak amount of time in seconds sync replication has been eligible to the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationActiveSyncIneligiblePeakTime** | Pointer to **int32** | The peak amount of time in seconds sync replication has been ineligible to the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationActiveSyncQueuedAsAsyncMsgCount** | Pointer to **int64** | The number of sync messages queued as async to the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationActiveSyncQueuedMsgCount** | Pointer to **int64** | The number of sync messages queued to the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationActiveTransitionToSyncIneligibleCount** | Pointer to **int64** | The number of sync replication ineligible transitions to the replication standby remote Message VPN. Available since 2.12. | [optional] 
**ReplicationBridgeAuthenticationBasicClientUsername** | Pointer to **string** | The Client Username the replication Bridge uses to login to the remote Message VPN. Available since 2.12. | [optional] 
**ReplicationBridgeAuthenticationScheme** | Pointer to **string** | The authentication scheme for the replication Bridge in the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;basic\&quot; - Basic Authentication Scheme (via username and password). \&quot;client-certificate\&quot; - Client Certificate Authentication Scheme (via certificate file or content). &lt;/pre&gt;  Available since 2.12. | [optional] 
**ReplicationBridgeBoundToQueue** | Pointer to **bool** | Indicates whether the local replication Bridge is bound to the Queue in the remote Message VPN. Available since 2.12. | [optional] 
**ReplicationBridgeCompressedDataEnabled** | Pointer to **bool** | Indicates whether compression is used for the replication Bridge. Available since 2.12. | [optional] 
**ReplicationBridgeEgressFlowWindowSize** | Pointer to **int64** | The size of the window used for guaranteed messages published to the replication Bridge, in messages. Available since 2.12. | [optional] 
**ReplicationBridgeName** | Pointer to **string** | The name of the local replication Bridge in the Message VPN. Available since 2.12. | [optional] 
**ReplicationBridgeRetryDelay** | Pointer to **int64** | The number of seconds that must pass before retrying the replication Bridge connection. Available since 2.12. | [optional] 
**ReplicationBridgeTlsEnabled** | Pointer to **bool** | Indicates whether encryption (TLS) is enabled for the replication Bridge connection. Available since 2.12. | [optional] 
**ReplicationBridgeUnidirectionalClientProfileName** | Pointer to **string** | The Client Profile for the unidirectional replication Bridge in the Message VPN. It is used only for the TCP parameters. Available since 2.12. | [optional] 
**ReplicationBridgeUp** | Pointer to **bool** | Indicates whether the local replication Bridge is operationally up in the Message VPN. Available since 2.12. | [optional] 
**ReplicationEnabled** | Pointer to **bool** | Indicates whether replication is enabled for the Message VPN. Available since 2.12. | [optional] 
**ReplicationQueueBound** | Pointer to **bool** | Indicates whether the remote replication Bridge is bound to the Queue in the Message VPN. Available since 2.12. | [optional] 
**ReplicationQueueMaxMsgSpoolUsage** | Pointer to **int64** | The maximum message spool usage by the replication Bridge local Queue (quota), in megabytes. Available since 2.12. | [optional] 
**ReplicationQueueRejectMsgToSenderOnDiscardEnabled** | Pointer to **bool** | Indicates whether messages discarded on this replication Bridge Queue are rejected back to the sender. Available since 2.12. | [optional] 
**ReplicationRejectMsgWhenSyncIneligibleEnabled** | Pointer to **bool** | Indicates whether guaranteed messages published to synchronously replicated Topics are rejected back to the sender when synchronous replication becomes ineligible. Available since 2.12. | [optional] 
**ReplicationRemoteBridgeName** | Pointer to **string** | The name of the remote replication Bridge in the Message VPN. Available since 2.12. | [optional] 
**ReplicationRemoteBridgeUp** | Pointer to **bool** | Indicates whether the remote replication Bridge is operationally up in the Message VPN. Available since 2.12. | [optional] 
**ReplicationRole** | Pointer to **string** | The replication role for the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;active\&quot; - Assume the Active role in replication for the Message VPN. \&quot;standby\&quot; - Assume the Standby role in replication for the Message VPN. &lt;/pre&gt;  Available since 2.12. | [optional] 
**ReplicationStandbyAckPropOutOfSeqRxMsgCount** | Pointer to **int64** | The number of acknowledgement messages received out of sequence from the replication active remote Message VPN. Available since 2.12. | [optional] 
**ReplicationStandbyAckPropRxMsgCount** | Pointer to **int64** | The number of acknowledgement messages received from the replication active remote Message VPN. Available since 2.12. | [optional] 
**ReplicationStandbyReconcileRequestTxMsgCount** | Pointer to **int64** | The number of reconcile request messages transmitted to the replication active remote Message VPN. Available since 2.12. | [optional] 
**ReplicationStandbyRxMsgCount** | Pointer to **int64** | The number of messages received from the replication active remote Message VPN. Available since 2.12. | [optional] 
**ReplicationStandbyTransactionRequestCount** | Pointer to **int64** | The number of transaction requests received from the replication active remote Message VPN. Available since 2.12. | [optional] 
**ReplicationStandbyTransactionRequestFailureCount** | Pointer to **int64** | The number of transaction requests received from the replication active remote Message VPN that failed. Available since 2.12. | [optional] 
**ReplicationStandbyTransactionRequestSuccessCount** | Pointer to **int64** | The number of transaction requests received from the replication active remote Message VPN that succeeded. Available since 2.12. | [optional] 
**ReplicationSyncEligible** | Pointer to **bool** | Indicates whether sync replication is eligible in the Message VPN. Available since 2.12. | [optional] 
**ReplicationTransactionMode** | Pointer to **string** | Indicates whether synchronous or asynchronous replication mode is used for all transactions within the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;sync\&quot; - Messages are acknowledged when replicated (spooled remotely). \&quot;async\&quot; - Messages are acknowledged when pending replication (spooled locally). &lt;/pre&gt;  Available since 2.12. | [optional] 
**RestTlsServerCertEnforceTrustedCommonNameEnabled** | Pointer to **bool** | Indicates whether the Common Name (CN) in the server certificate from the remote REST Consumer is validated. Deprecated since 2.17. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**RestTlsServerCertMaxChainDepth** | Pointer to **int64** | The maximum depth for a REST Consumer server certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. | [optional] 
**RestTlsServerCertValidateDateEnabled** | Pointer to **bool** | Indicates whether the \&quot;Not Before\&quot; and \&quot;Not After\&quot; validity dates in the REST Consumer server certificate are checked. | [optional] 
**RestTlsServerCertValidateNameEnabled** | Pointer to **bool** | Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the remote REST Consumer. If enabled, the name used to connect to the remote REST Consumer is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. Available since 2.17. | [optional] 
**RxByteCount** | Pointer to **int64** | The amount of messages received from clients by the Message VPN, in bytes (B). Available since 2.12. | [optional] 
**RxByteRate** | Pointer to **int64** | The current message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13. | [optional] 
**RxCompressedByteCount** | Pointer to **int64** | The amount of compressed messages received by the Message VPN, in bytes (B). Available since 2.12. | [optional] 
**RxCompressedByteRate** | Pointer to **int64** | The current compressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12. | [optional] 
**RxCompressionRatio** | Pointer to **string** | The compression ratio for messages received by the message VPN. Available since 2.12. | [optional] 
**RxMsgCount** | Pointer to **int64** | The number of messages received from clients by the Message VPN. Available since 2.12. | [optional] 
**RxMsgRate** | Pointer to **int64** | The current message rate received by the Message VPN, in messages per second (msg/sec). Available since 2.13. | [optional] 
**RxUncompressedByteCount** | Pointer to **int64** | The amount of uncompressed messages received by the Message VPN, in bytes (B). Available since 2.12. | [optional] 
**RxUncompressedByteRate** | Pointer to **int64** | The current uncompressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12. | [optional] 
**SempOverMsgBusAdminClientEnabled** | Pointer to **bool** | Indicates whether the \&quot;admin\&quot; level \&quot;client\&quot; commands are enabled for SEMP over the message bus in the Message VPN. | [optional] 
**SempOverMsgBusAdminDistributedCacheEnabled** | Pointer to **bool** | Indicates whether the \&quot;admin\&quot; level \&quot;Distributed Cache\&quot; commands are enabled for SEMP over the message bus in the Message VPN. | [optional] 
**SempOverMsgBusAdminEnabled** | Pointer to **bool** | Indicates whether the \&quot;admin\&quot; level commands are enabled for SEMP over the message bus in the Message VPN. | [optional] 
**SempOverMsgBusEnabled** | Pointer to **bool** | Indicates whether SEMP over the message bus is enabled in the Message VPN. | [optional] 
**SempOverMsgBusShowEnabled** | Pointer to **bool** | Indicates whether the \&quot;show\&quot; level commands are enabled for SEMP over the message bus in the Message VPN. | [optional] 
**ServiceAmqpMaxConnectionCount** | Pointer to **int64** | The maximum number of AMQP client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform. | [optional] 
**ServiceAmqpPlainTextCompressed** | Pointer to **bool** | Indicates whether the AMQP Service is compressed in the Message VPN. | [optional] 
**ServiceAmqpPlainTextEnabled** | Pointer to **bool** | Indicates whether the AMQP Service is enabled in the Message VPN. | [optional] 
**ServiceAmqpPlainTextFailureReason** | Pointer to **string** | The reason for the AMQP Service failure in the Message VPN. | [optional] 
**ServiceAmqpPlainTextListenPort** | Pointer to **int64** | The port number for plain-text AMQP clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. | [optional] 
**ServiceAmqpPlainTextUp** | Pointer to **bool** | Indicates whether the AMQP Service is operationally up in the Message VPN. | [optional] 
**ServiceAmqpTlsCompressed** | Pointer to **bool** | Indicates whether the TLS related AMQP Service is compressed in the Message VPN. | [optional] 
**ServiceAmqpTlsEnabled** | Pointer to **bool** | Indicates whether encryption (TLS) is enabled for AMQP clients in the Message VPN. | [optional] 
**ServiceAmqpTlsFailureReason** | Pointer to **string** | The reason for the TLS related AMQP Service failure in the Message VPN. | [optional] 
**ServiceAmqpTlsListenPort** | Pointer to **int64** | The port number for AMQP clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. | [optional] 
**ServiceAmqpTlsUp** | Pointer to **bool** | Indicates whether the TLS related AMQP Service is operationally up in the Message VPN. | [optional] 
**ServiceMqttAuthenticationClientCertRequest** | Pointer to **string** | Determines when to request a client certificate from an incoming MQTT client connecting via a TLS port. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;never\&quot; - Never ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;when-enabled-in-message-vpn\&quot; - Only ask for a client-certificate if client certificate authentication is enabled under \&quot;message-vpn &gt;  authentication &gt; client-certificate &gt; shutdown\&quot;. &lt;/pre&gt;  Available since 2.21. | [optional] 
**ServiceMqttMaxConnectionCount** | Pointer to **int64** | The maximum number of MQTT client connections that can be simultaneously connected to the Message VPN. | [optional] 
**ServiceMqttPlainTextCompressed** | Pointer to **bool** | Indicates whether the MQTT Service is compressed in the Message VPN. | [optional] 
**ServiceMqttPlainTextEnabled** | Pointer to **bool** | Indicates whether the MQTT Service is enabled in the Message VPN. | [optional] 
**ServiceMqttPlainTextFailureReason** | Pointer to **string** | The reason for the MQTT Service failure in the Message VPN. | [optional] 
**ServiceMqttPlainTextListenPort** | Pointer to **int64** | The port number for plain-text MQTT clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. | [optional] 
**ServiceMqttPlainTextUp** | Pointer to **bool** | Indicates whether the MQTT Service is operationally up in the Message VPN. | [optional] 
**ServiceMqttTlsCompressed** | Pointer to **bool** | Indicates whether the TLS related MQTT Service is compressed in the Message VPN. | [optional] 
**ServiceMqttTlsEnabled** | Pointer to **bool** | Indicates whether encryption (TLS) is enabled for MQTT clients in the Message VPN. | [optional] 
**ServiceMqttTlsFailureReason** | Pointer to **string** | The reason for the TLS related MQTT Service failure in the Message VPN. | [optional] 
**ServiceMqttTlsListenPort** | Pointer to **int64** | The port number for MQTT clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. | [optional] 
**ServiceMqttTlsUp** | Pointer to **bool** | Indicates whether the TLS related MQTT Service is operationally up in the Message VPN. | [optional] 
**ServiceMqttTlsWebSocketCompressed** | Pointer to **bool** | Indicates whether the TLS related Web transport MQTT Service is compressed in the Message VPN. | [optional] 
**ServiceMqttTlsWebSocketEnabled** | Pointer to **bool** | Indicates whether encryption (TLS) is enabled for MQTT Web clients in the Message VPN. | [optional] 
**ServiceMqttTlsWebSocketFailureReason** | Pointer to **string** | The reason for the TLS related Web transport MQTT Service failure in the Message VPN. | [optional] 
**ServiceMqttTlsWebSocketListenPort** | Pointer to **int64** | The port number for MQTT clients that connect to the Message VPN using WebSocket over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. | [optional] 
**ServiceMqttTlsWebSocketUp** | Pointer to **bool** | Indicates whether the TLS related Web transport MQTT Service is operationally up in the Message VPN. | [optional] 
**ServiceMqttWebSocketCompressed** | Pointer to **bool** | Indicates whether the Web transport related MQTT Service is compressed in the Message VPN. | [optional] 
**ServiceMqttWebSocketEnabled** | Pointer to **bool** | Indicates whether the Web transport for the SMF Service is enabled in the Message VPN. | [optional] 
**ServiceMqttWebSocketFailureReason** | Pointer to **string** | The reason for the Web transport related MQTT Service failure in the Message VPN. | [optional] 
**ServiceMqttWebSocketListenPort** | Pointer to **int64** | The port number for plain-text MQTT clients that connect to the Message VPN using WebSocket. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. | [optional] 
**ServiceMqttWebSocketUp** | Pointer to **bool** | Indicates whether the Web transport related MQTT Service is operationally up in the Message VPN. | [optional] 
**ServiceRestIncomingAuthenticationClientCertRequest** | Pointer to **string** | Determines when to request a client certificate from an incoming REST Producer connecting via a TLS port. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;never\&quot; - Never ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;when-enabled-in-message-vpn\&quot; - Only ask for a client-certificate if client certificate authentication is enabled under \&quot;message-vpn &gt;  authentication &gt; client-certificate &gt; shutdown\&quot;. &lt;/pre&gt;  Available since 2.21. | [optional] 
**ServiceRestIncomingAuthorizationHeaderHandling** | Pointer to **string** | The handling of Authorization headers for incoming REST connections. The allowed values and their meaning are:  &lt;pre&gt; \&quot;drop\&quot; - Do not attach the Authorization header to the message as a user property. This configuration is most secure. \&quot;forward\&quot; - Forward the Authorization header, attaching it to the message as a user property in the same way as other headers. For best security, use the drop setting. \&quot;legacy\&quot; - If the Authorization header was used for authentication to the broker, do not attach it to the message. If the Authorization header was not used for authentication to the broker, attach it to the message as a user property in the same way as other headers. For best security, use the drop setting. &lt;/pre&gt;  Available since 2.19. | [optional] 
**ServiceRestIncomingMaxConnectionCount** | Pointer to **int64** | The maximum number of REST incoming client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform. | [optional] 
**ServiceRestIncomingPlainTextCompressed** | Pointer to **bool** | Indicates whether the incoming REST Service is compressed in the Message VPN. | [optional] 
**ServiceRestIncomingPlainTextEnabled** | Pointer to **bool** | Indicates whether the REST Service is enabled in the Message VPN for incoming clients. | [optional] 
**ServiceRestIncomingPlainTextFailureReason** | Pointer to **string** | The reason for the incoming REST Service failure in the Message VPN. | [optional] 
**ServiceRestIncomingPlainTextListenPort** | Pointer to **int64** | The port number for incoming plain-text REST clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. | [optional] 
**ServiceRestIncomingPlainTextUp** | Pointer to **bool** | Indicates whether the incoming REST Service is operationally up in the Message VPN. | [optional] 
**ServiceRestIncomingTlsCompressed** | Pointer to **bool** | Indicates whether the TLS related incoming REST Service is compressed in the Message VPN. | [optional] 
**ServiceRestIncomingTlsEnabled** | Pointer to **bool** | Indicates whether encryption (TLS) is enabled for incoming REST clients in the Message VPN. | [optional] 
**ServiceRestIncomingTlsFailureReason** | Pointer to **string** | The reason for the TLS related incoming REST Service failure in the Message VPN. | [optional] 
**ServiceRestIncomingTlsListenPort** | Pointer to **int64** | The port number for incoming REST clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled. | [optional] 
**ServiceRestIncomingTlsUp** | Pointer to **bool** | Indicates whether the TLS related incoming REST Service is operationally up in the Message VPN. | [optional] 
**ServiceRestMode** | Pointer to **string** | The REST service mode for incoming REST clients that connect to the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;gateway\&quot; - Act as a message gateway through which REST messages are propagated. \&quot;messaging\&quot; - Act as a message broker on which REST messages are queued. &lt;/pre&gt;  | [optional] 
**ServiceRestOutgoingMaxConnectionCount** | Pointer to **int64** | The maximum number of REST Consumer (outgoing) client connections that can be simultaneously connected to the Message VPN. | [optional] 
**ServiceSmfMaxConnectionCount** | Pointer to **int64** | The maximum number of SMF client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform. | [optional] 
**ServiceSmfPlainTextEnabled** | Pointer to **bool** | Indicates whether the SMF Service is enabled in the Message VPN. | [optional] 
**ServiceSmfPlainTextFailureReason** | Pointer to **string** | The reason for the SMF Service failure in the Message VPN. | [optional] 
**ServiceSmfPlainTextUp** | Pointer to **bool** | Indicates whether the SMF Service is operationally up in the Message VPN. | [optional] 
**ServiceSmfTlsEnabled** | Pointer to **bool** | Indicates whether encryption (TLS) is enabled for SMF clients in the Message VPN. | [optional] 
**ServiceSmfTlsFailureReason** | Pointer to **string** | The reason for the TLS related SMF Service failure in the Message VPN. | [optional] 
**ServiceSmfTlsUp** | Pointer to **bool** | Indicates whether the TLS related SMF Service is operationally up in the Message VPN. | [optional] 
**ServiceWebAuthenticationClientCertRequest** | Pointer to **string** | Determines when to request a client certificate from a Web Transport client connecting via a TLS port. The allowed values and their meaning are:  &lt;pre&gt; \&quot;always\&quot; - Always ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;never\&quot; - Never ask for a client certificate regardless of the \&quot;message-vpn &gt; authentication &gt; client-certificate &gt; shutdown\&quot; configuration. \&quot;when-enabled-in-message-vpn\&quot; - Only ask for a client-certificate if client certificate authentication is enabled under \&quot;message-vpn &gt;  authentication &gt; client-certificate &gt; shutdown\&quot;. &lt;/pre&gt;  Available since 2.21. | [optional] 
**ServiceWebMaxConnectionCount** | Pointer to **int64** | The maximum number of Web Transport client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform. | [optional] 
**ServiceWebPlainTextEnabled** | Pointer to **bool** | Indicates whether the Web transport for the SMF Service is enabled in the Message VPN. | [optional] 
**ServiceWebPlainTextFailureReason** | Pointer to **string** | The reason for the Web transport related SMF Service failure in the Message VPN. | [optional] 
**ServiceWebPlainTextUp** | Pointer to **bool** | Indicates whether the Web transport for the SMF Service is operationally up in the Message VPN. | [optional] 
**ServiceWebTlsEnabled** | Pointer to **bool** | Indicates whether TLS is enabled for SMF clients in the Message VPN that use the Web transport. | [optional] 
**ServiceWebTlsFailureReason** | Pointer to **string** | The reason for the TLS related Web transport SMF Service failure in the Message VPN. | [optional] 
**ServiceWebTlsUp** | Pointer to **bool** | Indicates whether the TLS related Web transport SMF Service is operationally up in the Message VPN. | [optional] 
**State** | Pointer to **string** | The operational state of the local Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;up\&quot; - The Message VPN is operationally up. \&quot;down\&quot; - The Message VPN is operationally down. \&quot;standby\&quot; - The Message VPN is operationally replication standby. &lt;/pre&gt;  | [optional] 
**SubscriptionExportProgress** | Pointer to **int64** | The progress of the subscription export task, in percent complete. | [optional] 
**SystemManager** | Pointer to **bool** | Indicates whether the Message VPN is the system manager for handling system level SEMP get requests and system level event publishing. | [optional] 
**TlsAllowDowngradeToPlainTextEnabled** | Pointer to **bool** | Indicates whether SMF clients connected to the Message VPN are allowed to downgrade their connections from TLS to plain text. | [optional] 
**TlsAverageRxByteRate** | Pointer to **int64** | The one minute average of the TLS message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13. | [optional] 
**TlsAverageTxByteRate** | Pointer to **int64** | The one minute average of the TLS message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13. | [optional] 
**TlsRxByteCount** | Pointer to **int64** | The amount of TLS messages received by the Message VPN, in bytes (B). Available since 2.13. | [optional] 
**TlsRxByteRate** | Pointer to **int64** | The current TLS message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13. | [optional] 
**TlsTxByteCount** | Pointer to **int64** | The amount of TLS messages transmitted by the Message VPN, in bytes (B). Available since 2.13. | [optional] 
**TlsTxByteRate** | Pointer to **int64** | The current TLS message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13. | [optional] 
**TxByteCount** | Pointer to **int64** | The amount of messages transmitted to clients by the Message VPN, in bytes (B). Available since 2.12. | [optional] 
**TxByteRate** | Pointer to **int64** | The current message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13. | [optional] 
**TxCompressedByteCount** | Pointer to **int64** | The amount of compressed messages transmitted by the Message VPN, in bytes (B). Available since 2.12. | [optional] 
**TxCompressedByteRate** | Pointer to **int64** | The current compressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12. | [optional] 
**TxCompressionRatio** | Pointer to **string** | The compression ratio for messages transmitted by the message VPN. Available since 2.12. | [optional] 
**TxMsgCount** | Pointer to **int64** | The number of messages transmitted to clients by the Message VPN. Available since 2.12. | [optional] 
**TxMsgRate** | Pointer to **int64** | The current message rate transmitted by the Message VPN, in messages per second (msg/sec). Available since 2.13. | [optional] 
**TxUncompressedByteCount** | Pointer to **int64** | The amount of uncompressed messages transmitted by the Message VPN, in bytes (B). Available since 2.12. | [optional] 
**TxUncompressedByteRate** | Pointer to **int64** | The current uncompressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12. | [optional] 

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

### GetAverageRxByteRate

`func (o *MsgVpn) GetAverageRxByteRate() int64`

GetAverageRxByteRate returns the AverageRxByteRate field if non-nil, zero value otherwise.

### GetAverageRxByteRateOk

`func (o *MsgVpn) GetAverageRxByteRateOk() (*int64, bool)`

GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxByteRate

`func (o *MsgVpn) SetAverageRxByteRate(v int64)`

SetAverageRxByteRate sets AverageRxByteRate field to given value.

### HasAverageRxByteRate

`func (o *MsgVpn) HasAverageRxByteRate() bool`

HasAverageRxByteRate returns a boolean if a field has been set.

### GetAverageRxCompressedByteRate

`func (o *MsgVpn) GetAverageRxCompressedByteRate() int64`

GetAverageRxCompressedByteRate returns the AverageRxCompressedByteRate field if non-nil, zero value otherwise.

### GetAverageRxCompressedByteRateOk

`func (o *MsgVpn) GetAverageRxCompressedByteRateOk() (*int64, bool)`

GetAverageRxCompressedByteRateOk returns a tuple with the AverageRxCompressedByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxCompressedByteRate

`func (o *MsgVpn) SetAverageRxCompressedByteRate(v int64)`

SetAverageRxCompressedByteRate sets AverageRxCompressedByteRate field to given value.

### HasAverageRxCompressedByteRate

`func (o *MsgVpn) HasAverageRxCompressedByteRate() bool`

HasAverageRxCompressedByteRate returns a boolean if a field has been set.

### GetAverageRxMsgRate

`func (o *MsgVpn) GetAverageRxMsgRate() int64`

GetAverageRxMsgRate returns the AverageRxMsgRate field if non-nil, zero value otherwise.

### GetAverageRxMsgRateOk

`func (o *MsgVpn) GetAverageRxMsgRateOk() (*int64, bool)`

GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxMsgRate

`func (o *MsgVpn) SetAverageRxMsgRate(v int64)`

SetAverageRxMsgRate sets AverageRxMsgRate field to given value.

### HasAverageRxMsgRate

`func (o *MsgVpn) HasAverageRxMsgRate() bool`

HasAverageRxMsgRate returns a boolean if a field has been set.

### GetAverageRxUncompressedByteRate

`func (o *MsgVpn) GetAverageRxUncompressedByteRate() int64`

GetAverageRxUncompressedByteRate returns the AverageRxUncompressedByteRate field if non-nil, zero value otherwise.

### GetAverageRxUncompressedByteRateOk

`func (o *MsgVpn) GetAverageRxUncompressedByteRateOk() (*int64, bool)`

GetAverageRxUncompressedByteRateOk returns a tuple with the AverageRxUncompressedByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxUncompressedByteRate

`func (o *MsgVpn) SetAverageRxUncompressedByteRate(v int64)`

SetAverageRxUncompressedByteRate sets AverageRxUncompressedByteRate field to given value.

### HasAverageRxUncompressedByteRate

`func (o *MsgVpn) HasAverageRxUncompressedByteRate() bool`

HasAverageRxUncompressedByteRate returns a boolean if a field has been set.

### GetAverageTxByteRate

`func (o *MsgVpn) GetAverageTxByteRate() int64`

GetAverageTxByteRate returns the AverageTxByteRate field if non-nil, zero value otherwise.

### GetAverageTxByteRateOk

`func (o *MsgVpn) GetAverageTxByteRateOk() (*int64, bool)`

GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxByteRate

`func (o *MsgVpn) SetAverageTxByteRate(v int64)`

SetAverageTxByteRate sets AverageTxByteRate field to given value.

### HasAverageTxByteRate

`func (o *MsgVpn) HasAverageTxByteRate() bool`

HasAverageTxByteRate returns a boolean if a field has been set.

### GetAverageTxCompressedByteRate

`func (o *MsgVpn) GetAverageTxCompressedByteRate() int64`

GetAverageTxCompressedByteRate returns the AverageTxCompressedByteRate field if non-nil, zero value otherwise.

### GetAverageTxCompressedByteRateOk

`func (o *MsgVpn) GetAverageTxCompressedByteRateOk() (*int64, bool)`

GetAverageTxCompressedByteRateOk returns a tuple with the AverageTxCompressedByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxCompressedByteRate

`func (o *MsgVpn) SetAverageTxCompressedByteRate(v int64)`

SetAverageTxCompressedByteRate sets AverageTxCompressedByteRate field to given value.

### HasAverageTxCompressedByteRate

`func (o *MsgVpn) HasAverageTxCompressedByteRate() bool`

HasAverageTxCompressedByteRate returns a boolean if a field has been set.

### GetAverageTxMsgRate

`func (o *MsgVpn) GetAverageTxMsgRate() int64`

GetAverageTxMsgRate returns the AverageTxMsgRate field if non-nil, zero value otherwise.

### GetAverageTxMsgRateOk

`func (o *MsgVpn) GetAverageTxMsgRateOk() (*int64, bool)`

GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxMsgRate

`func (o *MsgVpn) SetAverageTxMsgRate(v int64)`

SetAverageTxMsgRate sets AverageTxMsgRate field to given value.

### HasAverageTxMsgRate

`func (o *MsgVpn) HasAverageTxMsgRate() bool`

HasAverageTxMsgRate returns a boolean if a field has been set.

### GetAverageTxUncompressedByteRate

`func (o *MsgVpn) GetAverageTxUncompressedByteRate() int64`

GetAverageTxUncompressedByteRate returns the AverageTxUncompressedByteRate field if non-nil, zero value otherwise.

### GetAverageTxUncompressedByteRateOk

`func (o *MsgVpn) GetAverageTxUncompressedByteRateOk() (*int64, bool)`

GetAverageTxUncompressedByteRateOk returns a tuple with the AverageTxUncompressedByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxUncompressedByteRate

`func (o *MsgVpn) SetAverageTxUncompressedByteRate(v int64)`

SetAverageTxUncompressedByteRate sets AverageTxUncompressedByteRate field to given value.

### HasAverageTxUncompressedByteRate

`func (o *MsgVpn) HasAverageTxUncompressedByteRate() bool`

HasAverageTxUncompressedByteRate returns a boolean if a field has been set.

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

### GetConfigSyncLocalKey

`func (o *MsgVpn) GetConfigSyncLocalKey() string`

GetConfigSyncLocalKey returns the ConfigSyncLocalKey field if non-nil, zero value otherwise.

### GetConfigSyncLocalKeyOk

`func (o *MsgVpn) GetConfigSyncLocalKeyOk() (*string, bool)`

GetConfigSyncLocalKeyOk returns a tuple with the ConfigSyncLocalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSyncLocalKey

`func (o *MsgVpn) SetConfigSyncLocalKey(v string)`

SetConfigSyncLocalKey sets ConfigSyncLocalKey field to given value.

### HasConfigSyncLocalKey

`func (o *MsgVpn) HasConfigSyncLocalKey() bool`

HasConfigSyncLocalKey returns a boolean if a field has been set.

### GetConfigSyncLocalLastResult

`func (o *MsgVpn) GetConfigSyncLocalLastResult() string`

GetConfigSyncLocalLastResult returns the ConfigSyncLocalLastResult field if non-nil, zero value otherwise.

### GetConfigSyncLocalLastResultOk

`func (o *MsgVpn) GetConfigSyncLocalLastResultOk() (*string, bool)`

GetConfigSyncLocalLastResultOk returns a tuple with the ConfigSyncLocalLastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSyncLocalLastResult

`func (o *MsgVpn) SetConfigSyncLocalLastResult(v string)`

SetConfigSyncLocalLastResult sets ConfigSyncLocalLastResult field to given value.

### HasConfigSyncLocalLastResult

`func (o *MsgVpn) HasConfigSyncLocalLastResult() bool`

HasConfigSyncLocalLastResult returns a boolean if a field has been set.

### GetConfigSyncLocalRole

`func (o *MsgVpn) GetConfigSyncLocalRole() string`

GetConfigSyncLocalRole returns the ConfigSyncLocalRole field if non-nil, zero value otherwise.

### GetConfigSyncLocalRoleOk

`func (o *MsgVpn) GetConfigSyncLocalRoleOk() (*string, bool)`

GetConfigSyncLocalRoleOk returns a tuple with the ConfigSyncLocalRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSyncLocalRole

`func (o *MsgVpn) SetConfigSyncLocalRole(v string)`

SetConfigSyncLocalRole sets ConfigSyncLocalRole field to given value.

### HasConfigSyncLocalRole

`func (o *MsgVpn) HasConfigSyncLocalRole() bool`

HasConfigSyncLocalRole returns a boolean if a field has been set.

### GetConfigSyncLocalState

`func (o *MsgVpn) GetConfigSyncLocalState() string`

GetConfigSyncLocalState returns the ConfigSyncLocalState field if non-nil, zero value otherwise.

### GetConfigSyncLocalStateOk

`func (o *MsgVpn) GetConfigSyncLocalStateOk() (*string, bool)`

GetConfigSyncLocalStateOk returns a tuple with the ConfigSyncLocalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSyncLocalState

`func (o *MsgVpn) SetConfigSyncLocalState(v string)`

SetConfigSyncLocalState sets ConfigSyncLocalState field to given value.

### HasConfigSyncLocalState

`func (o *MsgVpn) HasConfigSyncLocalState() bool`

HasConfigSyncLocalState returns a boolean if a field has been set.

### GetConfigSyncLocalTimeInState

`func (o *MsgVpn) GetConfigSyncLocalTimeInState() int32`

GetConfigSyncLocalTimeInState returns the ConfigSyncLocalTimeInState field if non-nil, zero value otherwise.

### GetConfigSyncLocalTimeInStateOk

`func (o *MsgVpn) GetConfigSyncLocalTimeInStateOk() (*int32, bool)`

GetConfigSyncLocalTimeInStateOk returns a tuple with the ConfigSyncLocalTimeInState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSyncLocalTimeInState

`func (o *MsgVpn) SetConfigSyncLocalTimeInState(v int32)`

SetConfigSyncLocalTimeInState sets ConfigSyncLocalTimeInState field to given value.

### HasConfigSyncLocalTimeInState

`func (o *MsgVpn) HasConfigSyncLocalTimeInState() bool`

HasConfigSyncLocalTimeInState returns a boolean if a field has been set.

### GetControlRxByteCount

`func (o *MsgVpn) GetControlRxByteCount() int64`

GetControlRxByteCount returns the ControlRxByteCount field if non-nil, zero value otherwise.

### GetControlRxByteCountOk

`func (o *MsgVpn) GetControlRxByteCountOk() (*int64, bool)`

GetControlRxByteCountOk returns a tuple with the ControlRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRxByteCount

`func (o *MsgVpn) SetControlRxByteCount(v int64)`

SetControlRxByteCount sets ControlRxByteCount field to given value.

### HasControlRxByteCount

`func (o *MsgVpn) HasControlRxByteCount() bool`

HasControlRxByteCount returns a boolean if a field has been set.

### GetControlRxMsgCount

`func (o *MsgVpn) GetControlRxMsgCount() int64`

GetControlRxMsgCount returns the ControlRxMsgCount field if non-nil, zero value otherwise.

### GetControlRxMsgCountOk

`func (o *MsgVpn) GetControlRxMsgCountOk() (*int64, bool)`

GetControlRxMsgCountOk returns a tuple with the ControlRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRxMsgCount

`func (o *MsgVpn) SetControlRxMsgCount(v int64)`

SetControlRxMsgCount sets ControlRxMsgCount field to given value.

### HasControlRxMsgCount

`func (o *MsgVpn) HasControlRxMsgCount() bool`

HasControlRxMsgCount returns a boolean if a field has been set.

### GetControlTxByteCount

`func (o *MsgVpn) GetControlTxByteCount() int64`

GetControlTxByteCount returns the ControlTxByteCount field if non-nil, zero value otherwise.

### GetControlTxByteCountOk

`func (o *MsgVpn) GetControlTxByteCountOk() (*int64, bool)`

GetControlTxByteCountOk returns a tuple with the ControlTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTxByteCount

`func (o *MsgVpn) SetControlTxByteCount(v int64)`

SetControlTxByteCount sets ControlTxByteCount field to given value.

### HasControlTxByteCount

`func (o *MsgVpn) HasControlTxByteCount() bool`

HasControlTxByteCount returns a boolean if a field has been set.

### GetControlTxMsgCount

`func (o *MsgVpn) GetControlTxMsgCount() int64`

GetControlTxMsgCount returns the ControlTxMsgCount field if non-nil, zero value otherwise.

### GetControlTxMsgCountOk

`func (o *MsgVpn) GetControlTxMsgCountOk() (*int64, bool)`

GetControlTxMsgCountOk returns a tuple with the ControlTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTxMsgCount

`func (o *MsgVpn) SetControlTxMsgCount(v int64)`

SetControlTxMsgCount sets ControlTxMsgCount field to given value.

### HasControlTxMsgCount

`func (o *MsgVpn) HasControlTxMsgCount() bool`

HasControlTxMsgCount returns a boolean if a field has been set.

### GetCounter

`func (o *MsgVpn) GetCounter() MsgVpnCounter`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *MsgVpn) GetCounterOk() (*MsgVpnCounter, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *MsgVpn) SetCounter(v MsgVpnCounter)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *MsgVpn) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetDataRxByteCount

`func (o *MsgVpn) GetDataRxByteCount() int64`

GetDataRxByteCount returns the DataRxByteCount field if non-nil, zero value otherwise.

### GetDataRxByteCountOk

`func (o *MsgVpn) GetDataRxByteCountOk() (*int64, bool)`

GetDataRxByteCountOk returns a tuple with the DataRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxByteCount

`func (o *MsgVpn) SetDataRxByteCount(v int64)`

SetDataRxByteCount sets DataRxByteCount field to given value.

### HasDataRxByteCount

`func (o *MsgVpn) HasDataRxByteCount() bool`

HasDataRxByteCount returns a boolean if a field has been set.

### GetDataRxMsgCount

`func (o *MsgVpn) GetDataRxMsgCount() int64`

GetDataRxMsgCount returns the DataRxMsgCount field if non-nil, zero value otherwise.

### GetDataRxMsgCountOk

`func (o *MsgVpn) GetDataRxMsgCountOk() (*int64, bool)`

GetDataRxMsgCountOk returns a tuple with the DataRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxMsgCount

`func (o *MsgVpn) SetDataRxMsgCount(v int64)`

SetDataRxMsgCount sets DataRxMsgCount field to given value.

### HasDataRxMsgCount

`func (o *MsgVpn) HasDataRxMsgCount() bool`

HasDataRxMsgCount returns a boolean if a field has been set.

### GetDataTxByteCount

`func (o *MsgVpn) GetDataTxByteCount() int64`

GetDataTxByteCount returns the DataTxByteCount field if non-nil, zero value otherwise.

### GetDataTxByteCountOk

`func (o *MsgVpn) GetDataTxByteCountOk() (*int64, bool)`

GetDataTxByteCountOk returns a tuple with the DataTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxByteCount

`func (o *MsgVpn) SetDataTxByteCount(v int64)`

SetDataTxByteCount sets DataTxByteCount field to given value.

### HasDataTxByteCount

`func (o *MsgVpn) HasDataTxByteCount() bool`

HasDataTxByteCount returns a boolean if a field has been set.

### GetDataTxMsgCount

`func (o *MsgVpn) GetDataTxMsgCount() int64`

GetDataTxMsgCount returns the DataTxMsgCount field if non-nil, zero value otherwise.

### GetDataTxMsgCountOk

`func (o *MsgVpn) GetDataTxMsgCountOk() (*int64, bool)`

GetDataTxMsgCountOk returns a tuple with the DataTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxMsgCount

`func (o *MsgVpn) SetDataTxMsgCount(v int64)`

SetDataTxMsgCount sets DataTxMsgCount field to given value.

### HasDataTxMsgCount

`func (o *MsgVpn) HasDataTxMsgCount() bool`

HasDataTxMsgCount returns a boolean if a field has been set.

### GetDiscardedRxMsgCount

`func (o *MsgVpn) GetDiscardedRxMsgCount() int64`

GetDiscardedRxMsgCount returns the DiscardedRxMsgCount field if non-nil, zero value otherwise.

### GetDiscardedRxMsgCountOk

`func (o *MsgVpn) GetDiscardedRxMsgCountOk() (*int64, bool)`

GetDiscardedRxMsgCountOk returns a tuple with the DiscardedRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedRxMsgCount

`func (o *MsgVpn) SetDiscardedRxMsgCount(v int64)`

SetDiscardedRxMsgCount sets DiscardedRxMsgCount field to given value.

### HasDiscardedRxMsgCount

`func (o *MsgVpn) HasDiscardedRxMsgCount() bool`

HasDiscardedRxMsgCount returns a boolean if a field has been set.

### GetDiscardedTxMsgCount

`func (o *MsgVpn) GetDiscardedTxMsgCount() int64`

GetDiscardedTxMsgCount returns the DiscardedTxMsgCount field if non-nil, zero value otherwise.

### GetDiscardedTxMsgCountOk

`func (o *MsgVpn) GetDiscardedTxMsgCountOk() (*int64, bool)`

GetDiscardedTxMsgCountOk returns a tuple with the DiscardedTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedTxMsgCount

`func (o *MsgVpn) SetDiscardedTxMsgCount(v int64)`

SetDiscardedTxMsgCount sets DiscardedTxMsgCount field to given value.

### HasDiscardedTxMsgCount

`func (o *MsgVpn) HasDiscardedTxMsgCount() bool`

HasDiscardedTxMsgCount returns a boolean if a field has been set.

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

### GetFailureReason

`func (o *MsgVpn) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MsgVpn) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *MsgVpn) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *MsgVpn) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

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

### GetLoginRxMsgCount

`func (o *MsgVpn) GetLoginRxMsgCount() int64`

GetLoginRxMsgCount returns the LoginRxMsgCount field if non-nil, zero value otherwise.

### GetLoginRxMsgCountOk

`func (o *MsgVpn) GetLoginRxMsgCountOk() (*int64, bool)`

GetLoginRxMsgCountOk returns a tuple with the LoginRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginRxMsgCount

`func (o *MsgVpn) SetLoginRxMsgCount(v int64)`

SetLoginRxMsgCount sets LoginRxMsgCount field to given value.

### HasLoginRxMsgCount

`func (o *MsgVpn) HasLoginRxMsgCount() bool`

HasLoginRxMsgCount returns a boolean if a field has been set.

### GetLoginTxMsgCount

`func (o *MsgVpn) GetLoginTxMsgCount() int64`

GetLoginTxMsgCount returns the LoginTxMsgCount field if non-nil, zero value otherwise.

### GetLoginTxMsgCountOk

`func (o *MsgVpn) GetLoginTxMsgCountOk() (*int64, bool)`

GetLoginTxMsgCountOk returns a tuple with the LoginTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginTxMsgCount

`func (o *MsgVpn) SetLoginTxMsgCount(v int64)`

SetLoginTxMsgCount sets LoginTxMsgCount field to given value.

### HasLoginTxMsgCount

`func (o *MsgVpn) HasLoginTxMsgCount() bool`

HasLoginTxMsgCount returns a boolean if a field has been set.

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

### GetMaxEffectiveEndpointCount

`func (o *MsgVpn) GetMaxEffectiveEndpointCount() int32`

GetMaxEffectiveEndpointCount returns the MaxEffectiveEndpointCount field if non-nil, zero value otherwise.

### GetMaxEffectiveEndpointCountOk

`func (o *MsgVpn) GetMaxEffectiveEndpointCountOk() (*int32, bool)`

GetMaxEffectiveEndpointCountOk returns a tuple with the MaxEffectiveEndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEffectiveEndpointCount

`func (o *MsgVpn) SetMaxEffectiveEndpointCount(v int32)`

SetMaxEffectiveEndpointCount sets MaxEffectiveEndpointCount field to given value.

### HasMaxEffectiveEndpointCount

`func (o *MsgVpn) HasMaxEffectiveEndpointCount() bool`

HasMaxEffectiveEndpointCount returns a boolean if a field has been set.

### GetMaxEffectiveRxFlowCount

`func (o *MsgVpn) GetMaxEffectiveRxFlowCount() int32`

GetMaxEffectiveRxFlowCount returns the MaxEffectiveRxFlowCount field if non-nil, zero value otherwise.

### GetMaxEffectiveRxFlowCountOk

`func (o *MsgVpn) GetMaxEffectiveRxFlowCountOk() (*int32, bool)`

GetMaxEffectiveRxFlowCountOk returns a tuple with the MaxEffectiveRxFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEffectiveRxFlowCount

`func (o *MsgVpn) SetMaxEffectiveRxFlowCount(v int32)`

SetMaxEffectiveRxFlowCount sets MaxEffectiveRxFlowCount field to given value.

### HasMaxEffectiveRxFlowCount

`func (o *MsgVpn) HasMaxEffectiveRxFlowCount() bool`

HasMaxEffectiveRxFlowCount returns a boolean if a field has been set.

### GetMaxEffectiveSubscriptionCount

`func (o *MsgVpn) GetMaxEffectiveSubscriptionCount() int64`

GetMaxEffectiveSubscriptionCount returns the MaxEffectiveSubscriptionCount field if non-nil, zero value otherwise.

### GetMaxEffectiveSubscriptionCountOk

`func (o *MsgVpn) GetMaxEffectiveSubscriptionCountOk() (*int64, bool)`

GetMaxEffectiveSubscriptionCountOk returns a tuple with the MaxEffectiveSubscriptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEffectiveSubscriptionCount

`func (o *MsgVpn) SetMaxEffectiveSubscriptionCount(v int64)`

SetMaxEffectiveSubscriptionCount sets MaxEffectiveSubscriptionCount field to given value.

### HasMaxEffectiveSubscriptionCount

`func (o *MsgVpn) HasMaxEffectiveSubscriptionCount() bool`

HasMaxEffectiveSubscriptionCount returns a boolean if a field has been set.

### GetMaxEffectiveTransactedSessionCount

`func (o *MsgVpn) GetMaxEffectiveTransactedSessionCount() int32`

GetMaxEffectiveTransactedSessionCount returns the MaxEffectiveTransactedSessionCount field if non-nil, zero value otherwise.

### GetMaxEffectiveTransactedSessionCountOk

`func (o *MsgVpn) GetMaxEffectiveTransactedSessionCountOk() (*int32, bool)`

GetMaxEffectiveTransactedSessionCountOk returns a tuple with the MaxEffectiveTransactedSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEffectiveTransactedSessionCount

`func (o *MsgVpn) SetMaxEffectiveTransactedSessionCount(v int32)`

SetMaxEffectiveTransactedSessionCount sets MaxEffectiveTransactedSessionCount field to given value.

### HasMaxEffectiveTransactedSessionCount

`func (o *MsgVpn) HasMaxEffectiveTransactedSessionCount() bool`

HasMaxEffectiveTransactedSessionCount returns a boolean if a field has been set.

### GetMaxEffectiveTransactionCount

`func (o *MsgVpn) GetMaxEffectiveTransactionCount() int32`

GetMaxEffectiveTransactionCount returns the MaxEffectiveTransactionCount field if non-nil, zero value otherwise.

### GetMaxEffectiveTransactionCountOk

`func (o *MsgVpn) GetMaxEffectiveTransactionCountOk() (*int32, bool)`

GetMaxEffectiveTransactionCountOk returns a tuple with the MaxEffectiveTransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEffectiveTransactionCount

`func (o *MsgVpn) SetMaxEffectiveTransactionCount(v int32)`

SetMaxEffectiveTransactionCount sets MaxEffectiveTransactionCount field to given value.

### HasMaxEffectiveTransactionCount

`func (o *MsgVpn) HasMaxEffectiveTransactionCount() bool`

HasMaxEffectiveTransactionCount returns a boolean if a field has been set.

### GetMaxEffectiveTxFlowCount

`func (o *MsgVpn) GetMaxEffectiveTxFlowCount() int32`

GetMaxEffectiveTxFlowCount returns the MaxEffectiveTxFlowCount field if non-nil, zero value otherwise.

### GetMaxEffectiveTxFlowCountOk

`func (o *MsgVpn) GetMaxEffectiveTxFlowCountOk() (*int32, bool)`

GetMaxEffectiveTxFlowCountOk returns a tuple with the MaxEffectiveTxFlowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEffectiveTxFlowCount

`func (o *MsgVpn) SetMaxEffectiveTxFlowCount(v int32)`

SetMaxEffectiveTxFlowCount sets MaxEffectiveTxFlowCount field to given value.

### HasMaxEffectiveTxFlowCount

`func (o *MsgVpn) HasMaxEffectiveTxFlowCount() bool`

HasMaxEffectiveTxFlowCount returns a boolean if a field has been set.

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

### GetMsgReplayActiveCount

`func (o *MsgVpn) GetMsgReplayActiveCount() int32`

GetMsgReplayActiveCount returns the MsgReplayActiveCount field if non-nil, zero value otherwise.

### GetMsgReplayActiveCountOk

`func (o *MsgVpn) GetMsgReplayActiveCountOk() (*int32, bool)`

GetMsgReplayActiveCountOk returns a tuple with the MsgReplayActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgReplayActiveCount

`func (o *MsgVpn) SetMsgReplayActiveCount(v int32)`

SetMsgReplayActiveCount sets MsgReplayActiveCount field to given value.

### HasMsgReplayActiveCount

`func (o *MsgVpn) HasMsgReplayActiveCount() bool`

HasMsgReplayActiveCount returns a boolean if a field has been set.

### GetMsgReplayFailedCount

`func (o *MsgVpn) GetMsgReplayFailedCount() int32`

GetMsgReplayFailedCount returns the MsgReplayFailedCount field if non-nil, zero value otherwise.

### GetMsgReplayFailedCountOk

`func (o *MsgVpn) GetMsgReplayFailedCountOk() (*int32, bool)`

GetMsgReplayFailedCountOk returns a tuple with the MsgReplayFailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgReplayFailedCount

`func (o *MsgVpn) SetMsgReplayFailedCount(v int32)`

SetMsgReplayFailedCount sets MsgReplayFailedCount field to given value.

### HasMsgReplayFailedCount

`func (o *MsgVpn) HasMsgReplayFailedCount() bool`

HasMsgReplayFailedCount returns a boolean if a field has been set.

### GetMsgReplayInitializingCount

`func (o *MsgVpn) GetMsgReplayInitializingCount() int32`

GetMsgReplayInitializingCount returns the MsgReplayInitializingCount field if non-nil, zero value otherwise.

### GetMsgReplayInitializingCountOk

`func (o *MsgVpn) GetMsgReplayInitializingCountOk() (*int32, bool)`

GetMsgReplayInitializingCountOk returns a tuple with the MsgReplayInitializingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgReplayInitializingCount

`func (o *MsgVpn) SetMsgReplayInitializingCount(v int32)`

SetMsgReplayInitializingCount sets MsgReplayInitializingCount field to given value.

### HasMsgReplayInitializingCount

`func (o *MsgVpn) HasMsgReplayInitializingCount() bool`

HasMsgReplayInitializingCount returns a boolean if a field has been set.

### GetMsgReplayPendingCompleteCount

`func (o *MsgVpn) GetMsgReplayPendingCompleteCount() int32`

GetMsgReplayPendingCompleteCount returns the MsgReplayPendingCompleteCount field if non-nil, zero value otherwise.

### GetMsgReplayPendingCompleteCountOk

`func (o *MsgVpn) GetMsgReplayPendingCompleteCountOk() (*int32, bool)`

GetMsgReplayPendingCompleteCountOk returns a tuple with the MsgReplayPendingCompleteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgReplayPendingCompleteCount

`func (o *MsgVpn) SetMsgReplayPendingCompleteCount(v int32)`

SetMsgReplayPendingCompleteCount sets MsgReplayPendingCompleteCount field to given value.

### HasMsgReplayPendingCompleteCount

`func (o *MsgVpn) HasMsgReplayPendingCompleteCount() bool`

HasMsgReplayPendingCompleteCount returns a boolean if a field has been set.

### GetMsgSpoolMsgCount

`func (o *MsgVpn) GetMsgSpoolMsgCount() int64`

GetMsgSpoolMsgCount returns the MsgSpoolMsgCount field if non-nil, zero value otherwise.

### GetMsgSpoolMsgCountOk

`func (o *MsgVpn) GetMsgSpoolMsgCountOk() (*int64, bool)`

GetMsgSpoolMsgCountOk returns a tuple with the MsgSpoolMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolMsgCount

`func (o *MsgVpn) SetMsgSpoolMsgCount(v int64)`

SetMsgSpoolMsgCount sets MsgSpoolMsgCount field to given value.

### HasMsgSpoolMsgCount

`func (o *MsgVpn) HasMsgSpoolMsgCount() bool`

HasMsgSpoolMsgCount returns a boolean if a field has been set.

### GetMsgSpoolRxMsgCount

`func (o *MsgVpn) GetMsgSpoolRxMsgCount() int64`

GetMsgSpoolRxMsgCount returns the MsgSpoolRxMsgCount field if non-nil, zero value otherwise.

### GetMsgSpoolRxMsgCountOk

`func (o *MsgVpn) GetMsgSpoolRxMsgCountOk() (*int64, bool)`

GetMsgSpoolRxMsgCountOk returns a tuple with the MsgSpoolRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolRxMsgCount

`func (o *MsgVpn) SetMsgSpoolRxMsgCount(v int64)`

SetMsgSpoolRxMsgCount sets MsgSpoolRxMsgCount field to given value.

### HasMsgSpoolRxMsgCount

`func (o *MsgVpn) HasMsgSpoolRxMsgCount() bool`

HasMsgSpoolRxMsgCount returns a boolean if a field has been set.

### GetMsgSpoolTxMsgCount

`func (o *MsgVpn) GetMsgSpoolTxMsgCount() int64`

GetMsgSpoolTxMsgCount returns the MsgSpoolTxMsgCount field if non-nil, zero value otherwise.

### GetMsgSpoolTxMsgCountOk

`func (o *MsgVpn) GetMsgSpoolTxMsgCountOk() (*int64, bool)`

GetMsgSpoolTxMsgCountOk returns a tuple with the MsgSpoolTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolTxMsgCount

`func (o *MsgVpn) SetMsgSpoolTxMsgCount(v int64)`

SetMsgSpoolTxMsgCount sets MsgSpoolTxMsgCount field to given value.

### HasMsgSpoolTxMsgCount

`func (o *MsgVpn) HasMsgSpoolTxMsgCount() bool`

HasMsgSpoolTxMsgCount returns a boolean if a field has been set.

### GetMsgSpoolUsage

`func (o *MsgVpn) GetMsgSpoolUsage() int64`

GetMsgSpoolUsage returns the MsgSpoolUsage field if non-nil, zero value otherwise.

### GetMsgSpoolUsageOk

`func (o *MsgVpn) GetMsgSpoolUsageOk() (*int64, bool)`

GetMsgSpoolUsageOk returns a tuple with the MsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolUsage

`func (o *MsgVpn) SetMsgSpoolUsage(v int64)`

SetMsgSpoolUsage sets MsgSpoolUsage field to given value.

### HasMsgSpoolUsage

`func (o *MsgVpn) HasMsgSpoolUsage() bool`

HasMsgSpoolUsage returns a boolean if a field has been set.

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

### GetRate

`func (o *MsgVpn) GetRate() MsgVpnRate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MsgVpn) GetRateOk() (*MsgVpnRate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MsgVpn) SetRate(v MsgVpnRate)`

SetRate sets Rate field to given value.

### HasRate

`func (o *MsgVpn) HasRate() bool`

HasRate returns a boolean if a field has been set.

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

### GetReplicationActiveAckPropTxMsgCount

`func (o *MsgVpn) GetReplicationActiveAckPropTxMsgCount() int64`

GetReplicationActiveAckPropTxMsgCount returns the ReplicationActiveAckPropTxMsgCount field if non-nil, zero value otherwise.

### GetReplicationActiveAckPropTxMsgCountOk

`func (o *MsgVpn) GetReplicationActiveAckPropTxMsgCountOk() (*int64, bool)`

GetReplicationActiveAckPropTxMsgCountOk returns a tuple with the ReplicationActiveAckPropTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveAckPropTxMsgCount

`func (o *MsgVpn) SetReplicationActiveAckPropTxMsgCount(v int64)`

SetReplicationActiveAckPropTxMsgCount sets ReplicationActiveAckPropTxMsgCount field to given value.

### HasReplicationActiveAckPropTxMsgCount

`func (o *MsgVpn) HasReplicationActiveAckPropTxMsgCount() bool`

HasReplicationActiveAckPropTxMsgCount returns a boolean if a field has been set.

### GetReplicationActiveAsyncQueuedMsgCount

`func (o *MsgVpn) GetReplicationActiveAsyncQueuedMsgCount() int64`

GetReplicationActiveAsyncQueuedMsgCount returns the ReplicationActiveAsyncQueuedMsgCount field if non-nil, zero value otherwise.

### GetReplicationActiveAsyncQueuedMsgCountOk

`func (o *MsgVpn) GetReplicationActiveAsyncQueuedMsgCountOk() (*int64, bool)`

GetReplicationActiveAsyncQueuedMsgCountOk returns a tuple with the ReplicationActiveAsyncQueuedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveAsyncQueuedMsgCount

`func (o *MsgVpn) SetReplicationActiveAsyncQueuedMsgCount(v int64)`

SetReplicationActiveAsyncQueuedMsgCount sets ReplicationActiveAsyncQueuedMsgCount field to given value.

### HasReplicationActiveAsyncQueuedMsgCount

`func (o *MsgVpn) HasReplicationActiveAsyncQueuedMsgCount() bool`

HasReplicationActiveAsyncQueuedMsgCount returns a boolean if a field has been set.

### GetReplicationActiveLocallyConsumedMsgCount

`func (o *MsgVpn) GetReplicationActiveLocallyConsumedMsgCount() int64`

GetReplicationActiveLocallyConsumedMsgCount returns the ReplicationActiveLocallyConsumedMsgCount field if non-nil, zero value otherwise.

### GetReplicationActiveLocallyConsumedMsgCountOk

`func (o *MsgVpn) GetReplicationActiveLocallyConsumedMsgCountOk() (*int64, bool)`

GetReplicationActiveLocallyConsumedMsgCountOk returns a tuple with the ReplicationActiveLocallyConsumedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveLocallyConsumedMsgCount

`func (o *MsgVpn) SetReplicationActiveLocallyConsumedMsgCount(v int64)`

SetReplicationActiveLocallyConsumedMsgCount sets ReplicationActiveLocallyConsumedMsgCount field to given value.

### HasReplicationActiveLocallyConsumedMsgCount

`func (o *MsgVpn) HasReplicationActiveLocallyConsumedMsgCount() bool`

HasReplicationActiveLocallyConsumedMsgCount returns a boolean if a field has been set.

### GetReplicationActiveMateFlowCongestedPeakTime

`func (o *MsgVpn) GetReplicationActiveMateFlowCongestedPeakTime() int32`

GetReplicationActiveMateFlowCongestedPeakTime returns the ReplicationActiveMateFlowCongestedPeakTime field if non-nil, zero value otherwise.

### GetReplicationActiveMateFlowCongestedPeakTimeOk

`func (o *MsgVpn) GetReplicationActiveMateFlowCongestedPeakTimeOk() (*int32, bool)`

GetReplicationActiveMateFlowCongestedPeakTimeOk returns a tuple with the ReplicationActiveMateFlowCongestedPeakTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveMateFlowCongestedPeakTime

`func (o *MsgVpn) SetReplicationActiveMateFlowCongestedPeakTime(v int32)`

SetReplicationActiveMateFlowCongestedPeakTime sets ReplicationActiveMateFlowCongestedPeakTime field to given value.

### HasReplicationActiveMateFlowCongestedPeakTime

`func (o *MsgVpn) HasReplicationActiveMateFlowCongestedPeakTime() bool`

HasReplicationActiveMateFlowCongestedPeakTime returns a boolean if a field has been set.

### GetReplicationActiveMateFlowNotCongestedPeakTime

`func (o *MsgVpn) GetReplicationActiveMateFlowNotCongestedPeakTime() int32`

GetReplicationActiveMateFlowNotCongestedPeakTime returns the ReplicationActiveMateFlowNotCongestedPeakTime field if non-nil, zero value otherwise.

### GetReplicationActiveMateFlowNotCongestedPeakTimeOk

`func (o *MsgVpn) GetReplicationActiveMateFlowNotCongestedPeakTimeOk() (*int32, bool)`

GetReplicationActiveMateFlowNotCongestedPeakTimeOk returns a tuple with the ReplicationActiveMateFlowNotCongestedPeakTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveMateFlowNotCongestedPeakTime

`func (o *MsgVpn) SetReplicationActiveMateFlowNotCongestedPeakTime(v int32)`

SetReplicationActiveMateFlowNotCongestedPeakTime sets ReplicationActiveMateFlowNotCongestedPeakTime field to given value.

### HasReplicationActiveMateFlowNotCongestedPeakTime

`func (o *MsgVpn) HasReplicationActiveMateFlowNotCongestedPeakTime() bool`

HasReplicationActiveMateFlowNotCongestedPeakTime returns a boolean if a field has been set.

### GetReplicationActivePromotedQueuedMsgCount

`func (o *MsgVpn) GetReplicationActivePromotedQueuedMsgCount() int64`

GetReplicationActivePromotedQueuedMsgCount returns the ReplicationActivePromotedQueuedMsgCount field if non-nil, zero value otherwise.

### GetReplicationActivePromotedQueuedMsgCountOk

`func (o *MsgVpn) GetReplicationActivePromotedQueuedMsgCountOk() (*int64, bool)`

GetReplicationActivePromotedQueuedMsgCountOk returns a tuple with the ReplicationActivePromotedQueuedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActivePromotedQueuedMsgCount

`func (o *MsgVpn) SetReplicationActivePromotedQueuedMsgCount(v int64)`

SetReplicationActivePromotedQueuedMsgCount sets ReplicationActivePromotedQueuedMsgCount field to given value.

### HasReplicationActivePromotedQueuedMsgCount

`func (o *MsgVpn) HasReplicationActivePromotedQueuedMsgCount() bool`

HasReplicationActivePromotedQueuedMsgCount returns a boolean if a field has been set.

### GetReplicationActiveReconcileRequestRxMsgCount

`func (o *MsgVpn) GetReplicationActiveReconcileRequestRxMsgCount() int64`

GetReplicationActiveReconcileRequestRxMsgCount returns the ReplicationActiveReconcileRequestRxMsgCount field if non-nil, zero value otherwise.

### GetReplicationActiveReconcileRequestRxMsgCountOk

`func (o *MsgVpn) GetReplicationActiveReconcileRequestRxMsgCountOk() (*int64, bool)`

GetReplicationActiveReconcileRequestRxMsgCountOk returns a tuple with the ReplicationActiveReconcileRequestRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveReconcileRequestRxMsgCount

`func (o *MsgVpn) SetReplicationActiveReconcileRequestRxMsgCount(v int64)`

SetReplicationActiveReconcileRequestRxMsgCount sets ReplicationActiveReconcileRequestRxMsgCount field to given value.

### HasReplicationActiveReconcileRequestRxMsgCount

`func (o *MsgVpn) HasReplicationActiveReconcileRequestRxMsgCount() bool`

HasReplicationActiveReconcileRequestRxMsgCount returns a boolean if a field has been set.

### GetReplicationActiveSyncEligiblePeakTime

`func (o *MsgVpn) GetReplicationActiveSyncEligiblePeakTime() int32`

GetReplicationActiveSyncEligiblePeakTime returns the ReplicationActiveSyncEligiblePeakTime field if non-nil, zero value otherwise.

### GetReplicationActiveSyncEligiblePeakTimeOk

`func (o *MsgVpn) GetReplicationActiveSyncEligiblePeakTimeOk() (*int32, bool)`

GetReplicationActiveSyncEligiblePeakTimeOk returns a tuple with the ReplicationActiveSyncEligiblePeakTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveSyncEligiblePeakTime

`func (o *MsgVpn) SetReplicationActiveSyncEligiblePeakTime(v int32)`

SetReplicationActiveSyncEligiblePeakTime sets ReplicationActiveSyncEligiblePeakTime field to given value.

### HasReplicationActiveSyncEligiblePeakTime

`func (o *MsgVpn) HasReplicationActiveSyncEligiblePeakTime() bool`

HasReplicationActiveSyncEligiblePeakTime returns a boolean if a field has been set.

### GetReplicationActiveSyncIneligiblePeakTime

`func (o *MsgVpn) GetReplicationActiveSyncIneligiblePeakTime() int32`

GetReplicationActiveSyncIneligiblePeakTime returns the ReplicationActiveSyncIneligiblePeakTime field if non-nil, zero value otherwise.

### GetReplicationActiveSyncIneligiblePeakTimeOk

`func (o *MsgVpn) GetReplicationActiveSyncIneligiblePeakTimeOk() (*int32, bool)`

GetReplicationActiveSyncIneligiblePeakTimeOk returns a tuple with the ReplicationActiveSyncIneligiblePeakTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveSyncIneligiblePeakTime

`func (o *MsgVpn) SetReplicationActiveSyncIneligiblePeakTime(v int32)`

SetReplicationActiveSyncIneligiblePeakTime sets ReplicationActiveSyncIneligiblePeakTime field to given value.

### HasReplicationActiveSyncIneligiblePeakTime

`func (o *MsgVpn) HasReplicationActiveSyncIneligiblePeakTime() bool`

HasReplicationActiveSyncIneligiblePeakTime returns a boolean if a field has been set.

### GetReplicationActiveSyncQueuedAsAsyncMsgCount

`func (o *MsgVpn) GetReplicationActiveSyncQueuedAsAsyncMsgCount() int64`

GetReplicationActiveSyncQueuedAsAsyncMsgCount returns the ReplicationActiveSyncQueuedAsAsyncMsgCount field if non-nil, zero value otherwise.

### GetReplicationActiveSyncQueuedAsAsyncMsgCountOk

`func (o *MsgVpn) GetReplicationActiveSyncQueuedAsAsyncMsgCountOk() (*int64, bool)`

GetReplicationActiveSyncQueuedAsAsyncMsgCountOk returns a tuple with the ReplicationActiveSyncQueuedAsAsyncMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveSyncQueuedAsAsyncMsgCount

`func (o *MsgVpn) SetReplicationActiveSyncQueuedAsAsyncMsgCount(v int64)`

SetReplicationActiveSyncQueuedAsAsyncMsgCount sets ReplicationActiveSyncQueuedAsAsyncMsgCount field to given value.

### HasReplicationActiveSyncQueuedAsAsyncMsgCount

`func (o *MsgVpn) HasReplicationActiveSyncQueuedAsAsyncMsgCount() bool`

HasReplicationActiveSyncQueuedAsAsyncMsgCount returns a boolean if a field has been set.

### GetReplicationActiveSyncQueuedMsgCount

`func (o *MsgVpn) GetReplicationActiveSyncQueuedMsgCount() int64`

GetReplicationActiveSyncQueuedMsgCount returns the ReplicationActiveSyncQueuedMsgCount field if non-nil, zero value otherwise.

### GetReplicationActiveSyncQueuedMsgCountOk

`func (o *MsgVpn) GetReplicationActiveSyncQueuedMsgCountOk() (*int64, bool)`

GetReplicationActiveSyncQueuedMsgCountOk returns a tuple with the ReplicationActiveSyncQueuedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveSyncQueuedMsgCount

`func (o *MsgVpn) SetReplicationActiveSyncQueuedMsgCount(v int64)`

SetReplicationActiveSyncQueuedMsgCount sets ReplicationActiveSyncQueuedMsgCount field to given value.

### HasReplicationActiveSyncQueuedMsgCount

`func (o *MsgVpn) HasReplicationActiveSyncQueuedMsgCount() bool`

HasReplicationActiveSyncQueuedMsgCount returns a boolean if a field has been set.

### GetReplicationActiveTransitionToSyncIneligibleCount

`func (o *MsgVpn) GetReplicationActiveTransitionToSyncIneligibleCount() int64`

GetReplicationActiveTransitionToSyncIneligibleCount returns the ReplicationActiveTransitionToSyncIneligibleCount field if non-nil, zero value otherwise.

### GetReplicationActiveTransitionToSyncIneligibleCountOk

`func (o *MsgVpn) GetReplicationActiveTransitionToSyncIneligibleCountOk() (*int64, bool)`

GetReplicationActiveTransitionToSyncIneligibleCountOk returns a tuple with the ReplicationActiveTransitionToSyncIneligibleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationActiveTransitionToSyncIneligibleCount

`func (o *MsgVpn) SetReplicationActiveTransitionToSyncIneligibleCount(v int64)`

SetReplicationActiveTransitionToSyncIneligibleCount sets ReplicationActiveTransitionToSyncIneligibleCount field to given value.

### HasReplicationActiveTransitionToSyncIneligibleCount

`func (o *MsgVpn) HasReplicationActiveTransitionToSyncIneligibleCount() bool`

HasReplicationActiveTransitionToSyncIneligibleCount returns a boolean if a field has been set.

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

### GetReplicationBridgeBoundToQueue

`func (o *MsgVpn) GetReplicationBridgeBoundToQueue() bool`

GetReplicationBridgeBoundToQueue returns the ReplicationBridgeBoundToQueue field if non-nil, zero value otherwise.

### GetReplicationBridgeBoundToQueueOk

`func (o *MsgVpn) GetReplicationBridgeBoundToQueueOk() (*bool, bool)`

GetReplicationBridgeBoundToQueueOk returns a tuple with the ReplicationBridgeBoundToQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeBoundToQueue

`func (o *MsgVpn) SetReplicationBridgeBoundToQueue(v bool)`

SetReplicationBridgeBoundToQueue sets ReplicationBridgeBoundToQueue field to given value.

### HasReplicationBridgeBoundToQueue

`func (o *MsgVpn) HasReplicationBridgeBoundToQueue() bool`

HasReplicationBridgeBoundToQueue returns a boolean if a field has been set.

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

### GetReplicationBridgeName

`func (o *MsgVpn) GetReplicationBridgeName() string`

GetReplicationBridgeName returns the ReplicationBridgeName field if non-nil, zero value otherwise.

### GetReplicationBridgeNameOk

`func (o *MsgVpn) GetReplicationBridgeNameOk() (*string, bool)`

GetReplicationBridgeNameOk returns a tuple with the ReplicationBridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeName

`func (o *MsgVpn) SetReplicationBridgeName(v string)`

SetReplicationBridgeName sets ReplicationBridgeName field to given value.

### HasReplicationBridgeName

`func (o *MsgVpn) HasReplicationBridgeName() bool`

HasReplicationBridgeName returns a boolean if a field has been set.

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

### GetReplicationBridgeUp

`func (o *MsgVpn) GetReplicationBridgeUp() bool`

GetReplicationBridgeUp returns the ReplicationBridgeUp field if non-nil, zero value otherwise.

### GetReplicationBridgeUpOk

`func (o *MsgVpn) GetReplicationBridgeUpOk() (*bool, bool)`

GetReplicationBridgeUpOk returns a tuple with the ReplicationBridgeUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationBridgeUp

`func (o *MsgVpn) SetReplicationBridgeUp(v bool)`

SetReplicationBridgeUp sets ReplicationBridgeUp field to given value.

### HasReplicationBridgeUp

`func (o *MsgVpn) HasReplicationBridgeUp() bool`

HasReplicationBridgeUp returns a boolean if a field has been set.

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

### GetReplicationQueueBound

`func (o *MsgVpn) GetReplicationQueueBound() bool`

GetReplicationQueueBound returns the ReplicationQueueBound field if non-nil, zero value otherwise.

### GetReplicationQueueBoundOk

`func (o *MsgVpn) GetReplicationQueueBoundOk() (*bool, bool)`

GetReplicationQueueBoundOk returns a tuple with the ReplicationQueueBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationQueueBound

`func (o *MsgVpn) SetReplicationQueueBound(v bool)`

SetReplicationQueueBound sets ReplicationQueueBound field to given value.

### HasReplicationQueueBound

`func (o *MsgVpn) HasReplicationQueueBound() bool`

HasReplicationQueueBound returns a boolean if a field has been set.

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

### GetReplicationRemoteBridgeName

`func (o *MsgVpn) GetReplicationRemoteBridgeName() string`

GetReplicationRemoteBridgeName returns the ReplicationRemoteBridgeName field if non-nil, zero value otherwise.

### GetReplicationRemoteBridgeNameOk

`func (o *MsgVpn) GetReplicationRemoteBridgeNameOk() (*string, bool)`

GetReplicationRemoteBridgeNameOk returns a tuple with the ReplicationRemoteBridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRemoteBridgeName

`func (o *MsgVpn) SetReplicationRemoteBridgeName(v string)`

SetReplicationRemoteBridgeName sets ReplicationRemoteBridgeName field to given value.

### HasReplicationRemoteBridgeName

`func (o *MsgVpn) HasReplicationRemoteBridgeName() bool`

HasReplicationRemoteBridgeName returns a boolean if a field has been set.

### GetReplicationRemoteBridgeUp

`func (o *MsgVpn) GetReplicationRemoteBridgeUp() bool`

GetReplicationRemoteBridgeUp returns the ReplicationRemoteBridgeUp field if non-nil, zero value otherwise.

### GetReplicationRemoteBridgeUpOk

`func (o *MsgVpn) GetReplicationRemoteBridgeUpOk() (*bool, bool)`

GetReplicationRemoteBridgeUpOk returns a tuple with the ReplicationRemoteBridgeUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRemoteBridgeUp

`func (o *MsgVpn) SetReplicationRemoteBridgeUp(v bool)`

SetReplicationRemoteBridgeUp sets ReplicationRemoteBridgeUp field to given value.

### HasReplicationRemoteBridgeUp

`func (o *MsgVpn) HasReplicationRemoteBridgeUp() bool`

HasReplicationRemoteBridgeUp returns a boolean if a field has been set.

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

### GetReplicationStandbyAckPropOutOfSeqRxMsgCount

`func (o *MsgVpn) GetReplicationStandbyAckPropOutOfSeqRxMsgCount() int64`

GetReplicationStandbyAckPropOutOfSeqRxMsgCount returns the ReplicationStandbyAckPropOutOfSeqRxMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyAckPropOutOfSeqRxMsgCountOk

`func (o *MsgVpn) GetReplicationStandbyAckPropOutOfSeqRxMsgCountOk() (*int64, bool)`

GetReplicationStandbyAckPropOutOfSeqRxMsgCountOk returns a tuple with the ReplicationStandbyAckPropOutOfSeqRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyAckPropOutOfSeqRxMsgCount

`func (o *MsgVpn) SetReplicationStandbyAckPropOutOfSeqRxMsgCount(v int64)`

SetReplicationStandbyAckPropOutOfSeqRxMsgCount sets ReplicationStandbyAckPropOutOfSeqRxMsgCount field to given value.

### HasReplicationStandbyAckPropOutOfSeqRxMsgCount

`func (o *MsgVpn) HasReplicationStandbyAckPropOutOfSeqRxMsgCount() bool`

HasReplicationStandbyAckPropOutOfSeqRxMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyAckPropRxMsgCount

`func (o *MsgVpn) GetReplicationStandbyAckPropRxMsgCount() int64`

GetReplicationStandbyAckPropRxMsgCount returns the ReplicationStandbyAckPropRxMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyAckPropRxMsgCountOk

`func (o *MsgVpn) GetReplicationStandbyAckPropRxMsgCountOk() (*int64, bool)`

GetReplicationStandbyAckPropRxMsgCountOk returns a tuple with the ReplicationStandbyAckPropRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyAckPropRxMsgCount

`func (o *MsgVpn) SetReplicationStandbyAckPropRxMsgCount(v int64)`

SetReplicationStandbyAckPropRxMsgCount sets ReplicationStandbyAckPropRxMsgCount field to given value.

### HasReplicationStandbyAckPropRxMsgCount

`func (o *MsgVpn) HasReplicationStandbyAckPropRxMsgCount() bool`

HasReplicationStandbyAckPropRxMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyReconcileRequestTxMsgCount

`func (o *MsgVpn) GetReplicationStandbyReconcileRequestTxMsgCount() int64`

GetReplicationStandbyReconcileRequestTxMsgCount returns the ReplicationStandbyReconcileRequestTxMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyReconcileRequestTxMsgCountOk

`func (o *MsgVpn) GetReplicationStandbyReconcileRequestTxMsgCountOk() (*int64, bool)`

GetReplicationStandbyReconcileRequestTxMsgCountOk returns a tuple with the ReplicationStandbyReconcileRequestTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyReconcileRequestTxMsgCount

`func (o *MsgVpn) SetReplicationStandbyReconcileRequestTxMsgCount(v int64)`

SetReplicationStandbyReconcileRequestTxMsgCount sets ReplicationStandbyReconcileRequestTxMsgCount field to given value.

### HasReplicationStandbyReconcileRequestTxMsgCount

`func (o *MsgVpn) HasReplicationStandbyReconcileRequestTxMsgCount() bool`

HasReplicationStandbyReconcileRequestTxMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyRxMsgCount

`func (o *MsgVpn) GetReplicationStandbyRxMsgCount() int64`

GetReplicationStandbyRxMsgCount returns the ReplicationStandbyRxMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyRxMsgCountOk

`func (o *MsgVpn) GetReplicationStandbyRxMsgCountOk() (*int64, bool)`

GetReplicationStandbyRxMsgCountOk returns a tuple with the ReplicationStandbyRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyRxMsgCount

`func (o *MsgVpn) SetReplicationStandbyRxMsgCount(v int64)`

SetReplicationStandbyRxMsgCount sets ReplicationStandbyRxMsgCount field to given value.

### HasReplicationStandbyRxMsgCount

`func (o *MsgVpn) HasReplicationStandbyRxMsgCount() bool`

HasReplicationStandbyRxMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyTransactionRequestCount

`func (o *MsgVpn) GetReplicationStandbyTransactionRequestCount() int64`

GetReplicationStandbyTransactionRequestCount returns the ReplicationStandbyTransactionRequestCount field if non-nil, zero value otherwise.

### GetReplicationStandbyTransactionRequestCountOk

`func (o *MsgVpn) GetReplicationStandbyTransactionRequestCountOk() (*int64, bool)`

GetReplicationStandbyTransactionRequestCountOk returns a tuple with the ReplicationStandbyTransactionRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyTransactionRequestCount

`func (o *MsgVpn) SetReplicationStandbyTransactionRequestCount(v int64)`

SetReplicationStandbyTransactionRequestCount sets ReplicationStandbyTransactionRequestCount field to given value.

### HasReplicationStandbyTransactionRequestCount

`func (o *MsgVpn) HasReplicationStandbyTransactionRequestCount() bool`

HasReplicationStandbyTransactionRequestCount returns a boolean if a field has been set.

### GetReplicationStandbyTransactionRequestFailureCount

`func (o *MsgVpn) GetReplicationStandbyTransactionRequestFailureCount() int64`

GetReplicationStandbyTransactionRequestFailureCount returns the ReplicationStandbyTransactionRequestFailureCount field if non-nil, zero value otherwise.

### GetReplicationStandbyTransactionRequestFailureCountOk

`func (o *MsgVpn) GetReplicationStandbyTransactionRequestFailureCountOk() (*int64, bool)`

GetReplicationStandbyTransactionRequestFailureCountOk returns a tuple with the ReplicationStandbyTransactionRequestFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyTransactionRequestFailureCount

`func (o *MsgVpn) SetReplicationStandbyTransactionRequestFailureCount(v int64)`

SetReplicationStandbyTransactionRequestFailureCount sets ReplicationStandbyTransactionRequestFailureCount field to given value.

### HasReplicationStandbyTransactionRequestFailureCount

`func (o *MsgVpn) HasReplicationStandbyTransactionRequestFailureCount() bool`

HasReplicationStandbyTransactionRequestFailureCount returns a boolean if a field has been set.

### GetReplicationStandbyTransactionRequestSuccessCount

`func (o *MsgVpn) GetReplicationStandbyTransactionRequestSuccessCount() int64`

GetReplicationStandbyTransactionRequestSuccessCount returns the ReplicationStandbyTransactionRequestSuccessCount field if non-nil, zero value otherwise.

### GetReplicationStandbyTransactionRequestSuccessCountOk

`func (o *MsgVpn) GetReplicationStandbyTransactionRequestSuccessCountOk() (*int64, bool)`

GetReplicationStandbyTransactionRequestSuccessCountOk returns a tuple with the ReplicationStandbyTransactionRequestSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyTransactionRequestSuccessCount

`func (o *MsgVpn) SetReplicationStandbyTransactionRequestSuccessCount(v int64)`

SetReplicationStandbyTransactionRequestSuccessCount sets ReplicationStandbyTransactionRequestSuccessCount field to given value.

### HasReplicationStandbyTransactionRequestSuccessCount

`func (o *MsgVpn) HasReplicationStandbyTransactionRequestSuccessCount() bool`

HasReplicationStandbyTransactionRequestSuccessCount returns a boolean if a field has been set.

### GetReplicationSyncEligible

`func (o *MsgVpn) GetReplicationSyncEligible() bool`

GetReplicationSyncEligible returns the ReplicationSyncEligible field if non-nil, zero value otherwise.

### GetReplicationSyncEligibleOk

`func (o *MsgVpn) GetReplicationSyncEligibleOk() (*bool, bool)`

GetReplicationSyncEligibleOk returns a tuple with the ReplicationSyncEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSyncEligible

`func (o *MsgVpn) SetReplicationSyncEligible(v bool)`

SetReplicationSyncEligible sets ReplicationSyncEligible field to given value.

### HasReplicationSyncEligible

`func (o *MsgVpn) HasReplicationSyncEligible() bool`

HasReplicationSyncEligible returns a boolean if a field has been set.

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

### GetRxByteCount

`func (o *MsgVpn) GetRxByteCount() int64`

GetRxByteCount returns the RxByteCount field if non-nil, zero value otherwise.

### GetRxByteCountOk

`func (o *MsgVpn) GetRxByteCountOk() (*int64, bool)`

GetRxByteCountOk returns a tuple with the RxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteCount

`func (o *MsgVpn) SetRxByteCount(v int64)`

SetRxByteCount sets RxByteCount field to given value.

### HasRxByteCount

`func (o *MsgVpn) HasRxByteCount() bool`

HasRxByteCount returns a boolean if a field has been set.

### GetRxByteRate

`func (o *MsgVpn) GetRxByteRate() int64`

GetRxByteRate returns the RxByteRate field if non-nil, zero value otherwise.

### GetRxByteRateOk

`func (o *MsgVpn) GetRxByteRateOk() (*int64, bool)`

GetRxByteRateOk returns a tuple with the RxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteRate

`func (o *MsgVpn) SetRxByteRate(v int64)`

SetRxByteRate sets RxByteRate field to given value.

### HasRxByteRate

`func (o *MsgVpn) HasRxByteRate() bool`

HasRxByteRate returns a boolean if a field has been set.

### GetRxCompressedByteCount

`func (o *MsgVpn) GetRxCompressedByteCount() int64`

GetRxCompressedByteCount returns the RxCompressedByteCount field if non-nil, zero value otherwise.

### GetRxCompressedByteCountOk

`func (o *MsgVpn) GetRxCompressedByteCountOk() (*int64, bool)`

GetRxCompressedByteCountOk returns a tuple with the RxCompressedByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxCompressedByteCount

`func (o *MsgVpn) SetRxCompressedByteCount(v int64)`

SetRxCompressedByteCount sets RxCompressedByteCount field to given value.

### HasRxCompressedByteCount

`func (o *MsgVpn) HasRxCompressedByteCount() bool`

HasRxCompressedByteCount returns a boolean if a field has been set.

### GetRxCompressedByteRate

`func (o *MsgVpn) GetRxCompressedByteRate() int64`

GetRxCompressedByteRate returns the RxCompressedByteRate field if non-nil, zero value otherwise.

### GetRxCompressedByteRateOk

`func (o *MsgVpn) GetRxCompressedByteRateOk() (*int64, bool)`

GetRxCompressedByteRateOk returns a tuple with the RxCompressedByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxCompressedByteRate

`func (o *MsgVpn) SetRxCompressedByteRate(v int64)`

SetRxCompressedByteRate sets RxCompressedByteRate field to given value.

### HasRxCompressedByteRate

`func (o *MsgVpn) HasRxCompressedByteRate() bool`

HasRxCompressedByteRate returns a boolean if a field has been set.

### GetRxCompressionRatio

`func (o *MsgVpn) GetRxCompressionRatio() string`

GetRxCompressionRatio returns the RxCompressionRatio field if non-nil, zero value otherwise.

### GetRxCompressionRatioOk

`func (o *MsgVpn) GetRxCompressionRatioOk() (*string, bool)`

GetRxCompressionRatioOk returns a tuple with the RxCompressionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxCompressionRatio

`func (o *MsgVpn) SetRxCompressionRatio(v string)`

SetRxCompressionRatio sets RxCompressionRatio field to given value.

### HasRxCompressionRatio

`func (o *MsgVpn) HasRxCompressionRatio() bool`

HasRxCompressionRatio returns a boolean if a field has been set.

### GetRxMsgCount

`func (o *MsgVpn) GetRxMsgCount() int64`

GetRxMsgCount returns the RxMsgCount field if non-nil, zero value otherwise.

### GetRxMsgCountOk

`func (o *MsgVpn) GetRxMsgCountOk() (*int64, bool)`

GetRxMsgCountOk returns a tuple with the RxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgCount

`func (o *MsgVpn) SetRxMsgCount(v int64)`

SetRxMsgCount sets RxMsgCount field to given value.

### HasRxMsgCount

`func (o *MsgVpn) HasRxMsgCount() bool`

HasRxMsgCount returns a boolean if a field has been set.

### GetRxMsgRate

`func (o *MsgVpn) GetRxMsgRate() int64`

GetRxMsgRate returns the RxMsgRate field if non-nil, zero value otherwise.

### GetRxMsgRateOk

`func (o *MsgVpn) GetRxMsgRateOk() (*int64, bool)`

GetRxMsgRateOk returns a tuple with the RxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgRate

`func (o *MsgVpn) SetRxMsgRate(v int64)`

SetRxMsgRate sets RxMsgRate field to given value.

### HasRxMsgRate

`func (o *MsgVpn) HasRxMsgRate() bool`

HasRxMsgRate returns a boolean if a field has been set.

### GetRxUncompressedByteCount

`func (o *MsgVpn) GetRxUncompressedByteCount() int64`

GetRxUncompressedByteCount returns the RxUncompressedByteCount field if non-nil, zero value otherwise.

### GetRxUncompressedByteCountOk

`func (o *MsgVpn) GetRxUncompressedByteCountOk() (*int64, bool)`

GetRxUncompressedByteCountOk returns a tuple with the RxUncompressedByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxUncompressedByteCount

`func (o *MsgVpn) SetRxUncompressedByteCount(v int64)`

SetRxUncompressedByteCount sets RxUncompressedByteCount field to given value.

### HasRxUncompressedByteCount

`func (o *MsgVpn) HasRxUncompressedByteCount() bool`

HasRxUncompressedByteCount returns a boolean if a field has been set.

### GetRxUncompressedByteRate

`func (o *MsgVpn) GetRxUncompressedByteRate() int64`

GetRxUncompressedByteRate returns the RxUncompressedByteRate field if non-nil, zero value otherwise.

### GetRxUncompressedByteRateOk

`func (o *MsgVpn) GetRxUncompressedByteRateOk() (*int64, bool)`

GetRxUncompressedByteRateOk returns a tuple with the RxUncompressedByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxUncompressedByteRate

`func (o *MsgVpn) SetRxUncompressedByteRate(v int64)`

SetRxUncompressedByteRate sets RxUncompressedByteRate field to given value.

### HasRxUncompressedByteRate

`func (o *MsgVpn) HasRxUncompressedByteRate() bool`

HasRxUncompressedByteRate returns a boolean if a field has been set.

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

### GetServiceAmqpPlainTextCompressed

`func (o *MsgVpn) GetServiceAmqpPlainTextCompressed() bool`

GetServiceAmqpPlainTextCompressed returns the ServiceAmqpPlainTextCompressed field if non-nil, zero value otherwise.

### GetServiceAmqpPlainTextCompressedOk

`func (o *MsgVpn) GetServiceAmqpPlainTextCompressedOk() (*bool, bool)`

GetServiceAmqpPlainTextCompressedOk returns a tuple with the ServiceAmqpPlainTextCompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpPlainTextCompressed

`func (o *MsgVpn) SetServiceAmqpPlainTextCompressed(v bool)`

SetServiceAmqpPlainTextCompressed sets ServiceAmqpPlainTextCompressed field to given value.

### HasServiceAmqpPlainTextCompressed

`func (o *MsgVpn) HasServiceAmqpPlainTextCompressed() bool`

HasServiceAmqpPlainTextCompressed returns a boolean if a field has been set.

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

### GetServiceAmqpPlainTextFailureReason

`func (o *MsgVpn) GetServiceAmqpPlainTextFailureReason() string`

GetServiceAmqpPlainTextFailureReason returns the ServiceAmqpPlainTextFailureReason field if non-nil, zero value otherwise.

### GetServiceAmqpPlainTextFailureReasonOk

`func (o *MsgVpn) GetServiceAmqpPlainTextFailureReasonOk() (*string, bool)`

GetServiceAmqpPlainTextFailureReasonOk returns a tuple with the ServiceAmqpPlainTextFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpPlainTextFailureReason

`func (o *MsgVpn) SetServiceAmqpPlainTextFailureReason(v string)`

SetServiceAmqpPlainTextFailureReason sets ServiceAmqpPlainTextFailureReason field to given value.

### HasServiceAmqpPlainTextFailureReason

`func (o *MsgVpn) HasServiceAmqpPlainTextFailureReason() bool`

HasServiceAmqpPlainTextFailureReason returns a boolean if a field has been set.

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

### GetServiceAmqpPlainTextUp

`func (o *MsgVpn) GetServiceAmqpPlainTextUp() bool`

GetServiceAmqpPlainTextUp returns the ServiceAmqpPlainTextUp field if non-nil, zero value otherwise.

### GetServiceAmqpPlainTextUpOk

`func (o *MsgVpn) GetServiceAmqpPlainTextUpOk() (*bool, bool)`

GetServiceAmqpPlainTextUpOk returns a tuple with the ServiceAmqpPlainTextUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpPlainTextUp

`func (o *MsgVpn) SetServiceAmqpPlainTextUp(v bool)`

SetServiceAmqpPlainTextUp sets ServiceAmqpPlainTextUp field to given value.

### HasServiceAmqpPlainTextUp

`func (o *MsgVpn) HasServiceAmqpPlainTextUp() bool`

HasServiceAmqpPlainTextUp returns a boolean if a field has been set.

### GetServiceAmqpTlsCompressed

`func (o *MsgVpn) GetServiceAmqpTlsCompressed() bool`

GetServiceAmqpTlsCompressed returns the ServiceAmqpTlsCompressed field if non-nil, zero value otherwise.

### GetServiceAmqpTlsCompressedOk

`func (o *MsgVpn) GetServiceAmqpTlsCompressedOk() (*bool, bool)`

GetServiceAmqpTlsCompressedOk returns a tuple with the ServiceAmqpTlsCompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpTlsCompressed

`func (o *MsgVpn) SetServiceAmqpTlsCompressed(v bool)`

SetServiceAmqpTlsCompressed sets ServiceAmqpTlsCompressed field to given value.

### HasServiceAmqpTlsCompressed

`func (o *MsgVpn) HasServiceAmqpTlsCompressed() bool`

HasServiceAmqpTlsCompressed returns a boolean if a field has been set.

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

### GetServiceAmqpTlsFailureReason

`func (o *MsgVpn) GetServiceAmqpTlsFailureReason() string`

GetServiceAmqpTlsFailureReason returns the ServiceAmqpTlsFailureReason field if non-nil, zero value otherwise.

### GetServiceAmqpTlsFailureReasonOk

`func (o *MsgVpn) GetServiceAmqpTlsFailureReasonOk() (*string, bool)`

GetServiceAmqpTlsFailureReasonOk returns a tuple with the ServiceAmqpTlsFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpTlsFailureReason

`func (o *MsgVpn) SetServiceAmqpTlsFailureReason(v string)`

SetServiceAmqpTlsFailureReason sets ServiceAmqpTlsFailureReason field to given value.

### HasServiceAmqpTlsFailureReason

`func (o *MsgVpn) HasServiceAmqpTlsFailureReason() bool`

HasServiceAmqpTlsFailureReason returns a boolean if a field has been set.

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

### GetServiceAmqpTlsUp

`func (o *MsgVpn) GetServiceAmqpTlsUp() bool`

GetServiceAmqpTlsUp returns the ServiceAmqpTlsUp field if non-nil, zero value otherwise.

### GetServiceAmqpTlsUpOk

`func (o *MsgVpn) GetServiceAmqpTlsUpOk() (*bool, bool)`

GetServiceAmqpTlsUpOk returns a tuple with the ServiceAmqpTlsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAmqpTlsUp

`func (o *MsgVpn) SetServiceAmqpTlsUp(v bool)`

SetServiceAmqpTlsUp sets ServiceAmqpTlsUp field to given value.

### HasServiceAmqpTlsUp

`func (o *MsgVpn) HasServiceAmqpTlsUp() bool`

HasServiceAmqpTlsUp returns a boolean if a field has been set.

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

### GetServiceMqttPlainTextCompressed

`func (o *MsgVpn) GetServiceMqttPlainTextCompressed() bool`

GetServiceMqttPlainTextCompressed returns the ServiceMqttPlainTextCompressed field if non-nil, zero value otherwise.

### GetServiceMqttPlainTextCompressedOk

`func (o *MsgVpn) GetServiceMqttPlainTextCompressedOk() (*bool, bool)`

GetServiceMqttPlainTextCompressedOk returns a tuple with the ServiceMqttPlainTextCompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttPlainTextCompressed

`func (o *MsgVpn) SetServiceMqttPlainTextCompressed(v bool)`

SetServiceMqttPlainTextCompressed sets ServiceMqttPlainTextCompressed field to given value.

### HasServiceMqttPlainTextCompressed

`func (o *MsgVpn) HasServiceMqttPlainTextCompressed() bool`

HasServiceMqttPlainTextCompressed returns a boolean if a field has been set.

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

### GetServiceMqttPlainTextFailureReason

`func (o *MsgVpn) GetServiceMqttPlainTextFailureReason() string`

GetServiceMqttPlainTextFailureReason returns the ServiceMqttPlainTextFailureReason field if non-nil, zero value otherwise.

### GetServiceMqttPlainTextFailureReasonOk

`func (o *MsgVpn) GetServiceMqttPlainTextFailureReasonOk() (*string, bool)`

GetServiceMqttPlainTextFailureReasonOk returns a tuple with the ServiceMqttPlainTextFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttPlainTextFailureReason

`func (o *MsgVpn) SetServiceMqttPlainTextFailureReason(v string)`

SetServiceMqttPlainTextFailureReason sets ServiceMqttPlainTextFailureReason field to given value.

### HasServiceMqttPlainTextFailureReason

`func (o *MsgVpn) HasServiceMqttPlainTextFailureReason() bool`

HasServiceMqttPlainTextFailureReason returns a boolean if a field has been set.

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

### GetServiceMqttPlainTextUp

`func (o *MsgVpn) GetServiceMqttPlainTextUp() bool`

GetServiceMqttPlainTextUp returns the ServiceMqttPlainTextUp field if non-nil, zero value otherwise.

### GetServiceMqttPlainTextUpOk

`func (o *MsgVpn) GetServiceMqttPlainTextUpOk() (*bool, bool)`

GetServiceMqttPlainTextUpOk returns a tuple with the ServiceMqttPlainTextUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttPlainTextUp

`func (o *MsgVpn) SetServiceMqttPlainTextUp(v bool)`

SetServiceMqttPlainTextUp sets ServiceMqttPlainTextUp field to given value.

### HasServiceMqttPlainTextUp

`func (o *MsgVpn) HasServiceMqttPlainTextUp() bool`

HasServiceMqttPlainTextUp returns a boolean if a field has been set.

### GetServiceMqttTlsCompressed

`func (o *MsgVpn) GetServiceMqttTlsCompressed() bool`

GetServiceMqttTlsCompressed returns the ServiceMqttTlsCompressed field if non-nil, zero value otherwise.

### GetServiceMqttTlsCompressedOk

`func (o *MsgVpn) GetServiceMqttTlsCompressedOk() (*bool, bool)`

GetServiceMqttTlsCompressedOk returns a tuple with the ServiceMqttTlsCompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttTlsCompressed

`func (o *MsgVpn) SetServiceMqttTlsCompressed(v bool)`

SetServiceMqttTlsCompressed sets ServiceMqttTlsCompressed field to given value.

### HasServiceMqttTlsCompressed

`func (o *MsgVpn) HasServiceMqttTlsCompressed() bool`

HasServiceMqttTlsCompressed returns a boolean if a field has been set.

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

### GetServiceMqttTlsFailureReason

`func (o *MsgVpn) GetServiceMqttTlsFailureReason() string`

GetServiceMqttTlsFailureReason returns the ServiceMqttTlsFailureReason field if non-nil, zero value otherwise.

### GetServiceMqttTlsFailureReasonOk

`func (o *MsgVpn) GetServiceMqttTlsFailureReasonOk() (*string, bool)`

GetServiceMqttTlsFailureReasonOk returns a tuple with the ServiceMqttTlsFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttTlsFailureReason

`func (o *MsgVpn) SetServiceMqttTlsFailureReason(v string)`

SetServiceMqttTlsFailureReason sets ServiceMqttTlsFailureReason field to given value.

### HasServiceMqttTlsFailureReason

`func (o *MsgVpn) HasServiceMqttTlsFailureReason() bool`

HasServiceMqttTlsFailureReason returns a boolean if a field has been set.

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

### GetServiceMqttTlsUp

`func (o *MsgVpn) GetServiceMqttTlsUp() bool`

GetServiceMqttTlsUp returns the ServiceMqttTlsUp field if non-nil, zero value otherwise.

### GetServiceMqttTlsUpOk

`func (o *MsgVpn) GetServiceMqttTlsUpOk() (*bool, bool)`

GetServiceMqttTlsUpOk returns a tuple with the ServiceMqttTlsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttTlsUp

`func (o *MsgVpn) SetServiceMqttTlsUp(v bool)`

SetServiceMqttTlsUp sets ServiceMqttTlsUp field to given value.

### HasServiceMqttTlsUp

`func (o *MsgVpn) HasServiceMqttTlsUp() bool`

HasServiceMqttTlsUp returns a boolean if a field has been set.

### GetServiceMqttTlsWebSocketCompressed

`func (o *MsgVpn) GetServiceMqttTlsWebSocketCompressed() bool`

GetServiceMqttTlsWebSocketCompressed returns the ServiceMqttTlsWebSocketCompressed field if non-nil, zero value otherwise.

### GetServiceMqttTlsWebSocketCompressedOk

`func (o *MsgVpn) GetServiceMqttTlsWebSocketCompressedOk() (*bool, bool)`

GetServiceMqttTlsWebSocketCompressedOk returns a tuple with the ServiceMqttTlsWebSocketCompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttTlsWebSocketCompressed

`func (o *MsgVpn) SetServiceMqttTlsWebSocketCompressed(v bool)`

SetServiceMqttTlsWebSocketCompressed sets ServiceMqttTlsWebSocketCompressed field to given value.

### HasServiceMqttTlsWebSocketCompressed

`func (o *MsgVpn) HasServiceMqttTlsWebSocketCompressed() bool`

HasServiceMqttTlsWebSocketCompressed returns a boolean if a field has been set.

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

### GetServiceMqttTlsWebSocketFailureReason

`func (o *MsgVpn) GetServiceMqttTlsWebSocketFailureReason() string`

GetServiceMqttTlsWebSocketFailureReason returns the ServiceMqttTlsWebSocketFailureReason field if non-nil, zero value otherwise.

### GetServiceMqttTlsWebSocketFailureReasonOk

`func (o *MsgVpn) GetServiceMqttTlsWebSocketFailureReasonOk() (*string, bool)`

GetServiceMqttTlsWebSocketFailureReasonOk returns a tuple with the ServiceMqttTlsWebSocketFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttTlsWebSocketFailureReason

`func (o *MsgVpn) SetServiceMqttTlsWebSocketFailureReason(v string)`

SetServiceMqttTlsWebSocketFailureReason sets ServiceMqttTlsWebSocketFailureReason field to given value.

### HasServiceMqttTlsWebSocketFailureReason

`func (o *MsgVpn) HasServiceMqttTlsWebSocketFailureReason() bool`

HasServiceMqttTlsWebSocketFailureReason returns a boolean if a field has been set.

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

### GetServiceMqttTlsWebSocketUp

`func (o *MsgVpn) GetServiceMqttTlsWebSocketUp() bool`

GetServiceMqttTlsWebSocketUp returns the ServiceMqttTlsWebSocketUp field if non-nil, zero value otherwise.

### GetServiceMqttTlsWebSocketUpOk

`func (o *MsgVpn) GetServiceMqttTlsWebSocketUpOk() (*bool, bool)`

GetServiceMqttTlsWebSocketUpOk returns a tuple with the ServiceMqttTlsWebSocketUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttTlsWebSocketUp

`func (o *MsgVpn) SetServiceMqttTlsWebSocketUp(v bool)`

SetServiceMqttTlsWebSocketUp sets ServiceMqttTlsWebSocketUp field to given value.

### HasServiceMqttTlsWebSocketUp

`func (o *MsgVpn) HasServiceMqttTlsWebSocketUp() bool`

HasServiceMqttTlsWebSocketUp returns a boolean if a field has been set.

### GetServiceMqttWebSocketCompressed

`func (o *MsgVpn) GetServiceMqttWebSocketCompressed() bool`

GetServiceMqttWebSocketCompressed returns the ServiceMqttWebSocketCompressed field if non-nil, zero value otherwise.

### GetServiceMqttWebSocketCompressedOk

`func (o *MsgVpn) GetServiceMqttWebSocketCompressedOk() (*bool, bool)`

GetServiceMqttWebSocketCompressedOk returns a tuple with the ServiceMqttWebSocketCompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttWebSocketCompressed

`func (o *MsgVpn) SetServiceMqttWebSocketCompressed(v bool)`

SetServiceMqttWebSocketCompressed sets ServiceMqttWebSocketCompressed field to given value.

### HasServiceMqttWebSocketCompressed

`func (o *MsgVpn) HasServiceMqttWebSocketCompressed() bool`

HasServiceMqttWebSocketCompressed returns a boolean if a field has been set.

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

### GetServiceMqttWebSocketFailureReason

`func (o *MsgVpn) GetServiceMqttWebSocketFailureReason() string`

GetServiceMqttWebSocketFailureReason returns the ServiceMqttWebSocketFailureReason field if non-nil, zero value otherwise.

### GetServiceMqttWebSocketFailureReasonOk

`func (o *MsgVpn) GetServiceMqttWebSocketFailureReasonOk() (*string, bool)`

GetServiceMqttWebSocketFailureReasonOk returns a tuple with the ServiceMqttWebSocketFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttWebSocketFailureReason

`func (o *MsgVpn) SetServiceMqttWebSocketFailureReason(v string)`

SetServiceMqttWebSocketFailureReason sets ServiceMqttWebSocketFailureReason field to given value.

### HasServiceMqttWebSocketFailureReason

`func (o *MsgVpn) HasServiceMqttWebSocketFailureReason() bool`

HasServiceMqttWebSocketFailureReason returns a boolean if a field has been set.

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

### GetServiceMqttWebSocketUp

`func (o *MsgVpn) GetServiceMqttWebSocketUp() bool`

GetServiceMqttWebSocketUp returns the ServiceMqttWebSocketUp field if non-nil, zero value otherwise.

### GetServiceMqttWebSocketUpOk

`func (o *MsgVpn) GetServiceMqttWebSocketUpOk() (*bool, bool)`

GetServiceMqttWebSocketUpOk returns a tuple with the ServiceMqttWebSocketUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMqttWebSocketUp

`func (o *MsgVpn) SetServiceMqttWebSocketUp(v bool)`

SetServiceMqttWebSocketUp sets ServiceMqttWebSocketUp field to given value.

### HasServiceMqttWebSocketUp

`func (o *MsgVpn) HasServiceMqttWebSocketUp() bool`

HasServiceMqttWebSocketUp returns a boolean if a field has been set.

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

### GetServiceRestIncomingPlainTextCompressed

`func (o *MsgVpn) GetServiceRestIncomingPlainTextCompressed() bool`

GetServiceRestIncomingPlainTextCompressed returns the ServiceRestIncomingPlainTextCompressed field if non-nil, zero value otherwise.

### GetServiceRestIncomingPlainTextCompressedOk

`func (o *MsgVpn) GetServiceRestIncomingPlainTextCompressedOk() (*bool, bool)`

GetServiceRestIncomingPlainTextCompressedOk returns a tuple with the ServiceRestIncomingPlainTextCompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingPlainTextCompressed

`func (o *MsgVpn) SetServiceRestIncomingPlainTextCompressed(v bool)`

SetServiceRestIncomingPlainTextCompressed sets ServiceRestIncomingPlainTextCompressed field to given value.

### HasServiceRestIncomingPlainTextCompressed

`func (o *MsgVpn) HasServiceRestIncomingPlainTextCompressed() bool`

HasServiceRestIncomingPlainTextCompressed returns a boolean if a field has been set.

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

### GetServiceRestIncomingPlainTextFailureReason

`func (o *MsgVpn) GetServiceRestIncomingPlainTextFailureReason() string`

GetServiceRestIncomingPlainTextFailureReason returns the ServiceRestIncomingPlainTextFailureReason field if non-nil, zero value otherwise.

### GetServiceRestIncomingPlainTextFailureReasonOk

`func (o *MsgVpn) GetServiceRestIncomingPlainTextFailureReasonOk() (*string, bool)`

GetServiceRestIncomingPlainTextFailureReasonOk returns a tuple with the ServiceRestIncomingPlainTextFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingPlainTextFailureReason

`func (o *MsgVpn) SetServiceRestIncomingPlainTextFailureReason(v string)`

SetServiceRestIncomingPlainTextFailureReason sets ServiceRestIncomingPlainTextFailureReason field to given value.

### HasServiceRestIncomingPlainTextFailureReason

`func (o *MsgVpn) HasServiceRestIncomingPlainTextFailureReason() bool`

HasServiceRestIncomingPlainTextFailureReason returns a boolean if a field has been set.

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

### GetServiceRestIncomingPlainTextUp

`func (o *MsgVpn) GetServiceRestIncomingPlainTextUp() bool`

GetServiceRestIncomingPlainTextUp returns the ServiceRestIncomingPlainTextUp field if non-nil, zero value otherwise.

### GetServiceRestIncomingPlainTextUpOk

`func (o *MsgVpn) GetServiceRestIncomingPlainTextUpOk() (*bool, bool)`

GetServiceRestIncomingPlainTextUpOk returns a tuple with the ServiceRestIncomingPlainTextUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingPlainTextUp

`func (o *MsgVpn) SetServiceRestIncomingPlainTextUp(v bool)`

SetServiceRestIncomingPlainTextUp sets ServiceRestIncomingPlainTextUp field to given value.

### HasServiceRestIncomingPlainTextUp

`func (o *MsgVpn) HasServiceRestIncomingPlainTextUp() bool`

HasServiceRestIncomingPlainTextUp returns a boolean if a field has been set.

### GetServiceRestIncomingTlsCompressed

`func (o *MsgVpn) GetServiceRestIncomingTlsCompressed() bool`

GetServiceRestIncomingTlsCompressed returns the ServiceRestIncomingTlsCompressed field if non-nil, zero value otherwise.

### GetServiceRestIncomingTlsCompressedOk

`func (o *MsgVpn) GetServiceRestIncomingTlsCompressedOk() (*bool, bool)`

GetServiceRestIncomingTlsCompressedOk returns a tuple with the ServiceRestIncomingTlsCompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingTlsCompressed

`func (o *MsgVpn) SetServiceRestIncomingTlsCompressed(v bool)`

SetServiceRestIncomingTlsCompressed sets ServiceRestIncomingTlsCompressed field to given value.

### HasServiceRestIncomingTlsCompressed

`func (o *MsgVpn) HasServiceRestIncomingTlsCompressed() bool`

HasServiceRestIncomingTlsCompressed returns a boolean if a field has been set.

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

### GetServiceRestIncomingTlsFailureReason

`func (o *MsgVpn) GetServiceRestIncomingTlsFailureReason() string`

GetServiceRestIncomingTlsFailureReason returns the ServiceRestIncomingTlsFailureReason field if non-nil, zero value otherwise.

### GetServiceRestIncomingTlsFailureReasonOk

`func (o *MsgVpn) GetServiceRestIncomingTlsFailureReasonOk() (*string, bool)`

GetServiceRestIncomingTlsFailureReasonOk returns a tuple with the ServiceRestIncomingTlsFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingTlsFailureReason

`func (o *MsgVpn) SetServiceRestIncomingTlsFailureReason(v string)`

SetServiceRestIncomingTlsFailureReason sets ServiceRestIncomingTlsFailureReason field to given value.

### HasServiceRestIncomingTlsFailureReason

`func (o *MsgVpn) HasServiceRestIncomingTlsFailureReason() bool`

HasServiceRestIncomingTlsFailureReason returns a boolean if a field has been set.

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

### GetServiceRestIncomingTlsUp

`func (o *MsgVpn) GetServiceRestIncomingTlsUp() bool`

GetServiceRestIncomingTlsUp returns the ServiceRestIncomingTlsUp field if non-nil, zero value otherwise.

### GetServiceRestIncomingTlsUpOk

`func (o *MsgVpn) GetServiceRestIncomingTlsUpOk() (*bool, bool)`

GetServiceRestIncomingTlsUpOk returns a tuple with the ServiceRestIncomingTlsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRestIncomingTlsUp

`func (o *MsgVpn) SetServiceRestIncomingTlsUp(v bool)`

SetServiceRestIncomingTlsUp sets ServiceRestIncomingTlsUp field to given value.

### HasServiceRestIncomingTlsUp

`func (o *MsgVpn) HasServiceRestIncomingTlsUp() bool`

HasServiceRestIncomingTlsUp returns a boolean if a field has been set.

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

### GetServiceSmfPlainTextFailureReason

`func (o *MsgVpn) GetServiceSmfPlainTextFailureReason() string`

GetServiceSmfPlainTextFailureReason returns the ServiceSmfPlainTextFailureReason field if non-nil, zero value otherwise.

### GetServiceSmfPlainTextFailureReasonOk

`func (o *MsgVpn) GetServiceSmfPlainTextFailureReasonOk() (*string, bool)`

GetServiceSmfPlainTextFailureReasonOk returns a tuple with the ServiceSmfPlainTextFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfPlainTextFailureReason

`func (o *MsgVpn) SetServiceSmfPlainTextFailureReason(v string)`

SetServiceSmfPlainTextFailureReason sets ServiceSmfPlainTextFailureReason field to given value.

### HasServiceSmfPlainTextFailureReason

`func (o *MsgVpn) HasServiceSmfPlainTextFailureReason() bool`

HasServiceSmfPlainTextFailureReason returns a boolean if a field has been set.

### GetServiceSmfPlainTextUp

`func (o *MsgVpn) GetServiceSmfPlainTextUp() bool`

GetServiceSmfPlainTextUp returns the ServiceSmfPlainTextUp field if non-nil, zero value otherwise.

### GetServiceSmfPlainTextUpOk

`func (o *MsgVpn) GetServiceSmfPlainTextUpOk() (*bool, bool)`

GetServiceSmfPlainTextUpOk returns a tuple with the ServiceSmfPlainTextUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfPlainTextUp

`func (o *MsgVpn) SetServiceSmfPlainTextUp(v bool)`

SetServiceSmfPlainTextUp sets ServiceSmfPlainTextUp field to given value.

### HasServiceSmfPlainTextUp

`func (o *MsgVpn) HasServiceSmfPlainTextUp() bool`

HasServiceSmfPlainTextUp returns a boolean if a field has been set.

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

### GetServiceSmfTlsFailureReason

`func (o *MsgVpn) GetServiceSmfTlsFailureReason() string`

GetServiceSmfTlsFailureReason returns the ServiceSmfTlsFailureReason field if non-nil, zero value otherwise.

### GetServiceSmfTlsFailureReasonOk

`func (o *MsgVpn) GetServiceSmfTlsFailureReasonOk() (*string, bool)`

GetServiceSmfTlsFailureReasonOk returns a tuple with the ServiceSmfTlsFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfTlsFailureReason

`func (o *MsgVpn) SetServiceSmfTlsFailureReason(v string)`

SetServiceSmfTlsFailureReason sets ServiceSmfTlsFailureReason field to given value.

### HasServiceSmfTlsFailureReason

`func (o *MsgVpn) HasServiceSmfTlsFailureReason() bool`

HasServiceSmfTlsFailureReason returns a boolean if a field has been set.

### GetServiceSmfTlsUp

`func (o *MsgVpn) GetServiceSmfTlsUp() bool`

GetServiceSmfTlsUp returns the ServiceSmfTlsUp field if non-nil, zero value otherwise.

### GetServiceSmfTlsUpOk

`func (o *MsgVpn) GetServiceSmfTlsUpOk() (*bool, bool)`

GetServiceSmfTlsUpOk returns a tuple with the ServiceSmfTlsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSmfTlsUp

`func (o *MsgVpn) SetServiceSmfTlsUp(v bool)`

SetServiceSmfTlsUp sets ServiceSmfTlsUp field to given value.

### HasServiceSmfTlsUp

`func (o *MsgVpn) HasServiceSmfTlsUp() bool`

HasServiceSmfTlsUp returns a boolean if a field has been set.

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

### GetServiceWebPlainTextFailureReason

`func (o *MsgVpn) GetServiceWebPlainTextFailureReason() string`

GetServiceWebPlainTextFailureReason returns the ServiceWebPlainTextFailureReason field if non-nil, zero value otherwise.

### GetServiceWebPlainTextFailureReasonOk

`func (o *MsgVpn) GetServiceWebPlainTextFailureReasonOk() (*string, bool)`

GetServiceWebPlainTextFailureReasonOk returns a tuple with the ServiceWebPlainTextFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebPlainTextFailureReason

`func (o *MsgVpn) SetServiceWebPlainTextFailureReason(v string)`

SetServiceWebPlainTextFailureReason sets ServiceWebPlainTextFailureReason field to given value.

### HasServiceWebPlainTextFailureReason

`func (o *MsgVpn) HasServiceWebPlainTextFailureReason() bool`

HasServiceWebPlainTextFailureReason returns a boolean if a field has been set.

### GetServiceWebPlainTextUp

`func (o *MsgVpn) GetServiceWebPlainTextUp() bool`

GetServiceWebPlainTextUp returns the ServiceWebPlainTextUp field if non-nil, zero value otherwise.

### GetServiceWebPlainTextUpOk

`func (o *MsgVpn) GetServiceWebPlainTextUpOk() (*bool, bool)`

GetServiceWebPlainTextUpOk returns a tuple with the ServiceWebPlainTextUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebPlainTextUp

`func (o *MsgVpn) SetServiceWebPlainTextUp(v bool)`

SetServiceWebPlainTextUp sets ServiceWebPlainTextUp field to given value.

### HasServiceWebPlainTextUp

`func (o *MsgVpn) HasServiceWebPlainTextUp() bool`

HasServiceWebPlainTextUp returns a boolean if a field has been set.

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

### GetServiceWebTlsFailureReason

`func (o *MsgVpn) GetServiceWebTlsFailureReason() string`

GetServiceWebTlsFailureReason returns the ServiceWebTlsFailureReason field if non-nil, zero value otherwise.

### GetServiceWebTlsFailureReasonOk

`func (o *MsgVpn) GetServiceWebTlsFailureReasonOk() (*string, bool)`

GetServiceWebTlsFailureReasonOk returns a tuple with the ServiceWebTlsFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebTlsFailureReason

`func (o *MsgVpn) SetServiceWebTlsFailureReason(v string)`

SetServiceWebTlsFailureReason sets ServiceWebTlsFailureReason field to given value.

### HasServiceWebTlsFailureReason

`func (o *MsgVpn) HasServiceWebTlsFailureReason() bool`

HasServiceWebTlsFailureReason returns a boolean if a field has been set.

### GetServiceWebTlsUp

`func (o *MsgVpn) GetServiceWebTlsUp() bool`

GetServiceWebTlsUp returns the ServiceWebTlsUp field if non-nil, zero value otherwise.

### GetServiceWebTlsUpOk

`func (o *MsgVpn) GetServiceWebTlsUpOk() (*bool, bool)`

GetServiceWebTlsUpOk returns a tuple with the ServiceWebTlsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWebTlsUp

`func (o *MsgVpn) SetServiceWebTlsUp(v bool)`

SetServiceWebTlsUp sets ServiceWebTlsUp field to given value.

### HasServiceWebTlsUp

`func (o *MsgVpn) HasServiceWebTlsUp() bool`

HasServiceWebTlsUp returns a boolean if a field has been set.

### GetState

`func (o *MsgVpn) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MsgVpn) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MsgVpn) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MsgVpn) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubscriptionExportProgress

`func (o *MsgVpn) GetSubscriptionExportProgress() int64`

GetSubscriptionExportProgress returns the SubscriptionExportProgress field if non-nil, zero value otherwise.

### GetSubscriptionExportProgressOk

`func (o *MsgVpn) GetSubscriptionExportProgressOk() (*int64, bool)`

GetSubscriptionExportProgressOk returns a tuple with the SubscriptionExportProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExportProgress

`func (o *MsgVpn) SetSubscriptionExportProgress(v int64)`

SetSubscriptionExportProgress sets SubscriptionExportProgress field to given value.

### HasSubscriptionExportProgress

`func (o *MsgVpn) HasSubscriptionExportProgress() bool`

HasSubscriptionExportProgress returns a boolean if a field has been set.

### GetSystemManager

`func (o *MsgVpn) GetSystemManager() bool`

GetSystemManager returns the SystemManager field if non-nil, zero value otherwise.

### GetSystemManagerOk

`func (o *MsgVpn) GetSystemManagerOk() (*bool, bool)`

GetSystemManagerOk returns a tuple with the SystemManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemManager

`func (o *MsgVpn) SetSystemManager(v bool)`

SetSystemManager sets SystemManager field to given value.

### HasSystemManager

`func (o *MsgVpn) HasSystemManager() bool`

HasSystemManager returns a boolean if a field has been set.

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

### GetTlsAverageRxByteRate

`func (o *MsgVpn) GetTlsAverageRxByteRate() int64`

GetTlsAverageRxByteRate returns the TlsAverageRxByteRate field if non-nil, zero value otherwise.

### GetTlsAverageRxByteRateOk

`func (o *MsgVpn) GetTlsAverageRxByteRateOk() (*int64, bool)`

GetTlsAverageRxByteRateOk returns a tuple with the TlsAverageRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAverageRxByteRate

`func (o *MsgVpn) SetTlsAverageRxByteRate(v int64)`

SetTlsAverageRxByteRate sets TlsAverageRxByteRate field to given value.

### HasTlsAverageRxByteRate

`func (o *MsgVpn) HasTlsAverageRxByteRate() bool`

HasTlsAverageRxByteRate returns a boolean if a field has been set.

### GetTlsAverageTxByteRate

`func (o *MsgVpn) GetTlsAverageTxByteRate() int64`

GetTlsAverageTxByteRate returns the TlsAverageTxByteRate field if non-nil, zero value otherwise.

### GetTlsAverageTxByteRateOk

`func (o *MsgVpn) GetTlsAverageTxByteRateOk() (*int64, bool)`

GetTlsAverageTxByteRateOk returns a tuple with the TlsAverageTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAverageTxByteRate

`func (o *MsgVpn) SetTlsAverageTxByteRate(v int64)`

SetTlsAverageTxByteRate sets TlsAverageTxByteRate field to given value.

### HasTlsAverageTxByteRate

`func (o *MsgVpn) HasTlsAverageTxByteRate() bool`

HasTlsAverageTxByteRate returns a boolean if a field has been set.

### GetTlsRxByteCount

`func (o *MsgVpn) GetTlsRxByteCount() int64`

GetTlsRxByteCount returns the TlsRxByteCount field if non-nil, zero value otherwise.

### GetTlsRxByteCountOk

`func (o *MsgVpn) GetTlsRxByteCountOk() (*int64, bool)`

GetTlsRxByteCountOk returns a tuple with the TlsRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsRxByteCount

`func (o *MsgVpn) SetTlsRxByteCount(v int64)`

SetTlsRxByteCount sets TlsRxByteCount field to given value.

### HasTlsRxByteCount

`func (o *MsgVpn) HasTlsRxByteCount() bool`

HasTlsRxByteCount returns a boolean if a field has been set.

### GetTlsRxByteRate

`func (o *MsgVpn) GetTlsRxByteRate() int64`

GetTlsRxByteRate returns the TlsRxByteRate field if non-nil, zero value otherwise.

### GetTlsRxByteRateOk

`func (o *MsgVpn) GetTlsRxByteRateOk() (*int64, bool)`

GetTlsRxByteRateOk returns a tuple with the TlsRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsRxByteRate

`func (o *MsgVpn) SetTlsRxByteRate(v int64)`

SetTlsRxByteRate sets TlsRxByteRate field to given value.

### HasTlsRxByteRate

`func (o *MsgVpn) HasTlsRxByteRate() bool`

HasTlsRxByteRate returns a boolean if a field has been set.

### GetTlsTxByteCount

`func (o *MsgVpn) GetTlsTxByteCount() int64`

GetTlsTxByteCount returns the TlsTxByteCount field if non-nil, zero value otherwise.

### GetTlsTxByteCountOk

`func (o *MsgVpn) GetTlsTxByteCountOk() (*int64, bool)`

GetTlsTxByteCountOk returns a tuple with the TlsTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTxByteCount

`func (o *MsgVpn) SetTlsTxByteCount(v int64)`

SetTlsTxByteCount sets TlsTxByteCount field to given value.

### HasTlsTxByteCount

`func (o *MsgVpn) HasTlsTxByteCount() bool`

HasTlsTxByteCount returns a boolean if a field has been set.

### GetTlsTxByteRate

`func (o *MsgVpn) GetTlsTxByteRate() int64`

GetTlsTxByteRate returns the TlsTxByteRate field if non-nil, zero value otherwise.

### GetTlsTxByteRateOk

`func (o *MsgVpn) GetTlsTxByteRateOk() (*int64, bool)`

GetTlsTxByteRateOk returns a tuple with the TlsTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTxByteRate

`func (o *MsgVpn) SetTlsTxByteRate(v int64)`

SetTlsTxByteRate sets TlsTxByteRate field to given value.

### HasTlsTxByteRate

`func (o *MsgVpn) HasTlsTxByteRate() bool`

HasTlsTxByteRate returns a boolean if a field has been set.

### GetTxByteCount

`func (o *MsgVpn) GetTxByteCount() int64`

GetTxByteCount returns the TxByteCount field if non-nil, zero value otherwise.

### GetTxByteCountOk

`func (o *MsgVpn) GetTxByteCountOk() (*int64, bool)`

GetTxByteCountOk returns a tuple with the TxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteCount

`func (o *MsgVpn) SetTxByteCount(v int64)`

SetTxByteCount sets TxByteCount field to given value.

### HasTxByteCount

`func (o *MsgVpn) HasTxByteCount() bool`

HasTxByteCount returns a boolean if a field has been set.

### GetTxByteRate

`func (o *MsgVpn) GetTxByteRate() int64`

GetTxByteRate returns the TxByteRate field if non-nil, zero value otherwise.

### GetTxByteRateOk

`func (o *MsgVpn) GetTxByteRateOk() (*int64, bool)`

GetTxByteRateOk returns a tuple with the TxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteRate

`func (o *MsgVpn) SetTxByteRate(v int64)`

SetTxByteRate sets TxByteRate field to given value.

### HasTxByteRate

`func (o *MsgVpn) HasTxByteRate() bool`

HasTxByteRate returns a boolean if a field has been set.

### GetTxCompressedByteCount

`func (o *MsgVpn) GetTxCompressedByteCount() int64`

GetTxCompressedByteCount returns the TxCompressedByteCount field if non-nil, zero value otherwise.

### GetTxCompressedByteCountOk

`func (o *MsgVpn) GetTxCompressedByteCountOk() (*int64, bool)`

GetTxCompressedByteCountOk returns a tuple with the TxCompressedByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCompressedByteCount

`func (o *MsgVpn) SetTxCompressedByteCount(v int64)`

SetTxCompressedByteCount sets TxCompressedByteCount field to given value.

### HasTxCompressedByteCount

`func (o *MsgVpn) HasTxCompressedByteCount() bool`

HasTxCompressedByteCount returns a boolean if a field has been set.

### GetTxCompressedByteRate

`func (o *MsgVpn) GetTxCompressedByteRate() int64`

GetTxCompressedByteRate returns the TxCompressedByteRate field if non-nil, zero value otherwise.

### GetTxCompressedByteRateOk

`func (o *MsgVpn) GetTxCompressedByteRateOk() (*int64, bool)`

GetTxCompressedByteRateOk returns a tuple with the TxCompressedByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCompressedByteRate

`func (o *MsgVpn) SetTxCompressedByteRate(v int64)`

SetTxCompressedByteRate sets TxCompressedByteRate field to given value.

### HasTxCompressedByteRate

`func (o *MsgVpn) HasTxCompressedByteRate() bool`

HasTxCompressedByteRate returns a boolean if a field has been set.

### GetTxCompressionRatio

`func (o *MsgVpn) GetTxCompressionRatio() string`

GetTxCompressionRatio returns the TxCompressionRatio field if non-nil, zero value otherwise.

### GetTxCompressionRatioOk

`func (o *MsgVpn) GetTxCompressionRatioOk() (*string, bool)`

GetTxCompressionRatioOk returns a tuple with the TxCompressionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxCompressionRatio

`func (o *MsgVpn) SetTxCompressionRatio(v string)`

SetTxCompressionRatio sets TxCompressionRatio field to given value.

### HasTxCompressionRatio

`func (o *MsgVpn) HasTxCompressionRatio() bool`

HasTxCompressionRatio returns a boolean if a field has been set.

### GetTxMsgCount

`func (o *MsgVpn) GetTxMsgCount() int64`

GetTxMsgCount returns the TxMsgCount field if non-nil, zero value otherwise.

### GetTxMsgCountOk

`func (o *MsgVpn) GetTxMsgCountOk() (*int64, bool)`

GetTxMsgCountOk returns a tuple with the TxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgCount

`func (o *MsgVpn) SetTxMsgCount(v int64)`

SetTxMsgCount sets TxMsgCount field to given value.

### HasTxMsgCount

`func (o *MsgVpn) HasTxMsgCount() bool`

HasTxMsgCount returns a boolean if a field has been set.

### GetTxMsgRate

`func (o *MsgVpn) GetTxMsgRate() int64`

GetTxMsgRate returns the TxMsgRate field if non-nil, zero value otherwise.

### GetTxMsgRateOk

`func (o *MsgVpn) GetTxMsgRateOk() (*int64, bool)`

GetTxMsgRateOk returns a tuple with the TxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgRate

`func (o *MsgVpn) SetTxMsgRate(v int64)`

SetTxMsgRate sets TxMsgRate field to given value.

### HasTxMsgRate

`func (o *MsgVpn) HasTxMsgRate() bool`

HasTxMsgRate returns a boolean if a field has been set.

### GetTxUncompressedByteCount

`func (o *MsgVpn) GetTxUncompressedByteCount() int64`

GetTxUncompressedByteCount returns the TxUncompressedByteCount field if non-nil, zero value otherwise.

### GetTxUncompressedByteCountOk

`func (o *MsgVpn) GetTxUncompressedByteCountOk() (*int64, bool)`

GetTxUncompressedByteCountOk returns a tuple with the TxUncompressedByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxUncompressedByteCount

`func (o *MsgVpn) SetTxUncompressedByteCount(v int64)`

SetTxUncompressedByteCount sets TxUncompressedByteCount field to given value.

### HasTxUncompressedByteCount

`func (o *MsgVpn) HasTxUncompressedByteCount() bool`

HasTxUncompressedByteCount returns a boolean if a field has been set.

### GetTxUncompressedByteRate

`func (o *MsgVpn) GetTxUncompressedByteRate() int64`

GetTxUncompressedByteRate returns the TxUncompressedByteRate field if non-nil, zero value otherwise.

### GetTxUncompressedByteRateOk

`func (o *MsgVpn) GetTxUncompressedByteRateOk() (*int64, bool)`

GetTxUncompressedByteRateOk returns a tuple with the TxUncompressedByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxUncompressedByteRate

`func (o *MsgVpn) SetTxUncompressedByteRate(v int64)`

SetTxUncompressedByteRate sets TxUncompressedByteRate field to given value.

### HasTxUncompressedByteRate

`func (o *MsgVpn) HasTxUncompressedByteRate() bool`

HasTxUncompressedByteRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


