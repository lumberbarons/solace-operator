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

// MsgVpnTopicEndpoint struct for MsgVpnTopicEndpoint
type MsgVpnTopicEndpoint struct {
	// The access type for delivering messages to consumer flows bound to the Topic Endpoint. The allowed values and their meaning are:  <pre> \"exclusive\" - Exclusive delivery of messages to the first bound consumer flow. \"non-exclusive\" - Non-exclusive delivery of messages to all bound consumer flows in a round-robin fashion. </pre>
	AccessType *string `json:"accessType,omitempty"`
	// The number of Topic Endpoint bind failures due to being already bound.
	AlreadyBoundBindFailureCount *int64 `json:"alreadyBoundBindFailureCount,omitempty"`
	// The one minute average of the message rate received by the Topic Endpoint, in bytes per second (B/sec).
	AverageRxByteRate *int64 `json:"averageRxByteRate,omitempty"`
	// The one minute average of the message rate received by the Topic Endpoint, in messages per second (msg/sec).
	AverageRxMsgRate *int64 `json:"averageRxMsgRate,omitempty"`
	// The one minute average of the message rate transmitted by the Topic Endpoint, in bytes per second (B/sec).
	AverageTxByteRate *int64 `json:"averageTxByteRate,omitempty"`
	// The one minute average of the message rate transmitted by the Topic Endpoint, in messages per second (msg/sec).
	AverageTxMsgRate *int64 `json:"averageTxMsgRate,omitempty"`
	// The number of consumer requests to bind to the Topic Endpoint.
	BindRequestCount *int64 `json:"bindRequestCount,omitempty"`
	// The number of successful consumer requests to bind to the Topic Endpoint.
	BindSuccessCount *int64 `json:"bindSuccessCount,omitempty"`
	// The forwarding mode of the Topic Endpoint at bind time. The allowed values and their meaning are:  <pre> \"store-and-forward\" - Deliver messages using the guaranteed data path. \"cut-through\" - Deliver messages using the direct and guaranteed data paths for lower latency. </pre>
	BindTimeForwardingMode *string `json:"bindTimeForwardingMode,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to being denied by the Client Profile.
	ClientProfileDeniedDiscardedMsgCount *int64 `json:"clientProfileDeniedDiscardedMsgCount,omitempty"`
	// Indicates whether the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN is enabled.
	ConsumerAckPropagationEnabled *bool `json:"consumerAckPropagationEnabled,omitempty"`
	// Indicates whether the Topic Endpoint was created by a management API (CLI or SEMP).
	CreatedByManagement *bool `json:"createdByManagement,omitempty"`
	// The name of the Dead Message Queue (DMQ) used by the Topic Endpoint.
	DeadMsgQueue *string `json:"deadMsgQueue,omitempty"`
	// The number of guaranteed messages deleted from the Topic Endpoint.
	DeletedMsgCount *int64 `json:"deletedMsgCount,omitempty"`
	// Enable or disable the ability for client applications to query the message delivery count of messages received from the Topic Endpoint. This is a controlled availability feature. Please contact Solace to find out if this feature is supported for your use case. Available since 2.19.
	DeliveryCountEnabled *bool `json:"deliveryCountEnabled,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to a destination group error.
	DestinationGroupErrorDiscardedMsgCount *int64 `json:"destinationGroupErrorDiscardedMsgCount,omitempty"`
	// The destination topic of the Topic Endpoint.
	DestinationTopic *string `json:"destinationTopic,omitempty"`
	// The number of Topic Endpoint bind failures due to being disabled.
	DisabledBindFailureCount *int64 `json:"disabledBindFailureCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to it being disabled.
	DisabledDiscardedMsgCount *int64 `json:"disabledDiscardedMsgCount,omitempty"`
	// Indicates whether the Topic Endpoint is durable and not temporary.
	Durable *bool `json:"durable,omitempty"`
	// Indicates whether the transmission of messages from the Topic Endpoint is enabled.
	EgressEnabled                           *bool           `json:"egressEnabled,omitempty"`
	EventBindCountThreshold                 *EventThreshold `json:"eventBindCountThreshold,omitempty"`
	EventRejectLowPriorityMsgLimitThreshold *EventThreshold `json:"eventRejectLowPriorityMsgLimitThreshold,omitempty"`
	EventSpoolUsageThreshold                *EventThreshold `json:"eventSpoolUsageThreshold,omitempty"`
	// The highest identifier (ID) of guaranteed messages in the Topic Endpoint that were acknowledged.
	HighestAckedMsgId *int64 `json:"highestAckedMsgId,omitempty"`
	// The highest identifier (ID) of guaranteed messages in the Topic Endpoint.
	HighestMsgId *int64 `json:"highestMsgId,omitempty"`
	// The number of acknowledgement messages received by the Topic Endpoint that are in the process of updating and deleting associated guaranteed messages.
	InProgressAckMsgCount *int64 `json:"inProgressAckMsgCount,omitempty"`
	// Indicates whether the reception of messages to the Topic Endpoint is enabled.
	IngressEnabled *bool `json:"ingressEnabled,omitempty"`
	// The number of Topic Endpoint bind failures due to an invalid selector.
	InvalidSelectorBindFailureCount *int64 `json:"invalidSelectorBindFailureCount,omitempty"`
	// The timestamp of the last completed replay for the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastReplayCompleteTime *int32 `json:"lastReplayCompleteTime,omitempty"`
	// The reason for the last replay failure for the Topic Endpoint.
	LastReplayFailureReason *string `json:"lastReplayFailureReason,omitempty"`
	// The timestamp of the last replay failure for the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastReplayFailureTime *int32 `json:"lastReplayFailureTime,omitempty"`
	// The timestamp of the last replay started for the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastReplayStartTime *int32 `json:"lastReplayStartTime,omitempty"`
	// The timestamp of the last replayed message transmitted by the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time).
	LastReplayedMsgTxTime *int32 `json:"lastReplayedMsgTxTime,omitempty"`
	// The identifier (ID) of the last message examined by the Topic Endpoint selector.
	LastSelectorExaminedMsgId *int64 `json:"lastSelectorExaminedMsgId,omitempty"`
	// The identifier (ID) of the last guaranteed message spooled in the Topic Endpoint.
	LastSpooledMsgId *int64 `json:"lastSpooledMsgId,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to low priority message congestion control.
	LowPriorityMsgCongestionDiscardedMsgCount *int64 `json:"lowPriorityMsgCongestionDiscardedMsgCount,omitempty"`
	// The state of the low priority message congestion in the Topic Endpoint. The allowed values and their meaning are:  <pre> \"disabled\" - Messages are not being checked for priority. \"not-congested\" - Low priority messages are being stored and delivered. \"congested\" - Low priority messages are being discarded. </pre>
	LowPriorityMsgCongestionState *string `json:"lowPriorityMsgCongestionState,omitempty"`
	// The lowest identifier (ID) of guaranteed messages in the Topic Endpoint that were acknowledged.
	LowestAckedMsgId *int64 `json:"lowestAckedMsgId,omitempty"`
	// The lowest identifier (ID) of guaranteed messages in the Topic Endpoint.
	LowestMsgId *int64 `json:"lowestMsgId,omitempty"`
	// The maximum number of consumer flows that can bind to the Topic Endpoint.
	MaxBindCount *int64 `json:"maxBindCount,omitempty"`
	// The number of Topic Endpoint bind failures due to the maximum bind count being exceeded.
	MaxBindCountExceededBindFailureCount *int64 `json:"maxBindCountExceededBindFailureCount,omitempty"`
	// The maximum number of messages delivered but not acknowledged per flow for the Topic Endpoint.
	MaxDeliveredUnackedMsgsPerFlow *int64 `json:"maxDeliveredUnackedMsgsPerFlow,omitempty"`
	// The effective maximum number of consumer flows that can bind to the Topic Endpoint.
	MaxEffectiveBindCount *int32 `json:"maxEffectiveBindCount,omitempty"`
	// The maximum message size allowed in the Topic Endpoint, in bytes (B).
	MaxMsgSize *int32 `json:"maxMsgSize,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum message size being exceeded.
	MaxMsgSizeExceededDiscardedMsgCount *int64 `json:"maxMsgSizeExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum message spool usage being exceeded.
	MaxMsgSpoolUsageExceededDiscardedMsgCount *int64 `json:"maxMsgSpoolUsageExceededDiscardedMsgCount,omitempty"`
	// The maximum number of times the Topic Endpoint will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever.
	MaxRedeliveryCount *int64 `json:"maxRedeliveryCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum redelivery attempts being exceeded.
	MaxRedeliveryExceededDiscardedMsgCount *int64 `json:"maxRedeliveryExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum redelivery attempts being exceeded and failing to move to the Dead Message Queue (DMQ).
	MaxRedeliveryExceededToDmqFailedMsgCount *int64 `json:"maxRedeliveryExceededToDmqFailedMsgCount,omitempty"`
	// The number of guaranteed messages moved to the Dead Message Queue (DMQ) by the Topic Endpoint due to the maximum redelivery attempts being exceeded.
	MaxRedeliveryExceededToDmqMsgCount *int64 `json:"maxRedeliveryExceededToDmqMsgCount,omitempty"`
	// The maximum message spool usage allowed by the Topic Endpoint, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking.
	MaxSpoolUsage *int64 `json:"maxSpoolUsage,omitempty"`
	// The maximum time in seconds a message can stay in the Topic Endpoint when `respectTtlEnabled` is `\"true\"`. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the `maxTtl` configured for the Topic Endpoint, is exceeded. A value of 0 disables expiry.
	MaxTtl *int64 `json:"maxTtl,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum time-to-live (TTL) in hops being exceeded. The TTL hop count is incremented when the message crosses a bridge.
	MaxTtlExceededDiscardedMsgCount *int64 `json:"maxTtlExceededDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum time-to-live (TTL) timestamp expiring.
	MaxTtlExpiredDiscardedMsgCount *int64 `json:"maxTtlExpiredDiscardedMsgCount,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to the maximum time-to-live (TTL) timestamp expiring and failing to move to the Dead Message Queue (DMQ).
	MaxTtlExpiredToDmqFailedMsgCount *int64 `json:"maxTtlExpiredToDmqFailedMsgCount,omitempty"`
	// The number of guaranteed messages moved to the Dead Message Queue (DMQ) by the Topic Endpoint due to the maximum time-to-live (TTL) timestamp expiring.
	MaxTtlExpiredToDmqMsgCount *int64 `json:"maxTtlExpiredToDmqMsgCount,omitempty"`
	// The message spool peak usage by the Topic Endpoint, in bytes (B).
	MsgSpoolPeakUsage *int64 `json:"msgSpoolPeakUsage,omitempty"`
	// The message spool usage by the Topic Endpoint, in bytes (B).
	MsgSpoolUsage *int64 `json:"msgSpoolUsage,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// The name of the network topic for the Topic Endpoint.
	NetworkTopic *string `json:"networkTopic,omitempty"`
	// The number of guaranteed messages discarded by the Topic Endpoint due to no local delivery being requested.
	NoLocalDeliveryDiscardedMsgCount *int64 `json:"noLocalDeliveryDiscardedMsgCount,omitempty"`
	// The number of Topic Endpoint bind failures due to other reasons.
	OtherBindFailureCount *int64 `json:"otherBindFailureCount,omitempty"`
	// The Client Username that owns the Topic Endpoint and has permission equivalent to `\"delete\"`.
	Owner *string `json:"owner,omitempty"`
	// The permission level for all consumers of the Topic Endpoint, excluding the owner. The allowed values and their meaning are:  <pre> \"no-access\" - Disallows all access. \"read-only\" - Read-only access to the messages. \"consume\" - Consume (read and remove) messages. \"modify-topic\" - Consume messages or modify the topic/selector. \"delete\" - Consume messages, modify the topic/selector or delete the Client created endpoint altogether. </pre>
	Permission *string `json:"permission,omitempty"`
	// The number of guaranteed messages transmitted by the Topic Endpoint for redelivery.
	RedeliveredMsgCount *int64 `json:"redeliveredMsgCount,omitempty"`
	// Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the topic-endpoint more than once. Available since 2.18.
	RedeliveryEnabled *bool `json:"redeliveryEnabled,omitempty"`
	// Indicates whether the checking of low priority messages against the `rejectLowPriorityMsgLimit` is enabled.
	RejectLowPriorityMsgEnabled *bool `json:"rejectLowPriorityMsgEnabled,omitempty"`
	// The number of messages of any priority in the Topic Endpoint above which low priority messages are not admitted but higher priority messages are allowed.
	RejectLowPriorityMsgLimit *int64 `json:"rejectLowPriorityMsgLimit,omitempty"`
	// Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The allowed values and their meaning are:  <pre> \"always\" - Always return a negative acknowledgment (NACK) to the sending client on message discard. \"when-topic-endpoint-enabled\" - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Topic Endpoint is enabled. \"never\" - Never return a negative acknowledgment (NACK) to the sending client on message discard. </pre>
	RejectMsgToSenderOnDiscardBehavior *string `json:"rejectMsgToSenderOnDiscardBehavior,omitempty"`
	// The number of replays that failed for the Topic Endpoint.
	ReplayFailureCount *int64 `json:"replayFailureCount,omitempty"`
	// The number of replays started for the Topic Endpoint.
	ReplayStartCount *int64 `json:"replayStartCount,omitempty"`
	// The state of replay for the Topic Endpoint. The allowed values and their meaning are:  <pre> \"initializing\" - All messages are being deleted from the endpoint before replay starts. \"active\" - Subscription matching logged messages are being replayed to the endpoint. \"pending-complete\" - Replay is complete, but final accounting is in progress. \"complete\" - Replay and all related activities are complete. \"failed\" - Replay has failed and is waiting for an unbind response. </pre>
	ReplayState *string `json:"replayState,omitempty"`
	// The number of replays that succeeded for the Topic Endpoint.
	ReplaySuccessCount *int64 `json:"replaySuccessCount,omitempty"`
	// The number of replayed messages transmitted by the Topic Endpoint and acked by all consumers.
	ReplayedAckedMsgCount *int64 `json:"replayedAckedMsgCount,omitempty"`
	// The number of replayed messages transmitted by the Topic Endpoint.
	ReplayedTxMsgCount *int64 `json:"replayedTxMsgCount,omitempty"`
	// The number of acknowledgement messages propagated by the Topic Endpoint to the replication standby remote Message VPN.
	ReplicationActiveAckPropTxMsgCount *int64 `json:"replicationActiveAckPropTxMsgCount,omitempty"`
	// The number of propagated acknowledgement messages received by the Topic Endpoint from the replication active remote Message VPN.
	ReplicationStandbyAckPropRxMsgCount *int64 `json:"replicationStandbyAckPropRxMsgCount,omitempty"`
	// The number of messages acknowledged in the Topic Endpoint by acknowledgement propagation from the replication active remote Message VPN.
	ReplicationStandbyAckedByAckPropMsgCount *int64 `json:"replicationStandbyAckedByAckPropMsgCount,omitempty"`
	// The number of messages received by the Topic Endpoint from the replication active remote Message VPN.
	ReplicationStandbyRxMsgCount *int64 `json:"replicationStandbyRxMsgCount,omitempty"`
	// Indicates whether message priorities are respected. When enabled, messages contained in the Topic Endpoint are delivered in priority order, from 9 (highest) to 0 (lowest).
	RespectMsgPriorityEnabled *bool `json:"respectMsgPriorityEnabled,omitempty"`
	// Indicates whether the time-to-live (TTL) for messages in the Topic Endpoint is respected. When enabled, expired messages are discarded or moved to the DMQ.
	RespectTtlEnabled *bool `json:"respectTtlEnabled,omitempty"`
	// The current message rate received by the Topic Endpoint, in bytes per second (B/sec).
	RxByteRate *int32 `json:"rxByteRate,omitempty"`
	// The current message rate received by the Topic Endpoint, in messages per second (msg/sec).
	RxMsgRate *int64 `json:"rxMsgRate,omitempty"`
	// Indicates whether the Topic Endpoint has a selector to filter received messages.
	RxSelector *bool `json:"rxSelector,omitempty"`
	// The value of the receive selector for the Topic Endpoint.
	Selector *string `json:"selector,omitempty"`
	// The number of guaranteed messages examined by the Topic Endpoint selector.
	SelectorExaminedMsgCount *int64 `json:"selectorExaminedMsgCount,omitempty"`
	// The number of guaranteed messages for which the Topic Endpoint selector matched.
	SelectorMatchedMsgCount *int64 `json:"selectorMatchedMsgCount,omitempty"`
	// The number of guaranteed messages for which the Topic Endpoint selector did not match.
	SelectorNotMatchedMsgCount *int64 `json:"selectorNotMatchedMsgCount,omitempty"`
	// The amount of guaranteed messages that were spooled in the Topic Endpoint, in bytes (B).
	SpooledByteCount *int64 `json:"spooledByteCount,omitempty"`
	// The number of guaranteed messages that were spooled in the Topic Endpoint.
	SpooledMsgCount *int64 `json:"spooledMsgCount,omitempty"`
	// The name of the Topic Endpoint.
	TopicEndpointName *string `json:"topicEndpointName,omitempty"`
	// The number of guaranteed messages that were retransmitted by the Topic Endpoint at the transport layer as part of a single delivery attempt. Available since 2.18.
	TransportRetransmitMsgCount *int64 `json:"transportRetransmitMsgCount,omitempty"`
	// The current message rate transmitted by the Topic Endpoint, in bytes per second (B/sec).
	TxByteRate *int64 `json:"txByteRate,omitempty"`
	// The current message rate transmitted by the Topic Endpoint, in messages per second (msg/sec).
	TxMsgRate *int64 `json:"txMsgRate,omitempty"`
	// The number of guaranteed messages in the Topic Endpoint that have been transmitted but not acknowledged by all consumers.
	TxUnackedMsgCount *int64 `json:"txUnackedMsgCount,omitempty"`
	// The virtual router used by the Topic Endpoint. The allowed values and their meaning are:  <pre> \"primary\" - The endpoint belongs to the primary virtual router. \"backup\" - The endpoint belongs to the backup virtual router. </pre>
	VirtualRouter *string `json:"virtualRouter,omitempty"`
}

