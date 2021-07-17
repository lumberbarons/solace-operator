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

// MsgVpnMqttSessionCounter The counters for the MQTT Session. Deprecated since 2.13. All attributes in this object have been moved to the MsgVpnMqttSession object.
type MsgVpnMqttSessionCounter struct {
	// The number of MQTT connect acknowledgment (CONNACK) refused response packets transmitted to the Client. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttConnackErrorTxCount *int64 `json:"mqttConnackErrorTxCount,omitempty"`
	// The number of MQTT connect acknowledgment (CONNACK) accepted response packets transmitted to the Client. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttConnackTxCount *int64 `json:"mqttConnackTxCount,omitempty"`
	// The number of MQTT connect (CONNECT) request packets received from the Client. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttConnectRxCount *int64 `json:"mqttConnectRxCount,omitempty"`
	// The number of MQTT disconnect (DISCONNECT) request packets received from the Client. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttDisconnectRxCount *int64 `json:"mqttDisconnectRxCount,omitempty"`
	// The number of MQTT publish complete (PUBCOMP) packets transmitted to the Client in response to a PUBREL packet. These packets are the fourth and final packet of a QoS 2 protocol exchange. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttPubcompTxCount *int64 `json:"mqttPubcompTxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 0 message delivery. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttPublishQos0RxCount *int64 `json:"mqttPublishQos0RxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 0 message delivery. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttPublishQos0TxCount *int64 `json:"mqttPublishQos0TxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 1 message delivery. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttPublishQos1RxCount *int64 `json:"mqttPublishQos1RxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets transmitted to the Client for QoS 1 message delivery. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttPublishQos1TxCount *int64 `json:"mqttPublishQos1TxCount,omitempty"`
	// The number of MQTT publish message (PUBLISH) request packets received from the Client for QoS 2 message delivery. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttPublishQos2RxCount *int64 `json:"mqttPublishQos2RxCount,omitempty"`
	// The number of MQTT publish received (PUBREC) packets transmitted to the Client in response to a PUBLISH packet with QoS 2. These packets are the second packet of a QoS 2 protocol exchange. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttPubrecTxCount *int64 `json:"mqttPubrecTxCount,omitempty"`
	// The number of MQTT publish release (PUBREL) packets received from the Client in response to a PUBREC packet. These packets are the third packet of a QoS 2 protocol exchange. Deprecated since 2.13. This attribute has been moved to the MsgVpnMqttSession object.
	MqttPubrelRxCount *int64 `json:"mqttPubrelRxCount,omitempty"`
}

// NewMsgVpnMqttSessionCounter instantiates a new MsgVpnMqttSessionCounter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnMqttSessionCounter() *MsgVpnMqttSessionCounter {
	this := MsgVpnMqttSessionCounter{}
	return &this
}

// NewMsgVpnMqttSessionCounterWithDefaults instantiates a new MsgVpnMqttSessionCounter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnMqttSessionCounterWithDefaults() *MsgVpnMqttSessionCounter {
	this := MsgVpnMqttSessionCounter{}
	return &this
}

// GetMqttConnackErrorTxCount returns the MqttConnackErrorTxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttConnackErrorTxCount() int64 {
	if o == nil || o.MqttConnackErrorTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttConnackErrorTxCount
}

// GetMqttConnackErrorTxCountOk returns a tuple with the MqttConnackErrorTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttConnackErrorTxCountOk() (*int64, bool) {
	if o == nil || o.MqttConnackErrorTxCount == nil {
		return nil, false
	}
	return o.MqttConnackErrorTxCount, true
}

// HasMqttConnackErrorTxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttConnackErrorTxCount() bool {
	if o != nil && o.MqttConnackErrorTxCount != nil {
		return true
	}

	return false
}

// SetMqttConnackErrorTxCount gets a reference to the given int64 and assigns it to the MqttConnackErrorTxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttConnackErrorTxCount(v int64) {
	o.MqttConnackErrorTxCount = &v
}

// GetMqttConnackTxCount returns the MqttConnackTxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttConnackTxCount() int64 {
	if o == nil || o.MqttConnackTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttConnackTxCount
}

// GetMqttConnackTxCountOk returns a tuple with the MqttConnackTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttConnackTxCountOk() (*int64, bool) {
	if o == nil || o.MqttConnackTxCount == nil {
		return nil, false
	}
	return o.MqttConnackTxCount, true
}

// HasMqttConnackTxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttConnackTxCount() bool {
	if o != nil && o.MqttConnackTxCount != nil {
		return true
	}

	return false
}

