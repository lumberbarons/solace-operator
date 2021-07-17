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

// MsgVpnMqttSession struct for MsgVpnMqttSession
type MsgVpnMqttSession struct {
	// Indicates whether the Client requested a clean (newly created) MQTT Session when connecting. If not clean (already existing), then previously stored messages for QoS 1 subscriptions are delivered.
	Clean *bool `json:"clean,omitempty"`
	// The name of the MQTT Session Client.
	ClientName *string                   `json:"clientName,omitempty"`
	Counter    *MsgVpnMqttSessionCounter `json:"counter,omitempty"`
	// Indicates whether the MQTT Session was created by a Management API.
	CreatedByManagement *bool `json:"createdByManagement,omitempty"`
	// Indicates whether the MQTT Session is durable. Disconnected durable MQTT Sessions are deleted when their expiry time is reached. Disconnected non-durable MQTT Sessions are deleted immediately. Available since 2.21.
	Durable *bool `json:"durable,omitempty"`
	// Indicates whether the MQTT Session is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The timestamp of when the disconnected MQTT session expires and is deleted. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). A value of 0 indicates that the session is either connected, or will never expire. Available since 2.21.
	ExpiryTime *int64 `json:"expiryTime,omitempty"`
	// The maximum size of a packet, including all headers and payload, that the Client has signaled it is willing to accept. A value of zero indicates no limit. Note that there are other broker settings which may further limit packet size. Available since 2.21.
	MaxPacketSize *int64 `json:"maxPacketSize,omitempty"`
	// The number of MQTT connect acknowledgment (CONNACK) refused response packets transmitted to the Client. Available since 2.13.
	MqttConnackErrorTxCount *int64 `json:"mqttConnackErrorTxCount,omitempty"`
	// The number of MQTT connect acknowledgment (CONNACK) accepted response packets transmitted to the Client. Available since 2.13.
	MqttConnackTxCount *int64 `json:"mqttConnackTxCount,omitempty"`
	// The number of MQTT connect (CONNECT) request packets received from the Client. Available since 2.13.
	MqttConnectRxCount *int64 `json:"mqttConnectRxCount,omitempty"`
	// The number of MQTT disconnect (DISCONNECT) request packets received from the Client. Available since 2.13.
	MqttDisconnectRxCount *int64 `json:"mqttDisconnectRxCount,omitempty"`
	// The number of MQTT publish complete (PUBCOMP) packets transmitted to the Client in response to a PUBREL packet. These packets are the fourth and final packet of a QoS 2 protocol exchange. Available since 2.13.
	MqttPubcompTxCount *int64 `json:"mqttPubcompTxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 0 message delivery. Available since 2.13.
	MqttPublishQos0RxCount *int64 `json:"mqttPublishQos0RxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 0 message delivery. Available since 2.13.
	MqttPublishQos0TxCount *int64 `json:"mqttPublishQos0TxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 1 message delivery. Available since 2.13.
	MqttPublishQos1RxCount *int64 `json:"mqttPublishQos1RxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 1 message delivery. Available since 2.13.
	MqttPublishQos1TxCount *int64 `json:"mqttPublishQos1TxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 2 message delivery. Available since 2.13.
	MqttPublishQos2RxCount *int64 `json:"mqttPublishQos2RxCount,omitempty"`
	// The number of MQTT publish received (PUBREC) packets transmitted to the Client in response to a PUBLISH packet with QoS 2. These packets are the second packet of a QoS 2 protocol exchange. Available since 2.13.
	MqttPubrecTxCount *int64 `json:"mqttPubrecTxCount,omitempty"`
	// The number of MQTT publish release (PUBREL) packets received from the Client in response to a PUBREC packet. These packets are the third packet of a QoS 2 protocol exchange. Available since 2.13.
	MqttPubrelRxCount *int64 `json:"mqttPubrelRxCount,omitempty"`
	// The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
	MqttSessionClientId *string `json:"mqttSessionClientId,omitempty"`
	// The virtual router of the MQTT Session. The allowed values and their meaning are:  <pre> \"primary\" - The MQTT Session belongs to the primary virtual router. \"backup\" - The MQTT Session belongs to the backup virtual router. </pre>
	MqttSessionVirtualRouter *string `json:"mqttSessionVirtualRouter,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// The Client Username which owns the MQTT Session.
	Owner *string `json:"owner,omitempty"`
	// Indicates whether consumer acknowledgements (ACKs) received on the active replication Message VPN are propagated to the standby replication Message VPN. Available since 2.14.
	QueueConsumerAckPropagationEnabled *bool `json:"queueConsumerAckPropagationEnabled,omitempty"`
	// The name of the Dead Message Queue (DMQ) used by the MQTT Session Queue. Available since 2.14.
	QueueDeadMsgQueue                            *string         `json:"queueDeadMsgQueue,omitempty"`
	QueueEventBindCountThreshold                 *EventThreshold `json:"queueEventBindCountThreshold,omitempty"`
	QueueEventMsgSpoolUsageThreshold             *EventThreshold `json:"queueEventMsgSpoolUsageThreshold,omitempty"`
	QueueEventRejectLowPriorityMsgLimitThreshold *EventThreshold `json:"queueEventRejectLowPriorityMsgLimitThreshold,omitempty"`
	// The maximum number of consumer flows that can bind to the MQTT Session Queue. Available since 2.14.
	QueueMaxBindCount *int64 `json:"queueMaxBindCount,omitempty"`
	// The maximum number of messages delivered but not acknowledged per flow for the MQTT Session Queue. Available since 2.14.
	QueueMaxDeliveredUnackedMsgsPerFlow *int64 `json:"queueMaxDeliveredUnackedMsgsPerFlow,omitempty"`
	// The maximum message size allowed in the MQTT Session Queue, in bytes (B). Available since 2.14.
	QueueMaxMsgSize *int32 `json:"queueMaxMsgSize,omitempty"`
	// The maximum message spool usage allowed by the MQTT Session Queue, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. Available since 2.14.
	QueueMaxMsgSpoolUsage *int64 `json:"queueMaxMsgSpoolUsage,omitempty"`
	// The maximum number of times the MQTT Session Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. Available since 2.14.
	QueueMaxRedeliveryCount *int64 `json:"queueMaxRedeliveryCount,omitempty"`
	// The maximum time in seconds a message can stay in the MQTT Session Queue when `queueRespectTtlEnabled` is `\"true\"`. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the `queueMaxTtl` configured for the MQTT Session Queue, is exceeded. A value of 0 disables expiry. Available since 2.14.
	QueueMaxTtl *int64 `json:"queueMaxTtl,omitempty"`
	// The name of the MQTT Session Queue.
	QueueName *string `json:"queueName,omitempty"`
	// Indicates whether to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. Available since 2.14.
	QueueRejectLowPriorityMsgEnabled *bool `json:"queueRejectLowPriorityMsgEnabled,omitempty"`
	// The number of messages of any priority in the MQTT Session Queue above which low priority messages are not admitted but higher priority messages are allowed. Available since 2.14.
	QueueRejectLowPriorityMsgLimit *int64 `json:"queueRejectLowPriorityMsgLimit,omitempty"`
	// Indicates whether negative acknowledgements (NACKs) are returned to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The allowed values and their meaning are:  <pre> \"always\" - Always return a negative acknowledgment (NACK) to the sending client on message discard. \"when-queue-enabled\" - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Queue is enabled. \"never\" - Never return a negative acknowledgment (NACK) to the sending client on message discard. </pre>  Available since 2.14.
	QueueRejectMsgToSenderOnDiscardBehavior *string `json:"queueRejectMsgToSenderOnDiscardBehavior,omitempty"`
	// Indicates whether the time-to-live (TTL) for messages in the MQTT Session Queue is respected. When enabled, expired messages are discarded or moved to the DMQ. Available since 2.14.
	QueueRespectTtlEnabled *bool `json:"queueRespectTtlEnabled,omitempty"`
	// The maximum number of outstanding QoS1 and QoS2 messages that the Client has signaled it is willing to accept. Note that there are other broker settings which may further limit the number of outstanding messasges. Available since 2.21.
	RxMax *int64 `json:"rxMax,omitempty"`
	// Indicates whether the MQTT Session has the Will message specified by the Client. The Will message is published if the Client disconnects without sending the MQTT DISCONNECT packet.
	Will *bool `json:"will,omitempty"`
}

