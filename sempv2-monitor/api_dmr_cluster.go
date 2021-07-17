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

// DmrClusterApiService DmrClusterApi service
type DmrClusterApiService service

type DmrClusterApiApiGetDmrClusterRequest struct {
	ctx            _context.Context
	ApiService     *DmrClusterApiService
	dmrClusterName string
	select_        *[]string
}

func (r DmrClusterApiApiGetDmrClusterRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterRequest) Execute() (DmrClusterResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterExecute(r)
}

/*
 * GetDmrCluster Get a Cluster object.
 * Get a Cluster object.

A Cluster is a provisioned object on a message broker that contains global DMR configuration parameters.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|
tlsServerCertEnforceTrustedCommonNameEnabled||x



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @return DmrClusterApiApiGetDmrClusterRequest
*/
func (a *DmrClusterApiService) GetDmrCluster(ctx _context.Context, dmrClusterName string) DmrClusterApiApiGetDmrClusterRequest {
	return DmrClusterApiApiGetDmrClusterRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterResponse
 */
func (a *DmrClusterApiService) GetDmrClusterExecute(r DmrClusterApiApiGetDmrClusterRequest) (DmrClusterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrCluster")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)

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

type DmrClusterApiApiGetDmrClusterLinkRequest struct {
	ctx            _context.Context
	ApiService     *DmrClusterApiService
	dmrClusterName string
	remoteNodeName string
	select_        *[]string
}

func (r DmrClusterApiApiGetDmrClusterLinkRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterLinkRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterLinkRequest) Execute() (DmrClusterLinkResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterLinkExecute(r)
}

/*
 * GetDmrClusterLink Get a Link object.
 * Get a Link object.

A Link connects nodes (either within a Cluster or between two different Clusters) and allows them to exchange topology information, subscriptions and data.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|
remoteNodeName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @param remoteNodeName The name of the node at the remote end of the Link.
 * @return DmrClusterApiApiGetDmrClusterLinkRequest
*/
func (a *DmrClusterApiService) GetDmrClusterLink(ctx _context.Context, dmrClusterName string, remoteNodeName string) DmrClusterApiApiGetDmrClusterLinkRequest {
	return DmrClusterApiApiGetDmrClusterLinkRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
		remoteNodeName: remoteNodeName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterLinkResponse
 */
func (a *DmrClusterApiService) GetDmrClusterLinkExecute(r DmrClusterApiApiGetDmrClusterLinkRequest) (DmrClusterLinkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusterLink")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}/links/{remoteNodeName}"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)
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

type DmrClusterApiApiGetDmrClusterLinkChannelRequest struct {
	ctx            _context.Context
	ApiService     *DmrClusterApiService
	dmrClusterName string
	remoteNodeName string
	msgVpnName     string
	select_        *[]string
}

func (r DmrClusterApiApiGetDmrClusterLinkChannelRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterLinkChannelRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterLinkChannelRequest) Execute() (DmrClusterLinkChannelResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterLinkChannelExecute(r)
}

/*
 * GetDmrClusterLinkChannel Get a Cluster Link Channels object.
 * Get a Cluster Link Channels object.

A Channel is a connection between this broker and a remote node in the Cluster.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|
msgVpnName|x|
remoteNodeName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @param remoteNodeName The name of the node at the remote end of the Link.
 * @param msgVpnName The name of the Message VPN.
 * @return DmrClusterApiApiGetDmrClusterLinkChannelRequest
*/
func (a *DmrClusterApiService) GetDmrClusterLinkChannel(ctx _context.Context, dmrClusterName string, remoteNodeName string, msgVpnName string) DmrClusterApiApiGetDmrClusterLinkChannelRequest {
	return DmrClusterApiApiGetDmrClusterLinkChannelRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
		remoteNodeName: remoteNodeName,
		msgVpnName:     msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterLinkChannelResponse
 */
func (a *DmrClusterApiService) GetDmrClusterLinkChannelExecute(r DmrClusterApiApiGetDmrClusterLinkChannelRequest) (DmrClusterLinkChannelResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkChannelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusterLinkChannel")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}/links/{remoteNodeName}/channels/{msgVpnName}"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteNodeName"+"}", _neturl.PathEscape(parameterToString(r.remoteNodeName, "")), -1)
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

type DmrClusterApiApiGetDmrClusterLinkChannelsRequest struct {
	ctx            _context.Context
	ApiService     *DmrClusterApiService
	dmrClusterName string
	remoteNodeName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r DmrClusterApiApiGetDmrClusterLinkChannelsRequest) Count(count int32) DmrClusterApiApiGetDmrClusterLinkChannelsRequest {
	r.count = &count
	return r
}
func (r DmrClusterApiApiGetDmrClusterLinkChannelsRequest) Cursor(cursor string) DmrClusterApiApiGetDmrClusterLinkChannelsRequest {
	r.cursor = &cursor
	return r
}
func (r DmrClusterApiApiGetDmrClusterLinkChannelsRequest) Where(where []string) DmrClusterApiApiGetDmrClusterLinkChannelsRequest {
	r.where = &where
	return r
}
func (r DmrClusterApiApiGetDmrClusterLinkChannelsRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterLinkChannelsRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterLinkChannelsRequest) Execute() (DmrClusterLinkChannelsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterLinkChannelsExecute(r)
}

/*
 * GetDmrClusterLinkChannels Get a list of Cluster Link Channels objects.
 * Get a list of Cluster Link Channels objects.

A Channel is a connection between this broker and a remote node in the Cluster.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|
msgVpnName|x|
remoteNodeName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @param remoteNodeName The name of the node at the remote end of the Link.
 * @return DmrClusterApiApiGetDmrClusterLinkChannelsRequest
*/
func (a *DmrClusterApiService) GetDmrClusterLinkChannels(ctx _context.Context, dmrClusterName string, remoteNodeName string) DmrClusterApiApiGetDmrClusterLinkChannelsRequest {
	return DmrClusterApiApiGetDmrClusterLinkChannelsRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
		remoteNodeName: remoteNodeName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterLinkChannelsResponse
 */
func (a *DmrClusterApiService) GetDmrClusterLinkChannelsExecute(r DmrClusterApiApiGetDmrClusterLinkChannelsRequest) (DmrClusterLinkChannelsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkChannelsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusterLinkChannels")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}/links/{remoteNodeName}/channels"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteNodeName"+"}", _neturl.PathEscape(parameterToString(r.remoteNodeName, "")), -1)

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

type DmrClusterApiApiGetDmrClusterLinkRemoteAddressRequest struct {
	ctx            _context.Context
	ApiService     *DmrClusterApiService
	dmrClusterName string
	remoteNodeName string
	remoteAddress  string
	select_        *[]string
}

func (r DmrClusterApiApiGetDmrClusterLinkRemoteAddressRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterLinkRemoteAddressRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterLinkRemoteAddressRequest) Execute() (DmrClusterLinkRemoteAddressResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterLinkRemoteAddressExecute(r)
}

/*
 * GetDmrClusterLinkRemoteAddress Get a Remote Address object.
 * Get a Remote Address object.

Each Remote Address, consisting of a FQDN or IP address and optional port, is used to connect to the remote node for this Link. Up to 4 addresses may be provided for each Link, and will be tried on a round-robin basis.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|
remoteAddress|x|
remoteNodeName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @param remoteNodeName The name of the node at the remote end of the Link.
 * @param remoteAddress The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed).
 * @return DmrClusterApiApiGetDmrClusterLinkRemoteAddressRequest
*/
func (a *DmrClusterApiService) GetDmrClusterLinkRemoteAddress(ctx _context.Context, dmrClusterName string, remoteNodeName string, remoteAddress string) DmrClusterApiApiGetDmrClusterLinkRemoteAddressRequest {
	return DmrClusterApiApiGetDmrClusterLinkRemoteAddressRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
		remoteNodeName: remoteNodeName,
		remoteAddress:  remoteAddress,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterLinkRemoteAddressResponse
 */
func (a *DmrClusterApiService) GetDmrClusterLinkRemoteAddressExecute(r DmrClusterApiApiGetDmrClusterLinkRemoteAddressRequest) (DmrClusterLinkRemoteAddressResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkRemoteAddressResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusterLinkRemoteAddress")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses/{remoteAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteNodeName"+"}", _neturl.PathEscape(parameterToString(r.remoteNodeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteAddress"+"}", _neturl.PathEscape(parameterToString(r.remoteAddress, "")), -1)

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

type DmrClusterApiApiGetDmrClusterLinkRemoteAddressesRequest struct {
	ctx            _context.Context
	ApiService     *DmrClusterApiService
	dmrClusterName string
	remoteNodeName string
	where          *[]string
	select_        *[]string
}

func (r DmrClusterApiApiGetDmrClusterLinkRemoteAddressesRequest) Where(where []string) DmrClusterApiApiGetDmrClusterLinkRemoteAddressesRequest {
	r.where = &where
	return r
}
func (r DmrClusterApiApiGetDmrClusterLinkRemoteAddressesRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterLinkRemoteAddressesRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterLinkRemoteAddressesRequest) Execute() (DmrClusterLinkRemoteAddressesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterLinkRemoteAddressesExecute(r)
}

/*
 * GetDmrClusterLinkRemoteAddresses Get a list of Remote Address objects.
 * Get a list of Remote Address objects.

Each Remote Address, consisting of a FQDN or IP address and optional port, is used to connect to the remote node for this Link. Up to 4 addresses may be provided for each Link, and will be tried on a round-robin basis.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|
remoteAddress|x|
remoteNodeName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @param remoteNodeName The name of the node at the remote end of the Link.
 * @return DmrClusterApiApiGetDmrClusterLinkRemoteAddressesRequest
*/
func (a *DmrClusterApiService) GetDmrClusterLinkRemoteAddresses(ctx _context.Context, dmrClusterName string, remoteNodeName string) DmrClusterApiApiGetDmrClusterLinkRemoteAddressesRequest {
	return DmrClusterApiApiGetDmrClusterLinkRemoteAddressesRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
		remoteNodeName: remoteNodeName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterLinkRemoteAddressesResponse
 */
func (a *DmrClusterApiService) GetDmrClusterLinkRemoteAddressesExecute(r DmrClusterApiApiGetDmrClusterLinkRemoteAddressesRequest) (DmrClusterLinkRemoteAddressesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkRemoteAddressesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusterLinkRemoteAddresses")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteNodeName"+"}", _neturl.PathEscape(parameterToString(r.remoteNodeName, "")), -1)

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

type DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest struct {
	ctx                  _context.Context
	ApiService           *DmrClusterApiService
	dmrClusterName       string
	remoteNodeName       string
	tlsTrustedCommonName string
	select_              *[]string
}

func (r DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest) Execute() (DmrClusterLinkTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterLinkTlsTrustedCommonNameExecute(r)
}

/*
 * GetDmrClusterLinkTlsTrustedCommonName Get a Trusted Common Name object.
 * Get a Trusted Common Name object.

The Trusted Common Names for the Link are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|x
remoteNodeName|x|x
tlsTrustedCommonName|x|x



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @param remoteNodeName The name of the node at the remote end of the Link.
 * @param tlsTrustedCommonName The expected trusted common name of the remote certificate.
 * @return DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest
*/
func (a *DmrClusterApiService) GetDmrClusterLinkTlsTrustedCommonName(ctx _context.Context, dmrClusterName string, remoteNodeName string, tlsTrustedCommonName string) DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest {
	return DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest{
		ApiService:           a,
		ctx:                  ctx,
		dmrClusterName:       dmrClusterName,
		remoteNodeName:       remoteNodeName,
		tlsTrustedCommonName: tlsTrustedCommonName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterLinkTlsTrustedCommonNameResponse
 */
func (a *DmrClusterApiService) GetDmrClusterLinkTlsTrustedCommonNameExecute(r DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest) (DmrClusterLinkTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkTlsTrustedCommonNameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusterLinkTlsTrustedCommonName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteNodeName"+"}", _neturl.PathEscape(parameterToString(r.remoteNodeName, "")), -1)
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

type DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest struct {
	ctx            _context.Context
	ApiService     *DmrClusterApiService
	dmrClusterName string
	remoteNodeName string
	where          *[]string
	select_        *[]string
}

func (r DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest) Where(where []string) DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest {
	r.where = &where
	return r
}
func (r DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest) Execute() (DmrClusterLinkTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterLinkTlsTrustedCommonNamesExecute(r)
}

/*
 * GetDmrClusterLinkTlsTrustedCommonNames Get a list of Trusted Common Name objects.
 * Get a list of Trusted Common Name objects.

The Trusted Common Names for the Link are used by encrypted transports to verify the name in the certificate presented by the remote node. They must include the common name of the remote node's server certificate or client certificate, depending upon the initiator of the connection.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|x
remoteNodeName|x|x
tlsTrustedCommonName|x|x



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @param remoteNodeName The name of the node at the remote end of the Link.
 * @return DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest
*/
func (a *DmrClusterApiService) GetDmrClusterLinkTlsTrustedCommonNames(ctx _context.Context, dmrClusterName string, remoteNodeName string) DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest {
	return DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
		remoteNodeName: remoteNodeName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterLinkTlsTrustedCommonNamesResponse
 */
func (a *DmrClusterApiService) GetDmrClusterLinkTlsTrustedCommonNamesExecute(r DmrClusterApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest) (DmrClusterLinkTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkTlsTrustedCommonNamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusterLinkTlsTrustedCommonNames")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"remoteNodeName"+"}", _neturl.PathEscape(parameterToString(r.remoteNodeName, "")), -1)

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

type DmrClusterApiApiGetDmrClusterLinksRequest struct {
	ctx            _context.Context
	ApiService     *DmrClusterApiService
	dmrClusterName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r DmrClusterApiApiGetDmrClusterLinksRequest) Count(count int32) DmrClusterApiApiGetDmrClusterLinksRequest {
	r.count = &count
	return r
}
func (r DmrClusterApiApiGetDmrClusterLinksRequest) Cursor(cursor string) DmrClusterApiApiGetDmrClusterLinksRequest {
	r.cursor = &cursor
	return r
}
func (r DmrClusterApiApiGetDmrClusterLinksRequest) Where(where []string) DmrClusterApiApiGetDmrClusterLinksRequest {
	r.where = &where
	return r
}
func (r DmrClusterApiApiGetDmrClusterLinksRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterLinksRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterLinksRequest) Execute() (DmrClusterLinksResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterLinksExecute(r)
}

/*
 * GetDmrClusterLinks Get a list of Link objects.
 * Get a list of Link objects.

A Link connects nodes (either within a Cluster or between two different Clusters) and allows them to exchange topology information, subscriptions and data.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|
remoteNodeName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @return DmrClusterApiApiGetDmrClusterLinksRequest
*/
func (a *DmrClusterApiService) GetDmrClusterLinks(ctx _context.Context, dmrClusterName string) DmrClusterApiApiGetDmrClusterLinksRequest {
	return DmrClusterApiApiGetDmrClusterLinksRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterLinksResponse
 */
func (a *DmrClusterApiService) GetDmrClusterLinksExecute(r DmrClusterApiApiGetDmrClusterLinksRequest) (DmrClusterLinksResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusterLinks")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)

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

type DmrClusterApiApiGetDmrClusterTopologyIssueRequest struct {
	ctx            _context.Context
	ApiService     *DmrClusterApiService
	dmrClusterName string
	topologyIssue  string
	select_        *[]string
}

func (r DmrClusterApiApiGetDmrClusterTopologyIssueRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterTopologyIssueRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterTopologyIssueRequest) Execute() (DmrClusterTopologyIssueResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterTopologyIssueExecute(r)
}

/*
 * GetDmrClusterTopologyIssue Get a Cluster Topology Issue object.
 * Get a Cluster Topology Issue object.

A Cluster Topology Issue indicates incorrect or inconsistent configuration within the DMR network. Such issues will cause messages to be misdelivered or lost.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|
topologyIssue|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @param topologyIssue The topology issue discovered in the Cluster. A topology issue indicates incorrect or inconsistent configuration within the DMR network. Such issues will cause messages to be misdelivered or lost.
 * @return DmrClusterApiApiGetDmrClusterTopologyIssueRequest
*/
func (a *DmrClusterApiService) GetDmrClusterTopologyIssue(ctx _context.Context, dmrClusterName string, topologyIssue string) DmrClusterApiApiGetDmrClusterTopologyIssueRequest {
	return DmrClusterApiApiGetDmrClusterTopologyIssueRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
		topologyIssue:  topologyIssue,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterTopologyIssueResponse
 */
func (a *DmrClusterApiService) GetDmrClusterTopologyIssueExecute(r DmrClusterApiApiGetDmrClusterTopologyIssueRequest) (DmrClusterTopologyIssueResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterTopologyIssueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusterTopologyIssue")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}/topologyIssues/{topologyIssue}"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topologyIssue"+"}", _neturl.PathEscape(parameterToString(r.topologyIssue, "")), -1)

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

type DmrClusterApiApiGetDmrClusterTopologyIssuesRequest struct {
	ctx            _context.Context
	ApiService     *DmrClusterApiService
	dmrClusterName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r DmrClusterApiApiGetDmrClusterTopologyIssuesRequest) Count(count int32) DmrClusterApiApiGetDmrClusterTopologyIssuesRequest {
	r.count = &count
	return r
}
func (r DmrClusterApiApiGetDmrClusterTopologyIssuesRequest) Cursor(cursor string) DmrClusterApiApiGetDmrClusterTopologyIssuesRequest {
	r.cursor = &cursor
	return r
}
func (r DmrClusterApiApiGetDmrClusterTopologyIssuesRequest) Where(where []string) DmrClusterApiApiGetDmrClusterTopologyIssuesRequest {
	r.where = &where
	return r
}
func (r DmrClusterApiApiGetDmrClusterTopologyIssuesRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClusterTopologyIssuesRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClusterTopologyIssuesRequest) Execute() (DmrClusterTopologyIssuesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClusterTopologyIssuesExecute(r)
}

/*
 * GetDmrClusterTopologyIssues Get a list of Cluster Topology Issue objects.
 * Get a list of Cluster Topology Issue objects.

A Cluster Topology Issue indicates incorrect or inconsistent configuration within the DMR network. Such issues will cause messages to be misdelivered or lost.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|
topologyIssue|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dmrClusterName The name of the Cluster.
 * @return DmrClusterApiApiGetDmrClusterTopologyIssuesRequest
*/
func (a *DmrClusterApiService) GetDmrClusterTopologyIssues(ctx _context.Context, dmrClusterName string) DmrClusterApiApiGetDmrClusterTopologyIssuesRequest {
	return DmrClusterApiApiGetDmrClusterTopologyIssuesRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterTopologyIssuesResponse
 */
func (a *DmrClusterApiService) GetDmrClusterTopologyIssuesExecute(r DmrClusterApiApiGetDmrClusterTopologyIssuesRequest) (DmrClusterTopologyIssuesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterTopologyIssuesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusterTopologyIssues")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters/{dmrClusterName}/topologyIssues"
	localVarPath = strings.Replace(localVarPath, "{"+"dmrClusterName"+"}", _neturl.PathEscape(parameterToString(r.dmrClusterName, "")), -1)

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

type DmrClusterApiApiGetDmrClustersRequest struct {
	ctx        _context.Context
	ApiService *DmrClusterApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r DmrClusterApiApiGetDmrClustersRequest) Count(count int32) DmrClusterApiApiGetDmrClustersRequest {
	r.count = &count
	return r
}
func (r DmrClusterApiApiGetDmrClustersRequest) Cursor(cursor string) DmrClusterApiApiGetDmrClustersRequest {
	r.cursor = &cursor
	return r
}
func (r DmrClusterApiApiGetDmrClustersRequest) Where(where []string) DmrClusterApiApiGetDmrClustersRequest {
	r.where = &where
	return r
}
func (r DmrClusterApiApiGetDmrClustersRequest) Select_(select_ []string) DmrClusterApiApiGetDmrClustersRequest {
	r.select_ = &select_
	return r
}

func (r DmrClusterApiApiGetDmrClustersRequest) Execute() (DmrClustersResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDmrClustersExecute(r)
}

/*
 * GetDmrClusters Get a list of Cluster objects.
 * Get a list of Cluster objects.

A Cluster is a provisioned object on a message broker that contains global DMR configuration parameters.


Attribute|Identifying|Deprecated
:---|:---:|:---:
dmrClusterName|x|
tlsServerCertEnforceTrustedCommonNameEnabled||x



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return DmrClusterApiApiGetDmrClustersRequest
*/
func (a *DmrClusterApiService) GetDmrClusters(ctx _context.Context) DmrClusterApiApiGetDmrClustersRequest {
	return DmrClusterApiApiGetDmrClustersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return DmrClustersResponse
 */
func (a *DmrClusterApiService) GetDmrClustersExecute(r DmrClusterApiApiGetDmrClustersRequest) (DmrClustersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClustersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DmrClusterApiService.GetDmrClusters")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dmrClusters"

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