// NewMsgVpnTopicEndpoint instantiates a new MsgVpnTopicEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnTopicEndpoint() *MsgVpnTopicEndpoint {
	this := MsgVpnTopicEndpoint{}
	return &this
}

// NewMsgVpnTopicEndpointWithDefaults instantiates a new MsgVpnTopicEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnTopicEndpointWithDefaults() *MsgVpnTopicEndpoint {
	this := MsgVpnTopicEndpoint{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetAccessType() string {
	if o == nil || o.AccessType == nil {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetAccessTypeOk() (*string, bool) {
	if o == nil || o.AccessType == nil {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasAccessType() bool {
	if o != nil && o.AccessType != nil {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *MsgVpnTopicEndpoint) SetAccessType(v string) {
	o.AccessType = &v
}

// GetAlreadyBoundBindFailureCount returns the AlreadyBoundBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetAlreadyBoundBindFailureCount() int64 {
	if o == nil || o.AlreadyBoundBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.AlreadyBoundBindFailureCount
}

// GetAlreadyBoundBindFailureCountOk returns a tuple with the AlreadyBoundBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetAlreadyBoundBindFailureCountOk() (*int64, bool) {
	if o == nil || o.AlreadyBoundBindFailureCount == nil {
		return nil, false
	}
	return o.AlreadyBoundBindFailureCount, true
}

// HasAlreadyBoundBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasAlreadyBoundBindFailureCount() bool {
	if o != nil && o.AlreadyBoundBindFailureCount != nil {
		return true
	}

	return false
}

// SetAlreadyBoundBindFailureCount gets a reference to the given int64 and assigns it to the AlreadyBoundBindFailureCount field.
func (o *MsgVpnTopicEndpoint) SetAlreadyBoundBindFailureCount(v int64) {
	o.AlreadyBoundBindFailureCount = &v
}

// GetAverageRxByteRate returns the AverageRxByteRate field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetAverageRxByteRate() int64 {
	if o == nil || o.AverageRxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRxByteRate
}

// GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetAverageRxByteRateOk() (*int64, bool) {
	if o == nil || o.AverageRxByteRate == nil {
		return nil, false
	}
	return o.AverageRxByteRate, true
}

// HasAverageRxByteRate returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasAverageRxByteRate() bool {
	if o != nil && o.AverageRxByteRate != nil {
		return true
	}

	return false
}

// SetAverageRxByteRate gets a reference to the given int64 and assigns it to the AverageRxByteRate field.
func (o *MsgVpnTopicEndpoint) SetAverageRxByteRate(v int64) {
	o.AverageRxByteRate = &v
}

// GetAverageRxMsgRate returns the AverageRxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetAverageRxMsgRate() int64 {
	if o == nil || o.AverageRxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageRxMsgRate
}

// GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetAverageRxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageRxMsgRate == nil {
		return nil, false
	}
	return o.AverageRxMsgRate, true
}

// HasAverageRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasAverageRxMsgRate() bool {
	if o != nil && o.AverageRxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageRxMsgRate gets a reference to the given int64 and assigns it to the AverageRxMsgRate field.
func (o *MsgVpnTopicEndpoint) SetAverageRxMsgRate(v int64) {
	o.AverageRxMsgRate = &v
}

// GetAverageTxByteRate returns the AverageTxByteRate field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetAverageTxByteRate() int64 {
	if o == nil || o.AverageTxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageTxByteRate
}

// GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetAverageTxByteRateOk() (*int64, bool) {
	if o == nil || o.AverageTxByteRate == nil {
		return nil, false
	}
	return o.AverageTxByteRate, true
}

// HasAverageTxByteRate returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasAverageTxByteRate() bool {
	if o != nil && o.AverageTxByteRate != nil {
		return true
	}

	return false
}

// SetAverageTxByteRate gets a reference to the given int64 and assigns it to the AverageTxByteRate field.
func (o *MsgVpnTopicEndpoint) SetAverageTxByteRate(v int64) {
	o.AverageTxByteRate = &v
}

// GetAverageTxMsgRate returns the AverageTxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetAverageTxMsgRate() int64 {
	if o == nil || o.AverageTxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.AverageTxMsgRate
}

// GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetAverageTxMsgRateOk() (*int64, bool) {
	if o == nil || o.AverageTxMsgRate == nil {
		return nil, false
	}
	return o.AverageTxMsgRate, true
}

// HasAverageTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasAverageTxMsgRate() bool {
	if o != nil && o.AverageTxMsgRate != nil {
		return true
	}

	return false
}

// SetAverageTxMsgRate gets a reference to the given int64 and assigns it to the AverageTxMsgRate field.
func (o *MsgVpnTopicEndpoint) SetAverageTxMsgRate(v int64) {
	o.AverageTxMsgRate = &v
}

// GetBindRequestCount returns the BindRequestCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetBindRequestCount() int64 {
	if o == nil || o.BindRequestCount == nil {
		var ret int64
		return ret
	}
	return *o.BindRequestCount
}

// GetBindRequestCountOk returns a tuple with the BindRequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetBindRequestCountOk() (*int64, bool) {
	if o == nil || o.BindRequestCount == nil {
		return nil, false
	}
	return o.BindRequestCount, true
}

// HasBindRequestCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasBindRequestCount() bool {
	if o != nil && o.BindRequestCount != nil {
		return true
	}

	return false
}

// SetBindRequestCount gets a reference to the given int64 and assigns it to the BindRequestCount field.
func (o *MsgVpnTopicEndpoint) SetBindRequestCount(v int64) {
	o.BindRequestCount = &v
}

// GetBindSuccessCount returns the BindSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetBindSuccessCount() int64 {
	if o == nil || o.BindSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.BindSuccessCount
}

// GetBindSuccessCountOk returns a tuple with the BindSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetBindSuccessCountOk() (*int64, bool) {
	if o == nil || o.BindSuccessCount == nil {
		return nil, false
	}
	return o.BindSuccessCount, true
}

// HasBindSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasBindSuccessCount() bool {
	if o != nil && o.BindSuccessCount != nil {
		return true
	}

	return false
}

// SetBindSuccessCount gets a reference to the given int64 and assigns it to the BindSuccessCount field.
func (o *MsgVpnTopicEndpoint) SetBindSuccessCount(v int64) {
	o.BindSuccessCount = &v
}

// GetBindTimeForwardingMode returns the BindTimeForwardingMode field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetBindTimeForwardingMode() string {
	if o == nil || o.BindTimeForwardingMode == nil {
		var ret string
		return ret
	}
	return *o.BindTimeForwardingMode
}

// GetBindTimeForwardingModeOk returns a tuple with the BindTimeForwardingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetBindTimeForwardingModeOk() (*string, bool) {
	if o == nil || o.BindTimeForwardingMode == nil {
		return nil, false
	}
	return o.BindTimeForwardingMode, true
}

// HasBindTimeForwardingMode returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasBindTimeForwardingMode() bool {
	if o != nil && o.BindTimeForwardingMode != nil {
		return true
	}

	return false
}

// SetBindTimeForwardingMode gets a reference to the given string and assigns it to the BindTimeForwardingMode field.
func (o *MsgVpnTopicEndpoint) SetBindTimeForwardingMode(v string) {
	o.BindTimeForwardingMode = &v
}

// GetClientProfileDeniedDiscardedMsgCount returns the ClientProfileDeniedDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetClientProfileDeniedDiscardedMsgCount() int64 {
	if o == nil || o.ClientProfileDeniedDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ClientProfileDeniedDiscardedMsgCount
}

// GetClientProfileDeniedDiscardedMsgCountOk returns a tuple with the ClientProfileDeniedDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetClientProfileDeniedDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.ClientProfileDeniedDiscardedMsgCount == nil {
		return nil, false
	}
	return o.ClientProfileDeniedDiscardedMsgCount, true
}

// HasClientProfileDeniedDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasClientProfileDeniedDiscardedMsgCount() bool {
	if o != nil && o.ClientProfileDeniedDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetClientProfileDeniedDiscardedMsgCount gets a reference to the given int64 and assigns it to the ClientProfileDeniedDiscardedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetClientProfileDeniedDiscardedMsgCount(v int64) {
	o.ClientProfileDeniedDiscardedMsgCount = &v
}

// GetConsumerAckPropagationEnabled returns the ConsumerAckPropagationEnabled field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetConsumerAckPropagationEnabled() bool {
	if o == nil || o.ConsumerAckPropagationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ConsumerAckPropagationEnabled
}

// GetConsumerAckPropagationEnabledOk returns a tuple with the ConsumerAckPropagationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetConsumerAckPropagationEnabledOk() (*bool, bool) {
	if o == nil || o.ConsumerAckPropagationEnabled == nil {
		return nil, false
	}
	return o.ConsumerAckPropagationEnabled, true
}

// HasConsumerAckPropagationEnabled returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasConsumerAckPropagationEnabled() bool {
	if o != nil && o.ConsumerAckPropagationEnabled != nil {
		return true
	}

	return false
}

// SetConsumerAckPropagationEnabled gets a reference to the given bool and assigns it to the ConsumerAckPropagationEnabled field.
func (o *MsgVpnTopicEndpoint) SetConsumerAckPropagationEnabled(v bool) {
	o.ConsumerAckPropagationEnabled = &v
}

// GetCreatedByManagement returns the CreatedByManagement field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetCreatedByManagement() bool {
	if o == nil || o.CreatedByManagement == nil {
		var ret bool
		return ret
	}
	return *o.CreatedByManagement
}

// GetCreatedByManagementOk returns a tuple with the CreatedByManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetCreatedByManagementOk() (*bool, bool) {
	if o == nil || o.CreatedByManagement == nil {
		return nil, false
	}
	return o.CreatedByManagement, true
}

// HasCreatedByManagement returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasCreatedByManagement() bool {
	if o != nil && o.CreatedByManagement != nil {
		return true
	}

	return false
}

// SetCreatedByManagement gets a reference to the given bool and assigns it to the CreatedByManagement field.
func (o *MsgVpnTopicEndpoint) SetCreatedByManagement(v bool) {
	o.CreatedByManagement = &v
}

// GetDeadMsgQueue returns the DeadMsgQueue field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetDeadMsgQueue() string {
	if o == nil || o.DeadMsgQueue == nil {
		var ret string
		return ret
	}
	return *o.DeadMsgQueue
}

// GetDeadMsgQueueOk returns a tuple with the DeadMsgQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetDeadMsgQueueOk() (*string, bool) {
	if o == nil || o.DeadMsgQueue == nil {
		return nil, false
	}
	return o.DeadMsgQueue, true
}

// HasDeadMsgQueue returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasDeadMsgQueue() bool {
	if o != nil && o.DeadMsgQueue != nil {
		return true
	}

	return false
}

// SetDeadMsgQueue gets a reference to the given string and assigns it to the DeadMsgQueue field.
func (o *MsgVpnTopicEndpoint) SetDeadMsgQueue(v string) {
	o.DeadMsgQueue = &v
}

// GetDeletedMsgCount returns the DeletedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetDeletedMsgCount() int64 {
	if o == nil || o.DeletedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DeletedMsgCount
}

// GetDeletedMsgCountOk returns a tuple with the DeletedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetDeletedMsgCountOk() (*int64, bool) {
	if o == nil || o.DeletedMsgCount == nil {
		return nil, false
	}
	return o.DeletedMsgCount, true
}

// HasDeletedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasDeletedMsgCount() bool {
	if o != nil && o.DeletedMsgCount != nil {
		return true
	}

	return false
}

// SetDeletedMsgCount gets a reference to the given int64 and assigns it to the DeletedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetDeletedMsgCount(v int64) {
	o.DeletedMsgCount = &v
}

// GetDeliveryCountEnabled returns the DeliveryCountEnabled field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetDeliveryCountEnabled() bool {
	if o == nil || o.DeliveryCountEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DeliveryCountEnabled
}

// GetDeliveryCountEnabledOk returns a tuple with the DeliveryCountEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetDeliveryCountEnabledOk() (*bool, bool) {
	if o == nil || o.DeliveryCountEnabled == nil {
		return nil, false
	}
	return o.DeliveryCountEnabled, true
}

// HasDeliveryCountEnabled returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasDeliveryCountEnabled() bool {
	if o != nil && o.DeliveryCountEnabled != nil {
		return true
	}

	return false
}

// SetDeliveryCountEnabled gets a reference to the given bool and assigns it to the DeliveryCountEnabled field.
func (o *MsgVpnTopicEndpoint) SetDeliveryCountEnabled(v bool) {
	o.DeliveryCountEnabled = &v
}

// GetDestinationGroupErrorDiscardedMsgCount returns the DestinationGroupErrorDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetDestinationGroupErrorDiscardedMsgCount() int64 {
	if o == nil || o.DestinationGroupErrorDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DestinationGroupErrorDiscardedMsgCount
}

// GetDestinationGroupErrorDiscardedMsgCountOk returns a tuple with the DestinationGroupErrorDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetDestinationGroupErrorDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.DestinationGroupErrorDiscardedMsgCount == nil {
		return nil, false
	}
	return o.DestinationGroupErrorDiscardedMsgCount, true
}

// HasDestinationGroupErrorDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasDestinationGroupErrorDiscardedMsgCount() bool {
	if o != nil && o.DestinationGroupErrorDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetDestinationGroupErrorDiscardedMsgCount gets a reference to the given int64 and assigns it to the DestinationGroupErrorDiscardedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetDestinationGroupErrorDiscardedMsgCount(v int64) {
	o.DestinationGroupErrorDiscardedMsgCount = &v
}

// GetDestinationTopic returns the DestinationTopic field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetDestinationTopic() string {
	if o == nil || o.DestinationTopic == nil {
		var ret string
		return ret
	}
	return *o.DestinationTopic
}

// GetDestinationTopicOk returns a tuple with the DestinationTopic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetDestinationTopicOk() (*string, bool) {
	if o == nil || o.DestinationTopic == nil {
		return nil, false
	}
	return o.DestinationTopic, true
}

// HasDestinationTopic returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasDestinationTopic() bool {
	if o != nil && o.DestinationTopic != nil {
		return true
	}

	return false
}

// SetDestinationTopic gets a reference to the given string and assigns it to the DestinationTopic field.
func (o *MsgVpnTopicEndpoint) SetDestinationTopic(v string) {
	o.DestinationTopic = &v
}

// GetDisabledBindFailureCount returns the DisabledBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetDisabledBindFailureCount() int64 {
	if o == nil || o.DisabledBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.DisabledBindFailureCount
}

// GetDisabledBindFailureCountOk returns a tuple with the DisabledBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetDisabledBindFailureCountOk() (*int64, bool) {
	if o == nil || o.DisabledBindFailureCount == nil {
		return nil, false
	}
	return o.DisabledBindFailureCount, true
}

// HasDisabledBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasDisabledBindFailureCount() bool {
	if o != nil && o.DisabledBindFailureCount != nil {
		return true
	}

	return false
}

// SetDisabledBindFailureCount gets a reference to the given int64 and assigns it to the DisabledBindFailureCount field.
func (o *MsgVpnTopicEndpoint) SetDisabledBindFailureCount(v int64) {
	o.DisabledBindFailureCount = &v
}

// GetDisabledDiscardedMsgCount returns the DisabledDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetDisabledDiscardedMsgCount() int64 {
	if o == nil || o.DisabledDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.DisabledDiscardedMsgCount
}

// GetDisabledDiscardedMsgCountOk returns a tuple with the DisabledDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetDisabledDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.DisabledDiscardedMsgCount == nil {
		return nil, false
	}
	return o.DisabledDiscardedMsgCount, true
}

// HasDisabledDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasDisabledDiscardedMsgCount() bool {
	if o != nil && o.DisabledDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetDisabledDiscardedMsgCount gets a reference to the given int64 and assigns it to the DisabledDiscardedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetDisabledDiscardedMsgCount(v int64) {
	o.DisabledDiscardedMsgCount = &v
}

// GetDurable returns the Durable field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetDurable() bool {
	if o == nil || o.Durable == nil {
		var ret bool
		return ret
	}
	return *o.Durable
}

// GetDurableOk returns a tuple with the Durable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetDurableOk() (*bool, bool) {
	if o == nil || o.Durable == nil {
		return nil, false
	}
	return o.Durable, true
}

// HasDurable returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasDurable() bool {
	if o != nil && o.Durable != nil {
		return true
	}

	return false
}

// SetDurable gets a reference to the given bool and assigns it to the Durable field.
func (o *MsgVpnTopicEndpoint) SetDurable(v bool) {
	o.Durable = &v
}

// GetEgressEnabled returns the EgressEnabled field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetEgressEnabled() bool {
	if o == nil || o.EgressEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EgressEnabled
}

// GetEgressEnabledOk returns a tuple with the EgressEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetEgressEnabledOk() (*bool, bool) {
	if o == nil || o.EgressEnabled == nil {
		return nil, false
	}
	return o.EgressEnabled, true
}

// HasEgressEnabled returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasEgressEnabled() bool {
	if o != nil && o.EgressEnabled != nil {
		return true
	}

	return false
}

// SetEgressEnabled gets a reference to the given bool and assigns it to the EgressEnabled field.
func (o *MsgVpnTopicEndpoint) SetEgressEnabled(v bool) {
	o.EgressEnabled = &v
}

// GetEventBindCountThreshold returns the EventBindCountThreshold field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetEventBindCountThreshold() EventThreshold {
	if o == nil || o.EventBindCountThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventBindCountThreshold
}

// GetEventBindCountThresholdOk returns a tuple with the EventBindCountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetEventBindCountThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventBindCountThreshold == nil {
		return nil, false
	}
	return o.EventBindCountThreshold, true
}

// HasEventBindCountThreshold returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasEventBindCountThreshold() bool {
	if o != nil && o.EventBindCountThreshold != nil {
		return true
	}

	return false
}

// SetEventBindCountThreshold gets a reference to the given EventThreshold and assigns it to the EventBindCountThreshold field.
func (o *MsgVpnTopicEndpoint) SetEventBindCountThreshold(v EventThreshold) {
	o.EventBindCountThreshold = &v
}

// GetEventRejectLowPriorityMsgLimitThreshold returns the EventRejectLowPriorityMsgLimitThreshold field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetEventRejectLowPriorityMsgLimitThreshold() EventThreshold {
	if o == nil || o.EventRejectLowPriorityMsgLimitThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventRejectLowPriorityMsgLimitThreshold
}

// GetEventRejectLowPriorityMsgLimitThresholdOk returns a tuple with the EventRejectLowPriorityMsgLimitThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetEventRejectLowPriorityMsgLimitThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventRejectLowPriorityMsgLimitThreshold == nil {
		return nil, false
	}
	return o.EventRejectLowPriorityMsgLimitThreshold, true
}

// HasEventRejectLowPriorityMsgLimitThreshold returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasEventRejectLowPriorityMsgLimitThreshold() bool {
	if o != nil && o.EventRejectLowPriorityMsgLimitThreshold != nil {
		return true
	}

	return false
}

// SetEventRejectLowPriorityMsgLimitThreshold gets a reference to the given EventThreshold and assigns it to the EventRejectLowPriorityMsgLimitThreshold field.
func (o *MsgVpnTopicEndpoint) SetEventRejectLowPriorityMsgLimitThreshold(v EventThreshold) {
	o.EventRejectLowPriorityMsgLimitThreshold = &v
}

// GetEventSpoolUsageThreshold returns the EventSpoolUsageThreshold field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetEventSpoolUsageThreshold() EventThreshold {
	if o == nil || o.EventSpoolUsageThreshold == nil {
		var ret EventThreshold
		return ret
	}
	return *o.EventSpoolUsageThreshold
}

// GetEventSpoolUsageThresholdOk returns a tuple with the EventSpoolUsageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetEventSpoolUsageThresholdOk() (*EventThreshold, bool) {
	if o == nil || o.EventSpoolUsageThreshold == nil {
		return nil, false
	}
	return o.EventSpoolUsageThreshold, true
}

// HasEventSpoolUsageThreshold returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasEventSpoolUsageThreshold() bool {
	if o != nil && o.EventSpoolUsageThreshold != nil {
		return true
	}

	return false
}

// SetEventSpoolUsageThreshold gets a reference to the given EventThreshold and assigns it to the EventSpoolUsageThreshold field.
func (o *MsgVpnTopicEndpoint) SetEventSpoolUsageThreshold(v EventThreshold) {
	o.EventSpoolUsageThreshold = &v
}

// GetHighestAckedMsgId returns the HighestAckedMsgId field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetHighestAckedMsgId() int64 {
	if o == nil || o.HighestAckedMsgId == nil {
		var ret int64
		return ret
	}
	return *o.HighestAckedMsgId
}

// GetHighestAckedMsgIdOk returns a tuple with the HighestAckedMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetHighestAckedMsgIdOk() (*int64, bool) {
	if o == nil || o.HighestAckedMsgId == nil {
		return nil, false
	}
	return o.HighestAckedMsgId, true
}

// HasHighestAckedMsgId returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasHighestAckedMsgId() bool {
	if o != nil && o.HighestAckedMsgId != nil {
		return true
	}

	return false
}

// SetHighestAckedMsgId gets a reference to the given int64 and assigns it to the HighestAckedMsgId field.
func (o *MsgVpnTopicEndpoint) SetHighestAckedMsgId(v int64) {
	o.HighestAckedMsgId = &v
}

// GetHighestMsgId returns the HighestMsgId field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetHighestMsgId() int64 {
	if o == nil || o.HighestMsgId == nil {
		var ret int64
		return ret
	}
	return *o.HighestMsgId
}

// GetHighestMsgIdOk returns a tuple with the HighestMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetHighestMsgIdOk() (*int64, bool) {
	if o == nil || o.HighestMsgId == nil {
		return nil, false
	}
	return o.HighestMsgId, true
}

// HasHighestMsgId returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasHighestMsgId() bool {
	if o != nil && o.HighestMsgId != nil {
		return true
	}

	return false
}

// SetHighestMsgId gets a reference to the given int64 and assigns it to the HighestMsgId field.
func (o *MsgVpnTopicEndpoint) SetHighestMsgId(v int64) {
	o.HighestMsgId = &v
}

// GetInProgressAckMsgCount returns the InProgressAckMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetInProgressAckMsgCount() int64 {
	if o == nil || o.InProgressAckMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.InProgressAckMsgCount
}

// GetInProgressAckMsgCountOk returns a tuple with the InProgressAckMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetInProgressAckMsgCountOk() (*int64, bool) {
	if o == nil || o.InProgressAckMsgCount == nil {
		return nil, false
	}
	return o.InProgressAckMsgCount, true
}

// HasInProgressAckMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasInProgressAckMsgCount() bool {
	if o != nil && o.InProgressAckMsgCount != nil {
		return true
	}

	return false
}

// SetInProgressAckMsgCount gets a reference to the given int64 and assigns it to the InProgressAckMsgCount field.
func (o *MsgVpnTopicEndpoint) SetInProgressAckMsgCount(v int64) {
	o.InProgressAckMsgCount = &v
}

// GetIngressEnabled returns the IngressEnabled field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetIngressEnabled() bool {
	if o == nil || o.IngressEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IngressEnabled
}

// GetIngressEnabledOk returns a tuple with the IngressEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetIngressEnabledOk() (*bool, bool) {
	if o == nil || o.IngressEnabled == nil {
		return nil, false
	}
	return o.IngressEnabled, true
}

// HasIngressEnabled returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasIngressEnabled() bool {
	if o != nil && o.IngressEnabled != nil {
		return true
	}

	return false
}

// SetIngressEnabled gets a reference to the given bool and assigns it to the IngressEnabled field.
func (o *MsgVpnTopicEndpoint) SetIngressEnabled(v bool) {
	o.IngressEnabled = &v
}

// GetInvalidSelectorBindFailureCount returns the InvalidSelectorBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetInvalidSelectorBindFailureCount() int64 {
	if o == nil || o.InvalidSelectorBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.InvalidSelectorBindFailureCount
}

// GetInvalidSelectorBindFailureCountOk returns a tuple with the InvalidSelectorBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetInvalidSelectorBindFailureCountOk() (*int64, bool) {
	if o == nil || o.InvalidSelectorBindFailureCount == nil {
		return nil, false
	}
	return o.InvalidSelectorBindFailureCount, true
}

// HasInvalidSelectorBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasInvalidSelectorBindFailureCount() bool {
	if o != nil && o.InvalidSelectorBindFailureCount != nil {
		return true
	}

	return false
}

// SetInvalidSelectorBindFailureCount gets a reference to the given int64 and assigns it to the InvalidSelectorBindFailureCount field.
func (o *MsgVpnTopicEndpoint) SetInvalidSelectorBindFailureCount(v int64) {
	o.InvalidSelectorBindFailureCount = &v
}

// GetLastReplayCompleteTime returns the LastReplayCompleteTime field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLastReplayCompleteTime() int32 {
	if o == nil || o.LastReplayCompleteTime == nil {
		var ret int32
		return ret
	}
	return *o.LastReplayCompleteTime
}

// GetLastReplayCompleteTimeOk returns a tuple with the LastReplayCompleteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLastReplayCompleteTimeOk() (*int32, bool) {
	if o == nil || o.LastReplayCompleteTime == nil {
		return nil, false
	}
	return o.LastReplayCompleteTime, true
}

// HasLastReplayCompleteTime returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLastReplayCompleteTime() bool {
	if o != nil && o.LastReplayCompleteTime != nil {
		return true
	}

	return false
}

// SetLastReplayCompleteTime gets a reference to the given int32 and assigns it to the LastReplayCompleteTime field.
func (o *MsgVpnTopicEndpoint) SetLastReplayCompleteTime(v int32) {
	o.LastReplayCompleteTime = &v
}

// GetLastReplayFailureReason returns the LastReplayFailureReason field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLastReplayFailureReason() string {
	if o == nil || o.LastReplayFailureReason == nil {
		var ret string
		return ret
	}
	return *o.LastReplayFailureReason
}

