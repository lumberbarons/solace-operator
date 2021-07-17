/*
 * SEMP (Solace Element Management Protocol)
 *
 * SEMP (starting in `v2`, see note 1) is a RESTful API for configuring, monitoring, and administering a Solace PubSub+ broker.  SEMP uses URIs to address manageable **resources** of the Solace PubSub+ broker. Resources are individual **objects**, **collections** of objects, or (exclusively in the action API) **actions**. This document applies to the following API:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Monitoring|/SEMP/v2/monitor|Querying operational parameters|See note 2    The following APIs are also available:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Action|/SEMP/v2/action|Performing actions|See note 2 Configuration|/SEMP/v2/config|Reading and writing config state|See note 2    Resources are always nouns, with individual objects being singular and collections being plural.  Objects within a collection are identified by an `obj-id`, which follows the collection name with the form `collection-name/obj-id`.  Actions within an object are identified by an `action-id`, which follows the object name with the form `obj-id/action-id`.  Some examples:  ``` /SEMP/v2/config/msgVpns                        ; MsgVpn collection /SEMP/v2/config/msgVpns/a                      ; MsgVpn object named \"a\" /SEMP/v2/config/msgVpns/a/queues               ; Queue collection in MsgVpn \"a\" /SEMP/v2/config/msgVpns/a/queues/b             ; Queue object named \"b\" in MsgVpn \"a\" /SEMP/v2/action/msgVpns/a/queues/b/startReplay ; Action that starts a replay on Queue \"b\" in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients             ; Client collection in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients/c           ; Client object named \"c\" in MsgVpn \"a\" ```  ## Collection Resources  Collections are unordered lists of objects (unless described as otherwise), and are described by JSON arrays. Each item in the array represents an object in the same manner as the individual object would normally be represented. In the configuration API, the creation of a new object is done through its collection resource.  ## Object and Action Resources  Objects are composed of attributes, actions, collections, and other objects. They are described by JSON objects as name/value pairs. The collections and actions of an object are not contained directly in the object's JSON content; rather the content includes an attribute containing a URI which points to the collections and actions. These contained resources must be managed through this URI. At a minimum, every object has one or more identifying attributes, and its own `uri` attribute which contains the URI pointing to itself.  Actions are also composed of attributes, and are described by JSON objects as name/value pairs. Unlike objects, however, they are not members of a collection and cannot be retrieved, only performed. Actions only exist in the action API.  Attributes in an object or action may have any combination of the following properties:   Property|Meaning|Comments :---|:---|:--- Identifying|Attribute is involved in unique identification of the object, and appears in its URI| Required|Attribute must be provided in the request| Read-Only|Attribute can only be read, not written.|See note 3 Write-Only|Attribute can only be written, not read, unless the attribute is also opaque|See the documentation for the opaque property Requires-Disable|Attribute can only be changed when object is disabled| Deprecated|Attribute is deprecated, and will disappear in the next SEMP version| Opaque|Attribute can be set or retrieved in opaque form when the `opaquePassword` query parameter is present|See the `opaquePassword` query parameter documentation    In some requests, certain attributes may only be provided in certain combinations with other attributes:   Relationship|Meaning :---|:--- Requires|Attribute may only be changed by a request if a particular attribute or combination of attributes is also provided in the request Conflicts|Attribute may only be provided in a request if a particular attribute or combination of attributes is not also provided in the request    In the monitoring API, any non-identifying attribute may not be returned in a GET.  ## HTTP Methods  The following HTTP methods manipulate resources in accordance with these general principles. Note that some methods are only used in certain APIs:   Method|Resource|Meaning|Request Body|Response Body|Missing Request Attributes :---|:---|:---|:---|:---|:--- POST|Collection|Create object|Initial attribute values|Object attributes and metadata|Set to default PUT|Object|Create or replace object (see note 5)|New attribute values|Object attributes and metadata|Set to default, with certain exceptions (see note 4) PUT|Action|Performs action|Action arguments|Action metadata|N/A PATCH|Object|Update object|New attribute values|Object attributes and metadata|unchanged DELETE|Object|Delete object|Empty|Object metadata|N/A GET|Object|Get object|Empty|Object attributes and metadata|N/A GET|Collection|Get collection|Empty|Object attributes and collection metadata|N/A    ## Common Query Parameters  The following are some common query parameters that are supported by many method/URI combinations. Individual URIs may document additional parameters. Note that multiple query parameters can be used together in a single URI, separated by the ampersand character. For example:  ``` ; Request for the MsgVpns collection using two hypothetical query parameters ; \"q1\" and \"q2\" with values \"val1\" and \"val2\" respectively /SEMP/v2/monitor/msgVpns?q1=val1&q2=val2 ```  ### select  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. Use this query parameter to limit the size of the returned data for each returned object, return only those fields that are desired, or exclude fields that are not desired.  The value of `select` is a comma-separated list of attribute names. If the list contains attribute names that are not prefaced by `-`, only those attributes are included in the response. If the list contains attribute names that are prefaced by `-`, those attributes are excluded from the response. If the list contains both types, then the difference of the first set of attributes and the second set of attributes is returned. If the list is empty (i.e. `select=`), no attributes are returned.  All attributes that are prefaced by `-` must follow all attributes that are not prefaced by `-`. In addition, each attribute name in the list must match at least one attribute in the object.  Names may include the `*` wildcard (zero or more characters). Nested attribute names are supported using periods (e.g. `parentName.childName`).  Some examples:  ``` ; List of all MsgVpn names /SEMP/v2/monitor/msgVpns?select=msgVpnName ; List of all MsgVpn and their attributes except for their names /SEMP/v2/monitor/msgVpns?select=-msgVpnName ; Authentication attributes of MsgVpn \"finance\" /SEMP/v2/monitor/msgVpns/finance?select=authentication* ; All attributes of MsgVpn \"finance\" except for authentication attributes /SEMP/v2/monitor/msgVpns/finance?select=-authentication* ; Access related attributes of Queue \"orderQ\" of MsgVpn \"finance\" /SEMP/v2/monitor/msgVpns/finance/queues/orderQ?select=owner,permission ```  ### where  Include in the response only objects where certain conditions are true. Use this query parameter to limit which objects are returned to those whose attribute values meet the given conditions.  The value of `where` is a comma-separated list of expressions. All expressions must be true for the object to be included in the response. Each expression takes the form:  ``` expression  = attribute-name OP value OP          = '==' | '!=' | '&lt;' | '&gt;' | '&lt;=' | '&gt;=' ```  `value` may be a number, string, `true`, or `false`, as appropriate for the type of `attribute-name`. Greater-than and less-than comparisons only work for numbers. A `*` in a string `value` is interpreted as a wildcard (zero or more characters). Some examples:  ``` ; Only enabled MsgVpns /SEMP/v2/monitor/msgVpns?where=enabled==true ; Only MsgVpns using basic non-LDAP authentication /SEMP/v2/monitor/msgVpns?where=authenticationBasicEnabled==true,authenticationBasicType!=ldap ; Only MsgVpns that allow more than 100 client connections /SEMP/v2/monitor/msgVpns?where=maxConnectionCount>100 ; Only MsgVpns with msgVpnName starting with \"B\": /SEMP/v2/monitor/msgVpns?where=msgVpnName==B* ```  ### count  Limit the count of objects in the response. This can be useful to limit the size of the response for large collections. The minimum value for `count` is `1` and the default is `10`. There is also a per-collection maximum value to limit request handling time. For example:  ``` ; Up to 25 MsgVpns /SEMP/v2/monitor/msgVpns?count=25 ```  ### cursor  The cursor, or position, for the next page of objects. Cursors are opaque data that should not be created or interpreted by SEMP clients, and should only be used as described below.  When a request is made for a collection and there may be additional objects available for retrieval that are not included in the initial response, the response will include a `cursorQuery` field containing a cursor. The value of this field can be specified in the `cursor` query parameter of a subsequent request to retrieve the next page of objects. For convenience, an appropriate URI is constructed automatically by the broker and included in the `nextPageUri` field of the response. This URI can be used directly to retrieve the next page of objects.  ### opaquePassword  Attributes with the opaque property are also write-only and so cannot normally be retrieved in a GET. However, when a password is provided in the `opaquePassword` query parameter, attributes with the opaque property are retrieved in a GET in opaque form, encrypted with this password. The query parameter can also be used on a POST, PATCH, or PUT to set opaque attributes using opaque attribute values retrieved in a GET, so long as:  1. the same password that was used to retrieve the opaque attribute values is provided; and  2. the broker to which the request is being sent has the same major and minor SEMP version as the broker that produced the opaque attribute values.  The password provided in the query parameter must be a minimum of 8 characters and a maximum of 128 characters.  The query parameter can only be used in the configuration API, and only over HTTPS.  ## Authentication  When a client makes its first SEMPv2 request, it must supply a username and password using HTTP Basic authentication.  If authentication is successful, the broker returns a cookie containing a session key. The client can omit the username and password from subsequent requests, because the broker now uses the session cookie for authentication instead. When the session expires or is deleted, the client must provide the username and password again, and the broker creates a new session.  There are a limited number of session slots available on the broker. The broker returns 529 No SEMP Session Available if it is not able to allocate a session. For this reason, all clients that use SEMPv2 should support cookies.  If certain attributes—such as a user's password—are changed, the broker automatically deletes the affected sessions. These attributes are documented below. However, changes in external user configuration data stored on a RADIUS or LDAP server do not trigger the broker to delete the associated session(s), therefore you must do this manually, if required.  A client can retrieve its current session information using the /about/user endpoint, delete its own session using the /about/user/logout endpoint, and manage all sessions using the /sessions endpoint.  ## Help  Visit [our website](https://solace.com) to learn more about Solace.  You can also download the SEMP API specifications by clicking [here](https://solace.com/downloads/).  If you need additional support, please contact us at [support@solace.com](mailto:support@solace.com).  ## Notes  Note|Description :---:|:--- 1|This specification defines SEMP starting in \"v2\", and not the original SEMP \"v1\" interface. Request and response formats between \"v1\" and \"v2\" are entirely incompatible, although both protocols share a common port configuration on the Solace PubSub+ broker. They are differentiated by the initial portion of the URI path, one of either \"/SEMP/\" or \"/SEMP/v2/\" 2|This API is partially implemented. Only a subset of all objects are available. 3|Read-only attributes may appear in POST and PUT/PATCH requests. However, if a read-only attribute is not marked as identifying, it will be ignored during a PUT/PATCH. 4|On a PUT, if the SEMP user is not authorized to modify the attribute, its value is left unchanged rather than set to default. In addition, the values of write-only attributes are not set to their defaults on a PUT, except in the following two cases: there is a mutual requires relationship with another non-write-only attribute, both attributes are absent from the request, and the non-write-only attribute is not currently set to its default value; or the attribute is also opaque and the `opaquePassword` query parameter is provided in the request. 5|On a PUT, if the object does not exist, it is created first.
 *
 * API version: 2.21
 * Contact: support@solace.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MsgVpnBridge struct for MsgVpnBridge
type MsgVpnBridge struct {
	// The one minute average of the message rate received from the Bridge, in bytes per second (B/sec). Available since 2.13.
	AverageRxByteRate *int64 `json:"averageRxByteRate,omitempty"`
	// The one minute average of the message rate received from the Bridge, in messages per second (msg/sec). Available since 2.13.
	AverageRxMsgRate *int64 `json:"averageRxMsgRate,omitempty"`
	// The one minute average of the message rate transmitted to the Bridge, in bytes per second (B/sec). Available since 2.13.
	AverageTxByteRate *int64 `json:"averageTxByteRate,omitempty"`
	// The one minute average of the message rate transmitted to the Bridge, in messages per second (msg/sec). Available since 2.13.
	AverageTxMsgRate *int64 `json:"averageTxMsgRate,omitempty"`
	// Indicates whether the Bridge is bound to the queue in the remote Message VPN.
	BoundToQueue *bool `json:"boundToQueue,omitempty"`
	// The name of the Bridge.
	BridgeName *string `json:"bridgeName,omitempty"`
	// The virtual router of the Bridge. The allowed values and their meaning are:  <pre> \"primary\" - The Bridge is used for the primary virtual router. \"backup\" - The Bridge is used for the backup virtual router. \"auto\" - The Bridge is automatically assigned a virtual router at creation, depending on the broker's active-standby role. </pre>
	BridgeVirtualRouter *string `json:"bridgeVirtualRouter,omitempty"`
	// The name of the Client for the Bridge.
	ClientName *string `json:"clientName,omitempty"`
	// Indicates whether messages transmitted over the Bridge are compressed.
	Compressed *bool `json:"compressed,omitempty"`
	// The amount of client control messages received from the Bridge, in bytes (B). Available since 2.13.
	ControlRxByteCount *int64 `json:"controlRxByteCount,omitempty"`
	// The number of client control messages received from the Bridge. Available since 2.13.
	ControlRxMsgCount *int64 `json:"controlRxMsgCount,omitempty"`
	// The amount of client control messages transmitted to the Bridge, in bytes (B). Available since 2.13.
	ControlTxByteCount *int64 `json:"controlTxByteCount,omitempty"`
	// The number of client control messages transmitted to the Bridge. Available since 2.13.
	ControlTxMsgCount *int64               `json:"controlTxMsgCount,omitempty"`
	Counter           *MsgVpnBridgeCounter `json:"counter,omitempty"`
	// The amount of client data messages received from the Bridge, in bytes (B). Available since 2.13.
	DataRxByteCount *int64 `json:"dataRxByteCount,omitempty"`
	// The number of client data messages received from the Bridge. Available since 2.13.
	DataRxMsgCount *int64 `json:"dataRxMsgCount,omitempty"`
	// The amount of client data messages transmitted to the Bridge, in bytes (B). Available since 2.13.
	DataTxByteCount *int64 `json:"dataTxByteCount,omitempty"`
	// The number of client data messages transmitted to the Bridge. Available since 2.13.
	DataTxMsgCount *int64 `json:"dataTxMsgCount,omitempty"`
	// The number of messages discarded during reception from the Bridge. Available since 2.13.
	DiscardedRxMsgCount *int64 `json:"discardedRxMsgCount,omitempty"`
	// The number of messages discarded during transmission to the Bridge. Available since 2.13.
	DiscardedTxMsgCount *int64 `json:"discardedTxMsgCount,omitempty"`
	// Indicates whether the Bridge is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Indicates whether messages transmitted over the Bridge are encrypted with TLS.
	Encrypted *bool `json:"encrypted,omitempty"`
	// The establisher of the Bridge connection. The allowed values and their meaning are:  <pre> \"local\" - The Bridge connection was established by the local Message VPN. \"remote\" - The Bridge connection was established by the remote Message VPN. </pre>
	Establisher *string `json:"establisher,omitempty"`
	// The reason for the inbound connection failure from the Bridge. If there is no failure reason, an empty string (\"\") is returned.
	InboundFailureReason *string `json:"inboundFailureReason,omitempty"`
	// The state of the inbound connection from the Bridge. The allowed values and their meaning are:  <pre> \"init\" - The connection is initializing. \"disabled\" - The connection is disabled by configuration. \"enabled\" - The connection is enabled by configuration. \"prepare\" - The connection is operationally down. \"prepare-wait-to-connect\" - The connection is waiting to connect. \"prepare-fetching-dns\" - The domain name of the destination node is being resolved. \"not-ready\" - The connection is operationally down. \"not-ready-connecting\" - The connection is trying to connect. \"not-ready-handshaking\" - The connection is handshaking. \"not-ready-wait-next\" - The connection failed to connect and is waiting to retry. \"not-ready-wait-reuse\" - The connection is closing in order to reuse an existing connection. \"not-ready-wait-bridge-version-mismatch\" - The connection is closing because of a version mismatch. \"not-ready-wait-cleanup\" - The connection is closed and cleaning up. \"ready\" - The connection is operationally up. \"ready-subscribing\" - The connection is up and synchronizing subscriptions. \"ready-in-sync\" - The connection is up and subscriptions are synchronized. </pre>
	InboundState *string `json:"inboundState,omitempty"`
	// The ID of the last message transmitted to the Bridge.
	LastTxMsgId *int64 `json:"lastTxMsgId,omitempty"`
	// The physical interface on the local Message VPN host for connecting to the remote Message VPN.
	LocalInterface *string `json:"localInterface,omitempty"`
	// The name of the local queue for the Bridge.
	LocalQueueName *string `json:"localQueueName,omitempty"`
	// The number of login request messages received from the Bridge. Available since 2.13.
	LoginRxMsgCount *int64 `json:"loginRxMsgCount,omitempty"`
	// The number of login response messages transmitted to the Bridge. Available since 2.13.
	LoginTxMsgCount *int64 `json:"loginTxMsgCount,omitempty"`
	// The maximum time-to-live (TTL) in hops. Messages are discarded if their TTL exceeds this value.
	MaxTtl *int64 `json:"maxTtl,omitempty"`
	// The number of guaranteed messages received from the Bridge. Available since 2.13.
	MsgSpoolRxMsgCount *int64 `json:"msgSpoolRxMsgCount,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// The state of the outbound connection to the Bridge. The allowed values and their meaning are:  <pre> \"init\" - The connection is initializing. \"disabled\" - The connection is disabled by configuration. \"enabled\" - The connection is enabled by configuration. \"prepare\" - The connection is operationally down. \"prepare-wait-to-connect\" - The connection is waiting to connect. \"prepare-fetching-dns\" - The domain name of the destination node is being resolved. \"not-ready\" - The connection is operationally down. \"not-ready-connecting\" - The connection is trying to connect. \"not-ready-handshaking\" - The connection is handshaking. \"not-ready-wait-next\" - The connection failed to connect and is waiting to retry. \"not-ready-wait-reuse\" - The connection is closing in order to reuse an existing connection. \"not-ready-wait-bridge-version-mismatch\" - The connection is closing because of a version mismatch. \"not-ready-wait-cleanup\" - The connection is closed and cleaning up. \"ready\" - The connection is operationally up. \"ready-subscribing\" - The connection is up and synchronizing subscriptions. \"ready-in-sync\" - The connection is up and subscriptions are synchronized. </pre>
	OutboundState *string           `json:"outboundState,omitempty"`
	Rate          *MsgVpnBridgeRate `json:"rate,omitempty"`
	// The FQDN or IP address of the remote Message VPN.
	RemoteAddress *string `json:"remoteAddress,omitempty"`
	// The Client Username the Bridge uses to login to the remote Message VPN.
	RemoteAuthenticationBasicClientUsername *string `json:"remoteAuthenticationBasicClientUsername,omitempty"`
	// The authentication scheme for the remote Message VPN. The allowed values and their meaning are:  <pre> \"basic\" - Basic Authentication Scheme (via username and password). \"client-certificate\" - Client Certificate Authentication Scheme (via certificate file or content). </pre>
	RemoteAuthenticationScheme *string `json:"remoteAuthenticationScheme,omitempty"`
	// The maximum number of retry attempts to establish a connection to the remote Message VPN. A value of 0 means to retry forever.
	RemoteConnectionRetryCount *int64 `json:"remoteConnectionRetryCount,omitempty"`
	// The number of seconds the broker waits for the bridge connection to be established before attempting a new connection.
	RemoteConnectionRetryDelay *int64 `json:"remoteConnectionRetryDelay,omitempty"`
	// The priority for deliver-to-one (DTO) messages transmitted from the remote Message VPN. The allowed values and their meaning are:  <pre> \"p1\" - The 1st or highest priority. \"p2\" - The 2nd highest priority. \"p3\" - The 3rd highest priority. \"p4\" - The 4th highest priority. \"da\" - Ignore priority and deliver always. </pre>
	RemoteDeliverToOnePriority *string `json:"remoteDeliverToOnePriority,omitempty"`
	// The name of the remote Message VPN.
	RemoteMsgVpnName *string `json:"remoteMsgVpnName,omitempty"`
	// The name of the remote router.
	RemoteRouterName *string `json:"remoteRouterName,omitempty"`
	// The ID of the transmit flow for the connected remote Message VPN.
	RemoteTxFlowId *int32 `json:"remoteTxFlowId,omitempty"`
	// The amount of messages received from the Bridge, in bytes (B). Available since 2.13.
	RxByteCount *int64 `json:"rxByteCount,omitempty"`
	// The current message rate received from the Bridge, in bytes per second (B/sec). Available since 2.13.
	RxByteRate *int64 `json:"rxByteRate,omitempty"`
	// The category of the inbound connection failure from the Bridge. The allowed values and their meaning are:  <pre> \"no-failure\" - There is no bridge failure. \"local-configuration-problem\" - The bridge failure is a local configuration problem. \"local-operational-state-problem\" - The bridge failure is an operational state problem. </pre>  Available since 2.18.
	RxConnectionFailureCategory *string `json:"rxConnectionFailureCategory,omitempty"`
	// The number of messages received from the Bridge. Available since 2.13.
	RxMsgCount *int64 `json:"rxMsgCount,omitempty"`
	// The current message rate received from the Bridge, in messages per second (msg/sec). Available since 2.13.
	RxMsgRate *int64 `json:"rxMsgRate,omitempty"`
	// The colon-separated list of cipher suites supported for TLS connections to the remote Message VPN. The value \"default\" implies all supported suites ordered from most secure to least secure.
	TlsCipherSuiteList *string `json:"tlsCipherSuiteList,omitempty"`
	// Indicates whether the Bridge is configured to use the default cipher-suite list.
	TlsDefaultCipherSuiteList *bool `json:"tlsDefaultCipherSuiteList,omitempty"`
	// Indicates whether the TTL (hops) exceeded event has been raised.
	TtlExceededEventRaised *bool `json:"ttlExceededEventRaised,omitempty"`
	// The amount of messages transmitted to the Bridge, in bytes (B). Available since 2.13.
	TxByteCount *int64 `json:"txByteCount,omitempty"`
	// The current message rate transmitted to the Bridge, in bytes per second (B/sec). Available since 2.13.
	TxByteRate *int64 `json:"txByteRate,omitempty"`
	// The number of messages transmitted to the Bridge. Available since 2.13.
	TxMsgCount *int64 `json:"txMsgCount,omitempty"`
	// The current message rate transmitted to the Bridge, in messages per second (msg/sec). Available since 2.13.
	TxMsgRate *int64 `json:"txMsgRate,omitempty"`
	// The amount of time in seconds since the Bridge connected to the remote Message VPN.
	Uptime *int64 `json:"uptime,omitempty"`
}

// NewMsgVpnBridge instantiates a new MsgVpnBridge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnBridge() *MsgVpnBridge {
	this := MsgVpnBridge{}
	return &this
}

// NewMsgVpnBridgeWithDefaults instantiates a new MsgVpnBridge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnBridgeWithDefaults() *MsgVpnBridge {
	this := MsgVpnBridge{}
	return &this
}

// GetAverageRxByteRate returns the AverageRxByteRate field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetAverageRxByteRate() int64 {
	if o == nil || o.AverageRxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRxByteRate
}

// GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetAverageRxByteRateOk() (*int64, bool) {
	if o == nil || o.AverageRxByteRate == nil {
		return nil, false
	}
	return o.AverageRxByteRate, true
}

// HasAverageRxByteRate returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasAverageRxByteRate() bool {
	if o != nil && o.AverageRxByteRate != nil {
		return true
	}

	return false
}

// SetAverageRxByteRate gets a reference to the given int64 and assigns it to the AverageRxByteRate field.
func (o *MsgVpnBridge) SetAverageRxByteRate(v int64) {
	o.AverageRxByteRate = &v
}

// GetAverageRxMsgRate returns the AverageRxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetAverageRxMsgRate() int64 {
	if o == nil || o.AverageRxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRxMsgRate
}

// GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetAverageRxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageRxMsgRate == nil {
		return nil, false
	}
	return o.AverageRxMsgRate, true
}

// HasAverageRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasAverageRxMsgRate() bool {
	if o != nil && o.AverageRxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageRxMsgRate gets a reference to the given int64 and assigns it to the AverageRxMsgRate field.
func (o *MsgVpnBridge) SetAverageRxMsgRate(v int64) {
	o.AverageRxMsgRate = &v
}

// GetAverageTxByteRate returns the AverageTxByteRate field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetAverageTxByteRate() int64 {
	if o == nil || o.AverageTxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageTxByteRate
}

// GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetAverageTxByteRateOk() (*int64, bool) {
	if o == nil || o.AverageTxByteRate == nil {
		return nil, false
	}
	return o.AverageTxByteRate, true
}

// HasAverageTxByteRate returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasAverageTxByteRate() bool {
	if o != nil && o.AverageTxByteRate != nil {
		return true
	}

	return false
}

// SetAverageTxByteRate gets a reference to the given int64 and assigns it to the AverageTxByteRate field.
func (o *MsgVpnBridge) SetAverageTxByteRate(v int64) {
	o.AverageTxByteRate = &v
}

// GetAverageTxMsgRate returns the AverageTxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetAverageTxMsgRate() int64 {
	if o == nil || o.AverageTxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageTxMsgRate
}

// GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetAverageTxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageTxMsgRate == nil {
		return nil, false
	}
	return o.AverageTxMsgRate, true
}

// HasAverageTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasAverageTxMsgRate() bool {
	if o != nil && o.AverageTxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageTxMsgRate gets a reference to the given int64 and assigns it to the AverageTxMsgRate field.
func (o *MsgVpnBridge) SetAverageTxMsgRate(v int64) {
	o.AverageTxMsgRate = &v
}

// GetBoundToQueue returns the BoundToQueue field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetBoundToQueue() bool {
	if o == nil || o.BoundToQueue == nil {
		var ret bool
		return ret
	}
	return *o.BoundToQueue
}

// GetBoundToQueueOk returns a tuple with the BoundToQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetBoundToQueueOk() (*bool, bool) {
	if o == nil || o.BoundToQueue == nil {
		return nil, false
	}
	return o.BoundToQueue, true
}

// HasBoundToQueue returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasBoundToQueue() bool {
	if o != nil && o.BoundToQueue != nil {
		return true
	}

	return false
}

// SetBoundToQueue gets a reference to the given bool and assigns it to the BoundToQueue field.
func (o *MsgVpnBridge) SetBoundToQueue(v bool) {
	o.BoundToQueue = &v
}

// GetBridgeName returns the BridgeName field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetBridgeName() string {
	if o == nil || o.BridgeName == nil {
		var ret string
		return ret
	}
	return *o.BridgeName
}

// GetBridgeNameOk returns a tuple with the BridgeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetBridgeNameOk() (*string, bool) {
	if o == nil || o.BridgeName == nil {
		return nil, false
	}
	return o.BridgeName, true
}

// HasBridgeName returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasBridgeName() bool {
	if o != nil && o.BridgeName != nil {
		return true
	}

	return false
}

// SetBridgeName gets a reference to the given string and assigns it to the BridgeName field.
func (o *MsgVpnBridge) SetBridgeName(v string) {
	o.BridgeName = &v
}

// GetBridgeVirtualRouter returns the BridgeVirtualRouter field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetBridgeVirtualRouter() string {
	if o == nil || o.BridgeVirtualRouter == nil {
		var ret string
		return ret
	}
	return *o.BridgeVirtualRouter
}

// GetBridgeVirtualRouterOk returns a tuple with the BridgeVirtualRouter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetBridgeVirtualRouterOk() (*string, bool) {
	if o == nil || o.BridgeVirtualRouter == nil {
		return nil, false
	}
	return o.BridgeVirtualRouter, true
}

// HasBridgeVirtualRouter returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasBridgeVirtualRouter() bool {
	if o != nil && o.BridgeVirtualRouter != nil {
		return true
	}

	return false
}

// SetBridgeVirtualRouter gets a reference to the given string and assigns it to the BridgeVirtualRouter field.
func (o *MsgVpnBridge) SetBridgeVirtualRouter(v string) {
	o.BridgeVirtualRouter = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *MsgVpnBridge) SetClientName(v string) {
	o.ClientName = &v
}

// GetCompressed returns the Compressed field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetCompressed() bool {
	if o == nil || o.Compressed == nil {
		var ret bool
		return ret
	}
	return *o.Compressed
}

// GetCompressedOk returns a tuple with the Compressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetCompressedOk() (*bool, bool) {
	if o == nil || o.Compressed == nil {
		return nil, false
	}
	return o.Compressed, true
}

// HasCompressed returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasCompressed() bool {
	if o != nil && o.Compressed != nil {
		return true
	}

	return false
}

// SetCompressed gets a reference to the given bool and assigns it to the Compressed field.
func (o *MsgVpnBridge) SetCompressed(v bool) {
	o.Compressed = &v
}

// GetControlRxByteCount returns the ControlRxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetControlRxByteCount() int64 {
	if o == nil || o.ControlRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlRxByteCount
}

// GetControlRxByteCountOk returns a tuple with the ControlRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetControlRxByteCountOk() (*int64, bool) {
	if o == nil || o.ControlRxByteCount == nil {
		return nil, false
	}
	return o.ControlRxByteCount, true
}

// HasControlRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasControlRxByteCount() bool {
	if o != nil && o.ControlRxByteCount != nil {
		return true
	}

	return false
}

// SetControlRxByteCount gets a reference to the given int64 and assigns it to the ControlRxByteCount field.
func (o *MsgVpnBridge) SetControlRxByteCount(v int64) {
	o.ControlRxByteCount = &v
}

// GetControlRxMsgCount returns the ControlRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetControlRxMsgCount() int64 {
	if o == nil || o.ControlRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlRxMsgCount
}

// GetControlRxMsgCountOk returns a tuple with the ControlRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetControlRxMsgCountOk() (*int64, bool) {
	if o == nil || o.ControlRxMsgCount == nil {
		return nil, false
	}
	return o.ControlRxMsgCount, true
}

// HasControlRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasControlRxMsgCount() bool {
	if o != nil && o.ControlRxMsgCount != nil {
		return true
	}

	return false
}

// SetControlRxMsgCount gets a reference to the given int64 and assigns it to the ControlRxMsgCount field.
func (o *MsgVpnBridge) SetControlRxMsgCount(v int64) {
	o.ControlRxMsgCount = &v
}

// GetControlTxByteCount returns the ControlTxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetControlTxByteCount() int64 {
	if o == nil || o.ControlTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlTxByteCount
}

// GetControlTxByteCountOk returns a tuple with the ControlTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetControlTxByteCountOk() (*int64, bool) {
	if o == nil || o.ControlTxByteCount == nil {
		return nil, false
	}
	return o.ControlTxByteCount, true
}

// HasControlTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasControlTxByteCount() bool {
	if o != nil && o.ControlTxByteCount != nil {
		return true
	}

	return false
}

// SetControlTxByteCount gets a reference to the given int64 and assigns it to the ControlTxByteCount field.
func (o *MsgVpnBridge) SetControlTxByteCount(v int64) {
	o.ControlTxByteCount = &v
}

// GetControlTxMsgCount returns the ControlTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetControlTxMsgCount() int64 {
	if o == nil || o.ControlTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlTxMsgCount
}

// GetControlTxMsgCountOk returns a tuple with the ControlTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetControlTxMsgCountOk() (*int64, bool) {
	if o == nil || o.ControlTxMsgCount == nil {
		return nil, false
	}
	return o.ControlTxMsgCount, true
}

// HasControlTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasControlTxMsgCount() bool {
	if o != nil && o.ControlTxMsgCount != nil {
		return true
	}

	return false
}

// SetControlTxMsgCount gets a reference to the given int64 and assigns it to the ControlTxMsgCount field.
func (o *MsgVpnBridge) SetControlTxMsgCount(v int64) {
	o.ControlTxMsgCount = &v
}

// GetCounter returns the Counter field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetCounter() MsgVpnBridgeCounter {
	if o == nil || o.Counter == nil {
		var ret MsgVpnBridgeCounter
		return ret
	}
	return *o.Counter
}

// GetCounterOk returns a tuple with the Counter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetCounterOk() (*MsgVpnBridgeCounter, bool) {
	if o == nil || o.Counter == nil {
		return nil, false
	}
	return o.Counter, true
}

// HasCounter returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasCounter() bool {
	if o != nil && o.Counter != nil {
		return true
	}

	return false
}

// SetCounter gets a reference to the given MsgVpnBridgeCounter and assigns it to the Counter field.
func (o *MsgVpnBridge) SetCounter(v MsgVpnBridgeCounter) {
	o.Counter = &v
}

// GetDataRxByteCount returns the DataRxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetDataRxByteCount() int64 {
	if o == nil || o.DataRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.DataRxByteCount
}

// GetDataRxByteCountOk returns a tuple with the DataRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetDataRxByteCountOk() (*int64, bool) {
	if o == nil || o.DataRxByteCount == nil {
		return nil, false
	}
	return o.DataRxByteCount, true
}

// HasDataRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasDataRxByteCount() bool {
	if o != nil && o.DataRxByteCount != nil {
		return true
	}

	return false
}

// SetDataRxByteCount gets a reference to the given int64 and assigns it to the DataRxByteCount field.
func (o *MsgVpnBridge) SetDataRxByteCount(v int64) {
	o.DataRxByteCount = &v
}

// GetDataRxMsgCount returns the DataRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetDataRxMsgCount() int64 {
	if o == nil || o.DataRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DataRxMsgCount
}

// GetDataRxMsgCountOk returns a tuple with the DataRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetDataRxMsgCountOk() (*int64, bool) {
	if o == nil || o.DataRxMsgCount == nil {
		return nil, false
	}
	return o.DataRxMsgCount, true
}

// HasDataRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasDataRxMsgCount() bool {
	if o != nil && o.DataRxMsgCount != nil {
		return true
	}

	return false
}

// SetDataRxMsgCount gets a reference to the given int64 and assigns it to the DataRxMsgCount field.
func (o *MsgVpnBridge) SetDataRxMsgCount(v int64) {
	o.DataRxMsgCount = &v
}

// GetDataTxByteCount returns the DataTxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetDataTxByteCount() int64 {
	if o == nil || o.DataTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.DataTxByteCount
}

// GetDataTxByteCountOk returns a tuple with the DataTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetDataTxByteCountOk() (*int64, bool) {
	if o == nil || o.DataTxByteCount == nil {
		return nil, false
	}
	return o.DataTxByteCount, true
}

// HasDataTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasDataTxByteCount() bool {
	if o != nil && o.DataTxByteCount != nil {
		return true
	}

	return false
}

// SetDataTxByteCount gets a reference to the given int64 and assigns it to the DataTxByteCount field.
func (o *MsgVpnBridge) SetDataTxByteCount(v int64) {
	o.DataTxByteCount = &v
}

// GetDataTxMsgCount returns the DataTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetDataTxMsgCount() int64 {
	if o == nil || o.DataTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DataTxMsgCount
}

// GetDataTxMsgCountOk returns a tuple with the DataTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetDataTxMsgCountOk() (*int64, bool) {
	if o == nil || o.DataTxMsgCount == nil {
		return nil, false
	}
	return o.DataTxMsgCount, true
}

// HasDataTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasDataTxMsgCount() bool {
	if o != nil && o.DataTxMsgCount != nil {
		return true
	}

	return false
}

// SetDataTxMsgCount gets a reference to the given int64 and assigns it to the DataTxMsgCount field.
func (o *MsgVpnBridge) SetDataTxMsgCount(v int64) {
	o.DataTxMsgCount = &v
}

// GetDiscardedRxMsgCount returns the DiscardedRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetDiscardedRxMsgCount() int64 {
	if o == nil || o.DiscardedRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DiscardedRxMsgCount
}

// GetDiscardedRxMsgCountOk returns a tuple with the DiscardedRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetDiscardedRxMsgCountOk() (*int64, bool) {
	if o == nil || o.DiscardedRxMsgCount == nil {
		return nil, false
	}
	return o.DiscardedRxMsgCount, true
}

// HasDiscardedRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasDiscardedRxMsgCount() bool {
	if o != nil && o.DiscardedRxMsgCount != nil {
		return true
	}

	return false
}

// SetDiscardedRxMsgCount gets a reference to the given int64 and assigns it to the DiscardedRxMsgCount field.
func (o *MsgVpnBridge) SetDiscardedRxMsgCount(v int64) {
	o.DiscardedRxMsgCount = &v
}

// GetDiscardedTxMsgCount returns the DiscardedTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetDiscardedTxMsgCount() int64 {
	if o == nil || o.DiscardedTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DiscardedTxMsgCount
}

// GetDiscardedTxMsgCountOk returns a tuple with the DiscardedTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetDiscardedTxMsgCountOk() (*int64, bool) {
	if o == nil || o.DiscardedTxMsgCount == nil {
		return nil, false
	}
	return o.DiscardedTxMsgCount, true
}

// HasDiscardedTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasDiscardedTxMsgCount() bool {
	if o != nil && o.DiscardedTxMsgCount != nil {
		return true
	}

	return false
}

// SetDiscardedTxMsgCount gets a reference to the given int64 and assigns it to the DiscardedTxMsgCount field.
func (o *MsgVpnBridge) SetDiscardedTxMsgCount(v int64) {
	o.DiscardedTxMsgCount = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MsgVpnBridge) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetEncrypted() bool {
	if o == nil || o.Encrypted == nil {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetEncryptedOk() (*bool, bool) {
	if o == nil || o.Encrypted == nil {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasEncrypted() bool {
	if o != nil && o.Encrypted != nil {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *MsgVpnBridge) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetEstablisher returns the Establisher field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetEstablisher() string {
	if o == nil || o.Establisher == nil {
		var ret string
		return ret
	}
	return *o.Establisher
}

// GetEstablisherOk returns a tuple with the Establisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetEstablisherOk() (*string, bool) {
	if o == nil || o.Establisher == nil {
		return nil, false
	}
	return o.Establisher, true
}

// HasEstablisher returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasEstablisher() bool {
	if o != nil && o.Establisher != nil {
		return true
	}

	return false
}

// SetEstablisher gets a reference to the given string and assigns it to the Establisher field.
func (o *MsgVpnBridge) SetEstablisher(v string) {
	o.Establisher = &v
}

// GetInboundFailureReason returns the InboundFailureReason field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetInboundFailureReason() string {
	if o == nil || o.InboundFailureReason == nil {
		var ret string
		return ret
	}
	return *o.InboundFailureReason
}

// GetInboundFailureReasonOk returns a tuple with the InboundFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetInboundFailureReasonOk() (*string, bool) {
	if o == nil || o.InboundFailureReason == nil {
		return nil, false
	}
	return o.InboundFailureReason, true
}

// HasInboundFailureReason returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasInboundFailureReason() bool {
	if o != nil && o.InboundFailureReason != nil {
		return true
	}

	return false
}

// SetInboundFailureReason gets a reference to the given string and assigns it to the InboundFailureReason field.
func (o *MsgVpnBridge) SetInboundFailureReason(v string) {
	o.InboundFailureReason = &v
}

// GetInboundState returns the InboundState field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetInboundState() string {
	if o == nil || o.InboundState == nil {
		var ret string
		return ret
	}
	return *o.InboundState
}

// GetInboundStateOk returns a tuple with the InboundState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetInboundStateOk() (*string, bool) {
	if o == nil || o.InboundState == nil {
		return nil, false
	}
	return o.InboundState, true
}

// HasInboundState returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasInboundState() bool {
	if o != nil && o.InboundState != nil {
		return true
	}

	return false
}

// SetInboundState gets a reference to the given string and assigns it to the InboundState field.
func (o *MsgVpnBridge) SetInboundState(v string) {
	o.InboundState = &v
}

// GetLastTxMsgId returns the LastTxMsgId field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetLastTxMsgId() int64 {
	if o == nil || o.LastTxMsgId == nil {
		var ret int64
		return ret
	}
	return *o.LastTxMsgId
}

// GetLastTxMsgIdOk returns a tuple with the LastTxMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetLastTxMsgIdOk() (*int64, bool) {
	if o == nil || o.LastTxMsgId == nil {
		return nil, false
	}
	return o.LastTxMsgId, true
}

// HasLastTxMsgId returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasLastTxMsgId() bool {
	if o != nil && o.LastTxMsgId != nil {
		return true
	}

	return false
}

// SetLastTxMsgId gets a reference to the given int64 and assigns it to the LastTxMsgId field.
func (o *MsgVpnBridge) SetLastTxMsgId(v int64) {
	o.LastTxMsgId = &v
}

// GetLocalInterface returns the LocalInterface field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetLocalInterface() string {
	if o == nil || o.LocalInterface == nil {
		var ret string
		return ret
	}
	return *o.LocalInterface
}

// GetLocalInterfaceOk returns a tuple with the LocalInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetLocalInterfaceOk() (*string, bool) {
	if o == nil || o.LocalInterface == nil {
		return nil, false
	}
	return o.LocalInterface, true
}

// HasLocalInterface returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasLocalInterface() bool {
	if o != nil && o.LocalInterface != nil {
		return true
	}

	return false
}

// SetLocalInterface gets a reference to the given string and assigns it to the LocalInterface field.
func (o *MsgVpnBridge) SetLocalInterface(v string) {
	o.LocalInterface = &v
}

// GetLocalQueueName returns the LocalQueueName field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetLocalQueueName() string {
	if o == nil || o.LocalQueueName == nil {
		var ret string
		return ret
	}
	return *o.LocalQueueName
}

// GetLocalQueueNameOk returns a tuple with the LocalQueueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetLocalQueueNameOk() (*string, bool) {
	if o == nil || o.LocalQueueName == nil {
		return nil, false
	}
	return o.LocalQueueName, true
}

// HasLocalQueueName returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasLocalQueueName() bool {
	if o != nil && o.LocalQueueName != nil {
		return true
	}

	return false
}

// SetLocalQueueName gets a reference to the given string and assigns it to the LocalQueueName field.
func (o *MsgVpnBridge) SetLocalQueueName(v string) {
	o.LocalQueueName = &v
}

// GetLoginRxMsgCount returns the LoginRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetLoginRxMsgCount() int64 {
	if o == nil || o.LoginRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginRxMsgCount
}

// GetLoginRxMsgCountOk returns a tuple with the LoginRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetLoginRxMsgCountOk() (*int64, bool) {
	if o == nil || o.LoginRxMsgCount == nil {
		return nil, false
	}
	return o.LoginRxMsgCount, true
}

// HasLoginRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasLoginRxMsgCount() bool {
	if o != nil && o.LoginRxMsgCount != nil {
		return true
	}

	return false
}

// SetLoginRxMsgCount gets a reference to the given int64 and assigns it to the LoginRxMsgCount field.
func (o *MsgVpnBridge) SetLoginRxMsgCount(v int64) {
	o.LoginRxMsgCount = &v
}

// GetLoginTxMsgCount returns the LoginTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetLoginTxMsgCount() int64 {
	if o == nil || o.LoginTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginTxMsgCount
}

// GetLoginTxMsgCountOk returns a tuple with the LoginTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetLoginTxMsgCountOk() (*int64, bool) {
	if o == nil || o.LoginTxMsgCount == nil {
		return nil, false
	}
	return o.LoginTxMsgCount, true
}

// HasLoginTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasLoginTxMsgCount() bool {
	if o != nil && o.LoginTxMsgCount != nil {
		return true
	}

	return false
}

// SetLoginTxMsgCount gets a reference to the given int64 and assigns it to the LoginTxMsgCount field.
func (o *MsgVpnBridge) SetLoginTxMsgCount(v int64) {
	o.LoginTxMsgCount = &v
}

// GetMaxTtl returns the MaxTtl field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetMaxTtl() int64 {
	if o == nil || o.MaxTtl == nil {
		var ret int64
		return ret
	}
	return *o.MaxTtl
}

// GetMaxTtlOk returns a tuple with the MaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetMaxTtlOk() (*int64, bool) {
	if o == nil || o.MaxTtl == nil {
		return nil, false
	}
	return o.MaxTtl, true
}

// HasMaxTtl returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasMaxTtl() bool {
	if o != nil && o.MaxTtl != nil {
		return true
	}

	return false
}

// SetMaxTtl gets a reference to the given int64 and assigns it to the MaxTtl field.
func (o *MsgVpnBridge) SetMaxTtl(v int64) {
	o.MaxTtl = &v
}

// GetMsgSpoolRxMsgCount returns the MsgSpoolRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetMsgSpoolRxMsgCount() int64 {
	if o == nil || o.MsgSpoolRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MsgSpoolRxMsgCount
}

// GetMsgSpoolRxMsgCountOk returns a tuple with the MsgSpoolRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetMsgSpoolRxMsgCountOk() (*int64, bool) {
	if o == nil || o.MsgSpoolRxMsgCount == nil {
		return nil, false
	}
	return o.MsgSpoolRxMsgCount, true
}

// HasMsgSpoolRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasMsgSpoolRxMsgCount() bool {
	if o != nil && o.MsgSpoolRxMsgCount != nil {
		return true
	}

	return false
}

// SetMsgSpoolRxMsgCount gets a reference to the given int64 and assigns it to the MsgSpoolRxMsgCount field.
func (o *MsgVpnBridge) SetMsgSpoolRxMsgCount(v int64) {
	o.MsgSpoolRxMsgCount = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnBridge) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetOutboundState returns the OutboundState field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetOutboundState() string {
	if o == nil || o.OutboundState == nil {
		var ret string
		return ret
	}
	return *o.OutboundState
}

// GetOutboundStateOk returns a tuple with the OutboundState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetOutboundStateOk() (*string, bool) {
	if o == nil || o.OutboundState == nil {
		return nil, false
	}
	return o.OutboundState, true
}

// HasOutboundState returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasOutboundState() bool {
	if o != nil && o.OutboundState != nil {
		return true
	}

	return false
}

// SetOutboundState gets a reference to the given string and assigns it to the OutboundState field.
func (o *MsgVpnBridge) SetOutboundState(v string) {
	o.OutboundState = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRate() MsgVpnBridgeRate {
	if o == nil || o.Rate == nil {
		var ret MsgVpnBridgeRate
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRateOk() (*MsgVpnBridgeRate, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given MsgVpnBridgeRate and assigns it to the Rate field.
func (o *MsgVpnBridge) SetRate(v MsgVpnBridgeRate) {
	o.Rate = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteAddress() string {
	if o == nil || o.RemoteAddress == nil {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteAddressOk() (*string, bool) {
	if o == nil || o.RemoteAddress == nil {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteAddress() bool {
	if o != nil && o.RemoteAddress != nil {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *MsgVpnBridge) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetRemoteAuthenticationBasicClientUsername returns the RemoteAuthenticationBasicClientUsername field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteAuthenticationBasicClientUsername() string {
	if o == nil || o.RemoteAuthenticationBasicClientUsername == nil {
		var ret string
		return ret
	}
	return *o.RemoteAuthenticationBasicClientUsername
}

// GetRemoteAuthenticationBasicClientUsernameOk returns a tuple with the RemoteAuthenticationBasicClientUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteAuthenticationBasicClientUsernameOk() (*string, bool) {
	if o == nil || o.RemoteAuthenticationBasicClientUsername == nil {
		return nil, false
	}
	return o.RemoteAuthenticationBasicClientUsername, true
}

// HasRemoteAuthenticationBasicClientUsername returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteAuthenticationBasicClientUsername() bool {
	if o != nil && o.RemoteAuthenticationBasicClientUsername != nil {
		return true
	}

	return false
}

// SetRemoteAuthenticationBasicClientUsername gets a reference to the given string and assigns it to the RemoteAuthenticationBasicClientUsername field.
func (o *MsgVpnBridge) SetRemoteAuthenticationBasicClientUsername(v string) {
	o.RemoteAuthenticationBasicClientUsername = &v
}

// GetRemoteAuthenticationScheme returns the RemoteAuthenticationScheme field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteAuthenticationScheme() string {
	if o == nil || o.RemoteAuthenticationScheme == nil {
		var ret string
		return ret
	}
	return *o.RemoteAuthenticationScheme
}

// GetRemoteAuthenticationSchemeOk returns a tuple with the RemoteAuthenticationScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteAuthenticationSchemeOk() (*string, bool) {
	if o == nil || o.RemoteAuthenticationScheme == nil {
		return nil, false
	}
	return o.RemoteAuthenticationScheme, true
}

// HasRemoteAuthenticationScheme returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteAuthenticationScheme() bool {
	if o != nil && o.RemoteAuthenticationScheme != nil {
		return true
	}

	return false
}

// SetRemoteAuthenticationScheme gets a reference to the given string and assigns it to the RemoteAuthenticationScheme field.
func (o *MsgVpnBridge) SetRemoteAuthenticationScheme(v string) {
	o.RemoteAuthenticationScheme = &v
}

// GetRemoteConnectionRetryCount returns the RemoteConnectionRetryCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteConnectionRetryCount() int64 {
	if o == nil || o.RemoteConnectionRetryCount == nil {
		var ret int64
		return ret
	}
	return *o.RemoteConnectionRetryCount
}

// GetRemoteConnectionRetryCountOk returns a tuple with the RemoteConnectionRetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteConnectionRetryCountOk() (*int64, bool) {
	if o == nil || o.RemoteConnectionRetryCount == nil {
		return nil, false
	}
	return o.RemoteConnectionRetryCount, true
}

// HasRemoteConnectionRetryCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteConnectionRetryCount() bool {
	if o != nil && o.RemoteConnectionRetryCount != nil {
		return true
	}

	return false
}

// SetRemoteConnectionRetryCount gets a reference to the given int64 and assigns it to the RemoteConnectionRetryCount field.
func (o *MsgVpnBridge) SetRemoteConnectionRetryCount(v int64) {
	o.RemoteConnectionRetryCount = &v
}

// GetRemoteConnectionRetryDelay returns the RemoteConnectionRetryDelay field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteConnectionRetryDelay() int64 {
	if o == nil || o.RemoteConnectionRetryDelay == nil {
		var ret int64
		return ret
	}
	return *o.RemoteConnectionRetryDelay
}

// GetRemoteConnectionRetryDelayOk returns a tuple with the RemoteConnectionRetryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteConnectionRetryDelayOk() (*int64, bool) {
	if o == nil || o.RemoteConnectionRetryDelay == nil {
		return nil, false
	}
	return o.RemoteConnectionRetryDelay, true
}

// HasRemoteConnectionRetryDelay returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteConnectionRetryDelay() bool {
	if o != nil && o.RemoteConnectionRetryDelay != nil {
		return true
	}

	return false
}

// SetRemoteConnectionRetryDelay gets a reference to the given int64 and assigns it to the RemoteConnectionRetryDelay field.
func (o *MsgVpnBridge) SetRemoteConnectionRetryDelay(v int64) {
	o.RemoteConnectionRetryDelay = &v
}

// GetRemoteDeliverToOnePriority returns the RemoteDeliverToOnePriority field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteDeliverToOnePriority() string {
	if o == nil || o.RemoteDeliverToOnePriority == nil {
		var ret string
		return ret
	}
	return *o.RemoteDeliverToOnePriority
}

// GetRemoteDeliverToOnePriorityOk returns a tuple with the RemoteDeliverToOnePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteDeliverToOnePriorityOk() (*string, bool) {
	if o == nil || o.RemoteDeliverToOnePriority == nil {
		return nil, false
	}
	return o.RemoteDeliverToOnePriority, true
}

// HasRemoteDeliverToOnePriority returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteDeliverToOnePriority() bool {
	if o != nil && o.RemoteDeliverToOnePriority != nil {
		return true
	}

	return false
}

// SetRemoteDeliverToOnePriority gets a reference to the given string and assigns it to the RemoteDeliverToOnePriority field.
func (o *MsgVpnBridge) SetRemoteDeliverToOnePriority(v string) {
	o.RemoteDeliverToOnePriority = &v
}

// GetRemoteMsgVpnName returns the RemoteMsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteMsgVpnName() string {
	if o == nil || o.RemoteMsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.RemoteMsgVpnName
}

// GetRemoteMsgVpnNameOk returns a tuple with the RemoteMsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteMsgVpnNameOk() (*string, bool) {
	if o == nil || o.RemoteMsgVpnName == nil {
		return nil, false
	}
	return o.RemoteMsgVpnName, true
}

// HasRemoteMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteMsgVpnName() bool {
	if o != nil && o.RemoteMsgVpnName != nil {
		return true
	}

	return false
}

// SetRemoteMsgVpnName gets a reference to the given string and assigns it to the RemoteMsgVpnName field.
func (o *MsgVpnBridge) SetRemoteMsgVpnName(v string) {
	o.RemoteMsgVpnName = &v
}

// GetRemoteRouterName returns the RemoteRouterName field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteRouterName() string {
	if o == nil || o.RemoteRouterName == nil {
		var ret string
		return ret
	}
	return *o.RemoteRouterName
}

// GetRemoteRouterNameOk returns a tuple with the RemoteRouterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteRouterNameOk() (*string, bool) {
	if o == nil || o.RemoteRouterName == nil {
		return nil, false
	}
	return o.RemoteRouterName, true
}

// HasRemoteRouterName returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteRouterName() bool {
	if o != nil && o.RemoteRouterName != nil {
		return true
	}

	return false
}

// SetRemoteRouterName gets a reference to the given string and assigns it to the RemoteRouterName field.
func (o *MsgVpnBridge) SetRemoteRouterName(v string) {
	o.RemoteRouterName = &v
}

// GetRemoteTxFlowId returns the RemoteTxFlowId field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteTxFlowId() int32 {
	if o == nil || o.RemoteTxFlowId == nil {
		var ret int32
		return ret
	}
	return *o.RemoteTxFlowId
}

// GetRemoteTxFlowIdOk returns a tuple with the RemoteTxFlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteTxFlowIdOk() (*int32, bool) {
	if o == nil || o.RemoteTxFlowId == nil {
		return nil, false
	}
	return o.RemoteTxFlowId, true
}

// HasRemoteTxFlowId returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteTxFlowId() bool {
	if o != nil && o.RemoteTxFlowId != nil {
		return true
	}

	return false
}

// SetRemoteTxFlowId gets a reference to the given int32 and assigns it to the RemoteTxFlowId field.
func (o *MsgVpnBridge) SetRemoteTxFlowId(v int32) {
	o.RemoteTxFlowId = &v
}

// GetRxByteCount returns the RxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRxByteCount() int64 {
	if o == nil || o.RxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.RxByteCount
}

// GetRxByteCountOk returns a tuple with the RxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRxByteCountOk() (*int64, bool) {
	if o == nil || o.RxByteCount == nil {
		return nil, false
	}
	return o.RxByteCount, true
}

// HasRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRxByteCount() bool {
	if o != nil && o.RxByteCount != nil {
		return true
	}

	return false
}

// SetRxByteCount gets a reference to the given int64 and assigns it to the RxByteCount field.
func (o *MsgVpnBridge) SetRxByteCount(v int64) {
	o.RxByteCount = &v
}

// GetRxByteRate returns the RxByteRate field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRxByteRate() int64 {
	if o == nil || o.RxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.RxByteRate
}

// GetRxByteRateOk returns a tuple with the RxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRxByteRateOk() (*int64, bool) {
	if o == nil || o.RxByteRate == nil {
		return nil, false
	}
	return o.RxByteRate, true
}

// HasRxByteRate returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRxByteRate() bool {
	if o != nil && o.RxByteRate != nil {
		return true
	}

	return false
}

// SetRxByteRate gets a reference to the given int64 and assigns it to the RxByteRate field.
func (o *MsgVpnBridge) SetRxByteRate(v int64) {
	o.RxByteRate = &v
}

// GetRxConnectionFailureCategory returns the RxConnectionFailureCategory field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRxConnectionFailureCategory() string {
	if o == nil || o.RxConnectionFailureCategory == nil {
		var ret string
		return ret
	}
	return *o.RxConnectionFailureCategory
}

// GetRxConnectionFailureCategoryOk returns a tuple with the RxConnectionFailureCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRxConnectionFailureCategoryOk() (*string, bool) {
	if o == nil || o.RxConnectionFailureCategory == nil {
		return nil, false
	}
	return o.RxConnectionFailureCategory, true
}

// HasRxConnectionFailureCategory returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRxConnectionFailureCategory() bool {
	if o != nil && o.RxConnectionFailureCategory != nil {
		return true
	}

	return false
}

// SetRxConnectionFailureCategory gets a reference to the given string and assigns it to the RxConnectionFailureCategory field.
func (o *MsgVpnBridge) SetRxConnectionFailureCategory(v string) {
	o.RxConnectionFailureCategory = &v
}

// GetRxMsgCount returns the RxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRxMsgCount() int64 {
	if o == nil || o.RxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RxMsgCount
}

// GetRxMsgCountOk returns a tuple with the RxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRxMsgCountOk() (*int64, bool) {
	if o == nil || o.RxMsgCount == nil {
		return nil, false
	}
	return o.RxMsgCount, true
}

// HasRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRxMsgCount() bool {
	if o != nil && o.RxMsgCount != nil {
		return true
	}

	return false
}

// SetRxMsgCount gets a reference to the given int64 and assigns it to the RxMsgCount field.
func (o *MsgVpnBridge) SetRxMsgCount(v int64) {
	o.RxMsgCount = &v
}

// GetRxMsgRate returns the RxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRxMsgRate() int64 {
	if o == nil || o.RxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.RxMsgRate
}

// GetRxMsgRateOk returns a tuple with the RxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRxMsgRateOk() (*int64, bool) {
	if o == nil || o.RxMsgRate == nil {
		return nil, false
	}
	return o.RxMsgRate, true
}

// HasRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRxMsgRate() bool {
	if o != nil && o.RxMsgRate != nil {
		return true
	}

	return false
}

// SetRxMsgRate gets a reference to the given int64 and assigns it to the RxMsgRate field.
func (o *MsgVpnBridge) SetRxMsgRate(v int64) {
	o.RxMsgRate = &v
}

// GetTlsCipherSuiteList returns the TlsCipherSuiteList field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetTlsCipherSuiteList() string {
	if o == nil || o.TlsCipherSuiteList == nil {
		var ret string
		return ret
	}
	return *o.TlsCipherSuiteList
}

// GetTlsCipherSuiteListOk returns a tuple with the TlsCipherSuiteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetTlsCipherSuiteListOk() (*string, bool) {
	if o == nil || o.TlsCipherSuiteList == nil {
		return nil, false
	}
	return o.TlsCipherSuiteList, true
}

// HasTlsCipherSuiteList returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasTlsCipherSuiteList() bool {
	if o != nil && o.TlsCipherSuiteList != nil {
		return true
	}

	return false
}

// SetTlsCipherSuiteList gets a reference to the given string and assigns it to the TlsCipherSuiteList field.
func (o *MsgVpnBridge) SetTlsCipherSuiteList(v string) {
	o.TlsCipherSuiteList = &v
}

// GetTlsDefaultCipherSuiteList returns the TlsDefaultCipherSuiteList field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetTlsDefaultCipherSuiteList() bool {
	if o == nil || o.TlsDefaultCipherSuiteList == nil {
		var ret bool
		return ret
	}
	return *o.TlsDefaultCipherSuiteList
}

// GetTlsDefaultCipherSuiteListOk returns a tuple with the TlsDefaultCipherSuiteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetTlsDefaultCipherSuiteListOk() (*bool, bool) {
	if o == nil || o.TlsDefaultCipherSuiteList == nil {
		return nil, false
	}
	return o.TlsDefaultCipherSuiteList, true
}

// HasTlsDefaultCipherSuiteList returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasTlsDefaultCipherSuiteList() bool {
	if o != nil && o.TlsDefaultCipherSuiteList != nil {
		return true
	}

	return false
}

// SetTlsDefaultCipherSuiteList gets a reference to the given bool and assigns it to the TlsDefaultCipherSuiteList field.
func (o *MsgVpnBridge) SetTlsDefaultCipherSuiteList(v bool) {
	o.TlsDefaultCipherSuiteList = &v
}

// GetTtlExceededEventRaised returns the TtlExceededEventRaised field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetTtlExceededEventRaised() bool {
	if o == nil || o.TtlExceededEventRaised == nil {
		var ret bool
		return ret
	}
	return *o.TtlExceededEventRaised
}

// GetTtlExceededEventRaisedOk returns a tuple with the TtlExceededEventRaised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetTtlExceededEventRaisedOk() (*bool, bool) {
	if o == nil || o.TtlExceededEventRaised == nil {
		return nil, false
	}
	return o.TtlExceededEventRaised, true
}

// HasTtlExceededEventRaised returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasTtlExceededEventRaised() bool {
	if o != nil && o.TtlExceededEventRaised != nil {
		return true
	}

	return false
}

// SetTtlExceededEventRaised gets a reference to the given bool and assigns it to the TtlExceededEventRaised field.
func (o *MsgVpnBridge) SetTtlExceededEventRaised(v bool) {
	o.TtlExceededEventRaised = &v
}

// GetTxByteCount returns the TxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetTxByteCount() int64 {
	if o == nil || o.TxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.TxByteCount
}

// GetTxByteCountOk returns a tuple with the TxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetTxByteCountOk() (*int64, bool) {
	if o == nil || o.TxByteCount == nil {
		return nil, false
	}
	return o.TxByteCount, true
}

// HasTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasTxByteCount() bool {
	if o != nil && o.TxByteCount != nil {
		return true
	}

	return false
}

// SetTxByteCount gets a reference to the given int64 and assigns it to the TxByteCount field.
func (o *MsgVpnBridge) SetTxByteCount(v int64) {
	o.TxByteCount = &v
}

// GetTxByteRate returns the TxByteRate field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetTxByteRate() int64 {
	if o == nil || o.TxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.TxByteRate
}

// GetTxByteRateOk returns a tuple with the TxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetTxByteRateOk() (*int64, bool) {
	if o == nil || o.TxByteRate == nil {
		return nil, false
	}
	return o.TxByteRate, true
}

// HasTxByteRate returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasTxByteRate() bool {
	if o != nil && o.TxByteRate != nil {
		return true
	}

	return false
}

// SetTxByteRate gets a reference to the given int64 and assigns it to the TxByteRate field.
func (o *MsgVpnBridge) SetTxByteRate(v int64) {
	o.TxByteRate = &v
}

// GetTxMsgCount returns the TxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetTxMsgCount() int64 {
	if o == nil || o.TxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.TxMsgCount
}

// GetTxMsgCountOk returns a tuple with the TxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetTxMsgCountOk() (*int64, bool) {
	if o == nil || o.TxMsgCount == nil {
		return nil, false
	}
	return o.TxMsgCount, true
}

// HasTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasTxMsgCount() bool {
	if o != nil && o.TxMsgCount != nil {
		return true
	}

	return false
}

// SetTxMsgCount gets a reference to the given int64 and assigns it to the TxMsgCount field.
func (o *MsgVpnBridge) SetTxMsgCount(v int64) {
	o.TxMsgCount = &v
}

// GetTxMsgRate returns the TxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetTxMsgRate() int64 {
	if o == nil || o.TxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.TxMsgRate
}

// GetTxMsgRateOk returns a tuple with the TxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetTxMsgRateOk() (*int64, bool) {
	if o == nil || o.TxMsgRate == nil {
		return nil, false
	}
	return o.TxMsgRate, true
}

// HasTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasTxMsgRate() bool {
	if o != nil && o.TxMsgRate != nil {
		return true
	}

	return false
}

// SetTxMsgRate gets a reference to the given int64 and assigns it to the TxMsgRate field.
func (o *MsgVpnBridge) SetTxMsgRate(v int64) {
	o.TxMsgRate = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetUptime() int64 {
	if o == nil || o.Uptime == nil {
		var ret int64
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetUptimeOk() (*int64, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given int64 and assigns it to the Uptime field.
func (o *MsgVpnBridge) SetUptime(v int64) {
	o.Uptime = &v
}

func (o MsgVpnBridge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AverageRxByteRate != nil {
		toSerialize["averageRxByteRate"] = o.AverageRxByteRate
	}
	if o.AverageRxMsgRate != nil {
		toSerialize["averageRxMsgRate"] = o.AverageRxMsgRate
	}
	if o.AverageTxByteRate != nil {
		toSerialize["averageTxByteRate"] = o.AverageTxByteRate
	}
	if o.AverageTxMsgRate != nil {
		toSerialize["averageTxMsgRate"] = o.AverageTxMsgRate
	}
	if o.BoundToQueue != nil {
		toSerialize["boundToQueue"] = o.BoundToQueue
	}
	if o.BridgeName != nil {
		toSerialize["bridgeName"] = o.BridgeName
	}
	if o.BridgeVirtualRouter != nil {
		toSerialize["bridgeVirtualRouter"] = o.BridgeVirtualRouter
	}
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.Compressed != nil {
		toSerialize["compressed"] = o.Compressed
	}
	if o.ControlRxByteCount != nil {
		toSerialize["controlRxByteCount"] = o.ControlRxByteCount
	}
	if o.ControlRxMsgCount != nil {
		toSerialize["controlRxMsgCount"] = o.ControlRxMsgCount
	}
	if o.ControlTxByteCount != nil {
		toSerialize["controlTxByteCount"] = o.ControlTxByteCount
	}
	if o.ControlTxMsgCount != nil {
		toSerialize["controlTxMsgCount"] = o.ControlTxMsgCount
	}
	if o.Counter != nil {
		toSerialize["counter"] = o.Counter
	}
	if o.DataRxByteCount != nil {
		toSerialize["dataRxByteCount"] = o.DataRxByteCount
	}
	if o.DataRxMsgCount != nil {
		toSerialize["dataRxMsgCount"] = o.DataRxMsgCount
	}
	if o.DataTxByteCount != nil {
		toSerialize["dataTxByteCount"] = o.DataTxByteCount
	}
	if o.DataTxMsgCount != nil {
		toSerialize["dataTxMsgCount"] = o.DataTxMsgCount
	}
	if o.DiscardedRxMsgCount != nil {
		toSerialize["discardedRxMsgCount"] = o.DiscardedRxMsgCount
	}
	if o.DiscardedTxMsgCount != nil {
		toSerialize["discardedTxMsgCount"] = o.DiscardedTxMsgCount
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Encrypted != nil {
		toSerialize["encrypted"] = o.Encrypted
	}
	if o.Establisher != nil {
		toSerialize["establisher"] = o.Establisher
	}
	if o.InboundFailureReason != nil {
		toSerialize["inboundFailureReason"] = o.InboundFailureReason
	}
	if o.InboundState != nil {
		toSerialize["inboundState"] = o.InboundState
	}
	if o.LastTxMsgId != nil {
		toSerialize["lastTxMsgId"] = o.LastTxMsgId
	}
	if o.LocalInterface != nil {
		toSerialize["localInterface"] = o.LocalInterface
	}
	if o.LocalQueueName != nil {
		toSerialize["localQueueName"] = o.LocalQueueName
	}
	if o.LoginRxMsgCount != nil {
		toSerialize["loginRxMsgCount"] = o.LoginRxMsgCount
	}
	if o.LoginTxMsgCount != nil {
		toSerialize["loginTxMsgCount"] = o.LoginTxMsgCount
	}
	if o.MaxTtl != nil {
		toSerialize["maxTtl"] = o.MaxTtl
	}
	if o.MsgSpoolRxMsgCount != nil {
		toSerialize["msgSpoolRxMsgCount"] = o.MsgSpoolRxMsgCount
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.OutboundState != nil {
		toSerialize["outboundState"] = o.OutboundState
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.RemoteAddress != nil {
		toSerialize["remoteAddress"] = o.RemoteAddress
	}
	if o.RemoteAuthenticationBasicClientUsername != nil {
		toSerialize["remoteAuthenticationBasicClientUsername"] = o.RemoteAuthenticationBasicClientUsername
	}
	if o.RemoteAuthenticationScheme != nil {
		toSerialize["remoteAuthenticationScheme"] = o.RemoteAuthenticationScheme
	}
	if o.RemoteConnectionRetryCount != nil {
		toSerialize["remoteConnectionRetryCount"] = o.RemoteConnectionRetryCount
	}
	if o.RemoteConnectionRetryDelay != nil {
		toSerialize["remoteConnectionRetryDelay"] = o.RemoteConnectionRetryDelay
	}
	if o.RemoteDeliverToOnePriority != nil {
		toSerialize["remoteDeliverToOnePriority"] = o.RemoteDeliverToOnePriority
	}
	if o.RemoteMsgVpnName != nil {
		toSerialize["remoteMsgVpnName"] = o.RemoteMsgVpnName
	}
	if o.RemoteRouterName != nil {
		toSerialize["remoteRouterName"] = o.RemoteRouterName
	}
	if o.RemoteTxFlowId != nil {
		toSerialize["remoteTxFlowId"] = o.RemoteTxFlowId
	}
	if o.RxByteCount != nil {
		toSerialize["rxByteCount"] = o.RxByteCount
	}
	if o.RxByteRate != nil {
		toSerialize["rxByteRate"] = o.RxByteRate
	}
	if o.RxConnectionFailureCategory != nil {
		toSerialize["rxConnectionFailureCategory"] = o.RxConnectionFailureCategory
	}
	if o.RxMsgCount != nil {
		toSerialize["rxMsgCount"] = o.RxMsgCount
	}
	if o.RxMsgRate != nil {
		toSerialize["rxMsgRate"] = o.RxMsgRate
	}
	if o.TlsCipherSuiteList != nil {
		toSerialize["tlsCipherSuiteList"] = o.TlsCipherSuiteList
	}
	if o.TlsDefaultCipherSuiteList != nil {
		toSerialize["tlsDefaultCipherSuiteList"] = o.TlsDefaultCipherSuiteList
	}
	if o.TtlExceededEventRaised != nil {
		toSerialize["ttlExceededEventRaised"] = o.TtlExceededEventRaised
	}
	if o.TxByteCount != nil {
		toSerialize["txByteCount"] = o.TxByteCount
	}
	if o.TxByteRate != nil {
		toSerialize["txByteRate"] = o.TxByteRate
	}
	if o.TxMsgCount != nil {
		toSerialize["txMsgCount"] = o.TxMsgCount
	}
	if o.TxMsgRate != nil {
		toSerialize["txMsgRate"] = o.TxMsgRate
	}
	if o.Uptime != nil {
		toSerialize["uptime"] = o.Uptime
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnBridge struct {
	value *MsgVpnBridge
	isSet bool
}

func (v NullableMsgVpnBridge) Get() *MsgVpnBridge {
	return v.value
}

func (v *NullableMsgVpnBridge) Set(val *MsgVpnBridge) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnBridge) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnBridge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnBridge(val *MsgVpnBridge) *NullableMsgVpnBridge {
	return &NullableMsgVpnBridge{value: val, isSet: true}
}

func (v NullableMsgVpnBridge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnBridge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
