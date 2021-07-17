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

// MsgVpnClient struct for MsgVpnClient
type MsgVpnClient struct {
	// The name of the access control list (ACL) profile of the Client.
	AclProfileName *string `json:"aclProfileName,omitempty"`
	// The name of the original MsgVpn which the client signaled in. Available since 2.14.
	AliasedFromMsgVpnName *string `json:"aliasedFromMsgVpnName,omitempty"`
	// The number of Client bind failures due to endpoint being already bound.
	AlreadyBoundBindFailureCount *int64 `json:"alreadyBoundBindFailureCount,omitempty"`
	// The name of the authorization group of the Client.
	AuthorizationGroupName *string `json:"authorizationGroupName,omitempty"`
	// The one minute average of the message rate received from the Client, in bytes per second (B/sec).
	AverageRxByteRate *int64 `json:"averageRxByteRate,omitempty"`
	// The one minute average of the message rate received from the Client, in messages per second (msg/sec).
	AverageRxMsgRate *int64 `json:"averageRxMsgRate,omitempty"`
	// The one minute average of the message rate transmitted to the Client, in bytes per second (B/sec).
	AverageTxByteRate *int64 `json:"averageTxByteRate,omitempty"`
	// The one minute average of the message rate transmitted to the Client, in messages per second (msg/sec).
	AverageTxMsgRate *int64 `json:"averageTxMsgRate,omitempty"`
	// The number of Client requests to bind to an endpoint.
	BindRequestCount *int64 `json:"bindRequestCount,omitempty"`
	// The number of successful Client requests to bind to an endpoint.
	BindSuccessCount *int64 `json:"bindSuccessCount,omitempty"`
	// The IP address and port of the Client.
	ClientAddress *string `json:"clientAddress,omitempty"`
	// The identifier (ID) of the Client.
	ClientId *int32 `json:"clientId,omitempty"`
	// The name of the Client.
	ClientName *string `json:"clientName,omitempty"`
	// The name of the client profile of the Client.
	ClientProfileName *string `json:"clientProfileName,omitempty"`
	// The client username of the Client used for authorization.
	ClientUsername *string `json:"clientUsername,omitempty"`
	// The amount of client control messages received from the Client, in bytes (B).
	ControlRxByteCount *int64 `json:"controlRxByteCount,omitempty"`
	// The number of client control messages received from the Client.
	ControlRxMsgCount *int64 `json:"controlRxMsgCount,omitempty"`
	// The amount of client control messages transmitted to the Client, in bytes (B).
	ControlTxByteCount *int64 `json:"controlTxByteCount,omitempty"`
	// The number of client control messages transmitted to the Client.
	ControlTxMsgCount *int64 `json:"controlTxMsgCount,omitempty"`
	// The number of Client bind failures due to being denied cut-through forwarding.
	CutThroughDeniedBindFailureCount *int64 `json:"cutThroughDeniedBindFailureCount,omitempty"`
	// The amount of client data messages received from the Client, in bytes (B).
	DataRxByteCount *int64 `json:"dataRxByteCount,omitempty"`
	// The number of client data messages received from the Client.
	DataRxMsgCount *int64 `json:"dataRxMsgCount,omitempty"`
	// The amount of client data messages transmitted to the Client, in bytes (B).
	DataTxByteCount *int64 `json:"dataTxByteCount,omitempty"`
	// The number of client data messages transmitted to the Client.
	DataTxMsgCount *int64 `json:"dataTxMsgCount,omitempty"`
	// The description text of the Client.
	Description *string `json:"description,omitempty"`
	// The number of Client bind failures due to endpoint being disabled.
	DisabledBindFailureCount *int64 `json:"disabledBindFailureCount,omitempty"`
	// The priority of the Client's subscriptions for receiving deliver-to-one (DTO) messages published on the local broker.
	DtoLocalPriority *int32 `json:"dtoLocalPriority,omitempty"`
	// The priority of the Client's subscriptions for receiving deliver-to-one (DTO) messages published on a remote broker.
	DtoNetworkPriority *int32 `json:"dtoNetworkPriority,omitempty"`
	// Indicates whether message eliding is enabled for the Client.
	Eliding *bool `json:"eliding,omitempty"`
	// The number of topics requiring message eliding for the Client.
	ElidingTopicCount *int32 `json:"elidingTopicCount,omitempty"`
	// The peak number of topics requiring message eliding for the Client.
	ElidingTopicPeakCount *int32 `json:"elidingTopicPeakCount,omitempty"`
	// The number of Client bind failures due to being denied guaranteed messaging.
	GuaranteedDeniedBindFailureCount *int64 `json:"guaranteedDeniedBindFailureCount,omitempty"`
	// The number of Client bind failures due to an invalid selector.
	InvalidSelectorBindFailureCount *int64 `json:"invalidSelectorBindFailureCount,omitempty"`
	// Indicates whether keepalive messages from the Client are received by the broker. Applicable for SMF and MQTT clients only. Available since 2.19.
	Keepalive *bool `json:"keepalive,omitempty"`
	// The maximum period of time the broker will accept inactivity from the Client before disconnecting, in seconds. Available since 2.19.
	KeepaliveEffectiveTimeout *int32 `json:"keepaliveEffectiveTimeout,omitempty"`
	// Indicates whether the large-message event has been raised for the Client.
	LargeMsgEventRaised *bool `json:"largeMsgEventRaised,omitempty"`
	// The number of login request messages received from the Client.
	LoginRxMsgCount *int64 `json:"loginRxMsgCount,omitempty"`
	// The number of login response messages transmitted to the Client.
	LoginTxMsgCount *int64 `json:"loginTxMsgCount,omitempty"`
	// The number of Client bind failures due to the endpoint maximum bind count being exceeded.
	MaxBindCountExceededBindFailureCount *int64 `json:"maxBindCountExceededBindFailureCount,omitempty"`
	// Indicates whether the max-eliding-topic-count event has been raised for the Client.
	MaxElidingTopicCountEventRaised *bool `json:"maxElidingTopicCountEventRaised,omitempty"`
	// The number of MQTT connect acknowledgment (CONNACK) refused response packets transmitted to the Client.
	MqttConnackErrorTxCount *int64 `json:"mqttConnackErrorTxCount,omitempty"`
	// The number of MQTT connect acknowledgment (CONNACK) accepted response packets transmitted to the Client.
	MqttConnackTxCount *int64 `json:"mqttConnackTxCount,omitempty"`
	// The number of MQTT connect (CONNECT) request packets received from the Client.
	MqttConnectRxCount *int64 `json:"mqttConnectRxCount,omitempty"`
	// The number of MQTT disconnect (DISCONNECT) request packets received from the Client.
	MqttDisconnectRxCount *int64 `json:"mqttDisconnectRxCount,omitempty"`
	// The number of MQTT ping request (PINGREQ) packets received from the Client.
	MqttPingreqRxCount *int64 `json:"mqttPingreqRxCount,omitempty"`
	// The number of MQTT ping response (PINGRESP) packets transmitted to the Client.
	MqttPingrespTxCount *int64 `json:"mqttPingrespTxCount,omitempty"`
	// The number of MQTT publish acknowledgement (PUBACK) response packets received from the Client.
	MqttPubackRxCount *int64 `json:"mqttPubackRxCount,omitempty"`
	// The number of MQTT publish acknowledgement (PUBACK) response packets transmitted to the Client.
	MqttPubackTxCount *int64 `json:"mqttPubackTxCount,omitempty"`
	// The number of MQTT publish complete (PUBCOMP) packets transmitted to the Client in response to a PUBREL packet. These packets are the fourth and final packet of a QoS 2 protocol exchange.
	MqttPubcompTxCount *int64 `json:"mqttPubcompTxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 0 message delivery.
	MqttPublishQos0RxCount *int64 `json:"mqttPublishQos0RxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 0 message delivery.
	MqttPublishQos0TxCount *int64 `json:"mqttPublishQos0TxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 1 message delivery.
	MqttPublishQos1RxCount *int64 `json:"mqttPublishQos1RxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 1 message delivery.
	MqttPublishQos1TxCount *int64 `json:"mqttPublishQos1TxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 2 message delivery.
	MqttPublishQos2RxCount *int64 `json:"mqttPublishQos2RxCount,omitempty"`
	// The number of MQTT publish received (PUBREC) packets transmitted to the Client in response to a PUBLISH packet with QoS 2. These packets are the second packet of a QoS 2 protocol exchange.
	MqttPubrecTxCount *int64 `json:"mqttPubrecTxCount,omitempty"`
	// The number of MQTT publish release (PUBREL) packets received from the Client in response to a PUBREC packet. These packets are the third packet of a QoS 2 protocol exchange.
	MqttPubrelRxCount *int64 `json:"mqttPubrelRxCount,omitempty"`
	// The number of MQTT subscribe acknowledgement (SUBACK) failure response packets transmitted to the Client.
	MqttSubackErrorTxCount *int64 `json:"mqttSubackErrorTxCount,omitempty"`
	// The number of MQTT subscribe acknowledgement (SUBACK) response packets transmitted to the Client.
	MqttSubackTxCount *int64 `json:"mqttSubackTxCount,omitempty"`
	// The number of MQTT subscribe (SUBSCRIBE) request packets received from the Client to create one or more topic subscriptions.
	MqttSubscribeRxCount *int64 `json:"mqttSubscribeRxCount,omitempty"`
	// The number of MQTT unsubscribe acknowledgement (UNSUBACK) response packets transmitted to the Client.
	MqttUnsubackTxCount *int64 `json:"mqttUnsubackTxCount,omitempty"`
	// The number of MQTT unsubscribe (UNSUBSCRIBE) request packets received from the Client to remove one or more topic subscriptions.
	MqttUnsubscribeRxCount *int64 `json:"mqttUnsubscribeRxCount,omitempty"`
	// The number of messages from the Client discarded due to message spool congestion primarily caused by message promotion.
	MsgSpoolCongestionRxDiscardedMsgCount *int64 `json:"msgSpoolCongestionRxDiscardedMsgCount,omitempty"`
	// The number of messages from the Client discarded by the message spool.
	MsgSpoolRxDiscardedMsgCount *int64 `json:"msgSpoolRxDiscardedMsgCount,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// Indicates whether not to deliver messages to the Client if it published them.
	NoLocalDelivery *bool `json:"noLocalDelivery,omitempty"`
	// The number of messages from the Client discarded due to no matching subscription found.
	NoSubscriptionMatchRxDiscardedMsgCount *int64 `json:"noSubscriptionMatchRxDiscardedMsgCount,omitempty"`
	// The original value of the client username used for Client authentication.
	OriginalClientUsername *string `json:"originalClientUsername,omitempty"`
	// The number of Client bind failures due to other reasons.
	OtherBindFailureCount *int64 `json:"otherBindFailureCount,omitempty"`
	// The platform the Client application software was built for, which may include the OS and API type.
	Platform *string `json:"platform,omitempty"`
	// The number of messages from the Client discarded due to the publish topic being denied by the Access Control List (ACL) profile.
	PublishTopicAclRxDiscardedMsgCount *int64 `json:"publishTopicAclRxDiscardedMsgCount,omitempty"`
	// The amount of HTTP request messages received from the Client, in bytes (B).
	RestHttpRequestRxByteCount *int64 `json:"restHttpRequestRxByteCount,omitempty"`
	// The number of HTTP request messages received from the Client.
	RestHttpRequestRxMsgCount *int64 `json:"restHttpRequestRxMsgCount,omitempty"`
	// The amount of HTTP request messages transmitted to the Client, in bytes (B).
	RestHttpRequestTxByteCount *int64 `json:"restHttpRequestTxByteCount,omitempty"`
	// The number of HTTP request messages transmitted to the Client.
	RestHttpRequestTxMsgCount *int64 `json:"restHttpRequestTxMsgCount,omitempty"`
	// The number of HTTP client/server error response messages received from the Client.
	RestHttpResponseErrorRxMsgCount *int64 `json:"restHttpResponseErrorRxMsgCount,omitempty"`
	// The number of HTTP client/server error response messages transmitted to the Client.
	RestHttpResponseErrorTxMsgCount *int64 `json:"restHttpResponseErrorTxMsgCount,omitempty"`
	// The amount of HTTP response messages received from the Client, in bytes (B).
	RestHttpResponseRxByteCount *int64 `json:"restHttpResponseRxByteCount,omitempty"`
	// The number of HTTP response messages received from the Client.
	RestHttpResponseRxMsgCount *int64 `json:"restHttpResponseRxMsgCount,omitempty"`
	// The number of HTTP successful response messages received from the Client.
	RestHttpResponseSuccessRxMsgCount *int64 `json:"restHttpResponseSuccessRxMsgCount,omitempty"`
	// The number of HTTP successful response messages transmitted to the Client.
	RestHttpResponseSuccessTxMsgCount *int64 `json:"restHttpResponseSuccessTxMsgCount,omitempty"`
	// The number of HTTP wait for reply timeout response messages received from the Client.
	RestHttpResponseTimeoutRxMsgCount *int64 `json:"restHttpResponseTimeoutRxMsgCount,omitempty"`
	// The number of HTTP wait for reply timeout response messages transmitted to the Client.
	RestHttpResponseTimeoutTxMsgCount *int64 `json:"restHttpResponseTimeoutTxMsgCount,omitempty"`
	// The amount of HTTP response messages transmitted to the Client, in bytes (B).
	RestHttpResponseTxByteCount *int64 `json:"restHttpResponseTxByteCount,omitempty"`
	// The number of HTTP response messages transmitted to the Client.
	RestHttpResponseTxMsgCount *int64 `json:"restHttpResponseTxMsgCount,omitempty"`
	// The amount of messages received from the Client, in bytes (B).
	RxByteCount *int64 `json:"rxByteCount,omitempty"`
	// The current message rate received from the Client, in bytes per second (B/sec).
	RxByteRate *int64 `json:"rxByteRate,omitempty"`
	// The number of messages discarded during reception from the Client.
	RxDiscardedMsgCount *int64 `json:"rxDiscardedMsgCount,omitempty"`
	// The number of messages received from the Client.
	RxMsgCount *int64 `json:"rxMsgCount,omitempty"`
	// The current message rate received from the Client, in messages per second (msg/sec).
	RxMsgRate *int64 `json:"rxMsgRate,omitempty"`
	// The timestamp of when the Client will be disconnected by the broker. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.13.
	ScheduledDisconnectTime *int32 `json:"scheduledDisconnectTime,omitempty"`
	// Indicates whether the Client is a slow subscriber and blocks for a few seconds when receiving messages.
	SlowSubscriber *bool `json:"slowSubscriber,omitempty"`
	// The date the Client application software was built.
	SoftwareDate *string `json:"softwareDate,omitempty"`
	// The version of the Client application software.
	SoftwareVersion *string `json:"softwareVersion,omitempty"`
	// The description of the TLS cipher used by the Client, which may include cipher name, key exchange and encryption algorithms.
	TlsCipherDescription *string `json:"tlsCipherDescription,omitempty"`
	// Indicates whether the Client TLS connection was downgraded to plain-text to increase performance.
	TlsDowngradedToPlainText *bool `json:"tlsDowngradedToPlainText,omitempty"`
	// The version of TLS used by the Client.
	TlsVersion *string `json:"tlsVersion,omitempty"`
	// The number of messages from the Client discarded due to an error while parsing the publish topic.
	TopicParseErrorRxDiscardedMsgCount *int64 `json:"topicParseErrorRxDiscardedMsgCount,omitempty"`
	// The amount of messages transmitted to the Client, in bytes (B).
	TxByteCount *int64 `json:"txByteCount,omitempty"`
	// The current message rate transmitted to the Client, in bytes per second (B/sec).
	TxByteRate *int64 `json:"txByteRate,omitempty"`
	// The number of messages discarded during transmission to the Client.
	TxDiscardedMsgCount *int64 `json:"txDiscardedMsgCount,omitempty"`
	// The number of messages transmitted to the Client.
	TxMsgCount *int64 `json:"txMsgCount,omitempty"`
	// The current message rate transmitted to the Client, in messages per second (msg/sec).
	TxMsgRate *int64 `json:"txMsgRate,omitempty"`
	// The amount of time in seconds since the Client connected.
	Uptime *int32 `json:"uptime,omitempty"`
	// The user description for the Client, which may include computer name and process ID.
	User *string `json:"user,omitempty"`
	// The virtual router used by the Client. The allowed values and their meaning are:  <pre> \"primary\" - The Client is using the primary virtual router. \"backup\" - The Client is using the backup virtual router. \"internal\" - The Client is using the internal virtual router. \"unknown\" - The Client virtual router is unknown. </pre>
	VirtualRouter *string `json:"virtualRouter,omitempty"`
	// The maximum web transport timeout for the Client being inactive, in seconds.
	WebInactiveTimeout *int32 `json:"webInactiveTimeout,omitempty"`
	// The maximum web transport message payload size which excludes the size of the message header, in bytes.
	WebMaxPayload *int64 `json:"webMaxPayload,omitempty"`
	// The number of messages from the Client discarded due to an error while parsing the web message.
	WebParseErrorRxDiscardedMsgCount *int64 `json:"webParseErrorRxDiscardedMsgCount,omitempty"`
	// The remaining web transport timeout for the Client being inactive, in seconds.
	WebRemainingTimeout *int32 `json:"webRemainingTimeout,omitempty"`
	// The amount of web transport messages received from the Client, in bytes (B).
	WebRxByteCount *int64 `json:"webRxByteCount,omitempty"`
	// The type of encoding used during reception from the Client. The allowed values and their meaning are:  <pre> \"binary\" - The Client is using binary encoding. \"base64\" - The Client is using base64 encoding. \"illegal\" - The Client is using an illegal encoding type. </pre>
	WebRxEncoding *string `json:"webRxEncoding,omitempty"`
	// The number of web transport messages received from the Client.
	WebRxMsgCount *int64 `json:"webRxMsgCount,omitempty"`
	// The type of web transport used during reception from the Client. The allowed values and their meaning are:  <pre> \"ws-binary\" - The Client is using WebSocket binary transport. \"http-binary-streaming\" - The Client is using HTTP binary streaming transport. \"http-binary\" - The Client is using HTTP binary transport. \"http-base64\" - The Client is using HTTP base64 transport. </pre>
	WebRxProtocol *string `json:"webRxProtocol,omitempty"`
	// The number of web transport requests received from the Client (HTTP only). Not available for WebSockets.
	WebRxRequestCount *int64 `json:"webRxRequestCount,omitempty"`
	// The number of web transport responses transmitted to the Client on the receive connection (HTTP only). Not available for WebSockets.
	WebRxResponseCount *int64 `json:"webRxResponseCount,omitempty"`
	// The TCP state of the receive connection from the Client. When fully operational, should be: established. See RFC 793 for further details. The allowed values and their meaning are:  <pre> \"closed\" - No connection state at all. \"listen\" - Waiting for a connection request from any remote TCP and port. \"syn-sent\" - Waiting for a matching connection request after having sent a connection request. \"syn-received\" - Waiting for a confirming connection request acknowledgment after having both received and sent a connection request. \"established\" - An open connection, data received can be delivered to the user. \"close-wait\" - Waiting for a connection termination request from the local user. \"fin-wait-1\" - Waiting for a connection termination request from the remote TCP, or an acknowledgment of the connection termination request previously sent. \"closing\" - Waiting for a connection termination request acknowledgment from the remote TCP. \"last-ack\" - Waiting for an acknowledgment of the connection termination request previously sent to the remote TCP. \"fin-wait-2\" - Waiting for a connection termination request from the remote TCP. \"time-wait\" - Waiting for enough time to pass to be sure the remote TCP received the acknowledgment of its connection termination request. </pre>
	WebRxTcpState *string `json:"webRxTcpState,omitempty"`
	// The description of the TLS cipher received from the Client, which may include cipher name, key exchange and encryption algorithms.
	WebRxTlsCipherDescription *string `json:"webRxTlsCipherDescription,omitempty"`
	// The version of TLS used during reception from the Client.
	WebRxTlsVersion *string `json:"webRxTlsVersion,omitempty"`
	// The identifier (ID) of the web transport session for the Client.
	WebSessionId *string `json:"webSessionId,omitempty"`
	// The amount of web transport messages transmitted to the Client, in bytes (B).
	WebTxByteCount *int64 `json:"webTxByteCount,omitempty"`
	// The type of encoding used during transmission to the Client. The allowed values and their meaning are:  <pre> \"binary\" - The Client is using binary encoding. \"base64\" - The Client is using base64 encoding. \"illegal\" - The Client is using an illegal encoding type. </pre>
	WebTxEncoding *string `json:"webTxEncoding,omitempty"`
	// The number of web transport messages transmitted to the Client.
	WebTxMsgCount *int64 `json:"webTxMsgCount,omitempty"`
	// The type of web transport used during transmission to the Client. The allowed values and their meaning are:  <pre> \"ws-binary\" - The Client is using WebSocket binary transport. \"http-binary-streaming\" - The Client is using HTTP binary streaming transport. \"http-binary\" - The Client is using HTTP binary transport. \"http-base64\" - The Client is using HTTP base64 transport. </pre>
	WebTxProtocol *string `json:"webTxProtocol,omitempty"`
	// The number of web transport requests transmitted to the Client (HTTP only). Not available for WebSockets.
	WebTxRequestCount *int64 `json:"webTxRequestCount,omitempty"`
	// The number of web transport responses received from the Client on the transmit connection (HTTP only). Not available for WebSockets.
	WebTxResponseCount *int64 `json:"webTxResponseCount,omitempty"`
	// The TCP state of the transmit connection to the Client. When fully operational, should be: established. See RFC 793 for further details. The allowed values and their meaning are:  <pre> \"closed\" - No connection state at all. \"listen\" - Waiting for a connection request from any remote TCP and port. \"syn-sent\" - Waiting for a matching connection request after having sent a connection request. \"syn-received\" - Waiting for a confirming connection request acknowledgment after having both received and sent a connection request. \"established\" - An open connection, data received can be delivered to the user. \"close-wait\" - Waiting for a connection termination request from the local user. \"fin-wait-1\" - Waiting for a connection termination request from the remote TCP, or an acknowledgment of the connection termination request previously sent. \"closing\" - Waiting for a connection termination request acknowledgment from the remote TCP. \"last-ack\" - Waiting for an acknowledgment of the connection termination request previously sent to the remote TCP. \"fin-wait-2\" - Waiting for a connection termination request from the remote TCP. \"time-wait\" - Waiting for enough time to pass to be sure the remote TCP received the acknowledgment of its connection termination request. </pre>
	WebTxTcpState *string `json:"webTxTcpState,omitempty"`
	// The description of the TLS cipher transmitted to the Client, which may include cipher name, key exchange and encryption algorithms.
	WebTxTlsCipherDescription *string `json:"webTxTlsCipherDescription,omitempty"`
	// The version of TLS used during transmission to the Client.
	WebTxTlsVersion *string `json:"webTxTlsVersion,omitempty"`
}

