# \DmrBridgeApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnDmrBridge**](DmrBridgeApi.md#CreateMsgVpnDmrBridge) | **Post** /msgVpns/{msgVpnName}/dmrBridges | Create a DMR Bridge object.
[**DeleteMsgVpnDmrBridge**](DmrBridgeApi.md#DeleteMsgVpnDmrBridge) | **Delete** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Delete a DMR Bridge object.
[**GetMsgVpnDmrBridge**](DmrBridgeApi.md#GetMsgVpnDmrBridge) | **Get** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Get a DMR Bridge object.
[**GetMsgVpnDmrBridges**](DmrBridgeApi.md#GetMsgVpnDmrBridges) | **Get** /msgVpns/{msgVpnName}/dmrBridges | Get a list of DMR Bridge objects.
[**ReplaceMsgVpnDmrBridge**](DmrBridgeApi.md#ReplaceMsgVpnDmrBridge) | **Put** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Replace a DMR Bridge object.
[**UpdateMsgVpnDmrBridge**](DmrBridgeApi.md#UpdateMsgVpnDmrBridge) | **Patch** /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName} | Update a DMR Bridge object.



## CreateMsgVpnDmrBridge

> MsgVpnDmrBridgeResponse CreateMsgVpnDmrBridge(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a DMR Bridge object.



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
    body := *openapiclient.NewMsgVpnDmrBridge() // MsgVpnDmrBridge | The DMR Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrBridgeApi.CreateMsgVpnDmrBridge(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrBridgeApi.CreateMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDmrBridge`: MsgVpnDmrBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrBridgeApi.CreateMsgVpnDmrBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDmrBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnDmrBridge**](MsgVpnDmrBridge.md) | The DMR Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnDmrBridge

> SempMetaOnlyResponse DeleteMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName).Execute()

Delete a DMR Bridge object.



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
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the DMR Bridge.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrBridgeApi.DeleteMsgVpnDmrBridge(context.Background(), msgVpnName, remoteNodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrBridgeApi.DeleteMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDmrBridge`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrBridgeApi.DeleteMsgVpnDmrBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the DMR Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDmrBridgeRequest struct via the builder pattern


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


## GetMsgVpnDmrBridge

> MsgVpnDmrBridgeResponse GetMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a DMR Bridge object.



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
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the DMR Bridge.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrBridgeApi.GetMsgVpnDmrBridge(context.Background(), msgVpnName, remoteNodeName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrBridgeApi.GetMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDmrBridge`: MsgVpnDmrBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrBridgeApi.GetMsgVpnDmrBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the DMR Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDmrBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDmrBridges

> MsgVpnDmrBridgesResponse GetMsgVpnDmrBridges(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of DMR Bridge objects.



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
    resp, r, err := api_client.DmrBridgeApi.GetMsgVpnDmrBridges(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrBridgeApi.GetMsgVpnDmrBridges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDmrBridges`: MsgVpnDmrBridgesResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrBridgeApi.GetMsgVpnDmrBridges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDmrBridgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgesResponse**](MsgVpnDmrBridgesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnDmrBridge

> MsgVpnDmrBridgeResponse ReplaceMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a DMR Bridge object.



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
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the DMR Bridge.
    body := *openapiclient.NewMsgVpnDmrBridge() // MsgVpnDmrBridge | The DMR Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrBridgeApi.ReplaceMsgVpnDmrBridge(context.Background(), msgVpnName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrBridgeApi.ReplaceMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnDmrBridge`: MsgVpnDmrBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrBridgeApi.ReplaceMsgVpnDmrBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the DMR Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnDmrBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnDmrBridge**](MsgVpnDmrBridge.md) | The DMR Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnDmrBridge

> MsgVpnDmrBridgeResponse UpdateMsgVpnDmrBridge(ctx, msgVpnName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a DMR Bridge object.



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
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the DMR Bridge.
    body := *openapiclient.NewMsgVpnDmrBridge() // MsgVpnDmrBridge | The DMR Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrBridgeApi.UpdateMsgVpnDmrBridge(context.Background(), msgVpnName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrBridgeApi.UpdateMsgVpnDmrBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnDmrBridge`: MsgVpnDmrBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrBridgeApi.UpdateMsgVpnDmrBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the DMR Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnDmrBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnDmrBridge**](MsgVpnDmrBridge.md) | The DMR Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDmrBridgeResponse**](MsgVpnDmrBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

