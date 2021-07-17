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

// MsgVpnCollections struct for MsgVpnCollections
type MsgVpnCollections struct {
	AclProfiles                  *MsgVpnCollectionsAclprofiles                  `json:"aclProfiles,omitempty"`
	AuthenticationOauthProviders *MsgVpnCollectionsAuthenticationoauthproviders `json:"authenticationOauthProviders,omitempty"`
	AuthorizationGroups          *MsgVpnCollectionsAuthorizationgroups          `json:"authorizationGroups,omitempty"`
	Bridges                      *MsgVpnCollectionsBridges                      `json:"bridges,omitempty"`
	ClientProfiles               *MsgVpnCollectionsClientprofiles               `json:"clientProfiles,omitempty"`
	ClientUsernames              *MsgVpnCollectionsClientusernames              `json:"clientUsernames,omitempty"`
	Clients                      *MsgVpnCollectionsClients                      `json:"clients,omitempty"`
	ConfigSyncRemoteNodes        *MsgVpnCollectionsConfigsyncremotenodes        `json:"configSyncRemoteNodes,omitempty"`
	DistributedCaches            *MsgVpnCollectionsDistributedcaches            `json:"distributedCaches,omitempty"`
	DmrBridges                   *MsgVpnCollectionsDmrbridges                   `json:"dmrBridges,omitempty"`
	JndiConnectionFactories      *MsgVpnCollectionsJndiconnectionfactories      `json:"jndiConnectionFactories,omitempty"`
	JndiQueues                   *MsgVpnCollectionsJndiqueues                   `json:"jndiQueues,omitempty"`
	JndiTopics                   *MsgVpnCollectionsJnditopics                   `json:"jndiTopics,omitempty"`
	MqttRetainCaches             *MsgVpnCollectionsMqttretaincaches             `json:"mqttRetainCaches,omitempty"`
	MqttSessions                 *MsgVpnCollectionsMqttsessions                 `json:"mqttSessions,omitempty"`
	QueueTemplates               *MsgVpnCollectionsQueuetemplates               `json:"queueTemplates,omitempty"`
	Queues                       *MsgVpnCollectionsQueues                       `json:"queues,omitempty"`
	ReplayLogs                   *MsgVpnCollectionsReplaylogs                   `json:"replayLogs,omitempty"`
	ReplicatedTopics             *MsgVpnCollectionsReplicatedtopics             `json:"replicatedTopics,omitempty"`
	RestDeliveryPoints           *MsgVpnCollectionsRestdeliverypoints           `json:"restDeliveryPoints,omitempty"`
	TopicEndpointTemplates       *MsgVpnCollectionsTopicendpointtemplates       `json:"topicEndpointTemplates,omitempty"`
	TopicEndpoints               *MsgVpnCollectionsTopicendpoints               `json:"topicEndpoints,omitempty"`
	Transactions                 *MsgVpnCollectionsTransactions                 `json:"transactions,omitempty"`
}

// NewMsgVpnCollections instantiates a new MsgVpnCollections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnCollections() *MsgVpnCollections {
	this := MsgVpnCollections{}
	return &this
}

// NewMsgVpnCollectionsWithDefaults instantiates a new MsgVpnCollections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnCollectionsWithDefaults() *MsgVpnCollections {
	this := MsgVpnCollections{}
	return &this
}

// GetAclProfiles returns the AclProfiles field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetAclProfiles() MsgVpnCollectionsAclprofiles {
	if o == nil || o.AclProfiles == nil {
		var ret MsgVpnCollectionsAclprofiles
		return ret
	}
	return *o.AclProfiles
}

// GetAclProfilesOk returns a tuple with the AclProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetAclProfilesOk() (*MsgVpnCollectionsAclprofiles, bool) {
	if o == nil || o.AclProfiles == nil {
		return nil, false
	}
	return o.AclProfiles, true
}

// HasAclProfiles returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasAclProfiles() bool {
	if o != nil && o.AclProfiles != nil {
		return true
	}

	return false
}

// SetAclProfiles gets a reference to the given MsgVpnCollectionsAclprofiles and assigns it to the AclProfiles field.
func (o *MsgVpnCollections) SetAclProfiles(v MsgVpnCollectionsAclprofiles) {
	o.AclProfiles = &v
}