// NewMsgVpnMqttSession instantiates a new MsgVpnMqttSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnMqttSession() *MsgVpnMqttSession {
	this := MsgVpnMqttSession{}
	return &this
}

// NewMsgVpnMqttSessionWithDefaults instantiates a new MsgVpnMqttSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnMqttSessionWithDefaults() *MsgVpnMqttSession {
	this := MsgVpnMqttSession{}
	return &this
}

// GetClean returns the Clean field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetClean() bool {
	if o == nil || o.Clean == nil {
		var ret bool
		return ret
	}
	return *o.Clean
}

// GetCleanOk returns a tuple with the Clean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetCleanOk() (*bool, bool) {
	if o == nil || o.Clean == nil {
		return nil, false
	}
	return o.Clean, true
}

// HasClean returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasClean() bool {
	if o != nil && o.Clean != nil {
		return true
	}

	return false
}

// SetClean gets a reference to the given bool and assigns it to the Clean field.
func (o *MsgVpnMqttSession) SetClean(v bool) {
	o.Clean = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *MsgVpnMqttSession) SetClientName(v string) {
	o.ClientName = &v
}

// GetCounter returns the Counter field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetCounter() MsgVpnMqttSessionCounter {
	if o == nil || o.Counter == nil {
		var ret MsgVpnMqttSessionCounter
		return ret
	}
	return *o.Counter
}

// GetCounterOk returns a tuple with the Counter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetCounterOk() (*MsgVpnMqttSessionCounter, bool) {
	if o == nil || o.Counter == nil {
		return nil, false
	}
	return o.Counter, true
}

// HasCounter returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasCounter() bool {
	if o != nil && o.Counter != nil {
		return true
	}

	return false
}

// SetCounter gets a reference to the given MsgVpnMqttSessionCounter and assigns it to the Counter field.
func (o *MsgVpnMqttSession) SetCounter(v MsgVpnMqttSessionCounter) {
	o.Counter = &v
}

// GetCreatedByManagement returns the CreatedByManagement field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetCreatedByManagement() bool {
	if o == nil || o.CreatedByManagement == nil {
		var ret bool
		return ret
	}
	return *o.CreatedByManagement
}

// GetCreatedByManagementOk returns a tuple with the CreatedByManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetCreatedByManagementOk() (*bool, bool) {
	if o == nil || o.CreatedByManagement == nil {
		return nil, false
	}
	return o.CreatedByManagement, true
}

// HasCreatedByManagement returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasCreatedByManagement() bool {
	if o != nil && o.CreatedByManagement != nil {
		return true
	}

	return false
}

// SetCreatedByManagement gets a reference to the given bool and assigns it to the CreatedByManagement field.
func (o *MsgVpnMqttSession) SetCreatedByManagement(v bool) {
	o.CreatedByManagement = &v
}

// GetDurable returns the Durable field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetDurable() bool {
	if o == nil || o.Durable == nil {
		var ret bool
		return ret
	}
	return *o.Durable
}

// GetDurableOk returns a tuple with the Durable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetDurableOk() (*bool, bool) {
	if o == nil || o.Durable == nil {
		return nil, false
	}
	return o.Durable, true
}

// HasDurable returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasDurable() bool {
	if o != nil && o.Durable != nil {
		return true
	}

	return false
}