// GetLastReplayFailureReasonOk returns a tuple with the LastReplayFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLastReplayFailureReasonOk() (*string, bool) {
	if o == nil || o.LastReplayFailureReason == nil {
		return nil, false
	}
	return o.LastReplayFailureReason, true
}

// HasLastReplayFailureReason returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLastReplayFailureReason() bool {
	if o != nil && o.LastReplayFailureReason != nil {
		return true
	}

	return false
}

// SetLastReplayFailureReason gets a reference to the given string and assigns it to the LastReplayFailureReason field.
func (o *MsgVpnTopicEndpoint) SetLastReplayFailureReason(v string) {
	o.LastReplayFailureReason = &v
}

// GetLastReplayFailureTime returns the LastReplayFailureTime field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLastReplayFailureTime() int32 {
	if o == nil || o.LastReplayFailureTime == nil {
		var ret int32
		return ret
	}
	return *o.LastReplayFailureTime
}

// GetLastReplayFailureTimeOk returns a tuple with the LastReplayFailureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLastReplayFailureTimeOk() (*int32, bool) {
	if o == nil || o.LastReplayFailureTime == nil {
		return nil, false
	}
	return o.LastReplayFailureTime, true
}

// HasLastReplayFailureTime returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLastReplayFailureTime() bool {
	if o != nil && o.LastReplayFailureTime != nil {
		return true
	}

	return false
}

// SetLastReplayFailureTime gets a reference to the given int32 and assigns it to the LastReplayFailureTime field.
func (o *MsgVpnTopicEndpoint) SetLastReplayFailureTime(v int32) {
	o.LastReplayFailureTime = &v
}

// GetLastReplayStartTime returns the LastReplayStartTime field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLastReplayStartTime() int32 {
	if o == nil || o.LastReplayStartTime == nil {
		var ret int32
		return ret
	}
	return *o.LastReplayStartTime
}

// GetLastReplayStartTimeOk returns a tuple with the LastReplayStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLastReplayStartTimeOk() (*int32, bool) {
	if o == nil || o.LastReplayStartTime == nil {
		return nil, false
	}
	return o.LastReplayStartTime, true
}

// HasLastReplayStartTime returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLastReplayStartTime() bool {
	if o != nil && o.LastReplayStartTime != nil {
		return true
	}

	return false
}

// SetLastReplayStartTime gets a reference to the given int32 and assigns it to the LastReplayStartTime field.
func (o *MsgVpnTopicEndpoint) SetLastReplayStartTime(v int32) {
	o.LastReplayStartTime = &v
}

// GetLastReplayedMsgTxTime returns the LastReplayedMsgTxTime field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLastReplayedMsgTxTime() int32 {
	if o == nil || o.LastReplayedMsgTxTime == nil {
		var ret int32
		return ret
	}
	return *o.LastReplayedMsgTxTime
}

// GetLastReplayedMsgTxTimeOk returns a tuple with the LastReplayedMsgTxTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLastReplayedMsgTxTimeOk() (*int32, bool) {
	if o == nil || o.LastReplayedMsgTxTime == nil {
		return nil, false
	}
	return o.LastReplayedMsgTxTime, true
}

// HasLastReplayedMsgTxTime returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLastReplayedMsgTxTime() bool {
	if o != nil && o.LastReplayedMsgTxTime != nil {
		return true
	}

	return false
}

// SetLastReplayedMsgTxTime gets a reference to the given int32 and assigns it to the LastReplayedMsgTxTime field.
func (o *MsgVpnTopicEndpoint) SetLastReplayedMsgTxTime(v int32) {
	o.LastReplayedMsgTxTime = &v
}

// GetLastSelectorExaminedMsgId returns the LastSelectorExaminedMsgId field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLastSelectorExaminedMsgId() int64 {
	if o == nil || o.LastSelectorExaminedMsgId == nil {
		var ret int64
		return ret
	}
	return *o.LastSelectorExaminedMsgId
}

// GetLastSelectorExaminedMsgIdOk returns a tuple with the LastSelectorExaminedMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLastSelectorExaminedMsgIdOk() (*int64, bool) {
	if o == nil || o.LastSelectorExaminedMsgId == nil {
		return nil, false
	}
	return o.LastSelectorExaminedMsgId, true
}

// HasLastSelectorExaminedMsgId returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLastSelectorExaminedMsgId() bool {
	if o != nil && o.LastSelectorExaminedMsgId != nil {
		return true
	}

	return false
}

// SetLastSelectorExaminedMsgId gets a reference to the given int64 and assigns it to the LastSelectorExaminedMsgId field.
func (o *MsgVpnTopicEndpoint) SetLastSelectorExaminedMsgId(v int64) {
	o.LastSelectorExaminedMsgId = &v
}

// GetLastSpooledMsgId returns the LastSpooledMsgId field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLastSpooledMsgId() int64 {
	if o == nil || o.LastSpooledMsgId == nil {
		var ret int64
		return ret
	}
	return *o.LastSpooledMsgId
}

// GetLastSpooledMsgIdOk returns a tuple with the LastSpooledMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLastSpooledMsgIdOk() (*int64, bool) {
	if o == nil || o.LastSpooledMsgId == nil {
		return nil, false
	}
	return o.LastSpooledMsgId, true
}

// HasLastSpooledMsgId returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLastSpooledMsgId() bool {
	if o != nil && o.LastSpooledMsgId != nil {
		return true
	}

	return false
}

// SetLastSpooledMsgId gets a reference to the given int64 and assigns it to the LastSpooledMsgId field.
func (o *MsgVpnTopicEndpoint) SetLastSpooledMsgId(v int64) {
	o.LastSpooledMsgId = &v
}

// GetLowPriorityMsgCongestionDiscardedMsgCount returns the LowPriorityMsgCongestionDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLowPriorityMsgCongestionDiscardedMsgCount() int64 {
	if o == nil || o.LowPriorityMsgCongestionDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.LowPriorityMsgCongestionDiscardedMsgCount
}

// GetLowPriorityMsgCongestionDiscardedMsgCountOk returns a tuple with the LowPriorityMsgCongestionDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLowPriorityMsgCongestionDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.LowPriorityMsgCongestionDiscardedMsgCount == nil {
		return nil, false
	}
	return o.LowPriorityMsgCongestionDiscardedMsgCount, true
}

// HasLowPriorityMsgCongestionDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLowPriorityMsgCongestionDiscardedMsgCount() bool {
	if o != nil && o.LowPriorityMsgCongestionDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetLowPriorityMsgCongestionDiscardedMsgCount gets a reference to the given int64 and assigns it to the LowPriorityMsgCongestionDiscardedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetLowPriorityMsgCongestionDiscardedMsgCount(v int64) {
	o.LowPriorityMsgCongestionDiscardedMsgCount = &v
}

// GetLowPriorityMsgCongestionState returns the LowPriorityMsgCongestionState field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLowPriorityMsgCongestionState() string {
	if o == nil || o.LowPriorityMsgCongestionState == nil {
		var ret string
		return ret
	}
	return *o.LowPriorityMsgCongestionState
}

// GetLowPriorityMsgCongestionStateOk returns a tuple with the LowPriorityMsgCongestionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLowPriorityMsgCongestionStateOk() (*string, bool) {
	if o == nil || o.LowPriorityMsgCongestionState == nil {
		return nil, false
	}
	return o.LowPriorityMsgCongestionState, true
}

// HasLowPriorityMsgCongestionState returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLowPriorityMsgCongestionState() bool {
	if o != nil && o.LowPriorityMsgCongestionState != nil {
		return true
	}

	return false
}

// SetLowPriorityMsgCongestionState gets a reference to the given string and assigns it to the LowPriorityMsgCongestionState field.
func (o *MsgVpnTopicEndpoint) SetLowPriorityMsgCongestionState(v string) {
	o.LowPriorityMsgCongestionState = &v
}

// GetLowestAckedMsgId returns the LowestAckedMsgId field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLowestAckedMsgId() int64 {
	if o == nil || o.LowestAckedMsgId == nil {
		var ret int64
		return ret
	}
	return *o.LowestAckedMsgId
}

// GetLowestAckedMsgIdOk returns a tuple with the LowestAckedMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLowestAckedMsgIdOk() (*int64, bool) {
	if o == nil || o.LowestAckedMsgId == nil {
		return nil, false
	}
	return o.LowestAckedMsgId, true
}

// HasLowestAckedMsgId returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLowestAckedMsgId() bool {
	if o != nil && o.LowestAckedMsgId != nil {
		return true
	}

	return false
}

// SetLowestAckedMsgId gets a reference to the given int64 and assigns it to the LowestAckedMsgId field.
func (o *MsgVpnTopicEndpoint) SetLowestAckedMsgId(v int64) {
	o.LowestAckedMsgId = &v
}

// GetLowestMsgId returns the LowestMsgId field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetLowestMsgId() int64 {
	if o == nil || o.LowestMsgId == nil {
		var ret int64
		return ret
	}
	return *o.LowestMsgId
}

// GetLowestMsgIdOk returns a tuple with the LowestMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetLowestMsgIdOk() (*int64, bool) {
	if o == nil || o.LowestMsgId == nil {
		return nil, false
	}
	return o.LowestMsgId, true
}

// HasLowestMsgId returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasLowestMsgId() bool {
	if o != nil && o.LowestMsgId != nil {
		return true
	}

	return false
}

// SetLowestMsgId gets a reference to the given int64 and assigns it to the LowestMsgId field.
func (o *MsgVpnTopicEndpoint) SetLowestMsgId(v int64) {
	o.LowestMsgId = &v
}

// GetMaxBindCount returns the MaxBindCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxBindCount() int64 {
	if o == nil || o.MaxBindCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxBindCount
}

// GetMaxBindCountOk returns a tuple with the MaxBindCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxBindCountOk() (*int64, bool) {
	if o == nil || o.MaxBindCount == nil {
		return nil, false
	}
	return o.MaxBindCount, true
}

// HasMaxBindCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxBindCount() bool {
	if o != nil && o.MaxBindCount != nil {
		return true
	}

	return false
}

// SetMaxBindCount gets a reference to the given int64 and assigns it to the MaxBindCount field.
func (o *MsgVpnTopicEndpoint) SetMaxBindCount(v int64) {
	o.MaxBindCount = &v
}

// GetMaxBindCountExceededBindFailureCount returns the MaxBindCountExceededBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxBindCountExceededBindFailureCount() int64 {
	if o == nil || o.MaxBindCountExceededBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxBindCountExceededBindFailureCount
}

// GetMaxBindCountExceededBindFailureCountOk returns a tuple with the MaxBindCountExceededBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxBindCountExceededBindFailureCountOk() (*int64, bool) {
	if o == nil || o.MaxBindCountExceededBindFailureCount == nil {
		return nil, false
	}
	return o.MaxBindCountExceededBindFailureCount, true
}

// HasMaxBindCountExceededBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxBindCountExceededBindFailureCount() bool {
	if o != nil && o.MaxBindCountExceededBindFailureCount != nil {
		return true
	}

	return false
}

// SetMaxBindCountExceededBindFailureCount gets a reference to the given int64 and assigns it to the MaxBindCountExceededBindFailureCount field.
func (o *MsgVpnTopicEndpoint) SetMaxBindCountExceededBindFailureCount(v int64) {
	o.MaxBindCountExceededBindFailureCount = &v
}

// GetMaxDeliveredUnackedMsgsPerFlow returns the MaxDeliveredUnackedMsgsPerFlow field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxDeliveredUnackedMsgsPerFlow() int64 {
	if o == nil || o.MaxDeliveredUnackedMsgsPerFlow == nil {
		var ret int64
		return ret
	}
	return *o.MaxDeliveredUnackedMsgsPerFlow
}

// GetMaxDeliveredUnackedMsgsPerFlowOk returns a tuple with the MaxDeliveredUnackedMsgsPerFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxDeliveredUnackedMsgsPerFlowOk() (*int64, bool) {
	if o == nil || o.MaxDeliveredUnackedMsgsPerFlow == nil {
		return nil, false
	}
	return o.MaxDeliveredUnackedMsgsPerFlow, true
}

// HasMaxDeliveredUnackedMsgsPerFlow returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxDeliveredUnackedMsgsPerFlow() bool {
	if o != nil && o.MaxDeliveredUnackedMsgsPerFlow != nil {
		return true
	}

	return false
}

// SetMaxDeliveredUnackedMsgsPerFlow gets a reference to the given int64 and assigns it to the MaxDeliveredUnackedMsgsPerFlow field.
func (o *MsgVpnTopicEndpoint) SetMaxDeliveredUnackedMsgsPerFlow(v int64) {
	o.MaxDeliveredUnackedMsgsPerFlow = &v
}

// GetMaxEffectiveBindCount returns the MaxEffectiveBindCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxEffectiveBindCount() int32 {
	if o == nil || o.MaxEffectiveBindCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxEffectiveBindCount
}

// GetMaxEffectiveBindCountOk returns a tuple with the MaxEffectiveBindCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxEffectiveBindCountOk() (*int32, bool) {
	if o == nil || o.MaxEffectiveBindCount == nil {
		return nil, false
	}
	return o.MaxEffectiveBindCount, true
}

// HasMaxEffectiveBindCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxEffectiveBindCount() bool {
	if o != nil && o.MaxEffectiveBindCount != nil {
		return true
	}

	return false
}

// SetMaxEffectiveBindCount gets a reference to the given int32 and assigns it to the MaxEffectiveBindCount field.
func (o *MsgVpnTopicEndpoint) SetMaxEffectiveBindCount(v int32) {
	o.MaxEffectiveBindCount = &v
}

// GetMaxMsgSize returns the MaxMsgSize field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxMsgSize() int32 {
	if o == nil || o.MaxMsgSize == nil {
		var ret int32
		return ret
	}
	return *o.MaxMsgSize
}

// GetMaxMsgSizeOk returns a tuple with the MaxMsgSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxMsgSizeOk() (*int32, bool) {
	if o == nil || o.MaxMsgSize == nil {
		return nil, false
	}
	return o.MaxMsgSize, true
}

// HasMaxMsgSize returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxMsgSize() bool {
	if o != nil && o.MaxMsgSize != nil {
		return true
	}

	return false
}

// SetMaxMsgSize gets a reference to the given int32 and assigns it to the MaxMsgSize field.
func (o *MsgVpnTopicEndpoint) SetMaxMsgSize(v int32) {
	o.MaxMsgSize = &v
}

// GetMaxMsgSizeExceededDiscardedMsgCount returns the MaxMsgSizeExceededDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxMsgSizeExceededDiscardedMsgCount() int64 {
	if o == nil || o.MaxMsgSizeExceededDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxMsgSizeExceededDiscardedMsgCount
}

