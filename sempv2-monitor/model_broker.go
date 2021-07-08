/*
 * SEMP (Solace Element Management Protocol)
 *
 * SEMP (starting in `v2`, see note 1) is a RESTful API for configuring, monitoring, and administering a Solace PubSub+ broker.  SEMP uses URIs to address manageable **resources** of the Solace PubSub+ broker. Resources are individual **objects**, **collections** of objects, or (exclusively in the action API) **actions**. This document applies to the following API:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Monitoring|/SEMP/v2/monitor|Querying operational parameters|See note 2    The following APIs are also available:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Action|/SEMP/v2/action|Performing actions|See note 2 Configuration|/SEMP/v2/config|Reading and writing config state|See note 2    Resources are always nouns, with individual objects being singular and collections being plural.  Objects within a collection are identified by an `obj-id`, which follows the collection name with the form `collection-name/obj-id`.  Actions within an object are identified by an `action-id`, which follows the object name with the form `obj-id/action-id`.  Some examples:  ``` /SEMP/v2/config/msgVpns                        ; MsgVpn collection /SEMP/v2/config/msgVpns/a                      ; MsgVpn object named \"a\" /SEMP/v2/config/msgVpns/a/queues               ; Queue collection in MsgVpn \"a\" /SEMP/v2/config/msgVpns/a/queues/b             ; Queue object named \"b\" in MsgVpn \"a\" /SEMP/v2/action/msgVpns/a/queues/b/startReplay ; Action that starts a replay on Queue \"b\" in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients             ; Client collection in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients/c           ; Client object named \"c\" in MsgVpn \"a\" ```  ## Collection Resources  Collections are unordered lists of objects (unless described as otherwise), and are described by JSON arrays. Each item in the array represents an object in the same manner as the individual object would normally be represented. In the configuration API, the creation of a new object is done through its collection resource.  ## Object and Action Resources  Objects are composed of attributes, actions, collections, and other objects. They are described by JSON objects as name/value pairs. The collections and actions of an object are not contained directly in the object's JSON content; rather the content includes an attribute containing a URI which points to the collections and actions. These contained resources must be managed through this URI. At a minimum, every object has one or more identifying attributes, and its own `uri` attribute which contains the URI pointing to itself.  Actions are also composed of attributes, and are described by JSON objects as name/value pairs. Unlike objects, however, they are not members of a collection and cannot be retrieved, only performed. Actions only exist in the action API.  Attributes in an object or action may have any combination of the following properties:   Property|Meaning|Comments :---|:---|:--- Identifying|Attribute is involved in unique identification of the object, and appears in its URI| Required|Attribute must be provided in the request| Read-Only|Attribute can only be read, not written.|See note 3 Write-Only|Attribute can only be written, not read, unless the attribute is also opaque|See the documentation for the opaque property Requires-Disable|Attribute can only be changed when object is disabled| Deprecated|Attribute is deprecated, and will disappear in the next SEMP version| Opaque|Attribute can be set or retrieved in opaque form when the `opaquePassword` query parameter is present|See the `opaquePassword` query parameter documentation    In some requests, certain attributes may only be provided in certain combinations with other attributes:   Relationship|Meaning :---|:--- Requires|Attribute may only be changed by a request if a particular attribute or combination of attributes is also provided in the request Conflicts|Attribute may only be provided in a request if a particular attribute or combination of attributes is not also provided in the request    In the monitoring API, any non-identifying attribute may not be returned in a GET.  ## HTTP Methods  The following HTTP methods manipulate resources in accordance with these general principles. Note that some methods are only used in certain APIs:   Method|Resource|Meaning|Request Body|Response Body|Missing Request Attributes :---|:---|:---|:---|:---|:--- POST|Collection|Create object|Initial attribute values|Object attributes and metadata|Set to default PUT|Object|Create or replace object (see note 5)|New attribute values|Object attributes and metadata|Set to default, with certain exceptions (see note 4) PUT|Action|Performs action|Action arguments|Action metadata|N/A PATCH|Object|Update object|New attribute values|Object attributes and metadata|unchanged DELETE|Object|Delete object|Empty|Object metadata|N/A GET|Object|Get object|Empty|Object attributes and metadata|N/A GET|Collection|Get collection|Empty|Object attributes and collection metadata|N/A    ## Common Query Parameters  The following are some common query parameters that are supported by many method/URI combinations. Individual URIs may document additional parameters. Note that multiple query parameters can be used together in a single URI, separated by the ampersand character. For example:  ``` ; Request for the MsgVpns collection using two hypothetical query parameters ; \"q1\" and \"q2\" with values \"val1\" and \"val2\" respectively /SEMP/v2/monitor/msgVpns?q1=val1&q2=val2 ```  ### select  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. Use this query parameter to limit the size of the returned data for each returned object, return only those fields that are desired, or exclude fields that are not desired.  The value of `select` is a comma-separated list of attribute names. If the list contains attribute names that are not prefaced by `-`, only those attributes are included in the response. If the list contains attribute names that are prefaced by `-`, those attributes are excluded from the response. If the list contains both types, then the difference of the first set of attributes and the second set of attributes is returned. If the list is empty (i.e. `select=`), no attributes are returned.  All attributes that are prefaced by `-` must follow all attributes that are not prefaced by `-`. In addition, each attribute name in the list must match at least one attribute in the object.  Names may include the `*` wildcard (zero or more characters). Nested attribute names are supported using periods (e.g. `parentName.childName`).  Some examples:  ``` ; List of all MsgVpn names /SEMP/v2/monitor/msgVpns?select=msgVpnName ; List of all MsgVpn and their attributes except for their names /SEMP/v2/monitor/msgVpns?select=-msgVpnName ; Authentication attributes of MsgVpn \"finance\" /SEMP/v2/monitor/msgVpns/finance?select=authentication* ; All attributes of MsgVpn \"finance\" except for authentication attributes /SEMP/v2/monitor/msgVpns/finance?select=-authentication* ; Access related attributes of Queue \"orderQ\" of MsgVpn \"finance\" /SEMP/v2/monitor/msgVpns/finance/queues/orderQ?select=owner,permission ```  ### where  Include in the response only objects where certain conditions are true. Use this query parameter to limit which objects are returned to those whose attribute values meet the given conditions.  The value of `where` is a comma-separated list of expressions. All expressions must be true for the object to be included in the response. Each expression takes the form:  ``` expression  = attribute-name OP value OP          = '==' | '!=' | '&lt;' | '&gt;' | '&lt;=' | '&gt;=' ```  `value` may be a number, string, `true`, or `false`, as appropriate for the type of `attribute-name`. Greater-than and less-than comparisons only work for numbers. A `*` in a string `value` is interpreted as a wildcard (zero or more characters). Some examples:  ``` ; Only enabled MsgVpns /SEMP/v2/monitor/msgVpns?where=enabled==true ; Only MsgVpns using basic non-LDAP authentication /SEMP/v2/monitor/msgVpns?where=authenticationBasicEnabled==true,authenticationBasicType!=ldap ; Only MsgVpns that allow more than 100 client connections /SEMP/v2/monitor/msgVpns?where=maxConnectionCount>100 ; Only MsgVpns with msgVpnName starting with \"B\": /SEMP/v2/monitor/msgVpns?where=msgVpnName==B* ```  ### count  Limit the count of objects in the response. This can be useful to limit the size of the response for large collections. The minimum value for `count` is `1` and the default is `10`. There is also a per-collection maximum value to limit request handling time. For example:  ``` ; Up to 25 MsgVpns /SEMP/v2/monitor/msgVpns?count=25 ```  ### cursor  The cursor, or position, for the next page of objects. Cursors are opaque data that should not be created or interpreted by SEMP clients, and should only be used as described below.  When a request is made for a collection and there may be additional objects available for retrieval that are not included in the initial response, the response will include a `cursorQuery` field containing a cursor. The value of this field can be specified in the `cursor` query parameter of a subsequent request to retrieve the next page of objects. For convenience, an appropriate URI is constructed automatically by the broker and included in the `nextPageUri` field of the response. This URI can be used directly to retrieve the next page of objects.  ### opaquePassword  Attributes with the opaque property are also write-only and so cannot normally be retrieved in a GET. However, when a password is provided in the `opaquePassword` query parameter, attributes with the opaque property are retrieved in a GET in opaque form, encrypted with this password. The query parameter can also be used on a POST, PATCH, or PUT to set opaque attributes using opaque attribute values retrieved in a GET, so long as:  1. the same password that was used to retrieve the opaque attribute values is provided; and  2. the broker to which the request is being sent has the same major and minor SEMP version as the broker that produced the opaque attribute values.  The password provided in the query parameter must be a minimum of 8 characters and a maximum of 128 characters.  The query parameter can only be used in the configuration API, and only over HTTPS.  ## Authentication  When a client makes its first SEMPv2 request, it must supply a username and password using HTTP Basic authentication.  If authentication is successful, the broker returns a cookie containing a session key. The client can omit the username and password from subsequent requests, because the broker now uses the session cookie for authentication instead. When the session expires or is deleted, the client must provide the username and password again, and the broker creates a new session.  There are a limited number of session slots available on the broker. The broker returns 529 No SEMP Session Available if it is not able to allocate a session. For this reason, all clients that use SEMPv2 should support cookies.  If certain attributes—such as a user's password—are changed, the broker automatically deletes the affected sessions. These attributes are documented below. However, changes in external user configuration data stored on a RADIUS or LDAP server do not trigger the broker to delete the associated session(s), therefore you must do this manually, if required.  A client can retrieve its current session information using the /about/user endpoint, delete its own session using the /about/user/logout endpoint, and manage all sessions using the /sessions endpoint.  ## Help  Visit [our website](https://solace.com) to learn more about Solace.  You can also download the SEMP API specifications by clicking [here](https://solace.com/downloads/).  If you need additional support, please contact us at [support@solace.com](mailto:support@solace.com).  ## Notes  Note|Description :---:|:--- 1|This specification defines SEMP starting in \"v2\", and not the original SEMP \"v1\" interface. Request and response formats between \"v1\" and \"v2\" are entirely incompatible, although both protocols share a common port configuration on the Solace PubSub+ broker. They are differentiated by the initial portion of the URI path, one of either \"/SEMP/\" or \"/SEMP/v2/\" 2|This API is partially implemented. Only a subset of all objects are available. 3|Read-only attributes may appear in POST and PUT/PATCH requests. However, if a read-only attribute is not marked as identifying, it will be ignored during a PUT/PATCH. 4|On a PUT, if the SEMP user is not authorized to modify the attribute, its value is left unchanged rather than set to default. In addition, the values of write-only attributes are not set to their defaults on a PUT, except in the following two cases: there is a mutual requires relationship with another non-write-only attribute, both attributes are absent from the request, and the non-write-only attribute is not currently set to its default value; or the attribute is also opaque and the `opaquePassword` query parameter is provided in the request. 5|On a PUT, if the object does not exist, it is created first.
 *
 * API version: 2.21
 * Contact: support@solace.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Broker struct {
	// The client certificate revocation checking mode used when a client authenticates with a client certificate. The allowed values and their meaning are:  <pre> \"none\" - Do not perform any certificate revocation checking. \"ocsp\" - Use the Open Certificate Status Protcol (OCSP) for certificate revocation checking. \"crl\" - Use Certificate Revocation Lists (CRL) for certificate revocation checking. \"ocsp-crl\" - Use OCSP first, but if OCSP fails to return an unambiguous result, then check via CRL. </pre>
	AuthClientCertRevocationCheckMode string `json:"authClientCertRevocationCheckMode,omitempty"`
	// The one minute average of the message rate received by the Broker, in bytes per second (B/sec). Available since 2.14.
	AverageRxByteRate int64 `json:"averageRxByteRate,omitempty"`
	// The one minute average of the compressed message rate received by the Broker, in bytes per second (B/sec). Available since 2.14.
	AverageRxCompressedByteRate int64 `json:"averageRxCompressedByteRate,omitempty"`
	// The one minute average of the message rate received by the Broker, in messages per second (msg/sec). Available since 2.14.
	AverageRxMsgRate int64 `json:"averageRxMsgRate,omitempty"`
	// The one minute average of the uncompressed message rate received by the Broker, in bytes per second (B/sec). Available since 2.14.
	AverageRxUncompressedByteRate int64 `json:"averageRxUncompressedByteRate,omitempty"`
	// The one minute average of the message rate transmitted by the Broker, in bytes per second (B/sec). Available since 2.14.
	AverageTxByteRate int64 `json:"averageTxByteRate,omitempty"`
	// The one minute average of the compressed message rate transmitted by the Broker, in bytes per second (B/sec). Available since 2.14.
	AverageTxCompressedByteRate int64 `json:"averageTxCompressedByteRate,omitempty"`
	// The one minute average of the message rate transmitted by the Broker, in messages per second (msg/sec). Available since 2.14.
	AverageTxMsgRate int64 `json:"averageTxMsgRate,omitempty"`
	// The one minute average of the uncompressed message rate transmitted by the Broker, in bytes per second (B/sec). Available since 2.14.
	AverageTxUncompressedByteRate int64 `json:"averageTxUncompressedByteRate,omitempty"`
	// The current CSPF version. Available since 2.17.
	CspfVersion int32 `json:"cspfVersion,omitempty"`
	// An approximation of the amount of disk space consumed, but not used, by the persisted data. Calculated as a percentage of total space. Available since 2.18.
	GuaranteedMsgingDefragmentationEstimatedFragmentation int64 `json:"guaranteedMsgingDefragmentationEstimatedFragmentation,omitempty"`
	// An approximation of the amount of disk space recovered upon a successfully completed execution of a defragmentation operation. Expressed in MB. Available since 2.18.
	GuaranteedMsgingDefragmentationEstimatedRecoverableSpace int32 `json:"guaranteedMsgingDefragmentationEstimatedRecoverableSpace,omitempty"`
	// A timestamp reflecting when the last defragmentation completed. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.18.
	GuaranteedMsgingDefragmentationLastCompletedOn int32 `json:"guaranteedMsgingDefragmentationLastCompletedOn,omitempty"`
	// How much of the message spool was visited during the last defragmentation operation. This number reflects the percentage of the message spool visited in terms of disk space (as opposed to, for example, spool files). Available since 2.18.
	GuaranteedMsgingDefragmentationLastCompletionPercentage int64 `json:"guaranteedMsgingDefragmentationLastCompletionPercentage,omitempty"`
	// Reflects how the last defragmentation operation completed. The allowed values and their meaning are:  <pre> \"success\" - Defragmentation completed successfully. \"unmovable-local-transaction\" - Defragmentation stopped after encountering an unmovable local transaction. \"unmovable-xa-transaction\" - Defragmentation stopped after encountering an unmovable XA transaction. \"incomplete\" - Defragmentation stopped prematurely. \"stopped-by-administrator\" - Defragmentation stopped by administrator. </pre>  Available since 2.18.
	GuaranteedMsgingDefragmentationLastExitCondition string `json:"guaranteedMsgingDefragmentationLastExitCondition,omitempty"`
	// Optional additional information regarding the exit condition of the last defragmentation operation. Available since 2.18.
	GuaranteedMsgingDefragmentationLastExitConditionInformation string `json:"guaranteedMsgingDefragmentationLastExitConditionInformation,omitempty"`
	// Defragmentation status of guaranteed messaging. The allowed values and their meaning are:  <pre> \"idle\" - Defragmentation is not currently running. \"pending\" - Degfragmentation is preparing to run. \"active\" - Defragmentation is in progress. </pre>  Available since 2.18.
	GuaranteedMsgingDefragmentationStatus string `json:"guaranteedMsgingDefragmentationStatus,omitempty"`
	// The estimated completion percentage of a defragmentation operation currently in progress. Only valid if the defragmentation status is \"Active\". Available since 2.18.
	GuaranteedMsgingDefragmentationStatusActiveCompletionPercentage int64 `json:"guaranteedMsgingDefragmentationStatusActiveCompletionPercentage,omitempty"`
	// Enable or disable Guaranteed Messaging. Available since 2.18.
	GuaranteedMsgingEnabled                                      bool                     `json:"guaranteedMsgingEnabled,omitempty"`
	GuaranteedMsgingEventCacheUsageThreshold                     *EventThreshold          `json:"guaranteedMsgingEventCacheUsageThreshold,omitempty"`
	GuaranteedMsgingEventDeliveredUnackedThreshold               *EventThresholdByPercent `json:"guaranteedMsgingEventDeliveredUnackedThreshold,omitempty"`
	GuaranteedMsgingEventDiskUsageThreshold                      *EventThresholdByPercent `json:"guaranteedMsgingEventDiskUsageThreshold,omitempty"`
	GuaranteedMsgingEventEgressFlowCountThreshold                *EventThreshold          `json:"guaranteedMsgingEventEgressFlowCountThreshold,omitempty"`
	GuaranteedMsgingEventEndpointCountThreshold                  *EventThreshold          `json:"guaranteedMsgingEventEndpointCountThreshold,omitempty"`
	GuaranteedMsgingEventIngressFlowCountThreshold               *EventThreshold          `json:"guaranteedMsgingEventIngressFlowCountThreshold,omitempty"`
	GuaranteedMsgingEventMsgCountThreshold                       *EventThresholdByPercent `json:"guaranteedMsgingEventMsgCountThreshold,omitempty"`
	GuaranteedMsgingEventMsgSpoolFileCountThreshold              *EventThresholdByPercent `json:"guaranteedMsgingEventMsgSpoolFileCountThreshold,omitempty"`
	GuaranteedMsgingEventMsgSpoolUsageThreshold                  *EventThreshold          `json:"guaranteedMsgingEventMsgSpoolUsageThreshold,omitempty"`
	GuaranteedMsgingEventTransactedSessionCountThreshold         *EventThreshold          `json:"guaranteedMsgingEventTransactedSessionCountThreshold,omitempty"`
	GuaranteedMsgingEventTransactedSessionResourceCountThreshold *EventThresholdByPercent `json:"guaranteedMsgingEventTransactedSessionResourceCountThreshold,omitempty"`
	GuaranteedMsgingEventTransactionCountThreshold               *EventThreshold          `json:"guaranteedMsgingEventTransactionCountThreshold,omitempty"`
	// Guaranteed messaging cache usage limit. Expressed as a maximum percentage of the NAB's egress queueing. resources that the guaranteed message cache is allowed to use. Available since 2.18.
	GuaranteedMsgingMaxCacheUsage int32 `json:"guaranteedMsgingMaxCacheUsage,omitempty"`
	// The maximum total message spool usage allowed across all VPNs on this broker, in megabytes. Recommendation: the maximum value should be less than 90% of the disk space allocated for the guaranteed message spool. Available since 2.18.
	GuaranteedMsgingMaxMsgSpoolUsage int64 `json:"guaranteedMsgingMaxMsgSpoolUsage,omitempty"`
	// The maximum time, in milliseconds, that can be tolerated for remote acknowledgement of synchronization messages before which the remote system will be considered out of sync. Available since 2.18.
	GuaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout int64 `json:"guaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout,omitempty"`
	// The maximum time, in milliseconds, that can be tolerated for remote disk writes before which the remote system will be considered out of sync. Available since 2.18.
	GuaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout int64 `json:"guaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout,omitempty"`
	// Operational status of guaranteed messaging. The allowed values and their meaning are:  <pre> \"disabled\" - The operational status of guaranteed messaging is Disabled. \"not-ready\" - The operational status of guaranteed messaging is NotReady. \"standby\" - The operational status of guaranteed messaging is Standby. \"activating\" - The operational status of guaranteed messaging is Activating. \"active\" - The operational status of guaranteed messaging is Active. </pre>  Available since 2.18.
	GuaranteedMsgingOperationalStatus string `json:"guaranteedMsgingOperationalStatus,omitempty"`
	// The replication compatibility mode for the router. The default value is `\"legacy\"`. The allowed values and their meaning are:\"legacy\" - All transactions originated by clients are replicated to the standby site without using transactions.\"transacted\" - All transactions originated by clients are replicated to the standby site using transactions. The allowed values and their meaning are:  <pre> \"legacy\" - All transactions originated by clients are replicated to the standby site without using transactions. \"transacted\" - All transactions originated by clients are replicated to the standby site using transactions. </pre>  Available since 2.18.
	GuaranteedMsgingTransactionReplicationCompatibilityMode string `json:"guaranteedMsgingTransactionReplicationCompatibilityMode,omitempty"`
	// The amount of messages received from clients by the Broker, in bytes (B). Available since 2.14.
	RxByteCount int64 `json:"rxByteCount,omitempty"`
	// The current message rate received by the Broker, in bytes per second (B/sec). Available since 2.14.
	RxByteRate int64 `json:"rxByteRate,omitempty"`
	// The amount of compressed messages received by the Broker, in bytes (B). Available since 2.14.
	RxCompressedByteCount int64 `json:"rxCompressedByteCount,omitempty"`
	// The current compressed message rate received by the Broker, in bytes per second (B/sec). Available since 2.14.
	RxCompressedByteRate int64 `json:"rxCompressedByteRate,omitempty"`
	// The compression ratio for messages received by the Broker. Available since 2.14.
	RxCompressionRatio string `json:"rxCompressionRatio,omitempty"`
	// The number of messages received from clients by the Broker. Available since 2.14.
	RxMsgCount int64 `json:"rxMsgCount,omitempty"`
	// The current message rate received by the Broker, in messages per second (msg/sec). Available since 2.14.
	RxMsgRate int64 `json:"rxMsgRate,omitempty"`
	// The amount of uncompressed messages received by the Broker, in bytes (B). Available since 2.14.
	RxUncompressedByteCount int64 `json:"rxUncompressedByteCount,omitempty"`
	// The current uncompressed message rate received by the Broker, in bytes per second (B/sec). Available since 2.14.
	RxUncompressedByteRate int64 `json:"rxUncompressedByteRate,omitempty"`
	// Enable or disable the AMQP service. When disabled new AMQP Clients may not connect through the global or per-VPN AMQP listen-ports, and all currently connected AMQP Clients are immediately disconnected. Available since 2.17.
	ServiceAmqpEnabled bool `json:"serviceAmqpEnabled,omitempty"`
	// TCP port number that AMQP clients can use to connect to the broker using raw TCP over TLS. Available since 2.17.
	ServiceAmqpTlsListenPort             int64           `json:"serviceAmqpTlsListenPort,omitempty"`
	ServiceEventConnectionCountThreshold *EventThreshold `json:"serviceEventConnectionCountThreshold,omitempty"`
	// Enable or disable the health-check service. Available since 2.17.
	ServiceHealthCheckEnabled bool `json:"serviceHealthCheckEnabled,omitempty"`
	// The port number for the health-check service. The port must be unique across the message backbone. The health-check service must be disabled to change the port. Available since 2.17.
	ServiceHealthCheckListenPort int64 `json:"serviceHealthCheckListenPort,omitempty"`
	// Enable or disable the mate-link service. Available since 2.17.
	ServiceMateLinkEnabled bool `json:"serviceMateLinkEnabled,omitempty"`
	// The port number for the mate-link service. The port must be unique across the message backbone. The mate-link service must be disabled to change the port. Available since 2.17.
	ServiceMateLinkListenPort int64 `json:"serviceMateLinkListenPort,omitempty"`
	// Enable or disable the MQTT service. When disabled new MQTT Clients may not connect through the per-VPN MQTT listen-ports, and all currently connected MQTT Clients are immediately disconnected. Available since 2.17.
	ServiceMqttEnabled bool `json:"serviceMqttEnabled,omitempty"`
	// Enable or disable the msg-backbone service. When disabled new Clients may not connect through global or per-VPN listen-ports, and all currently connected Clients are immediately disconnected. Available since 2.17.
	ServiceMsgBackboneEnabled bool `json:"serviceMsgBackboneEnabled,omitempty"`
	// Enable or disable the redundancy service. Available since 2.17.
	ServiceRedundancyEnabled bool `json:"serviceRedundancyEnabled,omitempty"`
	// The first listen-port used for the redundancy service. Redundancy uses this port and the subsequent 2 ports. These port must be unique across the message backbone. The redundancy service must be disabled to change this port. Available since 2.17.
	ServiceRedundancyFirstListenPort                 int64           `json:"serviceRedundancyFirstListenPort,omitempty"`
	ServiceRestEventOutgoingConnectionCountThreshold *EventThreshold `json:"serviceRestEventOutgoingConnectionCountThreshold,omitempty"`
	// Enable or disable the REST service incoming connections on the router. Available since 2.17.
	ServiceRestIncomingEnabled bool `json:"serviceRestIncomingEnabled,omitempty"`
	// Enable or disable the REST service outgoing connections on the router. Available since 2.17.
	ServiceRestOutgoingEnabled bool `json:"serviceRestOutgoingEnabled,omitempty"`
	// Enable or disable extended SEMP timeouts for paged GETs. When a request times out, it returns the current page of content, even if the page is not full.  When enabled, the timeout is 60 seconds. When disabled, the timeout is 5 seconds.  The recommended setting is disabled (no legacy-timeout).  This parameter is intended as a temporary workaround to be used until SEMP clients can handle short pages.  This setting will be removed in a future release. Available since 2.18.
	ServiceSempLegacyTimeoutEnabled bool `json:"serviceSempLegacyTimeoutEnabled,omitempty"`
	// Enable or disable plain-text SEMP service. Available since 2.17.
	ServiceSempPlainTextEnabled bool `json:"serviceSempPlainTextEnabled,omitempty"`
	// The TCP port for plain-text SEMP client connections. Available since 2.17.
	ServiceSempPlainTextListenPort int64 `json:"serviceSempPlainTextListenPort,omitempty"`
	// The session idle timeout, in minutes. Sessions will be invalidated if there is no activity in this period of time. Available since 2.21.
	ServiceSempSessionIdleTimeout int32 `json:"serviceSempSessionIdleTimeout,omitempty"`
	// The maximum lifetime of a session, in minutes. Sessions will be invalidated after this period of time, regardless of activity. Available since 2.21.
	ServiceSempSessionMaxLifetime int32 `json:"serviceSempSessionMaxLifetime,omitempty"`
	// Enable or disable TLS SEMP service. Available since 2.17.
	ServiceSempTlsEnabled bool `json:"serviceSempTlsEnabled,omitempty"`
	// The TCP port for TLS SEMP client connections. Available since 2.17.
	ServiceSempTlsListenPort int64 `json:"serviceSempTlsListenPort,omitempty"`
	// TCP port number that SMF clients can use to connect to the broker using raw compression TCP. Available since 2.17.
	ServiceSmfCompressionListenPort int64 `json:"serviceSmfCompressionListenPort,omitempty"`
	// Enable or disable the SMF service. When disabled new SMF Clients may not connect through the global listen-ports, and all currently connected SMF Clients are immediately disconnected. Available since 2.17.
	ServiceSmfEnabled                       bool            `json:"serviceSmfEnabled,omitempty"`
	ServiceSmfEventConnectionCountThreshold *EventThreshold `json:"serviceSmfEventConnectionCountThreshold,omitempty"`
	// TCP port number that SMF clients can use to connect to the broker using raw TCP. Available since 2.17.
	ServiceSmfPlainTextListenPort int64 `json:"serviceSmfPlainTextListenPort,omitempty"`
	// TCP port number that SMF clients can use to connect to the broker using raw routing control TCP. Available since 2.17.
	ServiceSmfRoutingControlListenPort int64 `json:"serviceSmfRoutingControlListenPort,omitempty"`
	// TCP port number that SMF clients can use to connect to the broker using raw TCP over TLS. Available since 2.17.
	ServiceSmfTlsListenPort                 int64           `json:"serviceSmfTlsListenPort,omitempty"`
	ServiceTlsEventConnectionCountThreshold *EventThreshold `json:"serviceTlsEventConnectionCountThreshold,omitempty"`
	// Enable or disable the web-transport service. When disabled new web-transport Clients may not connect through the global listen-ports, and all currently connected web-transport Clients are immediately disconnected. Available since 2.17.
	ServiceWebTransportEnabled bool `json:"serviceWebTransportEnabled,omitempty"`
	// The TCP port for plain-text WEB client connections. Available since 2.17.
	ServiceWebTransportPlainTextListenPort int64 `json:"serviceWebTransportPlainTextListenPort,omitempty"`
	// The TCP port for TLS WEB client connections. Available since 2.17.
	ServiceWebTransportTlsListenPort int64 `json:"serviceWebTransportTlsListenPort,omitempty"`
	// Used to specify the Web URL suffix that will be used by Web clients when communicating with the broker. Available since 2.17.
	ServiceWebTransportWebUrlSuffix string `json:"serviceWebTransportWebUrlSuffix,omitempty"`
	// Indicates whether TLS version 1.1 connections are blocked. When blocked, all existing incoming and outgoing TLS 1.1 connections with Clients, SEMP users, and LDAP servers remain connected while new connections are blocked. Note that support for TLS 1.1 will eventually be discontinued, at which time TLS 1.1 connections will be blocked regardless of this setting.
	TlsBlockVersion11Enabled bool `json:"tlsBlockVersion11Enabled,omitempty"`
	// The colon-separated list of default cipher suites for TLS management connections.
	TlsCipherSuiteManagementDefaultList string `json:"tlsCipherSuiteManagementDefaultList,omitempty"`
	// The colon-separated list of cipher suites used for TLS management connections (e.g. SEMP, LDAP). The value \"default\" implies all supported suites ordered from most secure to least secure.
	TlsCipherSuiteManagementList string `json:"tlsCipherSuiteManagementList,omitempty"`
	// The colon-separated list of supported cipher suites for TLS management connections.
	TlsCipherSuiteManagementSupportedList string `json:"tlsCipherSuiteManagementSupportedList,omitempty"`
	// The colon-separated list of default cipher suites for TLS data connections.
	TlsCipherSuiteMsgBackboneDefaultList string `json:"tlsCipherSuiteMsgBackboneDefaultList,omitempty"`
	// The colon-separated list of cipher suites used for TLS data connections (e.g. client pub/sub). The value \"default\" implies all supported suites ordered from most secure to least secure.
	TlsCipherSuiteMsgBackboneList string `json:"tlsCipherSuiteMsgBackboneList,omitempty"`
	// The colon-separated list of supported cipher suites for TLS data connections.
	TlsCipherSuiteMsgBackboneSupportedList string `json:"tlsCipherSuiteMsgBackboneSupportedList,omitempty"`
	// The colon-separated list of default cipher suites for TLS secure shell connections.
	TlsCipherSuiteSecureShellDefaultList string `json:"tlsCipherSuiteSecureShellDefaultList,omitempty"`
	// The colon-separated list of cipher suites used for TLS secure shell connections (e.g. SSH, SFTP, SCP). The value \"default\" implies all supported suites ordered from most secure to least secure.
	TlsCipherSuiteSecureShellList string `json:"tlsCipherSuiteSecureShellList,omitempty"`
	// The colon-separated list of supported cipher suites for TLS secure shell connections.
	TlsCipherSuiteSecureShellSupportedList string `json:"tlsCipherSuiteSecureShellSupportedList,omitempty"`
	// Indicates whether protection against the CRIME exploit is enabled. When enabled, TLS+compressed messaging performance is degraded. This protection should only be disabled if sufficient ACL and authentication features are being employed such that a potential attacker does not have sufficient access to trigger the exploit.
	TlsCrimeExploitProtectionEnabled bool `json:"tlsCrimeExploitProtectionEnabled,omitempty"`
	// Enable or disable the standard domain certificate authority list. Available since 2.19.
	TlsStandardDomainCertificateAuthoritiesEnabled bool `json:"tlsStandardDomainCertificateAuthoritiesEnabled,omitempty"`
	// The TLS ticket lifetime in seconds. When a client connects with TLS, a session with a session ticket is created using the TLS ticket lifetime which determines how long the client has to resume the session.
	TlsTicketLifetime int32 `json:"tlsTicketLifetime,omitempty"`
	// The comma-separated list of supported TLS versions.
	TlsVersionSupportedList string `json:"tlsVersionSupportedList,omitempty"`
	// The amount of messages transmitted to clients by the Broker, in bytes (B). Available since 2.14.
	TxByteCount int64 `json:"txByteCount,omitempty"`
	// The current message rate transmitted by the Broker, in bytes per second (B/sec). Available since 2.14.
	TxByteRate int64 `json:"txByteRate,omitempty"`
	// The amount of compressed messages transmitted by the Broker, in bytes (B). Available since 2.14.
	TxCompressedByteCount int64 `json:"txCompressedByteCount,omitempty"`
	// The current compressed message rate transmitted by the Broker, in bytes per second (B/sec). Available since 2.14.
	TxCompressedByteRate int64 `json:"txCompressedByteRate,omitempty"`
	// The compression ratio for messages transmitted by the Broker. Available since 2.14.
	TxCompressionRatio string `json:"txCompressionRatio,omitempty"`
	// The number of messages transmitted to clients by the Broker. Available since 2.14.
	TxMsgCount int64 `json:"txMsgCount,omitempty"`
	// The current message rate transmitted by the Broker, in messages per second (msg/sec). Available since 2.14.
	TxMsgRate int64 `json:"txMsgRate,omitempty"`
	// The amount of uncompressed messages transmitted by the Broker, in bytes (B). Available since 2.14.
	TxUncompressedByteCount int64 `json:"txUncompressedByteCount,omitempty"`
	// The current uncompressed message rate transmitted by the Broker, in bytes per second (B/sec). Available since 2.14.
	TxUncompressedByteRate int64 `json:"txUncompressedByteRate,omitempty"`
}
