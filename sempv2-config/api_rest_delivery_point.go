/*
 * SEMP (Solace Element Management Protocol)
 *
 * SEMP (starting in `v2`, see note 1) is a RESTful API for configuring, monitoring, and administering a Solace PubSub+ broker.  SEMP uses URIs to address manageable **resources** of the Solace PubSub+ broker. Resources are individual **objects**, **collections** of objects, or (exclusively in the action API) **actions**. This document applies to the following API:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Configuration|/SEMP/v2/config|Reading and writing config state|See note 2    The following APIs are also available:   API|Base Path|Purpose|Comments :---|:---|:---|:--- Action|/SEMP/v2/action|Performing actions|See note 2 Monitoring|/SEMP/v2/monitor|Querying operational parameters|See note 2    Resources are always nouns, with individual objects being singular and collections being plural.  Objects within a collection are identified by an `obj-id`, which follows the collection name with the form `collection-name/obj-id`.  Actions within an object are identified by an `action-id`, which follows the object name with the form `obj-id/action-id`.  Some examples:  ``` /SEMP/v2/config/msgVpns                        ; MsgVpn collection /SEMP/v2/config/msgVpns/a                      ; MsgVpn object named \"a\" /SEMP/v2/config/msgVpns/a/queues               ; Queue collection in MsgVpn \"a\" /SEMP/v2/config/msgVpns/a/queues/b             ; Queue object named \"b\" in MsgVpn \"a\" /SEMP/v2/action/msgVpns/a/queues/b/startReplay ; Action that starts a replay on Queue \"b\" in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients             ; Client collection in MsgVpn \"a\" /SEMP/v2/monitor/msgVpns/a/clients/c           ; Client object named \"c\" in MsgVpn \"a\" ```  ## Collection Resources  Collections are unordered lists of objects (unless described as otherwise), and are described by JSON arrays. Each item in the array represents an object in the same manner as the individual object would normally be represented. In the configuration API, the creation of a new object is done through its collection resource.  ## Object and Action Resources  Objects are composed of attributes, actions, collections, and other objects. They are described by JSON objects as name/value pairs. The collections and actions of an object are not contained directly in the object's JSON content; rather the content includes an attribute containing a URI which points to the collections and actions. These contained resources must be managed through this URI. At a minimum, every object has one or more identifying attributes, and its own `uri` attribute which contains the URI pointing to itself.  Actions are also composed of attributes, and are described by JSON objects as name/value pairs. Unlike objects, however, they are not members of a collection and cannot be retrieved, only performed. Actions only exist in the action API.  Attributes in an object or action may have any combination of the following properties:   Property|Meaning|Comments :---|:---|:--- Identifying|Attribute is involved in unique identification of the object, and appears in its URI| Required|Attribute must be provided in the request| Read-Only|Attribute can only be read, not written.|See note 3 Write-Only|Attribute can only be written, not read, unless the attribute is also opaque|See the documentation for the opaque property Requires-Disable|Attribute can only be changed when object is disabled| Deprecated|Attribute is deprecated, and will disappear in the next SEMP version| Opaque|Attribute can be set or retrieved in opaque form when the `opaquePassword` query parameter is present|See the `opaquePassword` query parameter documentation    In some requests, certain attributes may only be provided in certain combinations with other attributes:   Relationship|Meaning :---|:--- Requires|Attribute may only be changed by a request if a particular attribute or combination of attributes is also provided in the request Conflicts|Attribute may only be provided in a request if a particular attribute or combination of attributes is not also provided in the request    In the monitoring API, any non-identifying attribute may not be returned in a GET.  ## HTTP Methods  The following HTTP methods manipulate resources in accordance with these general principles. Note that some methods are only used in certain APIs:   Method|Resource|Meaning|Request Body|Response Body|Missing Request Attributes :---|:---|:---|:---|:---|:--- POST|Collection|Create object|Initial attribute values|Object attributes and metadata|Set to default PUT|Object|Create or replace object (see note 5)|New attribute values|Object attributes and metadata|Set to default, with certain exceptions (see note 4) PUT|Action|Performs action|Action arguments|Action metadata|N/A PATCH|Object|Update object|New attribute values|Object attributes and metadata|unchanged DELETE|Object|Delete object|Empty|Object metadata|N/A GET|Object|Get object|Empty|Object attributes and metadata|N/A GET|Collection|Get collection|Empty|Object attributes and collection metadata|N/A    ## Common Query Parameters  The following are some common query parameters that are supported by many method/URI combinations. Individual URIs may document additional parameters. Note that multiple query parameters can be used together in a single URI, separated by the ampersand character. For example:  ``` ; Request for the MsgVpns collection using two hypothetical query parameters ; \"q1\" and \"q2\" with values \"val1\" and \"val2\" respectively /SEMP/v2/config/msgVpns?q1=val1&q2=val2 ```  ### select  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. Use this query parameter to limit the size of the returned data for each returned object, return only those fields that are desired, or exclude fields that are not desired.  The value of `select` is a comma-separated list of attribute names. If the list contains attribute names that are not prefaced by `-`, only those attributes are included in the response. If the list contains attribute names that are prefaced by `-`, those attributes are excluded from the response. If the list contains both types, then the difference of the first set of attributes and the second set of attributes is returned. If the list is empty (i.e. `select=`), no attributes are returned.  All attributes that are prefaced by `-` must follow all attributes that are not prefaced by `-`. In addition, each attribute name in the list must match at least one attribute in the object.  Names may include the `*` wildcard (zero or more characters). Nested attribute names are supported using periods (e.g. `parentName.childName`).  Some examples:  ``` ; List of all MsgVpn names /SEMP/v2/config/msgVpns?select=msgVpnName ; List of all MsgVpn and their attributes except for their names /SEMP/v2/config/msgVpns?select=-msgVpnName ; Authentication attributes of MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance?select=authentication* ; All attributes of MsgVpn \"finance\" except for authentication attributes /SEMP/v2/config/msgVpns/finance?select=-authentication* ; Access related attributes of Queue \"orderQ\" of MsgVpn \"finance\" /SEMP/v2/config/msgVpns/finance/queues/orderQ?select=owner,permission ```  ### where  Include in the response only objects where certain conditions are true. Use this query parameter to limit which objects are returned to those whose attribute values meet the given conditions.  The value of `where` is a comma-separated list of expressions. All expressions must be true for the object to be included in the response. Each expression takes the form:  ``` expression  = attribute-name OP value OP          = '==' | '!=' | '&lt;' | '&gt;' | '&lt;=' | '&gt;=' ```  `value` may be a number, string, `true`, or `false`, as appropriate for the type of `attribute-name`. Greater-than and less-than comparisons only work for numbers. A `*` in a string `value` is interpreted as a wildcard (zero or more characters). Some examples:  ``` ; Only enabled MsgVpns /SEMP/v2/config/msgVpns?where=enabled==true ; Only MsgVpns using basic non-LDAP authentication /SEMP/v2/config/msgVpns?where=authenticationBasicEnabled==true,authenticationBasicType!=ldap ; Only MsgVpns that allow more than 100 client connections /SEMP/v2/config/msgVpns?where=maxConnectionCount>100 ; Only MsgVpns with msgVpnName starting with \"B\": /SEMP/v2/config/msgVpns?where=msgVpnName==B* ```  ### count  Limit the count of objects in the response. This can be useful to limit the size of the response for large collections. The minimum value for `count` is `1` and the default is `10`. There is also a per-collection maximum value to limit request handling time. For example:  ``` ; Up to 25 MsgVpns /SEMP/v2/config/msgVpns?count=25 ```  ### cursor  The cursor, or position, for the next page of objects. Cursors are opaque data that should not be created or interpreted by SEMP clients, and should only be used as described below.  When a request is made for a collection and there may be additional objects available for retrieval that are not included in the initial response, the response will include a `cursorQuery` field containing a cursor. The value of this field can be specified in the `cursor` query parameter of a subsequent request to retrieve the next page of objects. For convenience, an appropriate URI is constructed automatically by the broker and included in the `nextPageUri` field of the response. This URI can be used directly to retrieve the next page of objects.  ### opaquePassword  Attributes with the opaque property are also write-only and so cannot normally be retrieved in a GET. However, when a password is provided in the `opaquePassword` query parameter, attributes with the opaque property are retrieved in a GET in opaque form, encrypted with this password. The query parameter can also be used on a POST, PATCH, or PUT to set opaque attributes using opaque attribute values retrieved in a GET, so long as:  1. the same password that was used to retrieve the opaque attribute values is provided; and  2. the broker to which the request is being sent has the same major and minor SEMP version as the broker that produced the opaque attribute values.  The password provided in the query parameter must be a minimum of 8 characters and a maximum of 128 characters.  The query parameter can only be used in the configuration API, and only over HTTPS.  ## Authentication  When a client makes its first SEMPv2 request, it must supply a username and password using HTTP Basic authentication.  If authentication is successful, the broker returns a cookie containing a session key. The client can omit the username and password from subsequent requests, because the broker now uses the session cookie for authentication instead. When the session expires or is deleted, the client must provide the username and password again, and the broker creates a new session.  There are a limited number of session slots available on the broker. The broker returns 529 No SEMP Session Available if it is not able to allocate a session. For this reason, all clients that use SEMPv2 should support cookies.  If certain attributes—such as a user's password—are changed, the broker automatically deletes the affected sessions. These attributes are documented below. However, changes in external user configuration data stored on a RADIUS or LDAP server do not trigger the broker to delete the associated session(s), therefore you must do this manually, if required.  A client can retrieve its current session information using the /about/user endpoint, delete its own session using the /about/user/logout endpoint, and manage all sessions using the /sessions endpoint.  ## Help  Visit [our website](https://solace.com) to learn more about Solace.  You can also download the SEMP API specifications by clicking [here](https://solace.com/downloads/).  If you need additional support, please contact us at [support@solace.com](mailto:support@solace.com).  ## Notes  Note|Description :---:|:--- 1|This specification defines SEMP starting in \"v2\", and not the original SEMP \"v1\" interface. Request and response formats between \"v1\" and \"v2\" are entirely incompatible, although both protocols share a common port configuration on the Solace PubSub+ broker. They are differentiated by the initial portion of the URI path, one of either \"/SEMP/\" or \"/SEMP/v2/\" 2|This API is partially implemented. Only a subset of all objects are available. 3|Read-only attributes may appear in POST and PUT/PATCH requests. However, if a read-only attribute is not marked as identifying, it will be ignored during a PUT/PATCH. 4|On a PUT, if the SEMP user is not authorized to modify the attribute, its value is left unchanged rather than set to default. In addition, the values of write-only attributes are not set to their defaults on a PUT, except in the following two cases: there is a mutual requires relationship with another non-write-only attribute, both attributes are absent from the request, and the non-write-only attribute is not currently set to its default value; or the attribute is also opaque and the `opaquePassword` query parameter is provided in the request. 5|On a PUT, if the object does not exist, it is created first.
 *
 * API version: 2.21
 * Contact: support@solace.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type RestDeliveryPointApiService service

/*
RestDeliveryPointApiService Create a REST Delivery Point object.
Create a REST Delivery Point object. Any attribute missing from the request will be set to its default value.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| restDeliveryPointName|x|x||||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The REST Delivery Point object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param optional nil or *RestDeliveryPointApiCreateMsgVpnRestDeliveryPointOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointResponse
*/

