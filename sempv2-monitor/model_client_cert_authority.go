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

// ClientCertAuthority struct for ClientCertAuthority
type ClientCertAuthority struct {
	// The name of the Certificate Authority.
	CertAuthorityName *string `json:"certAuthorityName,omitempty"`
	// The PEM formatted content for the trusted root certificate of a client Certificate Authority.
	CertContent *string `json:"certContent,omitempty"`
	// The scheduled CRL refresh day(s), specified as \"daily\" or a comma-separated list of days. Days must be specified as \"Sun\", \"Mon\", \"Tue\", \"Wed\", \"Thu\", \"Fri\", or \"Sat\", with no spaces, and in sorted order from Sunday to Saturday.
	CrlDayList *string `json:"crlDayList,omitempty"`
	// The timestamp of the last successful CRL download. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	CrlLastDownloadTime *int32 `json:"crlLastDownloadTime,omitempty"`
	// The reason for the last CRL failure.
	CrlLastFailureReason *string `json:"crlLastFailureReason,omitempty"`
	// The timestamp of the last CRL failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	CrlLastFailureTime *int32 `json:"crlLastFailureTime,omitempty"`
	// The scheduled time of the next CRL download. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	CrlNextDownloadTime *int32 `json:"crlNextDownloadTime,omitempty"`
	// The scheduled CRL refresh time(s), specified as \"hourly\" or a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times must be in sorted order from 0:00 to 23:59.
	CrlTimeList *string `json:"crlTimeList,omitempty"`
	// Indicates whether CRL revocation checking is operationally up.
	CrlUp *bool `json:"crlUp,omitempty"`
	// The URL for the CRL source. This is a required attribute for CRL to be operational and the URL must be complete with http:// included.
	CrlUrl *string `json:"crlUrl,omitempty"`
	// The reason for the last OCSP failure.
	OcspLastFailureReason *string `json:"ocspLastFailureReason,omitempty"`
	// The timestamp of the last OCSP failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	OcspLastFailureTime *int32 `json:"ocspLastFailureTime,omitempty"`
	// The URL involved in the last OCSP failure.
	OcspLastFailureUrl *string `json:"ocspLastFailureUrl,omitempty"`
	// Indicates whether a non-responder certificate is allowed to sign an OCSP response. Typically used with an OCSP override URL in cases where a single certificate is used to sign client certificates and OCSP responses.
	OcspNonResponderCertEnabled *bool `json:"ocspNonResponderCertEnabled,omitempty"`
	// The OCSP responder URL to use for overriding the one supplied in the client certificate. The URL must be complete with http:// included.
	OcspOverrideUrl *string `json:"ocspOverrideUrl,omitempty"`
	// The timeout in seconds to receive a response from the OCSP responder after sending a request or making the initial connection attempt.
	OcspTimeout *int64 `json:"ocspTimeout,omitempty"`
	// Indicates whether Certificate Authority revocation checking is enabled.
	RevocationCheckEnabled *bool `json:"revocationCheckEnabled,omitempty"`
}

// NewClientCertAuthority instantiates a new ClientCertAuthority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientCertAuthority() *ClientCertAuthority {
	this := ClientCertAuthority{}
	return &this
}

// NewClientCertAuthorityWithDefaults instantiates a new ClientCertAuthority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientCertAuthorityWithDefaults() *ClientCertAuthority {
	this := ClientCertAuthority{}
	return &this
}

// GetCertAuthorityName returns the CertAuthorityName field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetCertAuthorityName() string {
	if o == nil || o.CertAuthorityName == nil {
		var ret string
		return ret
	}
	return *o.CertAuthorityName
}

// GetCertAuthorityNameOk returns a tuple with the CertAuthorityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetCertAuthorityNameOk() (*string, bool) {
	if o == nil || o.CertAuthorityName == nil {
		return nil, false
	}
	return o.CertAuthorityName, true
}

// HasCertAuthorityName returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasCertAuthorityName() bool {
	if o != nil && o.CertAuthorityName != nil {
		return true
	}

	return false
}

// SetCertAuthorityName gets a reference to the given string and assigns it to the CertAuthorityName field.
func (o *ClientCertAuthority) SetCertAuthorityName(v string) {
	o.CertAuthorityName = &v
}