// GetMaxMsgSizeExceededDiscardedMsgCountOk returns a tuple with the MaxMsgSizeExceededDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxMsgSizeExceededDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxMsgSizeExceededDiscardedMsgCount == nil {
		return nil, false
	}
	return o.MaxMsgSizeExceededDiscardedMsgCount, true
}

// HasMaxMsgSizeExceededDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxMsgSizeExceededDiscardedMsgCount() bool {
	if o != nil && o.MaxMsgSizeExceededDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetMaxMsgSizeExceededDiscardedMsgCount gets a reference to the given int64 and assigns it to the MaxMsgSizeExceededDiscardedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetMaxMsgSizeExceededDiscardedMsgCount(v int64) {
	o.MaxMsgSizeExceededDiscardedMsgCount = &v
}

// GetMaxMsgSpoolUsageExceededDiscardedMsgCount returns the MaxMsgSpoolUsageExceededDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxMsgSpoolUsageExceededDiscardedMsgCount() int64 {
	if o == nil || o.MaxMsgSpoolUsageExceededDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxMsgSpoolUsageExceededDiscardedMsgCount
}

// GetMaxMsgSpoolUsageExceededDiscardedMsgCountOk returns a tuple with the MaxMsgSpoolUsageExceededDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxMsgSpoolUsageExceededDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxMsgSpoolUsageExceededDiscardedMsgCount == nil {
		return nil, false
	}
	return o.MaxMsgSpoolUsageExceededDiscardedMsgCount, true
}

// HasMaxMsgSpoolUsageExceededDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxMsgSpoolUsageExceededDiscardedMsgCount() bool {
	if o != nil && o.MaxMsgSpoolUsageExceededDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetMaxMsgSpoolUsageExceededDiscardedMsgCount gets a reference to the given int64 and assigns it to the MaxMsgSpoolUsageExceededDiscardedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetMaxMsgSpoolUsageExceededDiscardedMsgCount(v int64) {
	o.MaxMsgSpoolUsageExceededDiscardedMsgCount = &v
}

// GetMaxRedeliveryCount returns the MaxRedeliveryCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryCount() int64 {
	if o == nil || o.MaxRedeliveryCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxRedeliveryCount
}

// GetMaxRedeliveryCountOk returns a tuple with the MaxRedeliveryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryCountOk() (*int64, bool) {
	if o == nil || o.MaxRedeliveryCount == nil {
		return nil, false
	}
	return o.MaxRedeliveryCount, true
}

// HasMaxRedeliveryCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxRedeliveryCount() bool {
	if o != nil && o.MaxRedeliveryCount != nil {
		return true
	}

	return false
}

// SetMaxRedeliveryCount gets a reference to the given int64 and assigns it to the MaxRedeliveryCount field.
func (o *MsgVpnTopicEndpoint) SetMaxRedeliveryCount(v int64) {
	o.MaxRedeliveryCount = &v
}

// GetMaxRedeliveryExceededDiscardedMsgCount returns the MaxRedeliveryExceededDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededDiscardedMsgCount() int64 {
	if o == nil || o.MaxRedeliveryExceededDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxRedeliveryExceededDiscardedMsgCount
}

// GetMaxRedeliveryExceededDiscardedMsgCountOk returns a tuple with the MaxRedeliveryExceededDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxRedeliveryExceededDiscardedMsgCount == nil {
		return nil, false
	}
	return o.MaxRedeliveryExceededDiscardedMsgCount, true
}

// HasMaxRedeliveryExceededDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxRedeliveryExceededDiscardedMsgCount() bool {
	if o != nil && o.MaxRedeliveryExceededDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetMaxRedeliveryExceededDiscardedMsgCount gets a reference to the given int64 and assigns it to the MaxRedeliveryExceededDiscardedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetMaxRedeliveryExceededDiscardedMsgCount(v int64) {
	o.MaxRedeliveryExceededDiscardedMsgCount = &v
}

// GetMaxRedeliveryExceededToDmqFailedMsgCount returns the MaxRedeliveryExceededToDmqFailedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededToDmqFailedMsgCount() int64 {
	if o == nil || o.MaxRedeliveryExceededToDmqFailedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxRedeliveryExceededToDmqFailedMsgCount
}

// GetMaxRedeliveryExceededToDmqFailedMsgCountOk returns a tuple with the MaxRedeliveryExceededToDmqFailedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededToDmqFailedMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxRedeliveryExceededToDmqFailedMsgCount == nil {
		return nil, false
	}
	return o.MaxRedeliveryExceededToDmqFailedMsgCount, true
}

// HasMaxRedeliveryExceededToDmqFailedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxRedeliveryExceededToDmqFailedMsgCount() bool {
	if o != nil && o.MaxRedeliveryExceededToDmqFailedMsgCount != nil {
		return true
	}

	return false
}

// SetMaxRedeliveryExceededToDmqFailedMsgCount gets a reference to the given int64 and assigns it to the MaxRedeliveryExceededToDmqFailedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetMaxRedeliveryExceededToDmqFailedMsgCount(v int64) {
	o.MaxRedeliveryExceededToDmqFailedMsgCount = &v
}

// GetMaxRedeliveryExceededToDmqMsgCount returns the MaxRedeliveryExceededToDmqMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededToDmqMsgCount() int64 {
	if o == nil || o.MaxRedeliveryExceededToDmqMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxRedeliveryExceededToDmqMsgCount
}

// GetMaxRedeliveryExceededToDmqMsgCountOk returns a tuple with the MaxRedeliveryExceededToDmqMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxRedeliveryExceededToDmqMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxRedeliveryExceededToDmqMsgCount == nil {
		return nil, false
	}
	return o.MaxRedeliveryExceededToDmqMsgCount, true
}

// HasMaxRedeliveryExceededToDmqMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxRedeliveryExceededToDmqMsgCount() bool {
	if o != nil && o.MaxRedeliveryExceededToDmqMsgCount != nil {
		return true
	}

	return false
}

// SetMaxRedeliveryExceededToDmqMsgCount gets a reference to the given int64 and assigns it to the MaxRedeliveryExceededToDmqMsgCount field.
func (o *MsgVpnTopicEndpoint) SetMaxRedeliveryExceededToDmqMsgCount(v int64) {
	o.MaxRedeliveryExceededToDmqMsgCount = &v
}

// GetMaxSpoolUsage returns the MaxSpoolUsage field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxSpoolUsage() int64 {
	if o == nil || o.MaxSpoolUsage == nil {
		var ret int64
		return ret
	}
	return *o.MaxSpoolUsage
}

// GetMaxSpoolUsageOk returns a tuple with the MaxSpoolUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxSpoolUsageOk() (*int64, bool) {
	if o == nil || o.MaxSpoolUsage == nil {
		return nil, false
	}
	return o.MaxSpoolUsage, true
}

// HasMaxSpoolUsage returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxSpoolUsage() bool {
	if o != nil && o.MaxSpoolUsage != nil {
		return true
	}

	return false
}

// SetMaxSpoolUsage gets a reference to the given int64 and assigns it to the MaxSpoolUsage field.
func (o *MsgVpnTopicEndpoint) SetMaxSpoolUsage(v int64) {
	o.MaxSpoolUsage = &v
}

// GetMaxTtl returns the MaxTtl field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxTtl() int64 {
	if o == nil || o.MaxTtl == nil {
		var ret int64
		return ret
	}
	return *o.MaxTtl
}

// GetMaxTtlOk returns a tuple with the MaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxTtlOk() (*int64, bool) {
	if o == nil || o.MaxTtl == nil {
		return nil, false
	}
	return o.MaxTtl, true
}

// HasMaxTtl returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxTtl() bool {
	if o != nil && o.MaxTtl != nil {
		return true
	}

	return false
}

// SetMaxTtl gets a reference to the given int64 and assigns it to the MaxTtl field.
func (o *MsgVpnTopicEndpoint) SetMaxTtl(v int64) {
	o.MaxTtl = &v
}

// GetMaxTtlExceededDiscardedMsgCount returns the MaxTtlExceededDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxTtlExceededDiscardedMsgCount() int64 {
	if o == nil || o.MaxTtlExceededDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxTtlExceededDiscardedMsgCount
}

// GetMaxTtlExceededDiscardedMsgCountOk returns a tuple with the MaxTtlExceededDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxTtlExceededDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxTtlExceededDiscardedMsgCount == nil {
		return nil, false
	}
	return o.MaxTtlExceededDiscardedMsgCount, true
}

// HasMaxTtlExceededDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxTtlExceededDiscardedMsgCount() bool {
	if o != nil && o.MaxTtlExceededDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetMaxTtlExceededDiscardedMsgCount gets a reference to the given int64 and assigns it to the MaxTtlExceededDiscardedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetMaxTtlExceededDiscardedMsgCount(v int64) {
	o.MaxTtlExceededDiscardedMsgCount = &v
}

// GetMaxTtlExpiredDiscardedMsgCount returns the MaxTtlExpiredDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredDiscardedMsgCount() int64 {
	if o == nil || o.MaxTtlExpiredDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxTtlExpiredDiscardedMsgCount
}

// GetMaxTtlExpiredDiscardedMsgCountOk returns a tuple with the MaxTtlExpiredDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxTtlExpiredDiscardedMsgCount == nil {
		return nil, false
	}
	return o.MaxTtlExpiredDiscardedMsgCount, true
}

// HasMaxTtlExpiredDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxTtlExpiredDiscardedMsgCount() bool {
	if o != nil && o.MaxTtlExpiredDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetMaxTtlExpiredDiscardedMsgCount gets a reference to the given int64 and assigns it to the MaxTtlExpiredDiscardedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetMaxTtlExpiredDiscardedMsgCount(v int64) {
	o.MaxTtlExpiredDiscardedMsgCount = &v
}

// GetMaxTtlExpiredToDmqFailedMsgCount returns the MaxTtlExpiredToDmqFailedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredToDmqFailedMsgCount() int64 {
	if o == nil || o.MaxTtlExpiredToDmqFailedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxTtlExpiredToDmqFailedMsgCount
}

// GetMaxTtlExpiredToDmqFailedMsgCountOk returns a tuple with the MaxTtlExpiredToDmqFailedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredToDmqFailedMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxTtlExpiredToDmqFailedMsgCount == nil {
		return nil, false
	}
	return o.MaxTtlExpiredToDmqFailedMsgCount, true
}

// HasMaxTtlExpiredToDmqFailedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxTtlExpiredToDmqFailedMsgCount() bool {
	if o != nil && o.MaxTtlExpiredToDmqFailedMsgCount != nil {
		return true
	}

	return false
}

// SetMaxTtlExpiredToDmqFailedMsgCount gets a reference to the given int64 and assigns it to the MaxTtlExpiredToDmqFailedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetMaxTtlExpiredToDmqFailedMsgCount(v int64) {
	o.MaxTtlExpiredToDmqFailedMsgCount = &v
}

// GetMaxTtlExpiredToDmqMsgCount returns the MaxTtlExpiredToDmqMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredToDmqMsgCount() int64 {
	if o == nil || o.MaxTtlExpiredToDmqMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxTtlExpiredToDmqMsgCount
}

// GetMaxTtlExpiredToDmqMsgCountOk returns a tuple with the MaxTtlExpiredToDmqMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMaxTtlExpiredToDmqMsgCountOk() (*int64, bool) {
	if o == nil || o.MaxTtlExpiredToDmqMsgCount == nil {
		return nil, false
	}
	return o.MaxTtlExpiredToDmqMsgCount, true
}

// HasMaxTtlExpiredToDmqMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMaxTtlExpiredToDmqMsgCount() bool {
	if o != nil && o.MaxTtlExpiredToDmqMsgCount != nil {
		return true
	}

	return false
}

// SetMaxTtlExpiredToDmqMsgCount gets a reference to the given int64 and assigns it to the MaxTtlExpiredToDmqMsgCount field.
func (o *MsgVpnTopicEndpoint) SetMaxTtlExpiredToDmqMsgCount(v int64) {
	o.MaxTtlExpiredToDmqMsgCount = &v
}

// GetMsgSpoolPeakUsage returns the MsgSpoolPeakUsage field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMsgSpoolPeakUsage() int64 {
	if o == nil || o.MsgSpoolPeakUsage == nil {
		var ret int64
		return ret
	}
	return *o.MsgSpoolPeakUsage
}

// GetMsgSpoolPeakUsageOk returns a tuple with the MsgSpoolPeakUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMsgSpoolPeakUsageOk() (*int64, bool) {
	if o == nil || o.MsgSpoolPeakUsage == nil {
		return nil, false
	}
	return o.MsgSpoolPeakUsage, true
}

// HasMsgSpoolPeakUsage returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMsgSpoolPeakUsage() bool {
	if o != nil && o.MsgSpoolPeakUsage != nil {
		return true
	}

	return false
}

// SetMsgSpoolPeakUsage gets a reference to the given int64 and assigns it to the MsgSpoolPeakUsage field.
func (o *MsgVpnTopicEndpoint) SetMsgSpoolPeakUsage(v int64) {
	o.MsgSpoolPeakUsage = &v
}

// GetMsgSpoolUsage returns the MsgSpoolUsage field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMsgSpoolUsage() int64 {
	if o == nil || o.MsgSpoolUsage == nil {
		var ret int64
		return ret
	}
	return *o.MsgSpoolUsage
}

// GetMsgSpoolUsageOk returns a tuple with the MsgSpoolUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMsgSpoolUsageOk() (*int64, bool) {
	if o == nil || o.MsgSpoolUsage == nil {
		return nil, false
	}
	return o.MsgSpoolUsage, true
}

// HasMsgSpoolUsage returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMsgSpoolUsage() bool {
	if o != nil && o.MsgSpoolUsage != nil {
		return true
	}

	return false
}

// SetMsgSpoolUsage gets a reference to the given int64 and assigns it to the MsgSpoolUsage field.
func (o *MsgVpnTopicEndpoint) SetMsgSpoolUsage(v int64) {
	o.MsgSpoolUsage = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnTopicEndpoint) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetNetworkTopic returns the NetworkTopic field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetNetworkTopic() string {
	if o == nil || o.NetworkTopic == nil {
		var ret string
		return ret
	}
	return *o.NetworkTopic
}

// GetNetworkTopicOk returns a tuple with the NetworkTopic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetNetworkTopicOk() (*string, bool) {
	if o == nil || o.NetworkTopic == nil {
		return nil, false
	}
	return o.NetworkTopic, true
}

// HasNetworkTopic returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasNetworkTopic() bool {
	if o != nil && o.NetworkTopic != nil {
		return true
	}

	return false
}

// SetNetworkTopic gets a reference to the given string and assigns it to the NetworkTopic field.
func (o *MsgVpnTopicEndpoint) SetNetworkTopic(v string) {
	o.NetworkTopic = &v
}

