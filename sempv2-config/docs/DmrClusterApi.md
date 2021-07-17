# \DmrClusterApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDmrCluster**](DmrClusterApi.md#CreateDmrCluster) | **Post** /dmrClusters | Create a Cluster object.
[**CreateDmrClusterLink**](DmrClusterApi.md#CreateDmrClusterLink) | **Post** /dmrClusters/{dmrClusterName}/links | Create a Link object.
[**CreateDmrClusterLinkRemoteAddress**](DmrClusterApi.md#CreateDmrClusterLinkRemoteAddress) | **Post** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses | Create a Remote Address object.
[**CreateDmrClusterLinkTlsTrustedCommonName**](DmrClusterApi.md#CreateDmrClusterLinkTlsTrustedCommonName) | **Post** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames | Create a Trusted Common Name object.
[**DeleteDmrCluster**](DmrClusterApi.md#DeleteDmrCluster) | **Delete** /dmrClusters/{dmrClusterName} | Delete a Cluster object.
[**DeleteDmrClusterLink**](DmrClusterApi.md#DeleteDmrClusterLink) | **Delete** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Delete a Link object.
[**DeleteDmrClusterLinkRemoteAddress**](DmrClusterApi.md#DeleteDmrClusterLinkRemoteAddress) | **Delete** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses/{remoteAddress} | Delete a Remote Address object.
[**DeleteDmrClusterLinkTlsTrustedCommonName**](DmrClusterApi.md#DeleteDmrClusterLinkTlsTrustedCommonName) | **Delete** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Delete a Trusted Common Name object.
[**GetDmrCluster**](DmrClusterApi.md#GetDmrCluster) | **Get** /dmrClusters/{dmrClusterName} | Get a Cluster object.
[**GetDmrClusterLink**](DmrClusterApi.md#GetDmrClusterLink) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Get a Link object.
[**GetDmrClusterLinkRemoteAddress**](DmrClusterApi.md#GetDmrClusterLinkRemoteAddress) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses/{remoteAddress} | Get a Remote Address object.
[**GetDmrClusterLinkRemoteAddresses**](DmrClusterApi.md#GetDmrClusterLinkRemoteAddresses) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/remoteAddresses | Get a list of Remote Address objects.
[**GetDmrClusterLinkTlsTrustedCommonName**](DmrClusterApi.md#GetDmrClusterLinkTlsTrustedCommonName) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName} | Get a Trusted Common Name object.
[**GetDmrClusterLinkTlsTrustedCommonNames**](DmrClusterApi.md#GetDmrClusterLinkTlsTrustedCommonNames) | **Get** /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames | Get a list of Trusted Common Name objects.
[**GetDmrClusterLinks**](DmrClusterApi.md#GetDmrClusterLinks) | **Get** /dmrClusters/{dmrClusterName}/links | Get a list of Link objects.
[**GetDmrClusters**](DmrClusterApi.md#GetDmrClusters) | **Get** /dmrClusters | Get a list of Cluster objects.
[**ReplaceDmrCluster**](DmrClusterApi.md#ReplaceDmrCluster) | **Put** /dmrClusters/{dmrClusterName} | Replace a Cluster object.
[**ReplaceDmrClusterLink**](DmrClusterApi.md#ReplaceDmrClusterLink) | **Put** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Replace a Link object.
[**UpdateDmrCluster**](DmrClusterApi.md#UpdateDmrCluster) | **Patch** /dmrClusters/{dmrClusterName} | Update a Cluster object.
[**UpdateDmrClusterLink**](DmrClusterApi.md#UpdateDmrClusterLink) | **Patch** /dmrClusters/{dmrClusterName}/links/{remoteNodeName} | Update a Link object.



## CreateDmrCluster

> DmrClusterResponse CreateDmrCluster(ctx).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Cluster object.



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
    body := *openapiclient.NewDmrCluster() // DmrCluster | The Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.CreateDmrCluster(context.Background()).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.CreateDmrCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmrCluster`: DmrClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.CreateDmrCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmrClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DmrCluster**](DmrCluster.md) | The Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDmrClusterLink

> DmrClusterLinkResponse CreateDmrClusterLink(ctx, dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Link object.



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
    body := *openapiclient.NewDmrClusterLink() // DmrClusterLink | The Link object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.CreateDmrClusterLink(context.Background(), dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.CreateDmrClusterLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmrClusterLink`: DmrClusterLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.CreateDmrClusterLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmrClusterLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DmrClusterLink**](DmrClusterLink.md) | The Link object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDmrClusterLinkRemoteAddress

> DmrClusterLinkRemoteAddressResponse CreateDmrClusterLinkRemoteAddress(ctx, dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Remote Address object.



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
    body := *openapiclient.NewDmrClusterLinkRemoteAddress() // DmrClusterLinkRemoteAddress | The Remote Address object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.CreateDmrClusterLinkRemoteAddress(context.Background(), dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.CreateDmrClusterLinkRemoteAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmrClusterLinkRemoteAddress`: DmrClusterLinkRemoteAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.CreateDmrClusterLinkRemoteAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmrClusterLinkRemoteAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DmrClusterLinkRemoteAddress**](DmrClusterLinkRemoteAddress.md) | The Remote Address object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkRemoteAddressResponse**](DmrClusterLinkRemoteAddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDmrClusterLinkTlsTrustedCommonName

> DmrClusterLinkTlsTrustedCommonNameResponse CreateDmrClusterLinkTlsTrustedCommonName(ctx, dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    body := *openapiclient.NewDmrClusterLinkTlsTrustedCommonName() // DmrClusterLinkTlsTrustedCommonName | The Trusted Common Name object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.CreateDmrClusterLinkTlsTrustedCommonName(context.Background(), dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.CreateDmrClusterLinkTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDmrClusterLinkTlsTrustedCommonName`: DmrClusterLinkTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.CreateDmrClusterLinkTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDmrClusterLinkTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DmrClusterLinkTlsTrustedCommonName**](DmrClusterLinkTlsTrustedCommonName.md) | The Trusted Common Name object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkTlsTrustedCommonNameResponse**](DmrClusterLinkTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDmrCluster

> SempMetaOnlyResponse DeleteDmrCluster(ctx, dmrClusterName).Execute()

Delete a Cluster object.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.DeleteDmrCluster(context.Background(), dmrClusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.DeleteDmrCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDmrCluster`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.DeleteDmrCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmrClusterRequest struct via the builder pattern


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


## DeleteDmrClusterLink

> SempMetaOnlyResponse DeleteDmrClusterLink(ctx, dmrClusterName, remoteNodeName).Execute()

Delete a Link object.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.DeleteDmrClusterLink(context.Background(), dmrClusterName, remoteNodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.DeleteDmrClusterLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDmrClusterLink`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.DeleteDmrClusterLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmrClusterLinkRequest struct via the builder pattern


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


## DeleteDmrClusterLinkRemoteAddress

> SempMetaOnlyResponse DeleteDmrClusterLinkRemoteAddress(ctx, dmrClusterName, remoteNodeName, remoteAddress).Execute()

Delete a Remote Address object.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.DeleteDmrClusterLinkRemoteAddress(context.Background(), dmrClusterName, remoteNodeName, remoteAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.DeleteDmrClusterLinkRemoteAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDmrClusterLinkRemoteAddress`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.DeleteDmrClusterLinkRemoteAddress`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteDmrClusterLinkRemoteAddressRequest struct via the builder pattern


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


## DeleteDmrClusterLinkTlsTrustedCommonName

> SempMetaOnlyResponse DeleteDmrClusterLinkTlsTrustedCommonName(ctx, dmrClusterName, remoteNodeName, tlsTrustedCommonName).Execute()

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
    dmrClusterName := "dmrClusterName_example" // string | The name of the Cluster.
    remoteNodeName := "remoteNodeName_example" // string | The name of the node at the remote end of the Link.
    tlsTrustedCommonName := "tlsTrustedCommonName_example" // string | The expected trusted common name of the remote certificate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.DeleteDmrClusterLinkTlsTrustedCommonName(context.Background(), dmrClusterName, remoteNodeName, tlsTrustedCommonName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.DeleteDmrClusterLinkTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDmrClusterLinkTlsTrustedCommonName`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.DeleteDmrClusterLinkTlsTrustedCommonName`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteDmrClusterLinkTlsTrustedCommonNameRequest struct via the builder pattern


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


## GetDmrCluster

> DmrClusterResponse GetDmrCluster(ctx, dmrClusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()

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
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrCluster(context.Background(), dmrClusterName).OpaquePassword(opaquePassword).Select_(select_).Execute()
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

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

> DmrClusterLinkResponse GetDmrClusterLink(ctx, dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Select_(select_).Execute()

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
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLink(context.Background(), dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Select_(select_).Execute()
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


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
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


## GetDmrClusterLinkRemoteAddress

> DmrClusterLinkRemoteAddressResponse GetDmrClusterLinkRemoteAddress(ctx, dmrClusterName, remoteNodeName, remoteAddress).OpaquePassword(opaquePassword).Select_(select_).Execute()

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
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinkRemoteAddress(context.Background(), dmrClusterName, remoteNodeName, remoteAddress).OpaquePassword(opaquePassword).Select_(select_).Execute()
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



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

> DmrClusterLinkRemoteAddressesResponse GetDmrClusterLinkRemoteAddresses(ctx, dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

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
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinkRemoteAddresses(context.Background(), dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
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


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

> DmrClusterLinkTlsTrustedCommonNameResponse GetDmrClusterLinkTlsTrustedCommonName(ctx, dmrClusterName, remoteNodeName, tlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()

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
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinkTlsTrustedCommonName(context.Background(), dmrClusterName, remoteNodeName, tlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()
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



 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

> DmrClusterLinkTlsTrustedCommonNamesResponse GetDmrClusterLinkTlsTrustedCommonNames(ctx, dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

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
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinkTlsTrustedCommonNames(context.Background(), dmrClusterName, remoteNodeName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
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


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
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

> DmrClusterLinksResponse GetDmrClusterLinks(ctx, dmrClusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

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
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusterLinks(context.Background(), dmrClusterName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
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
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
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


## GetDmrClusters

> DmrClustersResponse GetDmrClusters(ctx).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

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
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.GetDmrClusters(context.Background()).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
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
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
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


## ReplaceDmrCluster

> DmrClusterResponse ReplaceDmrCluster(ctx, dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Cluster object.



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
    body := *openapiclient.NewDmrCluster() // DmrCluster | The Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.ReplaceDmrCluster(context.Background(), dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.ReplaceDmrCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceDmrCluster`: DmrClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.ReplaceDmrCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDmrClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DmrCluster**](DmrCluster.md) | The Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDmrClusterLink

> DmrClusterLinkResponse ReplaceDmrClusterLink(ctx, dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Link object.



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
    body := *openapiclient.NewDmrClusterLink() // DmrClusterLink | The Link object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.ReplaceDmrClusterLink(context.Background(), dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.ReplaceDmrClusterLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceDmrClusterLink`: DmrClusterLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.ReplaceDmrClusterLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDmrClusterLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DmrClusterLink**](DmrClusterLink.md) | The Link object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDmrCluster

> DmrClusterResponse UpdateDmrCluster(ctx, dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Cluster object.



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
    body := *openapiclient.NewDmrCluster() // DmrCluster | The Cluster object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.UpdateDmrCluster(context.Background(), dmrClusterName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.UpdateDmrCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDmrCluster`: DmrClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.UpdateDmrCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDmrClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DmrCluster**](DmrCluster.md) | The Cluster object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterResponse**](DmrClusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDmrClusterLink

> DmrClusterLinkResponse UpdateDmrClusterLink(ctx, dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Link object.



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
    body := *openapiclient.NewDmrClusterLink() // DmrClusterLink | The Link object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DmrClusterApi.UpdateDmrClusterLink(context.Background(), dmrClusterName, remoteNodeName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DmrClusterApi.UpdateDmrClusterLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDmrClusterLink`: DmrClusterLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `DmrClusterApi.UpdateDmrClusterLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dmrClusterName** | **string** | The name of the Cluster. | 
**remoteNodeName** | **string** | The name of the node at the remote end of the Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDmrClusterLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DmrClusterLink**](DmrClusterLink.md) | The Link object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**DmrClusterLinkResponse**](DmrClusterLinkResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