// GetCertContent returns the CertContent field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetCertContent() string {
	if o == nil || o.CertContent == nil {
		var ret string
		return ret
	}
	return *o.CertContent
}

// GetCertContentOk returns a tuple with the CertContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetCertContentOk() (*string, bool) {
	if o == nil || o.CertContent == nil {
		return nil, false
	}
	return o.CertContent, true
}

// HasCertContent returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasCertContent() bool {
	if o != nil && o.CertContent != nil {
		return true
	}

	return false
}

// SetCertContent gets a reference to the given string and assigns it to the CertContent field.
func (o *ClientCertAuthority) SetCertContent(v string) {
	o.CertContent = &v
}

// GetCrlDayList returns the CrlDayList field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetCrlDayList() string {
	if o == nil || o.CrlDayList == nil {
		var ret string
		return ret
	}
	return *o.CrlDayList
}

// GetCrlDayListOk returns a tuple with the CrlDayList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetCrlDayListOk() (*string, bool) {
	if o == nil || o.CrlDayList == nil {
		return nil, false
	}
	return o.CrlDayList, true
}

// HasCrlDayList returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasCrlDayList() bool {
	if o != nil && o.CrlDayList != nil {
		return true
	}

	return false
}

// SetCrlDayList gets a reference to the given string and assigns it to the CrlDayList field.
func (o *ClientCertAuthority) SetCrlDayList(v string) {
	o.CrlDayList = &v
}

// GetCrlLastDownloadTime returns the CrlLastDownloadTime field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetCrlLastDownloadTime() int32 {
	if o == nil || o.CrlLastDownloadTime == nil {
		var ret int32
		return ret
	}
	return *o.CrlLastDownloadTime
}

// GetCrlLastDownloadTimeOk returns a tuple with the CrlLastDownloadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetCrlLastDownloadTimeOk() (*int32, bool) {
	if o == nil || o.CrlLastDownloadTime == nil {
		return nil, false
	}
	return o.CrlLastDownloadTime, true
}

// HasCrlLastDownloadTime returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasCrlLastDownloadTime() bool {
	if o != nil && o.CrlLastDownloadTime != nil {
		return true
	}

	return false
}

// SetCrlLastDownloadTime gets a reference to the given int32 and assigns it to the CrlLastDownloadTime field.
func (o *ClientCertAuthority) SetCrlLastDownloadTime(v int32) {
	o.CrlLastDownloadTime = &v
}

// GetCrlLastFailureReason returns the CrlLastFailureReason field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetCrlLastFailureReason() string {
	if o == nil || o.CrlLastFailureReason == nil {
		var ret string
		return ret
	}
	return *o.CrlLastFailureReason
}

// GetCrlLastFailureReasonOk returns a tuple with the CrlLastFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetCrlLastFailureReasonOk() (*string, bool) {
	if o == nil || o.CrlLastFailureReason == nil {
		return nil, false
	}
	return o.CrlLastFailureReason, true
}

// HasCrlLastFailureReason returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasCrlLastFailureReason() bool {
	if o != nil && o.CrlLastFailureReason != nil {
		return true
	}

	return false
}

// SetCrlLastFailureReason gets a reference to the given string and assigns it to the CrlLastFailureReason field.
func (o *ClientCertAuthority) SetCrlLastFailureReason(v string) {
	o.CrlLastFailureReason = &v
}

// GetCrlLastFailureTime returns the CrlLastFailureTime field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetCrlLastFailureTime() int32 {
	if o == nil || o.CrlLastFailureTime == nil {
		var ret int32
		return ret
	}
	return *o.CrlLastFailureTime
}

// GetCrlLastFailureTimeOk returns a tuple with the CrlLastFailureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetCrlLastFailureTimeOk() (*int32, bool) {
	if o == nil || o.CrlLastFailureTime == nil {
		return nil, false
	}
	return o.CrlLastFailureTime, true
}

// HasCrlLastFailureTime returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasCrlLastFailureTime() bool {
	if o != nil && o.CrlLastFailureTime != nil {
		return true
	}

	return false
}

// SetCrlLastFailureTime gets a reference to the given int32 and assigns it to the CrlLastFailureTime field.
func (o *ClientCertAuthority) SetCrlLastFailureTime(v int32) {
	o.CrlLastFailureTime = &v
}

