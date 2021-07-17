# \ClientProfileApi

All URIs are relative to *http://www.solace.com/SEMP/v2/config*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsgVpnClientProfile**](ClientProfileApi.md#CreateMsgVpnClientProfile) | **Post** /msgVpns/{msgVpnName}/clientProfiles | Create a Client Profile object.
[**DeleteMsgVpnClientProfile**](ClientProfileApi.md#DeleteMsgVpnClientProfile) | **Delete** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Delete a Client Profile object.
[**GetMsgVpnClientProfile**](ClientProfileApi.md#GetMsgVpnClientProfile) | **Get** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Get a Client Profile object.
[**GetMsgVpnClientProfiles**](ClientProfileApi.md#GetMsgVpnClientProfiles) | **Get** /msgVpns/{msgVpnName}/clientProfiles | Get a list of Client Profile objects.
[**ReplaceMsgVpnClientProfile**](ClientProfileApi.md#ReplaceMsgVpnClientProfile) | **Put** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Replace a Client Profile object.
[**UpdateMsgVpnClientProfile**](ClientProfileApi.md#UpdateMsgVpnClientProfile) | **Patch** /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName} | Update a Client Profile object.



## CreateMsgVpnClientProfile

> MsgVpnClientProfileResponse CreateMsgVpnClientProfile(ctx, msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Create a Client Profile object.



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
    body := *openapiclient.NewMsgVpnClientProfile() // MsgVpnClientProfile | The Client Profile object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientProfileApi.CreateMsgVpnClientProfile(context.Background(), msgVpnName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfileApi.CreateMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMsgVpnClientProfile`: MsgVpnClientProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfileApi.CreateMsgVpnClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMsgVpnClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md) | The Client Profile object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMsgVpnClientProfile

> SempMetaOnlyResponse DeleteMsgVpnClientProfile(ctx, msgVpnName, clientProfileName).Execute()

Delete a Client Profile object.



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
    clientProfileName := "clientProfileName_example" // string | The name of the Client Profile.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientProfileApi.DeleteMsgVpnClientProfile(context.Background(), msgVpnName, clientProfileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfileApi.DeleteMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMsgVpnClientProfile`: SempMetaOnlyResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfileApi.DeleteMsgVpnClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientProfileName** | **string** | The name of the Client Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMsgVpnClientProfileRequest struct via the builder pattern


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


## GetMsgVpnClientProfile

> MsgVpnClientProfileResponse GetMsgVpnClientProfile(ctx, msgVpnName, clientProfileName).OpaquePassword(opaquePassword).Select_(select_).Execute()

Get a Client Profile object.



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
    clientProfileName := "clientProfileName_example" // string | The name of the Client Profile.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientProfileApi.GetMsgVpnClientProfile(context.Background(), msgVpnName, clientProfileName).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfileApi.GetMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientProfile`: MsgVpnClientProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfileApi.GetMsgVpnClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientProfileName** | **string** | The name of the Client Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientProfiles

> MsgVpnClientProfilesResponse GetMsgVpnClientProfiles(ctx, msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()

Get a list of Client Profile objects.



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
    resp, r, err := api_client.ClientProfileApi.GetMsgVpnClientProfiles(context.Background(), msgVpnName).Count(count).Cursor(cursor).OpaquePassword(opaquePassword).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfileApi.GetMsgVpnClientProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientProfiles`: MsgVpnClientProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfileApi.GetMsgVpnClientProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfilesResponse**](MsgVpnClientProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceMsgVpnClientProfile

> MsgVpnClientProfileResponse ReplaceMsgVpnClientProfile(ctx, msgVpnName, clientProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Replace a Client Profile object.



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
    clientProfileName := "clientProfileName_example" // string | The name of the Client Profile.
    body := *openapiclient.NewMsgVpnClientProfile() // MsgVpnClientProfile | The Client Profile object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientProfileApi.ReplaceMsgVpnClientProfile(context.Background(), msgVpnName, clientProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfileApi.ReplaceMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceMsgVpnClientProfile`: MsgVpnClientProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfileApi.ReplaceMsgVpnClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientProfileName** | **string** | The name of the Client Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceMsgVpnClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md) | The Client Profile object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMsgVpnClientProfile

> MsgVpnClientProfileResponse UpdateMsgVpnClientProfile(ctx, msgVpnName, clientProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()

Update a Client Profile object.



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
    clientProfileName := "clientProfileName_example" // string | The name of the Client Profile.
    body := *openapiclient.NewMsgVpnClientProfile() // MsgVpnClientProfile | The Client Profile object's attributes.
    opaquePassword := "opaquePassword_example" // string | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientProfileApi.UpdateMsgVpnClientProfile(context.Background(), msgVpnName, clientProfileName).Body(body).OpaquePassword(opaquePassword).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientProfileApi.UpdateMsgVpnClientProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsgVpnClientProfile`: MsgVpnClientProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientProfileApi.UpdateMsgVpnClientProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientProfileName** | **string** | The name of the Client Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsgVpnClientProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**MsgVpnClientProfile**](MsgVpnClientProfile.md) | The Client Profile object&#39;s attributes. | 
 **opaquePassword** | **string** | Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the &#x60;opaquePassword&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientProfileResponse**](MsgVpnClientProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