// GetAuthenticationOauthProviders returns the AuthenticationOauthProviders field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetAuthenticationOauthProviders() MsgVpnCollectionsAuthenticationoauthproviders {
	if o == nil || o.AuthenticationOauthProviders == nil {
		var ret MsgVpnCollectionsAuthenticationoauthproviders
		return ret
	}
	return *o.AuthenticationOauthProviders
}

// GetAuthenticationOauthProvidersOk returns a tuple with the AuthenticationOauthProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetAuthenticationOauthProvidersOk() (*MsgVpnCollectionsAuthenticationoauthproviders, bool) {
	if o == nil || o.AuthenticationOauthProviders == nil {
		return nil, false
	}
	return o.AuthenticationOauthProviders, true
}

// HasAuthenticationOauthProviders returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasAuthenticationOauthProviders() bool {
	if o != nil && o.AuthenticationOauthProviders != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthProviders gets a reference to the given MsgVpnCollectionsAuthenticationoauthproviders and assigns it to the AuthenticationOauthProviders field.
func (o *MsgVpnCollections) SetAuthenticationOauthProviders(v MsgVpnCollectionsAuthenticationoauthproviders) {
	o.AuthenticationOauthProviders = &v
}

// GetAuthorizationGroups returns the AuthorizationGroups field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetAuthorizationGroups() MsgVpnCollectionsAuthorizationgroups {
	if o == nil || o.AuthorizationGroups == nil {
		var ret MsgVpnCollectionsAuthorizationgroups
		return ret
	}
	return *o.AuthorizationGroups
}

// GetAuthorizationGroupsOk returns a tuple with the AuthorizationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetAuthorizationGroupsOk() (*MsgVpnCollectionsAuthorizationgroups, bool) {
	if o == nil || o.AuthorizationGroups == nil {
		return nil, false
	}
	return o.AuthorizationGroups, true
}

// HasAuthorizationGroups returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasAuthorizationGroups() bool {
	if o != nil && o.AuthorizationGroups != nil {
		return true
	}

	return false
}

// SetAuthorizationGroups gets a reference to the given MsgVpnCollectionsAuthorizationgroups and assigns it to the AuthorizationGroups field.
func (o *MsgVpnCollections) SetAuthorizationGroups(v MsgVpnCollectionsAuthorizationgroups) {
	o.AuthorizationGroups = &v
}

// GetBridges returns the Bridges field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetBridges() MsgVpnCollectionsBridges {
	if o == nil || o.Bridges == nil {
		var ret MsgVpnCollectionsBridges
		return ret
	}
	return *o.Bridges
}

// GetBridgesOk returns a tuple with the Bridges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetBridgesOk() (*MsgVpnCollectionsBridges, bool) {
	if o == nil || o.Bridges == nil {
		return nil, false
	}
	return o.Bridges, true
}

// HasBridges returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasBridges() bool {
	if o != nil && o.Bridges != nil {
		return true
	}

	return false
}

// SetBridges gets a reference to the given MsgVpnCollectionsBridges and assigns it to the Bridges field.
func (o *MsgVpnCollections) SetBridges(v MsgVpnCollectionsBridges) {
	o.Bridges = &v
}

// GetClientProfiles returns the ClientProfiles field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetClientProfiles() MsgVpnCollectionsClientprofiles {
	if o == nil || o.ClientProfiles == nil {
		var ret MsgVpnCollectionsClientprofiles
		return ret
	}
	return *o.ClientProfiles
}

// GetClientProfilesOk returns a tuple with the ClientProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetClientProfilesOk() (*MsgVpnCollectionsClientprofiles, bool) {
	if o == nil || o.ClientProfiles == nil {
		return nil, false
	}
	return o.ClientProfiles, true
}

// HasClientProfiles returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasClientProfiles() bool {
	if o != nil && o.ClientProfiles != nil {
		return true
	}

	return false
}

// SetClientProfiles gets a reference to the given MsgVpnCollectionsClientprofiles and assigns it to the ClientProfiles field.
func (o *MsgVpnCollections) SetClientProfiles(v MsgVpnCollectionsClientprofiles) {
	o.ClientProfiles = &v
}

