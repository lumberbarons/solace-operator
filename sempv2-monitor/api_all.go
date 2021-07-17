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

// AllApiService AllApi service
type AllApiService service

type AllApiApiGetAboutRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	select_    *[]string
}

func (r AllApiApiGetAboutRequest) Select_(select_ []string) AllApiApiGetAboutRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetAboutRequest) Execute() (AboutResponse, *_nethttp.Response, error) {
	return r.ApiService.GetAboutExecute(r)
}

/*
 * GetAbout Get an About object.
 * Get an About object.

This provides metadata about the SEMP API, such as the version of the API supported by the broker.



A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation.

This has been available since 2.13.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetAboutRequest
*/
func (a *AllApiService) GetAbout(ctx _context.Context) AllApiApiGetAboutRequest {
	return AllApiApiGetAboutRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return AboutResponse
 */
func (a *AllApiService) GetAboutExecute(r AllApiApiGetAboutRequest) (AboutResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AboutResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetAbout")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/about"

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

type AllApiApiGetAboutApiRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	select_    *[]string
}

func (r AllApiApiGetAboutApiRequest) Select_(select_ []string) AllApiApiGetAboutApiRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetAboutApiRequest) Execute() (AboutApiResponse, *_nethttp.Response, error) {
	return r.ApiService.GetAboutApiExecute(r)
}

/*
 * GetAboutApi Get an API Description object.
 * Get an API Description object.

The API Description object provides metadata about the SEMP API.



A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetAboutApiRequest
*/
func (a *AllApiService) GetAboutApi(ctx _context.Context) AllApiApiGetAboutApiRequest {
	return AllApiApiGetAboutApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return AboutApiResponse
 */
func (a *AllApiService) GetAboutApiExecute(r AllApiApiGetAboutApiRequest) (AboutApiResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AboutApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetAboutApi")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/about/api"

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

type AllApiApiGetAboutUserRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	select_    *[]string
}

func (r AllApiApiGetAboutUserRequest) Select_(select_ []string) AllApiApiGetAboutUserRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetAboutUserRequest) Execute() (AboutUserResponse, *_nethttp.Response, error) {
	return r.ApiService.GetAboutUserExecute(r)
}

/*
 * GetAboutUser Get a User object.
 * Get a User object.

Session and access level information about the user accessing the SEMP API.



A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetAboutUserRequest
*/
func (a *AllApiService) GetAboutUser(ctx _context.Context) AllApiApiGetAboutUserRequest {
	return AllApiApiGetAboutUserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return AboutUserResponse
 */
func (a *AllApiService) GetAboutUserExecute(r AllApiApiGetAboutUserRequest) (AboutUserResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AboutUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetAboutUser")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/about/user"

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

type AllApiApiGetAboutUserMsgVpnRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	select_    *[]string
}

func (r AllApiApiGetAboutUserMsgVpnRequest) Select_(select_ []string) AllApiApiGetAboutUserMsgVpnRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetAboutUserMsgVpnRequest) Execute() (AboutUserMsgVpnResponse, *_nethttp.Response, error) {
	return r.ApiService.GetAboutUserMsgVpnExecute(r)
}

/*
 * GetAboutUserMsgVpn Get a User Message VPN object.
 * Get a User Message VPN object.

This provides information about the Message VPN access level for the username used to access the SEMP API.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @return AllApiApiGetAboutUserMsgVpnRequest
*/
func (a *AllApiService) GetAboutUserMsgVpn(ctx _context.Context, msgVpnName string) AllApiApiGetAboutUserMsgVpnRequest {
	return AllApiApiGetAboutUserMsgVpnRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return AboutUserMsgVpnResponse
 */
func (a *AllApiService) GetAboutUserMsgVpnExecute(r AllApiApiGetAboutUserMsgVpnRequest) (AboutUserMsgVpnResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AboutUserMsgVpnResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetAboutUserMsgVpn")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/about/user/msgVpns/{msgVpnName}"
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

type AllApiApiGetAboutUserMsgVpnsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetAboutUserMsgVpnsRequest) Count(count int32) AllApiApiGetAboutUserMsgVpnsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetAboutUserMsgVpnsRequest) Cursor(cursor string) AllApiApiGetAboutUserMsgVpnsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetAboutUserMsgVpnsRequest) Where(where []string) AllApiApiGetAboutUserMsgVpnsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetAboutUserMsgVpnsRequest) Select_(select_ []string) AllApiApiGetAboutUserMsgVpnsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetAboutUserMsgVpnsRequest) Execute() (AboutUserMsgVpnsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetAboutUserMsgVpnsExecute(r)
}

/*
 * GetAboutUserMsgVpns Get a list of User Message VPN objects.
 * Get a list of User Message VPN objects.

This provides information about the Message VPN access level for the username used to access the SEMP API.


Attribute|Identifying|Deprecated
:---|:---:|:---:
msgVpnName|x|



A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation.

This has been available since 2.11.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetAboutUserMsgVpnsRequest
*/
func (a *AllApiService) GetAboutUserMsgVpns(ctx _context.Context) AllApiApiGetAboutUserMsgVpnsRequest {
	return AllApiApiGetAboutUserMsgVpnsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return AboutUserMsgVpnsResponse
 */
func (a *AllApiService) GetAboutUserMsgVpnsExecute(r AllApiApiGetAboutUserMsgVpnsRequest) (AboutUserMsgVpnsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AboutUserMsgVpnsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetAboutUserMsgVpns")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/about/user/msgVpns"

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

type AllApiApiGetBrokerRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	select_    *[]string
}

func (r AllApiApiGetBrokerRequest) Select_(select_ []string) AllApiApiGetBrokerRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetBrokerRequest) Execute() (BrokerResponse, *_nethttp.Response, error) {
	return r.ApiService.GetBrokerExecute(r)
}

/*
 * GetBroker Get a Broker object.
 * Get a Broker object.

This object contains global configuration for the message broker.



A SEMP client authorized with a minimum access scope/level of "global/none" is required to perform this operation. Requests which include the following attributes require greater access scope/level:


Attribute|Access Scope/Level
:---|:---:
averageRxByteRate|global/read-only
averageRxCompressedByteRate|global/read-only
averageRxMsgRate|global/read-only
averageRxUncompressedByteRate|global/read-only
averageTxByteRate|global/read-only
averageTxCompressedByteRate|global/read-only
averageTxMsgRate|global/read-only
averageTxUncompressedByteRate|global/read-only
cspfVersion|global/read-only
guaranteedMsgingDefragmentationEstimatedFragmentation|global/read-only
guaranteedMsgingDefragmentationEstimatedRecoverableSpace|global/read-only
guaranteedMsgingDefragmentationLastCompletedOn|global/read-only
guaranteedMsgingDefragmentationLastCompletionPercentage|global/read-only
guaranteedMsgingDefragmentationLastExitCondition|global/read-only
guaranteedMsgingDefragmentationLastExitConditionInformation|global/read-only
guaranteedMsgingDefragmentationStatus|global/read-only
guaranteedMsgingDefragmentationStatusActiveCompletionPercentage|global/read-only
guaranteedMsgingEnabled|global/read-only
guaranteedMsgingEventCacheUsageThreshold.clearPercent|global/read-only
guaranteedMsgingEventCacheUsageThreshold.clearValue|global/read-only
guaranteedMsgingEventCacheUsageThreshold.setPercent|global/read-only
guaranteedMsgingEventCacheUsageThreshold.setValue|global/read-only
guaranteedMsgingEventDeliveredUnackedThreshold.clearPercent|global/read-only
guaranteedMsgingEventDeliveredUnackedThreshold.setPercent|global/read-only
guaranteedMsgingEventDiskUsageThreshold.clearPercent|global/read-only
guaranteedMsgingEventDiskUsageThreshold.setPercent|global/read-only
guaranteedMsgingEventEgressFlowCountThreshold.clearPercent|global/read-only
guaranteedMsgingEventEgressFlowCountThreshold.clearValue|global/read-only
guaranteedMsgingEventEgressFlowCountThreshold.setPercent|global/read-only
guaranteedMsgingEventEgressFlowCountThreshold.setValue|global/read-only
guaranteedMsgingEventEndpointCountThreshold.clearPercent|global/read-only
guaranteedMsgingEventEndpointCountThreshold.clearValue|global/read-only
guaranteedMsgingEventEndpointCountThreshold.setPercent|global/read-only
guaranteedMsgingEventEndpointCountThreshold.setValue|global/read-only
guaranteedMsgingEventIngressFlowCountThreshold.clearPercent|global/read-only
guaranteedMsgingEventIngressFlowCountThreshold.clearValue|global/read-only
guaranteedMsgingEventIngressFlowCountThreshold.setPercent|global/read-only
guaranteedMsgingEventIngressFlowCountThreshold.setValue|global/read-only
guaranteedMsgingEventMsgCountThreshold.clearPercent|global/read-only
guaranteedMsgingEventMsgCountThreshold.setPercent|global/read-only
guaranteedMsgingEventMsgSpoolFileCountThreshold.clearPercent|global/read-only
guaranteedMsgingEventMsgSpoolFileCountThreshold.setPercent|global/read-only
guaranteedMsgingEventMsgSpoolUsageThreshold.clearPercent|global/read-only
guaranteedMsgingEventMsgSpoolUsageThreshold.clearValue|global/read-only
guaranteedMsgingEventMsgSpoolUsageThreshold.setPercent|global/read-only
guaranteedMsgingEventMsgSpoolUsageThreshold.setValue|global/read-only
guaranteedMsgingEventTransactedSessionCountThreshold.clearPercent|global/read-only
guaranteedMsgingEventTransactedSessionCountThreshold.clearValue|global/read-only
guaranteedMsgingEventTransactedSessionCountThreshold.setPercent|global/read-only
guaranteedMsgingEventTransactedSessionCountThreshold.setValue|global/read-only
guaranteedMsgingEventTransactedSessionResourceCountThreshold.clearPercent|global/read-only
guaranteedMsgingEventTransactedSessionResourceCountThreshold.setPercent|global/read-only
guaranteedMsgingEventTransactionCountThreshold.clearPercent|global/read-only
guaranteedMsgingEventTransactionCountThreshold.clearValue|global/read-only
guaranteedMsgingEventTransactionCountThreshold.setPercent|global/read-only
guaranteedMsgingEventTransactionCountThreshold.setValue|global/read-only
guaranteedMsgingMaxCacheUsage|global/read-only
guaranteedMsgingMaxMsgSpoolUsage|global/read-only
guaranteedMsgingMsgSpoolSyncMirroredMsgAckTimeout|global/read-only
guaranteedMsgingMsgSpoolSyncMirroredSpoolFileAckTimeout|global/read-only
guaranteedMsgingOperationalStatus|global/read-only
guaranteedMsgingTransactionReplicationCompatibilityMode|global/read-only
rxByteCount|global/read-only
rxByteRate|global/read-only
rxCompressedByteCount|global/read-only
rxCompressedByteRate|global/read-only
rxCompressionRatio|global/read-only
rxMsgCount|global/read-only
rxMsgRate|global/read-only
rxUncompressedByteCount|global/read-only
rxUncompressedByteRate|global/read-only
serviceAmqpEnabled|global/read-only
serviceAmqpTlsListenPort|global/read-only
serviceEventConnectionCountThreshold.clearPercent|global/read-only
serviceEventConnectionCountThreshold.clearValue|global/read-only
serviceEventConnectionCountThreshold.setPercent|global/read-only
serviceEventConnectionCountThreshold.setValue|global/read-only
serviceHealthCheckEnabled|global/read-only
serviceHealthCheckListenPort|global/read-only
serviceMateLinkEnabled|global/read-only
serviceMateLinkListenPort|global/read-only
serviceMqttEnabled|global/read-only
serviceMsgBackboneEnabled|global/read-only
serviceRedundancyEnabled|global/read-only
serviceRedundancyFirstListenPort|global/read-only
serviceRestEventOutgoingConnectionCountThreshold.clearPercent|global/read-only
serviceRestEventOutgoingConnectionCountThreshold.clearValue|global/read-only
serviceRestEventOutgoingConnectionCountThreshold.setPercent|global/read-only
serviceRestEventOutgoingConnectionCountThreshold.setValue|global/read-only
serviceRestIncomingEnabled|global/read-only
serviceRestOutgoingEnabled|global/read-only
serviceSempLegacyTimeoutEnabled|global/read-only
serviceSempPlainTextEnabled|global/read-only
serviceSempPlainTextListenPort|global/read-only
serviceSempSessionIdleTimeout|global/read-only
serviceSempSessionMaxLifetime|global/read-only
serviceSempTlsEnabled|global/read-only
serviceSempTlsListenPort|global/read-only
serviceSmfCompressionListenPort|global/read-only
serviceSmfEnabled|global/read-only
serviceSmfEventConnectionCountThreshold.clearPercent|global/read-only
serviceSmfEventConnectionCountThreshold.clearValue|global/read-only
serviceSmfEventConnectionCountThreshold.setPercent|global/read-only
serviceSmfEventConnectionCountThreshold.setValue|global/read-only
serviceSmfPlainTextListenPort|global/read-only
serviceSmfRoutingControlListenPort|global/read-only
serviceSmfTlsListenPort|global/read-only
serviceTlsEventConnectionCountThreshold.clearPercent|global/read-only
serviceTlsEventConnectionCountThreshold.clearValue|global/read-only
serviceTlsEventConnectionCountThreshold.setPercent|global/read-only
serviceTlsEventConnectionCountThreshold.setValue|global/read-only
serviceWebTransportEnabled|global/read-only
serviceWebTransportPlainTextListenPort|global/read-only
serviceWebTransportTlsListenPort|global/read-only
serviceWebTransportWebUrlSuffix|global/read-only
tlsBlockVersion11Enabled|global/read-only
tlsCipherSuiteManagementDefaultList|global/read-only
tlsCipherSuiteManagementList|global/read-only
tlsCipherSuiteManagementSupportedList|vpn/read-only
tlsCipherSuiteMsgBackboneDefaultList|global/read-only
tlsCipherSuiteMsgBackboneList|global/read-only
tlsCipherSuiteMsgBackboneSupportedList|vpn/read-only
tlsCipherSuiteSecureShellDefaultList|global/read-only
tlsCipherSuiteSecureShellList|global/read-only
tlsCipherSuiteSecureShellSupportedList|vpn/read-only
tlsCrimeExploitProtectionEnabled|global/read-only
tlsStandardDomainCertificateAuthoritiesEnabled|vpn/read-only
tlsTicketLifetime|global/read-only
tlsVersionSupportedList|vpn/read-only
txByteCount|global/read-only
txByteRate|global/read-only
txCompressedByteCount|global/read-only
txCompressedByteRate|global/read-only
txCompressionRatio|global/read-only
txMsgCount|global/read-only
txMsgRate|global/read-only
txUncompressedByteCount|global/read-only
txUncompressedByteRate|global/read-only



This has been available since 2.13.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetBrokerRequest
*/
func (a *AllApiService) GetBroker(ctx _context.Context) AllApiApiGetBrokerRequest {
	return AllApiApiGetBrokerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return BrokerResponse
 */
func (a *AllApiService) GetBrokerExecute(r AllApiApiGetBrokerRequest) (BrokerResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BrokerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetBroker")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/"

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

type AllApiApiGetCertAuthoritiesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetCertAuthoritiesRequest) Count(count int32) AllApiApiGetCertAuthoritiesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetCertAuthoritiesRequest) Cursor(cursor string) AllApiApiGetCertAuthoritiesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetCertAuthoritiesRequest) Where(where []string) AllApiApiGetCertAuthoritiesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetCertAuthoritiesRequest) Select_(select_ []string) AllApiApiGetCertAuthoritiesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetCertAuthoritiesRequest) Execute() (CertAuthoritiesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetCertAuthoritiesExecute(r)
}

/*
 * GetCertAuthorities Get a list of Certificate Authority objects.
 * Get a list of Certificate Authority objects.

Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|x
certContent||x
crlDayList||x
crlLastDownloadTime||x
crlLastFailureReason||x
crlLastFailureTime||x
crlNextDownloadTime||x
crlTimeList||x
crlUp||x
crlUrl||x
ocspLastFailureReason||x
ocspLastFailureTime||x
ocspLastFailureUrl||x
ocspNonResponderCertEnabled||x
ocspOverrideUrl||x
ocspTimeout||x
revocationCheckEnabled||x



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetCertAuthoritiesRequest
*/
func (a *AllApiService) GetCertAuthorities(ctx _context.Context) AllApiApiGetCertAuthoritiesRequest {
	return AllApiApiGetCertAuthoritiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return CertAuthoritiesResponse
 */
func (a *AllApiService) GetCertAuthoritiesExecute(r AllApiApiGetCertAuthoritiesRequest) (CertAuthoritiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CertAuthoritiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetCertAuthorities")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certAuthorities"

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

type AllApiApiGetCertAuthorityRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	certAuthorityName string
	select_           *[]string
}

func (r AllApiApiGetCertAuthorityRequest) Select_(select_ []string) AllApiApiGetCertAuthorityRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetCertAuthorityRequest) Execute() (CertAuthorityResponse, *_nethttp.Response, error) {
	return r.ApiService.GetCertAuthorityExecute(r)
}

