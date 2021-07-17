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

// MsgVpnDistributedCacheClusterInstanceRate The rates associated with the Cache Instance. Deprecated since 2.13. All attributes in this object have been moved to the MsgVpnDistributedCacheClusterInstance object.
type MsgVpnDistributedCacheClusterInstanceRate struct {
	// The peak of the one minute average of the data message rate received by the Cache Instance, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	AverageDataRxBytePeakRate *int64 `json:"averageDataRxBytePeakRate,omitempty"`
	// The one minute average of the data message rate received by the Cache Instance, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	AverageDataRxByteRate *int64 `json:"averageDataRxByteRate,omitempty"`
	// The peak of the one minute average of the data message rate received by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	AverageDataRxMsgPeakRate *int64 `json:"averageDataRxMsgPeakRate,omitempty"`
	// The one minute average of the data message rate received by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	AverageDataRxMsgRate *int64 `json:"averageDataRxMsgRate,omitempty"`
	// The peak of the one minute average of the data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	AverageDataTxMsgPeakRate *int64 `json:"averageDataTxMsgPeakRate,omitempty"`
	// The one minute average of the data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	AverageDataTxMsgRate *int64 `json:"averageDataTxMsgRate,omitempty"`
	// The peak of the one minute average of the request rate received by the Cache Instance, in requests per second (req/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	AverageRequestRxPeakRate *int64 `json:"averageRequestRxPeakRate,omitempty"`
	// The one minute average of the request rate received by the Cache Instance, in requests per second (req/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	AverageRequestRxRate *int64 `json:"averageRequestRxRate,omitempty"`
	// The data message peak rate received by the Cache Instance, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	DataRxBytePeakRate *int64 `json:"dataRxBytePeakRate,omitempty"`
	// The data message rate received by the Cache Instance, in bytes per second (B/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	DataRxByteRate *int64 `json:"dataRxByteRate,omitempty"`
	// The data message peak rate received by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	DataRxMsgPeakRate *int64 `json:"dataRxMsgPeakRate,omitempty"`
	// The data message rate received by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	DataRxMsgRate *int64 `json:"dataRxMsgRate,omitempty"`
	// The data message peak rate transmitted by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	DataTxMsgPeakRate *int64 `json:"dataTxMsgPeakRate,omitempty"`
	// The data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	DataTxMsgRate *int64 `json:"dataTxMsgRate,omitempty"`
	// The request peak rate received by the Cache Instance, in requests per second (req/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	RequestRxPeakRate *int64 `json:"requestRxPeakRate,omitempty"`
	// The request rate received by the Cache Instance, in requests per second (req/sec). Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object.
	RequestRxRate *int64 `json:"requestRxRate,omitempty"`
}

// NewMsgVpnDistributedCacheClusterInstanceRate instantiates a new MsgVpnDistributedCacheClusterInstanceRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnDistributedCacheClusterInstanceRate() *MsgVpnDistributedCacheClusterInstanceRate {
	this := MsgVpnDistributedCacheClusterInstanceRate{}
	return &this
}

// NewMsgVpnDistributedCacheClusterInstanceRateWithDefaults instantiates a new MsgVpnDistributedCacheClusterInstanceRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnDistributedCacheClusterInstanceRateWithDefaults() *MsgVpnDistributedCacheClusterInstanceRate {
	this := MsgVpnDistributedCacheClusterInstanceRate{}
	return &this
}

// GetAverageDataRxBytePeakRate returns the AverageDataRxBytePeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxBytePeakRate() int64 {
	if o == nil || o.AverageDataRxBytePeakRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataRxBytePeakRate
}

// GetAverageDataRxBytePeakRateOk returns a tuple with the AverageDataRxBytePeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxBytePeakRateOk() (*int64, bool) {
	if o == nil || o.AverageDataRxBytePeakRate == nil {
		return nil, false
	}
	return o.AverageDataRxBytePeakRate, true
}

// HasAverageDataRxBytePeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataRxBytePeakRate() bool {
	if o != nil && o.AverageDataRxBytePeakRate != nil {
		return true
	}

	return false
}

// SetAverageDataRxBytePeakRate gets a reference to the given int64 and assigns it to the AverageDataRxBytePeakRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataRxBytePeakRate(v int64) {
	o.AverageDataRxBytePeakRate = &v
}

// GetAverageDataRxByteRate returns the AverageDataRxByteRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxByteRate() int64 {
	if o == nil || o.AverageDataRxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataRxByteRate
}

// GetAverageDataRxByteRateOk returns a tuple with the AverageDataRxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxByteRateOk() (*int64, bool) {
	if o == nil || o.AverageDataRxByteRate == nil {
		return nil, false
	}
	return o.AverageDataRxByteRate, true
}

// HasAverageDataRxByteRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataRxByteRate() bool {
	if o != nil && o.AverageDataRxByteRate != nil {
		return true
	}

	return false
}

// SetAverageDataRxByteRate gets a reference to the given int64 and assigns it to the AverageDataRxByteRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataRxByteRate(v int64) {
	o.AverageDataRxByteRate = &v
}

// GetAverageDataRxMsgPeakRate returns the AverageDataRxMsgPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxMsgPeakRate() int64 {
	if o == nil || o.AverageDataRxMsgPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataRxMsgPeakRate
}

// GetAverageDataRxMsgPeakRateOk returns a tuple with the AverageDataRxMsgPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxMsgPeakRateOk() (*int64, bool) {
	if o == nil || o.AverageDataRxMsgPeakRate == nil {
		return nil, false
	}
	return o.AverageDataRxMsgPeakRate, true
}

// HasAverageDataRxMsgPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataRxMsgPeakRate() bool {
	if o != nil && o.AverageDataRxMsgPeakRate != nil {
		return true
	}

	return false
}

// SetAverageDataRxMsgPeakRate gets a reference to the given int64 and assigns it to the AverageDataRxMsgPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataRxMsgPeakRate(v int64) {
	o.AverageDataRxMsgPeakRate = &v
}

// GetAverageDataRxMsgRate returns the AverageDataRxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxMsgRate() int64 {
	if o == nil || o.AverageDataRxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataRxMsgRate
}

// GetAverageDataRxMsgRateOk returns a tuple with the AverageDataRxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataRxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageDataRxMsgRate == nil {
		return nil, false
	}
	return o.AverageDataRxMsgRate, true
}

// HasAverageDataRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataRxMsgRate() bool {
	if o != nil && o.AverageDataRxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageDataRxMsgRate gets a reference to the given int64 and assigns it to the AverageDataRxMsgRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataRxMsgRate(v int64) {
	o.AverageDataRxMsgRate = &v
}

// GetAverageDataTxMsgPeakRate returns the AverageDataTxMsgPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataTxMsgPeakRate() int64 {
	if o == nil || o.AverageDataTxMsgPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataTxMsgPeakRate
}

// GetAverageDataTxMsgPeakRateOk returns a tuple with the AverageDataTxMsgPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataTxMsgPeakRateOk() (*int64, bool) {
	if o == nil || o.AverageDataTxMsgPeakRate == nil {
		return nil, false
	}
	return o.AverageDataTxMsgPeakRate, true
}

// HasAverageDataTxMsgPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataTxMsgPeakRate() bool {
	if o != nil && o.AverageDataTxMsgPeakRate != nil {
		return true
	}

	return false
}

// SetAverageDataTxMsgPeakRate gets a reference to the given int64 and assigns it to the AverageDataTxMsgPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataTxMsgPeakRate(v int64) {
	o.AverageDataTxMsgPeakRate = &v
}

// GetAverageDataTxMsgRate returns the AverageDataTxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataTxMsgRate() int64 {
	if o == nil || o.AverageDataTxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataTxMsgRate
}

