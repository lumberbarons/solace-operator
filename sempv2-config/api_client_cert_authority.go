/*
 * SEMP (Solace Element Management Protocol)
 *
 * SEMP (starting in `v2`, see note 1) is a RESTful API for configuring, monitoring, and administering a Solace PubSub+ broker.  SEMP uses URIs to address manageable **resources** of the Solace PubSub+ broker. Resources are individual **objects**, **collections** of objects, or (exclusively in the action API) **actions**. This document applies to the following API:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Configuration|/SEMP/v2/config|Reading and writing config state|See note 2    The following APIs are also available:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Action|/SEMP/v2/action|Performing actions|See note 2 Monitoring|/SEMP/v2/monitor|Querying operational parameters|See note 2    Resources are always nouns, with individual objects being singular and collections being plural.  Objects within a collection are identified by an `obj-id`, which follows the collection name with the form `collection-name/obj-id`.  Actions within an object are identified by an `action-id`, which follows the object name with the form `obj-id/action-id`.  Some examples:  ``` /SEMP/v2/config/msgVpns                        ; MsgVpn collection /SEMP/v2/config/msgVpns/a                      ; MsgVpn object named \"a\" /SEMP/v2/config/msgVpns/a/queues               ; Queue collection in MsgVpn \"a\" /SEMP/v2/config/msgVpns/a/queues/b             ; Queue object named \"b\" in MsgVpn \"a\" /SEMP/v2/action/msgVpns/a/queues/b/startReplay ; Action that starts a replay on Queue \"b\" in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients             ; Client collection in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients/c           ; Client object named \"c\" in MsgVpn \"a\" ```  ## Collection Resources  Collections are unordered lists of objects (unless described as otherwise), and are described by JSON arrays. Each item in the array represents an object in the same manner as the individual object would normally be represented. In the configuration API, the creation of a new object is done through its collection resource.  ## Object and Action Resources  Objects are composed of attributes, actions, collections, and other objects. They are described by JSON objects as name/value pairs. The collections and actions of an object are not contained directly in the object's JSON content; rather the content includes an attribute containing a URI which points to the collections and actions. These contained resources must be managed through this URI. At a minimum, every object has one or more identifying attributes, and its own `uri` attribute which contains the URI pointing to itself.  Actions are also composed of attributes, and are described by JSON objects as name/value pairs. Unlike objects, however, they are not members of a collection and cannot be retrieved, only performed. Actions only exist in the action API.  Attributes in an object or action may have any combination of the following properties:   Property|Meaning|Comments :---|:---|:--- Identifying|Attribute is involved in unique identification of the object, and appears in its URI| Required|Attribute must be provided in the request| Read-Only|Attribute can only be read, not written.|See note 3 Write-Only|Attribute can only be written, not read, unless the attribute is also opaque|See the documentation for the opaque property Requires-Disable|Attribute can only be changed when object is disabled| Deprecated|Attribute is deprecated, and will disappear in the next SEMP version| Opaque|Attribute can be set or retrieved in opaque form when the `opaquePassword` query parameter is present|See the `opaquePassword` query parameter documentation    In some requests, certain attributes may only be provided in certain combinations with other attributes:   Relationship|Meaning :---|:--- Requires|Attribute may only be changed by a request if a particular attribute or combination of attributes is also provided in the request Conflicts|Attribute may only be provided in a request if a particular attribute or combination of attributes is not also provided in the request    In the monitoring API, any non-identifying attribute may not be returned in a GET.  ## HTTP Methods  The following HTTP methods manipulate resources in accordance with these general principles. Note that some methods are only used in certain APIs:   Method|Resource|Meaning|Request Body|Response Body|Missing Request Attributes :---|:---|:---|:---|:---|:--- POST|Collection|Create object|Initial attribute values|Object attributes and metadata|Set to default PUT|Object|Create or replace object (see note 5)|New attribute values|Object attributes and metadata|Set to default, with certain exceptions (see note 4) PUT|Action|Performs action|Action arguments|Action metadata|N/A PATCH|Object|Update object|New attribute values|Object attributes and metadata|unchanged DELETE|Object|Delete object|Empty|Object metadata|N/A GET|Object|Get object|Empty|Object attributes and metadata|N/A GET|Collection|Get collection|Empty|Object attributes and collection metadata|N/A    ## Common Query Parameters  The following are some common query parameters that are supported by many method/URI combinations. Individual URIs may document additional parameters. Note that multiple query parameters can be used together in a single URI, separated by the ampersand character. For example:  ``` ; Request for the MsgVpns collection using two hypothetical query parameters ; \"q1\" and \"q2\" with values \"val1\" and \"val2\" respectively /SEMP/v2/config/msgVpns?q1=val1&q2=val2 ```  ### select  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. Use this query parameter to limit the size of the returned data for each returned object, return only those fields that are desired, or exclude fields that are not desired.  The value of `select` is a comma-separated list of attribute names. If the list contains attribute names that are not prefaced by `-`, only those attributes are included in the response. If the list contains attribute names that are prefaced by `-`, those attributes are excluded from the response. If the list contains both types, then the difference of the first set of attributes and the second set of attributes is returned. If the list is empty (i.e. `select=`), no attributes are returned.  All attributes that are prefaced by `-` must follow all attributes that are not prefaced by `-`. In addition, each attribute name in the list must match at least one attribute in the object.  Names may include the `*` wildcard (zero or more characters). Nested attribute names are supported using periods (e.g. `parentName.childName`).  Some examples:  ``` ; List of all MsgVpn names /SEMP/v2/config/msgVpns?select=msgVpnName ; List of all MsgVpn and their attributes except for their names /SEMP/v2/config/msgVpns?select=-msgVpnName ; Authentication attributes of MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance?select=authentication* ; All attributes of MsgVpn \"finance\" except for authentication attributes /SEMP/v2/config/msgVpns/finance?select=-authentication* ; Access related attributes of Queue \"orderQ\" of MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance/queues/orderQ?select=owner,permission ```  ### where  Include in the response only objects where certain conditions are true. Use this query parameter to limit which objects are returned to those whose attribute values meet the given conditions.  The value of `where` is a comma-separated list of expressions. All expressions must be true for the object to be included in the response. Each expression takes the form:  ``` expression  = attribute-name OP value OP          = '==' | '!=' | '&lt;' | '&gt;' | '&lt;=' | '&gt;=' ```  `value` may be a number, string, `true`, or `false`, as appropriate for the type of `attribute-name`. Greater-than and less-than comparisons only work for numbers. A `*` in a string `value` is interpreted as a wildcard (zero or more characters). Some examples:  ``` ; Only enabled MsgVpns /SEMP/v2/config/msgVpns?where=enabled==true ; Only MsgVpns using basic non-LDAP authentication /SEMP/v2/config/msgVpns?where=authenticationBasicEnabled==true,authenticationBasicType!=ldap ; Only MsgVpns that allow more than 100 client connections /SEMP/v2/config/msgVpns?where=maxConnectionCount>100 ; Only MsgVpns with msgVpnName starting with \"B\": /SEMP/v2/config/msgVpns?where=msgVpnName==B* ```  ### count  Limit the count of objects in the response. This can be useful to limit the size of the response for large collections. The minimum value for `count` is `1` and the default is `10`. There is also a per-collection maximum value to limit request handling time. For example:  ``` ; Up to 25 MsgVpns /SEMP/v2/config/msgVpns?count=25 ```  ### cursor  The cursor, or position, for the next page of objects. Cursors are opaque data that should not be created or interpreted by SEMP clients, and should only be used as described below.  When a request is made for a collection and there may be additional objects available for retrieval that are not included in the initial response, the response will include a `cursorQuery` field containing a cursor. The value of this field can be specified in the `cursor` query parameter of a subsequent request to retrieve the next page of objects. For convenience, an appropriate URI is constructed automatically by the broker and included in the `nextPageUri` field of the response. This URI can be used directly to retrieve the next page of objects.  ### opaquePassword  Attributes with the opaque property are also write-only and so cannot normally be retrieved in a GET. However, when a password is provided in the `opaquePassword` query parameter, attributes with the opaque property are retrieved in a GET in opaque form, encrypted with this password. The query parameter can also be used on a POST, PATCH, or PUT to set opaque attributes using opaque attribute values retrieved in a GET, so long as:  1. the same password that was used to retrieve the opaque attribute values is provided; and  2. the broker to which the request is being sent has the same major and minor SEMP version as the broker that produced the opaque attribute values.  The password provided in the query parameter must be a minimum of 8 characters and a maximum of 128 characters.  The query parameter can only be used in the configuration API, and only over HTTPS.  ## Authentication  When a client makes its first SEMPv2 request, it must supply a username and password using HTTP Basic authentication.  If authentication is successful, the broker returns a cookie containing a session key. The client can omit the username and password from subsequent requests, because the broker now uses the session cookie for authentication instead. When the session expires or is deleted, the client must provide the username and password again, and the broker creates a new session.  There are a limited number of session slots available on the broker. The broker returns 529 No SEMP Session Available if it is not able to allocate a session. For this reason, all clients that use SEMPv2 should support cookies.  If certain attributes—such as a user's password—are changed, the broker automatically deletes the affected sessions. These attributes are documented below. However, changes in external user configuration data stored on a RADIUS or LDAP server do not trigger the broker to delete the associated session(s), therefore you must do this manually, if required.  A client can retrieve its current session information using the /about/user endpoint, delete its own session using the /about/user/logout endpoint, and manage all sessions using the /sessions endpoint.  ## Help  Visit [our website](https://solace.com) to learn more about Solace.  You can also download the SEMP API specifications by clicking [here](https://solace.com/downloads/).  If you need additional support, please contact us at [support@solace.com](mailto:support@solace.com).  ## Notes  Note|Description :---:|:--- 1|This specification defines SEMP starting in \"v2\", and not the original SEMP \"v1\" interface. Request and response formats between \"v1\" and \"v2\" are entirely incompatible, although both protocols share a common port configuration on the Solace PubSub+ broker. They are differentiated by the initial portion of the URI path, one of either \"/SEMP/\" or \"/SEMP/v2/\" 2|This API is partially implemented. Only a subset of all objects are available. 3|Read-only attributes may appear in POST and PUT/PATCH requests. However, if a read-only attribute is not marked as identifying, it will be ignored during a PUT/PATCH. 4|On a PUT, if the SEMP user is not authorized to modify the attribute, its value is left unchanged rather than set to default. In addition, the values of write-only attributes are not set to their defaults on a PUT, except in the following two cases: there is a mutual requires relationship with another non-write-only attribute, both attributes are absent from the request, and the non-write-only attribute is not currently set to its default value; or the attribute is also opaque and the `opaquePassword` query parameter is provided in the request. 5|On a PUT, if the object does not exist, it is created first.
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

// ClientCertAuthorityApiService ClientCertAuthorityApi service
type ClientCertAuthorityApiService service

type ClientCertAuthorityApiApiCreateClientCertAuthorityRequest struct {
	ctx            _context.Context
	ApiService     *ClientCertAuthorityApiService
	body           *ClientCertAuthority
	opaquePassword *string
	select_        *[]string
}

func (r ClientCertAuthorityApiApiCreateClientCertAuthorityRequest) Body(body ClientCertAuthority) ClientCertAuthorityApiApiCreateClientCertAuthorityRequest {
	r.body = &body
	return r
}
func (r ClientCertAuthorityApiApiCreateClientCertAuthorityRequest) OpaquePassword(opaquePassword string) ClientCertAuthorityApiApiCreateClientCertAuthorityRequest {
	r.opaquePassword = &opaquePassword
	return r
}
func (r ClientCertAuthorityApiApiCreateClientCertAuthorityRequest) Select_(select_ []string) ClientCertAuthorityApiApiCreateClientCertAuthorityRequest {
	r.select_ = &select_
	return r
}

func (r ClientCertAuthorityApiApiCreateClientCertAuthorityRequest) Execute() (ClientCertAuthorityResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateClientCertAuthorityExecute(r)
}

/*
 * CreateClientCertAuthority Create a Client Certificate Authority object.
 * Create a Client Certificate Authority object. Any attribute missing from the request will be set to its default value.

Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.


Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque
:---|:---:|:---:|:---:|:---:|:---:|:---:
certAuthorityName|x|x||||



The following attributes in the request may only be provided in certain combinations with other attributes:


Class|Attribute|Requires|Conflicts
:---|:---|:---|:---
ClientCertAuthority|crlDayList|crlTimeList|
ClientCertAuthority|crlTimeList|crlDayList|



A SEMP client authorized with a minimum access scope/level of "global/admin" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ClientCertAuthorityApiApiCreateClientCertAuthorityRequest
*/
func (a *ClientCertAuthorityApiService) CreateClientCertAuthority(ctx _context.Context) ClientCertAuthorityApiApiCreateClientCertAuthorityRequest {
	return ClientCertAuthorityApiApiCreateClientCertAuthorityRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthorityResponse
 */
func (a *ClientCertAuthorityApiService) CreateClientCertAuthorityExecute(r ClientCertAuthorityApiApiCreateClientCertAuthorityRequest) (ClientCertAuthorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCertAuthorityApiService.CreateClientCertAuthority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.opaquePassword != nil {
		localVarQueryParams.Add("opaquePassword", parameterToString(*r.opaquePassword, ""))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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

type ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest struct {
	ctx               _context.Context
	ApiService        *ClientCertAuthorityApiService
	certAuthorityName string
	body              *ClientCertAuthorityOcspTlsTrustedCommonName
	opaquePassword    *string
	select_           *[]string
}

func (r ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest) Body(body ClientCertAuthorityOcspTlsTrustedCommonName) ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest {
	r.body = &body
	return r
}
func (r ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest) OpaquePassword(opaquePassword string) ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest {
	r.opaquePassword = &opaquePassword
	return r
}
func (r ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest) Select_(select_ []string) ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest {
	r.select_ = &select_
	return r
}

func (r ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest) Execute() (ClientCertAuthorityOcspTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateClientCertAuthorityOcspTlsTrustedCommonNameExecute(r)
}

/*
 * CreateClientCertAuthorityOcspTlsTrustedCommonName Create an OCSP Responder Trusted Common Name object.
 * Create an OCSP Responder Trusted Common Name object. Any attribute missing from the request will be set to its default value.

When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.


Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque
:---|:---:|:---:|:---:|:---:|:---:|:---:
certAuthorityName|x||x|||
ocspTlsTrustedCommonName|x|x||||



A SEMP client authorized with a minimum access scope/level of "global/admin" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest
*/
func (a *ClientCertAuthorityApiService) CreateClientCertAuthorityOcspTlsTrustedCommonName(ctx _context.Context, certAuthorityName string) ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest {
	return ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthorityOcspTlsTrustedCommonNameResponse
 */
func (a *ClientCertAuthorityApiService) CreateClientCertAuthorityOcspTlsTrustedCommonNameExecute(r ClientCertAuthorityApiApiCreateClientCertAuthorityOcspTlsTrustedCommonNameRequest) (ClientCertAuthorityOcspTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthorityOcspTlsTrustedCommonNameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCertAuthorityApiService.CreateClientCertAuthorityOcspTlsTrustedCommonName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.opaquePassword != nil {
		localVarQueryParams.Add("opaquePassword", parameterToString(*r.opaquePassword, ""))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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

type ClientCertAuthorityApiApiDeleteClientCertAuthorityRequest struct {
	ctx               _context.Context
	ApiService        *ClientCertAuthorityApiService
	certAuthorityName string
}

func (r ClientCertAuthorityApiApiDeleteClientCertAuthorityRequest) Execute() (SempMetaOnlyResponse, *_nethttp.Response, error) {
	return r.ApiService.DeleteClientCertAuthorityExecute(r)
}

/*
 * DeleteClientCertAuthority Delete a Client Certificate Authority object.
 * Delete a Client Certificate Authority object.

Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.

A SEMP client authorized with a minimum access scope/level of "global/admin" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return ClientCertAuthorityApiApiDeleteClientCertAuthorityRequest
*/
func (a *ClientCertAuthorityApiService) DeleteClientCertAuthority(ctx _context.Context, certAuthorityName string) ClientCertAuthorityApiApiDeleteClientCertAuthorityRequest {
	return ClientCertAuthorityApiApiDeleteClientCertAuthorityRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return SempMetaOnlyResponse
 */
func (a *ClientCertAuthorityApiService) DeleteClientCertAuthorityExecute(r ClientCertAuthorityApiApiDeleteClientCertAuthorityRequest) (SempMetaOnlyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SempMetaOnlyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCertAuthorityApiService.DeleteClientCertAuthority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ClientCertAuthorityApiApiDeleteClientCertAuthorityOcspTlsTrustedCommonNameRequest struct {
	ctx                      _context.Context
	ApiService               *ClientCertAuthorityApiService
	certAuthorityName        string
	ocspTlsTrustedCommonName string
}

func (r ClientCertAuthorityApiApiDeleteClientCertAuthorityOcspTlsTrustedCommonNameRequest) Execute() (SempMetaOnlyResponse, *_nethttp.Response, error) {
	return r.ApiService.DeleteClientCertAuthorityOcspTlsTrustedCommonNameExecute(r)
}

/*
 * DeleteClientCertAuthorityOcspTlsTrustedCommonName Delete an OCSP Responder Trusted Common Name object.
 * Delete an OCSP Responder Trusted Common Name object.

When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.

A SEMP client authorized with a minimum access scope/level of "global/admin" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @param ocspTlsTrustedCommonName The expected Trusted Common Name of the OCSP responder remote certificate.
 * @return ClientCertAuthorityApiApiDeleteClientCertAuthorityOcspTlsTrustedCommonNameRequest
*/
func (a *ClientCertAuthorityApiService) DeleteClientCertAuthorityOcspTlsTrustedCommonName(ctx _context.Context, certAuthorityName string, ocspTlsTrustedCommonName string) ClientCertAuthorityApiApiDeleteClientCertAuthorityOcspTlsTrustedCommonNameRequest {
	return ClientCertAuthorityApiApiDeleteClientCertAuthorityOcspTlsTrustedCommonNameRequest{
		ApiService:               a,
		ctx:                      ctx,
		certAuthorityName:        certAuthorityName,
		ocspTlsTrustedCommonName: ocspTlsTrustedCommonName,
	}
}

/*
 * Execute executes the request
 * @return SempMetaOnlyResponse
 */
func (a *ClientCertAuthorityApiService) DeleteClientCertAuthorityOcspTlsTrustedCommonNameExecute(r ClientCertAuthorityApiApiDeleteClientCertAuthorityOcspTlsTrustedCommonNameRequest) (SempMetaOnlyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SempMetaOnlyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCertAuthorityApiService.DeleteClientCertAuthorityOcspTlsTrustedCommonName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ocspTlsTrustedCommonName"+"}", _neturl.PathEscape(parameterToString(r.ocspTlsTrustedCommonName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest struct {
	ctx            _context.Context
	ApiService     *ClientCertAuthorityApiService
	count          *int32
	cursor         *string
	opaquePassword *string
	where          *[]string
	select_        *[]string
}

func (r ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest) Count(count int32) ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest {
	r.count = &count
	return r
}
func (r ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest) Cursor(cursor string) ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest {
	r.cursor = &cursor
	return r
}
func (r ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest) OpaquePassword(opaquePassword string) ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest {
	r.opaquePassword = &opaquePassword
	return r
}
func (r ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest) Where(where []string) ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest {
	r.where = &where
	return r
}
func (r ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest) Select_(select_ []string) ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest {
	r.select_ = &select_
	return r
}

func (r ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest) Execute() (ClientCertAuthoritiesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetClientCertAuthoritiesExecute(r)
}

/*
 * GetClientCertAuthorities Get a list of Client Certificate Authority objects.
 * Get a list of Client Certificate Authority objects.

Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.


Attribute|Identifying|Write-Only|Deprecated|Opaque
:---|:---:|:---:|:---:|:---:
certAuthorityName|x|||



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest
*/
func (a *ClientCertAuthorityApiService) GetClientCertAuthorities(ctx _context.Context) ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest {
	return ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthoritiesResponse
 */
func (a *ClientCertAuthorityApiService) GetClientCertAuthoritiesExecute(r ClientCertAuthorityApiApiGetClientCertAuthoritiesRequest) (ClientCertAuthoritiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthoritiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCertAuthorityApiService.GetClientCertAuthorities")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.opaquePassword != nil {
		localVarQueryParams.Add("opaquePassword", parameterToString(*r.opaquePassword, ""))
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

type ClientCertAuthorityApiApiGetClientCertAuthorityRequest struct {
	ctx               _context.Context
	ApiService        *ClientCertAuthorityApiService
	certAuthorityName string
	opaquePassword    *string
	select_           *[]string
}

func (r ClientCertAuthorityApiApiGetClientCertAuthorityRequest) OpaquePassword(opaquePassword string) ClientCertAuthorityApiApiGetClientCertAuthorityRequest {
	r.opaquePassword = &opaquePassword
	return r
}
func (r ClientCertAuthorityApiApiGetClientCertAuthorityRequest) Select_(select_ []string) ClientCertAuthorityApiApiGetClientCertAuthorityRequest {
	r.select_ = &select_
	return r
}

func (r ClientCertAuthorityApiApiGetClientCertAuthorityRequest) Execute() (ClientCertAuthorityResponse, *_nethttp.Response, error) {
	return r.ApiService.GetClientCertAuthorityExecute(r)
}

/*
 * GetClientCertAuthority Get a Client Certificate Authority object.
 * Get a Client Certificate Authority object.

Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.


Attribute|Identifying|Write-Only|Deprecated|Opaque
:---|:---:|:---:|:---:|:---:
certAuthorityName|x|||



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return ClientCertAuthorityApiApiGetClientCertAuthorityRequest
*/
func (a *ClientCertAuthorityApiService) GetClientCertAuthority(ctx _context.Context, certAuthorityName string) ClientCertAuthorityApiApiGetClientCertAuthorityRequest {
	return ClientCertAuthorityApiApiGetClientCertAuthorityRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthorityResponse
 */
func (a *ClientCertAuthorityApiService) GetClientCertAuthorityExecute(r ClientCertAuthorityApiApiGetClientCertAuthorityRequest) (ClientCertAuthorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCertAuthorityApiService.GetClientCertAuthority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.opaquePassword != nil {
		localVarQueryParams.Add("opaquePassword", parameterToString(*r.opaquePassword, ""))
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

type ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest struct {
	ctx                      _context.Context
	ApiService               *ClientCertAuthorityApiService
	certAuthorityName        string
	ocspTlsTrustedCommonName string
	opaquePassword           *string
	select_                  *[]string
}

func (r ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest) OpaquePassword(opaquePassword string) ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest {
	r.opaquePassword = &opaquePassword
	return r
}
func (r ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest) Select_(select_ []string) ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest {
	r.select_ = &select_
	return r
}

func (r ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest) Execute() (ClientCertAuthorityOcspTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	return r.ApiService.GetClientCertAuthorityOcspTlsTrustedCommonNameExecute(r)
}

/*
 * GetClientCertAuthorityOcspTlsTrustedCommonName Get an OCSP Responder Trusted Common Name object.
 * Get an OCSP Responder Trusted Common Name object.

When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.


Attribute|Identifying|Write-Only|Deprecated|Opaque
:---|:---:|:---:|:---:|:---:
certAuthorityName|x|||
ocspTlsTrustedCommonName|x|||



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @param ocspTlsTrustedCommonName The expected Trusted Common Name of the OCSP responder remote certificate.
 * @return ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest
*/
func (a *ClientCertAuthorityApiService) GetClientCertAuthorityOcspTlsTrustedCommonName(ctx _context.Context, certAuthorityName string, ocspTlsTrustedCommonName string) ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest {
	return ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest{
		ApiService:               a,
		ctx:                      ctx,
		certAuthorityName:        certAuthorityName,
		ocspTlsTrustedCommonName: ocspTlsTrustedCommonName,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthorityOcspTlsTrustedCommonNameResponse
 */
func (a *ClientCertAuthorityApiService) GetClientCertAuthorityOcspTlsTrustedCommonNameExecute(r ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest) (ClientCertAuthorityOcspTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthorityOcspTlsTrustedCommonNameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCertAuthorityApiService.GetClientCertAuthorityOcspTlsTrustedCommonName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ocspTlsTrustedCommonName"+"}", _neturl.PathEscape(parameterToString(r.ocspTlsTrustedCommonName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.opaquePassword != nil {
		localVarQueryParams.Add("opaquePassword", parameterToString(*r.opaquePassword, ""))
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

type ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest struct {
	ctx               _context.Context
	ApiService        *ClientCertAuthorityApiService
	certAuthorityName string
	opaquePassword    *string
	where             *[]string
	select_           *[]string
}

func (r ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest) OpaquePassword(opaquePassword string) ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest {
	r.opaquePassword = &opaquePassword
	return r
}
func (r ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest) Where(where []string) ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest {
	r.where = &where
	return r
}
func (r ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest) Select_(select_ []string) ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest {
	r.select_ = &select_
	return r
}

func (r ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest) Execute() (ClientCertAuthorityOcspTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetClientCertAuthorityOcspTlsTrustedCommonNamesExecute(r)
}

/*
 * GetClientCertAuthorityOcspTlsTrustedCommonNames Get a list of OCSP Responder Trusted Common Name objects.
 * Get a list of OCSP Responder Trusted Common Name objects.

When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.


Attribute|Identifying|Write-Only|Deprecated|Opaque
:---|:---:|:---:|:---:|:---:
certAuthorityName|x|||
ocspTlsTrustedCommonName|x|||



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest
*/
func (a *ClientCertAuthorityApiService) GetClientCertAuthorityOcspTlsTrustedCommonNames(ctx _context.Context, certAuthorityName string) ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest {
	return ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthorityOcspTlsTrustedCommonNamesResponse
 */
func (a *ClientCertAuthorityApiService) GetClientCertAuthorityOcspTlsTrustedCommonNamesExecute(r ClientCertAuthorityApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest) (ClientCertAuthorityOcspTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthorityOcspTlsTrustedCommonNamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCertAuthorityApiService.GetClientCertAuthorityOcspTlsTrustedCommonNames")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.opaquePassword != nil {
		localVarQueryParams.Add("opaquePassword", parameterToString(*r.opaquePassword, ""))
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

type ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest struct {
	ctx               _context.Context
	ApiService        *ClientCertAuthorityApiService
	certAuthorityName string
	body              *ClientCertAuthority
	opaquePassword    *string
	select_           *[]string
}

func (r ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest) Body(body ClientCertAuthority) ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest {
	r.body = &body
	return r
}
func (r ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest) OpaquePassword(opaquePassword string) ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest {
	r.opaquePassword = &opaquePassword
	return r
}
func (r ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest) Select_(select_ []string) ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest {
	r.select_ = &select_
	return r
}

func (r ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest) Execute() (ClientCertAuthorityResponse, *_nethttp.Response, error) {
	return r.ApiService.ReplaceClientCertAuthorityExecute(r)
}

/*
 * ReplaceClientCertAuthority Replace a Client Certificate Authority object.
 * Replace a Client Certificate Authority object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.

Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.


Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated|Opaque
:---|:---:|:---:|:---:|:---:|:---:|:---:
certAuthorityName|x|x||||
crlUrl||||x||



The following attributes in the request may only be provided in certain combinations with other attributes:


Class|Attribute|Requires|Conflicts
:---|:---|:---|:---
ClientCertAuthority|crlDayList|crlTimeList|
ClientCertAuthority|crlTimeList|crlDayList|



A SEMP client authorized with a minimum access scope/level of "global/admin" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest
*/
func (a *ClientCertAuthorityApiService) ReplaceClientCertAuthority(ctx _context.Context, certAuthorityName string) ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest {
	return ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthorityResponse
 */
func (a *ClientCertAuthorityApiService) ReplaceClientCertAuthorityExecute(r ClientCertAuthorityApiApiReplaceClientCertAuthorityRequest) (ClientCertAuthorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCertAuthorityApiService.ReplaceClientCertAuthority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.opaquePassword != nil {
		localVarQueryParams.Add("opaquePassword", parameterToString(*r.opaquePassword, ""))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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

type ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest struct {
	ctx               _context.Context
	ApiService        *ClientCertAuthorityApiService
	certAuthorityName string
	body              *ClientCertAuthority
	opaquePassword    *string
	select_           *[]string
}

func (r ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest) Body(body ClientCertAuthority) ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest {
	r.body = &body
	return r
}
func (r ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest) OpaquePassword(opaquePassword string) ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest {
	r.opaquePassword = &opaquePassword
	return r
}
func (r ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest) Select_(select_ []string) ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest {
	r.select_ = &select_
	return r
}

func (r ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest) Execute() (ClientCertAuthorityResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateClientCertAuthorityExecute(r)
}

/*
 * UpdateClientCertAuthority Update a Client Certificate Authority object.
 * Update a Client Certificate Authority object. Any attribute missing from the request will be left unchanged.

Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.


Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated|Opaque
:---|:---:|:---:|:---:|:---:|:---:|:---:
certAuthorityName|x|x||||
crlUrl||||x||



The following attributes in the request may only be provided in certain combinations with other attributes:


Class|Attribute|Requires|Conflicts
:---|:---|:---|:---
ClientCertAuthority|crlDayList|crlTimeList|
ClientCertAuthority|crlTimeList|crlDayList|



A SEMP client authorized with a minimum access scope/level of "global/admin" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest
*/
func (a *ClientCertAuthorityApiService) UpdateClientCertAuthority(ctx _context.Context, certAuthorityName string) ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest {
	return ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthorityResponse
 */
func (a *ClientCertAuthorityApiService) UpdateClientCertAuthorityExecute(r ClientCertAuthorityApiApiUpdateClientCertAuthorityRequest) (ClientCertAuthorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientCertAuthorityApiService.UpdateClientCertAuthority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.opaquePassword != nil {
		localVarQueryParams.Add("opaquePassword", parameterToString(*r.opaquePassword, ""))
	}
	if r.select_ != nil {
		localVarQueryParams.Add("select", parameterToString(*r.select_, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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