// NewMsgVpnClient instantiates a new MsgVpnClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnClient() *MsgVpnClient {
	this := MsgVpnClient{}
	return &this
}

// NewMsgVpnClientWithDefaults instantiates a new MsgVpnClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnClientWithDefaults() *MsgVpnClient {
	this := MsgVpnClient{}
	return &this
}

// GetAclProfileName returns the AclProfileName field value if set, zero value otherwise.
func (o *MsgVpnClient) GetAclProfileName() string {
	if o == nil || o.AclProfileName == nil {
		var ret string
		return ret
	}
	return *o.AclProfileName
}

// GetAclProfileNameOk returns a tuple with the AclProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetAclProfileNameOk() (*string, bool) {
	if o == nil || o.AclProfileName == nil {
		return nil, false
	}
	return o.AclProfileName, true
}

// HasAclProfileName returns a boolean if a field has been set.
func (o *MsgVpnClient) HasAclProfileName() bool {
	if o != nil && o.AclProfileName != nil {
		return true
	}

	return false
}

// SetAclProfileName gets a reference to the given string and assigns it to the AclProfileName field.
func (o *MsgVpnClient) SetAclProfileName(v string) {
	o.AclProfileName = &v
}

// GetAliasedFromMsgVpnName returns the AliasedFromMsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnClient) GetAliasedFromMsgVpnName() string {
	if o == nil || o.AliasedFromMsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.AliasedFromMsgVpnName
}

// GetAliasedFromMsgVpnNameOk returns a tuple with the AliasedFromMsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetAliasedFromMsgVpnNameOk() (*string, bool) {
	if o == nil || o.AliasedFromMsgVpnName == nil {
		return nil, false
	}
	return o.AliasedFromMsgVpnName, true
}

// HasAliasedFromMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnClient) HasAliasedFromMsgVpnName() bool {
	if o != nil && o.AliasedFromMsgVpnName != nil {
		return true
	}

	return false
}

// SetAliasedFromMsgVpnName gets a reference to the given string and assigns it to the AliasedFromMsgVpnName field.
func (o *MsgVpnClient) SetAliasedFromMsgVpnName(v string) {
	o.AliasedFromMsgVpnName = &v
}

// GetAlreadyBoundBindFailureCount returns the AlreadyBoundBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetAlreadyBoundBindFailureCount() int64 {
	if o == nil || o.AlreadyBoundBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.AlreadyBoundBindFailureCount
}

// GetAlreadyBoundBindFailureCountOk returns a tuple with the AlreadyBoundBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetAlreadyBoundBindFailureCountOk() (*int64, bool) {
	if o == nil || o.AlreadyBoundBindFailureCount == nil {
		return nil, false
	}
	return o.AlreadyBoundBindFailureCount, true
}

// HasAlreadyBoundBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasAlreadyBoundBindFailureCount() bool {
	if o != nil && o.AlreadyBoundBindFailureCount != nil {
		return true
	}

	return false
}

// SetAlreadyBoundBindFailureCount gets a reference to the given int64 and assigns it to the AlreadyBoundBindFailureCount field.
func (o *MsgVpnClient) SetAlreadyBoundBindFailureCount(v int64) {
	o.AlreadyBoundBindFailureCount = &v
}

// GetAuthorizationGroupName returns the AuthorizationGroupName field value if set, zero value otherwise.
func (o *MsgVpnClient) GetAuthorizationGroupName() string {
	if o == nil || o.AuthorizationGroupName == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationGroupName
}

// GetAuthorizationGroupNameOk returns a tuple with the AuthorizationGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetAuthorizationGroupNameOk() (*string, bool) {
	if o == nil || o.AuthorizationGroupName == nil {
		return nil, false
	}
	return o.AuthorizationGroupName, true
}

// HasAuthorizationGroupName returns a boolean if a field has been set.
func (o *MsgVpnClient) HasAuthorizationGroupName() bool {
	if o != nil && o.AuthorizationGroupName != nil {
		return true
	}

	return false
}

// SetAuthorizationGroupName gets a reference to the given string and assigns it to the AuthorizationGroupName field.
func (o *MsgVpnClient) SetAuthorizationGroupName(v string) {
	o.AuthorizationGroupName = &v
}

// GetAverageRxByteRate returns the AverageRxByteRate field value if set, zero value otherwise.
func (o *MsgVpnClient) GetAverageRxByteRate() int64 {
	if o == nil || o.AverageRxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRxByteRate
}

// GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetAverageRxByteRateOk() (*int64, bool) {
	if o == nil || o.AverageRxByteRate == nil {
		return nil, false
	}
	return o.AverageRxByteRate, true
}

// HasAverageRxByteRate returns a boolean if a field has been set.
func (o *MsgVpnClient) HasAverageRxByteRate() bool {
	if o != nil && o.AverageRxByteRate != nil {
		return true
	}

	return false
}

// SetAverageRxByteRate gets a reference to the given int64 and assigns it to the AverageRxByteRate field.
func (o *MsgVpnClient) SetAverageRxByteRate(v int64) {
	o.AverageRxByteRate = &v
}

// GetAverageRxMsgRate returns the AverageRxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnClient) GetAverageRxMsgRate() int64 {
	if o == nil || o.AverageRxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRxMsgRate
}

// GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetAverageRxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageRxMsgRate == nil {
		return nil, false
	}
	return o.AverageRxMsgRate, true
}

// HasAverageRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnClient) HasAverageRxMsgRate() bool {
	if o != nil && o.AverageRxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageRxMsgRate gets a reference to the given int64 and assigns it to the AverageRxMsgRate field.
func (o *MsgVpnClient) SetAverageRxMsgRate(v int64) {
	o.AverageRxMsgRate = &v
}

// GetAverageTxByteRate returns the AverageTxByteRate field value if set, zero value otherwise.
func (o *MsgVpnClient) GetAverageTxByteRate() int64 {
	if o == nil || o.AverageTxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageTxByteRate
}

// GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetAverageTxByteRateOk() (*int64, bool) {
	if o == nil || o.AverageTxByteRate == nil {
		return nil, false
	}
	return o.AverageTxByteRate, true
}

// HasAverageTxByteRate returns a boolean if a field has been set.
func (o *MsgVpnClient) HasAverageTxByteRate() bool {
	if o != nil && o.AverageTxByteRate != nil {
		return true
	}

	return false
}

// SetAverageTxByteRate gets a reference to the given int64 and assigns it to the AverageTxByteRate field.
func (o *MsgVpnClient) SetAverageTxByteRate(v int64) {
	o.AverageTxByteRate = &v
}

// GetAverageTxMsgRate returns the AverageTxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnClient) GetAverageTxMsgRate() int64 {
	if o == nil || o.AverageTxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageTxMsgRate
}

// GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetAverageTxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageTxMsgRate == nil {
		return nil, false
	}
	return o.AverageTxMsgRate, true
}

// HasAverageTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnClient) HasAverageTxMsgRate() bool {
	if o != nil && o.AverageTxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageTxMsgRate gets a reference to the given int64 and assigns it to the AverageTxMsgRate field.
func (o *MsgVpnClient) SetAverageTxMsgRate(v int64) {
	o.AverageTxMsgRate = &v
}

// GetBindRequestCount returns the BindRequestCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetBindRequestCount() int64 {
	if o == nil || o.BindRequestCount == nil {
		var ret int64
		return ret
	}
	return *o.BindRequestCount
}

// GetBindRequestCountOk returns a tuple with the BindRequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetBindRequestCountOk() (*int64, bool) {
	if o == nil || o.BindRequestCount == nil {
		return nil, false
	}
	return o.BindRequestCount, true
}

// HasBindRequestCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasBindRequestCount() bool {
	if o != nil && o.BindRequestCount != nil {
		return true
	}

	return false
}

// SetBindRequestCount gets a reference to the given int64 and assigns it to the BindRequestCount field.
func (o *MsgVpnClient) SetBindRequestCount(v int64) {
	o.BindRequestCount = &v
}

// GetBindSuccessCount returns the BindSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetBindSuccessCount() int64 {
	if o == nil || o.BindSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.BindSuccessCount
}

// GetBindSuccessCountOk returns a tuple with the BindSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetBindSuccessCountOk() (*int64, bool) {
	if o == nil || o.BindSuccessCount == nil {
		return nil, false
	}
	return o.BindSuccessCount, true
}

// HasBindSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasBindSuccessCount() bool {
	if o != nil && o.BindSuccessCount != nil {
		return true
	}

	return false
}

// SetBindSuccessCount gets a reference to the given int64 and assigns it to the BindSuccessCount field.
func (o *MsgVpnClient) SetBindSuccessCount(v int64) {
	o.BindSuccessCount = &v
}

// GetClientAddress returns the ClientAddress field value if set, zero value otherwise.
func (o *MsgVpnClient) GetClientAddress() string {
	if o == nil || o.ClientAddress == nil {
		var ret string
		return ret
	}
	return *o.ClientAddress
}

// GetClientAddressOk returns a tuple with the ClientAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetClientAddressOk() (*string, bool) {
	if o == nil || o.ClientAddress == nil {
		return nil, false
	}
	return o.ClientAddress, true
}

// HasClientAddress returns a boolean if a field has been set.
func (o *MsgVpnClient) HasClientAddress() bool {
	if o != nil && o.ClientAddress != nil {
		return true
	}

	return false
}

// SetClientAddress gets a reference to the given string and assigns it to the ClientAddress field.
func (o *MsgVpnClient) SetClientAddress(v string) {
	o.ClientAddress = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *MsgVpnClient) GetClientId() int32 {
	if o == nil || o.ClientId == nil {
		var ret int32
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetClientIdOk() (*int32, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *MsgVpnClient) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int32 and assigns it to the ClientId field.
func (o *MsgVpnClient) SetClientId(v int32) {
	o.ClientId = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *MsgVpnClient) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *MsgVpnClient) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *MsgVpnClient) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientProfileName returns the ClientProfileName field value if set, zero value otherwise.
func (o *MsgVpnClient) GetClientProfileName() string {
	if o == nil || o.ClientProfileName == nil {
		var ret string
		return ret
	}
	return *o.ClientProfileName
}

// GetClientProfileNameOk returns a tuple with the ClientProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetClientProfileNameOk() (*string, bool) {
	if o == nil || o.ClientProfileName == nil {
		return nil, false
	}
	return o.ClientProfileName, true
}

// HasClientProfileName returns a boolean if a field has been set.
func (o *MsgVpnClient) HasClientProfileName() bool {
	if o != nil && o.ClientProfileName != nil {
		return true
	}

	return false
}

// SetClientProfileName gets a reference to the given string and assigns it to the ClientProfileName field.
func (o *MsgVpnClient) SetClientProfileName(v string) {
	o.ClientProfileName = &v
}

// GetClientUsername returns the ClientUsername field value if set, zero value otherwise.
func (o *MsgVpnClient) GetClientUsername() string {
	if o == nil || o.ClientUsername == nil {
		var ret string
		return ret
	}
	return *o.ClientUsername
}

// GetClientUsernameOk returns a tuple with the ClientUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetClientUsernameOk() (*string, bool) {
	if o == nil || o.ClientUsername == nil {
		return nil, false
	}
	return o.ClientUsername, true
}

// HasClientUsername returns a boolean if a field has been set.
func (o *MsgVpnClient) HasClientUsername() bool {
	if o != nil && o.ClientUsername != nil {
		return true
	}

	return false
}

// SetClientUsername gets a reference to the given string and assigns it to the ClientUsername field.
func (o *MsgVpnClient) SetClientUsername(v string) {
	o.ClientUsername = &v
}

// GetControlRxByteCount returns the ControlRxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetControlRxByteCount() int64 {
	if o == nil || o.ControlRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlRxByteCount
}

// GetControlRxByteCountOk returns a tuple with the ControlRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetControlRxByteCountOk() (*int64, bool) {
	if o == nil || o.ControlRxByteCount == nil {
		return nil, false
	}
	return o.ControlRxByteCount, true
}

// HasControlRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasControlRxByteCount() bool {
	if o != nil && o.ControlRxByteCount != nil {
		return true
	}

	return false
}

// SetControlRxByteCount gets a reference to the given int64 and assigns it to the ControlRxByteCount field.
func (o *MsgVpnClient) SetControlRxByteCount(v int64) {
	o.ControlRxByteCount = &v
}

// GetControlRxMsgCount returns the ControlRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetControlRxMsgCount() int64 {
	if o == nil || o.ControlRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlRxMsgCount
}

// GetControlRxMsgCountOk returns a tuple with the ControlRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetControlRxMsgCountOk() (*int64, bool) {
	if o == nil || o.ControlRxMsgCount == nil {
		return nil, false
	}
	return o.ControlRxMsgCount, true
}

// HasControlRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasControlRxMsgCount() bool {
	if o != nil && o.ControlRxMsgCount != nil {
		return true
	}

	return false
}

// SetControlRxMsgCount gets a reference to the given int64 and assigns it to the ControlRxMsgCount field.
func (o *MsgVpnClient) SetControlRxMsgCount(v int64) {
	o.ControlRxMsgCount = &v
}

// GetControlTxByteCount returns the ControlTxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetControlTxByteCount() int64 {
	if o == nil || o.ControlTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlTxByteCount
}

// GetControlTxByteCountOk returns a tuple with the ControlTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetControlTxByteCountOk() (*int64, bool) {
	if o == nil || o.ControlTxByteCount == nil {
		return nil, false
	}
	return o.ControlTxByteCount, true
}

// HasControlTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasControlTxByteCount() bool {
	if o != nil && o.ControlTxByteCount != nil {
		return true
	}

	return false
}