type RestDeliveryPointApiCreateMsgVpnRestDeliveryPointOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) CreateMsgVpnRestDeliveryPoint(ctx context.Context, body MsgVpnRestDeliveryPoint, msgVpnName string, localVarOptionals *RestDeliveryPointApiCreateMsgVpnRestDeliveryPointOpts) (MsgVpnRestDeliveryPointResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Create a Queue Binding object.
Create a Queue Binding object. Any attribute missing from the request will be set to its default value.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| queueBindingName|x|x|||| restDeliveryPointName|x||x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The Queue Binding object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param optional nil or *RestDeliveryPointApiCreateMsgVpnRestDeliveryPointQueueBindingOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointQueueBindingResponse
*/

type RestDeliveryPointApiCreateMsgVpnRestDeliveryPointQueueBindingOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) CreateMsgVpnRestDeliveryPointQueueBinding(ctx context.Context, body MsgVpnRestDeliveryPointQueueBinding, msgVpnName string, restDeliveryPointName string, localVarOptionals *RestDeliveryPointApiCreateMsgVpnRestDeliveryPointQueueBindingOpts) (MsgVpnRestDeliveryPointQueueBindingResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointQueueBindingResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointQueueBindingResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Create a REST Consumer object.
Create a REST Consumer object. Any attribute missing from the request will be set to its default value.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationClientCertContent||||x||x authenticationClientCertPassword||||x|| authenticationHttpBasicPassword||||x||x authenticationHttpHeaderValue||||x||x authenticationOauthClientSecret||||x||x authenticationOauthJwtSecretKey||||x||x msgVpnName|x||x||| restConsumerName|x|x|||| restDeliveryPointName|x||x|||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationClientCertPassword|authenticationClientCertContent| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The REST Consumer object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param optional nil or *RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumerResponse
*/

type RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) CreateMsgVpnRestDeliveryPointRestConsumer(ctx context.Context, body MsgVpnRestDeliveryPointRestConsumer, msgVpnName string, restDeliveryPointName string, localVarOptionals *RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerOpts) (MsgVpnRestDeliveryPointRestConsumerResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumerResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Create a Claim object.
Create a Claim object. Any attribute missing from the request will be set to its default value.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||| oauthJwtClaimName|x|x|||| oauthJwtClaimValue||x|||| restConsumerName|x||x||| restDeliveryPointName|x||x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.21.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The Claim object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param optional nil or *RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
*/

type RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) CreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx context.Context, body MsgVpnRestDeliveryPointRestConsumerOauthJwtClaim, msgVpnName string, restDeliveryPointName string, restConsumerName string, localVarOptionals *RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts) (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Create a Trusted Common Name object.
Create a Trusted Common Name object. Any attribute missing from the request will be set to its default value.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer&#x27;s server certificate.   Attribute|Identifying|Required|Read-Only|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x||x||x| restConsumerName|x||x||x| restDeliveryPointName|x||x||x| tlsTrustedCommonName|x|x|||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The Trusted Common Name object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param optional nil or *RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
*/

type RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) CreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx context.Context, body MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName, msgVpnName string, restDeliveryPointName string, restConsumerName string, localVarOptionals *RestDeliveryPointApiCreateMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts) (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Delete a REST Delivery Point object.
Delete a REST Delivery Point object.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
@return SempMetaOnlyResponse
*/
func (a *RestDeliveryPointApiService) DeleteMsgVpnRestDeliveryPoint(ctx context.Context, msgVpnName string, restDeliveryPointName string) (SempMetaOnlyResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SempMetaOnlyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Delete a Queue Binding object.
Delete a Queue Binding object.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param queueBindingName The name of a queue in the Message VPN.
@return SempMetaOnlyResponse
*/
func (a *RestDeliveryPointApiService) DeleteMsgVpnRestDeliveryPointQueueBinding(ctx context.Context, msgVpnName string, restDeliveryPointName string, queueBindingName string) (SempMetaOnlyResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SempMetaOnlyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueBindingName"+"}", fmt.Sprintf("%v", queueBindingName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Delete a REST Consumer object.
Delete a REST Consumer object.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
@return SempMetaOnlyResponse
*/
func (a *RestDeliveryPointApiService) DeleteMsgVpnRestDeliveryPointRestConsumer(ctx context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string) (SempMetaOnlyResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SempMetaOnlyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Delete a Claim object.
Delete a Claim object.  A Claim is added to the JWT sent to the OAuth token request endpoint.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.21.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param oauthJwtClaimName The name of the additional claim. Cannot be \&quot;exp\&quot;, \&quot;iat\&quot;, or \&quot;jti\&quot;.
@return SempMetaOnlyResponse
*/
func (a *RestDeliveryPointApiService) DeleteMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, oauthJwtClaimName string) (SempMetaOnlyResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SempMetaOnlyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"oauthJwtClaimName"+"}", fmt.Sprintf("%v", oauthJwtClaimName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Delete a Trusted Common Name object.
Delete a Trusted Common Name object.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer&#x27;s server certificate.  A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param tlsTrustedCommonName The expected trusted common name of the remote certificate.
@return SempMetaOnlyResponse
*/
func (a *RestDeliveryPointApiService) DeleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, tlsTrustedCommonName string) (SempMetaOnlyResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue SempMetaOnlyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tlsTrustedCommonName"+"}", fmt.Sprintf("%v", tlsTrustedCommonName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Get a REST Delivery Point object.
Get a REST Delivery Point object.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param optional nil or *RestDeliveryPointApiGetMsgVpnRestDeliveryPointOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointResponse
*/

type RestDeliveryPointApiGetMsgVpnRestDeliveryPointOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) GetMsgVpnRestDeliveryPoint(ctx context.Context, msgVpnName string, restDeliveryPointName string, localVarOptionals *RestDeliveryPointApiGetMsgVpnRestDeliveryPointOpts) (MsgVpnRestDeliveryPointResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Get a Queue Binding object.
Get a Queue Binding object.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param queueBindingName The name of a queue in the Message VPN.
 * @param optional nil or *RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointQueueBindingResponse
*/

type RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) GetMsgVpnRestDeliveryPointQueueBinding(ctx context.Context, msgVpnName string, restDeliveryPointName string, queueBindingName string, localVarOptionals *RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingOpts) (MsgVpnRestDeliveryPointQueueBindingResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointQueueBindingResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueBindingName"+"}", fmt.Sprintf("%v", queueBindingName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointQueueBindingResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Get a list of Queue Binding objects.
Get a list of Queue Binding objects.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| queueBindingName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param optional nil or *RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingsOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointQueueBindingsResponse
*/

type RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingsOpts struct {
	Count          optional.Int32
	Cursor         optional.String
	OpaquePassword optional.String
	Where          optional.Interface
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) GetMsgVpnRestDeliveryPointQueueBindings(ctx context.Context, msgVpnName string, restDeliveryPointName string, localVarOptionals *RestDeliveryPointApiGetMsgVpnRestDeliveryPointQueueBindingsOpts) (MsgVpnRestDeliveryPointQueueBindingsResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointQueueBindingsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointQueueBindingsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Get a REST Consumer object.
Get a REST Consumer object.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationClientCertContent||x||x authenticationClientCertPassword||x|| authenticationHttpBasicPassword||x||x authenticationHttpHeaderValue||x||x authenticationOauthClientSecret||x||x authenticationOauthJwtSecretKey||x||x msgVpnName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param optional nil or *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumerResponse
*/

type RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) GetMsgVpnRestDeliveryPointRestConsumer(ctx context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, localVarOptionals *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOpts) (MsgVpnRestDeliveryPointRestConsumerResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumerResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Get a Claim object.
Get a Claim object.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthJwtClaimName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.21.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param oauthJwtClaimName The name of the additional claim. Cannot be \&quot;exp\&quot;, \&quot;iat\&quot;, or \&quot;jti\&quot;.
 * @param optional nil or *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
*/

type RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaim(ctx context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, oauthJwtClaimName string, localVarOptionals *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimOpts) (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims/{oauthJwtClaimName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"oauthJwtClaimName"+"}", fmt.Sprintf("%v", oauthJwtClaimName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Get a list of Claim objects.
Get a list of Claim objects.  A Claim is added to the JWT sent to the OAuth token request endpoint.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| oauthJwtClaimName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.21.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param optional nil or *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse
*/

type RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsOpts struct {
	Count          optional.Int32
	Cursor         optional.String
	OpaquePassword optional.String
	Where          optional.Interface
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) GetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaims(ctx context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, localVarOptionals *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsOpts) (MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/oauthJwtClaims"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumerOauthJwtClaimsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Get a Trusted Common Name object.
Get a Trusted Common Name object.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer&#x27;s server certificate.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||x| restConsumerName|x||x| restDeliveryPointName|x||x| tlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param tlsTrustedCommonName The expected trusted common name of the remote certificate.
 * @param optional nil or *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
*/

type RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName(ctx context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, tlsTrustedCommonName string, localVarOptionals *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOpts) (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tlsTrustedCommonName"+"}", fmt.Sprintf("%v", tlsTrustedCommonName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Get a list of Trusted Common Name objects.
Get a list of Trusted Common Name objects.  The Trusted Common Names for the REST Consumer are used by encrypted transports to verify the name in the certificate presented by the remote REST consumer. They must include the common name of the remote REST consumer&#x27;s server certificate.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||x| restConsumerName|x||x| restDeliveryPointName|x||x| tlsTrustedCommonName|x||x|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been deprecated since (will be deprecated in next SEMP version). Common Name validation has been replaced by Server Certificate Name validation.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param optional nil or *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse
*/

type RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesOpts struct {
	OpaquePassword optional.String
	Where          optional.Interface
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) GetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNames(ctx context.Context, msgVpnName string, restDeliveryPointName string, restConsumerName string, localVarOptionals *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesOpts) (MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNamesResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Get a list of REST Consumer objects.
Get a list of REST Consumer objects.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: authenticationClientCertContent||x||x authenticationClientCertPassword||x|| authenticationHttpBasicPassword||x||x authenticationHttpHeaderValue||x||x authenticationOauthClientSecret||x||x authenticationOauthJwtSecretKey||x||x msgVpnName|x||| restConsumerName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param optional nil or *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumersOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumersResponse
*/

type RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumersOpts struct {
	Count          optional.Int32
	Cursor         optional.String
	OpaquePassword optional.String
	Where          optional.Interface
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) GetMsgVpnRestDeliveryPointRestConsumers(ctx context.Context, msgVpnName string, restDeliveryPointName string, localVarOptionals *RestDeliveryPointApiGetMsgVpnRestDeliveryPointRestConsumersOpts) (MsgVpnRestDeliveryPointRestConsumersResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumersResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Get a list of REST Delivery Point objects.
Get a list of REST Delivery Point objects.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Write-Only|Deprecated|Opaque :---|:---:|:---:|:---:|:---: msgVpnName|x||| restDeliveryPointName|x|||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-only\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param msgVpnName The name of the Message VPN.
 * @param optional nil or *RestDeliveryPointApiGetMsgVpnRestDeliveryPointsOpts - Optional Parameters:
     * @param "Count" (optional.Int32) -  Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter.
     * @param "Cursor" (optional.String) -  The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter.
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Where" (optional.Interface of []string) -  Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointsResponse
*/

type RestDeliveryPointApiGetMsgVpnRestDeliveryPointsOpts struct {
	Count          optional.Int32
	Cursor         optional.String
	OpaquePassword optional.String
	Where          optional.Interface
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) GetMsgVpnRestDeliveryPoints(ctx context.Context, msgVpnName string, localVarOptionals *RestDeliveryPointApiGetMsgVpnRestDeliveryPointsOpts) (MsgVpnRestDeliveryPointsResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Where.IsSet() {
		localVarQueryParams.Add("where", parameterToString(localVarOptionals.Where.Value(), "csv"))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Replace a REST Delivery Point object.
Replace a REST Delivery Point object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: clientProfileName||||x|| msgVpnName|x|x|||| restDeliveryPointName|x|x||||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The REST Delivery Point object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param optional nil or *RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointResponse
*/

type RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) ReplaceMsgVpnRestDeliveryPoint(ctx context.Context, body MsgVpnRestDeliveryPoint, msgVpnName string, restDeliveryPointName string, localVarOptionals *RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointOpts) (MsgVpnRestDeliveryPointResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Replace a Queue Binding object.
Replace a Queue Binding object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x|||| queueBindingName|x|x|||| restDeliveryPointName|x|x||||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The Queue Binding object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param queueBindingName The name of a queue in the Message VPN.
 * @param optional nil or *RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointQueueBindingOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointQueueBindingResponse
*/

type RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointQueueBindingOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) ReplaceMsgVpnRestDeliveryPointQueueBinding(ctx context.Context, body MsgVpnRestDeliveryPointQueueBinding, msgVpnName string, restDeliveryPointName string, queueBindingName string, localVarOptionals *RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointQueueBindingOpts) (MsgVpnRestDeliveryPointQueueBindingResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointQueueBindingResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueBindingName"+"}", fmt.Sprintf("%v", queueBindingName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointQueueBindingResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Replace a REST Consumer object.
Replace a REST Consumer object. Any attribute missing from the request will be set to its default value, subject to the exceptions in note 4.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationClientCertContent|||x|x||x authenticationClientCertPassword|||x|x|| authenticationHttpBasicPassword|||x|x||x authenticationHttpBasicUsername||||x|| authenticationHttpHeaderValue|||x|||x authenticationOauthClientId||||x|| authenticationOauthClientScope||||x|| authenticationOauthClientSecret|||x|x||x authenticationOauthClientTokenEndpoint||||x|| authenticationOauthJwtSecretKey|||x|x||x authenticationOauthJwtTokenEndpoint||||x|| authenticationScheme||||x|| msgVpnName|x|x|||| outgoingConnectionCount||||x|| remoteHost||||x|| remotePort||||x|| restConsumerName|x|x|||| restDeliveryPointName|x|x|||| tlsCipherSuiteList||||x|| tlsEnabled||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationClientCertPassword|authenticationClientCertContent| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The REST Consumer object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param optional nil or *RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointRestConsumerOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumerResponse
*/

type RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointRestConsumerOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) ReplaceMsgVpnRestDeliveryPointRestConsumer(ctx context.Context, body MsgVpnRestDeliveryPointRestConsumer, msgVpnName string, restDeliveryPointName string, restConsumerName string, localVarOptionals *RestDeliveryPointApiReplaceMsgVpnRestDeliveryPointRestConsumerOpts) (MsgVpnRestDeliveryPointRestConsumerResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Put")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumerResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Update a REST Delivery Point object.
Update a REST Delivery Point object. Any attribute missing from the request will be left unchanged.  A REST Delivery Point manages delivery of messages from queues to a named list of REST Consumers.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: clientProfileName||||x|| msgVpnName|x|x|||| restDeliveryPointName|x|x||||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The REST Delivery Point object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param optional nil or *RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointResponse
*/

type RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) UpdateMsgVpnRestDeliveryPoint(ctx context.Context, body MsgVpnRestDeliveryPoint, msgVpnName string, restDeliveryPointName string, localVarOptionals *RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointOpts) (MsgVpnRestDeliveryPointResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Patch")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Update a Queue Binding object.
Update a Queue Binding object. Any attribute missing from the request will be left unchanged.  A Queue Binding for a REST Delivery Point attracts messages to be delivered to REST consumers. If the queue does not exist it can be created subsequently, and once the queue is operational the broker performs the queue binding. Removing the queue binding does not delete the queue itself. Similarly, removing the queue does not remove the queue binding, which fails until the queue is recreated or the queue binding is deleted.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: msgVpnName|x|x|||| queueBindingName|x|x|||| restDeliveryPointName|x|x||||    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The Queue Binding object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param queueBindingName The name of a queue in the Message VPN.
 * @param optional nil or *RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointQueueBindingOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointQueueBindingResponse
*/

type RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointQueueBindingOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) UpdateMsgVpnRestDeliveryPointQueueBinding(ctx context.Context, body MsgVpnRestDeliveryPointQueueBinding, msgVpnName string, restDeliveryPointName string, queueBindingName string, localVarOptionals *RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointQueueBindingOpts) (MsgVpnRestDeliveryPointQueueBindingResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Patch")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointQueueBindingResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings/{queueBindingName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueBindingName"+"}", fmt.Sprintf("%v", queueBindingName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointQueueBindingResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
RestDeliveryPointApiService Update a REST Consumer object.
Update a REST Consumer object. Any attribute missing from the request will be left unchanged.  REST Consumer objects establish HTTP connectivity to REST consumer applications who wish to receive messages from a broker.   Attribute|Identifying|Read-Only|Write-Only|Requires-Disable|Deprecated|Opaque :---|:---:|:---:|:---:|:---:|:---:|:---: authenticationClientCertContent|||x|x||x authenticationClientCertPassword|||x|x|| authenticationHttpBasicPassword|||x|x||x authenticationHttpBasicUsername||||x|| authenticationHttpHeaderValue|||x|||x authenticationOauthClientId||||x|| authenticationOauthClientScope||||x|| authenticationOauthClientSecret|||x|x||x authenticationOauthClientTokenEndpoint||||x|| authenticationOauthJwtSecretKey|||x|x||x authenticationOauthJwtTokenEndpoint||||x|| authenticationScheme||||x|| msgVpnName|x|x|||| outgoingConnectionCount||||x|| remoteHost||||x|| remotePort||||x|| restConsumerName|x|x|||| restDeliveryPointName|x|x|||| tlsCipherSuiteList||||x|| tlsEnabled||||x||    The following attributes in the request may only be provided in certain combinations with other attributes:   Class|Attribute|Requires|Conflicts :---|:---|:---|:--- MsgVpnRestDeliveryPointRestConsumer|authenticationClientCertPassword|authenticationClientCertContent| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicPassword|authenticationHttpBasicUsername| MsgVpnRestDeliveryPointRestConsumer|authenticationHttpBasicUsername|authenticationHttpBasicPassword| MsgVpnRestDeliveryPointRestConsumer|remotePort|tlsEnabled| MsgVpnRestDeliveryPointRestConsumer|tlsEnabled|remotePort|    A SEMP client authorized with a minimum access scope/level of \&quot;vpn/read-write\&quot; is required to perform this operation.  This has been available since 2.0.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body The REST Consumer object&#x27;s attributes.
 * @param msgVpnName The name of the Message VPN.
 * @param restDeliveryPointName The name of the REST Delivery Point.
 * @param restConsumerName The name of the REST Consumer.
 * @param optional nil or *RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointRestConsumerOpts - Optional Parameters:
     * @param "OpaquePassword" (optional.String) -  Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter.
     * @param "Select_" (optional.Interface of []string) -  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter.
@return MsgVpnRestDeliveryPointRestConsumerResponse
*/

type RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointRestConsumerOpts struct {
	OpaquePassword optional.String
	Select_        optional.Interface
}

func (a *RestDeliveryPointApiService) UpdateMsgVpnRestDeliveryPointRestConsumer(ctx context.Context, body MsgVpnRestDeliveryPointRestConsumer, msgVpnName string, restDeliveryPointName string, restConsumerName string, localVarOptionals *RestDeliveryPointApiUpdateMsgVpnRestDeliveryPointRestConsumerOpts) (MsgVpnRestDeliveryPointRestConsumerResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Patch")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue MsgVpnRestDeliveryPointRestConsumerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}"
	localVarPath = strings.Replace(localVarPath, "{"+"msgVpnName"+"}", fmt.Sprintf("%v", msgVpnName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restDeliveryPointName"+"}", fmt.Sprintf("%v", restDeliveryPointName), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restConsumerName"+"}", fmt.Sprintf("%v", restConsumerName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.OpaquePassword.IsSet() {
		localVarQueryParams.Add("opaquePassword", parameterToString(localVarOptionals.OpaquePassword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Select_.IsSet() {
		localVarQueryParams.Add("select", parameterToString(localVarOptionals.Select_.Value(), "csv"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v MsgVpnRestDeliveryPointRestConsumerResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v SempMetaOnlyResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
