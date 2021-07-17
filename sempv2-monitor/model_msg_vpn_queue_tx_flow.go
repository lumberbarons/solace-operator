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

// MsgVpnQueueTxFlow struct for MsgVpnQueueTxFlow
type MsgVpnQueueTxFlow struct {
	// The number of guaranteed messages delivered and acknowledged by the consumer.
	AckedMsgCount *int64 `json:"ackedMsgCount,omitempty"`
	// The timestamp of when the bound Flow became active. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	ActivationTime *int32 `json:"activationTime,omitempty"`
	// The activity state of the Flow. The allowed values and their meaning are:  <pre> \"active-browser\" - The Flow is active as a browser. \"active-consumer\" - The Flow is active as a consumer. \"inactive\" - The Flow is inactive. </pre>
	ActivityState *string `json:"activityState,omitempty"`
	// The state of updating the consumer with the Flow activity. The allowed values and their meaning are:  <pre> \"in-progress\" - The Flow is in the process of updating the client with its activity state. \"synchronized\" - The Flow has updated the client with its activity state. \"not-requested\" - The Flow has not been requested by the client to provide activity updates. </pre>
	ActivityUpdateState *string `json:"activityUpdateState,omitempty"`
	// The timestamp of when the Flow bound to the Queue. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	BindTime *int32 `json:"bindTime,omitempty"`
	// The name of the Client.
	ClientName *string `json:"clientName,omitempty"`
	// Indicates whether redelivery requests can be received as negative acknowledgements (NACKs) from the consumer. Applicable only to REST consumers.
	ConsumerRedeliveryRequestAllowed *bool `json:"consumerRedeliveryRequestAllowed,omitempty"`
	// The number of guaranteed messages that used cut-through delivery and are acknowledged by the consumer.
	CutThroughAckedMsgCount *int64 `json:"cutThroughAckedMsgCount,omitempty"`
	// The delivery state of the Flow. The allowed values and their meaning are:  <pre> \"closed\" - The Flow is unbound. \"opened\" - The Flow is bound but inactive. \"unbinding\" - The Flow received an unbind request. \"handshaking\" - The Flow is handshaking to become active. \"deliver-cut-through\" - The Flow is streaming messages using direct+guaranteed delivery. \"deliver-from-input-stream\" - The Flow is streaming messages using guaranteed delivery. \"deliver-from-memory\" - The Flow throttled causing message delivery from memory (RAM). \"deliver-from-spool\" - The Flow stalled causing message delivery from spool (ADB or disk). </pre>
	DeliveryState *string `json:"deliveryState,omitempty"`
	// The identifier (ID) of the Flow.
	FlowId *int64 `json:"flowId,omitempty"`
	// The highest identifier (ID) of message transmitted and waiting for acknowledgement.
	HighestAckPendingMsgId *int64 `json:"highestAckPendingMsgId,omitempty"`
	// The identifier (ID) of the last message transmitted and acknowledged by the consumer.
	LastAckedMsgId *int64 `json:"lastAckedMsgId,omitempty"`
	// The identifier (ID) of the last message examined by the Flow selector.
	LastSelectorExaminedMsgId *int64 `json:"lastSelectorExaminedMsgId,omitempty"`
	// The lowest identifier (ID) of message transmitted and waiting for acknowledgement.
	LowestAckPendingMsgId *int64 `json:"lowestAckPendingMsgId,omitempty"`
	// The number of guaranteed messages that exceeded the maximum number of delivered unacknowledged messages.
	MaxUnackedMsgsExceededMsgCount *int64 `json:"maxUnackedMsgsExceededMsgCount,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// Indicates whether not to deliver messages to a consumer that published them.
	NoLocalDelivery *bool `json:"noLocalDelivery,omitempty"`
	// The name of the Queue.
	QueueName *string `json:"queueName,omitempty"`
	// The number of guaranteed messages that were redelivered.
	RedeliveredMsgCount *int64 `json:"redeliveredMsgCount,omitempty"`
	// The number of consumer requests via negative acknowledgements (NACKs) to redeliver guaranteed messages.
	RedeliveryRequestCount *int64 `json:"redeliveryRequestCount,omitempty"`
	// The value of the Flow selector.
	Selector *string `json:"selector,omitempty"`
	// The number of guaranteed messages examined by the Flow selector.
	SelectorExaminedMsgCount *int64 `json:"selectorExaminedMsgCount,omitempty"`
	// The number of guaranteed messages for which the Flow selector matched.
	SelectorMatchedMsgCount *int64 `json:"selectorMatchedMsgCount,omitempty"`
	// The number of guaranteed messages for which the Flow selector did not match.
	SelectorNotMatchedMsgCount *int64 `json:"selectorNotMatchedMsgCount,omitempty"`
	// The name of the Transacted Session for the Flow.
	SessionName *string `json:"sessionName,omitempty"`
	// The number of guaranteed messages that used store and forward delivery and are acknowledged by the consumer.
	StoreAndForwardAckedMsgCount *int64 `json:"storeAndForwardAckedMsgCount,omitempty"`
	// The number of guaranteed messages that were retransmitted at the transport layer as part of a single delivery attempt. Available since 2.18.
	TransportRetransmitMsgCount *int64 `json:"transportRetransmitMsgCount,omitempty"`
	// The number of guaranteed messages delivered but not yet acknowledged by the consumer.
	UnackedMsgCount *int64 `json:"unackedMsgCount,omitempty"`
	// The number of guaranteed messages using the available window size.
	UsedWindowSize *int64 `json:"usedWindowSize,omitempty"`
	// The number of times the window for guaranteed messages was filled and closed before an acknowledgement was received.
	WindowClosedCount *int64 `json:"windowClosedCount,omitempty"`
	// The number of outstanding guaranteed messages that can be transmitted over the Flow before an acknowledgement is received.
	WindowSize *int64 `json:"windowSize,omitempty"`
}

// NewMsgVpnQueueTxFlow instantiates a new MsgVpnQueueTxFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnQueueTxFlow() *MsgVpnQueueTxFlow {
	this := MsgVpnQueueTxFlow{}
	return &this
}

// NewMsgVpnQueueTxFlowWithDefaults instantiates a new MsgVpnQueueTxFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnQueueTxFlowWithDefaults() *MsgVpnQueueTxFlow {
	this := MsgVpnQueueTxFlow{}
	return &this
}

// GetAckedMsgCount returns the AckedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetAckedMsgCount() int64 {
	if o == nil || o.AckedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.AckedMsgCount
}

// GetAckedMsgCountOk returns a tuple with the AckedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetAckedMsgCountOk() (*int64, bool) {
	if o == nil || o.AckedMsgCount == nil {
		return nil, false
	}
	return o.AckedMsgCount, true
}

// HasAckedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasAckedMsgCount() bool {
	if o != nil && o.AckedMsgCount != nil {
		return true
	}

	return false
}

// SetAckedMsgCount gets a reference to the given int64 and assigns it to the AckedMsgCount field.
func (o *MsgVpnQueueTxFlow) SetAckedMsgCount(v int64) {
	o.AckedMsgCount = &v
}

// GetActivationTime returns the ActivationTime field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetActivationTime() int32 {
	if o == nil || o.ActivationTime == nil {
		var ret int32
		return ret
	}
	return *o.ActivationTime
}

// GetActivationTimeOk returns a tuple with the ActivationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetActivationTimeOk() (*int32, bool) {
	if o == nil || o.ActivationTime == nil {
		return nil, false
	}
	return o.ActivationTime, true
}

// HasActivationTime returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasActivationTime() bool {
	if o != nil && o.ActivationTime != nil {
		return true
	}

	return false
}

// SetActivationTime gets a reference to the given int32 and assigns it to the ActivationTime field.
func (o *MsgVpnQueueTxFlow) SetActivationTime(v int32) {
	o.ActivationTime = &v
}

// GetActivityState returns the ActivityState field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetActivityState() string {
	if o == nil || o.ActivityState == nil {
		var ret string
		return ret
	}
	return *o.ActivityState
}

// GetActivityStateOk returns a tuple with the ActivityState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetActivityStateOk() (*string, bool) {
	if o == nil || o.ActivityState == nil {
		return nil, false
	}
	return o.ActivityState, true
}

// HasActivityState returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasActivityState() bool {
	if o != nil && o.ActivityState != nil {
		return true
	}

	return false
}

// SetActivityState gets a reference to the given string and assigns it to the ActivityState field.
func (o *MsgVpnQueueTxFlow) SetActivityState(v string) {
	o.ActivityState = &v
}

// GetActivityUpdateState returns the ActivityUpdateState field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetActivityUpdateState() string {
	if o == nil || o.ActivityUpdateState == nil {
		var ret string
		return ret
	}
	return *o.ActivityUpdateState
}

// GetActivityUpdateStateOk returns a tuple with the ActivityUpdateState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetActivityUpdateStateOk() (*string, bool) {
	if o == nil || o.ActivityUpdateState == nil {
		return nil, false
	}
	return o.ActivityUpdateState, true
}

// HasActivityUpdateState returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasActivityUpdateState() bool {
	if o != nil && o.ActivityUpdateState != nil {
		return true
	}

	return false
}

// SetActivityUpdateState gets a reference to the given string and assigns it to the ActivityUpdateState field.
func (o *MsgVpnQueueTxFlow) SetActivityUpdateState(v string) {
	o.ActivityUpdateState = &v
}

// GetBindTime returns the BindTime field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetBindTime() int32 {
	if o == nil || o.BindTime == nil {
		var ret int32
		return ret
	}
	return *o.BindTime
}

// GetBindTimeOk returns a tuple with the BindTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetBindTimeOk() (*int32, bool) {
	if o == nil || o.BindTime == nil {
		return nil, false
	}
	return o.BindTime, true
}

// HasBindTime returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasBindTime() bool {
	if o != nil && o.BindTime != nil {
		return true
	}

	return false
}

// SetBindTime gets a reference to the given int32 and assigns it to the BindTime field.
func (o *MsgVpnQueueTxFlow) SetBindTime(v int32) {
	o.BindTime = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *MsgVpnQueueTxFlow) SetClientName(v string) {
	o.ClientName = &v
}

// GetConsumerRedeliveryRequestAllowed returns the ConsumerRedeliveryRequestAllowed field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetConsumerRedeliveryRequestAllowed() bool {
	if o == nil || o.ConsumerRedeliveryRequestAllowed == nil {
		var ret bool
		return ret
	}
	return *o.ConsumerRedeliveryRequestAllowed
}

// GetConsumerRedeliveryRequestAllowedOk returns a tuple with the ConsumerRedeliveryRequestAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetConsumerRedeliveryRequestAllowedOk() (*bool, bool) {
	if o == nil || o.ConsumerRedeliveryRequestAllowed == nil {
		return nil, false
	}
	return o.ConsumerRedeliveryRequestAllowed, true
}

// HasConsumerRedeliveryRequestAllowed returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasConsumerRedeliveryRequestAllowed() bool {
	if o != nil && o.ConsumerRedeliveryRequestAllowed != nil {
		return true
	}

	return false
}

// SetConsumerRedeliveryRequestAllowed gets a reference to the given bool and assigns it to the ConsumerRedeliveryRequestAllowed field.
func (o *MsgVpnQueueTxFlow) SetConsumerRedeliveryRequestAllowed(v bool) {
	o.ConsumerRedeliveryRequestAllowed = &v
}

// GetCutThroughAckedMsgCount returns the CutThroughAckedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetCutThroughAckedMsgCount() int64 {
	if o == nil || o.CutThroughAckedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.CutThroughAckedMsgCount
}

// GetCutThroughAckedMsgCountOk returns a tuple with the CutThroughAckedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetCutThroughAckedMsgCountOk() (*int64, bool) {
	if o == nil || o.CutThroughAckedMsgCount == nil {
		return nil, false
	}
	return o.CutThroughAckedMsgCount, true
}

// HasCutThroughAckedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasCutThroughAckedMsgCount() bool {
	if o != nil && o.CutThroughAckedMsgCount != nil {
		return true
	}

	return false
}

// SetCutThroughAckedMsgCount gets a reference to the given int64 and assigns it to the CutThroughAckedMsgCount field.
func (o *MsgVpnQueueTxFlow) SetCutThroughAckedMsgCount(v int64) {
	o.CutThroughAckedMsgCount = &v
}

// GetDeliveryState returns the DeliveryState field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetDeliveryState() string {
	if o == nil || o.DeliveryState == nil {
		var ret string
		return ret
	}
	return *o.DeliveryState
}

// GetDeliveryStateOk returns a tuple with the DeliveryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetDeliveryStateOk() (*string, bool) {
	if o == nil || o.DeliveryState == nil {
		return nil, false
	}
	return o.DeliveryState, true
}

// HasDeliveryState returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasDeliveryState() bool {
	if o != nil && o.DeliveryState != nil {
		return true
	}

	return false
}

// SetDeliveryState gets a reference to the given string and assigns it to the DeliveryState field.
func (o *MsgVpnQueueTxFlow) SetDeliveryState(v string) {
	o.DeliveryState = &v
}

// GetFlowId returns the FlowId field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetFlowId() int64 {
	if o == nil || o.FlowId == nil {
		var ret int64
		return ret
	}
	return *o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetFlowIdOk() (*int64, bool) {
	if o == nil || o.FlowId == nil {
		return nil, false
	}
	return o.FlowId, true
}

// HasFlowId returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasFlowId() bool {
	if o != nil && o.FlowId != nil {
		return true
	}

	return false
}

// SetFlowId gets a reference to the given int64 and assigns it to the FlowId field.
func (o *MsgVpnQueueTxFlow) SetFlowId(v int64) {
	o.FlowId = &v
}

// GetHighestAckPendingMsgId returns the HighestAckPendingMsgId field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetHighestAckPendingMsgId() int64 {
	if o == nil || o.HighestAckPendingMsgId == nil {
		var ret int64
		return ret
	}
	return *o.HighestAckPendingMsgId
}

// GetHighestAckPendingMsgIdOk returns a tuple with the HighestAckPendingMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetHighestAckPendingMsgIdOk() (*int64, bool) {
	if o == nil || o.HighestAckPendingMsgId == nil {
		return nil, false
	}
	return o.HighestAckPendingMsgId, true
}

// HasHighestAckPendingMsgId returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasHighestAckPendingMsgId() bool {
	if o != nil && o.HighestAckPendingMsgId != nil {
		return true
	}

	return false
}

// SetHighestAckPendingMsgId gets a reference to the given int64 and assigns it to the HighestAckPendingMsgId field.
func (o *MsgVpnQueueTxFlow) SetHighestAckPendingMsgId(v int64) {
	o.HighestAckPendingMsgId = &v
}

// GetLastAckedMsgId returns the LastAckedMsgId field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetLastAckedMsgId() int64 {
	if o == nil || o.LastAckedMsgId == nil {
		var ret int64
		return ret
	}
	return *o.LastAckedMsgId
}

// GetLastAckedMsgIdOk returns a tuple with the LastAckedMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetLastAckedMsgIdOk() (*int64, bool) {
	if o == nil || o.LastAckedMsgId == nil {
		return nil, false
	}
	return o.LastAckedMsgId, true
}

// HasLastAckedMsgId returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasLastAckedMsgId() bool {
	if o != nil && o.LastAckedMsgId != nil {
		return true
	}

	return false
}

// SetLastAckedMsgId gets a reference to the given int64 and assigns it to the LastAckedMsgId field.
func (o *MsgVpnQueueTxFlow) SetLastAckedMsgId(v int64) {
	o.LastAckedMsgId = &v
}

// GetLastSelectorExaminedMsgId returns the LastSelectorExaminedMsgId field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetLastSelectorExaminedMsgId() int64 {
	if o == nil || o.LastSelectorExaminedMsgId == nil {
		var ret int64
		return ret
	}
	return *o.LastSelectorExaminedMsgId
}

// GetLastSelectorExaminedMsgIdOk returns a tuple with the LastSelectorExaminedMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetLastSelectorExaminedMsgIdOk() (*int64, bool) {
	if o == nil || o.LastSelectorExaminedMsgId == nil {
		return nil, false
	}
	return o.LastSelectorExaminedMsgId, true
}

// HasLastSelectorExaminedMsgId returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasLastSelectorExaminedMsgId() bool {
	if o != nil && o.LastSelectorExaminedMsgId != nil {
		return true
	}

	return false
}

// SetLastSelectorExaminedMsgId gets a reference to the given int64 and assigns it to the LastSelectorExaminedMsgId field.
func (o *MsgVpnQueueTxFlow) SetLastSelectorExaminedMsgId(v int64) {
	o.LastSelectorExaminedMsgId = &v
}

// GetLowestAckPendingMsgId returns the LowestAckPendingMsgId field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetLowestAckPendingMsgId() int64 {
	if o == nil || o.LowestAckPendingMsgId == nil {
		var ret int64
		return ret
	}
	return *o.LowestAckPendingMsgId
}

// GetLowestAckPendingMsgIdOk returns a tuple with the LowestAckPendingMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetLowestAckPendingMsgIdOk() (*int64, bool) {
	if o == nil || o.LowestAckPendingMsgId == nil {
		return nil, false
	}
	return o.LowestAckPendingMsgId, true
}

// HasLowestAckPendingMsgId returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasLowestAckPendingMsgId() bool {
	if o != nil && o.LowestAckPendingMsgId != nil {
		return true
	}

	return false
}

// SetLowestAckPendingMsgId gets a reference to the given int64 and assigns it to the LowestAckPendingMsgId field.
func (o *MsgVpnQueueTxFlow) SetLowestAckPendingMsgId(v int64) {
	o.LowestAckPendingMsgId = &v
}

// GetMaxUnackedMsgsExceededMsgCount returns the MaxUnackedMsgsExceededMsgCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetMaxUnackedMsgsExceededMsgCount() int64 {
	if o == nil || o.MaxUnackedMsgsExceededMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxUnackedMsgsExceededMsgCount
}

// GetMaxUnackedMsgsExceededMsgCountOk returns a tuple with the MaxUnackedMsgsExceededMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetMaxUnackedMsgsExceededMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxUnackedMsgsExceededMsgCount == nil {
		return nil, false
	}
	return o.MaxUnackedMsgsExceededMsgCount, true
}

// HasMaxUnackedMsgsExceededMsgCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasMaxUnackedMsgsExceededMsgCount() bool {
	if o != nil && o.MaxUnackedMsgsExceededMsgCount != nil {
		return true
	}

	return false
}

// SetMaxUnackedMsgsExceededMsgCount gets a reference to the given int64 and assigns it to the MaxUnackedMsgsExceededMsgCount field.
func (o *MsgVpnQueueTxFlow) SetMaxUnackedMsgsExceededMsgCount(v int64) {
	o.MaxUnackedMsgsExceededMsgCount = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnQueueTxFlow) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetNoLocalDelivery returns the NoLocalDelivery field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetNoLocalDelivery() bool {
	if o == nil || o.NoLocalDelivery == nil {
		var ret bool
		return ret
	}
	return *o.NoLocalDelivery
}

// GetNoLocalDeliveryOk returns a tuple with the NoLocalDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetNoLocalDeliveryOk() (*bool, bool) {
	if o == nil || o.NoLocalDelivery == nil {
		return nil, false
	}
	return o.NoLocalDelivery, true
}

// HasNoLocalDelivery returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasNoLocalDelivery() bool {
	if o != nil && o.NoLocalDelivery != nil {
		return true
	}

	return false
}

// SetNoLocalDelivery gets a reference to the given bool and assigns it to the NoLocalDelivery field.
func (o *MsgVpnQueueTxFlow) SetNoLocalDelivery(v bool) {
	o.NoLocalDelivery = &v
}

// GetQueueName returns the QueueName field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetQueueName() string {
	if o == nil || o.QueueName == nil {
		var ret string
		return ret
	}
	return *o.QueueName
}

// GetQueueNameOk returns a tuple with the QueueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetQueueNameOk() (*string, bool) {
	if o == nil || o.QueueName == nil {
		return nil, false
	}
	return o.QueueName, true
}

// HasQueueName returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasQueueName() bool {
	if o != nil && o.QueueName != nil {
		return true
	}

	return false
}

// SetQueueName gets a reference to the given string and assigns it to the QueueName field.
func (o *MsgVpnQueueTxFlow) SetQueueName(v string) {
	o.QueueName = &v
}

// GetRedeliveredMsgCount returns the RedeliveredMsgCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetRedeliveredMsgCount() int64 {
	if o == nil || o.RedeliveredMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RedeliveredMsgCount
}

// GetRedeliveredMsgCountOk returns a tuple with the RedeliveredMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetRedeliveredMsgCountOk() (*int64, bool) {
	if o == nil || o.RedeliveredMsgCount == nil {
		return nil, false
	}
	return o.RedeliveredMsgCount, true
}

// HasRedeliveredMsgCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasRedeliveredMsgCount() bool {
	if o != nil && o.RedeliveredMsgCount != nil {
		return true
	}

	return false
}

// SetRedeliveredMsgCount gets a reference to the given int64 and assigns it to the RedeliveredMsgCount field.
func (o *MsgVpnQueueTxFlow) SetRedeliveredMsgCount(v int64) {
	o.RedeliveredMsgCount = &v
}

// GetRedeliveryRequestCount returns the RedeliveryRequestCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetRedeliveryRequestCount() int64 {
	if o == nil || o.RedeliveryRequestCount == nil {
		var ret int64
		return ret
	}
	return *o.RedeliveryRequestCount
}

// GetRedeliveryRequestCountOk returns a tuple with the RedeliveryRequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetRedeliveryRequestCountOk() (*int64, bool) {
	if o == nil || o.RedeliveryRequestCount == nil {
		return nil, false
	}
	return o.RedeliveryRequestCount, true
}

// HasRedeliveryRequestCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasRedeliveryRequestCount() bool {
	if o != nil && o.RedeliveryRequestCount != nil {
		return true
	}

	return false
}

// SetRedeliveryRequestCount gets a reference to the given int64 and assigns it to the RedeliveryRequestCount field.
func (o *MsgVpnQueueTxFlow) SetRedeliveryRequestCount(v int64) {
	o.RedeliveryRequestCount = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *MsgVpnQueueTxFlow) SetSelector(v string) {
	o.Selector = &v
}

// GetSelectorExaminedMsgCount returns the SelectorExaminedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetSelectorExaminedMsgCount() int64 {
	if o == nil || o.SelectorExaminedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SelectorExaminedMsgCount
}

// GetSelectorExaminedMsgCountOk returns a tuple with the SelectorExaminedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetSelectorExaminedMsgCountOk() (*int64, bool) {
	if o == nil || o.SelectorExaminedMsgCount == nil {
		return nil, false
	}
	return o.SelectorExaminedMsgCount, true
}

// HasSelectorExaminedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasSelectorExaminedMsgCount() bool {
	if o != nil && o.SelectorExaminedMsgCount != nil {
		return true
	}

	return false
}

// SetSelectorExaminedMsgCount gets a reference to the given int64 and assigns it to the SelectorExaminedMsgCount field.
func (o *MsgVpnQueueTxFlow) SetSelectorExaminedMsgCount(v int64) {
	o.SelectorExaminedMsgCount = &v
}

// GetSelectorMatchedMsgCount returns the SelectorMatchedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetSelectorMatchedMsgCount() int64 {
	if o == nil || o.SelectorMatchedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SelectorMatchedMsgCount
}

// GetSelectorMatchedMsgCountOk returns a tuple with the SelectorMatchedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetSelectorMatchedMsgCountOk() (*int64, bool) {
	if o == nil || o.SelectorMatchedMsgCount == nil {
		return nil, false
	}
	return o.SelectorMatchedMsgCount, true
}

// HasSelectorMatchedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasSelectorMatchedMsgCount() bool {
	if o != nil && o.SelectorMatchedMsgCount != nil {
		return true
	}

	return false
}

// SetSelectorMatchedMsgCount gets a reference to the given int64 and assigns it to the SelectorMatchedMsgCount field.
func (o *MsgVpnQueueTxFlow) SetSelectorMatchedMsgCount(v int64) {
	o.SelectorMatchedMsgCount = &v
}

// GetSelectorNotMatchedMsgCount returns the SelectorNotMatchedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetSelectorNotMatchedMsgCount() int64 {
	if o == nil || o.SelectorNotMatchedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SelectorNotMatchedMsgCount
}

// GetSelectorNotMatchedMsgCountOk returns a tuple with the SelectorNotMatchedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetSelectorNotMatchedMsgCountOk() (*int64, bool) {
	if o == nil || o.SelectorNotMatchedMsgCount == nil {
		return nil, false
	}
	return o.SelectorNotMatchedMsgCount, true
}

// HasSelectorNotMatchedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasSelectorNotMatchedMsgCount() bool {
	if o != nil && o.SelectorNotMatchedMsgCount != nil {
		return true
	}

	return false
}

// SetSelectorNotMatchedMsgCount gets a reference to the given int64 and assigns it to the SelectorNotMatchedMsgCount field.
func (o *MsgVpnQueueTxFlow) SetSelectorNotMatchedMsgCount(v int64) {
	o.SelectorNotMatchedMsgCount = &v
}

// GetSessionName returns the SessionName field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetSessionName() string {
	if o == nil || o.SessionName == nil {
		var ret string
		return ret
	}
	return *o.SessionName
}

// GetSessionNameOk returns a tuple with the SessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetSessionNameOk() (*string, bool) {
	if o == nil || o.SessionName == nil {
		return nil, false
	}
	return o.SessionName, true
}

// HasSessionName returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasSessionName() bool {
	if o != nil && o.SessionName != nil {
		return true
	}

	return false
}

// SetSessionName gets a reference to the given string and assigns it to the SessionName field.
func (o *MsgVpnQueueTxFlow) SetSessionName(v string) {
	o.SessionName = &v
}

// GetStoreAndForwardAckedMsgCount returns the StoreAndForwardAckedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetStoreAndForwardAckedMsgCount() int64 {
	if o == nil || o.StoreAndForwardAckedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.StoreAndForwardAckedMsgCount
}

// GetStoreAndForwardAckedMsgCountOk returns a tuple with the StoreAndForwardAckedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetStoreAndForwardAckedMsgCountOk() (*int64, bool) {
	if o == nil || o.StoreAndForwardAckedMsgCount == nil {
		return nil, false
	}
	return o.StoreAndForwardAckedMsgCount, true
}

// HasStoreAndForwardAckedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasStoreAndForwardAckedMsgCount() bool {
	if o != nil && o.StoreAndForwardAckedMsgCount != nil {
		return true
	}

	return false
}

// SetStoreAndForwardAckedMsgCount gets a reference to the given int64 and assigns it to the StoreAndForwardAckedMsgCount field.
func (o *MsgVpnQueueTxFlow) SetStoreAndForwardAckedMsgCount(v int64) {
	o.StoreAndForwardAckedMsgCount = &v
}

// GetTransportRetransmitMsgCount returns the TransportRetransmitMsgCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetTransportRetransmitMsgCount() int64 {
	if o == nil || o.TransportRetransmitMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.TransportRetransmitMsgCount
}

// GetTransportRetransmitMsgCountOk returns a tuple with the TransportRetransmitMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetTransportRetransmitMsgCountOk() (*int64, bool) {
	if o == nil || o.TransportRetransmitMsgCount == nil {
		return nil, false
	}
	return o.TransportRetransmitMsgCount, true
}

// HasTransportRetransmitMsgCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasTransportRetransmitMsgCount() bool {
	if o != nil && o.TransportRetransmitMsgCount != nil {
		return true
	}

	return false
}

// SetTransportRetransmitMsgCount gets a reference to the given int64 and assigns it to the TransportRetransmitMsgCount field.
func (o *MsgVpnQueueTxFlow) SetTransportRetransmitMsgCount(v int64) {
	o.TransportRetransmitMsgCount = &v
}

// GetUnackedMsgCount returns the UnackedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetUnackedMsgCount() int64 {
	if o == nil || o.UnackedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.UnackedMsgCount
}

// GetUnackedMsgCountOk returns a tuple with the UnackedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetUnackedMsgCountOk() (*int64, bool) {
	if o == nil || o.UnackedMsgCount == nil {
		return nil, false
	}
	return o.UnackedMsgCount, true
}

// HasUnackedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasUnackedMsgCount() bool {
	if o != nil && o.UnackedMsgCount != nil {
		return true
	}

	return false
}

// SetUnackedMsgCount gets a reference to the given int64 and assigns it to the UnackedMsgCount field.
func (o *MsgVpnQueueTxFlow) SetUnackedMsgCount(v int64) {
	o.UnackedMsgCount = &v
}

// GetUsedWindowSize returns the UsedWindowSize field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetUsedWindowSize() int64 {
	if o == nil || o.UsedWindowSize == nil {
		var ret int64
		return ret
	}
	return *o.UsedWindowSize
}

// GetUsedWindowSizeOk returns a tuple with the UsedWindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetUsedWindowSizeOk() (*int64, bool) {
	if o == nil || o.UsedWindowSize == nil {
		return nil, false
	}
	return o.UsedWindowSize, true
}

// HasUsedWindowSize returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasUsedWindowSize() bool {
	if o != nil && o.UsedWindowSize != nil {
		return true
	}

	return false
}

// SetUsedWindowSize gets a reference to the given int64 and assigns it to the UsedWindowSize field.
func (o *MsgVpnQueueTxFlow) SetUsedWindowSize(v int64) {
	o.UsedWindowSize = &v
}

// GetWindowClosedCount returns the WindowClosedCount field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetWindowClosedCount() int64 {
	if o == nil || o.WindowClosedCount == nil {
		var ret int64
		return ret
	}
	return *o.WindowClosedCount
}

// GetWindowClosedCountOk returns a tuple with the WindowClosedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetWindowClosedCountOk() (*int64, bool) {
	if o == nil || o.WindowClosedCount == nil {
		return nil, false
	}
	return o.WindowClosedCount, true
}

// HasWindowClosedCount returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasWindowClosedCount() bool {
	if o != nil && o.WindowClosedCount != nil {
		return true
	}

	return false
}

// SetWindowClosedCount gets a reference to the given int64 and assigns it to the WindowClosedCount field.
func (o *MsgVpnQueueTxFlow) SetWindowClosedCount(v int64) {
	o.WindowClosedCount = &v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *MsgVpnQueueTxFlow) GetWindowSize() int64 {
	if o == nil || o.WindowSize == nil {
		var ret int64
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnQueueTxFlow) GetWindowSizeOk() (*int64, bool) {
	if o == nil || o.WindowSize == nil {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *MsgVpnQueueTxFlow) HasWindowSize() bool {
	if o != nil && o.WindowSize != nil {
		return true
	}

	return false
}

// SetWindowSize gets a reference to the given int64 and assigns it to the WindowSize field.
func (o *MsgVpnQueueTxFlow) SetWindowSize(v int64) {
	o.WindowSize = &v
}

func (o MsgVpnQueueTxFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AckedMsgCount != nil {
		toSerialize["ackedMsgCount"] = o.AckedMsgCount
	}
	if o.ActivationTime != nil {
		toSerialize["activationTime"] = o.ActivationTime
	}
	if o.ActivityState != nil {
		toSerialize["activityState"] = o.ActivityState
	}
	if o.ActivityUpdateState != nil {
		toSerialize["activityUpdateState"] = o.ActivityUpdateState
	}
	if o.BindTime != nil {
		toSerialize["bindTime"] = o.BindTime
	}
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.ConsumerRedeliveryRequestAllowed != nil {
		toSerialize["consumerRedeliveryRequestAllowed"] = o.ConsumerRedeliveryRequestAllowed
	}
	if o.CutThroughAckedMsgCount != nil {
		toSerialize["cutThroughAckedMsgCount"] = o.CutThroughAckedMsgCount
	}
	if o.DeliveryState != nil {
		toSerialize["deliveryState"] = o.DeliveryState
	}
	if o.FlowId != nil {
		toSerialize["flowId"] = o.FlowId
	}
	if o.HighestAckPendingMsgId != nil {
		toSerialize["highestAckPendingMsgId"] = o.HighestAckPendingMsgId
	}
	if o.LastAckedMsgId != nil {
		toSerialize["lastAckedMsgId"] = o.LastAckedMsgId
	}
	if o.LastSelectorExaminedMsgId != nil {
		toSerialize["lastSelectorExaminedMsgId"] = o.LastSelectorExaminedMsgId
	}
	if o.LowestAckPendingMsgId != nil {
		toSerialize["lowestAckPendingMsgId"] = o.LowestAckPendingMsgId
	}
	if o.MaxUnackedMsgsExceededMsgCount != nil {
		toSerialize["maxUnackedMsgsExceededMsgCount"] = o.MaxUnackedMsgsExceededMsgCount
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.NoLocalDelivery != nil {
		toSerialize["noLocalDelivery"] = o.NoLocalDelivery
	}
	if o.QueueName != nil {
		toSerialize["queueName"] = o.QueueName
	}
	if o.RedeliveredMsgCount != nil {
		toSerialize["redeliveredMsgCount"] = o.RedeliveredMsgCount
	}
	if o.RedeliveryRequestCount != nil {
		toSerialize["redeliveryRequestCount"] = o.RedeliveryRequestCount
	}
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	if o.SelectorExaminedMsgCount != nil {
		toSerialize["selectorExaminedMsgCount"] = o.SelectorExaminedMsgCount
	}
	if o.SelectorMatchedMsgCount != nil {
		toSerialize["selectorMatchedMsgCount"] = o.SelectorMatchedMsgCount
	}
	if o.SelectorNotMatchedMsgCount != nil {
		toSerialize["selectorNotMatchedMsgCount"] = o.SelectorNotMatchedMsgCount
	}
	if o.SessionName != nil {
		toSerialize["sessionName"] = o.SessionName
	}
	if o.StoreAndForwardAckedMsgCount != nil {
		toSerialize["storeAndForwardAckedMsgCount"] = o.StoreAndForwardAckedMsgCount
	}
	if o.TransportRetransmitMsgCount != nil {
		toSerialize["transportRetransmitMsgCount"] = o.TransportRetransmitMsgCount
	}
	if o.UnackedMsgCount != nil {
		toSerialize["unackedMsgCount"] = o.UnackedMsgCount
	}
	if o.UsedWindowSize != nil {
		toSerialize["usedWindowSize"] = o.UsedWindowSize
	}
	if o.WindowClosedCount != nil {
		toSerialize["windowClosedCount"] = o.WindowClosedCount
	}
	if o.WindowSize != nil {
		toSerialize["windowSize"] = o.WindowSize
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnQueueTxFlow struct {
	value *MsgVpnQueueTxFlow
	isSet bool
}

func (v NullableMsgVpnQueueTxFlow) Get() *MsgVpnQueueTxFlow {
	return v.value
}

func (v *NullableMsgVpnQueueTxFlow) Set(val *MsgVpnQueueTxFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnQueueTxFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnQueueTxFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnQueueTxFlow(val *MsgVpnQueueTxFlow) *NullableMsgVpnQueueTxFlow {
	return &NullableMsgVpnQueueTxFlow{value: val, isSet: true}
}

func (v NullableMsgVpnQueueTxFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnQueueTxFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