// SetMqttConnackTxCount gets a reference to the given int64 and assigns it to the MqttConnackTxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttConnackTxCount(v int64) {
	o.MqttConnackTxCount = &v
}

// GetMqttConnectRxCount returns the MqttConnectRxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttConnectRxCount() int64 {
	if o == nil || o.MqttConnectRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttConnectRxCount
}

// GetMqttConnectRxCountOk returns a tuple with the MqttConnectRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttConnectRxCountOk() (*int64, bool) {
	if o == nil || o.MqttConnectRxCount == nil {
		return nil, false
	}
	return o.MqttConnectRxCount, true
}

// HasMqttConnectRxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttConnectRxCount() bool {
	if o != nil && o.MqttConnectRxCount != nil {
		return true
	}

	return false
}

// SetMqttConnectRxCount gets a reference to the given int64 and assigns it to the MqttConnectRxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttConnectRxCount(v int64) {
	o.MqttConnectRxCount = &v
}

// GetMqttDisconnectRxCount returns the MqttDisconnectRxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttDisconnectRxCount() int64 {
	if o == nil || o.MqttDisconnectRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttDisconnectRxCount
}

// GetMqttDisconnectRxCountOk returns a tuple with the MqttDisconnectRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttDisconnectRxCountOk() (*int64, bool) {
	if o == nil || o.MqttDisconnectRxCount == nil {
		return nil, false
	}
	return o.MqttDisconnectRxCount, true
}

// HasMqttDisconnectRxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttDisconnectRxCount() bool {
	if o != nil && o.MqttDisconnectRxCount != nil {
		return true
	}

	return false
}

// SetMqttDisconnectRxCount gets a reference to the given int64 and assigns it to the MqttDisconnectRxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttDisconnectRxCount(v int64) {
	o.MqttDisconnectRxCount = &v
}

// GetMqttPubcompTxCount returns the MqttPubcompTxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttPubcompTxCount() int64 {
	if o == nil || o.MqttPubcompTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubcompTxCount
}

// GetMqttPubcompTxCountOk returns a tuple with the MqttPubcompTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttPubcompTxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubcompTxCount == nil {
		return nil, false
	}
	return o.MqttPubcompTxCount, true
}

// HasMqttPubcompTxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttPubcompTxCount() bool {
	if o != nil && o.MqttPubcompTxCount != nil {
		return true
	}

	return false
}

// SetMqttPubcompTxCount gets a reference to the given int64 and assigns it to the MqttPubcompTxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttPubcompTxCount(v int64) {
	o.MqttPubcompTxCount = &v
}

// GetMqttPublishQos0RxCount returns the MqttPublishQos0RxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos0RxCount() int64 {
	if o == nil || o.MqttPublishQos0RxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos0RxCount
}

// GetMqttPublishQos0RxCountOk returns a tuple with the MqttPublishQos0RxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos0RxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos0RxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos0RxCount, true
}

// HasMqttPublishQos0RxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttPublishQos0RxCount() bool {
	if o != nil && o.MqttPublishQos0RxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos0RxCount gets a reference to the given int64 and assigns it to the MqttPublishQos0RxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttPublishQos0RxCount(v int64) {
	o.MqttPublishQos0RxCount = &v
}

// GetMqttPublishQos0TxCount returns the MqttPublishQos0TxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos0TxCount() int64 {
	if o == nil || o.MqttPublishQos0TxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos0TxCount
}

// GetMqttPublishQos0TxCountOk returns a tuple with the MqttPublishQos0TxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos0TxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos0TxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos0TxCount, true
}

// HasMqttPublishQos0TxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttPublishQos0TxCount() bool {
	if o != nil && o.MqttPublishQos0TxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos0TxCount gets a reference to the given int64 and assigns it to the MqttPublishQos0TxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttPublishQos0TxCount(v int64) {
	o.MqttPublishQos0TxCount = &v
}

// GetMqttPublishQos1RxCount returns the MqttPublishQos1RxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos1RxCount() int64 {
	if o == nil || o.MqttPublishQos1RxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos1RxCount
}

// GetMqttPublishQos1RxCountOk returns a tuple with the MqttPublishQos1RxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos1RxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos1RxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos1RxCount, true
}

