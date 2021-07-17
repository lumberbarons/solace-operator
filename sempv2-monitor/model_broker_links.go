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

// BrokerLinks struct for BrokerLinks
type BrokerLinks struct {
	// The URI of this Broker's About object.
	AboutUri *string `json:"aboutUri,omitempty"`
	// The URI of this Broker's collection of Certificate Authority objects. Deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.
	CertAuthoritiesUri *string `json:"certAuthoritiesUri,omitempty"`
	// The URI of this Broker's collection of Client Certificate Authority objects. Available since 2.19.
	ClientCertAuthoritiesUri *string `json:"clientCertAuthoritiesUri,omitempty"`
	// The URI of this Broker's collection of Cluster objects. Available since 2.11.
	DmrClustersUri *string `json:"dmrClustersUri,omitempty"`
	// The URI of this Broker's collection of Domain Certificate Authority objects. Available since 2.19.
	DomainCertAuthoritiesUri *string `json:"domainCertAuthoritiesUri,omitempty"`
	// The URI of this Broker's collection of Message VPN objects. Available since 2.11.
	MsgVpnsUri *string `json:"msgVpnsUri,omitempty"`
	// The URI of this Broker's collection of Session objects. Available since 2.21.
	SessionsUri *string `json:"sessionsUri,omitempty"`
	// The URI of this Broker's collection of Standard Domain Certificate Authority objects. Available since 2.19.
	StandardDomainCertAuthoritiesUri *string `json:"standardDomainCertAuthoritiesUri,omitempty"`
	// The URI of this Broker object.
	Uri *string `json:"uri,omitempty"`
	// The URI of this Broker's collection of Virtual Hostname objects. Available since 2.17.
	VirtualHostnamesUri *string `json:"virtualHostnamesUri,omitempty"`
}

// NewBrokerLinks instantiates a new BrokerLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerLinks() *BrokerLinks {
	this := BrokerLinks{}
	return &this
}

// NewBrokerLinksWithDefaults instantiates a new BrokerLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerLinksWithDefaults() *BrokerLinks {
	this := BrokerLinks{}
	return &this
}

// GetAboutUri returns the AboutUri field value if set, zero value otherwise.
func (o *BrokerLinks) GetAboutUri() string {
	if o == nil || o.AboutUri == nil {
		var ret string
		return ret
	}
	return *o.AboutUri
}

// GetAboutUriOk returns a tuple with the AboutUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerLinks) GetAboutUriOk() (*string, bool) {
	if o == nil || o.AboutUri == nil {
		return nil, false
	}
	return o.AboutUri, true
}

// HasAboutUri returns a boolean if a field has been set.
func (o *BrokerLinks) HasAboutUri() bool {
	if o != nil && o.AboutUri != nil {
		return true
	}

	return false
}

// SetAboutUri gets a reference to the given string and assigns it to the AboutUri field.
func (o *BrokerLinks) SetAboutUri(v string) {
	o.AboutUri = &v
}

// GetCertAuthoritiesUri returns the CertAuthoritiesUri field value if set, zero value otherwise.
func (o *BrokerLinks) GetCertAuthoritiesUri() string {
	if o == nil || o.CertAuthoritiesUri == nil {
		var ret string
		return ret
	}
	return *o.CertAuthoritiesUri
}

// GetCertAuthoritiesUriOk returns a tuple with the CertAuthoritiesUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerLinks) GetCertAuthoritiesUriOk() (*string, bool) {
	if o == nil || o.CertAuthoritiesUri == nil {
		return nil, false
	}
	return o.CertAuthoritiesUri, true
}

// HasCertAuthoritiesUri returns a boolean if a field has been set.
func (o *BrokerLinks) HasCertAuthoritiesUri() bool {
	if o != nil && o.CertAuthoritiesUri != nil {
		return true
	}

	return false
}

// SetCertAuthoritiesUri gets a reference to the given string and assigns it to the CertAuthoritiesUri field.
func (o *BrokerLinks) SetCertAuthoritiesUri(v string) {
	o.CertAuthoritiesUri = &v
}

// GetClientCertAuthoritiesUri returns the ClientCertAuthoritiesUri field value if set, zero value otherwise.
func (o *BrokerLinks) GetClientCertAuthoritiesUri() string {
	if o == nil || o.ClientCertAuthoritiesUri == nil {
		var ret string
		return ret
	}
	return *o.ClientCertAuthoritiesUri
}

