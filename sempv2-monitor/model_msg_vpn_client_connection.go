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

// MsgVpnClientConnection struct for MsgVpnClientConnection
type MsgVpnClientConnection struct {
	// The IP address and TCP port on the Client side of the Client Connection.
	ClientAddress *string `json:"clientAddress,omitempty"`
	// The name of the Client.
	ClientName *string `json:"clientName,omitempty"`
	// Indicates whether compression is enabled for the Client Connection.
	Compression *bool `json:"compression,omitempty"`
	// Indicates whether encryption (TLS) is enabled for the Client Connection.
	Encryption *bool `json:"encryption,omitempty"`
	// The number of TCP fast retransmits due to duplicate acknowledgments (ACKs). See RFC 5681 for further details.
	FastRetransmitCount *int32 `json:"fastRetransmitCount,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// The number of bytes currently in the receive queue for the Client Connection.
	RxQueueByteCount *int32 `json:"rxQueueByteCount,omitempty"`
	// The number of TCP segments received from the Client Connection out of order.
	SegmentReceivedOutOfOrderCount *int32 `json:"segmentReceivedOutOfOrderCount,omitempty"`
	// The TCP smoothed round-trip time (SRTT) for the Client Connection, in nanoseconds. See RFC 2988 for further details.
	SmoothedRoundTripTime *int64 `json:"smoothedRoundTripTime,omitempty"`
	// The TCP state of the Client Connection. When fully operational, should be: established. See RFC 793 for further details. The allowed values and their meaning are:  <pre> \"closed\" - No connection state at all. \"listen\" - Waiting for a connection request from any remote TCP and port. \"syn-sent\" - Waiting for a matching connection request after having sent a connection request. \"syn-received\" - Waiting for a confirming connection request acknowledgment after having both received and sent a connection request. \"established\" - An open connection, data received can be delivered to the user. \"close-wait\" - Waiting for a connection termination request from the local user. \"fin-wait-1\" - Waiting for a connection termination request from the remote TCP, or an acknowledgment of the connection termination request previously sent. \"closing\" - Waiting for a connection termination request acknowledgment from the remote TCP. \"last-ack\" - Waiting for an acknowledgment of the connection termination request previously sent to the remote TCP. \"fin-wait-2\" - Waiting for a connection termination request from the remote TCP. \"time-wait\" - Waiting for enough time to pass to be sure the remote TCP received the acknowledgment of its connection termination request. </pre>
	TcpState *string `json:"tcpState,omitempty"`
	// The number of TCP segments retransmitted due to timeout awaiting an acknowledgement (ACK). See RFC 793 for further details.
	TimedRetransmitCount *int32 `json:"timedRetransmitCount,omitempty"`
	// The number of bytes currently in the transmit queue for the Client Connection.
	TxQueueByteCount *int32 `json:"txQueueByteCount,omitempty"`
	// The amount of time in seconds since the Client Connection was established.
	Uptime *int64 `json:"uptime,omitempty"`
}

// NewMsgVpnClientConnection instantiates a new MsgVpnClientConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnClientConnection() *MsgVpnClientConnection {
	this := MsgVpnClientConnection{}
	return &this
}

// NewMsgVpnClientConnectionWithDefaults instantiates a new MsgVpnClientConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnClientConnectionWithDefaults() *MsgVpnClientConnection {
	this := MsgVpnClientConnection{}
	return &this
}

// GetClientAddress returns the ClientAddress field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetClientAddress() string {
	if o == nil || o.ClientAddress == nil {
		var ret string
		return ret
	}
	return *o.ClientAddress
}

// GetClientAddressOk returns a tuple with the ClientAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetClientAddressOk() (*string, bool) {
	if o == nil || o.ClientAddress == nil {
		return nil, false
	}
	return o.ClientAddress, true
}

// HasClientAddress returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasClientAddress() bool {
	if o != nil && o.ClientAddress != nil {
		return true
	}

	return false
}

// SetClientAddress gets a reference to the given string and assigns it to the ClientAddress field.
func (o *MsgVpnClientConnection) SetClientAddress(v string) {
	o.ClientAddress = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *MsgVpnClientConnection) SetClientName(v string) {
	o.ClientName = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetCompression() bool {
	if o == nil || o.Compression == nil {
		var ret bool
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetCompressionOk() (*bool, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasCompression() bool {
	if o != nil && o.Compression != nil {
		return true
	}

	return false
}

// SetCompression gets a reference to the given bool and assigns it to the Compression field.
func (o *MsgVpnClientConnection) SetCompression(v bool) {
	o.Compression = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetEncryption() bool {
	if o == nil || o.Encryption == nil {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetEncryptionOk() (*bool, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasEncryption() bool {
	if o != nil && o.Encryption != nil {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *MsgVpnClientConnection) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetFastRetransmitCount returns the FastRetransmitCount field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetFastRetransmitCount() int32 {
	if o == nil || o.FastRetransmitCount == nil {
		var ret int32
		return ret
	}
	return *o.FastRetransmitCount
}

// GetFastRetransmitCountOk returns a tuple with the FastRetransmitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetFastRetransmitCountOk() (*int32, bool) {
	if o == nil || o.FastRetransmitCount == nil {
		return nil, false
	}
	return o.FastRetransmitCount, true
}

// HasFastRetransmitCount returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasFastRetransmitCount() bool {
	if o != nil && o.FastRetransmitCount != nil {
		return true
	}

	return false
}

// SetFastRetransmitCount gets a reference to the given int32 and assigns it to the FastRetransmitCount field.
func (o *MsgVpnClientConnection) SetFastRetransmitCount(v int32) {
	o.FastRetransmitCount = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnClientConnection) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetRxQueueByteCount returns the RxQueueByteCount field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetRxQueueByteCount() int32 {
	if o == nil || o.RxQueueByteCount == nil {
		var ret int32
		return ret
	}
	return *o.RxQueueByteCount
}

// GetRxQueueByteCountOk returns a tuple with the RxQueueByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetRxQueueByteCountOk() (*int32, bool) {
	if o == nil || o.RxQueueByteCount == nil {
		return nil, false
	}
	return o.RxQueueByteCount, true
}

// HasRxQueueByteCount returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasRxQueueByteCount() bool {
	if o != nil && o.RxQueueByteCount != nil {
		return true
	}

	return false
}

// SetRxQueueByteCount gets a reference to the given int32 and assigns it to the RxQueueByteCount field.
func (o *MsgVpnClientConnection) SetRxQueueByteCount(v int32) {
	o.RxQueueByteCount = &v
}

// GetSegmentReceivedOutOfOrderCount returns the SegmentReceivedOutOfOrderCount field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetSegmentReceivedOutOfOrderCount() int32 {
	if o == nil || o.SegmentReceivedOutOfOrderCount == nil {
		var ret int32
		return ret
	}
	return *o.SegmentReceivedOutOfOrderCount
}

// GetSegmentReceivedOutOfOrderCountOk returns a tuple with the SegmentReceivedOutOfOrderCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetSegmentReceivedOutOfOrderCountOk() (*int32, bool) {
	if o == nil || o.SegmentReceivedOutOfOrderCount == nil {
		return nil, false
	}
	return o.SegmentReceivedOutOfOrderCount, true
}

// HasSegmentReceivedOutOfOrderCount returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasSegmentReceivedOutOfOrderCount() bool {
	if o != nil && o.SegmentReceivedOutOfOrderCount != nil {
		return true
	}

	return false
}

// SetSegmentReceivedOutOfOrderCount gets a reference to the given int32 and assigns it to the SegmentReceivedOutOfOrderCount field.
func (o *MsgVpnClientConnection) SetSegmentReceivedOutOfOrderCount(v int32) {
	o.SegmentReceivedOutOfOrderCount = &v
}

// GetSmoothedRoundTripTime returns the SmoothedRoundTripTime field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetSmoothedRoundTripTime() int64 {
	if o == nil || o.SmoothedRoundTripTime == nil {
		var ret int64
		return ret
	}
	return *o.SmoothedRoundTripTime
}

// GetSmoothedRoundTripTimeOk returns a tuple with the SmoothedRoundTripTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetSmoothedRoundTripTimeOk() (*int64, bool) {
	if o == nil || o.SmoothedRoundTripTime == nil {
		return nil, false
	}
	return o.SmoothedRoundTripTime, true
}

// HasSmoothedRoundTripTime returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasSmoothedRoundTripTime() bool {
	if o != nil && o.SmoothedRoundTripTime != nil {
		return true
	}

	return false
}

// SetSmoothedRoundTripTime gets a reference to the given int64 and assigns it to the SmoothedRoundTripTime field.
func (o *MsgVpnClientConnection) SetSmoothedRoundTripTime(v int64) {
	o.SmoothedRoundTripTime = &v
}

// GetTcpState returns the TcpState field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetTcpState() string {
	if o == nil || o.TcpState == nil {
		var ret string
		return ret
	}
	return *o.TcpState
}

// GetTcpStateOk returns a tuple with the TcpState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetTcpStateOk() (*string, bool) {
	if o == nil || o.TcpState == nil {
		return nil, false
	}
	return o.TcpState, true
}

// HasTcpState returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasTcpState() bool {
	if o != nil && o.TcpState != nil {
		return true
	}

	return false
}

// SetTcpState gets a reference to the given string and assigns it to the TcpState field.
func (o *MsgVpnClientConnection) SetTcpState(v string) {
	o.TcpState = &v
}

// GetTimedRetransmitCount returns the TimedRetransmitCount field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetTimedRetransmitCount() int32 {
	if o == nil || o.TimedRetransmitCount == nil {
		var ret int32
		return ret
	}
	return *o.TimedRetransmitCount
}

// GetTimedRetransmitCountOk returns a tuple with the TimedRetransmitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetTimedRetransmitCountOk() (*int32, bool) {
	if o == nil || o.TimedRetransmitCount == nil {
		return nil, false
	}
	return o.TimedRetransmitCount, true
}

// HasTimedRetransmitCount returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasTimedRetransmitCount() bool {
	if o != nil && o.TimedRetransmitCount != nil {
		return true
	}

	return false
}

// SetTimedRetransmitCount gets a reference to the given int32 and assigns it to the TimedRetransmitCount field.
func (o *MsgVpnClientConnection) SetTimedRetransmitCount(v int32) {
	o.TimedRetransmitCount = &v
}

// GetTxQueueByteCount returns the TxQueueByteCount field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetTxQueueByteCount() int32 {
	if o == nil || o.TxQueueByteCount == nil {
		var ret int32
		return ret
	}
	return *o.TxQueueByteCount
}

// GetTxQueueByteCountOk returns a tuple with the TxQueueByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetTxQueueByteCountOk() (*int32, bool) {
	if o == nil || o.TxQueueByteCount == nil {
		return nil, false
	}
	return o.TxQueueByteCount, true
}

// HasTxQueueByteCount returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasTxQueueByteCount() bool {
	if o != nil && o.TxQueueByteCount != nil {
		return true
	}

	return false
}

// SetTxQueueByteCount gets a reference to the given int32 and assigns it to the TxQueueByteCount field.
func (o *MsgVpnClientConnection) SetTxQueueByteCount(v int32) {
	o.TxQueueByteCount = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *MsgVpnClientConnection) GetUptime() int64 {
	if o == nil || o.Uptime == nil {
		var ret int64
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientConnection) GetUptimeOk() (*int64, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *MsgVpnClientConnection) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given int64 and assigns it to the Uptime field.
func (o *MsgVpnClientConnection) SetUptime(v int64) {
	o.Uptime = &v
}

func (o MsgVpnClientConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientAddress != nil {
		toSerialize["clientAddress"] = o.ClientAddress
	}
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	if o.Encryption != nil {
		toSerialize["encryption"] = o.Encryption
	}
	if o.FastRetransmitCount != nil {
		toSerialize["fastRetransmitCount"] = o.FastRetransmitCount
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.RxQueueByteCount != nil {
		toSerialize["rxQueueByteCount"] = o.RxQueueByteCount
	}
	if o.SegmentReceivedOutOfOrderCount != nil {
		toSerialize["segmentReceivedOutOfOrderCount"] = o.SegmentReceivedOutOfOrderCount
	}
	if o.SmoothedRoundTripTime != nil {
		toSerialize["smoothedRoundTripTime"] = o.SmoothedRoundTripTime
	}
	if o.TcpState != nil {
		toSerialize["tcpState"] = o.TcpState
	}
	if o.TimedRetransmitCount != nil {
		toSerialize["timedRetransmitCount"] = o.TimedRetransmitCount
	}
	if o.TxQueueByteCount != nil {
		toSerialize["txQueueByteCount"] = o.TxQueueByteCount
	}
	if o.Uptime != nil {
		toSerialize["uptime"] = o.Uptime
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnClientConnection struct {
	value *MsgVpnClientConnection
	isSet bool
}

func (v NullableMsgVpnClientConnection) Get() *MsgVpnClientConnection {
	return v.value
}

func (v *NullableMsgVpnClientConnection) Set(val *MsgVpnClientConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnClientConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnClientConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnClientConnection(val *MsgVpnClientConnection) *NullableMsgVpnClientConnection {
	return &NullableMsgVpnClientConnection{value: val, isSet: true}
}

func (v NullableMsgVpnClientConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnClientConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