// GetClientUsernames returns the ClientUsernames field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetClientUsernames() MsgVpnCollectionsClientusernames {
	if o == nil || o.ClientUsernames == nil {
		var ret MsgVpnCollectionsClientusernames
		return ret
	}
	return *o.ClientUsernames
}

// GetClientUsernamesOk returns a tuple with the ClientUsernames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetClientUsernamesOk() (*MsgVpnCollectionsClientusernames, bool) {
	if o == nil || o.ClientUsernames == nil {
		return nil, false
	}
	return o.ClientUsernames, true
}

// HasClientUsernames returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasClientUsernames() bool {
	if o != nil && o.ClientUsernames != nil {
		return true
	}

	return false
}

// SetClientUsernames gets a reference to the given MsgVpnCollectionsClientusernames and assigns it to the ClientUsernames field.
func (o *MsgVpnCollections) SetClientUsernames(v MsgVpnCollectionsClientusernames) {
	o.ClientUsernames = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetClients() MsgVpnCollectionsClients {
	if o == nil || o.Clients == nil {
		var ret MsgVpnCollectionsClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetClientsOk() (*MsgVpnCollectionsClients, bool) {
	if o == nil || o.Clients == nil {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasClients() bool {
	if o != nil && o.Clients != nil {
		return true
	}

	return false
}

// SetClients gets a reference to the given MsgVpnCollectionsClients and assigns it to the Clients field.
func (o *MsgVpnCollections) SetClients(v MsgVpnCollectionsClients) {
	o.Clients = &v
}

// GetConfigSyncRemoteNodes returns the ConfigSyncRemoteNodes field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetConfigSyncRemoteNodes() MsgVpnCollectionsConfigsyncremotenodes {
	if o == nil || o.ConfigSyncRemoteNodes == nil {
		var ret MsgVpnCollectionsConfigsyncremotenodes
		return ret
	}
	return *o.ConfigSyncRemoteNodes
}

// GetConfigSyncRemoteNodesOk returns a tuple with the ConfigSyncRemoteNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetConfigSyncRemoteNodesOk() (*MsgVpnCollectionsConfigsyncremotenodes, bool) {
	if o == nil || o.ConfigSyncRemoteNodes == nil {
		return nil, false
	}
	return o.ConfigSyncRemoteNodes, true
}

// HasConfigSyncRemoteNodes returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasConfigSyncRemoteNodes() bool {
	if o != nil && o.ConfigSyncRemoteNodes != nil {
		return true
	}

	return false
}

// SetConfigSyncRemoteNodes gets a reference to the given MsgVpnCollectionsConfigsyncremotenodes and assigns it to the ConfigSyncRemoteNodes field.
func (o *MsgVpnCollections) SetConfigSyncRemoteNodes(v MsgVpnCollectionsConfigsyncremotenodes) {
	o.ConfigSyncRemoteNodes = &v
}

// GetDistributedCaches returns the DistributedCaches field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetDistributedCaches() MsgVpnCollectionsDistributedcaches {
	if o == nil || o.DistributedCaches == nil {
		var ret MsgVpnCollectionsDistributedcaches
		return ret
	}
	return *o.DistributedCaches
}

// GetDistributedCachesOk returns a tuple with the DistributedCaches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetDistributedCachesOk() (*MsgVpnCollectionsDistributedcaches, bool) {
	if o == nil || o.DistributedCaches == nil {
		return nil, false
	}
	return o.DistributedCaches, true
}

// HasDistributedCaches returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasDistributedCaches() bool {
	if o != nil && o.DistributedCaches != nil {
		return true
	}

	return false
}

// SetDistributedCaches gets a reference to the given MsgVpnCollectionsDistributedcaches and assigns it to the DistributedCaches field.
func (o *MsgVpnCollections) SetDistributedCaches(v MsgVpnCollectionsDistributedcaches) {
	o.DistributedCaches = &v
}

// GetDmrBridges returns the DmrBridges field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetDmrBridges() MsgVpnCollectionsDmrbridges {
	if o == nil || o.DmrBridges == nil {
		var ret MsgVpnCollectionsDmrbridges
		return ret
	}
	return *o.DmrBridges
}