// SetDurable gets a reference to the given bool and assigns it to the Durable field.
func (o *MsgVpnMqttSession) SetDurable(v bool) {
	o.Durable = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MsgVpnMqttSession) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetExpiryTime() int64 {
	if o == nil || o.ExpiryTime == nil {
		var ret int64
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetExpiryTimeOk() (*int64, bool) {
	if o == nil || o.ExpiryTime == nil {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasExpiryTime() bool {
	if o != nil && o.ExpiryTime != nil {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given int64 and assigns it to the ExpiryTime field.
func (o *MsgVpnMqttSession) SetExpiryTime(v int64) {
	o.ExpiryTime = &v
}

// GetMaxPacketSize returns the MaxPacketSize field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMaxPacketSize() int64 {
	if o == nil || o.MaxPacketSize == nil {
		var ret int64
		return ret
	}
	return *o.MaxPacketSize
}

// GetMaxPacketSizeOk returns a tuple with the MaxPacketSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMaxPacketSizeOk() (*int64, bool) {
	if o == nil || o.MaxPacketSize == nil {
		return nil, false
	}
	return o.MaxPacketSize, true
}

// HasMaxPacketSize returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMaxPacketSize() bool {
	if o != nil && o.MaxPacketSize != nil {
		return true
	}

	return false
}

// SetMaxPacketSize gets a reference to the given int64 and assigns it to the MaxPacketSize field.
func (o *MsgVpnMqttSession) SetMaxPacketSize(v int64) {
	o.MaxPacketSize = &v
}

// GetMqttConnackErrorTxCount returns the MqttConnackErrorTxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttConnackErrorTxCount() int64 {
	if o == nil || o.MqttConnackErrorTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttConnackErrorTxCount
}

// GetMqttConnackErrorTxCountOk returns a tuple with the MqttConnackErrorTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttConnackErrorTxCountOk() (*int64, bool) {
	if o == nil || o.MqttConnackErrorTxCount == nil {
		return nil, false
	}
	return o.MqttConnackErrorTxCount, true
}

// HasMqttConnackErrorTxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttConnackErrorTxCount() bool {
	if o != nil && o.MqttConnackErrorTxCount != nil {
		return true
	}

	return false
}

// SetMqttConnackErrorTxCount gets a reference to the given int64 and assigns it to the MqttConnackErrorTxCount field.
func (o *MsgVpnMqttSession) SetMqttConnackErrorTxCount(v int64) {
	o.MqttConnackErrorTxCount = &v
}

// GetMqttConnackTxCount returns the MqttConnackTxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttConnackTxCount() int64 {
	if o == nil || o.MqttConnackTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttConnackTxCount
}

// GetMqttConnackTxCountOk returns a tuple with the MqttConnackTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttConnackTxCountOk() (*int64, bool) {
	if o == nil || o.MqttConnackTxCount == nil {
		return nil, false
	}
	return o.MqttConnackTxCount, true
}

// HasMqttConnackTxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttConnackTxCount() bool {
	if o != nil && o.MqttConnackTxCount != nil {
		return true
	}

	return false
}

// SetMqttConnackTxCount gets a reference to the given int64 and assigns it to the MqttConnackTxCount field.
func (o *MsgVpnMqttSession) SetMqttConnackTxCount(v int64) {
	o.MqttConnackTxCount = &v
}

// GetMqttConnectRxCount returns the MqttConnectRxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttConnectRxCount() int64 {
	if o == nil || o.MqttConnectRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttConnectRxCount
}

// GetMqttConnectRxCountOk returns a tuple with the MqttConnectRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttConnectRxCountOk() (*int64, bool) {
	if o == nil || o.MqttConnectRxCount == nil {
		return nil, false
	}
	return o.MqttConnectRxCount, true
}

// HasMqttConnectRxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttConnectRxCount() bool {
	if o != nil && o.MqttConnectRxCount != nil {
		return true
	}

	return false
}

// SetMqttConnectRxCount gets a reference to the given int64 and assigns it to the MqttConnectRxCount field.
func (o *MsgVpnMqttSession) SetMqttConnectRxCount(v int64) {
	o.MqttConnectRxCount = &v
}

// GetMqttDisconnectRxCount returns the MqttDisconnectRxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttDisconnectRxCount() int64 {
	if o == nil || o.MqttDisconnectRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttDisconnectRxCount
}

// GetMqttDisconnectRxCountOk returns a tuple with the MqttDisconnectRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttDisconnectRxCountOk() (*int64, bool) {
	if o == nil || o.MqttDisconnectRxCount == nil {
		return nil, false
	}
	return o.MqttDisconnectRxCount, true
}

// HasMqttDisconnectRxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttDisconnectRxCount() bool {
	if o != nil && o.MqttDisconnectRxCount != nil {
		return true
	}

	return false
}

// SetMqttDisconnectRxCount gets a reference to the given int64 and assigns it to the MqttDisconnectRxCount field.
func (o *MsgVpnMqttSession) SetMqttDisconnectRxCount(v int64) {
	o.MqttDisconnectRxCount = &v
}

// GetMqttPubcompTxCount returns the MqttPubcompTxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttPubcompTxCount() int64 {
	if o == nil || o.MqttPubcompTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubcompTxCount
}

// GetMqttPubcompTxCountOk returns a tuple with the MqttPubcompTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttPubcompTxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubcompTxCount == nil {
		return nil, false
	}
	return o.MqttPubcompTxCount, true
}

// HasMqttPubcompTxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttPubcompTxCount() bool {
	if o != nil && o.MqttPubcompTxCount != nil {
		return true
	}

	return false
}

// SetMqttPubcompTxCount gets a reference to the given int64 and assigns it to the MqttPubcompTxCount field.
func (o *MsgVpnMqttSession) SetMqttPubcompTxCount(v int64) {
	o.MqttPubcompTxCount = &v
}

// GetMqttPublishQos0RxCount returns the MqttPublishQos0RxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttPublishQos0RxCount() int64 {
	if o == nil || o.MqttPublishQos0RxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos0RxCount
}

// GetMqttPublishQos0RxCountOk returns a tuple with the MqttPublishQos0RxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttPublishQos0RxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos0RxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos0RxCount, true
}

// HasMqttPublishQos0RxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttPublishQos0RxCount() bool {
	if o != nil && o.MqttPublishQos0RxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos0RxCount gets a reference to the given int64 and assigns it to the MqttPublishQos0RxCount field.
func (o *MsgVpnMqttSession) SetMqttPublishQos0RxCount(v int64) {
	o.MqttPublishQos0RxCount = &v
}