/*
 * GetCertAuthority Get a Certificate Authority object.
 * Get a Certificate Authority object.

Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|x
certContent||x
crlDayList||x
crlLastDownloadTime||x
crlLastFailureReason||x
crlLastFailureTime||x
crlNextDownloadTime||x
crlTimeList||x
crlUp||x
crlUrl||x
ocspLastFailureReason||x
ocspLastFailureTime||x
ocspLastFailureUrl||x
ocspNonResponderCertEnabled||x
ocspOverrideUrl||x
ocspTimeout||x
revocationCheckEnabled||x



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been deprecated since 2.19. Replaced by clientCertAuthorities and domainCertAuthorities.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return AllApiApiGetCertAuthorityRequest
*/
func (a *AllApiService) GetCertAuthority(ctx _context.Context, certAuthorityName string) AllApiApiGetCertAuthorityRequest {
	return AllApiApiGetCertAuthorityRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return CertAuthorityResponse
 */
func (a *AllApiService) GetCertAuthorityExecute(r AllApiApiGetCertAuthorityRequest) (CertAuthorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CertAuthorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetCertAuthority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certAuthorities/{certAuthorityName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

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

type AllApiApiGetCertAuthorityOcspTlsTrustedCommonNameRequest struct {
	ctx                      _context.Context
	ApiService               *AllApiService
	certAuthorityName        string
	ocspTlsTrustedCommonName string
	select_                  *[]string
}

func (r AllApiApiGetCertAuthorityOcspTlsTrustedCommonNameRequest) Select_(select_ []string) AllApiApiGetCertAuthorityOcspTlsTrustedCommonNameRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetCertAuthorityOcspTlsTrustedCommonNameRequest) Execute() (CertAuthorityOcspTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	return r.ApiService.GetCertAuthorityOcspTlsTrustedCommonNameExecute(r)
}

/*
 * GetCertAuthorityOcspTlsTrustedCommonName Get an OCSP Responder Trusted Common Name object.
 * Get an OCSP Responder Trusted Common Name object.

When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|x
ocspTlsTrustedCommonName|x|x



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been deprecated since 2.19. Replaced by clientCertAuthorities.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @param ocspTlsTrustedCommonName The expected Trusted Common Name of the OCSP responder remote certificate.
 * @return AllApiApiGetCertAuthorityOcspTlsTrustedCommonNameRequest
*/
func (a *AllApiService) GetCertAuthorityOcspTlsTrustedCommonName(ctx _context.Context, certAuthorityName string, ocspTlsTrustedCommonName string) AllApiApiGetCertAuthorityOcspTlsTrustedCommonNameRequest {
	return AllApiApiGetCertAuthorityOcspTlsTrustedCommonNameRequest{
		ApiService:               a,
		ctx:                      ctx,
		certAuthorityName:        certAuthorityName,
		ocspTlsTrustedCommonName: ocspTlsTrustedCommonName,
	}
}

/*
 * Execute executes the request
 * @return CertAuthorityOcspTlsTrustedCommonNameResponse
 */
func (a *AllApiService) GetCertAuthorityOcspTlsTrustedCommonNameExecute(r AllApiApiGetCertAuthorityOcspTlsTrustedCommonNameRequest) (CertAuthorityOcspTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CertAuthorityOcspTlsTrustedCommonNameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetCertAuthorityOcspTlsTrustedCommonName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ocspTlsTrustedCommonName"+"}", _neturl.PathEscape(parameterToString(r.ocspTlsTrustedCommonName, "")), -1)

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

type AllApiApiGetCertAuthorityOcspTlsTrustedCommonNamesRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	certAuthorityName string
	where             *[]string
	select_           *[]string
}

func (r AllApiApiGetCertAuthorityOcspTlsTrustedCommonNamesRequest) Where(where []string) AllApiApiGetCertAuthorityOcspTlsTrustedCommonNamesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetCertAuthorityOcspTlsTrustedCommonNamesRequest) Select_(select_ []string) AllApiApiGetCertAuthorityOcspTlsTrustedCommonNamesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetCertAuthorityOcspTlsTrustedCommonNamesRequest) Execute() (CertAuthorityOcspTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetCertAuthorityOcspTlsTrustedCommonNamesExecute(r)
}

/*
 * GetCertAuthorityOcspTlsTrustedCommonNames Get a list of OCSP Responder Trusted Common Name objects.
 * Get a list of OCSP Responder Trusted Common Name objects.

When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|x
ocspTlsTrustedCommonName|x|x



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been deprecated since 2.19. Replaced by clientCertAuthorities.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return AllApiApiGetCertAuthorityOcspTlsTrustedCommonNamesRequest
*/
func (a *AllApiService) GetCertAuthorityOcspTlsTrustedCommonNames(ctx _context.Context, certAuthorityName string) AllApiApiGetCertAuthorityOcspTlsTrustedCommonNamesRequest {
	return AllApiApiGetCertAuthorityOcspTlsTrustedCommonNamesRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return CertAuthorityOcspTlsTrustedCommonNamesResponse
 */
func (a *AllApiService) GetCertAuthorityOcspTlsTrustedCommonNamesExecute(r AllApiApiGetCertAuthorityOcspTlsTrustedCommonNamesRequest) (CertAuthorityOcspTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CertAuthorityOcspTlsTrustedCommonNamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetCertAuthorityOcspTlsTrustedCommonNames")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

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

type AllApiApiGetClientCertAuthoritiesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetClientCertAuthoritiesRequest) Count(count int32) AllApiApiGetClientCertAuthoritiesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetClientCertAuthoritiesRequest) Cursor(cursor string) AllApiApiGetClientCertAuthoritiesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetClientCertAuthoritiesRequest) Where(where []string) AllApiApiGetClientCertAuthoritiesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetClientCertAuthoritiesRequest) Select_(select_ []string) AllApiApiGetClientCertAuthoritiesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetClientCertAuthoritiesRequest) Execute() (ClientCertAuthoritiesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetClientCertAuthoritiesExecute(r)
}

/*
 * GetClientCertAuthorities Get a list of Client Certificate Authority objects.
 * Get a list of Client Certificate Authority objects.

Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetClientCertAuthoritiesRequest
*/
func (a *AllApiService) GetClientCertAuthorities(ctx _context.Context) AllApiApiGetClientCertAuthoritiesRequest {
	return AllApiApiGetClientCertAuthoritiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthoritiesResponse
 */
func (a *AllApiService) GetClientCertAuthoritiesExecute(r AllApiApiGetClientCertAuthoritiesRequest) (ClientCertAuthoritiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthoritiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetClientCertAuthorities")
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

type AllApiApiGetClientCertAuthorityRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	certAuthorityName string
	select_           *[]string
}

func (r AllApiApiGetClientCertAuthorityRequest) Select_(select_ []string) AllApiApiGetClientCertAuthorityRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetClientCertAuthorityRequest) Execute() (ClientCertAuthorityResponse, *_nethttp.Response, error) {
	return r.ApiService.GetClientCertAuthorityExecute(r)
}

/*
 * GetClientCertAuthority Get a Client Certificate Authority object.
 * Get a Client Certificate Authority object.

Clients can authenticate with the message broker over TLS by presenting a valid client certificate. The message broker authenticates the client certificate by constructing a full certificate chain (from the client certificate to intermediate CAs to a configured root CA). The intermediate CAs in this chain can be provided by the client, or configured in the message broker. The root CA must be configured on the message broker.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return AllApiApiGetClientCertAuthorityRequest
*/
func (a *AllApiService) GetClientCertAuthority(ctx _context.Context, certAuthorityName string) AllApiApiGetClientCertAuthorityRequest {
	return AllApiApiGetClientCertAuthorityRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthorityResponse
 */
func (a *AllApiService) GetClientCertAuthorityExecute(r AllApiApiGetClientCertAuthorityRequest) (ClientCertAuthorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetClientCertAuthority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

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

type AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest struct {
	ctx                      _context.Context
	ApiService               *AllApiService
	certAuthorityName        string
	ocspTlsTrustedCommonName string
	select_                  *[]string
}

func (r AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest) Select_(select_ []string) AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest) Execute() (ClientCertAuthorityOcspTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	return r.ApiService.GetClientCertAuthorityOcspTlsTrustedCommonNameExecute(r)
}

/*
 * GetClientCertAuthorityOcspTlsTrustedCommonName Get an OCSP Responder Trusted Common Name object.
 * Get an OCSP Responder Trusted Common Name object.

When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|
ocspTlsTrustedCommonName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @param ocspTlsTrustedCommonName The expected Trusted Common Name of the OCSP responder remote certificate.
 * @return AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest
*/
func (a *AllApiService) GetClientCertAuthorityOcspTlsTrustedCommonName(ctx _context.Context, certAuthorityName string, ocspTlsTrustedCommonName string) AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest {
	return AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest{
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
func (a *AllApiService) GetClientCertAuthorityOcspTlsTrustedCommonNameExecute(r AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNameRequest) (ClientCertAuthorityOcspTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthorityOcspTlsTrustedCommonNameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetClientCertAuthorityOcspTlsTrustedCommonName")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ocspTlsTrustedCommonName"+"}", _neturl.PathEscape(parameterToString(r.ocspTlsTrustedCommonName, "")), -1)

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

type AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	certAuthorityName string
	where             *[]string
	select_           *[]string
}

func (r AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest) Where(where []string) AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest) Select_(select_ []string) AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest) Execute() (ClientCertAuthorityOcspTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetClientCertAuthorityOcspTlsTrustedCommonNamesExecute(r)
}