// GetAverageDataTxMsgRateOk returns a tuple with the AverageDataTxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageDataTxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageDataTxMsgRate == nil {
		return nil, false
	}
	return o.AverageDataTxMsgRate, true
}

// HasAverageDataTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageDataTxMsgRate() bool {
	if o != nil && o.AverageDataTxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageDataTxMsgRate gets a reference to the given int64 and assigns it to the AverageDataTxMsgRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageDataTxMsgRate(v int64) {
	o.AverageDataTxMsgRate = &v
}

// GetAverageRequestRxPeakRate returns the AverageRequestRxPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageRequestRxPeakRate() int64 {
	if o == nil || o.AverageRequestRxPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRequestRxPeakRate
}

// GetAverageRequestRxPeakRateOk returns a tuple with the AverageRequestRxPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageRequestRxPeakRateOk() (*int64, bool) {
	if o == nil || o.AverageRequestRxPeakRate == nil {
		return nil, false
	}
	return o.AverageRequestRxPeakRate, true
}

// HasAverageRequestRxPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageRequestRxPeakRate() bool {
	if o != nil && o.AverageRequestRxPeakRate != nil {
		return true
	}

	return false
}

// SetAverageRequestRxPeakRate gets a reference to the given int64 and assigns it to the AverageRequestRxPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageRequestRxPeakRate(v int64) {
	o.AverageRequestRxPeakRate = &v
}

// GetAverageRequestRxRate returns the AverageRequestRxRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageRequestRxRate() int64 {
	if o == nil || o.AverageRequestRxRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRequestRxRate
}

// GetAverageRequestRxRateOk returns a tuple with the AverageRequestRxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetAverageRequestRxRateOk() (*int64, bool) {
	if o == nil || o.AverageRequestRxRate == nil {
		return nil, false
	}
	return o.AverageRequestRxRate, true
}

// HasAverageRequestRxRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasAverageRequestRxRate() bool {
	if o != nil && o.AverageRequestRxRate != nil {
		return true
	}

	return false
}

// SetAverageRequestRxRate gets a reference to the given int64 and assigns it to the AverageRequestRxRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetAverageRequestRxRate(v int64) {
	o.AverageRequestRxRate = &v
}

// GetDataRxBytePeakRate returns the DataRxBytePeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxBytePeakRate() int64 {
	if o == nil || o.DataRxBytePeakRate == nil {
		var ret int64
		return ret
	}
	return *o.DataRxBytePeakRate
}

// GetDataRxBytePeakRateOk returns a tuple with the DataRxBytePeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxBytePeakRateOk() (*int64, bool) {
	if o == nil || o.DataRxBytePeakRate == nil {
		return nil, false
	}
	return o.DataRxBytePeakRate, true
}

// HasDataRxBytePeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataRxBytePeakRate() bool {
	if o != nil && o.DataRxBytePeakRate != nil {
		return true
	}

	return false
}

// SetDataRxBytePeakRate gets a reference to the given int64 and assigns it to the DataRxBytePeakRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataRxBytePeakRate(v int64) {
	o.DataRxBytePeakRate = &v
}

// GetDataRxByteRate returns the DataRxByteRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxByteRate() int64 {
	if o == nil || o.DataRxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.DataRxByteRate
}

// GetDataRxByteRateOk returns a tuple with the DataRxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxByteRateOk() (*int64, bool) {
	if o == nil || o.DataRxByteRate == nil {
		return nil, false
	}
	return o.DataRxByteRate, true
}

// HasDataRxByteRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataRxByteRate() bool {
	if o != nil && o.DataRxByteRate != nil {
		return true
	}

	return false
}

// SetDataRxByteRate gets a reference to the given int64 and assigns it to the DataRxByteRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataRxByteRate(v int64) {
	o.DataRxByteRate = &v
}

// GetDataRxMsgPeakRate returns the DataRxMsgPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxMsgPeakRate() int64 {
	if o == nil || o.DataRxMsgPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.DataRxMsgPeakRate
}