// SetControlTxByteCount gets a reference to the given int64 and assigns it to the ControlTxByteCount field.
func (o *MsgVpnClient) SetControlTxByteCount(v int64) {
	o.ControlTxByteCount = &v
}

// GetControlTxMsgCount returns the ControlTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetControlTxMsgCount() int64 {
	if o == nil || o.ControlTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlTxMsgCount
}

// GetControlTxMsgCountOk returns a tuple with the ControlTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetControlTxMsgCountOk() (*int64, bool) {
	if o == nil || o.ControlTxMsgCount == nil {
		return nil, false
	}
	return o.ControlTxMsgCount, true
}

// HasControlTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasControlTxMsgCount() bool {
	if o != nil && o.ControlTxMsgCount != nil {
		return true
	}

	return false
}

// SetControlTxMsgCount gets a reference to the given int64 and assigns it to the ControlTxMsgCount field.
func (o *MsgVpnClient) SetControlTxMsgCount(v int64) {
	o.ControlTxMsgCount = &v
}

// GetCutThroughDeniedBindFailureCount returns the CutThroughDeniedBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetCutThroughDeniedBindFailureCount() int64 {
	if o == nil || o.CutThroughDeniedBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.CutThroughDeniedBindFailureCount
}

// GetCutThroughDeniedBindFailureCountOk returns a tuple with the CutThroughDeniedBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetCutThroughDeniedBindFailureCountOk() (*int64, bool) {
	if o == nil || o.CutThroughDeniedBindFailureCount == nil {
		return nil, false
	}
	return o.CutThroughDeniedBindFailureCount, true
}

// HasCutThroughDeniedBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasCutThroughDeniedBindFailureCount() bool {
	if o != nil && o.CutThroughDeniedBindFailureCount != nil {
		return true
	}

	return false
}

// SetCutThroughDeniedBindFailureCount gets a reference to the given int64 and assigns it to the CutThroughDeniedBindFailureCount field.
func (o *MsgVpnClient) SetCutThroughDeniedBindFailureCount(v int64) {
	o.CutThroughDeniedBindFailureCount = &v
}

// GetDataRxByteCount returns the DataRxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetDataRxByteCount() int64 {
	if o == nil || o.DataRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.DataRxByteCount
}

// GetDataRxByteCountOk returns a tuple with the DataRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetDataRxByteCountOk() (*int64, bool) {
	if o == nil || o.DataRxByteCount == nil {
		return nil, false
	}
	return o.DataRxByteCount, true
}

// HasDataRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasDataRxByteCount() bool {
	if o != nil && o.DataRxByteCount != nil {
		return true
	}

	return false
}

// SetDataRxByteCount gets a reference to the given int64 and assigns it to the DataRxByteCount field.
func (o *MsgVpnClient) SetDataRxByteCount(v int64) {
	o.DataRxByteCount = &v
}

// GetDataRxMsgCount returns the DataRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetDataRxMsgCount() int64 {
	if o == nil || o.DataRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DataRxMsgCount
}

// GetDataRxMsgCountOk returns a tuple with the DataRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetDataRxMsgCountOk() (*int64, bool) {
	if o == nil || o.DataRxMsgCount == nil {
		return nil, false
	}
	return o.DataRxMsgCount, true
}

// HasDataRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasDataRxMsgCount() bool {
	if o != nil && o.DataRxMsgCount != nil {
		return true
	}

	return false
}

// SetDataRxMsgCount gets a reference to the given int64 and assigns it to the DataRxMsgCount field.
func (o *MsgVpnClient) SetDataRxMsgCount(v int64) {
	o.DataRxMsgCount = &v
}

// GetDataTxByteCount returns the DataTxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetDataTxByteCount() int64 {
	if o == nil || o.DataTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.DataTxByteCount
}

// GetDataTxByteCountOk returns a tuple with the DataTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetDataTxByteCountOk() (*int64, bool) {
	if o == nil || o.DataTxByteCount == nil {
		return nil, false
	}
	return o.DataTxByteCount, true
}

// HasDataTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasDataTxByteCount() bool {
	if o != nil && o.DataTxByteCount != nil {
		return true
	}

	return false
}

// SetDataTxByteCount gets a reference to the given int64 and assigns it to the DataTxByteCount field.
func (o *MsgVpnClient) SetDataTxByteCount(v int64) {
	o.DataTxByteCount = &v
}

// GetDataTxMsgCount returns the DataTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetDataTxMsgCount() int64 {
	if o == nil || o.DataTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DataTxMsgCount
}

// GetDataTxMsgCountOk returns a tuple with the DataTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetDataTxMsgCountOk() (*int64, bool) {
	if o == nil || o.DataTxMsgCount == nil {
		return nil, false
	}
	return o.DataTxMsgCount, true
}

// HasDataTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasDataTxMsgCount() bool {
	if o != nil && o.DataTxMsgCount != nil {
		return true
	}

	return false
}

// SetDataTxMsgCount gets a reference to the given int64 and assigns it to the DataTxMsgCount field.
func (o *MsgVpnClient) SetDataTxMsgCount(v int64) {
	o.DataTxMsgCount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MsgVpnClient) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MsgVpnClient) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MsgVpnClient) SetDescription(v string) {
	o.Description = &v
}

// GetDisabledBindFailureCount returns the DisabledBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetDisabledBindFailureCount() int64 {
	if o == nil || o.DisabledBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.DisabledBindFailureCount
}

// GetDisabledBindFailureCountOk returns a tuple with the DisabledBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetDisabledBindFailureCountOk() (*int64, bool) {
	if o == nil || o.DisabledBindFailureCount == nil {
		return nil, false
	}
	return o.DisabledBindFailureCount, true
}

// HasDisabledBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasDisabledBindFailureCount() bool {
	if o != nil && o.DisabledBindFailureCount != nil {
		return true
	}

	return false
}

// SetDisabledBindFailureCount gets a reference to the given int64 and assigns it to the DisabledBindFailureCount field.
func (o *MsgVpnClient) SetDisabledBindFailureCount(v int64) {
	o.DisabledBindFailureCount = &v
}

// GetDtoLocalPriority returns the DtoLocalPriority field value if set, zero value otherwise.
func (o *MsgVpnClient) GetDtoLocalPriority() int32 {
	if o == nil || o.DtoLocalPriority == nil {
		var ret int32
		return ret
	}
	return *o.DtoLocalPriority
}

// GetDtoLocalPriorityOk returns a tuple with the DtoLocalPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetDtoLocalPriorityOk() (*int32, bool) {
	if o == nil || o.DtoLocalPriority == nil {
		return nil, false
	}
	return o.DtoLocalPriority, true
}

// HasDtoLocalPriority returns a boolean if a field has been set.
func (o *MsgVpnClient) HasDtoLocalPriority() bool {
	if o != nil && o.DtoLocalPriority != nil {
		return true
	}

	return false
}

// SetDtoLocalPriority gets a reference to the given int32 and assigns it to the DtoLocalPriority field.
func (o *MsgVpnClient) SetDtoLocalPriority(v int32) {
	o.DtoLocalPriority = &v
}

// GetDtoNetworkPriority returns the DtoNetworkPriority field value if set, zero value otherwise.
func (o *MsgVpnClient) GetDtoNetworkPriority() int32 {
	if o == nil || o.DtoNetworkPriority == nil {
		var ret int32
		return ret
	}
	return *o.DtoNetworkPriority
}

// GetDtoNetworkPriorityOk returns a tuple with the DtoNetworkPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetDtoNetworkPriorityOk() (*int32, bool) {
	if o == nil || o.DtoNetworkPriority == nil {
		return nil, false
	}
	return o.DtoNetworkPriority, true
}

// HasDtoNetworkPriority returns a boolean if a field has been set.
func (o *MsgVpnClient) HasDtoNetworkPriority() bool {
	if o != nil && o.DtoNetworkPriority != nil {
		return true
	}

	return false
}

// SetDtoNetworkPriority gets a reference to the given int32 and assigns it to the DtoNetworkPriority field.
func (o *MsgVpnClient) SetDtoNetworkPriority(v int32) {
	o.DtoNetworkPriority = &v
}

// GetEliding returns the Eliding field value if set, zero value otherwise.
func (o *MsgVpnClient) GetEliding() bool {
	if o == nil || o.Eliding == nil {
		var ret bool
		return ret
	}
	return *o.Eliding
}

// GetElidingOk returns a tuple with the Eliding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetElidingOk() (*bool, bool) {
	if o == nil || o.Eliding == nil {
		return nil, false
	}
	return o.Eliding, true
}

// HasEliding returns a boolean if a field has been set.
func (o *MsgVpnClient) HasEliding() bool {
	if o != nil && o.Eliding != nil {
		return true
	}

	return false
}

// SetEliding gets a reference to the given bool and assigns it to the Eliding field.
func (o *MsgVpnClient) SetEliding(v bool) {
	o.Eliding = &v
}

// GetElidingTopicCount returns the ElidingTopicCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetElidingTopicCount() int32 {
	if o == nil || o.ElidingTopicCount == nil {
		var ret int32
		return ret
	}
	return *o.ElidingTopicCount
}

// GetElidingTopicCountOk returns a tuple with the ElidingTopicCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetElidingTopicCountOk() (*int32, bool) {
	if o == nil || o.ElidingTopicCount == nil {
		return nil, false
	}
	return o.ElidingTopicCount, true
}

// HasElidingTopicCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasElidingTopicCount() bool {
	if o != nil && o.ElidingTopicCount != nil {
		return true
	}

	return false
}

// SetElidingTopicCount gets a reference to the given int32 and assigns it to the ElidingTopicCount field.
func (o *MsgVpnClient) SetElidingTopicCount(v int32) {
	o.ElidingTopicCount = &v
}

// GetElidingTopicPeakCount returns the ElidingTopicPeakCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetElidingTopicPeakCount() int32 {
	if o == nil || o.ElidingTopicPeakCount == nil {
		var ret int32
		return ret
	}
	return *o.ElidingTopicPeakCount
}

// GetElidingTopicPeakCountOk returns a tuple with the ElidingTopicPeakCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetElidingTopicPeakCountOk() (*int32, bool) {
	if o == nil || o.ElidingTopicPeakCount == nil {
		return nil, false
	}
	return o.ElidingTopicPeakCount, true
}

// HasElidingTopicPeakCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasElidingTopicPeakCount() bool {
	if o != nil && o.ElidingTopicPeakCount != nil {
		return true
	}

	return false
}

// SetElidingTopicPeakCount gets a reference to the given int32 and assigns it to the ElidingTopicPeakCount field.
func (o *MsgVpnClient) SetElidingTopicPeakCount(v int32) {
	o.ElidingTopicPeakCount = &v
}

// GetGuaranteedDeniedBindFailureCount returns the GuaranteedDeniedBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetGuaranteedDeniedBindFailureCount() int64 {
	if o == nil || o.GuaranteedDeniedBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.GuaranteedDeniedBindFailureCount
}

// GetGuaranteedDeniedBindFailureCountOk returns a tuple with the GuaranteedDeniedBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetGuaranteedDeniedBindFailureCountOk() (*int64, bool) {
	if o == nil || o.GuaranteedDeniedBindFailureCount == nil {
		return nil, false
	}
	return o.GuaranteedDeniedBindFailureCount, true
}

// HasGuaranteedDeniedBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasGuaranteedDeniedBindFailureCount() bool {
	if o != nil && o.GuaranteedDeniedBindFailureCount != nil {
		return true
	}

	return false
}

// SetGuaranteedDeniedBindFailureCount gets a reference to the given int64 and assigns it to the GuaranteedDeniedBindFailureCount field.
func (o *MsgVpnClient) SetGuaranteedDeniedBindFailureCount(v int64) {
	o.GuaranteedDeniedBindFailureCount = &v
}

// GetInvalidSelectorBindFailureCount returns the InvalidSelectorBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetInvalidSelectorBindFailureCount() int64 {
	if o == nil || o.InvalidSelectorBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.InvalidSelectorBindFailureCount
}

// GetInvalidSelectorBindFailureCountOk returns a tuple with the InvalidSelectorBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetInvalidSelectorBindFailureCountOk() (*int64, bool) {
	if o == nil || o.InvalidSelectorBindFailureCount == nil {
		return nil, false
	}
	return o.InvalidSelectorBindFailureCount, true
}

// HasInvalidSelectorBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasInvalidSelectorBindFailureCount() bool {
	if o != nil && o.InvalidSelectorBindFailureCount != nil {
		return true
	}

	return false
}

// SetInvalidSelectorBindFailureCount gets a reference to the given int64 and assigns it to the InvalidSelectorBindFailureCount field.
func (o *MsgVpnClient) SetInvalidSelectorBindFailureCount(v int64) {
	o.InvalidSelectorBindFailureCount = &v
}

// GetKeepalive returns the Keepalive field value if set, zero value otherwise.
func (o *MsgVpnClient) GetKeepalive() bool {
	if o == nil || o.Keepalive == nil {
		var ret bool
		return ret
	}
	return *o.Keepalive
}

// GetKeepaliveOk returns a tuple with the Keepalive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetKeepaliveOk() (*bool, bool) {
	if o == nil || o.Keepalive == nil {
		return nil, false
	}
	return o.Keepalive, true
}

// HasKeepalive returns a boolean if a field has been set.
func (o *MsgVpnClient) HasKeepalive() bool {
	if o != nil && o.Keepalive != nil {
		return true
	}

	return false
}

// SetKeepalive gets a reference to the given bool and assigns it to the Keepalive field.
func (o *MsgVpnClient) SetKeepalive(v bool) {
	o.Keepalive = &v
}

// GetKeepaliveEffectiveTimeout returns the KeepaliveEffectiveTimeout field value if set, zero value otherwise.
func (o *MsgVpnClient) GetKeepaliveEffectiveTimeout() int32 {
	if o == nil || o.KeepaliveEffectiveTimeout == nil {
		var ret int32
		return ret
	}
	return *o.KeepaliveEffectiveTimeout
}

// GetKeepaliveEffectiveTimeoutOk returns a tuple with the KeepaliveEffectiveTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetKeepaliveEffectiveTimeoutOk() (*int32, bool) {
	if o == nil || o.KeepaliveEffectiveTimeout == nil {
		return nil, false
	}
	return o.KeepaliveEffectiveTimeout, true
}

// HasKeepaliveEffectiveTimeout returns a boolean if a field has been set.
func (o *MsgVpnClient) HasKeepaliveEffectiveTimeout() bool {
	if o != nil && o.KeepaliveEffectiveTimeout != nil {
		return true
	}

	return false
}

// SetKeepaliveEffectiveTimeout gets a reference to the given int32 and assigns it to the KeepaliveEffectiveTimeout field.
func (o *MsgVpnClient) SetKeepaliveEffectiveTimeout(v int32) {
	o.KeepaliveEffectiveTimeout = &v
}

// GetLargeMsgEventRaised returns the LargeMsgEventRaised field value if set, zero value otherwise.
func (o *MsgVpnClient) GetLargeMsgEventRaised() bool {
	if o == nil || o.LargeMsgEventRaised == nil {
		var ret bool
		return ret
	}
	return *o.LargeMsgEventRaised
}

// GetLargeMsgEventRaisedOk returns a tuple with the LargeMsgEventRaised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetLargeMsgEventRaisedOk() (*bool, bool) {
	if o == nil || o.LargeMsgEventRaised == nil {
		return nil, false
	}
	return o.LargeMsgEventRaised, true
}

// HasLargeMsgEventRaised returns a boolean if a field has been set.
func (o *MsgVpnClient) HasLargeMsgEventRaised() bool {
	if o != nil && o.LargeMsgEventRaised != nil {
		return true
	}

	return false
}

// SetLargeMsgEventRaised gets a reference to the given bool and assigns it to the LargeMsgEventRaised field.
func (o *MsgVpnClient) SetLargeMsgEventRaised(v bool) {
	o.LargeMsgEventRaised = &v
}

// GetLoginRxMsgCount returns the LoginRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetLoginRxMsgCount() int64 {
	if o == nil || o.LoginRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginRxMsgCount
}

// GetLoginRxMsgCountOk returns a tuple with the LoginRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetLoginRxMsgCountOk() (*int64, bool) {
	if o == nil || o.LoginRxMsgCount == nil {
		return nil, false
	}
	return o.LoginRxMsgCount, true
}

// HasLoginRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasLoginRxMsgCount() bool {
	if o != nil && o.LoginRxMsgCount != nil {
		return true
	}

	return false
}

// SetLoginRxMsgCount gets a reference to the given int64 and assigns it to the LoginRxMsgCount field.
func (o *MsgVpnClient) SetLoginRxMsgCount(v int64) {
	o.LoginRxMsgCount = &v
}

// GetLoginTxMsgCount returns the LoginTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetLoginTxMsgCount() int64 {
	if o == nil || o.LoginTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginTxMsgCount
}

// GetLoginTxMsgCountOk returns a tuple with the LoginTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetLoginTxMsgCountOk() (*int64, bool) {
	if o == nil || o.LoginTxMsgCount == nil {
		return nil, false
	}
	return o.LoginTxMsgCount, true
}

// HasLoginTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasLoginTxMsgCount() bool {
	if o != nil && o.LoginTxMsgCount != nil {
		return true
	}

	return false
}

// SetLoginTxMsgCount gets a reference to the given int64 and assigns it to the LoginTxMsgCount field.
func (o *MsgVpnClient) SetLoginTxMsgCount(v int64) {
	o.LoginTxMsgCount = &v
}

// GetMaxBindCountExceededBindFailureCount returns the MaxBindCountExceededBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMaxBindCountExceededBindFailureCount() int64 {
	if o == nil || o.MaxBindCountExceededBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxBindCountExceededBindFailureCount
}

// GetMaxBindCountExceededBindFailureCountOk returns a tuple with the MaxBindCountExceededBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMaxBindCountExceededBindFailureCountOk() (*int64, bool) {
	if o == nil || o.MaxBindCountExceededBindFailureCount == nil {
		return nil, false
	}
	return o.MaxBindCountExceededBindFailureCount, true
}

// HasMaxBindCountExceededBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMaxBindCountExceededBindFailureCount() bool {
	if o != nil && o.MaxBindCountExceededBindFailureCount != nil {
		return true
	}

	return false
}