// GetNoLocalDeliveryDiscardedMsgCount returns the NoLocalDeliveryDiscardedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetNoLocalDeliveryDiscardedMsgCount() int64 {
	if o == nil || o.NoLocalDeliveryDiscardedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.NoLocalDeliveryDiscardedMsgCount
}

// GetNoLocalDeliveryDiscardedMsgCountOk returns a tuple with the NoLocalDeliveryDiscardedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetNoLocalDeliveryDiscardedMsgCountOk() (*int64, bool) {
	if o == nil || o.NoLocalDeliveryDiscardedMsgCount == nil {
		return nil, false
	}
	return o.NoLocalDeliveryDiscardedMsgCount, true
}

// HasNoLocalDeliveryDiscardedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasNoLocalDeliveryDiscardedMsgCount() bool {
	if o != nil && o.NoLocalDeliveryDiscardedMsgCount != nil {
		return true
	}

	return false
}

// SetNoLocalDeliveryDiscardedMsgCount gets a reference to the given int64 and assigns it to the NoLocalDeliveryDiscardedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetNoLocalDeliveryDiscardedMsgCount(v int64) {
	o.NoLocalDeliveryDiscardedMsgCount = &v
}

// GetOtherBindFailureCount returns the OtherBindFailureCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetOtherBindFailureCount() int64 {
	if o == nil || o.OtherBindFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.OtherBindFailureCount
}

// GetOtherBindFailureCountOk returns a tuple with the OtherBindFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetOtherBindFailureCountOk() (*int64, bool) {
	if o == nil || o.OtherBindFailureCount == nil {
		return nil, false
	}
	return o.OtherBindFailureCount, true
}

// HasOtherBindFailureCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasOtherBindFailureCount() bool {
	if o != nil && o.OtherBindFailureCount != nil {
		return true
	}

	return false
}

// SetOtherBindFailureCount gets a reference to the given int64 and assigns it to the OtherBindFailureCount field.
func (o *MsgVpnTopicEndpoint) SetOtherBindFailureCount(v int64) {
	o.OtherBindFailureCount = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *MsgVpnTopicEndpoint) SetOwner(v string) {
	o.Owner = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *MsgVpnTopicEndpoint) SetPermission(v string) {
	o.Permission = &v
}

// GetRedeliveredMsgCount returns the RedeliveredMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetRedeliveredMsgCount() int64 {
	if o == nil || o.RedeliveredMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RedeliveredMsgCount
}

// GetRedeliveredMsgCountOk returns a tuple with the RedeliveredMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetRedeliveredMsgCountOk() (*int64, bool) {
	if o == nil || o.RedeliveredMsgCount == nil {
		return nil, false
	}
	return o.RedeliveredMsgCount, true
}

// HasRedeliveredMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasRedeliveredMsgCount() bool {
	if o != nil && o.RedeliveredMsgCount != nil {
		return true
	}

	return false
}

// SetRedeliveredMsgCount gets a reference to the given int64 and assigns it to the RedeliveredMsgCount field.
func (o *MsgVpnTopicEndpoint) SetRedeliveredMsgCount(v int64) {
	o.RedeliveredMsgCount = &v
}

// GetRedeliveryEnabled returns the RedeliveryEnabled field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetRedeliveryEnabled() bool {
	if o == nil || o.RedeliveryEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RedeliveryEnabled
}

// GetRedeliveryEnabledOk returns a tuple with the RedeliveryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetRedeliveryEnabledOk() (*bool, bool) {
	if o == nil || o.RedeliveryEnabled == nil {
		return nil, false
	}
	return o.RedeliveryEnabled, true
}

// HasRedeliveryEnabled returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasRedeliveryEnabled() bool {
	if o != nil && o.RedeliveryEnabled != nil {
		return true
	}

	return false
}

// SetRedeliveryEnabled gets a reference to the given bool and assigns it to the RedeliveryEnabled field.
func (o *MsgVpnTopicEndpoint) SetRedeliveryEnabled(v bool) {
	o.RedeliveryEnabled = &v
}

// GetRejectLowPriorityMsgEnabled returns the RejectLowPriorityMsgEnabled field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetRejectLowPriorityMsgEnabled() bool {
	if o == nil || o.RejectLowPriorityMsgEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RejectLowPriorityMsgEnabled
}

// GetRejectLowPriorityMsgEnabledOk returns a tuple with the RejectLowPriorityMsgEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetRejectLowPriorityMsgEnabledOk() (*bool, bool) {
	if o == nil || o.RejectLowPriorityMsgEnabled == nil {
		return nil, false
	}
	return o.RejectLowPriorityMsgEnabled, true
}

// HasRejectLowPriorityMsgEnabled returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasRejectLowPriorityMsgEnabled() bool {
	if o != nil && o.RejectLowPriorityMsgEnabled != nil {
		return true
	}

	return false
}

// SetRejectLowPriorityMsgEnabled gets a reference to the given bool and assigns it to the RejectLowPriorityMsgEnabled field.
func (o *MsgVpnTopicEndpoint) SetRejectLowPriorityMsgEnabled(v bool) {
	o.RejectLowPriorityMsgEnabled = &v
}

// GetRejectLowPriorityMsgLimit returns the RejectLowPriorityMsgLimit field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetRejectLowPriorityMsgLimit() int64 {
	if o == nil || o.RejectLowPriorityMsgLimit == nil {
		var ret int64
		return ret
	}
	return *o.RejectLowPriorityMsgLimit
}

// GetRejectLowPriorityMsgLimitOk returns a tuple with the RejectLowPriorityMsgLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetRejectLowPriorityMsgLimitOk() (*int64, bool) {
	if o == nil || o.RejectLowPriorityMsgLimit == nil {
		return nil, false
	}
	return o.RejectLowPriorityMsgLimit, true
}

// HasRejectLowPriorityMsgLimit returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasRejectLowPriorityMsgLimit() bool {
	if o != nil && o.RejectLowPriorityMsgLimit != nil {
		return true
	}

	return false
}

// SetRejectLowPriorityMsgLimit gets a reference to the given int64 and assigns it to the RejectLowPriorityMsgLimit field.
func (o *MsgVpnTopicEndpoint) SetRejectLowPriorityMsgLimit(v int64) {
	o.RejectLowPriorityMsgLimit = &v
}

// GetRejectMsgToSenderOnDiscardBehavior returns the RejectMsgToSenderOnDiscardBehavior field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetRejectMsgToSenderOnDiscardBehavior() string {
	if o == nil || o.RejectMsgToSenderOnDiscardBehavior == nil {
		var ret string
		return ret
	}
	return *o.RejectMsgToSenderOnDiscardBehavior
}

// GetRejectMsgToSenderOnDiscardBehaviorOk returns a tuple with the RejectMsgToSenderOnDiscardBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetRejectMsgToSenderOnDiscardBehaviorOk() (*string, bool) {
	if o == nil || o.RejectMsgToSenderOnDiscardBehavior == nil {
		return nil, false
	}
	return o.RejectMsgToSenderOnDiscardBehavior, true
}

// HasRejectMsgToSenderOnDiscardBehavior returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasRejectMsgToSenderOnDiscardBehavior() bool {
	if o != nil && o.RejectMsgToSenderOnDiscardBehavior != nil {
		return true
	}

	return false
}

// SetRejectMsgToSenderOnDiscardBehavior gets a reference to the given string and assigns it to the RejectMsgToSenderOnDiscardBehavior field.
func (o *MsgVpnTopicEndpoint) SetRejectMsgToSenderOnDiscardBehavior(v string) {
	o.RejectMsgToSenderOnDiscardBehavior = &v
}

// GetReplayFailureCount returns the ReplayFailureCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetReplayFailureCount() int64 {
	if o == nil || o.ReplayFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplayFailureCount
}

// GetReplayFailureCountOk returns a tuple with the ReplayFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetReplayFailureCountOk() (*int64, bool) {
	if o == nil || o.ReplayFailureCount == nil {
		return nil, false
	}
	return o.ReplayFailureCount, true
}

// HasReplayFailureCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasReplayFailureCount() bool {
	if o != nil && o.ReplayFailureCount != nil {
		return true
	}

	return false
}

// SetReplayFailureCount gets a reference to the given int64 and assigns it to the ReplayFailureCount field.
func (o *MsgVpnTopicEndpoint) SetReplayFailureCount(v int64) {
	o.ReplayFailureCount = &v
}

// GetReplayStartCount returns the ReplayStartCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetReplayStartCount() int64 {
	if o == nil || o.ReplayStartCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplayStartCount
}

// GetReplayStartCountOk returns a tuple with the ReplayStartCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetReplayStartCountOk() (*int64, bool) {
	if o == nil || o.ReplayStartCount == nil {
		return nil, false
	}
	return o.ReplayStartCount, true
}

// HasReplayStartCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasReplayStartCount() bool {
	if o != nil && o.ReplayStartCount != nil {
		return true
	}

	return false
}

// SetReplayStartCount gets a reference to the given int64 and assigns it to the ReplayStartCount field.
func (o *MsgVpnTopicEndpoint) SetReplayStartCount(v int64) {
	o.ReplayStartCount = &v
}

// GetReplayState returns the ReplayState field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetReplayState() string {
	if o == nil || o.ReplayState == nil {
		var ret string
		return ret
	}
	return *o.ReplayState
}

// GetReplayStateOk returns a tuple with the ReplayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetReplayStateOk() (*string, bool) {
	if o == nil || o.ReplayState == nil {
		return nil, false
	}
	return o.ReplayState, true
}

// HasReplayState returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasReplayState() bool {
	if o != nil && o.ReplayState != nil {
		return true
	}

	return false
}

// SetReplayState gets a reference to the given string and assigns it to the ReplayState field.
func (o *MsgVpnTopicEndpoint) SetReplayState(v string) {
	o.ReplayState = &v
}

// GetReplaySuccessCount returns the ReplaySuccessCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetReplaySuccessCount() int64 {
	if o == nil || o.ReplaySuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplaySuccessCount
}

// GetReplaySuccessCountOk returns a tuple with the ReplaySuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetReplaySuccessCountOk() (*int64, bool) {
	if o == nil || o.ReplaySuccessCount == nil {
		return nil, false
	}
	return o.ReplaySuccessCount, true
}

// HasReplaySuccessCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasReplaySuccessCount() bool {
	if o != nil && o.ReplaySuccessCount != nil {
		return true
	}

	return false
}

// SetReplaySuccessCount gets a reference to the given int64 and assigns it to the ReplaySuccessCount field.
func (o *MsgVpnTopicEndpoint) SetReplaySuccessCount(v int64) {
	o.ReplaySuccessCount = &v
}

// GetReplayedAckedMsgCount returns the ReplayedAckedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetReplayedAckedMsgCount() int64 {
	if o == nil || o.ReplayedAckedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplayedAckedMsgCount
}

// GetReplayedAckedMsgCountOk returns a tuple with the ReplayedAckedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetReplayedAckedMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplayedAckedMsgCount == nil {
		return nil, false
	}
	return o.ReplayedAckedMsgCount, true
}

// HasReplayedAckedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasReplayedAckedMsgCount() bool {
	if o != nil && o.ReplayedAckedMsgCount != nil {
		return true
	}

	return false
}

// SetReplayedAckedMsgCount gets a reference to the given int64 and assigns it to the ReplayedAckedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetReplayedAckedMsgCount(v int64) {
	o.ReplayedAckedMsgCount = &v
}

// GetReplayedTxMsgCount returns the ReplayedTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetReplayedTxMsgCount() int64 {
	if o == nil || o.ReplayedTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplayedTxMsgCount
}

// GetReplayedTxMsgCountOk returns a tuple with the ReplayedTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetReplayedTxMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplayedTxMsgCount == nil {
		return nil, false
	}
	return o.ReplayedTxMsgCount, true
}

// HasReplayedTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasReplayedTxMsgCount() bool {
	if o != nil && o.ReplayedTxMsgCount != nil {
		return true
	}

	return false
}

// SetReplayedTxMsgCount gets a reference to the given int64 and assigns it to the ReplayedTxMsgCount field.
func (o *MsgVpnTopicEndpoint) SetReplayedTxMsgCount(v int64) {
	o.ReplayedTxMsgCount = &v
}

// GetReplicationActiveAckPropTxMsgCount returns the ReplicationActiveAckPropTxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetReplicationActiveAckPropTxMsgCount() int64 {
	if o == nil || o.ReplicationActiveAckPropTxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationActiveAckPropTxMsgCount
}

// GetReplicationActiveAckPropTxMsgCountOk returns a tuple with the ReplicationActiveAckPropTxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetReplicationActiveAckPropTxMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationActiveAckPropTxMsgCount == nil {
		return nil, false
	}
	return o.ReplicationActiveAckPropTxMsgCount, true
}

// HasReplicationActiveAckPropTxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasReplicationActiveAckPropTxMsgCount() bool {
	if o != nil && o.ReplicationActiveAckPropTxMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationActiveAckPropTxMsgCount gets a reference to the given int64 and assigns it to the ReplicationActiveAckPropTxMsgCount field.
func (o *MsgVpnTopicEndpoint) SetReplicationActiveAckPropTxMsgCount(v int64) {
	o.ReplicationActiveAckPropTxMsgCount = &v
}

// GetReplicationStandbyAckPropRxMsgCount returns the ReplicationStandbyAckPropRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetReplicationStandbyAckPropRxMsgCount() int64 {
	if o == nil || o.ReplicationStandbyAckPropRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyAckPropRxMsgCount
}

// GetReplicationStandbyAckPropRxMsgCountOk returns a tuple with the ReplicationStandbyAckPropRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetReplicationStandbyAckPropRxMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyAckPropRxMsgCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyAckPropRxMsgCount, true
}

// HasReplicationStandbyAckPropRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasReplicationStandbyAckPropRxMsgCount() bool {
	if o != nil && o.ReplicationStandbyAckPropRxMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyAckPropRxMsgCount gets a reference to the given int64 and assigns it to the ReplicationStandbyAckPropRxMsgCount field.
func (o *MsgVpnTopicEndpoint) SetReplicationStandbyAckPropRxMsgCount(v int64) {
	o.ReplicationStandbyAckPropRxMsgCount = &v
}

// GetReplicationStandbyAckedByAckPropMsgCount returns the ReplicationStandbyAckedByAckPropMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetReplicationStandbyAckedByAckPropMsgCount() int64 {
	if o == nil || o.ReplicationStandbyAckedByAckPropMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyAckedByAckPropMsgCount
}

// GetReplicationStandbyAckedByAckPropMsgCountOk returns a tuple with the ReplicationStandbyAckedByAckPropMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetReplicationStandbyAckedByAckPropMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyAckedByAckPropMsgCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyAckedByAckPropMsgCount, true
}

// HasReplicationStandbyAckedByAckPropMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasReplicationStandbyAckedByAckPropMsgCount() bool {
	if o != nil && o.ReplicationStandbyAckedByAckPropMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyAckedByAckPropMsgCount gets a reference to the given int64 and assigns it to the ReplicationStandbyAckedByAckPropMsgCount field.