// GetMqttPublishQos0TxCount returns the MqttPublishQos0TxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttPublishQos0TxCount() int64 {
	if o == nil || o.MqttPublishQos0TxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos0TxCount
}

// GetMqttPublishQos0TxCountOk returns a tuple with the MqttPublishQos0TxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttPublishQos0TxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos0TxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos0TxCount, true
}

// HasMqttPublishQos0TxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttPublishQos0TxCount() bool {
	if o != nil && o.MqttPublishQos0TxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos0TxCount gets a reference to the given int64 and assigns it to the MqttPublishQos0TxCount field.
func (o *MsgVpnMqttSession) SetMqttPublishQos0TxCount(v int64) {
	o.MqttPublishQos0TxCount = &v
}

// GetMqttPublishQos1RxCount returns the MqttPublishQos1RxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttPublishQos1RxCount() int64 {
	if o == nil || o.MqttPublishQos1RxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos1RxCount
}

// GetMqttPublishQos1RxCountOk returns a tuple with the MqttPublishQos1RxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttPublishQos1RxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos1RxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos1RxCount, true
}

// HasMqttPublishQos1RxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttPublishQos1RxCount() bool {
	if o != nil && o.MqttPublishQos1RxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos1RxCount gets a reference to the given int64 and assigns it to the MqttPublishQos1RxCount field.
func (o *MsgVpnMqttSession) SetMqttPublishQos1RxCount(v int64) {
	o.MqttPublishQos1RxCount = &v
}

// GetMqttPublishQos1TxCount returns the MqttPublishQos1TxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttPublishQos1TxCount() int64 {
	if o == nil || o.MqttPublishQos1TxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos1TxCount
}

// GetMqttPublishQos1TxCountOk returns a tuple with the MqttPublishQos1TxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttPublishQos1TxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos1TxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos1TxCount, true
}

// HasMqttPublishQos1TxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttPublishQos1TxCount() bool {
	if o != nil && o.MqttPublishQos1TxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos1TxCount gets a reference to the given int64 and assigns it to the MqttPublishQos1TxCount field.
func (o *MsgVpnMqttSession) SetMqttPublishQos1TxCount(v int64) {
	o.MqttPublishQos1TxCount = &v
}

// GetMqttPublishQos2RxCount returns the MqttPublishQos2RxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttPublishQos2RxCount() int64 {
	if o == nil || o.MqttPublishQos2RxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos2RxCount
}

// GetMqttPublishQos2RxCountOk returns a tuple with the MqttPublishQos2RxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttPublishQos2RxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos2RxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos2RxCount, true
}

// HasMqttPublishQos2RxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttPublishQos2RxCount() bool {
	if o != nil && o.MqttPublishQos2RxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos2RxCount gets a reference to the given int64 and assigns it to the MqttPublishQos2RxCount field.
func (o *MsgVpnMqttSession) SetMqttPublishQos2RxCount(v int64) {
	o.MqttPublishQos2RxCount = &v
}

// GetMqttPubrecTxCount returns the MqttPubrecTxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttPubrecTxCount() int64 {
	if o == nil || o.MqttPubrecTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubrecTxCount
}

// GetMqttPubrecTxCountOk returns a tuple with the MqttPubrecTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttPubrecTxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubrecTxCount == nil {
		return nil, false
	}
	return o.MqttPubrecTxCount, true
}

// HasMqttPubrecTxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttPubrecTxCount() bool {
	if o != nil && o.MqttPubrecTxCount != nil {
		return true
	}

	return false
}

// SetMqttPubrecTxCount gets a reference to the given int64 and assigns it to the MqttPubrecTxCount field.
func (o *MsgVpnMqttSession) SetMqttPubrecTxCount(v int64) {
	o.MqttPubrecTxCount = &v
}

// GetMqttPubrelRxCount returns the MqttPubrelRxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttPubrelRxCount() int64 {
	if o == nil || o.MqttPubrelRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubrelRxCount
}

// GetMqttPubrelRxCountOk returns a tuple with the MqttPubrelRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttPubrelRxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubrelRxCount == nil {
		return nil, false
	}
	return o.MqttPubrelRxCount, true
}

// HasMqttPubrelRxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttPubrelRxCount() bool {
	if o != nil && o.MqttPubrelRxCount != nil {
		return true
	}

	return false
}

// SetMqttPubrelRxCount gets a reference to the given int64 and assigns it to the MqttPubrelRxCount field.
func (o *MsgVpnMqttSession) SetMqttPubrelRxCount(v int64) {
	o.MqttPubrelRxCount = &v
}

// GetMqttSessionClientId returns the MqttSessionClientId field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttSessionClientId() string {
	if o == nil || o.MqttSessionClientId == nil {
		var ret string
		return ret
	}
	return *o.MqttSessionClientId
}

// GetMqttSessionClientIdOk returns a tuple with the MqttSessionClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttSessionClientIdOk() (*string, bool) {
	if o == nil || o.MqttSessionClientId == nil {
		return nil, false
	}
	return o.MqttSessionClientId, true
}

// HasMqttSessionClientId returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttSessionClientId() bool {
	if o != nil && o.MqttSessionClientId != nil {
		return true
	}

	return false
}

// SetMqttSessionClientId gets a reference to the given string and assigns it to the MqttSessionClientId field.
func (o *MsgVpnMqttSession) SetMqttSessionClientId(v string) {
	o.MqttSessionClientId = &v
}

// GetMqttSessionVirtualRouter returns the MqttSessionVirtualRouter field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMqttSessionVirtualRouter() string {
	if o == nil || o.MqttSessionVirtualRouter == nil {
		var ret string
		return ret
	}
	return *o.MqttSessionVirtualRouter
}

// GetMqttSessionVirtualRouterOk returns a tuple with the MqttSessionVirtualRouter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMqttSessionVirtualRouterOk() (*string, bool) {
	if o == nil || o.MqttSessionVirtualRouter == nil {
		return nil, false
	}
	return o.MqttSessionVirtualRouter, true
}

