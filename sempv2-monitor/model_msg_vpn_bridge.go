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

type MsgVpnBridge struct {
	// The one minute average of the message rate received from the Bridge, in bytes per second (B/sec). Available since 2.13.
	AverageRxByteRate int64 `json:"averageRxByteRate,omitempty"`
	// The one minute average of the message rate received from the Bridge, in messages per second (msg/sec). Available since 2.13.
	AverageRxMsgRate int64 `json:"averageRxMsgRate,omitempty"`
	// The one minute average of the message rate transmitted to the Bridge, in bytes per second (B/sec). Available since 2.13.
	AverageTxByteRate int64 `json:"averageTxByteRate,omitempty"`
	// The one minute average of the message rate transmitted to the Bridge, in messages per second (msg/sec). Available since 2.13.
	AverageTxMsgRate int64 `json:"averageTxMsgRate,omitempty"`
	// Indicates whether the Bridge is bound to the queue in the remote Message VPN.
	BoundToQueue bool `json:"boundToQueue,omitempty"`
	// The name of the Bridge.
	BridgeName string `json:"bridgeName,omitempty"`
	// The virtual router of the Bridge. The allowed values and their meaning are:  <pre> \"primary\" - The Bridge is used for the primary virtual router. \"backup\" - The Bridge is used for the backup virtual router. \"auto\" - The Bridge is automatically assigned a virtual router at creation, depending on the broker's active-standby role. </pre>
	BridgeVirtualRouter string `json:"bridgeVirtualRouter,omitempty"`
	// The name of the Client for the Bridge.
	ClientName string `json:"clientName,omitempty"`
	// Indicates whether messages transmitted over the Bridge are compressed.
	Compressed bool `json:"compressed,omitempty"`
	// The amount of client control messages received from the Bridge, in bytes (B). Available since 2.13.
	ControlRxByteCount int64 `json:"controlRxByteCount,omitempty"`
	// The number of client control messages received from the Bridge. Available since 2.13.
	ControlRxMsgCount int64 `json:"controlRxMsgCount,omitempty"`
	// The amount of client control messages transmitted to the Bridge, in bytes (B). Available since 2.13.
	ControlTxByteCount int64 `json:"controlTxByteCount,omitempty"`
	// The number of client control messages transmitted to the Bridge. Available since 2.13.
	ControlTxMsgCount int64                `json:"controlTxMsgCount,omitempty"`
	Counter           *MsgVpnBridgeCounter `json:"counter,omitempty"`
	// The amount of client data messages received from the Bridge, in bytes (B). Available since 2.13.
	DataRxByteCount int64 `json:"dataRxByteCount,omitempty"`
	// The number of client data messages received from the Bridge. Available since 2.13.
	DataRxMsgCount int64 `json:"dataRxMsgCount,omitempty"`
	// The amount of client data messages transmitted to the Bridge, in bytes (B). Available since 2.13.
	DataTxByteCount int64 `json:"dataTxByteCount,omitempty"`
	// The number of client data messages transmitted to the Bridge. Available since 2.13.
	DataTxMsgCount int64 `json:"dataTxMsgCount,omitempty"`
	// The number of messages discarded during reception from the Bridge. Available since 2.13.
	DiscardedRxMsgCount int64 `json:"discardedRxMsgCount,omitempty"`
	// The number of messages discarded during transmission to the Bridge. Available since 2.13.
	DiscardedTxMsgCount int64 `json:"discardedTxMsgCount,omitempty"`
	// Indicates whether the Bridge is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Indicates whether messages transmitted over the Bridge are encrypted with TLS.
	Encrypted bool `json:"encrypted,omitempty"`
	// The establisher of the Bridge connection. The allowed values and their meaning are:  <pre> \"local\" - The Bridge connection was established by the local Message VPN. \"remote\" - The Bridge connection was established by the remote Message VPN. </pre>
	Establisher string `json:"establisher,omitempty"`
	// The reason for the inbound connection failure from the Bridge. If there is no failure reason, an empty string (\"\") is returned.
	InboundFailureReason string `json:"inboundFailureReason,omitempty"`
	// The state of the inbound connection from the Bridge. The allowed values and their meaning are:  <pre> \"init\" - The connection is initializing. \"disabled\" - The connection is disabled by configuration. \"enabled\" - The connection is enabled by configuration. \"prepare\" - The connection is operationally down. \"prepare-wait-to-connect\" - The connection is waiting to connect. \"prepare-fetching-dns\" - The domain name of the destination node is being resolved. \"not-ready\" - The connection is operationally down. \"not-ready-connecting\" - The connection is trying to connect. \"not-ready-handshaking\" - The connection is handshaking. \"not-ready-wait-next\" - The connection failed to connect and is waiting to retry. \"not-ready-wait-reuse\" - The connection is closing in order to reuse an existing connection. \"not-ready-wait-bridge-version-mismatch\" - The connection is closing because of a version mismatch. \"not-ready-wait-cleanup\" - The connection is closed and cleaning up. \"ready\" - The connection is operationally up. \"ready-subscribing\" - The connection is up and synchronizing subscriptions. \"ready-in-sync\" - The connection is up and subscriptions are synchronized. </pre>
	InboundState string `json:"inboundState,omitempty"`
	// The ID of the last message transmitted to the Bridge.
	LastTxMsgId int64 `json:"lastTxMsgId,omitempty"`
	// The physical interface on the local Message VPN host for connecting to the remote Message VPN.
	LocalInterface string `json:"localInterface,omitempty"`
	// The name of the local queue for the Bridge.
	LocalQueueName string `json:"localQueueName,omitempty"`
	// The number of login request messages received from the Bridge. Available since 2.13.
	LoginRxMsgCount int64 `json:"loginRxMsgCount,omitempty"`
	// The number of login response messages transmitted to the Bridge. Available since 2.13.
	LoginTxMsgCount int64 `json:"loginTxMsgCount,omitempty"`
	// The maximum time-to-live (TTL) in hops. Messages are discarded if their TTL exceeds this value.
	MaxTtl int64 `json:"maxTtl,omitempty"`
	// The number of guaranteed messages received from the Bridge. Available since 2.13.
	MsgSpoolRxMsgCount int64 `json:"msgSpoolRxMsgCount,omitempty"`
	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`
	// The state of the outbound connection to the Bridge. The allowed values and their meaning are:  <pre> \"init\" - The connection is initializing. \"disabled\" - The connection is disabled by configuration. \"enabled\" - The connection is enabled by configuration. \"prepare\" - The connection is operationally down. \"prepare-wait-to-connect\" - The connection is waiting to connect. \"prepare-fetching-dns\" - The domain name of the destination node is being resolved. \"not-ready\" - The connection is operationally down. \"not-ready-connecting\" - The connection is trying to connect. \"not-ready-handshaking\" - The connection is handshaking. \"not-ready-wait-next\" - The connection failed to connect and is waiting to retry. \"not-ready-wait-reuse\" - The connection is closing in order to reuse an existing connection. \"not-ready-wait-bridge-version-mismatch\" - The connection is closing because of a version mismatch. \"not-ready-wait-cleanup\" - The connection is closed and cleaning up. \"ready\" - The connection is operationally up. \"ready-subscribing\" - The connection is up and synchronizing subscriptions. \"ready-in-sync\" - The connection is up and subscriptions are synchronized. </pre>
	OutboundState string            `json:"outboundState,omitempty"`
	Rate          *MsgVpnBridgeRate `json:"rate,omitempty"`
	// The FQDN or IP address of the remote Message VPN.
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// The Client Username the Bridge uses to login to the remote Message VPN.
	RemoteAuthenticationBasicClientUsername string `json:"remoteAuthenticationBasicClientUsername,omitempty"`
	// The authentication scheme for the remote Message VPN. The allowed values and their meaning are:  <pre> \"basic\" - Basic Authentication Scheme (via username and password). \"client-certificate\" - Client Certificate Authentication Scheme (via certificate file or content). </pre>
	RemoteAuthenticationScheme string `json:"remoteAuthenticationScheme,omitempty"`
	// The maximum number of retry attempts to establish a connection to the remote Message VPN. A value of 0 means to retry forever.
	RemoteConnectionRetryCount int64 `json:"remoteConnectionRetryCount,omitempty"`
	// The number of seconds the broker waits for the bridge connection to be established before attempting a new connection.
	RemoteConnectionRetryDelay int64 `json:"remoteConnectionRetryDelay,omitempty"`
	// The priority for deliver-to-one (DTO) messages transmitted from the remote Message VPN. The allowed values and their meaning are:  <pre> \"p1\" - The 1st or highest priority. \"p2\" - The 2nd highest priority. \"p3\" - The 3rd highest priority. \"p4\" - The 4th highest priority. \"da\" - Ignore priority and deliver always. </pre>
	RemoteDeliverToOnePriority string `json:"remoteDeliverToOnePriority,omitempty"`
	// The name of the remote Message VPN.
	RemoteMsgVpnName string `json:"remoteMsgVpnName,omitempty"`
	// The name of the remote router.
	RemoteRouterName string `json:"remoteRouterName,omitempty"`
	// The ID of the transmit flow for the connected remote Message VPN.
	RemoteTxFlowId int32 `json:"remoteTxFlowId,omitempty"`
	// The amount of messages received from the Bridge, in bytes (B). Available since 2.13.
	RxByteCount int64 `json:"rxByteCount,omitempty"`
	// The current message rate received from the Bridge, in bytes per second (B/sec). Available since 2.13.
	RxByteRate int64 `json:"rxByteRate,omitempty"`
	// The category of the inbound connection failure from the Bridge. The allowed values and their meaning are:  <pre> \"no-failure\" - There is no bridge failure. \"local-configuration-problem\" - The bridge failure is a local configuration problem. \"local-operational-state-problem\" - The bridge failure is an operational state problem. </pre>  Available since 2.18.
	RxConnectionFailureCategory string `json:"rxConnectionFailureCategory,omitempty"`
	// The number of messages received from the Bridge. Available since 2.13.
	RxMsgCount int64 `json:"rxMsgCount,omitempty"`
	// The current message rate received from the Bridge, in messages per second (msg/sec). Available since 2.13.
	RxMsgRate int64 `json:"rxMsgRate,omitempty"`
	// The colon-separated list of cipher suites supported for TLS connections to the remote Message VPN. The value \"default\" implies all supported suites ordered from most secure to least secure.
	TlsCipherSuiteList string `json:"tlsCipherSuiteList,omitempty"`
	// Indicates whether the Bridge is configured to use the default cipher-suite list.
	TlsDefaultCipherSuiteList bool `json:"tlsDefaultCipherSuiteList,omitempty"`
	// Indicates whether the TTL (hops) exceeded event has been raised.
	TtlExceededEventRaised bool `json:"ttlExceededEventRaised,omitempty"`
	// The amount of messages transmitted to the Bridge, in bytes (B). Available since 2.13.
	TxByteCount int64 `json:"txByteCount,omitempty"`
	// The current message rate transmitted to the Bridge, in bytes per second (B/sec). Available since 2.13.
	TxByteRate int64 `json:"txByteRate,omitempty"`
	// The number of messages transmitted to the Bridge. Available since 2.13.
	TxMsgCount int64 `json:"txMsgCount,omitempty"`
	// The current message rate transmitted to the Bridge, in messages per second (msg/sec). Available since 2.13.
	TxMsgRate int64 `json:"txMsgRate,omitempty"`
	// The amount of time in seconds since the Bridge connected to the remote Message VPN.
	Uptime int64 `json:"uptime,omitempty"`
}