// GetClientCertAuthoritiesUriOk returns a tuple with the ClientCertAuthoritiesUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerLinks) GetClientCertAuthoritiesUriOk() (*string, bool) {
	if o == nil || o.ClientCertAuthoritiesUri == nil {
		return nil, false
	}
	return o.ClientCertAuthoritiesUri, true
}

// HasClientCertAuthoritiesUri returns a boolean if a field has been set.
func (o *BrokerLinks) HasClientCertAuthoritiesUri() bool {
	if o != nil && o.ClientCertAuthoritiesUri != nil {
		return true
	}

	return false
}

// SetClientCertAuthoritiesUri gets a reference to the given string and assigns it to the ClientCertAuthoritiesUri field.
func (o *BrokerLinks) SetClientCertAuthoritiesUri(v string) {
	o.ClientCertAuthoritiesUri = &v
}

// GetDmrClustersUri returns the DmrClustersUri field value if set, zero value otherwise.
func (o *BrokerLinks) GetDmrClustersUri() string {
	if o == nil || o.DmrClustersUri == nil {
		var ret string
		return ret
	}
	return *o.DmrClustersUri
}

// GetDmrClustersUriOk returns a tuple with the DmrClustersUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerLinks) GetDmrClustersUriOk() (*string, bool) {
	if o == nil || o.DmrClustersUri == nil {
		return nil, false
	}
	return o.DmrClustersUri, true
}

// HasDmrClustersUri returns a boolean if a field has been set.
func (o *BrokerLinks) HasDmrClustersUri() bool {
	if o != nil && o.DmrClustersUri != nil {
		return true
	}

	return false
}

// SetDmrClustersUri gets a reference to the given string and assigns it to the DmrClustersUri field.
func (o *BrokerLinks) SetDmrClustersUri(v string) {
	o.DmrClustersUri = &v
}

// GetDomainCertAuthoritiesUri returns the DomainCertAuthoritiesUri field value if set, zero value otherwise.
func (o *BrokerLinks) GetDomainCertAuthoritiesUri() string {
	if o == nil || o.DomainCertAuthoritiesUri == nil {
		var ret string
		return ret
	}
	return *o.DomainCertAuthoritiesUri
}

// GetDomainCertAuthoritiesUriOk returns a tuple with the DomainCertAuthoritiesUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerLinks) GetDomainCertAuthoritiesUriOk() (*string, bool) {
	if o == nil || o.DomainCertAuthoritiesUri == nil {
		return nil, false
	}
	return o.DomainCertAuthoritiesUri, true
}

// HasDomainCertAuthoritiesUri returns a boolean if a field has been set.
func (o *BrokerLinks) HasDomainCertAuthoritiesUri() bool {
	if o != nil && o.DomainCertAuthoritiesUri != nil {
		return true
	}

	return false
}

// SetDomainCertAuthoritiesUri gets a reference to the given string and assigns it to the DomainCertAuthoritiesUri field.
func (o *BrokerLinks) SetDomainCertAuthoritiesUri(v string) {
	o.DomainCertAuthoritiesUri = &v
}

// GetMsgVpnsUri returns the MsgVpnsUri field value if set, zero value otherwise.
func (o *BrokerLinks) GetMsgVpnsUri() string {
	if o == nil || o.MsgVpnsUri == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnsUri
}

// GetMsgVpnsUriOk returns a tuple with the MsgVpnsUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerLinks) GetMsgVpnsUriOk() (*string, bool) {
	if o == nil || o.MsgVpnsUri == nil {
		return nil, false
	}
	return o.MsgVpnsUri, true
}

// HasMsgVpnsUri returns a boolean if a field has been set.
func (o *BrokerLinks) HasMsgVpnsUri() bool {
	if o != nil && o.MsgVpnsUri != nil {
		return true
	}

	return false
}

// SetMsgVpnsUri gets a reference to the given string and assigns it to the MsgVpnsUri field.
func (o *BrokerLinks) SetMsgVpnsUri(v string) {
	o.MsgVpnsUri = &v
}

// GetSessionsUri returns the SessionsUri field value if set, zero value otherwise.
func (o *BrokerLinks) GetSessionsUri() string {
	if o == nil || o.SessionsUri == nil {
		var ret string
		return ret
	}
	return *o.SessionsUri
}

// GetSessionsUriOk returns a tuple with the SessionsUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerLinks) GetSessionsUriOk() (*string, bool) {
	if o == nil || o.SessionsUri == nil {
		return nil, false
	}
	return o.SessionsUri, true
}

// HasSessionsUri returns a boolean if a field has been set.
func (o *BrokerLinks) HasSessionsUri() bool {
	if o != nil && o.SessionsUri != nil {
		return true
	}

	return false
}