// GetDataRxMsgPeakRateOk returns a tuple with the DataRxMsgPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxMsgPeakRateOk() (*int64, bool) {
	if o == nil || o.DataRxMsgPeakRate == nil {
		return nil, false
	}
	return o.DataRxMsgPeakRate, true
}

// HasDataRxMsgPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataRxMsgPeakRate() bool {
	if o != nil && o.DataRxMsgPeakRate != nil {
		return true
	}

	return false
}

// SetDataRxMsgPeakRate gets a reference to the given int64 and assigns it to the DataRxMsgPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataRxMsgPeakRate(v int64) {
	o.DataRxMsgPeakRate = &v
}

// GetDataRxMsgRate returns the DataRxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxMsgRate() int64 {
	if o == nil || o.DataRxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.DataRxMsgRate
}

// GetDataRxMsgRateOk returns a tuple with the DataRxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataRxMsgRateOk() (*int64, bool) {
	if o == nil || o.DataRxMsgRate == nil {
		return nil, false
	}
	return o.DataRxMsgRate, true
}

// HasDataRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataRxMsgRate() bool {
	if o != nil && o.DataRxMsgRate != nil {
		return true
	}

	return false
}

// SetDataRxMsgRate gets a reference to the given int64 and assigns it to the DataRxMsgRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataRxMsgRate(v int64) {
	o.DataRxMsgRate = &v
}

// GetDataTxMsgPeakRate returns the DataTxMsgPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataTxMsgPeakRate() int64 {
	if o == nil || o.DataTxMsgPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.DataTxMsgPeakRate
}

// GetDataTxMsgPeakRateOk returns a tuple with the DataTxMsgPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataTxMsgPeakRateOk() (*int64, bool) {
	if o == nil || o.DataTxMsgPeakRate == nil {
		return nil, false
	}
	return o.DataTxMsgPeakRate, true
}

// HasDataTxMsgPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataTxMsgPeakRate() bool {
	if o != nil && o.DataTxMsgPeakRate != nil {
		return true
	}

	return false
}

// SetDataTxMsgPeakRate gets a reference to the given int64 and assigns it to the DataTxMsgPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataTxMsgPeakRate(v int64) {
	o.DataTxMsgPeakRate = &v
}

// GetDataTxMsgRate returns the DataTxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataTxMsgRate() int64 {
	if o == nil || o.DataTxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.DataTxMsgRate
}

// GetDataTxMsgRateOk returns a tuple with the DataTxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetDataTxMsgRateOk() (*int64, bool) {
	if o == nil || o.DataTxMsgRate == nil {
		return nil, false
	}
	return o.DataTxMsgRate, true
}

// HasDataTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasDataTxMsgRate() bool {
	if o != nil && o.DataTxMsgRate != nil {
		return true
	}

	return false
}

// SetDataTxMsgRate gets a reference to the given int64 and assigns it to the DataTxMsgRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetDataTxMsgRate(v int64) {
	o.DataTxMsgRate = &v
}

// GetRequestRxPeakRate returns the RequestRxPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetRequestRxPeakRate() int64 {
	if o == nil || o.RequestRxPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.RequestRxPeakRate
}

// GetRequestRxPeakRateOk returns a tuple with the RequestRxPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetRequestRxPeakRateOk() (*int64, bool) {
	if o == nil || o.RequestRxPeakRate == nil {
		return nil, false
	}
	return o.RequestRxPeakRate, true
}

// HasRequestRxPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasRequestRxPeakRate() bool {
	if o != nil && o.RequestRxPeakRate != nil {
		return true
	}

	return false
}

// SetRequestRxPeakRate gets a reference to the given int64 and assigns it to the RequestRxPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetRequestRxPeakRate(v int64) {
	o.RequestRxPeakRate = &v
}

// GetRequestRxRate returns the RequestRxRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetRequestRxRate() int64 {
	if o == nil || o.RequestRxRate == nil {
		var ret int64
		return ret
	}
	return *o.RequestRxRate
}