// HasMqttSessionVirtualRouter returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMqttSessionVirtualRouter() bool {
	if o != nil && o.MqttSessionVirtualRouter != nil {
		return true
	}

	return false
}

// SetMqttSessionVirtualRouter gets a reference to the given string and assigns it to the MqttSessionVirtualRouter field.
func (o *MsgVpnMqttSession) SetMqttSessionVirtualRouter(v string) {
	o.MqttSessionVirtualRouter = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnMqttSession) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *MsgVpnMqttSession) SetOwner(v string) {
	o.Owner = &v
}

// GetQueueConsumerAckPropagationEnabled returns the QueueConsumerAckPropagationEnabled field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueConsumerAckPropagationEnabled() bool {
	if o == nil || o.QueueConsumerAckPropagationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.QueueConsumerAckPropagationEnabled
}

// GetQueueConsumerAckPropagationEnabledOk returns a tuple with the QueueConsumerAckPropagationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueConsumerAckPropagationEnabledOk() (*bool, bool) {
	if o == nil || o.QueueConsumerAckPropagationEnabled == nil {
		return nil, false
	}
	return o.QueueConsumerAckPropagationEnabled, true
}

// HasQueueConsumerAckPropagationEnabled returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueConsumerAckPropagationEnabled() bool {
	if o != nil && o.QueueConsumerAckPropagationEnabled != nil {
		return true
	}

	return false
}

// SetQueueConsumerAckPropagationEnabled gets a reference to the given bool and assigns it to the QueueConsumerAckPropagationEnabled field.
func (o *MsgVpnMqttSession) SetQueueConsumerAckPropagationEnabled(v bool) {
	o.QueueConsumerAckPropagationEnabled = &v
}

// GetQueueDeadMsgQueue returns the QueueDeadMsgQueue field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueDeadMsgQueue() string {
	if o == nil || o.QueueDeadMsgQueue == nil {
		var ret string
		return ret
	}
	return *o.QueueDeadMsgQueue
}

// GetQueueDeadMsgQueueOk returns a tuple with the QueueDeadMsgQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueDeadMsgQueueOk() (*string, bool) {
	if o == nil || o.QueueDeadMsgQueue == nil {
		return nil, false
	}
	return o.QueueDeadMsgQueue, true
}

// HasQueueDeadMsgQueue returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueDeadMsgQueue() bool {
	if o != nil && o.QueueDeadMsgQueue != nil {
		return true
	}

	return false
}

// SetQueueDeadMsgQueue gets a reference to the given string and assigns it to the QueueDeadMsgQueue field.
func (o *MsgVpnMqttSession) SetQueueDeadMsgQueue(v string) {
	o.QueueDeadMsgQueue = &v
}

// GetQueueEventBindCountThreshold returns the QueueEventBindCountThreshold field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueEventBindCountThreshold() EventThreshold {
	if o == nil || o.QueueEventBindCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.QueueEventBindCountThreshold
}

// GetQueueEventBindCountThresholdOk returns a tuple with the QueueEventBindCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueEventBindCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.QueueEventBindCountThreshold == nil {
		return nil, false
	}
	return o.QueueEventBindCountThreshold, true
}

// HasQueueEventBindCountThreshold returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueEventBindCountThreshold() bool {
	if o != nil && o.QueueEventBindCountThreshold != nil {
		return true
	}

	return false
}

// SetQueueEventBindCountThreshold gets a reference to the given EventThreshold and assigns it to the QueueEventBindCountThreshold field.
func (o *MsgVpnMqttSession) SetQueueEventBindCountThreshold(v EventThreshold) {
	o.QueueEventBindCountThreshold = &v
}

// GetQueueEventMsgSpoolUsageThreshold returns the QueueEventMsgSpoolUsageThreshold field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueEventMsgSpoolUsageThreshold() EventThreshold {
	if o == nil || o.QueueEventMsgSpoolUsageThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.QueueEventMsgSpoolUsageThreshold
}

// GetQueueEventMsgSpoolUsageThresholdOk returns a tuple with the QueueEventMsgSpoolUsageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueEventMsgSpoolUsageThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.QueueEventMsgSpoolUsageThreshold == nil {
		return nil, false
	}
	return o.QueueEventMsgSpoolUsageThreshold, true
}

// HasQueueEventMsgSpoolUsageThreshold returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueEventMsgSpoolUsageThreshold() bool {
	if o != nil && o.QueueEventMsgSpoolUsageThreshold != nil {
		return true
	}

	return false
}

// SetQueueEventMsgSpoolUsageThreshold gets a reference to the given EventThreshold and assigns it to the QueueEventMsgSpoolUsageThreshold field.
func (o *MsgVpnMqttSession) SetQueueEventMsgSpoolUsageThreshold(v EventThreshold) {
	o.QueueEventMsgSpoolUsageThreshold = &v
}

// GetQueueEventRejectLowPriorityMsgLimitThreshold returns the QueueEventRejectLowPriorityMsgLimitThreshold field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueEventRejectLowPriorityMsgLimitThreshold() EventThreshold {
	if o == nil || o.QueueEventRejectLowPriorityMsgLimitThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.QueueEventRejectLowPriorityMsgLimitThreshold
}

// GetQueueEventRejectLowPriorityMsgLimitThresholdOk returns a tuple with the QueueEventRejectLowPriorityMsgLimitThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueEventRejectLowPriorityMsgLimitThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.QueueEventRejectLowPriorityMsgLimitThreshold == nil {
		return nil, false
	}
	return o.QueueEventRejectLowPriorityMsgLimitThreshold, true
}

// HasQueueEventRejectLowPriorityMsgLimitThreshold returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueEventRejectLowPriorityMsgLimitThreshold() bool {
	if o != nil && o.QueueEventRejectLowPriorityMsgLimitThreshold != nil {
		return true
	}

	return false
}

