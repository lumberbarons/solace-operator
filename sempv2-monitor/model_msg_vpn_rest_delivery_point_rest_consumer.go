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

// MsgVpnRestDeliveryPointRestConsumer struct for MsgVpnRestDeliveryPointRestConsumer
type MsgVpnRestDeliveryPointRestConsumer struct {
	// The username that the REST Consumer will use to login to the REST host.
	AuthenticationHttpBasicUsername *string `json:"authenticationHttpBasicUsername,omitempty"`
	// The authentication header name. Available since 2.15.
	AuthenticationHttpHeaderName *string `json:"authenticationHttpHeaderName,omitempty"`
	// The OAuth client ID. Available since 2.19.
	AuthenticationOauthClientId *string `json:"authenticationOauthClientId,omitempty"`
	// The reason for the most recent OAuth token retrieval failure. Available since 2.19.
	AuthenticationOauthClientLastFailureReason *string `json:"authenticationOauthClientLastFailureReason,omitempty"`
	// The time of the last OAuth token retrieval failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.19.
	AuthenticationOauthClientLastFailureTime *int32 `json:"authenticationOauthClientLastFailureTime,omitempty"`
	// The OAuth scope. Available since 2.19.
	AuthenticationOauthClientScope *string `json:"authenticationOauthClientScope,omitempty"`
	// The OAuth token endpoint URL that the REST Consumer will use to request a token for login to the REST host. Must begin with \"https\". Available since 2.19.
	AuthenticationOauthClientTokenEndpoint *string `json:"authenticationOauthClientTokenEndpoint,omitempty"`
	// The validity duration of the OAuth token. Available since 2.19.
	AuthenticationOauthClientTokenLifetime *int64 `json:"authenticationOauthClientTokenLifetime,omitempty"`
	// The time at which the broker requested the token from the OAuth token endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.19.
	AuthenticationOauthClientTokenRetrievedTime *int32 `json:"authenticationOauthClientTokenRetrievedTime,omitempty"`
	// The current state of the current OAuth token. The allowed values and their meaning are:  <pre> \"valid\" - The token is valid. \"invalid\" - The token is invalid. </pre>  Available since 2.19.
	AuthenticationOauthClientTokenState *string `json:"authenticationOauthClientTokenState,omitempty"`
	// The reason for the most recent OAuth token retrieval failure. Available since 2.21.
	AuthenticationOauthJwtLastFailureReason *string `json:"authenticationOauthJwtLastFailureReason,omitempty"`
	// The time of the last OAuth token retrieval failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21.
	AuthenticationOauthJwtLastFailureTime *int32 `json:"authenticationOauthJwtLastFailureTime,omitempty"`
	// The OAuth token endpoint URL that the REST Consumer will use to request a token for login to the REST host. Available since 2.21.
	AuthenticationOauthJwtTokenEndpoint *string `json:"authenticationOauthJwtTokenEndpoint,omitempty"`
	// The validity duration of the OAuth token. Available since 2.21.
	AuthenticationOauthJwtTokenLifetime *int64 `json:"authenticationOauthJwtTokenLifetime,omitempty"`
	// The time at which the broker requested the token from the OAuth token endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21.
	AuthenticationOauthJwtTokenRetrievedTime *int32 `json:"authenticationOauthJwtTokenRetrievedTime,omitempty"`
	// The current state of the current OAuth token. The allowed values and their meaning are:  <pre> \"valid\" - The token is valid. \"invalid\" - The token is invalid. </pre>  Available since 2.21.
	AuthenticationOauthJwtTokenState *string `json:"authenticationOauthJwtTokenState,omitempty"`
	// The authentication scheme used by the REST Consumer to login to the REST host. The allowed values and their meaning are:  <pre> \"none\" - Login with no authentication. This may be useful for anonymous connections or when a REST Consumer does not require authentication. \"http-basic\" - Login with a username and optional password according to HTTP Basic authentication as per RFC2616. \"client-certificate\" - Login with a client TLS certificate as per RFC5246. Client certificate authentication is only available on TLS connections. \"http-header\" - Login with a specified HTTP header. \"oauth-client\" - Login with OAuth 2.0 client credentials. \"oauth-jwt\" - Login with OAuth (RFC 7523 JWT Profile). \"transparent\" - Login using the Authorization header from the message properties, if present. Transparent authentication passes along existing Authorization header metadata instead of discarding it. Note that if the message is coming from a REST producer, the REST service must be configured to forward the Authorization header. </pre>
	AuthenticationScheme *string                                     `json:"authenticationScheme,omitempty"`
	Counter              *MsgVpnRestDeliveryPointRestConsumerCounter `json:"counter,omitempty"`
	// Indicates whether the REST Consumer is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The HTTP method to use (POST or PUT). This is used only when operating in the REST service \"messaging\" mode and is ignored in \"gateway\" mode. The allowed values and their meaning are:  <pre> \"post\" - Use the POST HTTP method. \"put\" - Use the PUT HTTP method. </pre>  Available since 2.17.
	HttpMethod *string `json:"httpMethod,omitempty"`
	// The number of HTTP request messages transmitted to the REST Consumer to close the connection. Available since 2.13.
	HttpRequestConnectionCloseTxMsgCount *int64 `json:"httpRequestConnectionCloseTxMsgCount,omitempty"`
	// The number of HTTP request messages transmitted to the REST Consumer that are waiting for a response. Available since 2.13.
	HttpRequestOutstandingTxMsgCount *int64 `json:"httpRequestOutstandingTxMsgCount,omitempty"`
	// The number of HTTP request messages transmitted to the REST Consumer that have timed out. Available since 2.13.
	HttpRequestTimedOutTxMsgCount *int64 `json:"httpRequestTimedOutTxMsgCount,omitempty"`
	// The amount of HTTP request messages transmitted to the REST Consumer, in bytes (B). Available since 2.13.
	HttpRequestTxByteCount *int64 `json:"httpRequestTxByteCount,omitempty"`
	// The number of HTTP request messages transmitted to the REST Consumer. Available since 2.13.
	HttpRequestTxMsgCount *int64 `json:"httpRequestTxMsgCount,omitempty"`
	// The number of HTTP client/server error response messages received from the REST Consumer. Available since 2.13.
	HttpResponseErrorRxMsgCount *int64 `json:"httpResponseErrorRxMsgCount,omitempty"`
	// The amount of HTTP response messages received from the REST Consumer, in bytes (B). Available since 2.13.
	HttpResponseRxByteCount *int64 `json:"httpResponseRxByteCount,omitempty"`
	// The number of HTTP response messages received from the REST Consumer. Available since 2.13.
	HttpResponseRxMsgCount *int64 `json:"httpResponseRxMsgCount,omitempty"`
	// The number of HTTP successful response messages received from the REST Consumer. Available since 2.13.
	HttpResponseSuccessRxMsgCount *int64 `json:"httpResponseSuccessRxMsgCount,omitempty"`
	// The local endpoint at the time of the last connection failure.
	LastConnectionFailureLocalEndpoint *string `json:"lastConnectionFailureLocalEndpoint,omitempty"`
	// The reason for the last connection failure between local and remote endpoints.
	LastConnectionFailureReason *string `json:"lastConnectionFailureReason,omitempty"`
	// The remote endpoint at the time of the last connection failure.
	LastConnectionFailureRemoteEndpoint *string `json:"lastConnectionFailureRemoteEndpoint,omitempty"`
	// The timestamp of the last connection failure between local and remote endpoints. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastConnectionFailureTime *int32 `json:"lastConnectionFailureTime,omitempty"`
	// The reason for the last REST Consumer failure.
	LastFailureReason *string `json:"lastFailureReason,omitempty"`
	// The timestamp of the last REST Consumer failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastFailureTime *int32 `json:"lastFailureTime,omitempty"`
	// The interface that will be used for all outgoing connections associated with the REST Consumer. When unspecified, an interface is automatically chosen.
	LocalInterface *string `json:"localInterface,omitempty"`
	// The maximum amount of time (in seconds) to wait for an HTTP POST response from the REST Consumer. Once this time is exceeded, the TCP connection is reset.
	MaxPostWaitTime *int32 `json:"maxPostWaitTime,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// The number of concurrent TCP connections open to the REST Consumer.
	OutgoingConnectionCount *int32 `json:"outgoingConnectionCount,omitempty"`
	// The IP address or DNS name for the REST Consumer.
	RemoteHost *string `json:"remoteHost,omitempty"`
	// The number of outgoing connections for the REST Consumer that are up.
	RemoteOutgoingConnectionUpCount *int64 `json:"remoteOutgoingConnectionUpCount,omitempty"`
	// The port associated with the host of the REST Consumer.
	RemotePort *int64 `json:"remotePort,omitempty"`
	// The name of the REST Consumer.
	RestConsumerName *string `json:"restConsumerName,omitempty"`
	// The name of the REST Delivery Point.
	RestDeliveryPointName *string `json:"restDeliveryPointName,omitempty"`
	// The number of seconds that must pass before retrying the remote REST Consumer connection.
	RetryDelay *int32 `json:"retryDelay,omitempty"`
	// The colon-separated list of cipher suites the REST Consumer uses in its encrypted connection. The value `\"default\"` implies all supported suites ordered from most secure to least secure. The list of default cipher suites is available in the `tlsCipherSuiteMsgBackboneDefaultList` attribute of the Broker object in the Monitoring API. The REST Consumer should choose the first suite from this list that it supports.
	TlsCipherSuiteList *string `json:"tlsCipherSuiteList,omitempty"`
	// Indicates whether encryption (TLS) is enabled for the REST Consumer.
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`
	// Indicates whether the operational state of the REST Consumer is up.
	Up *bool `json:"up,omitempty"`
}

// NewMsgVpnRestDeliveryPointRestConsumer instantiates a new MsgVpnRestDeliveryPointRestConsumer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnRestDeliveryPointRestConsumer() *MsgVpnRestDeliveryPointRestConsumer {
	this := MsgVpnRestDeliveryPointRestConsumer{}
	return &this
}

// NewMsgVpnRestDeliveryPointRestConsumerWithDefaults instantiates a new MsgVpnRestDeliveryPointRestConsumer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnRestDeliveryPointRestConsumerWithDefaults() *MsgVpnRestDeliveryPointRestConsumer {
	this := MsgVpnRestDeliveryPointRestConsumer{}
	return &this
}

// GetAuthenticationHttpBasicUsername returns the AuthenticationHttpBasicUsername field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpBasicUsername() string {
	if o == nil || o.AuthenticationHttpBasicUsername == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationHttpBasicUsername
}

// GetAuthenticationHttpBasicUsernameOk returns a tuple with the AuthenticationHttpBasicUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpBasicUsernameOk() (*string, bool) {
	if o == nil || o.AuthenticationHttpBasicUsername == nil {
		return nil, false
	}
	return o.AuthenticationHttpBasicUsername, true
}

// HasAuthenticationHttpBasicUsername returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationHttpBasicUsername() bool {
	if o != nil && o.AuthenticationHttpBasicUsername != nil {
		return true
	}

	return false
}

// SetAuthenticationHttpBasicUsername gets a reference to the given string and assigns it to the AuthenticationHttpBasicUsername field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationHttpBasicUsername(v string) {
	o.AuthenticationHttpBasicUsername = &v
}

// GetAuthenticationHttpHeaderName returns the AuthenticationHttpHeaderName field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpHeaderName() string {
	if o == nil || o.AuthenticationHttpHeaderName == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationHttpHeaderName
}

// GetAuthenticationHttpHeaderNameOk returns a tuple with the AuthenticationHttpHeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpHeaderNameOk() (*string, bool) {
	if o == nil || o.AuthenticationHttpHeaderName == nil {
		return nil, false
	}
	return o.AuthenticationHttpHeaderName, true
}

// HasAuthenticationHttpHeaderName returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationHttpHeaderName() bool {
	if o != nil && o.AuthenticationHttpHeaderName != nil {
		return true
	}

	return false
}

// SetAuthenticationHttpHeaderName gets a reference to the given string and assigns it to the AuthenticationHttpHeaderName field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationHttpHeaderName(v string) {
	o.AuthenticationHttpHeaderName = &v
}

// GetAuthenticationOauthClientId returns the AuthenticationOauthClientId field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientId() string {
	if o == nil || o.AuthenticationOauthClientId == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationOauthClientId
}

// GetAuthenticationOauthClientIdOk returns a tuple with the AuthenticationOauthClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientIdOk() (*string, bool) {
	if o == nil || o.AuthenticationOauthClientId == nil {
		return nil, false
	}
	return o.AuthenticationOauthClientId, true
}

// HasAuthenticationOauthClientId returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientId() bool {
	if o != nil && o.AuthenticationOauthClientId != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthClientId gets a reference to the given string and assigns it to the AuthenticationOauthClientId field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientId(v string) {
	o.AuthenticationOauthClientId = &v
}

// GetAuthenticationOauthClientLastFailureReason returns the AuthenticationOauthClientLastFailureReason field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientLastFailureReason() string {
	if o == nil || o.AuthenticationOauthClientLastFailureReason == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationOauthClientLastFailureReason
}

// GetAuthenticationOauthClientLastFailureReasonOk returns a tuple with the AuthenticationOauthClientLastFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientLastFailureReasonOk() (*string, bool) {
	if o == nil || o.AuthenticationOauthClientLastFailureReason == nil {
		return nil, false
	}
	return o.AuthenticationOauthClientLastFailureReason, true
}

// HasAuthenticationOauthClientLastFailureReason returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientLastFailureReason() bool {
	if o != nil && o.AuthenticationOauthClientLastFailureReason != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthClientLastFailureReason gets a reference to the given string and assigns it to the AuthenticationOauthClientLastFailureReason field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientLastFailureReason(v string) {
	o.AuthenticationOauthClientLastFailureReason = &v
}

// GetAuthenticationOauthClientLastFailureTime returns the AuthenticationOauthClientLastFailureTime field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientLastFailureTime() int32 {
	if o == nil || o.AuthenticationOauthClientLastFailureTime == nil {
		var ret int32
		return ret
	}
	return *o.AuthenticationOauthClientLastFailureTime
}

// GetAuthenticationOauthClientLastFailureTimeOk returns a tuple with the AuthenticationOauthClientLastFailureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientLastFailureTimeOk() (*int32, bool) {
	if o == nil || o.AuthenticationOauthClientLastFailureTime == nil {
		return nil, false
	}
	return o.AuthenticationOauthClientLastFailureTime, true
}

// HasAuthenticationOauthClientLastFailureTime returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientLastFailureTime() bool {
	if o != nil && o.AuthenticationOauthClientLastFailureTime != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthClientLastFailureTime gets a reference to the given int32 and assigns it to the AuthenticationOauthClientLastFailureTime field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientLastFailureTime(v int32) {
	o.AuthenticationOauthClientLastFailureTime = &v
}

// GetAuthenticationOauthClientScope returns the AuthenticationOauthClientScope field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientScope() string {
	if o == nil || o.AuthenticationOauthClientScope == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationOauthClientScope
}

// GetAuthenticationOauthClientScopeOk returns a tuple with the AuthenticationOauthClientScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientScopeOk() (*string, bool) {
	if o == nil || o.AuthenticationOauthClientScope == nil {
		return nil, false
	}
	return o.AuthenticationOauthClientScope, true
}

// HasAuthenticationOauthClientScope returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientScope() bool {
	if o != nil && o.AuthenticationOauthClientScope != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthClientScope gets a reference to the given string and assigns it to the AuthenticationOauthClientScope field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientScope(v string) {
	o.AuthenticationOauthClientScope = &v
}

// GetAuthenticationOauthClientTokenEndpoint returns the AuthenticationOauthClientTokenEndpoint field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenEndpoint() string {
	if o == nil || o.AuthenticationOauthClientTokenEndpoint == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationOauthClientTokenEndpoint
}

// GetAuthenticationOauthClientTokenEndpointOk returns a tuple with the AuthenticationOauthClientTokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenEndpointOk() (*string, bool) {
	if o == nil || o.AuthenticationOauthClientTokenEndpoint == nil {
		return nil, false
	}
	return o.AuthenticationOauthClientTokenEndpoint, true
}

// HasAuthenticationOauthClientTokenEndpoint returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientTokenEndpoint() bool {
	if o != nil && o.AuthenticationOauthClientTokenEndpoint != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthClientTokenEndpoint gets a reference to the given string and assigns it to the AuthenticationOauthClientTokenEndpoint field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientTokenEndpoint(v string) {
	o.AuthenticationOauthClientTokenEndpoint = &v
}

// GetAuthenticationOauthClientTokenLifetime returns the AuthenticationOauthClientTokenLifetime field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenLifetime() int64 {
	if o == nil || o.AuthenticationOauthClientTokenLifetime == nil {
		var ret int64
		return ret
	}
	return *o.AuthenticationOauthClientTokenLifetime
}

// GetAuthenticationOauthClientTokenLifetimeOk returns a tuple with the AuthenticationOauthClientTokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenLifetimeOk() (*int64, bool) {
	if o == nil || o.AuthenticationOauthClientTokenLifetime == nil {
		return nil, false
	}
	return o.AuthenticationOauthClientTokenLifetime, true
}

// HasAuthenticationOauthClientTokenLifetime returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientTokenLifetime() bool {
	if o != nil && o.AuthenticationOauthClientTokenLifetime != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthClientTokenLifetime gets a reference to the given int64 and assigns it to the AuthenticationOauthClientTokenLifetime field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientTokenLifetime(v int64) {
	o.AuthenticationOauthClientTokenLifetime = &v
}

// GetAuthenticationOauthClientTokenRetrievedTime returns the AuthenticationOauthClientTokenRetrievedTime field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenRetrievedTime() int32 {
	if o == nil || o.AuthenticationOauthClientTokenRetrievedTime == nil {
		var ret int32
		return ret
	}
	return *o.AuthenticationOauthClientTokenRetrievedTime
}

// GetAuthenticationOauthClientTokenRetrievedTimeOk returns a tuple with the AuthenticationOauthClientTokenRetrievedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenRetrievedTimeOk() (*int32, bool) {
	if o == nil || o.AuthenticationOauthClientTokenRetrievedTime == nil {
		return nil, false
	}
	return o.AuthenticationOauthClientTokenRetrievedTime, true
}

// HasAuthenticationOauthClientTokenRetrievedTime returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientTokenRetrievedTime() bool {
	if o != nil && o.AuthenticationOauthClientTokenRetrievedTime != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthClientTokenRetrievedTime gets a reference to the given int32 and assigns it to the AuthenticationOauthClientTokenRetrievedTime field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientTokenRetrievedTime(v int32) {
	o.AuthenticationOauthClientTokenRetrievedTime = &v
}

// GetAuthenticationOauthClientTokenState returns the AuthenticationOauthClientTokenState field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenState() string {
	if o == nil || o.AuthenticationOauthClientTokenState == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationOauthClientTokenState
}

// GetAuthenticationOauthClientTokenStateOk returns a tuple with the AuthenticationOauthClientTokenState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenStateOk() (*string, bool) {
	if o == nil || o.AuthenticationOauthClientTokenState == nil {
		return nil, false
	}
	return o.AuthenticationOauthClientTokenState, true
}

// HasAuthenticationOauthClientTokenState returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientTokenState() bool {
	if o != nil && o.AuthenticationOauthClientTokenState != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthClientTokenState gets a reference to the given string and assigns it to the AuthenticationOauthClientTokenState field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientTokenState(v string) {
	o.AuthenticationOauthClientTokenState = &v
}

// GetAuthenticationOauthJwtLastFailureReason returns the AuthenticationOauthJwtLastFailureReason field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtLastFailureReason() string {
	if o == nil || o.AuthenticationOauthJwtLastFailureReason == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationOauthJwtLastFailureReason
}

// GetAuthenticationOauthJwtLastFailureReasonOk returns a tuple with the AuthenticationOauthJwtLastFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtLastFailureReasonOk() (*string, bool) {
	if o == nil || o.AuthenticationOauthJwtLastFailureReason == nil {
		return nil, false
	}
	return o.AuthenticationOauthJwtLastFailureReason, true
}

// HasAuthenticationOauthJwtLastFailureReason returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtLastFailureReason() bool {
	if o != nil && o.AuthenticationOauthJwtLastFailureReason != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthJwtLastFailureReason gets a reference to the given string and assigns it to the AuthenticationOauthJwtLastFailureReason field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtLastFailureReason(v string) {
	o.AuthenticationOauthJwtLastFailureReason = &v
}

// GetAuthenticationOauthJwtLastFailureTime returns the AuthenticationOauthJwtLastFailureTime field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtLastFailureTime() int32 {
	if o == nil || o.AuthenticationOauthJwtLastFailureTime == nil {
		var ret int32
		return ret
	}
	return *o.AuthenticationOauthJwtLastFailureTime
}

// GetAuthenticationOauthJwtLastFailureTimeOk returns a tuple with the AuthenticationOauthJwtLastFailureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtLastFailureTimeOk() (*int32, bool) {
	if o == nil || o.AuthenticationOauthJwtLastFailureTime == nil {
		return nil, false
	}
	return o.AuthenticationOauthJwtLastFailureTime, true
}

// HasAuthenticationOauthJwtLastFailureTime returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtLastFailureTime() bool {
	if o != nil && o.AuthenticationOauthJwtLastFailureTime != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthJwtLastFailureTime gets a reference to the given int32 and assigns it to the AuthenticationOauthJwtLastFailureTime field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtLastFailureTime(v int32) {
	o.AuthenticationOauthJwtLastFailureTime = &v
}

// GetAuthenticationOauthJwtTokenEndpoint returns the AuthenticationOauthJwtTokenEndpoint field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenEndpoint() string {
	if o == nil || o.AuthenticationOauthJwtTokenEndpoint == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationOauthJwtTokenEndpoint
}

// GetAuthenticationOauthJwtTokenEndpointOk returns a tuple with the AuthenticationOauthJwtTokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenEndpointOk() (*string, bool) {
	if o == nil || o.AuthenticationOauthJwtTokenEndpoint == nil {
		return nil, false
	}
	return o.AuthenticationOauthJwtTokenEndpoint, true
}

// HasAuthenticationOauthJwtTokenEndpoint returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtTokenEndpoint() bool {
	if o != nil && o.AuthenticationOauthJwtTokenEndpoint != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthJwtTokenEndpoint gets a reference to the given string and assigns it to the AuthenticationOauthJwtTokenEndpoint field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtTokenEndpoint(v string) {
	o.AuthenticationOauthJwtTokenEndpoint = &v
}

// GetAuthenticationOauthJwtTokenLifetime returns the AuthenticationOauthJwtTokenLifetime field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenLifetime() int64 {
	if o == nil || o.AuthenticationOauthJwtTokenLifetime == nil {
		var ret int64
		return ret
	}
	return *o.AuthenticationOauthJwtTokenLifetime
}

// GetAuthenticationOauthJwtTokenLifetimeOk returns a tuple with the AuthenticationOauthJwtTokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenLifetimeOk() (*int64, bool) {
	if o == nil || o.AuthenticationOauthJwtTokenLifetime == nil {
		return nil, false
	}
	return o.AuthenticationOauthJwtTokenLifetime, true
}

// HasAuthenticationOauthJwtTokenLifetime returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtTokenLifetime() bool {
	if o != nil && o.AuthenticationOauthJwtTokenLifetime != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthJwtTokenLifetime gets a reference to the given int64 and assigns it to the AuthenticationOauthJwtTokenLifetime field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtTokenLifetime(v int64) {
	o.AuthenticationOauthJwtTokenLifetime = &v
}

// GetAuthenticationOauthJwtTokenRetrievedTime returns the AuthenticationOauthJwtTokenRetrievedTime field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenRetrievedTime() int32 {
	if o == nil || o.AuthenticationOauthJwtTokenRetrievedTime == nil {
		var ret int32
		return ret
	}
	return *o.AuthenticationOauthJwtTokenRetrievedTime
}

// GetAuthenticationOauthJwtTokenRetrievedTimeOk returns a tuple with the AuthenticationOauthJwtTokenRetrievedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenRetrievedTimeOk() (*int32, bool) {
	if o == nil || o.AuthenticationOauthJwtTokenRetrievedTime == nil {
		return nil, false
	}
	return o.AuthenticationOauthJwtTokenRetrievedTime, true
}

// HasAuthenticationOauthJwtTokenRetrievedTime returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtTokenRetrievedTime() bool {
	if o != nil && o.AuthenticationOauthJwtTokenRetrievedTime != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthJwtTokenRetrievedTime gets a reference to the given int32 and assigns it to the AuthenticationOauthJwtTokenRetrievedTime field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtTokenRetrievedTime(v int32) {
	o.AuthenticationOauthJwtTokenRetrievedTime = &v
}

// GetAuthenticationOauthJwtTokenState returns the AuthenticationOauthJwtTokenState field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenState() string {
	if o == nil || o.AuthenticationOauthJwtTokenState == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationOauthJwtTokenState
}

// GetAuthenticationOauthJwtTokenStateOk returns a tuple with the AuthenticationOauthJwtTokenState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenStateOk() (*string, bool) {
	if o == nil || o.AuthenticationOauthJwtTokenState == nil {
		return nil, false
	}
	return o.AuthenticationOauthJwtTokenState, true
}

// HasAuthenticationOauthJwtTokenState returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtTokenState() bool {
	if o != nil && o.AuthenticationOauthJwtTokenState != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthJwtTokenState gets a reference to the given string and assigns it to the AuthenticationOauthJwtTokenState field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtTokenState(v string) {
	o.AuthenticationOauthJwtTokenState = &v
}

// GetAuthenticationScheme returns the AuthenticationScheme field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationScheme() string {
	if o == nil || o.AuthenticationScheme == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationScheme
}

// GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationSchemeOk() (*string, bool) {
	if o == nil || o.AuthenticationScheme == nil {
		return nil, false
	}
	return o.AuthenticationScheme, true
}

// HasAuthenticationScheme returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationScheme() bool {
	if o != nil && o.AuthenticationScheme != nil {
		return true
	}

	return false
}

// SetAuthenticationScheme gets a reference to the given string and assigns it to the AuthenticationScheme field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationScheme(v string) {
	o.AuthenticationScheme = &v
}

// GetCounter returns the Counter field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetCounter() MsgVpnRestDeliveryPointRestConsumerCounter {
	if o == nil || o.Counter == nil {
		var ret MsgVpnRestDeliveryPointRestConsumerCounter
		return ret
	}
	return *o.Counter
}

// GetCounterOk returns a tuple with the Counter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetCounterOk() (*MsgVpnRestDeliveryPointRestConsumerCounter, bool) {
	if o == nil || o.Counter == nil {
		return nil, false
	}
	return o.Counter, true
}

// HasCounter returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasCounter() bool {
	if o != nil && o.Counter != nil {
		return true
	}

	return false
}

// SetCounter gets a reference to the given MsgVpnRestDeliveryPointRestConsumerCounter and assigns it to the Counter field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetCounter(v MsgVpnRestDeliveryPointRestConsumerCounter) {
	o.Counter = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpMethod() string {
	if o == nil || o.HttpMethod == nil {
		var ret string
		return ret
	}
	return *o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpMethodOk() (*string, bool) {
	if o == nil || o.HttpMethod == nil {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpMethod() bool {
	if o != nil && o.HttpMethod != nil {
		return true
	}

	return false
}

// SetHttpMethod gets a reference to the given string and assigns it to the HttpMethod field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpMethod(v string) {
	o.HttpMethod = &v
}

// GetHttpRequestConnectionCloseTxMsgCount returns the HttpRequestConnectionCloseTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestConnectionCloseTxMsgCount() int64 {
	if o == nil || o.HttpRequestConnectionCloseTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.HttpRequestConnectionCloseTxMsgCount
}

// GetHttpRequestConnectionCloseTxMsgCountOk returns a tuple with the HttpRequestConnectionCloseTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestConnectionCloseTxMsgCountOk() (*int64, bool) {
	if o == nil || o.HttpRequestConnectionCloseTxMsgCount == nil {
		return nil, false
	}
	return o.HttpRequestConnectionCloseTxMsgCount, true
}

// HasHttpRequestConnectionCloseTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpRequestConnectionCloseTxMsgCount() bool {
	if o != nil && o.HttpRequestConnectionCloseTxMsgCount != nil {
		return true
	}

	return false
}

// SetHttpRequestConnectionCloseTxMsgCount gets a reference to the given int64 and assigns it to the HttpRequestConnectionCloseTxMsgCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpRequestConnectionCloseTxMsgCount(v int64) {
	o.HttpRequestConnectionCloseTxMsgCount = &v
}

// GetHttpRequestOutstandingTxMsgCount returns the HttpRequestOutstandingTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestOutstandingTxMsgCount() int64 {
	if o == nil || o.HttpRequestOutstandingTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.HttpRequestOutstandingTxMsgCount
}

// GetHttpRequestOutstandingTxMsgCountOk returns a tuple with the HttpRequestOutstandingTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestOutstandingTxMsgCountOk() (*int64, bool) {
	if o == nil || o.HttpRequestOutstandingTxMsgCount == nil {
		return nil, false
	}
	return o.HttpRequestOutstandingTxMsgCount, true
}

// HasHttpRequestOutstandingTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpRequestOutstandingTxMsgCount() bool {
	if o != nil && o.HttpRequestOutstandingTxMsgCount != nil {
		return true
	}

	return false
}

// SetHttpRequestOutstandingTxMsgCount gets a reference to the given int64 and assigns it to the HttpRequestOutstandingTxMsgCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpRequestOutstandingTxMsgCount(v int64) {
	o.HttpRequestOutstandingTxMsgCount = &v
}

// GetHttpRequestTimedOutTxMsgCount returns the HttpRequestTimedOutTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTimedOutTxMsgCount() int64 {
	if o == nil || o.HttpRequestTimedOutTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.HttpRequestTimedOutTxMsgCount
}

// GetHttpRequestTimedOutTxMsgCountOk returns a tuple with the HttpRequestTimedOutTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTimedOutTxMsgCountOk() (*int64, bool) {
	if o == nil || o.HttpRequestTimedOutTxMsgCount == nil {
		return nil, false
	}
	return o.HttpRequestTimedOutTxMsgCount, true
}

// HasHttpRequestTimedOutTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpRequestTimedOutTxMsgCount() bool {
	if o != nil && o.HttpRequestTimedOutTxMsgCount != nil {
		return true
	}

	return false
}

// SetHttpRequestTimedOutTxMsgCount gets a reference to the given int64 and assigns it to the HttpRequestTimedOutTxMsgCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpRequestTimedOutTxMsgCount(v int64) {
	o.HttpRequestTimedOutTxMsgCount = &v
}

// GetHttpRequestTxByteCount returns the HttpRequestTxByteCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTxByteCount() int64 {
	if o == nil || o.HttpRequestTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.HttpRequestTxByteCount
}

// GetHttpRequestTxByteCountOk returns a tuple with the HttpRequestTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTxByteCountOk() (*int64, bool) {
	if o == nil || o.HttpRequestTxByteCount == nil {
		return nil, false
	}
	return o.HttpRequestTxByteCount, true
}

// HasHttpRequestTxByteCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpRequestTxByteCount() bool {
	if o != nil && o.HttpRequestTxByteCount != nil {
		return true
	}

	return false
}

// SetHttpRequestTxByteCount gets a reference to the given int64 and assigns it to the HttpRequestTxByteCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpRequestTxByteCount(v int64) {
	o.HttpRequestTxByteCount = &v
}

// GetHttpRequestTxMsgCount returns the HttpRequestTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTxMsgCount() int64 {
	if o == nil || o.HttpRequestTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.HttpRequestTxMsgCount
}

// GetHttpRequestTxMsgCountOk returns a tuple with the HttpRequestTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTxMsgCountOk() (*int64, bool) {
	if o == nil || o.HttpRequestTxMsgCount == nil {
		return nil, false
	}
	return o.HttpRequestTxMsgCount, true
}

// HasHttpRequestTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpRequestTxMsgCount() bool {
	if o != nil && o.HttpRequestTxMsgCount != nil {
		return true
	}

	return false
}

// SetHttpRequestTxMsgCount gets a reference to the given int64 and assigns it to the HttpRequestTxMsgCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpRequestTxMsgCount(v int64) {
	o.HttpRequestTxMsgCount = &v
}

// GetHttpResponseErrorRxMsgCount returns the HttpResponseErrorRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseErrorRxMsgCount() int64 {
	if o == nil || o.HttpResponseErrorRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.HttpResponseErrorRxMsgCount
}

// GetHttpResponseErrorRxMsgCountOk returns a tuple with the HttpResponseErrorRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseErrorRxMsgCountOk() (*int64, bool) {
	if o == nil || o.HttpResponseErrorRxMsgCount == nil {
		return nil, false
	}
	return o.HttpResponseErrorRxMsgCount, true
}

// HasHttpResponseErrorRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpResponseErrorRxMsgCount() bool {
	if o != nil && o.HttpResponseErrorRxMsgCount != nil {
		return true
	}

	return false
}

// SetHttpResponseErrorRxMsgCount gets a reference to the given int64 and assigns it to the HttpResponseErrorRxMsgCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpResponseErrorRxMsgCount(v int64) {
	o.HttpResponseErrorRxMsgCount = &v
}

// GetHttpResponseRxByteCount returns the HttpResponseRxByteCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseRxByteCount() int64 {
	if o == nil || o.HttpResponseRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.HttpResponseRxByteCount
}

// GetHttpResponseRxByteCountOk returns a tuple with the HttpResponseRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseRxByteCountOk() (*int64, bool) {
	if o == nil || o.HttpResponseRxByteCount == nil {
		return nil, false
	}
	return o.HttpResponseRxByteCount, true
}

// HasHttpResponseRxByteCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpResponseRxByteCount() bool {
	if o != nil && o.HttpResponseRxByteCount != nil {
		return true
	}

	return false
}

// SetHttpResponseRxByteCount gets a reference to the given int64 and assigns it to the HttpResponseRxByteCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpResponseRxByteCount(v int64) {
	o.HttpResponseRxByteCount = &v
}

// GetHttpResponseRxMsgCount returns the HttpResponseRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseRxMsgCount() int64 {
	if o == nil || o.HttpResponseRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.HttpResponseRxMsgCount
}

// GetHttpResponseRxMsgCountOk returns a tuple with the HttpResponseRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseRxMsgCountOk() (*int64, bool) {
	if o == nil || o.HttpResponseRxMsgCount == nil {
		return nil, false
	}
	return o.HttpResponseRxMsgCount, true
}

// HasHttpResponseRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpResponseRxMsgCount() bool {
	if o != nil && o.HttpResponseRxMsgCount != nil {
		return true
	}

	return false
}

// SetHttpResponseRxMsgCount gets a reference to the given int64 and assigns it to the HttpResponseRxMsgCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpResponseRxMsgCount(v int64) {
	o.HttpResponseRxMsgCount = &v
}

// GetHttpResponseSuccessRxMsgCount returns the HttpResponseSuccessRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseSuccessRxMsgCount() int64 {
	if o == nil || o.HttpResponseSuccessRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.HttpResponseSuccessRxMsgCount
}

// GetHttpResponseSuccessRxMsgCountOk returns a tuple with the HttpResponseSuccessRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseSuccessRxMsgCountOk() (*int64, bool) {
	if o == nil || o.HttpResponseSuccessRxMsgCount == nil {
		return nil, false
	}
	return o.HttpResponseSuccessRxMsgCount, true
}

// HasHttpResponseSuccessRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpResponseSuccessRxMsgCount() bool {
	if o != nil && o.HttpResponseSuccessRxMsgCount != nil {
		return true
	}

	return false
}

// SetHttpResponseSuccessRxMsgCount gets a reference to the given int64 and assigns it to the HttpResponseSuccessRxMsgCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpResponseSuccessRxMsgCount(v int64) {
	o.HttpResponseSuccessRxMsgCount = &v
}

// GetLastConnectionFailureLocalEndpoint returns the LastConnectionFailureLocalEndpoint field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureLocalEndpoint() string {
	if o == nil || o.LastConnectionFailureLocalEndpoint == nil {
		var ret string
		return ret
	}
	return *o.LastConnectionFailureLocalEndpoint
}

// GetLastConnectionFailureLocalEndpointOk returns a tuple with the LastConnectionFailureLocalEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureLocalEndpointOk() (*string, bool) {
	if o == nil || o.LastConnectionFailureLocalEndpoint == nil {
		return nil, false
	}
	return o.LastConnectionFailureLocalEndpoint, true
}

// HasLastConnectionFailureLocalEndpoint returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastConnectionFailureLocalEndpoint() bool {
	if o != nil && o.LastConnectionFailureLocalEndpoint != nil {
		return true
	}

	return false
}

// SetLastConnectionFailureLocalEndpoint gets a reference to the given string and assigns it to the LastConnectionFailureLocalEndpoint field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastConnectionFailureLocalEndpoint(v string) {
	o.LastConnectionFailureLocalEndpoint = &v
}

// GetLastConnectionFailureReason returns the LastConnectionFailureReason field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureReason() string {
	if o == nil || o.LastConnectionFailureReason == nil {
		var ret string
		return ret
	}
	return *o.LastConnectionFailureReason
}

// GetLastConnectionFailureReasonOk returns a tuple with the LastConnectionFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureReasonOk() (*string, bool) {
	if o == nil || o.LastConnectionFailureReason == nil {
		return nil, false
	}
	return o.LastConnectionFailureReason, true
}

// HasLastConnectionFailureReason returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastConnectionFailureReason() bool {
	if o != nil && o.LastConnectionFailureReason != nil {
		return true
	}

	return false
}

// SetLastConnectionFailureReason gets a reference to the given string and assigns it to the LastConnectionFailureReason field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastConnectionFailureReason(v string) {
	o.LastConnectionFailureReason = &v
}

// GetLastConnectionFailureRemoteEndpoint returns the LastConnectionFailureRemoteEndpoint field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureRemoteEndpoint() string {
	if o == nil || o.LastConnectionFailureRemoteEndpoint == nil {
		var ret string
		return ret
	}
	return *o.LastConnectionFailureRemoteEndpoint
}

// GetLastConnectionFailureRemoteEndpointOk returns a tuple with the LastConnectionFailureRemoteEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureRemoteEndpointOk() (*string, bool) {
	if o == nil || o.LastConnectionFailureRemoteEndpoint == nil {
		return nil, false
	}
	return o.LastConnectionFailureRemoteEndpoint, true
}

// HasLastConnectionFailureRemoteEndpoint returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastConnectionFailureRemoteEndpoint() bool {
	if o != nil && o.LastConnectionFailureRemoteEndpoint != nil {
		return true
	}

	return false
}

// SetLastConnectionFailureRemoteEndpoint gets a reference to the given string and assigns it to the LastConnectionFailureRemoteEndpoint field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastConnectionFailureRemoteEndpoint(v string) {
	o.LastConnectionFailureRemoteEndpoint = &v
}

// GetLastConnectionFailureTime returns the LastConnectionFailureTime field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureTime() int32 {
	if o == nil || o.LastConnectionFailureTime == nil {
		var ret int32
		return ret
	}
	return *o.LastConnectionFailureTime
}

// GetLastConnectionFailureTimeOk returns a tuple with the LastConnectionFailureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureTimeOk() (*int32, bool) {
	if o == nil || o.LastConnectionFailureTime == nil {
		return nil, false
	}
	return o.LastConnectionFailureTime, true
}

// HasLastConnectionFailureTime returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastConnectionFailureTime() bool {
	if o != nil && o.LastConnectionFailureTime != nil {
		return true
	}

	return false
}

// SetLastConnectionFailureTime gets a reference to the given int32 and assigns it to the LastConnectionFailureTime field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastConnectionFailureTime(v int32) {
	o.LastConnectionFailureTime = &v
}

// GetLastFailureReason returns the LastFailureReason field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastFailureReason() string {
	if o == nil || o.LastFailureReason == nil {
		var ret string
		return ret
	}
	return *o.LastFailureReason
}

// GetLastFailureReasonOk returns a tuple with the LastFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastFailureReasonOk() (*string, bool) {
	if o == nil || o.LastFailureReason == nil {
		return nil, false
	}
	return o.LastFailureReason, true
}

// HasLastFailureReason returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastFailureReason() bool {
	if o != nil && o.LastFailureReason != nil {
		return true
	}

	return false
}

// SetLastFailureReason gets a reference to the given string and assigns it to the LastFailureReason field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastFailureReason(v string) {
	o.LastFailureReason = &v
}

// GetLastFailureTime returns the LastFailureTime field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastFailureTime() int32 {
	if o == nil || o.LastFailureTime == nil {
		var ret int32
		return ret
	}
	return *o.LastFailureTime
}

// GetLastFailureTimeOk returns a tuple with the LastFailureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastFailureTimeOk() (*int32, bool) {
	if o == nil || o.LastFailureTime == nil {
		return nil, false
	}
	return o.LastFailureTime, true
}

// HasLastFailureTime returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastFailureTime() bool {
	if o != nil && o.LastFailureTime != nil {
		return true
	}

	return false
}

// SetLastFailureTime gets a reference to the given int32 and assigns it to the LastFailureTime field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastFailureTime(v int32) {
	o.LastFailureTime = &v
}

// GetLocalInterface returns the LocalInterface field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLocalInterface() string {
	if o == nil || o.LocalInterface == nil {
		var ret string
		return ret
	}
	return *o.LocalInterface
}

// GetLocalInterfaceOk returns a tuple with the LocalInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetLocalInterfaceOk() (*string, bool) {
	if o == nil || o.LocalInterface == nil {
		return nil, false
	}
	return o.LocalInterface, true
}

// HasLocalInterface returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasLocalInterface() bool {
	if o != nil && o.LocalInterface != nil {
		return true
	}

	return false
}

// SetLocalInterface gets a reference to the given string and assigns it to the LocalInterface field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetLocalInterface(v string) {
	o.LocalInterface = &v
}

// GetMaxPostWaitTime returns the MaxPostWaitTime field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetMaxPostWaitTime() int32 {
	if o == nil || o.MaxPostWaitTime == nil {
		var ret int32
		return ret
	}
	return *o.MaxPostWaitTime
}

// GetMaxPostWaitTimeOk returns a tuple with the MaxPostWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetMaxPostWaitTimeOk() (*int32, bool) {
	if o == nil || o.MaxPostWaitTime == nil {
		return nil, false
	}
	return o.MaxPostWaitTime, true
}

// HasMaxPostWaitTime returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasMaxPostWaitTime() bool {
	if o != nil && o.MaxPostWaitTime != nil {
		return true
	}

	return false
}

// SetMaxPostWaitTime gets a reference to the given int32 and assigns it to the MaxPostWaitTime field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetMaxPostWaitTime(v int32) {
	o.MaxPostWaitTime = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetOutgoingConnectionCount returns the OutgoingConnectionCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetOutgoingConnectionCount() int32 {
	if o == nil || o.OutgoingConnectionCount == nil {
		var ret int32
		return ret
	}
	return *o.OutgoingConnectionCount
}

// GetOutgoingConnectionCountOk returns a tuple with the OutgoingConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetOutgoingConnectionCountOk() (*int32, bool) {
	if o == nil || o.OutgoingConnectionCount == nil {
		return nil, false
	}
	return o.OutgoingConnectionCount, true
}

// HasOutgoingConnectionCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasOutgoingConnectionCount() bool {
	if o != nil && o.OutgoingConnectionCount != nil {
		return true
	}

	return false
}

// SetOutgoingConnectionCount gets a reference to the given int32 and assigns it to the OutgoingConnectionCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetOutgoingConnectionCount(v int32) {
	o.OutgoingConnectionCount = &v
}

// GetRemoteHost returns the RemoteHost field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemoteHost() string {
	if o == nil || o.RemoteHost == nil {
		var ret string
		return ret
	}
	return *o.RemoteHost
}

// GetRemoteHostOk returns a tuple with the RemoteHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemoteHostOk() (*string, bool) {
	if o == nil || o.RemoteHost == nil {
		return nil, false
	}
	return o.RemoteHost, true
}

// HasRemoteHost returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasRemoteHost() bool {
	if o != nil && o.RemoteHost != nil {
		return true
	}

	return false
}

// SetRemoteHost gets a reference to the given string and assigns it to the RemoteHost field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetRemoteHost(v string) {
	o.RemoteHost = &v
}

// GetRemoteOutgoingConnectionUpCount returns the RemoteOutgoingConnectionUpCount field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemoteOutgoingConnectionUpCount() int64 {
	if o == nil || o.RemoteOutgoingConnectionUpCount == nil {
		var ret int64
		return ret
	}
	return *o.RemoteOutgoingConnectionUpCount
}

// GetRemoteOutgoingConnectionUpCountOk returns a tuple with the RemoteOutgoingConnectionUpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemoteOutgoingConnectionUpCountOk() (*int64, bool) {
	if o == nil || o.RemoteOutgoingConnectionUpCount == nil {
		return nil, false
	}
	return o.RemoteOutgoingConnectionUpCount, true
}

// HasRemoteOutgoingConnectionUpCount returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasRemoteOutgoingConnectionUpCount() bool {
	if o != nil && o.RemoteOutgoingConnectionUpCount != nil {
		return true
	}

	return false
}

// SetRemoteOutgoingConnectionUpCount gets a reference to the given int64 and assigns it to the RemoteOutgoingConnectionUpCount field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetRemoteOutgoingConnectionUpCount(v int64) {
	o.RemoteOutgoingConnectionUpCount = &v
}

// GetRemotePort returns the RemotePort field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemotePort() int64 {
	if o == nil || o.RemotePort == nil {
		var ret int64
		return ret
	}
	return *o.RemotePort
}

// GetRemotePortOk returns a tuple with the RemotePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemotePortOk() (*int64, bool) {
	if o == nil || o.RemotePort == nil {
		return nil, false
	}
	return o.RemotePort, true
}

// HasRemotePort returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasRemotePort() bool {
	if o != nil && o.RemotePort != nil {
		return true
	}

	return false
}

// SetRemotePort gets a reference to the given int64 and assigns it to the RemotePort field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetRemotePort(v int64) {
	o.RemotePort = &v
}

// GetRestConsumerName returns the RestConsumerName field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRestConsumerName() string {
	if o == nil || o.RestConsumerName == nil {
		var ret string
		return ret
	}
	return *o.RestConsumerName
}

// GetRestConsumerNameOk returns a tuple with the RestConsumerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRestConsumerNameOk() (*string, bool) {
	if o == nil || o.RestConsumerName == nil {
		return nil, false
	}
	return o.RestConsumerName, true
}

// HasRestConsumerName returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasRestConsumerName() bool {
	if o != nil && o.RestConsumerName != nil {
		return true
	}

	return false
}

// SetRestConsumerName gets a reference to the given string and assigns it to the RestConsumerName field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetRestConsumerName(v string) {
	o.RestConsumerName = &v
}

// GetRestDeliveryPointName returns the RestDeliveryPointName field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRestDeliveryPointName() string {
	if o == nil || o.RestDeliveryPointName == nil {
		var ret string
		return ret
	}
	return *o.RestDeliveryPointName
}

// GetRestDeliveryPointNameOk returns a tuple with the RestDeliveryPointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRestDeliveryPointNameOk() (*string, bool) {
	if o == nil || o.RestDeliveryPointName == nil {
		return nil, false
	}
	return o.RestDeliveryPointName, true
}

// HasRestDeliveryPointName returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasRestDeliveryPointName() bool {
	if o != nil && o.RestDeliveryPointName != nil {
		return true
	}

	return false
}

// SetRestDeliveryPointName gets a reference to the given string and assigns it to the RestDeliveryPointName field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetRestDeliveryPointName(v string) {
	o.RestDeliveryPointName = &v
}

// GetRetryDelay returns the RetryDelay field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRetryDelay() int32 {
	if o == nil || o.RetryDelay == nil {
		var ret int32
		return ret
	}
	return *o.RetryDelay
}

// GetRetryDelayOk returns a tuple with the RetryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetRetryDelayOk() (*int32, bool) {
	if o == nil || o.RetryDelay == nil {
		return nil, false
	}
	return o.RetryDelay, true
}

// HasRetryDelay returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasRetryDelay() bool {
	if o != nil && o.RetryDelay != nil {
		return true
	}

	return false
}

// SetRetryDelay gets a reference to the given int32 and assigns it to the RetryDelay field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetRetryDelay(v int32) {
	o.RetryDelay = &v
}

// GetTlsCipherSuiteList returns the TlsCipherSuiteList field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetTlsCipherSuiteList() string {
	if o == nil || o.TlsCipherSuiteList == nil {
		var ret string
		return ret
	}
	return *o.TlsCipherSuiteList
}

// GetTlsCipherSuiteListOk returns a tuple with the TlsCipherSuiteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetTlsCipherSuiteListOk() (*string, bool) {
	if o == nil || o.TlsCipherSuiteList == nil {
		return nil, false
	}
	return o.TlsCipherSuiteList, true
}

// HasTlsCipherSuiteList returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasTlsCipherSuiteList() bool {
	if o != nil && o.TlsCipherSuiteList != nil {
		return true
	}

	return false
}

// SetTlsCipherSuiteList gets a reference to the given string and assigns it to the TlsCipherSuiteList field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetTlsCipherSuiteList(v string) {
	o.TlsCipherSuiteList = &v
}

// GetTlsEnabled returns the TlsEnabled field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetTlsEnabled() bool {
	if o == nil || o.TlsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TlsEnabled
}

// GetTlsEnabledOk returns a tuple with the TlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetTlsEnabledOk() (*bool, bool) {
	if o == nil || o.TlsEnabled == nil {
		return nil, false
	}
	return o.TlsEnabled, true
}

// HasTlsEnabled returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasTlsEnabled() bool {
	if o != nil && o.TlsEnabled != nil {
		return true
	}

	return false
}

// SetTlsEnabled gets a reference to the given bool and assigns it to the TlsEnabled field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetTlsEnabled(v bool) {
	o.TlsEnabled = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetUp() bool {
	if o == nil || o.Up == nil {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) GetUpOk() (*bool, bool) {
	if o == nil || o.Up == nil {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *MsgVpnRestDeliveryPointRestConsumer) HasUp() bool {
	if o != nil && o.Up != nil {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *MsgVpnRestDeliveryPointRestConsumer) SetUp(v bool) {
	o.Up = &v
}

func (o MsgVpnRestDeliveryPointRestConsumer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationHttpBasicUsername != nil {
		toSerialize["authenticationHttpBasicUsername"] = o.AuthenticationHttpBasicUsername
	}
	if o.AuthenticationHttpHeaderName != nil {
		toSerialize["authenticationHttpHeaderName"] = o.AuthenticationHttpHeaderName
	}
	if o.AuthenticationOauthClientId != nil {
		toSerialize["authenticationOauthClientId"] = o.AuthenticationOauthClientId
	}
	if o.AuthenticationOauthClientLastFailureReason != nil {
		toSerialize["authenticationOauthClientLastFailureReason"] = o.AuthenticationOauthClientLastFailureReason
	}
	if o.AuthenticationOauthClientLastFailureTime != nil {
		toSerialize["authenticationOauthClientLastFailureTime"] = o.AuthenticationOauthClientLastFailureTime
	}
	if o.AuthenticationOauthClientScope != nil {
		toSerialize["authenticationOauthClientScope"] = o.AuthenticationOauthClientScope
	}
	if o.AuthenticationOauthClientTokenEndpoint != nil {
		toSerialize["authenticationOauthClientTokenEndpoint"] = o.AuthenticationOauthClientTokenEndpoint
	}
	if o.AuthenticationOauthClientTokenLifetime != nil {
		toSerialize["authenticationOauthClientTokenLifetime"] = o.AuthenticationOauthClientTokenLifetime
	}
	if o.AuthenticationOauthClientTokenRetrievedTime != nil {
		toSerialize["authenticationOauthClientTokenRetrievedTime"] = o.AuthenticationOauthClientTokenRetrievedTime
	}
	if o.AuthenticationOauthClientTokenState != nil {
		toSerialize["authenticationOauthClientTokenState"] = o.AuthenticationOauthClientTokenState
	}
	if o.AuthenticationOauthJwtLastFailureReason != nil {
		toSerialize["authenticationOauthJwtLastFailureReason"] = o.AuthenticationOauthJwtLastFailureReason
	}
	if o.AuthenticationOauthJwtLastFailureTime != nil {
		toSerialize["authenticationOauthJwtLastFailureTime"] = o.AuthenticationOauthJwtLastFailureTime
	}
	if o.AuthenticationOauthJwtTokenEndpoint != nil {
		toSerialize["authenticationOauthJwtTokenEndpoint"] = o.AuthenticationOauthJwtTokenEndpoint
	}
	if o.AuthenticationOauthJwtTokenLifetime != nil {
		toSerialize["authenticationOauthJwtTokenLifetime"] = o.AuthenticationOauthJwtTokenLifetime
	}
	if o.AuthenticationOauthJwtTokenRetrievedTime != nil {
		toSerialize["authenticationOauthJwtTokenRetrievedTime"] = o.AuthenticationOauthJwtTokenRetrievedTime
	}
	if o.AuthenticationOauthJwtTokenState != nil {
		toSerialize["authenticationOauthJwtTokenState"] = o.AuthenticationOauthJwtTokenState
	}
	if o.AuthenticationScheme != nil {
		toSerialize["authenticationScheme"] = o.AuthenticationScheme
	}
	if o.Counter != nil {
		toSerialize["counter"] = o.Counter
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.HttpMethod != nil {
		toSerialize["httpMethod"] = o.HttpMethod
	}
	if o.HttpRequestConnectionCloseTxMsgCount != nil {
		toSerialize["httpRequestConnectionCloseTxMsgCount"] = o.HttpRequestConnectionCloseTxMsgCount
	}
	if o.HttpRequestOutstandingTxMsgCount != nil {
		toSerialize["httpRequestOutstandingTxMsgCount"] = o.HttpRequestOutstandingTxMsgCount
	}
	if o.HttpRequestTimedOutTxMsgCount != nil {
		toSerialize["httpRequestTimedOutTxMsgCount"] = o.HttpRequestTimedOutTxMsgCount
	}
	if o.HttpRequestTxByteCount != nil {
		toSerialize["httpRequestTxByteCount"] = o.HttpRequestTxByteCount
	}
	if o.HttpRequestTxMsgCount != nil {
		toSerialize["httpRequestTxMsgCount"] = o.HttpRequestTxMsgCount
	}
	if o.HttpResponseErrorRxMsgCount != nil {
		toSerialize["httpResponseErrorRxMsgCount"] = o.HttpResponseErrorRxMsgCount
	}
	if o.HttpResponseRxByteCount != nil {
		toSerialize["httpResponseRxByteCount"] = o.HttpResponseRxByteCount
	}
	if o.HttpResponseRxMsgCount != nil {
		toSerialize["httpResponseRxMsgCount"] = o.HttpResponseRxMsgCount
	}
	if o.HttpResponseSuccessRxMsgCount != nil {
		toSerialize["httpResponseSuccessRxMsgCount"] = o.HttpResponseSuccessRxMsgCount
	}
	if o.LastConnectionFailureLocalEndpoint != nil {
		toSerialize["lastConnectionFailureLocalEndpoint"] = o.LastConnectionFailureLocalEndpoint
	}
	if o.LastConnectionFailureReason != nil {
		toSerialize["lastConnectionFailureReason"] = o.LastConnectionFailureReason
	}
	if o.LastConnectionFailureRemoteEndpoint != nil {
		toSerialize["lastConnectionFailureRemoteEndpoint"] = o.LastConnectionFailureRemoteEndpoint
	}
	if o.LastConnectionFailureTime != nil {
		toSerialize["lastConnectionFailureTime"] = o.LastConnectionFailureTime
	}
	if o.LastFailureReason != nil {
		toSerialize["lastFailureReason"] = o.LastFailureReason
	}
	if o.LastFailureTime != nil {
		toSerialize["lastFailureTime"] = o.LastFailureTime
	}
	if o.LocalInterface != nil {
		toSerialize["localInterface"] = o.LocalInterface
	}
	if o.MaxPostWaitTime != nil {
		toSerialize["maxPostWaitTime"] = o.MaxPostWaitTime
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.OutgoingConnectionCount != nil {
		toSerialize["outgoingConnectionCount"] = o.OutgoingConnectionCount
	}
	if o.RemoteHost != nil {
		toSerialize["remoteHost"] = o.RemoteHost
	}
	if o.RemoteOutgoingConnectionUpCount != nil {
		toSerialize["remoteOutgoingConnectionUpCount"] = o.RemoteOutgoingConnectionUpCount
	}
	if o.RemotePort != nil {
		toSerialize["remotePort"] = o.RemotePort
	}
	if o.RestConsumerName != nil {
		toSerialize["restConsumerName"] = o.RestConsumerName
	}
	if o.RestDeliveryPointName != nil {
		toSerialize["restDeliveryPointName"] = o.RestDeliveryPointName
	}
	if o.RetryDelay != nil {
		toSerialize["retryDelay"] = o.RetryDelay
	}
	if o.TlsCipherSuiteList != nil {
		toSerialize["tlsCipherSuiteList"] = o.TlsCipherSuiteList
	}
	if o.TlsEnabled != nil {
		toSerialize["tlsEnabled"] = o.TlsEnabled
	}
	if o.Up != nil {
		toSerialize["up"] = o.Up
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnRestDeliveryPointRestConsumer struct {
	value *MsgVpnRestDeliveryPointRestConsumer
	isSet bool
}

func (v NullableMsgVpnRestDeliveryPointRestConsumer) Get() *MsgVpnRestDeliveryPointRestConsumer {
	return v.value
}

func (v *NullableMsgVpnRestDeliveryPointRestConsumer) Set(val *MsgVpnRestDeliveryPointRestConsumer) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnRestDeliveryPointRestConsumer) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnRestDeliveryPointRestConsumer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnRestDeliveryPointRestConsumer(val *MsgVpnRestDeliveryPointRestConsumer) *NullableMsgVpnRestDeliveryPointRestConsumer {
	return &NullableMsgVpnRestDeliveryPointRestConsumer{value: val, isSet: true}
}

func (v NullableMsgVpnRestDeliveryPointRestConsumer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnRestDeliveryPointRestConsumer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