// GetDmrBridgesOk returns a tuple with the DmrBridges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetDmrBridgesOk() (*MsgVpnCollectionsDmrbridges, bool) {
	if o == nil || o.DmrBridges == nil {
		return nil, false
	}
	return o.DmrBridges, true
}

// HasDmrBridges returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasDmrBridges() bool {
	if o != nil && o.DmrBridges != nil {
		return true
	}

	return false
}

// SetDmrBridges gets a reference to the given MsgVpnCollectionsDmrbridges and assigns it to the DmrBridges field.
func (o *MsgVpnCollections) SetDmrBridges(v MsgVpnCollectionsDmrbridges) {
	o.DmrBridges = &v
}

// GetJndiConnectionFactories returns the JndiConnectionFactories field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetJndiConnectionFactories() MsgVpnCollectionsJndiconnectionfactories {
	if o == nil || o.JndiConnectionFactories == nil {
		var ret MsgVpnCollectionsJndiconnectionfactories
		return ret
	}
	return *o.JndiConnectionFactories
}

// GetJndiConnectionFactoriesOk returns a tuple with the JndiConnectionFactories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetJndiConnectionFactoriesOk() (*MsgVpnCollectionsJndiconnectionfactories, bool) {
	if o == nil || o.JndiConnectionFactories == nil {
		return nil, false
	}
	return o.JndiConnectionFactories, true
}

// HasJndiConnectionFactories returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasJndiConnectionFactories() bool {
	if o != nil && o.JndiConnectionFactories != nil {
		return true
	}

	return false
}

// SetJndiConnectionFactories gets a reference to the given MsgVpnCollectionsJndiconnectionfactories and assigns it to the JndiConnectionFactories field.
func (o *MsgVpnCollections) SetJndiConnectionFactories(v MsgVpnCollectionsJndiconnectionfactories) {
	o.JndiConnectionFactories = &v
}

// GetJndiQueues returns the JndiQueues field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetJndiQueues() MsgVpnCollectionsJndiqueues {
	if o == nil || o.JndiQueues == nil {
		var ret MsgVpnCollectionsJndiqueues
		return ret
	}
	return *o.JndiQueues
}

// GetJndiQueuesOk returns a tuple with the JndiQueues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetJndiQueuesOk() (*MsgVpnCollectionsJndiqueues, bool) {
	if o == nil || o.JndiQueues == nil {
		return nil, false
	}
	return o.JndiQueues, true
}

// HasJndiQueues returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasJndiQueues() bool {
	if o != nil && o.JndiQueues != nil {
		return true
	}

	return false
}

// SetJndiQueues gets a reference to the given MsgVpnCollectionsJndiqueues and assigns it to the JndiQueues field.
func (o *MsgVpnCollections) SetJndiQueues(v MsgVpnCollectionsJndiqueues) {
	o.JndiQueues = &v
}

// GetJndiTopics returns the JndiTopics field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetJndiTopics() MsgVpnCollectionsJnditopics {
	if o == nil || o.JndiTopics == nil {
		var ret MsgVpnCollectionsJnditopics
		return ret
	}
	return *o.JndiTopics
}

// GetJndiTopicsOk returns a tuple with the JndiTopics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetJndiTopicsOk() (*MsgVpnCollectionsJnditopics, bool) {
	if o == nil || o.JndiTopics == nil {
		return nil, false
	}
	return o.JndiTopics, true
}

// HasJndiTopics returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasJndiTopics() bool {
	if o != nil && o.JndiTopics != nil {
		return true
	}

	return false
}

// SetJndiTopics gets a reference to the given MsgVpnCollectionsJnditopics and assigns it to the JndiTopics field.
func (o *MsgVpnCollections) SetJndiTopics(v MsgVpnCollectionsJnditopics) {
	o.JndiTopics = &v
}

// GetMqttRetainCaches returns the MqttRetainCaches field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetMqttRetainCaches() MsgVpnCollectionsMqttretaincaches {
	if o == nil || o.MqttRetainCaches == nil {
		var ret MsgVpnCollectionsMqttretaincaches
		return ret
	}
	return *o.MqttRetainCaches
}

