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

// MsgVpnBridgeCounter The counters for the Bridge. Deprecated since 2.13. All attributes in this object have been moved to the MsgVpnBridge object.
type MsgVpnBridgeCounter struct {
	// The amount of client control messages received from the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	ControlRxByteCount *int64 `json:"controlRxByteCount,omitempty"`
	// The number of client control messages received from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	ControlRxMsgCount *int64 `json:"controlRxMsgCount,omitempty"`
	// The amount of client control messages transmitted to the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	ControlTxByteCount *int64 `json:"controlTxByteCount,omitempty"`
	// The number of client control messages transmitted to the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	ControlTxMsgCount *int64 `json:"controlTxMsgCount,omitempty"`
	// The amount of client data messages received from the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	DataRxByteCount *int64 `json:"dataRxByteCount,omitempty"`
	// The number of client data messages received from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	DataRxMsgCount *int64 `json:"dataRxMsgCount,omitempty"`
	// The amount of client data messages transmitted to the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	DataTxByteCount *int64 `json:"dataTxByteCount,omitempty"`
	// The number of client data messages transmitted to the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	DataTxMsgCount *int64 `json:"dataTxMsgCount,omitempty"`
	// The number of messages discarded during reception from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	DiscardedRxMsgCount *int64 `json:"discardedRxMsgCount,omitempty"`
	// The number of messages discarded during transmission to the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	DiscardedTxMsgCount *int64 `json:"discardedTxMsgCount,omitempty"`
	// The number of login request messages received from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	LoginRxMsgCount *int64 `json:"loginRxMsgCount,omitempty"`
	// The number of login response messages transmitted to the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	LoginTxMsgCount *int64 `json:"loginTxMsgCount,omitempty"`
	// The number of guaranteed messages received from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	MsgSpoolRxMsgCount *int64 `json:"msgSpoolRxMsgCount,omitempty"`
	// The amount of messages received from the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	RxByteCount *int64 `json:"rxByteCount,omitempty"`
	// The number of messages received from the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	RxMsgCount *int64 `json:"rxMsgCount,omitempty"`
	// The amount of messages transmitted to the Bridge, in bytes (B). Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	TxByteCount *int64 `json:"txByteCount,omitempty"`
	// The number of messages transmitted to the Bridge. Deprecated since 2.13. This attribute has been moved to the MsgVpnBridge object.
	TxMsgCount *int64 `json:"txMsgCount,omitempty"`
}

// NewMsgVpnBridgeCounter instantiates a new MsgVpnBridgeCounter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnBridgeCounter() *MsgVpnBridgeCounter {
	this := MsgVpnBridgeCounter{}
	return &this
}

// NewMsgVpnBridgeCounterWithDefaults instantiates a new MsgVpnBridgeCounter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnBridgeCounterWithDefaults() *MsgVpnBridgeCounter {
	this := MsgVpnBridgeCounter{}
	return &this
}

// GetControlRxByteCount returns the ControlRxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetControlRxByteCount() int64 {
	if o == nil || o.ControlRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlRxByteCount
}

// GetControlRxByteCountOk returns a tuple with the ControlRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetControlRxByteCountOk() (*int64, bool) {
	if o == nil || o.ControlRxByteCount == nil {
		return nil, false
	}
	return o.ControlRxByteCount, true
}

// HasControlRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasControlRxByteCount() bool {
	if o != nil && o.ControlRxByteCount != nil {
		return true
	}

	return false
}

// SetControlRxByteCount gets a reference to the given int64 and assigns it to the ControlRxByteCount field.
func (o *MsgVpnBridgeCounter) SetControlRxByteCount(v int64) {
	o.ControlRxByteCount = &v
}

// GetControlRxMsgCount returns the ControlRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetControlRxMsgCount() int64 {
	if o == nil || o.ControlRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlRxMsgCount
}

// GetControlRxMsgCountOk returns a tuple with the ControlRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetControlRxMsgCountOk() (*int64, bool) {
	if o == nil || o.ControlRxMsgCount == nil {
		return nil, false
	}
	return o.ControlRxMsgCount, true
}

// HasControlRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasControlRxMsgCount() bool {
	if o != nil && o.ControlRxMsgCount != nil {
		return true
	}

	return false
}

// SetControlRxMsgCount gets a reference to the given int64 and assigns it to the ControlRxMsgCount field.
func (o *MsgVpnBridgeCounter) SetControlRxMsgCount(v int64) {
	o.ControlRxMsgCount = &v
}

