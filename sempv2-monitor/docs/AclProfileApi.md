# \AclProfileApi

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnAclProfile**](AclProfileApi.md#GetMsgVpnAclProfile) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName} | Get an ACL Profile object.
[**GetMsgVpnAclProfileClientConnectException**](AclProfileApi.md#GetMsgVpnAclProfileClientConnectException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions/{clientConnectExceptionAddress} | Get a Client Connect Exception object.
[**GetMsgVpnAclProfileClientConnectExceptions**](AclProfileApi.md#GetMsgVpnAclProfileClientConnectExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/clientConnectExceptions | Get a list of Client Connect Exception objects.
[**GetMsgVpnAclProfilePublishException**](AclProfileApi.md#GetMsgVpnAclProfilePublishException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic} | Get a Publish Topic Exception object.
[**GetMsgVpnAclProfilePublishExceptions**](AclProfileApi.md#GetMsgVpnAclProfilePublishExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions | Get a list of Publish Topic Exception objects.
[**GetMsgVpnAclProfilePublishTopicException**](AclProfileApi.md#GetMsgVpnAclProfilePublishTopicException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions/{publishTopicExceptionSyntax},{publishTopicException} | Get a Publish Topic Exception object.
[**GetMsgVpnAclProfilePublishTopicExceptions**](AclProfileApi.md#GetMsgVpnAclProfilePublishTopicExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions | Get a list of Publish Topic Exception objects.
[**GetMsgVpnAclProfileSubscribeException**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions/{topicSyntax},{subscribeExceptionTopic} | Get a Subscribe Topic Exception object.
[**GetMsgVpnAclProfileSubscribeExceptions**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions | Get a list of Subscribe Topic Exception objects.
[**GetMsgVpnAclProfileSubscribeShareNameException**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeShareNameException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions/{subscribeShareNameExceptionSyntax},{subscribeShareNameException} | Get a Subscribe Share Name Exception object.
[**GetMsgVpnAclProfileSubscribeShareNameExceptions**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeShareNameExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeShareNameExceptions | Get a list of Subscribe Share Name Exception objects.
[**GetMsgVpnAclProfileSubscribeTopicException**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeTopicException) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions/{subscribeTopicExceptionSyntax},{subscribeTopicException} | Get a Subscribe Topic Exception object.
[**GetMsgVpnAclProfileSubscribeTopicExceptions**](AclProfileApi.md#GetMsgVpnAclProfileSubscribeTopicExceptions) | **Get** /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeTopicExceptions | Get a list of Subscribe Topic Exception objects.
[**GetMsgVpnAclProfiles**](AclProfileApi.md#GetMsgVpnAclProfiles) | **Get** /msgVpns/{msgVpnName}/aclProfiles | Get a list of ACL Profile objects.



## GetMsgVpnAclProfile

> MsgVpnAclProfileResponse GetMsgVpnAclProfile(ctx, msgVpnName, aclProfileName).Select_(select_).Execute()

Get an ACL Profile object.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfile(context.Background(), msgVpnName, aclProfileName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfile`: MsgVpnAclProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileResponse**](MsgVpnAclProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileClientConnectException

> MsgVpnAclProfileClientConnectExceptionResponse GetMsgVpnAclProfileClientConnectException(ctx, msgVpnName, aclProfileName, clientConnectExceptionAddress).Select_(select_).Execute()

Get a Client Connect Exception object.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    clientConnectExceptionAddress := "clientConnectExceptionAddress_example" // string | The IP address/netmask of the client connect exception in CIDR form.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfileClientConnectException(context.Background(), msgVpnName, aclProfileName, clientConnectExceptionAddress).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfileClientConnectException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileClientConnectException`: MsgVpnAclProfileClientConnectExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfileClientConnectException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**clientConnectExceptionAddress** | **string** | The IP address/netmask of the client connect exception in CIDR form. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileClientConnectExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionResponse**](MsgVpnAclProfileClientConnectExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileClientConnectExceptions

> MsgVpnAclProfileClientConnectExceptionsResponse GetMsgVpnAclProfileClientConnectExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Connect Exception objects.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfileClientConnectExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfileClientConnectExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileClientConnectExceptions`: MsgVpnAclProfileClientConnectExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfileClientConnectExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileClientConnectExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileClientConnectExceptionsResponse**](MsgVpnAclProfileClientConnectExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfilePublishException

> MsgVpnAclProfilePublishExceptionResponse GetMsgVpnAclProfilePublishException(ctx, msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic).Select_(select_).Execute()

Get a Publish Topic Exception object.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    topicSyntax := "topicSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    publishExceptionTopic := "publishExceptionTopic_example" // string | The topic for the exception to the default action taken. May include wildcard characters.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfilePublishException(context.Background(), msgVpnName, aclProfileName, topicSyntax, publishExceptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfilePublishException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishException`: MsgVpnAclProfilePublishExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfilePublishException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**topicSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**publishExceptionTopic** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfilePublishExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionResponse**](MsgVpnAclProfilePublishExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfilePublishExceptions

> MsgVpnAclProfilePublishExceptionsResponse GetMsgVpnAclProfilePublishExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Publish Topic Exception objects.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfilePublishExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfilePublishExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishExceptions`: MsgVpnAclProfilePublishExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfilePublishExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfilePublishExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishExceptionsResponse**](MsgVpnAclProfilePublishExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfilePublishTopicException

> MsgVpnAclProfilePublishTopicExceptionResponse GetMsgVpnAclProfilePublishTopicException(ctx, msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException).Select_(select_).Execute()

Get a Publish Topic Exception object.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    publishTopicExceptionSyntax := "publishTopicExceptionSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    publishTopicException := "publishTopicException_example" // string | The topic for the exception to the default action taken. May include wildcard characters.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfilePublishTopicException(context.Background(), msgVpnName, aclProfileName, publishTopicExceptionSyntax, publishTopicException).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfilePublishTopicException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishTopicException`: MsgVpnAclProfilePublishTopicExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfilePublishTopicException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**publishTopicExceptionSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**publishTopicException** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfilePublishTopicExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionResponse**](MsgVpnAclProfilePublishTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfilePublishTopicExceptions

> MsgVpnAclProfilePublishTopicExceptionsResponse GetMsgVpnAclProfilePublishTopicExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Publish Topic Exception objects.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfilePublishTopicExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfilePublishTopicExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfilePublishTopicExceptions`: MsgVpnAclProfilePublishTopicExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfilePublishTopicExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfilePublishTopicExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilePublishTopicExceptionsResponse**](MsgVpnAclProfilePublishTopicExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeException

> MsgVpnAclProfileSubscribeExceptionResponse GetMsgVpnAclProfileSubscribeException(ctx, msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic).Select_(select_).Execute()

Get a Subscribe Topic Exception object.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    topicSyntax := "topicSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    subscribeExceptionTopic := "subscribeExceptionTopic_example" // string | The topic for the exception to the default action taken. May include wildcard characters.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfileSubscribeException(context.Background(), msgVpnName, aclProfileName, topicSyntax, subscribeExceptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfileSubscribeException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeException`: MsgVpnAclProfileSubscribeExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfileSubscribeException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**topicSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**subscribeExceptionTopic** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionResponse**](MsgVpnAclProfileSubscribeExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeExceptions

> MsgVpnAclProfileSubscribeExceptionsResponse GetMsgVpnAclProfileSubscribeExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Subscribe Topic Exception objects.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfileSubscribeExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfileSubscribeExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeExceptions`: MsgVpnAclProfileSubscribeExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfileSubscribeExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeExceptionsResponse**](MsgVpnAclProfileSubscribeExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeShareNameException

> MsgVpnAclProfileSubscribeShareNameExceptionResponse GetMsgVpnAclProfileSubscribeShareNameException(ctx, msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException).Select_(select_).Execute()

Get a Subscribe Share Name Exception object.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    subscribeShareNameExceptionSyntax := "subscribeShareNameExceptionSyntax_example" // string | The syntax of the subscribe share name for the exception to the default action taken.
    subscribeShareNameException := "subscribeShareNameException_example" // string | The subscribe share name exception to the default action taken. May include wildcard characters.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfileSubscribeShareNameException(context.Background(), msgVpnName, aclProfileName, subscribeShareNameExceptionSyntax, subscribeShareNameException).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfileSubscribeShareNameException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeShareNameException`: MsgVpnAclProfileSubscribeShareNameExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfileSubscribeShareNameException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**subscribeShareNameExceptionSyntax** | **string** | The syntax of the subscribe share name for the exception to the default action taken. | 
**subscribeShareNameException** | **string** | The subscribe share name exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeShareNameExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionResponse**](MsgVpnAclProfileSubscribeShareNameExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeShareNameExceptions

> MsgVpnAclProfileSubscribeShareNameExceptionsResponse GetMsgVpnAclProfileSubscribeShareNameExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Subscribe Share Name Exception objects.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfileSubscribeShareNameExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfileSubscribeShareNameExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeShareNameExceptions`: MsgVpnAclProfileSubscribeShareNameExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfileSubscribeShareNameExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeShareNameExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeShareNameExceptionsResponse**](MsgVpnAclProfileSubscribeShareNameExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeTopicException

> MsgVpnAclProfileSubscribeTopicExceptionResponse GetMsgVpnAclProfileSubscribeTopicException(ctx, msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException).Select_(select_).Execute()

Get a Subscribe Topic Exception object.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    subscribeTopicExceptionSyntax := "subscribeTopicExceptionSyntax_example" // string | The syntax of the topic for the exception to the default action taken.
    subscribeTopicException := "subscribeTopicException_example" // string | The topic for the exception to the default action taken. May include wildcard characters.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfileSubscribeTopicException(context.Background(), msgVpnName, aclProfileName, subscribeTopicExceptionSyntax, subscribeTopicException).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfileSubscribeTopicException``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeTopicException`: MsgVpnAclProfileSubscribeTopicExceptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfileSubscribeTopicException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 
**subscribeTopicExceptionSyntax** | **string** | The syntax of the topic for the exception to the default action taken. | 
**subscribeTopicException** | **string** | The topic for the exception to the default action taken. May include wildcard characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeTopicExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionResponse**](MsgVpnAclProfileSubscribeTopicExceptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfileSubscribeTopicExceptions

> MsgVpnAclProfileSubscribeTopicExceptionsResponse GetMsgVpnAclProfileSubscribeTopicExceptions(ctx, msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Subscribe Topic Exception objects.



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
    aclProfileName := "aclProfileName_example" // string | The name of the ACL Profile.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfileSubscribeTopicExceptions(context.Background(), msgVpnName, aclProfileName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfileSubscribeTopicExceptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfileSubscribeTopicExceptions`: MsgVpnAclProfileSubscribeTopicExceptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfileSubscribeTopicExceptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**aclProfileName** | **string** | The name of the ACL Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfileSubscribeTopicExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfileSubscribeTopicExceptionsResponse**](MsgVpnAclProfileSubscribeTopicExceptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnAclProfiles

> MsgVpnAclProfilesResponse GetMsgVpnAclProfiles(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of ACL Profile objects.



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
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AclProfileApi.GetMsgVpnAclProfiles(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AclProfileApi.GetMsgVpnAclProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnAclProfiles`: MsgVpnAclProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `AclProfileApi.GetMsgVpnAclProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnAclProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnAclProfilesResponse**](MsgVpnAclProfilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

