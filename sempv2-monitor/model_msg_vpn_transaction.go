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

// MsgVpnTransaction struct for MsgVpnTransaction
type MsgVpnTransaction struct {
	// The identifier (ID) of the Client.
	ClientId *int32 `json:"clientId,omitempty"`
	// The name of the Client.
	ClientName *string `json:"clientName,omitempty"`
	// The username of the Client.
	ClientUsername *string `json:"clientUsername,omitempty"`
	// The number of seconds before an idle Transaction may be automatically rolled back and freed.
	IdleTimeout *int32 `json:"idleTimeout,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// Indicates whether the Transaction is replicated.
	Replicated *bool `json:"replicated,omitempty"`
	// The name of the Transacted Session for the Transaction.
	SessionName *string `json:"sessionName,omitempty"`
	// The state of the Transaction. The allowed values and their meaning are:  <pre> \"active\" - The Transaction was started. \"suspended\" - The Transaction was suspended. \"idle\" - The Transaction was ended. \"prepared\" - The Transaction was prepared. \"complete\" - The Transaction was committed or rolled back. </pre>
	State *string `json:"state,omitempty"`
	// The number of seconds the Transaction has remained in the current state.
	TimeInState *int32 `json:"timeInState,omitempty"`
	// The type of Transaction. The allowed values and their meaning are:  <pre> \"xa\" - The Transaction is an XA Transaction. \"local\" - The Transaction is a local Transaction. </pre>
	Type *string `json:"type,omitempty"`
	// The identifier (ID) of the Transaction.
	Xid *string `json:"xid,omitempty"`
}

// NewMsgVpnTransaction instantiates a new MsgVpnTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnTransaction() *MsgVpnTransaction {
	this := MsgVpnTransaction{}
	return &this
}

// NewMsgVpnTransactionWithDefaults instantiates a new MsgVpnTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnTransactionWithDefaults() *MsgVpnTransaction {
	this := MsgVpnTransaction{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetClientId() int32 {
	if o == nil || o.ClientId == nil {
		var ret int32
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetClientIdOk() (*int32, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int32 and assigns it to the ClientId field.
func (o *MsgVpnTransaction) SetClientId(v int32) {
	o.ClientId = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *MsgVpnTransaction) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientUsername returns the ClientUsername field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetClientUsername() string {
	if o == nil || o.ClientUsername == nil {
		var ret string
		return ret
	}
	return *o.ClientUsername
}

// GetClientUsernameOk returns a tuple with the ClientUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetClientUsernameOk() (*string, bool) {
	if o == nil || o.ClientUsername == nil {
		return nil, false
	}
	return o.ClientUsername, true
}

// HasClientUsername returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasClientUsername() bool {
	if o != nil && o.ClientUsername != nil {
		return true
	}

	return false
}

// SetClientUsername gets a reference to the given string and assigns it to the ClientUsername field.
func (o *MsgVpnTransaction) SetClientUsername(v string) {
	o.ClientUsername = &v
}

// GetIdleTimeout returns the IdleTimeout field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetIdleTimeout() int32 {
	if o == nil || o.IdleTimeout == nil {
		var ret int32
		return ret
	}
	return *o.IdleTimeout
}

// GetIdleTimeoutOk returns a tuple with the IdleTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetIdleTimeoutOk() (*int32, bool) {
	if o == nil || o.IdleTimeout == nil {
		return nil, false
	}
	return o.IdleTimeout, true
}

// HasIdleTimeout returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasIdleTimeout() bool {
	if o != nil && o.IdleTimeout != nil {
		return true
	}

	return false
}

// SetIdleTimeout gets a reference to the given int32 and assigns it to the IdleTimeout field.
func (o *MsgVpnTransaction) SetIdleTimeout(v int32) {
	o.IdleTimeout = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnTransaction) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetReplicated returns the Replicated field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetReplicated() bool {
	if o == nil || o.Replicated == nil {
		var ret bool
		return ret
	}
	return *o.Replicated
}

// GetReplicatedOk returns a tuple with the Replicated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetReplicatedOk() (*bool, bool) {
	if o == nil || o.Replicated == nil {
		return nil, false
	}
	return o.Replicated, true
}

// HasReplicated returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasReplicated() bool {
	if o != nil && o.Replicated != nil {
		return true
	}

	return false
}

// SetReplicated gets a reference to the given bool and assigns it to the Replicated field.
func (o *MsgVpnTransaction) SetReplicated(v bool) {
	o.Replicated = &v
}

// GetSessionName returns the SessionName field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetSessionName() string {
	if o == nil || o.SessionName == nil {
		var ret string
		return ret
	}
	return *o.SessionName
}

// GetSessionNameOk returns a tuple with the SessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetSessionNameOk() (*string, bool) {
	if o == nil || o.SessionName == nil {
		return nil, false
	}
	return o.SessionName, true
}

// HasSessionName returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasSessionName() bool {
	if o != nil && o.SessionName != nil {
		return true
	}

	return false
}

// SetSessionName gets a reference to the given string and assigns it to the SessionName field.
func (o *MsgVpnTransaction) SetSessionName(v string) {
	o.SessionName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MsgVpnTransaction) SetState(v string) {
	o.State = &v
}

// GetTimeInState returns the TimeInState field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetTimeInState() int32 {
	if o == nil || o.TimeInState == nil {
		var ret int32
		return ret
	}
	return *o.TimeInState
}

// GetTimeInStateOk returns a tuple with the TimeInState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetTimeInStateOk() (*int32, bool) {
	if o == nil || o.TimeInState == nil {
		return nil, false
	}
	return o.TimeInState, true
}

// HasTimeInState returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasTimeInState() bool {
	if o != nil && o.TimeInState != nil {
		return true
	}

	return false
}

// SetTimeInState gets a reference to the given int32 and assigns it to the TimeInState field.
func (o *MsgVpnTransaction) SetTimeInState(v int32) {
	o.TimeInState = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MsgVpnTransaction) SetType(v string) {
	o.Type = &v
}

// GetXid returns the Xid field value if set, zero value otherwise.
func (o *MsgVpnTransaction) GetXid() string {
	if o == nil || o.Xid == nil {
		var ret string
		return ret
	}
	return *o.Xid
}

// GetXidOk returns a tuple with the Xid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTransaction) GetXidOk() (*string, bool) {
	if o == nil || o.Xid == nil {
		return nil, false
	}
	return o.Xid, true
}

// HasXid returns a boolean if a field has been set.
func (o *MsgVpnTransaction) HasXid() bool {
	if o != nil && o.Xid != nil {
		return true
	}

	return false
}

// SetXid gets a reference to the given string and assigns it to the Xid field.
func (o *MsgVpnTransaction) SetXid(v string) {
	o.Xid = &v
}

func (o MsgVpnTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.ClientUsername != nil {
		toSerialize["clientUsername"] = o.ClientUsername
	}
	if o.IdleTimeout != nil {
		toSerialize["idleTimeout"] = o.IdleTimeout
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.Replicated != nil {
		toSerialize["replicated"] = o.Replicated
	}
	if o.SessionName != nil {
		toSerialize["sessionName"] = o.SessionName
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.TimeInState != nil {
		toSerialize["timeInState"] = o.TimeInState
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Xid != nil {
		toSerialize["xid"] = o.Xid
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnTransaction struct {
	value *MsgVpnTransaction
	isSet bool
}

func (v NullableMsgVpnTransaction) Get() *MsgVpnTransaction {
	return v.value
}

func (v *NullableMsgVpnTransaction) Set(val *MsgVpnTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnTransaction(val *MsgVpnTransaction) *NullableMsgVpnTransaction {
	return &NullableMsgVpnTransaction{value: val, isSet: true}
}

func (v NullableMsgVpnTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
