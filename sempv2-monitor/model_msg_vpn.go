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

// MsgVpn struct for MsgVpn
type MsgVpn struct {
	// The name of another Message VPN which this Message VPN is an alias for. Available since 2.14.
	Alias *string `json:"alias,omitempty"`
	// Indicates whether basic authentication is enabled for clients connecting to the Message VPN.
	AuthenticationBasicEnabled *bool `json:"authenticationBasicEnabled,omitempty"`
	// The name of the RADIUS or LDAP Profile to use for basic authentication.
	AuthenticationBasicProfileName *string `json:"authenticationBasicProfileName,omitempty"`
	// The RADIUS domain to use for basic authentication.
	AuthenticationBasicRadiusDomain *string `json:"authenticationBasicRadiusDomain,omitempty"`
	// The type of basic authentication to use for clients connecting to the Message VPN. The allowed values and their meaning are:  <pre> \"internal\" - Internal database. Authentication is against Client Usernames. \"ldap\" - LDAP authentication. An LDAP profile name must be provided. \"radius\" - RADIUS authentication. A RADIUS profile name must be provided. \"none\" - No authentication. Anonymous login allowed. </pre>
	AuthenticationBasicType *string `json:"authenticationBasicType,omitempty"`
	// Indicates whether a client is allowed to specify a Client Username via the API connect method. When disabled, the certificate CN (Common Name) is always used.
	AuthenticationClientCertAllowApiProvidedUsernameEnabled *bool `json:"authenticationClientCertAllowApiProvidedUsernameEnabled,omitempty"`
	// Indicates whether client certificate authentication is enabled in the Message VPN.
	AuthenticationClientCertEnabled *bool `json:"authenticationClientCertEnabled,omitempty"`
	// The maximum depth for a client certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate.
	AuthenticationClientCertMaxChainDepth *int64 `json:"authenticationClientCertMaxChainDepth,omitempty"`
	// The desired behavior for client certificate revocation checking. The allowed values and their meaning are:  <pre> \"allow-all\" - Allow the client to authenticate, the result of client certificate revocation check is ignored. \"allow-unknown\" - Allow the client to authenticate even if the revocation status of his certificate cannot be determined. \"allow-valid\" - Allow the client to authenticate only when the revocation check returned an explicit positive response. </pre>
	AuthenticationClientCertRevocationCheckMode *string `json:"authenticationClientCertRevocationCheckMode,omitempty"`
	// The field from the client certificate to use as the client username. The allowed values and their meaning are:  <pre> \"certificate-thumbprint\" - The username is computed as the SHA-1 hash over the entire DER-encoded contents of the client certificate. \"common-name\" - The username is extracted from the certificate's first instance of the Common Name attribute in the Subject DN. \"common-name-last\" - The username is extracted from the certificate's last instance of the Common Name attribute in the Subject DN. \"subject-alternate-name-msupn\" - The username is extracted from the certificate's Other Name type of the Subject Alternative Name and must have the msUPN signature. \"uid\" - The username is extracted from the certificate's first instance of the User Identifier attribute in the Subject DN. \"uid-last\" - The username is extracted from the certificate's last instance of the User Identifier attribute in the Subject DN. </pre>
	AuthenticationClientCertUsernameSource *string `json:"authenticationClientCertUsernameSource,omitempty"`
	// Indicates whether the \"Not Before\" and \"Not After\" validity dates in the client certificate are checked.
	AuthenticationClientCertValidateDateEnabled *bool `json:"authenticationClientCertValidateDateEnabled,omitempty"`
	// Indicates whether a client is allowed to specify a Client Username via the API connect method. When disabled, the Kerberos Principal name is always used.
	AuthenticationKerberosAllowApiProvidedUsernameEnabled *bool `json:"authenticationKerberosAllowApiProvidedUsernameEnabled,omitempty"`
	// Indicates whether Kerberos authentication is enabled in the Message VPN.
	AuthenticationKerberosEnabled *bool `json:"authenticationKerberosEnabled,omitempty"`
	// The name of the provider to use when the client does not supply a provider name. Available since 2.13.
	AuthenticationOauthDefaultProviderName *string `json:"authenticationOauthDefaultProviderName,omitempty"`
	// Indicates whether OAuth authentication is enabled. Available since 2.13.
	AuthenticationOauthEnabled *bool `json:"authenticationOauthEnabled,omitempty"`
	// The name of the attribute that is retrieved from the LDAP server as part of the LDAP search when authorizing a client connecting to the Message VPN.
	AuthorizationLdapGroupMembershipAttributeName *string `json:"authorizationLdapGroupMembershipAttributeName,omitempty"`
	// Indicates whether client-username domain trimming for LDAP lookups of client connections is enabled. Available since 2.13.
	AuthorizationLdapTrimClientUsernameDomainEnabled *bool `json:"authorizationLdapTrimClientUsernameDomainEnabled,omitempty"`
	// The name of the LDAP Profile to use for client authorization.
	AuthorizationProfileName *string `json:"authorizationProfileName,omitempty"`
	// The type of authorization to use for clients connecting to the Message VPN. The allowed values and their meaning are:  <pre> \"ldap\" - LDAP authorization. \"internal\" - Internal authorization. </pre>
	AuthorizationType *string `json:"authorizationType,omitempty"`
	// The one minute average of the message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	AverageRxByteRate *int64 `json:"averageRxByteRate,omitempty"`
	// The one minute average of the compressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	AverageRxCompressedByteRate *int64 `json:"averageRxCompressedByteRate,omitempty"`
	// The one minute average of the message rate received by the Message VPN, in messages per second (msg/sec). Available since 2.13.
	AverageRxMsgRate *int64 `json:"averageRxMsgRate,omitempty"`
	// The one minute average of the uncompressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	AverageRxUncompressedByteRate *int64 `json:"averageRxUncompressedByteRate,omitempty"`
	// The one minute average of the message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	AverageTxByteRate *int64 `json:"averageTxByteRate,omitempty"`
	// The one minute average of the compressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	AverageTxCompressedByteRate *int64 `json:"averageTxCompressedByteRate,omitempty"`
	// The one minute average of the message rate transmitted by the Message VPN, in messages per second (msg/sec). Available since 2.13.
	AverageTxMsgRate *int64 `json:"averageTxMsgRate,omitempty"`
	// The one minute average of the uncompressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	AverageTxUncompressedByteRate *int64 `json:"averageTxUncompressedByteRate,omitempty"`
	// Indicates whether the Common Name (CN) in the server certificate from the remote broker is validated for the Bridge. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
	BridgingTlsServerCertEnforceTrustedCommonNameEnabled *bool `json:"bridgingTlsServerCertEnforceTrustedCommonNameEnabled,omitempty"`
	// The maximum depth for a server certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate.
	BridgingTlsServerCertMaxChainDepth *int64 `json:"bridgingTlsServerCertMaxChainDepth,omitempty"`
	// Indicates whether the \"Not Before\" and \"Not After\" validity dates in the server certificate are checked.
	BridgingTlsServerCertValidateDateEnabled *bool `json:"bridgingTlsServerCertValidateDateEnabled,omitempty"`
	// Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the bridge. If enabled, the name used to connect to the bridge is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. Available since 2.18.
	BridgingTlsServerCertValidateNameEnabled *bool `json:"bridgingTlsServerCertValidateNameEnabled,omitempty"`
	// The key for the config sync table of the local Message VPN. Available since 2.12.
	ConfigSyncLocalKey *string `json:"configSyncLocalKey,omitempty"`
	// The result of the last operation on the config sync table of the local Message VPN. Available since 2.12.
	ConfigSyncLocalLastResult *string `json:"configSyncLocalLastResult,omitempty"`
	// The role of the config sync table of the local Message VPN. The allowed values and their meaning are:  <pre> \"unknown\" - The role is unknown. \"primary\" - Acts as the primary source of config data. \"replica\" - Acts as a replica of the primary config data. </pre>  Available since 2.12.
	ConfigSyncLocalRole *string `json:"configSyncLocalRole,omitempty"`
	// The state of the config sync table of the local Message VPN. The allowed values and their meaning are:  <pre> \"unknown\" - The state is unknown. \"in-sync\" - The config data is synchronized between Message VPNs. \"reconciling\" - The config data is reconciling between Message VPNs. \"blocked\" - The config data is blocked from reconciling due to an error. \"out-of-sync\" - The config data is out of sync between Message VPNs. \"down\" - The state is down due to configuration. </pre>  Available since 2.12.
	ConfigSyncLocalState *string `json:"configSyncLocalState,omitempty"`
	// The amount of time in seconds the config sync table of the local Message VPN has been in the current state. Available since 2.12.
	ConfigSyncLocalTimeInState *int32 `json:"configSyncLocalTimeInState,omitempty"`
	// The amount of client control messages received from clients by the Message VPN, in bytes (B). Available since 2.13.
	ControlRxByteCount *int64 `json:"controlRxByteCount,omitempty"`
	// The number of client control messages received from clients by the Message VPN. Available since 2.13.
	ControlRxMsgCount *int64 `json:"controlRxMsgCount,omitempty"`
	// The amount of client control messages transmitted to clients by the Message VPN, in bytes (B). Available since 2.13.
	ControlTxByteCount *int64 `json:"controlTxByteCount,omitempty"`
	// The number of client control messages transmitted to clients by the Message VPN. Available since 2.13.
	ControlTxMsgCount *int64         `json:"controlTxMsgCount,omitempty"`
	Counter           *MsgVpnCounter `json:"counter,omitempty"`
	// The amount of client data messages received from clients by the Message VPN, in bytes (B). Available since 2.13.
	DataRxByteCount *int64 `json:"dataRxByteCount,omitempty"`
	// The number of client data messages received from clients by the Message VPN. Available since 2.13.
	DataRxMsgCount *int64 `json:"dataRxMsgCount,omitempty"`
	// The amount of client data messages transmitted to clients by the Message VPN, in bytes (B). Available since 2.13.
	DataTxByteCount *int64 `json:"dataTxByteCount,omitempty"`
	// The number of client data messages transmitted to clients by the Message VPN. Available since 2.13.
	DataTxMsgCount *int64 `json:"dataTxMsgCount,omitempty"`
	// The number of messages discarded during reception by the Message VPN. Available since 2.13.
	DiscardedRxMsgCount *int64 `json:"discardedRxMsgCount,omitempty"`
	// The number of messages discarded during transmission by the Message VPN. Available since 2.13.
	DiscardedTxMsgCount *int64 `json:"discardedTxMsgCount,omitempty"`
	// Indicates whether managing of cache instances over the message bus is enabled in the Message VPN.
	DistributedCacheManagementEnabled *bool `json:"distributedCacheManagementEnabled,omitempty"`
	// Indicates whether Dynamic Message Routing (DMR) is enabled for the Message VPN.
	DmrEnabled *bool `json:"dmrEnabled,omitempty"`
	// Indicates whether the Message VPN is enabled.
	Enabled                        *bool                  `json:"enabled,omitempty"`
	EventConnectionCountThreshold  *EventThreshold        `json:"eventConnectionCountThreshold,omitempty"`
	EventEgressFlowCountThreshold  *EventThreshold        `json:"eventEgressFlowCountThreshold,omitempty"`
	EventEgressMsgRateThreshold    *EventThresholdByValue `json:"eventEgressMsgRateThreshold,omitempty"`
	EventEndpointCountThreshold    *EventThreshold        `json:"eventEndpointCountThreshold,omitempty"`
	EventIngressFlowCountThreshold *EventThreshold        `json:"eventIngressFlowCountThreshold,omitempty"`
	EventIngressMsgRateThreshold   *EventThresholdByValue `json:"eventIngressMsgRateThreshold,omitempty"`
	// Exceeding this message size in kilobytes (KB) triggers a corresponding Event in the Message VPN.
	EventLargeMsgThreshold *int64 `json:"eventLargeMsgThreshold,omitempty"`
	// The value of the prefix applied to all published Events in the Message VPN.
	EventLogTag                 *string         `json:"eventLogTag,omitempty"`
	EventMsgSpoolUsageThreshold *EventThreshold `json:"eventMsgSpoolUsageThreshold,omitempty"`
	// Indicates whether client Events are published in the Message VPN.
	EventPublishClientEnabled *bool `json:"eventPublishClientEnabled,omitempty"`
	// Indicates whether Message VPN Events are published in the Message VPN.
	EventPublishMsgVpnEnabled *bool `json:"eventPublishMsgVpnEnabled,omitempty"`
	// The mode of subscription Events published in the Message VPN. The allowed values and their meaning are:  <pre> \"off\" - Disable client level event message publishing. \"on-with-format-v1\" - Enable client level event message publishing with format v1. \"on-with-no-unsubscribe-events-on-disconnect-format-v1\" - As \"on-with-format-v1\", but unsubscribe events are not generated when a client disconnects. Unsubscribe events are still raised when a client explicitly unsubscribes from its subscriptions. \"on-with-format-v2\" - Enable client level event message publishing with format v2. \"on-with-no-unsubscribe-events-on-disconnect-format-v2\" - As \"on-with-format-v2\", but unsubscribe events are not generated when a client disconnects. Unsubscribe events are still raised when a client explicitly unsubscribes from its subscriptions. </pre>
	EventPublishSubscriptionMode *string `json:"eventPublishSubscriptionMode,omitempty"`
	// Indicates whether Message VPN Events are published in the MQTT format.
	EventPublishTopicFormatMqttEnabled *bool `json:"eventPublishTopicFormatMqttEnabled,omitempty"`
	// Indicates whether Message VPN Events are published in the SMF format.
	EventPublishTopicFormatSmfEnabled                *bool           `json:"eventPublishTopicFormatSmfEnabled,omitempty"`
	EventServiceAmqpConnectionCountThreshold         *EventThreshold `json:"eventServiceAmqpConnectionCountThreshold,omitempty"`
	EventServiceMqttConnectionCountThreshold         *EventThreshold `json:"eventServiceMqttConnectionCountThreshold,omitempty"`
	EventServiceRestIncomingConnectionCountThreshold *EventThreshold `json:"eventServiceRestIncomingConnectionCountThreshold,omitempty"`
	EventServiceSmfConnectionCountThreshold          *EventThreshold `json:"eventServiceSmfConnectionCountThreshold,omitempty"`
	EventServiceWebConnectionCountThreshold          *EventThreshold `json:"eventServiceWebConnectionCountThreshold,omitempty"`
	EventSubscriptionCountThreshold                  *EventThreshold `json:"eventSubscriptionCountThreshold,omitempty"`
	EventTransactedSessionCountThreshold             *EventThreshold `json:"eventTransactedSessionCountThreshold,omitempty"`
	EventTransactionCountThreshold                   *EventThreshold `json:"eventTransactionCountThreshold,omitempty"`
	// Indicates whether exports of subscriptions to other routers in the network over neighbour links is enabled in the Message VPN.
	ExportSubscriptionsEnabled *bool `json:"exportSubscriptionsEnabled,omitempty"`
	// The reason for the Message VPN failure.
	FailureReason *string `json:"failureReason,omitempty"`
	// Indicates whether the JNDI access for clients is enabled in the Message VPN.
	JndiEnabled *bool `json:"jndiEnabled,omitempty"`
	// The number of login request messages received by the Message VPN. Available since 2.13.
	LoginRxMsgCount *int64 `json:"loginRxMsgCount,omitempty"`
	// The number of login response messages transmitted by the Message VPN. Available since 2.13.
	LoginTxMsgCount *int64 `json:"loginTxMsgCount,omitempty"`
	// The maximum number of client connections to the Message VPN.
	MaxConnectionCount *int64 `json:"maxConnectionCount,omitempty"`
	// The effective maximum number of Queues and Topic Endpoints allowed in the Message VPN.
	MaxEffectiveEndpointCount *int32 `json:"maxEffectiveEndpointCount,omitempty"`
	// The effective maximum number of receive flows allowed in the Message VPN.
	MaxEffectiveRxFlowCount *int32 `json:"maxEffectiveRxFlowCount,omitempty"`
	// The effective maximum number of subscriptions allowed in the Message VPN.
	MaxEffectiveSubscriptionCount *int64 `json:"maxEffectiveSubscriptionCount,omitempty"`
	// The effective maximum number of transacted sessions allowed in the Message VPN.
	MaxEffectiveTransactedSessionCount *int32 `json:"maxEffectiveTransactedSessionCount,omitempty"`
	// The effective maximum number of transactions allowed in the Message VPN.
	MaxEffectiveTransactionCount *int32 `json:"maxEffectiveTransactionCount,omitempty"`
	// The effective maximum number of transmit flows allowed in the Message VPN.
	MaxEffectiveTxFlowCount *int32 `json:"maxEffectiveTxFlowCount,omitempty"`
	// The maximum number of transmit flows that can be created in the Message VPN.
	MaxEgressFlowCount *int64 `json:"maxEgressFlowCount,omitempty"`
	// The maximum number of Queues and Topic Endpoints that can be created in the Message VPN.
	MaxEndpointCount *int64 `json:"maxEndpointCount,omitempty"`
	// The maximum number of receive flows that can be created in the Message VPN.
	MaxIngressFlowCount *int64 `json:"maxIngressFlowCount,omitempty"`
	// The maximum message spool usage by the Message VPN, in megabytes.
	MaxMsgSpoolUsage *int64 `json:"maxMsgSpoolUsage,omitempty"`
	// The maximum number of local client subscriptions that can be added to the Message VPN. This limit is not enforced when a subscription is added using a management interface, such as CLI or SEMP.
	MaxSubscriptionCount *int64 `json:"maxSubscriptionCount,omitempty"`
	// The maximum number of transacted sessions that can be created in the Message VPN.
	MaxTransactedSessionCount *int64 `json:"maxTransactedSessionCount,omitempty"`
	// The maximum number of transactions that can be created in the Message VPN.
	MaxTransactionCount *int64 `json:"maxTransactionCount,omitempty"`
	// The maximum total memory usage of the MQTT Retain feature for this Message VPN, in MB. If the maximum memory is reached, any arriving retain messages that require more memory are discarded. A value of -1 indicates that the memory is bounded only by the global max memory limit. A value of 0 prevents MQTT Retain from becoming operational.
	MqttRetainMaxMemory *int32 `json:"mqttRetainMaxMemory,omitempty"`
	// The number of message replays that are currently active in the Message VPN.
	MsgReplayActiveCount *int32 `json:"msgReplayActiveCount,omitempty"`
	// The number of message replays that are currently failed in the Message VPN.
	MsgReplayFailedCount *int32 `json:"msgReplayFailedCount,omitempty"`
	// The number of message replays that are currently initializing in the Message VPN.
	MsgReplayInitializingCount *int32 `json:"msgReplayInitializingCount,omitempty"`
	// The number of message replays that are pending complete in the Message VPN.
	MsgReplayPendingCompleteCount *int32 `json:"msgReplayPendingCompleteCount,omitempty"`
	// The current number of messages spooled (persisted in the Message Spool) in the Message VPN. Available since 2.14.
	MsgSpoolMsgCount *int64 `json:"msgSpoolMsgCount,omitempty"`
	// The number of guaranteed messages received by the Message VPN. Available since 2.13.
	MsgSpoolRxMsgCount *int64 `json:"msgSpoolRxMsgCount,omitempty"`
	// The number of guaranteed messages transmitted by the Message VPN. One message to multiple clients is counted as one message. Available since 2.13.
	MsgSpoolTxMsgCount *int64 `json:"msgSpoolTxMsgCount,omitempty"`
	// The current message spool usage by the Message VPN, in bytes (B).
	MsgSpoolUsage *int64 `json:"msgSpoolUsage,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string     `json:"msgVpnName,omitempty"`
	Rate       *MsgVpnRate `json:"rate,omitempty"`
	// The acknowledgement (ACK) propagation interval for the replication Bridge, in number of replicated messages. Available since 2.12.
	ReplicationAckPropagationIntervalMsgCount *int64 `json:"replicationAckPropagationIntervalMsgCount,omitempty"`
	// The number of acknowledgement messages propagated to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveAckPropTxMsgCount *int64 `json:"replicationActiveAckPropTxMsgCount,omitempty"`
	// The number of async messages queued to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveAsyncQueuedMsgCount *int64 `json:"replicationActiveAsyncQueuedMsgCount,omitempty"`
	// The number of messages consumed in the replication active local Message VPN. Available since 2.12.
	ReplicationActiveLocallyConsumedMsgCount *int64 `json:"replicationActiveLocallyConsumedMsgCount,omitempty"`
	// The peak amount of time in seconds the message flow has been congested to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveMateFlowCongestedPeakTime *int32 `json:"replicationActiveMateFlowCongestedPeakTime,omitempty"`
	// The peak amount of time in seconds the message flow has not been congested to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveMateFlowNotCongestedPeakTime *int32 `json:"replicationActiveMateFlowNotCongestedPeakTime,omitempty"`
	// The number of promoted messages queued to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActivePromotedQueuedMsgCount *int64 `json:"replicationActivePromotedQueuedMsgCount,omitempty"`
	// The number of reconcile request messages received from the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveReconcileRequestRxMsgCount *int64 `json:"replicationActiveReconcileRequestRxMsgCount,omitempty"`
	// The peak amount of time in seconds sync replication has been eligible to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveSyncEligiblePeakTime *int32 `json:"replicationActiveSyncEligiblePeakTime,omitempty"`
	// The peak amount of time in seconds sync replication has been ineligible to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveSyncIneligiblePeakTime *int32 `json:"replicationActiveSyncIneligiblePeakTime,omitempty"`
	// The number of sync messages queued as async to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveSyncQueuedAsAsyncMsgCount *int64 `json:"replicationActiveSyncQueuedAsAsyncMsgCount,omitempty"`
	// The number of sync messages queued to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveSyncQueuedMsgCount *int64 `json:"replicationActiveSyncQueuedMsgCount,omitempty"`
	// The number of sync replication ineligible transitions to the replication standby remote Message VPN. Available since 2.12.
	ReplicationActiveTransitionToSyncIneligibleCount *int64 `json:"replicationActiveTransitionToSyncIneligibleCount,omitempty"`
	// The Client Username the replication Bridge uses to login to the remote Message VPN. Available since 2.12.
	ReplicationBridgeAuthenticationBasicClientUsername *string `json:"replicationBridgeAuthenticationBasicClientUsername,omitempty"`
	// The authentication scheme for the replication Bridge in the Message VPN. The allowed values and their meaning are:  <pre> \"basic\" - Basic Authentication Scheme (via username and password). \"client-certificate\" - Client Certificate Authentication Scheme (via certificate file or content). </pre>  Available since 2.12.
	ReplicationBridgeAuthenticationScheme *string `json:"replicationBridgeAuthenticationScheme,omitempty"`
	// Indicates whether the local replication Bridge is bound to the Queue in the remote Message VPN. Available since 2.12.
	ReplicationBridgeBoundToQueue *bool `json:"replicationBridgeBoundToQueue,omitempty"`
	// Indicates whether compression is used for the replication Bridge. Available since 2.12.
	ReplicationBridgeCompressedDataEnabled *bool `json:"replicationBridgeCompressedDataEnabled,omitempty"`
	// The size of the window used for guaranteed messages published to the replication Bridge, in messages. Available since 2.12.
	ReplicationBridgeEgressFlowWindowSize *int64 `json:"replicationBridgeEgressFlowWindowSize,omitempty"`
	// The name of the local replication Bridge in the Message VPN. Available since 2.12.
	ReplicationBridgeName *string `json:"replicationBridgeName,omitempty"`
	// The number of seconds that must pass before retrying the replication Bridge connection. Available since 2.12.
	ReplicationBridgeRetryDelay *int64 `json:"replicationBridgeRetryDelay,omitempty"`
	// Indicates whether encryption (TLS) is enabled for the replication Bridge connection. Available since 2.12.
	ReplicationBridgeTlsEnabled *bool `json:"replicationBridgeTlsEnabled,omitempty"`
	// The Client Profile for the unidirectional replication Bridge in the Message VPN. It is used only for the TCP parameters. Available since 2.12.
	ReplicationBridgeUnidirectionalClientProfileName *string `json:"replicationBridgeUnidirectionalClientProfileName,omitempty"`
	// Indicates whether the local replication Bridge is operationally up in the Message VPN. Available since 2.12.
	ReplicationBridgeUp *bool `json:"replicationBridgeUp,omitempty"`
	// Indicates whether replication is enabled for the Message VPN. Available since 2.12.
	ReplicationEnabled *bool `json:"replicationEnabled,omitempty"`
	// Indicates whether the remote replication Bridge is bound to the Queue in the Message VPN. Available since 2.12.
	ReplicationQueueBound *bool `json:"replicationQueueBound,omitempty"`
	// The maximum message spool usage by the replication Bridge local Queue (quota), in megabytes. Available since 2.12.
	ReplicationQueueMaxMsgSpoolUsage *int64 `json:"replicationQueueMaxMsgSpoolUsage,omitempty"`
	// Indicates whether messages discarded on this replication Bridge Queue are rejected back to the sender. Available since 2.12.
	ReplicationQueueRejectMsgToSenderOnDiscardEnabled *bool `json:"replicationQueueRejectMsgToSenderOnDiscardEnabled,omitempty"`
	// Indicates whether guaranteed messages published to synchronously replicated Topics are rejected back to the sender when synchronous replication becomes ineligible. Available since 2.12.
	ReplicationRejectMsgWhenSyncIneligibleEnabled *bool `json:"replicationRejectMsgWhenSyncIneligibleEnabled,omitempty"`
	// The name of the remote replication Bridge in the Message VPN. Available since 2.12.
	ReplicationRemoteBridgeName *string `json:"replicationRemoteBridgeName,omitempty"`
	// Indicates whether the remote replication Bridge is operationally up in the Message VPN. Available since 2.12.
	ReplicationRemoteBridgeUp *bool `json:"replicationRemoteBridgeUp,omitempty"`
	// The replication role for the Message VPN. The allowed values and their meaning are:  <pre> \"active\" - Assume the Active role in replication for the Message VPN. \"standby\" - Assume the Standby role in replication for the Message VPN. </pre>  Available since 2.12.
	ReplicationRole *string `json:"replicationRole,omitempty"`
	// The number of acknowledgement messages received out of sequence from the replication active remote Message VPN. Available since 2.12.
	ReplicationStandbyAckPropOutOfSeqRxMsgCount *int64 `json:"replicationStandbyAckPropOutOfSeqRxMsgCount,omitempty"`
	// The number of acknowledgement messages received from the replication active remote Message VPN. Available since 2.12.
	ReplicationStandbyAckPropRxMsgCount *int64 `json:"replicationStandbyAckPropRxMsgCount,omitempty"`
	// The number of reconcile request messages transmitted to the replication active remote Message VPN. Available since 2.12.
	ReplicationStandbyReconcileRequestTxMsgCount *int64 `json:"replicationStandbyReconcileRequestTxMsgCount,omitempty"`
	// The number of messages received from the replication active remote Message VPN. Available since 2.12.
	ReplicationStandbyRxMsgCount *int64 `json:"replicationStandbyRxMsgCount,omitempty"`
	// The number of transaction requests received from the replication active remote Message VPN. Available since 2.12.
	ReplicationStandbyTransactionRequestCount *int64 `json:"replicationStandbyTransactionRequestCount,omitempty"`
	// The number of transaction requests received from the replication active remote Message VPN that failed. Available since 2.12.
	ReplicationStandbyTransactionRequestFailureCount *int64 `json:"replicationStandbyTransactionRequestFailureCount,omitempty"`
	// The number of transaction requests received from the replication active remote Message VPN that succeeded. Available since 2.12.
	ReplicationStandbyTransactionRequestSuccessCount *int64 `json:"replicationStandbyTransactionRequestSuccessCount,omitempty"`
	// Indicates whether sync replication is eligible in the Message VPN. Available since 2.12.
	ReplicationSyncEligible *bool `json:"replicationSyncEligible,omitempty"`
	// Indicates whether synchronous or asynchronous replication mode is used for all transactions within the Message VPN. The allowed values and their meaning are:  <pre> \"sync\" - Messages are acknowledged when replicated (spooled remotely). \"async\" - Messages are acknowledged when pending replication (spooled locally). </pre>  Available since 2.12.
	ReplicationTransactionMode *string `json:"replicationTransactionMode,omitempty"`
	// Indicates whether the Common Name (CN) in the server certificate from the remote REST Consumer is validated. Deprecated since 2.17. Common Name validation has been replaced by Server Certificate Name validation.
	RestTlsServerCertEnforceTrustedCommonNameEnabled *bool `json:"restTlsServerCertEnforceTrustedCommonNameEnabled,omitempty"`
	// The maximum depth for a REST Consumer server certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate.
	RestTlsServerCertMaxChainDepth *int64 `json:"restTlsServerCertMaxChainDepth,omitempty"`
	// Indicates whether the \"Not Before\" and \"Not After\" validity dates in the REST Consumer server certificate are checked.
	RestTlsServerCertValidateDateEnabled *bool `json:"restTlsServerCertValidateDateEnabled,omitempty"`
	// Enable or disable the standard TLS authentication mechanism of verifying the name used to connect to the remote REST Consumer. If enabled, the name used to connect to the remote REST Consumer is checked against the names specified in the certificate returned by the remote router. Legacy Common Name validation is not performed if Server Certificate Name Validation is enabled, even if Common Name validation is also enabled. Available since 2.17.
	RestTlsServerCertValidateNameEnabled *bool `json:"restTlsServerCertValidateNameEnabled,omitempty"`
	// The amount of messages received from clients by the Message VPN, in bytes (B). Available since 2.12.
	RxByteCount *int64 `json:"rxByteCount,omitempty"`
	// The current message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	RxByteRate *int64 `json:"rxByteRate,omitempty"`
	// The amount of compressed messages received by the Message VPN, in bytes (B). Available since 2.12.
	RxCompressedByteCount *int64 `json:"rxCompressedByteCount,omitempty"`
	// The current compressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	RxCompressedByteRate *int64 `json:"rxCompressedByteRate,omitempty"`
	// The compression ratio for messages received by the message VPN. Available since 2.12.
	RxCompressionRatio *string `json:"rxCompressionRatio,omitempty"`
	// The number of messages received from clients by the Message VPN. Available since 2.12.
	RxMsgCount *int64 `json:"rxMsgCount,omitempty"`
	// The current message rate received by the Message VPN, in messages per second (msg/sec). Available since 2.13.
	RxMsgRate *int64 `json:"rxMsgRate,omitempty"`
	// The amount of uncompressed messages received by the Message VPN, in bytes (B). Available since 2.12.
	RxUncompressedByteCount *int64 `json:"rxUncompressedByteCount,omitempty"`
	// The current uncompressed message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	RxUncompressedByteRate *int64 `json:"rxUncompressedByteRate,omitempty"`
	// Indicates whether the \"admin\" level \"client\" commands are enabled for SEMP over the message bus in the Message VPN.
	SempOverMsgBusAdminClientEnabled *bool `json:"sempOverMsgBusAdminClientEnabled,omitempty"`
	// Indicates whether the \"admin\" level \"Distributed Cache\" commands are enabled for SEMP over the message bus in the Message VPN.
	SempOverMsgBusAdminDistributedCacheEnabled *bool `json:"sempOverMsgBusAdminDistributedCacheEnabled,omitempty"`
	// Indicates whether the \"admin\" level commands are enabled for SEMP over the message bus in the Message VPN.
	SempOverMsgBusAdminEnabled *bool `json:"sempOverMsgBusAdminEnabled,omitempty"`
	// Indicates whether SEMP over the message bus is enabled in the Message VPN.
	SempOverMsgBusEnabled *bool `json:"sempOverMsgBusEnabled,omitempty"`
	// Indicates whether the \"show\" level commands are enabled for SEMP over the message bus in the Message VPN.
	SempOverMsgBusShowEnabled *bool `json:"sempOverMsgBusShowEnabled,omitempty"`
	// The maximum number of AMQP client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform.
	ServiceAmqpMaxConnectionCount *int64 `json:"serviceAmqpMaxConnectionCount,omitempty"`
	// Indicates whether the AMQP Service is compressed in the Message VPN.
	ServiceAmqpPlainTextCompressed *bool `json:"serviceAmqpPlainTextCompressed,omitempty"`
	// Indicates whether the AMQP Service is enabled in the Message VPN.
	ServiceAmqpPlainTextEnabled *bool `json:"serviceAmqpPlainTextEnabled,omitempty"`
	// The reason for the AMQP Service failure in the Message VPN.
	ServiceAmqpPlainTextFailureReason *string `json:"serviceAmqpPlainTextFailureReason,omitempty"`
	// The port number for plain-text AMQP clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceAmqpPlainTextListenPort *int64 `json:"serviceAmqpPlainTextListenPort,omitempty"`
	// Indicates whether the AMQP Service is operationally up in the Message VPN.
	ServiceAmqpPlainTextUp *bool `json:"serviceAmqpPlainTextUp,omitempty"`
	// Indicates whether the TLS related AMQP Service is compressed in the Message VPN.
	ServiceAmqpTlsCompressed *bool `json:"serviceAmqpTlsCompressed,omitempty"`
	// Indicates whether encryption (TLS) is enabled for AMQP clients in the Message VPN.
	ServiceAmqpTlsEnabled *bool `json:"serviceAmqpTlsEnabled,omitempty"`
	// The reason for the TLS related AMQP Service failure in the Message VPN.
	ServiceAmqpTlsFailureReason *string `json:"serviceAmqpTlsFailureReason,omitempty"`
	// The port number for AMQP clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceAmqpTlsListenPort *int64 `json:"serviceAmqpTlsListenPort,omitempty"`
	// Indicates whether the TLS related AMQP Service is operationally up in the Message VPN.
	ServiceAmqpTlsUp *bool `json:"serviceAmqpTlsUp,omitempty"`
	// Determines when to request a client certificate from an incoming MQTT client connecting via a TLS port. The allowed values and their meaning are:  <pre> \"always\" - Always ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"never\" - Never ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"when-enabled-in-message-vpn\" - Only ask for a client-certificate if client certificate authentication is enabled under \"message-vpn >  authentication > client-certificate > shutdown\". </pre>  Available since 2.21.
	ServiceMqttAuthenticationClientCertRequest *string `json:"serviceMqttAuthenticationClientCertRequest,omitempty"`
	// The maximum number of MQTT client connections that can be simultaneously connected to the Message VPN.
	ServiceMqttMaxConnectionCount *int64 `json:"serviceMqttMaxConnectionCount,omitempty"`
	// Indicates whether the MQTT Service is compressed in the Message VPN.
	ServiceMqttPlainTextCompressed *bool `json:"serviceMqttPlainTextCompressed,omitempty"`
	// Indicates whether the MQTT Service is enabled in the Message VPN.
	ServiceMqttPlainTextEnabled *bool `json:"serviceMqttPlainTextEnabled,omitempty"`
	// The reason for the MQTT Service failure in the Message VPN.
	ServiceMqttPlainTextFailureReason *string `json:"serviceMqttPlainTextFailureReason,omitempty"`
	// The port number for plain-text MQTT clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceMqttPlainTextListenPort *int64 `json:"serviceMqttPlainTextListenPort,omitempty"`
	// Indicates whether the MQTT Service is operationally up in the Message VPN.
	ServiceMqttPlainTextUp *bool `json:"serviceMqttPlainTextUp,omitempty"`
	// Indicates whether the TLS related MQTT Service is compressed in the Message VPN.
	ServiceMqttTlsCompressed *bool `json:"serviceMqttTlsCompressed,omitempty"`
	// Indicates whether encryption (TLS) is enabled for MQTT clients in the Message VPN.
	ServiceMqttTlsEnabled *bool `json:"serviceMqttTlsEnabled,omitempty"`
	// The reason for the TLS related MQTT Service failure in the Message VPN.
	ServiceMqttTlsFailureReason *string `json:"serviceMqttTlsFailureReason,omitempty"`
	// The port number for MQTT clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceMqttTlsListenPort *int64 `json:"serviceMqttTlsListenPort,omitempty"`
	// Indicates whether the TLS related MQTT Service is operationally up in the Message VPN.
	ServiceMqttTlsUp *bool `json:"serviceMqttTlsUp,omitempty"`
	// Indicates whether the TLS related Web transport MQTT Service is compressed in the Message VPN.
	ServiceMqttTlsWebSocketCompressed *bool `json:"serviceMqttTlsWebSocketCompressed,omitempty"`
	// Indicates whether encryption (TLS) is enabled for MQTT Web clients in the Message VPN.
	ServiceMqttTlsWebSocketEnabled *bool `json:"serviceMqttTlsWebSocketEnabled,omitempty"`
	// The reason for the TLS related Web transport MQTT Service failure in the Message VPN.
	ServiceMqttTlsWebSocketFailureReason *string `json:"serviceMqttTlsWebSocketFailureReason,omitempty"`
	// The port number for MQTT clients that connect to the Message VPN using WebSocket over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceMqttTlsWebSocketListenPort *int64 `json:"serviceMqttTlsWebSocketListenPort,omitempty"`
	// Indicates whether the TLS related Web transport MQTT Service is operationally up in the Message VPN.
	ServiceMqttTlsWebSocketUp *bool `json:"serviceMqttTlsWebSocketUp,omitempty"`
	// Indicates whether the Web transport related MQTT Service is compressed in the Message VPN.
	ServiceMqttWebSocketCompressed *bool `json:"serviceMqttWebSocketCompressed,omitempty"`
	// Indicates whether the Web transport for the SMF Service is enabled in the Message VPN.
	ServiceMqttWebSocketEnabled *bool `json:"serviceMqttWebSocketEnabled,omitempty"`
	// The reason for the Web transport related MQTT Service failure in the Message VPN.
	ServiceMqttWebSocketFailureReason *string `json:"serviceMqttWebSocketFailureReason,omitempty"`
	// The port number for plain-text MQTT clients that connect to the Message VPN using WebSocket. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceMqttWebSocketListenPort *int64 `json:"serviceMqttWebSocketListenPort,omitempty"`
	// Indicates whether the Web transport related MQTT Service is operationally up in the Message VPN.
	ServiceMqttWebSocketUp *bool `json:"serviceMqttWebSocketUp,omitempty"`
	// Determines when to request a client certificate from an incoming REST Producer connecting via a TLS port. The allowed values and their meaning are:  <pre> \"always\" - Always ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"never\" - Never ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"when-enabled-in-message-vpn\" - Only ask for a client-certificate if client certificate authentication is enabled under \"message-vpn >  authentication > client-certificate > shutdown\". </pre>  Available since 2.21.
	ServiceRestIncomingAuthenticationClientCertRequest *string `json:"serviceRestIncomingAuthenticationClientCertRequest,omitempty"`
	// The handling of Authorization headers for incoming REST connections. The allowed values and their meaning are:  <pre> \"drop\" - Do not attach the Authorization header to the message as a user property. This configuration is most secure. \"forward\" - Forward the Authorization header, attaching it to the message as a user property in the same way as other headers. For best security, use the drop setting. \"legacy\" - If the Authorization header was used for authentication to the broker, do not attach it to the message. If the Authorization header was not used for authentication to the broker, attach it to the message as a user property in the same way as other headers. For best security, use the drop setting. </pre>  Available since 2.19.
	ServiceRestIncomingAuthorizationHeaderHandling *string `json:"serviceRestIncomingAuthorizationHeaderHandling,omitempty"`
	// The maximum number of REST incoming client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform.
	ServiceRestIncomingMaxConnectionCount *int64 `json:"serviceRestIncomingMaxConnectionCount,omitempty"`
	// Indicates whether the incoming REST Service is compressed in the Message VPN.
	ServiceRestIncomingPlainTextCompressed *bool `json:"serviceRestIncomingPlainTextCompressed,omitempty"`
	// Indicates whether the REST Service is enabled in the Message VPN for incoming clients.
	ServiceRestIncomingPlainTextEnabled *bool `json:"serviceRestIncomingPlainTextEnabled,omitempty"`
	// The reason for the incoming REST Service failure in the Message VPN.
	ServiceRestIncomingPlainTextFailureReason *string `json:"serviceRestIncomingPlainTextFailureReason,omitempty"`
	// The port number for incoming plain-text REST clients that connect to the Message VPN. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceRestIncomingPlainTextListenPort *int64 `json:"serviceRestIncomingPlainTextListenPort,omitempty"`
	// Indicates whether the incoming REST Service is operationally up in the Message VPN.
	ServiceRestIncomingPlainTextUp *bool `json:"serviceRestIncomingPlainTextUp,omitempty"`
	// Indicates whether the TLS related incoming REST Service is compressed in the Message VPN.
	ServiceRestIncomingTlsCompressed *bool `json:"serviceRestIncomingTlsCompressed,omitempty"`
	// Indicates whether encryption (TLS) is enabled for incoming REST clients in the Message VPN.
	ServiceRestIncomingTlsEnabled *bool `json:"serviceRestIncomingTlsEnabled,omitempty"`
	// The reason for the TLS related incoming REST Service failure in the Message VPN.
	ServiceRestIncomingTlsFailureReason *string `json:"serviceRestIncomingTlsFailureReason,omitempty"`
	// The port number for incoming REST clients that connect to the Message VPN over TLS. The port must be unique across the message backbone. A value of 0 means that the listen-port is unassigned and cannot be enabled.
	ServiceRestIncomingTlsListenPort *int64 `json:"serviceRestIncomingTlsListenPort,omitempty"`
	// Indicates whether the TLS related incoming REST Service is operationally up in the Message VPN.
	ServiceRestIncomingTlsUp *bool `json:"serviceRestIncomingTlsUp,omitempty"`
	// The REST service mode for incoming REST clients that connect to the Message VPN. The allowed values and their meaning are:  <pre> \"gateway\" - Act as a message gateway through which REST messages are propagated. \"messaging\" - Act as a message broker on which REST messages are queued. </pre>
	ServiceRestMode *string `json:"serviceRestMode,omitempty"`
	// The maximum number of REST Consumer (outgoing) client connections that can be simultaneously connected to the Message VPN.
	ServiceRestOutgoingMaxConnectionCount *int64 `json:"serviceRestOutgoingMaxConnectionCount,omitempty"`
	// The maximum number of SMF client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform.
	ServiceSmfMaxConnectionCount *int64 `json:"serviceSmfMaxConnectionCount,omitempty"`
	// Indicates whether the SMF Service is enabled in the Message VPN.
	ServiceSmfPlainTextEnabled *bool `json:"serviceSmfPlainTextEnabled,omitempty"`
	// The reason for the SMF Service failure in the Message VPN.
	ServiceSmfPlainTextFailureReason *string `json:"serviceSmfPlainTextFailureReason,omitempty"`
	// Indicates whether the SMF Service is operationally up in the Message VPN.
	ServiceSmfPlainTextUp *bool `json:"serviceSmfPlainTextUp,omitempty"`
	// Indicates whether encryption (TLS) is enabled for SMF clients in the Message VPN.
	ServiceSmfTlsEnabled *bool `json:"serviceSmfTlsEnabled,omitempty"`
	// The reason for the TLS related SMF Service failure in the Message VPN.
	ServiceSmfTlsFailureReason *string `json:"serviceSmfTlsFailureReason,omitempty"`
	// Indicates whether the TLS related SMF Service is operationally up in the Message VPN.
	ServiceSmfTlsUp *bool `json:"serviceSmfTlsUp,omitempty"`
	// Determines when to request a client certificate from a Web Transport client connecting via a TLS port. The allowed values and their meaning are:  <pre> \"always\" - Always ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"never\" - Never ask for a client certificate regardless of the \"message-vpn > authentication > client-certificate > shutdown\" configuration. \"when-enabled-in-message-vpn\" - Only ask for a client-certificate if client certificate authentication is enabled under \"message-vpn >  authentication > client-certificate > shutdown\". </pre>  Available since 2.21.
	ServiceWebAuthenticationClientCertRequest *string `json:"serviceWebAuthenticationClientCertRequest,omitempty"`
	// The maximum number of Web Transport client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the platform.
	ServiceWebMaxConnectionCount *int64 `json:"serviceWebMaxConnectionCount,omitempty"`
	// Indicates whether the Web transport for the SMF Service is enabled in the Message VPN.
	ServiceWebPlainTextEnabled *bool `json:"serviceWebPlainTextEnabled,omitempty"`
	// The reason for the Web transport related SMF Service failure in the Message VPN.
	ServiceWebPlainTextFailureReason *string `json:"serviceWebPlainTextFailureReason,omitempty"`
	// Indicates whether the Web transport for the SMF Service is operationally up in the Message VPN.
	ServiceWebPlainTextUp *bool `json:"serviceWebPlainTextUp,omitempty"`
	// Indicates whether TLS is enabled for SMF clients in the Message VPN that use the Web transport.
	ServiceWebTlsEnabled *bool `json:"serviceWebTlsEnabled,omitempty"`
	// The reason for the TLS related Web transport SMF Service failure in the Message VPN.
	ServiceWebTlsFailureReason *string `json:"serviceWebTlsFailureReason,omitempty"`
	// Indicates whether the TLS related Web transport SMF Service is operationally up in the Message VPN.
	ServiceWebTlsUp *bool `json:"serviceWebTlsUp,omitempty"`
	// The operational state of the local Message VPN. The allowed values and their meaning are:  <pre> \"up\" - The Message VPN is operationally up. \"down\" - The Message VPN is operationally down. \"standby\" - The Message VPN is operationally replication standby. </pre>
	State *string `json:"state,omitempty"`
	// The progress of the subscription export task, in percent complete.
	SubscriptionExportProgress *int64 `json:"subscriptionExportProgress,omitempty"`
	// Indicates whether the Message VPN is the system manager for handling system level SEMP get requests and system level event publishing.
	SystemManager *bool `json:"systemManager,omitempty"`
	// Indicates whether SMF clients connected to the Message VPN are allowed to downgrade their connections from TLS to plain text.
	TlsAllowDowngradeToPlainTextEnabled *bool `json:"tlsAllowDowngradeToPlainTextEnabled,omitempty"`
	// The one minute average of the TLS message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	TlsAverageRxByteRate *int64 `json:"tlsAverageRxByteRate,omitempty"`
	// The one minute average of the TLS message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	TlsAverageTxByteRate *int64 `json:"tlsAverageTxByteRate,omitempty"`
	// The amount of TLS messages received by the Message VPN, in bytes (B). Available since 2.13.
	TlsRxByteCount *int64 `json:"tlsRxByteCount,omitempty"`
	// The current TLS message rate received by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	TlsRxByteRate *int64 `json:"tlsRxByteRate,omitempty"`
	// The amount of TLS messages transmitted by the Message VPN, in bytes (B). Available since 2.13.
	TlsTxByteCount *int64 `json:"tlsTxByteCount,omitempty"`
	// The current TLS message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	TlsTxByteRate *int64 `json:"tlsTxByteRate,omitempty"`
	// The amount of messages transmitted to clients by the Message VPN, in bytes (B). Available since 2.12.
	TxByteCount *int64 `json:"txByteCount,omitempty"`
	// The current message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.13.
	TxByteRate *int64 `json:"txByteRate,omitempty"`
	// The amount of compressed messages transmitted by the Message VPN, in bytes (B). Available since 2.12.
	TxCompressedByteCount *int64 `json:"txCompressedByteCount,omitempty"`
	// The current compressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	TxCompressedByteRate *int64 `json:"txCompressedByteRate,omitempty"`
	// The compression ratio for messages transmitted by the message VPN. Available since 2.12.
	TxCompressionRatio *string `json:"txCompressionRatio,omitempty"`
	// The number of messages transmitted to clients by the Message VPN. Available since 2.12.
	TxMsgCount *int64 `json:"txMsgCount,omitempty"`
	// The current message rate transmitted by the Message VPN, in messages per second (msg/sec). Available since 2.13.
	TxMsgRate *int64 `json:"txMsgRate,omitempty"`
	// The amount of uncompressed messages transmitted by the Message VPN, in bytes (B). Available since 2.12.
	TxUncompressedByteCount *int64 `json:"txUncompressedByteCount,omitempty"`
	// The current uncompressed message rate transmitted by the Message VPN, in bytes per second (B/sec). Available since 2.12.
	TxUncompressedByteRate *int64 `json:"txUncompressedByteRate,omitempty"`
}

// NewMsgVpn instantiates a new MsgVpn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpn() *MsgVpn {
	this := MsgVpn{}
	return &this
}

// NewMsgVpnWithDefaults instantiates a new MsgVpn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnWithDefaults() *MsgVpn {
	this := MsgVpn{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *MsgVpn) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *MsgVpn) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *MsgVpn) SetAlias(v string) {
	o.Alias = &v
}

// GetAuthenticationBasicEnabled returns the AuthenticationBasicEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationBasicEnabled() bool {
	if o == nil || o.AuthenticationBasicEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthenticationBasicEnabled
}

// GetAuthenticationBasicEnabledOk returns a tuple with the AuthenticationBasicEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationBasicEnabledOk() (*bool, bool) {
	if o == nil || o.AuthenticationBasicEnabled == nil {
		return nil, false
	}
	return o.AuthenticationBasicEnabled, true
}

// HasAuthenticationBasicEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationBasicEnabled() bool {
	if o != nil && o.AuthenticationBasicEnabled != nil {
		return true
	}

	return false
}

// SetAuthenticationBasicEnabled gets a reference to the given bool and assigns it to the AuthenticationBasicEnabled field.
func (o *MsgVpn) SetAuthenticationBasicEnabled(v bool) {
	o.AuthenticationBasicEnabled = &v
}

// GetAuthenticationBasicProfileName returns the AuthenticationBasicProfileName field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationBasicProfileName() string {
	if o == nil || o.AuthenticationBasicProfileName == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationBasicProfileName
}

// GetAuthenticationBasicProfileNameOk returns a tuple with the AuthenticationBasicProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationBasicProfileNameOk() (*string, bool) {
	if o == nil || o.AuthenticationBasicProfileName == nil {
		return nil, false
	}
	return o.AuthenticationBasicProfileName, true
}

// HasAuthenticationBasicProfileName returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationBasicProfileName() bool {
	if o != nil && o.AuthenticationBasicProfileName != nil {
		return true
	}

	return false
}

// SetAuthenticationBasicProfileName gets a reference to the given string and assigns it to the AuthenticationBasicProfileName field.
func (o *MsgVpn) SetAuthenticationBasicProfileName(v string) {
	o.AuthenticationBasicProfileName = &v
}

// GetAuthenticationBasicRadiusDomain returns the AuthenticationBasicRadiusDomain field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationBasicRadiusDomain() string {
	if o == nil || o.AuthenticationBasicRadiusDomain == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationBasicRadiusDomain
}

// GetAuthenticationBasicRadiusDomainOk returns a tuple with the AuthenticationBasicRadiusDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationBasicRadiusDomainOk() (*string, bool) {
	if o == nil || o.AuthenticationBasicRadiusDomain == nil {
		return nil, false
	}
	return o.AuthenticationBasicRadiusDomain, true
}

// HasAuthenticationBasicRadiusDomain returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationBasicRadiusDomain() bool {
	if o != nil && o.AuthenticationBasicRadiusDomain != nil {
		return true
	}

	return false
}

// SetAuthenticationBasicRadiusDomain gets a reference to the given string and assigns it to the AuthenticationBasicRadiusDomain field.
func (o *MsgVpn) SetAuthenticationBasicRadiusDomain(v string) {
	o.AuthenticationBasicRadiusDomain = &v
}

// GetAuthenticationBasicType returns the AuthenticationBasicType field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationBasicType() string {
	if o == nil || o.AuthenticationBasicType == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationBasicType
}

// GetAuthenticationBasicTypeOk returns a tuple with the AuthenticationBasicType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationBasicTypeOk() (*string, bool) {
	if o == nil || o.AuthenticationBasicType == nil {
		return nil, false
	}
	return o.AuthenticationBasicType, true
}

// HasAuthenticationBasicType returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationBasicType() bool {
	if o != nil && o.AuthenticationBasicType != nil {
		return true
	}

	return false
}

// SetAuthenticationBasicType gets a reference to the given string and assigns it to the AuthenticationBasicType field.
func (o *MsgVpn) SetAuthenticationBasicType(v string) {
	o.AuthenticationBasicType = &v
}

// GetAuthenticationClientCertAllowApiProvidedUsernameEnabled returns the AuthenticationClientCertAllowApiProvidedUsernameEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationClientCertAllowApiProvidedUsernameEnabled() bool {
	if o == nil || o.AuthenticationClientCertAllowApiProvidedUsernameEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthenticationClientCertAllowApiProvidedUsernameEnabled
}

// GetAuthenticationClientCertAllowApiProvidedUsernameEnabledOk returns a tuple with the AuthenticationClientCertAllowApiProvidedUsernameEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationClientCertAllowApiProvidedUsernameEnabledOk() (*bool, bool) {
	if o == nil || o.AuthenticationClientCertAllowApiProvidedUsernameEnabled == nil {
		return nil, false
	}
	return o.AuthenticationClientCertAllowApiProvidedUsernameEnabled, true
}

// HasAuthenticationClientCertAllowApiProvidedUsernameEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationClientCertAllowApiProvidedUsernameEnabled() bool {
	if o != nil && o.AuthenticationClientCertAllowApiProvidedUsernameEnabled != nil {
		return true
	}

	return false
}

// SetAuthenticationClientCertAllowApiProvidedUsernameEnabled gets a reference to the given bool and assigns it to the AuthenticationClientCertAllowApiProvidedUsernameEnabled field.
func (o *MsgVpn) SetAuthenticationClientCertAllowApiProvidedUsernameEnabled(v bool) {
	o.AuthenticationClientCertAllowApiProvidedUsernameEnabled = &v
}

// GetAuthenticationClientCertEnabled returns the AuthenticationClientCertEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationClientCertEnabled() bool {
	if o == nil || o.AuthenticationClientCertEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthenticationClientCertEnabled
}

// GetAuthenticationClientCertEnabledOk returns a tuple with the AuthenticationClientCertEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationClientCertEnabledOk() (*bool, bool) {
	if o == nil || o.AuthenticationClientCertEnabled == nil {
		return nil, false
	}
	return o.AuthenticationClientCertEnabled, true
}

// HasAuthenticationClientCertEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationClientCertEnabled() bool {
	if o != nil && o.AuthenticationClientCertEnabled != nil {
		return true
	}

	return false
}

// SetAuthenticationClientCertEnabled gets a reference to the given bool and assigns it to the AuthenticationClientCertEnabled field.
func (o *MsgVpn) SetAuthenticationClientCertEnabled(v bool) {
	o.AuthenticationClientCertEnabled = &v
}

// GetAuthenticationClientCertMaxChainDepth returns the AuthenticationClientCertMaxChainDepth field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationClientCertMaxChainDepth() int64 {
	if o == nil || o.AuthenticationClientCertMaxChainDepth == nil {
		var ret int64
		return ret
	}
	return *o.AuthenticationClientCertMaxChainDepth
}

// GetAuthenticationClientCertMaxChainDepthOk returns a tuple with the AuthenticationClientCertMaxChainDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationClientCertMaxChainDepthOk() (*int64, bool) {
	if o == nil || o.AuthenticationClientCertMaxChainDepth == nil {
		return nil, false
	}
	return o.AuthenticationClientCertMaxChainDepth, true
}

// HasAuthenticationClientCertMaxChainDepth returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationClientCertMaxChainDepth() bool {
	if o != nil && o.AuthenticationClientCertMaxChainDepth != nil {
		return true
	}

	return false
}

// SetAuthenticationClientCertMaxChainDepth gets a reference to the given int64 and assigns it to the AuthenticationClientCertMaxChainDepth field.
func (o *MsgVpn) SetAuthenticationClientCertMaxChainDepth(v int64) {
	o.AuthenticationClientCertMaxChainDepth = &v
}

// GetAuthenticationClientCertRevocationCheckMode returns the AuthenticationClientCertRevocationCheckMode field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationClientCertRevocationCheckMode() string {
	if o == nil || o.AuthenticationClientCertRevocationCheckMode == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationClientCertRevocationCheckMode
}

// GetAuthenticationClientCertRevocationCheckModeOk returns a tuple with the AuthenticationClientCertRevocationCheckMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationClientCertRevocationCheckModeOk() (*string, bool) {
	if o == nil || o.AuthenticationClientCertRevocationCheckMode == nil {
		return nil, false
	}
	return o.AuthenticationClientCertRevocationCheckMode, true
}

// HasAuthenticationClientCertRevocationCheckMode returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationClientCertRevocationCheckMode() bool {
	if o != nil && o.AuthenticationClientCertRevocationCheckMode != nil {
		return true
	}

	return false
}

// SetAuthenticationClientCertRevocationCheckMode gets a reference to the given string and assigns it to the AuthenticationClientCertRevocationCheckMode field.
func (o *MsgVpn) SetAuthenticationClientCertRevocationCheckMode(v string) {
	o.AuthenticationClientCertRevocationCheckMode = &v
}

// GetAuthenticationClientCertUsernameSource returns the AuthenticationClientCertUsernameSource field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationClientCertUsernameSource() string {
	if o == nil || o.AuthenticationClientCertUsernameSource == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationClientCertUsernameSource
}

// GetAuthenticationClientCertUsernameSourceOk returns a tuple with the AuthenticationClientCertUsernameSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationClientCertUsernameSourceOk() (*string, bool) {
	if o == nil || o.AuthenticationClientCertUsernameSource == nil {
		return nil, false
	}
	return o.AuthenticationClientCertUsernameSource, true
}

// HasAuthenticationClientCertUsernameSource returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationClientCertUsernameSource() bool {
	if o != nil && o.AuthenticationClientCertUsernameSource != nil {
		return true
	}

	return false
}

// SetAuthenticationClientCertUsernameSource gets a reference to the given string and assigns it to the AuthenticationClientCertUsernameSource field.
func (o *MsgVpn) SetAuthenticationClientCertUsernameSource(v string) {
	o.AuthenticationClientCertUsernameSource = &v
}

// GetAuthenticationClientCertValidateDateEnabled returns the AuthenticationClientCertValidateDateEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationClientCertValidateDateEnabled() bool {
	if o == nil || o.AuthenticationClientCertValidateDateEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthenticationClientCertValidateDateEnabled
}

// GetAuthenticationClientCertValidateDateEnabledOk returns a tuple with the AuthenticationClientCertValidateDateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationClientCertValidateDateEnabledOk() (*bool, bool) {
	if o == nil || o.AuthenticationClientCertValidateDateEnabled == nil {
		return nil, false
	}
	return o.AuthenticationClientCertValidateDateEnabled, true
}

// HasAuthenticationClientCertValidateDateEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationClientCertValidateDateEnabled() bool {
	if o != nil && o.AuthenticationClientCertValidateDateEnabled != nil {
		return true
	}

	return false
}

// SetAuthenticationClientCertValidateDateEnabled gets a reference to the given bool and assigns it to the AuthenticationClientCertValidateDateEnabled field.
func (o *MsgVpn) SetAuthenticationClientCertValidateDateEnabled(v bool) {
	o.AuthenticationClientCertValidateDateEnabled = &v
}

// GetAuthenticationKerberosAllowApiProvidedUsernameEnabled returns the AuthenticationKerberosAllowApiProvidedUsernameEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationKerberosAllowApiProvidedUsernameEnabled() bool {
	if o == nil || o.AuthenticationKerberosAllowApiProvidedUsernameEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthenticationKerberosAllowApiProvidedUsernameEnabled
}

// GetAuthenticationKerberosAllowApiProvidedUsernameEnabledOk returns a tuple with the AuthenticationKerberosAllowApiProvidedUsernameEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationKerberosAllowApiProvidedUsernameEnabledOk() (*bool, bool) {
	if o == nil || o.AuthenticationKerberosAllowApiProvidedUsernameEnabled == nil {
		return nil, false
	}
	return o.AuthenticationKerberosAllowApiProvidedUsernameEnabled, true
}

// HasAuthenticationKerberosAllowApiProvidedUsernameEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationKerberosAllowApiProvidedUsernameEnabled() bool {
	if o != nil && o.AuthenticationKerberosAllowApiProvidedUsernameEnabled != nil {
		return true
	}

	return false
}

// SetAuthenticationKerberosAllowApiProvidedUsernameEnabled gets a reference to the given bool and assigns it to the AuthenticationKerberosAllowApiProvidedUsernameEnabled field.
func (o *MsgVpn) SetAuthenticationKerberosAllowApiProvidedUsernameEnabled(v bool) {
	o.AuthenticationKerberosAllowApiProvidedUsernameEnabled = &v
}

// GetAuthenticationKerberosEnabled returns the AuthenticationKerberosEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationKerberosEnabled() bool {
	if o == nil || o.AuthenticationKerberosEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthenticationKerberosEnabled
}

// GetAuthenticationKerberosEnabledOk returns a tuple with the AuthenticationKerberosEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationKerberosEnabledOk() (*bool, bool) {
	if o == nil || o.AuthenticationKerberosEnabled == nil {
		return nil, false
	}
	return o.AuthenticationKerberosEnabled, true
}

// HasAuthenticationKerberosEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationKerberosEnabled() bool {
	if o != nil && o.AuthenticationKerberosEnabled != nil {
		return true
	}

	return false
}

// SetAuthenticationKerberosEnabled gets a reference to the given bool and assigns it to the AuthenticationKerberosEnabled field.
func (o *MsgVpn) SetAuthenticationKerberosEnabled(v bool) {
	o.AuthenticationKerberosEnabled = &v
}

// GetAuthenticationOauthDefaultProviderName returns the AuthenticationOauthDefaultProviderName field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationOauthDefaultProviderName() string {
	if o == nil || o.AuthenticationOauthDefaultProviderName == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationOauthDefaultProviderName
}

// GetAuthenticationOauthDefaultProviderNameOk returns a tuple with the AuthenticationOauthDefaultProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationOauthDefaultProviderNameOk() (*string, bool) {
	if o == nil || o.AuthenticationOauthDefaultProviderName == nil {
		return nil, false
	}
	return o.AuthenticationOauthDefaultProviderName, true
}

// HasAuthenticationOauthDefaultProviderName returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationOauthDefaultProviderName() bool {
	if o != nil && o.AuthenticationOauthDefaultProviderName != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthDefaultProviderName gets a reference to the given string and assigns it to the AuthenticationOauthDefaultProviderName field.
func (o *MsgVpn) SetAuthenticationOauthDefaultProviderName(v string) {
	o.AuthenticationOauthDefaultProviderName = &v
}

// GetAuthenticationOauthEnabled returns the AuthenticationOauthEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthenticationOauthEnabled() bool {
	if o == nil || o.AuthenticationOauthEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthenticationOauthEnabled
}

// GetAuthenticationOauthEnabledOk returns a tuple with the AuthenticationOauthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthenticationOauthEnabledOk() (*bool, bool) {
	if o == nil || o.AuthenticationOauthEnabled == nil {
		return nil, false
	}
	return o.AuthenticationOauthEnabled, true
}

// HasAuthenticationOauthEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthenticationOauthEnabled() bool {
	if o != nil && o.AuthenticationOauthEnabled != nil {
		return true
	}

	return false
}

// SetAuthenticationOauthEnabled gets a reference to the given bool and assigns it to the AuthenticationOauthEnabled field.
func (o *MsgVpn) SetAuthenticationOauthEnabled(v bool) {
	o.AuthenticationOauthEnabled = &v
}

// GetAuthorizationLdapGroupMembershipAttributeName returns the AuthorizationLdapGroupMembershipAttributeName field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthorizationLdapGroupMembershipAttributeName() string {
	if o == nil || o.AuthorizationLdapGroupMembershipAttributeName == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationLdapGroupMembershipAttributeName
}

// GetAuthorizationLdapGroupMembershipAttributeNameOk returns a tuple with the AuthorizationLdapGroupMembershipAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthorizationLdapGroupMembershipAttributeNameOk() (*string, bool) {
	if o == nil || o.AuthorizationLdapGroupMembershipAttributeName == nil {
		return nil, false
	}
	return o.AuthorizationLdapGroupMembershipAttributeName, true
}

// HasAuthorizationLdapGroupMembershipAttributeName returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthorizationLdapGroupMembershipAttributeName() bool {
	if o != nil && o.AuthorizationLdapGroupMembershipAttributeName != nil {
		return true
	}

	return false
}

// SetAuthorizationLdapGroupMembershipAttributeName gets a reference to the given string and assigns it to the AuthorizationLdapGroupMembershipAttributeName field.
func (o *MsgVpn) SetAuthorizationLdapGroupMembershipAttributeName(v string) {
	o.AuthorizationLdapGroupMembershipAttributeName = &v
}

// GetAuthorizationLdapTrimClientUsernameDomainEnabled returns the AuthorizationLdapTrimClientUsernameDomainEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthorizationLdapTrimClientUsernameDomainEnabled() bool {
	if o == nil || o.AuthorizationLdapTrimClientUsernameDomainEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthorizationLdapTrimClientUsernameDomainEnabled
}

// GetAuthorizationLdapTrimClientUsernameDomainEnabledOk returns a tuple with the AuthorizationLdapTrimClientUsernameDomainEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthorizationLdapTrimClientUsernameDomainEnabledOk() (*bool, bool) {
	if o == nil || o.AuthorizationLdapTrimClientUsernameDomainEnabled == nil {
		return nil, false
	}
	return o.AuthorizationLdapTrimClientUsernameDomainEnabled, true
}

// HasAuthorizationLdapTrimClientUsernameDomainEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthorizationLdapTrimClientUsernameDomainEnabled() bool {
	if o != nil && o.AuthorizationLdapTrimClientUsernameDomainEnabled != nil {
		return true
	}

	return false
}

// SetAuthorizationLdapTrimClientUsernameDomainEnabled gets a reference to the given bool and assigns it to the AuthorizationLdapTrimClientUsernameDomainEnabled field.
func (o *MsgVpn) SetAuthorizationLdapTrimClientUsernameDomainEnabled(v bool) {
	o.AuthorizationLdapTrimClientUsernameDomainEnabled = &v
}

// GetAuthorizationProfileName returns the AuthorizationProfileName field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthorizationProfileName() string {
	if o == nil || o.AuthorizationProfileName == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationProfileName
}

// GetAuthorizationProfileNameOk returns a tuple with the AuthorizationProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthorizationProfileNameOk() (*string, bool) {
	if o == nil || o.AuthorizationProfileName == nil {
		return nil, false
	}
	return o.AuthorizationProfileName, true
}

// HasAuthorizationProfileName returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthorizationProfileName() bool {
	if o != nil && o.AuthorizationProfileName != nil {
		return true
	}

	return false
}

// SetAuthorizationProfileName gets a reference to the given string and assigns it to the AuthorizationProfileName field.
func (o *MsgVpn) SetAuthorizationProfileName(v string) {
	o.AuthorizationProfileName = &v
}

// GetAuthorizationType returns the AuthorizationType field value if set, zero value otherwise.
func (o *MsgVpn) GetAuthorizationType() string {
	if o == nil || o.AuthorizationType == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationType
}

// GetAuthorizationTypeOk returns a tuple with the AuthorizationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAuthorizationTypeOk() (*string, bool) {
	if o == nil || o.AuthorizationType == nil {
		return nil, false
	}
	return o.AuthorizationType, true
}

// HasAuthorizationType returns a boolean if a field has been set.
func (o *MsgVpn) HasAuthorizationType() bool {
	if o != nil && o.AuthorizationType != nil {
		return true
	}

	return false
}

// SetAuthorizationType gets a reference to the given string and assigns it to the AuthorizationType field.
func (o *MsgVpn) SetAuthorizationType(v string) {
	o.AuthorizationType = &v
}

// GetAverageRxByteRate returns the AverageRxByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetAverageRxByteRate() int64 {
	if o == nil || o.AverageRxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRxByteRate
}

// GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAverageRxByteRateOk() (*int64, bool) {
	if o == nil || o.AverageRxByteRate == nil {
		return nil, false
	}
	return o.AverageRxByteRate, true
}

// HasAverageRxByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasAverageRxByteRate() bool {
	if o != nil && o.AverageRxByteRate != nil {
		return true
	}

	return false
}

// SetAverageRxByteRate gets a reference to the given int64 and assigns it to the AverageRxByteRate field.
func (o *MsgVpn) SetAverageRxByteRate(v int64) {
	o.AverageRxByteRate = &v
}

// GetAverageRxCompressedByteRate returns the AverageRxCompressedByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetAverageRxCompressedByteRate() int64 {
	if o == nil || o.AverageRxCompressedByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRxCompressedByteRate
}

// GetAverageRxCompressedByteRateOk returns a tuple with the AverageRxCompressedByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAverageRxCompressedByteRateOk() (*int64, bool) {
	if o == nil || o.AverageRxCompressedByteRate == nil {
		return nil, false
	}
	return o.AverageRxCompressedByteRate, true
}

// HasAverageRxCompressedByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasAverageRxCompressedByteRate() bool {
	if o != nil && o.AverageRxCompressedByteRate != nil {
		return true
	}

	return false
}

// SetAverageRxCompressedByteRate gets a reference to the given int64 and assigns it to the AverageRxCompressedByteRate field.
func (o *MsgVpn) SetAverageRxCompressedByteRate(v int64) {
	o.AverageRxCompressedByteRate = &v
}

// GetAverageRxMsgRate returns the AverageRxMsgRate field value if set, zero value otherwise.
func (o *MsgVpn) GetAverageRxMsgRate() int64 {
	if o == nil || o.AverageRxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRxMsgRate
}

// GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAverageRxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageRxMsgRate == nil {
		return nil, false
	}
	return o.AverageRxMsgRate, true
}

// HasAverageRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpn) HasAverageRxMsgRate() bool {
	if o != nil && o.AverageRxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageRxMsgRate gets a reference to the given int64 and assigns it to the AverageRxMsgRate field.
func (o *MsgVpn) SetAverageRxMsgRate(v int64) {
	o.AverageRxMsgRate = &v
}

// GetAverageRxUncompressedByteRate returns the AverageRxUncompressedByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetAverageRxUncompressedByteRate() int64 {
	if o == nil || o.AverageRxUncompressedByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRxUncompressedByteRate
}

// GetAverageRxUncompressedByteRateOk returns a tuple with the AverageRxUncompressedByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAverageRxUncompressedByteRateOk() (*int64, bool) {
	if o == nil || o.AverageRxUncompressedByteRate == nil {
		return nil, false
	}
	return o.AverageRxUncompressedByteRate, true
}

// HasAverageRxUncompressedByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasAverageRxUncompressedByteRate() bool {
	if o != nil && o.AverageRxUncompressedByteRate != nil {
		return true
	}

	return false
}

// SetAverageRxUncompressedByteRate gets a reference to the given int64 and assigns it to the AverageRxUncompressedByteRate field.
func (o *MsgVpn) SetAverageRxUncompressedByteRate(v int64) {
	o.AverageRxUncompressedByteRate = &v
}

// GetAverageTxByteRate returns the AverageTxByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetAverageTxByteRate() int64 {
	if o == nil || o.AverageTxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageTxByteRate
}

// GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAverageTxByteRateOk() (*int64, bool) {
	if o == nil || o.AverageTxByteRate == nil {
		return nil, false
	}
	return o.AverageTxByteRate, true
}

// HasAverageTxByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasAverageTxByteRate() bool {
	if o != nil && o.AverageTxByteRate != nil {
		return true
	}

	return false
}

// SetAverageTxByteRate gets a reference to the given int64 and assigns it to the AverageTxByteRate field.
func (o *MsgVpn) SetAverageTxByteRate(v int64) {
	o.AverageTxByteRate = &v
}

// GetAverageTxCompressedByteRate returns the AverageTxCompressedByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetAverageTxCompressedByteRate() int64 {
	if o == nil || o.AverageTxCompressedByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageTxCompressedByteRate
}

// GetAverageTxCompressedByteRateOk returns a tuple with the AverageTxCompressedByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAverageTxCompressedByteRateOk() (*int64, bool) {
	if o == nil || o.AverageTxCompressedByteRate == nil {
		return nil, false
	}
	return o.AverageTxCompressedByteRate, true
}

// HasAverageTxCompressedByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasAverageTxCompressedByteRate() bool {
	if o != nil && o.AverageTxCompressedByteRate != nil {
		return true
	}

	return false
}

// SetAverageTxCompressedByteRate gets a reference to the given int64 and assigns it to the AverageTxCompressedByteRate field.
func (o *MsgVpn) SetAverageTxCompressedByteRate(v int64) {
	o.AverageTxCompressedByteRate = &v
}

// GetAverageTxMsgRate returns the AverageTxMsgRate field value if set, zero value otherwise.
func (o *MsgVpn) GetAverageTxMsgRate() int64 {
	if o == nil || o.AverageTxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageTxMsgRate
}

// GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAverageTxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageTxMsgRate == nil {
		return nil, false
	}
	return o.AverageTxMsgRate, true
}

// HasAverageTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpn) HasAverageTxMsgRate() bool {
	if o != nil && o.AverageTxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageTxMsgRate gets a reference to the given int64 and assigns it to the AverageTxMsgRate field.
func (o *MsgVpn) SetAverageTxMsgRate(v int64) {
	o.AverageTxMsgRate = &v
}

// GetAverageTxUncompressedByteRate returns the AverageTxUncompressedByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetAverageTxUncompressedByteRate() int64 {
	if o == nil || o.AverageTxUncompressedByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageTxUncompressedByteRate
}

// GetAverageTxUncompressedByteRateOk returns a tuple with the AverageTxUncompressedByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetAverageTxUncompressedByteRateOk() (*int64, bool) {
	if o == nil || o.AverageTxUncompressedByteRate == nil {
		return nil, false
	}
	return o.AverageTxUncompressedByteRate, true
}

// HasAverageTxUncompressedByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasAverageTxUncompressedByteRate() bool {
	if o != nil && o.AverageTxUncompressedByteRate != nil {
		return true
	}

	return false
}

// SetAverageTxUncompressedByteRate gets a reference to the given int64 and assigns it to the AverageTxUncompressedByteRate field.
func (o *MsgVpn) SetAverageTxUncompressedByteRate(v int64) {
	o.AverageTxUncompressedByteRate = &v
}

// GetBridgingTlsServerCertEnforceTrustedCommonNameEnabled returns the BridgingTlsServerCertEnforceTrustedCommonNameEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetBridgingTlsServerCertEnforceTrustedCommonNameEnabled() bool {
	if o == nil || o.BridgingTlsServerCertEnforceTrustedCommonNameEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BridgingTlsServerCertEnforceTrustedCommonNameEnabled
}

// GetBridgingTlsServerCertEnforceTrustedCommonNameEnabledOk returns a tuple with the BridgingTlsServerCertEnforceTrustedCommonNameEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetBridgingTlsServerCertEnforceTrustedCommonNameEnabledOk() (*bool, bool) {
	if o == nil || o.BridgingTlsServerCertEnforceTrustedCommonNameEnabled == nil {
		return nil, false
	}
	return o.BridgingTlsServerCertEnforceTrustedCommonNameEnabled, true
}

// HasBridgingTlsServerCertEnforceTrustedCommonNameEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasBridgingTlsServerCertEnforceTrustedCommonNameEnabled() bool {
	if o != nil && o.BridgingTlsServerCertEnforceTrustedCommonNameEnabled != nil {
		return true
	}

	return false
}

// SetBridgingTlsServerCertEnforceTrustedCommonNameEnabled gets a reference to the given bool and assigns it to the BridgingTlsServerCertEnforceTrustedCommonNameEnabled field.
func (o *MsgVpn) SetBridgingTlsServerCertEnforceTrustedCommonNameEnabled(v bool) {
	o.BridgingTlsServerCertEnforceTrustedCommonNameEnabled = &v
}

// GetBridgingTlsServerCertMaxChainDepth returns the BridgingTlsServerCertMaxChainDepth field value if set, zero value otherwise.
func (o *MsgVpn) GetBridgingTlsServerCertMaxChainDepth() int64 {
	if o == nil || o.BridgingTlsServerCertMaxChainDepth == nil {
		var ret int64
		return ret
	}
	return *o.BridgingTlsServerCertMaxChainDepth
}

// GetBridgingTlsServerCertMaxChainDepthOk returns a tuple with the BridgingTlsServerCertMaxChainDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetBridgingTlsServerCertMaxChainDepthOk() (*int64, bool) {
	if o == nil || o.BridgingTlsServerCertMaxChainDepth == nil {
		return nil, false
	}
	return o.BridgingTlsServerCertMaxChainDepth, true
}

// HasBridgingTlsServerCertMaxChainDepth returns a boolean if a field has been set.
func (o *MsgVpn) HasBridgingTlsServerCertMaxChainDepth() bool {
	if o != nil && o.BridgingTlsServerCertMaxChainDepth != nil {
		return true
	}

	return false
}

// SetBridgingTlsServerCertMaxChainDepth gets a reference to the given int64 and assigns it to the BridgingTlsServerCertMaxChainDepth field.
func (o *MsgVpn) SetBridgingTlsServerCertMaxChainDepth(v int64) {
	o.BridgingTlsServerCertMaxChainDepth = &v
}

// GetBridgingTlsServerCertValidateDateEnabled returns the BridgingTlsServerCertValidateDateEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetBridgingTlsServerCertValidateDateEnabled() bool {
	if o == nil || o.BridgingTlsServerCertValidateDateEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BridgingTlsServerCertValidateDateEnabled
}

// GetBridgingTlsServerCertValidateDateEnabledOk returns a tuple with the BridgingTlsServerCertValidateDateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetBridgingTlsServerCertValidateDateEnabledOk() (*bool, bool) {
	if o == nil || o.BridgingTlsServerCertValidateDateEnabled == nil {
		return nil, false
	}
	return o.BridgingTlsServerCertValidateDateEnabled, true
}

// HasBridgingTlsServerCertValidateDateEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasBridgingTlsServerCertValidateDateEnabled() bool {
	if o != nil && o.BridgingTlsServerCertValidateDateEnabled != nil {
		return true
	}

	return false
}

// SetBridgingTlsServerCertValidateDateEnabled gets a reference to the given bool and assigns it to the BridgingTlsServerCertValidateDateEnabled field.
func (o *MsgVpn) SetBridgingTlsServerCertValidateDateEnabled(v bool) {
	o.BridgingTlsServerCertValidateDateEnabled = &v
}

// GetBridgingTlsServerCertValidateNameEnabled returns the BridgingTlsServerCertValidateNameEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetBridgingTlsServerCertValidateNameEnabled() bool {
	if o == nil || o.BridgingTlsServerCertValidateNameEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BridgingTlsServerCertValidateNameEnabled
}

// GetBridgingTlsServerCertValidateNameEnabledOk returns a tuple with the BridgingTlsServerCertValidateNameEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetBridgingTlsServerCertValidateNameEnabledOk() (*bool, bool) {
	if o == nil || o.BridgingTlsServerCertValidateNameEnabled == nil {
		return nil, false
	}
	return o.BridgingTlsServerCertValidateNameEnabled, true
}

// HasBridgingTlsServerCertValidateNameEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasBridgingTlsServerCertValidateNameEnabled() bool {
	if o != nil && o.BridgingTlsServerCertValidateNameEnabled != nil {
		return true
	}

	return false
}

// SetBridgingTlsServerCertValidateNameEnabled gets a reference to the given bool and assigns it to the BridgingTlsServerCertValidateNameEnabled field.
func (o *MsgVpn) SetBridgingTlsServerCertValidateNameEnabled(v bool) {
	o.BridgingTlsServerCertValidateNameEnabled = &v
}

// GetConfigSyncLocalKey returns the ConfigSyncLocalKey field value if set, zero value otherwise.
func (o *MsgVpn) GetConfigSyncLocalKey() string {
	if o == nil || o.ConfigSyncLocalKey == nil {
		var ret string
		return ret
	}
	return *o.ConfigSyncLocalKey
}

// GetConfigSyncLocalKeyOk returns a tuple with the ConfigSyncLocalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetConfigSyncLocalKeyOk() (*string, bool) {
	if o == nil || o.ConfigSyncLocalKey == nil {
		return nil, false
	}
	return o.ConfigSyncLocalKey, true
}

// HasConfigSyncLocalKey returns a boolean if a field has been set.
func (o *MsgVpn) HasConfigSyncLocalKey() bool {
	if o != nil && o.ConfigSyncLocalKey != nil {
		return true
	}

	return false
}

// SetConfigSyncLocalKey gets a reference to the given string and assigns it to the ConfigSyncLocalKey field.
func (o *MsgVpn) SetConfigSyncLocalKey(v string) {
	o.ConfigSyncLocalKey = &v
}

// GetConfigSyncLocalLastResult returns the ConfigSyncLocalLastResult field value if set, zero value otherwise.
func (o *MsgVpn) GetConfigSyncLocalLastResult() string {
	if o == nil || o.ConfigSyncLocalLastResult == nil {
		var ret string
		return ret
	}
	return *o.ConfigSyncLocalLastResult
}

// GetConfigSyncLocalLastResultOk returns a tuple with the ConfigSyncLocalLastResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetConfigSyncLocalLastResultOk() (*string, bool) {
	if o == nil || o.ConfigSyncLocalLastResult == nil {
		return nil, false
	}
	return o.ConfigSyncLocalLastResult, true
}

// HasConfigSyncLocalLastResult returns a boolean if a field has been set.
func (o *MsgVpn) HasConfigSyncLocalLastResult() bool {
	if o != nil && o.ConfigSyncLocalLastResult != nil {
		return true
	}

	return false
}

// SetConfigSyncLocalLastResult gets a reference to the given string and assigns it to the ConfigSyncLocalLastResult field.
func (o *MsgVpn) SetConfigSyncLocalLastResult(v string) {
	o.ConfigSyncLocalLastResult = &v
}

// GetConfigSyncLocalRole returns the ConfigSyncLocalRole field value if set, zero value otherwise.
func (o *MsgVpn) GetConfigSyncLocalRole() string {
	if o == nil || o.ConfigSyncLocalRole == nil {
		var ret string
		return ret
	}
	return *o.ConfigSyncLocalRole
}

// GetConfigSyncLocalRoleOk returns a tuple with the ConfigSyncLocalRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetConfigSyncLocalRoleOk() (*string, bool) {
	if o == nil || o.ConfigSyncLocalRole == nil {
		return nil, false
	}
	return o.ConfigSyncLocalRole, true
}

// HasConfigSyncLocalRole returns a boolean if a field has been set.
func (o *MsgVpn) HasConfigSyncLocalRole() bool {
	if o != nil && o.ConfigSyncLocalRole != nil {
		return true
	}

	return false
}

// SetConfigSyncLocalRole gets a reference to the given string and assigns it to the ConfigSyncLocalRole field.
func (o *MsgVpn) SetConfigSyncLocalRole(v string) {
	o.ConfigSyncLocalRole = &v
}

// GetConfigSyncLocalState returns the ConfigSyncLocalState field value if set, zero value otherwise.
func (o *MsgVpn) GetConfigSyncLocalState() string {
	if o == nil || o.ConfigSyncLocalState == nil {
		var ret string
		return ret
	}
	return *o.ConfigSyncLocalState
}

// GetConfigSyncLocalStateOk returns a tuple with the ConfigSyncLocalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetConfigSyncLocalStateOk() (*string, bool) {
	if o == nil || o.ConfigSyncLocalState == nil {
		return nil, false
	}
	return o.ConfigSyncLocalState, true
}

// HasConfigSyncLocalState returns a boolean if a field has been set.
func (o *MsgVpn) HasConfigSyncLocalState() bool {
	if o != nil && o.ConfigSyncLocalState != nil {
		return true
	}

	return false
}

// SetConfigSyncLocalState gets a reference to the given string and assigns it to the ConfigSyncLocalState field.
func (o *MsgVpn) SetConfigSyncLocalState(v string) {
	o.ConfigSyncLocalState = &v
}

// GetConfigSyncLocalTimeInState returns the ConfigSyncLocalTimeInState field value if set, zero value otherwise.
func (o *MsgVpn) GetConfigSyncLocalTimeInState() int32 {
	if o == nil || o.ConfigSyncLocalTimeInState == nil {
		var ret int32
		return ret
	}
	return *o.ConfigSyncLocalTimeInState
}

// GetConfigSyncLocalTimeInStateOk returns a tuple with the ConfigSyncLocalTimeInState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetConfigSyncLocalTimeInStateOk() (*int32, bool) {
	if o == nil || o.ConfigSyncLocalTimeInState == nil {
		return nil, false
	}
	return o.ConfigSyncLocalTimeInState, true
}

// HasConfigSyncLocalTimeInState returns a boolean if a field has been set.
func (o *MsgVpn) HasConfigSyncLocalTimeInState() bool {
	if o != nil && o.ConfigSyncLocalTimeInState != nil {
		return true
	}

	return false
}

// SetConfigSyncLocalTimeInState gets a reference to the given int32 and assigns it to the ConfigSyncLocalTimeInState field.
func (o *MsgVpn) SetConfigSyncLocalTimeInState(v int32) {
	o.ConfigSyncLocalTimeInState = &v
}

// GetControlRxByteCount returns the ControlRxByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetControlRxByteCount() int64 {
	if o == nil || o.ControlRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlRxByteCount
}

// GetControlRxByteCountOk returns a tuple with the ControlRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetControlRxByteCountOk() (*int64, bool) {
	if o == nil || o.ControlRxByteCount == nil {
		return nil, false
	}
	return o.ControlRxByteCount, true
}

// HasControlRxByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasControlRxByteCount() bool {
	if o != nil && o.ControlRxByteCount != nil {
		return true
	}

	return false
}

// SetControlRxByteCount gets a reference to the given int64 and assigns it to the ControlRxByteCount field.
func (o *MsgVpn) SetControlRxByteCount(v int64) {
	o.ControlRxByteCount = &v
}

// GetControlRxMsgCount returns the ControlRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetControlRxMsgCount() int64 {
	if o == nil || o.ControlRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlRxMsgCount
}

// GetControlRxMsgCountOk returns a tuple with the ControlRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetControlRxMsgCountOk() (*int64, bool) {
	if o == nil || o.ControlRxMsgCount == nil {
		return nil, false
	}
	return o.ControlRxMsgCount, true
}

// HasControlRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasControlRxMsgCount() bool {
	if o != nil && o.ControlRxMsgCount != nil {
		return true
	}

	return false
}

// SetControlRxMsgCount gets a reference to the given int64 and assigns it to the ControlRxMsgCount field.
func (o *MsgVpn) SetControlRxMsgCount(v int64) {
	o.ControlRxMsgCount = &v
}

// GetControlTxByteCount returns the ControlTxByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetControlTxByteCount() int64 {
	if o == nil || o.ControlTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlTxByteCount
}

// GetControlTxByteCountOk returns a tuple with the ControlTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetControlTxByteCountOk() (*int64, bool) {
	if o == nil || o.ControlTxByteCount == nil {
		return nil, false
	}
	return o.ControlTxByteCount, true
}

// HasControlTxByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasControlTxByteCount() bool {
	if o != nil && o.ControlTxByteCount != nil {
		return true
	}

	return false
}

// SetControlTxByteCount gets a reference to the given int64 and assigns it to the ControlTxByteCount field.
func (o *MsgVpn) SetControlTxByteCount(v int64) {
	o.ControlTxByteCount = &v
}

// GetControlTxMsgCount returns the ControlTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetControlTxMsgCount() int64 {
	if o == nil || o.ControlTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ControlTxMsgCount
}

// GetControlTxMsgCountOk returns a tuple with the ControlTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetControlTxMsgCountOk() (*int64, bool) {
	if o == nil || o.ControlTxMsgCount == nil {
		return nil, false
	}
	return o.ControlTxMsgCount, true
}

// HasControlTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasControlTxMsgCount() bool {
	if o != nil && o.ControlTxMsgCount != nil {
		return true
	}

	return false
}

// SetControlTxMsgCount gets a reference to the given int64 and assigns it to the ControlTxMsgCount field.
func (o *MsgVpn) SetControlTxMsgCount(v int64) {
	o.ControlTxMsgCount = &v
}

// GetCounter returns the Counter field value if set, zero value otherwise.
func (o *MsgVpn) GetCounter() MsgVpnCounter {
	if o == nil || o.Counter == nil {
		var ret MsgVpnCounter
		return ret
	}
	return *o.Counter
}

// GetCounterOk returns a tuple with the Counter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetCounterOk() (*MsgVpnCounter, bool) {
	if o == nil || o.Counter == nil {
		return nil, false
	}
	return o.Counter, true
}

// HasCounter returns a boolean if a field has been set.
func (o *MsgVpn) HasCounter() bool {
	if o != nil && o.Counter != nil {
		return true
	}

	return false
}

// SetCounter gets a reference to the given MsgVpnCounter and assigns it to the Counter field.
func (o *MsgVpn) SetCounter(v MsgVpnCounter) {
	o.Counter = &v
}

// GetDataRxByteCount returns the DataRxByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetDataRxByteCount() int64 {
	if o == nil || o.DataRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.DataRxByteCount
}

// GetDataRxByteCountOk returns a tuple with the DataRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetDataRxByteCountOk() (*int64, bool) {
	if o == nil || o.DataRxByteCount == nil {
		return nil, false
	}
	return o.DataRxByteCount, true
}

// HasDataRxByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasDataRxByteCount() bool {
	if o != nil && o.DataRxByteCount != nil {
		return true
	}

	return false
}

// SetDataRxByteCount gets a reference to the given int64 and assigns it to the DataRxByteCount field.
func (o *MsgVpn) SetDataRxByteCount(v int64) {
	o.DataRxByteCount = &v
}

// GetDataRxMsgCount returns the DataRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetDataRxMsgCount() int64 {
	if o == nil || o.DataRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DataRxMsgCount
}

// GetDataRxMsgCountOk returns a tuple with the DataRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetDataRxMsgCountOk() (*int64, bool) {
	if o == nil || o.DataRxMsgCount == nil {
		return nil, false
	}
	return o.DataRxMsgCount, true
}

// HasDataRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasDataRxMsgCount() bool {
	if o != nil && o.DataRxMsgCount != nil {
		return true
	}

	return false
}

// SetDataRxMsgCount gets a reference to the given int64 and assigns it to the DataRxMsgCount field.
func (o *MsgVpn) SetDataRxMsgCount(v int64) {
	o.DataRxMsgCount = &v
}

// GetDataTxByteCount returns the DataTxByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetDataTxByteCount() int64 {
	if o == nil || o.DataTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.DataTxByteCount
}

// GetDataTxByteCountOk returns a tuple with the DataTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetDataTxByteCountOk() (*int64, bool) {
	if o == nil || o.DataTxByteCount == nil {
		return nil, false
	}
	return o.DataTxByteCount, true
}

// HasDataTxByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasDataTxByteCount() bool {
	if o != nil && o.DataTxByteCount != nil {
		return true
	}

	return false
}

// SetDataTxByteCount gets a reference to the given int64 and assigns it to the DataTxByteCount field.
func (o *MsgVpn) SetDataTxByteCount(v int64) {
	o.DataTxByteCount = &v
}

// GetDataTxMsgCount returns the DataTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetDataTxMsgCount() int64 {
	if o == nil || o.DataTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DataTxMsgCount
}

// GetDataTxMsgCountOk returns a tuple with the DataTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetDataTxMsgCountOk() (*int64, bool) {
	if o == nil || o.DataTxMsgCount == nil {
		return nil, false
	}
	return o.DataTxMsgCount, true
}

// HasDataTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasDataTxMsgCount() bool {
	if o != nil && o.DataTxMsgCount != nil {
		return true
	}

	return false
}

// SetDataTxMsgCount gets a reference to the given int64 and assigns it to the DataTxMsgCount field.
func (o *MsgVpn) SetDataTxMsgCount(v int64) {
	o.DataTxMsgCount = &v
}

// GetDiscardedRxMsgCount returns the DiscardedRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetDiscardedRxMsgCount() int64 {
	if o == nil || o.DiscardedRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DiscardedRxMsgCount
}

// GetDiscardedRxMsgCountOk returns a tuple with the DiscardedRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetDiscardedRxMsgCountOk() (*int64, bool) {
	if o == nil || o.DiscardedRxMsgCount == nil {
		return nil, false
	}
	return o.DiscardedRxMsgCount, true
}

// HasDiscardedRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasDiscardedRxMsgCount() bool {
	if o != nil && o.DiscardedRxMsgCount != nil {
		return true
	}

	return false
}

// SetDiscardedRxMsgCount gets a reference to the given int64 and assigns it to the DiscardedRxMsgCount field.
func (o *MsgVpn) SetDiscardedRxMsgCount(v int64) {
	o.DiscardedRxMsgCount = &v
}

// GetDiscardedTxMsgCount returns the DiscardedTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetDiscardedTxMsgCount() int64 {
	if o == nil || o.DiscardedTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DiscardedTxMsgCount
}

// GetDiscardedTxMsgCountOk returns a tuple with the DiscardedTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetDiscardedTxMsgCountOk() (*int64, bool) {
	if o == nil || o.DiscardedTxMsgCount == nil {
		return nil, false
	}
	return o.DiscardedTxMsgCount, true
}

// HasDiscardedTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasDiscardedTxMsgCount() bool {
	if o != nil && o.DiscardedTxMsgCount != nil {
		return true
	}

	return false
}

// SetDiscardedTxMsgCount gets a reference to the given int64 and assigns it to the DiscardedTxMsgCount field.
func (o *MsgVpn) SetDiscardedTxMsgCount(v int64) {
	o.DiscardedTxMsgCount = &v
}

// GetDistributedCacheManagementEnabled returns the DistributedCacheManagementEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetDistributedCacheManagementEnabled() bool {
	if o == nil || o.DistributedCacheManagementEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DistributedCacheManagementEnabled
}

// GetDistributedCacheManagementEnabledOk returns a tuple with the DistributedCacheManagementEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetDistributedCacheManagementEnabledOk() (*bool, bool) {
	if o == nil || o.DistributedCacheManagementEnabled == nil {
		return nil, false
	}
	return o.DistributedCacheManagementEnabled, true
}

// HasDistributedCacheManagementEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasDistributedCacheManagementEnabled() bool {
	if o != nil && o.DistributedCacheManagementEnabled != nil {
		return true
	}

	return false
}

// SetDistributedCacheManagementEnabled gets a reference to the given bool and assigns it to the DistributedCacheManagementEnabled field.
func (o *MsgVpn) SetDistributedCacheManagementEnabled(v bool) {
	o.DistributedCacheManagementEnabled = &v
}

// GetDmrEnabled returns the DmrEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetDmrEnabled() bool {
	if o == nil || o.DmrEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DmrEnabled
}

// GetDmrEnabledOk returns a tuple with the DmrEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetDmrEnabledOk() (*bool, bool) {
	if o == nil || o.DmrEnabled == nil {
		return nil, false
	}
	return o.DmrEnabled, true
}

// HasDmrEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasDmrEnabled() bool {
	if o != nil && o.DmrEnabled != nil {
		return true
	}

	return false
}

// SetDmrEnabled gets a reference to the given bool and assigns it to the DmrEnabled field.
func (o *MsgVpn) SetDmrEnabled(v bool) {
	o.DmrEnabled = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MsgVpn) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MsgVpn) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventConnectionCountThreshold returns the EventConnectionCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventConnectionCountThreshold() EventThreshold {
	if o == nil || o.EventConnectionCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventConnectionCountThreshold
}

// GetEventConnectionCountThresholdOk returns a tuple with the EventConnectionCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventConnectionCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventConnectionCountThreshold == nil {
		return nil, false
	}
	return o.EventConnectionCountThreshold, true
}

// HasEventConnectionCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventConnectionCountThreshold() bool {
	if o != nil && o.EventConnectionCountThreshold != nil {
		return true
	}

	return false
}

// SetEventConnectionCountThreshold gets a reference to the given EventThreshold and assigns it to the EventConnectionCountThreshold field.
func (o *MsgVpn) SetEventConnectionCountThreshold(v EventThreshold) {
	o.EventConnectionCountThreshold = &v
}

// GetEventEgressFlowCountThreshold returns the EventEgressFlowCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventEgressFlowCountThreshold() EventThreshold {
	if o == nil || o.EventEgressFlowCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventEgressFlowCountThreshold
}

// GetEventEgressFlowCountThresholdOk returns a tuple with the EventEgressFlowCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventEgressFlowCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventEgressFlowCountThreshold == nil {
		return nil, false
	}
	return o.EventEgressFlowCountThreshold, true
}

// HasEventEgressFlowCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventEgressFlowCountThreshold() bool {
	if o != nil && o.EventEgressFlowCountThreshold != nil {
		return true
	}

	return false
}

// SetEventEgressFlowCountThreshold gets a reference to the given EventThreshold and assigns it to the EventEgressFlowCountThreshold field.
func (o *MsgVpn) SetEventEgressFlowCountThreshold(v EventThreshold) {
	o.EventEgressFlowCountThreshold = &v
}

// GetEventEgressMsgRateThreshold returns the EventEgressMsgRateThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventEgressMsgRateThreshold() EventThresholdByValue {
	if o == nil || o.EventEgressMsgRateThreshold == nil {
		var ret EventThresholdByValue
		return ret
	}
	return *o.EventEgressMsgRateThreshold
}

// GetEventEgressMsgRateThresholdOk returns a tuple with the EventEgressMsgRateThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventEgressMsgRateThresholdOk() (*EventThresholdByValue, bool) {
	if o == nil || o.EventEgressMsgRateThreshold == nil {
		return nil, false
	}
	return o.EventEgressMsgRateThreshold, true
}

// HasEventEgressMsgRateThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventEgressMsgRateThreshold() bool {
	if o != nil && o.EventEgressMsgRateThreshold != nil {
		return true
	}

	return false
}

// SetEventEgressMsgRateThreshold gets a reference to the given EventThresholdByValue and assigns it to the EventEgressMsgRateThreshold field.
func (o *MsgVpn) SetEventEgressMsgRateThreshold(v EventThresholdByValue) {
	o.EventEgressMsgRateThreshold = &v
}

// GetEventEndpointCountThreshold returns the EventEndpointCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventEndpointCountThreshold() EventThreshold {
	if o == nil || o.EventEndpointCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventEndpointCountThreshold
}

// GetEventEndpointCountThresholdOk returns a tuple with the EventEndpointCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventEndpointCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventEndpointCountThreshold == nil {
		return nil, false
	}
	return o.EventEndpointCountThreshold, true
}

// HasEventEndpointCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventEndpointCountThreshold() bool {
	if o != nil && o.EventEndpointCountThreshold != nil {
		return true
	}

	return false
}

// SetEventEndpointCountThreshold gets a reference to the given EventThreshold and assigns it to the EventEndpointCountThreshold field.
func (o *MsgVpn) SetEventEndpointCountThreshold(v EventThreshold) {
	o.EventEndpointCountThreshold = &v
}

// GetEventIngressFlowCountThreshold returns the EventIngressFlowCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventIngressFlowCountThreshold() EventThreshold {
	if o == nil || o.EventIngressFlowCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventIngressFlowCountThreshold
}

// GetEventIngressFlowCountThresholdOk returns a tuple with the EventIngressFlowCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventIngressFlowCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventIngressFlowCountThreshold == nil {
		return nil, false
	}
	return o.EventIngressFlowCountThreshold, true
}

// HasEventIngressFlowCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventIngressFlowCountThreshold() bool {
	if o != nil && o.EventIngressFlowCountThreshold != nil {
		return true
	}

	return false
}

// SetEventIngressFlowCountThreshold gets a reference to the given EventThreshold and assigns it to the EventIngressFlowCountThreshold field.
func (o *MsgVpn) SetEventIngressFlowCountThreshold(v EventThreshold) {
	o.EventIngressFlowCountThreshold = &v
}

// GetEventIngressMsgRateThreshold returns the EventIngressMsgRateThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventIngressMsgRateThreshold() EventThresholdByValue {
	if o == nil || o.EventIngressMsgRateThreshold == nil {
		var ret EventThresholdByValue
		return ret
	}
	return *o.EventIngressMsgRateThreshold
}

// GetEventIngressMsgRateThresholdOk returns a tuple with the EventIngressMsgRateThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventIngressMsgRateThresholdOk() (*EventThresholdByValue, bool) {
	if o == nil || o.EventIngressMsgRateThreshold == nil {
		return nil, false
	}
	return o.EventIngressMsgRateThreshold, true
}

// HasEventIngressMsgRateThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventIngressMsgRateThreshold() bool {
	if o != nil && o.EventIngressMsgRateThreshold != nil {
		return true
	}

	return false
}

// SetEventIngressMsgRateThreshold gets a reference to the given EventThresholdByValue and assigns it to the EventIngressMsgRateThreshold field.
func (o *MsgVpn) SetEventIngressMsgRateThreshold(v EventThresholdByValue) {
	o.EventIngressMsgRateThreshold = &v
}

// GetEventLargeMsgThreshold returns the EventLargeMsgThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventLargeMsgThreshold() int64 {
	if o == nil || o.EventLargeMsgThreshold == nil {
		var ret int64
		return ret
	}
	return *o.EventLargeMsgThreshold
}

// GetEventLargeMsgThresholdOk returns a tuple with the EventLargeMsgThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventLargeMsgThresholdOk() (*int64, bool) {
	if o == nil || o.EventLargeMsgThreshold == nil {
		return nil, false
	}
	return o.EventLargeMsgThreshold, true
}

// HasEventLargeMsgThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventLargeMsgThreshold() bool {
	if o != nil && o.EventLargeMsgThreshold != nil {
		return true
	}

	return false
}

// SetEventLargeMsgThreshold gets a reference to the given int64 and assigns it to the EventLargeMsgThreshold field.
func (o *MsgVpn) SetEventLargeMsgThreshold(v int64) {
	o.EventLargeMsgThreshold = &v
}

// GetEventLogTag returns the EventLogTag field value if set, zero value otherwise.
func (o *MsgVpn) GetEventLogTag() string {
	if o == nil || o.EventLogTag == nil {
		var ret string
		return ret
	}
	return *o.EventLogTag
}

// GetEventLogTagOk returns a tuple with the EventLogTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventLogTagOk() (*string, bool) {
	if o == nil || o.EventLogTag == nil {
		return nil, false
	}
	return o.EventLogTag, true
}

// HasEventLogTag returns a boolean if a field has been set.
func (o *MsgVpn) HasEventLogTag() bool {
	if o != nil && o.EventLogTag != nil {
		return true
	}

	return false
}

// SetEventLogTag gets a reference to the given string and assigns it to the EventLogTag field.
func (o *MsgVpn) SetEventLogTag(v string) {
	o.EventLogTag = &v
}

// GetEventMsgSpoolUsageThreshold returns the EventMsgSpoolUsageThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventMsgSpoolUsageThreshold() EventThreshold {
	if o == nil || o.EventMsgSpoolUsageThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventMsgSpoolUsageThreshold
}

// GetEventMsgSpoolUsageThresholdOk returns a tuple with the EventMsgSpoolUsageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventMsgSpoolUsageThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventMsgSpoolUsageThreshold == nil {
		return nil, false
	}
	return o.EventMsgSpoolUsageThreshold, true
}

// HasEventMsgSpoolUsageThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventMsgSpoolUsageThreshold() bool {
	if o != nil && o.EventMsgSpoolUsageThreshold != nil {
		return true
	}

	return false
}

// SetEventMsgSpoolUsageThreshold gets a reference to the given EventThreshold and assigns it to the EventMsgSpoolUsageThreshold field.
func (o *MsgVpn) SetEventMsgSpoolUsageThreshold(v EventThreshold) {
	o.EventMsgSpoolUsageThreshold = &v
}

// GetEventPublishClientEnabled returns the EventPublishClientEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetEventPublishClientEnabled() bool {
	if o == nil || o.EventPublishClientEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventPublishClientEnabled
}

// GetEventPublishClientEnabledOk returns a tuple with the EventPublishClientEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventPublishClientEnabledOk() (*bool, bool) {
	if o == nil || o.EventPublishClientEnabled == nil {
		return nil, false
	}
	return o.EventPublishClientEnabled, true
}

// HasEventPublishClientEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasEventPublishClientEnabled() bool {
	if o != nil && o.EventPublishClientEnabled != nil {
		return true
	}

	return false
}

// SetEventPublishClientEnabled gets a reference to the given bool and assigns it to the EventPublishClientEnabled field.
func (o *MsgVpn) SetEventPublishClientEnabled(v bool) {
	o.EventPublishClientEnabled = &v
}

// GetEventPublishMsgVpnEnabled returns the EventPublishMsgVpnEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetEventPublishMsgVpnEnabled() bool {
	if o == nil || o.EventPublishMsgVpnEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventPublishMsgVpnEnabled
}

// GetEventPublishMsgVpnEnabledOk returns a tuple with the EventPublishMsgVpnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventPublishMsgVpnEnabledOk() (*bool, bool) {
	if o == nil || o.EventPublishMsgVpnEnabled == nil {
		return nil, false
	}
	return o.EventPublishMsgVpnEnabled, true
}

// HasEventPublishMsgVpnEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasEventPublishMsgVpnEnabled() bool {
	if o != nil && o.EventPublishMsgVpnEnabled != nil {
		return true
	}

	return false
}

// SetEventPublishMsgVpnEnabled gets a reference to the given bool and assigns it to the EventPublishMsgVpnEnabled field.
func (o *MsgVpn) SetEventPublishMsgVpnEnabled(v bool) {
	o.EventPublishMsgVpnEnabled = &v
}

// GetEventPublishSubscriptionMode returns the EventPublishSubscriptionMode field value if set, zero value otherwise.
func (o *MsgVpn) GetEventPublishSubscriptionMode() string {
	if o == nil || o.EventPublishSubscriptionMode == nil {
		var ret string
		return ret
	}
	return *o.EventPublishSubscriptionMode
}

// GetEventPublishSubscriptionModeOk returns a tuple with the EventPublishSubscriptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventPublishSubscriptionModeOk() (*string, bool) {
	if o == nil || o.EventPublishSubscriptionMode == nil {
		return nil, false
	}
	return o.EventPublishSubscriptionMode, true
}

// HasEventPublishSubscriptionMode returns a boolean if a field has been set.
func (o *MsgVpn) HasEventPublishSubscriptionMode() bool {
	if o != nil && o.EventPublishSubscriptionMode != nil {
		return true
	}

	return false
}

// SetEventPublishSubscriptionMode gets a reference to the given string and assigns it to the EventPublishSubscriptionMode field.
func (o *MsgVpn) SetEventPublishSubscriptionMode(v string) {
	o.EventPublishSubscriptionMode = &v
}

// GetEventPublishTopicFormatMqttEnabled returns the EventPublishTopicFormatMqttEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetEventPublishTopicFormatMqttEnabled() bool {
	if o == nil || o.EventPublishTopicFormatMqttEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventPublishTopicFormatMqttEnabled
}

// GetEventPublishTopicFormatMqttEnabledOk returns a tuple with the EventPublishTopicFormatMqttEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventPublishTopicFormatMqttEnabledOk() (*bool, bool) {
	if o == nil || o.EventPublishTopicFormatMqttEnabled == nil {
		return nil, false
	}
	return o.EventPublishTopicFormatMqttEnabled, true
}

// HasEventPublishTopicFormatMqttEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasEventPublishTopicFormatMqttEnabled() bool {
	if o != nil && o.EventPublishTopicFormatMqttEnabled != nil {
		return true
	}

	return false
}

// SetEventPublishTopicFormatMqttEnabled gets a reference to the given bool and assigns it to the EventPublishTopicFormatMqttEnabled field.
func (o *MsgVpn) SetEventPublishTopicFormatMqttEnabled(v bool) {
	o.EventPublishTopicFormatMqttEnabled = &v
}

// GetEventPublishTopicFormatSmfEnabled returns the EventPublishTopicFormatSmfEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetEventPublishTopicFormatSmfEnabled() bool {
	if o == nil || o.EventPublishTopicFormatSmfEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EventPublishTopicFormatSmfEnabled
}

// GetEventPublishTopicFormatSmfEnabledOk returns a tuple with the EventPublishTopicFormatSmfEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventPublishTopicFormatSmfEnabledOk() (*bool, bool) {
	if o == nil || o.EventPublishTopicFormatSmfEnabled == nil {
		return nil, false
	}
	return o.EventPublishTopicFormatSmfEnabled, true
}

// HasEventPublishTopicFormatSmfEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasEventPublishTopicFormatSmfEnabled() bool {
	if o != nil && o.EventPublishTopicFormatSmfEnabled != nil {
		return true
	}

	return false
}

// SetEventPublishTopicFormatSmfEnabled gets a reference to the given bool and assigns it to the EventPublishTopicFormatSmfEnabled field.
func (o *MsgVpn) SetEventPublishTopicFormatSmfEnabled(v bool) {
	o.EventPublishTopicFormatSmfEnabled = &v
}

// GetEventServiceAmqpConnectionCountThreshold returns the EventServiceAmqpConnectionCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventServiceAmqpConnectionCountThreshold() EventThreshold {
	if o == nil || o.EventServiceAmqpConnectionCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventServiceAmqpConnectionCountThreshold
}

// GetEventServiceAmqpConnectionCountThresholdOk returns a tuple with the EventServiceAmqpConnectionCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventServiceAmqpConnectionCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventServiceAmqpConnectionCountThreshold == nil {
		return nil, false
	}
	return o.EventServiceAmqpConnectionCountThreshold, true
}

// HasEventServiceAmqpConnectionCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventServiceAmqpConnectionCountThreshold() bool {
	if o != nil && o.EventServiceAmqpConnectionCountThreshold != nil {
		return true
	}

	return false
}

// SetEventServiceAmqpConnectionCountThreshold gets a reference to the given EventThreshold and assigns it to the EventServiceAmqpConnectionCountThreshold field.
func (o *MsgVpn) SetEventServiceAmqpConnectionCountThreshold(v EventThreshold) {
	o.EventServiceAmqpConnectionCountThreshold = &v
}

// GetEventServiceMqttConnectionCountThreshold returns the EventServiceMqttConnectionCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventServiceMqttConnectionCountThreshold() EventThreshold {
	if o == nil || o.EventServiceMqttConnectionCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventServiceMqttConnectionCountThreshold
}

// GetEventServiceMqttConnectionCountThresholdOk returns a tuple with the EventServiceMqttConnectionCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventServiceMqttConnectionCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventServiceMqttConnectionCountThreshold == nil {
		return nil, false
	}
	return o.EventServiceMqttConnectionCountThreshold, true
}

// HasEventServiceMqttConnectionCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventServiceMqttConnectionCountThreshold() bool {
	if o != nil && o.EventServiceMqttConnectionCountThreshold != nil {
		return true
	}

	return false
}

// SetEventServiceMqttConnectionCountThreshold gets a reference to the given EventThreshold and assigns it to the EventServiceMqttConnectionCountThreshold field.
func (o *MsgVpn) SetEventServiceMqttConnectionCountThreshold(v EventThreshold) {
	o.EventServiceMqttConnectionCountThreshold = &v
}

// GetEventServiceRestIncomingConnectionCountThreshold returns the EventServiceRestIncomingConnectionCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventServiceRestIncomingConnectionCountThreshold() EventThreshold {
	if o == nil || o.EventServiceRestIncomingConnectionCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventServiceRestIncomingConnectionCountThreshold
}

// GetEventServiceRestIncomingConnectionCountThresholdOk returns a tuple with the EventServiceRestIncomingConnectionCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventServiceRestIncomingConnectionCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventServiceRestIncomingConnectionCountThreshold == nil {
		return nil, false
	}
	return o.EventServiceRestIncomingConnectionCountThreshold, true
}

// HasEventServiceRestIncomingConnectionCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventServiceRestIncomingConnectionCountThreshold() bool {
	if o != nil && o.EventServiceRestIncomingConnectionCountThreshold != nil {
		return true
	}

	return false
}

// SetEventServiceRestIncomingConnectionCountThreshold gets a reference to the given EventThreshold and assigns it to the EventServiceRestIncomingConnectionCountThreshold field.
func (o *MsgVpn) SetEventServiceRestIncomingConnectionCountThreshold(v EventThreshold) {
	o.EventServiceRestIncomingConnectionCountThreshold = &v
}

// GetEventServiceSmfConnectionCountThreshold returns the EventServiceSmfConnectionCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventServiceSmfConnectionCountThreshold() EventThreshold {
	if o == nil || o.EventServiceSmfConnectionCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventServiceSmfConnectionCountThreshold
}

// GetEventServiceSmfConnectionCountThresholdOk returns a tuple with the EventServiceSmfConnectionCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventServiceSmfConnectionCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventServiceSmfConnectionCountThreshold == nil {
		return nil, false
	}
	return o.EventServiceSmfConnectionCountThreshold, true
}

// HasEventServiceSmfConnectionCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventServiceSmfConnectionCountThreshold() bool {
	if o != nil && o.EventServiceSmfConnectionCountThreshold != nil {
		return true
	}

	return false
}

// SetEventServiceSmfConnectionCountThreshold gets a reference to the given EventThreshold and assigns it to the EventServiceSmfConnectionCountThreshold field.
func (o *MsgVpn) SetEventServiceSmfConnectionCountThreshold(v EventThreshold) {
	o.EventServiceSmfConnectionCountThreshold = &v
}

// GetEventServiceWebConnectionCountThreshold returns the EventServiceWebConnectionCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventServiceWebConnectionCountThreshold() EventThreshold {
	if o == nil || o.EventServiceWebConnectionCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventServiceWebConnectionCountThreshold
}

// GetEventServiceWebConnectionCountThresholdOk returns a tuple with the EventServiceWebConnectionCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventServiceWebConnectionCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventServiceWebConnectionCountThreshold == nil {
		return nil, false
	}
	return o.EventServiceWebConnectionCountThreshold, true
}

// HasEventServiceWebConnectionCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventServiceWebConnectionCountThreshold() bool {
	if o != nil && o.EventServiceWebConnectionCountThreshold != nil {
		return true
	}

	return false
}

// SetEventServiceWebConnectionCountThreshold gets a reference to the given EventThreshold and assigns it to the EventServiceWebConnectionCountThreshold field.
func (o *MsgVpn) SetEventServiceWebConnectionCountThreshold(v EventThreshold) {
	o.EventServiceWebConnectionCountThreshold = &v
}

// GetEventSubscriptionCountThreshold returns the EventSubscriptionCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventSubscriptionCountThreshold() EventThreshold {
	if o == nil || o.EventSubscriptionCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventSubscriptionCountThreshold
}

// GetEventSubscriptionCountThresholdOk returns a tuple with the EventSubscriptionCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventSubscriptionCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventSubscriptionCountThreshold == nil {
		return nil, false
	}
	return o.EventSubscriptionCountThreshold, true
}

// HasEventSubscriptionCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventSubscriptionCountThreshold() bool {
	if o != nil && o.EventSubscriptionCountThreshold != nil {
		return true
	}

	return false
}

// SetEventSubscriptionCountThreshold gets a reference to the given EventThreshold and assigns it to the EventSubscriptionCountThreshold field.
func (o *MsgVpn) SetEventSubscriptionCountThreshold(v EventThreshold) {
	o.EventSubscriptionCountThreshold = &v
}

// GetEventTransactedSessionCountThreshold returns the EventTransactedSessionCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventTransactedSessionCountThreshold() EventThreshold {
	if o == nil || o.EventTransactedSessionCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventTransactedSessionCountThreshold
}

// GetEventTransactedSessionCountThresholdOk returns a tuple with the EventTransactedSessionCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventTransactedSessionCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventTransactedSessionCountThreshold == nil {
		return nil, false
	}
	return o.EventTransactedSessionCountThreshold, true
}

// HasEventTransactedSessionCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventTransactedSessionCountThreshold() bool {
	if o != nil && o.EventTransactedSessionCountThreshold != nil {
		return true
	}

	return false
}

// SetEventTransactedSessionCountThreshold gets a reference to the given EventThreshold and assigns it to the EventTransactedSessionCountThreshold field.
func (o *MsgVpn) SetEventTransactedSessionCountThreshold(v EventThreshold) {
	o.EventTransactedSessionCountThreshold = &v
}

// GetEventTransactionCountThreshold returns the EventTransactionCountThreshold field value if set, zero value otherwise.
func (o *MsgVpn) GetEventTransactionCountThreshold() EventThreshold {
	if o == nil || o.EventTransactionCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventTransactionCountThreshold
}

// GetEventTransactionCountThresholdOk returns a tuple with the EventTransactionCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetEventTransactionCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventTransactionCountThreshold == nil {
		return nil, false
	}
	return o.EventTransactionCountThreshold, true
}

// HasEventTransactionCountThreshold returns a boolean if a field has been set.
func (o *MsgVpn) HasEventTransactionCountThreshold() bool {
	if o != nil && o.EventTransactionCountThreshold != nil {
		return true
	}

	return false
}

// SetEventTransactionCountThreshold gets a reference to the given EventThreshold and assigns it to the EventTransactionCountThreshold field.
func (o *MsgVpn) SetEventTransactionCountThreshold(v EventThreshold) {
	o.EventTransactionCountThreshold = &v
}

// GetExportSubscriptionsEnabled returns the ExportSubscriptionsEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetExportSubscriptionsEnabled() bool {
	if o == nil || o.ExportSubscriptionsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ExportSubscriptionsEnabled
}

// GetExportSubscriptionsEnabledOk returns a tuple with the ExportSubscriptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetExportSubscriptionsEnabledOk() (*bool, bool) {
	if o == nil || o.ExportSubscriptionsEnabled == nil {
		return nil, false
	}
	return o.ExportSubscriptionsEnabled, true
}

// HasExportSubscriptionsEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasExportSubscriptionsEnabled() bool {
	if o != nil && o.ExportSubscriptionsEnabled != nil {
		return true
	}

	return false
}

// SetExportSubscriptionsEnabled gets a reference to the given bool and assigns it to the ExportSubscriptionsEnabled field.
func (o *MsgVpn) SetExportSubscriptionsEnabled(v bool) {
	o.ExportSubscriptionsEnabled = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *MsgVpn) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetJndiEnabled returns the JndiEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetJndiEnabled() bool {
	if o == nil || o.JndiEnabled == nil {
		var ret bool
		return ret
	}
	return *o.JndiEnabled
}

// GetJndiEnabledOk returns a tuple with the JndiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetJndiEnabledOk() (*bool, bool) {
	if o == nil || o.JndiEnabled == nil {
		return nil, false
	}
	return o.JndiEnabled, true
}

// HasJndiEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasJndiEnabled() bool {
	if o != nil && o.JndiEnabled != nil {
		return true
	}

	return false
}

// SetJndiEnabled gets a reference to the given bool and assigns it to the JndiEnabled field.
func (o *MsgVpn) SetJndiEnabled(v bool) {
	o.JndiEnabled = &v
}

// GetLoginRxMsgCount returns the LoginRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetLoginRxMsgCount() int64 {
	if o == nil || o.LoginRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginRxMsgCount
}

// GetLoginRxMsgCountOk returns a tuple with the LoginRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetLoginRxMsgCountOk() (*int64, bool) {
	if o == nil || o.LoginRxMsgCount == nil {
		return nil, false
	}
	return o.LoginRxMsgCount, true
}

// HasLoginRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasLoginRxMsgCount() bool {
	if o != nil && o.LoginRxMsgCount != nil {
		return true
	}

	return false
}

// SetLoginRxMsgCount gets a reference to the given int64 and assigns it to the LoginRxMsgCount field.
func (o *MsgVpn) SetLoginRxMsgCount(v int64) {
	o.LoginRxMsgCount = &v
}

// GetLoginTxMsgCount returns the LoginTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetLoginTxMsgCount() int64 {
	if o == nil || o.LoginTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LoginTxMsgCount
}

// GetLoginTxMsgCountOk returns a tuple with the LoginTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetLoginTxMsgCountOk() (*int64, bool) {
	if o == nil || o.LoginTxMsgCount == nil {
		return nil, false
	}
	return o.LoginTxMsgCount, true
}

// HasLoginTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasLoginTxMsgCount() bool {
	if o != nil && o.LoginTxMsgCount != nil {
		return true
	}

	return false
}

// SetLoginTxMsgCount gets a reference to the given int64 and assigns it to the LoginTxMsgCount field.
func (o *MsgVpn) SetLoginTxMsgCount(v int64) {
	o.LoginTxMsgCount = &v
}

// GetMaxConnectionCount returns the MaxConnectionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxConnectionCount() int64 {
	if o == nil || o.MaxConnectionCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxConnectionCount
}

// GetMaxConnectionCountOk returns a tuple with the MaxConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxConnectionCountOk() (*int64, bool) {
	if o == nil || o.MaxConnectionCount == nil {
		return nil, false
	}
	return o.MaxConnectionCount, true
}

// HasMaxConnectionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxConnectionCount() bool {
	if o != nil && o.MaxConnectionCount != nil {
		return true
	}

	return false
}

// SetMaxConnectionCount gets a reference to the given int64 and assigns it to the MaxConnectionCount field.
func (o *MsgVpn) SetMaxConnectionCount(v int64) {
	o.MaxConnectionCount = &v
}

// GetMaxEffectiveEndpointCount returns the MaxEffectiveEndpointCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxEffectiveEndpointCount() int32 {
	if o == nil || o.MaxEffectiveEndpointCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxEffectiveEndpointCount
}

// GetMaxEffectiveEndpointCountOk returns a tuple with the MaxEffectiveEndpointCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxEffectiveEndpointCountOk() (*int32, bool) {
	if o == nil || o.MaxEffectiveEndpointCount == nil {
		return nil, false
	}
	return o.MaxEffectiveEndpointCount, true
}

// HasMaxEffectiveEndpointCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxEffectiveEndpointCount() bool {
	if o != nil && o.MaxEffectiveEndpointCount != nil {
		return true
	}

	return false
}

// SetMaxEffectiveEndpointCount gets a reference to the given int32 and assigns it to the MaxEffectiveEndpointCount field.
func (o *MsgVpn) SetMaxEffectiveEndpointCount(v int32) {
	o.MaxEffectiveEndpointCount = &v
}

// GetMaxEffectiveRxFlowCount returns the MaxEffectiveRxFlowCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxEffectiveRxFlowCount() int32 {
	if o == nil || o.MaxEffectiveRxFlowCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxEffectiveRxFlowCount
}

// GetMaxEffectiveRxFlowCountOk returns a tuple with the MaxEffectiveRxFlowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxEffectiveRxFlowCountOk() (*int32, bool) {
	if o == nil || o.MaxEffectiveRxFlowCount == nil {
		return nil, false
	}
	return o.MaxEffectiveRxFlowCount, true
}

// HasMaxEffectiveRxFlowCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxEffectiveRxFlowCount() bool {
	if o != nil && o.MaxEffectiveRxFlowCount != nil {
		return true
	}

	return false
}

// SetMaxEffectiveRxFlowCount gets a reference to the given int32 and assigns it to the MaxEffectiveRxFlowCount field.
func (o *MsgVpn) SetMaxEffectiveRxFlowCount(v int32) {
	o.MaxEffectiveRxFlowCount = &v
}

// GetMaxEffectiveSubscriptionCount returns the MaxEffectiveSubscriptionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxEffectiveSubscriptionCount() int64 {
	if o == nil || o.MaxEffectiveSubscriptionCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxEffectiveSubscriptionCount
}

// GetMaxEffectiveSubscriptionCountOk returns a tuple with the MaxEffectiveSubscriptionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxEffectiveSubscriptionCountOk() (*int64, bool) {
	if o == nil || o.MaxEffectiveSubscriptionCount == nil {
		return nil, false
	}
	return o.MaxEffectiveSubscriptionCount, true
}

// HasMaxEffectiveSubscriptionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxEffectiveSubscriptionCount() bool {
	if o != nil && o.MaxEffectiveSubscriptionCount != nil {
		return true
	}

	return false
}

// SetMaxEffectiveSubscriptionCount gets a reference to the given int64 and assigns it to the MaxEffectiveSubscriptionCount field.
func (o *MsgVpn) SetMaxEffectiveSubscriptionCount(v int64) {
	o.MaxEffectiveSubscriptionCount = &v
}

// GetMaxEffectiveTransactedSessionCount returns the MaxEffectiveTransactedSessionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxEffectiveTransactedSessionCount() int32 {
	if o == nil || o.MaxEffectiveTransactedSessionCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxEffectiveTransactedSessionCount
}

// GetMaxEffectiveTransactedSessionCountOk returns a tuple with the MaxEffectiveTransactedSessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxEffectiveTransactedSessionCountOk() (*int32, bool) {
	if o == nil || o.MaxEffectiveTransactedSessionCount == nil {
		return nil, false
	}
	return o.MaxEffectiveTransactedSessionCount, true
}

// HasMaxEffectiveTransactedSessionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxEffectiveTransactedSessionCount() bool {
	if o != nil && o.MaxEffectiveTransactedSessionCount != nil {
		return true
	}

	return false
}

// SetMaxEffectiveTransactedSessionCount gets a reference to the given int32 and assigns it to the MaxEffectiveTransactedSessionCount field.
func (o *MsgVpn) SetMaxEffectiveTransactedSessionCount(v int32) {
	o.MaxEffectiveTransactedSessionCount = &v
}

// GetMaxEffectiveTransactionCount returns the MaxEffectiveTransactionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxEffectiveTransactionCount() int32 {
	if o == nil || o.MaxEffectiveTransactionCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxEffectiveTransactionCount
}

// GetMaxEffectiveTransactionCountOk returns a tuple with the MaxEffectiveTransactionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxEffectiveTransactionCountOk() (*int32, bool) {
	if o == nil || o.MaxEffectiveTransactionCount == nil {
		return nil, false
	}
	return o.MaxEffectiveTransactionCount, true
}

// HasMaxEffectiveTransactionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxEffectiveTransactionCount() bool {
	if o != nil && o.MaxEffectiveTransactionCount != nil {
		return true
	}

	return false
}

// SetMaxEffectiveTransactionCount gets a reference to the given int32 and assigns it to the MaxEffectiveTransactionCount field.
func (o *MsgVpn) SetMaxEffectiveTransactionCount(v int32) {
	o.MaxEffectiveTransactionCount = &v
}

// GetMaxEffectiveTxFlowCount returns the MaxEffectiveTxFlowCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxEffectiveTxFlowCount() int32 {
	if o == nil || o.MaxEffectiveTxFlowCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxEffectiveTxFlowCount
}

// GetMaxEffectiveTxFlowCountOk returns a tuple with the MaxEffectiveTxFlowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxEffectiveTxFlowCountOk() (*int32, bool) {
	if o == nil || o.MaxEffectiveTxFlowCount == nil {
		return nil, false
	}
	return o.MaxEffectiveTxFlowCount, true
}

// HasMaxEffectiveTxFlowCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxEffectiveTxFlowCount() bool {
	if o != nil && o.MaxEffectiveTxFlowCount != nil {
		return true
	}

	return false
}

// SetMaxEffectiveTxFlowCount gets a reference to the given int32 and assigns it to the MaxEffectiveTxFlowCount field.
func (o *MsgVpn) SetMaxEffectiveTxFlowCount(v int32) {
	o.MaxEffectiveTxFlowCount = &v
}

// GetMaxEgressFlowCount returns the MaxEgressFlowCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxEgressFlowCount() int64 {
	if o == nil || o.MaxEgressFlowCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxEgressFlowCount
}

// GetMaxEgressFlowCountOk returns a tuple with the MaxEgressFlowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxEgressFlowCountOk() (*int64, bool) {
	if o == nil || o.MaxEgressFlowCount == nil {
		return nil, false
	}
	return o.MaxEgressFlowCount, true
}

// HasMaxEgressFlowCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxEgressFlowCount() bool {
	if o != nil && o.MaxEgressFlowCount != nil {
		return true
	}

	return false
}

// SetMaxEgressFlowCount gets a reference to the given int64 and assigns it to the MaxEgressFlowCount field.
func (o *MsgVpn) SetMaxEgressFlowCount(v int64) {
	o.MaxEgressFlowCount = &v
}

// GetMaxEndpointCount returns the MaxEndpointCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxEndpointCount() int64 {
	if o == nil || o.MaxEndpointCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxEndpointCount
}

// GetMaxEndpointCountOk returns a tuple with the MaxEndpointCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxEndpointCountOk() (*int64, bool) {
	if o == nil || o.MaxEndpointCount == nil {
		return nil, false
	}
	return o.MaxEndpointCount, true
}

// HasMaxEndpointCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxEndpointCount() bool {
	if o != nil && o.MaxEndpointCount != nil {
		return true
	}

	return false
}

// SetMaxEndpointCount gets a reference to the given int64 and assigns it to the MaxEndpointCount field.
func (o *MsgVpn) SetMaxEndpointCount(v int64) {
	o.MaxEndpointCount = &v
}

// GetMaxIngressFlowCount returns the MaxIngressFlowCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxIngressFlowCount() int64 {
	if o == nil || o.MaxIngressFlowCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxIngressFlowCount
}

// GetMaxIngressFlowCountOk returns a tuple with the MaxIngressFlowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxIngressFlowCountOk() (*int64, bool) {
	if o == nil || o.MaxIngressFlowCount == nil {
		return nil, false
	}
	return o.MaxIngressFlowCount, true
}

// HasMaxIngressFlowCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxIngressFlowCount() bool {
	if o != nil && o.MaxIngressFlowCount != nil {
		return true
	}

	return false
}

// SetMaxIngressFlowCount gets a reference to the given int64 and assigns it to the MaxIngressFlowCount field.
func (o *MsgVpn) SetMaxIngressFlowCount(v int64) {
	o.MaxIngressFlowCount = &v
}

// GetMaxMsgSpoolUsage returns the MaxMsgSpoolUsage field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxMsgSpoolUsage() int64 {
	if o == nil || o.MaxMsgSpoolUsage == nil {
		var ret int64
		return ret
	}
	return *o.MaxMsgSpoolUsage
}

// GetMaxMsgSpoolUsageOk returns a tuple with the MaxMsgSpoolUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxMsgSpoolUsageOk() (*int64, bool) {
	if o == nil || o.MaxMsgSpoolUsage == nil {
		return nil, false
	}
	return o.MaxMsgSpoolUsage, true
}

// HasMaxMsgSpoolUsage returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxMsgSpoolUsage() bool {
	if o != nil && o.MaxMsgSpoolUsage != nil {
		return true
	}

	return false
}

// SetMaxMsgSpoolUsage gets a reference to the given int64 and assigns it to the MaxMsgSpoolUsage field.
func (o *MsgVpn) SetMaxMsgSpoolUsage(v int64) {
	o.MaxMsgSpoolUsage = &v
}

// GetMaxSubscriptionCount returns the MaxSubscriptionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxSubscriptionCount() int64 {
	if o == nil || o.MaxSubscriptionCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxSubscriptionCount
}

// GetMaxSubscriptionCountOk returns a tuple with the MaxSubscriptionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxSubscriptionCountOk() (*int64, bool) {
	if o == nil || o.MaxSubscriptionCount == nil {
		return nil, false
	}
	return o.MaxSubscriptionCount, true
}

// HasMaxSubscriptionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxSubscriptionCount() bool {
	if o != nil && o.MaxSubscriptionCount != nil {
		return true
	}

	return false
}

// SetMaxSubscriptionCount gets a reference to the given int64 and assigns it to the MaxSubscriptionCount field.
func (o *MsgVpn) SetMaxSubscriptionCount(v int64) {
	o.MaxSubscriptionCount = &v
}

// GetMaxTransactedSessionCount returns the MaxTransactedSessionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxTransactedSessionCount() int64 {
	if o == nil || o.MaxTransactedSessionCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxTransactedSessionCount
}

// GetMaxTransactedSessionCountOk returns a tuple with the MaxTransactedSessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxTransactedSessionCountOk() (*int64, bool) {
	if o == nil || o.MaxTransactedSessionCount == nil {
		return nil, false
	}
	return o.MaxTransactedSessionCount, true
}

// HasMaxTransactedSessionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxTransactedSessionCount() bool {
	if o != nil && o.MaxTransactedSessionCount != nil {
		return true
	}

	return false
}

// SetMaxTransactedSessionCount gets a reference to the given int64 and assigns it to the MaxTransactedSessionCount field.
func (o *MsgVpn) SetMaxTransactedSessionCount(v int64) {
	o.MaxTransactedSessionCount = &v
}

// GetMaxTransactionCount returns the MaxTransactionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMaxTransactionCount() int64 {
	if o == nil || o.MaxTransactionCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxTransactionCount
}

// GetMaxTransactionCountOk returns a tuple with the MaxTransactionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMaxTransactionCountOk() (*int64, bool) {
	if o == nil || o.MaxTransactionCount == nil {
		return nil, false
	}
	return o.MaxTransactionCount, true
}

// HasMaxTransactionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMaxTransactionCount() bool {
	if o != nil && o.MaxTransactionCount != nil {
		return true
	}

	return false
}

// SetMaxTransactionCount gets a reference to the given int64 and assigns it to the MaxTransactionCount field.
func (o *MsgVpn) SetMaxTransactionCount(v int64) {
	o.MaxTransactionCount = &v
}

// GetMqttRetainMaxMemory returns the MqttRetainMaxMemory field value if set, zero value otherwise.
func (o *MsgVpn) GetMqttRetainMaxMemory() int32 {
	if o == nil || o.MqttRetainMaxMemory == nil {
		var ret int32
		return ret
	}
	return *o.MqttRetainMaxMemory
}

// GetMqttRetainMaxMemoryOk returns a tuple with the MqttRetainMaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMqttRetainMaxMemoryOk() (*int32, bool) {
	if o == nil || o.MqttRetainMaxMemory == nil {
		return nil, false
	}
	return o.MqttRetainMaxMemory, true
}

// HasMqttRetainMaxMemory returns a boolean if a field has been set.
func (o *MsgVpn) HasMqttRetainMaxMemory() bool {
	if o != nil && o.MqttRetainMaxMemory != nil {
		return true
	}

	return false
}

// SetMqttRetainMaxMemory gets a reference to the given int32 and assigns it to the MqttRetainMaxMemory field.
func (o *MsgVpn) SetMqttRetainMaxMemory(v int32) {
	o.MqttRetainMaxMemory = &v
}

// GetMsgReplayActiveCount returns the MsgReplayActiveCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMsgReplayActiveCount() int32 {
	if o == nil || o.MsgReplayActiveCount == nil {
		var ret int32
		return ret
	}
	return *o.MsgReplayActiveCount
}

// GetMsgReplayActiveCountOk returns a tuple with the MsgReplayActiveCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMsgReplayActiveCountOk() (*int32, bool) {
	if o == nil || o.MsgReplayActiveCount == nil {
		return nil, false
	}
	return o.MsgReplayActiveCount, true
}

// HasMsgReplayActiveCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMsgReplayActiveCount() bool {
	if o != nil && o.MsgReplayActiveCount != nil {
		return true
	}

	return false
}

// SetMsgReplayActiveCount gets a reference to the given int32 and assigns it to the MsgReplayActiveCount field.
func (o *MsgVpn) SetMsgReplayActiveCount(v int32) {
	o.MsgReplayActiveCount = &v
}

// GetMsgReplayFailedCount returns the MsgReplayFailedCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMsgReplayFailedCount() int32 {
	if o == nil || o.MsgReplayFailedCount == nil {
		var ret int32
		return ret
	}
	return *o.MsgReplayFailedCount
}

// GetMsgReplayFailedCountOk returns a tuple with the MsgReplayFailedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMsgReplayFailedCountOk() (*int32, bool) {
	if o == nil || o.MsgReplayFailedCount == nil {
		return nil, false
	}
	return o.MsgReplayFailedCount, true
}

// HasMsgReplayFailedCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMsgReplayFailedCount() bool {
	if o != nil && o.MsgReplayFailedCount != nil {
		return true
	}

	return false
}

// SetMsgReplayFailedCount gets a reference to the given int32 and assigns it to the MsgReplayFailedCount field.
func (o *MsgVpn) SetMsgReplayFailedCount(v int32) {
	o.MsgReplayFailedCount = &v
}

// GetMsgReplayInitializingCount returns the MsgReplayInitializingCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMsgReplayInitializingCount() int32 {
	if o == nil || o.MsgReplayInitializingCount == nil {
		var ret int32
		return ret
	}
	return *o.MsgReplayInitializingCount
}

// GetMsgReplayInitializingCountOk returns a tuple with the MsgReplayInitializingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMsgReplayInitializingCountOk() (*int32, bool) {
	if o == nil || o.MsgReplayInitializingCount == nil {
		return nil, false
	}
	return o.MsgReplayInitializingCount, true
}

// HasMsgReplayInitializingCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMsgReplayInitializingCount() bool {
	if o != nil && o.MsgReplayInitializingCount != nil {
		return true
	}

	return false
}

// SetMsgReplayInitializingCount gets a reference to the given int32 and assigns it to the MsgReplayInitializingCount field.
func (o *MsgVpn) SetMsgReplayInitializingCount(v int32) {
	o.MsgReplayInitializingCount = &v
}

// GetMsgReplayPendingCompleteCount returns the MsgReplayPendingCompleteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMsgReplayPendingCompleteCount() int32 {
	if o == nil || o.MsgReplayPendingCompleteCount == nil {
		var ret int32
		return ret
	}
	return *o.MsgReplayPendingCompleteCount
}

// GetMsgReplayPendingCompleteCountOk returns a tuple with the MsgReplayPendingCompleteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMsgReplayPendingCompleteCountOk() (*int32, bool) {
	if o == nil || o.MsgReplayPendingCompleteCount == nil {
		return nil, false
	}
	return o.MsgReplayPendingCompleteCount, true
}

// HasMsgReplayPendingCompleteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMsgReplayPendingCompleteCount() bool {
	if o != nil && o.MsgReplayPendingCompleteCount != nil {
		return true
	}

	return false
}

// SetMsgReplayPendingCompleteCount gets a reference to the given int32 and assigns it to the MsgReplayPendingCompleteCount field.
func (o *MsgVpn) SetMsgReplayPendingCompleteCount(v int32) {
	o.MsgReplayPendingCompleteCount = &v
}

// GetMsgSpoolMsgCount returns the MsgSpoolMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMsgSpoolMsgCount() int64 {
	if o == nil || o.MsgSpoolMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MsgSpoolMsgCount
}

// GetMsgSpoolMsgCountOk returns a tuple with the MsgSpoolMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMsgSpoolMsgCountOk() (*int64, bool) {
	if o == nil || o.MsgSpoolMsgCount == nil {
		return nil, false
	}
	return o.MsgSpoolMsgCount, true
}

// HasMsgSpoolMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMsgSpoolMsgCount() bool {
	if o != nil && o.MsgSpoolMsgCount != nil {
		return true
	}

	return false
}

// SetMsgSpoolMsgCount gets a reference to the given int64 and assigns it to the MsgSpoolMsgCount field.
func (o *MsgVpn) SetMsgSpoolMsgCount(v int64) {
	o.MsgSpoolMsgCount = &v
}

// GetMsgSpoolRxMsgCount returns the MsgSpoolRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMsgSpoolRxMsgCount() int64 {
	if o == nil || o.MsgSpoolRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MsgSpoolRxMsgCount
}

// GetMsgSpoolRxMsgCountOk returns a tuple with the MsgSpoolRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMsgSpoolRxMsgCountOk() (*int64, bool) {
	if o == nil || o.MsgSpoolRxMsgCount == nil {
		return nil, false
	}
	return o.MsgSpoolRxMsgCount, true
}

// HasMsgSpoolRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMsgSpoolRxMsgCount() bool {
	if o != nil && o.MsgSpoolRxMsgCount != nil {
		return true
	}

	return false
}

// SetMsgSpoolRxMsgCount gets a reference to the given int64 and assigns it to the MsgSpoolRxMsgCount field.
func (o *MsgVpn) SetMsgSpoolRxMsgCount(v int64) {
	o.MsgSpoolRxMsgCount = &v
}

// GetMsgSpoolTxMsgCount returns the MsgSpoolTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetMsgSpoolTxMsgCount() int64 {
	if o == nil || o.MsgSpoolTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MsgSpoolTxMsgCount
}

// GetMsgSpoolTxMsgCountOk returns a tuple with the MsgSpoolTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMsgSpoolTxMsgCountOk() (*int64, bool) {
	if o == nil || o.MsgSpoolTxMsgCount == nil {
		return nil, false
	}
	return o.MsgSpoolTxMsgCount, true
}

// HasMsgSpoolTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasMsgSpoolTxMsgCount() bool {
	if o != nil && o.MsgSpoolTxMsgCount != nil {
		return true
	}

	return false
}

// SetMsgSpoolTxMsgCount gets a reference to the given int64 and assigns it to the MsgSpoolTxMsgCount field.
func (o *MsgVpn) SetMsgSpoolTxMsgCount(v int64) {
	o.MsgSpoolTxMsgCount = &v
}

// GetMsgSpoolUsage returns the MsgSpoolUsage field value if set, zero value otherwise.
func (o *MsgVpn) GetMsgSpoolUsage() int64 {
	if o == nil || o.MsgSpoolUsage == nil {
		var ret int64
		return ret
	}
	return *o.MsgSpoolUsage
}

// GetMsgSpoolUsageOk returns a tuple with the MsgSpoolUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMsgSpoolUsageOk() (*int64, bool) {
	if o == nil || o.MsgSpoolUsage == nil {
		return nil, false
	}
	return o.MsgSpoolUsage, true
}

// HasMsgSpoolUsage returns a boolean if a field has been set.
func (o *MsgVpn) HasMsgSpoolUsage() bool {
	if o != nil && o.MsgSpoolUsage != nil {
		return true
	}

	return false
}

// SetMsgSpoolUsage gets a reference to the given int64 and assigns it to the MsgSpoolUsage field.
func (o *MsgVpn) SetMsgSpoolUsage(v int64) {
	o.MsgSpoolUsage = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpn) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpn) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpn) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *MsgVpn) GetRate() MsgVpnRate {
	if o == nil || o.Rate == nil {
		var ret MsgVpnRate
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRateOk() (*MsgVpnRate, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *MsgVpn) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given MsgVpnRate and assigns it to the Rate field.
func (o *MsgVpn) SetRate(v MsgVpnRate) {
	o.Rate = &v
}

// GetReplicationAckPropagationIntervalMsgCount returns the ReplicationAckPropagationIntervalMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationAckPropagationIntervalMsgCount() int64 {
	if o == nil || o.ReplicationAckPropagationIntervalMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationAckPropagationIntervalMsgCount
}

// GetReplicationAckPropagationIntervalMsgCountOk returns a tuple with the ReplicationAckPropagationIntervalMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationAckPropagationIntervalMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationAckPropagationIntervalMsgCount == nil {
		return nil, false
	}
	return o.ReplicationAckPropagationIntervalMsgCount, true
}

// HasReplicationAckPropagationIntervalMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationAckPropagationIntervalMsgCount() bool {
	if o != nil && o.ReplicationAckPropagationIntervalMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationAckPropagationIntervalMsgCount gets a reference to the given int64 and assigns it to the ReplicationAckPropagationIntervalMsgCount field.
func (o *MsgVpn) SetReplicationAckPropagationIntervalMsgCount(v int64) {
	o.ReplicationAckPropagationIntervalMsgCount = &v
}

// GetReplicationActiveAckPropTxMsgCount returns the ReplicationActiveAckPropTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveAckPropTxMsgCount() int64 {
	if o == nil || o.ReplicationActiveAckPropTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationActiveAckPropTxMsgCount
}

// GetReplicationActiveAckPropTxMsgCountOk returns a tuple with the ReplicationActiveAckPropTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveAckPropTxMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationActiveAckPropTxMsgCount == nil {
		return nil, false
	}
	return o.ReplicationActiveAckPropTxMsgCount, true
}

// HasReplicationActiveAckPropTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveAckPropTxMsgCount() bool {
	if o != nil && o.ReplicationActiveAckPropTxMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationActiveAckPropTxMsgCount gets a reference to the given int64 and assigns it to the ReplicationActiveAckPropTxMsgCount field.
func (o *MsgVpn) SetReplicationActiveAckPropTxMsgCount(v int64) {
	o.ReplicationActiveAckPropTxMsgCount = &v
}

// GetReplicationActiveAsyncQueuedMsgCount returns the ReplicationActiveAsyncQueuedMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveAsyncQueuedMsgCount() int64 {
	if o == nil || o.ReplicationActiveAsyncQueuedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationActiveAsyncQueuedMsgCount
}

// GetReplicationActiveAsyncQueuedMsgCountOk returns a tuple with the ReplicationActiveAsyncQueuedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveAsyncQueuedMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationActiveAsyncQueuedMsgCount == nil {
		return nil, false
	}
	return o.ReplicationActiveAsyncQueuedMsgCount, true
}

// HasReplicationActiveAsyncQueuedMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveAsyncQueuedMsgCount() bool {
	if o != nil && o.ReplicationActiveAsyncQueuedMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationActiveAsyncQueuedMsgCount gets a reference to the given int64 and assigns it to the ReplicationActiveAsyncQueuedMsgCount field.
func (o *MsgVpn) SetReplicationActiveAsyncQueuedMsgCount(v int64) {
	o.ReplicationActiveAsyncQueuedMsgCount = &v
}

// GetReplicationActiveLocallyConsumedMsgCount returns the ReplicationActiveLocallyConsumedMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveLocallyConsumedMsgCount() int64 {
	if o == nil || o.ReplicationActiveLocallyConsumedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationActiveLocallyConsumedMsgCount
}

// GetReplicationActiveLocallyConsumedMsgCountOk returns a tuple with the ReplicationActiveLocallyConsumedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveLocallyConsumedMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationActiveLocallyConsumedMsgCount == nil {
		return nil, false
	}
	return o.ReplicationActiveLocallyConsumedMsgCount, true
}

// HasReplicationActiveLocallyConsumedMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveLocallyConsumedMsgCount() bool {
	if o != nil && o.ReplicationActiveLocallyConsumedMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationActiveLocallyConsumedMsgCount gets a reference to the given int64 and assigns it to the ReplicationActiveLocallyConsumedMsgCount field.
func (o *MsgVpn) SetReplicationActiveLocallyConsumedMsgCount(v int64) {
	o.ReplicationActiveLocallyConsumedMsgCount = &v
}

// GetReplicationActiveMateFlowCongestedPeakTime returns the ReplicationActiveMateFlowCongestedPeakTime field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveMateFlowCongestedPeakTime() int32 {
	if o == nil || o.ReplicationActiveMateFlowCongestedPeakTime == nil {
		var ret int32
		return ret
	}
	return *o.ReplicationActiveMateFlowCongestedPeakTime
}

// GetReplicationActiveMateFlowCongestedPeakTimeOk returns a tuple with the ReplicationActiveMateFlowCongestedPeakTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveMateFlowCongestedPeakTimeOk() (*int32, bool) {
	if o == nil || o.ReplicationActiveMateFlowCongestedPeakTime == nil {
		return nil, false
	}
	return o.ReplicationActiveMateFlowCongestedPeakTime, true
}

// HasReplicationActiveMateFlowCongestedPeakTime returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveMateFlowCongestedPeakTime() bool {
	if o != nil && o.ReplicationActiveMateFlowCongestedPeakTime != nil {
		return true
	}

	return false
}

// SetReplicationActiveMateFlowCongestedPeakTime gets a reference to the given int32 and assigns it to the ReplicationActiveMateFlowCongestedPeakTime field.
func (o *MsgVpn) SetReplicationActiveMateFlowCongestedPeakTime(v int32) {
	o.ReplicationActiveMateFlowCongestedPeakTime = &v
}

// GetReplicationActiveMateFlowNotCongestedPeakTime returns the ReplicationActiveMateFlowNotCongestedPeakTime field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveMateFlowNotCongestedPeakTime() int32 {
	if o == nil || o.ReplicationActiveMateFlowNotCongestedPeakTime == nil {
		var ret int32
		return ret
	}
	return *o.ReplicationActiveMateFlowNotCongestedPeakTime
}

// GetReplicationActiveMateFlowNotCongestedPeakTimeOk returns a tuple with the ReplicationActiveMateFlowNotCongestedPeakTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveMateFlowNotCongestedPeakTimeOk() (*int32, bool) {
	if o == nil || o.ReplicationActiveMateFlowNotCongestedPeakTime == nil {
		return nil, false
	}
	return o.ReplicationActiveMateFlowNotCongestedPeakTime, true
}

// HasReplicationActiveMateFlowNotCongestedPeakTime returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveMateFlowNotCongestedPeakTime() bool {
	if o != nil && o.ReplicationActiveMateFlowNotCongestedPeakTime != nil {
		return true
	}

	return false
}

// SetReplicationActiveMateFlowNotCongestedPeakTime gets a reference to the given int32 and assigns it to the ReplicationActiveMateFlowNotCongestedPeakTime field.
func (o *MsgVpn) SetReplicationActiveMateFlowNotCongestedPeakTime(v int32) {
	o.ReplicationActiveMateFlowNotCongestedPeakTime = &v
}

// GetReplicationActivePromotedQueuedMsgCount returns the ReplicationActivePromotedQueuedMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActivePromotedQueuedMsgCount() int64 {
	if o == nil || o.ReplicationActivePromotedQueuedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationActivePromotedQueuedMsgCount
}

// GetReplicationActivePromotedQueuedMsgCountOk returns a tuple with the ReplicationActivePromotedQueuedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActivePromotedQueuedMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationActivePromotedQueuedMsgCount == nil {
		return nil, false
	}
	return o.ReplicationActivePromotedQueuedMsgCount, true
}

// HasReplicationActivePromotedQueuedMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActivePromotedQueuedMsgCount() bool {
	if o != nil && o.ReplicationActivePromotedQueuedMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationActivePromotedQueuedMsgCount gets a reference to the given int64 and assigns it to the ReplicationActivePromotedQueuedMsgCount field.
func (o *MsgVpn) SetReplicationActivePromotedQueuedMsgCount(v int64) {
	o.ReplicationActivePromotedQueuedMsgCount = &v
}

// GetReplicationActiveReconcileRequestRxMsgCount returns the ReplicationActiveReconcileRequestRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveReconcileRequestRxMsgCount() int64 {
	if o == nil || o.ReplicationActiveReconcileRequestRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationActiveReconcileRequestRxMsgCount
}

// GetReplicationActiveReconcileRequestRxMsgCountOk returns a tuple with the ReplicationActiveReconcileRequestRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveReconcileRequestRxMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationActiveReconcileRequestRxMsgCount == nil {
		return nil, false
	}
	return o.ReplicationActiveReconcileRequestRxMsgCount, true
}

// HasReplicationActiveReconcileRequestRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveReconcileRequestRxMsgCount() bool {
	if o != nil && o.ReplicationActiveReconcileRequestRxMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationActiveReconcileRequestRxMsgCount gets a reference to the given int64 and assigns it to the ReplicationActiveReconcileRequestRxMsgCount field.
func (o *MsgVpn) SetReplicationActiveReconcileRequestRxMsgCount(v int64) {
	o.ReplicationActiveReconcileRequestRxMsgCount = &v
}

// GetReplicationActiveSyncEligiblePeakTime returns the ReplicationActiveSyncEligiblePeakTime field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveSyncEligiblePeakTime() int32 {
	if o == nil || o.ReplicationActiveSyncEligiblePeakTime == nil {
		var ret int32
		return ret
	}
	return *o.ReplicationActiveSyncEligiblePeakTime
}

// GetReplicationActiveSyncEligiblePeakTimeOk returns a tuple with the ReplicationActiveSyncEligiblePeakTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveSyncEligiblePeakTimeOk() (*int32, bool) {
	if o == nil || o.ReplicationActiveSyncEligiblePeakTime == nil {
		return nil, false
	}
	return o.ReplicationActiveSyncEligiblePeakTime, true
}

// HasReplicationActiveSyncEligiblePeakTime returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveSyncEligiblePeakTime() bool {
	if o != nil && o.ReplicationActiveSyncEligiblePeakTime != nil {
		return true
	}

	return false
}

// SetReplicationActiveSyncEligiblePeakTime gets a reference to the given int32 and assigns it to the ReplicationActiveSyncEligiblePeakTime field.
func (o *MsgVpn) SetReplicationActiveSyncEligiblePeakTime(v int32) {
	o.ReplicationActiveSyncEligiblePeakTime = &v
}

// GetReplicationActiveSyncIneligiblePeakTime returns the ReplicationActiveSyncIneligiblePeakTime field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveSyncIneligiblePeakTime() int32 {
	if o == nil || o.ReplicationActiveSyncIneligiblePeakTime == nil {
		var ret int32
		return ret
	}
	return *o.ReplicationActiveSyncIneligiblePeakTime
}

// GetReplicationActiveSyncIneligiblePeakTimeOk returns a tuple with the ReplicationActiveSyncIneligiblePeakTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveSyncIneligiblePeakTimeOk() (*int32, bool) {
	if o == nil || o.ReplicationActiveSyncIneligiblePeakTime == nil {
		return nil, false
	}
	return o.ReplicationActiveSyncIneligiblePeakTime, true
}

// HasReplicationActiveSyncIneligiblePeakTime returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveSyncIneligiblePeakTime() bool {
	if o != nil && o.ReplicationActiveSyncIneligiblePeakTime != nil {
		return true
	}

	return false
}

// SetReplicationActiveSyncIneligiblePeakTime gets a reference to the given int32 and assigns it to the ReplicationActiveSyncIneligiblePeakTime field.
func (o *MsgVpn) SetReplicationActiveSyncIneligiblePeakTime(v int32) {
	o.ReplicationActiveSyncIneligiblePeakTime = &v
}

// GetReplicationActiveSyncQueuedAsAsyncMsgCount returns the ReplicationActiveSyncQueuedAsAsyncMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveSyncQueuedAsAsyncMsgCount() int64 {
	if o == nil || o.ReplicationActiveSyncQueuedAsAsyncMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationActiveSyncQueuedAsAsyncMsgCount
}

// GetReplicationActiveSyncQueuedAsAsyncMsgCountOk returns a tuple with the ReplicationActiveSyncQueuedAsAsyncMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveSyncQueuedAsAsyncMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationActiveSyncQueuedAsAsyncMsgCount == nil {
		return nil, false
	}
	return o.ReplicationActiveSyncQueuedAsAsyncMsgCount, true
}

// HasReplicationActiveSyncQueuedAsAsyncMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveSyncQueuedAsAsyncMsgCount() bool {
	if o != nil && o.ReplicationActiveSyncQueuedAsAsyncMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationActiveSyncQueuedAsAsyncMsgCount gets a reference to the given int64 and assigns it to the ReplicationActiveSyncQueuedAsAsyncMsgCount field.
func (o *MsgVpn) SetReplicationActiveSyncQueuedAsAsyncMsgCount(v int64) {
	o.ReplicationActiveSyncQueuedAsAsyncMsgCount = &v
}

// GetReplicationActiveSyncQueuedMsgCount returns the ReplicationActiveSyncQueuedMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveSyncQueuedMsgCount() int64 {
	if o == nil || o.ReplicationActiveSyncQueuedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationActiveSyncQueuedMsgCount
}

// GetReplicationActiveSyncQueuedMsgCountOk returns a tuple with the ReplicationActiveSyncQueuedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveSyncQueuedMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationActiveSyncQueuedMsgCount == nil {
		return nil, false
	}
	return o.ReplicationActiveSyncQueuedMsgCount, true
}

// HasReplicationActiveSyncQueuedMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveSyncQueuedMsgCount() bool {
	if o != nil && o.ReplicationActiveSyncQueuedMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationActiveSyncQueuedMsgCount gets a reference to the given int64 and assigns it to the ReplicationActiveSyncQueuedMsgCount field.
func (o *MsgVpn) SetReplicationActiveSyncQueuedMsgCount(v int64) {
	o.ReplicationActiveSyncQueuedMsgCount = &v
}

// GetReplicationActiveTransitionToSyncIneligibleCount returns the ReplicationActiveTransitionToSyncIneligibleCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationActiveTransitionToSyncIneligibleCount() int64 {
	if o == nil || o.ReplicationActiveTransitionToSyncIneligibleCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationActiveTransitionToSyncIneligibleCount
}

// GetReplicationActiveTransitionToSyncIneligibleCountOk returns a tuple with the ReplicationActiveTransitionToSyncIneligibleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationActiveTransitionToSyncIneligibleCountOk() (*int64, bool) {
	if o == nil || o.ReplicationActiveTransitionToSyncIneligibleCount == nil {
		return nil, false
	}
	return o.ReplicationActiveTransitionToSyncIneligibleCount, true
}

// HasReplicationActiveTransitionToSyncIneligibleCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationActiveTransitionToSyncIneligibleCount() bool {
	if o != nil && o.ReplicationActiveTransitionToSyncIneligibleCount != nil {
		return true
	}

	return false
}

// SetReplicationActiveTransitionToSyncIneligibleCount gets a reference to the given int64 and assigns it to the ReplicationActiveTransitionToSyncIneligibleCount field.
func (o *MsgVpn) SetReplicationActiveTransitionToSyncIneligibleCount(v int64) {
	o.ReplicationActiveTransitionToSyncIneligibleCount = &v
}

// GetReplicationBridgeAuthenticationBasicClientUsername returns the ReplicationBridgeAuthenticationBasicClientUsername field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationBridgeAuthenticationBasicClientUsername() string {
	if o == nil || o.ReplicationBridgeAuthenticationBasicClientUsername == nil {
		var ret string
		return ret
	}
	return *o.ReplicationBridgeAuthenticationBasicClientUsername
}

// GetReplicationBridgeAuthenticationBasicClientUsernameOk returns a tuple with the ReplicationBridgeAuthenticationBasicClientUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationBridgeAuthenticationBasicClientUsernameOk() (*string, bool) {
	if o == nil || o.ReplicationBridgeAuthenticationBasicClientUsername == nil {
		return nil, false
	}
	return o.ReplicationBridgeAuthenticationBasicClientUsername, true
}

// HasReplicationBridgeAuthenticationBasicClientUsername returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationBridgeAuthenticationBasicClientUsername() bool {
	if o != nil && o.ReplicationBridgeAuthenticationBasicClientUsername != nil {
		return true
	}

	return false
}

// SetReplicationBridgeAuthenticationBasicClientUsername gets a reference to the given string and assigns it to the ReplicationBridgeAuthenticationBasicClientUsername field.
func (o *MsgVpn) SetReplicationBridgeAuthenticationBasicClientUsername(v string) {
	o.ReplicationBridgeAuthenticationBasicClientUsername = &v
}

// GetReplicationBridgeAuthenticationScheme returns the ReplicationBridgeAuthenticationScheme field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationBridgeAuthenticationScheme() string {
	if o == nil || o.ReplicationBridgeAuthenticationScheme == nil {
		var ret string
		return ret
	}
	return *o.ReplicationBridgeAuthenticationScheme
}

// GetReplicationBridgeAuthenticationSchemeOk returns a tuple with the ReplicationBridgeAuthenticationScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationBridgeAuthenticationSchemeOk() (*string, bool) {
	if o == nil || o.ReplicationBridgeAuthenticationScheme == nil {
		return nil, false
	}
	return o.ReplicationBridgeAuthenticationScheme, true
}

// HasReplicationBridgeAuthenticationScheme returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationBridgeAuthenticationScheme() bool {
	if o != nil && o.ReplicationBridgeAuthenticationScheme != nil {
		return true
	}

	return false
}

// SetReplicationBridgeAuthenticationScheme gets a reference to the given string and assigns it to the ReplicationBridgeAuthenticationScheme field.
func (o *MsgVpn) SetReplicationBridgeAuthenticationScheme(v string) {
	o.ReplicationBridgeAuthenticationScheme = &v
}

// GetReplicationBridgeBoundToQueue returns the ReplicationBridgeBoundToQueue field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationBridgeBoundToQueue() bool {
	if o == nil || o.ReplicationBridgeBoundToQueue == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationBridgeBoundToQueue
}

// GetReplicationBridgeBoundToQueueOk returns a tuple with the ReplicationBridgeBoundToQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationBridgeBoundToQueueOk() (*bool, bool) {
	if o == nil || o.ReplicationBridgeBoundToQueue == nil {
		return nil, false
	}
	return o.ReplicationBridgeBoundToQueue, true
}

// HasReplicationBridgeBoundToQueue returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationBridgeBoundToQueue() bool {
	if o != nil && o.ReplicationBridgeBoundToQueue != nil {
		return true
	}

	return false
}

// SetReplicationBridgeBoundToQueue gets a reference to the given bool and assigns it to the ReplicationBridgeBoundToQueue field.
func (o *MsgVpn) SetReplicationBridgeBoundToQueue(v bool) {
	o.ReplicationBridgeBoundToQueue = &v
}

// GetReplicationBridgeCompressedDataEnabled returns the ReplicationBridgeCompressedDataEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationBridgeCompressedDataEnabled() bool {
	if o == nil || o.ReplicationBridgeCompressedDataEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationBridgeCompressedDataEnabled
}

// GetReplicationBridgeCompressedDataEnabledOk returns a tuple with the ReplicationBridgeCompressedDataEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationBridgeCompressedDataEnabledOk() (*bool, bool) {
	if o == nil || o.ReplicationBridgeCompressedDataEnabled == nil {
		return nil, false
	}
	return o.ReplicationBridgeCompressedDataEnabled, true
}

// HasReplicationBridgeCompressedDataEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationBridgeCompressedDataEnabled() bool {
	if o != nil && o.ReplicationBridgeCompressedDataEnabled != nil {
		return true
	}

	return false
}

// SetReplicationBridgeCompressedDataEnabled gets a reference to the given bool and assigns it to the ReplicationBridgeCompressedDataEnabled field.
func (o *MsgVpn) SetReplicationBridgeCompressedDataEnabled(v bool) {
	o.ReplicationBridgeCompressedDataEnabled = &v
}

// GetReplicationBridgeEgressFlowWindowSize returns the ReplicationBridgeEgressFlowWindowSize field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationBridgeEgressFlowWindowSize() int64 {
	if o == nil || o.ReplicationBridgeEgressFlowWindowSize == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationBridgeEgressFlowWindowSize
}

// GetReplicationBridgeEgressFlowWindowSizeOk returns a tuple with the ReplicationBridgeEgressFlowWindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationBridgeEgressFlowWindowSizeOk() (*int64, bool) {
	if o == nil || o.ReplicationBridgeEgressFlowWindowSize == nil {
		return nil, false
	}
	return o.ReplicationBridgeEgressFlowWindowSize, true
}

// HasReplicationBridgeEgressFlowWindowSize returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationBridgeEgressFlowWindowSize() bool {
	if o != nil && o.ReplicationBridgeEgressFlowWindowSize != nil {
		return true
	}

	return false
}

// SetReplicationBridgeEgressFlowWindowSize gets a reference to the given int64 and assigns it to the ReplicationBridgeEgressFlowWindowSize field.
func (o *MsgVpn) SetReplicationBridgeEgressFlowWindowSize(v int64) {
	o.ReplicationBridgeEgressFlowWindowSize = &v
}

// GetReplicationBridgeName returns the ReplicationBridgeName field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationBridgeName() string {
	if o == nil || o.ReplicationBridgeName == nil {
		var ret string
		return ret
	}
	return *o.ReplicationBridgeName
}

// GetReplicationBridgeNameOk returns a tuple with the ReplicationBridgeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationBridgeNameOk() (*string, bool) {
	if o == nil || o.ReplicationBridgeName == nil {
		return nil, false
	}
	return o.ReplicationBridgeName, true
}

// HasReplicationBridgeName returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationBridgeName() bool {
	if o != nil && o.ReplicationBridgeName != nil {
		return true
	}

	return false
}

// SetReplicationBridgeName gets a reference to the given string and assigns it to the ReplicationBridgeName field.
func (o *MsgVpn) SetReplicationBridgeName(v string) {
	o.ReplicationBridgeName = &v
}

// GetReplicationBridgeRetryDelay returns the ReplicationBridgeRetryDelay field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationBridgeRetryDelay() int64 {
	if o == nil || o.ReplicationBridgeRetryDelay == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationBridgeRetryDelay
}

// GetReplicationBridgeRetryDelayOk returns a tuple with the ReplicationBridgeRetryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationBridgeRetryDelayOk() (*int64, bool) {
	if o == nil || o.ReplicationBridgeRetryDelay == nil {
		return nil, false
	}
	return o.ReplicationBridgeRetryDelay, true
}

// HasReplicationBridgeRetryDelay returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationBridgeRetryDelay() bool {
	if o != nil && o.ReplicationBridgeRetryDelay != nil {
		return true
	}

	return false
}

// SetReplicationBridgeRetryDelay gets a reference to the given int64 and assigns it to the ReplicationBridgeRetryDelay field.
func (o *MsgVpn) SetReplicationBridgeRetryDelay(v int64) {
	o.ReplicationBridgeRetryDelay = &v
}

// GetReplicationBridgeTlsEnabled returns the ReplicationBridgeTlsEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationBridgeTlsEnabled() bool {
	if o == nil || o.ReplicationBridgeTlsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationBridgeTlsEnabled
}

// GetReplicationBridgeTlsEnabledOk returns a tuple with the ReplicationBridgeTlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationBridgeTlsEnabledOk() (*bool, bool) {
	if o == nil || o.ReplicationBridgeTlsEnabled == nil {
		return nil, false
	}
	return o.ReplicationBridgeTlsEnabled, true
}

// HasReplicationBridgeTlsEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationBridgeTlsEnabled() bool {
	if o != nil && o.ReplicationBridgeTlsEnabled != nil {
		return true
	}

	return false
}

// SetReplicationBridgeTlsEnabled gets a reference to the given bool and assigns it to the ReplicationBridgeTlsEnabled field.
func (o *MsgVpn) SetReplicationBridgeTlsEnabled(v bool) {
	o.ReplicationBridgeTlsEnabled = &v
}

// GetReplicationBridgeUnidirectionalClientProfileName returns the ReplicationBridgeUnidirectionalClientProfileName field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationBridgeUnidirectionalClientProfileName() string {
	if o == nil || o.ReplicationBridgeUnidirectionalClientProfileName == nil {
		var ret string
		return ret
	}
	return *o.ReplicationBridgeUnidirectionalClientProfileName
}

// GetReplicationBridgeUnidirectionalClientProfileNameOk returns a tuple with the ReplicationBridgeUnidirectionalClientProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationBridgeUnidirectionalClientProfileNameOk() (*string, bool) {
	if o == nil || o.ReplicationBridgeUnidirectionalClientProfileName == nil {
		return nil, false
	}
	return o.ReplicationBridgeUnidirectionalClientProfileName, true
}

// HasReplicationBridgeUnidirectionalClientProfileName returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationBridgeUnidirectionalClientProfileName() bool {
	if o != nil && o.ReplicationBridgeUnidirectionalClientProfileName != nil {
		return true
	}

	return false
}

// SetReplicationBridgeUnidirectionalClientProfileName gets a reference to the given string and assigns it to the ReplicationBridgeUnidirectionalClientProfileName field.
func (o *MsgVpn) SetReplicationBridgeUnidirectionalClientProfileName(v string) {
	o.ReplicationBridgeUnidirectionalClientProfileName = &v
}

// GetReplicationBridgeUp returns the ReplicationBridgeUp field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationBridgeUp() bool {
	if o == nil || o.ReplicationBridgeUp == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationBridgeUp
}

// GetReplicationBridgeUpOk returns a tuple with the ReplicationBridgeUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationBridgeUpOk() (*bool, bool) {
	if o == nil || o.ReplicationBridgeUp == nil {
		return nil, false
	}
	return o.ReplicationBridgeUp, true
}

// HasReplicationBridgeUp returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationBridgeUp() bool {
	if o != nil && o.ReplicationBridgeUp != nil {
		return true
	}

	return false
}

// SetReplicationBridgeUp gets a reference to the given bool and assigns it to the ReplicationBridgeUp field.
func (o *MsgVpn) SetReplicationBridgeUp(v bool) {
	o.ReplicationBridgeUp = &v
}

// GetReplicationEnabled returns the ReplicationEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationEnabled() bool {
	if o == nil || o.ReplicationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationEnabled
}

// GetReplicationEnabledOk returns a tuple with the ReplicationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationEnabledOk() (*bool, bool) {
	if o == nil || o.ReplicationEnabled == nil {
		return nil, false
	}
	return o.ReplicationEnabled, true
}

// HasReplicationEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationEnabled() bool {
	if o != nil && o.ReplicationEnabled != nil {
		return true
	}

	return false
}

// SetReplicationEnabled gets a reference to the given bool and assigns it to the ReplicationEnabled field.
func (o *MsgVpn) SetReplicationEnabled(v bool) {
	o.ReplicationEnabled = &v
}

// GetReplicationQueueBound returns the ReplicationQueueBound field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationQueueBound() bool {
	if o == nil || o.ReplicationQueueBound == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationQueueBound
}

// GetReplicationQueueBoundOk returns a tuple with the ReplicationQueueBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationQueueBoundOk() (*bool, bool) {
	if o == nil || o.ReplicationQueueBound == nil {
		return nil, false
	}
	return o.ReplicationQueueBound, true
}

// HasReplicationQueueBound returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationQueueBound() bool {
	if o != nil && o.ReplicationQueueBound != nil {
		return true
	}

	return false
}

// SetReplicationQueueBound gets a reference to the given bool and assigns it to the ReplicationQueueBound field.
func (o *MsgVpn) SetReplicationQueueBound(v bool) {
	o.ReplicationQueueBound = &v
}

// GetReplicationQueueMaxMsgSpoolUsage returns the ReplicationQueueMaxMsgSpoolUsage field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationQueueMaxMsgSpoolUsage() int64 {
	if o == nil || o.ReplicationQueueMaxMsgSpoolUsage == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationQueueMaxMsgSpoolUsage
}

// GetReplicationQueueMaxMsgSpoolUsageOk returns a tuple with the ReplicationQueueMaxMsgSpoolUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationQueueMaxMsgSpoolUsageOk() (*int64, bool) {
	if o == nil || o.ReplicationQueueMaxMsgSpoolUsage == nil {
		return nil, false
	}
	return o.ReplicationQueueMaxMsgSpoolUsage, true
}

// HasReplicationQueueMaxMsgSpoolUsage returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationQueueMaxMsgSpoolUsage() bool {
	if o != nil && o.ReplicationQueueMaxMsgSpoolUsage != nil {
		return true
	}

	return false
}

// SetReplicationQueueMaxMsgSpoolUsage gets a reference to the given int64 and assigns it to the ReplicationQueueMaxMsgSpoolUsage field.
func (o *MsgVpn) SetReplicationQueueMaxMsgSpoolUsage(v int64) {
	o.ReplicationQueueMaxMsgSpoolUsage = &v
}

// GetReplicationQueueRejectMsgToSenderOnDiscardEnabled returns the ReplicationQueueRejectMsgToSenderOnDiscardEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationQueueRejectMsgToSenderOnDiscardEnabled() bool {
	if o == nil || o.ReplicationQueueRejectMsgToSenderOnDiscardEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationQueueRejectMsgToSenderOnDiscardEnabled
}

// GetReplicationQueueRejectMsgToSenderOnDiscardEnabledOk returns a tuple with the ReplicationQueueRejectMsgToSenderOnDiscardEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationQueueRejectMsgToSenderOnDiscardEnabledOk() (*bool, bool) {
	if o == nil || o.ReplicationQueueRejectMsgToSenderOnDiscardEnabled == nil {
		return nil, false
	}
	return o.ReplicationQueueRejectMsgToSenderOnDiscardEnabled, true
}

// HasReplicationQueueRejectMsgToSenderOnDiscardEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationQueueRejectMsgToSenderOnDiscardEnabled() bool {
	if o != nil && o.ReplicationQueueRejectMsgToSenderOnDiscardEnabled != nil {
		return true
	}

	return false
}

// SetReplicationQueueRejectMsgToSenderOnDiscardEnabled gets a reference to the given bool and assigns it to the ReplicationQueueRejectMsgToSenderOnDiscardEnabled field.
func (o *MsgVpn) SetReplicationQueueRejectMsgToSenderOnDiscardEnabled(v bool) {
	o.ReplicationQueueRejectMsgToSenderOnDiscardEnabled = &v
}

// GetReplicationRejectMsgWhenSyncIneligibleEnabled returns the ReplicationRejectMsgWhenSyncIneligibleEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationRejectMsgWhenSyncIneligibleEnabled() bool {
	if o == nil || o.ReplicationRejectMsgWhenSyncIneligibleEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationRejectMsgWhenSyncIneligibleEnabled
}

// GetReplicationRejectMsgWhenSyncIneligibleEnabledOk returns a tuple with the ReplicationRejectMsgWhenSyncIneligibleEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationRejectMsgWhenSyncIneligibleEnabledOk() (*bool, bool) {
	if o == nil || o.ReplicationRejectMsgWhenSyncIneligibleEnabled == nil {
		return nil, false
	}
	return o.ReplicationRejectMsgWhenSyncIneligibleEnabled, true
}

// HasReplicationRejectMsgWhenSyncIneligibleEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationRejectMsgWhenSyncIneligibleEnabled() bool {
	if o != nil && o.ReplicationRejectMsgWhenSyncIneligibleEnabled != nil {
		return true
	}

	return false
}

// SetReplicationRejectMsgWhenSyncIneligibleEnabled gets a reference to the given bool and assigns it to the ReplicationRejectMsgWhenSyncIneligibleEnabled field.
func (o *MsgVpn) SetReplicationRejectMsgWhenSyncIneligibleEnabled(v bool) {
	o.ReplicationRejectMsgWhenSyncIneligibleEnabled = &v
}

// GetReplicationRemoteBridgeName returns the ReplicationRemoteBridgeName field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationRemoteBridgeName() string {
	if o == nil || o.ReplicationRemoteBridgeName == nil {
		var ret string
		return ret
	}
	return *o.ReplicationRemoteBridgeName
}

// GetReplicationRemoteBridgeNameOk returns a tuple with the ReplicationRemoteBridgeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationRemoteBridgeNameOk() (*string, bool) {
	if o == nil || o.ReplicationRemoteBridgeName == nil {
		return nil, false
	}
	return o.ReplicationRemoteBridgeName, true
}

// HasReplicationRemoteBridgeName returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationRemoteBridgeName() bool {
	if o != nil && o.ReplicationRemoteBridgeName != nil {
		return true
	}

	return false
}

// SetReplicationRemoteBridgeName gets a reference to the given string and assigns it to the ReplicationRemoteBridgeName field.
func (o *MsgVpn) SetReplicationRemoteBridgeName(v string) {
	o.ReplicationRemoteBridgeName = &v
}

// GetReplicationRemoteBridgeUp returns the ReplicationRemoteBridgeUp field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationRemoteBridgeUp() bool {
	if o == nil || o.ReplicationRemoteBridgeUp == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationRemoteBridgeUp
}

// GetReplicationRemoteBridgeUpOk returns a tuple with the ReplicationRemoteBridgeUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationRemoteBridgeUpOk() (*bool, bool) {
	if o == nil || o.ReplicationRemoteBridgeUp == nil {
		return nil, false
	}
	return o.ReplicationRemoteBridgeUp, true
}

// HasReplicationRemoteBridgeUp returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationRemoteBridgeUp() bool {
	if o != nil && o.ReplicationRemoteBridgeUp != nil {
		return true
	}

	return false
}

// SetReplicationRemoteBridgeUp gets a reference to the given bool and assigns it to the ReplicationRemoteBridgeUp field.
func (o *MsgVpn) SetReplicationRemoteBridgeUp(v bool) {
	o.ReplicationRemoteBridgeUp = &v
}

// GetReplicationRole returns the ReplicationRole field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationRole() string {
	if o == nil || o.ReplicationRole == nil {
		var ret string
		return ret
	}
	return *o.ReplicationRole
}

// GetReplicationRoleOk returns a tuple with the ReplicationRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationRoleOk() (*string, bool) {
	if o == nil || o.ReplicationRole == nil {
		return nil, false
	}
	return o.ReplicationRole, true
}

// HasReplicationRole returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationRole() bool {
	if o != nil && o.ReplicationRole != nil {
		return true
	}

	return false
}

// SetReplicationRole gets a reference to the given string and assigns it to the ReplicationRole field.
func (o *MsgVpn) SetReplicationRole(v string) {
	o.ReplicationRole = &v
}

// GetReplicationStandbyAckPropOutOfSeqRxMsgCount returns the ReplicationStandbyAckPropOutOfSeqRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationStandbyAckPropOutOfSeqRxMsgCount() int64 {
	if o == nil || o.ReplicationStandbyAckPropOutOfSeqRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyAckPropOutOfSeqRxMsgCount
}

// GetReplicationStandbyAckPropOutOfSeqRxMsgCountOk returns a tuple with the ReplicationStandbyAckPropOutOfSeqRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationStandbyAckPropOutOfSeqRxMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyAckPropOutOfSeqRxMsgCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyAckPropOutOfSeqRxMsgCount, true
}

// HasReplicationStandbyAckPropOutOfSeqRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationStandbyAckPropOutOfSeqRxMsgCount() bool {
	if o != nil && o.ReplicationStandbyAckPropOutOfSeqRxMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyAckPropOutOfSeqRxMsgCount gets a reference to the given int64 and assigns it to the ReplicationStandbyAckPropOutOfSeqRxMsgCount field.
func (o *MsgVpn) SetReplicationStandbyAckPropOutOfSeqRxMsgCount(v int64) {
	o.ReplicationStandbyAckPropOutOfSeqRxMsgCount = &v
}

// GetReplicationStandbyAckPropRxMsgCount returns the ReplicationStandbyAckPropRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationStandbyAckPropRxMsgCount() int64 {
	if o == nil || o.ReplicationStandbyAckPropRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyAckPropRxMsgCount
}

// GetReplicationStandbyAckPropRxMsgCountOk returns a tuple with the ReplicationStandbyAckPropRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationStandbyAckPropRxMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyAckPropRxMsgCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyAckPropRxMsgCount, true
}

// HasReplicationStandbyAckPropRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationStandbyAckPropRxMsgCount() bool {
	if o != nil && o.ReplicationStandbyAckPropRxMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyAckPropRxMsgCount gets a reference to the given int64 and assigns it to the ReplicationStandbyAckPropRxMsgCount field.
func (o *MsgVpn) SetReplicationStandbyAckPropRxMsgCount(v int64) {
	o.ReplicationStandbyAckPropRxMsgCount = &v
}

// GetReplicationStandbyReconcileRequestTxMsgCount returns the ReplicationStandbyReconcileRequestTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationStandbyReconcileRequestTxMsgCount() int64 {
	if o == nil || o.ReplicationStandbyReconcileRequestTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyReconcileRequestTxMsgCount
}

// GetReplicationStandbyReconcileRequestTxMsgCountOk returns a tuple with the ReplicationStandbyReconcileRequestTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationStandbyReconcileRequestTxMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyReconcileRequestTxMsgCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyReconcileRequestTxMsgCount, true
}

// HasReplicationStandbyReconcileRequestTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationStandbyReconcileRequestTxMsgCount() bool {
	if o != nil && o.ReplicationStandbyReconcileRequestTxMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyReconcileRequestTxMsgCount gets a reference to the given int64 and assigns it to the ReplicationStandbyReconcileRequestTxMsgCount field.
func (o *MsgVpn) SetReplicationStandbyReconcileRequestTxMsgCount(v int64) {
	o.ReplicationStandbyReconcileRequestTxMsgCount = &v
}

// GetReplicationStandbyRxMsgCount returns the ReplicationStandbyRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationStandbyRxMsgCount() int64 {
	if o == nil || o.ReplicationStandbyRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyRxMsgCount
}

// GetReplicationStandbyRxMsgCountOk returns a tuple with the ReplicationStandbyRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationStandbyRxMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyRxMsgCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyRxMsgCount, true
}

// HasReplicationStandbyRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationStandbyRxMsgCount() bool {
	if o != nil && o.ReplicationStandbyRxMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyRxMsgCount gets a reference to the given int64 and assigns it to the ReplicationStandbyRxMsgCount field.
func (o *MsgVpn) SetReplicationStandbyRxMsgCount(v int64) {
	o.ReplicationStandbyRxMsgCount = &v
}

// GetReplicationStandbyTransactionRequestCount returns the ReplicationStandbyTransactionRequestCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationStandbyTransactionRequestCount() int64 {
	if o == nil || o.ReplicationStandbyTransactionRequestCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyTransactionRequestCount
}

// GetReplicationStandbyTransactionRequestCountOk returns a tuple with the ReplicationStandbyTransactionRequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationStandbyTransactionRequestCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyTransactionRequestCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyTransactionRequestCount, true
}

// HasReplicationStandbyTransactionRequestCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationStandbyTransactionRequestCount() bool {
	if o != nil && o.ReplicationStandbyTransactionRequestCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyTransactionRequestCount gets a reference to the given int64 and assigns it to the ReplicationStandbyTransactionRequestCount field.
func (o *MsgVpn) SetReplicationStandbyTransactionRequestCount(v int64) {
	o.ReplicationStandbyTransactionRequestCount = &v
}

// GetReplicationStandbyTransactionRequestFailureCount returns the ReplicationStandbyTransactionRequestFailureCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationStandbyTransactionRequestFailureCount() int64 {
	if o == nil || o.ReplicationStandbyTransactionRequestFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyTransactionRequestFailureCount
}

// GetReplicationStandbyTransactionRequestFailureCountOk returns a tuple with the ReplicationStandbyTransactionRequestFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationStandbyTransactionRequestFailureCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyTransactionRequestFailureCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyTransactionRequestFailureCount, true
}

// HasReplicationStandbyTransactionRequestFailureCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationStandbyTransactionRequestFailureCount() bool {
	if o != nil && o.ReplicationStandbyTransactionRequestFailureCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyTransactionRequestFailureCount gets a reference to the given int64 and assigns it to the ReplicationStandbyTransactionRequestFailureCount field.
func (o *MsgVpn) SetReplicationStandbyTransactionRequestFailureCount(v int64) {
	o.ReplicationStandbyTransactionRequestFailureCount = &v
}

// GetReplicationStandbyTransactionRequestSuccessCount returns the ReplicationStandbyTransactionRequestSuccessCount field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationStandbyTransactionRequestSuccessCount() int64 {
	if o == nil || o.ReplicationStandbyTransactionRequestSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyTransactionRequestSuccessCount
}

// GetReplicationStandbyTransactionRequestSuccessCountOk returns a tuple with the ReplicationStandbyTransactionRequestSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationStandbyTransactionRequestSuccessCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyTransactionRequestSuccessCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyTransactionRequestSuccessCount, true
}

// HasReplicationStandbyTransactionRequestSuccessCount returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationStandbyTransactionRequestSuccessCount() bool {
	if o != nil && o.ReplicationStandbyTransactionRequestSuccessCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyTransactionRequestSuccessCount gets a reference to the given int64 and assigns it to the ReplicationStandbyTransactionRequestSuccessCount field.
func (o *MsgVpn) SetReplicationStandbyTransactionRequestSuccessCount(v int64) {
	o.ReplicationStandbyTransactionRequestSuccessCount = &v
}

// GetReplicationSyncEligible returns the ReplicationSyncEligible field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationSyncEligible() bool {
	if o == nil || o.ReplicationSyncEligible == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationSyncEligible
}

// GetReplicationSyncEligibleOk returns a tuple with the ReplicationSyncEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationSyncEligibleOk() (*bool, bool) {
	if o == nil || o.ReplicationSyncEligible == nil {
		return nil, false
	}
	return o.ReplicationSyncEligible, true
}

// HasReplicationSyncEligible returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationSyncEligible() bool {
	if o != nil && o.ReplicationSyncEligible != nil {
		return true
	}

	return false
}

// SetReplicationSyncEligible gets a reference to the given bool and assigns it to the ReplicationSyncEligible field.
func (o *MsgVpn) SetReplicationSyncEligible(v bool) {
	o.ReplicationSyncEligible = &v
}

// GetReplicationTransactionMode returns the ReplicationTransactionMode field value if set, zero value otherwise.
func (o *MsgVpn) GetReplicationTransactionMode() string {
	if o == nil || o.ReplicationTransactionMode == nil {
		var ret string
		return ret
	}
	return *o.ReplicationTransactionMode
}

// GetReplicationTransactionModeOk returns a tuple with the ReplicationTransactionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetReplicationTransactionModeOk() (*string, bool) {
	if o == nil || o.ReplicationTransactionMode == nil {
		return nil, false
	}
	return o.ReplicationTransactionMode, true
}

// HasReplicationTransactionMode returns a boolean if a field has been set.
func (o *MsgVpn) HasReplicationTransactionMode() bool {
	if o != nil && o.ReplicationTransactionMode != nil {
		return true
	}

	return false
}

// SetReplicationTransactionMode gets a reference to the given string and assigns it to the ReplicationTransactionMode field.
func (o *MsgVpn) SetReplicationTransactionMode(v string) {
	o.ReplicationTransactionMode = &v
}

// GetRestTlsServerCertEnforceTrustedCommonNameEnabled returns the RestTlsServerCertEnforceTrustedCommonNameEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetRestTlsServerCertEnforceTrustedCommonNameEnabled() bool {
	if o == nil || o.RestTlsServerCertEnforceTrustedCommonNameEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RestTlsServerCertEnforceTrustedCommonNameEnabled
}

// GetRestTlsServerCertEnforceTrustedCommonNameEnabledOk returns a tuple with the RestTlsServerCertEnforceTrustedCommonNameEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRestTlsServerCertEnforceTrustedCommonNameEnabledOk() (*bool, bool) {
	if o == nil || o.RestTlsServerCertEnforceTrustedCommonNameEnabled == nil {
		return nil, false
	}
	return o.RestTlsServerCertEnforceTrustedCommonNameEnabled, true
}

// HasRestTlsServerCertEnforceTrustedCommonNameEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasRestTlsServerCertEnforceTrustedCommonNameEnabled() bool {
	if o != nil && o.RestTlsServerCertEnforceTrustedCommonNameEnabled != nil {
		return true
	}

	return false
}

// SetRestTlsServerCertEnforceTrustedCommonNameEnabled gets a reference to the given bool and assigns it to the RestTlsServerCertEnforceTrustedCommonNameEnabled field.
func (o *MsgVpn) SetRestTlsServerCertEnforceTrustedCommonNameEnabled(v bool) {
	o.RestTlsServerCertEnforceTrustedCommonNameEnabled = &v
}

// GetRestTlsServerCertMaxChainDepth returns the RestTlsServerCertMaxChainDepth field value if set, zero value otherwise.
func (o *MsgVpn) GetRestTlsServerCertMaxChainDepth() int64 {
	if o == nil || o.RestTlsServerCertMaxChainDepth == nil {
		var ret int64
		return ret
	}
	return *o.RestTlsServerCertMaxChainDepth
}

// GetRestTlsServerCertMaxChainDepthOk returns a tuple with the RestTlsServerCertMaxChainDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRestTlsServerCertMaxChainDepthOk() (*int64, bool) {
	if o == nil || o.RestTlsServerCertMaxChainDepth == nil {
		return nil, false
	}
	return o.RestTlsServerCertMaxChainDepth, true
}

// HasRestTlsServerCertMaxChainDepth returns a boolean if a field has been set.
func (o *MsgVpn) HasRestTlsServerCertMaxChainDepth() bool {
	if o != nil && o.RestTlsServerCertMaxChainDepth != nil {
		return true
	}

	return false
}

// SetRestTlsServerCertMaxChainDepth gets a reference to the given int64 and assigns it to the RestTlsServerCertMaxChainDepth field.
func (o *MsgVpn) SetRestTlsServerCertMaxChainDepth(v int64) {
	o.RestTlsServerCertMaxChainDepth = &v
}

// GetRestTlsServerCertValidateDateEnabled returns the RestTlsServerCertValidateDateEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetRestTlsServerCertValidateDateEnabled() bool {
	if o == nil || o.RestTlsServerCertValidateDateEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RestTlsServerCertValidateDateEnabled
}

// GetRestTlsServerCertValidateDateEnabledOk returns a tuple with the RestTlsServerCertValidateDateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRestTlsServerCertValidateDateEnabledOk() (*bool, bool) {
	if o == nil || o.RestTlsServerCertValidateDateEnabled == nil {
		return nil, false
	}
	return o.RestTlsServerCertValidateDateEnabled, true
}

// HasRestTlsServerCertValidateDateEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasRestTlsServerCertValidateDateEnabled() bool {
	if o != nil && o.RestTlsServerCertValidateDateEnabled != nil {
		return true
	}

	return false
}

// SetRestTlsServerCertValidateDateEnabled gets a reference to the given bool and assigns it to the RestTlsServerCertValidateDateEnabled field.
func (o *MsgVpn) SetRestTlsServerCertValidateDateEnabled(v bool) {
	o.RestTlsServerCertValidateDateEnabled = &v
}

// GetRestTlsServerCertValidateNameEnabled returns the RestTlsServerCertValidateNameEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetRestTlsServerCertValidateNameEnabled() bool {
	if o == nil || o.RestTlsServerCertValidateNameEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RestTlsServerCertValidateNameEnabled
}

// GetRestTlsServerCertValidateNameEnabledOk returns a tuple with the RestTlsServerCertValidateNameEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRestTlsServerCertValidateNameEnabledOk() (*bool, bool) {
	if o == nil || o.RestTlsServerCertValidateNameEnabled == nil {
		return nil, false
	}
	return o.RestTlsServerCertValidateNameEnabled, true
}

// HasRestTlsServerCertValidateNameEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasRestTlsServerCertValidateNameEnabled() bool {
	if o != nil && o.RestTlsServerCertValidateNameEnabled != nil {
		return true
	}

	return false
}

// SetRestTlsServerCertValidateNameEnabled gets a reference to the given bool and assigns it to the RestTlsServerCertValidateNameEnabled field.
func (o *MsgVpn) SetRestTlsServerCertValidateNameEnabled(v bool) {
	o.RestTlsServerCertValidateNameEnabled = &v
}

// GetRxByteCount returns the RxByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetRxByteCount() int64 {
	if o == nil || o.RxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.RxByteCount
}

// GetRxByteCountOk returns a tuple with the RxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRxByteCountOk() (*int64, bool) {
	if o == nil || o.RxByteCount == nil {
		return nil, false
	}
	return o.RxByteCount, true
}

// HasRxByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasRxByteCount() bool {
	if o != nil && o.RxByteCount != nil {
		return true
	}

	return false
}

// SetRxByteCount gets a reference to the given int64 and assigns it to the RxByteCount field.
func (o *MsgVpn) SetRxByteCount(v int64) {
	o.RxByteCount = &v
}

// GetRxByteRate returns the RxByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetRxByteRate() int64 {
	if o == nil || o.RxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.RxByteRate
}

// GetRxByteRateOk returns a tuple with the RxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRxByteRateOk() (*int64, bool) {
	if o == nil || o.RxByteRate == nil {
		return nil, false
	}
	return o.RxByteRate, true
}

// HasRxByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasRxByteRate() bool {
	if o != nil && o.RxByteRate != nil {
		return true
	}

	return false
}

// SetRxByteRate gets a reference to the given int64 and assigns it to the RxByteRate field.
func (o *MsgVpn) SetRxByteRate(v int64) {
	o.RxByteRate = &v
}

// GetRxCompressedByteCount returns the RxCompressedByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetRxCompressedByteCount() int64 {
	if o == nil || o.RxCompressedByteCount == nil {
		var ret int64
		return ret
	}
	return *o.RxCompressedByteCount
}

// GetRxCompressedByteCountOk returns a tuple with the RxCompressedByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRxCompressedByteCountOk() (*int64, bool) {
	if o == nil || o.RxCompressedByteCount == nil {
		return nil, false
	}
	return o.RxCompressedByteCount, true
}

// HasRxCompressedByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasRxCompressedByteCount() bool {
	if o != nil && o.RxCompressedByteCount != nil {
		return true
	}

	return false
}

// SetRxCompressedByteCount gets a reference to the given int64 and assigns it to the RxCompressedByteCount field.
func (o *MsgVpn) SetRxCompressedByteCount(v int64) {
	o.RxCompressedByteCount = &v
}

// GetRxCompressedByteRate returns the RxCompressedByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetRxCompressedByteRate() int64 {
	if o == nil || o.RxCompressedByteRate == nil {
		var ret int64
		return ret
	}
	return *o.RxCompressedByteRate
}

// GetRxCompressedByteRateOk returns a tuple with the RxCompressedByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRxCompressedByteRateOk() (*int64, bool) {
	if o == nil || o.RxCompressedByteRate == nil {
		return nil, false
	}
	return o.RxCompressedByteRate, true
}

// HasRxCompressedByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasRxCompressedByteRate() bool {
	if o != nil && o.RxCompressedByteRate != nil {
		return true
	}

	return false
}

// SetRxCompressedByteRate gets a reference to the given int64 and assigns it to the RxCompressedByteRate field.
func (o *MsgVpn) SetRxCompressedByteRate(v int64) {
	o.RxCompressedByteRate = &v
}

// GetRxCompressionRatio returns the RxCompressionRatio field value if set, zero value otherwise.
func (o *MsgVpn) GetRxCompressionRatio() string {
	if o == nil || o.RxCompressionRatio == nil {
		var ret string
		return ret
	}
	return *o.RxCompressionRatio
}

// GetRxCompressionRatioOk returns a tuple with the RxCompressionRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRxCompressionRatioOk() (*string, bool) {
	if o == nil || o.RxCompressionRatio == nil {
		return nil, false
	}
	return o.RxCompressionRatio, true
}

// HasRxCompressionRatio returns a boolean if a field has been set.
func (o *MsgVpn) HasRxCompressionRatio() bool {
	if o != nil && o.RxCompressionRatio != nil {
		return true
	}

	return false
}

// SetRxCompressionRatio gets a reference to the given string and assigns it to the RxCompressionRatio field.
func (o *MsgVpn) SetRxCompressionRatio(v string) {
	o.RxCompressionRatio = &v
}

// GetRxMsgCount returns the RxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetRxMsgCount() int64 {
	if o == nil || o.RxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RxMsgCount
}

// GetRxMsgCountOk returns a tuple with the RxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRxMsgCountOk() (*int64, bool) {
	if o == nil || o.RxMsgCount == nil {
		return nil, false
	}
	return o.RxMsgCount, true
}

// HasRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasRxMsgCount() bool {
	if o != nil && o.RxMsgCount != nil {
		return true
	}

	return false
}

// SetRxMsgCount gets a reference to the given int64 and assigns it to the RxMsgCount field.
func (o *MsgVpn) SetRxMsgCount(v int64) {
	o.RxMsgCount = &v
}

// GetRxMsgRate returns the RxMsgRate field value if set, zero value otherwise.
func (o *MsgVpn) GetRxMsgRate() int64 {
	if o == nil || o.RxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.RxMsgRate
}

// GetRxMsgRateOk returns a tuple with the RxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRxMsgRateOk() (*int64, bool) {
	if o == nil || o.RxMsgRate == nil {
		return nil, false
	}
	return o.RxMsgRate, true
}

// HasRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpn) HasRxMsgRate() bool {
	if o != nil && o.RxMsgRate != nil {
		return true
	}

	return false
}

// SetRxMsgRate gets a reference to the given int64 and assigns it to the RxMsgRate field.
func (o *MsgVpn) SetRxMsgRate(v int64) {
	o.RxMsgRate = &v
}

// GetRxUncompressedByteCount returns the RxUncompressedByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetRxUncompressedByteCount() int64 {
	if o == nil || o.RxUncompressedByteCount == nil {
		var ret int64
		return ret
	}
	return *o.RxUncompressedByteCount
}

// GetRxUncompressedByteCountOk returns a tuple with the RxUncompressedByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRxUncompressedByteCountOk() (*int64, bool) {
	if o == nil || o.RxUncompressedByteCount == nil {
		return nil, false
	}
	return o.RxUncompressedByteCount, true
}

// HasRxUncompressedByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasRxUncompressedByteCount() bool {
	if o != nil && o.RxUncompressedByteCount != nil {
		return true
	}

	return false
}

// SetRxUncompressedByteCount gets a reference to the given int64 and assigns it to the RxUncompressedByteCount field.
func (o *MsgVpn) SetRxUncompressedByteCount(v int64) {
	o.RxUncompressedByteCount = &v
}

// GetRxUncompressedByteRate returns the RxUncompressedByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetRxUncompressedByteRate() int64 {
	if o == nil || o.RxUncompressedByteRate == nil {
		var ret int64
		return ret
	}
	return *o.RxUncompressedByteRate
}

// GetRxUncompressedByteRateOk returns a tuple with the RxUncompressedByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetRxUncompressedByteRateOk() (*int64, bool) {
	if o == nil || o.RxUncompressedByteRate == nil {
		return nil, false
	}
	return o.RxUncompressedByteRate, true
}

// HasRxUncompressedByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasRxUncompressedByteRate() bool {
	if o != nil && o.RxUncompressedByteRate != nil {
		return true
	}

	return false
}

// SetRxUncompressedByteRate gets a reference to the given int64 and assigns it to the RxUncompressedByteRate field.
func (o *MsgVpn) SetRxUncompressedByteRate(v int64) {
	o.RxUncompressedByteRate = &v
}

// GetSempOverMsgBusAdminClientEnabled returns the SempOverMsgBusAdminClientEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetSempOverMsgBusAdminClientEnabled() bool {
	if o == nil || o.SempOverMsgBusAdminClientEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SempOverMsgBusAdminClientEnabled
}

// GetSempOverMsgBusAdminClientEnabledOk returns a tuple with the SempOverMsgBusAdminClientEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetSempOverMsgBusAdminClientEnabledOk() (*bool, bool) {
	if o == nil || o.SempOverMsgBusAdminClientEnabled == nil {
		return nil, false
	}
	return o.SempOverMsgBusAdminClientEnabled, true
}

// HasSempOverMsgBusAdminClientEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasSempOverMsgBusAdminClientEnabled() bool {
	if o != nil && o.SempOverMsgBusAdminClientEnabled != nil {
		return true
	}

	return false
}

// SetSempOverMsgBusAdminClientEnabled gets a reference to the given bool and assigns it to the SempOverMsgBusAdminClientEnabled field.
func (o *MsgVpn) SetSempOverMsgBusAdminClientEnabled(v bool) {
	o.SempOverMsgBusAdminClientEnabled = &v
}

// GetSempOverMsgBusAdminDistributedCacheEnabled returns the SempOverMsgBusAdminDistributedCacheEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetSempOverMsgBusAdminDistributedCacheEnabled() bool {
	if o == nil || o.SempOverMsgBusAdminDistributedCacheEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SempOverMsgBusAdminDistributedCacheEnabled
}

// GetSempOverMsgBusAdminDistributedCacheEnabledOk returns a tuple with the SempOverMsgBusAdminDistributedCacheEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetSempOverMsgBusAdminDistributedCacheEnabledOk() (*bool, bool) {
	if o == nil || o.SempOverMsgBusAdminDistributedCacheEnabled == nil {
		return nil, false
	}
	return o.SempOverMsgBusAdminDistributedCacheEnabled, true
}

// HasSempOverMsgBusAdminDistributedCacheEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasSempOverMsgBusAdminDistributedCacheEnabled() bool {
	if o != nil && o.SempOverMsgBusAdminDistributedCacheEnabled != nil {
		return true
	}

	return false
}

// SetSempOverMsgBusAdminDistributedCacheEnabled gets a reference to the given bool and assigns it to the SempOverMsgBusAdminDistributedCacheEnabled field.
func (o *MsgVpn) SetSempOverMsgBusAdminDistributedCacheEnabled(v bool) {
	o.SempOverMsgBusAdminDistributedCacheEnabled = &v
}

// GetSempOverMsgBusAdminEnabled returns the SempOverMsgBusAdminEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetSempOverMsgBusAdminEnabled() bool {
	if o == nil || o.SempOverMsgBusAdminEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SempOverMsgBusAdminEnabled
}

// GetSempOverMsgBusAdminEnabledOk returns a tuple with the SempOverMsgBusAdminEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetSempOverMsgBusAdminEnabledOk() (*bool, bool) {
	if o == nil || o.SempOverMsgBusAdminEnabled == nil {
		return nil, false
	}
	return o.SempOverMsgBusAdminEnabled, true
}

// HasSempOverMsgBusAdminEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasSempOverMsgBusAdminEnabled() bool {
	if o != nil && o.SempOverMsgBusAdminEnabled != nil {
		return true
	}

	return false
}

// SetSempOverMsgBusAdminEnabled gets a reference to the given bool and assigns it to the SempOverMsgBusAdminEnabled field.
func (o *MsgVpn) SetSempOverMsgBusAdminEnabled(v bool) {
	o.SempOverMsgBusAdminEnabled = &v
}

// GetSempOverMsgBusEnabled returns the SempOverMsgBusEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetSempOverMsgBusEnabled() bool {
	if o == nil || o.SempOverMsgBusEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SempOverMsgBusEnabled
}

// GetSempOverMsgBusEnabledOk returns a tuple with the SempOverMsgBusEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetSempOverMsgBusEnabledOk() (*bool, bool) {
	if o == nil || o.SempOverMsgBusEnabled == nil {
		return nil, false
	}
	return o.SempOverMsgBusEnabled, true
}

// HasSempOverMsgBusEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasSempOverMsgBusEnabled() bool {
	if o != nil && o.SempOverMsgBusEnabled != nil {
		return true
	}

	return false
}

// SetSempOverMsgBusEnabled gets a reference to the given bool and assigns it to the SempOverMsgBusEnabled field.
func (o *MsgVpn) SetSempOverMsgBusEnabled(v bool) {
	o.SempOverMsgBusEnabled = &v
}

// GetSempOverMsgBusShowEnabled returns the SempOverMsgBusShowEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetSempOverMsgBusShowEnabled() bool {
	if o == nil || o.SempOverMsgBusShowEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SempOverMsgBusShowEnabled
}

// GetSempOverMsgBusShowEnabledOk returns a tuple with the SempOverMsgBusShowEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetSempOverMsgBusShowEnabledOk() (*bool, bool) {
	if o == nil || o.SempOverMsgBusShowEnabled == nil {
		return nil, false
	}
	return o.SempOverMsgBusShowEnabled, true
}

// HasSempOverMsgBusShowEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasSempOverMsgBusShowEnabled() bool {
	if o != nil && o.SempOverMsgBusShowEnabled != nil {
		return true
	}

	return false
}

// SetSempOverMsgBusShowEnabled gets a reference to the given bool and assigns it to the SempOverMsgBusShowEnabled field.
func (o *MsgVpn) SetSempOverMsgBusShowEnabled(v bool) {
	o.SempOverMsgBusShowEnabled = &v
}

// GetServiceAmqpMaxConnectionCount returns the ServiceAmqpMaxConnectionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpMaxConnectionCount() int64 {
	if o == nil || o.ServiceAmqpMaxConnectionCount == nil {
		var ret int64
		return ret
	}
	return *o.ServiceAmqpMaxConnectionCount
}

// GetServiceAmqpMaxConnectionCountOk returns a tuple with the ServiceAmqpMaxConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpMaxConnectionCountOk() (*int64, bool) {
	if o == nil || o.ServiceAmqpMaxConnectionCount == nil {
		return nil, false
	}
	return o.ServiceAmqpMaxConnectionCount, true
}

// HasServiceAmqpMaxConnectionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpMaxConnectionCount() bool {
	if o != nil && o.ServiceAmqpMaxConnectionCount != nil {
		return true
	}

	return false
}

// SetServiceAmqpMaxConnectionCount gets a reference to the given int64 and assigns it to the ServiceAmqpMaxConnectionCount field.
func (o *MsgVpn) SetServiceAmqpMaxConnectionCount(v int64) {
	o.ServiceAmqpMaxConnectionCount = &v
}

// GetServiceAmqpPlainTextCompressed returns the ServiceAmqpPlainTextCompressed field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpPlainTextCompressed() bool {
	if o == nil || o.ServiceAmqpPlainTextCompressed == nil {
		var ret bool
		return ret
	}
	return *o.ServiceAmqpPlainTextCompressed
}

// GetServiceAmqpPlainTextCompressedOk returns a tuple with the ServiceAmqpPlainTextCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpPlainTextCompressedOk() (*bool, bool) {
	if o == nil || o.ServiceAmqpPlainTextCompressed == nil {
		return nil, false
	}
	return o.ServiceAmqpPlainTextCompressed, true
}

// HasServiceAmqpPlainTextCompressed returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpPlainTextCompressed() bool {
	if o != nil && o.ServiceAmqpPlainTextCompressed != nil {
		return true
	}

	return false
}

// SetServiceAmqpPlainTextCompressed gets a reference to the given bool and assigns it to the ServiceAmqpPlainTextCompressed field.
func (o *MsgVpn) SetServiceAmqpPlainTextCompressed(v bool) {
	o.ServiceAmqpPlainTextCompressed = &v
}

// GetServiceAmqpPlainTextEnabled returns the ServiceAmqpPlainTextEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpPlainTextEnabled() bool {
	if o == nil || o.ServiceAmqpPlainTextEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceAmqpPlainTextEnabled
}

// GetServiceAmqpPlainTextEnabledOk returns a tuple with the ServiceAmqpPlainTextEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpPlainTextEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceAmqpPlainTextEnabled == nil {
		return nil, false
	}
	return o.ServiceAmqpPlainTextEnabled, true
}

// HasServiceAmqpPlainTextEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpPlainTextEnabled() bool {
	if o != nil && o.ServiceAmqpPlainTextEnabled != nil {
		return true
	}

	return false
}

// SetServiceAmqpPlainTextEnabled gets a reference to the given bool and assigns it to the ServiceAmqpPlainTextEnabled field.
func (o *MsgVpn) SetServiceAmqpPlainTextEnabled(v bool) {
	o.ServiceAmqpPlainTextEnabled = &v
}

// GetServiceAmqpPlainTextFailureReason returns the ServiceAmqpPlainTextFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpPlainTextFailureReason() string {
	if o == nil || o.ServiceAmqpPlainTextFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceAmqpPlainTextFailureReason
}

// GetServiceAmqpPlainTextFailureReasonOk returns a tuple with the ServiceAmqpPlainTextFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpPlainTextFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceAmqpPlainTextFailureReason == nil {
		return nil, false
	}
	return o.ServiceAmqpPlainTextFailureReason, true
}

// HasServiceAmqpPlainTextFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpPlainTextFailureReason() bool {
	if o != nil && o.ServiceAmqpPlainTextFailureReason != nil {
		return true
	}

	return false
}

// SetServiceAmqpPlainTextFailureReason gets a reference to the given string and assigns it to the ServiceAmqpPlainTextFailureReason field.
func (o *MsgVpn) SetServiceAmqpPlainTextFailureReason(v string) {
	o.ServiceAmqpPlainTextFailureReason = &v
}

// GetServiceAmqpPlainTextListenPort returns the ServiceAmqpPlainTextListenPort field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpPlainTextListenPort() int64 {
	if o == nil || o.ServiceAmqpPlainTextListenPort == nil {
		var ret int64
		return ret
	}
	return *o.ServiceAmqpPlainTextListenPort
}

// GetServiceAmqpPlainTextListenPortOk returns a tuple with the ServiceAmqpPlainTextListenPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpPlainTextListenPortOk() (*int64, bool) {
	if o == nil || o.ServiceAmqpPlainTextListenPort == nil {
		return nil, false
	}
	return o.ServiceAmqpPlainTextListenPort, true
}

// HasServiceAmqpPlainTextListenPort returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpPlainTextListenPort() bool {
	if o != nil && o.ServiceAmqpPlainTextListenPort != nil {
		return true
	}

	return false
}

// SetServiceAmqpPlainTextListenPort gets a reference to the given int64 and assigns it to the ServiceAmqpPlainTextListenPort field.
func (o *MsgVpn) SetServiceAmqpPlainTextListenPort(v int64) {
	o.ServiceAmqpPlainTextListenPort = &v
}

// GetServiceAmqpPlainTextUp returns the ServiceAmqpPlainTextUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpPlainTextUp() bool {
	if o == nil || o.ServiceAmqpPlainTextUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceAmqpPlainTextUp
}

// GetServiceAmqpPlainTextUpOk returns a tuple with the ServiceAmqpPlainTextUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpPlainTextUpOk() (*bool, bool) {
	if o == nil || o.ServiceAmqpPlainTextUp == nil {
		return nil, false
	}
	return o.ServiceAmqpPlainTextUp, true
}

// HasServiceAmqpPlainTextUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpPlainTextUp() bool {
	if o != nil && o.ServiceAmqpPlainTextUp != nil {
		return true
	}

	return false
}

// SetServiceAmqpPlainTextUp gets a reference to the given bool and assigns it to the ServiceAmqpPlainTextUp field.
func (o *MsgVpn) SetServiceAmqpPlainTextUp(v bool) {
	o.ServiceAmqpPlainTextUp = &v
}

// GetServiceAmqpTlsCompressed returns the ServiceAmqpTlsCompressed field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpTlsCompressed() bool {
	if o == nil || o.ServiceAmqpTlsCompressed == nil {
		var ret bool
		return ret
	}
	return *o.ServiceAmqpTlsCompressed
}

// GetServiceAmqpTlsCompressedOk returns a tuple with the ServiceAmqpTlsCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpTlsCompressedOk() (*bool, bool) {
	if o == nil || o.ServiceAmqpTlsCompressed == nil {
		return nil, false
	}
	return o.ServiceAmqpTlsCompressed, true
}

// HasServiceAmqpTlsCompressed returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpTlsCompressed() bool {
	if o != nil && o.ServiceAmqpTlsCompressed != nil {
		return true
	}

	return false
}

// SetServiceAmqpTlsCompressed gets a reference to the given bool and assigns it to the ServiceAmqpTlsCompressed field.
func (o *MsgVpn) SetServiceAmqpTlsCompressed(v bool) {
	o.ServiceAmqpTlsCompressed = &v
}

// GetServiceAmqpTlsEnabled returns the ServiceAmqpTlsEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpTlsEnabled() bool {
	if o == nil || o.ServiceAmqpTlsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceAmqpTlsEnabled
}

// GetServiceAmqpTlsEnabledOk returns a tuple with the ServiceAmqpTlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpTlsEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceAmqpTlsEnabled == nil {
		return nil, false
	}
	return o.ServiceAmqpTlsEnabled, true
}

// HasServiceAmqpTlsEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpTlsEnabled() bool {
	if o != nil && o.ServiceAmqpTlsEnabled != nil {
		return true
	}

	return false
}

// SetServiceAmqpTlsEnabled gets a reference to the given bool and assigns it to the ServiceAmqpTlsEnabled field.
func (o *MsgVpn) SetServiceAmqpTlsEnabled(v bool) {
	o.ServiceAmqpTlsEnabled = &v
}

// GetServiceAmqpTlsFailureReason returns the ServiceAmqpTlsFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpTlsFailureReason() string {
	if o == nil || o.ServiceAmqpTlsFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceAmqpTlsFailureReason
}

// GetServiceAmqpTlsFailureReasonOk returns a tuple with the ServiceAmqpTlsFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpTlsFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceAmqpTlsFailureReason == nil {
		return nil, false
	}
	return o.ServiceAmqpTlsFailureReason, true
}

// HasServiceAmqpTlsFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpTlsFailureReason() bool {
	if o != nil && o.ServiceAmqpTlsFailureReason != nil {
		return true
	}

	return false
}

// SetServiceAmqpTlsFailureReason gets a reference to the given string and assigns it to the ServiceAmqpTlsFailureReason field.
func (o *MsgVpn) SetServiceAmqpTlsFailureReason(v string) {
	o.ServiceAmqpTlsFailureReason = &v
}

// GetServiceAmqpTlsListenPort returns the ServiceAmqpTlsListenPort field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpTlsListenPort() int64 {
	if o == nil || o.ServiceAmqpTlsListenPort == nil {
		var ret int64
		return ret
	}
	return *o.ServiceAmqpTlsListenPort
}

// GetServiceAmqpTlsListenPortOk returns a tuple with the ServiceAmqpTlsListenPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpTlsListenPortOk() (*int64, bool) {
	if o == nil || o.ServiceAmqpTlsListenPort == nil {
		return nil, false
	}
	return o.ServiceAmqpTlsListenPort, true
}

// HasServiceAmqpTlsListenPort returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpTlsListenPort() bool {
	if o != nil && o.ServiceAmqpTlsListenPort != nil {
		return true
	}

	return false
}

// SetServiceAmqpTlsListenPort gets a reference to the given int64 and assigns it to the ServiceAmqpTlsListenPort field.
func (o *MsgVpn) SetServiceAmqpTlsListenPort(v int64) {
	o.ServiceAmqpTlsListenPort = &v
}

// GetServiceAmqpTlsUp returns the ServiceAmqpTlsUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceAmqpTlsUp() bool {
	if o == nil || o.ServiceAmqpTlsUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceAmqpTlsUp
}

// GetServiceAmqpTlsUpOk returns a tuple with the ServiceAmqpTlsUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceAmqpTlsUpOk() (*bool, bool) {
	if o == nil || o.ServiceAmqpTlsUp == nil {
		return nil, false
	}
	return o.ServiceAmqpTlsUp, true
}

// HasServiceAmqpTlsUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceAmqpTlsUp() bool {
	if o != nil && o.ServiceAmqpTlsUp != nil {
		return true
	}

	return false
}

// SetServiceAmqpTlsUp gets a reference to the given bool and assigns it to the ServiceAmqpTlsUp field.
func (o *MsgVpn) SetServiceAmqpTlsUp(v bool) {
	o.ServiceAmqpTlsUp = &v
}

// GetServiceMqttAuthenticationClientCertRequest returns the ServiceMqttAuthenticationClientCertRequest field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttAuthenticationClientCertRequest() string {
	if o == nil || o.ServiceMqttAuthenticationClientCertRequest == nil {
		var ret string
		return ret
	}
	return *o.ServiceMqttAuthenticationClientCertRequest
}

// GetServiceMqttAuthenticationClientCertRequestOk returns a tuple with the ServiceMqttAuthenticationClientCertRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttAuthenticationClientCertRequestOk() (*string, bool) {
	if o == nil || o.ServiceMqttAuthenticationClientCertRequest == nil {
		return nil, false
	}
	return o.ServiceMqttAuthenticationClientCertRequest, true
}

// HasServiceMqttAuthenticationClientCertRequest returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttAuthenticationClientCertRequest() bool {
	if o != nil && o.ServiceMqttAuthenticationClientCertRequest != nil {
		return true
	}

	return false
}

// SetServiceMqttAuthenticationClientCertRequest gets a reference to the given string and assigns it to the ServiceMqttAuthenticationClientCertRequest field.
func (o *MsgVpn) SetServiceMqttAuthenticationClientCertRequest(v string) {
	o.ServiceMqttAuthenticationClientCertRequest = &v
}

// GetServiceMqttMaxConnectionCount returns the ServiceMqttMaxConnectionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttMaxConnectionCount() int64 {
	if o == nil || o.ServiceMqttMaxConnectionCount == nil {
		var ret int64
		return ret
	}
	return *o.ServiceMqttMaxConnectionCount
}

// GetServiceMqttMaxConnectionCountOk returns a tuple with the ServiceMqttMaxConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttMaxConnectionCountOk() (*int64, bool) {
	if o == nil || o.ServiceMqttMaxConnectionCount == nil {
		return nil, false
	}
	return o.ServiceMqttMaxConnectionCount, true
}

// HasServiceMqttMaxConnectionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttMaxConnectionCount() bool {
	if o != nil && o.ServiceMqttMaxConnectionCount != nil {
		return true
	}

	return false
}

// SetServiceMqttMaxConnectionCount gets a reference to the given int64 and assigns it to the ServiceMqttMaxConnectionCount field.
func (o *MsgVpn) SetServiceMqttMaxConnectionCount(v int64) {
	o.ServiceMqttMaxConnectionCount = &v
}

// GetServiceMqttPlainTextCompressed returns the ServiceMqttPlainTextCompressed field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttPlainTextCompressed() bool {
	if o == nil || o.ServiceMqttPlainTextCompressed == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttPlainTextCompressed
}

// GetServiceMqttPlainTextCompressedOk returns a tuple with the ServiceMqttPlainTextCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttPlainTextCompressedOk() (*bool, bool) {
	if o == nil || o.ServiceMqttPlainTextCompressed == nil {
		return nil, false
	}
	return o.ServiceMqttPlainTextCompressed, true
}

// HasServiceMqttPlainTextCompressed returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttPlainTextCompressed() bool {
	if o != nil && o.ServiceMqttPlainTextCompressed != nil {
		return true
	}

	return false
}

// SetServiceMqttPlainTextCompressed gets a reference to the given bool and assigns it to the ServiceMqttPlainTextCompressed field.
func (o *MsgVpn) SetServiceMqttPlainTextCompressed(v bool) {
	o.ServiceMqttPlainTextCompressed = &v
}

// GetServiceMqttPlainTextEnabled returns the ServiceMqttPlainTextEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttPlainTextEnabled() bool {
	if o == nil || o.ServiceMqttPlainTextEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttPlainTextEnabled
}

// GetServiceMqttPlainTextEnabledOk returns a tuple with the ServiceMqttPlainTextEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttPlainTextEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceMqttPlainTextEnabled == nil {
		return nil, false
	}
	return o.ServiceMqttPlainTextEnabled, true
}

// HasServiceMqttPlainTextEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttPlainTextEnabled() bool {
	if o != nil && o.ServiceMqttPlainTextEnabled != nil {
		return true
	}

	return false
}

// SetServiceMqttPlainTextEnabled gets a reference to the given bool and assigns it to the ServiceMqttPlainTextEnabled field.
func (o *MsgVpn) SetServiceMqttPlainTextEnabled(v bool) {
	o.ServiceMqttPlainTextEnabled = &v
}

// GetServiceMqttPlainTextFailureReason returns the ServiceMqttPlainTextFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttPlainTextFailureReason() string {
	if o == nil || o.ServiceMqttPlainTextFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceMqttPlainTextFailureReason
}

// GetServiceMqttPlainTextFailureReasonOk returns a tuple with the ServiceMqttPlainTextFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttPlainTextFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceMqttPlainTextFailureReason == nil {
		return nil, false
	}
	return o.ServiceMqttPlainTextFailureReason, true
}

// HasServiceMqttPlainTextFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttPlainTextFailureReason() bool {
	if o != nil && o.ServiceMqttPlainTextFailureReason != nil {
		return true
	}

	return false
}

// SetServiceMqttPlainTextFailureReason gets a reference to the given string and assigns it to the ServiceMqttPlainTextFailureReason field.
func (o *MsgVpn) SetServiceMqttPlainTextFailureReason(v string) {
	o.ServiceMqttPlainTextFailureReason = &v
}

// GetServiceMqttPlainTextListenPort returns the ServiceMqttPlainTextListenPort field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttPlainTextListenPort() int64 {
	if o == nil || o.ServiceMqttPlainTextListenPort == nil {
		var ret int64
		return ret
	}
	return *o.ServiceMqttPlainTextListenPort
}

// GetServiceMqttPlainTextListenPortOk returns a tuple with the ServiceMqttPlainTextListenPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttPlainTextListenPortOk() (*int64, bool) {
	if o == nil || o.ServiceMqttPlainTextListenPort == nil {
		return nil, false
	}
	return o.ServiceMqttPlainTextListenPort, true
}

// HasServiceMqttPlainTextListenPort returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttPlainTextListenPort() bool {
	if o != nil && o.ServiceMqttPlainTextListenPort != nil {
		return true
	}

	return false
}

// SetServiceMqttPlainTextListenPort gets a reference to the given int64 and assigns it to the ServiceMqttPlainTextListenPort field.
func (o *MsgVpn) SetServiceMqttPlainTextListenPort(v int64) {
	o.ServiceMqttPlainTextListenPort = &v
}

// GetServiceMqttPlainTextUp returns the ServiceMqttPlainTextUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttPlainTextUp() bool {
	if o == nil || o.ServiceMqttPlainTextUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttPlainTextUp
}

// GetServiceMqttPlainTextUpOk returns a tuple with the ServiceMqttPlainTextUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttPlainTextUpOk() (*bool, bool) {
	if o == nil || o.ServiceMqttPlainTextUp == nil {
		return nil, false
	}
	return o.ServiceMqttPlainTextUp, true
}

// HasServiceMqttPlainTextUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttPlainTextUp() bool {
	if o != nil && o.ServiceMqttPlainTextUp != nil {
		return true
	}

	return false
}

// SetServiceMqttPlainTextUp gets a reference to the given bool and assigns it to the ServiceMqttPlainTextUp field.
func (o *MsgVpn) SetServiceMqttPlainTextUp(v bool) {
	o.ServiceMqttPlainTextUp = &v
}

// GetServiceMqttTlsCompressed returns the ServiceMqttTlsCompressed field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttTlsCompressed() bool {
	if o == nil || o.ServiceMqttTlsCompressed == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttTlsCompressed
}

// GetServiceMqttTlsCompressedOk returns a tuple with the ServiceMqttTlsCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttTlsCompressedOk() (*bool, bool) {
	if o == nil || o.ServiceMqttTlsCompressed == nil {
		return nil, false
	}
	return o.ServiceMqttTlsCompressed, true
}

// HasServiceMqttTlsCompressed returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttTlsCompressed() bool {
	if o != nil && o.ServiceMqttTlsCompressed != nil {
		return true
	}

	return false
}

// SetServiceMqttTlsCompressed gets a reference to the given bool and assigns it to the ServiceMqttTlsCompressed field.
func (o *MsgVpn) SetServiceMqttTlsCompressed(v bool) {
	o.ServiceMqttTlsCompressed = &v
}

// GetServiceMqttTlsEnabled returns the ServiceMqttTlsEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttTlsEnabled() bool {
	if o == nil || o.ServiceMqttTlsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttTlsEnabled
}

// GetServiceMqttTlsEnabledOk returns a tuple with the ServiceMqttTlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttTlsEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceMqttTlsEnabled == nil {
		return nil, false
	}
	return o.ServiceMqttTlsEnabled, true
}

// HasServiceMqttTlsEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttTlsEnabled() bool {
	if o != nil && o.ServiceMqttTlsEnabled != nil {
		return true
	}

	return false
}

// SetServiceMqttTlsEnabled gets a reference to the given bool and assigns it to the ServiceMqttTlsEnabled field.
func (o *MsgVpn) SetServiceMqttTlsEnabled(v bool) {
	o.ServiceMqttTlsEnabled = &v
}

// GetServiceMqttTlsFailureReason returns the ServiceMqttTlsFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttTlsFailureReason() string {
	if o == nil || o.ServiceMqttTlsFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceMqttTlsFailureReason
}

// GetServiceMqttTlsFailureReasonOk returns a tuple with the ServiceMqttTlsFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttTlsFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceMqttTlsFailureReason == nil {
		return nil, false
	}
	return o.ServiceMqttTlsFailureReason, true
}

// HasServiceMqttTlsFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttTlsFailureReason() bool {
	if o != nil && o.ServiceMqttTlsFailureReason != nil {
		return true
	}

	return false
}

// SetServiceMqttTlsFailureReason gets a reference to the given string and assigns it to the ServiceMqttTlsFailureReason field.
func (o *MsgVpn) SetServiceMqttTlsFailureReason(v string) {
	o.ServiceMqttTlsFailureReason = &v
}

// GetServiceMqttTlsListenPort returns the ServiceMqttTlsListenPort field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttTlsListenPort() int64 {
	if o == nil || o.ServiceMqttTlsListenPort == nil {
		var ret int64
		return ret
	}
	return *o.ServiceMqttTlsListenPort
}

// GetServiceMqttTlsListenPortOk returns a tuple with the ServiceMqttTlsListenPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttTlsListenPortOk() (*int64, bool) {
	if o == nil || o.ServiceMqttTlsListenPort == nil {
		return nil, false
	}
	return o.ServiceMqttTlsListenPort, true
}

// HasServiceMqttTlsListenPort returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttTlsListenPort() bool {
	if o != nil && o.ServiceMqttTlsListenPort != nil {
		return true
	}

	return false
}

// SetServiceMqttTlsListenPort gets a reference to the given int64 and assigns it to the ServiceMqttTlsListenPort field.
func (o *MsgVpn) SetServiceMqttTlsListenPort(v int64) {
	o.ServiceMqttTlsListenPort = &v
}

// GetServiceMqttTlsUp returns the ServiceMqttTlsUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttTlsUp() bool {
	if o == nil || o.ServiceMqttTlsUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttTlsUp
}

// GetServiceMqttTlsUpOk returns a tuple with the ServiceMqttTlsUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttTlsUpOk() (*bool, bool) {
	if o == nil || o.ServiceMqttTlsUp == nil {
		return nil, false
	}
	return o.ServiceMqttTlsUp, true
}

// HasServiceMqttTlsUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttTlsUp() bool {
	if o != nil && o.ServiceMqttTlsUp != nil {
		return true
	}

	return false
}

// SetServiceMqttTlsUp gets a reference to the given bool and assigns it to the ServiceMqttTlsUp field.
func (o *MsgVpn) SetServiceMqttTlsUp(v bool) {
	o.ServiceMqttTlsUp = &v
}

// GetServiceMqttTlsWebSocketCompressed returns the ServiceMqttTlsWebSocketCompressed field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttTlsWebSocketCompressed() bool {
	if o == nil || o.ServiceMqttTlsWebSocketCompressed == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttTlsWebSocketCompressed
}

// GetServiceMqttTlsWebSocketCompressedOk returns a tuple with the ServiceMqttTlsWebSocketCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttTlsWebSocketCompressedOk() (*bool, bool) {
	if o == nil || o.ServiceMqttTlsWebSocketCompressed == nil {
		return nil, false
	}
	return o.ServiceMqttTlsWebSocketCompressed, true
}

// HasServiceMqttTlsWebSocketCompressed returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttTlsWebSocketCompressed() bool {
	if o != nil && o.ServiceMqttTlsWebSocketCompressed != nil {
		return true
	}

	return false
}

// SetServiceMqttTlsWebSocketCompressed gets a reference to the given bool and assigns it to the ServiceMqttTlsWebSocketCompressed field.
func (o *MsgVpn) SetServiceMqttTlsWebSocketCompressed(v bool) {
	o.ServiceMqttTlsWebSocketCompressed = &v
}

// GetServiceMqttTlsWebSocketEnabled returns the ServiceMqttTlsWebSocketEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttTlsWebSocketEnabled() bool {
	if o == nil || o.ServiceMqttTlsWebSocketEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttTlsWebSocketEnabled
}

// GetServiceMqttTlsWebSocketEnabledOk returns a tuple with the ServiceMqttTlsWebSocketEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttTlsWebSocketEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceMqttTlsWebSocketEnabled == nil {
		return nil, false
	}
	return o.ServiceMqttTlsWebSocketEnabled, true
}

// HasServiceMqttTlsWebSocketEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttTlsWebSocketEnabled() bool {
	if o != nil && o.ServiceMqttTlsWebSocketEnabled != nil {
		return true
	}

	return false
}

// SetServiceMqttTlsWebSocketEnabled gets a reference to the given bool and assigns it to the ServiceMqttTlsWebSocketEnabled field.
func (o *MsgVpn) SetServiceMqttTlsWebSocketEnabled(v bool) {
	o.ServiceMqttTlsWebSocketEnabled = &v
}

// GetServiceMqttTlsWebSocketFailureReason returns the ServiceMqttTlsWebSocketFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttTlsWebSocketFailureReason() string {
	if o == nil || o.ServiceMqttTlsWebSocketFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceMqttTlsWebSocketFailureReason
}

// GetServiceMqttTlsWebSocketFailureReasonOk returns a tuple with the ServiceMqttTlsWebSocketFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttTlsWebSocketFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceMqttTlsWebSocketFailureReason == nil {
		return nil, false
	}
	return o.ServiceMqttTlsWebSocketFailureReason, true
}

// HasServiceMqttTlsWebSocketFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttTlsWebSocketFailureReason() bool {
	if o != nil && o.ServiceMqttTlsWebSocketFailureReason != nil {
		return true
	}

	return false
}

// SetServiceMqttTlsWebSocketFailureReason gets a reference to the given string and assigns it to the ServiceMqttTlsWebSocketFailureReason field.
func (o *MsgVpn) SetServiceMqttTlsWebSocketFailureReason(v string) {
	o.ServiceMqttTlsWebSocketFailureReason = &v
}

// GetServiceMqttTlsWebSocketListenPort returns the ServiceMqttTlsWebSocketListenPort field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttTlsWebSocketListenPort() int64 {
	if o == nil || o.ServiceMqttTlsWebSocketListenPort == nil {
		var ret int64
		return ret
	}
	return *o.ServiceMqttTlsWebSocketListenPort
}

// GetServiceMqttTlsWebSocketListenPortOk returns a tuple with the ServiceMqttTlsWebSocketListenPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttTlsWebSocketListenPortOk() (*int64, bool) {
	if o == nil || o.ServiceMqttTlsWebSocketListenPort == nil {
		return nil, false
	}
	return o.ServiceMqttTlsWebSocketListenPort, true
}

// HasServiceMqttTlsWebSocketListenPort returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttTlsWebSocketListenPort() bool {
	if o != nil && o.ServiceMqttTlsWebSocketListenPort != nil {
		return true
	}

	return false
}

// SetServiceMqttTlsWebSocketListenPort gets a reference to the given int64 and assigns it to the ServiceMqttTlsWebSocketListenPort field.
func (o *MsgVpn) SetServiceMqttTlsWebSocketListenPort(v int64) {
	o.ServiceMqttTlsWebSocketListenPort = &v
}

// GetServiceMqttTlsWebSocketUp returns the ServiceMqttTlsWebSocketUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttTlsWebSocketUp() bool {
	if o == nil || o.ServiceMqttTlsWebSocketUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttTlsWebSocketUp
}

// GetServiceMqttTlsWebSocketUpOk returns a tuple with the ServiceMqttTlsWebSocketUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttTlsWebSocketUpOk() (*bool, bool) {
	if o == nil || o.ServiceMqttTlsWebSocketUp == nil {
		return nil, false
	}
	return o.ServiceMqttTlsWebSocketUp, true
}

// HasServiceMqttTlsWebSocketUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttTlsWebSocketUp() bool {
	if o != nil && o.ServiceMqttTlsWebSocketUp != nil {
		return true
	}

	return false
}

// SetServiceMqttTlsWebSocketUp gets a reference to the given bool and assigns it to the ServiceMqttTlsWebSocketUp field.
func (o *MsgVpn) SetServiceMqttTlsWebSocketUp(v bool) {
	o.ServiceMqttTlsWebSocketUp = &v
}

// GetServiceMqttWebSocketCompressed returns the ServiceMqttWebSocketCompressed field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttWebSocketCompressed() bool {
	if o == nil || o.ServiceMqttWebSocketCompressed == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttWebSocketCompressed
}

// GetServiceMqttWebSocketCompressedOk returns a tuple with the ServiceMqttWebSocketCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttWebSocketCompressedOk() (*bool, bool) {
	if o == nil || o.ServiceMqttWebSocketCompressed == nil {
		return nil, false
	}
	return o.ServiceMqttWebSocketCompressed, true
}

// HasServiceMqttWebSocketCompressed returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttWebSocketCompressed() bool {
	if o != nil && o.ServiceMqttWebSocketCompressed != nil {
		return true
	}

	return false
}

// SetServiceMqttWebSocketCompressed gets a reference to the given bool and assigns it to the ServiceMqttWebSocketCompressed field.
func (o *MsgVpn) SetServiceMqttWebSocketCompressed(v bool) {
	o.ServiceMqttWebSocketCompressed = &v
}

// GetServiceMqttWebSocketEnabled returns the ServiceMqttWebSocketEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttWebSocketEnabled() bool {
	if o == nil || o.ServiceMqttWebSocketEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttWebSocketEnabled
}

// GetServiceMqttWebSocketEnabledOk returns a tuple with the ServiceMqttWebSocketEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttWebSocketEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceMqttWebSocketEnabled == nil {
		return nil, false
	}
	return o.ServiceMqttWebSocketEnabled, true
}

// HasServiceMqttWebSocketEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttWebSocketEnabled() bool {
	if o != nil && o.ServiceMqttWebSocketEnabled != nil {
		return true
	}

	return false
}

// SetServiceMqttWebSocketEnabled gets a reference to the given bool and assigns it to the ServiceMqttWebSocketEnabled field.
func (o *MsgVpn) SetServiceMqttWebSocketEnabled(v bool) {
	o.ServiceMqttWebSocketEnabled = &v
}

// GetServiceMqttWebSocketFailureReason returns the ServiceMqttWebSocketFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttWebSocketFailureReason() string {
	if o == nil || o.ServiceMqttWebSocketFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceMqttWebSocketFailureReason
}

// GetServiceMqttWebSocketFailureReasonOk returns a tuple with the ServiceMqttWebSocketFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttWebSocketFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceMqttWebSocketFailureReason == nil {
		return nil, false
	}
	return o.ServiceMqttWebSocketFailureReason, true
}

// HasServiceMqttWebSocketFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttWebSocketFailureReason() bool {
	if o != nil && o.ServiceMqttWebSocketFailureReason != nil {
		return true
	}

	return false
}

// SetServiceMqttWebSocketFailureReason gets a reference to the given string and assigns it to the ServiceMqttWebSocketFailureReason field.
func (o *MsgVpn) SetServiceMqttWebSocketFailureReason(v string) {
	o.ServiceMqttWebSocketFailureReason = &v
}

// GetServiceMqttWebSocketListenPort returns the ServiceMqttWebSocketListenPort field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttWebSocketListenPort() int64 {
	if o == nil || o.ServiceMqttWebSocketListenPort == nil {
		var ret int64
		return ret
	}
	return *o.ServiceMqttWebSocketListenPort
}

// GetServiceMqttWebSocketListenPortOk returns a tuple with the ServiceMqttWebSocketListenPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttWebSocketListenPortOk() (*int64, bool) {
	if o == nil || o.ServiceMqttWebSocketListenPort == nil {
		return nil, false
	}
	return o.ServiceMqttWebSocketListenPort, true
}

// HasServiceMqttWebSocketListenPort returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttWebSocketListenPort() bool {
	if o != nil && o.ServiceMqttWebSocketListenPort != nil {
		return true
	}

	return false
}

// SetServiceMqttWebSocketListenPort gets a reference to the given int64 and assigns it to the ServiceMqttWebSocketListenPort field.
func (o *MsgVpn) SetServiceMqttWebSocketListenPort(v int64) {
	o.ServiceMqttWebSocketListenPort = &v
}

// GetServiceMqttWebSocketUp returns the ServiceMqttWebSocketUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceMqttWebSocketUp() bool {
	if o == nil || o.ServiceMqttWebSocketUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceMqttWebSocketUp
}

// GetServiceMqttWebSocketUpOk returns a tuple with the ServiceMqttWebSocketUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceMqttWebSocketUpOk() (*bool, bool) {
	if o == nil || o.ServiceMqttWebSocketUp == nil {
		return nil, false
	}
	return o.ServiceMqttWebSocketUp, true
}

// HasServiceMqttWebSocketUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceMqttWebSocketUp() bool {
	if o != nil && o.ServiceMqttWebSocketUp != nil {
		return true
	}

	return false
}

// SetServiceMqttWebSocketUp gets a reference to the given bool and assigns it to the ServiceMqttWebSocketUp field.
func (o *MsgVpn) SetServiceMqttWebSocketUp(v bool) {
	o.ServiceMqttWebSocketUp = &v
}

// GetServiceRestIncomingAuthenticationClientCertRequest returns the ServiceRestIncomingAuthenticationClientCertRequest field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingAuthenticationClientCertRequest() string {
	if o == nil || o.ServiceRestIncomingAuthenticationClientCertRequest == nil {
		var ret string
		return ret
	}
	return *o.ServiceRestIncomingAuthenticationClientCertRequest
}

// GetServiceRestIncomingAuthenticationClientCertRequestOk returns a tuple with the ServiceRestIncomingAuthenticationClientCertRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingAuthenticationClientCertRequestOk() (*string, bool) {
	if o == nil || o.ServiceRestIncomingAuthenticationClientCertRequest == nil {
		return nil, false
	}
	return o.ServiceRestIncomingAuthenticationClientCertRequest, true
}

// HasServiceRestIncomingAuthenticationClientCertRequest returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingAuthenticationClientCertRequest() bool {
	if o != nil && o.ServiceRestIncomingAuthenticationClientCertRequest != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingAuthenticationClientCertRequest gets a reference to the given string and assigns it to the ServiceRestIncomingAuthenticationClientCertRequest field.
func (o *MsgVpn) SetServiceRestIncomingAuthenticationClientCertRequest(v string) {
	o.ServiceRestIncomingAuthenticationClientCertRequest = &v
}

// GetServiceRestIncomingAuthorizationHeaderHandling returns the ServiceRestIncomingAuthorizationHeaderHandling field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingAuthorizationHeaderHandling() string {
	if o == nil || o.ServiceRestIncomingAuthorizationHeaderHandling == nil {
		var ret string
		return ret
	}
	return *o.ServiceRestIncomingAuthorizationHeaderHandling
}

// GetServiceRestIncomingAuthorizationHeaderHandlingOk returns a tuple with the ServiceRestIncomingAuthorizationHeaderHandling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingAuthorizationHeaderHandlingOk() (*string, bool) {
	if o == nil || o.ServiceRestIncomingAuthorizationHeaderHandling == nil {
		return nil, false
	}
	return o.ServiceRestIncomingAuthorizationHeaderHandling, true
}

// HasServiceRestIncomingAuthorizationHeaderHandling returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingAuthorizationHeaderHandling() bool {
	if o != nil && o.ServiceRestIncomingAuthorizationHeaderHandling != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingAuthorizationHeaderHandling gets a reference to the given string and assigns it to the ServiceRestIncomingAuthorizationHeaderHandling field.
func (o *MsgVpn) SetServiceRestIncomingAuthorizationHeaderHandling(v string) {
	o.ServiceRestIncomingAuthorizationHeaderHandling = &v
}

// GetServiceRestIncomingMaxConnectionCount returns the ServiceRestIncomingMaxConnectionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingMaxConnectionCount() int64 {
	if o == nil || o.ServiceRestIncomingMaxConnectionCount == nil {
		var ret int64
		return ret
	}
	return *o.ServiceRestIncomingMaxConnectionCount
}

// GetServiceRestIncomingMaxConnectionCountOk returns a tuple with the ServiceRestIncomingMaxConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingMaxConnectionCountOk() (*int64, bool) {
	if o == nil || o.ServiceRestIncomingMaxConnectionCount == nil {
		return nil, false
	}
	return o.ServiceRestIncomingMaxConnectionCount, true
}

// HasServiceRestIncomingMaxConnectionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingMaxConnectionCount() bool {
	if o != nil && o.ServiceRestIncomingMaxConnectionCount != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingMaxConnectionCount gets a reference to the given int64 and assigns it to the ServiceRestIncomingMaxConnectionCount field.
func (o *MsgVpn) SetServiceRestIncomingMaxConnectionCount(v int64) {
	o.ServiceRestIncomingMaxConnectionCount = &v
}

// GetServiceRestIncomingPlainTextCompressed returns the ServiceRestIncomingPlainTextCompressed field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingPlainTextCompressed() bool {
	if o == nil || o.ServiceRestIncomingPlainTextCompressed == nil {
		var ret bool
		return ret
	}
	return *o.ServiceRestIncomingPlainTextCompressed
}

// GetServiceRestIncomingPlainTextCompressedOk returns a tuple with the ServiceRestIncomingPlainTextCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingPlainTextCompressedOk() (*bool, bool) {
	if o == nil || o.ServiceRestIncomingPlainTextCompressed == nil {
		return nil, false
	}
	return o.ServiceRestIncomingPlainTextCompressed, true
}

// HasServiceRestIncomingPlainTextCompressed returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingPlainTextCompressed() bool {
	if o != nil && o.ServiceRestIncomingPlainTextCompressed != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingPlainTextCompressed gets a reference to the given bool and assigns it to the ServiceRestIncomingPlainTextCompressed field.
func (o *MsgVpn) SetServiceRestIncomingPlainTextCompressed(v bool) {
	o.ServiceRestIncomingPlainTextCompressed = &v
}

// GetServiceRestIncomingPlainTextEnabled returns the ServiceRestIncomingPlainTextEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingPlainTextEnabled() bool {
	if o == nil || o.ServiceRestIncomingPlainTextEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceRestIncomingPlainTextEnabled
}

// GetServiceRestIncomingPlainTextEnabledOk returns a tuple with the ServiceRestIncomingPlainTextEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingPlainTextEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceRestIncomingPlainTextEnabled == nil {
		return nil, false
	}
	return o.ServiceRestIncomingPlainTextEnabled, true
}

// HasServiceRestIncomingPlainTextEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingPlainTextEnabled() bool {
	if o != nil && o.ServiceRestIncomingPlainTextEnabled != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingPlainTextEnabled gets a reference to the given bool and assigns it to the ServiceRestIncomingPlainTextEnabled field.
func (o *MsgVpn) SetServiceRestIncomingPlainTextEnabled(v bool) {
	o.ServiceRestIncomingPlainTextEnabled = &v
}

// GetServiceRestIncomingPlainTextFailureReason returns the ServiceRestIncomingPlainTextFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingPlainTextFailureReason() string {
	if o == nil || o.ServiceRestIncomingPlainTextFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceRestIncomingPlainTextFailureReason
}

// GetServiceRestIncomingPlainTextFailureReasonOk returns a tuple with the ServiceRestIncomingPlainTextFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingPlainTextFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceRestIncomingPlainTextFailureReason == nil {
		return nil, false
	}
	return o.ServiceRestIncomingPlainTextFailureReason, true
}

// HasServiceRestIncomingPlainTextFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingPlainTextFailureReason() bool {
	if o != nil && o.ServiceRestIncomingPlainTextFailureReason != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingPlainTextFailureReason gets a reference to the given string and assigns it to the ServiceRestIncomingPlainTextFailureReason field.
func (o *MsgVpn) SetServiceRestIncomingPlainTextFailureReason(v string) {
	o.ServiceRestIncomingPlainTextFailureReason = &v
}

// GetServiceRestIncomingPlainTextListenPort returns the ServiceRestIncomingPlainTextListenPort field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingPlainTextListenPort() int64 {
	if o == nil || o.ServiceRestIncomingPlainTextListenPort == nil {
		var ret int64
		return ret
	}
	return *o.ServiceRestIncomingPlainTextListenPort
}

// GetServiceRestIncomingPlainTextListenPortOk returns a tuple with the ServiceRestIncomingPlainTextListenPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingPlainTextListenPortOk() (*int64, bool) {
	if o == nil || o.ServiceRestIncomingPlainTextListenPort == nil {
		return nil, false
	}
	return o.ServiceRestIncomingPlainTextListenPort, true
}

// HasServiceRestIncomingPlainTextListenPort returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingPlainTextListenPort() bool {
	if o != nil && o.ServiceRestIncomingPlainTextListenPort != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingPlainTextListenPort gets a reference to the given int64 and assigns it to the ServiceRestIncomingPlainTextListenPort field.
func (o *MsgVpn) SetServiceRestIncomingPlainTextListenPort(v int64) {
	o.ServiceRestIncomingPlainTextListenPort = &v
}

// GetServiceRestIncomingPlainTextUp returns the ServiceRestIncomingPlainTextUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingPlainTextUp() bool {
	if o == nil || o.ServiceRestIncomingPlainTextUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceRestIncomingPlainTextUp
}

// GetServiceRestIncomingPlainTextUpOk returns a tuple with the ServiceRestIncomingPlainTextUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingPlainTextUpOk() (*bool, bool) {
	if o == nil || o.ServiceRestIncomingPlainTextUp == nil {
		return nil, false
	}
	return o.ServiceRestIncomingPlainTextUp, true
}

// HasServiceRestIncomingPlainTextUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingPlainTextUp() bool {
	if o != nil && o.ServiceRestIncomingPlainTextUp != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingPlainTextUp gets a reference to the given bool and assigns it to the ServiceRestIncomingPlainTextUp field.
func (o *MsgVpn) SetServiceRestIncomingPlainTextUp(v bool) {
	o.ServiceRestIncomingPlainTextUp = &v
}

// GetServiceRestIncomingTlsCompressed returns the ServiceRestIncomingTlsCompressed field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingTlsCompressed() bool {
	if o == nil || o.ServiceRestIncomingTlsCompressed == nil {
		var ret bool
		return ret
	}
	return *o.ServiceRestIncomingTlsCompressed
}

// GetServiceRestIncomingTlsCompressedOk returns a tuple with the ServiceRestIncomingTlsCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingTlsCompressedOk() (*bool, bool) {
	if o == nil || o.ServiceRestIncomingTlsCompressed == nil {
		return nil, false
	}
	return o.ServiceRestIncomingTlsCompressed, true
}

// HasServiceRestIncomingTlsCompressed returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingTlsCompressed() bool {
	if o != nil && o.ServiceRestIncomingTlsCompressed != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingTlsCompressed gets a reference to the given bool and assigns it to the ServiceRestIncomingTlsCompressed field.
func (o *MsgVpn) SetServiceRestIncomingTlsCompressed(v bool) {
	o.ServiceRestIncomingTlsCompressed = &v
}

// GetServiceRestIncomingTlsEnabled returns the ServiceRestIncomingTlsEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingTlsEnabled() bool {
	if o == nil || o.ServiceRestIncomingTlsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceRestIncomingTlsEnabled
}

// GetServiceRestIncomingTlsEnabledOk returns a tuple with the ServiceRestIncomingTlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingTlsEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceRestIncomingTlsEnabled == nil {
		return nil, false
	}
	return o.ServiceRestIncomingTlsEnabled, true
}

// HasServiceRestIncomingTlsEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingTlsEnabled() bool {
	if o != nil && o.ServiceRestIncomingTlsEnabled != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingTlsEnabled gets a reference to the given bool and assigns it to the ServiceRestIncomingTlsEnabled field.
func (o *MsgVpn) SetServiceRestIncomingTlsEnabled(v bool) {
	o.ServiceRestIncomingTlsEnabled = &v
}

// GetServiceRestIncomingTlsFailureReason returns the ServiceRestIncomingTlsFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingTlsFailureReason() string {
	if o == nil || o.ServiceRestIncomingTlsFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceRestIncomingTlsFailureReason
}

// GetServiceRestIncomingTlsFailureReasonOk returns a tuple with the ServiceRestIncomingTlsFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingTlsFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceRestIncomingTlsFailureReason == nil {
		return nil, false
	}
	return o.ServiceRestIncomingTlsFailureReason, true
}

// HasServiceRestIncomingTlsFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingTlsFailureReason() bool {
	if o != nil && o.ServiceRestIncomingTlsFailureReason != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingTlsFailureReason gets a reference to the given string and assigns it to the ServiceRestIncomingTlsFailureReason field.
func (o *MsgVpn) SetServiceRestIncomingTlsFailureReason(v string) {
	o.ServiceRestIncomingTlsFailureReason = &v
}

// GetServiceRestIncomingTlsListenPort returns the ServiceRestIncomingTlsListenPort field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingTlsListenPort() int64 {
	if o == nil || o.ServiceRestIncomingTlsListenPort == nil {
		var ret int64
		return ret
	}
	return *o.ServiceRestIncomingTlsListenPort
}

// GetServiceRestIncomingTlsListenPortOk returns a tuple with the ServiceRestIncomingTlsListenPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingTlsListenPortOk() (*int64, bool) {
	if o == nil || o.ServiceRestIncomingTlsListenPort == nil {
		return nil, false
	}
	return o.ServiceRestIncomingTlsListenPort, true
}

// HasServiceRestIncomingTlsListenPort returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingTlsListenPort() bool {
	if o != nil && o.ServiceRestIncomingTlsListenPort != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingTlsListenPort gets a reference to the given int64 and assigns it to the ServiceRestIncomingTlsListenPort field.
func (o *MsgVpn) SetServiceRestIncomingTlsListenPort(v int64) {
	o.ServiceRestIncomingTlsListenPort = &v
}

// GetServiceRestIncomingTlsUp returns the ServiceRestIncomingTlsUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestIncomingTlsUp() bool {
	if o == nil || o.ServiceRestIncomingTlsUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceRestIncomingTlsUp
}

// GetServiceRestIncomingTlsUpOk returns a tuple with the ServiceRestIncomingTlsUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestIncomingTlsUpOk() (*bool, bool) {
	if o == nil || o.ServiceRestIncomingTlsUp == nil {
		return nil, false
	}
	return o.ServiceRestIncomingTlsUp, true
}

// HasServiceRestIncomingTlsUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestIncomingTlsUp() bool {
	if o != nil && o.ServiceRestIncomingTlsUp != nil {
		return true
	}

	return false
}

// SetServiceRestIncomingTlsUp gets a reference to the given bool and assigns it to the ServiceRestIncomingTlsUp field.
func (o *MsgVpn) SetServiceRestIncomingTlsUp(v bool) {
	o.ServiceRestIncomingTlsUp = &v
}

// GetServiceRestMode returns the ServiceRestMode field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestMode() string {
	if o == nil || o.ServiceRestMode == nil {
		var ret string
		return ret
	}
	return *o.ServiceRestMode
}

// GetServiceRestModeOk returns a tuple with the ServiceRestMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestModeOk() (*string, bool) {
	if o == nil || o.ServiceRestMode == nil {
		return nil, false
	}
	return o.ServiceRestMode, true
}

// HasServiceRestMode returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestMode() bool {
	if o != nil && o.ServiceRestMode != nil {
		return true
	}

	return false
}

// SetServiceRestMode gets a reference to the given string and assigns it to the ServiceRestMode field.
func (o *MsgVpn) SetServiceRestMode(v string) {
	o.ServiceRestMode = &v
}

// GetServiceRestOutgoingMaxConnectionCount returns the ServiceRestOutgoingMaxConnectionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceRestOutgoingMaxConnectionCount() int64 {
	if o == nil || o.ServiceRestOutgoingMaxConnectionCount == nil {
		var ret int64
		return ret
	}
	return *o.ServiceRestOutgoingMaxConnectionCount
}

// GetServiceRestOutgoingMaxConnectionCountOk returns a tuple with the ServiceRestOutgoingMaxConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceRestOutgoingMaxConnectionCountOk() (*int64, bool) {
	if o == nil || o.ServiceRestOutgoingMaxConnectionCount == nil {
		return nil, false
	}
	return o.ServiceRestOutgoingMaxConnectionCount, true
}

// HasServiceRestOutgoingMaxConnectionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceRestOutgoingMaxConnectionCount() bool {
	if o != nil && o.ServiceRestOutgoingMaxConnectionCount != nil {
		return true
	}

	return false
}

// SetServiceRestOutgoingMaxConnectionCount gets a reference to the given int64 and assigns it to the ServiceRestOutgoingMaxConnectionCount field.
func (o *MsgVpn) SetServiceRestOutgoingMaxConnectionCount(v int64) {
	o.ServiceRestOutgoingMaxConnectionCount = &v
}

// GetServiceSmfMaxConnectionCount returns the ServiceSmfMaxConnectionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceSmfMaxConnectionCount() int64 {
	if o == nil || o.ServiceSmfMaxConnectionCount == nil {
		var ret int64
		return ret
	}
	return *o.ServiceSmfMaxConnectionCount
}

// GetServiceSmfMaxConnectionCountOk returns a tuple with the ServiceSmfMaxConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceSmfMaxConnectionCountOk() (*int64, bool) {
	if o == nil || o.ServiceSmfMaxConnectionCount == nil {
		return nil, false
	}
	return o.ServiceSmfMaxConnectionCount, true
}

// HasServiceSmfMaxConnectionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceSmfMaxConnectionCount() bool {
	if o != nil && o.ServiceSmfMaxConnectionCount != nil {
		return true
	}

	return false
}

// SetServiceSmfMaxConnectionCount gets a reference to the given int64 and assigns it to the ServiceSmfMaxConnectionCount field.
func (o *MsgVpn) SetServiceSmfMaxConnectionCount(v int64) {
	o.ServiceSmfMaxConnectionCount = &v
}

// GetServiceSmfPlainTextEnabled returns the ServiceSmfPlainTextEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceSmfPlainTextEnabled() bool {
	if o == nil || o.ServiceSmfPlainTextEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceSmfPlainTextEnabled
}

// GetServiceSmfPlainTextEnabledOk returns a tuple with the ServiceSmfPlainTextEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceSmfPlainTextEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceSmfPlainTextEnabled == nil {
		return nil, false
	}
	return o.ServiceSmfPlainTextEnabled, true
}

// HasServiceSmfPlainTextEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceSmfPlainTextEnabled() bool {
	if o != nil && o.ServiceSmfPlainTextEnabled != nil {
		return true
	}

	return false
}

// SetServiceSmfPlainTextEnabled gets a reference to the given bool and assigns it to the ServiceSmfPlainTextEnabled field.
func (o *MsgVpn) SetServiceSmfPlainTextEnabled(v bool) {
	o.ServiceSmfPlainTextEnabled = &v
}

// GetServiceSmfPlainTextFailureReason returns the ServiceSmfPlainTextFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceSmfPlainTextFailureReason() string {
	if o == nil || o.ServiceSmfPlainTextFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceSmfPlainTextFailureReason
}

// GetServiceSmfPlainTextFailureReasonOk returns a tuple with the ServiceSmfPlainTextFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceSmfPlainTextFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceSmfPlainTextFailureReason == nil {
		return nil, false
	}
	return o.ServiceSmfPlainTextFailureReason, true
}

// HasServiceSmfPlainTextFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceSmfPlainTextFailureReason() bool {
	if o != nil && o.ServiceSmfPlainTextFailureReason != nil {
		return true
	}

	return false
}

// SetServiceSmfPlainTextFailureReason gets a reference to the given string and assigns it to the ServiceSmfPlainTextFailureReason field.
func (o *MsgVpn) SetServiceSmfPlainTextFailureReason(v string) {
	o.ServiceSmfPlainTextFailureReason = &v
}

// GetServiceSmfPlainTextUp returns the ServiceSmfPlainTextUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceSmfPlainTextUp() bool {
	if o == nil || o.ServiceSmfPlainTextUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceSmfPlainTextUp
}

// GetServiceSmfPlainTextUpOk returns a tuple with the ServiceSmfPlainTextUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceSmfPlainTextUpOk() (*bool, bool) {
	if o == nil || o.ServiceSmfPlainTextUp == nil {
		return nil, false
	}
	return o.ServiceSmfPlainTextUp, true
}

// HasServiceSmfPlainTextUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceSmfPlainTextUp() bool {
	if o != nil && o.ServiceSmfPlainTextUp != nil {
		return true
	}

	return false
}

// SetServiceSmfPlainTextUp gets a reference to the given bool and assigns it to the ServiceSmfPlainTextUp field.
func (o *MsgVpn) SetServiceSmfPlainTextUp(v bool) {
	o.ServiceSmfPlainTextUp = &v
}

// GetServiceSmfTlsEnabled returns the ServiceSmfTlsEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceSmfTlsEnabled() bool {
	if o == nil || o.ServiceSmfTlsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceSmfTlsEnabled
}

// GetServiceSmfTlsEnabledOk returns a tuple with the ServiceSmfTlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceSmfTlsEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceSmfTlsEnabled == nil {
		return nil, false
	}
	return o.ServiceSmfTlsEnabled, true
}

// HasServiceSmfTlsEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceSmfTlsEnabled() bool {
	if o != nil && o.ServiceSmfTlsEnabled != nil {
		return true
	}

	return false
}

// SetServiceSmfTlsEnabled gets a reference to the given bool and assigns it to the ServiceSmfTlsEnabled field.
func (o *MsgVpn) SetServiceSmfTlsEnabled(v bool) {
	o.ServiceSmfTlsEnabled = &v
}

// GetServiceSmfTlsFailureReason returns the ServiceSmfTlsFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceSmfTlsFailureReason() string {
	if o == nil || o.ServiceSmfTlsFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceSmfTlsFailureReason
}

// GetServiceSmfTlsFailureReasonOk returns a tuple with the ServiceSmfTlsFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceSmfTlsFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceSmfTlsFailureReason == nil {
		return nil, false
	}
	return o.ServiceSmfTlsFailureReason, true
}

// HasServiceSmfTlsFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceSmfTlsFailureReason() bool {
	if o != nil && o.ServiceSmfTlsFailureReason != nil {
		return true
	}

	return false
}

// SetServiceSmfTlsFailureReason gets a reference to the given string and assigns it to the ServiceSmfTlsFailureReason field.
func (o *MsgVpn) SetServiceSmfTlsFailureReason(v string) {
	o.ServiceSmfTlsFailureReason = &v
}

// GetServiceSmfTlsUp returns the ServiceSmfTlsUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceSmfTlsUp() bool {
	if o == nil || o.ServiceSmfTlsUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceSmfTlsUp
}

// GetServiceSmfTlsUpOk returns a tuple with the ServiceSmfTlsUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceSmfTlsUpOk() (*bool, bool) {
	if o == nil || o.ServiceSmfTlsUp == nil {
		return nil, false
	}
	return o.ServiceSmfTlsUp, true
}

// HasServiceSmfTlsUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceSmfTlsUp() bool {
	if o != nil && o.ServiceSmfTlsUp != nil {
		return true
	}

	return false
}

// SetServiceSmfTlsUp gets a reference to the given bool and assigns it to the ServiceSmfTlsUp field.
func (o *MsgVpn) SetServiceSmfTlsUp(v bool) {
	o.ServiceSmfTlsUp = &v
}

// GetServiceWebAuthenticationClientCertRequest returns the ServiceWebAuthenticationClientCertRequest field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceWebAuthenticationClientCertRequest() string {
	if o == nil || o.ServiceWebAuthenticationClientCertRequest == nil {
		var ret string
		return ret
	}
	return *o.ServiceWebAuthenticationClientCertRequest
}

// GetServiceWebAuthenticationClientCertRequestOk returns a tuple with the ServiceWebAuthenticationClientCertRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceWebAuthenticationClientCertRequestOk() (*string, bool) {
	if o == nil || o.ServiceWebAuthenticationClientCertRequest == nil {
		return nil, false
	}
	return o.ServiceWebAuthenticationClientCertRequest, true
}

// HasServiceWebAuthenticationClientCertRequest returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceWebAuthenticationClientCertRequest() bool {
	if o != nil && o.ServiceWebAuthenticationClientCertRequest != nil {
		return true
	}

	return false
}

// SetServiceWebAuthenticationClientCertRequest gets a reference to the given string and assigns it to the ServiceWebAuthenticationClientCertRequest field.
func (o *MsgVpn) SetServiceWebAuthenticationClientCertRequest(v string) {
	o.ServiceWebAuthenticationClientCertRequest = &v
}

// GetServiceWebMaxConnectionCount returns the ServiceWebMaxConnectionCount field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceWebMaxConnectionCount() int64 {
	if o == nil || o.ServiceWebMaxConnectionCount == nil {
		var ret int64
		return ret
	}
	return *o.ServiceWebMaxConnectionCount
}

// GetServiceWebMaxConnectionCountOk returns a tuple with the ServiceWebMaxConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceWebMaxConnectionCountOk() (*int64, bool) {
	if o == nil || o.ServiceWebMaxConnectionCount == nil {
		return nil, false
	}
	return o.ServiceWebMaxConnectionCount, true
}

// HasServiceWebMaxConnectionCount returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceWebMaxConnectionCount() bool {
	if o != nil && o.ServiceWebMaxConnectionCount != nil {
		return true
	}

	return false
}

// SetServiceWebMaxConnectionCount gets a reference to the given int64 and assigns it to the ServiceWebMaxConnectionCount field.
func (o *MsgVpn) SetServiceWebMaxConnectionCount(v int64) {
	o.ServiceWebMaxConnectionCount = &v
}

// GetServiceWebPlainTextEnabled returns the ServiceWebPlainTextEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceWebPlainTextEnabled() bool {
	if o == nil || o.ServiceWebPlainTextEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceWebPlainTextEnabled
}

// GetServiceWebPlainTextEnabledOk returns a tuple with the ServiceWebPlainTextEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceWebPlainTextEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceWebPlainTextEnabled == nil {
		return nil, false
	}
	return o.ServiceWebPlainTextEnabled, true
}

// HasServiceWebPlainTextEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceWebPlainTextEnabled() bool {
	if o != nil && o.ServiceWebPlainTextEnabled != nil {
		return true
	}

	return false
}

// SetServiceWebPlainTextEnabled gets a reference to the given bool and assigns it to the ServiceWebPlainTextEnabled field.
func (o *MsgVpn) SetServiceWebPlainTextEnabled(v bool) {
	o.ServiceWebPlainTextEnabled = &v
}

// GetServiceWebPlainTextFailureReason returns the ServiceWebPlainTextFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceWebPlainTextFailureReason() string {
	if o == nil || o.ServiceWebPlainTextFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceWebPlainTextFailureReason
}

// GetServiceWebPlainTextFailureReasonOk returns a tuple with the ServiceWebPlainTextFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceWebPlainTextFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceWebPlainTextFailureReason == nil {
		return nil, false
	}
	return o.ServiceWebPlainTextFailureReason, true
}

// HasServiceWebPlainTextFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceWebPlainTextFailureReason() bool {
	if o != nil && o.ServiceWebPlainTextFailureReason != nil {
		return true
	}

	return false
}

// SetServiceWebPlainTextFailureReason gets a reference to the given string and assigns it to the ServiceWebPlainTextFailureReason field.
func (o *MsgVpn) SetServiceWebPlainTextFailureReason(v string) {
	o.ServiceWebPlainTextFailureReason = &v
}

// GetServiceWebPlainTextUp returns the ServiceWebPlainTextUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceWebPlainTextUp() bool {
	if o == nil || o.ServiceWebPlainTextUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceWebPlainTextUp
}

// GetServiceWebPlainTextUpOk returns a tuple with the ServiceWebPlainTextUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceWebPlainTextUpOk() (*bool, bool) {
	if o == nil || o.ServiceWebPlainTextUp == nil {
		return nil, false
	}
	return o.ServiceWebPlainTextUp, true
}

// HasServiceWebPlainTextUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceWebPlainTextUp() bool {
	if o != nil && o.ServiceWebPlainTextUp != nil {
		return true
	}

	return false
}

// SetServiceWebPlainTextUp gets a reference to the given bool and assigns it to the ServiceWebPlainTextUp field.
func (o *MsgVpn) SetServiceWebPlainTextUp(v bool) {
	o.ServiceWebPlainTextUp = &v
}

// GetServiceWebTlsEnabled returns the ServiceWebTlsEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceWebTlsEnabled() bool {
	if o == nil || o.ServiceWebTlsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServiceWebTlsEnabled
}

// GetServiceWebTlsEnabledOk returns a tuple with the ServiceWebTlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceWebTlsEnabledOk() (*bool, bool) {
	if o == nil || o.ServiceWebTlsEnabled == nil {
		return nil, false
	}
	return o.ServiceWebTlsEnabled, true
}

// HasServiceWebTlsEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceWebTlsEnabled() bool {
	if o != nil && o.ServiceWebTlsEnabled != nil {
		return true
	}

	return false
}

// SetServiceWebTlsEnabled gets a reference to the given bool and assigns it to the ServiceWebTlsEnabled field.
func (o *MsgVpn) SetServiceWebTlsEnabled(v bool) {
	o.ServiceWebTlsEnabled = &v
}

// GetServiceWebTlsFailureReason returns the ServiceWebTlsFailureReason field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceWebTlsFailureReason() string {
	if o == nil || o.ServiceWebTlsFailureReason == nil {
		var ret string
		return ret
	}
	return *o.ServiceWebTlsFailureReason
}

// GetServiceWebTlsFailureReasonOk returns a tuple with the ServiceWebTlsFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceWebTlsFailureReasonOk() (*string, bool) {
	if o == nil || o.ServiceWebTlsFailureReason == nil {
		return nil, false
	}
	return o.ServiceWebTlsFailureReason, true
}

// HasServiceWebTlsFailureReason returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceWebTlsFailureReason() bool {
	if o != nil && o.ServiceWebTlsFailureReason != nil {
		return true
	}

	return false
}

// SetServiceWebTlsFailureReason gets a reference to the given string and assigns it to the ServiceWebTlsFailureReason field.
func (o *MsgVpn) SetServiceWebTlsFailureReason(v string) {
	o.ServiceWebTlsFailureReason = &v
}

// GetServiceWebTlsUp returns the ServiceWebTlsUp field value if set, zero value otherwise.
func (o *MsgVpn) GetServiceWebTlsUp() bool {
	if o == nil || o.ServiceWebTlsUp == nil {
		var ret bool
		return ret
	}
	return *o.ServiceWebTlsUp
}

// GetServiceWebTlsUpOk returns a tuple with the ServiceWebTlsUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetServiceWebTlsUpOk() (*bool, bool) {
	if o == nil || o.ServiceWebTlsUp == nil {
		return nil, false
	}
	return o.ServiceWebTlsUp, true
}

// HasServiceWebTlsUp returns a boolean if a field has been set.
func (o *MsgVpn) HasServiceWebTlsUp() bool {
	if o != nil && o.ServiceWebTlsUp != nil {
		return true
	}

	return false
}

// SetServiceWebTlsUp gets a reference to the given bool and assigns it to the ServiceWebTlsUp field.
func (o *MsgVpn) SetServiceWebTlsUp(v bool) {
	o.ServiceWebTlsUp = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MsgVpn) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MsgVpn) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MsgVpn) SetState(v string) {
	o.State = &v
}

// GetSubscriptionExportProgress returns the SubscriptionExportProgress field value if set, zero value otherwise.
func (o *MsgVpn) GetSubscriptionExportProgress() int64 {
	if o == nil || o.SubscriptionExportProgress == nil {
		var ret int64
		return ret
	}
	return *o.SubscriptionExportProgress
}

// GetSubscriptionExportProgressOk returns a tuple with the SubscriptionExportProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetSubscriptionExportProgressOk() (*int64, bool) {
	if o == nil || o.SubscriptionExportProgress == nil {
		return nil, false
	}
	return o.SubscriptionExportProgress, true
}

// HasSubscriptionExportProgress returns a boolean if a field has been set.
func (o *MsgVpn) HasSubscriptionExportProgress() bool {
	if o != nil && o.SubscriptionExportProgress != nil {
		return true
	}

	return false
}

// SetSubscriptionExportProgress gets a reference to the given int64 and assigns it to the SubscriptionExportProgress field.
func (o *MsgVpn) SetSubscriptionExportProgress(v int64) {
	o.SubscriptionExportProgress = &v
}

// GetSystemManager returns the SystemManager field value if set, zero value otherwise.
func (o *MsgVpn) GetSystemManager() bool {
	if o == nil || o.SystemManager == nil {
		var ret bool
		return ret
	}
	return *o.SystemManager
}

// GetSystemManagerOk returns a tuple with the SystemManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetSystemManagerOk() (*bool, bool) {
	if o == nil || o.SystemManager == nil {
		return nil, false
	}
	return o.SystemManager, true
}

// HasSystemManager returns a boolean if a field has been set.
func (o *MsgVpn) HasSystemManager() bool {
	if o != nil && o.SystemManager != nil {
		return true
	}

	return false
}

// SetSystemManager gets a reference to the given bool and assigns it to the SystemManager field.
func (o *MsgVpn) SetSystemManager(v bool) {
	o.SystemManager = &v
}

// GetTlsAllowDowngradeToPlainTextEnabled returns the TlsAllowDowngradeToPlainTextEnabled field value if set, zero value otherwise.
func (o *MsgVpn) GetTlsAllowDowngradeToPlainTextEnabled() bool {
	if o == nil || o.TlsAllowDowngradeToPlainTextEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TlsAllowDowngradeToPlainTextEnabled
}

// GetTlsAllowDowngradeToPlainTextEnabledOk returns a tuple with the TlsAllowDowngradeToPlainTextEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTlsAllowDowngradeToPlainTextEnabledOk() (*bool, bool) {
	if o == nil || o.TlsAllowDowngradeToPlainTextEnabled == nil {
		return nil, false
	}
	return o.TlsAllowDowngradeToPlainTextEnabled, true
}

// HasTlsAllowDowngradeToPlainTextEnabled returns a boolean if a field has been set.
func (o *MsgVpn) HasTlsAllowDowngradeToPlainTextEnabled() bool {
	if o != nil && o.TlsAllowDowngradeToPlainTextEnabled != nil {
		return true
	}

	return false
}

// SetTlsAllowDowngradeToPlainTextEnabled gets a reference to the given bool and assigns it to the TlsAllowDowngradeToPlainTextEnabled field.
func (o *MsgVpn) SetTlsAllowDowngradeToPlainTextEnabled(v bool) {
	o.TlsAllowDowngradeToPlainTextEnabled = &v
}

// GetTlsAverageRxByteRate returns the TlsAverageRxByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetTlsAverageRxByteRate() int64 {
	if o == nil || o.TlsAverageRxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.TlsAverageRxByteRate
}

// GetTlsAverageRxByteRateOk returns a tuple with the TlsAverageRxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTlsAverageRxByteRateOk() (*int64, bool) {
	if o == nil || o.TlsAverageRxByteRate == nil {
		return nil, false
	}
	return o.TlsAverageRxByteRate, true
}

// HasTlsAverageRxByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasTlsAverageRxByteRate() bool {
	if o != nil && o.TlsAverageRxByteRate != nil {
		return true
	}

	return false
}

// SetTlsAverageRxByteRate gets a reference to the given int64 and assigns it to the TlsAverageRxByteRate field.
func (o *MsgVpn) SetTlsAverageRxByteRate(v int64) {
	o.TlsAverageRxByteRate = &v
}

// GetTlsAverageTxByteRate returns the TlsAverageTxByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetTlsAverageTxByteRate() int64 {
	if o == nil || o.TlsAverageTxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.TlsAverageTxByteRate
}

// GetTlsAverageTxByteRateOk returns a tuple with the TlsAverageTxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTlsAverageTxByteRateOk() (*int64, bool) {
	if o == nil || o.TlsAverageTxByteRate == nil {
		return nil, false
	}
	return o.TlsAverageTxByteRate, true
}

// HasTlsAverageTxByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasTlsAverageTxByteRate() bool {
	if o != nil && o.TlsAverageTxByteRate != nil {
		return true
	}

	return false
}

// SetTlsAverageTxByteRate gets a reference to the given int64 and assigns it to the TlsAverageTxByteRate field.
func (o *MsgVpn) SetTlsAverageTxByteRate(v int64) {
	o.TlsAverageTxByteRate = &v
}

// GetTlsRxByteCount returns the TlsRxByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetTlsRxByteCount() int64 {
	if o == nil || o.TlsRxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.TlsRxByteCount
}

// GetTlsRxByteCountOk returns a tuple with the TlsRxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTlsRxByteCountOk() (*int64, bool) {
	if o == nil || o.TlsRxByteCount == nil {
		return nil, false
	}
	return o.TlsRxByteCount, true
}

// HasTlsRxByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasTlsRxByteCount() bool {
	if o != nil && o.TlsRxByteCount != nil {
		return true
	}

	return false
}

// SetTlsRxByteCount gets a reference to the given int64 and assigns it to the TlsRxByteCount field.
func (o *MsgVpn) SetTlsRxByteCount(v int64) {
	o.TlsRxByteCount = &v
}

// GetTlsRxByteRate returns the TlsRxByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetTlsRxByteRate() int64 {
	if o == nil || o.TlsRxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.TlsRxByteRate
}

// GetTlsRxByteRateOk returns a tuple with the TlsRxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTlsRxByteRateOk() (*int64, bool) {
	if o == nil || o.TlsRxByteRate == nil {
		return nil, false
	}
	return o.TlsRxByteRate, true
}

// HasTlsRxByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasTlsRxByteRate() bool {
	if o != nil && o.TlsRxByteRate != nil {
		return true
	}

	return false
}

// SetTlsRxByteRate gets a reference to the given int64 and assigns it to the TlsRxByteRate field.
func (o *MsgVpn) SetTlsRxByteRate(v int64) {
	o.TlsRxByteRate = &v
}

// GetTlsTxByteCount returns the TlsTxByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetTlsTxByteCount() int64 {
	if o == nil || o.TlsTxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.TlsTxByteCount
}

// GetTlsTxByteCountOk returns a tuple with the TlsTxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTlsTxByteCountOk() (*int64, bool) {
	if o == nil || o.TlsTxByteCount == nil {
		return nil, false
	}
	return o.TlsTxByteCount, true
}

// HasTlsTxByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasTlsTxByteCount() bool {
	if o != nil && o.TlsTxByteCount != nil {
		return true
	}

	return false
}

// SetTlsTxByteCount gets a reference to the given int64 and assigns it to the TlsTxByteCount field.
func (o *MsgVpn) SetTlsTxByteCount(v int64) {
	o.TlsTxByteCount = &v
}

// GetTlsTxByteRate returns the TlsTxByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetTlsTxByteRate() int64 {
	if o == nil || o.TlsTxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.TlsTxByteRate
}

// GetTlsTxByteRateOk returns a tuple with the TlsTxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTlsTxByteRateOk() (*int64, bool) {
	if o == nil || o.TlsTxByteRate == nil {
		return nil, false
	}
	return o.TlsTxByteRate, true
}

// HasTlsTxByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasTlsTxByteRate() bool {
	if o != nil && o.TlsTxByteRate != nil {
		return true
	}

	return false
}

// SetTlsTxByteRate gets a reference to the given int64 and assigns it to the TlsTxByteRate field.
func (o *MsgVpn) SetTlsTxByteRate(v int64) {
	o.TlsTxByteRate = &v
}

// GetTxByteCount returns the TxByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetTxByteCount() int64 {
	if o == nil || o.TxByteCount == nil {
		var ret int64
		return ret
	}
	return *o.TxByteCount
}

// GetTxByteCountOk returns a tuple with the TxByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTxByteCountOk() (*int64, bool) {
	if o == nil || o.TxByteCount == nil {
		return nil, false
	}
	return o.TxByteCount, true
}

// HasTxByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasTxByteCount() bool {
	if o != nil && o.TxByteCount != nil {
		return true
	}

	return false
}

// SetTxByteCount gets a reference to the given int64 and assigns it to the TxByteCount field.
func (o *MsgVpn) SetTxByteCount(v int64) {
	o.TxByteCount = &v
}

// GetTxByteRate returns the TxByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetTxByteRate() int64 {
	if o == nil || o.TxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.TxByteRate
}

// GetTxByteRateOk returns a tuple with the TxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTxByteRateOk() (*int64, bool) {
	if o == nil || o.TxByteRate == nil {
		return nil, false
	}
	return o.TxByteRate, true
}

// HasTxByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasTxByteRate() bool {
	if o != nil && o.TxByteRate != nil {
		return true
	}

	return false
}

// SetTxByteRate gets a reference to the given int64 and assigns it to the TxByteRate field.
func (o *MsgVpn) SetTxByteRate(v int64) {
	o.TxByteRate = &v
}

// GetTxCompressedByteCount returns the TxCompressedByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetTxCompressedByteCount() int64 {
	if o == nil || o.TxCompressedByteCount == nil {
		var ret int64
		return ret
	}
	return *o.TxCompressedByteCount
}

// GetTxCompressedByteCountOk returns a tuple with the TxCompressedByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTxCompressedByteCountOk() (*int64, bool) {
	if o == nil || o.TxCompressedByteCount == nil {
		return nil, false
	}
	return o.TxCompressedByteCount, true
}

// HasTxCompressedByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasTxCompressedByteCount() bool {
	if o != nil && o.TxCompressedByteCount != nil {
		return true
	}

	return false
}

// SetTxCompressedByteCount gets a reference to the given int64 and assigns it to the TxCompressedByteCount field.
func (o *MsgVpn) SetTxCompressedByteCount(v int64) {
	o.TxCompressedByteCount = &v
}

// GetTxCompressedByteRate returns the TxCompressedByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetTxCompressedByteRate() int64 {
	if o == nil || o.TxCompressedByteRate == nil {
		var ret int64
		return ret
	}
	return *o.TxCompressedByteRate
}

// GetTxCompressedByteRateOk returns a tuple with the TxCompressedByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTxCompressedByteRateOk() (*int64, bool) {
	if o == nil || o.TxCompressedByteRate == nil {
		return nil, false
	}
	return o.TxCompressedByteRate, true
}

// HasTxCompressedByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasTxCompressedByteRate() bool {
	if o != nil && o.TxCompressedByteRate != nil {
		return true
	}

	return false
}

// SetTxCompressedByteRate gets a reference to the given int64 and assigns it to the TxCompressedByteRate field.
func (o *MsgVpn) SetTxCompressedByteRate(v int64) {
	o.TxCompressedByteRate = &v
}

// GetTxCompressionRatio returns the TxCompressionRatio field value if set, zero value otherwise.
func (o *MsgVpn) GetTxCompressionRatio() string {
	if o == nil || o.TxCompressionRatio == nil {
		var ret string
		return ret
	}
	return *o.TxCompressionRatio
}

// GetTxCompressionRatioOk returns a tuple with the TxCompressionRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTxCompressionRatioOk() (*string, bool) {
	if o == nil || o.TxCompressionRatio == nil {
		return nil, false
	}
	return o.TxCompressionRatio, true
}

// HasTxCompressionRatio returns a boolean if a field has been set.
func (o *MsgVpn) HasTxCompressionRatio() bool {
	if o != nil && o.TxCompressionRatio != nil {
		return true
	}

	return false
}

// SetTxCompressionRatio gets a reference to the given string and assigns it to the TxCompressionRatio field.
func (o *MsgVpn) SetTxCompressionRatio(v string) {
	o.TxCompressionRatio = &v
}

// GetTxMsgCount returns the TxMsgCount field value if set, zero value otherwise.
func (o *MsgVpn) GetTxMsgCount() int64 {
	if o == nil || o.TxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.TxMsgCount
}

// GetTxMsgCountOk returns a tuple with the TxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTxMsgCountOk() (*int64, bool) {
	if o == nil || o.TxMsgCount == nil {
		return nil, false
	}
	return o.TxMsgCount, true
}

// HasTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpn) HasTxMsgCount() bool {
	if o != nil && o.TxMsgCount != nil {
		return true
	}

	return false
}

// SetTxMsgCount gets a reference to the given int64 and assigns it to the TxMsgCount field.
func (o *MsgVpn) SetTxMsgCount(v int64) {
	o.TxMsgCount = &v
}

// GetTxMsgRate returns the TxMsgRate field value if set, zero value otherwise.
func (o *MsgVpn) GetTxMsgRate() int64 {
	if o == nil || o.TxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.TxMsgRate
}

// GetTxMsgRateOk returns a tuple with the TxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTxMsgRateOk() (*int64, bool) {
	if o == nil || o.TxMsgRate == nil {
		return nil, false
	}
	return o.TxMsgRate, true
}

// HasTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpn) HasTxMsgRate() bool {
	if o != nil && o.TxMsgRate != nil {
		return true
	}

	return false
}

// SetTxMsgRate gets a reference to the given int64 and assigns it to the TxMsgRate field.
func (o *MsgVpn) SetTxMsgRate(v int64) {
	o.TxMsgRate = &v
}

// GetTxUncompressedByteCount returns the TxUncompressedByteCount field value if set, zero value otherwise.
func (o *MsgVpn) GetTxUncompressedByteCount() int64 {
	if o == nil || o.TxUncompressedByteCount == nil {
		var ret int64
		return ret
	}
	return *o.TxUncompressedByteCount
}

// GetTxUncompressedByteCountOk returns a tuple with the TxUncompressedByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTxUncompressedByteCountOk() (*int64, bool) {
	if o == nil || o.TxUncompressedByteCount == nil {
		return nil, false
	}
	return o.TxUncompressedByteCount, true
}

// HasTxUncompressedByteCount returns a boolean if a field has been set.
func (o *MsgVpn) HasTxUncompressedByteCount() bool {
	if o != nil && o.TxUncompressedByteCount != nil {
		return true
	}

	return false
}

// SetTxUncompressedByteCount gets a reference to the given int64 and assigns it to the TxUncompressedByteCount field.
func (o *MsgVpn) SetTxUncompressedByteCount(v int64) {
	o.TxUncompressedByteCount = &v
}

// GetTxUncompressedByteRate returns the TxUncompressedByteRate field value if set, zero value otherwise.
func (o *MsgVpn) GetTxUncompressedByteRate() int64 {
	if o == nil || o.TxUncompressedByteRate == nil {
		var ret int64
		return ret
	}
	return *o.TxUncompressedByteRate
}

// GetTxUncompressedByteRateOk returns a tuple with the TxUncompressedByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpn) GetTxUncompressedByteRateOk() (*int64, bool) {
	if o == nil || o.TxUncompressedByteRate == nil {
		return nil, false
	}
	return o.TxUncompressedByteRate, true
}

// HasTxUncompressedByteRate returns a boolean if a field has been set.
func (o *MsgVpn) HasTxUncompressedByteRate() bool {
	if o != nil && o.TxUncompressedByteRate != nil {
		return true
	}

	return false
}

// SetTxUncompressedByteRate gets a reference to the given int64 and assigns it to the TxUncompressedByteRate field.
func (o *MsgVpn) SetTxUncompressedByteRate(v int64) {
	o.TxUncompressedByteRate = &v
}

func (o MsgVpn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.AuthenticationBasicEnabled != nil {
		toSerialize["authenticationBasicEnabled"] = o.AuthenticationBasicEnabled
	}
	if o.AuthenticationBasicProfileName != nil {
		toSerialize["authenticationBasicProfileName"] = o.AuthenticationBasicProfileName
	}
	if o.AuthenticationBasicRadiusDomain != nil {
		toSerialize["authenticationBasicRadiusDomain"] = o.AuthenticationBasicRadiusDomain
	}
	if o.AuthenticationBasicType != nil {
		toSerialize["authenticationBasicType"] = o.AuthenticationBasicType
	}
	if o.AuthenticationClientCertAllowApiProvidedUsernameEnabled != nil {
		toSerialize["authenticationClientCertAllowApiProvidedUsernameEnabled"] = o.AuthenticationClientCertAllowApiProvidedUsernameEnabled
	}
	if o.AuthenticationClientCertEnabled != nil {
		toSerialize["authenticationClientCertEnabled"] = o.AuthenticationClientCertEnabled
	}
	if o.AuthenticationClientCertMaxChainDepth != nil {
		toSerialize["authenticationClientCertMaxChainDepth"] = o.AuthenticationClientCertMaxChainDepth
	}
	if o.AuthenticationClientCertRevocationCheckMode != nil {
		toSerialize["authenticationClientCertRevocationCheckMode"] = o.AuthenticationClientCertRevocationCheckMode
	}
	if o.AuthenticationClientCertUsernameSource != nil {
		toSerialize["authenticationClientCertUsernameSource"] = o.AuthenticationClientCertUsernameSource
	}
	if o.AuthenticationClientCertValidateDateEnabled != nil {
		toSerialize["authenticationClientCertValidateDateEnabled"] = o.AuthenticationClientCertValidateDateEnabled
	}
	if o.AuthenticationKerberosAllowApiProvidedUsernameEnabled != nil {
		toSerialize["authenticationKerberosAllowApiProvidedUsernameEnabled"] = o.AuthenticationKerberosAllowApiProvidedUsernameEnabled
	}
	if o.AuthenticationKerberosEnabled != nil {
		toSerialize["authenticationKerberosEnabled"] = o.AuthenticationKerberosEnabled
	}
	if o.AuthenticationOauthDefaultProviderName != nil {
		toSerialize["authenticationOauthDefaultProviderName"] = o.AuthenticationOauthDefaultProviderName
	}
	if o.AuthenticationOauthEnabled != nil {
		toSerialize["authenticationOauthEnabled"] = o.AuthenticationOauthEnabled
	}
	if o.AuthorizationLdapGroupMembershipAttributeName != nil {
		toSerialize["authorizationLdapGroupMembershipAttributeName"] = o.AuthorizationLdapGroupMembershipAttributeName
	}
	if o.AuthorizationLdapTrimClientUsernameDomainEnabled != nil {
		toSerialize["authorizationLdapTrimClientUsernameDomainEnabled"] = o.AuthorizationLdapTrimClientUsernameDomainEnabled
	}
	if o.AuthorizationProfileName != nil {
		toSerialize["authorizationProfileName"] = o.AuthorizationProfileName
	}
	if o.AuthorizationType != nil {
		toSerialize["authorizationType"] = o.AuthorizationType
	}
	if o.AverageRxByteRate != nil {
		toSerialize["averageRxByteRate"] = o.AverageRxByteRate
	}
	if o.AverageRxCompressedByteRate != nil {
		toSerialize["averageRxCompressedByteRate"] = o.AverageRxCompressedByteRate
	}
	if o.AverageRxMsgRate != nil {
		toSerialize["averageRxMsgRate"] = o.AverageRxMsgRate
	}
	if o.AverageRxUncompressedByteRate != nil {
		toSerialize["averageRxUncompressedByteRate"] = o.AverageRxUncompressedByteRate
	}
	if o.AverageTxByteRate != nil {
		toSerialize["averageTxByteRate"] = o.AverageTxByteRate
	}
	if o.AverageTxCompressedByteRate != nil {
		toSerialize["averageTxCompressedByteRate"] = o.AverageTxCompressedByteRate
	}
	if o.AverageTxMsgRate != nil {
		toSerialize["averageTxMsgRate"] = o.AverageTxMsgRate
	}
	if o.AverageTxUncompressedByteRate != nil {
		toSerialize["averageTxUncompressedByteRate"] = o.AverageTxUncompressedByteRate
	}
	if o.BridgingTlsServerCertEnforceTrustedCommonNameEnabled != nil {
		toSerialize["bridgingTlsServerCertEnforceTrustedCommonNameEnabled"] = o.BridgingTlsServerCertEnforceTrustedCommonNameEnabled
	}
	if o.BridgingTlsServerCertMaxChainDepth != nil {
		toSerialize["bridgingTlsServerCertMaxChainDepth"] = o.BridgingTlsServerCertMaxChainDepth
	}
	if o.BridgingTlsServerCertValidateDateEnabled != nil {
		toSerialize["bridgingTlsServerCertValidateDateEnabled"] = o.BridgingTlsServerCertValidateDateEnabled
	}
	if o.BridgingTlsServerCertValidateNameEnabled != nil {
		toSerialize["bridgingTlsServerCertValidateNameEnabled"] = o.BridgingTlsServerCertValidateNameEnabled
	}
	if o.ConfigSyncLocalKey != nil {
		toSerialize["configSyncLocalKey"] = o.ConfigSyncLocalKey
	}
	if o.ConfigSyncLocalLastResult != nil {
		toSerialize["configSyncLocalLastResult"] = o.ConfigSyncLocalLastResult
	}
	if o.ConfigSyncLocalRole != nil {
		toSerialize["configSyncLocalRole"] = o.ConfigSyncLocalRole
	}
	if o.ConfigSyncLocalState != nil {
		toSerialize["configSyncLocalState"] = o.ConfigSyncLocalState
	}
	if o.ConfigSyncLocalTimeInState != nil {
		toSerialize["configSyncLocalTimeInState"] = o.ConfigSyncLocalTimeInState
	}
	if o.ControlRxByteCount != nil {
		toSerialize["controlRxByteCount"] = o.ControlRxByteCount
	}
	if o.ControlRxMsgCount != nil {
		toSerialize["controlRxMsgCount"] = o.ControlRxMsgCount
	}
	if o.ControlTxByteCount != nil {
		toSerialize["controlTxByteCount"] = o.ControlTxByteCount
	}
	if o.ControlTxMsgCount != nil {
		toSerialize["controlTxMsgCount"] = o.ControlTxMsgCount
	}
	if o.Counter != nil {
		toSerialize["counter"] = o.Counter
	}
	if o.DataRxByteCount != nil {
		toSerialize["dataRxByteCount"] = o.DataRxByteCount
	}
	if o.DataRxMsgCount != nil {
		toSerialize["dataRxMsgCount"] = o.DataRxMsgCount
	}
	if o.DataTxByteCount != nil {
		toSerialize["dataTxByteCount"] = o.DataTxByteCount
	}
	if o.DataTxMsgCount != nil {
		toSerialize["dataTxMsgCount"] = o.DataTxMsgCount
	}
	if o.DiscardedRxMsgCount != nil {
		toSerialize["discardedRxMsgCount"] = o.DiscardedRxMsgCount
	}
	if o.DiscardedTxMsgCount != nil {
		toSerialize["discardedTxMsgCount"] = o.DiscardedTxMsgCount
	}
	if o.DistributedCacheManagementEnabled != nil {
		toSerialize["distributedCacheManagementEnabled"] = o.DistributedCacheManagementEnabled
	}
	if o.DmrEnabled != nil {
		toSerialize["dmrEnabled"] = o.DmrEnabled
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.EventConnectionCountThreshold != nil {
		toSerialize["eventConnectionCountThreshold"] = o.EventConnectionCountThreshold
	}
	if o.EventEgressFlowCountThreshold != nil {
		toSerialize["eventEgressFlowCountThreshold"] = o.EventEgressFlowCountThreshold
	}
	if o.EventEgressMsgRateThreshold != nil {
		toSerialize["eventEgressMsgRateThreshold"] = o.EventEgressMsgRateThreshold
	}
	if o.EventEndpointCountThreshold != nil {
		toSerialize["eventEndpointCountThreshold"] = o.EventEndpointCountThreshold
	}
	if o.EventIngressFlowCountThreshold != nil {
		toSerialize["eventIngressFlowCountThreshold"] = o.EventIngressFlowCountThreshold
	}
	if o.EventIngressMsgRateThreshold != nil {
		toSerialize["eventIngressMsgRateThreshold"] = o.EventIngressMsgRateThreshold
	}
	if o.EventLargeMsgThreshold != nil {
		toSerialize["eventLargeMsgThreshold"] = o.EventLargeMsgThreshold
	}
	if o.EventLogTag != nil {
		toSerialize["eventLogTag"] = o.EventLogTag
	}
	if o.EventMsgSpoolUsageThreshold != nil {
		toSerialize["eventMsgSpoolUsageThreshold"] = o.EventMsgSpoolUsageThreshold
	}
	if o.EventPublishClientEnabled != nil {
		toSerialize["eventPublishClientEnabled"] = o.EventPublishClientEnabled
	}
	if o.EventPublishMsgVpnEnabled != nil {
		toSerialize["eventPublishMsgVpnEnabled"] = o.EventPublishMsgVpnEnabled
	}
	if o.EventPublishSubscriptionMode != nil {
		toSerialize["eventPublishSubscriptionMode"] = o.EventPublishSubscriptionMode
	}
	if o.EventPublishTopicFormatMqttEnabled != nil {
		toSerialize["eventPublishTopicFormatMqttEnabled"] = o.EventPublishTopicFormatMqttEnabled
	}
	if o.EventPublishTopicFormatSmfEnabled != nil {
		toSerialize["eventPublishTopicFormatSmfEnabled"] = o.EventPublishTopicFormatSmfEnabled
	}
	if o.EventServiceAmqpConnectionCountThreshold != nil {
		toSerialize["eventServiceAmqpConnectionCountThreshold"] = o.EventServiceAmqpConnectionCountThreshold
	}
	if o.EventServiceMqttConnectionCountThreshold != nil {
		toSerialize["eventServiceMqttConnectionCountThreshold"] = o.EventServiceMqttConnectionCountThreshold
	}
	if o.EventServiceRestIncomingConnectionCountThreshold != nil {
		toSerialize["eventServiceRestIncomingConnectionCountThreshold"] = o.EventServiceRestIncomingConnectionCountThreshold
	}
	if o.EventServiceSmfConnectionCountThreshold != nil {
		toSerialize["eventServiceSmfConnectionCountThreshold"] = o.EventServiceSmfConnectionCountThreshold
	}
	if o.EventServiceWebConnectionCountThreshold != nil {
		toSerialize["eventServiceWebConnectionCountThreshold"] = o.EventServiceWebConnectionCountThreshold
	}
	if o.EventSubscriptionCountThreshold != nil {
		toSerialize["eventSubscriptionCountThreshold"] = o.EventSubscriptionCountThreshold
	}
	if o.EventTransactedSessionCountThreshold != nil {
		toSerialize["eventTransactedSessionCountThreshold"] = o.EventTransactedSessionCountThreshold
	}
	if o.EventTransactionCountThreshold != nil {
		toSerialize["eventTransactionCountThreshold"] = o.EventTransactionCountThreshold
	}
	if o.ExportSubscriptionsEnabled != nil {
		toSerialize["exportSubscriptionsEnabled"] = o.ExportSubscriptionsEnabled
	}
	if o.FailureReason != nil {
		toSerialize["failureReason"] = o.FailureReason
	}
	if o.JndiEnabled != nil {
		toSerialize["jndiEnabled"] = o.JndiEnabled
	}
	if o.LoginRxMsgCount != nil {
		toSerialize["loginRxMsgCount"] = o.LoginRxMsgCount
	}
	if o.LoginTxMsgCount != nil {
		toSerialize["loginTxMsgCount"] = o.LoginTxMsgCount
	}
	if o.MaxConnectionCount != nil {
		toSerialize["maxConnectionCount"] = o.MaxConnectionCount
	}
	if o.MaxEffectiveEndpointCount != nil {
		toSerialize["maxEffectiveEndpointCount"] = o.MaxEffectiveEndpointCount
	}
	if o.MaxEffectiveRxFlowCount != nil {
		toSerialize["maxEffectiveRxFlowCount"] = o.MaxEffectiveRxFlowCount
	}
	if o.MaxEffectiveSubscriptionCount != nil {
		toSerialize["maxEffectiveSubscriptionCount"] = o.MaxEffectiveSubscriptionCount
	}
	if o.MaxEffectiveTransactedSessionCount != nil {
		toSerialize["maxEffectiveTransactedSessionCount"] = o.MaxEffectiveTransactedSessionCount
	}
	if o.MaxEffectiveTransactionCount != nil {
		toSerialize["maxEffectiveTransactionCount"] = o.MaxEffectiveTransactionCount
	}
	if o.MaxEffectiveTxFlowCount != nil {
		toSerialize["maxEffectiveTxFlowCount"] = o.MaxEffectiveTxFlowCount
	}
	if o.MaxEgressFlowCount != nil {
		toSerialize["maxEgressFlowCount"] = o.MaxEgressFlowCount
	}
	if o.MaxEndpointCount != nil {
		toSerialize["maxEndpointCount"] = o.MaxEndpointCount
	}
	if o.MaxIngressFlowCount != nil {
		toSerialize["maxIngressFlowCount"] = o.MaxIngressFlowCount
	}
	if o.MaxMsgSpoolUsage != nil {
		toSerialize["maxMsgSpoolUsage"] = o.MaxMsgSpoolUsage
	}
	if o.MaxSubscriptionCount != nil {
		toSerialize["maxSubscriptionCount"] = o.MaxSubscriptionCount
	}
	if o.MaxTransactedSessionCount != nil {
		toSerialize["maxTransactedSessionCount"] = o.MaxTransactedSessionCount
	}
	if o.MaxTransactionCount != nil {
		toSerialize["maxTransactionCount"] = o.MaxTransactionCount
	}
	if o.MqttRetainMaxMemory != nil {
		toSerialize["mqttRetainMaxMemory"] = o.MqttRetainMaxMemory
	}
	if o.MsgReplayActiveCount != nil {
		toSerialize["msgReplayActiveCount"] = o.MsgReplayActiveCount
	}
	if o.MsgReplayFailedCount != nil {
		toSerialize["msgReplayFailedCount"] = o.MsgReplayFailedCount
	}
	if o.MsgReplayInitializingCount != nil {
		toSerialize["msgReplayInitializingCount"] = o.MsgReplayInitializingCount
	}
	if o.MsgReplayPendingCompleteCount != nil {
		toSerialize["msgReplayPendingCompleteCount"] = o.MsgReplayPendingCompleteCount
	}
	if o.MsgSpoolMsgCount != nil {
		toSerialize["msgSpoolMsgCount"] = o.MsgSpoolMsgCount
	}
	if o.MsgSpoolRxMsgCount != nil {
		toSerialize["msgSpoolRxMsgCount"] = o.MsgSpoolRxMsgCount
	}
	if o.MsgSpoolTxMsgCount != nil {
		toSerialize["msgSpoolTxMsgCount"] = o.MsgSpoolTxMsgCount
	}
	if o.MsgSpoolUsage != nil {
		toSerialize["msgSpoolUsage"] = o.MsgSpoolUsage
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.ReplicationAckPropagationIntervalMsgCount != nil {
		toSerialize["replicationAckPropagationIntervalMsgCount"] = o.ReplicationAckPropagationIntervalMsgCount
	}
	if o.ReplicationActiveAckPropTxMsgCount != nil {
		toSerialize["replicationActiveAckPropTxMsgCount"] = o.ReplicationActiveAckPropTxMsgCount
	}
	if o.ReplicationActiveAsyncQueuedMsgCount != nil {
		toSerialize["replicationActiveAsyncQueuedMsgCount"] = o.ReplicationActiveAsyncQueuedMsgCount
	}
	if o.ReplicationActiveLocallyConsumedMsgCount != nil {
		toSerialize["replicationActiveLocallyConsumedMsgCount"] = o.ReplicationActiveLocallyConsumedMsgCount
	}
	if o.ReplicationActiveMateFlowCongestedPeakTime != nil {
		toSerialize["replicationActiveMateFlowCongestedPeakTime"] = o.ReplicationActiveMateFlowCongestedPeakTime
	}
	if o.ReplicationActiveMateFlowNotCongestedPeakTime != nil {
		toSerialize["replicationActiveMateFlowNotCongestedPeakTime"] = o.ReplicationActiveMateFlowNotCongestedPeakTime
	}
	if o.ReplicationActivePromotedQueuedMsgCount != nil {
		toSerialize["replicationActivePromotedQueuedMsgCount"] = o.ReplicationActivePromotedQueuedMsgCount
	}
	if o.ReplicationActiveReconcileRequestRxMsgCount != nil {
		toSerialize["replicationActiveReconcileRequestRxMsgCount"] = o.ReplicationActiveReconcileRequestRxMsgCount
	}
	if o.ReplicationActiveSyncEligiblePeakTime != nil {
		toSerialize["replicationActiveSyncEligiblePeakTime"] = o.ReplicationActiveSyncEligiblePeakTime
	}
	if o.ReplicationActiveSyncIneligiblePeakTime != nil {
		toSerialize["replicationActiveSyncIneligiblePeakTime"] = o.ReplicationActiveSyncIneligiblePeakTime
	}
	if o.ReplicationActiveSyncQueuedAsAsyncMsgCount != nil {
		toSerialize["replicationActiveSyncQueuedAsAsyncMsgCount"] = o.ReplicationActiveSyncQueuedAsAsyncMsgCount
	}
	if o.ReplicationActiveSyncQueuedMsgCount != nil {
		toSerialize["replicationActiveSyncQueuedMsgCount"] = o.ReplicationActiveSyncQueuedMsgCount
	}
	if o.ReplicationActiveTransitionToSyncIneligibleCount != nil {
		toSerialize["replicationActiveTransitionToSyncIneligibleCount"] = o.ReplicationActiveTransitionToSyncIneligibleCount
	}
	if o.ReplicationBridgeAuthenticationBasicClientUsername != nil {
		toSerialize["replicationBridgeAuthenticationBasicClientUsername"] = o.ReplicationBridgeAuthenticationBasicClientUsername
	}
	if o.ReplicationBridgeAuthenticationScheme != nil {
		toSerialize["replicationBridgeAuthenticationScheme"] = o.ReplicationBridgeAuthenticationScheme
	}
	if o.ReplicationBridgeBoundToQueue != nil {
		toSerialize["replicationBridgeBoundToQueue"] = o.ReplicationBridgeBoundToQueue
	}
	if o.ReplicationBridgeCompressedDataEnabled != nil {
		toSerialize["replicationBridgeCompressedDataEnabled"] = o.ReplicationBridgeCompressedDataEnabled
	}
	if o.ReplicationBridgeEgressFlowWindowSize != nil {
		toSerialize["replicationBridgeEgressFlowWindowSize"] = o.ReplicationBridgeEgressFlowWindowSize
	}
	if o.ReplicationBridgeName != nil {
		toSerialize["replicationBridgeName"] = o.ReplicationBridgeName
	}
	if o.ReplicationBridgeRetryDelay != nil {
		toSerialize["replicationBridgeRetryDelay"] = o.ReplicationBridgeRetryDelay
	}
	if o.ReplicationBridgeTlsEnabled != nil {
		toSerialize["replicationBridgeTlsEnabled"] = o.ReplicationBridgeTlsEnabled
	}
	if o.ReplicationBridgeUnidirectionalClientProfileName != nil {
		toSerialize["replicationBridgeUnidirectionalClientProfileName"] = o.ReplicationBridgeUnidirectionalClientProfileName
	}
	if o.ReplicationBridgeUp != nil {
		toSerialize["replicationBridgeUp"] = o.ReplicationBridgeUp
	}
	if o.ReplicationEnabled != nil {
		toSerialize["replicationEnabled"] = o.ReplicationEnabled
	}
	if o.ReplicationQueueBound != nil {
		toSerialize["replicationQueueBound"] = o.ReplicationQueueBound
	}
	if o.ReplicationQueueMaxMsgSpoolUsage != nil {
		toSerialize["replicationQueueMaxMsgSpoolUsage"] = o.ReplicationQueueMaxMsgSpoolUsage
	}
	if o.ReplicationQueueRejectMsgToSenderOnDiscardEnabled != nil {
		toSerialize["replicationQueueRejectMsgToSenderOnDiscardEnabled"] = o.ReplicationQueueRejectMsgToSenderOnDiscardEnabled
	}
	if o.ReplicationRejectMsgWhenSyncIneligibleEnabled != nil {
		toSerialize["replicationRejectMsgWhenSyncIneligibleEnabled"] = o.ReplicationRejectMsgWhenSyncIneligibleEnabled
	}
	if o.ReplicationRemoteBridgeName != nil {
		toSerialize["replicationRemoteBridgeName"] = o.ReplicationRemoteBridgeName
	}
	if o.ReplicationRemoteBridgeUp != nil {
		toSerialize["replicationRemoteBridgeUp"] = o.ReplicationRemoteBridgeUp
	}
	if o.ReplicationRole != nil {
		toSerialize["replicationRole"] = o.ReplicationRole
	}
	if o.ReplicationStandbyAckPropOutOfSeqRxMsgCount != nil {
		toSerialize["replicationStandbyAckPropOutOfSeqRxMsgCount"] = o.ReplicationStandbyAckPropOutOfSeqRxMsgCount
	}
	if o.ReplicationStandbyAckPropRxMsgCount != nil {
		toSerialize["replicationStandbyAckPropRxMsgCount"] = o.ReplicationStandbyAckPropRxMsgCount
	}
	if o.ReplicationStandbyReconcileRequestTxMsgCount != nil {
		toSerialize["replicationStandbyReconcileRequestTxMsgCount"] = o.ReplicationStandbyReconcileRequestTxMsgCount
	}
	if o.ReplicationStandbyRxMsgCount != nil {
		toSerialize["replicationStandbyRxMsgCount"] = o.ReplicationStandbyRxMsgCount
	}
	if o.ReplicationStandbyTransactionRequestCount != nil {
		toSerialize["replicationStandbyTransactionRequestCount"] = o.ReplicationStandbyTransactionRequestCount
	}
	if o.ReplicationStandbyTransactionRequestFailureCount != nil {
		toSerialize["replicationStandbyTransactionRequestFailureCount"] = o.ReplicationStandbyTransactionRequestFailureCount
	}
	if o.ReplicationStandbyTransactionRequestSuccessCount != nil {
		toSerialize["replicationStandbyTransactionRequestSuccessCount"] = o.ReplicationStandbyTransactionRequestSuccessCount
	}
	if o.ReplicationSyncEligible != nil {
		toSerialize["replicationSyncEligible"] = o.ReplicationSyncEligible
	}
	if o.ReplicationTransactionMode != nil {
		toSerialize["replicationTransactionMode"] = o.ReplicationTransactionMode
	}
	if o.RestTlsServerCertEnforceTrustedCommonNameEnabled != nil {
		toSerialize["restTlsServerCertEnforceTrustedCommonNameEnabled"] = o.RestTlsServerCertEnforceTrustedCommonNameEnabled
	}
	if o.RestTlsServerCertMaxChainDepth != nil {
		toSerialize["restTlsServerCertMaxChainDepth"] = o.RestTlsServerCertMaxChainDepth
	}
	if o.RestTlsServerCertValidateDateEnabled != nil {
		toSerialize["restTlsServerCertValidateDateEnabled"] = o.RestTlsServerCertValidateDateEnabled
	}
	if o.RestTlsServerCertValidateNameEnabled != nil {
		toSerialize["restTlsServerCertValidateNameEnabled"] = o.RestTlsServerCertValidateNameEnabled
	}
	if o.RxByteCount != nil {
		toSerialize["rxByteCount"] = o.RxByteCount
	}
	if o.RxByteRate != nil {
		toSerialize["rxByteRate"] = o.RxByteRate
	}
	if o.RxCompressedByteCount != nil {
		toSerialize["rxCompressedByteCount"] = o.RxCompressedByteCount
	}
	if o.RxCompressedByteRate != nil {
		toSerialize["rxCompressedByteRate"] = o.RxCompressedByteRate
	}
	if o.RxCompressionRatio != nil {
		toSerialize["rxCompressionRatio"] = o.RxCompressionRatio
	}
	if o.RxMsgCount != nil {
		toSerialize["rxMsgCount"] = o.RxMsgCount
	}
	if o.RxMsgRate != nil {
		toSerialize["rxMsgRate"] = o.RxMsgRate
	}
	if o.RxUncompressedByteCount != nil {
		toSerialize["rxUncompressedByteCount"] = o.RxUncompressedByteCount
	}
	if o.RxUncompressedByteRate != nil {
		toSerialize["rxUncompressedByteRate"] = o.RxUncompressedByteRate
	}
	if o.SempOverMsgBusAdminClientEnabled != nil {
		toSerialize["sempOverMsgBusAdminClientEnabled"] = o.SempOverMsgBusAdminClientEnabled
	}
	if o.SempOverMsgBusAdminDistributedCacheEnabled != nil {
		toSerialize["sempOverMsgBusAdminDistributedCacheEnabled"] = o.SempOverMsgBusAdminDistributedCacheEnabled
	}
	if o.SempOverMsgBusAdminEnabled != nil {
		toSerialize["sempOverMsgBusAdminEnabled"] = o.SempOverMsgBusAdminEnabled
	}
	if o.SempOverMsgBusEnabled != nil {
		toSerialize["sempOverMsgBusEnabled"] = o.SempOverMsgBusEnabled
	}
	if o.SempOverMsgBusShowEnabled != nil {
		toSerialize["sempOverMsgBusShowEnabled"] = o.SempOverMsgBusShowEnabled
	}
	if o.ServiceAmqpMaxConnectionCount != nil {
		toSerialize["serviceAmqpMaxConnectionCount"] = o.ServiceAmqpMaxConnectionCount
	}
	if o.ServiceAmqpPlainTextCompressed != nil {
		toSerialize["serviceAmqpPlainTextCompressed"] = o.ServiceAmqpPlainTextCompressed
	}
	if o.ServiceAmqpPlainTextEnabled != nil {
		toSerialize["serviceAmqpPlainTextEnabled"] = o.ServiceAmqpPlainTextEnabled
	}
	if o.ServiceAmqpPlainTextFailureReason != nil {
		toSerialize["serviceAmqpPlainTextFailureReason"] = o.ServiceAmqpPlainTextFailureReason
	}
	if o.ServiceAmqpPlainTextListenPort != nil {
		toSerialize["serviceAmqpPlainTextListenPort"] = o.ServiceAmqpPlainTextListenPort
	}
	if o.ServiceAmqpPlainTextUp != nil {
		toSerialize["serviceAmqpPlainTextUp"] = o.ServiceAmqpPlainTextUp
	}
	if o.ServiceAmqpTlsCompressed != nil {
		toSerialize["serviceAmqpTlsCompressed"] = o.ServiceAmqpTlsCompressed
	}
	if o.ServiceAmqpTlsEnabled != nil {
		toSerialize["serviceAmqpTlsEnabled"] = o.ServiceAmqpTlsEnabled
	}
	if o.ServiceAmqpTlsFailureReason != nil {
		toSerialize["serviceAmqpTlsFailureReason"] = o.ServiceAmqpTlsFailureReason
	}
	if o.ServiceAmqpTlsListenPort != nil {
		toSerialize["serviceAmqpTlsListenPort"] = o.ServiceAmqpTlsListenPort
	}
	if o.ServiceAmqpTlsUp != nil {
		toSerialize["serviceAmqpTlsUp"] = o.ServiceAmqpTlsUp
	}
	if o.ServiceMqttAuthenticationClientCertRequest != nil {
		toSerialize["serviceMqttAuthenticationClientCertRequest"] = o.ServiceMqttAuthenticationClientCertRequest
	}
	if o.ServiceMqttMaxConnectionCount != nil {
		toSerialize["serviceMqttMaxConnectionCount"] = o.ServiceMqttMaxConnectionCount
	}
	if o.ServiceMqttPlainTextCompressed != nil {
		toSerialize["serviceMqttPlainTextCompressed"] = o.ServiceMqttPlainTextCompressed
	}
	if o.ServiceMqttPlainTextEnabled != nil {
		toSerialize["serviceMqttPlainTextEnabled"] = o.ServiceMqttPlainTextEnabled
	}
	if o.ServiceMqttPlainTextFailureReason != nil {
		toSerialize["serviceMqttPlainTextFailureReason"] = o.ServiceMqttPlainTextFailureReason
	}
	if o.ServiceMqttPlainTextListenPort != nil {
		toSerialize["serviceMqttPlainTextListenPort"] = o.ServiceMqttPlainTextListenPort
	}
	if o.ServiceMqttPlainTextUp != nil {
		toSerialize["serviceMqttPlainTextUp"] = o.ServiceMqttPlainTextUp
	}
	if o.ServiceMqttTlsCompressed != nil {
		toSerialize["serviceMqttTlsCompressed"] = o.ServiceMqttTlsCompressed
	}
	if o.ServiceMqttTlsEnabled != nil {
		toSerialize["serviceMqttTlsEnabled"] = o.ServiceMqttTlsEnabled
	}
	if o.ServiceMqttTlsFailureReason != nil {
		toSerialize["serviceMqttTlsFailureReason"] = o.ServiceMqttTlsFailureReason
	}
	if o.ServiceMqttTlsListenPort != nil {
		toSerialize["serviceMqttTlsListenPort"] = o.ServiceMqttTlsListenPort
	}
	if o.ServiceMqttTlsUp != nil {
		toSerialize["serviceMqttTlsUp"] = o.ServiceMqttTlsUp
	}
	if o.ServiceMqttTlsWebSocketCompressed != nil {
		toSerialize["serviceMqttTlsWebSocketCompressed"] = o.ServiceMqttTlsWebSocketCompressed
	}
	if o.ServiceMqttTlsWebSocketEnabled != nil {
		toSerialize["serviceMqttTlsWebSocketEnabled"] = o.ServiceMqttTlsWebSocketEnabled
	}
	if o.ServiceMqttTlsWebSocketFailureReason != nil {
		toSerialize["serviceMqttTlsWebSocketFailureReason"] = o.ServiceMqttTlsWebSocketFailureReason
	}
	if o.ServiceMqttTlsWebSocketListenPort != nil {
		toSerialize["serviceMqttTlsWebSocketListenPort"] = o.ServiceMqttTlsWebSocketListenPort
	}
	if o.ServiceMqttTlsWebSocketUp != nil {
		toSerialize["serviceMqttTlsWebSocketUp"] = o.ServiceMqttTlsWebSocketUp
	}
	if o.ServiceMqttWebSocketCompressed != nil {
		toSerialize["serviceMqttWebSocketCompressed"] = o.ServiceMqttWebSocketCompressed
	}
	if o.ServiceMqttWebSocketEnabled != nil {
		toSerialize["serviceMqttWebSocketEnabled"] = o.ServiceMqttWebSocketEnabled
	}
	if o.ServiceMqttWebSocketFailureReason != nil {
		toSerialize["serviceMqttWebSocketFailureReason"] = o.ServiceMqttWebSocketFailureReason
	}
	if o.ServiceMqttWebSocketListenPort != nil {
		toSerialize["serviceMqttWebSocketListenPort"] = o.ServiceMqttWebSocketListenPort
	}
	if o.ServiceMqttWebSocketUp != nil {
		toSerialize["serviceMqttWebSocketUp"] = o.ServiceMqttWebSocketUp
	}
	if o.ServiceRestIncomingAuthenticationClientCertRequest != nil {
		toSerialize["serviceRestIncomingAuthenticationClientCertRequest"] = o.ServiceRestIncomingAuthenticationClientCertRequest
	}
	if o.ServiceRestIncomingAuthorizationHeaderHandling != nil {
		toSerialize["serviceRestIncomingAuthorizationHeaderHandling"] = o.ServiceRestIncomingAuthorizationHeaderHandling
	}
	if o.ServiceRestIncomingMaxConnectionCount != nil {
		toSerialize["serviceRestIncomingMaxConnectionCount"] = o.ServiceRestIncomingMaxConnectionCount
	}
	if o.ServiceRestIncomingPlainTextCompressed != nil {
		toSerialize["serviceRestIncomingPlainTextCompressed"] = o.ServiceRestIncomingPlainTextCompressed
	}
	if o.ServiceRestIncomingPlainTextEnabled != nil {
		toSerialize["serviceRestIncomingPlainTextEnabled"] = o.ServiceRestIncomingPlainTextEnabled
	}
	if o.ServiceRestIncomingPlainTextFailureReason != nil {
		toSerialize["serviceRestIncomingPlainTextFailureReason"] = o.ServiceRestIncomingPlainTextFailureReason
	}
	if o.ServiceRestIncomingPlainTextListenPort != nil {
		toSerialize["serviceRestIncomingPlainTextListenPort"] = o.ServiceRestIncomingPlainTextListenPort
	}
	if o.ServiceRestIncomingPlainTextUp != nil {
		toSerialize["serviceRestIncomingPlainTextUp"] = o.ServiceRestIncomingPlainTextUp
	}
	if o.ServiceRestIncomingTlsCompressed != nil {
		toSerialize["serviceRestIncomingTlsCompressed"] = o.ServiceRestIncomingTlsCompressed
	}
	if o.ServiceRestIncomingTlsEnabled != nil {
		toSerialize["serviceRestIncomingTlsEnabled"] = o.ServiceRestIncomingTlsEnabled
	}
	if o.ServiceRestIncomingTlsFailureReason != nil {
		toSerialize["serviceRestIncomingTlsFailureReason"] = o.ServiceRestIncomingTlsFailureReason
	}
	if o.ServiceRestIncomingTlsListenPort != nil {
		toSerialize["serviceRestIncomingTlsListenPort"] = o.ServiceRestIncomingTlsListenPort
	}
	if o.ServiceRestIncomingTlsUp != nil {
		toSerialize["serviceRestIncomingTlsUp"] = o.ServiceRestIncomingTlsUp
	}
	if o.ServiceRestMode != nil {
		toSerialize["serviceRestMode"] = o.ServiceRestMode
	}
	if o.ServiceRestOutgoingMaxConnectionCount != nil {
		toSerialize["serviceRestOutgoingMaxConnectionCount"] = o.ServiceRestOutgoingMaxConnectionCount
	}
	if o.ServiceSmfMaxConnectionCount != nil {
		toSerialize["serviceSmfMaxConnectionCount"] = o.ServiceSmfMaxConnectionCount
	}
	if o.ServiceSmfPlainTextEnabled != nil {
		toSerialize["serviceSmfPlainTextEnabled"] = o.ServiceSmfPlainTextEnabled
	}
	if o.ServiceSmfPlainTextFailureReason != nil {
		toSerialize["serviceSmfPlainTextFailureReason"] = o.ServiceSmfPlainTextFailureReason
	}
	if o.ServiceSmfPlainTextUp != nil {
		toSerialize["serviceSmfPlainTextUp"] = o.ServiceSmfPlainTextUp
	}
	if o.ServiceSmfTlsEnabled != nil {
		toSerialize["serviceSmfTlsEnabled"] = o.ServiceSmfTlsEnabled
	}
	if o.ServiceSmfTlsFailureReason != nil {
		toSerialize["serviceSmfTlsFailureReason"] = o.ServiceSmfTlsFailureReason
	}
	if o.ServiceSmfTlsUp != nil {
		toSerialize["serviceSmfTlsUp"] = o.ServiceSmfTlsUp
	}
	if o.ServiceWebAuthenticationClientCertRequest != nil {
		toSerialize["serviceWebAuthenticationClientCertRequest"] = o.ServiceWebAuthenticationClientCertRequest
	}
	if o.ServiceWebMaxConnectionCount != nil {
		toSerialize["serviceWebMaxConnectionCount"] = o.ServiceWebMaxConnectionCount
	}
	if o.ServiceWebPlainTextEnabled != nil {
		toSerialize["serviceWebPlainTextEnabled"] = o.ServiceWebPlainTextEnabled
	}
	if o.ServiceWebPlainTextFailureReason != nil {
		toSerialize["serviceWebPlainTextFailureReason"] = o.ServiceWebPlainTextFailureReason
	}
	if o.ServiceWebPlainTextUp != nil {
		toSerialize["serviceWebPlainTextUp"] = o.ServiceWebPlainTextUp
	}
	if o.ServiceWebTlsEnabled != nil {
		toSerialize["serviceWebTlsEnabled"] = o.ServiceWebTlsEnabled
	}
	if o.ServiceWebTlsFailureReason != nil {
		toSerialize["serviceWebTlsFailureReason"] = o.ServiceWebTlsFailureReason
	}
	if o.ServiceWebTlsUp != nil {
		toSerialize["serviceWebTlsUp"] = o.ServiceWebTlsUp
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.SubscriptionExportProgress != nil {
		toSerialize["subscriptionExportProgress"] = o.SubscriptionExportProgress
	}
	if o.SystemManager != nil {
		toSerialize["systemManager"] = o.SystemManager
	}
	if o.TlsAllowDowngradeToPlainTextEnabled != nil {
		toSerialize["tlsAllowDowngradeToPlainTextEnabled"] = o.TlsAllowDowngradeToPlainTextEnabled
	}
	if o.TlsAverageRxByteRate != nil {
		toSerialize["tlsAverageRxByteRate"] = o.TlsAverageRxByteRate
	}
	if o.TlsAverageTxByteRate != nil {
		toSerialize["tlsAverageTxByteRate"] = o.TlsAverageTxByteRate
	}
	if o.TlsRxByteCount != nil {
		toSerialize["tlsRxByteCount"] = o.TlsRxByteCount
	}
	if o.TlsRxByteRate != nil {
		toSerialize["tlsRxByteRate"] = o.TlsRxByteRate
	}
	if o.TlsTxByteCount != nil {
		toSerialize["tlsTxByteCount"] = o.TlsTxByteCount
	}
	if o.TlsTxByteRate != nil {
		toSerialize["tlsTxByteRate"] = o.TlsTxByteRate
	}
	if o.TxByteCount != nil {
		toSerialize["txByteCount"] = o.TxByteCount
	}
	if o.TxByteRate != nil {
		toSerialize["txByteRate"] = o.TxByteRate
	}
	if o.TxCompressedByteCount != nil {
		toSerialize["txCompressedByteCount"] = o.TxCompressedByteCount
	}
	if o.TxCompressedByteRate != nil {
		toSerialize["txCompressedByteRate"] = o.TxCompressedByteRate
	}
	if o.TxCompressionRatio != nil {
		toSerialize["txCompressionRatio"] = o.TxCompressionRatio
	}
	if o.TxMsgCount != nil {
		toSerialize["txMsgCount"] = o.TxMsgCount
	}
	if o.TxMsgRate != nil {
		toSerialize["txMsgRate"] = o.TxMsgRate
	}
	if o.TxUncompressedByteCount != nil {
		toSerialize["txUncompressedByteCount"] = o.TxUncompressedByteCount
	}
	if o.TxUncompressedByteRate != nil {
		toSerialize["txUncompressedByteRate"] = o.TxUncompressedByteRate
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpn struct {
	value *MsgVpn
	isSet bool
}

func (v NullableMsgVpn) Get() *MsgVpn {
	return v.value
}

func (v *NullableMsgVpn) Set(val *MsgVpn) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpn) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpn(val *MsgVpn) *NullableMsgVpn {
	return &NullableMsgVpn{value: val, isSet: true}
}

func (v NullableMsgVpn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
