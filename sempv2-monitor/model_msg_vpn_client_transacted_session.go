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

// MsgVpnClientTransactedSession struct for MsgVpnClientTransactedSession
type MsgVpnClientTransactedSession struct {
	// The name of the Client.
	ClientName *string `json:"clientName,omitempty"`
	// The number of transactions committed within the Transacted Session.
	CommitCount *int64 `json:"commitCount,omitempty"`
	// The number of transaction commit operations that failed.
	CommitFailureCount *int64 `json:"commitFailureCount,omitempty"`
	// The number of transaction commit operations that succeeded.
	CommitSuccessCount *int64 `json:"commitSuccessCount,omitempty"`
	// The number of messages consumed within the Transacted Session.
	ConsumedMsgCount *int64 `json:"consumedMsgCount,omitempty"`
	// The number of transaction end fail operations that failed.
	EndFailFailureCount *int64 `json:"endFailFailureCount,omitempty"`
	// The number of transaction end fail operations that succeeded.
	EndFailSuccessCount *int64 `json:"endFailSuccessCount,omitempty"`
	// The number of transaction end operations that failed.
	EndFailureCount *int64 `json:"endFailureCount,omitempty"`
	// The number of transaction end rollback operations that failed.
	EndRollbackFailureCount *int64 `json:"endRollbackFailureCount,omitempty"`
	// The number of transaction end rollback operations that succeeded.
	EndRollbackSuccessCount *int64 `json:"endRollbackSuccessCount,omitempty"`
	// The number of transaction end operations that succeeded.
	EndSuccessCount *int64 `json:"endSuccessCount,omitempty"`
	// The number of transactions that failed within the Transacted Session.
	FailureCount *int64 `json:"failureCount,omitempty"`
	// The number of transaction forget operations that failed.
	ForgetFailureCount *int64 `json:"forgetFailureCount,omitempty"`
	// The number of transaction forget operations that succeeded.
	ForgetSuccessCount *int64 `json:"forgetSuccessCount,omitempty"`
	// The name of the Message VPN.
	MsgVpnName *string `json:"msgVpnName,omitempty"`
	// The number of transaction one-phase commit operations that failed.
	OnePhaseCommitFailureCount *int64 `json:"onePhaseCommitFailureCount,omitempty"`
	// The number of transaction one-phase commit operations that succeeded.
	OnePhaseCommitSuccessCount *int64 `json:"onePhaseCommitSuccessCount,omitempty"`
	// The number of messages to be consumed when the transaction is committed.
	PendingConsumedMsgCount *int32 `json:"pendingConsumedMsgCount,omitempty"`
	// The number of messages to be published when the transaction is committed.
	PendingPublishedMsgCount *int32 `json:"pendingPublishedMsgCount,omitempty"`
	// The number of transaction prepare operations that failed.
	PrepareFailureCount *int64 `json:"prepareFailureCount,omitempty"`
	// The number of transaction prepare operations that succeeded.
	PrepareSuccessCount *int64 `json:"prepareSuccessCount,omitempty"`
	// The state of the previous transaction. The allowed values and their meaning are:  <pre> \"none\" - The previous transaction had no state. \"committed\" - The previous transaction was committed. \"rolled-back\" - The previous transaction was rolled back. \"failed\" - The previous transaction failed. </pre>
	PreviousTransactionState *string `json:"previousTransactionState,omitempty"`
	// The number of messages published within the Transacted Session.
	PublishedMsgCount *int64 `json:"publishedMsgCount,omitempty"`
	// The number of transaction resume operations that failed.
	ResumeFailureCount *int64 `json:"resumeFailureCount,omitempty"`
	// The number of transaction resume operations that succeeded.
	ResumeSuccessCount *int64 `json:"resumeSuccessCount,omitempty"`
	// The number of messages retrieved within the Transacted Session.
	RetrievedMsgCount *int64 `json:"retrievedMsgCount,omitempty"`
	// The number of transactions rolled back within the Transacted Session.
	RollbackCount *int64 `json:"rollbackCount,omitempty"`
	// The number of transaction rollback operations that failed.
	RollbackFailureCount *int64 `json:"rollbackFailureCount,omitempty"`
	// The number of transaction rollback operations that succeeded.
	RollbackSuccessCount *int64 `json:"rollbackSuccessCount,omitempty"`
	// The name of the Transacted Session.
	SessionName *string `json:"sessionName,omitempty"`
	// The number of messages spooled within the Transacted Session.
	SpooledMsgCount *int64 `json:"spooledMsgCount,omitempty"`
	// The number of transaction start operations that failed.
	StartFailureCount *int64 `json:"startFailureCount,omitempty"`
	// The number of transaction start operations that succeeded.
	StartSuccessCount *int64 `json:"startSuccessCount,omitempty"`
	// The number of transactions that succeeded within the Transacted Session.
	SuccessCount *int64 `json:"successCount,omitempty"`
	// The number of transaction suspend operations that failed.
	SuspendFailureCount *int64 `json:"suspendFailureCount,omitempty"`
	// The number of transaction suspend operations that succeeded.
	SuspendSuccessCount *int64 `json:"suspendSuccessCount,omitempty"`
	// The identifier (ID) of the transaction in the Transacted Session.
	TransactionId *int32 `json:"transactionId,omitempty"`
	// The state of the current transaction. The allowed values and their meaning are:  <pre> \"in-progress\" - The current transaction is in progress. \"committing\" - The current transaction is committing. \"rolling-back\" - The current transaction is rolling back. \"failing\" - The current transaction is failing. </pre>
	TransactionState *string `json:"transactionState,omitempty"`
	// The number of transaction two-phase commit operations that failed.
	TwoPhaseCommitFailureCount *int64 `json:"twoPhaseCommitFailureCount,omitempty"`
	// The number of transaction two-phase commit operations that succeeded.
	TwoPhaseCommitSuccessCount *int64 `json:"twoPhaseCommitSuccessCount,omitempty"`
}