// SetMaxBindCountExceededBindFailureCount gets a reference to the given int64 and assigns it to the MaxBindCountExceededBindFailureCount field.
func (o *MsgVpnClient) SetMaxBindCountExceededBindFailureCount(v int64) {
	o.MaxBindCountExceededBindFailureCount = &v
}

// GetMaxElidingTopicCountEventRaised returns the MaxElidingTopicCountEventRaised field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMaxElidingTopicCountEventRaised() bool {
	if o == nil || o.MaxElidingTopicCountEventRaised == nil {
		var ret bool
		return ret
	}
	return *o.MaxElidingTopicCountEventRaised
}

// GetMaxElidingTopicCountEventRaisedOk returns a tuple with the MaxElidingTopicCountEventRaised field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMaxElidingTopicCountEventRaisedOk() (*bool, bool) {
	if o == nil || o.MaxElidingTopicCountEventRaised == nil {
		return nil, false
	}
	return o.MaxElidingTopicCountEventRaised, true
}

// HasMaxElidingTopicCountEventRaised returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMaxElidingTopicCountEventRaised() bool {
	if o != nil && o.MaxElidingTopicCountEventRaised != nil {
		return true
	}

	return false
}

// SetMaxElidingTopicCountEventRaised gets a reference to the given bool and assigns it to the MaxElidingTopicCountEventRaised field.
func (o *MsgVpnClient) SetMaxElidingTopicCountEventRaised(v bool) {
	o.MaxElidingTopicCountEventRaised = &v
}

// GetMqttConnackErrorTxCount returns the MqttConnackErrorTxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttConnackErrorTxCount() int64 {
	if o == nil || o.MqttConnackErrorTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttConnackErrorTxCount
}

// GetMqttConnackErrorTxCountOk returns a tuple with the MqttConnackErrorTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttConnackErrorTxCountOk() (*int64, bool) {
	if o == nil || o.MqttConnackErrorTxCount == nil {
		return nil, false
	}
	return o.MqttConnackErrorTxCount, true
}

// HasMqttConnackErrorTxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttConnackErrorTxCount() bool {
	if o != nil && o.MqttConnackErrorTxCount != nil {
		return true
	}

	return false
}

// SetMqttConnackErrorTxCount gets a reference to the given int64 and assigns it to the MqttConnackErrorTxCount field.
func (o *MsgVpnClient) SetMqttConnackErrorTxCount(v int64) {
	o.MqttConnackErrorTxCount = &v
}

// GetMqttConnackTxCount returns the MqttConnackTxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttConnackTxCount() int64 {
	if o == nil || o.MqttConnackTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttConnackTxCount
}

// GetMqttConnackTxCountOk returns a tuple with the MqttConnackTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttConnackTxCountOk() (*int64, bool) {
	if o == nil || o.MqttConnackTxCount == nil {
		return nil, false
	}
	return o.MqttConnackTxCount, true
}

// HasMqttConnackTxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttConnackTxCount() bool {
	if o != nil && o.MqttConnackTxCount != nil {
		return true
	}

	return false
}

// SetMqttConnackTxCount gets a reference to the given int64 and assigns it to the MqttConnackTxCount field.
func (o *MsgVpnClient) SetMqttConnackTxCount(v int64) {
	o.MqttConnackTxCount = &v
}

// GetMqttConnectRxCount returns the MqttConnectRxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttConnectRxCount() int64 {
	if o == nil || o.MqttConnectRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttConnectRxCount
}

// GetMqttConnectRxCountOk returns a tuple with the MqttConnectRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttConnectRxCountOk() (*int64, bool) {
	if o == nil || o.MqttConnectRxCount == nil {
		return nil, false
	}
	return o.MqttConnectRxCount, true
}

// HasMqttConnectRxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttConnectRxCount() bool {
	if o != nil && o.MqttConnectRxCount != nil {
		return true
	}

	return false
}

// SetMqttConnectRxCount gets a reference to the given int64 and assigns it to the MqttConnectRxCount field.
func (o *MsgVpnClient) SetMqttConnectRxCount(v int64) {
	o.MqttConnectRxCount = &v
}

// GetMqttDisconnectRxCount returns the MqttDisconnectRxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttDisconnectRxCount() int64 {
	if o == nil || o.MqttDisconnectRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttDisconnectRxCount
}

// GetMqttDisconnectRxCountOk returns a tuple with the MqttDisconnectRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttDisconnectRxCountOk() (*int64, bool) {
	if o == nil || o.MqttDisconnectRxCount == nil {
		return nil, false
	}
	return o.MqttDisconnectRxCount, true
}

// HasMqttDisconnectRxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttDisconnectRxCount() bool {
	if o != nil && o.MqttDisconnectRxCount != nil {
		return true
	}

	return false
}

// SetMqttDisconnectRxCount gets a reference to the given int64 and assigns it to the MqttDisconnectRxCount field.
func (o *MsgVpnClient) SetMqttDisconnectRxCount(v int64) {
	o.MqttDisconnectRxCount = &v
}

// GetMqttPingreqRxCount returns the MqttPingreqRxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPingreqRxCount() int64 {
	if o == nil || o.MqttPingreqRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPingreqRxCount
}

// GetMqttPingreqRxCountOk returns a tuple with the MqttPingreqRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPingreqRxCountOk() (*int64, bool) {
	if o == nil || o.MqttPingreqRxCount == nil {
		return nil, false
	}
	return o.MqttPingreqRxCount, true
}

// HasMqttPingreqRxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPingreqRxCount() bool {
	if o != nil && o.MqttPingreqRxCount != nil {
		return true
	}

	return false
}

// SetMqttPingreqRxCount gets a reference to the given int64 and assigns it to the MqttPingreqRxCount field.
func (o *MsgVpnClient) SetMqttPingreqRxCount(v int64) {
	o.MqttPingreqRxCount = &v
}

// GetMqttPingrespTxCount returns the MqttPingrespTxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPingrespTxCount() int64 {
	if o == nil || o.MqttPingrespTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPingrespTxCount
}

// GetMqttPingrespTxCountOk returns a tuple with the MqttPingrespTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPingrespTxCountOk() (*int64, bool) {
	if o == nil || o.MqttPingrespTxCount == nil {
		return nil, false
	}
	return o.MqttPingrespTxCount, true
}

// HasMqttPingrespTxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPingrespTxCount() bool {
	if o != nil && o.MqttPingrespTxCount != nil {
		return true
	}

	return false
}

// SetMqttPingrespTxCount gets a reference to the given int64 and assigns it to the MqttPingrespTxCount field.
func (o *MsgVpnClient) SetMqttPingrespTxCount(v int64) {
	o.MqttPingrespTxCount = &v
}

// GetMqttPubackRxCount returns the MqttPubackRxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPubackRxCount() int64 {
	if o == nil || o.MqttPubackRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubackRxCount
}

// GetMqttPubackRxCountOk returns a tuple with the MqttPubackRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPubackRxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubackRxCount == nil {
		return nil, false
	}
	return o.MqttPubackRxCount, true
}

// HasMqttPubackRxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPubackRxCount() bool {
	if o != nil && o.MqttPubackRxCount != nil {
		return true
	}

	return false
}

// SetMqttPubackRxCount gets a reference to the given int64 and assigns it to the MqttPubackRxCount field.
func (o *MsgVpnClient) SetMqttPubackRxCount(v int64) {
	o.MqttPubackRxCount = &v
}

// GetMqttPubackTxCount returns the MqttPubackTxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPubackTxCount() int64 {
	if o == nil || o.MqttPubackTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubackTxCount
}

// GetMqttPubackTxCountOk returns a tuple with the MqttPubackTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPubackTxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubackTxCount == nil {
		return nil, false
	}
	return o.MqttPubackTxCount, true
}

// HasMqttPubackTxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPubackTxCount() bool {
	if o != nil && o.MqttPubackTxCount != nil {
		return true
	}

	return false
}

// SetMqttPubackTxCount gets a reference to the given int64 and assigns it to the MqttPubackTxCount field.
func (o *MsgVpnClient) SetMqttPubackTxCount(v int64) {
	o.MqttPubackTxCount = &v
}

// GetMqttPubcompTxCount returns the MqttPubcompTxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPubcompTxCount() int64 {
	if o == nil || o.MqttPubcompTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubcompTxCount
}

// GetMqttPubcompTxCountOk returns a tuple with the MqttPubcompTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPubcompTxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubcompTxCount == nil {
		return nil, false
	}
	return o.MqttPubcompTxCount, true
}

// HasMqttPubcompTxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPubcompTxCount() bool {
	if o != nil && o.MqttPubcompTxCount != nil {
		return true
	}

	return false
}

// SetMqttPubcompTxCount gets a reference to the given int64 and assigns it to the MqttPubcompTxCount field.
func (o *MsgVpnClient) SetMqttPubcompTxCount(v int64) {
	o.MqttPubcompTxCount = &v
}

// GetMqttPublishQos0RxCount returns the MqttPublishQos0RxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPublishQos0RxCount() int64 {
	if o == nil || o.MqttPublishQos0RxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos0RxCount
}

// GetMqttPublishQos0RxCountOk returns a tuple with the MqttPublishQos0RxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPublishQos0RxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos0RxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos0RxCount, true
}

// HasMqttPublishQos0RxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPublishQos0RxCount() bool {
	if o != nil && o.MqttPublishQos0RxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos0RxCount gets a reference to the given int64 and assigns it to the MqttPublishQos0RxCount field.
func (o *MsgVpnClient) SetMqttPublishQos0RxCount(v int64) {
	o.MqttPublishQos0RxCount = &v
}

// GetMqttPublishQos0TxCount returns the MqttPublishQos0TxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPublishQos0TxCount() int64 {
	if o == nil || o.MqttPublishQos0TxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos0TxCount
}

// GetMqttPublishQos0TxCountOk returns a tuple with the MqttPublishQos0TxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPublishQos0TxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos0TxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos0TxCount, true
}

// HasMqttPublishQos0TxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPublishQos0TxCount() bool {
	if o != nil && o.MqttPublishQos0TxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos0TxCount gets a reference to the given int64 and assigns it to the MqttPublishQos0TxCount field.
func (o *MsgVpnClient) SetMqttPublishQos0TxCount(v int64) {
	o.MqttPublishQos0TxCount = &v
}

// GetMqttPublishQos1RxCount returns the MqttPublishQos1RxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPublishQos1RxCount() int64 {
	if o == nil || o.MqttPublishQos1RxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos1RxCount
}

// GetMqttPublishQos1RxCountOk returns a tuple with the MqttPublishQos1RxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPublishQos1RxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos1RxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos1RxCount, true
}

// HasMqttPublishQos1RxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPublishQos1RxCount() bool {
	if o != nil && o.MqttPublishQos1RxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos1RxCount gets a reference to the given int64 and assigns it to the MqttPublishQos1RxCount field.
func (o *MsgVpnClient) SetMqttPublishQos1RxCount(v int64) {
	o.MqttPublishQos1RxCount = &v
}

// GetMqttPublishQos1TxCount returns the MqttPublishQos1TxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPublishQos1TxCount() int64 {
	if o == nil || o.MqttPublishQos1TxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos1TxCount
}

// GetMqttPublishQos1TxCountOk returns a tuple with the MqttPublishQos1TxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPublishQos1TxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos1TxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos1TxCount, true
}

// HasMqttPublishQos1TxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPublishQos1TxCount() bool {
	if o != nil && o.MqttPublishQos1TxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos1TxCount gets a reference to the given int64 and assigns it to the MqttPublishQos1TxCount field.
func (o *MsgVpnClient) SetMqttPublishQos1TxCount(v int64) {
	o.MqttPublishQos1TxCount = &v
}

// GetMqttPublishQos2RxCount returns the MqttPublishQos2RxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPublishQos2RxCount() int64 {
	if o == nil || o.MqttPublishQos2RxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos2RxCount
}

// GetMqttPublishQos2RxCountOk returns a tuple with the MqttPublishQos2RxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPublishQos2RxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos2RxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos2RxCount, true
}

// HasMqttPublishQos2RxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPublishQos2RxCount() bool {
	if o != nil && o.MqttPublishQos2RxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos2RxCount gets a reference to the given int64 and assigns it to the MqttPublishQos2RxCount field.
func (o *MsgVpnClient) SetMqttPublishQos2RxCount(v int64) {
	o.MqttPublishQos2RxCount = &v
}

// GetMqttPubrecTxCount returns the MqttPubrecTxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPubrecTxCount() int64 {
	if o == nil || o.MqttPubrecTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubrecTxCount
}

// GetMqttPubrecTxCountOk returns a tuple with the MqttPubrecTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPubrecTxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubrecTxCount == nil {
		return nil, false
	}
	return o.MqttPubrecTxCount, true
}

// HasMqttPubrecTxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPubrecTxCount() bool {
	if o != nil && o.MqttPubrecTxCount != nil {
		return true
	}

	return false
}

// SetMqttPubrecTxCount gets a reference to the given int64 and assigns it to the MqttPubrecTxCount field.
func (o *MsgVpnClient) SetMqttPubrecTxCount(v int64) {
	o.MqttPubrecTxCount = &v
}

// GetMqttPubrelRxCount returns the MqttPubrelRxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttPubrelRxCount() int64 {
	if o == nil || o.MqttPubrelRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubrelRxCount
}

// GetMqttPubrelRxCountOk returns a tuple with the MqttPubrelRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttPubrelRxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubrelRxCount == nil {
		return nil, false
	}
	return o.MqttPubrelRxCount, true
}

// HasMqttPubrelRxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttPubrelRxCount() bool {
	if o != nil && o.MqttPubrelRxCount != nil {
		return true
	}

	return false
}

// SetMqttPubrelRxCount gets a reference to the given int64 and assigns it to the MqttPubrelRxCount field.
func (o *MsgVpnClient) SetMqttPubrelRxCount(v int64) {
	o.MqttPubrelRxCount = &v
}

// GetMqttSubackErrorTxCount returns the MqttSubackErrorTxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttSubackErrorTxCount() int64 {
	if o == nil || o.MqttSubackErrorTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttSubackErrorTxCount
}

// GetMqttSubackErrorTxCountOk returns a tuple with the MqttSubackErrorTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttSubackErrorTxCountOk() (*int64, bool) {
	if o == nil || o.MqttSubackErrorTxCount == nil {
		return nil, false
	}
	return o.MqttSubackErrorTxCount, true
}

// HasMqttSubackErrorTxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttSubackErrorTxCount() bool {
	if o != nil && o.MqttSubackErrorTxCount != nil {
		return true
	}

	return false
}

// SetMqttSubackErrorTxCount gets a reference to the given int64 and assigns it to the MqttSubackErrorTxCount field.
func (o *MsgVpnClient) SetMqttSubackErrorTxCount(v int64) {
	o.MqttSubackErrorTxCount = &v
}

// GetMqttSubackTxCount returns the MqttSubackTxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttSubackTxCount() int64 {
	if o == nil || o.MqttSubackTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttSubackTxCount
}

// GetMqttSubackTxCountOk returns a tuple with the MqttSubackTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttSubackTxCountOk() (*int64, bool) {
	if o == nil || o.MqttSubackTxCount == nil {
		return nil, false
	}
	return o.MqttSubackTxCount, true
}

// HasMqttSubackTxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttSubackTxCount() bool {
	if o != nil && o.MqttSubackTxCount != nil {
		return true
	}

	return false
}

// SetMqttSubackTxCount gets a reference to the given int64 and assigns it to the MqttSubackTxCount field.
func (o *MsgVpnClient) SetMqttSubackTxCount(v int64) {
	o.MqttSubackTxCount = &v
}

// GetMqttSubscribeRxCount returns the MqttSubscribeRxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttSubscribeRxCount() int64 {
	if o == nil || o.MqttSubscribeRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttSubscribeRxCount
}

// GetMqttSubscribeRxCountOk returns a tuple with the MqttSubscribeRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttSubscribeRxCountOk() (*int64, bool) {
	if o == nil || o.MqttSubscribeRxCount == nil {
		return nil, false
	}
	return o.MqttSubscribeRxCount, true
}

// HasMqttSubscribeRxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttSubscribeRxCount() bool {
	if o != nil && o.MqttSubscribeRxCount != nil {
		return true
	}

	return false
}

// SetMqttSubscribeRxCount gets a reference to the given int64 and assigns it to the MqttSubscribeRxCount field.
func (o *MsgVpnClient) SetMqttSubscribeRxCount(v int64) {
	o.MqttSubscribeRxCount = &v
}

// GetMqttUnsubackTxCount returns the MqttUnsubackTxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttUnsubackTxCount() int64 {
	if o == nil || o.MqttUnsubackTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttUnsubackTxCount
}

// GetMqttUnsubackTxCountOk returns a tuple with the MqttUnsubackTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttUnsubackTxCountOk() (*int64, bool) {
	if o == nil || o.MqttUnsubackTxCount == nil {
		return nil, false
	}
	return o.MqttUnsubackTxCount, true
}

// HasMqttUnsubackTxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttUnsubackTxCount() bool {
	if o != nil && o.MqttUnsubackTxCount != nil {
		return true
	}

	return false
}

// SetMqttUnsubackTxCount gets a reference to the given int64 and assigns it to the MqttUnsubackTxCount field.
func (o *MsgVpnClient) SetMqttUnsubackTxCount(v int64) {
	o.MqttUnsubackTxCount = &v
}

// GetMqttUnsubscribeRxCount returns the MqttUnsubscribeRxCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMqttUnsubscribeRxCount() int64 {
	if o == nil || o.MqttUnsubscribeRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttUnsubscribeRxCount
}

// GetMqttUnsubscribeRxCountOk returns a tuple with the MqttUnsubscribeRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMqttUnsubscribeRxCountOk() (*int64, bool) {
	if o == nil || o.MqttUnsubscribeRxCount == nil {
		return nil, false
	}
	return o.MqttUnsubscribeRxCount, true
}

// HasMqttUnsubscribeRxCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMqttUnsubscribeRxCount() bool {
	if o != nil && o.MqttUnsubscribeRxCount != nil {
		return true
	}

	return false
}

// SetMqttUnsubscribeRxCount gets a reference to the given int64 and assigns it to the MqttUnsubscribeRxCount field.
func (o *MsgVpnClient) SetMqttUnsubscribeRxCount(v int64) {
	o.MqttUnsubscribeRxCount = &v
}

