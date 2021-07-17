# \StandardDomainCertAuthorityApi

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStandardDomainCertAuthorities**](StandardDomainCertAuthorityApi.md#GetStandardDomainCertAuthorities) | **Get** /standardDomainCertAuthorities | Get a list of Standard Domain Certificate Authority objects.
[**GetStandardDomainCertAuthority**](StandardDomainCertAuthorityApi.md#GetStandardDomainCertAuthority) | **Get** /standardDomainCertAuthorities/{certAuthorityName} | Get a Standard Domain Certificate Authority object.



## GetStandardDomainCertAuthorities

> StandardDomainCertAuthoritiesResponse GetStandardDomainCertAuthorities(ctx).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Standard Domain Certificate Authority objects.



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
    resp, r, err := api_client.StandardDomainCertAuthorityApi.GetStandardDomainCertAuthorities(context.Background()).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandardDomainCertAuthorityApi.GetStandardDomainCertAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStandardDomainCertAuthorities`: StandardDomainCertAuthoritiesResponse
    fmt.Fprintf(os.Stdout, "Response from `StandardDomainCertAuthorityApi.GetStandardDomainCertAuthorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStandardDomainCertAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**StandardDomainCertAuthoritiesResponse**](StandardDomainCertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStandardDomainCertAuthority

> StandardDomainCertAuthorityResponse GetStandardDomainCertAuthority(ctx, certAuthorityName).Select_(select_).Execute()

Get a Standard Domain Certificate Authority object.



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
    certAuthorityName := "certAuthorityName_example" // string | The name of the Certificate Authority.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StandardDomainCertAuthorityApi.GetStandardDomainCertAuthority(context.Background(), certAuthorityName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandardDomainCertAuthorityApi.GetStandardDomainCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStandardDomainCertAuthority`: StandardDomainCertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `StandardDomainCertAuthorityApi.GetStandardDomainCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStandardDomainCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**StandardDomainCertAuthorityResponse**](StandardDomainCertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