// HasMqttPublishQos1RxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttPublishQos1RxCount() bool {
	if o != nil && o.MqttPublishQos1RxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos1RxCount gets a reference to the given int64 and assigns it to the MqttPublishQos1RxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttPublishQos1RxCount(v int64) {
	o.MqttPublishQos1RxCount = &v
}

// GetMqttPublishQos1TxCount returns the MqttPublishQos1TxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos1TxCount() int64 {
	if o == nil || o.MqttPublishQos1TxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos1TxCount
}

// GetMqttPublishQos1TxCountOk returns a tuple with the MqttPublishQos1TxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos1TxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos1TxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos1TxCount, true
}

// HasMqttPublishQos1TxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttPublishQos1TxCount() bool {
	if o != nil && o.MqttPublishQos1TxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos1TxCount gets a reference to the given int64 and assigns it to the MqttPublishQos1TxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttPublishQos1TxCount(v int64) {
	o.MqttPublishQos1TxCount = &v
}

// GetMqttPublishQos2RxCount returns the MqttPublishQos2RxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos2RxCount() int64 {
	if o == nil || o.MqttPublishQos2RxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPublishQos2RxCount
}

// GetMqttPublishQos2RxCountOk returns a tuple with the MqttPublishQos2RxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttPublishQos2RxCountOk() (*int64, bool) {
	if o == nil || o.MqttPublishQos2RxCount == nil {
		return nil, false
	}
	return o.MqttPublishQos2RxCount, true
}

// HasMqttPublishQos2RxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttPublishQos2RxCount() bool {
	if o != nil && o.MqttPublishQos2RxCount != nil {
		return true
	}

	return false
}

// SetMqttPublishQos2RxCount gets a reference to the given int64 and assigns it to the MqttPublishQos2RxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttPublishQos2RxCount(v int64) {
	o.MqttPublishQos2RxCount = &v
}

// GetMqttPubrecTxCount returns the MqttPubrecTxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttPubrecTxCount() int64 {
	if o == nil || o.MqttPubrecTxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubrecTxCount
}

// GetMqttPubrecTxCountOk returns a tuple with the MqttPubrecTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttPubrecTxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubrecTxCount == nil {
		return nil, false
	}
	return o.MqttPubrecTxCount, true
}

// HasMqttPubrecTxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttPubrecTxCount() bool {
	if o != nil && o.MqttPubrecTxCount != nil {
		return true
	}

	return false
}

// SetMqttPubrecTxCount gets a reference to the given int64 and assigns it to the MqttPubrecTxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttPubrecTxCount(v int64) {
	o.MqttPubrecTxCount = &v
}

// GetMqttPubrelRxCount returns the MqttPubrelRxCount field value if set, zero value otherwise.
func (o *MsgVpnMqttSessionCounter) GetMqttPubrelRxCount() int64 {
	if o == nil || o.MqttPubrelRxCount == nil {
		var ret int64
		return ret
	}
	return *o.MqttPubrelRxCount
}

// GetMqttPubrelRxCountOk returns a tuple with the MqttPubrelRxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttSessionCounter) GetMqttPubrelRxCountOk() (*int64, bool) {
	if o == nil || o.MqttPubrelRxCount == nil {
		return nil, false
	}
	return o.MqttPubrelRxCount, true
}

// HasMqttPubrelRxCount returns a boolean if a field has been set.
func (o *MsgVpnMqttSessionCounter) HasMqttPubrelRxCount() bool {
	if o != nil && o.MqttPubrelRxCount != nil {
		return true
	}

	return false
}

// SetMqttPubrelRxCount gets a reference to the given int64 and assigns it to the MqttPubrelRxCount field.
func (o *MsgVpnMqttSessionCounter) SetMqttPubrelRxCount(v int64) {
	o.MqttPubrelRxCount = &v
}

func (o MsgVpnMqttSessionCounter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullableMsgVpnMqttSessionCounter struct {
	value *MsgVpnMqttSessionCounter
	isSet bool
}

func (v NullableMsgVpnMqttSessionCounter) Get() *MsgVpnMqttSessionCounter {
	return v.value
}

func (v *NullableMsgVpnMqttSessionCounter) Set(val *MsgVpnMqttSessionCounter) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnMqttSessionCounter) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnMqttSessionCounter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnMqttSessionCounter(val *MsgVpnMqttSessionCounter) *NullableMsgVpnMqttSessionCounter {
	return &NullableMsgVpnMqttSessionCounter{value: val, isSet: true}
}

func (v NullableMsgVpnMqttSessionCounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnMqttSessionCounter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