// NewMsgVpnClientTransactedSession instantiates a new MsgVpnClientTransactedSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsgVpnClientTransactedSession() *MsgVpnClientTransactedSession {
	this := MsgVpnClientTransactedSession{}
	return &this
}

// NewMsgVpnClientTransactedSessionWithDefaults instantiates a new MsgVpnClientTransactedSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsgVpnClientTransactedSessionWithDefaults() *MsgVpnClientTransactedSession {
	this := MsgVpnClientTransactedSession{}
	return &this
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *MsgVpnClientTransactedSession) SetClientName(v string) {
	o.ClientName = &v
}

// GetCommitCount returns the CommitCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetCommitCount() int64 {
	if o == nil || o.CommitCount == nil {
		var ret int64
		return ret
	}
	return *o.CommitCount
}

// GetCommitCountOk returns a tuple with the CommitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetCommitCountOk() (*int64, bool) {
	if o == nil || o.CommitCount == nil {
		return nil, false
	}
	return o.CommitCount, true
}

// HasCommitCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasCommitCount() bool {
	if o != nil && o.CommitCount != nil {
		return true
	}

	return false
}

// SetCommitCount gets a reference to the given int64 and assigns it to the CommitCount field.
func (o *MsgVpnClientTransactedSession) SetCommitCount(v int64) {
	o.CommitCount = &v
}

// GetCommitFailureCount returns the CommitFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetCommitFailureCount() int64 {
	if o == nil || o.CommitFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.CommitFailureCount
}

// GetCommitFailureCountOk returns a tuple with the CommitFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetCommitFailureCountOk() (*int64, bool) {
	if o == nil || o.CommitFailureCount == nil {
		return nil, false
	}
	return o.CommitFailureCount, true
}

// HasCommitFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasCommitFailureCount() bool {
	if o != nil && o.CommitFailureCount != nil {
		return true
	}

	return false
}

// SetCommitFailureCount gets a reference to the given int64 and assigns it to the CommitFailureCount field.
func (o *MsgVpnClientTransactedSession) SetCommitFailureCount(v int64) {
	o.CommitFailureCount = &v
}

// GetCommitSuccessCount returns the CommitSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetCommitSuccessCount() int64 {
	if o == nil || o.CommitSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.CommitSuccessCount
}

// GetCommitSuccessCountOk returns a tuple with the CommitSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetCommitSuccessCountOk() (*int64, bool) {
	if o == nil || o.CommitSuccessCount == nil {
		return nil, false
	}
	return o.CommitSuccessCount, true
}

// HasCommitSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasCommitSuccessCount() bool {
	if o != nil && o.CommitSuccessCount != nil {
		return true
	}

	return false
}

// SetCommitSuccessCount gets a reference to the given int64 and assigns it to the CommitSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetCommitSuccessCount(v int64) {
	o.CommitSuccessCount = &v
}

// GetConsumedMsgCount returns the ConsumedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetConsumedMsgCount() int64 {
	if o == nil || o.ConsumedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.ConsumedMsgCount
}

// GetConsumedMsgCountOk returns a tuple with the ConsumedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetConsumedMsgCountOk() (*int64, bool) {
	if o == nil || o.ConsumedMsgCount == nil {
		return nil, false
	}
	return o.ConsumedMsgCount, true
}

// HasConsumedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasConsumedMsgCount() bool {
	if o != nil && o.ConsumedMsgCount != nil {
		return true
	}

	return false
}

// SetConsumedMsgCount gets a reference to the given int64 and assigns it to the ConsumedMsgCount field.
func (o *MsgVpnClientTransactedSession) SetConsumedMsgCount(v int64) {
	o.ConsumedMsgCount = &v
}

// GetEndFailFailureCount returns the EndFailFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetEndFailFailureCount() int64 {
	if o == nil || o.EndFailFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.EndFailFailureCount
}

// GetEndFailFailureCountOk returns a tuple with the EndFailFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetEndFailFailureCountOk() (*int64, bool) {
	if o == nil || o.EndFailFailureCount == nil {
		return nil, false
	}
	return o.EndFailFailureCount, true
}

// HasEndFailFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasEndFailFailureCount() bool {
	if o != nil && o.EndFailFailureCount != nil {
		return true
	}

	return false
}

// SetEndFailFailureCount gets a reference to the given int64 and assigns it to the EndFailFailureCount field.
func (o *MsgVpnClientTransactedSession) SetEndFailFailureCount(v int64) {
	o.EndFailFailureCount = &v
}

// GetEndFailSuccessCount returns the EndFailSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetEndFailSuccessCount() int64 {
	if o == nil || o.EndFailSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.EndFailSuccessCount
}

// GetEndFailSuccessCountOk returns a tuple with the EndFailSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetEndFailSuccessCountOk() (*int64, bool) {
	if o == nil || o.EndFailSuccessCount == nil {
		return nil, false
	}
	return o.EndFailSuccessCount, true
}

// HasEndFailSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasEndFailSuccessCount() bool {
	if o != nil && o.EndFailSuccessCount != nil {
		return true
	}

	return false
}