// GetControlTxByteCount returns the ControlTxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetControlTxByteCount() int64 {
	if o == nil || o.ControlTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlTxByteCount
}

// GetControlTxByteCountOk returns a tuple with the ControlTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetControlTxByteCountOk() (*int64, bool) {
	if o == nil || o.ControlTxByteCount == nil {
		return nil, false
	}
	return o.ControlTxByteCount, true
}

// HasControlTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasControlTxByteCount() bool {
	if o != nil && o.ControlTxByteCount != nil {
		return true
	}

	return false
}

// SetControlTxByteCount gets a reference to the given int64 and assigns it to the ControlTxByteCount field.
func (o *MsgVpnBridgeCounter) SetControlTxByteCount(v int64) {
	o.ControlTxByteCount = &v
}

// GetControlTxMsgCount returns the ControlTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetControlTxMsgCount() int64 {
	if o == nil || o.ControlTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlTxMsgCount
}

// GetControlTxMsgCountOk returns a tuple with the ControlTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetControlTxMsgCountOk() (*int64, bool) {
	if o == nil || o.ControlTxMsgCount == nil {
		return nil, false
	}
	return o.ControlTxMsgCount, true
}

// HasControlTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasControlTxMsgCount() bool {
	if o != nil && o.ControlTxMsgCount != nil {
		return true
	}

	return false
}

// SetControlTxMsgCount gets a reference to the given int64 and assigns it to the ControlTxMsgCount field.
func (o *MsgVpnBridgeCounter) SetControlTxMsgCount(v int64) {
	o.ControlTxMsgCount = &v
}

// GetDataRxByteCount returns the DataRxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetDataRxByteCount() int64 {
	if o == nil || o.DataRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.DataRxByteCount
}

// GetDataRxByteCountOk returns a tuple with the DataRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetDataRxByteCountOk() (*int64, bool) {
	if o == nil || o.DataRxByteCount == nil {
		return nil, false
	}
	return o.DataRxByteCount, true
}

// HasDataRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasDataRxByteCount() bool {
	if o != nil && o.DataRxByteCount != nil {
		return true
	}

	return false
}

// SetDataRxByteCount gets a reference to the given int64 and assigns it to the DataRxByteCount field.
func (o *MsgVpnBridgeCounter) SetDataRxByteCount(v int64) {
	o.DataRxByteCount = &v
}

// GetDataRxMsgCount returns the DataRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetDataRxMsgCount() int64 {
	if o == nil || o.DataRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DataRxMsgCount
}

// GetDataRxMsgCountOk returns a tuple with the DataRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetDataRxMsgCountOk() (*int64, bool) {
	if o == nil || o.DataRxMsgCount == nil {
		return nil, false
	}
	return o.DataRxMsgCount, true
}

// HasDataRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasDataRxMsgCount() bool {
	if o != nil && o.DataRxMsgCount != nil {
		return true
	}

	return false
}

// SetDataRxMsgCount gets a reference to the given int64 and assigns it to the DataRxMsgCount field.
func (o *MsgVpnBridgeCounter) SetDataRxMsgCount(v int64) {
	o.DataRxMsgCount = &v
}

// GetDataTxByteCount returns the DataTxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetDataTxByteCount() int64 {
	if o == nil || o.DataTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.DataTxByteCount
}

// GetDataTxByteCountOk returns a tuple with the DataTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetDataTxByteCountOk() (*int64, bool) {
	if o == nil || o.DataTxByteCount == nil {
		return nil, false
	}
	return o.DataTxByteCount, true
}

// HasDataTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasDataTxByteCount() bool {
	if o != nil && o.DataTxByteCount != nil {
		return true
	}

	return false
}

// SetDataTxByteCount gets a reference to the given int64 and assigns it to the DataTxByteCount field.
func (o *MsgVpnBridgeCounter) SetDataTxByteCount(v int64) {
	o.DataTxByteCount = &v
}

// GetDataTxMsgCount returns the DataTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetDataTxMsgCount() int64 {
	if o == nil || o.DataTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DataTxMsgCount
}

// GetDataTxMsgCountOk returns a tuple with the DataTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetDataTxMsgCountOk() (*int64, bool) {
	if o == nil || o.DataTxMsgCount == nil {
		return nil, false
	}
	return o.DataTxMsgCount, true
}

// HasDataTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasDataTxMsgCount() bool {
	if o != nil && o.DataTxMsgCount != nil {
		return true
	}

	return false
}

// SetDataTxMsgCount gets a reference to the given int64 and assigns it to the DataTxMsgCount field.
func (o *MsgVpnBridgeCounter) SetDataTxMsgCount(v int64) {
	o.DataTxMsgCount = &v
}