func (o *MsgVpnTopicEndpoint) SetReplicationStandbyAckedByAckPropMsgCount(v int64) {
	o.ReplicationStandbyAckedByAckPropMsgCount = &v
}

// GetReplicationStandbyRxMsgCount returns the ReplicationStandbyRxMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetReplicationStandbyRxMsgCount() int64 {
	if o == nil || o.ReplicationStandbyRxMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationStandbyRxMsgCount
}

// GetReplicationStandbyRxMsgCountOk returns a tuple with the ReplicationStandbyRxMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetReplicationStandbyRxMsgCountOk() (*int64, bool) {
	if o == nil || o.ReplicationStandbyRxMsgCount == nil {
		return nil, false
	}
	return o.ReplicationStandbyRxMsgCount, true
}

// HasReplicationStandbyRxMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasReplicationStandbyRxMsgCount() bool {
	if o != nil && o.ReplicationStandbyRxMsgCount != nil {
		return true
	}

	return false
}

// SetReplicationStandbyRxMsgCount gets a reference to the given int64 and assigns it to the ReplicationStandbyRxMsgCount field.
func (o *MsgVpnTopicEndpoint) SetReplicationStandbyRxMsgCount(v int64) {
	o.ReplicationStandbyRxMsgCount = &v
}

// GetRespectMsgPriorityEnabled returns the RespectMsgPriorityEnabled field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetRespectMsgPriorityEnabled() bool {
	if o == nil || o.RespectMsgPriorityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RespectMsgPriorityEnabled
}

// GetRespectMsgPriorityEnabledOk returns a tuple with the RespectMsgPriorityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetRespectMsgPriorityEnabledOk() (*bool, bool) {
	if o == nil || o.RespectMsgPriorityEnabled == nil {
		return nil, false
	}
	return o.RespectMsgPriorityEnabled, true
}

// HasRespectMsgPriorityEnabled returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasRespectMsgPriorityEnabled() bool {
	if o != nil && o.RespectMsgPriorityEnabled != nil {
		return true
	}

	return false
}

// SetRespectMsgPriorityEnabled gets a reference to the given bool and assigns it to the RespectMsgPriorityEnabled field.
func (o *MsgVpnTopicEndpoint) SetRespectMsgPriorityEnabled(v bool) {
	o.RespectMsgPriorityEnabled = &v
}

// GetRespectTtlEnabled returns the RespectTtlEnabled field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetRespectTtlEnabled() bool {
	if o == nil || o.RespectTtlEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RespectTtlEnabled
}

// GetRespectTtlEnabledOk returns a tuple with the RespectTtlEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetRespectTtlEnabledOk() (*bool, bool) {
	if o == nil || o.RespectTtlEnabled == nil {
		return nil, false
	}
	return o.RespectTtlEnabled, true
}

// HasRespectTtlEnabled returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasRespectTtlEnabled() bool {
	if o != nil && o.RespectTtlEnabled != nil {
		return true
	}

	return false
}

// SetRespectTtlEnabled gets a reference to the given bool and assigns it to the RespectTtlEnabled field.
func (o *MsgVpnTopicEndpoint) SetRespectTtlEnabled(v bool) {
	o.RespectTtlEnabled = &v
}

// GetRxByteRate returns the RxByteRate field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetRxByteRate() int32 {
	if o == nil || o.RxByteRate == nil {
		var ret int32
		return ret
	}
	return *o.RxByteRate
}

// GetRxByteRateOk returns a tuple with the RxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetRxByteRateOk() (*int32, bool) {
	if o == nil || o.RxByteRate == nil {
		return nil, false
	}
	return o.RxByteRate, true
}

// HasRxByteRate returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasRxByteRate() bool {
	if o != nil && o.RxByteRate != nil {
		return true
	}

	return false
}

// SetRxByteRate gets a reference to the given int32 and assigns it to the RxByteRate field.
func (o *MsgVpnTopicEndpoint) SetRxByteRate(v int32) {
	o.RxByteRate = &v
}

// GetRxMsgRate returns the RxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetRxMsgRate() int64 {
	if o == nil || o.RxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.RxMsgRate
}

// GetRxMsgRateOk returns a tuple with the RxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetRxMsgRateOk() (*int64, bool) {
	if o == nil || o.RxMsgRate == nil {
		return nil, false
	}
	return o.RxMsgRate, true
}

// HasRxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasRxMsgRate() bool {
	if o != nil && o.RxMsgRate != nil {
		return true
	}

	return false
}

// SetRxMsgRate gets a reference to the given int64 and assigns it to the RxMsgRate field.
func (o *MsgVpnTopicEndpoint) SetRxMsgRate(v int64) {
	o.RxMsgRate = &v
}

// GetRxSelector returns the RxSelector field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetRxSelector() bool {
	if o == nil || o.RxSelector == nil {
		var ret bool
		return ret
	}
	return *o.RxSelector
}

// GetRxSelectorOk returns a tuple with the RxSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetRxSelectorOk() (*bool, bool) {
	if o == nil || o.RxSelector == nil {
		return nil, false
	}
	return o.RxSelector, true
}

// HasRxSelector returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasRxSelector() bool {
	if o != nil && o.RxSelector != nil {
		return true
	}

	return false
}

// SetRxSelector gets a reference to the given bool and assigns it to the RxSelector field.
func (o *MsgVpnTopicEndpoint) SetRxSelector(v bool) {
	o.RxSelector = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *MsgVpnTopicEndpoint) SetSelector(v string) {
	o.Selector = &v
}

// GetSelectorExaminedMsgCount returns the SelectorExaminedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetSelectorExaminedMsgCount() int64 {
	if o == nil || o.SelectorExaminedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SelectorExaminedMsgCount
}

// GetSelectorExaminedMsgCountOk returns a tuple with the SelectorExaminedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetSelectorExaminedMsgCountOk() (*int64, bool) {
	if o == nil || o.SelectorExaminedMsgCount == nil {
		return nil, false
	}
	return o.SelectorExaminedMsgCount, true
}

// HasSelectorExaminedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasSelectorExaminedMsgCount() bool {
	if o != nil && o.SelectorExaminedMsgCount != nil {
		return true
	}

	return false
}

// SetSelectorExaminedMsgCount gets a reference to the given int64 and assigns it to the SelectorExaminedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetSelectorExaminedMsgCount(v int64) {
	o.SelectorExaminedMsgCount = &v
}

// GetSelectorMatchedMsgCount returns the SelectorMatchedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetSelectorMatchedMsgCount() int64 {
	if o == nil || o.SelectorMatchedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SelectorMatchedMsgCount
}

// GetSelectorMatchedMsgCountOk returns a tuple with the SelectorMatchedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetSelectorMatchedMsgCountOk() (*int64, bool) {
	if o == nil || o.SelectorMatchedMsgCount == nil {
		return nil, false
	}
	return o.SelectorMatchedMsgCount, true
}

// HasSelectorMatchedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasSelectorMatchedMsgCount() bool {
	if o != nil && o.SelectorMatchedMsgCount != nil {
		return true
	}

	return false
}

// SetSelectorMatchedMsgCount gets a reference to the given int64 and assigns it to the SelectorMatchedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetSelectorMatchedMsgCount(v int64) {
	o.SelectorMatchedMsgCount = &v
}

// GetSelectorNotMatchedMsgCount returns the SelectorNotMatchedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetSelectorNotMatchedMsgCount() int64 {
	if o == nil || o.SelectorNotMatchedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SelectorNotMatchedMsgCount
}

// GetSelectorNotMatchedMsgCountOk returns a tuple with the SelectorNotMatchedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetSelectorNotMatchedMsgCountOk() (*int64, bool) {
	if o == nil || o.SelectorNotMatchedMsgCount == nil {
		return nil, false
	}
	return o.SelectorNotMatchedMsgCount, true
}

// HasSelectorNotMatchedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasSelectorNotMatchedMsgCount() bool {
	if o != nil && o.SelectorNotMatchedMsgCount != nil {
		return true
	}

	return false
}

// SetSelectorNotMatchedMsgCount gets a reference to the given int64 and assigns it to the SelectorNotMatchedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetSelectorNotMatchedMsgCount(v int64) {
	o.SelectorNotMatchedMsgCount = &v
}

// GetSpooledByteCount returns the SpooledByteCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetSpooledByteCount() int64 {
	if o == nil || o.SpooledByteCount == nil {
		var ret int64
		return ret
	}
	return *o.SpooledByteCount
}

// GetSpooledByteCountOk returns a tuple with the SpooledByteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetSpooledByteCountOk() (*int64, bool) {
	if o == nil || o.SpooledByteCount == nil {
		return nil, false
	}
	return o.SpooledByteCount, true
}

// HasSpooledByteCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasSpooledByteCount() bool {
	if o != nil && o.SpooledByteCount != nil {
		return true
	}

	return false
}

// SetSpooledByteCount gets a reference to the given int64 and assigns it to the SpooledByteCount field.
func (o *MsgVpnTopicEndpoint) SetSpooledByteCount(v int64) {
	o.SpooledByteCount = &v
}

// GetSpooledMsgCount returns the SpooledMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetSpooledMsgCount() int64 {
	if o == nil || o.SpooledMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SpooledMsgCount
}

// GetSpooledMsgCountOk returns a tuple with the SpooledMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetSpooledMsgCountOk() (*int64, bool) {
	if o == nil || o.SpooledMsgCount == nil {
		return nil, false
	}
	return o.SpooledMsgCount, true
}

// HasSpooledMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasSpooledMsgCount() bool {
	if o != nil && o.SpooledMsgCount != nil {
		return true
	}

	return false
}

// SetSpooledMsgCount gets a reference to the given int64 and assigns it to the SpooledMsgCount field.
func (o *MsgVpnTopicEndpoint) SetSpooledMsgCount(v int64) {
	o.SpooledMsgCount = &v
}

// GetTopicEndpointName returns the TopicEndpointName field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetTopicEndpointName() string {
	if o == nil || o.TopicEndpointName == nil {
		var ret string
		return ret
	}
	return *o.TopicEndpointName
}

// GetTopicEndpointNameOk returns a tuple with the TopicEndpointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetTopicEndpointNameOk() (*string, bool) {
	if o == nil || o.TopicEndpointName == nil {
		return nil, false
	}
	return o.TopicEndpointName, true
}

// HasTopicEndpointName returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasTopicEndpointName() bool {
	if o != nil && o.TopicEndpointName != nil {
		return true
	}

	return false
}

// SetTopicEndpointName gets a reference to the given string and assigns it to the TopicEndpointName field.
func (o *MsgVpnTopicEndpoint) SetTopicEndpointName(v string) {
	o.TopicEndpointName = &v
}

// GetTransportRetransmitMsgCount returns the TransportRetransmitMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetTransportRetransmitMsgCount() int64 {
	if o == nil || o.TransportRetransmitMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.TransportRetransmitMsgCount
}

// GetTransportRetransmitMsgCountOk returns a tuple with the TransportRetransmitMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetTransportRetransmitMsgCountOk() (*int64, bool) {
	if o == nil || o.TransportRetransmitMsgCount == nil {
		return nil, false
	}
	return o.TransportRetransmitMsgCount, true
}

// HasTransportRetransmitMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasTransportRetransmitMsgCount() bool {
	if o != nil && o.TransportRetransmitMsgCount != nil {
		return true
	}

	return false
}

// SetTransportRetransmitMsgCount gets a reference to the given int64 and assigns it to the TransportRetransmitMsgCount field.
func (o *MsgVpnTopicEndpoint) SetTransportRetransmitMsgCount(v int64) {
	o.TransportRetransmitMsgCount = &v
}

// GetTxByteRate returns the TxByteRate field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetTxByteRate() int64 {
	if o == nil || o.TxByteRate == nil {
		var ret int64
		return ret
	}
	return *o.TxByteRate
}

// GetTxByteRateOk returns a tuple with the TxByteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetTxByteRateOk() (*int64, bool) {
	if o == nil || o.TxByteRate == nil {
		return nil, false
	}
	return o.TxByteRate, true
}

// HasTxByteRate returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasTxByteRate() bool {
	if o != nil && o.TxByteRate != nil {
		return true
	}

	return false
}

// SetTxByteRate gets a reference to the given int64 and assigns it to the TxByteRate field.
func (o *MsgVpnTopicEndpoint) SetTxByteRate(v int64) {
	o.TxByteRate = &v
}

// GetTxMsgRate returns the TxMsgRate field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetTxMsgRate() int64 {
	if o == nil || o.TxMsgRate == nil {
		var ret int64
		return ret
	}
	return *o.TxMsgRate
}

// GetTxMsgRateOk returns a tuple with the TxMsgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetTxMsgRateOk() (*int64, bool) {
	if o == nil || o.TxMsgRate == nil {
		return nil, false
	}
	return o.TxMsgRate, true
}

// HasTxMsgRate returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasTxMsgRate() bool {
	if o != nil && o.TxMsgRate != nil {
		return true
	}

	return false
}

// SetTxMsgRate gets a reference to the given int64 and assigns it to the TxMsgRate field.
func (o *MsgVpnTopicEndpoint) SetTxMsgRate(v int64) {
	o.TxMsgRate = &v
}

// GetTxUnackedMsgCount returns the TxUnackedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetTxUnackedMsgCount() int64 {
	if o == nil || o.TxUnackedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.TxUnackedMsgCount
}

// GetTxUnackedMsgCountOk returns a tuple with the TxUnackedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetTxUnackedMsgCountOk() (*int64, bool) {
	if o == nil || o.TxUnackedMsgCount == nil {
		return nil, false
	}
	return o.TxUnackedMsgCount, true
}

// HasTxUnackedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasTxUnackedMsgCount() bool {
	if o != nil && o.TxUnackedMsgCount != nil {
		return true
	}

	return false
}

// SetTxUnackedMsgCount gets a reference to the given int64 and assigns it to the TxUnackedMsgCount field.
func (o *MsgVpnTopicEndpoint) SetTxUnackedMsgCount(v int64) {
	o.TxUnackedMsgCount = &v
}

// GetVirtualRouter returns the VirtualRouter field value if set, zero value otherwise.
func (o *MsgVpnTopicEndpoint) GetVirtualRouter() string {
	if o == nil || o.VirtualRouter == nil {
		var ret string
		return ret
	}
	return *o.VirtualRouter
}

// GetVirtualRouterOk returns a tuple with the VirtualRouter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnTopicEndpoint) GetVirtualRouterOk() (*string, bool) {
	if o == nil || o.VirtualRouter == nil {
		return nil, false
	}
	return o.VirtualRouter, true
}

// HasVirtualRouter returns a boolean if a field has been set.
func (o *MsgVpnTopicEndpoint) HasVirtualRouter() bool {
	if o != nil && o.VirtualRouter != nil {
		return true
	}

	return false
}