// SetEndFailSuccessCount gets a reference to the given int64 and assigns it to the EndFailSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetEndFailSuccessCount(v int64) {
	o.EndFailSuccessCount = &v
}

// GetEndFailureCount returns the EndFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetEndFailureCount() int64 {
	if o == nil || o.EndFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.EndFailureCount
}

// GetEndFailureCountOk returns a tuple with the EndFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetEndFailureCountOk() (*int64, bool) {
	if o == nil || o.EndFailureCount == nil {
		return nil, false
	}
	return o.EndFailureCount, true
}

// HasEndFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasEndFailureCount() bool {
	if o != nil && o.EndFailureCount != nil {
		return true
	}

	return false
}

// SetEndFailureCount gets a reference to the given int64 and assigns it to the EndFailureCount field.
func (o *MsgVpnClientTransactedSession) SetEndFailureCount(v int64) {
	o.EndFailureCount = &v
}

// GetEndRollbackFailureCount returns the EndRollbackFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetEndRollbackFailureCount() int64 {
	if o == nil || o.EndRollbackFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.EndRollbackFailureCount
}

// GetEndRollbackFailureCountOk returns a tuple with the EndRollbackFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetEndRollbackFailureCountOk() (*int64, bool) {
	if o == nil || o.EndRollbackFailureCount == nil {
		return nil, false
	}
	return o.EndRollbackFailureCount, true
}

// HasEndRollbackFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasEndRollbackFailureCount() bool {
	if o != nil && o.EndRollbackFailureCount != nil {
		return true
	}

	return false
}

// SetEndRollbackFailureCount gets a reference to the given int64 and assigns it to the EndRollbackFailureCount field.
func (o *MsgVpnClientTransactedSession) SetEndRollbackFailureCount(v int64) {
	o.EndRollbackFailureCount = &v
}

// GetEndRollbackSuccessCount returns the EndRollbackSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetEndRollbackSuccessCount() int64 {
	if o == nil || o.EndRollbackSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.EndRollbackSuccessCount
}

// GetEndRollbackSuccessCountOk returns a tuple with the EndRollbackSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetEndRollbackSuccessCountOk() (*int64, bool) {
	if o == nil || o.EndRollbackSuccessCount == nil {
		return nil, false
	}
	return o.EndRollbackSuccessCount, true
}

// HasEndRollbackSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasEndRollbackSuccessCount() bool {
	if o != nil && o.EndRollbackSuccessCount != nil {
		return true
	}

	return false
}

// SetEndRollbackSuccessCount gets a reference to the given int64 and assigns it to the EndRollbackSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetEndRollbackSuccessCount(v int64) {
	o.EndRollbackSuccessCount = &v
}

// GetEndSuccessCount returns the EndSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetEndSuccessCount() int64 {
	if o == nil || o.EndSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.EndSuccessCount
}

// GetEndSuccessCountOk returns a tuple with the EndSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetEndSuccessCountOk() (*int64, bool) {
	if o == nil || o.EndSuccessCount == nil {
		return nil, false
	}
	return o.EndSuccessCount, true
}

// HasEndSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasEndSuccessCount() bool {
	if o != nil && o.EndSuccessCount != nil {
		return true
	}

	return false
}

// SetEndSuccessCount gets a reference to the given int64 and assigns it to the EndSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetEndSuccessCount(v int64) {
	o.EndSuccessCount = &v
}

// GetFailureCount returns the FailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetFailureCount() int64 {
	if o == nil || o.FailureCount == nil {
		var ret int64
		return ret
	}
	return *o.FailureCount
}

// GetFailureCountOk returns a tuple with the FailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetFailureCountOk() (*int64, bool) {
	if o == nil || o.FailureCount == nil {
		return nil, false
	}
	return o.FailureCount, true
}

// HasFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasFailureCount() bool {
	if o != nil && o.FailureCount != nil {
		return true
	}

	return false
}

// SetFailureCount gets a reference to the given int64 and assigns it to the FailureCount field.
func (o *MsgVpnClientTransactedSession) SetFailureCount(v int64) {
	o.FailureCount = &v
}

// GetForgetFailureCount returns the ForgetFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetForgetFailureCount() int64 {
	if o == nil || o.ForgetFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.ForgetFailureCount
}

// GetForgetFailureCountOk returns a tuple with the ForgetFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetForgetFailureCountOk() (*int64, bool) {
	if o == nil || o.ForgetFailureCount == nil {
		return nil, false
	}
	return o.ForgetFailureCount, true
}

// HasForgetFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasForgetFailureCount() bool {
	if o != nil && o.ForgetFailureCount != nil {
		return true
	}

	return false
}

// SetForgetFailureCount gets a reference to the given int64 and assigns it to the ForgetFailureCount field.
func (o *MsgVpnClientTransactedSession) SetForgetFailureCount(v int64) {
	o.ForgetFailureCount = &v
}

// GetForgetSuccessCount returns the ForgetSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetForgetSuccessCount() int64 {
	if o == nil || o.ForgetSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.ForgetSuccessCount
}

// GetForgetSuccessCountOk returns a tuple with the ForgetSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetForgetSuccessCountOk() (*int64, bool) {
	if o == nil || o.ForgetSuccessCount == nil {
		return nil, false
	}
	return o.ForgetSuccessCount, true
}

// HasForgetSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasForgetSuccessCount() bool {
	if o != nil && o.ForgetSuccessCount != nil {
		return true
	}

	return false
}

// SetForgetSuccessCount gets a reference to the given int64 and assigns it to the ForgetSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetForgetSuccessCount(v int64) {
	o.ForgetSuccessCount = &v
}

