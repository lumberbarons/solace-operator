# \CertAuthorityApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertAuthority**](CertAuthorityApi.md#CreateCertAuthority) | **Post** /certAuthorities | Create a Certificate Authority object.
[**CreateCertAuthorityOcspTlsTrustedCommonName**](CertAuthorityApi.md#CreateCertAuthorityOcspTlsTrustedCommonName) | **Post** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Create an OCSP Responder Trusted Common Name object.
[**DeleteCertAuthority**](CertAuthorityApi.md#DeleteCertAuthority) | **Delete** /certAuthorities/{certAuthorityName} | Delete a Certificate Authority object.
[**DeleteCertAuthorityOcspTlsTrustedCommonName**](CertAuthorityApi.md#DeleteCertAuthorityOcspTlsTrustedCommonName) | **Delete** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Delete an OCSP Responder Trusted Common Name object.
[**GetCertAuthorities**](CertAuthorityApi.md#GetCertAuthorities) | **Get** /certAuthorities | Get a list of Certificate Authority objects.
[**GetCertAuthority**](CertAuthorityApi.md#GetCertAuthority) | **Get** /certAuthorities/{certAuthorityName} | Get a Certificate Authority object.
[**GetCertAuthorityOcspTlsTrustedCommonName**](CertAuthorityApi.md#GetCertAuthorityOcspTlsTrustedCommonName) | **Get** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames/{ocspTlsTrustedCommonName} | Get an OCSP Responder Trusted Common Name object.
[**GetCertAuthorityOcspTlsTrustedCommonNames**](CertAuthorityApi.md#GetCertAuthorityOcspTlsTrustedCommonNames) | **Get** /certAuthorities/{certAuthorityName}/ocspTlsTrustedCommonNames | Get a list of OCSP Responder Trusted Common Name objects.
[**ReplaceCertAuthority**](CertAuthorityApi.md#ReplaceCertAuthority) | **Put** /certAuthorities/{certAuthorityName} | Replace a Certificate Authority object.
[**UpdateCertAuthority**](CertAuthorityApi.md#UpdateCertAuthority) | **Patch** /certAuthorities/{certAuthorityName} | Update a Certificate Authority object.



## CreateCertAuthority

> CertAuthorityResponse CreateCertAuthority(ctx).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Certificate Authority object.



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
    body := *openapiclient.NewCertAuthority() // CertAuthority | The Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertAuthorityApi.CreateCertAuthority(context.Background()).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertAuthorityApi.CreateCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertAuthority`: CertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertAuthorityApi.CreateCertAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CertAuthority**](CertAuthority.md) | The Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertAuthorityOcspTlsTrustedCommonName

> CertAuthorityOcspTlsTrustedCommonNameResponse CreateCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create an OCSP Responder Trusted Common Name object.



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
    body := *openapiclient.NewCertAuthorityOcspTlsTrustedCommonName() // CertAuthorityOcspTlsTrustedCommonName | The OCSP Responder Trusted Common Name object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertAuthorityApi.CreateCertAuthorityOcspTlsTrustedCommonName(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertAuthorityApi.CreateCertAuthorityOcspTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertAuthorityOcspTlsTrustedCommonName`: CertAuthorityOcspTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `CertAuthorityApi.CreateCertAuthorityOcspTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertAuthorityOcspTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CertAuthorityOcspTlsTrustedCommonName**](CertAuthorityOcspTlsTrustedCommonName.md) | The OCSP Responder Trusted Common Name object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityOcspTlsTrustedCommonNameResponse**](CertAuthorityOcspTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertAuthority

> SempMetaOnlyResponse DeleteCertAuthority(ctx, certAuthorityName).Execute()

Delete a Certificate Authority object.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertAuthorityApi.DeleteCertAuthority(context.Background(), certAuthorityName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertAuthorityApi.DeleteCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCertAuthority`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `CertAuthorityApi.DeleteCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertAuthorityRequest struct via the builder pattern


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


## DeleteCertAuthorityOcspTlsTrustedCommonName

> SempMetaOnlyResponse DeleteCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName, ocspTlsTrustedCommonName).Execute()

Delete an OCSP Responder Trusted Common Name object.



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
    ocspTlsTrustedCommonName := "ocspTlsTrustedCommonName_example" // string | The expected Trusted Common Name of the OCSP responder remote certificate.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertAuthorityApi.DeleteCertAuthorityOcspTlsTrustedCommonName(context.Background(), certAuthorityName, ocspTlsTrustedCommonName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertAuthorityApi.DeleteCertAuthorityOcspTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCertAuthorityOcspTlsTrustedCommonName`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `CertAuthorityApi.DeleteCertAuthorityOcspTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 
**ocspTlsTrustedCommonName** | **string** | The expected Trusted Common Name of the OCSP responder remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertAuthorityOcspTlsTrustedCommonNameRequest struct via the builder pattern


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


## GetCertAuthorities

> CertAuthoritiesResponse GetCertAuthorities(ctx).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Certificate Authority objects.



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
    resp, r, err := api_client.CertAuthorityApi.GetCertAuthorities(context.Background()).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertAuthorityApi.GetCertAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertAuthorities`: CertAuthoritiesResponse
    fmt.Fprintf(os.Stdout, "Response from `CertAuthorityApi.GetCertAuthorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthoritiesResponse**](CertAuthoritiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertAuthority

> CertAuthorityResponse GetCertAuthority(ctx, certAuthorityName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Certificate Authority object.



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
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertAuthorityApi.GetCertAuthority(context.Background(), certAuthorityName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertAuthorityApi.GetCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertAuthority`: CertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertAuthorityApi.GetCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertAuthorityOcspTlsTrustedCommonName

> CertAuthorityOcspTlsTrustedCommonNameResponse GetCertAuthorityOcspTlsTrustedCommonName(ctx, certAuthorityName, ocspTlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get an OCSP Responder Trusted Common Name object.



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
    ocspTlsTrustedCommonName := "ocspTlsTrustedCommonName_example" // string | The expected Trusted Common Name of the OCSP responder remote certificate.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertAuthorityApi.GetCertAuthorityOcspTlsTrustedCommonName(context.Background(), certAuthorityName, ocspTlsTrustedCommonName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertAuthorityApi.GetCertAuthorityOcspTlsTrustedCommonName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertAuthorityOcspTlsTrustedCommonName`: CertAuthorityOcspTlsTrustedCommonNameResponse
    fmt.Fprintf(os.Stdout, "Response from `CertAuthorityApi.GetCertAuthorityOcspTlsTrustedCommonName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 
**ocspTlsTrustedCommonName** | **string** | The expected Trusted Common Name of the OCSP responder remote certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertAuthorityOcspTlsTrustedCommonNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityOcspTlsTrustedCommonNameResponse**](CertAuthorityOcspTlsTrustedCommonNameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertAuthorityOcspTlsTrustedCommonNames

> CertAuthorityOcspTlsTrustedCommonNamesResponse GetCertAuthorityOcspTlsTrustedCommonNames(ctx, certAuthorityName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of OCSP Responder Trusted Common Name objects.



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
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertAuthorityApi.GetCertAuthorityOcspTlsTrustedCommonNames(context.Background(), certAuthorityName).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertAuthorityApi.GetCertAuthorityOcspTlsTrustedCommonNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertAuthorityOcspTlsTrustedCommonNames`: CertAuthorityOcspTlsTrustedCommonNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `CertAuthorityApi.GetCertAuthorityOcspTlsTrustedCommonNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertAuthorityOcspTlsTrustedCommonNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityOcspTlsTrustedCommonNamesResponse**](CertAuthorityOcspTlsTrustedCommonNamesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCertAuthority

> CertAuthorityResponse ReplaceCertAuthority(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Certificate Authority object.



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
    body := *openapiclient.NewCertAuthority() // CertAuthority | The Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertAuthorityApi.ReplaceCertAuthority(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertAuthorityApi.ReplaceCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCertAuthority`: CertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertAuthorityApi.ReplaceCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CertAuthority**](CertAuthority.md) | The Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertAuthority

> CertAuthorityResponse UpdateCertAuthority(ctx, certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Certificate Authority object.



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
    body := *openapiclient.NewCertAuthority() // CertAuthority | The Certificate Authority object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertAuthorityApi.UpdateCertAuthority(context.Background(), certAuthorityName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertAuthorityApi.UpdateCertAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertAuthority`: CertAuthorityResponse
    fmt.Fprintf(os.Stdout, "Response from `CertAuthorityApi.UpdateCertAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certAuthorityName** | **string** | The name of the Certificate Authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CertAuthority**](CertAuthority.md) | The Certificate Authority object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**CertAuthorityResponse**](CertAuthorityResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

