# \MqttSessionApi

All URIs are relative to *http://www.solace.com/SEMP/v2/monitor*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsgVpnMqttSession**](MqttSessionApi.md#GetMsgVpnMqttSession) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter} | Get an MQTT Session object.
[**GetMsgVpnMqttSessionSubscription**](MqttSessionApi.md#GetMsgVpnMqttSessionSubscription) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic} | Get a Subscription object.
[**GetMsgVpnMqttSessionSubscriptions**](MqttSessionApi.md#GetMsgVpnMqttSessionSubscriptions) | **Get** /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions | Get a list of Subscription objects.
[**GetMsgVpnMqttSessions**](MqttSessionApi.md#GetMsgVpnMqttSessions) | **Get** /msgVpns/{msgVpnName}/mqttSessions | Get a list of MQTT Session objects.



## GetMsgVpnMqttSession

> MsgVpnMqttSessionResponse GetMsgVpnMqttSession(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Select_(select_).Execute()

Get an MQTT Session object.



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
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MqttSessionApi.GetMsgVpnMqttSession(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttSessionApi.GetMsgVpnMqttSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSession`: MsgVpnMqttSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `MqttSessionApi.GetMsgVpnMqttSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnMqttSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionResponse**](MsgVpnMqttSessionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnMqttSessionSubscription

> MsgVpnMqttSessionSubscriptionResponse GetMsgVpnMqttSessionSubscription(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).Select_(select_).Execute()

Get a Subscription object.



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
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    subscriptionTopic := "subscriptionTopic_example" // string | The MQTT subscription topic.
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MqttSessionApi.GetMsgVpnMqttSessionSubscription(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter, subscriptionTopic).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttSessionApi.GetMsgVpnMqttSessionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSessionSubscription`: MsgVpnMqttSessionSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `MqttSessionApi.GetMsgVpnMqttSessionSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 
**subscriptionTopic** | **string** | The MQTT subscription topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnMqttSessionSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionResponse**](MsgVpnMqttSessionSubscriptionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnMqttSessionSubscriptions

> MsgVpnMqttSessionSubscriptionsResponse GetMsgVpnMqttSessionSubscriptions(ctx, msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of Subscription objects.



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
    mqttSessionClientId := "mqttSessionClientId_example" // string | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.
    mqttSessionVirtualRouter := "mqttSessionVirtualRouter_example" // string | The virtual router of the MQTT Session.
    count := int32(56) // int32 | Limit the count of objects in the response. See the documentation for the `count` parameter. (optional) (default to 10)
    cursor := "cursor_example" // string | The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter. (optional)
    where := []string{"Inner_example"} // []string | Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter. (optional)
    select_ := []string{"Inner_example"} // []string | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MqttSessionApi.GetMsgVpnMqttSessionSubscriptions(context.Background(), msgVpnName, mqttSessionClientId, mqttSessionVirtualRouter).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttSessionApi.GetMsgVpnMqttSessionSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSessionSubscriptions`: MsgVpnMqttSessionSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MqttSessionApi.GetMsgVpnMqttSessionSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 
**mqttSessionClientId** | **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | 
**mqttSessionVirtualRouter** | **string** | The virtual router of the MQTT Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnMqttSessionSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionSubscriptionsResponse**](MsgVpnMqttSessionSubscriptionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsgVpnMqttSessions

> MsgVpnMqttSessionsResponse GetMsgVpnMqttSessions(ctx, msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()

Get a list of MQTT Session objects.



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
    resp, r, err := api_client.MqttSessionApi.GetMsgVpnMqttSessions(context.Background(), msgVpnName).Count(count).Cursor(cursor).Where(where).Select_(select_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttSessionApi.GetMsgVpnMqttSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsgVpnMqttSessions`: MsgVpnMqttSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MqttSessionApi.GetMsgVpnMqttSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgVpnName** | **string** | The name of the Message VPN. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsgVpnMqttSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Limit the count of objects in the response. See the documentation for the &#x60;count&#x60; parameter. | [default to 10]
 **cursor** | **string** | The cursor, or position, for the next page of objects. See the documentation for the &#x60;cursor&#x60; parameter. | 
 **where** | **[]string** | Include in the response only objects where certain conditions are true. See the the documentation for the &#x60;where&#x60; parameter. | 
 **select_** | **[]string** | Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the &#x60;select&#x60; parameter. | 

### Return type

[**MsgVpnMqttSessionsResponse**](MsgVpnMqttSessionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