// SetVirtualRouter gets a reference to the given string and assigns it to the VirtualRouter field.
func (o *MsgVpnTopicEndpoint) SetVirtualRouter(v string) {
	o.VirtualRouter = &v
}

func (o MsgVpnTopicEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessType != nil {
		toSerialize["accessType"] = o.AccessType
	}
	if o.AlreadyBoundBindFailureCount != nil {
		toSerialize["alreadyBoundBindFailureCount"] = o.AlreadyBoundBindFailureCount
	}
	if o.AverageRxByteRate != nil {
		toSerialize["averageRxByteRate"] = o.AverageRxByteRate
	}
	if o.AverageRxMsgRate != nil {
		toSerialize["averageRxMsgRate"] = o.AverageRxMsgRate
	}
	if o.AverageTxByteRate != nil {
		toSerialize["averageTxByteRate"] = o.AverageTxByteRate
	}
	if o.AverageTxMsgRate != nil {
		toSerialize["averageTxMsgRate"] = o.AverageTxMsgRate
	}
	if o.BindRequestCount != nil {
		toSerialize["bindRequestCount"] = o.BindRequestCount
	}
	if o.BindSuccessCount != nil {
		toSerialize["bindSuccessCount"] = o.BindSuccessCount
	}
	if o.BindTimeForwardingMode != nil {
		toSerialize["bindTimeForwardingMode"] = o.BindTimeForwardingMode
	}
	if o.ClientProfileDeniedDiscardedMsgCount != nil {
		toSerialize["clientProfileDeniedDiscardedMsgCount"] = o.ClientProfileDeniedDiscardedMsgCount
	}
	if o.ConsumerAckPropagationEnabled != nil {
		toSerialize["consumerAckPropagationEnabled"] = o.ConsumerAckPropagationEnabled
	}
	if o.CreatedByManagement != nil {
		toSerialize["createdByManagement"] = o.CreatedByManagement
	}
	if o.DeadMsgQueue != nil {
		toSerialize["deadMsgQueue"] = o.DeadMsgQueue
	}
	if o.DeletedMsgCount != nil {
		toSerialize["deletedMsgCount"] = o.DeletedMsgCount
	}
	if o.DeliveryCountEnabled != nil {
		toSerialize["deliveryCountEnabled"] = o.DeliveryCountEnabled
	}
	if o.DestinationGroupErrorDiscardedMsgCount != nil {
		toSerialize["destinationGroupErrorDiscardedMsgCount"] = o.DestinationGroupErrorDiscardedMsgCount
	}
	if o.DestinationTopic != nil {
		toSerialize["destinationTopic"] = o.DestinationTopic
	}
	if o.DisabledBindFailureCount != nil {
		toSerialize["disabledBindFailureCount"] = o.DisabledBindFailureCount
	}
	if o.DisabledDiscardedMsgCount != nil {
		toSerialize["disabledDiscardedMsgCount"] = o.DisabledDiscardedMsgCount
	}
	if o.Durable != nil {
		toSerialize["durable"] = o.Durable
	}
	if o.EgressEnabled != nil {
		toSerialize["egressEnabled"] = o.EgressEnabled
	}
	if o.EventBindCountThreshold != nil {
		toSerialize["eventBindCountThreshold"] = o.EventBindCountThreshold
	}
	if o.EventRejectLowPriorityMsgLimitThreshold != nil {
		toSerialize["eventRejectLowPriorityMsgLimitThreshold"] = o.EventRejectLowPriorityMsgLimitThreshold
	}
	if o.EventSpoolUsageThreshold != nil {
		toSerialize["eventSpoolUsageThreshold"] = o.EventSpoolUsageThreshold
	}
	if o.HighestAckedMsgId != nil {
		toSerialize["highestAckedMsgId"] = o.HighestAckedMsgId
	}
	if o.HighestMsgId != nil {
		toSerialize["highestMsgId"] = o.HighestMsgId
	}
	if o.InProgressAckMsgCount != nil {
		toSerialize["inProgressAckMsgCount"] = o.InProgressAckMsgCount
	}
	if o.IngressEnabled != nil {
		toSerialize["ingressEnabled"] = o.IngressEnabled
	}
	if o.InvalidSelectorBindFailureCount != nil {
		toSerialize["invalidSelectorBindFailureCount"] = o.InvalidSelectorBindFailureCount
	}
	if o.LastReplayCompleteTime != nil {
		toSerialize["lastReplayCompleteTime"] = o.LastReplayCompleteTime
	}
	if o.LastReplayFailureReason != nil {
		toSerialize["lastReplayFailureReason"] = o.LastReplayFailureReason
	}
	if o.LastReplayFailureTime != nil {
		toSerialize["lastReplayFailureTime"] = o.LastReplayFailureTime
	}
	if o.LastReplayStartTime != nil {
		toSerialize["lastReplayStartTime"] = o.LastReplayStartTime
	}
	if o.LastReplayedMsgTxTime != nil {
		toSerialize["lastReplayedMsgTxTime"] = o.LastReplayedMsgTxTime
	}
	if o.LastSelectorExaminedMsgId != nil {
		toSerialize["lastSelectorExaminedMsgId"] = o.LastSelectorExaminedMsgId
	}
	if o.LastSpooledMsgId != nil {
		toSerialize["lastSpooledMsgId"] = o.LastSpooledMsgId
	}
	if o.LowPriorityMsgCongestionDiscardedMsgCount != nil {
		toSerialize["lowPriorityMsgCongestionDiscardedMsgCount"] = o.LowPriorityMsgCongestionDiscardedMsgCount
	}
	if o.LowPriorityMsgCongestionState != nil {
		toSerialize["lowPriorityMsgCongestionState"] = o.LowPriorityMsgCongestionState
	}
	if o.LowestAckedMsgId != nil {
		toSerialize["lowestAckedMsgId"] = o.LowestAckedMsgId
	}
	if o.LowestMsgId != nil {
		toSerialize["lowestMsgId"] = o.LowestMsgId
	}
	if o.MaxBindCount != nil {
		toSerialize["maxBindCount"] = o.MaxBindCount
	}
	if o.MaxBindCountExceededBindFailureCount != nil {
		toSerialize["maxBindCountExceededBindFailureCount"] = o.MaxBindCountExceededBindFailureCount
	}
	if o.MaxDeliveredUnackedMsgsPerFlow != nil {
		toSerialize["maxDeliveredUnackedMsgsPerFlow"] = o.MaxDeliveredUnackedMsgsPerFlow
	}
	if o.MaxEffectiveBindCount != nil {
		toSerialize["maxEffectiveBindCount"] = o.MaxEffectiveBindCount
	}
	if o.MaxMsgSize != nil {
		toSerialize["maxMsgSize"] = o.MaxMsgSize
	}
	if o.MaxMsgSizeExceededDiscardedMsgCount != nil {
		toSerialize["maxMsgSizeExceededDiscardedMsgCount"] = o.MaxMsgSizeExceededDiscardedMsgCount
	}
	if o.MaxMsgSpoolUsageExceededDiscardedMsgCount != nil {
		toSerialize["maxMsgSpoolUsageExceededDiscardedMsgCount"] = o.MaxMsgSpoolUsageExceededDiscardedMsgCount
	}
	if o.MaxRedeliveryCount != nil {
		toSerialize["maxRedeliveryCount"] = o.MaxRedeliveryCount
	}
	if o.MaxRedeliveryExceededDiscardedMsgCount != nil {
		toSerialize["maxRedeliveryExceededDiscardedMsgCount"] = o.MaxRedeliveryExceededDiscardedMsgCount
	}
	if o.MaxRedeliveryExceededToDmqFailedMsgCount != nil {
		toSerialize["maxRedeliveryExceededToDmqFailedMsgCount"] = o.MaxRedeliveryExceededToDmqFailedMsgCount
	}
	if o.MaxRedeliveryExceededToDmqMsgCount != nil {
		toSerialize["maxRedeliveryExceededToDmqMsgCount"] = o.MaxRedeliveryExceededToDmqMsgCount
	}
	if o.MaxSpoolUsage != nil {
		toSerialize["maxSpoolUsage"] = o.MaxSpoolUsage
	}
	if o.MaxTtl != nil {
		toSerialize["maxTtl"] = o.MaxTtl
	}
	if o.MaxTtlExceededDiscardedMsgCount != nil {
		toSerialize["maxTtlExceededDiscardedMsgCount"] = o.MaxTtlExceededDiscardedMsgCount
	}
	if o.MaxTtlExpiredDiscardedMsgCount != nil {
		toSerialize["maxTtlExpiredDiscardedMsgCount"] = o.MaxTtlExpiredDiscardedMsgCount
	}
	if o.MaxTtlExpiredToDmqFailedMsgCount != nil {
		toSerialize["maxTtlExpiredToDmqFailedMsgCount"] = o.MaxTtlExpiredToDmqFailedMsgCount
	}
	if o.MaxTtlExpiredToDmqMsgCount != nil {
		toSerialize["maxTtlExpiredToDmqMsgCount"] = o.MaxTtlExpiredToDmqMsgCount
	}
	if o.MsgSpoolPeakUsage != nil {
		toSerialize["msgSpoolPeakUsage"] = o.MsgSpoolPeakUsage
	}
	if o.MsgSpoolUsage != nil {
		toSerialize["msgSpoolUsage"] = o.MsgSpoolUsage
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.NetworkTopic != nil {
		toSerialize["networkTopic"] = o.NetworkTopic
	}
	if o.NoLocalDeliveryDiscardedMsgCount != nil {
		toSerialize["noLocalDeliveryDiscardedMsgCount"] = o.NoLocalDeliveryDiscardedMsgCount
	}
	if o.OtherBindFailureCount != nil {
		toSerialize["otherBindFailureCount"] = o.OtherBindFailureCount
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.RedeliveredMsgCount != nil {
		toSerialize["redeliveredMsgCount"] = o.RedeliveredMsgCount
	}
	if o.RedeliveryEnabled != nil {
		toSerialize["redeliveryEnabled"] = o.RedeliveryEnabled
	}
	if o.RejectLowPriorityMsgEnabled != nil {
		toSerialize["rejectLowPriorityMsgEnabled"] = o.RejectLowPriorityMsgEnabled
	}
	if o.RejectLowPriorityMsgLimit != nil {
		toSerialize["rejectLowPriorityMsgLimit"] = o.RejectLowPriorityMsgLimit
	}
	if o.RejectMsgToSenderOnDiscardBehavior != nil {
		toSerialize["rejectMsgToSenderOnDiscardBehavior"] = o.RejectMsgToSenderOnDiscardBehavior
	}
	if o.ReplayFailureCount != nil {
		toSerialize["replayFailureCount"] = o.ReplayFailureCount
	}
	if o.ReplayStartCount != nil {
		toSerialize["replayStartCount"] = o.ReplayStartCount
	}
	if o.ReplayState != nil {
		toSerialize["replayState"] = o.ReplayState
	}
	if o.ReplaySuccessCount != nil {
		toSerialize["replaySuccessCount"] = o.ReplaySuccessCount
	}
	if o.ReplayedAckedMsgCount != nil {
		toSerialize["replayedAckedMsgCount"] = o.ReplayedAckedMsgCount
	}
	if o.ReplayedTxMsgCount != nil {
		toSerialize["replayedTxMsgCount"] = o.ReplayedTxMsgCount
	}
	if o.ReplicationActiveAckPropTxMsgCount != nil {
		toSerialize["replicationActiveAckPropTxMsgCount"] = o.ReplicationActiveAckPropTxMsgCount
	}
	if o.ReplicationStandbyAckPropRxMsgCount != nil {
		toSerialize["replicationStandbyAckPropRxMsgCount"] = o.ReplicationStandbyAckPropRxMsgCount
	}
	if o.ReplicationStandbyAckedByAckPropMsgCount != nil {
		toSerialize["replicationStandbyAckedByAckPropMsgCount"] = o.ReplicationStandbyAckedByAckPropMsgCount
	}
	if o.ReplicationStandbyRxMsgCount != nil {
		toSerialize["replicationStandbyRxMsgCount"] = o.ReplicationStandbyRxMsgCount
	}
	if o.RespectMsgPriorityEnabled != nil {
		toSerialize["respectMsgPriorityEnabled"] = o.RespectMsgPriorityEnabled
	}
	if o.RespectTtlEnabled != nil {
		toSerialize["respectTtlEnabled"] = o.RespectTtlEnabled
	}
	if o.RxByteRate != nil {
		toSerialize["rxByteRate"] = o.RxByteRate
	}
	if o.RxMsgRate != nil {
		toSerialize["rxMsgRate"] = o.RxMsgRate
	}
	if o.RxSelector != nil {
		toSerialize["rxSelector"] = o.RxSelector
	}
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	if o.SelectorExaminedMsgCount != nil {
		toSerialize["selectorExaminedMsgCount"] = o.SelectorExaminedMsgCount
	}
	if o.SelectorMatchedMsgCount != nil {
		toSerialize["selectorMatchedMsgCount"] = o.SelectorMatchedMsgCount
	}
	if o.SelectorNotMatchedMsgCount != nil {
		toSerialize["selectorNotMatchedMsgCount"] = o.SelectorNotMatchedMsgCount
	}
	if o.SpooledByteCount != nil {
		toSerialize["spooledByteCount"] = o.SpooledByteCount
	}
	if o.SpooledMsgCount != nil {
		toSerialize["spooledMsgCount"] = o.SpooledMsgCount
	}
	if o.TopicEndpointName != nil {
		toSerialize["topicEndpointName"] = o.TopicEndpointName
	}
	if o.TransportRetransmitMsgCount != nil {
		toSerialize["transportRetransmitMsgCount"] = o.TransportRetransmitMsgCount
	}
	if o.TxByteRate != nil {
		toSerialize["txByteRate"] = o.TxByteRate
	}
	if o.TxMsgRate != nil {
		toSerialize["txMsgRate"] = o.TxMsgRate
	}
	if o.TxUnackedMsgCount != nil {
		toSerialize["txUnackedMsgCount"] = o.TxUnackedMsgCount
	}
	if o.VirtualRouter != nil {
		toSerialize["virtualRouter"] = o.VirtualRouter
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnTopicEndpoint struct {
	value *MsgVpnTopicEndpoint
	isSet bool
}

func (v NullableMsgVpnTopicEndpoint) Get() *MsgVpnTopicEndpoint {
	return v.value
}

func (v *NullableMsgVpnTopicEndpoint) Set(val *MsgVpnTopicEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnTopicEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnTopicEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnTopicEndpoint(val *MsgVpnTopicEndpoint) *NullableMsgVpnTopicEndpoint {
	return &NullableMsgVpnTopicEndpoint{value: val, isSet: true}
}

func (v NullableMsgVpnTopicEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnTopicEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