// GetCrlNextDownloadTime returns the CrlNextDownloadTime field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetCrlNextDownloadTime() int32 {
	if o == nil || o.CrlNextDownloadTime == nil {
		var ret int32
		return ret
	}
	return *o.CrlNextDownloadTime
}

// GetCrlNextDownloadTimeOk returns a tuple with the CrlNextDownloadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetCrlNextDownloadTimeOk() (*int32, bool) {
	if o == nil || o.CrlNextDownloadTime == nil {
		return nil, false
	}
	return o.CrlNextDownloadTime, true
}

// HasCrlNextDownloadTime returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasCrlNextDownloadTime() bool {
	if o != nil && o.CrlNextDownloadTime != nil {
		return true
	}

	return false
}

// SetCrlNextDownloadTime gets a reference to the given int32 and assigns it to the CrlNextDownloadTime field.
func (o *ClientCertAuthority) SetCrlNextDownloadTime(v int32) {
	o.CrlNextDownloadTime = &v
}

// GetCrlTimeList returns the CrlTimeList field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetCrlTimeList() string {
	if o == nil || o.CrlTimeList == nil {
		var ret string
		return ret
	}
	return *o.CrlTimeList
}

// GetCrlTimeListOk returns a tuple with the CrlTimeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetCrlTimeListOk() (*string, bool) {
	if o == nil || o.CrlTimeList == nil {
		return nil, false
	}
	return o.CrlTimeList, true
}

// HasCrlTimeList returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasCrlTimeList() bool {
	if o != nil && o.CrlTimeList != nil {
		return true
	}

	return false
}

// SetCrlTimeList gets a reference to the given string and assigns it to the CrlTimeList field.
func (o *ClientCertAuthority) SetCrlTimeList(v string) {
	o.CrlTimeList = &v
}

// GetCrlUp returns the CrlUp field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetCrlUp() bool {
	if o == nil || o.CrlUp == nil {
		var ret bool
		return ret
	}
	return *o.CrlUp
}

// GetCrlUpOk returns a tuple with the CrlUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetCrlUpOk() (*bool, bool) {
	if o == nil || o.CrlUp == nil {
		return nil, false
	}
	return o.CrlUp, true
}

// HasCrlUp returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasCrlUp() bool {
	if o != nil && o.CrlUp != nil {
		return true
	}

	return false
}

// SetCrlUp gets a reference to the given bool and assigns it to the CrlUp field.
func (o *ClientCertAuthority) SetCrlUp(v bool) {
	o.CrlUp = &v
}

// GetCrlUrl returns the CrlUrl field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetCrlUrl() string {
	if o == nil || o.CrlUrl == nil {
		var ret string
		return ret
	}
	return *o.CrlUrl
}

// GetCrlUrlOk returns a tuple with the CrlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetCrlUrlOk() (*string, bool) {
	if o == nil || o.CrlUrl == nil {
		return nil, false
	}
	return o.CrlUrl, true
}

// HasCrlUrl returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasCrlUrl() bool {
	if o != nil && o.CrlUrl != nil {
		return true
	}

	return false
}

// SetCrlUrl gets a reference to the given string and assigns it to the CrlUrl field.
func (o *ClientCertAuthority) SetCrlUrl(v string) {
	o.CrlUrl = &v
}

// GetOcspLastFailureReason returns the OcspLastFailureReason field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetOcspLastFailureReason() string {
	if o == nil || o.OcspLastFailureReason == nil {
		var ret string
		return ret
	}
	return *o.OcspLastFailureReason
}

// GetOcspLastFailureReasonOk returns a tuple with the OcspLastFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetOcspLastFailureReasonOk() (*string, bool) {
	if o == nil || o.OcspLastFailureReason == nil {
		return nil, false
	}
	return o.OcspLastFailureReason, true
}

// HasOcspLastFailureReason returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasOcspLastFailureReason() bool {
	if o != nil && o.OcspLastFailureReason != nil {
		return true
	}

	return false
}

// SetOcspLastFailureReason gets a reference to the given string and assigns it to the OcspLastFailureReason field.
func (o *ClientCertAuthority) SetOcspLastFailureReason(v string) {
	o.OcspLastFailureReason = &v
}