// GetDiscardedRxMsgCount returns the DiscardedRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetDiscardedRxMsgCount() int64 {
	if o == nil || o.DiscardedRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DiscardedRxMsgCount
}

// GetDiscardedRxMsgCountOk returns a tuple with the DiscardedRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetDiscardedRxMsgCountOk() (*int64, bool) {
	if o == nil || o.DiscardedRxMsgCount == nil {
		return nil, false
	}
	return o.DiscardedRxMsgCount, true
}

// HasDiscardedRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasDiscardedRxMsgCount() bool {
	if o != nil && o.DiscardedRxMsgCount != nil {
		return true
	}

	return false
}

// SetDiscardedRxMsgCount gets a reference to the given int64 and assigns it to the DiscardedRxMsgCount field.
func (o *MsgVpnBridgeCounter) SetDiscardedRxMsgCount(v int64) {
	o.DiscardedRxMsgCount = &v
}

// GetDiscardedTxMsgCount returns the DiscardedTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetDiscardedTxMsgCount() int64 {
	if o == nil || o.DiscardedTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DiscardedTxMsgCount
}

// GetDiscardedTxMsgCountOk returns a tuple with the DiscardedTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetDiscardedTxMsgCountOk() (*int64, bool) {
	if o == nil || o.DiscardedTxMsgCount == nil {
		return nil, false
	}
	return o.DiscardedTxMsgCount, true
}

// HasDiscardedTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasDiscardedTxMsgCount() bool {
	if o != nil && o.DiscardedTxMsgCount != nil {
		return true
	}

	return false
}

// SetDiscardedTxMsgCount gets a reference to the given int64 and assigns it to the DiscardedTxMsgCount field.
func (o *MsgVpnBridgeCounter) SetDiscardedTxMsgCount(v int64) {
	o.DiscardedTxMsgCount = &v
}

// GetLoginRxMsgCount returns the LoginRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetLoginRxMsgCount() int64 {
	if o == nil || o.LoginRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginRxMsgCount
}

// GetLoginRxMsgCountOk returns a tuple with the LoginRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetLoginRxMsgCountOk() (*int64, bool) {
	if o == nil || o.LoginRxMsgCount == nil {
		return nil, false
	}
	return o.LoginRxMsgCount, true
}

// HasLoginRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasLoginRxMsgCount() bool {
	if o != nil && o.LoginRxMsgCount != nil {
		return true
	}

	return false
}

// SetLoginRxMsgCount gets a reference to the given int64 and assigns it to the LoginRxMsgCount field.
func (o *MsgVpnBridgeCounter) SetLoginRxMsgCount(v int64) {
	o.LoginRxMsgCount = &v
}

// GetLoginTxMsgCount returns the LoginTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetLoginTxMsgCount() int64 {
	if o == nil || o.LoginTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginTxMsgCount
}

// GetLoginTxMsgCountOk returns a tuple with the LoginTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetLoginTxMsgCountOk() (*int64, bool) {
	if o == nil || o.LoginTxMsgCount == nil {
		return nil, false
	}
	return o.LoginTxMsgCount, true
}

// HasLoginTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasLoginTxMsgCount() bool {
	if o != nil && o.LoginTxMsgCount != nil {
		return true
	}

	return false
}

// SetLoginTxMsgCount gets a reference to the given int64 and assigns it to the LoginTxMsgCount field.
func (o *MsgVpnBridgeCounter) SetLoginTxMsgCount(v int64) {
	o.LoginTxMsgCount = &v
}

// GetMsgSpoolRxMsgCount returns the MsgSpoolRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetMsgSpoolRxMsgCount() int64 {
	if o == nil || o.MsgSpoolRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MsgSpoolRxMsgCount
}

// GetMsgSpoolRxMsgCountOk returns a tuple with the MsgSpoolRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetMsgSpoolRxMsgCountOk() (*int64, bool) {
	if o == nil || o.MsgSpoolRxMsgCount == nil {
		return nil, false
	}
	return o.MsgSpoolRxMsgCount, true
}

// HasMsgSpoolRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasMsgSpoolRxMsgCount() bool {
	if o != nil && o.MsgSpoolRxMsgCount != nil {
		return true
	}

	return false
}

// SetMsgSpoolRxMsgCount gets a reference to the given int64 and assigns it to the MsgSpoolRxMsgCount field.
func (o *MsgVpnBridgeCounter) SetMsgSpoolRxMsgCount(v int64) {
	o.MsgSpoolRxMsgCount = &v
}

