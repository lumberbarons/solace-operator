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

// MsgVpnMqttRetainCache struct for MsgVpnMqttRetainCache
type MsgVpnMqttRetainCache struct {
	// The name of the backup Cache Instance associated with this MQTT Retain Cache.
	BackupCacheInstance *string `json:"backupCacheInstance,omitempty"`
	// The reason why the backup cache associated with this MQTT Retain Cache is operationally down, if any.
	BackupFailureReason *string `json:"backupFailureReason,omitempty"`
	// Indicates whether the backup cache associated with this MQTT Retain Cache is operationally up.
	BackupUp *bool `json:"backupUp,omitempty"`
	// The number of seconds that the backup cache associated with this MQTT Retain Cache has been operationally up.
	BackupUptime *int32 `json:"backupUptime,omitempty"`
	// The name of the Cache Cluster associated with this MQTT Retain Cache.
	CacheCluster *string `json:"cacheCluster,omitempty"`
	// The name of the MQTT Retain Cache.
	CacheName *string `json:"cacheName,omitempty"`
	// The name of the Distributed Cache associated with this MQTT Retain Cache.
	DistributedCache *string `json:"distributedCache,omitempty"`
	// Indicates whether this MQTT Retain Cache is enabled. When the cache is disabled, neither retain messages nor retain requests will be delivered by the cache. However, live retain messages will continue to be delivered to currently connected MQTT clients.
	Enabled *bool `json:"enabled,omitempty"`
	// The reason why this MQTT Retain Cache is operationally down, if any.
	FailureReason *string `json:"failureReason,omitempty"`
	// The message lifetime, in seconds. If a message remains cached for the duration of its lifetime, the cache will remove the message. A lifetime of 0 results in the message being retained indefinitely.
	MsgLifetime *int64 `json:"msgLifetime,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// The name of the primary Cache Instance associated with this MQTT Retain Cache.
	PrimaryCacheInstance *string `json:"primaryCacheInstance,omitempty"`
	// The reason why the primary cache associated with this MQTT Retain Cache is operationally down, if any.
	PrimaryFailureReason *string `json:"primaryFailureReason,omitempty"`
	// Indicates whether the primary cache associated with this MQTT Retain Cache is operationally up.
	PrimaryUp *bool `json:"primaryUp,omitempty"`
	// The number of seconds that the primary cache associated with this MQTT Retain Cache has been operationally up.
	PrimaryUptime *int32 `json:"primaryUptime,omitempty"`
	// Indicates whether this MQTT Retain Cache is operationally up.
	Up *bool `json:"up,omitempty"`
	// The number of seconds that the MQTT Retain Cache has been operationally up.
	Uptime *int32 `json:"uptime,omitempty"`
}

// NewMsgVpnMqttRetainCache instantiates a new MsgVpnMqttRetainCache object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnMqttRetainCache() *MsgVpnMqttRetainCache {
	this := MsgVpnMqttRetainCache{}
	return &this
}

// NewMsgVpnMqttRetainCacheWithDefaults instantiates a new MsgVpnMqttRetainCache object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnMqttRetainCacheWithDefaults() *MsgVpnMqttRetainCache {
	this := MsgVpnMqttRetainCache{}
	return &this
}

// GetBackupCacheInstance returns the BackupCacheInstance field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetBackupCacheInstance() string {
	if o == nil || o.BackupCacheInstance == nil {
		var ret string
		return ret
	}
	return *o.BackupCacheInstance
}

// GetBackupCacheInstanceOk returns a tuple with the BackupCacheInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetBackupCacheInstanceOk() (*string, bool) {
	if o == nil || o.BackupCacheInstance == nil {
		return nil, false
	}
	return o.BackupCacheInstance, true
}

// HasBackupCacheInstance returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasBackupCacheInstance() bool {
	if o != nil && o.BackupCacheInstance != nil {
		return true
	}

	return false
}

// SetBackupCacheInstance gets a reference to the given string and assigns it to the BackupCacheInstance field.
func (o *MsgVpnMqttRetainCache) SetBackupCacheInstance(v string) {
	o.BackupCacheInstance = &v
}

// GetBackupFailureReason returns the BackupFailureReason field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetBackupFailureReason() string {
	if o == nil || o.BackupFailureReason == nil {
		var ret string
		return ret
	}
	return *o.BackupFailureReason
}

// GetBackupFailureReasonOk returns a tuple with the BackupFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetBackupFailureReasonOk() (*string, bool) {
	if o == nil || o.BackupFailureReason == nil {
		return nil, false
	}
	return o.BackupFailureReason, true
}

// HasBackupFailureReason returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasBackupFailureReason() bool {
	if o != nil && o.BackupFailureReason != nil {
		return true
	}

	return false
}

// SetBackupFailureReason gets a reference to the given string and assigns it to the BackupFailureReason field.
func (o *MsgVpnMqttRetainCache) SetBackupFailureReason(v string) {
	o.BackupFailureReason = &v
}

// GetBackupUp returns the BackupUp field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetBackupUp() bool {
	if o == nil || o.BackupUp == nil {
		var ret bool
		return ret
	}
	return *o.BackupUp
}

// GetBackupUpOk returns a tuple with the BackupUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetBackupUpOk() (*bool, bool) {
	if o == nil || o.BackupUp == nil {
		return nil, false
	}
	return o.BackupUp, true
}

// HasBackupUp returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasBackupUp() bool {
	if o != nil && o.BackupUp != nil {
		return true
	}

	return false
}

// SetBackupUp gets a reference to the given bool and assigns it to the BackupUp field.
func (o *MsgVpnMqttRetainCache) SetBackupUp(v bool) {
	o.BackupUp = &v
}

// GetBackupUptime returns the BackupUptime field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetBackupUptime() int32 {
	if o == nil || o.BackupUptime == nil {
		var ret int32
		return ret
	}
	return *o.BackupUptime
}

// GetBackupUptimeOk returns a tuple with the BackupUptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetBackupUptimeOk() (*int32, bool) {
	if o == nil || o.BackupUptime == nil {
		return nil, false
	}
	return o.BackupUptime, true
}

// HasBackupUptime returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasBackupUptime() bool {
	if o != nil && o.BackupUptime != nil {
		return true
	}

	return false
}

// SetBackupUptime gets a reference to the given int32 and assigns it to the BackupUptime field.
func (o *MsgVpnMqttRetainCache) SetBackupUptime(v int32) {
	o.BackupUptime = &v
}

// GetCacheCluster returns the CacheCluster field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetCacheCluster() string {
	if o == nil || o.CacheCluster == nil {
		var ret string
		return ret
	}
	return *o.CacheCluster
}

// GetCacheClusterOk returns a tuple with the CacheCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetCacheClusterOk() (*string, bool) {
	if o == nil || o.CacheCluster == nil {
		return nil, false
	}
	return o.CacheCluster, true
}

// HasCacheCluster returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasCacheCluster() bool {
	if o != nil && o.CacheCluster != nil {
		return true
	}

	return false
}

// SetCacheCluster gets a reference to the given string and assigns it to the CacheCluster field.
func (o *MsgVpnMqttRetainCache) SetCacheCluster(v string) {
	o.CacheCluster = &v
}

// GetCacheName returns the CacheName field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetCacheName() string {
	if o == nil || o.CacheName == nil {
		var ret string
		return ret
	}
	return *o.CacheName
}

// GetCacheNameOk returns a tuple with the CacheName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetCacheNameOk() (*string, bool) {
	if o == nil || o.CacheName == nil {
		return nil, false
	}
	return o.CacheName, true
}

// HasCacheName returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasCacheName() bool {
	if o != nil && o.CacheName != nil {
		return true
	}

	return false
}

// SetCacheName gets a reference to the given string and assigns it to the CacheName field.
func (o *MsgVpnMqttRetainCache) SetCacheName(v string) {
	o.CacheName = &v
}

// GetDistributedCache returns the DistributedCache field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetDistributedCache() string {
	if o == nil || o.DistributedCache == nil {
		var ret string
		return ret
	}
	return *o.DistributedCache
}

// GetDistributedCacheOk returns a tuple with the DistributedCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetDistributedCacheOk() (*string, bool) {
	if o == nil || o.DistributedCache == nil {
		return nil, false
	}
	return o.DistributedCache, true
}

// HasDistributedCache returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasDistributedCache() bool {
	if o != nil && o.DistributedCache != nil {
		return true
	}

	return false
}

// SetDistributedCache gets a reference to the given string and assigns it to the DistributedCache field.
func (o *MsgVpnMqttRetainCache) SetDistributedCache(v string) {
	o.DistributedCache = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MsgVpnMqttRetainCache) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *MsgVpnMqttRetainCache) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetMsgLifetime returns the MsgLifetime field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetMsgLifetime() int64 {
	if o == nil || o.MsgLifetime == nil {
		var ret int64
		return ret
	}
	return *o.MsgLifetime
}

// GetMsgLifetimeOk returns a tuple with the MsgLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetMsgLifetimeOk() (*int64, bool) {
	if o == nil || o.MsgLifetime == nil {
		return nil, false
	}
	return o.MsgLifetime, true
}

// HasMsgLifetime returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasMsgLifetime() bool {
	if o != nil && o.MsgLifetime != nil {
		return true
	}

	return false
}

// SetMsgLifetime gets a reference to the given int64 and assigns it to the MsgLifetime field.
func (o *MsgVpnMqttRetainCache) SetMsgLifetime(v int64) {
	o.MsgLifetime = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnMqttRetainCache) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetPrimaryCacheInstance returns the PrimaryCacheInstance field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetPrimaryCacheInstance() string {
	if o == nil || o.PrimaryCacheInstance == nil {
		var ret string
		return ret
	}
	return *o.PrimaryCacheInstance
}

// GetPrimaryCacheInstanceOk returns a tuple with the PrimaryCacheInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetPrimaryCacheInstanceOk() (*string, bool) {
	if o == nil || o.PrimaryCacheInstance == nil {
		return nil, false
	}
	return o.PrimaryCacheInstance, true
}

// HasPrimaryCacheInstance returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasPrimaryCacheInstance() bool {
	if o != nil && o.PrimaryCacheInstance != nil {
		return true
	}

	return false
}

// SetPrimaryCacheInstance gets a reference to the given string and assigns it to the PrimaryCacheInstance field.
func (o *MsgVpnMqttRetainCache) SetPrimaryCacheInstance(v string) {
	o.PrimaryCacheInstance = &v
}

// GetPrimaryFailureReason returns the PrimaryFailureReason field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetPrimaryFailureReason() string {
	if o == nil || o.PrimaryFailureReason == nil {
		var ret string
		return ret
	}
	return *o.PrimaryFailureReason
}

// GetPrimaryFailureReasonOk returns a tuple with the PrimaryFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetPrimaryFailureReasonOk() (*string, bool) {
	if o == nil || o.PrimaryFailureReason == nil {
		return nil, false
	}
	return o.PrimaryFailureReason, true
}

// HasPrimaryFailureReason returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasPrimaryFailureReason() bool {
	if o != nil && o.PrimaryFailureReason != nil {
		return true
	}

	return false
}

// SetPrimaryFailureReason gets a reference to the given string and assigns it to the PrimaryFailureReason field.
func (o *MsgVpnMqttRetainCache) SetPrimaryFailureReason(v string) {
	o.PrimaryFailureReason = &v
}

// GetPrimaryUp returns the PrimaryUp field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetPrimaryUp() bool {
	if o == nil || o.PrimaryUp == nil {
		var ret bool
		return ret
	}
	return *o.PrimaryUp
}

// GetPrimaryUpOk returns a tuple with the PrimaryUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetPrimaryUpOk() (*bool, bool) {
	if o == nil || o.PrimaryUp == nil {
		return nil, false
	}
	return o.PrimaryUp, true
}

// HasPrimaryUp returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasPrimaryUp() bool {
	if o != nil && o.PrimaryUp != nil {
		return true
	}

	return false
}

// SetPrimaryUp gets a reference to the given bool and assigns it to the PrimaryUp field.
func (o *MsgVpnMqttRetainCache) SetPrimaryUp(v bool) {
	o.PrimaryUp = &v
}

// GetPrimaryUptime returns the PrimaryUptime field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetPrimaryUptime() int32 {
	if o == nil || o.PrimaryUptime == nil {
		var ret int32
		return ret
	}
	return *o.PrimaryUptime
}

// GetPrimaryUptimeOk returns a tuple with the PrimaryUptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetPrimaryUptimeOk() (*int32, bool) {
	if o == nil || o.PrimaryUptime == nil {
		return nil, false
	}
	return o.PrimaryUptime, true
}

// HasPrimaryUptime returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasPrimaryUptime() bool {
	if o != nil && o.PrimaryUptime != nil {
		return true
	}

	return false
}

// SetPrimaryUptime gets a reference to the given int32 and assigns it to the PrimaryUptime field.
func (o *MsgVpnMqttRetainCache) SetPrimaryUptime(v int32) {
	o.PrimaryUptime = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetUp() bool {
	if o == nil || o.Up == nil {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetUpOk() (*bool, bool) {
	if o == nil || o.Up == nil {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasUp() bool {
	if o != nil && o.Up != nil {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *MsgVpnMqttRetainCache) SetUp(v bool) {
	o.Up = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *MsgVpnMqttRetainCache) GetUptime() int32 {
	if o == nil || o.Uptime == nil {
		var ret int32
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnMqttRetainCache) GetUptimeOk() (*int32, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *MsgVpnMqttRetainCache) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given int32 and assigns it to the Uptime field.
func (o *MsgVpnMqttRetainCache) SetUptime(v int32) {
	o.Uptime = &v
}

func (o MsgVpnMqttRetainCache) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupCacheInstance != nil {
		toSerialize["backupCacheInstance"] = o.BackupCacheInstance
	}
	if o.BackupFailureReason != nil {
		toSerialize["backupFailureReason"] = o.BackupFailureReason
	}
	if o.BackupUp != nil {
		toSerialize["backupUp"] = o.BackupUp
	}
	if o.BackupUptime != nil {
		toSerialize["backupUptime"] = o.BackupUptime
	}
	if o.CacheCluster != nil {
		toSerialize["cacheCluster"] = o.CacheCluster
	}
	if o.CacheName != nil {
		toSerialize["cacheName"] = o.CacheName
	}
	if o.DistributedCache != nil {
		toSerialize["distributedCache"] = o.DistributedCache
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FailureReason != nil {
		toSerialize["failureReason"] = o.FailureReason
	}
	if o.MsgLifetime != nil {
		toSerialize["msgLifetime"] = o.MsgLifetime
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.PrimaryCacheInstance != nil {
		toSerialize["primaryCacheInstance"] = o.PrimaryCacheInstance
	}
	if o.PrimaryFailureReason != nil {
		toSerialize["primaryFailureReason"] = o.PrimaryFailureReason
	}
	if o.PrimaryUp != nil {
		toSerialize["primaryUp"] = o.PrimaryUp
	}
	if o.PrimaryUptime != nil {
		toSerialize["primaryUptime"] = o.PrimaryUptime
	}
	if o.Up != nil {
		toSerialize["up"] = o.Up
	}
	if o.Uptime != nil {
		toSerialize["uptime"] = o.Uptime
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnMqttRetainCache struct {
	value *MsgVpnMqttRetainCache
	isSet bool
}

func (v NullableMsgVpnMqttRetainCache) Get() *MsgVpnMqttRetainCache {
	return v.value
}

func (v *NullableMsgVpnMqttRetainCache) Set(val *MsgVpnMqttRetainCache) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnMqttRetainCache) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnMqttRetainCache) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnMqttRetainCache(val *MsgVpnMqttRetainCache) *NullableMsgVpnMqttRetainCache {
	return &NullableMsgVpnMqttRetainCache{value: val, isSet: true}
}

func (v NullableMsgVpnMqttRetainCache) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnMqttRetainCache) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