// GetOcspLastFailureTime returns the OcspLastFailureTime field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetOcspLastFailureTime() int32 {
	if o == nil || o.OcspLastFailureTime == nil {
		var ret int32
		return ret
	}
	return *o.OcspLastFailureTime
}

// GetOcspLastFailureTimeOk returns a tuple with the OcspLastFailureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetOcspLastFailureTimeOk() (*int32, bool) {
	if o == nil || o.OcspLastFailureTime == nil {
		return nil, false
	}
	return o.OcspLastFailureTime, true
}

// HasOcspLastFailureTime returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasOcspLastFailureTime() bool {
	if o != nil && o.OcspLastFailureTime != nil {
		return true
	}

	return false
}

// SetOcspLastFailureTime gets a reference to the given int32 and assigns it to the OcspLastFailureTime field.
func (o *ClientCertAuthority) SetOcspLastFailureTime(v int32) {
	o.OcspLastFailureTime = &v
}

// GetOcspLastFailureUrl returns the OcspLastFailureUrl field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetOcspLastFailureUrl() string {
	if o == nil || o.OcspLastFailureUrl == nil {
		var ret string
		return ret
	}
	return *o.OcspLastFailureUrl
}

// GetOcspLastFailureUrlOk returns a tuple with the OcspLastFailureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetOcspLastFailureUrlOk() (*string, bool) {
	if o == nil || o.OcspLastFailureUrl == nil {
		return nil, false
	}
	return o.OcspLastFailureUrl, true
}

// HasOcspLastFailureUrl returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasOcspLastFailureUrl() bool {
	if o != nil && o.OcspLastFailureUrl != nil {
		return true
	}

	return false
}

// SetOcspLastFailureUrl gets a reference to the given string and assigns it to the OcspLastFailureUrl field.
func (o *ClientCertAuthority) SetOcspLastFailureUrl(v string) {
	o.OcspLastFailureUrl = &v
}

// GetOcspNonResponderCertEnabled returns the OcspNonResponderCertEnabled field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetOcspNonResponderCertEnabled() bool {
	if o == nil || o.OcspNonResponderCertEnabled == nil {
		var ret bool
		return ret
	}
	return *o.OcspNonResponderCertEnabled
}

// GetOcspNonResponderCertEnabledOk returns a tuple with the OcspNonResponderCertEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetOcspNonResponderCertEnabledOk() (*bool, bool) {
	if o == nil || o.OcspNonResponderCertEnabled == nil {
		return nil, false
	}
	return o.OcspNonResponderCertEnabled, true
}

// HasOcspNonResponderCertEnabled returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasOcspNonResponderCertEnabled() bool {
	if o != nil && o.OcspNonResponderCertEnabled != nil {
		return true
	}

	return false
}

// SetOcspNonResponderCertEnabled gets a reference to the given bool and assigns it to the OcspNonResponderCertEnabled field.
func (o *ClientCertAuthority) SetOcspNonResponderCertEnabled(v bool) {
	o.OcspNonResponderCertEnabled = &v
}

// GetOcspOverrideUrl returns the OcspOverrideUrl field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetOcspOverrideUrl() string {
	if o == nil || o.OcspOverrideUrl == nil {
		var ret string
		return ret
	}
	return *o.OcspOverrideUrl
}

// GetOcspOverrideUrlOk returns a tuple with the OcspOverrideUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetOcspOverrideUrlOk() (*string, bool) {
	if o == nil || o.OcspOverrideUrl == nil {
		return nil, false
	}
	return o.OcspOverrideUrl, true
}

// HasOcspOverrideUrl returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasOcspOverrideUrl() bool {
	if o != nil && o.OcspOverrideUrl != nil {
		return true
	}

	return false
}

// SetOcspOverrideUrl gets a reference to the given string and assigns it to the OcspOverrideUrl field.
func (o *ClientCertAuthority) SetOcspOverrideUrl(v string) {
	o.OcspOverrideUrl = &v
}

// GetOcspTimeout returns the OcspTimeout field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetOcspTimeout() int64 {
	if o == nil || o.OcspTimeout == nil {
		var ret int64
		return ret
	}
	return *o.OcspTimeout
}

