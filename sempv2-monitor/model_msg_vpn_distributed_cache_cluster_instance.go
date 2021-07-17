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

// MsgVpnDistributedCacheClusterInstance struct for MsgVpnDistributedCacheClusterInstance
type MsgVpnDistributedCacheClusterInstance struct {
	// Indicates whether auto-start for the Cache Instance is enabled, and the Cache Instance will automatically attempt to transition from the Stopped operational state to Up whenever it restarts or reconnects to the message broker.
	AutoStartEnabled *bool `json:"autoStartEnabled,omitempty"`
	// The peak of the one minute average of the data message rate received by the Cache Instance, in bytes per second (B/sec). Available since 2.13.
	AverageDataRxBytePeakRate *int64 `json:"averageDataRxBytePeakRate,omitempty"`
	// The one minute average of the data message rate received by the Cache Instance, in bytes per second (B/sec). Available since 2.13.
	AverageDataRxByteRate *int64 `json:"averageDataRxByteRate,omitempty"`
	// The peak of the one minute average of the data message rate received by the Cache Instance, in messages per second (msg/sec). Available since 2.13.
	AverageDataRxMsgPeakRate *int64 `json:"averageDataRxMsgPeakRate,omitempty"`
	// The one minute average of the data message rate received by the Cache Instance, in messages per second (msg/sec). Available since 2.13.
	AverageDataRxMsgRate *int64 `json:"averageDataRxMsgRate,omitempty"`
	// The peak of the one minute average of the data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Available since 2.13.
	AverageDataTxMsgPeakRate *int64 `json:"averageDataTxMsgPeakRate,omitempty"`
	// The one minute average of the data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Available since 2.13.
	AverageDataTxMsgRate *int64 `json:"averageDataTxMsgRate,omitempty"`
	// The peak of the one minute average of the request rate received by the Cache Instance, in requests per second (req/sec). Available since 2.13.
	AverageRequestRxPeakRate *int64 `json:"averageRequestRxPeakRate,omitempty"`
	// The one minute average of the request rate received by the Cache Instance, in requests per second (req/sec). Available since 2.13.
	AverageRequestRxRate *int64 `json:"averageRequestRxRate,omitempty"`
	// The name of the Distributed Cache.
	CacheName *string `json:"cacheName,omitempty"`
	// The name of the Cache Cluster.
	ClusterName *string                                       `json:"clusterName,omitempty"`
	Counter     *MsgVpnDistributedCacheClusterInstanceCounter `json:"counter,omitempty"`
	// The data message peak rate received by the Cache Instance, in bytes per second (B/sec). Available since 2.13.
	DataRxBytePeakRate *int64 `json:"dataRxBytePeakRate,omitempty"`
	// The data message rate received by the Cache Instance, in bytes per second (B/sec). Available since 2.13.
	DataRxByteRate *int64 `json:"dataRxByteRate,omitempty"`
	// The data message peak rate received by the Cache Instance, in messages per second (msg/sec). Available since 2.13.
	DataRxMsgPeakRate *int64 `json:"dataRxMsgPeakRate,omitempty"`
	// The data message rate received by the Cache Instance, in messages per second (msg/sec). Available since 2.13.
	DataRxMsgRate *int64 `json:"dataRxMsgRate,omitempty"`
	// The data message peak rate transmitted by the Cache Instance, in messages per second (msg/sec). Available since 2.13.
	DataTxMsgPeakRate *int64 `json:"dataTxMsgPeakRate,omitempty"`
	// The data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Available since 2.13.
	DataTxMsgRate *int64 `json:"dataTxMsgRate,omitempty"`
	// Indicates whether the Cache Instance is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the Cache Instance.
	InstanceName *string `json:"instanceName,omitempty"`
	// The timestamp of when the Cache Instance last registered with the message broker. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastRegisteredTime *int32 `json:"lastRegisteredTime,omitempty"`
	// The timestamp of the last heartbeat message received from the Cache Instance. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastRxHeartbeatTime *int32 `json:"lastRxHeartbeatTime,omitempty"`
	// The timestamp of the last request for setting the lost message indication received from the Cache Instance. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastRxSetLostMsgTime *int32 `json:"lastRxSetLostMsgTime,omitempty"`
	// The reason why the Cache Instance was last stopped.
	LastStoppedReason *string `json:"lastStoppedReason,omitempty"`
	// The timestamp of when the Cache Instance was last stopped. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastStoppedTime *int32 `json:"lastStoppedTime,omitempty"`
	// The timestamp of the last request for clearing the lost message indication transmitted to the Cache Instance. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastTxClearLostMsgTime *int32 `json:"lastTxClearLostMsgTime,omitempty"`
	// The memory usage of the Cache Instance, in megabytes (MB).
	MemoryUsage *int32 `json:"memoryUsage,omitempty"`
	// The number of messages cached for the Cache Instance. Available since 2.13.
	MsgCount *int64 `json:"msgCount,omitempty"`
	// The number of messages cached peak for the Cache Instance. Available since 2.13.
	MsgPeakCount *int64 `json:"msgPeakCount,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// Indicates whether one or more messages were lost by the Cache Instance.
	MsgsLost *bool                                      `json:"msgsLost,omitempty"`
	Rate     *MsgVpnDistributedCacheClusterInstanceRate `json:"rate,omitempty"`
	// The received request message queue depth for the Cache Instance. Available since 2.13.
	RequestQueueDepthCount *int64 `json:"requestQueueDepthCount,omitempty"`
	// The received request message queue depth peak for the Cache Instance. Available since 2.13.
	RequestQueueDepthPeakCount *int64 `json:"requestQueueDepthPeakCount,omitempty"`
	// The request peak rate received by the Cache Instance, in requests per second (req/sec). Available since 2.13.
	RequestRxPeakRate *int64 `json:"requestRxPeakRate,omitempty"`
	// The request rate received by the Cache Instance, in requests per second (req/sec). Available since 2.13.
	RequestRxRate *int64 `json:"requestRxRate,omitempty"`
	// The operational state of the Cache Instance. The allowed values and their meaning are:  <pre> \"invalid\" - The Cache Instance state is invalid. \"down\" - The Cache Instance is operationally down. \"stopped\" - The Cache Instance has stopped processing cache requests. \"stopped-lost-msg\" - The Cache Instance has stopped due to a lost message. \"register\" - The Cache Instance is registering with the broker. \"config-sync\" - The Cache Instance is synchronizing its configuration with the broker. \"cluster-sync\" - The Cache Instance is synchronizing its messages with the Cache Cluster. \"up\" - The Cache Instance is operationally up. \"backup\" - The Cache Instance is backing up its messages to disk. \"restore\" - The Cache Instance is restoring its messages from disk. \"not-available\" - The Cache Instance state is not available. </pre>
	State *string `json:"state,omitempty"`
	// Indicates whether stop-on-lost-message is enabled, and the Cache Instance will transition to the Stopped operational state upon losing a message. When Stopped, it cannot accept or respond to cache requests, but continues to cache messages.
	StopOnLostMsgEnabled *bool `json:"stopOnLostMsgEnabled,omitempty"`
	// The number of topics cached for the Cache Instance. Available since 2.13.
	TopicCount *int64 `json:"topicCount,omitempty"`
	// The number of topics cached peak for the Cache Instance. Available since 2.13.
	TopicPeakCount *int64 `json:"topicPeakCount,omitempty"`
}

// NewMsgVpnDistributedCacheClusterInstance instantiates a new MsgVpnDistributedCacheClusterInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnDistributedCacheClusterInstance() *MsgVpnDistributedCacheClusterInstance {
	this := MsgVpnDistributedCacheClusterInstance{}
	return &this
}

// NewMsgVpnDistributedCacheClusterInstanceWithDefaults instantiates a new MsgVpnDistributedCacheClusterInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnDistributedCacheClusterInstanceWithDefaults() *MsgVpnDistributedCacheClusterInstance {
	this := MsgVpnDistributedCacheClusterInstance{}
	return &this
}

// GetAutoStartEnabled returns the AutoStartEnabled field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetAutoStartEnabled() bool {
	if o == nil || o.AutoStartEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AutoStartEnabled
}

// GetAutoStartEnabledOk returns a tuple with the AutoStartEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetAutoStartEnabledOk() (*bool, bool) {
	if o == nil || o.AutoStartEnabled == nil {
		return nil, false
	}
	return o.AutoStartEnabled, true
}

// HasAutoStartEnabled returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasAutoStartEnabled() bool {
	if o != nil && o.AutoStartEnabled != nil {
		return true
	}

	return false
}

// SetAutoStartEnabled gets a reference to the given bool and assigns it to the AutoStartEnabled field.
func (o *MsgVpnDistributedCacheClusterInstance) SetAutoStartEnabled(v bool) {
	o.AutoStartEnabled = &v
}

// GetAverageDataRxBytePeakRate returns the AverageDataRxBytePeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxBytePeakRate() int64 {
	if o == nil || o.AverageDataRxBytePeakRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataRxBytePeakRate
}

// GetAverageDataRxBytePeakRateOk returns a tuple with the AverageDataRxBytePeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxBytePeakRateOk() (*int64, bool) {
	if o == nil || o.AverageDataRxBytePeakRate == nil {
		return nil, false
	}
	return o.AverageDataRxBytePeakRate, true
}

// HasAverageDataRxBytePeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataRxBytePeakRate() bool {
	if o != nil && o.AverageDataRxBytePeakRate != nil {
		return true
	}

	return false
}

// SetAverageDataRxBytePeakRate gets a reference to the given int64 and assigns it to the AverageDataRxBytePeakRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataRxBytePeakRate(v int64) {
	o.AverageDataRxBytePeakRate = &v
}

// GetAverageDataRxByteRate returns the AverageDataRxByteRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxByteRate() int64 {
	if o == nil || o.AverageDataRxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataRxByteRate
}

// GetAverageDataRxByteRateOk returns a tuple with the AverageDataRxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxByteRateOk() (*int64, bool) {
	if o == nil || o.AverageDataRxByteRate == nil {
		return nil, false
	}
	return o.AverageDataRxByteRate, true
}

// HasAverageDataRxByteRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataRxByteRate() bool {
	if o != nil && o.AverageDataRxByteRate != nil {
		return true
	}

	return false
}

// SetAverageDataRxByteRate gets a reference to the given int64 and assigns it to the AverageDataRxByteRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataRxByteRate(v int64) {
	o.AverageDataRxByteRate = &v
}

// GetAverageDataRxMsgPeakRate returns the AverageDataRxMsgPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxMsgPeakRate() int64 {
	if o == nil || o.AverageDataRxMsgPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataRxMsgPeakRate
}

// GetAverageDataRxMsgPeakRateOk returns a tuple with the AverageDataRxMsgPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxMsgPeakRateOk() (*int64, bool) {
	if o == nil || o.AverageDataRxMsgPeakRate == nil {
		return nil, false
	}
	return o.AverageDataRxMsgPeakRate, true
}

// HasAverageDataRxMsgPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataRxMsgPeakRate() bool {
	if o != nil && o.AverageDataRxMsgPeakRate != nil {
		return true
	}

	return false
}

// SetAverageDataRxMsgPeakRate gets a reference to the given int64 and assigns it to the AverageDataRxMsgPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataRxMsgPeakRate(v int64) {
	o.AverageDataRxMsgPeakRate = &v
}

// GetAverageDataRxMsgRate returns the AverageDataRxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxMsgRate() int64 {
	if o == nil || o.AverageDataRxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataRxMsgRate
}

// GetAverageDataRxMsgRateOk returns a tuple with the AverageDataRxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageDataRxMsgRate == nil {
		return nil, false
	}
	return o.AverageDataRxMsgRate, true
}

// HasAverageDataRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataRxMsgRate() bool {
	if o != nil && o.AverageDataRxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageDataRxMsgRate gets a reference to the given int64 and assigns it to the AverageDataRxMsgRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataRxMsgRate(v int64) {
	o.AverageDataRxMsgRate = &v
}

// GetAverageDataTxMsgPeakRate returns the AverageDataTxMsgPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataTxMsgPeakRate() int64 {
	if o == nil || o.AverageDataTxMsgPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataTxMsgPeakRate
}

// GetAverageDataTxMsgPeakRateOk returns a tuple with the AverageDataTxMsgPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataTxMsgPeakRateOk() (*int64, bool) {
	if o == nil || o.AverageDataTxMsgPeakRate == nil {
		return nil, false
	}
	return o.AverageDataTxMsgPeakRate, true
}

// HasAverageDataTxMsgPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataTxMsgPeakRate() bool {
	if o != nil && o.AverageDataTxMsgPeakRate != nil {
		return true
	}

	return false
}

// SetAverageDataTxMsgPeakRate gets a reference to the given int64 and assigns it to the AverageDataTxMsgPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataTxMsgPeakRate(v int64) {
	o.AverageDataTxMsgPeakRate = &v
}

// GetAverageDataTxMsgRate returns the AverageDataTxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataTxMsgRate() int64 {
	if o == nil || o.AverageDataTxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageDataTxMsgRate
}

// GetAverageDataTxMsgRateOk returns a tuple with the AverageDataTxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataTxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageDataTxMsgRate == nil {
		return nil, false
	}
	return o.AverageDataTxMsgRate, true
}

// HasAverageDataTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataTxMsgRate() bool {
	if o != nil && o.AverageDataTxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageDataTxMsgRate gets a reference to the given int64 and assigns it to the AverageDataTxMsgRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataTxMsgRate(v int64) {
	o.AverageDataTxMsgRate = &v
}

// GetAverageRequestRxPeakRate returns the AverageRequestRxPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageRequestRxPeakRate() int64 {
	if o == nil || o.AverageRequestRxPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRequestRxPeakRate
}

// GetAverageRequestRxPeakRateOk returns a tuple with the AverageRequestRxPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageRequestRxPeakRateOk() (*int64, bool) {
	if o == nil || o.AverageRequestRxPeakRate == nil {
		return nil, false
	}
	return o.AverageRequestRxPeakRate, true
}

// HasAverageRequestRxPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasAverageRequestRxPeakRate() bool {
	if o != nil && o.AverageRequestRxPeakRate != nil {
		return true
	}

	return false
}

// SetAverageRequestRxPeakRate gets a reference to the given int64 and assigns it to the AverageRequestRxPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetAverageRequestRxPeakRate(v int64) {
	o.AverageRequestRxPeakRate = &v
}

// GetAverageRequestRxRate returns the AverageRequestRxRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageRequestRxRate() int64 {
	if o == nil || o.AverageRequestRxRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRequestRxRate
}

// GetAverageRequestRxRateOk returns a tuple with the AverageRequestRxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetAverageRequestRxRateOk() (*int64, bool) {
	if o == nil || o.AverageRequestRxRate == nil {
		return nil, false
	}
	return o.AverageRequestRxRate, true
}

// HasAverageRequestRxRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasAverageRequestRxRate() bool {
	if o != nil && o.AverageRequestRxRate != nil {
		return true
	}

	return false
}

// SetAverageRequestRxRate gets a reference to the given int64 and assigns it to the AverageRequestRxRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetAverageRequestRxRate(v int64) {
	o.AverageRequestRxRate = &v
}

// GetCacheName returns the CacheName field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetCacheName() string {
	if o == nil || o.CacheName == nil {
		var ret string
		return ret
	}
	return *o.CacheName
}

// GetCacheNameOk returns a tuple with the CacheName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetCacheNameOk() (*string, bool) {
	if o == nil || o.CacheName == nil {
		return nil, false
	}
	return o.CacheName, true
}

// HasCacheName returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasCacheName() bool {
	if o != nil && o.CacheName != nil {
		return true
	}

	return false
}

// SetCacheName gets a reference to the given string and assigns it to the CacheName field.
func (o *MsgVpnDistributedCacheClusterInstance) SetCacheName(v string) {
	o.CacheName = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *MsgVpnDistributedCacheClusterInstance) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCounter returns the Counter field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetCounter() MsgVpnDistributedCacheClusterInstanceCounter {
	if o == nil || o.Counter == nil {
		var ret MsgVpnDistributedCacheClusterInstanceCounter
		return ret
	}
	return *o.Counter
}

// GetCounterOk returns a tuple with the Counter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetCounterOk() (*MsgVpnDistributedCacheClusterInstanceCounter, bool) {
	if o == nil || o.Counter == nil {
		return nil, false
	}
	return o.Counter, true
}

// HasCounter returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasCounter() bool {
	if o != nil && o.Counter != nil {
		return true
	}

	return false
}

// SetCounter gets a reference to the given MsgVpnDistributedCacheClusterInstanceCounter and assigns it to the Counter field.
func (o *MsgVpnDistributedCacheClusterInstance) SetCounter(v MsgVpnDistributedCacheClusterInstanceCounter) {
	o.Counter = &v
}

// GetDataRxBytePeakRate returns the DataRxBytePeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxBytePeakRate() int64 {
	if o == nil || o.DataRxBytePeakRate == nil {
		var ret int64
		return ret
	}
	return *o.DataRxBytePeakRate
}

// GetDataRxBytePeakRateOk returns a tuple with the DataRxBytePeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxBytePeakRateOk() (*int64, bool) {
	if o == nil || o.DataRxBytePeakRate == nil {
		return nil, false
	}
	return o.DataRxBytePeakRate, true
}

// HasDataRxBytePeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasDataRxBytePeakRate() bool {
	if o != nil && o.DataRxBytePeakRate != nil {
		return true
	}

	return false
}

// SetDataRxBytePeakRate gets a reference to the given int64 and assigns it to the DataRxBytePeakRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetDataRxBytePeakRate(v int64) {
	o.DataRxBytePeakRate = &v
}

// GetDataRxByteRate returns the DataRxByteRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxByteRate() int64 {
	if o == nil || o.DataRxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.DataRxByteRate
}

// GetDataRxByteRateOk returns a tuple with the DataRxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxByteRateOk() (*int64, bool) {
	if o == nil || o.DataRxByteRate == nil {
		return nil, false
	}
	return o.DataRxByteRate, true
}

// HasDataRxByteRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasDataRxByteRate() bool {
	if o != nil && o.DataRxByteRate != nil {
		return true
	}

	return false
}

// SetDataRxByteRate gets a reference to the given int64 and assigns it to the DataRxByteRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetDataRxByteRate(v int64) {
	o.DataRxByteRate = &v
}

// GetDataRxMsgPeakRate returns the DataRxMsgPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxMsgPeakRate() int64 {
	if o == nil || o.DataRxMsgPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.DataRxMsgPeakRate
}

// GetDataRxMsgPeakRateOk returns a tuple with the DataRxMsgPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxMsgPeakRateOk() (*int64, bool) {
	if o == nil || o.DataRxMsgPeakRate == nil {
		return nil, false
	}
	return o.DataRxMsgPeakRate, true
}

// HasDataRxMsgPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasDataRxMsgPeakRate() bool {
	if o != nil && o.DataRxMsgPeakRate != nil {
		return true
	}

	return false
}

// SetDataRxMsgPeakRate gets a reference to the given int64 and assigns it to the DataRxMsgPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetDataRxMsgPeakRate(v int64) {
	o.DataRxMsgPeakRate = &v
}

// GetDataRxMsgRate returns the DataRxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxMsgRate() int64 {
	if o == nil || o.DataRxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.DataRxMsgRate
}

// GetDataRxMsgRateOk returns a tuple with the DataRxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxMsgRateOk() (*int64, bool) {
	if o == nil || o.DataRxMsgRate == nil {
		return nil, false
	}
	return o.DataRxMsgRate, true
}

// HasDataRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasDataRxMsgRate() bool {
	if o != nil && o.DataRxMsgRate != nil {
		return true
	}

	return false
}

// SetDataRxMsgRate gets a reference to the given int64 and assigns it to the DataRxMsgRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetDataRxMsgRate(v int64) {
	o.DataRxMsgRate = &v
}

// GetDataTxMsgPeakRate returns the DataTxMsgPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataTxMsgPeakRate() int64 {
	if o == nil || o.DataTxMsgPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.DataTxMsgPeakRate
}

// GetDataTxMsgPeakRateOk returns a tuple with the DataTxMsgPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataTxMsgPeakRateOk() (*int64, bool) {
	if o == nil || o.DataTxMsgPeakRate == nil {
		return nil, false
	}
	return o.DataTxMsgPeakRate, true
}

// HasDataTxMsgPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasDataTxMsgPeakRate() bool {
	if o != nil && o.DataTxMsgPeakRate != nil {
		return true
	}

	return false
}

// SetDataTxMsgPeakRate gets a reference to the given int64 and assigns it to the DataTxMsgPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetDataTxMsgPeakRate(v int64) {
	o.DataTxMsgPeakRate = &v
}

// GetDataTxMsgRate returns the DataTxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataTxMsgRate() int64 {
	if o == nil || o.DataTxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.DataTxMsgRate
}

// GetDataTxMsgRateOk returns a tuple with the DataTxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetDataTxMsgRateOk() (*int64, bool) {
	if o == nil || o.DataTxMsgRate == nil {
		return nil, false
	}
	return o.DataTxMsgRate, true
}

// HasDataTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasDataTxMsgRate() bool {
	if o != nil && o.DataTxMsgRate != nil {
		return true
	}

	return false
}

// SetDataTxMsgRate gets a reference to the given int64 and assigns it to the DataTxMsgRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetDataTxMsgRate(v int64) {
	o.DataTxMsgRate = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MsgVpnDistributedCacheClusterInstance) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasInstanceName() bool {
	if o != nil && o.InstanceName != nil {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *MsgVpnDistributedCacheClusterInstance) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetLastRegisteredTime returns the LastRegisteredTime field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastRegisteredTime() int32 {
	if o == nil || o.LastRegisteredTime == nil {
		var ret int32
		return ret
	}
	return *o.LastRegisteredTime
}

// GetLastRegisteredTimeOk returns a tuple with the LastRegisteredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastRegisteredTimeOk() (*int32, bool) {
	if o == nil || o.LastRegisteredTime == nil {
		return nil, false
	}
	return o.LastRegisteredTime, true
}

// HasLastRegisteredTime returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasLastRegisteredTime() bool {
	if o != nil && o.LastRegisteredTime != nil {
		return true
	}

	return false
}

// SetLastRegisteredTime gets a reference to the given int32 and assigns it to the LastRegisteredTime field.
func (o *MsgVpnDistributedCacheClusterInstance) SetLastRegisteredTime(v int32) {
	o.LastRegisteredTime = &v
}

// GetLastRxHeartbeatTime returns the LastRxHeartbeatTime field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastRxHeartbeatTime() int32 {
	if o == nil || o.LastRxHeartbeatTime == nil {
		var ret int32
		return ret
	}
	return *o.LastRxHeartbeatTime
}

// GetLastRxHeartbeatTimeOk returns a tuple with the LastRxHeartbeatTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastRxHeartbeatTimeOk() (*int32, bool) {
	if o == nil || o.LastRxHeartbeatTime == nil {
		return nil, false
	}
	return o.LastRxHeartbeatTime, true
}

// HasLastRxHeartbeatTime returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasLastRxHeartbeatTime() bool {
	if o != nil && o.LastRxHeartbeatTime != nil {
		return true
	}

	return false
}

// SetLastRxHeartbeatTime gets a reference to the given int32 and assigns it to the LastRxHeartbeatTime field.
func (o *MsgVpnDistributedCacheClusterInstance) SetLastRxHeartbeatTime(v int32) {
	o.LastRxHeartbeatTime = &v
}

// GetLastRxSetLostMsgTime returns the LastRxSetLostMsgTime field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastRxSetLostMsgTime() int32 {
	if o == nil || o.LastRxSetLostMsgTime == nil {
		var ret int32
		return ret
	}
	return *o.LastRxSetLostMsgTime
}

// GetLastRxSetLostMsgTimeOk returns a tuple with the LastRxSetLostMsgTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastRxSetLostMsgTimeOk() (*int32, bool) {
	if o == nil || o.LastRxSetLostMsgTime == nil {
		return nil, false
	}
	return o.LastRxSetLostMsgTime, true
}

// HasLastRxSetLostMsgTime returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasLastRxSetLostMsgTime() bool {
	if o != nil && o.LastRxSetLostMsgTime != nil {
		return true
	}

	return false
}

// SetLastRxSetLostMsgTime gets a reference to the given int32 and assigns it to the LastRxSetLostMsgTime field.
func (o *MsgVpnDistributedCacheClusterInstance) SetLastRxSetLostMsgTime(v int32) {
	o.LastRxSetLostMsgTime = &v
}

// GetLastStoppedReason returns the LastStoppedReason field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastStoppedReason() string {
	if o == nil || o.LastStoppedReason == nil {
		var ret string
		return ret
	}
	return *o.LastStoppedReason
}

// GetLastStoppedReasonOk returns a tuple with the LastStoppedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastStoppedReasonOk() (*string, bool) {
	if o == nil || o.LastStoppedReason == nil {
		return nil, false
	}
	return o.LastStoppedReason, true
}

// HasLastStoppedReason returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasLastStoppedReason() bool {
	if o != nil && o.LastStoppedReason != nil {
		return true
	}

	return false
}

// SetLastStoppedReason gets a reference to the given string and assigns it to the LastStoppedReason field.
func (o *MsgVpnDistributedCacheClusterInstance) SetLastStoppedReason(v string) {
	o.LastStoppedReason = &v
}

// GetLastStoppedTime returns the LastStoppedTime field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastStoppedTime() int32 {
	if o == nil || o.LastStoppedTime == nil {
		var ret int32
		return ret
	}
	return *o.LastStoppedTime
}

// GetLastStoppedTimeOk returns a tuple with the LastStoppedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastStoppedTimeOk() (*int32, bool) {
	if o == nil || o.LastStoppedTime == nil {
		return nil, false
	}
	return o.LastStoppedTime, true
}

// HasLastStoppedTime returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasLastStoppedTime() bool {
	if o != nil && o.LastStoppedTime != nil {
		return true
	}

	return false
}

// SetLastStoppedTime gets a reference to the given int32 and assigns it to the LastStoppedTime field.
func (o *MsgVpnDistributedCacheClusterInstance) SetLastStoppedTime(v int32) {
	o.LastStoppedTime = &v
}

// GetLastTxClearLostMsgTime returns the LastTxClearLostMsgTime field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastTxClearLostMsgTime() int32 {
	if o == nil || o.LastTxClearLostMsgTime == nil {
		var ret int32
		return ret
	}
	return *o.LastTxClearLostMsgTime
}

// GetLastTxClearLostMsgTimeOk returns a tuple with the LastTxClearLostMsgTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetLastTxClearLostMsgTimeOk() (*int32, bool) {
	if o == nil || o.LastTxClearLostMsgTime == nil {
		return nil, false
	}
	return o.LastTxClearLostMsgTime, true
}

// HasLastTxClearLostMsgTime returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasLastTxClearLostMsgTime() bool {
	if o != nil && o.LastTxClearLostMsgTime != nil {
		return true
	}

	return false
}

// SetLastTxClearLostMsgTime gets a reference to the given int32 and assigns it to the LastTxClearLostMsgTime field.
func (o *MsgVpnDistributedCacheClusterInstance) SetLastTxClearLostMsgTime(v int32) {
	o.LastTxClearLostMsgTime = &v
}

// GetMemoryUsage returns the MemoryUsage field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetMemoryUsage() int32 {
	if o == nil || o.MemoryUsage == nil {
		var ret int32
		return ret
	}
	return *o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetMemoryUsageOk() (*int32, bool) {
	if o == nil || o.MemoryUsage == nil {
		return nil, false
	}
	return o.MemoryUsage, true
}

// HasMemoryUsage returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasMemoryUsage() bool {
	if o != nil && o.MemoryUsage != nil {
		return true
	}

	return false
}

// SetMemoryUsage gets a reference to the given int32 and assigns it to the MemoryUsage field.
func (o *MsgVpnDistributedCacheClusterInstance) SetMemoryUsage(v int32) {
	o.MemoryUsage = &v
}

// GetMsgCount returns the MsgCount field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetMsgCount() int64 {
	if o == nil || o.MsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MsgCount
}

// GetMsgCountOk returns a tuple with the MsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetMsgCountOk() (*int64, bool) {
	if o == nil || o.MsgCount == nil {
		return nil, false
	}
	return o.MsgCount, true
}

// HasMsgCount returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasMsgCount() bool {
	if o != nil && o.MsgCount != nil {
		return true
	}

	return false
}

// SetMsgCount gets a reference to the given int64 and assigns it to the MsgCount field.
func (o *MsgVpnDistributedCacheClusterInstance) SetMsgCount(v int64) {
	o.MsgCount = &v
}

// GetMsgPeakCount returns the MsgPeakCount field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetMsgPeakCount() int64 {
	if o == nil || o.MsgPeakCount == nil {
		var ret int64
		return ret
	}
	return *o.MsgPeakCount
}

// GetMsgPeakCountOk returns a tuple with the MsgPeakCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetMsgPeakCountOk() (*int64, bool) {
	if o == nil || o.MsgPeakCount == nil {
		return nil, false
	}
	return o.MsgPeakCount, true
}

// HasMsgPeakCount returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasMsgPeakCount() bool {
	if o != nil && o.MsgPeakCount != nil {
		return true
	}

	return false
}

// SetMsgPeakCount gets a reference to the given int64 and assigns it to the MsgPeakCount field.
func (o *MsgVpnDistributedCacheClusterInstance) SetMsgPeakCount(v int64) {
	o.MsgPeakCount = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnDistributedCacheClusterInstance) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetMsgsLost returns the MsgsLost field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetMsgsLost() bool {
	if o == nil || o.MsgsLost == nil {
		var ret bool
		return ret
	}
	return *o.MsgsLost
}

// GetMsgsLostOk returns a tuple with the MsgsLost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetMsgsLostOk() (*bool, bool) {
	if o == nil || o.MsgsLost == nil {
		return nil, false
	}
	return o.MsgsLost, true
}

// HasMsgsLost returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasMsgsLost() bool {
	if o != nil && o.MsgsLost != nil {
		return true
	}

	return false
}

// SetMsgsLost gets a reference to the given bool and assigns it to the MsgsLost field.
func (o *MsgVpnDistributedCacheClusterInstance) SetMsgsLost(v bool) {
	o.MsgsLost = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetRate() MsgVpnDistributedCacheClusterInstanceRate {
	if o == nil || o.Rate == nil {
		var ret MsgVpnDistributedCacheClusterInstanceRate
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetRateOk() (*MsgVpnDistributedCacheClusterInstanceRate, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given MsgVpnDistributedCacheClusterInstanceRate and assigns it to the Rate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetRate(v MsgVpnDistributedCacheClusterInstanceRate) {
	o.Rate = &v
}

// GetRequestQueueDepthCount returns the RequestQueueDepthCount field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetRequestQueueDepthCount() int64 {
	if o == nil || o.RequestQueueDepthCount == nil {
		var ret int64
		return ret
	}
	return *o.RequestQueueDepthCount
}

// GetRequestQueueDepthCountOk returns a tuple with the RequestQueueDepthCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetRequestQueueDepthCountOk() (*int64, bool) {
	if o == nil || o.RequestQueueDepthCount == nil {
		return nil, false
	}
	return o.RequestQueueDepthCount, true
}

// HasRequestQueueDepthCount returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasRequestQueueDepthCount() bool {
	if o != nil && o.RequestQueueDepthCount != nil {
		return true
	}

	return false
}

// SetRequestQueueDepthCount gets a reference to the given int64 and assigns it to the RequestQueueDepthCount field.
func (o *MsgVpnDistributedCacheClusterInstance) SetRequestQueueDepthCount(v int64) {
	o.RequestQueueDepthCount = &v
}

// GetRequestQueueDepthPeakCount returns the RequestQueueDepthPeakCount field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetRequestQueueDepthPeakCount() int64 {
	if o == nil || o.RequestQueueDepthPeakCount == nil {
		var ret int64
		return ret
	}
	return *o.RequestQueueDepthPeakCount
}

// GetRequestQueueDepthPeakCountOk returns a tuple with the RequestQueueDepthPeakCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetRequestQueueDepthPeakCountOk() (*int64, bool) {
	if o == nil || o.RequestQueueDepthPeakCount == nil {
		return nil, false
	}
	return o.RequestQueueDepthPeakCount, true
}

// HasRequestQueueDepthPeakCount returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasRequestQueueDepthPeakCount() bool {
	if o != nil && o.RequestQueueDepthPeakCount != nil {
		return true
	}

	return false
}

// SetRequestQueueDepthPeakCount gets a reference to the given int64 and assigns it to the RequestQueueDepthPeakCount field.
func (o *MsgVpnDistributedCacheClusterInstance) SetRequestQueueDepthPeakCount(v int64) {
	o.RequestQueueDepthPeakCount = &v
}

// GetRequestRxPeakRate returns the RequestRxPeakRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetRequestRxPeakRate() int64 {
	if o == nil || o.RequestRxPeakRate == nil {
		var ret int64
		return ret
	}
	return *o.RequestRxPeakRate
}

// GetRequestRxPeakRateOk returns a tuple with the RequestRxPeakRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetRequestRxPeakRateOk() (*int64, bool) {
	if o == nil || o.RequestRxPeakRate == nil {
		return nil, false
	}
	return o.RequestRxPeakRate, true
}

// HasRequestRxPeakRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasRequestRxPeakRate() bool {
	if o != nil && o.RequestRxPeakRate != nil {
		return true
	}

	return false
}

// SetRequestRxPeakRate gets a reference to the given int64 and assigns it to the RequestRxPeakRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetRequestRxPeakRate(v int64) {
	o.RequestRxPeakRate = &v
}

// GetRequestRxRate returns the RequestRxRate field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetRequestRxRate() int64 {
	if o == nil || o.RequestRxRate == nil {
		var ret int64
		return ret
	}
	return *o.RequestRxRate
}

// GetRequestRxRateOk returns a tuple with the RequestRxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetRequestRxRateOk() (*int64, bool) {
	if o == nil || o.RequestRxRate == nil {
		return nil, false
	}
	return o.RequestRxRate, true
}

// HasRequestRxRate returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasRequestRxRate() bool {
	if o != nil && o.RequestRxRate != nil {
		return true
	}

	return false
}

// SetRequestRxRate gets a reference to the given int64 and assigns it to the RequestRxRate field.
func (o *MsgVpnDistributedCacheClusterInstance) SetRequestRxRate(v int64) {
	o.RequestRxRate = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MsgVpnDistributedCacheClusterInstance) SetState(v string) {
	o.State = &v
}

// GetStopOnLostMsgEnabled returns the StopOnLostMsgEnabled field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetStopOnLostMsgEnabled() bool {
	if o == nil || o.StopOnLostMsgEnabled == nil {
		var ret bool
		return ret
	}
	return *o.StopOnLostMsgEnabled
}

// GetStopOnLostMsgEnabledOk returns a tuple with the StopOnLostMsgEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetStopOnLostMsgEnabledOk() (*bool, bool) {
	if o == nil || o.StopOnLostMsgEnabled == nil {
		return nil, false
	}
	return o.StopOnLostMsgEnabled, true
}

// HasStopOnLostMsgEnabled returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasStopOnLostMsgEnabled() bool {
	if o != nil && o.StopOnLostMsgEnabled != nil {
		return true
	}

	return false
}

// SetStopOnLostMsgEnabled gets a reference to the given bool and assigns it to the StopOnLostMsgEnabled field.
func (o *MsgVpnDistributedCacheClusterInstance) SetStopOnLostMsgEnabled(v bool) {
	o.StopOnLostMsgEnabled = &v
}

// GetTopicCount returns the TopicCount field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetTopicCount() int64 {
	if o == nil || o.TopicCount == nil {
		var ret int64
		return ret
	}
	return *o.TopicCount
}

// GetTopicCountOk returns a tuple with the TopicCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetTopicCountOk() (*int64, bool) {
	if o == nil || o.TopicCount == nil {
		return nil, false
	}
	return o.TopicCount, true
}

// HasTopicCount returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasTopicCount() bool {
	if o != nil && o.TopicCount != nil {
		return true
	}

	return false
}

// SetTopicCount gets a reference to the given int64 and assigns it to the TopicCount field.
func (o *MsgVpnDistributedCacheClusterInstance) SetTopicCount(v int64) {
	o.TopicCount = &v
}

// GetTopicPeakCount returns the TopicPeakCount field value if set, zero value otherwise.
func (o *MsgVpnDistributedCacheClusterInstance) GetTopicPeakCount() int64 {
	if o == nil || o.TopicPeakCount == nil {
		var ret int64
		return ret
	}
	return *o.TopicPeakCount
}

// GetTopicPeakCountOk returns a tuple with the TopicPeakCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnDistributedCacheClusterInstance) GetTopicPeakCountOk() (*int64, bool) {
	if o == nil || o.TopicPeakCount == nil {
		return nil, false
	}
	return o.TopicPeakCount, true
}

// HasTopicPeakCount returns a boolean if a field has been set.
func (o *MsgVpnDistributedCacheClusterInstance) HasTopicPeakCount() bool {
	if o != nil && o.TopicPeakCount != nil {
		return true
	}

	return false
}

// SetTopicPeakCount gets a reference to the given int64 and assigns it to the TopicPeakCount field.
func (o *MsgVpnDistributedCacheClusterInstance) SetTopicPeakCount(v int64) {
	o.TopicPeakCount = &v
}

func (o MsgVpnDistributedCacheClusterInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoStartEnabled != nil {
		toSerialize["autoStartEnabled"] = o.AutoStartEnabled
	}
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
	if o.CacheName != nil {
		toSerialize["cacheName"] = o.CacheName
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.Counter != nil {
		toSerialize["counter"] = o.Counter
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
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}
	if o.LastRegisteredTime != nil {
		toSerialize["lastRegisteredTime"] = o.LastRegisteredTime
	}
	if o.LastRxHeartbeatTime != nil {
		toSerialize["lastRxHeartbeatTime"] = o.LastRxHeartbeatTime
	}
	if o.LastRxSetLostMsgTime != nil {
		toSerialize["lastRxSetLostMsgTime"] = o.LastRxSetLostMsgTime
	}
	if o.LastStoppedReason != nil {
		toSerialize["lastStoppedReason"] = o.LastStoppedReason
	}
	if o.LastStoppedTime != nil {
		toSerialize["lastStoppedTime"] = o.LastStoppedTime
	}
	if o.LastTxClearLostMsgTime != nil {
		toSerialize["lastTxClearLostMsgTime"] = o.LastTxClearLostMsgTime
	}
	if o.MemoryUsage != nil {
		toSerialize["memoryUsage"] = o.MemoryUsage
	}
	if o.MsgCount != nil {
		toSerialize["msgCount"] = o.MsgCount
	}
	if o.MsgPeakCount != nil {
		toSerialize["msgPeakCount"] = o.MsgPeakCount
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.MsgsLost != nil {
		toSerialize["msgsLost"] = o.MsgsLost
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.RequestQueueDepthCount != nil {
		toSerialize["requestQueueDepthCount"] = o.RequestQueueDepthCount
	}
	if o.RequestQueueDepthPeakCount != nil {
		toSerialize["requestQueueDepthPeakCount"] = o.RequestQueueDepthPeakCount
	}
	if o.RequestRxPeakRate != nil {
		toSerialize["requestRxPeakRate"] = o.RequestRxPeakRate
	}
	if o.RequestRxRate != nil {
		toSerialize["requestRxRate"] = o.RequestRxRate
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StopOnLostMsgEnabled != nil {
		toSerialize["stopOnLostMsgEnabled"] = o.StopOnLostMsgEnabled
	}
	if o.TopicCount != nil {
		toSerialize["topicCount"] = o.TopicCount
	}
	if o.TopicPeakCount != nil {
		toSerialize["topicPeakCount"] = o.TopicPeakCount
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnDistributedCacheClusterInstance struct {
	value *MsgVpnDistributedCacheClusterInstance
	isSet bool
}

func (v NullableMsgVpnDistributedCacheClusterInstance) Get() *MsgVpnDistributedCacheClusterInstance {
	return v.value
}

func (v *NullableMsgVpnDistributedCacheClusterInstance) Set(val *MsgVpnDistributedCacheClusterInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnDistributedCacheClusterInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnDistributedCacheClusterInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnDistributedCacheClusterInstance(val *MsgVpnDistributedCacheClusterInstance) *NullableMsgVpnDistributedCacheClusterInstance {
	return &NullableMsgVpnDistributedCacheClusterInstance{value: val, isSet: true}
}

func (v NullableMsgVpnDistributedCacheClusterInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnDistributedCacheClusterInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
