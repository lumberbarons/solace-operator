# \DmrClusterApi

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDmrCluster**](DmrClusterApi.md#GetDmrCluster) | **Get** /dmrClusters/{dmrClusterName} | Get a Cluster object.
[**GetDmrClusterLink**](DmrClusterApi.md#GetDmrClusterLink) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Get a Link object.
[**GetDmrClusterLinkChannel**](DmrClusterApi.md#GetDmrClusterLinkChannel) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/channels/{msgVpnName} | Get a Cluster Link Channels object.
[**GetDmrClusterLinkChannels**](DmrClusterApi.md#GetDmrClusterLinkChannels) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/channels | Get a list of Cluster Link Channels objects.
[**GetDmrClusterLinkRemoteAddress**](DmrClusterApi.md#GetDmrClusterLinkRemoteAddress) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses/{remoteAddress} | Get a Remote Address object.
[**GetDmrClusterLinkRemoteAddresses**](DmrClusterApi.md#GetDmrClusterLinkRemoteAddresses) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses | Get a list of Remote Address objects.
[**GetDmrClusterLinkTlsTrustedCommonName**](DmrClusterApi.md#GetDmrClusterLinkTlsTrustedCommonName) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetDmrClusterLinkTlsTrustedCommonNames**](DmrClusterApi.md#GetDmrClusterLinkTlsTrustedCommonNames) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetDmrClusterLinks**](DmrClusterApi.md#GetDmrClusterLinks) | **Get** /dmrClusters/{dmrClusterName}/links | Get a list of Link objects.
[**GetDmrClusterTopologyIssue**](DmrClusterApi.md#GetDmrClusterTopologyIssue) | **Get** /dmrClusters/{dmrClusterName}/topologyIssues/{topologyIssue} | Get a Cluster Topology Issue object.
[**GetDmrClusterTopologyIssues**](DmrClusterApi.md#GetDmrClusterTopologyIssues) | **Get** /dmrClusters/{dmrClusterName}/topologyIssues | Get a list of Cluster Topology Issue objects.
[**GetDmrClusters**](DmrClusterApi.md#GetDmrClusters) | **Get** /dmrClusters | Get a list of Cluster objects.



## GetDmrCluster

> DmrClusterResponse GetDmrCluster(ctx, dmrClusterName).Select_(select_).Execute()

Get a Cluster object.



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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrCluster(context.Background(), dmrClusterName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrCluster`: DmrClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLink

> DmrClusterLinkResponse GetDmrClusterLink(ctx, dmrClusterName, remoteNodeName).Select_(select_).Execute()

Get a Link object.



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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLink(context.Background(), dmrClusterName, remoteNodeName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusterLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLink`: DmrClusterLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusterLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinkChannel

> DmrClusterLinkChannelResponse GetDmrClusterLinkChannel(ctx, dmrClusterName, remoteNodeName, msgVpnName).Select_(select_).Execute()

Get a Cluster Link Channels object.



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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    msgVpnName := "msgVpnName_example" // string | The name of the Message VPN.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinkChannel(context.Background(), dmrClusterName, remoteNodeName, msgVpnName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusterLinkChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinkChannel`: DmrClusterLinkChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusterLinkChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkChannelResponse**](DmrClusterLinkChannelResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinkChannels

> DmrClusterLinkChannelsResponse GetDmrClusterLinkChannels(ctx, dmrClusterName, remoteNodeName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Cluster Link Channels objects.



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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinkChannels(context.Background(), dmrClusterName, remoteNodeName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusterLinkChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinkChannels`: DmrClusterLinkChannelsResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusterLinkChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkChannelsResponse**](DmrClusterLinkChannelsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinkRemoteAddress

> DmrClusterLinkRemoteAddressResponse GetDmrClusterLinkRemoteAddress(ctx, dmrClusterName, remoteNodeName, remoteAddress).Select_(select_).Execute()

Get a Remote Address object.



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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    remoteAddress := "remoteAddress_example" // string | The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed).
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinkRemoteAddress(context.Background(), dmrClusterName, remoteNodeName, remoteAddress).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusterLinkRemoteAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinkRemoteAddress`: DmrClusterLinkRemoteAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusterLinkRemoteAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 
**remoteAddress** | **string** | The FQDN or IP address (and optional port) of the remote node. If a port is not provided, it will vary based on the transport encoding: 55555 (plain-text), 55443 (encrypted), or 55003 (compressed). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkRemoteAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkRemoteAddressResponse**](DmrClusterLinkRemoteAddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinkRemoteAddresses

> DmrClusterLinkRemoteAddressesResponse GetDmrClusterLinkRemoteAddresses(ctx, dmrClusterName, remoteNodeName).Where(where).Select_(select_).Execute()

Get a list of Remote Address objects.



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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinkRemoteAddresses(context.Background(), dmrClusterName, remoteNodeName).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusterLinkRemoteAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinkRemoteAddresses`: DmrClusterLinkRemoteAddressesResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusterLinkRemoteAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkRemoteAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkRemoteAddressesResponse**](DmrClusterLinkRemoteAddressesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinkTlsTrustedCommonName

> DmrClusterLinkTlsTrustedCommonNameResponse GetDmrClusterLinkTlsTrustedCommonName(ctx, dmrClusterName, remoteNodeName, tlsTrustedCommonName).Select_(select_).Execute()

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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    tlsTrustedCommonName := "tlsTrustedCommonName_example" // string | The expected trusted common name of the remote certificate.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinkTlsTrustedCommonName(context.Background(), dmrClusterName, remoteNodeName, tlsTrustedCommonName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusterLinkTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinkTlsTrustedCommonName`: DmrClusterLinkTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusterLinkTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 
**tlsTrustedCommonName** | **string** | The expected trusted common name of the remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkTlsTrustedCommonNameResponse**](DmrClusterLinkTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinkTlsTrustedCommonNames

> DmrClusterLinkTlsTrustedCommonNamesResponse GetDmrClusterLinkTlsTrustedCommonNames(ctx, dmrClusterName, remoteNodeName).Where(where).Select_(select_).Execute()

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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinkTlsTrustedCommonNames(context.Background(), dmrClusterName, remoteNodeName).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusterLinkTlsTrustedCommonNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinkTlsTrustedCommonNames`: DmrClusterLinkTlsTrustedCommonNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusterLinkTlsTrustedCommonNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinkTlsTrustedCommonNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkTlsTrustedCommonNamesResponse**](DmrClusterLinkTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterLinks

> DmrClusterLinksResponse GetDmrClusterLinks(ctx, dmrClusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Link objects.



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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinks(context.Background(), dmrClusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusterLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterLinks`: DmrClusterLinksResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusterLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinksResponse**](DmrClusterLinksResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterTopologyIssue

> DmrClusterTopologyIssueResponse GetDmrClusterTopologyIssue(ctx, dmrClusterName, topologyIssue).Select_(select_).Execute()

Get a Cluster Topology Issue object.



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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    topologyIssue := "topologyIssue_example" // string | The topology issue discovered in the Cluster. A topology issue indicates incorrect or inconsistent configuration within the DMR network. Such issues will cause messages to be misdelivered or lost.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterTopologyIssue(context.Background(), dmrClusterName, topologyIssue).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusterTopologyIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterTopologyIssue`: DmrClusterTopologyIssueResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusterTopologyIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**topologyIssue** | **string** | The topology issue discovered in the Cluster. A topology issue indicates incorrect or inconsistent configuration within the DMR network. Such issues will cause messages to be misdelivered or lost. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterTopologyIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterTopologyIssueResponse**](DmrClusterTopologyIssueResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusterTopologyIssues

> DmrClusterTopologyIssuesResponse GetDmrClusterTopologyIssues(ctx, dmrClusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Cluster Topology Issue objects.



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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterTopologyIssues(context.Background(), dmrClusterName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusterTopologyIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusterTopologyIssues`: DmrClusterTopologyIssuesResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusterTopologyIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClusterTopologyIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterTopologyIssuesResponse**](DmrClusterTopologyIssuesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmrClusters

> DmrClustersResponse GetDmrClusters(ctx).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Cluster objects.



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
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusters(context.Background()).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.GetDmrClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDmrClusters`: DmrClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.GetDmrClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDmrClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClustersResponse**](DmrClustersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