// GetMsgSpoolCongestionRxDiscardedMsgCount returns the MsgSpoolCongestionRxDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMsgSpoolCongestionRxDiscardedMsgCount() int64 {
	if o == nil || o.MsgSpoolCongestionRxDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MsgSpoolCongestionRxDiscardedMsgCount
}

// GetMsgSpoolCongestionRxDiscardedMsgCountOk returns a tuple with the MsgSpoolCongestionRxDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMsgSpoolCongestionRxDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.MsgSpoolCongestionRxDiscardedMsgCount == nil {
		return nil, false
	}
	return o.MsgSpoolCongestionRxDiscardedMsgCount, true
}

// HasMsgSpoolCongestionRxDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMsgSpoolCongestionRxDiscardedMsgCount() bool {
	if o != nil && o.MsgSpoolCongestionRxDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetMsgSpoolCongestionRxDiscardedMsgCount gets a reference to the given int64 and assigns it to the MsgSpoolCongestionRxDiscardedMsgCount field.
func (o *MsgVpnClient) SetMsgSpoolCongestionRxDiscardedMsgCount(v int64) {
	o.MsgSpoolCongestionRxDiscardedMsgCount = &v
}

// GetMsgSpoolRxDiscardedMsgCount returns the MsgSpoolRxDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMsgSpoolRxDiscardedMsgCount() int64 {
	if o == nil || o.MsgSpoolRxDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MsgSpoolRxDiscardedMsgCount
}

// GetMsgSpoolRxDiscardedMsgCountOk returns a tuple with the MsgSpoolRxDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMsgSpoolRxDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.MsgSpoolRxDiscardedMsgCount == nil {
		return nil, false
	}
	return o.MsgSpoolRxDiscardedMsgCount, true
}

// HasMsgSpoolRxDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMsgSpoolRxDiscardedMsgCount() bool {
	if o != nil && o.MsgSpoolRxDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetMsgSpoolRxDiscardedMsgCount gets a reference to the given int64 and assigns it to the MsgSpoolRxDiscardedMsgCount field.
func (o *MsgVpnClient) SetMsgSpoolRxDiscardedMsgCount(v int64) {
	o.MsgSpoolRxDiscardedMsgCount = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnClient) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnClient) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnClient) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetNoLocalDelivery returns the NoLocalDelivery field value if set, zero value otherwise.
func (o *MsgVpnClient) GetNoLocalDelivery() bool {
	if o == nil || o.NoLocalDelivery == nil {
		var ret bool
		return ret
	}
	return *o.NoLocalDelivery
}

// GetNoLocalDeliveryOk returns a tuple with the NoLocalDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetNoLocalDeliveryOk() (*bool, bool) {
	if o == nil || o.NoLocalDelivery == nil {
		return nil, false
	}
	return o.NoLocalDelivery, true
}

// HasNoLocalDelivery returns a boolean if a field has been set.
func (o *MsgVpnClient) HasNoLocalDelivery() bool {
	if o != nil && o.NoLocalDelivery != nil {
		return true
	}

	return false
}

// SetNoLocalDelivery gets a reference to the given bool and assigns it to the NoLocalDelivery field.
func (o *MsgVpnClient) SetNoLocalDelivery(v bool) {
	o.NoLocalDelivery = &v
}

// GetNoSubscriptionMatchRxDiscardedMsgCount returns the NoSubscriptionMatchRxDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetNoSubscriptionMatchRxDiscardedMsgCount() int64 {
	if o == nil || o.NoSubscriptionMatchRxDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.NoSubscriptionMatchRxDiscardedMsgCount
}

// GetNoSubscriptionMatchRxDiscardedMsgCountOk returns a tuple with the NoSubscriptionMatchRxDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetNoSubscriptionMatchRxDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.NoSubscriptionMatchRxDiscardedMsgCount == nil {
		return nil, false
	}
	return o.NoSubscriptionMatchRxDiscardedMsgCount, true
}

// HasNoSubscriptionMatchRxDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasNoSubscriptionMatchRxDiscardedMsgCount() bool {
	if o != nil && o.NoSubscriptionMatchRxDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetNoSubscriptionMatchRxDiscardedMsgCount gets a reference to the given int64 and assigns it to the NoSubscriptionMatchRxDiscardedMsgCount field.
func (o *MsgVpnClient) SetNoSubscriptionMatchRxDiscardedMsgCount(v int64) {
	o.NoSubscriptionMatchRxDiscardedMsgCount = &v
}

// GetOriginalClientUsername returns the OriginalClientUsername field value if set, zero value otherwise.
func (o *MsgVpnClient) GetOriginalClientUsername() string {
	if o == nil || o.OriginalClientUsername == nil {
		var ret string
		return ret
	}
	return *o.OriginalClientUsername
}

// GetOriginalClientUsernameOk returns a tuple with the OriginalClientUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetOriginalClientUsernameOk() (*string, bool) {
	if o == nil || o.OriginalClientUsername == nil {
		return nil, false
	}
	return o.OriginalClientUsername, true
}

// HasOriginalClientUsername returns a boolean if a field has been set.
func (o *MsgVpnClient) HasOriginalClientUsername() bool {
	if o != nil && o.OriginalClientUsername != nil {
		return true
	}

	return false
}

// SetOriginalClientUsername gets a reference to the given string and assigns it to the OriginalClientUsername field.
func (o *MsgVpnClient) SetOriginalClientUsername(v string) {
	o.OriginalClientUsername = &v
}

// GetOtherBindFailureCount returns the OtherBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetOtherBindFailureCount() int64 {
	if o == nil || o.OtherBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.OtherBindFailureCount
}

// GetOtherBindFailureCountOk returns a tuple with the OtherBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetOtherBindFailureCountOk() (*int64, bool) {
	if o == nil || o.OtherBindFailureCount == nil {
		return nil, false
	}
	return o.OtherBindFailureCount, true
}

// HasOtherBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasOtherBindFailureCount() bool {
	if o != nil && o.OtherBindFailureCount != nil {
		return true
	}

	return false
}

// SetOtherBindFailureCount gets a reference to the given int64 and assigns it to the OtherBindFailureCount field.
func (o *MsgVpnClient) SetOtherBindFailureCount(v int64) {
	o.OtherBindFailureCount = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *MsgVpnClient) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *MsgVpnClient) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *MsgVpnClient) SetPlatform(v string) {
	o.Platform = &v
}

// GetPublishTopicAclRxDiscardedMsgCount returns the PublishTopicAclRxDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetPublishTopicAclRxDiscardedMsgCount() int64 {
	if o == nil || o.PublishTopicAclRxDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.PublishTopicAclRxDiscardedMsgCount
}

// GetPublishTopicAclRxDiscardedMsgCountOk returns a tuple with the PublishTopicAclRxDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetPublishTopicAclRxDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.PublishTopicAclRxDiscardedMsgCount == nil {
		return nil, false
	}
	return o.PublishTopicAclRxDiscardedMsgCount, true
}

// HasPublishTopicAclRxDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasPublishTopicAclRxDiscardedMsgCount() bool {
	if o != nil && o.PublishTopicAclRxDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetPublishTopicAclRxDiscardedMsgCount gets a reference to the given int64 and assigns it to the PublishTopicAclRxDiscardedMsgCount field.
func (o *MsgVpnClient) SetPublishTopicAclRxDiscardedMsgCount(v int64) {
	o.PublishTopicAclRxDiscardedMsgCount = &v
}

// GetRestHttpRequestRxByteCount returns the RestHttpRequestRxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpRequestRxByteCount() int64 {
	if o == nil || o.RestHttpRequestRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpRequestRxByteCount
}

// GetRestHttpRequestRxByteCountOk returns a tuple with the RestHttpRequestRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpRequestRxByteCountOk() (*int64, bool) {
	if o == nil || o.RestHttpRequestRxByteCount == nil {
		return nil, false
	}
	return o.RestHttpRequestRxByteCount, true
}

// HasRestHttpRequestRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpRequestRxByteCount() bool {
	if o != nil && o.RestHttpRequestRxByteCount != nil {
		return true
	}

	return false
}

// SetRestHttpRequestRxByteCount gets a reference to the given int64 and assigns it to the RestHttpRequestRxByteCount field.
func (o *MsgVpnClient) SetRestHttpRequestRxByteCount(v int64) {
	o.RestHttpRequestRxByteCount = &v
}

// GetRestHttpRequestRxMsgCount returns the RestHttpRequestRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpRequestRxMsgCount() int64 {
	if o == nil || o.RestHttpRequestRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpRequestRxMsgCount
}

// GetRestHttpRequestRxMsgCountOk returns a tuple with the RestHttpRequestRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpRequestRxMsgCountOk() (*int64, bool) {
	if o == nil || o.RestHttpRequestRxMsgCount == nil {
		return nil, false
	}
	return o.RestHttpRequestRxMsgCount, true
}

// HasRestHttpRequestRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpRequestRxMsgCount() bool {
	if o != nil && o.RestHttpRequestRxMsgCount != nil {
		return true
	}

	return false
}

// SetRestHttpRequestRxMsgCount gets a reference to the given int64 and assigns it to the RestHttpRequestRxMsgCount field.
func (o *MsgVpnClient) SetRestHttpRequestRxMsgCount(v int64) {
	o.RestHttpRequestRxMsgCount = &v
}

// GetRestHttpRequestTxByteCount returns the RestHttpRequestTxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpRequestTxByteCount() int64 {
	if o == nil || o.RestHttpRequestTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpRequestTxByteCount
}

// GetRestHttpRequestTxByteCountOk returns a tuple with the RestHttpRequestTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpRequestTxByteCountOk() (*int64, bool) {
	if o == nil || o.RestHttpRequestTxByteCount == nil {
		return nil, false
	}
	return o.RestHttpRequestTxByteCount, true
}

// HasRestHttpRequestTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpRequestTxByteCount() bool {
	if o != nil && o.RestHttpRequestTxByteCount != nil {
		return true
	}

	return false
}

// SetRestHttpRequestTxByteCount gets a reference to the given int64 and assigns it to the RestHttpRequestTxByteCount field.
func (o *MsgVpnClient) SetRestHttpRequestTxByteCount(v int64) {
	o.RestHttpRequestTxByteCount = &v
}

// GetRestHttpRequestTxMsgCount returns the RestHttpRequestTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpRequestTxMsgCount() int64 {
	if o == nil || o.RestHttpRequestTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpRequestTxMsgCount
}

// GetRestHttpRequestTxMsgCountOk returns a tuple with the RestHttpRequestTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpRequestTxMsgCountOk() (*int64, bool) {
	if o == nil || o.RestHttpRequestTxMsgCount == nil {
		return nil, false
	}
	return o.RestHttpRequestTxMsgCount, true
}

// HasRestHttpRequestTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpRequestTxMsgCount() bool {
	if o != nil && o.RestHttpRequestTxMsgCount != nil {
		return true
	}

	return false
}

// SetRestHttpRequestTxMsgCount gets a reference to the given int64 and assigns it to the RestHttpRequestTxMsgCount field.
func (o *MsgVpnClient) SetRestHttpRequestTxMsgCount(v int64) {
	o.RestHttpRequestTxMsgCount = &v
}

// GetRestHttpResponseErrorRxMsgCount returns the RestHttpResponseErrorRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpResponseErrorRxMsgCount() int64 {
	if o == nil || o.RestHttpResponseErrorRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpResponseErrorRxMsgCount
}

// GetRestHttpResponseErrorRxMsgCountOk returns a tuple with the RestHttpResponseErrorRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpResponseErrorRxMsgCountOk() (*int64, bool) {
	if o == nil || o.RestHttpResponseErrorRxMsgCount == nil {
		return nil, false
	}
	return o.RestHttpResponseErrorRxMsgCount, true
}

// HasRestHttpResponseErrorRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpResponseErrorRxMsgCount() bool {
	if o != nil && o.RestHttpResponseErrorRxMsgCount != nil {
		return true
	}

	return false
}

// SetRestHttpResponseErrorRxMsgCount gets a reference to the given int64 and assigns it to the RestHttpResponseErrorRxMsgCount field.
func (o *MsgVpnClient) SetRestHttpResponseErrorRxMsgCount(v int64) {
	o.RestHttpResponseErrorRxMsgCount = &v
}

// GetRestHttpResponseErrorTxMsgCount returns the RestHttpResponseErrorTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpResponseErrorTxMsgCount() int64 {
	if o == nil || o.RestHttpResponseErrorTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpResponseErrorTxMsgCount
}

// GetRestHttpResponseErrorTxMsgCountOk returns a tuple with the RestHttpResponseErrorTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpResponseErrorTxMsgCountOk() (*int64, bool) {
	if o == nil || o.RestHttpResponseErrorTxMsgCount == nil {
		return nil, false
	}
	return o.RestHttpResponseErrorTxMsgCount, true
}

// HasRestHttpResponseErrorTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpResponseErrorTxMsgCount() bool {
	if o != nil && o.RestHttpResponseErrorTxMsgCount != nil {
		return true
	}

	return false
}

// SetRestHttpResponseErrorTxMsgCount gets a reference to the given int64 and assigns it to the RestHttpResponseErrorTxMsgCount field.
func (o *MsgVpnClient) SetRestHttpResponseErrorTxMsgCount(v int64) {
	o.RestHttpResponseErrorTxMsgCount = &v
}

// GetRestHttpResponseRxByteCount returns the RestHttpResponseRxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpResponseRxByteCount() int64 {
	if o == nil || o.RestHttpResponseRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpResponseRxByteCount
}

// GetRestHttpResponseRxByteCountOk returns a tuple with the RestHttpResponseRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpResponseRxByteCountOk() (*int64, bool) {
	if o == nil || o.RestHttpResponseRxByteCount == nil {
		return nil, false
	}
	return o.RestHttpResponseRxByteCount, true
}

// HasRestHttpResponseRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpResponseRxByteCount() bool {
	if o != nil && o.RestHttpResponseRxByteCount != nil {
		return true
	}

	return false
}

// SetRestHttpResponseRxByteCount gets a reference to the given int64 and assigns it to the RestHttpResponseRxByteCount field.
func (o *MsgVpnClient) SetRestHttpResponseRxByteCount(v int64) {
	o.RestHttpResponseRxByteCount = &v
}

// GetRestHttpResponseRxMsgCount returns the RestHttpResponseRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpResponseRxMsgCount() int64 {
	if o == nil || o.RestHttpResponseRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpResponseRxMsgCount
}

// GetRestHttpResponseRxMsgCountOk returns a tuple with the RestHttpResponseRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpResponseRxMsgCountOk() (*int64, bool) {
	if o == nil || o.RestHttpResponseRxMsgCount == nil {
		return nil, false
	}
	return o.RestHttpResponseRxMsgCount, true
}

// HasRestHttpResponseRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpResponseRxMsgCount() bool {
	if o != nil && o.RestHttpResponseRxMsgCount != nil {
		return true
	}

	return false
}

// SetRestHttpResponseRxMsgCount gets a reference to the given int64 and assigns it to the RestHttpResponseRxMsgCount field.
func (o *MsgVpnClient) SetRestHttpResponseRxMsgCount(v int64) {
	o.RestHttpResponseRxMsgCount = &v
}

// GetRestHttpResponseSuccessRxMsgCount returns the RestHttpResponseSuccessRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpResponseSuccessRxMsgCount() int64 {
	if o == nil || o.RestHttpResponseSuccessRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpResponseSuccessRxMsgCount
}

// GetRestHttpResponseSuccessRxMsgCountOk returns a tuple with the RestHttpResponseSuccessRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpResponseSuccessRxMsgCountOk() (*int64, bool) {
	if o == nil || o.RestHttpResponseSuccessRxMsgCount == nil {
		return nil, false
	}
	return o.RestHttpResponseSuccessRxMsgCount, true
}

// HasRestHttpResponseSuccessRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpResponseSuccessRxMsgCount() bool {
	if o != nil && o.RestHttpResponseSuccessRxMsgCount != nil {
		return true
	}

	return false
}

// SetRestHttpResponseSuccessRxMsgCount gets a reference to the given int64 and assigns it to the RestHttpResponseSuccessRxMsgCount field.
func (o *MsgVpnClient) SetRestHttpResponseSuccessRxMsgCount(v int64) {
	o.RestHttpResponseSuccessRxMsgCount = &v
}

// GetRestHttpResponseSuccessTxMsgCount returns the RestHttpResponseSuccessTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpResponseSuccessTxMsgCount() int64 {
	if o == nil || o.RestHttpResponseSuccessTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpResponseSuccessTxMsgCount
}

// GetRestHttpResponseSuccessTxMsgCountOk returns a tuple with the RestHttpResponseSuccessTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpResponseSuccessTxMsgCountOk() (*int64, bool) {
	if o == nil || o.RestHttpResponseSuccessTxMsgCount == nil {
		return nil, false
	}
	return o.RestHttpResponseSuccessTxMsgCount, true
}

// HasRestHttpResponseSuccessTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpResponseSuccessTxMsgCount() bool {
	if o != nil && o.RestHttpResponseSuccessTxMsgCount != nil {
		return true
	}

	return false
}

// SetRestHttpResponseSuccessTxMsgCount gets a reference to the given int64 and assigns it to the RestHttpResponseSuccessTxMsgCount field.
func (o *MsgVpnClient) SetRestHttpResponseSuccessTxMsgCount(v int64) {
	o.RestHttpResponseSuccessTxMsgCount = &v
}

// GetRestHttpResponseTimeoutRxMsgCount returns the RestHttpResponseTimeoutRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpResponseTimeoutRxMsgCount() int64 {
	if o == nil || o.RestHttpResponseTimeoutRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpResponseTimeoutRxMsgCount
}

// GetRestHttpResponseTimeoutRxMsgCountOk returns a tuple with the RestHttpResponseTimeoutRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpResponseTimeoutRxMsgCountOk() (*int64, bool) {
	if o == nil || o.RestHttpResponseTimeoutRxMsgCount == nil {
		return nil, false
	}
	return o.RestHttpResponseTimeoutRxMsgCount, true
}

// HasRestHttpResponseTimeoutRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpResponseTimeoutRxMsgCount() bool {
	if o != nil && o.RestHttpResponseTimeoutRxMsgCount != nil {
		return true
	}

	return false
}

