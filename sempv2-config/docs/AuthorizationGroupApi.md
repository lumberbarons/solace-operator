# \AuthorizationGroupApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnAuthorizationGroup**](AuthorizationGroupApi.md#CreateMsgVpnAuthorizationGroup) | **Post** /msgVpns/{msgVpnName}/authorizationGroups | Create an LDAP Authorization Group object.
[**DeleteMsgVpnAuthorizationGroup**](AuthorizationGroupApi.md#DeleteMsgVpnAuthorizationGroup) | **Delete** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Delete an LDAP Authorization Group object.
[**GetMsgVpnAuthorizationGroup**](AuthorizationGroupApi.md#GetMsgVpnAuthorizationGroup) | **Get** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Get an LDAP Authorization Group object.
[**GetMsgVpnAuthorizationGroups**](AuthorizationGroupApi.md#GetMsgVpnAuthorizationGroups) | **Get** /msgVpns/{msgVpnName}/authorizationGroups | Get a list of LDAP Authorization Group objects.
[**ReplaceMsgVpnAuthorizationGroup**](AuthorizationGroupApi.md#ReplaceMsgVpnAuthorizationGroup) | **Put** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Replace an LDAP Authorization Group object.
[**UpdateMsgVpnAuthorizationGroup**](AuthorizationGroupApi.md#UpdateMsgVpnAuthorizationGroup) | **Patch** /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName} | Update an LDAP Authorization Group object.



## CreateMsgVpnAuthorizationGroup

> MsgVpnAuthorizationGroupResponse CreateMsgVpnAuthorizationGroup(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create an LDAP Authorization Group object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    body := *openapiclient.NewMsgVpnAuthorizationGroup() // MsgVpnAuthorizationGroup | The LDAP Authorization Group object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationGroupApi.CreateMsgVpnAuthorizationGroup(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationGroupApi.CreateMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnAuthorizationGroup`: MsgVpnAuthorizationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationGroupApi.CreateMsgVpnAuthorizationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnAuthorizationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnAuthorizationGroup**](MsgVpnAuthorizationGroup.md) | The LDAP Authorization Group object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnAuthorizationGroup

> SempMetaOnlyResponse DeleteMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName).Execute()

Delete an LDAP Authorization Group object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    authorizationGroupName := "authorizationGroupName_example" // string | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\\#,lab,com'.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationGroupApi.DeleteMsgVpnAuthorizationGroup(context.Background(), msgVpnName, authorizationGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationGroupApi.DeleteMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnAuthorizationGroup`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationGroupApi.DeleteMsgVpnAuthorizationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**authorizationGroupName** | **string** | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as &#39;#&#39;, &#39;+&#39;, &#39;;&#39;, &#39;&#x3D;&#39; as the value of the group name returned from the LDAP server might prepend those characters with &#39;\\&#39;. For example a group name called &#39;test#,lab,com&#39; will be returned from the LDAP server as &#39;test\\#,lab,com&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnAuthorizationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SempMetaOnlyResponse**](SempMetaOnlyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAuthorizationGroup

> MsgVpnAuthorizationGroupResponse GetMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an LDAP Authorization Group object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    authorizationGroupName := "authorizationGroupName_example" // string | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\\#,lab,com'.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationGroupApi.GetMsgVpnAuthorizationGroup(context.Background(), msgVpnName, authorizationGroupName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationGroupApi.GetMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAuthorizationGroup`: MsgVpnAuthorizationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationGroupApi.GetMsgVpnAuthorizationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**authorizationGroupName** | **string** | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as &#39;#&#39;, &#39;+&#39;, &#39;;&#39;, &#39;&#x3D;&#39; as the value of the group name returned from the LDAP server might prepend those characters with &#39;\\&#39;. For example a group name called &#39;test#,lab,com&#39; will be returned from the LDAP server as &#39;test\\#,lab,com&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAuthorizationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAuthorizationGroups

> MsgVpnAuthorizationGroupsResponse GetMsgVpnAuthorizationGroups(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of LDAP Authorization Group objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationGroupApi.GetMsgVpnAuthorizationGroups(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationGroupApi.GetMsgVpnAuthorizationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAuthorizationGroups`: MsgVpnAuthorizationGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationGroupApi.GetMsgVpnAuthorizationGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAuthorizationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupsResponse**](MsgVpnAuthorizationGroupsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnAuthorizationGroup

> MsgVpnAuthorizationGroupResponse ReplaceMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace an LDAP Authorization Group object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    authorizationGroupName := "authorizationGroupName_example" // string | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\\#,lab,com'.
    body := *openapiclient.NewMsgVpnAuthorizationGroup() // MsgVpnAuthorizationGroup | The LDAP Authorization Group object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationGroupApi.ReplaceMsgVpnAuthorizationGroup(context.Background(), msgVpnName, authorizationGroupName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationGroupApi.ReplaceMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnAuthorizationGroup`: MsgVpnAuthorizationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationGroupApi.ReplaceMsgVpnAuthorizationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**authorizationGroupName** | **string** | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as &#39;#&#39;, &#39;+&#39;, &#39;;&#39;, &#39;&#x3D;&#39; as the value of the group name returned from the LDAP server might prepend those characters with &#39;\\&#39;. For example a group name called &#39;test#,lab,com&#39; will be returned from the LDAP server as &#39;test\\#,lab,com&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnAuthorizationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAuthorizationGroup**](MsgVpnAuthorizationGroup.md) | The LDAP Authorization Group object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnAuthorizationGroup

> MsgVpnAuthorizationGroupResponse UpdateMsgVpnAuthorizationGroup(ctx, msgVpnName, authorizationGroupName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update an LDAP Authorization Group object.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    authorizationGroupName := "authorizationGroupName_example" // string | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as '#', '+', ';', '=' as the value of the group name returned from the LDAP server might prepend those characters with '\\'. For example a group name called 'test#,lab,com' will be returned from the LDAP server as 'test\\#,lab,com'.
    body := *openapiclient.NewMsgVpnAuthorizationGroup() // MsgVpnAuthorizationGroup | The LDAP Authorization Group object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationGroupApi.UpdateMsgVpnAuthorizationGroup(context.Background(), msgVpnName, authorizationGroupName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationGroupApi.UpdateMsgVpnAuthorizationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnAuthorizationGroup`: MsgVpnAuthorizationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationGroupApi.UpdateMsgVpnAuthorizationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**authorizationGroupName** | **string** | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as &#39;#&#39;, &#39;+&#39;, &#39;;&#39;, &#39;&#x3D;&#39; as the value of the group name returned from the LDAP server might prepend those characters with &#39;\\&#39;. For example a group name called &#39;test#,lab,com&#39; will be returned from the LDAP server as &#39;test\\#,lab,com&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnAuthorizationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnAuthorizationGroup**](MsgVpnAuthorizationGroup.md) | The LDAP Authorization Group object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAuthorizationGroupResponse**](MsgVpnAuthorizationGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

