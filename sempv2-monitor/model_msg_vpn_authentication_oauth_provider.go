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

// MsgVpnAuthenticationOauthProvider struct for MsgVpnAuthenticationOauthProvider
type MsgVpnAuthenticationOauthProvider struct {
	// The audience claim name, indicating which part of the object to use for determining the audience.
	AudienceClaimName *string `json:"audienceClaimName,omitempty"`
	// The audience claim source, indicating where to search for the audience value. The allowed values and their meaning are:  <pre> \"access-token\" - The OAuth v2 access_token. \"id-token\" - The OpenID Connect id_token. \"introspection\" - The result of introspecting the OAuth v2 access_token. </pre>
	AudienceClaimSource *string `json:"audienceClaimSource,omitempty"`
	// The required audience value for a token to be considered valid.
	AudienceClaimValue *string `json:"audienceClaimValue,omitempty"`
	// Indicates whether audience validation is enabled.
	AudienceValidationEnabled *bool `json:"audienceValidationEnabled,omitempty"`
	// The number of OAuth Provider client authentications that succeeded.
	AuthenticationSuccessCount *int64 `json:"authenticationSuccessCount,omitempty"`
	// The authorization group claim name, indicating which part of the object to use for determining the authorization group.
	AuthorizationGroupClaimName *string `json:"authorizationGroupClaimName,omitempty"`
	// The authorization group claim source, indicating where to search for the authorization group name. The allowed values and their meaning are:  <pre> \"access-token\" - The OAuth v2 access_token. \"id-token\" - The OpenID Connect id_token. \"introspection\" - The result of introspecting the OAuth v2 access_token. </pre>
	AuthorizationGroupClaimSource *string `json:"authorizationGroupClaimSource,omitempty"`
	// Indicates whether OAuth based authorization is enabled and the configured authorization type for OAuth clients is overridden.
	AuthorizationGroupEnabled *bool `json:"authorizationGroupEnabled,omitempty"`
	// Indicates whether clients are disconnected when their tokens expire.
	DisconnectOnTokenExpirationEnabled *bool `json:"disconnectOnTokenExpirationEnabled,omitempty"`
	// Indicates whether OAuth Provider client authentication is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The reason for the last JWKS public key refresh failure.
	JwksLastRefreshFailureReason *string `json:"jwksLastRefreshFailureReason,omitempty"`
	// The timestamp of the last JWKS public key refresh failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	JwksLastRefreshFailureTime *int32 `json:"jwksLastRefreshFailureTime,omitempty"`
	// The timestamp of the last JWKS public key refresh success. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	JwksLastRefreshTime *int32 `json:"jwksLastRefreshTime,omitempty"`
	// The timestamp of the next scheduled JWKS public key refresh. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	JwksNextScheduledRefreshTime *int32 `json:"jwksNextScheduledRefreshTime,omitempty"`
	// The number of JWKS public key refresh failures.
	JwksRefreshFailureCount *int64 `json:"jwksRefreshFailureCount,omitempty"`
	// The number of seconds between forced JWKS public key refreshing.
	JwksRefreshInterval *int32 `json:"jwksRefreshInterval,omitempty"`
	// The URI where the OAuth provider publishes its JWKS public keys.
	JwksUri *string `json:"jwksUri,omitempty"`
	// The number of login failures due to an incorrect audience value.
	LoginFailureIncorrectAudienceValueCount *int64 `json:"loginFailureIncorrectAudienceValueCount,omitempty"`
	// The number of login failures due to an invalid audience value.
	LoginFailureInvalidAudienceValueCount *int64 `json:"loginFailureInvalidAudienceValueCount,omitempty"`
	// The number of login failures due to an invalid authorization group value (zero-length or non-string).
	LoginFailureInvalidAuthorizationGroupValueCount *int64 `json:"loginFailureInvalidAuthorizationGroupValueCount,omitempty"`
	// The number of login failures due to an invalid JWT signature.
	LoginFailureInvalidJwtSignatureCount *int64 `json:"loginFailureInvalidJwtSignatureCount,omitempty"`
	// The number of login failures due to an invalid username value.
	LoginFailureInvalidUsernameValueCount *int64 `json:"loginFailureInvalidUsernameValueCount,omitempty"`
	// The number of login failures due to a mismatched username.
	LoginFailureMismatchedUsernameCount *int64 `json:"loginFailureMismatchedUsernameCount,omitempty"`
	// The number of login failures due to a missing audience claim.
	LoginFailureMissingAudienceCount *int64 `json:"loginFailureMissingAudienceCount,omitempty"`
	// The number of login failures due to a missing JSON Web Key (JWK).
	LoginFailureMissingJwkCount *int64 `json:"loginFailureMissingJwkCount,omitempty"`
	// The number of login failures due to a missing or invalid token.
	LoginFailureMissingOrInvalidTokenCount *int64 `json:"loginFailureMissingOrInvalidTokenCount,omitempty"`
	// The number of login failures due to a missing username.
	LoginFailureMissingUsernameCount *int64 `json:"loginFailureMissingUsernameCount,omitempty"`
	// The number of login failures due to a token being expired.
	LoginFailureTokenExpiredCount *int64 `json:"loginFailureTokenExpiredCount,omitempty"`
	// The number of login failures due to a token introspection error response.
	LoginFailureTokenIntrospectionErroredCount *int64 `json:"loginFailureTokenIntrospectionErroredCount,omitempty"`
	// The number of login failures due to a failure to complete the token introspection.
	LoginFailureTokenIntrospectionFailureCount *int64 `json:"loginFailureTokenIntrospectionFailureCount,omitempty"`
	// The number of login failures due to a token introspection HTTPS error.
	LoginFailureTokenIntrospectionHttpsErrorCount *int64 `json:"loginFailureTokenIntrospectionHttpsErrorCount,omitempty"`
	// The number of login failures due to a token introspection response being invalid.
	LoginFailureTokenIntrospectionInvalidCount *int64 `json:"loginFailureTokenIntrospectionInvalidCount,omitempty"`
	// The number of login failures due to a token introspection timeout.
	LoginFailureTokenIntrospectionTimeoutCount *int64 `json:"loginFailureTokenIntrospectionTimeoutCount,omitempty"`
	// The number of login failures due to a token not being valid yet.
	LoginFailureTokenNotValidYetCount *int64 `json:"loginFailureTokenNotValidYetCount,omitempty"`
	// The number of login failures due to an unsupported algorithm.
	LoginFailureUnsupportedAlgCount *int64 `json:"loginFailureUnsupportedAlgCount,omitempty"`
	// The number of clients that did not provide an authorization group claim value when expected.
	MissingAuthorizationGroupCount *int64 `json:"missingAuthorizationGroupCount,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// The name of the OAuth Provider.
	OauthProviderName *string `json:"oauthProviderName,omitempty"`
	// Indicates whether to ignore time limits and accept tokens that are not yet valid or are no longer valid.
	TokenIgnoreTimeLimitsEnabled *bool `json:"tokenIgnoreTimeLimitsEnabled,omitempty"`
	// The one minute average of the time required to complete a token introspection, in milliseconds (ms).
	TokenIntrospectionAverageTime *int32 `json:"tokenIntrospectionAverageTime,omitempty"`
	// The reason for the last token introspection failure.
	TokenIntrospectionLastFailureReason *string `json:"tokenIntrospectionLastFailureReason,omitempty"`
	// The timestamp of the last token introspection failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	TokenIntrospectionLastFailureTime *int32 `json:"tokenIntrospectionLastFailureTime,omitempty"`
	// The parameter name used to identify the token during access token introspection. A standards compliant OAuth introspection server expects \"token\".
	TokenIntrospectionParameterName *string `json:"tokenIntrospectionParameterName,omitempty"`
	// The number of token introspection successes.
	TokenIntrospectionSuccessCount *int64 `json:"tokenIntrospectionSuccessCount,omitempty"`
	// The maximum time in seconds a token introspection is allowed to take.
	TokenIntrospectionTimeout *int32 `json:"tokenIntrospectionTimeout,omitempty"`
	// The token introspection URI of the OAuth authentication server.
	TokenIntrospectionUri *string `json:"tokenIntrospectionUri,omitempty"`
	// The username to use when logging into the token introspection URI.
	TokenIntrospectionUsername *string `json:"tokenIntrospectionUsername,omitempty"`
	// The username claim name, indicating which part of the object to use for determining the username.
	UsernameClaimName *string `json:"usernameClaimName,omitempty"`
	// The username claim source, indicating where to search for the username value. The allowed values and their meaning are:  <pre> \"access-token\" - The OAuth v2 access_token. \"id-token\" - The OpenID Connect id_token. \"introspection\" - The result of introspecting the OAuth v2 access_token. </pre>
	UsernameClaimSource *string `json:"usernameClaimSource,omitempty"`
	// Indicates whether the API provided username will be validated against the username calculated from the token(s).
	UsernameValidateEnabled *bool `json:"usernameValidateEnabled,omitempty"`
}

// NewMsgVpnAuthenticationOauthProvider instantiates a new MsgVpnAuthenticationOauthProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnAuthenticationOauthProvider() *MsgVpnAuthenticationOauthProvider {
	this := MsgVpnAuthenticationOauthProvider{}
	return &this
}

// NewMsgVpnAuthenticationOauthProviderWithDefaults instantiates a new MsgVpnAuthenticationOauthProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnAuthenticationOauthProviderWithDefaults() *MsgVpnAuthenticationOauthProvider {
	this := MsgVpnAuthenticationOauthProvider{}
	return &this
}

// GetAudienceClaimName returns the AudienceClaimName field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimName() string {
	if o == nil || o.AudienceClaimName == nil {
		var ret string
		return ret
	}
	return *o.AudienceClaimName
}

// GetAudienceClaimNameOk returns a tuple with the AudienceClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimNameOk() (*string, bool) {
	if o == nil || o.AudienceClaimName == nil {
		return nil, false
	}
	return o.AudienceClaimName, true
}

// HasAudienceClaimName returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasAudienceClaimName() bool {
	if o != nil && o.AudienceClaimName != nil {
		return true
	}

	return false
}

// SetAudienceClaimName gets a reference to the given string and assigns it to the AudienceClaimName field.
func (o *MsgVpnAuthenticationOauthProvider) SetAudienceClaimName(v string) {
	o.AudienceClaimName = &v
}

// GetAudienceClaimSource returns the AudienceClaimSource field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimSource() string {
	if o == nil || o.AudienceClaimSource == nil {
		var ret string
		return ret
	}
	return *o.AudienceClaimSource
}

// GetAudienceClaimSourceOk returns a tuple with the AudienceClaimSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimSourceOk() (*string, bool) {
	if o == nil || o.AudienceClaimSource == nil {
		return nil, false
	}
	return o.AudienceClaimSource, true
}

// HasAudienceClaimSource returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasAudienceClaimSource() bool {
	if o != nil && o.AudienceClaimSource != nil {
		return true
	}

	return false
}

// SetAudienceClaimSource gets a reference to the given string and assigns it to the AudienceClaimSource field.
func (o *MsgVpnAuthenticationOauthProvider) SetAudienceClaimSource(v string) {
	o.AudienceClaimSource = &v
}

// GetAudienceClaimValue returns the AudienceClaimValue field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimValue() string {
	if o == nil || o.AudienceClaimValue == nil {
		var ret string
		return ret
	}
	return *o.AudienceClaimValue
}

// GetAudienceClaimValueOk returns a tuple with the AudienceClaimValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimValueOk() (*string, bool) {
	if o == nil || o.AudienceClaimValue == nil {
		return nil, false
	}
	return o.AudienceClaimValue, true
}

// HasAudienceClaimValue returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasAudienceClaimValue() bool {
	if o != nil && o.AudienceClaimValue != nil {
		return true
	}

	return false
}

// SetAudienceClaimValue gets a reference to the given string and assigns it to the AudienceClaimValue field.
func (o *MsgVpnAuthenticationOauthProvider) SetAudienceClaimValue(v string) {
	o.AudienceClaimValue = &v
}

// GetAudienceValidationEnabled returns the AudienceValidationEnabled field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetAudienceValidationEnabled() bool {
	if o == nil || o.AudienceValidationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AudienceValidationEnabled
}

// GetAudienceValidationEnabledOk returns a tuple with the AudienceValidationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetAudienceValidationEnabledOk() (*bool, bool) {
	if o == nil || o.AudienceValidationEnabled == nil {
		return nil, false
	}
	return o.AudienceValidationEnabled, true
}

// HasAudienceValidationEnabled returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasAudienceValidationEnabled() bool {
	if o != nil && o.AudienceValidationEnabled != nil {
		return true
	}

	return false
}

// SetAudienceValidationEnabled gets a reference to the given bool and assigns it to the AudienceValidationEnabled field.
func (o *MsgVpnAuthenticationOauthProvider) SetAudienceValidationEnabled(v bool) {
	o.AudienceValidationEnabled = &v
}

// GetAuthenticationSuccessCount returns the AuthenticationSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetAuthenticationSuccessCount() int64 {
	if o == nil || o.AuthenticationSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.AuthenticationSuccessCount
}

// GetAuthenticationSuccessCountOk returns a tuple with the AuthenticationSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetAuthenticationSuccessCountOk() (*int64, bool) {
	if o == nil || o.AuthenticationSuccessCount == nil {
		return nil, false
	}
	return o.AuthenticationSuccessCount, true
}

// HasAuthenticationSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasAuthenticationSuccessCount() bool {
	if o != nil && o.AuthenticationSuccessCount != nil {
		return true
	}

	return false
}

// SetAuthenticationSuccessCount gets a reference to the given int64 and assigns it to the AuthenticationSuccessCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetAuthenticationSuccessCount(v int64) {
	o.AuthenticationSuccessCount = &v
}

// GetAuthorizationGroupClaimName returns the AuthorizationGroupClaimName field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupClaimName() string {
	if o == nil || o.AuthorizationGroupClaimName == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationGroupClaimName
}

// GetAuthorizationGroupClaimNameOk returns a tuple with the AuthorizationGroupClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupClaimNameOk() (*string, bool) {
	if o == nil || o.AuthorizationGroupClaimName == nil {
		return nil, false
	}
	return o.AuthorizationGroupClaimName, true
}

// HasAuthorizationGroupClaimName returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasAuthorizationGroupClaimName() bool {
	if o != nil && o.AuthorizationGroupClaimName != nil {
		return true
	}

	return false
}

// SetAuthorizationGroupClaimName gets a reference to the given string and assigns it to the AuthorizationGroupClaimName field.
func (o *MsgVpnAuthenticationOauthProvider) SetAuthorizationGroupClaimName(v string) {
	o.AuthorizationGroupClaimName = &v
}

// GetAuthorizationGroupClaimSource returns the AuthorizationGroupClaimSource field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupClaimSource() string {
	if o == nil || o.AuthorizationGroupClaimSource == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationGroupClaimSource
}

// GetAuthorizationGroupClaimSourceOk returns a tuple with the AuthorizationGroupClaimSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupClaimSourceOk() (*string, bool) {
	if o == nil || o.AuthorizationGroupClaimSource == nil {
		return nil, false
	}
	return o.AuthorizationGroupClaimSource, true
}

// HasAuthorizationGroupClaimSource returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasAuthorizationGroupClaimSource() bool {
	if o != nil && o.AuthorizationGroupClaimSource != nil {
		return true
	}

	return false
}

// SetAuthorizationGroupClaimSource gets a reference to the given string and assigns it to the AuthorizationGroupClaimSource field.
func (o *MsgVpnAuthenticationOauthProvider) SetAuthorizationGroupClaimSource(v string) {
	o.AuthorizationGroupClaimSource = &v
}

// GetAuthorizationGroupEnabled returns the AuthorizationGroupEnabled field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupEnabled() bool {
	if o == nil || o.AuthorizationGroupEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthorizationGroupEnabled
}

// GetAuthorizationGroupEnabledOk returns a tuple with the AuthorizationGroupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupEnabledOk() (*bool, bool) {
	if o == nil || o.AuthorizationGroupEnabled == nil {
		return nil, false
	}
	return o.AuthorizationGroupEnabled, true
}

// HasAuthorizationGroupEnabled returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasAuthorizationGroupEnabled() bool {
	if o != nil && o.AuthorizationGroupEnabled != nil {
		return true
	}

	return false
}

// SetAuthorizationGroupEnabled gets a reference to the given bool and assigns it to the AuthorizationGroupEnabled field.
func (o *MsgVpnAuthenticationOauthProvider) SetAuthorizationGroupEnabled(v bool) {
	o.AuthorizationGroupEnabled = &v
}

// GetDisconnectOnTokenExpirationEnabled returns the DisconnectOnTokenExpirationEnabled field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetDisconnectOnTokenExpirationEnabled() bool {
	if o == nil || o.DisconnectOnTokenExpirationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DisconnectOnTokenExpirationEnabled
}

// GetDisconnectOnTokenExpirationEnabledOk returns a tuple with the DisconnectOnTokenExpirationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetDisconnectOnTokenExpirationEnabledOk() (*bool, bool) {
	if o == nil || o.DisconnectOnTokenExpirationEnabled == nil {
		return nil, false
	}
	return o.DisconnectOnTokenExpirationEnabled, true
}

// HasDisconnectOnTokenExpirationEnabled returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasDisconnectOnTokenExpirationEnabled() bool {
	if o != nil && o.DisconnectOnTokenExpirationEnabled != nil {
		return true
	}

	return false
}

// SetDisconnectOnTokenExpirationEnabled gets a reference to the given bool and assigns it to the DisconnectOnTokenExpirationEnabled field.
func (o *MsgVpnAuthenticationOauthProvider) SetDisconnectOnTokenExpirationEnabled(v bool) {
	o.DisconnectOnTokenExpirationEnabled = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MsgVpnAuthenticationOauthProvider) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetJwksLastRefreshFailureReason returns the JwksLastRefreshFailureReason field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshFailureReason() string {
	if o == nil || o.JwksLastRefreshFailureReason == nil {
		var ret string
		return ret
	}
	return *o.JwksLastRefreshFailureReason
}

// GetJwksLastRefreshFailureReasonOk returns a tuple with the JwksLastRefreshFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshFailureReasonOk() (*string, bool) {
	if o == nil || o.JwksLastRefreshFailureReason == nil {
		return nil, false
	}
	return o.JwksLastRefreshFailureReason, true
}

// HasJwksLastRefreshFailureReason returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasJwksLastRefreshFailureReason() bool {
	if o != nil && o.JwksLastRefreshFailureReason != nil {
		return true
	}

	return false
}

// SetJwksLastRefreshFailureReason gets a reference to the given string and assigns it to the JwksLastRefreshFailureReason field.
func (o *MsgVpnAuthenticationOauthProvider) SetJwksLastRefreshFailureReason(v string) {
	o.JwksLastRefreshFailureReason = &v
}

// GetJwksLastRefreshFailureTime returns the JwksLastRefreshFailureTime field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshFailureTime() int32 {
	if o == nil || o.JwksLastRefreshFailureTime == nil {
		var ret int32
		return ret
	}
	return *o.JwksLastRefreshFailureTime
}

// GetJwksLastRefreshFailureTimeOk returns a tuple with the JwksLastRefreshFailureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshFailureTimeOk() (*int32, bool) {
	if o == nil || o.JwksLastRefreshFailureTime == nil {
		return nil, false
	}
	return o.JwksLastRefreshFailureTime, true
}

// HasJwksLastRefreshFailureTime returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasJwksLastRefreshFailureTime() bool {
	if o != nil && o.JwksLastRefreshFailureTime != nil {
		return true
	}

	return false
}

// SetJwksLastRefreshFailureTime gets a reference to the given int32 and assigns it to the JwksLastRefreshFailureTime field.
func (o *MsgVpnAuthenticationOauthProvider) SetJwksLastRefreshFailureTime(v int32) {
	o.JwksLastRefreshFailureTime = &v
}

// GetJwksLastRefreshTime returns the JwksLastRefreshTime field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshTime() int32 {
	if o == nil || o.JwksLastRefreshTime == nil {
		var ret int32
		return ret
	}
	return *o.JwksLastRefreshTime
}

// GetJwksLastRefreshTimeOk returns a tuple with the JwksLastRefreshTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshTimeOk() (*int32, bool) {
	if o == nil || o.JwksLastRefreshTime == nil {
		return nil, false
	}
	return o.JwksLastRefreshTime, true
}

// HasJwksLastRefreshTime returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasJwksLastRefreshTime() bool {
	if o != nil && o.JwksLastRefreshTime != nil {
		return true
	}

	return false
}

// SetJwksLastRefreshTime gets a reference to the given int32 and assigns it to the JwksLastRefreshTime field.
func (o *MsgVpnAuthenticationOauthProvider) SetJwksLastRefreshTime(v int32) {
	o.JwksLastRefreshTime = &v
}

// GetJwksNextScheduledRefreshTime returns the JwksNextScheduledRefreshTime field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksNextScheduledRefreshTime() int32 {
	if o == nil || o.JwksNextScheduledRefreshTime == nil {
		var ret int32
		return ret
	}
	return *o.JwksNextScheduledRefreshTime
}

// GetJwksNextScheduledRefreshTimeOk returns a tuple with the JwksNextScheduledRefreshTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksNextScheduledRefreshTimeOk() (*int32, bool) {
	if o == nil || o.JwksNextScheduledRefreshTime == nil {
		return nil, false
	}
	return o.JwksNextScheduledRefreshTime, true
}

// HasJwksNextScheduledRefreshTime returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasJwksNextScheduledRefreshTime() bool {
	if o != nil && o.JwksNextScheduledRefreshTime != nil {
		return true
	}

	return false
}

// SetJwksNextScheduledRefreshTime gets a reference to the given int32 and assigns it to the JwksNextScheduledRefreshTime field.
func (o *MsgVpnAuthenticationOauthProvider) SetJwksNextScheduledRefreshTime(v int32) {
	o.JwksNextScheduledRefreshTime = &v
}

// GetJwksRefreshFailureCount returns the JwksRefreshFailureCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksRefreshFailureCount() int64 {
	if o == nil || o.JwksRefreshFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.JwksRefreshFailureCount
}

// GetJwksRefreshFailureCountOk returns a tuple with the JwksRefreshFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksRefreshFailureCountOk() (*int64, bool) {
	if o == nil || o.JwksRefreshFailureCount == nil {
		return nil, false
	}
	return o.JwksRefreshFailureCount, true
}

// HasJwksRefreshFailureCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasJwksRefreshFailureCount() bool {
	if o != nil && o.JwksRefreshFailureCount != nil {
		return true
	}

	return false
}

// SetJwksRefreshFailureCount gets a reference to the given int64 and assigns it to the JwksRefreshFailureCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetJwksRefreshFailureCount(v int64) {
	o.JwksRefreshFailureCount = &v
}

// GetJwksRefreshInterval returns the JwksRefreshInterval field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksRefreshInterval() int32 {
	if o == nil || o.JwksRefreshInterval == nil {
		var ret int32
		return ret
	}
	return *o.JwksRefreshInterval
}

// GetJwksRefreshIntervalOk returns a tuple with the JwksRefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksRefreshIntervalOk() (*int32, bool) {
	if o == nil || o.JwksRefreshInterval == nil {
		return nil, false
	}
	return o.JwksRefreshInterval, true
}

// HasJwksRefreshInterval returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasJwksRefreshInterval() bool {
	if o != nil && o.JwksRefreshInterval != nil {
		return true
	}

	return false
}

// SetJwksRefreshInterval gets a reference to the given int32 and assigns it to the JwksRefreshInterval field.
func (o *MsgVpnAuthenticationOauthProvider) SetJwksRefreshInterval(v int32) {
	o.JwksRefreshInterval = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksUri() string {
	if o == nil || o.JwksUri == nil {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetJwksUriOk() (*string, bool) {
	if o == nil || o.JwksUri == nil {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasJwksUri() bool {
	if o != nil && o.JwksUri != nil {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *MsgVpnAuthenticationOauthProvider) SetJwksUri(v string) {
	o.JwksUri = &v
}

// GetLoginFailureIncorrectAudienceValueCount returns the LoginFailureIncorrectAudienceValueCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureIncorrectAudienceValueCount() int64 {
	if o == nil || o.LoginFailureIncorrectAudienceValueCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureIncorrectAudienceValueCount
}

// GetLoginFailureIncorrectAudienceValueCountOk returns a tuple with the LoginFailureIncorrectAudienceValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureIncorrectAudienceValueCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureIncorrectAudienceValueCount == nil {
		return nil, false
	}
	return o.LoginFailureIncorrectAudienceValueCount, true
}

// HasLoginFailureIncorrectAudienceValueCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureIncorrectAudienceValueCount() bool {
	if o != nil && o.LoginFailureIncorrectAudienceValueCount != nil {
		return true
	}

	return false
}

// SetLoginFailureIncorrectAudienceValueCount gets a reference to the given int64 and assigns it to the LoginFailureIncorrectAudienceValueCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureIncorrectAudienceValueCount(v int64) {
	o.LoginFailureIncorrectAudienceValueCount = &v
}

// GetLoginFailureInvalidAudienceValueCount returns the LoginFailureInvalidAudienceValueCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidAudienceValueCount() int64 {
	if o == nil || o.LoginFailureInvalidAudienceValueCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureInvalidAudienceValueCount
}

// GetLoginFailureInvalidAudienceValueCountOk returns a tuple with the LoginFailureInvalidAudienceValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidAudienceValueCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureInvalidAudienceValueCount == nil {
		return nil, false
	}
	return o.LoginFailureInvalidAudienceValueCount, true
}

// HasLoginFailureInvalidAudienceValueCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureInvalidAudienceValueCount() bool {
	if o != nil && o.LoginFailureInvalidAudienceValueCount != nil {
		return true
	}

	return false
}

// SetLoginFailureInvalidAudienceValueCount gets a reference to the given int64 and assigns it to the LoginFailureInvalidAudienceValueCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureInvalidAudienceValueCount(v int64) {
	o.LoginFailureInvalidAudienceValueCount = &v
}

// GetLoginFailureInvalidAuthorizationGroupValueCount returns the LoginFailureInvalidAuthorizationGroupValueCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidAuthorizationGroupValueCount() int64 {
	if o == nil || o.LoginFailureInvalidAuthorizationGroupValueCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureInvalidAuthorizationGroupValueCount
}

// GetLoginFailureInvalidAuthorizationGroupValueCountOk returns a tuple with the LoginFailureInvalidAuthorizationGroupValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidAuthorizationGroupValueCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureInvalidAuthorizationGroupValueCount == nil {
		return nil, false
	}
	return o.LoginFailureInvalidAuthorizationGroupValueCount, true
}

// HasLoginFailureInvalidAuthorizationGroupValueCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureInvalidAuthorizationGroupValueCount() bool {
	if o != nil && o.LoginFailureInvalidAuthorizationGroupValueCount != nil {
		return true
	}

	return false
}

// SetLoginFailureInvalidAuthorizationGroupValueCount gets a reference to the given int64 and assigns it to the LoginFailureInvalidAuthorizationGroupValueCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureInvalidAuthorizationGroupValueCount(v int64) {
	o.LoginFailureInvalidAuthorizationGroupValueCount = &v
}

// GetLoginFailureInvalidJwtSignatureCount returns the LoginFailureInvalidJwtSignatureCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidJwtSignatureCount() int64 {
	if o == nil || o.LoginFailureInvalidJwtSignatureCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureInvalidJwtSignatureCount
}

// GetLoginFailureInvalidJwtSignatureCountOk returns a tuple with the LoginFailureInvalidJwtSignatureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidJwtSignatureCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureInvalidJwtSignatureCount == nil {
		return nil, false
	}
	return o.LoginFailureInvalidJwtSignatureCount, true
}

// HasLoginFailureInvalidJwtSignatureCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureInvalidJwtSignatureCount() bool {
	if o != nil && o.LoginFailureInvalidJwtSignatureCount != nil {
		return true
	}

	return false
}

// SetLoginFailureInvalidJwtSignatureCount gets a reference to the given int64 and assigns it to the LoginFailureInvalidJwtSignatureCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureInvalidJwtSignatureCount(v int64) {
	o.LoginFailureInvalidJwtSignatureCount = &v
}

// GetLoginFailureInvalidUsernameValueCount returns the LoginFailureInvalidUsernameValueCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidUsernameValueCount() int64 {
	if o == nil || o.LoginFailureInvalidUsernameValueCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureInvalidUsernameValueCount
}

// GetLoginFailureInvalidUsernameValueCountOk returns a tuple with the LoginFailureInvalidUsernameValueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidUsernameValueCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureInvalidUsernameValueCount == nil {
		return nil, false
	}
	return o.LoginFailureInvalidUsernameValueCount, true
}

// HasLoginFailureInvalidUsernameValueCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureInvalidUsernameValueCount() bool {
	if o != nil && o.LoginFailureInvalidUsernameValueCount != nil {
		return true
	}

	return false
}

// SetLoginFailureInvalidUsernameValueCount gets a reference to the given int64 and assigns it to the LoginFailureInvalidUsernameValueCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureInvalidUsernameValueCount(v int64) {
	o.LoginFailureInvalidUsernameValueCount = &v
}

// GetLoginFailureMismatchedUsernameCount returns the LoginFailureMismatchedUsernameCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMismatchedUsernameCount() int64 {
	if o == nil || o.LoginFailureMismatchedUsernameCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureMismatchedUsernameCount
}

// GetLoginFailureMismatchedUsernameCountOk returns a tuple with the LoginFailureMismatchedUsernameCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMismatchedUsernameCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureMismatchedUsernameCount == nil {
		return nil, false
	}
	return o.LoginFailureMismatchedUsernameCount, true
}

// HasLoginFailureMismatchedUsernameCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureMismatchedUsernameCount() bool {
	if o != nil && o.LoginFailureMismatchedUsernameCount != nil {
		return true
	}

	return false
}

// SetLoginFailureMismatchedUsernameCount gets a reference to the given int64 and assigns it to the LoginFailureMismatchedUsernameCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureMismatchedUsernameCount(v int64) {
	o.LoginFailureMismatchedUsernameCount = &v
}

// GetLoginFailureMissingAudienceCount returns the LoginFailureMissingAudienceCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingAudienceCount() int64 {
	if o == nil || o.LoginFailureMissingAudienceCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureMissingAudienceCount
}

// GetLoginFailureMissingAudienceCountOk returns a tuple with the LoginFailureMissingAudienceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingAudienceCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureMissingAudienceCount == nil {
		return nil, false
	}
	return o.LoginFailureMissingAudienceCount, true
}

// HasLoginFailureMissingAudienceCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureMissingAudienceCount() bool {
	if o != nil && o.LoginFailureMissingAudienceCount != nil {
		return true
	}

	return false
}

// SetLoginFailureMissingAudienceCount gets a reference to the given int64 and assigns it to the LoginFailureMissingAudienceCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureMissingAudienceCount(v int64) {
	o.LoginFailureMissingAudienceCount = &v
}

// GetLoginFailureMissingJwkCount returns the LoginFailureMissingJwkCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingJwkCount() int64 {
	if o == nil || o.LoginFailureMissingJwkCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureMissingJwkCount
}

// GetLoginFailureMissingJwkCountOk returns a tuple with the LoginFailureMissingJwkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingJwkCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureMissingJwkCount == nil {
		return nil, false
	}
	return o.LoginFailureMissingJwkCount, true
}

// HasLoginFailureMissingJwkCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureMissingJwkCount() bool {
	if o != nil && o.LoginFailureMissingJwkCount != nil {
		return true
	}

	return false
}

// SetLoginFailureMissingJwkCount gets a reference to the given int64 and assigns it to the LoginFailureMissingJwkCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureMissingJwkCount(v int64) {
	o.LoginFailureMissingJwkCount = &v
}

// GetLoginFailureMissingOrInvalidTokenCount returns the LoginFailureMissingOrInvalidTokenCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingOrInvalidTokenCount() int64 {
	if o == nil || o.LoginFailureMissingOrInvalidTokenCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureMissingOrInvalidTokenCount
}

// GetLoginFailureMissingOrInvalidTokenCountOk returns a tuple with the LoginFailureMissingOrInvalidTokenCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingOrInvalidTokenCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureMissingOrInvalidTokenCount == nil {
		return nil, false
	}
	return o.LoginFailureMissingOrInvalidTokenCount, true
}

// HasLoginFailureMissingOrInvalidTokenCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureMissingOrInvalidTokenCount() bool {
	if o != nil && o.LoginFailureMissingOrInvalidTokenCount != nil {
		return true
	}

	return false
}

// SetLoginFailureMissingOrInvalidTokenCount gets a reference to the given int64 and assigns it to the LoginFailureMissingOrInvalidTokenCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureMissingOrInvalidTokenCount(v int64) {
	o.LoginFailureMissingOrInvalidTokenCount = &v
}

// GetLoginFailureMissingUsernameCount returns the LoginFailureMissingUsernameCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingUsernameCount() int64 {
	if o == nil || o.LoginFailureMissingUsernameCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureMissingUsernameCount
}

// GetLoginFailureMissingUsernameCountOk returns a tuple with the LoginFailureMissingUsernameCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingUsernameCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureMissingUsernameCount == nil {
		return nil, false
	}
	return o.LoginFailureMissingUsernameCount, true
}

// HasLoginFailureMissingUsernameCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureMissingUsernameCount() bool {
	if o != nil && o.LoginFailureMissingUsernameCount != nil {
		return true
	}

	return false
}

// SetLoginFailureMissingUsernameCount gets a reference to the given int64 and assigns it to the LoginFailureMissingUsernameCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureMissingUsernameCount(v int64) {
	o.LoginFailureMissingUsernameCount = &v
}

// GetLoginFailureTokenExpiredCount returns the LoginFailureTokenExpiredCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenExpiredCount() int64 {
	if o == nil || o.LoginFailureTokenExpiredCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureTokenExpiredCount
}

// GetLoginFailureTokenExpiredCountOk returns a tuple with the LoginFailureTokenExpiredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenExpiredCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureTokenExpiredCount == nil {
		return nil, false
	}
	return o.LoginFailureTokenExpiredCount, true
}

// HasLoginFailureTokenExpiredCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenExpiredCount() bool {
	if o != nil && o.LoginFailureTokenExpiredCount != nil {
		return true
	}

	return false
}

// SetLoginFailureTokenExpiredCount gets a reference to the given int64 and assigns it to the LoginFailureTokenExpiredCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenExpiredCount(v int64) {
	o.LoginFailureTokenExpiredCount = &v
}

// GetLoginFailureTokenIntrospectionErroredCount returns the LoginFailureTokenIntrospectionErroredCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionErroredCount() int64 {
	if o == nil || o.LoginFailureTokenIntrospectionErroredCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureTokenIntrospectionErroredCount
}

// GetLoginFailureTokenIntrospectionErroredCountOk returns a tuple with the LoginFailureTokenIntrospectionErroredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionErroredCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureTokenIntrospectionErroredCount == nil {
		return nil, false
	}
	return o.LoginFailureTokenIntrospectionErroredCount, true
}

// HasLoginFailureTokenIntrospectionErroredCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenIntrospectionErroredCount() bool {
	if o != nil && o.LoginFailureTokenIntrospectionErroredCount != nil {
		return true
	}

	return false
}

// SetLoginFailureTokenIntrospectionErroredCount gets a reference to the given int64 and assigns it to the LoginFailureTokenIntrospectionErroredCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenIntrospectionErroredCount(v int64) {
	o.LoginFailureTokenIntrospectionErroredCount = &v
}

// GetLoginFailureTokenIntrospectionFailureCount returns the LoginFailureTokenIntrospectionFailureCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionFailureCount() int64 {
	if o == nil || o.LoginFailureTokenIntrospectionFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureTokenIntrospectionFailureCount
}

// GetLoginFailureTokenIntrospectionFailureCountOk returns a tuple with the LoginFailureTokenIntrospectionFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionFailureCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureTokenIntrospectionFailureCount == nil {
		return nil, false
	}
	return o.LoginFailureTokenIntrospectionFailureCount, true
}

// HasLoginFailureTokenIntrospectionFailureCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenIntrospectionFailureCount() bool {
	if o != nil && o.LoginFailureTokenIntrospectionFailureCount != nil {
		return true
	}

	return false
}

// SetLoginFailureTokenIntrospectionFailureCount gets a reference to the given int64 and assigns it to the LoginFailureTokenIntrospectionFailureCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenIntrospectionFailureCount(v int64) {
	o.LoginFailureTokenIntrospectionFailureCount = &v
}

// GetLoginFailureTokenIntrospectionHttpsErrorCount returns the LoginFailureTokenIntrospectionHttpsErrorCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionHttpsErrorCount() int64 {
	if o == nil || o.LoginFailureTokenIntrospectionHttpsErrorCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureTokenIntrospectionHttpsErrorCount
}

// GetLoginFailureTokenIntrospectionHttpsErrorCountOk returns a tuple with the LoginFailureTokenIntrospectionHttpsErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionHttpsErrorCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureTokenIntrospectionHttpsErrorCount == nil {
		return nil, false
	}
	return o.LoginFailureTokenIntrospectionHttpsErrorCount, true
}

// HasLoginFailureTokenIntrospectionHttpsErrorCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenIntrospectionHttpsErrorCount() bool {
	if o != nil && o.LoginFailureTokenIntrospectionHttpsErrorCount != nil {
		return true
	}

	return false
}

// SetLoginFailureTokenIntrospectionHttpsErrorCount gets a reference to the given int64 and assigns it to the LoginFailureTokenIntrospectionHttpsErrorCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenIntrospectionHttpsErrorCount(v int64) {
	o.LoginFailureTokenIntrospectionHttpsErrorCount = &v
}

// GetLoginFailureTokenIntrospectionInvalidCount returns the LoginFailureTokenIntrospectionInvalidCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionInvalidCount() int64 {
	if o == nil || o.LoginFailureTokenIntrospectionInvalidCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureTokenIntrospectionInvalidCount
}

// GetLoginFailureTokenIntrospectionInvalidCountOk returns a tuple with the LoginFailureTokenIntrospectionInvalidCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionInvalidCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureTokenIntrospectionInvalidCount == nil {
		return nil, false
	}
	return o.LoginFailureTokenIntrospectionInvalidCount, true
}

// HasLoginFailureTokenIntrospectionInvalidCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenIntrospectionInvalidCount() bool {
	if o != nil && o.LoginFailureTokenIntrospectionInvalidCount != nil {
		return true
	}

	return false
}

// SetLoginFailureTokenIntrospectionInvalidCount gets a reference to the given int64 and assigns it to the LoginFailureTokenIntrospectionInvalidCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenIntrospectionInvalidCount(v int64) {
	o.LoginFailureTokenIntrospectionInvalidCount = &v
}

// GetLoginFailureTokenIntrospectionTimeoutCount returns the LoginFailureTokenIntrospectionTimeoutCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionTimeoutCount() int64 {
	if o == nil || o.LoginFailureTokenIntrospectionTimeoutCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureTokenIntrospectionTimeoutCount
}

// GetLoginFailureTokenIntrospectionTimeoutCountOk returns a tuple with the LoginFailureTokenIntrospectionTimeoutCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionTimeoutCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureTokenIntrospectionTimeoutCount == nil {
		return nil, false
	}
	return o.LoginFailureTokenIntrospectionTimeoutCount, true
}

// HasLoginFailureTokenIntrospectionTimeoutCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenIntrospectionTimeoutCount() bool {
	if o != nil && o.LoginFailureTokenIntrospectionTimeoutCount != nil {
		return true
	}

	return false
}

// SetLoginFailureTokenIntrospectionTimeoutCount gets a reference to the given int64 and assigns it to the LoginFailureTokenIntrospectionTimeoutCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenIntrospectionTimeoutCount(v int64) {
	o.LoginFailureTokenIntrospectionTimeoutCount = &v
}

// GetLoginFailureTokenNotValidYetCount returns the LoginFailureTokenNotValidYetCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenNotValidYetCount() int64 {
	if o == nil || o.LoginFailureTokenNotValidYetCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureTokenNotValidYetCount
}

// GetLoginFailureTokenNotValidYetCountOk returns a tuple with the LoginFailureTokenNotValidYetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenNotValidYetCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureTokenNotValidYetCount == nil {
		return nil, false
	}
	return o.LoginFailureTokenNotValidYetCount, true
}

// HasLoginFailureTokenNotValidYetCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenNotValidYetCount() bool {
	if o != nil && o.LoginFailureTokenNotValidYetCount != nil {
		return true
	}

	return false
}

// SetLoginFailureTokenNotValidYetCount gets a reference to the given int64 and assigns it to the LoginFailureTokenNotValidYetCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenNotValidYetCount(v int64) {
	o.LoginFailureTokenNotValidYetCount = &v
}

// GetLoginFailureUnsupportedAlgCount returns the LoginFailureUnsupportedAlgCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureUnsupportedAlgCount() int64 {
	if o == nil || o.LoginFailureUnsupportedAlgCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginFailureUnsupportedAlgCount
}

// GetLoginFailureUnsupportedAlgCountOk returns a tuple with the LoginFailureUnsupportedAlgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureUnsupportedAlgCountOk() (*int64, bool) {
	if o == nil || o.LoginFailureUnsupportedAlgCount == nil {
		return nil, false
	}
	return o.LoginFailureUnsupportedAlgCount, true
}

// HasLoginFailureUnsupportedAlgCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureUnsupportedAlgCount() bool {
	if o != nil && o.LoginFailureUnsupportedAlgCount != nil {
		return true
	}

	return false
}

// SetLoginFailureUnsupportedAlgCount gets a reference to the given int64 and assigns it to the LoginFailureUnsupportedAlgCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureUnsupportedAlgCount(v int64) {
	o.LoginFailureUnsupportedAlgCount = &v
}

// GetMissingAuthorizationGroupCount returns the MissingAuthorizationGroupCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetMissingAuthorizationGroupCount() int64 {
	if o == nil || o.MissingAuthorizationGroupCount == nil {
		var ret int64
		return ret
	}
	return *o.MissingAuthorizationGroupCount
}

// GetMissingAuthorizationGroupCountOk returns a tuple with the MissingAuthorizationGroupCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetMissingAuthorizationGroupCountOk() (*int64, bool) {
	if o == nil || o.MissingAuthorizationGroupCount == nil {
		return nil, false
	}
	return o.MissingAuthorizationGroupCount, true
}

// HasMissingAuthorizationGroupCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasMissingAuthorizationGroupCount() bool {
	if o != nil && o.MissingAuthorizationGroupCount != nil {
		return true
	}

	return false
}

// SetMissingAuthorizationGroupCount gets a reference to the given int64 and assigns it to the MissingAuthorizationGroupCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetMissingAuthorizationGroupCount(v int64) {
	o.MissingAuthorizationGroupCount = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnAuthenticationOauthProvider) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetOauthProviderName returns the OauthProviderName field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetOauthProviderName() string {
	if o == nil || o.OauthProviderName == nil {
		var ret string
		return ret
	}
	return *o.OauthProviderName
}

// GetOauthProviderNameOk returns a tuple with the OauthProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetOauthProviderNameOk() (*string, bool) {
	if o == nil || o.OauthProviderName == nil {
		return nil, false
	}
	return o.OauthProviderName, true
}

// HasOauthProviderName returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasOauthProviderName() bool {
	if o != nil && o.OauthProviderName != nil {
		return true
	}

	return false
}

// SetOauthProviderName gets a reference to the given string and assigns it to the OauthProviderName field.
func (o *MsgVpnAuthenticationOauthProvider) SetOauthProviderName(v string) {
	o.OauthProviderName = &v
}

// GetTokenIgnoreTimeLimitsEnabled returns the TokenIgnoreTimeLimitsEnabled field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIgnoreTimeLimitsEnabled() bool {
	if o == nil || o.TokenIgnoreTimeLimitsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TokenIgnoreTimeLimitsEnabled
}

// GetTokenIgnoreTimeLimitsEnabledOk returns a tuple with the TokenIgnoreTimeLimitsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIgnoreTimeLimitsEnabledOk() (*bool, bool) {
	if o == nil || o.TokenIgnoreTimeLimitsEnabled == nil {
		return nil, false
	}
	return o.TokenIgnoreTimeLimitsEnabled, true
}

// HasTokenIgnoreTimeLimitsEnabled returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasTokenIgnoreTimeLimitsEnabled() bool {
	if o != nil && o.TokenIgnoreTimeLimitsEnabled != nil {
		return true
	}

	return false
}

// SetTokenIgnoreTimeLimitsEnabled gets a reference to the given bool and assigns it to the TokenIgnoreTimeLimitsEnabled field.
func (o *MsgVpnAuthenticationOauthProvider) SetTokenIgnoreTimeLimitsEnabled(v bool) {
	o.TokenIgnoreTimeLimitsEnabled = &v
}

// GetTokenIntrospectionAverageTime returns the TokenIntrospectionAverageTime field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionAverageTime() int32 {
	if o == nil || o.TokenIntrospectionAverageTime == nil {
		var ret int32
		return ret
	}
	return *o.TokenIntrospectionAverageTime
}

// GetTokenIntrospectionAverageTimeOk returns a tuple with the TokenIntrospectionAverageTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionAverageTimeOk() (*int32, bool) {
	if o == nil || o.TokenIntrospectionAverageTime == nil {
		return nil, false
	}
	return o.TokenIntrospectionAverageTime, true
}

// HasTokenIntrospectionAverageTime returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionAverageTime() bool {
	if o != nil && o.TokenIntrospectionAverageTime != nil {
		return true
	}

	return false
}

// SetTokenIntrospectionAverageTime gets a reference to the given int32 and assigns it to the TokenIntrospectionAverageTime field.
func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionAverageTime(v int32) {
	o.TokenIntrospectionAverageTime = &v
}

// GetTokenIntrospectionLastFailureReason returns the TokenIntrospectionLastFailureReason field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionLastFailureReason() string {
	if o == nil || o.TokenIntrospectionLastFailureReason == nil {
		var ret string
		return ret
	}
	return *o.TokenIntrospectionLastFailureReason
}

// GetTokenIntrospectionLastFailureReasonOk returns a tuple with the TokenIntrospectionLastFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionLastFailureReasonOk() (*string, bool) {
	if o == nil || o.TokenIntrospectionLastFailureReason == nil {
		return nil, false
	}
	return o.TokenIntrospectionLastFailureReason, true
}

// HasTokenIntrospectionLastFailureReason returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionLastFailureReason() bool {
	if o != nil && o.TokenIntrospectionLastFailureReason != nil {
		return true
	}

	return false
}

// SetTokenIntrospectionLastFailureReason gets a reference to the given string and assigns it to the TokenIntrospectionLastFailureReason field.
func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionLastFailureReason(v string) {
	o.TokenIntrospectionLastFailureReason = &v
}

// GetTokenIntrospectionLastFailureTime returns the TokenIntrospectionLastFailureTime field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionLastFailureTime() int32 {
	if o == nil || o.TokenIntrospectionLastFailureTime == nil {
		var ret int32
		return ret
	}
	return *o.TokenIntrospectionLastFailureTime
}

// GetTokenIntrospectionLastFailureTimeOk returns a tuple with the TokenIntrospectionLastFailureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionLastFailureTimeOk() (*int32, bool) {
	if o == nil || o.TokenIntrospectionLastFailureTime == nil {
		return nil, false
	}
	return o.TokenIntrospectionLastFailureTime, true
}

// HasTokenIntrospectionLastFailureTime returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionLastFailureTime() bool {
	if o != nil && o.TokenIntrospectionLastFailureTime != nil {
		return true
	}

	return false
}

// SetTokenIntrospectionLastFailureTime gets a reference to the given int32 and assigns it to the TokenIntrospectionLastFailureTime field.
func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionLastFailureTime(v int32) {
	o.TokenIntrospectionLastFailureTime = &v
}

// GetTokenIntrospectionParameterName returns the TokenIntrospectionParameterName field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionParameterName() string {
	if o == nil || o.TokenIntrospectionParameterName == nil {
		var ret string
		return ret
	}
	return *o.TokenIntrospectionParameterName
}

// GetTokenIntrospectionParameterNameOk returns a tuple with the TokenIntrospectionParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionParameterNameOk() (*string, bool) {
	if o == nil || o.TokenIntrospectionParameterName == nil {
		return nil, false
	}
	return o.TokenIntrospectionParameterName, true
}

// HasTokenIntrospectionParameterName returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionParameterName() bool {
	if o != nil && o.TokenIntrospectionParameterName != nil {
		return true
	}

	return false
}

// SetTokenIntrospectionParameterName gets a reference to the given string and assigns it to the TokenIntrospectionParameterName field.
func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionParameterName(v string) {
	o.TokenIntrospectionParameterName = &v
}

// GetTokenIntrospectionSuccessCount returns the TokenIntrospectionSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionSuccessCount() int64 {
	if o == nil || o.TokenIntrospectionSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.TokenIntrospectionSuccessCount
}

// GetTokenIntrospectionSuccessCountOk returns a tuple with the TokenIntrospectionSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionSuccessCountOk() (*int64, bool) {
	if o == nil || o.TokenIntrospectionSuccessCount == nil {
		return nil, false
	}
	return o.TokenIntrospectionSuccessCount, true
}

// HasTokenIntrospectionSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionSuccessCount() bool {
	if o != nil && o.TokenIntrospectionSuccessCount != nil {
		return true
	}

	return false
}

// SetTokenIntrospectionSuccessCount gets a reference to the given int64 and assigns it to the TokenIntrospectionSuccessCount field.
func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionSuccessCount(v int64) {
	o.TokenIntrospectionSuccessCount = &v
}

// GetTokenIntrospectionTimeout returns the TokenIntrospectionTimeout field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionTimeout() int32 {
	if o == nil || o.TokenIntrospectionTimeout == nil {
		var ret int32
		return ret
	}
	return *o.TokenIntrospectionTimeout
}

// GetTokenIntrospectionTimeoutOk returns a tuple with the TokenIntrospectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionTimeoutOk() (*int32, bool) {
	if o == nil || o.TokenIntrospectionTimeout == nil {
		return nil, false
	}
	return o.TokenIntrospectionTimeout, true
}

// HasTokenIntrospectionTimeout returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionTimeout() bool {
	if o != nil && o.TokenIntrospectionTimeout != nil {
		return true
	}

	return false
}

// SetTokenIntrospectionTimeout gets a reference to the given int32 and assigns it to the TokenIntrospectionTimeout field.
func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionTimeout(v int32) {
	o.TokenIntrospectionTimeout = &v
}

// GetTokenIntrospectionUri returns the TokenIntrospectionUri field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionUri() string {
	if o == nil || o.TokenIntrospectionUri == nil {
		var ret string
		return ret
	}
	return *o.TokenIntrospectionUri
}

// GetTokenIntrospectionUriOk returns a tuple with the TokenIntrospectionUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionUriOk() (*string, bool) {
	if o == nil || o.TokenIntrospectionUri == nil {
		return nil, false
	}
	return o.TokenIntrospectionUri, true
}

// HasTokenIntrospectionUri returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionUri() bool {
	if o != nil && o.TokenIntrospectionUri != nil {
		return true
	}

	return false
}

// SetTokenIntrospectionUri gets a reference to the given string and assigns it to the TokenIntrospectionUri field.
func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionUri(v string) {
	o.TokenIntrospectionUri = &v
}

// GetTokenIntrospectionUsername returns the TokenIntrospectionUsername field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionUsername() string {
	if o == nil || o.TokenIntrospectionUsername == nil {
		var ret string
		return ret
	}
	return *o.TokenIntrospectionUsername
}

// GetTokenIntrospectionUsernameOk returns a tuple with the TokenIntrospectionUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionUsernameOk() (*string, bool) {
	if o == nil || o.TokenIntrospectionUsername == nil {
		return nil, false
	}
	return o.TokenIntrospectionUsername, true
}

// HasTokenIntrospectionUsername returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionUsername() bool {
	if o != nil && o.TokenIntrospectionUsername != nil {
		return true
	}

	return false
}

// SetTokenIntrospectionUsername gets a reference to the given string and assigns it to the TokenIntrospectionUsername field.
func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionUsername(v string) {
	o.TokenIntrospectionUsername = &v
}

// GetUsernameClaimName returns the UsernameClaimName field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetUsernameClaimName() string {
	if o == nil || o.UsernameClaimName == nil {
		var ret string
		return ret
	}
	return *o.UsernameClaimName
}

// GetUsernameClaimNameOk returns a tuple with the UsernameClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetUsernameClaimNameOk() (*string, bool) {
	if o == nil || o.UsernameClaimName == nil {
		return nil, false
	}
	return o.UsernameClaimName, true
}

// HasUsernameClaimName returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasUsernameClaimName() bool {
	if o != nil && o.UsernameClaimName != nil {
		return true
	}

	return false
}

// SetUsernameClaimName gets a reference to the given string and assigns it to the UsernameClaimName field.
func (o *MsgVpnAuthenticationOauthProvider) SetUsernameClaimName(v string) {
	o.UsernameClaimName = &v
}

// GetUsernameClaimSource returns the UsernameClaimSource field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetUsernameClaimSource() string {
	if o == nil || o.UsernameClaimSource == nil {
		var ret string
		return ret
	}
	return *o.UsernameClaimSource
}

// GetUsernameClaimSourceOk returns a tuple with the UsernameClaimSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetUsernameClaimSourceOk() (*string, bool) {
	if o == nil || o.UsernameClaimSource == nil {
		return nil, false
	}
	return o.UsernameClaimSource, true
}

// HasUsernameClaimSource returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasUsernameClaimSource() bool {
	if o != nil && o.UsernameClaimSource != nil {
		return true
	}

	return false
}

// SetUsernameClaimSource gets a reference to the given string and assigns it to the UsernameClaimSource field.
func (o *MsgVpnAuthenticationOauthProvider) SetUsernameClaimSource(v string) {
	o.UsernameClaimSource = &v
}

// GetUsernameValidateEnabled returns the UsernameValidateEnabled field value if set, zero value otherwise.
func (o *MsgVpnAuthenticationOauthProvider) GetUsernameValidateEnabled() bool {
	if o == nil || o.UsernameValidateEnabled == nil {
		var ret bool
		return ret
	}
	return *o.UsernameValidateEnabled
}

// GetUsernameValidateEnabledOk returns a tuple with the UsernameValidateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnAuthenticationOauthProvider) GetUsernameValidateEnabledOk() (*bool, bool) {
	if o == nil || o.UsernameValidateEnabled == nil {
		return nil, false
	}
	return o.UsernameValidateEnabled, true
}

// HasUsernameValidateEnabled returns a boolean if a field has been set.
func (o *MsgVpnAuthenticationOauthProvider) HasUsernameValidateEnabled() bool {
	if o != nil && o.UsernameValidateEnabled != nil {
		return true
	}

	return false
}

// SetUsernameValidateEnabled gets a reference to the given bool and assigns it to the UsernameValidateEnabled field.
func (o *MsgVpnAuthenticationOauthProvider) SetUsernameValidateEnabled(v bool) {
	o.UsernameValidateEnabled = &v
}

func (o MsgVpnAuthenticationOauthProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AudienceClaimName != nil {
		toSerialize["audienceClaimName"] = o.AudienceClaimName
	}
	if o.AudienceClaimSource != nil {
		toSerialize["audienceClaimSource"] = o.AudienceClaimSource
	}
	if o.AudienceClaimValue != nil {
		toSerialize["audienceClaimValue"] = o.AudienceClaimValue
	}
	if o.AudienceValidationEnabled != nil {
		toSerialize["audienceValidationEnabled"] = o.AudienceValidationEnabled
	}
	if o.AuthenticationSuccessCount != nil {
		toSerialize["authenticationSuccessCount"] = o.AuthenticationSuccessCount
	}
	if o.AuthorizationGroupClaimName != nil {
		toSerialize["authorizationGroupClaimName"] = o.AuthorizationGroupClaimName
	}
	if o.AuthorizationGroupClaimSource != nil {
		toSerialize["authorizationGroupClaimSource"] = o.AuthorizationGroupClaimSource
	}
	if o.AuthorizationGroupEnabled != nil {
		toSerialize["authorizationGroupEnabled"] = o.AuthorizationGroupEnabled
	}
	if o.DisconnectOnTokenExpirationEnabled != nil {
		toSerialize["disconnectOnTokenExpirationEnabled"] = o.DisconnectOnTokenExpirationEnabled
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.JwksLastRefreshFailureReason != nil {
		toSerialize["jwksLastRefreshFailureReason"] = o.JwksLastRefreshFailureReason
	}
	if o.JwksLastRefreshFailureTime != nil {
		toSerialize["jwksLastRefreshFailureTime"] = o.JwksLastRefreshFailureTime
	}
	if o.JwksLastRefreshTime != nil {
		toSerialize["jwksLastRefreshTime"] = o.JwksLastRefreshTime
	}
	if o.JwksNextScheduledRefreshTime != nil {
		toSerialize["jwksNextScheduledRefreshTime"] = o.JwksNextScheduledRefreshTime
	}
	if o.JwksRefreshFailureCount != nil {
		toSerialize["jwksRefreshFailureCount"] = o.JwksRefreshFailureCount
	}
	if o.JwksRefreshInterval != nil {
		toSerialize["jwksRefreshInterval"] = o.JwksRefreshInterval
	}
	if o.JwksUri != nil {
		toSerialize["jwksUri"] = o.JwksUri
	}
	if o.LoginFailureIncorrectAudienceValueCount != nil {
		toSerialize["loginFailureIncorrectAudienceValueCount"] = o.LoginFailureIncorrectAudienceValueCount
	}
	if o.LoginFailureInvalidAudienceValueCount != nil {
		toSerialize["loginFailureInvalidAudienceValueCount"] = o.LoginFailureInvalidAudienceValueCount
	}
	if o.LoginFailureInvalidAuthorizationGroupValueCount != nil {
		toSerialize["loginFailureInvalidAuthorizationGroupValueCount"] = o.LoginFailureInvalidAuthorizationGroupValueCount
	}
	if o.LoginFailureInvalidJwtSignatureCount != nil {
		toSerialize["loginFailureInvalidJwtSignatureCount"] = o.LoginFailureInvalidJwtSignatureCount
	}
	if o.LoginFailureInvalidUsernameValueCount != nil {
		toSerialize["loginFailureInvalidUsernameValueCount"] = o.LoginFailureInvalidUsernameValueCount
	}
	if o.LoginFailureMismatchedUsernameCount != nil {
		toSerialize["loginFailureMismatchedUsernameCount"] = o.LoginFailureMismatchedUsernameCount
	}
	if o.LoginFailureMissingAudienceCount != nil {
		toSerialize["loginFailureMissingAudienceCount"] = o.LoginFailureMissingAudienceCount
	}
	if o.LoginFailureMissingJwkCount != nil {
		toSerialize["loginFailureMissingJwkCount"] = o.LoginFailureMissingJwkCount
	}
	if o.LoginFailureMissingOrInvalidTokenCount != nil {
		toSerialize["loginFailureMissingOrInvalidTokenCount"] = o.LoginFailureMissingOrInvalidTokenCount
	}
	if o.LoginFailureMissingUsernameCount != nil {
		toSerialize["loginFailureMissingUsernameCount"] = o.LoginFailureMissingUsernameCount
	}
	if o.LoginFailureTokenExpiredCount != nil {
		toSerialize["loginFailureTokenExpiredCount"] = o.LoginFailureTokenExpiredCount
	}
	if o.LoginFailureTokenIntrospectionErroredCount != nil {
		toSerialize["loginFailureTokenIntrospectionErroredCount"] = o.LoginFailureTokenIntrospectionErroredCount
	}
	if o.LoginFailureTokenIntrospectionFailureCount != nil {
		toSerialize["loginFailureTokenIntrospectionFailureCount"] = o.LoginFailureTokenIntrospectionFailureCount
	}
	if o.LoginFailureTokenIntrospectionHttpsErrorCount != nil {
		toSerialize["loginFailureTokenIntrospectionHttpsErrorCount"] = o.LoginFailureTokenIntrospectionHttpsErrorCount
	}
	if o.LoginFailureTokenIntrospectionInvalidCount != nil {
		toSerialize["loginFailureTokenIntrospectionInvalidCount"] = o.LoginFailureTokenIntrospectionInvalidCount
	}
	if o.LoginFailureTokenIntrospectionTimeoutCount != nil {
		toSerialize["loginFailureTokenIntrospectionTimeoutCount"] = o.LoginFailureTokenIntrospectionTimeoutCount
	}
	if o.LoginFailureTokenNotValidYetCount != nil {
		toSerialize["loginFailureTokenNotValidYetCount"] = o.LoginFailureTokenNotValidYetCount
	}
	if o.LoginFailureUnsupportedAlgCount != nil {
		toSerialize["loginFailureUnsupportedAlgCount"] = o.LoginFailureUnsupportedAlgCount
	}
	if o.MissingAuthorizationGroupCount != nil {
		toSerialize["missingAuthorizationGroupCount"] = o.MissingAuthorizationGroupCount
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.OauthProviderName != nil {
		toSerialize["oauthProviderName"] = o.OauthProviderName
	}
	if o.TokenIgnoreTimeLimitsEnabled != nil {
		toSerialize["tokenIgnoreTimeLimitsEnabled"] = o.TokenIgnoreTimeLimitsEnabled
	}
	if o.TokenIntrospectionAverageTime != nil {
		toSerialize["tokenIntrospectionAverageTime"] = o.TokenIntrospectionAverageTime
	}
	if o.TokenIntrospectionLastFailureReason != nil {
		toSerialize["tokenIntrospectionLastFailureReason"] = o.TokenIntrospectionLastFailureReason
	}
	if o.TokenIntrospectionLastFailureTime != nil {
		toSerialize["tokenIntrospectionLastFailureTime"] = o.TokenIntrospectionLastFailureTime
	}
	if o.TokenIntrospectionParameterName != nil {
		toSerialize["tokenIntrospectionParameterName"] = o.TokenIntrospectionParameterName
	}
	if o.TokenIntrospectionSuccessCount != nil {
		toSerialize["tokenIntrospectionSuccessCount"] = o.TokenIntrospectionSuccessCount
	}
	if o.TokenIntrospectionTimeout != nil {
		toSerialize["tokenIntrospectionTimeout"] = o.TokenIntrospectionTimeout
	}
	if o.TokenIntrospectionUri != nil {
		toSerialize["tokenIntrospectionUri"] = o.TokenIntrospectionUri
	}
	if o.TokenIntrospectionUsername != nil {
		toSerialize["tokenIntrospectionUsername"] = o.TokenIntrospectionUsername
	}
	if o.UsernameClaimName != nil {
		toSerialize["usernameClaimName"] = o.UsernameClaimName
	}
	if o.UsernameClaimSource != nil {
		toSerialize["usernameClaimSource"] = o.UsernameClaimSource
	}
	if o.UsernameValidateEnabled != nil {
		toSerialize["usernameValidateEnabled"] = o.UsernameValidateEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnAuthenticationOauthProvider struct {
	value *MsgVpnAuthenticationOauthProvider
	isSet bool
}

func (v NullableMsgVpnAuthenticationOauthProvider) Get() *MsgVpnAuthenticationOauthProvider {
	return v.value
}

func (v *NullableMsgVpnAuthenticationOauthProvider) Set(val *MsgVpnAuthenticationOauthProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnAuthenticationOauthProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnAuthenticationOauthProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnAuthenticationOauthProvider(val *MsgVpnAuthenticationOauthProvider) *NullableMsgVpnAuthenticationOauthProvider {
	return &NullableMsgVpnAuthenticationOauthProvider{value: val, isSet: true}
}

func (v NullableMsgVpnAuthenticationOauthProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnAuthenticationOauthProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
