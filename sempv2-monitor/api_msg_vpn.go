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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// MsgVpnApiService MsgVpnApi service
type MsgVpnApiService service

type MsgVpnApiApiGetMsgVpnRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRequest) Execute() (MsgVpnResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnExecute(r)
}

/*
 * GetMsgVpn Get a Message VPN object.
 * Get a Message VPN object.

Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgingTlsServerCertEnforceTrustedCommonNameEnabled||x
counter.controlRxByteCount||x
counter.controlRxMsgCount||x
counter.controlTxByteCount||x
counter.controlTxMsgCount||x
counter.dataRxByteCount||x
counter.dataRxMsgCount||x
counter.dataTxByteCount||x
counter.dataTxMsgCount||x
counter.discardedRxMsgCount||x
counter.discardedTxMsgCount||x
counter.loginRxMsgCount||x
counter.loginTxMsgCount||x
counter.msgSpoolRxMsgCount||x
counter.msgSpoolTxMsgCount||x
counter.tlsRxByteCount||x
counter.tlsTxByteCount||x
msgVpnName|x|
rate.averageRxByteRate||x
rate.averageRxMsgRate||x
rate.averageTxByteRate||x
rate.averageTxMsgRate||x
rate.rxByteRate||x
rate.rxMsgRate||x
rate.tlsAverageRxByteRate||x
rate.tlsAverageTxByteRate||x
rate.tlsRxByteRate||x
rate.tlsTxByteRate||x
rate.txByteRate||x
rate.txMsgRate||x
restTlsServerCertEnforceTrustedCommonNameEnabled||x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnRequest
*/
func (a *MsgVpnApiService) GetMsgVpn(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnRequest {
	return MsgVpnApiApiGetMsgVpnRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnResponse
 */
func (a *MsgVpnApiService) GetMsgVpnExecute(r MsgVpnApiApiGetMsgVpnRequest) (MsgVpnResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpn")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfileRequest struct {
	ctx            _context.Context
	ApiService     *MsgVpnApiService
	msgVpnName     string
	aclProfileName string
	select_        *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfileRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfileRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfileRequest) Execute() (MsgVpnAclProfileResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfileExecute(r)
}

/*
 * GetMsgVpnAclProfile Get an ACL Profile object.
 * Get an ACL Profile object.

An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @return MsgVpnApiApiGetMsgVpnAclProfileRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfile(ctx _context.Context, msgVpnName string, aclProfileName string) MsgVpnApiApiGetMsgVpnAclProfileRequest {
	return MsgVpnApiApiGetMsgVpnAclProfileRequest{
		ApiService:     a,
		ctx:            ctx,
		msgVpnName:     msgVpnName,
		aclProfileName: aclProfileName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfileResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfileExecute(r MsgVpnApiApiGetMsgVpnAclProfileRequest) (MsgVpnAclProfileResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionRequest struct {
	ctx                           _context.Context
	ApiService                    *MsgVpnApiService
	msgVpnName                    string
	aclProfileName                string
	clientConnectExceptionAddress string
	select_                       *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionRequest) Execute() (MsgVpnAclProfileClientConnectExceptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfileClientConnectExceptionExecute(r)
}

/*
 * GetMsgVpnAclProfileClientConnectException Get a Client Connect Exception object.
 * Get a Client Connect Exception object.

A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|
clientConnectExceptionAddress|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @param clientConnectExceptionAddress The IP address/netmask of the client connect exception in CIDR form.
 * @return MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfileClientConnectException(ctx _context.Context, msgVpnName string, aclProfileName string, clientConnectExceptionAddress string) MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionRequest {
	return MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionRequest{
		ApiService:                    a,
		ctx:                           ctx,
		msgVpnName:                    msgVpnName,
		aclProfileName:                aclProfileName,
		clientConnectExceptionAddress: clientConnectExceptionAddress,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfileClientConnectExceptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfileClientConnectExceptionExecute(r MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionRequest) (MsgVpnAclProfileClientConnectExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileClientConnectExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfileClientConnectException")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientConnectExceptionAddress"+"}", _neturl.PathEscape(parameterToString(r.clientConnectExceptionAddress, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *MsgVpnApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) Execute() (MsgVpnAclProfileClientConnectExceptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfileClientConnectExceptionsExecute(r)
}

/*
 * GetMsgVpnAclProfileClientConnectExceptions Get a list of Client Connect Exception objects.
 * Get a list of Client Connect Exception objects.

A Client Connect Exception is an exception to the default action to take when a client using the ACL Profile connects to the Message VPN. Exceptions must be expressed as an IP address/netmask in CIDR form.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|
clientConnectExceptionAddress|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @return MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfileClientConnectExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest {
	return MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest{
		ApiService:     a,
		ctx:            ctx,
		msgVpnName:     msgVpnName,
		aclProfileName: aclProfileName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfileClientConnectExceptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfileClientConnectExceptionsExecute(r MsgVpnApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) (MsgVpnAclProfileClientConnectExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileClientConnectExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfileClientConnectExceptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	aclProfileName        string
	topicSyntax           string
	publishExceptionTopic string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionRequest) Execute() (MsgVpnAclProfilePublishExceptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfilePublishExceptionExecute(r)
}

/*
 * GetMsgVpnAclProfilePublishException Get a Publish Topic Exception object.
 * Get a Publish Topic Exception object.

A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|x
msgVpnName|x|x
publishExceptionTopic|x|x
topicSyntax|x|x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been deprecated since 2.14. Replaced by publishTopicExceptions.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @param topicSyntax The syntax of the topic for the exception to the default action taken.
 * @param publishExceptionTopic The topic for the exception to the default action taken. May include wildcard characters.
 * @return MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfilePublishException(ctx _context.Context, msgVpnName string, aclProfileName string, topicSyntax string, publishExceptionTopic string) MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionRequest {
	return MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		aclProfileName:        aclProfileName,
		topicSyntax:           topicSyntax,
		publishExceptionTopic: publishExceptionTopic,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfilePublishExceptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfilePublishExceptionExecute(r MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionRequest) (MsgVpnAclProfilePublishExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfilePublishExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfilePublishException")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicSyntax"+"}", _neturl.PathEscape(parameterToString(r.topicSyntax, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"publishExceptionTopic"+"}", _neturl.PathEscape(parameterToString(r.publishExceptionTopic, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *MsgVpnApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest) Execute() (MsgVpnAclProfilePublishExceptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfilePublishExceptionsExecute(r)
}

/*
 * GetMsgVpnAclProfilePublishExceptions Get a list of Publish Topic Exception objects.
 * Get a list of Publish Topic Exception objects.

A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|x
msgVpnName|x|x
publishExceptionTopic|x|x
topicSyntax|x|x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been deprecated since 2.14. Replaced by publishTopicExceptions.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @return MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfilePublishExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest {
	return MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest{
		ApiService:     a,
		ctx:            ctx,
		msgVpnName:     msgVpnName,
		aclProfileName: aclProfileName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfilePublishExceptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfilePublishExceptionsExecute(r MsgVpnApiApiGetMsgVpnAclProfilePublishExceptionsRequest) (MsgVpnAclProfilePublishExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfilePublishExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfilePublishExceptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest struct {
	ctx                         _context.Context
	ApiService                  *MsgVpnApiService
	msgVpnName                  string
	aclProfileName              string
	publishTopicExceptionSyntax string
	publishTopicException       string
	select_                     *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest) Execute() (MsgVpnAclProfilePublishTopicExceptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfilePublishTopicExceptionExecute(r)
}

/*
 * GetMsgVpnAclProfilePublishTopicException Get a Publish Topic Exception object.
 * Get a Publish Topic Exception object.

A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|
msgVpnName|x|
publishTopicException|x|
publishTopicExceptionSyntax|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.14.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @param publishTopicExceptionSyntax The syntax of the topic for the exception to the default action taken.
 * @param publishTopicException The topic for the exception to the default action taken. May include wildcard characters.
 * @return MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfilePublishTopicException(ctx _context.Context, msgVpnName string, aclProfileName string, publishTopicExceptionSyntax string, publishTopicException string) MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest {
	return MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest{
		ApiService:                  a,
		ctx:                         ctx,
		msgVpnName:                  msgVpnName,
		aclProfileName:              aclProfileName,
		publishTopicExceptionSyntax: publishTopicExceptionSyntax,
		publishTopicException:       publishTopicException,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfilePublishTopicExceptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfilePublishTopicExceptionExecute(r MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest) (MsgVpnAclProfilePublishTopicExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfilePublishTopicExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfilePublishTopicException")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"publishTopicExceptionSyntax"+"}", _neturl.PathEscape(parameterToString(r.publishTopicExceptionSyntax, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"publishTopicException"+"}", _neturl.PathEscape(parameterToString(r.publishTopicException, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *MsgVpnApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) Execute() (MsgVpnAclProfilePublishTopicExceptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfilePublishTopicExceptionsExecute(r)
}

/*
 * GetMsgVpnAclProfilePublishTopicExceptions Get a list of Publish Topic Exception objects.
 * Get a list of Publish Topic Exception objects.

A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|
msgVpnName|x|
publishTopicException|x|
publishTopicExceptionSyntax|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.14.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @return MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfilePublishTopicExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest {
	return MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest{
		ApiService:     a,
		ctx:            ctx,
		msgVpnName:     msgVpnName,
		aclProfileName: aclProfileName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfilePublishTopicExceptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfilePublishTopicExceptionsExecute(r MsgVpnApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) (MsgVpnAclProfilePublishTopicExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfilePublishTopicExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfilePublishTopicExceptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionRequest struct {
	ctx                     _context.Context
	ApiService              *MsgVpnApiService
	msgVpnName              string
	aclProfileName          string
	topicSyntax             string
	subscribeExceptionTopic string
	select_                 *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionRequest) Execute() (MsgVpnAclProfileSubscribeExceptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfileSubscribeExceptionExecute(r)
}

/*
 * GetMsgVpnAclProfileSubscribeException Get a Subscribe Topic Exception object.
 * Get a Subscribe Topic Exception object.

A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|x
msgVpnName|x|x
subscribeExceptionTopic|x|x
topicSyntax|x|x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @param topicSyntax The syntax of the topic for the exception to the default action taken.
 * @param subscribeExceptionTopic The topic for the exception to the default action taken. May include wildcard characters.
 * @return MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeException(ctx _context.Context, msgVpnName string, aclProfileName string, topicSyntax string, subscribeExceptionTopic string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionRequest {
	return MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionRequest{
		ApiService:              a,
		ctx:                     ctx,
		msgVpnName:              msgVpnName,
		aclProfileName:          aclProfileName,
		topicSyntax:             topicSyntax,
		subscribeExceptionTopic: subscribeExceptionTopic,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfileSubscribeExceptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeExceptionExecute(r MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionRequest) (MsgVpnAclProfileSubscribeExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfileSubscribeException")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicSyntax"+"}", _neturl.PathEscape(parameterToString(r.topicSyntax, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscribeExceptionTopic"+"}", _neturl.PathEscape(parameterToString(r.subscribeExceptionTopic, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *MsgVpnApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) Execute() (MsgVpnAclProfileSubscribeExceptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfileSubscribeExceptionsExecute(r)
}

/*
 * GetMsgVpnAclProfileSubscribeExceptions Get a list of Subscribe Topic Exception objects.
 * Get a list of Subscribe Topic Exception objects.

A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|x
msgVpnName|x|x
subscribeExceptionTopic|x|x
topicSyntax|x|x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been deprecated since 2.14. Replaced by subscribeTopicExceptions.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @return MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest {
	return MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest{
		ApiService:     a,
		ctx:            ctx,
		msgVpnName:     msgVpnName,
		aclProfileName: aclProfileName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfileSubscribeExceptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeExceptionsExecute(r MsgVpnApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) (MsgVpnAclProfileSubscribeExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfileSubscribeExceptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest struct {
	ctx                               _context.Context
	ApiService                        *MsgVpnApiService
	msgVpnName                        string
	aclProfileName                    string
	subscribeShareNameExceptionSyntax string
	subscribeShareNameException       string
	select_                           *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest) Execute() (MsgVpnAclProfileSubscribeShareNameExceptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfileSubscribeShareNameExceptionExecute(r)
}

/*
 * GetMsgVpnAclProfileSubscribeShareNameException Get a Subscribe Share Name Exception object.
 * Get a Subscribe Share Name Exception object.

A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|
msgVpnName|x|
subscribeShareNameException|x|
subscribeShareNameExceptionSyntax|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.14.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @param subscribeShareNameExceptionSyntax The syntax of the subscribe share name for the exception to the default action taken.
 * @param subscribeShareNameException The subscribe share name exception to the default action taken. May include wildcard characters.
 * @return MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeShareNameException(ctx _context.Context, msgVpnName string, aclProfileName string, subscribeShareNameExceptionSyntax string, subscribeShareNameException string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest {
	return MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest{
		ApiService:                        a,
		ctx:                               ctx,
		msgVpnName:                        msgVpnName,
		aclProfileName:                    aclProfileName,
		subscribeShareNameExceptionSyntax: subscribeShareNameExceptionSyntax,
		subscribeShareNameException:       subscribeShareNameException,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfileSubscribeShareNameExceptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeShareNameExceptionExecute(r MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest) (MsgVpnAclProfileSubscribeShareNameExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeShareNameExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfileSubscribeShareNameException")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscribeShareNameExceptionSyntax"+"}", _neturl.PathEscape(parameterToString(r.subscribeShareNameExceptionSyntax, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscribeShareNameException"+"}", _neturl.PathEscape(parameterToString(r.subscribeShareNameException, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *MsgVpnApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) Execute() (MsgVpnAclProfileSubscribeShareNameExceptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfileSubscribeShareNameExceptionsExecute(r)
}

/*
 * GetMsgVpnAclProfileSubscribeShareNameExceptions Get a list of Subscribe Share Name Exception objects.
 * Get a list of Subscribe Share Name Exception objects.

A Subscribe Share Name Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|
msgVpnName|x|
subscribeShareNameException|x|
subscribeShareNameExceptionSyntax|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.14.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @return MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeShareNameExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest {
	return MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest{
		ApiService:     a,
		ctx:            ctx,
		msgVpnName:     msgVpnName,
		aclProfileName: aclProfileName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfileSubscribeShareNameExceptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeShareNameExceptionsExecute(r MsgVpnApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) (MsgVpnAclProfileSubscribeShareNameExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeShareNameExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfileSubscribeShareNameExceptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest struct {
	ctx                           _context.Context
	ApiService                    *MsgVpnApiService
	msgVpnName                    string
	aclProfileName                string
	subscribeTopicExceptionSyntax string
	subscribeTopicException       string
	select_                       *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest) Execute() (MsgVpnAclProfileSubscribeTopicExceptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfileSubscribeTopicExceptionExecute(r)
}

/*
 * GetMsgVpnAclProfileSubscribeTopicException Get a Subscribe Topic Exception object.
 * Get a Subscribe Topic Exception object.

A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|
msgVpnName|x|
subscribeTopicException|x|
subscribeTopicExceptionSyntax|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.14.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @param subscribeTopicExceptionSyntax The syntax of the topic for the exception to the default action taken.
 * @param subscribeTopicException The topic for the exception to the default action taken. May include wildcard characters.
 * @return MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeTopicException(ctx _context.Context, msgVpnName string, aclProfileName string, subscribeTopicExceptionSyntax string, subscribeTopicException string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest {
	return MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest{
		ApiService:                    a,
		ctx:                           ctx,
		msgVpnName:                    msgVpnName,
		aclProfileName:                aclProfileName,
		subscribeTopicExceptionSyntax: subscribeTopicExceptionSyntax,
		subscribeTopicException:       subscribeTopicException,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfileSubscribeTopicExceptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeTopicExceptionExecute(r MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest) (MsgVpnAclProfileSubscribeTopicExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeTopicExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfileSubscribeTopicException")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscribeTopicExceptionSyntax"+"}", _neturl.PathEscape(parameterToString(r.subscribeTopicExceptionSyntax, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscribeTopicException"+"}", _neturl.PathEscape(parameterToString(r.subscribeTopicException, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *MsgVpnApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) Execute() (MsgVpnAclProfileSubscribeTopicExceptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfileSubscribeTopicExceptionsExecute(r)
}

/*
 * GetMsgVpnAclProfileSubscribeTopicExceptions Get a list of Subscribe Topic Exception objects.
 * Get a list of Subscribe Topic Exception objects.

A Subscribe Topic Exception is an exception to the default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|
msgVpnName|x|
subscribeTopicException|x|
subscribeTopicExceptionSyntax|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.14.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param aclProfileName The name of the ACL Profile.
 * @return MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeTopicExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest {
	return MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest{
		ApiService:     a,
		ctx:            ctx,
		msgVpnName:     msgVpnName,
		aclProfileName: aclProfileName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfileSubscribeTopicExceptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfileSubscribeTopicExceptionsExecute(r MsgVpnApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) (MsgVpnAclProfileSubscribeTopicExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeTopicExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfileSubscribeTopicExceptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"aclProfileName"+"}", _neturl.PathEscape(parameterToString(r.aclProfileName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAclProfilesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnAclProfilesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnAclProfilesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfilesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnAclProfilesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfilesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnAclProfilesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnAclProfilesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAclProfilesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAclProfilesRequest) Execute() (MsgVpnAclProfilesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAclProfilesExecute(r)
}

/*
 * GetMsgVpnAclProfiles Get a list of ACL Profile objects.
 * Get a list of ACL Profile objects.

An ACL Profile controls whether an authenticated client is permitted to establish a connection with the message broker or permitted to publish and subscribe to specific topics.


Attribute|Identifying|Deprecated
:---|:---:|:---:
aclProfileName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnAclProfilesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAclProfiles(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnAclProfilesRequest {
	return MsgVpnApiApiGetMsgVpnAclProfilesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfilesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAclProfilesExecute(r MsgVpnApiApiGetMsgVpnAclProfilesRequest) (MsgVpnAclProfilesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAclProfiles")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/aclProfiles"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAuthenticationOauthProviderRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	oauthProviderName string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnAuthenticationOauthProviderRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAuthenticationOauthProviderRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAuthenticationOauthProviderRequest) Execute() (MsgVpnAuthenticationOauthProviderResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAuthenticationOauthProviderExecute(r)
}

/*
 * GetMsgVpnAuthenticationOauthProvider Get an OAuth Provider object.
 * Get an OAuth Provider object.

OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
oauthProviderName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.13.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param oauthProviderName The name of the OAuth Provider.
 * @return MsgVpnApiApiGetMsgVpnAuthenticationOauthProviderRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAuthenticationOauthProvider(ctx _context.Context, msgVpnName string, oauthProviderName string) MsgVpnApiApiGetMsgVpnAuthenticationOauthProviderRequest {
	return MsgVpnApiApiGetMsgVpnAuthenticationOauthProviderRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		oauthProviderName: oauthProviderName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAuthenticationOauthProviderResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAuthenticationOauthProviderExecute(r MsgVpnApiApiGetMsgVpnAuthenticationOauthProviderRequest) (MsgVpnAuthenticationOauthProviderResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAuthenticationOauthProviderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAuthenticationOauthProvider")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/authenticationOauthProviders/{oauthProviderName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"oauthProviderName"+"}", _neturl.PathEscape(parameterToString(r.oauthProviderName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest) Count(count int32) MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest) Where(where []string) MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest) Execute() (MsgVpnAuthenticationOauthProvidersResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAuthenticationOauthProvidersExecute(r)
}

/*
 * GetMsgVpnAuthenticationOauthProviders Get a list of OAuth Provider objects.
 * Get a list of OAuth Provider objects.

OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
oauthProviderName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.13.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAuthenticationOauthProviders(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest {
	return MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAuthenticationOauthProvidersResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAuthenticationOauthProvidersExecute(r MsgVpnApiApiGetMsgVpnAuthenticationOauthProvidersRequest) (MsgVpnAuthenticationOauthProvidersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAuthenticationOauthProvidersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAuthenticationOauthProviders")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/authenticationOauthProviders"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAuthorizationGroupRequest struct {
	ctx                    _context.Context
	ApiService             *MsgVpnApiService
	msgVpnName             string
	authorizationGroupName string
	select_                *[]string
}

func (r MsgVpnApiApiGetMsgVpnAuthorizationGroupRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAuthorizationGroupRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAuthorizationGroupRequest) Execute() (MsgVpnAuthorizationGroupResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAuthorizationGroupExecute(r)
}

/*
 * GetMsgVpnAuthorizationGroup Get an LDAP Authorization Group object.
 * Get an LDAP Authorization Group object.

To use client authorization groups configured on an external LDAP server to provide client authorizations, LDAP Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the LDAP server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.


Attribute|Identifying|Deprecated
:---|:---:|:---:
authorizationGroupName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param authorizationGroupName The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\\#,lab,com'.
 * @return MsgVpnApiApiGetMsgVpnAuthorizationGroupRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAuthorizationGroup(ctx _context.Context, msgVpnName string, authorizationGroupName string) MsgVpnApiApiGetMsgVpnAuthorizationGroupRequest {
	return MsgVpnApiApiGetMsgVpnAuthorizationGroupRequest{
		ApiService:             a,
		ctx:                    ctx,
		msgVpnName:             msgVpnName,
		authorizationGroupName: authorizationGroupName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAuthorizationGroupResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAuthorizationGroupExecute(r MsgVpnApiApiGetMsgVpnAuthorizationGroupRequest) (MsgVpnAuthorizationGroupResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAuthorizationGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAuthorizationGroup")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"authorizationGroupName"+"}", _neturl.PathEscape(parameterToString(r.authorizationGroupName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest) Execute() (MsgVpnAuthorizationGroupsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnAuthorizationGroupsExecute(r)
}

/*
 * GetMsgVpnAuthorizationGroups Get a list of LDAP Authorization Group objects.
 * Get a list of LDAP Authorization Group objects.

To use client authorization groups configured on an external LDAP server to provide client authorizations, LDAP Authorization Group objects must be created on the Message VPN that match the authorization groups provisioned on the LDAP server. These objects must be configured with the client profiles and ACL profiles that will be assigned to the clients that belong to those authorization groups. A newly created group is placed at the end of the group list which is the lowest priority.


Attribute|Identifying|Deprecated
:---|:---:|:---:
authorizationGroupName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnAuthorizationGroups(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest {
	return MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAuthorizationGroupsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnAuthorizationGroupsExecute(r MsgVpnApiApiGetMsgVpnAuthorizationGroupsRequest) (MsgVpnAuthorizationGroupsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAuthorizationGroupsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnAuthorizationGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/authorizationGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnBridgeRequest struct {
	ctx                 _context.Context
	ApiService          *MsgVpnApiService
	msgVpnName          string
	bridgeName          string
	bridgeVirtualRouter string
	select_             *[]string
}

func (r MsgVpnApiApiGetMsgVpnBridgeRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnBridgeRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnBridgeRequest) Execute() (MsgVpnBridgeResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnBridgeExecute(r)
}

/*
 * GetMsgVpnBridge Get a Bridge object.
 * Get a Bridge object.

Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgeName|x|
bridgeVirtualRouter|x|
counter.controlRxByteCount||x
counter.controlRxMsgCount||x
counter.controlTxByteCount||x
counter.controlTxMsgCount||x
counter.dataRxByteCount||x
counter.dataRxMsgCount||x
counter.dataTxByteCount||x
counter.dataTxMsgCount||x
counter.discardedRxMsgCount||x
counter.discardedTxMsgCount||x
counter.loginRxMsgCount||x
counter.loginTxMsgCount||x
counter.msgSpoolRxMsgCount||x
counter.rxByteCount||x
counter.rxMsgCount||x
counter.txByteCount||x
counter.txMsgCount||x
msgVpnName|x|
rate.averageRxByteRate||x
rate.averageRxMsgRate||x
rate.averageTxByteRate||x
rate.averageTxMsgRate||x
rate.rxByteRate||x
rate.rxMsgRate||x
rate.txByteRate||x
rate.txMsgRate||x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param bridgeName The name of the Bridge.
 * @param bridgeVirtualRouter The virtual router of the Bridge.
 * @return MsgVpnApiApiGetMsgVpnBridgeRequest
*/
func (a *MsgVpnApiService) GetMsgVpnBridge(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string) MsgVpnApiApiGetMsgVpnBridgeRequest {
	return MsgVpnApiApiGetMsgVpnBridgeRequest{
		ApiService:          a,
		ctx:                 ctx,
		msgVpnName:          msgVpnName,
		bridgeName:          bridgeName,
		bridgeVirtualRouter: bridgeVirtualRouter,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgeResponse
 */
func (a *MsgVpnApiService) GetMsgVpnBridgeExecute(r MsgVpnApiApiGetMsgVpnBridgeRequest) (MsgVpnBridgeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnBridge")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", _neturl.PathEscape(parameterToString(r.bridgeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.bridgeVirtualRouter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionRequest struct {
	ctx                    _context.Context
	ApiService             *MsgVpnApiService
	msgVpnName             string
	bridgeName             string
	bridgeVirtualRouter    string
	localSubscriptionTopic string
	select_                *[]string
}

func (r MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionRequest) Execute() (MsgVpnBridgeLocalSubscriptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnBridgeLocalSubscriptionExecute(r)
}

/*
 * GetMsgVpnBridgeLocalSubscription Get a Bridge Local Subscriptions object.
 * Get a Bridge Local Subscriptions object.

A Local Subscription is a topic subscription used by a remote Message VPN Bridge to attract messages from this broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgeName|x|
bridgeVirtualRouter|x|
localSubscriptionTopic|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param bridgeName The name of the Bridge.
 * @param bridgeVirtualRouter The virtual router of the Bridge.
 * @param localSubscriptionTopic The topic of the Bridge local subscription.
 * @return MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnBridgeLocalSubscription(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string, localSubscriptionTopic string) MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionRequest {
	return MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionRequest{
		ApiService:             a,
		ctx:                    ctx,
		msgVpnName:             msgVpnName,
		bridgeName:             bridgeName,
		bridgeVirtualRouter:    bridgeVirtualRouter,
		localSubscriptionTopic: localSubscriptionTopic,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgeLocalSubscriptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnBridgeLocalSubscriptionExecute(r MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionRequest) (MsgVpnBridgeLocalSubscriptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeLocalSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnBridgeLocalSubscription")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions/{localSubscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", _neturl.PathEscape(parameterToString(r.bridgeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.bridgeVirtualRouter, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"localSubscriptionTopic"+"}", _neturl.PathEscape(parameterToString(r.localSubscriptionTopic, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest struct {
	ctx                 _context.Context
	ApiService          *MsgVpnApiService
	msgVpnName          string
	bridgeName          string
	bridgeVirtualRouter string
	count               *int32
	cursor              *string
	where               *[]string
	select_             *[]string
}

func (r MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) Execute() (MsgVpnBridgeLocalSubscriptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnBridgeLocalSubscriptionsExecute(r)
}

/*
 * GetMsgVpnBridgeLocalSubscriptions Get a list of Bridge Local Subscriptions objects.
 * Get a list of Bridge Local Subscriptions objects.

A Local Subscription is a topic subscription used by a remote Message VPN Bridge to attract messages from this broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgeName|x|
bridgeVirtualRouter|x|
localSubscriptionTopic|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param bridgeName The name of the Bridge.
 * @param bridgeVirtualRouter The virtual router of the Bridge.
 * @return MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnBridgeLocalSubscriptions(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string) MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest {
	return MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest{
		ApiService:          a,
		ctx:                 ctx,
		msgVpnName:          msgVpnName,
		bridgeName:          bridgeName,
		bridgeVirtualRouter: bridgeVirtualRouter,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgeLocalSubscriptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnBridgeLocalSubscriptionsExecute(r MsgVpnApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) (MsgVpnBridgeLocalSubscriptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeLocalSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnBridgeLocalSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/localSubscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", _neturl.PathEscape(parameterToString(r.bridgeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.bridgeVirtualRouter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	bridgeName            string
	bridgeVirtualRouter   string
	remoteMsgVpnName      string
	remoteMsgVpnLocation  string
	remoteMsgVpnInterface string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnRequest) Execute() (MsgVpnBridgeRemoteMsgVpnResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnBridgeRemoteMsgVpnExecute(r)
}

/*
 * GetMsgVpnBridgeRemoteMsgVpn Get a Remote Message VPN object.
 * Get a Remote Message VPN object.

The Remote Message VPN is the Message VPN that the Bridge connects to.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgeName|x|
bridgeVirtualRouter|x|
msgVpnName|x|
remoteMsgVpnInterface|x|
remoteMsgVpnLocation|x|
remoteMsgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param bridgeName The name of the Bridge.
 * @param bridgeVirtualRouter The virtual router of the Bridge.
 * @param remoteMsgVpnName The name of the remote Message VPN.
 * @param remoteMsgVpnLocation The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \"v:\").
 * @param remoteMsgVpnInterface The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remoteMsgVpnLocation` must not be a virtual router name.
 * @return MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnRequest
*/
func (a *MsgVpnApiService) GetMsgVpnBridgeRemoteMsgVpn(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string, remoteMsgVpnName string, remoteMsgVpnLocation string, remoteMsgVpnInterface string) MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnRequest {
	return MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		bridgeName:            bridgeName,
		bridgeVirtualRouter:   bridgeVirtualRouter,
		remoteMsgVpnName:      remoteMsgVpnName,
		remoteMsgVpnLocation:  remoteMsgVpnLocation,
		remoteMsgVpnInterface: remoteMsgVpnInterface,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgeRemoteMsgVpnResponse
 */
func (a *MsgVpnApiService) GetMsgVpnBridgeRemoteMsgVpnExecute(r MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnRequest) (MsgVpnBridgeRemoteMsgVpnResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeRemoteMsgVpnResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnBridgeRemoteMsgVpn")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", _neturl.PathEscape(parameterToString(r.bridgeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.bridgeVirtualRouter, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnName"+"}", _neturl.PathEscape(parameterToString(r.remoteMsgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnLocation"+"}", _neturl.PathEscape(parameterToString(r.remoteMsgVpnLocation, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteMsgVpnInterface"+"}", _neturl.PathEscape(parameterToString(r.remoteMsgVpnInterface, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest struct {
	ctx                 _context.Context
	ApiService          *MsgVpnApiService
	msgVpnName          string
	bridgeName          string
	bridgeVirtualRouter string
	where               *[]string
	select_             *[]string
}

func (r MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest) Execute() (MsgVpnBridgeRemoteMsgVpnsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnBridgeRemoteMsgVpnsExecute(r)
}

/*
 * GetMsgVpnBridgeRemoteMsgVpns Get a list of Remote Message VPN objects.
 * Get a list of Remote Message VPN objects.

The Remote Message VPN is the Message VPN that the Bridge connects to.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgeName|x|
bridgeVirtualRouter|x|
msgVpnName|x|
remoteMsgVpnInterface|x|
remoteMsgVpnLocation|x|
remoteMsgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param bridgeName The name of the Bridge.
 * @param bridgeVirtualRouter The virtual router of the Bridge.
 * @return MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnBridgeRemoteMsgVpns(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string) MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest {
	return MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest{
		ApiService:          a,
		ctx:                 ctx,
		msgVpnName:          msgVpnName,
		bridgeName:          bridgeName,
		bridgeVirtualRouter: bridgeVirtualRouter,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgeRemoteMsgVpnsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnBridgeRemoteMsgVpnsExecute(r MsgVpnApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest) (MsgVpnBridgeRemoteMsgVpnsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeRemoteMsgVpnsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnBridgeRemoteMsgVpns")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", _neturl.PathEscape(parameterToString(r.bridgeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.bridgeVirtualRouter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionRequest struct {
	ctx                     _context.Context
	ApiService              *MsgVpnApiService
	msgVpnName              string
	bridgeName              string
	bridgeVirtualRouter     string
	remoteSubscriptionTopic string
	select_                 *[]string
}

func (r MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionRequest) Execute() (MsgVpnBridgeRemoteSubscriptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnBridgeRemoteSubscriptionExecute(r)
}

/*
 * GetMsgVpnBridgeRemoteSubscription Get a Remote Subscription object.
 * Get a Remote Subscription object.

A Remote Subscription is a topic subscription used by the Message VPN Bridge to attract messages from the remote message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgeName|x|
bridgeVirtualRouter|x|
msgVpnName|x|
remoteSubscriptionTopic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param bridgeName The name of the Bridge.
 * @param bridgeVirtualRouter The virtual router of the Bridge.
 * @param remoteSubscriptionTopic The topic of the Bridge remote subscription.
 * @return MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnBridgeRemoteSubscription(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string, remoteSubscriptionTopic string) MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionRequest {
	return MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionRequest{
		ApiService:              a,
		ctx:                     ctx,
		msgVpnName:              msgVpnName,
		bridgeName:              bridgeName,
		bridgeVirtualRouter:     bridgeVirtualRouter,
		remoteSubscriptionTopic: remoteSubscriptionTopic,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgeRemoteSubscriptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnBridgeRemoteSubscriptionExecute(r MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionRequest) (MsgVpnBridgeRemoteSubscriptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeRemoteSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnBridgeRemoteSubscription")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", _neturl.PathEscape(parameterToString(r.bridgeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.bridgeVirtualRouter, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteSubscriptionTopic"+"}", _neturl.PathEscape(parameterToString(r.remoteSubscriptionTopic, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest struct {
	ctx                 _context.Context
	ApiService          *MsgVpnApiService
	msgVpnName          string
	bridgeName          string
	bridgeVirtualRouter string
	count               *int32
	cursor              *string
	where               *[]string
	select_             *[]string
}

func (r MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) Execute() (MsgVpnBridgeRemoteSubscriptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnBridgeRemoteSubscriptionsExecute(r)
}

/*
 * GetMsgVpnBridgeRemoteSubscriptions Get a list of Remote Subscription objects.
 * Get a list of Remote Subscription objects.

A Remote Subscription is a topic subscription used by the Message VPN Bridge to attract messages from the remote message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgeName|x|
bridgeVirtualRouter|x|
msgVpnName|x|
remoteSubscriptionTopic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param bridgeName The name of the Bridge.
 * @param bridgeVirtualRouter The virtual router of the Bridge.
 * @return MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnBridgeRemoteSubscriptions(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string) MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest {
	return MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest{
		ApiService:          a,
		ctx:                 ctx,
		msgVpnName:          msgVpnName,
		bridgeName:          bridgeName,
		bridgeVirtualRouter: bridgeVirtualRouter,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgeRemoteSubscriptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnBridgeRemoteSubscriptionsExecute(r MsgVpnApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) (MsgVpnBridgeRemoteSubscriptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeRemoteSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnBridgeRemoteSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", _neturl.PathEscape(parameterToString(r.bridgeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.bridgeVirtualRouter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest struct {
	ctx                  _context.Context
	ApiService           *MsgVpnApiService
	msgVpnName           string
	bridgeName           string
	bridgeVirtualRouter  string
	tlsTrustedCommonName string
	select_              *[]string
}

func (r MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest) Execute() (MsgVpnBridgeTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnBridgeTlsTrustedCommonNameExecute(r)
}

/*
 * GetMsgVpnBridgeTlsTrustedCommonName Get a Trusted Common Name object.
 * Get a Trusted Common Name object.

The Trusted Common Names for the Bridge are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgeName|x|x
bridgeVirtualRouter|x|x
msgVpnName|x|x
tlsTrustedCommonName|x|x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param bridgeName The name of the Bridge.
 * @param bridgeVirtualRouter The virtual router of the Bridge.
 * @param tlsTrustedCommonName The expected trusted common name of the remote certificate.
 * @return MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest
*/
func (a *MsgVpnApiService) GetMsgVpnBridgeTlsTrustedCommonName(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string, tlsTrustedCommonName string) MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest {
	return MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest{
		ApiService:           a,
		ctx:                  ctx,
		msgVpnName:           msgVpnName,
		bridgeName:           bridgeName,
		bridgeVirtualRouter:  bridgeVirtualRouter,
		tlsTrustedCommonName: tlsTrustedCommonName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgeTlsTrustedCommonNameResponse
 */
func (a *MsgVpnApiService) GetMsgVpnBridgeTlsTrustedCommonNameExecute(r MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest) (MsgVpnBridgeTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeTlsTrustedCommonNameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnBridgeTlsTrustedCommonName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", _neturl.PathEscape(parameterToString(r.bridgeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.bridgeVirtualRouter, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tlsTrustedCommonName"+"}", _neturl.PathEscape(parameterToString(r.tlsTrustedCommonName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest struct {
	ctx                 _context.Context
	ApiService          *MsgVpnApiService
	msgVpnName          string
	bridgeName          string
	bridgeVirtualRouter string
	where               *[]string
	select_             *[]string
}

func (r MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest) Execute() (MsgVpnBridgeTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnBridgeTlsTrustedCommonNamesExecute(r)
}

/*
 * GetMsgVpnBridgeTlsTrustedCommonNames Get a list of Trusted Common Name objects.
 * Get a list of Trusted Common Name objects.

The Trusted Common Names for the Bridge are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgeName|x|x
bridgeVirtualRouter|x|x
msgVpnName|x|x
tlsTrustedCommonName|x|x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param bridgeName The name of the Bridge.
 * @param bridgeVirtualRouter The virtual router of the Bridge.
 * @return MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnBridgeTlsTrustedCommonNames(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string) MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest {
	return MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest{
		ApiService:          a,
		ctx:                 ctx,
		msgVpnName:          msgVpnName,
		bridgeName:          bridgeName,
		bridgeVirtualRouter: bridgeVirtualRouter,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgeTlsTrustedCommonNamesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnBridgeTlsTrustedCommonNamesExecute(r MsgVpnApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest) (MsgVpnBridgeTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeTlsTrustedCommonNamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnBridgeTlsTrustedCommonNames")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeName"+"}", _neturl.PathEscape(parameterToString(r.bridgeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bridgeVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.bridgeVirtualRouter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnBridgesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnBridgesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnBridgesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnBridgesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnBridgesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnBridgesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnBridgesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnBridgesRequest) Execute() (MsgVpnBridgesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnBridgesExecute(r)
}

/*
 * GetMsgVpnBridges Get a list of Bridge objects.
 * Get a list of Bridge objects.

Bridges can be used to link two Message VPNs so that messages published to one Message VPN that match the topic subscriptions set for the bridge are also delivered to the linked Message VPN.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgeName|x|
bridgeVirtualRouter|x|
counter.controlRxByteCount||x
counter.controlRxMsgCount||x
counter.controlTxByteCount||x
counter.controlTxMsgCount||x
counter.dataRxByteCount||x
counter.dataRxMsgCount||x
counter.dataTxByteCount||x
counter.dataTxMsgCount||x
counter.discardedRxMsgCount||x
counter.discardedTxMsgCount||x
counter.loginRxMsgCount||x
counter.loginTxMsgCount||x
counter.msgSpoolRxMsgCount||x
counter.rxByteCount||x
counter.rxMsgCount||x
counter.txByteCount||x
counter.txMsgCount||x
msgVpnName|x|
rate.averageRxByteRate||x
rate.averageRxMsgRate||x
rate.averageTxByteRate||x
rate.averageTxMsgRate||x
rate.rxByteRate||x
rate.rxMsgRate||x
rate.txByteRate||x
rate.txMsgRate||x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnBridgesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnBridges(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnBridgesRequest {
	return MsgVpnApiApiGetMsgVpnBridgesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnBridgesExecute(r MsgVpnApiApiGetMsgVpnBridgesRequest) (MsgVpnBridgesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnBridges")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/bridges"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	clientName string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientRequest) Execute() (MsgVpnClientResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientExecute(r)
}

/*
 * GetMsgVpnClient Get a Client object.
 * Get a Client object.

Applications or devices that connect to message brokers to send and/or receive messages are represented as Clients.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @return MsgVpnApiApiGetMsgVpnClientRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClient(ctx _context.Context, msgVpnName string, clientName string) MsgVpnApiApiGetMsgVpnClientRequest {
	return MsgVpnApiApiGetMsgVpnClientRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		clientName: clientName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientExecute(r MsgVpnApiApiGetMsgVpnClientRequest) (MsgVpnClientResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClient")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientConnectionRequest struct {
	ctx           _context.Context
	ApiService    *MsgVpnApiService
	msgVpnName    string
	clientName    string
	clientAddress string
	select_       *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientConnectionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientConnectionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientConnectionRequest) Execute() (MsgVpnClientConnectionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientConnectionExecute(r)
}

/*
 * GetMsgVpnClientConnection Get a Client Connection object.
 * Get a Client Connection object.

A Client Connection represents the Transmission Control Protocol (TCP) connection the Client uses to communicate with the message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientAddress|x|
clientName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @param clientAddress The IP address and TCP port on the Client side of the Client Connection.
 * @return MsgVpnApiApiGetMsgVpnClientConnectionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientConnection(ctx _context.Context, msgVpnName string, clientName string, clientAddress string) MsgVpnApiApiGetMsgVpnClientConnectionRequest {
	return MsgVpnApiApiGetMsgVpnClientConnectionRequest{
		ApiService:    a,
		ctx:           ctx,
		msgVpnName:    msgVpnName,
		clientName:    clientName,
		clientAddress: clientAddress,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientConnectionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientConnectionExecute(r MsgVpnApiApiGetMsgVpnClientConnectionRequest) (MsgVpnClientConnectionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientConnectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientConnection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}/connections/{clientAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientAddress"+"}", _neturl.PathEscape(parameterToString(r.clientAddress, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientConnectionsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	clientName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientConnectionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnClientConnectionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientConnectionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnClientConnectionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientConnectionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnClientConnectionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientConnectionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientConnectionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientConnectionsRequest) Execute() (MsgVpnClientConnectionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientConnectionsExecute(r)
}

/*
 * GetMsgVpnClientConnections Get a list of Client Connection objects.
 * Get a list of Client Connection objects.

A Client Connection represents the Transmission Control Protocol (TCP) connection the Client uses to communicate with the message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientAddress|x|
clientName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @return MsgVpnApiApiGetMsgVpnClientConnectionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientConnections(ctx _context.Context, msgVpnName string, clientName string) MsgVpnApiApiGetMsgVpnClientConnectionsRequest {
	return MsgVpnApiApiGetMsgVpnClientConnectionsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		clientName: clientName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientConnectionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientConnectionsExecute(r MsgVpnApiApiGetMsgVpnClientConnectionsRequest) (MsgVpnClientConnectionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientConnectionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientConnections")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}/connections"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientProfileRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	clientProfileName string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientProfileRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientProfileRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientProfileRequest) Execute() (MsgVpnClientProfileResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientProfileExecute(r)
}

/*
 * GetMsgVpnClientProfile Get a Client Profile object.
 * Get a Client Profile object.

Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.


Attribute|Identifying|Deprecated
:---|:---:|:---:
apiQueueManagementCopyFromOnCreateName||x
apiTopicEndpointManagementCopyFromOnCreateName||x
clientProfileName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientProfileName The name of the Client Profile.
 * @return MsgVpnApiApiGetMsgVpnClientProfileRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientProfile(ctx _context.Context, msgVpnName string, clientProfileName string) MsgVpnApiApiGetMsgVpnClientProfileRequest {
	return MsgVpnApiApiGetMsgVpnClientProfileRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		clientProfileName: clientProfileName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientProfileResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientProfileExecute(r MsgVpnApiApiGetMsgVpnClientProfileRequest) (MsgVpnClientProfileResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientProfileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientProfile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clientProfiles/{clientProfileName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientProfileName"+"}", _neturl.PathEscape(parameterToString(r.clientProfileName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientProfilesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientProfilesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnClientProfilesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientProfilesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnClientProfilesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientProfilesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnClientProfilesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientProfilesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientProfilesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientProfilesRequest) Execute() (MsgVpnClientProfilesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientProfilesExecute(r)
}

/*
 * GetMsgVpnClientProfiles Get a list of Client Profile objects.
 * Get a list of Client Profile objects.

Client Profiles are used to assign common configuration properties to clients that have been successfully authorized.


Attribute|Identifying|Deprecated
:---|:---:|:---:
apiQueueManagementCopyFromOnCreateName||x
apiTopicEndpointManagementCopyFromOnCreateName||x
clientProfileName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnClientProfilesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientProfiles(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnClientProfilesRequest {
	return MsgVpnApiApiGetMsgVpnClientProfilesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientProfilesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientProfilesExecute(r MsgVpnApiApiGetMsgVpnClientProfilesRequest) (MsgVpnClientProfilesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientProfilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientProfiles")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clientProfiles"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientRxFlowRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	clientName string
	flowId     string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientRxFlowRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientRxFlowRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientRxFlowRequest) Execute() (MsgVpnClientRxFlowResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientRxFlowExecute(r)
}

/*
 * GetMsgVpnClientRxFlow Get a Client Receive Flow object.
 * Get a Client Receive Flow object.

Client Receive Flows are used by clients to publish Guaranteed messages to a message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientName|x|
flowId|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @param flowId The identifier (ID) of the flow.
 * @return MsgVpnApiApiGetMsgVpnClientRxFlowRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientRxFlow(ctx _context.Context, msgVpnName string, clientName string, flowId string) MsgVpnApiApiGetMsgVpnClientRxFlowRequest {
	return MsgVpnApiApiGetMsgVpnClientRxFlowRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		clientName: clientName,
		flowId:     flowId,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientRxFlowResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientRxFlowExecute(r MsgVpnApiApiGetMsgVpnClientRxFlowRequest) (MsgVpnClientRxFlowResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientRxFlowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientRxFlow")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}/rxFlows/{flowId}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"flowId"+"}", _neturl.PathEscape(parameterToString(r.flowId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientRxFlowsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	clientName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientRxFlowsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnClientRxFlowsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientRxFlowsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnClientRxFlowsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientRxFlowsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnClientRxFlowsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientRxFlowsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientRxFlowsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientRxFlowsRequest) Execute() (MsgVpnClientRxFlowsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientRxFlowsExecute(r)
}

/*
 * GetMsgVpnClientRxFlows Get a list of Client Receive Flow objects.
 * Get a list of Client Receive Flow objects.

Client Receive Flows are used by clients to publish Guaranteed messages to a message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientName|x|
flowId|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @return MsgVpnApiApiGetMsgVpnClientRxFlowsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientRxFlows(ctx _context.Context, msgVpnName string, clientName string) MsgVpnApiApiGetMsgVpnClientRxFlowsRequest {
	return MsgVpnApiApiGetMsgVpnClientRxFlowsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		clientName: clientName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientRxFlowsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientRxFlowsExecute(r MsgVpnApiApiGetMsgVpnClientRxFlowsRequest) (MsgVpnClientRxFlowsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientRxFlowsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientRxFlows")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}/rxFlows"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientSubscriptionRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	clientName        string
	subscriptionTopic string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientSubscriptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientSubscriptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientSubscriptionRequest) Execute() (MsgVpnClientSubscriptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientSubscriptionExecute(r)
}

/*
 * GetMsgVpnClientSubscription Get a Client Subscription object.
 * Get a Client Subscription object.

Once clients are authenticated on the message broker they can add and remove Client Subscriptions for Direct messages published to the Message VPN to which they have connected.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientName|x|
msgVpnName|x|
subscriptionTopic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @param subscriptionTopic The topic of the Subscription.
 * @return MsgVpnApiApiGetMsgVpnClientSubscriptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientSubscription(ctx _context.Context, msgVpnName string, clientName string, subscriptionTopic string) MsgVpnApiApiGetMsgVpnClientSubscriptionRequest {
	return MsgVpnApiApiGetMsgVpnClientSubscriptionRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		clientName:        clientName,
		subscriptionTopic: subscriptionTopic,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientSubscriptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientSubscriptionExecute(r MsgVpnApiApiGetMsgVpnClientSubscriptionRequest) (MsgVpnClientSubscriptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientSubscription")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}/subscriptions/{subscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionTopic"+"}", _neturl.PathEscape(parameterToString(r.subscriptionTopic, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	clientName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest) Execute() (MsgVpnClientSubscriptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientSubscriptionsExecute(r)
}

/*
 * GetMsgVpnClientSubscriptions Get a list of Client Subscription objects.
 * Get a list of Client Subscription objects.

Once clients are authenticated on the message broker they can add and remove Client Subscriptions for Direct messages published to the Message VPN to which they have connected.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientName|x|
msgVpnName|x|
subscriptionTopic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @return MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientSubscriptions(ctx _context.Context, msgVpnName string, clientName string) MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest {
	return MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		clientName: clientName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientSubscriptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientSubscriptionsExecute(r MsgVpnApiApiGetMsgVpnClientSubscriptionsRequest) (MsgVpnClientSubscriptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientTransactedSessionRequest struct {
	ctx         _context.Context
	ApiService  *MsgVpnApiService
	msgVpnName  string
	clientName  string
	sessionName string
	select_     *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientTransactedSessionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientTransactedSessionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientTransactedSessionRequest) Execute() (MsgVpnClientTransactedSessionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientTransactedSessionExecute(r)
}

/*
 * GetMsgVpnClientTransactedSession Get a Client Transacted Session object.
 * Get a Client Transacted Session object.

Transacted Sessions enable clients to group multiple message send and/or receive operations together in single, atomic units known as local transactions.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientName|x|
msgVpnName|x|
sessionName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @param sessionName The name of the Transacted Session.
 * @return MsgVpnApiApiGetMsgVpnClientTransactedSessionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientTransactedSession(ctx _context.Context, msgVpnName string, clientName string, sessionName string) MsgVpnApiApiGetMsgVpnClientTransactedSessionRequest {
	return MsgVpnApiApiGetMsgVpnClientTransactedSessionRequest{
		ApiService:  a,
		ctx:         ctx,
		msgVpnName:  msgVpnName,
		clientName:  clientName,
		sessionName: sessionName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientTransactedSessionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientTransactedSessionExecute(r MsgVpnApiApiGetMsgVpnClientTransactedSessionRequest) (MsgVpnClientTransactedSessionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientTransactedSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientTransactedSession")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionName"+"}", _neturl.PathEscape(parameterToString(r.sessionName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	clientName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest) Execute() (MsgVpnClientTransactedSessionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientTransactedSessionsExecute(r)
}

/*
 * GetMsgVpnClientTransactedSessions Get a list of Client Transacted Session objects.
 * Get a list of Client Transacted Session objects.

Transacted Sessions enable clients to group multiple message send and/or receive operations together in single, atomic units known as local transactions.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientName|x|
msgVpnName|x|
sessionName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @return MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientTransactedSessions(ctx _context.Context, msgVpnName string, clientName string) MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest {
	return MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		clientName: clientName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientTransactedSessionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientTransactedSessionsExecute(r MsgVpnApiApiGetMsgVpnClientTransactedSessionsRequest) (MsgVpnClientTransactedSessionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientTransactedSessionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientTransactedSessions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientTxFlowRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	clientName string
	flowId     string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientTxFlowRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientTxFlowRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientTxFlowRequest) Execute() (MsgVpnClientTxFlowResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientTxFlowExecute(r)
}

/*
 * GetMsgVpnClientTxFlow Get a Client Transmit Flow object.
 * Get a Client Transmit Flow object.

Client Transmit Flows are used by clients to consume Guaranteed messages from a message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientName|x|
flowId|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @param flowId The identifier (ID) of the flow.
 * @return MsgVpnApiApiGetMsgVpnClientTxFlowRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientTxFlow(ctx _context.Context, msgVpnName string, clientName string, flowId string) MsgVpnApiApiGetMsgVpnClientTxFlowRequest {
	return MsgVpnApiApiGetMsgVpnClientTxFlowRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		clientName: clientName,
		flowId:     flowId,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientTxFlowResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientTxFlowExecute(r MsgVpnApiApiGetMsgVpnClientTxFlowRequest) (MsgVpnClientTxFlowResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientTxFlowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientTxFlow")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}/txFlows/{flowId}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"flowId"+"}", _neturl.PathEscape(parameterToString(r.flowId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientTxFlowsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	clientName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientTxFlowsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnClientTxFlowsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientTxFlowsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnClientTxFlowsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientTxFlowsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnClientTxFlowsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientTxFlowsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientTxFlowsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientTxFlowsRequest) Execute() (MsgVpnClientTxFlowsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientTxFlowsExecute(r)
}

/*
 * GetMsgVpnClientTxFlows Get a list of Client Transmit Flow objects.
 * Get a list of Client Transmit Flow objects.

Client Transmit Flows are used by clients to consume Guaranteed messages from a message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientName|x|
flowId|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientName The name of the Client.
 * @return MsgVpnApiApiGetMsgVpnClientTxFlowsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientTxFlows(ctx _context.Context, msgVpnName string, clientName string) MsgVpnApiApiGetMsgVpnClientTxFlowsRequest {
	return MsgVpnApiApiGetMsgVpnClientTxFlowsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		clientName: clientName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientTxFlowsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientTxFlowsExecute(r MsgVpnApiApiGetMsgVpnClientTxFlowsRequest) (MsgVpnClientTxFlowsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientTxFlowsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientTxFlows")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients/{clientName}/txFlows"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientName"+"}", _neturl.PathEscape(parameterToString(r.clientName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientUsernameRequest struct {
	ctx            _context.Context
	ApiService     *MsgVpnApiService
	msgVpnName     string
	clientUsername string
	select_        *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientUsernameRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientUsernameRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientUsernameRequest) Execute() (MsgVpnClientUsernameResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientUsernameExecute(r)
}

/*
 * GetMsgVpnClientUsername Get a Client Username object.
 * Get a Client Username object.

A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientUsername|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param clientUsername The name of the Client Username.
 * @return MsgVpnApiApiGetMsgVpnClientUsernameRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientUsername(ctx _context.Context, msgVpnName string, clientUsername string) MsgVpnApiApiGetMsgVpnClientUsernameRequest {
	return MsgVpnApiApiGetMsgVpnClientUsernameRequest{
		ApiService:     a,
		ctx:            ctx,
		msgVpnName:     msgVpnName,
		clientUsername: clientUsername,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientUsernameResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientUsernameExecute(r MsgVpnApiApiGetMsgVpnClientUsernameRequest) (MsgVpnClientUsernameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientUsernameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientUsername")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clientUsernames/{clientUsername}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientUsername"+"}", _neturl.PathEscape(parameterToString(r.clientUsername, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientUsernamesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientUsernamesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnClientUsernamesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientUsernamesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnClientUsernamesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientUsernamesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnClientUsernamesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientUsernamesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientUsernamesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientUsernamesRequest) Execute() (MsgVpnClientUsernamesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientUsernamesExecute(r)
}

/*
 * GetMsgVpnClientUsernames Get a list of Client Username objects.
 * Get a list of Client Username objects.

A client is only authorized to connect to a Message VPN that is associated with a Client Username that the client has been assigned.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientUsername|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnClientUsernamesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClientUsernames(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnClientUsernamesRequest {
	return MsgVpnApiApiGetMsgVpnClientUsernamesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientUsernamesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientUsernamesExecute(r MsgVpnApiApiGetMsgVpnClientUsernamesRequest) (MsgVpnClientUsernamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientUsernamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClientUsernames")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clientUsernames"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnClientsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnClientsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnClientsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnClientsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnClientsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnClientsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnClientsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnClientsRequest) Execute() (MsgVpnClientsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnClientsExecute(r)
}

/*
 * GetMsgVpnClients Get a list of Client objects.
 * Get a list of Client objects.

Applications or devices that connect to message brokers to send and/or receive messages are represented as Clients.


Attribute|Identifying|Deprecated
:---|:---:|:---:
clientName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnClientsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnClients(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnClientsRequest {
	return MsgVpnApiApiGetMsgVpnClientsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnClientsExecute(r MsgVpnApiApiGetMsgVpnClientsRequest) (MsgVpnClientsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnClients")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/clients"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodeRequest struct {
	ctx            _context.Context
	ApiService     *MsgVpnApiService
	msgVpnName     string
	remoteNodeName string
	select_        *[]string
}

func (r MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodeRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodeRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodeRequest) Execute() (MsgVpnConfigSyncRemoteNodeResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnConfigSyncRemoteNodeExecute(r)
}

/*
 * GetMsgVpnConfigSyncRemoteNode Get a Config Sync Remote Node object.
 * Get a Config Sync Remote Node object.

A Config Sync Remote Node object contains information about the status of the table for this Message VPN with respect to a remote node.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
remoteNodeName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param remoteNodeName The name of the Config Sync Remote Node.
 * @return MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodeRequest
*/
func (a *MsgVpnApiService) GetMsgVpnConfigSyncRemoteNode(ctx _context.Context, msgVpnName string, remoteNodeName string) MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodeRequest {
	return MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodeRequest{
		ApiService:     a,
		ctx:            ctx,
		msgVpnName:     msgVpnName,
		remoteNodeName: remoteNodeName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnConfigSyncRemoteNodeResponse
 */
func (a *MsgVpnApiService) GetMsgVpnConfigSyncRemoteNodeExecute(r MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodeRequest) (MsgVpnConfigSyncRemoteNodeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnConfigSyncRemoteNodeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnConfigSyncRemoteNode")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/configSyncRemoteNodes/{remoteNodeName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteNodeName"+"}", _neturl.PathEscape(parameterToString(r.remoteNodeName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest) Execute() (MsgVpnConfigSyncRemoteNodesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnConfigSyncRemoteNodesExecute(r)
}

/*
 * GetMsgVpnConfigSyncRemoteNodes Get a list of Config Sync Remote Node objects.
 * Get a list of Config Sync Remote Node objects.

A Config Sync Remote Node object contains information about the status of the table for this Message VPN with respect to a remote node.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
remoteNodeName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnConfigSyncRemoteNodes(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest {
	return MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnConfigSyncRemoteNodesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnConfigSyncRemoteNodesExecute(r MsgVpnApiApiGetMsgVpnConfigSyncRemoteNodesRequest) (MsgVpnConfigSyncRemoteNodesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnConfigSyncRemoteNodesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnConfigSyncRemoteNodes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/configSyncRemoteNodes"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	cacheName  string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheRequest) Execute() (MsgVpnDistributedCacheResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheExecute(r)
}

/*
 * GetMsgVpnDistributedCache Get a Distributed Cache object.
 * Get a Distributed Cache object.

A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCache(ctx _context.Context, msgVpnName string, cacheName string) MsgVpnApiApiGetMsgVpnDistributedCacheRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		cacheName:  cacheName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheRequest) (MsgVpnDistributedCacheResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCache")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterRequest struct {
	ctx         _context.Context
	ApiService  *MsgVpnApiService
	msgVpnName  string
	cacheName   string
	clusterName string
	select_     *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterRequest) Execute() (MsgVpnDistributedCacheClusterResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterExecute(r)
}

/*
 * GetMsgVpnDistributedCacheCluster Get a Cache Cluster object.
 * Get a Cache Cluster object.

A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheCluster(ctx _context.Context, msgVpnName string, cacheName string, clusterName string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterRequest{
		ApiService:  a,
		ctx:         ctx,
		msgVpnName:  msgVpnName,
		cacheName:   cacheName,
		clusterName: clusterName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterRequest) (MsgVpnDistributedCacheClusterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheCluster")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest struct {
	ctx             _context.Context
	ApiService      *MsgVpnApiService
	msgVpnName      string
	cacheName       string
	clusterName     string
	homeClusterName string
	select_         *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest) Execute() (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster Get a Home Cache Cluster object.
 * Get a Home Cache Cluster object.

A Home Cache Cluster is a Cache Cluster that is the "definitive" Cache Cluster for a given topic in the context of the Global Caching feature.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
homeClusterName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param homeClusterName The name of the remote Home Cache Cluster.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, homeClusterName string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest{
		ApiService:      a,
		ctx:             ctx,
		msgVpnName:      msgVpnName,
		cacheName:       cacheName,
		clusterName:     clusterName,
		homeClusterName: homeClusterName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest) (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"homeClusterName"+"}", _neturl.PathEscape(parameterToString(r.homeClusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest struct {
	ctx             _context.Context
	ApiService      *MsgVpnApiService
	msgVpnName      string
	cacheName       string
	clusterName     string
	homeClusterName string
	topicPrefix     string
	select_         *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest) Execute() (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix Get a Topic Prefix object.
 * Get a Topic Prefix object.

A Topic Prefix is a prefix for a global topic that is available from the containing Home Cache Cluster.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
homeClusterName|x|
msgVpnName|x|
topicPrefix|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param homeClusterName The name of the remote Home Cache Cluster.
 * @param topicPrefix A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/>) is implied at the end of the prefix.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, homeClusterName string, topicPrefix string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest{
		ApiService:      a,
		ctx:             ctx,
		msgVpnName:      msgVpnName,
		cacheName:       cacheName,
		clusterName:     clusterName,
		homeClusterName: homeClusterName,
		topicPrefix:     topicPrefix,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest) (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"homeClusterName"+"}", _neturl.PathEscape(parameterToString(r.homeClusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicPrefix"+"}", _neturl.PathEscape(parameterToString(r.topicPrefix, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest struct {
	ctx             _context.Context
	ApiService      *MsgVpnApiService
	msgVpnName      string
	cacheName       string
	clusterName     string
	homeClusterName string
	count           *int32
	cursor          *string
	where           *[]string
	select_         *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) Execute() (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes Get a list of Topic Prefix objects.
 * Get a list of Topic Prefix objects.

A Topic Prefix is a prefix for a global topic that is available from the containing Home Cache Cluster.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
homeClusterName|x|
msgVpnName|x|
topicPrefix|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param homeClusterName The name of the remote Home Cache Cluster.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, homeClusterName string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest{
		ApiService:      a,
		ctx:             ctx,
		msgVpnName:      msgVpnName,
		cacheName:       cacheName,
		clusterName:     clusterName,
		homeClusterName: homeClusterName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"homeClusterName"+"}", _neturl.PathEscape(parameterToString(r.homeClusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest struct {
	ctx         _context.Context
	ApiService  *MsgVpnApiService
	msgVpnName  string
	cacheName   string
	clusterName string
	count       *int32
	cursor      *string
	where       *[]string
	select_     *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) Count(count int32) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) Where(where []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) Execute() (MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters Get a list of Home Cache Cluster objects.
 * Get a list of Home Cache Cluster objects.

A Home Cache Cluster is a Cache Cluster that is the "definitive" Cache Cluster for a given topic in the context of the Global Caching feature.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
homeClusterName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters(ctx _context.Context, msgVpnName string, cacheName string, clusterName string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest{
		ApiService:  a,
		ctx:         ctx,
		msgVpnName:  msgVpnName,
		cacheName:   cacheName,
		clusterName: clusterName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) (MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRequest struct {
	ctx          _context.Context
	ApiService   *MsgVpnApiService
	msgVpnName   string
	cacheName    string
	clusterName  string
	instanceName string
	select_      *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRequest) Execute() (MsgVpnDistributedCacheClusterInstanceResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterInstanceExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterInstance Get a Cache Instance object.
 * Get a Cache Instance object.

A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
counter.msgCount||x
counter.msgPeakCount||x
counter.requestQueueDepthCount||x
counter.requestQueueDepthPeakCount||x
counter.topicCount||x
counter.topicPeakCount||x
instanceName|x|
msgVpnName|x|
rate.averageDataRxBytePeakRate||x
rate.averageDataRxByteRate||x
rate.averageDataRxMsgPeakRate||x
rate.averageDataRxMsgRate||x
rate.averageDataTxMsgPeakRate||x
rate.averageDataTxMsgRate||x
rate.averageRequestRxPeakRate||x
rate.averageRequestRxRate||x
rate.dataRxBytePeakRate||x
rate.dataRxByteRate||x
rate.dataRxMsgPeakRate||x
rate.dataRxMsgRate||x
rate.dataTxMsgPeakRate||x
rate.dataTxMsgRate||x
rate.requestRxPeakRate||x
rate.requestRxRate||x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param instanceName The name of the Cache Instance.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstance(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRequest{
		ApiService:   a,
		ctx:          ctx,
		msgVpnName:   msgVpnName,
		cacheName:    cacheName,
		clusterName:  clusterName,
		instanceName: instanceName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterInstanceResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstanceExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRequest) (MsgVpnDistributedCacheClusterInstanceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterInstance")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", _neturl.PathEscape(parameterToString(r.instanceName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest struct {
	ctx             _context.Context
	ApiService      *MsgVpnApiService
	msgVpnName      string
	cacheName       string
	clusterName     string
	instanceName    string
	homeClusterName string
	select_         *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest) Execute() (MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster Get a Remote Home Cache Cluster object.
 * Get a Remote Home Cache Cluster object.

A Remote Home Cache Cluster is a Home Cache Cluster that the Cache Instance is communicating with in the context of the Global Caching feature.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
homeClusterName|x|
instanceName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param instanceName The name of the Cache Instance.
 * @param homeClusterName The name of the remote Home Cache Cluster.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string, homeClusterName string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest{
		ApiService:      a,
		ctx:             ctx,
		msgVpnName:      msgVpnName,
		cacheName:       cacheName,
		clusterName:     clusterName,
		instanceName:    instanceName,
		homeClusterName: homeClusterName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest) (MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters/{homeClusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", _neturl.PathEscape(parameterToString(r.instanceName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"homeClusterName"+"}", _neturl.PathEscape(parameterToString(r.homeClusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest struct {
	ctx          _context.Context
	ApiService   *MsgVpnApiService
	msgVpnName   string
	cacheName    string
	clusterName  string
	instanceName string
	count        *int32
	cursor       *string
	where        *[]string
	select_      *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) Count(count int32) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) Where(where []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) Execute() (MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters Get a list of Remote Home Cache Cluster objects.
 * Get a list of Remote Home Cache Cluster objects.

A Remote Home Cache Cluster is a Home Cache Cluster that the Cache Instance is communicating with in the context of the Global Caching feature.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
homeClusterName|x|
instanceName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param instanceName The name of the Cache Instance.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest{
		ApiService:   a,
		ctx:          ctx,
		msgVpnName:   msgVpnName,
		cacheName:    cacheName,
		clusterName:  clusterName,
		instanceName: instanceName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) (MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteGlobalCachingHomeClusters"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", _neturl.PathEscape(parameterToString(r.instanceName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest struct {
	ctx          _context.Context
	ApiService   *MsgVpnApiService
	msgVpnName   string
	cacheName    string
	clusterName  string
	instanceName string
	topic        string
	select_      *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest) Execute() (MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteTopicExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterInstanceRemoteTopic Get a Remote Topic object.
 * Get a Remote Topic object.

A Remote Topic is a topic for which the Cache Instance has cached messages.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
instanceName|x|
msgVpnName|x|
topic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param instanceName The name of the Cache Instance.
 * @param topic The value of the remote Topic.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteTopic(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string, topic string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest{
		ApiService:   a,
		ctx:          ctx,
		msgVpnName:   msgVpnName,
		cacheName:    cacheName,
		clusterName:  clusterName,
		instanceName: instanceName,
		topic:        topic,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteTopicExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest) (MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteTopic")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics/{topic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", _neturl.PathEscape(parameterToString(r.instanceName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topic"+"}", _neturl.PathEscape(parameterToString(r.topic, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest struct {
	ctx          _context.Context
	ApiService   *MsgVpnApiService
	msgVpnName   string
	cacheName    string
	clusterName  string
	instanceName string
	count        *int32
	cursor       *string
	where        *[]string
	select_      *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) Execute() (MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteTopicsExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterInstanceRemoteTopics Get a list of Remote Topic objects.
 * Get a list of Remote Topic objects.

A Remote Topic is a topic for which the Cache Instance has cached messages.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
instanceName|x|
msgVpnName|x|
topic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param instanceName The name of the Cache Instance.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteTopics(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest{
		ApiService:   a,
		ctx:          ctx,
		msgVpnName:   msgVpnName,
		cacheName:    cacheName,
		clusterName:  clusterName,
		instanceName: instanceName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteTopicsExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) (MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteTopics")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}/remoteTopics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceName"+"}", _neturl.PathEscape(parameterToString(r.instanceName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest struct {
	ctx         _context.Context
	ApiService  *MsgVpnApiService
	msgVpnName  string
	cacheName   string
	clusterName string
	count       *int32
	cursor      *string
	where       *[]string
	select_     *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) Execute() (MsgVpnDistributedCacheClusterInstancesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterInstancesExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterInstances Get a list of Cache Instance objects.
 * Get a list of Cache Instance objects.

A Cache Instance is a single Cache process that belongs to a single Cache Cluster. A Cache Instance object provisioned on the broker is used to disseminate configuration information to the Cache process. Cache Instances listen for and cache live data messages that match the topic subscriptions configured for their parent Cache Cluster.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
counter.msgCount||x
counter.msgPeakCount||x
counter.requestQueueDepthCount||x
counter.requestQueueDepthPeakCount||x
counter.topicCount||x
counter.topicPeakCount||x
instanceName|x|
msgVpnName|x|
rate.averageDataRxBytePeakRate||x
rate.averageDataRxByteRate||x
rate.averageDataRxMsgPeakRate||x
rate.averageDataRxMsgRate||x
rate.averageDataTxMsgPeakRate||x
rate.averageDataTxMsgRate||x
rate.averageRequestRxPeakRate||x
rate.averageRequestRxRate||x
rate.dataRxBytePeakRate||x
rate.dataRxByteRate||x
rate.dataRxMsgPeakRate||x
rate.dataRxMsgRate||x
rate.dataTxMsgPeakRate||x
rate.dataTxMsgRate||x
rate.requestRxPeakRate||x
rate.requestRxRate||x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstances(ctx _context.Context, msgVpnName string, cacheName string, clusterName string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest{
		ApiService:  a,
		ctx:         ctx,
		msgVpnName:  msgVpnName,
		cacheName:   cacheName,
		clusterName: clusterName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterInstancesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterInstancesExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) (MsgVpnDistributedCacheClusterInstancesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstancesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterInstances")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicRequest struct {
	ctx         _context.Context
	ApiService  *MsgVpnApiService
	msgVpnName  string
	cacheName   string
	clusterName string
	topic       string
	select_     *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicRequest) Execute() (MsgVpnDistributedCacheClusterTopicResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterTopicExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterTopic Get a Topic object.
 * Get a Topic object.

The Cache Instances that belong to the containing Cache Cluster will cache any messages published to topics that match a Topic Subscription.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
msgVpnName|x|
topic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @param topic The value of the Topic in the form a/b/c.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterTopic(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, topic string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicRequest{
		ApiService:  a,
		ctx:         ctx,
		msgVpnName:  msgVpnName,
		cacheName:   cacheName,
		clusterName: clusterName,
		topic:       topic,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterTopicResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterTopicExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicRequest) (MsgVpnDistributedCacheClusterTopicResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterTopicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterTopic")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topic"+"}", _neturl.PathEscape(parameterToString(r.topic, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest struct {
	ctx         _context.Context
	ApiService  *MsgVpnApiService
	msgVpnName  string
	cacheName   string
	clusterName string
	count       *int32
	cursor      *string
	where       *[]string
	select_     *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) Execute() (MsgVpnDistributedCacheClusterTopicsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClusterTopicsExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusterTopics Get a list of Topic objects.
 * Get a list of Topic objects.

The Cache Instances that belong to the containing Cache Cluster will cache any messages published to topics that match a Topic Subscription.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
msgVpnName|x|
topic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @param clusterName The name of the Cache Cluster.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterTopics(ctx _context.Context, msgVpnName string, cacheName string, clusterName string) MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest{
		ApiService:  a,
		ctx:         ctx,
		msgVpnName:  msgVpnName,
		cacheName:   cacheName,
		clusterName: clusterName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClusterTopicsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusterTopicsExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) (MsgVpnDistributedCacheClusterTopicsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterTopicsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusterTopics")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(parameterToString(r.clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	cacheName  string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest) Count(count int32) MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest) Where(where []string) MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest) Execute() (MsgVpnDistributedCacheClustersResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCacheClustersExecute(r)
}

/*
 * GetMsgVpnDistributedCacheClusters Get a list of Cache Cluster objects.
 * Get a list of Cache Cluster objects.

A Cache Cluster is a collection of one or more Cache Instances that subscribe to exactly the same topics. Cache Instances are grouped together in a Cache Cluster for the purpose of fault tolerance and load balancing. As published messages are received, the message broker message bus sends these live data messages to the Cache Instances in the Cache Cluster. This enables client cache requests to be served by any of Cache Instances in the Cache Cluster.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
clusterName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the Distributed Cache.
 * @return MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClusters(ctx _context.Context, msgVpnName string, cacheName string) MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		cacheName:  cacheName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCacheClustersResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCacheClustersExecute(r MsgVpnApiApiGetMsgVpnDistributedCacheClustersRequest) (MsgVpnDistributedCacheClustersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClustersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCacheClusters")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDistributedCachesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnDistributedCachesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnDistributedCachesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCachesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnDistributedCachesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCachesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnDistributedCachesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnDistributedCachesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDistributedCachesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDistributedCachesRequest) Execute() (MsgVpnDistributedCachesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDistributedCachesExecute(r)
}

/*
 * GetMsgVpnDistributedCaches Get a list of Distributed Cache objects.
 * Get a list of Distributed Cache objects.

A Distributed Cache is a collection of one or more Cache Clusters that belong to the same Message VPN. Each Cache Cluster in a Distributed Cache is configured to subscribe to a different set of topics. This effectively divides up the configured topic space, to provide scaling to very large topic spaces or very high cached message throughput.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnDistributedCachesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDistributedCaches(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnDistributedCachesRequest {
	return MsgVpnApiApiGetMsgVpnDistributedCachesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCachesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDistributedCachesExecute(r MsgVpnApiApiGetMsgVpnDistributedCachesRequest) (MsgVpnDistributedCachesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCachesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDistributedCaches")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/distributedCaches"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDmrBridgeRequest struct {
	ctx            _context.Context
	ApiService     *MsgVpnApiService
	msgVpnName     string
	remoteNodeName string
	select_        *[]string
}

func (r MsgVpnApiApiGetMsgVpnDmrBridgeRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDmrBridgeRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDmrBridgeRequest) Execute() (MsgVpnDmrBridgeResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDmrBridgeExecute(r)
}

/*
 * GetMsgVpnDmrBridge Get a DMR Bridge object.
 * Get a DMR Bridge object.

A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
remoteNodeName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param remoteNodeName The name of the node at the remote end of the DMR Bridge.
 * @return MsgVpnApiApiGetMsgVpnDmrBridgeRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDmrBridge(ctx _context.Context, msgVpnName string, remoteNodeName string) MsgVpnApiApiGetMsgVpnDmrBridgeRequest {
	return MsgVpnApiApiGetMsgVpnDmrBridgeRequest{
		ApiService:     a,
		ctx:            ctx,
		msgVpnName:     msgVpnName,
		remoteNodeName: remoteNodeName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDmrBridgeResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDmrBridgeExecute(r MsgVpnApiApiGetMsgVpnDmrBridgeRequest) (MsgVpnDmrBridgeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDmrBridgeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDmrBridge")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteNodeName"+"}", _neturl.PathEscape(parameterToString(r.remoteNodeName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnDmrBridgesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnDmrBridgesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnDmrBridgesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnDmrBridgesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnDmrBridgesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnDmrBridgesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnDmrBridgesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnDmrBridgesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnDmrBridgesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnDmrBridgesRequest) Execute() (MsgVpnDmrBridgesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnDmrBridgesExecute(r)
}

/*
 * GetMsgVpnDmrBridges Get a list of DMR Bridge objects.
 * Get a list of DMR Bridge objects.

A DMR Bridge is required to establish a data channel over a corresponding external link to the remote node for a given Message VPN. Each DMR Bridge identifies which external link the Message VPN should use, and what the name of the equivalent Message VPN at the remote node is.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
remoteNodeName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnDmrBridgesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnDmrBridges(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnDmrBridgesRequest {
	return MsgVpnApiApiGetMsgVpnDmrBridgesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDmrBridgesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnDmrBridgesExecute(r MsgVpnApiApiGetMsgVpnDmrBridgesRequest) (MsgVpnDmrBridgesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDmrBridgesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnDmrBridges")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/dmrBridges"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest) Execute() (MsgVpnJndiConnectionFactoriesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnJndiConnectionFactoriesExecute(r)
}

/*
 * GetMsgVpnJndiConnectionFactories Get a list of JNDI Connection Factory objects.
 * Get a list of JNDI Connection Factory objects.

The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.


Attribute|Identifying|Deprecated
:---|:---:|:---:
connectionFactoryName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnJndiConnectionFactories(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest {
	return MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnJndiConnectionFactoriesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnJndiConnectionFactoriesExecute(r MsgVpnApiApiGetMsgVpnJndiConnectionFactoriesRequest) (MsgVpnJndiConnectionFactoriesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiConnectionFactoriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnJndiConnectionFactories")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/jndiConnectionFactories"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnJndiConnectionFactoryRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	connectionFactoryName string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnJndiConnectionFactoryRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnJndiConnectionFactoryRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnJndiConnectionFactoryRequest) Execute() (MsgVpnJndiConnectionFactoryResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnJndiConnectionFactoryExecute(r)
}

/*
 * GetMsgVpnJndiConnectionFactory Get a JNDI Connection Factory object.
 * Get a JNDI Connection Factory object.

The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.


Attribute|Identifying|Deprecated
:---|:---:|:---:
connectionFactoryName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param connectionFactoryName The name of the JMS Connection Factory.
 * @return MsgVpnApiApiGetMsgVpnJndiConnectionFactoryRequest
*/
func (a *MsgVpnApiService) GetMsgVpnJndiConnectionFactory(ctx _context.Context, msgVpnName string, connectionFactoryName string) MsgVpnApiApiGetMsgVpnJndiConnectionFactoryRequest {
	return MsgVpnApiApiGetMsgVpnJndiConnectionFactoryRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		connectionFactoryName: connectionFactoryName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnJndiConnectionFactoryResponse
 */
func (a *MsgVpnApiService) GetMsgVpnJndiConnectionFactoryExecute(r MsgVpnApiApiGetMsgVpnJndiConnectionFactoryRequest) (MsgVpnJndiConnectionFactoryResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiConnectionFactoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnJndiConnectionFactory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connectionFactoryName"+"}", _neturl.PathEscape(parameterToString(r.connectionFactoryName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnJndiQueueRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	queueName  string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnJndiQueueRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnJndiQueueRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnJndiQueueRequest) Execute() (MsgVpnJndiQueueResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnJndiQueueExecute(r)
}

/*
 * GetMsgVpnJndiQueue Get a JNDI Queue object.
 * Get a JNDI Queue object.

The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
queueName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueName The JNDI name of the JMS Queue.
 * @return MsgVpnApiApiGetMsgVpnJndiQueueRequest
*/
func (a *MsgVpnApiService) GetMsgVpnJndiQueue(ctx _context.Context, msgVpnName string, queueName string) MsgVpnApiApiGetMsgVpnJndiQueueRequest {
	return MsgVpnApiApiGetMsgVpnJndiQueueRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		queueName:  queueName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnJndiQueueResponse
 */
func (a *MsgVpnApiService) GetMsgVpnJndiQueueExecute(r MsgVpnApiApiGetMsgVpnJndiQueueRequest) (MsgVpnJndiQueueResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiQueueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnJndiQueue")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/jndiQueues/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", _neturl.PathEscape(parameterToString(r.queueName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnJndiQueuesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnJndiQueuesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnJndiQueuesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnJndiQueuesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnJndiQueuesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnJndiQueuesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnJndiQueuesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnJndiQueuesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnJndiQueuesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnJndiQueuesRequest) Execute() (MsgVpnJndiQueuesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnJndiQueuesExecute(r)
}

/*
 * GetMsgVpnJndiQueues Get a list of JNDI Queue objects.
 * Get a list of JNDI Queue objects.

The message broker provides an internal JNDI store for provisioned Queue objects that clients can access through JNDI lookups.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
queueName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnJndiQueuesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnJndiQueues(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnJndiQueuesRequest {
	return MsgVpnApiApiGetMsgVpnJndiQueuesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnJndiQueuesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnJndiQueuesExecute(r MsgVpnApiApiGetMsgVpnJndiQueuesRequest) (MsgVpnJndiQueuesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiQueuesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnJndiQueues")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/jndiQueues"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnJndiTopicRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	topicName  string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnJndiTopicRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnJndiTopicRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnJndiTopicRequest) Execute() (MsgVpnJndiTopicResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnJndiTopicExecute(r)
}

/*
 * GetMsgVpnJndiTopic Get a JNDI Topic object.
 * Get a JNDI Topic object.

The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
topicName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param topicName The JNDI name of the JMS Topic.
 * @return MsgVpnApiApiGetMsgVpnJndiTopicRequest
*/
func (a *MsgVpnApiService) GetMsgVpnJndiTopic(ctx _context.Context, msgVpnName string, topicName string) MsgVpnApiApiGetMsgVpnJndiTopicRequest {
	return MsgVpnApiApiGetMsgVpnJndiTopicRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		topicName:  topicName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnJndiTopicResponse
 */
func (a *MsgVpnApiService) GetMsgVpnJndiTopicExecute(r MsgVpnApiApiGetMsgVpnJndiTopicRequest) (MsgVpnJndiTopicResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiTopicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnJndiTopic")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/jndiTopics/{topicName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicName"+"}", _neturl.PathEscape(parameterToString(r.topicName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnJndiTopicsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnJndiTopicsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnJndiTopicsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnJndiTopicsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnJndiTopicsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnJndiTopicsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnJndiTopicsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnJndiTopicsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnJndiTopicsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnJndiTopicsRequest) Execute() (MsgVpnJndiTopicsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnJndiTopicsExecute(r)
}

/*
 * GetMsgVpnJndiTopics Get a list of JNDI Topic objects.
 * Get a list of JNDI Topic objects.

The message broker provides an internal JNDI store for provisioned Topic objects that clients can access through JNDI lookups.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
topicName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnJndiTopicsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnJndiTopics(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnJndiTopicsRequest {
	return MsgVpnApiApiGetMsgVpnJndiTopicsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnJndiTopicsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnJndiTopicsExecute(r MsgVpnApiApiGetMsgVpnJndiTopicsRequest) (MsgVpnJndiTopicsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiTopicsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnJndiTopics")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/jndiTopics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnMqttRetainCacheRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	cacheName  string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnMqttRetainCacheRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnMqttRetainCacheRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnMqttRetainCacheRequest) Execute() (MsgVpnMqttRetainCacheResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnMqttRetainCacheExecute(r)
}

/*
 * GetMsgVpnMqttRetainCache Get an MQTT Retain Cache object.
 * Get an MQTT Retain Cache object.

Using MQTT retained messages allows publishing MQTT clients to indicate that a message must be stored for later delivery to subscribing clients when those subscribing clients add subscriptions matching the retained message's topic. An MQTT Retain Cache processes all retained messages for a Message VPN.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param cacheName The name of the MQTT Retain Cache.
 * @return MsgVpnApiApiGetMsgVpnMqttRetainCacheRequest
*/
func (a *MsgVpnApiService) GetMsgVpnMqttRetainCache(ctx _context.Context, msgVpnName string, cacheName string) MsgVpnApiApiGetMsgVpnMqttRetainCacheRequest {
	return MsgVpnApiApiGetMsgVpnMqttRetainCacheRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		cacheName:  cacheName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnMqttRetainCacheResponse
 */
func (a *MsgVpnApiService) GetMsgVpnMqttRetainCacheExecute(r MsgVpnApiApiGetMsgVpnMqttRetainCacheRequest) (MsgVpnMqttRetainCacheResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttRetainCacheResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnMqttRetainCache")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cacheName"+"}", _neturl.PathEscape(parameterToString(r.cacheName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest) Execute() (MsgVpnMqttRetainCachesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnMqttRetainCachesExecute(r)
}

/*
 * GetMsgVpnMqttRetainCaches Get a list of MQTT Retain Cache objects.
 * Get a list of MQTT Retain Cache objects.

Using MQTT retained messages allows publishing MQTT clients to indicate that a message must be stored for later delivery to subscribing clients when those subscribing clients add subscriptions matching the retained message's topic. An MQTT Retain Cache processes all retained messages for a Message VPN.


Attribute|Identifying|Deprecated
:---|:---:|:---:
cacheName|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnMqttRetainCaches(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest {
	return MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnMqttRetainCachesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnMqttRetainCachesExecute(r MsgVpnApiApiGetMsgVpnMqttRetainCachesRequest) (MsgVpnMqttRetainCachesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttRetainCachesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnMqttRetainCaches")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/mqttRetainCaches"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnMqttSessionRequest struct {
	ctx                      _context.Context
	ApiService               *MsgVpnApiService
	msgVpnName               string
	mqttSessionClientId      string
	mqttSessionVirtualRouter string
	select_                  *[]string
}

func (r MsgVpnApiApiGetMsgVpnMqttSessionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnMqttSessionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnMqttSessionRequest) Execute() (MsgVpnMqttSessionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnMqttSessionExecute(r)
}

/*
 * GetMsgVpnMqttSession Get an MQTT Session object.
 * Get an MQTT Session object.

An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).


Attribute|Identifying|Deprecated
:---|:---:|:---:
counter.mqttConnackErrorTxCount||x
counter.mqttConnackTxCount||x
counter.mqttConnectRxCount||x
counter.mqttDisconnectRxCount||x
counter.mqttPubcompTxCount||x
counter.mqttPublishQos0RxCount||x
counter.mqttPublishQos0TxCount||x
counter.mqttPublishQos1RxCount||x
counter.mqttPublishQos1TxCount||x
counter.mqttPublishQos2RxCount||x
counter.mqttPubrecTxCount||x
counter.mqttPubrelRxCount||x
mqttSessionClientId|x|
mqttSessionVirtualRouter|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param mqttSessionClientId The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
 * @param mqttSessionVirtualRouter The virtual router of the MQTT Session.
 * @return MsgVpnApiApiGetMsgVpnMqttSessionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnMqttSession(ctx _context.Context, msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string) MsgVpnApiApiGetMsgVpnMqttSessionRequest {
	return MsgVpnApiApiGetMsgVpnMqttSessionRequest{
		ApiService:               a,
		ctx:                      ctx,
		msgVpnName:               msgVpnName,
		mqttSessionClientId:      mqttSessionClientId,
		mqttSessionVirtualRouter: mqttSessionVirtualRouter,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnMqttSessionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnMqttSessionExecute(r MsgVpnApiApiGetMsgVpnMqttSessionRequest) (MsgVpnMqttSessionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnMqttSession")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", _neturl.PathEscape(parameterToString(r.mqttSessionClientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.mqttSessionVirtualRouter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionRequest struct {
	ctx                      _context.Context
	ApiService               *MsgVpnApiService
	msgVpnName               string
	mqttSessionClientId      string
	mqttSessionVirtualRouter string
	subscriptionTopic        string
	select_                  *[]string
}

func (r MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionRequest) Execute() (MsgVpnMqttSessionSubscriptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnMqttSessionSubscriptionExecute(r)
}

/*
 * GetMsgVpnMqttSessionSubscription Get a Subscription object.
 * Get a Subscription object.

An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.


Attribute|Identifying|Deprecated
:---|:---:|:---:
mqttSessionClientId|x|
mqttSessionVirtualRouter|x|
msgVpnName|x|
subscriptionTopic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param mqttSessionClientId The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
 * @param mqttSessionVirtualRouter The virtual router of the MQTT Session.
 * @param subscriptionTopic The MQTT subscription topic.
 * @return MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnMqttSessionSubscription(ctx _context.Context, msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, subscriptionTopic string) MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionRequest {
	return MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionRequest{
		ApiService:               a,
		ctx:                      ctx,
		msgVpnName:               msgVpnName,
		mqttSessionClientId:      mqttSessionClientId,
		mqttSessionVirtualRouter: mqttSessionVirtualRouter,
		subscriptionTopic:        subscriptionTopic,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnMqttSessionSubscriptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnMqttSessionSubscriptionExecute(r MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionRequest) (MsgVpnMqttSessionSubscriptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttSessionSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnMqttSessionSubscription")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", _neturl.PathEscape(parameterToString(r.mqttSessionClientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.mqttSessionVirtualRouter, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionTopic"+"}", _neturl.PathEscape(parameterToString(r.subscriptionTopic, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest struct {
	ctx                      _context.Context
	ApiService               *MsgVpnApiService
	msgVpnName               string
	mqttSessionClientId      string
	mqttSessionVirtualRouter string
	count                    *int32
	cursor                   *string
	where                    *[]string
	select_                  *[]string
}

func (r MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest) Execute() (MsgVpnMqttSessionSubscriptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnMqttSessionSubscriptionsExecute(r)
}

/*
 * GetMsgVpnMqttSessionSubscriptions Get a list of Subscription objects.
 * Get a list of Subscription objects.

An MQTT session contains a client's QoS 0 and QoS 1 subscription sets. On creation, a subscription defaults to QoS 0.


Attribute|Identifying|Deprecated
:---|:---:|:---:
mqttSessionClientId|x|
mqttSessionVirtualRouter|x|
msgVpnName|x|
subscriptionTopic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param mqttSessionClientId The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
 * @param mqttSessionVirtualRouter The virtual router of the MQTT Session.
 * @return MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnMqttSessionSubscriptions(ctx _context.Context, msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string) MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest {
	return MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest{
		ApiService:               a,
		ctx:                      ctx,
		msgVpnName:               msgVpnName,
		mqttSessionClientId:      mqttSessionClientId,
		mqttSessionVirtualRouter: mqttSessionVirtualRouter,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnMqttSessionSubscriptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnMqttSessionSubscriptionsExecute(r MsgVpnApiApiGetMsgVpnMqttSessionSubscriptionsRequest) (MsgVpnMqttSessionSubscriptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttSessionSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnMqttSessionSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionClientId"+"}", _neturl.PathEscape(parameterToString(r.mqttSessionClientId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mqttSessionVirtualRouter"+"}", _neturl.PathEscape(parameterToString(r.mqttSessionVirtualRouter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnMqttSessionsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnMqttSessionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnMqttSessionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnMqttSessionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnMqttSessionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnMqttSessionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnMqttSessionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnMqttSessionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnMqttSessionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnMqttSessionsRequest) Execute() (MsgVpnMqttSessionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnMqttSessionsExecute(r)
}

/*
 * GetMsgVpnMqttSessions Get a list of MQTT Session objects.
 * Get a list of MQTT Session objects.

An MQTT Session object is a virtual representation of an MQTT client connection. An MQTT session holds the state of an MQTT client (that is, it is used to contain a client's QoS 0 and QoS 1 subscription sets and any undelivered QoS 1 messages).


Attribute|Identifying|Deprecated
:---|:---:|:---:
counter.mqttConnackErrorTxCount||x
counter.mqttConnackTxCount||x
counter.mqttConnectRxCount||x
counter.mqttDisconnectRxCount||x
counter.mqttPubcompTxCount||x
counter.mqttPublishQos0RxCount||x
counter.mqttPublishQos0TxCount||x
counter.mqttPublishQos1RxCount||x
counter.mqttPublishQos1TxCount||x
counter.mqttPublishQos2RxCount||x
counter.mqttPubrecTxCount||x
counter.mqttPubrelRxCount||x
mqttSessionClientId|x|
mqttSessionVirtualRouter|x|
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnMqttSessionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnMqttSessions(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnMqttSessionsRequest {
	return MsgVpnApiApiGetMsgVpnMqttSessionsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnMqttSessionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnMqttSessionsExecute(r MsgVpnApiApiGetMsgVpnMqttSessionsRequest) (MsgVpnMqttSessionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttSessionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnMqttSessions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/mqttSessions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueueRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	queueName  string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueueRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueueRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueueRequest) Execute() (MsgVpnQueueResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueueExecute(r)
}

/*
 * GetMsgVpnQueue Get a Queue object.
 * Get a Queue object.

A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
queueName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueName The name of the Queue.
 * @return MsgVpnApiApiGetMsgVpnQueueRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueue(ctx _context.Context, msgVpnName string, queueName string) MsgVpnApiApiGetMsgVpnQueueRequest {
	return MsgVpnApiApiGetMsgVpnQueueRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		queueName:  queueName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueueResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueueExecute(r MsgVpnApiApiGetMsgVpnQueueRequest) (MsgVpnQueueResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueue")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queues/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", _neturl.PathEscape(parameterToString(r.queueName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueueMsgRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	queueName  string
	msgId      string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueueMsgRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueueMsgRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueueMsgRequest) Execute() (MsgVpnQueueMsgResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueueMsgExecute(r)
}

/*
 * GetMsgVpnQueueMsg Get a Queue Message object.
 * Get a Queue Message object.

A Queue Message is a packet of information sent from producers to consumers using the Queue.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgId|x|
msgVpnName|x|
queueName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueName The name of the Queue.
 * @param msgId The identifier (ID) of the Message.
 * @return MsgVpnApiApiGetMsgVpnQueueMsgRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueueMsg(ctx _context.Context, msgVpnName string, queueName string, msgId string) MsgVpnApiApiGetMsgVpnQueueMsgRequest {
	return MsgVpnApiApiGetMsgVpnQueueMsgRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		queueName:  queueName,
		msgId:      msgId,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueueMsgResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueueMsgExecute(r MsgVpnApiApiGetMsgVpnQueueMsgRequest) (MsgVpnQueueMsgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueMsgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueueMsg")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/msgs/{msgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", _neturl.PathEscape(parameterToString(r.queueName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"msgId"+"}", _neturl.PathEscape(parameterToString(r.msgId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueueMsgsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	queueName  string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueueMsgsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnQueueMsgsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueMsgsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnQueueMsgsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueMsgsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnQueueMsgsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueMsgsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueueMsgsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueueMsgsRequest) Execute() (MsgVpnQueueMsgsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueueMsgsExecute(r)
}

/*
 * GetMsgVpnQueueMsgs Get a list of Queue Message objects.
 * Get a list of Queue Message objects.

A Queue Message is a packet of information sent from producers to consumers using the Queue.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgId|x|
msgVpnName|x|
queueName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueName The name of the Queue.
 * @return MsgVpnApiApiGetMsgVpnQueueMsgsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueueMsgs(ctx _context.Context, msgVpnName string, queueName string) MsgVpnApiApiGetMsgVpnQueueMsgsRequest {
	return MsgVpnApiApiGetMsgVpnQueueMsgsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		queueName:  queueName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueueMsgsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueueMsgsExecute(r MsgVpnApiApiGetMsgVpnQueueMsgsRequest) (MsgVpnQueueMsgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueMsgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueueMsgs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/msgs"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", _neturl.PathEscape(parameterToString(r.queueName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	queueName  string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest) Execute() (MsgVpnQueuePrioritiesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueuePrioritiesExecute(r)
}

/*
 * GetMsgVpnQueuePriorities Get a list of Queue Priority objects.
 * Get a list of Queue Priority objects.

Queues can optionally support priority message delivery; all messages of a higher priority are delivered before any messages of a lower priority. A Priority object contains information about the number and size of the messages with a particular priority in the Queue.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
priority|x|
queueName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueName The name of the Queue.
 * @return MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueuePriorities(ctx _context.Context, msgVpnName string, queueName string) MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest {
	return MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		queueName:  queueName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueuePrioritiesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueuePrioritiesExecute(r MsgVpnApiApiGetMsgVpnQueuePrioritiesRequest) (MsgVpnQueuePrioritiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueuePrioritiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueuePriorities")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/priorities"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", _neturl.PathEscape(parameterToString(r.queueName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueuePriorityRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	queueName  string
	priority   string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueuePriorityRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueuePriorityRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueuePriorityRequest) Execute() (MsgVpnQueuePriorityResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueuePriorityExecute(r)
}

/*
 * GetMsgVpnQueuePriority Get a Queue Priority object.
 * Get a Queue Priority object.

Queues can optionally support priority message delivery; all messages of a higher priority are delivered before any messages of a lower priority. A Priority object contains information about the number and size of the messages with a particular priority in the Queue.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
priority|x|
queueName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueName The name of the Queue.
 * @param priority The level of the Priority, from 9 (highest) to 0 (lowest).
 * @return MsgVpnApiApiGetMsgVpnQueuePriorityRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueuePriority(ctx _context.Context, msgVpnName string, queueName string, priority string) MsgVpnApiApiGetMsgVpnQueuePriorityRequest {
	return MsgVpnApiApiGetMsgVpnQueuePriorityRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		queueName:  queueName,
		priority:   priority,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueuePriorityResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueuePriorityExecute(r MsgVpnApiApiGetMsgVpnQueuePriorityRequest) (MsgVpnQueuePriorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueuePriorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueuePriority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/priorities/{priority}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", _neturl.PathEscape(parameterToString(r.queueName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"priority"+"}", _neturl.PathEscape(parameterToString(r.priority, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueueSubscriptionRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	queueName         string
	subscriptionTopic string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueueSubscriptionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueueSubscriptionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueueSubscriptionRequest) Execute() (MsgVpnQueueSubscriptionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueueSubscriptionExecute(r)
}

/*
 * GetMsgVpnQueueSubscription Get a Queue Subscription object.
 * Get a Queue Subscription object.

One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
queueName|x|
subscriptionTopic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueName The name of the Queue.
 * @param subscriptionTopic The topic of the Subscription.
 * @return MsgVpnApiApiGetMsgVpnQueueSubscriptionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueueSubscription(ctx _context.Context, msgVpnName string, queueName string, subscriptionTopic string) MsgVpnApiApiGetMsgVpnQueueSubscriptionRequest {
	return MsgVpnApiApiGetMsgVpnQueueSubscriptionRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		queueName:         queueName,
		subscriptionTopic: subscriptionTopic,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueueSubscriptionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueueSubscriptionExecute(r MsgVpnApiApiGetMsgVpnQueueSubscriptionRequest) (MsgVpnQueueSubscriptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueueSubscription")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", _neturl.PathEscape(parameterToString(r.queueName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionTopic"+"}", _neturl.PathEscape(parameterToString(r.subscriptionTopic, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	queueName  string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest) Execute() (MsgVpnQueueSubscriptionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueueSubscriptionsExecute(r)
}

/*
 * GetMsgVpnQueueSubscriptions Get a list of Queue Subscription objects.
 * Get a list of Queue Subscription objects.

One or more Queue Subscriptions can be added to a durable queue so that Guaranteed messages published to matching topics are also delivered to and spooled by the queue.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
queueName|x|
subscriptionTopic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueName The name of the Queue.
 * @return MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueueSubscriptions(ctx _context.Context, msgVpnName string, queueName string) MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest {
	return MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		queueName:  queueName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueueSubscriptionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueueSubscriptionsExecute(r MsgVpnApiApiGetMsgVpnQueueSubscriptionsRequest) (MsgVpnQueueSubscriptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueueSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", _neturl.PathEscape(parameterToString(r.queueName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueueTemplateRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	queueTemplateName string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueueTemplateRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueueTemplateRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueueTemplateRequest) Execute() (MsgVpnQueueTemplateResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueueTemplateExecute(r)
}

/*
 * GetMsgVpnQueueTemplate Get a Queue Template object.
 * Get a Queue Template object.

A Queue Template provides a mechanism for specifying the initial state for client created queues.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
queueTemplateName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.14.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueTemplateName The name of the Queue Template.
 * @return MsgVpnApiApiGetMsgVpnQueueTemplateRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueueTemplate(ctx _context.Context, msgVpnName string, queueTemplateName string) MsgVpnApiApiGetMsgVpnQueueTemplateRequest {
	return MsgVpnApiApiGetMsgVpnQueueTemplateRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		queueTemplateName: queueTemplateName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueueTemplateResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueueTemplateExecute(r MsgVpnApiApiGetMsgVpnQueueTemplateRequest) (MsgVpnQueueTemplateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueTemplateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueueTemplate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueTemplateName"+"}", _neturl.PathEscape(parameterToString(r.queueTemplateName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueueTemplatesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueueTemplatesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnQueueTemplatesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueTemplatesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnQueueTemplatesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueTemplatesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnQueueTemplatesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueTemplatesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueueTemplatesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueueTemplatesRequest) Execute() (MsgVpnQueueTemplatesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueueTemplatesExecute(r)
}

/*
 * GetMsgVpnQueueTemplates Get a list of Queue Template objects.
 * Get a list of Queue Template objects.

A Queue Template provides a mechanism for specifying the initial state for client created queues.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
queueTemplateName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.14.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnQueueTemplatesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueueTemplates(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnQueueTemplatesRequest {
	return MsgVpnApiApiGetMsgVpnQueueTemplatesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueueTemplatesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueueTemplatesExecute(r MsgVpnApiApiGetMsgVpnQueueTemplatesRequest) (MsgVpnQueueTemplatesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueTemplatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueueTemplates")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queueTemplates"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueueTxFlowRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	queueName  string
	flowId     string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueueTxFlowRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueueTxFlowRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueueTxFlowRequest) Execute() (MsgVpnQueueTxFlowResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueueTxFlowExecute(r)
}

/*
 * GetMsgVpnQueueTxFlow Get a Queue Transmit Flow object.
 * Get a Queue Transmit Flow object.

Queue Transmit Flows are used by clients to consume Guaranteed messages from a Queue.


Attribute|Identifying|Deprecated
:---|:---:|:---:
flowId|x|
msgVpnName|x|
queueName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueName The name of the Queue.
 * @param flowId The identifier (ID) of the Flow.
 * @return MsgVpnApiApiGetMsgVpnQueueTxFlowRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueueTxFlow(ctx _context.Context, msgVpnName string, queueName string, flowId string) MsgVpnApiApiGetMsgVpnQueueTxFlowRequest {
	return MsgVpnApiApiGetMsgVpnQueueTxFlowRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		queueName:  queueName,
		flowId:     flowId,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueueTxFlowResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueueTxFlowExecute(r MsgVpnApiApiGetMsgVpnQueueTxFlowRequest) (MsgVpnQueueTxFlowResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueTxFlowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueueTxFlow")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/txFlows/{flowId}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", _neturl.PathEscape(parameterToString(r.queueName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"flowId"+"}", _neturl.PathEscape(parameterToString(r.flowId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	queueName  string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest) Execute() (MsgVpnQueueTxFlowsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueueTxFlowsExecute(r)
}

/*
 * GetMsgVpnQueueTxFlows Get a list of Queue Transmit Flow objects.
 * Get a list of Queue Transmit Flow objects.

Queue Transmit Flows are used by clients to consume Guaranteed messages from a Queue.


Attribute|Identifying|Deprecated
:---|:---:|:---:
flowId|x|
msgVpnName|x|
queueName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param queueName The name of the Queue.
 * @return MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueueTxFlows(ctx _context.Context, msgVpnName string, queueName string) MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest {
	return MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		queueName:  queueName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueueTxFlowsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueueTxFlowsExecute(r MsgVpnApiApiGetMsgVpnQueueTxFlowsRequest) (MsgVpnQueueTxFlowsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueTxFlowsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueueTxFlows")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queues/{queueName}/txFlows"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", _neturl.PathEscape(parameterToString(r.queueName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnQueuesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnQueuesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnQueuesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueuesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnQueuesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueuesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnQueuesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnQueuesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnQueuesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnQueuesRequest) Execute() (MsgVpnQueuesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnQueuesExecute(r)
}

/*
 * GetMsgVpnQueues Get a list of Queue objects.
 * Get a list of Queue objects.

A Queue acts as both a destination that clients can publish messages to, and as an endpoint that clients can bind consumers to and consume messages from.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
queueName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnQueuesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnQueues(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnQueuesRequest {
	return MsgVpnApiApiGetMsgVpnQueuesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueuesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnQueuesExecute(r MsgVpnApiApiGetMsgVpnQueuesRequest) (MsgVpnQueuesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueuesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnQueues")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/queues"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnReplayLogRequest struct {
	ctx           _context.Context
	ApiService    *MsgVpnApiService
	msgVpnName    string
	replayLogName string
	select_       *[]string
}

func (r MsgVpnApiApiGetMsgVpnReplayLogRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnReplayLogRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnReplayLogRequest) Execute() (MsgVpnReplayLogResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnReplayLogExecute(r)
}

/*
 * GetMsgVpnReplayLog Get a Replay Log object.
 * Get a Replay Log object.

When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
replayLogName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param replayLogName The name of the Replay Log.
 * @return MsgVpnApiApiGetMsgVpnReplayLogRequest
*/
func (a *MsgVpnApiService) GetMsgVpnReplayLog(ctx _context.Context, msgVpnName string, replayLogName string) MsgVpnApiApiGetMsgVpnReplayLogRequest {
	return MsgVpnApiApiGetMsgVpnReplayLogRequest{
		ApiService:    a,
		ctx:           ctx,
		msgVpnName:    msgVpnName,
		replayLogName: replayLogName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnReplayLogResponse
 */
func (a *MsgVpnApiService) GetMsgVpnReplayLogExecute(r MsgVpnApiApiGetMsgVpnReplayLogRequest) (MsgVpnReplayLogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplayLogResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnReplayLog")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/replayLogs/{replayLogName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"replayLogName"+"}", _neturl.PathEscape(parameterToString(r.replayLogName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnReplayLogMsgRequest struct {
	ctx           _context.Context
	ApiService    *MsgVpnApiService
	msgVpnName    string
	replayLogName string
	msgId         string
	select_       *[]string
}

func (r MsgVpnApiApiGetMsgVpnReplayLogMsgRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnReplayLogMsgRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnReplayLogMsgRequest) Execute() (MsgVpnReplayLogMsgResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnReplayLogMsgExecute(r)
}

/*
 * GetMsgVpnReplayLogMsg Get a Message object.
 * Get a Message object.

A Message is a packet of information sent from producers to consumers. Messages are the central units of information that clients exchange using the message broker and which are cached in the Replay Log.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgId|x|
msgVpnName|x|
replayLogName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param replayLogName The name of the Replay Log.
 * @param msgId The identifier (ID) of the message.
 * @return MsgVpnApiApiGetMsgVpnReplayLogMsgRequest
*/
func (a *MsgVpnApiService) GetMsgVpnReplayLogMsg(ctx _context.Context, msgVpnName string, replayLogName string, msgId string) MsgVpnApiApiGetMsgVpnReplayLogMsgRequest {
	return MsgVpnApiApiGetMsgVpnReplayLogMsgRequest{
		ApiService:    a,
		ctx:           ctx,
		msgVpnName:    msgVpnName,
		replayLogName: replayLogName,
		msgId:         msgId,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnReplayLogMsgResponse
 */
func (a *MsgVpnApiService) GetMsgVpnReplayLogMsgExecute(r MsgVpnApiApiGetMsgVpnReplayLogMsgRequest) (MsgVpnReplayLogMsgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplayLogMsgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnReplayLogMsg")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs/{msgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"replayLogName"+"}", _neturl.PathEscape(parameterToString(r.replayLogName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"msgId"+"}", _neturl.PathEscape(parameterToString(r.msgId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest struct {
	ctx           _context.Context
	ApiService    *MsgVpnApiService
	msgVpnName    string
	replayLogName string
	count         *int32
	cursor        *string
	where         *[]string
	select_       *[]string
}

func (r MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest) Execute() (MsgVpnReplayLogMsgsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnReplayLogMsgsExecute(r)
}

/*
 * GetMsgVpnReplayLogMsgs Get a list of Message objects.
 * Get a list of Message objects.

A Message is a packet of information sent from producers to consumers. Messages are the central units of information that clients exchange using the message broker and which are cached in the Replay Log.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgId|x|
msgVpnName|x|
replayLogName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param replayLogName The name of the Replay Log.
 * @return MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnReplayLogMsgs(ctx _context.Context, msgVpnName string, replayLogName string) MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest {
	return MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest{
		ApiService:    a,
		ctx:           ctx,
		msgVpnName:    msgVpnName,
		replayLogName: replayLogName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnReplayLogMsgsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnReplayLogMsgsExecute(r MsgVpnApiApiGetMsgVpnReplayLogMsgsRequest) (MsgVpnReplayLogMsgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplayLogMsgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnReplayLogMsgs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/replayLogs/{replayLogName}/msgs"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"replayLogName"+"}", _neturl.PathEscape(parameterToString(r.replayLogName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnReplayLogsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnReplayLogsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnReplayLogsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnReplayLogsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnReplayLogsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnReplayLogsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnReplayLogsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnReplayLogsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnReplayLogsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnReplayLogsRequest) Execute() (MsgVpnReplayLogsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnReplayLogsExecute(r)
}

/*
 * GetMsgVpnReplayLogs Get a list of Replay Log objects.
 * Get a list of Replay Log objects.

When the Message Replay feature is enabled, message brokers store persistent messages in a Replay Log. These messages are kept until the log is full, after which the oldest messages are removed to free up space for new messages.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
replayLogName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnReplayLogsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnReplayLogs(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnReplayLogsRequest {
	return MsgVpnApiApiGetMsgVpnReplayLogsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnReplayLogsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnReplayLogsExecute(r MsgVpnApiApiGetMsgVpnReplayLogsRequest) (MsgVpnReplayLogsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplayLogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnReplayLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/replayLogs"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnReplicatedTopicRequest struct {
	ctx             _context.Context
	ApiService      *MsgVpnApiService
	msgVpnName      string
	replicatedTopic string
	select_         *[]string
}

func (r MsgVpnApiApiGetMsgVpnReplicatedTopicRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnReplicatedTopicRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnReplicatedTopicRequest) Execute() (MsgVpnReplicatedTopicResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnReplicatedTopicExecute(r)
}

/*
 * GetMsgVpnReplicatedTopic Get a Replicated Topic object.
 * Get a Replicated Topic object.

To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
replicatedTopic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param replicatedTopic The topic for applying replication. Published messages matching this topic will be replicated to the standby site.
 * @return MsgVpnApiApiGetMsgVpnReplicatedTopicRequest
*/
func (a *MsgVpnApiService) GetMsgVpnReplicatedTopic(ctx _context.Context, msgVpnName string, replicatedTopic string) MsgVpnApiApiGetMsgVpnReplicatedTopicRequest {
	return MsgVpnApiApiGetMsgVpnReplicatedTopicRequest{
		ApiService:      a,
		ctx:             ctx,
		msgVpnName:      msgVpnName,
		replicatedTopic: replicatedTopic,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnReplicatedTopicResponse
 */
func (a *MsgVpnApiService) GetMsgVpnReplicatedTopicExecute(r MsgVpnApiApiGetMsgVpnReplicatedTopicRequest) (MsgVpnReplicatedTopicResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplicatedTopicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnReplicatedTopic")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"replicatedTopic"+"}", _neturl.PathEscape(parameterToString(r.replicatedTopic, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest) Execute() (MsgVpnReplicatedTopicsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnReplicatedTopicsExecute(r)
}

/*
 * GetMsgVpnReplicatedTopics Get a list of Replicated Topic objects.
 * Get a list of Replicated Topic objects.

To indicate which messages should be replicated between the active and standby site, a Replicated Topic subscription must be configured on a Message VPN. If a published message matches both a replicated topic and an endpoint on the active site, then the message is replicated to the standby site.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
replicatedTopic|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnReplicatedTopics(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest {
	return MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnReplicatedTopicsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnReplicatedTopicsExecute(r MsgVpnApiApiGetMsgVpnReplicatedTopicsRequest) (MsgVpnReplicatedTopicsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplicatedTopicsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnReplicatedTopics")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/replicatedTopics"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnRestDeliveryPointRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	restDeliveryPointName string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRequest) Execute() (MsgVpnRestDeliveryPointResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnRestDeliveryPointExecute(r)
}

/*
 * GetMsgVpnRestDeliveryPoint Get a REST Delivery Point object.
 * Get a REST Delivery Point object.

A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
restDeliveryPointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @return MsgVpnApiApiGetMsgVpnRestDeliveryPointRequest
*/
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPoint(ctx _context.Context, msgVpnName string, restDeliveryPointName string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRequest {
	return MsgVpnApiApiGetMsgVpnRestDeliveryPointRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		restDeliveryPointName: restDeliveryPointName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointResponse
 */
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointExecute(r MsgVpnApiApiGetMsgVpnRestDeliveryPointRequest) (MsgVpnRestDeliveryPointResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnRestDeliveryPoint")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", _neturl.PathEscape(parameterToString(r.restDeliveryPointName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	restDeliveryPointName string
	queueBindingName      string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest) Execute() (MsgVpnRestDeliveryPointQueueBindingResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnRestDeliveryPointQueueBindingExecute(r)
}

/*
 * GetMsgVpnRestDeliveryPointQueueBinding Get a Queue Binding object.
 * Get a Queue Binding object.

A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
queueBindingName|x|
restDeliveryPointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param queueBindingName The name of a queue in the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest
*/
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointQueueBinding(ctx _context.Context, msgVpnName string, restDeliveryPointName string, queueBindingName string) MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest {
	return MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		restDeliveryPointName: restDeliveryPointName,
		queueBindingName:      queueBindingName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointQueueBindingResponse
 */
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointQueueBindingExecute(r MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest) (MsgVpnRestDeliveryPointQueueBindingResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointQueueBindingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnRestDeliveryPointQueueBinding")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", _neturl.PathEscape(parameterToString(r.restDeliveryPointName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueBindingName"+"}", _neturl.PathEscape(parameterToString(r.queueBindingName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	restDeliveryPointName string
	count                 *int32
	cursor                *string
	where                 *[]string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) Execute() (MsgVpnRestDeliveryPointQueueBindingsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnRestDeliveryPointQueueBindingsExecute(r)
}

/*
 * GetMsgVpnRestDeliveryPointQueueBindings Get a list of Queue Binding objects.
 * Get a list of Queue Binding objects.

A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
queueBindingName|x|
restDeliveryPointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @return MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointQueueBindings(ctx _context.Context, msgVpnName string, restDeliveryPointName string) MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest {
	return MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		restDeliveryPointName: restDeliveryPointName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointQueueBindingsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointQueueBindingsExecute(r MsgVpnApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) (MsgVpnRestDeliveryPointQueueBindingsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointQueueBindingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnRestDeliveryPointQueueBindings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", _neturl.PathEscape(parameterToString(r.restDeliveryPointName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	restDeliveryPointName string
	restConsumerName      string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest) Execute() (MsgVpnRestDeliveryPointRestConsumerResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnRestDeliveryPointRestConsumerExecute(r)
}

/*
 * GetMsgVpnRestDeliveryPointRestConsumer Get a REST Consumer object.
 * Get a REST Consumer object.

REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
counter.httpRequestConnectionCloseTxMsgCount||x
counter.httpRequestOutstandingTxMsgCount||x
counter.httpRequestTimedOutTxMsgCount||x
counter.httpRequestTxByteCount||x
counter.httpRequestTxMsgCount||x
counter.httpResponseErrorRxMsgCount||x
counter.httpResponseRxByteCount||x
counter.httpResponseRxMsgCount||x
counter.httpResponseSuccessRxMsgCount||x
msgVpnName|x|
restConsumerName|x|
restDeliveryPointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest
*/
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumer(ctx _context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest {
	return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		restDeliveryPointName: restDeliveryPointName,
		restConsumerName:      restConsumerName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointRestConsumerResponse
 */
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumerExecute(r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest) (MsgVpnRestDeliveryPointRestConsumerResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnRestDeliveryPointRestConsumer")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", _neturl.PathEscape(parameterToString(r.restDeliveryPointName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", _neturl.PathEscape(parameterToString(r.restConsumerName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	restDeliveryPointName string
	restConsumerName      string
	oauthJwtClaimName     string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest) Execute() (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimExecute(r)
}

/*
 * GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim Get a Claim object.
 * Get a Claim object.

A Claim is added to the JWT sent to the OAuth token request endpoint.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
oauthJwtClaimName|x|
restConsumerName|x|
restDeliveryPointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.21.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param oauthJwtClaimName The name of the additional claim. Cannot be \"exp\", \"iat\", or \"jti\".
 * @return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest
*/
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx _context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, oauthJwtClaimName string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest {
	return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		restDeliveryPointName: restDeliveryPointName,
		restConsumerName:      restConsumerName,
		oauthJwtClaimName:     oauthJwtClaimName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
 */
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimExecute(r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest) (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", _neturl.PathEscape(parameterToString(r.restDeliveryPointName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", _neturl.PathEscape(parameterToString(r.restConsumerName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"oauthJwtClaimName"+"}", _neturl.PathEscape(parameterToString(r.oauthJwtClaimName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	restDeliveryPointName string
	restConsumerName      string
	count                 *int32
	cursor                *string
	where                 *[]string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) Execute() (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsExecute(r)
}

/*
 * GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims Get a list of Claim objects.
 * Get a list of Claim objects.

A Claim is added to the JWT sent to the OAuth token request endpoint.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
oauthJwtClaimName|x|
restConsumerName|x|
restDeliveryPointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.21.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims(ctx _context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest {
	return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		restDeliveryPointName: restDeliveryPointName,
		restConsumerName:      restConsumerName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsExecute(r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", _neturl.PathEscape(parameterToString(r.restDeliveryPointName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", _neturl.PathEscape(parameterToString(r.restConsumerName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	restDeliveryPointName string
	restConsumerName      string
	tlsTrustedCommonName  string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest) Execute() (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameExecute(r)
}

/*
 * GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName Get a Trusted Common Name object.
 * Get a Trusted Common Name object.

The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|x
restConsumerName|x|x
restDeliveryPointName|x|x
tlsTrustedCommonName|x|x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been deprecated since 2.17. Common Name validation has been replaced by Server Certificate Name validation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param tlsTrustedCommonName The expected trusted common name of the remote certificate.
 * @return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest
*/
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx _context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, tlsTrustedCommonName string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest {
	return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		restDeliveryPointName: restDeliveryPointName,
		restConsumerName:      restConsumerName,
		tlsTrustedCommonName:  tlsTrustedCommonName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
 */
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameExecute(r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest) (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", _neturl.PathEscape(parameterToString(r.restDeliveryPointName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", _neturl.PathEscape(parameterToString(r.restConsumerName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tlsTrustedCommonName"+"}", _neturl.PathEscape(parameterToString(r.tlsTrustedCommonName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	restDeliveryPointName string
	restConsumerName      string
	where                 *[]string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest) Execute() (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesExecute(r)
}

/*
 * GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames Get a list of Trusted Common Name objects.
 * Get a list of Trusted Common Name objects.

The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer's server certificate.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|x
restConsumerName|x|x
restDeliveryPointName|x|x
tlsTrustedCommonName|x|x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been deprecated since 2.17. Common Name validation has been replaced by Server Certificate Name validation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames(ctx _context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest {
	return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		restDeliveryPointName: restDeliveryPointName,
		restConsumerName:      restConsumerName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesExecute(r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest) (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", _neturl.PathEscape(parameterToString(r.restDeliveryPointName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", _neturl.PathEscape(parameterToString(r.restConsumerName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest struct {
	ctx                   _context.Context
	ApiService            *MsgVpnApiService
	msgVpnName            string
	restDeliveryPointName string
	count                 *int32
	cursor                *string
	where                 *[]string
	select_               *[]string
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) Count(count int32) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) Where(where []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) Execute() (MsgVpnRestDeliveryPointRestConsumersResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnRestDeliveryPointRestConsumersExecute(r)
}

/*
 * GetMsgVpnRestDeliveryPointRestConsumers Get a list of REST Consumer objects.
 * Get a list of REST Consumer objects.

REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
counter.httpRequestConnectionCloseTxMsgCount||x
counter.httpRequestOutstandingTxMsgCount||x
counter.httpRequestTimedOutTxMsgCount||x
counter.httpRequestTxByteCount||x
counter.httpRequestTxMsgCount||x
counter.httpResponseErrorRxMsgCount||x
counter.httpResponseRxByteCount||x
counter.httpResponseRxMsgCount||x
counter.httpResponseSuccessRxMsgCount||x
msgVpnName|x|
restConsumerName|x|
restDeliveryPointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest
*/
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumers(ctx _context.Context, msgVpnName string, restDeliveryPointName string) MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest {
	return MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest{
		ApiService:            a,
		ctx:                   ctx,
		msgVpnName:            msgVpnName,
		restDeliveryPointName: restDeliveryPointName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointRestConsumersResponse
 */
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointRestConsumersExecute(r MsgVpnApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) (MsgVpnRestDeliveryPointRestConsumersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnRestDeliveryPointRestConsumers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", _neturl.PathEscape(parameterToString(r.restDeliveryPointName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest) Execute() (MsgVpnRestDeliveryPointsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnRestDeliveryPointsExecute(r)
}

/*
 * GetMsgVpnRestDeliveryPoints Get a list of REST Delivery Point objects.
 * Get a list of REST Delivery Point objects.

A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
restDeliveryPointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPoints(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest {
	return MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnRestDeliveryPointsExecute(r MsgVpnApiApiGetMsgVpnRestDeliveryPointsRequest) (MsgVpnRestDeliveryPointsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnRestDeliveryPoints")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTopicEndpointRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	topicEndpointName string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTopicEndpointRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointRequest) Execute() (MsgVpnTopicEndpointResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTopicEndpointExecute(r)
}

/*
 * GetMsgVpnTopicEndpoint Get a Topic Endpoint object.
 * Get a Topic Endpoint object.

A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
topicEndpointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param topicEndpointName The name of the Topic Endpoint.
 * @return MsgVpnApiApiGetMsgVpnTopicEndpointRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTopicEndpoint(ctx _context.Context, msgVpnName string, topicEndpointName string) MsgVpnApiApiGetMsgVpnTopicEndpointRequest {
	return MsgVpnApiApiGetMsgVpnTopicEndpointRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		topicEndpointName: topicEndpointName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointExecute(r MsgVpnApiApiGetMsgVpnTopicEndpointRequest) (MsgVpnTopicEndpointResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTopicEndpoint")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", _neturl.PathEscape(parameterToString(r.topicEndpointName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTopicEndpointMsgRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	topicEndpointName string
	msgId             string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointMsgRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTopicEndpointMsgRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointMsgRequest) Execute() (MsgVpnTopicEndpointMsgResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTopicEndpointMsgExecute(r)
}

/*
 * GetMsgVpnTopicEndpointMsg Get a Topic Endpoint Message object.
 * Get a Topic Endpoint Message object.

A Topic Endpoint Message is a packet of information sent from producers to consumers using the Topic Endpoint.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgId|x|
msgVpnName|x|
topicEndpointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param topicEndpointName The name of the Topic Endpoint.
 * @param msgId The identifier (ID) of the Message.
 * @return MsgVpnApiApiGetMsgVpnTopicEndpointMsgRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointMsg(ctx _context.Context, msgVpnName string, topicEndpointName string, msgId string) MsgVpnApiApiGetMsgVpnTopicEndpointMsgRequest {
	return MsgVpnApiApiGetMsgVpnTopicEndpointMsgRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		topicEndpointName: topicEndpointName,
		msgId:             msgId,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointMsgResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointMsgExecute(r MsgVpnApiApiGetMsgVpnTopicEndpointMsgRequest) (MsgVpnTopicEndpointMsgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointMsgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTopicEndpointMsg")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs/{msgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", _neturl.PathEscape(parameterToString(r.topicEndpointName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"msgId"+"}", _neturl.PathEscape(parameterToString(r.msgId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	topicEndpointName string
	count             *int32
	cursor            *string
	where             *[]string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest) Execute() (MsgVpnTopicEndpointMsgsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTopicEndpointMsgsExecute(r)
}

/*
 * GetMsgVpnTopicEndpointMsgs Get a list of Topic Endpoint Message objects.
 * Get a list of Topic Endpoint Message objects.

A Topic Endpoint Message is a packet of information sent from producers to consumers using the Topic Endpoint.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgId|x|
msgVpnName|x|
topicEndpointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param topicEndpointName The name of the Topic Endpoint.
 * @return MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointMsgs(ctx _context.Context, msgVpnName string, topicEndpointName string) MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest {
	return MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		topicEndpointName: topicEndpointName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointMsgsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointMsgsExecute(r MsgVpnApiApiGetMsgVpnTopicEndpointMsgsRequest) (MsgVpnTopicEndpointMsgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointMsgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTopicEndpointMsgs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/msgs"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", _neturl.PathEscape(parameterToString(r.topicEndpointName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	topicEndpointName string
	count             *int32
	cursor            *string
	where             *[]string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest) Execute() (MsgVpnTopicEndpointPrioritiesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTopicEndpointPrioritiesExecute(r)
}

/*
 * GetMsgVpnTopicEndpointPriorities Get a list of Topic Endpoint Priority objects.
 * Get a list of Topic Endpoint Priority objects.

Topic Endpoints can optionally support priority message delivery; all messages of a higher priority are delivered before any messages of a lower priority. A Priority object contains information about the number and size of the messages with a particular priority in the Topic Endpoint.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
priority|x|
topicEndpointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param topicEndpointName The name of the Topic Endpoint.
 * @return MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointPriorities(ctx _context.Context, msgVpnName string, topicEndpointName string) MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest {
	return MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		topicEndpointName: topicEndpointName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointPrioritiesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointPrioritiesExecute(r MsgVpnApiApiGetMsgVpnTopicEndpointPrioritiesRequest) (MsgVpnTopicEndpointPrioritiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointPrioritiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTopicEndpointPriorities")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", _neturl.PathEscape(parameterToString(r.topicEndpointName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTopicEndpointPriorityRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	topicEndpointName string
	priority          string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointPriorityRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTopicEndpointPriorityRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointPriorityRequest) Execute() (MsgVpnTopicEndpointPriorityResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTopicEndpointPriorityExecute(r)
}

/*
 * GetMsgVpnTopicEndpointPriority Get a Topic Endpoint Priority object.
 * Get a Topic Endpoint Priority object.

Topic Endpoints can optionally support priority message delivery; all messages of a higher priority are delivered before any messages of a lower priority. A Priority object contains information about the number and size of the messages with a particular priority in the Topic Endpoint.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
priority|x|
topicEndpointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param topicEndpointName The name of the Topic Endpoint.
 * @param priority The level of the Priority, from 9 (highest) to 0 (lowest).
 * @return MsgVpnApiApiGetMsgVpnTopicEndpointPriorityRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointPriority(ctx _context.Context, msgVpnName string, topicEndpointName string, priority string) MsgVpnApiApiGetMsgVpnTopicEndpointPriorityRequest {
	return MsgVpnApiApiGetMsgVpnTopicEndpointPriorityRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		topicEndpointName: topicEndpointName,
		priority:          priority,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointPriorityResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointPriorityExecute(r MsgVpnApiApiGetMsgVpnTopicEndpointPriorityRequest) (MsgVpnTopicEndpointPriorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointPriorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTopicEndpointPriority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/priorities/{priority}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", _neturl.PathEscape(parameterToString(r.topicEndpointName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"priority"+"}", _neturl.PathEscape(parameterToString(r.priority, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTopicEndpointTemplateRequest struct {
	ctx                       _context.Context
	ApiService                *MsgVpnApiService
	msgVpnName                string
	topicEndpointTemplateName string
	select_                   *[]string
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointTemplateRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTopicEndpointTemplateRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointTemplateRequest) Execute() (MsgVpnTopicEndpointTemplateResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTopicEndpointTemplateExecute(r)
}

/*
 * GetMsgVpnTopicEndpointTemplate Get a Topic Endpoint Template object.
 * Get a Topic Endpoint Template object.

A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
topicEndpointTemplateName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.14.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param topicEndpointTemplateName The name of the Topic Endpoint Template.
 * @return MsgVpnApiApiGetMsgVpnTopicEndpointTemplateRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointTemplate(ctx _context.Context, msgVpnName string, topicEndpointTemplateName string) MsgVpnApiApiGetMsgVpnTopicEndpointTemplateRequest {
	return MsgVpnApiApiGetMsgVpnTopicEndpointTemplateRequest{
		ApiService:                a,
		ctx:                       ctx,
		msgVpnName:                msgVpnName,
		topicEndpointTemplateName: topicEndpointTemplateName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointTemplateResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointTemplateExecute(r MsgVpnApiApiGetMsgVpnTopicEndpointTemplateRequest) (MsgVpnTopicEndpointTemplateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointTemplateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTopicEndpointTemplate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/topicEndpointTemplates/{topicEndpointTemplateName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointTemplateName"+"}", _neturl.PathEscape(parameterToString(r.topicEndpointTemplateName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest) Count(count int32) MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest) Where(where []string) MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest) Execute() (MsgVpnTopicEndpointTemplatesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTopicEndpointTemplatesExecute(r)
}

/*
 * GetMsgVpnTopicEndpointTemplates Get a list of Topic Endpoint Template objects.
 * Get a list of Topic Endpoint Template objects.

A Topic Endpoint Template provides a mechanism for specifying the initial state for client created topic endpoints.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
topicEndpointTemplateName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.14.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointTemplates(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest {
	return MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointTemplatesResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointTemplatesExecute(r MsgVpnApiApiGetMsgVpnTopicEndpointTemplatesRequest) (MsgVpnTopicEndpointTemplatesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointTemplatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTopicEndpointTemplates")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/topicEndpointTemplates"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	topicEndpointName string
	flowId            string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowRequest) Execute() (MsgVpnTopicEndpointTxFlowResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTopicEndpointTxFlowExecute(r)
}

/*
 * GetMsgVpnTopicEndpointTxFlow Get a Topic Endpoint Transmit Flow object.
 * Get a Topic Endpoint Transmit Flow object.

Topic Endpoint Transmit Flows are used by clients to consume Guaranteed messages from a Topic Endpoint.


Attribute|Identifying|Deprecated
:---|:---:|:---:
flowId|x|
msgVpnName|x|
topicEndpointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param topicEndpointName The name of the Topic Endpoint.
 * @param flowId The identifier (ID) of the Flow.
 * @return MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointTxFlow(ctx _context.Context, msgVpnName string, topicEndpointName string, flowId string) MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowRequest {
	return MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		topicEndpointName: topicEndpointName,
		flowId:            flowId,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointTxFlowResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointTxFlowExecute(r MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowRequest) (MsgVpnTopicEndpointTxFlowResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointTxFlowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTopicEndpointTxFlow")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows/{flowId}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", _neturl.PathEscape(parameterToString(r.topicEndpointName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"flowId"+"}", _neturl.PathEscape(parameterToString(r.flowId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest struct {
	ctx               _context.Context
	ApiService        *MsgVpnApiService
	msgVpnName        string
	topicEndpointName string
	count             *int32
	cursor            *string
	where             *[]string
	select_           *[]string
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest) Execute() (MsgVpnTopicEndpointTxFlowsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTopicEndpointTxFlowsExecute(r)
}

/*
 * GetMsgVpnTopicEndpointTxFlows Get a list of Topic Endpoint Transmit Flow objects.
 * Get a list of Topic Endpoint Transmit Flow objects.

Topic Endpoint Transmit Flows are used by clients to consume Guaranteed messages from a Topic Endpoint.


Attribute|Identifying|Deprecated
:---|:---:|:---:
flowId|x|
msgVpnName|x|
topicEndpointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param topicEndpointName The name of the Topic Endpoint.
 * @return MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointTxFlows(ctx _context.Context, msgVpnName string, topicEndpointName string) MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest {
	return MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest{
		ApiService:        a,
		ctx:               ctx,
		msgVpnName:        msgVpnName,
		topicEndpointName: topicEndpointName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointTxFlowsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointTxFlowsExecute(r MsgVpnApiApiGetMsgVpnTopicEndpointTxFlowsRequest) (MsgVpnTopicEndpointTxFlowsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointTxFlowsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTopicEndpointTxFlows")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}/txFlows"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topicEndpointName"+"}", _neturl.PathEscape(parameterToString(r.topicEndpointName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTopicEndpointsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnTopicEndpointsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnTopicEndpointsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnTopicEndpointsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnTopicEndpointsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTopicEndpointsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTopicEndpointsRequest) Execute() (MsgVpnTopicEndpointsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTopicEndpointsExecute(r)
}

/*
 * GetMsgVpnTopicEndpoints Get a list of Topic Endpoint objects.
 * Get a list of Topic Endpoint objects.

A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
topicEndpointName|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnTopicEndpointsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTopicEndpoints(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnTopicEndpointsRequest {
	return MsgVpnApiApiGetMsgVpnTopicEndpointsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTopicEndpointsExecute(r MsgVpnApiApiGetMsgVpnTopicEndpointsRequest) (MsgVpnTopicEndpointsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTopicEndpoints")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/topicEndpoints"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTransactionRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	xid        string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnTransactionRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTransactionRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTransactionRequest) Execute() (MsgVpnTransactionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTransactionExecute(r)
}

/*
 * GetMsgVpnTransaction Get a Replicated Local Transaction or XA Transaction object.
 * Get a Replicated Local Transaction or XA Transaction object.

Transactions can be used to group a set of Guaranteed messages to be published or consumed or both as an atomic unit of work.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
xid|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param xid The identifier (ID) of the Transaction.
 * @return MsgVpnApiApiGetMsgVpnTransactionRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTransaction(ctx _context.Context, msgVpnName string, xid string) MsgVpnApiApiGetMsgVpnTransactionRequest {
	return MsgVpnApiApiGetMsgVpnTransactionRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		xid:        xid,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTransactionResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTransactionExecute(r MsgVpnApiApiGetMsgVpnTransactionRequest) (MsgVpnTransactionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTransaction")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/transactions/{xid}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"xid"+"}", _neturl.PathEscape(parameterToString(r.xid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTransactionConsumerMsgRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	xid        string
	msgId      string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnTransactionConsumerMsgRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTransactionConsumerMsgRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTransactionConsumerMsgRequest) Execute() (MsgVpnTransactionConsumerMsgResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTransactionConsumerMsgExecute(r)
}

/*
 * GetMsgVpnTransactionConsumerMsg Get a Transaction Consumer Message object.
 * Get a Transaction Consumer Message object.

A Transaction Consumer Message is a message that will be consumed as part of this Transaction once the Transaction is committed.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgId|x|
msgVpnName|x|
xid|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param xid The identifier (ID) of the Transaction.
 * @param msgId The identifier (ID) of the Message.
 * @return MsgVpnApiApiGetMsgVpnTransactionConsumerMsgRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTransactionConsumerMsg(ctx _context.Context, msgVpnName string, xid string, msgId string) MsgVpnApiApiGetMsgVpnTransactionConsumerMsgRequest {
	return MsgVpnApiApiGetMsgVpnTransactionConsumerMsgRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		xid:        xid,
		msgId:      msgId,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTransactionConsumerMsgResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTransactionConsumerMsgExecute(r MsgVpnApiApiGetMsgVpnTransactionConsumerMsgRequest) (MsgVpnTransactionConsumerMsgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionConsumerMsgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTransactionConsumerMsg")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs/{msgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"xid"+"}", _neturl.PathEscape(parameterToString(r.xid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"msgId"+"}", _neturl.PathEscape(parameterToString(r.msgId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	xid        string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest) Execute() (MsgVpnTransactionConsumerMsgsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTransactionConsumerMsgsExecute(r)
}

/*
 * GetMsgVpnTransactionConsumerMsgs Get a list of Transaction Consumer Message objects.
 * Get a list of Transaction Consumer Message objects.

A Transaction Consumer Message is a message that will be consumed as part of this Transaction once the Transaction is committed.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgId|x|
msgVpnName|x|
xid|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param xid The identifier (ID) of the Transaction.
 * @return MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTransactionConsumerMsgs(ctx _context.Context, msgVpnName string, xid string) MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest {
	return MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		xid:        xid,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTransactionConsumerMsgsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTransactionConsumerMsgsExecute(r MsgVpnApiApiGetMsgVpnTransactionConsumerMsgsRequest) (MsgVpnTransactionConsumerMsgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionConsumerMsgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTransactionConsumerMsgs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/transactions/{xid}/consumerMsgs"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"xid"+"}", _neturl.PathEscape(parameterToString(r.xid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTransactionPublisherMsgRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	xid        string
	msgId      string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnTransactionPublisherMsgRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTransactionPublisherMsgRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTransactionPublisherMsgRequest) Execute() (MsgVpnTransactionPublisherMsgResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTransactionPublisherMsgExecute(r)
}

/*
 * GetMsgVpnTransactionPublisherMsg Get a Transaction Publisher Message object.
 * Get a Transaction Publisher Message object.

A Transaction Publisher Message is a message that will be published as part of this Transaction once the Transaction is committed.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgId|x|
msgVpnName|x|
xid|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param xid The identifier (ID) of the Transaction.
 * @param msgId The identifier (ID) of the Message.
 * @return MsgVpnApiApiGetMsgVpnTransactionPublisherMsgRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTransactionPublisherMsg(ctx _context.Context, msgVpnName string, xid string, msgId string) MsgVpnApiApiGetMsgVpnTransactionPublisherMsgRequest {
	return MsgVpnApiApiGetMsgVpnTransactionPublisherMsgRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		xid:        xid,
		msgId:      msgId,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTransactionPublisherMsgResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTransactionPublisherMsgExecute(r MsgVpnApiApiGetMsgVpnTransactionPublisherMsgRequest) (MsgVpnTransactionPublisherMsgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionPublisherMsgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTransactionPublisherMsg")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs/{msgId}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"xid"+"}", _neturl.PathEscape(parameterToString(r.xid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"msgId"+"}", _neturl.PathEscape(parameterToString(r.msgId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	xid        string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest) Execute() (MsgVpnTransactionPublisherMsgsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTransactionPublisherMsgsExecute(r)
}

/*
 * GetMsgVpnTransactionPublisherMsgs Get a list of Transaction Publisher Message objects.
 * Get a list of Transaction Publisher Message objects.

A Transaction Publisher Message is a message that will be published as part of this Transaction once the Transaction is committed.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgId|x|
msgVpnName|x|
xid|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param xid The identifier (ID) of the Transaction.
 * @return MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTransactionPublisherMsgs(ctx _context.Context, msgVpnName string, xid string) MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest {
	return MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
		xid:        xid,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTransactionPublisherMsgsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTransactionPublisherMsgsExecute(r MsgVpnApiApiGetMsgVpnTransactionPublisherMsgsRequest) (MsgVpnTransactionPublisherMsgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionPublisherMsgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTransactionPublisherMsgs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/transactions/{xid}/publisherMsgs"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"xid"+"}", _neturl.PathEscape(parameterToString(r.xid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnTransactionsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnTransactionsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnTransactionsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnTransactionsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnTransactionsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnTransactionsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnTransactionsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnTransactionsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnTransactionsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnTransactionsRequest) Execute() (MsgVpnTransactionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnTransactionsExecute(r)
}

/*
 * GetMsgVpnTransactions Get a list of Replicated Local Transaction or XA Transaction objects.
 * Get a list of Replicated Local Transaction or XA Transaction objects.

Transactions can be used to group a set of Guaranteed messages to be published or consumed or both as an atomic unit of work.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|
xid|x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.12.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return MsgVpnApiApiGetMsgVpnTransactionsRequest
*/
func (a *MsgVpnApiService) GetMsgVpnTransactions(ctx _context.Context, msgVpnName string) MsgVpnApiApiGetMsgVpnTransactionsRequest {
	return MsgVpnApiApiGetMsgVpnTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTransactionsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnTransactionsExecute(r MsgVpnApiApiGetMsgVpnTransactionsRequest) (MsgVpnTransactionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpnTransactions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns/{msgVpnName}/transactions"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", _neturl.PathEscape(parameterToString(r.msgVpnName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MsgVpnApiApiGetMsgVpnsRequest struct {
	ctx        _context.Context
	ApiService *MsgVpnApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r MsgVpnApiApiGetMsgVpnsRequest) Count(count int32) MsgVpnApiApiGetMsgVpnsRequest {
	r.count = &count
	return r
}
func (r MsgVpnApiApiGetMsgVpnsRequest) Cursor(cursor string) MsgVpnApiApiGetMsgVpnsRequest {
	r.cursor = &cursor
	return r
}
func (r MsgVpnApiApiGetMsgVpnsRequest) Where(where []string) MsgVpnApiApiGetMsgVpnsRequest {
	r.where = &where
	return r
}
func (r MsgVpnApiApiGetMsgVpnsRequest) Select_(select_ []string) MsgVpnApiApiGetMsgVpnsRequest {
	r.select_ = &select_
	return r
}

func (r MsgVpnApiApiGetMsgVpnsRequest) Execute() (MsgVpnsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetMsgVpnsExecute(r)
}

/*
 * GetMsgVpns Get a list of Message VPN objects.
 * Get a list of Message VPN objects.

Message VPNs (Virtual Private Networks) allow for the segregation of topic space and clients. They also group clients connecting to a network of message brokers, such that messages published within a particular group are only visible to that group's clients.


Attribute|Identifying|Deprecated
:---|:---:|:---:
bridgingTlsServerCertEnforceTrustedCommonNameEnabled||x
counter.controlRxByteCount||x
counter.controlRxMsgCount||x
counter.controlTxByteCount||x
counter.controlTxMsgCount||x
counter.dataRxByteCount||x
counter.dataRxMsgCount||x
counter.dataTxByteCount||x
counter.dataTxMsgCount||x
counter.discardedRxMsgCount||x
counter.discardedTxMsgCount||x
counter.loginRxMsgCount||x
counter.loginTxMsgCount||x
counter.msgSpoolRxMsgCount||x
counter.msgSpoolTxMsgCount||x
counter.tlsRxByteCount||x
counter.tlsTxByteCount||x
msgVpnName|x|
rate.averageRxByteRate||x
rate.averageRxMsgRate||x
rate.averageTxByteRate||x
rate.averageTxMsgRate||x
rate.rxByteRate||x
rate.rxMsgRate||x
rate.tlsAverageRxByteRate||x
rate.tlsAverageTxByteRate||x
rate.tlsRxByteRate||x
rate.tlsTxByteRate||x
rate.txByteRate||x
rate.txMsgRate||x
restTlsServerCertEnforceTrustedCommonNameEnabled||x



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return MsgVpnApiApiGetMsgVpnsRequest
*/
func (a *MsgVpnApiService) GetMsgVpns(ctx _context.Context) MsgVpnApiApiGetMsgVpnsRequest {
	return MsgVpnApiApiGetMsgVpnsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnsResponse
 */
func (a *MsgVpnApiService) GetMsgVpnsExecute(r MsgVpnApiApiGetMsgVpnsRequest) (MsgVpnsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MsgVpnApiService.GetMsgVpns")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/msgVpns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, "csv"))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v SempMetaOnlyResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
