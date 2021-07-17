# \DistributedCacheApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnDistributedCache**](DistributedCacheApi.md#CreateMsgVpnDistributedCache) | **Post** /msgVpns/{msgVpnName}/distributedCaches | Create a Distributed Cache object.
[**CreateMsgVpnDistributedCacheCluster**](DistributedCacheApi.md#CreateMsgVpnDistributedCacheCluster) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Create a Cache Cluster object.
[**CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](DistributedCacheApi.md#CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Create a Home Cache Cluster object.
[**CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](DistributedCacheApi.md#CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Create a Topic Prefix object.
[**CreateMsgVpnDistributedCacheClusterInstance**](DistributedCacheApi.md#CreateMsgVpnDistributedCacheClusterInstance) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Create a Cache Instance object.
[**CreateMsgVpnDistributedCacheClusterTopic**](DistributedCacheApi.md#CreateMsgVpnDistributedCacheClusterTopic) | **Post** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics | Create a Topic object.
[**DeleteMsgVpnDistributedCache**](DistributedCacheApi.md#DeleteMsgVpnDistributedCache) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Delete a Distributed Cache object.
[**DeleteMsgVpnDistributedCacheCluster**](DistributedCacheApi.md#DeleteMsgVpnDistributedCacheCluster) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Delete a Cache Cluster object.
[**DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](DistributedCacheApi.md#DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Delete a Home Cache Cluster object.
[**DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](DistributedCacheApi.md#DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Delete a Topic Prefix object.
[**DeleteMsgVpnDistributedCacheClusterInstance**](DistributedCacheApi.md#DeleteMsgVpnDistributedCacheClusterInstance) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Delete a Cache Instance object.
[**DeleteMsgVpnDistributedCacheClusterTopic**](DistributedCacheApi.md#DeleteMsgVpnDistributedCacheClusterTopic) | **Delete** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic} | Delete a Topic object.
[**GetMsgVpnDistributedCache**](DistributedCacheApi.md#GetMsgVpnDistributedCache) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Get a Distributed Cache object.
[**GetMsgVpnDistributedCacheCluster**](DistributedCacheApi.md#GetMsgVpnDistributedCacheCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Get a Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName} | Get a Home Cache Cluster object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix} | Get a Topic Prefix object.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes | Get a list of Topic Prefix objects.
[**GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters | Get a list of Home Cache Cluster objects.
[**GetMsgVpnDistributedCacheClusterInstance**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusterInstance) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Get a Cache Instance object.
[**GetMsgVpnDistributedCacheClusterInstances**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusterInstances) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances | Get a list of Cache Instance objects.
[**GetMsgVpnDistributedCacheClusterTopic**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusterTopic) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic} | Get a Topic object.
[**GetMsgVpnDistributedCacheClusterTopics**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusterTopics) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics | Get a list of Topic objects.
[**GetMsgVpnDistributedCacheClusters**](DistributedCacheApi.md#GetMsgVpnDistributedCacheClusters) | **Get** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters | Get a list of Cache Cluster objects.
[**GetMsgVpnDistributedCaches**](DistributedCacheApi.md#GetMsgVpnDistributedCaches) | **Get** /msgVpns/{msgVpnName}/distributedCaches | Get a list of Distributed Cache objects.
[**ReplaceMsgVpnDistributedCache**](DistributedCacheApi.md#ReplaceMsgVpnDistributedCache) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Replace a Distributed Cache object.
[**ReplaceMsgVpnDistributedCacheCluster**](DistributedCacheApi.md#ReplaceMsgVpnDistributedCacheCluster) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Replace a Cache Cluster object.
[**ReplaceMsgVpnDistributedCacheClusterInstance**](DistributedCacheApi.md#ReplaceMsgVpnDistributedCacheClusterInstance) | **Put** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Replace a Cache Instance object.
[**UpdateMsgVpnDistributedCache**](DistributedCacheApi.md#UpdateMsgVpnDistributedCache) | **Patch** /msgVpns/{msgVpnName}/distributedCaches/{cacheName} | Update a Distributed Cache object.
[**UpdateMsgVpnDistributedCacheCluster**](DistributedCacheApi.md#UpdateMsgVpnDistributedCacheCluster) | **Patch** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName} | Update a Cache Cluster object.
[**UpdateMsgVpnDistributedCacheClusterInstance**](DistributedCacheApi.md#UpdateMsgVpnDistributedCacheClusterInstance) | **Patch** /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName} | Update a Cache Instance object.



## CreateMsgVpnDistributedCache

> MsgVpnDistributedCacheResponse CreateMsgVpnDistributedCache(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Distributed Cache object.



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
    body := *openapiclient.NewMsgVpnDistributedCache() // MsgVpnDistributedCache | The Distributed Cache object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.CreateMsgVpnDistributedCache(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.CreateMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCache`: MsgVpnDistributedCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.CreateMsgVpnDistributedCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnDistributedCache**](MsgVpnDistributedCache.md) | The Distributed Cache object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCacheCluster

> MsgVpnDistributedCacheClusterResponse CreateMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Cache Cluster object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    body := *openapiclient.NewMsgVpnDistributedCacheCluster() // MsgVpnDistributedCacheCluster | The Cache Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.CreateMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.CreateMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCacheCluster`: MsgVpnDistributedCacheClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.CreateMsgVpnDistributedCacheCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnDistributedCacheCluster**](MsgVpnDistributedCacheCluster.md) | The Cache Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Home Cache Cluster object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterGlobalCachingHomeCluster() // MsgVpnDistributedCacheClusterGlobalCachingHomeCluster | The Home Cache Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(context.Background(), msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnDistributedCacheClusterGlobalCachingHomeCluster**](MsgVpnDistributedCacheClusterGlobalCachingHomeCluster.md) | The Home Cache Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx, msgVpnName, cacheName, clusterName, homeClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Topic Prefix object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix() // MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix | The Topic Prefix object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix.md) | The Topic Prefix object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCacheClusterInstance

> MsgVpnDistributedCacheClusterInstanceResponse CreateMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Cache Instance object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterInstance() // MsgVpnDistributedCacheClusterInstance | The Cache Instance object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.CreateMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.CreateMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCacheClusterInstance`: MsgVpnDistributedCacheClusterInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.CreateMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheClusterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnDistributedCacheClusterInstance**](MsgVpnDistributedCacheClusterInstance.md) | The Cache Instance object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMsgVpnDistributedCacheClusterTopic

> MsgVpnDistributedCacheClusterTopicResponse CreateMsgVpnDistributedCacheClusterTopic(ctx, msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Topic object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterTopic() // MsgVpnDistributedCacheClusterTopic | The Topic object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.CreateMsgVpnDistributedCacheClusterTopic(context.Background(), msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.CreateMsgVpnDistributedCacheClusterTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnDistributedCacheClusterTopic`: MsgVpnDistributedCacheClusterTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.CreateMsgVpnDistributedCacheClusterTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnDistributedCacheClusterTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnDistributedCacheClusterTopic**](MsgVpnDistributedCacheClusterTopic.md) | The Topic object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterTopicResponse**](MsgVpnDistributedCacheClusterTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnDistributedCache

> SempMetaOnlyResponse DeleteMsgVpnDistributedCache(ctx, msgVpnName, cacheName).Execute()

Delete a Distributed Cache object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.DeleteMsgVpnDistributedCache(context.Background(), msgVpnName, cacheName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.DeleteMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCache`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.DeleteMsgVpnDistributedCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheRequest struct via the builder pattern


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


## DeleteMsgVpnDistributedCacheCluster

> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName).Execute()

Delete a Cache Cluster object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.DeleteMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.DeleteMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCacheCluster`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.DeleteMsgVpnDistributedCacheCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheClusterRequest struct via the builder pattern


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


## DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster

> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName, homeClusterName).Execute()

Delete a Home Cache Cluster object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest struct via the builder pattern


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


## DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix

> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx, msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix).Execute()

Delete a Topic Prefix object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    topicPrefix := "topicPrefix_example" // string | A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/>) is implied at the end of the prefix.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 
**topicPrefix** | **string** | A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/&gt;) is implied at the end of the prefix. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest struct via the builder pattern


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


## DeleteMsgVpnDistributedCacheClusterInstance

> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName).Execute()

Delete a Cache Instance object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    instanceName := "instanceName_example" // string | The name of the Cache Instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName, instanceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCacheClusterInstance`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**instanceName** | **string** | The name of the Cache Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheClusterInstanceRequest struct via the builder pattern


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


## DeleteMsgVpnDistributedCacheClusterTopic

> SempMetaOnlyResponse DeleteMsgVpnDistributedCacheClusterTopic(ctx, msgVpnName, cacheName, clusterName, topic).Execute()

Delete a Topic object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    topic := "topic_example" // string | The value of the Topic in the form a/b/c.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterTopic(context.Background(), msgVpnName, cacheName, clusterName, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnDistributedCacheClusterTopic`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.DeleteMsgVpnDistributedCacheClusterTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**topic** | **string** | The value of the Topic in the form a/b/c. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnDistributedCacheClusterTopicRequest struct via the builder pattern


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


## GetMsgVpnDistributedCache

> MsgVpnDistributedCacheResponse GetMsgVpnDistributedCache(ctx, msgVpnName, cacheName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Distributed Cache object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCache(context.Background(), msgVpnName, cacheName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCache`: MsgVpnDistributedCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheCluster

> MsgVpnDistributedCacheClusterResponse GetMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Cache Cluster object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName, clusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheCluster`: MsgVpnDistributedCacheClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCacheCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(ctx, msgVpnName, cacheName, clusterName, homeClusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Home Cache Cluster object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(ctx, msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Topic Prefix object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    topicPrefix := "topicPrefix_example" // string | A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/>) is implied at the end of the prefix.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName, topicPrefix).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 
**topicPrefix** | **string** | A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/&gt;) is implied at the end of the prefix. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes

> MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes(ctx, msgVpnName, cacheName, clusterName, homeClusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Topic Prefix objects.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    homeClusterName := "homeClusterName_example" // string | The name of the remote Home Cache Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes(context.Background(), msgVpnName, cacheName, clusterName, homeClusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes`: MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**homeClusterName** | **string** | The name of the remote Home Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters

> MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters(ctx, msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Home Cache Cluster objects.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters(context.Background(), msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters`: MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterGlobalCachingHomeClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse**](MsgVpnDistributedCacheClusterGlobalCachingHomeClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterInstance

> MsgVpnDistributedCacheClusterInstanceResponse GetMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Cache Instance object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    instanceName := "instanceName_example" // string | The name of the Cache Instance.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName, instanceName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterInstance`: MsgVpnDistributedCacheClusterInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**instanceName** | **string** | The name of the Cache Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterInstances

> MsgVpnDistributedCacheClusterInstancesResponse GetMsgVpnDistributedCacheClusterInstances(ctx, msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Cache Instance objects.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCacheClusterInstances(context.Background(), msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCacheClusterInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterInstances`: MsgVpnDistributedCacheClusterInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCacheClusterInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstancesResponse**](MsgVpnDistributedCacheClusterInstancesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterTopic

> MsgVpnDistributedCacheClusterTopicResponse GetMsgVpnDistributedCacheClusterTopic(ctx, msgVpnName, cacheName, clusterName, topic).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Topic object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    topic := "topic_example" // string | The value of the Topic in the form a/b/c.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCacheClusterTopic(context.Background(), msgVpnName, cacheName, clusterName, topic).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCacheClusterTopic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterTopic`: MsgVpnDistributedCacheClusterTopicResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCacheClusterTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**topic** | **string** | The value of the Topic in the form a/b/c. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterTopicResponse**](MsgVpnDistributedCacheClusterTopicResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusterTopics

> MsgVpnDistributedCacheClusterTopicsResponse GetMsgVpnDistributedCacheClusterTopics(ctx, msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Topic objects.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCacheClusterTopics(context.Background(), msgVpnName, cacheName, clusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCacheClusterTopics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusterTopics`: MsgVpnDistributedCacheClusterTopicsResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCacheClusterTopics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClusterTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterTopicsResponse**](MsgVpnDistributedCacheClusterTopicsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCacheClusters

> MsgVpnDistributedCacheClustersResponse GetMsgVpnDistributedCacheClusters(ctx, msgVpnName, cacheName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Cache Cluster objects.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCacheClusters(context.Background(), msgVpnName, cacheName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCacheClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCacheClusters`: MsgVpnDistributedCacheClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCacheClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCacheClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClustersResponse**](MsgVpnDistributedCacheClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnDistributedCaches

> MsgVpnDistributedCachesResponse GetMsgVpnDistributedCaches(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Distributed Cache objects.



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
    resp, r, err := api_client.DistributedCacheApi.GetMsgVpnDistributedCaches(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.GetMsgVpnDistributedCaches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnDistributedCaches`: MsgVpnDistributedCachesResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.GetMsgVpnDistributedCaches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnDistributedCachesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCachesResponse**](MsgVpnDistributedCachesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnDistributedCache

> MsgVpnDistributedCacheResponse ReplaceMsgVpnDistributedCache(ctx, msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Distributed Cache object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    body := *openapiclient.NewMsgVpnDistributedCache() // MsgVpnDistributedCache | The Distributed Cache object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.ReplaceMsgVpnDistributedCache(context.Background(), msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.ReplaceMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnDistributedCache`: MsgVpnDistributedCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.ReplaceMsgVpnDistributedCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnDistributedCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnDistributedCache**](MsgVpnDistributedCache.md) | The Distributed Cache object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnDistributedCacheCluster

> MsgVpnDistributedCacheClusterResponse ReplaceMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Cache Cluster object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheCluster() // MsgVpnDistributedCacheCluster | The Cache Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.ReplaceMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.ReplaceMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnDistributedCacheCluster`: MsgVpnDistributedCacheClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.ReplaceMsgVpnDistributedCacheCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnDistributedCacheClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnDistributedCacheCluster**](MsgVpnDistributedCacheCluster.md) | The Cache Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnDistributedCacheClusterInstance

> MsgVpnDistributedCacheClusterInstanceResponse ReplaceMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Cache Instance object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    instanceName := "instanceName_example" // string | The name of the Cache Instance.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterInstance() // MsgVpnDistributedCacheClusterInstance | The Cache Instance object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.ReplaceMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName, instanceName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.ReplaceMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnDistributedCacheClusterInstance`: MsgVpnDistributedCacheClusterInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.ReplaceMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**instanceName** | **string** | The name of the Cache Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnDistributedCacheClusterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**MsgVpnDistributedCacheClusterInstance**](MsgVpnDistributedCacheClusterInstance.md) | The Cache Instance object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnDistributedCache

> MsgVpnDistributedCacheResponse UpdateMsgVpnDistributedCache(ctx, msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Distributed Cache object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    body := *openapiclient.NewMsgVpnDistributedCache() // MsgVpnDistributedCache | The Distributed Cache object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.UpdateMsgVpnDistributedCache(context.Background(), msgVpnName, cacheName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.UpdateMsgVpnDistributedCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnDistributedCache`: MsgVpnDistributedCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.UpdateMsgVpnDistributedCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnDistributedCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnDistributedCache**](MsgVpnDistributedCache.md) | The Distributed Cache object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheResponse**](MsgVpnDistributedCacheResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnDistributedCacheCluster

> MsgVpnDistributedCacheClusterResponse UpdateMsgVpnDistributedCacheCluster(ctx, msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Cache Cluster object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    body := *openapiclient.NewMsgVpnDistributedCacheCluster() // MsgVpnDistributedCacheCluster | The Cache Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.UpdateMsgVpnDistributedCacheCluster(context.Background(), msgVpnName, cacheName, clusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.UpdateMsgVpnDistributedCacheCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnDistributedCacheCluster`: MsgVpnDistributedCacheClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.UpdateMsgVpnDistributedCacheCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnDistributedCacheClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**MsgVpnDistributedCacheCluster**](MsgVpnDistributedCacheCluster.md) | The Cache Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterResponse**](MsgVpnDistributedCacheClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnDistributedCacheClusterInstance

> MsgVpnDistributedCacheClusterInstanceResponse UpdateMsgVpnDistributedCacheClusterInstance(ctx, msgVpnName, cacheName, clusterName, instanceName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Cache Instance object.



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
    cacheName := "cacheName_example" // string | The name of the Distributed Cache.
    clusterName := "clusterName_example" // string | The name of the Cache Cluster.
    instanceName := "instanceName_example" // string | The name of the Cache Instance.
    body := *openapiclient.NewMsgVpnDistributedCacheClusterInstance() // MsgVpnDistributedCacheClusterInstance | The Cache Instance object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedCacheApi.UpdateMsgVpnDistributedCacheClusterInstance(context.Background(), msgVpnName, cacheName, clusterName, instanceName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedCacheApi.UpdateMsgVpnDistributedCacheClusterInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnDistributedCacheClusterInstance`: MsgVpnDistributedCacheClusterInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DistributedCacheApi.UpdateMsgVpnDistributedCacheClusterInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**cacheName** | **string** | The name of the Distributed Cache. | 
**clusterName** | **string** | The name of the Cache Cluster. | 
**instanceName** | **string** | The name of the Cache Instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnDistributedCacheClusterInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**MsgVpnDistributedCacheClusterInstance**](MsgVpnDistributedCacheClusterInstance.md) | The Cache Instance object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnDistributedCacheClusterInstanceResponse**](MsgVpnDistributedCacheClusterInstanceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

