/*
 * SEMP (Solace Element Management Protocol)
 *
 * SEMP (starting in `v2`, see note 1) is a RESTful API for configuring, monitoring, and administering a Solace PubSub+ broker.  SEMP uses URIs to address manageable **resources** of the Solace PubSub+ broker. Resources are individual **objects**, **collections** of objects, or (exclusively in the action API) **actions**. This document applies to the following API:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Configuration|/SEMP/v2/config|Reading and writing config state|See note 2    The following APIs are also available:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Action|/SEMP/v2/action|Performing actions|See note 2 Monitoring|/SEMP/v2/monitor|Querying operational parameters|See note 2    Resources are always nouns, with individual objects being singular and collections being plural.  Objects within a collection are identified by an `obj-id`, which follows the collection name with the form `collection-name/obj-id`.  Actions within an object are identified by an `action-id`, which follows the object name with the form `obj-id/action-id`.  Some examples:  ``` /SEMP/v2/config/msgVpns                        ; MsgVpn collection /SEMP/v2/config/msgVpns/a                      ; MsgVpn object named \"a\" /SEMP/v2/config/msgVpns/a/queues               ; Queue collection in MsgVpn \"a\" /SEMP/v2/config/msgVpns/a/queues/b             ; Queue object named \"b\" in MsgVpn \"a\" /SEMP/v2/action/msgVpns/a/queues/b/startReplay ; Action that starts a replay on Queue \"b\" in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients             ; Client collection in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients/c           ; Client object named \"c\" in MsgVpn \"a\" ```  ## Collection Resources  Collections are unordered lists of objects (unless described as otherwise), and are described by JSON arrays. Each item in the array represents an object in the same manner as the individual object would normally be represented. In the configuration API, the creation of a new object is done through its collection resource.  ## Object and Action Resources  Objects are composed of attributes, actions, collections, and other objects. They are described by JSON objects as name/value pairs. The collections and actions of an object are not contained directly in the object's JSON content; rather the content includes an attribute containing a URI which points to the collections and actions. These contained resources must be managed through this URI. At a minimum, every object has one or more identifying attributes, and its own `uri` attribute which contains the URI pointing to itself.  Actions are also composed of attributes, and are described by JSON objects as name/value pairs. Unlike objects, however, they are not members of a collection and cannot be retrieved, only performed. Actions only exist in the action API.  Attributes in an object or action may have any combination of the following properties:   Property|Meaning|Comments :---|:---|:--- Identifying|Attribute is involved in unique identification of the object, and appears in its URI| Required|Attribute must be provided in the request| Read-Only|Attribute can only be read, not written.|See note 3 Write-Only|Attribute can only be written, not read, unless the attribute is also opaque|See the documentation for the opaque property Requires-Disable|Attribute can only be changed when object is disabled| Deprecated|Attribute is deprecated, and will disappear in the next SEMP version| Opaque|Attribute can be set or retrieved in opaque form when the `opaquePassword` query parameter is present|See the `opaquePassword` query parameter documentation    In some requests, certain attributes may only be provided in certain combinations with other attributes:   Relationship|Meaning :---|:--- Requires|Attribute may only be changed by a request if a particular attribute or combination of attributes is also provided in the request Conflicts|Attribute may only be provided in a request if a particular attribute or combination of attributes is not also provided in the request    In the monitoring API, any non-identifying attribute may not be returned in a GET.  ## HTTP Methods  The following HTTP methods manipulate resources in accordance with these general principles. Note that some methods are only used in certain APIs:   Method|Resource|Meaning|Request Body|Response Body|Missing Request Attributes :---|:---|:---|:---|:---|:--- POST|Collection|Create object|Initial attribute values|Object attributes and metadata|Set to default PUT|Object|Create or replace object (see note 5)|New attribute values|Object attributes and metadata|Set to default, with certain exceptions (see note 4) PUT|Action|Performs action|Action arguments|Action metadata|N/A PATCH|Object|Update object|New attribute values|Object attributes and metadata|unchanged DELETE|Object|Delete object|Empty|Object metadata|N/A GET|Object|Get object|Empty|Object attributes and metadata|N/A GET|Collection|Get collection|Empty|Object attributes and collection metadata|N/A    ## Common Query Parameters  The following are some common query parameters that are supported by many method/URI combinations. Individual URIs may document additional parameters. Note that multiple query parameters can be used together in a single URI, separated by the ampersand character. For example:  ``` ; Request for the MsgVpns collection using two hypothetical query parameters ; \"q1\" and \"q2\" with values \"val1\" and \"val2\" respectively /SEMP/v2/config/msgVpns?q1=val1&q2=val2 ```  ### select  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. Use this query parameter to limit the size of the returned data for each returned object, return only those fields that are desired, or exclude fields that are not desired.  The value of `select` is a comma-separated list of attribute names. If the list contains attribute names that are not prefaced by `-`, only those attributes are included in the response. If the list contains attribute names that are prefaced by `-`, those attributes are excluded from the response. If the list contains both types, then the difference of the first set of attributes and the second set of attributes is returned. If the list is empty (i.e. `select=`), no attributes are returned.  All attributes that are prefaced by `-` must follow all attributes that are not prefaced by `-`. In addition, each attribute name in the list must match at least one attribute in the object.  Names may include the `*` wildcard (zero or more characters). Nested attribute names are supported using periods (e.g. `parentName.childName`).  Some examples:  ``` ; List of all MsgVpn names /SEMP/v2/config/msgVpns?select=msgVpnName ; List of all MsgVpn and their attributes except for their names /SEMP/v2/config/msgVpns?select=-msgVpnName ; Authentication attributes of MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance?select=authentication* ; All attributes of MsgVpn \"finance\" except for authentication attributes /SEMP/v2/config/msgVpns/finance?select=-authentication* ; Access related attributes of Queue \"orderQ\" of MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance/queues/orderQ?select=owner,permission ```  ### where  Include in the response only objects where certain conditions are true. Use this query parameter to limit which objects are returned to those whose attribute values meet the given conditions.  The value of `where` is a comma-separated list of expressions. All expressions must be true for the object to be included in the response. Each expression takes the form:  ``` expression  = attribute-name OP value OP          = '==' | '!=' | '&lt;' | '&gt;' | '&lt;=' | '&gt;=' ```  `value` may be a number, string, `true`, or `false`, as appropriate for the type of `attribute-name`. Greater-than and less-than comparisons only work for numbers. A `*` in a string `value` is interpreted as a wildcard (zero or more characters). Some examples:  ``` ; Only enabled MsgVpns /SEMP/v2/config/msgVpns?where=enabled==true ; Only MsgVpns using basic non-LDAP authentication /SEMP/v2/config/msgVpns?where=authenticationBasicEnabled==true,authenticationBasicType!=ldap ; Only MsgVpns that allow more than 100 client connections /SEMP/v2/config/msgVpns?where=maxConnectionCount>100 ; Only MsgVpns with msgVpnName starting with \"B\": /SEMP/v2/config/msgVpns?where=msgVpnName==B* ```  ### count  Limit the count of objects in the response. This can be useful to limit the size of the response for large collections. The minimum value for `count` is `1` and the default is `10`. There is also a per-collection maximum value to limit request handling time. For example:  ``` ; Up to 25 MsgVpns /SEMP/v2/config/msgVpns?count=25 ```  ### cursor  The cursor, or position, for the next page of objects. Cursors are opaque data that should not be created or interpreted by SEMP clients, and should only be used as described below.  When a request is made for a collection and there may be additional objects available for retrieval that are not included in the initial response, the response will include a `cursorQuery` field containing a cursor. The value of this field can be specified in the `cursor` query parameter of a subsequent request to retrieve the next page of objects. For convenience, an appropriate URI is constructed automatically by the broker and included in the `nextPageUri` field of the response. This URI can be used directly to retrieve the next page of objects.  ### opaquePassword  Attributes with the opaque property are also write-only and so cannot normally be retrieved in a GET. However, when a password is provided in the `opaquePassword` query parameter, attributes with the opaque property are retrieved in a GET in opaque form, encrypted with this password. The query parameter can also be used on a POST, PATCH, or PUT to set opaque attributes using opaque attribute values retrieved in a GET, so long as:  1. the same password that was used to retrieve the opaque attribute values is provided; and  2. the broker to which the request is being sent has the same major and minor SEMP version as the broker that produced the opaque attribute values.  The password provided in the query parameter must be a minimum of 8 characters and a maximum of 128 characters.  The query parameter can only be used in the configuration API, and only over HTTPS.  ## Authentication  When a client makes its first SEMPv2 request, it must supply a username and password using HTTP Basic authentication.  If authentication is successful, the broker returns a cookie containing a session key. The client can omit the username and password from subsequent requests, because the broker now uses the session cookie for authentication instead. When the session expires or is deleted, the client must provide the username and password again, and the broker creates a new session.  There are a limited number of session slots available on the broker. The broker returns 529 No SEMP Session Available if it is not able to allocate a session. For this reason, all clients that use SEMPv2 should support cookies.  If certain attributes—such as a user's password—are changed, the broker automatically deletes the affected sessions. These attributes are documented below. However, changes in external user configuration data stored on a RADIUS or LDAP server do not trigger the broker to delete the associated session(s), therefore you must do this manually, if required.  A client can retrieve its current session information using the /about/user endpoint, delete its own session using the /about/user/logout endpoint, and manage all sessions using the /sessions endpoint.  ## Help  Visit [our website](https://solace.com) to learn more about Solace.  You can also download the SEMP API specifications by clicking [here](https://solace.com/downloads/).  If you need additional support, please contact us at [support@solace.com](mailto:support@solace.com).  ## Notes  Note|Description :---:|:--- 1|This specification defines SEMP starting in \"v2\", and not the original SEMP \"v1\" interface. Request and response formats between \"v1\" and \"v2\" are entirely incompatible, although both protocols share a common port configuration on the Solace PubSub+ broker. They are differentiated by the initial portion of the URI path, one of either \"/SEMP/\" or \"/SEMP/v2/\" 2|This API is partially implemented. Only a subset of all objects are available. 3|Read-only attributes may appear in POST and PUT/PATCH requests. However, if a read-only attribute is not marked as identifying, it will be ignored during a PUT/PATCH. 4|On a PUT, if the SEMP user is not authorized to modify the attribute, its value is left unchanged rather than set to default. In addition, the values of write-only attributes are not set to their defaults on a PUT, except in the following two cases: there is a mutual requires relationship with another non-write-only attribute, both attributes are absent from the request, and the non-write-only attribute is not currently set to its default value; or the attribute is also opaque and the `opaquePassword` query parameter is provided in the request. 5|On a PUT, if the object does not exist, it is created first.
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
	// The name of the Bridge.
	BridgeName *string `json:"bridgeName,omitempty"`
	// The virtual router of the Bridge. The allowed values and their meaning are:  <pre> \"primary\" - The Bridge is used for the primary virtual router. \"backup\" - The Bridge is used for the backup virtual router. \"auto\" - The Bridge is automatically assigned a virtual router at creation, depending on the broker's active-standby role. </pre>
	BridgeVirtualRouter *string `json:"bridgeVirtualRouter,omitempty"`
	// Enable or disable the Bridge. The default value is `false`.
	Enabled *bool `json:"enabled,omitempty"`
	// The maximum time-to-live (TTL) in hops. Messages are discarded if their TTL exceeds this value. The default value is `8`.
	MaxTtl *int64 `json:"maxTtl,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// The Client Username the Bridge uses to login to the remote Message VPN. The default value is `\"\"`.
	RemoteAuthenticationBasicClientUsername *string `json:"remoteAuthenticationBasicClientUsername,omitempty"`
	// The password for the Client Username. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is `\"\"`.
	RemoteAuthenticationBasicPassword *string `json:"remoteAuthenticationBasicPassword,omitempty"`
	// The PEM formatted content for the client certificate used by the Bridge to login to the remote Message VPN. It must consist of a private key and between one and three certificates comprising the certificate trust chain. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is `\"\"`. Available since 2.9.
	RemoteAuthenticationClientCertContent *string `json:"remoteAuthenticationClientCertContent,omitempty"`
	// The password for the client certificate. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is `\"\"`. Available since 2.9.
	RemoteAuthenticationClientCertPassword *string `json:"remoteAuthenticationClientCertPassword,omitempty"`
	// The authentication scheme for the remote Message VPN. The default value is `\"basic\"`. The allowed values and their meaning are:  <pre> \"basic\" - Basic Authentication Scheme (via username and password). \"client-certificate\" - Client Certificate Authentication Scheme (via certificate file or content). </pre>
	RemoteAuthenticationScheme *string `json:"remoteAuthenticationScheme,omitempty"`
	// The maximum number of retry attempts to establish a connection to the remote Message VPN. A value of 0 means to retry forever. The default value is `0`.
	RemoteConnectionRetryCount *int64 `json:"remoteConnectionRetryCount,omitempty"`
	// The number of seconds the broker waits for the bridge connection to be established before attempting a new connection. The default value is `3`.
	RemoteConnectionRetryDelay *int64 `json:"remoteConnectionRetryDelay,omitempty"`
	// The priority for deliver-to-one (DTO) messages transmitted from the remote Message VPN. The default value is `\"p1\"`. The allowed values and their meaning are:  <pre> \"p1\" - The 1st or highest priority. \"p2\" - The 2nd highest priority. \"p3\" - The 3rd highest priority. \"p4\" - The 4th highest priority. \"da\" - Ignore priority and deliver always. </pre>
	RemoteDeliverToOnePriority *string `json:"remoteDeliverToOnePriority,omitempty"`
	// The colon-separated list of cipher suites supported for TLS connections to the remote Message VPN. The value \"default\" implies all supported suites ordered from most secure to least secure. The default value is `\"default\"`.
	TlsCipherSuiteList *string `json:"tlsCipherSuiteList,omitempty"`
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

// GetRemoteAuthenticationBasicPassword returns the RemoteAuthenticationBasicPassword field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteAuthenticationBasicPassword() string {
	if o == nil || o.RemoteAuthenticationBasicPassword == nil {
		var ret string
		return ret
	}
	return *o.RemoteAuthenticationBasicPassword
}

// GetRemoteAuthenticationBasicPasswordOk returns a tuple with the RemoteAuthenticationBasicPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteAuthenticationBasicPasswordOk() (*string, bool) {
	if o == nil || o.RemoteAuthenticationBasicPassword == nil {
		return nil, false
	}
	return o.RemoteAuthenticationBasicPassword, true
}

// HasRemoteAuthenticationBasicPassword returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteAuthenticationBasicPassword() bool {
	if o != nil && o.RemoteAuthenticationBasicPassword != nil {
		return true
	}

	return false
}

// SetRemoteAuthenticationBasicPassword gets a reference to the given string and assigns it to the RemoteAuthenticationBasicPassword field.
func (o *MsgVpnBridge) SetRemoteAuthenticationBasicPassword(v string) {
	o.RemoteAuthenticationBasicPassword = &v
}

// GetRemoteAuthenticationClientCertContent returns the RemoteAuthenticationClientCertContent field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteAuthenticationClientCertContent() string {
	if o == nil || o.RemoteAuthenticationClientCertContent == nil {
		var ret string
		return ret
	}
	return *o.RemoteAuthenticationClientCertContent
}

// GetRemoteAuthenticationClientCertContentOk returns a tuple with the RemoteAuthenticationClientCertContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteAuthenticationClientCertContentOk() (*string, bool) {
	if o == nil || o.RemoteAuthenticationClientCertContent == nil {
		return nil, false
	}
	return o.RemoteAuthenticationClientCertContent, true
}

// HasRemoteAuthenticationClientCertContent returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteAuthenticationClientCertContent() bool {
	if o != nil && o.RemoteAuthenticationClientCertContent != nil {
		return true
	}

	return false
}

// SetRemoteAuthenticationClientCertContent gets a reference to the given string and assigns it to the RemoteAuthenticationClientCertContent field.
func (o *MsgVpnBridge) SetRemoteAuthenticationClientCertContent(v string) {
	o.RemoteAuthenticationClientCertContent = &v
}

// GetRemoteAuthenticationClientCertPassword returns the RemoteAuthenticationClientCertPassword field value if set, zero value otherwise.
func (o *MsgVpnBridge) GetRemoteAuthenticationClientCertPassword() string {
	if o == nil || o.RemoteAuthenticationClientCertPassword == nil {
		var ret string
		return ret
	}
	return *o.RemoteAuthenticationClientCertPassword
}

// GetRemoteAuthenticationClientCertPasswordOk returns a tuple with the RemoteAuthenticationClientCertPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnBridge) GetRemoteAuthenticationClientCertPasswordOk() (*string, bool) {
	if o == nil || o.RemoteAuthenticationClientCertPassword == nil {
		return nil, false
	}
	return o.RemoteAuthenticationClientCertPassword, true
}

// HasRemoteAuthenticationClientCertPassword returns a boolean if a field has been set.
func (o *MsgVpnBridge) HasRemoteAuthenticationClientCertPassword() bool {
	if o != nil && o.RemoteAuthenticationClientCertPassword != nil {
		return true
	}

	return false
}

// SetRemoteAuthenticationClientCertPassword gets a reference to the given string and assigns it to the RemoteAuthenticationClientCertPassword field.
func (o *MsgVpnBridge) SetRemoteAuthenticationClientCertPassword(v string) {
	o.RemoteAuthenticationClientCertPassword = &v
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

func (o MsgVpnBridge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BridgeName != nil {
		toSerialize["bridgeName"] = o.BridgeName
	}
	if o.BridgeVirtualRouter != nil {
		toSerialize["bridgeVirtualRouter"] = o.BridgeVirtualRouter
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MaxTtl != nil {
		toSerialize["maxTtl"] = o.MaxTtl
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.RemoteAuthenticationBasicClientUsername != nil {
		toSerialize["remoteAuthenticationBasicClientUsername"] = o.RemoteAuthenticationBasicClientUsername
	}
	if o.RemoteAuthenticationBasicPassword != nil {
		toSerialize["remoteAuthenticationBasicPassword"] = o.RemoteAuthenticationBasicPassword
	}
	if o.RemoteAuthenticationClientCertContent != nil {
		toSerialize["remoteAuthenticationClientCertContent"] = o.RemoteAuthenticationClientCertContent
	}
	if o.RemoteAuthenticationClientCertPassword != nil {
		toSerialize["remoteAuthenticationClientCertPassword"] = o.RemoteAuthenticationClientCertPassword
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
	if o.TlsCipherSuiteList != nil {
		toSerialize["tlsCipherSuiteList"] = o.TlsCipherSuiteList
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