// SetSessionsUri gets a reference to the given string and assigns it to the SessionsUri field.
func (o *BrokerLinks) SetSessionsUri(v string) {
	o.SessionsUri = &v
}

// GetStandardDomainCertAuthoritiesUri returns the StandardDomainCertAuthoritiesUri field value if set, zero value otherwise.
func (o *BrokerLinks) GetStandardDomainCertAuthoritiesUri() string {
	if o == nil || o.StandardDomainCertAuthoritiesUri == nil {
		var ret string
		return ret
	}
	return *o.StandardDomainCertAuthoritiesUri
}

// GetStandardDomainCertAuthoritiesUriOk returns a tuple with the StandardDomainCertAuthoritiesUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerLinks) GetStandardDomainCertAuthoritiesUriOk() (*string, bool) {
	if o == nil || o.StandardDomainCertAuthoritiesUri == nil {
		return nil, false
	}
	return o.StandardDomainCertAuthoritiesUri, true
}

// HasStandardDomainCertAuthoritiesUri returns a boolean if a field has been set.
func (o *BrokerLinks) HasStandardDomainCertAuthoritiesUri() bool {
	if o != nil && o.StandardDomainCertAuthoritiesUri != nil {
		return true
	}

	return false
}

// SetStandardDomainCertAuthoritiesUri gets a reference to the given string and assigns it to the StandardDomainCertAuthoritiesUri field.
func (o *BrokerLinks) SetStandardDomainCertAuthoritiesUri(v string) {
	o.StandardDomainCertAuthoritiesUri = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BrokerLinks) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerLinks) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BrokerLinks) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BrokerLinks) SetUri(v string) {
	o.Uri = &v
}

// GetVirtualHostnamesUri returns the VirtualHostnamesUri field value if set, zero value otherwise.
func (o *BrokerLinks) GetVirtualHostnamesUri() string {
	if o == nil || o.VirtualHostnamesUri == nil {
		var ret string
		return ret
	}
	return *o.VirtualHostnamesUri
}

// GetVirtualHostnamesUriOk returns a tuple with the VirtualHostnamesUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerLinks) GetVirtualHostnamesUriOk() (*string, bool) {
	if o == nil || o.VirtualHostnamesUri == nil {
		return nil, false
	}
	return o.VirtualHostnamesUri, true
}

// HasVirtualHostnamesUri returns a boolean if a field has been set.
func (o *BrokerLinks) HasVirtualHostnamesUri() bool {
	if o != nil && o.VirtualHostnamesUri != nil {
		return true
	}

	return false
}

// SetVirtualHostnamesUri gets a reference to the given string and assigns it to the VirtualHostnamesUri field.
func (o *BrokerLinks) SetVirtualHostnamesUri(v string) {
	o.VirtualHostnamesUri = &v
}

func (o BrokerLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AboutUri != nil {
		toSerialize["aboutUri"] = o.AboutUri
	}
	if o.CertAuthoritiesUri != nil {
		toSerialize["certAuthoritiesUri"] = o.CertAuthoritiesUri
	}
	if o.ClientCertAuthoritiesUri != nil {
		toSerialize["clientCertAuthoritiesUri"] = o.ClientCertAuthoritiesUri
	}
	if o.DmrClustersUri != nil {
		toSerialize["dmrClustersUri"] = o.DmrClustersUri
	}
	if o.DomainCertAuthoritiesUri != nil {
		toSerialize["domainCertAuthoritiesUri"] = o.DomainCertAuthoritiesUri
	}
	if o.MsgVpnsUri != nil {
		toSerialize["msgVpnsUri"] = o.MsgVpnsUri
	}
	if o.SessionsUri != nil {
		toSerialize["sessionsUri"] = o.SessionsUri
	}
	if o.StandardDomainCertAuthoritiesUri != nil {
		toSerialize["standardDomainCertAuthoritiesUri"] = o.StandardDomainCertAuthoritiesUri
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	if o.VirtualHostnamesUri != nil {
		toSerialize["virtualHostnamesUri"] = o.VirtualHostnamesUri
	}
	return json.Marshal(toSerialize)
}

type NullableBrokerLinks struct {
	value *BrokerLinks
	isSet bool
}

func (v NullableBrokerLinks) Get() *BrokerLinks {
	return v.value
}

func (v *NullableBrokerLinks) Set(val *BrokerLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerLinks(val *BrokerLinks) *NullableBrokerLinks {
	return &NullableBrokerLinks{value: val, isSet: true}
}

func (v NullableBrokerLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