// GetRequestRxRateOk returns a tuple with the RequestRxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) GetRequestRxRateOk() (*int64, bool) {
	if o == nil || o.RequestRxRate == nil {
		return nil, false
	}
	return o.RequestRxRate, true
}

// HasRequestRxRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstanceRate) HasRequestRxRate() bool {
	if o != nil && o.RequestRxRate != nil {
		return true
	}

	return false
}

// SetRequestRxRate gets a reference to the given int64 and assigns it to the RequestRxRate field.
func (o *MsgVpnDistributedCacheClusterInstanceRate) SetRequestRxRate(v int64) {
	o.RequestRxRate = &v
}

func (o MsgVpnDistributedCacheClusterInstanceRate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AverageDataRxBytePeakRate != nil {
		toSerialize["averageDataRxBytePeakRate"] = o.AverageDataRxBytePeakRate
	}
	if o.AverageDataRxByteRate != nil {
		toSerialize["averageDataRxByteRate"] = o.AverageDataRxByteRate
	}
	if o.AverageDataRxMsgPeakRate != nil {
		toSerialize["averageDataRxMsgPeakRate"] = o.AverageDataRxMsgPeakRate
	}
	if o.AverageDataRxMsgRate != nil {
		toSerialize["averageDataRxMsgRate"] = o.AverageDataRxMsgRate
	}
	if o.AverageDataTxMsgPeakRate != nil {
		toSerialize["averageDataTxMsgPeakRate"] = o.AverageDataTxMsgPeakRate
	}
	if o.AverageDataTxMsgRate != nil {
		toSerialize["averageDataTxMsgRate"] = o.AverageDataTxMsgRate
	}
	if o.AverageRequestRxPeakRate != nil {
		toSerialize["averageRequestRxPeakRate"] = o.AverageRequestRxPeakRate
	}
	if o.AverageRequestRxRate != nil {
		toSerialize["averageRequestRxRate"] = o.AverageRequestRxRate
	}
	if o.DataRxBytePeakRate != nil {
		toSerialize["dataRxBytePeakRate"] = o.DataRxBytePeakRate
	}
	if o.DataRxByteRate != nil {
		toSerialize["dataRxByteRate"] = o.DataRxByteRate
	}
	if o.DataRxMsgPeakRate != nil {
		toSerialize["dataRxMsgPeakRate"] = o.DataRxMsgPeakRate
	}
	if o.DataRxMsgRate != nil {
		toSerialize["dataRxMsgRate"] = o.DataRxMsgRate
	}
	if o.DataTxMsgPeakRate != nil {
		toSerialize["dataTxMsgPeakRate"] = o.DataTxMsgPeakRate
	}
	if o.DataTxMsgRate != nil {
		toSerialize["dataTxMsgRate"] = o.DataTxMsgRate
	}
	if o.RequestRxPeakRate != nil {
		toSerialize["requestRxPeakRate"] = o.RequestRxPeakRate
	}
	if o.RequestRxRate != nil {
		toSerialize["requestRxRate"] = o.RequestRxRate
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnDistributedCacheClusterInstanceRate struct {
	value *MsgVpnDistributedCacheClusterInstanceRate
	isSet bool
}

func (v NullableMsgVpnDistributedCacheClusterInstanceRate) Get() *MsgVpnDistributedCacheClusterInstanceRate {
	return v.value
}

func (v *NullableMsgVpnDistributedCacheClusterInstanceRate) Set(val *MsgVpnDistributedCacheClusterInstanceRate) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnDistributedCacheClusterInstanceRate) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnDistributedCacheClusterInstanceRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnDistributedCacheClusterInstanceRate(val *MsgVpnDistributedCacheClusterInstanceRate) *NullableMsgVpnDistributedCacheClusterInstanceRate {
	return &NullableMsgVpnDistributedCacheClusterInstanceRate{value: val, isSet: true}
}

func (v NullableMsgVpnDistributedCacheClusterInstanceRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnDistributedCacheClusterInstanceRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