// GetMqttRetainCachesOk returns a tuple with the MqttRetainCaches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetMqttRetainCachesOk() (*MsgVpnCollectionsMqttretaincaches, bool) {
	if o == nil || o.MqttRetainCaches == nil {
		return nil, false
	}
	return o.MqttRetainCaches, true
}

// HasMqttRetainCaches returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasMqttRetainCaches() bool {
	if o != nil && o.MqttRetainCaches != nil {
		return true
	}

	return false
}

// SetMqttRetainCaches gets a reference to the given MsgVpnCollectionsMqttretaincaches and assigns it to the MqttRetainCaches field.
func (o *MsgVpnCollections) SetMqttRetainCaches(v MsgVpnCollectionsMqttretaincaches) {
	o.MqttRetainCaches = &v
}

// GetMqttSessions returns the MqttSessions field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetMqttSessions() MsgVpnCollectionsMqttsessions {
	if o == nil || o.MqttSessions == nil {
		var ret MsgVpnCollectionsMqttsessions
		return ret
	}
	return *o.MqttSessions
}

// GetMqttSessionsOk returns a tuple with the MqttSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetMqttSessionsOk() (*MsgVpnCollectionsMqttsessions, bool) {
	if o == nil || o.MqttSessions == nil {
		return nil, false
	}
	return o.MqttSessions, true
}

// HasMqttSessions returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasMqttSessions() bool {
	if o != nil && o.MqttSessions != nil {
		return true
	}

	return false
}

// SetMqttSessions gets a reference to the given MsgVpnCollectionsMqttsessions and assigns it to the MqttSessions field.
func (o *MsgVpnCollections) SetMqttSessions(v MsgVpnCollectionsMqttsessions) {
	o.MqttSessions = &v
}

// GetQueueTemplates returns the QueueTemplates field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetQueueTemplates() MsgVpnCollectionsQueuetemplates {
	if o == nil || o.QueueTemplates == nil {
		var ret MsgVpnCollectionsQueuetemplates
		return ret
	}
	return *o.QueueTemplates
}

// GetQueueTemplatesOk returns a tuple with the QueueTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetQueueTemplatesOk() (*MsgVpnCollectionsQueuetemplates, bool) {
	if o == nil || o.QueueTemplates == nil {
		return nil, false
	}
	return o.QueueTemplates, true
}

// HasQueueTemplates returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasQueueTemplates() bool {
	if o != nil && o.QueueTemplates != nil {
		return true
	}

	return false
}

// SetQueueTemplates gets a reference to the given MsgVpnCollectionsQueuetemplates and assigns it to the QueueTemplates field.
func (o *MsgVpnCollections) SetQueueTemplates(v MsgVpnCollectionsQueuetemplates) {
	o.QueueTemplates = &v
}

// GetQueues returns the Queues field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetQueues() MsgVpnCollectionsQueues {
	if o == nil || o.Queues == nil {
		var ret MsgVpnCollectionsQueues
		return ret
	}
	return *o.Queues
}

// GetQueuesOk returns a tuple with the Queues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetQueuesOk() (*MsgVpnCollectionsQueues, bool) {
	if o == nil || o.Queues == nil {
		return nil, false
	}
	return o.Queues, true
}

// HasQueues returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasQueues() bool {
	if o != nil && o.Queues != nil {
		return true
	}

	return false
}

// SetQueues gets a reference to the given MsgVpnCollectionsQueues and assigns it to the Queues field.
func (o *MsgVpnCollections) SetQueues(v MsgVpnCollectionsQueues) {
	o.Queues = &v
}

// GetReplayLogs returns the ReplayLogs field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetReplayLogs() MsgVpnCollectionsReplaylogs {
	if o == nil || o.ReplayLogs == nil {
		var ret MsgVpnCollectionsReplaylogs
		return ret
	}
	return *o.ReplayLogs
}

// GetReplayLogsOk returns a tuple with the ReplayLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetReplayLogsOk() (*MsgVpnCollectionsReplaylogs, bool) {
	if o == nil || o.ReplayLogs == nil {
		return nil, false
	}
	return o.ReplayLogs, true
}

// HasReplayLogs returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasReplayLogs() bool {
	if o != nil && o.ReplayLogs != nil {
		return true
	}

	return false
}