// GetOcspTimeoutOk returns a tuple with the OcspTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetOcspTimeoutOk() (*int64, bool) {
	if o == nil || o.OcspTimeout == nil {
		return nil, false
	}
	return o.OcspTimeout, true
}

// HasOcspTimeout returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasOcspTimeout() bool {
	if o != nil && o.OcspTimeout != nil {
		return true
	}

	return false
}

// SetOcspTimeout gets a reference to the given int64 and assigns it to the OcspTimeout field.
func (o *ClientCertAuthority) SetOcspTimeout(v int64) {
	o.OcspTimeout = &v
}

// GetRevocationCheckEnabled returns the RevocationCheckEnabled field value if set, zero value otherwise.
func (o *ClientCertAuthority) GetRevocationCheckEnabled() bool {
	if o == nil || o.RevocationCheckEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RevocationCheckEnabled
}

// GetRevocationCheckEnabledOk returns a tuple with the RevocationCheckEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCertAuthority) GetRevocationCheckEnabledOk() (*bool, bool) {
	if o == nil || o.RevocationCheckEnabled == nil {
		return nil, false
	}
	return o.RevocationCheckEnabled, true
}

// HasRevocationCheckEnabled returns a boolean if a field has been set.
func (o *ClientCertAuthority) HasRevocationCheckEnabled() bool {
	if o != nil && o.RevocationCheckEnabled != nil {
		return true
	}

	return false
}

// SetRevocationCheckEnabled gets a reference to the given bool and assigns it to the RevocationCheckEnabled field.
func (o *ClientCertAuthority) SetRevocationCheckEnabled(v bool) {
	o.RevocationCheckEnabled = &v
}

func (o ClientCertAuthority) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertAuthorityName != nil {
		toSerialize["certAuthorityName"] = o.CertAuthorityName
	}
	if o.CertContent != nil {
		toSerialize["certContent"] = o.CertContent
	}
	if o.CrlDayList != nil {
		toSerialize["crlDayList"] = o.CrlDayList
	}
	if o.CrlLastDownloadTime != nil {
		toSerialize["crlLastDownloadTime"] = o.CrlLastDownloadTime
	}
	if o.CrlLastFailureReason != nil {
		toSerialize["crlLastFailureReason"] = o.CrlLastFailureReason
	}
	if o.CrlLastFailureTime != nil {
		toSerialize["crlLastFailureTime"] = o.CrlLastFailureTime
	}
	if o.CrlNextDownloadTime != nil {
		toSerialize["crlNextDownloadTime"] = o.CrlNextDownloadTime
	}
	if o.CrlTimeList != nil {
		toSerialize["crlTimeList"] = o.CrlTimeList
	}
	if o.CrlUp != nil {
		toSerialize["crlUp"] = o.CrlUp
	}
	if o.CrlUrl != nil {
		toSerialize["crlUrl"] = o.CrlUrl
	}
	if o.OcspLastFailureReason != nil {
		toSerialize["ocspLastFailureReason"] = o.OcspLastFailureReason
	}
	if o.OcspLastFailureTime != nil {
		toSerialize["ocspLastFailureTime"] = o.OcspLastFailureTime
	}
	if o.OcspLastFailureUrl != nil {
		toSerialize["ocspLastFailureUrl"] = o.OcspLastFailureUrl
	}
	if o.OcspNonResponderCertEnabled != nil {
		toSerialize["ocspNonResponderCertEnabled"] = o.OcspNonResponderCertEnabled
	}
	if o.OcspOverrideUrl != nil {
		toSerialize["ocspOverrideUrl"] = o.OcspOverrideUrl
	}
	if o.OcspTimeout != nil {
		toSerialize["ocspTimeout"] = o.OcspTimeout
	}
	if o.RevocationCheckEnabled != nil {
		toSerialize["revocationCheckEnabled"] = o.RevocationCheckEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableClientCertAuthority struct {
	value *ClientCertAuthority
	isSet bool
}

func (v NullableClientCertAuthority) Get() *ClientCertAuthority {
	return v.value
}

func (v *NullableClientCertAuthority) Set(val *ClientCertAuthority) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCertAuthority) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCertAuthority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCertAuthority(val *ClientCertAuthority) *NullableClientCertAuthority {
	return &NullableClientCertAuthority{value: val, isSet: true}
}

func (v NullableClientCertAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCertAuthority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
