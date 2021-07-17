# \BridgeApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnBridge**](BridgeApi.md#CreateMsgVpnBridge) | **Post** /msgVpns/{msgVpnName}/bridges | Create a Bridge object.
[**CreateMsgVpnBridgeRemoteMsgVpn**](BridgeApi.md#CreateMsgVpnBridgeRemoteMsgVpn) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Create a Remote Message VPN object.
[**CreateMsgVpnBridgeRemoteSubscription**](BridgeApi.md#CreateMsgVpnBridgeRemoteSubscription) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Create a Remote Subscription object.
[**CreateMsgVpnBridgeTlsTrustedCommonName**](BridgeApi.md#CreateMsgVpnBridgeTlsTrustedCommonName) | **Post** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Create a Trusted Common Name object.
[**DeleteMsgVpnBridge**](BridgeApi.md#DeleteMsgVpnBridge) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Delete a Bridge object.
[**DeleteMsgVpnBridgeRemoteMsgVpn**](BridgeApi.md#DeleteMsgVpnBridgeRemoteMsgVpn) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Delete a Remote Message VPN object.
[**DeleteMsgVpnBridgeRemoteSubscription**](BridgeApi.md#DeleteMsgVpnBridgeRemoteSubscription) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Delete a Remote Subscription object.
[**DeleteMsgVpnBridgeTlsTrustedCommonName**](BridgeApi.md#DeleteMsgVpnBridgeTlsTrustedCommonName) | **Delete** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Delete a Trusted Common Name object.
[**GetMsgVpnBridge**](BridgeApi.md#GetMsgVpnBridge) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Get a Bridge object.
[**GetMsgVpnBridgeRemoteMsgVpn**](BridgeApi.md#GetMsgVpnBridgeRemoteMsgVpn) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Get a Remote Message VPN object.
[**GetMsgVpnBridgeRemoteMsgVpns**](BridgeApi.md#GetMsgVpnBridgeRemoteMsgVpns) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns | Get a list of Remote Message VPN objects.
[**GetMsgVpnBridgeRemoteSubscription**](BridgeApi.md#GetMsgVpnBridgeRemoteSubscription) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions/{remoteSubscriptionTopic} | Get a Remote Subscription object.
[**GetMsgVpnBridgeRemoteSubscriptions**](BridgeApi.md#GetMsgVpnBridgeRemoteSubscriptions) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteSubscriptions | Get a list of Remote Subscription objects.
[**GetMsgVpnBridgeTlsTrustedCommonName**](BridgeApi.md#GetMsgVpnBridgeTlsTrustedCommonName) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetMsgVpnBridgeTlsTrustedCommonNames**](BridgeApi.md#GetMsgVpnBridgeTlsTrustedCommonNames) | **Get** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetMsgVpnBridges**](BridgeApi.md#GetMsgVpnBridges) | **Get** /msgVpns/{msgVpnName}/bridges | Get a list of Bridge objects.
[**ReplaceMsgVpnBridge**](BridgeApi.md#ReplaceMsgVpnBridge) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Replace a Bridge object.
[**ReplaceMsgVpnBridgeRemoteMsgVpn**](BridgeApi.md#ReplaceMsgVpnBridgeRemoteMsgVpn) | **Put** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Replace a Remote Message VPN object.
[**UpdateMsgVpnBridge**](BridgeApi.md#UpdateMsgVpnBridge) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter} | Update a Bridge object.
[**UpdateMsgVpnBridgeRemoteMsgVpn**](BridgeApi.md#UpdateMsgVpnBridgeRemoteMsgVpn) | **Patch** /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface} | Update a Remote Message VPN object.



## CreateMsgVpnBridge

> MsgVpnBridgeResponse CreateMsgVpnBridge(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Bridge object.



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
    body := *openapiclient.NewMsgVpnBridge() // MsgVpnBridge | The Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.CreateMsgVpnBridge(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.CreateMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnBridge`: MsgVpnBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.CreateMsgVpnBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnBridge**](MsgVpnBridge.md) | The Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnBridgeRemoteMsgVpn

> MsgVpnBridgeRemoteMsgVpnResponse CreateMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Remote Message VPN object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    body := *openapiclient.NewMsgVpnBridgeRemoteMsgVpn() // MsgVpnBridgeRemoteMsgVpn | The Remote Message VPN object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.CreateMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.CreateMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnBridgeRemoteMsgVpn`: MsgVpnBridgeRemoteMsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.CreateMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnBridgeRemoteMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md) | The Remote Message VPN object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnBridgeRemoteSubscription

> MsgVpnBridgeRemoteSubscriptionResponse CreateMsgVpnBridgeRemoteSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Remote Subscription object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    body := *openapiclient.NewMsgVpnBridgeRemoteSubscription() // MsgVpnBridgeRemoteSubscription | The Remote Subscription object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.CreateMsgVpnBridgeRemoteSubscription(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.CreateMsgVpnBridgeRemoteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnBridgeRemoteSubscription`: MsgVpnBridgeRemoteSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.CreateMsgVpnBridgeRemoteSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnBridgeRemoteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnBridgeRemoteSubscription**](MsgVpnBridgeRemoteSubscription.md) | The Remote Subscription object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteSubscriptionResponse**](MsgVpnBridgeRemoteSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnBridgeTlsTrustedCommonName

> MsgVpnBridgeTlsTrustedCommonNameResponse CreateMsgVpnBridgeTlsTrustedCommonName(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Trusted Common Name object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    body := *openapiclient.NewMsgVpnBridgeTlsTrustedCommonName() // MsgVpnBridgeTlsTrustedCommonName | The Trusted Common Name object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.CreateMsgVpnBridgeTlsTrustedCommonName(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.CreateMsgVpnBridgeTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnBridgeTlsTrustedCommonName`: MsgVpnBridgeTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.CreateMsgVpnBridgeTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnBridgeTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnBridgeTlsTrustedCommonName**](MsgVpnBridgeTlsTrustedCommonName.md) | The Trusted Common Name object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNameResponse**](MsgVpnBridgeTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnBridge

> SempMetaOnlyResponse DeleteMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Execute()

Delete a Bridge object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.DeleteMsgVpnBridge(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.DeleteMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnBridge`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.DeleteMsgVpnBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnBridgeRequest struct via the builder pattern


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


## DeleteMsgVpnBridgeRemoteMsgVpn

> SempMetaOnlyResponse DeleteMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Execute()

Delete a Remote Message VPN object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteMsgVpnName := "remoteMsgVpnName_example" // string | The name of the remote Message VPN.
    remoteMsgVpnLocation := "remoteMsgVpnLocation_example" // string | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \"v:\").
    remoteMsgVpnInterface := "remoteMsgVpnInterface_example" // string | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remoteMsgVpnLocation` must not be a virtual router name.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.DeleteMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.DeleteMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnBridgeRemoteMsgVpn`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.DeleteMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteMsgVpnName** | **string** | The name of the remote Message VPN. | 
**remoteMsgVpnLocation** | **string** | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
**remoteMsgVpnInterface** | **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnBridgeRemoteMsgVpnRequest struct via the builder pattern


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


## DeleteMsgVpnBridgeRemoteSubscription

> SempMetaOnlyResponse DeleteMsgVpnBridgeRemoteSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic).Execute()

Delete a Remote Subscription object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteSubscriptionTopic := "remoteSubscriptionTopic_example" // string | The topic of the Bridge remote subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.DeleteMsgVpnBridgeRemoteSubscription(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.DeleteMsgVpnBridgeRemoteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnBridgeRemoteSubscription`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.DeleteMsgVpnBridgeRemoteSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteSubscriptionTopic** | **string** | The topic of the Bridge remote subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnBridgeRemoteSubscriptionRequest struct via the builder pattern


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


## DeleteMsgVpnBridgeTlsTrustedCommonName

> SempMetaOnlyResponse DeleteMsgVpnBridgeTlsTrustedCommonName(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName).Execute()

Delete a Trusted Common Name object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    tlsTrustedCommonName := "tlsTrustedCommonName_example" // string | The expected trusted common name of the remote certificate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.DeleteMsgVpnBridgeTlsTrustedCommonName(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.DeleteMsgVpnBridgeTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnBridgeTlsTrustedCommonName`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.DeleteMsgVpnBridgeTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**tlsTrustedCommonName** | **string** | The expected trusted common name of the remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnBridgeTlsTrustedCommonNameRequest struct via the builder pattern


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


## GetMsgVpnBridge

> MsgVpnBridgeResponse GetMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Bridge object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.GetMsgVpnBridge(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.GetMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridge`: MsgVpnBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.GetMsgVpnBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeRemoteMsgVpn

> MsgVpnBridgeRemoteMsgVpnResponse GetMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Remote Message VPN object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteMsgVpnName := "remoteMsgVpnName_example" // string | The name of the remote Message VPN.
    remoteMsgVpnLocation := "remoteMsgVpnLocation_example" // string | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \"v:\").
    remoteMsgVpnInterface := "remoteMsgVpnInterface_example" // string | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remoteMsgVpnLocation` must not be a virtual router name.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.GetMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.GetMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteMsgVpn`: MsgVpnBridgeRemoteMsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.GetMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteMsgVpnName** | **string** | The name of the remote Message VPN. | 
**remoteMsgVpnLocation** | **string** | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
**remoteMsgVpnInterface** | **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeRemoteMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeRemoteMsgVpns

> MsgVpnBridgeRemoteMsgVpnsResponse GetMsgVpnBridgeRemoteMsgVpns(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Remote Message VPN objects.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.GetMsgVpnBridgeRemoteMsgVpns(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.GetMsgVpnBridgeRemoteMsgVpns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteMsgVpns`: MsgVpnBridgeRemoteMsgVpnsResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.GetMsgVpnBridgeRemoteMsgVpns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeRemoteMsgVpnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnsResponse**](MsgVpnBridgeRemoteMsgVpnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeRemoteSubscription

> MsgVpnBridgeRemoteSubscriptionResponse GetMsgVpnBridgeRemoteSubscription(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Remote Subscription object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteSubscriptionTopic := "remoteSubscriptionTopic_example" // string | The topic of the Bridge remote subscription.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.GetMsgVpnBridgeRemoteSubscription(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteSubscriptionTopic).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.GetMsgVpnBridgeRemoteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteSubscription`: MsgVpnBridgeRemoteSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.GetMsgVpnBridgeRemoteSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteSubscriptionTopic** | **string** | The topic of the Bridge remote subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeRemoteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteSubscriptionResponse**](MsgVpnBridgeRemoteSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeRemoteSubscriptions

> MsgVpnBridgeRemoteSubscriptionsResponse GetMsgVpnBridgeRemoteSubscriptions(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Remote Subscription objects.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.GetMsgVpnBridgeRemoteSubscriptions(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.GetMsgVpnBridgeRemoteSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeRemoteSubscriptions`: MsgVpnBridgeRemoteSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.GetMsgVpnBridgeRemoteSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeRemoteSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteSubscriptionsResponse**](MsgVpnBridgeRemoteSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeTlsTrustedCommonName

> MsgVpnBridgeTlsTrustedCommonNameResponse GetMsgVpnBridgeTlsTrustedCommonName(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Trusted Common Name object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    tlsTrustedCommonName := "tlsTrustedCommonName_example" // string | The expected trusted common name of the remote certificate.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.GetMsgVpnBridgeTlsTrustedCommonName(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, tlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.GetMsgVpnBridgeTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeTlsTrustedCommonName`: MsgVpnBridgeTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.GetMsgVpnBridgeTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**tlsTrustedCommonName** | **string** | The expected trusted common name of the remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNameResponse**](MsgVpnBridgeTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridgeTlsTrustedCommonNames

> MsgVpnBridgeTlsTrustedCommonNamesResponse GetMsgVpnBridgeTlsTrustedCommonNames(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Trusted Common Name objects.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.GetMsgVpnBridgeTlsTrustedCommonNames(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.GetMsgVpnBridgeTlsTrustedCommonNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridgeTlsTrustedCommonNames`: MsgVpnBridgeTlsTrustedCommonNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.GetMsgVpnBridgeTlsTrustedCommonNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgeTlsTrustedCommonNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeTlsTrustedCommonNamesResponse**](MsgVpnBridgeTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnBridges

> MsgVpnBridgesResponse GetMsgVpnBridges(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Bridge objects.



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
    resp, r, err := api_client.BridgeApi.GetMsgVpnBridges(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.GetMsgVpnBridges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnBridges`: MsgVpnBridgesResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.GetMsgVpnBridges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnBridgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgesResponse**](MsgVpnBridgesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnBridge

> MsgVpnBridgeResponse ReplaceMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Bridge object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    body := *openapiclient.NewMsgVpnBridge() // MsgVpnBridge | The Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.ReplaceMsgVpnBridge(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.ReplaceMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnBridge`: MsgVpnBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.ReplaceMsgVpnBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnBridge**](MsgVpnBridge.md) | The Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnBridgeRemoteMsgVpn

> MsgVpnBridgeRemoteMsgVpnResponse ReplaceMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Remote Message VPN object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteMsgVpnName := "remoteMsgVpnName_example" // string | The name of the remote Message VPN.
    remoteMsgVpnLocation := "remoteMsgVpnLocation_example" // string | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \"v:\").
    remoteMsgVpnInterface := "remoteMsgVpnInterface_example" // string | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remoteMsgVpnLocation` must not be a virtual router name.
    body := *openapiclient.NewMsgVpnBridgeRemoteMsgVpn() // MsgVpnBridgeRemoteMsgVpn | The Remote Message VPN object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.ReplaceMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.ReplaceMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnBridgeRemoteMsgVpn`: MsgVpnBridgeRemoteMsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.ReplaceMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteMsgVpnName** | **string** | The name of the remote Message VPN. | 
**remoteMsgVpnLocation** | **string** | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
**remoteMsgVpnInterface** | **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnBridgeRemoteMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md) | The Remote Message VPN object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnBridge

> MsgVpnBridgeResponse UpdateMsgVpnBridge(ctx, msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Bridge object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    body := *openapiclient.NewMsgVpnBridge() // MsgVpnBridge | The Bridge object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.UpdateMsgVpnBridge(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.UpdateMsgVpnBridge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnBridge`: MsgVpnBridgeResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.UpdateMsgVpnBridge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnBridgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnBridge**](MsgVpnBridge.md) | The Bridge object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeResponse**](MsgVpnBridgeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnBridgeRemoteMsgVpn

> MsgVpnBridgeRemoteMsgVpnResponse UpdateMsgVpnBridgeRemoteMsgVpn(ctx, msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Remote Message VPN object.



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
    bridgeName := "bridgeName_example" // string | The name of the Bridge.
    bridgeVirtualRouter := "bridgeVirtualRouter_example" // string | The virtual router of the Bridge.
    remoteMsgVpnName := "remoteMsgVpnName_example" // string | The name of the remote Message VPN.
    remoteMsgVpnLocation := "remoteMsgVpnLocation_example" // string | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \"v:\").
    remoteMsgVpnInterface := "remoteMsgVpnInterface_example" // string | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remoteMsgVpnLocation` must not be a virtual router name.
    body := *openapiclient.NewMsgVpnBridgeRemoteMsgVpn() // MsgVpnBridgeRemoteMsgVpn | The Remote Message VPN object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BridgeApi.UpdateMsgVpnBridgeRemoteMsgVpn(context.Background(), msgVpnName, bridgeName, bridgeVirtualRouter, remoteMsgVpnName, remoteMsgVpnLocation, remoteMsgVpnInterface).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BridgeApi.UpdateMsgVpnBridgeRemoteMsgVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnBridgeRemoteMsgVpn`: MsgVpnBridgeRemoteMsgVpnResponse
    fmt.Fprintf(os.Stdout, "Response from `BridgeApi.UpdateMsgVpnBridgeRemoteMsgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**bridgeName** | **string** | The name of the Bridge. | 
**bridgeVirtualRouter** | **string** | The virtual router of the Bridge. | 
**remoteMsgVpnName** | **string** | The name of the remote Message VPN. | 
**remoteMsgVpnLocation** | **string** | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | 
**remoteMsgVpnInterface** | **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnBridgeRemoteMsgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | [**MsgVpnBridgeRemoteMsgVpn**](MsgVpnBridgeRemoteMsgVpn.md) | The Remote Message VPN object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnBridgeRemoteMsgVpnResponse**](MsgVpnBridgeRemoteMsgVpnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