/*
 * GetClientCertAuthorityOcspTlsTrustedCommonNames Get a list of OCSP Responder Trusted Common Name objects.
 * Get a list of OCSP Responder Trusted Common Name objects.

When an OCSP override URL is configured, the OCSP responder will be required to sign the OCSP responses with certificates issued to these Trusted Common Names. A maximum of 8 common names can be configured as valid response signers.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|
ocspTlsTrustedCommonName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest
*/
func (a *AllApiService) GetClientCertAuthorityOcspTlsTrustedCommonNames(ctx _context.Context, certAuthorityName string) AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest {
	return AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return ClientCertAuthorityOcspTlsTrustedCommonNamesResponse
 */
func (a *AllApiService) GetClientCertAuthorityOcspTlsTrustedCommonNamesExecute(r AllApiApiGetClientCertAuthorityOcspTlsTrustedCommonNamesRequest) (ClientCertAuthorityOcspTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ClientCertAuthorityOcspTlsTrustedCommonNamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetClientCertAuthorityOcspTlsTrustedCommonNames")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clientCertAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

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

type AllApiApiGetDmrClusterRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	dmrClusterName string
	select_        *[]string
}

func (r AllApiApiGetDmrClusterRequest) Select_(select_ []string) AllApiApiGetDmrClusterRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterRequest) Execute() (DmrClusterResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterRequest
*/
func (a *AllApiService) GetDmrCluster(ctx _context.Context, dmrClusterName string) AllApiApiGetDmrClusterRequest {
	return AllApiApiGetDmrClusterRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterResponse
 */
func (a *AllApiService) GetDmrClusterExecute(r AllApiApiGetDmrClusterRequest) (DmrClusterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrCluster")
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

type AllApiApiGetDmrClusterLinkRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	dmrClusterName string
	remoteNodeName string
	select_        *[]string
}

func (r AllApiApiGetDmrClusterLinkRequest) Select_(select_ []string) AllApiApiGetDmrClusterLinkRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterLinkRequest) Execute() (DmrClusterLinkResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterLinkRequest
*/
func (a *AllApiService) GetDmrClusterLink(ctx _context.Context, dmrClusterName string, remoteNodeName string) AllApiApiGetDmrClusterLinkRequest {
	return AllApiApiGetDmrClusterLinkRequest{
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
func (a *AllApiService) GetDmrClusterLinkExecute(r AllApiApiGetDmrClusterLinkRequest) (DmrClusterLinkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusterLink")
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

type AllApiApiGetDmrClusterLinkChannelRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	dmrClusterName string
	remoteNodeName string
	msgVpnName     string
	select_        *[]string
}

func (r AllApiApiGetDmrClusterLinkChannelRequest) Select_(select_ []string) AllApiApiGetDmrClusterLinkChannelRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterLinkChannelRequest) Execute() (DmrClusterLinkChannelResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterLinkChannelRequest
*/
func (a *AllApiService) GetDmrClusterLinkChannel(ctx _context.Context, dmrClusterName string, remoteNodeName string, msgVpnName string) AllApiApiGetDmrClusterLinkChannelRequest {
	return AllApiApiGetDmrClusterLinkChannelRequest{
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
func (a *AllApiService) GetDmrClusterLinkChannelExecute(r AllApiApiGetDmrClusterLinkChannelRequest) (DmrClusterLinkChannelResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkChannelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusterLinkChannel")
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

type AllApiApiGetDmrClusterLinkChannelsRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	dmrClusterName string
	remoteNodeName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetDmrClusterLinkChannelsRequest) Count(count int32) AllApiApiGetDmrClusterLinkChannelsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetDmrClusterLinkChannelsRequest) Cursor(cursor string) AllApiApiGetDmrClusterLinkChannelsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetDmrClusterLinkChannelsRequest) Where(where []string) AllApiApiGetDmrClusterLinkChannelsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetDmrClusterLinkChannelsRequest) Select_(select_ []string) AllApiApiGetDmrClusterLinkChannelsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterLinkChannelsRequest) Execute() (DmrClusterLinkChannelsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterLinkChannelsRequest
*/
func (a *AllApiService) GetDmrClusterLinkChannels(ctx _context.Context, dmrClusterName string, remoteNodeName string) AllApiApiGetDmrClusterLinkChannelsRequest {
	return AllApiApiGetDmrClusterLinkChannelsRequest{
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
func (a *AllApiService) GetDmrClusterLinkChannelsExecute(r AllApiApiGetDmrClusterLinkChannelsRequest) (DmrClusterLinkChannelsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkChannelsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusterLinkChannels")
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

type AllApiApiGetDmrClusterLinkRemoteAddressRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	dmrClusterName string
	remoteNodeName string
	remoteAddress  string
	select_        *[]string
}

func (r AllApiApiGetDmrClusterLinkRemoteAddressRequest) Select_(select_ []string) AllApiApiGetDmrClusterLinkRemoteAddressRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterLinkRemoteAddressRequest) Execute() (DmrClusterLinkRemoteAddressResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterLinkRemoteAddressRequest
*/
func (a *AllApiService) GetDmrClusterLinkRemoteAddress(ctx _context.Context, dmrClusterName string, remoteNodeName string, remoteAddress string) AllApiApiGetDmrClusterLinkRemoteAddressRequest {
	return AllApiApiGetDmrClusterLinkRemoteAddressRequest{
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
func (a *AllApiService) GetDmrClusterLinkRemoteAddressExecute(r AllApiApiGetDmrClusterLinkRemoteAddressRequest) (DmrClusterLinkRemoteAddressResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkRemoteAddressResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusterLinkRemoteAddress")
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

type AllApiApiGetDmrClusterLinkRemoteAddressesRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	dmrClusterName string
	remoteNodeName string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetDmrClusterLinkRemoteAddressesRequest) Where(where []string) AllApiApiGetDmrClusterLinkRemoteAddressesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetDmrClusterLinkRemoteAddressesRequest) Select_(select_ []string) AllApiApiGetDmrClusterLinkRemoteAddressesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterLinkRemoteAddressesRequest) Execute() (DmrClusterLinkRemoteAddressesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterLinkRemoteAddressesRequest
*/
func (a *AllApiService) GetDmrClusterLinkRemoteAddresses(ctx _context.Context, dmrClusterName string, remoteNodeName string) AllApiApiGetDmrClusterLinkRemoteAddressesRequest {
	return AllApiApiGetDmrClusterLinkRemoteAddressesRequest{
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
func (a *AllApiService) GetDmrClusterLinkRemoteAddressesExecute(r AllApiApiGetDmrClusterLinkRemoteAddressesRequest) (DmrClusterLinkRemoteAddressesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkRemoteAddressesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusterLinkRemoteAddresses")
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

type AllApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest struct {
	ctx                  _context.Context
	ApiService           *AllApiService
	dmrClusterName       string
	remoteNodeName       string
	tlsTrustedCommonName string
	select_              *[]string
}

func (r AllApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest) Select_(select_ []string) AllApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest) Execute() (DmrClusterLinkTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest
*/
func (a *AllApiService) GetDmrClusterLinkTlsTrustedCommonName(ctx _context.Context, dmrClusterName string, remoteNodeName string, tlsTrustedCommonName string) AllApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest {
	return AllApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest{
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
func (a *AllApiService) GetDmrClusterLinkTlsTrustedCommonNameExecute(r AllApiApiGetDmrClusterLinkTlsTrustedCommonNameRequest) (DmrClusterLinkTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkTlsTrustedCommonNameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusterLinkTlsTrustedCommonName")
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

type AllApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	dmrClusterName string
	remoteNodeName string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest) Where(where []string) AllApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest) Select_(select_ []string) AllApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest) Execute() (DmrClusterLinkTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest
*/
func (a *AllApiService) GetDmrClusterLinkTlsTrustedCommonNames(ctx _context.Context, dmrClusterName string, remoteNodeName string) AllApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest {
	return AllApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest{
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
func (a *AllApiService) GetDmrClusterLinkTlsTrustedCommonNamesExecute(r AllApiApiGetDmrClusterLinkTlsTrustedCommonNamesRequest) (DmrClusterLinkTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinkTlsTrustedCommonNamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusterLinkTlsTrustedCommonNames")
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

type AllApiApiGetDmrClusterLinksRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	dmrClusterName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetDmrClusterLinksRequest) Count(count int32) AllApiApiGetDmrClusterLinksRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetDmrClusterLinksRequest) Cursor(cursor string) AllApiApiGetDmrClusterLinksRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetDmrClusterLinksRequest) Where(where []string) AllApiApiGetDmrClusterLinksRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetDmrClusterLinksRequest) Select_(select_ []string) AllApiApiGetDmrClusterLinksRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterLinksRequest) Execute() (DmrClusterLinksResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterLinksRequest
*/
func (a *AllApiService) GetDmrClusterLinks(ctx _context.Context, dmrClusterName string) AllApiApiGetDmrClusterLinksRequest {
	return AllApiApiGetDmrClusterLinksRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterLinksResponse
 */
func (a *AllApiService) GetDmrClusterLinksExecute(r AllApiApiGetDmrClusterLinksRequest) (DmrClusterLinksResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterLinksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusterLinks")
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

type AllApiApiGetDmrClusterTopologyIssueRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	dmrClusterName string
	topologyIssue  string
	select_        *[]string
}

func (r AllApiApiGetDmrClusterTopologyIssueRequest) Select_(select_ []string) AllApiApiGetDmrClusterTopologyIssueRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterTopologyIssueRequest) Execute() (DmrClusterTopologyIssueResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterTopologyIssueRequest
*/
func (a *AllApiService) GetDmrClusterTopologyIssue(ctx _context.Context, dmrClusterName string, topologyIssue string) AllApiApiGetDmrClusterTopologyIssueRequest {
	return AllApiApiGetDmrClusterTopologyIssueRequest{
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
func (a *AllApiService) GetDmrClusterTopologyIssueExecute(r AllApiApiGetDmrClusterTopologyIssueRequest) (DmrClusterTopologyIssueResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterTopologyIssueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusterTopologyIssue")
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

type AllApiApiGetDmrClusterTopologyIssuesRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	dmrClusterName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetDmrClusterTopologyIssuesRequest) Count(count int32) AllApiApiGetDmrClusterTopologyIssuesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetDmrClusterTopologyIssuesRequest) Cursor(cursor string) AllApiApiGetDmrClusterTopologyIssuesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetDmrClusterTopologyIssuesRequest) Where(where []string) AllApiApiGetDmrClusterTopologyIssuesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetDmrClusterTopologyIssuesRequest) Select_(select_ []string) AllApiApiGetDmrClusterTopologyIssuesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClusterTopologyIssuesRequest) Execute() (DmrClusterTopologyIssuesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClusterTopologyIssuesRequest
*/
func (a *AllApiService) GetDmrClusterTopologyIssues(ctx _context.Context, dmrClusterName string) AllApiApiGetDmrClusterTopologyIssuesRequest {
	return AllApiApiGetDmrClusterTopologyIssuesRequest{
		ApiService:     a,
		ctx:            ctx,
		dmrClusterName: dmrClusterName,
	}
}

/*
 * Execute executes the request
 * @return DmrClusterTopologyIssuesResponse
 */
func (a *AllApiService) GetDmrClusterTopologyIssuesExecute(r AllApiApiGetDmrClusterTopologyIssuesRequest) (DmrClusterTopologyIssuesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClusterTopologyIssuesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusterTopologyIssues")
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

type AllApiApiGetDmrClustersRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetDmrClustersRequest) Count(count int32) AllApiApiGetDmrClustersRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetDmrClustersRequest) Cursor(cursor string) AllApiApiGetDmrClustersRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetDmrClustersRequest) Where(where []string) AllApiApiGetDmrClustersRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetDmrClustersRequest) Select_(select_ []string) AllApiApiGetDmrClustersRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDmrClustersRequest) Execute() (DmrClustersResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetDmrClustersRequest
*/
func (a *AllApiService) GetDmrClusters(ctx _context.Context) AllApiApiGetDmrClustersRequest {
	return AllApiApiGetDmrClustersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return DmrClustersResponse
 */
func (a *AllApiService) GetDmrClustersExecute(r AllApiApiGetDmrClustersRequest) (DmrClustersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DmrClustersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDmrClusters")
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

type AllApiApiGetDomainCertAuthoritiesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetDomainCertAuthoritiesRequest) Count(count int32) AllApiApiGetDomainCertAuthoritiesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetDomainCertAuthoritiesRequest) Cursor(cursor string) AllApiApiGetDomainCertAuthoritiesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetDomainCertAuthoritiesRequest) Where(where []string) AllApiApiGetDomainCertAuthoritiesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetDomainCertAuthoritiesRequest) Select_(select_ []string) AllApiApiGetDomainCertAuthoritiesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDomainCertAuthoritiesRequest) Execute() (DomainCertAuthoritiesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDomainCertAuthoritiesExecute(r)
}

/*
 * GetDomainCertAuthorities Get a list of Domain Certificate Authority objects.
 * Get a list of Domain Certificate Authority objects.

Certificate Authorities trusted for domain verification.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetDomainCertAuthoritiesRequest
*/
func (a *AllApiService) GetDomainCertAuthorities(ctx _context.Context) AllApiApiGetDomainCertAuthoritiesRequest {
	return AllApiApiGetDomainCertAuthoritiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return DomainCertAuthoritiesResponse
 */
func (a *AllApiService) GetDomainCertAuthoritiesExecute(r AllApiApiGetDomainCertAuthoritiesRequest) (DomainCertAuthoritiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DomainCertAuthoritiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDomainCertAuthorities")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domainCertAuthorities"

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

type AllApiApiGetDomainCertAuthorityRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	certAuthorityName string
	select_           *[]string
}

func (r AllApiApiGetDomainCertAuthorityRequest) Select_(select_ []string) AllApiApiGetDomainCertAuthorityRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetDomainCertAuthorityRequest) Execute() (DomainCertAuthorityResponse, *_nethttp.Response, error) {
	return r.ApiService.GetDomainCertAuthorityExecute(r)
}

/*
 * GetDomainCertAuthority Get a Domain Certificate Authority object.
 * Get a Domain Certificate Authority object.

Certificate Authorities trusted for domain verification.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return AllApiApiGetDomainCertAuthorityRequest
*/
func (a *AllApiService) GetDomainCertAuthority(ctx _context.Context, certAuthorityName string) AllApiApiGetDomainCertAuthorityRequest {
	return AllApiApiGetDomainCertAuthorityRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return DomainCertAuthorityResponse
 */
func (a *AllApiService) GetDomainCertAuthorityExecute(r AllApiApiGetDomainCertAuthorityRequest) (DomainCertAuthorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DomainCertAuthorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetDomainCertAuthority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/domainCertAuthorities/{certAuthorityName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

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

type AllApiApiGetMsgVpnRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnRequest) Select_(select_ []string) AllApiApiGetMsgVpnRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRequest) Execute() (MsgVpnResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRequest
*/
func (a *AllApiService) GetMsgVpn(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnRequest {
	return AllApiApiGetMsgVpnRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnResponse
 */
func (a *AllApiService) GetMsgVpnExecute(r AllApiApiGetMsgVpnRequest) (MsgVpnResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpn")
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

type AllApiApiGetMsgVpnAclProfileRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	msgVpnName     string
	aclProfileName string
	select_        *[]string
}

func (r AllApiApiGetMsgVpnAclProfileRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfileRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfileRequest) Execute() (MsgVpnAclProfileResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfileRequest
*/
func (a *AllApiService) GetMsgVpnAclProfile(ctx _context.Context, msgVpnName string, aclProfileName string) AllApiApiGetMsgVpnAclProfileRequest {
	return AllApiApiGetMsgVpnAclProfileRequest{
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
func (a *AllApiService) GetMsgVpnAclProfileExecute(r AllApiApiGetMsgVpnAclProfileRequest) (MsgVpnAclProfileResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfile")
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

type AllApiApiGetMsgVpnAclProfileClientConnectExceptionRequest struct {
	ctx                           _context.Context
	ApiService                    *AllApiService
	msgVpnName                    string
	aclProfileName                string
	clientConnectExceptionAddress string
	select_                       *[]string
}

func (r AllApiApiGetMsgVpnAclProfileClientConnectExceptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfileClientConnectExceptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfileClientConnectExceptionRequest) Execute() (MsgVpnAclProfileClientConnectExceptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfileClientConnectExceptionRequest
*/
func (a *AllApiService) GetMsgVpnAclProfileClientConnectException(ctx _context.Context, msgVpnName string, aclProfileName string, clientConnectExceptionAddress string) AllApiApiGetMsgVpnAclProfileClientConnectExceptionRequest {
	return AllApiApiGetMsgVpnAclProfileClientConnectExceptionRequest{
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
func (a *AllApiService) GetMsgVpnAclProfileClientConnectExceptionExecute(r AllApiApiGetMsgVpnAclProfileClientConnectExceptionRequest) (MsgVpnAclProfileClientConnectExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileClientConnectExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfileClientConnectException")
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

type AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) Count(count int32) AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) Where(where []string) AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) Execute() (MsgVpnAclProfileClientConnectExceptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest
*/
func (a *AllApiService) GetMsgVpnAclProfileClientConnectExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest {
	return AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest{
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
func (a *AllApiService) GetMsgVpnAclProfileClientConnectExceptionsExecute(r AllApiApiGetMsgVpnAclProfileClientConnectExceptionsRequest) (MsgVpnAclProfileClientConnectExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileClientConnectExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfileClientConnectExceptions")
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

type AllApiApiGetMsgVpnAclProfilePublishExceptionRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	aclProfileName        string
	topicSyntax           string
	publishExceptionTopic string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnAclProfilePublishExceptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfilePublishExceptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfilePublishExceptionRequest) Execute() (MsgVpnAclProfilePublishExceptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfilePublishExceptionRequest
*/
func (a *AllApiService) GetMsgVpnAclProfilePublishException(ctx _context.Context, msgVpnName string, aclProfileName string, topicSyntax string, publishExceptionTopic string) AllApiApiGetMsgVpnAclProfilePublishExceptionRequest {
	return AllApiApiGetMsgVpnAclProfilePublishExceptionRequest{
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
func (a *AllApiService) GetMsgVpnAclProfilePublishExceptionExecute(r AllApiApiGetMsgVpnAclProfilePublishExceptionRequest) (MsgVpnAclProfilePublishExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfilePublishExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfilePublishException")
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

type AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest) Count(count int32) AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest) Where(where []string) AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest) Execute() (MsgVpnAclProfilePublishExceptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest
*/
func (a *AllApiService) GetMsgVpnAclProfilePublishExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest {
	return AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest{
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
func (a *AllApiService) GetMsgVpnAclProfilePublishExceptionsExecute(r AllApiApiGetMsgVpnAclProfilePublishExceptionsRequest) (MsgVpnAclProfilePublishExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfilePublishExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfilePublishExceptions")
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

type AllApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest struct {
	ctx                         _context.Context
	ApiService                  *AllApiService
	msgVpnName                  string
	aclProfileName              string
	publishTopicExceptionSyntax string
	publishTopicException       string
	select_                     *[]string
}

func (r AllApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest) Execute() (MsgVpnAclProfilePublishTopicExceptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest
*/
func (a *AllApiService) GetMsgVpnAclProfilePublishTopicException(ctx _context.Context, msgVpnName string, aclProfileName string, publishTopicExceptionSyntax string, publishTopicException string) AllApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest {
	return AllApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest{
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
func (a *AllApiService) GetMsgVpnAclProfilePublishTopicExceptionExecute(r AllApiApiGetMsgVpnAclProfilePublishTopicExceptionRequest) (MsgVpnAclProfilePublishTopicExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfilePublishTopicExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfilePublishTopicException")
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

type AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) Count(count int32) AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) Where(where []string) AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) Execute() (MsgVpnAclProfilePublishTopicExceptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest
*/
func (a *AllApiService) GetMsgVpnAclProfilePublishTopicExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest {
	return AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest{
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
func (a *AllApiService) GetMsgVpnAclProfilePublishTopicExceptionsExecute(r AllApiApiGetMsgVpnAclProfilePublishTopicExceptionsRequest) (MsgVpnAclProfilePublishTopicExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfilePublishTopicExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfilePublishTopicExceptions")
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

type AllApiApiGetMsgVpnAclProfileSubscribeExceptionRequest struct {
	ctx                     _context.Context
	ApiService              *AllApiService
	msgVpnName              string
	aclProfileName          string
	topicSyntax             string
	subscribeExceptionTopic string
	select_                 *[]string
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeExceptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfileSubscribeExceptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeExceptionRequest) Execute() (MsgVpnAclProfileSubscribeExceptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfileSubscribeExceptionRequest
*/
func (a *AllApiService) GetMsgVpnAclProfileSubscribeException(ctx _context.Context, msgVpnName string, aclProfileName string, topicSyntax string, subscribeExceptionTopic string) AllApiApiGetMsgVpnAclProfileSubscribeExceptionRequest {
	return AllApiApiGetMsgVpnAclProfileSubscribeExceptionRequest{
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
func (a *AllApiService) GetMsgVpnAclProfileSubscribeExceptionExecute(r AllApiApiGetMsgVpnAclProfileSubscribeExceptionRequest) (MsgVpnAclProfileSubscribeExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfileSubscribeException")
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

type AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) Count(count int32) AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) Where(where []string) AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) Execute() (MsgVpnAclProfileSubscribeExceptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest
*/
func (a *AllApiService) GetMsgVpnAclProfileSubscribeExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest {
	return AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest{
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
func (a *AllApiService) GetMsgVpnAclProfileSubscribeExceptionsExecute(r AllApiApiGetMsgVpnAclProfileSubscribeExceptionsRequest) (MsgVpnAclProfileSubscribeExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfileSubscribeExceptions")
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

type AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest struct {
	ctx                               _context.Context
	ApiService                        *AllApiService
	msgVpnName                        string
	aclProfileName                    string
	subscribeShareNameExceptionSyntax string
	subscribeShareNameException       string
	select_                           *[]string
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest) Execute() (MsgVpnAclProfileSubscribeShareNameExceptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest
*/
func (a *AllApiService) GetMsgVpnAclProfileSubscribeShareNameException(ctx _context.Context, msgVpnName string, aclProfileName string, subscribeShareNameExceptionSyntax string, subscribeShareNameException string) AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest {
	return AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest{
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
func (a *AllApiService) GetMsgVpnAclProfileSubscribeShareNameExceptionExecute(r AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest) (MsgVpnAclProfileSubscribeShareNameExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeShareNameExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfileSubscribeShareNameException")
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

type AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) Count(count int32) AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) Where(where []string) AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) Execute() (MsgVpnAclProfileSubscribeShareNameExceptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest
*/
func (a *AllApiService) GetMsgVpnAclProfileSubscribeShareNameExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest {
	return AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest{
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
func (a *AllApiService) GetMsgVpnAclProfileSubscribeShareNameExceptionsExecute(r AllApiApiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest) (MsgVpnAclProfileSubscribeShareNameExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeShareNameExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfileSubscribeShareNameExceptions")
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

type AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest struct {
	ctx                           _context.Context
	ApiService                    *AllApiService
	msgVpnName                    string
	aclProfileName                string
	subscribeTopicExceptionSyntax string
	subscribeTopicException       string
	select_                       *[]string
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest) Execute() (MsgVpnAclProfileSubscribeTopicExceptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest
*/
func (a *AllApiService) GetMsgVpnAclProfileSubscribeTopicException(ctx _context.Context, msgVpnName string, aclProfileName string, subscribeTopicExceptionSyntax string, subscribeTopicException string) AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest {
	return AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest{
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
func (a *AllApiService) GetMsgVpnAclProfileSubscribeTopicExceptionExecute(r AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionRequest) (MsgVpnAclProfileSubscribeTopicExceptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeTopicExceptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfileSubscribeTopicException")
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

type AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	msgVpnName     string
	aclProfileName string
	count          *int32
	cursor         *string
	where          *[]string
	select_        *[]string
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) Count(count int32) AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) Where(where []string) AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) Execute() (MsgVpnAclProfileSubscribeTopicExceptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest
*/
func (a *AllApiService) GetMsgVpnAclProfileSubscribeTopicExceptions(ctx _context.Context, msgVpnName string, aclProfileName string) AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest {
	return AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest{
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
func (a *AllApiService) GetMsgVpnAclProfileSubscribeTopicExceptionsExecute(r AllApiApiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest) (MsgVpnAclProfileSubscribeTopicExceptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfileSubscribeTopicExceptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfileSubscribeTopicExceptions")
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

type AllApiApiGetMsgVpnAclProfilesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnAclProfilesRequest) Count(count int32) AllApiApiGetMsgVpnAclProfilesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnAclProfilesRequest) Cursor(cursor string) AllApiApiGetMsgVpnAclProfilesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnAclProfilesRequest) Where(where []string) AllApiApiGetMsgVpnAclProfilesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnAclProfilesRequest) Select_(select_ []string) AllApiApiGetMsgVpnAclProfilesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAclProfilesRequest) Execute() (MsgVpnAclProfilesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAclProfilesRequest
*/
func (a *AllApiService) GetMsgVpnAclProfiles(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnAclProfilesRequest {
	return AllApiApiGetMsgVpnAclProfilesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAclProfilesResponse
 */
func (a *AllApiService) GetMsgVpnAclProfilesExecute(r AllApiApiGetMsgVpnAclProfilesRequest) (MsgVpnAclProfilesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAclProfilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAclProfiles")
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

type AllApiApiGetMsgVpnAuthenticationOauthProviderRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	oauthProviderName string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnAuthenticationOauthProviderRequest) Select_(select_ []string) AllApiApiGetMsgVpnAuthenticationOauthProviderRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAuthenticationOauthProviderRequest) Execute() (MsgVpnAuthenticationOauthProviderResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAuthenticationOauthProviderRequest
*/
func (a *AllApiService) GetMsgVpnAuthenticationOauthProvider(ctx _context.Context, msgVpnName string, oauthProviderName string) AllApiApiGetMsgVpnAuthenticationOauthProviderRequest {
	return AllApiApiGetMsgVpnAuthenticationOauthProviderRequest{
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
func (a *AllApiService) GetMsgVpnAuthenticationOauthProviderExecute(r AllApiApiGetMsgVpnAuthenticationOauthProviderRequest) (MsgVpnAuthenticationOauthProviderResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAuthenticationOauthProviderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAuthenticationOauthProvider")
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

type AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest) Count(count int32) AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest) Cursor(cursor string) AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest) Where(where []string) AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest) Select_(select_ []string) AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest) Execute() (MsgVpnAuthenticationOauthProvidersResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest
*/
func (a *AllApiService) GetMsgVpnAuthenticationOauthProviders(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest {
	return AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAuthenticationOauthProvidersResponse
 */
func (a *AllApiService) GetMsgVpnAuthenticationOauthProvidersExecute(r AllApiApiGetMsgVpnAuthenticationOauthProvidersRequest) (MsgVpnAuthenticationOauthProvidersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAuthenticationOauthProvidersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAuthenticationOauthProviders")
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

type AllApiApiGetMsgVpnAuthorizationGroupRequest struct {
	ctx                    _context.Context
	ApiService             *AllApiService
	msgVpnName             string
	authorizationGroupName string
	select_                *[]string
}

func (r AllApiApiGetMsgVpnAuthorizationGroupRequest) Select_(select_ []string) AllApiApiGetMsgVpnAuthorizationGroupRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAuthorizationGroupRequest) Execute() (MsgVpnAuthorizationGroupResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAuthorizationGroupRequest
*/
func (a *AllApiService) GetMsgVpnAuthorizationGroup(ctx _context.Context, msgVpnName string, authorizationGroupName string) AllApiApiGetMsgVpnAuthorizationGroupRequest {
	return AllApiApiGetMsgVpnAuthorizationGroupRequest{
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
func (a *AllApiService) GetMsgVpnAuthorizationGroupExecute(r AllApiApiGetMsgVpnAuthorizationGroupRequest) (MsgVpnAuthorizationGroupResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAuthorizationGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAuthorizationGroup")
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

type AllApiApiGetMsgVpnAuthorizationGroupsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnAuthorizationGroupsRequest) Count(count int32) AllApiApiGetMsgVpnAuthorizationGroupsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnAuthorizationGroupsRequest) Cursor(cursor string) AllApiApiGetMsgVpnAuthorizationGroupsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnAuthorizationGroupsRequest) Where(where []string) AllApiApiGetMsgVpnAuthorizationGroupsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnAuthorizationGroupsRequest) Select_(select_ []string) AllApiApiGetMsgVpnAuthorizationGroupsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnAuthorizationGroupsRequest) Execute() (MsgVpnAuthorizationGroupsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnAuthorizationGroupsRequest
*/
func (a *AllApiService) GetMsgVpnAuthorizationGroups(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnAuthorizationGroupsRequest {
	return AllApiApiGetMsgVpnAuthorizationGroupsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnAuthorizationGroupsResponse
 */
func (a *AllApiService) GetMsgVpnAuthorizationGroupsExecute(r AllApiApiGetMsgVpnAuthorizationGroupsRequest) (MsgVpnAuthorizationGroupsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnAuthorizationGroupsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnAuthorizationGroups")
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

type AllApiApiGetMsgVpnBridgeRequest struct {
	ctx                 _context.Context
	ApiService          *AllApiService
	msgVpnName          string
	bridgeName          string
	bridgeVirtualRouter string
	select_             *[]string
}

func (r AllApiApiGetMsgVpnBridgeRequest) Select_(select_ []string) AllApiApiGetMsgVpnBridgeRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnBridgeRequest) Execute() (MsgVpnBridgeResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnBridgeRequest
*/
func (a *AllApiService) GetMsgVpnBridge(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string) AllApiApiGetMsgVpnBridgeRequest {
	return AllApiApiGetMsgVpnBridgeRequest{
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
func (a *AllApiService) GetMsgVpnBridgeExecute(r AllApiApiGetMsgVpnBridgeRequest) (MsgVpnBridgeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnBridge")
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

type AllApiApiGetMsgVpnBridgeLocalSubscriptionRequest struct {
	ctx                    _context.Context
	ApiService             *AllApiService
	msgVpnName             string
	bridgeName             string
	bridgeVirtualRouter    string
	localSubscriptionTopic string
	select_                *[]string
}

func (r AllApiApiGetMsgVpnBridgeLocalSubscriptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnBridgeLocalSubscriptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnBridgeLocalSubscriptionRequest) Execute() (MsgVpnBridgeLocalSubscriptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnBridgeLocalSubscriptionRequest
*/
func (a *AllApiService) GetMsgVpnBridgeLocalSubscription(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string, localSubscriptionTopic string) AllApiApiGetMsgVpnBridgeLocalSubscriptionRequest {
	return AllApiApiGetMsgVpnBridgeLocalSubscriptionRequest{
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
func (a *AllApiService) GetMsgVpnBridgeLocalSubscriptionExecute(r AllApiApiGetMsgVpnBridgeLocalSubscriptionRequest) (MsgVpnBridgeLocalSubscriptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeLocalSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnBridgeLocalSubscription")
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

type AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest struct {
	ctx                 _context.Context
	ApiService          *AllApiService
	msgVpnName          string
	bridgeName          string
	bridgeVirtualRouter string
	count               *int32
	cursor              *string
	where               *[]string
	select_             *[]string
}

func (r AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) Count(count int32) AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) Where(where []string) AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) Execute() (MsgVpnBridgeLocalSubscriptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest
*/
func (a *AllApiService) GetMsgVpnBridgeLocalSubscriptions(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string) AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest {
	return AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest{
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
func (a *AllApiService) GetMsgVpnBridgeLocalSubscriptionsExecute(r AllApiApiGetMsgVpnBridgeLocalSubscriptionsRequest) (MsgVpnBridgeLocalSubscriptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeLocalSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnBridgeLocalSubscriptions")
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

type AllApiApiGetMsgVpnBridgeRemoteMsgVpnRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	bridgeName            string
	bridgeVirtualRouter   string
	remoteMsgVpnName      string
	remoteMsgVpnLocation  string
	remoteMsgVpnInterface string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnBridgeRemoteMsgVpnRequest) Select_(select_ []string) AllApiApiGetMsgVpnBridgeRemoteMsgVpnRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnBridgeRemoteMsgVpnRequest) Execute() (MsgVpnBridgeRemoteMsgVpnResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnBridgeRemoteMsgVpnRequest
*/
func (a *AllApiService) GetMsgVpnBridgeRemoteMsgVpn(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string, remoteMsgVpnName string, remoteMsgVpnLocation string, remoteMsgVpnInterface string) AllApiApiGetMsgVpnBridgeRemoteMsgVpnRequest {
	return AllApiApiGetMsgVpnBridgeRemoteMsgVpnRequest{
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
func (a *AllApiService) GetMsgVpnBridgeRemoteMsgVpnExecute(r AllApiApiGetMsgVpnBridgeRemoteMsgVpnRequest) (MsgVpnBridgeRemoteMsgVpnResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeRemoteMsgVpnResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnBridgeRemoteMsgVpn")
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

type AllApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest struct {
	ctx                 _context.Context
	ApiService          *AllApiService
	msgVpnName          string
	bridgeName          string
	bridgeVirtualRouter string
	where               *[]string
	select_             *[]string
}

func (r AllApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest) Where(where []string) AllApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest) Select_(select_ []string) AllApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest) Execute() (MsgVpnBridgeRemoteMsgVpnsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest
*/
func (a *AllApiService) GetMsgVpnBridgeRemoteMsgVpns(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string) AllApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest {
	return AllApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest{
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
func (a *AllApiService) GetMsgVpnBridgeRemoteMsgVpnsExecute(r AllApiApiGetMsgVpnBridgeRemoteMsgVpnsRequest) (MsgVpnBridgeRemoteMsgVpnsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeRemoteMsgVpnsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnBridgeRemoteMsgVpns")
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

type AllApiApiGetMsgVpnBridgeRemoteSubscriptionRequest struct {
	ctx                     _context.Context
	ApiService              *AllApiService
	msgVpnName              string
	bridgeName              string
	bridgeVirtualRouter     string
	remoteSubscriptionTopic string
	select_                 *[]string
}

func (r AllApiApiGetMsgVpnBridgeRemoteSubscriptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnBridgeRemoteSubscriptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnBridgeRemoteSubscriptionRequest) Execute() (MsgVpnBridgeRemoteSubscriptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnBridgeRemoteSubscriptionRequest
*/
func (a *AllApiService) GetMsgVpnBridgeRemoteSubscription(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string, remoteSubscriptionTopic string) AllApiApiGetMsgVpnBridgeRemoteSubscriptionRequest {
	return AllApiApiGetMsgVpnBridgeRemoteSubscriptionRequest{
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
func (a *AllApiService) GetMsgVpnBridgeRemoteSubscriptionExecute(r AllApiApiGetMsgVpnBridgeRemoteSubscriptionRequest) (MsgVpnBridgeRemoteSubscriptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeRemoteSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnBridgeRemoteSubscription")
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

type AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest struct {
	ctx                 _context.Context
	ApiService          *AllApiService
	msgVpnName          string
	bridgeName          string
	bridgeVirtualRouter string
	count               *int32
	cursor              *string
	where               *[]string
	select_             *[]string
}

func (r AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) Count(count int32) AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) Where(where []string) AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) Execute() (MsgVpnBridgeRemoteSubscriptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest
*/
func (a *AllApiService) GetMsgVpnBridgeRemoteSubscriptions(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string) AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest {
	return AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest{
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
func (a *AllApiService) GetMsgVpnBridgeRemoteSubscriptionsExecute(r AllApiApiGetMsgVpnBridgeRemoteSubscriptionsRequest) (MsgVpnBridgeRemoteSubscriptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeRemoteSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnBridgeRemoteSubscriptions")
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

type AllApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest struct {
	ctx                  _context.Context
	ApiService           *AllApiService
	msgVpnName           string
	bridgeName           string
	bridgeVirtualRouter  string
	tlsTrustedCommonName string
	select_              *[]string
}

func (r AllApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest) Select_(select_ []string) AllApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest) Execute() (MsgVpnBridgeTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest
*/
func (a *AllApiService) GetMsgVpnBridgeTlsTrustedCommonName(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string, tlsTrustedCommonName string) AllApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest {
	return AllApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest{
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
func (a *AllApiService) GetMsgVpnBridgeTlsTrustedCommonNameExecute(r AllApiApiGetMsgVpnBridgeTlsTrustedCommonNameRequest) (MsgVpnBridgeTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeTlsTrustedCommonNameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnBridgeTlsTrustedCommonName")
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

type AllApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest struct {
	ctx                 _context.Context
	ApiService          *AllApiService
	msgVpnName          string
	bridgeName          string
	bridgeVirtualRouter string
	where               *[]string
	select_             *[]string
}

func (r AllApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest) Where(where []string) AllApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest) Select_(select_ []string) AllApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest) Execute() (MsgVpnBridgeTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest
*/
func (a *AllApiService) GetMsgVpnBridgeTlsTrustedCommonNames(ctx _context.Context, msgVpnName string, bridgeName string, bridgeVirtualRouter string) AllApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest {
	return AllApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest{
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
func (a *AllApiService) GetMsgVpnBridgeTlsTrustedCommonNamesExecute(r AllApiApiGetMsgVpnBridgeTlsTrustedCommonNamesRequest) (MsgVpnBridgeTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgeTlsTrustedCommonNamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnBridgeTlsTrustedCommonNames")
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

type AllApiApiGetMsgVpnBridgesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnBridgesRequest) Count(count int32) AllApiApiGetMsgVpnBridgesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnBridgesRequest) Cursor(cursor string) AllApiApiGetMsgVpnBridgesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnBridgesRequest) Where(where []string) AllApiApiGetMsgVpnBridgesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnBridgesRequest) Select_(select_ []string) AllApiApiGetMsgVpnBridgesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnBridgesRequest) Execute() (MsgVpnBridgesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnBridgesRequest
*/
func (a *AllApiService) GetMsgVpnBridges(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnBridgesRequest {
	return AllApiApiGetMsgVpnBridgesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnBridgesResponse
 */
func (a *AllApiService) GetMsgVpnBridgesExecute(r AllApiApiGetMsgVpnBridgesRequest) (MsgVpnBridgesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnBridgesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnBridges")
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

type AllApiApiGetMsgVpnClientRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	clientName string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientRequest) Execute() (MsgVpnClientResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientRequest
*/
func (a *AllApiService) GetMsgVpnClient(ctx _context.Context, msgVpnName string, clientName string) AllApiApiGetMsgVpnClientRequest {
	return AllApiApiGetMsgVpnClientRequest{
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
func (a *AllApiService) GetMsgVpnClientExecute(r AllApiApiGetMsgVpnClientRequest) (MsgVpnClientResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClient")
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

type AllApiApiGetMsgVpnClientConnectionRequest struct {
	ctx           _context.Context
	ApiService    *AllApiService
	msgVpnName    string
	clientName    string
	clientAddress string
	select_       *[]string
}

func (r AllApiApiGetMsgVpnClientConnectionRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientConnectionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientConnectionRequest) Execute() (MsgVpnClientConnectionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientConnectionRequest
*/
func (a *AllApiService) GetMsgVpnClientConnection(ctx _context.Context, msgVpnName string, clientName string, clientAddress string) AllApiApiGetMsgVpnClientConnectionRequest {
	return AllApiApiGetMsgVpnClientConnectionRequest{
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
func (a *AllApiService) GetMsgVpnClientConnectionExecute(r AllApiApiGetMsgVpnClientConnectionRequest) (MsgVpnClientConnectionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientConnectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientConnection")
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

type AllApiApiGetMsgVpnClientConnectionsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	clientName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientConnectionsRequest) Count(count int32) AllApiApiGetMsgVpnClientConnectionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnClientConnectionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnClientConnectionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnClientConnectionsRequest) Where(where []string) AllApiApiGetMsgVpnClientConnectionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnClientConnectionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientConnectionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientConnectionsRequest) Execute() (MsgVpnClientConnectionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientConnectionsRequest
*/
func (a *AllApiService) GetMsgVpnClientConnections(ctx _context.Context, msgVpnName string, clientName string) AllApiApiGetMsgVpnClientConnectionsRequest {
	return AllApiApiGetMsgVpnClientConnectionsRequest{
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
func (a *AllApiService) GetMsgVpnClientConnectionsExecute(r AllApiApiGetMsgVpnClientConnectionsRequest) (MsgVpnClientConnectionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientConnectionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientConnections")
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

type AllApiApiGetMsgVpnClientProfileRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	clientProfileName string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnClientProfileRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientProfileRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientProfileRequest) Execute() (MsgVpnClientProfileResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientProfileRequest
*/
func (a *AllApiService) GetMsgVpnClientProfile(ctx _context.Context, msgVpnName string, clientProfileName string) AllApiApiGetMsgVpnClientProfileRequest {
	return AllApiApiGetMsgVpnClientProfileRequest{
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
func (a *AllApiService) GetMsgVpnClientProfileExecute(r AllApiApiGetMsgVpnClientProfileRequest) (MsgVpnClientProfileResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientProfileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientProfile")
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

type AllApiApiGetMsgVpnClientProfilesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientProfilesRequest) Count(count int32) AllApiApiGetMsgVpnClientProfilesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnClientProfilesRequest) Cursor(cursor string) AllApiApiGetMsgVpnClientProfilesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnClientProfilesRequest) Where(where []string) AllApiApiGetMsgVpnClientProfilesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnClientProfilesRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientProfilesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientProfilesRequest) Execute() (MsgVpnClientProfilesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientProfilesRequest
*/
func (a *AllApiService) GetMsgVpnClientProfiles(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnClientProfilesRequest {
	return AllApiApiGetMsgVpnClientProfilesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientProfilesResponse
 */
func (a *AllApiService) GetMsgVpnClientProfilesExecute(r AllApiApiGetMsgVpnClientProfilesRequest) (MsgVpnClientProfilesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientProfilesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientProfiles")
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

type AllApiApiGetMsgVpnClientRxFlowRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	clientName string
	flowId     string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientRxFlowRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientRxFlowRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientRxFlowRequest) Execute() (MsgVpnClientRxFlowResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientRxFlowRequest
*/
func (a *AllApiService) GetMsgVpnClientRxFlow(ctx _context.Context, msgVpnName string, clientName string, flowId string) AllApiApiGetMsgVpnClientRxFlowRequest {
	return AllApiApiGetMsgVpnClientRxFlowRequest{
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
func (a *AllApiService) GetMsgVpnClientRxFlowExecute(r AllApiApiGetMsgVpnClientRxFlowRequest) (MsgVpnClientRxFlowResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientRxFlowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientRxFlow")
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

type AllApiApiGetMsgVpnClientRxFlowsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	clientName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientRxFlowsRequest) Count(count int32) AllApiApiGetMsgVpnClientRxFlowsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnClientRxFlowsRequest) Cursor(cursor string) AllApiApiGetMsgVpnClientRxFlowsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnClientRxFlowsRequest) Where(where []string) AllApiApiGetMsgVpnClientRxFlowsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnClientRxFlowsRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientRxFlowsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientRxFlowsRequest) Execute() (MsgVpnClientRxFlowsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientRxFlowsRequest
*/
func (a *AllApiService) GetMsgVpnClientRxFlows(ctx _context.Context, msgVpnName string, clientName string) AllApiApiGetMsgVpnClientRxFlowsRequest {
	return AllApiApiGetMsgVpnClientRxFlowsRequest{
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
func (a *AllApiService) GetMsgVpnClientRxFlowsExecute(r AllApiApiGetMsgVpnClientRxFlowsRequest) (MsgVpnClientRxFlowsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientRxFlowsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientRxFlows")
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

type AllApiApiGetMsgVpnClientSubscriptionRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	clientName        string
	subscriptionTopic string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnClientSubscriptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientSubscriptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientSubscriptionRequest) Execute() (MsgVpnClientSubscriptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientSubscriptionRequest
*/
func (a *AllApiService) GetMsgVpnClientSubscription(ctx _context.Context, msgVpnName string, clientName string, subscriptionTopic string) AllApiApiGetMsgVpnClientSubscriptionRequest {
	return AllApiApiGetMsgVpnClientSubscriptionRequest{
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
func (a *AllApiService) GetMsgVpnClientSubscriptionExecute(r AllApiApiGetMsgVpnClientSubscriptionRequest) (MsgVpnClientSubscriptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientSubscription")
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

type AllApiApiGetMsgVpnClientSubscriptionsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	clientName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientSubscriptionsRequest) Count(count int32) AllApiApiGetMsgVpnClientSubscriptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnClientSubscriptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnClientSubscriptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnClientSubscriptionsRequest) Where(where []string) AllApiApiGetMsgVpnClientSubscriptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnClientSubscriptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientSubscriptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientSubscriptionsRequest) Execute() (MsgVpnClientSubscriptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientSubscriptionsRequest
*/
func (a *AllApiService) GetMsgVpnClientSubscriptions(ctx _context.Context, msgVpnName string, clientName string) AllApiApiGetMsgVpnClientSubscriptionsRequest {
	return AllApiApiGetMsgVpnClientSubscriptionsRequest{
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
func (a *AllApiService) GetMsgVpnClientSubscriptionsExecute(r AllApiApiGetMsgVpnClientSubscriptionsRequest) (MsgVpnClientSubscriptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientSubscriptions")
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

type AllApiApiGetMsgVpnClientTransactedSessionRequest struct {
	ctx         _context.Context
	ApiService  *AllApiService
	msgVpnName  string
	clientName  string
	sessionName string
	select_     *[]string
}

func (r AllApiApiGetMsgVpnClientTransactedSessionRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientTransactedSessionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientTransactedSessionRequest) Execute() (MsgVpnClientTransactedSessionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientTransactedSessionRequest
*/
func (a *AllApiService) GetMsgVpnClientTransactedSession(ctx _context.Context, msgVpnName string, clientName string, sessionName string) AllApiApiGetMsgVpnClientTransactedSessionRequest {
	return AllApiApiGetMsgVpnClientTransactedSessionRequest{
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
func (a *AllApiService) GetMsgVpnClientTransactedSessionExecute(r AllApiApiGetMsgVpnClientTransactedSessionRequest) (MsgVpnClientTransactedSessionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientTransactedSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientTransactedSession")
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

type AllApiApiGetMsgVpnClientTransactedSessionsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	clientName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientTransactedSessionsRequest) Count(count int32) AllApiApiGetMsgVpnClientTransactedSessionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnClientTransactedSessionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnClientTransactedSessionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnClientTransactedSessionsRequest) Where(where []string) AllApiApiGetMsgVpnClientTransactedSessionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnClientTransactedSessionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientTransactedSessionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientTransactedSessionsRequest) Execute() (MsgVpnClientTransactedSessionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientTransactedSessionsRequest
*/
func (a *AllApiService) GetMsgVpnClientTransactedSessions(ctx _context.Context, msgVpnName string, clientName string) AllApiApiGetMsgVpnClientTransactedSessionsRequest {
	return AllApiApiGetMsgVpnClientTransactedSessionsRequest{
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
func (a *AllApiService) GetMsgVpnClientTransactedSessionsExecute(r AllApiApiGetMsgVpnClientTransactedSessionsRequest) (MsgVpnClientTransactedSessionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientTransactedSessionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientTransactedSessions")
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

type AllApiApiGetMsgVpnClientTxFlowRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	clientName string
	flowId     string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientTxFlowRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientTxFlowRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientTxFlowRequest) Execute() (MsgVpnClientTxFlowResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientTxFlowRequest
*/
func (a *AllApiService) GetMsgVpnClientTxFlow(ctx _context.Context, msgVpnName string, clientName string, flowId string) AllApiApiGetMsgVpnClientTxFlowRequest {
	return AllApiApiGetMsgVpnClientTxFlowRequest{
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
func (a *AllApiService) GetMsgVpnClientTxFlowExecute(r AllApiApiGetMsgVpnClientTxFlowRequest) (MsgVpnClientTxFlowResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientTxFlowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientTxFlow")
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

type AllApiApiGetMsgVpnClientTxFlowsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	clientName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientTxFlowsRequest) Count(count int32) AllApiApiGetMsgVpnClientTxFlowsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnClientTxFlowsRequest) Cursor(cursor string) AllApiApiGetMsgVpnClientTxFlowsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnClientTxFlowsRequest) Where(where []string) AllApiApiGetMsgVpnClientTxFlowsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnClientTxFlowsRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientTxFlowsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientTxFlowsRequest) Execute() (MsgVpnClientTxFlowsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientTxFlowsRequest
*/
func (a *AllApiService) GetMsgVpnClientTxFlows(ctx _context.Context, msgVpnName string, clientName string) AllApiApiGetMsgVpnClientTxFlowsRequest {
	return AllApiApiGetMsgVpnClientTxFlowsRequest{
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
func (a *AllApiService) GetMsgVpnClientTxFlowsExecute(r AllApiApiGetMsgVpnClientTxFlowsRequest) (MsgVpnClientTxFlowsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientTxFlowsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientTxFlows")
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

type AllApiApiGetMsgVpnClientUsernameRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	msgVpnName     string
	clientUsername string
	select_        *[]string
}

func (r AllApiApiGetMsgVpnClientUsernameRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientUsernameRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientUsernameRequest) Execute() (MsgVpnClientUsernameResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientUsernameRequest
*/
func (a *AllApiService) GetMsgVpnClientUsername(ctx _context.Context, msgVpnName string, clientUsername string) AllApiApiGetMsgVpnClientUsernameRequest {
	return AllApiApiGetMsgVpnClientUsernameRequest{
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
func (a *AllApiService) GetMsgVpnClientUsernameExecute(r AllApiApiGetMsgVpnClientUsernameRequest) (MsgVpnClientUsernameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientUsernameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientUsername")
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

type AllApiApiGetMsgVpnClientUsernamesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientUsernamesRequest) Count(count int32) AllApiApiGetMsgVpnClientUsernamesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnClientUsernamesRequest) Cursor(cursor string) AllApiApiGetMsgVpnClientUsernamesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnClientUsernamesRequest) Where(where []string) AllApiApiGetMsgVpnClientUsernamesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnClientUsernamesRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientUsernamesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientUsernamesRequest) Execute() (MsgVpnClientUsernamesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientUsernamesRequest
*/
func (a *AllApiService) GetMsgVpnClientUsernames(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnClientUsernamesRequest {
	return AllApiApiGetMsgVpnClientUsernamesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientUsernamesResponse
 */
func (a *AllApiService) GetMsgVpnClientUsernamesExecute(r AllApiApiGetMsgVpnClientUsernamesRequest) (MsgVpnClientUsernamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientUsernamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClientUsernames")
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

type AllApiApiGetMsgVpnClientsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnClientsRequest) Count(count int32) AllApiApiGetMsgVpnClientsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnClientsRequest) Cursor(cursor string) AllApiApiGetMsgVpnClientsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnClientsRequest) Where(where []string) AllApiApiGetMsgVpnClientsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnClientsRequest) Select_(select_ []string) AllApiApiGetMsgVpnClientsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnClientsRequest) Execute() (MsgVpnClientsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnClientsRequest
*/
func (a *AllApiService) GetMsgVpnClients(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnClientsRequest {
	return AllApiApiGetMsgVpnClientsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnClientsResponse
 */
func (a *AllApiService) GetMsgVpnClientsExecute(r AllApiApiGetMsgVpnClientsRequest) (MsgVpnClientsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnClientsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnClients")
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

type AllApiApiGetMsgVpnConfigSyncRemoteNodeRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	msgVpnName     string
	remoteNodeName string
	select_        *[]string
}

func (r AllApiApiGetMsgVpnConfigSyncRemoteNodeRequest) Select_(select_ []string) AllApiApiGetMsgVpnConfigSyncRemoteNodeRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnConfigSyncRemoteNodeRequest) Execute() (MsgVpnConfigSyncRemoteNodeResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnConfigSyncRemoteNodeRequest
*/
func (a *AllApiService) GetMsgVpnConfigSyncRemoteNode(ctx _context.Context, msgVpnName string, remoteNodeName string) AllApiApiGetMsgVpnConfigSyncRemoteNodeRequest {
	return AllApiApiGetMsgVpnConfigSyncRemoteNodeRequest{
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
func (a *AllApiService) GetMsgVpnConfigSyncRemoteNodeExecute(r AllApiApiGetMsgVpnConfigSyncRemoteNodeRequest) (MsgVpnConfigSyncRemoteNodeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnConfigSyncRemoteNodeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnConfigSyncRemoteNode")
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

type AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest) Count(count int32) AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest) Cursor(cursor string) AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest) Where(where []string) AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest) Select_(select_ []string) AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest) Execute() (MsgVpnConfigSyncRemoteNodesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest
*/
func (a *AllApiService) GetMsgVpnConfigSyncRemoteNodes(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest {
	return AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnConfigSyncRemoteNodesResponse
 */
func (a *AllApiService) GetMsgVpnConfigSyncRemoteNodesExecute(r AllApiApiGetMsgVpnConfigSyncRemoteNodesRequest) (MsgVpnConfigSyncRemoteNodesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnConfigSyncRemoteNodesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnConfigSyncRemoteNodes")
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

type AllApiApiGetMsgVpnDistributedCacheRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	cacheName  string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheRequest) Execute() (MsgVpnDistributedCacheResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCache(ctx _context.Context, msgVpnName string, cacheName string) AllApiApiGetMsgVpnDistributedCacheRequest {
	return AllApiApiGetMsgVpnDistributedCacheRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheExecute(r AllApiApiGetMsgVpnDistributedCacheRequest) (MsgVpnDistributedCacheResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCache")
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

type AllApiApiGetMsgVpnDistributedCacheClusterRequest struct {
	ctx         _context.Context
	ApiService  *AllApiService
	msgVpnName  string
	cacheName   string
	clusterName string
	select_     *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterRequest) Execute() (MsgVpnDistributedCacheClusterResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheCluster(ctx _context.Context, msgVpnName string, cacheName string, clusterName string) AllApiApiGetMsgVpnDistributedCacheClusterRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterExecute(r AllApiApiGetMsgVpnDistributedCacheClusterRequest) (MsgVpnDistributedCacheClusterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheCluster")
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

type AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest struct {
	ctx             _context.Context
	ApiService      *AllApiService
	msgVpnName      string
	cacheName       string
	clusterName     string
	homeClusterName string
	select_         *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest) Execute() (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, homeClusterName string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterExecute(r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest) (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster")
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

type AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest struct {
	ctx             _context.Context
	ApiService      *AllApiService
	msgVpnName      string
	cacheName       string
	clusterName     string
	homeClusterName string
	topicPrefix     string
	select_         *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest) Execute() (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, homeClusterName string, topicPrefix string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixExecute(r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest) (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix")
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

type AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest struct {
	ctx             _context.Context
	ApiService      *AllApiService
	msgVpnName      string
	cacheName       string
	clusterName     string
	homeClusterName string
	count           *int32
	cursor          *string
	where           *[]string
	select_         *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) Count(count int32) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) Cursor(cursor string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) Where(where []string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) Execute() (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, homeClusterName string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesExecute(r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest) (MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes")
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

type AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest struct {
	ctx         _context.Context
	ApiService  *AllApiService
	msgVpnName  string
	cacheName   string
	clusterName string
	count       *int32
	cursor      *string
	where       *[]string
	select_     *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) Count(count int32) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) Cursor(cursor string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) Where(where []string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) Execute() (MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters(ctx _context.Context, msgVpnName string, cacheName string, clusterName string) AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersExecute(r AllApiApiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest) (MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters")
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

type AllApiApiGetMsgVpnDistributedCacheClusterInstanceRequest struct {
	ctx          _context.Context
	ApiService   *AllApiService
	msgVpnName   string
	cacheName    string
	clusterName  string
	instanceName string
	select_      *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRequest) Execute() (MsgVpnDistributedCacheClusterInstanceResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterInstanceRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstance(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterInstanceRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstanceExecute(r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRequest) (MsgVpnDistributedCacheClusterInstanceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstanceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterInstance")
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

type AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest struct {
	ctx             _context.Context
	ApiService      *AllApiService
	msgVpnName      string
	cacheName       string
	clusterName     string
	instanceName    string
	homeClusterName string
	select_         *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest) Execute() (MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string, homeClusterName string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterExecute(r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterRequest) (MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster")
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

type AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest struct {
	ctx          _context.Context
	ApiService   *AllApiService
	msgVpnName   string
	cacheName    string
	clusterName  string
	instanceName string
	count        *int32
	cursor       *string
	where        *[]string
	select_      *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) Count(count int32) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) Cursor(cursor string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) Where(where []string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) Execute() (MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersExecute(r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersRequest) (MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClustersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusters")
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

type AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest struct {
	ctx          _context.Context
	ApiService   *AllApiService
	msgVpnName   string
	cacheName    string
	clusterName  string
	instanceName string
	topic        string
	select_      *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest) Execute() (MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteTopic(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string, topic string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteTopicExecute(r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicRequest) (MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstanceRemoteTopicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteTopic")
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

type AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest struct {
	ctx          _context.Context
	ApiService   *AllApiService
	msgVpnName   string
	cacheName    string
	clusterName  string
	instanceName string
	count        *int32
	cursor       *string
	where        *[]string
	select_      *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) Count(count int32) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) Cursor(cursor string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) Where(where []string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) Execute() (MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteTopics(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, instanceName string) AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstanceRemoteTopicsExecute(r AllApiApiGetMsgVpnDistributedCacheClusterInstanceRemoteTopicsRequest) (MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstanceRemoteTopicsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterInstanceRemoteTopics")
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

type AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest struct {
	ctx         _context.Context
	ApiService  *AllApiService
	msgVpnName  string
	cacheName   string
	clusterName string
	count       *int32
	cursor      *string
	where       *[]string
	select_     *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) Count(count int32) AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) Cursor(cursor string) AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) Where(where []string) AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) Execute() (MsgVpnDistributedCacheClusterInstancesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstances(ctx _context.Context, msgVpnName string, cacheName string, clusterName string) AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterInstancesExecute(r AllApiApiGetMsgVpnDistributedCacheClusterInstancesRequest) (MsgVpnDistributedCacheClusterInstancesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterInstancesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterInstances")
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

type AllApiApiGetMsgVpnDistributedCacheClusterTopicRequest struct {
	ctx         _context.Context
	ApiService  *AllApiService
	msgVpnName  string
	cacheName   string
	clusterName string
	topic       string
	select_     *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterTopicRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterTopicRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterTopicRequest) Execute() (MsgVpnDistributedCacheClusterTopicResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterTopicRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterTopic(ctx _context.Context, msgVpnName string, cacheName string, clusterName string, topic string) AllApiApiGetMsgVpnDistributedCacheClusterTopicRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterTopicRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterTopicExecute(r AllApiApiGetMsgVpnDistributedCacheClusterTopicRequest) (MsgVpnDistributedCacheClusterTopicResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterTopicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterTopic")
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

type AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest struct {
	ctx         _context.Context
	ApiService  *AllApiService
	msgVpnName  string
	cacheName   string
	clusterName string
	count       *int32
	cursor      *string
	where       *[]string
	select_     *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) Count(count int32) AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) Cursor(cursor string) AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) Where(where []string) AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) Execute() (MsgVpnDistributedCacheClusterTopicsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusterTopics(ctx _context.Context, msgVpnName string, cacheName string, clusterName string) AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest {
	return AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClusterTopicsExecute(r AllApiApiGetMsgVpnDistributedCacheClusterTopicsRequest) (MsgVpnDistributedCacheClusterTopicsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClusterTopicsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusterTopics")
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

type AllApiApiGetMsgVpnDistributedCacheClustersRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	cacheName  string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnDistributedCacheClustersRequest) Count(count int32) AllApiApiGetMsgVpnDistributedCacheClustersRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClustersRequest) Cursor(cursor string) AllApiApiGetMsgVpnDistributedCacheClustersRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClustersRequest) Where(where []string) AllApiApiGetMsgVpnDistributedCacheClustersRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnDistributedCacheClustersRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCacheClustersRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCacheClustersRequest) Execute() (MsgVpnDistributedCacheClustersResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCacheClustersRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCacheClusters(ctx _context.Context, msgVpnName string, cacheName string) AllApiApiGetMsgVpnDistributedCacheClustersRequest {
	return AllApiApiGetMsgVpnDistributedCacheClustersRequest{
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
func (a *AllApiService) GetMsgVpnDistributedCacheClustersExecute(r AllApiApiGetMsgVpnDistributedCacheClustersRequest) (MsgVpnDistributedCacheClustersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCacheClustersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCacheClusters")
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

type AllApiApiGetMsgVpnDistributedCachesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnDistributedCachesRequest) Count(count int32) AllApiApiGetMsgVpnDistributedCachesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnDistributedCachesRequest) Cursor(cursor string) AllApiApiGetMsgVpnDistributedCachesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnDistributedCachesRequest) Where(where []string) AllApiApiGetMsgVpnDistributedCachesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnDistributedCachesRequest) Select_(select_ []string) AllApiApiGetMsgVpnDistributedCachesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDistributedCachesRequest) Execute() (MsgVpnDistributedCachesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDistributedCachesRequest
*/
func (a *AllApiService) GetMsgVpnDistributedCaches(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnDistributedCachesRequest {
	return AllApiApiGetMsgVpnDistributedCachesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDistributedCachesResponse
 */
func (a *AllApiService) GetMsgVpnDistributedCachesExecute(r AllApiApiGetMsgVpnDistributedCachesRequest) (MsgVpnDistributedCachesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDistributedCachesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDistributedCaches")
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

type AllApiApiGetMsgVpnDmrBridgeRequest struct {
	ctx            _context.Context
	ApiService     *AllApiService
	msgVpnName     string
	remoteNodeName string
	select_        *[]string
}

func (r AllApiApiGetMsgVpnDmrBridgeRequest) Select_(select_ []string) AllApiApiGetMsgVpnDmrBridgeRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDmrBridgeRequest) Execute() (MsgVpnDmrBridgeResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDmrBridgeRequest
*/
func (a *AllApiService) GetMsgVpnDmrBridge(ctx _context.Context, msgVpnName string, remoteNodeName string) AllApiApiGetMsgVpnDmrBridgeRequest {
	return AllApiApiGetMsgVpnDmrBridgeRequest{
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
func (a *AllApiService) GetMsgVpnDmrBridgeExecute(r AllApiApiGetMsgVpnDmrBridgeRequest) (MsgVpnDmrBridgeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDmrBridgeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDmrBridge")
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

type AllApiApiGetMsgVpnDmrBridgesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnDmrBridgesRequest) Count(count int32) AllApiApiGetMsgVpnDmrBridgesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnDmrBridgesRequest) Cursor(cursor string) AllApiApiGetMsgVpnDmrBridgesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnDmrBridgesRequest) Where(where []string) AllApiApiGetMsgVpnDmrBridgesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnDmrBridgesRequest) Select_(select_ []string) AllApiApiGetMsgVpnDmrBridgesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnDmrBridgesRequest) Execute() (MsgVpnDmrBridgesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnDmrBridgesRequest
*/
func (a *AllApiService) GetMsgVpnDmrBridges(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnDmrBridgesRequest {
	return AllApiApiGetMsgVpnDmrBridgesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnDmrBridgesResponse
 */
func (a *AllApiService) GetMsgVpnDmrBridgesExecute(r AllApiApiGetMsgVpnDmrBridgesRequest) (MsgVpnDmrBridgesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnDmrBridgesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnDmrBridges")
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

type AllApiApiGetMsgVpnJndiConnectionFactoriesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnJndiConnectionFactoriesRequest) Count(count int32) AllApiApiGetMsgVpnJndiConnectionFactoriesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnJndiConnectionFactoriesRequest) Cursor(cursor string) AllApiApiGetMsgVpnJndiConnectionFactoriesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnJndiConnectionFactoriesRequest) Where(where []string) AllApiApiGetMsgVpnJndiConnectionFactoriesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnJndiConnectionFactoriesRequest) Select_(select_ []string) AllApiApiGetMsgVpnJndiConnectionFactoriesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnJndiConnectionFactoriesRequest) Execute() (MsgVpnJndiConnectionFactoriesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnJndiConnectionFactoriesRequest
*/
func (a *AllApiService) GetMsgVpnJndiConnectionFactories(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnJndiConnectionFactoriesRequest {
	return AllApiApiGetMsgVpnJndiConnectionFactoriesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnJndiConnectionFactoriesResponse
 */
func (a *AllApiService) GetMsgVpnJndiConnectionFactoriesExecute(r AllApiApiGetMsgVpnJndiConnectionFactoriesRequest) (MsgVpnJndiConnectionFactoriesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiConnectionFactoriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnJndiConnectionFactories")
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

type AllApiApiGetMsgVpnJndiConnectionFactoryRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	connectionFactoryName string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnJndiConnectionFactoryRequest) Select_(select_ []string) AllApiApiGetMsgVpnJndiConnectionFactoryRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnJndiConnectionFactoryRequest) Execute() (MsgVpnJndiConnectionFactoryResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnJndiConnectionFactoryRequest
*/
func (a *AllApiService) GetMsgVpnJndiConnectionFactory(ctx _context.Context, msgVpnName string, connectionFactoryName string) AllApiApiGetMsgVpnJndiConnectionFactoryRequest {
	return AllApiApiGetMsgVpnJndiConnectionFactoryRequest{
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
func (a *AllApiService) GetMsgVpnJndiConnectionFactoryExecute(r AllApiApiGetMsgVpnJndiConnectionFactoryRequest) (MsgVpnJndiConnectionFactoryResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiConnectionFactoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnJndiConnectionFactory")
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

type AllApiApiGetMsgVpnJndiQueueRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	queueName  string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnJndiQueueRequest) Select_(select_ []string) AllApiApiGetMsgVpnJndiQueueRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnJndiQueueRequest) Execute() (MsgVpnJndiQueueResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnJndiQueueRequest
*/
func (a *AllApiService) GetMsgVpnJndiQueue(ctx _context.Context, msgVpnName string, queueName string) AllApiApiGetMsgVpnJndiQueueRequest {
	return AllApiApiGetMsgVpnJndiQueueRequest{
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
func (a *AllApiService) GetMsgVpnJndiQueueExecute(r AllApiApiGetMsgVpnJndiQueueRequest) (MsgVpnJndiQueueResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiQueueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnJndiQueue")
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

type AllApiApiGetMsgVpnJndiQueuesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnJndiQueuesRequest) Count(count int32) AllApiApiGetMsgVpnJndiQueuesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnJndiQueuesRequest) Cursor(cursor string) AllApiApiGetMsgVpnJndiQueuesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnJndiQueuesRequest) Where(where []string) AllApiApiGetMsgVpnJndiQueuesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnJndiQueuesRequest) Select_(select_ []string) AllApiApiGetMsgVpnJndiQueuesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnJndiQueuesRequest) Execute() (MsgVpnJndiQueuesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnJndiQueuesRequest
*/
func (a *AllApiService) GetMsgVpnJndiQueues(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnJndiQueuesRequest {
	return AllApiApiGetMsgVpnJndiQueuesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnJndiQueuesResponse
 */
func (a *AllApiService) GetMsgVpnJndiQueuesExecute(r AllApiApiGetMsgVpnJndiQueuesRequest) (MsgVpnJndiQueuesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiQueuesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnJndiQueues")
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

type AllApiApiGetMsgVpnJndiTopicRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	topicName  string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnJndiTopicRequest) Select_(select_ []string) AllApiApiGetMsgVpnJndiTopicRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnJndiTopicRequest) Execute() (MsgVpnJndiTopicResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnJndiTopicRequest
*/
func (a *AllApiService) GetMsgVpnJndiTopic(ctx _context.Context, msgVpnName string, topicName string) AllApiApiGetMsgVpnJndiTopicRequest {
	return AllApiApiGetMsgVpnJndiTopicRequest{
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
func (a *AllApiService) GetMsgVpnJndiTopicExecute(r AllApiApiGetMsgVpnJndiTopicRequest) (MsgVpnJndiTopicResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiTopicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnJndiTopic")
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

type AllApiApiGetMsgVpnJndiTopicsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnJndiTopicsRequest) Count(count int32) AllApiApiGetMsgVpnJndiTopicsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnJndiTopicsRequest) Cursor(cursor string) AllApiApiGetMsgVpnJndiTopicsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnJndiTopicsRequest) Where(where []string) AllApiApiGetMsgVpnJndiTopicsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnJndiTopicsRequest) Select_(select_ []string) AllApiApiGetMsgVpnJndiTopicsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnJndiTopicsRequest) Execute() (MsgVpnJndiTopicsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnJndiTopicsRequest
*/
func (a *AllApiService) GetMsgVpnJndiTopics(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnJndiTopicsRequest {
	return AllApiApiGetMsgVpnJndiTopicsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnJndiTopicsResponse
 */
func (a *AllApiService) GetMsgVpnJndiTopicsExecute(r AllApiApiGetMsgVpnJndiTopicsRequest) (MsgVpnJndiTopicsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnJndiTopicsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnJndiTopics")
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

type AllApiApiGetMsgVpnMqttRetainCacheRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	cacheName  string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnMqttRetainCacheRequest) Select_(select_ []string) AllApiApiGetMsgVpnMqttRetainCacheRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnMqttRetainCacheRequest) Execute() (MsgVpnMqttRetainCacheResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnMqttRetainCacheRequest
*/
func (a *AllApiService) GetMsgVpnMqttRetainCache(ctx _context.Context, msgVpnName string, cacheName string) AllApiApiGetMsgVpnMqttRetainCacheRequest {
	return AllApiApiGetMsgVpnMqttRetainCacheRequest{
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
func (a *AllApiService) GetMsgVpnMqttRetainCacheExecute(r AllApiApiGetMsgVpnMqttRetainCacheRequest) (MsgVpnMqttRetainCacheResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttRetainCacheResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnMqttRetainCache")
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

type AllApiApiGetMsgVpnMqttRetainCachesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnMqttRetainCachesRequest) Count(count int32) AllApiApiGetMsgVpnMqttRetainCachesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnMqttRetainCachesRequest) Cursor(cursor string) AllApiApiGetMsgVpnMqttRetainCachesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnMqttRetainCachesRequest) Where(where []string) AllApiApiGetMsgVpnMqttRetainCachesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnMqttRetainCachesRequest) Select_(select_ []string) AllApiApiGetMsgVpnMqttRetainCachesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnMqttRetainCachesRequest) Execute() (MsgVpnMqttRetainCachesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnMqttRetainCachesRequest
*/
func (a *AllApiService) GetMsgVpnMqttRetainCaches(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnMqttRetainCachesRequest {
	return AllApiApiGetMsgVpnMqttRetainCachesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnMqttRetainCachesResponse
 */
func (a *AllApiService) GetMsgVpnMqttRetainCachesExecute(r AllApiApiGetMsgVpnMqttRetainCachesRequest) (MsgVpnMqttRetainCachesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttRetainCachesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnMqttRetainCaches")
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

type AllApiApiGetMsgVpnMqttSessionRequest struct {
	ctx                      _context.Context
	ApiService               *AllApiService
	msgVpnName               string
	mqttSessionClientId      string
	mqttSessionVirtualRouter string
	select_                  *[]string
}

func (r AllApiApiGetMsgVpnMqttSessionRequest) Select_(select_ []string) AllApiApiGetMsgVpnMqttSessionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnMqttSessionRequest) Execute() (MsgVpnMqttSessionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnMqttSessionRequest
*/
func (a *AllApiService) GetMsgVpnMqttSession(ctx _context.Context, msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string) AllApiApiGetMsgVpnMqttSessionRequest {
	return AllApiApiGetMsgVpnMqttSessionRequest{
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
func (a *AllApiService) GetMsgVpnMqttSessionExecute(r AllApiApiGetMsgVpnMqttSessionRequest) (MsgVpnMqttSessionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnMqttSession")
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

type AllApiApiGetMsgVpnMqttSessionSubscriptionRequest struct {
	ctx                      _context.Context
	ApiService               *AllApiService
	msgVpnName               string
	mqttSessionClientId      string
	mqttSessionVirtualRouter string
	subscriptionTopic        string
	select_                  *[]string
}

func (r AllApiApiGetMsgVpnMqttSessionSubscriptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnMqttSessionSubscriptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnMqttSessionSubscriptionRequest) Execute() (MsgVpnMqttSessionSubscriptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnMqttSessionSubscriptionRequest
*/
func (a *AllApiService) GetMsgVpnMqttSessionSubscription(ctx _context.Context, msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string, subscriptionTopic string) AllApiApiGetMsgVpnMqttSessionSubscriptionRequest {
	return AllApiApiGetMsgVpnMqttSessionSubscriptionRequest{
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
func (a *AllApiService) GetMsgVpnMqttSessionSubscriptionExecute(r AllApiApiGetMsgVpnMqttSessionSubscriptionRequest) (MsgVpnMqttSessionSubscriptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttSessionSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnMqttSessionSubscription")
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

type AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest struct {
	ctx                      _context.Context
	ApiService               *AllApiService
	msgVpnName               string
	mqttSessionClientId      string
	mqttSessionVirtualRouter string
	count                    *int32
	cursor                   *string
	where                    *[]string
	select_                  *[]string
}

func (r AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest) Count(count int32) AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest) Where(where []string) AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest) Execute() (MsgVpnMqttSessionSubscriptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest
*/
func (a *AllApiService) GetMsgVpnMqttSessionSubscriptions(ctx _context.Context, msgVpnName string, mqttSessionClientId string, mqttSessionVirtualRouter string) AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest {
	return AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest{
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
func (a *AllApiService) GetMsgVpnMqttSessionSubscriptionsExecute(r AllApiApiGetMsgVpnMqttSessionSubscriptionsRequest) (MsgVpnMqttSessionSubscriptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttSessionSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnMqttSessionSubscriptions")
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

type AllApiApiGetMsgVpnMqttSessionsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnMqttSessionsRequest) Count(count int32) AllApiApiGetMsgVpnMqttSessionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnMqttSessionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnMqttSessionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnMqttSessionsRequest) Where(where []string) AllApiApiGetMsgVpnMqttSessionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnMqttSessionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnMqttSessionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnMqttSessionsRequest) Execute() (MsgVpnMqttSessionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnMqttSessionsRequest
*/
func (a *AllApiService) GetMsgVpnMqttSessions(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnMqttSessionsRequest {
	return AllApiApiGetMsgVpnMqttSessionsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnMqttSessionsResponse
 */
func (a *AllApiService) GetMsgVpnMqttSessionsExecute(r AllApiApiGetMsgVpnMqttSessionsRequest) (MsgVpnMqttSessionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnMqttSessionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnMqttSessions")
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

type AllApiApiGetMsgVpnQueueRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	queueName  string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnQueueRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueueRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueueRequest) Execute() (MsgVpnQueueResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueueRequest
*/
func (a *AllApiService) GetMsgVpnQueue(ctx _context.Context, msgVpnName string, queueName string) AllApiApiGetMsgVpnQueueRequest {
	return AllApiApiGetMsgVpnQueueRequest{
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
func (a *AllApiService) GetMsgVpnQueueExecute(r AllApiApiGetMsgVpnQueueRequest) (MsgVpnQueueResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueue")
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

type AllApiApiGetMsgVpnQueueMsgRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	queueName  string
	msgId      string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnQueueMsgRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueueMsgRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueueMsgRequest) Execute() (MsgVpnQueueMsgResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueueMsgRequest
*/
func (a *AllApiService) GetMsgVpnQueueMsg(ctx _context.Context, msgVpnName string, queueName string, msgId string) AllApiApiGetMsgVpnQueueMsgRequest {
	return AllApiApiGetMsgVpnQueueMsgRequest{
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
func (a *AllApiService) GetMsgVpnQueueMsgExecute(r AllApiApiGetMsgVpnQueueMsgRequest) (MsgVpnQueueMsgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueMsgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueueMsg")
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

type AllApiApiGetMsgVpnQueueMsgsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	queueName  string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnQueueMsgsRequest) Count(count int32) AllApiApiGetMsgVpnQueueMsgsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnQueueMsgsRequest) Cursor(cursor string) AllApiApiGetMsgVpnQueueMsgsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnQueueMsgsRequest) Where(where []string) AllApiApiGetMsgVpnQueueMsgsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnQueueMsgsRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueueMsgsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueueMsgsRequest) Execute() (MsgVpnQueueMsgsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueueMsgsRequest
*/
func (a *AllApiService) GetMsgVpnQueueMsgs(ctx _context.Context, msgVpnName string, queueName string) AllApiApiGetMsgVpnQueueMsgsRequest {
	return AllApiApiGetMsgVpnQueueMsgsRequest{
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
func (a *AllApiService) GetMsgVpnQueueMsgsExecute(r AllApiApiGetMsgVpnQueueMsgsRequest) (MsgVpnQueueMsgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueMsgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueueMsgs")
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

type AllApiApiGetMsgVpnQueuePrioritiesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	queueName  string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnQueuePrioritiesRequest) Count(count int32) AllApiApiGetMsgVpnQueuePrioritiesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnQueuePrioritiesRequest) Cursor(cursor string) AllApiApiGetMsgVpnQueuePrioritiesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnQueuePrioritiesRequest) Where(where []string) AllApiApiGetMsgVpnQueuePrioritiesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnQueuePrioritiesRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueuePrioritiesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueuePrioritiesRequest) Execute() (MsgVpnQueuePrioritiesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueuePrioritiesRequest
*/
func (a *AllApiService) GetMsgVpnQueuePriorities(ctx _context.Context, msgVpnName string, queueName string) AllApiApiGetMsgVpnQueuePrioritiesRequest {
	return AllApiApiGetMsgVpnQueuePrioritiesRequest{
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
func (a *AllApiService) GetMsgVpnQueuePrioritiesExecute(r AllApiApiGetMsgVpnQueuePrioritiesRequest) (MsgVpnQueuePrioritiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueuePrioritiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueuePriorities")
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

type AllApiApiGetMsgVpnQueuePriorityRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	queueName  string
	priority   string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnQueuePriorityRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueuePriorityRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueuePriorityRequest) Execute() (MsgVpnQueuePriorityResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueuePriorityRequest
*/
func (a *AllApiService) GetMsgVpnQueuePriority(ctx _context.Context, msgVpnName string, queueName string, priority string) AllApiApiGetMsgVpnQueuePriorityRequest {
	return AllApiApiGetMsgVpnQueuePriorityRequest{
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
func (a *AllApiService) GetMsgVpnQueuePriorityExecute(r AllApiApiGetMsgVpnQueuePriorityRequest) (MsgVpnQueuePriorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueuePriorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueuePriority")
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

type AllApiApiGetMsgVpnQueueSubscriptionRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	queueName         string
	subscriptionTopic string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnQueueSubscriptionRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueueSubscriptionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueueSubscriptionRequest) Execute() (MsgVpnQueueSubscriptionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueueSubscriptionRequest
*/
func (a *AllApiService) GetMsgVpnQueueSubscription(ctx _context.Context, msgVpnName string, queueName string, subscriptionTopic string) AllApiApiGetMsgVpnQueueSubscriptionRequest {
	return AllApiApiGetMsgVpnQueueSubscriptionRequest{
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
func (a *AllApiService) GetMsgVpnQueueSubscriptionExecute(r AllApiApiGetMsgVpnQueueSubscriptionRequest) (MsgVpnQueueSubscriptionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueueSubscription")
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

type AllApiApiGetMsgVpnQueueSubscriptionsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	queueName  string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnQueueSubscriptionsRequest) Count(count int32) AllApiApiGetMsgVpnQueueSubscriptionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnQueueSubscriptionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnQueueSubscriptionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnQueueSubscriptionsRequest) Where(where []string) AllApiApiGetMsgVpnQueueSubscriptionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnQueueSubscriptionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueueSubscriptionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueueSubscriptionsRequest) Execute() (MsgVpnQueueSubscriptionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueueSubscriptionsRequest
*/
func (a *AllApiService) GetMsgVpnQueueSubscriptions(ctx _context.Context, msgVpnName string, queueName string) AllApiApiGetMsgVpnQueueSubscriptionsRequest {
	return AllApiApiGetMsgVpnQueueSubscriptionsRequest{
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
func (a *AllApiService) GetMsgVpnQueueSubscriptionsExecute(r AllApiApiGetMsgVpnQueueSubscriptionsRequest) (MsgVpnQueueSubscriptionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueSubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueueSubscriptions")
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

type AllApiApiGetMsgVpnQueueTemplateRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	queueTemplateName string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnQueueTemplateRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueueTemplateRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueueTemplateRequest) Execute() (MsgVpnQueueTemplateResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueueTemplateRequest
*/
func (a *AllApiService) GetMsgVpnQueueTemplate(ctx _context.Context, msgVpnName string, queueTemplateName string) AllApiApiGetMsgVpnQueueTemplateRequest {
	return AllApiApiGetMsgVpnQueueTemplateRequest{
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
func (a *AllApiService) GetMsgVpnQueueTemplateExecute(r AllApiApiGetMsgVpnQueueTemplateRequest) (MsgVpnQueueTemplateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueTemplateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueueTemplate")
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

type AllApiApiGetMsgVpnQueueTemplatesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnQueueTemplatesRequest) Count(count int32) AllApiApiGetMsgVpnQueueTemplatesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnQueueTemplatesRequest) Cursor(cursor string) AllApiApiGetMsgVpnQueueTemplatesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnQueueTemplatesRequest) Where(where []string) AllApiApiGetMsgVpnQueueTemplatesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnQueueTemplatesRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueueTemplatesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueueTemplatesRequest) Execute() (MsgVpnQueueTemplatesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueueTemplatesRequest
*/
func (a *AllApiService) GetMsgVpnQueueTemplates(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnQueueTemplatesRequest {
	return AllApiApiGetMsgVpnQueueTemplatesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueueTemplatesResponse
 */
func (a *AllApiService) GetMsgVpnQueueTemplatesExecute(r AllApiApiGetMsgVpnQueueTemplatesRequest) (MsgVpnQueueTemplatesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueTemplatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueueTemplates")
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

type AllApiApiGetMsgVpnQueueTxFlowRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	queueName  string
	flowId     string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnQueueTxFlowRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueueTxFlowRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueueTxFlowRequest) Execute() (MsgVpnQueueTxFlowResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueueTxFlowRequest
*/
func (a *AllApiService) GetMsgVpnQueueTxFlow(ctx _context.Context, msgVpnName string, queueName string, flowId string) AllApiApiGetMsgVpnQueueTxFlowRequest {
	return AllApiApiGetMsgVpnQueueTxFlowRequest{
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
func (a *AllApiService) GetMsgVpnQueueTxFlowExecute(r AllApiApiGetMsgVpnQueueTxFlowRequest) (MsgVpnQueueTxFlowResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueTxFlowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueueTxFlow")
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

type AllApiApiGetMsgVpnQueueTxFlowsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	queueName  string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnQueueTxFlowsRequest) Count(count int32) AllApiApiGetMsgVpnQueueTxFlowsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnQueueTxFlowsRequest) Cursor(cursor string) AllApiApiGetMsgVpnQueueTxFlowsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnQueueTxFlowsRequest) Where(where []string) AllApiApiGetMsgVpnQueueTxFlowsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnQueueTxFlowsRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueueTxFlowsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueueTxFlowsRequest) Execute() (MsgVpnQueueTxFlowsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueueTxFlowsRequest
*/
func (a *AllApiService) GetMsgVpnQueueTxFlows(ctx _context.Context, msgVpnName string, queueName string) AllApiApiGetMsgVpnQueueTxFlowsRequest {
	return AllApiApiGetMsgVpnQueueTxFlowsRequest{
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
func (a *AllApiService) GetMsgVpnQueueTxFlowsExecute(r AllApiApiGetMsgVpnQueueTxFlowsRequest) (MsgVpnQueueTxFlowsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueueTxFlowsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueueTxFlows")
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

type AllApiApiGetMsgVpnQueuesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnQueuesRequest) Count(count int32) AllApiApiGetMsgVpnQueuesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnQueuesRequest) Cursor(cursor string) AllApiApiGetMsgVpnQueuesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnQueuesRequest) Where(where []string) AllApiApiGetMsgVpnQueuesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnQueuesRequest) Select_(select_ []string) AllApiApiGetMsgVpnQueuesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnQueuesRequest) Execute() (MsgVpnQueuesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnQueuesRequest
*/
func (a *AllApiService) GetMsgVpnQueues(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnQueuesRequest {
	return AllApiApiGetMsgVpnQueuesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnQueuesResponse
 */
func (a *AllApiService) GetMsgVpnQueuesExecute(r AllApiApiGetMsgVpnQueuesRequest) (MsgVpnQueuesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnQueuesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnQueues")
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

type AllApiApiGetMsgVpnReplayLogRequest struct {
	ctx           _context.Context
	ApiService    *AllApiService
	msgVpnName    string
	replayLogName string
	select_       *[]string
}

func (r AllApiApiGetMsgVpnReplayLogRequest) Select_(select_ []string) AllApiApiGetMsgVpnReplayLogRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnReplayLogRequest) Execute() (MsgVpnReplayLogResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnReplayLogRequest
*/
func (a *AllApiService) GetMsgVpnReplayLog(ctx _context.Context, msgVpnName string, replayLogName string) AllApiApiGetMsgVpnReplayLogRequest {
	return AllApiApiGetMsgVpnReplayLogRequest{
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
func (a *AllApiService) GetMsgVpnReplayLogExecute(r AllApiApiGetMsgVpnReplayLogRequest) (MsgVpnReplayLogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplayLogResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnReplayLog")
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

type AllApiApiGetMsgVpnReplayLogMsgRequest struct {
	ctx           _context.Context
	ApiService    *AllApiService
	msgVpnName    string
	replayLogName string
	msgId         string
	select_       *[]string
}

func (r AllApiApiGetMsgVpnReplayLogMsgRequest) Select_(select_ []string) AllApiApiGetMsgVpnReplayLogMsgRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnReplayLogMsgRequest) Execute() (MsgVpnReplayLogMsgResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnReplayLogMsgRequest
*/
func (a *AllApiService) GetMsgVpnReplayLogMsg(ctx _context.Context, msgVpnName string, replayLogName string, msgId string) AllApiApiGetMsgVpnReplayLogMsgRequest {
	return AllApiApiGetMsgVpnReplayLogMsgRequest{
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
func (a *AllApiService) GetMsgVpnReplayLogMsgExecute(r AllApiApiGetMsgVpnReplayLogMsgRequest) (MsgVpnReplayLogMsgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplayLogMsgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnReplayLogMsg")
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

type AllApiApiGetMsgVpnReplayLogMsgsRequest struct {
	ctx           _context.Context
	ApiService    *AllApiService
	msgVpnName    string
	replayLogName string
	count         *int32
	cursor        *string
	where         *[]string
	select_       *[]string
}

func (r AllApiApiGetMsgVpnReplayLogMsgsRequest) Count(count int32) AllApiApiGetMsgVpnReplayLogMsgsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnReplayLogMsgsRequest) Cursor(cursor string) AllApiApiGetMsgVpnReplayLogMsgsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnReplayLogMsgsRequest) Where(where []string) AllApiApiGetMsgVpnReplayLogMsgsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnReplayLogMsgsRequest) Select_(select_ []string) AllApiApiGetMsgVpnReplayLogMsgsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnReplayLogMsgsRequest) Execute() (MsgVpnReplayLogMsgsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnReplayLogMsgsRequest
*/
func (a *AllApiService) GetMsgVpnReplayLogMsgs(ctx _context.Context, msgVpnName string, replayLogName string) AllApiApiGetMsgVpnReplayLogMsgsRequest {
	return AllApiApiGetMsgVpnReplayLogMsgsRequest{
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
func (a *AllApiService) GetMsgVpnReplayLogMsgsExecute(r AllApiApiGetMsgVpnReplayLogMsgsRequest) (MsgVpnReplayLogMsgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplayLogMsgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnReplayLogMsgs")
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

type AllApiApiGetMsgVpnReplayLogsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnReplayLogsRequest) Count(count int32) AllApiApiGetMsgVpnReplayLogsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnReplayLogsRequest) Cursor(cursor string) AllApiApiGetMsgVpnReplayLogsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnReplayLogsRequest) Where(where []string) AllApiApiGetMsgVpnReplayLogsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnReplayLogsRequest) Select_(select_ []string) AllApiApiGetMsgVpnReplayLogsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnReplayLogsRequest) Execute() (MsgVpnReplayLogsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnReplayLogsRequest
*/
func (a *AllApiService) GetMsgVpnReplayLogs(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnReplayLogsRequest {
	return AllApiApiGetMsgVpnReplayLogsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnReplayLogsResponse
 */
func (a *AllApiService) GetMsgVpnReplayLogsExecute(r AllApiApiGetMsgVpnReplayLogsRequest) (MsgVpnReplayLogsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplayLogsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnReplayLogs")
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

type AllApiApiGetMsgVpnReplicatedTopicRequest struct {
	ctx             _context.Context
	ApiService      *AllApiService
	msgVpnName      string
	replicatedTopic string
	select_         *[]string
}

func (r AllApiApiGetMsgVpnReplicatedTopicRequest) Select_(select_ []string) AllApiApiGetMsgVpnReplicatedTopicRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnReplicatedTopicRequest) Execute() (MsgVpnReplicatedTopicResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnReplicatedTopicRequest
*/
func (a *AllApiService) GetMsgVpnReplicatedTopic(ctx _context.Context, msgVpnName string, replicatedTopic string) AllApiApiGetMsgVpnReplicatedTopicRequest {
	return AllApiApiGetMsgVpnReplicatedTopicRequest{
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
func (a *AllApiService) GetMsgVpnReplicatedTopicExecute(r AllApiApiGetMsgVpnReplicatedTopicRequest) (MsgVpnReplicatedTopicResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplicatedTopicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnReplicatedTopic")
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

type AllApiApiGetMsgVpnReplicatedTopicsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnReplicatedTopicsRequest) Count(count int32) AllApiApiGetMsgVpnReplicatedTopicsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnReplicatedTopicsRequest) Cursor(cursor string) AllApiApiGetMsgVpnReplicatedTopicsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnReplicatedTopicsRequest) Where(where []string) AllApiApiGetMsgVpnReplicatedTopicsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnReplicatedTopicsRequest) Select_(select_ []string) AllApiApiGetMsgVpnReplicatedTopicsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnReplicatedTopicsRequest) Execute() (MsgVpnReplicatedTopicsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnReplicatedTopicsRequest
*/
func (a *AllApiService) GetMsgVpnReplicatedTopics(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnReplicatedTopicsRequest {
	return AllApiApiGetMsgVpnReplicatedTopicsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnReplicatedTopicsResponse
 */
func (a *AllApiService) GetMsgVpnReplicatedTopicsExecute(r AllApiApiGetMsgVpnReplicatedTopicsRequest) (MsgVpnReplicatedTopicsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnReplicatedTopicsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnReplicatedTopics")
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

type AllApiApiGetMsgVpnRestDeliveryPointRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	restDeliveryPointName string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRequest) Select_(select_ []string) AllApiApiGetMsgVpnRestDeliveryPointRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRequest) Execute() (MsgVpnRestDeliveryPointResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRestDeliveryPointRequest
*/
func (a *AllApiService) GetMsgVpnRestDeliveryPoint(ctx _context.Context, msgVpnName string, restDeliveryPointName string) AllApiApiGetMsgVpnRestDeliveryPointRequest {
	return AllApiApiGetMsgVpnRestDeliveryPointRequest{
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
func (a *AllApiService) GetMsgVpnRestDeliveryPointExecute(r AllApiApiGetMsgVpnRestDeliveryPointRequest) (MsgVpnRestDeliveryPointResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnRestDeliveryPoint")
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

type AllApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	restDeliveryPointName string
	queueBindingName      string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest) Select_(select_ []string) AllApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest) Execute() (MsgVpnRestDeliveryPointQueueBindingResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest
*/
func (a *AllApiService) GetMsgVpnRestDeliveryPointQueueBinding(ctx _context.Context, msgVpnName string, restDeliveryPointName string, queueBindingName string) AllApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest {
	return AllApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest{
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
func (a *AllApiService) GetMsgVpnRestDeliveryPointQueueBindingExecute(r AllApiApiGetMsgVpnRestDeliveryPointQueueBindingRequest) (MsgVpnRestDeliveryPointQueueBindingResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointQueueBindingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnRestDeliveryPointQueueBinding")
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

type AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	restDeliveryPointName string
	count                 *int32
	cursor                *string
	where                 *[]string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) Count(count int32) AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) Cursor(cursor string) AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) Where(where []string) AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) Select_(select_ []string) AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) Execute() (MsgVpnRestDeliveryPointQueueBindingsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest
*/
func (a *AllApiService) GetMsgVpnRestDeliveryPointQueueBindings(ctx _context.Context, msgVpnName string, restDeliveryPointName string) AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest {
	return AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest{
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
func (a *AllApiService) GetMsgVpnRestDeliveryPointQueueBindingsExecute(r AllApiApiGetMsgVpnRestDeliveryPointQueueBindingsRequest) (MsgVpnRestDeliveryPointQueueBindingsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointQueueBindingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnRestDeliveryPointQueueBindings")
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

type AllApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	restDeliveryPointName string
	restConsumerName      string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest) Select_(select_ []string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest) Execute() (MsgVpnRestDeliveryPointRestConsumerResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest
*/
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumer(ctx _context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest {
	return AllApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest{
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
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumerExecute(r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerRequest) (MsgVpnRestDeliveryPointRestConsumerResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnRestDeliveryPointRestConsumer")
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

type AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	restDeliveryPointName string
	restConsumerName      string
	oauthJwtClaimName     string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest) Select_(select_ []string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest) Execute() (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest
*/
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx _context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, oauthJwtClaimName string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest {
	return AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest{
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
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimExecute(r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimRequest) (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim")
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

type AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	restDeliveryPointName string
	restConsumerName      string
	count                 *int32
	cursor                *string
	where                 *[]string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) Count(count int32) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) Cursor(cursor string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) Where(where []string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) Select_(select_ []string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) Execute() (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest
*/
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims(ctx _context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest {
	return AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest{
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
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsExecute(r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsRequest) (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims")
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

type AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	restDeliveryPointName string
	restConsumerName      string
	tlsTrustedCommonName  string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest) Select_(select_ []string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest) Execute() (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest
*/
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx _context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, tlsTrustedCommonName string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest {
	return AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest{
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
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameExecute(r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameRequest) (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName")
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

type AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	restDeliveryPointName string
	restConsumerName      string
	where                 *[]string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest) Where(where []string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest) Select_(select_ []string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest) Execute() (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest
*/
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames(ctx _context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest {
	return AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest{
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
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesExecute(r AllApiApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesRequest) (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames")
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

type AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest struct {
	ctx                   _context.Context
	ApiService            *AllApiService
	msgVpnName            string
	restDeliveryPointName string
	count                 *int32
	cursor                *string
	where                 *[]string
	select_               *[]string
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) Count(count int32) AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) Cursor(cursor string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) Where(where []string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) Select_(select_ []string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) Execute() (MsgVpnRestDeliveryPointRestConsumersResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest
*/
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumers(ctx _context.Context, msgVpnName string, restDeliveryPointName string) AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest {
	return AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest{
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
func (a *AllApiService) GetMsgVpnRestDeliveryPointRestConsumersExecute(r AllApiApiGetMsgVpnRestDeliveryPointRestConsumersRequest) (MsgVpnRestDeliveryPointRestConsumersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointRestConsumersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnRestDeliveryPointRestConsumers")
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

type AllApiApiGetMsgVpnRestDeliveryPointsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnRestDeliveryPointsRequest) Count(count int32) AllApiApiGetMsgVpnRestDeliveryPointsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointsRequest) Cursor(cursor string) AllApiApiGetMsgVpnRestDeliveryPointsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointsRequest) Where(where []string) AllApiApiGetMsgVpnRestDeliveryPointsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnRestDeliveryPointsRequest) Select_(select_ []string) AllApiApiGetMsgVpnRestDeliveryPointsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnRestDeliveryPointsRequest) Execute() (MsgVpnRestDeliveryPointsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnRestDeliveryPointsRequest
*/
func (a *AllApiService) GetMsgVpnRestDeliveryPoints(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnRestDeliveryPointsRequest {
	return AllApiApiGetMsgVpnRestDeliveryPointsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnRestDeliveryPointsResponse
 */
func (a *AllApiService) GetMsgVpnRestDeliveryPointsExecute(r AllApiApiGetMsgVpnRestDeliveryPointsRequest) (MsgVpnRestDeliveryPointsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnRestDeliveryPointsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnRestDeliveryPoints")
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

type AllApiApiGetMsgVpnTopicEndpointRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	topicEndpointName string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnTopicEndpointRequest) Select_(select_ []string) AllApiApiGetMsgVpnTopicEndpointRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTopicEndpointRequest) Execute() (MsgVpnTopicEndpointResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTopicEndpointRequest
*/
func (a *AllApiService) GetMsgVpnTopicEndpoint(ctx _context.Context, msgVpnName string, topicEndpointName string) AllApiApiGetMsgVpnTopicEndpointRequest {
	return AllApiApiGetMsgVpnTopicEndpointRequest{
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
func (a *AllApiService) GetMsgVpnTopicEndpointExecute(r AllApiApiGetMsgVpnTopicEndpointRequest) (MsgVpnTopicEndpointResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTopicEndpoint")
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

type AllApiApiGetMsgVpnTopicEndpointMsgRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	topicEndpointName string
	msgId             string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnTopicEndpointMsgRequest) Select_(select_ []string) AllApiApiGetMsgVpnTopicEndpointMsgRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTopicEndpointMsgRequest) Execute() (MsgVpnTopicEndpointMsgResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTopicEndpointMsgRequest
*/
func (a *AllApiService) GetMsgVpnTopicEndpointMsg(ctx _context.Context, msgVpnName string, topicEndpointName string, msgId string) AllApiApiGetMsgVpnTopicEndpointMsgRequest {
	return AllApiApiGetMsgVpnTopicEndpointMsgRequest{
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
func (a *AllApiService) GetMsgVpnTopicEndpointMsgExecute(r AllApiApiGetMsgVpnTopicEndpointMsgRequest) (MsgVpnTopicEndpointMsgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointMsgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTopicEndpointMsg")
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

type AllApiApiGetMsgVpnTopicEndpointMsgsRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	topicEndpointName string
	count             *int32
	cursor            *string
	where             *[]string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnTopicEndpointMsgsRequest) Count(count int32) AllApiApiGetMsgVpnTopicEndpointMsgsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointMsgsRequest) Cursor(cursor string) AllApiApiGetMsgVpnTopicEndpointMsgsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointMsgsRequest) Where(where []string) AllApiApiGetMsgVpnTopicEndpointMsgsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointMsgsRequest) Select_(select_ []string) AllApiApiGetMsgVpnTopicEndpointMsgsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTopicEndpointMsgsRequest) Execute() (MsgVpnTopicEndpointMsgsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTopicEndpointMsgsRequest
*/
func (a *AllApiService) GetMsgVpnTopicEndpointMsgs(ctx _context.Context, msgVpnName string, topicEndpointName string) AllApiApiGetMsgVpnTopicEndpointMsgsRequest {
	return AllApiApiGetMsgVpnTopicEndpointMsgsRequest{
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
func (a *AllApiService) GetMsgVpnTopicEndpointMsgsExecute(r AllApiApiGetMsgVpnTopicEndpointMsgsRequest) (MsgVpnTopicEndpointMsgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointMsgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTopicEndpointMsgs")
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

type AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	topicEndpointName string
	count             *int32
	cursor            *string
	where             *[]string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest) Count(count int32) AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest) Cursor(cursor string) AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest) Where(where []string) AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest) Select_(select_ []string) AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest) Execute() (MsgVpnTopicEndpointPrioritiesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest
*/
func (a *AllApiService) GetMsgVpnTopicEndpointPriorities(ctx _context.Context, msgVpnName string, topicEndpointName string) AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest {
	return AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest{
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
func (a *AllApiService) GetMsgVpnTopicEndpointPrioritiesExecute(r AllApiApiGetMsgVpnTopicEndpointPrioritiesRequest) (MsgVpnTopicEndpointPrioritiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointPrioritiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTopicEndpointPriorities")
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

type AllApiApiGetMsgVpnTopicEndpointPriorityRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	topicEndpointName string
	priority          string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnTopicEndpointPriorityRequest) Select_(select_ []string) AllApiApiGetMsgVpnTopicEndpointPriorityRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTopicEndpointPriorityRequest) Execute() (MsgVpnTopicEndpointPriorityResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTopicEndpointPriorityRequest
*/
func (a *AllApiService) GetMsgVpnTopicEndpointPriority(ctx _context.Context, msgVpnName string, topicEndpointName string, priority string) AllApiApiGetMsgVpnTopicEndpointPriorityRequest {
	return AllApiApiGetMsgVpnTopicEndpointPriorityRequest{
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
func (a *AllApiService) GetMsgVpnTopicEndpointPriorityExecute(r AllApiApiGetMsgVpnTopicEndpointPriorityRequest) (MsgVpnTopicEndpointPriorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointPriorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTopicEndpointPriority")
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

type AllApiApiGetMsgVpnTopicEndpointTemplateRequest struct {
	ctx                       _context.Context
	ApiService                *AllApiService
	msgVpnName                string
	topicEndpointTemplateName string
	select_                   *[]string
}

func (r AllApiApiGetMsgVpnTopicEndpointTemplateRequest) Select_(select_ []string) AllApiApiGetMsgVpnTopicEndpointTemplateRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTopicEndpointTemplateRequest) Execute() (MsgVpnTopicEndpointTemplateResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTopicEndpointTemplateRequest
*/
func (a *AllApiService) GetMsgVpnTopicEndpointTemplate(ctx _context.Context, msgVpnName string, topicEndpointTemplateName string) AllApiApiGetMsgVpnTopicEndpointTemplateRequest {
	return AllApiApiGetMsgVpnTopicEndpointTemplateRequest{
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
func (a *AllApiService) GetMsgVpnTopicEndpointTemplateExecute(r AllApiApiGetMsgVpnTopicEndpointTemplateRequest) (MsgVpnTopicEndpointTemplateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointTemplateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTopicEndpointTemplate")
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

type AllApiApiGetMsgVpnTopicEndpointTemplatesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnTopicEndpointTemplatesRequest) Count(count int32) AllApiApiGetMsgVpnTopicEndpointTemplatesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointTemplatesRequest) Cursor(cursor string) AllApiApiGetMsgVpnTopicEndpointTemplatesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointTemplatesRequest) Where(where []string) AllApiApiGetMsgVpnTopicEndpointTemplatesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointTemplatesRequest) Select_(select_ []string) AllApiApiGetMsgVpnTopicEndpointTemplatesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTopicEndpointTemplatesRequest) Execute() (MsgVpnTopicEndpointTemplatesResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTopicEndpointTemplatesRequest
*/
func (a *AllApiService) GetMsgVpnTopicEndpointTemplates(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnTopicEndpointTemplatesRequest {
	return AllApiApiGetMsgVpnTopicEndpointTemplatesRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointTemplatesResponse
 */
func (a *AllApiService) GetMsgVpnTopicEndpointTemplatesExecute(r AllApiApiGetMsgVpnTopicEndpointTemplatesRequest) (MsgVpnTopicEndpointTemplatesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointTemplatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTopicEndpointTemplates")
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

type AllApiApiGetMsgVpnTopicEndpointTxFlowRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	topicEndpointName string
	flowId            string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnTopicEndpointTxFlowRequest) Select_(select_ []string) AllApiApiGetMsgVpnTopicEndpointTxFlowRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTopicEndpointTxFlowRequest) Execute() (MsgVpnTopicEndpointTxFlowResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTopicEndpointTxFlowRequest
*/
func (a *AllApiService) GetMsgVpnTopicEndpointTxFlow(ctx _context.Context, msgVpnName string, topicEndpointName string, flowId string) AllApiApiGetMsgVpnTopicEndpointTxFlowRequest {
	return AllApiApiGetMsgVpnTopicEndpointTxFlowRequest{
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
func (a *AllApiService) GetMsgVpnTopicEndpointTxFlowExecute(r AllApiApiGetMsgVpnTopicEndpointTxFlowRequest) (MsgVpnTopicEndpointTxFlowResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointTxFlowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTopicEndpointTxFlow")
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

type AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	msgVpnName        string
	topicEndpointName string
	count             *int32
	cursor            *string
	where             *[]string
	select_           *[]string
}

func (r AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest) Count(count int32) AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest) Cursor(cursor string) AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest) Where(where []string) AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest) Select_(select_ []string) AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest) Execute() (MsgVpnTopicEndpointTxFlowsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest
*/
func (a *AllApiService) GetMsgVpnTopicEndpointTxFlows(ctx _context.Context, msgVpnName string, topicEndpointName string) AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest {
	return AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest{
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
func (a *AllApiService) GetMsgVpnTopicEndpointTxFlowsExecute(r AllApiApiGetMsgVpnTopicEndpointTxFlowsRequest) (MsgVpnTopicEndpointTxFlowsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointTxFlowsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTopicEndpointTxFlows")
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

type AllApiApiGetMsgVpnTopicEndpointsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnTopicEndpointsRequest) Count(count int32) AllApiApiGetMsgVpnTopicEndpointsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointsRequest) Cursor(cursor string) AllApiApiGetMsgVpnTopicEndpointsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointsRequest) Where(where []string) AllApiApiGetMsgVpnTopicEndpointsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnTopicEndpointsRequest) Select_(select_ []string) AllApiApiGetMsgVpnTopicEndpointsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTopicEndpointsRequest) Execute() (MsgVpnTopicEndpointsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTopicEndpointsRequest
*/
func (a *AllApiService) GetMsgVpnTopicEndpoints(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnTopicEndpointsRequest {
	return AllApiApiGetMsgVpnTopicEndpointsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTopicEndpointsResponse
 */
func (a *AllApiService) GetMsgVpnTopicEndpointsExecute(r AllApiApiGetMsgVpnTopicEndpointsRequest) (MsgVpnTopicEndpointsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTopicEndpointsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTopicEndpoints")
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

type AllApiApiGetMsgVpnTransactionRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	xid        string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnTransactionRequest) Select_(select_ []string) AllApiApiGetMsgVpnTransactionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTransactionRequest) Execute() (MsgVpnTransactionResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTransactionRequest
*/
func (a *AllApiService) GetMsgVpnTransaction(ctx _context.Context, msgVpnName string, xid string) AllApiApiGetMsgVpnTransactionRequest {
	return AllApiApiGetMsgVpnTransactionRequest{
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
func (a *AllApiService) GetMsgVpnTransactionExecute(r AllApiApiGetMsgVpnTransactionRequest) (MsgVpnTransactionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTransaction")
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

type AllApiApiGetMsgVpnTransactionConsumerMsgRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	xid        string
	msgId      string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnTransactionConsumerMsgRequest) Select_(select_ []string) AllApiApiGetMsgVpnTransactionConsumerMsgRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTransactionConsumerMsgRequest) Execute() (MsgVpnTransactionConsumerMsgResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTransactionConsumerMsgRequest
*/
func (a *AllApiService) GetMsgVpnTransactionConsumerMsg(ctx _context.Context, msgVpnName string, xid string, msgId string) AllApiApiGetMsgVpnTransactionConsumerMsgRequest {
	return AllApiApiGetMsgVpnTransactionConsumerMsgRequest{
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
func (a *AllApiService) GetMsgVpnTransactionConsumerMsgExecute(r AllApiApiGetMsgVpnTransactionConsumerMsgRequest) (MsgVpnTransactionConsumerMsgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionConsumerMsgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTransactionConsumerMsg")
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

type AllApiApiGetMsgVpnTransactionConsumerMsgsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	xid        string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnTransactionConsumerMsgsRequest) Count(count int32) AllApiApiGetMsgVpnTransactionConsumerMsgsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnTransactionConsumerMsgsRequest) Cursor(cursor string) AllApiApiGetMsgVpnTransactionConsumerMsgsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnTransactionConsumerMsgsRequest) Where(where []string) AllApiApiGetMsgVpnTransactionConsumerMsgsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnTransactionConsumerMsgsRequest) Select_(select_ []string) AllApiApiGetMsgVpnTransactionConsumerMsgsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTransactionConsumerMsgsRequest) Execute() (MsgVpnTransactionConsumerMsgsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTransactionConsumerMsgsRequest
*/
func (a *AllApiService) GetMsgVpnTransactionConsumerMsgs(ctx _context.Context, msgVpnName string, xid string) AllApiApiGetMsgVpnTransactionConsumerMsgsRequest {
	return AllApiApiGetMsgVpnTransactionConsumerMsgsRequest{
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
func (a *AllApiService) GetMsgVpnTransactionConsumerMsgsExecute(r AllApiApiGetMsgVpnTransactionConsumerMsgsRequest) (MsgVpnTransactionConsumerMsgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionConsumerMsgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTransactionConsumerMsgs")
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

type AllApiApiGetMsgVpnTransactionPublisherMsgRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	xid        string
	msgId      string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnTransactionPublisherMsgRequest) Select_(select_ []string) AllApiApiGetMsgVpnTransactionPublisherMsgRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTransactionPublisherMsgRequest) Execute() (MsgVpnTransactionPublisherMsgResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTransactionPublisherMsgRequest
*/
func (a *AllApiService) GetMsgVpnTransactionPublisherMsg(ctx _context.Context, msgVpnName string, xid string, msgId string) AllApiApiGetMsgVpnTransactionPublisherMsgRequest {
	return AllApiApiGetMsgVpnTransactionPublisherMsgRequest{
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
func (a *AllApiService) GetMsgVpnTransactionPublisherMsgExecute(r AllApiApiGetMsgVpnTransactionPublisherMsgRequest) (MsgVpnTransactionPublisherMsgResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionPublisherMsgResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTransactionPublisherMsg")
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

type AllApiApiGetMsgVpnTransactionPublisherMsgsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	xid        string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnTransactionPublisherMsgsRequest) Count(count int32) AllApiApiGetMsgVpnTransactionPublisherMsgsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnTransactionPublisherMsgsRequest) Cursor(cursor string) AllApiApiGetMsgVpnTransactionPublisherMsgsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnTransactionPublisherMsgsRequest) Where(where []string) AllApiApiGetMsgVpnTransactionPublisherMsgsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnTransactionPublisherMsgsRequest) Select_(select_ []string) AllApiApiGetMsgVpnTransactionPublisherMsgsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTransactionPublisherMsgsRequest) Execute() (MsgVpnTransactionPublisherMsgsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTransactionPublisherMsgsRequest
*/
func (a *AllApiService) GetMsgVpnTransactionPublisherMsgs(ctx _context.Context, msgVpnName string, xid string) AllApiApiGetMsgVpnTransactionPublisherMsgsRequest {
	return AllApiApiGetMsgVpnTransactionPublisherMsgsRequest{
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
func (a *AllApiService) GetMsgVpnTransactionPublisherMsgsExecute(r AllApiApiGetMsgVpnTransactionPublisherMsgsRequest) (MsgVpnTransactionPublisherMsgsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionPublisherMsgsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTransactionPublisherMsgs")
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

type AllApiApiGetMsgVpnTransactionsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	msgVpnName string
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnTransactionsRequest) Count(count int32) AllApiApiGetMsgVpnTransactionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnTransactionsRequest) Cursor(cursor string) AllApiApiGetMsgVpnTransactionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnTransactionsRequest) Where(where []string) AllApiApiGetMsgVpnTransactionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnTransactionsRequest) Select_(select_ []string) AllApiApiGetMsgVpnTransactionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnTransactionsRequest) Execute() (MsgVpnTransactionsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnTransactionsRequest
*/
func (a *AllApiService) GetMsgVpnTransactions(ctx _context.Context, msgVpnName string) AllApiApiGetMsgVpnTransactionsRequest {
	return AllApiApiGetMsgVpnTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
		msgVpnName: msgVpnName,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnTransactionsResponse
 */
func (a *AllApiService) GetMsgVpnTransactionsExecute(r AllApiApiGetMsgVpnTransactionsRequest) (MsgVpnTransactionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnTransactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpnTransactions")
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

type AllApiApiGetMsgVpnsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetMsgVpnsRequest) Count(count int32) AllApiApiGetMsgVpnsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetMsgVpnsRequest) Cursor(cursor string) AllApiApiGetMsgVpnsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetMsgVpnsRequest) Where(where []string) AllApiApiGetMsgVpnsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetMsgVpnsRequest) Select_(select_ []string) AllApiApiGetMsgVpnsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetMsgVpnsRequest) Execute() (MsgVpnsResponse, *_nethttp.Response, error) {
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
 * @return AllApiApiGetMsgVpnsRequest
*/
func (a *AllApiService) GetMsgVpns(ctx _context.Context) AllApiApiGetMsgVpnsRequest {
	return AllApiApiGetMsgVpnsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return MsgVpnsResponse
 */
func (a *AllApiService) GetMsgVpnsExecute(r AllApiApiGetMsgVpnsRequest) (MsgVpnsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MsgVpnsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetMsgVpns")
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

type AllApiApiGetSessionRequest struct {
	ctx             _context.Context
	ApiService      *AllApiService
	sessionUsername string
	sessionId       string
	select_         *[]string
}

func (r AllApiApiGetSessionRequest) Select_(select_ []string) AllApiApiGetSessionRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetSessionRequest) Execute() (SessionResponse, *_nethttp.Response, error) {
	return r.ApiService.GetSessionExecute(r)
}

/*
 * GetSession Get a Session object.
 * Get a Session object.

Administrative sessions for configuration and monitoring.


Attribute|Identifying|Deprecated
:---|:---:|:---:
sessionId|x|
sessionUsername|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.21.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sessionUsername The username used for authorization.
 * @param sessionId The unique identifier for the session.
 * @return AllApiApiGetSessionRequest
*/
func (a *AllApiService) GetSession(ctx _context.Context, sessionUsername string, sessionId string) AllApiApiGetSessionRequest {
	return AllApiApiGetSessionRequest{
		ApiService:      a,
		ctx:             ctx,
		sessionUsername: sessionUsername,
		sessionId:       sessionId,
	}
}

/*
 * Execute executes the request
 * @return SessionResponse
 */
func (a *AllApiService) GetSessionExecute(r AllApiApiGetSessionRequest) (SessionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetSession")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions/{sessionUsername},{sessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionUsername"+"}", _neturl.PathEscape(parameterToString(r.sessionUsername, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", _neturl.PathEscape(parameterToString(r.sessionId, "")), -1)

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

type AllApiApiGetSessionsRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetSessionsRequest) Count(count int32) AllApiApiGetSessionsRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetSessionsRequest) Cursor(cursor string) AllApiApiGetSessionsRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetSessionsRequest) Where(where []string) AllApiApiGetSessionsRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetSessionsRequest) Select_(select_ []string) AllApiApiGetSessionsRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetSessionsRequest) Execute() (SessionsResponse, *_nethttp.Response, error) {
	return r.ApiService.GetSessionsExecute(r)
}

/*
 * GetSessions Get a list of Session objects.
 * Get a list of Session objects.

Administrative sessions for configuration and monitoring.


Attribute|Identifying|Deprecated
:---|:---:|:---:
sessionId|x|
sessionUsername|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.21.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetSessionsRequest
*/
func (a *AllApiService) GetSessions(ctx _context.Context) AllApiApiGetSessionsRequest {
	return AllApiApiGetSessionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return SessionsResponse
 */
func (a *AllApiService) GetSessionsExecute(r AllApiApiGetSessionsRequest) (SessionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SessionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetSessions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sessions"

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

type AllApiApiGetStandardDomainCertAuthoritiesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetStandardDomainCertAuthoritiesRequest) Count(count int32) AllApiApiGetStandardDomainCertAuthoritiesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetStandardDomainCertAuthoritiesRequest) Cursor(cursor string) AllApiApiGetStandardDomainCertAuthoritiesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetStandardDomainCertAuthoritiesRequest) Where(where []string) AllApiApiGetStandardDomainCertAuthoritiesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetStandardDomainCertAuthoritiesRequest) Select_(select_ []string) AllApiApiGetStandardDomainCertAuthoritiesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetStandardDomainCertAuthoritiesRequest) Execute() (StandardDomainCertAuthoritiesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetStandardDomainCertAuthoritiesExecute(r)
}

/*
 * GetStandardDomainCertAuthorities Get a list of Standard Domain Certificate Authority objects.
 * Get a list of Standard Domain Certificate Authority objects.

Standard Certificate Authorities trusted for domain verification.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetStandardDomainCertAuthoritiesRequest
*/
func (a *AllApiService) GetStandardDomainCertAuthorities(ctx _context.Context) AllApiApiGetStandardDomainCertAuthoritiesRequest {
	return AllApiApiGetStandardDomainCertAuthoritiesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return StandardDomainCertAuthoritiesResponse
 */
func (a *AllApiService) GetStandardDomainCertAuthoritiesExecute(r AllApiApiGetStandardDomainCertAuthoritiesRequest) (StandardDomainCertAuthoritiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StandardDomainCertAuthoritiesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetStandardDomainCertAuthorities")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/standardDomainCertAuthorities"

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

type AllApiApiGetStandardDomainCertAuthorityRequest struct {
	ctx               _context.Context
	ApiService        *AllApiService
	certAuthorityName string
	select_           *[]string
}

func (r AllApiApiGetStandardDomainCertAuthorityRequest) Select_(select_ []string) AllApiApiGetStandardDomainCertAuthorityRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetStandardDomainCertAuthorityRequest) Execute() (StandardDomainCertAuthorityResponse, *_nethttp.Response, error) {
	return r.ApiService.GetStandardDomainCertAuthorityExecute(r)
}

/*
 * GetStandardDomainCertAuthority Get a Standard Domain Certificate Authority object.
 * Get a Standard Domain Certificate Authority object.

Standard Certificate Authorities trusted for domain verification.


Attribute|Identifying|Deprecated
:---|:---:|:---:
certAuthorityName|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.19.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param certAuthorityName The name of the Certificate Authority.
 * @return AllApiApiGetStandardDomainCertAuthorityRequest
*/
func (a *AllApiService) GetStandardDomainCertAuthority(ctx _context.Context, certAuthorityName string) AllApiApiGetStandardDomainCertAuthorityRequest {
	return AllApiApiGetStandardDomainCertAuthorityRequest{
		ApiService:        a,
		ctx:               ctx,
		certAuthorityName: certAuthorityName,
	}
}

/*
 * Execute executes the request
 * @return StandardDomainCertAuthorityResponse
 */
func (a *AllApiService) GetStandardDomainCertAuthorityExecute(r AllApiApiGetStandardDomainCertAuthorityRequest) (StandardDomainCertAuthorityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StandardDomainCertAuthorityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetStandardDomainCertAuthority")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/standardDomainCertAuthorities/{certAuthorityName}"
	localVarPath = strings.Replace(localVarPath, "{"+"certAuthorityName"+"}", _neturl.PathEscape(parameterToString(r.certAuthorityName, "")), -1)

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

type AllApiApiGetVirtualHostnameRequest struct {
	ctx             _context.Context
	ApiService      *AllApiService
	virtualHostname string
	select_         *[]string
}

func (r AllApiApiGetVirtualHostnameRequest) Select_(select_ []string) AllApiApiGetVirtualHostnameRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetVirtualHostnameRequest) Execute() (VirtualHostnameResponse, *_nethttp.Response, error) {
	return r.ApiService.GetVirtualHostnameExecute(r)
}

/*
 * GetVirtualHostname Get a Virtual Hostname object.
 * Get a Virtual Hostname object.

A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.

Clients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.

This mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.


Attribute|Identifying|Deprecated
:---|:---:|:---:
virtualHostname|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.17.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param virtualHostname The virtual hostname.
 * @return AllApiApiGetVirtualHostnameRequest
*/
func (a *AllApiService) GetVirtualHostname(ctx _context.Context, virtualHostname string) AllApiApiGetVirtualHostnameRequest {
	return AllApiApiGetVirtualHostnameRequest{
		ApiService:      a,
		ctx:             ctx,
		virtualHostname: virtualHostname,
	}
}

/*
 * Execute executes the request
 * @return VirtualHostnameResponse
 */
func (a *AllApiService) GetVirtualHostnameExecute(r AllApiApiGetVirtualHostnameRequest) (VirtualHostnameResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VirtualHostnameResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetVirtualHostname")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/virtualHostnames/{virtualHostname}"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualHostname"+"}", _neturl.PathEscape(parameterToString(r.virtualHostname, "")), -1)

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

type AllApiApiGetVirtualHostnamesRequest struct {
	ctx        _context.Context
	ApiService *AllApiService
	count      *int32
	cursor     *string
	where      *[]string
	select_    *[]string
}

func (r AllApiApiGetVirtualHostnamesRequest) Count(count int32) AllApiApiGetVirtualHostnamesRequest {
	r.count = &count
	return r
}
func (r AllApiApiGetVirtualHostnamesRequest) Cursor(cursor string) AllApiApiGetVirtualHostnamesRequest {
	r.cursor = &cursor
	return r
}
func (r AllApiApiGetVirtualHostnamesRequest) Where(where []string) AllApiApiGetVirtualHostnamesRequest {
	r.where = &where
	return r
}
func (r AllApiApiGetVirtualHostnamesRequest) Select_(select_ []string) AllApiApiGetVirtualHostnamesRequest {
	r.select_ = &select_
	return r
}

func (r AllApiApiGetVirtualHostnamesRequest) Execute() (VirtualHostnamesResponse, *_nethttp.Response, error) {
	return r.ApiService.GetVirtualHostnamesExecute(r)
}

/*
 * GetVirtualHostnames Get a list of Virtual Hostname objects.
 * Get a list of Virtual Hostname objects.

A Virtual Hostname is a provisioned object on a message broker that contains a Virtual Hostname to Message VPN mapping.

Clients which connect to a global (as opposed to per Message VPN) port and provides this hostname will be directed to its corresponding Message VPN. A case-insentive match is performed on the full client-provided hostname against the configured virtual-hostname.

This mechanism is only supported for hostnames provided through the Server Name Indication (SNI) extension of TLS.


Attribute|Identifying|Deprecated
:---|:---:|:---:
virtualHostname|x|



A SEMP client authorized with a minimum access scope/level of "global/read-only" is required to perform this operation.

This has been available since 2.17.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return AllApiApiGetVirtualHostnamesRequest
*/
func (a *AllApiService) GetVirtualHostnames(ctx _context.Context) AllApiApiGetVirtualHostnamesRequest {
	return AllApiApiGetVirtualHostnamesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return VirtualHostnamesResponse
 */
func (a *AllApiService) GetVirtualHostnamesExecute(r AllApiApiGetVirtualHostnamesRequest) (VirtualHostnamesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VirtualHostnamesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllApiService.GetVirtualHostnames")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/virtualHostnames"

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
