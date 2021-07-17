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

// MsgVpnClientRxFlow struct for MsgVpnClientRxFlow
type MsgVpnClientRxFlow struct {
	// The name of the Client.
	ClientName *string `json:"clientName,omitempty"`
	// The timestamp of when the Flow from the Client connected.
	ConnectTime *int32 `json:"connectTime,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to a destination group error.
	DestinationGroupErrorDiscardedMsgCount *int64 `json:"destinationGroupErrorDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to being a duplicate.
	DuplicateDiscardedMsgCount *int64 `json:"duplicateDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to an eligible endpoint destination being disabled.
	EndpointDisabledDiscardedMsgCount *int64 `json:"endpointDisabledDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to an eligible endpoint destination having its maximum message spool usage exceeded.
	EndpointUsageExceededDiscardedMsgCount *int64 `json:"endpointUsageExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to errors being detected.
	ErroredDiscardedMsgCount *int64 `json:"erroredDiscardedMsgCount,omitempty"`
	// The identifier (ID) of the flow.
	FlowId *int64 `json:"flowId,omitempty"`
	// The name of the Flow.
	FlowName *string `json:"flowName,omitempty"`
	// The number of guaranteed messages from the Flow.
	GuaranteedMsgCount *int64 `json:"guaranteedMsgCount,omitempty"`
	// The identifier (ID) of the last message received on the Flow.
	LastRxMsgId *int64 `json:"lastRxMsgId,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to the maximum number of messages allowed on the broker being exceeded.
	LocalMsgCountExceededDiscardedMsgCount *int64 `json:"localMsgCountExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to congestion of low priority messages.
	LowPriorityMsgCongestionDiscardedMsgCount *int64 `json:"lowPriorityMsgCongestionDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to the maximum allowed message size being exceeded.
	MaxMsgSizeExceededDiscardedMsgCount *int64 `json:"maxMsgSizeExceededDiscardedMsgCount,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to there being no eligible endpoint destination.
	NoEligibleDestinationsDiscardedMsgCount *int64 `json:"noEligibleDestinationsDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to no local delivery being requested.
	NoLocalDeliveryDiscardedMsgCount *int64 `json:"noLocalDeliveryDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to being incompatible with the forwarding mode of an eligible endpoint destination.
	NotCompatibleWithForwardingModeDiscardedMsgCount *int64 `json:"notCompatibleWithForwardingModeDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to being received out of order.
	OutOfOrderDiscardedMsgCount *int64 `json:"outOfOrderDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to being denied by the access control list (ACL) profile for the published topic.
	PublishAclDeniedDiscardedMsgCount *int64 `json:"publishAclDeniedDiscardedMsgCount,omitempty"`
	// The identifier (ID) of the publisher for the Flow.
	PublisherId *int64 `json:"publisherId,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to the destination queue not being found.
	QueueNotFoundDiscardedMsgCount *int64 `json:"queueNotFoundDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to the Message VPN being in the replication standby state.
	ReplicationStandbyDiscardedMsgCount *int64 `json:"replicationStandbyDiscardedMsgCount,omitempty"`
	// The name of the transacted session on the Flow.
	SessionName *string `json:"sessionName,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to the message time-to-live (TTL) count being exceeded. The message TTL count is the maximum number of times the message can cross a bridge between Message VPNs.
	SmfTtlExceededDiscardedMsgCount *int64 `json:"smfTtlExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to all available message spool file resources being used.
	SpoolFileLimitExceededDiscardedMsgCount *int64 `json:"spoolFileLimitExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to the message spool being not ready.
	SpoolNotReadyDiscardedMsgCount *int64 `json:"spoolNotReadyDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to a failure while spooling to the Assured Delivery Blade (ADB).
	SpoolToAdbFailDiscardedMsgCount *int64 `json:"spoolToAdbFailDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to a failure while spooling to the disk.
	SpoolToDiskFailDiscardedMsgCount *int64 `json:"spoolToDiskFailDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to the maximum message spool usage being exceeded.
	SpoolUsageExceededDiscardedMsgCount *int64 `json:"spoolUsageExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to synchronous replication being ineligible.
	SyncReplicationIneligibleDiscardedMsgCount *int64 `json:"syncReplicationIneligibleDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages from the Flow discarded due to being denied by the client profile.
	UserProfileDeniedGuaranteedDiscardedMsgCount *int64 `json:"userProfileDeniedGuaranteedDiscardedMsgCount,omitempty"`
	// The size of the window used for guaranteed messages sent on the Flow, in messages.
	WindowSize *int32 `json:"windowSize,omitempty"`
}

// NewMsgVpnClientRxFlow instantiates a new MsgVpnClientRxFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnClientRxFlow() *MsgVpnClientRxFlow {
	this := MsgVpnClientRxFlow{}
	return &this
}

// NewMsgVpnClientRxFlowWithDefaults instantiates a new MsgVpnClientRxFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnClientRxFlowWithDefaults() *MsgVpnClientRxFlow {
	this := MsgVpnClientRxFlow{}
	return &this
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *MsgVpnClientRxFlow) SetClientName(v string) {
	o.ClientName = &v
}

// GetConnectTime returns the ConnectTime field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetConnectTime() int32 {
	if o == nil || o.ConnectTime == nil {
		var ret int32
		return ret
	}
	return *o.ConnectTime
}

// GetConnectTimeOk returns a tuple with the ConnectTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetConnectTimeOk() (*int32, bool) {
	if o == nil || o.ConnectTime == nil {
		return nil, false
	}
	return o.ConnectTime, true
}

// HasConnectTime returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasConnectTime() bool {
	if o != nil && o.ConnectTime != nil {
		return true
	}

	return false
}

// SetConnectTime gets a reference to the given int32 and assigns it to the ConnectTime field.
func (o *MsgVpnClientRxFlow) SetConnectTime(v int32) {
	o.ConnectTime = &v
}

// GetDestinationGroupErrorDiscardedMsgCount returns the DestinationGroupErrorDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetDestinationGroupErrorDiscardedMsgCount() int64 {
	if o == nil || o.DestinationGroupErrorDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DestinationGroupErrorDiscardedMsgCount
}

// GetDestinationGroupErrorDiscardedMsgCountOk returns a tuple with the DestinationGroupErrorDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetDestinationGroupErrorDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.DestinationGroupErrorDiscardedMsgCount == nil {
		return nil, false
	}
	return o.DestinationGroupErrorDiscardedMsgCount, true
}

// HasDestinationGroupErrorDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasDestinationGroupErrorDiscardedMsgCount() bool {
	if o != nil && o.DestinationGroupErrorDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetDestinationGroupErrorDiscardedMsgCount gets a reference to the given int64 and assigns it to the DestinationGroupErrorDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetDestinationGroupErrorDiscardedMsgCount(v int64) {
	o.DestinationGroupErrorDiscardedMsgCount = &v
}

// GetDuplicateDiscardedMsgCount returns the DuplicateDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetDuplicateDiscardedMsgCount() int64 {
	if o == nil || o.DuplicateDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DuplicateDiscardedMsgCount
}

// GetDuplicateDiscardedMsgCountOk returns a tuple with the DuplicateDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetDuplicateDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.DuplicateDiscardedMsgCount == nil {
		return nil, false
	}
	return o.DuplicateDiscardedMsgCount, true
}

// HasDuplicateDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasDuplicateDiscardedMsgCount() bool {
	if o != nil && o.DuplicateDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetDuplicateDiscardedMsgCount gets a reference to the given int64 and assigns it to the DuplicateDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetDuplicateDiscardedMsgCount(v int64) {
	o.DuplicateDiscardedMsgCount = &v
}

// GetEndpointDisabledDiscardedMsgCount returns the EndpointDisabledDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetEndpointDisabledDiscardedMsgCount() int64 {
	if o == nil || o.EndpointDisabledDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.EndpointDisabledDiscardedMsgCount
}

// GetEndpointDisabledDiscardedMsgCountOk returns a tuple with the EndpointDisabledDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetEndpointDisabledDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.EndpointDisabledDiscardedMsgCount == nil {
		return nil, false
	}
	return o.EndpointDisabledDiscardedMsgCount, true
}

// HasEndpointDisabledDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasEndpointDisabledDiscardedMsgCount() bool {
	if o != nil && o.EndpointDisabledDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetEndpointDisabledDiscardedMsgCount gets a reference to the given int64 and assigns it to the EndpointDisabledDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetEndpointDisabledDiscardedMsgCount(v int64) {
	o.EndpointDisabledDiscardedMsgCount = &v
}

// GetEndpointUsageExceededDiscardedMsgCount returns the EndpointUsageExceededDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetEndpointUsageExceededDiscardedMsgCount() int64 {
	if o == nil || o.EndpointUsageExceededDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.EndpointUsageExceededDiscardedMsgCount
}

// GetEndpointUsageExceededDiscardedMsgCountOk returns a tuple with the EndpointUsageExceededDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetEndpointUsageExceededDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.EndpointUsageExceededDiscardedMsgCount == nil {
		return nil, false
	}
	return o.EndpointUsageExceededDiscardedMsgCount, true
}

// HasEndpointUsageExceededDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasEndpointUsageExceededDiscardedMsgCount() bool {
	if o != nil && o.EndpointUsageExceededDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetEndpointUsageExceededDiscardedMsgCount gets a reference to the given int64 and assigns it to the EndpointUsageExceededDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetEndpointUsageExceededDiscardedMsgCount(v int64) {
	o.EndpointUsageExceededDiscardedMsgCount = &v
}

// GetErroredDiscardedMsgCount returns the ErroredDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetErroredDiscardedMsgCount() int64 {
	if o == nil || o.ErroredDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ErroredDiscardedMsgCount
}

// GetErroredDiscardedMsgCountOk returns a tuple with the ErroredDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetErroredDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.ErroredDiscardedMsgCount == nil {
		return nil, false
	}
	return o.ErroredDiscardedMsgCount, true
}

// HasErroredDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasErroredDiscardedMsgCount() bool {
	if o != nil && o.ErroredDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetErroredDiscardedMsgCount gets a reference to the given int64 and assigns it to the ErroredDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetErroredDiscardedMsgCount(v int64) {
	o.ErroredDiscardedMsgCount = &v
}

// GetFlowId returns the FlowId field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetFlowId() int64 {
	if o == nil || o.FlowId == nil {
		var ret int64
		return ret
	}
	return *o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetFlowIdOk() (*int64, bool) {
	if o == nil || o.FlowId == nil {
		return nil, false
	}
	return o.FlowId, true
}

// HasFlowId returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasFlowId() bool {
	if o != nil && o.FlowId != nil {
		return true
	}

	return false
}

// SetFlowId gets a reference to the given int64 and assigns it to the FlowId field.
func (o *MsgVpnClientRxFlow) SetFlowId(v int64) {
	o.FlowId = &v
}

// GetFlowName returns the FlowName field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetFlowName() string {
	if o == nil || o.FlowName == nil {
		var ret string
		return ret
	}
	return *o.FlowName
}

// GetFlowNameOk returns a tuple with the FlowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetFlowNameOk() (*string, bool) {
	if o == nil || o.FlowName == nil {
		return nil, false
	}
	return o.FlowName, true
}

// HasFlowName returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasFlowName() bool {
	if o != nil && o.FlowName != nil {
		return true
	}

	return false
}

// SetFlowName gets a reference to the given string and assigns it to the FlowName field.
func (o *MsgVpnClientRxFlow) SetFlowName(v string) {
	o.FlowName = &v
}

// GetGuaranteedMsgCount returns the GuaranteedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetGuaranteedMsgCount() int64 {
	if o == nil || o.GuaranteedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.GuaranteedMsgCount
}

// GetGuaranteedMsgCountOk returns a tuple with the GuaranteedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetGuaranteedMsgCountOk() (*int64, bool) {
	if o == nil || o.GuaranteedMsgCount == nil {
		return nil, false
	}
	return o.GuaranteedMsgCount, true
}

// HasGuaranteedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasGuaranteedMsgCount() bool {
	if o != nil && o.GuaranteedMsgCount != nil {
		return true
	}

	return false
}

// SetGuaranteedMsgCount gets a reference to the given int64 and assigns it to the GuaranteedMsgCount field.
func (o *MsgVpnClientRxFlow) SetGuaranteedMsgCount(v int64) {
	o.GuaranteedMsgCount = &v
}

// GetLastRxMsgId returns the LastRxMsgId field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetLastRxMsgId() int64 {
	if o == nil || o.LastRxMsgId == nil {
		var ret int64
		return ret
	}
	return *o.LastRxMsgId
}

// GetLastRxMsgIdOk returns a tuple with the LastRxMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetLastRxMsgIdOk() (*int64, bool) {
	if o == nil || o.LastRxMsgId == nil {
		return nil, false
	}
	return o.LastRxMsgId, true
}

// HasLastRxMsgId returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasLastRxMsgId() bool {
	if o != nil && o.LastRxMsgId != nil {
		return true
	}

	return false
}

// SetLastRxMsgId gets a reference to the given int64 and assigns it to the LastRxMsgId field.
func (o *MsgVpnClientRxFlow) SetLastRxMsgId(v int64) {
	o.LastRxMsgId = &v
}

// GetLocalMsgCountExceededDiscardedMsgCount returns the LocalMsgCountExceededDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetLocalMsgCountExceededDiscardedMsgCount() int64 {
	if o == nil || o.LocalMsgCountExceededDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LocalMsgCountExceededDiscardedMsgCount
}

// GetLocalMsgCountExceededDiscardedMsgCountOk returns a tuple with the LocalMsgCountExceededDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetLocalMsgCountExceededDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.LocalMsgCountExceededDiscardedMsgCount == nil {
		return nil, false
	}
	return o.LocalMsgCountExceededDiscardedMsgCount, true
}

// HasLocalMsgCountExceededDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasLocalMsgCountExceededDiscardedMsgCount() bool {
	if o != nil && o.LocalMsgCountExceededDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetLocalMsgCountExceededDiscardedMsgCount gets a reference to the given int64 and assigns it to the LocalMsgCountExceededDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetLocalMsgCountExceededDiscardedMsgCount(v int64) {
	o.LocalMsgCountExceededDiscardedMsgCount = &v
}

// GetLowPriorityMsgCongestionDiscardedMsgCount returns the LowPriorityMsgCongestionDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetLowPriorityMsgCongestionDiscardedMsgCount() int64 {
	if o == nil || o.LowPriorityMsgCongestionDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LowPriorityMsgCongestionDiscardedMsgCount
}

// GetLowPriorityMsgCongestionDiscardedMsgCountOk returns a tuple with the LowPriorityMsgCongestionDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetLowPriorityMsgCongestionDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.LowPriorityMsgCongestionDiscardedMsgCount == nil {
		return nil, false
	}
	return o.LowPriorityMsgCongestionDiscardedMsgCount, true
}

// HasLowPriorityMsgCongestionDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasLowPriorityMsgCongestionDiscardedMsgCount() bool {
	if o != nil && o.LowPriorityMsgCongestionDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetLowPriorityMsgCongestionDiscardedMsgCount gets a reference to the given int64 and assigns it to the LowPriorityMsgCongestionDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetLowPriorityMsgCongestionDiscardedMsgCount(v int64) {
	o.LowPriorityMsgCongestionDiscardedMsgCount = &v
}

// GetMaxMsgSizeExceededDiscardedMsgCount returns the MaxMsgSizeExceededDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetMaxMsgSizeExceededDiscardedMsgCount() int64 {
	if o == nil || o.MaxMsgSizeExceededDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxMsgSizeExceededDiscardedMsgCount
}

// GetMaxMsgSizeExceededDiscardedMsgCountOk returns a tuple with the MaxMsgSizeExceededDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetMaxMsgSizeExceededDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxMsgSizeExceededDiscardedMsgCount == nil {
		return nil, false
	}
	return o.MaxMsgSizeExceededDiscardedMsgCount, true
}

// HasMaxMsgSizeExceededDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasMaxMsgSizeExceededDiscardedMsgCount() bool {
	if o != nil && o.MaxMsgSizeExceededDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetMaxMsgSizeExceededDiscardedMsgCount gets a reference to the given int64 and assigns it to the MaxMsgSizeExceededDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetMaxMsgSizeExceededDiscardedMsgCount(v int64) {
	o.MaxMsgSizeExceededDiscardedMsgCount = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnClientRxFlow) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetNoEligibleDestinationsDiscardedMsgCount returns the NoEligibleDestinationsDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetNoEligibleDestinationsDiscardedMsgCount() int64 {
	if o == nil || o.NoEligibleDestinationsDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.NoEligibleDestinationsDiscardedMsgCount
}

// GetNoEligibleDestinationsDiscardedMsgCountOk returns a tuple with the NoEligibleDestinationsDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetNoEligibleDestinationsDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.NoEligibleDestinationsDiscardedMsgCount == nil {
		return nil, false
	}
	return o.NoEligibleDestinationsDiscardedMsgCount, true
}

// HasNoEligibleDestinationsDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasNoEligibleDestinationsDiscardedMsgCount() bool {
	if o != nil && o.NoEligibleDestinationsDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetNoEligibleDestinationsDiscardedMsgCount gets a reference to the given int64 and assigns it to the NoEligibleDestinationsDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetNoEligibleDestinationsDiscardedMsgCount(v int64) {
	o.NoEligibleDestinationsDiscardedMsgCount = &v
}

// GetNoLocalDeliveryDiscardedMsgCount returns the NoLocalDeliveryDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetNoLocalDeliveryDiscardedMsgCount() int64 {
	if o == nil || o.NoLocalDeliveryDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.NoLocalDeliveryDiscardedMsgCount
}

// GetNoLocalDeliveryDiscardedMsgCountOk returns a tuple with the NoLocalDeliveryDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetNoLocalDeliveryDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.NoLocalDeliveryDiscardedMsgCount == nil {
		return nil, false
	}
	return o.NoLocalDeliveryDiscardedMsgCount, true
}

// HasNoLocalDeliveryDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasNoLocalDeliveryDiscardedMsgCount() bool {
	if o != nil && o.NoLocalDeliveryDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetNoLocalDeliveryDiscardedMsgCount gets a reference to the given int64 and assigns it to the NoLocalDeliveryDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetNoLocalDeliveryDiscardedMsgCount(v int64) {
	o.NoLocalDeliveryDiscardedMsgCount = &v
}

// GetNotCompatibleWithForwardingModeDiscardedMsgCount returns the NotCompatibleWithForwardingModeDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetNotCompatibleWithForwardingModeDiscardedMsgCount() int64 {
	if o == nil || o.NotCompatibleWithForwardingModeDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.NotCompatibleWithForwardingModeDiscardedMsgCount
}

// GetNotCompatibleWithForwardingModeDiscardedMsgCountOk returns a tuple with the NotCompatibleWithForwardingModeDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetNotCompatibleWithForwardingModeDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.NotCompatibleWithForwardingModeDiscardedMsgCount == nil {
		return nil, false
	}
	return o.NotCompatibleWithForwardingModeDiscardedMsgCount, true
}

// HasNotCompatibleWithForwardingModeDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasNotCompatibleWithForwardingModeDiscardedMsgCount() bool {
	if o != nil && o.NotCompatibleWithForwardingModeDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetNotCompatibleWithForwardingModeDiscardedMsgCount gets a reference to the given int64 and assigns it to the NotCompatibleWithForwardingModeDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetNotCompatibleWithForwardingModeDiscardedMsgCount(v int64) {
	o.NotCompatibleWithForwardingModeDiscardedMsgCount = &v
}

// GetOutOfOrderDiscardedMsgCount returns the OutOfOrderDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetOutOfOrderDiscardedMsgCount() int64 {
	if o == nil || o.OutOfOrderDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.OutOfOrderDiscardedMsgCount
}

// GetOutOfOrderDiscardedMsgCountOk returns a tuple with the OutOfOrderDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetOutOfOrderDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.OutOfOrderDiscardedMsgCount == nil {
		return nil, false
	}
	return o.OutOfOrderDiscardedMsgCount, true
}

// HasOutOfOrderDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasOutOfOrderDiscardedMsgCount() bool {
	if o != nil && o.OutOfOrderDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetOutOfOrderDiscardedMsgCount gets a reference to the given int64 and assigns it to the OutOfOrderDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetOutOfOrderDiscardedMsgCount(v int64) {
	o.OutOfOrderDiscardedMsgCount = &v
}

// GetPublishAclDeniedDiscardedMsgCount returns the PublishAclDeniedDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetPublishAclDeniedDiscardedMsgCount() int64 {
	if o == nil || o.PublishAclDeniedDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.PublishAclDeniedDiscardedMsgCount
}

// GetPublishAclDeniedDiscardedMsgCountOk returns a tuple with the PublishAclDeniedDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetPublishAclDeniedDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.PublishAclDeniedDiscardedMsgCount == nil {
		return nil, false
	}
	return o.PublishAclDeniedDiscardedMsgCount, true
}

// HasPublishAclDeniedDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasPublishAclDeniedDiscardedMsgCount() bool {
	if o != nil && o.PublishAclDeniedDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetPublishAclDeniedDiscardedMsgCount gets a reference to the given int64 and assigns it to the PublishAclDeniedDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetPublishAclDeniedDiscardedMsgCount(v int64) {
	o.PublishAclDeniedDiscardedMsgCount = &v
}

// GetPublisherId returns the PublisherId field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetPublisherId() int64 {
	if o == nil || o.PublisherId == nil {
		var ret int64
		return ret
	}
	return *o.PublisherId
}

// GetPublisherIdOk returns a tuple with the PublisherId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetPublisherIdOk() (*int64, bool) {
	if o == nil || o.PublisherId == nil {
		return nil, false
	}
	return o.PublisherId, true
}

// HasPublisherId returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasPublisherId() bool {
	if o != nil && o.PublisherId != nil {
		return true
	}

	return false
}

// SetPublisherId gets a reference to the given int64 and assigns it to the PublisherId field.
func (o *MsgVpnClientRxFlow) SetPublisherId(v int64) {
	o.PublisherId = &v
}

// GetQueueNotFoundDiscardedMsgCount returns the QueueNotFoundDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetQueueNotFoundDiscardedMsgCount() int64 {
	if o == nil || o.QueueNotFoundDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.QueueNotFoundDiscardedMsgCount
}

// GetQueueNotFoundDiscardedMsgCountOk returns a tuple with the QueueNotFoundDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetQueueNotFoundDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.QueueNotFoundDiscardedMsgCount == nil {
		return nil, false
	}
	return o.QueueNotFoundDiscardedMsgCount, true
}

// HasQueueNotFoundDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasQueueNotFoundDiscardedMsgCount() bool {
	if o != nil && o.QueueNotFoundDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetQueueNotFoundDiscardedMsgCount gets a reference to the given int64 and assigns it to the QueueNotFoundDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetQueueNotFoundDiscardedMsgCount(v int64) {
	o.QueueNotFoundDiscardedMsgCount = &v
}

// GetReplicationStandbyDiscardedMsgCount returns the ReplicationStandbyDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetReplicationStandbyDiscardedMsgCount() int64 {
	if o == nil || o.ReplicationStandbyDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyDiscardedMsgCount
}

// GetReplicationStandbyDiscardedMsgCountOk returns a tuple with the ReplicationStandbyDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetReplicationStandbyDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyDiscardedMsgCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyDiscardedMsgCount, true
}

// HasReplicationStandbyDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasReplicationStandbyDiscardedMsgCount() bool {
	if o != nil && o.ReplicationStandbyDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyDiscardedMsgCount gets a reference to the given int64 and assigns it to the ReplicationStandbyDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetReplicationStandbyDiscardedMsgCount(v int64) {
	o.ReplicationStandbyDiscardedMsgCount = &v
}

// GetSessionName returns the SessionName field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetSessionName() string {
	if o == nil || o.SessionName == nil {
		var ret string
		return ret
	}
	return *o.SessionName
}

// GetSessionNameOk returns a tuple with the SessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetSessionNameOk() (*string, bool) {
	if o == nil || o.SessionName == nil {
		return nil, false
	}
	return o.SessionName, true
}

// HasSessionName returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasSessionName() bool {
	if o != nil && o.SessionName != nil {
		return true
	}

	return false
}

// SetSessionName gets a reference to the given string and assigns it to the SessionName field.
func (o *MsgVpnClientRxFlow) SetSessionName(v string) {
	o.SessionName = &v
}

// GetSmfTtlExceededDiscardedMsgCount returns the SmfTtlExceededDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetSmfTtlExceededDiscardedMsgCount() int64 {
	if o == nil || o.SmfTtlExceededDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SmfTtlExceededDiscardedMsgCount
}

// GetSmfTtlExceededDiscardedMsgCountOk returns a tuple with the SmfTtlExceededDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetSmfTtlExceededDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.SmfTtlExceededDiscardedMsgCount == nil {
		return nil, false
	}
	return o.SmfTtlExceededDiscardedMsgCount, true
}

// HasSmfTtlExceededDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasSmfTtlExceededDiscardedMsgCount() bool {
	if o != nil && o.SmfTtlExceededDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetSmfTtlExceededDiscardedMsgCount gets a reference to the given int64 and assigns it to the SmfTtlExceededDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetSmfTtlExceededDiscardedMsgCount(v int64) {
	o.SmfTtlExceededDiscardedMsgCount = &v
}

// GetSpoolFileLimitExceededDiscardedMsgCount returns the SpoolFileLimitExceededDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetSpoolFileLimitExceededDiscardedMsgCount() int64 {
	if o == nil || o.SpoolFileLimitExceededDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SpoolFileLimitExceededDiscardedMsgCount
}

// GetSpoolFileLimitExceededDiscardedMsgCountOk returns a tuple with the SpoolFileLimitExceededDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetSpoolFileLimitExceededDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.SpoolFileLimitExceededDiscardedMsgCount == nil {
		return nil, false
	}
	return o.SpoolFileLimitExceededDiscardedMsgCount, true
}

// HasSpoolFileLimitExceededDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasSpoolFileLimitExceededDiscardedMsgCount() bool {
	if o != nil && o.SpoolFileLimitExceededDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetSpoolFileLimitExceededDiscardedMsgCount gets a reference to the given int64 and assigns it to the SpoolFileLimitExceededDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetSpoolFileLimitExceededDiscardedMsgCount(v int64) {
	o.SpoolFileLimitExceededDiscardedMsgCount = &v
}

// GetSpoolNotReadyDiscardedMsgCount returns the SpoolNotReadyDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetSpoolNotReadyDiscardedMsgCount() int64 {
	if o == nil || o.SpoolNotReadyDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SpoolNotReadyDiscardedMsgCount
}

// GetSpoolNotReadyDiscardedMsgCountOk returns a tuple with the SpoolNotReadyDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetSpoolNotReadyDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.SpoolNotReadyDiscardedMsgCount == nil {
		return nil, false
	}
	return o.SpoolNotReadyDiscardedMsgCount, true
}

// HasSpoolNotReadyDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasSpoolNotReadyDiscardedMsgCount() bool {
	if o != nil && o.SpoolNotReadyDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetSpoolNotReadyDiscardedMsgCount gets a reference to the given int64 and assigns it to the SpoolNotReadyDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetSpoolNotReadyDiscardedMsgCount(v int64) {
	o.SpoolNotReadyDiscardedMsgCount = &v
}

// GetSpoolToAdbFailDiscardedMsgCount returns the SpoolToAdbFailDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetSpoolToAdbFailDiscardedMsgCount() int64 {
	if o == nil || o.SpoolToAdbFailDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SpoolToAdbFailDiscardedMsgCount
}

// GetSpoolToAdbFailDiscardedMsgCountOk returns a tuple with the SpoolToAdbFailDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetSpoolToAdbFailDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.SpoolToAdbFailDiscardedMsgCount == nil {
		return nil, false
	}
	return o.SpoolToAdbFailDiscardedMsgCount, true
}

// HasSpoolToAdbFailDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasSpoolToAdbFailDiscardedMsgCount() bool {
	if o != nil && o.SpoolToAdbFailDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetSpoolToAdbFailDiscardedMsgCount gets a reference to the given int64 and assigns it to the SpoolToAdbFailDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetSpoolToAdbFailDiscardedMsgCount(v int64) {
	o.SpoolToAdbFailDiscardedMsgCount = &v
}

// GetSpoolToDiskFailDiscardedMsgCount returns the SpoolToDiskFailDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetSpoolToDiskFailDiscardedMsgCount() int64 {
	if o == nil || o.SpoolToDiskFailDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SpoolToDiskFailDiscardedMsgCount
}

// GetSpoolToDiskFailDiscardedMsgCountOk returns a tuple with the SpoolToDiskFailDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetSpoolToDiskFailDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.SpoolToDiskFailDiscardedMsgCount == nil {
		return nil, false
	}
	return o.SpoolToDiskFailDiscardedMsgCount, true
}

// HasSpoolToDiskFailDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasSpoolToDiskFailDiscardedMsgCount() bool {
	if o != nil && o.SpoolToDiskFailDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetSpoolToDiskFailDiscardedMsgCount gets a reference to the given int64 and assigns it to the SpoolToDiskFailDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetSpoolToDiskFailDiscardedMsgCount(v int64) {
	o.SpoolToDiskFailDiscardedMsgCount = &v
}

// GetSpoolUsageExceededDiscardedMsgCount returns the SpoolUsageExceededDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetSpoolUsageExceededDiscardedMsgCount() int64 {
	if o == nil || o.SpoolUsageExceededDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SpoolUsageExceededDiscardedMsgCount
}

// GetSpoolUsageExceededDiscardedMsgCountOk returns a tuple with the SpoolUsageExceededDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetSpoolUsageExceededDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.SpoolUsageExceededDiscardedMsgCount == nil {
		return nil, false
	}
	return o.SpoolUsageExceededDiscardedMsgCount, true
}

// HasSpoolUsageExceededDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasSpoolUsageExceededDiscardedMsgCount() bool {
	if o != nil && o.SpoolUsageExceededDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetSpoolUsageExceededDiscardedMsgCount gets a reference to the given int64 and assigns it to the SpoolUsageExceededDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetSpoolUsageExceededDiscardedMsgCount(v int64) {
	o.SpoolUsageExceededDiscardedMsgCount = &v
}

// GetSyncReplicationIneligibleDiscardedMsgCount returns the SyncReplicationIneligibleDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetSyncReplicationIneligibleDiscardedMsgCount() int64 {
	if o == nil || o.SyncReplicationIneligibleDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SyncReplicationIneligibleDiscardedMsgCount
}

// GetSyncReplicationIneligibleDiscardedMsgCountOk returns a tuple with the SyncReplicationIneligibleDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetSyncReplicationIneligibleDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.SyncReplicationIneligibleDiscardedMsgCount == nil {
		return nil, false
	}
	return o.SyncReplicationIneligibleDiscardedMsgCount, true
}

// HasSyncReplicationIneligibleDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasSyncReplicationIneligibleDiscardedMsgCount() bool {
	if o != nil && o.SyncReplicationIneligibleDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetSyncReplicationIneligibleDiscardedMsgCount gets a reference to the given int64 and assigns it to the SyncReplicationIneligibleDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetSyncReplicationIneligibleDiscardedMsgCount(v int64) {
	o.SyncReplicationIneligibleDiscardedMsgCount = &v
}

// GetUserProfileDeniedGuaranteedDiscardedMsgCount returns the UserProfileDeniedGuaranteedDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetUserProfileDeniedGuaranteedDiscardedMsgCount() int64 {
	if o == nil || o.UserProfileDeniedGuaranteedDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.UserProfileDeniedGuaranteedDiscardedMsgCount
}

// GetUserProfileDeniedGuaranteedDiscardedMsgCountOk returns a tuple with the UserProfileDeniedGuaranteedDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetUserProfileDeniedGuaranteedDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.UserProfileDeniedGuaranteedDiscardedMsgCount == nil {
		return nil, false
	}
	return o.UserProfileDeniedGuaranteedDiscardedMsgCount, true
}

// HasUserProfileDeniedGuaranteedDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasUserProfileDeniedGuaranteedDiscardedMsgCount() bool {
	if o != nil && o.UserProfileDeniedGuaranteedDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetUserProfileDeniedGuaranteedDiscardedMsgCount gets a reference to the given int64 and assigns it to the UserProfileDeniedGuaranteedDiscardedMsgCount field.
func (o *MsgVpnClientRxFlow) SetUserProfileDeniedGuaranteedDiscardedMsgCount(v int64) {
	o.UserProfileDeniedGuaranteedDiscardedMsgCount = &v
}

// GetWindowSize returns the WindowSize field value if set, zero value otherwise.
func (o *MsgVpnClientRxFlow) GetWindowSize() int32 {
	if o == nil || o.WindowSize == nil {
		var ret int32
		return ret
	}
	return *o.WindowSize
}

// GetWindowSizeOk returns a tuple with the WindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientRxFlow) GetWindowSizeOk() (*int32, bool) {
	if o == nil || o.WindowSize == nil {
		return nil, false
	}
	return o.WindowSize, true
}

// HasWindowSize returns a boolean if a field has been set.
func (o *MsgVpnClientRxFlow) HasWindowSize() bool {
	if o != nil && o.WindowSize != nil {
		return true
	}

	return false
}

// SetWindowSize gets a reference to the given int32 and assigns it to the WindowSize field.
func (o *MsgVpnClientRxFlow) SetWindowSize(v int32) {
	o.WindowSize = &v
}

func (o MsgVpnClientRxFlow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.ConnectTime != nil {
		toSerialize["connectTime"] = o.ConnectTime
	}
	if o.DestinationGroupErrorDiscardedMsgCount != nil {
		toSerialize["destinationGroupErrorDiscardedMsgCount"] = o.DestinationGroupErrorDiscardedMsgCount
	}
	if o.DuplicateDiscardedMsgCount != nil {
		toSerialize["duplicateDiscardedMsgCount"] = o.DuplicateDiscardedMsgCount
	}
	if o.EndpointDisabledDiscardedMsgCount != nil {
		toSerialize["endpointDisabledDiscardedMsgCount"] = o.EndpointDisabledDiscardedMsgCount
	}
	if o.EndpointUsageExceededDiscardedMsgCount != nil {
		toSerialize["endpointUsageExceededDiscardedMsgCount"] = o.EndpointUsageExceededDiscardedMsgCount
	}
	if o.ErroredDiscardedMsgCount != nil {
		toSerialize["erroredDiscardedMsgCount"] = o.ErroredDiscardedMsgCount
	}
	if o.FlowId != nil {
		toSerialize["flowId"] = o.FlowId
	}
	if o.FlowName != nil {
		toSerialize["flowName"] = o.FlowName
	}
	if o.GuaranteedMsgCount != nil {
		toSerialize["guaranteedMsgCount"] = o.GuaranteedMsgCount
	}
	if o.LastRxMsgId != nil {
		toSerialize["lastRxMsgId"] = o.LastRxMsgId
	}
	if o.LocalMsgCountExceededDiscardedMsgCount != nil {
		toSerialize["localMsgCountExceededDiscardedMsgCount"] = o.LocalMsgCountExceededDiscardedMsgCount
	}
	if o.LowPriorityMsgCongestionDiscardedMsgCount != nil {
		toSerialize["lowPriorityMsgCongestionDiscardedMsgCount"] = o.LowPriorityMsgCongestionDiscardedMsgCount
	}
	if o.MaxMsgSizeExceededDiscardedMsgCount != nil {
		toSerialize["maxMsgSizeExceededDiscardedMsgCount"] = o.MaxMsgSizeExceededDiscardedMsgCount
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.NoEligibleDestinationsDiscardedMsgCount != nil {
		toSerialize["noEligibleDestinationsDiscardedMsgCount"] = o.NoEligibleDestinationsDiscardedMsgCount
	}
	if o.NoLocalDeliveryDiscardedMsgCount != nil {
		toSerialize["noLocalDeliveryDiscardedMsgCount"] = o.NoLocalDeliveryDiscardedMsgCount
	}
	if o.NotCompatibleWithForwardingModeDiscardedMsgCount != nil {
		toSerialize["notCompatibleWithForwardingModeDiscardedMsgCount"] = o.NotCompatibleWithForwardingModeDiscardedMsgCount
	}
	if o.OutOfOrderDiscardedMsgCount != nil {
		toSerialize["outOfOrderDiscardedMsgCount"] = o.OutOfOrderDiscardedMsgCount
	}
	if o.PublishAclDeniedDiscardedMsgCount != nil {
		toSerialize["publishAclDeniedDiscardedMsgCount"] = o.PublishAclDeniedDiscardedMsgCount
	}
	if o.PublisherId != nil {
		toSerialize["publisherId"] = o.PublisherId
	}
	if o.QueueNotFoundDiscardedMsgCount != nil {
		toSerialize["queueNotFoundDiscardedMsgCount"] = o.QueueNotFoundDiscardedMsgCount
	}
	if o.ReplicationStandbyDiscardedMsgCount != nil {
		toSerialize["replicationStandbyDiscardedMsgCount"] = o.ReplicationStandbyDiscardedMsgCount
	}
	if o.SessionName != nil {
		toSerialize["sessionName"] = o.SessionName
	}
	if o.SmfTtlExceededDiscardedMsgCount != nil {
		toSerialize["smfTtlExceededDiscardedMsgCount"] = o.SmfTtlExceededDiscardedMsgCount
	}
	if o.SpoolFileLimitExceededDiscardedMsgCount != nil {
		toSerialize["spoolFileLimitExceededDiscardedMsgCount"] = o.SpoolFileLimitExceededDiscardedMsgCount
	}
	if o.SpoolNotReadyDiscardedMsgCount != nil {
		toSerialize["spoolNotReadyDiscardedMsgCount"] = o.SpoolNotReadyDiscardedMsgCount
	}
	if o.SpoolToAdbFailDiscardedMsgCount != nil {
		toSerialize["spoolToAdbFailDiscardedMsgCount"] = o.SpoolToAdbFailDiscardedMsgCount
	}
	if o.SpoolToDiskFailDiscardedMsgCount != nil {
		toSerialize["spoolToDiskFailDiscardedMsgCount"] = o.SpoolToDiskFailDiscardedMsgCount
	}
	if o.SpoolUsageExceededDiscardedMsgCount != nil {
		toSerialize["spoolUsageExceededDiscardedMsgCount"] = o.SpoolUsageExceededDiscardedMsgCount
	}
	if o.SyncReplicationIneligibleDiscardedMsgCount != nil {
		toSerialize["syncReplicationIneligibleDiscardedMsgCount"] = o.SyncReplicationIneligibleDiscardedMsgCount
	}
	if o.UserProfileDeniedGuaranteedDiscardedMsgCount != nil {
		toSerialize["userProfileDeniedGuaranteedDiscardedMsgCount"] = o.UserProfileDeniedGuaranteedDiscardedMsgCount
	}
	if o.WindowSize != nil {
		toSerialize["windowSize"] = o.WindowSize
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnClientRxFlow struct {
	value *MsgVpnClientRxFlow
	isSet bool
}

func (v NullableMsgVpnClientRxFlow) Get() *MsgVpnClientRxFlow {
	return v.value
}

func (v *NullableMsgVpnClientRxFlow) Set(val *MsgVpnClientRxFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnClientRxFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnClientRxFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnClientRxFlow(val *MsgVpnClientRxFlow) *NullableMsgVpnClientRxFlow {
	return &NullableMsgVpnClientRxFlow{value: val, isSet: true}
}

func (v NullableMsgVpnClientRxFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnClientRxFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