// SetQueueEventRejectLowPriorityMsgLimitThreshold gets a reference to the given EventThreshold and assigns it to the QueueEventRejectLowPriorityMsgLimitThreshold field.
func (o *MsgVpnMqttSession) SetQueueEventRejectLowPriorityMsgLimitThreshold(v EventThreshold) {
	o.QueueEventRejectLowPriorityMsgLimitThreshold = &v
}

// GetQueueMaxBindCount returns the QueueMaxBindCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueMaxBindCount() int64 {
	if o == nil || o.QueueMaxBindCount == nil {
		var ret int64
		return ret
	}
	return *o.QueueMaxBindCount
}

// GetQueueMaxBindCountOk returns a tuple with the QueueMaxBindCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueMaxBindCountOk() (*int64, bool) {
	if o == nil || o.QueueMaxBindCount == nil {
		return nil, false
	}
	return o.QueueMaxBindCount, true
}

// HasQueueMaxBindCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueMaxBindCount() bool {
	if o != nil && o.QueueMaxBindCount != nil {
		return true
	}

	return false
}

// SetQueueMaxBindCount gets a reference to the given int64 and assigns it to the QueueMaxBindCount field.
func (o *MsgVpnMqttSession) SetQueueMaxBindCount(v int64) {
	o.QueueMaxBindCount = &v
}

// GetQueueMaxDeliveredUnackedMsgsPerFlow returns the QueueMaxDeliveredUnackedMsgsPerFlow field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueMaxDeliveredUnackedMsgsPerFlow() int64 {
	if o == nil || o.QueueMaxDeliveredUnackedMsgsPerFlow == nil {
		var ret int64
		return ret
	}
	return *o.QueueMaxDeliveredUnackedMsgsPerFlow
}

// GetQueueMaxDeliveredUnackedMsgsPerFlowOk returns a tuple with the QueueMaxDeliveredUnackedMsgsPerFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueMaxDeliveredUnackedMsgsPerFlowOk() (*int64, bool) {
	if o == nil || o.QueueMaxDeliveredUnackedMsgsPerFlow == nil {
		return nil, false
	}
	return o.QueueMaxDeliveredUnackedMsgsPerFlow, true
}

// HasQueueMaxDeliveredUnackedMsgsPerFlow returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueMaxDeliveredUnackedMsgsPerFlow() bool {
	if o != nil && o.QueueMaxDeliveredUnackedMsgsPerFlow != nil {
		return true
	}

	return false
}

// SetQueueMaxDeliveredUnackedMsgsPerFlow gets a reference to the given int64 and assigns it to the QueueMaxDeliveredUnackedMsgsPerFlow field.
func (o *MsgVpnMqttSession) SetQueueMaxDeliveredUnackedMsgsPerFlow(v int64) {
	o.QueueMaxDeliveredUnackedMsgsPerFlow = &v
}

// GetQueueMaxMsgSize returns the QueueMaxMsgSize field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueMaxMsgSize() int32 {
	if o == nil || o.QueueMaxMsgSize == nil {
		var ret int32
		return ret
	}
	return *o.QueueMaxMsgSize
}

// GetQueueMaxMsgSizeOk returns a tuple with the QueueMaxMsgSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueMaxMsgSizeOk() (*int32, bool) {
	if o == nil || o.QueueMaxMsgSize == nil {
		return nil, false
	}
	return o.QueueMaxMsgSize, true
}

// HasQueueMaxMsgSize returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueMaxMsgSize() bool {
	if o != nil && o.QueueMaxMsgSize != nil {
		return true
	}

	return false
}

// SetQueueMaxMsgSize gets a reference to the given int32 and assigns it to the QueueMaxMsgSize field.
func (o *MsgVpnMqttSession) SetQueueMaxMsgSize(v int32) {
	o.QueueMaxMsgSize = &v
}

// GetQueueMaxMsgSpoolUsage returns the QueueMaxMsgSpoolUsage field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueMaxMsgSpoolUsage() int64 {
	if o == nil || o.QueueMaxMsgSpoolUsage == nil {
		var ret int64
		return ret
	}
	return *o.QueueMaxMsgSpoolUsage
}

// GetQueueMaxMsgSpoolUsageOk returns a tuple with the QueueMaxMsgSpoolUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueMaxMsgSpoolUsageOk() (*int64, bool) {
	if o == nil || o.QueueMaxMsgSpoolUsage == nil {
		return nil, false
	}
	return o.QueueMaxMsgSpoolUsage, true
}

// HasQueueMaxMsgSpoolUsage returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueMaxMsgSpoolUsage() bool {
	if o != nil && o.QueueMaxMsgSpoolUsage != nil {
		return true
	}

	return false
}

// SetQueueMaxMsgSpoolUsage gets a reference to the given int64 and assigns it to the QueueMaxMsgSpoolUsage field.
func (o *MsgVpnMqttSession) SetQueueMaxMsgSpoolUsage(v int64) {
	o.QueueMaxMsgSpoolUsage = &v
}

// GetQueueMaxRedeliveryCount returns the QueueMaxRedeliveryCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueMaxRedeliveryCount() int64 {
	if o == nil || o.QueueMaxRedeliveryCount == nil {
		var ret int64
		return ret
	}
	return *o.QueueMaxRedeliveryCount
}

// GetQueueMaxRedeliveryCountOk returns a tuple with the QueueMaxRedeliveryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueMaxRedeliveryCountOk() (*int64, bool) {
	if o == nil || o.QueueMaxRedeliveryCount == nil {
		return nil, false
	}
	return o.QueueMaxRedeliveryCount, true
}

// HasQueueMaxRedeliveryCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueMaxRedeliveryCount() bool {
	if o != nil && o.QueueMaxRedeliveryCount != nil {
		return true
	}

	return false
}

// SetQueueMaxRedeliveryCount gets a reference to the given int64 and assigns it to the QueueMaxRedeliveryCount field.
func (o *MsgVpnMqttSession) SetQueueMaxRedeliveryCount(v int64) {
	o.QueueMaxRedeliveryCount = &v
}

// GetQueueMaxTtl returns the QueueMaxTtl field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueMaxTtl() int64 {
	if o == nil || o.QueueMaxTtl == nil {
		var ret int64
		return ret
	}
	return *o.QueueMaxTtl
}