// SetRestHttpResponseTimeoutRxMsgCount gets a reference to the given int64 and assigns it to the RestHttpResponseTimeoutRxMsgCount field.
func (o *MsgVpnClient) SetRestHttpResponseTimeoutRxMsgCount(v int64) {
	o.RestHttpResponseTimeoutRxMsgCount = &v
}

// GetRestHttpResponseTimeoutTxMsgCount returns the RestHttpResponseTimeoutTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpResponseTimeoutTxMsgCount() int64 {
	if o == nil || o.RestHttpResponseTimeoutTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpResponseTimeoutTxMsgCount
}

// GetRestHttpResponseTimeoutTxMsgCountOk returns a tuple with the RestHttpResponseTimeoutTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpResponseTimeoutTxMsgCountOk() (*int64, bool) {
	if o == nil || o.RestHttpResponseTimeoutTxMsgCount == nil {
		return nil, false
	}
	return o.RestHttpResponseTimeoutTxMsgCount, true
}

// HasRestHttpResponseTimeoutTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpResponseTimeoutTxMsgCount() bool {
	if o != nil && o.RestHttpResponseTimeoutTxMsgCount != nil {
		return true
	}

	return false
}

// SetRestHttpResponseTimeoutTxMsgCount gets a reference to the given int64 and assigns it to the RestHttpResponseTimeoutTxMsgCount field.
func (o *MsgVpnClient) SetRestHttpResponseTimeoutTxMsgCount(v int64) {
	o.RestHttpResponseTimeoutTxMsgCount = &v
}

// GetRestHttpResponseTxByteCount returns the RestHttpResponseTxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpResponseTxByteCount() int64 {
	if o == nil || o.RestHttpResponseTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpResponseTxByteCount
}

// GetRestHttpResponseTxByteCountOk returns a tuple with the RestHttpResponseTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpResponseTxByteCountOk() (*int64, bool) {
	if o == nil || o.RestHttpResponseTxByteCount == nil {
		return nil, false
	}
	return o.RestHttpResponseTxByteCount, true
}

// HasRestHttpResponseTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpResponseTxByteCount() bool {
	if o != nil && o.RestHttpResponseTxByteCount != nil {
		return true
	}

	return false
}

// SetRestHttpResponseTxByteCount gets a reference to the given int64 and assigns it to the RestHttpResponseTxByteCount field.
func (o *MsgVpnClient) SetRestHttpResponseTxByteCount(v int64) {
	o.RestHttpResponseTxByteCount = &v
}

// GetRestHttpResponseTxMsgCount returns the RestHttpResponseTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRestHttpResponseTxMsgCount() int64 {
	if o == nil || o.RestHttpResponseTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RestHttpResponseTxMsgCount
}

// GetRestHttpResponseTxMsgCountOk returns a tuple with the RestHttpResponseTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRestHttpResponseTxMsgCountOk() (*int64, bool) {
	if o == nil || o.RestHttpResponseTxMsgCount == nil {
		return nil, false
	}
	return o.RestHttpResponseTxMsgCount, true
}

// HasRestHttpResponseTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRestHttpResponseTxMsgCount() bool {
	if o != nil && o.RestHttpResponseTxMsgCount != nil {
		return true
	}

	return false
}

// SetRestHttpResponseTxMsgCount gets a reference to the given int64 and assigns it to the RestHttpResponseTxMsgCount field.
func (o *MsgVpnClient) SetRestHttpResponseTxMsgCount(v int64) {
	o.RestHttpResponseTxMsgCount = &v
}

// GetRxByteCount returns the RxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRxByteCount() int64 {
	if o == nil || o.RxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.RxByteCount
}

// GetRxByteCountOk returns a tuple with the RxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRxByteCountOk() (*int64, bool) {
	if o == nil || o.RxByteCount == nil {
		return nil, false
	}
	return o.RxByteCount, true
}

// HasRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRxByteCount() bool {
	if o != nil && o.RxByteCount != nil {
		return true
	}

	return false
}

// SetRxByteCount gets a reference to the given int64 and assigns it to the RxByteCount field.
func (o *MsgVpnClient) SetRxByteCount(v int64) {
	o.RxByteCount = &v
}

// GetRxByteRate returns the RxByteRate field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRxByteRate() int64 {
	if o == nil || o.RxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.RxByteRate
}

// GetRxByteRateOk returns a tuple with the RxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRxByteRateOk() (*int64, bool) {
	if o == nil || o.RxByteRate == nil {
		return nil, false
	}
	return o.RxByteRate, true
}

// HasRxByteRate returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRxByteRate() bool {
	if o != nil && o.RxByteRate != nil {
		return true
	}

	return false
}

// SetRxByteRate gets a reference to the given int64 and assigns it to the RxByteRate field.
func (o *MsgVpnClient) SetRxByteRate(v int64) {
	o.RxByteRate = &v
}

// GetRxDiscardedMsgCount returns the RxDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRxDiscardedMsgCount() int64 {
	if o == nil || o.RxDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RxDiscardedMsgCount
}

// GetRxDiscardedMsgCountOk returns a tuple with the RxDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRxDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.RxDiscardedMsgCount == nil {
		return nil, false
	}
	return o.RxDiscardedMsgCount, true
}

// HasRxDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRxDiscardedMsgCount() bool {
	if o != nil && o.RxDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetRxDiscardedMsgCount gets a reference to the given int64 and assigns it to the RxDiscardedMsgCount field.
func (o *MsgVpnClient) SetRxDiscardedMsgCount(v int64) {
	o.RxDiscardedMsgCount = &v
}

// GetRxMsgCount returns the RxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRxMsgCount() int64 {
	if o == nil || o.RxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RxMsgCount
}

// GetRxMsgCountOk returns a tuple with the RxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRxMsgCountOk() (*int64, bool) {
	if o == nil || o.RxMsgCount == nil {
		return nil, false
	}
	return o.RxMsgCount, true
}

// HasRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRxMsgCount() bool {
	if o != nil && o.RxMsgCount != nil {
		return true
	}

	return false
}

// SetRxMsgCount gets a reference to the given int64 and assigns it to the RxMsgCount field.
func (o *MsgVpnClient) SetRxMsgCount(v int64) {
	o.RxMsgCount = &v
}

// GetRxMsgRate returns the RxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnClient) GetRxMsgRate() int64 {
	if o == nil || o.RxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.RxMsgRate
}

// GetRxMsgRateOk returns a tuple with the RxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetRxMsgRateOk() (*int64, bool) {
	if o == nil || o.RxMsgRate == nil {
		return nil, false
	}
	return o.RxMsgRate, true
}

// HasRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnClient) HasRxMsgRate() bool {
	if o != nil && o.RxMsgRate != nil {
		return true
	}

	return false
}

// SetRxMsgRate gets a reference to the given int64 and assigns it to the RxMsgRate field.
func (o *MsgVpnClient) SetRxMsgRate(v int64) {
	o.RxMsgRate = &v
}

// GetScheduledDisconnectTime returns the ScheduledDisconnectTime field value if set, zero value otherwise.
func (o *MsgVpnClient) GetScheduledDisconnectTime() int32 {
	if o == nil || o.ScheduledDisconnectTime == nil {
		var ret int32
		return ret
	}
	return *o.ScheduledDisconnectTime
}

// GetScheduledDisconnectTimeOk returns a tuple with the ScheduledDisconnectTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetScheduledDisconnectTimeOk() (*int32, bool) {
	if o == nil || o.ScheduledDisconnectTime == nil {
		return nil, false
	}
	return o.ScheduledDisconnectTime, true
}

// HasScheduledDisconnectTime returns a boolean if a field has been set.
func (o *MsgVpnClient) HasScheduledDisconnectTime() bool {
	if o != nil && o.ScheduledDisconnectTime != nil {
		return true
	}

	return false
}

// SetScheduledDisconnectTime gets a reference to the given int32 and assigns it to the ScheduledDisconnectTime field.
func (o *MsgVpnClient) SetScheduledDisconnectTime(v int32) {
	o.ScheduledDisconnectTime = &v
}

// GetSlowSubscriber returns the SlowSubscriber field value if set, zero value otherwise.
func (o *MsgVpnClient) GetSlowSubscriber() bool {
	if o == nil || o.SlowSubscriber == nil {
		var ret bool
		return ret
	}
	return *o.SlowSubscriber
}

// GetSlowSubscriberOk returns a tuple with the SlowSubscriber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetSlowSubscriberOk() (*bool, bool) {
	if o == nil || o.SlowSubscriber == nil {
		return nil, false
	}
	return o.SlowSubscriber, true
}

// HasSlowSubscriber returns a boolean if a field has been set.
func (o *MsgVpnClient) HasSlowSubscriber() bool {
	if o != nil && o.SlowSubscriber != nil {
		return true
	}

	return false
}

// SetSlowSubscriber gets a reference to the given bool and assigns it to the SlowSubscriber field.
func (o *MsgVpnClient) SetSlowSubscriber(v bool) {
	o.SlowSubscriber = &v
}

// GetSoftwareDate returns the SoftwareDate field value if set, zero value otherwise.
func (o *MsgVpnClient) GetSoftwareDate() string {
	if o == nil || o.SoftwareDate == nil {
		var ret string
		return ret
	}
	return *o.SoftwareDate
}

// GetSoftwareDateOk returns a tuple with the SoftwareDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetSoftwareDateOk() (*string, bool) {
	if o == nil || o.SoftwareDate == nil {
		return nil, false
	}
	return o.SoftwareDate, true
}

// HasSoftwareDate returns a boolean if a field has been set.
func (o *MsgVpnClient) HasSoftwareDate() bool {
	if o != nil && o.SoftwareDate != nil {
		return true
	}

	return false
}

// SetSoftwareDate gets a reference to the given string and assigns it to the SoftwareDate field.
func (o *MsgVpnClient) SetSoftwareDate(v string) {
	o.SoftwareDate = &v
}

// GetSoftwareVersion returns the SoftwareVersion field value if set, zero value otherwise.
func (o *MsgVpnClient) GetSoftwareVersion() string {
	if o == nil || o.SoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.SoftwareVersion
}

// GetSoftwareVersionOk returns a tuple with the SoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetSoftwareVersionOk() (*string, bool) {
	if o == nil || o.SoftwareVersion == nil {
		return nil, false
	}
	return o.SoftwareVersion, true
}

// HasSoftwareVersion returns a boolean if a field has been set.
func (o *MsgVpnClient) HasSoftwareVersion() bool {
	if o != nil && o.SoftwareVersion != nil {
		return true
	}

	return false
}

// SetSoftwareVersion gets a reference to the given string and assigns it to the SoftwareVersion field.
func (o *MsgVpnClient) SetSoftwareVersion(v string) {
	o.SoftwareVersion = &v
}

// GetTlsCipherDescription returns the TlsCipherDescription field value if set, zero value otherwise.
func (o *MsgVpnClient) GetTlsCipherDescription() string {
	if o == nil || o.TlsCipherDescription == nil {
		var ret string
		return ret
	}
	return *o.TlsCipherDescription
}

// GetTlsCipherDescriptionOk returns a tuple with the TlsCipherDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetTlsCipherDescriptionOk() (*string, bool) {
	if o == nil || o.TlsCipherDescription == nil {
		return nil, false
	}
	return o.TlsCipherDescription, true
}

// HasTlsCipherDescription returns a boolean if a field has been set.
func (o *MsgVpnClient) HasTlsCipherDescription() bool {
	if o != nil && o.TlsCipherDescription != nil {
		return true
	}

	return false
}

// SetTlsCipherDescription gets a reference to the given string and assigns it to the TlsCipherDescription field.
func (o *MsgVpnClient) SetTlsCipherDescription(v string) {
	o.TlsCipherDescription = &v
}

// GetTlsDowngradedToPlainText returns the TlsDowngradedToPlainText field value if set, zero value otherwise.
func (o *MsgVpnClient) GetTlsDowngradedToPlainText() bool {
	if o == nil || o.TlsDowngradedToPlainText == nil {
		var ret bool
		return ret
	}
	return *o.TlsDowngradedToPlainText
}

// GetTlsDowngradedToPlainTextOk returns a tuple with the TlsDowngradedToPlainText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetTlsDowngradedToPlainTextOk() (*bool, bool) {
	if o == nil || o.TlsDowngradedToPlainText == nil {
		return nil, false
	}
	return o.TlsDowngradedToPlainText, true
}

// HasTlsDowngradedToPlainText returns a boolean if a field has been set.
func (o *MsgVpnClient) HasTlsDowngradedToPlainText() bool {
	if o != nil && o.TlsDowngradedToPlainText != nil {
		return true
	}

	return false
}

// SetTlsDowngradedToPlainText gets a reference to the given bool and assigns it to the TlsDowngradedToPlainText field.
func (o *MsgVpnClient) SetTlsDowngradedToPlainText(v bool) {
	o.TlsDowngradedToPlainText = &v
}

// GetTlsVersion returns the TlsVersion field value if set, zero value otherwise.
func (o *MsgVpnClient) GetTlsVersion() string {
	if o == nil || o.TlsVersion == nil {
		var ret string
		return ret
	}
	return *o.TlsVersion
}

// GetTlsVersionOk returns a tuple with the TlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetTlsVersionOk() (*string, bool) {
	if o == nil || o.TlsVersion == nil {
		return nil, false
	}
	return o.TlsVersion, true
}

// HasTlsVersion returns a boolean if a field has been set.
func (o *MsgVpnClient) HasTlsVersion() bool {
	if o != nil && o.TlsVersion != nil {
		return true
	}

	return false
}

// SetTlsVersion gets a reference to the given string and assigns it to the TlsVersion field.
func (o *MsgVpnClient) SetTlsVersion(v string) {
	o.TlsVersion = &v
}

// GetTopicParseErrorRxDiscardedMsgCount returns the TopicParseErrorRxDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetTopicParseErrorRxDiscardedMsgCount() int64 {
	if o == nil || o.TopicParseErrorRxDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.TopicParseErrorRxDiscardedMsgCount
}

// GetTopicParseErrorRxDiscardedMsgCountOk returns a tuple with the TopicParseErrorRxDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetTopicParseErrorRxDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.TopicParseErrorRxDiscardedMsgCount == nil {
		return nil, false
	}
	return o.TopicParseErrorRxDiscardedMsgCount, true
}

// HasTopicParseErrorRxDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasTopicParseErrorRxDiscardedMsgCount() bool {
	if o != nil && o.TopicParseErrorRxDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetTopicParseErrorRxDiscardedMsgCount gets a reference to the given int64 and assigns it to the TopicParseErrorRxDiscardedMsgCount field.
func (o *MsgVpnClient) SetTopicParseErrorRxDiscardedMsgCount(v int64) {
	o.TopicParseErrorRxDiscardedMsgCount = &v
}

// GetTxByteCount returns the TxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetTxByteCount() int64 {
	if o == nil || o.TxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.TxByteCount
}

// GetTxByteCountOk returns a tuple with the TxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetTxByteCountOk() (*int64, bool) {
	if o == nil || o.TxByteCount == nil {
		return nil, false
	}
	return o.TxByteCount, true
}

// HasTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasTxByteCount() bool {
	if o != nil && o.TxByteCount != nil {
		return true
	}

	return false
}

// SetTxByteCount gets a reference to the given int64 and assigns it to the TxByteCount field.
func (o *MsgVpnClient) SetTxByteCount(v int64) {
	o.TxByteCount = &v
}

// GetTxByteRate returns the TxByteRate field value if set, zero value otherwise.
func (o *MsgVpnClient) GetTxByteRate() int64 {
	if o == nil || o.TxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.TxByteRate
}

// GetTxByteRateOk returns a tuple with the TxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetTxByteRateOk() (*int64, bool) {
	if o == nil || o.TxByteRate == nil {
		return nil, false
	}
	return o.TxByteRate, true
}

// HasTxByteRate returns a boolean if a field has been set.
func (o *MsgVpnClient) HasTxByteRate() bool {
	if o != nil && o.TxByteRate != nil {
		return true
	}

	return false
}

// SetTxByteRate gets a reference to the given int64 and assigns it to the TxByteRate field.
func (o *MsgVpnClient) SetTxByteRate(v int64) {
	o.TxByteRate = &v
}

// GetTxDiscardedMsgCount returns the TxDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetTxDiscardedMsgCount() int64 {
	if o == nil || o.TxDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.TxDiscardedMsgCount
}

// GetTxDiscardedMsgCountOk returns a tuple with the TxDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetTxDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.TxDiscardedMsgCount == nil {
		return nil, false
	}
	return o.TxDiscardedMsgCount, true
}

// HasTxDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasTxDiscardedMsgCount() bool {
	if o != nil && o.TxDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetTxDiscardedMsgCount gets a reference to the given int64 and assigns it to the TxDiscardedMsgCount field.
func (o *MsgVpnClient) SetTxDiscardedMsgCount(v int64) {
	o.TxDiscardedMsgCount = &v
}

// GetTxMsgCount returns the TxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetTxMsgCount() int64 {
	if o == nil || o.TxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.TxMsgCount
}

// GetTxMsgCountOk returns a tuple with the TxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetTxMsgCountOk() (*int64, bool) {
	if o == nil || o.TxMsgCount == nil {
		return nil, false
	}
	return o.TxMsgCount, true
}

// HasTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasTxMsgCount() bool {
	if o != nil && o.TxMsgCount != nil {
		return true
	}

	return false
}

// SetTxMsgCount gets a reference to the given int64 and assigns it to the TxMsgCount field.
func (o *MsgVpnClient) SetTxMsgCount(v int64) {
	o.TxMsgCount = &v
}

// GetTxMsgRate returns the TxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnClient) GetTxMsgRate() int64 {
	if o == nil || o.TxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.TxMsgRate
}

// GetTxMsgRateOk returns a tuple with the TxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetTxMsgRateOk() (*int64, bool) {
	if o == nil || o.TxMsgRate == nil {
		return nil, false
	}
	return o.TxMsgRate, true
}

// HasTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnClient) HasTxMsgRate() bool {
	if o != nil && o.TxMsgRate != nil {
		return true
	}

	return false
}

