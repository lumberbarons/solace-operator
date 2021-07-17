# \ClientApi

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnClient**](ClientApi.md#GetMsgVpnClient) | **Get** /msgVpns/{msgVpnName}/clients/{clientName} | Get a Client object.
[**GetMsgVpnClientConnection**](ClientApi.md#GetMsgVpnClientConnection) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections/{clientAddress} | Get a Client Connection object.
[**GetMsgVpnClientConnections**](ClientApi.md#GetMsgVpnClientConnections) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/connections | Get a list of Client Connection objects.
[**GetMsgVpnClientRxFlow**](ClientApi.md#GetMsgVpnClientRxFlow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows/{flowId} | Get a Client Receive Flow object.
[**GetMsgVpnClientRxFlows**](ClientApi.md#GetMsgVpnClientRxFlows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/rxFlows | Get a list of Client Receive Flow objects.
[**GetMsgVpnClientSubscription**](ClientApi.md#GetMsgVpnClientSubscription) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions/{subscriptionTopic} | Get a Client Subscription object.
[**GetMsgVpnClientSubscriptions**](ClientApi.md#GetMsgVpnClientSubscriptions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/subscriptions | Get a list of Client Subscription objects.
[**GetMsgVpnClientTransactedSession**](ClientApi.md#GetMsgVpnClientTransactedSession) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions/{sessionName} | Get a Client Transacted Session object.
[**GetMsgVpnClientTransactedSessions**](ClientApi.md#GetMsgVpnClientTransactedSessions) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/transactedSessions | Get a list of Client Transacted Session objects.
[**GetMsgVpnClientTxFlow**](ClientApi.md#GetMsgVpnClientTxFlow) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows/{flowId} | Get a Client Transmit Flow object.
[**GetMsgVpnClientTxFlows**](ClientApi.md#GetMsgVpnClientTxFlows) | **Get** /msgVpns/{msgVpnName}/clients/{clientName}/txFlows | Get a list of Client Transmit Flow objects.
[**GetMsgVpnClients**](ClientApi.md#GetMsgVpnClients) | **Get** /msgVpns/{msgVpnName}/clients | Get a list of Client objects.



## GetMsgVpnClient

> MsgVpnClientResponse GetMsgVpnClient(ctx, msgVpnName, clientName).Select_(select_).Execute()

Get a Client object.



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
    clientName := "clientName_example" // string | The name of the Client.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClient(context.Background(), msgVpnName, clientName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClient`: MsgVpnClientResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientResponse**](MsgVpnClientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientConnection

> MsgVpnClientConnectionResponse GetMsgVpnClientConnection(ctx, msgVpnName, clientName, clientAddress).Select_(select_).Execute()

Get a Client Connection object.



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
    clientName := "clientName_example" // string | The name of the Client.
    clientAddress := "clientAddress_example" // string | The IP address and TCP port on the Client side of the Client Connection.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClientConnection(context.Background(), msgVpnName, clientName, clientAddress).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClientConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientConnection`: MsgVpnClientConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClientConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 
**clientAddress** | **string** | The IP address and TCP port on the Client side of the Client Connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientConnectionResponse**](MsgVpnClientConnectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientConnections

> MsgVpnClientConnectionsResponse GetMsgVpnClientConnections(ctx, msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Connection objects.



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
    clientName := "clientName_example" // string | The name of the Client.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClientConnections(context.Background(), msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClientConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientConnections`: MsgVpnClientConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClientConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientConnectionsResponse**](MsgVpnClientConnectionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientRxFlow

> MsgVpnClientRxFlowResponse GetMsgVpnClientRxFlow(ctx, msgVpnName, clientName, flowId).Select_(select_).Execute()

Get a Client Receive Flow object.



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
    clientName := "clientName_example" // string | The name of the Client.
    flowId := "flowId_example" // string | The identifier (ID) of the flow.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClientRxFlow(context.Background(), msgVpnName, clientName, flowId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClientRxFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientRxFlow`: MsgVpnClientRxFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClientRxFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 
**flowId** | **string** | The identifier (ID) of the flow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientRxFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientRxFlowResponse**](MsgVpnClientRxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientRxFlows

> MsgVpnClientRxFlowsResponse GetMsgVpnClientRxFlows(ctx, msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Receive Flow objects.



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
    clientName := "clientName_example" // string | The name of the Client.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClientRxFlows(context.Background(), msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClientRxFlows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientRxFlows`: MsgVpnClientRxFlowsResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClientRxFlows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientRxFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientRxFlowsResponse**](MsgVpnClientRxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientSubscription

> MsgVpnClientSubscriptionResponse GetMsgVpnClientSubscription(ctx, msgVpnName, clientName, subscriptionTopic).Select_(select_).Execute()

Get a Client Subscription object.



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
    clientName := "clientName_example" // string | The name of the Client.
    subscriptionTopic := "subscriptionTopic_example" // string | The topic of the Subscription.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClientSubscription(context.Background(), msgVpnName, clientName, subscriptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClientSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientSubscription`: MsgVpnClientSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClientSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 
**subscriptionTopic** | **string** | The topic of the Subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientSubscriptionResponse**](MsgVpnClientSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientSubscriptions

> MsgVpnClientSubscriptionsResponse GetMsgVpnClientSubscriptions(ctx, msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Subscription objects.



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
    clientName := "clientName_example" // string | The name of the Client.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClientSubscriptions(context.Background(), msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClientSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientSubscriptions`: MsgVpnClientSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClientSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientSubscriptionsResponse**](MsgVpnClientSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientTransactedSession

> MsgVpnClientTransactedSessionResponse GetMsgVpnClientTransactedSession(ctx, msgVpnName, clientName, sessionName).Select_(select_).Execute()

Get a Client Transacted Session object.



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
    clientName := "clientName_example" // string | The name of the Client.
    sessionName := "sessionName_example" // string | The name of the Transacted Session.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClientTransactedSession(context.Background(), msgVpnName, clientName, sessionName).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClientTransactedSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientTransactedSession`: MsgVpnClientTransactedSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClientTransactedSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 
**sessionName** | **string** | The name of the Transacted Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientTransactedSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTransactedSessionResponse**](MsgVpnClientTransactedSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientTransactedSessions

> MsgVpnClientTransactedSessionsResponse GetMsgVpnClientTransactedSessions(ctx, msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Transacted Session objects.



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
    clientName := "clientName_example" // string | The name of the Client.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClientTransactedSessions(context.Background(), msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClientTransactedSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientTransactedSessions`: MsgVpnClientTransactedSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClientTransactedSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientTransactedSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTransactedSessionsResponse**](MsgVpnClientTransactedSessionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientTxFlow

> MsgVpnClientTxFlowResponse GetMsgVpnClientTxFlow(ctx, msgVpnName, clientName, flowId).Select_(select_).Execute()

Get a Client Transmit Flow object.



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
    clientName := "clientName_example" // string | The name of the Client.
    flowId := "flowId_example" // string | The identifier (ID) of the flow.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClientTxFlow(context.Background(), msgVpnName, clientName, flowId).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClientTxFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientTxFlow`: MsgVpnClientTxFlowResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClientTxFlow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 
**flowId** | **string** | The identifier (ID) of the flow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientTxFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTxFlowResponse**](MsgVpnClientTxFlowResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClientTxFlows

> MsgVpnClientTxFlowsResponse GetMsgVpnClientTxFlows(ctx, msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client Transmit Flow objects.



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
    clientName := "clientName_example" // string | The name of the Client.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClientApi.GetMsgVpnClientTxFlows(context.Background(), msgVpnName, clientName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClientTxFlows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClientTxFlows`: MsgVpnClientTxFlowsResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClientTxFlows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**clientName** | **string** | The name of the Client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientTxFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientTxFlowsResponse**](MsgVpnClientTxFlowsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnClients

> MsgVpnClientsResponse GetMsgVpnClients(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Client objects.



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
    resp, r, err := api_client.ClientApi.GetMsgVpnClients(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientApi.GetMsgVpnClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnClients`: MsgVpnClientsResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientApi.GetMsgVpnClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnClientsResponse**](MsgVpnClientsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