// GetQueueMaxTtlOk returns a tuple with the QueueMaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueMaxTtlOk() (*int64, bool) {
	if o == nil || o.QueueMaxTtl == nil {
		return nil, false
	}
	return o.QueueMaxTtl, true
}

// HasQueueMaxTtl returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueMaxTtl() bool {
	if o != nil && o.QueueMaxTtl != nil {
		return true
	}

	return false
}

// SetQueueMaxTtl gets a reference to the given int64 and assigns it to the QueueMaxTtl field.
func (o *MsgVpnMqttSession) SetQueueMaxTtl(v int64) {
	o.QueueMaxTtl = &v
}

// GetQueueName returns the QueueName field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueName() string {
	if o == nil || o.QueueName == nil {
		var ret string
		return ret
	}
	return *o.QueueName
}

// GetQueueNameOk returns a tuple with the QueueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueNameOk() (*string, bool) {
	if o == nil || o.QueueName == nil {
		return nil, false
	}
	return o.QueueName, true
}

// HasQueueName returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueName() bool {
	if o != nil && o.QueueName != nil {
		return true
	}

	return false
}

// SetQueueName gets a reference to the given string and assigns it to the QueueName field.
func (o *MsgVpnMqttSession) SetQueueName(v string) {
	o.QueueName = &v
}

// GetQueueRejectLowPriorityMsgEnabled returns the QueueRejectLowPriorityMsgEnabled field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueRejectLowPriorityMsgEnabled() bool {
	if o == nil || o.QueueRejectLowPriorityMsgEnabled == nil {
		var ret bool
		return ret
	}
	return *o.QueueRejectLowPriorityMsgEnabled
}

// GetQueueRejectLowPriorityMsgEnabledOk returns a tuple with the QueueRejectLowPriorityMsgEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueRejectLowPriorityMsgEnabledOk() (*bool, bool) {
	if o == nil || o.QueueRejectLowPriorityMsgEnabled == nil {
		return nil, false
	}
	return o.QueueRejectLowPriorityMsgEnabled, true
}

// HasQueueRejectLowPriorityMsgEnabled returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueRejectLowPriorityMsgEnabled() bool {
	if o != nil && o.QueueRejectLowPriorityMsgEnabled != nil {
		return true
	}

	return false
}

// SetQueueRejectLowPriorityMsgEnabled gets a reference to the given bool and assigns it to the QueueRejectLowPriorityMsgEnabled field.
func (o *MsgVpnMqttSession) SetQueueRejectLowPriorityMsgEnabled(v bool) {
	o.QueueRejectLowPriorityMsgEnabled = &v
}

// GetQueueRejectLowPriorityMsgLimit returns the QueueRejectLowPriorityMsgLimit field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueRejectLowPriorityMsgLimit() int64 {
	if o == nil || o.QueueRejectLowPriorityMsgLimit == nil {
		var ret int64
		return ret
	}
	return *o.QueueRejectLowPriorityMsgLimit
}

// GetQueueRejectLowPriorityMsgLimitOk returns a tuple with the QueueRejectLowPriorityMsgLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueRejectLowPriorityMsgLimitOk() (*int64, bool) {
	if o == nil || o.QueueRejectLowPriorityMsgLimit == nil {
		return nil, false
	}
	return o.QueueRejectLowPriorityMsgLimit, true
}

// HasQueueRejectLowPriorityMsgLimit returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueRejectLowPriorityMsgLimit() bool {
	if o != nil && o.QueueRejectLowPriorityMsgLimit != nil {
		return true
	}

	return false
}

// SetQueueRejectLowPriorityMsgLimit gets a reference to the given int64 and assigns it to the QueueRejectLowPriorityMsgLimit field.
func (o *MsgVpnMqttSession) SetQueueRejectLowPriorityMsgLimit(v int64) {
	o.QueueRejectLowPriorityMsgLimit = &v
}

// GetQueueRejectMsgToSenderOnDiscardBehavior returns the QueueRejectMsgToSenderOnDiscardBehavior field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueRejectMsgToSenderOnDiscardBehavior() string {
	if o == nil || o.QueueRejectMsgToSenderOnDiscardBehavior == nil {
		var ret string
		return ret
	}
	return *o.QueueRejectMsgToSenderOnDiscardBehavior
}

// GetQueueRejectMsgToSenderOnDiscardBehaviorOk returns a tuple with the QueueRejectMsgToSenderOnDiscardBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueRejectMsgToSenderOnDiscardBehaviorOk() (*string, bool) {
	if o == nil || o.QueueRejectMsgToSenderOnDiscardBehavior == nil {
		return nil, false
	}
	return o.QueueRejectMsgToSenderOnDiscardBehavior, true
}

// HasQueueRejectMsgToSenderOnDiscardBehavior returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueRejectMsgToSenderOnDiscardBehavior() bool {
	if o != nil && o.QueueRejectMsgToSenderOnDiscardBehavior != nil {
		return true
	}

	return false
}

// SetQueueRejectMsgToSenderOnDiscardBehavior gets a reference to the given string and assigns it to the QueueRejectMsgToSenderOnDiscardBehavior field.
func (o *MsgVpnMqttSession) SetQueueRejectMsgToSenderOnDiscardBehavior(v string) {
	o.QueueRejectMsgToSenderOnDiscardBehavior = &v
}

// GetQueueRespectTtlEnabled returns the QueueRespectTtlEnabled field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetQueueRespectTtlEnabled() bool {
	if o == nil || o.QueueRespectTtlEnabled == nil {
		var ret bool
		return ret
	}
	return *o.QueueRespectTtlEnabled
}

// GetQueueRespectTtlEnabledOk returns a tuple with the QueueRespectTtlEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetQueueRespectTtlEnabledOk() (*bool, bool) {
	if o == nil || o.QueueRespectTtlEnabled == nil {
		return nil, false
	}
	return o.QueueRespectTtlEnabled, true
}

// HasQueueRespectTtlEnabled returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasQueueRespectTtlEnabled() bool {
	if o != nil && o.QueueRespectTtlEnabled != nil {
		return true
	}

	return false
}