// SetTxMsgRate gets a reference to the given int64 and assigns it to the TxMsgRate field.
func (o *MsgVpnClient) SetTxMsgRate(v int64) {
	o.TxMsgRate = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *MsgVpnClient) GetUptime() int32 {
	if o == nil || o.Uptime == nil {
		var ret int32
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetUptimeOk() (*int32, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *MsgVpnClient) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given int32 and assigns it to the Uptime field.
func (o *MsgVpnClient) SetUptime(v int32) {
	o.Uptime = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MsgVpnClient) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MsgVpnClient) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *MsgVpnClient) SetUser(v string) {
	o.User = &v
}

// GetVirtualRouter returns the VirtualRouter field value if set, zero value otherwise.
func (o *MsgVpnClient) GetVirtualRouter() string {
	if o == nil || o.VirtualRouter == nil {
		var ret string
		return ret
	}
	return *o.VirtualRouter
}

// GetVirtualRouterOk returns a tuple with the VirtualRouter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetVirtualRouterOk() (*string, bool) {
	if o == nil || o.VirtualRouter == nil {
		return nil, false
	}
	return o.VirtualRouter, true
}

// HasVirtualRouter returns a boolean if a field has been set.
func (o *MsgVpnClient) HasVirtualRouter() bool {
	if o != nil && o.VirtualRouter != nil {
		return true
	}

	return false
}

// SetVirtualRouter gets a reference to the given string and assigns it to the VirtualRouter field.
func (o *MsgVpnClient) SetVirtualRouter(v string) {
	o.VirtualRouter = &v
}

// GetWebInactiveTimeout returns the WebInactiveTimeout field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebInactiveTimeout() int32 {
	if o == nil || o.WebInactiveTimeout == nil {
		var ret int32
		return ret
	}
	return *o.WebInactiveTimeout
}

// GetWebInactiveTimeoutOk returns a tuple with the WebInactiveTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebInactiveTimeoutOk() (*int32, bool) {
	if o == nil || o.WebInactiveTimeout == nil {
		return nil, false
	}
	return o.WebInactiveTimeout, true
}

// HasWebInactiveTimeout returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebInactiveTimeout() bool {
	if o != nil && o.WebInactiveTimeout != nil {
		return true
	}

	return false
}

// SetWebInactiveTimeout gets a reference to the given int32 and assigns it to the WebInactiveTimeout field.
func (o *MsgVpnClient) SetWebInactiveTimeout(v int32) {
	o.WebInactiveTimeout = &v
}

// GetWebMaxPayload returns the WebMaxPayload field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebMaxPayload() int64 {
	if o == nil || o.WebMaxPayload == nil {
		var ret int64
		return ret
	}
	return *o.WebMaxPayload
}

// GetWebMaxPayloadOk returns a tuple with the WebMaxPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebMaxPayloadOk() (*int64, bool) {
	if o == nil || o.WebMaxPayload == nil {
		return nil, false
	}
	return o.WebMaxPayload, true
}

// HasWebMaxPayload returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebMaxPayload() bool {
	if o != nil && o.WebMaxPayload != nil {
		return true
	}

	return false
}

// SetWebMaxPayload gets a reference to the given int64 and assigns it to the WebMaxPayload field.
func (o *MsgVpnClient) SetWebMaxPayload(v int64) {
	o.WebMaxPayload = &v
}

// GetWebParseErrorRxDiscardedMsgCount returns the WebParseErrorRxDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebParseErrorRxDiscardedMsgCount() int64 {
	if o == nil || o.WebParseErrorRxDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.WebParseErrorRxDiscardedMsgCount
}

// GetWebParseErrorRxDiscardedMsgCountOk returns a tuple with the WebParseErrorRxDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebParseErrorRxDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.WebParseErrorRxDiscardedMsgCount == nil {
		return nil, false
	}
	return o.WebParseErrorRxDiscardedMsgCount, true
}

// HasWebParseErrorRxDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebParseErrorRxDiscardedMsgCount() bool {
	if o != nil && o.WebParseErrorRxDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetWebParseErrorRxDiscardedMsgCount gets a reference to the given int64 and assigns it to the WebParseErrorRxDiscardedMsgCount field.
func (o *MsgVpnClient) SetWebParseErrorRxDiscardedMsgCount(v int64) {
	o.WebParseErrorRxDiscardedMsgCount = &v
}

// GetWebRemainingTimeout returns the WebRemainingTimeout field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebRemainingTimeout() int32 {
	if o == nil || o.WebRemainingTimeout == nil {
		var ret int32
		return ret
	}
	return *o.WebRemainingTimeout
}

// GetWebRemainingTimeoutOk returns a tuple with the WebRemainingTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebRemainingTimeoutOk() (*int32, bool) {
	if o == nil || o.WebRemainingTimeout == nil {
		return nil, false
	}
	return o.WebRemainingTimeout, true
}

// HasWebRemainingTimeout returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebRemainingTimeout() bool {
	if o != nil && o.WebRemainingTimeout != nil {
		return true
	}

	return false
}

// SetWebRemainingTimeout gets a reference to the given int32 and assigns it to the WebRemainingTimeout field.
func (o *MsgVpnClient) SetWebRemainingTimeout(v int32) {
	o.WebRemainingTimeout = &v
}

// GetWebRxByteCount returns the WebRxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebRxByteCount() int64 {
	if o == nil || o.WebRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.WebRxByteCount
}

// GetWebRxByteCountOk returns a tuple with the WebRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebRxByteCountOk() (*int64, bool) {
	if o == nil || o.WebRxByteCount == nil {
		return nil, false
	}
	return o.WebRxByteCount, true
}

// HasWebRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebRxByteCount() bool {
	if o != nil && o.WebRxByteCount != nil {
		return true
	}

	return false
}

// SetWebRxByteCount gets a reference to the given int64 and assigns it to the WebRxByteCount field.
func (o *MsgVpnClient) SetWebRxByteCount(v int64) {
	o.WebRxByteCount = &v
}

// GetWebRxEncoding returns the WebRxEncoding field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebRxEncoding() string {
	if o == nil || o.WebRxEncoding == nil {
		var ret string
		return ret
	}
	return *o.WebRxEncoding
}

// GetWebRxEncodingOk returns a tuple with the WebRxEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebRxEncodingOk() (*string, bool) {
	if o == nil || o.WebRxEncoding == nil {
		return nil, false
	}
	return o.WebRxEncoding, true
}

// HasWebRxEncoding returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebRxEncoding() bool {
	if o != nil && o.WebRxEncoding != nil {
		return true
	}

	return false
}

// SetWebRxEncoding gets a reference to the given string and assigns it to the WebRxEncoding field.
func (o *MsgVpnClient) SetWebRxEncoding(v string) {
	o.WebRxEncoding = &v
}

// GetWebRxMsgCount returns the WebRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebRxMsgCount() int64 {
	if o == nil || o.WebRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.WebRxMsgCount
}

// GetWebRxMsgCountOk returns a tuple with the WebRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebRxMsgCountOk() (*int64, bool) {
	if o == nil || o.WebRxMsgCount == nil {
		return nil, false
	}
	return o.WebRxMsgCount, true
}

// HasWebRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebRxMsgCount() bool {
	if o != nil && o.WebRxMsgCount != nil {
		return true
	}

	return false
}

// SetWebRxMsgCount gets a reference to the given int64 and assigns it to the WebRxMsgCount field.
func (o *MsgVpnClient) SetWebRxMsgCount(v int64) {
	o.WebRxMsgCount = &v
}

// GetWebRxProtocol returns the WebRxProtocol field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebRxProtocol() string {
	if o == nil || o.WebRxProtocol == nil {
		var ret string
		return ret
	}
	return *o.WebRxProtocol
}

// GetWebRxProtocolOk returns a tuple with the WebRxProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebRxProtocolOk() (*string, bool) {
	if o == nil || o.WebRxProtocol == nil {
		return nil, false
	}
	return o.WebRxProtocol, true
}

// HasWebRxProtocol returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebRxProtocol() bool {
	if o != nil && o.WebRxProtocol != nil {
		return true
	}

	return false
}

// SetWebRxProtocol gets a reference to the given string and assigns it to the WebRxProtocol field.
func (o *MsgVpnClient) SetWebRxProtocol(v string) {
	o.WebRxProtocol = &v
}

// GetWebRxRequestCount returns the WebRxRequestCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebRxRequestCount() int64 {
	if o == nil || o.WebRxRequestCount == nil {
		var ret int64
		return ret
	}
	return *o.WebRxRequestCount
}

// GetWebRxRequestCountOk returns a tuple with the WebRxRequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebRxRequestCountOk() (*int64, bool) {
	if o == nil || o.WebRxRequestCount == nil {
		return nil, false
	}
	return o.WebRxRequestCount, true
}

// HasWebRxRequestCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebRxRequestCount() bool {
	if o != nil && o.WebRxRequestCount != nil {
		return true
	}

	return false
}

// SetWebRxRequestCount gets a reference to the given int64 and assigns it to the WebRxRequestCount field.
func (o *MsgVpnClient) SetWebRxRequestCount(v int64) {
	o.WebRxRequestCount = &v
}

// GetWebRxResponseCount returns the WebRxResponseCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebRxResponseCount() int64 {
	if o == nil || o.WebRxResponseCount == nil {
		var ret int64
		return ret
	}
	return *o.WebRxResponseCount
}

// GetWebRxResponseCountOk returns a tuple with the WebRxResponseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebRxResponseCountOk() (*int64, bool) {
	if o == nil || o.WebRxResponseCount == nil {
		return nil, false
	}
	return o.WebRxResponseCount, true
}

// HasWebRxResponseCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebRxResponseCount() bool {
	if o != nil && o.WebRxResponseCount != nil {
		return true
	}

	return false
}

// SetWebRxResponseCount gets a reference to the given int64 and assigns it to the WebRxResponseCount field.
func (o *MsgVpnClient) SetWebRxResponseCount(v int64) {
	o.WebRxResponseCount = &v
}

// GetWebRxTcpState returns the WebRxTcpState field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebRxTcpState() string {
	if o == nil || o.WebRxTcpState == nil {
		var ret string
		return ret
	}
	return *o.WebRxTcpState
}

// GetWebRxTcpStateOk returns a tuple with the WebRxTcpState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebRxTcpStateOk() (*string, bool) {
	if o == nil || o.WebRxTcpState == nil {
		return nil, false
	}
	return o.WebRxTcpState, true
}

// HasWebRxTcpState returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebRxTcpState() bool {
	if o != nil && o.WebRxTcpState != nil {
		return true
	}

	return false
}

// SetWebRxTcpState gets a reference to the given string and assigns it to the WebRxTcpState field.
func (o *MsgVpnClient) SetWebRxTcpState(v string) {
	o.WebRxTcpState = &v
}

// GetWebRxTlsCipherDescription returns the WebRxTlsCipherDescription field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebRxTlsCipherDescription() string {
	if o == nil || o.WebRxTlsCipherDescription == nil {
		var ret string
		return ret
	}
	return *o.WebRxTlsCipherDescription
}

// GetWebRxTlsCipherDescriptionOk returns a tuple with the WebRxTlsCipherDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebRxTlsCipherDescriptionOk() (*string, bool) {
	if o == nil || o.WebRxTlsCipherDescription == nil {
		return nil, false
	}
	return o.WebRxTlsCipherDescription, true
}

// HasWebRxTlsCipherDescription returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebRxTlsCipherDescription() bool {
	if o != nil && o.WebRxTlsCipherDescription != nil {
		return true
	}

	return false
}

// SetWebRxTlsCipherDescription gets a reference to the given string and assigns it to the WebRxTlsCipherDescription field.
func (o *MsgVpnClient) SetWebRxTlsCipherDescription(v string) {
	o.WebRxTlsCipherDescription = &v
}

// GetWebRxTlsVersion returns the WebRxTlsVersion field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebRxTlsVersion() string {
	if o == nil || o.WebRxTlsVersion == nil {
		var ret string
		return ret
	}
	return *o.WebRxTlsVersion
}

// GetWebRxTlsVersionOk returns a tuple with the WebRxTlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebRxTlsVersionOk() (*string, bool) {
	if o == nil || o.WebRxTlsVersion == nil {
		return nil, false
	}
	return o.WebRxTlsVersion, true
}

// HasWebRxTlsVersion returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebRxTlsVersion() bool {
	if o != nil && o.WebRxTlsVersion != nil {
		return true
	}

	return false
}

// SetWebRxTlsVersion gets a reference to the given string and assigns it to the WebRxTlsVersion field.
func (o *MsgVpnClient) SetWebRxTlsVersion(v string) {
	o.WebRxTlsVersion = &v
}

// GetWebSessionId returns the WebSessionId field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebSessionId() string {
	if o == nil || o.WebSessionId == nil {
		var ret string
		return ret
	}
	return *o.WebSessionId
}

// GetWebSessionIdOk returns a tuple with the WebSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebSessionIdOk() (*string, bool) {
	if o == nil || o.WebSessionId == nil {
		return nil, false
	}
	return o.WebSessionId, true
}

// HasWebSessionId returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebSessionId() bool {
	if o != nil && o.WebSessionId != nil {
		return true
	}

	return false
}

// SetWebSessionId gets a reference to the given string and assigns it to the WebSessionId field.
func (o *MsgVpnClient) SetWebSessionId(v string) {
	o.WebSessionId = &v
}

// GetWebTxByteCount returns the WebTxByteCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebTxByteCount() int64 {
	if o == nil || o.WebTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.WebTxByteCount
}

// GetWebTxByteCountOk returns a tuple with the WebTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebTxByteCountOk() (*int64, bool) {
	if o == nil || o.WebTxByteCount == nil {
		return nil, false
	}
	return o.WebTxByteCount, true
}

// HasWebTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebTxByteCount() bool {
	if o != nil && o.WebTxByteCount != nil {
		return true
	}

	return false
}

// SetWebTxByteCount gets a reference to the given int64 and assigns it to the WebTxByteCount field.
func (o *MsgVpnClient) SetWebTxByteCount(v int64) {
	o.WebTxByteCount = &v
}

// GetWebTxEncoding returns the WebTxEncoding field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebTxEncoding() string {
	if o == nil || o.WebTxEncoding == nil {
		var ret string
		return ret
	}
	return *o.WebTxEncoding
}

// GetWebTxEncodingOk returns a tuple with the WebTxEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebTxEncodingOk() (*string, bool) {
	if o == nil || o.WebTxEncoding == nil {
		return nil, false
	}
	return o.WebTxEncoding, true
}

// HasWebTxEncoding returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebTxEncoding() bool {
	if o != nil && o.WebTxEncoding != nil {
		return true
	}

	return false
}

// SetWebTxEncoding gets a reference to the given string and assigns it to the WebTxEncoding field.
func (o *MsgVpnClient) SetWebTxEncoding(v string) {
	o.WebTxEncoding = &v
}

// GetWebTxMsgCount returns the WebTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebTxMsgCount() int64 {
	if o == nil || o.WebTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.WebTxMsgCount
}

// GetWebTxMsgCountOk returns a tuple with the WebTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebTxMsgCountOk() (*int64, bool) {
	if o == nil || o.WebTxMsgCount == nil {
		return nil, false
	}
	return o.WebTxMsgCount, true
}

// HasWebTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebTxMsgCount() bool {
	if o != nil && o.WebTxMsgCount != nil {
		return true
	}

	return false
}

// SetWebTxMsgCount gets a reference to the given int64 and assigns it to the WebTxMsgCount field.
func (o *MsgVpnClient) SetWebTxMsgCount(v int64) {
	o.WebTxMsgCount = &v
}

// GetWebTxProtocol returns the WebTxProtocol field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebTxProtocol() string {
	if o == nil || o.WebTxProtocol == nil {
		var ret string
		return ret
	}
	return *o.WebTxProtocol
}

// GetWebTxProtocolOk returns a tuple with the WebTxProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebTxProtocolOk() (*string, bool) {
	if o == nil || o.WebTxProtocol == nil {
		return nil, false
	}
	return o.WebTxProtocol, true
}

// HasWebTxProtocol returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebTxProtocol() bool {
	if o != nil && o.WebTxProtocol != nil {
		return true
	}

	return false
}

// SetWebTxProtocol gets a reference to the given string and assigns it to the WebTxProtocol field.
func (o *MsgVpnClient) SetWebTxProtocol(v string) {
	o.WebTxProtocol = &v
}

// GetWebTxRequestCount returns the WebTxRequestCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebTxRequestCount() int64 {
	if o == nil || o.WebTxRequestCount == nil {
		var ret int64
		return ret
	}
	return *o.WebTxRequestCount
}

// GetWebTxRequestCountOk returns a tuple with the WebTxRequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebTxRequestCountOk() (*int64, bool) {
	if o == nil || o.WebTxRequestCount == nil {
		return nil, false
	}
	return o.WebTxRequestCount, true
}

// HasWebTxRequestCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebTxRequestCount() bool {
	if o != nil && o.WebTxRequestCount != nil {
		return true
	}

	return false
}

// SetWebTxRequestCount gets a reference to the given int64 and assigns it to the WebTxRequestCount field.
func (o *MsgVpnClient) SetWebTxRequestCount(v int64) {
	o.WebTxRequestCount = &v
}

// GetWebTxResponseCount returns the WebTxResponseCount field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebTxResponseCount() int64 {
	if o == nil || o.WebTxResponseCount == nil {
		var ret int64
		return ret
	}
	return *o.WebTxResponseCount
}

// GetWebTxResponseCountOk returns a tuple with the WebTxResponseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebTxResponseCountOk() (*int64, bool) {
	if o == nil || o.WebTxResponseCount == nil {
		return nil, false
	}
	return o.WebTxResponseCount, true
}

// HasWebTxResponseCount returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebTxResponseCount() bool {
	if o != nil && o.WebTxResponseCount != nil {
		return true
	}

	return false
}

// SetWebTxResponseCount gets a reference to the given int64 and assigns it to the WebTxResponseCount field.
func (o *MsgVpnClient) SetWebTxResponseCount(v int64) {
	o.WebTxResponseCount = &v
}

// GetWebTxTcpState returns the WebTxTcpState field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebTxTcpState() string {
	if o == nil || o.WebTxTcpState == nil {
		var ret string
		return ret
	}
	return *o.WebTxTcpState
}