// SetReplayLogs gets a reference to the given MsgVpnCollectionsReplaylogs and assigns it to the ReplayLogs field.
func (o *MsgVpnCollections) SetReplayLogs(v MsgVpnCollectionsReplaylogs) {
	o.ReplayLogs = &v
}

// GetReplicatedTopics returns the ReplicatedTopics field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetReplicatedTopics() MsgVpnCollectionsReplicatedtopics {
	if o == nil || o.ReplicatedTopics == nil {
		var ret MsgVpnCollectionsReplicatedtopics
		return ret
	}
	return *o.ReplicatedTopics
}

// GetReplicatedTopicsOk returns a tuple with the ReplicatedTopics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetReplicatedTopicsOk() (*MsgVpnCollectionsReplicatedtopics, bool) {
	if o == nil || o.ReplicatedTopics == nil {
		return nil, false
	}
	return o.ReplicatedTopics, true
}

// HasReplicatedTopics returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasReplicatedTopics() bool {
	if o != nil && o.ReplicatedTopics != nil {
		return true
	}

	return false
}

// SetReplicatedTopics gets a reference to the given MsgVpnCollectionsReplicatedtopics and assigns it to the ReplicatedTopics field.
func (o *MsgVpnCollections) SetReplicatedTopics(v MsgVpnCollectionsReplicatedtopics) {
	o.ReplicatedTopics = &v
}

// GetRestDeliveryPoints returns the RestDeliveryPoints field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetRestDeliveryPoints() MsgVpnCollectionsRestdeliverypoints {
	if o == nil || o.RestDeliveryPoints == nil {
		var ret MsgVpnCollectionsRestdeliverypoints
		return ret
	}
	return *o.RestDeliveryPoints
}

// GetRestDeliveryPointsOk returns a tuple with the RestDeliveryPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetRestDeliveryPointsOk() (*MsgVpnCollectionsRestdeliverypoints, bool) {
	if o == nil || o.RestDeliveryPoints == nil {
		return nil, false
	}
	return o.RestDeliveryPoints, true
}

// HasRestDeliveryPoints returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasRestDeliveryPoints() bool {
	if o != nil && o.RestDeliveryPoints != nil {
		return true
	}

	return false
}

// SetRestDeliveryPoints gets a reference to the given MsgVpnCollectionsRestdeliverypoints and assigns it to the RestDeliveryPoints field.
func (o *MsgVpnCollections) SetRestDeliveryPoints(v MsgVpnCollectionsRestdeliverypoints) {
	o.RestDeliveryPoints = &v
}

// GetTopicEndpointTemplates returns the TopicEndpointTemplates field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetTopicEndpointTemplates() MsgVpnCollectionsTopicendpointtemplates {
	if o == nil || o.TopicEndpointTemplates == nil {
		var ret MsgVpnCollectionsTopicendpointtemplates
		return ret
	}
	return *o.TopicEndpointTemplates
}

// GetTopicEndpointTemplatesOk returns a tuple with the TopicEndpointTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetTopicEndpointTemplatesOk() (*MsgVpnCollectionsTopicendpointtemplates, bool) {
	if o == nil || o.TopicEndpointTemplates == nil {
		return nil, false
	}
	return o.TopicEndpointTemplates, true
}

// HasTopicEndpointTemplates returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasTopicEndpointTemplates() bool {
	if o != nil && o.TopicEndpointTemplates != nil {
		return true
	}

	return false
}

// SetTopicEndpointTemplates gets a reference to the given MsgVpnCollectionsTopicendpointtemplates and assigns it to the TopicEndpointTemplates field.
func (o *MsgVpnCollections) SetTopicEndpointTemplates(v MsgVpnCollectionsTopicendpointtemplates) {
	o.TopicEndpointTemplates = &v
}

// GetTopicEndpoints returns the TopicEndpoints field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetTopicEndpoints() MsgVpnCollectionsTopicendpoints {
	if o == nil || o.TopicEndpoints == nil {
		var ret MsgVpnCollectionsTopicendpoints
		return ret
	}
	return *o.TopicEndpoints
}

// GetTopicEndpointsOk returns a tuple with the TopicEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetTopicEndpointsOk() (*MsgVpnCollectionsTopicendpoints, bool) {
	if o == nil || o.TopicEndpoints == nil {
		return nil, false
	}
	return o.TopicEndpoints, true
}