// SetQueueRespectTtlEnabled gets a reference to the given bool and assigns it to the QueueRespectTtlEnabled field.
func (o *MsgVpnMqttSession) SetQueueRespectTtlEnabled(v bool) {
	o.QueueRespectTtlEnabled = &v
}

// GetRxMax returns the RxMax field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetRxMax() int64 {
	if o == nil || o.RxMax == nil {
		var ret int64
		return ret
	}
	return *o.RxMax
}

// GetRxMaxOk returns a tuple with the RxMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetRxMaxOk() (*int64, bool) {
	if o == nil || o.RxMax == nil {
		return nil, false
	}
	return o.RxMax, true
}

// HasRxMax returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasRxMax() bool {
	if o != nil && o.RxMax != nil {
		return true
	}

	return false
}

// SetRxMax gets a reference to the given int64 and assigns it to the RxMax field.
func (o *MsgVpnMqttSession) SetRxMax(v int64) {
	o.RxMax = &v
}

// GetWill returns the Will field value if set, zero value otherwise.
func (o *MsgVpnMqttSession) GetWill() bool {
	if o == nil || o.Will == nil {
		var ret bool
		return ret
	}
	return *o.Will
}

// GetWillOk returns a tuple with the Will field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSession) GetWillOk() (*bool, bool) {
	if o == nil || o.Will == nil {
		return nil, false
	}
	return o.Will, true
}

// HasWill returns a boolean if a field has been set.
func (o *MsgVpnMqttSession) HasWill() bool {
	if o != nil && o.Will != nil {
		return true
	}

	return false
}

// SetWill gets a reference to the given bool and assigns it to the Will field.
func (o *MsgVpnMqttSession) SetWill(v bool) {
	o.Will = &v
}

func (o MsgVpnMqttSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clean != nil {
		toSerialize["clean"] = o.Clean
	}
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.Counter != nil {
		toSerialize["counter"] = o.Counter
	}
	if o.CreatedByManagement != nil {
		toSerialize["createdByManagement"] = o.CreatedByManagement
	}
	if o.Durable != nil {
		toSerialize["durable"] = o.Durable
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExpiryTime != nil {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if o.MaxPacketSize != nil {
		toSerialize["maxPacketSize"] = o.MaxPacketSize
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
	if o.MqttSessionClientId != nil {
		toSerialize["mqttSessionClientId"] = o.MqttSessionClientId
	}
	if o.MqttSessionVirtualRouter != nil {
		toSerialize["mqttSessionVirtualRouter"] = o.MqttSessionVirtualRouter
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.QueueConsumerAckPropagationEnabled != nil {
		toSerialize["queueConsumerAckPropagationEnabled"] = o.QueueConsumerAckPropagationEnabled
	}
	if o.QueueDeadMsgQueue != nil {
		toSerialize["queueDeadMsgQueue"] = o.QueueDeadMsgQueue
	}
	if o.QueueEventBindCountThreshold != nil {
		toSerialize["queueEventBindCountThreshold"] = o.QueueEventBindCountThreshold
	}
	if o.QueueEventMsgSpoolUsageThreshold != nil {
		toSerialize["queueEventMsgSpoolUsageThreshold"] = o.QueueEventMsgSpoolUsageThreshold
	}
	if o.QueueEventRejectLowPriorityMsgLimitThreshold != nil {
		toSerialize["queueEventRejectLowPriorityMsgLimitThreshold"] = o.QueueEventRejectLowPriorityMsgLimitThreshold
	}
	if o.QueueMaxBindCount != nil {
		toSerialize["queueMaxBindCount"] = o.QueueMaxBindCount
	}
	if o.QueueMaxDeliveredUnackedMsgsPerFlow != nil {
		toSerialize["queueMaxDeliveredUnackedMsgsPerFlow"] = o.QueueMaxDeliveredUnackedMsgsPerFlow
	}
	if o.QueueMaxMsgSize != nil {
		toSerialize["queueMaxMsgSize"] = o.QueueMaxMsgSize
	}
	if o.QueueMaxMsgSpoolUsage != nil {
		toSerialize["queueMaxMsgSpoolUsage"] = o.QueueMaxMsgSpoolUsage
	}
	if o.QueueMaxRedeliveryCount != nil {
		toSerialize["queueMaxRedeliveryCount"] = o.QueueMaxRedeliveryCount
	}
	if o.QueueMaxTtl != nil {
		toSerialize["queueMaxTtl"] = o.QueueMaxTtl
	}
	if o.QueueName != nil {
		toSerialize["queueName"] = o.QueueName
	}
	if o.QueueRejectLowPriorityMsgEnabled != nil {
		toSerialize["queueRejectLowPriorityMsgEnabled"] = o.QueueRejectLowPriorityMsgEnabled
	}
	if o.QueueRejectLowPriorityMsgLimit != nil {
		toSerialize["queueRejectLowPriorityMsgLimit"] = o.QueueRejectLowPriorityMsgLimit
	}
	if o.QueueRejectMsgToSenderOnDiscardBehavior != nil {
		toSerialize["queueRejectMsgToSenderOnDiscardBehavior"] = o.QueueRejectMsgToSenderOnDiscardBehavior
	}
	if o.QueueRespectTtlEnabled != nil {
		toSerialize["queueRespectTtlEnabled"] = o.QueueRespectTtlEnabled
	}
	if o.RxMax != nil {
		toSerialize["rxMax"] = o.RxMax
	}
	if o.Will != nil {
		toSerialize["will"] = o.Will
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnMqttSession struct {
	value *MsgVpnMqttSession
	isSet bool
}

func (v NullableMsgVpnMqttSession) Get() *MsgVpnMqttSession {
	return v.value
}

func (v *NullableMsgVpnMqttSession) Set(val *MsgVpnMqttSession) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnMqttSession) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnMqttSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnMqttSession(val *MsgVpnMqttSession) *NullableMsgVpnMqttSession {
	return &NullableMsgVpnMqttSession{value: val, isSet: true}
}

func (v NullableMsgVpnMqttSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnMqttSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