// GetWebTxTcpStateOk returns a tuple with the WebTxTcpState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebTxTcpStateOk() (*string, bool) {
	if o == nil || o.WebTxTcpState == nil {
		return nil, false
	}
	return o.WebTxTcpState, true
}

// HasWebTxTcpState returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebTxTcpState() bool {
	if o != nil && o.WebTxTcpState != nil {
		return true
	}

	return false
}

// SetWebTxTcpState gets a reference to the given string and assigns it to the WebTxTcpState field.
func (o *MsgVpnClient) SetWebTxTcpState(v string) {
	o.WebTxTcpState = &v
}

// GetWebTxTlsCipherDescription returns the WebTxTlsCipherDescription field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebTxTlsCipherDescription() string {
	if o == nil || o.WebTxTlsCipherDescription == nil {
		var ret string
		return ret
	}
	return *o.WebTxTlsCipherDescription
}

// GetWebTxTlsCipherDescriptionOk returns a tuple with the WebTxTlsCipherDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebTxTlsCipherDescriptionOk() (*string, bool) {
	if o == nil || o.WebTxTlsCipherDescription == nil {
		return nil, false
	}
	return o.WebTxTlsCipherDescription, true
}

// HasWebTxTlsCipherDescription returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebTxTlsCipherDescription() bool {
	if o != nil && o.WebTxTlsCipherDescription != nil {
		return true
	}

	return false
}

// SetWebTxTlsCipherDescription gets a reference to the given string and assigns it to the WebTxTlsCipherDescription field.
func (o *MsgVpnClient) SetWebTxTlsCipherDescription(v string) {
	o.WebTxTlsCipherDescription = &v
}

// GetWebTxTlsVersion returns the WebTxTlsVersion field value if set, zero value otherwise.
func (o *MsgVpnClient) GetWebTxTlsVersion() string {
	if o == nil || o.WebTxTlsVersion == nil {
		var ret string
		return ret
	}
	return *o.WebTxTlsVersion
}

// GetWebTxTlsVersionOk returns a tuple with the WebTxTlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClient) GetWebTxTlsVersionOk() (*string, bool) {
	if o == nil || o.WebTxTlsVersion == nil {
		return nil, false
	}
	return o.WebTxTlsVersion, true
}

// HasWebTxTlsVersion returns a boolean if a field has been set.
func (o *MsgVpnClient) HasWebTxTlsVersion() bool {
	if o != nil && o.WebTxTlsVersion != nil {
		return true
	}

	return false
}

// SetWebTxTlsVersion gets a reference to the given string and assigns it to the WebTxTlsVersion field.
func (o *MsgVpnClient) SetWebTxTlsVersion(v string) {
	o.WebTxTlsVersion = &v
}

func (o MsgVpnClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AclProfileName != nil {
		toSerialize["aclProfileName"] = o.AclProfileName
	}
	if o.AliasedFromMsgVpnName != nil {
		toSerialize["aliasedFromMsgVpnName"] = o.AliasedFromMsgVpnName
	}
	if o.AlreadyBoundBindFailureCount != nil {
		toSerialize["alreadyBoundBindFailureCount"] = o.AlreadyBoundBindFailureCount
	}
	if o.AuthorizationGroupName != nil {
		toSerialize["authorizationGroupName"] = o.AuthorizationGroupName
	}
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
	if o.BindRequestCount != nil {
		toSerialize["bindRequestCount"] = o.BindRequestCount
	}
	if o.BindSuccessCount != nil {
		toSerialize["bindSuccessCount"] = o.BindSuccessCount
	}
	if o.ClientAddress != nil {
		toSerialize["clientAddress"] = o.ClientAddress
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.ClientProfileName != nil {
		toSerialize["clientProfileName"] = o.ClientProfileName
	}
	if o.ClientUsername != nil {
		toSerialize["clientUsername"] = o.ClientUsername
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
	if o.CutThroughDeniedBindFailureCount != nil {
		toSerialize["cutThroughDeniedBindFailureCount"] = o.CutThroughDeniedBindFailureCount
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisabledBindFailureCount != nil {
		toSerialize["disabledBindFailureCount"] = o.DisabledBindFailureCount
	}
	if o.DtoLocalPriority != nil {
		toSerialize["dtoLocalPriority"] = o.DtoLocalPriority
	}
	if o.DtoNetworkPriority != nil {
		toSerialize["dtoNetworkPriority"] = o.DtoNetworkPriority
	}
	if o.Eliding != nil {
		toSerialize["eliding"] = o.Eliding
	}
	if o.ElidingTopicCount != nil {
		toSerialize["elidingTopicCount"] = o.ElidingTopicCount
	}
	if o.ElidingTopicPeakCount != nil {
		toSerialize["elidingTopicPeakCount"] = o.ElidingTopicPeakCount
	}
	if o.GuaranteedDeniedBindFailureCount != nil {
		toSerialize["guaranteedDeniedBindFailureCount"] = o.GuaranteedDeniedBindFailureCount
	}
	if o.InvalidSelectorBindFailureCount != nil {
		toSerialize["invalidSelectorBindFailureCount"] = o.InvalidSelectorBindFailureCount
	}
	if o.Keepalive != nil {
		toSerialize["keepalive"] = o.Keepalive
	}
	if o.KeepaliveEffectiveTimeout != nil {
		toSerialize["keepaliveEffectiveTimeout"] = o.KeepaliveEffectiveTimeout
	}
	if o.LargeMsgEventRaised != nil {
		toSerialize["largeMsgEventRaised"] = o.LargeMsgEventRaised
	}
	if o.LoginRxMsgCount != nil {
		toSerialize["loginRxMsgCount"] = o.LoginRxMsgCount
	}
	if o.LoginTxMsgCount != nil {
		toSerialize["loginTxMsgCount"] = o.LoginTxMsgCount
	}
	if o.MaxBindCountExceededBindFailureCount != nil {
		toSerialize["maxBindCountExceededBindFailureCount"] = o.MaxBindCountExceededBindFailureCount
	}
	if o.MaxElidingTopicCountEventRaised != nil {
		toSerialize["maxElidingTopicCountEventRaised"] = o.MaxElidingTopicCountEventRaised
	}
	if o.MqttConnackErrorTxCount != nil {
		toSerialize["mqttConnackErrorTxCount"] = o.MqttConnackErrorTxCount
	}
	if o.MqttConnackTxCount != nil {
		toSerialize["mqttConnackTxCount"] = o.MqttConnackTxCount
	}
	if o.MqttConnectRxCount != nil {
		toSerialize["mqttConnectRxCount"] = o.MqttConnectRxCount
	}
	if o.MqttDisconnectRxCount != nil {
		toSerialize["mqttDisconnectRxCount"] = o.MqttDisconnectRxCount
	}
	if o.MqttPingreqRxCount != nil {
		toSerialize["mqttPingreqRxCount"] = o.MqttPingreqRxCount
	}
	if o.MqttPingrespTxCount != nil {
		toSerialize["mqttPingrespTxCount"] = o.MqttPingrespTxCount
	}
	if o.MqttPubackRxCount != nil {
		toSerialize["mqttPubackRxCount"] = o.MqttPubackRxCount
	}
	if o.MqttPubackTxCount != nil {
		toSerialize["mqttPubackTxCount"] = o.MqttPubackTxCount
	}
	if o.MqttPubcompTxCount != nil {
		toSerialize["mqttPubcompTxCount"] = o.MqttPubcompTxCount
	}
	if o.MqttPublishQos0RxCount != nil {
		toSerialize["mqttPublishQos0RxCount"] = o.MqttPublishQos0RxCount
	}
	if o.MqttPublishQos0TxCount != nil {
		toSerialize["mqttPublishQos0TxCount"] = o.MqttPublishQos0TxCount
	}
	if o.MqttPublishQos1RxCount != nil {
		toSerialize["mqttPublishQos1RxCount"] = o.MqttPublishQos1RxCount
	}
	if o.MqttPublishQos1TxCount != nil {
		toSerialize["mqttPublishQos1TxCount"] = o.MqttPublishQos1TxCount
	}
	if o.MqttPublishQos2RxCount != nil {
		toSerialize["mqttPublishQos2RxCount"] = o.MqttPublishQos2RxCount
	}
	if o.MqttPubrecTxCount != nil {
		toSerialize["mqttPubrecTxCount"] = o.MqttPubrecTxCount
	}
	if o.MqttPubrelRxCount != nil {
		toSerialize["mqttPubrelRxCount"] = o.MqttPubrelRxCount
	}
	if o.MqttSubackErrorTxCount != nil {
		toSerialize["mqttSubackErrorTxCount"] = o.MqttSubackErrorTxCount
	}
	if o.MqttSubackTxCount != nil {
		toSerialize["mqttSubackTxCount"] = o.MqttSubackTxCount
	}
	if o.MqttSubscribeRxCount != nil {
		toSerialize["mqttSubscribeRxCount"] = o.MqttSubscribeRxCount
	}
	if o.MqttUnsubackTxCount != nil {
		toSerialize["mqttUnsubackTxCount"] = o.MqttUnsubackTxCount
	}
	if o.MqttUnsubscribeRxCount != nil {
		toSerialize["mqttUnsubscribeRxCount"] = o.MqttUnsubscribeRxCount
	}
	if o.MsgSpoolCongestionRxDiscardedMsgCount != nil {
		toSerialize["msgSpoolCongestionRxDiscardedMsgCount"] = o.MsgSpoolCongestionRxDiscardedMsgCount
	}
	if o.MsgSpoolRxDiscardedMsgCount != nil {
		toSerialize["msgSpoolRxDiscardedMsgCount"] = o.MsgSpoolRxDiscardedMsgCount
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.NoLocalDelivery != nil {
		toSerialize["noLocalDelivery"] = o.NoLocalDelivery
	}
	if o.NoSubscriptionMatchRxDiscardedMsgCount != nil {
		toSerialize["noSubscriptionMatchRxDiscardedMsgCount"] = o.NoSubscriptionMatchRxDiscardedMsgCount
	}
	if o.OriginalClientUsername != nil {
		toSerialize["originalClientUsername"] = o.OriginalClientUsername
	}
	if o.OtherBindFailureCount != nil {
		toSerialize["otherBindFailureCount"] = o.OtherBindFailureCount
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.PublishTopicAclRxDiscardedMsgCount != nil {
		toSerialize["publishTopicAclRxDiscardedMsgCount"] = o.PublishTopicAclRxDiscardedMsgCount
	}
	if o.RestHttpRequestRxByteCount != nil {
		toSerialize["restHttpRequestRxByteCount"] = o.RestHttpRequestRxByteCount
	}
	if o.RestHttpRequestRxMsgCount != nil {
		toSerialize["restHttpRequestRxMsgCount"] = o.RestHttpRequestRxMsgCount
	}
	if o.RestHttpRequestTxByteCount != nil {
		toSerialize["restHttpRequestTxByteCount"] = o.RestHttpRequestTxByteCount
	}
	if o.RestHttpRequestTxMsgCount != nil {
		toSerialize["restHttpRequestTxMsgCount"] = o.RestHttpRequestTxMsgCount
	}
	if o.RestHttpResponseErrorRxMsgCount != nil {
		toSerialize["restHttpResponseErrorRxMsgCount"] = o.RestHttpResponseErrorRxMsgCount
	}
	if o.RestHttpResponseErrorTxMsgCount != nil {
		toSerialize["restHttpResponseErrorTxMsgCount"] = o.RestHttpResponseErrorTxMsgCount
	}
	if o.RestHttpResponseRxByteCount != nil {
		toSerialize["restHttpResponseRxByteCount"] = o.RestHttpResponseRxByteCount
	}
	if o.RestHttpResponseRxMsgCount != nil {
		toSerialize["restHttpResponseRxMsgCount"] = o.RestHttpResponseRxMsgCount
	}
	if o.RestHttpResponseSuccessRxMsgCount != nil {
		toSerialize["restHttpResponseSuccessRxMsgCount"] = o.RestHttpResponseSuccessRxMsgCount
	}
	if o.RestHttpResponseSuccessTxMsgCount != nil {
		toSerialize["restHttpResponseSuccessTxMsgCount"] = o.RestHttpResponseSuccessTxMsgCount
	}
	if o.RestHttpResponseTimeoutRxMsgCount != nil {
		toSerialize["restHttpResponseTimeoutRxMsgCount"] = o.RestHttpResponseTimeoutRxMsgCount
	}
	if o.RestHttpResponseTimeoutTxMsgCount != nil {
		toSerialize["restHttpResponseTimeoutTxMsgCount"] = o.RestHttpResponseTimeoutTxMsgCount
	}
	if o.RestHttpResponseTxByteCount != nil {
		toSerialize["restHttpResponseTxByteCount"] = o.RestHttpResponseTxByteCount
	}
	if o.RestHttpResponseTxMsgCount != nil {
		toSerialize["restHttpResponseTxMsgCount"] = o.RestHttpResponseTxMsgCount
	}
	if o.RxByteCount != nil {
		toSerialize["rxByteCount"] = o.RxByteCount
	}
	if o.RxByteRate != nil {
		toSerialize["rxByteRate"] = o.RxByteRate
	}
	if o.RxDiscardedMsgCount != nil {
		toSerialize["rxDiscardedMsgCount"] = o.RxDiscardedMsgCount
	}
	if o.RxMsgCount != nil {
		toSerialize["rxMsgCount"] = o.RxMsgCount
	}
	if o.RxMsgRate != nil {
		toSerialize["rxMsgRate"] = o.RxMsgRate
	}
	if o.ScheduledDisconnectTime != nil {
		toSerialize["scheduledDisconnectTime"] = o.ScheduledDisconnectTime
	}
	if o.SlowSubscriber != nil {
		toSerialize["slowSubscriber"] = o.SlowSubscriber
	}
	if o.SoftwareDate != nil {
		toSerialize["softwareDate"] = o.SoftwareDate
	}
	if o.SoftwareVersion != nil {
		toSerialize["softwareVersion"] = o.SoftwareVersion
	}
	if o.TlsCipherDescription != nil {
		toSerialize["tlsCipherDescription"] = o.TlsCipherDescription
	}
	if o.TlsDowngradedToPlainText != nil {
		toSerialize["tlsDowngradedToPlainText"] = o.TlsDowngradedToPlainText
	}
	if o.TlsVersion != nil {
		toSerialize["tlsVersion"] = o.TlsVersion
	}
	if o.TopicParseErrorRxDiscardedMsgCount != nil {
		toSerialize["topicParseErrorRxDiscardedMsgCount"] = o.TopicParseErrorRxDiscardedMsgCount
	}
	if o.TxByteCount != nil {
		toSerialize["txByteCount"] = o.TxByteCount
	}
	if o.TxByteRate != nil {
		toSerialize["txByteRate"] = o.TxByteRate
	}
	if o.TxDiscardedMsgCount != nil {
		toSerialize["txDiscardedMsgCount"] = o.TxDiscardedMsgCount
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
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.VirtualRouter != nil {
		toSerialize["virtualRouter"] = o.VirtualRouter
	}
	if o.WebInactiveTimeout != nil {
		toSerialize["webInactiveTimeout"] = o.WebInactiveTimeout
	}
	if o.WebMaxPayload != nil {
		toSerialize["webMaxPayload"] = o.WebMaxPayload
	}
	if o.WebParseErrorRxDiscardedMsgCount != nil {
		toSerialize["webParseErrorRxDiscardedMsgCount"] = o.WebParseErrorRxDiscardedMsgCount
	}
	if o.WebRemainingTimeout != nil {
		toSerialize["webRemainingTimeout"] = o.WebRemainingTimeout
	}
	if o.WebRxByteCount != nil {
		toSerialize["webRxByteCount"] = o.WebRxByteCount
	}
	if o.WebRxEncoding != nil {
		toSerialize["webRxEncoding"] = o.WebRxEncoding
	}
	if o.WebRxMsgCount != nil {
		toSerialize["webRxMsgCount"] = o.WebRxMsgCount
	}
	if o.WebRxProtocol != nil {
		toSerialize["webRxProtocol"] = o.WebRxProtocol
	}
	if o.WebRxRequestCount != nil {
		toSerialize["webRxRequestCount"] = o.WebRxRequestCount
	}
	if o.WebRxResponseCount != nil {
		toSerialize["webRxResponseCount"] = o.WebRxResponseCount
	}
	if o.WebRxTcpState != nil {
		toSerialize["webRxTcpState"] = o.WebRxTcpState
	}
	if o.WebRxTlsCipherDescription != nil {
		toSerialize["webRxTlsCipherDescription"] = o.WebRxTlsCipherDescription
	}
	if o.WebRxTlsVersion != nil {
		toSerialize["webRxTlsVersion"] = o.WebRxTlsVersion
	}
	if o.WebSessionId != nil {
		toSerialize["webSessionId"] = o.WebSessionId
	}
	if o.WebTxByteCount != nil {
		toSerialize["webTxByteCount"] = o.WebTxByteCount
	}
	if o.WebTxEncoding != nil {
		toSerialize["webTxEncoding"] = o.WebTxEncoding
	}
	if o.WebTxMsgCount != nil {
		toSerialize["webTxMsgCount"] = o.WebTxMsgCount
	}
	if o.WebTxProtocol != nil {
		toSerialize["webTxProtocol"] = o.WebTxProtocol
	}
	if o.WebTxRequestCount != nil {
		toSerialize["webTxRequestCount"] = o.WebTxRequestCount
	}
	if o.WebTxResponseCount != nil {
		toSerialize["webTxResponseCount"] = o.WebTxResponseCount
	}
	if o.WebTxTcpState != nil {
		toSerialize["webTxTcpState"] = o.WebTxTcpState
	}
	if o.WebTxTlsCipherDescription != nil {
		toSerialize["webTxTlsCipherDescription"] = o.WebTxTlsCipherDescription
	}
	if o.WebTxTlsVersion != nil {
		toSerialize["webTxTlsVersion"] = o.WebTxTlsVersion
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnClient struct {
	value *MsgVpnClient
	isSet bool
}

func (v NullableMsgVpnClient) Get() *MsgVpnClient {
	return v.value
}

func (v *NullableMsgVpnClient) Set(val *MsgVpnClient) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnClient) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnClient(val *MsgVpnClient) *NullableMsgVpnClient {
	return &NullableMsgVpnClient{value: val, isSet: true}
}

func (v NullableMsgVpnClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