// GetMsgVpnName returns the MsgVpnName field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetMsgVpnName() string {
	if o == nil || o.MsgVpnName == nil {
		var ret string
		return ret
	}
	return *o.MsgVpnName
}

// GetMsgVpnNameOk returns a tuple with the MsgVpnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetMsgVpnNameOk() (*string, bool) {
	if o == nil || o.MsgVpnName == nil {
		return nil, false
	}
	return o.MsgVpnName, true
}

// HasMsgVpnName returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasMsgVpnName() bool {
	if o != nil && o.MsgVpnName != nil {
		return true
	}

	return false
}

// SetMsgVpnName gets a reference to the given string and assigns it to the MsgVpnName field.
func (o *MsgVpnClientTransactedSession) SetMsgVpnName(v string) {
	o.MsgVpnName = &v
}

// GetOnePhaseCommitFailureCount returns the OnePhaseCommitFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetOnePhaseCommitFailureCount() int64 {
	if o == nil || o.OnePhaseCommitFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.OnePhaseCommitFailureCount
}

// GetOnePhaseCommitFailureCountOk returns a tuple with the OnePhaseCommitFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetOnePhaseCommitFailureCountOk() (*int64, bool) {
	if o == nil || o.OnePhaseCommitFailureCount == nil {
		return nil, false
	}
	return o.OnePhaseCommitFailureCount, true
}

// HasOnePhaseCommitFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasOnePhaseCommitFailureCount() bool {
	if o != nil && o.OnePhaseCommitFailureCount != nil {
		return true
	}

	return false
}

// SetOnePhaseCommitFailureCount gets a reference to the given int64 and assigns it to the OnePhaseCommitFailureCount field.
func (o *MsgVpnClientTransactedSession) SetOnePhaseCommitFailureCount(v int64) {
	o.OnePhaseCommitFailureCount = &v
}

// GetOnePhaseCommitSuccessCount returns the OnePhaseCommitSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetOnePhaseCommitSuccessCount() int64 {
	if o == nil || o.OnePhaseCommitSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.OnePhaseCommitSuccessCount
}

// GetOnePhaseCommitSuccessCountOk returns a tuple with the OnePhaseCommitSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetOnePhaseCommitSuccessCountOk() (*int64, bool) {
	if o == nil || o.OnePhaseCommitSuccessCount == nil {
		return nil, false
	}
	return o.OnePhaseCommitSuccessCount, true
}

// HasOnePhaseCommitSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasOnePhaseCommitSuccessCount() bool {
	if o != nil && o.OnePhaseCommitSuccessCount != nil {
		return true
	}

	return false
}

// SetOnePhaseCommitSuccessCount gets a reference to the given int64 and assigns it to the OnePhaseCommitSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetOnePhaseCommitSuccessCount(v int64) {
	o.OnePhaseCommitSuccessCount = &v
}

// GetPendingConsumedMsgCount returns the PendingConsumedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetPendingConsumedMsgCount() int32 {
	if o == nil || o.PendingConsumedMsgCount == nil {
		var ret int32
		return ret
	}
	return *o.PendingConsumedMsgCount
}

// GetPendingConsumedMsgCountOk returns a tuple with the PendingConsumedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetPendingConsumedMsgCountOk() (*int32, bool) {
	if o == nil || o.PendingConsumedMsgCount == nil {
		return nil, false
	}
	return o.PendingConsumedMsgCount, true
}

// HasPendingConsumedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasPendingConsumedMsgCount() bool {
	if o != nil && o.PendingConsumedMsgCount != nil {
		return true
	}

	return false
}

// SetPendingConsumedMsgCount gets a reference to the given int32 and assigns it to the PendingConsumedMsgCount field.
func (o *MsgVpnClientTransactedSession) SetPendingConsumedMsgCount(v int32) {
	o.PendingConsumedMsgCount = &v
}

// GetPendingPublishedMsgCount returns the PendingPublishedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetPendingPublishedMsgCount() int32 {
	if o == nil || o.PendingPublishedMsgCount == nil {
		var ret int32
		return ret
	}
	return *o.PendingPublishedMsgCount
}

// GetPendingPublishedMsgCountOk returns a tuple with the PendingPublishedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetPendingPublishedMsgCountOk() (*int32, bool) {
	if o == nil || o.PendingPublishedMsgCount == nil {
		return nil, false
	}
	return o.PendingPublishedMsgCount, true
}

// HasPendingPublishedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasPendingPublishedMsgCount() bool {
	if o != nil && o.PendingPublishedMsgCount != nil {
		return true
	}

	return false
}

// SetPendingPublishedMsgCount gets a reference to the given int32 and assigns it to the PendingPublishedMsgCount field.
func (o *MsgVpnClientTransactedSession) SetPendingPublishedMsgCount(v int32) {
	o.PendingPublishedMsgCount = &v
}

// GetPrepareFailureCount returns the PrepareFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetPrepareFailureCount() int64 {
	if o == nil || o.PrepareFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.PrepareFailureCount
}

// GetPrepareFailureCountOk returns a tuple with the PrepareFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetPrepareFailureCountOk() (*int64, bool) {
	if o == nil || o.PrepareFailureCount == nil {
		return nil, false
	}
	return o.PrepareFailureCount, true
}

// HasPrepareFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasPrepareFailureCount() bool {
	if o != nil && o.PrepareFailureCount != nil {
		return true
	}

	return false
}