// GetRxByteCount returns the RxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetRxByteCount() int64 {
	if o == nil || o.RxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.RxByteCount
}

// GetRxByteCountOk returns a tuple with the RxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetRxByteCountOk() (*int64, bool) {
	if o == nil || o.RxByteCount == nil {
		return nil, false
	}
	return o.RxByteCount, true
}

// HasRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasRxByteCount() bool {
	if o != nil && o.RxByteCount != nil {
		return true
	}

	return false
}

// SetRxByteCount gets a reference to the given int64 and assigns it to the RxByteCount field.
func (o *MsgVpnBridgeCounter) SetRxByteCount(v int64) {
	o.RxByteCount = &v
}

// GetRxMsgCount returns the RxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetRxMsgCount() int64 {
	if o == nil || o.RxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RxMsgCount
}

// GetRxMsgCountOk returns a tuple with the RxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetRxMsgCountOk() (*int64, bool) {
	if o == nil || o.RxMsgCount == nil {
		return nil, false
	}
	return o.RxMsgCount, true
}

// HasRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasRxMsgCount() bool {
	if o != nil && o.RxMsgCount != nil {
		return true
	}

	return false
}

// SetRxMsgCount gets a reference to the given int64 and assigns it to the RxMsgCount field.
func (o *MsgVpnBridgeCounter) SetRxMsgCount(v int64) {
	o.RxMsgCount = &v
}

// GetTxByteCount returns the TxByteCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetTxByteCount() int64 {
	if o == nil || o.TxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.TxByteCount
}

// GetTxByteCountOk returns a tuple with the TxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetTxByteCountOk() (*int64, bool) {
	if o == nil || o.TxByteCount == nil {
		return nil, false
	}
	return o.TxByteCount, true
}

// HasTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasTxByteCount() bool {
	if o != nil && o.TxByteCount != nil {
		return true
	}

	return false
}

// SetTxByteCount gets a reference to the given int64 and assigns it to the TxByteCount field.
func (o *MsgVpnBridgeCounter) SetTxByteCount(v int64) {
	o.TxByteCount = &v
}

// GetTxMsgCount returns the TxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnBridgeCounter) GetTxMsgCount() int64 {
	if o == nil || o.TxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.TxMsgCount
}

// GetTxMsgCountOk returns a tuple with the TxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridgeCounter) GetTxMsgCountOk() (*int64, bool) {
	if o == nil || o.TxMsgCount == nil {
		return nil, false
	}
	return o.TxMsgCount, true
}

// HasTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnBridgeCounter) HasTxMsgCount() bool {
	if o != nil && o.TxMsgCount != nil {
		return true
	}

	return false
}

// SetTxMsgCount gets a reference to the given int64 and assigns it to the TxMsgCount field.
func (o *MsgVpnBridgeCounter) SetTxMsgCount(v int64) {
	o.TxMsgCount = &v
}

func (o MsgVpnBridgeCounter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.LoginRxMsgCount != nil {
		toSerialize["loginRxMsgCount"] = o.LoginRxMsgCount
	}
	if o.LoginTxMsgCount != nil {
		toSerialize["loginTxMsgCount"] = o.LoginTxMsgCount
	}
	if o.MsgSpoolRxMsgCount != nil {
		toSerialize["msgSpoolRxMsgCount"] = o.MsgSpoolRxMsgCount
	}
	if o.RxByteCount != nil {
		toSerialize["rxByteCount"] = o.RxByteCount
	}
	if o.RxMsgCount != nil {
		toSerialize["rxMsgCount"] = o.RxMsgCount
	}
	if o.TxByteCount != nil {
		toSerialize["txByteCount"] = o.TxByteCount
	}
	if o.TxMsgCount != nil {
		toSerialize["txMsgCount"] = o.TxMsgCount
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnBridgeCounter struct {
	value *MsgVpnBridgeCounter
	isSet bool
}

func (v NullableMsgVpnBridgeCounter) Get() *MsgVpnBridgeCounter {
	return v.value
}

func (v *NullableMsgVpnBridgeCounter) Set(val *MsgVpnBridgeCounter) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnBridgeCounter) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnBridgeCounter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnBridgeCounter(val *MsgVpnBridgeCounter) *NullableMsgVpnBridgeCounter {
	return &NullableMsgVpnBridgeCounter{value: val, isSet: true}
}

func (v NullableMsgVpnBridgeCounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnBridgeCounter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