// HasTopicEndpoints returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasTopicEndpoints() bool {
	if o != nil && o.TopicEndpoints != nil {
		return true
	}

	return false
}

// SetTopicEndpoints gets a reference to the given MsgVpnCollectionsTopicendpoints and assigns it to the TopicEndpoints field.
func (o *MsgVpnCollections) SetTopicEndpoints(v MsgVpnCollectionsTopicendpoints) {
	o.TopicEndpoints = &v
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *MsgVpnCollections) GetTransactions() MsgVpnCollectionsTransactions {
	if o == nil || o.Transactions == nil {
		var ret MsgVpnCollectionsTransactions
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnCollections) GetTransactionsOk() (*MsgVpnCollectionsTransactions, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *MsgVpnCollections) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given MsgVpnCollectionsTransactions and assigns it to the Transactions field.
func (o *MsgVpnCollections) SetTransactions(v MsgVpnCollectionsTransactions) {
	o.Transactions = &v
}

func (o MsgVpnCollections) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AclProfiles != nil {
		toSerialize["aclProfiles"] = o.AclProfiles
	}
	if o.AuthenticationOauthProviders != nil {
		toSerialize["authenticationOauthProviders"] = o.AuthenticationOauthProviders
	}
	if o.AuthorizationGroups != nil {
		toSerialize["authorizationGroups"] = o.AuthorizationGroups
	}
	if o.Bridges != nil {
		toSerialize["bridges"] = o.Bridges
	}
	if o.ClientProfiles != nil {
		toSerialize["clientProfiles"] = o.ClientProfiles
	}
	if o.ClientUsernames != nil {
		toSerialize["clientUsernames"] = o.ClientUsernames
	}
	if o.Clients != nil {
		toSerialize["clients"] = o.Clients
	}
	if o.ConfigSyncRemoteNodes != nil {
		toSerialize["configSyncRemoteNodes"] = o.ConfigSyncRemoteNodes
	}
	if o.DistributedCaches != nil {
		toSerialize["distributedCaches"] = o.DistributedCaches
	}
	if o.DmrBridges != nil {
		toSerialize["dmrBridges"] = o.DmrBridges
	}
	if o.JndiConnectionFactories != nil {
		toSerialize["jndiConnectionFactories"] = o.JndiConnectionFactories
	}
	if o.JndiQueues != nil {
		toSerialize["jndiQueues"] = o.JndiQueues
	}
	if o.JndiTopics != nil {
		toSerialize["jndiTopics"] = o.JndiTopics
	}
	if o.MqttRetainCaches != nil {
		toSerialize["mqttRetainCaches"] = o.MqttRetainCaches
	}
	if o.MqttSessions != nil {
		toSerialize["mqttSessions"] = o.MqttSessions
	}
	if o.QueueTemplates != nil {
		toSerialize["queueTemplates"] = o.QueueTemplates
	}
	if o.Queues != nil {
		toSerialize["queues"] = o.Queues
	}
	if o.ReplayLogs != nil {
		toSerialize["replayLogs"] = o.ReplayLogs
	}
	if o.ReplicatedTopics != nil {
		toSerialize["replicatedTopics"] = o.ReplicatedTopics
	}
	if o.RestDeliveryPoints != nil {
		toSerialize["restDeliveryPoints"] = o.RestDeliveryPoints
	}
	if o.TopicEndpointTemplates != nil {
		toSerialize["topicEndpointTemplates"] = o.TopicEndpointTemplates
	}
	if o.TopicEndpoints != nil {
		toSerialize["topicEndpoints"] = o.TopicEndpoints
	}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnCollections struct {
	value *MsgVpnCollections
	isSet bool
}

func (v NullableMsgVpnCollections) Get() *MsgVpnCollections {
	return v.value
}

func (v *NullableMsgVpnCollections) Set(val *MsgVpnCollections) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnCollections) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnCollections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnCollections(val *MsgVpnCollections) *NullableMsgVpnCollections {
	return &NullableMsgVpnCollections{value: val, isSet: true}
}

func (v NullableMsgVpnCollections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnCollections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