// SetPrepareFailureCount gets a reference to the given int64 and assigns it to the PrepareFailureCount field.
func (o *MsgVpnClientTransactedSession) SetPrepareFailureCount(v int64) {
	o.PrepareFailureCount = &v
}

// GetPrepareSuccessCount returns the PrepareSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetPrepareSuccessCount() int64 {
	if o == nil || o.PrepareSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.PrepareSuccessCount
}

// GetPrepareSuccessCountOk returns a tuple with the PrepareSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetPrepareSuccessCountOk() (*int64, bool) {
	if o == nil || o.PrepareSuccessCount == nil {
		return nil, false
	}
	return o.PrepareSuccessCount, true
}

// HasPrepareSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasPrepareSuccessCount() bool {
	if o != nil && o.PrepareSuccessCount != nil {
		return true
	}

	return false
}

// SetPrepareSuccessCount gets a reference to the given int64 and assigns it to the PrepareSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetPrepareSuccessCount(v int64) {
	o.PrepareSuccessCount = &v
}

// GetPreviousTransactionState returns the PreviousTransactionState field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetPreviousTransactionState() string {
	if o == nil || o.PreviousTransactionState == nil {
		var ret string
		return ret
	}
	return *o.PreviousTransactionState
}

// GetPreviousTransactionStateOk returns a tuple with the PreviousTransactionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetPreviousTransactionStateOk() (*string, bool) {
	if o == nil || o.PreviousTransactionState == nil {
		return nil, false
	}
	return o.PreviousTransactionState, true
}

// HasPreviousTransactionState returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasPreviousTransactionState() bool {
	if o != nil && o.PreviousTransactionState != nil {
		return true
	}

	return false
}

// SetPreviousTransactionState gets a reference to the given string and assigns it to the PreviousTransactionState field.
func (o *MsgVpnClientTransactedSession) SetPreviousTransactionState(v string) {
	o.PreviousTransactionState = &v
}

// GetPublishedMsgCount returns the PublishedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetPublishedMsgCount() int64 {
	if o == nil || o.PublishedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.PublishedMsgCount
}

// GetPublishedMsgCountOk returns a tuple with the PublishedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetPublishedMsgCountOk() (*int64, bool) {
	if o == nil || o.PublishedMsgCount == nil {
		return nil, false
	}
	return o.PublishedMsgCount, true
}

// HasPublishedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasPublishedMsgCount() bool {
	if o != nil && o.PublishedMsgCount != nil {
		return true
	}

	return false
}

// SetPublishedMsgCount gets a reference to the given int64 and assigns it to the PublishedMsgCount field.
func (o *MsgVpnClientTransactedSession) SetPublishedMsgCount(v int64) {
	o.PublishedMsgCount = &v
}

// GetResumeFailureCount returns the ResumeFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetResumeFailureCount() int64 {
	if o == nil || o.ResumeFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.ResumeFailureCount
}

// GetResumeFailureCountOk returns a tuple with the ResumeFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetResumeFailureCountOk() (*int64, bool) {
	if o == nil || o.ResumeFailureCount == nil {
		return nil, false
	}
	return o.ResumeFailureCount, true
}

// HasResumeFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasResumeFailureCount() bool {
	if o != nil && o.ResumeFailureCount != nil {
		return true
	}

	return false
}

// SetResumeFailureCount gets a reference to the given int64 and assigns it to the ResumeFailureCount field.
func (o *MsgVpnClientTransactedSession) SetResumeFailureCount(v int64) {
	o.ResumeFailureCount = &v
}

// GetResumeSuccessCount returns the ResumeSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetResumeSuccessCount() int64 {
	if o == nil || o.ResumeSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.ResumeSuccessCount
}

// GetResumeSuccessCountOk returns a tuple with the ResumeSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetResumeSuccessCountOk() (*int64, bool) {
	if o == nil || o.ResumeSuccessCount == nil {
		return nil, false
	}
	return o.ResumeSuccessCount, true
}

// HasResumeSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasResumeSuccessCount() bool {
	if o != nil && o.ResumeSuccessCount != nil {
		return true
	}

	return false
}

// SetResumeSuccessCount gets a reference to the given int64 and assigns it to the ResumeSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetResumeSuccessCount(v int64) {
	o.ResumeSuccessCount = &v
}

// GetRetrievedMsgCount returns the RetrievedMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetRetrievedMsgCount() int64 {
	if o == nil || o.RetrievedMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.RetrievedMsgCount
}

// GetRetrievedMsgCountOk returns a tuple with the RetrievedMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetRetrievedMsgCountOk() (*int64, bool) {
	if o == nil || o.RetrievedMsgCount == nil {
		return nil, false
	}
	return o.RetrievedMsgCount, true
}

// HasRetrievedMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasRetrievedMsgCount() bool {
	if o != nil && o.RetrievedMsgCount != nil {
		return true
	}

	return false
}

// SetRetrievedMsgCount gets a reference to the given int64 and assigns it to the RetrievedMsgCount field.
func (o *MsgVpnClientTransactedSession) SetRetrievedMsgCount(v int64) {
	o.RetrievedMsgCount = &v
}

// GetRollbackCount returns the RollbackCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetRollbackCount() int64 {
	if o == nil || o.RollbackCount == nil {
		var ret int64
		return ret
	}
	return *o.RollbackCount
}

// GetRollbackCountOk returns a tuple with the RollbackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetRollbackCountOk() (*int64, bool) {
	if o == nil || o.RollbackCount == nil {
		return nil, false
	}
	return o.RollbackCount, true
}

// HasRollbackCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasRollbackCount() bool {
	if o != nil && o.RollbackCount != nil {
		return true
	}

	return false
}

// SetRollbackCount gets a reference to the given int64 and assigns it to the RollbackCount field.
func (o *MsgVpnClientTransactedSession) SetRollbackCount(v int64) {
	o.RollbackCount = &v
}

// GetRollbackFailureCount returns the RollbackFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetRollbackFailureCount() int64 {
	if o == nil || o.RollbackFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.RollbackFailureCount
}

// GetRollbackFailureCountOk returns a tuple with the RollbackFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetRollbackFailureCountOk() (*int64, bool) {
	if o == nil || o.RollbackFailureCount == nil {
		return nil, false
	}
	return o.RollbackFailureCount, true
}

// HasRollbackFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasRollbackFailureCount() bool {
	if o != nil && o.RollbackFailureCount != nil {
		return true
	}

	return false
}

// SetRollbackFailureCount gets a reference to the given int64 and assigns it to the RollbackFailureCount field.
func (o *MsgVpnClientTransactedSession) SetRollbackFailureCount(v int64) {
	o.RollbackFailureCount = &v
}

// GetRollbackSuccessCount returns the RollbackSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetRollbackSuccessCount() int64 {
	if o == nil || o.RollbackSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.RollbackSuccessCount
}

// GetRollbackSuccessCountOk returns a tuple with the RollbackSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetRollbackSuccessCountOk() (*int64, bool) {
	if o == nil || o.RollbackSuccessCount == nil {
		return nil, false
	}
	return o.RollbackSuccessCount, true
}

// HasRollbackSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasRollbackSuccessCount() bool {
	if o != nil && o.RollbackSuccessCount != nil {
		return true
	}

	return false
}

// SetRollbackSuccessCount gets a reference to the given int64 and assigns it to the RollbackSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetRollbackSuccessCount(v int64) {
	o.RollbackSuccessCount = &v
}

// GetSessionName returns the SessionName field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetSessionName() string {
	if o == nil || o.SessionName == nil {
		var ret string
		return ret
	}
	return *o.SessionName
}

// GetSessionNameOk returns a tuple with the SessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetSessionNameOk() (*string, bool) {
	if o == nil || o.SessionName == nil {
		return nil, false
	}
	return o.SessionName, true
}

// HasSessionName returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasSessionName() bool {
	if o != nil && o.SessionName != nil {
		return true
	}

	return false
}

// SetSessionName gets a reference to the given string and assigns it to the SessionName field.
func (o *MsgVpnClientTransactedSession) SetSessionName(v string) {
	o.SessionName = &v
}

// GetSpooledMsgCount returns the SpooledMsgCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetSpooledMsgCount() int64 {
	if o == nil || o.SpooledMsgCount == nil {
		var ret int64
		return ret
	}
	return *o.SpooledMsgCount
}

// GetSpooledMsgCountOk returns a tuple with the SpooledMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetSpooledMsgCountOk() (*int64, bool) {
	if o == nil || o.SpooledMsgCount == nil {
		return nil, false
	}
	return o.SpooledMsgCount, true
}

// HasSpooledMsgCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasSpooledMsgCount() bool {
	if o != nil && o.SpooledMsgCount != nil {
		return true
	}

	return false
}

// SetSpooledMsgCount gets a reference to the given int64 and assigns it to the SpooledMsgCount field.
func (o *MsgVpnClientTransactedSession) SetSpooledMsgCount(v int64) {
	o.SpooledMsgCount = &v
}

// GetStartFailureCount returns the StartFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetStartFailureCount() int64 {
	if o == nil || o.StartFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.StartFailureCount
}

// GetStartFailureCountOk returns a tuple with the StartFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetStartFailureCountOk() (*int64, bool) {
	if o == nil || o.StartFailureCount == nil {
		return nil, false
	}
	return o.StartFailureCount, true
}

// HasStartFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasStartFailureCount() bool {
	if o != nil && o.StartFailureCount != nil {
		return true
	}

	return false
}

// SetStartFailureCount gets a reference to the given int64 and assigns it to the StartFailureCount field.
func (o *MsgVpnClientTransactedSession) SetStartFailureCount(v int64) {
	o.StartFailureCount = &v
}

// GetStartSuccessCount returns the StartSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetStartSuccessCount() int64 {
	if o == nil || o.StartSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.StartSuccessCount
}

// GetStartSuccessCountOk returns a tuple with the StartSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetStartSuccessCountOk() (*int64, bool) {
	if o == nil || o.StartSuccessCount == nil {
		return nil, false
	}
	return o.StartSuccessCount, true
}

// HasStartSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasStartSuccessCount() bool {
	if o != nil && o.StartSuccessCount != nil {
		return true
	}

	return false
}

// SetStartSuccessCount gets a reference to the given int64 and assigns it to the StartSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetStartSuccessCount(v int64) {
	o.StartSuccessCount = &v
}

// GetSuccessCount returns the SuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetSuccessCount() int64 {
	if o == nil || o.SuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.SuccessCount
}

// GetSuccessCountOk returns a tuple with the SuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetSuccessCountOk() (*int64, bool) {
	if o == nil || o.SuccessCount == nil {
		return nil, false
	}
	return o.SuccessCount, true
}

// HasSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasSuccessCount() bool {
	if o != nil && o.SuccessCount != nil {
		return true
	}

	return false
}

// SetSuccessCount gets a reference to the given int64 and assigns it to the SuccessCount field.
func (o *MsgVpnClientTransactedSession) SetSuccessCount(v int64) {
	o.SuccessCount = &v
}

// GetSuspendFailureCount returns the SuspendFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetSuspendFailureCount() int64 {
	if o == nil || o.SuspendFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.SuspendFailureCount
}

// GetSuspendFailureCountOk returns a tuple with the SuspendFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetSuspendFailureCountOk() (*int64, bool) {
	if o == nil || o.SuspendFailureCount == nil {
		return nil, false
	}
	return o.SuspendFailureCount, true
}

// HasSuspendFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasSuspendFailureCount() bool {
	if o != nil && o.SuspendFailureCount != nil {
		return true
	}

	return false
}

// SetSuspendFailureCount gets a reference to the given int64 and assigns it to the SuspendFailureCount field.
func (o *MsgVpnClientTransactedSession) SetSuspendFailureCount(v int64) {
	o.SuspendFailureCount = &v
}

// GetSuspendSuccessCount returns the SuspendSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetSuspendSuccessCount() int64 {
	if o == nil || o.SuspendSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.SuspendSuccessCount
}

// GetSuspendSuccessCountOk returns a tuple with the SuspendSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetSuspendSuccessCountOk() (*int64, bool) {
	if o == nil || o.SuspendSuccessCount == nil {
		return nil, false
	}
	return o.SuspendSuccessCount, true
}

// HasSuspendSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasSuspendSuccessCount() bool {
	if o != nil && o.SuspendSuccessCount != nil {
		return true
	}

	return false
}

// SetSuspendSuccessCount gets a reference to the given int64 and assigns it to the SuspendSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetSuspendSuccessCount(v int64) {
	o.SuspendSuccessCount = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetTransactionId() int32 {
	if o == nil || o.TransactionId == nil {
		var ret int32
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetTransactionIdOk() (*int32, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given int32 and assigns it to the TransactionId field.
func (o *MsgVpnClientTransactedSession) SetTransactionId(v int32) {
	o.TransactionId = &v
}

// GetTransactionState returns the TransactionState field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetTransactionState() string {
	if o == nil || o.TransactionState == nil {
		var ret string
		return ret
	}
	return *o.TransactionState
}

// GetTransactionStateOk returns a tuple with the TransactionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetTransactionStateOk() (*string, bool) {
	if o == nil || o.TransactionState == nil {
		return nil, false
	}
	return o.TransactionState, true
}

// HasTransactionState returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasTransactionState() bool {
	if o != nil && o.TransactionState != nil {
		return true
	}

	return false
}

// SetTransactionState gets a reference to the given string and assigns it to the TransactionState field.
func (o *MsgVpnClientTransactedSession) SetTransactionState(v string) {
	o.TransactionState = &v
}

// GetTwoPhaseCommitFailureCount returns the TwoPhaseCommitFailureCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetTwoPhaseCommitFailureCount() int64 {
	if o == nil || o.TwoPhaseCommitFailureCount == nil {
		var ret int64
		return ret
	}
	return *o.TwoPhaseCommitFailureCount
}

// GetTwoPhaseCommitFailureCountOk returns a tuple with the TwoPhaseCommitFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetTwoPhaseCommitFailureCountOk() (*int64, bool) {
	if o == nil || o.TwoPhaseCommitFailureCount == nil {
		return nil, false
	}
	return o.TwoPhaseCommitFailureCount, true
}

// HasTwoPhaseCommitFailureCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasTwoPhaseCommitFailureCount() bool {
	if o != nil && o.TwoPhaseCommitFailureCount != nil {
		return true
	}

	return false
}

// SetTwoPhaseCommitFailureCount gets a reference to the given int64 and assigns it to the TwoPhaseCommitFailureCount field.
func (o *MsgVpnClientTransactedSession) SetTwoPhaseCommitFailureCount(v int64) {
	o.TwoPhaseCommitFailureCount = &v
}

// GetTwoPhaseCommitSuccessCount returns the TwoPhaseCommitSuccessCount field value if set, zero value otherwise.
func (o *MsgVpnClientTransactedSession) GetTwoPhaseCommitSuccessCount() int64 {
	if o == nil || o.TwoPhaseCommitSuccessCount == nil {
		var ret int64
		return ret
	}
	return *o.TwoPhaseCommitSuccessCount
}

// GetTwoPhaseCommitSuccessCountOk returns a tuple with the TwoPhaseCommitSuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MsgVpnClientTransactedSession) GetTwoPhaseCommitSuccessCountOk() (*int64, bool) {
	if o == nil || o.TwoPhaseCommitSuccessCount == nil {
		return nil, false
	}
	return o.TwoPhaseCommitSuccessCount, true
}

// HasTwoPhaseCommitSuccessCount returns a boolean if a field has been set.
func (o *MsgVpnClientTransactedSession) HasTwoPhaseCommitSuccessCount() bool {
	if o != nil && o.TwoPhaseCommitSuccessCount != nil {
		return true
	}

	return false
}

// SetTwoPhaseCommitSuccessCount gets a reference to the given int64 and assigns it to the TwoPhaseCommitSuccessCount field.
func (o *MsgVpnClientTransactedSession) SetTwoPhaseCommitSuccessCount(v int64) {
	o.TwoPhaseCommitSuccessCount = &v
}

func (o MsgVpnClientTransactedSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.CommitCount != nil {
		toSerialize["commitCount"] = o.CommitCount
	}
	if o.CommitFailureCount != nil {
		toSerialize["commitFailureCount"] = o.CommitFailureCount
	}
	if o.CommitSuccessCount != nil {
		toSerialize["commitSuccessCount"] = o.CommitSuccessCount
	}
	if o.ConsumedMsgCount != nil {
		toSerialize["consumedMsgCount"] = o.ConsumedMsgCount
	}
	if o.EndFailFailureCount != nil {
		toSerialize["endFailFailureCount"] = o.EndFailFailureCount
	}
	if o.EndFailSuccessCount != nil {
		toSerialize["endFailSuccessCount"] = o.EndFailSuccessCount
	}
	if o.EndFailureCount != nil {
		toSerialize["endFailureCount"] = o.EndFailureCount
	}
	if o.EndRollbackFailureCount != nil {
		toSerialize["endRollbackFailureCount"] = o.EndRollbackFailureCount
	}
	if o.EndRollbackSuccessCount != nil {
		toSerialize["endRollbackSuccessCount"] = o.EndRollbackSuccessCount
	}
	if o.EndSuccessCount != nil {
		toSerialize["endSuccessCount"] = o.EndSuccessCount
	}
	if o.FailureCount != nil {
		toSerialize["failureCount"] = o.FailureCount
	}
	if o.ForgetFailureCount != nil {
		toSerialize["forgetFailureCount"] = o.ForgetFailureCount
	}
	if o.ForgetSuccessCount != nil {
		toSerialize["forgetSuccessCount"] = o.ForgetSuccessCount
	}
	if o.MsgVpnName != nil {
		toSerialize["msgVpnName"] = o.MsgVpnName
	}
	if o.OnePhaseCommitFailureCount != nil {
		toSerialize["onePhaseCommitFailureCount"] = o.OnePhaseCommitFailureCount
	}
	if o.OnePhaseCommitSuccessCount != nil {
		toSerialize["onePhaseCommitSuccessCount"] = o.OnePhaseCommitSuccessCount
	}
	if o.PendingConsumedMsgCount != nil {
		toSerialize["pendingConsumedMsgCount"] = o.PendingConsumedMsgCount
	}
	if o.PendingPublishedMsgCount != nil {
		toSerialize["pendingPublishedMsgCount"] = o.PendingPublishedMsgCount
	}
	if o.PrepareFailureCount != nil {
		toSerialize["prepareFailureCount"] = o.PrepareFailureCount
	}
	if o.PrepareSuccessCount != nil {
		toSerialize["prepareSuccessCount"] = o.PrepareSuccessCount
	}
	if o.PreviousTransactionState != nil {
		toSerialize["previousTransactionState"] = o.PreviousTransactionState
	}
	if o.PublishedMsgCount != nil {
		toSerialize["publishedMsgCount"] = o.PublishedMsgCount
	}
	if o.ResumeFailureCount != nil {
		toSerialize["resumeFailureCount"] = o.ResumeFailureCount
	}
	if o.ResumeSuccessCount != nil {
		toSerialize["resumeSuccessCount"] = o.ResumeSuccessCount
	}
	if o.RetrievedMsgCount != nil {
		toSerialize["retrievedMsgCount"] = o.RetrievedMsgCount
	}
	if o.RollbackCount != nil {
		toSerialize["rollbackCount"] = o.RollbackCount
	}
	if o.RollbackFailureCount != nil {
		toSerialize["rollbackFailureCount"] = o.RollbackFailureCount
	}
	if o.RollbackSuccessCount != nil {
		toSerialize["rollbackSuccessCount"] = o.RollbackSuccessCount
	}
	if o.SessionName != nil {
		toSerialize["sessionName"] = o.SessionName
	}
	if o.SpooledMsgCount != nil {
		toSerialize["spooledMsgCount"] = o.SpooledMsgCount
	}
	if o.StartFailureCount != nil {
		toSerialize["startFailureCount"] = o.StartFailureCount
	}
	if o.StartSuccessCount != nil {
		toSerialize["startSuccessCount"] = o.StartSuccessCount
	}
	if o.SuccessCount != nil {
		toSerialize["successCount"] = o.SuccessCount
	}
	if o.SuspendFailureCount != nil {
		toSerialize["suspendFailureCount"] = o.SuspendFailureCount
	}
	if o.SuspendSuccessCount != nil {
		toSerialize["suspendSuccessCount"] = o.SuspendSuccessCount
	}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	if o.TransactionState != nil {
		toSerialize["transactionState"] = o.TransactionState
	}
	if o.TwoPhaseCommitFailureCount != nil {
		toSerialize["twoPhaseCommitFailureCount"] = o.TwoPhaseCommitFailureCount
	}
	if o.TwoPhaseCommitSuccessCount != nil {
		toSerialize["twoPhaseCommitSuccessCount"] = o.TwoPhaseCommitSuccessCount
	}
	return json.Marshal(toSerialize)
}

type NullableMsgVpnClientTransactedSession struct {
	value *MsgVpnClientTransactedSession
	isSet bool
}

func (v NullableMsgVpnClientTransactedSession) Get() *MsgVpnClientTransactedSession {
	return v.value
}

func (v *NullableMsgVpnClientTransactedSession) Set(val *MsgVpnClientTransactedSession) {
	v.value = val
	v.isSet = true
}

func (v NullableMsgVpnClientTransactedSession) IsSet() bool {
	return v.isSet
}

func (v *NullableMsgVpnClientTransactedSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsgVpnClientTransactedSession(val *MsgVpnClientTransactedSession) *NullableMsgVpnClientTransactedSession {
	return &NullableMsgVpnClientTransactedSession{value: val, isSet: true}
}

func (v NullableMsgVpnClientTransactedSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsgVpnClientTransactedSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
